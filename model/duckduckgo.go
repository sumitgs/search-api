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
	AbstractSource string
	AbstractURL    string
}

func Do(url string) ([]byte, error) {

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func DuckQuery(query string) (*Message, error) {

	url := EncodeDuckURL(query)

	body, err := Do(url)

	if err != nil {
		return nil, err
	}

	message := &Message{}

	if err = message.Decode(body); err != nil {
		return nil, err
	}

	return message, nil
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
