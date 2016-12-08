package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/word", GetMastered).Methods("GET").Queries("mastered", "true")
	router.HandleFunc("/word", GetLearning).Methods("GET").Queries("mastered", "false")
	router.HandleFunc("/word", GetWords).Methods("GET")
	router.HandleFunc("/word/{word}", GetWord).Methods("GET")
	router.HandleFunc("/word", SetWord).Methods("POST")
	// router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	router.HandleFunc("/", Index)
	return router
}
