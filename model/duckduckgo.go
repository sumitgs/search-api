package model

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/search-api/util"
)

var duckBaseUrl = "https://api.duckduckgo.com/?q=%s&format=json"

type Message struct {
	AbstractSource string   `json:"AbstractSource,omitempty"`
	AbstractURL    string   `json:"AbstractURL,omitempty"`
	Err            ApiError `json:"ApiError,omitempty"`
}

func DuckQuery(query string, ch chan Message) {

	url := EncodeDuckURL(query)

	body, err := util.Do(url)

	if err != nil {
		ch <- Message{
			Err: ApiError{"Internal Server Error", 500},
		}
	}

	message := &Message{}

	if err = message.Decode(body); err != nil {
		ch <- Message{
			Err: ApiError{"Internal Server Error", 500},
		}
	}

	ch <- *message
}

func EncodeDuckURL(query string) string {
	queryEnc := url.QueryEscape(query)

	return fmt.Sprintf(duckBaseUrl, queryEnc)
}

func (message *Message) Decode(body []byte) error {
	if err := json.Unmarshal(body, message); err != nil {
		return err
	}
	return nil
}
