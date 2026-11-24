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

const groupByExternalIdDataSourceName = "group_by_external_id"

var _ datasource.DataSourceWithConfigure = &GroupByExternalIdDataSource{}

func DataSourceGroupByExternalId() datasource.DataSource {
	return &GroupByExternalIdDataSource{}
}

type GroupByExternalIdDataSource struct {
	Client *common.DatabricksClient
}

// GroupByExternalIdData resolves a Databricks group from an external ID
// assigned by the customer's IdP, at either the account or workspace level.
type GroupByExternalIdData struct {
	iamv2_tf.Group
	// Whether to use the account or workspace API. Inferred from the
	// provider's host when unset; required on a unified host.
	Api types.String `tfsdk:"api"`
	tfschema.Namespace
}

func (GroupByExternalIdData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"provider_config": reflect.TypeOf(tfschema.ProviderConfigData{}),
	}
}

func (GroupByExternalIdData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["external_id"] = attrs["external_id"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["group_id"] = attrs["group_id"].SetComputed()
	attrs["group_name"] = attrs["group_name"].SetComputed()
	attrs["api"] = attrs["api"].SetOptional()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

func (d *GroupByExternalIdDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(groupByExternalIdDataSourceName)
}

func (d *GroupByExternalIdDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, GroupByExternalIdData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Resolves a Databricks group from an external ID assigned by the customer's IdP. " +
			"If no group with the given external ID exists yet, Databricks creates one as a side effect of this call.",
		Attributes: attrs,
		Blocks:     blocks,
	}
}

func (d *GroupByExternalIdDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if d.Client == nil {
		d.Client = pluginfwcommon.ConfigureDataSource(req, resp)
	}
}

func (d *GroupByExternalIdDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, groupByExternalIdDataSourceName)

	var data GroupByExternalIdData
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
	var group *iamv2.Group
	if isAccount {
		a, clientDiags := d.Client.GetAccountClient()
		resp.Diagnostics.Append(clientDiags...)
		if resp.Diagnostics.HasError() {
			return
		}
		response, err := a.IamV2.ResolveGroup(ctx, iamv2.ResolveGroupRequest{ExternalId: externalId})
		if err != nil {
			resp.Diagnostics.AddError("failed to resolve group by external id", err.Error())
			return
		}
		group = response.Group
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
		response, err := w.WorkspaceIamV2.ResolveGroupProxy(ctx, iamv2.ResolveGroupProxyRequest{ExternalId: externalId})
		if err != nil {
			resp.Diagnostics.AddError("failed to resolve group by external id", err.Error())
			return
		}
		group = response.Group
	}

	if group == nil {
		resp.Diagnostics.AddError("failed to resolve group by external id",
			fmt.Sprintf("no group was returned for external_id %q", externalId))
		return
	}

	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, group, &data)...)
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
