package logger

import (
	"context"
	"log"
)

type SimpleLogger struct {
	Level Level
}

func (l *SimpleLogger) Enabled(_ context.Context, level Level) bool {
	return level >= l.Level
}

func (l *SimpleLogger) Tracef(ctx context.Context, format string, v ...any) {
	if !l.Enabled(ctx, LevelTrace) {
		return
	}
	log.Printf("[TRACE] "+format, v...)
}

func (l *SimpleLogger) Debugf(ctx context.Context, format string, v ...any) {
	if !l.Enabled(ctx, LevelDebug) {
		return
	}
	log.Printf("[DEBUG] "+format, v...)
}

func (l *SimpleLogger) Infof(ctx context.Context, format string, v ...any) {
	if !l.Enabled(ctx, LevelInfo) {
		return
	}
	log.Printf("[INFO] "+format, v...)
}

func (l *SimpleLogger) Warnf(ctx context.Context, format string, v ...any) {
	if !l.Enabled(ctx, LevelWarn) {
		return
	}
	log.Printf("[WARN] "+format, v...)
}

func (l *SimpleLogger) Errorf(ctx context.Context, format string, v ...any) {
	if !l.Enabled(ctx, LevelError) {
		return
	}
	log.Printf("[ERROR] "+format, v...)
}
