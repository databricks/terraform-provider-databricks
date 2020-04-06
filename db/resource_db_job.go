package db

import (
	"bytes"
	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"sort"
	"strconv"
)

func resourceJob() *schema.Resource {
	return &schema.Resource{
		Create: resourceJobCreate,
		Read:   resourceJobRead,
		Update: resourceJobUpdate,
		Delete: resourceJobDelete,

		Schema: map[string]*schema.Schema{
			"existing_cluster_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cluster_num_workers": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cluster_autoscale": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
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
			},
			"cluster_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cluster_spark_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cluster_spark_conf": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"cluster_aws_attributes": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"availability": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"zone_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"spot_bid_price_percent": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"instance_profile_arn": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"first_on_demand": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"ebs_volume_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"ebs_volume_count": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"ebs_volume_size": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"cluster_driver_node_type_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cluster_node_type_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cluster_ssh_public_keys": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
				//	TODO: Validate less than 10 values
			},
			"cluster_custom_tags": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"cluster_log_conf": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
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
							Type:     schema.TypeString,
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
			"cluster_init_scripts": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
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
				Set: clusterInitScriptHash,
				//	Validate less than 10 values
			},
			"cluster_docker_image": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"url": {
							Type:     schema.TypeString,
							Optional: true,
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
			"cluster_spark_env_vars": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"cluster_autotermination_minutes": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cluster_enable_elastic_disk": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"cluster_instance_pool_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"library_jar": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"library_egg": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"library_whl": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"library_pypi": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
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
					},
				},
				Set: mapKeysHash,
			},
			"library_maven": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
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
					},
				},
				Set: mapKeysHash,
			},
			"library_cran": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
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
					},
				},
				Set: mapKeysHash,
			},
			"notebook_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"notebook_base_parameters": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"jar_uri": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"jar_main_class_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"jar_parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"python_file": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"python_parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"spark_submit_parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"email_on_start": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"email_on_success": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"email_on_failure": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"no_alert_for_skipped_runs": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"timeout_seconds": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_retries": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"min_retry_interval_millis": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"retry_on_timeout": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"schedule": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"quartz_cron_expression": {
							Type:     schema.TypeString,
							Required: true,
						},
						"timezone_id": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"max_concurrent_runs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"job_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"creator_user_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_time": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceJobCreate(d *schema.ResourceData, m interface{}) error {

	client := m.(service.DBApiClient)

	jobSettings := parseSchemaToJobSettings(d)
	job, err := client.Jobs().Create(jobSettings)
	if err != nil {
		return err
	}
	log.Println(job.JobId)
	d.SetId(strconv.Itoa(int(job.JobId)))
	return resourceJobRead(d, m)
}

