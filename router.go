package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

// NewRouter creates an object with the REST API
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/word", GetMastered).Methods("GET").Queries("mastered", "true")
	router.HandleFunc("/word", GetLearning).Methods("GET").Queries("mastered", "false")
	router.HandleFunc("/word", GetWords).Methods("GET")
	router.HandleFunc("/word/{word}", GetWord).Methods("GET")
	router.HandleFunc("/word/{word}", DeleteWord).Methods("DELETE")
	// router.HandleFunc("/word/{word}", PutWord).Methods("PUT")
	router.HandleFunc("/word", SetWord).Methods("POST")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	router.HandleFunc("/", Index)
	return router
}
