package exporter

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/commands"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/jobs"
	"github.com/databricks/terraform-provider-databricks/libraries"
	"github.com/databricks/terraform-provider-databricks/permissions"
	"github.com/databricks/terraform-provider-databricks/pipelines"
	"github.com/databricks/terraform-provider-databricks/policies"
	"github.com/databricks/terraform-provider-databricks/pools"
	"github.com/databricks/terraform-provider-databricks/provider"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/databricks/terraform-provider-databricks/repos"
	"github.com/databricks/terraform-provider-databricks/scim"
	"github.com/databricks/terraform-provider-databricks/secrets"
	"github.com/databricks/terraform-provider-databricks/storage"
	"github.com/databricks/terraform-provider-databricks/workspace"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
)

func importContextForTest() *importContext {
	p := provider.DatabricksProvider()
	return &importContext{
		Importables:              resourcesMap,
		Resources:                p.ResourcesMap,
		testEmits:                map[string]bool{},
		nameFixes:                nameFixes,
		waitGroup:                &sync.WaitGroup{},
		allUsers:                 map[string]scim.User{},
		allSps:                   map[string]scim.User{},
		channels:                 makeResourcesChannels(p),
		exportDeletedUsersAssets: false,
		ignoredResources:         map[string]struct{}{},
	}
}

func importContextForTestWithClient(ctx context.Context, client *common.DatabricksClient) *importContext {
	ic := importContextForTest()
	ic.Client = client
	ic.Context = ctx
	if client.Config.IsAccountClient() {
		ic.accountClient, _ = client.AccountClient()
	} else {
		ic.workspaceClient, _ = client.WorkspaceClient()
	}
	return ic
}

func TestInstancePool(t *testing.T) {
	d := pools.ResourceInstancePool().TestResourceData()
	d.Set("instance_pool_name", "blah-bah")
	ic := importContextForTest()

	name := resourcesMap["databricks_instance_pool"].Name(ic, d)
	assert.Equal(t, "blah-bah", name)

	d = pools.ResourceInstancePool().TestResourceData()
	d.SetId("abc-bcd-def")
	name = resourcesMap["databricks_instance_pool"].Name(ic, d)
	assert.Equal(t, "def", name)

	ic.meAdmin = true
	err := resourcesMap["databricks_instance_pool"].Import(ic, &resource{
		ID:   "abc",
		Data: d,
	})
	assert.NoError(t, err)
	assert.True(t, ic.testEmits["databricks_permissions[inst_pool_def] (id: /instance-pools/abc)"])
}

func TestClusterPolicy(t *testing.T) {
	d := policies.ResourceClusterPolicy().TestResourceData()
	d.Set("name", "bcd")
	definition := map[string]map[string]string{
		"aws_attributes.instance_profile_arn": {
			"value": "def",
		},
		"instance_pool_id": {
			"defaultValue": "efg",
		},
		"init_scripts.0.dbfs.destination": {
			"type":  "fixed",
			"value": "dbfs:/FileStore/init-script.sh",
		},
	}
	policy, _ := json.Marshal(definition)
	d.Set("definition", string(policy))
	ic := importContextForTest()
	ic.meAdmin = true
	err := ic.Importables["databricks_cluster_policy"].Import(ic, &resource{
		ID:   "abc",
		Data: d,
	})
	assert.NoError(t, err)
	assert.Len(t, ic.testEmits, 4)
	assert.True(t, ic.testEmits["databricks_permissions[cluster_policy_bcd] (id: /cluster-policies/abc)"])
	assert.True(t, ic.testEmits["databricks_instance_pool[<unknown>] (id: efg)"])
	assert.True(t, ic.testEmits["databricks_instance_profile[<unknown>] (id: def)"])
	assert.True(t, ic.testEmits["databricks_dbfs_file[<unknown>] (id: dbfs:/FileStore/init-script.sh)"])
}

func TestPredefinedClusterPolicy(t *testing.T) {
	d := policies.ResourceClusterPolicy().TestResourceData()
	d.Set("policy_family_id", "job-cluster")
	policy, _ := json.Marshal(map[string]map[string]string{})
	d.Set("definition", string(policy))
	ic := importContextForTest()
	ic.builtInPolicies = map[string]compute.PolicyFamily{
		"job-cluster": {Name: "Job Compute"},
	}
	r := resource{ID: "abc", Data: d}
	err := ic.Importables["databricks_cluster_policy"].Import(ic, &r)
	assert.NoError(t, err)
	assert.Equal(t, "data", r.Mode)
	assert.Equal(t, "", r.Data.Get("definition").(string))
}

