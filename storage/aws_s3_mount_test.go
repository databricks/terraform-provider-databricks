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

// Test interface compliance via compile time error
var _ Mount = (*AWSIamMount)(nil)

const testS3BucketName = "test-s3-bucket"
const testS3BucketPath = "s3a://" + testS3BucketName

func TestResourceAwsS3MountCreate(t *testing.T) {
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
		Resource: ResourceAWSS3Mount(),
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
			"cluster_id":     "this_cluster",
			"mount_name":     "this_mount",
			"s3_bucket_name": testS3BucketName,
		},
		Create: true,
	}.Apply(t)
	require.NoError(t, err, err)
	assert.Equal(t, "this_mount", d.Id())
	assert.Equal(t, testS3BucketPath, d.Get("source"))
}

func TestResourceAwsS3MountCreate_nothing_specified(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource: ResourceAWSS3Mount(),
		State: map[string]interface{}{
			"mount_name":     "this_mount",
			"s3_bucket_name": testS3BucketName,
		},
		Create: true,
	}.Apply(t)
	require.EqualError(t, err, "Either cluster_id or instance_profile must be specified")
}

func TestResourceAwsS3MountCreate_invalid_arn(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource: ResourceAWSS3Mount(),
		State: map[string]interface{}{
			"mount_name":       "this_mount",
			"s3_bucket_name":   testS3BucketName,
			"instance_profile": "this_mount",
		},
		Create: true,
	}.Apply(t)
	require.EqualError(t, err, "arn: invalid prefix")
}

func TestResourceAwsS3MountRead(t *testing.T) {
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
		Resource: ResourceAWSS3Mount(),
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
			"cluster_id":     "this_cluster",
			"mount_name":     "this_mount",
			"s3_bucket_name": testS3BucketName,
		},
		ID:   "this_mount",
		Read: true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "this_mount", d.Id())
	assert.Equal(t, testS3BucketPath, d.Get("source"))
}

func TestResourceAwsS3MountRead_NotFound(t *testing.T) {
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
		Resource: ResourceAWSS3Mount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := internal.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			return common.CommandResults{
				ResultType: "error",
				Summary:    "Mount not found",
			}
		},
		State: map[string]interface{}{
			"cluster_id":     "this_cluster",
			"mount_name":     "this_mount",
			"s3_bucket_name": testS3BucketName,
		},
		ID:      "this_mount",
		Read:    true,
		Removed: true,
	}.ApplyNoError(t)
}

func TestResourceAwsS3MountRead_Error(t *testing.T) {
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
		Resource: ResourceAWSS3Mount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := internal.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			return common.CommandResults{
				ResultType: "error",
				Summary:    "Some error",
			}
		},
		State: map[string]interface{}{
			"cluster_id":     "this_cluster",
			"mount_name":     "this_mount",
			"s3_bucket_name": testS3BucketName,
		},
		ID:   "this_mount",
		Read: true,
	}.Apply(t)
	require.EqualError(t, err, "Some error")
	assert.Equal(t, "this_mount", d.Id())
	assert.Equal(t, "", d.Get("source"))
}

func TestResourceAwsS3MountDelete(t *testing.T) {
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
		Resource: ResourceAWSS3Mount(),
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
			"cluster_id":     "this_cluster",
			"mount_name":     "this_mount",
			"s3_bucket_name": testS3BucketName,
		},
		ID:     "this_mount",
		Delete: true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "this_mount", d.Id())
	assert.Equal(t, "", d.Get("source"))
}

func TestAwsAccS3Mount(t *testing.T) {
	client := common.NewClientFromEnvironment()
	instanceProfile := qa.GetEnvOrSkipTest(t, "TEST_EC2_INSTANCE_PROFILE")
	ctx := context.WithValue(context.Background(), common.Current, t.Name())
	instanceProfilesAPI := identity.NewInstanceProfilesAPI(ctx, client)
	instanceProfilesAPI.Synchronized(instanceProfile, func() bool {
		if err := instanceProfilesAPI.Create(instanceProfile); err != nil {
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
		}, AWSIamMount{
			S3BucketName: bucket,
		})
		return true
	})
}
