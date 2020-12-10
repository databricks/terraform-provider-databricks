package sanity

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/databrickslabs/databricks-terraform/access"
	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/compute"
	"github.com/databrickslabs/databricks-terraform/identity"
	"github.com/databrickslabs/databricks-terraform/mws"
	"github.com/databrickslabs/databricks-terraform/storage"
	"github.com/databrickslabs/databricks-terraform/workspace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/stretchr/testify/assert"
)

func TestDummy(t *testing.T) {
	MissingResourceChecks([]MissingResourceCheck{
		{
			Name: "Dummy",
			ReadFunc: func() error {
				return common.APIError{
					StatusCode: 404,
				}
			},
		},
	}).Verify(t)
}

func TestAccMutiworkspaceUsedFromNormalMode(t *testing.T) {
	if cloudEnv, ok := os.LookupEnv("CLOUD_ENV"); !ok || cloudEnv == "MWS" {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' set")
	}
	client := common.CommonEnvironmentClient()
	checkCheck := func(_ interface{}, err error) {
		assert.Error(t, err)
		a, ok := err.(common.APIError)
		assert.True(t, ok)
		assert.Equal(t, "INCORRECT_CONFIGURATION", a.ErrorCode)
	}
	checkCheck(mws.NewCredentialsAPI(client).List("_"))
	checkCheck(mws.NewNetworksAPI(client).List("_"))
	checkCheck(mws.NewStorageConfigurationsAPI(client).List("_"))
	checkCheck(mws.NewWorkspacesAPI(client).List("_"))
}

func TestAccMissingResourcesInWorkspace(t *testing.T) {
	cloudEnv, ok := os.LookupEnv("CLOUD_ENV")
	if !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' set")
	}
	randStringID := acctest.RandStringFromCharSet(16, "0123456789abcdef")
	randIntID := 2000000 + acctest.RandIntRange(100000, 20000000)
	randomInstancePoolID := fmt.Sprintf("%v-%v-%s-pool-%s", acctest.RandIntRange(1000, 9999),
		acctest.RandIntRange(100000, 999999), acctest.RandString(6), acctest.RandString(8))
	client := common.CommonEnvironmentClient()

	ctx := context.Background()
	tests := []MissingResourceCheck{
		{
			Name: "Tokens",
			ReadFunc: func() error {
				_, err := identity.NewTokensAPI(client).Read(randStringID)
				return err
			},
		},
		{
			Name: "Secret Scopes",
			ReadFunc: func() error {
				_, err := access.NewSecretScopesAPI(client).Read(randStringID)
				return err
			},
		},
		{
			Name: "Secrets",
			ReadFunc: func() error {
				_, err := access.NewSecretsAPI(client).Read(randStringID, randStringID)
				return err
			},
		},
		{
			Name: "Secret ACLs",
			ReadFunc: func() error {
				_, err := access.NewSecretAclsAPI(client).Read(randStringID, randStringID)
				return err
			},
		},
		{
			Name: "Notebooks",
			ReadFunc: func() error {
				_, err := workspace.NewNotebooksAPI(ctx, client).Read("/" + randStringID)
				return err
			},
		},
		{
			Name: "Instance Pools",
			ReadFunc: func() error {
				_, err := compute.NewInstancePoolsAPI(client).Read(randomInstancePoolID)
				return err
			},
		},
		{
			Name: "Clusters",
			ReadFunc: func() error {
				_, err := compute.NewClustersAPI(client).Get(randStringID)
				return err
			},
		},
		{
			Name: "DBFS Files",
			ReadFunc: func() error {
				_, err := storage.NewDBFSAPI(client).Read("/" + randStringID)
				return err
			},
		},
		{
			Name: "Groups",
			ReadFunc: func() error {
				_, err := identity.NewGroupsAPI(ctx, client).Read(randStringID)
				return err
			},
		},
		{
			Name: "Users",
			ReadFunc: func() error {
				_, err := identity.NewUsersAPI(ctx, client).Read(randStringID)
				return err
			},
		},
		{
			Name: "Cluster Policies",
			ReadFunc: func() error {
				_, err := compute.NewClusterPoliciesAPI(client).Get(randStringID)
				return err
			},
		},
		{
			Name: "Cluster Policies Delete",
			ReadFunc: func() error {
				return compute.NewClusterPoliciesAPI(client).Delete(randStringID)
			},
		},
		{
			Name: "Jobs",
			ReadFunc: func() error {
				_, err := compute.NewJobsAPI(client).Read(strconv.Itoa(randIntID))
				return err
			},
		},
	}
	if cloudEnv == "AWS" {
		tests = append(tests, MissingResourceCheck{
			Name: "Instance Profiles",
			ReadFunc: func() error {
				_, err := identity.NewInstanceProfilesAPI(client).Read(randStringID)
				return err
			},
		})
	}
	MissingResourceChecks(tests).Verify(t)
}
