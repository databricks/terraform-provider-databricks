package acceptance

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/service/scim"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// https://github.com/databricks/terraform-provider-databricks/issues/1097
func TestAccForceUserImport(t *testing.T) {
	username := qa.RandomEmail()
	workspaceLevel(t, step{
		Template: `data "databricks_current_user" "me" {}`,
		Check: func(s *terraform.State) error {
			w, err := databricks.NewWorkspaceClient()
			if err != nil {
				return err
			}
			ctx := context.Background()
			// cleanup of this user will be handled by terraform
			logger.Infof("Creating conflicting user")
			_, err = w.Users.Create(ctx, scim.User{
				Active:     true,
				UserName:   username,
				ExternalId: qa.RandomName("ext-id"),
			})
			if err != nil {
				return err
			}
			return nil
		},
	}, step{
		Template: `resource "databricks_user" "this" {
			user_name = "` + username + `"
			force     = true
		}`,
	})
}

func TestAccUserResource(t *testing.T) {
	differentUsers := `
	resource "databricks_user" "first" {
		user_name = "tf-eerste+{var.RANDOM}@example.com"
		display_name = "Eerste {var.RANDOM}"
	}

	resource "databricks_user" "second" {
		user_name = "tf-tweede+{var.RANDOM}@example.com"
		display_name = "Tweede {var.RANDOM}"
		allow_cluster_create = true
	}

	resource "databricks_user" "third" {
		user_name = "tf-derde+{var.RANDOM}@example.com"
		display_name = "Derde {var.RANDOM}"
		allow_instance_pool_create = true
	}`
	workspaceLevel(t, step{
		Template: differentUsers,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_user.first", "allow_cluster_create", "false"),
			resource.TestCheckResourceAttr("databricks_user.first", "allow_instance_pool_create", "false"),
			resource.TestCheckResourceAttr("databricks_user.second", "allow_cluster_create", "true"),
			resource.TestCheckResourceAttr("databricks_user.second", "allow_instance_pool_create", "false"),
			resource.TestCheckResourceAttr("databricks_user.third", "allow_cluster_create", "false"),
			resource.TestCheckResourceAttr("databricks_user.third", "allow_instance_pool_create", "true"),
		),
	}, step{
		Template: differentUsers,
	})
}
