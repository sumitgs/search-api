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
	Items []googleResponse `json:"items,omitempty"`
	Err   ApiError         `json:"apperror,omitempty"`
}

func DoGoogleQuery(url string) ([]byte, error) {

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

func GoogleQuery(query string, ch chan GoogleResponses) {

	apikey := ""

	url := EncodeGoogleURL(query, apikey)
	fmt.Println("Url: ", url)

	body, err := Do(url)

	if err != nil {
		// TODO
	}

	googleRes := &GoogleResponses{}

	if err = googleRes.Decode(body); err != nil {
		// TODO
	}

	ch <- *googleRes
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
	return fmt.Sprintf(googleBaseURL, apikey, queryEnc)
}

func (googleResponses *GoogleResponses) Decode(body []byte) error {
	if err := json.Unmarshal(body, googleResponses); err != nil {
		return err
	}
	return nil
}
