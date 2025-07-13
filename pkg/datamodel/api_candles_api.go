package datamodel

type ApiCandlesQuery struct {
	Symbol    string `json:"symbol"`    // 币对
	Interval  string `json:"interval"`  // 时间间隔
	StartTime *int64 `json:"startTime"` // 开始时间
	EndTime   *int64 `json:"endTime"`   // 结束时间
	Limit     *int   `json:"limit"`     // 返回结果限制数量
}

type ApiCandlesResponse struct {
	OpenTime                 uint64 `json:"openTime"`                 // 开盘时间
	Open                     string `json:"open"`                     // 开盘价
	High                     string `json:"high"`                     // 最高价
	Low                      string `json:"low"`                      // 最低价
	Close                    string `json:"close"`                    // 收盘价
	Volume                   string `json:"volume"`                   // 成交量
	CloseTime                uint64 `json:"closeTime"`                // 收盘时间
	QuoteAssetVolume         string `json:"quoteAssetVolume"`         // 报价资产成交量
	NumberOfTrades           int64  `json:"numberOfTrades"`           // 交易次数
	TakerBuyBaseAssetVolume  string `json:"takerBuyBaseAssetVolume"`  // 买入成交量
	TakerBuyQuoteAssetVolume string `json:"takerBuyQuoteAssetVolume"` // 买入报价资产成交量
	Ignore                   string `json:"ignore"`                   // 忽略
}
