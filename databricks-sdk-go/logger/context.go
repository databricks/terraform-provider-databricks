package logger

import "context"

type logger int

var loggerKey logger

// NewContext returns a new Context that carries the specified logger.
func NewContext(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

// FromContext returns the Logger value stored in ctx, if any.
func FromContext(ctx context.Context) (Logger, bool) {
	u, ok := ctx.Value(loggerKey).(Logger)
	return u, ok
}
