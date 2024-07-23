package pipelines

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"time"

	"reflect"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/pipelines"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func PrintStruct(s interface{}) {
	v := reflect.ValueOf(s)
	t := v.Type()

	fmt.Printf("Struct: %s\n", t.Name())
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		fmt.Printf("%s: %v\n", field.Name, value)
	}
}

// DefaultTimeout is the default amount of time that Terraform will wait when creating, updating and deleting pipelines.
const DefaultTimeout = 20 * time.Minute

func adjustForceSendFields(clusterList *[]pipelines.PipelineCluster) {
	for i := range *clusterList {
		cluster := &((*clusterList)[i])
		// TF Go SDK doesn't differentiate between the default and not set values.
		// If nothing is specified, DLT creates a cluster with enhanced autoscaling
		// from 1 to 5 nodes, which is different than sending a request for zero workers.
		// The solution here is to look for the Spark configuration to determine
		// if the user only wants a single node cluster (only master, no workers).
		if cluster.SparkConf["spark.databricks.cluster.profile"] == "singleNode" {
			cluster.ForceSendFields = append(cluster.ForceSendFields, "NumWorkers")
		}
	}
}

func Create(w *databricks.WorkspaceClient, ctx context.Context, d *schema.ResourceData, timeout time.Duration) error {
	var createPipelineRequest createPipelineRequestStruct
	common.DataToStructPointer(d, pipelineSchema, &createPipelineRequest)
	PrintStruct(createPipelineRequest.CreatePipeline)
	adjustForceSendFields(&createPipelineRequest.Clusters)

	createdPipeline, err := w.Pipelines.Create(ctx, createPipelineRequest.CreatePipeline)
	if err != nil {
		return err
	}
	var id string
	if d.Get("dry_run").(bool) {
		id = createdPipeline.EffectiveSettings.Id
		d.SetId(id)
		return nil
	} else {
		id = createdPipeline.PipelineId
	}
	// _, err = w.Pipelines.WaitGetPipelineRunning(ctx, id, timeout, nil)
	err = waitForState(w, ctx, id, timeout, pipelines.PipelineStateRunning)
	if err != nil {
		log.Printf("[INFO] Pipeline creation failed, attempting to clean up pipeline %s", id)
		err2 := Delete(w, ctx, id, timeout)
		if err2 != nil {
			log.Printf("[WARN] Unable to delete pipeline %s; this resource needs to be manually cleaned up", id)
			return fmt.Errorf("multiple errors occurred when creating pipeline. Error while waiting for creation: \"%v\"; error while attempting to clean up failed pipeline: \"%v\"", err, err2)
		}
		log.Printf("[INFO] Successfully cleaned up pipeline %s", id)
		return err
	}
	d.SetId(id)
	d.Set("pipeline_id", id)
	return nil
}

func Read(w *databricks.WorkspaceClient, ctx context.Context, id string) (*pipelines.GetPipelineResponse, error) {
	return w.Pipelines.Get(ctx, pipelines.GetPipelineRequest{
		PipelineId: id,
	})
}

func Update(w *databricks.WorkspaceClient, ctx context.Context, d *schema.ResourceData, timeout time.Duration) error {
	var updatePipelineRequest updatePipelineRequestStruct
	fmt.Println("ExpectedLastModified", d.Get("expected_last_modified"))
	// fmt.Println(d.Get("expected_last_modified"))
	common.DataToStructPointer(d, pipelineSchema, &updatePipelineRequest)
	PrintStruct(updatePipelineRequest.EditPipeline)
	// fmt.Println("updatePipelineRequest", updatePipelineRequest)
	adjustForceSendFields(&updatePipelineRequest.Clusters)

	err := w.Pipelines.Update(ctx, updatePipelineRequest.EditPipeline)
	if err != nil {
		return err
	}
	return waitForState(w, ctx, d.Id(), timeout, pipelines.PipelineStateRunning)
}

func Delete(w *databricks.WorkspaceClient, ctx context.Context, id string, timeout time.Duration) error {
	err := w.Pipelines.Delete(ctx, pipelines.DeletePipelineRequest{
		PipelineId: id,
	})
	if err != nil {
		return err
	}
	return retry.RetryContext(ctx, timeout,
		func() *retry.RetryError {
			i, err := Read(w, ctx, id)
			if err != nil {
				if apierr.IsMissing(err) {
					return nil
				}
				return retry.NonRetryableError(err)
			}
			message := fmt.Sprintf("Pipeline %s is in state %s, not yet deleted", id, i.State)
			log.Printf("[DEBUG] %s", message)
			return retry.RetryableError(fmt.Errorf(message))
		})
}

