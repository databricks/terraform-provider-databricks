package scim

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestConsumerAccessCombinedWithOtherAccessEntitlements(t *testing.T) {
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
			combinations := []struct {
				name   string
				fields []string
				values []string
			}{
				{"sql", []string{"databricks_sql_access"}, []string{"databricks-sql-access"}},
				{"workspace", []string{"workspace_access"}, []string{"workspace-access"}},
				{"sql_and_workspace", []string{"databricks_sql_access", "workspace_access"},
					[]string{"databricks-sql-access", "workspace-access"}},
			}
			for _, c := range combinations {
				t.Run(c.name, func(t *testing.T) {
					config := map[string]any{"workspace_consume": true}
					for k, v := range tt.identity {
						config[k] = v
					}
					for _, field := range c.fields {
						config[field] = true
					}
					resource := &schema.Resource{Schema: tt.schema}
					require.Empty(t, resource.Validate(terraform.NewResourceConfigRaw(config)))

					expected := []ComplexValue{}
					for _, value := range c.values {
						expected = append(expected, ComplexValue{Value: value})
					}
					expected = append(expected, ComplexValue{Value: "workspace-consume"})
					data := schema.TestResourceDataRaw(t, tt.schema, config)
					entitlements := readEntitlementsFromData(data)
					require.Equal(t, expected, []ComplexValue(entitlements))

					readData := schema.TestResourceDataRaw(t, tt.schema, nil)
					require.NoError(t, entitlements.readIntoData(readData))
					for _, field := range append(c.fields, "workspace_consume") {
						require.True(t, readData.Get(field).(bool), field)
					}
				})
			}
		})
	}
}
