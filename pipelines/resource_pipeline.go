package pipelines

import (
	"context"
	"errors"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/pipelines"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

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
	adjustForceSendFields(&createPipelineRequest.Clusters)

	createdPipeline, err := w.Pipelines.Create(ctx, createPipelineRequest.CreatePipeline)
	if err != nil {
		return err
	}
	id := createdPipeline.PipelineId
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
	return nil
}

func Read(w *databricks.WorkspaceClient, ctx context.Context, id string) (*pipelines.GetPipelineResponse, error) {
	return w.Pipelines.Get(ctx, pipelines.GetPipelineRequest{
		PipelineId: id,
	})
}

func Update(w *databricks.WorkspaceClient, ctx context.Context, d *schema.ResourceData, timeout time.Duration) error {
	var updatePipelineRequest updatePipelineRequestStruct
	common.DataToStructPointer(d, pipelineSchema, &updatePipelineRequest)
	updatePipelineRequest.EditPipeline.PipelineId = d.Id()
	adjustForceSendFields(&updatePipelineRequest.Clusters)
	// Workspaces not enrolled in the private preview must not send run_as in the update request.
	// If run_as was persisted in state because of a `terraform refresh`, there will not be a planned change
	// as long as the user hasn't specified a value.
	if !d.HasChange("run_as") {
		updatePipelineRequest.RunAs = nil
	}

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
			return retry.RetryableError(errors.New(message))
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
			return retry.RetryableError(errors.New(message))
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

type Pipeline struct {
	pipelines.PipelineSpec
	AllowDuplicateNames  bool                                `json:"allow_duplicate_names,omitempty"`
	Cause                string                              `json:"cause,omitempty"`
	ClusterId            string                              `json:"cluster_id,omitempty"`
	CreatorUserName      string                              `json:"creator_user_name,omitempty"`
	Health               pipelines.GetPipelineResponseHealth `json:"health,omitempty"`
	LastModified         int64                               `json:"last_modified,omitempty"`
	LatestUpdates        []pipelines.UpdateStateInfo         `json:"latest_updates,omitempty"`
	RunAs                pipelines.RunAs                     `json:"run_as,omitempty"`
	RunAsUserName        string                              `json:"run_as_user_name,omitempty"`
	ExpectedLastModified int64                               `json:"expected_last_modified,omitempty"`
	State                pipelines.PipelineState             `json:"state,omitempty"`
	// Provides the URL to the pipeline in the Databricks UI.
	URL string `json:"url,omitempty"`
}

func (Pipeline) Aliases() map[string]map[string]string {
	return map[string]map[string]string{
		"pipelines.Pipeline":     aliasMap,
		"pipelines.PipelineSpec": aliasMap,
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
	s.SchemaPath("gateway_definition", "connection_id").SetForceNew()
	s.SchemaPath("gateway_definition", "gateway_storage_catalog").SetForceNew()
	s.SchemaPath("gateway_definition", "gateway_storage_schema").SetForceNew()
	s.SchemaPath("ingestion_definition", "connection_name").SetForceNew()
	s.SchemaPath("ingestion_definition", "ingestion_gateway_id").SetForceNew()

	// Required fields
	s.SchemaPath("library", "glob", "include").SetRequired()
	s.SchemaPath("library", "notebook", "path").SetRequired()
	s.SchemaPath("library", "file", "path").SetRequired()

	// Computed fields
	s.SchemaPath("cluster", "node_type_id").SetComputed()
	s.SchemaPath("cluster", "driver_node_type_id").SetComputed()
	s.SchemaPath("cluster", "enable_local_disk_encryption").SetComputed()

	for _, field := range []string{"id", "state", "latest_updates", "last_modified",
		"health", "cause", "cluster_id", "creator_user_name", "run_as", "url", "run_as_user_name"} {
		s.SchemaPath(field).SetComputed()
	}

	// customize event_log
	s.SchemaPath("event_log", "name").SetRequired()
	s.SchemaPath("event_log", "catalog").SetComputed()
	s.SchemaPath("event_log", "schema").SetComputed()

	// SuppressDiff fields
	s.SchemaPath("edition").SetSuppressDiff()
	s.SchemaPath("channel").SetSuppressDiff()
	s.SchemaPath("cluster", "spark_conf").SetCustomSuppressDiff(clusters.SparkConfDiffSuppressFunc)
	s.SchemaPath("cluster", "aws_attributes", "zone_id").SetCustomSuppressDiff(clusters.ZoneDiffSuppress)
	s.SchemaPath("cluster", "autoscale", "mode").SetCustomSuppressDiff(common.EqualFoldDiffSuppress)
	s.SchemaPath("edition").SetCustomSuppressDiff(common.EqualFoldDiffSuppress)
	s.SchemaPath("storage").SetCustomSuppressDiff(suppressStorageDiff)

	// As of 6th Nov 2024, the DLT API only normalizes the catalog name when creating
	// a pipeline. So we only ignore the equal fold diff for the catalog name and not other
	// UC resources like target, schema or ingestion_definition.connection_name.
	s.SchemaPath("catalog").SetCustomSuppressDiff(common.EqualFoldDiffSuppress)

	// Deprecated fields
	s.SchemaPath("cluster", "init_scripts", "dbfs").SetDeprecated(clusters.DbfsDeprecationWarning)
	s.SchemaPath("library", "whl").SetDeprecated("The 'whl' field is deprecated")

	// Delete fields
	s.SchemaPath("cluster", "gcp_attributes").RemoveField("use_preemptible_executors")
	s.SchemaPath("cluster", "gcp_attributes").RemoveField("boot_disk_size")

	// Default values
	s.SchemaPath("edition").SetDefault("ADVANCED")
	s.SchemaPath("channel").SetDefault("CURRENT")

	// ConflictsWith fields
	s.SchemaPath("storage").SetConflictsWith([]string{"catalog"})
	s.SchemaPath("catalog").SetConflictsWith([]string{"storage"})
	s.SchemaPath("ingestion_definition", "connection_name").SetConflictsWith([]string{"ingestion_definition.0.ingestion_gateway_id"})
	s.SchemaPath("target").SetConflictsWith([]string{"schema"})
	s.SchemaPath("schema").SetConflictsWith([]string{"target"})

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

	// RequiredWith fields
	s.SchemaPath("ingestion_definition").SetRequiredWith([]string{"ingestion_definition.0.objects"})

	return s
}

var pipelineSchema = common.StructToSchema(Pipeline{}, nil)

func ResourcePipeline() common.Resource {
	return common.Resource{
		Schema: pipelineSchema,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return Create(w, ctx, d, d.Timeout(schema.TimeoutCreate))
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			readPipeline, err := Read(w, ctx, d.Id())

			if err != nil {
				return err
			}
			if readPipeline.Spec == nil {
				return fmt.Errorf("pipeline spec is nil for '%v'", readPipeline.PipelineId)
			}
			p := Pipeline{
				PipelineSpec:    *readPipeline.Spec,
				Cause:           readPipeline.Cause,
				ClusterId:       readPipeline.ClusterId,
				CreatorUserName: readPipeline.CreatorUserName,
				Health:          readPipeline.Health,
				LastModified:    readPipeline.LastModified,
				LatestUpdates:   readPipeline.LatestUpdates,
				RunAsUserName:   readPipeline.RunAsUserName,
				State:           readPipeline.State,
				// Provides the URL to the pipeline in the Databricks UI.
				URL: c.FormatURL("#joblist/pipelines/", d.Id()),
			}
			if readPipeline.RunAsUserName != "" {
				if common.StringIsUUID(readPipeline.RunAsUserName) {
					p.RunAs = pipelines.RunAs{
						ServicePrincipalName: readPipeline.RunAsUserName,
					}
				} else {
					p.RunAs = pipelines.RunAs{
						UserName: readPipeline.RunAsUserName,
					}
				}
			}
			return common.StructToData(p, pipelineSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return Update(w, ctx, d, d.Timeout(schema.TimeoutUpdate))

		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return Delete(w, ctx, d.Id(), d.Timeout(schema.TimeoutDelete))

		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(DefaultTimeout),
		},
	}
}
