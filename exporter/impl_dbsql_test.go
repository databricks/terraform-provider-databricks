package exporter

import (
	"context"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestExportSqlEndpointSettings(t *testing.T) {
	for _, tc := range []struct{ warehouseType, autoStop, timeout, maxDbu string }{
		{"PRO", "30", "0", "0"}, {"PRO", "0", "0", "0"},
		{"CLASSIC", "30", "0", "0"}, {"CLASSIC", "0", "0", "0"},
		{"REALTIME", "30", "60", "96"}, {"REALTIME", "0", "0", "0"},
	} {
		t.Run(tc.warehouseType+"/autostop="+tc.autoStop, func(t *testing.T) {
			ic := importContextForTest()
			r := &resource{Resource: "databricks_sql_endpoint", ID: "abc", Name: "warehouse"}
			r.Data = ic.Resources[r.Resource].Data(&terraform.InstanceState{
				ID: r.ID,
				Attributes: map[string]string{
					"name": "warehouse", "cluster_size": "Small", "warehouse_type": tc.warehouseType,
					"auto_stop_mins": tc.autoStop, "statement_timeout": tc.timeout, "max_dbu": tc.maxDbu,
					"customer_version_label": "2026.10",
				},
			})
			f := hclwrite.NewEmptyFile()
			body := f.Body().AppendNewBlock("resource", []string{r.Resource, r.Name}).Body()
			require.NoError(t, ic.unifiedDataToHcl(ic.Importables[r.Resource], nil, r, body))
			expected := map[string]string{"auto_stop_mins": tc.autoStop}
			if tc.warehouseType == "REALTIME" {
				expected["statement_timeout"] = tc.timeout
				if tc.maxDbu != "0" {
					expected["max_dbu"] = tc.maxDbu
				}
			}
			for _, field := range []string{"auto_stop_mins", "statement_timeout", "max_dbu", "customer_version_label"} {
				attr := body.GetAttribute(field)
				if value, ok := expected[field]; ok {
					require.NotNil(t, attr, field)
					assert.Equal(t, value, strings.TrimSpace(string(attr.Expr().BuildTokens(nil).Bytes())), field)
				} else {
					assert.Nil(t, attr, field)
				}
			}
		})
	}
}

func TestDefaultWarehouseOverrideId(t *testing.T) {
	assert.Equal(t, "123", defaultWarehouseOverrideId("default-warehouse-overrides/123"))
	assert.Equal(t, "me", defaultWarehouseOverrideId("default-warehouse-overrides/me"))
	assert.Equal(t, "", defaultWarehouseOverrideId("default-warehouse-overrides/"))
	assert.Equal(t, "", defaultWarehouseOverrideId("123"))
	assert.Equal(t, "", defaultWarehouseOverrideId("warehouses/123"))
}

func TestListWarehouseDefaultOverrides(t *testing.T) {
	qa.MockWorkspaceApply(t, func(mw *mocks.MockWorkspaceClient) {
		mw.GetMockWarehousesAPI().EXPECT().
			ListDefaultWarehouseOverrides(mock.Anything, sql.ListDefaultWarehouseOverridesRequest{}).
			Return(createIteratorFromSlice([]sql.DefaultWarehouseOverride{
				{
					DefaultWarehouseOverrideId: "123",
					Name:                       "default-warehouse-overrides/123",
					Type:                       sql.DefaultWarehouseOverrideTypeCustom,
					WarehouseId:                "warehouse-abc",
				},
				{
					DefaultWarehouseOverrideId: "456",
					Name:                       "default-warehouse-overrides/456",
					Type:                       sql.DefaultWarehouseOverrideTypeLastSelected,
				},
			}))
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("sql-endpoints")

		err := listWarehouseDefaultOverrides(ic)
		assert.NoError(t, err)
		assert.True(t, ic.testEmits["databricks_warehouses_default_warehouse_override[<unknown>] (id: default-warehouse-overrides/123)"])
		assert.True(t, ic.testEmits["databricks_warehouses_default_warehouse_override[<unknown>] (id: default-warehouse-overrides/456)"])
	})
}

func TestImportWarehouseDefaultOverride(t *testing.T) {
	ic := importContextForTest()
	ic.enableServices("sql-endpoints,users")

	// The override ID component is the user ID, which is emitted as a dependency.
	err := importWarehouseDefaultOverride(ic, &resource{
		Resource: "databricks_warehouses_default_warehouse_override",
		ID:       "default-warehouse-overrides/123",
	})
	assert.NoError(t, err)
	assert.True(t, ic.testEmits["databricks_user[<unknown>] (id: 123)"],
		"Should emit the user the override belongs to")

	// The special "me" identifier doesn't map to a specific user, so nothing is emitted.
	ic = importContextForTest()
	ic.enableServices("sql-endpoints,users")
	err = importWarehouseDefaultOverride(ic, &resource{
		Resource: "databricks_warehouses_default_warehouse_override",
		ID:       "default-warehouse-overrides/me",
	})
	assert.NoError(t, err)
	assert.Empty(t, ic.testEmits, "Should not emit a user for the 'me' identifier")
}
