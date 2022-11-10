package exporter

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/commands"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/jobs"
	"github.com/databricks/terraform-provider-databricks/permissions"
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
		Importables: resourcesMap,
		Resources:   p.ResourcesMap,
		Files:       map[string]*hclwrite.File{},
		testEmits:   map[string]bool{},
		nameFixes:   nameFixes,
	}
}

func TestInstancePool(t *testing.T) {
	d := pools.ResourceInstancePool().TestResourceData()
	d.Set("instance_pool_name", "blah-bah")
	name := resourcesMap["databricks_instance_pool"].Name(d)
	assert.Equal(t, "blah-bah", name)

	d = pools.ResourceInstancePool().TestResourceData()
	d.SetId("abc-bcd-def")
	name = resourcesMap["databricks_instance_pool"].Name(d)
	assert.Equal(t, "def", name)

	ic := importContextForTest()
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
	}
	policy, _ := json.Marshal(definition)
	d.Set("definition", string(policy))
	ic := importContextForTest()
	err := ic.Importables["databricks_cluster_policy"].Import(ic, &resource{
		ID:   "abc",
		Data: d,
	})
	assert.NoError(t, err)
	assert.Len(t, ic.testEmits, 3)
	assert.True(t, ic.testEmits["databricks_permissions[clust_policy_bcd] (id: /cluster-policies/abc)"])
	assert.True(t, ic.testEmits["databricks_instance_pool[<unknown>] (id: efg)"])
	assert.True(t, ic.testEmits["databricks_instance_profile[<unknown>] (id: def)"])
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
	r := &resource{
		Value:     "foo",
		Attribute: "display_name",
	}
	err := ic.Importables["databricks_group"].Search(ic, r)
	assert.NoError(t, err)
	assert.Equal(t, "123", r.ID)

	err = ic.Importables["databricks_group"].Import(ic, r)
	assert.NoError(t, err)
	assert.Len(t, ic.testEmits, 4)
	assert.True(t, ic.testEmits["databricks_group_instance_profile[<unknown>] (id: 123|abc)"])
	assert.True(t, ic.testEmits["databricks_instance_profile[<unknown>] (id: abc)"])
	assert.True(t, ic.testEmits["databricks_group[<unknown>] (id: parent-group)"])
	assert.True(t, ic.testEmits["databricks_group_member[_parent-group_foo] (id: parent-group|123)"])
}

func TestPermissions(t *testing.T) {
	p := permissions.ResourcePermissions()
	d := p.TestResourceData()
	d.SetId("abc")
	ic := importContextForTest()
	name := ic.Importables["databricks_permissions"].Name(d)
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
		},
	}, p.Schema, d)
	assert.NoError(t, err)

	err = ic.Importables["databricks_permissions"].Import(ic, &resource{
		Data: d,
	})
	assert.NoError(t, err)
	assert.Len(t, ic.testEmits, 2)
	assert.True(t, ic.testEmits["databricks_user[<unknown>] (user_name: a)"])
	assert.True(t, ic.testEmits["databricks_group[<unknown>] (display_name: b)"])
}

func TestSecretScope(t *testing.T) {
	d := secrets.ResourceSecretScope().TestResourceData()
	d.Set("name", "abc")
	ic := importContextForTest()
	name := ic.Importables["databricks_secret_scope"].Name(d)
	assert.Equal(t, "abc", name)
}

func TestInstancePoolNameFromID(t *testing.T) {
	d := pools.ResourceInstancePool().TestResourceData()
	d.SetId("a-b-c")
	d.Set("instance_pool_name", "")
	assert.Equal(t, "c", resourcesMap["databricks_instance_pool"].Name(d))
}

func TestClusterNameFromID(t *testing.T) {
	d := clusters.ResourceCluster().TestResourceData()
	d.SetId("a-b-c")
	assert.Equal(t, "c", resourcesMap["databricks_cluster"].Name(d))
}

func TestClusterListFails(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/list",
			Status:   404,
			Response: common.NotFound("nope"),
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTest()
		ic.Client = client
		ic.Context = ctx
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
		ic := importContextForTest()
		ic.Client = client
		ic.Context = ctx
		ic.match = "bcd"
		err := resourcesMap["databricks_cluster"].List(ic)
		assert.NoError(t, err)
		assert.Equal(t, 0, len(ic.testEmits))
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

func TestJobList_FailGetRuns(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/jobs/runs/list?completed_only=true&job_id=1&limit=1",
			Status:   404,
			Response: common.NotFound("nope"),
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTest()
		ic.Client = client
		ic.Context = ctx
		ic.importJobs([]jobs.Job{
			{
				JobID: 1,
				Settings: &jobs.JobSettings{
					Name: "abc",
				},
			},
		})
		assert.Equal(t, 0, len(ic.testEmits))
	})
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
			Resource:     "/api/2.0/preview/scim/v2/Groups?",
			Status:       404,
			Response:     common.NotFound("nope"),
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTest()
		ic.Client = client
		ic.Context = ctx
		err := resourcesMap["databricks_group"].List(ic)
		assert.EqualError(t, err, "nope")

		err = resourcesMap["databricks_group"].Search(ic, &resource{
			ID: "nonsense",
		})
		assert.EqualError(t, err, "nope")

		err = resourcesMap["databricks_group"].Import(ic, &resource{
			ID: "nonsense",
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
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			Method:       "GET",
			Resource:     "/api/2.0/preview/scim/v2/Users?filter=userName%20eq%20%27dbc%27",
			Status:       404,
			Response:     common.NotFound("nope"),
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTest()
		ic.Client = client
		ic.Context = ctx

		d := scim.ResourceUser().TestResourceData()
		d.Set("user_name", "dbc")
		r := &resource{
			Attribute: "display_name",
			Value:     "dbc",
			Data:      d,
		}
		err := resourcesMap["databricks_user"].Search(ic, r)
		assert.EqualError(t, err, "nope")

		err = resourcesMap["databricks_user"].Import(ic, r)
		assert.EqualError(t, err, "nope")
	})
}

