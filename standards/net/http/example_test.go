package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"testing"
)

// NOTE(docs): Unless otherwise noted, these are defined in RFC 7231 section 4.3.
const (
	MethodGet     = "GET"
	MethodHead    = "HEAD"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH" // RFC 5789
	MethodDelete  = "DELETE"
	MethodConnect = "CONNECT"
	MethodOptions = "OPTIONS"
	MethodTrace   = "TRACE"
)

// NOTE(docs): HTTP status codes as registered with IANA. See:
// https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
const (
	StatusContinue           = 100 // RFC 9110, 15.2.1
	StatusSwitchingProtocols = 101 // RFC 9110, 15.2.2
	StatusProcessing         = 102 // RFC 2518, 10.1
	StatusEarlyHints         = 103 // RFC 8297

	StatusOK                   = 200 // RFC 9110, 15.3.1
	StatusCreated              = 201 // RFC 9110, 15.3.2
	StatusAccepted             = 202 // RFC 9110, 15.3.3
	StatusNonAuthoritativeInfo = 203 // RFC 9110, 15.3.4
	StatusNoContent            = 204 // RFC 9110, 15.3.5
	StatusResetContent         = 205 // RFC 9110, 15.3.6
	StatusPartialContent       = 206 // RFC 9110, 15.3.7
	StatusMultiStatus          = 207 // RFC 4918, 11.1
	StatusAlreadyReported      = 208 // RFC 5842, 7.1
	StatusIMUsed               = 226 // RFC 3229, 10.4.1

	StatusMultipleChoices  = 300 // RFC 9110, 15.4.1
	StatusMovedPermanently = 301 // RFC 9110, 15.4.2
	StatusFound            = 302 // RFC 9110, 15.4.3
	StatusSeeOther         = 303 // RFC 9110, 15.4.4
	StatusNotModified      = 304 // RFC 9110, 15.4.5
	StatusUseProxy         = 305 // RFC 9110, 15.4.6

	StatusTemporaryRedirect = 307 // RFC 9110, 15.4.8
	StatusPermanentRedirect = 308 // RFC 9110, 15.4.9

	StatusBadRequest                   = 400 // RFC 9110, 15.5.1
	StatusUnauthorized                 = 401 // RFC 9110, 15.5.2
	StatusPaymentRequired              = 402 // RFC 9110, 15.5.3
	StatusForbidden                    = 403 // RFC 9110, 15.5.4
	StatusNotFound                     = 404 // RFC 9110, 15.5.5
	StatusMethodNotAllowed             = 405 // RFC 9110, 15.5.6
	StatusNotAcceptable                = 406 // RFC 9110, 15.5.7
	StatusProxyAuthRequired            = 407 // RFC 9110, 15.5.8
	StatusRequestTimeout               = 408 // RFC 9110, 15.5.9
	StatusConflict                     = 409 // RFC 9110, 15.5.10
	StatusGone                         = 410 // RFC 9110, 15.5.11
	StatusLengthRequired               = 411 // RFC 9110, 15.5.12
	StatusPreconditionFailed           = 412 // RFC 9110, 15.5.13
	StatusRequestEntityTooLarge        = 413 // RFC 9110, 15.5.14
	StatusRequestURITooLong            = 414 // RFC 9110, 15.5.15
	StatusUnsupportedMediaType         = 415 // RFC 9110, 15.5.16
	StatusRequestedRangeNotSatisfiable = 416 // RFC 9110, 15.5.17
	StatusExpectationFailed            = 417 // RFC 9110, 15.5.18
	StatusTeapot                       = 418 // RFC 9110, 15.5.19 (Unused)
	StatusMisdirectedRequest           = 421 // RFC 9110, 15.5.20
	StatusUnprocessableEntity          = 422 // RFC 9110, 15.5.21
	StatusLocked                       = 423 // RFC 4918, 11.3
	StatusFailedDependency             = 424 // RFC 4918, 11.4
	StatusTooEarly                     = 425 // RFC 8470, 5.2.
	StatusUpgradeRequired              = 426 // RFC 9110, 15.5.22
	StatusPreconditionRequired         = 428 // RFC 6585, 3
	StatusTooManyRequests              = 429 // RFC 6585, 4
	StatusRequestHeaderFieldsTooLarge  = 431 // RFC 6585, 5
	StatusUnavailableForLegalReasons   = 451 // RFC 7725, 3

	StatusInternalServerError           = 500 // RFC 9110, 15.6.1
	StatusNotImplemented                = 501 // RFC 9110, 15.6.2
	StatusBadGateway                    = 502 // RFC 9110, 15.6.3
	StatusServiceUnavailable            = 503 // RFC 9110, 15.6.4
	StatusGatewayTimeout                = 504 // RFC 9110, 15.6.5
	StatusHTTPVersionNotSupported       = 505 // RFC 9110, 15.6.6
	StatusVariantAlsoNegotiates         = 506 // RFC 2295, 8.1
	StatusInsufficientStorage           = 507 // RFC 4918, 11.5
	StatusLoopDetected                  = 508 // RFC 5842, 7.2
	StatusNotExtended                   = 510 // RFC 2774, 7
	StatusNetworkAuthenticationRequired = 511 // RFC 6585, 6
)

