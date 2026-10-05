// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package mason_session

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

const dataSourcesName = "mason_sessions"

var _ datasource.DataSourceWithConfigure = &SessionsDataSource{}

func DataSourceSessions() datasource.DataSource {
	return &SessionsDataSource{}
}

// SessionsData extends the main model with additional fields.
type SessionsData struct {
	Mason types.List `tfsdk:"sessions"`
	// Filter expression. Supported fields include `actor_id` and `metadata`;
	// for example, `actor_id = "support-customer-123"`.
	Filter types.String `tfsdk:"filter"`
	// Sort order. Defaults to `last_activity_time desc`. Page-token
	// continuation is exactly-once when ordering by `create_time` (immutable);
	// ordering by `last_activity_time` is best-effort, because that value
	// changes as a session gains activity, so a session updated between page
	// requests may be repeated or skipped. To enumerate every session exactly
	// once, order by `create_time`.
	OrderBy types.String `tfsdk:"order_by"`
	// Maximum number of sessions to return. Defaults to 10; must be between 1
	// and 100.
	PageSize types.Int64 `tfsdk:"page_size"`
	// Resource name of the containing session store, in the form
	// `session-stores/{session_store_id}`.
	Parent             types.String `tfsdk:"parent"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (SessionsData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"sessions":        reflect.TypeOf(SessionData{}),
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m SessionsData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["filter"] = attrs["filter"].SetOptional()
	attrs["order_by"] = attrs["order_by"].SetOptional()

	attrs["sessions"] = attrs["sessions"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type SessionsDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *SessionsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *SessionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, SessionsData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Session",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *SessionsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *SessionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config SessionsData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest mason.ListSessionsRequest
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

	response, err := client.Mason.ListSessionsAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list mason_sessions", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var session SessionData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &session)...)
		if resp.Diagnostics.HasError() {
			return
		}
		session.ProviderConfigData = config.ProviderConfigData

		results = append(results, session.ToObjectValue(ctx))
	}

	config.Mason = types.ListValueMust(SessionData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
