package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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

func GoogleResourceQuery(ctx context.Context, queryParameter string, responseCh chan model.GoogleResponses) {
	request, err := http.NewRequest("GET", EncodeGoogleURL(queryParameter, apikey), nil)

	if err != nil {
		responseCh <- model.GoogleResponses{
			Err: model.ApiError{err.Error(), http.StatusInternalServerError},
		}
		return
	}

	// Issue the HTTP request and handle the respons. Request is cancelled if context is closed.
	var googleResp model.GoogleResponses
	err = util.HttpDo(ctx, request, func(response *http.Response, err error) error {
		if err != nil {
			return err
		}
		defer response.Body.Close()

		if err = json.NewDecoder(response.Body).Decode(&googleResp); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		responseCh <- model.GoogleResponses{
			Err: model.ApiError{err.Error(), http.StatusInternalServerError},
		}
	} else {
		responseCh <- googleResp
	}

}

// encode URL for given query parameter
func EncodeGoogleURL(queryParameter, apikey string) string {
	queryEnc := url.QueryEscape(queryParameter)

	return fmt.Sprintf(googleBaseURL, apikey, queryEnc)
}
