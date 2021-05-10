package compute

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/databrickslabs/terraform-provider-databricks/common"
)

// DefaultProvisionTimeout ...
const DefaultProvisionTimeout = 30 * time.Minute

var clusterSchema = resourceClusterSchema()

// ResourceCluster - returns Cluster resource description
func ResourceCluster() *schema.Resource {
	return common.Resource{
		Create: resourceClusterCreate,
		Read:   resourceClusterRead,
		Update: resourceClusterUpdate,
		Delete: func(ctx context.Context,
			d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewClustersAPI(ctx, c).PermanentDelete(d.Id())
		},
		Schema:        clusterSchema,
		SchemaVersion: 2,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(DefaultProvisionTimeout),
			Update: schema.DefaultTimeout(DefaultProvisionTimeout),
			Delete: schema.DefaultTimeout(DefaultProvisionTimeout),
		},
	}.ToResource()
}

func sparkConfDiffSuppressFunc(k, old, new string, d *schema.ResourceData) bool {
	isPossiblyLegacyConfig := k == "spark_conf.%" && old == "1" && new == "0"
	isLegacyConfig := k == "spark_conf.spark.databricks.delta.preview.enabled"
	if isPossiblyLegacyConfig || isLegacyConfig {
		log.Printf("[DEBUG] Suppressing diff for k=%#v old=%#v new=%#v", k, old, new)
		return true
	}
	return false
}

func resourceClusterSchema() map[string]*schema.Schema {
	return common.StructToSchema(Cluster{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		s["spark_conf"].DiffSuppressFunc = sparkConfDiffSuppressFunc
		// adds `libraries` configuration block
		s["library"] = common.StructToSchema(ClusterLibraryList{},
			func(ss map[string]*schema.Schema) map[string]*schema.Schema {
				return ss
			})["library"]

		p, err := common.SchemaPath(s, "docker_image", "basic_auth", "password")
		if err == nil {
			p.Sensitive = true
		}
		s["autotermination_minutes"].Default = 60
		s["idempotency_token"].ForceNew = true
		s["cluster_id"] = &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		}
		s["aws_attributes"].ConflictsWith = []string{"azure_attributes", "gcp_attributes"}
		s["azure_attributes"].ConflictsWith = []string{"aws_attributes", "gcp_attributes"}
		s["gcp_attributes"].ConflictsWith = []string{"aws_attributes", "azure_attributes"}
		s["aws_attributes"].DiffSuppressFunc = makeEmptyBlockSuppressFunc("aws_attributes.#")
		s["azure_attributes"].DiffSuppressFunc = makeEmptyBlockSuppressFunc("azure_attributes.#")
		s["gcp_attributes"].DiffSuppressFunc = makeEmptyBlockSuppressFunc("gcp_attributes.#")

		s["is_pinned"] = &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				if old == "" && new == "false" {
					return true
				}
				return old == new
			},
		}
		s["state"] = &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		}
		s["default_tags"] = &schema.Schema{
			Type:     schema.TypeMap,
			Computed: true,
		}
		s["num_workers"] = &schema.Schema{
			Type:             schema.TypeInt,
			Optional:         true,
			Default:          0,
			ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(0)),
		}
		s["url"] = &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		}
		return s
	})
}

func validateClusterDefinition(cluster Cluster) error {
	if cluster.NumWorkers > 0 || cluster.Autoscale != nil {
		return nil
	}
	profile := cluster.SparkConf["spark.databricks.cluster.profile"]
	master := cluster.SparkConf["spark.master"]
	resourceClass := cluster.CustomTags["ResourceClass"]
	if profile == "singleNode" && strings.HasPrefix(master, "local") && resourceClass == "SingleNode" {
		return nil
	}
	return fmt.Errorf("NumWorkers could be 0 only for SingleNode clusters. See https://docs.databricks.com/clusters/single-node.html for more details")
}

