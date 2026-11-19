package common

import (
	"context"
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/assert"
)

type ownerRollbackContextKey struct{}

func TestOwnerRollbackContextOutlivesCancelledParentUntilTimeout(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		parent, cancelParent := context.WithTimeout(
			context.WithValue(t.Context(), ownerRollbackContextKey{}, "apply"), time.Second)
		cancelParent()

		ctx, cancel := OwnerRollbackContext(parent)
		defer cancel()
		assert.Equal(t, "apply", ctx.Value(ownerRollbackContextKey{}))
		assert.NoError(t, ctx.Err())
		deadline, ok := ctx.Deadline()
		assert.True(t, ok)
		assert.Equal(t, time.Now().Add(OwnerRollbackTimeout), deadline)

		start := time.Now()
		<-ctx.Done()
		assert.Equal(t, OwnerRollbackTimeout, time.Since(start))
		assert.ErrorIs(t, ctx.Err(), context.DeadlineExceeded)
	})
}
