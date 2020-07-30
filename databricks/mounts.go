package databricks

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceMounts() *schema.Resource {
	return nil // TODO: make all mounts datasource
}

// Mount exposes generic url & extra config map options
type Mount interface {
	Source() string
	Config() map[string]string
}

// MountPoint is something actionable
type MountPoint struct {
	exec      service.CommandExecutor
	clusterID string
	name      string
}

// Source ...
func (mp MountPoint) Source() (string, error) {
	return mp.exec.Execute(mp.clusterID, "python", fmt.Sprintf(`
		dbutils.fs.refreshMounts()
		for mount in dbutils.fs.mounts():
			if mount.mountPoint == "/mnt/%s":
				dbutils.notebook.exit(mount.source)
		raise Exception("Mount not found")
	`, mp.name))
}

// Delete ...
func (mp MountPoint) Delete() error {
	_, err := mp.exec.Execute(mp.clusterID, "python", fmt.Sprintf(`
		mount_point = "/mnt/%s"
		dbutils.fs.unmount(mount_point)
		dbutils.fs.refreshMounts()
		for mount in dbutils.fs.mounts():
			if mount.mountPoint == mount_point:
				raise Exception("Failed to unmount")
		dbutils.notebook.exit("success")
	`, mp.name))
	return err
}

// Mount ...
func (mp MountPoint) Mount(mo Mount) (source string, err error) {
	extraConfigs, err := json.Marshal(mo.Config())
	if err != nil {
		return
	}
	b := regexp.MustCompile(`"\{secrets/([^/]+)/([^\}]+)\}"`)
	extraConfigs = b.ReplaceAll(extraConfigs, []byte(`dbutils.secrets.get("$1", "$2")`))
	source, err = mp.exec.Execute(mp.clusterID, "python", fmt.Sprintf(`
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
		mount_source = safe_mount("/mnt/%s", "%v", %s)
		dbutils.notebook.exit(mount_source)
	`, mp.name, mo.Source(), extraConfigs))
	return
}

func mountCluster(mf func() Mount, d *schema.ResourceData, m interface{},
	r *schema.Resource) (mo Mount, mp MountPoint, err error) {
	mo = mf()
	client := m.(*service.DatabricksClient)
	err = readStructFromData([]string{}, d, &mo, r)
	if err != nil {
		return
	}
	name := d.Get("mount_name").(string)
	d.SetId(name)
	clusterID := d.Get("cluster_id").(string)
	if clusterID == "" {
		cluster, err2 := client.Clusters().GetOrCreateRunningCluster("terraform-mount")
		if err2 != nil {
			err = err2
			return
		}
		clusterID = cluster.ClusterID
	} else {
		err = changeClusterIntoRunningState(clusterID, client)
		if isClusterMissing(err, clusterID) {
			log.Printf("[WANR] Unable to get mount %s state on %s", name, clusterID)
			return
		}
		if err != nil {
			return
		}
	}
	mp = MountPoint{
		exec:      client.Commands(),
		clusterID: clusterID,
		name:      name,
	}
	return
}

func mountCreate(mf func() Mount, r *schema.Resource) schema.CreateFunc {
	return func(d *schema.ResourceData, m interface{}) (err error) {
		mo, mp, err := mountCluster(mf, d, m, r)
		if err != nil {
			return
		}
		source, err := mp.Mount(mo)
		if err != nil {
			return
		}
		return d.Set("source", source)
	}
}

func mountRead(mf func() Mount, r *schema.Resource) schema.ReadFunc {
	return func(d *schema.ResourceData, m interface{}) (err error) {
		_, mp, err := mountCluster(mf, d, m, r)
		if err != nil {
			return
		}
		source, err := mp.Source()
		if err == errors.New("Mount not found") {
			// remove resource
			d.SetId("")
			return nil
		}
		if err != nil {
			return
		}
		return d.Set("source", source)
	}
}

func mountDelete(mf func() Mount, r *schema.Resource) schema.DeleteFunc {
	return func(d *schema.ResourceData, m interface{}) (err error) {
		_, mp, err := mountCluster(mf, d, m, r)
		if err != nil {
			return
		}
		return mp.Delete()
	}
}

// ValidateMountDirectory is a ValidateFunc that ensures the mount directory starts with a '/'
func ValidateMountDirectory(val interface{}, key string) (warns []string, errs []error) {
	v := val.(string)
	if v != "" && !strings.HasPrefix(v, "/") {
		return nil, []error{fmt.Errorf("%s must start with /, got: %s", key, v)}
	}
	return nil, nil
}
