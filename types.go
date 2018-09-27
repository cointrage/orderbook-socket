package orderbook

type OrderBook struct {
	LastUpdateId int64
	Asks [][]float64
	Bids [][]float64
}
