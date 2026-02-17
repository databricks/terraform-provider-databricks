// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package apps_space

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/service/apps"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/apps_tf"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const resourceName = "apps_space"

var _ resource.ResourceWithConfigure = &SpaceResource{}

func ResourceSpace() resource.Resource {
	return &SpaceResource{}
}

type SpaceResource struct {
	Client *autogen.DatabricksClient
}

// ProviderConfig contains the fields to configure the provider.
type ProviderConfig struct {
	WorkspaceID types.String `tfsdk:"workspace_id"`
}

// ApplySchemaCustomizations applies the schema customizations to the ProviderConfig type.
func (r ProviderConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()
	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(
		stringplanmodifier.RequiresReplaceIf(ProviderConfigWorkspaceIDPlanModifier, "", ""))

	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.StringAttributeBuilder).AddValidator(stringvalidator.LengthAtLeast(1))
	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.StringAttributeBuilder).AddValidator(
		stringvalidator.RegexMatches(regexp.MustCompile(`^[1-9]\d*$`), "workspace_id must be a positive integer without leading zeros"))
	return attrs
}

// ProviderConfigWorkspaceIDPlanModifier is plan modifier for the workspace_id field.
// Resource requires replacement if the workspace_id changes from one non-empty value to another.
func ProviderConfigWorkspaceIDPlanModifier(ctx context.Context, req planmodifier.StringRequest, resp *stringplanmodifier.RequiresReplaceIfFuncResponse) {
	// Require replacement if workspace_id changes from one non-empty value to another
	oldValue := req.StateValue.ValueString()
	newValue := req.PlanValue.ValueString()

	if oldValue != "" && newValue != "" && oldValue != newValue {
		resp.RequiresReplace = true
	}
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// ProviderConfig struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (r ProviderConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProviderConfig
// only implements ToObjectValue() and Type().
func (r ProviderConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		r.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": r.WorkspaceID,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (r ProviderConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.StringType,
		},
	}
}

