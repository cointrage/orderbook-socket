// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orderbook/orderbook.proto

package orderbook

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Exchange int32

const (
	Exchange_UnknownExchange Exchange = 0
	Exchange_Idex            Exchange = 1
	Exchange_EtherDelta      Exchange = 2
	Exchange_ForkDelta       Exchange = 3
	Exchange_Binance         Exchange = 4
	Exchange_HitBTC          Exchange = 5
	Exchange_Mercatox        Exchange = 6
	Exchange_Huobi           Exchange = 7
	Exchange_Hadax           Exchange = 8
	Exchange_Bitfinex        Exchange = 9
	Exchange_Ethfinex        Exchange = 10
	Exchange_Coinex          Exchange = 11
	Exchange_Kucoin          Exchange = 12
	Exchange_Okex            Exchange = 13
	Exchange_Bittrex         Exchange = 14
	Exchange_Bibox           Exchange = 15
	Exchange_Bilaxy          Exchange = 16
	Exchange_Tidex           Exchange = 17
	Exchange_Ddex            Exchange = 18
	Exchange_Coinegg         Exchange = 19
	Exchange_Lbank           Exchange = 20
	Exchange_BigOne          Exchange = 21
	Exchange_Coinbene        Exchange = 22
	Exchange_Liqui           Exchange = 23
	Exchange_Cobinhood       Exchange = 24
	Exchange_Gate            Exchange = 25
	Exchange_Bcex            Exchange = 26
	Exchange_Kraken          Exchange = 27
	Exchange_Liquid          Exchange = 28
	Exchange_Coinbase        Exchange = 29
	Exchange_Poloniex        Exchange = 30
	Exchange_Bitz            Exchange = 31
	Exchange_Switcheo        Exchange = 32
	Exchange_Idax            Exchange = 33
	Exchange_Tokenstore      Exchange = 34
	Exchange_Bithumb         Exchange = 35
	Exchange_BithumbDex      Exchange = 36
)

var Exchange_name = map[int32]string{
	0:  "UnknownExchange",
	1:  "Idex",
	2:  "EtherDelta",
	3:  "ForkDelta",
	4:  "Binance",
	5:  "HitBTC",
	6:  "Mercatox",
	7:  "Huobi",
	8:  "Hadax",
	9:  "Bitfinex",
	10: "Ethfinex",
	11: "Coinex",
	12: "Kucoin",
	13: "Okex",
	14: "Bittrex",
	15: "Bibox",
	16: "Bilaxy",
	17: "Tidex",
	18: "Ddex",
	19: "Coinegg",
	20: "Lbank",
	21: "BigOne",
	22: "Coinbene",
	23: "Liqui",
	24: "Cobinhood",
	25: "Gate",
	26: "Bcex",
	27: "Kraken",
	28: "Liquid",
	29: "Coinbase",
	30: "Poloniex",
	31: "Bitz",
	32: "Switcheo",
	33: "Idax",
	34: "Tokenstore",
	35: "Bithumb",
	36: "BithumbDex",
}

var Exchange_value = map[string]int32{
	"UnknownExchange": 0,
	"Idex":            1,
	"EtherDelta":      2,
	"ForkDelta":       3,
	"Binance":         4,
	"HitBTC":          5,
	"Mercatox":        6,
	"Huobi":           7,
	"Hadax":           8,
	"Bitfinex":        9,
	"Ethfinex":        10,
	"Coinex":          11,
	"Kucoin":          12,
	"Okex":            13,
	"Bittrex":         14,
	"Bibox":           15,
	"Bilaxy":          16,
	"Tidex":           17,
	"Ddex":            18,
	"Coinegg":         19,
	"Lbank":           20,
	"BigOne":          21,
	"Coinbene":        22,
	"Liqui":           23,
	"Cobinhood":       24,
	"Gate":            25,
	"Bcex":            26,
	"Kraken":          27,
	"Liquid":          28,
	"Coinbase":        29,
	"Poloniex":        30,
	"Bitz":            31,
	"Switcheo":        32,
	"Idax":            33,
	"Tokenstore":      34,
	"Bithumb":         35,
	"BithumbDex":      36,
}

func (x Exchange) String() string {
	return proto.EnumName(Exchange_name, int32(x))
}

func (Exchange) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_233e91af0692daf8, []int{0}
}

type Market int32

