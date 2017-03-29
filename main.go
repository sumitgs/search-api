package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/search-api/api"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", api.Search())

	log.Fatal(http.ListenAndServe(":2324", router))
}
