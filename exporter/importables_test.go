package exporter

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"sync"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/databricks-sdk-go/service/iam"
	sdk_jobs "github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/databricks-sdk-go/service/sharing"
	sdk_workspace "github.com/databricks/databricks-sdk-go/service/workspace"
	tfcatalog "github.com/databricks/terraform-provider-databricks/catalog"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/commands"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/jobs"
	"github.com/databricks/terraform-provider-databricks/permissions"
	"github.com/databricks/terraform-provider-databricks/pipelines"
	"github.com/databricks/terraform-provider-databricks/policies"
	"github.com/databricks/terraform-provider-databricks/pools"
	"github.com/databricks/terraform-provider-databricks/provider"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/databricks/terraform-provider-databricks/repos"
	"github.com/databricks/terraform-provider-databricks/scim"
	"github.com/databricks/terraform-provider-databricks/secrets"
	tfsharing "github.com/databricks/terraform-provider-databricks/sharing"
	"github.com/databricks/terraform-provider-databricks/storage"
	"github.com/databricks/terraform-provider-databricks/workspace"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/maps"
)

func importContextForTest() *importContext {
	p := provider.DatabricksProvider()
	supportedResources := maps.Keys(resourcesMap)
	return &importContext{
		Importables:               resourcesMap,
		Resources:                 p.ResourcesMap,
		testEmits:                 map[string]bool{},
		nameFixes:                 nameFixes,
		waitGroup:                 &sync.WaitGroup{},
		allUsers:                  map[string]scim.User{},
		allSps:                    map[string]scim.User{},
		channels:                  makeResourcesChannels(),
		oldWorkspaceObjectMapping: map[int64]string{},
		exportDeletedUsersAssets:  false,
		ignoredResources:          map[string]struct{}{},
		deletedResources:          map[string]struct{}{},
		State:                     newStateApproximation(supportedResources),
		emittedUsers:              map[string]struct{}{},
		userOrSpDirectories:       map[string]bool{},
		defaultChannel:            make(resourceChannel, defaultChannelSize),
		services:                  map[string]struct{}{},
		listing:                   map[string]struct{}{},
		tfvars:                    map[string]string{},
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
	d := pools.ResourceInstancePool().ToResource().TestResourceData()
	d.Set("instance_pool_name", "blah-bah")
	ic := importContextForTest()
	ic.enableServices("access,pools")

	name := resourcesMap["databricks_instance_pool"].Name(ic, d)
	assert.Equal(t, "blah-bah", name)

	d = pools.ResourceInstancePool().ToResource().TestResourceData()
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
	d := policies.ResourceClusterPolicy().ToResource().TestResourceData()
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
	ic.enableServices("storage,pools,policies,access")
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
	d := policies.ResourceClusterPolicy().ToResource().TestResourceData()
	d.Set("policy_family_id", "job-cluster")
	d.Set("name", "Job Compute")
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
	ic.enableServices("groups,access")
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
	d := scim.ResourceGroup().ToResource().TestResourceData()
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
	d := p.ToResource().TestResourceData()
	d.SetId("abc")
	ic := importContextForTest()
	ic.enableServices("access,users,groups")
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
	d := secrets.ResourceSecretScope().ToResource().TestResourceData()
	d.Set("name", "abc")
	ic := importContextForTest()
	name := ic.Importables["databricks_secret_scope"].Name(ic, d)
	assert.Equal(t, "abc_a9993e3647", name)
}

func TestInstancePoolNameFromID(t *testing.T) {
	ic := importContextForTest()
	d := pools.ResourceInstancePool().ToResource().TestResourceData()
	d.SetId("a-b-c")
	d.Set("instance_pool_name", "")
	assert.Equal(t, "c", resourcesMap["databricks_instance_pool"].Name(ic, d))
}

func TestClusterNameFromID(t *testing.T) {
	ic := importContextForTest()
	d := clusters.ResourceCluster().ToResource().TestResourceData()
	d.SetId("a-b-c")
	assert.Equal(t, "c", resourcesMap["databricks_cluster"].Name(ic, d))
}

func TestRepoName(t *testing.T) {
	ic := importContextForTest()
	d := repos.ResourceRepo().ToResource().TestResourceData()
	d.SetId("12345")
	// Repo without path
	assert.Equal(t, "repo_12345", resourcesMap["databricks_repo"].Name(ic, d))
	// Repo with path
	d.Set("path", "/Repos/user/test")
	assert.Equal(t, "user_test_12345", resourcesMap["databricks_repo"].Name(ic, d))
}

func TestRepoIgnore(t *testing.T) {
	ic := importContextForTest()
	d := repos.ResourceRepo().ToResource().TestResourceData()
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
	d := pipelines.ResourcePipeline().ToResource().TestResourceData()
	d.SetId("12345")
	r := &resource{ID: "12345", Data: d}
	// job without libraries
	assert.True(t, resourcesMap["databricks_pipeline"].Ignore(ic, r))
	assert.Equal(t, 1, len(ic.ignoredResources))
}

func TestJobsIgnore(t *testing.T) {
	ic := importContextForTest()
	d := jobs.ResourceJob().ToResource().TestResourceData()
	d.SetId("12345")
	r := &resource{ID: "12345", Data: d}
	// job without tasks
	assert.True(t, resourcesMap["databricks_job"].Ignore(ic, r))
	assert.Equal(t, 1, len(ic.ignoredResources))
}

func TestJobName(t *testing.T) {
	ic := importContextForTest()
	d := jobs.ResourceJob().ToResource().TestResourceData()
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
			Response: compute.ClusterLibraryStatuses{
				LibraryStatuses: []compute.LibraryFullStatus{
					{
						Library: &compute.Library{
							Whl: "foo.whl",
						},
						Status: "INSTALLED",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		d := clusters.ResourceCluster().ToResource().TestResourceData()
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
		d := clusters.ResourceCluster().ToResource().TestResourceData()
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
		ic.enableServices("pools")
		ic.match = "bcd"
		err := resourcesMap["databricks_instance_pool"].List(ic)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(ic.testEmits))
	})
}

func TestJobListNoNameMatchOrFromBundles(t *testing.T) {
	ic := importContextForTest()
	ic.match = "bcd"
	ic.importJobs([]jobs.Job{
		{
			Settings: &jobs.JobSettings{
				Name: "abc",
			},
		},
		{
			Settings: &jobs.JobSettings{
				Name:     "bcd",
				EditMode: "UI_LOCKED",
				Deployment: &sdk_jobs.JobDeployment{
					Kind: "BUNDLE",
				},
			},
		},
	})
	assert.Equal(t, 0, len(ic.testEmits))
}

func TestClusterPolicyWrongDef(t *testing.T) {
	d := policies.ResourceClusterPolicy().ToResource().TestResourceData()
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
	d := policies.ResourceClusterPolicy().ToResource().TestResourceData()
	d.Set("name", "abc")
	d.Set("definition", `{"foo": {}}`)
	ic := importContextForTest()
	ic.enableServices("access")
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
		d := scim.ResourceGroup().ToResource().TestResourceData()
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
		d := scim.ResourceUser().ToResource().TestResourceData()
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
		d := scim.ResourceServicePrincipal().ToResource().TestResourceData()
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
		d := scim.ResourceServicePrincipal().ToResource().TestResourceData()
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
		d2 := scim.ResourceServicePrincipal().ToResource().TestResourceData()
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
	d := scim.ResourceUser().ToResource().TestResourceData()
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
	d := repos.ResourceRepo().ToResource().TestResourceData()
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
		d := scim.ResourceUser().ToResource().TestResourceData()
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

			Response: sdk_workspace.ListScopesResponse{
				Scopes: []sdk_workspace.SecretScope{
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
		ic.enableServices("policies")
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
		ic.enableServices("policies")
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
	ic.enableServices("mounts,access")
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
	d := workspace.ResourceGlobalInitScript().ToResource().TestResourceData()
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
			Resource: "/api/2.0/global-init-scripts/sad-emoji?",
			Response: compute.GlobalInitScriptDetailsWithContent{
				Name:   "x.sh",
				Script: "🥺",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/global-init-scripts/second?",
			Response: compute.GlobalInitScriptDetailsWithContent{
				Name:   "x.sh",
				Script: "YWJj",
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
		ic.enableServices(services)
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
						ObjectID:   123,
						Language:   "PYTHON",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/export?format=SOURCE&path=%2FFirst%2FSecond",
			Response: workspace.ExportPath{
				Content: "YWJj",
			},
			ReuseRequest: true,
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
						ObjectID:   123,
						ObjectType: "NOTEBOOK",
						Path:       "/First/Second",
						Language:   "PYTHON",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/export?format=JUPYTER&path=%2FFirst%2FSecond",
			Response: workspace.ExportPath{
				Content: "YWJj",
			},
			ReuseRequest: true,
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
						ObjectID:   123,
						ObjectType: "NOTEBOOK",
						Path:       "/Fir\"st\\/Second",
						Language:   "PYTHON",
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
			ReuseRequest: true,
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/export?format=SOURCE&path=%2FFir%22st%5C%2FSecond",
			Response: workspace.ExportPath{
				Content: "YWJj",
			},
			ReuseRequest: true,
		},
	}, "notebooks,directories", true, func(ic *importContext) {
		ic.notebooksFormat = "SOURCE"
		ic.enableServices("notebooks")
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

func TestGlobalInitScriptGeneration(t *testing.T) {
	testGenerate(t, []qa.HTTPFixture{
		{
			Method:       "GET",
			ReuseRequest: true,
			Resource:     "/api/2.0/global-init-scripts/a?",
			Response: compute.GlobalInitScriptDetailsWithContent{
				ScriptId: "a",
				Name:     "New: Importing ^ Things",
				Enabled:  true,
				Script:   "YWJj",
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
		  source  = "${path.module}/global_init_scripts/new_importing_things.sh"
		  name    = "New: Importing ^ Things"
		  enabled = true
		}`), getGeneratedFile(ic, "workspace"))
	})
}

func TestSecretGeneration(t *testing.T) {
	testGenerate(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/secrets/list?scope=a",
			Response: sdk_workspace.ListSecretsResponse{
				Secrets: []sdk_workspace.SecretMetadata{
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

func TestDbfsFileGeneration(t *testing.T) {
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
		  source = "${path.module}/dbfs_files/_0cc175b9c0f1b6a831c399e269772661_a"
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
		ic.enableServices("dlt")
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
		ic.enableServices("uc-system-schemas")
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
	ic.enableServices("uc-artifact-allowlist")
	ic.currentMetastore = currentMetastoreResponse
	err := resourcesMap["databricks_artifact_allowlist"].List(ic)
	assert.NoError(t, err)
	assert.Equal(t, len(ic.testEmits), 3)
}

func TestEmitSqlParent(t *testing.T) {
	ic := importContextForTest()
	ic.enableServices("directories")
	ic.emitSqlParentDirectory("")
	assert.Equal(t, 0, len(ic.testEmits))
	ic.emitSqlParentDirectory("folders/12345")
	assert.Equal(t, 1, len(ic.testEmits))
	assert.Contains(t, ic.testEmits, "databricks_directory[<unknown>] (object_id: 12345)")
}

func TestEmitFilesFromSlice(t *testing.T) {
	ic := importContextForTest()
	ic.enableServices("storage,notebooks")
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
	ic.enableServices("storage,notebooks")
	ic.emitFilesFromMap(map[string]string{
		"k1": "dbfs:/FileStore/test.txt",
		"k2": "/Workspace/Shared/test.txt",
		"k3": "nothing",
	})
	assert.Equal(t, 2, len(ic.testEmits))
	assert.Contains(t, ic.testEmits, "databricks_dbfs_file[<unknown>] (id: dbfs:/FileStore/test.txt)")
	assert.Contains(t, ic.testEmits, "databricks_workspace_file[<unknown>] (id: /Shared/test.txt)")
}

func TestStorageCredentialListFails(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.1/unity-catalog/storage-credentials?",
			Status:   200,
			Response: &catalog.ListStorageCredentialsResponse{},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		err := resourcesMap["databricks_storage_credential"].List(ic)
		assert.NoError(t, err)
	})
}

func TestImportStorageCredentialGrants(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			Method:       "GET",
			Status:       200,
			Resource:     "/api/2.1/unity-catalog/permissions/storage_credential/abc",
			Response: catalog.PermissionsList{
				PrivilegeAssignments: []catalog.PrivilegeAssignment{
					{
						Principal:  "principal",
						Privileges: []catalog.Privilege{"CREATE EXTERNAL LOCATION"},
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		d := tfcatalog.ResourceStorageCredential().ToResource().TestResourceData()
		d.SetId("abc")
		err := resourcesMap["databricks_storage_credential"].Import(ic, &resource{
			ID:   "abc",
			Data: d,
		})
		assert.NoError(t, err)
	})
}

func TestExternalLocationListFails(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.1/unity-catalog/external-locations?",
			Status:   200,
			Response: &catalog.ListExternalLocationsResponse{},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		err := resourcesMap["databricks_external_location"].List(ic)
		assert.NoError(t, err)
	})
}

func TestImportExternalLocationGrants(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			Method:       "GET",
			Status:       200,
			Resource:     "/api/2.1/unity-catalog/permissions/external-locations/abc",
			Response: catalog.PermissionsList{
				PrivilegeAssignments: []catalog.PrivilegeAssignment{
					{
						Principal:  "principal",
						Privileges: []catalog.Privilege{"ALL PRIVILEGES"},
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		d := tfcatalog.ResourceExternalLocation().ToResource().TestResourceData()
		d.SetId("abc")
		err := resourcesMap["databricks_external_location"].Import(ic, &resource{
			ID:   "abc",
			Data: d,
		})
		assert.NoError(t, err)
	})
}

// TODO: restore it when support for Account-level tests is added
// func TestListMetastores(t *testing.T) {
// 	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
// 		{
// 			ReuseRequest: true,
// 			Method:       "GET",
// 			Resource:     "/api/2.1/unity-catalog/metastores",
// 			Response: catalog.ListMetastoresResponse{
// 				Metastores: []catalog.MetastoreInfo{
// 					{
// 						Name:        "test",
// 						MetastoreId: "1234",
// 					},
// 				},
// 			},
// 		},
// 	}, func(ctx context.Context, client *common.DatabricksClient) {
// 		ic := importContextForTestWithClient(ctx, client)
// 		ic.enableServices("uc-metastores")
// 		err := resourcesMap["databricks_metastore"].List(ic)
// 		assert.NoError(t, err)
// 		require.Equal(t, 1, len(ic.testEmits))
// 		assert.True(t, ic.testEmits["databricks_metastore[<unknown>] (id: 1234)"])
// 	})
// }

func TestListCatalogs(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			Method:       "GET",
			Resource:     "/api/2.1/unity-catalog/catalogs?",
			Response: catalog.ListCatalogsResponse{
				Catalogs: []catalog.CatalogInfo{
					{
						Name:        "cat1",
						CatalogType: "MANAGED_CATALOG",
					},
					{
						Name:        "cat2",
						CatalogType: "UNKNOWN",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("uc-catalogs")
		ic.currentMetastore = currentMetastoreResponse
		err := resourcesMap["databricks_catalog"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 1, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_catalog[cat1_test_MANAGED_CATALOG] (id: cat1)"])
	})
}

func TestImportManagedCatalog(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.1/unity-catalog/schemas?catalog_name=ctest",
			Response: catalog.ListSchemasResponse{
				Schemas: []catalog.SchemaInfo{
					{
						CatalogType: "MANAGED_CATALOG",
						Name:        "schema1",
						FullName:    "ctest.schema1",
					},
					{
						CatalogType: "MANAGED_CATALOG",
						Name:        "information_schema",
						FullName:    "ctest.schema1",
					},
					{
						CatalogType: "UNKNOWN",
						FullName:    "ctest.schema2",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("uc-catalogs,uc-grants,uc-schemas")
		ic.currentMetastore = currentMetastoreResponse
		d := tfcatalog.ResourceCatalog().ToResource().TestResourceData()
		d.SetId("ctest")
		d.Set("name", "ctest")
		err := resourcesMap["databricks_catalog"].Import(ic, &resource{
			ID:   "ctest",
			Data: d,
		})
		assert.NoError(t, err)
		require.Equal(t, 2, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_grants[<unknown>] (id: catalog/ctest)"])
		assert.True(t, ic.testEmits["databricks_schema[<unknown>] (id: ctest.schema1)"])
	})
}

func TestImportForeignCatalog(t *testing.T) {
	ic := importContextForTest()
	ic.enableServices("uc-catalogs,uc-grants,uc-connections")
	ic.currentMetastore = currentMetastoreResponse
	d := tfcatalog.ResourceCatalog().ToResource().TestResourceData()
	d.SetId("fctest")
	d.Set("metastore_id", "1234")
	d.Set("connection_name", "conn")
	d.Set("name", "fctest")
	err := resourcesMap["databricks_catalog"].Import(ic, &resource{
		ID:   "fctest",
		Data: d,
	})
	assert.NoError(t, err)
	require.Equal(t, 2, len(ic.testEmits))
	assert.True(t, ic.testEmits["databricks_grants[<unknown>] (id: catalog/fctest)"])
	assert.True(t, ic.testEmits["databricks_connection[<unknown>] (id: 1234|conn)"])
}

func TestImportIsolatedManagedCatalog(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.1/unity-catalog/schemas?catalog_name=ctest",
			Response: catalog.ListSchemasResponse{},
		},
		{
			Method:   "GET",
			Resource: "/api/2.1/unity-catalog/bindings/catalog/ctest?",
			Response: catalog.WorkspaceBindingsResponse{
				Bindings: []catalog.WorkspaceBinding{
					{
						BindingType: "BINDING_TYPE_READ",
						WorkspaceId: 1234,
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("uc-catalogs,uc-grants,uc-schemas")
		ic.currentMetastore = currentMetastoreResponse
		d := tfcatalog.ResourceCatalog().ToResource().TestResourceData()
		d.SetId("ctest")
		d.Set("name", "ctest")
		d.Set("isolation_mode", "ISOLATED")
		err := resourcesMap["databricks_catalog"].Import(ic, &resource{
			ID:   "ctest",
			Data: d,
		})
		assert.NoError(t, err)
		require.Equal(t, 2, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_grants[<unknown>] (id: catalog/ctest)"])
		assert.True(t, ic.testEmits["databricks_catalog_workspace_binding[catalog_ctest_ws_1234] (id: 1234|catalog|ctest)"])
	})
}

func TestImportSchema(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.1/unity-catalog/models?catalog_name=ctest&schema_name=stest",
			Response: catalog.ListRegisteredModelsResponse{
				RegisteredModels: []catalog.RegisteredModelInfo{
					{
						Name:     "model1",
						FullName: "ctest.stest.model1",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.1/unity-catalog/volumes?catalog_name=ctest&schema_name=stest",
			Response: catalog.ListVolumesResponseContent{
				Volumes: []catalog.VolumeInfo{
					{
						Name:     "volume1",
						FullName: "ctest.stest.volume1",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.1/unity-catalog/tables?catalog_name=ctest&schema_name=stest",
			Response: catalog.ListTablesResponse{
				Tables: []catalog.TableInfo{
					{
						Name:      "table1",
						TableType: "MANAGED",
						FullName:  "ctest.stest.table1",
					},
					{
						Name:      "table2",
						TableType: "UNKNOWN",
						FullName:  "ctest.stest.table2",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("uc-catalogs,uc-grants,uc-schemas,uc-volumes,uc-models,uc-tables")
		ic.currentMetastore = currentMetastoreResponse
		d := tfcatalog.ResourceSchema().ToResource().TestResourceData()
		d.SetId("ctest.stest")
		d.Set("catalog_name", "ctest")
		d.Set("name", "stest")
		err := resourcesMap["databricks_schema"].Import(ic, &resource{
			ID:   "ctest.stest",
			Data: d,
		})
		assert.NoError(t, err)
		require.Equal(t, 5, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_grants[<unknown>] (id: schema/ctest.stest)"])
		assert.True(t, ic.testEmits["databricks_catalog[<unknown>] (id: ctest)"])
		assert.True(t, ic.testEmits["databricks_registered_model[<unknown>] (id: ctest.stest.model1)"])
		assert.True(t, ic.testEmits["databricks_volume[<unknown>] (id: ctest.stest.volume1)"])
		assert.True(t, ic.testEmits["databricks_sql_table[<unknown>] (id: ctest.stest.table1)"])
	})
}

func TestImportShare(t *testing.T) {
	ic := importContextForTest()
	ic.enableServices("uc-grants,uc-volumes,uc-models,uc-tables")
	d := tfsharing.ResourceShare().ToResource().TestResourceData()
	scm := tfsharing.ResourceShare().Schema
	share := tfsharing.ShareInfo{
		Name: "stest",
		Objects: []tfsharing.SharedDataObject{
			{
				DataObjectType: "TABLE",
				Name:           "ctest.stest.table1",
			},
			{
				DataObjectType: "MODEL",
				Name:           "ctest.stest.model1",
			},
			{
				DataObjectType: "VOLUME",
				Name:           "ctest.stest.vol1",
			},
			{
				DataObjectType: "NOTEBOOK",
				Name:           "Test",
			},
		},
	}
	d.MarkNewResource()
	err := common.StructToData(share, scm, d)
	require.NoError(t, err)
	err = resourcesMap["databricks_share"].Import(ic, &resource{
		ID:   "stest",
		Data: d,
	})
	assert.NoError(t, err)
	require.Equal(t, 4, len(ic.testEmits))
	assert.True(t, ic.testEmits["databricks_grants[<unknown>] (id: share/stest)"])
	assert.True(t, ic.testEmits["databricks_registered_model[<unknown>] (id: ctest.stest.model1)"])
	assert.True(t, ic.testEmits["databricks_volume[<unknown>] (id: ctest.stest.vol1)"])
	assert.True(t, ic.testEmits["databricks_sql_table[<unknown>] (id: ctest.stest.table1)"])
}

func TestConnections(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			Method:       "GET",
			Resource:     "/api/2.1/unity-catalog/connections",
			Response: catalog.ListConnectionsResponse{
				Connections: []catalog.ConnectionInfo{
					{
						Name:        "test",
						MetastoreId: "12345",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("uc-connections,uc-grants")
		// Test Listing
		err := resourcesMap["databricks_connection"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 1, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_connection[<unknown>] (id: 12345|test)"])
		// Test Importing
		d := tfcatalog.ResourceConnection().ToResource().TestResourceData()
		d.SetId("ctest")
		d.Set("name", "ctest")
		err = resourcesMap["databricks_connection"].Import(ic, &resource{
			ID:   "ctest",
			Data: d,
		})
		assert.NoError(t, err)
		require.Equal(t, 2, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_grants[<unknown>] (id: foreign_connection/ctest)"])
	})
}

func TestListExternalLocations(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			Method:       "GET",
			Resource:     "/api/2.1/unity-catalog/external-locations?",
			Response: catalog.ListExternalLocationsResponse{
				ExternalLocations: []catalog.ExternalLocationInfo{
					{
						Name: "test",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("uc-external-locations,uc-storage-credentials,uc-grants")
		ic.currentMetastore = currentMetastoreResponse
		// Test listing
		err := resourcesMap["databricks_external_location"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 1, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_external_location[<unknown>] (id: test)"])
		// Test import
		d := tfcatalog.ResourceExternalLocation().ToResource().TestResourceData()
		d.SetId("ext_loc")
		d.Set("credential_name", "stest")
		err = resourcesMap["databricks_external_location"].Import(ic, &resource{
			ID:   "ext_loc",
			Data: d,
		})
		assert.NoError(t, err)
		require.Equal(t, 3, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_grants[<unknown>] (id: external_location/ext_loc)"])
		assert.True(t, ic.testEmits["databricks_storage_credential[<unknown>] (id: stest)"])
	})
}

func TestStorageCredentials(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			Method:       "GET",
			Resource:     "/api/2.1/unity-catalog/storage-credentials?",
			Response: catalog.ListStorageCredentialsResponse{
				StorageCredentials: []catalog.StorageCredentialInfo{
					{
						Name: "test",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("uc-storage-credentials,uc-grants")
		ic.currentMetastore = currentMetastoreResponse
		// Test listing
		err := resourcesMap["databricks_storage_credential"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 1, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_storage_credential[<unknown>] (id: test)"])
		// Test import
		err = resourcesMap["databricks_storage_credential"].Import(ic, &resource{
			ID: "1234",
		})
		assert.NoError(t, err)
		require.Equal(t, 2, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_grants[<unknown>] (id: storage_credential/1234)"])
	})
}

func TestListRecipients(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			Method:       "GET",
			Resource:     "/api/2.1/unity-catalog/recipients?",
			Response: sharing.ListRecipientsResponse{
				Recipients: []sharing.RecipientInfo{
					{
						Name: "test",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("uc-shares")
		ic.currentMetastore = currentMetastoreResponse
		err := resourcesMap["databricks_recipient"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 1, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_recipient[<unknown>] (id: test)"])
	})
}

func TestVolumes(t *testing.T) {
	ic := importContextForTest()
	ic.enableServices("uc-volumes,uc-catalogs,uc-schemas,uc-grants")
	// Test importing
	d := tfcatalog.ResourceVolume().ToResource().TestResourceData()
	d.SetId("vtest")
	d.Set("catalog_name", "ctest")
	d.Set("schema_name", "stest")
	err := resourcesMap["databricks_volume"].Import(ic, &resource{
		ID:   "vtest",
		Data: d,
	})
	assert.NoError(t, err)
	require.Equal(t, 2, len(ic.testEmits))
	assert.True(t, ic.testEmits["databricks_grants[<unknown>] (id: volume/vtest)"])
	assert.True(t, ic.testEmits["databricks_schema[<unknown>] (id: ctest.stest)"])

	//
	shouldOmitFunc := resourcesMap["databricks_volume"].ShouldOmitField
	require.NotNil(t, shouldOmitFunc)
	scm := tfcatalog.ResourceVolume().Schema
	assert.False(t, shouldOmitFunc(nil, "volume_type", scm["volume_type"], d))
	assert.False(t, shouldOmitFunc(nil, "name", scm["name"], d))
	d.Set("volume_type", "MANAGED")
	d.Set("storage_location", "s3://abc/")
	assert.False(t, shouldOmitFunc(nil, "volume_type", scm["volume_type"], d))
	assert.True(t, shouldOmitFunc(nil, "storage_location", scm["storage_location"], d))
}

func TestSqlTables(t *testing.T) {
	ic := importContextForTest()
	ic.enableServices("uc-tables,uc-catalogs,uc-schemas,uc-grants")
	// Test importing
	d := tfcatalog.ResourceSqlTable().ToResource().TestResourceData()
	d.SetId("ttest")
	d.Set("catalog_name", "ctest")
	d.Set("schema_name", "stest")
	err := resourcesMap["databricks_sql_table"].Import(ic, &resource{
		ID:   "ttest",
		Data: d,
	})
	assert.NoError(t, err)
	require.Equal(t, 2, len(ic.testEmits))
	assert.True(t, ic.testEmits["databricks_grants[<unknown>] (id: table/ttest)"])
	assert.True(t, ic.testEmits["databricks_schema[<unknown>] (id: ctest.stest)"])

	//
	shouldOmitFunc := resourcesMap["databricks_sql_table"].ShouldOmitField
	require.NotNil(t, shouldOmitFunc)
	scm := tfcatalog.ResourceSqlTable().Schema
	assert.False(t, shouldOmitFunc(nil, "table_type", scm["table_type"], d))
	assert.False(t, shouldOmitFunc(nil, "name", scm["name"], d))
	d.Set("table_type", "MANAGED")
	d.Set("storage_location", "s3://abc/")
	assert.False(t, shouldOmitFunc(nil, "table_type", scm["table_type"], d))
	assert.True(t, shouldOmitFunc(nil, "storage_location", scm["storage_location"], d))
	assert.True(t, shouldOmitFunc(nil, "storage_location", scm["storage_location"], d))
}

func TestRegisteredModels(t *testing.T) {
	ic := importContextForTest()
	ic.enableServices("uc-models,uc-catalogs,uc-schemas,uc-grants")
	// Test importing
	d := tfcatalog.ResourceRegisteredModel().ToResource().TestResourceData()
	d.SetId("mtest")
	d.Set("catalog_name", "ctest")
	d.Set("schema_name", "stest")
	err := resourcesMap["databricks_registered_model"].Import(ic, &resource{
		ID:   "mtest",
		Data: d,
	})
	assert.NoError(t, err)
	require.Equal(t, 2, len(ic.testEmits))
	assert.True(t, ic.testEmits["databricks_grants[<unknown>] (id: model/mtest)"])
	assert.True(t, ic.testEmits["databricks_schema[<unknown>] (id: ctest.stest)"])

	//
	shouldOmitFunc := resourcesMap["databricks_registered_model"].ShouldOmitField
	require.NotNil(t, shouldOmitFunc)
	scm := tfcatalog.ResourceRegisteredModel().Schema
	assert.True(t, shouldOmitFunc(nil, "storage_location", scm["storage_location"], d))
	d.Set("storage_location", "s3://abc/")
	assert.False(t, shouldOmitFunc(nil, "storage_location", scm["storage_location"], d))
	assert.False(t, shouldOmitFunc(nil, "name", scm["name"], d))

	ic.currentMetastore = currentMetastoreResponse
	d.Set("storage_location", "s3://abc/"+currentMetastoreResponse.MetastoreId+"/models/123456")
	assert.True(t, shouldOmitFunc(ic, "storage_location", scm["storage_location"], d))
}

func TestListShares(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			Method:       "GET",
			Resource:     "/api/2.1/unity-catalog/shares",
			Response: sharing.ListSharesResponse{
				Shares: []sharing.ShareInfo{
					{
						Name: "test",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("uc-shares")
		ic.currentMetastore = currentMetastoreResponse
		err := resourcesMap["databricks_share"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 1, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_share[<unknown>] (id: test)"])
	})
}

func TestAuxUcFunctions(t *testing.T) {
	// Metastore Assignment
	d := tfcatalog.ResourceMetastoreAssignment().ToResource().TestResourceData()
	d.Set("workspace_id", 123)
	assert.Equal(t, "ws_123", resourcesMap["databricks_metastore_assignment"].Name(nil, d))

	shouldOmitFunc := resourcesMap["databricks_metastore_assignment"].ShouldOmitField
	require.NotNil(t, shouldOmitFunc)
	d.Set("default_catalog_name", "")

	scm := tfcatalog.ResourceMetastoreAssignment().Schema
	assert.True(t, shouldOmitFunc(nil, "default_catalog_name", scm["default_catalog_name"], d))
	assert.False(t, shouldOmitFunc(nil, "metastore_id", scm["metastore_id"], d))

	// Metastore
	d = tfcatalog.ResourceMetastore().ToResource().TestResourceData()
	d.SetId("1234")
	assert.Equal(t, "1234", resourcesMap["databricks_metastore"].Name(nil, d))
	d.Set("name", "test")
	assert.Equal(t, "test", resourcesMap["databricks_metastore"].Name(nil, d))

	shouldOmitFunc = resourcesMap["databricks_metastore"].ShouldOmitField
	require.NotNil(t, shouldOmitFunc)
	scm = tfcatalog.ResourceMetastore().Schema
	assert.True(t, shouldOmitFunc(nil, "default_data_access_config_id", scm["default_data_access_config_id"], d))
	assert.True(t, shouldOmitFunc(nil, "owner", scm["owner"], d))
	d.Set("owner", "test")
	assert.False(t, shouldOmitFunc(nil, "owner", scm["owner"], d))
	assert.False(t, shouldOmitFunc(nil, "name", scm["name"], d))

	// Connections
	d = tfcatalog.ResourceConnection().ToResource().TestResourceData()
	d.SetId("1234")
	assert.Equal(t, "1234", resourcesMap["databricks_connection"].Name(nil, d))
	d.Set("name", "test")
	d.Set("connection_type", "db")
	assert.Equal(t, "db_test", resourcesMap["databricks_connection"].Name(nil, d))

	// Catalogs
	d = tfcatalog.ResourceCatalog().ToResource().TestResourceData()
	d.SetId("test")
	shouldOmitFunc = resourcesMap["databricks_catalog"].ShouldOmitField
	require.NotNil(t, shouldOmitFunc)
	scm = tfcatalog.ResourceCatalog().Schema
	d.Set("isolation_mode", "OPEN")
	assert.True(t, shouldOmitFunc(nil, "isolation_mode", scm["isolation_mode"], d))
	d.Set("isolation_mode", "ISOLATED")
	assert.False(t, shouldOmitFunc(nil, "isolation_mode", scm["isolation_mode"], d))
	assert.False(t, shouldOmitFunc(nil, "name", scm["name"], d))
}

func TestImportUcVolumeFile(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			Method:       "GET",
			Resource:     "/api/2.0/fs/files/Volumes/main/default/wheels/some.whl?",
			Response:     "test",
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
		defer os.RemoveAll(tmpDir)
		os.Mkdir(tmpDir, 0700)
		ic.Directory = tmpDir
		ic.enableServices("storage")
		ic.currentMetastore = currentMetastoreResponse

		file_path := "/Volumes/main/default/wheels/some.whl"
		d := storage.ResourceFile().ToResource().TestResourceData()
		d.SetId(file_path)
		err := resourcesMap["databricks_file"].Import(ic, &resource{
			ID:   file_path,
			Data: d,
		})
		assert.NoError(t, err)
		assert.Equal(t, file_path, d.Get("path"))
		assert.Equal(t, "uc_files/main/default/wheels/some.whl", d.Get("source"))
		// Testing auxiliary functions
		shouldOmitFunc := resourcesMap["databricks_file"].ShouldOmitField
		require.NotNil(t, shouldOmitFunc)
		scm := storage.ResourceFile().Schema
		assert.True(t, shouldOmitFunc(ic, "md5", scm["md5"], d))
		assert.False(t, shouldOmitFunc(ic, "path", scm["path"], d))

		assert.Equal(t, "main/default/wheels/some.whl_f27badf8", resourcesMap["databricks_file"].Name(nil, d))
	})
}

func sortStringsCopy(s []string) []string {
	c := make([]string, len(s))
	copy(c, s)
	sort.Strings(c)
	return c
}

func TestImportGrants(t *testing.T) {
	ic := importContextForTest()

	s := ic.Resources["databricks_grants"].Schema
	d := tfcatalog.ResourceGrants().ToResource().TestResourceData()
	id := "metastore/1234"
	d.SetId(id)
	d.MarkNewResource()
	r := &resource{ID: id, Data: d}
	err := resourcesMap["databricks_grants"].Import(ic, r)
	assert.NoError(t, err)

	// Test ignore function
	assert.True(t, resourcesMap["databricks_grants"].Ignore(ic, r))

	var pList tfcatalog.PermissionsList
	common.DataToStructPointer(r.Data, s, &pList)
	assert.Empty(t, pList.Assignments)

	// Test with a filled user name and no owner
	ic.meUserName = "user@domain.com"
	d.Set("catalog", "1234")
	err = resourcesMap["databricks_grants"].Import(ic, r)
	assert.NoError(t, err)
	common.DataToStructPointer(r.Data, s, &pList)
	require.Equal(t, 0, len(pList.Assignments))

	// Test with a filled user name and permissions
	r.AddExtraData("owner", "otheruser@domain.com")
	err = resourcesMap["databricks_grants"].Import(ic, r)
	assert.NoError(t, err)
	common.DataToStructPointer(r.Data, s, &pList)
	require.Equal(t, 1, len(pList.Assignments))
	assert.Equal(t, ic.meUserName, pList.Assignments[0].Principal)
	assert.Equal(t, sortStringsCopy(grantsPrivilegesToAdd["catalog"]), sortStringsCopy(pList.Assignments[0].Privileges))

	// Test with a filled user name and permissions
	d.Set("metastore", "")
	d.Set("catalog", "test")
	pList.Assignments = []tfcatalog.PrivilegeAssignment{
		{Principal: ic.meUserName, Privileges: []string{"USE_CATALOG", "USE_SCHEMA"}},
	}
	common.StructToData(pList, s, d)
	err = resourcesMap["databricks_grants"].Import(ic, r)
	assert.NoError(t, err)
	common.DataToStructPointer(r.Data, s, &pList)
	require.Equal(t, 1, len(pList.Assignments))
	assert.Equal(t, ic.meUserName, pList.Assignments[0].Principal)
	assert.Equal(t, sortStringsCopy(append([]string{"USE_CATALOG", "USE_SCHEMA"}, grantsPrivilegesToAdd["catalog"]...)),
		sortStringsCopy(pList.Assignments[0].Privileges))

	// Test with a filled user name and unsupported objects
	d.Set("catalog", "")
	d.Set("model", "test")
	err = resourcesMap["databricks_grants"].Import(ic, r)
	assert.NoError(t, err)
}
