package logger

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoggerContext(t *testing.T) {
	logger, ok := FromContext(context.Background())
	assert.Nil(t, logger)
	assert.False(t, ok)

	ctx := NewContext(context.Background(), &SimpleLogger{})
	logger, ok = FromContext(ctx)
	assert.NotNil(t, logger)
	assert.True(t, ok)
}
