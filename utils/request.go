package utils

import (
	"fmt"
	"io"
	"net/http"
)

type Request struct {
}

func Httpget(url string, header map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	fmt.Println(req.Header)
	res, err := http.DefaultClient.Do(req)
	return res, err
}
func Httppost(url string, header map[string]string, payload io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	res, err := http.DefaultClient.Do(req)
	return res, err
}
