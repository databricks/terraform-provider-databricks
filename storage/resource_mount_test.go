package storage

import (
	"context"
	"strings"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/compute"
	"github.com/databrickslabs/terraform-provider-databricks/identity"
	"github.com/databrickslabs/terraform-provider-databricks/internal"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ============================== AWS S3 Tests ==============================

// Test interface compliance via compile time error
var _ Mount = (*S3IamMount)(nil)

var sparkVersionsResponse = compute.SparkVersionsList{
	SparkVersions: []compute.SparkVersion{
		{
			Version:     "7.3.x-scala2.12",
			Description: "7.3 LTS (includes Apache Spark 3.0.1, Scala 2.12)",
		},
	},
}

var nodeListResponse = compute.NodeTypeList{
	NodeTypes: []compute.NodeType{
		{
			NodeTypeID:     "Standard_F4s",
			InstanceTypeID: "Standard_F4s",
			MemoryMB:       8192,
			NumCores:       4,
			NodeInstanceType: &compute.NodeInstanceType{
				LocalDisks:      1,
				LocalDiskSizeGB: 16,
				LocalNVMeDisks:  0,
			},
		},
	},
}

func TestResourceAwsS3MountGenericCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=this_cluster",
				Response: compute.ClusterInfo{
					State: compute.ClusterStateRunning,
					AwsAttributes: &compute.AwsAttributes{
						InstanceProfileArn: "abc",
					},
				},
			},
		},
		Resource: ResourceDatabricksMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := internal.TrimLeadingWhitespace(commandStr)
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
		State: map[string]interface{}{
			"cluster_id": "this_cluster",
			"mount_name": "this_mount",
			"s3": []interface{}{map[string]interface{}{
				"bucket_name": testS3BucketName,
			}},
		},
		Create: true,
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, "this_mount", d.Id())
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
				Response: compute.ClusterInfo{
					ClusterID: "abcd",
					State:     compute.ClusterStateRunning,
					AwsAttributes: &compute.AwsAttributes{
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
				Response: compute.ClusterList{
					Clusters: []compute.ClusterInfo{
						{
							ClusterID: "abcd",
							State:     compute.ClusterStateRunning,
							AwsAttributes: &compute.AwsAttributes{
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
				ExpectedRequest: compute.Cluster{
					NodeTypeID: "Standard_F4s",
					AwsAttributes: &compute.AwsAttributes{
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
				Response: compute.ClusterID{
					ClusterID: "abcd",
				},
			},
		},
		Resource: ResourceDatabricksMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := internal.TrimLeadingWhitespace(commandStr)
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
		State: map[string]interface{}{
			"mount_name": "this_mount",
			"s3": []interface{}{map[string]interface{}{
				"bucket_name":      testS3BucketName,
				"instance_profile": instance_profile,
			}},
		},
		Create: true,
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, "this_mount", d.Id())
	assert.Equal(t, "abcd", d.Get("cluster_id"))
	assert.Equal(t, testS3BucketPath, d.Get("source"))
}

func TestResourceAwsS3MountGenericCreate_nothing_specified(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource: ResourceDatabricksMount(),
		State: map[string]interface{}{
			"mount_name": "this_mount",
			"s3": []interface{}{map[string]interface{}{
				"bucket_name": testS3BucketName,
			}},
		},
		Create: true,
	}.Apply(t)
	require.EqualError(t, err, "either cluster_id or instance_profile must be specified to mount S3 bucket")
}

func TestResourceAwsS3MountGenericCreate_invalid_arn(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource: ResourceDatabricksMount(),
		State: map[string]interface{}{
			"mount_name": "this_mount",
			"s3": []interface{}{map[string]interface{}{
				"bucket_name":      testS3BucketName,
				"instance_profile": "this_mount",
			}},
		},
		Create: true,
	}.Apply(t)
	require.EqualError(t, err, "invalid arn: this_mount")
}

func TestResourceAwsS3MountGenericRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=this_cluster",
				Response: compute.ClusterInfo{
					State: compute.ClusterStateRunning,
					AwsAttributes: &compute.AwsAttributes{
						InstanceProfileArn: "abc",
					},
				},
			},
		},
		Resource: ResourceDatabricksMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := internal.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			assert.Contains(t, trunc, "dbutils.fs.mounts()")
			assert.Contains(t, trunc, `mount.mountPoint == "/mnt/this_mount"`)
			return common.CommandResults{
				ResultType: "text",
				Data:       testS3BucketPath,
			}
		},
		State: map[string]interface{}{
			"cluster_id": "this_cluster",
			"mount_name": "this_mount",
			"s3": []interface{}{map[string]interface{}{
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
				Response: compute.ClusterInfo{
					State: compute.ClusterStateRunning,
					AwsAttributes: &compute.AwsAttributes{
						InstanceProfileArn: "abc",
					},
				},
			},
		},
		Resource: ResourceDatabricksMount(),
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
			"mount_name": "this_mount",
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
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=this_cluster",
				Response: compute.ClusterInfo{
					State: compute.ClusterStateRunning,
					AwsAttributes: &compute.AwsAttributes{
						InstanceProfileArn: "abc",
					},
				},
			},
		},
		Resource: ResourceDatabricksMount(),
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
			"mount_name": "this_mount",
			"s3": []interface{}{map[string]interface{}{
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
				Response: compute.ClusterInfo{
					State: compute.ClusterStateRunning,
					AwsAttributes: &compute.AwsAttributes{
						InstanceProfileArn: "abc",
					},
				},
			},
		},
		Resource: ResourceDatabricksMount(),
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
			"mount_name": "this_mount",
			"s3": []interface{}{map[string]interface{}{
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

func TestAwsAccS3MountGeneric(t *testing.T) {
	client := common.NewClientFromEnvironment()
	instanceProfile := qa.GetEnvOrSkipTest(t, "TEST_EC2_INSTANCE_PROFILE")
	ctx := context.WithValue(context.Background(), common.Current, t.Name())
	instanceProfilesAPI := identity.NewInstanceProfilesAPI(ctx, client)
	instanceProfilesAPI.Synchronized(instanceProfile, func() bool {
		if err := instanceProfilesAPI.Create(identity.InstanceProfileInfo{
			InstanceProfileArn: instanceProfile,
		}); err != nil {
			return false
		}
		bucket := qa.GetEnvOrSkipTest(t, "TEST_S3_BUCKET")
		client := compute.CommonEnvironmentClientWithRealCommandExecutor()
		clustersAPI := compute.NewClustersAPI(ctx, client)
		clusterInfo, err := GetOrCreateMountingClusterWithInstanceProfile(
			clustersAPI, instanceProfile)
		require.NoError(t, err)
		defer func() {
			err = clustersAPI.PermanentDelete(clusterInfo.ClusterID)
			assert.NoError(t, err)
			err = instanceProfilesAPI.Delete(instanceProfile)
			assert.NoError(t, err)
		}()
		testMounting(t, MountPoint{
			exec:      client.CommandExecutor(ctx),
			clusterID: clusterInfo.ClusterID,
			name:      qa.RandomName("t"),
		}, GenericMount{
			S3: &S3IamMount{BucketName: bucket},
		})
		return true
	})
}

// ============================== ADLS Gen1 Tests ==============================

func TestAzureAccADLSv1MountGeneric(t *testing.T) {
	client, mp := mountPointThroughReusedCluster(t)
	if !client.IsAzureClientSecretSet() {
		t.Skip("Test is meant only for client-secret conf Azure")
	}
	storageResource := qa.GetEnvOrSkipTest(t, "TEST_DATA_LAKE_STORE_NAME")
	testWithNewSecretScope(t, func(scope, key string) {
		testMounting(t, mp,
			GenericMount{Adl: &AzureADLSGen1MountGeneric{
				ClientID:        client.AzureClientID,
				TenantID:        client.AzureTenantID,
				PrefixType:      "dfs.adls",
				StorageResource: storageResource,
				Directory:       "/",
				SecretScope:     scope,
				SecretKey:       key,
			}})
	}, client, mp.name, client.AzureClientSecret)
}

func TestResourceAdlsGen1MountGeneric_Create(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=this_cluster",
				Response: compute.ClusterInfo{
					State: compute.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceDatabricksMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := internal.TrimLeadingWhitespace(commandStr)
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
		State: map[string]interface{}{
			"cluster_id": "this_cluster",
			"mount_name": "this_mount",
			"adl": []interface{}{map[string]interface{}{
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
	require.NoError(t, err, err)
	assert.Equal(t, "this_mount", d.Id())
	assert.Equal(t, testS3BucketPath, d.Get("source"))
}

// ============================== ADLS Gen2 Tests ==============================

func TestAzureAccADLSv2MountGeneric(t *testing.T) {
	client, mp := mountPointThroughReusedCluster(t)
	if !client.IsAzureClientSecretSet() {
		t.Skip("Test is meant only for client-secret conf Azure")
	}
	storageAccountName := qa.GetEnvOrSkipTest(t, "TEST_STORAGE_V2_ACCOUNT")
	container := qa.GetEnvOrSkipTest(t, "TEST_STORAGE_V2_ABFSS")
	testWithNewSecretScope(t, func(scope, key string) {
		testMounting(t, mp, GenericMount{Abfs: &AzureADLSGen2MountGeneric{
			ClientID:             client.AzureClientID,
			TenantID:             client.AzureTenantID,
			StorageAccountName:   storageAccountName,
			ContainerName:        container,
			SecretScope:          scope,
			SecretKey:            key,
			InitializeFileSystem: true,
			Directory:            "/",
		},
		})
	}, client, mp.name, client.AzureClientSecret)
}

func TestResourceAdlsGen2MountGeneric_Create(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=this_cluster",
				Response: compute.ClusterInfo{
					State: compute.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceDatabricksMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := internal.TrimLeadingWhitespace(commandStr)
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
		State: map[string]interface{}{
			"cluster_id": "this_cluster",
			"mount_name": "this_mount",
			"abfs": []interface{}{map[string]interface{}{
				"storage_account_name":   "test-adls-gen2",
				"container_name":         "e",
				"tenant_id":              "a",
				"client_id":              "b",
				"client_secret_scope":    "c",
				"client_secret_key":      "d",
				"initialize_file_system": true,
			}}},
		Create: true,
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, "this_mount", d.Id())
	assert.Equal(t, "abfss://e@test-adls-gen2.dfs.core.windows.net", d.Get("source"))
}

// ============================== Azure Blob Storage Tests ==============================

func TestResourceAzureBlobMountCreateGeneric(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=b",
				Response: compute.ClusterInfo{
					State: compute.ClusterStateRunning,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=b",
				Response: compute.ClusterInfo{
					State: compute.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceDatabricksMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := internal.TrimLeadingWhitespace(commandStr)
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
		State: map[string]interface{}{
			"cluster_id": "b",
			"mount_name": "e",
			"wasb": []interface{}{map[string]interface{}{
				"auth_type":            "ACCESS_KEY",
				"storage_account_name": "f",
				"token_secret_key":     "g",
				"token_secret_scope":   "h",
				"container_name":       "c",
				"directory":            "/d",
			},
			}},
		Create: true,
	}.Apply(t)
	require.NoError(t, err, err) // TODO: global search-replace for NoError
	assert.Equal(t, "e", d.Id())
	assert.Equal(t, "wasbs://c@f.blob.core.windows.net/d", d.Get("source"))
}

func TestResourceAzureBlobMountCreateGeneric_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=b",
				Response: compute.ClusterInfo{
					State: compute.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceDatabricksMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			return common.CommandResults{
				ResultType: "error",
				Summary:    "Some error",
			}
		},
		State: map[string]interface{}{
			"cluster_id": "b",
			"mount_name": "e",
			"wasb": []interface{}{map[string]interface{}{
				"container_name":       "c",
				"auth_type":            "ACCESS_KEY",
				"directory":            "/d",
				"storage_account_name": "f",
				"token_secret_key":     "g",
				"token_secret_scope":   "h",
			}}},
		Create: true,
	}.Apply(t)
	require.EqualError(t, err, "Some error")
	assert.Equal(t, "e", d.Id())
	assert.Equal(t, "", d.Get("source"))
}

func TestResourceAzureBlobMountGeneric_Read(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=b",
				Response: compute.ClusterInfo{
					State: compute.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceDatabricksMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := internal.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			assert.Contains(t, trunc, "dbutils.fs.mounts()")
			assert.Contains(t, trunc, `mount.mountPoint == "/mnt/e"`)
			return common.CommandResults{
				ResultType: "text",
				Data:       "wasbs://c@f.blob.core.windows.net/d",
			}
		},
		State: map[string]interface{}{
			"cluster_id": "b",
			"mount_name": "e",
			"wasb": []interface{}{map[string]interface{}{
				"auth_type":            "ACCESS_KEY",
				"container_name":       "c",
				"directory":            "/d",
				"storage_account_name": "f",
				"token_secret_key":     "g",
				"token_secret_scope":   "h",
			}},
		},
		ID:   "e",
		Read: true,
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
				Response: compute.ClusterInfo{
					State: compute.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceDatabricksMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := internal.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			return common.CommandResults{
				ResultType: "error",
				Summary:    "Mount not found",
			}
		},
		State: map[string]interface{}{
			"cluster_id": "b",
			"mount_name": "e",
			"wasb": []interface{}{map[string]interface{}{
				"auth_type":            "ACCESS_KEY",
				"container_name":       "c",
				"directory":            "/d",
				"storage_account_name": "f",
				"token_secret_key":     "g",
				"token_secret_scope":   "h",
			}},
		},
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
				Response: compute.ClusterInfo{
					State: compute.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceDatabricksMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := internal.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			return common.CommandResults{
				ResultType: "error",
				Summary:    "Some error",
			}
		},
		State: map[string]interface{}{
			"cluster_id": "b",
			"mount_name": "e",
			"wasb": []interface{}{map[string]interface{}{
				"auth_type":            "ACCESS_KEY",
				"container_name":       "c",
				"directory":            "/d",
				"storage_account_name": "f",
				"token_secret_key":     "g",
				"token_secret_scope":   "h",
			}},
		},
		ID:   "e",
		Read: true,
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
				Response: compute.ClusterInfo{
					State: compute.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceDatabricksMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := internal.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			assert.Contains(t, trunc, "dbutils.fs.unmount(mount_point)")
			return common.CommandResults{
				ResultType: "Text",
				Data:       "",
			}
		},
		State: map[string]interface{}{
			"cluster_id": "b",
			"mount_name": "e",
			"wasb": []interface{}{map[string]interface{}{
				"auth_type":            "ACCESS_KEY",
				"container_name":       "c",
				"directory":            "/d",
				"storage_account_name": "f",
				"token_secret_key":     "g",
				"token_secret_scope":   "h",
			}},
		},
		ID:     "e",
		Delete: true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "e", d.Id())
	assert.Equal(t, "", d.Get("source"))
}

func TestAzureAccBlobMountGeneric(t *testing.T) {
	client, mp := mountPointThroughReusedCluster(t)
	storageAccountName := qa.GetEnvOrSkipTest(t, "TEST_STORAGE_V2_ACCOUNT")
	accountKey := qa.GetEnvOrSkipTest(t, "TEST_STORAGE_V2_KEY")
	container := qa.GetEnvOrSkipTest(t, "TEST_STORAGE_V2_WASBS")
	testWithNewSecretScope(t, func(scope, key string) {
		testMounting(t, mp, GenericMount{Wasb: &AzureBlobMountGeneric{
			StorageAccountName: storageAccountName,
			ContainerName:      container,
			SecretScope:        scope,
			SecretKey:          key,
			Directory:          "/",
		}})
	}, client, mp.name, accountKey)
}

// ============================== Google Cloud Storage Tests ==============================

const testGcsBucketPath = "gs://" + testS3BucketName

func TestResourceGcsMountGenericCreate_WithCluster(t *testing.T) {
	google_account := "acc@acc-dbx.iam.gserviceaccount.com"
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=this_cluster",
				Response: compute.ClusterInfo{
					State: compute.ClusterStateRunning,
					GcpAttributes: &compute.GcpAttributes{
						GoogleServiceAccount: google_account,
					},
				},
			},
		},
		Resource: ResourceDatabricksMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := internal.TrimLeadingWhitespace(commandStr)
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
		State: map[string]interface{}{
			"cluster_id": "this_cluster",
			"mount_name": "this_mount",
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

func TestResourceGcsMountGenericCreate_WithServiceAccount(t *testing.T) {
	google_account := "acc@acc-dbx.iam.gserviceaccount.com"
	clusterName := "terraform-mount-gcs-bcb24f32098efa4172f435adbed2dae2"
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=abcd",
				Response: compute.ClusterInfo{
					ClusterID: "abcd",
					State:     compute.ClusterStateRunning,
					GcpAttributes: &compute.GcpAttributes{
						GoogleServiceAccount: google_account,
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
				Response: compute.ClusterList{
					Clusters: []compute.ClusterInfo{
						{
							ClusterID: "abcd",
							State:     compute.ClusterStateRunning,
							GcpAttributes: &compute.GcpAttributes{
								GoogleServiceAccount: google_account,
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
				ExpectedRequest: compute.Cluster{
					NodeTypeID: "Standard_F4s",
					GcpAttributes: &compute.GcpAttributes{
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
				Response: compute.ClusterID{
					ClusterID: "abcd",
				},
			},
		},
		Resource: ResourceDatabricksMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := internal.TrimLeadingWhitespace(commandStr)
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
		State: map[string]interface{}{
			"mount_name": "this_mount",
			"gs": []interface{}{map[string]interface{}{
				"bucket_name":     testS3BucketName,
				"service_account": google_account,
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
		Resource: ResourceDatabricksMount(),
		State: map[string]interface{}{
			"mount_name": "this_mount",
			"gs": []interface{}{map[string]interface{}{
				"bucket_name": testS3BucketName,
			}},
		},
		Create: true,
	}.Apply(t)
	require.EqualError(t, err, "either cluster_id or service_account must be specified to mount GCS bucket")
}

// TODO: implement it
// func TestGcpAccGcsMount(t *testing.T) {
// 	client, mp := mountPointThroughReusedCluster(t)
// 	storageAccountName := qa.GetEnvOrSkipTest(t, "TEST_STORAGE_V2_ACCOUNT")
// 	accountKey := qa.GetEnvOrSkipTest(t, "TEST_STORAGE_V2_KEY")
// 	container := qa.GetEnvOrSkipTest(t, "TEST_STORAGE_V2_WASBS")
// 	testWithNewSecretScope(t, func(scope, key string) {
// 		testMounting(t, mp, GenericMount{Wasb: &AzureBlobMount{
// 			StorageAccountName: storageAccountName,
// 			ContainerName:      container,
// 			SecretScope:        scope,
// 			SecretKey:          key,
// 			Directory:          "/",
// 		}})
// 	}, client, mp.name, accountKey)
// }

// ============================== Tests for Generic Configuration options ==============================

func TestResourceMountGenericCreate_WithUriAndOpts(t *testing.T) {
	abfssPath := "abfss://test@$test.dfs.core.windows.net"
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=this_cluster",
				Response: compute.ClusterInfo{
					State: compute.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceDatabricksMount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := internal.TrimLeadingWhitespace(commandStr)
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
		State: map[string]interface{}{
			"cluster_id": "this_cluster",
			"mount_name": "this_mount",
			"uri":        abfssPath,
			"extra_configs": map[string]interface{}{
				"fs.azure.account.auth.type": "CustomAccessToken",
			},
		},
		Create: true,
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, "this_mount", d.Id())
	assert.Equal(t, abfssPath, d.Get("source"))
}
