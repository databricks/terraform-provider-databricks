package storage

import (
	"fmt"
	"testing"

	"github.com/databrickslabs/databricks-terraform/internal"

	"github.com/databrickslabs/databricks-terraform/common"
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

const expectedCommandResp = "done"

func testMountFuncHelper(t *testing.T, mountFunc func(mp MountPoint, mount Mount) (string, error), mount Mount,
	mountName, expectedCommand string) {
	c := common.DatabricksClient{
		Host:  ".",
		Token: ".",
	}
	err := c.Configure()
	assert.NoError(t, err)

	var called bool

	c.WithCommandMock(func(commandStr string) (s string, e error) {
		called = true
		assert.Equal(t, internal.TrimLeadingWhitespace(expectedCommand), internal.TrimLeadingWhitespace(commandStr))
		return expectedCommandResp, nil
	})

	mp := MountPoint{
		exec:      c.CommandExecutor(),
		clusterID: "random_cluster_id",
		name:      mountName,
	}

	resp, err := mountFunc(mp, mount)
	assert.NoError(t, err)
	assert.True(t, called, "mocked command was not invoked")
	assert.Equal(t, expectedCommandResp, resp)
}

type mockMount struct{}

func (t mockMount) Source() string            { return "fake-mount" }
func (t mockMount) Config() map[string]string { return map[string]string{"fake-key": "fake-value"} }

func TestMountPoint_Mount(t *testing.T) {
	mount := mockMount{}
	expectedMountSource := "fake-mount"
	expectedMountConfig := `{"fake-key":"fake-value"}`
	mountName := "this_mount"
	expectedCommand := fmt.Sprintf(`
		def safe_mount(mount_point, mount_source, configs):
			for mount in dbutils.fs.mounts():
				if mount.mountPoint == mount_point and mount.source == mount_source:
					return
			try:
				dbutils.fs.mount(mount_source, mount_point, extra_configs=configs)
				dbutils.fs.refreshMounts()
				dbutils.fs.ls(mount_point)
				return mount_source
			except Exception as e:
				try:
					dbutils.fs.unmount(mount_point)
				except Exception as e2:
					print("Failed to unmount", e2)
				raise e
		mount_source = safe_mount("/mnt/%s", %q, %s)
		dbutils.notebook.exit(mount_source)
	`, mountName, expectedMountSource, expectedMountConfig)
	testMountFuncHelper(t, func(mp MountPoint, mount Mount) (s string, e error) {
		return mp.Mount(mount)
	}, mount, mountName, expectedCommand)
}

func TestMountPoint_Source(t *testing.T) {
	mountName := "this_mount"
	expectedCommand := fmt.Sprintf(`
		dbutils.fs.refreshMounts()
		for mount in dbutils.fs.mounts():
			if mount.mountPoint == "/mnt/%s":
				dbutils.notebook.exit(mount.source)
		raise Exception("Mount not found")
	`, mountName)
	testMountFuncHelper(t, func(mp MountPoint, mount Mount) (s string, e error) {
		return mp.Source()
	}, nil, mountName, expectedCommand)
}

func TestMountPoint_Delete(t *testing.T) {
	mountName := "this_mount"
	expectedCommand := fmt.Sprintf(`
		mount_point = "/mnt/%s"
		dbutils.fs.unmount(mount_point)
		dbutils.fs.refreshMounts()
		for mount in dbutils.fs.mounts():
			if mount.mountPoint == mount_point:
				raise Exception("Failed to unmount")
		dbutils.notebook.exit("success")
	`, mountName)
	testMountFuncHelper(t, func(mp MountPoint, mount Mount) (s string, e error) {
		return expectedCommandResp, mp.Delete()
	}, nil, mountName, expectedCommand)
}
