package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/word", GetWords)
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	return router
}
