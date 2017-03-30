package util

import (
	"io/ioutil"
	"net/http"
)

// Do the HTTP requests against a body
func Do(url string) ([]byte, error) {

	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
