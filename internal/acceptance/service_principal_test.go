package acceptance

import (
	"context"
	"fmt"
	"strings"
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
	loadWorkspaceEnv(t)
	if !isAzure(t) {
		skipf(t)("Test only valid for Azure")
	}
	uuid := createUuid()
	template := `
	resource "databricks_service_principal" "a" {
		application_id = "` + uuid + `"
		force_delete_home_dir = true
	}`
	var spId string
	workspaceLevel(t, step{
		Template: template,
		Check: func(s *terraform.State) error {
			spId = s.RootModule().Resources["databricks_service_principal.a"].Primary.Attributes["application_id"]
			return nil
		},
	}, step{
		Template: template,
		Destroy:  true,
		Check: func(s *terraform.State) error {
			w, err := databricks.NewWorkspaceClient()
			if err != nil {
				return err
			}
			ctx := context.Background()
			_, err = w.Workspace.GetStatusByPath(ctx, fmt.Sprintf("/Users/%v", spId))
			if err != nil && !apierr.IsMissing(err) {
				return err
			}
			return nil
		},
	})
}

func TestAccServicePrinicpalHomeDeleteNotDeleted(t *testing.T) {
	loadWorkspaceEnv(t)
	if !isAzure(t) {
		skipf(t)("Test only valid for Azure")
	}
	uuid := createUuid()
	template := `
	resource "databricks_service_principal" "a" {
		application_id = "` + uuid + `"
		force_delete_home_dir = false 
	}`
	var appId string
	workspaceLevel(t, step{
		Template: template,
		Check: func(s *terraform.State) error {
			appId = s.RootModule().Resources["databricks_service_principal.a"].Primary.Attributes["application_id"]
			return provisionHomeFolder(context.Background(), s, "databricks_service_principal.a", appId)
		},
	}, step{
		Template: template,
		Destroy:  true,
		Check: func(s *terraform.State) error {
			w, err := databricks.NewWorkspaceClient()
			if err != nil {
				return err
			}
			ctx := context.Background()
			_, err = w.Workspace.GetStatusByPath(ctx, fmt.Sprintf("/Users/%v", appId))
			return err
		},
	})
}

func TestMwsAccServicePrincipalResourceOnAzure(t *testing.T) {
	GetEnvOrSkipTest(t, "ARM_CLIENT_ID")
	azureSpnRenamed := strings.ReplaceAll(azureSpn, `"SPN `, `"SPN Renamed `)
	accountLevel(t, step{
		Template: azureSpn,
	}, step{
		Template: azureSpnRenamed,
	})
}

func TestAccServicePrincipalResourceOnAzure(t *testing.T) {
	GetEnvOrSkipTest(t, "ARM_CLIENT_ID")
	azureSpnRenamed := strings.ReplaceAll(azureSpn, `"SPN `, `"SPN Renamed `)
	workspaceLevel(t, step{
		Template: azureSpn,
	}, step{
		Template: azureSpnRenamed,
	})
}

func TestMwsAccServicePrincipalResourceOnAws(t *testing.T) {
	GetEnvOrSkipTest(t, "TEST_ROOT_BUCKET")
	awsSpnRenamed := strings.ReplaceAll(awsSpn, `"SPN `, `"SPN Renamed `)
	accountLevel(t, step{
		Template: awsSpn,
	}, step{
		Template: awsSpnRenamed,
	})
}

func TestAccServicePrincipalResourceOnAws(t *testing.T) {
	GetEnvOrSkipTest(t, "TEST_EC2_INSTANCE_PROFILE")
	awsSpnRenamed := strings.ReplaceAll(awsSpn, `"SPN `, `"SPN Renamed `)
	workspaceLevel(t, step{
		Template: awsSpn,
	}, step{
		Template: awsSpnRenamed,
	})
}
