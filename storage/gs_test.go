package storage

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrValidateClusterForGoogleStorage_Failures(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			MatchAny:     true,
			ReuseRequest: true,
			Status:       404,
			Response: &apierr.APIError{
				ErrorCode:  "NOT_FOUND",
				StatusCode: 404,
				Message:    "nope",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		d := ResourceMount().ToResource().TestResourceData()
		err := createOrValidateClusterForGoogleStorage(ctx, client, d, "a", "")
		assert.EqualError(t, err, "cannot re-create mounting cluster: cannot determine smallest node type: nope")

		err = createOrValidateClusterForGoogleStorage(ctx, client, d, "", "b")
		assert.EqualError(t, err, "cannot create mounting cluster: cannot determine smallest node type: nope")
	})
}

func TestCreateOrValidateClusterForGoogleStorage_FailsOnErrorGettingCluster(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=my-cluster",
			Status:   500,
			Response: apierr.APIError{
				ErrorCode:  "SERVER_ERROR",
				StatusCode: 500,
				Message:    "Server error",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		d := ResourceMount().ToResource().TestResourceData()
		err := createOrValidateClusterForGoogleStorage(ctx, client, d, "my-cluster", "service-account")
		assert.EqualError(t, err, "cannot get mounting cluster: Server error")
	})
}
