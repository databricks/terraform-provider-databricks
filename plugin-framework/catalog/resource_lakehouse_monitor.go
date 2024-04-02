package pluginframework

import (
	"context"
	"time"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginframework "github.com/databricks/terraform-provider-databricks/plugin-framework"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

const lakehouseMonitorDefaultProvisionTimeout = 15 * time.Minute

type LakehouseMonitorResource struct {
	pluginframework.DatabricksResource
}

func (r *LakehouseMonitorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
}

func (r *LakehouseMonitorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	req = resource.SchemaRequest{}
}

func (r *LakehouseMonitorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	c := common.DatabricksClient{}
	w, err := c.WorkspaceClient()
	if err != nil {
		resp.Diagnostics.AddError("Failed to get workspace client", err.Error())
		return
	}
	var create catalog.CreateMonitor
	common.DataToStructPointer(d, monitorSchema, &create)
	create.FullName = d.Get("table_name").(string)

	endpoint, err := w.LakehouseMonitors.Create(ctx, create)
	if err != nil {
		return err
	}
	err = WaitForMonitor(w, ctx, create.FullName)
	if err != nil {
		return err
	}
	d.SetId(endpoint.TableName)
	return nil
}

func (r *LakehouseMonitorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (r *LakehouseMonitorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *LakehouseMonitorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
