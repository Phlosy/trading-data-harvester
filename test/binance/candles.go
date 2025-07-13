package main

import (
	"fmt"
	"trading-data-harvester/pkg/binance"
	"trading-data-harvester/pkg/datamodel"
)

func main() {
	limit := 1
	query := datamodel.ApiCandlesQuery{
		Symbol:   "BTCUSDT",
		Interval: "3d",
		Limit:    &limit,
	}

	candles, err := binance.GetCandlestickData(query)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(candles)
}
