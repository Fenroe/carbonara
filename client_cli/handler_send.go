package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"golang.design/x/clipboard"
)

func handlerSend(s *state, _ command) error {
	type request struct {
		Content string `json:"content"`
	}

	url := fmt.Sprintf("%s/api/clips", s.apiurl)
	err := clipboard.Init()
	if err != nil {
		return err
	}
	cbdata := clipboard.Read(clipboard.FmtText)
	if len(cbdata) > 65535 {
		cbdata = cbdata[:65535]
		fmt.Println("Your clipboard data has been trimmed to fit into storage")
	}
	content := string(cbdata)
	body := request{
		Content: content,
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		return errors.New("couldn't marshal json")
	}

	// Define the handler function for sending the request
	handler := func(token string) (*http.Response, error) {
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			return nil, errors.New("couldn't create HTTP request")
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))

		return s.client.Do(req)
	}

	// Use the withAuth middleware to handle authentication and retries
	res, err := withAuth(s, handler)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	fmt.Println("Your clipboard data has been sent")
	return nil
}
