package mws

import (
	"context"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/databrickslabs/databricks-terraform/internal/sanity"
)

func TestMwsAccMissingResources(t *testing.T) {
	if cloudEnv, ok := os.LookupEnv("CLOUD_ENV"); !ok || cloudEnv != "MWS" {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV=MWS' is set.")
	}
	mwsAcctID := qa.GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID")
	if mwsAcctID == "" {
		t.Skip("Must have DATABRICKS_ACCOUNT_ID environment variable set.")
	}
	randStringID := acctest.RandString(10)
	randIntID := 2000000 + acctest.RandIntRange(100000, 20000000)

	ctx := context.Background()
	client := common.CommonEnvironmentClient()
	tests := []sanity.MissingResourceCheck{
		{
			Name: "Credential",
			ReadFunc: func() error {
				_, err := NewCredentialsAPI(client).Read(mwsAcctID, randStringID)
				return err
			},
		},
		{
			Name: "Network",
			ReadFunc: func() error {
				_, err := NewNetworksAPI(ctx, client).Read(mwsAcctID, randStringID)
				return err
			},
		},
		{
			Name: "Storage",
			ReadFunc: func() error {
				_, err := NewStorageConfigurationsAPI(client).Read(mwsAcctID, randStringID)
				return err
			},
		},
		{
			Name: "Customer Managed Key",
			ReadFunc: func() error {
				_, err := NewCustomerManagedKeysAPI(client).Read(mwsAcctID, randStringID)
				return err
			},
		},
		{
			Name: "Workspace",
			ReadFunc: func() error {
				_, err := NewWorkspacesAPI(ctx, client).Read(mwsAcctID, int64(randIntID))
				return err
			},
		},
	}
	sanity.MissingResourceChecks(tests).Verify(t)
}
