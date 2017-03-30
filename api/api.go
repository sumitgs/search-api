package api

import (
	"encoding/json"
	"net/http"

	"time"

	"github.com/search-api/model"
)

func Search() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")

		googleAPIResCh := getGoogleSearch(query, model.GoogleQuery)
		twitterAPIResCh := getTwitterSearch(query, model.TwitterQuery)
		duckduckGoAPIResCh := getDuckDuckGoSearch(query, model.DuckQuery)
		defer func() {
			close(googleAPIResCh)
			close(twitterAPIResCh)
			close(duckduckGoAPIResCh)
		}()

		response := model.SearchResponse{
			Query:      query,
			Google:     <-googleAPIResCh,
			DuckDuckGO: <-duckduckGoAPIResCh,
			Twitter:    <-twitterAPIResCh,
		}

		w.Header().Add("Content-Type", "application/json")
		b, err := json.Marshal(response)
		if err != nil {
			// TODO
		}
		w.Write(b)
	}
}

func getGoogleSearch(query string, s func(string, chan model.GoogleResponses)) chan model.GoogleResponses {
	ch := make(chan model.GoogleResponses)
	timer := time.NewTimer(1 * time.Second)

	go s(query, ch)

	responseCh := make(chan model.GoogleResponses)

	go func() {
		select {
		case res := <-ch:
			responseCh <- res
		case <-timer.C:
			responseCh <- model.GoogleResponses{
				Err: model.ApiError{"timeout occured", 500},
			}
		}
	}()

	return responseCh
}

func getTwitterSearch(query string, s func(string, chan model.Tweets)) chan model.Tweets {
	ch := make(chan model.Tweets)
	timer := time.NewTimer(1 * time.Second)

	go s(query, ch)

	responseCh := make(chan model.Tweets)

	go func() {
		select {
		case res := <-ch:
			responseCh <- res
		case <-timer.C:
			responseCh <- model.Tweets{
				Err: model.ApiError{"timeout occured", 500},
			}
		}

	}()

	return responseCh
}

func getDuckDuckGoSearch(query string, s func(string, chan model.Message)) chan model.Message {
	ch := make(chan model.Message)
	timer := time.NewTimer(1 * time.Second)

	go s(query, ch)

	responseCh := make(chan model.Message)

	go func() {
		select {
		case res := <-ch:
			responseCh <- res
		case <-timer.C:
			responseCh <- model.Message{
				Err: model.ApiError{"timeout occured", 500},
			}
		}

	}()

	return responseCh
}