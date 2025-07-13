package body2db

import (
	"trading-data-harvester/pkg/datamodel"
	utils "trading-data-harvester/utils/convert"
)

func CandlesBody2DB(exchange string, query datamodel.ApiCandlesQuery, candles []datamodel.ApiCandlesResponse) []datamodel.CandleDBModel {
	dbCandles := make([]datamodel.CandleDBModel, len(candles))
	for i, candle := range candles {
		dbCandles[i] = datamodel.CandleDBModel{
			Exchange:       exchange,
			Symbol:         query.Symbol,
			Interval:       query.Interval,
			OpenTime:       candle.OpenTime,
			EndTime:        candle.CloseTime,
			Open:           utils.StringToFloat64(candle.Open),
			High:           utils.StringToFloat64(candle.High),
			Low:            utils.StringToFloat64(candle.Low),
			Close:          utils.StringToFloat64(candle.Close),
			Volume:         utils.StringToFloat64(candle.Volume),
			QuoteVolume:    utils.StringToFloat64(candle.QuoteAssetVolume),
			Trades:         uint64(candle.NumberOfTrades),
			BuyVolume:      utils.StringToFloat64(candle.TakerBuyBaseAssetVolume),
			BuyQuoteVolume: utils.StringToFloat64(candle.TakerBuyQuoteAssetVolume),
			TimeStamp:      utils.Unix2Time(int64(candle.OpenTime)),
		}
	}
	return dbCandles
}
