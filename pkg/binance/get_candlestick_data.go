package binance

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"trading-data-harvester/pkg/datamodel"
)

func GetCandlestickData(query datamodel.ApiCandlesQuery) ([]datamodel.ApiCandlesResponse, error) {
	url := fmt.Sprintf("%s%s?%s", BaseURL, GetCandlestickDataPath, BuildQuery(query))

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var rawCandles [][]interface{}
	err = json.Unmarshal(body, &rawCandles)
	if err != nil {
		return nil, err
	}

	var candles []datamodel.ApiCandlesResponse
	for _, rawCandle := range rawCandles {
		candle := datamodel.ApiCandlesResponse{
			OpenTime:                 int64(rawCandle[0].(float64)),
			Open:                     rawCandle[1].(string),
			High:                     rawCandle[2].(string),
			Low:                      rawCandle[3].(string),
			Close:                    rawCandle[4].(string),
			Volume:                   rawCandle[5].(string),
			CloseTime:                int64(rawCandle[6].(float64)),
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

func BuildQuery(query datamodel.ApiCandlesQuery) string {
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
