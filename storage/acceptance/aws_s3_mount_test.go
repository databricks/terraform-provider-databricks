package acceptance

import (
	"context"
	"fmt"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/compute"
	"github.com/databrickslabs/terraform-provider-databricks/identity"
	"github.com/databrickslabs/terraform-provider-databricks/internal/acceptance"
	"github.com/databrickslabs/terraform-provider-databricks/qa"
	. "github.com/databrickslabs/terraform-provider-databricks/storage"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"
)

func getRunningClusterWithInstanceProfile(t *testing.T, client *common.DatabricksClient) (compute.ClusterInfo, error) {
	clusterName := "TerraformIntegrationTestIAM"
	ctx := context.Background()
	clustersAPI := compute.NewClustersAPI(ctx, client)
	instanceProfile := qa.GetEnvOrSkipTest(t, "TEST_EC2_INSTANCE_PROFILE")
	return clustersAPI.GetOrCreateRunningCluster(clusterName, compute.Cluster{
		NumWorkers:             1,
		ClusterName:            clusterName,
		SparkVersion:           clustersAPI.LatestSparkVersionOrDefault(compute.SparkVersionRequest{Latest: true, LongTermSupport: true}),
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
	ctx := context.WithValue(context.Background(), common.Current, t.Name())
	instanceProfilesAPI := identity.NewInstanceProfilesAPI(ctx, client)
	instanceProfilesAPI.Synchronized(arn, func() bool {
		if instanceProfilesAPI.IsRegistered(arn) {
			return false
		}
		config := qa.EnvironmentTemplate(t, `
		resource "databricks_instance_profile" "this" {
			instance_profile_arn = "{env.TEST_EC2_INSTANCE_PROFILE}"
		}
		data "databricks_spark_version" "latest" {
		}
		resource "databricks_cluster" "this" {
			cluster_name = "ready-{var.RANDOM}"
			spark_version = data.databricks_spark_version.latest.id
			instance_pool_id = "{var.COMMON_INSTANCE_POOL_ID}"
			autotermination_minutes = 5
			num_workers = 1
			aws_attributes {
				instance_profile_arn = databricks_instance_profile.this.id
			}
		}
		resource "databricks_aws_s3_mount" "mount" {
			cluster_id     = databricks_cluster.this.id
			mount_name     = "{var.RANDOM}"
			s3_bucket_name = "{env.TEST_S3_BUCKET}"
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
			},
		})
		return true
	})
}

func TestAwsAccS3IamMount_NoClusterGiven(t *testing.T) {
	client := common.NewClientFromEnvironment()
	arn := qa.GetEnvOrSkipTest(t, "TEST_EC2_INSTANCE_PROFILE")
	ctx := context.WithValue(context.Background(), common.Current, t.Name())
	instanceProfilesAPI := identity.NewInstanceProfilesAPI(ctx, client)
	instanceProfilesAPI.Synchronized(arn, func() bool {
		config := qa.EnvironmentTemplate(t, `
		resource "databricks_instance_profile" "this" {
			instance_profile_arn = "{env.TEST_EC2_INSTANCE_PROFILE}"
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

						ctx := context.Background()
						mp := NewMountPoint(client.CommandExecutor(ctx),
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
		return true
	})
}
