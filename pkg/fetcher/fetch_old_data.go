package fetcher

import (
	"database/sql"
	"time"

	"trading-data-harvester/common"
	"trading-data-harvester/pkg/binance"
	"trading-data-harvester/pkg/convert/body2db"
	"trading-data-harvester/pkg/datamodel"
	"trading-data-harvester/pkg/tradingdb/writer"
)

// FetchOldCandles 获取所有蜡烛数据
func FetchOldCandles(startTime time.Time, exchange string, symbol string, db *sql.DB) error {
	// 倒序获取所有蜡烛数据，直到返回值为空
	nowTime := time.Now().UTC()

	limit := 1000
	query := datamodel.ApiCandlesQuery{
		Symbol: symbol,
		Limit:  &limit,
	}

	for _, interval := range common.Intervals {
		query.Interval = interval

		// 根据interval计算时间间隔的毫秒数
		var intervalMs int64
		switch interval {
		case "1m":
			intervalMs = 60 * 1000
		case "3m":
			intervalMs = 3 * 60 * 1000
		case "5m":
			intervalMs = 5 * 60 * 1000
		case "15m":
			intervalMs = 15 * 60 * 1000
		case "30m":
			intervalMs = 30 * 60 * 1000
		case "1h":
			intervalMs = 60 * 60 * 1000
		case "2h":
			intervalMs = 2 * 60 * 60 * 1000
		case "4h":
			intervalMs = 4 * 60 * 60 * 1000
		case "6h":
			intervalMs = 6 * 60 * 60 * 1000
		case "8h":
			intervalMs = 8 * 60 * 60 * 1000
		case "12h":
			intervalMs = 12 * 60 * 60 * 1000
		case "1d":
			intervalMs = 24 * 60 * 60 * 1000
		case "3d":
			intervalMs = 3 * 24 * 60 * 60 * 1000
		case "1w":
			intervalMs = 7 * 24 * 60 * 60 * 1000
		case "1M":
			intervalMs = 30 * 24 * 60 * 60 * 1000
		}

		startTime := nowTime.UnixMilli() - int64(limit)*intervalMs
		endTime := nowTime.UnixMilli() - int64(limit)*intervalMs

		for {
			// 计算startTime，往前移动limit * intervalMs的时间
			startTime = endTime - int64(limit)*intervalMs

			// 如果startTime小于0，设置为0
			if startTime < 0 {
				startTime = 0
			}

			query.StartTime = &startTime
			query.EndTime = &endTime

			// 获取数据
			rawCandles, err := binance.GetCandlestickData(query)
			if err != nil {
				return err
			}

			// 如果返回的数据为空，说明已经获取完所有数据
			if len(rawCandles) == 0 {
				break
			}

			// 将rawCandles转换为CandleDBModel
			candles := body2db.CandlesBody2DB(exchange, query, rawCandles)

			err = writer.WriteCandles(db, candles)
			if err != nil {
				return err
			}

			// 更新endTime为当前批次的startTime，继续获取更早的数据
			endTime = startTime - 1
			// 如果startTime已经是0，说明已经获取到最早的数据
			if startTime == 0 {
				break
			}
		}

	}

	return nil
}
