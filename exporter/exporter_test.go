package exporter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/access"
	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/compute"
	"github.com/databrickslabs/terraform-provider-databricks/identity"
	"github.com/databrickslabs/terraform-provider-databricks/qa"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccImporter(t *testing.T) {
	if _, ok := os.LookupEnv("VSCODE_PID"); !ok {
		t.Skip("Running this test is only meant for IDE")
	}
	log.SetOutput(&levelWriter{"[INFO]", "[ERROR]", "[WARN]", "[DEBUG]"})
	err := Run()
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

// nolint
func getJSONObject(filename string) interface{} {
	data, _ := ioutil.ReadFile(filename)
	var obj map[string]interface{}
	err := json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Printf("[ERROR] error! err=%v\n", err)
		fmt.Printf("[ERROR] data=%s\n", string(data))
	}
	return obj
}

// TODO: don't craft the answers manually, especially complex ones, but instead read them from the JSON files?
// TODO: add test for outputing import.sh into directory without permissions, like `/bin`

func TestImportingUsersGroupsSecretScopes(t *testing.T) {
	defer common.CleanupEnvironment()()
	_, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Groups?",
			Response: identity.GroupList{
				Resources: []identity.ScimGroup{
					// TODO: add another user for which there is no filter resut
					{ID: "a", DisplayName: "admins",
						Members: []identity.GroupMember{
							{Display: "test@test.com", Value: "123", Ref: "Users/123"},
							{Display: "Test group", Value: "f", Ref: "Groups/f"},
						},
					},
					{ID: "b", DisplayName: "users"},
					{ID: "c", DisplayName: "test"},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Groups/a",
			Response: identity.ScimGroup{ID: "a", DisplayName: "admins",
				Members: []identity.GroupMember{
					{Display: "test@test.com", Value: "123", Ref: "Users/123"},
					{Display: "Test group", Value: "f", Ref: "Groups/f"},
				},
			},
			ReuseRequest: true,
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Groups/b",
			Response: identity.ScimGroup{ID: "b", DisplayName: "users"},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Groups/c",
			Response: identity.ScimGroup{ID: "c", DisplayName: "test",
				Groups: []identity.GroupMember{
					{Display: "admins", Value: "a", Ref: "Groups/a", Type: "direct"},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Groups/f",
			Response: identity.ScimGroup{ID: "f", DisplayName: "nested"},
		},
		// TODO: add groups to the output
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Users/123",
			Response: identity.ScimUser{ID: "123", DisplayName: "test@test.com", UserName: "test@test.com"},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Users?filter=userName%20eq%20%27test@test.com%27",
			Response: identity.UserList{
				Resources: []identity.ScimUser{
					{ID: "123", DisplayName: "test@test.com", UserName: "test@test.com"},
				},
				StartIndex:   1,
				TotalResults: 1,
				ItemsPerPage: 1,
			},
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
			Method:       "GET",
			Resource:     "/api/2.0/secrets/scopes/list",
			ReuseRequest: true,
			Response: access.SecretScopeList{
				Scopes: []access.SecretScope{
					{Name: "a"},
				},
			},
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/secrets/list?scope=a",
			ReuseRequest: true,
			Response: access.SecretsList{
				Secrets: []access.SecretMetadata{
					{Key: "b"},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/secrets/acls/list?scope=a",
			Response: access.SecretScopeACL{
				Items: []access.ACLItem{
					{Permission: "MANAGE", Principal: "test"},
					{Permission: "READ", Principal: "users"},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/secrets/acls/list?scope=a",
			Response: access.SecretScopeACL{
				Items: []access.ACLItem{
					{Permission: "MANAGE", Principal: "test"},
					{Permission: "READ", Principal: "users"},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/secrets/acls/get?principal=test&scope=a",
			Response: access.ACLItem{Permission: "MANAGE", Principal: "test"},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/secrets/acls/get?principal=users&scope=a",
			Response: access.ACLItem{Permission: "READ", Principal: "users"},
		},
	})
	require.NoError(t, err)
	defer server.Close()

	os.Setenv("DATABRICKS_HOST", server.URL)
	os.Setenv("DATABRICKS_TOKEN", "..")

	tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
	defer os.RemoveAll(tmpDir)

	err = Run("-directory", tmpDir)
	assert.NoError(t, err)
	// TODO: check the content of the generated files
}

func TestImportingNoResourcesError(t *testing.T) {
	defer common.CleanupEnvironment()()
	_, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Groups?",
			Response: identity.GroupList{Resources: []identity.ScimGroup{}},
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
			Method:       "GET",
			Resource:     "/api/2.0/secrets/scopes/list",
			ReuseRequest: true,
			Response: access.SecretScopeList{
				Scopes: []access.SecretScope{},
			},
		},
	})
	require.NoError(t, err)
	defer server.Close()

	os.Setenv("DATABRICKS_HOST", server.URL)
	os.Setenv("DATABRICKS_TOKEN", "..")

	tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
	defer os.RemoveAll(tmpDir)

	err = Run("-directory", tmpDir)
	assert.EqualError(t, err, "No resources to import")
}

func TestImportingClusters(t *testing.T) {
	defer common.CleanupEnvironment()()
	_, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Groups?",
			Response: identity.GroupList{Resources: []identity.ScimGroup{}},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/jobs/list",
			Response: compute.JobList{},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/list",
			Response: getJSONObject("test-data/clusters-list-response.json"),
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=test1",
			Response: getJSONObject("test-data/get-cluster-test1-response.json"),
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/clusters/events",
			ExpectedRequest: compute.EventsRequest{
				ClusterID:  "test1",
				Order:      "DESC",
				EventTypes: []compute.ClusterEventType{"PINNED", "UNPINNED"},
				Limit:      1,
			},
			Response: compute.EventDetails{},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/libraries/cluster-status?cluster_id=test1",
			Response: getJSONObject("test-data/libraries-cluster-status-test1.json"),
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/permissions/clusters/test1",
			Response: getJSONObject("test-data/get-cluster-permissions-test1-response.json"),
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/dbfs/get-status?path=dbfs:%2FFileStore%2Fjars%2Ftest.jar",
			ReuseRequest: true,
			Response:     getJSONObject("test-data/get-dbfs-library-status.json"),
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/dbfs/read?length=1000000&path=dbfs%3A%2FFileStore%2Fjars%2Ftest.jar",
			ReuseRequest: true,
			Response:     getJSONObject("test-data/get-dbfs-library-data.json"),
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=test2",
			Response: getJSONObject("test-data/get-cluster-test2-response.json"),
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/clusters/events",
			ExpectedRequest: compute.EventsRequest{
				ClusterID:  "test2",
				Order:      "DESC",
				EventTypes: []compute.ClusterEventType{"PINNED", "UNPINNED"},
				Limit:      1,
			},
			Response: compute.EventDetails{},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/libraries/cluster-status?cluster_id=test2",
			Response: getJSONObject("test-data/libraries-cluster-status-test2.json"),
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/permissions/clusters/test2",
			Response: getJSONObject("test-data/get-cluster-permissions-test2-response.json"),
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/policies/clusters/get?policy_id=123",
			Response: getJSONObject("test-data/get-cluster-policy.json"),
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/permissions/cluster-policies/123",
			Response: getJSONObject("test-data/get-cluster-policy-permissions.json"),
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=awscluster",
			Response: getJSONObject("test-data/get-cluster-awscluster-response.json"),
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/clusters/events",
			ExpectedRequest: compute.EventsRequest{
				ClusterID:  "awscluster",
				Order:      "DESC",
				EventTypes: []compute.ClusterEventType{"PINNED", "UNPINNED"},
				Limit:      1,
			},
			Response: compute.EventDetails{},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/libraries/cluster-status?cluster_id=awscluster",
			Response: getJSONObject("test-data/libraries-cluster-status-test2.json"),
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/permissions/clusters/awscluster",
			Response: getJSONObject("test-data/get-cluster-permissions-awscluster-response.json"),
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/instance-profiles/list",
			Response: getJSONObject("test-data/list-instance-profiles.json"),
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/preview/scim/v2/Me",
			ReuseRequest: true,
			Response:     identity.ScimUser{ID: "a", DisplayName: "test@test.com"},
		},
	})
	require.NoError(t, err)
	defer server.Close()

	os.Setenv("DATABRICKS_HOST", server.URL)
	os.Setenv("DATABRICKS_TOKEN", "..")

	tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
	defer os.RemoveAll(tmpDir)

	err = Run("-directory", tmpDir, "-listing", "compute")
	assert.NoError(t, err)
}

