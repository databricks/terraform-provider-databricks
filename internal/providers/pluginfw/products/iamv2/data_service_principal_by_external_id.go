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

const servicePrincipalByExternalIdDataSourceName = "service_principal_by_external_id"

var _ datasource.DataSourceWithConfigure = &ServicePrincipalByExternalIdDataSource{}

func DataSourceServicePrincipalByExternalId() datasource.DataSource {
	return &ServicePrincipalByExternalIdDataSource{}
}

type ServicePrincipalByExternalIdDataSource struct {
	Client *common.DatabricksClient
}

// ServicePrincipalByExternalIdData resolves a Databricks service principal
// from an external ID assigned by the customer's IdP, at either the account
// or workspace level.
type ServicePrincipalByExternalIdData struct {
	iamv2_tf.ServicePrincipal
	// Whether to use the account or workspace API. Inferred from the
	// provider's host when unset; required on a unified host.
	Api types.String `tfsdk:"api"`
	tfschema.Namespace
}

func (ServicePrincipalByExternalIdData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"provider_config": reflect.TypeOf(tfschema.ProviderConfigData{}),
	}
}

func (ServicePrincipalByExternalIdData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["external_id"] = attrs["external_id"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["account_sp_status"] = attrs["account_sp_status"].SetComputed()
	attrs["application_id"] = attrs["application_id"].SetComputed()
	attrs["display_name"] = attrs["display_name"].SetComputed()
	attrs["service_principal_id"] = attrs["service_principal_id"].SetComputed()
	attrs["api"] = attrs["api"].SetOptional()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

func (d *ServicePrincipalByExternalIdDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(servicePrincipalByExternalIdDataSourceName)
}

func (d *ServicePrincipalByExternalIdDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, ServicePrincipalByExternalIdData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Resolves a Databricks service principal from an external ID assigned by the customer's IdP. " +
			"If no service principal with the given external ID exists yet, Databricks creates one as a side effect of this call.",
		Attributes: attrs,
		Blocks:     blocks,
	}
}

func (d *ServicePrincipalByExternalIdDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if d.Client == nil {
		d.Client = pluginfwcommon.ConfigureDataSource(req, resp)
	}
}

func (d *ServicePrincipalByExternalIdDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, servicePrincipalByExternalIdDataSourceName)

	var data ServicePrincipalByExternalIdData
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
	var servicePrincipal *iamv2.ServicePrincipal
	if isAccount {
		a, clientDiags := d.Client.GetAccountClient()
		resp.Diagnostics.Append(clientDiags...)
		if resp.Diagnostics.HasError() {
			return
		}
		response, err := a.IamV2.ResolveServicePrincipal(ctx, iamv2.ResolveServicePrincipalRequest{ExternalId: externalId})
		if err != nil {
			resp.Diagnostics.AddError("failed to resolve service principal by external id", err.Error())
			return
		}
		servicePrincipal = response.ServicePrincipal
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
		response, err := w.WorkspaceIamV2.ResolveServicePrincipalProxy(ctx, iamv2.ResolveServicePrincipalProxyRequest{ExternalId: externalId})
		if err != nil {
			resp.Diagnostics.AddError("failed to resolve service principal by external id", err.Error())
			return
		}
		servicePrincipal = response.ServicePrincipal
	}

	if servicePrincipal == nil {
		resp.Diagnostics.AddError("failed to resolve service principal by external id",
			fmt.Sprintf("no service principal was returned for external_id %q", externalId))
		return
	}

	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, servicePrincipal, &data)...)
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
