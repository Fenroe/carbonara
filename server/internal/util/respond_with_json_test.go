package util

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestRespondWithJSON(t *testing.T) {
	// mock data
	recorder := httptest.NewRecorder()
	type body struct {
		Message string `json:"message"`
	}
	statusCode := 200
	payload := body{
		Message: "Success!",
	}
	RespondWithJSON(recorder, statusCode, payload)
	res := recorder.Result()
	defer res.Body.Close()
	// check status code
	if res.StatusCode != statusCode {
		t.Errorf("mismatch on field Statuscode; expected %d, but got %d", statusCode, res.StatusCode)
	}
	// check body
	resBody := body{}
	err := json.NewDecoder(res.Body).Decode(&resBody)
	if err != nil {
		t.Errorf("could not marshal response data: %s", err)
	}
	if resBody.Message != payload.Message {
		t.Errorf("mismatch on field Error; expected %s, but got %s", payload.Message, resBody.Message)
	}
}
