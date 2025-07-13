package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"trading-data-harvester/pkg/fetcher"
	"trading-data-harvester/pkg/tradingdb/connect"
)

func main() {
	// 设置日志输出到文件
	logFile, err := os.OpenFile("fetch_all_data.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("无法创建日志文件:", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	startTime := time.Now()
	log.Printf("开始获取数据，开始时间: %s", startTime.Format("2006-01-02 15:04:05"))

	db, err := connect.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = fetcher.FetchAllCandles("binance", "BTCUSDT", db)
	if err != nil {
		log.Fatal(err)
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	log.Printf("数据获取完成，结束时间: %s", endTime.Format("2006-01-02 15:04:05"))
	log.Printf("总耗时: %v", duration)

	fmt.Println("Candles written successfully")
}