const (
	Market_UnknownMarket Market = 0
	Market_BTC           Market = 1
	Market_ETH           Market = 2
	Market_USDT          Market = 3
	Market_TUSD          Market = 4
	Market_USD           Market = 5
	Market_USDC          Market = 6
	Market_GUSD          Market = 7
	Market_BNB           Market = 8
	Market_BIX           Market = 9
	Market_DAI           Market = 10
	Market_BCH           Market = 11
	Market_CET           Market = 12
	Market_QTUM          Market = 13
	Market_NEO           Market = 14
	Market_KCS           Market = 15
	Market_EOS           Market = 16
	Market_EUR           Market = 17
	Market_GBP           Market = 18
	Market_JPY           Market = 19
	Market_COB           Market = 20
	Market_PAX           Market = 21
	Market_GAS           Market = 22
	Market_SWTH          Market = 23
	Market_KRW           Market = 24
)

var Market_name = map[int32]string{
	0:  "UnknownMarket",
	1:  "BTC",
	2:  "ETH",
	3:  "USDT",
	4:  "TUSD",
	5:  "USD",
	6:  "USDC",
	7:  "GUSD",
	8:  "BNB",
	9:  "BIX",
	10: "DAI",
	11: "BCH",
	12: "CET",
	13: "QTUM",
	14: "NEO",
	15: "KCS",
	16: "EOS",
	17: "EUR",
	18: "GBP",
	19: "JPY",
	20: "COB",
	21: "PAX",
	22: "GAS",
	23: "SWTH",
	24: "KRW",
}

var Market_value = map[string]int32{
	"UnknownMarket": 0,
	"BTC":           1,
	"ETH":           2,
	"USDT":          3,
	"TUSD":          4,
	"USD":           5,
	"USDC":          6,
	"GUSD":          7,
	"BNB":           8,
	"BIX":           9,
	"DAI":           10,
	"BCH":           11,
	"CET":           12,
	"QTUM":          13,
	"NEO":           14,
	"KCS":           15,
	"EOS":           16,
	"EUR":           17,
	"GBP":           18,
	"JPY":           19,
	"COB":           20,
	"PAX":           21,
	"GAS":           22,
	"SWTH":          23,
	"KRW":           24,
}

func (x Market) String() string {
	return proto.EnumName(Market_name, int32(x))
}

func (Market) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_233e91af0692daf8, []int{1}
}

type OrderBook struct {
	LastUpdateId         int64    `protobuf:"varint,1,opt,name=LastUpdateId,proto3" json:"LastUpdateId,omitempty"`
	Exchange             Exchange `protobuf:"varint,2,opt,name=Exchange,proto3,enum=orderbook.Exchange" json:"Exchange,omitempty"`
	Market               Market   `protobuf:"varint,3,opt,name=Market,proto3,enum=orderbook.Market" json:"Market,omitempty"`
	Ticker               string   `protobuf:"bytes,4,opt,name=Ticker,proto3" json:"Ticker,omitempty"`
	Asks                 []*Order `protobuf:"bytes,5,rep,name=Asks,proto3" json:"Asks,omitempty"`
	Bids                 []*Order `protobuf:"bytes,6,rep,name=Bids,proto3" json:"Bids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderBook) Reset()         { *m = OrderBook{} }
func (m *OrderBook) String() string { return proto.CompactTextString(m) }
func (*OrderBook) ProtoMessage()    {}
func (*OrderBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_233e91af0692daf8, []int{0}
}

func (m *OrderBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderBook.Unmarshal(m, b)
}
func (m *OrderBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderBook.Marshal(b, m, deterministic)
}
func (m *OrderBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderBook.Merge(m, src)
}
func (m *OrderBook) XXX_Size() int {
	return xxx_messageInfo_OrderBook.Size(m)
}
func (m *OrderBook) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderBook.DiscardUnknown(m)
}

var xxx_messageInfo_OrderBook proto.InternalMessageInfo

func (m *OrderBook) GetLastUpdateId() int64 {
	if m != nil {
		return m.LastUpdateId
	}
	return 0
}

func (m *OrderBook) GetExchange() Exchange {
	if m != nil {
		return m.Exchange
	}
	return Exchange_UnknownExchange
}

func (m *OrderBook) GetMarket() Market {
	if m != nil {
		return m.Market
	}
	return Market_UnknownMarket
}

func (m *OrderBook) GetTicker() string {
	if m != nil {
		return m.Ticker
	}
	return ""
}

func (m *OrderBook) GetAsks() []*Order {
	if m != nil {
		return m.Asks
	}
	return nil
}

func (m *OrderBook) GetBids() []*Order {
	if m != nil {
		return m.Bids
	}
	return nil
}

type OrderBookDiff struct {
	LastUpdateId         int64    `protobuf:"varint,1,opt,name=LastUpdateId,proto3" json:"LastUpdateId,omitempty"`
	Exchange             Exchange `protobuf:"varint,2,opt,name=Exchange,proto3,enum=orderbook.Exchange" json:"Exchange,omitempty"`
	Market               Market   `protobuf:"varint,3,opt,name=Market,proto3,enum=orderbook.Market" json:"Market,omitempty"`
	Ticker               string   `protobuf:"bytes,4,opt,name=Ticker,proto3" json:"Ticker,omitempty"`
	Asks                 []*Order `protobuf:"bytes,5,rep,name=Asks,proto3" json:"Asks,omitempty"`
	Bids                 []*Order `protobuf:"bytes,6,rep,name=Bids,proto3" json:"Bids,omitempty"`
	FirstUpdateId        int64    `protobuf:"varint,7,opt,name=FirstUpdateId,proto3" json:"FirstUpdateId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderBookDiff) Reset()         { *m = OrderBookDiff{} }
func (m *OrderBookDiff) String() string { return proto.CompactTextString(m) }
func (*OrderBookDiff) ProtoMessage()    {}
func (*OrderBookDiff) Descriptor() ([]byte, []int) {
	return fileDescriptor_233e91af0692daf8, []int{1}
}

func (m *OrderBookDiff) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderBookDiff.Unmarshal(m, b)
}
func (m *OrderBookDiff) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderBookDiff.Marshal(b, m, deterministic)
}
func (m *OrderBookDiff) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderBookDiff.Merge(m, src)
}
func (m *OrderBookDiff) XXX_Size() int {
	return xxx_messageInfo_OrderBookDiff.Size(m)
}
func (m *OrderBookDiff) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderBookDiff.DiscardUnknown(m)
}

