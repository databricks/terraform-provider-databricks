package clusters

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/libraries"
)

// DefaultProvisionTimeout ...
const DefaultProvisionTimeout = 30 * time.Minute

const DbfsDeprecationWarning = "For init scripts use 'volumes', 'workspace' or cloud storage location instead of 'dbfs'."

var clusterSchema = resourceClusterSchemaProvider()

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

type ClusterResourceProvider struct{}

func (ClusterResourceProvider) UnderlyingType() compute.ClusterDetails {
	return compute.ClusterDetails{}
}

func (ClusterResourceProvider) Aliases() map[string]string {
	return map[string]string{"cluster_mount_infos": "cluster_mount_info"}
}

func makeEmptyBlockSuppressFunc() func(k, old, new string, d *schema.ResourceData) bool {
	return func(k, old, new string, d *schema.ResourceData) bool {
		if strings.HasSuffix(k, ".#") && old == "1" && new == "0" {
			log.Printf("[DEBUG] Suppressing diff for k=%#v platform=%#v config=%#v", k, old, new)
			return true
		}
		return false
	}
}

func diffSuppressor(zero string) func(k, old, new string, d *schema.ResourceData) bool {
	return func(k, old, new string, d *schema.ResourceData) bool {
		if new == zero && old != zero {
			log.Printf("[DEBUG] Suppressing diff for %v: platform=%#v config=%#v", k, old, new)
			return true
		}
		return false
	}
}

func (ClusterResourceProvider) TfOverlay() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enable_elastic_disk": {
			Computed: true,
		},
		"enable_local_disk_encryption": {
			Computed: true,
		},
		"node_type_id": {
			ConflictsWith: []string{"driver_instance_pool_id", "instance_pool_id"},
			Computed:      true,
		},
		"driver_node_type_id": {
			ConflictsWith: []string{"driver_instance_pool_id", "instance_pool_id"},
			Computed:      true,
		},
		"driver_instance_pool_id": {
			ConflictsWith: []string{"driver_node_type_id", "node_type_id"},
			Computed:      true,
		},
		"ssh_public_keys": {
			MaxItems: 10,
		},
		"init_scripts": {
			MaxItems: 10,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"dbfs": {
						Deprecated: DbfsDeprecationWarning,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"destination": {
									Required: true,
								},
							},
						},
					},
					"s3": {
						Deprecated: DbfsDeprecationWarning,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"destination": {
									Required: true,
								},
							},
						},
					},
				},
			},
		},
		"workload_type": {
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"clients": {
						MinItems: 5,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"notebooks": {
									Optional: true,
									Default:  true,
								},
								"jobs": {
									Optional: true,
									Default:  true,
								},
							},
						},
					},
				},
			},
		},
		"idempotency_token": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		"data_security_mode": {
			DiffSuppressFunc: diffSuppressor(""),
		},
		"docker_image": {
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"url": {
						Required: true,
					},
					"basic_auth": {
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"password": {
									Sensitive: true,
									Required:  true,
								},
								"username": {
									Required: true,
								},
							},
						},
					},
				},
			},
		},
		"spark_conf": {
			DiffSuppressFunc: SparkConfDiffSuppressFunc,
		},
		"aws_attributes": {
			DiffSuppressFunc: makeEmptyBlockSuppressFunc(),
			ConflictsWith:    []string{"azure_attributes", "gcp_attributes"},
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"zone_id": {
						DiffSuppressFunc: ZoneDiffSuppress,
					},
				},
			},
		},
		"azure_attributes": {
			DiffSuppressFunc: makeEmptyBlockSuppressFunc(),
			ConflictsWith:    []string{"aws_attributes", "gcp_attributes"},
		},
		"gcp_attributes": {
			DiffSuppressFunc: makeEmptyBlockSuppressFunc(),
			ConflictsWith:    []string{"azure_attributes", "aws_attributes"},
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"use_preemptible_executors": {
						Type:       schema.TypeBool,
						Optional:   true,
						Deprecated: "Please use 'availability' instead.",
					},
					"zone_id": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		"library": common.StructToSchema(libraries.ClusterLibraryList{},
			func(ss map[string]*schema.Schema) map[string]*schema.Schema {
				ss["library"].Set = func(i any) int {
					lib := libraries.NewLibraryFromInstanceState(i)
					return schema.HashString(lib.String())
				}
				return ss
			})["library"],
		"autotermination_minutes": {
			Default: 60,
		},
		"autoscale": {
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"max_workers": {
						Optional: true,
					},
					"min_workers": {
						Optional: true,
					},
				},
			},
		},
		"apply_policy_default_values": {
			Optional: true,
			Type:     schema.TypeBool,
		},
		"cluster_log_conf": {
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"dbfs": {
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"destination": {
									Required: true,
								},
							},
						},
					},
					"s3": {
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"destination": {
									Required: true,
								},
							},
						},
					},
				},
			},
		},
		"spark_version": {
			Required: true,
		},
		"cluster_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"instance_pool_id": {
			ConflictsWith: []string{"driver_node_type_id", "node_type_id"},
		},
		"runtime_engine": {
			ValidateFunc: validation.StringInSlice([]string{"PHOTON", "STANDARD"}, false),
		},
		"is_pinned": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				if old == "" && new == "false" {
					return true
				}
				return old == new
			},
		},
		"state": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"default_tags": {
			Type:     schema.TypeMap,
			Computed: true,
		},
		"num_workers": {
			Type:             schema.TypeInt,
			Optional:         true,
			Default:          0,
			ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(0)),
		},
		"url": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func resourceClusterSchemaProvider() map[string]*schema.Schema {
	return common.ResourceProviderStructToSchema[compute.ClusterDetails](ClusterResourceProvider{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		return s
	})
}

