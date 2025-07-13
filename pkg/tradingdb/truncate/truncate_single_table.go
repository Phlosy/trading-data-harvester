package truncate

import (
	"database/sql"
	"fmt"
)

func TruncateSingleTable(db *sql.DB, database string, table string) error {
	query := fmt.Sprintf("TRUNCATE TABLE %s.%s", database, table)
	_, err := db.Exec(query)
	return err
}
