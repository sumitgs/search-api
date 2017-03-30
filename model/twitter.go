package model

import "encoding/json"

type TwitterTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type Tweet struct {
	Text string
}

type Tweets struct {
	Statuses []Tweet  `json:"statuses,omitempty"`
	Err      ApiError `json:"ApiError,omitempty"`
}

func (tweet *Tweets) Decode(body []byte) error {
	if err := json.Unmarshal(body, tweet); err != nil {
		return err
	}
	return nil
}