// TODO: add spark_submit_task job
func TestImportingJobs(t *testing.T) {
	defer common.CleanupEnvironment()()
	_, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Groups?",
			Response: identity.GroupList{Resources: []identity.ScimGroup{}},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/jobs/list",
			Response: getJSONObject("test-data/get-jobs-list.json"),
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/jobs/get?job_id=14",
			Response: getJSONObject("test-data/get-job-14.json"),
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/permissions/jobs/14",
			Response: getJSONObject("test-data/get-job-permissions-14.json"),
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/jobs/get?job_id=12",
			Response: getJSONObject("test-data/get-job-12.json"),
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/permissions/jobs/12",
			Response: getJSONObject("test-data/get-job-permissions-12.json"),
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=test2",
			Response: getJSONObject("test-data/get-cluster-test2-response.json"),
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/clusters/events",
			ExpectedRequest: compute.EventsRequest{
				ClusterID:  "test2",
				Order:      "DESC",
				EventTypes: []compute.ClusterEventType{"PINNED", "UNPINNED"},
				Limit:      1,
			},
			Response: compute.EventDetails{},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/libraries/cluster-status?cluster_id=test2",
			Response: getJSONObject("test-data/libraries-cluster-status-test2.json"),
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/permissions/clusters/test2",
			Response: getJSONObject("test-data/get-cluster-permissions-test2-response.json"),
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/policies/clusters/get?policy_id=123",
			Response: getJSONObject("test-data/get-cluster-policy.json"),
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/permissions/cluster-policies/123",
			Response: getJSONObject("test-data/get-cluster-policy-permissions.json"),
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/jobs/get?job_id=15",
			Response: getJSONObject("test-data/get-job-15.json"),
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/permissions/jobs/15",
			Response: getJSONObject("test-data/get-job-permissions-15.json"),
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/dbfs/get-status?path=dbfs:%2FFileStore%2Fjars%2Ftest.jar",
			ReuseRequest: true,
			Response:     getJSONObject("test-data/get-dbfs-library-status.json"),
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/dbfs/read?length=1000000&path=dbfs%3A%2FFileStore%2Fjars%2Ftest.jar",
			ReuseRequest: true,
			Response:     getJSONObject("test-data/get-dbfs-library-data.json"),
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/instance-pools/get?instance_pool_id=pool1",
			ReuseRequest: true,
			Response:     getJSONObject("test-data/get-instance-pool1.json"),
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/preview/permissions/instance-pools/pool1",
			ReuseRequest: true,
			Response:     getJSONObject("test-data/get-instance-pool1-permissions.json"),
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/jobs/get?job_id=16",
			Response: getJSONObject("test-data/get-job-16.json"),
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/permissions/jobs/16",
			Response: getJSONObject("test-data/get-job-permissions-16.json"),
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/jobs/get?job_id=17",
			Response: getJSONObject("test-data/get-job-17.json"),
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/permissions/jobs/17",
			Response: getJSONObject("test-data/get-job-permissions-17.json"),
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/preview/scim/v2/Me",
			ReuseRequest: true,
			Response:     identity.ScimUser{ID: "a", DisplayName: "test@test.com"},
		},
	})
	require.NoError(t, err)
	defer server.Close()

	os.Setenv("DATABRICKS_HOST", server.URL)
	os.Setenv("DATABRICKS_TOKEN", "..")

	tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
	defer os.RemoveAll(tmpDir)

	err = Run("-directory", tmpDir, "-listing", "jobs")
	assert.NoError(t, err)
}

