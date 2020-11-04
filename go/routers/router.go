package routers

import (
	"github.com/gorilla/mux"
)

// InitRoutes initialize all routes of this REST API service
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetUpURLs(router)
	return router
}
