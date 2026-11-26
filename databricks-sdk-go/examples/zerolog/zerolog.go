package main

import (
	"context"

	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/rs/zerolog"
)

// zerologAdapter makes an zerolog.Logger usable with the Databricks SDK.
type zerologAdapter struct {
	zerolog.Logger
}

func (z *zerologAdapter) Enabled(_ context.Context, level logger.Level) bool {
	switch level {
	case logger.LevelTrace:
		return z.Logger.GetLevel() <= zerolog.TraceLevel
	case logger.LevelDebug:
		return z.Logger.GetLevel() <= zerolog.DebugLevel
	case logger.LevelInfo:
		return z.Logger.GetLevel() <= zerolog.InfoLevel
	case logger.LevelWarn:
		return z.Logger.GetLevel() <= zerolog.WarnLevel
	case logger.LevelError:
		return z.Logger.GetLevel() <= zerolog.ErrorLevel
	default:
		return true
	}
}

func (s *zerologAdapter) Tracef(_ context.Context, format string, v ...any) {
	s.Logger.Trace().Msgf(format, v...)
}

func (s *zerologAdapter) Debugf(_ context.Context, format string, v ...any) {
	s.Logger.Debug().Msgf(format, v...)
}

func (s *zerologAdapter) Infof(_ context.Context, format string, v ...any) {
	s.Logger.Info().Msgf(format, v...)
}

func (s *zerologAdapter) Warnf(_ context.Context, format string, v ...any) {
	s.Logger.Warn().Msgf(format, v...)
}

func (s *zerologAdapter) Errorf(_ context.Context, format string, v ...any) {
	s.Logger.Error().Msgf(format, v...)
}
