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

type App struct {
	apps.App
	// The mode of which the deployment will manage the source code.
	Mode string `json:"mode,omitempty"`
	// The workspace file system path of the source code used to create the app
	// deployment. This is different from
	// `deployment_artifacts.source_code_path`, which is the path used by the
	// deployed app. The former refers to the original source code location of
	// the app in the workspace during deployment creation, whereas the latter
	// provides a system generated stable snapshotted source code path used by
	// the deployment.
	SourceCodePath string `json:"source_code_path,omitempty"`
}

func (App) CustomizeSchema(s *common.CustomizableSchema) *common.CustomizableSchema {

	// Required fields & validation
	s.SchemaPath("name").SetRequired().SetForceNew().SetValidateFunc(validation.StringLenBetween(2, 30))
	s.SchemaPath("mode").SetDefault("SNAPSHOT").SetValidateFunc(validation.StringInSlice([]string{"SNAPSHOT", "AUTO_SYNC"}, false))

	// Computed fields
	for _, p := range []string{"active_deployment", "app_status", "compute_status", "create_time", "creator",
		"default_source_code_path", "pending_deployment", "service_principal_id", "service_principal_name",
		"update_time", "updater", "url"} {
		s.SchemaPath(p).SetComputed()
	}

	// SuppressDiff
	s.SchemaPath("source_code_path").SetCustomSuppressDiff(common.WorkspacePathPrefixDiffSuppress)

	return s
}

var appSchema = common.StructToSchema(App{}, nil)

func ResourceApp() common.Resource {
	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var createApp apps.CreateAppRequest
			var createAppDeployment apps.CreateAppDeploymentRequest
			common.DataToStructPointer(d, appSchema, &createApp)
			common.DataToStructPointer(d, appSchema, &createAppDeployment)
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}

			// create the app, which does not require the source code path yet
			wait, err := w.Apps.Create(ctx, createApp)
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
			// now deploy the app, using the source code path
			createAppDeployment.AppName = app.Name
			waitDeploy, err := w.Apps.Deploy(ctx, createAppDeployment)
			if err != nil {
				return err
			}
			_, err = waitDeploy.GetWithTimeout(d.Timeout(schema.TimeoutCreate) - deleteCallTimeout)
			if err != nil {
				log.Printf("[ERROR] Error waiting for app to be deployed: %s", err.Error())
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
			return common.StructToData(app, appSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var update apps.UpdateAppRequest
			common.DataToStructPointer(d, appSchema, &update)
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			_, err = w.Apps.Update(ctx, update)
			if err != nil {
				return err
			}
			if d.HasChanges("source_code_path", "mode") {
				_, err = w.Apps.Deploy(ctx, apps.CreateAppDeploymentRequest{
					AppName:        d.Id(),
					Mode:           apps.AppDeploymentMode(d.Get("mode").(string)),
					SourceCodePath: d.Get("source_code_path").(string),
				})
				if err != nil {
					return err
				}
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
			Create: schema.DefaultTimeout(defaultEndpointProvisionTimeout),
		},
	}
}
