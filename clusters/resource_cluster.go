package clusters

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/libraries"
)

const DefaultProvisionTimeout = 30 * time.Minute
const DbfsDeprecationWarning = "For init scripts use 'volumes', 'workspace' or cloud storage location instead of 'dbfs'."

var clusterSchema = resourceClusterSchema()
var clusterSchemaVersion = 4

const (
	unsupportedExceptCreateEditClusterSpecErr = "unsupported type %T, must be one of %scompute.CreateCluster, %scompute.ClusterSpec or %scompute.EditCluster. Please report this issue to the GitHub repo"
)

func ResourceCluster() common.Resource {
	return common.Resource{
		Create: resourceClusterCreate,
		Read:   resourceClusterRead,
		Update: resourceClusterUpdate,
		Delete: resourceClusterDelete,
		Schema: clusterSchema,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff) error {
			if isSingleNode, ok := d.GetOk("is_single_node"); ok && isSingleNode.(bool) {
				return singleNodeClusterChangesCustomizeDiff(d)
			}
			return nil
		},
		SchemaVersion: clusterSchemaVersion,
		Timeouts:      resourceClusterTimeouts(),
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    clusterSchemaV0(),
				Version: 3,
				Upgrade: removeZeroAwsEbsVolumeAttributes,
			},
		},
	}
}

// the API automatically sets the `ResourceClass` key in `custom_tags` and two other keys in the `spark_conf`.
// If the user hasn't set these explicitly in their config, the plan marks these keys for removal.
// This method copies the values for these keys from state to the plan.
// This needs to be done in addition to setting these attributes as computed; otherwise, this customization
// won't take effect for users who have set additional `spark_conf` or `custom_tags`.
func singleNodeClusterChangesCustomizeDiff(d *schema.ResourceDiff) error {
	autoConfigAttributes := map[string][]string{
		"custom_tags": {"ResourceClass"},
		"spark_conf":  {"spark.databricks.cluster.profile", "spark.master"},
	}

	for key, attributes := range autoConfigAttributes {
		if !d.HasChange(key) {
			continue
		}

		o, n := d.GetChange(key)
		old, okOld := o.(map[string]interface{})
		new, okNew := n.(map[string]interface{})

		if !okNew || !okOld {
			return fmt.Errorf("internal type casting error n: %T, o: %T", n, o)
		}

		log.Printf("[DEBUG] values for key %s, old: %v, new: %v", key, old, new)

		for _, attribute := range attributes {
			if _, exists := new[attribute]; exists && new[attribute] != nil {
				continue
			}

			new[attribute] = old[attribute]
		}

		if err := d.SetNew(key, new); err != nil {
			return err
		}
	}

	return nil
}

func clusterSchemaV0() cty.Type {
	return (&schema.Resource{
		Schema: clusterSchema}).CoreConfigSchema().ImpliedType()
}