func resourceClusterCreate(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
	var cluster Cluster
	clusters := NewClustersAPI(ctx, c)
	err := common.DataToStructPointer(d, clusterSchema, &cluster)
	if err != nil {
		return err
	}
	if err = validateClusterDefinition(cluster); err != nil {
		return err
	}
	modifyClusterRequest(&cluster)
	clusterInfo, err := clusters.Create(cluster)
	if err != nil {
		return err
	}
	d.SetId(clusterInfo.ClusterID)
	d.Set("cluster_id", clusterInfo.ClusterID)
	isPinned, ok := d.GetOk("is_pinned")
	if ok && isPinned.(bool) {
		err = clusters.Pin(clusterInfo.ClusterID)
		if err != nil {
			return err
		}
	}
	var libraryList ClusterLibraryList
	if err = common.DataToStructPointer(d, clusterSchema, &libraryList); err != nil {
		return err
	}
	librariesAPI := NewLibrariesAPI(ctx, c)
	if len(libraryList.Libraries) > 0 {
		if err = librariesAPI.Install(libraryList); err != nil {
			return err
		}
		if _, err := waitForLibrariesInstalled(librariesAPI, clusterInfo); err != nil {
			return err
		}
	}
	return nil
}

func setPinnedStatus(d *schema.ResourceData, clusterAPI ClustersAPI) error {
	events, err := clusterAPI.Events(EventsRequest{
		ClusterID:  d.Id(),
		Limit:      1,
		Order:      SortDescending,
		EventTypes: []ClusterEventType{EvTypePinned, EvTypeUnpinned},
		MaxItems:   1,
	})
	if err != nil {
		return err
	}
	pinnedEvent := EvTypeUnpinned
	if len(events) > 0 {
		pinnedEvent = events[0].Type
	}
	return d.Set("is_pinned", pinnedEvent == EvTypePinned)
}

func resourceClusterRead(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
	clusterAPI := NewClustersAPI(ctx, c)
	clusterInfo, err := clusterAPI.Get(d.Id())
	if err != nil {
		return err
	}
	if err = common.StructToData(clusterInfo, clusterSchema, d); err != nil {
		return err
	}
	if err = setPinnedStatus(d, clusterAPI); err != nil {
		return err
	}
	d.Set("url", c.FormatURL("#setting/clusters/", d.Id(), "/configuration"))
	librariesAPI := NewLibrariesAPI(ctx, c)
	libsClusterStatus, err := waitForLibrariesInstalled(librariesAPI, clusterInfo)
	if err != nil {
		return err
	}
	libList := libsClusterStatus.ToLibraryList()
	return common.StructToData(libList, clusterSchema, d)
}

