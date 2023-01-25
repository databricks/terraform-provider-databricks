package acceptance

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/workspace"
	"github.com/stretchr/testify/assert"
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
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			resource "databricks_service_principal" "abc" {
				application_id    = "abc"
				display_name = "abc"
				delete_home_dir = true
			}`,
		},
		{
			Callback: func(ctx context.Context, client *common.DatabricksClient, id string) error {
				_, err := workspace.NewNotebooksAPI(ctx, client).Read("/Users/abc")
				assert.NotEqual(t, err, nil)
				return nil
			},
		},
	})
}
func TestAccServicePrincipalHomeDeleteNotDeleted(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	t.Parallel()
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			resource "databricks_service_principal" "abc" {
				application_id    = "abc"
				display_name = "abc"
				delete_hom_dir = false 
			}`,
		},
		{
			Callback: func(ctx context.Context, client *common.DatabricksClient, id string) error {
				_, err := workspace.NewNotebooksAPI(ctx, client).Read("/Users/abc")
				assert.Equal(t, err, nil)
				return nil
			},
		},
	})
}