func ExampleCanonicalHeaderKey() {
	canon := func(s string) {
		fmt.Printf("key:\t\t%q\ncanonical:\t%q\n\n", s, http.CanonicalHeaderKey(s))
	}
	canon("if-match")
	canon("keep-alive")
	canon("accept")
	canon("authorization")
	canon("access-control-allow-origin")
	// spaces and bad bytes just return the string
	canon("ac\x00cept")
	canon("if match")
	// Output:
	// key:		"if-match"
	// canonical:	"If-Match"
	//
	// key:		"keep-alive"
	// canonical:	"Keep-Alive"
	//
	// key:		"accept"
	// canonical:	"Accept"
	//
	// key:		"authorization"
	// canonical:	"Authorization"
	//
	// key:		"access-control-allow-origin"
	// canonical:	"Access-Control-Allow-Origin"
	//
	// key:		"ac\x00cept"
	// canonical:	"ac\x00cept"
	//
	// key:		"if match"
	// canonical:	"if match"
}

func ExampleDetectContentType() {
	sniff := func(b []byte) {
		fmt.Printf("type: %s\n", http.DetectContentType(b))
	}
	sniff([]byte("<!DOCTYPE html><html></html>"))
	sniff([]byte(`%PDF-1.7%âãÏÓ 1 0 obj<</Annots 47 0 R/Contents 2 0 R/CropBox[0 0 446.25 631.5]/Group<</CS/DeviceRGB/I true/S/Transparency/Type/Group>>/MediaBox[0 0 446.25 631.5]/Parent 95 0 R/Resources<</ExtGState<</GS0 133 0 R>>/XObject<</Fm0 12 0 R>>>>/Rotate 0/Type/Page>>endobj2 0 obj<</Filter/FlateDecode/Length 46>>stream H*ä2T0 `))
	sniff([]byte("GIF89a]]]S¼ÛHVZÐÐÐQQQLnxVVV```ÅÅÅWÄãKfn¢¢¢ZZZÜÜÜ¦^^MMMI^d¹¹¹çççNóóóGúúúQ¶Ö²³´¼¼¼®®®"))
	sniff([]byte("RIFF\x00\x00\x00\x00WEBPVP8X \x00\x00\x00\x00\x00\x00Ç\x00\x00¡\x00\x00ANIM\x00\x00\x00ÿÿÿÿ\x00\x00ANMFüF\x00\x00\x00\x00\x00\x00\x00\x00Ç\x00\x00¡\x00\x00B\x00\x00VP8"))
	fmt.Println()

	// NOTE(jay): Not all MIME types are recognized.
	sniff([]byte(`
yaml:
  laml:
    ding: "dong"`[1:])) // application/yaml
	sniff([]byte(`{"this":"is","json":true}`)) // application/json
	// Output:
	// type: text/html; charset=utf-8
	// type: application/pdf
	// type: image/gif
	// type: image/webp
	//
	// type: text/plain; charset=utf-8
	// type: text/plain; charset=utf-8
}

func ExampleError() {
	rr := httptest.NewRecorder()
	http.Error(rr, "there was an error processing your request", http.StatusInternalServerError)
	fmt.Println("status:", rr.Result().Status, "body:", rr.Body.String())

	// XXX(jay): [http.Error] does NOT end the request, so you need to `return` before
	// writing anything to the [http.ResponseWriter] by accident.
	rr.Write([]byte("Here's the good payload! ..."))
	fmt.Println("status:", rr.Result().Status, "body:", rr.Body.String())

	// Output:
	// status: 500 Internal Server Error body: there was an error processing your request
	//
	// status: 500 Internal Server Error body: there was an error processing your request
	// Here's the good payload! ...
}

func ExampleGet() {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte(`<genericRecord><attrListMap><attrList attrCode="SSN"><attr><fieldMap><field name="idnumber"><value xsi:type="xs:string" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xs="http://www.w3.org/2001/XMLSchema">555667777</value></field></fieldMap></attr></attrList><attrList attrCode="PATNAME"><attr><fieldMap><field name="onmFirst"><value xsi:type="xs:string" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xs="http://www.w3.org/2001/XMLSchema">Mario</value></field></fieldMap></attr></attrList><attrList attrCode="BIRTHDT"><attr><fieldMap><field name="dateval"><value xsi:type="xs:string" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xs="http://www.w3.org/2001/XMLSchema">1912-02-14</value></field></fieldMap></attr></attrList><attrList attrCode="HOMEPHON"><attr><fieldMap><field name="phnumber"><value xsi:type="xs:string" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xs="http://www.w3.org/2001/XMLSchema">101-123-7654</value></field></fieldMap></attr></attrList></attrListMap></genericRecord>`))
		case http.MethodPost:
			// noop 😹
		}
	}))

	resp, err := http.Get(svr.URL)
	if err != nil { // How da heyll?
		panic(err)
	}
	defer resp.Body.Close()

	type genRec struct {
		XMLName     xml.Name `xml:"genericRecord"`
		AttrListMap struct {
			XMLName  xml.Name `xml:"attrListMap"`
			AttrList []struct {
				XMLName  xml.Name `xml:"attrList"`
				AttrCode string   `xml:"attrCode"`
				Attr     struct {
					XMLName  xml.Name `xml:"attr"`
					FieldMap struct {
						XMLName xml.Name `xml:"fieldMap"`
						Field   struct {
							XMLName xml.Name `xml:"field"`
							Name    string   `xml:"name,attr"`
							Value   string   `xml:"value"`
						} `xml:"field"`
					}
				}
			} `xml:"attrList"`
		}
	}

	data := genRec{}
	xml.NewDecoder(resp.Body).Decode(&data)
	for _, v := range data.AttrListMap.AttrList {
		fmt.Println("isn't XML great:", v.Attr.FieldMap.Field.Name, v.Attr.FieldMap.Field.Value)
	}
	// Output:
	// isn't XML great: idnumber 555667777
	// isn't XML great: onmFirst Mario
	// isn't XML great: dateval 1912-02-14
	// isn't XML great: phnumber 101-123-7654
}

