package model

import "encoding/json"

type googleResponse struct {
	Title string
	Link  string
}

// structure containing response returned by Google search API for a resouce query
type GoogleResponses struct {
	Items []googleResponse `json:"items,omitempty"`
	Err   ApiError         `json:"ApiError,omitempty"`
}

// Decode a response given a HTTP response body
func (googleResponses *GoogleResponses) Decode(body []byte) error {
	if err := json.Unmarshal(body, googleResponses); err != nil {
		return err
	}
	return nil
}
