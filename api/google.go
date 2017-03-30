package api

import (
	"fmt"
	"net/url"
	"os"

	"github.com/search-api/model"
	"github.com/search-api/util"
)

var googleBaseURL = "https://www.googleapis.com/customsearch/v1?key=%s&cx=017576662512468239146:omuauf_lfve&q=%s"
var apikey string

// set credential necessary for authentication
func SetGoogleCredential() {
	apikey = os.Getenv("GOOGLE_API_KEY")
}

// query Google search API for given query parameter and returns response in a channel
func GoogleResourceQuery(queryParameter string, responseCh chan model.GoogleResponses) {

	url := EncodeGoogleURL(queryParameter, apikey)

	body, err := util.Do(url)

	if err != nil {
		responseCh <- model.GoogleResponses{
			Err: model.ApiError{"Internal Server Error", 500},
		}
	}

	googleRes := &model.GoogleResponses{}

	if err = googleRes.Decode(body); err != nil {
		responseCh <- model.GoogleResponses{
			Err: model.ApiError{"Internal Server Error", 500},
		}
	}

	responseCh <- *googleRes
}

// encode URL for given query parameter
func EncodeGoogleURL(queryParameter, apikey string) string {
	queryEnc := url.QueryEscape(queryParameter)

	return fmt.Sprintf(googleBaseURL, apikey, queryEnc)
}
