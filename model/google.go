package model

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/search-api/util"
)

var googleBaseURL = "https://www.googleapis.com/customsearch/v1?key=%s&cx=017576662512468239146:omuauf_lfve&q=%s"

type googleResponse struct {
	Title string
	Link  string
}

type GoogleResponses struct {
	Items []googleResponse `json:"items,omitempty"`
	Err   ApiError         `json:"ApiError,omitempty"`
}

func GoogleQuery(query string, ch chan GoogleResponses) {

	apikey := ""

	url := EncodeGoogleURL(query, apikey)

	body, err := util.Do(url)

	if err != nil {
		ch <- GoogleResponses{
			Err: ApiError{"Internal Server Error", 500},
		}
	}

	googleRes := &GoogleResponses{}

	if err = googleRes.Decode(body); err != nil {
		ch <- GoogleResponses{
			Err: ApiError{"Internal Server Error", 500},
		}
	}

	ch <- *googleRes
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