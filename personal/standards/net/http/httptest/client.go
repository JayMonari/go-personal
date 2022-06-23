package main

import (
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	url string
}

func NewClient(url string) Client {
	return Client{url: url}
}

func (c Client) UpperCase(word string) (string, error) {
	res, err := http.Get(c.url + "/upper?words=" + word)
	if err != nil {
		return "", fmt.Errorf("%w unable to complete Get request", err)
	}
	defer res.Body.Close()

	out, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("%w unable to read response data", err)
	}
	return string(out), nil
}
