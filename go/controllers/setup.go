package controllers

import (
	"encoding/json"
	"net/http"
)

// response is the response data model
type response struct {
	Status string      `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

// writeErrorResponse writes an error response to the response writer
func writeErrorResponse(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response{
		Status: "ERROR",
		Msg:    msg,
		Data:   nil,
	})
}

// writeOkResponse writes an ok response to the response writer
func writeOkResponse(w http.ResponseWriter, msg string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response{
		Status: "OK",
		Msg:    msg,
		Data:   data,
	})
}
