package middleware

import (
	"context"
	"net/http"
	"transaction-service/pkg/constants"
	"transaction-service/pkg/logger"

	"github.com/google/uuid"
)

// RequestIDMiddleware is a middleware that adds a request ID to the context
func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New().String()
		ctx := context.WithValue(r.Context(), constants.RequestId, requestID)
		r = r.WithContext(ctx)
		logger.WithCtx(ctx).Info("Request started")
		next.ServeHTTP(w, r)
	})
}
