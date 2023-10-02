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

func TestConvertToMap(t *testing.T) {
	m := convertToMap("value1", "value2")
	assert.Equal(t, "value1", m["0"])
	assert.Equal(t, "value2", m["1"])
}

func TestSetLogger(t *testing.T) {
	SetLogger()
	assert.IsType(t, &TfLogger{}, goLogger.DefaultLogger)
}
