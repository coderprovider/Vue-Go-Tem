package reader

import (
	"database/sql"
	"fmt"
)

func ReadSQLite() {
	db, err := sql.Open("sqlite3", "test")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into test (mytext, myinteger) values ('Success', '42')",
		"Apple", 72000)
	if err != nil {
		panic(err)
	}

	lastInsertId, _ := result.LastInsertId()
	rowsAffecred, _ := result.RowsAffected()

	fmt.Println("SQL last insert ID: ", lastInsertId)
	fmt.Println("SQL rows affected: ", rowsAffecred)
}
