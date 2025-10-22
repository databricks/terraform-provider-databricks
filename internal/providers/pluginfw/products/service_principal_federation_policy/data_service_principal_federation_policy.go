// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package service_principal_federation_policy

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/oauth2"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/oauth2_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceName = "service_principal_federation_policy"

var _ datasource.DataSourceWithConfigure = &FederationPolicyDataSource{}

func DataSourceFederationPolicy() datasource.DataSource {
	return &FederationPolicyDataSource{}
}

type FederationPolicyDataSource struct {
	Client *autogen.DatabricksClient
}

// FederationPolicyData extends the main model with additional fields.
type FederationPolicyData struct {
	// Creation time of the federation policy.
	CreateTime types.String `tfsdk:"create_time"`
	// Description of the federation policy.
	Description types.String `tfsdk:"description"`
	// Resource name for the federation policy. Example values include
	// `accounts/<account-id>/federationPolicies/my-federation-policy` for
	// Account Federation Policies, and
	// `accounts/<account-id>/servicePrincipals/<service-principal-id>/federationPolicies/my-federation-policy`
	// for Service Principal Federation Policies. Typically an output parameter,
	// which does not need to be specified in create or update requests. If
	// specified in a request, must match the value in the request URL.
	Name types.String `tfsdk:"name"`

	OidcPolicy types.Object `tfsdk:"oidc_policy"`
	// The ID of the federation policy. Output only.
	PolicyId types.String `tfsdk:"policy_id"`
	// The service principal ID that this federation policy applies to. Output
	// only. Only set for service principal federation policies.
	ServicePrincipalId types.Int64 `tfsdk:"service_principal_id"`
	// Unique, immutable id of the federation policy.
	Uid types.String `tfsdk:"uid"`
	// Last update time of the federation policy.
	UpdateTime types.String `tfsdk:"update_time"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// FederationPolicyData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m FederationPolicyData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"oidc_policy": reflect.TypeOf(oauth2_tf.OidcFederationPolicy{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FederationPolicyData
// only implements ToObjectValue() and Type().
func (m FederationPolicyData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time":          m.CreateTime,
			"description":          m.Description,
			"name":                 m.Name,
			"oidc_policy":          m.OidcPolicy,
			"policy_id":            m.PolicyId,
			"service_principal_id": m.ServicePrincipalId,
			"uid":                  m.Uid,
			"update_time":          m.UpdateTime,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m FederationPolicyData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"create_time": types.StringType,
			"description":          types.StringType,
			"name":                 types.StringType,
			"oidc_policy":          oauth2_tf.OidcFederationPolicy{}.Type(ctx),
			"policy_id":            types.StringType,
			"service_principal_id": types.Int64Type,
			"uid":                  types.StringType,
			"update_time":          types.StringType,
		},
	}
}

func (m FederationPolicyData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["name"] = attrs["name"].SetComputed()
	attrs["oidc_policy"] = attrs["oidc_policy"].SetOptional()
	attrs["policy_id"] = attrs["policy_id"].SetComputed()
	attrs["service_principal_id"] = attrs["service_principal_id"].SetComputed()
	attrs["uid"] = attrs["uid"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()

	return attrs
}

func (r *FederationPolicyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *FederationPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, FederationPolicyData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks FederationPolicy",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *FederationPolicyDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *FederationPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config FederationPolicyData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest oauth2.GetServicePrincipalFederationPolicyRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetAccountClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.ServicePrincipalFederationPolicy.Get(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get service_principal_federation_policy", err.Error())
		return
	}

	var newState FederationPolicyData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
