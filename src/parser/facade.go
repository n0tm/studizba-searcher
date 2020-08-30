package parser

import (
	"fmt"
	"net/http"
)

func makeRequest(url string) (*http.Response, error) {
	response, err := http.Get(url)
	if err != nil {
		return response, err
	}

	if response.StatusCode != 200 {
		return response, fmt.Errorf("status code error: %d %s", response.StatusCode, response.Status)
	}

	return response, err
}
