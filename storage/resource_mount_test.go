package storage

import (
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/commands"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ============================== AWS S3 Tests ==============================

// Test interface compliance via compile time error
var _ Mount = (*S3IamMount)(nil)

var sparkVersionsResponse = compute.GetSparkVersionsResponse{
	Versions: []compute.SparkVersion{
		{
			Key:  "7.3.x-scala2.12",
			Name: "7.3 LTS (includes Apache Spark 3.0.1, Scala 2.12)",
		},
	},
}

var nodeListResponse = compute.ListNodeTypesResponse{
	NodeTypes: []compute.NodeType{
		{
			NodeTypeId:     "Standard_F4s",
			InstanceTypeId: "Standard_F4s",
			MemoryMb:       8192,
			NumCores:       4,
			NodeInstanceType: &compute.NodeInstanceType{
				LocalDisks:      1,
				LocalDiskSizeGb: 16,
			},
		},
	},
}

func TestS3MountDefaults(t *testing.T) {
	s := ResourceDatabricksMountSchema()
	d := schema.TestResourceDataRaw(t, s, map[string]any{})
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{})
	defer server.Close()
	require.NoError(t, err)

	err = S3IamMount{}.ValidateAndApplyDefaults(d, client)
	qa.AssertErrorStartsWith(t, err, "'name' is not detected & it's impossible to infer it")

	d = schema.TestResourceDataRaw(t, s, map[string]any{"name": "test"})
	err = S3IamMount{}.ValidateAndApplyDefaults(d, client)
	require.NoError(t, err)
	assert.Equal(t, d.Get("name").(string), "test")
	d = schema.TestResourceDataRaw(t, s, map[string]any{})
	err = S3IamMount{BucketName: "abc"}.ValidateAndApplyDefaults(d, client)
	require.NoError(t, err)
	assert.Equal(t, d.Get("name").(string), "abc")
}

func TestResourceAwsS3MountGenericCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
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
			trunc := commands.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			if strings.HasPrefix(trunc, "def safe_mount") {
				assert.Contains(t, trunc, testS3BucketPath) // bucketname
				assert.Contains(t, trunc, `{}`)             // empty brackets for empty config
			}
			assert.Contains(t, trunc, "/mnt/this_mount")
			return common.CommandResults{
				ResultType: "text",
				Data:       testS3BucketPath,
			}
		},
		State: map[string]any{
			"cluster_id": "this_cluster",
			"name":       "this_mount",
			"s3": []any{map[string]any{
				"bucket_name": testS3BucketName,
			}},
		},
		Create: true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "this_mount", d.Id())
	assert.Equal(t, testS3BucketPath, d.Get("source"))
}

func TestResourceAwsS3MountGenericCreateWithInvalidClusterId(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=this_cluster",
				Status:       404,
			},
		},
		Resource: ResourceMount(),
		State: map[string]any{
			"cluster_id": "this_cluster",
			"name":       "this_mount",
			"s3": []any{map[string]any{
				"bucket_name": testS3BucketName,
			}},
		},
		Create: true,
	}.Apply(t)
	assert.EqualError(t, err, "instance profile is required to re-create mounting cluster")
}

func TestResourceAwsS3MountGenericReadWithInvalidClusterId(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=this_cluster",
				Status:       404,
			},
		},
		Resource: ResourceMount(),
		InstanceState: map[string]string{
			"cluster_id": "this_cluster",
			"name":       "this_mount",
			"source":     testS3BucketPath,
		},
		ID:      "this_id_should_be_unset",
		Removed: true,
		Read:    true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "", d.Id())
}

func TestResourceAwsS3MountGenericDeleteWithInvalidClusterId(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=this_cluster",
				Status:       404,
			},
		},
		Resource: ResourceMount(),
		InstanceState: map[string]string{
			"cluster_id": "this_cluster",
			"name":       "this_mount",
			"source":     testS3BucketPath,
		},
		ID:      "this_id_should_be_unset",
		Removed: true,
		Delete:  true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "", d.Id())
}

func TestResourceAwsS3MountGenericCreate_NoName(t *testing.T) {
	d, err := qa.ResourceFixture{
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
			trunc := commands.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			if strings.HasPrefix(trunc, "def safe_mount") {
				assert.Contains(t, trunc, testS3BucketPath) // bucketname
				assert.Contains(t, trunc, `{}`)             // empty brackets for empty config
			}
			assert.Contains(t, trunc, "/mnt/"+testS3BucketName)
			return common.CommandResults{
				ResultType: "text",
				Data:       testS3BucketPath,
			}
		},
		State: map[string]any{
			"cluster_id": "this_cluster",
			"s3": []any{map[string]any{
				"bucket_name": testS3BucketName,
			}},
		},
		Create: true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, testS3BucketName, d.Id())
	assert.Equal(t, testS3BucketPath, d.Get("source"))
}

