package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"golang.design/x/clipboard"
)

func handlerSync(s *state, _ command) error {
	type responseClip struct {
		Content string `json:"content"`
	}

	type response struct {
		Clips []responseClip `json:"clips"`
	}

	url := fmt.Sprintf("%s/api/clips", s.apiurl)
	err := clipboard.Init()
	if err != nil {
		return err
	}
	// Define the handler function for sending the request
	handler := func(token string) (*http.Response, error) {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, errors.New("couldn't create HTTP request")
		}
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))

		return s.client.Do(req)
	}

	// Use the withAuth middleware to handle authentication and retries
	res, err := withAuth(s, handler)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	resBody := response{}
	if err = json.Unmarshal(data, &resBody); err != nil {
		return err
	}
	latestClip := resBody.Clips[0]
	clipboard.Write(clipboard.FmtText, []byte(latestClip.Content))
	// At this point the clipboard is blank
	// For whatever reason, reading the clipboard fixes that
	clipboard.Read(clipboard.FmtText)
	fmt.Println("Your clipboard data has been synced")
	return nil
}
