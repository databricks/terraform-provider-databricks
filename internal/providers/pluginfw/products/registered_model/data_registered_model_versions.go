package registered_model

import (
	"context"
	"fmt"
	"reflect"

	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/catalog_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func DataSourceRegisteredModelVersions() datasource.DataSource {
	return &RegisteredModelVersionsDataSource{}
}

var _ datasource.DataSourceWithConfigure = &RegisteredModelVersionsDataSource{}

type RegisteredModelVersionsDataSource struct {
	Client *common.DatabricksClient
}

type RegisteredModelVersionsData struct {
	FullName      types.String `tfsdk:"full_name"`
	ModelVersions types.List   `tfsdk:"model_versions"`
}

func (RegisteredModelVersionsData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["full_name"] = attrs["full_name"].SetRequired()
	attrs["model_versions"] = attrs["model_versions"].SetOptional().SetComputed()
	return attrs
}

func (RegisteredModelVersionsData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model_versions": reflect.TypeOf(catalog_tf.ModelVersionInfo_SdkV2{}),
	}
}

func (d *RegisteredModelVersionsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "databricks_registered_model_versions"
}

func (d *RegisteredModelVersionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, RegisteredModelVersionsData{}, nil)
	resp.Schema = schema.Schema{
		Attributes: attrs,
		Blocks:     blocks,
	}
}

func (d *RegisteredModelVersionsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if d.Client == nil {
		d.Client = pluginfwcommon.ConfigureDataSource(req, resp)
	}
}

func (d *RegisteredModelVersionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	w, diags := d.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var registeredModelVersions RegisteredModelVersionsData
	diags = req.Config.Get(ctx, &registeredModelVersions)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	modelFullName := registeredModelVersions.FullName.ValueString()
	modelVersions, err := w.ModelVersions.ListByFullName(ctx, modelFullName)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("failed to list model versions for registered model %s", modelFullName), err.Error())
		return
	}
	var tfModelVersions []attr.Value
	for _, modelVersionSdk := range modelVersions.ModelVersions {
		var modelVersion catalog_tf.ModelVersionInfo_SdkV2
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, modelVersionSdk, &modelVersion)...)
		if resp.Diagnostics.HasError() {
			return
		}
		tfModelVersions = append(tfModelVersions, modelVersion.ToObjectValue(ctx))
	}
	registeredModelVersions.ModelVersions = types.ListValueMust(catalog_tf.ModelVersionInfo_SdkV2{}.Type(ctx), tfModelVersions)
	resp.Diagnostics.Append(resp.State.Set(ctx, registeredModelVersions)...)
}
