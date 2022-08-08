package main

import (
	"context"
	"flag"
	"net/http"
	"os"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptr "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	var (
		port  = flag.String("port", ":9001", "HTTP port")
		proxy = flag.String("proxy", "",
			"Optional ocmma-separated list of URLS to proxy uppercase requests")
	)
	flag.Parse()

	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "port", *port, "caller", log.DefaultCaller)
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{})

	var svc StringService
	svc = proxyingMiddleware(context.Background(), *proxy, logger)(svc)
	svc = stringService{}
	svc = loggingMiddleware{logger, svc}
	svc = instrumentingMiddleware{requestCount, requestLatency, countResult, svc}

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
	http.Handle("/metrics", promhttp.Handler())
	logger.Log("msg", "HTTP", "addr", *port)
	logger.Log("err", http.ListenAndServe(*port, nil))
}
