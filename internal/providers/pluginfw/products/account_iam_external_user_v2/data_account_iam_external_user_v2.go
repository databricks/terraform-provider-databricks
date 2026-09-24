// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package account_iam_external_user_v2

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/iamv2"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/iamv2_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceName = "account_iam_external_user_v2"

var _ datasource.DataSourceWithConfigure = &ExternalUserDataSource{}

func DataSourceExternalUser() datasource.DataSource {
	return &ExternalUserDataSource{}
}

type ExternalUserDataSource struct {
	Client *autogen.DatabricksClient
}

// ExternalUserData extends the main model with additional fields.
type ExternalUserData struct {
	// The parent account ID, from Databricks.
	AccountId types.String `tfsdk:"account_id"`
	// The activity status of the user in the Databricks account.
	AccountUserStatus types.String `tfsdk:"account_user_status"`
	// Display name of the user from the customer's IdP.
	DisplayName types.String `tfsdk:"display_name"`
	// The external ID of the user in the customer's IdP.
	ExternalUserId types.String `tfsdk:"external_user_id"`
	// The full name of the user, from the customer's IdP.
	FullName types.Object `tfsdk:"full_name"`
	// Internal userId of the user in Databricks.
	InternalId types.String `tfsdk:"internal_id"`
	// The resource name of the external user. The format depends on the API
	// that returned it: - Account-scoped:
	// accounts/{account_id}/external-users/{external_user_id} -
	// Workspace-scoped: external-users/{external_user_id}
	Name types.String `tfsdk:"name"`
	// Username/email of the user, from Databricks.
	Username types.String `tfsdk:"username"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// ExternalUserData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m ExternalUserData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"full_name": reflect.TypeOf(iamv2_tf.FullName{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalUserData
// only implements ToObjectValue() and Type().
func (m ExternalUserData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":          m.AccountId,
			"account_user_status": m.AccountUserStatus,
			"display_name":        m.DisplayName,
			"external_user_id":    m.ExternalUserId,
			"full_name":           m.FullName,
			"internal_id":         m.InternalId,
			"name":                m.Name,
			"username":            m.Username,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m ExternalUserData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":          types.StringType,
			"account_user_status": types.StringType,
			"display_name":        types.StringType,
			"external_user_id":    types.StringType,
			"full_name":           iamv2_tf.FullName{}.Type(ctx),
			"internal_id":         types.StringType,
			"name":                types.StringType,
			"username":            types.StringType,
		},
	}
}

func (m ExternalUserData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetComputed()
	attrs["account_user_status"] = attrs["account_user_status"].SetComputed()
	attrs["display_name"] = attrs["display_name"].SetComputed()
	attrs["external_user_id"] = attrs["external_user_id"].SetComputed()
	attrs["full_name"] = attrs["full_name"].SetComputed()
	attrs["internal_id"] = attrs["internal_id"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["username"] = attrs["username"].SetComputed()

	return attrs
}

func (r *ExternalUserDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *ExternalUserDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, ExternalUserData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks ExternalUser",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *ExternalUserDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *ExternalUserDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config ExternalUserData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest iamv2.GetExternalUserRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetAccountClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.IamV2.GetExternalUser(ctx, readRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to get account_iam_external_user_v2", err.Error())
		return
	}

	var newState ExternalUserData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
