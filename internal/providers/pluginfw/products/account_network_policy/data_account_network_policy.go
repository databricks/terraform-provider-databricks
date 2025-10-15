// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package account_network_policy

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/settings_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceName = "account_network_policy"

var _ datasource.DataSourceWithConfigure = &AccountNetworkPolicyDataSource{}

func DataSourceAccountNetworkPolicy() datasource.DataSource {
	return &AccountNetworkPolicyDataSource{}
}

type AccountNetworkPolicyDataSource struct {
	Client *autogen.DatabricksClient
}

// AccountNetworkPolicyData extends the main model with additional fields.
type AccountNetworkPolicyData struct {
	// The associated account ID for this Network Policy object.
	AccountId types.String `tfsdk:"account_id"`
	// The network policies applying for egress traffic.
	Egress types.Object `tfsdk:"egress"`
	// The unique identifier for the network policy.
	NetworkPolicyId types.String `tfsdk:"network_policy_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// AccountNetworkPolicyData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m AccountNetworkPolicyData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"egress": reflect.TypeOf(settings_tf.NetworkPolicyEgress{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccountNetworkPolicyData
// only implements ToObjectValue() and Type().
func (m AccountNetworkPolicyData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"account_id": m.AccountId,
			"egress":            m.Egress,
			"network_policy_id": m.NetworkPolicyId,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m AccountNetworkPolicyData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"account_id": types.StringType,
			"egress":            settings_tf.NetworkPolicyEgress{}.Type(ctx),
			"network_policy_id": types.StringType,
		},
	}
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *AccountNetworkPolicyData) SyncFieldsDuringRead(ctx context.Context, from AccountNetworkPolicyData) {
}

func (m AccountNetworkPolicyData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetOptional()
	attrs["egress"] = attrs["egress"].SetOptional()
	attrs["network_policy_id"] = attrs["network_policy_id"].SetOptional()

	return attrs
}

func (r *AccountNetworkPolicyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *AccountNetworkPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, AccountNetworkPolicyData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks AccountNetworkPolicy",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *AccountNetworkPolicyDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *AccountNetworkPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config AccountNetworkPolicyData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest settings.GetNetworkPolicyRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetAccountClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.NetworkPolicies.GetNetworkPolicyRpc(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get account_network_policy", err.Error())
		return
	}

	var newState AccountNetworkPolicyData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, config)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
