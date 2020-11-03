package util

import (
	"bytes"
	"encoding/json"
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
