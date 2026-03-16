package common

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContextWithApiLevel(t *testing.T) {
	ctx := context.Background()
	assert.Equal(t, "", ApiLevelFromContext(ctx))

	ctx = ContextWithApiLevel(ctx, ApiLevelAccount)
	assert.Equal(t, ApiLevelAccount, ApiLevelFromContext(ctx))

	ctx = ContextWithApiLevel(ctx, ApiLevelWorkspace)
	assert.Equal(t, ApiLevelWorkspace, ApiLevelFromContext(ctx))
}

func TestApiLevelFromContextEmpty(t *testing.T) {
	ctx := context.Background()
	assert.Equal(t, "", ApiLevelFromContext(ctx))
}

func TestValidateApiField(t *testing.T) {
	// Validation is done by the schema ValidateFunc, so we just test the constants
	assert.Equal(t, "account", ApiLevelAccount)
	assert.Equal(t, "workspace", ApiLevelWorkspace)
}
