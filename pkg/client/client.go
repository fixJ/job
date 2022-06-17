package client

import (
	"bytes"
	"net/http"
)

func DoPost(url string, data []byte) ([]byte, error) {
	post, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	body := make([]byte, post.ContentLength)
	post.Body.Read(body)
	return body, nil
}
