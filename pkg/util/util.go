package util

import (
	"bytes"
	"encoding/json"
	xj "github.com/basgys/goxml2json"
	"io"
	"log"
	"net/http"
	"net/url"
)

func BytesToPrettyJsonString(data []byte) string {
	b := bytes.Buffer{}
	if err := json.Indent(&b, data, "  ", "  "); err != nil {
		log.Printf("Error in BytesToPrettyJsonString : %s", err)
		return ""
	}
	return b.String()
}

func ObjectToString(data interface{}) string {
	b, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error in ObjectToString : %s", err)
		return ""
	}
	return BytesToPrettyJsonString(b)
}

func ReqGet(url string, queries map[string]string) (*http.Response, error) {
	fullURL := makeFullURL(url, queries)
	return http.Get(fullURL)
}

func makeFullURL(urlPath string, queries map[string]string) string {
	params := url.Values{}
	for key, val := range queries {
		params.Add(key, val)
	}
	return urlPath + "?" + params.Encode()
}

func Xml2json(data io.ReadCloser) (*bytes.Buffer, error) {
	buf, err := xj.Convert(data)
	if err != nil {
		return nil, err
	}
	return buf, err
}