func TestResourceAwsS3MountGenericCreate_WithInstanceProfile(t *testing.T) {
	instance_profile := "arn:aws:iam::1234567:instance-profile/s3-access"
	clusterName := "terraform-mount-s3-access"
	d, err := qa.ResourceFixture{
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
				Response: map[string]any{},
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
								ZoneID:             "auto",
							},
							AutoterminationMinutes: 10,
							SparkConf: map[string]string{
								"spark.databricks.cluster.profile": "singleNode",
								"spark.master":                     "local[*]", "spark.scheduler.mode": "FIFO",
							},
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
				Resource:     "/api/2.1/clusters/spark-versions",
				Response:     sparkVersionsResponse,
				ReuseRequest: true,
			},
			{
				Method:       "GET",
				Resource:     "/api/2.1/clusters/list-node-types",
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
						ZoneID:             "auto",
					},
					AutoterminationMinutes: 10,
					SparkConf: map[string]string{
						"spark.databricks.cluster.profile": "singleNode",
						"spark.master":                     "local[*]", "spark.scheduler.mode": "FIFO",
					},
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
			trunc := commands.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			if strings.HasPrefix(trunc, "def safe_mount") {
				assert.Contains(t, trunc, testS3BucketPath) // bucketname
				assert.Contains(t, trunc, `{}`)             // empty brackets for empty config
			}
			assert.Contains(t, trunc, "/mnt/this_mount")
			return common.CommandResults{
				ResultType: "text",
				Data:       testS3BucketPath,
			}
		},
		State: map[string]any{
			"name": "this_mount",
			"s3": []any{map[string]any{
				"bucket_name":      testS3BucketName,
				"instance_profile": instance_profile,
			}},
		},
		Create: true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "this_mount", d.Id())
	assert.Equal(t, "abcd", d.Get("cluster_id"))
	assert.Equal(t, testS3BucketPath, d.Get("source"))
}

func TestResourceAwsS3MountGenericCreate_nothing_specified(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource: ResourceMount(),
		State: map[string]any{
			"name": "this_mount",
			"s3": []any{map[string]any{
				"bucket_name": testS3BucketName,
			}},
		},
		Create: true,
	}.Apply(t)
	require.EqualError(t, err, "either cluster_id or instance_profile must be specified to mount S3 bucket")
}

func TestResourceAwsS3MountGenericCreate_invalid_arn(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource: ResourceMount(),
		State: map[string]any{
			"name": "this_mount",
			"s3": []any{map[string]any{
				"bucket_name":      testS3BucketName,
				"instance_profile": "this_mount",
			}},
		},
		Create: true,
	}.Apply(t)
	require.EqualError(t, err, "mount via profile: invalid arn: this_mount")
}

func TestResourceAwsS3MountGenericRead(t *testing.T) {
	d, err := qa.ResourceFixture{
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
			trunc := commands.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			assert.Contains(t, trunc, "dbutils.fs.mounts()")
			assert.Contains(t, trunc, `mount.mountPoint == "/mnt/this_mount"`)
			return common.CommandResults{
				ResultType: "text",
				Data:       testS3BucketPath,
			}
		},
		State: map[string]any{
			"cluster_id": "this_cluster",
			"name":       "this_mount",
			"s3": []any{map[string]any{
				"bucket_name": testS3BucketName,
			}},
		},
		ID:   "this_mount",
		Read: true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "this_mount", d.Id())
	assert.Equal(t, testS3BucketPath, d.Get("source"))
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
			trunc := commands.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			return common.CommandResults{
				ResultType: "error",
				Summary:    "Mount not found",
			}
		},
		State: map[string]any{
			"cluster_id": "this_cluster",
			"name":       "this_mount",
			"s3": []any{map[string]any{
				"bucket_name": testS3BucketName,
			}},
		},
		ID:      "this_mount",
		Read:    true,
		Removed: true,
	}.ApplyNoError(t)
}

func TestResourceAwsS3MountGenericRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
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
			trunc := commands.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			return common.CommandResults{
				ResultType: "error",
				Summary:    "Some error",
			}
		},
		State: map[string]any{
			"cluster_id": "this_cluster",
			"name":       "this_mount",
			"s3": []any{map[string]any{
				"bucket_name": testS3BucketName,
			}},
		},
		ID:   "this_mount",
		Read: true,
	}.Apply(t)
	require.EqualError(t, err, "Some error")
	assert.Equal(t, "this_mount", d.Id())
	assert.Equal(t, "", d.Get("source"))
}

func TestResourceAwsS3MountDeleteGeneric(t *testing.T) {
	d, err := qa.ResourceFixture{
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
			trunc := commands.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			assert.Contains(t, trunc, "/mnt/this_mount")
			assert.Contains(t, trunc, "dbutils.fs.unmount(mount_point)")
			return common.CommandResults{
				ResultType: "text",
				Data:       "",
			}
		},
		State: map[string]any{
			"cluster_id": "this_cluster",
			"name":       "this_mount",
			"s3": []any{map[string]any{
				"bucket_name": testS3BucketName,
			}},
		},
		ID:     "this_mount",
		Delete: true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "this_mount", d.Id())
	assert.Equal(t, "", d.Get("source"))
}

// ============================== ADLS Gen1 Tests ==============================

