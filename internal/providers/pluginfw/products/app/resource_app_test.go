package app

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAppResourceSchema_SpaceInheritedFieldsAreComputed(t *testing.T) {
	ctx := context.Background()
	r := ResourceApp()
	resp := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, resp)

	for _, field := range []string{"resources", "user_api_scopes", "budget_policy_id"} {
		attr, ok := resp.Schema.Attributes[field]
		require.True(t, ok, "field %s should exist in schema", field)
		assert.True(t, attr.IsOptional(), "field %s should be optional", field)
		assert.True(t, attr.IsComputed(), "field %s should be computed (server-populated from space)", field)
	}
}
