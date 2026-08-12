package exporter

import (
	"bytes"
	"context"
	"log"
	"testing"

	sdklogger "github.com/databricks/databricks-sdk-go/logger"
	"github.com/stretchr/testify/assert"
)

func TestDetermineExporterLogLevel(t *testing.T) {
	tests := []struct {
		name        string
		debug       bool
		trace       bool
		tfLog       string
		wantName    string
		wantSDK     sdklogger.Level
		wantPrefix  []string
		wantWarning string
	}{
		{
			name:       "default info",
			wantName:   "INFO",
			wantSDK:    sdklogger.LevelInfo,
			wantPrefix: []string{"[INFO]", "[WARN]", "[ERROR]"},
		},
		{
			name:       "tf log error",
			tfLog:      "ERROR",
			wantName:   "ERROR",
			wantSDK:    sdklogger.LevelError,
			wantPrefix: []string{"[ERROR]"},
		},
		{
			name:       "tf log warn",
			tfLog:      "warn",
			wantName:   "WARN",
			wantSDK:    sdklogger.LevelWarn,
			wantPrefix: []string{"[WARN]", "[ERROR]"},
		},
		{
			name:       "tf log debug",
			tfLog:      "DEBUG",
			wantName:   "DEBUG",
			wantSDK:    sdklogger.LevelDebug,
			wantPrefix: []string{"[DEBUG]", "[INFO]", "[WARN]", "[ERROR]"},
		},
		{
			name:       "tf log trace",
			tfLog:      "TRACE",
			wantName:   "TRACE",
			wantSDK:    sdklogger.LevelTrace,
			wantPrefix: []string{"[TRACE]", "[DEBUG]", "[INFO]", "[WARN]", "[ERROR]"},
		},
		{
			name:       "debug overrides tf log",
			debug:      true,
			tfLog:      "ERROR",
			wantName:   "DEBUG",
			wantSDK:    sdklogger.LevelDebug,
			wantPrefix: []string{"[DEBUG]", "[INFO]", "[WARN]", "[ERROR]"},
		},
		{
			name:       "trace overrides debug",
			debug:      true,
			trace:      true,
			tfLog:      "ERROR",
			wantName:   "TRACE",
			wantSDK:    sdklogger.LevelTrace,
			wantPrefix: []string{"[TRACE]", "[DEBUG]", "[INFO]", "[WARN]", "[ERROR]"},
		},
		{
			name:        "invalid tf log",
			tfLog:       "verbose",
			wantName:    "INFO",
			wantSDK:     sdklogger.LevelInfo,
			wantPrefix:  []string{"[INFO]", "[WARN]", "[ERROR]"},
			wantWarning: `Invalid TF_LOG value "verbose"; defaulting exporter logging to INFO. Valid values are: TRACE, DEBUG, INFO, WARN, ERROR`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, warning := determineExporterLogLevel(tc.debug, tc.trace, tc.tfLog)
			assert.Equal(t, tc.wantName, got.name)
			assert.Equal(t, tc.wantSDK, got.sdkLevel)
			assert.Equal(t, tc.wantPrefix, got.prefixes)
			assert.Equal(t, tc.wantWarning, warning)
		})
	}
}

func TestConfigureExporterLoggingConfiguresSdkLogger(t *testing.T) {
	previousOutput := log.Writer()
	previousLogger := sdklogger.DefaultLogger
	t.Cleanup(func() {
		log.SetOutput(previousOutput)
		sdklogger.DefaultLogger = previousLogger
	})

	t.Setenv("TF_LOG", "DEBUG")
	level := configureExporterLogging(false, false)

	assert.Equal(t, "DEBUG", level.name)
	assert.True(t, sdklogger.DefaultLogger.Enabled(context.Background(), sdklogger.LevelDebug))
	assert.False(t, sdklogger.DefaultLogger.Enabled(context.Background(), sdklogger.LevelTrace))
}

func TestLevelWriterFiltersAndFormatsMessages(t *testing.T) {
	var buf bytes.Buffer
	writer := &levelWriter{
		prefixes: []string{"[INFO]"},
		output:   &buf,
	}

	message := []byte("[DEBUG] hidden\n")
	n, err := writer.Write(message)
	assert.NoError(t, err)
	assert.Equal(t, len(message), n)
	assert.Empty(t, buf.String())

	message = []byte("[INFO] visible\n")
	n, err = writer.Write(message)
	assert.NoError(t, err)
	assert.Equal(t, len(message), n)
	assert.Contains(t, buf.String(), "[INFO] ")
	assert.Contains(t, buf.String(), ": visible\n")
}
