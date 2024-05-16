package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func buttonHandler(w http.ResponseWriter, r *http.Request) {
	var decoded string

	// TODO: remove json decoder
	err := json.NewDecoder(r.Body).Decode(&decoded)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Got the following message: %s\n", decoded)
}

func main() {
	http.HandleFunc("/api/hello", buttonHandler)

	// Serve static files from the frontend/dist directory.
	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)

	// Start the server.
	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}