func TestUserImportSkipNonDirectGroups(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			Method:       "GET",
			Resource:     "/api/2.0/preview/scim/v2/Users?filter=userName%20eq%20%27dbc%27",
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
		ic := importContextForTest()
		ic.Client = client
		ic.Context = ctx

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
		ic := importContextForTest()
		ic.Client = client
		ic.Context = ctx
		ic.match = "bcd"
		err := resourcesMap["databricks_secret_scope"].List(ic)
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

func TestGlobalInitScriptNameFromId(t *testing.T) {
	d := workspace.ResourceGlobalInitScript().TestResourceData()
	d.SetId("abc")
	assert.Equal(t, "abc", resourcesMap["databricks_global_init_script"].Name(d))
}

func TestGlobalInitScriptsErrors(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			MatchAny:     true,
			Status:       404,
			Response:     common.NotFound("nope"),
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTest()
		ic.Client = client
		ic.Context = ctx
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
		ic := importContextForTest()
		ic.Client = client
		ic.Context = ctx
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

func TestRepoIdForName(t *testing.T) {
	d := repos.ResourceRepo().TestResourceData()
	d.SetId("x")
	assert.Equal(t, "x", resourcesMap["databricks_repo"].Name(d))
}

func TestRepoListFails(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			MatchAny:     true,
			Status:       404,
			Response:     common.NotFound("nope"),
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTest()
		ic.Client = client
		ic.Context = ctx
		err := resourcesMap["databricks_repo"].List(ic)
		assert.EqualError(t, err, "nope")
	})
}

func testGenerate(t *testing.T, fixtures []qa.HTTPFixture, services string, cb func(*importContext)) {
	qa.HTTPFixturesApply(t, fixtures, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTest()
		ic.Directory = fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
		defer os.RemoveAll(ic.Directory)
		ic.Client = client
		ic.Context = ctx
		ic.testEmits = nil
		ic.importing = map[string]bool{}
		ic.variables = map[string]string{}
		ic.services = services
		cb(ic)
	})
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
	}, "notebooks", func(ic *importContext) {
		err := resourcesMap["databricks_notebook"].List(ic)
		assert.NoError(t, err)

		ic.generateHclForResources(nil)
		assert.Equal(t, commands.TrimLeadingWhitespace(`
		resource "databricks_notebook" "first_second_123" {
		  source = "${path.module}/notebooks/First/Second.py"
		  path   = "/First/Second"
		}`), string(ic.Files["notebooks"].Bytes()))
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
	}, "workspace", func(ic *importContext) {
		ic.Emit(&resource{
			Resource: "databricks_global_init_script",
			ID:       "a",
		})

		ic.generateHclForResources(nil)
		assert.Equal(t, commands.TrimLeadingWhitespace(`
		resource "databricks_global_init_script" "new_importing_things" {
		  source  = "${path.module}/files/new_importing_things.sh"
		  name    = "New: Importing ^ Things"
		  enabled = true
		}`), string(ic.Files["workspace"].Bytes()))
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
	}, "secrets", func(ic *importContext) {
		ic.Emit(&resource{
			Resource: "databricks_secret",
			ID:       "a|||b",
		})

		ic.generateHclForResources(nil)
		assert.Equal(t, commands.TrimLeadingWhitespace(`
		resource "databricks_secret" "a_b" {
		  string_value = var.string_value_a_b
		  scope        = "a"
		  key          = "b"
		}`), string(ic.Files["secrets"].Bytes()))
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
	}, "storage", func(ic *importContext) {
		ic.Emit(&resource{
			Resource: "databricks_dbfs_file",
			ID:       "a",
		})

		ic.generateHclForResources(nil)
		assert.Equal(t, commands.TrimLeadingWhitespace(`
		resource "databricks_dbfs_file" "_a_0cc175b9c0f1b6a831c399e269772661" {
		  source = "${path.module}/files/_a_0cc175b9c0f1b6a831c399e269772661"
		  path   = "a"
		}`), string(ic.Files["storage"].Bytes()))
	})
}

func TestSqlListObjects(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/sql/queries",
			Response: dbsqlListResponse{PageSize: 1, Page: 1, TotalCount: 2,
				Results: []map[string]any{{"key1": "value1"}}},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/sql/queries?page=2&page_size=1",
			Response: dbsqlListResponse{PageSize: 1, Page: 2, TotalCount: 2,
				Results: []map[string]any{{"key2": "value2"}}},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTest()
		ic.Client = client
		ic.Context = ctx
		answer, err := dbsqlListObjects(ic, "/preview/sql/queries")
		assert.NoError(t, err)
		assert.Len(t, answer, 2)
	})
}
