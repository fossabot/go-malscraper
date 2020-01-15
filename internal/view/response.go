package view

import (
	"encoding/json"
	"net/http"
)

// RespondWithJSON writes json response format.
func RespondWithJSON(w http.ResponseWriter, statusCode int, statusMessage string, data interface{}) {
	response := struct {
		Status  int         `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{statusCode, statusMessage, data}

	responseJSON, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(responseJSON)
}

// RespondWithCSS writes css response format (for user cover).
func RespondWithCSS(w http.ResponseWriter, statusCode int, statusMessage string, data string) {
	w.Header().Set("Content-Type", "text/css")
	w.WriteHeader(statusCode)

	if statusCode != 200 {
		w.Write([]byte(statusMessage))
	} else {
		w.Write([]byte(data))
	}
}