// Space extends the main model with additional fields.
type Space struct {
	// The creation time of the app space. Formatted timestamp in ISO 6801.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// The email of the user that created the app space.
	Creator types.String `tfsdk:"creator"`
	// The description of the app space.
	Description types.String `tfsdk:"description"`
	// The effective usage policy ID used by apps in the space.
	EffectiveUsagePolicyId types.String `tfsdk:"effective_usage_policy_id"`
	// The effective api scopes granted to the user access token.
	EffectiveUserApiScopes types.List `tfsdk:"effective_user_api_scopes"`
	// The unique identifier of the app space.
	Id types.String `tfsdk:"id"`
	// The name of the app space. The name must contain only lowercase
	// alphanumeric characters and hyphens. It must be unique within the
	// workspace.
	Name types.String `tfsdk:"name"`
	// The OAuth2 app client ID for the app space.
	Oauth2AppClientId types.String `tfsdk:"oauth2_app_client_id"`
	// The OAuth2 app integration ID for the app space.
	Oauth2AppIntegrationId types.String `tfsdk:"oauth2_app_integration_id"`
	// Resources for the app space. Resources configured at the space level are
	// available to all apps in the space.
	Resources types.List `tfsdk:"resources"`
	// The service principal client ID for the app space.
	ServicePrincipalClientId types.String `tfsdk:"service_principal_client_id"`
	// The service principal ID for the app space.
	ServicePrincipalId types.Int64 `tfsdk:"service_principal_id"`
	// The service principal name for the app space.
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// The status of the app space.
	Status types.Object `tfsdk:"status"`
	// The update time of the app space. Formatted timestamp in ISO 6801.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
	// The email of the user that last updated the app space.
	Updater types.String `tfsdk:"updater"`
	// The usage policy ID for managing cost at the space level.
	UsagePolicyId types.String `tfsdk:"usage_policy_id"`
	// OAuth scopes for apps in the space.
	UserApiScopes  types.List   `tfsdk:"user_api_scopes"`
	ProviderConfig types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// Space struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m Space) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"effective_user_api_scopes": reflect.TypeOf(types.String{}),
		"resources":                 reflect.TypeOf(apps_tf.AppResource{}),
		"status":                    reflect.TypeOf(apps_tf.SpaceStatus{}),
		"user_api_scopes":           reflect.TypeOf(types.String{}),
		"provider_config":           reflect.TypeOf(ProviderConfig{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Space
// only implements ToObjectValue() and Type().
func (m Space) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"create_time": m.CreateTime,
			"creator":                     m.Creator,
			"description":                 m.Description,
			"effective_usage_policy_id":   m.EffectiveUsagePolicyId,
			"effective_user_api_scopes":   m.EffectiveUserApiScopes,
			"id":                          m.Id,
			"name":                        m.Name,
			"oauth2_app_client_id":        m.Oauth2AppClientId,
			"oauth2_app_integration_id":   m.Oauth2AppIntegrationId,
			"resources":                   m.Resources,
			"service_principal_client_id": m.ServicePrincipalClientId,
			"service_principal_id":        m.ServicePrincipalId,
			"service_principal_name":      m.ServicePrincipalName,
			"status":                      m.Status,
			"update_time":                 m.UpdateTime,
			"updater":                     m.Updater,
			"usage_policy_id":             m.UsagePolicyId,
			"user_api_scopes":             m.UserApiScopes,

			"provider_config": m.ProviderConfig,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m Space) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"create_time": timetypes.RFC3339{}.Type(ctx),
			"creator":                   types.StringType,
			"description":               types.StringType,
			"effective_usage_policy_id": types.StringType,
			"effective_user_api_scopes": basetypes.ListType{
				ElemType: types.StringType,
			},
			"id":                        types.StringType,
			"name":                      types.StringType,
			"oauth2_app_client_id":      types.StringType,
			"oauth2_app_integration_id": types.StringType,
			"resources": basetypes.ListType{
				ElemType: apps_tf.AppResource{}.Type(ctx),
			},
			"service_principal_client_id": types.StringType,
			"service_principal_id":        types.Int64Type,
			"service_principal_name":      types.StringType,
			"status":                      apps_tf.SpaceStatus{}.Type(ctx),
			"update_time":                 timetypes.RFC3339{}.Type(ctx),
			"updater":                     types.StringType,
			"usage_policy_id":             types.StringType,
			"user_api_scopes": basetypes.ListType{
				ElemType: types.StringType,
			},

			"provider_config": ProviderConfig{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *Space) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Space) {
	if !from.EffectiveUserApiScopes.IsNull() && !from.EffectiveUserApiScopes.IsUnknown() && to.EffectiveUserApiScopes.IsNull() && len(from.EffectiveUserApiScopes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EffectiveUserApiScopes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EffectiveUserApiScopes = from.EffectiveUserApiScopes
	}
	if !from.Resources.IsNull() && !from.Resources.IsUnknown() && to.Resources.IsNull() && len(from.Resources.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Resources, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Resources = from.Resources
	}
	if !from.Status.IsNull() && !from.Status.IsUnknown() {
		if toStatus, ok := to.GetStatus(ctx); ok {
			if fromStatus, ok := from.GetStatus(ctx); ok {
				// Recursively sync the fields of Status
				toStatus.SyncFieldsDuringCreateOrUpdate(ctx, fromStatus)
				to.SetStatus(ctx, toStatus)
			}
		}
	}
	if !from.UserApiScopes.IsNull() && !from.UserApiScopes.IsUnknown() && to.UserApiScopes.IsNull() && len(from.UserApiScopes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for UserApiScopes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.UserApiScopes = from.UserApiScopes
	}
	to.ProviderConfig = from.ProviderConfig

}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *Space) SyncFieldsDuringRead(ctx context.Context, from Space) {
	if !from.EffectiveUserApiScopes.IsNull() && !from.EffectiveUserApiScopes.IsUnknown() && to.EffectiveUserApiScopes.IsNull() && len(from.EffectiveUserApiScopes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EffectiveUserApiScopes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EffectiveUserApiScopes = from.EffectiveUserApiScopes
	}
	if !from.Resources.IsNull() && !from.Resources.IsUnknown() && to.Resources.IsNull() && len(from.Resources.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Resources, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Resources = from.Resources
	}
	if !from.Status.IsNull() && !from.Status.IsUnknown() {
		if toStatus, ok := to.GetStatus(ctx); ok {
			if fromStatus, ok := from.GetStatus(ctx); ok {
				toStatus.SyncFieldsDuringRead(ctx, fromStatus)
				to.SetStatus(ctx, toStatus)
			}
		}
	}
	if !from.UserApiScopes.IsNull() && !from.UserApiScopes.IsUnknown() && to.UserApiScopes.IsNull() && len(from.UserApiScopes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for UserApiScopes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.UserApiScopes = from.UserApiScopes
	}
	to.ProviderConfig = from.ProviderConfig

}

func (m Space) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["creator"] = attrs["creator"].SetComputed()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["effective_usage_policy_id"] = attrs["effective_usage_policy_id"].SetComputed()
	attrs["effective_user_api_scopes"] = attrs["effective_user_api_scopes"].SetComputed()
	attrs["id"] = attrs["id"].SetComputed()
	attrs["name"] = attrs["name"].SetComputed()
	attrs["oauth2_app_client_id"] = attrs["oauth2_app_client_id"].SetComputed()
	attrs["oauth2_app_integration_id"] = attrs["oauth2_app_integration_id"].SetComputed()
	attrs["resources"] = attrs["resources"].SetOptional()
	attrs["service_principal_client_id"] = attrs["service_principal_client_id"].SetComputed()
	attrs["service_principal_id"] = attrs["service_principal_id"].SetComputed()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetComputed()
	attrs["status"] = attrs["status"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["updater"] = attrs["updater"].SetComputed()
	attrs["usage_policy_id"] = attrs["usage_policy_id"].SetOptional()
	attrs["user_api_scopes"] = attrs["user_api_scopes"].SetOptional()

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

// GetEffectiveUserApiScopes returns the value of the EffectiveUserApiScopes field in Space as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *Space) GetEffectiveUserApiScopes(ctx context.Context) ([]types.String, bool) {
	if m.EffectiveUserApiScopes.IsNull() || m.EffectiveUserApiScopes.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.EffectiveUserApiScopes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveUserApiScopes sets the value of the EffectiveUserApiScopes field in Space.
func (m *Space) SetEffectiveUserApiScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_user_api_scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EffectiveUserApiScopes = types.ListValueMust(t, vs)
}

