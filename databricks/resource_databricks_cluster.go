package databricks

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceClusterCreate,
		Read:   resourceClusterRead,
		Update: resourceClusterUpdate,
		Delete: resourceClusterDelete,

		Schema: map[string]*schema.Schema{
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
				Elem: &schema.Resource{
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
										AtLeastOneOf: []string{
											"cluster_log_conf.0.s3.0.region",
											"cluster_log_conf.0.s3.0.endpoint",
										},
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
							ExactlyOneOf: []string{
								"cluster_log_conf.0.dbfs",
								"cluster_log_conf.0.s3",
							},
						},
					},
				},
			},
			"init_scripts": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 10,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dbfs": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									//TODO: Validate that destination has dbfs:// prefix
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
						},
					},
				},
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
			"library_jar": {
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
			"library_egg": {
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
			"library_whl": {
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
			"library_pypi": {
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
			"library_maven": {
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
			"library_cran": {
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
		},
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
	client := m.(*service.DBApiClient)

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
		if err != nil {
			return err
		}
	}

	return resourceClusterRead(d, m)
}

func resourceClusterRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)
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
	if err != nil {
		return err
	}

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

	if clusterInfo.SparkConf != nil {
		err = d.Set("spark_conf", clusterInfo.SparkConf)
		if err != nil {
			return err
		}
	}

	_, ok := d.GetOk("aws_attributes")
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
		if err != nil {
			return err
		}
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
		clusterLogConfList := []interface{}{}
		clusterLogConfListItem := map[string]interface{}{}
		if clusterInfo.ClusterLogConf.Dbfs != nil {
			dbfsList := []interface{}{}
			dbfsListItem := map[string]interface{}{}
			dbfsListItem["destination"] = clusterInfo.ClusterLogConf.Dbfs.Destination
			dbfsList = append(dbfsList, dbfsListItem)
			clusterLogConfListItem["dbfs"] = dbfsList
		}
		if clusterInfo.ClusterLogConf.S3 != nil {
			s3List := []interface{}{}
			s3ListItem := map[string]interface{}{}
			s3ListItem["destination"] = clusterInfo.ClusterLogConf.S3.Destination
			s3ListItem["region"] = clusterInfo.ClusterLogConf.S3.Region
			s3ListItem["endpoint"] = clusterInfo.ClusterLogConf.S3.Endpoint
			s3ListItem["enable_encryption"] = clusterInfo.ClusterLogConf.S3.EnableEncryption
			s3ListItem["encryption_type"] = clusterInfo.ClusterLogConf.S3.EncryptionType
			s3ListItem["kms_key"] = clusterInfo.ClusterLogConf.S3.KmsKey
			s3ListItem["canned_acl"] = clusterInfo.ClusterLogConf.S3.CannedACL
			s3List = append(s3List, s3ListItem)
			clusterLogConfListItem["s3"] = s3List
		}
		clusterLogConfList = append(clusterLogConfList, clusterLogConfListItem)
		err = d.Set("cluster_log_conf", clusterLogConfList)
		if err != nil {
			return err
		}
	}

	// Handle reading init scripts
	if clusterInfo.InitScripts != nil && len(clusterInfo.InitScripts) > 0 {
		listOfInitScripts := []interface{}{}
		for _, v := range clusterInfo.InitScripts {
			initScriptListItem := map[string]interface{}{}
			if v.Dbfs != nil {
				dbfsList := []interface{}{}
				dbfsListItem := map[string]interface{}{}
				dbfsListItem["destination"] = v.Dbfs.Destination
				dbfsList = append(dbfsList, dbfsListItem)
				initScriptListItem["dbfs"] = dbfsList
			}
			if v.S3 != nil {
				s3List := []interface{}{}
				s3ListItem := map[string]interface{}{}
				s3ListItem["destination"] = v.S3.Destination
				s3ListItem["region"] = v.S3.Region
				s3ListItem["endpoint"] = v.S3.Endpoint
				s3ListItem["enable_encryption"] = v.S3.EnableEncryption
				s3ListItem["encryption_type"] = v.S3.EncryptionType
				s3ListItem["kms_key"] = v.S3.KmsKey
				s3ListItem["canned_acl"] = v.S3.CannedACL
				s3List = append(s3List, s3ListItem)
				initScriptListItem["s3"] = s3List
			}
			listOfInitScripts = append(listOfInitScripts, initScriptListItem)
		}
		err = d.Set("init_scripts", listOfInitScripts)
		if err != nil {
			return err
		}
	}

	if clusterInfo.DockerImage != nil {
		dockerImageList := []interface{}{}
		dockerImageListItem := map[string]interface{}{}
		dockerImageListItem["url"] = clusterInfo.DockerImage.URL
		if clusterInfo.DockerImage.BasicAuth != nil {
			basicAuthList := []interface{}{}
			basicAuthListItem := map[string]interface{}{}
			basicAuthListItem["username"] = clusterInfo.DockerImage.BasicAuth.Username
			basicAuthListItem["password"] = clusterInfo.DockerImage.BasicAuth.Password
			basicAuthList = append(basicAuthList, basicAuthListItem)
			dockerImageListItem["basic_auth"] = basicAuthList
		}

		dockerImageList = append(dockerImageList, dockerImageListItem)
		err = d.Set("docker_image", dockerImageList)
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
		err := d.Set("instance_pool_id", clusterInfo.InstancePoolID)
		if err != nil {
			return err
		}
	}

	if _, ok := d.GetOk("single_user_name"); ok {
		err := d.Set("single_user_name", clusterInfo.SingleUserName)
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

	err = d.Set("cluster_id", clusterInfo.ClusterID)
	if err != nil {
		return err
	}

	err = d.Set("state", string(clusterInfo.State))
	if err != nil {
		return err
	}

	err = d.Set("state_message", clusterInfo.StateMessage)

	return err
}

func calculateLibraryChanges(new []model.Library, old []model.Library) ([]model.Library, []model.Library) {
	newDictionary := map[string]model.Library{}
	newKeys := []string{}

	for _, library := range new {
		switch {
		case len(library.Whl) > 0:
			newDictionary[library.Whl] = library
			newKeys = append(newKeys, library.Whl)
		case len(library.Egg) > 0:
			newDictionary[library.Egg] = library
			newKeys = append(newKeys, library.Egg)
		case len(library.Jar) > 0:
			newDictionary[library.Jar] = library
			newKeys = append(newKeys, library.Jar)
		case library.Pypi != nil && len(library.Pypi.Package) > 0:
			newDictionary[library.Pypi.Package+library.Pypi.Repo] = library
			newKeys = append(newKeys, library.Pypi.Package+library.Pypi.Repo)
		case library.Maven != nil && len(library.Maven.Coordinates) > 0:
			newDictionary[library.Maven.Coordinates+library.Maven.Repo+strings.Join(library.Maven.Exclusions, "")] = library
			newKeys = append(newKeys, library.Maven.Coordinates+library.Maven.Repo+strings.Join(library.Maven.Exclusions, ""))
		case library.Cran != nil && len(library.Cran.Package) > 0:
			newDictionary[library.Cran.Package+library.Cran.Repo] = library
			newKeys = append(newKeys, library.Cran.Package+library.Cran.Repo)
		}
	}

	oldDictionary := map[string]model.Library{}
	oldKeys := []string{}
	for _, library := range old {
		switch {
		case len(library.Whl) > 0:
			oldDictionary[library.Whl] = library
			oldKeys = append(oldKeys, library.Whl)
		case len(library.Egg) > 0:
			oldDictionary[library.Egg] = library
			oldKeys = append(oldKeys, library.Egg)
		case len(library.Jar) > 0:
			oldDictionary[library.Jar] = library
			oldKeys = append(oldKeys, library.Jar)
		case library.Pypi != nil && len(library.Pypi.Package) > 0:
			oldDictionary[library.Pypi.Package+library.Pypi.Repo] = library
			oldKeys = append(oldKeys, library.Pypi.Package+library.Pypi.Repo)
		case library.Maven != nil && len(library.Maven.Coordinates) > 0:
			oldDictionary[library.Maven.Coordinates+library.Maven.Repo+strings.Join(library.Maven.Exclusions, "")] = library
			oldKeys = append(oldKeys, library.Maven.Coordinates+library.Maven.Repo+strings.Join(library.Maven.Exclusions, ""))
		case library.Cran != nil && len(library.Cran.Package) > 0:
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
	client := m.(*service.DBApiClient)
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

	switch {
	case model.ContainsClusterState([]model.ClusterState{model.ClusterState(model.ClusterStateTerminated)}, clusterState):
		cluster := parseSchemaToCluster(d, "")
		cluster.ClusterID = id
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
	case model.ContainsClusterState([]model.ClusterState{model.ClusterState(model.ClusterStateRunning)}, clusterState):
		cluster := parseSchemaToCluster(d, "")
		cluster.ClusterID = id

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
	case model.ContainsClusterState([]model.ClusterState{model.ClusterStatePending, model.ClusterStateResizing}, clusterState):
		err = client.Clusters().WaitForClusterRunning(clusterInfo.ClusterID, 30, 120)
		if err != nil {
			return err
		}
		cluster := parseSchemaToCluster(d, "")
		cluster.ClusterID = id

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

	return errors.New("unable to edit cluster due to cluster state not being in a runnable/terminated state")
}

func resourceClusterDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DBApiClient)

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
		if zoneID, ok := awsAttributesMap["zone_id"]; ok {
			awsAttributes.ZoneID = zoneID.(string)
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
			awsAttributes.EbsVolumeCount = int32(ebsVolumeCount.(int))
		}
		if ebsVolumeSize, ok := awsAttributesMap["ebs_volume_size"]; ok {
			//val, _ := strconv.ParseInt(ebsVolumeSize.(string), 10, 32)
			awsAttributes.EbsVolumeSize = int32(ebsVolumeSize.(int))
		}
		cluster.AwsAttributes = &awsAttributes
	}

	//Deal with driver node type id
	if driverNodeTypeID, ok := d.GetOk(schemaAttPrefix + "driver_node_type_id"); ok {
		cluster.DriverNodeTypeID = driverNodeTypeID.(string)
	}

	//Deal with worker node type id
	if nodeTypeID, ok := d.GetOk(schemaAttPrefix + "node_type_id"); ok {
		cluster.NodeTypeID = nodeTypeID.(string)
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
		// Fetch the cluster log config list; this should only be one item
		clusterLogConfItemList := clusterLogConfSet.([]interface{})
		if len(clusterLogConfItemList) > 0 {
			// Fetch the first item and convert it to map to identify dbfs vs s3
			clusterLogConfItemMap := clusterLogConfItemList[0].(map[string]interface{})

			// DBFS PATH
			if dbfsSpec, ok := clusterLogConfItemMap["dbfs"]; ok {
				// Fetch the first item for dbfs config
				dbfsStorageItemList := dbfsSpec.([]interface{})
				if len(dbfsStorageItemList) > 0 {
					dbfsStorage := model.DbfsStorageInfo{}
					dbfsStorageMap := dbfsStorageItemList[0].(map[string]interface{})
					dbfsStorage.Destination = dbfsStorageMap["destination"].(string)
					clusterLogConf.Dbfs = &dbfsStorage
				}
			}
			// S3 PATH
			if s3Spec, ok := clusterLogConfItemMap["s3"]; ok {
				// Fetch the first item for s3 config
				s3StorageItemList := s3Spec.([]interface{})
				if len(s3StorageItemList) > 0 {
					s3StorageMap := s3StorageItemList[0].(map[string]interface{})
					s3Storage := model.S3StorageInfo{}
					if s3Destination, ok := s3StorageMap["destination"]; ok {
						s3Storage.Destination = s3Destination.(string)
					}
					if s3Region, ok := s3StorageMap["region"]; ok {
						s3Storage.Region = s3Region.(string)
					}
					if s3Endpoint, ok := s3StorageMap["endpoint"]; ok {
						s3Storage.Endpoint = s3Endpoint.(string)
					}
					if s3EnableEncryption, ok := s3StorageMap["enable_encryption"]; ok {
						s3Storage.EnableEncryption = s3EnableEncryption.(bool)
					}
					if s3EncrptionType, ok := s3StorageMap["encryption_type"]; ok {
						s3Storage.EncryptionType = s3EncrptionType.(string)
					}
					if s3KMSKey, ok := s3StorageMap["kms_key"]; ok {
						s3Storage.KmsKey = s3KMSKey.(string)
					}
					if s3CannedACL, ok := s3StorageMap["canned_acl"]; ok {
						s3Storage.CannedACL = s3CannedACL.(string)
					}
					clusterLogConf.S3 = &s3Storage
				}
			}
		}
		if clusterLogConf.S3 != nil || clusterLogConf.Dbfs != nil {
			cluster.ClusterLogConf = &clusterLogConf
		}
	}

	//Deal with worker init script setup
	if initScripts, ok := d.GetOk(schemaAttPrefix + "init_scripts"); ok {
		initScripts := initScripts.([]interface{})
		initScriptsLocations := []model.StorageInfo{}
		for _, v := range initScripts {
			initScript := v.(map[string]interface{})
			initScriptsConf := model.StorageInfo{}
			if dbfsSpec, ok := initScript["dbfs"]; ok {
				// Fetch the first item for dbfs config
				dbfsStorageItemList := dbfsSpec.([]interface{})
				if len(dbfsStorageItemList) > 0 {
					dbfsStorage := model.DbfsStorageInfo{}
					dbfsStorageMap := dbfsStorageItemList[0].(map[string]interface{})
					dbfsStorage.Destination = dbfsStorageMap["destination"].(string)
					initScriptsConf.Dbfs = &dbfsStorage
				}
			}
			if s3Spec, ok := initScript["s3"]; ok {
				// Fetch the first item for s3 config
				s3StorageItemList := s3Spec.([]interface{})
				if len(s3StorageItemList) > 0 {
					s3StorageMap := s3StorageItemList[0].(map[string]interface{})
					s3Storage := model.S3StorageInfo{}
					if s3Destination, ok := s3StorageMap["destination"]; ok {
						s3Storage.Destination = s3Destination.(string)
					}
					if s3Region, ok := s3StorageMap["region"]; ok {
						s3Storage.Region = s3Region.(string)
					}
					if s3Endpoint, ok := s3StorageMap["endpoint"]; ok {
						s3Storage.Endpoint = s3Endpoint.(string)
					}
					if s3EnableEncryption, ok := s3StorageMap["enable_encryption"]; ok {
						s3Storage.EnableEncryption = s3EnableEncryption.(bool)
					}
					if s3EncrptionType, ok := s3StorageMap["encryption_type"]; ok {
						s3Storage.EncryptionType = s3EncrptionType.(string)
					}
					if s3KMSKey, ok := s3StorageMap["kms_key"]; ok {
						s3Storage.KmsKey = s3KMSKey.(string)
					}
					if s3CannedACL, ok := s3StorageMap["canned_acl"]; ok {
						s3Storage.CannedACL = s3CannedACL.(string)
					}
					initScriptsConf.S3 = &s3Storage
				}
			}
			initScriptsLocations = append(initScriptsLocations, initScriptsConf)
		}
		cluster.InitScripts = initScriptsLocations
	}

	//Deal with docker image for DCS
	dockerImageData := model.DockerImage{}
	if dockerImageList, ok := d.GetOk(schemaAttPrefix + "docker_image"); ok {
		//dockerImageConf := getMapFromOneItemSet(dockerImageList)
		dockerImageListInterface := dockerImageList.([]interface{})
		if len(dockerImageListInterface) > 0 {
			dockerImageListItem := dockerImageListInterface[0].(map[string]interface{})
			if url, ok := dockerImageListItem["url"]; ok {
				dockerImageData.URL = url.(string)
			}
			if basicAuthList, ok := dockerImageListItem["basic_auth"]; ok {
				basicAuthListInterface := basicAuthList.([]interface{})
				if len(basicAuthListInterface) > 0 {
					basicAuthListItem := basicAuthListInterface[0].(map[string]interface{})
					dockerAuthData := model.DockerBasicAuth{}
					if username, ok := basicAuthListItem["username"]; ok {
						dockerAuthData.Username = username.(string)
					}
					if password, ok := basicAuthListItem["password"]; ok {
						dockerAuthData.Password = password.(string)
					}
					dockerImageData.BasicAuth = &dockerAuthData
				}
			}
			cluster.DockerImage = &dockerImageData
		}
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
		cluster.InstancePoolID = instancePoolID.(string)
	}

	//Deal with single user name
	if singleUserName, ok := d.GetOk(schemaAttPrefix + "single_user_name"); ok {
		cluster.SingleUserName = singleUserName.(string)
	}

	//Deal with idempotency token
	if idempotencyToken, ok := d.GetOk(schemaAttPrefix + "idempotency_token"); ok {
		cluster.IdempotencyToken = idempotencyToken.(string)
	}
	return cluster
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