func TestResourceAdlsGen1MountGeneric_Create(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=this_cluster",
				Response: clusters.ClusterInfo{
					State: clusters.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := commands.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			if strings.HasPrefix(trunc, "def safe_mount") {
				assert.Contains(t, trunc, "adl://test-adls.azuredatalakestore.net")
				assert.Contains(t, trunc, `"fs.adl.oauth2.credential":dbutils.secrets.get("c", "d")`)
			}
			assert.Contains(t, trunc, "/mnt/this_mount")
			return common.CommandResults{
				ResultType: "text",
				Data:       testS3BucketPath,
			}
		},
		State: map[string]any{
			"cluster_id": "this_cluster",
			"name":       "this_mount",
			"adl": []any{map[string]any{
				"storage_resource_name": "test-adls",
				"tenant_id":             "a",
				"client_id":             "b",
				"client_secret_scope":   "c",
				"client_secret_key":     "d",
				"spark_conf_prefix":     "fs.adl",
			}},
		},
		Create: true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "this_mount", d.Id())
}

func TestResourceAdlsGen1MountGeneric_Create_ResourceID(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=this_cluster",
				Response: clusters.ClusterInfo{
					State: clusters.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := commands.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			if strings.HasPrefix(trunc, "def safe_mount") {
				assert.Contains(t, trunc, "adl://gen1.azuredatalakestore.net")
				assert.Contains(t, trunc, `"fs.adl.oauth2.credential":dbutils.secrets.get("c", "d")`)
			}
			assert.Contains(t, trunc, "/mnt/gen1")
			return common.CommandResults{
				ResultType: "text",
				Data:       testS3BucketPath,
			}
		},
		State: map[string]any{
			"cluster_id":  "this_cluster",
			"resource_id": "/subscriptions/123/resourceGroups/some-rg/providers/Microsoft.DataLakeStore/accounts/gen1",
			"adl": []any{map[string]any{
				"tenant_id":           "a",
				"client_id":           "b",
				"client_secret_scope": "c",
				"client_secret_key":   "d",
				"spark_conf_prefix":   "fs.adl",
			}},
		},
		Create: true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "gen1", d.Id())
}

func TestResourceAdlsGen1MountGeneric_Create_ResourceID_Error1(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{},
		Resource: ResourceMount(),
		State: map[string]any{
			"resource_id": "/subscriptions/123/resourceGroups/some-rg/providers/Microsoft.DataLakeStore/acc/gen1",
			"adl": []any{map[string]any{
				"tenant_id":           "a",
				"client_id":           "b",
				"client_secret_scope": "c",
				"client_secret_key":   "d",
				"spark_conf_prefix":   "fs.adl",
			}},
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "incorrect resource type or provider in resource_id:")
}

func TestResourceAdlsGen1MountGeneric_Create_ResourceID_Error2(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{},
		Resource: ResourceMount(),
		State: map[string]any{
			"adl": []any{map[string]any{
				"tenant_id":           "a",
				"client_id":           "b",
				"client_secret_scope": "c",
				"client_secret_key":   "d",
				"spark_conf_prefix":   "fs.adl",
			}},
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "storage_resource_name is empty, and resource_id or uri aren't specified")
}

func TestResourceAdlsGen1MountGeneric_Create_NoTenantID_Error(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource: ResourceMount(),
		Azure:    true,
		State: map[string]any{
			"resource_id": "/subscriptions/123/resourceGroups/some-rg/providers/Microsoft.DataLakeStore/accounts/gen1",
			"cluster_id":  "this_cluster",
			"name":        "this_mount",
			"adl": []any{map[string]any{
				"client_id":           "b",
				"client_secret_scope": "c",
				"client_secret_key":   "d",
				"spark_conf_prefix":   "fs.adl",
			}},
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "tenant_id is not defined, and we can't extract it: token contains an invalid number of segments")
}

func TestResourceAdlsGen1MountGeneric_Create_NoTenantID_Error_EmptyTenant(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource: ResourceMount(),
		Token:    "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE2MzU3ODQxNTYsImV4cCI6MTY2NzMyMDE1NiwiYXVkIjoid3d3LmV4YW1wbGUuY29tIiwic3ViIjoianJvY2tldEBleGFtcGxlLmNvbSIsInRpZCI6IiAgIn0.faxuGAFghVxa1epYnovOxoQrzju7-z_EJj3oZtwIxdk",
		Azure:    true,
		State: map[string]any{
			"resource_id": "/subscriptions/123/resourceGroups/some-rg/providers/Microsoft.DataLakeStore/accounts/gen1",
			"cluster_id":  "this_cluster",
			"name":        "this_mount",
			"adl": []any{map[string]any{
				"client_id":           "b",
				"client_secret_scope": "c",
				"client_secret_key":   "d",
				"spark_conf_prefix":   "fs.adl",
			}},
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "tenant_id is not defined, and we can't extract it: tenant_id isn't provided & we can't detect it")
}

// ============================== ADLS Gen2 Tests ==============================

