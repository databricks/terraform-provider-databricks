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

const defaultAppProvisionTimeout = 10 * time.Minute
const deleteCallTimeout = 10 * time.Second

var appAliasMap = map[string]string{
	"resources": "resource",
}

type appStruct struct {
	apps.App
}

func (appStruct) Aliases() map[string]map[string]string {
	return map[string]map[string]string{
		"apps.appStruct": appAliasMap,
	}
}

func (appStruct) CustomizeSchema(s *common.CustomizableSchema) *common.CustomizableSchema {

	// Required fields & validation
	s.SchemaPath("name").SetRequired().SetForceNew().SetValidateFunc(validation.StringLenBetween(2, 30))

	// Computed fields
	for _, p := range []string{"active_deployment", "app_status", "compute_status", "create_time", "creator",
		"default_source_code_path", "pending_deployment", "service_principal_id", "service_principal_name",
		"update_time", "updater", "url"} {
		s.SchemaPath(p).SetComputed()
	}
	return s
}

var appSchema = common.StructToSchema(appStruct{}, nil)

func ResourceApp() common.Resource {
	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var createApp appStruct
			common.DataToStructPointer(d, appSchema, &createApp)
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}

			// create the app
			wait, err := w.Apps.Create(ctx, apps.CreateAppRequest{
				App: &createApp.App,
			})
			if err != nil {
				return err
			}
			app, err := wait.GetWithTimeout(d.Timeout(schema.TimeoutCreate) - deleteCallTimeout)
			if err != nil {
				log.Printf("[ERROR] Error waiting for app to be created: %s", err.Error())
				_, nestedErr := w.Apps.DeleteByName(ctx, createApp.Name)
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
			return common.StructToData(appStruct{App: *app}, appSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var update appStruct
			common.DataToStructPointer(d, appSchema, &update)
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			_, err = w.Apps.Update(ctx, apps.UpdateAppRequest{
				App:  &update.App,
				Name: d.Id(),
			})
			if err != nil {
				return err
			}
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			_, err = w.Apps.DeleteByName(ctx, d.Id())
			return err
		},
		Schema: appSchema,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(defaultAppProvisionTimeout),
		},
	}
}
