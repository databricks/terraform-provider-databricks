// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package domain

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/domains"
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

const dataSourcesName = "domains"

var _ datasource.DataSourceWithConfigure = &DomainsDataSource{}

func DataSourceDomains() datasource.DataSource {
	return &DomainsDataSource{}
}

// DomainsData extends the main model with additional fields.
type DomainsData struct {
	Domains types.List `tfsdk:"domains"`

	PageSize types.Int64 `tfsdk:"page_size"`
	// Filter by parent domain. - Absent: return all domains regardless of
	// hierarchy. - Present: return only direct children of the specified
	// domain.
	ParentDomainId     types.String `tfsdk:"parent_domain_id"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (DomainsData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"domains":         reflect.TypeOf(DomainData{}),
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

func (m DomainsData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["parent_domain_id"] = attrs["parent_domain_id"].SetOptional()

	attrs["domains"] = attrs["domains"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

type DomainsDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *DomainsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *DomainsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, DomainsData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Domain",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *DomainsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *DomainsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config DomainsData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest domains.ListDomainsRequest
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

	response, err := client.Domains.ListDomainsAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list domains", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var domain DomainData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &domain)...)
		if resp.Diagnostics.HasError() {
			return
		}
		domain.ProviderConfigData = config.ProviderConfigData

		results = append(results, domain.ToObjectValue(ctx))
	}

	config.Domains = types.ListValueMust(DomainData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
