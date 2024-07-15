package pipelines

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/pipelines"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Pipeline struct {
	// not setting group:node_type and max_items:10 for fields in pipelineCluster
	// Cluster->Autoscale->Mode (difference in type)
	// Cluster->AwsAttributes->Availability (difference in type)
	// Cluster->enable_local_disk_encryption (not in new)
	// In cluster, ID in old Id in new
	// Cluster->init_scripts->Abfss (difference in type)
	// not setting alias for Clusters in pipelineCluster
	// Libraries->Whl does not exist in new
	// Not setting alias for Libraries
	// Not setting alias for Notifications
	// Did not set min_items:1 for sub-fields of Notifications
	// Why is serverless Optional in Serverless
	pipelines.PipelineSpec
	AllowDuplicateNames bool `json:"allow_duplicate_names,omitempty"`
	DryRun              bool `json:"dry_run,omitempty"`
	// An optional message detailing the cause of the pipeline state.
	Cause string `json:"cause,omitempty"`
	// The ID of the cluster that the pipeline is running on.
	ClusterId string `json:"cluster_id,omitempty"`
	// The username of the pipeline creator.
	CreatorUserName string `json:"creator_user_name,omitempty"`
	// The health of a pipeline.
	Health pipelines.GetPipelineResponseHealth `json:"health,omitempty"`
	// The last time the pipeline settings were modified or created.
	LastModified int64 `json:"last_modified,omitempty"`
	// Status of the latest updates for the pipeline. Ordered with the newest
	// update first.
	LatestUpdates []pipelines.UpdateStateInfo `json:"latest_updates,omitempty"`
	// Username of the user that the pipeline will run on behalf of.
	RunAsUserName string `json:"run_as_user_name,omitempty"`
	// The pipeline state.
	State pipelines.PipelineState `json:"state,omitempty"`
}

func (Pipeline) CustomizeSchema(s *common.CustomizableSchema) *common.CustomizableSchema {
	// Required fields
	s.SchemaPath("name").SetRequired()
	s.SchemaPath("storage").SetRequired()
	s.SchemaPath("clusters").SetRequired()
	s.SchemaPath("continuous").SetRequired()

	// ForceNew fields
	s.SchemaPath("storage").SetForceNew()
	s.SchemaPath("catalog").SetForceNew()

	// Computed fields
	s.SchemaPath("id").SetComputed()
	s.SchemaPath("clusters", "node_type_id").SetComputed()
	s.SchemaPath("clusters", "driver_node_type_id").SetComputed()
	s.SchemaPath("clusters", "enable_local_disk_encryption").SetComputed()

	// SuppressDiff fields
	s.SchemaPath("edition").SetSuppressDiff()
	s.SchemaPath("channel").SetSuppressDiff()
	s.SchemaPath("clusters", "spark_conf").SetCustomSuppressDiff(clusters.SparkConfDiffSuppressFunc)
	s.SchemaPath("clusters", "aws_attributes", "zone_id").SetCustomSuppressDiff(clusters.ZoneDiffSuppress)
	s.SchemaPath("clusters", "autoscale", "mode").SetCustomSuppressDiff(common.EqualFoldDiffSuppress)

	// Deprecated fields
	s.SchemaPath("clusters", "init_scripts", "dbfs").SetDeprecated(clusters.DbfsDeprecationWarning)

	// Delete fields
	s.SchemaPath("clusters", "gcp_attributes").RemoveField("use_preemptible_executors")
	s.SchemaPath("clusters", "gcp_attributes").RemoveField("boot_disk_size")

	// Default values
	s.SchemaPath("edition").SetDefault("ADVANCED")
	s.SchemaPath("channel").SetDefault("CURRENT")

	// ConflictsWith fields
	s.SchemaPath("storage").SetConflictsWith([]string{"catalog"})
	s.SchemaPath("catalog").SetConflictsWith([]string{"storage"})

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
			var createPipelineRequest pipelines.CreatePipeline
			common.DataToStructPointer(d, pipelineSchema, &createPipelineRequest)
			createdPipeline, err := w.Pipelines.Create(ctx, createPipelineRequest)
			if err != nil {
				return err
			}
			d.SetId(createdPipeline.PipelineId)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			readPipeline, err := w.Pipelines.Get(ctx, pipelines.GetPipelineRequest{
				PipelineId: d.Id(),
			})
			if err != nil {
				return err
			}
			return common.StructToData(readPipeline, pipelineSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var updatePipelineRequest pipelines.EditPipeline
			common.DataToStructPointer(d, pipelineSchema, &updatePipelineRequest)
			return w.Pipelines.Update(ctx, updatePipelineRequest)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return w.Pipelines.Delete(ctx, pipelines.DeletePipelineRequest{
				PipelineId: d.Id(),
			})
		},
	}
}
