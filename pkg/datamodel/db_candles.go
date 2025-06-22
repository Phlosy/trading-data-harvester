package datamodel

import "time"

type CandleDBModel struct {
	Exchange       string    // 交易所名称（如 binance）
	Symbol         string    // 币对（如 BTC/USDT）
	Interval       string    // K线周期（如 1m, 5m）
	OpenTime       time.Time // 开始时间，使用 Unix 时间戳
	EndTime        time.Time // 结束时间，使用 Unix 时间戳
	Open           float64   // 开盘价
	High           float64   // 最高价
	Low            float64   // 最低价
	Close          float64   // 收盘价
	Volume         float64   // 成交量
	QuoteVolume    float64   // 成交额
	Trades         uint64    // 成交笔数
	BuyVolume      float64   // 主动买入量
	BuyQuoteVolume float64   // 主动买入额
}
