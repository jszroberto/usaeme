package main

import (
	"encoding/json"
	"fmt"
	"github.com/uber-go/zap"
	"html/template"
	"net/http"
	"os"
)

const (
	DEFAULT_PORT = "8080"
)

func main() {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = DEFAULT_PORT
	}

	fmt.Println("Starting app on port" + port + "\n")
	http.ListenAndServe(":"+port, NewRouter())
}
