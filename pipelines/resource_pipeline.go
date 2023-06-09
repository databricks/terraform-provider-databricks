package pipelines

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/libraries"
)

// DefaultTimeout is the default amount of time that Terraform will wait when creating, updating and deleting pipelines.
const DefaultTimeout = 20 * time.Minute

// dltAutoScale is a struct the describes auto scaling for DLT clusters
type dltAutoScale struct {
	MinWorkers int32  `json:"min_workers,omitempty"`
	MaxWorkers int32  `json:"max_workers,omitempty"`
	Mode       string `json:"mode,omitempty"`
}

// We separate this struct from Cluster for two reasons:
// 1. Pipeline clusters include a `Label` field.
// 2. Spark version is not required (and shouldn't be specified) for pipeline clusters.
// 3. num_workers is optional, and there is no single-node support for pipelines clusters.
type pipelineCluster struct {
	Label string `json:"label,omitempty"` // used only by pipelines

	NumWorkers int32         `json:"num_workers,omitempty" tf:"group:size"`
	Autoscale  *dltAutoScale `json:"autoscale,omitempty" tf:"group:size"`

	NodeTypeID           string                    `json:"node_type_id,omitempty" tf:"group:node_type,computed"`
	DriverNodeTypeID     string                    `json:"driver_node_type_id,omitempty" tf:"computed"`
	InstancePoolID       string                    `json:"instance_pool_id,omitempty" tf:"group:node_type"`
	DriverInstancePoolID string                    `json:"driver_instance_pool_id,omitempty"`
	AwsAttributes        *clusters.AwsAttributes   `json:"aws_attributes,omitempty"`
	GcpAttributes        *clusters.GcpAttributes   `json:"gcp_attributes,omitempty"`
	AzureAttributes      *clusters.AzureAttributes `json:"azure_attributes,omitempty"`

	EnableLocalDiskEncryption bool `json:"enable_local_disk_encryption,omitempty" tf:"computed"`

	PolicyID                 string `json:"policy_id,omitempty"`
	ApplyPolicyDefaultValues bool   `json:"apply_policy_default_values,omitempty"`

	SparkConf    map[string]string `json:"spark_conf,omitempty"`
	SparkEnvVars map[string]string `json:"spark_env_vars,omitempty"`
	CustomTags   map[string]string `json:"custom_tags,omitempty"`

	SSHPublicKeys  []string                         `json:"ssh_public_keys,omitempty" tf:"max_items:10"`
	InitScripts    []clusters.InitScriptStorageInfo `json:"init_scripts,omitempty" tf:"max_items:10"` // TODO: tf:alias
	ClusterLogConf *clusters.StorageInfo            `json:"cluster_log_conf,omitempty"`
}

type NotebookLibrary struct {
	Path string `json:"path"`
}

type FileLibrary struct {
	Path string `json:"path"`
}

type PipelineLibrary struct {
	Jar      string           `json:"jar,omitempty"`
	Maven    *libraries.Maven `json:"maven,omitempty"`
	Whl      string           `json:"whl,omitempty"`
	Notebook *NotebookLibrary `json:"notebook,omitempty"`
	File     *FileLibrary     `json:"file,omitempty"`
}

type filters struct {
	Include []string `json:"include,omitempty"`
	Exclude []string `json:"exclude,omitempty"`
}

type Notification struct {
	EmailRecipients []string `json:"email_recipients" tf:"min_items:1"`
	Alerts          []string `json:"alerts" tf:"min_items:1"`
}

