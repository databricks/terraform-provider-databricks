// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package account_iam_external_group_v2

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

const dataSourceName = "account_iam_external_group_v2"

var _ datasource.DataSourceWithConfigure = &ExternalGroupDataSource{}

func DataSourceExternalGroup() datasource.DataSource {
	return &ExternalGroupDataSource{}
}

type ExternalGroupDataSource struct {
	Client *autogen.DatabricksClient
}

// ExternalGroupData extends the main model with additional fields.
type ExternalGroupData struct {
	// The parent account ID, from Databricks.
	AccountId types.String `tfsdk:"account_id"`
	// Display name of the group from the customer's IdP.
	DisplayName types.String `tfsdk:"display_name"`
	// The external ID of the group in the customer's IdP.
	ExternalGroupId types.String `tfsdk:"external_group_id"`
	// Internal groupId of the group in Databricks.
	InternalId types.String `tfsdk:"internal_id"`
	// The resource name of the external group. The format depends on the API
	// that returned it: - Account-scoped:
	// accounts/{account_id}/external-groups/{external_group_id} -
	// Workspace-scoped: external-groups/{external_group_id}
	Name types.String `tfsdk:"name"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// ExternalGroupData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m ExternalGroupData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalGroupData
// only implements ToObjectValue() and Type().
func (m ExternalGroupData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":        m.AccountId,
			"display_name":      m.DisplayName,
			"external_group_id": m.ExternalGroupId,
			"internal_id":       m.InternalId,
			"name":              m.Name,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m ExternalGroupData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":        types.StringType,
			"display_name":      types.StringType,
			"external_group_id": types.StringType,
			"internal_id":       types.StringType,
			"name":              types.StringType,
		},
	}
}

func (m ExternalGroupData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["display_name"] = attrs["display_name"].SetComputed()
	attrs["external_group_id"] = attrs["external_group_id"].SetComputed()
	attrs["internal_id"] = attrs["internal_id"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

func (r *ExternalGroupDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *ExternalGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, ExternalGroupData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks ExternalGroup",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *ExternalGroupDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *ExternalGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config ExternalGroupData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest iamv2.GetExternalGroupRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetAccountClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.IamV2.GetExternalGroup(ctx, readRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to get account_iam_external_group_v2", err.Error())
		return
	}

	var newState ExternalGroupData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
