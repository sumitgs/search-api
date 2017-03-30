package main

import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"
	"os"

	"github.com/gorilla/mux"
	"github.com/search-api/api"
)

type Cfg struct {
	Port string
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", api.Search())

	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	cfg := Cfg{}
	err := decoder.Decode(&cfg)
	if err != nil {
		fmt.Println("error: ", err.Error())
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), router))
}
