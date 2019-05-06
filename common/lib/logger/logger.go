// Package logger contains some helper functions
// for managing loggers in go applications.
// The package uses the zap logging library.
package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewDevelopment creates a new development zap.Logger instance.
func NewDevelopment() *zap.Logger {
	log, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	return log
}

// NewProduction creates a new production
// zap.Logger instance with a custom timestamp
// field which is compatible with the EFK/ELK stack.
func NewProduction() *zap.Logger {
	cfg := zap.NewProductionEncoderConfig()
	cfg.TimeKey = "@timestamp"
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder

	log := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg),
		zapcore.Lock(os.Stderr),
		zap.NewAtomicLevel(),
	))

	log = log.With(zap.String("lang", "go"))

	return log
}
