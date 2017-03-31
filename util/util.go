package util

import (
	"context"
	"net/http"
)

// httpDo issues the HTTP request and calls f with the response, if context is closed while the request or f is running
// HttpDo cancels the request, waits for f to finish and returns context error. And if it is not closed, then returns function f's response
func HttpDo(ctx context.Context, request *http.Request, f func(*http.Response, error) error) error {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	c := make(chan error, 1)

	go func() {
		c <- f(client.Do(request))
	}()

	select {
	case <-ctx.Done():
		tr.CancelRequest(request)
		<-c
		return ctx.Err()
	case err := <-c:
		return err
	}
}
