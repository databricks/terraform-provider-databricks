package registered_model

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/catalog_tf"
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
	FullName      types.String                  `tfsdk:"full_name"`
	ModelVersions []catalog_tf.ModelVersionInfo `tfsdk:"model_versions" tf:"optional,computed"`
}

func (d *RegisteredModelVersionsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "databricks_registered_model_versions"
}

func (d *RegisteredModelVersionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(RegisteredModelVersionsData{}, nil)
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
	for _, modelVersionSdk := range modelVersions.ModelVersions {
		var modelVersion catalog_tf.ModelVersionInfo
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, modelVersionSdk, &modelVersion)...)
		if resp.Diagnostics.HasError() {
			return
		}
		registeredModelVersions.ModelVersions = append(registeredModelVersions.ModelVersions, modelVersion)
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, registeredModelVersions)...)
}