func TestImportingWithError(t *testing.T) {
	err := Run("-directory", "/bin/sh", "-services", "groups,users")
	assert.EqualError(t, err, "The path /bin/sh is not a directory")

	err = Run("-directory", "/bin/abcd", "-services", "groups,users")
	assert.EqualError(t, err, "Can't create directory /bin/abcd")
}

func TestImportingSecrets(t *testing.T) {
	defer common.CleanupEnvironment()()
	_, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Groups?",
			Response: identity.GroupList{Resources: []identity.ScimGroup{}},
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
			Method:       "GET",
			Resource:     "/api/2.0/secrets/scopes/list",
			ReuseRequest: true,
			Response:     getJSONObject("test-data/secret-scopes-response.json"),
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/secrets/list?scope=some-kv-scope",
			ReuseRequest: true,
			Response:     getJSONObject("test-data/secret-scopes-list-scope-response.json"),
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/secrets/acls/list?scope=some-kv-scope",
			ReuseRequest: true,
			Response:     getJSONObject("test-data/secret-scopes-list-scope-acls-response.json"),
		},
		{
			Method:       "GET",
			Resource:     "/api/2.0/secrets/acls/get?principal=test%40test.com&scope=some-kv-scope",
			ReuseRequest: true,
			Response:     getJSONObject("test-data/secret-scopes-get-principal-response.json"),
		},
	})
	require.NoError(t, err)
	defer server.Close()

	os.Setenv("DATABRICKS_HOST", server.URL)
	os.Setenv("DATABRICKS_TOKEN", "..")

	tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
	defer os.RemoveAll(tmpDir)

	err = Run("-directory", tmpDir, "-listing", "secrets", "-generateProviderDeclaration", "true")
	assert.NoError(t, err)
}

