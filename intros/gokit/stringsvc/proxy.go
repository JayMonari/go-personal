package main

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	httptrp "github.com/go-kit/kit/transport/http"
	"github.com/sony/gobreaker"
	"golang.org/x/time/rate"
)

// proxymw implments StringService, forwarding Uppercase requests to the
// provided endpoint, and serving all other (i.e. Count) requests via the next
// [StringService].
type proxymw struct {
	ctx       context.Context
	next      StringService     // Serve most request from this endpoint...
	uppercase endpoint.Endpoint // ...except Uppercase
}

func (mw proxymw) Uppercase(s string) (string, error) {
	resp, err := mw.uppercase(context.Background(), uppercaseRequest{S: s})
	if err != nil {
		return "", err
	}
	r := resp.(uppercaseResponse)
	if r.Err != "" {
		return r.V, errors.New(r.Err)
	}
	return r.V, nil
}

func (mw proxymw) Count(s string) int { return mw.next.Count(s) }

func proxyingMiddleware(ctx context.Context, instances string, l log.Logger,
) ServiceMiddleware {
	if instances == "" {
		l.Log("proxy_to", "none")
		return func(next StringService) StringService { return next }
	}
	var (
		qps        = 100                    // beyond which we return an error
		maxAttempt = 3                      // per request
		maxTime    = 250 * time.Millisecond // before giving up
	)
	var (
		instanceSlc = split(instances)
		endpointer  sd.FixedEndpointer
	)
	l.Log("proxy_to", fmt.Sprint(instanceSlc))
	for _, i := range instanceSlc {
		var e endpoint.Endpoint
		e = makeUppercaseProxy(ctx, i)
		e = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(e)
		e = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), qps))(e)
		endpointer = append(endpointer, e)
	}
	retry := lb.Retry(maxAttempt, maxTime, lb.NewRoundRobin(endpointer))
	return func(next StringService) StringService {
		return proxymw{
			ctx:       ctx,
			next:      next,
			uppercase: retry,
		}
	}
}

func makeUppercaseProxy(ctx context.Context, instance string) endpoint.Endpoint {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		panic(err)
	}
	if u.Path == "" {
		u.Path = "/uppercase"
	}
	return httptrp.NewClient(
		"GET",
		u,
		encodeRequest,
		decodeUppercaseResponse,
	).Endpoint()
}

func split(s string) []string {
	a := strings.Split(s, ",")
	for i := range a {
		a[i] = strings.TrimSpace(a[i])
	}
	return a
}
