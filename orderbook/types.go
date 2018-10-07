package orderbook

import (
	"time"
)

type OrderBook struct {
	LastUpdateId int64
	Exchange, Market, Ticker string
	Asks [][]float64
	Bids [][]float64
}

type OrderBookDiff struct {
	FirstUpdateId int64
	OrderBook
}

type Message struct {
	Ts time.Time
	Type string
	Data string
}
