package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"runtime"

	_ "github.com/mattn/go-sqlite3"
)

type vueMessage struct {
	Message string `json:"message"`
}

type IdPost struct {
	Id int64 `json:"id"`
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

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	var decoded IdPost

	err := json.NewDecoder(r.Body).Decode(&decoded)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Uploaded data: %d\n", decoded.Id)
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

	// Server startup

	http.HandleFunc("/api/hello", buttonHandler)
	http.HandleFunc("/api/upload", uploadHandler)

	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)

	// Post last created ID

	/*
		marshalled, err := json.Marshal(IdPost{Id: lastInsertId})
		if err != nil {
			panic(err)
		}


			resp, err := http.Post("/api/upload", "application/json", bytes.NewReader(marshalled))
			if err != nil {
				panic(err)
			}

			fmt.Println("POST response:", resp)
	*/

	// Server listen

	// fmt.Println("Server listening on port 3000")

	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Panic(err)
	}

	go func() {
		log.Panic(
			http.Serve(ln, nil),
		)
	}()

	runtime.Goexit()
	fmt.Println("Exit")

}
