package catalog

import (
	"context"
	"time"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const DefaultProvisionTimeout = 15 * time.Minute

func ResourceLakehouseMonitor() common.Resource {
	monitorSchema := common.StructToSchema(
		catalog.MonitorInfo{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			common.CustomizeSchemaPath(m, "assets_dir").SetRequired()
			common.CustomizeSchemaPath(m, "output_schema_name").SetRequired()
			common.CustomizeSchemaPath(m, "table_name").SetRequired()
			common.CustomizeSchemaPath(m).AddNewField("skip_builtin_dashboard", &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Required: false,
			})
			common.CustomizeSchemaPath(m).AddNewField("warehouse_id", &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Required: false,
			})
			return m
		},
	)

	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}

			var create catalog.CreateMonitor
			common.DataToStructPointer(d, monitorSchema, &create)
			create.FullName = d.Get("table_name").(string)

			endpoint, err := w.LakehouseMonitors.Create(ctx, create)
			if err != nil {
				return err
			}
			d.SetId(endpoint.TableName)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			endpoint, err := w.LakehouseMonitors.GetByFullName(ctx, d.Id())
			if err != nil {
				return err

			}
			err = common.StructToData(endpoint, monitorSchema, d)
			if err != nil {
				return err
			}
			return nil
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var update catalog.UpdateMonitor
			common.DataToStructPointer(d, monitorSchema, &update)
			update.FullName = d.Get("table_name").(string)
			_, err = w.LakehouseMonitors.Update(ctx, update)
			return err
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return w.LakehouseMonitors.DeleteByFullName(ctx, d.Id())
		},
		Schema: monitorSchema,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(DefaultProvisionTimeout),
		},
	}
}
