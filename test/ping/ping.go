package main

import (
	"fmt"
	"log"
	"trading-data-harvester/pkg/tradingdb/connect"
)

func main() {
	db, err := connect.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to ClickHouse: %v", err)
	}
	defer db.Close()

	// ping clickhouse
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping ClickHouse: %v", err)
	}

	fmt.Println("Connected to ClickHouse successfully")
}
