package main

import (
	"fmt"
	"trading-data-harvester/pkg/binance"
	"trading-data-harvester/pkg/datamodel"
)

func main() {
	limit := 5
	query := datamodel.ApiCandlesQuery{
		Symbol:   "BTCUSDT",
		Interval: "1m",
		Limit:    &limit,
	}

	candles, err := binance.GetCandlestickData(query)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(candles)
}
