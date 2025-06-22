package writer

import (
	"database/sql"
	"trading-data-harvester/pkg/datamodel"
)

func WriteCandles(db *sql.DB, candles []datamodel.CandleDBModel) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(
		"INSERT INTO candles (exchange, symbol, interval, open_time, end_time, open, high, low, close, volume, quote_volume, trades, buy_volume, buy_quote_volume) " +
			"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, candle := range candles {
		_, err := stmt.Exec(
			candle.Exchange, candle.Symbol, candle.Interval, candle.OpenTime, candle.EndTime,
			candle.Open, candle.High, candle.Low, candle.Close, candle.Volume,
			candle.QuoteVolume, candle.Trades, candle.BuyVolume, candle.BuyQuoteVolume,
		)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
