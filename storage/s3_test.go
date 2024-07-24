package storage

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/clusters"
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
			Response: apierr.NotFound("cluster deleted"),
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

func TestPreprocessS3MountOnDeletedClusterWorks(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=removed-cluster",
			Status:   404,
			Response: apierr.NotFound("cluster deleted"),
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
				ClusterName:  "terraform-mount-s3-access",
				SparkVersion: "11.3.x-scala2.12",
				NumWorkers:   0,
				NodeTypeID:   "i3.xlarge",
				AwsAttributes: &clusters.AwsAttributes{
					Availability:       "SPOT",
					InstanceProfileArn: "arn:aws:iam::1234567:instance-profile/s3-access",
					ZoneID:             "auto",
				},
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
		r := ResourceMount()
		d := r.ToResource().TestResourceData()
		d.MarkNewResource()
		common.StructToData(GenericMount{
			URI:       "s3://bucket",
			ClusterID: "removed-cluster",
			S3: &S3IamMount{
				InstanceProfile: "arn:aws:iam::1234567:instance-profile/s3-access",
			},
		}, r.Schema, d)
		err := preprocessS3MountGeneric(ctx, r.Schema, d, client)
		assert.NoError(t, err)
		assert.Equal(t, "new-cluster", d.Get("cluster_id"))
	})
}
