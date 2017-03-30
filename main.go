package main

import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"
	"os"

	"github.com/search-api/api"
)

type Cfg struct {
	Port string
}

func main() {

	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	cfg := Cfg{}
	err := decoder.Decode(&cfg)
	if err != nil {
		fmt.Println("error: ", err.Error())
	}

	http.HandleFunc("/", api.Search())
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), nil))
}
