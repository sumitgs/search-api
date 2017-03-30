package model

import "encoding/json"

type Message struct {
	AbstractSource string   `json:"AbstractSource,omitempty"`
	AbstractURL    string   `json:"AbstractURL,omitempty"`
	Err            ApiError `json:"ApiError,omitempty"`
}

func (message *Message) Decode(body []byte) error {
	if err := json.Unmarshal(body, message); err != nil {
		return err
	}
	return nil
}
