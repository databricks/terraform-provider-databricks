package cluster

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/compute_tf"
	"github.com/databricks/terraform-provider-databricks/libraries"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
)

const resourceName = "cluster"
const DefaultProvisionTimeout = 30 * time.Minute
const DbfsDeprecationWarning = "For init scripts use 'volumes', 'workspace' or cloud storage location instead of 'dbfs'."

var clusterSchemaVersion = 4 // TODO: Is this needed?

const (
	numWorkerErr                              = "NumWorkers could be 0 only for SingleNode clusters. See https://docs.databricks.com/clusters/single-node.html for more details"
	unsupportedExceptCreateEditClusterSpecErr = "unsupported type %T, must be one of %scompute.CreateCluster, %scompute.ClusterSpec or %scompute.EditCluster. Please report this issue to the GitHub repo"
	createTimeout                             = DefaultProvisionTimeout
	updateTimeout                             = DefaultProvisionTimeout
	deleteTimeout                             = DefaultProvisionTimeout
)

var _ resource.ResourceWithConfigure = &ClusterResource{}
var _ resource.ResourceWithUpgradeState = &ClusterResource{}

type ClusterSpecExtended struct {
	compute_tf.ClusterSpec
	Libraries              []compute_tf.Library `tfsdk:"libraries" tf:"optional,slice_set,alias:library"`
	DefaultTags            types.Map            `tfsdk:"default_tags" tf:"computed"`
	IsPinned               types.Bool           `tfsdk:"is_pinned" tf:"optional"`
	NoWait                 types.Bool           `tfsdk:"no_wait" tf:"optional"`
	State                  types.String         `tfsdk:"state" tf:"computed"`
	Url                    types.String         `tfsdk:"url" tf:"computed"`
	AutoterminationMinutes types.Int32          `tfsdk:"autotermination_minutes" tf:"optional"`
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

func SparkConfDiffSuppressFunc(k, old, new string) bool {
	isPossiblyLegacyConfig := (k == "spark_conf.%" && old == "1" && new == "0")
	isLegacyConfig := (k == "spark_conf.spark.databricks.delta.preview.enabled")
	if isPossiblyLegacyConfig || isLegacyConfig {
		log.Printf("[DEBUG] Suppressing diff for k=%#v old=%#v new=%#v", k, old, new)
		return true
	}
	return false
}

func ZoneDiffSuppress(k, old, new string) bool {
	if old != "" && (new == "auto" || new == "") {
		log.Printf("[INFO] Suppressing diff on availability zone")
		return true
	}
	return false
}

// This method is a duplicate of Validate() in clusters/clusters_api.go that uses Go SDK.
// Long term, Validate() in clusters_api.go will be removed once all the resources using clusters are migrated to Go SDK.
func Validate(cluster any) error {
	var profile, master, resourceClass string
	switch c := cluster.(type) {
	case compute.CreateCluster:
		if c.NumWorkers > 0 || c.Autoscale != nil {
			return nil
		}
		profile = c.SparkConf["spark.databricks.cluster.profile"]
		master = c.SparkConf["spark.master"]
		resourceClass = c.CustomTags["ResourceClass"]
	case compute.EditCluster:
		if c.NumWorkers > 0 || c.Autoscale != nil {
			return nil
		}
		profile = c.SparkConf["spark.databricks.cluster.profile"]
		master = c.SparkConf["spark.master"]
		resourceClass = c.CustomTags["ResourceClass"]
	case compute.ClusterSpec:
		if c.NumWorkers > 0 || c.Autoscale != nil {
			return nil
		}
		profile = c.SparkConf["spark.databricks.cluster.profile"]
		master = c.SparkConf["spark.master"]
		resourceClass = c.CustomTags["ResourceClass"]
	default:
		return fmt.Errorf(unsupportedExceptCreateEditClusterSpecErr, cluster, "", "", "")
	}
	if profile == "singleNode" && strings.HasPrefix(master, "local") && resourceClass == "SingleNode" {
		return nil
	}
	return errors.New(numWorkerErr)
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

func SetForceSendFieldsForCluster(cluster any) error {
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

func setPinnedStatus(ctx context.Context, clusterId string, clusterAPI compute.ClustersInterface, resp *resource.ReadResponse) {
	events, err := clusterAPI.EventsAll(ctx, compute.GetEvents{
		ClusterId:  clusterId,
		Limit:      1,
		Order:      compute.GetEventsOrderDesc,
		EventTypes: []compute.EventType{compute.EventTypePinned, compute.EventTypeUnpinned},
	})
	if err != nil {
		resp.Diagnostics.AddError("failed to get cluster events", err.Error())
		return
	}
	pinnedEvent := compute.EventTypeUnpinned
	if len(events) > 0 {
		pinnedEvent = events[0].Type
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("is_pinned"), pinnedEvent == compute.EventTypePinned)...)
}

func hasClusterConfigChanged(req resource.UpdateRequest, resp *resource.UpdateResponse) bool {
	for k := range req.Config {
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

func ResourceCluster() resource.Resource {
	return &ClusterResource{}
}

type ClusterResource struct {
	Client *common.DatabricksClient
}

func (r *ClusterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksStagingName(resourceName)
}

// TODO
func (r *ClusterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ClusterSpecExtended{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		return c
	})
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Cluster",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *ClusterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if r.Client == nil && req.ProviderData != nil {
		r.Client = pluginfwcommon.ConfigureResource(req, resp)
	}
}

func (r *ClusterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("cluster_id"), req, resp)
}

func (r *ClusterResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		3: {StateUpgrader: removeZeroAwsEbsVolumeAttributes},
	}
}

