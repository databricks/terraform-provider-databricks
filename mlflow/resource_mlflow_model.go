package mlflow

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/mlflow"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceMlflowModel() *schema.Resource {
	s := common.StructToSchema(
		mlflow.RegisteredModel{},
		func(s map[string]*schema.Schema) map[string]*schema.Schema {
			delete(s, "latest_versions")
			s["name"] = &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			}
			return s
		})

	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var req mlflow.CreateRegisteredModelRequest
			common.DataToStructPointer(d, s, &req)
			res, err := w.RegisteredModels.Create(ctx, req)
			if err != nil {
				return err
			}
			d.SetId(res.RegisteredModel.Name)
			d.Set("registered_model_id", res.RegisteredModel.Name)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			model, err := w.RegisteredModels.GetByName(ctx, d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(model, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			err = w.RegisteredModels.Update(ctx, mlflow.UpdateRegisteredModelRequest{
				Description: d.Get("description").(string),
				Name:        d.Get("name").(string),
			})
			return err
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			name := d.Get("name").(string)
			w, err := c.WorkspaceClient()
			_ = name
			if err != nil {
				return err
			}
			err = w.RegisteredModels.Delete(ctx, mlflow.DeleteRegisteredModelRequest{
				Name: d.Get("name").(string),
			})
			return err
		},
		Schema: s,
	}.ToResource()
}