var xxx_messageInfo_OrderBookDiff proto.InternalMessageInfo

func (m *OrderBookDiff) GetLastUpdateId() int64 {
	if m != nil {
		return m.LastUpdateId
	}
	return 0
}

func (m *OrderBookDiff) GetExchange() Exchange {
	if m != nil {
		return m.Exchange
	}
	return Exchange_UnknownExchange
}

func (m *OrderBookDiff) GetMarket() Market {
	if m != nil {
		return m.Market
	}
	return Market_UnknownMarket
}

func (m *OrderBookDiff) GetTicker() string {
	if m != nil {
		return m.Ticker
	}
	return ""
}

func (m *OrderBookDiff) GetAsks() []*Order {
	if m != nil {
		return m.Asks
	}
	return nil
}

func (m *OrderBookDiff) GetBids() []*Order {
	if m != nil {
		return m.Bids
	}
	return nil
}

func (m *OrderBookDiff) GetFirstUpdateId() int64 {
	if m != nil {
		return m.FirstUpdateId
	}
	return 0
}

type Order struct {
	Price                float32  `protobuf:"fixed32,1,opt,name=Price,proto3" json:"Price,omitempty"`
	Quantity             float32  `protobuf:"fixed32,2,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_233e91af0692daf8, []int{2}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Order) GetQuantity() float32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func init() {
	proto.RegisterEnum("orderbook.Exchange", Exchange_name, Exchange_value)
	proto.RegisterEnum("orderbook.Market", Market_name, Market_value)
	proto.RegisterType((*OrderBook)(nil), "orderbook.OrderBook")
	proto.RegisterType((*OrderBookDiff)(nil), "orderbook.OrderBookDiff")
	proto.RegisterType((*Order)(nil), "orderbook.Order")
}

func init() { proto.RegisterFile("orderbook/orderbook.proto", fileDescriptor_233e91af0692daf8) }

var fileDescriptor_233e91af0692daf8 = []byte{
	// 694 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x54, 0xcb, 0x56, 0x22, 0x49,
	0x10, 0xb5, 0x78, 0x93, 0x08, 0x06, 0x89, 0x8f, 0xd2, 0x79, 0x31, 0x8c, 0x0b, 0xc6, 0x85, 0x9e,
	0xe3, 0xac, 0x66, 0x49, 0x01, 0x0a, 0xe3, 0x03, 0x2c, 0x8a, 0xa3, 0xb3, 0xcc, 0xaa, 0x4a, 0x21,
	0x4f, 0xd9, 0x99, 0x76, 0x91, 0x1c, 0xcb, 0xfe, 0x93, 0xfe, 0xba, 0xfe, 0x84, 0x5e, 0xf7, 0xae,
	0x4f, 0x64, 0x21, 0xea, 0xa2, 0xbf, 0xa0, 0x77, 0xf7, 0xde, 0xb8, 0x71, 0x2b, 0x22, 0x83, 0x03,
	0xd9, 0x57, 0x71, 0xc8, 0x63, 0x5f, 0xa9, 0xe8, 0x64, 0x8d, 0x8e, 0x1f, 0x63, 0xa5, 0x15, 0x2d,
	0xaf, 0x85, 0xd6, 0x57, 0x8b, 0x94, 0x47, 0xc8, 0x1c, 0xa5, 0x22, 0xda, 0x22, 0x9b, 0x97, 0x6c,
	0xa1, 0xa7, 0x8f, 0x21, 0xd3, 0x7c, 0x18, 0xda, 0x56, 0xd3, 0x6a, 0x67, 0xdd, 0x77, 0x1a, 0x3d,
	0x21, 0xa5, 0x7e, 0x12, 0xcc, 0x99, 0x9c, 0x71, 0x3b, 0xd3, 0xb4, 0xda, 0xb5, 0xd3, 0xc6, 0xf1,
	0xeb, 0x07, 0x5e, 0x4a, 0xee, 0xda, 0x44, 0xff, 0x26, 0x85, 0x2b, 0x16, 0x47, 0x5c, 0xdb, 0x59,
	0x63, 0xaf, 0xbf, 0xb1, 0xa7, 0x05, 0x77, 0x65, 0xa0, 0xbb, 0xa4, 0xe0, 0x89, 0x20, 0xe2, 0xb1,
	0x9d, 0x6b, 0x5a, 0xed, 0xb2, 0xbb, 0x62, 0xf4, 0x90, 0xe4, 0x3a, 0x8b, 0x68, 0x61, 0xe7, 0x9b,
	0xd9, 0x76, 0xe5, 0x14, 0xde, 0x04, 0x98, 0xd9, 0x5d, 0x53, 0x45, 0x97, 0x23, 0xc2, 0x85, 0x5d,
	0xf8, 0x91, 0x0b, 0xab, 0xad, 0xcf, 0x19, 0x52, 0x5d, 0x6f, 0xdc, 0x13, 0xf7, 0xf7, 0x3f, 0xc3,
	0xd6, 0xf4, 0x90, 0x54, 0xcf, 0x44, 0xfc, 0x66, 0xc9, 0xa2, 0x59, 0xf2, 0xbd, 0xd8, 0xfa, 0x97,
	0xe4, 0x4d, 0x13, 0xdd, 0x26, 0xf9, 0x71, 0x2c, 0x02, 0x6e, 0xde, 0x22, 0xe3, 0xa6, 0x84, 0x1e,
	0x90, 0xd2, 0xcd, 0x92, 0x49, 0x2d, 0xf4, 0xb3, 0x79, 0x84, 0x8c, 0xbb, 0xe6, 0x47, 0x5f, 0xb2,
	0xaf, 0x2f, 0x44, 0x1b, 0x64, 0x6b, 0x2a, 0x23, 0xa9, 0x9e, 0xe4, 0x8b, 0x04, 0x1b, 0xb4, 0x44,
	0x72, 0xc3, 0x90, 0x27, 0x60, 0xd1, 0x1a, 0x21, 0x7d, 0x3d, 0xe7, 0x71, 0x8f, 0x3f, 0x68, 0x06,
	0x19, 0x5a, 0x25, 0xe5, 0x33, 0x15, 0x47, 0x29, 0xcd, 0xd2, 0x0a, 0x29, 0x3a, 0x42, 0x32, 0x19,
	0x70, 0xc8, 0x51, 0x42, 0x0a, 0x03, 0xa1, 0x1d, 0xaf, 0x0b, 0x79, 0xba, 0x49, 0x4a, 0x57, 0x3c,
	0x0e, 0x98, 0x56, 0x09, 0x14, 0x68, 0x99, 0xe4, 0x07, 0x4b, 0xe5, 0x0b, 0x28, 0x1a, 0xc8, 0x42,
	0x96, 0x40, 0x09, 0x3d, 0x8e, 0xd0, 0xf7, 0x42, 0xf2, 0x04, 0xca, 0xc8, 0xfa, 0x7a, 0x9e, 0x32,
	0x82, 0x59, 0x5d, 0x65, 0x70, 0x05, 0xf1, 0xc5, 0x32, 0x50, 0x42, 0xc2, 0x26, 0x4e, 0x36, 0x8a,
	0x78, 0x02, 0xd5, 0xf4, 0xd3, 0x5a, 0xc7, 0x3c, 0x81, 0x1a, 0xa6, 0x3a, 0xc2, 0x57, 0x09, 0x6c,
	0xa1, 0xdb, 0x11, 0x0f, 0x2c, 0x79, 0x06, 0x40, 0xd9, 0x13, 0xb8, 0x48, 0x1d, 0x1b, 0x7b, 0x88,
	0x28, 0x36, 0x9a, 0xe8, 0xd9, 0x0c, 0x1a, 0xe8, 0xb8, 0xf4, 0x99, 0x8c, 0x60, 0x3b, 0x6d, 0x9c,
	0x8d, 0x24, 0x87, 0x1d, 0x1c, 0x06, 0x3d, 0x3e, 0x97, 0x1c, 0x76, 0x8d, 0x49, 0x7c, 0x5c, 0x0a,
	0xd8, 0xc3, 0xfd, 0xbb, 0xca, 0x17, 0x72, 0xae, 0x54, 0x08, 0x36, 0xa6, 0x9e, 0x33, 0xcd, 0x61,
	0x1f, 0x91, 0x13, 0xf0, 0x04, 0x0e, 0xcc, 0xb8, 0x31, 0x8b, 0xb8, 0x84, 0x5f, 0x10, 0x9b, 0xce,
	0x10, 0x7e, 0x5d, 0x67, 0xb2, 0x05, 0x87, 0xdf, 0x90, 0x8d, 0xd5, 0x83, 0x92, 0x82, 0x27, 0xf0,
	0xbb, 0xe9, 0x16, 0xfa, 0x13, 0xfc, 0x81, 0xfa, 0xe4, 0x49, 0xe8, 0x60, 0xce, 0x15, 0x34, 0xd3,
	0x43, 0xb0, 0x04, 0xfe, 0xc4, 0x43, 0x78, 0x2a, 0xe2, 0x72, 0xa1, 0x55, 0xcc, 0xa1, 0xb5, 0x5a,
	0x7f, 0xbe, 0xfc, 0xe0, 0xc3, 0x5f, 0x58, 0x5c, 0x91, 0x1e, 0x4f, 0xe0, 0xf0, 0xe8, 0x9b, 0xf5,
	0xf2, 0x93, 0xa6, 0x75, 0x52, 0x5d, 0xdd, 0x37, 0x15, 0x60, 0x83, 0x16, 0x49, 0x16, 0x8f, 0x64,
	0x21, 0xe8, 0x7b, 0x03, 0xc8, 0xe0, 0x67, 0xa6, 0x93, 0x9e, 0x07, 0x59, 0x44, 0xde, 0x74, 0xd2,
	0x83, 0x1c, 0x16, 0x11, 0xe4, 0x57, 0xc5, 0x2e, 0x14, 0xcc, 0xb6, 0xa8, 0x15, 0x4d, 0xc4, 0xb5,
	0x03, 0x25, 0x03, 0x86, 0x77, 0x50, 0x46, 0xd0, 0xeb, 0x0c, 0x81, 0x18, 0xa5, 0x3b, 0x80, 0x0a,
	0x82, 0x6e, 0xdf, 0x4b, 0x6f, 0x76, 0xe3, 0x4d, 0xaf, 0xa0, 0x8a, 0xd2, 0x75, 0x7f, 0x04, 0x35,
	0x04, 0x17, 0xdd, 0x09, 0x6c, 0x99, 0x11, 0x46, 0x13, 0x00, 0x03, 0xa6, 0x2e, 0xd4, 0x11, 0x9c,
	0x3b, 0x63, 0xa0, 0x08, 0xfe, 0x1b, 0xff, 0x0f, 0x0d, 0x13, 0x34, 0x72, 0x60, 0x1b, 0xc1, 0xb8,
	0x73, 0x07, 0x3b, 0xc6, 0xd3, 0x99, 0xc0, 0x2e, 0x46, 0x4f, 0x6e, 0xbd, 0x01, 0xec, 0x99, 0x44,
	0xf7, 0x16, 0x6c, 0xbf, 0x60, 0xfe, 0x38, 0xff, 0xf9, 0x1e, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x99,
	0x09, 0x1b, 0x55, 0x05, 0x00, 0x00,
}
