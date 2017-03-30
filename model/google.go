package model

import "encoding/json"

type googleResponse struct {
	Title string
	Link  string
}

type GoogleResponses struct {
	Items []googleResponse `json:"items,omitempty"`
	Err   ApiError         `json:"ApiError,omitempty"`
}

func (googleResponses *GoogleResponses) Decode(body []byte) error {
	if err := json.Unmarshal(body, googleResponses); err != nil {
		return err
	}
	return nil
}
