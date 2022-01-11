package storage

import (
	"fmt"
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

// ============================== AWS S3 Tests ==============================

// Test interface compliance via compile time error
var _ Mount = (*S3IamMount)(nil)

var sparkVersionsResponse = clusters.SparkVersionsList{
	SparkVersions: []clusters.SparkVersion{
		{
			Version:     "7.3.x-scala2.12",
			Description: "7.3 LTS (includes Apache Spark 3.0.1, Scala 2.12)",
		},
	},
}

var nodeListResponse = clusters.NodeTypeList{
	NodeTypes: []clusters.NodeType{
		{
			NodeTypeID:     "Standard_F4s",
			InstanceTypeID: "Standard_F4s",
			MemoryMB:       8192,
			NumCores:       4,
			NodeInstanceType: &clusters.NodeInstanceType{
				LocalDisks:      1,
				LocalDiskSizeGB: 16,
				LocalNVMeDisks:  0,
			},
		},
	},
}

func TestS3MountDefaults(t *testing.T) {
	s := ResourceDatabricksMountSchema()
	d := schema.TestResourceDataRaw(t, s, map[string]interface{}{})
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{})
	defer server.Close()
	require.NoError(t, err, err)

	err = S3IamMount{}.ValidateAndApplyDefaults(d, client)
	qa.AssertErrorStartsWith(t, err, "'name' is not detected & it's impossible to infer it")

	d = schema.TestResourceDataRaw(t, s, map[string]interface{}{"name": "test"})
	err = S3IamMount{}.ValidateAndApplyDefaults(d, client)
	require.NoError(t, err, err)
	assert.Equal(t, d.Get("name").(string), "test")
	d = schema.TestResourceDataRaw(t, s, map[string]interface{}{})
	err = S3IamMount{BucketName: "abc"}.ValidateAndApplyDefaults(d, client)
	require.NoError(t, err, err)
	assert.Equal(t, d.Get("name").(string), "abc")
}

func mockMountInfo(source, hash string) common.CommandResults {
	return common.CommandResults{
		ResultType: "text",
		Data:       fmt.Sprintf(
			`{"source":"%s","config_hash":"%s"}`, 
			source, hash),
	}
}

func TestResourceAwsS3MountGenericCreate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=this_cluster",
				Response: clusters.ClusterInfo{
					State: clusters.ClusterStateRunning,
					AwsAttributes: &clusters.AwsAttributes{
						InstanceProfileArn: "abc",
					},
				},
			},
		},
		Resource: ResourceMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := internal.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			if strings.HasPrefix(trunc, "def safe_mount") {
				assert.Contains(t, trunc, testS3BucketPath) // bucketname
				assert.Contains(t, trunc, `{}`)             // empty brackets for empty config
			}
			assert.Contains(t, trunc, "/mnt/this_mount")
			return mockMountInfo(testS3BucketPath, "a1b2c3")
		},
		State: map[string]interface{}{
			"cluster_id": "this_cluster",
			"name":       "this_mount",
			"s3": []interface{}{map[string]interface{}{
				"bucket_name": testS3BucketName,
			}},
		},
		Create: true,
	}.ApplyNoError(t)
}

func TestResourceAwsS3MountGenericCreate_NoName(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=this_cluster",
				Response: clusters.ClusterInfo{
					State: clusters.ClusterStateRunning,
					AwsAttributes: &clusters.AwsAttributes{
						InstanceProfileArn: "abc",
					},
				},
			},
		},
		Resource: ResourceMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := internal.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			if strings.HasPrefix(trunc, "def safe_mount") {
				assert.Contains(t, trunc, testS3BucketPath) // bucketname
				assert.Contains(t, trunc, `{}`)             // empty brackets for empty config
			}
			assert.Contains(t, trunc, "/mnt/"+testS3BucketName)
			return mockMountInfo(testS3BucketPath, "a1b2c3")
		},
		State: map[string]interface{}{
			"cluster_id": "this_cluster",
			"s3": []interface{}{map[string]interface{}{
				"bucket_name": testS3BucketName,
			}},
		},
		Create: true,
	}.ApplyNoError(t)
}

func TestResourceAwsS3MountGenericCreate_WithInstanceProfile(t *testing.T) {
	instance_profile := "arn:aws:iam::1234567:instance-profile/s3-access"
	clusterName := "terraform-mount-s3-access"
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=abcd",
				Response: clusters.ClusterInfo{
					ClusterID: "abcd",
					State:     clusters.ClusterStateRunning,
					AwsAttributes: &clusters.AwsAttributes{
						InstanceProfileArn: instance_profile,
						Availability:       "SPOT",
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
							AwsAttributes: &clusters.AwsAttributes{
								InstanceProfileArn: instance_profile,
								Availability:       "SPOT",
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
					AwsAttributes: &clusters.AwsAttributes{
						InstanceProfileArn: instance_profile,
						Availability:       "SPOT",
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
				assert.Contains(t, trunc, testS3BucketPath) // bucketname
				assert.Contains(t, trunc, `{}`)             // empty brackets for empty config
			}
			assert.Contains(t, trunc, "/mnt/this_mount")
			return mockMountInfo(testS3BucketPath, "a1b2c3")
		},
		State: map[string]interface{}{
			"name": "this_mount",
			"s3": []interface{}{map[string]interface{}{
				"bucket_name":      testS3BucketName,
				"instance_profile": instance_profile,
			}},
		},
		Create: true,
	}.ApplyNoError(t)
}

func TestResourceAwsS3MountGenericCreate_nothing_specified(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceMount(),
		State: map[string]interface{}{
			"name": "this_mount",
			"s3": []interface{}{map[string]interface{}{
				"bucket_name": testS3BucketName,
			}},
		},
		Create: true,
	}.ExpectError(t, "either cluster_id or instance_profile must be specified to mount S3 bucket")
}

