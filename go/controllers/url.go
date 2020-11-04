package controllers

import (
	"net/http"

	"github.com/jansvabik/go-nodejs-performance/go/data"
)

// GetURLs gets all orders from database and sends it to client
func GetURLs(w http.ResponseWriter, r *http.Request) {
	orders, err := data.GetURLs()
	if err != nil {
		writeErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeOkResponse(w, "Full URL list successfully retrieved.", orders)
}
