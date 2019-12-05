package main

import (
	"bytes"
	"net/http"
	"net/url"
)

func jsonForwardToUrl(jsonStr string, method string, destinationUrl string) error {
	_, err := url.Parse(destinationUrl)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, destinationUrl, bytes.NewBufferString(jsonStr))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	return resp.Body.Close()
}
