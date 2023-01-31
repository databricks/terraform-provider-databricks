package acceptance

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/workspace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestMwsAccServicePrincipalResourceOnAzure(t *testing.T) {
	TestAccServicePrincipalResourceOnAzure(t)
}
func TestAccServicePrincipalResourceOnAzure(t *testing.T) {
	if cloud, ok := os.LookupEnv("CLOUD_ENV"); !ok || !strings.Contains(cloud, "azure") {
		t.Skip("Test is only for CLOUD_ENV=azure")
	}
	t.Parallel()
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `resource "databricks_service_principal" "this" {
				application_id = "00000000-1234-5678-0000-000000000001"
				display_name = "SPN {var.RANDOM}"
				force = true
			}`,
		},
	})
}
func TestMwsAccServicePrincipalResourceOnAws(t *testing.T) {
	TestAccServicePrincipalResourceOnAws(t)
}
func TestAccServicePrincipalResourceOnAws(t *testing.T) {
	if cloud, ok := os.LookupEnv("CLOUD_ENV"); !ok || cloud != "aws" {
		t.Skip("Test is only for CLOUD_ENV=aws")
	}
	t.Parallel()
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `resource "databricks_service_principal" "this" {
				display_name = "SPN {var.RANDOM}"
			}`,
		},
	})
}

func TestAccServicePrincipalHomeDeleteSuccess(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	t.Parallel()
	appId := "12345a67-8b9c-0d1e-23fa-4567b89cde04"
	os.Setenv("appId", appId)
	ctx := context.Background()
	client := common.CommonEnvironmentClient()
	notebooksAPI := workspace.NewNotebooksAPI(ctx, client)
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			resource "databricks_service_principal" "a" {
				application_id = "{env.appId}"
				delete_home_dir = true
			}`,
			Check: func(s *terraform.State) error {
				return nil
			},
		},
		{
			Template: `
			resource "databricks_service_principal" "b" {
				application_id = "12345a67-8b9c-0d1e-23fa-4567b89cde10"
			}
			`,
			Check: func(s *terraform.State) error {
				_, err := notebooksAPI.Read(fmt.Sprintf("/Users/%v", appId))
				if err != nil {
					if strings.Contains(err.Error(), "doesn't exist") {
						return nil
					}
					return err
				}
				return nil
			},
		},
	})
}

func TestAccServicePrinicpalHomeDeleteNotDeleted(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	t.Parallel()
	appId := "12345a67-8b9c-0d1e-23fa-4567b89cde99"
	os.Setenv("appId", appId)
	ctx := context.Background()
	client := common.CommonEnvironmentClient()
	notebooksAPI := workspace.NewNotebooksAPI(ctx, client)
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			resource "databricks_service_principal" "a" {
				application_id = "{env.appId}"
			}`,
			Check: func(s *terraform.State) error {
				return nil
			},
		},
		{
			Template: `
			resource "databricks_service_principal" "b" {
				application_id = "12345a67-8b9c-0d1e-23fa-4567b89cde12"
			}
			`,
			Check: func(s *terraform.State) error {
				_, err := notebooksAPI.Read(fmt.Sprintf("/Users/%v", appId))
				return err
			},
		},
	})
}
