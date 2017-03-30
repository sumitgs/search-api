package main

import (
	"fmt"
	"log"
	"net/http"

	"os"

	"github.com/search-api/api"
)

func main() {
	http.HandleFunc("/", api.Search())

	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
