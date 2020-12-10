package acceptance

import (
	"context"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/identity"
	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/databrickslabs/databricks-terraform/internal/qa"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"

	"os"
	"testing"
)

func TestAwsAccGroupInstanceProfileResource(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	// TODO: add instance profile lock around these
	config := qa.EnvironmentTemplate(t, `
	resource "databricks_instance_profile" "this" {
		instance_profile_arn = "{env.TEST_EC2_INSTANCE_PROFILE}"
	}
	resource "databricks_group" "this" {
		display_name = "tf-{var.RANDOM}"
	}
	resource "databricks_group_instance_profile" "this" {
		group_id = databricks_group.this.id
		instance_profile_id = databricks_instance_profile.this.id
	}`)

	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config:  config,
				Destroy: false,
			},
			{
				Config: config,
				PreConfig: func() {
					client := common.CommonEnvironmentClient()
					ctx := context.Background()
					role := qa.FirstKeyValue(t, config, "instance_profile_arn")
					err := identity.NewInstanceProfilesAPI(ctx, client).Delete(role)
					assert.NoError(t, err, err)
				},
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
				Destroy:            false,
			},
			{
				Config: config,
				PreConfig: func() {
					client := common.CommonEnvironmentClient()
					ctx := context.Background()
					groupID := qa.FirstKeyValue(t, config, "group_id")
					err := identity.NewGroupsAPI(ctx, client).Delete(groupID)
					assert.NoError(t, err, err)
				},
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
				Destroy:            false,
			},
		},
	})
}
