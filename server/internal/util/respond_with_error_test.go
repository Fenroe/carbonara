package util

import (
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"
)

func TestRespondWithError(t *testing.T) {
	// mock data
	recorder := httptest.NewRecorder()
	statusCodse := 500
	message := "this is an error message"
	testErr := errors.New("internal server error")
	// test
	RespondWithError(recorder, statusCodse, message, testErr)
	res := recorder.Result()
	defer res.Body.Close()
	// check status code
	if res.StatusCode != statusCodse {
		t.Errorf("mismatch on field Statuscode; expected %d, but got %d", statusCodse, res.StatusCode)
	}
	// check body
	type response struct {
		Error string `json:"error"`
	}
	resBody := response{}
	err := json.NewDecoder(res.Body).Decode(&resBody)
	if err != nil {
		t.Errorf("could not marshal response data: %s", err)
	}
	if resBody.Error != message {
		t.Errorf("mismatch on field Error; expected %s, but got %s", message, resBody.Error)
	}
}
