package compute

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/databrickslabs/terraform-provider-databricks/common"
)

// We separate this struct from Cluster for two reasons:
// 1. Pipeline clusters include a `Label` field.
// 2. Spark version is not required (and shouldn't be specified) for pipeline clusters.
// 3. num_workers is optional, and there is no single-node support for pipelines clusters.
type pipelineCluster struct {
	Label string `json:"label,omitempty"` // used only by pipelines

	NumWorkers int32      `json:"num_workers,omitempty" tf:"group:size"`
	Autoscale  *AutoScale `json:"autoscale,omitempty" tf:"group:size"`

	NodeTypeID       string         `json:"node_type_id,omitempty" tf:"group:node_type,computed"`
	DriverNodeTypeID string         `json:"driver_node_type_id,omitempty" tf:"conflicts:instance_pool_id,computed"`
	InstancePoolID   string         `json:"instance_pool_id,omitempty" tf:"group:node_type"`
	AwsAttributes    *AwsAttributes `json:"aws_attributes,omitempty" tf:"conflicts:instance_pool_id"`

	SparkConf    map[string]string `json:"spark_conf,omitempty"`
	SparkEnvVars map[string]string `json:"spark_env_vars,omitempty"`
	CustomTags   map[string]string `json:"custom_tags,omitempty"`

	SSHPublicKeys  []string                `json:"ssh_public_keys,omitempty" tf:"max_items:10"`
	InitScripts    []InitScriptStorageInfo `json:"init_scripts,omitempty" tf:"max_items:10"` // TODO: tf:alias
	ClusterLogConf *StorageInfo            `json:"cluster_log_conf,omitempty"`
}

type notebookLibrary struct {
	Path string `json:"path"`
}

type pipelineLibrary struct {
	Jar      string           `json:"jar,omitempty"`
	Maven    *Maven           `json:"maven,omitempty"`
	Whl      string           `json:"whl,omitempty"`
	Notebook *notebookLibrary `json:"notebook,omitempty"`
}

type filters struct {
	Include []string `json:"include,omitempty"`
	Exclude []string `json:"exclude,omitempty"`
}

type pipelineSpec struct {
	ID                  string            `json:"id,omitempty" tf:"computed"`
	Name                string            `json:"name,omitempty"`
	Storage             string            `json:"storage,omitempty"`
	Configuration       map[string]string `json:"configuration,omitempty"`
	Clusters            []pipelineCluster `json:"clusters,omitempty"`
	Libraries           []pipelineLibrary `json:"libraries,omitempty"`
	Filters             *filters          `json:"filters"`
	Continuous          bool              `json:"continuous,omitempty"`
	AllowDuplicateNames bool              `json:"allow_duplicate_names,omitempty"`
}

type createPipelineResponse struct {
	PipelineID string `json:"pipeline_id"`
}

// PipelineState ...
type PipelineState string

// Constants for PipelineStates
const (
	StateDeploying  PipelineState = "DEPLOYING"
	StateStarting   PipelineState = "STARTING"
	StateRunning    PipelineState = "RUNNING"
	StateStopping   PipelineState = "STOPPPING"
	StateDeleted    PipelineState = "DELETED"
	StateRecovering PipelineState = "RECOVERING"
	StateFailed     PipelineState = "FAILED"
	StateResetting  PipelineState = "RESETTING"
	StateIdle       PipelineState = "IDLE"
)

// PipelineHealthStatus ...
type PipelineHealthStatus string

// Constants for PipelineHealthStatus
const (
	HealthStatusHealthy   PipelineHealthStatus = "HEALTHY"
	HealthStatusUnhealthy PipelineHealthStatus = "UNHEALTHY"
)

type pipelineInfo struct {
	PipelineID string                `json:"pipeline_id"`
	Spec       *pipelineSpec         `json:"spec"`
	State      *PipelineState        `json:"state"`
	Cause      string                `json:"cause"`
	ClusterID  string                `json:"cluster_id"`
	Name       string                `json:"name"`
	Health     *PipelineHealthStatus `json:"health"`
}

type pipelinesAPI struct {
	client *common.DatabricksClient
	ctx    context.Context
}

