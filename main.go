package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/search-api/model"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", getTwitterResp)

	log.Fatal(http.ListenAndServe(":2324", router))
}

func getDuckDuckGoResp(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query().Get("q")
	message, _ := model.DuckQuery(query)

	w.Write([]byte(model.MessageToJson(message)))

}

func getGoogleResp(w http.ResponseWriter, r *http.Request) {
	var apikey = ""
	query := r.URL.Query().Get("q")
	response, _ := model.GoogleQuery(query, apikey)

	w.Write([]byte(model.ResponseToJSON(response)))
}

func getTwitterResp(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query().Get("q")
	response, _ := model.TwitterQuery(query)

	fmt.Println("on the ai: ", response)

	w.Write([]byte(model.TweetToJson(response)))
}