func removeZeroAwsEbsVolumeAttributes(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	newState := map[string]any{}
	for k, v := range rawState {
		switch k {
		case "aws_attributes":
			awsAttributes, ok := v.(map[string]any)
			if !ok {
				continue
			}

			if awsAttributes["ebs_volume_count"] == 0 {
				log.Printf("[INFO] remove zero ebs_volume_count")
				delete(awsAttributes, "ebs_volume_count")
			}
			if awsAttributes["ebs_volume_iops"] == 0 {
				log.Printf("[INFO] remove zero ebs_volume_iops")
				delete(awsAttributes, "ebs_volume_iops")
			}
			if awsAttributes["ebs_volume_size"] == 0 {
				log.Printf("[INFO] remove zero ebs_volume_size")
				delete(awsAttributes, "ebs_volume_size")
			}

			newState[k] = awsAttributes
		default:
			newState[k] = v
		}
	}
	return newState, nil
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

// This method is a duplicate of ModifyRequestOnInstancePool() in clusters/clusters_api.go that uses Go SDK.
// Long term, ModifyRequestOnInstancePool() in clusters_api.go will be removed once all the resources using clusters are migrated to Go SDK.
func ModifyRequestOnInstancePool(cluster any) error {
	switch c := cluster.(type) {
	case *compute.ClusterSpec:
		// Instance profile id does not exist or not set
		if c.InstancePoolId == "" {
			// Worker must use an instance pool if driver uses an instance pool,
			// therefore empty the computed value for driver instance pool.
			c.DriverInstancePoolId = ""
			return nil
		}
		if c.AwsAttributes != nil {
			// Reset AwsAttributes
			awsAttributes := compute.AwsAttributes{
				InstanceProfileArn: c.AwsAttributes.InstanceProfileArn,
			}
			c.AwsAttributes = &awsAttributes
		}
		if c.AzureAttributes != nil {
			c.AzureAttributes = &compute.AzureAttributes{}
		}
		if c.GcpAttributes != nil {
			gcpAttributes := compute.GcpAttributes{
				GoogleServiceAccount: c.GcpAttributes.GoogleServiceAccount,
			}
			c.GcpAttributes = &gcpAttributes
		}
		c.EnableElasticDisk = false
		c.NodeTypeId = ""
		c.DriverNodeTypeId = ""
		return nil
	case *compute.CreateCluster:
		// Instance profile id does not exist or not set
		if c.InstancePoolId == "" {
			// Worker must use an instance pool if driver uses an instance pool,
			// therefore empty the computed value for driver instance pool.
			c.DriverInstancePoolId = ""
			return nil
		}
		if c.AwsAttributes != nil {
			// Reset AwsAttributes
			awsAttributes := compute.AwsAttributes{
				InstanceProfileArn: c.AwsAttributes.InstanceProfileArn,
			}
			c.AwsAttributes = &awsAttributes
		}
		if c.AzureAttributes != nil {
			c.AzureAttributes = &compute.AzureAttributes{}
		}
		if c.GcpAttributes != nil {
			gcpAttributes := compute.GcpAttributes{
				GoogleServiceAccount: c.GcpAttributes.GoogleServiceAccount,
			}
			c.GcpAttributes = &gcpAttributes
		}
		c.EnableElasticDisk = false
		c.NodeTypeId = ""
		c.DriverNodeTypeId = ""
		return nil
	case *compute.EditCluster:
		// Instance profile id does not exist or not set
		if c.InstancePoolId == "" {
			// Worker must use an instance pool if driver uses an instance pool,
			// therefore empty the computed value for driver instance pool.
			c.DriverInstancePoolId = ""
			return nil
		}
		if c.AwsAttributes != nil {
			// Reset AwsAttributes
			awsAttributes := compute.AwsAttributes{
				InstanceProfileArn: c.AwsAttributes.InstanceProfileArn,
			}
			c.AwsAttributes = &awsAttributes
		}
		if c.AzureAttributes != nil {
			c.AzureAttributes = &compute.AzureAttributes{}
		}
		if c.GcpAttributes != nil {
			gcpAttributes := compute.GcpAttributes{
				GoogleServiceAccount: c.GcpAttributes.GoogleServiceAccount,
			}
			c.GcpAttributes = &gcpAttributes
		}
		c.EnableElasticDisk = false
		c.NodeTypeId = ""
		c.DriverNodeTypeId = ""
		return nil
	default:
		return fmt.Errorf(unsupportedExceptCreateEditClusterSpecErr, cluster, "*", "*", "*")
	}
}

// This method is a duplicate of FixInstancePoolChangeIfAny(d *schema.ResourceData) in clusters/clusters_api.go that uses Go SDK.
// Long term, FixInstancePoolChangeIfAny(d *schema.ResourceData) in clusters_api.go will be removed once all the resources using clusters are migrated to Go SDK.
// https://github.com/databricks/terraform-provider-databricks/issues/824
func FixInstancePoolChangeIfAny(d *schema.ResourceData, cluster any) error {
	switch c := cluster.(type) {
	case *compute.ClusterSpec:
		oldInstancePool, newInstancePool := d.GetChange("instance_pool_id")
		oldDriverPool, newDriverPool := d.GetChange("driver_instance_pool_id")
		if oldInstancePool != newInstancePool &&
			oldDriverPool == oldInstancePool &&
			oldDriverPool == newDriverPool {
			c.DriverInstancePoolId = c.InstancePoolId
		}
		return nil
	case *compute.EditCluster:
		oldInstancePool, newInstancePool := d.GetChange("instance_pool_id")
		oldDriverPool, newDriverPool := d.GetChange("driver_instance_pool_id")
		if oldInstancePool != newInstancePool &&
			oldDriverPool == oldInstancePool &&
			oldDriverPool == newDriverPool {
			c.DriverInstancePoolId = c.InstancePoolId
		}
		return nil
	default:
		return fmt.Errorf(unsupportedExceptCreateEditClusterSpecErr, cluster, "*", "*", "*")
	}
}

func SetForceSendFieldsForCluster(cluster any, d *schema.ResourceData) error {
	switch c := cluster.(type) {
	case *compute.ClusterSpec:
		// Used in jobs.
		if c.Autoscale == nil {
			c.ForceSendFields = append(c.ForceSendFields, "NumWorkers")
		}
		// Workload type is not relevant in jobs clusters.
		return nil
	case *compute.CreateCluster:
		if c.Autoscale == nil {
			c.ForceSendFields = append(c.ForceSendFields, "NumWorkers")
		}
		// If workload type is set by the user, the fields within Clients should always be sent.
		// These default to true if not set.
		if c.WorkloadType != nil {
			c.WorkloadType.Clients.ForceSendFields = []string{"Jobs", "Notebooks"}
		}
		return nil
	case *compute.EditCluster:
		if c.Autoscale == nil {
			c.ForceSendFields = append(c.ForceSendFields, "NumWorkers")
		}
		// If workload type is set by the user, the fields within Clients should always be sent.
		// These default to true if not set.
		if c.WorkloadType != nil {
			c.WorkloadType.Clients.ForceSendFields = []string{"Jobs", "Notebooks"}
		}
		return nil
	default:
		return fmt.Errorf(unsupportedExceptCreateEditClusterSpecErr, cluster, "*", "*", "*")
	}
}

type LibraryWithAlias struct {
	Libraries []compute.Library `json:"libraries,omitempty" tf:"slice_set,alias:library"`
}

type ClusterSpec struct {
	compute.ClusterSpec
	LibraryWithAlias
}

func (ClusterSpec) CustomizeSchemaResourceSpecific(s *common.CustomizableSchema) *common.CustomizableSchema {
	s.AddNewField("default_tags", &schema.Schema{
		Type:     schema.TypeMap,
		Computed: true,
	})
	s.AddNewField("is_pinned", &schema.Schema{
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
	s.AddNewField("no_wait", &schema.Schema{
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
	s.AddNewField("state", &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	})
	s.AddNewField("url", &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	})
	s.AddNewField("autotermination_minutes", &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Default:  60,
	})
	return s
}

func (ClusterSpec) CustomizeSchema(s *common.CustomizableSchema) *common.CustomizableSchema {
	s.SchemaPath("enable_elastic_disk").SetComputed()
	s.SchemaPath("enable_local_disk_encryption").SetComputed()
	s.SchemaPath("node_type_id").SetComputed().SetConflictsWith([]string{"driver_instance_pool_id", "instance_pool_id"})
	s.SchemaPath("driver_node_type_id").SetComputed().SetConflictsWith([]string{"driver_instance_pool_id", "instance_pool_id"})
	s.SchemaPath("driver_instance_pool_id").SetComputed().SetConflictsWith([]string{"driver_node_type_id", "node_type_id"})
	s.SchemaPath("ssh_public_keys").SetMaxItems(10)
	s.SchemaPath("init_scripts").SetMaxItems(10)
	s.SchemaPath("init_scripts", "dbfs").SetDeprecated(DbfsDeprecationWarning)
	s.SchemaPath("init_scripts", "dbfs", "destination").SetRequired()
	s.SchemaPath("init_scripts", "s3", "destination").SetRequired()
	s.SchemaPath("init_scripts", "volumes", "destination").SetRequired()
	s.SchemaPath("init_scripts", "workspace", "destination").SetRequired()
	s.SchemaPath("workload_type", "clients").SetRequired()
	s.SchemaPath("workload_type", "clients", "notebooks").SetDefault(true)
	s.SchemaPath("workload_type", "clients", "jobs").SetDefault(true)
	s.SchemaPath("library").Schema.Set = func(i any) int {
		lib := libraries.NewLibraryFromInstanceState(i)
		return schema.HashString(lib.String())
	}
	s.AddNewField("idempotency_token", &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		ForceNew: true,
	})
	s.SchemaPath("data_security_mode").SetSuppressDiff()
	s.SchemaPath("docker_image", "url").SetRequired()
	s.SchemaPath("docker_image", "basic_auth", "password").SetRequired().SetSensitive()
	s.SchemaPath("docker_image", "basic_auth", "username").SetRequired()
	s.SchemaPath("spark_conf").SetCustomSuppressDiff(SparkConfDiffSuppressFunc).SetComputed().SetOptional()
	s.SchemaPath("custom_tags").SetComputed().SetOptional()
	s.SchemaPath("aws_attributes").SetSuppressDiff().SetConflictsWith([]string{"azure_attributes", "gcp_attributes"})
	s.SchemaPath("aws_attributes", "zone_id").SetCustomSuppressDiff(ZoneDiffSuppress)
	s.SchemaPath("azure_attributes").SetSuppressDiff().SetConflictsWith([]string{"aws_attributes", "gcp_attributes"})
	s.SchemaPath("gcp_attributes").SetSuppressDiff().SetConflictsWith([]string{"aws_attributes", "azure_attributes"})
	s.SchemaPath("autoscale", "max_workers").SetOptional()
	s.SchemaPath("autoscale", "min_workers").SetOptional()
	s.SchemaPath("cluster_log_conf", "dbfs", "destination").SetRequired()
	s.SchemaPath("cluster_log_conf", "s3", "destination").SetRequired()
	s.SchemaPath("spark_version").SetRequired()
	s.AddNewField("cluster_id", &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	})
	s.SchemaPath("instance_pool_id").SetConflictsWith([]string{"driver_node_type_id", "node_type_id"})
	s.SchemaPath("runtime_engine").SetValidateFunc(validation.StringInSlice([]string{"PHOTON", "STANDARD"}, false))
	s.SchemaPath("num_workers").SetDefault(0).SetValidateDiagFunc(validation.ToDiagFunc(validation.IntAtLeast(0)))
	s.AddNewField("cluster_mount_info", &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: common.StructToSchema(MountInfo{}, nil),
		},
	})
	// Adding it back in the resource specific customization function because it is not relevant for other resources.
	s.RemoveField("autotermination_minutes")

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
	var createClusterRequest compute.CreateCluster
	common.DataToStructPointer(d, clusterSchema, &createClusterRequest)
	if err = ModifyRequestOnInstancePool(&createClusterRequest); err != nil {
		return err
	}
	SetForceSendFieldsForCluster(&createClusterRequest, d)
	if createClusterRequest.GcpAttributes != nil {
		if _, ok := d.GetOkExists("gcp_attributes.0.local_ssd_count"); ok {
			createClusterRequest.GcpAttributes.ForceSendFields = []string{"LocalSsdCount"}
		}
	}
	clusterWaiter, err := clusters.Create(ctx, createClusterRequest)
	if err != nil {
		return err
	}

	d.SetId(clusterWaiter.ClusterId)
	d.Set("cluster_id", d.Id())
	isPinned, ok := d.GetOk("is_pinned")
	if ok && isPinned.(bool) {
		err = clusters.PinByClusterId(ctx, d.Id())
		if err != nil {
			return err
		}
	}

	var cluster ClusterSpec
	common.DataToStructPointer(d, clusterSchema, &cluster)
	if len(cluster.Libraries) > 0 {
		if err = w.Libraries.Install(ctx, compute.InstallLibraries{
			ClusterId: d.Id(),
			Libraries: cluster.Libraries,
		}); err != nil {
			return err
		}
	}

	// If there is a no_wait flag set to true, don't wait for the cluster to be created
	noWait, ok := d.GetOk("no_wait")
	if ok && noWait.(bool) {
		return nil
	}

	clusterInfo, err := clusterWaiter.GetWithTimeout(timeout)
	if err != nil {
		// In case of "ERROR" or "TERMINATED" state, WaitGetClusterRunning returns an error and we should delete the cluster before returning
		deleteError := resourceClusterDelete(ctx, d, c)
		if deleteError != nil {
			return fmt.Errorf("failed to create cluster: %v and failed to delete it during cleanup: %v", err, deleteError)
		}
		return err
	}

	if len(cluster.Libraries) > 0 {
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
	clusterDetails := clusterAPI.List(ctx, compute.ListClustersRequest{
		FilterBy: &compute.ListClustersFilterBy{
			IsPinned: true,
		},
		PageSize: 100, // pinned cluster limit - just get all of them
	})

	for clusterDetails.HasNext(ctx) {
		detail, err := clusterDetails.Next(ctx)
		if err != nil {
			return err
		}
		if detail.ClusterId == d.Id() {
			return d.Set("is_pinned", true)
		}
	}
	return d.Set("is_pinned", false)
}

