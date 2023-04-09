package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"cloud.google.com/go/firestore"
)

func TestAPIHandler(t *testing.T) {
	client, err := firestore.NewClient(context.Background(), "project-resume-24")
	if err != nil {
		t.Fatalf("Failed to create Firestore client: %v", err)
	}

	// create request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// create response recorder
	rr := httptest.NewRecorder()

	// create handler and serve the request
	handler := APIHandler(client)
	handler.ServeHTTP(rr, req)

	// check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// check if the value in the response is an integer
	var data map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := data["value"].(float64); !ok {
		t.Errorf("handler returned unexpected value type: got %T want float64", data["value"])
	}
}
