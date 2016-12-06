package main

import (
	"errors"
	"fmt"
	"github.com/cloudfoundry-community/go-cfenv"
	"html/template"
	"log"
	"net/http"
	"os"
)

const (
	DEFAULT_PORT = "8080"
)

var index = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/index.html",
))

func helloworld(w http.ResponseWriter, req *http.Request) {
	_, err := connectDB()
	if err != nil {
		panic(fmt.Sprintf("%v", err.Error()))
	}
	fmt.Printf("Stablished connection with Database")
	index.Execute(w, nil)
}

func connectDB() (*DB, error) {
	appEnv, err := cfenv.Current()
	if err != nil {
		return &DB{}, err
	}
	if redis_services, ok := appEnv.Services["compose-for-redis"]; ok {
		db := NewDatabase(redis_services[0].Credentials["uri"].(string))
		return db, db.Ping()
	} else {
		return &DB{}, errors.New("compose-for-redis service not bound to the app")
	}
}

func main() {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = DEFAULT_PORT
	}

	http.HandleFunc("/", helloworld)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Printf("Starting app on port %+v\n", port)
	http.ListenAndServe(":"+port, nil)
}
