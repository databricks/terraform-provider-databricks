package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"

	"github.com/databrickslabs/terraform-provider-databricks/clusters"
	"github.com/databrickslabs/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Mount exposes generic url & extra config map options
type Mount interface {
	Source() string
	Config(client *common.DatabricksClient) map[string]string

	Name() string
	ValidateAndApplyDefaults(d *schema.ResourceData, client *common.DatabricksClient) error
}

// MountPoint is something actionable
type MountPoint struct {
	Exec           common.CommandExecutor
	ClusterID      string
	Name           string
	EncryptionType string
}

// Source returns mountpoint source
func (mp MountPoint) Source() (string, error) {
	result := mp.Exec.Execute(mp.ClusterID, "python", fmt.Sprintf(`
		dbutils.fs.refreshMounts()
		for mount in dbutils.fs.mounts():
			if mount.mountPoint == "/mnt/%s":
				dbutils.notebook.exit(mount.source)
		raise Exception("Mount not found")
	`, mp.Name))
	return result.Text(), result.Err()
}

// Delete removes mount from workspace
func (mp MountPoint) Delete() error {
	result := mp.Exec.Execute(mp.ClusterID, "python", fmt.Sprintf(`
		found = False
		mount_point = "/mnt/%s"
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
	`, mp.Name))
	return result.Err()
}

// Mount mounts object store on workspace
func (mp MountPoint) Mount(mo Mount, client *common.DatabricksClient) (source string, err error) {
	extraConfigs, err := json.Marshal(mo.Config(client))
	if err != nil {
		return
	}
	secretsRe := regexp.MustCompile(`"\{\{secrets/([^/]+)/([^\}]+)\}\}"`)
	extraConfigs = secretsRe.ReplaceAll(extraConfigs, []byte(`dbutils.secrets.get("$1", "$2")`))
	sparkConfRe := regexp.MustCompile(`"\{\{sparkconf/([^\}]+)\}\}"`)
	extraConfigs = sparkConfRe.ReplaceAll(extraConfigs, []byte(`spark.conf.get("$1")`))
	command := fmt.Sprintf(`
		def safe_mount(mount_point, mount_source, configs, encryptionType):
			for mount in dbutils.fs.mounts():
				if mount.mountPoint == mount_point and mount.source == mount_source:
					return
			try:
				dbutils.fs.mount(mount_source, mount_point, extra_configs=configs, encryption_type=encryptionType)
				dbutils.fs.refreshMounts()
				dbutils.fs.ls(mount_point)
				return mount_source
			except Exception as e:
				try:
					dbutils.fs.unmount(mount_point)
				except Exception as e2:
					print("Failed to unmount", e2)
				raise e
		mount_source = safe_mount("/mnt/%s", "%v", %s, "%s")
		dbutils.notebook.exit(mount_source)
	`, mp.Name, mo.Source(), extraConfigs, mp.EncryptionType) // lgtm[go/unsafe-quoting]
	result := mp.Exec.Execute(mp.ClusterID, "python", command)
	return result.Text(), result.Err()
}

func commonMountResource(tpl Mount, s map[string]*schema.Schema) *schema.Resource {
	resource := &schema.Resource{
		SchemaVersion: 2,
		Schema:        s,
	}
	// nolint should be a bigger context-aware refactor
	resource.CreateContext = mountCreate(tpl, resource)
	resource.ReadContext = mountRead(tpl, resource)
	resource.DeleteContext = mountDelete(tpl, resource)
	resource.Importer = &schema.ResourceImporter{
		StateContext: schema.ImportStatePassthroughContext,
	}
	return resource
}

func deprecatedMountTesource(r *schema.Resource) *schema.Resource {
	r.DeprecationMessage = "Resource is deprecated and will be removed in further versions. " +
		"Please rewrite configuration using `databricks_mount` resource. More info at " +
		"https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs/" +
		"resources/mount#migration-from-other-mount-resources"
	return r
}

// NewMountPoint returns new mount point config
func NewMountPoint(executor common.CommandExecutor, name, clusterID string) MountPoint {
	return MountPoint{
		Exec:      executor,
		ClusterID: clusterID,
		Name:      name,
	}
}

func getCommonClusterObject(clustersAPI clusters.ClustersAPI, clusterName string) clusters.Cluster {
	return clusters.Cluster{
		NumWorkers:  0,
		ClusterName: clusterName,
		SparkVersion: clustersAPI.LatestSparkVersionOrDefault(
			clusters.SparkVersionRequest{
				Latest:          true,
				LongTermSupport: true,
			}),
		NodeTypeID: clustersAPI.GetSmallestNodeType(
			clusters.NodeTypeRequest{
				LocalDisk: true,
			}),
		AutoterminationMinutes: 10,
		SparkConf: map[string]string{
			"spark.master":                     "local[*]",
			"spark.databricks.cluster.profile": "singleNode",
			"spark.scheduler.mode":             "FIFO",
		},
		CustomTags: map[string]string{
			"ResourceClass": "SingleNode",
		},
	}
}

