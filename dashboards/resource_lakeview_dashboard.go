package dashboards

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/dashboards"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceLakeviewDashboard manages lakeview dashboards
func ResourceLakeviewDashboard() common.Resource {
	s := common.StructToSchema(dashboards.Dashboard{},
		func(s map[string]*schema.Schema) map[string]*schema.Schema {
			s["create_time"].Computed = true
			s["dashboard_id"].Computed = true
			s["display_name"].Optional = false
			s["display_name"].Required = true
			s["etag"].Computed = true
			s["lifecycle_state"].Computed = true
			s["parent_path"].Optional = false
			s["parent_path"].Required = true
			s["parent_path"].ForceNew = true
			s["path"].Computed = true
			s["serialized_dashboard"].ConflictsWith = []string{"file_path"}
			s["serialized_dashboard"].DiffSuppressFunc = func(k, old, new string, d *schema.ResourceData) bool {
				_, newHash, err := common.ReadSerializedJsonContent(new, d.Get("file_path").(string))
				if err != nil {
					return false
				}
				return d.Get("md5").(string) == newHash && !d.Get("dashboard_change_detected").(bool)
			}
			s["update_time"].Computed = true
			s["warehouse_id"].Optional = false
			s["warehouse_id"].Required = true

			s["file_path"] = &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"serialized_dashboard"},
			}
			s["md5"] = &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			}
			s["dashboard_change_detected"] = &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			}
			s["embed_credentials"] = &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			}
			return s
		})
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var newDashboardRequest dashboards.CreateDashboardRequest
			common.DataToStructPointer(d, s, &newDashboardRequest)
			content, md5Hash, err := common.ReadSerializedJsonContent(d.Get("serialized_dashboard").(string), d.Get("file_path").(string))
			if err != nil {
				return err
			}
			d.Set("md5", md5Hash)
			newDashboardRequest.SerializedDashboard = content
			createdDashboard, err := w.Lakeview.Create(ctx, newDashboardRequest)
			if err != nil {
				return err
			}

			d.Set("etag", createdDashboard.Etag)

			// We need to 'Force send' the EmbedCredentials field because it is 'omitempty' and if it is not set, it will be ignored. This is a workaround to force the field to be sent if the user wants to set 'embed_credentials' to false.
			_, err = w.Lakeview.Publish(ctx, dashboards.PublishRequest{
				DashboardId:      createdDashboard.DashboardId,
				WarehouseId:      d.Get("warehouse_id").(string),
				EmbedCredentials: d.Get("embed_credentials").(bool),
				ForceSendFields:  []string{"EmbedCredentials"},
			})
			if err != nil {
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
			return common.StructToData(resp, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var updateDashboardRequest dashboards.UpdateDashboardRequest
			common.DataToStructPointer(d, s, &updateDashboardRequest)
			updateDashboardRequest.DashboardId = d.Get("dashboard_id").(string)
			content, md5Hash, err := common.ReadSerializedJsonContent(d.Get("serialized_dashboard").(string), d.Get("file_path").(string))
			if err != nil {
				return err
			}
			d.Set("md5", md5Hash)
			updateDashboardRequest.SerializedDashboard = content
			updatedDashboard, err := w.Lakeview.Update(ctx, updateDashboardRequest)
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
