package orderbook

import (
	"fmt"
	"sort"
)

type byPriceAsc [][]float64

func (x byPriceAsc) Len() int { return len(x) }
func (x byPriceAsc) Less(i, j int) bool { return x[i][0] < x[j][0] }
func (x byPriceAsc) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

// makes compact representation of differences between two orderbooks
func MakeDiff(o1 *OrderBook, o2 *OrderBook2) (*OrderBookDiff, error) {
	if o1 == nil || o2 == nil {
		return nil, fmt.Errorf("orderbooks should not be nil")
	}

	diff := &OrderBookDiff{
		FirstUpdateId: (*o1).LastUpdateId,
		LastUpdateId: (*o2).LastUpdateI	d,
	}

	// building asks diff
	asks := make([][]float64)
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
			asks = append(asks, []float64{ask1[0], 0})
		}
	}

	for _, ask2 := range (*o2).Asks {

		found := false
		for _, ask1 := range (*o1).Asks {
			if ask2[0] == ask1[0] && ask2[1] != ask1[1] {
				// updating price level
				asks = append(asks, []float64{ask2[0], ask2[1]})
				break
			}
		}

		if !found {
			// adding price level
			asks = append(asks, []float64{ask2[0], ask2[1]})
		}
	}

	// sorting by price level
	sort.Sort(byPriceAsc(asks))
	diff.Asks = asks

	// building bids diff
	bids := make([][]float64)
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
			bids = append(bids, []float64{bid1[0], 0})
		}
	}

	for _, bid2 := range (*o2).Bids {

		found := false
		for _, bid1 := range (*o1).Bids {
			if bid2[0] == bid1[0] && bid2[1] != bid1[1] {
				// updating price level
				bids = append(bids, []float64{bid2[0], bid2[1]})
				break
			}
		}

		if !found {
			// adding price level
			bids = append(bids, []float64{bid2[0], bid2[1]})
		}
	}

	// sorting by price level
	sort.Sort(sort.Reverse(byPriceAsc(bids)))
	diff.Bids = bids

	return diff, nil
}