package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
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

func GetWord(w http.ResponseWriter, req *http.Request) {

	log := NewLogger()

	vars := mux.Vars(req)

	wordID := vars["word"]

	db, err := connectDB(log)
	if err != nil {
		log.Error("Can't reach database", zap.Error(err))
	}
	word, err := db.GetWord(wordID)
	if err != nil {
		log.Error("Can't get word", zap.Error(err))
		http.Error(w, "Not found", 404)
	} else {

		json.NewEncoder(w).Encode(word)
	}

}

func SetWord(w http.ResponseWriter, req *http.Request) {

	log := NewLogger()

	var word Word

	if req.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(req.Body).Decode(&word)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	db, err := connectDB(log)
	if err != nil {
		log.Error("Can't reach database", zap.Error(err))
		http.Error(w, "Can't reach database:"+err.Error(), 500)
	}
	err = db.SetWord(word)
	if err != nil {
		log.Error("Can't set word")
		http.Error(w, "Can't set word:"+err.Error(), 400)
	}
}
