package storage

import (
	"context"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/clusters"
	"github.com/databrickslabs/terraform-provider-databricks/internal"

	"github.com/databrickslabs/terraform-provider-databricks/qa"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/stretchr/testify/assert"
)

func TestValidateMountDirectory(t *testing.T) {
	testCases := []struct {
		directory  string
		errorCount int
	}{
		{"", 0},
		{"/directory", 0},
		{"directory", 1},
	}
	for _, tc := range testCases {
		_, errs := ValidateMountDirectory(tc.directory, "key")

		assert.Lenf(t, errs, tc.errorCount, "directory '%s' does not generate the expected error count", tc.directory)
	}
}

func testMountFuncHelper(t *testing.T, mountFunc func(mp MountPoint, mount Mount) (string, error), mount Mount,
	mountName, expectedCommand string) {
	c := common.DatabricksClient{
		Host:  ".",
		Token: ".",
	}
	err := c.Configure()
	assert.NoError(t, err)

	var called bool

	c.WithCommandMock(func(commandStr string) common.CommandResults {
		called = true
		assert.Equal(t, internal.TrimLeadingWhitespace(expectedCommand), internal.TrimLeadingWhitespace(commandStr))
		return common.CommandResults{
			ResultType: "text",
			Data:       `{"source":"done", "config_hash":"abc123"}`,
		}
	})

	ctx := context.Background()
	mp := MountPoint{
		Exec:      c.CommandExecutor(ctx),
		ClusterID: "random_cluster_id",
		Name:      mountName,
	}

	resp, err := mountFunc(mp, mount)
	assert.NoError(t, err)
	assert.True(t, called, "mocked command was not invoked")
	assert.Equal(t, "done", resp)
}

type mockMount struct{}

func (t mockMount) Source() string { return "dummy://things" }
func (t mockMount) Name() string   { return "fake-mount" }
func (t mockMount) Config(client *common.DatabricksClient) map[string]string {
	return map[string]string{"creds": "{{secrets/scope/key}}"}
}
func (m mockMount) ValidateAndApplyDefaults(d *schema.ResourceData, client *common.DatabricksClient) error {
	return nil
}

func TestMountPoint_Mount(t *testing.T) {
	testMountFuncHelper(t, func(mp MountPoint, mount Mount) (string, error) {
		info, err := mp.Mount(mount, &common.DatabricksClient{
			Host:  ".",
			Token: ".",
		})
		return info.Source, err
	}, mockMount{}, "dummy", `
		import json, hashlib
		def safe_mount(mount_point, mount_source, configs, encryptionType):
			for mount in dbutils.fs.mounts():
				if mount.mountPoint == mount_point and mount.source == mount_source:
					return
			try:
				dbutils.fs.mount(mount_source, mount_point, 
					extra_configs=configs, 
					encryption_type=encryptionType)
				dbutils.fs.refreshMounts()
				dbutils.fs.ls(mount_point)
				return mount_source
			except Exception as e:
				try:
					dbutils.fs.unmount(mount_point)
				except Exception as e2:
					print("Failed to unmount", e2)
				raise e
		extra_configs = {"creds":dbutils.secrets.get("scope", "key")}
		mount_source = safe_mount("/mnt/dummy", "dummy://things", extra_configs, "")
		dbutils.notebook.exit(json.dumps({
			"source": mount_source,
			"config_hash": `+configHashCode+`
		}))
	`)
}

func TestMountPoint_Source(t *testing.T) {
	testMountFuncHelper(t, func(mp MountPoint, mount Mount) (string, error) {
		i, err := mp.Source(mount, &common.DatabricksClient{
			Host:  ".",
			Token: ".",
		})
		return i.Source, err
	}, mockMount{}, "dummy", `
		import json, hashlib
		dbutils.fs.refreshMounts()
		extra_configs = {"creds":dbutils.secrets.get("scope", "key")}
		for mount in dbutils.fs.mounts():
			if mount.mountPoint == "/mnt/dummy":
				dbutils.notebook.exit(json.dumps({
					"source": mount.source,
					"config_hash": `+configHashCode+`
				}))
		raise Exception("Mount not found")
	`)
}

func TestMountPoint_Update(t *testing.T) {
	testMountFuncHelper(t, func(mp MountPoint, mount Mount) (s string, e error) {
		i, err := mp.Update(mount, &common.DatabricksClient{
			Host:  ".",
			Token: ".",
		})
		return i.Source, err
	}, mockMount{}, "dummy", `
		extra_configs = {"creds":dbutils.secrets.get("scope", "key")}
		dbutils.fs.updateMount(
			mount_point = "/mnt/dummy",
			source = "dummy://things",
			extra_configs = extra_configs)
		dbutils.notebook.exit(json.dumps({
			"source": mount_source,
			"config_hash": `+configHashCode+`
		}))
	`)
}

func TestMountPoint_Delete(t *testing.T) {
	testMountFuncHelper(t, func(mp MountPoint, mount Mount) (s string, e error) {
		return "done", mp.Delete()
	}, mockMount{}, "dummy", `
		found = False
		mount_point = "/mnt/dummy"
		dbutils.fs.refreshMounts()
		for mount in dbutils.fs.mounts():
			if mount.mountPoint == mount_point:
				found = True
		if not found:
			dbutils.notebook.exit("success")
		dbutils.fs.unmount(mount_point)
		dbutils.fs.refreshMounts()
		for mount in dbutils.fs.mounts():
			if mount.mountPoint == mount_point:
				raise Exception("Failed to unmount")
		dbutils.notebook.exit("success")
	`)
}

