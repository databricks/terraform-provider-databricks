package common

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

func TestStringIsUUID(t *testing.T) {
	assert.True(t, StringIsUUID("3f670caf-9a4b-4479-8143-1a0878da8f57"))
	assert.False(t, StringIsUUID("abc"))
}

func TestGetTerraformVersionFromContext(t *testing.T) {
	assert.Equal(t, "unknown", GetTerraformVersionFromContext(context.Background()))

	//
	p := &schema.Provider{}
	p.TerraformVersion = "exporter"
	ctx := context.WithValue(context.Background(), Provider, p)
	assert.Equal(t, "exporter", GetTerraformVersionFromContext(ctx))

	//
	assert.True(t, IsExporter(ctx))
}

func TestSuppressDiffWhitespaceChange(t *testing.T) {
	assert.True(t, SuppressDiffWhitespaceChange("k", "value", "  value  ", nil))
	assert.False(t, SuppressDiffWhitespaceChange("k", "value", "new_value", nil))
}

func TestSuppressCaseSensitivity(t *testing.T) {
	assert.True(t, SuppressCaseSensitivity("k", "value", "VALUE", nil))
	assert.False(t, SuppressCaseSensitivity("k", "value", "new_value", nil))
}
