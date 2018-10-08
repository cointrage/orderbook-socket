package orderbook

import (
	"fmt"
	"sort"
)

type byPriceAsc [][]float32

func (x byPriceAsc) Len() int { return len(x) }
func (x byPriceAsc) Less(i, j int) bool { return x[i][0] < x[j][0] }
func (x byPriceAsc) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

// Copies orderbook values to new structure
func Copy(book *OrderBook) (*OrderBook) {
	
	if book == nil {
		return nil
	}

	cp := &OrderBook{
		LastUpdateId: (*book).LastUpdateId,
		Exchange: (*book).Exchange,
		Market: (*book).Market,
		Ticker: (*book).Ticker,
		Asks: make([][]float32, 0),
		Bids: make([][]float32, 0),
	}

	for _, ask := range (*book).Asks {
		cp.Asks = append(cp.Asks, []float32{ask[0], ask[1]})
	}

	for _, bid := range (*book).Bids {
		cp.Bids = append(cp.Bids, []float32{bid[0], bid[1]})
	}

	return cp
}

// Makes compact representation of differences between two orderbooks
func MakeDiff(o1 *OrderBook, o2 *OrderBook) (*OrderBookDiff, error) {
	if o1 == nil || o2 == nil {
		return nil, fmt.Errorf("orderbooks should be defined")
	}

	diff := &OrderBookDiff{
		FirstUpdateId: (*o1).LastUpdateId,
		OrderBook: OrderBook{
			LastUpdateId: (*o2).LastUpdateId,
		},
	}

	// building asks diff
	asks := make([][]float32, 0)
	for _, ask1 := range (*o1).Asks {

		found := false
		for _, ask2 := range (*o2).Asks {
			if ask1[0] == ask2[0] {
				found = true
				break
			}
		}

		if !found {
			// removing the price level
			asks = append(asks, []float32{ask1[0], 0})
		}
	}

	for _, ask2 := range (*o2).Asks {

		found := false
		for _, ask1 := range (*o1).Asks {
			if ask2[0] == ask1[0] && ask2[1] != ask1[1] {
				// updating price level
				asks = append(asks, []float32{ask2[0], ask2[1]})
				break
			}
		}

		if !found {
			// adding price level
			asks = append(asks, []float32{ask2[0], ask2[1]})
		}
	}

	// sorting by price level
	sort.Sort(byPriceAsc(asks))
	diff.Asks = asks

	// building bids diff
	bids := make([][]float32, 0)
	for _, bid1 := range (*o1).Bids {

		found := false
		for _, bid2 := range (*o2).Bids {
			if bid1[0] == bid2[0] {
				found = true
				break
			}
		}

		if !found {
			// removing the price level
			bids = append(bids, []float32{bid1[0], 0})
		}
	}

	for _, bid2 := range (*o2).Bids {

		found := false
		for _, bid1 := range (*o1).Bids {
			if bid2[0] == bid1[0] && bid2[1] != bid1[1] {
				// updating price level
				bids = append(bids, []float32{bid2[0], bid2[1]})
				break
			}
		}

		if !found {
			// adding price level
			bids = append(bids, []float32{bid2[0], bid2[1]})
		}
	}

	// sorting by price level
	sort.Sort(sort.Reverse(byPriceAsc(bids)))
	diff.Bids = bids

	return diff, nil
}

// Applies diff update to orderbook in place 
func ApplyDiff(book *OrderBook, diff *OrderBookDiff) (error) {
	if book == nil || diff == nil {
		return fmt.Errorf("both arguments should be defined")
	}

	// updating asks
	for _, ask := range (*diff).Asks {
		if len((*book).Asks) == 0 || (*book).Asks[len((*book).Asks) - 1][0] < ask[0] {
			// inserting at the end
			(*book).Asks = append((*book).Asks, []float32{ask[0], ask[1]})
			continue
		}

		for i := 0; i < len((*book).Asks); i += 1 {
			oask := (*book).Asks[i] 
			if oask[0] < ask[0] {
				continue
			}

			if oask[0] == ask[0] {
				// updating or deleting a price level
				if ask[1] == 0 {
					// removing price level
					(*book).Asks = append((*book).Asks[:i], (*book).Asks[i+1:]...)
				} else {
					// updating quantity
					oask[1] = ask[1]
				}
			} else {

				if ask[1] == 0 {
					// nothing to do
					continue
				}

				// inserting a new price level at (i)-th position
				(*book).Asks = append((*book).Asks, []float32{})
				copy((*book).Asks[i+1:], (*book).Asks[i:])
				(*book).Asks[i] = []float32{ask[0], ask[1]}
			}

			break
		}
	}

	// updating bids
	for _, bid := range (*diff).Bids {
		if len((*book).Bids) == 0 || (*book).Bids[len((*book).Bids) - 1][0] > bid[0] {
			// inserting at the end
			(*book).Bids = append((*book).Bids, []float32{bid[0], bid[1]})
			continue
		}

		for i := 0; i < len((*book).Bids); i += 1 {
			obid := (*book).Bids[i] 
			if obid[0] > bid[0] {
				continue
			}

			if obid[0] == bid[0] {
				// updating or deleting a price level
				if bid[1] == 0 {
					// removing price level
					(*book).Bids = append((*book).Bids[:i], (*book).Bids[i+1:]...)
				} else {
					// updating quantity
					obid[1] = bid[1]
				}
			} else {

				if bid[1] == 0 {
					// nothing to do
					continue
				}

				// inserting a new price level at (i)-th position
				(*book).Bids = append((*book).Bids, []float32{})
				copy((*book).Bids[i+1:], (*book).Bids[i:])
				(*book).Bids[i] = []float32{bid[0], bid[1]}
			}

			break
		}
	}

	(*book).LastUpdateId = (*diff).LastUpdateId

	return nil
}

func Print(orderbook *OrderBook) {

	fmt.Printf("\033[0;0H\n")

	for i := 4; i >= 0; i -= 1 {
		fmt.Printf("%0.10f \t %0.10f\n", orderbook.Asks[i][0], orderbook.Asks[i][1])
	}

	fmt.Println("------------------------------")

	for _, bid := range orderbook.Bids[:5] {
		fmt.Printf("%0.10f \t %0.10f\n", bid[0], bid[1])
	}
}