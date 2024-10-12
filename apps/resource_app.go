package apps

import (
	"context"
	"log"
	"time"

	"github.com/databricks/databricks-sdk-go/service/apps"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const defaultEndpointProvisionTimeout = 75 * time.Minute
const deleteCallTimeout = 10 * time.Second

func ResourceApp() common.Resource {
	s := common.StructToSchema(apps.App{}, func(m map[string]*schema.Schema) map[string]*schema.Schema {
		common.CustomizeSchemaPath(m, "name").SetRequired().SetForceNew().SetValidateFunc(validation.StringLenBetween(2, 30))
		for _, p := range []string{"active_deployment", "app_status", "compute_status", "create_time", "creator",
			"default_source_code_path", "pending_deployment", "service_principal_id", "service_principal_name",
			"update_time", "updater", "url"} {
			common.CustomizeSchemaPath(m, p).SetComputed()
		}
		return m
	})
	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var create apps.CreateAppRequest
			common.DataToStructPointer(d, s, &create)
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			wait, err := w.Apps.Create(ctx, create)
			if err != nil {
				return err
			}
			app, err := wait.GetWithTimeout(d.Timeout(schema.TimeoutCreate) - deleteCallTimeout)
			if err != nil {
				log.Printf("[ERROR] Error waiting for app to be created: %s", err.Error())
				_, nestedErr := w.Apps.DeleteByName(ctx, create.Name)
				if nestedErr != nil {
					log.Printf("[ERROR] Error cleaning up app: %s", nestedErr.Error())
				}
				return err
			}
			d.SetId(app.Name)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			app, err := w.Apps.GetByName(ctx, d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(app, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var update apps.UpdateAppRequest
			common.DataToStructPointer(d, s, &update)
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			app, err := w.Apps.Update(ctx, update)
			if err != nil {
				return err
			}
			return common.StructToData(app, s, d)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			_, err = w.Apps.DeleteByName(ctx, d.Id())
			return err
		},
		Schema: s,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(defaultEndpointProvisionTimeout),
		},
	}
}
