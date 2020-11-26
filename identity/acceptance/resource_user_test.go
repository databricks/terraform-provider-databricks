package acceptance

import (
	"context"
	"os"
	"testing"

	. "github.com/databrickslabs/databricks-terraform/identity"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccUserResource(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	config := qa.EnvironmentTemplate(t, `
	resource "databricks_user" "first" {
		user_name = "eerste+{var.RANDOM}@example.com"
		display_name = "Eerste {var.RANDOM}"
	}

	resource "databricks_user" "second" {
		user_name = "tweede+{var.RANDOM}@example.com"
		display_name = "Tweede {var.RANDOM}"
		allow_cluster_create = true
	}

	resource "databricks_user" "third" {
		user_name = "derde+{var.RANDOM}@example.com"
		display_name = "Derde {var.RANDOM}"
		allow_instance_pool_create = true
	}
	
	resource "databricks_group" "first" {
		display_name = "Vierde {var.RANDOM}"
	}`)
	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("databricks_user.first", "allow_cluster_create", "false"),
					resource.TestCheckResourceAttr("databricks_user.first", "allow_instance_pool_create", "false"),
					resource.TestCheckResourceAttr("databricks_user.second", "allow_cluster_create", "true"),
					resource.TestCheckResourceAttr("databricks_user.second", "allow_instance_pool_create", "false"),
					resource.TestCheckResourceAttr("databricks_user.third", "allow_cluster_create", "false"),
					resource.TestCheckResourceAttr("databricks_user.third", "allow_instance_pool_create", "true"),
					func(s *terraform.State) error {
						r := s.RootModule().Resources
						client := common.CommonEnvironmentClient()
						ctx := context.Background()
						return NewGroupsAPI(ctx, client).Patch(r["databricks_group.first"].Primary.ID, []string{
							r["databricks_user.first"].Primary.ID,
							r["databricks_user.second"].Primary.ID,
						}, nil, GroupMembersPath)
					},
				),
			},
			{
				Config: config,
			},
		},
	})
}
