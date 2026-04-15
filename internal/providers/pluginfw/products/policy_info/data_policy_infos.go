// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package policy_info

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/catalog_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
)

const dataSourcesName = "policy_infos"

var _ datasource.DataSourceWithConfigure = &PolicyInfosDataSource{}

func DataSourcePolicyInfos() datasource.DataSource {
	return &PolicyInfosDataSource{}
}

// PolicyInfosData extends the main model with additional fields.
type PolicyInfosData struct {
	Policies types.List `tfsdk:"policies"`
    // Optional. Whether to include policies defined on parent securables. By
    // default, the inherited policies are not included.
	IncludeInherited types.Bool `tfsdk:"include_inherited"`
    // Optional. Maximum number of policies to return on a single page (page
    // length). - When not set or set to 0, the page length is set to a server
    // configured value (recommended); - When set to a value greater than 0, the
    // page length is the minimum of this value and a server configured value;
	MaxResults types.Int64 `tfsdk:"max_results"`
	OnSecurableType types.String `tfsdk:"on_securable_type"`
	OnSecurableFullname types.String `tfsdk:"on_securable_fullname"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
	
}

func (PolicyInfosData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"policies": reflect.TypeOf(PolicyInfoData{}),
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
		
	}
}

func (m PolicyInfosData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {attrs["include_inherited"] = attrs["include_inherited"].SetOptional()
attrs["max_results"] = attrs["max_results"].SetOptional()

	attrs["policies"] = attrs["policies"].SetComputed()
	attrs["on_securable_type"] = attrs["on_securable_type"].SetRequired()
	attrs["on_securable_fullname"] = attrs["on_securable_fullname"].SetRequired()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	
	return attrs
}

type PolicyInfosDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *PolicyInfosDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *PolicyInfosDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, PolicyInfosData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks PolicyInfo",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *PolicyInfosDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *PolicyInfosDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
    ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config PolicyInfosData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest catalog.ListPoliciesRequest
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

	response, err := client.Policies.ListPoliciesAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list policy_infos", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var policy_info PolicyInfoData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &policy_info)...)
		if resp.Diagnostics.HasError() {
			return
		}
		policy_info.ProviderConfigData = config.ProviderConfigData
		
		results = append(results, policy_info.ToObjectValue(ctx))
	}

	config.Policies = types.ListValueMust(PolicyInfoData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
}