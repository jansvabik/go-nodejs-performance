package routers

import (
	"github.com/gorilla/mux"
	"github.com/jansvabik/go-nodejs-performance/go/controllers"
)

// SetUpURLs sets up routes for url resource
func SetUpURLs(router *mux.Router) *mux.Router {
	router.HandleFunc("/", controllers.GetURLs).Methods("GET")
	return router
}
