package pipelines

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/pipelines"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

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
