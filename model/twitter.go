package model

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type TwitterTokenResponse struct {
	AccessToken string `json:"access_token"`
}

var twitterBaseUrl = "https://api.twitter.com/1.1/search/tweets.json?count=11&q=%s"
var twitterOAuthUrl = "https://api.twitter.com/oauth2/token"

type Tweet struct {
	Text string
}

type Tweets struct {
	Statuses []Tweet  `json:"statuses,omitempty"`
	Err      ApiError `json:"ApiError,omitempty"`
}

func TwitterQuery(query string, ch chan Tweets) {

	client := &http.Client{}

	ConsumerKey := os.Getenv("TCK")
	ConsumerSecret := os.Getenv("TCS")

	src := ConsumerKey + ":" + ConsumerSecret
	encodedTwitterKey := base64.StdEncoding.EncodeToString([]byte(src))

	requestBody := bytes.NewReader([]byte("grant_type=client_credentials"))

	req, err := http.NewRequest("POST", twitterOAuthUrl, requestBody)
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", encodedTwitterKey))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

	resp, err := client.Do(req)
	if err != nil {
		ch <- Tweets{
			Err: ApiError{"Bad Request", 400},
		}
	}

	twitterToken := TwitterTokenResponse{}
	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&twitterToken)
	if err != nil {
		ch <- Tweets{
			Err: ApiError{"Internal Server Error", 500},
		}
	}

	bearerToken := twitterToken.AccessToken

	clients := &http.Client{}
	request, err := http.NewRequest("GET", EncodeTwitterURL(query), nil)

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", bearerToken))
	response, err := clients.Do(request)
	if err != nil {
		ch <- Tweets{
			Err: ApiError{"Bad Request", 400},
		}
	}

	tweets := &Tweets{}

	body, err := ioutil.ReadAll(response.Body)
	if err = tweets.Decode(body); err != nil {
		ch <- Tweets{
			Err: ApiError{"Internal Server Error", 500},
		}
	}

	ch <- *tweets
}

func EncodeTwitterURL(query string) string {
	queryEnc := url.QueryEscape(query)

	return fmt.Sprintf(twitterBaseUrl, queryEnc)
}

func (tweet *Tweets) Decode(body []byte) error {
	if err := json.Unmarshal(body, tweet); err != nil {
		return err
	}
	return nil
}
