package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"os"
	"encoding/json"
)

func TestWelcomeAPI(t *testing.T) {

	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	requestResponse := httptest.NewRecorder()
	handler := http.HandlerFunc(hello)
	handler.ServeHTTP(requestResponse, request)

	if status := requestResponse.Code; status != http.StatusOK {
		t.Errorf("API returned wrong status code, Returned: %v want %v", status, http.StatusOK)
	}

	expected := `{"Welcome":"Hello World!"}`

	if requestResponse.Body.String() != expected {
		t.Errorf("API returned unexpected JSON message, Returned: %v Expected %v", requestResponse.Body.String(), expected)
	}
}

func TestHeathAPI(t *testing.T) {

	request, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	requestResponse := httptest.NewRecorder()
	handler := http.HandlerFunc(healthCheck)
	handler.ServeHTTP(requestResponse, request)

	if status := requestResponse.Code; status != http.StatusOK {
		t.Errorf("API returned wrong status code, Returned: %v want %v", status, http.StatusOK)
	}

	expected := `{"health":"Still alive!"}`

	if requestResponse.Body.String() != expected {
		t.Errorf("API returned unexpected JSON message, Returned: %v Expected %v", requestResponse.Body.String(), expected)
	}
}

func TestGetInfoAPI(t *testing.T) {
	
	request, err := http.NewRequest("GET", "/getInfo", nil)
	if err != nil {
		t.Fatal(err)
	}

	requestResponse := httptest.NewRecorder()
	handler := http.HandlerFunc(getInfo)
	handler.ServeHTTP(requestResponse, request)

	if status := requestResponse.Code; status != http.StatusOK {
		t.Errorf("API returned wrong HTTP status code, Returned: %v Expected: %v", status, http.StatusOK)
	}

	myInfo := info{
		Version: os.Getenv("VERSION"),
		LastCommitSHA: os.Getenv("VERSION"),
		Description:   "pre-interview technical test",
	}

	myApp := app{
		AppName: []info{
			myInfo,
		},
	}

	infoJSON, err := json.Marshal(myApp)
	
	if requestResponse.Body.String() != string(infoJSON) {
		t.Errorf("API returned unexpected JSON message, Returned: %v  Expected: %v", requestResponse.Body.String(), infoJSON)
	}
}