func resourceJobRead(d *schema.ResourceData, m interface{}) error {
	client := m.(service.DBApiClient)
	id := d.Id()
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return err
	}

	job, err := client.Jobs().Read(idInt)
	if err != nil {
		return err
	}

	if _, ok := d.GetOk("existing_cluster_id"); ok {
		err := d.Set("existing_cluster_id", job.Settings.ExistingClusterId)
		if err != nil {
			return err
		}
	}

	if _, ok := d.GetOk("name"); ok {
		err := d.Set("name", job.Settings.Name)
		if err != nil {
			return err
		}
	}

	if job.Settings.NewCluster != nil {
		if _, ok := d.GetOk("cluster_num_workers"); ok {
			err := d.Set("cluster_num_workers", job.Settings.NewCluster.NumWorkers)
			if err != nil {
				return err
			}
		}

		if _, ok := d.GetOk("cluster_autoscale"); ok {
			if job.Settings.NewCluster.Autoscale != nil {
				autoscale := map[string]string{}
				autoscale["min_workers"] = strconv.Itoa(int(job.Settings.NewCluster.Autoscale.MinWorkers))
				autoscale["max_workers"] = strconv.Itoa(int(job.Settings.NewCluster.Autoscale.MaxWorkers))
				err := d.Set("cluster_autoscale", autoscale)
				if err != nil {
					return err
				}
			}
		}

		if _, ok := d.GetOk("cluster_name"); ok && len(job.Settings.NewCluster.ClusterName) > 0 {
			err = d.Set("cluster_name", job.Settings.NewCluster.ClusterName)
			if err != nil {
				return err
			}
		}

		if _, ok := d.GetOk("cluster_spark_version"); ok && len(job.Settings.NewCluster.SparkVersion) > 0 {
			err = d.Set("cluster_spark_version", job.Settings.NewCluster.SparkVersion)
			if err != nil {
				return err
			}
		}

		if _, ok := d.GetOk("cluster_spark_conf"); ok && job.Settings.NewCluster.SparkConf != nil && len(job.Settings.NewCluster.SparkConf) > 0 {
			err = d.Set("cluster_spark_conf", job.Settings.NewCluster.SparkConf)
			if err != nil {
				return err
			}
		}

		if job.Settings.NewCluster.AwsAttributes != nil {
			awsAtts := map[string]string{}
			if _, ok := d.GetOk("cluster_aws_attributes.availability"); ok {
				awsAtts["availability"] = string(job.Settings.NewCluster.AwsAttributes.Availability)
			}
			if _, ok := d.GetOk("cluster_aws_attributes.zone_id"); ok {
				awsAtts["zone_id"] = job.Settings.NewCluster.AwsAttributes.ZoneID
			}
			if _, ok := d.GetOk("cluster_aws_attributes.spot_bid_price_percent"); ok {
				awsAtts["spot_bid_price_percent"] = strconv.Itoa(int(job.Settings.NewCluster.AwsAttributes.SpotBidPricePercent))
			}
			if _, ok := d.GetOk("cluster_aws_attributes.instance_profile_arn"); ok {
				awsAtts["instance_profile_arn"] = job.Settings.NewCluster.AwsAttributes.InstanceProfileArn
			}
			if _, ok := d.GetOk("cluster_aws_attributes.first_on_demand"); ok {
				awsAtts["first_on_demand"] = strconv.Itoa(int(job.Settings.NewCluster.AwsAttributes.FirstOnDemand))
			}
			if _, ok := d.GetOk("cluster_aws_attributes.ebs_volume_type"); ok {
				awsAtts["ebs_volume_type"] = string(job.Settings.NewCluster.AwsAttributes.EbsVolumeType)
			}
			if _, ok := d.GetOk("cluster_aws_attributes.ebs_volume_count"); ok {
				awsAtts["ebs_volume_count"] = strconv.Itoa(int(job.Settings.NewCluster.AwsAttributes.EbsVolumeCount))
			}
			if _, ok := d.GetOk("cluster_aws_attributes.ebs_volume_size"); ok {
				awsAtts["ebs_volume_size"] = strconv.Itoa(int(job.Settings.NewCluster.AwsAttributes.EbsVolumeSize))
			}
			err = d.Set("cluster_aws_attributes", awsAtts)
			if err != nil {
				return err
			}
		}

		if _, ok := d.GetOk("cluster_driver_node_type_id"); ok && len(job.Settings.NewCluster.DriverNodeTypeID) > 0 {
			err = d.Set("cluster_driver_node_type_id", job.Settings.NewCluster.DriverNodeTypeID)
			if err != nil {
				return err
			}
		}

		if _, ok := d.GetOk("cluster_node_type_id"); ok && len(job.Settings.NewCluster.NodeTypeID) > 0 {
			err = d.Set("cluster_node_type_id", job.Settings.NewCluster.NodeTypeID)
			if err != nil {
				return err
			}
		}

		if len(job.Settings.NewCluster.SSHPublicKeys) > 0 {
			err = d.Set("cluster_ssh_public_keys", job.Settings.NewCluster.SSHPublicKeys)
			if err != nil {
				return err
			}
		}

		if len(job.Settings.NewCluster.CustomTags) > 0 {
			err = d.Set("cluster_custom_tags", job.Settings.NewCluster.CustomTags)
			if err != nil {
				return err
			}
		}

		if job.Settings.NewCluster.ClusterLogConf != nil {
			clusterLogConf := map[string]string{}
			if job.Settings.NewCluster.ClusterLogConf.Dbfs != nil {
				clusterLogConf["dbfs_destination"] = job.Settings.NewCluster.ClusterLogConf.Dbfs.Destination
			} else {
				if _, ok := d.GetOk("cluster_log_conf.s3_destination"); ok {
					clusterLogConf["s3_destination"] = job.Settings.NewCluster.ClusterLogConf.S3.Destination
				}
				if _, ok := d.GetOk("cluster_log_conf.s3_region"); ok {
					clusterLogConf["s3_region"] = job.Settings.NewCluster.ClusterLogConf.S3.Region
				}
				if _, ok := d.GetOk("cluster_log_conf.s3_endpoint"); ok {
					clusterLogConf["s3_endpoint"] = job.Settings.NewCluster.ClusterLogConf.S3.Endpoint
				}
				if _, ok := d.GetOk("cluster_log_conf.s3_enable_encryption"); ok {
					clusterLogConf["s3_enable_encryption"] = strconv.FormatBool(job.Settings.NewCluster.ClusterLogConf.S3.EnableEncryption)
				}
				if _, ok := d.GetOk("cluster_log_conf.s3_encryption_type"); ok {
					clusterLogConf["s3_encryption_type"] = job.Settings.NewCluster.ClusterLogConf.S3.EncryptionType
				}
				if _, ok := d.GetOk("cluster_log_conf.s3_kms_key"); ok {
					clusterLogConf["s3_kms_key"] = job.Settings.NewCluster.ClusterLogConf.S3.KmsKey
				}
				if _, ok := d.GetOk("cluster_log_conf.s3_canned_acl"); ok {
					clusterLogConf["s3_canned_acl"] = job.Settings.NewCluster.ClusterLogConf.S3.CannedACL
				}
			}
			err = d.Set("cluster_log_conf", clusterLogConf)
			if err != nil {
				return err
			}
		}

		if len(job.Settings.NewCluster.InitScripts) > 0 {
			var listOfInitScripts []map[string]string
			for _, v := range job.Settings.NewCluster.InitScripts {
				initScriptStorageConfig := map[string]string{}
				if v.Dbfs != nil {
					initScriptStorageConfig["dbfs_destination"] = v.Dbfs.Destination
				} else {
					if len(v.S3.Destination) > 0 {
						initScriptStorageConfig["s3_destination"] = v.S3.Destination
					}
					if len(v.S3.Region) > 0 {
						initScriptStorageConfig["s3_region"] = v.S3.Region
					}
					if len(v.S3.Endpoint) > 0 {
						initScriptStorageConfig["s3_endpoint"] = v.S3.Endpoint
					}
				}
				listOfInitScripts = append(listOfInitScripts, initScriptStorageConfig)
			}
			err = d.Set("init_scripts", listOfInitScripts)
			if err != nil {
				return err
			}
		}

		if job.Settings.NewCluster.DockerImage != nil {
			dockerImage := map[string]string{}
			dockerImage["url"] = job.Settings.NewCluster.DockerImage.Url
			if job.Settings.NewCluster.DockerImage.BasicAuth != nil {
				dockerImage["username"] = job.Settings.NewCluster.DockerImage.BasicAuth.Username
				dockerImage["password"] = job.Settings.NewCluster.DockerImage.BasicAuth.Password
			}
			err = d.Set("cluster_aws_attributes", dockerImage)
			if err != nil {
				return err
			}
		}

		if len(job.Settings.NewCluster.SparkEnvVars) > 0 {
			err = d.Set("cluster_spark_env_vars", job.Settings.NewCluster.SparkEnvVars)
			if err != nil {
				return err
			}
		}

		err = d.Set("cluster_enable_elastic_disk", job.Settings.NewCluster.EnableElasticDisk)
		if err != nil {
			return err
		}

		err = d.Set("cluster_enable_elastic_disk", job.Settings.NewCluster.EnableElasticDisk)
		if err != nil {
			return err
		}
	}

	if _, ok := d.GetOk("instance_pool_id"); ok {
		err := d.Set("instance_pool_id", job.Settings.NewCluster.InstancePoolId)
		if err != nil {
			return err
		}
	}

	libraries := job.Settings.Libraries
	var jars []string
	var eggs []string
	var whls []string
	var pypi []map[string]string
	var maven []map[string]interface{}
	var cran []map[string]string

	for _, lib := range libraries {
		if len(lib.Jar) > 0 {
			jars = append(jars, lib.Jar)
		}
		if len(lib.Egg) > 0 {
			eggs = append(eggs, lib.Jar)
		}
		if len(lib.Whl) > 0 {
			whls = append(whls, lib.Whl)
		}
		if lib.Pypi != nil {
			pypiPackage := map[string]string{}
			if len(lib.Pypi.Package) > 0 {
				pypiPackage["package"] = lib.Pypi.Package
			}
			if len(lib.Pypi.Repo) > 0 {
				pypiPackage["repo"] = lib.Pypi.Repo
			}
			pypi = append(pypi, pypiPackage)
		}
		if lib.Maven != nil {
			mvnPackage := map[string]interface{}{}
			if len(lib.Maven.Coordinates) > 0 {
				mvnPackage["coordinates"] = lib.Maven.Coordinates
			}
			if len(lib.Maven.Repo) > 0 {
				mvnPackage["repo"] = lib.Maven.Repo
			}
			if len(lib.Maven.Exclusions) > 0 {
				mvnPackage["exclusions"] = lib.Maven.Exclusions
			}
			maven = append(maven, mvnPackage)
		}
		if lib.Cran != nil {
			cranPackage := map[string]string{}
			if len(lib.Cran.Package) > 0 {
				cranPackage["package"] = lib.Cran.Package
			}
			if len(lib.Cran.Repo) > 0 {
				cranPackage["repo"] = lib.Cran.Repo
			}
			cran = append(cran, cranPackage)
		}
	}

	if len(jars) > 0 {
		err := d.Set("library_jar", jars)
		if err != nil {
			return err
		}
	}

	if len(eggs) > 0 {
		err := d.Set("library_egg", eggs)
		if err != nil {
			return err
		}
	}

	if len(whls) > 0 {
		err := d.Set("library_whl", whls)
		if err != nil {
			return err
		}
	}

	if len(pypi) > 0 {
		err := d.Set("library_pypi", pypi)
		if err != nil {
			return err
		}
	}
	if len(maven) > 0 {
		err := d.Set("library_maven", maven)
		if err != nil {
			return err
		}
	}
	if len(cran) > 0 {
		err := d.Set("library_cran", cran)
		if err != nil {
			return err
		}
	}

	if _, ok := d.GetOk("notebook_path"); ok {
		err := d.Set("notebook_path", job.Settings.NotebookTask.NotebookPath)
		if err != nil {
			return err
		}
	}

	if _, ok := d.GetOk("notebook_base_parameters"); ok {
		err := d.Set("notebook_base_parameters", job.Settings.NotebookTask.BaseParameters)
		if err != nil {
			return err
		}
	}

	if _, ok := d.GetOk("jar_uri"); ok {
		err := d.Set("jar_uri", job.Settings.SparkJarTask.JarUri)
		if err != nil {
			return err
		}
	}

	if _, ok := d.GetOk("jar_main_class_name"); ok {
		err := d.Set("jar_main_class_name", job.Settings.SparkJarTask.MainClassName)
		if err != nil {
			return err
		}
	}

	if _, ok := d.GetOk("jar_parameters"); ok {
		err := d.Set("jar_parameters", job.Settings.SparkJarTask.Parameters)
		if err != nil {
			return err
		}
	}

	if _, ok := d.GetOk("python_file"); ok {
		err := d.Set("python_file", job.Settings.SparkPythonTask.PythonFile)
		if err != nil {
			return err
		}
	}

	if _, ok := d.GetOk("python_parameters"); ok {
		err := d.Set("python_parameters", job.Settings.SparkPythonTask.Parameters)
		if err != nil {
			return err
		}
	}

	if _, ok := d.GetOk("spark_submit_parameters"); ok {
		err := d.Set("spark_submit_parameters", job.Settings.SparkSubmitTask.Parameters)
		if err != nil {
			return err
		}
	}

	if _, ok := d.GetOk("email_on_start"); ok {
		err := d.Set("email_on_start", job.Settings.EmailNotifications.OnStart)
		if err != nil {
			return err
		}
	}

	if _, ok := d.GetOk("email_on_success"); ok {
		err := d.Set("email_on_success", job.Settings.EmailNotifications.OnSuccess)
		if err != nil {
			return err
		}
	}

	if _, ok := d.GetOk("email_on_failure"); ok {
		err := d.Set("email_on_failure", job.Settings.EmailNotifications.OnFailure)
		if err != nil {
			return err
		}
	}

	if _, ok := d.GetOk("no_alert_for_skipped_runs"); ok {
		err := d.Set("no_alert_for_skipped_runs", job.Settings.EmailNotifications.NoAlertForSkippedRuns)
		if err != nil {
			return err
		}
	}

	if _, ok := d.GetOk("timeout_seconds"); ok {
		err := d.Set("timeout_seconds", job.Settings.TimeoutSeconds)
		if err != nil {
			return err
		}
	}

	if _, ok := d.GetOk("max_retries"); ok {
		err := d.Set("max_retries", job.Settings.MaxRetries)
		if err != nil {
			return err
		}
	}

	if _, ok := d.GetOk("min_retry_interval_millis"); ok {
		err := d.Set("min_retry_interval_millis", job.Settings.MinRetryIntervalMillis)
		if err != nil {
			return err
		}
	}

	if _, ok := d.GetOk("retry_on_timeout"); ok {
		err := d.Set("retry_on_timeout", job.Settings.RetryOnTimeout)
		if err != nil {
			return err
		}
	}

	if scheduleMap, ok := d.GetOk("schedule"); ok {
		schedule := scheduleMap.(map[string]interface{})
		sched := map[string]string{}
		if _, ok := schedule["quartz_cron_expression"]; ok {
			sched["quartz_cron_expression"] = job.Settings.Schedule.QuartzCronExpression
		}
		if _, ok := schedule["timezone_id"]; ok {
			sched["timezone_id"] = job.Settings.Schedule.TimezoneId
		}
		err := d.Set("schedule", sched)
		if err != nil {
			return err
		}
	}

	if _, ok := d.GetOk("max_concurrent_runs"); ok {
		err := d.Set("max_concurrent_runs", job.Settings.MaxConcurrentRuns)
		if err != nil {
			return err
		}
	}

	if _, ok := d.GetOk("job_id"); ok {
		err := d.Set("job_id", job.JobId)
		if err != nil {
			return err
		}
	}

	if _, ok := d.GetOk("creator_user_name"); ok {
		err := d.Set("creator_user_name", job.CreatorUserName)
		if err != nil {
			return err
		}
	}

	if _, ok := d.GetOk("created_time"); ok {
		err := d.Set("created_time", job.CreatedTime)
		if err != nil {
			return err
		}
	}

	return err
}

func resourceJobUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(service.DBApiClient)
	id := d.Id()
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return err
	}
	jobSettings := parseSchemaToJobSettings(d)

	err = client.Jobs().Update(idInt, jobSettings)
	if err != nil {
		return err
	}
	return resourceJobRead(d, m)
}

func resourceJobDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(service.DBApiClient)
	id := d.Id()
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return err
	}
	err = client.Jobs().Delete(idInt)
	return err
}

func parseSchemaToJobSettings(d *schema.ResourceData) model.JobSettings {

	var jobSettings model.JobSettings

	if existingClusterId, ok := d.GetOk("existing_cluster_id"); ok {
		jobSettings.ExistingClusterId = existingClusterId.(string)
	}

	cluster := parseSchemaToNewJobCluster(d)
	jobSettings.NewCluster = cluster

	if name, ok := d.GetOk("name"); ok {
		jobSettings.Name = name.(string)
	}
	libraries := parseSchemaToLibraries(d)
	jobSettings.Libraries = libraries

	nbTask := parseSchemaToNotebookTask(d)
	jobSettings.NotebookTask = nbTask

	jarTask := parseSchemaToSparkJarTask(d)
	jobSettings.SparkJarTask = jarTask

	pyTask := parseSchemaToSparkPythonTask(d)
	jobSettings.SparkPythonTask = pyTask

	sparkSubmitTask := parseSchemaToSparkSubmitTask(d)
	jobSettings.SparkSubmitTask = sparkSubmitTask

	var email model.JobEmailNotifications
	if emailOnStart, ok := d.GetOk("email_on_start"); ok {
		email.OnStart = convertListInterfaceToString(emailOnStart.(*schema.Set).List())
	}
	if emailOnSuccess, ok := d.GetOk("email_on_success"); ok {
		email.OnSuccess = convertListInterfaceToString(emailOnSuccess.(*schema.Set).List())
	}
	if emailOnFailure, ok := d.GetOk("email_on_failure"); ok {
		email.OnFailure = convertListInterfaceToString(emailOnFailure.(*schema.Set).List())
	}
	if noAlertForSkippedRuns, ok := d.GetOk("no_alert_for_skipped_runs"); ok {
		email.NoAlertForSkippedRuns = noAlertForSkippedRuns.(bool)
	}
	jobSettings.EmailNotifications = &email

	if timeoutSeconds, ok := d.GetOk("timeout_seconds"); ok {
		intVal, _ := timeoutSeconds.(int)
		jobSettings.TimeoutSeconds = int32(intVal)
	}

	if maxRetries, ok := d.GetOk("max_retries"); ok {
		intVal, _ := maxRetries.(int)
		jobSettings.MaxRetries = int32(intVal)
	}

	if minRetryIntervalMillis, ok := d.GetOk("min_retry_interval_millis"); ok {
		intVal, _ := minRetryIntervalMillis.(int)
		jobSettings.MinRetryIntervalMillis = int32(intVal)
	}

	if retryOnTimeout, ok := d.GetOk("retry_on_timeout"); ok {
		jobSettings.RetryOnTimeout = retryOnTimeout.(bool)
	}
	if schedule, ok := d.GetOk("schedule"); ok {
		schedule := convertMapStringInterfaceToStringString(schedule.(map[string]interface{}))
		jobSettings.Schedule = &model.CronSchedule{
			QuartzCronExpression: schedule["quartz_cron_expression"],
			TimezoneId:           schedule["timezone_id"],
		}

	}
	if maxConcurrentRuns, ok := d.GetOk("max_concurrent_runs"); ok {
		intVal, _ := maxConcurrentRuns.(int)
		jobSettings.MaxConcurrentRuns = int32(intVal)
	}
	return jobSettings
}

