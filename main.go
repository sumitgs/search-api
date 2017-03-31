package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/search-api/api"
)

func main() {

	http.HandleFunc("/", api.Search)

	api.SetGoogleCredential()
	api.SetTwitterCredential()
	err := api.SetTwitterBearerToken()

	if err != nil {
		fmt.Println("cannot set twitter Bearer Token, error is: ", err.Error())
		return
	}

	fmt.Println("ready to go")

	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
