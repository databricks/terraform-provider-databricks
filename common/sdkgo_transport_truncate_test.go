package common

import (
	"context"
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// debug_truncate_bytes shortens the values of a logged body in the debug line of both
// clients. The default of 96 bytes would log this volume's values whole.
func TestSdkGoClientTruncatesLoggedBodiesLikeClassicClient(t *testing.T) {
	sdkGoLogger, classicLogger := &capturedLogger{}, &capturedLogger{}
	// readBothWays reads through sdk-go first.
	next := []*capturedLogger{sdkGoLogger, classicLogger}
	newContext := func() context.Context {
		l := next[0]
		next = next[1:]
		return logger.NewContext(context.Background(), l)
	}
	truncate := func(cfg *config.Config) { cfg.DebugTruncateBytes = 4 }
	r := readBothWays(t, newContext, truncate, respond(http.StatusOK, sdkGoTestVolume))
	require.NoError(t, r.err)
	require.NoError(t, r.classicErr)
	read := "DEBUG GET " + sdkGoTestVolumePath
	sdkGoLines := sdkGoLogger.matching(read, `"full_name": "main... (14 more bytes)"`)
	require.Len(t, sdkGoLines, 1, "%q", sdkGoLogger.lines)
	assert.Equal(t, classicLogger.matching(read), sdkGoLines)
}