// func resourceClusterSchema() map[string]*schema.Schema {
// 	return common.StructToSchema(Cluster{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
// 		s["spark_conf"].DiffSuppressFunc = SparkConfDiffSuppressFunc
// 		common.MustSchemaPath(s, "aws_attributes", "zone_id").DiffSuppressFunc = ZoneDiffSuppress
// 		common.MustSchemaPath(s, "gcp_attributes", "use_preemptible_executors").Deprecated = "Please use 'availability' instead."

// 		common.MustSchemaPath(s, "init_scripts", "dbfs").Deprecated = DbfsDeprecationWarning

// 		// adds `library` configuration block
// 		s["library"] = common.StructToSchema(libraries.ClusterLibraryList{},
// 			func(ss map[string]*schema.Schema) map[string]*schema.Schema {
// 				ss["library"].Set = func(i any) int {
// 					lib := libraries.NewLibraryFromInstanceState(i)
// 					return schema.HashString(lib.String())
// 				}
// 				return ss
// 			})["library"]

// 		s["autotermination_minutes"].Default = 60
// 		s["cluster_id"] = &schema.Schema{
// 			Type:     schema.TypeString,
// 			Optional: true,
// 			Computed: true,
// 		}
// 		s["aws_attributes"].ConflictsWith = []string{"azure_attributes", "gcp_attributes"}
// 		s["azure_attributes"].ConflictsWith = []string{"aws_attributes", "gcp_attributes"}
// 		s["gcp_attributes"].ConflictsWith = []string{"aws_attributes", "azure_attributes"}
// 		s["instance_pool_id"].ConflictsWith = []string{"driver_node_type_id", "node_type_id"}
// 		s["driver_instance_pool_id"].ConflictsWith = []string{"driver_node_type_id", "node_type_id"}
// 		s["driver_node_type_id"].ConflictsWith = []string{"driver_instance_pool_id", "instance_pool_id"}
// 		s["node_type_id"].ConflictsWith = []string{"driver_instance_pool_id", "instance_pool_id"}

// 		s["runtime_engine"].ValidateFunc = validation.StringInSlice([]string{"PHOTON", "STANDARD"}, false)

// 		s["is_pinned"] = &schema.Schema{
// 			Type:     schema.TypeBool,
// 			Optional: true,
// 			Default:  false,
// 			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
// 				if old == "" && new == "false" {
// 					return true
// 				}
// 				return old == new
// 			},
// 		}
// 		s["state"] = &schema.Schema{
// 			Type:     schema.TypeString,
// 			Computed: true,
// 		}
// 		s["default_tags"] = &schema.Schema{
// 			Type:     schema.TypeMap,
// 			Computed: true,
// 		}
// 		s["num_workers"] = &schema.Schema{
// 			Type:             schema.TypeInt,
// 			Optional:         true,
// 			Default:          0,
// 			ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(0)),
// 		}
// 		s["url"] = &schema.Schema{
// 			Type:     schema.TypeString,
// 			Computed: true,
// 		}
// 		return s
// 	})
// }

