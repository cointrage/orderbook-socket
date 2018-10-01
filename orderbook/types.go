package orderbook

type OrderBook struct {
	LastUpdateId int64
	Market, Ticker string
	Asks [][]float64
	Bids [][]float64
}

type OrderBookDiff struct {
	FirstUpdateId int64
	OrderBook
}

type Message struct {
	Ts int64
	Type string
	Data string
}
