package orderbook

import (
	"fmt"
	"sort"
)

type byPriceAsc []*Order

func (x byPriceAsc) Len() int { return len(x) }
func (x byPriceAsc) Less(i, j int) bool { return x[i].Price < x[j].Price }
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
		Asks: make([]*Order, 0),
		Bids: make([]*Order, 0),
	}

	for _, ask := range (*book).Asks {
		cp.Asks = append(cp.Asks, &Order{Price: ask.Price, Quantity: ask.Quantity})
	}

	for _, bid := range (*book).Bids {
		cp.Bids = append(cp.Bids, &Order{Price: bid.Price, Quantity: bid.Quantity})
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
		LastUpdateId: (*o2).LastUpdateId,
	}

	// building asks diff
	asks := make([]*Order, 0)
	for _, ask1 := range (*o1).Asks {

		found := false
		for _, ask2 := range (*o2).Asks {
			if ask1.Price == ask2.Price {
				found = true
				break
			}
		}

		if !found {
			// removing the price level
			asks = append(asks, &Order{Price: ask1.Price, Quantity: 0})
		}
	}

	for _, ask2 := range (*o2).Asks {

		found := false
		for _, ask1 := range (*o1).Asks {
			if ask2.Price == ask1.Price && ask2.Quantity != ask1.Quantity {
				// updating price level
				asks = append(asks, &Order{Price: ask2.Price, Quantity: ask2.Quantity})
				break
			}
		}

		if !found {
			// adding price level
			asks = append(asks, &Order{Price: ask2.Price, Quantity: ask2.Quantity})
		}
	}

	// sorting by price level
	sort.Sort(byPriceAsc(asks))
	diff.Asks = asks

	// building bids diff
	bids := make([]*Order, 0)
	for _, bid1 := range (*o1).Bids {

		found := false
		for _, bid2 := range (*o2).Bids {
			if bid1.Price == bid2.Price {
				found = true
				break
			}
		}

		if !found {
			// removing the price level
			bids = append(bids, &Order{Price: bid1.Price, Quantity: 0})
		}
	}

	for _, bid2 := range (*o2).Bids {

		found := false
		for _, bid1 := range (*o1).Bids {
			if bid2.Price == bid1.Price && bid2.Quantity != bid1.Quantity {
				// updating price level
				bids = append(bids, &Order{Price: bid2.Price, Quantity: bid2.Quantity})
				break
			}
		}

		if !found {
			// adding price level
			bids = append(bids, &Order{Price: bid2.Price, Quantity: bid2.Quantity})
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
		if len((*book).Asks) == 0 || (*book).Asks[len((*book).Asks) - 1].Price < ask.Price {
			// inserting at the end
			(*book).Asks = append((*book).Asks, &Order{Price: ask.Price, Quantity: ask.Quantity})
			continue
		}

		for i := 0; i < len((*book).Asks); i += 1 {
			oask := (*book).Asks[i] 
			if oask.Price < ask.Price {
				continue
			}

			if oask.Price == ask.Price {
				// updating or deleting a price level
				if ask.Quantity == 0 {
					// removing price level
					(*book).Asks = append((*book).Asks[:i], (*book).Asks[i+1:]...)
				} else {
					// updating quantity
					oask.Quantity = ask.Quantity
				}
			} else {

				if ask.Quantity == 0 {
					// nothing to do
					continue
				}

				// inserting a new price level at (i)-th position
				(*book).Asks = append((*book).Asks, nil)
				copy((*book).Asks[i+1:], (*book).Asks[i:])
				(*book).Asks[i] = &Order{Price: ask.Price, Quantity: ask.Quantity}
			}

			break
		}
	}

	// updating bids
	for _, bid := range (*diff).Bids {
		if len((*book).Bids) == 0 || (*book).Bids[len((*book).Bids) - 1].Price > bid.Price {
			// inserting at the end
			(*book).Bids = append((*book).Bids, &Order{Price: bid.Price, Quantity: bid.Quantity})
			continue
		}

		for i := 0; i < len((*book).Bids); i += 1 {
			obid := (*book).Bids[i] 
			if obid.Price > bid.Price {
				continue
			}

			if obid.Price == bid.Price {
				// updating or deleting a price level
				if bid.Quantity == 0 {
					// removing price level
					(*book).Bids = append((*book).Bids[:i], (*book).Bids[i+1:]...)
				} else {
					// updating quantity
					obid.Quantity = bid.Quantity
				}
			} else {

				if bid.Quantity == 0 {
					// nothing to do
					continue
				}

				// inserting a new price level at (i)-th position
				(*book).Bids = append((*book).Bids, nil)
				copy((*book).Bids[i+1:], (*book).Bids[i:])
				(*book).Bids[i] = &Order{Price: bid.Price, Quantity: bid.Quantity}
			}

			break
		}
	}

	(*book).LastUpdateId = (*diff).LastUpdateId

	return nil
}

func Print(orderbook *OrderBook) {

	fmt.Printf("\033[0;0H\n")

	ln := 5
	if len(orderbook.Asks) < 5 {
		ln = len(orderbook.Asks)
	}

	for i := ln - 1; i >= 0; i -= 1 {
		fmt.Printf("%0.10f \t %0.10f\n", orderbook.Asks[i].Price, orderbook.Asks[i].Quantity)
	}

	fmt.Println("------------------------------")

	ln = 5
	if len(orderbook.Bids) < 5 {
		ln = len(orderbook.Asks)
	}
	
	for i := 0; i < ln; i += 1 {
		fmt.Printf("%0.10f \t %0.10f\n", orderbook.Bids[i].Price, orderbook.Bids[i].Quantity)
	}
}