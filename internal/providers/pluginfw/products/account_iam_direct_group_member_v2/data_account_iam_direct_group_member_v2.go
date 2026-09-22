// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package account_iam_direct_group_member_v2

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/iamv2"
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

const dataSourceName = "account_iam_direct_group_member_v2"

var _ datasource.DataSourceWithConfigure = &DirectGroupMemberDataSource{}

func DataSourceDirectGroupMember() datasource.DataSource {
	return &DirectGroupMemberDataSource{}
}

type DirectGroupMemberDataSource struct {
	Client *autogen.DatabricksClient
}

// DirectGroupMemberData extends the main model with additional fields.
type DirectGroupMemberData struct {
	// Display name of the principal.
	DisplayName types.String `tfsdk:"display_name"`
	// The external ID of the principal in Databricks.
	ExternalId types.String `tfsdk:"external_id"`
	// The internal ID of the group this member belongs to.
	GroupId types.Int64 `tfsdk:"group_id"`
	// The source of group membership (internal or from identity provider).
	MembershipSource types.String `tfsdk:"membership_source"`
	// Internal ID of the principal in Databricks.
	PrincipalId types.Int64 `tfsdk:"principal_id"`
	// The type of the principal (user/service principal/group).
	PrincipalType types.String `tfsdk:"principal_type"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// DirectGroupMemberData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m DirectGroupMemberData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DirectGroupMemberData
// only implements ToObjectValue() and Type().
func (m DirectGroupMemberData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"display_name":      m.DisplayName,
			"external_id":       m.ExternalId,
			"group_id":          m.GroupId,
			"membership_source": m.MembershipSource,
			"principal_id":      m.PrincipalId,
			"principal_type":    m.PrincipalType,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m DirectGroupMemberData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"display_name":      types.StringType,
			"external_id":       types.StringType,
			"group_id":          types.Int64Type,
			"membership_source": types.StringType,
			"principal_id":      types.Int64Type,
			"principal_type":    types.StringType,
		},
	}
}

func (m DirectGroupMemberData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["display_name"] = attrs["display_name"].SetComputed()
	attrs["external_id"] = attrs["external_id"].SetComputed()
	attrs["group_id"] = attrs["group_id"].SetRequired()
	attrs["membership_source"] = attrs["membership_source"].SetComputed()
	attrs["principal_id"] = attrs["principal_id"].SetRequired()
	attrs["principal_type"] = attrs["principal_type"].SetComputed()

	return attrs
}

func (r *DirectGroupMemberDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *DirectGroupMemberDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, DirectGroupMemberData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks DirectGroupMember",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *DirectGroupMemberDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *DirectGroupMemberDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config DirectGroupMemberData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest iamv2.GetDirectGroupMemberRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetAccountClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.IamV2.GetDirectGroupMember(ctx, readRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to get account_iam_direct_group_member_v2", err.Error())
		return
	}

	var newState DirectGroupMemberData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