func TestGroup(t *testing.T) {
	ic := importContextForTest()
	ic.allGroups = []scim.Group{
		{
			DisplayName: "foo",
			ID:          "123",
			Roles: []scim.ComplexValue{
				{
					Value: "abc",
					Type:  "direct",
				},
			},
			Members: []scim.ComplexValue{
				// this is just for log line printing
				{Value: "a001"},
				{Value: "a002"},
				{Value: "a003"},
				{Value: "a004"},
				{Value: "a005"},
				{Value: "a006"},
				{Value: "a007"},
				{Value: "a008"},
				{Value: "a009"},
				{Value: "a010"},
				{Value: "a011"},
			},
			Groups: []scim.ComplexValue{
				{
					Value: "parent-group",
					Type:  "direct",
				},
			},
		},
	}
	d := scim.ResourceGroup().TestResourceData()
	d.Set("display_name", "foo")
	r := &resource{
		Value:     "foo",
		Attribute: "display_name",
		Data:      d,
	}
	err := ic.Importables["databricks_group"].Search(ic, r)
	assert.NoError(t, err)
	assert.Equal(t, "123", r.ID)

	err = ic.Importables["databricks_group"].Import(ic, r)
	assert.NoError(t, err)
	assert.Len(t, ic.testEmits, 4)
	assert.True(t, ic.testEmits["databricks_group_role[<unknown>] (id: 123|abc)"])
	assert.True(t, ic.testEmits["databricks_instance_profile[<unknown>] (id: abc)"])
	assert.True(t, ic.testEmits["databricks_group[<unknown>] (id: parent-group)"])
	assert.True(t, ic.testEmits["databricks_group_member[_parent-group_foo] (id: parent-group|123)"])
}

func TestPermissions(t *testing.T) {
	p := permissions.ResourcePermissions()
	d := p.TestResourceData()
	d.SetId("abc")
	ic := importContextForTest()
	name := ic.Importables["databricks_permissions"].Name(ic, d)
	assert.Equal(t, "abc", name)

	d.MarkNewResource()
	err := common.StructToData(permissions.PermissionsEntity{
		AccessControlList: []permissions.AccessControlChange{
			{
				UserName: "a",
			},
			{
				GroupName: "b",
			},
			{
				ServicePrincipalName: "123",
			},
		},
	}, p.Schema, d)
	assert.NoError(t, err)

	err = ic.Importables["databricks_permissions"].Import(ic, &resource{
		Data: d,
	})
	assert.NoError(t, err)
	assert.Len(t, ic.testEmits, 3)
	assert.True(t, ic.testEmits["databricks_user[<unknown>] (user_name: a)"])
	assert.True(t, ic.testEmits["databricks_group[<unknown>] (display_name: b)"])
	assert.True(t, ic.testEmits["databricks_service_principal[<unknown>] (application_id: 123)"])
}

func TestSecretScope(t *testing.T) {
	d := secrets.ResourceSecretScope().TestResourceData()
	d.Set("name", "abc")
	ic := importContextForTest()
	name := ic.Importables["databricks_secret_scope"].Name(ic, d)
	assert.Equal(t, "abc_a9993e3647", name)
}

func TestInstancePoolNameFromID(t *testing.T) {
	ic := importContextForTest()
	d := pools.ResourceInstancePool().TestResourceData()
	d.SetId("a-b-c")
	d.Set("instance_pool_name", "")
	assert.Equal(t, "c", resourcesMap["databricks_instance_pool"].Name(ic, d))
}

func TestClusterNameFromID(t *testing.T) {
	ic := importContextForTest()
	d := clusters.ResourceCluster().TestResourceData()
	d.SetId("a-b-c")
	assert.Equal(t, "c", resourcesMap["databricks_cluster"].Name(ic, d))
}

func TestRepoName(t *testing.T) {
	ic := importContextForTest()
	d := repos.ResourceRepo().TestResourceData()
	d.SetId("12345")
	// Repo without path
	assert.Equal(t, "repo_12345", resourcesMap["databricks_repo"].Name(ic, d))
	// Repo with path
	d.Set("path", "/Repos/user/test")
	assert.Equal(t, "user_test_12345", resourcesMap["databricks_repo"].Name(ic, d))
}

func TestRepoIgnore(t *testing.T) {
	ic := importContextForTest()
	d := repos.ResourceRepo().TestResourceData()
	d.SetId("12345")
	d.Set("path", "/Repos/user/test")
	r := &resource{ID: "12345", Data: d}
	// Repo without URL
	assert.True(t, resourcesMap["databricks_repo"].Ignore(ic, r))
	assert.Equal(t, 1, len(ic.ignoredResources))
	// Repo with URL
	d.Set("url", "https://github.com/abc/abc.git")
	assert.False(t, resourcesMap["databricks_repo"].Ignore(ic, r))
}

func TestDLTIgnore(t *testing.T) {
	ic := importContextForTest()
	d := pipelines.ResourcePipeline().TestResourceData()
	d.SetId("12345")
	r := &resource{ID: "12345", Data: d}
	// job without libraries
	assert.True(t, resourcesMap["databricks_pipeline"].Ignore(ic, r))
	assert.Equal(t, 1, len(ic.ignoredResources))
}

func TestJobsIgnore(t *testing.T) {
	ic := importContextForTest()
	d := jobs.ResourceJob().TestResourceData()
	d.SetId("12345")
	r := &resource{ID: "12345", Data: d}
	// job without tasks
	assert.True(t, resourcesMap["databricks_job"].Ignore(ic, r))
	assert.Equal(t, 1, len(ic.ignoredResources))
}

func TestJobName(t *testing.T) {
	ic := importContextForTest()
	d := jobs.ResourceJob().TestResourceData()
	d.SetId("12345")
	// job without name
	assert.Equal(t, "job_12345", resourcesMap["databricks_job"].Name(ic, d))
	// job with name
	d.Set("name", "test@1pm")
	assert.Equal(t, "test_1pm_12345", resourcesMap["databricks_job"].Name(ic, d))
}

