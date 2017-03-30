package model

import "encoding/json"

// structure containing response returned by multiple sites
type SearchResponse struct {
	Query      string
	Google     GoogleResponses
	DuckDuckGO Message
	Twitter    Tweets
}

type ApiError struct {
	Message string
	Code    int
}

func SearchToJSON(s *SearchResponse) string {
	b, err := json.Marshal(s)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}
