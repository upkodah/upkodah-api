package util

import (
	"bytes"
	"encoding/json"
	"errors"
	xj "github.com/basgys/goxml2json"
	"io"
	"log"
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

func Xml2json(data io.ReadCloser) (*bytes.Buffer, error) {
	buf, err := xj.Convert(data)
	if err != nil {
		return nil, errors.New("Xml2json/" + err.Error())
	}
	return buf, err
}
