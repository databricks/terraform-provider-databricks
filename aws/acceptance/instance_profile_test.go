package acceptance

import (
	"context"

	"github.com/databrickslabs/terraform-provider-databricks/aws"
	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/internal/acceptance"
	"github.com/databrickslabs/terraform-provider-databricks/qa"

	"os"
	"testing"
)

func TestAwsAccGroupInstanceProfileResource(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	ctx := context.WithValue(context.Background(), common.Current, t.Name())
	client := common.CommonEnvironmentClient()
	arn := qa.GetEnvOrSkipTest(t, "TEST_EC2_INSTANCE_PROFILE")
	instanceProfilesAPI := aws.NewInstanceProfilesAPI(ctx, client)
	instanceProfilesAPI.Synchronized(arn, func() bool {
		if instanceProfilesAPI.IsRegistered(arn) {
			return false
		}
		acceptance.Test(t, []acceptance.Step{
			{
				Template: `
				resource "databricks_instance_profile" "this" {
					instance_profile_arn = "{env.TEST_EC2_INSTANCE_PROFILE}"
				}
				resource "databricks_group" "this" {
					display_name = "tf-{var.RANDOM}"
				}
				resource "databricks_group_instance_profile" "this" {
					group_id = databricks_group.this.id
					instance_profile_id = databricks_instance_profile.this.id
				}`,
			},
		})
		return true
	})
}
