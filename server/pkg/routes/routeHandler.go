package routes

import (
	"database/sql"
	"lxdAssessmentServer/ent"
	"lxdAssessmentServer/pkg/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func RouteHandlers(client *ent.Client, db *sql.DB) *mux.Router {

	h := controllers.NewHandler(client, db)

	r := mux.NewRouter().StrictSlash(true)
	r.Use(middleWare)
	s := r.PathPrefix("/1.0").Subrouter()
	s.HandleFunc("/test", h.Test).Methods("GET")
	s.HandleFunc("/book", h.CreateBook).Methods("POST")
	s.HandleFunc("/book", h.DeleteBook).Methods("DELETE")
	s.HandleFunc("/books", h.GetBooks).Methods("GET")
	s.HandleFunc("/collection", h.CreateCollection).Methods("POST")
	s.HandleFunc("/collection", h.DeleteCollection).Methods("DELETE")
	s.HandleFunc("/collections", h.GetCollections).Methods("GET")
	return r
}

func middleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
