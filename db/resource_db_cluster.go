package db

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"sort"
	"strings"
	"time"
)

func resourceCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceClusterCreate,
		Read:   resourceClusterRead,
		Update: resourceClusterUpdate,
		Delete: resourceClusterDelete,

		Schema: map[string]*schema.Schema{
			"num_workers": &schema.Schema{
				Type:          schema.TypeInt,
				Optional:      true,
				ConflictsWith: []string{"autoscale"},
			},
			"autoscale": &schema.Schema{
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
			"cluster_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spark_version": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"spark_conf": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"aws_attributes": &schema.Schema{
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
			"driver_node_type_id": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"instance_pool_id"},
			},
			"node_type_id": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"instance_pool_id"},
			},
			"ssh_public_keys": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				//	TODO: Validate less than 10 values
			},
			"custom_tags": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"cluster_log_conf": &schema.Schema{
				Type:       schema.TypeSet,
				Optional:   true,
				MaxItems:   1,
				ConfigMode: schema.SchemaConfigModeAttr,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dbfs_destination": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"s3_destination": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"s3_region": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"s3_endpoint": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"s3_enable_encryption": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"s3_encryption_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"s3_kms_key": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"s3_canned_acl": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"init_scripts": &schema.Schema{
				Type:       schema.TypeSet,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dbfs_destination": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"s3_destination": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"s3_region": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"s3_endpoint": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"docker_image": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"url": {
							Type:     schema.TypeString,
							Required: true,
						},
						"username": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"password": {
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
						},
					},
				},
			},
			"spark_env_vars": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"autotermination_minutes": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
				//Computed: true,
			},
			"enable_elastic_disk": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"instance_pool_id": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"node_type_id", "driver_node_type_id", "aws_attributes"},
			},
			"idempotency_token": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"library_jar": &schema.Schema{
				Type:       schema.TypeSet,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"path": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"messages": {
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
					},
				},
			},
			"library_egg": &schema.Schema{
				Type:       schema.TypeSet,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"path": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"messages": {
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
					},
				},
			},
			"library_whl": &schema.Schema{
				Type:       schema.TypeSet,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"path": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"messages": {
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
					},
				},
			},
			"library_pypi": &schema.Schema{
				Type:       schema.TypeSet,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"package": {
							Type:     schema.TypeString,
							Required: true,
						},
						"repo": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"messages": {
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
					},
				},
			},
			"library_maven": &schema.Schema{
				Type:       schema.TypeSet,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"coordinates": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"repo": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"exclusions": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"messages": {
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
					},
				},
			},
			"library_cran": &schema.Schema{
				Type:       schema.TypeSet,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"package": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"repo": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"messages": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"cluster_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_tags": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"state_message": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		CustomizeDiff: customdiff.All(
			customdiff.ValidateValue("cluster_log_conf", validateClusterLogConf),
			customdiff.ValidateValue("init_scripts", validateInitScripts),
			customdiff.ComputedIf("aws_attributes", func(d *schema.ResourceDiff, meta interface{}) bool {
				log.Println("rdiff aws_attributes")
				log.Println(d.GetChange("aws_attributes"))
				return false
			}),
		),
	}
}

func convertListInterfaceToString(m []interface{}) []string {
	response := []string{}
	for _, v := range m {
		if v != nil {
			response = append(response, v.(string))
		}
	}
	return response
}

func resourceClusterCreate(d *schema.ResourceData, m interface{}) error {

	client := m.(service.DBApiClient)

	cluster := parseSchemaToCluster(d, "")
	libraries := parseSchemaToClusterLibraries(d)

	clusterInfo, err := client.Clusters().Create(cluster)
	if err != nil {
		return err
	}

	d.SetId(clusterInfo.ClusterID)

	if idempotencyToken, ok := d.GetOk("idempotency_token"); ok {
		err := d.Set("idempotency_token", idempotencyToken)
		if err != nil {
			return err
		}
	}

	err = client.Clusters().WaitForClusterRunning(clusterInfo.ClusterID, 30, 120)
	if err != nil {
		return err
	}

	if len(libraries) > 0 {
		err = client.Libraries().Create(clusterInfo.ClusterID, libraries)
	}

	return resourceClusterRead(d, m)
}

