package storage_test

import (
	"context"
	"fmt"
	"testing"

	"golang.org/x/exp/maps"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var mountHcl = `
data "databricks_spark_version" "latest" {}

# Test cluster to create the mount using.
resource "databricks_cluster" "this" {
	cluster_name = "acc-test-mounts-{var.STICKY_RANDOM}"
	spark_version = data.databricks_spark_version.latest.id
	instance_pool_id = "{env.TEST_INSTANCE_POOL_ID}"
	num_workers = 1

	aws_attributes {
		instance_profile_arn = "{env.TEST_INSTANCE_PROFILE_ARN}"
	}
}

resource "databricks_mount" "my_mount" {
	name = "test-mount-{var.STICKY_RANDOM}"
	cluster_id = databricks_cluster.this.id

	s3 {
		bucket_name      = "{env.TEST_S3_BUCKET_NAME}"
	}
}`

func TestAccCreateDatabricksMount(t *testing.T) {
	acceptance.WorkspaceLevel(t,
		acceptance.Step{
			Template: mountHcl,
		})
}

func TestAccCreateDatabricksMountIsFineOnClusterRecreate(t *testing.T) {
	clusterId1 := ""
	clusterId2 := ""

	acceptance.WorkspaceLevel(t,
		// Step 1 creates the cluster and mount.
		acceptance.Step{
			Template: mountHcl,
			Check: func(s *terraform.State) error {
				resources := s.RootModule().Resources
				cluster := resources["databricks_cluster.this"]
				if cluster == nil {
					return fmt.Errorf("expected to find databricks_cluster.this in resources keys: %v", maps.Keys(resources))
				}
				clusterId1 = cluster.Primary.ID

				// Assert cluster id is not empty. This is later used to ensure
				// the cluster has been recreated.
				assert.NotEmpty(t, clusterId1)

				// Assert the mount points to the created cluster.
				mount := resources["databricks_mount.my_mount"]
				assert.Equal(t, clusterId1, mount.Primary.Attributes["cluster_id"])
				return nil
			},
		},
		// Step 2: Manually delete the cluster, and then reapply the config. The mount
		// will be recreated in this case.
		acceptance.Step{
			PreConfig: func() {
				w, err := databricks.NewWorkspaceClient(&databricks.Config{})
				require.NoError(t, err)
				err = w.Clusters.PermanentDeleteByClusterId(context.Background(), clusterId1)
				assert.NoError(t, err, "failed to delete the cluster, id: "+clusterId1)
			},
			Template: mountHcl,
			Check: func(s *terraform.State) error {
				resources := s.RootModule().Resources
				cluster := resources["databricks_cluster.this"]
				if cluster == nil {
					return fmt.Errorf("expected to find databricks_cluster.this in resources keys: %v", maps.Keys(resources))
				}
				clusterId2 = cluster.Primary.ID

				// Assert cluster was indeed recreated
				assert.NotEmpty(t, clusterId1)
				assert.NotEmpty(t, clusterId2)
				assert.NotEqual(t, clusterId1, clusterId2)

				// Assert the mount points to the newly recreated cluster.
				mount := resources["databricks_mount.my_mount"]
				assert.Equal(t, clusterId2, mount.Primary.Attributes["cluster_id"])
				return nil
			},
		},
	)
}