func ExampleHandle() {
	// NOTE(jay): Compare with [http.HandleFunc]
	inSomeOtherFile := func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("or we can declare a handler from a different file and pass it in."))
		}
	}
	// NOTE(jay): Ordered by specificity. Meaning if a pattern cannot be found it will go
	// down ⤵  the list to find a general enough matching pattern.
	http.Handle("short.localhost/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("we can inline a http.Handler"))
	}))
	http.Handle("/short", inSomeOtherFile())
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("I catch 🪝 EVERYTHING.🤔 Seems like a good place for a 404 page"))
	}))

	go http.ListenAndServe(":9001", nil)

	cpbody := func(url string) {
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		io.Copy(os.Stdout, resp.Body)
		fmt.Println()
	}
	cpbody("http://short.localhost:9001/short")
	cpbody("http://localhost:9001/short")
	cpbody("http://localhost:9001/catch-me")
	// Output:
	// we can inline a http.Handler
	// or we can declare a handler from a different file and pass it in.
	// I catch 🪝 EVERYTHING.🤔 Seems like a good place for a 404 page
}

func ExampleHandle_panic() {
	// NOTE(jay): Compare with [http.HandleFunc]
	defer func() { // to top
		if r := recover(); r != nil {
			fmt.Println("Panic! at the func🕺", r)
		}
	}()
	defer func() { // bottom
		if r := recover(); r != nil {
			fmt.Println("Panic! at the func🕺", r)
			http.Handle("/nil-handler", nil) // 💀❗
		}
	}()
	defer func() { // read from
		if r := recover(); r != nil {
			fmt.Println("Panic! at the func🕺", r)
			http.Handle("/registered", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("already registered 💀❗"))
			}))
		}
	}()

	http.Handle("/registered", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	http.Handle("", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("empty string 💀❗"))
	}))
	// Output:
	// Panic! at the func🕺 http: invalid pattern
	// Panic! at the func🕺 http: multiple registrations for /registered
	// Panic! at the func🕺 http: nil handler
}

func ExampleHandleFunc() {
	// NOTE(jay): Compare with [http.Handle]
	inSomeOtherFile := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("leaving a trailing slash '/' makes it grab everything"))
	}
	// NOTE(jay): Ordered by specificity. Meaning if a pattern cannot be found it will go
	// down ⤵  the list to find a general enough matching pattern.
	http.HandleFunc("smol.localhost/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("we can specify subdomains"))
	})
	http.HandleFunc("/smol/", inSomeOtherFile)
	http.HandleFunc("/smol", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("I only catch what exactly matches."))
	})

	go http.ListenAndServe(":9001", nil)

	cpbody := func(url string) {
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		io.Copy(os.Stdout, resp.Body)
		fmt.Println()
	}
	cpbody("http://smol.localhost:9001/smol")
	cpbody("http://localhost:9001/smol/catch-me")
	cpbody("http://localhost:9001/smol")
	// Output:
	// we can specify subdomains
	// leaving a trailing slash '/' makes it grab everything
	// I only catch what exactly matches.
}

func ExampleHead() {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodHead:
			w.Header().Add("Cache-Control", "Max-Age=3600")
			w.Header().Add("Content-Type", "text/html; charset=utf-8")
			w.Header().Add("Etag", "32e86370dc1c7c634880b37f2273e26408495cc3b3bb47004558e19d490cc8a5")
			w.Header().Add("Last-Modified", "Mon, 17 Oct 2022 22:51:04 GMT")
			w.Header().Add("Strict-Transport-Security", "Max-Age=31556926")
			w.Header().Add("Accept-Ranges", "bytes")
			w.Header().Add("Date", "Wed, 26 Oct 2022 22:56:25 GMT")
			w.Header().Add("X-Served-By", "Cache-Ewr18170-EWR")
			w.Header().Add("X-Cache", "MISS")
			w.Header().Add("X-Cache-Hits", "0")
			w.Header().Add("X-Timer", "S1666824985.282013,VS0,VE92")
			w.Header().Add("Vary", "X-Fh-Requested-Host, Accept-Encoding")
			w.Header().Add("Alt-Svc", `h3=":443";ma=86400,h3-29=":443";ma=86400,h3-27=":443";ma=86400`)
			w.Header().Add("Content-Length", "8543")
			w.WriteHeader(http.StatusOK)
		}
	}))
	resp, err := http.Head(srv.URL)
	fmt.Println("no error:", err)
	defer resp.Body.Close()

	fmt.Println("HEAD contents:")
	for k, v := range resp.Header {
		fmt.Println("key:", k, "value:", v)
	}
	io.Copy(os.Stdout, resp.Body) // This yields nothing.
	// Unordered output:
	// no error: <nil>
	// HEAD contents:
	// key: Content-Type value: [text/html; charset=utf-8]
	// key: Last-Modified value: [Mon, 17 Oct 2022 22:51:04 GMT]
	// key: X-Cache value: [MISS]
	// key: Accept-Ranges value: [bytes]
	// key: Etag value: [32e86370dc1c7c634880b37f2273e26408495cc3b3bb47004558e19d490cc8a5]
	// key: Vary value: [X-Fh-Requested-Host, Accept-Encoding]
	// key: Date value: [Wed, 26 Oct 2022 22:56:25 GMT]
	// key: Content-Length value: [8543]
	// key: X-Served-By value: [Cache-Ewr18170-EWR]
	// key: X-Timer value: [S1666824985.282013,VS0,VE92]
	// key: Alt-Svc value: [h3=":443";ma=86400,h3-29=":443";ma=86400,h3-27=":443";ma=86400]
	// key: Strict-Transport-Security value: [Max-Age=31556926]
	// key: X-Cache-Hits value: [0]
	// key: Cache-Control value: [Max-Age=3600]
}

