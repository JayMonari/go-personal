package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestDoubleHandler(t *testing.T) {
	tt := []struct {
		name   string
		value  string
		double int
		status int
		err    string
	}{
		{name: "double of two", value: "2", double: 4, status: http.StatusOK},
		{name: "not a number", value: "x", err: "not a number: x", status: http.StatusBadRequest},
		{name: "missing value", value: "", err: "missing value", status: http.StatusBadRequest},
	}
	svr := httptest.NewServer(handler())
	defer svr.Close()
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res, err := http.Get(fmt.Sprintf("%s/double?v="+tc.value, svr.URL))
			if err != nil && err.Error() == tc.err {
				t.Fatalf("could not send GET request: %v", err)
			}
			defer res.Body.Close()
			if res.StatusCode != tc.status {
				t.Errorf("expected %v; got %v", tc.status, res.StatusCode)
			}

			b, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read response: %v", err)
			}
			bs := string(bytes.TrimSpace(b))
			if tc.err != "" {
				if tc.err != bs {
					t.Fatalf("expected err %v; got %v", tc.err, bs)
				}
				return
			}
			d, err := strconv.Atoi(bs)
			if err != nil {
				t.Fatalf("expected an integer; got %s", b)
			}
			if d != tc.double {
				t.Fatalf("expected double to be %v; got %v", tc.double, d)
			}
		})
	}
}
