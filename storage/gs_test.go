package storage

import (
	"context"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/clusters"
	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrValidateClusterForGoogleStorage_Failures(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			MatchAny:     true,
			ReuseRequest: true,
			Status:       404,
			Response:     common.NotFound("nope"),
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		d := ResourceMount().TestResourceData()
		err := createOrValidateClusterForGoogleStorage(ctx, client, d, "a", "")
		assert.EqualError(t, err, "cannot re-create mounting cluster: nope")

		err = createOrValidateClusterForGoogleStorage(ctx, client, d, "", "b")
		assert.EqualError(t, err, "cannot create mounting cluster: nope")
	})
}

func TestCreateOrValidateClusterForGoogleStorage_WorksOnDeletedCluster(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=removed-cluster",
			Status:   404,
			Response: common.NotFound("cluster deleted"),
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/list",
			Response: clusters.ClusterList{
				Clusters: []clusters.ClusterInfo{},
			},
		},
		{
			ReuseRequest: true,
			Method:       "GET",
			Resource:     "/api/2.0/clusters/spark-versions",
		},
		{
			ReuseRequest: true,
			Method:       "GET",
			Resource:     "/api/2.0/clusters/list-node-types",
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/clusters/create",
			ExpectedRequest: clusters.Cluster{
				CustomTags: map[string]string{
					"ResourceClass": "SingleNode",
				},
				ClusterName: "terraform-mount-gcs-03a56ec1d1576b505aabf088337cbf36",
				GcpAttributes: &clusters.GcpAttributes{
					GoogleServiceAccount: "service-account",
				},
				SparkVersion:           "7.3.x-scala2.12",
				NumWorkers:             0,
				NodeTypeID:             "i3.xlarge",
				AutoterminationMinutes: 10,
				SparkConf: map[string]string{
					"spark.databricks.cluster.profile": "singleNode",
					"spark.master":                     "local[*]",
					"spark.scheduler.mode":             "FIFO",
				},
			},
			Response: clusters.ClusterID{
				ClusterID: "new-cluster",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=new-cluster",
			Response: clusters.ClusterInfo{
				ClusterID:    "new-cluster",
				State:        "RUNNING",
				StateMessage: "created",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		d := ResourceMount().TestResourceData()
		err := createOrValidateClusterForGoogleStorage(ctx, client, d, "removed-cluster", "service-account")
		assert.NoError(t, err)
		assert.Equal(t, "new-cluster", d.Get("cluster_id"))
	})
}