func ExampleListenAndServe() {
	errs := make(chan error)
	go func() {
		errs <- http.ListenAndServe(":9001", nil)
	}()
	go func() {
		errs <- http.ListenAndServe(":xxxx", nil)
	}()
	go func() {
		errs <- http.ListenAndServe("xxxx", nil)
	}()
	fmt.Println(<-errs)
	fmt.Println(<-errs)
	fmt.Println(<-errs)
	// Unordered output:
	// listen tcp: address xxxx: missing port in address
	// listen tcp :9001: bind: address already in use
	// listen tcp: lookup tcp/xxxx: Servname not supported for ai_socktype
}

// TODO(JAY): COME BACK WHEN SMARTER
// func ExampleListenAndServeTLS() {
// 	// NOTE(jay): Want to make your own server certificate and key? Follow this quick
// 	// tutorial: https://scriptcrunch.com/create-ca-tls-ssl-certificates-keys/
// 	home, err := os.UserHomeDir()
// 	if err != nil {
// 		fmt.Println(err)
// 		panic(err)
// 	}
// 	go http.ListenAndServeTLS(":9002",
// 		path.Join(home, "src", "openssl", "srv.crt"),
// 		path.Join(home, "src", "openssl", "srv.key"),
// 		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			w.Write([]byte("🔒Securely guarded"))
// 		}))
// 	resp, err := http.Get("https://localhost:9002")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("error:", err)
// 	defer resp.Body.Close()
// 	io.Copy(os.Stdout, resp.Body)
// 	// Output:
// }

func ExampleMaxBytesReader() {
	http.HandleFunc("/max-bytes-rdr", func(w http.ResponseWriter, r *http.Request) {
		data := make([]byte, 1<<5)
		n, err := http.MaxBytesReader(w, r.Body, 1<<4).Read(data)
		if mberr, ok := err.(*http.MaxBytesError); ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("could only write " + strconv.Itoa(n) + " bytes: " + mberr.Error()))
		}
	})

	resp, err := http.Post(
		"http://localhost:9001/max-bytes-rdr",
		"application/octet-stream",
		bytes.NewReader(make([]byte, 1<<8))) // 👈 larger than 1<<4
	fmt.Println("request error:", err)
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
	// Output:
	// request error: <nil>
	// could only write 16 bytes: http: request body too large
}

func ExampleNewRequest() {
	// NOTE(jay): run `go doc http.MethodGet` in the terminal to see all methods you can
	// use. Or you can see them at the top of this file.

	// NOTE(jay): It's better practice to be using [http.NewRequestWithContext] as one of
	// the core mechanics of the http package is leveraged heavily off of [context.Context].
	// And in fact if you look at [http.NewRequest] it calls [http.NewRequestWithContext]
	// and passes in [context.Background] as the context.
	req, err := http.NewRequest(http.MethodDelete,
		"https://subdomain.domain.com/path/param?query=string#fragment",
		bytes.NewReader([]byte("Act like there's lots of data in here.")))
	if err != nil {
		panic(err)
	}
	var mt http.Request

	fmt.Printf(`
|             empty            new
Method:         %q              %q
URL:            %s              %q
Proto:          %q              %q
ProtoMajor:     %d              %d
ProtoMinor:     %d              %d
Header:         %v              %v
Body !nil:      %t              %t
Host:           %q              %q`[1:],
		mt.Method, req.Method,
		mt.URL, req.URL,
		mt.Proto, req.Proto,
		mt.ProtoMajor, req.ProtoMajor,
		mt.ProtoMinor, req.ProtoMinor,
		mt.Header, req.Header,
		mt.Body != nil, req.Body != nil,
		mt.Host, req.Host,
	)
	// Output:
	// |             empty            new
	// Method:         ""              "DELETE"
	// URL:            <nil>              "https://subdomain.domain.com/path/param?query=string#fragment"
	// Proto:          ""              "HTTP/1.1"
	// ProtoMajor:     0              1
	// ProtoMinor:     0              1
	// Header:         map[]              map[]
	// Body !nil:      false              true
	// Host:           ""              "subdomain.domain.com"
}

func ExampleNewRequestWithContext() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://gophergo.dev", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(req.Context())
	// Output:
	// context.Background.WithCancel
}

