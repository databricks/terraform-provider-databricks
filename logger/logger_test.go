package logger

import (
	"context"
	"testing"

	goLogger "github.com/databricks/databricks-sdk-go/logger"
	"github.com/stretchr/testify/assert"
)

func TestTfLogger_Enabled(t *testing.T) {
	l := &TfLogger{}
	assert.True(t, l.Enabled(context.Background(), goLogger.LevelInfo))
}

func TestSetLogger(t *testing.T) {
	SetLogger(context.Background())
	assert.IsType(t, &TfLogger{}, goLogger.DefaultLogger)
}
