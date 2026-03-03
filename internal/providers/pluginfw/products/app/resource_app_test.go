package app

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAppResourceSchema_DescriptionIsOptionalAndComputed(t *testing.T) {
	r := ResourceApp()
	resp := &resource.SchemaResponse{}
	r.Schema(context.Background(), resource.SchemaRequest{}, resp)
	require.False(t, resp.Diagnostics.HasError(), "schema should not have errors")

	descAttr, ok := resp.Schema.Attributes["description"]
	require.True(t, ok, "description attribute must exist in schema")

	strAttr, ok := descAttr.(schema.StringAttribute)
	require.True(t, ok, "description attribute must be a StringAttribute")

	assert.True(t, strAttr.Optional, "description must be Optional")
	assert.True(t, strAttr.Computed, "description must be Computed")
	assert.NotEmpty(t, strAttr.PlanModifiers, "description must have plan modifiers (UseStateForUnknown)")
}
