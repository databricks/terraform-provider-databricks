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
			d.Set("registered_model_id", res.RegisteredModel.Name) // alias
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var req mlflow.GetRegisteredModelRequest
			common.DataToStructPointer(d, s, &req)
			res, err := w.RegisteredModels.Get(ctx, req)
			if err != nil {
				return err
			}
			return common.StructToData(res, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var req mlflow.UpdateRegisteredModelRequest
			common.DataToStructPointer(d, s, &req)
			return w.RegisteredModels.Update(ctx, req)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var req mlflow.DeleteRegisteredModelRequest
			common.DataToStructPointer(d, s, &req)
			return w.RegisteredModels.Delete(ctx, req)
		},
		Schema: s,
	}.ToResource()
}
