package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var duckBaseUrl = "https://api.duckduckgo.com/?q=%s&format=json"

type Message struct {
	AbstractSource string   `json:"AbstractSource,omitempty"`
	AbstractURL    string   `json:"AbstractURL,omitempty"`
	Err            ApiError `json:"apperror,omitempty"`
}

func Do(url string) ([]byte, error) {

	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func DuckQuery(query string, ch chan Message) {

	url := EncodeDuckURL(query)

	body, err := Do(url)

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
