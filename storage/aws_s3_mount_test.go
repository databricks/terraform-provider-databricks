package storage

import (
	"strings"
	"testing"

	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/commands"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test interface compliance via compile time error
var _ Mount = (*AWSIamMount)(nil)

const (
	testS3BucketName = "test-s3-bucket"
	testS3BucketPath = "s3a://" + testS3BucketName
)

func TestResourceAwsS3MountCreate(t *testing.T) {
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
		Resource: ResourceAWSS3Mount(),
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
			"cluster_id":     "this_cluster",
			"mount_name":     "this_mount",
			"s3_bucket_name": testS3BucketName,
		},
		Create: true,
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "this_mount", d.Id())
	assert.Equal(t, testS3BucketPath, d.Get("source"))
}

func TestResourceAwsS3MountCreate_nothing_specified(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource: ResourceAWSS3Mount(),
		State: map[string]any{
			"mount_name":     "this_mount",
			"s3_bucket_name": testS3BucketName,
		},
		Create: true,
	}.Apply(t)
	require.EqualError(t, err, "either cluster_id or instance_profile must be specified")
}

func TestResourceAwsS3MountCreate_invalid_arn(t *testing.T) {
	_, err := qa.ResourceFixture{
		Resource: ResourceAWSS3Mount(),
		State: map[string]any{
			"mount_name":       "this_mount",
			"s3_bucket_name":   testS3BucketName,
			"instance_profile": "this_mount",
		},
		Create: true,
	}.Apply(t)
	require.EqualError(t, err, "mount via profile: invalid arn: this_mount")
}

func TestResourceAwsS3MountRead(t *testing.T) {
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
		Resource: ResourceAWSS3Mount(),
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
				Response: clusters.ClusterInfo{
					State: clusters.ClusterStateRunning,
					AwsAttributes: &clusters.AwsAttributes{
						InstanceProfileArn: "abc",
					},
				},
			},
		},
		Resource: ResourceAWSS3Mount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := commands.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			return common.CommandResults{
				ResultType: "error",
				Summary:    "Mount not found",
			}
		},
		State: map[string]any{
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
				Response: clusters.ClusterInfo{
					State: clusters.ClusterStateRunning,
					AwsAttributes: &clusters.AwsAttributes{
						InstanceProfileArn: "abc",
					},
				},
			},
		},
		Resource: ResourceAWSS3Mount(),
		CommandMock: func(commandStr string) common.CommandResults {
			trunc := commands.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			return common.CommandResults{
				ResultType: "error",
				Summary:    "Some error",
			}
		},
		State: map[string]any{
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
				Response: clusters.ClusterInfo{
					State: clusters.ClusterStateRunning,
					AwsAttributes: &clusters.AwsAttributes{
						InstanceProfileArn: "abc",
					},
				},
			},
		},
		Resource: ResourceAWSS3Mount(),
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