func TestImportClusterLibraries(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			Method:       "GET",
			Status:       200,
			Resource:     "/api/2.0/libraries/cluster-status?cluster_id=abc",
			Response: libraries.ClusterLibraryStatuses{
				LibraryStatuses: []libraries.LibraryStatus{
					{
						Library: &libraries.Library{
							Whl: "foo.whl",
						},
						Status: "INSTALLED",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		d := clusters.ResourceCluster().TestResourceData()
		d.SetId("abc")
		err := resourcesMap["databricks_cluster"].Import(ic, &resource{
			ID:   "abc",
			Data: d,
		})
		assert.NoError(t, err)
	})
}

func TestImportClusterLibrariesFails(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			Method:       "GET",
			Status:       404,
			Resource:     "/api/2.0/libraries/cluster-status?cluster_id=abc",
			Response:     apierr.NotFound("nope"),
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		d := clusters.ResourceCluster().TestResourceData()
		d.SetId("abc")
		err := resourcesMap["databricks_cluster"].Import(ic, &resource{
			ID:   "abc",
			Data: d,
		})
		assert.EqualError(t, err, "nope")
	})
}

func TestClusterListFails(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/list",
			Status:   404,
			Response: apierr.NotFound("nope"),
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		err := resourcesMap["databricks_cluster"].List(ic)
		assert.EqualError(t, err, "nope")
	})
}

func TestClusterList_NoNameMatch(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/list",
			Response: clusters.ClusterList{
				Clusters: []clusters.ClusterInfo{
					{
						ClusterName: "abc",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.match = "bcd"
		err := resourcesMap["databricks_cluster"].List(ic)
		assert.NoError(t, err)
		assert.Equal(t, 0, len(ic.testEmits))
	})
}

func TestInstnacePoolsListWithMatch(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/instance-pools/list",
			Response: compute.ListInstancePools{
				InstancePools: []compute.InstancePoolAndStats{
					{
						InstancePoolName: "test",
						InstancePoolId:   "test",
					},
					{
						InstancePoolName: "bcd",
						InstancePoolId:   "bcd",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.match = "bcd"
		err := resourcesMap["databricks_instance_pool"].List(ic)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(ic.testEmits))
	})
}

func TestJobListNoNameMatch(t *testing.T) {
	ic := importContextForTest()
	ic.match = "bcd"
	ic.importJobs([]jobs.Job{
		{
			Settings: &jobs.JobSettings{
				Name: "abc",
			},
		},
	})
	assert.Equal(t, 0, len(ic.testEmits))
}

func TestClusterPolicyWrongDef(t *testing.T) {
	d := policies.ResourceClusterPolicy().TestResourceData()
	d.Set("name", "abc")
	d.Set("definition", "..")
	ic := importContextForTest()
	err := resourcesMap["databricks_cluster_policy"].Import(ic, &resource{
		ID:   "x",
		Data: d,
	})
	assert.EqualError(t, err, "invalid character '.' looking for beginning of value")
}

func TestClusterPolicyNoValues(t *testing.T) {
	d := policies.ResourceClusterPolicy().TestResourceData()
	d.Set("name", "abc")
	d.Set("definition", `{"foo": {}}`)
	ic := importContextForTest()
	ic.meAdmin = true
	err := resourcesMap["databricks_cluster_policy"].Import(ic, &resource{
		ID:   "x",
		Data: d,
	})
	assert.NoError(t, err)
	assert.Equal(t, 1, len(ic.testEmits))
}

func TestGroupCacheError(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			Method:       "GET",
			Resource:     "/api/2.0/preview/scim/v2/Groups?attributes=id&count=100&startIndex=1",
			Status:       404,
			Response:     apierr.NotFound("nope"),
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		err := resourcesMap["databricks_group"].List(ic)
		assert.EqualError(t, err, "nope")

		err = resourcesMap["databricks_group"].Search(ic, &resource{
			ID: "nonsense",
		})
		assert.EqualError(t, err, "nope")
		d := scim.ResourceGroup().TestResourceData()
		d.Set("display_name", "nonsense")
		err = resourcesMap["databricks_group"].Import(ic, &resource{
			ID:   "nonsense",
			Data: d,
		})
		assert.EqualError(t, err, "nope")
	})
}

func TestGroupListNoNameMatch(t *testing.T) {
	ic := importContextForTest()
	ic.match = "bcd"
	ic.allGroups = []scim.Group{
		{
			DisplayName: "abc",
		},
	}
	err := resourcesMap["databricks_group"].List(ic)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(ic.testEmits))
}

func TestGroupSearchNoMatch(t *testing.T) {
	ic := importContextForTest()
	ic.allGroups = []scim.Group{
		{
			DisplayName: "abc",
		},
	}
	r := &resource{
		Attribute: "display_name",
		Value:     "dbc",
	}
	err := resourcesMap["databricks_group"].Search(ic, r)
	assert.NoError(t, err)
	assert.Equal(t, "", r.ID)
}

func TestUserSearchFails(t *testing.T) {
	userFixture := qa.ListUsersFixtures([]iam.User{})
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		userFixture[0],
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		d := scim.ResourceUser().TestResourceData()
		d.Set("user_name", "dbc")
		r := &resource{
			Attribute: "display_name",
			Value:     "dbc",
			Data:      d,
		}
		err := resourcesMap["databricks_user"].Search(ic, r)
		assert.EqualError(t, err, "there is no user 'dbc'")

		err = resourcesMap["databricks_user"].Import(ic, r)
		assert.EqualError(t, err, "user dbc is not found")
	})
}