func waitForLibrariesInstalled(
	libraries LibrariesAPI, clusterInfo ClusterInfo) (result *ClusterLibraryStatuses, err error) {
	err = resource.RetryContext(libraries.context, 30*time.Minute, func() *resource.RetryError {
		libsClusterStatus, err := libraries.ClusterStatus(clusterInfo.ClusterID)
		if ae, ok := err.(common.APIError); ok && ae.IsMissing() {
			// eventual consistency error
			return resource.RetryableError(err)
		}
		if err != nil {
			return resource.NonRetryableError(err)
		}
		if !clusterInfo.IsRunningOrResizing() {
			log.Printf("[INFO] Cluster %#v (%s) is currently not running, so just returning list of %d libraries",
				clusterInfo.ClusterName, clusterInfo.ClusterID, len(libsClusterStatus.LibraryStatuses))
			result = &libsClusterStatus
			return nil
		}
		retry, err := libsClusterStatus.IsRetryNeeded()
		if retry {
			return resource.RetryableError(err)
		}
		if err != nil {
			return resource.NonRetryableError(err)
		}
		result = &libsClusterStatus
		return nil
	})
	return
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
	clusters := NewClustersAPI(ctx, c)
	clusterID := d.Id()
	cluster := Cluster{ClusterID: clusterID}
	err := common.DataToStructPointer(d, clusterSchema, &cluster)
	if err != nil {
		return err
	}
	var clusterInfo ClusterInfo
	if hasClusterConfigChanged(d) {
		log.Printf("[DEBUG] Cluster state has changed!")
		err = validateClusterDefinition(cluster)
		if err != nil {
			return err
		}
		modifyClusterRequest(&cluster)
		clusterInfo, err = clusters.Edit(cluster)
		if err != nil {
			return err
		}
	} else {
		clusterInfo, err = clusters.Get(clusterID)
		if err != nil {
			return err
		}
	}
	oldPinned, newPinned := d.GetChange("is_pinned")
	if oldPinned.(bool) != newPinned.(bool) {
		log.Printf("[DEBUG] Update: is_pinned. Old: %v, New: %v", oldPinned, newPinned)
		if newPinned.(bool) {
			err = clusters.Pin(clusterID)
		} else {
			err = clusters.Unpin(clusterID)
		}
		if err != nil {
			return err
		}
	}

	var libraryList ClusterLibraryList
	if err = common.DataToStructPointer(d, clusterSchema, &libraryList); err != nil {
		return err
	}
	librariesAPI := NewLibrariesAPI(ctx, c)
	libsClusterStatus, err := librariesAPI.ClusterStatus(clusterID)
	if err != nil {
		return err
	}
	libraryList.ClusterID = clusterID
	libsToInstall, libsToUninstall := libraryList.Diff(libsClusterStatus)
	if len(libsToUninstall.Libraries) > 0 || len(libsToInstall.Libraries) > 0 {
		tmpClusterInfo := clusterInfo
		if !clusterInfo.IsRunningOrResizing() {
			tmpClusterInfo, err = clusters.StartAndGetInfo(clusterID)
			if err != nil {
				return err
			}
		}
		if err = updateLibraries(librariesAPI, tmpClusterInfo, libsToInstall, libsToUninstall); err != nil {
			return err
		}
		if clusterInfo.State == ClusterStateTerminated {
			log.Printf("[INFO] %s was in TERMINATED state, so terminating it again", clusterID)
			if err = clusters.Terminate(clusterID); err != nil {
				return err
			}
		}
	}
	return nil
}

// modifyClusterRequest helps remove all request fields that should not be submitted when instance pool is selected.
func modifyClusterRequest(clusterModel *Cluster) {
	// Instance profile id does not exist or not set
	if clusterModel.InstancePoolID == "" {
		return
	}
	if clusterModel.AwsAttributes != nil {
		// Reset AwsAttributes
		awsAttributes := AwsAttributes{
			InstanceProfileArn: clusterModel.AwsAttributes.InstanceProfileArn,
		}
		clusterModel.AwsAttributes = &awsAttributes
	}
	if clusterModel.AzureAttributes != nil {
		clusterModel.AzureAttributes = nil
	}
	if clusterModel.GcpAttributes != nil {
		gcpAttributes := GcpAttributes{
			GoogleServiceAccount: clusterModel.GcpAttributes.GoogleServiceAccount,
		}
		clusterModel.GcpAttributes = &gcpAttributes
	}
	clusterModel.EnableElasticDisk = false
	clusterModel.NodeTypeID = ""
	clusterModel.DriverNodeTypeID = ""
}

func updateLibraries(libraries LibrariesAPI, clusterInfo ClusterInfo,
	libsToInstall, libsToUninstall ClusterLibraryList) error {
	if len(libsToUninstall.Libraries) > 0 {
		err := libraries.Uninstall(libsToUninstall)
		if err != nil {
			return err
		}
	}
	if len(libsToInstall.Libraries) > 0 {
		err := libraries.Install(libsToInstall)
		if err != nil {
			return err
		}
	}
	_, err := waitForLibrariesInstalled(libraries, clusterInfo)
	return err
}
