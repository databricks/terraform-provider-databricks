// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package disaster_recovery_stable_url

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/disasterrecovery"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const dataSourcesName = "disaster_recovery_stable_urls"

var _ datasource.DataSourceWithConfigure = &StableUrlsDataSource{}

func DataSourceStableUrls() datasource.DataSource {
	return &StableUrlsDataSource{}
}

// StableUrlsData extends the main model with additional fields.
type StableUrlsData struct {
	DisasterRecovery types.List `tfsdk:"stable_urls"`
	// Maximum number of stable URLs to return per page. Default: 50, maximum:
	// 100.
	PageSize types.Int64 `tfsdk:"page_size"`
	// The parent resource. Format: accounts/{account_id}.
	Parent types.String `tfsdk:"parent"`
}

func (StableUrlsData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"stable_urls": reflect.TypeOf(StableUrlData{}),
	}
}

func (m StableUrlsData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	attrs["stable_urls"] = attrs["stable_urls"].SetComputed()
	return attrs
}

type StableUrlsDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *StableUrlsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *StableUrlsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, StableUrlsData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks StableUrl",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *StableUrlsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *StableUrlsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config StableUrlsData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest disasterrecovery.ListStableUrlsRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetAccountClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.DisasterRecovery.ListStableUrlsAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list disaster_recovery_stable_urls", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var stable_url StableUrlData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &stable_url)...)
		if resp.Diagnostics.HasError() {
			return
		}
		results = append(results, stable_url.ToObjectValue(ctx))
	}

	config.DisasterRecovery = types.ListValueMust(StableUrlData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
}
