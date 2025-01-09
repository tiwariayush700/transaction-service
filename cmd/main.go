package main

import (
	"context"
	"transaction-service/pkg/logger"
)

func main() {
	ctx := context.Background()

	// Example without requestId
	logger.LogWithRequestId(ctx).Info("First log without requestId")

	// Example with requestId
	ctxWithRequestId := context.WithValue(ctx, "requestId", "12345")
	logger.LogWithRequestId(ctxWithRequestId).Error("Message")
}