func ExampleNewRequestWithContext_errors() {
	// NOTE(jay): All errors can be returned for [http.NewRequest] as well, besides the one
	// concerning context, of course.
	urlerrs := func(ctx context.Context, method, url string, body io.Reader) {
		_, err := http.NewRequestWithContext(ctx, method, url, body)
		fmt.Println(err)
	}
	urlerrs(context.Background(), "/GET", "https://gophergo.dev", nil)
	urlerrs(context.Background(), "GET", "https://gophergo.dev", nil)
	urlerrs(nil, "GET", "https://gophergo.dev", nil)
	// Output:
	// net/http: invalid method "/GET"
	// parse "\rhttps://gophergo.dev": net/url: invalid control character in URL
	// net/http: nil Context
}

func ExampleNotFound() {
	http.HandleFunc("/not-found", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("I was found!"))
	})
	http.HandleFunc("/not-found/", func(w http.ResponseWriter, r *http.Request) {
		// NOTE(jay): This would be better on the pattern "/" as that's a great to catch
		// literally everything, but it was already used in the ExampleHandle func.
		http.NotFound(w, r)
	})
	resp, err := http.Get("http://localhost:9001/not-found/does-not-exist")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
	// Output:
	// 404 page not found
}

func ExampleParseHTTPVersion() {
	parsed := func(s string) {
		major, minor, ok := http.ParseHTTPVersion(s)
		fmt.Printf("major: %v, minor: %v, ok? %t\n", major, minor, ok)
	}
	parsed("HTTP/1.1")
	parsed("HTTP/2") // XXX(jay): Needs minor version
	parsed("HTTP/2.0")
	// Output:
	// major: 1, minor: 1, ok? true
	// major: 0, minor: 0, ok? false
	// major: 2, minor: 0, ok? true
}

func ExampleParseTime() {
	// NOTE(jay): For an API that isn't just specific to HTTP time formats look into
	// [time.Parse] and [time.ParseInLocation] and [time.Layout] for a full list of already
	// made time format layouts to use.
	http.HandleFunc("/time-plz", func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Header.Get("Date") != "":
			t, err := http.ParseTime(r.Header.Get("Date"))
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			w.Write([]byte(t.String()))
		}
	})
	tyme := func(date string) {
		req, err := http.NewRequestWithContext(
			context.TODO(), http.MethodOptions, "http://localhost:9001/time-plz", nil)
		if err != nil {
			panic(err)
		}
		req.Header.Set(http.CanonicalHeaderKey("date"), date)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		fmt.Print(date, " parsed to\t")
		io.Copy(os.Stdout, resp.Body)
		fmt.Println()
		resp.Body.Close()
	}
	tyme("Fri, 28 Oct 2022 16:32:55 GMT")  // [http.TimeFormat] RFC5322 -- Preferred!
	tyme("Mon Jun 23 16:15:04 2053")       // ANSIC -- Deprecated
	tyme("Monday, 23-Jun-53 16:15:04 CDT") // RFC850 -- Deprecated
	tyme("2053-06-23T16:15:04-05:00")      // RFC3339 -- ERROR
	// Output:
	// Fri, 28 Oct 2022 16:32:55 GMT parsed to	2022-10-28 16:32:55 +0000 UTC
	// Mon Jun 23 16:15:04 2053 parsed to	2053-06-23 16:15:04 +0000 UTC
	// Monday, 23-Jun-53 16:15:04 CDT parsed to	2053-06-23 16:15:04 +0000 CDT
	// 2053-06-23T16:15:04-05:00 parsed to	parsing time "2053-06-23T16:15:04-05:00" as "Mon Jan _2 15:04:05 2006": cannot parse "2053-06-23T16:15:04-05:00" as "Mon"
}

func ExamplePost() {
	http.HandleFunc("/base64", func(w http.ResponseWriter, r *http.Request) {
		jdata, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		enc := base64.NewEncoder(base64.StdEncoding, w)
		enc.Write(jdata)
		enc.Close()
	})
	data := []byte(`{"someData":"to be base64 encoded"}`)
	fmt.Println("base64 encoded data looks like:")
	resp, err := http.Post("http://localhost:9001/base64", "application/json", bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
	// Output:
	// base64 encoded data looks like:
	// eyJzb21lRGF0YSI6InRvIGJlIGJhc2U2NCBlbmNvZGVkIn0=
}

func ExamplePostForm() {
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		if !r.Form.Has("grant_type") || !r.Form.Has("client_id") || !r.Form.Has("client_secret") {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("This server cannot do anything for such a request."))
			return
		}
		// Do some really cool bcrypting or something....
		w.Write([]byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.w8lhMhWeaCAkODcic4tMmiXq_Ym7Zm9a6aBQGepKlAo"))
	})
	resp, err := http.PostForm("http://localhost:9001/login", url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {"my_super_cool_id"},
		"client_secret": {"p29gMFOxLKEu"},
	})
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
	// Output:
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.w8lhMhWeaCAkODcic4tMmiXq_Ym7Zm9a6aBQGepKlAo
}

func TestMain(m *testing.M) {
	os.Setenv("HTTPS_PROXY", "http://localhost:9001/sproxy")
	os.Exit(m.Run())
}

func ExampleProxyFromEnvironment() {
	// NOTE(jay): See TestMain for setting the necessary environment variables.
	pxy := func(url string) {
		req, err := http.NewRequest(http.MethodTrace, url, nil)
		if err != nil {
			panic(err)
		}
		u, err := http.ProxyFromEnvironment(req)
		fmt.Println("URL:", u, "error:", err)
	}
	pxy("https://gophergo.dev")
	pxy("http://gophergo.dev")    // If nothing is defined in Env: <nil> <nil>
	pxy("https://localhost:9001") // Special case for localhost
	// Output:
	// URL: http://localhost:9001/sproxy error: <nil>
	// URL: <nil> error: <nil>
	// URL: <nil> error: <nil>
}

