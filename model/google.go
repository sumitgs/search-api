package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var googleBaseURL = "https://www.googleapis.com/customsearch/v1?key=%s&cx=017576662512468239146:omuauf_lfve&q=%s"

type googleResponse struct {
	Title string
	Link  string
}

type GoogleResponses struct {
	Items []googleResponse
}

func DoGoogleQuery(url string) ([]byte, error) {
	fmt.Println("I am  here 1")

	response, err := http.Get(url)
	fmt.Println("I am here 2")
	if err != nil {
		return nil, err
	}

	fmt.Println("I am here 3")
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func GoogleQuery(query, apikey string) (*GoogleResponses, error) {

	url := EncodeGoogleURL(query, apikey)
	fmt.Println("Url: ", url)

	body, err := Do(url)

	if err != nil {
		return nil, err
	}

	message := &GoogleResponses{}

	if err = message.Decode(body); err != nil {
		return nil, err
	}

	return message, nil
}

func ResponseToJSON(m *GoogleResponses) string {
	b, err := json.Marshal(m)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func EncodeGoogleURL(query, apikey string) string {
	queryEnc := url.QueryEscape(query)
	// if strings.HasPrefix(query, "!") {
	// 	return fmt.Sprintf(baseUrl, queryEnc, "&no_redirect=1")
	// }
	return fmt.Sprintf(googleBaseURL, apikey, queryEnc)
}

func (googleResponses *GoogleResponses) Decode(body []byte) error {
	if err := json.Unmarshal(body, googleResponses); err != nil {
		return err
	}
	return nil
}
