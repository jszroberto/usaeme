package main

import (
	"encoding/json"
	"github.com/uber-go/zap"
	"html/template"
	"net/http"
)

var index = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/index.html",
))

func Index(w http.ResponseWriter, req *http.Request) {
	index.Execute(w, nil)
}

func GetWords(w http.ResponseWriter, req *http.Request) {

	log := NewLogger()

	db, err := connectDB(log)
	if err != nil {
		log.Error("Can't reach database", zap.Error(err))
	}

	words, err := db.GetWords()
	if err != nil {
		log.Error("Can't get words", zap.Error(err))
	}

	json.NewEncoder(w).Encode(words)
}
