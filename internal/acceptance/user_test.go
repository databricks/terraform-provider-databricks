package acceptance

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/service/iam"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// https://github.com/databricks/terraform-provider-databricks/issues/1097
func TestAccForceUserImport(t *testing.T) {
	username := qa.RandomEmail()
	WorkspaceLevel(t, Step{
		Template: `data "databricks_current_user" "me" {}`,
		Check: func(s *terraform.State) error {
			w, err := databricks.NewWorkspaceClient()
			if err != nil {
				return err
			}
			ctx := context.Background()
			// cleanup of this user will be handled by terraform
			logger.Infof(ctx, "Creating conflicting user")
			_, err = w.Users.Create(ctx, iam.User{
				Active:     true,
				UserName:   username,
				ExternalId: qa.RandomName("ext-id"),
			})
			if err != nil {
				return err
			}
			return nil
		},
	}, Step{
		Template: `resource "databricks_user" "this" {
			user_name = "` + username + `"
			force     = true
		}`,
	})
}

func TestAccUserHomeDeleteHasNoEffectInAccount(t *testing.T) {
	username := qa.RandomEmail()
	AccountLevel(t, Step{
		Template: `
		resource "databricks_user" "first" {
			user_name = "` + username + `"
			force_delete_home_dir = true
		}`,
	}, Step{
		Template: `
		resource "databricks_user" "second" {
			user_name = "{var.RANDOM}@example.com"
		}`,
	})
}

func TestAccUserHomeDelete(t *testing.T) {
	username := qa.RandomEmail()
	template := `
	resource "databricks_user" "first" {
		user_name = "` + username + `"
		force_delete_home_dir = true
	}`
	WorkspaceLevel(t, Step{
		Template: template,
	}, Step{
		Template: template,
		Destroy:  true,
		Check: func(s *terraform.State) error {
			w, err := databricks.NewWorkspaceClient()
			if err != nil {
				return err
			}
			ctx := context.Background()
			_, err = w.Workspace.GetStatusByPath(ctx, fmt.Sprintf("/Users/%v", username))
			if err != nil {
				targetErr := fmt.Sprintf("Path (/Users/%v) doesn't exist", username)
				if strings.Contains(err.Error(), targetErr) {
					return nil
				}
				return err
			}
			return nil
		},
	})
}

func provisionHomeFolder(ctx context.Context, s *terraform.State, tfAttribute, username string) error {
	client, err := client.New(&config.Config{})
	if err != nil {
		return err
	}
	userId := s.Modules[0].Resources[tfAttribute].Primary.ID
	return client.Do(ctx, "PUT", fmt.Sprintf("/api/2.0/workspace/user/%s/homefolder", userId), nil, nil, map[string]any{
		"user": map[string]any{
			"user_id":  userId,
			"username": username,
		},
	}, nil)
}

func TestAccUserHomeDeleteNotDeleted(t *testing.T) {
	username := qa.RandomEmail()
	template := `
	resource "databricks_user" "a" {
		user_name = "` + username + `"
	}`
	WorkspaceLevel(t, Step{
		Template: template,
		Check: func(s *terraform.State) error {
			return provisionHomeFolder(context.Background(), s, "databricks_user.a", username)
		},
	}, Step{
		Template: template,
		Destroy:  true,
		Check: func(s *terraform.State) error {
			w, err := databricks.NewWorkspaceClient()
			if err != nil {
				return err
			}
			ctx := context.Background()
			_, err = w.Workspace.GetStatusByPath(ctx, fmt.Sprintf("/Users/%v", username))
			return err
		},
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
	}
	`
	WorkspaceLevel(t, Step{
		Template: differentUsers,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_user.first", "allow_cluster_create", "false"),
			resource.TestCheckResourceAttr("databricks_user.first", "allow_instance_pool_create", "false"),
			resource.TestCheckResourceAttr("databricks_user.second", "allow_cluster_create", "true"),
			resource.TestCheckResourceAttr("databricks_user.second", "allow_instance_pool_create", "false"),
			resource.TestCheckResourceAttr("databricks_user.third", "allow_cluster_create", "false"),
			resource.TestCheckResourceAttr("databricks_user.third", "allow_instance_pool_create", "true"),
		),
	}, Step{
		Template: differentUsers,
	})
}

func TestAccUserResourceCaseInsensitive(t *testing.T) {
	username := "CSTF-" + qa.RandomEmail()
	csUser := `resource "databricks_user" "first" {
		user_name = "` + username + `"
		}`
	WorkspaceLevel(t, Step{
		Template: csUser,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_user.first", "user_name", strings.ToLower(username)),
		),
	}, Step{
		Template: csUser,
	})
}
