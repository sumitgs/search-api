package api

import (
	"fmt"
	"net/url"
	"os"

	"github.com/search-api/model"
	"github.com/search-api/util"
)

var googleBaseURL = "https://www.googleapis.com/customsearch/v1?key=%s&cx=017576662512468239146:omuauf_lfve&q=%s"

func GoogleQuery(query string, ch chan model.GoogleResponses) {

	apikey := os.Getenv("GAK")

	url := EncodeGoogleURL(query, apikey)

	body, err := util.Do(url)

	if err != nil {
		ch <- model.GoogleResponses{
			Err: model.ApiError{"Internal Server Error", 500},
		}
	}

	googleRes := &model.GoogleResponses{}

	if err = googleRes.Decode(body); err != nil {
		ch <- model.GoogleResponses{
			Err: model.ApiError{"Internal Server Error", 500},
		}
	}

	ch <- *googleRes
}

func EncodeGoogleURL(query, apikey string) string {
	queryEnc := url.QueryEscape(query)

	return fmt.Sprintf(googleBaseURL, apikey, queryEnc)
}
