package main

import (
	"context"
	"transaction-service/pkg/logger"
	"transaction-service/pkg/server"
)

func main() {
	ctx := context.Background()

	// Example without requestId
	logger.WithCtx(ctx).Info("First log without requestId")

	// Example with requestId
	ctxWithRequestId := context.WithValue(ctx, "requestId", "12345")
	logger.WithCtx(ctxWithRequestId).Error("Message")

	server.Start(ctx)
}
