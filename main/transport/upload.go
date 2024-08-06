package transport

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type IdPost struct {
	Id int64 `json:"id"`
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	var decoded IdPost

	err := json.NewDecoder(r.Body).Decode(&decoded)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Uploaded data: %d\n", decoded.Id)
}
