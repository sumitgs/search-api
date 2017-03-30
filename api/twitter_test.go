package api

import (
	"testing"

	"github.com/search-api/model"
)

func Test_TwiiterCredentials(t *testing.T) {
	SetTwitterCredential()
	if ConsumerKey == "" || ConsumerSecret == "" {
		t.Errorf("at least one credential is empty, subsequent test will fail")
	}
}

func Test_TwitterResourceQuery(t *testing.T) {
	responseChannel := make(chan model.Tweets)

	go TwitterResourceQuery("barcelona", responseChannel)

	tweets := <-responseChannel

	if len(tweets.Statuses) < 1 {
		t.Fatalf("Expected at least one tweet")
	}

}
