package mws_private_access_settings

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/provisioning"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSchemaLookupFieldsAreOptionalComputed(t *testing.T) {
	ctx := context.Background()
	attrs, _ := tfschema.DataSourceStructToSchemaMap(ctx, MwsPrivateAccessSettingsData{}, nil)

	lookupFields := []string{"private_access_settings_id", "private_access_settings_name"}
	for _, field := range lookupFields {
		attr, ok := attrs[field]
		require.True(t, ok, "attribute %s must exist in schema", field)
		strAttr, ok := attr.(schema.StringAttribute)
		require.True(t, ok, "%s must be a StringAttribute", field)
		assert.True(t, strAttr.Optional, "%s must be optional", field)
		assert.True(t, strAttr.Computed, "%s must be computed", field)
		assert.False(t, strAttr.Required, "%s must not be required", field)
	}
}

func TestSchemaComputedFields(t *testing.T) {
	ctx := context.Background()
	attrs, _ := tfschema.DataSourceStructToSchemaMap(ctx, MwsPrivateAccessSettingsData{}, nil)

	computedFields := []string{
		"account_id",
		"allowed_vpc_endpoint_ids",
		"private_access_level",
		"public_access_enabled",
		"region",
	}
	for _, field := range computedFields {
		attr, ok := attrs[field]
		require.True(t, ok, "attribute %s must exist in schema", field)
		switch a := attr.(type) {
		case schema.StringAttribute:
			assert.True(t, a.Computed, "%s must be computed", field)
			assert.False(t, a.Optional, "%s must not be optional", field)
		case schema.BoolAttribute:
			assert.True(t, a.Computed, "%s must be computed", field)
			assert.False(t, a.Optional, "%s must not be optional", field)
		case schema.ListAttribute:
			assert.True(t, a.Computed, "%s must be computed", field)
			assert.False(t, a.Optional, "%s must not be optional", field)
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

func TestNoPrivateAccessSettingsError(t *testing.T) {
	ctx := context.Background()
	name := "nonexistent"
	matches := []provisioning.PrivateAccessSettings{}
	dd := validatePrivateAccessSettingsList(ctx, matches, name)
	expected := diag.Diagnostics{diag.NewErrorDiagnostic(fmt.Sprintf("there is no private access settings with name '%s'", name), "")}
	assert.True(t, dd.HasError())
	assert.Equal(t, expected, dd)
}

func TestMultiplePrivateAccessSettingsError(t *testing.T) {
	ctx := context.Background()
	name := "duplicate"
	settings := []provisioning.PrivateAccessSettings{
		{PrivateAccessSettingsName: "duplicate", PrivateAccessSettingsId: "id-1"},
		{PrivateAccessSettingsName: "duplicate", PrivateAccessSettingsId: "id-2"},
	}
	dd := validatePrivateAccessSettingsList(ctx, settings, name)
	assert.True(t, dd.HasError())
	assert.Contains(t, dd[0].Summary(), "more than one")
}

func TestSinglePrivateAccessSettingsSuccess(t *testing.T) {
	ctx := context.Background()
	name := "my-pas"
	settings := []provisioning.PrivateAccessSettings{
		{PrivateAccessSettingsName: "my-pas", PrivateAccessSettingsId: "id-1"},
	}
	dd := validatePrivateAccessSettingsList(ctx, settings, name)
	assert.False(t, dd.HasError())
}
