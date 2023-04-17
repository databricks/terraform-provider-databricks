package mlflow

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/mlflow"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceMlflowExperiment() *schema.Resource {
	s := common.StructToSchema(
		mlflow.Experiment{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			return m
		})

	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			exp, err := w.Experiments.Create(ctx, mlflow.CreateExperiment{
				ArtifactLocation: d.Get("artifact_location").(string),
				Name:             d.Get("name").(string),
			})
			if err != nil {
				return err
			}
			d.SetId(exp.ExperimentId)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			exp, err := w.Experiments.GetByExperimentId(ctx, d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(exp, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			err = w.Experiments.Update(ctx, mlflow.UpdateExperiment{
				ExperimentId: d.Id(),
				NewName:      d.Get("name").(string), // TODO:check this
			})
			return err
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			err = w.Experiments.DeleteByExperimentId(ctx, d.Id())
			return err

		},
		StateUpgraders: []schema.StateUpgrader{},
		Schema:         s,
		SchemaVersion:  0,
		Timeouts:       &schema.ResourceTimeout{},
	}.ToResource()
}
