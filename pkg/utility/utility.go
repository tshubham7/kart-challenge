package utility

import (
	"context"
	"os"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type contextKey string

const (
	KeyLogger contextKey = "LOGGER"

	KeyCorrelationId = "CORRELATION_ID"
)

func NewLogger(ctx context.Context) (context.Context, *logrus.Logger) {
	var log = logrus.New()
	log.Formatter = new(logrus.JSONFormatter)
	log.Level = logrus.InfoLevel
	log.Out = os.Stdout

	log = log.WithField(KeyCorrelationId, uuid.NewString()).Logger

	ctx = context.WithValue(ctx, KeyLogger, log)
	return ctx, log
}

func GetCurrentLogger(ctx context.Context) *logrus.Logger {
	if log := ctx.Value(KeyLogger); log != nil {
		return log.(*logrus.Logger)
	}

	_, log := NewLogger(ctx)
	return log
}
