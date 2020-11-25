package importer

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/databrickslabs/databricks-terraform/access"
	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/compute"
	"github.com/databrickslabs/databricks-terraform/identity"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
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
				Parents: []identity.GroupMember{
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
					{ID: "123", DisplayName: "test@test.com", UserName: "test@test.com",
						Groups: []identity.GroupMember{
							{Display: "admins", Value: "a", Ref: "Groups/a", Type: "direct"},
						},
					},
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
			Response: compute.ClusterList{
				Clusters: []compute.ClusterInfo{
					{ClusterID: "test", ClusterName: "test", SparkVersion: "spark-3.0"},
					{ClusterID: "running", ClusterName: "running", SparkVersion: "spark-3.0"},
				},
			},
		},
		// test for cluster 'test'
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=test",
			Response: compute.ClusterInfo{
				ClusterID: "test", ClusterName: "test", SparkVersion: "spark-3.0",
				State: "TERMINATED", LastActivityTime: 1000,
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/clusters/events",
			ExpectedRequest: compute.EventsRequest{
				ClusterID:  "test",
				Order:      "DESC",
				EventTypes: []compute.ClusterEventType{"PINNED", "UNPINNED"},
				Limit:      1,
			},
			Response: compute.EventDetails{},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/libraries/cluster-status?cluster_id=test",
			Response: compute.ClusterLibraryStatuses{
				ClusterID: "test",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/permissions/clusters/test",
			Response: access.ObjectACL{},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Me",
			Response: identity.ScimUser{},
		},
		// tests for cluster 'running'
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=running",
			Response: compute.ClusterInfo{
				ClusterID: "running", ClusterName: "running", SparkVersion: "spark-3.0",
				State: "RUNNING", LastActivityTime: time.Now().Unix() - 100,
			},
		},
		{
			Method:   "POST",
			Resource: "/api/2.0/clusters/events",
			ExpectedRequest: compute.EventsRequest{
				ClusterID:  "running",
				Order:      "DESC",
				EventTypes: []compute.ClusterEventType{"PINNED", "UNPINNED"},
				Limit:      1,
			},
			Response: compute.EventDetails{},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/libraries/cluster-status?cluster_id=running",
			Response: compute.ClusterLibraryStatuses{
				ClusterID: "running",
				LibraryStatuses: []compute.LibraryStatus{
					{Library: &compute.Library{
						Maven: &compute.Maven{
							Coordinates: "com.microsoft.azure:azure-eventhubs-spark_2.12:2.3.17",
							Repo:        "https://mvnrepository.com",
							Exclusions:  []string{"abc", "def"},
						},
					},
						Status: "RUNNING",
					},
					{Library: &compute.Library{
						Pypi: &compute.PyPi{
							Package: "plotly==4.8.2",
							Repo:    "https://mvnrepository.com",
						},
					},
						Status: "RUNNING",
					},
					{Library: &compute.Library{
						Cran: &compute.Cran{
							Package: "ggplot",
							Repo:    "https://mvnrepository.com",
						},
					},
						Status: "RUNNING",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/permissions/clusters/running",
			Response: access.ObjectACL{
				ObjectID: "/clusters/running", ObjectType: "cluster",
				AccessControlList: []access.AccessControl{
					{UserName: "test@test.com",
						AllPermissions: []access.Permission{
							{PermissionLevel: "CAN_MANAGE", Inherited: false},
						},
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Me",
			Response: identity.ScimUser{ID: "a", DisplayName: "test@test.com"},
		},
		// the rest of the lists
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
	assert.NoError(t, err)
}

func TestImportingWithError(t *testing.T) {
	err := Run("-directory", "/bin/sh", "-services", "groups,users")
	assert.EqualError(t, err, "The path /bin/sh is not a directory")

	err = Run("-directory", "/bin/abcd", "-services", "groups,users")
	assert.EqualError(t, err, "Can't create directory /bin/abcd")
}

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
