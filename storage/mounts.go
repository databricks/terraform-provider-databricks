package storage

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/compute"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Mount exposes generic url & extra config map options
type Mount interface {
	Source() string
	Config() map[string]string
}

// MountPoint is something actionable
type MountPoint struct {
	exec      common.CommandExecutor
	clusterID string
	name      string
}

// Source returns mountpoint source
func (mp MountPoint) Source() (string, error) {
	return mp.exec.Execute(mp.clusterID, "python", fmt.Sprintf(`
		dbutils.fs.refreshMounts()
		for mount in dbutils.fs.mounts():
			if mount.mountPoint == "/mnt/%s":
				dbutils.notebook.exit(mount.source)
		raise Exception("Mount not found")
	`, mp.name))
}

// Delete removes mount from workspace
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

// Mount mounts object store on workspace
func (mp MountPoint) Mount(mo Mount) (source string, err error) {
	extraConfigs, err := json.Marshal(mo.Config())
	if err != nil {
		return
	}
	b := regexp.MustCompile(`"\{secrets/([^/]+)/([^\}]+)\}"`)
	extraConfigs = b.ReplaceAll(extraConfigs, []byte(`dbutils.secrets.get("$1", "$2")`))
	command := fmt.Sprintf(`
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
	`, mp.name, mo.Source(), extraConfigs)
	source, err = mp.exec.Execute(mp.clusterID, "python", command)
	return
}

func commonMountResource(tpl Mount, s map[string]*schema.Schema) *schema.Resource {
	resource := &schema.Resource{Schema: s, SchemaVersion: 2}
	// nolint should be a bigger context-aware refactor
	resource.Create = mountCreate(tpl, resource)
	resource.Read = mountRead(tpl, resource)
	resource.Delete = mountDelete(tpl, resource)
	return resource
}

// NewMountPoint returns new mount point config
func NewMountPoint(client *common.DatabricksClient, name, clusterID string) MountPoint {
	executor := client.CommandExecutor()
	return MountPoint{
		// todo: fix
		exec:      executor,
		clusterID: clusterID,
		name:      name,
	}
}

func getMountingClusterID(client *common.DatabricksClient, clusterID string) (string, error) {
	if clusterID == "" {
		cluster, err := compute.NewClustersAPI(client).GetOrCreateRunningCluster("terraform-mount")
		if err != nil {
			return "", err
		}
		return cluster.ClusterID, nil
	}
	clusterInfo, err := compute.NewClustersAPI(client).Get(clusterID)
	if err != nil {
		return "", err
	}
	if !clusterInfo.IsRunningOrResizing() {
		err = compute.NewClustersAPI(client).Start(clusterInfo.ClusterID)
		if err != nil {
			return "", err
		}
	}
	return clusterID, nil
}

func mountCluster(tpl interface{}, d *schema.ResourceData, m interface{},
	r *schema.Resource) (Mount, MountPoint, error) {
	var mountPoint MountPoint
	var mountConfig Mount

	client := m.(*common.DatabricksClient)
	mountPoint.exec = client.CommandExecutor()

	clusterID := d.Get("cluster_id").(string)
	clusterID, err := getMountingClusterID(client, clusterID)
	if err != nil {
		return mountConfig, mountPoint, err
	}

	mountType := reflect.TypeOf(tpl)
	mountTypePointer := reflect.New(mountType)
	mountReflectValue := mountTypePointer.Elem()
	err = internal.DataToReflectValue(d, r, mountReflectValue)
	if err != nil {
		return mountConfig, mountPoint, err
	}
	mountInterface := mountReflectValue.Interface()
	mountConfig = mountInterface.(Mount)

	name := d.Get("mount_name").(string)
	d.SetId(name)

	return mountConfig, NewMountPoint(client, name, clusterID), nil
}

// returns resource create mount for object store on workspace
func mountCreate(tpl interface{}, r *schema.Resource) func(*schema.ResourceData, interface{}) error {
	return func(d *schema.ResourceData, m interface{}) (err error) {
		mountConfig, mountPoint, err := mountCluster(tpl, d, m, r)
		if err != nil {
			return
		}
		log.Printf("[INFO] Mounting %s at /mnt/%s", mountConfig.Source(), d.Id())
		source, err := mountPoint.Mount(mountConfig)
		if err != nil {
			return
		}
		err = d.Set("source", source)
		if err != nil {
			return
		}
		return readMountSource(mountPoint, d)
	}
}

// reads and sets source of the mount
func readMountSource(mp MountPoint, d *schema.ResourceData) error {
	source, err := mp.Source()
	if err != nil {
		if err.Error() == "Mount not found" {
			log.Printf("[INFO] /mnt/%s is not mounted", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}
	return d.Set("source", source)
}

// return resource reader function
func mountRead(tpl Mount, r *schema.Resource) schema.ReadFunc {
	return func(d *schema.ResourceData, m interface{}) (err error) {
		_, mp, err := mountCluster(tpl, d, m, r)
		if err != nil {
			return
		}
		return readMountSource(mp, d)
	}
}

// returns delete resource function
func mountDelete(tpl Mount, r *schema.Resource) schema.DeleteFunc {
	return func(d *schema.ResourceData, m interface{}) (err error) {
		_, mp, err := mountCluster(tpl, d, m, r)
		if err != nil {
			return
		}
		log.Printf("[INFO] Unmounting /mnt/%s", d.Id())
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
