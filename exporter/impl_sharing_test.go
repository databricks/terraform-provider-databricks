package exporter

import (
	"context"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	sdk_sharing "github.com/databricks/databricks-sdk-go/service/sharing"
	tf_sharing "github.com/databricks/terraform-provider-databricks/sharing"
)

func TestImportShare(t *testing.T) {
	ic := importContextForTest()
	ic.enableServices("uc-grants,uc-volumes,uc-models,uc-tables")
	d := tf_sharing.ResourceShare().ToResource().TestResourceData()
	scm := tf_sharing.ResourceShare().Schema
	share := tf_sharing.ShareInfo{
		ShareInfo: sdk_sharing.ShareInfo{
			Name: "stest",
			Objects: []sdk_sharing.SharedDataObject{
				{
					DataObjectType: "TABLE",
					Name:           "ctest.stest.table1",
				},
				{
					DataObjectType: "MODEL",
					Name:           "ctest.stest.model1",
				},
				{
					DataObjectType: "VOLUME",
					Name:           "ctest.stest.vol1",
				},
				{
					DataObjectType: "NOTEBOOK",
					Name:           "Test",
				},
			},
		},
	}
	d.MarkNewResource()
	err := common.StructToData(share, scm, d)
	require.NoError(t, err)
	err = resourcesMap["databricks_share"].Import(ic, &resource{
		ID:   "stest",
		Data: d,
	})
	assert.NoError(t, err)
	require.Equal(t, 4, len(ic.testEmits))
	assert.True(t, ic.testEmits["databricks_grants[<unknown>] (id: share/stest)"])
	assert.True(t, ic.testEmits["databricks_registered_model[<unknown>] (id: ctest.stest.model1)"])
	assert.True(t, ic.testEmits["databricks_volume[<unknown>] (id: ctest.stest.vol1)"])
	assert.True(t, ic.testEmits["databricks_sql_table[<unknown>] (id: ctest.stest.table1)"])
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