func resourceClusterRead(d *schema.ResourceData, m interface{}) error {
	client := m.(service.DBApiClient)
	id := d.Id()

	clusterInfo, err := client.Clusters().Get(id)
	if err != nil {
		if isClusterMissing(err.Error(), id) {
			log.Printf("Missing cluster with id: %s.", id)
			d.SetId("")
			return nil
		}
		return err
	}

	librariesStatuses, err := client.Libraries().List(id)

	var jars []map[string]interface{}
	var eggs []map[string]interface{}
	var whls []map[string]interface{}
	var pypi []map[string]interface{}
	var maven []map[string]interface{}
	var cran []map[string]interface{}

	for _, lib := range librariesStatuses {
		if len(lib.Library.Jar) > 0 {
			jarPackage := map[string]interface{}{}
			jarPackage["path"] = lib.Library.Jar
			jarPackage["messages"] = lib.Messages
			jarPackage["status"] = lib.Status
			jars = append(jars, jarPackage)
		}
		if len(lib.Library.Egg) > 0 {
			eggsPackage := map[string]interface{}{}
			eggsPackage["path"] = lib.Library.Egg
			eggsPackage["messages"] = lib.Messages
			eggsPackage["status"] = lib.Status
			eggs = append(eggs, eggsPackage)
		}
		if len(lib.Library.Whl) > 0 {
			whlPackage := map[string]interface{}{}
			whlPackage["path"] = lib.Library.Whl
			whlPackage["messages"] = lib.Messages
			whlPackage["status"] = lib.Status
			whls = append(whls, whlPackage)
		}
		if lib.Library.Pypi != nil {
			pypiPackage := map[string]interface{}{}
			if len(lib.Library.Pypi.Package) > 0 {
				pypiPackage["package"] = lib.Library.Pypi.Package
				pypiPackage["messages"] = lib.Messages
				pypiPackage["status"] = lib.Status
			}
			if len(lib.Library.Pypi.Repo) > 0 {
				pypiPackage["repo"] = lib.Library.Pypi.Repo
			}
			pypi = append(pypi, pypiPackage)
		}
		if lib.Library.Maven != nil {
			mvnPackage := map[string]interface{}{}
			if len(lib.Library.Maven.Coordinates) > 0 {
				mvnPackage["coordinates"] = lib.Library.Maven.Coordinates
				mvnPackage["messages"] = lib.Messages
				mvnPackage["status"] = lib.Status
			}
			if len(lib.Library.Maven.Repo) > 0 {
				mvnPackage["repo"] = lib.Library.Maven.Repo
			}
			if len(lib.Library.Maven.Exclusions) > 0 {
				mvnPackage["exclusions"] = lib.Library.Maven.Exclusions
			}
			maven = append(maven, mvnPackage)
		}
		if lib.Library.Cran != nil {
			cranPackage := map[string]interface{}{}
			if len(lib.Library.Cran.Package) > 0 {
				cranPackage["package"] = lib.Library.Cran.Package
				cranPackage["messages"] = lib.Messages
				cranPackage["status"] = lib.Status
			}
			if len(lib.Library.Cran.Repo) > 0 {
				cranPackage["repo"] = lib.Library.Cran.Repo
			}
			cran = append(cran, cranPackage)
		}
	}

	err = d.Set("library_jar", jars)
	if err != nil {
		return err
	}

	err = d.Set("library_egg", eggs)
	if err != nil {
		return err
	}

	err = d.Set("library_whl", whls)
	if err != nil {
		return err
	}

	err = d.Set("library_pypi", pypi)
	if err != nil {
		return err
	}
	err = d.Set("library_maven", maven)
	if err != nil {
		return err
	}

	err = d.Set("library_cran", cran)
	if err != nil {
		return err
	}

	if _, ok := d.GetOk("num_workers"); ok {
		err := d.Set("num_workers", clusterInfo.NumWorkers)
		if err != nil {
			return err
		}
	}

	if clusterInfo.AutoScale != nil {
		autoscale := map[string]int{}
		autoscale["min_workers"] = int(clusterInfo.AutoScale.MinWorkers)
		autoscale["max_workers"] = int(clusterInfo.AutoScale.MaxWorkers)
		autoscaleSet := []map[string]int{autoscale}
		err := d.Set("autoscale", autoscaleSet)
		if err != nil {
			return err
		}
	}

	err = d.Set("cluster_name", clusterInfo.ClusterName)
	if err != nil {
		return err
	}

	err = d.Set("spark_version", clusterInfo.SparkVersion)
	if err != nil {
		return err
	}

	if len(clusterInfo.SparkConf) >= 0 {
		err = d.Set("spark_conf", clusterInfo.SparkConf)
		if err != nil {
			return err
		}
	}

	attributes, ok := d.GetOk("aws_attributes")
	log.Println("AWS ATTRIBUTES")
	log.Println(attributes)
	log.Println(d.GetChange("aws_attributes"))
	log.Println(d.HasChange("aws_attributes"))
	if ok && clusterInfo.AwsAttributes != nil {
		awsAtts := map[string]interface{}{}
		awsAtts["availability"] = string(clusterInfo.AwsAttributes.Availability)
		awsAtts["zone_id"] = clusterInfo.AwsAttributes.ZoneID
		awsAtts["spot_bid_price_percent"] = int(clusterInfo.AwsAttributes.SpotBidPricePercent)
		awsAtts["instance_profile_arn"] = clusterInfo.AwsAttributes.InstanceProfileArn
		awsAtts["first_on_demand"] = int(clusterInfo.AwsAttributes.FirstOnDemand)
		awsAtts["ebs_volume_type"] = string(clusterInfo.AwsAttributes.EbsVolumeType)
		awsAtts["ebs_volume_count"] = int(clusterInfo.AwsAttributes.EbsVolumeCount)
		awsAtts["ebs_volume_size"] = int(clusterInfo.AwsAttributes.EbsVolumeSize)
		awsAttsSet := []map[string]interface{}{awsAtts}
		err = d.Set("aws_attributes", awsAttsSet)
		if err != nil {
			return err
		}
	} else {
		err = d.Set("aws_attributes", nil)
	}

	err = d.Set("driver_node_type_id", clusterInfo.DriverNodeTypeID)
	if err != nil {
		return err
	}

	err = d.Set("node_type_id", clusterInfo.NodeTypeID)
	if err != nil {
		return err
	}

	if len(clusterInfo.SSHPublicKeys) > 0 {
		err = d.Set("ssh_public_keys", clusterInfo.SSHPublicKeys)
		if err != nil {
			return err
		}
	}

	if len(clusterInfo.CustomTags) > 0 {
		err = d.Set("custom_tags", clusterInfo.CustomTags)
		if err != nil {
			return err
		}
	}

	if clusterInfo.ClusterLogConf != nil {
		clusterLogConfMap := map[string]interface{}{}
		if clusterInfo.ClusterLogConf.Dbfs != nil {
			clusterLogConfMap["dbfs_destination"] = clusterInfo.ClusterLogConf.Dbfs.Destination
		}
		if clusterInfo.ClusterLogConf.S3 != nil {
			clusterLogConfMap["s3_destination"] = clusterInfo.ClusterLogConf.S3.Destination
			clusterLogConfMap["s3_region"] = clusterInfo.ClusterLogConf.S3.Region
			clusterLogConfMap["s3_endpoint"] = clusterInfo.ClusterLogConf.S3.Endpoint
			clusterLogConfMap["s3_enable_encryption"] = clusterInfo.ClusterLogConf.S3.EnableEncryption
			clusterLogConfMap["s3_encryption_type"] = clusterInfo.ClusterLogConf.S3.EncryptionType
			clusterLogConfMap["s3_kms_key"] = clusterInfo.ClusterLogConf.S3.KmsKey
			clusterLogConfMap["s3_canned_acl"] = clusterInfo.ClusterLogConf.S3.CannedACL
		}
		clusterLogConfSet := []map[string]interface{}{clusterLogConfMap}
		err = d.Set("cluster_log_conf", clusterLogConfSet)
		if err != nil {
			return err
		}
	}

	if len(clusterInfo.InitScripts) > 0 {
		var listOfInitScripts []map[string]string
		for _, v := range clusterInfo.InitScripts {
			initScriptStorageConfig := map[string]string{}
			if v.Dbfs != nil {

				initScriptStorageConfig["dbfs_destination"] = v.Dbfs.Destination
			} else {
				initScriptStorageConfig["s3_destination"] = v.S3.Destination
				initScriptStorageConfig["s3_region"] = v.S3.Region
				initScriptStorageConfig["s3_endpoint"] = v.S3.Endpoint
			}
			listOfInitScripts = append(listOfInitScripts, initScriptStorageConfig)
		}
		err = d.Set("init_scripts", listOfInitScripts)
		if err != nil {
			return err
		}
	}

	if clusterInfo.DockerImage != nil {
		dockerImage := map[string]string{}
		dockerImage["url"] = clusterInfo.DockerImage.Url
		if clusterInfo.DockerImage.BasicAuth != nil {
			dockerImage["username"] = clusterInfo.DockerImage.BasicAuth.Username
			dockerImage["password"] = clusterInfo.DockerImage.BasicAuth.Password
		}
		dockerImageSet := []map[string]string{dockerImage}
		err = d.Set("docker_image", dockerImageSet)
		if err != nil {
			return err
		}
	}

	if len(clusterInfo.SparkEnvVars) > 0 {
		err = d.Set("spark_env_vars", clusterInfo.SparkEnvVars)
		if err != nil {
			return err
		}
	}

	err = d.Set("autotermination_minutes", clusterInfo.AutoterminationMinutes)
	if err != nil {
		return err
	}

	err = d.Set("enable_elastic_disk", clusterInfo.EnableElasticDisk)
	if err != nil {
		return err
	}

	err = d.Set("enable_elastic_disk", clusterInfo.EnableElasticDisk)
	if err != nil {
		return err
	}

	if _, ok := d.GetOk("instance_pool_id"); ok {
		err := d.Set("instance_pool_id", clusterInfo.InstancePoolId)
		if err != nil {
			return err
		}
	}

	if len(clusterInfo.DefaultTags) > 0 {
		err = d.Set("default_tags", clusterInfo.DefaultTags)
		if err != nil {
			return err
		}
	}
	err = d.Set("state", string(clusterInfo.State))
	if err != nil {
		return err
	}

	err = d.Set("state_message", string(clusterInfo.StateMessage))

	return nil
}

