package api

import (
	"testing"

	"github.com/search-api/model"
)

func Test_DuckResourceQuery(t *testing.T) {

	responseChannel := make(chan model.Message)
	go DuckResourceQuery("barcelona", responseChannel)

	message := <-responseChannel
	if message.AbstractSource == "" {
		t.Fatalf("Expectred response from DuckDuckGo API")
	}

}
