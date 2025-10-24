package pools

import (
	"context"
	"fmt"
	"reflect"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/compute_tf"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const dataSourceName = "instance_pool"

func DataSourceInstancePool() datasource.DataSource {
	return &InstancePoolDataSource{}
}

var _ datasource.DataSourceWithConfigure = &InstancePoolDataSource{}

type InstancePoolDataSource struct {
	Client *common.DatabricksClient
}

type InstancePoolInfo struct {
	Id       types.String `tfsdk:"id"`
	Name     types.String `tfsdk:"name"`
	PoolInfo types.Object `tfsdk:"pool_info"`
}

func (InstancePoolInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["pool_info"] = attrs["pool_info"].SetComputed()

	return attrs
}

func (InstancePoolInfo) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"pool_info": reflect.TypeOf(compute_tf.InstancePoolAndStats{}),
	}
}

func (d *InstancePoolDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(dataSourceName)
}

func (d *InstancePoolDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, InstancePoolInfo{}, nil)
	resp.Schema = schema.Schema{
		Attributes: attrs,
		Blocks:     blocks,
	}
}

func (d *InstancePoolDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if d.Client == nil {
		d.Client = pluginfwcommon.ConfigureDataSource(req, resp)
	}
}

func (d *InstancePoolDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)
	w, diags := d.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var poolInfo InstancePoolInfo
	resp.Diagnostics.Append(req.Config.Get(ctx, &poolInfo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	poolName := poolInfo.Name.ValueString()
	pool, err := d.getInstancePoolByName(ctx, w, poolName)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("failed to get instance pool '%s'", poolName), err.Error())
		return
	}

	var tfPool compute_tf.InstancePoolAndStats
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, pool, &tfPool)...)
	if resp.Diagnostics.HasError() {
		return
	}

	poolInfo.Id = types.StringValue(pool.InstancePoolId)
	poolInfo.Name = types.StringValue(pool.InstancePoolName)
	poolInfo.PoolInfo = tfPool.ToObjectValue(ctx)
	resp.Diagnostics.Append(resp.State.Set(ctx, poolInfo)...)
}

func (d *InstancePoolDataSource) getInstancePoolByName(ctx context.Context, w *databricks.WorkspaceClient, poolName string) (*compute.InstancePoolAndStats, error) {
	pool, err := w.InstancePools.GetByInstancePoolName(ctx, poolName)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
