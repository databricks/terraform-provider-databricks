package sharing

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceNameShares = "shares"

type SharesList struct {
	Shares types.List `tfsdk:"shares"`
	tfschema.Namespace
}

func (s SharesList) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"shares":          reflect.TypeOf(types.String{}),
		"provider_config": reflect.TypeOf(tfschema.ProviderConfigData{}),
	}
}

func (s SharesList) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"shares":          types.ListType{ElemType: types.StringType},
			"provider_config": tfschema.ProviderConfigData{}.Type(ctx),
		},
	}
}

func (s SharesList) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["shares"] = attrs["shares"].SetComputed().SetOptional()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	return attrs
}

func DataSourceShares() datasource.DataSource {
	return &SharesDataSource{}
}

var _ datasource.DataSourceWithConfigure = &SharesDataSource{}

type SharesDataSource struct {
	Client *common.DatabricksClient
}

func (d *SharesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(dataSourceNameShares)
}

func (d *SharesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, SharesList{}, nil)
	resp.Schema = schema.Schema{
		Attributes: attrs,
		Blocks:     blocks,
	}
}

func (d *SharesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if d.Client == nil {
		d.Client = pluginfwcommon.ConfigureDataSource(req, resp)
	}
}

func (d *SharesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceNameShares)

	var config SharesList
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var workspaceID string
	if !config.ProviderConfig.IsNull() {
		var namespace tfschema.ProviderConfigData
		resp.Diagnostics.Append(config.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
			UnhandledNullAsEmpty:    true,
			UnhandledUnknownAsEmpty: true,
		})...)
		if resp.Diagnostics.HasError() {
			return
		}
		workspaceID = namespace.WorkspaceID.ValueString()
	}
	w, clientDiags := d.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, workspaceID)

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	shares, err := w.Shares.ListAll(ctx, sharing.ListSharesRequest{})
	if err != nil {
		resp.Diagnostics.AddError("Failed to fetch shares", err.Error())
		return
	}

	shareNames := make([]attr.Value, len(shares))
	for i, share := range shares {
		shareNames[i] = types.StringValue(share.Name)
	}

	newState := SharesList{
		Shares: types.ListValueMust(types.StringType, shareNames),
		Namespace: tfschema.Namespace{
			ProviderConfig: config.ProviderConfig,
		},
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