func parseSchemaToNotebookTask(d *schema.ResourceData) *model.NotebookTask {
	var notebookTask model.NotebookTask
	if path, ok := d.GetOk("notebook_path"); ok {
		notebookTask.NotebookPath = path.(string)
	}

	if notebookParams, ok := d.GetOk("notebook_base_parameters"); ok {
		notebookTask.BaseParameters = convertMapStringInterfaceToStringString(notebookParams.(map[string]interface{}))
	}
	return &notebookTask
}

func parseSchemaToSparkJarTask(d *schema.ResourceData) *model.SparkJarTask {
	var sparkJarTask model.SparkJarTask
	if uri, ok := d.GetOk("jar_uri"); ok {
		sparkJarTask.JarUri = uri.(string)
	}
	if cName, ok := d.GetOk("jar_main_class_name"); ok {
		sparkJarTask.MainClassName = cName.(string)
	} else {
		return nil
	}

	if jarParams, ok := d.GetOk("jar_parameters"); ok {
		sparkJarTask.Parameters = convertListInterfaceToString(jarParams.([]interface{}))
	}
	return &sparkJarTask
}

func parseSchemaToSparkPythonTask(d *schema.ResourceData) *model.SparkPythonTask {
	var sparkPythonTask model.SparkPythonTask
	if file, ok := d.GetOk("python_file"); ok {
		sparkPythonTask.PythonFile = file.(string)
	} else {
		return nil
	}

	if pythonParams, ok := d.GetOk("python_parameters"); ok {
		sparkPythonTask.Parameters = convertListInterfaceToString(pythonParams.([]interface{}))
	}
	return &sparkPythonTask
}

