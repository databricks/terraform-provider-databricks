package registered_model

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/catalog_tf"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func DataSourceRegisteredModel() datasource.DataSource {
	return &RegisteredModelDataSource{}
}

var _ datasource.DataSourceWithConfigure = &RegisteredModelDataSource{}

type RegisteredModelDataSource struct {
	Client *common.DatabricksClient
}

type RegisteredModelData struct {
	FullName             types.String                    `tfsdk:"full_name"`
	IncludeModelVersions types.Bool                      `tfsdk:"include_model_versions" tf:"optional"`
	ModelInfo            *catalog_tf.RegisteredModelInfo `tfsdk:"model_info" tf:"optional,computed"`
	ModelVersions        []catalog_tf.ModelVersionInfo   `tfsdk:"model_versions" tf:"optional,computed"`
}

func (d *RegisteredModelDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "databricks_registered_model"
}

func (d *RegisteredModelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: tfschema.DataSourceStructToSchemaMap(RegisteredModelData{}, nil),
	}
}

func (d *RegisteredModelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if d.Client == nil {
		d.Client = pluginfwcommon.ConfigureDataSource(req, resp)
	}
}

func (d *RegisteredModelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	w, diags := d.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var registeredModel RegisteredModelData
	diags = req.Config.Get(ctx, &registeredModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	modelFullName := registeredModel.FullName.ValueString()
	modelInfoSdk, err := w.RegisteredModels.Get(ctx, catalog.GetRegisteredModelRequest{
		FullName:       modelFullName,
		IncludeAliases: true,
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
	registeredModel.ModelInfo = &modelInfo
	registeredModel.ModelVersions = make([]catalog_tf.ModelVersionInfo, 0)
	if registeredModel.IncludeModelVersions.ValueBool() {
		modelVersions, err := w.ModelVersions.ListByFullName(ctx, modelFullName)
		if err != nil {
			resp.Diagnostics.AddError(fmt.Sprintf("failed to list model versions for registered model %s", modelFullName), err.Error())
			return
		}
		for _, modelVersionSdk := range modelVersions.ModelVersions {
			var modelVersion catalog_tf.ModelVersionInfo
			resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, modelVersionSdk, &modelVersion)...)
			if resp.Diagnostics.HasError() {
				return
			}
			registeredModel.ModelVersions = append(registeredModel.ModelVersions, modelVersion)
		}
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, registeredModel)...)
}
