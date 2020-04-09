package db

import (
	"fmt"
	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
	"strings"
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
			"new_cluster": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"num_workers": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"autoscale": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							MaxItems: 1,
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
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"spark_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"node_type_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssh_public_keys": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
							//	TODO: Validate less than 10 values
						},
						"custom_tags": &schema.Schema{
							Type:     schema.TypeMap,
							Optional: true,
						},
						"cluster_log_conf": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							MaxItems: 1,
							//ConfigMode: schema.SchemaConfigModeAttr,
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
							Type:     schema.TypeSet,
							Optional: true,
							//ConfigMode: schema.SchemaConfigModeAttr,
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
							//	Validate less than 10 values
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
							Computed: true,
						},
						"enable_elastic_disk": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"instance_pool_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
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
			},
			"notebook_path": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"jar_main_class_name", "spark_submit_parameters", "python_file"},
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
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"python_file", "notebook_path", "spark_submit_parameters"},
			},
			"jar_parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"python_file": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"jar_main_class_name", "notebook_path", "spark_submit_parameters"},
			},
			"python_parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"spark_submit_parameters": &schema.Schema{
				Type:          schema.TypeList,
				Optional:      true,
				Elem:          &schema.Schema{Type: schema.TypeString},
				ConflictsWith: []string{"jar_main_class_name", "notebook_path", "python_file"},
			},
			"email_notifications": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// Suppress order change for the list
					if k == "email_notifications.#" {
						return true
					}
					if old != new {
						return false
					}
					return true
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"on_start": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
						"on_success": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"on_failure": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"no_alert_for_skipped_runs": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
				MaxItems: 1,
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
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
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
		if isJobMissing(err.Error(), id) {
			log.Printf("Missing job with id: %s.", id)
			d.SetId("")
			return nil
		}
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
		newClusterSettings := map[string]interface{}{}
		newClusterSettings["num_workers"] = job.Settings.NewCluster.NumWorkers

		autoscale := map[string]int{}
		autoscale["min_workers"] = int(job.Settings.NewCluster.Autoscale.MinWorkers)
		autoscale["max_workers"] = int(job.Settings.NewCluster.Autoscale.MaxWorkers)
		autoscaleSet := []map[string]int{autoscale}
		newClusterSettings["autoscale"] = autoscaleSet

		newClusterSettings["name"] = job.Settings.NewCluster.ClusterName

		newClusterSettings["spark_version"] = job.Settings.NewCluster.SparkVersion

		newClusterSettings["spark_conf"] = job.Settings.NewCluster.SparkConf

		awsAtts := map[string]interface{}{}
		awsAtts["availability"] = string(job.Settings.NewCluster.AwsAttributes.Availability)
		awsAtts["zone_id"] = job.Settings.NewCluster.AwsAttributes.ZoneID
		awsAtts["spot_bid_price_percent"] = int(job.Settings.NewCluster.AwsAttributes.SpotBidPricePercent)
		awsAtts["instance_profile_arn"] = job.Settings.NewCluster.AwsAttributes.InstanceProfileArn
		awsAtts["first_on_demand"] = int(job.Settings.NewCluster.AwsAttributes.FirstOnDemand)
		awsAtts["ebs_volume_type"] = string(job.Settings.NewCluster.AwsAttributes.EbsVolumeType)
		awsAtts["ebs_volume_count"] = int(job.Settings.NewCluster.AwsAttributes.EbsVolumeCount)
		awsAtts["ebs_volume_size"] = int(job.Settings.NewCluster.AwsAttributes.EbsVolumeSize)
		awsAttsSet := []map[string]interface{}{awsAtts}
		newClusterSettings["aws_attributes"] = awsAttsSet

		newClusterSettings["driver_node_type_id"] = job.Settings.NewCluster.DriverNodeTypeID

		newClusterSettings["node_type_id"] = job.Settings.NewCluster.DriverNodeTypeID

		newClusterSettings["ssh_public_keys"] = job.Settings.NewCluster.SSHPublicKeys

		newClusterSettings["custom_tags"] = job.Settings.NewCluster.CustomTags

		clusterLogConfMap := map[string]interface{}{}
		if job.Settings.NewCluster.ClusterLogConf != nil {
			if job.Settings.NewCluster.ClusterLogConf.Dbfs != nil {
				clusterLogConfMap["dbfs_destination"] = job.Settings.NewCluster.ClusterLogConf.Dbfs.Destination
			}

			if job.Settings.NewCluster.ClusterLogConf.S3 != nil {
				clusterLogConfMap["s3_destination"] = job.Settings.NewCluster.ClusterLogConf.S3.Destination
				clusterLogConfMap["s3_region"] = job.Settings.NewCluster.ClusterLogConf.S3.Region
				clusterLogConfMap["s3_endpoint"] = job.Settings.NewCluster.ClusterLogConf.S3.Endpoint
				clusterLogConfMap["s3_enable_encryption"] = job.Settings.NewCluster.ClusterLogConf.S3.EnableEncryption
				clusterLogConfMap["s3_encryption_type"] = job.Settings.NewCluster.ClusterLogConf.S3.EncryptionType
				clusterLogConfMap["s3_kms_key"] = job.Settings.NewCluster.ClusterLogConf.S3.KmsKey
				clusterLogConfMap["s3_canned_acl"] = job.Settings.NewCluster.ClusterLogConf.S3.CannedACL
			}
			clusterLogConfSet := []map[string]interface{}{clusterLogConfMap}
			newClusterSettings["cluster_log_conf"] = clusterLogConfSet
		} else {
			newClusterSettings["cluster_log_conf"] = nil
		}

		var listOfInitScripts []map[string]string
		for _, v := range job.Settings.NewCluster.InitScripts {
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
		newClusterSettings["init_scripts"] = listOfInitScripts

		dockerImage := map[string]string{}
		dockerImage["url"] = job.Settings.NewCluster.DockerImage.Url
		if job.Settings.NewCluster.DockerImage.BasicAuth != nil {
			dockerImage["username"] = job.Settings.NewCluster.DockerImage.BasicAuth.Username
			dockerImage["password"] = job.Settings.NewCluster.DockerImage.BasicAuth.Password
		}
		dockerImageSet := []map[string]string{dockerImage}
		newClusterSettings["docker_image"] = dockerImageSet

		newClusterSettings["spark_env_vars"] = job.Settings.NewCluster.SparkEnvVars

		newClusterSettings["autotermination_minutes"] = job.Settings.NewCluster.AutoterminationMinutes

		newClusterSettings["enable_elastic_disk"] = job.Settings.NewCluster.EnableElasticDisk

		newClusterSettings["instance_pool_id"] = job.Settings.NewCluster.InstancePoolId
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

	if job.Settings.NotebookTask != nil {
		err = d.Set("notebook_path", job.Settings.NotebookTask.NotebookPath)
		if err != nil {
			return err
		}

		err = d.Set("notebook_base_parameters", job.Settings.NotebookTask.BaseParameters)
		if err != nil {
			return err
		}
	} else {
		err = d.Set("notebook_path", nil)
		if err != nil {
			return err
		}
		err = d.Set("notebook_base_parameters", nil)
		if err != nil {
			return err
		}
	}

	if job.Settings.SparkJarTask != nil {
		err = d.Set("jar_uri", job.Settings.SparkJarTask.JarUri)
		if err != nil {
			return err
		}

		err = d.Set("jar_main_class_name", job.Settings.SparkJarTask.MainClassName)
		if err != nil {
			return err
		}

		err = d.Set("jar_parameters", job.Settings.SparkJarTask.Parameters)
		if err != nil {
			return err
		}
	} else {
		err = d.Set("jar_uri", nil)
		if err != nil {
			return err
		}
		err = d.Set("jar_main_class_name", nil)
		if err != nil {
			return err
		}
		err = d.Set("jar_parameters", nil)
		if err != nil {
			return err
		}
	}

	if job.Settings.SparkPythonTask != nil {
		err = d.Set("python_file", job.Settings.SparkPythonTask.PythonFile)
		if err != nil {
			return err
		}
		err = d.Set("python_parameters", job.Settings.SparkPythonTask.Parameters)
		if err != nil {
			return err
		}
	} else {
		err = d.Set("python_file", nil)
		if err != nil {
			return err
		}
		err = d.Set("python_parameters", nil)
		if err != nil {
			return err
		}
	}
	if job.Settings.SparkSubmitTask != nil {
		err = d.Set("spark_submit_parameters", job.Settings.SparkSubmitTask.Parameters)
		if err != nil {
			return err
		}
	} else {
		err = d.Set("spark_submit_parameters", nil)
		if err != nil {
			return err
		}
	}

	if job.Settings.EmailNotifications != nil {
		emailNotifiactions := map[string]interface{}{}
		if job.Settings.EmailNotifications.OnStart != nil {
			emailNotifiactions["on_start"] = job.Settings.EmailNotifications.OnStart
		} else {
			emailNotifiactions["on_start"] = nil
		}
		if job.Settings.EmailNotifications.OnFailure != nil {
			emailNotifiactions["on_failure"] = job.Settings.EmailNotifications.OnFailure
		} else {
			emailNotifiactions["on_failure"] = nil
		}
		if job.Settings.EmailNotifications.OnSuccess != nil {
			emailNotifiactions["on_success"] = job.Settings.EmailNotifications.OnSuccess
		} else {
			emailNotifiactions["on_success"] = nil
		}
		emailNotifiactions["no_alert_for_skipped_runs"] = job.Settings.EmailNotifications.NoAlertForSkippedRuns

		emailNotifiactionsSet := []map[string]interface{}{emailNotifiactions}
		err = d.Set("email_notifications", emailNotifiactionsSet)
		if err != nil {
			return err
		}
	} else {
		err = d.Set("email_notifications", nil)
		if err != nil {
			return err
		}
	}

	err = d.Set("timeout_seconds", job.Settings.TimeoutSeconds)
	if err != nil {
		return err
	}

	err = d.Set("max_retries", job.Settings.MaxRetries)
	if err != nil {
		return err
	}

	err = d.Set("min_retry_interval_millis", job.Settings.MinRetryIntervalMillis)
	if err != nil {
		return err
	}

	err = d.Set("retry_on_timeout", job.Settings.RetryOnTimeout)
	if err != nil {
		return err
	}

	if job.Settings.Schedule != nil {
		sched := map[string]string{}
		sched["quartz_cron_expression"] = job.Settings.Schedule.QuartzCronExpression

		sched["timezone_id"] = job.Settings.Schedule.TimezoneId

		schedSet := []map[string]string{sched}
		err = d.Set("schedule", schedSet)
		if err != nil {
			return err
		}
	} else {
		err = d.Set("schedule", nil)
		if err != nil {
			return err
		}
	}

	err = d.Set("max_concurrent_runs", job.Settings.MaxConcurrentRuns)
	if err != nil {
		return err
	}

	err = d.Set("job_id", job.JobId)
	if err != nil {
		return err
	}

	err = d.Set("creator_user_name", job.CreatorUserName)
	if err != nil {
		return err
	}

	err = d.Set("created_time", job.CreatedTime)
	if err != nil {
		return err
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

	cluster := parseSchemaToCluster(d, "new_cluster.0.")
	log.Println("Parse cluster")
	log.Println(cluster)
	log.Println(d.Get("new_cluster"))
	if numWorkers, ok := d.GetOk("new_cluster[0].num_workers"); ok {
		log.Println("num_workers")
		log.Println(numWorkers)
	}
	if numWorkers, ok := d.GetOk("new_cluster[0].autoscale"); ok {
		log.Println("num_workers")
		log.Println(numWorkers)
	}
	jobSettings.NewCluster = &cluster

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

	if emailNotificationsList, ok := d.GetOk("email_notifications"); ok {
		var email model.JobEmailNotifications
		emailNotificationsMap := getMapFromOneItemList(emailNotificationsList)
		if emailOnStart, ok := emailNotificationsMap["on_start"]; ok {
			email.OnStart = convertListInterfaceToString(emailOnStart.(*schema.Set).List())
		}
		if emailOnSuccess, ok := emailNotificationsMap["on_success"]; ok {
			email.OnSuccess = convertListInterfaceToString(emailOnSuccess.(*schema.Set).List())
		}
		if emailOnFailure, ok := emailNotificationsMap["on_failure"]; ok {
			email.OnFailure = convertListInterfaceToString(emailOnFailure.(*schema.Set).List())
		}
		if noAlertForSkippedRuns, ok := emailNotificationsMap["no_alert_for_skipped_runs"]; ok {
			email.NoAlertForSkippedRuns = noAlertForSkippedRuns.(bool)
		}
		jobSettings.EmailNotifications = &email

	}

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
		scheduleMap := getMapFromOneItemSet(schedule)
		jobSettings.Schedule = &model.CronSchedule{
			QuartzCronExpression: scheduleMap["quartz_cron_expression"].(string),
			TimezoneId:           scheduleMap["timezone_id"].(string),
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

func parseSchemaToLibraries(d *schema.ResourceData) []model.Library {
	var libraryList []model.Library
	if jars, ok := d.GetOk("library_jar"); ok {
		libraries := jars.(*schema.Set).List()
		for _, library := range libraries {
			thisLibrary := model.Library{
				Jar: library.(string),
			}
			libraryList = append(libraryList, thisLibrary)
		}
	}
	if eggs, ok := d.GetOk("library_egg"); ok {
		libraries := eggs.(*schema.Set).List()
		for _, library := range libraries {
			thisLibrary := model.Library{
				Egg: library.(string),
			}
			libraryList = append(libraryList, thisLibrary)
		}
	}
	if whls, ok := d.GetOk("library_whl"); ok {
		libraries := whls.(*schema.Set).List()
		for _, library := range libraries {
			thisLibrary := model.Library{
				Whl: library.(string),
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

func isJobMissing(errorMsg, resourceId string) bool {
	return strings.Contains(errorMsg, "INVALID_PARAMETER_VALUE") &&
		strings.Contains(errorMsg, fmt.Sprintf("Job %s does not exist.", resourceId))
}