// TODO(jay): Maybe comeback to this.
func ExampleProxyURL() {
	u, err := url.Parse("https://fixedURL.com/never/changes?only=for#Transport")
	if err != nil {
		panic(err)
	}
	t := http.Transport{
		Proxy: http.ProxyURL(u),
	}
	fmt.Println(t.Proxy(httptest.NewRequest("", "/", nil)))
	// Output:
	// https://fixedURL.com/never/changes?only=for#Transport <nil>
}

func ExampleReadRequest() {
	// NOTE(jay): Much like a JSON/XML/PDF/Database Row/Document/GRPC payload isn't a Go
	// struct in any form, neither is a HTTP Request, but we all know they have a
	// specification (RFC9110). Therefore we can think of [http.ReadRequest] like we do
	// about [json.Unmarshal], to turn something foreign into something Go.
	req, err := http.ReadRequest(bufio.NewReader(bytes.NewReader([]byte(`
GET /fun-with-funcs?fun=yes HTTP/1.1
Host: gophergo.dev
User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:106.0) Gecko/20100101 Firefox/106.0
Accept: */*
Accept-Language: en-US,en;q=0.5
Accept-Encoding: gzip, deflate, br
Content-Type: text/plain;charset=UTF-8
Content-Length: 2255
Origin: https://gophergo.dev
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: same-site
Authorization: Bearer Q5USI2pTti09BXCBFI4bsR_59LEYiEQOIHuJ6PEAIF8Wlby89cwD8uObJdTMYP-neNjLpzVhMxz6ADfT9FRp2HA0_-HoZYOP9Z5a-8HdnGVtRVAOjDgin
Referer: https://gophergo.dev/
Connection: keep-alive
TE: trailers

`[1:]))))
	fmt.Printf("Method: %q URL: %q Proto: %q Major.Minor: %d.%d\nHost: %q\nContent Length: %d\nURI: %q\nHeaders: %v\n",
		req.Method, req.URL, req.Proto, req.ProtoMajor, req.ProtoMinor,
		req.Host,
		req.ContentLength,
		req.RequestURI,
		req.Header) // XXX(jay): Maps are unorderd the Output may fail due to that.
	req, err = http.ReadRequest(bufio.NewReader(bytes.NewReader([]byte(`
POST /breaks/because/protocol HTTP/2
Host: gophergo.dev

`[1:]))))
	fmt.Println(req, err)
	// Output:
	// Method: "GET" URL: "/fun-with-funcs?fun=yes" Proto: "HTTP/1.1" Major.Minor: 1.1
	// Host: "gophergo.dev"
	// Content Length: 2255
	// URI: "/fun-with-funcs?fun=yes"
	// Headers: map[Accept:[*/*] Accept-Encoding:[gzip, deflate, br] Accept-Language:[en-US,en;q=0.5] Authorization:[Bearer Q5USI2pTti09BXCBFI4bsR_59LEYiEQOIHuJ6PEAIF8Wlby89cwD8uObJdTMYP-neNjLpzVhMxz6ADfT9FRp2HA0_-HoZYOP9Z5a-8HdnGVtRVAOjDgin] Connection:[keep-alive] Content-Length:[2255] Content-Type:[text/plain;charset=UTF-8] Origin:[https://gophergo.dev] Referer:[https://gophergo.dev/] Sec-Fetch-Dest:[empty] Sec-Fetch-Mode:[cors] Sec-Fetch-Site:[same-site] Te:[trailers] User-Agent:[Mozilla/5.0 (X11; Linux x86_64; rv:106.0) Gecko/20100101 Firefox/106.0]]
	// <nil> malformed HTTP version "HTTP/2"
}

