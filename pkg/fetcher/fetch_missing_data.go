package fetcher

import (
	"database/sql"
	"fmt"
	"time"
	"trading-data-harvester/pkg/binance"
	"trading-data-harvester/pkg/datamodel"
	"trading-data-harvester/pkg/tradingdb/writer"
	utils "trading-data-harvester/utils/convert"
)

func FetchMissingData(db *sql.DB, gaps *[]datamodel.ValidatorTimeGap) error {
	var missingCandles *[]datamodel.CandleDBModel = &[]datamodel.CandleDBModel{}

	length := len(*gaps)

	// 遍历gaps，调用api获取数据
	for i, gap := range *gaps {
		// 调用api获取数据
		limit := 1000

		endTime := gap.MissingTo - 1

		candles, err := binance.GetCandlestickData(datamodel.ApiCandlesQuery{
			Symbol:    gap.Symbol,
			Interval:  gap.Interval,
			StartTime: &gap.MissingFrom,
			EndTime:   &endTime,
			Limit:     &limit,
		})

		fmt.Printf("进度: %d/%d\n", i+1, length)

		if err != nil {
			return err
		}
		for _, candle := range candles {
			*missingCandles = append(*missingCandles, datamodel.CandleDBModel{
				Exchange:       gap.Exchange,
				Symbol:         gap.Symbol,
				Interval:       gap.Interval,
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
				TimeStamp:      time.UnixMilli(int64(candle.CloseTime)),
			})
		}
	}

	err := writer.WriteCandles(db, *missingCandles)
	if err != nil {
		return err
	}

	return nil
}
