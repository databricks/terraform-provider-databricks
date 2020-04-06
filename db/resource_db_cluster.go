package db

import (
	"bytes"
	"errors"
	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
)

func resourceCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceClusterCreate,
		Read:   resourceClusterRead,
		Update: resourceClusterUpdate,
		Delete: resourceClusterDelete,

		Schema: map[string]*schema.Schema{
			"num_workers": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"autoscale": &schema.Schema{
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
			"spark_version": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"spark_conf": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"aws_attributes": &schema.Schema{
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
			"driver_node_type_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"node_type_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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
			"init_scripts": &schema.Schema{
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
			"docker_image": &schema.Schema{
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
			"idempotency_token": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"cluster_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_tags": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
			},
		},
	}
}

func convertListInterfaceToString(m []interface{}) []string {
	response := []string{}
	for i, v := range m {
		if v != nil {
			response[i] = v.(string)
		}
	}
	return response
}

func resourceClusterCreate(d *schema.ResourceData, m interface{}) error {

	client := m.(service.DBApiClient)

	cluster := parseSchemaToCluster(d)

	clusterInfo, err := client.Clusters().Create(cluster)
	if err != nil {
		return err
	}

	err = client.Clusters().WaitForClusterRunning(clusterInfo.ClusterID, 30, 120)
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

	return resourceClusterRead(d, m)
}

func resourceClusterRead(d *schema.ResourceData, m interface{}) error {
	client := m.(service.DBApiClient)
	id := d.Id()

	clusterInfo, err := client.Clusters().Get(id)
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
		autoscale := map[string]string{}
		autoscale["min_workers"] = strconv.Itoa(int(clusterInfo.AutoScale.MinWorkers))
		autoscale["max_workers"] = strconv.Itoa(int(clusterInfo.AutoScale.MaxWorkers))
		err := d.Set("autoscale", autoscale)
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

	if len(clusterInfo.SparkConf) > 0 {
		err = d.Set("spark_conf", clusterInfo.SparkConf)
		if err != nil {
			return err
		}
	}

	if clusterInfo.AwsAttributes != nil {
		awsAtts := map[string]string{}
		if _, ok := d.GetOk("aws_attributes.availability"); ok {
			awsAtts["availability"] = string(clusterInfo.AwsAttributes.Availability)
		}
		if _, ok := d.GetOk("aws_attributes.zone_id"); ok {
			awsAtts["zone_id"] = clusterInfo.AwsAttributes.ZoneID
		}
		if _, ok := d.GetOk("aws_attributes.spot_bid_price_percent"); ok {
			awsAtts["spot_bid_price_percent"] = strconv.Itoa(int(clusterInfo.AwsAttributes.SpotBidPricePercent))
		}
		if _, ok := d.GetOk("aws_attributes.instance_profile_arn"); ok {
			awsAtts["instance_profile_arn"] = clusterInfo.AwsAttributes.InstanceProfileArn
		}
		if _, ok := d.GetOk("aws_attributes.first_on_demand"); ok {
			awsAtts["first_on_demand"] = strconv.Itoa(int(clusterInfo.AwsAttributes.FirstOnDemand))
		}
		if _, ok := d.GetOk("aws_attributes.ebs_volume_type"); ok {
			awsAtts["ebs_volume_type"] = string(clusterInfo.AwsAttributes.EbsVolumeType)
		}
		if _, ok := d.GetOk("aws_attributes.ebs_volume_count"); ok {
			awsAtts["ebs_volume_count"] = strconv.Itoa(int(clusterInfo.AwsAttributes.EbsVolumeCount))
		}
		if _, ok := d.GetOk("aws_attributes.ebs_volume_size"); ok {
			awsAtts["ebs_volume_size"] = strconv.Itoa(int(clusterInfo.AwsAttributes.EbsVolumeSize))
		}
		err = d.Set("aws_attributes", awsAtts)
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
		clusterLogConf := map[string]string{}
		if clusterInfo.ClusterLogConf.Dbfs != nil {
			clusterLogConf["dbfs_destination"] = clusterInfo.ClusterLogConf.Dbfs.Destination
		} else {
			if _, ok := d.GetOk("cluster_log_conf.s3_destination"); ok {
				clusterLogConf["s3_destination"] = clusterInfo.ClusterLogConf.S3.Destination
			}
			if _, ok := d.GetOk("cluster_log_conf.s3_region"); ok {
				clusterLogConf["s3_region"] = clusterInfo.ClusterLogConf.S3.Region
			}
			if _, ok := d.GetOk("cluster_log_conf.s3_endpoint"); ok {
				clusterLogConf["s3_endpoint"] = clusterInfo.ClusterLogConf.S3.Endpoint
			}
			if _, ok := d.GetOk("cluster_log_conf.s3_enable_encryption"); ok {
				clusterLogConf["s3_enable_encryption"] = strconv.FormatBool(clusterInfo.ClusterLogConf.S3.EnableEncryption)
			}
			if _, ok := d.GetOk("cluster_log_conf.s3_encryption_type"); ok {
				clusterLogConf["s3_encryption_type"] = clusterInfo.ClusterLogConf.S3.EncryptionType
			}
			if _, ok := d.GetOk("cluster_log_conf.s3_kms_key"); ok {
				clusterLogConf["s3_kms_key"] = clusterInfo.ClusterLogConf.S3.KmsKey
			}
			if _, ok := d.GetOk("cluster_log_conf.s3_canned_acl"); ok {
				clusterLogConf["s3_canned_acl"] = clusterInfo.ClusterLogConf.S3.CannedACL
			}
		}
		err = d.Set("cluster_log_conf", clusterLogConf)
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

	if clusterInfo.DockerImage != nil {
		dockerImage := map[string]string{}
		dockerImage["url"] = clusterInfo.DockerImage.Url
		if clusterInfo.DockerImage.BasicAuth != nil {
			dockerImage["username"] = clusterInfo.DockerImage.BasicAuth.Username
			dockerImage["password"] = clusterInfo.DockerImage.BasicAuth.Password
		}
		err = d.Set("aws_attributes", dockerImage)
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
	}

	return nil
}

func resourceClusterUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(service.DBApiClient)
	id := d.Id()
	clusterInfo, err := client.Clusters().Get(id)
	if err != nil {
		return err
	}
	clusterState := clusterInfo.State

	if model.ContainsClusterState([]model.ClusterState{model.ClusterState(model.ClusterStateTerminated)}, clusterState) {
		cluster := parseSchemaToCluster(d)
		cluster.ClusterId = id
		err := client.Clusters().Edit(cluster)
		if err != nil {
			return err
		}
		return resourceClusterRead(d, m)
	} else if model.ContainsClusterState([]model.ClusterState{model.ClusterState(model.ClusterStateRunning)}, clusterState) {
		cluster := parseSchemaToCluster(d)
		cluster.ClusterId = id
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
		cluster := parseSchemaToCluster(d)
		cluster.ClusterId = id
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

func parseSchemaToCluster(d *schema.ResourceData) model.Cluster {
	cluster := model.Cluster{}

	//Deal with Num workers
	if numWorkers, ok := d.GetOk("num_workers"); ok {
		cluster.NumWorkers = int32(numWorkers.(int))
	}

	//Deal with auto scaling options
	autoScale := model.AutoScale{}
	if autoscale, ok := d.GetOk("autoscale"); ok {
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
	}

	//Deal with cluster name
	if clusterName, ok := d.GetOk("cluster_name"); ok {
		cluster.ClusterName = clusterName.(string)
	}

	//Deal with spark versions
	if sparkVersion, ok := d.GetOk("spark_version"); ok {
		cluster.SparkVersion = sparkVersion.(string)
	}

	//Deal with spark confs
	if sparkConf, ok := d.GetOk("spark_conf"); ok {
		cluster.SparkConf = convertMapStringInterfaceToStringString(sparkConf.(map[string]interface{}))
	}

	//Deal with aws attributes for aws deployment
	awsAttributes := model.AwsAttributes{}

	if awsAttributesSchema, ok := d.GetOk("aws_attributes"); ok {
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
	if driverNodeTypeId, ok := d.GetOk("driver_node_type_id"); ok {
		cluster.DriverNodeTypeID = driverNodeTypeId.(string)
	}

	//Deal with worker node type id
	if nodeTypeId, ok := d.GetOk("node_type_id"); ok {
		cluster.NodeTypeID = nodeTypeId.(string)
	}

	//Deal with worker ssh public keys
	if sshPublicKeys, ok := d.GetOk("ssh_public_keys"); ok {
		cluster.SSHPublicKeys = convertListInterfaceToString(sshPublicKeys.(*schema.Set).List())
	}

	//Deal with worker custom tags
	if customTags, ok := d.GetOk("custom_tags"); ok {
		tags := customTags.(map[string]interface{})
		cluster.CustomTags = convertMapStringInterfaceToStringString(tags)
	}

	//Deal with worker cluster log config
	clusterLogConf := model.StorageInfo{}
	if clusterLogConfSchema, ok := d.GetOk("cluster_log_conf"); ok {
		clusterLogConfMap := clusterLogConfSchema.(map[string]interface{})
		if dbfsDestination, ok := clusterLogConfMap["dbfs_destination"]; ok {
			dbfsStorage := model.DbfsStorageInfo{}
			dbfsStorage.Destination = dbfsDestination.(string)
			clusterLogConf.Dbfs = &dbfsStorage
		} else {
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
	if initScripts, ok := d.GetOk("init_scripts"); ok {
		initScripts := initScripts.(*schema.Set).List()
		initScriptsLocations := []model.StorageInfo{}
		for _, v := range initScripts {
			initScript := v.(map[string]interface{})
			storageInfo := model.StorageInfo{}
			if dbfsDestination, ok := initScript["dbfs_destination"]; ok {
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
	if dockerImage, ok := d.GetOk("docker_image"); ok {
		dockerImageConf := dockerImage.(map[string]interface{})
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
	if sparkEnv, ok := d.GetOk("spark_env_vars"); ok {
		cluster.SparkEnvVars = convertMapStringInterfaceToStringString(sparkEnv.(map[string]interface{}))
	}

	//Deal with auto termination minutes
	if autoTerminationMinutes, ok := d.GetOk("autotermination_minutes"); ok {
		cluster.AutoterminationMinutes = int32(autoTerminationMinutes.(int))
	}

	//Deal with enable elastic disk
	if enableElasticDisk, ok := d.GetOk("enable_elastic_disk"); ok {
		cluster.EnableElasticDisk = enableElasticDisk.(bool)
	}

	//Deal with instance pool id
	if instancePoolID, ok := d.GetOk("instance_pool_id"); ok {
		cluster.InstancePoolId = instancePoolID.(string)
	}

	//Deal with idempotency token
	if idempotencyToken, ok := d.GetOk("idempotency_token"); ok {
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
