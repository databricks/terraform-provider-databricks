package logger

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLogger(t *testing.T) {
	t.Cleanup(func() {
		DefaultLogger = &SimpleLogger{}
	})

	t1 := &SimpleLogger{}
	t2 := &SimpleLogger{}
	DefaultLogger = t1

	logger := Get(context.Background())
	assert.Equal(t, logger, t1)

	ctx := NewContext(context.Background(), t2)
	logger = Get(ctx)
	assert.Equal(t, logger, t2)
}

func TestWhenInfoLevelThenDebugDisabled(t *testing.T) {
	t.Cleanup(func() {
		DefaultLogger = &SimpleLogger{}
	})

	infoLevelLogger := &SimpleLogger{
		Level: LevelInfo,
	}
	debugEnabled := infoLevelLogger.Enabled(context.Background(), LevelDebug)
	assert.False(t, debugEnabled)
}

func TestWhenInfoLevelThenErrorEnabled(t *testing.T) {
	infoLevelLogger := &SimpleLogger{
		Level: LevelInfo,
	}

	errorEnabled := infoLevelLogger.Enabled(context.Background(), LevelError)
	assert.True(t, errorEnabled)
}

func TestDefaultLevelInfo(t *testing.T) {
	logger := &SimpleLogger{}
	assert.EqualValues(t, LevelInfo, logger.Level)
}
