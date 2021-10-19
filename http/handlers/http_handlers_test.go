package handlers

import (
	"bytes"
	"encoding/json"
	"go-challenge/mocks"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestHandler_HandleTeachRequest(t *testing.T) {
	const (
		success = iota
		invalidBody
	)

	testCases := []struct {
		Name    string
		Payload interface{}
		Type    int
	}{
		{
			Name: "Test successful handling of data to trigram",
			// Build payload
			Payload: func() []byte {
				data := "getting a new test to try and see if mankind can be saved"
				payload, _ := json.Marshal(data)
				return payload
			}(),
			Type: success,
		},
		{
			Name: "Test invalid body",
			// Build payload
			Payload: func() []byte {
				data := 6
				payload, _ := json.Marshal(data)
				return payload
			}(),
			Type: invalidBody,
		},
	}

	for _, testCase := range testCases {
		controller := gomock.NewController(t)
		defer controller.Finish()

		trigramMock := mocks.NewMockTrigramIntf(controller)

		// HTTP handler options
		handlersOpts := &HandlerOptions{
			Trigram: trigramMock,
		}

		h := NewHTTPHandler(handlersOpts)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/learn", bytes.NewReader(testCase.Payload.([]byte)))
		t.Run(testCase.Name, func(t *testing.T) {
			switch testCase.Type {
			case success:
				trigramMock.EXPECT().Add(string(testCase.Payload.([]byte))).Return(true).MinTimes(1)
				h.HandleTeachRequest(w, r)
			case invalidBody:
				trigramMock.EXPECT().Add(string(testCase.Payload.([]byte))).Return(false).MinTimes(1)
				h.HandleTeachRequest(w, r)
			}
		})
	}
}

func TestHandler_HandleFetchRequest(t *testing.T) {
	const (
		success = iota
		invalidQuery
	)

	testCases := []struct {
		Name  string
		Query int
		Word  string
		Type  int
	}{
		{
			Name:  "Test successful fetching of data from trigram",
			Query: 6,
			Type:  success,
			Word:  "getting a new test to try and see if mankind can be saved",
		},
		{
			Name: "Test invalid body",
			Type: invalidQuery,
		},
	}

	for _, testCase := range testCases {
		controller := gomock.NewController(t)
		defer controller.Finish()

		trigramMock := mocks.NewMockTrigramIntf(controller)

		// HTTP handler options
		handlersOpts := &HandlerOptions{
			Trigram: trigramMock,
		}

		h := NewHTTPHandler(handlersOpts)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/generate?complexity="+strconv.Itoa(testCase.Query), nil)
		t.Run(testCase.Name, func(t *testing.T) {
			switch testCase.Type {
			case success:
				trigramMock.EXPECT().Query(testCase.Query).Return(testCase.Word).MinTimes(1)
				h.HandleFetchRequest(w, r)
			case invalidQuery:
				trigramMock.EXPECT().Query(testCase.Query).Return(testCase.Word).MinTimes(1)
				h.HandleFetchRequest(w, r)
			}
		})
	}
}
