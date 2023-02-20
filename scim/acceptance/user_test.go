package acceptance

import (
	"context"
	"os"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/scim"
	"github.com/stretchr/testify/assert"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestMwsAccForceUserImport(t *testing.T) {
	TestAccForceUserImport(t)
}

// https://github.com/databricks/terraform-provider-databricks/issues/1097
func TestAccForceUserImport(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	username := qa.RandomEmail()
	t.Setenv("TEST_USERNAME", username)
	ctx := context.Background()
	client := common.CommonEnvironmentClient()
	usersAPI := scim.NewUsersAPI(ctx, client)
	user, err := usersAPI.Create(scim.User{
		Active:     true,
		UserName:   username,
		ExternalID: qa.RandomName("ext-id"),
	})
	assert.NoError(t, err)
	defer usersAPI.Delete(user.ID)
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `resource "databricks_user" "this" {
				user_name = "{env.TEST_USERNAME}"
				force     = true
			}`,
		},
	})
}
func TestMwsAccUserResource(t *testing.T) {
	TestAccUserResource(t)
}
func TestAccUserResource(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	t.Parallel()
	config := acceptance.EnvironmentTemplate(t, `
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
				),
			},
			{
				Config: config,
			},
		},
	})
}
