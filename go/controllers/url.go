package controllers

import (
	"net/http"

	"github.com/jansvabik/go-nodejs-performance/go/data"
)

// GetList gets all orders from database and sends it to client
func GetList(w http.ResponseWriter, r *http.Request) {
	orders, err := data.GetList()
	if err != nil {
		writeErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeOkResponse(w, "Full URL list successfully retrieved.", orders)
}

// Redir does the redirect to the shortened URL target
func Redir(w http.ResponseWriter, r *http.Request) {
	writeErrorResponse(w, "Not implemented yet", http.StatusInternalServerError)
}

// Create creates new document and saves it in database
func Create(w http.ResponseWriter, r *http.Request) {
	writeErrorResponse(w, "Not implemented yet", http.StatusInternalServerError)
}

// Update updates the target of specified document in database
func Update(w http.ResponseWriter, r *http.Request) {
	writeErrorResponse(w, "Not implemented yet", http.StatusInternalServerError)
}

// Delete deletes the specified document from database
func Delete(w http.ResponseWriter, r *http.Request) {
	writeErrorResponse(w, "Not implemented yet", http.StatusInternalServerError)
}
