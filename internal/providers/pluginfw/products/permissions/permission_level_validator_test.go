package permissions

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermissionLevelValidator_Description(t *testing.T) {
	validator := ValidatePermissionLevel()

	desc := validator.Description(context.Background())
	assert.NotEmpty(t, desc)
	assert.Contains(t, desc, "permission level")

	mdDesc := validator.MarkdownDescription(context.Background())
	assert.NotEmpty(t, mdDesc)
	assert.Contains(t, mdDesc, "permission level")
}

// Note: Full validation testing with config is done in acceptance tests
// The validator requires access to the full config to determine object type,
// which is complex to mock in unit tests but straightforward in integration tests