type PipelineSpec struct {
	ID                  string            `json:"id,omitempty" tf:"computed"`
	Name                string            `json:"name,omitempty"`
	Storage             string            `json:"storage,omitempty" tf:"force_new"`
	Catalog             string            `json:"catalog,omitempty" tf:"force_new"`
	Configuration       map[string]string `json:"configuration,omitempty"`
	Clusters            []pipelineCluster `json:"clusters,omitempty" tf:"alias:cluster"`
	Libraries           []PipelineLibrary `json:"libraries,omitempty" tf:"slice_set,alias:library"`
	Filters             *filters          `json:"filters,omitempty"`
	Continuous          bool              `json:"continuous,omitempty"`
	Development         bool              `json:"development,omitempty"`
	AllowDuplicateNames bool              `json:"allow_duplicate_names,omitempty"`
	Target              string            `json:"target,omitempty"`
	Photon              bool              `json:"photon,omitempty"`
	Edition             string            `json:"edition,omitempty" tf:"suppress_diff,default:ADVANCED"`
	Channel             string            `json:"channel,omitempty" tf:"suppress_diff,default:CURRENT"`
	Notifications       []Notification    `json:"notifications,omitempty" tf:"alias:notification"`
	Serverless          bool              `json:"serverless" tf:"optional"`
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

type PipelineInfo struct {
	PipelineID      string                `json:"pipeline_id"`
	Spec            *PipelineSpec         `json:"spec"`
	State           *PipelineState        `json:"state"`
	Cause           string                `json:"cause"`
	ClusterID       string                `json:"cluster_id"`
	Name            string                `json:"name"`
	Health          *PipelineHealthStatus `json:"health"`
	CreatorUserName string                `json:"creator_user_name"`
}

type PipelineUpdateStateInfo struct {
	UpdateID     string         `json:"update_id"`
	State        *PipelineState `json:"state"`
	CreationTime string         `json:"creation_time"`
}

type PipelineStateInfo struct {
	PipelineID      string                    `json:"pipeline_id"`
	State           *PipelineState            `json:"state"`
	ClusterID       string                    `json:"cluster_id"`
	Name            string                    `json:"name"`
	Health          *PipelineHealthStatus     `json:"health"`
	CreatorUserName string                    `json:"creator_user_name"`
	RunAsUserName   string                    `json:"run_as_user_name"`
	LatestUpdates   []PipelineUpdateStateInfo `json:"latest_updates,omitempty"`
}

type PipelineListResponse struct {
	Statuses      []PipelineStateInfo `json:"statuses"`
	NextPageToken string              `json:"next_page_token,omitempty"`
	PrevPageToken string              `json:"prev_page_token,omitempty"`
}

type PipelinesAPI struct {
	client *common.DatabricksClient
	ctx    context.Context
}

func NewPipelinesAPI(ctx context.Context, m any) PipelinesAPI {
	return PipelinesAPI{m.(*common.DatabricksClient), ctx}
}

func (a PipelinesAPI) Create(s PipelineSpec, timeout time.Duration) (string, error) {
	var resp createPipelineResponse
	err := a.client.Post(a.ctx, "/pipelines", s, &resp)
	if err != nil {
		return "", err
	}
	id := resp.PipelineID
	err = a.waitForState(id, timeout, StateRunning)
	if err != nil {
		log.Printf("[INFO] Pipeline creation failed, attempting to clean up pipeline %s", id)
		err2 := a.Delete(id, timeout)
		if err2 != nil {
			log.Printf("[WARN] Unable to delete pipeline %s; this resource needs to be manually cleaned up", id)
			return "", fmt.Errorf("multiple errors occurred when creating pipeline. Error while waiting for creation: \"%v\"; error while attempting to clean up failed pipeline: \"%v\"", err, err2)
		}
		log.Printf("[INFO] Successfully cleaned up pipeline %s", id)
		return "", err
	}
	return id, nil
}

func (a PipelinesAPI) Read(id string) (p PipelineInfo, err error) {
	err = a.client.Get(a.ctx, "/pipelines/"+id, nil, &p)
	return
}

func (a PipelinesAPI) Update(id string, s PipelineSpec, timeout time.Duration) error {
	err := a.client.Put(a.ctx, "/pipelines/"+id, s)
	if err != nil {
		return err
	}
	return a.waitForState(id, timeout, StateRunning)
}

func (a PipelinesAPI) Delete(id string, timeout time.Duration) error {
	err := a.client.Delete(a.ctx, "/pipelines/"+id, map[string]string{})
	if err != nil {
		return err
	}
	return resource.RetryContext(a.ctx, timeout,
		func() *resource.RetryError {
			i, err := a.Read(id)
			if err != nil {
				if apierr.IsMissing(err) {
					return nil
				}
				return resource.NonRetryableError(err)
			}
			message := fmt.Sprintf("Pipeline %s is in state %s, not yet deleted", id, *i.State)
			log.Printf("[DEBUG] %s", message)
			return resource.RetryableError(fmt.Errorf(message))
		})
}

// List returns a list of the DLT pipelines. List could be filtered by name
func (a PipelinesAPI) List(pageSize int, filter string) ([]PipelineStateInfo, error) {
	payload := map[string]any{"max_results": pageSize}
	if filter != "" {
		payload["filter"] = filter
	}
	result := []PipelineStateInfo{}

	for {
		var resp PipelineListResponse
		err := a.client.Get(a.ctx, "/pipelines", payload, &resp)
		if err != nil {
			return []PipelineStateInfo{}, err
		}
		result = append(result, resp.Statuses...)
		if resp.NextPageToken == "" {
			break
		}
		payload["page_token"] = resp.NextPageToken
	}

	return result, nil
}

func (a PipelinesAPI) waitForState(id string, timeout time.Duration, desiredState PipelineState) error {
	return resource.RetryContext(a.ctx, timeout,
		func() *resource.RetryError {
			i, err := a.Read(id)
			if err != nil {
				return resource.NonRetryableError(err)
			}
			state := *i.State
			if state == desiredState {
				return nil
			}
			if state == StateFailed {
				return resource.NonRetryableError(fmt.Errorf("pipeline %s has failed", id))
			}
			if !i.Spec.Continuous {
				// continuous pipelines just need a non-FAILED check
				return nil
			}
			message := fmt.Sprintf("Pipeline %s is in state %s, not yet in state %s", id, state, desiredState)
			log.Printf("[DEBUG] %s", message)
			return resource.RetryableError(fmt.Errorf(message))
		})
}

func suppressStorageDiff(k, old, new string, d *schema.ResourceData) bool {
	defaultStorageRegex := regexp.MustCompile(
		`^dbfs:/pipelines/[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)
	res := defaultStorageRegex.MatchString(old)
	if new == "" && res {
		log.Printf("[DEBUG] Suppressing diff for %v: platform=%#v config=%#v", k, old, new)
		return true
	}
	return false
}

func adjustPipelineResourceSchema(m map[string]*schema.Schema) map[string]*schema.Schema {
	cluster, _ := m["cluster"].Elem.(*schema.Resource)
	clustersSchema := cluster.Schema
	clustersSchema["spark_conf"].DiffSuppressFunc = clusters.SparkConfDiffSuppressFunc
	common.MustSchemaPath(clustersSchema,
		"aws_attributes", "zone_id").DiffSuppressFunc = clusters.ZoneDiffSuppress
	common.MustSchemaPath(clustersSchema, "autoscale", "mode").DiffSuppressFunc = common.EqualFoldDiffSuppress

	gcpAttributes, _ := clustersSchema["gcp_attributes"].Elem.(*schema.Resource)
	gcpAttributesSchema := gcpAttributes.Schema
	delete(gcpAttributesSchema, "use_preemptible_executors")
	delete(gcpAttributesSchema, "boot_disk_size")

	m["library"].MinItems = 1
	m["url"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
	m["channel"].ValidateFunc = validation.StringInSlice([]string{"current", "preview"}, true)
	m["edition"].ValidateFunc = validation.StringInSlice([]string{"pro", "core", "advanced"}, true)
	m["edition"].DiffSuppressFunc = common.EqualFoldDiffSuppress

	m["storage"].DiffSuppressFunc = suppressStorageDiff
	m["storage"].ConflictsWith = []string{"catalog"}
	m["catalog"].ConflictsWith = []string{"storage"}

	return m
}

// ResourcePipeline defines the Terraform resource for pipelines.
func ResourcePipeline() *schema.Resource {
	var pipelineSchema = common.StructToSchema(PipelineSpec{}, adjustPipelineResourceSchema)
	return common.Resource{
		Schema: pipelineSchema,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var s PipelineSpec
			common.DataToStructPointer(d, pipelineSchema, &s)
			api := NewPipelinesAPI(ctx, c)
			id, err := api.Create(s, d.Timeout(schema.TimeoutCreate))
			if err != nil {
				return err
			}
			d.SetId(id)
			d.Set("url", c.FormatURL("#joblist/pipelines/", d.Id()))
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			i, err := NewPipelinesAPI(ctx, c).Read(d.Id())
			if err != nil {
				return err
			}
			if i.Spec == nil {
				return fmt.Errorf("pipeline spec is nil for '%v'", i.PipelineID)
			}
			return common.StructToData(*i.Spec, pipelineSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var s PipelineSpec
			common.DataToStructPointer(d, pipelineSchema, &s)
			return NewPipelinesAPI(ctx, c).Update(d.Id(), s, d.Timeout(schema.TimeoutUpdate))
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			api := NewPipelinesAPI(ctx, c)
			return api.Delete(d.Id(), d.Timeout(schema.TimeoutDelete))
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(DefaultTimeout),
		},
	}.ToResource()
}
