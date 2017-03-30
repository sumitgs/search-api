package model

import "encoding/json"

type TwitterTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type Tweet struct {
	Text string
}

// structure containing tweets returned by Twitter search API for a resouce query
type Tweets struct {
	Statuses []Tweet  `json:"statuses,omitempty"`
	Err      ApiError `json:"ApiError,omitempty"`
}

// Decode a tweet given a HTTP response body
func (tweet *Tweets) Decode(body []byte) error {
	if err := json.Unmarshal(body, tweet); err != nil {
		return err
	}
	return nil
}