func waitForState(w *databricks.WorkspaceClient, ctx context.Context, id string, timeout time.Duration, desiredState pipelines.PipelineState) error {
	return retry.RetryContext(ctx, timeout,
		func() *retry.RetryError {
			i, err := Read(w, ctx, id)
			if err != nil {
				return retry.NonRetryableError(err)
			}
			state := i.State
			if state == desiredState {
				return nil
			}
			if state == pipelines.PipelineStateFailed {
				return retry.NonRetryableError(fmt.Errorf("pipeline %s has failed", id))
			}
			if !i.Spec.Continuous {
				// continuous pipelines just need a non-FAILED check
				return nil
			}
			message := fmt.Sprintf("Pipeline %s is in state %s, not yet in state %s", id, state, desiredState)
			log.Printf("[DEBUG] %s", message)
			return retry.RetryableError(fmt.Errorf(message))
		})
}

type createPipelineRequestStruct struct {
	pipelines.CreatePipeline
}

var aliasMap = map[string]string{
	"clusters":      "cluster",
	"libraries":     "library",
	"notifications": "notification",
}

func (createPipelineRequestStruct) Aliases() map[string]map[string]string {
	return map[string]map[string]string{
		"pipelines.createPipelineRequestStruct": aliasMap,
	}
}

func (createPipelineRequestStruct) CustomizeSchema(s *common.CustomizableSchema) *common.CustomizableSchema {
	return s
}

type updatePipelineRequestStruct struct {
	pipelines.EditPipeline
}

func (updatePipelineRequestStruct) Aliases() map[string]map[string]string {
	return map[string]map[string]string{
		"pipelines.updatePipelineRequestStruct": aliasMap,
	}
}

func (updatePipelineRequestStruct) CustomizeSchema(s *common.CustomizableSchema) *common.CustomizableSchema {
	return s
}

// Name not keeping twice
type Pipeline struct {
	pipelines.PipelineSpec
	PipelineId           string                              `json:"pipeline_id,omitempty"`
	AllowDuplicateNames  bool                                `json:"allow_duplicate_names,omitempty"`
	DryRun               bool                                `json:"dry_run,omitempty"`
	Cause                string                              `json:"cause,omitempty"`
	ClusterId            string                              `json:"cluster_id,omitempty"`
	CreatorUserName      string                              `json:"creator_user_name,omitempty"`
	Health               pipelines.GetPipelineResponseHealth `json:"health,omitempty"`
	LastModified         int64                               `json:"last_modified,omitempty"`
	LatestUpdates        []pipelines.UpdateStateInfo         `json:"latest_updates,omitempty"`
	RunAsUserName        string                              `json:"run_as_user_name,omitempty"`
	ExpectedLastModified int64                               `json:"expected_last_modified,omitempty"`
	State                pipelines.PipelineState             `json:"state,omitempty"`
	// Provides the URL to the pipeline in the Databricks UI.
	URL string `json:"url,omitempty"`
}

