package helpers

import (
	"database/sql"

	"github.com/akash-searce/product-catalog/dbconnect"
)

func QueryRun(query string, v ...any) (*sql.Rows, error) {
	db := dbconnect.ConnectToDB()
	var err error
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(v...)
	return rows, err
}
