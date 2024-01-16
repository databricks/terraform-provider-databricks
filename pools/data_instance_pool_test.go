package pools

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDataSourceInstancePool(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/instance-pools/list",
				Response: InstancePoolList{
					InstancePools: []InstancePoolAndStats{
						{
							InstancePoolID:   "abc",
							InstancePoolName: "pool",
							NodeTypeID:       "node-type",
						},
					},
				},
			},
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceInstancePool(),
		ID:          ".",
		State: map[string]any{
			"name": "pool",
		},
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
	assert.NotNil(t, d.Get("pool_info"))
	assert.Equal(t, "node-type", d.Get("pool_info.0.node_type_id").(string))
}

func TestDataSourceInstancePoolsGetPool(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/instance-pools/list",
			Status:   404,
			Response: apierr.APIError{
				Message: "searching_error",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/instance-pools/list",
			Response: InstancePoolList{},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		poolsAPI := NewInstancePoolsAPI(ctx, client)

		_, err := getPool(poolsAPI, "searching_error")
		assert.EqualError(t, err, "searching_error")

		_, err = getPool(poolsAPI, "unknown")
		assert.EqualError(t, err, "instance pool 'unknown' doesn't exist")
	})
}

func TestDataSourceInstancePool_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/instance-pools/list",
				Response: InstancePoolList{
					InstancePools: []InstancePoolAndStats{},
				},
			},
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceInstancePool(),
		ID:          ".",
		State: map[string]any{
			"name": "Unknown",
		},
	}.ExpectError(t, "instance pool 'Unknown' doesn't exist")
}