func TestResourceAdlsGen2MountGeneric_Create(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=this_cluster",
				Response: clusters.ClusterInfo{
					State: clusters.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := commands.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			if strings.HasPrefix(trunc, "def safe_mount") {
				assert.Contains(t, trunc, "abfss://e@test-adls-gen2.dfs.core.windows.net")
				assert.Contains(t, trunc, `"fs.azure.account.oauth2.client.secret":dbutils.secrets.get("c", "d")`)
			}
			assert.Contains(t, trunc, "/mnt/this_mount")
			return common.CommandResults{
				ResultType: "text",
				Data:       "abfss://e@test-adls-gen2.dfs.core.windows.net",
			}
		},
		State: map[string]any{
			"cluster_id": "this_cluster",
			"name":       "this_mount",
			"abfs": []any{map[string]any{
				"storage_account_name":   "test-adls-gen2",
				"container_name":         "e",
				"tenant_id":              "a",
				"client_id":              "b",
				"client_secret_scope":    "c",
				"client_secret_key":      "d",
				"initialize_file_system": true,
			}},
		},
		Azure:  true,
		Create: true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "this_mount", d.Id())
	assert.Equal(t, "abfss://e@test-adls-gen2.dfs.core.windows.net", d.Get("source"))
}

func TestResourceAdlsGen2MountGeneric_Create_ResourceID(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=this_cluster",
				Response: clusters.ClusterInfo{
					State: clusters.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := commands.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			if strings.HasPrefix(trunc, "def safe_mount") {
				assert.Contains(t, trunc, "abfss://e@test-adls-gen2.dfs.core.windows.net")
				assert.Contains(t, trunc, `"fs.azure.account.oauth2.client.secret":dbutils.secrets.get("c", "d")`)
			}
			assert.Contains(t, trunc, "/mnt/e")
			return common.CommandResults{
				ResultType: "text",
				Data:       "abfss://e@test-adls-gen2.dfs.core.windows.net",
			}
		},
		State: map[string]any{
			"cluster_id":  "this_cluster",
			"resource_id": "/subscriptions/123/resourceGroups/some-rg/providers/Microsoft.Storage/storageAccounts/test-adls-gen2/blobServices/default/containers/e",
			"abfs": []any{map[string]any{
				"tenant_id":              "a",
				"client_id":              "b",
				"client_secret_scope":    "c",
				"client_secret_key":      "d",
				"initialize_file_system": true,
			}},
		},
		Create: true,
		Azure:  true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "e", d.Id())
	assert.Equal(t, "abfss://e@test-adls-gen2.dfs.core.windows.net", d.Get("source"))
}

func TestResourceAdlsGen2MountGeneric_Create_NoTenantID_SPN(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=this_cluster",
				Response: clusters.ClusterInfo{
					State: clusters.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := commands.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			if strings.HasPrefix(trunc, "def safe_mount") {
				assert.Contains(t, trunc, "abfss://e@test-adls-gen2.dfs.core.windows.net")
				assert.Contains(t, trunc, `"fs.azure.account.oauth2.client.secret":dbutils.secrets.get("c", "d")`)
				assert.Contains(t, trunc, "https://login.microsoftonline.com/c/oauth2/token")
			}
			assert.Contains(t, trunc, "/mnt/this_mount")
			return common.CommandResults{
				ResultType: "text",
				Data:       "abfss://e@test-adls-gen2.dfs.core.windows.net",
			}
		},
		AzureSPN: true,
		State: map[string]any{
			"cluster_id": "this_cluster",
			"name":       "this_mount",
			"abfs": []any{map[string]any{
				"storage_account_name":   "test-adls-gen2",
				"container_name":         "e",
				"client_id":              "b",
				"client_secret_scope":    "c",
				"client_secret_key":      "d",
				"initialize_file_system": true,
			}},
		},
		Create: true,
		Azure:  true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "this_mount", d.Id())
	assert.Equal(t, "abfss://e@test-adls-gen2.dfs.core.windows.net", d.Get("source"))
}

func TestResourceAdlsGen2MountGeneric_Create_NoTenantID_CLI(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=this_cluster",
				Response: clusters.ClusterInfo{
					State: clusters.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := commands.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			if strings.HasPrefix(trunc, "def safe_mount") {
				assert.Contains(t, trunc, "abfss://e@test-adls-gen2.dfs.core.windows.net")
				assert.Contains(t, trunc, `"fs.azure.account.oauth2.client.secret":dbutils.secrets.get("c", "d")`)
				assert.Contains(t, trunc, "https://login.microsoftonline.com/abc/oauth2/token")
			}
			assert.Contains(t, trunc, "/mnt/this_mount")
			return common.CommandResults{
				ResultType: "text",
				Data:       "abfss://e@test-adls-gen2.dfs.core.windows.net",
			}
		},
		Azure: true,
		// sample JWT token for testing
		Token: "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE2MzU2OTU4MzksImV4cCI6MTY2NzIzMTgzOSwiYXVkIjoid3d3LmV4YW1wbGUuY29tIiwic3ViIjoianJvY2tldEBleGFtcGxlLmNvbSIsInRpZCI6ImFiYyJ9._G1DrR4DspidISpsra8UnecV_FV4zMlJDtSNzaS0UxI",
		State: map[string]any{
			"cluster_id": "this_cluster",
			"name":       "this_mount",
			"abfs": []any{map[string]any{
				"storage_account_name":   "test-adls-gen2",
				"container_name":         "e",
				"client_id":              "b",
				"client_secret_scope":    "c",
				"client_secret_key":      "d",
				"initialize_file_system": true,
			}},
		},
		Create: true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "this_mount", d.Id())
	assert.Equal(t, "abfss://e@test-adls-gen2.dfs.core.windows.net", d.Get("source"))
}

