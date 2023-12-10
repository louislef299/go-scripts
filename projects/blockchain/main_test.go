package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculateHash(t *testing.T) {
	block := Block{Index: 0, Timestamp: "2022-01-01T00:00:00Z", BPM: 60, PrevHash: "", Hash: ""}
	hash := calculateHash(block)

	expectedHash := "8e7a3116cfa94e83c18b7c2d95e54416181362390fbce93dde43572ae487cce1"
	if hash != expectedHash {
		t.Errorf("calculateHash returned incorrect hash: got %v want %v",
			hash, expectedHash)
	}
}

func TestHandleGetBlockchain(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleGetBlockchain)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if rr.Body.String() == "" {
		t.Errorf("handler returned unexpected body: got an empty body")
	}
}
