package storage

import (
	"errors"
	"strings"
	"testing"

	"github.com/databrickslabs/databricks-terraform/compute"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databrickslabs/databricks-terraform/internal/qa"
)

// Test interface compliance via compile time error
var _ Mount = (*AWSIamMount)(nil)

const testS3BucketName = "test-s3-bucket"
const testS3BucketPath = "s3a://" + testS3BucketName

func TestResourceAwsS3MountCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=this_cluster",
				Response: compute.ClusterInfo{
					State: compute.ClusterStateRunning,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=this_cluster",
				Response: compute.ClusterInfo{
					State: compute.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceAWSS3Mount(),
		CommandMock: func(commandStr string) (string, error) {
			trunc := internal.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			if strings.HasPrefix(trunc, "def safe_mount") {
				assert.Contains(t, trunc, testS3BucketPath) // bucketname
				assert.Contains(t, trunc, `{}`)             // empty brackets for empty config
			}
			assert.Contains(t, trunc, "/mnt/this_mount")
			return testS3BucketPath, nil
		},
		State: map[string]interface{}{
			"cluster_id":     "this_cluster",
			"mount_name":     "this_mount",
			"s3_bucket_name": testS3BucketName,
		},
		Create: true,
	}.Apply(t)
	require.NoError(t, err, err) // TODO: global search-replace for NoError
	assert.Equal(t, "this_mount", d.Id())
	assert.Equal(t, testS3BucketPath, d.Get("source"))
}

func TestResourceAwsS3MountCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=this_cluster",
				Response: compute.ClusterInfo{
					State: compute.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceAWSS3Mount(),
		CommandMock: func(commandStr string) (string, error) {
			return "", errors.New("Some error")
		},
		State: map[string]interface{}{
			"cluster_id":     "this_cluster",
			"mount_name":     "this_mount",
			"s3_bucket_name": testS3BucketName,
		},
		Create: true,
	}.Apply(t)
	require.EqualError(t, err, "Some error")
	assert.Equal(t, "this_mount", d.Id())
	assert.Equal(t, "", d.Get("source"))
}

func TestResourceAwsS3MountRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=this_cluster",
				Response: compute.ClusterInfo{
					State: compute.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceAWSS3Mount(),
		CommandMock: func(commandStr string) (string, error) {
			trunc := internal.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			assert.Contains(t, trunc, "dbutils.fs.mounts()")
			assert.Contains(t, trunc, `mount.mountPoint == "/mnt/this_mount"`)
			return testS3BucketPath, nil
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
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=this_cluster",
				Response: compute.ClusterInfo{
					State: compute.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceAWSS3Mount(),
		CommandMock: func(commandStr string) (string, error) {
			trunc := internal.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			return "", errors.New("Mount not found")
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
	assert.Equal(t, "", d.Id())
	assert.Equal(t, "", d.Get("source"))
}

func TestResourceAwsS3MountRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=this_cluster",
				Response: compute.ClusterInfo{
					State: compute.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceAWSS3Mount(),
		CommandMock: func(commandStr string) (string, error) {
			trunc := internal.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			return "", errors.New("Some error")
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
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=this_cluster",
				Response: compute.ClusterInfo{
					State: compute.ClusterStateRunning,
				},
			},
		},
		Resource: ResourceAWSS3Mount(),
		CommandMock: func(commandStr string) (string, error) {
			trunc := internal.TrimLeadingWhitespace(commandStr)
			t.Logf("Received command:\n%s", trunc)
			assert.Contains(t, trunc, "/mnt/this_mount")
			assert.Contains(t, trunc, "dbutils.fs.unmount(mount_point)")
			return "", nil
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

func TestAwsAccBlobMount(t *testing.T) {
	_, mp := mountPointThroughReusedCluster(t)
	bucket := qa.GetEnvOrSkipTest(t, "TEST_S3_BUCKET")
	testMounting(t, mp, AWSIamMount{
		S3BucketName: bucket,
	})
}