func TestResourceAdlsGen2MountGeneric_Create_NoTenantID_Error(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource: ResourceMount(),
		Azure:    true,
		State: map[string]any{
			"name": "this_mount",
			"abfs": []any{map[string]any{
				"storage_account_name":   "test-adls-gen2",
				"container_name":         "e",
				"client_id":              "b",
				"client_secret_scope":    "c",
				"client_secret_key":      "d",
				"initialize_file_system": true,
			}},
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "tenant_id is not defined, and we can't extract it: token contains an invalid number of segments")
}

func TestResourceAdlsGen2MountGeneric_Create_NoTenantID_Error_EmptyTenant(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource: ResourceMount(),
		Token:    "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE2MzU3ODQxNTYsImV4cCI6MTY2NzMyMDE1NiwiYXVkIjoid3d3LmV4YW1wbGUuY29tIiwic3ViIjoianJvY2tldEBleGFtcGxlLmNvbSIsInRpZCI6IiAgIn0.faxuGAFghVxa1epYnovOxoQrzju7-z_EJj3oZtwIxdk",
		Azure:    true,
		State: map[string]any{
			"name": "this_mount",
			"abfs": []any{map[string]any{
				"storage_account_name":   "test-adls-gen2",
				"container_name":         "e",
				"client_id":              "b",
				"client_secret_scope":    "c",
				"client_secret_key":      "d",
				"initialize_file_system": true,
			}},
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "tenant_id is not defined, and we can't extract it: tenant_id isn't provided & we can't detect it")
}

// ============================== Azure Blob Storage Tests ==============================

func TestResourceAzureBlobMountCreateGeneric(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=b",
				Response: clusters.ClusterInfo{
					State: clusters.ClusterStateRunning,
				},
				ReuseRequest: true,
			},
		},
		Resource: ResourceMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := commands.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)

			if strings.HasPrefix(trunc, "def safe_mount") {
				assert.Contains(t, trunc, "wasbs://c@f.blob.core.windows.net/d")
				assert.Contains(t, trunc, `"fs.azure.account.key.f.blob.core.windows.net":dbutils.secrets.get("h", "g")`)
			}
			assert.Contains(t, trunc, "/mnt/e")
			return common.CommandResults{
				ResultType: "text",
				Data:       "wasbs://c@f.blob.core.windows.net/d",
			}
		},
		State: map[string]any{
			"cluster_id": "b",
			"name":       "e",
			"wasb": []any{
				map[string]any{
					"auth_type":            "ACCESS_KEY",
					"storage_account_name": "f",
					"token_secret_key":     "g",
					"token_secret_scope":   "h",
					"container_name":       "c",
					"directory":            "/d",
				},
			},
		},
		Create: true,
		Azure:  true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "e", d.Id())
	assert.Equal(t, "wasbs://c@f.blob.core.windows.net/d", d.Get("source"))
}

func TestResourceAzureBlobMountCreateGeneric_SAS(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=b",
				Response: clusters.ClusterInfo{
					State: clusters.ClusterStateRunning,
				},
				ReuseRequest: true,
			},
		},
		Resource: ResourceMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := commands.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)

			if strings.HasPrefix(trunc, "def safe_mount") {
				assert.Contains(t, trunc, "wasbs://c@f.blob.core.windows.net/d")
				assert.Contains(t, trunc, `"fs.azure.sas.c.f.blob.core.windows.net":dbutils.secrets.get("h", "g")`)
			}
			assert.Contains(t, trunc, "/mnt/e")
			return common.CommandResults{
				ResultType: "text",
				Data:       "wasbs://c@f.blob.core.windows.net/d",
			}
		},
		State: map[string]any{
			"cluster_id": "b",
			"name":       "e",
			"wasb": []any{
				map[string]any{
					"auth_type":            "SAS",
					"storage_account_name": "f",
					"token_secret_key":     "g",
					"token_secret_scope":   "h",
					"container_name":       "c",
					"directory":            "/d",
				},
			},
		},
		Azure:  true,
		Create: true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "e", d.Id())
	assert.Equal(t, "wasbs://c@f.blob.core.windows.net/d", d.Get("source"))
}

