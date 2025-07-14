package main

import (
	"fmt"
	"log"
	"trading-data-harvester/pkg/binance"
	"trading-data-harvester/pkg/convert/body2db"
	"trading-data-harvester/pkg/datamodel"
	"trading-data-harvester/pkg/tradingdb/connect"
	"trading-data-harvester/pkg/tradingdb/writer"
)

func main() {
	db, err := connect.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	limit := 1
	query := datamodel.ApiCandlesQuery{
		Symbol:   "BTCUSDT",
		Interval: "1m",
		Limit:    &limit,
	}

	rawCandles, err := binance.GetCandlestickData(query)
	if err != nil {
		log.Fatal(err)
	}

	candles := body2db.CandlesBody2DB("binance", query, rawCandles)

	fmt.Println(candles)

	err = writer.WriteCandles(db, candles)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Candles written successfully")
}
