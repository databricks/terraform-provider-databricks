package clusters

import (
	"context"
	"log"
	"sort"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/libraries"
)

const DefaultProvisionTimeout = 30 * time.Minute
const DbfsDeprecationWarning = "For init scripts use 'volumes', 'workspace' or cloud storage location instead of 'dbfs'."

var clusterSchema = resourceClusterSchema()
var clusterSchemaWithRequiredClusterId = resourceClusterSchemaWithRequiredClusterId()
var clusterSchemaVersion = 2

func ResourceCluster() common.Resource {
	return common.Resource{
		Create:        resourceClusterCreate,
		Read:          resourceClusterRead,
		Update:        resourceClusterUpdate,
		Delete:        resourceClusterDelete,
		Schema:        clusterSchema,
		SchemaVersion: clusterSchemaVersion,
		Timeouts:      resourceClusterTimeouts(),
	}
}

func resourceClusterTimeouts() *schema.ResourceTimeout {
	return &schema.ResourceTimeout{
		Create: schema.DefaultTimeout(DefaultProvisionTimeout),
		Update: schema.DefaultTimeout(DefaultProvisionTimeout),
		Delete: schema.DefaultTimeout(DefaultProvisionTimeout),
	}

}

func SparkConfDiffSuppressFunc(k, old, new string, d *schema.ResourceData) bool {
	isPossiblyLegacyConfig := k == "spark_conf.%" && old == "1" && new == "0"
	isLegacyConfig := k == "spark_conf.spark.databricks.delta.preview.enabled"
	if isPossiblyLegacyConfig || isLegacyConfig {
		log.Printf("[DEBUG] Suppressing diff for k=%#v old=%#v new=%#v", k, old, new)
		return true
	}
	return false
}

func ZoneDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	if old != "" && (new == "auto" || new == "") {
		log.Printf("[INFO] Suppressing diff on availability zone")
		return true
	}
	return false
}

type ClusterSpec struct {
	compute.ClusterSpec
}

