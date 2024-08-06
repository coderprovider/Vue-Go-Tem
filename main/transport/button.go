package transport

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type vueMessage struct {
	Message string `json:"message"`
}

func ButtonHandler(w http.ResponseWriter, r *http.Request) {
	var decoded vueMessage

	err := json.NewDecoder(r.Body).Decode(&decoded)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Got the following message: %s\n", decoded.Message)
}
