package tools

import (
	"io"
	"net/http"
)

func RequestCreator(method, url string, reader io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		return &http.Response{}, err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &http.Response{}, err
	}

	return resp, nil
}
