package main

import (
	"log"
	"trading-data-harvester/pkg/tradingdb/connect"
	"trading-data-harvester/pkg/tradingdb/truncate"
)

func main() {
	db, err := connect.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	truncate.TruncateSingleTable(db, "crypto", "candles")
}