func resourceClusterRead(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
	w, err := c.WorkspaceClient()
	if err != nil {
		return err
	}
	clusterAPI := w.Clusters
	clusterInfo, err := clusterAPI.GetByClusterId(ctx, d.Id())
	if err != nil {
		return wrapMissingClusterError(err, d.Id())
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
	return common.StructToData(LibraryWithAlias{
		Libraries: libList.Libraries,
	}, clusterSchema, d)
}

func hasClusterConfigChanged(d *schema.ResourceData) bool {
	for k := range clusterSchema {
		// TODO: create a map if we'll add more non-cluster config parameters in the future
		if k == "library" || k == "is_pinned" || k == "no_wait" {
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
	var cluster compute.EditCluster
	common.DataToStructPointer(d, clusterSchema, &cluster)
	clusterId := d.Id()
	cluster.ClusterId = clusterId
	var clusterInfo *compute.ClusterDetails

	if hasClusterConfigChanged(d) {
		log.Printf("[DEBUG] Cluster state has changed!")
		if err = ModifyRequestOnInstancePool(&cluster); err != nil {
			return err
		}
		err = FixInstancePoolChangeIfAny(d, &cluster)
		if err != nil {
			return err
		}

		// We can only call the resize api if the cluster is in the running state
		// and only the cluster size (ie num_workers OR autoscale) is being changed
		hasNumWorkersChanged := d.HasChange("num_workers")
		hasAutoscaleChanged := d.HasChange("autoscale")
		hasOnlyResizeClusterConfigChanged := true
		for k := range clusterSchema {
			if k == "library" ||
				k == "is_pinned" ||
				k == "no_wait" ||
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
			return wrapMissingClusterError(err, d.Id())
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
				ClusterId:       clusterId,
				NumWorkers:      cluster.NumWorkers,
				ForceSendFields: []string{"NumWorkers"},
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
			SetForceSendFieldsForCluster(&cluster, d)

			err = retry.RetryContext(ctx, 15*time.Minute, func() *retry.RetryError {
				_, err = clusters.Edit(ctx, cluster)
				if err == nil {
					return nil
				}
				var apiErr *apierr.APIError
				// Only Running and Terminated clusters can be modified. In particular, autoscaling clusters cannot be modified
				// while the resizing is ongoing. We retry in this case. Scaling can take several minutes.
				if errors.As(err, &apiErr) && apiErr.ErrorCode == "INVALID_STATE" {
					return retry.RetryableError(fmt.Errorf("cluster %s cannot be modified in its current state", clusterId))
				}
				return retry.NonRetryableError(err)
			})

		}
		if err != nil {
			return err
		}
	} else {
		_, err = clusters.GetByClusterId(ctx, clusterId)
		if err != nil {
			return wrapMissingClusterError(err, d.Id())
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
	libsClusterStatus, err := w.Libraries.ClusterStatusByClusterId(ctx, clusterId)
	if err != nil {
		return err
	}

	var clusterLibraries LibraryWithAlias
	common.DataToStructPointer(d, clusterSchema, &clusterLibraries)
	libsToInstall, libsToUninstall := libraries.GetLibrariesToInstallAndUninstall(clusterLibraries.Libraries, libsClusterStatus)

	clusterInfo, err = clusters.GetByClusterId(ctx, clusterId)
	if err != nil {
		return wrapMissingClusterError(err, d.Id())
	}
	if len(libsToUninstall) > 0 || len(libsToInstall) > 0 {
		if !clusterInfo.IsRunningOrResizing() {
			if _, err = clusters.StartByClusterIdAndWait(ctx, clusterId); err != nil {
				return err
			}
		}
		// clusters.StartAndGetInfo() always returns a running cluster
		// or errors out, so we just know the cluster is active.
		err = w.Libraries.UpdateAndWait(ctx, compute.Update{
			ClusterId: clusterId,
			Install:   libsToInstall,
			Uninstall: libsToUninstall,
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
	err = w.Clusters.PermanentDeleteByClusterId(ctx, d.Id())
	if err == nil || apierr.IsMissing(err) {
		return nil
	}
	if !strings.Contains(err.Error(), "unpin the cluster first") {
		return err
	}
	err = w.Clusters.UnpinByClusterId(ctx, d.Id())
	if err != nil {
		return err
	}
	return w.Clusters.PermanentDeleteByClusterId(ctx, d.Id())
}

func init() {
	common.RegisterResourceProvider(compute.ClusterSpec{}, ClusterSpec{})
	common.RegisterResourceProvider(compute.Library{}, LibraryResource{})
}