func (r *ClusterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx, cancel := context.WithTimeout(ctx, createTimeout)
	defer cancel()
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)
	w, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	start := time.Now()
	timeout := d.Timeout(schema.TimeoutCreate)
	w, err := c.WorkspaceClient()
	if err != nil {
		return err
	}
	clusters := w.Clusters
	var createClusterRequest compute.CreateCluster
	common.DataToStructPointer(d, clusterSchema, &createClusterRequest)
	if err := Validate(createClusterRequest); err != nil {
		return err
	}
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

func (r *ClusterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)
	w, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var clusterId types.String
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("cluster_id"), &clusterId)...)
	if resp.Diagnostics.HasError() {
		return
	}
	clusterAPI := w.Clusters
	clusterInfo, err := clusterAPI.GetByClusterId(ctx, clusterId.ValueString())
	if err != nil {
		err = clusters.wrapMissingClusterError(err, clusterId.ValueString())
		resp.Diagnostics.AddError("failed to get cluster", err.Error())
		return
	}
	if err = common.StructToData(clusterInfo, clusterSchema, d); err != nil {
		return err
	}
	setPinnedStatus(ctx, clusterId.ValueString(), clusterAPI, resp)
	if resp.Diagnostics.HasError() {
		return
	}

	d.Set("url", c.FormatURL("#setting/clusters/", d.Id(), "/configuration"))
	shouldSkipLibrariesRead := !common.IsExporter(ctx)
	if d.Get("library.#").(int) == 0 && shouldSkipLibrariesRead {
		// don't add externally added libraries, if config has no `library {}` blocks
		// TODO: check if it still works fine with importing. Perhaps os.Setenv will do the trick
		return nil
	}

	libsClusterStatus, err := libraries.WaitForLibrariesInstalledSdk(ctx, w, compute.Wait{
		ClusterID: clusterId.ValueString(),
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

func (r *ClusterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx, cancel := context.WithTimeout(ctx, updateTimeout)
	defer cancel()
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)
	w, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	clusters := w.Clusters
	var cluster compute.EditCluster
	common.DataToStructPointer(d, clusterSchema, &cluster)
	clusterId := d.Id()
	cluster.ClusterId = clusterId
	var clusterInfo *compute.ClusterDetails

	if hasClusterConfigChanged(req, resp) {
		log.Printf("[DEBUG] Cluster state has changed!")
		if err := Validate(cluster); err != nil {
			return err
		}
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

func (r *ClusterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx, cancel := context.WithTimeout(ctx, deleteTimeout)
	defer cancel()
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)
	w, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var clusterId types.String
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("cluster_id"), &clusterId)...)
	if resp.Diagnostics.HasError() {
		return
	}
	err := w.Clusters.PermanentDeleteByClusterId(ctx, clusterId.ValueString())
	if err == nil {
		return
	}
	if !strings.Contains(err.Error(), "unpin the cluster first") {
		resp.Diagnostics.AddError("failed to delete cluster", err.Error())
		return
	}
	err = w.Clusters.UnpinByClusterId(ctx, clusterId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to unpin cluster", err.Error())
		return
	}
	err = w.Clusters.PermanentDeleteByClusterId(ctx, clusterId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to delete cluster", err.Error())
		return
	}
}
