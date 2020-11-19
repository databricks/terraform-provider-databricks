package importer

import (
	"log"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/access"
	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/compute"
	"github.com/databrickslabs/databricks-terraform/identity"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccImporter(t *testing.T) {
	if _, ok := os.LookupEnv("VSCODE_PID"); !ok {
		t.Skip("Running this test is only meant for IDE")
	}
	log.SetOutput(&levelWriter{"[INFO]", "[ERROR]", "[WARN]"})
	c := common.NewClientFromEnvironment()
	err := newImportContext(c).Run()
	assert.NoError(t, err)
}

func TestAccImportIdentity(t *testing.T) {
	if _, ok := os.LookupEnv("VSCODE_PID"); !ok {
		t.Skip("Running this test is only meant for IDE")
	}
	err := Run("-directory", "/tmp/data-group", "-services", "groups,users")
	assert.NoError(t, err)
}

func TestAccImportSecrets(t *testing.T) {
	if _, ok := os.LookupEnv("VSCODE_PID"); !ok {
		t.Skip("Running this test is only meant for IDE")
	}
	err := Run("-directory", "/tmp/data-group", "-services", "secrets,users,groups")
	assert.NoError(t, err)
}

func TestAccImportJobs(t *testing.T) {
	if _, ok := os.LookupEnv("VSCODE_PID"); !ok {
		t.Skip("Running this test is only meant for IDE")
	}
	err := Run(
		"-directory=/tmp/data-group",
		"-match=dbxtest7-sample",
		"-listing=jobs",
		"-last-active-days=30")
	assert.NoError(t, err)
}

func TestImporting(t *testing.T) {
	defer common.CleanupEnvironment()()
	_, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Groups?",
			Response: identity.GroupList{},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/jobs/list",
			Response: compute.JobList{},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/list",
			Response: compute.ClusterList{},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/secrets/scopes/list",
			Response: access.SecretScopeList{},
		},
	})
	require.NoError(t, err)
	defer server.Close()

	os.Setenv("DATABRICKS_HOST", server.URL)
	os.Setenv("DATABRICKS_TOKEN", "..")

	err = Run()
	assert.EqualError(t, err, "No resources to import")
}

func TestResourceName(t *testing.T) {
	ic := newImportContext(&common.DatabricksClient{})
	norm := ic.ResourceName(&resource{
		Name: "9721431b_bcd3_4526_b90f_f5de2befec8c-dbutils_extensions_2_11_0_0_1-18dc8.jar",
	}, &schema.ResourceData{})
	assert.Equal(t, "dbutils_extensions_jar", norm)

	norm = ic.ResourceName(&resource{
		Name: "9721431b_bcd3_4526_b90f_f5de2befec8c|8737798193",
	}, &schema.ResourceData{})
	assert.Equal(t, "r7322b058678", norm)

	norm = ic.ResourceName(&resource{
		Name: "General Policy - All Users",
	}, &schema.ResourceData{})
	assert.Equal(t, "general_policy_all_users", norm)
}