func calculateLibraryChanges(new []model.Library, old []model.Library) ([]model.Library, []model.Library) {
	newDictionary := map[string]model.Library{}
	newKeys := []string{}

	for _, library := range new {
		if len(library.Whl) > 0 {
			newDictionary[library.Whl] = library
			newKeys = append(newKeys, library.Whl)
		} else if len(library.Egg) > 0 {
			newDictionary[library.Egg] = library
			newKeys = append(newKeys, library.Egg)
		} else if len(library.Jar) > 0 {
			newDictionary[library.Jar] = library
			newKeys = append(newKeys, library.Jar)
		} else if len(library.Pypi.Package) > 0 {
			newDictionary[library.Pypi.Package+library.Pypi.Repo] = library
			newKeys = append(newKeys, library.Pypi.Package+library.Pypi.Repo)
		} else if len(library.Maven.Coordinates) > 0 {
			newDictionary[library.Maven.Coordinates+library.Maven.Repo+strings.Join(library.Maven.Exclusions, "")] = library
			newKeys = append(newKeys, library.Maven.Coordinates+library.Maven.Repo+strings.Join(library.Maven.Exclusions, ""))
		} else if len(library.Cran.Package) > 0 {
			newDictionary[library.Cran.Package+library.Cran.Repo] = library
			newKeys = append(newKeys, library.Cran.Package+library.Cran.Repo)
		}
	}

	oldDictionary := map[string]model.Library{}
	oldKeys := []string{}
	for _, library := range old {
		if len(library.Whl) > 0 {
			oldDictionary[library.Whl] = library
			oldKeys = append(oldKeys, library.Whl)
		} else if len(library.Egg) > 0 {
			oldDictionary[library.Egg] = library
			oldKeys = append(oldKeys, library.Egg)
		} else if len(library.Jar) > 0 {
			oldDictionary[library.Jar] = library
			oldKeys = append(oldKeys, library.Jar)
		} else if len(library.Pypi.Package) > 0 {
			oldDictionary[library.Pypi.Package+library.Pypi.Repo] = library
			oldKeys = append(oldKeys, library.Pypi.Package+library.Pypi.Repo)
		} else if len(library.Maven.Coordinates) > 0 {
			oldDictionary[library.Maven.Coordinates+library.Maven.Repo+strings.Join(library.Maven.Exclusions, "")] = library
			oldKeys = append(oldKeys, library.Maven.Coordinates+library.Maven.Repo+strings.Join(library.Maven.Exclusions, ""))
		} else if len(library.Cran.Package) > 0 {
			oldDictionary[library.Cran.Package+library.Cran.Repo] = library
			oldKeys = append(oldKeys, library.Cran.Package+library.Cran.Repo)
		}
	}

	installLibrariesKeys := diff(newKeys, oldKeys)
	uninstallLibrariesKeys := diff(oldKeys, newKeys)

	installLibraries := []model.Library{}
	for _, key := range installLibrariesKeys {
		installLibraries = append(installLibraries, newDictionary[key])
	}

	uninstallLibraries := []model.Library{}
	for _, key := range uninstallLibrariesKeys {
		uninstallLibraries = append(uninstallLibraries, oldDictionary[key])
	}
	return installLibraries, uninstallLibraries
}