func resourceClusterCreate(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
	var cluster Cluster
	start := time.Now()
	timeout := d.Timeout(schema.TimeoutCreate)
	clusters := NewClustersAPI(ctx, c)
	common.DataToStructPointer(d, clusterSchema, &cluster)
	if err := cluster.Validate(); err != nil {
		return err
	}
	cluster.ModifyRequestOnInstancePool()
	// TODO: propagate d.Timeout(schema.TimeoutCreate)
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
	var libraryList libraries.ClusterLibraryList
	common.DataToStructPointer(d, clusterSchema, &libraryList)
	libs := libraries.NewLibrariesAPI(ctx, c)
	if len(libraryList.Libraries) > 0 {
		if err = libs.Install(libraryList); err != nil {
			return err
		}
		_, err := libs.WaitForLibrariesInstalled(libraries.Wait{
			ClusterID: d.Id(),
			Timeout:   timeout - time.Since(start),
			IsRunning: clusterInfo.IsRunningOrResizing(),
			IsRefresh: false,
		})
		if err != nil {
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
	shouldSkipLibrariesRead := !common.IsExporter(ctx)
	if d.Get("library.#").(int) == 0 && shouldSkipLibrariesRead {
		// don't add externally added libraries, if config has no `library {}` blocks
		// TODO: check if it still works fine with importing. Perhaps os.Setenv will do the trick
		return nil
	}
	librariesAPI := libraries.NewLibrariesAPI(ctx, c)
	libsClusterStatus, err := librariesAPI.WaitForLibrariesInstalled(libraries.Wait{
		ClusterID: d.Id(),
		Timeout:   d.Timeout(schema.TimeoutRead),
		IsRunning: clusterInfo.IsRunningOrResizing(),
		IsRefresh: true,
	})
	if err != nil {
		return err
	}
	libList := libsClusterStatus.ToLibraryList()
	return common.StructToData(libList, clusterSchema, d)
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
	common.DataToStructPointer(d, clusterSchema, &cluster)
	var clusterInfo ClusterInfo
	var err error

	if hasClusterConfigChanged(d) {
		log.Printf("[DEBUG] Cluster state has changed!")
		if err := cluster.Validate(); err != nil {
			return err
		}
		cluster.ModifyRequestOnInstancePool()
		cluster.FixInstancePoolChangeIfAny(d)

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
		clusterInfo, err = clusters.Get(clusterID)
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
			clusterInfo, err = clusters.Resize(ResizeRequest{
				ClusterID:  clusterID,
				NumWorkers: cluster.NumWorkers,
			})
		} else if isAutoscaleConfigResizeForAutoscalingCluster ||
			isNonAutoScalingToAutoscalingResize {
			clusterInfo, err = clusters.Resize(ResizeRequest{
				ClusterID: clusterID,
				AutoScale: cluster.Autoscale,
			})
		} else {
			clusterInfo, err = clusters.Edit(cluster)
		}
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
	oldNumLibs, newNumLibs := d.GetChange("library.#")
	if oldNumLibs == newNumLibs && oldNumLibs.(int) == 0 {
		// don't add externally added libraries, if config has no `library {}` blocks
		return nil
	}
	var libraryList libraries.ClusterLibraryList
	common.DataToStructPointer(d, clusterSchema, &libraryList)
	librariesAPI := libraries.NewLibrariesAPI(ctx, c)
	libsClusterStatus, err := librariesAPI.ClusterStatus(clusterID)
	if err != nil {
		return err
	}
	libraryList.ClusterID = clusterID
	libsToInstall, libsToUninstall := libraryList.Diff(libsClusterStatus)
	if len(libsToUninstall.Libraries) > 0 || len(libsToInstall.Libraries) > 0 {
		if !clusterInfo.IsRunningOrResizing() {
			if _, err = clusters.StartAndGetInfo(clusterID); err != nil {
				return err
			}
		}
		// clusters.StartAndGetInfo() always returns a running cluster
		// or errors out, so we just know the cluster is active.
		err = librariesAPI.UpdateLibraries(clusterID, libsToInstall, libsToUninstall,
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
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
