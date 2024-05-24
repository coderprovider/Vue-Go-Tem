package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"runtime"

	reader "github.com/chnmk/vue-go-playground/main/sql"

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
	// Read SQL
	reader.ReadSQLite()

	// Run handlers
	http.HandleFunc("/api/hello", buttonHandler)
	http.HandleFunc("/api/upload", uploadHandler)

	// Serve frontend app
	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)

	// Create listener
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Panic(err)
	}

	// Serve in a goroutine
	go func() {
		log.Panic(
			http.Serve(ln, nil),
		)
	}()

	// Keep server goroutine alive until exit
	runtime.Goexit()
	fmt.Println("Exit")

}
