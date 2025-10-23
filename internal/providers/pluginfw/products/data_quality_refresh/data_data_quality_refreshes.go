// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package data_quality_refresh

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/dataquality"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const dataSourcesName = "data_quality_refreshes"

var _ datasource.DataSourceWithConfigure = &RefreshesDataSource{}

func DataSourceRefreshes() datasource.DataSource {
	return &RefreshesDataSource{}
}

// RefreshesData extends the main model with additional fields.
type RefreshesData struct {
	DataQuality types.List `tfsdk:"refreshes"`

	PageSize   types.Int64  `tfsdk:"page_size"`
	ObjectType types.String `tfsdk:"object_type"`
	ObjectId   types.String `tfsdk:"object_id"`
}

func (RefreshesData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"refreshes": reflect.TypeOf(RefreshData{}),
	}
}

func (m RefreshesData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()

	attrs["refreshes"] = attrs["refreshes"].SetComputed()
	attrs["object_type"] = attrs["object_type"].SetRequired()
	attrs["object_id"] = attrs["object_id"].SetRequired()
	return attrs
}

type RefreshesDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *RefreshesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *RefreshesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, RefreshesData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Refresh",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *RefreshesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *RefreshesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config RefreshesData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest dataquality.ListRefreshRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.DataQuality.ListRefreshAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list data_quality_refreshes", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var refresh RefreshData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &refresh)...)
		if resp.Diagnostics.HasError() {
			return
		}
		results = append(results, refresh.ToObjectValue(ctx))
	}

	config.DataQuality = types.ListValueMust(RefreshData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
}