func TestSpnSearchFails(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		qa.ListServicePrincipalsFixtures([]iam.ServicePrincipal{})[0],
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		d := scim.ResourceServicePrincipal().TestResourceData()
		d.Set("application_id", "dbc")
		r := &resource{
			Attribute: "application_id",
			Value:     "dbc",
			Data:      d,
		}
		err := resourcesMap["databricks_service_principal"].Search(ic, r)
		assert.EqualError(t, err, "there is no service principal 'dbc'")

		err = resourcesMap["databricks_service_principal"].Import(ic, r)
		assert.EqualError(t, err, "service principal dbc is not found")
	})
}

func TestSpnSearchSuccess(t *testing.T) {
	spFixture := qa.ListServicePrincipalsFixtures([]iam.ServicePrincipal{
		{
			Id: "321", DisplayName: "spn", ApplicationId: "dbc",
		},
	})
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		spFixture[0],
		spFixture[1],
		{
			ReuseRequest: true,
			Method:       "GET",
			Resource:     "/api/2.0/preview/scim/v2/ServicePrincipals/321?attributes=userName,displayName,active,externalId,entitlements,groups,roles",
			Response:     scim.User{ID: "321", DisplayName: "spn", ApplicationID: "dbc"},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		d := scim.ResourceServicePrincipal().TestResourceData()
		d.Set("application_id", "dbc")
		d.Set("display_name", "dbc")
		r := &resource{
			Attribute: "application_id",
			Value:     "dbc",
			Data:      d,
		}
		err := resourcesMap["databricks_service_principal"].Search(ic, r)
		assert.NoError(t, err)

		err = resourcesMap["databricks_service_principal"].Import(ic, r)
		assert.NoError(t, err)

		assert.True(t, resourcesMap["databricks_service_principal"].ShouldOmitField(ic, "application_id",
			scim.ResourceServicePrincipal().Schema["application_id"], d))
		ic.Client.Config.Host = "https://abc.azuredatabricks.net"
		assert.True(t, resourcesMap["databricks_service_principal"].ShouldOmitField(ic, "display_name",
			scim.ResourceServicePrincipal().Schema["display_name"], d))

		// test for different branches in Name function
		d2 := scim.ResourceServicePrincipal().TestResourceData()
		d2.SetId("123")
		d2.Set("application_id", "dbc")
		assert.Equal(t, "dbc_123", resourcesMap["databricks_service_principal"].Name(ic, d2))
		d2.Set("application_id", "60622399-fd3f-4faf-8810-bf08b225cf3b")
		assert.Equal(t, "60622399_123", resourcesMap["databricks_service_principal"].Name(ic, d2))

		d2.Set("display_name", "abc")
		assert.Equal(t, "abc_123", resourcesMap["databricks_service_principal"].Name(ic, d2))
	})
}

func TestShouldOmitForUsers(t *testing.T) {
	d := scim.ResourceUser().TestResourceData()
	d.SetId("user1")
	d.Set("user_name", "user@domain.com")
	d.Set("display_name", "")
	assert.True(t, resourcesMap["databricks_user"].ShouldOmitField(nil, "display_name",
		scim.ResourceUser().Schema["application_id"], d))
	d.Set("display_name", "user@domain.com")
	assert.True(t, resourcesMap["databricks_user"].ShouldOmitField(nil, "display_name",
		scim.ResourceUser().Schema["application_id"], d))
	d.Set("display_name", "Some user")
	assert.False(t, resourcesMap["databricks_user"].ShouldOmitField(nil, "display_name",
		scim.ResourceUser().Schema["application_id"], d))
}

func TestShouldOmitFoRepos(t *testing.T) {
	d := repos.ResourceRepo().TestResourceData()
	d.SetId("1234")
	d.Set("path", "/Repos/Test/repo")
	assert.False(t, resourcesMap["databricks_repo"].ShouldOmitField(nil, "path",
		repos.ResourceRepo().Schema["path"], d))
	assert.True(t, resourcesMap["databricks_repo"].ShouldOmitField(nil, "branch",
		repos.ResourceRepo().Schema["branch"], d))
	assert.True(t, resourcesMap["databricks_repo"].ShouldOmitField(nil, "tag",
		repos.ResourceRepo().Schema["tag"], d))
	d.Set("branch", "test")
	assert.False(t, resourcesMap["databricks_repo"].ShouldOmitField(nil, "branch",
		repos.ResourceRepo().Schema["branch"], d))
	d.Set("tag", "v123")
	assert.False(t, resourcesMap["databricks_repo"].ShouldOmitField(nil, "tag",
		repos.ResourceRepo().Schema["tag"], d))
}

