package acceptance

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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
	LoadWorkspaceEnv(t)
	if !IsAzure(t) {
		skipf(t)("Test only valid for Azure")
	}
	uuid := createUuid()
	template := `
	resource "databricks_service_principal" "a" {
		application_id = "` + uuid + `"
		force_delete_home_dir = true
	}`
	var spId string
	WorkspaceLevel(t, Step{
		Template: template,
		Check: func(s *terraform.State) error {
			spId = s.RootModule().Resources["databricks_service_principal.a"].Primary.Attributes["application_id"]
			return nil
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
			_, err = w.Workspace.GetStatusByPath(ctx, fmt.Sprintf("/Users/%v", spId))
			if err != nil && !apierr.IsMissing(err) {
				return err
			}
			return nil
		},
	})
}

func TestAccServicePrinicpalHomeDeleteNotDeleted(t *testing.T) {
	LoadWorkspaceEnv(t)
	if !IsAzure(t) {
		skipf(t)("Test only valid for Azure")
	}
	uuid := createUuid()
	template := `
	resource "databricks_service_principal" "a" {
		application_id = "` + uuid + `"
		force_delete_home_dir = false 
	}`
	var appId string
	WorkspaceLevel(t, Step{
		Template: template,
		Check: func(s *terraform.State) error {
			appId = s.RootModule().Resources["databricks_service_principal.a"].Primary.Attributes["application_id"]
			return provisionHomeFolder(context.Background(), s, "databricks_service_principal.a", appId)
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
			_, err = w.Workspace.GetStatusByPath(ctx, fmt.Sprintf("/Users/%v", appId))
			return err
		},
	})
}

func TestMwsAccServicePrincipalResourceOnAzure(t *testing.T) {
	GetEnvOrSkipTest(t, "ARM_CLIENT_ID")
	azureSpnRenamed := strings.ReplaceAll(azureSpn, `"SPN `, `"SPN Renamed `)
	AccountLevel(t, Step{
		Template: azureSpn,
	}, Step{
		Template: azureSpnRenamed,
	})
}

func TestAccServicePrincipalResourceOnAzure(t *testing.T) {
	GetEnvOrSkipTest(t, "ARM_CLIENT_ID")
	azureSpnRenamed := strings.ReplaceAll(azureSpn, `"SPN `, `"SPN Renamed `)
	WorkspaceLevel(t, Step{
		Template: azureSpn,
	}, Step{
		Template: azureSpnRenamed,
	})
}

func TestMwsAccServicePrincipalResourceOnAws(t *testing.T) {
	GetEnvOrSkipTest(t, "TEST_ROOT_BUCKET")
	awsSpnRenamed := strings.ReplaceAll(awsSpn, `"SPN `, `"SPN Renamed `)
	AccountLevel(t, Step{
		Template: awsSpn,
	}, Step{
		Template: awsSpnRenamed,
	})
}

func TestAccServicePrincipalResourceOnAws(t *testing.T) {
	GetEnvOrSkipTest(t, "TEST_EC2_INSTANCE_PROFILE")
	awsSpnRenamed := strings.ReplaceAll(awsSpn, `"SPN `, `"SPN Renamed `)
	WorkspaceLevel(t, Step{
		Template: awsSpn,
	}, Step{
		Template: awsSpnRenamed,
	})
}