func TestDeletedMountClusterRecreates(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Status:   404,
		},
		{
			Method:       "GET",
			ReuseRequest: true,
			Resource:     "/api/2.0/clusters/list",
			Response:     map[string]interface{}{},
		},
		{
			Method:       "GET",
			ReuseRequest: true,
			Resource:     "/api/2.0/clusters/spark-versions",
			Response: clusters.SparkVersionsList{
				SparkVersions: []clusters.SparkVersion{
					{
						Version:     "7.1.x-cpu-ml-scala2.12",
						Description: "7.1 ML (includes Apache Spark 3.0.0, Scala 2.12)",
					},
				},
			},
		},
		{
			Method:       "GET",
			ReuseRequest: true,
			Resource:     "/api/2.0/clusters/list-node-types",
			Response: clusters.NodeTypeList{
				NodeTypes: []clusters.NodeType{
					{
						NodeTypeID:     "Standard_F4s",
						InstanceTypeID: "Standard_F4s",
						MemoryMB:       8192,
						NumCores:       4,
						NodeInstanceType: &clusters.NodeInstanceType{
							LocalDisks:      1,
							InstanceTypeID:  "Standard_F4s",
							LocalDiskSizeGB: 16,
							LocalNVMeDisks:  0,
						},
					},
				},
			},
		},
		{
			Method:       "POST",
			ReuseRequest: true,
			Resource:     "/api/2.0/clusters/create",
			ExpectedRequest: clusters.Cluster{
				AutoterminationMinutes: 10,
				ClusterName:            "terraform-mount",
				NodeTypeID:             "Standard_F4s",
				SparkVersion:           "7.3.x-scala2.12",
				CustomTags: map[string]string{
					"ResourceClass": "SingleNode",
				},
				SparkConf: map[string]string{
					"spark.databricks.cluster.profile": "singleNode",
					"spark.master":                     "local[*]",
					"spark.scheduler.mode":             "FIFO",
				},
			},
			Response: clusters.ClusterID{
				ClusterID: "bcd",
			},
		},
		{
			Method:       "GET",
			ReuseRequest: true,
			Resource:     "/api/2.0/clusters/get?cluster_id=bcd",
			Response: clusters.ClusterInfo{
				ClusterID: "bcd",
				State:     "RUNNING",
				SparkConf: map[string]string{
					"spark.databricks.acl.dfAclsEnabled": "true",
					"spark.databricks.cluster.profile":   "singleNode",
					"spark.scheduler.mode":               "FIFO",
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		clusterID, err := getMountingClusterID(ctx, client, "abc")
		assert.NoError(t, err)
		assert.Equal(t, "bcd", clusterID)
	})
}

func TestOldMountImplementations(t *testing.T) {
	n := "test"
	m1 := AzureADLSGen2Mount{ContainerName: n}
	assert.Equal(t, m1.Name(), n)
	assert.Nil(t, m1.ValidateAndApplyDefaults(nil, nil))

	m2 := AzureBlobMount{ContainerName: n}
	assert.Equal(t, m2.Name(), n)
	assert.Nil(t, m2.ValidateAndApplyDefaults(nil, nil))

	m3 := AzureADLSGen1Mount{StorageResource: n}
	assert.Equal(t, m3.Name(), n)
	assert.Nil(t, m3.ValidateAndApplyDefaults(nil, nil))

	m4 := AWSIamMount{S3BucketName: n}
	assert.Equal(t, m4.Name(), n)
	assert.Nil(t, m4.ValidateAndApplyDefaults(nil, nil))

}

type sampleCommand string

func (s sampleCommand) Execute(clusterID, language, commandStr string) common.CommandResults {
	return common.CommandResults{
		ResultType: "text",
		Data:       string(s),
	}
}

func TestNewMountPoint(t *testing.T) {
	mp := NewMountPoint(sampleCommand("abc"), "abc", "bcd")
	r := mp.Exec.Execute("a", "b", "c")
	assert.Equal(t, "abc", (&r).Text())
}

func TestGetMountingClusterID_Failures(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=def",
			Status:   518,
			Response: common.APIError{
				Message: "😤",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=terminated",
			Response: clusters.ClusterInfo{
				ClusterID: "failing-id",
				State:     clusters.ClusterStateTerminated,
			},
		},
		{
			MatchAny:     true,
			ReuseRequest: true,
			Status:       404,
			Response:     common.NotFound("nope"),
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		// no mounting cluster given, try creating it
		_, err := getMountingClusterID(ctx, client, "")
		assert.EqualError(t, err, "failed to get mouting cluster: nope")

		// mounting cluster given, but it's removed already
		_, err = getMountingClusterID(ctx, client, "bcd")
		assert.EqualError(t, err, "failed to get mouting cluster: nope")

		// some other error happens
		_, err = getMountingClusterID(ctx, client, "def")
		assert.EqualError(t, err, "failed to re-create mounting cluster: 😤")

		_, err = getMountingClusterID(ctx, client, "terminated")
		assert.EqualError(t, err, "failed to start mounting cluster: nope")
	})
}

func TestMountCRD(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			MatchAny:     true,
			ReuseRequest: true,
			Status:       404,
			Response:     common.NotFound("nope"),
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		r := ResourceMount()
		d := r.TestResourceData()
		client.WithCommandMock(func(commandStr string) common.CommandResults {
			return common.CommandResults{}
		})
		diags := mountCreate(nil, r)(ctx, d, client)
		assert.True(t, diags.HasError())
		assert.Equal(t, "failed to get mouting cluster: nope", diags[0].Summary)

		diags = mountRead(nil, r)(ctx, d, client)
		assert.True(t, diags.HasError())
		assert.Equal(t, "failed to get mouting cluster: nope", diags[0].Summary)

		diags = mountDelete(nil, r)(ctx, d, client)
		assert.True(t, diags.HasError())
		assert.Equal(t, "failed to get mouting cluster: nope", diags[0].Summary)
	})
}
