package acceptance

import (
	"context"
	"fmt"
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
			appId := s.RootModule().Resources["databricks_service_principal.a"].Primary.Attributes["application_id"]
			_, err = w.Workspace.GetByPath(ctx, fmt.Sprintf("/Users/%v", appId))
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
				force_delete_home_dir = true
			}`,
		Check: func(s *terraform.State) error {
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
			appId := s.RootModule().Resources["databricks_service_principal.a"].Primary.Attributes["application_id"]
			_, err = w.Workspace.GetByPath(ctx, fmt.Sprintf("/Users/%v", appId))
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
