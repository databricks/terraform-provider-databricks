package databricks

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
)

func resourceCluster() *schema.Resource {
	clusterSchema := clusterSchema()
	clusterSchema["library_jar"] = librarySchema("path")
	clusterSchema["library_egg"] = librarySchema("path")
	clusterSchema["library_whl"] = librarySchema("path")
	clusterSchema["library_pypi"] = librarySchema("package", "repo")
	clusterSchema["library_cran"] = librarySchema("package", "repo")
	return &schema.Resource{
		Create: resourceClusterCreate,
		Read:   resourceClusterRead,
		Update: resourceClusterUpdate,
		Delete: resourceClusterDelete,
		Schema: clusterSchema,
		// see usage at https://github.com/terraform-providers/terraform-provider-aws/blob/9900515b28b413505e5dd7c8d8ea258c59515f32/aws/resource_aws_vpc_peering_connection.go#L288
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
	}
}

func storageInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dbfs": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"s3": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						//TODO: Validate that destination has s3:// prefix
						"destination": {
							Type:     schema.TypeString,
							Required: true,
						},
						"region": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"endpoint": {
							Type:     schema.TypeString,
							Optional: true,
							// AtLeastOneOf: []string{
							// 	"cluster_log_conf.0.s3.0.region",
							// 	"cluster_log_conf.0.s3.0.endpoint",
							// },
						},
						"enable_encryption": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"encryption_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"kms_key": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"canned_acl": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
				// ExactlyOneOf: []string{
				// 	"cluster_log_conf.0.dbfs",
				// 	"cluster_log_conf.0.s3",
				// },
			},
		},
	}
}

func librarySchema(dims ...string) *schema.Schema {
	fields := map[string]*schema.Schema{
		"messages": {
			// consider removing it...
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Computed: true,
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
	for _, dim := range dims {
		fields[dim] = &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		}
	}
	return &schema.Schema{
		Type:       schema.TypeSet,
		Optional:   true,
		ConfigMode: schema.SchemaConfigModeAttr,
		Elem: &schema.Resource{
			Schema: fields,
		},
	}
}

func clusterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"num_workers": {
			Type:          schema.TypeInt,
			Optional:      true,
			ConflictsWith: []string{"autoscale"},
		},
		"autoscale": {
			Type:       schema.TypeSet,
			Optional:   true,
			MaxItems:   1,
			ConfigMode: schema.SchemaConfigModeAttr,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"min_workers": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"max_workers": {
						Type:     schema.TypeInt,
						Optional: true,
					},
				},
			},
			ConflictsWith: []string{"num_workers"},
		},
		"cluster_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"spark_version": {
			Type:     schema.TypeString,
			Required: true,
		},
		"spark_conf": {
			Type:     schema.TypeMap,
			Optional: true,
		},
		"aws_attributes": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"availability": {
						Type:     schema.TypeString,
						Optional: true,
						Default:  "SPOT_WITH_FALLBACK",
					},
					"zone_id": {
						Type:     schema.TypeString,
						Required: true,
					},
					"spot_bid_price_percent": {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  "100",
					},
					"instance_profile_arn": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"first_on_demand": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"ebs_volume_type": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"ebs_volume_count": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"ebs_volume_size": {
						Type:     schema.TypeInt,
						Optional: true,
					},
				},
			},
		},
		"driver_node_type_id": {
			Type:          schema.TypeString,
			Optional:      true,
			Computed:      true,
			ConflictsWith: []string{"instance_pool_id"},
		},
		"node_type_id": {
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"instance_pool_id"},
			AtLeastOneOf:  []string{"instance_pool_id"},
		},
		"ssh_public_keys": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
			//	TODO: Validate less than 10 values
		},
		"custom_tags": {
			Type:     schema.TypeMap,
			Optional: true,
		},
		"cluster_log_conf": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			//ConfigMode: schema.SchemaConfigModeAttr,
			Elem: storageInfoSchema(),
		},
		"init_scripts": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 10,
			Elem:     storageInfoSchema(),
		},
		"docker_image": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"url": {
						Type:     schema.TypeString,
						Required: true,
					},
					"basic_auth": {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"username": {
									Type:     schema.TypeString,
									Required: true,
								},
								"password": {
									Type:      schema.TypeString,
									Required:  true,
									Sensitive: true,
								},
							},
						},
					},
				},
			},
		},
		"spark_env_vars": {
			Type:     schema.TypeMap,
			Optional: true,
		},
		"autotermination_minutes": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
			//Computed: true,
		},
		"enable_elastic_disk": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"instance_pool_id": {
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"node_type_id", "driver_node_type_id", "aws_attributes"},
			AtLeastOneOf:  []string{"node_type_id"},
		},
		"idempotency_token": {
			Type:     schema.TypeInt,
			Optional: true,
			ForceNew: true,
		},
		// here drops support for mvn exclusions, will re-add if someone will complain
		"library_maven": librarySchema("coordinates", "repo"),
		"cluster_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"default_tags": {
			Type:     schema.TypeMap,
			Computed: true,
		},
		"state": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"state_message": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"single_user_name": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func resourceClusterCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)
	clusters := client.Clusters()
	cluster := model.Cluster{}
	err := readStructFromData([]string{}, d, &cluster, resourceCluster())
	if err != nil {
		return err
	}
	clusterInfo, err := clusters.Create(cluster)
	if err != nil {
		return err
	}
	d.SetId(clusterInfo.ClusterID)
	err = clusters.WaitForClusterRunning(clusterInfo.ClusterID)
	if err != nil {
		return err
	}
	err = readClusterLibraryListFromData(d).Apply(client.Libraries().Install)
	if err != nil {
		return err
	}
	return resourceClusterRead(d, m)
}

func resourceClusterRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)
	libraries := client.Libraries()
	clusterID := d.Id()
	clusterInfo, err := client.Clusters().Get(clusterID)
	if isClusterMissing(err, clusterID) {
		log.Printf("Missing cluster with id: %s.", clusterID)
		d.SetId("")
		return nil
	}
	if err != nil {
		return err
	}
	libsClusterStatus, err := waitForLibrariesInstalled(libraries, clusterID)
	if err != nil {
		return err
	}
	err = libsClusterStatus.Apply(d.Set)
	if err != nil {
		return err
	}
	return readResourceFromStruct(clusterInfo, d, []string{}, resourceCluster())
}

func waitForLibrariesInstalled(
	libraries service.LibrariesAPI, clusterID string) (*model.ClusterLibraryStatuses, error) {
	var result *model.ClusterLibraryStatuses
	return result, resource.Retry(5*time.Minute, func() *resource.RetryError {
		libsClusterStatus, err := libraries.ClusterStatus(clusterID)
		if isClusterMissing(err, clusterID) {
			// eventual consistency error
			return resource.RetryableError(err)
		}
		if err != nil {
			return resource.NonRetryableError(err)
		}
		retry, err := libsClusterStatus.IsRetryNeeded()
		if retry {
			return resource.RetryableError(err)
		}
		if err != nil {
			return resource.NonRetryableError(err)
		}
		// TODO: let this be reviewed by someone. looks okay, though...
		result = &libsClusterStatus
		return nil
	})
}

func readClusterLibraryListFromData(d *schema.ResourceData) *model.ClusterLibraryList {
	cll := model.ClusterLibraryList{ClusterID: d.Id()}
	for _, n := range []string{"library_jar", "library_egg",
		"library_whl", "library_pypi", "library_maven",
		"library_cran"} {
		if libs, ok := d.GetOk(n); ok {
			for _, l := range libs.(*schema.Set).List() {
				cll.AddLibraryFromMap(n, l.(map[string]interface{}))
			}
		}
	}
	return &cll
}

func resourceClusterUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)
	clusterID := d.Id()
	clusters := client.Clusters()
	libraries := client.Libraries()
	clusterInfo, err := clusters.Get(clusterID)
	if err != nil {
		return err
	}
	libsInConfig := readClusterLibraryListFromData(d)
	libsClusterStatus, err := libraries.ClusterStatus(clusterID)
	if err != nil {
		return err
	}
	libsToInstall, libsToUninstall := libsInConfig.Diff(libsClusterStatus)
	switch clusterInfo.State {
	case model.ClusterStateTerminated:
		err = editCluster(clusters, d)
		if err != nil {
			return err
		}
		if len(libsToUninstall.Libraries) > 0 || len(libsToInstall.Libraries) > 0 {
			err = clusters.Start(clusterID)
			if err != nil {
				return err
			}
			err := updateLibraries(libraries, libsToInstall, libsToUninstall)
			if err != nil {
				return err
			}
			err = clusters.Terminate(clusterID)
			if err != nil {
				return err
			}
		}
		return resourceClusterRead(d, m)
	case model.ClusterStatePending, model.ClusterStateResizing:
		err = clusters.WaitForClusterRunning(clusterID)
		if err != nil {
			return err
		}
		fallthrough
	case model.ClusterStateRunning:
		err = editCluster(clusters, d)
		if err != nil {
			return err
		}
		err = clusters.WaitForClusterRunning(clusterID)
		if err != nil {
			return err
		}
		err = updateLibraries(libraries, libsToInstall, libsToUninstall)
		if err != nil {
			return err
		}
		return resourceClusterRead(d, m)
	}
	return fmt.Errorf("%s is not in runnable or terminated state: %s", clusterID, clusterInfo.State)
}

func editCluster(clusters service.ClustersAPI, d *schema.ResourceData) error {
	cluster := model.Cluster{ClusterID: d.Id()}
	err := readStructFromData([]string{}, d, &cluster, resourceCluster())
	if err != nil {
		return err
	}
	return clusters.Edit(cluster)
}

func updateLibraries(libraries service.LibrariesAPI, libsToInstall, libsToUninstall model.ClusterLibraryList) error {
	if len(libsToUninstall.Libraries) > 0 {
		err := libraries.Uninstall(libsToUninstall)
		if err != nil {
			return err
		}
		_, err = waitForLibrariesInstalled(libraries, libsToInstall.ClusterID)
		if err != nil {
			return err
		}
	}
	if len(libsToInstall.Libraries) > 0 {
		err := libraries.Install(libsToInstall)
		if err != nil {
			return err
		}
		_, err = waitForLibrariesInstalled(libraries, libsToUninstall.ClusterID)
		if err != nil {
			return err
		}
	}
	return nil
}

func resourceClusterDelete(d *schema.ResourceData, m interface{}) error {
	return m.(*service.DBApiClient).Clusters().PermanentDelete(d.Id())
}

func isClusterMissing(err error, resourceID string) bool {
	if apiErr, ok := err.(service.APIError); ok {
		return apiErr.IsMissing() || strings.Contains(err.Error(),
			fmt.Sprintf("Cluster %s does not exist", resourceID))
	}
	return false
}
