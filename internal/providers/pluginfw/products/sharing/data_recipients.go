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
)

const dataSourceNameRecipients = "recipients"

type RecipientsList struct {
	Recipients types.List `tfsdk:"recipients"`
	tfschema.Namespace
}

func (s RecipientsList) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"recipients":      reflect.TypeOf(types.String{}),
		"provider_config": reflect.TypeOf(tfschema.ProviderConfigData{}),
	}
}

func (s RecipientsList) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"recipients":      types.ListType{ElemType: types.StringType},
			"provider_config": tfschema.ProviderConfigData{}.Type(ctx),
		},
	}
}

func (s RecipientsList) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["recipients"] = attrs["recipients"].SetComputed().SetOptional()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	return attrs
}

func DataSourceRecipients() datasource.DataSource {
	return &RecipientsDataSource{}
}

var _ datasource.DataSourceWithConfigure = &RecipientsDataSource{}

type RecipientsDataSource struct {
	Client *common.DatabricksClient
}

func (d *RecipientsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(dataSourceNameRecipients)
}

func (d *RecipientsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, RecipientsList{}, nil)
	resp.Schema = schema.Schema{
		Attributes: attrs,
		Blocks:     blocks,
	}
}

func (d *RecipientsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if d.Client == nil {
		d.Client = pluginfwcommon.ConfigureDataSource(req, resp)
	}
}

func (d *RecipientsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceNameRecipients)

	var config RecipientsList
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	workspaceID, diags := tfschema.GetWorkspaceIDDataSource(ctx, config.ProviderConfig)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	w, clientDiags := d.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, workspaceID)

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	recipients, err := w.Recipients.ListAll(ctx, sharing.ListRecipientsRequest{})
	if err != nil {
		resp.Diagnostics.AddError("Failed to fetch recipients", err.Error())
		return
	}

	recipientNames := make([]attr.Value, len(recipients))
	for i, recipient := range recipients {
		recipientNames[i] = types.StringValue(recipient.Name)
	}

	newState := RecipientsList{
		Recipients: types.ListValueMust(types.StringType, recipientNames),
		Namespace: tfschema.Namespace{
			ProviderConfig: config.ProviderConfig,
		},
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