// This needs to be in OpenAPI spec compute.LibraryList.
// We have compute.InstallLibraries and compute.UninstallLibraries but no LibraryList
type LibraryList struct {
	ClusterId string            `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	Libraries []compute.Library `json:"libraries,omitempty" tf:"slice_set,alias:library"`
}

func (cll *LibraryList) Sort() {
	sort.Slice(cll.Libraries, func(i, j int) bool {
		return cll.Libraries[i].String() < cll.Libraries[j].String()
	})
}

func (ClusterSpec) CustomizeSchema(s map[string]*schema.Schema) map[string]*schema.Schema {
	common.CustomizeSchemaPath(s, "cluster_source").SetReadOnly()
	common.CustomizeSchemaPath(s, "enable_elastic_disk").SetComputed()
	common.CustomizeSchemaPath(s, "enable_local_disk_encryption").SetComputed()
	common.CustomizeSchemaPath(s, "node_type_id").SetComputed().SetConflictsWith([]string{"driver_instance_pool_id", "instance_pool_id"})
	common.CustomizeSchemaPath(s, "driver_node_type_id").SetComputed().SetConflictsWith([]string{"driver_instance_pool_id", "instance_pool_id"})
	common.CustomizeSchemaPath(s, "driver_instance_pool_id").SetComputed().SetConflictsWith([]string{"driver_node_type_id", "node_type_id"})
	common.CustomizeSchemaPath(s, "ssh_public_keys").SetMaxItems(10)
	common.CustomizeSchemaPath(s, "init_scripts").SetMaxItems(10)
	common.CustomizeSchemaPath(s, "init_scripts", "dbfs").SetDeprecated(DbfsDeprecationWarning)
	common.CustomizeSchemaPath(s, "init_scripts", "dbfs", "destination").SetRequired()
	common.CustomizeSchemaPath(s, "init_scripts", "s3", "destination").SetRequired()
	common.CustomizeSchemaPath(s, "init_scripts", "volumes", "destination").SetRequired()
	common.CustomizeSchemaPath(s, "init_scripts", "workspace", "destination").SetRequired()
	common.CustomizeSchemaPath(s, "workload_type", "clients").SetRequired()
	common.CustomizeSchemaPath(s, "workload_type", "clients", "notebooks").SetDefault(true)
	common.CustomizeSchemaPath(s, "workload_type", "clients", "jobs").SetDefault(true)
	common.CustomizeSchemaPath(s).AddNewField("idempotency_token", &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		ForceNew: true,
	})
	common.CustomizeSchemaPath(s, "data_security_mode").SetSuppressDiff()
	common.CustomizeSchemaPath(s, "docker_image", "url").SetRequired()
	common.CustomizeSchemaPath(s, "docker_image", "basic_auth", "password").SetRequired().SetSensitive()
	common.CustomizeSchemaPath(s, "docker_image", "basic_auth", "username").SetRequired()
	common.CustomizeSchemaPath(s, "spark_conf").SetCustomSuppressDiff(SparkConfDiffSuppressFunc)
	common.CustomizeSchemaPath(s, "aws_attributes").SetSuppressDiff().SetConflictsWith([]string{"azure_attributes", "gcp_attributes"})
	common.CustomizeSchemaPath(s, "aws_attributes", "zone_id").SetCustomSuppressDiff(ZoneDiffSuppress)
	common.CustomizeSchemaPath(s, "azure_attributes").SetSuppressDiff().SetConflictsWith([]string{"aws_attributes", "gcp_attributes"})
	common.CustomizeSchemaPath(s, "gcp_attributes").SetSuppressDiff().SetConflictsWith([]string{"aws_attributes", "azure_attributes"})

	common.CustomizeSchemaPath(s).AddNewField("library", common.StructToSchema(libraries.ClusterLibraryList{},
		func(ss map[string]*schema.Schema) map[string]*schema.Schema {
			ss["library"].Set = func(i any) int {
				lib := libraries.NewLibraryFromInstanceState(i)
				return schema.HashString(lib.String())
			}
			return ss
		})["library"])

	common.CustomizeSchemaPath(s, "autotermination_minutes").SetDefault(60)
	common.CustomizeSchemaPath(s, "autoscale", "max_workers").SetOptional()
	common.CustomizeSchemaPath(s, "autoscale", "min_workers").SetOptional()
	common.CustomizeSchemaPath(s, "cluster_log_conf", "dbfs", "destination").SetRequired()
	common.CustomizeSchemaPath(s, "cluster_log_conf", "s3", "destination").SetRequired()
	common.CustomizeSchemaPath(s, "spark_version").SetRequired()
	common.CustomizeSchemaPath(s).AddNewField("cluster_id", &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	})
	common.CustomizeSchemaPath(s).AddNewField("default_tags", &schema.Schema{
		Type:     schema.TypeMap,
		Computed: true,
	})
	common.CustomizeSchemaPath(s).AddNewField("state", &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	})
	common.CustomizeSchemaPath(s, "instance_pool_id").SetConflictsWith([]string{"driver_node_type_id", "node_type_id"})
	common.CustomizeSchemaPath(s, "runtime_engine").SetValidateFunc(validation.StringInSlice([]string{"PHOTON", "STANDARD"}, false))
	common.CustomizeSchemaPath(s).AddNewField("is_pinned", &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Default:  false,
		DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
			if old == "" && new == "false" {
				return true
			}
			return old == new
		},
	})
	common.CustomizeSchemaPath(s).AddNewField("url", &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	})
	common.CustomizeSchemaPath(s, "num_workers").SetDefault(0).SetValidateDiagFunc(validation.ToDiagFunc(validation.IntAtLeast(0)))
	common.CustomizeSchemaPath(s).AddNewField("cluster_mount_info", &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: common.StructToSchema(MountInfo{}, nil),
		},
	})

	return s
}

func resourceClusterSchemaWithRequiredClusterId() map[string]*schema.Schema {
	s := resourceClusterSchema()
	common.CustomizeSchemaPath(s, "cluster_id").SetRequired()
	return s
}

func resourceClusterSchema() map[string]*schema.Schema {
	return common.StructToSchema(ClusterSpec{}, nil)
}

func resourceClusterCreate(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
	start := time.Now()
	timeout := d.Timeout(schema.TimeoutCreate)
	w, err := c.WorkspaceClient()
	if err != nil {
		return err
	}
	clusters := w.Clusters
	var cluster compute.CreateCluster
	common.DataToStructPointer(d, clusterSchema, &cluster)
	if err := ValidateCluster(cluster); err != nil {
		return err
	}
	ModifyRequestOnInstancePool(&cluster)
	// TODO: propagate d.Timeout(schema.TimeoutCreate)
	clusterInfo, err := clusters.CreateAndWait(ctx, cluster)
	if err != nil {
		return err
	}
	d.SetId(clusterInfo.ClusterId)
	d.Set("cluster_id", d.Id())
	isPinned, ok := d.GetOk("is_pinned")
	if ok && isPinned.(bool) {
		err = clusters.PinByClusterId(ctx, d.Id())
		if err != nil {
			return err
		}
	}

	var libraryList LibraryList
	libraryList.ClusterId = d.Id()
	common.DataToStructPointer(d, clusterSchema, &libraryList)
	if len(libraryList.Libraries) > 0 {
		if err = w.Libraries.Install(ctx, compute.InstallLibraries{
			ClusterId: d.Id(),
			Libraries: libraryList.Libraries,
		}); err != nil {
			return err
		}
		_, err := libraries.WaitForLibrariesInstalledSdk(ctx, w, compute.Wait{
			ClusterID: d.Id(),
			IsRunning: clusterInfo.IsRunningOrResizing(),
			IsRefresh: false,
		}, timeout-time.Since(start))
		if err != nil {
			return err
		}
	}
	return nil
}

func setPinnedStatus(ctx context.Context, d *schema.ResourceData, clusterAPI compute.ClustersInterface) error {
	events, err := clusterAPI.EventsAll(ctx, compute.GetEvents{
		ClusterId:  d.Id(),
		Limit:      1,
		Order:      compute.GetEventsOrderDesc,
		EventTypes: []compute.EventType{compute.EventTypePinned, compute.EventTypeUnpinned},
	})
	if err != nil {
		return err
	}
	pinnedEvent := compute.EventTypeUnpinned
	if len(events) > 0 {
		pinnedEvent = events[0].Type
	}
	return d.Set("is_pinned", pinnedEvent == compute.EventTypePinned)
}

func resourceClusterRead(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
	w, err := c.WorkspaceClient()
	if err != nil {
		return err
	}
	clusterAPI := w.Clusters
	clusterInfo, err := clusterAPI.GetByClusterId(ctx, d.Id())
	if err != nil {
		return err
	}
	if err = common.StructToData(clusterInfo, clusterSchema, d); err != nil {
		return err
	}
	if err = setPinnedStatus(ctx, d, clusterAPI); err != nil {
		return err
	}

	d.Set("url", c.FormatURL("#setting/clusters/", d.Id(), "/configuration"))
	shouldSkipLibrariesRead := !common.IsExporter(ctx)
	if d.Get("library.#").(int) == 0 && shouldSkipLibrariesRead {
		// don't add externally added libraries, if config has no `library {}` blocks
		// TODO: check if it still works fine with importing. Perhaps os.Setenv will do the trick
		return nil
	}

	libsClusterStatus, err := libraries.WaitForLibrariesInstalledSdk(ctx, w, compute.Wait{
		ClusterID: d.Id(),
		IsRunning: clusterInfo.IsRunningOrResizing(),
		IsRefresh: true,
	}, d.Timeout(schema.TimeoutRead))
	if err != nil {
		return err
	}
	libList := libsClusterStatus.ToLibraryList()
	return common.StructToData(libList, clusterSchemaWithRequiredClusterId, d)
}

func hasClusterConfigChanged(d *schema.ResourceData) bool {
	for k := range clusterSchema {
		// TODO: create a map if we'll add more non-cluster config parameters in the future
		if k == "library" || k == "is_pinned" {
			continue
		}
		if d.HasChange(k) {
			return true
		}
	}
	return false
}

func resourceClusterUpdate(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
	w, err := c.WorkspaceClient()
	if err != nil {
		return err
	}
	clusters := w.Clusters
	var cluster compute.CreateCluster
	common.DataToStructPointer(d, clusterSchema, &cluster)
	clusterId := d.Id()
	var clusterInfo *compute.ClusterDetails

	if hasClusterConfigChanged(d) {
		log.Printf("[DEBUG] Cluster state has changed!")
		if err := ValidateCluster(cluster); err != nil {
			return err
		}
		ModifyRequestOnInstancePool(&cluster)
		FixInstancePoolChangeIfAny(d, cluster)

		// We can only call the resize api if the cluster is in the running state
		// and only the cluster size (ie num_workers OR autoscale) is being changed
		hasNumWorkersChanged := d.HasChange("num_workers")
		hasAutoscaleChanged := d.HasChange("autoscale")
		hasOnlyResizeClusterConfigChanged := true
		for k := range clusterSchema {
			if k == "library" ||
				k == "is_pinned" ||
				k == "num_workers" ||
				k == "autoscale" {
				continue
			}
			if d.HasChange(k) {
				hasOnlyResizeClusterConfigChanged = false
			}
		}
		clusterInfo, err = clusters.GetByClusterId(ctx, clusterId)
		if err != nil {
			return err
		}

		isNumWorkersResizeForNonAutoscalingCluster := hasOnlyResizeClusterConfigChanged &&
			hasNumWorkersChanged &&
			!hasAutoscaleChanged &&
			clusterInfo.State == ClusterStateRunning
		isAutoScalingToNonAutoscalingResize := hasOnlyResizeClusterConfigChanged &&
			hasAutoscaleChanged &&
			hasNumWorkersChanged &&
			cluster.Autoscale == nil &&
			clusterInfo.State == ClusterStateRunning
		isAutoscaleConfigResizeForAutoscalingCluster := hasOnlyResizeClusterConfigChanged &&
			hasAutoscaleChanged &&
			!hasNumWorkersChanged &&
			clusterInfo.State == ClusterStateRunning
		isNonAutoScalingToAutoscalingResize := hasOnlyResizeClusterConfigChanged &&
			hasAutoscaleChanged &&
			hasNumWorkersChanged &&
			cluster.Autoscale != nil &&
			clusterInfo.State == ClusterStateRunning

		// We prefer to use the resize API in cases when only the number of
		// workers is changed because a resizing cluster can still serve queries
		if isNumWorkersResizeForNonAutoscalingCluster ||
			isAutoScalingToNonAutoscalingResize {
			_, err = clusters.Resize(ctx, compute.ResizeCluster{
				ClusterId:  clusterId,
				NumWorkers: cluster.NumWorkers,
			})
			if err != nil {
				return err
			}
		} else if isAutoscaleConfigResizeForAutoscalingCluster ||
			isNonAutoScalingToAutoscalingResize {
			_, err = clusters.Resize(ctx, compute.ResizeCluster{
				ClusterId: clusterId,
				Autoscale: cluster.Autoscale,
			})
		} else {
			var editCluster compute.EditCluster
			editCluster.ClusterId = clusterId
			common.DataToStructPointer(d, clusterSchemaWithRequiredClusterId, &editCluster)
			_, err = clusters.Edit(ctx, editCluster)
		}
		if err != nil {
			return err
		}
	}
	oldPinned, newPinned := d.GetChange("is_pinned")
	if oldPinned.(bool) != newPinned.(bool) {
		log.Printf("[DEBUG] Update: is_pinned. Old: %v, New: %v", oldPinned, newPinned)
		if newPinned.(bool) {
			err = clusters.PinByClusterId(ctx, clusterId)
		} else {
			err = clusters.UnpinByClusterId(ctx, clusterId)
		}
		if err != nil {
			return err
		}
	}
	oldNumLibs, newNumLibs := d.GetChange("library.#")
	if oldNumLibs == newNumLibs && oldNumLibs.(int) == 0 {
		// don't add externally added libraries, if config has no `library {}` blocks
		return nil
	}
	var libraryList LibraryList
	libraryList.ClusterId = clusterId
	common.DataToStructPointer(d, clusterSchema, &libraryList)
	libsClusterStatus, err := w.Libraries.ClusterStatusByClusterId(ctx, clusterId)
	if err != nil {
		return err
	}
	libraryList.ClusterId = clusterId
	libsToInstall, libsToUninstall := libraries.GetLibrariesToInstallAndUninstall(compute.InstallLibraries{
		ClusterId: clusterId,
		Libraries: libraryList.Libraries,
	}, libsClusterStatus)
	clusterInfo, err = clusters.GetByClusterId(ctx, clusterId)
	if err != nil {
		return err
	}
	if len(libsToUninstall.Libraries) > 0 || len(libsToInstall.Libraries) > 0 {
		if !clusterInfo.IsRunningOrResizing() {
			if _, err = clusters.StartByClusterIdAndWait(ctx, clusterId); err != nil {
				return err
			}
		}
		// clusters.StartAndGetInfo() always returns a running cluster
		// or errors out, so we just know the cluster is active.
		// TODO: Wait needed here?
		err = w.Libraries.UpdateAndWait(ctx, compute.Update{
			ClusterId: clusterId,
			Install:   libsToInstall.Libraries,
			Uninstall: libsToUninstall.Libraries,
		})
		if err != nil {
			return err
		}
		if clusterInfo.State == ClusterStateTerminated {
			log.Printf("[INFO] %s was in TERMINATED state, so terminating it again", clusterId)
			if err = clusters.DeleteByClusterId(ctx, clusterId); err != nil {
				return err
			}
		}
	}
	return nil
}

func resourceClusterDelete(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
	w, err := c.WorkspaceClient()
	if err != nil {
		return err
	}
	return w.Clusters.PermanentDeleteByClusterId(ctx, d.Id())
}

func init() {
	common.RegisterResourceProvider(compute.ClusterSpec{}, ClusterSpec{})
	common.RegisterResourceProvider(compute.Library{}, LibraryResource{})
}
