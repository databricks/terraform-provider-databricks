package mlflow

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceMlflowModel() *schema.Resource {
	s := common.StructToSchema(
		ml.Model{},
		func(s map[string]*schema.Schema) map[string]*schema.Schema {
			delete(s, "latest_versions")
			s["name"] = &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			}
			s["registered_model_id"] = &schema.Schema{
				Computed: true,
				Type:     schema.TypeString,
			}
			return s
		})

	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var req ml.CreateModelRequest
			common.DataToStructPointer(d, s, &req)
			res, err := w.ModelRegistry.CreateModel(ctx, req)
			if err != nil {
				return err
			}
			model, err := w.ModelRegistry.GetModel(ctx, ml.GetModelRequest{
				Name: res.RegisteredModel.Name,
			})
			if err != nil {
				return err
			}
			d.SetId(model.RegisteredModelDatabricks.Name)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var req ml.GetModelRequest
			common.DataToStructPointer(d, s, &req)
			req.Name = d.Id()
			res, err := w.ModelRegistry.GetModel(ctx, req)
			if err != nil {
				return err
			}
			err = common.StructToData(res, s, d)
			if err != nil {
				return err
			}
			d.Set("registered_model_id", res.RegisteredModelDatabricks.Id) // alias
			return nil
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var req ml.UpdateModelRequest
			common.DataToStructPointer(d, s, &req)
			return w.ModelRegistry.UpdateModel(ctx, req)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var req ml.DeleteModelRequest
			common.DataToStructPointer(d, s, &req)
			req.Name = d.Id()
			return w.ModelRegistry.DeleteModel(ctx, req)
		},
		Schema: s,
	}.ToResource()
}
