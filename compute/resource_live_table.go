package compute

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/databrickslabs/terraform-provider-databricks/common"
)

// Live Tables Models

type PipelinesAwsAttributes struct {
	ZoneId             string `json:"zone_id,omitempty" tf:"optional"`
	InstanceProfileArn string `json:"instance_profile_arn,omitempty" tf:"optional"`
}

type PipelinesDbfsStorageInfo struct {
	Destination string `json:"destination,omitempty" tf:"optional"`
}

type PipelinesLocalFileInfo struct {
	Destination string `json:"destination,omitempty" tf:"optional"`
}

type PipelinesS3StorageInfo struct {
	Destination      string `json:"destination,omitempty" tf:"optional"`
	Region           string `json:"region,omitempty" tf:"optional"`
	Endpoint         string `json:"endpoint,omitempty" tf:"optional"`
	EnableEncryption bool   `json:"enable_encryption,omitempty" tf:"optional,default:false"`
	EncryptionType   string `json:"encryption_type,omitempty" tf:"optional"`
	KmsKey           string `json:"kms_key,omitempty" tf:"optional"`
	CannedAcl        string `jsono:"canned_acl,omitempty" tf:"optional"`
}

type PipelinesClusterLogConf struct {
	Dbfs *PipelinesDbfsStorageInfo `json:"dbfs,omitempty" tf:"optional"`
	S3   *PipelinesS3StorageInfo   `json:"s3,omitempty" tf:"optional"`
}

type PipelinesInitScriptInfo struct {
	Dbfs *PipelinesDbfsStorageInfo `json:"dbfs,omitempty" tf:"optional"`
	S3   *PipelinesS3StorageInfo   `json:"s3,omitempty" tf:"optional"`
	File *PipelinesLocalFileInfo   `json:"file,omitempty" tf:"optional"`
}

type PipelinesClusterAttributes struct {
	SparkConf        map[string]string        `json:"spark_conf,omitempty"`
	AwsAttributes    *PipelinesAwsAttributes  `json:"aws_attributes,omitempty"`
	NodeTypeId       string                   `json:"node_type_id,omitempty"`
	DriverNodeTypeId string                   `json:"driver_node_type_id,omitempty"`
	SshPublicKeys    string                   `json:"ssh_public_keys,omitempty"`
	CustomTags       map[string]string        `json:"custom_tags,omitempty"`
	ClusterLogConf   *PipelinesClusterLogConf `json:"cluster_log_conf,omitempty"`
	SparkEnvVars     map[string]string        `json:"spark_env_vars,omitempty"`
	InstancePoolId   string                   `json:"instance_pool_id,omitempty"`
}

type PipelinesAutoScale struct {
	MinWorkers int `json:"min_workers"`
	MaxWorkers int `json:"max_workers"`
}

type PipelinesClusterSize struct {
	NumWorkers int                 `json:"num_workers,omitempty"`
	Autoscale  *PipelinesAutoScale `json:"autoscale,omitempty"`
}

type PipelinesNewCluster struct {
	PipelinesClusterAttributes
	PipelinesClusterSize
}

type PipelineCluster struct {
	Label string `json:"label,omitempty"`
	PipelinesNewCluster
}

type PipelinesMavenLibrary struct {
	Coordinates string   `json:"coordinates"`
	Repo        string   `json:"repo,omitempty"`
	Exclusions  []string `json:"exclusions,omitempty"`
}

type NotebookLibrary struct {
	Path string `json:"path"`
}

type PipelineLibrary struct {
	Jar      string                 `json:"jar,omitempty"`
	Maven    *PipelinesMavenLibrary `json:"maven,omitempty"`
	Whl      string                 `json:"whl,omitempty"`
	Notebook *NotebookLibrary       `json:"notebook,omitempty"`
}

type Filters struct {
	Include []string `json:"include,omitempty"`
	Exclude []string `json:"exclude,omitempty"`
}

