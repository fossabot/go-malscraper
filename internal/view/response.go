package view

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Response represents response model that will be converted to json.
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// RespondWithJSON writes json response format.
func RespondWithJSON(w http.ResponseWriter, statusCode int, statusMessage string, data interface{}) {
	response := Response{
		Status:  statusCode,
		Message: statusMessage,
		Data:    data,
	}

	responseJSON, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(responseJSON)))
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
