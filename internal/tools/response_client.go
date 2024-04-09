package tools

import (
	"net/http"
)

func RequestCreator(method, url string, regNum string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return &http.Response{}, err
	}

	q := req.URL.Query()
	q.Add("regNum", regNum)

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return &http.Response{}, err
	}

	return resp, nil
}
