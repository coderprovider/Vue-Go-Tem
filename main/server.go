package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type vueMessage struct {
	Message string `json:"message"`
}

func buttonHandler(w http.ResponseWriter, r *http.Request) {
	var decoded vueMessage

	err := json.NewDecoder(r.Body).Decode(&decoded)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Got the following message: %s\n", decoded.Message)
}

func main() {

	// Read SQL database

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

	// Start server

	http.HandleFunc("/api/hello", buttonHandler)

	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)

	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)

	// Post last created ID

	type IdPost struct {
		id int64
	}

	marshalled, err := json.Marshal(IdPost{id: lastInsertId})

	resp, err := http.Post("/api/upload", "application/json", bytes.NewReader(marshalled))
	if err != nil {
		panic(err)
	}

	fmt.Println("POST response:", resp)
}