func TestResourceAwsS3MountGenericCreate_invalid_arn(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceMount(),
		State: map[string]interface{}{
			"name": "this_mount",
			"s3": []interface{}{map[string]interface{}{
				"bucket_name":      testS3BucketName,
				"instance_profile": "this_mount",
			}},
		},
		Create: true,
	}.ExpectError(t, "invalid arn: this_mount")
}

func TestResourceAwsS3MountGenericRead(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=this_cluster",
				Response: clusters.ClusterInfo{
					State: clusters.ClusterStateRunning,
					AwsAttributes: &clusters.AwsAttributes{
						InstanceProfileArn: "abc",
					},
				},
			},
		},
		Resource: ResourceMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := internal.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			assert.Contains(t, trunc, "dbutils.fs.mounts()")
			assert.Contains(t, trunc, `mount.mountPoint == "/mnt/this_mount"`)
			return mockMountInfo(testS3BucketPath, "a1b2c3")
		},
		State: map[string]interface{}{
			"cluster_id": "this_cluster",
			"name":       "this_mount",
			"s3": []interface{}{map[string]interface{}{
				"bucket_name": testS3BucketName,
			}},
		},
		ID:   "this_mount",
		Read: true,
	}.ApplyNoError(t)
}

func TestResourceAwsS3MountGenericRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=this_cluster",
				Response: clusters.ClusterInfo{
					State: clusters.ClusterStateRunning,
					AwsAttributes: &clusters.AwsAttributes{
						InstanceProfileArn: "abc",
					},
				},
			},
		},
		Resource: ResourceMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := internal.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			return common.CommandResults{
				ResultType: "error",
				Summary:    "Mount not found",
			}
		},
		State: map[string]interface{}{
			"cluster_id": "this_cluster",
			"name":       "this_mount",
			"s3": []interface{}{map[string]interface{}{
				"bucket_name": testS3BucketName,
			}},
		},
		ID:      "this_mount",
		Read:    true,
		Removed: true,
	}.ApplyNoError(t)
}

func TestResourceAwsS3MountGenericRead_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=this_cluster",
				Response: clusters.ClusterInfo{
					State: clusters.ClusterStateRunning,
					AwsAttributes: &clusters.AwsAttributes{
						InstanceProfileArn: "abc",
					},
				},
			},
		},
		Resource: ResourceMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := internal.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			return common.CommandResults{
				ResultType: "error",
				Summary:    "Some error",
			}
		},
		State: map[string]interface{}{
			"cluster_id": "this_cluster",
			"name":       "this_mount",
			"s3": []interface{}{map[string]interface{}{
				"bucket_name": testS3BucketName,
			}},
		},
		ID:   "this_mount",
		Read: true,
	}.ExpectError(t, "Some error")
}

func TestResourceAzureStorageUpdate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=this_cluster",
				Response: clusters.ClusterInfo{
					State: clusters.ClusterStateRunning,
					AwsAttributes: &clusters.AwsAttributes{
						InstanceProfileArn: "abc",
					},
				},
			},
		},
		Resource: ResourceMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := internal.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			require.True(t, strings.Contains(commandStr, "dbutils.fs.updateMount"), 
				"Cannot find update mount call")
			return mockMountInfo(testS3BucketPath, "a1b2c3")
		},
		HCL: `
		cluster_id = "this_cluster"
		name = "this_mount"
		abfs {
			client_id = "a"
			tenant_id = "b"
			client_secret_key = "c"
			client_secret_scope = "d"
			container_name = "e"
			storage_account_name = "f"
		}
		`,
		InstanceState: map[string]string{
			"name": "this_mount",
			"cluster_id": "this_cluster",
		},
		ID:      "this_mount",
		Update:    true,
	}.ApplyNoError(t)
}

func TestResourceAwsS3MountDeleteGeneric(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=this_cluster",
				Response: clusters.ClusterInfo{
					State: clusters.ClusterStateRunning,
					AwsAttributes: &clusters.AwsAttributes{
						InstanceProfileArn: "abc",
					},
				},
			},
		},
		Resource: ResourceMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := internal.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			assert.Contains(t, trunc, "/mnt/this_mount")
			assert.Contains(t, trunc, "dbutils.fs.unmount(mount_point)")
			return common.CommandResults{
				ResultType: "text",
				Data:       "",
			}
		},
		State: map[string]interface{}{
			"cluster_id": "this_cluster",
			"name":       "this_mount",
			"s3": []interface{}{map[string]interface{}{
				"bucket_name": testS3BucketName,
			}},
		},
		ID:     "this_mount",
		Delete: true,
	}.ApplyNoError(t)
}
