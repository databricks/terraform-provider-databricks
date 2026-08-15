package mws_private_access_settings

import (
	"context"
	"reflect"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSchemaPrivateAccessSettingsIdIsRequired(t *testing.T) {
	ctx := context.Background()
	attrs, _ := tfschema.DataSourceStructToSchemaMap(ctx, MwsPrivateAccessSettingsData{}, nil)

	idAttr, ok := attrs["private_access_settings_id"]
	require.True(t, ok, "private_access_settings_id attribute must exist in schema")
	strAttr, ok := idAttr.(schema.StringAttribute)
	require.True(t, ok, "private_access_settings_id must be a StringAttribute")
	assert.True(t, strAttr.Required, "private_access_settings_id must be required")
	assert.False(t, strAttr.Computed, "private_access_settings_id must not be computed")
}

func TestSchemaComputedFields(t *testing.T) {
	ctx := context.Background()
	attrs, _ := tfschema.DataSourceStructToSchemaMap(ctx, MwsPrivateAccessSettingsData{}, nil)

	computedFields := []string{
		"account_id",
		"private_access_level",
		"private_access_settings_name",
		"public_access_enabled",
		"region",
	}
	for _, field := range computedFields {
		attr, ok := attrs[field]
		require.True(t, ok, "attribute %s must exist in schema", field)
		switch a := attr.(type) {
		case schema.StringAttribute:
			assert.True(t, a.Computed, "%s must be computed", field)
			assert.False(t, a.Required, "%s must not be required", field)
		case schema.BoolAttribute:
			assert.True(t, a.Computed, "%s must be computed", field)
			assert.False(t, a.Required, "%s must not be required", field)
		default:
			t.Errorf("unexpected attribute type for %s: %T", field, attr)
		}
	}
}

func TestGetComplexFieldTypes(t *testing.T) {
	ctx := context.Background()
	m := MwsPrivateAccessSettingsData{}
	complexTypes := m.GetComplexFieldTypes(ctx)

	assert.Len(t, complexTypes, 1)
	assert.Contains(t, complexTypes, "allowed_vpc_endpoint_ids")
	assert.Equal(t, reflect.TypeOf(types.String{}), complexTypes["allowed_vpc_endpoint_ids"])
}
