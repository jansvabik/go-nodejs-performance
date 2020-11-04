package routers

import (
	"github.com/gorilla/mux"
	"github.com/jansvabik/go-nodejs-performance/go/controllers"
)

// SetUpURLs sets up routes for url resource
func SetUpURLs(router *mux.Router) *mux.Router {
	router.HandleFunc("/", controllers.GetList).Methods("GET")
	router.HandleFunc("/{url}/", controllers.Redir).Methods("GET")
	router.HandleFunc("/", controllers.Create).Methods("POST")
	router.HandleFunc("/{url}/", controllers.Update).Methods("PATCH")
	router.HandleFunc("/{url}/", controllers.Delete).Methods("DELETE")
	return router
}
