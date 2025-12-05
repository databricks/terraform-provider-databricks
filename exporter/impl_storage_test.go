package exporter

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/commands"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/databricks/terraform-provider-databricks/storage"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
)

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
