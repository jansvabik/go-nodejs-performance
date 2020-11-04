package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jansvabik/go-nodejs-performance/go/data"
)

// GetList gets all URLs from database and sends it to client
func GetList(w http.ResponseWriter, r *http.Request) {
	URLs, err := data.GetList()
	if err != nil {
		writeErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeOkResponse(w, "Full URL list successfully retrieved.", URLs)
}

// Redir does the redirect to the shortened URL target
func Redir(w http.ResponseWriter, r *http.Request) {
	URLID := mux.Vars(r)["url"]
	URL, err := data.GetByURL(URLID)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	http.Redirect(w, r, URL.Target, http.StatusPermanentRedirect)
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
