package routes

import (
	"lxdAssessmentServer/pkg/controllers"

	"github.com/gorilla/mux"
)

func RouteHandlers() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)
	s := r.PathPrefix("/1.0").Subrouter()
	s.HandleFunc("/test", controllers.Test).Methods("GET")

	return r
}
