package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"os"
	"strings"
)

// NOTE(docs): when reading malformed chunked data with lines that are too long
var ErrLineTooLong = errors.New("header line too long")

func ExampleDumpRequest() {
	req, _ := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"http://example.com/some/path?query=string&fragment=coming#rightnow",
		bytes.NewReader([]byte("The body of the request that will be outbound soon.")),
	)
	req.Header.Add("Host", "example.com")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:106.0) Gecko/20100101 Firefox/106.0")
	req.Header.Add("Accept", "image/avif,image/webp,*/*")
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Referer", "https://example.com/static/css/main.343e270f.css")

	// NOTE(jay): What gets dumped can be found in [http.Request.Write]
	//
	//   This method consults the following fields of the request:
	//     Host
	//     URL
	//     Method (defaults to "GET")
	//     Header
	//     ContentLength
	//     TransferEncoding
	//     Body
	b, err := httputil.DumpRequest(req, true)
	// Need to remove carriage returns ('\r') for Output
	fmt.Printf("err: %v\n\n%s\n\n", err, bytes.ReplaceAll(b, []byte{'\r'}, nil))

	rd := make([]byte, 1<<8)
	n, err := req.Body.Read(rd)
	fmt.Printf("body was replaced because %d bytes were read and there is no error: %v", n, err)
	// Output:
	// err: <nil>
	//
	// POST /some/path?query=string&fragment=coming HTTP/1.1
	// Host: example.com
	// Accept: image/avif,image/webp,*/*
	// Accept-Encoding: gzip, deflate, br
	// Accept-Language: en-US,en;q=0.5
	// Referer: https://example.com./static/css/main.343e270f.css
	// User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:106.0) Gecko/20100101 Firefox/106.0
	//
	// The body of the request that will be outbound soon.
	//
	// body was replaced because 51 bytes were read and there is no error: <nil>
}

func ExampleDumpRequestOut() {
	// NOTE(jay): Adds other headers that [http.Transport] would add like User-Agent and
	// Content-Length
	req, _ := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"http://example.com/some/path?query=string&fragment=coming#rightnow",
		bytes.NewReader([]byte("The body of the request that will be outbound soon.")),
	)
	req.Header.Add("Host", "example.com")
	req.Header.Add("Accept", "image/avif,image/webp,*/*")
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Referer", "https://example.com./static/css/main.343e270f.css")
	b, err := httputil.DumpRequestOut(req, true)
	fmt.Printf("err: %v\n\n%s\n\n", err, bytes.ReplaceAll(b, []byte{'\r'}, nil))
	// Output:
	// err: <nil>
	//
	// POST /some/path?query=string&fragment=coming HTTP/1.1
	// Host: example.com
	// User-Agent: Go-http-client/1.1
	// Content-Length: 51
	// Accept: image/avif,image/webp,*/*
	// Accept-Encoding: gzip, deflate, br
	// Accept-Language: en-US,en;q=0.5
	// Referer: https://example.com./static/css/main.343e270f.css
	//
	// The body of the request that will be outbound soon.
}

func ExampleDumpResponse() {
	// Output:
}

func ExampleNewChunkedReader() {
	// NOTE(docs): NewChunkedReader is not needed by normal applications. The http package
	// automatically decodes chunking when reading response bodies.
	r := httputil.NewChunkedReader(bytes.NewReader([]byte(
		"7\r\nChunked\r\n8\r\n-Encoded\r\nb\r\n-Transfered\r\n0\r\n")))
	// b := make([]byte, 1<<8)
	// n, err := r.Read(b)
	// fmt.Println(n, err)
	// n, err = r.Read(b)
	// fmt.Println(n, err, string(b))
	// n, err = r.Read(b)
	// fmt.Println(n, err)
	io.Copy(os.Stdout, r)

	// Output:
	// Chunked-Encoded-Transfered
}

func ExampleNewChunkedWriter() {
	// NOTE(docs): NewChunkedWriter is not needed by normal applications. The http package
	// adds chunking automatically if handlers don't set a Content-Length header. Using
	// NewChunkedWriter inside a handler would result in double chunking or chunking with a
	// Content-Length length, both of which are wrong.
	rr := httptest.NewRecorder()
	cw := httputil.NewChunkedWriter(rr)
	cw.Write([]byte("chunky chunky chunky"))
	cw.Write([]byte("data data data"))
	cw.Write([]byte("another one and another one"))
	cw.Write([]byte("another one and another one"))
	cw.Close()
	// NOTE(docs): Closing the returned chunkedWriter sends the final 0-length chunk that
	// marks the end of the stream but does not send the final CRLF that appears after
	// trailers; trailers and the last CRLF must be written separately.
	rr.Write([]byte("Expires: Mon, 02 Jan 2006 15:04:05 GMT\r\n")) // trailer
	rr.Write([]byte("\r\n"))
	fmt.Println(strings.ReplaceAll(rr.Body.String(), "\r", ""))
	// Output:
	// 14
	// chunky chunky chunky
	// e
	// data data data
	// 1b
	// another one and another one
	// 1b
	// another one and another one
	// 0
	// Expires: Mon, 02 Jan 2006 15:04:05 GMT
}

// type BufferPool interface{ ... }
// type ReverseProxy struct{ ... }

// u, _ := url.Parse("http://localhost:9001")
// httputil.NewSingleHostReverseProxy(u)
func ExampleNewSingleHostReverseProxy() {
	// Output:
}

var (
	// NOTE(jay): DEPRECATED! Don't use.
	DEPRECATED_ErrPersistEOF = ""

	DEPRECATED_NewServerConn = func() {
		// NOTE(jay): DEPRECATED!!! Use [http.Server] instead.
		// type ServerConn struct{ ... }
	}

	DEPRECATED_NewClientConn = func() {
		// NOTE(jay): DEPRECATED!!! Use [http.Client] or [http.Transport] instead.
		// type ClientConn struct{ ... }
	}

	DEPRECATED_NewProxyClientConn = func() {
		// NOTE(jay): DEPRECATED!!! Use [http.Client] or [http.Transport] instead.
		// type ClientConn struct{ ... }
	}
)