func TestUserImportSkipNonDirectGroups(t *testing.T) {
	userFixture := qa.ListUsersFixtures([]iam.User{
		{
			UserName: "dbc",
			Id:       "321",
		},
	})
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		userFixture[0],
		userFixture[1],
		{
			ReuseRequest: true,
			Method:       "GET",
			Resource:     "/api/2.0/preview/scim/v2/Users/321?attributes=id,userName,displayName,active,externalId,entitlements,groups,roles",
			Response: scim.UserList{
				Resources: []scim.User{
					{
						Groups: []scim.ComplexValue{
							{
								Display: "x",
								Value:   "y",
							},
						},
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		d := scim.ResourceUser().TestResourceData()
		d.Set("user_name", "dbc")
		r := &resource{
			Attribute: "display_name",
			Value:     "dbc",
			Data:      d,
		}
		err := resourcesMap["databricks_user"].Import(ic, r)
		assert.NoError(t, err)
		assert.Equal(t, 0, len(ic.testEmits))
	})
}

func TestSecretScopeListNoNameMatch(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/secrets/scopes/list",

			Response: secrets.SecretScopeList{
				Scopes: []secrets.SecretScope{
					{
						Name: "abc",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.match = "bcd"
		err := resourcesMap["databricks_secret_scope"].List(ic)
		assert.NoError(t, err)
		assert.Equal(t, 0, len(ic.testEmits))
	})
}

func TestPoliciesListing(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/policies/clusters/list?",
			Response: compute.ListPoliciesResponse{
				Policies: []compute.Policy{
					{
						Name:           "Personal Compute",
						PolicyFamilyId: "personal-vm",
						PolicyId:       "123",
					},
					{
						Name:     "abcd",
						PolicyId: "456",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/policy-families?",
			Response: compute.ListPolicyFamiliesResponse{
				PolicyFamilies: []compute.PolicyFamily{
					{
						Name:           "Personal Compute",
						PolicyFamilyId: "personal-vm",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		err := resourcesMap["databricks_cluster_policy"].List(ic)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(ic.testEmits))
	})
}

func TestPoliciesListNoNameMatch(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/policies/clusters/list?",
			Response: compute.ListPoliciesResponse{
				Policies: []compute.Policy{
					{
						Name: "Personal Compute",
					},
					{
						Name: "abcd",
					},
				},
			},
		},
		emptyPolicyFamilies,
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.match = "bcd"
		err := resourcesMap["databricks_cluster_policy"].List(ic)
		assert.NoError(t, err)
		assert.Equal(t, 0, len(ic.testEmits))
	})
}

func TestAwsS3MountProfile(t *testing.T) {
	ic := importContextForTest()
	ic.mounts = true
	ic.match = "abc"
	ic.mountMap = map[string]mount{}
	ic.mountMap["/mnt/abc"] = mount{
		URL:             "s3a://abc",
		InstanceProfile: "bcd",
	}
	ic.mountMap["/mnt/def"] = mount{
		URL:             "s3a://def",
		InstanceProfile: "bcd",
	}
	err := resourcesMap["databricks_mount"].List(ic)
	assert.NoError(t, err)
	assert.Len(t, ic.testEmits, 2)
	assert.True(t, ic.testEmits["databricks_instance_profile[<unknown>] (id: bcd)"])
	assert.True(t, ic.testEmits["databricks_mount[<unknown>] (id: /mnt/abc)"])
}

func TestMountsBodyGeneration(t *testing.T) {
	ic := importContextForTest()
	ic.mounts = true
	ic.match = "abc"
	ic.mountMap = map[string]mount{}
	ic.variables = map[string]string{}
	ic.mountMap["/mnt/abc"] = mount{
		URL:             "s3a://abc",
		InstanceProfile: "bcd",
	}
	ic.mountMap["/mnt/def"] = mount{
		URL:       "s3a://def",
		ClusterID: "bcd",
	}
	ic.mountMap["/mnt/gcs"] = mount{
		URL:       "gs://gcs/dir",
		ClusterID: "bcd",
	}
	ic.mountMap["/mnt/abfss"] = mount{
		URL: "abfss://test@test.dfs.core.windows.net/directory",
	}
	ic.mountMap["/mnt/wasbs"] = mount{
		URL: "wasbs://test@test.blob.core.windows.net/directory",
	}
	ic.mountMap["/mnt/adls"] = mount{
		URL: "adls://test.dfs.core.windows.net/directory",
	}
	ic.mountMap["/mnt/dbfs"] = mount{
		URL: "dbfs:/directory",
	}

	//
	f := hclwrite.NewEmptyFile()
	body := f.Body()

	err := generateMountBody(ic, body, &resource{
		ID:       "/mnt/abc",
		Name:     "abc",
		Resource: "databricks_mount",
	})
	assert.NoError(t, err)

	err = generateMountBody(ic, body, &resource{
		ID:       "/mnt/def",
		Name:     "def",
		Resource: "databricks_mount",
	})
	assert.NoError(t, err)

	err = generateMountBody(ic, body, &resource{
		ID:       "/mnt/abfss",
		Name:     "abfss",
		Resource: "databricks_mount",
	})
	assert.NoError(t, err)

	err = generateMountBody(ic, body, &resource{
		ID:       "/mnt/gcs",
		Name:     "gcs",
		Resource: "databricks_mount",
	})
	assert.NoError(t, err)

	err = generateMountBody(ic, body, &resource{
		ID:       "/mnt/adls",
		Name:     "adls",
		Resource: "databricks_mount",
	})
	assert.NoError(t, err)

	err = generateMountBody(ic, body, &resource{
		ID:       "/mnt/wasbs",
		Name:     "wasbs",
		Resource: "databricks_mount",
	})
	assert.NoError(t, err)

	err = generateMountBody(ic, body, &resource{
		ID:       "/mnt/dbfs",
		Name:     "dbfs",
		Resource: "databricks_mount",
	})
	assert.EqualError(t, err, "no matching handler for: dbfs:/directory")
}

func TestGlobalInitScriptNameFromId(t *testing.T) {
	ic := importContextForTest()
	d := workspace.ResourceGlobalInitScript().TestResourceData()
	d.SetId("abc")
	assert.Equal(t, "abc", resourcesMap["databricks_global_init_script"].Name(ic, d))
}

func TestGlobalInitScriptsErrors(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			MatchAny:     true,
			Status:       404,
			Response:     apierr.NotFound("nope"),
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		err := resourcesMap["databricks_global_init_script"].List(ic)
		assert.EqualError(t, err, "nope")

		err = resourcesMap["databricks_global_init_script"].Import(ic, &resource{
			ID: "abc",
		})
		assert.EqualError(t, err, "nope")
	})
}

func TestGlobalInitScriptsBodyErrors(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/global-init-scripts/sad-emoji",
			Response: workspace.GlobalInitScriptInfo{
				Name:          "x.sh",
				ContentBase64: "🥺",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/global-init-scripts/second",
			Response: workspace.GlobalInitScriptInfo{
				Name:          "x.sh",
				ContentBase64: "YWJj",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		err := resourcesMap["databricks_global_init_script"].Import(ic, &resource{
			ID: "sad-emoji",
		})
		assert.EqualError(t, err, "illegal base64 data at input byte 0")

		err = resourcesMap["databricks_global_init_script"].Import(ic, &resource{
			ID: "second",
		})
		assert.NotNil(t, err) // no exact match because of OS diffs
	})
}

func TestRepoListFails(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			MatchAny:     true,
			Status:       404,
			Response:     apierr.NotFound("nope"),
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		err := resourcesMap["databricks_repo"].List(ic)
		assert.EqualError(t, err, "nope")
	})
}

