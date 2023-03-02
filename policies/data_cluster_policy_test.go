package policies

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDataSourceClusterPolicy(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/policies/clusters/list",
				Response: ClusterPolicyList{
					Policies: []ClusterPolicy{
						{
							PolicyID:   "abc",
							Name:       "policy",
							Definition: `{"abc":"123"}`,
						},
					},
				},
			},
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceClusterPolicy(),
		ID:          ".",
		State: map[string]any{
			"name": "policy",
		},
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "abc", d.Id())
	assert.Equal(t, `{"abc":"123"}`, d.Get("definition").(string))
}

func TestDataSourceClusterPolicyNotFound(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/policies/clusters/list",
			Status:   404,
			Response: apierr.APIError{
				Message: "searching_error",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/policies/clusters/list",
			Response: ClusterPolicyList{},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		poolsAPI := NewClusterPoliciesAPI(ctx, client)

		_, err := getPolicy(poolsAPI, "searching_error")
		assert.EqualError(t, err, "searching_error")

		_, err = getPolicy(poolsAPI, "unknown")
		assert.EqualError(t, err, "cluster policy 'unknown' wasn't found")
	})
}