func TestResourceAzureBlobMountCreateGeneric_Resource_ID(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=b",
				Response: clusters.ClusterInfo{
					State: clusters.ClusterStateRunning,
				},
				ReuseRequest: true,
			},
		},
		Resource: ResourceMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := commands.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)

			if strings.HasPrefix(trunc, "def safe_mount") {
				assert.Contains(t, trunc, "wasbs://c@f.blob.core.windows.net/d")
				assert.Contains(t, trunc, `"fs.azure.account.key.f.blob.core.windows.net":dbutils.secrets.get("h", "g")`)
			}
			assert.Contains(t, trunc, "/mnt/c")
			return common.CommandResults{
				ResultType: "text",
				Data:       "wasbs://c@f.blob.core.windows.net/d",
			}
		},
		State: map[string]any{
			"cluster_id":  "b",
			"resource_id": "/subscriptions/123/resourceGroups/some-rg/providers/Microsoft.Storage/storageAccounts/f/blobServices/default/containers/c",
			"wasb": []any{
				map[string]any{
					"auth_type":          "ACCESS_KEY",
					"token_secret_key":   "g",
					"token_secret_scope": "h",
					"directory":          "/d",
				},
			},
		},
		Azure:  true,
		Create: true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "c", d.Id())
	assert.Equal(t, "wasbs://c@f.blob.core.windows.net/d", d.Get("source"))
}

func TestResourceAzureBlobMountCreateGeneric_Resource_ID_Error(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource: ResourceMount(),
		State: map[string]any{
			"cluster_id":  "b",
			"resource_id": "abc",
			"wasb": []any{
				map[string]any{
					"auth_type":          "ACCESS_KEY",
					"token_secret_key":   "g",
					"token_secret_scope": "h",
					"directory":          "/d",
				},
			},
		},
		Azure:  true,
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "parsing failed for abc. Invalid container resource Id format")
}

func TestResourceAzureBlobMountCreateGeneric_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=b",
				Response: clusters.ClusterInfo{
					State: clusters.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			return common.CommandResults{
				ResultType: "error",
				Summary:    "Some error",
			}
		},
		State: map[string]any{
			"cluster_id": "b",
			"name":       "e",
			"wasb": []any{map[string]any{
				"container_name":       "c",
				"auth_type":            "ACCESS_KEY",
				"directory":            "/d",
				"storage_account_name": "f",
				"token_secret_key":     "g",
				"token_secret_scope":   "h",
			}},
		},
		Azure:  true,
		Create: true,
	}.Apply(t)
	require.EqualError(t, err, "Some error")
	assert.Equal(t, "e", d.Id())
	assert.Equal(t, "", d.Get("source"))
}

func TestResourceAzureBlobMountCreateGeneric_Error_NoResourceID(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=b",
				Response: clusters.ClusterInfo{
					State: clusters.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			return common.CommandResults{
				ResultType: "error",
				Summary:    "Some error",
			}
		},
		State: map[string]any{
			"cluster_id": "b",
			"name":       "e",
			"wasb": []any{map[string]any{
				"auth_type":          "ACCESS_KEY",
				"directory":          "/d",
				"token_secret_key":   "g",
				"token_secret_scope": "h",
			}},
		},
		Azure:  true,
		Create: true,
	}.Apply(t)
	require.EqualError(t, err, "container_name or storage_account_name are empty, and resource_id or uri aren't specified")
	assert.Equal(t, "", d.Get("source"))
}

func TestResourceAzureBlobMountGeneric_Read(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=b",
				Response: clusters.ClusterInfo{
					State: clusters.ClusterStateRunning,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=b",
				Response: clusters.ClusterInfo{
					State: clusters.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := commands.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			assert.Contains(t, trunc, "dbutils.fs.mounts()")
			assert.Contains(t, trunc, `mount.mountPoint == "/mnt/e"`)
			return common.CommandResults{
				ResultType: "text",
				Data:       "wasbs://c@f.blob.core.windows.net/d",
			}
		},
		State: map[string]any{
			"cluster_id": "b",
			"name":       "e",
			"wasb": []any{map[string]any{
				"auth_type":            "ACCESS_KEY",
				"container_name":       "c",
				"directory":            "/d",
				"storage_account_name": "f",
				"token_secret_key":     "g",
				"token_secret_scope":   "h",
			}},
		},
		Azure: true,
		ID:    "e",
		Read:  true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "e", d.Id())
	assert.Equal(t, "wasbs://c@f.blob.core.windows.net/d", d.Get("source"))
}

func TestResourceAzureBlobMountGenericRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=b",
				Response: clusters.ClusterInfo{
					State: clusters.ClusterStateRunning,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=b",
				Response: clusters.ClusterInfo{
					State: clusters.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := commands.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			return common.CommandResults{
				ResultType: "error",
				Summary:    "Mount not found",
			}
		},
		State: map[string]any{
			"cluster_id": "b",
			"name":       "e",
			"wasb": []any{map[string]any{
				"auth_type":            "ACCESS_KEY",
				"container_name":       "c",
				"directory":            "/d",
				"storage_account_name": "f",
				"token_secret_key":     "g",
				"token_secret_scope":   "h",
			}},
		},
		Azure:   true,
		ID:      "e",
		Read:    true,
		Removed: true,
	}.ApplyNoError(t)
}

func TestResourceAzureBlobMountGenericRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=b",
				Response: clusters.ClusterInfo{
					State: clusters.ClusterStateRunning,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=b",
				Response: clusters.ClusterInfo{
					State: clusters.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := commands.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			return common.CommandResults{
				ResultType: "error",
				Summary:    "Some error",
			}
		},
		State: map[string]any{
			"cluster_id": "b",
			"name":       "e",
			"wasb": []any{map[string]any{
				"auth_type":            "ACCESS_KEY",
				"container_name":       "c",
				"directory":            "/d",
				"storage_account_name": "f",
				"token_secret_key":     "g",
				"token_secret_scope":   "h",
			}},
		},
		Azure: true,
		ID:    "e",
		Read:  true,
	}.Apply(t)
	require.EqualError(t, err, "Some error")
	assert.Equal(t, "e", d.Id())
	assert.Equal(t, "", d.Get("source"))
}

func TestResourceAzureBlobMountGenericDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=b",
				Response: clusters.ClusterInfo{
					State: clusters.ClusterStateRunning,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=b",
				Response: clusters.ClusterInfo{
					State: clusters.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := commands.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			assert.Contains(t, trunc, "dbutils.fs.unmount(mount_point)")
			return common.CommandResults{
				ResultType: "Text",
				Data:       "",
			}
		},
		State: map[string]any{
			"cluster_id": "b",
			"name":       "e",
			"wasb": []any{map[string]any{
				"auth_type":            "ACCESS_KEY",
				"container_name":       "c",
				"directory":            "/d",
				"storage_account_name": "f",
				"token_secret_key":     "g",
				"token_secret_scope":   "h",
			}},
		},
		Azure:  true,
		ID:     "e",
		Delete: true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "e", d.Id())
	assert.Equal(t, "", d.Get("source"))
}

// ============================== Google Cloud Storage Tests ==============================

const testGcsBucketPath = "gs://" + testS3BucketName

func TestGSMountDefaults(t *testing.T) {
	s := ResourceDatabricksMountSchema()
	d := schema.TestResourceDataRaw(t, s, map[string]any{})
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{})
	defer server.Close()
	require.NoError(t, err)

	err = GSMount{}.ValidateAndApplyDefaults(d, client)
	qa.AssertErrorStartsWith(t, err, "'name' is not detected & it's impossible to infer it")

	d = schema.TestResourceDataRaw(t, s, map[string]any{"name": "test"})
	err = GSMount{}.ValidateAndApplyDefaults(d, client)
	require.NoError(t, err)
	assert.Equal(t, d.Get("name").(string), "test")
	d = schema.TestResourceDataRaw(t, s, map[string]any{})
	err = GSMount{BucketName: "abc"}.ValidateAndApplyDefaults(d, client)
	require.NoError(t, err)
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
			trunc := commands.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			if strings.HasPrefix(trunc, "def safe_mount") {
				assert.Contains(t, trunc, testGcsBucketPath) // bucketname
				assert.Contains(t, trunc, `{}`)              // empty brackets for empty config
			}
			assert.Contains(t, trunc, "/mnt/this_mount")
			return common.CommandResults{
				ResultType: "text",
				Data:       testGcsBucketPath,
			}
		},
		State: map[string]any{
			"cluster_id": "this_cluster",
			"name":       "this_mount",
			"gs": []any{map[string]any{
				"bucket_name": testS3BucketName,
			}},
		},
		Create: true,
	}.Apply(t)
	require.NoError(t, err)
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
			trunc := commands.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			if strings.HasPrefix(trunc, "def safe_mount") {
				assert.Contains(t, trunc, testGcsBucketPath) // bucketname
				assert.Contains(t, trunc, `{}`)              // empty brackets for empty config
			}
			assert.Contains(t, trunc, "/mnt/"+testS3BucketName)
			return common.CommandResults{
				ResultType: "text",
				Data:       testGcsBucketPath,
			}
		},
		State: map[string]any{
			"cluster_id": "this_cluster",
			"gs": []any{map[string]any{
				"bucket_name": testS3BucketName,
			}},
		},
		Create: true,
	}.Apply(t)
	require.NoError(t, err)
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
				Response: map[string]any{},
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
							SparkConf: map[string]string{
								"spark.databricks.cluster.profile": "singleNode",
								"spark.master":                     "local[*]", "spark.scheduler.mode": "FIFO",
							},
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
				Resource:     "/api/2.1/clusters/spark-versions",
				Response:     sparkVersionsResponse,
				ReuseRequest: true,
			},
			{
				Method:       "GET",
				Resource:     "/api/2.1/clusters/list-node-types",
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
					SparkConf: map[string]string{
						"spark.databricks.cluster.profile": "singleNode",
						"spark.master":                     "local[*]", "spark.scheduler.mode": "FIFO",
					},
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
			trunc := commands.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			if strings.HasPrefix(trunc, "def safe_mount") {
				assert.Contains(t, trunc, testGcsBucketPath) // bucketname
				assert.Contains(t, trunc, `{}`)              // empty brackets for empty config
			}
			assert.Contains(t, trunc, "/mnt/this_mount")
			return common.CommandResults{
				ResultType: "text",
				Data:       testGcsBucketPath,
			}
		},
		State: map[string]any{
			"name": "this_mount",
			"gs": []any{map[string]any{
				"bucket_name":     testS3BucketName,
				"service_account": googleAccount,
			}},
		},
		Create: true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "this_mount", d.Id())
	assert.Equal(t, "abcd", d.Get("cluster_id"))
	assert.Equal(t, testGcsBucketPath, d.Get("source"))
}