func newPipelinesAPI(ctx context.Context, m interface{}) pipelinesAPI {
	return pipelinesAPI{m.(*common.DatabricksClient), ctx}
}

func (a pipelinesAPI) create(s pipelineSpec, allowDuplicateNames bool) (string, error) {
	var id createPipelineResponse
	err := a.client.Post(a.ctx, "/pipelines", s, &id)
	return id.PipelineID, err
}

func (a pipelinesAPI) read(id string) (p pipelineInfo, err error) {
	err = a.client.Get(a.ctx, "/pipelines/"+id, nil, &p)
	return
}

func (a pipelinesAPI) update(id string, s pipelineSpec, allowDuplicateNames bool) error {
	return a.client.Put(a.ctx, "/pipelines/"+id, s)
}

func (a pipelinesAPI) delete(id string) error {
	return a.client.Delete(a.ctx, "/pipelines/"+id, map[string]string{})
}

func (a pipelinesAPI) defaultTimeout() time.Duration {
	return 30 * time.Minute
}

func (a pipelinesAPI) retryFunc(id string, desiredState PipelineState, lastState *PipelineState) *resource.RetryError {
	s, err := a.read(id)
	if err != nil {
		if e, ok := err.(common.APIError); ok && e.ErrorCode == "RESOURCE_DOES_NOT_EXIST" {
			*lastState = StateDeleted
		} else {
			return resource.NonRetryableError(err)
		}
	} else {
		lastState = s.State
	}
	log.Printf("[DEBUG] Pipeline %s is in state %s", id, *lastState)
	if *lastState == desiredState {
		return nil
	}
	return resource.RetryableError(fmt.Errorf("Pipeline %s is in state %s, not yet in state %s", id, *lastState, desiredState))
}

func (a pipelinesAPI) waitForState(id string, desiredState PipelineState) (PipelineState, error) {
	var lastState PipelineState
	return lastState, resource.RetryContext(a.ctx, a.defaultTimeout(), func() *resource.RetryError { return a.retryFunc(id, desiredState, &lastState) })
}

func adjustPipelineResourceSchema(m map[string]*schema.Schema) map[string]*schema.Schema {
	clusters, _ := m["clusters"].Elem.(*schema.Resource)
	clustersSchema := clusters.Schema
	clustersSchema["spark_conf"].DiffSuppressFunc = sparkConfDiffSuppressFunc

	awsAttributes, _ := clustersSchema["aws_attributes"].Elem.(*schema.Resource)
	awsAttributesSchema := awsAttributes.Schema
	delete(awsAttributesSchema, "first_on_demand")
	delete(awsAttributesSchema, "availability")
	delete(awsAttributesSchema, "spot_bid_price_percent")
	delete(awsAttributesSchema, "ebs_volume_type")
	delete(awsAttributesSchema, "ebs_volume_count")
	delete(awsAttributesSchema, "ebs_volume_size")

	m["libraries"].MinItems = 1

	return m
}

// ResourcePipeline ...
func ResourcePipeline() *schema.Resource {
	var pipelineSchema = common.StructToSchema(pipelineSpec{}, adjustPipelineResourceSchema)
	return common.Resource{
		Schema: pipelineSchema,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var s pipelineSpec
			err := common.DataToStructPointer(d, pipelineSchema, &s)
			if err != nil {
				return err
			}
			api := newPipelinesAPI(ctx, c)
			id, err := api.create(s, false)
			if err != nil {
				return err
			}
			_, err = api.waitForState(id, StateRunning)
			if err != nil {
				return err
			}
			d.SetId(id)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			i, err := newPipelinesAPI(ctx, c).read(d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(*i.Spec, pipelineSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var s pipelineSpec
			err := common.DataToStructPointer(d, pipelineSchema, &s)
			if err != nil {
				return err
			}
			return newPipelinesAPI(ctx, c).update(d.Id(), s, false)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			api := newPipelinesAPI(ctx, c)
			err := api.delete(d.Id())
			if err != nil {
				return err
			}
			_, err = api.waitForState(d.Id(), StateDeleted)
			return err
		},
	}.ToResource()
}
