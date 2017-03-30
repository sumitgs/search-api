package api

import (
	"fmt"
	"net/url"

	"github.com/search-api/model"
	"github.com/search-api/util"
)

var duckBaseUrl = "https://api.duckduckgo.com/?q=%s&format=json"

// query the DuckDuckGo search API for given query parameter and passes response in channel
func DuckResourceQuery(queryParameter string, responseCh chan model.Message) {
	url := EncodeDuckURL(queryParameter)
	body, err := util.Do(url)
	if err != nil {
		responseCh <- model.Message{
			Err: model.ApiError{"Internal Server Error", 500},
		}
	}
	message := &model.Message{}
	if err = message.Decode(body); err != nil {
		responseCh <- model.Message{
			Err: model.ApiError{"Internal Server Error", 500},
		}
	}
	responseCh <- *message
}

// Encode URL given a query parameter
func EncodeDuckURL(queryParameter string) string {
	queryEnc := url.QueryEscape(queryParameter)

	return fmt.Sprintf(duckBaseUrl, queryEnc)
}
