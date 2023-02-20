package acceptance

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go"
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
	appId := "12345a67-8b9c-0d1e-23fa-4567b89cde04"
	appId2 := "22345a67-8b9c-0d1e-23fa-4567b89cde04"
	os.Setenv("appId", appId)
	os.Setenv("appId2", appId2)
	workspaceLevel(t, step{
		Template: `
			resource "databricks_service_principal" "a" {
				application_id = "{env.appId}"
				force_delete_home_dir = true
			}`,
		Check: func(s *terraform.State) error {
			return nil
		},
	}, step{
		Template: `
			resource "databricks_service_principal" "b" {
				application_id = "{env.appId2}"
			}
			`,
		Check: func(s *terraform.State) error {
			w, err := databricks.NewWorkspaceClient()
			if err != nil {
				return err
			}
			ctx := context.Background()
			_, err = w.Workspace.GetByPath(ctx, fmt.Sprintf("/Users/%v", appId))
			if err != nil {
				if strings.Contains(err.Error(), "doesn't exist") {
					return nil
				}
				return err
			}
			return nil
		},
	})
}

func TestAccServicePrinicpalHomeDeleteNotDeleted(t *testing.T) {
	appId := "12343a67-8b9c-0d1e-23fa-4567b89cde99"
	appId2 := "22343a67-8b9c-0d1e-23fa-4567b89cde99"
	os.Setenv("appId", appId)
	os.Setenv("appId2", appId2)
	workspaceLevel(t, step{
		Template: `
			resource "databricks_service_principal" "a" {
				application_id = "{env.appId}"
				force_delete_home_dir = true
			}`,
		Check: func(s *terraform.State) error {
			return nil
		},
	}, step{
		Template: `
			resource "databricks_service_principal" "b" {
				application_id = "{env.appId2}"
			}
			`,
		Check: func(s *terraform.State) error {
			w, err := databricks.NewWorkspaceClient()
			if err != nil {
				return err
			}
			ctx := context.Background()
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