type PipelineSpec struct {
	Id            string            `json:"id,omitempty"`
	Name          string            `json:"name,omitempty"`
	Storage       string            `json:"storage,omitempty"`
	Configuration map[string]string `json:"configuration,omitempty"`
	Clusters      []PipelineCluster `json:"clusters,omitempty"`
	Libraries     []PipelineLibrary `json:"libraries,omitempty"`
	Filters       []Filters         `json:"filters"`
	Continuous    bool              `json:"continuous,omitempty"`
}

type PipelineState string

const (
	StatusDeploying  PipelineState = "DEPLOYING"
	StatusStarting   PipelineState = "STARTING"
	StatusRunning    PipelineState = "RUNNING"
	StatusStopping   PipelineState = "STOPPPING"
	StatusDeleted    PipelineState = "DELETED"
	StatusRecovering PipelineState = "RECOVERING"
	StatusFailed     PipelineState = "FAILED"
	StatusResetting  PipelineState = "RESETTING"
	StatusIdle       PipelineState = "IDLE"
)

type PipelineHealthStatus string

const (
	HealthStatusHealthy   PipelineHealthStatus = "HEALTHY"
	HealthStatusUnhealthy PipelineHealthStatus = "UNHEALTHY"
)

type PipelineInfo struct {
	PipelineId string                `json:"pipeline_id"`
	Spec       *PipelineSpec         `json:"spec"`
	State      *PipelineState        `json:"state"`
	Cause      string                `json:"cause"`
	ClusterId  string                `json:"cluster_id"`
	Name       string                `json:"name"`
	Health     *PipelineHealthStatus `json:"health"`
}

// Live Tables API Client

type LiveTablesAPI struct {
	client *common.DatabricksClient
	ctx    context.Context
}

type LiveTableId string

func NewLiveTablesAPI(ctx context.Context, m interface{}) LiveTablesAPI {
	return LiveTablesAPI{m.(*common.DatabricksClient), ctx}
}

func (a LiveTablesAPI) Create(s PipelineSpec, allowDuplicateNames bool) (LiveTableId, error) {
	var id LiveTableId
	err := a.client.Post(a.ctx, "/pipelines", map[string]interface{}{
		"spec":                  s,
		"allow_duplicate_names": allowDuplicateNames,
	}, &id)
	return id, err
}

func (a LiveTablesAPI) Read(id LiveTableId) (p PipelineInfo, err error) {
	err = a.client.Get(a.ctx, "/pipelines/"+string(id), nil, &p)
	return
}

func (a LiveTablesAPI) Update(id LiveTableId, s PipelineSpec, allowDuplicateNames bool) error {
	return a.client.Put(a.ctx, "/pipelines/"+string(id), map[string]interface{}{
		"spec":                  s,
		"allow_duplicate_names": allowDuplicateNames,
	})
}

func (a LiveTablesAPI) Delete(id LiveTableId) error {
	return a.client.Delete(a.ctx, "/pipelines/"+string(id), nil)
}

// Live Tables Resource Definition

var pipelineSchema = common.StructToSchema(PipelineSpec{}, func(m map[string]*schema.Schema) map[string]*schema.Schema { return m })

func ResourceLiveTable() *schema.Resource {
	return common.Resource{
		Schema:        pipelineSchema,
		SchemaVersion: 2,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var s PipelineSpec
			err := common.DataToStructPointer(d, pipelineSchema, &s)
			if err != nil {
				return err
			}
			id, err := NewLiveTablesAPI(ctx, c).Create(s, false)
			if err != nil {
				return err
			}
			d.SetId(string(id))
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			i, err := NewLiveTablesAPI(ctx, c).Read(LiveTableId(d.Id()))
			if err != nil {
				return err
			}
			return common.StructToData(i.Spec, pipelineSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var s PipelineSpec
			err := common.DataToStructPointer(d, pipelineSchema, &s)
			if err != nil {
				return err
			}
			return NewLiveTablesAPI(ctx, c).Update(LiveTableId(d.Id()), s, false)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewLiveTablesAPI(ctx, c).Delete(LiveTableId(d.Id()))
		},
	}.ToResource()
}
