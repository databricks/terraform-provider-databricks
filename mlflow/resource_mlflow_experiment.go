package mlflow

import (
	"context"
	"log"
	"strings"

	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func experimentNameSuppressDiff(k, old, new string, d *schema.ResourceData) bool {

	if strings.TrimSuffix(strings.TrimPrefix(new, "/Workspace"), "/") == strings.TrimSuffix(strings.TrimPrefix(old, "/Workspace"), "/") {
		log.Printf("[DEBUG] Ignoring configuration drift from %s to %s", old, new)
		return true
	}
	return false
}

func ResourceMlflowExperiment() common.Resource {
	s := common.StructToSchema(
		ml.Experiment{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			for _, p := range []string{"creation_time", "experiment_id", "last_update_time", "lifecycle_stage"} {
				common.CustomizeSchemaPath(m, p).SetComputed()
			}
			common.CustomizeSchemaPath(m, "artifact_location").SetForceNew().SetSuppressDiff()
			common.CustomizeSchemaPath(m, "name").SetRequired().SetCustomSuppressDiff(experimentNameSuppressDiff)
			return m
		})

	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var create ml.CreateExperiment
			common.DataToStructPointer(d, s, &create)
			response, err := w.Experiments.CreateExperiment(ctx, create)
			if err != nil {
				return err
			}
			d.SetId(response.ExperimentId)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			e, err := w.Experiments.GetExperiment(ctx, ml.GetExperimentRequest{ExperimentId: d.Id()})
			if err != nil {
				return err
			}
			return common.StructToData(e.Experiment, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var update ml.UpdateExperiment
			common.DataToStructPointer(d, s, &update)
			return w.Experiments.UpdateExperiment(ctx, update)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return w.Experiments.DeleteExperiment(ctx, ml.DeleteExperiment{ExperimentId: d.Id()})
		},
		StateUpgraders: []schema.StateUpgrader{},
		Schema:         s,
		SchemaVersion:  0,
		Timeouts:       &schema.ResourceTimeout{},
	}
}
