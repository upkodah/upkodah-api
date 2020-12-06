package util

import (
	"errors"
	"net/http"
	"net/url"
)

func ReqGet(url string, queries map[string]string) (*http.Response, error) {
	fullURL := makeFullURL(url, queries)
	return http.Get(fullURL)
}

func ReqGetWithHeader(url string, queries map[string]string, header map[string]string) (*http.Response, error) {
	fullURL := makeFullURL(url, queries)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, errors.New("ReqGetWithHeader/" + err.Error())
	}

	for key, val := range header {
		req.Header.Add(key, val)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, errors.New("ReqGetWithHeader/" + err.Error())
	}
	return res, nil
}

func makeFullURL(urlPath string, queries map[string]string) string {
	params := url.Values{}
	for key, val := range queries {
		params.Add(key, val)
	}
	return urlPath + "?" + params.Encode()
}
