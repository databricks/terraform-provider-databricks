package main

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/logger"
	"golang.org/x/exp/slog"
)

// slogAdapter makes an slog.Logger usable with the Databricks SDK.
type slogAdapter struct {
	*slog.Logger
}

func (s *slogAdapter) Enabled(ctx context.Context, level logger.Level) bool {
	switch level {
	case logger.LevelTrace:
		// Note: slog doesn't have a default trace level.
		// An application can define their own fine grained levels
		// and use those here, if needed.
		return s.Logger.Enabled(ctx, slog.LevelDebug)
	case logger.LevelDebug:
		return s.Logger.Enabled(ctx, slog.LevelDebug)
	case logger.LevelInfo:
		return s.Logger.Enabled(ctx, slog.LevelInfo)
	case logger.LevelWarn:
		return s.Logger.Enabled(ctx, slog.LevelWarn)
	case logger.LevelError:
		return s.Logger.Enabled(ctx, slog.LevelError)
	default:
		return true
	}
}

func (s *slogAdapter) Tracef(ctx context.Context, format string, v ...any) {
	s.DebugContext(ctx, fmt.Sprintf(format, v...))
}

func (s *slogAdapter) Debugf(ctx context.Context, format string, v ...any) {
	s.DebugContext(ctx, fmt.Sprintf(format, v...))
}

func (s *slogAdapter) Infof(ctx context.Context, format string, v ...any) {
	s.InfoContext(ctx, fmt.Sprintf(format, v...))
}

func (s *slogAdapter) Warnf(ctx context.Context, format string, v ...any) {
	s.WarnContext(ctx, fmt.Sprintf(format, v...))
}

func (s *slogAdapter) Errorf(ctx context.Context, format string, v ...any) {
	s.ErrorContext(ctx, fmt.Sprintf(format, v...), nil)
}
