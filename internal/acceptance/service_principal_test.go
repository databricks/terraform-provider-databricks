package acceptance

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const azureSpn = `resource "databricks_service_principal" "this" {
	application_id = "{var.RANDOM_UUID}"
	display_name = "SPN {var.RANDOM}"
	force = true
}`

const awsSpn = `resource "databricks_service_principal" "this" {
	display_name = "SPN {var.RANDOM}"
}`

func TestAccServicePrincipalHomeDeleteSuccess(t *testing.T) {
	GetEnvOrSkipTest(t, "ARM_CLIENT_ID")
	workspaceLevel(t, step{
		Template: `
			resource "databricks_service_principal" "a" {
				application_id = "{var.RANDOM_UUID}"
				force_delete_home_dir = true
			}`,
		Check: func(s *terraform.State) error {
			appId := s.RootModule().Resources["databricks_service_principal.a"].Primary.Attributes["application_id"]
			os.Setenv("application_id_a", appId)
			return nil
		},
	}, step{
		Template: `
			resource "databricks_service_principal" "b" {
				application_id = "{var.RANDOM_UUID}"
			}
			`,
		Check: func(s *terraform.State) error {
			w, err := databricks.NewWorkspaceClient()
			if err != nil {
				return err
			}
			ctx := context.Background()
			_, err = w.Workspace.GetStatusByPath(ctx, fmt.Sprintf("/Users/%v", os.Getenv("application_id_a")))
			os.Remove("application_id_a")
			if err != nil {
				if apierr.IsMissing(err) {
					return nil
				}
				return err
			}
			return nil
		},
	})
}

func TestAccServicePrinicpalHomeDeleteNotDeleted(t *testing.T) {
	GetEnvOrSkipTest(t, "ARM_CLIENT_ID")
	workspaceLevel(t, step{
		Template: `
			resource "databricks_service_principal" "a" {
				application_id = "{var.RANDOM_UUID}"
				force_delete_home_dir = false 
			}`,
		Check: func(s *terraform.State) error {
			appId := s.RootModule().Resources["databricks_service_principal.a"].Primary.Attributes["application_id"]
			os.Setenv("application_id_a", appId)
			return nil
		},
	}, step{
		Template: `
			resource "databricks_service_principal" "b" {
				application_id = "{var.RANDOM_UUID}"
			}
			`,
		Check: func(s *terraform.State) error {
			w, err := databricks.NewWorkspaceClient()
			if err != nil {
				return err
			}
			ctx := context.Background()
			appId := os.Getenv("application_id_a")
			_, err = w.Workspace.GetStatusByPath(ctx, fmt.Sprintf("/Users/%v", appId))
			os.Remove("application_id_a")
			return err
		},
	})
}

func TestMwsAccServicePrincipalResourceOnAzure(t *testing.T) {
	GetEnvOrSkipTest(t, "ARM_CLIENT_ID")
	accountLevel(t, step{
		Template: azureSpn,
	})
}

func TestAccServicePrincipalResourceOnAzure(t *testing.T) {
	GetEnvOrSkipTest(t, "ARM_CLIENT_ID")
	workspaceLevel(t, step{
		Template: azureSpn,
	})
}

func TestMwsAccServicePrincipalResourceOnAws(t *testing.T) {
	GetEnvOrSkipTest(t, "TEST_ROOT_BUCKET")
	accountLevel(t, step{
		Template: awsSpn,
	})
}

func TestAccServicePrincipalResourceOnAws(t *testing.T) {
	GetEnvOrSkipTest(t, "TEST_EC2_INSTANCE_PROFILE")
	workspaceLevel(t, step{
		Template: awsSpn,
	})
}
