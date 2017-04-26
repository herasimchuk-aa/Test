package main

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
)

func main() {
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