// func TestImportingMounts(t *testing.T) {
// 	defer common.CleanupEnvironment()()
// 	_, server, err := qa.HttpFixtureClient(t, []qa.HTTPFixture{
// 		{
// 			Method:   "GET",
// 			Resource: "/api/2.0/clusters/list",
// 			Response: compute.ClusterList{},
// 		},
// 		{
// 			Method:   "GET",
// 			Resource: "/api/2.0/clusters/list-node-types",
// 			Response: compute.NodeTypeList{},
// 		},
// 	})
// 	require.NoError(t, err)
// 	defer server.Close()

// 	os.Setenv("DATABRICKS_HOST", server.URL)
// 	os.Setenv("DATABRICKS_TOKEN", "..")

// 	tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
// 	defer os.RemoveAll(tmpDir)

// 	err = Run("-directory", tmpDir, "-listing", "mounts", "-mounts", "true")
// 	assert.NoError(t, err)
// }

func TestResourceName(t *testing.T) {
	ic := newImportContext(&common.DatabricksClient{})
	norm := ic.ResourceName(&resource{
		Name: "9721431b_bcd3_4526_b90f_f5de2befec8c-dbutils_extensions_2_11_0_0_1-18dc8.jar",
	})
	assert.Equal(t, "dbutils_extensions_jar", norm)

	norm = ic.ResourceName(&resource{
		Name: "9721431b_bcd3_4526_b90f_f5de2befec8c|8737798193",
	})
	assert.Equal(t, "r7322b058678", norm)

	norm = ic.ResourceName(&resource{
		Name: "General Policy - All Users",
	})
	assert.Equal(t, "general_policy_all_users", norm)
}
