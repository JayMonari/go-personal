package main

import (
	"io"
	"net/http/httptest"
	"testing"
)

func TestHandleUpperCase(t *testing.T) {
	req := httptest.NewRequest("", "/upper?word=abc", nil)
	w := httptest.NewRecorder()
	handleUpperCase(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != "ABC" {
		t.Errorf("want: ABC got: %s", data)
	}
}
