package fetcher

import (
	"database/sql"
	"fmt"
	"time"

	"trading-data-harvester/common"
	"trading-data-harvester/pkg/binance"
	"trading-data-harvester/pkg/convert/body2db"
	"trading-data-harvester/pkg/datamodel"
	"trading-data-harvester/pkg/tradingdb/writer"
	utils "trading-data-harvester/utils/convert"
)

// FetchAllCandles 获取所有蜡烛数据
func FetchAllCandles(exchange string, symbol string, db *sql.DB) error {
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
		intervalMs := utils.Interval2Ms(interval)

		// 如果intervalMs为0，说明interval不合法，跳过
		if intervalMs == 0 {
			continue
		}

		startTime := nowTime.UnixMilli() - intervalMs
		endTime := nowTime.UnixMilli() - intervalMs
		// TODO: 删除
		fmt.Println("nowTime:", nowTime.UnixMilli())
		fmt.Println("startTime:", startTime)
		fmt.Println("endTime:", endTime)

		for {
			// 计算startTime，往前移动limit * intervalMs的时间
			startTime = endTime - int64(limit)*intervalMs

			// 如果startTime小于0，设置为0
			if startTime < 0 {
				startTime = 0
			}
			if endTime < 0 {
				endTime = 0
			}

			query.StartTime = &startTime
			query.EndTime = &endTime

			// TODO: 删除
			tempStartTime := time.UnixMilli(startTime).Format("2006-01-02 15:04:05")
			tempEndTime := time.UnixMilli(endTime).Format("2006-01-02 15:04:05")
			fmt.Println("tempStartTime:", tempStartTime, "tempEndTime:", tempEndTime)
			fmt.Println(query)

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
