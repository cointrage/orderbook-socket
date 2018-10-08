package orderbook

import (
	"fmt"
	"time"
)

type ExchangeName uint8
const (
	Idex = iota
	EtherDelta
	ForkDelta
	Binance
	HitBTC
	Mercatox
	Huobi
	Hadax
	Bitfinex
	Ethfinex
	Coinex
	Kucoin
	Okex
	Bittrex
	Bibox
	Bilaxy
	Tidex
	Ddex
	Coinegg
	Lbank
	BigOne
	Coinbene
	Liqui
	Cobinhood
	Gate
	Bcex
	Kraken
)

type MarketName uint8
const (
	BTC = iota
	ETH
	USDT
	TUSD
	USD
)

type OrderBook struct {
	LastUpdateId int64
	Exchange ExchangeName
	Market MarketName
	Ticker string
	Asks [][]float32
	Bids [][]float32
}

type OrderBookDiff struct {
	FirstUpdateId int64
	OrderBook
}

type MessageType uint8
const (
    MessageBook MessageType = iota
    MessageDiff
    MessageSubscribe
)

type Message struct {
	Ts time.Time
	Type MessageType
	Data string
}

func ParseMarket(market string) (MarketName, error) {
	switch market {
	case "BTC":
		return BTC, nil
	case "ETH":
		return ETH, nil
	case "USDT":
		return USDT, nil
	case "TUSD":
		return TUSD, nil
	case "USD":
		return USD, nil
	default:
		return 0, fmt.Errorf("could not parse market: %s", market)
	}
}
