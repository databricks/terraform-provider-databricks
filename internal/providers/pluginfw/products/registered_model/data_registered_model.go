package registered_model

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
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceName = "registered_model"

func DataSourceRegisteredModel() datasource.DataSource {
	return &RegisteredModelDataSource{}
}

var _ datasource.DataSourceWithConfigure = &RegisteredModelDataSource{}

type RegisteredModelDataSource struct {
	Client *common.DatabricksClient
}

type RegisteredModelData struct {
	FullName       types.String `tfsdk:"full_name"`
	IncludeAliases types.Bool   `tfsdk:"include_aliases"`
	IncludeBrowse  types.Bool   `tfsdk:"include_browse"`
	ModelInfo      types.List   `tfsdk:"model_info"`
	tfschema.Namespace
}

func (RegisteredModelData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["full_name"] = attrs["full_name"].SetRequired()
	attrs["include_aliases"] = attrs["include_aliases"].SetOptional()
	attrs["include_browse"] = attrs["include_browse"].SetOptional()
	attrs["model_info"] = attrs["model_info"].SetOptional().SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	return attrs
}

func (RegisteredModelData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model_info":      reflect.TypeOf(catalog_tf.RegisteredModelInfo_SdkV2{}),
		"provider_config": reflect.TypeOf(tfschema.ProviderConfigData{}),
	}
}

func (d *RegisteredModelDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(dataSourceName)
}

func (d *RegisteredModelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, RegisteredModelData{}, nil)
	resp.Schema = schema.Schema{
		Attributes: attrs,
		Blocks:     blocks,
	}
}

func (d *RegisteredModelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if d.Client == nil {
		d.Client = pluginfwcommon.ConfigureDataSource(req, resp)
	}
}

func (d *RegisteredModelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var registeredModel RegisteredModelData
	diags := req.Config.Get(ctx, &registeredModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	workspaceID, diags := tfschema.GetWorkspaceIDDataSource(ctx, registeredModel.ProviderConfig)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	w, diags := d.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, workspaceID)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	modelFullName := registeredModel.FullName.ValueString()
	modelInfoSdk, err := w.RegisteredModels.Get(ctx, catalog.GetRegisteredModelRequest{
		FullName:       modelFullName,
		IncludeAliases: registeredModel.IncludeAliases.ValueBool(),
		IncludeBrowse:  registeredModel.IncludeBrowse.ValueBool(),
	})
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
		}
		resp.Diagnostics.AddError(fmt.Sprintf("failed to get registered model %s", modelFullName), err.Error())
		return
	}
	var modelInfo catalog_tf.RegisteredModelInfo
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, modelInfoSdk, &modelInfo)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if modelInfo.Aliases.IsNull() {
		var d diag.Diagnostics
		modelInfo.Aliases, d = basetypes.NewListValueFrom(ctx, modelInfo.Aliases.ElementType(ctx), []catalog_tf.RegisteredModelAlias{})
		resp.Diagnostics.Append(d...)
	}
	registeredModel.ModelInfo = types.ListValueMust(catalog_tf.RegisteredModelInfo{}.Type(ctx), []attr.Value{modelInfo.ToObjectValue(ctx)})
	resp.Diagnostics.Append(resp.State.Set(ctx, registeredModel)...)
}
