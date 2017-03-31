package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/search-api/model"
)

// Search handles API query by forwarding query to Google Search API, DuckDuckGo Search API and Twitter Search API. The search is cancelled
// after the timeout.
func Search(w http.ResponseWriter, r *http.Request) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	ctx, cancel = context.WithTimeout(context.Background(), 7*time.Second)

	defer cancel()

	queryPatameter := r.URL.Query().Get("q")
	if queryPatameter == "" {
		http.Error(w, "no query", http.StatusBadRequest)
		return
	}

	searchResponse := &model.SearchResponse{}
	searchResponse.Query = queryPatameter

	// Run the DuckDuckGo search and store its response in Response Channel
	duckduckGoAPIRespCh := make(chan model.Message)
	go DuckResourceQuery(ctx, queryPatameter, duckduckGoAPIRespCh)

	// Run the Google search and store its response in Response Channel
	googleAPIRespCh := make(chan model.GoogleResponses)
	go GoogleResourceQuery(ctx, queryPatameter, googleAPIRespCh)

	// Run the Twitter search and store its response in Response Channel
	twitterAPIRespCh := make(chan model.Tweets)
	go TwitterResourceQuery(ctx, queryPatameter, twitterAPIRespCh)

	searchResponse.DuckDuckGO = <-duckduckGoAPIRespCh
	searchResponse.Google = <-googleAPIRespCh
	searchResponse.Twitter = <-twitterAPIRespCh

	w.Header().Add("Content-Type", "application/json")

	resp, err := json.Marshal(searchResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(resp)

}
