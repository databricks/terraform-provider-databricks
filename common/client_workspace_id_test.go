package common

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setCachedWorkspaceID checks the cache again under its lock, so an operation that
// waited there while another one resolved the workspace ID sends no /Me of its own.
func TestSetCachedWorkspaceIDSkipsMeWhenAnotherOperationSetIt(t *testing.T) {
	server, calls := serveSequence(t)
	c := newSdkGoTestClient(t, server, nil)
	w, err := c.WorkspaceClient()
	require.NoError(t, err)
	// Another operation resolved the ID while this one waited for the lock.
	c.SetCachedWorkspaceID(67890)

	require.NoError(t, c.setCachedWorkspaceID(context.Background(), w))

	assert.Equal(t, int64(67890), c.cachedWorkspaceID)
	assert.EqualValues(t, 0, calls.Load(), "no /Me call")
}
