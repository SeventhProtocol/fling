package core

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func Bootstrap() bool {

	db, err := sql.Open("sqlite3", "fling.db")

	if err != nil {
		fmt.Printf("Couldn't open the SQLiteDB. It's useless to continue. It's over!\n")
		os.Exit(1)
	}

	db.Close()

	return false
}
