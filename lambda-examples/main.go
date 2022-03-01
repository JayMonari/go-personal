package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"
)

type Event struct {
	Name string `json:"name"`
}

var logger *zap.Logger

func init() {
	l, _ := zap.NewProduction()
	logger = l
	defer logger.Sync()
}

func Handler(ctx context.Context, e Event) error {
	logger.Info("in the handler", zap.Any("event", e))
	return nil
}

func main() {
	lambda.Start(Handler)
}