func parseSchemaToSparkSubmitTask(d *schema.ResourceData) *model.SparkSubmitTask {
	var sparkSubmitTask model.SparkSubmitTask
	if sparkSubmitParams, ok := d.GetOk("spark_submit_parameters"); ok {
		sparkSubmitTask.Parameters = convertListInterfaceToString(sparkSubmitParams.([]interface{}))
	} else {
		return nil
	}
	return &sparkSubmitTask
}

func parseSchemaToLibraries(d *schema.ResourceData) []model.JobLibrary {
	var libraryList []model.JobLibrary
	if jars, ok := d.GetOk("library_jar"); ok {
		libraries := jars.(*schema.Set).List()
		for _, library := range libraries {
			thisLibrary := model.JobLibrary{
				Jar: library.(string),
			}
			libraryList = append(libraryList, thisLibrary)
		}
	}
	if eggs, ok := d.GetOk("library_egg"); ok {
		libraries := eggs.(*schema.Set).List()
		for _, library := range libraries {
			thisLibrary := model.JobLibrary{
				Egg: library.(string),
			}
			libraryList = append(libraryList, thisLibrary)
		}
	}
	if whls, ok := d.GetOk("library_whl"); ok {
		libraries := whls.(*schema.Set).List()
		for _, library := range libraries {
			thisLibrary := model.JobLibrary{
				Whl: library.(string),
			}
			libraryList = append(libraryList, thisLibrary)
		}
	}
	if pypis, ok := d.GetOk("library_pypi"); ok {
		libraries := pypis.(*schema.Set).List()
		for _, library := range libraries {
			libraryMap := library.(map[string]interface{})
			var pypi model.PyPiLibrary
			if pkg, ok := libraryMap["package"]; ok {
				pypi.Package = pkg.(string)
			}
			if repo, ok := libraryMap["repo"]; ok {
				pypi.Repo = repo.(string)
			}
			thisLibrary := model.JobLibrary{
				Pypi: &pypi,
			}
			libraryList = append(libraryList, thisLibrary)
		}
	}
	if mavens, ok := d.GetOk("library_maven"); ok {
		libraries := mavens.(*schema.Set).List()
		for _, library := range libraries {
			libraryMap := library.(map[string]interface{})
			var maven model.MavenLibrary
			if coordinates, ok := libraryMap["coordinates"]; ok {
				maven.Coordinates = coordinates.(string)
			}
			if repo, ok := libraryMap["repo"]; ok {
				maven.Repo = repo.(string)
			}
			if exclusions, ok := libraryMap["exclusions"]; ok {
				maven.Exclusions = convertListInterfaceToString(exclusions.([]interface{}))
			}
			thisLibrary := model.JobLibrary{
				Maven: &maven,
			}
			libraryList = append(libraryList, thisLibrary)
		}
	}
	if crans, ok := d.GetOk("library_cran"); ok {
		libraries := crans.(*schema.Set).List()
		for _, library := range libraries {
			libraryMap := library.(map[string]interface{})
			var cran model.CranLibrary
			if pkg, ok := libraryMap["package"]; ok {
				cran.Package = pkg.(string)
			}
			if repo, ok := libraryMap["repo"]; ok {
				cran.Repo = repo.(string)
			}
			thisLibrary := model.JobLibrary{
				Cran: &cran,
			}
			libraryList = append(libraryList, thisLibrary)
		}
	}
	return libraryList
}

