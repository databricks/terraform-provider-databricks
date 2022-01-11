package storage

import (
	"context"
	"strings"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/clusters"
	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/internal"
	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testGcsBucketPath = "gs://" + testS3BucketName

func TestGSMountDefaults(t *testing.T) {
	s := ResourceDatabricksMountSchema()
	d := schema.TestResourceDataRaw(t, s, map[string]interface{}{})
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{})
	defer server.Close()
	require.NoError(t, err, err)

	err = GSMount{}.ValidateAndApplyDefaults(d, client)
	qa.AssertErrorStartsWith(t, err, "'name' is not detected & it's impossible to infer it")

	d = schema.TestResourceDataRaw(t, s, map[string]interface{}{"name": "test"})
	err = GSMount{}.ValidateAndApplyDefaults(d, client)
	require.NoError(t, err, err)
	assert.Equal(t, d.Get("name").(string), "test")
	d = schema.TestResourceDataRaw(t, s, map[string]interface{}{})
	err = GSMount{BucketName: "abc"}.ValidateAndApplyDefaults(d, client)
	require.NoError(t, err, err)
	assert.Equal(t, d.Get("name").(string), "abc")
}

func TestResourceGcsMountGenericCreate_WithCluster(t *testing.T) {
	google_account := "acc@acc-dbx.iam.gserviceaccount.com"
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=this_cluster",
				Response: clusters.ClusterInfo{
					State: clusters.ClusterStateRunning,
					GcpAttributes: &clusters.GcpAttributes{
						GoogleServiceAccount: google_account,
					},
				},
			},
		},
		Resource: ResourceMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := internal.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			if strings.HasPrefix(trunc, "def safe_mount") {
				assert.Contains(t, trunc, testGcsBucketPath) // bucketname
				assert.Contains(t, trunc, `{}`)              // empty brackets for empty config
			}
			assert.Contains(t, trunc, "/mnt/this_mount")
			return mockMountInfo(testGcsBucketPath, "a1b2c3")
		},
		State: map[string]interface{}{
			"cluster_id": "this_cluster",
			"name":       "this_mount",
			"gs": []interface{}{map[string]interface{}{
				"bucket_name": testS3BucketName,
			}},
		},
		Create: true,
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, "this_mount", d.Id())
	assert.Equal(t, testGcsBucketPath, d.Get("source"))
}

func TestResourceGcsMountGenericCreate_WithCluster_NoName(t *testing.T) {
	google_account := "acc@acc-dbx.iam.gserviceaccount.com"
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=this_cluster",
				Response: clusters.ClusterInfo{
					State: clusters.ClusterStateRunning,
					GcpAttributes: &clusters.GcpAttributes{
						GoogleServiceAccount: google_account,
					},
				},
			},
		},
		Resource: ResourceMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := internal.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			if strings.HasPrefix(trunc, "def safe_mount") {
				assert.Contains(t, trunc, testGcsBucketPath) // bucketname
				assert.Contains(t, trunc, `{}`)              // empty brackets for empty config
			}
			assert.Contains(t, trunc, "/mnt/"+testS3BucketName)
			return mockMountInfo(testGcsBucketPath, "a1b2c3")
		},
		State: map[string]interface{}{
			"cluster_id": "this_cluster",
			"gs": []interface{}{map[string]interface{}{
				"bucket_name": testS3BucketName,
			}},
		},
		Create: true,
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, testS3BucketName, d.Id())
	assert.Equal(t, testGcsBucketPath, d.Get("source"))
}

func TestResourceGcsMountGenericCreate_WithServiceAccount(t *testing.T) {
	googleAccount := "acc@acc-dbx.iam.gserviceaccount.com"
	clusterName := "terraform-mount-gcs-bcb24f32098efa4172f435adbed2dae2"
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=abcd",
				Response: clusters.ClusterInfo{
					ClusterID: "abcd",
					State:     clusters.ClusterStateRunning,
					GcpAttributes: &clusters.GcpAttributes{
						GoogleServiceAccount: googleAccount,
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/list",
				Response: map[string]interface{}{},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/list",
				Response: clusters.ClusterList{
					Clusters: []clusters.ClusterInfo{
						{
							ClusterID: "abcd",
							State:     clusters.ClusterStateRunning,
							GcpAttributes: &clusters.GcpAttributes{
								GoogleServiceAccount: googleAccount,
							},
							AutoterminationMinutes: 10,
							SparkConf: map[string]string{"spark.databricks.cluster.profile": "singleNode",
								"spark.master": "local[*]", "spark.scheduler.mode": "FIFO"},
							CustomTags:   map[string]string{"ResourceClass": "SingleNode"},
							ClusterName:  clusterName,
							SparkVersion: "7.3.x-scala2.12",
							NumWorkers:   0,
						},
					},
				},
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/clusters/spark-versions",
				Response:     sparkVersionsResponse,
				ReuseRequest: true,
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/clusters/list-node-types",
				ReuseRequest: true,
				Response:     nodeListResponse,
			},
			{
				Method:       "POST",
				Resource:     "/api/2.0/clusters/create",
				ReuseRequest: true,
				ExpectedRequest: clusters.Cluster{
					NodeTypeID: "Standard_F4s",
					GcpAttributes: &clusters.GcpAttributes{
						GoogleServiceAccount: "acc@acc-dbx.iam.gserviceaccount.com",
					},
					AutoterminationMinutes: 10,
					SparkConf: map[string]string{"spark.databricks.cluster.profile": "singleNode",
						"spark.master": "local[*]", "spark.scheduler.mode": "FIFO"},
					CustomTags:   map[string]string{"ResourceClass": "SingleNode"},
					ClusterName:  clusterName,
					SparkVersion: "7.3.x-scala2.12",
					NumWorkers:   0,
				},
				Response: clusters.ClusterID{
					ClusterID: "abcd",
				},
			},
		},
		Resource: ResourceMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := internal.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			if strings.HasPrefix(trunc, "def safe_mount") {
				assert.Contains(t, trunc, testGcsBucketPath) // bucketname
				assert.Contains(t, trunc, `{}`)              // empty brackets for empty config
			}
			assert.Contains(t, trunc, "/mnt/this_mount")
			return mockMountInfo(testGcsBucketPath, "a1b2c3")
		},
		State: map[string]interface{}{
			"name": "this_mount",
			"gs": []interface{}{map[string]interface{}{
				"bucket_name":     testS3BucketName,
				"service_account": googleAccount,
			}},
		},
		Create: true,
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, "this_mount", d.Id())
	assert.Equal(t, "abcd", d.Get("cluster_id"))
	assert.Equal(t, testGcsBucketPath, d.Get("source"))
}

func TestResourceGcsMountGenericCreate_nothing_specified(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource: ResourceMount(),
		State: map[string]interface{}{
			"name": "this_mount",
			"gs": []interface{}{map[string]interface{}{
				"bucket_name": testS3BucketName,
			}},
		},
		Create: true,
	}.Apply(t)
	require.EqualError(t, err, "either cluster_id or service_account must be specified to mount GCS bucket")
}

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
		assert.EqualError(t, err, "cannot get mounting cluster: nope")

		err = createOrValidateClusterForGoogleStorage(ctx, client, d, "", "b")
		assert.EqualError(t, err, "cannot create mounting cluster: nope")
	})
}
