package api

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"encoding/json"

	"github.com/search-api/model"
	"github.com/search-api/util"
)

var duckBaseUrl = "https://api.duckduckgo.com/?q=%s&format=json"

// query the DuckDuckGo search API for given query parameter and passes response in channel
func DuckResourceQuery(ctx context.Context, queryParameter string, responseCh chan model.Message) {

	request, err := http.NewRequest("GET", EncodeDuckURL(queryParameter), nil)

	if err != nil {
		responseCh <- model.Message{
			Err: model.ApiError{err.Error(), http.StatusInternalServerError},
		}
		return
	}

	// Issue the HTTP request and handle the respons. Request is cancelled if context is closed.
	var message model.Message
	err = util.HttpDo(ctx, request, func(response *http.Response, err error) error {
		if err != nil {
			// this is request error
			return err
		}
		defer response.Body.Close()

		if err = json.NewDecoder(response.Body).Decode(&message); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		responseCh <- model.Message{
			Err: model.ApiError{err.Error(), http.StatusInternalServerError},
		}
	} else {
		responseCh <- message
	}

}

// Encode URL given a query parameter
func EncodeDuckURL(queryParameter string) string {
	queryEnc := url.QueryEscape(queryParameter)

	return fmt.Sprintf(duckBaseUrl, queryEnc)
}
