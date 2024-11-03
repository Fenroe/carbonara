package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func doPostRequest[T any](s *state, body T, url string) (*http.Response, error) {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return res, fmt.Errorf("unexpected HTTP status: %s", res.Status)
	}
	return res, err
}
