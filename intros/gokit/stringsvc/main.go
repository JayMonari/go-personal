package main

import (
	"net/http"
	"os"

	httptr "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	var svc StringService
	svc = stringService{}
	svc = loggingMiddleware{logger, svc}
	uppercaseHandler := httptr.NewServer(
		makeUpperCaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	)
	countHandler := httptr.NewServer(
		makeCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	http.ListenAndServe(":9001", nil)
}
