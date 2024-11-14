package catalog

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

func DataSourceFunctions() datasource.DataSource {
	return &FunctionsDataSource{}
}

var _ datasource.DataSourceWithConfigure = &FunctionsDataSource{}

type FunctionsDataSource struct {
	Client *common.DatabricksClient
}

type FunctionsData struct {
	CatalogName   types.String              `tfsdk:"catalog_name"`
	SchemaName    types.String              `tfsdk:"schema_name"`
	IncludeBrowse types.Bool                `tfsdk:"include_browse" tf:"optional"`
	Functions     []catalog_tf.FunctionInfo `tfsdk:"functions" tf:"optional,computed"`
}

func (d *FunctionsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "databricks_functions"
}

func (d *FunctionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(FunctionsData{}, nil)
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
	w, diags := d.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var functions FunctionsData
	diags = req.Config.Get(ctx, &functions)
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
	for _, functionSdk := range functionsInfosSdk {
		var function catalog_tf.FunctionInfo
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, functionSdk, &function)...)
		if resp.Diagnostics.HasError() {
			return
		}
		functions.Functions = append(functions.Functions, function)
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, functions)...)
}