func TestResourceGcsMountGenericCreate_nothing_specified(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource: ResourceMount(),
		State: map[string]any{
			"name": "this_mount",
			"gs": []any{map[string]any{
				"bucket_name": testS3BucketName,
			}},
		},
		Create: true,
	}.Apply(t)
	require.EqualError(t, err, "either cluster_id or service_account must be specified to mount GCS bucket")
}

// ============================== Tests for Generic Configuration options ==============================

func TestResourceMountGenericCreate_WithUriAndOpts(t *testing.T) {
	abfssPath := "abfss://test@$test.dfs.core.windows.net"
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=this_cluster",
				Response: clusters.ClusterInfo{
					State: clusters.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := commands.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			if strings.HasPrefix(trunc, "def safe_mount") {
				assert.Contains(t, trunc, abfssPath) // URI
				assert.Contains(t, trunc, `{"fs.azure.account.auth.type":"CustomAccessToken"}`)
			}
			assert.Contains(t, trunc, "/mnt/this_mount")
			return common.CommandResults{
				ResultType: "text",
				Data:       abfssPath,
			}
		},
		State: map[string]any{
			"cluster_id": "this_cluster",
			"name":       "this_mount",
			"uri":        abfssPath,
			"extra_configs": map[string]any{
				"fs.azure.account.auth.type": "CustomAccessToken",
			},
		},
		Create: true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "this_mount", d.Id())
	assert.Equal(t, abfssPath, d.Get("source"))
}

func TestNames(t *testing.T) {
	mount_name := "abc"
	gm := GenericMount{MountName: mount_name}
	assert.Equal(t, gm.Name(), mount_name)
	assert.Equal(t, GenericMount{}.Name(), "")
	gm = GenericMount{Abfs: &AzureADLSGen2MountGeneric{ContainerName: mount_name}}
	assert.Equal(t, gm.Name(), mount_name)
	gm = GenericMount{Wasb: &AzureBlobMountGeneric{ContainerName: mount_name}}
	assert.Equal(t, gm.Name(), mount_name)
	gm = GenericMount{Adl: &AzureADLSGen1MountGeneric{StorageResource: mount_name}}
	assert.Equal(t, gm.Name(), mount_name)
}

func TestARMParsing(t *testing.T) {
	acc, container, err := parseStorageContainerId("/subscriptions/5363c143-2af7-4fb5-8a9d-ab1b2c8e756e/resourceGroups/test-rg/providers/Microsoft.Storage/storageAccounts/lrs-acc/blobServices/default/containers/test")
	require.NoError(t, err)
	assert.Equal(t, acc, "lrs-acc")
	assert.Equal(t, container, "test")
}

func TestARMParsingError(t *testing.T) {
	_, _, err := parseStorageContainerId("abc")
	qa.AssertErrorStartsWith(t, err, "parsing failed for ")
}

func TestARMParsing2(t *testing.T) {
	res, err := parseAzureResourceID("/subscriptions/6369c148-f8a9-4fb5-8a9d-ac1b2c8e756e/resourceGroups/alexott-rg/providers/Microsoft.DataLakeStore/accounts/aottgen1")
	require.NoError(t, err)
	assert.Equal(t, res.resourceName, "aottgen1")
}

func TestGenericMountDefaults(t *testing.T) {
	s := ResourceDatabricksMountSchema()
	d := schema.TestResourceDataRaw(t, s, map[string]any{})
	client, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{})
	defer server.Close()
	require.NoError(t, err)

	gm := GenericMount{MountName: "test"}
	err = common.StructToData(gm, s, d)
	require.NoError(t, err)
	err = gm.ValidateAndApplyDefaults(d, client)
	qa.AssertErrorStartsWith(t, err, "value of uri is not specified or empty")

	d = schema.TestResourceDataRaw(t, s, map[string]any{"uri": "s3://abc/"})
	err = gm.ValidateAndApplyDefaults(d, client)
	qa.AssertErrorStartsWith(t, err, "value of name is not specified or empty")

	gm = GenericMount{Abfs: &AzureADLSGen2MountGeneric{}}
	d = schema.TestResourceDataRaw(t, s, map[string]any{"abfs": map[string]any{}})
	err = gm.ValidateAndApplyDefaults(d, client)
	qa.AssertErrorStartsWith(t, err, "container_name or storage_account_name are empty, and resource_id or uri aren't specified")
}
