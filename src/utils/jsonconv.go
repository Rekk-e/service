package service_utils

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	message string
}

func Return_json(w http.ResponseWriter, jsonResponse []byte, jsonError error) {

	if jsonError != nil {
		json.NewEncoder(w).Encode(Error{"json convert error"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
