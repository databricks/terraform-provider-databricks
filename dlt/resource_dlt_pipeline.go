package pipelines

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/pipelines"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Pipeline struct {
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

	// Optional fields
	s.SchemaPath("configuration").SetOptional()
	s.SchemaPath("libraries").SetOptional()
	s.SchemaPath("target").SetOptional()
	s.SchemaPath("filters").SetOptional()
	s.SchemaPath("continuous").SetOptional()
	s.SchemaPath("development").SetOptional()
	s.SchemaPath("photon").SetOptional()
	s.SchemaPath("edition").SetOptional()
	s.SchemaPath("channel").SetOptional()
	s.SchemaPath("catalog").SetOptional()
	s.SchemaPath("notifications").SetOptional()
	s.SchemaPath("deployment").SetOptional()
	s.SchemaPath("allow_duplicate_names").SetDefault(false)
	s.SchemaPath("dry_run").SetOptional()

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
