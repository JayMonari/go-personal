package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
)

func ExamplePostForm() {
	c := http.Client{
		Transport: nil,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return nil
		},
		Jar:     nil,
		Timeout: 0,
	}
	srv := httptest.NewServer(http.HandlerFunc(func(
		w http.ResponseWriter, r *http.Request,
	) {
		w.Write([]byte(`Bearer YmFzZTY0IGVuY29kZSB0aGlzIGludG8gYWxsIG9mIHRoZSB0aGluZ3M`))
	}))
	defer srv.Close()

	req, err := http.NewRequest(http.MethodGet, "https://example.com", nil)
	req.SetBasicAuth("username", "password")
	req.Header.Add("Keep-Alive", "timeout=5,max=1000")
	raw, _ := httputil.DumpRequestOut(req, false)
	fmt.Println(string(raw))
	resp, err := c.PostForm(srv.URL, url.Values{
		"grant_type": []string{"client_credentials"},
		"scope":      []string{},
	})
	for _, b := range []byte{0x20, 0x25, 0x26, 0x2B, 0xC2, 0xA3, 0xE2, 0x82, 0xAC} {
		fmt.Printf("char: %c\n", b)
	}
	fmt.Println()
	if err != nil {
		fmt.Printf("could not connect to server: %v\n", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	fmt.Printf("%s", data)
	// Output:
}
