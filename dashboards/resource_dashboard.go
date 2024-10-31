package dashboards

import (
	"context"
	"log"
	"strings"

	"github.com/databricks/databricks-sdk-go/service/dashboards"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Dashboard struct {
	dashboards.Dashboard
	EmbedCredentials        bool   `json:"embed_credentials,omitempty"`
	FilePath                string `json:"file_path,omitempty"`
	Md5                     string `json:"md5,omitempty"`
	DashboardChangeDetected bool   `json:"dashboard_change_detected,omitempty"`
}

func customDiffSerializedDashboard(k, old, new string, d *schema.ResourceData) bool {
	_, newHash, err := common.ReadSerializedJsonContent(new, d.Get("file_path").(string))
	if err != nil {
		return false
	}
	return d.Get("md5").(string) == newHash && !d.Get("dashboard_change_detected").(bool)
}

func (Dashboard) CustomizeSchema(s *common.CustomizableSchema) *common.CustomizableSchema {
	// Required fields
	s.SchemaPath("display_name").SetRequired()
	s.SchemaPath("parent_path").SetRequired()
	s.SchemaPath("warehouse_id").SetRequired()

	// Computed fields
	s.SchemaPath("create_time").SetComputed()
	s.SchemaPath("dashboard_id").SetComputed()
	s.SchemaPath("etag").SetComputed()
	s.SchemaPath("lifecycle_state").SetComputed()
	s.SchemaPath("path").SetComputed()
	s.SchemaPath("update_time").SetComputed()
	s.SchemaPath("md5").SetComputed()

	// ForceNew fields
	s.SchemaPath("parent_path").SetCustomSuppressDiff(common.WorkspacePathPrefixDiffSuppress).SetForceNew()

	// ConflictsWith fields
	s.SchemaPath("serialized_dashboard").SetConflictsWith([]string{"file_path"})
	s.SchemaPath("file_path").SetConflictsWith([]string{"serialized_dashboard"})

	// Default values
	s.SchemaPath("embed_credentials").SetDefault(true)

	// DiffSuppressFunc
	s.SchemaPath("serialized_dashboard").SetCustomSuppressDiff(customDiffSerializedDashboard)

	return s
}

var dashboardSchema = common.StructToSchema(Dashboard{}, nil)

// ResourceDashboard manages dashboards
func ResourceDashboard() common.Resource {
	return common.Resource{
		Schema: dashboardSchema,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var dashboard dashboards.Dashboard
			common.DataToStructPointer(d, dashboardSchema, &dashboard)
			content, md5Hash, err := common.ReadSerializedJsonContent(d.Get("serialized_dashboard").(string), d.Get("file_path").(string))
			if err != nil {
				return err
			}
			d.Set("md5", md5Hash)
			dashboard.SerializedDashboard = content
			createdDashboard, err := w.Lakeview.Create(ctx, dashboards.CreateDashboardRequest{Dashboard: &dashboard})
			if err != nil && isParentDoesntExistError(err) {
				log.Printf("[DEBUG] Parent folder '%s' doesn't exist, creating...", dashboard.ParentPath)
				err = w.Workspace.MkdirsByPath(ctx, dashboard.ParentPath)
				if err != nil {
					return err
				}
				createdDashboard, err = w.Lakeview.Create(ctx, dashboards.CreateDashboardRequest{Dashboard: &dashboard})
			}
			if err != nil {
				return err
			}

			d.Set("etag", createdDashboard.Etag)

			// We need to 'Force send' the EmbedCredentials field because it is 'omitempty' and if it is not set, it will be ignored. This is a workaround to force the field to be sent if the user wants to set 'embed_credentials' to false.
			_, err = w.Lakeview.Publish(ctx, dashboards.PublishRequest{
				DashboardId:      createdDashboard.DashboardId,
				WarehouseId:      createdDashboard.WarehouseId,
				EmbedCredentials: d.Get("embed_credentials").(bool),
				ForceSendFields:  []string{"EmbedCredentials"},
			})
			if err != nil {
				// If the publish fails, we should delete the dashboard to avoid leaving it in a bad state.
				deleteErr := w.Lakeview.Trash(ctx, dashboards.TrashDashboardRequest{
					DashboardId: createdDashboard.DashboardId,
				})
				if deleteErr != nil {
					return deleteErr
				}
				return err
			}

			d.SetId(createdDashboard.DashboardId)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			dashboard := dashboards.GetDashboardRequest{
				DashboardId: d.Id(),
			}
			resp, err := w.Lakeview.Get(ctx, dashboard)
			if err != nil {
				return err
			}
			d.Set("dashboard_change_detected", (resp.Etag != d.Get("etag").(string)))
			return common.StructToData(resp, dashboardSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var dashboard dashboards.Dashboard
			common.DataToStructPointer(d, dashboardSchema, &dashboard)
			dashboard.DashboardId = d.Id()
			content, md5Hash, err := common.ReadSerializedJsonContent(d.Get("serialized_dashboard").(string), d.Get("file_path").(string))
			if err != nil {
				return err
			}
			d.Set("md5", md5Hash)
			dashboard.SerializedDashboard = content
			updatedDashboard, err := w.Lakeview.Update(ctx, dashboards.UpdateDashboardRequest{
				DashboardId: dashboard.DashboardId,
				Dashboard: &dashboard,
			})
			if err != nil {
				return err
			}
			d.Set("etag", updatedDashboard.Etag)

			// We need to 'Force send' the EmbedCredentials field because it is 'omitempty' and if it is not set, it will be ignored. This is a workaround to force the field to be sent if the user wants to set 'embed_credentials' to false.
			_, err = w.Lakeview.Publish(ctx, dashboards.PublishRequest{
				DashboardId:      d.Id(),
				WarehouseId:      d.Get("warehouse_id").(string),
				EmbedCredentials: d.Get("embed_credentials").(bool),
				ForceSendFields:  []string{"EmbedCredentials"},
			})
			if err != nil {
				// If the publish fails, we should delete the dashboard to avoid leaving it in a bad state.
				deleteErr := w.Lakeview.Trash(ctx, dashboards.TrashDashboardRequest{
					DashboardId: d.Id(),
				})
				if deleteErr != nil {
					return deleteErr
				}
				return err
			}

			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return w.Lakeview.Trash(ctx, dashboards.TrashDashboardRequest{
				DashboardId: d.Id(),
			})
		},
	}
}

func isParentDoesntExistError(err error) bool {
	errStr := err.Error()
	return strings.HasPrefix(errStr, "Path (") && strings.HasSuffix(errStr, ") doesn't exist.")
}
