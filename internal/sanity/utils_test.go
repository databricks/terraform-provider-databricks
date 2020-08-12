package sanity

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/databrickslabs/databricks-terraform/access"
	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/compute"
	"github.com/databrickslabs/databricks-terraform/identity"
	"github.com/databrickslabs/databricks-terraform/storage"
	"github.com/databrickslabs/databricks-terraform/workspace"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
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

func TestAccMissingResourcesInWorkspace(t *testing.T) {
	cloudEnv, ok := os.LookupEnv("CLOUD_ENV")
	if !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' set")
	}
	randStringID := acctest.RandString(10)
	randIntID := 2000000 + acctest.RandIntRange(100000, 20000000)
	randomClusterPolicyID := fmt.Sprintf("400E9E9E9A%d", acctest.RandIntRange(100000, 999999))
	randomInstancePoolID := fmt.Sprintf("%v-%v-%s-pool-%s", acctest.RandIntRange(1000, 9999),
		acctest.RandIntRange(100000, 999999), acctest.RandString(6), acctest.RandString(8))
	client := common.CommonEnvironmentClient()
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
				_, err := workspace.NewNotebooksAPI(client).Read("/" + randStringID)
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
				_, err := identity.NewGroupsAPI(client).Read(randStringID)
				return err
			},
		},
		{
			Name: "Users",
			ReadFunc: func() error {
				_, err := identity.NewUsersAPI(client).Read(randStringID)
				return err
			},
		},
		{
			Name: "Cluster Policies",
			ReadFunc: func() error {
				_, err := compute.NewClusterPoliciesAPI(client).Get(randomClusterPolicyID)
				return err
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
