package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

var index = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/index.html",
))

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/word", GetWords)
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	return router
}

func Index(w http.ResponseWriter, req *http.Request) {

	log := zap.New(
		zap.NewJSONEncoder(zap.NoTime()), // drop timestamps in tests
	)

	db, err := connectDB(log)
	if err != nil {
		panic(fmt.Sprintf("%v", err.Error()))
	}

	log.Info("Stablished connection with Database")

	db.GetWords()

	index.Execute(w, nil)
}

func GetWords(w http.ResponseWriter, req *http.Request) {

	log := zap.New(
		zap.NewJSONEncoder(zap.NoTime()), // drop timestamps in tests
	)

	db, err := connectDB(log)
	if err != nil {
		log.Error("Can't reach database", zap.Error(err))
	}

	log.Info("Stablished connection with Database")

	words, err := db.GetWords()
	if err != nil {
		log.Error("Can't get words", zap.Error(err))
	}

	json.NewEncoder(w).Encode(words)
}
