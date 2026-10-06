// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package mason_managed_memory_entry

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

const dataSourcesName = "mason_managed_memory_entries"

var _ datasource.DataSourceWithConfigure = &ManagedMemoryEntriesDataSource{}

func DataSourceManagedMemoryEntries() datasource.DataSource {
	return &ManagedMemoryEntriesDataSource{}
}

// ManagedMemoryEntriesData extends the main model with additional fields.
type ManagedMemoryEntriesData struct {
	Mason types.List `tfsdk:"managed_memory_entries"`
	// Customer-provided identifier for the actor whose entries are listed.
	ActorId types.String `tfsdk:"actor_id"`
	// Maximum number of entries to return. The service may return fewer entries
	// than requested. Defaults to 10; must be between 1 and 100.
	PageSize types.Int64 `tfsdk:"page_size"`
	// Managed memory store whose entries are listed, in the form
	// `memory-stores/{managed_memory_store_id}`.
	Parent types.String `tfsdk:"parent"`
	// Optional path prefix used to restrict entries within the actor partition.
	PathPrefix types.String `tfsdk:"path_prefix"`
	// Fields to return in each entry, using proto field names such as `content`
	// (not `contents`). An omitted or empty mask returns each full entry,
	// including `content`; a non-empty mask returns only the requested fields.
	ReadMask types.String `tfsdk:"read_mask"`
	// Optional session identifier. When set, only entries with this exact
	// `session_id` are returned. Omitted-session (cross-session) entries are
	// not included. Ignored when path is set.
	SessionId          types.String `tfsdk:"session_id"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (ManagedMemoryEntriesData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"managed_memory_entries": reflect.TypeOf(ManagedMemoryEntryData{}),
		"provider_config":        reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m ManagedMemoryEntriesData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["actor_id"] = attrs["actor_id"].SetRequired()
	attrs["path_prefix"] = attrs["path_prefix"].SetOptional()
	attrs["session_id"] = attrs["session_id"].SetOptional()
	attrs["read_mask"] = attrs["read_mask"].SetOptional()

	attrs["managed_memory_entries"] = attrs["managed_memory_entries"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type ManagedMemoryEntriesDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *ManagedMemoryEntriesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *ManagedMemoryEntriesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, ManagedMemoryEntriesData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks ManagedMemoryEntry",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *ManagedMemoryEntriesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *ManagedMemoryEntriesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config ManagedMemoryEntriesData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest mason.ListManagedMemoryEntriesRequest
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

	response, err := client.Mason.ListMemoriesAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list mason_managed_memory_entries", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var managed_memory_entry ManagedMemoryEntryData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &managed_memory_entry)...)
		if resp.Diagnostics.HasError() {
			return
		}
		managed_memory_entry.ProviderConfigData = config.ProviderConfigData

		results = append(results, managed_memory_entry.ToObjectValue(ctx))
	}

	config.Mason = types.ListValueMust(ManagedMemoryEntryData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
