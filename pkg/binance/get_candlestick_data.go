package binance

import (
	"encoding/json"
	"fmt"
	"trading-data-harvester/pkg/datamodel"
)

// GetCandlestickData 获取K线数据
func GetCandlestickData(query datamodel.ApiCandlesQuery) ([]datamodel.ApiCandlesResponse, error) {
	url := fmt.Sprintf("%s%s?%s", BaseURL, GetCandlestickDataPath, buildQuery(query))

	headers := map[string]string{
		"User-Agent": "Go-binance-client",
		"Connection": "keep-alive",
	}

	body, err := DoGet(url, headers)
	if err != nil {
		return nil, fmt.Errorf("请求 Binance 接口失败: %w", err)
	}

	var rawCandles [][]interface{}
	if err := json.Unmarshal(body, &rawCandles); err != nil {
		return nil, fmt.Errorf("JSON 解析失败: %w, body: %s", err, string(body))
	}

	var candles []datamodel.ApiCandlesResponse
	for _, rawCandle := range rawCandles {
		candle := datamodel.ApiCandlesResponse{
			OpenTime:                 uint64(rawCandle[0].(float64)),
			Open:                     rawCandle[1].(string),
			High:                     rawCandle[2].(string),
			Low:                      rawCandle[3].(string),
			Close:                    rawCandle[4].(string),
			Volume:                   rawCandle[5].(string),
			CloseTime:                uint64(rawCandle[6].(float64)),
			QuoteAssetVolume:         rawCandle[7].(string),
			NumberOfTrades:           int64(rawCandle[8].(float64)),
			TakerBuyBaseAssetVolume:  rawCandle[9].(string),
			TakerBuyQuoteAssetVolume: rawCandle[10].(string),
			Ignore:                   rawCandle[11].(string),
		}
		candles = append(candles, candle)
	}

	return candles, nil
}

func buildQuery(query datamodel.ApiCandlesQuery) string {
	queryString := fmt.Sprintf("symbol=%s&interval=%s", query.Symbol, query.Interval)
	if query.Limit != nil {
		queryString += fmt.Sprintf("&limit=%d", *query.Limit)
	}
	if query.StartTime != nil {
		queryString += fmt.Sprintf("&startTime=%d", *query.StartTime)
	}
	if query.EndTime != nil {
		queryString += fmt.Sprintf("&endTime=%d", *query.EndTime)
	}
	return queryString
}