func (Pipeline) Aliases() map[string]map[string]string {
	return map[string]map[string]string{
		"pipelines.Pipeline": aliasMap,
	}

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

func (Pipeline) CustomizeSchema(s *common.CustomizableSchema) *common.CustomizableSchema {

	// ForceNew fields
	s.SchemaPath("storage").SetForceNew()
	s.SchemaPath("catalog").SetForceNew()

	// Computed fields
	s.SchemaPath("id").SetComputed()
	s.SchemaPath("pipeline_id").SetComputed()
	s.SchemaPath("cluster", "node_type_id").SetComputed()
	s.SchemaPath("cluster", "driver_node_type_id").SetComputed()
	// s.SchemaPath("cluster", "enable_local_disk_encryption").SetComputed()
	s.SchemaPath("url").SetComputed()
	// Manually set the computed fields
	s.SchemaPath("state").SetComputed()
	s.SchemaPath("latest_updates").SetComputed()
	s.SchemaPath("last_modified").SetComputed()
	s.SchemaPath("health").SetComputed()
	s.SchemaPath("cause").SetComputed()
	s.SchemaPath("cluster_id").SetComputed()
	s.SchemaPath("creator_user_name").SetComputed()

	// SuppressDiff fields
	s.SchemaPath("edition").SetSuppressDiff()
	s.SchemaPath("channel").SetSuppressDiff()
	s.SchemaPath("cluster", "spark_conf").SetCustomSuppressDiff(clusters.SparkConfDiffSuppressFunc)
	s.SchemaPath("cluster", "aws_attributes", "zone_id").SetCustomSuppressDiff(clusters.ZoneDiffSuppress)
	s.SchemaPath("cluster", "autoscale", "mode").SetCustomSuppressDiff(common.EqualFoldDiffSuppress)
	s.SchemaPath("edition").SetCustomSuppressDiff(common.EqualFoldDiffSuppress)
	s.SchemaPath("storage").SetCustomSuppressDiff(suppressStorageDiff)

	// Deprecated fields
	s.SchemaPath("cluster", "init_scripts", "dbfs").SetDeprecated(clusters.DbfsDeprecationWarning)

	// Delete fields
	s.SchemaPath("cluster", "gcp_attributes").RemoveField("use_preemptible_executors")
	s.SchemaPath("cluster", "gcp_attributes").RemoveField("boot_disk_size")

	// Default values
	s.SchemaPath("edition").SetDefault("ADVANCED")
	s.SchemaPath("channel").SetDefault("CURRENT")

	// ConflictsWith fields
	s.SchemaPath("storage").SetConflictsWith([]string{"catalog"})
	s.SchemaPath("catalog").SetConflictsWith([]string{"storage"})

	// MinItems fields
	s.SchemaPath("library").SetMinItems(1)
	s.SchemaPath("notification", "email_recipients").SetMinItems(1)
	s.SchemaPath("notification", "alerts").SetMinItems(1)

	// MaxItems fields
	s.SchemaPath("cluster", "ssh_public_keys").SetMaxItems(10)
	s.SchemaPath("cluster", "init_scripts").SetMaxItems(10)

	// ValidateFunc fields
	s.SchemaPath("channel").SetValidateFunc(validation.StringInSlice([]string{"current", "preview"}, true))
	s.SchemaPath("edition").SetValidateFunc(validation.StringInSlice([]string{"pro", "core", "advanced"}, true))

	return s
}

var pipelineSchema = common.StructToSchema(Pipeline{}, nil)

// type workspaceClient struct {
// 	*databricks.WorkspaceClient
// }

func ResourcePipeline() common.Resource {
	return common.Resource{
		Schema: pipelineSchema,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			// wsClient := workspaceClient{w}
			err = Create(w, ctx, d, d.Timeout(schema.TimeoutCreate))
			if err != nil {
				return err
			}
			d.Set("url", c.FormatURL("#joblist/pipelines/", d.Id()))
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			// wsClient := workspaceClient{w}
			readPipeline, err := Read(w, ctx, d.Id())

			if err != nil {
				return err
			}
			if readPipeline.Spec == nil {
				return fmt.Errorf("pipeline spec is nil for '%v'", readPipeline.PipelineId)
			}
			// fmt.Println("readPipeline.Spec", readPipeline.Spec)
			return common.StructToData(readPipeline.Spec, pipelineSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			// wsClient := workspaceClient{w}
			return Update(w, ctx, d, d.Timeout(schema.TimeoutUpdate))

		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			// wsClient := workspaceClient{w}
			return Delete(w, ctx, d.Id(), d.Timeout(schema.TimeoutDelete))

		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(DefaultTimeout),
		},
	}
}

// Notes:
// Cluster->enable_local_disk_encryption (not in new)
// Libraries->Whl does not exist in new
// not setting group:node_type for fields in pipelineCluster
// In cluster, ID in old Id in new
// Cluster->Autoscale->Mode (difference in type)
// Cluster->AwsAttributes->Availability (difference in type)
// Cluster->init_scripts->Abfss (difference in type)
// Why is serverless Optional in Serverless
// What to do of dryrun as it affects create response
// AllowDuplicateNames is given a default value in new
// Cluster->AwsAttributes->EbsVolumeIops, EbsVolumeThroughput new
// Cluster->AzureAttributes->LogAnalyticsInfo new
// Gatewaydefinitions and ingestiondefinition in new
// LastestUpdates
// dry_run is in new - wrote unit test for it
// RunAsUserName - checked
// ExpectedLastModified - Test written

// PipelineStateInfo
// PipelineListResponse
