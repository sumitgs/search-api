package api

import (
	"fmt"
	"net/url"

	"github.com/search-api/model"
	"github.com/search-api/util"
)

var duckBaseUrl = "https://api.duckduckgo.com/?q=%s&format=json"

func DuckQuery(query string, ch chan model.Message) {

	url := EncodeDuckURL(query)

	body, err := util.Do(url)

	if err != nil {
		ch <- model.Message{
			Err: model.ApiError{"Internal Server Error", 500},
		}
	}

	message := &model.Message{}

	if err = message.Decode(body); err != nil {
		ch <- model.Message{
			Err: model.ApiError{"Internal Server Error", 500},
		}
	}

	ch <- *message
}

func EncodeDuckURL(query string) string {
	queryEnc := url.QueryEscape(query)

	return fmt.Sprintf(duckBaseUrl, queryEnc)
}
