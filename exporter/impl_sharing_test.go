package exporter

import (
	"context"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	sdk_sharing "github.com/databricks/databricks-sdk-go/service/sharing"
)

func TestImportShare(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.1/unity-catalog/shares/test_share?include_shared_data=true",
			Response: sdk_sharing.ShareInfo{
				Name: "test_share",
				Objects: []sdk_sharing.SharedDataObject{
					{
						Name:           "catalog.schema.table1",
						DataObjectType: "TABLE",
					},
					{
						Name:           "catalog.schema.model1",
						DataObjectType: "MODEL",
					},
					{
						Name:           "catalog.schema.volume1",
						DataObjectType: "VOLUME",
					},
					{
						Name:           "catalog.schema1",
						DataObjectType: "SCHEMA",
					},
					{
						Name:           "catalog.schema.notebook1",
						DataObjectType: "NOTEBOOK_FILE",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("uc-shares,uc-grants,uc-volumes,uc-models,uc-tables,uc-schemas")

		// Read the share resource first
		r := &resource{
			Resource: "databricks_share",
			ID:       "test_share",
		}

		// Read the resource to populate its data
		ir := resourcesMap["databricks_share"]
		wrapper := ic.readPluginFrameworkResource(r, ir)
		r.DataWrapper = wrapper

		// Now import it
		err := ir.Import(ic, r)
		assert.NoError(t, err)

		// Verify the correct dependencies are emitted
		// Should emit: grants, table, model, volume, schema (5 total)
		// NOTEBOOK_FILE type should be logged but not emitted as a resource
		require.Equal(t, 5, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_grants[<unknown>] (id: share/test_share)"])
		assert.True(t, ic.testEmits["databricks_sql_table[<unknown>] (id: catalog.schema.table1)"])
		assert.True(t, ic.testEmits["databricks_registered_model[<unknown>] (id: catalog.schema.model1)"])
		assert.True(t, ic.testEmits["databricks_volume[<unknown>] (id: catalog.schema.volume1)"])
		assert.True(t, ic.testEmits["databricks_schema[<unknown>] (id: catalog.schema1)"])
	})
}

func TestListShares(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			Method:       "GET",
			Resource:     "/api/2.1/unity-catalog/shares?",
			Response: sdk_sharing.ListSharesResponse{
				Shares: []sdk_sharing.ShareInfo{
					{
						Name: "test",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("uc-shares")
		ic.currentMetastore = currentMetastoreResponse
		err := resourcesMap["databricks_share"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 1, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_share[<unknown>] (id: test)"])
	})
}
