package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/search-api/model"
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

// query Twitter search API for given query parameter and returns response in a channel
func TwitterResourceQuery(queryParameter string, responseCh chan model.Tweets) {

	client := &http.Client{}

	TwitterKey := ConsumerKey + ":" + ConsumerSecret
	encodedTwitterKey := base64.StdEncoding.EncodeToString([]byte(TwitterKey))

	requestBody := bytes.NewReader([]byte("grant_type=client_credentials"))

	req, err := http.NewRequest("POST", twitterOAuthUrl, requestBody)
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", encodedTwitterKey))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

	resp, err := client.Do(req)
	if err != nil {
		responseCh <- model.Tweets{
			Err: model.ApiError{"Bad Request", 400},
		}
	}

	twitterToken := model.TwitterTokenResponse{}
	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&twitterToken)
	if err != nil {
		responseCh <- model.Tweets{
			Err: model.ApiError{"Internal Server Error", 500},
		}
	}

	bearerToken := twitterToken.AccessToken

	clients := &http.Client{}
	request, err := http.NewRequest("GET", EncodeTwitterURL(queryParameter), nil)

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", bearerToken))
	response, err := clients.Do(request)
	if err != nil {
		responseCh <- model.Tweets{
			Err: model.ApiError{"Bad Request", 400},
		}
	}

	tweets := &model.Tweets{}

	body, err := ioutil.ReadAll(response.Body)
	if err = tweets.Decode(body); err != nil {
		responseCh <- model.Tweets{
			Err: model.ApiError{"Internal Server Error", 500},
		}
	}
	responseCh <- *tweets
}

// encode URL for given query parameter
func EncodeTwitterURL(queryParameter string) string {
	queryEnc := url.QueryEscape(queryParameter)

	return fmt.Sprintf(twitterBaseUrl, queryEnc)
}
