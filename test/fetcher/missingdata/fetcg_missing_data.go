package main

import (
	"fmt"
	"log"
	"trading-data-harvester/pkg/fetcher"
	"trading-data-harvester/pkg/tradingdb/connect"
	"trading-data-harvester/pkg/validator"
)

func main() {
	db, err := connect.Connect()
	if err != nil {
		log.Fatal(err)
	}

	for {
		var num int = 0
		gaps, err := validator.ValidateTimeIntegrity(db, "crypto", "candles", "binance", "1m", "BTCUSDT")
		if err != nil {
			log.Fatal(err)
		}

		if len(*gaps) == 0 {
			break
		}

		// TODO: 删除
		fmt.Println(gaps)

		err = fetcher.FetchMissingData(db, gaps)
		if err != nil {
			log.Fatal(err)
		}

		num++
		fmt.Printf("第 %d 次获取数据\n", num)

	}
}
