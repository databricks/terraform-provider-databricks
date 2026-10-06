// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package mason_managed_memory_store

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/mason"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourcesName = "mason_managed_memory_stores"

var _ datasource.DataSourceWithConfigure = &ManagedMemoryStoresDataSource{}

func DataSourceManagedMemoryStores() datasource.DataSource {
	return &ManagedMemoryStoresDataSource{}
}

// ManagedMemoryStoresData extends the main model with additional fields.
type ManagedMemoryStoresData struct {
	Mason types.List `tfsdk:"managed_memory_stores"`
	// Maximum number of stores to return. The service may return fewer stores
	// than requested. Defaults to 10; must be between 1 and 100.
	PageSize           types.Int64  `tfsdk:"page_size"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (ManagedMemoryStoresData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"managed_memory_stores": reflect.TypeOf(ManagedMemoryStoreData{}),
		"provider_config":       reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m ManagedMemoryStoresData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()

	attrs["managed_memory_stores"] = attrs["managed_memory_stores"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type ManagedMemoryStoresDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *ManagedMemoryStoresDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *ManagedMemoryStoresDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, ManagedMemoryStoresData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks ManagedMemoryStore",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *ManagedMemoryStoresDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *ManagedMemoryStoresDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config ManagedMemoryStoresData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest mason.ListManagedMemoryStoresRequest
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

	response, err := client.Mason.ListMemoryStoresAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list mason_managed_memory_stores", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var managed_memory_store ManagedMemoryStoreData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &managed_memory_store)...)
		if resp.Diagnostics.HasError() {
			return
		}
		managed_memory_store.ProviderConfigData = config.ProviderConfigData

		results = append(results, managed_memory_store.ToObjectValue(ctx))
	}

	config.Mason = types.ListValueMust(ManagedMemoryStoreData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
