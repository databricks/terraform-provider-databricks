package storage_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/databricks/terraform-provider-databricks/storage"
)

const (
	mountClusterID = "test-cluster"
	mountName      = "test-mount"
	mountSource    = "s3a://test-bucket"
)

func mountClusterFixture() qa.HTTPFixture {
	return qa.HTTPFixture{
		Method:       "GET",
		ReuseRequest: true,
		Resource:     "/api/2.0/clusters/get?cluster_id=" + mountClusterID,
		Response: clusters.ClusterInfo{
			ClusterID: mountClusterID,
			State:     clusters.ClusterStateRunning,
			AwsAttributes: &clusters.AwsAttributes{
				InstanceProfileArn: "arn:aws:iam::123456789012:instance-profile/test",
			},
		},
	}
}

func mountCommandMock(t *testing.T) common.CommandMock {
	t.Helper()
	return func(command string) common.CommandResults {
		assert.Contains(t, command, "/mnt/"+mountName)
		if strings.Contains(command, "safe_mount") {
			assert.Contains(t, command, mountSource)
		}
		return common.CommandResults{ResultType: "text", Data: mountSource}
	}
}

func TestCreateDatabricksMount(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    []qa.HTTPFixture{mountClusterFixture()},
		Resource:    storage.ResourceMount(),
		CommandMock: mountCommandMock(t),
		State: map[string]any{
			"cluster_id": mountClusterID,
			"name":       mountName,
			"s3": []any{map[string]any{
				"bucket_name": "test-bucket",
			}},
		},
		Create: true,
	}.ApplyAndExpectData(t, map[string]any{
		"id":         mountName,
		"cluster_id": mountClusterID,
		"source":     mountSource,
	})
}

func TestCreateDatabricksMountWithProviderConfig(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:            []qa.HTTPFixture{mountClusterFixture()},
		ProviderWorkspaceID: "12345",
		Resource:            storage.ResourceMount(),
		CommandMock:         mountCommandMock(t),
		State: map[string]any{
			"cluster_id": mountClusterID,
			"name":       mountName,
			"provider_config": []any{map[string]any{
				"workspace_id": "12345",
			}},
			"s3": []any{map[string]any{
				"bucket_name": "test-bucket",
			}},
		},
		Create: true,
	}.ApplyAndExpectData(t, map[string]any{
		"id":         mountName,
		"cluster_id": mountClusterID,
		"source":     mountSource,
	})
}

func TestDatabricksMountReplacedWithCluster(t *testing.T) {
	r := storage.ResourceMount().ToResource()
	require.Contains(t, r.Schema, "cluster_id")
	assert.True(t, r.Schema["cluster_id"].ForceNew)
}
