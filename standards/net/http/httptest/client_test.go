package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestClientUpperCase(t *testing.T) {
	expected := "dummy data"
	fh, _ := os.Create("output")
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
		fmt.Fprintf(os.Stdout, expected)
		fmt.Fprintf(fh, expected)
	}))
	defer svr.Close()

	c := NewClient(svr.URL)
	res, err := c.UpperCase("anything")
	if err != nil {
		t.Errorf("err should be nil, got %v", err)
	}
	if strings.TrimSpace(res) != expected {
		t.Errorf("want: %s got: %s", expected, res)
	}
}