func testGenerate(t *testing.T, fixtures []qa.HTTPFixture, services string, asAdmin bool, cb func(*importContext)) {
	qa.HTTPFixturesApply(t, fixtures, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.Directory = fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
		os.MkdirAll(ic.Directory, 0755)
		defer os.RemoveAll(ic.Directory)
		ic.testEmits = nil
		ic.meAdmin = asAdmin
		ic.importing = map[string]bool{}
		ic.variables = map[string]string{}
		ic.services = strings.Split(services, ",")
		ic.startImportChannels()
		cb(ic)
	})
}

func getGeneratedFile(ic *importContext, service string) string {
	fileName := fmt.Sprintf("%s/%s.tf", ic.Directory, service)
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Printf("[ERROR] can't read file %s", fileName)
		return ""
	}
	return string(content)
}

func TestNotebookGeneration(t *testing.T) {
	testGenerate(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/list?path=%2F",
			Response: workspace.ObjectList{
				Objects: []workspace.ObjectStatus{
					{
						Path:       "/Repos/Foo/Bar",
						ObjectType: "NOTEBOOK",
					},
					{
						Path:       "/First/Second",
						ObjectType: "NOTEBOOK",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/get-status?path=%2FFirst%2FSecond",
			Response: workspace.ObjectStatus{
				ObjectID:   123,
				ObjectType: "NOTEBOOK",
				Path:       "/First/Second",
				Language:   "PYTHON",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/export?format=SOURCE&path=%2FFirst%2FSecond",
			Response: workspace.ExportPath{
				Content: "YWJj",
			},
		},
	}, "notebooks", false, func(ic *importContext) {
		ic.notebooksFormat = "SOURCE"
		err := resourcesMap["databricks_notebook"].List(ic)
		assert.NoError(t, err)
		ic.waitGroup.Wait()
		ic.closeImportChannels()
		ic.generateAndWriteResources(nil)
		assert.Equal(t, commands.TrimLeadingWhitespace(`
		resource "databricks_notebook" "first_second_123" {
		  source = "${path.module}/notebooks/First/Second_123.py"
		  path   = "/First/Second"
		}`), getGeneratedFile(ic, "notebooks"))
	})
}

func TestNotebookGenerationJupyter(t *testing.T) {
	testGenerate(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/list?path=%2F",
			Response: workspace.ObjectList{
				Objects: []workspace.ObjectStatus{
					{
						Path:       "/Repos/Foo/Bar",
						ObjectType: "NOTEBOOK",
					},
					{
						Path:       "/First/Second",
						ObjectType: "NOTEBOOK",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/get-status?path=%2FFirst%2FSecond",
			Response: workspace.ObjectStatus{
				ObjectID:   123,
				ObjectType: "NOTEBOOK",
				Path:       "/First/Second",
				Language:   "PYTHON",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/export?format=JUPYTER&path=%2FFirst%2FSecond",
			Response: workspace.ExportPath{
				Content: "YWJj",
			},
		},
	}, "notebooks", false, func(ic *importContext) {
		ic.notebooksFormat = "JUPYTER"
		err := resourcesMap["databricks_notebook"].List(ic)
		assert.NoError(t, err)
		ic.waitGroup.Wait()
		ic.closeImportChannels()
		ic.generateAndWriteResources(nil)
		assert.Equal(t, commands.TrimLeadingWhitespace(`
		resource "databricks_notebook" "first_second_123" {
		  source   = "${path.module}/notebooks/First/Second_123.ipynb"
		  path     = "/First/Second"
		  language = "PYTHON"
		  format   = "JUPYTER"
		}`), getGeneratedFile(ic, "notebooks"))
	})
}

func TestNotebookGenerationBadCharacters(t *testing.T) {
	testGenerate(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/list?path=%2F",
			Response: workspace.ObjectList{
				Objects: []workspace.ObjectStatus{
					{
						Path:       "/Repos/Foo/Bar",
						ObjectType: "NOTEBOOK",
					},
					{
						Path:       "/Fir\"st\\/Second",
						ObjectType: "NOTEBOOK",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/get-status?path=%2FFir%22st%5C",

			Response: workspace.ObjectStatus{
				ObjectID:   124,
				ObjectType: "DIRECTORY",
				Path:       "/Fir\"st\\",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/get-status?path=%2FFir%22st%5C%2FSecond",
			Response: workspace.ObjectStatus{
				ObjectID:   123,
				ObjectType: "NOTEBOOK",
				Path:       "/Fir\"st\\/Second",
				Language:   "PYTHON",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/export?format=SOURCE&path=%2FFir%22st%5C%2FSecond",
			Response: workspace.ExportPath{
				Content: "YWJj",
			},
		},
	}, "notebooks,directories", true, func(ic *importContext) {
		ic.notebooksFormat = "SOURCE"
		ic.services = []string{"notebooks"}
		err := resourcesMap["databricks_notebook"].List(ic)
		assert.NoError(t, err)
		ic.waitGroup.Wait()
		ic.closeImportChannels()
		ic.generateAndWriteResources(nil)
		assert.Equal(t, commands.TrimLeadingWhitespace(`
		resource "databricks_notebook" "fir_st_second_123" {
		  source = "${path.module}/notebooks/Fir_st_/Second_123.py"
		  path   = "/Fir\"st\\/Second"
		}`), getGeneratedFile(ic, "notebooks"))
	})
}

func TestDirectoryGeneration(t *testing.T) {
	testGenerate(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/list?path=%2F",
			Response: workspace.ObjectList{
				Objects: []workspace.ObjectStatus{
					{
						ObjectID:   1234,
						Path:       "/first",
						ObjectType: "DIRECTORY",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/list?path=%2Ffirst",
			Response: workspace.ObjectList{
				Objects: []workspace.ObjectStatus{
					{},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/get-status?path=%2Ffirst",
			Response: workspace.ObjectStatus{
				ObjectID:   1234,
				ObjectType: "DIRECTORY",
				Path:       "/first",
			},
		},
	}, "directories", false, func(ic *importContext) {
		err := resourcesMap["databricks_directory"].List(ic)
		assert.NoError(t, err)

		ic.waitGroup.Wait()
		ic.closeImportChannels()
		ic.generateAndWriteResources(nil)
		assert.Equal(t, commands.TrimLeadingWhitespace(`
		resource "databricks_directory" "first_1234" {
		  path = "/first"
		}`), getGeneratedFile(ic, "directories"))
	})
}

func TestGlobalInitScriptGen(t *testing.T) {
	testGenerate(t, []qa.HTTPFixture{
		{
			Method:       "GET",
			ReuseRequest: true,
			Resource:     "/api/2.0/global-init-scripts/a",
			Response: workspace.GlobalInitScriptInfo{
				Name:          "New: Importing ^ Things",
				Enabled:       true,
				ContentBase64: "YWJj",
			},
		},
	}, "workspace", false, func(ic *importContext) {
		ic.Emit(&resource{
			Resource: "databricks_global_init_script",
			ID:       "a",
		})

		ic.waitGroup.Wait()
		ic.closeImportChannels()
		ic.generateAndWriteResources(nil)
		assert.Equal(t, commands.TrimLeadingWhitespace(`
		resource "databricks_global_init_script" "new_importing_things" {
		  source  = "${path.module}/files/new_importing_things.sh"
		  name    = "New: Importing ^ Things"
		  enabled = true
		}`), getGeneratedFile(ic, "workspace"))
	})
}

func TestSecretGen(t *testing.T) {
	testGenerate(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/secrets/list?scope=a",

			Response: secrets.SecretsList{
				Secrets: []secrets.SecretMetadata{
					{
						Key: "b",
					},
				},
			},
		},
	}, "secrets", false, func(ic *importContext) {
		ic.Emit(&resource{
			Resource: "databricks_secret",
			ID:       "a|||b",
		})

		ic.waitGroup.Wait()
		ic.closeImportChannels()
		ic.generateAndWriteResources(nil)
		assert.Equal(t, commands.TrimLeadingWhitespace(`
		resource "databricks_secret" "a_b_eb2980a5a2" {
		  string_value = var.string_value_a_b_eb2980a5a2
		  scope        = "a"
		  key          = "b"
		}`), getGeneratedFile(ic, "secrets"))
	})
}

func TestDbfsFileGen(t *testing.T) {
	testGenerate(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/dbfs/get-status?path=a",
			Response: storage.FileInfo{
				Path: "a",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/dbfs/read?length=1000000&path=a",
			Response: storage.ReadResponse{
				Data:      "YWJj",
				BytesRead: 3,
			},
		},
	}, "storage", false, func(ic *importContext) {
		ic.Emit(&resource{
			Resource: "databricks_dbfs_file",
			ID:       "a",
		})

		ic.waitGroup.Wait()
		ic.closeImportChannels()
		ic.generateAndWriteResources(nil)
		assert.Equal(t, commands.TrimLeadingWhitespace(`
		resource "databricks_dbfs_file" "_0cc175b9c0f1b6a831c399e269772661_a" {
		  source = "${path.module}/files/_0cc175b9c0f1b6a831c399e269772661_a"
		  path   = "a"
		}`), getGeneratedFile(ic, "storage"))
	})
}

func TestSqlListObjects(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/sql/queries?page_size=100",
			Response: dbsqlListResponse{PageSize: 1, Page: 1, TotalCount: 2,
				Results: []map[string]any{{"key1": "value1"}}},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/sql/queries?page=2&page_size=100",
			Response: dbsqlListResponse{PageSize: 1, Page: 2, TotalCount: 2,
				Results: []map[string]any{{"key2": "value2"}}},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		answer, err := dbsqlListObjects(ic, "/preview/sql/queries")
		assert.NoError(t, err)
		assert.Len(t, answer, 2)
	})
}

func TestIncrementalListDLT(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/pipelines?max_results=50",
			Response: pipelines.PipelineListResponse{
				Statuses: []pipelines.PipelineStateInfo{
					{
						PipelineID: "abc",
						Name:       "abc",
					},
					{
						PipelineID: "def",
						Name:       "def",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/pipelines/abc",
			Response: pipelines.PipelineInfo{
				PipelineID:   "abc",
				Name:         "abc",
				LastModified: 1681466931226,
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/pipelines/def",
			Response: pipelines.PipelineInfo{
				PipelineID:   "def",
				Name:         "def",
				LastModified: 1690156900000,
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.incremental = true
		ic.updatedSinceStr = "2023-07-24T00:00:00Z"
		ic.updatedSinceMs = 1690156700000
		err := resourcesMap["databricks_pipeline"].List(ic)
		ic.waitGroup.Wait()
		ic.closeImportChannels()
		assert.NoError(t, err)
		assert.Equal(t, 1, len(ic.testEmits))
	})
}

func TestListSystemSchemasSuccess(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		currentMetastoreSuccess,
		{
			Method:   "GET",
			Resource: fmt.Sprintf("/api/2.1/unity-catalog/metastores/%s/systemschemas?", currentMetastoreResponse.MetastoreId),
			Response: catalog.ListSystemSchemasResponse{
				Schemas: []catalog.SystemSchemaInfo{
					{
						Schema: "access",
						State:  catalog.SystemSchemaInfoStateEnableCompleted,
					},
					{
						Schema: "marketplace",
						State:  catalog.SystemSchemaInfoStateAvailable,
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.currentMetastore = currentMetastoreResponse
		err := resourcesMap["databricks_system_schema"].List(ic)
		assert.NoError(t, err)
		assert.Equal(t, len(ic.testEmits), 1)
	})
}

func TestListSystemSchemasErrorGetMetastore(t *testing.T) {
	ic := importContextForTest()
	err := resourcesMap["databricks_system_schema"].List(ic)
	assert.EqualError(t, err, "there is no UC metastore information")
}

func TestListSystemSchemasErrorListing(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: fmt.Sprintf("/api/2.1/unity-catalog/metastores/%s/systemschemas?", currentMetastoreResponse.MetastoreId),
			Status:   404,
			Response: apierr.NotFound("nope"),
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.currentMetastore = currentMetastoreResponse
		err := resourcesMap["databricks_system_schema"].List(ic)
		assert.EqualError(t, err, "nope")
	})
}

func TestListUcAllowListError(t *testing.T) {
	ic := importContextForTest()
	err := resourcesMap["databricks_artifact_allowlist"].List(ic)
	assert.EqualError(t, err, "there is no UC metastore information")
}

func TestListUcAllowListSuccess(t *testing.T) {
	ic := importContextForTest()
	ic.currentMetastore = currentMetastoreResponse
	err := resourcesMap["databricks_artifact_allowlist"].List(ic)
	assert.NoError(t, err)
	assert.Equal(t, len(ic.testEmits), 3)
}

func TestEmitSqlParent(t *testing.T) {
	ic := importContextForTest()
	ic.emitSqlParentDirectory("")
	assert.Equal(t, 0, len(ic.testEmits))
	ic.emitSqlParentDirectory("folders/12345")
	assert.Equal(t, 1, len(ic.testEmits))
	assert.Contains(t, ic.testEmits, "databricks_directory[<unknown>] (object_id: 12345)")
}

func TestEmitFilesFromSlice(t *testing.T) {
	ic := importContextForTest()
	ic.emitFilesFromSlice([]string{
		"dbfs:/FileStore/test.txt",
		"/Workspace/Shared/test.txt",
		"nothing",
	})
	assert.Equal(t, 2, len(ic.testEmits))
	assert.Contains(t, ic.testEmits, "databricks_dbfs_file[<unknown>] (id: dbfs:/FileStore/test.txt)")
	assert.Contains(t, ic.testEmits, "databricks_workspace_file[<unknown>] (id: /Shared/test.txt)")
}

func TestEmitFilesFromMap(t *testing.T) {
	ic := importContextForTest()
	ic.emitFilesFromMap(map[string]string{
		"k1": "dbfs:/FileStore/test.txt",
		"k2": "/Workspace/Shared/test.txt",
		"k3": "nothing",
	})
	assert.Equal(t, 2, len(ic.testEmits))
	assert.Contains(t, ic.testEmits, "databricks_dbfs_file[<unknown>] (id: dbfs:/FileStore/test.txt)")
	assert.Contains(t, ic.testEmits, "databricks_workspace_file[<unknown>] (id: /Shared/test.txt)")
}
