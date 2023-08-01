package main

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "./transfers.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