func parseLibraryStatusListToLibraries(libraryStatuses []model.LibraryStatus) []model.Library {
	var libraries []model.Library
	for _, libraryStatus := range libraryStatuses {
		libraries = append(libraries, *libraryStatus.Library)
	}
	return libraries
}

func parseSchemaToClusterLibraries(d *schema.ResourceData) []model.Library {
	var libraryList []model.Library
	if jars, ok := d.GetOk("library_jar"); ok {
		libraries := jars.(*schema.Set).List()
		for _, library := range libraries {
			libraryMap := library.(map[string]interface{})
			thisLibrary := model.Library{
				Jar: libraryMap["path"].(string),
			}
			libraryList = append(libraryList, thisLibrary)
		}
	}
	if eggs, ok := d.GetOk("library_egg"); ok {
		libraries := eggs.(*schema.Set).List()
		for _, library := range libraries {
			libraryMap := library.(map[string]interface{})
			thisLibrary := model.Library{
				Egg: libraryMap["path"].(string),
			}
			libraryList = append(libraryList, thisLibrary)
		}
	}
	if whls, ok := d.GetOk("library_whl"); ok {
		libraries := whls.(*schema.Set).List()
		for _, library := range libraries {
			libraryMap := library.(map[string]interface{})
			thisLibrary := model.Library{
				Whl: libraryMap["path"].(string),
			}
			libraryList = append(libraryList, thisLibrary)
		}
	}
	if pypis, ok := d.GetOk("library_pypi"); ok {
		libraries := pypis.(*schema.Set).List()
		for _, library := range libraries {
			libraryMap := library.(map[string]interface{})
			var pypi model.PyPi
			if pkg, ok := libraryMap["package"]; ok {
				pypi.Package = pkg.(string)
			}
			if repo, ok := libraryMap["repo"]; ok {
				pypi.Repo = repo.(string)
			}
			thisLibrary := model.Library{
				Pypi: &pypi,
			}
			libraryList = append(libraryList, thisLibrary)
		}
	}
	if mavens, ok := d.GetOk("library_maven"); ok {
		libraries := mavens.(*schema.Set).List()
		for _, library := range libraries {
			libraryMap := library.(map[string]interface{})
			var maven model.Maven
			if coordinates, ok := libraryMap["coordinates"]; ok {
				maven.Coordinates = coordinates.(string)
			}
			if repo, ok := libraryMap["repo"]; ok {
				maven.Repo = repo.(string)
			}
			if exclusions, ok := libraryMap["exclusions"]; ok {
				maven.Exclusions = convertListInterfaceToString(exclusions.([]interface{}))
			}
			thisLibrary := model.Library{
				Maven: &maven,
			}
			libraryList = append(libraryList, thisLibrary)
		}
	}
	if crans, ok := d.GetOk("library_cran"); ok {
		libraries := crans.(*schema.Set).List()
		for _, library := range libraries {
			libraryMap := library.(map[string]interface{})
			var cran model.Cran
			if pkg, ok := libraryMap["package"]; ok {
				cran.Package = pkg.(string)
			}
			if repo, ok := libraryMap["repo"]; ok {
				cran.Repo = repo.(string)
			}
			thisLibrary := model.Library{
				Cran: &cran,
			}
			libraryList = append(libraryList, thisLibrary)
		}
	}
	return libraryList
}

func resourceClusterUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(service.DBApiClient)
	id := d.Id()
	clusterInfo, err := client.Clusters().Get(id)
	if err != nil {
		return err
	}
	libraryStatuses, err := client.Libraries().List(id)
	if err != nil {
		return err
	}
	newLibraries := parseSchemaToLibraries(d)
	oldLibraries := parseLibraryStatusListToLibraries(libraryStatuses)

	installs, uninstalls := calculateLibraryChanges(newLibraries, oldLibraries)

	clusterState := clusterInfo.State

	if model.ContainsClusterState([]model.ClusterState{model.ClusterState(model.ClusterStateTerminated)}, clusterState) {
		cluster := parseSchemaToCluster(d, "")
		cluster.ClusterId = id
		err := client.Clusters().Edit(cluster)
		if err != nil {
			return err
		}
		if len(uninstalls) > 0 {
			err = client.Libraries().Delete(id, uninstalls)
			if err != nil {
				return err
			}
		}
		if len(installs) > 0 {
			err := client.Clusters().Start(id)
			if err != nil {
				return err
			}
			err = client.Clusters().WaitForClusterRunning(clusterInfo.ClusterID, 30, 120)
			if err != nil {
				return err
			}
			err = client.Libraries().Create(id, installs)
			if err != nil {
				return err
			}
			time.Sleep(10 * time.Second)
			err = client.Clusters().Terminate(id)
			if err != nil {
				return err
			}
			err = client.Clusters().WaitForClusterTerminated(clusterInfo.ClusterID, 30, 120)
			if err != nil {
				return err
			}
		}
		return resourceClusterRead(d, m)
	} else if model.ContainsClusterState([]model.ClusterState{model.ClusterState(model.ClusterStateRunning)}, clusterState) {
		cluster := parseSchemaToCluster(d, "")
		cluster.ClusterId = id

		if len(installs) > 0 {
			err = client.Libraries().Create(id, installs)
			if err != nil {
				return err
			}
			time.Sleep(5 * time.Second)
		}

		if len(uninstalls) > 0 {
			err = client.Libraries().Delete(id, uninstalls)
			if err != nil {
				return err
			}
			time.Sleep(1 * time.Second)
		}
		err := client.Clusters().Edit(cluster)
		if err != nil {
			return err
		}

		err = client.Clusters().WaitForClusterRunning(clusterInfo.ClusterID, 30, 120)
		if err != nil {
			return err
		}
		return resourceClusterRead(d, m)
	} else if model.ContainsClusterState([]model.ClusterState{model.ClusterStatePending,
		model.ClusterStateResizing}, clusterState) {
		err = client.Clusters().WaitForClusterRunning(clusterInfo.ClusterID, 30, 120)
		if err != nil {
			return err
		}
		cluster := parseSchemaToCluster(d, "")
		cluster.ClusterId = id

		if len(installs) > 0 {
			err = client.Libraries().Create(id, installs)
			if err != nil {
				return err
			}
		}
		if len(uninstalls) > 0 {
			err = client.Libraries().Delete(id, uninstalls)
			if err != nil {
				return err
			}
			time.Sleep(5 * time.Second)
		}
		err := client.Clusters().Edit(cluster)
		if err != nil {
			return err
		}

		err = client.Clusters().WaitForClusterRunning(clusterInfo.ClusterID, 30, 120)
		if err != nil {
			return err
		}
		return resourceClusterRead(d, m)
	}

	return errors.New("Unable to edit cluster due to cluster state not being in a runnable/terminated state.")
}

func resourceClusterDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(service.DBApiClient)

	err := client.Clusters().PermanentDelete(id)
	if err != nil {
		return err
	}
	return err
}

func parseSchemaToCluster(d *schema.ResourceData, schemaAttPrefix string) model.Cluster {
	cluster := model.Cluster{}

	//Deal with Num workers
	if numWorkers, ok := d.GetOk(schemaAttPrefix + "num_workers"); ok {
		cluster.NumWorkers = int32(numWorkers.(int))
	}

	//Deal with auto scaling options
	autoScale := model.AutoScale{}
	if autoscale, ok := d.GetOk(schemaAttPrefix + "autoscale"); ok {
		autoScaleOptions := getMapFromOneItemSet(autoscale)
		if minWorkers, ok := autoScaleOptions["min_workers"]; ok {
			//minVal, _ := strconv.ParseInt(minWorkers.(string), 10, 32)
			autoScale.MinWorkers = int32(minWorkers.(int))
		}
		if maxWorkers, ok := autoScaleOptions["max_workers"]; ok {
			//maxVal, _ := strconv.ParseInt(maxWorkers.(string), 10, 32)
			autoScale.MaxWorkers = int32(maxWorkers.(int))
		}
		cluster.Autoscale = &autoScale
	}

	//Deal with cluster name
	if clusterName, ok := d.GetOk(schemaAttPrefix + "cluster_name"); ok {
		cluster.ClusterName = clusterName.(string)
	}

	//Deal with spark versions
	if sparkVersion, ok := d.GetOk(schemaAttPrefix + "spark_version"); ok {
		cluster.SparkVersion = sparkVersion.(string)
	}

	//Deal with spark confs
	if sparkConf, ok := d.GetOk(schemaAttPrefix + "spark_conf"); ok {
		cluster.SparkConf = convertMapStringInterfaceToStringString(sparkConf.(map[string]interface{}))
	}

	//Deal with aws attributes for aws deployment
	awsAttributes := model.AwsAttributes{}

	if awsAttributesSchema, ok := d.GetOk(schemaAttPrefix + "aws_attributes"); ok {
		awsAttributesMap := getMapFromOneItemList(awsAttributesSchema)
		if availability, ok := awsAttributesMap["availability"]; ok {
			awsAttributes.Availability = model.AwsAvailability(availability.(string))
		}
		if zoneId, ok := awsAttributesMap["zone_id"]; ok {
			awsAttributes.ZoneID = zoneId.(string)
		}
		if spotBidPricePercent, ok := awsAttributesMap["spot_bid_price_percent"]; ok {
			//val, _ := strconv.ParseInt(spotBidPricePercent.(string), 10, 32)
			awsAttributes.SpotBidPricePercent = int32(spotBidPricePercent.(int))
		}
		if instanceProfileArn, ok := awsAttributesMap["instance_profile_arn"]; ok {
			awsAttributes.InstanceProfileArn = instanceProfileArn.(string)
		}
		if firstOnDemand, ok := awsAttributesMap["first_on_demand"]; ok {
			//val, _ := strconv.ParseInt(firstOnDemand.(string), 10, 32)
			awsAttributes.FirstOnDemand = int32(firstOnDemand.(int))
		}
		if ebsVolumeType, ok := awsAttributesMap["ebs_volume_type"]; ok {
			awsAttributes.EbsVolumeType = model.EbsVolumeType(ebsVolumeType.(string))
		}
		if ebsVolumeCount, ok := awsAttributesMap["ebs_volume_count"]; ok {
			//val, _ := strconv.ParseInt(ebsVolumeCount.(string), 10, 32)
			awsAttributes.FirstOnDemand = int32(ebsVolumeCount.(int))
		}
		if ebsVolumeSize, ok := awsAttributesMap["ebs_volume_size"]; ok {
			//val, _ := strconv.ParseInt(ebsVolumeSize.(string), 10, 32)
			awsAttributes.EbsVolumeSize = int32(ebsVolumeSize.(int))
		}
		cluster.AwsAttributes = &awsAttributes
	}

	//Deal with driver node type id
	if driverNodeTypeId, ok := d.GetOk(schemaAttPrefix + "driver_node_type_id"); ok {
		cluster.DriverNodeTypeID = driverNodeTypeId.(string)
	}

	//Deal with worker node type id
	if nodeTypeId, ok := d.GetOk(schemaAttPrefix + "node_type_id"); ok {
		cluster.NodeTypeID = nodeTypeId.(string)
	}

	//Deal with worker ssh public keys
	if sshPublicKeys, ok := d.GetOk(schemaAttPrefix + "ssh_public_keys"); ok {
		cluster.SSHPublicKeys = convertListInterfaceToString(sshPublicKeys.(*schema.Set).List())
	}

	//Deal with worker custom tags
	if customTags, ok := d.GetOk(schemaAttPrefix + "custom_tags"); ok {
		tags := customTags.(map[string]interface{})
		cluster.CustomTags = convertMapStringInterfaceToStringString(tags)
	}

	//Deal with worker cluster log config
	clusterLogConf := model.StorageInfo{}
	if clusterLogConfSet, ok := d.GetOk(schemaAttPrefix + "cluster_log_conf"); ok {
		//clusterLogConfMap := dbfsClusterLogConfSet.(*schema.Set).List()[0].(map[string]interface{})
		clusterLogConfMap := getMapFromOneItemSet(clusterLogConfSet)
		dbfsDestination, dbfsOk := clusterLogConfMap["dbfs_destination"]
		if dbfsOk && len(dbfsDestination.(string)) > 0 {
			dbfsStorage := model.DbfsStorageInfo{}
			dbfsStorage.Destination = dbfsDestination.(string)
			clusterLogConf.Dbfs = &dbfsStorage
		}
		//cluster.ClusterLogConf = &clusterLogConf
		//}
		//if s3clusterLogConfSet, ok := d.GetOk("s3_cluster_log_conf"); ok {
		//clusterLogConfMap := s3clusterLogConfSet.(map[string]interface{})
		//s3ClusterLogConf := getMapFromOneItemSet(s3clusterLogConfSet)
		if (!dbfsOk || len(dbfsDestination.(string)) == 0) && len(clusterLogConfMap["s3_destination"].(string)) > 0 {
			s3Storage := model.S3StorageInfo{}
			if s3Destination, ok := clusterLogConfMap["s3_destination"]; ok {
				s3Storage.Destination = s3Destination.(string)
			}
			if s3Region, ok := clusterLogConfMap["s3_region"]; ok {
				s3Storage.Region = s3Region.(string)
			}
			if s3Endpoint, ok := clusterLogConfMap["s3_endpoint"]; ok {
				s3Storage.Endpoint = s3Endpoint.(string)
			}
			if s3EnableEncryption, ok := clusterLogConfMap["s3_enable_encryption"]; ok {
				//b, _ := strconv.ParseBool(s3EnableEncryption.(string))
				s3Storage.EnableEncryption = s3EnableEncryption.(bool)
			}
			if s3EncrptionType, ok := clusterLogConfMap["s3_encryption_type"]; ok {
				s3Storage.EncryptionType = s3EncrptionType.(string)
			}
			if s3KMSKey, ok := clusterLogConfMap["s3_kms_key"]; ok {
				s3Storage.KmsKey = s3KMSKey.(string)
			}
			if s3CannedACL, ok := clusterLogConfMap["s3_canned_acl"]; ok {
				s3Storage.CannedACL = s3CannedACL.(string)
			}
			clusterLogConf.S3 = &s3Storage
		}
		if clusterLogConf.S3 != nil || clusterLogConf.Dbfs != nil {
			cluster.ClusterLogConf = &clusterLogConf
		}
	}

	//Deal with worker init script setup
	if initScripts, ok := d.GetOk(schemaAttPrefix + "init_scripts"); ok {
		initScripts := initScripts.(*schema.Set).List()
		initScriptsLocations := []model.StorageInfo{}
		for _, v := range initScripts {
			initScript := v.(map[string]interface{})
			storageInfo := model.StorageInfo{}
			if dbfsDestination, ok := initScript["dbfs_destination"]; ok && len(dbfsDestination.(string)) > 0 {
				dbfsStorage := model.DbfsStorageInfo{}
				dbfsStorage.Destination = dbfsDestination.(string)
				storageInfo.Dbfs = &dbfsStorage
				initScriptsLocations = append(initScriptsLocations, storageInfo)
			} else {
				s3Storage := model.S3StorageInfo{}
				if s3Destination, ok := initScript["s3_destination"]; ok {
					s3Storage.Destination = s3Destination.(string)
				}
				if s3Region, ok := initScript["s3_region"]; ok {
					s3Storage.Region = s3Region.(string)
				}
				if s3Endpoint, ok := initScript["s3_endpoint"]; ok {
					s3Storage.Endpoint = s3Endpoint.(string)
				}
				storageInfo.S3 = &s3Storage
				initScriptsLocations = append(initScriptsLocations, storageInfo)
			}
		}
		cluster.InitScripts = initScriptsLocations

	}

	//Deal with docker image for DCS
	dockerImageData := model.DockerImage{}
	if dockerImageSet, ok := d.GetOk(schemaAttPrefix + "docker_image"); ok {
		dockerImageConf := getMapFromOneItemSet(dockerImageSet)
		if url, ok := dockerImageConf["url"]; ok {
			dockerImageData.Url = url.(string)
		}
		dockerAuthData := model.DockerBasicAuth{}
		username, userOk := dockerImageConf["username"]
		password, passOk := dockerImageConf["password"]
		if userOk && passOk {
			dockerAuthData.Username = username.(string)
			dockerAuthData.Password = password.(string)
			dockerImageData.BasicAuth = &dockerAuthData
		}
		cluster.DockerImage = &dockerImageData
	}

	//Deal with spark environment variables
	if sparkEnv, ok := d.GetOk(schemaAttPrefix + "spark_env_vars"); ok {
		cluster.SparkEnvVars = convertMapStringInterfaceToStringString(sparkEnv.(map[string]interface{}))
	}

	//Deal with auto termination minutes
	if autoTerminationMinutes, ok := d.GetOk(schemaAttPrefix + "autotermination_minutes"); ok {
		cluster.AutoterminationMinutes = int32(autoTerminationMinutes.(int))
	}

	//Deal with enable elastic disk
	if enableElasticDisk, ok := d.GetOk(schemaAttPrefix + "enable_elastic_disk"); ok {
		cluster.EnableElasticDisk = enableElasticDisk.(bool)
	}

	//Deal with instance pool id
	if instancePoolID, ok := d.GetOk(schemaAttPrefix + "instance_pool_id"); ok {
		cluster.InstancePoolId = instancePoolID.(string)
	}

	//Deal with idempotency token
	if idempotencyToken, ok := d.GetOk(schemaAttPrefix + "idempotency_token"); ok {
		cluster.IdempotencyToken = idempotencyToken.(string)
	}
	return cluster
}

func clusterInitScriptHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if v, ok := m["dbfs_destination"]; ok {
		buf.WriteString(v.(string))
	}
	if v, ok := m["s3_region"]; ok {
		buf.WriteString(v.(string))
	}
	if v, ok := m["s3_endpoint"]; ok {
		buf.WriteString(v.(string))
	}
	return hashcode.String(buf.String())
}

func createDynamicHash(k []string) func(v interface{}) int {
	return func(v interface{}) int {
		var buf bytes.Buffer
		m := v.(map[string]interface{})
		sort.Strings(k)
		for _, k := range k {
			value, ok := m[k].(string)
			if ok {
				if m[k] != nil && len(value) > 0 {
					buf.WriteString(value)
				}
			} else {
				list, ok := m[k].([]interface{})
				if ok {
					for _, val := range list {
						if val != nil && len(val.(string)) > 0 {
							buf.WriteString(val.(string))
						}
					}
				}
			}

		}
		return hashcode.String(buf.String())
	}
}

func getMapFromOneItemList(input interface{}) map[string]interface{} {
	inputList := input.([]interface{})
	if len(inputList) >= 1 {
		return inputList[0].(map[string]interface{})
	}
	return nil
}

func getMapFromOneItemSet(input interface{}) map[string]interface{} {
	inputList := input.(*schema.Set).List()
	if len(inputList) >= 1 {
		return inputList[0].(map[string]interface{})
	}
	return nil
}

func isClusterMissing(errorMsg, resourceId string) bool {
	return strings.Contains(errorMsg, "INVALID_PARAMETER_VALUE") &&
		strings.Contains(errorMsg, fmt.Sprintf("Cluster %s does not exist", resourceId))
}
