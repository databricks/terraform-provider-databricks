package catalog

import (
	"context"
	"fmt"
	"reflect"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/catalog_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceName = "functions"

func DataSourceFunctions() datasource.DataSource {
	return &FunctionsDataSource{}
}

var _ datasource.DataSourceWithConfigure = &FunctionsDataSource{}

type FunctionsDataSource struct {
	Client *common.DatabricksClient
}

type FunctionsData struct {
	CatalogName   types.String `tfsdk:"catalog_name"`
	SchemaName    types.String `tfsdk:"schema_name"`
	IncludeBrowse types.Bool   `tfsdk:"include_browse"`
	Functions     types.List   `tfsdk:"functions"`
	tfschema.Namespace
}

func (FunctionsData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["catalog_name"] = attrs["catalog_name"].SetRequired()
	attrs["schema_name"] = attrs["schema_name"].SetRequired()
	attrs["include_browse"] = attrs["include_browse"].SetOptional()
	attrs["functions"] = attrs["functions"].SetOptional().SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	return attrs
}

func (FunctionsData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"functions":       reflect.TypeOf(catalog_tf.FunctionInfo{}),
		"provider_config": reflect.TypeOf(tfschema.ProviderConfigData{}),
	}
}

func (d *FunctionsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(dataSourceName)
}

func (d *FunctionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, FunctionsData{}, nil)
	resp.Schema = schema.Schema{
		Attributes: attrs,
		Blocks:     blocks,
	}
}

func (d *FunctionsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if d.Client == nil {
		d.Client = pluginfwcommon.ConfigureDataSource(req, resp)
	}
}

func (d *FunctionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var functions FunctionsData
	diags := req.Config.Get(ctx, &functions)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var workspaceID string
	if !functions.ProviderConfig.IsNull() {
		var namespace tfschema.ProviderConfigData
		resp.Diagnostics.Append(functions.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
			UnhandledNullAsEmpty:    true,
			UnhandledUnknownAsEmpty: true,
		})...)
		if resp.Diagnostics.HasError() {
			return
		}
		workspaceID = namespace.WorkspaceID.ValueString()
	}

	w, diags := d.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, workspaceID)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	catalogName := functions.CatalogName.ValueString()
	schemaName := functions.SchemaName.ValueString()
	functionsInfosSdk, err := w.Functions.ListAll(ctx, catalog.ListFunctionsRequest{
		CatalogName:   catalogName,
		SchemaName:    schemaName,
		IncludeBrowse: functions.IncludeBrowse.ValueBool(),
	})
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
		}
		resp.Diagnostics.AddError(fmt.Sprintf("failed to get functions for %s.%s schema", catalogName, schemaName), err.Error())
		return
	}
	tfFunctions := []attr.Value{}
	for _, functionSdk := range functionsInfosSdk {
		var function catalog_tf.FunctionInfo
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, functionSdk, &function)...)
		if resp.Diagnostics.HasError() {
			return
		}
		tfFunctions = append(tfFunctions, function.ToObjectValue(ctx))
	}
	functions.Functions = types.ListValueMust(catalog_tf.FunctionInfo{}.Type(ctx), tfFunctions)
	resp.Diagnostics.Append(resp.State.Set(ctx, functions)...)
}
