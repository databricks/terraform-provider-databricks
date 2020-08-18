package storage

import (
	"fmt"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"

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

var expectedCommandResp = "done"

func testMountFuncHelper(t *testing.T, mountName, expectedCommand string, mount Mount,
	mountFunc func(mp MountPoint, mount Mount) (string, error)) {
	c := common.DatabricksClient{
		Host:  ".",
		Token: ".",
	}
	err := c.Configure()
	assert.NoError(t, err)

	var called bool

	c.WithCommandMock(func(commandStr string) (s string, e error) {
		called = true
		assert.Equal(t, expectedCommand, commandStr)
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

func testMountPointCreateHelper(t *testing.T, mount Mount, expectedMountSource, expectedMountConfig string) {
	randomMountName := qa.RandomName()
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
	`, randomMountName, expectedMountSource, expectedMountConfig)
	testMountFuncHelper(t, randomMountName, expectedCommand, mount, func(mp MountPoint, mount Mount) (s string, e error) {
		return mp.Mount(mount)
	})
}

func TestMountPoint_Source(t *testing.T) {
	randomMountName := qa.RandomName()
	expectedCommand := fmt.Sprintf(`
		dbutils.fs.refreshMounts()
		for mount in dbutils.fs.mounts():
			if mount.mountPoint == "/mnt/%s":
				dbutils.notebook.exit(mount.source)
		raise Exception("Mount not found")
	`, randomMountName)
	testMountFuncHelper(t, randomMountName, expectedCommand, nil,
		func(mp MountPoint, mount Mount) (s string, e error) {
			return mp.Source()
		})
}

func TestMountPoint_Delete(t *testing.T) {
	randomMountName := qa.RandomName()
	expectedCommand := fmt.Sprintf(`
		mount_point = "/mnt/%s"
		dbutils.fs.unmount(mount_point)
		dbutils.fs.refreshMounts()
		for mount in dbutils.fs.mounts():
			if mount.mountPoint == mount_point:
				raise Exception("Failed to unmount")
		dbutils.notebook.exit("success")
	`, randomMountName)
	testMountFuncHelper(t, randomMountName, expectedCommand, nil,
		func(mp MountPoint, mount Mount) (s string, e error) {
			return expectedCommandResp, mp.Delete()
		})
}
