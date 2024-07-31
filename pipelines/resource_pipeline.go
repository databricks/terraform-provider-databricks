package pipelines

import (
	"context"
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
	// Force send development (as default is true)
	createPipelineRequest.ForceSendFields = append(createPipelineRequest.ForceSendFields, "Development")

	createdPipeline, err := w.Pipelines.Create(ctx, createPipelineRequest.CreatePipeline)
	if err != nil {
		return err
	}
	var id string
	// If dry_run is set, the pipeline is not created and the ID is the effective settings ID.
	if d.Get("dry_run").(bool) {
		id = createdPipeline.EffectiveSettings.Id
		d.SetId(id)
		return fmt.Errorf("dry run succeeded; pipeline %s was not created", id)
	} else {
		id = createdPipeline.PipelineId
	}
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
	// Force send development (as default is true)
	updatePipelineRequest.ForceSendFields = append(updatePipelineRequest.ForceSendFields, "Development")

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

type Pipeline struct {
	pipelines.PipelineSpec
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
	s.SchemaPath("gateway_definition", "connection_id").SetForceNew()
	s.SchemaPath("gateway_definition", "gateway_storage_catalog").SetForceNew()
	s.SchemaPath("gateway_definition", "gateway_storage_schema").SetForceNew()
	s.SchemaPath("ingestion_definition", "connection_name").SetForceNew()
	s.SchemaPath("ingestion_definition", "ingestion_gateway_id").SetForceNew()

	// Computed fields
	s.SchemaPath("id").SetComputed()
	s.SchemaPath("cluster", "node_type_id").SetComputed()
	s.SchemaPath("cluster", "driver_node_type_id").SetComputed()
	s.SchemaPath("cluster", "enable_local_disk_encryption").SetComputed()
	s.SchemaPath("url").SetComputed()

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
	s.SchemaPath("library", "whl").SetDeprecated("The 'whl' field is deprecated")

	// Delete fields
	s.SchemaPath("cluster", "gcp_attributes").RemoveField("use_preemptible_executors")
	s.SchemaPath("cluster", "gcp_attributes").RemoveField("boot_disk_size")

	// Default values
	s.SchemaPath("edition").SetDefault("ADVANCED")
	s.SchemaPath("channel").SetDefault("CURRENT")
	s.SchemaPath(("development")).SetDefault(true)

	// ConflictsWith fields
	s.SchemaPath("storage").SetConflictsWith([]string{"catalog"})
	s.SchemaPath("catalog").SetConflictsWith([]string{"storage"})
	s.SchemaPath("ingestion_definition", "connection_name").SetConflictsWith([]string{"ingestion_definition.0.ingestion_gateway_id"})

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
	s.SchemaPath("gateway_definition").SetRequiredWith([]string{"gateway_definition.0.gateway_storage_name", "gateway_definition.0.gateway_storage_catalog", "gateway_definition.0.gateway_storage_schema"})
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
			readPipeline, err := Read(w, ctx, d.Id())

			if err != nil {
				return err
			}
			if readPipeline.Spec == nil {
				return fmt.Errorf("pipeline spec is nil for '%v'", readPipeline.PipelineId)
			}
			return common.StructToData(readPipeline.Spec, pipelineSchema, d)
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
