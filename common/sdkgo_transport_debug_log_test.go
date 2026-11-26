package common

import (
	"context"
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	sdkvolumes "github.com/databricks/sdk-go/uc/volumes/v1"
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

// The debug line of a create shows its request body, the same on both clients.
func TestSdkGoClientLogsWriteBodiesLikeClassicClient(t *testing.T) {
	server, _ := serveSequence(t, respond(http.StatusOK, sdkGoTestVolume),
		respond(http.StatusOK, sdkGoTestVolume))
	w := newVolumesTestWorkspaceClient(t, server, nil)
	sdkGoLogger, classicLogger := &capturedLogger{}, &capturedLogger{}
	catalogName, schemaName, name := "main", "replay", "volume"
	vc, err := SdkGoVolumesClient(context.Background(), w)
	require.NoError(t, err)
	_, err = vc.Create(logger.NewContext(context.Background(), sdkGoLogger),
		sdkvolumes.CreateVolumeRequest{CatalogName: &catalogName, SchemaName: &schemaName,
			Name: &name, VolumeType: sdkvolumes.VolumeType_Managed})
	require.NoError(t, err)
	_, err = w.Volumes.Create(logger.NewContext(context.Background(), classicLogger),
		catalog.CreateVolumeRequestContent{CatalogName: catalogName,
			SchemaName: schemaName, Name: name, VolumeType: catalog.VolumeTypeManaged})
	require.NoError(t, err)
	post := "DEBUG POST /api/2.1/unity-catalog/volumes"
	sdkGoLines := sdkGoLogger.matching(post, `"catalog_name": "main"`,
		`"volume_type": "MANAGED"`)
	require.Len(t, sdkGoLines, 1, "%q", sdkGoLogger.lines)
	assert.Equal(t, classicLogger.matching(post), sdkGoLines)
}

// quietLogger records every line it gets, but reports debug logging as off.
type quietLogger struct {
	capturedLogger
}

func (*quietLogger) Enabled(_ context.Context, level logger.Level) bool {
	return level > logger.LevelDebug
}

// With debug logging off, neither client writes the request's debug line.
func TestSdkGoClientSkipsTheDebugLineLikeClassicClientWhenDebugIsOff(t *testing.T) {
	sdkGoLogger, classicLogger := &quietLogger{}, &quietLogger{}
	// readBothWays reads through sdk-go first.
	next := []*quietLogger{sdkGoLogger, classicLogger}
	newContext := func() context.Context {
		l := next[0]
		next = next[1:]
		return logger.NewContext(context.Background(), l)
	}
	r := readBothWays(t, newContext, nil, respond(http.StatusOK, sdkGoTestVolume))
	require.NoError(t, r.err)
	require.NoError(t, r.classicErr)
	assert.Empty(t, sdkGoLogger.matching("DEBUG GET "+sdkGoTestVolumePath))
	assert.Empty(t, classicLogger.matching("DEBUG GET "+sdkGoTestVolumePath))
}
