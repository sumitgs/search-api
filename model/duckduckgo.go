package model

import "encoding/json"

// structure containing response returned by DuckDuckGo search API for a resouce query
type Message struct {
	AbstractSource string   `json:"AbstractSource,omitempty"`
	AbstractURL    string   `json:"AbstractURL,omitempty"`
	Err            ApiError `json:"ApiError,omitempty"`
}

// Decode a message given a HTTP response body
func (message *Message) Decode(body []byte) error {
	if err := json.Unmarshal(body, message); err != nil {
		return err
	}
	return nil
}
