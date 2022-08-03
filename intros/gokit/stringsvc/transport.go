package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't JSON-marshal
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}

func makeUpperCaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request any) (response any, err error) {
		v, err := svc.Uppercase(request.(uppercaseRequest).S)
		if err != nil {
			return uppercaseResponse{v, err.Error()}, nil
		}
		return uppercaseResponse{v, ""}, nil
	}
}

func makeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request any) (response any, err error) {
		return countResponse{svc.Count(request.(countRequest).S)}, nil
	}
}

func decodeUppercaseRequest(_ context.Context, r *http.Request) (any, error) {
	var req uppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (any, error) {
	var req countRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response any) error {
	return json.NewEncoder(w).Encode(response)
}
