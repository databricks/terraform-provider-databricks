package apps

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/databricks/databricks-sdk-go/service/apps"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	defaultAppProvisionTimeout = 10 * time.Minute
	deleteCallTimeout          = 10 * time.Second
)

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

	// Name is required and cannot be updated
	s.SchemaPath("name").SetRequired().SetForceNew()
	// Resources should be a set
	s.SchemaPath("resource").SetSliceSet()
	// Computed fields
	for _, p := range []string{
		"active_deployment",
		"app_status",
		"compute_status",
		"create_time",
		"creator",
		"default_source_code_path",
		"pending_deployment",
		"service_principal_client_id",
		"service_principal_id",
		"service_principal_name",
		"update_time",
		"updater",
		"url",
	} {
		s.SchemaPath(p).SetComputed()
	}
	return s
}

var appSchema = common.StructToSchema(appStruct{}, nil)

// each resource block should have exactly one resource type
func appHasExactlyOneOfResourceType(d *schema.ResourceData) bool {
	if _, ok := d.GetOk("resource"); ok {
		// resources is a TF set
		resources := d.Get("resource").(*schema.Set).List()
		for _, resource := range resources {
			resource := resource.(map[string]interface{})
			count := 0
			for _, v := range resource {
				// each resource type is stored as a list of maps. check for non-empty list
				if value, ok := v.([]interface{}); ok {
					if len(value) == 0 {
						continue
					}
					count++
				}
			}
			if count != 1 {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func ResourceApp() common.Resource {
	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var createApp appStruct
			common.DataToStructPointer(d, appSchema, &createApp)
			if appHasExactlyOneOfResourceType(d) == false {
				return errors.New("Exactly one resource type per resource block should be provided")
			}
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
			d.SetId(wait.Name)
			// wait for up to the create timeout, accounting for the deletion on failure.
			_, err = wait.GetWithTimeout(d.Timeout(schema.TimeoutCreate) - deleteCallTimeout)
			if err != nil {
				log.Printf("[ERROR] Error waiting for app to be created: %s", err.Error())
				_, nestedErr := w.Apps.DeleteByName(ctx, createApp.Name)
				if nestedErr != nil {
					log.Printf("[ERROR] Error cleaning up app: %s", nestedErr.Error())
				}
				return err
			}
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
			if appHasExactlyOneOfResourceType(d) == false {
				return errors.New("Exactly one resource type per resource block should be provided")
			}
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
