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

func ResourcePipeline() common.Resource {
	s := common.StructToSchema(pipelines.CreatePipeline{},
		func(s map[string]*schema.Schema) map[string]*schema.Schema {
			s["pipeline_id"].Computed = true
			s["name"].Optional = false
			s["name"].Required = true
			// does this require force new?
			s["storage"].Optional = false
			s["storage"].Required = true
			// see it (if optional, then remove)
			s["configuration"].Optional = true
			s["clusters"].Optional = false
			s["clusters"].Required = true
			// see it (if optional, then remove)
			s["libraries"].Optional = true
			// see it (if optional, then remove)
			s["target"].Optional = true

			s["filters"].Optional = true

			s["continuous"].Optional = true

			s["development"].Optional = true
			s["development"].Default = false

			s["photon"].Optional = true

			s["edition"].Optional = true

			s["channel"].Optional = true

			s["catalog"].Optional = true

			s["notifications"].Optional = true

			s["deployment"].Optional = true

			s["allow_duplicate_names"].Default = false

			s["dry_run"].Optional = true

			// s["spec"] = &schema.Schema{
			// 	Type: schema.TypeString,

			return s
		})
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var new_pipeline pipelines.CreatePipeline
			common.DataToStructPointer(d, s, &new_pipeline)
			created_pipeline, err := w.Pipelines.Create(ctx, new_pipeline)
			if err != nil {
				return err
			}
			d.SetId(created_pipeline.PipelineId)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			pipeline, err := w.Pipelines.Get(ctx, pipelines.GetPipelineRequest{
				PipelineId: d.Id(),
			})
			if err != nil {
				return err
			}
			return common.StructToData(pipeline, s, d)
		},
	}
}