func ExampleReadResponse() {
	// NOTE(jay): Much like a JSON/XML/PDF/Database Row/Document/GRPC payload isn't a Go
	// struct in any form, neither is a HTTP Response, but we all know they have a
	// specification (RFC9110). Therefore we can think of [http.ReadResponse] like we do
	// about [json.Unmarshal], to turn something foreign into something Go.
	resp, _ := http.ReadResponse(bufio.NewReader(bytes.NewReader([]byte(`
HTTP/1.1 206 Partial Content
Connection: keep-alive
Content-Length: 325659
Last-Modified: Wed, 13 Apr 2022 17:27:05 GMT
ETag: "29e0066f9f18a7832e052fa9ff8ad61c"
X-gopher-generation: 1649870825153460
X-gopher-metageneration: 1
X-gopher-stored-content-encoding: identity
X-gopher-stored-content-length: 7820224
Content-Type: application/octet-stream
Accept-Ranges: bytes
Age: 70
Content-Range: bytes 2277194-2602852/7820224
Date: Sat, 29 Oct 2022 14:41:03 GMT
X-Served-By: cache-chi-kigq8000164-CHI, cache-iad-kiad7000137-IAD
X-Cache: MISS, HIT
X-Cache-Hits: 0, 0
Access-Control-Allow-Origin: *
Cache-Control: max-age=315360000, no-transform

AAAU5W1vb2YAAAAQbWZoZAAAAAAAAAAKAAAUzXRyYWYAAAAcdGZoZAACACoAAAABAAAAAQAABAAAAAAAAAAAEHRmZHQAAAAAADyYAAAABtB0cnVuAAACAQAAAa8AABTtAAAC3gAAAtcAAALYAAAC3wAAAugAAALkAAAC5AAAAusAAALlAAAC8QAAA7sAAAKzAAACsgAA
AAAU5W1vb2YAAAAQbWZoZAAAAAAAAAAKAAAUzXRyYWYAAAAcdGZoZAACACoAAAABAAAAAQAABAAAAAAAAAAAEHRmZHQAAAAAADyYAAAABtB0cnVuAAACAQAAAa8AABTtAAAC3gAAAtcAAALYAAAC3wAAAugAAALkAAAC5AAAAusAAALlAAAC8QAAA7sAAAKzAAACsgAA
This is the body of the response.
AAAU5W1vb2YAAAAQbWZoZAAAAAAAAAAKAAAUzXRyYWYAAAAcdGZoZAACACoAAAABAAAAAQAABAAAAAAAAAAAEHRmZHQAAAAAADyYAAAABtB0cnVuAAACAQAAAa8AABTtAAAC3gAAAtcAAALYAAAC3wAAAugAAALkAAAC5AAAAusAAALlAAAC8QAAA7sAAAKzAAACsgAA
AAAU5W1vb2YAAAAQbWZoZAAAAAAAAAAKAAAUzXRyYWYAAAAcdGZoZAACACoAAAABAAAAAQAABAAAAAAAAAAAEHRmZHQAAAAAADyYAAAABtB0cnVuAAACAQAAAa8AABTtAAAC3gAAAtcAAALYAAAC3wAAAugAAALkAAAC5AAAAusAAALlAAAC8QAAA7sAAAKzAAACsgAA

`[1:]))), nil)
	fmt.Printf("Proto: %q Status: %q\nContent-Length: %d\nHeaders: %v\nBody:\n",
		resp.Proto, resp.Status,
		resp.ContentLength,
		resp.Header)
	io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()

	req, _ := http.NewRequest(http.MethodPatch, "https://gophergo.dev/", nil)
	resp, _ = http.ReadResponse(bufio.NewReader(bytes.NewReader([]byte(`
HTTP/1.1 204 No Content
content-length: 242471
vary: Origin
cross-origin-resource-policy: cross-origin
x-content-type-options: nosniff
server: gvs 1.0

`[1:]))), req)
	fmt.Printf("Request Method: %q and URL: %q\n", resp.Request.Method, resp.Request.URL)
	// Output:
	// Proto: "HTTP/1.1" Status: "206 Partial Content"
	// Content-Length: 325659
	// Headers: map[Accept-Ranges:[bytes] Access-Control-Allow-Origin:[*] Age:[70] Cache-Control:[max-age=315360000, no-transform] Connection:[keep-alive] Content-Length:[325659] Content-Range:[bytes 2277194-2602852/7820224] Content-Type:[application/octet-stream] Date:[Sat, 29 Oct 2022 14:41:03 GMT] Etag:["29e0066f9f18a7832e052fa9ff8ad61c"] Last-Modified:[Wed, 13 Apr 2022 17:27:05 GMT] X-Cache:[MISS, HIT] X-Cache-Hits:[0, 0] X-Gopher-Generation:[1649870825153460] X-Gopher-Metageneration:[1] X-Gopher-Stored-Content-Encoding:[identity] X-Gopher-Stored-Content-Length:[7820224] X-Served-By:[cache-chi-kigq8000164-CHI, cache-iad-kiad7000137-IAD]]
	// Body:
	// AAAU5W1vb2YAAAAQbWZoZAAAAAAAAAAKAAAUzXRyYWYAAAAcdGZoZAACACoAAAABAAAAAQAABAAAAAAAAAAAEHRmZHQAAAAAADyYAAAABtB0cnVuAAACAQAAAa8AABTtAAAC3gAAAtcAAALYAAAC3wAAAugAAALkAAAC5AAAAusAAALlAAAC8QAAA7sAAAKzAAACsgAA
	// AAAU5W1vb2YAAAAQbWZoZAAAAAAAAAAKAAAUzXRyYWYAAAAcdGZoZAACACoAAAABAAAAAQAABAAAAAAAAAAAEHRmZHQAAAAAADyYAAAABtB0cnVuAAACAQAAAa8AABTtAAAC3gAAAtcAAALYAAAC3wAAAugAAALkAAAC5AAAAusAAALlAAAC8QAAA7sAAAKzAAACsgAA
	// This is the body of the response.
	// AAAU5W1vb2YAAAAQbWZoZAAAAAAAAAAKAAAUzXRyYWYAAAAcdGZoZAACACoAAAABAAAAAQAABAAAAAAAAAAAEHRmZHQAAAAAADyYAAAABtB0cnVuAAACAQAAAa8AABTtAAAC3gAAAtcAAALYAAAC3wAAAugAAALkAAAC5AAAAusAAALlAAAC8QAAA7sAAAKzAAACsgAA
	// AAAU5W1vb2YAAAAQbWZoZAAAAAAAAAAKAAAUzXRyYWYAAAAcdGZoZAACACoAAAABAAAAAQAABAAAAAAAAAAAEHRmZHQAAAAAADyYAAAABtB0cnVuAAACAQAAAa8AABTtAAAC3gAAAtcAAALYAAAC3wAAAugAAALkAAAC5AAAAusAAALlAAAC8QAAA7sAAAKzAAACsgAA
	//
	// Request Method: "PATCH" and URL: "https://gophergo.dev/"
}