// GetResources returns the value of the Resources field in Space as
// a slice of apps_tf.AppResource values.
// If the field is unknown or null, the boolean return value is false.
func (m *Space) GetResources(ctx context.Context) ([]apps_tf.AppResource, bool) {
	if m.Resources.IsNull() || m.Resources.IsUnknown() {
		return nil, false
	}
	var v []apps_tf.AppResource
	d := m.Resources.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResources sets the value of the Resources field in Space.
func (m *Space) SetResources(ctx context.Context, v []apps_tf.AppResource) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["resources"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Resources = types.ListValueMust(t, vs)
}

// GetStatus returns the value of the Status field in Space as
// a apps_tf.SpaceStatus value.
// If the field is unknown or null, the boolean return value is false.
func (m *Space) GetStatus(ctx context.Context) (apps_tf.SpaceStatus, bool) {
	var e apps_tf.SpaceStatus
	if m.Status.IsNull() || m.Status.IsUnknown() {
		return e, false
	}
	var v apps_tf.SpaceStatus
	d := m.Status.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatus sets the value of the Status field in Space.
func (m *Space) SetStatus(ctx context.Context, v apps_tf.SpaceStatus) {
	vs := v.ToObjectValue(ctx)
	m.Status = vs
}

// GetUserApiScopes returns the value of the UserApiScopes field in Space as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *Space) GetUserApiScopes(ctx context.Context) ([]types.String, bool) {
	if m.UserApiScopes.IsNull() || m.UserApiScopes.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.UserApiScopes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUserApiScopes sets the value of the UserApiScopes field in Space.
func (m *Space) SetUserApiScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["user_api_scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.UserApiScopes = types.ListValueMust(t, vs)
}

func (r *SpaceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *SpaceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, Space{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks apps_space",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *SpaceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *SpaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan Space
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var space apps.Space

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &space)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := apps.CreateSpaceRequest{
		Space: space,
	}

	var namespace ProviderConfig
	resp.Diagnostics.Append(plan.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
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

	response, err := client.Apps.CreateSpace(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create apps_space", err.Error())
		return
	}

	var newState Space

	waitResponse, err := response.Wait(ctx)
	if err != nil {
		resp.Diagnostics.AddError("error waiting for apps_space to be ready", err.Error())
		return
	}

	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, waitResponse, &newState)...)

	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *SpaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState Space
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest apps.GetSpaceRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var namespace ProviderConfig
	resp.Diagnostics.Append(existingState.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
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
	response, err := client.Apps.GetSpace(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get apps_space", err.Error())
		return
	}

	var newState Space
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *SpaceResource) update(ctx context.Context, plan Space, diags *diag.Diagnostics, state *tfsdk.State) {
	var space apps.Space

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &space)...)
	if diags.HasError() {
		return
	}

	updateRequest := apps.UpdateSpaceRequest{
		Space:      space,
		Name:       plan.Name.ValueString(),
		UpdateMask: *fieldmask.New(strings.Split("description,resources,usage_policy_id,user_api_scopes", ",")),
	}

	var namespace ProviderConfig
	diags.Append(plan.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if diags.HasError() {
		return
	}
	client, clientDiags := r.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, namespace.WorkspaceID.ValueString())

	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}
	response, err := client.Apps.UpdateSpace(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update apps_space", err.Error())
		return
	}

	var newState Space

	waitResponse, err := response.Wait(ctx)
	if err != nil {
		diags.AddError("error waiting for apps_space update", err.Error())
		return
	}

	diags.Append(converters.GoSdkToTfSdkStruct(ctx, waitResponse, &newState)...)

	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *SpaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan Space
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *SpaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state Space
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest apps.DeleteSpaceRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var namespace ProviderConfig
	resp.Diagnostics.Append(state.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
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

	response, err := client.Apps.DeleteSpace(ctx, deleteRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to delete apps_space", err.Error())
		return
	}

	err = response.Wait(ctx)
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("error waiting for apps_space delete", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &SpaceResource{}

func (r *SpaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, ",")

	if len(parts) != 1 || parts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf(
				"Expected import identifier with format: name. Got: %q",
				req.ID,
			),
		)
		return
	}

	name := parts[0]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), name)...)
}
