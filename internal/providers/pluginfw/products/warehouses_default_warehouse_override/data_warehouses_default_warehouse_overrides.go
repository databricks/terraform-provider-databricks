// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package warehouses_default_warehouse_override

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/sql_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
)

const dataSourcesName = "warehouses_default_warehouse_overrides"

var _ datasource.DataSourceWithConfigure = &DefaultWarehouseOverridesDataSource{}

func DataSourceDefaultWarehouseOverrides() datasource.DataSource {
	return &DefaultWarehouseOverridesDataSource{}
}

// DefaultWarehouseOverridesData extends the main model with additional fields.
type DefaultWarehouseOverridesData struct {
	Warehouses types.List `tfsdk:"default_warehouse_overrides"`
    // The maximum number of overrides to return. The service may return fewer
    // than this value. If unspecified, at most 100 overrides will be returned.
    // The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize types.Int64 `tfsdk:"page_size"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
	
}

func (DefaultWarehouseOverridesData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"default_warehouse_overrides": reflect.TypeOf(DefaultWarehouseOverrideData{}),
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
		
	}
}

func (m DefaultWarehouseOverridesData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {attrs["page_size"] = attrs["page_size"].SetOptional()

	attrs["default_warehouse_overrides"] = attrs["default_warehouse_overrides"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	
	return attrs
}

type DefaultWarehouseOverridesDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *DefaultWarehouseOverridesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *DefaultWarehouseOverridesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, DefaultWarehouseOverridesData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks DefaultWarehouseOverride",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *DefaultWarehouseOverridesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *DefaultWarehouseOverridesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
    ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config DefaultWarehouseOverridesData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest sql.ListDefaultWarehouseOverridesRequest
    resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
    if resp.Diagnostics.HasError() {
        return
    }

	
	var namespace ProviderConfigData
	resp.Diagnostics.Append(config.ProviderConfigData.As(ctx, &namespace, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if resp.Diagnostics.HasError() {
		return
	}
	client, clientDiags := r.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, namespace.WorkspaceID.ValueString())
	
	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.Warehouses.ListDefaultWarehouseOverridesAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list warehouses_default_warehouse_overrides", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var default_warehouse_override DefaultWarehouseOverrideData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &default_warehouse_override)...)
		if resp.Diagnostics.HasError() {
			return
		}
		default_warehouse_override.ProviderConfigData = config.ProviderConfigData
		
		results = append(results, default_warehouse_override.ToObjectValue(ctx))
	}

	config.Warehouses = types.ListValueMust(DefaultWarehouseOverrideData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
}