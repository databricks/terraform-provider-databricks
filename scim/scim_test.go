package scim

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestConsumerSQLAccessValidation(t *testing.T) {
	tests := []struct {
		name     string
		schema   map[string]*schema.Schema
		identity map[string]any
	}{
		{"group_entitlements", ResourceEntitlements().Schema, map[string]any{"group_id": "abc"}},
		{"user_entitlements", ResourceEntitlements().Schema, map[string]any{"user_id": "abc"}},
		{"service_principal_entitlements", ResourceEntitlements().Schema, map[string]any{"service_principal_id": "abc"}},
		{"group", ResourceGroup().Schema, map[string]any{"display_name": "example"}},
		{"user", ResourceUser().Schema, map[string]any{"user_name": "example@example.com"}},
		{"service_principal", ResourceServicePrincipal().Schema, map[string]any{"display_name": "example"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := tt.identity
			config["databricks_sql_access"] = true
			config["workspace_consume"] = true
			resource := &schema.Resource{Schema: tt.schema}
			require.Empty(t, resource.Validate(terraform.NewResourceConfigRaw(config)))

			data := schema.TestResourceDataRaw(t, tt.schema, config)
			entitlements := readEntitlementsFromData(data)
			require.Equal(t, []ComplexValue{
				{Value: "databricks-sql-access"},
				{Value: "workspace-consume"},
			}, []ComplexValue(entitlements))
			readData := schema.TestResourceDataRaw(t, tt.schema, nil)
			require.NoError(t, entitlements.readIntoData(readData))
			require.True(t, readData.Get("databricks_sql_access").(bool))
			require.True(t, readData.Get("workspace_consume").(bool))

			config["workspace_access"] = true
			require.True(t, resource.Validate(terraform.NewResourceConfigRaw(config)).HasError(),
				"Consumer access must still conflict with Workspace access")
		})
	}
}