func getOrCreateMountingCluster(clustersAPI clusters.ClustersAPI) (string, error) {
	cluster, err := clustersAPI.GetOrCreateRunningCluster("terraform-mount", getCommonClusterObject(clustersAPI, "terraform-mount"))
	if err != nil {
		return "", fmt.Errorf("failed to get mouting cluster: %w", err)
	}
	return cluster.ClusterID, nil
}

func getMountingClusterID(ctx context.Context, client *common.DatabricksClient, clusterID string) (string, error) {
	clustersAPI := clusters.NewClustersAPI(ctx, client)
	if clusterID == "" {
		return getOrCreateMountingCluster(clustersAPI)
	}
	clusterInfo, err := clustersAPI.Get(clusterID)
	if common.IsMissing(err) {
		return getOrCreateMountingCluster(clustersAPI)
	}
	if err != nil {
		return "", fmt.Errorf("failed to re-create mounting cluster: %w", err)
	}
	if !clusterInfo.IsRunningOrResizing() {
		err = clustersAPI.Start(clusterInfo.ClusterID)
		if err != nil {
			return "", fmt.Errorf("failed to start mounting cluster: %w", err)
		}
	}
	return clusterID, nil
}

func mountCluster(ctx context.Context, tpl interface{}, d *schema.ResourceData,
	m interface{}, r *schema.Resource) (Mount, MountPoint, error) {
	var mountPoint MountPoint
	var mountConfig Mount

	client := m.(*common.DatabricksClient)
	mountPoint.Exec = client.CommandExecutor(ctx)

	clusterID := d.Get("cluster_id").(string)
	clusterID, err := getMountingClusterID(ctx, client, clusterID)
	if err != nil {
		return mountConfig, mountPoint, err
	}
	mountPoint.ClusterID = clusterID

	mountType := reflect.TypeOf(tpl)
	mountTypePointer := reflect.New(mountType)
	mountReflectValue := mountTypePointer.Elem()
	common.DataToReflectValue(d, r, mountReflectValue)
	mountInterface := mountReflectValue.Interface()
	mountConfig = mountInterface.(Mount)

	if name, ok := d.GetOk("mount_name"); ok && name.(string) != "" {
		mountPoint.Name = name.(string)
	} else if name, ok := d.GetOk("name"); ok && name.(string) != "" {
		mountPoint.Name = name.(string)
	} else {
		return mountConfig, mountPoint, fmt.Errorf("nor 'mount_name' or 'name' are set")
	}

	d.SetId(mountPoint.Name)

	if v := d.Get("encryption_type"); v != nil {
		mountPoint.EncryptionType = v.(string)
	}

	return mountConfig, mountPoint, nil
}

// returns resource create mount for object store on workspace
func mountCreate(tpl interface{}, r *schema.Resource) func(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		mountConfig, mountPoint, err := mountCluster(ctx, tpl, d, m, r)
		if err != nil {
			return diag.FromErr(err)
		}
		client := m.(*common.DatabricksClient)
		log.Printf("[INFO] Mounting %s at /mnt/%s", mountConfig.Source(), d.Id())
		source, err := mountPoint.Mount(mountConfig, client)
		if err != nil {
			return diag.FromErr(err)
		}
		d.Set("source", source)
		return readMountSource(ctx, mountPoint, d)
	}
}

// reads and sets source of the mount
func readMountSource(ctx context.Context, mp MountPoint, d *schema.ResourceData) diag.Diagnostics {
	source, err := mp.Source()
	if err != nil {
		if err.Error() == "Mount not found" {
			log.Printf("[INFO] /mnt/%s is not mounted", d.Id())
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}
	d.Set("source", source)
	return nil
}

// return resource reader function
func mountRead(tpl Mount, r *schema.Resource) schema.ReadContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		_, mp, err := mountCluster(ctx, tpl, d, m, r)
		if err != nil {
			return diag.FromErr(err)
		}
		return readMountSource(ctx, mp, d)
	}
}

// returns delete resource function
func mountDelete(tpl Mount, r *schema.Resource) schema.DeleteContextFunc {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		_, mp, err := mountCluster(ctx, tpl, d, m, r)
		if err != nil {
			return diag.FromErr(err)
		}
		log.Printf("[INFO] Unmounting /mnt/%s", d.Id())
		if err = mp.Delete(); err != nil {
			return diag.FromErr(err)
		}
		return nil
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