type PeekWriter struct{ http.ResponseWriter }

func ExampleRedirect() {
	http.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
		rr := httptest.NewRecorder()
		// [http.StatusSeeOther] is also viable.
		http.Redirect(rr, r, "http://localhost:9001/to", http.StatusMovedPermanently)
		http.Redirect(w, r, "http://localhost:9001/to", http.StatusFound)
		fmt.Println(rr.Header())
		fmt.Println(rr.Body.String()[:len(rr.Body.String())-1]) // remove trailing \n
	})
	http.HandleFunc("/to", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("r.Header: %v\n", r.Header)
		w.Write([]byte("Successfully redirected!"))
	})

	resp, err := http.Get("http://localhost:9001/redirect")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
	fmt.Println()

	rr := httptest.NewRecorder()
	rr.Header().Set("Content-Type", "")
	http.Redirect(rr, httptest.NewRequest("", "/", nil), "/", http.StatusSeeOther)
	fmt.Println(rr.Body.String()) // Nothing is written if `Content-Type` is set already.
	// Output:
	// map[Content-Type:[text/html; charset=utf-8] Location:[http://localhost:9001/to]]
	// <a href="http://localhost:9001/to">Moved Permanently</a>.
	//
	// r.Header: map[Accept-Encoding:[gzip] Referer:[http://localhost:9001/redirect] User-Agent:[Go-http-client/1.1]]
	// Successfully redirected!
}

func ExampleServe() {
	l, err := net.Listen("tcp", ":9002")
	if err != nil {
		panic(err)
	}

	go http.Serve(l, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO(jay): 
	}))
	// Output:
}

// func ExampleServeContent() {
// 	http.ServeContent()
// 	// Output:
// }
//
// func ExampleServeFile() {
// 	http.ServeFile()
// 	// Output:
// }
//
// func ExampleServeTLS() {
// 	http.ServeTLS()
// 	// Output:
// }
//
// func ExampleSetCookie() {
// 	http.SetCookie()
// 	// Output:
// }
//
// func ExampleStatusText() {
// 	http.StatusText()
// 	// Output:
// }
//
// // const MethodGet = "GET" ...
// // const StatusContinue = 100 ...
// // const DefaultMaxHeaderBytes = 1 << 20
// // const DefaultMaxIdleConnsPerHost = 2
// // const TimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"
// // const TrailerPrefix = "Trailer:"
// // var ErrNotSupported = &ProtocolError{ ... } ...
// // var ErrBodyNotAllowed = errors.New("http: request method or response status code does not allow body") ...
// // var ServerContextKey = &contextKey{ ... } ...
// // var DefaultClient = &Client{}
// // var DefaultServeMux = &defaultServeMux
// // var ErrAbortHandler = errors.New("net/http: abort Handler")
// // var ErrBodyReadAfterClose = errors.New("http: invalid Read on closed Body")
// // var ErrHandlerTimeout = errors.New("http: Handler timeout")
// // var ErrLineTooLong = internal.ErrLineTooLong
// // var ErrMissingFile = errors.New("http: no such file")
// // var ErrNoCookie = errors.New("http: named cookie not present")
// // var ErrNoLocation = errors.New("http: no Location header in response")
// // var ErrServerClosed = errors.New("http: Server closed")
// // var ErrSkipAltProtocol = errors.New("net/http: skip alternate protocol")
// // var ErrUseLastResponse = errors.New("net/http: use last response")
// // var NoBody = noBody{}
// // type Client struct{ ... }
// // type CloseNotifier interface{ ... }
// // type ConnState int
// //     const StateNew ConnState = iota ...
// // type Cookie struct{ ... }
// // type CookieJar interface{ ... }
// // type Dir string
// // type File interface{ ... }
// // type FileSystem interface{ ... }
// //     func FS(fsys fs.FS) FileSystem
// // type Flusher interface{ ... }
// // type Handler interface{ ... }
// //     func AllowQuerySemicolons(h Handler) Handler
// //     func FileServer(root FileSystem) Handler
// //     func MaxBytesHandler(h Handler, n int64) Handler
// //     func NotFoundHandler() Handler
// //     func RedirectHandler(url string, code int) Handler
// //     func StripPrefix(prefix string, h Handler) Handler
// //     func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler
// // type HandlerFunc func(ResponseWriter, *Request)
// // type Header map[string][]string
// // type Hijacker interface{ ... }
// // type MaxBytesError struct{ ... }
// // type ProtocolError struct{ ... }
// // type PushOptions struct{ ... }
// // type Pusher interface{ ... }
// // type Request struct{ ... }
// // type Response struct{ ... }
// // type ResponseWriter interface{ ... }
// // type RoundTripper interface{ ... }
// //     var DefaultTransport RoundTripper = &Transport{ ... }
// //     func NewFileTransport(fs FileSystem) RoundTripper
// // type SameSite int
// //     const SameSiteDefaultMode SameSite = iota + 1 ...
// // type ServeMux struct{ ... }
// //     func NewServeMux() *ServeMux
// // type Server struct{ ... }
// // type Transport struct{ ... }
