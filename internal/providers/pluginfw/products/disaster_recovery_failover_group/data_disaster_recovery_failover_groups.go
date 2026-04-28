// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package disaster_recovery_failover_group

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

const dataSourcesName = "disaster_recovery_failover_groups"

var _ datasource.DataSourceWithConfigure = &FailoverGroupsDataSource{}

func DataSourceFailoverGroups() datasource.DataSource {
	return &FailoverGroupsDataSource{}
}

// FailoverGroupsData extends the main model with additional fields.
type FailoverGroupsData struct {
	DisasterRecovery types.List `tfsdk:"failover_groups"`
	// Maximum number of failover groups to return per page. Default: 50,
	// maximum: 100.
	PageSize types.Int64 `tfsdk:"page_size"`
	// The parent resource. Format: accounts/{account_id}.
	Parent types.String `tfsdk:"parent"`
}

func (FailoverGroupsData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"failover_groups": reflect.TypeOf(FailoverGroupData{}),
	}
}

func (m FailoverGroupsData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	attrs["failover_groups"] = attrs["failover_groups"].SetComputed()
	return attrs
}

type FailoverGroupsDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *FailoverGroupsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *FailoverGroupsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, FailoverGroupsData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks FailoverGroup",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *FailoverGroupsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *FailoverGroupsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config FailoverGroupsData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest disasterrecovery.ListFailoverGroupsRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetAccountClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.DisasterRecovery.ListFailoverGroupsAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list disaster_recovery_failover_groups", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var failover_group FailoverGroupData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &failover_group)...)
		if resp.Diagnostics.HasError() {
			return
		}
		results = append(results, failover_group.ToObjectValue(ctx))
	}

	config.DisasterRecovery = types.ListValueMust(FailoverGroupData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
}
