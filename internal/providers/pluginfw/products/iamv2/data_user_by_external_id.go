package iamv2

import (
	"context"
	"fmt"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/iamv2"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/iamv2_tf"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const userByExternalIdDataSourceName = "user_by_external_id"

var _ datasource.DataSourceWithConfigure = &UserByExternalIdDataSource{}

func DataSourceUserByExternalId() datasource.DataSource {
	return &UserByExternalIdDataSource{}
}

type UserByExternalIdDataSource struct {
	Client *common.DatabricksClient
}

// UserByExternalIdData resolves a Databricks user from an external ID
// assigned by the customer's IdP, at either the account or workspace level.
type UserByExternalIdData struct {
	iamv2_tf.User
	// Whether to use the account or workspace API. Inferred from the
	// provider's host when unset; required on a unified host.
	Api types.String `tfsdk:"api"`
	tfschema.Namespace
}

func (UserByExternalIdData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"full_name":       reflect.TypeOf(iamv2_tf.UserFullName{}),
		"provider_config": reflect.TypeOf(tfschema.ProviderConfigData{}),
	}
}

func (UserByExternalIdData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["external_id"] = attrs["external_id"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["account_user_status"] = attrs["account_user_status"].SetComputed()
	attrs["full_name"] = attrs["full_name"].SetComputed()
	attrs["user_id"] = attrs["user_id"].SetComputed()
	attrs["username"] = attrs["username"].SetComputed()
	attrs["api"] = attrs["api"].SetOptional()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

func (d *UserByExternalIdDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(userByExternalIdDataSourceName)
}

func (d *UserByExternalIdDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, UserByExternalIdData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Resolves a Databricks user from an external ID assigned by the customer's IdP. " +
			"If no user with the given external ID exists yet, Databricks creates one as a side effect of this call.",
		Attributes: attrs,
		Blocks:     blocks,
	}
}

func (d *UserByExternalIdDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if d.Client == nil {
		d.Client = pluginfwcommon.ConfigureDataSource(req, resp)
	}
}

func (d *UserByExternalIdDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, userByExternalIdDataSourceName)

	var data UserByExternalIdData
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	isAccount, diags := resolveApiLevel(d.Client, data.Api)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	externalId := data.ExternalId.ValueString()
	var user *iamv2.User
	if isAccount {
		a, clientDiags := d.Client.GetAccountClient()
		resp.Diagnostics.Append(clientDiags...)
		if resp.Diagnostics.HasError() {
			return
		}
		response, err := a.IamV2.ResolveUser(ctx, iamv2.ResolveUserRequest{ExternalId: externalId})
		if err != nil {
			resp.Diagnostics.AddError("failed to resolve user by external id", err.Error())
			return
		}
		user = response.User
	} else {
		workspaceID, wsDiags := tfschema.GetWorkspaceIDDataSource(ctx, data.ProviderConfig)
		resp.Diagnostics.Append(wsDiags...)
		if resp.Diagnostics.HasError() {
			return
		}
		w, clientDiags := d.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, workspaceID)
		resp.Diagnostics.Append(clientDiags...)
		if resp.Diagnostics.HasError() {
			return
		}
		response, err := w.WorkspaceIamV2.ResolveUserProxy(ctx, iamv2.ResolveUserProxyRequest{ExternalId: externalId})
		if err != nil {
			resp.Diagnostics.AddError("failed to resolve user by external id", err.Error())
			return
		}
		user = response.User
	}

	if user == nil {
		resp.Diagnostics.AddError("failed to resolve user by external id",
			fmt.Sprintf("no user was returned for external_id %q", externalId))
		return
	}

	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, user, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !isAccount {
		resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, d.Client, data.ProviderConfig, &resp.State)...)
	}
}
