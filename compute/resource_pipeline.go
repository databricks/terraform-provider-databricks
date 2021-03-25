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
	Clusters            []Cluster         `json:"clusters,omitempty"`
	Libraries           []pipelineLibrary `json:"libraries,omitempty"`
	Filters             *filters          `json:"filters"`
	Continuous          bool              `json:"continuous,omitempty"`
	AllowDuplicateNames bool              `json:"allow_duplicate_names,omitempty"`
}

type createPipelineResponse struct {
	PipelineID string `json:"pipeline_id"`
}

type pipelineState string

const (
	statusDeploying  pipelineState = "DEPLOYING"
	statusStarting   pipelineState = "STARTING"
	statusRunning    pipelineState = "RUNNING"
	statusStopping   pipelineState = "STOPPPING"
	statusDeleted    pipelineState = "DELETED"
	statusRecovering pipelineState = "RECOVERING"
	statusFailed     pipelineState = "FAILED"
	statusResetting  pipelineState = "RESETTING"
	statusIdle       pipelineState = "IDLE"
)

type pipelineHealthStatus string

const (
	healthStatusHealthy   pipelineHealthStatus = "HEALTHY"
	healthStatusUnhealthy pipelineHealthStatus = "UNHEALTHY"
)

type pipelineInfo struct {
	PipelineID string                `json:"pipeline_id"`
	Spec       *pipelineSpec         `json:"spec"`
	State      *pipelineState        `json:"state"`
	Cause      string                `json:"cause"`
	ClusterID  string                `json:"cluster_id"`
	Name       string                `json:"name"`
	Health     *pipelineHealthStatus `json:"health"`
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

func (a pipelinesAPI) retryFunc(id string, desiredState pipelineState, lastState *pipelineState) *resource.RetryError {
	s, err := a.read(id)
	if err != nil {
		if e, ok := err.(common.APIError); ok && e.ErrorCode == "RESOURCE_DOES_NOT_EXIST" {
			*lastState = statusDeleted
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

func (a pipelinesAPI) waitForState(id string, desiredState pipelineState) (pipelineState, error) {
	var lastState pipelineState
	return lastState, resource.RetryContext(a.ctx, a.defaultTimeout(), func() *resource.RetryError { return a.retryFunc(id, desiredState, &lastState) })
}

func removeUnsupportedFields(m map[string]*schema.Schema) map[string]*schema.Schema {
	clusters, _ := m["clusters"].Elem.(*schema.Resource)
	clustersSchema := clusters.Schema
	delete(clustersSchema, "cluster_id")
	delete(clustersSchema, "cluster_name")
	delete(clustersSchema, "spark_version")
	delete(clustersSchema, "enable_elastic_disk")
	delete(clustersSchema, "enable_local_disk_encryption")
	delete(clustersSchema, "policy_id")
	delete(clustersSchema, "autotermination_minutes")
	delete(clustersSchema, "docker_image")
	delete(clustersSchema, "single_user_name")
	delete(clustersSchema, "idempotency_token")

	// Pipelines clusters are labeled
	clustersSchema["label"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	// num_workers is optional, but because the Cluster struct requires this field, we have to here as well.
	// Otherwise, we get the following error:
	// Error: clusters: Inconsistency: num_workers is optional, default is empty, but has no omitempty
	// clustersSchema["num_workers"].Required = false
	// clustersSchema["num_workers"].Optional = true
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
	var pipelineSchema = common.StructToSchema(pipelineSpec{}, removeUnsupportedFields)
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
			_, err = api.waitForState(id, statusRunning)
			d.SetId(string(id))
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
			_, err = api.waitForState(d.Id(), statusDeleted)
			return err
		},
	}.ToResource()
}
