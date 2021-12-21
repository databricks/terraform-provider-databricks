package exporter

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/clusters"
	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/jobs"
	"github.com/databrickslabs/terraform-provider-databricks/permissions"
	"github.com/databrickslabs/terraform-provider-databricks/policies"
	"github.com/databrickslabs/terraform-provider-databricks/pools"
	"github.com/databrickslabs/terraform-provider-databricks/provider"
	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/databrickslabs/terraform-provider-databricks/scim"
	"github.com/databrickslabs/terraform-provider-databricks/secrets"
	"github.com/databrickslabs/terraform-provider-databricks/storage"
	"github.com/stretchr/testify/assert"
)

func importContextForTest() *importContext {
	p := provider.DatabricksProvider()
	return &importContext{
		Importables: resourcesMap,
		Resources:   p.ResourcesMap,
		testEmits:   map[string]bool{},
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
	assert.True(t, ic.testEmits["databricks_group_member[_foo] (id: parent-group|123)"])
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

func TestDbfsFileCornerCases_ReadFail(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/dbfs/read?length=1000000&path=a",
			Status:   404,
			Response: common.NotFound("nope"),
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTest()
		ic.Client = client
		ic.Context = ctx
		err := resourcesMap["databricks_dbfs_file"].Body(ic, nil, &resource{
			ID: "a",
		})
		assert.EqualError(t, err, "nope")
	})
}

func TestDbfsFileCornerCases_WriteWrongDir(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/dbfs/read?length=1000000&path=a",
			Response: storage.ReadResponse{
				Data:      "YWJj",
				BytesRead: 3,
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTest()
		ic.Client = client
		ic.Context = ctx
		err := resourcesMap["databricks_dbfs_file"].Body(ic, nil, &resource{
			ID: "a",
		})
		assert.NotNil(t, err) // mustn't match direct OS error
	})
}

func TestDbfsFileCornerCases_WriteFileExists(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/dbfs/read?length=1000000&path=a",
			Response: storage.ReadResponse{
				Data:      "YWJj",
				BytesRead: 3,
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTest()
		ic.Client = client
		ic.Context = ctx
		ic.Directory = fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
		defer os.RemoveAll(ic.Directory)

		dstFile := fmt.Sprintf("%s/files/_abc_900150983cd24fb0d6963f7d28e17f72", ic.Directory)
		err := os.MkdirAll(dstFile, 0755)
		assert.NoError(t, err)

		d := storage.ResourceDBFSFile().TestResourceData()
		d.SetId("abc")

		err = resourcesMap["databricks_dbfs_file"].Body(ic, nil, &resource{
			ID:   "a",
			Data: d,
		})
		assert.Equal(t, err.Error(), fmt.Sprintf("open %s: is a directory", dstFile))
	})
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
		assert.Equal(t, 0, len(ic.importing))
	})
}

func TestJobListNoNameMatch(t *testing.T) {
	ic := importContextForTest()
	ic.match = "bcd"
	ic.importJobs(jobs.JobList{
		Jobs: []jobs.Job{
			{
				Settings: &jobs.JobSettings{
					Name: "abc",
				},
			},
		},
	})
	assert.Equal(t, 0, len(ic.importing))
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
		ic.importJobs(jobs.JobList{
			Jobs: []jobs.Job{
				{
					JobID: 1,
					Settings: &jobs.JobSettings{
						Name: "abc",
					},
				},
			},
		})
		assert.Equal(t, 0, len(ic.importing))
	})
}
