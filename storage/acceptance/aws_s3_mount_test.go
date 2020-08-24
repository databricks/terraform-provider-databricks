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

func TestAwsAccS3IamMount_NoClusterGiven(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	config := qa.EnvironmentTemplate(t, `
	resource "databricks_instance_profile" "this" {
		instance_profile_arn = "{env.TEST_EC2_INSTANCE_PROFILE}"
		skip_validation      = false
	}
	resource "databricks_aws_s3_mount" "mount" {
		mount_name        = "{var.RANDOM}"
		s3_bucket_name    = "{env.TEST_S3_BUCKET}"
		instance_profile  = databricks_instance_profile.this.id
	}`)
	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: mountResourceCheck("databricks_aws_s3_mount.mount", 
					func(client *common.DatabricksClient, mp MountPoint) error {
					source, err := mp.Source()
					assert.NoError(t, err)
					assert.Equal(t, fmt.Sprintf("s3a://%s", 
						qa.FirstKeyValue(t, config, "s3_bucket_name")), source)
					return nil
				}),
			},
			{
				PreConfig: func() {
					client := compute.CommonEnvironmentClientWithRealCommandExecutor()
					clusterInfo := compute.NewTinyClusterInCommonPoolPossiblyReused()
					mp := NewMountPoint(client,
						qa.FirstKeyValue(t, config, "mount_name"),
						clusterInfo.ClusterID)
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