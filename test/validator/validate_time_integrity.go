package main

import (
	"fmt"
	"log"
	"time"
	"trading-data-harvester/pkg/tradingdb/connect"
	"trading-data-harvester/pkg/validator"
)

func main() {
	db, err := connect.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	intervalMs := int64(60 * 1000) // 1分钟
	gaps, err := validator.ValidateTimeIntegrity(db, "kline_db", "btcusdt_1m", intervalMs)
	if err != nil {
		log.Fatalf("校验失败: %v", err)
	}
	if len(gaps) == 0 {
		fmt.Println("数据完整，无缺口")
	} else {
		fmt.Println("检测到缺失时间段：")
		for _, gap := range gaps {
			fmt.Printf("缺失 %s ~ %s\n", time.UnixMilli(gap.MissingFrom), time.UnixMilli(gap.MissingTo))
		}
	}

}
