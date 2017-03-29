package model

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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
	Statuses []Tweet
}

func DoTwitterQuery(url string) ([]byte, error) {

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

func TwitterQuery(query string) (*Tweets, error) {

	client := &http.Client{}

	ConsumerKey := ""
	ConsumerSecret := ""

	src := ConsumerKey + ":" + ConsumerSecret
	encodedTwitterKey := base64.StdEncoding.EncodeToString([]byte(src))

	requestBody := bytes.NewReader([]byte("grant_type=client_credentials"))

	req, err := http.NewRequest("POST", twitterOAuthUrl, requestBody)
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", encodedTwitterKey))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

	resp, err := client.Do(req)

	twitterToken := TwitterTokenResponse{}
	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&twitterToken)

	bearerToken := twitterToken.AccessToken

	clients := &http.Client{}
	request, err := http.NewRequest("GET", EncodeTwitterURL(query), nil)

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", bearerToken))
	response, err := clients.Do(request)

	tweets := &Tweets{}

	body, err := ioutil.ReadAll(response.Body)
	if err = tweets.Decode(body); err != nil {
		fmt.Println("err6: ", err.Error())
		return nil, err
	}

	return tweets, nil
}

func TweetToJson(t *Tweets) string {
	b, err := json.Marshal(t)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func EncodeTwitterURL(query string) string {
	queryEnc := url.QueryEscape(query)
	// if strings.HasPrefix(query, "!") {
	// 	return fmt.Sprintf(baseUrl, queryEnc, "&no_redirect=1")
	// }
	return fmt.Sprintf(twitterBaseUrl, queryEnc)
}

func (tweet *Tweets) Decode(body []byte) error {
	if err := json.Unmarshal(body, tweet); err != nil {
		return err
	}
	return nil
}
