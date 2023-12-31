package utils

import (
	"io"
	"net/http"
)

func FetchSchema(URL string) ([]byte, error) {
	var body []byte
	var client http.Client
	var err error
	resp, err := client.Get(URL)
	if err != nil {
		return body, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return body, err
		}
	}
	return body, err
}
