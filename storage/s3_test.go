package storage

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestPreprocessS3MountOnDeletedClusterNoInstanceProfileSpecifiedError(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=removed-cluster",
			Status:   404,
			Response: &apierr.APIError{
				ErrorCode:  "NOT_FOUND",
				StatusCode: 404,
				Message:    "cluster deleted",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		r := ResourceMount()
		d := r.ToResource().TestResourceData()
		d.Set("uri", "s3://bucket")
		d.Set("cluster_id", "removed-cluster")
		err := preprocessS3MountGeneric(ctx, r.Schema, d, client)
		assert.EqualError(t, err, "instance profile is required to re-create mounting cluster")
	})
}
