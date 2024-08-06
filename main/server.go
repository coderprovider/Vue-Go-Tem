package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"runtime"

	"github.com/chnmk/vue-go-playground/main/transport"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Read SQL
	// reader.ReadSQLite()

	// Run handlers
	http.HandleFunc("/api/hello", transport.ButtonHandler)
	http.HandleFunc("/api/upload", transport.UploadHandler)

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
