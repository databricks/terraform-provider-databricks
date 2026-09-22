package exporter

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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
