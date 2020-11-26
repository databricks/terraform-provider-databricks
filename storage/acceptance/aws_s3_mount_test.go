package acceptance

import (
	"context"
	"fmt"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/compute"
	"github.com/databrickslabs/databricks-terraform/identity"
	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	. "github.com/databrickslabs/databricks-terraform/storage"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"

	"github.com/databrickslabs/databricks-terraform/internal/qa"
)

func getRunningClusterWithInstanceProfile(t *testing.T, client *common.DatabricksClient) (compute.ClusterInfo, error) {
	clusterName := "TerraformIntegrationTestIAM"
	clustersAPI := compute.NewClustersAPI(client)
	instanceProfile := qa.GetEnvOrSkipTest(t, "TEST_EC2_INSTANCE_PROFILE")
	return clustersAPI.GetOrCreateRunningCluster(clusterName, compute.Cluster{
		NumWorkers:             1,
		ClusterName:            clusterName,
		SparkVersion:           compute.CommonRuntimeVersion(),
		InstancePoolID:         compute.CommonInstancePoolID(),
		AutoterminationMinutes: 10,
		AwsAttributes: &compute.AwsAttributes{
			InstanceProfileArn: instanceProfile,
		},
	})
}

func TestAwsAccS3IamMount_WithCluster(t *testing.T) {
	client := common.NewClientFromEnvironment()
	arn := qa.GetEnvOrSkipTest(t, "TEST_EC2_INSTANCE_PROFILE")
	ctx := context.Background()
	identity.NewInstanceProfilesAPI(ctx, client).Synchronized(arn, func() {
		config := qa.EnvironmentTemplate(t, `
		resource "databricks_instance_profile" "this" {
			instance_profile_arn = "{env.TEST_EC2_INSTANCE_PROFILE}"
			skip_validation      = false
		}
		resource "databricks_cluster" "this" {
			cluster_name = "ready-{var.RANDOM}"
			instance_pool_id = "{var.POOL}"
			spark_version    = "{var.SPARK}"
			  autotermination_minutes = 10
			num_workers = 1
			aws_attributes {
				instance_profile_arn = databricks_instance_profile.this.id
			}
		}
		resource "databricks_aws_s3_mount" "mount" {
			cluster_id     = databricks_cluster.this.id
			mount_name     = "{var.RANDOM}"
			s3_bucket_name = "{env.TEST_S3_BUCKET}"
		}`, map[string]string{
			"POOL":  compute.CommonInstancePoolID(),
			"SPARK": compute.CommonRuntimeVersion(),
		})
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
			},
		})
	})
}

func TestAwsAccS3IamMount_NoClusterGiven(t *testing.T) {
	client := common.NewClientFromEnvironment()
	arn := qa.GetEnvOrSkipTest(t, "TEST_EC2_INSTANCE_PROFILE")
	ctx := context.Background()
	identity.NewInstanceProfilesAPI(ctx, client).Synchronized(arn, func() {
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
						clusterInfo, err := getRunningClusterWithInstanceProfile(t, client)
						assert.NoError(t, err)

						mp := NewMountPoint(client,
							qa.FirstKeyValue(t, config, "mount_name"),
							clusterInfo.ClusterID)
						err = mp.Delete()
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
	})
}
