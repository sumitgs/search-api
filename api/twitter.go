package api

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/search-api/model"
	"github.com/search-api/util"
)

var twitterBaseUrl = "https://api.twitter.com/1.1/search/tweets.json?count=3&q=%s"
var twitterOAuthUrl = "https://api.twitter.com/oauth2/token"
var ConsumerKey string
var ConsumerSecret string

// set credential necessary for  authentication
func SetTwitterCredential() {
	ConsumerKey = os.Getenv("TWITTER_CONSUMER_KEY")
	ConsumerSecret = os.Getenv("TWITTER_CONSUMER_SECRET")
}

var bearerToken string

func SetTwitterBearerToken() error {
	client := &http.Client{}

	TwitterKey := ConsumerKey + ":" + ConsumerSecret
	encodedTwitterKey := base64.StdEncoding.EncodeToString([]byte(TwitterKey))

	requestBody := bytes.NewReader([]byte("grant_type=client_credentials"))

	req, err := http.NewRequest("POST", twitterOAuthUrl, requestBody)
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", encodedTwitterKey))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	twitterToken := model.TwitterTokenResponse{}
	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&twitterToken)
	if err != nil {
		return err
	}

	bearerToken = twitterToken.AccessToken
	return nil
}

// query Twitter search API for given query parameter and returns response in a channel
func TwitterResourceQuery(ctx context.Context, queryParameter string, responseCh chan model.Tweets) {

	request, err := http.NewRequest("GET", EncodeTwitterURL(queryParameter), nil)
	if err != nil {
		responseCh <- model.Tweets{
			Err: model.ApiError{err.Error(), http.StatusInternalServerError},
		}
		return
	}

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", bearerToken))

	// Issue the HTTP request and handle the respons. Request is cancelled if context is closed.
	var tweets model.Tweets
	err = util.HttpDo(ctx, request, func(response *http.Response, err error) error {
		if err != nil {
			return err
		}
		defer response.Body.Close()

		if err = json.NewDecoder(response.Body).Decode(&tweets); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		responseCh <- model.Tweets{
			Err: model.ApiError{err.Error(), http.StatusInternalServerError},
		}
	} else {
		responseCh <- tweets
	}
}

// encode URL for given query parameter
func EncodeTwitterURL(queryParameter string) string {
	queryEnc := url.QueryEscape(queryParameter)

	return fmt.Sprintf(twitterBaseUrl, queryEnc)
}
