package logger

import (
	"context"
	"github.com/sirupsen/logrus"
	"transaction-service/pkg/constants"
)

var log = logrus.New()

// LogWithRequestId logs a message with the requestId if it exists in the context
func LogWithRequestId(ctx context.Context) *logrus.Entry {
	if requestId, ok := ctx.Value(constants.RequestId).(string); ok {
		return log.WithField(constants.RequestId, requestId)
	}

	return log.WithContext(ctx)
}
