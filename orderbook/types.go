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
	USDC
)

func ParseMarket(market string) (MarketName) {
	switch market {
	case "BTC":
		return BTC
	case "ETH":
		return ETH
	case "USDT":
		return USDT
	case "TUSD":
		return TUSD
	case "USD":
		return USD
	case "USDC":
		return USDC
	default:
		panic(fmt.Errorf("could not parse market: %s", market))
	}
}

func (m MarketName) String() string {
	return [...]string{"BTC", "ETH", "USDT", "TUSD", "USD", "USDC"}[m]
}

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