func parseSchemaToNewJobCluster(d *schema.ResourceData) *model.Cluster {
	var cluster model.Cluster

	var workers bool
	var sparkVer bool
	var nodeType bool
	//Deal with Num workers
	if numWorkers, ok := d.GetOk("cluster_num_workers"); ok {
		workers = true
		cluster.NumWorkers = int32(numWorkers.(int))
	}

	//Deal with auto scaling options
	var autoScale model.AutoScale
	if autoscale, ok := d.GetOk("cluster_autoscale"); ok {
		autoScaleOptions := autoscale.(map[string]interface{})
		if minWorkers, ok := autoScaleOptions["min_workers"]; ok {
			minVal, _ := strconv.ParseInt(minWorkers.(string), 10, 32)
			autoScale.MinWorkers = int32(minVal)
		}
		if maxWorkers, ok := autoScaleOptions["max_workers"]; ok {
			maxVal, _ := strconv.ParseInt(maxWorkers.(string), 10, 32)
			autoScale.MaxWorkers = int32(maxVal)
		}
		cluster.Autoscale = &autoScale
		workers = true
	}

	//Deal with cluster name
	if clusterName, ok := d.GetOk("cluster_cluster_name"); ok {
		cluster.ClusterName = clusterName.(string)
	}

	//Deal with spark versions
	if sparkVersion, ok := d.GetOk("cluster_spark_version"); ok {
		cluster.SparkVersion = sparkVersion.(string)
		sparkVer = true
	}

	//Deal with spark confs
	if sparkConf, ok := d.GetOk("cluster_spark_conf"); ok {
		cluster.SparkConf = convertMapStringInterfaceToStringString(sparkConf.(map[string]interface{}))
	}

	//Deal with aws attributes for aws deployment
	var awsAttributes model.AwsAttributes
	if awsAttributesSchema, ok := d.GetOk("cluster_aws_attributes"); ok {
		awsAttributesMap := awsAttributesSchema.(map[string]interface{})
		if availability, ok := awsAttributesMap["availability"]; ok {
			awsAttributes.Availability = model.AwsAvailability(availability.(string))
		}
		if zoneId, ok := awsAttributesMap["zone_id"]; ok {
			awsAttributes.ZoneID = zoneId.(string)
		}
		if spotBidPricePercent, ok := awsAttributesMap["spot_bid_price_percent"]; ok {
			val, _ := strconv.ParseInt(spotBidPricePercent.(string), 10, 32)
			awsAttributes.SpotBidPricePercent = int32(val)
		} else {
			awsAttributes.SpotBidPricePercent = int32(100)
		}
		if instanceProfileArn, ok := awsAttributesMap["instance_profile_arn"]; ok {
			awsAttributes.InstanceProfileArn = instanceProfileArn.(string)
		}
		if firstOnDemand, ok := awsAttributesMap["first_on_demand"]; ok {
			val, _ := strconv.ParseInt(firstOnDemand.(string), 10, 32)
			awsAttributes.FirstOnDemand = int32(val)
		}
		if ebsVolumeType, ok := awsAttributesMap["ebs_volume_type"]; ok {
			awsAttributes.EbsVolumeType = model.EbsVolumeType(ebsVolumeType.(string))
		}
		if ebsVolumeCount, ok := awsAttributesMap["ebs_volume_count"]; ok {
			val, _ := strconv.ParseInt(ebsVolumeCount.(string), 10, 32)
			awsAttributes.FirstOnDemand = int32(val)
		}
		if ebsVolumeSize, ok := awsAttributesMap["ebs_volume_size"]; ok {
			val, _ := strconv.ParseInt(ebsVolumeSize.(string), 10, 32)
			awsAttributes.EbsVolumeSize = int32(val)
		}
		cluster.AwsAttributes = &awsAttributes
	}

	//Deal with driver node type id
	if driverNodeTypeId, ok := d.GetOk("cluster_driver_node_type_id"); ok {
		cluster.DriverNodeTypeID = driverNodeTypeId.(string)
	}

	//Deal with worker node type id
	if nodeTypeId, ok := d.GetOk("cluster_node_type_id"); ok {
		cluster.NodeTypeID = nodeTypeId.(string)
		nodeType = true
	}

	//Deal with worker ssh public keys
	if sshPublicKeys, ok := d.GetOk("cluster_ssh_public_keys"); ok {
		cluster.SSHPublicKeys = convertListInterfaceToString(sshPublicKeys.(*schema.Set).List())
	}

	//Deal with worker custom tags
	if customTags, ok := d.GetOk("cluster_custom_tags"); ok {
		tags := customTags.(map[string]interface{})
		cluster.CustomTags = convertMapStringInterfaceToStringString(tags)
	}

	//Deal with worker cluster log config
	var clusterLogConf model.StorageInfo
	if clusterLogConfSchema, ok := d.GetOk("cluster_log_conf"); ok {
		clusterLogConfMap := clusterLogConfSchema.(map[string]interface{})
		if dbfsDestination, ok := clusterLogConfMap["dbfs_destination"]; ok {
			var dbfsStorage model.DbfsStorageInfo
			dbfsStorage.Destination = dbfsDestination.(string)
			clusterLogConf.Dbfs = &dbfsStorage
		} else {
			var s3Storage model.S3StorageInfo
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
				b, _ := strconv.ParseBool(s3EnableEncryption.(string))
				s3Storage.EnableEncryption = b
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
		cluster.ClusterLogConf = &clusterLogConf
	}

	//Deal with worker init script setup
	if initScripts, ok := d.GetOk("cluster_init_scripts"); ok {
		initScripts := initScripts.(*schema.Set).List()
		var initScriptsLocations []model.StorageInfo
		for _, v := range initScripts {
			initScript := v.(map[string]interface{})
			var storageInfo model.StorageInfo
			if dbfsDestination, ok := initScript["dbfs_destination"]; ok {
				var dbfsStorage model.DbfsStorageInfo
				dbfsStorage.Destination = dbfsDestination.(string)
				storageInfo.Dbfs = &dbfsStorage
				initScriptsLocations = append(initScriptsLocations, storageInfo)
			} else {
				var s3Storage model.S3StorageInfo
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
	var dockerImageData model.DockerImage
	if dockerImage, ok := d.GetOk("cluster_docker_image"); ok {
		dockerImageConf := dockerImage.(map[string]interface{})
		if url, ok := dockerImageConf["url"]; ok {
			dockerImageData.Url = url.(string)
		}
		var dockerAuthData model.DockerBasicAuth
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
	if sparkEnv, ok := d.GetOk("cluster_spark_env_vars"); ok {
		cluster.SparkEnvVars = convertMapStringInterfaceToStringString(sparkEnv.(map[string]interface{}))
	}

	//Deal with enable elastic disk
	if enableElasticDisk, ok := d.GetOk("enable_elastic_disk"); ok {
		cluster.EnableElasticDisk = enableElasticDisk.(bool)
	}

	//Deal with instance pool id
	if instancePoolID, ok := d.GetOk("instance_pool_id"); ok {
		cluster.InstancePoolId = instancePoolID.(string)
	}

	if workers && sparkVer && nodeType {
		return &cluster
	} else {
		return nil
	}

}

func mapKeysHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	keys := make([]string, 0)
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		value, ok := m[k].(string)
		if ok {
			buf.WriteString(value)
		} else {
			list, ok := m[k].([]interface{})
			if ok {
				for _, val := range list {
					buf.WriteString(val.(string))
				}
			}
		}

	}
	return hashcode.String(buf.String())
}
