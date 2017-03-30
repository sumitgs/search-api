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
	Err            ApiError `json:"apperror,omitempty"`
}

func DuckQuery(query string, ch chan Message) {

	url := EncodeDuckURL(query)

	body, err := util.Do(url)

	if err != nil {
		// TODO
	}

	message := &Message{}

	if err = message.Decode(body); err != nil {
		// TODO
	}

	ch <- *message
}

func MessageToJson(m *Message) string {
	b, err := json.Marshal(m)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func EncodeDuckURL(query string) string {
	queryEnc := url.QueryEscape(query)
	// if strings.HasPrefix(query, "!") {
	// 	return fmt.Sprintf(baseUrl, queryEnc, "&no_redirect=1")
	// }
	return fmt.Sprintf(duckBaseUrl, queryEnc)
}

func (message *Message) Decode(body []byte) error {
	if err := json.Unmarshal(body, message); err != nil {
		return err
	}
	return nil
}
