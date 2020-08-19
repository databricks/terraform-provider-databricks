package acceptance

import (
	"fmt"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/compute"
	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	. "github.com/databrickslabs/databricks-terraform/storage"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/stretchr/testify/assert"

	"github.com/databrickslabs/databricks-terraform/internal/qa"
)

func TestAWSS3IamMount_correctly_mounts(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	randomMountName := fmt.Sprintf("tf-mount-test-%s", qa.RandomName())
	expectedS3Bucket := os.Getenv("TEST_S3_BUCKET_NAME")

	config := qa.EnvironmentTemplate(t, `
	resource "databricks_aws_s3_mount" "mount" {
		cluster_id			 = "{env.TEST_AWS_MOUNT_CLUSTER_ID}"
		mount_name           = "{var.RANDOM_MOUNT_NAME}"
		s3_bucket_name       = "{env.TEST_S3_BUCKET_NAME}"
	}`, map[string]string{"RANDOM_MOUNT_NAME": randomMountName})

	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: acceptance.ResourceCheck("databricks_aws_s3_mount.mount", func(client *common.DatabricksClient, id string) error {
					clusterInfo, err := compute.NewClustersAPI(client).GetOrCreateRunningCluster("TerraformIntegrationTest")
					assert.NoError(t, err)
					mp := NewMountPoint(client, id, clusterInfo.ClusterID)
					source, err := mp.Source()
					assert.NoError(t, err)
					assert.Equal(t, fmt.Sprintf("s3a://%s", expectedS3Bucket), source)
					return nil
				}),
			},
			{
				PreConfig: func() {
					client := common.CommonEnvironmentClient()
					clusterInfo := compute.NewTinyClusterInCommonPoolPossiblyReused()
					mp := NewMountPoint(client, randomMountName, clusterInfo.ClusterID)
					err := mp.Delete()
					assert.NoError(t, err)
				},
				Config:             config,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				// Prior PreConfig deleted the mount so this one should attempt to recreate the mount
				Config: config,
			},
		},
	})
}
