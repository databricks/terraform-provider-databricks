// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package apps_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type App struct {
	// The active deployment of the app. A deployment is considered active when
	// it has been deployed to the app compute.
	ActiveDeployment types.Object `tfsdk:"active_deployment"`

	AppStatus types.Object `tfsdk:"app_status"`

	BudgetPolicyId types.String `tfsdk:"budget_policy_id"`

	ComputeSize types.String `tfsdk:"compute_size"`

	ComputeStatus types.Object `tfsdk:"compute_status"`
	// The creation time of the app. Formatted timestamp in ISO 6801.
	CreateTime types.String `tfsdk:"create_time"`
	// The email of the user that created the app.
	Creator types.String `tfsdk:"creator"`
	// The default workspace file system path of the source code from which app
	// deployment are created. This field tracks the workspace source code path
	// of the last active deployment.
	DefaultSourceCodePath types.String `tfsdk:"default_source_code_path"`
	// The description of the app.
	Description types.String `tfsdk:"description"`

	EffectiveBudgetPolicyId types.String `tfsdk:"effective_budget_policy_id"`

	EffectiveUsagePolicyId types.String `tfsdk:"effective_usage_policy_id"`
	// The effective api scopes granted to the user access token.
	EffectiveUserApiScopes types.List `tfsdk:"effective_user_api_scopes"`
	// Git repository configuration for app deployments. When specified,
	// deployments can reference code from this repository by providing only the
	// git reference (branch, tag, or commit).
	GitRepository types.Object `tfsdk:"git_repository"`
	// The unique identifier of the app.
	Id types.String `tfsdk:"id"`
	// The name of the app. The name must contain only lowercase alphanumeric
	// characters and hyphens. It must be unique within the workspace.
	Name types.String `tfsdk:"name"`

	Oauth2AppClientId types.String `tfsdk:"oauth2_app_client_id"`

	Oauth2AppIntegrationId types.String `tfsdk:"oauth2_app_integration_id"`
	// The pending deployment of the app. A deployment is considered pending
	// when it is being prepared for deployment to the app compute.
	PendingDeployment types.Object `tfsdk:"pending_deployment"`
	// Resources for the app.
	Resources types.List `tfsdk:"resources"`

	ServicePrincipalClientId types.String `tfsdk:"service_principal_client_id"`

	ServicePrincipalId types.Int64 `tfsdk:"service_principal_id"`

	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// Name of the space this app belongs to.
	Space types.String `tfsdk:"space"`
	// The update time of the app. Formatted timestamp in ISO 6801.
	UpdateTime types.String `tfsdk:"update_time"`
	// The email of the user that last updated the app.
	Updater types.String `tfsdk:"updater"`
	// The URL of the app once it is deployed.
	Url types.String `tfsdk:"url"`

	UsagePolicyId types.String `tfsdk:"usage_policy_id"`

	UserApiScopes types.List `tfsdk:"user_api_scopes"`
}

func (to *App) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from App) {
	if !from.ActiveDeployment.IsNull() && !from.ActiveDeployment.IsUnknown() {
		if toActiveDeployment, ok := to.GetActiveDeployment(ctx); ok {
			if fromActiveDeployment, ok := from.GetActiveDeployment(ctx); ok {
				// Recursively sync the fields of ActiveDeployment
				toActiveDeployment.SyncFieldsDuringCreateOrUpdate(ctx, fromActiveDeployment)
				to.SetActiveDeployment(ctx, toActiveDeployment)
			}
		}
	}
	if !from.AppStatus.IsNull() && !from.AppStatus.IsUnknown() {
		if toAppStatus, ok := to.GetAppStatus(ctx); ok {
			if fromAppStatus, ok := from.GetAppStatus(ctx); ok {
				// Recursively sync the fields of AppStatus
				toAppStatus.SyncFieldsDuringCreateOrUpdate(ctx, fromAppStatus)
				to.SetAppStatus(ctx, toAppStatus)
			}
		}
	}
	if !from.ComputeStatus.IsNull() && !from.ComputeStatus.IsUnknown() {
		if toComputeStatus, ok := to.GetComputeStatus(ctx); ok {
			if fromComputeStatus, ok := from.GetComputeStatus(ctx); ok {
				// Recursively sync the fields of ComputeStatus
				toComputeStatus.SyncFieldsDuringCreateOrUpdate(ctx, fromComputeStatus)
				to.SetComputeStatus(ctx, toComputeStatus)
			}
		}
	}
	if !from.EffectiveUserApiScopes.IsNull() && !from.EffectiveUserApiScopes.IsUnknown() && to.EffectiveUserApiScopes.IsNull() && len(from.EffectiveUserApiScopes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EffectiveUserApiScopes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EffectiveUserApiScopes = from.EffectiveUserApiScopes
	}
	if !from.GitRepository.IsNull() && !from.GitRepository.IsUnknown() {
		if toGitRepository, ok := to.GetGitRepository(ctx); ok {
			if fromGitRepository, ok := from.GetGitRepository(ctx); ok {
				// Recursively sync the fields of GitRepository
				toGitRepository.SyncFieldsDuringCreateOrUpdate(ctx, fromGitRepository)
				to.SetGitRepository(ctx, toGitRepository)
			}
		}
	}
	if !from.PendingDeployment.IsNull() && !from.PendingDeployment.IsUnknown() {
		if toPendingDeployment, ok := to.GetPendingDeployment(ctx); ok {
			if fromPendingDeployment, ok := from.GetPendingDeployment(ctx); ok {
				// Recursively sync the fields of PendingDeployment
				toPendingDeployment.SyncFieldsDuringCreateOrUpdate(ctx, fromPendingDeployment)
				to.SetPendingDeployment(ctx, toPendingDeployment)
			}
		}
	}
	if !from.Resources.IsNull() && !from.Resources.IsUnknown() && to.Resources.IsNull() && len(from.Resources.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Resources, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Resources = from.Resources
	}
	if !from.UserApiScopes.IsNull() && !from.UserApiScopes.IsUnknown() && to.UserApiScopes.IsNull() && len(from.UserApiScopes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for UserApiScopes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.UserApiScopes = from.UserApiScopes
	}
}

func (to *App) SyncFieldsDuringRead(ctx context.Context, from App) {
	if !from.ActiveDeployment.IsNull() && !from.ActiveDeployment.IsUnknown() {
		if toActiveDeployment, ok := to.GetActiveDeployment(ctx); ok {
			if fromActiveDeployment, ok := from.GetActiveDeployment(ctx); ok {
				toActiveDeployment.SyncFieldsDuringRead(ctx, fromActiveDeployment)
				to.SetActiveDeployment(ctx, toActiveDeployment)
			}
		}
	}
	if !from.AppStatus.IsNull() && !from.AppStatus.IsUnknown() {
		if toAppStatus, ok := to.GetAppStatus(ctx); ok {
			if fromAppStatus, ok := from.GetAppStatus(ctx); ok {
				toAppStatus.SyncFieldsDuringRead(ctx, fromAppStatus)
				to.SetAppStatus(ctx, toAppStatus)
			}
		}
	}
	if !from.ComputeStatus.IsNull() && !from.ComputeStatus.IsUnknown() {
		if toComputeStatus, ok := to.GetComputeStatus(ctx); ok {
			if fromComputeStatus, ok := from.GetComputeStatus(ctx); ok {
				toComputeStatus.SyncFieldsDuringRead(ctx, fromComputeStatus)
				to.SetComputeStatus(ctx, toComputeStatus)
			}
		}
	}
	if !from.EffectiveUserApiScopes.IsNull() && !from.EffectiveUserApiScopes.IsUnknown() && to.EffectiveUserApiScopes.IsNull() && len(from.EffectiveUserApiScopes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EffectiveUserApiScopes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EffectiveUserApiScopes = from.EffectiveUserApiScopes
	}
	if !from.GitRepository.IsNull() && !from.GitRepository.IsUnknown() {
		if toGitRepository, ok := to.GetGitRepository(ctx); ok {
			if fromGitRepository, ok := from.GetGitRepository(ctx); ok {
				toGitRepository.SyncFieldsDuringRead(ctx, fromGitRepository)
				to.SetGitRepository(ctx, toGitRepository)
			}
		}
	}
	if !from.PendingDeployment.IsNull() && !from.PendingDeployment.IsUnknown() {
		if toPendingDeployment, ok := to.GetPendingDeployment(ctx); ok {
			if fromPendingDeployment, ok := from.GetPendingDeployment(ctx); ok {
				toPendingDeployment.SyncFieldsDuringRead(ctx, fromPendingDeployment)
				to.SetPendingDeployment(ctx, toPendingDeployment)
			}
		}
	}
	if !from.Resources.IsNull() && !from.Resources.IsUnknown() && to.Resources.IsNull() && len(from.Resources.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Resources, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Resources = from.Resources
	}
	if !from.UserApiScopes.IsNull() && !from.UserApiScopes.IsUnknown() && to.UserApiScopes.IsNull() && len(from.UserApiScopes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for UserApiScopes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.UserApiScopes = from.UserApiScopes
	}
}

func (m App) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["active_deployment"] = attrs["active_deployment"].SetComputed()
	attrs["app_status"] = attrs["app_status"].SetComputed()
	attrs["budget_policy_id"] = attrs["budget_policy_id"].SetOptional()
	attrs["compute_size"] = attrs["compute_size"].SetOptional()
	attrs["compute_status"] = attrs["compute_status"].SetComputed()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["creator"] = attrs["creator"].SetComputed()
	attrs["default_source_code_path"] = attrs["default_source_code_path"].SetComputed()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["effective_budget_policy_id"] = attrs["effective_budget_policy_id"].SetComputed()
	attrs["effective_usage_policy_id"] = attrs["effective_usage_policy_id"].SetComputed()
	attrs["effective_user_api_scopes"] = attrs["effective_user_api_scopes"].SetComputed()
	attrs["git_repository"] = attrs["git_repository"].SetOptional()
	attrs["id"] = attrs["id"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["oauth2_app_client_id"] = attrs["oauth2_app_client_id"].SetComputed()
	attrs["oauth2_app_integration_id"] = attrs["oauth2_app_integration_id"].SetComputed()
	attrs["pending_deployment"] = attrs["pending_deployment"].SetComputed()
	attrs["resources"] = attrs["resources"].SetOptional()
	attrs["service_principal_client_id"] = attrs["service_principal_client_id"].SetComputed()
	attrs["service_principal_id"] = attrs["service_principal_id"].SetComputed()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetComputed()
	attrs["space"] = attrs["space"].SetOptional()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["updater"] = attrs["updater"].SetComputed()
	attrs["url"] = attrs["url"].SetComputed()
	attrs["usage_policy_id"] = attrs["usage_policy_id"].SetOptional()
	attrs["user_api_scopes"] = attrs["user_api_scopes"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in App.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m App) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"active_deployment":         reflect.TypeOf(AppDeployment{}),
		"app_status":                reflect.TypeOf(ApplicationStatus{}),
		"compute_status":            reflect.TypeOf(ComputeStatus{}),
		"effective_user_api_scopes": reflect.TypeOf(types.String{}),
		"git_repository":            reflect.TypeOf(GitRepository{}),
		"pending_deployment":        reflect.TypeOf(AppDeployment{}),
		"resources":                 reflect.TypeOf(AppResource{}),
		"user_api_scopes":           reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, App
// only implements ToObjectValue() and Type().
func (m App) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"active_deployment":           m.ActiveDeployment,
			"app_status":                  m.AppStatus,
			"budget_policy_id":            m.BudgetPolicyId,
			"compute_size":                m.ComputeSize,
			"compute_status":              m.ComputeStatus,
			"create_time":                 m.CreateTime,
			"creator":                     m.Creator,
			"default_source_code_path":    m.DefaultSourceCodePath,
			"description":                 m.Description,
			"effective_budget_policy_id":  m.EffectiveBudgetPolicyId,
			"effective_usage_policy_id":   m.EffectiveUsagePolicyId,
			"effective_user_api_scopes":   m.EffectiveUserApiScopes,
			"git_repository":              m.GitRepository,
			"id":                          m.Id,
			"name":                        m.Name,
			"oauth2_app_client_id":        m.Oauth2AppClientId,
			"oauth2_app_integration_id":   m.Oauth2AppIntegrationId,
			"pending_deployment":          m.PendingDeployment,
			"resources":                   m.Resources,
			"service_principal_client_id": m.ServicePrincipalClientId,
			"service_principal_id":        m.ServicePrincipalId,
			"service_principal_name":      m.ServicePrincipalName,
			"space":                       m.Space,
			"update_time":                 m.UpdateTime,
			"updater":                     m.Updater,
			"url":                         m.Url,
			"usage_policy_id":             m.UsagePolicyId,
			"user_api_scopes":             m.UserApiScopes,
		})
}

// Type implements basetypes.ObjectValuable.
func (m App) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"active_deployment":          AppDeployment{}.Type(ctx),
			"app_status":                 ApplicationStatus{}.Type(ctx),
			"budget_policy_id":           types.StringType,
			"compute_size":               types.StringType,
			"compute_status":             ComputeStatus{}.Type(ctx),
			"create_time":                types.StringType,
			"creator":                    types.StringType,
			"default_source_code_path":   types.StringType,
			"description":                types.StringType,
			"effective_budget_policy_id": types.StringType,
			"effective_usage_policy_id":  types.StringType,
			"effective_user_api_scopes": basetypes.ListType{
				ElemType: types.StringType,
			},
			"git_repository":            GitRepository{}.Type(ctx),
			"id":                        types.StringType,
			"name":                      types.StringType,
			"oauth2_app_client_id":      types.StringType,
			"oauth2_app_integration_id": types.StringType,
			"pending_deployment":        AppDeployment{}.Type(ctx),
			"resources": basetypes.ListType{
				ElemType: AppResource{}.Type(ctx),
			},
			"service_principal_client_id": types.StringType,
			"service_principal_id":        types.Int64Type,
			"service_principal_name":      types.StringType,
			"space":                       types.StringType,
			"update_time":                 types.StringType,
			"updater":                     types.StringType,
			"url":                         types.StringType,
			"usage_policy_id":             types.StringType,
			"user_api_scopes": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetActiveDeployment returns the value of the ActiveDeployment field in App as
// a AppDeployment value.
// If the field is unknown or null, the boolean return value is false.
func (m *App) GetActiveDeployment(ctx context.Context) (AppDeployment, bool) {
	var e AppDeployment
	if m.ActiveDeployment.IsNull() || m.ActiveDeployment.IsUnknown() {
		return e, false
	}
	var v AppDeployment
	d := m.ActiveDeployment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetActiveDeployment sets the value of the ActiveDeployment field in App.
func (m *App) SetActiveDeployment(ctx context.Context, v AppDeployment) {
	vs := v.ToObjectValue(ctx)
	m.ActiveDeployment = vs
}

// GetAppStatus returns the value of the AppStatus field in App as
// a ApplicationStatus value.
// If the field is unknown or null, the boolean return value is false.
func (m *App) GetAppStatus(ctx context.Context) (ApplicationStatus, bool) {
	var e ApplicationStatus
	if m.AppStatus.IsNull() || m.AppStatus.IsUnknown() {
		return e, false
	}
	var v ApplicationStatus
	d := m.AppStatus.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAppStatus sets the value of the AppStatus field in App.
func (m *App) SetAppStatus(ctx context.Context, v ApplicationStatus) {
	vs := v.ToObjectValue(ctx)
	m.AppStatus = vs
}

// GetComputeStatus returns the value of the ComputeStatus field in App as
// a ComputeStatus value.
// If the field is unknown or null, the boolean return value is false.
func (m *App) GetComputeStatus(ctx context.Context) (ComputeStatus, bool) {
	var e ComputeStatus
	if m.ComputeStatus.IsNull() || m.ComputeStatus.IsUnknown() {
		return e, false
	}
	var v ComputeStatus
	d := m.ComputeStatus.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetComputeStatus sets the value of the ComputeStatus field in App.
func (m *App) SetComputeStatus(ctx context.Context, v ComputeStatus) {
	vs := v.ToObjectValue(ctx)
	m.ComputeStatus = vs
}

// GetEffectiveUserApiScopes returns the value of the EffectiveUserApiScopes field in App as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *App) GetEffectiveUserApiScopes(ctx context.Context) ([]types.String, bool) {
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

// SetEffectiveUserApiScopes sets the value of the EffectiveUserApiScopes field in App.
func (m *App) SetEffectiveUserApiScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_user_api_scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EffectiveUserApiScopes = types.ListValueMust(t, vs)
}

// GetGitRepository returns the value of the GitRepository field in App as
// a GitRepository value.
// If the field is unknown or null, the boolean return value is false.
func (m *App) GetGitRepository(ctx context.Context) (GitRepository, bool) {
	var e GitRepository
	if m.GitRepository.IsNull() || m.GitRepository.IsUnknown() {
		return e, false
	}
	var v GitRepository
	d := m.GitRepository.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGitRepository sets the value of the GitRepository field in App.
func (m *App) SetGitRepository(ctx context.Context, v GitRepository) {
	vs := v.ToObjectValue(ctx)
	m.GitRepository = vs
}

// GetPendingDeployment returns the value of the PendingDeployment field in App as
// a AppDeployment value.
// If the field is unknown or null, the boolean return value is false.
func (m *App) GetPendingDeployment(ctx context.Context) (AppDeployment, bool) {
	var e AppDeployment
	if m.PendingDeployment.IsNull() || m.PendingDeployment.IsUnknown() {
		return e, false
	}
	var v AppDeployment
	d := m.PendingDeployment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPendingDeployment sets the value of the PendingDeployment field in App.
func (m *App) SetPendingDeployment(ctx context.Context, v AppDeployment) {
	vs := v.ToObjectValue(ctx)
	m.PendingDeployment = vs
}

// GetResources returns the value of the Resources field in App as
// a slice of AppResource values.
// If the field is unknown or null, the boolean return value is false.
func (m *App) GetResources(ctx context.Context) ([]AppResource, bool) {
	if m.Resources.IsNull() || m.Resources.IsUnknown() {
		return nil, false
	}
	var v []AppResource
	d := m.Resources.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResources sets the value of the Resources field in App.
func (m *App) SetResources(ctx context.Context, v []AppResource) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["resources"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Resources = types.ListValueMust(t, vs)
}

// GetUserApiScopes returns the value of the UserApiScopes field in App as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *App) GetUserApiScopes(ctx context.Context) ([]types.String, bool) {
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

// SetUserApiScopes sets the value of the UserApiScopes field in App.
func (m *App) SetUserApiScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["user_api_scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.UserApiScopes = types.ListValueMust(t, vs)
}

type AppAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`

	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (to *AppAccessControlRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AppAccessControlRequest) {
}

func (to *AppAccessControlRequest) SyncFieldsDuringRead(ctx context.Context, from AppAccessControlRequest) {
}

func (m AppAccessControlRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppAccessControlRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AppAccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppAccessControlRequest
// only implements ToObjectValue() and Type().
func (m AppAccessControlRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_name":             m.GroupName,
			"permission_level":       m.PermissionLevel,
			"service_principal_name": m.ServicePrincipalName,
			"user_name":              m.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AppAccessControlRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type AppAccessControlResponse struct {
	// All permissions.
	AllPermissions types.List `tfsdk:"all_permissions"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name"`
	// name of the group
	GroupName types.String `tfsdk:"group_name"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (to *AppAccessControlResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AppAccessControlResponse) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (to *AppAccessControlResponse) SyncFieldsDuringRead(ctx context.Context, from AppAccessControlResponse) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (m AppAccessControlResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["all_permissions"] = attrs["all_permissions"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppAccessControlResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AppAccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(AppPermission{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppAccessControlResponse
// only implements ToObjectValue() and Type().
func (m AppAccessControlResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"all_permissions":        m.AllPermissions,
			"display_name":           m.DisplayName,
			"group_name":             m.GroupName,
			"service_principal_name": m.ServicePrincipalName,
			"user_name":              m.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AppAccessControlResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: AppPermission{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

// GetAllPermissions returns the value of the AllPermissions field in AppAccessControlResponse as
// a slice of AppPermission values.
// If the field is unknown or null, the boolean return value is false.
func (m *AppAccessControlResponse) GetAllPermissions(ctx context.Context) ([]AppPermission, bool) {
	if m.AllPermissions.IsNull() || m.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []AppPermission
	d := m.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in AppAccessControlResponse.
func (m *AppAccessControlResponse) SetAllPermissions(ctx context.Context, v []AppPermission) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllPermissions = types.ListValueMust(t, vs)
}

type AppDeployment struct {
	// The command with which to run the app. This will override the command
	// specified in the app.yaml file.
	Command types.List `tfsdk:"command"`
	// The creation time of the deployment. Formatted timestamp in ISO 6801.
	CreateTime types.String `tfsdk:"create_time"`
	// The email of the user creates the deployment.
	Creator types.String `tfsdk:"creator"`
	// The deployment artifacts for an app.
	DeploymentArtifacts types.Object `tfsdk:"deployment_artifacts"`
	// The unique id of the deployment.
	DeploymentId types.String `tfsdk:"deployment_id"`
	// The environment variables to set in the app runtime environment. This
	// will override the environment variables specified in the app.yaml file.
	EnvVars types.List `tfsdk:"env_vars"`
	// Git repository to use as the source for the app deployment.
	GitSource types.Object `tfsdk:"git_source"`
	// The mode of which the deployment will manage the source code.
	Mode types.String `tfsdk:"mode"`
	// The workspace file system path of the source code used to create the app
	// deployment. This is different from
	// `deployment_artifacts.source_code_path`, which is the path used by the
	// deployed app. The former refers to the original source code location of
	// the app in the workspace during deployment creation, whereas the latter
	// provides a system generated stable snapshotted source code path used by
	// the deployment.
	SourceCodePath types.String `tfsdk:"source_code_path"`
	// Status and status message of the deployment
	Status types.Object `tfsdk:"status"`
	// The update time of the deployment. Formatted timestamp in ISO 6801.
	UpdateTime types.String `tfsdk:"update_time"`
}

func (to *AppDeployment) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AppDeployment) {
	if !from.Command.IsNull() && !from.Command.IsUnknown() && to.Command.IsNull() && len(from.Command.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Command, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Command = from.Command
	}
	if !from.DeploymentArtifacts.IsNull() && !from.DeploymentArtifacts.IsUnknown() {
		if toDeploymentArtifacts, ok := to.GetDeploymentArtifacts(ctx); ok {
			if fromDeploymentArtifacts, ok := from.GetDeploymentArtifacts(ctx); ok {
				// Recursively sync the fields of DeploymentArtifacts
				toDeploymentArtifacts.SyncFieldsDuringCreateOrUpdate(ctx, fromDeploymentArtifacts)
				to.SetDeploymentArtifacts(ctx, toDeploymentArtifacts)
			}
		}
	}
	if !from.EnvVars.IsNull() && !from.EnvVars.IsUnknown() && to.EnvVars.IsNull() && len(from.EnvVars.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EnvVars, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EnvVars = from.EnvVars
	}
	if !from.GitSource.IsNull() && !from.GitSource.IsUnknown() {
		if toGitSource, ok := to.GetGitSource(ctx); ok {
			if fromGitSource, ok := from.GetGitSource(ctx); ok {
				// Recursively sync the fields of GitSource
				toGitSource.SyncFieldsDuringCreateOrUpdate(ctx, fromGitSource)
				to.SetGitSource(ctx, toGitSource)
			}
		}
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
}

func (to *AppDeployment) SyncFieldsDuringRead(ctx context.Context, from AppDeployment) {
	if !from.Command.IsNull() && !from.Command.IsUnknown() && to.Command.IsNull() && len(from.Command.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Command, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Command = from.Command
	}
	if !from.DeploymentArtifacts.IsNull() && !from.DeploymentArtifacts.IsUnknown() {
		if toDeploymentArtifacts, ok := to.GetDeploymentArtifacts(ctx); ok {
			if fromDeploymentArtifacts, ok := from.GetDeploymentArtifacts(ctx); ok {
				toDeploymentArtifacts.SyncFieldsDuringRead(ctx, fromDeploymentArtifacts)
				to.SetDeploymentArtifacts(ctx, toDeploymentArtifacts)
			}
		}
	}
	if !from.EnvVars.IsNull() && !from.EnvVars.IsUnknown() && to.EnvVars.IsNull() && len(from.EnvVars.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EnvVars, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EnvVars = from.EnvVars
	}
	if !from.GitSource.IsNull() && !from.GitSource.IsUnknown() {
		if toGitSource, ok := to.GetGitSource(ctx); ok {
			if fromGitSource, ok := from.GetGitSource(ctx); ok {
				toGitSource.SyncFieldsDuringRead(ctx, fromGitSource)
				to.SetGitSource(ctx, toGitSource)
			}
		}
	}
	if !from.Status.IsNull() && !from.Status.IsUnknown() {
		if toStatus, ok := to.GetStatus(ctx); ok {
			if fromStatus, ok := from.GetStatus(ctx); ok {
				toStatus.SyncFieldsDuringRead(ctx, fromStatus)
				to.SetStatus(ctx, toStatus)
			}
		}
	}
}

func (m AppDeployment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["command"] = attrs["command"].SetOptional()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["creator"] = attrs["creator"].SetComputed()
	attrs["deployment_artifacts"] = attrs["deployment_artifacts"].SetComputed()
	attrs["deployment_id"] = attrs["deployment_id"].SetOptional()
	attrs["env_vars"] = attrs["env_vars"].SetOptional()
	attrs["git_source"] = attrs["git_source"].SetOptional()
	attrs["mode"] = attrs["mode"].SetOptional()
	attrs["source_code_path"] = attrs["source_code_path"].SetOptional()
	attrs["status"] = attrs["status"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppDeployment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AppDeployment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"command":              reflect.TypeOf(types.String{}),
		"deployment_artifacts": reflect.TypeOf(AppDeploymentArtifacts{}),
		"env_vars":             reflect.TypeOf(EnvVar{}),
		"git_source":           reflect.TypeOf(GitSource{}),
		"status":               reflect.TypeOf(AppDeploymentStatus{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppDeployment
// only implements ToObjectValue() and Type().
func (m AppDeployment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"command":              m.Command,
			"create_time":          m.CreateTime,
			"creator":              m.Creator,
			"deployment_artifacts": m.DeploymentArtifacts,
			"deployment_id":        m.DeploymentId,
			"env_vars":             m.EnvVars,
			"git_source":           m.GitSource,
			"mode":                 m.Mode,
			"source_code_path":     m.SourceCodePath,
			"status":               m.Status,
			"update_time":          m.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AppDeployment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"command": basetypes.ListType{
				ElemType: types.StringType,
			},
			"create_time":          types.StringType,
			"creator":              types.StringType,
			"deployment_artifacts": AppDeploymentArtifacts{}.Type(ctx),
			"deployment_id":        types.StringType,
			"env_vars": basetypes.ListType{
				ElemType: EnvVar{}.Type(ctx),
			},
			"git_source":       GitSource{}.Type(ctx),
			"mode":             types.StringType,
			"source_code_path": types.StringType,
			"status":           AppDeploymentStatus{}.Type(ctx),
			"update_time":      types.StringType,
		},
	}
}

// GetCommand returns the value of the Command field in AppDeployment as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *AppDeployment) GetCommand(ctx context.Context) ([]types.String, bool) {
	if m.Command.IsNull() || m.Command.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Command.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCommand sets the value of the Command field in AppDeployment.
func (m *AppDeployment) SetCommand(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["command"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Command = types.ListValueMust(t, vs)
}

// GetDeploymentArtifacts returns the value of the DeploymentArtifacts field in AppDeployment as
// a AppDeploymentArtifacts value.
// If the field is unknown or null, the boolean return value is false.
func (m *AppDeployment) GetDeploymentArtifacts(ctx context.Context) (AppDeploymentArtifacts, bool) {
	var e AppDeploymentArtifacts
	if m.DeploymentArtifacts.IsNull() || m.DeploymentArtifacts.IsUnknown() {
		return e, false
	}
	var v AppDeploymentArtifacts
	d := m.DeploymentArtifacts.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDeploymentArtifacts sets the value of the DeploymentArtifacts field in AppDeployment.
func (m *AppDeployment) SetDeploymentArtifacts(ctx context.Context, v AppDeploymentArtifacts) {
	vs := v.ToObjectValue(ctx)
	m.DeploymentArtifacts = vs
}

// GetEnvVars returns the value of the EnvVars field in AppDeployment as
// a slice of EnvVar values.
// If the field is unknown or null, the boolean return value is false.
func (m *AppDeployment) GetEnvVars(ctx context.Context) ([]EnvVar, bool) {
	if m.EnvVars.IsNull() || m.EnvVars.IsUnknown() {
		return nil, false
	}
	var v []EnvVar
	d := m.EnvVars.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEnvVars sets the value of the EnvVars field in AppDeployment.
func (m *AppDeployment) SetEnvVars(ctx context.Context, v []EnvVar) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["env_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EnvVars = types.ListValueMust(t, vs)
}

// GetGitSource returns the value of the GitSource field in AppDeployment as
// a GitSource value.
// If the field is unknown or null, the boolean return value is false.
func (m *AppDeployment) GetGitSource(ctx context.Context) (GitSource, bool) {
	var e GitSource
	if m.GitSource.IsNull() || m.GitSource.IsUnknown() {
		return e, false
	}
	var v GitSource
	d := m.GitSource.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGitSource sets the value of the GitSource field in AppDeployment.
func (m *AppDeployment) SetGitSource(ctx context.Context, v GitSource) {
	vs := v.ToObjectValue(ctx)
	m.GitSource = vs
}

// GetStatus returns the value of the Status field in AppDeployment as
// a AppDeploymentStatus value.
// If the field is unknown or null, the boolean return value is false.
func (m *AppDeployment) GetStatus(ctx context.Context) (AppDeploymentStatus, bool) {
	var e AppDeploymentStatus
	if m.Status.IsNull() || m.Status.IsUnknown() {
		return e, false
	}
	var v AppDeploymentStatus
	d := m.Status.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatus sets the value of the Status field in AppDeployment.
func (m *AppDeployment) SetStatus(ctx context.Context, v AppDeploymentStatus) {
	vs := v.ToObjectValue(ctx)
	m.Status = vs
}

type AppDeploymentArtifacts struct {
	// The snapshotted workspace file system path of the source code loaded by
	// the deployed app.
	SourceCodePath types.String `tfsdk:"source_code_path"`
}

func (to *AppDeploymentArtifacts) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AppDeploymentArtifacts) {
}

func (to *AppDeploymentArtifacts) SyncFieldsDuringRead(ctx context.Context, from AppDeploymentArtifacts) {
}

func (m AppDeploymentArtifacts) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["source_code_path"] = attrs["source_code_path"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppDeploymentArtifacts.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AppDeploymentArtifacts) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppDeploymentArtifacts
// only implements ToObjectValue() and Type().
func (m AppDeploymentArtifacts) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"source_code_path": m.SourceCodePath,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AppDeploymentArtifacts) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"source_code_path": types.StringType,
		},
	}
}

type AppDeploymentStatus struct {
	// Message corresponding with the deployment state.
	Message types.String `tfsdk:"message"`
	// State of the deployment.
	State types.String `tfsdk:"state"`
}

func (to *AppDeploymentStatus) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AppDeploymentStatus) {
}

func (to *AppDeploymentStatus) SyncFieldsDuringRead(ctx context.Context, from AppDeploymentStatus) {
}

func (m AppDeploymentStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["message"] = attrs["message"].SetComputed()
	attrs["state"] = attrs["state"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppDeploymentStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AppDeploymentStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppDeploymentStatus
// only implements ToObjectValue() and Type().
func (m AppDeploymentStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"message": m.Message,
			"state":   m.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AppDeploymentStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"message": types.StringType,
			"state":   types.StringType,
		},
	}
}

// App manifest definition
type AppManifest struct {
	// Description of the app defined by manifest author / publisher
	Description types.String `tfsdk:"description"`
	// Name of the app defined by manifest author / publisher
	Name types.String `tfsdk:"name"`

	ResourceSpecs types.List `tfsdk:"resource_specs"`
	// The manifest schema version, for now only 1 is allowed
	Version types.Int64 `tfsdk:"version"`
}

func (to *AppManifest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AppManifest) {
	if !from.ResourceSpecs.IsNull() && !from.ResourceSpecs.IsUnknown() && to.ResourceSpecs.IsNull() && len(from.ResourceSpecs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ResourceSpecs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ResourceSpecs = from.ResourceSpecs
	}
}

func (to *AppManifest) SyncFieldsDuringRead(ctx context.Context, from AppManifest) {
	if !from.ResourceSpecs.IsNull() && !from.ResourceSpecs.IsUnknown() && to.ResourceSpecs.IsNull() && len(from.ResourceSpecs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ResourceSpecs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ResourceSpecs = from.ResourceSpecs
	}
}

func (m AppManifest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["resource_specs"] = attrs["resource_specs"].SetOptional()
	attrs["version"] = attrs["version"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppManifest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AppManifest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"resource_specs": reflect.TypeOf(AppManifestAppResourceSpec{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppManifest
// only implements ToObjectValue() and Type().
func (m AppManifest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":    m.Description,
			"name":           m.Name,
			"resource_specs": m.ResourceSpecs,
			"version":        m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AppManifest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"name":        types.StringType,
			"resource_specs": basetypes.ListType{
				ElemType: AppManifestAppResourceSpec{}.Type(ctx),
			},
			"version": types.Int64Type,
		},
	}
}

// GetResourceSpecs returns the value of the ResourceSpecs field in AppManifest as
// a slice of AppManifestAppResourceSpec values.
// If the field is unknown or null, the boolean return value is false.
func (m *AppManifest) GetResourceSpecs(ctx context.Context) ([]AppManifestAppResourceSpec, bool) {
	if m.ResourceSpecs.IsNull() || m.ResourceSpecs.IsUnknown() {
		return nil, false
	}
	var v []AppManifestAppResourceSpec
	d := m.ResourceSpecs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResourceSpecs sets the value of the ResourceSpecs field in AppManifest.
func (m *AppManifest) SetResourceSpecs(ctx context.Context, v []AppManifestAppResourceSpec) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["resource_specs"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ResourceSpecs = types.ListValueMust(t, vs)
}

type AppManifestAppResourceExperimentSpec struct {
	Permission types.String `tfsdk:"permission"`
}

func (to *AppManifestAppResourceExperimentSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AppManifestAppResourceExperimentSpec) {
}

func (to *AppManifestAppResourceExperimentSpec) SyncFieldsDuringRead(ctx context.Context, from AppManifestAppResourceExperimentSpec) {
}

func (m AppManifestAppResourceExperimentSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission"] = attrs["permission"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppManifestAppResourceExperimentSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AppManifestAppResourceExperimentSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppManifestAppResourceExperimentSpec
// only implements ToObjectValue() and Type().
func (m AppManifestAppResourceExperimentSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission": m.Permission,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AppManifestAppResourceExperimentSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission": types.StringType,
		},
	}
}

type AppManifestAppResourceJobSpec struct {
	// Permissions to grant on the Job. Supported permissions are: "CAN_MANAGE",
	// "IS_OWNER", "CAN_MANAGE_RUN", "CAN_VIEW".
	Permission types.String `tfsdk:"permission"`
}

func (to *AppManifestAppResourceJobSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AppManifestAppResourceJobSpec) {
}

func (to *AppManifestAppResourceJobSpec) SyncFieldsDuringRead(ctx context.Context, from AppManifestAppResourceJobSpec) {
}

func (m AppManifestAppResourceJobSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission"] = attrs["permission"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppManifestAppResourceJobSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AppManifestAppResourceJobSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppManifestAppResourceJobSpec
// only implements ToObjectValue() and Type().
func (m AppManifestAppResourceJobSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission": m.Permission,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AppManifestAppResourceJobSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission": types.StringType,
		},
	}
}

type AppManifestAppResourceSecretSpec struct {
	// Permission to grant on the secret scope. For secrets, only one permission
	// is allowed. Permission must be one of: "READ", "WRITE", "MANAGE".
	Permission types.String `tfsdk:"permission"`
}

func (to *AppManifestAppResourceSecretSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AppManifestAppResourceSecretSpec) {
}

func (to *AppManifestAppResourceSecretSpec) SyncFieldsDuringRead(ctx context.Context, from AppManifestAppResourceSecretSpec) {
}

func (m AppManifestAppResourceSecretSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission"] = attrs["permission"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppManifestAppResourceSecretSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AppManifestAppResourceSecretSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppManifestAppResourceSecretSpec
// only implements ToObjectValue() and Type().
func (m AppManifestAppResourceSecretSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission": m.Permission,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AppManifestAppResourceSecretSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission": types.StringType,
		},
	}
}

type AppManifestAppResourceServingEndpointSpec struct {
	// Permission to grant on the serving endpoint. Supported permissions are:
	// "CAN_MANAGE", "CAN_QUERY", "CAN_VIEW".
	Permission types.String `tfsdk:"permission"`
}

func (to *AppManifestAppResourceServingEndpointSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AppManifestAppResourceServingEndpointSpec) {
}

func (to *AppManifestAppResourceServingEndpointSpec) SyncFieldsDuringRead(ctx context.Context, from AppManifestAppResourceServingEndpointSpec) {
}

func (m AppManifestAppResourceServingEndpointSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission"] = attrs["permission"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppManifestAppResourceServingEndpointSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AppManifestAppResourceServingEndpointSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppManifestAppResourceServingEndpointSpec
// only implements ToObjectValue() and Type().
func (m AppManifestAppResourceServingEndpointSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission": m.Permission,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AppManifestAppResourceServingEndpointSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission": types.StringType,
		},
	}
}

// AppResource related fields are copied from app.proto but excludes resource
// identifiers (e.g. name, id, key, scope, etc.)
type AppManifestAppResourceSpec struct {
	// Description of the App Resource.
	Description types.String `tfsdk:"description"`

	ExperimentSpec types.Object `tfsdk:"experiment_spec"`

	JobSpec types.Object `tfsdk:"job_spec"`
	// Name of the App Resource.
	Name types.String `tfsdk:"name"`

	SecretSpec types.Object `tfsdk:"secret_spec"`

	ServingEndpointSpec types.Object `tfsdk:"serving_endpoint_spec"`

	SqlWarehouseSpec types.Object `tfsdk:"sql_warehouse_spec"`

	UcSecurableSpec types.Object `tfsdk:"uc_securable_spec"`
}

func (to *AppManifestAppResourceSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AppManifestAppResourceSpec) {
	if !from.ExperimentSpec.IsNull() && !from.ExperimentSpec.IsUnknown() {
		if toExperimentSpec, ok := to.GetExperimentSpec(ctx); ok {
			if fromExperimentSpec, ok := from.GetExperimentSpec(ctx); ok {
				// Recursively sync the fields of ExperimentSpec
				toExperimentSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromExperimentSpec)
				to.SetExperimentSpec(ctx, toExperimentSpec)
			}
		}
	}
	if !from.JobSpec.IsNull() && !from.JobSpec.IsUnknown() {
		if toJobSpec, ok := to.GetJobSpec(ctx); ok {
			if fromJobSpec, ok := from.GetJobSpec(ctx); ok {
				// Recursively sync the fields of JobSpec
				toJobSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromJobSpec)
				to.SetJobSpec(ctx, toJobSpec)
			}
		}
	}
	if !from.SecretSpec.IsNull() && !from.SecretSpec.IsUnknown() {
		if toSecretSpec, ok := to.GetSecretSpec(ctx); ok {
			if fromSecretSpec, ok := from.GetSecretSpec(ctx); ok {
				// Recursively sync the fields of SecretSpec
				toSecretSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromSecretSpec)
				to.SetSecretSpec(ctx, toSecretSpec)
			}
		}
	}
	if !from.ServingEndpointSpec.IsNull() && !from.ServingEndpointSpec.IsUnknown() {
		if toServingEndpointSpec, ok := to.GetServingEndpointSpec(ctx); ok {
			if fromServingEndpointSpec, ok := from.GetServingEndpointSpec(ctx); ok {
				// Recursively sync the fields of ServingEndpointSpec
				toServingEndpointSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromServingEndpointSpec)
				to.SetServingEndpointSpec(ctx, toServingEndpointSpec)
			}
		}
	}
	if !from.SqlWarehouseSpec.IsNull() && !from.SqlWarehouseSpec.IsUnknown() {
		if toSqlWarehouseSpec, ok := to.GetSqlWarehouseSpec(ctx); ok {
			if fromSqlWarehouseSpec, ok := from.GetSqlWarehouseSpec(ctx); ok {
				// Recursively sync the fields of SqlWarehouseSpec
				toSqlWarehouseSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromSqlWarehouseSpec)
				to.SetSqlWarehouseSpec(ctx, toSqlWarehouseSpec)
			}
		}
	}
	if !from.UcSecurableSpec.IsNull() && !from.UcSecurableSpec.IsUnknown() {
		if toUcSecurableSpec, ok := to.GetUcSecurableSpec(ctx); ok {
			if fromUcSecurableSpec, ok := from.GetUcSecurableSpec(ctx); ok {
				// Recursively sync the fields of UcSecurableSpec
				toUcSecurableSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromUcSecurableSpec)
				to.SetUcSecurableSpec(ctx, toUcSecurableSpec)
			}
		}
	}
}

func (to *AppManifestAppResourceSpec) SyncFieldsDuringRead(ctx context.Context, from AppManifestAppResourceSpec) {
	if !from.ExperimentSpec.IsNull() && !from.ExperimentSpec.IsUnknown() {
		if toExperimentSpec, ok := to.GetExperimentSpec(ctx); ok {
			if fromExperimentSpec, ok := from.GetExperimentSpec(ctx); ok {
				toExperimentSpec.SyncFieldsDuringRead(ctx, fromExperimentSpec)
				to.SetExperimentSpec(ctx, toExperimentSpec)
			}
		}
	}
	if !from.JobSpec.IsNull() && !from.JobSpec.IsUnknown() {
		if toJobSpec, ok := to.GetJobSpec(ctx); ok {
			if fromJobSpec, ok := from.GetJobSpec(ctx); ok {
				toJobSpec.SyncFieldsDuringRead(ctx, fromJobSpec)
				to.SetJobSpec(ctx, toJobSpec)
			}
		}
	}
	if !from.SecretSpec.IsNull() && !from.SecretSpec.IsUnknown() {
		if toSecretSpec, ok := to.GetSecretSpec(ctx); ok {
			if fromSecretSpec, ok := from.GetSecretSpec(ctx); ok {
				toSecretSpec.SyncFieldsDuringRead(ctx, fromSecretSpec)
				to.SetSecretSpec(ctx, toSecretSpec)
			}
		}
	}
	if !from.ServingEndpointSpec.IsNull() && !from.ServingEndpointSpec.IsUnknown() {
		if toServingEndpointSpec, ok := to.GetServingEndpointSpec(ctx); ok {
			if fromServingEndpointSpec, ok := from.GetServingEndpointSpec(ctx); ok {
				toServingEndpointSpec.SyncFieldsDuringRead(ctx, fromServingEndpointSpec)
				to.SetServingEndpointSpec(ctx, toServingEndpointSpec)
			}
		}
	}
	if !from.SqlWarehouseSpec.IsNull() && !from.SqlWarehouseSpec.IsUnknown() {
		if toSqlWarehouseSpec, ok := to.GetSqlWarehouseSpec(ctx); ok {
			if fromSqlWarehouseSpec, ok := from.GetSqlWarehouseSpec(ctx); ok {
				toSqlWarehouseSpec.SyncFieldsDuringRead(ctx, fromSqlWarehouseSpec)
				to.SetSqlWarehouseSpec(ctx, toSqlWarehouseSpec)
			}
		}
	}
	if !from.UcSecurableSpec.IsNull() && !from.UcSecurableSpec.IsUnknown() {
		if toUcSecurableSpec, ok := to.GetUcSecurableSpec(ctx); ok {
			if fromUcSecurableSpec, ok := from.GetUcSecurableSpec(ctx); ok {
				toUcSecurableSpec.SyncFieldsDuringRead(ctx, fromUcSecurableSpec)
				to.SetUcSecurableSpec(ctx, toUcSecurableSpec)
			}
		}
	}
}

func (m AppManifestAppResourceSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["experiment_spec"] = attrs["experiment_spec"].SetOptional()
	attrs["job_spec"] = attrs["job_spec"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["secret_spec"] = attrs["secret_spec"].SetOptional()
	attrs["serving_endpoint_spec"] = attrs["serving_endpoint_spec"].SetOptional()
	attrs["sql_warehouse_spec"] = attrs["sql_warehouse_spec"].SetOptional()
	attrs["uc_securable_spec"] = attrs["uc_securable_spec"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppManifestAppResourceSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AppManifestAppResourceSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"experiment_spec":       reflect.TypeOf(AppManifestAppResourceExperimentSpec{}),
		"job_spec":              reflect.TypeOf(AppManifestAppResourceJobSpec{}),
		"secret_spec":           reflect.TypeOf(AppManifestAppResourceSecretSpec{}),
		"serving_endpoint_spec": reflect.TypeOf(AppManifestAppResourceServingEndpointSpec{}),
		"sql_warehouse_spec":    reflect.TypeOf(AppManifestAppResourceSqlWarehouseSpec{}),
		"uc_securable_spec":     reflect.TypeOf(AppManifestAppResourceUcSecurableSpec{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppManifestAppResourceSpec
// only implements ToObjectValue() and Type().
func (m AppManifestAppResourceSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":           m.Description,
			"experiment_spec":       m.ExperimentSpec,
			"job_spec":              m.JobSpec,
			"name":                  m.Name,
			"secret_spec":           m.SecretSpec,
			"serving_endpoint_spec": m.ServingEndpointSpec,
			"sql_warehouse_spec":    m.SqlWarehouseSpec,
			"uc_securable_spec":     m.UcSecurableSpec,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AppManifestAppResourceSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":           types.StringType,
			"experiment_spec":       AppManifestAppResourceExperimentSpec{}.Type(ctx),
			"job_spec":              AppManifestAppResourceJobSpec{}.Type(ctx),
			"name":                  types.StringType,
			"secret_spec":           AppManifestAppResourceSecretSpec{}.Type(ctx),
			"serving_endpoint_spec": AppManifestAppResourceServingEndpointSpec{}.Type(ctx),
			"sql_warehouse_spec":    AppManifestAppResourceSqlWarehouseSpec{}.Type(ctx),
			"uc_securable_spec":     AppManifestAppResourceUcSecurableSpec{}.Type(ctx),
		},
	}
}

// GetExperimentSpec returns the value of the ExperimentSpec field in AppManifestAppResourceSpec as
// a AppManifestAppResourceExperimentSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *AppManifestAppResourceSpec) GetExperimentSpec(ctx context.Context) (AppManifestAppResourceExperimentSpec, bool) {
	var e AppManifestAppResourceExperimentSpec
	if m.ExperimentSpec.IsNull() || m.ExperimentSpec.IsUnknown() {
		return e, false
	}
	var v AppManifestAppResourceExperimentSpec
	d := m.ExperimentSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExperimentSpec sets the value of the ExperimentSpec field in AppManifestAppResourceSpec.
func (m *AppManifestAppResourceSpec) SetExperimentSpec(ctx context.Context, v AppManifestAppResourceExperimentSpec) {
	vs := v.ToObjectValue(ctx)
	m.ExperimentSpec = vs
}

// GetJobSpec returns the value of the JobSpec field in AppManifestAppResourceSpec as
// a AppManifestAppResourceJobSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *AppManifestAppResourceSpec) GetJobSpec(ctx context.Context) (AppManifestAppResourceJobSpec, bool) {
	var e AppManifestAppResourceJobSpec
	if m.JobSpec.IsNull() || m.JobSpec.IsUnknown() {
		return e, false
	}
	var v AppManifestAppResourceJobSpec
	d := m.JobSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetJobSpec sets the value of the JobSpec field in AppManifestAppResourceSpec.
func (m *AppManifestAppResourceSpec) SetJobSpec(ctx context.Context, v AppManifestAppResourceJobSpec) {
	vs := v.ToObjectValue(ctx)
	m.JobSpec = vs
}

// GetSecretSpec returns the value of the SecretSpec field in AppManifestAppResourceSpec as
// a AppManifestAppResourceSecretSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *AppManifestAppResourceSpec) GetSecretSpec(ctx context.Context) (AppManifestAppResourceSecretSpec, bool) {
	var e AppManifestAppResourceSecretSpec
	if m.SecretSpec.IsNull() || m.SecretSpec.IsUnknown() {
		return e, false
	}
	var v AppManifestAppResourceSecretSpec
	d := m.SecretSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSecretSpec sets the value of the SecretSpec field in AppManifestAppResourceSpec.
func (m *AppManifestAppResourceSpec) SetSecretSpec(ctx context.Context, v AppManifestAppResourceSecretSpec) {
	vs := v.ToObjectValue(ctx)
	m.SecretSpec = vs
}

// GetServingEndpointSpec returns the value of the ServingEndpointSpec field in AppManifestAppResourceSpec as
// a AppManifestAppResourceServingEndpointSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *AppManifestAppResourceSpec) GetServingEndpointSpec(ctx context.Context) (AppManifestAppResourceServingEndpointSpec, bool) {
	var e AppManifestAppResourceServingEndpointSpec
	if m.ServingEndpointSpec.IsNull() || m.ServingEndpointSpec.IsUnknown() {
		return e, false
	}
	var v AppManifestAppResourceServingEndpointSpec
	d := m.ServingEndpointSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServingEndpointSpec sets the value of the ServingEndpointSpec field in AppManifestAppResourceSpec.
func (m *AppManifestAppResourceSpec) SetServingEndpointSpec(ctx context.Context, v AppManifestAppResourceServingEndpointSpec) {
	vs := v.ToObjectValue(ctx)
	m.ServingEndpointSpec = vs
}

// GetSqlWarehouseSpec returns the value of the SqlWarehouseSpec field in AppManifestAppResourceSpec as
// a AppManifestAppResourceSqlWarehouseSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *AppManifestAppResourceSpec) GetSqlWarehouseSpec(ctx context.Context) (AppManifestAppResourceSqlWarehouseSpec, bool) {
	var e AppManifestAppResourceSqlWarehouseSpec
	if m.SqlWarehouseSpec.IsNull() || m.SqlWarehouseSpec.IsUnknown() {
		return e, false
	}
	var v AppManifestAppResourceSqlWarehouseSpec
	d := m.SqlWarehouseSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSqlWarehouseSpec sets the value of the SqlWarehouseSpec field in AppManifestAppResourceSpec.
func (m *AppManifestAppResourceSpec) SetSqlWarehouseSpec(ctx context.Context, v AppManifestAppResourceSqlWarehouseSpec) {
	vs := v.ToObjectValue(ctx)
	m.SqlWarehouseSpec = vs
}

// GetUcSecurableSpec returns the value of the UcSecurableSpec field in AppManifestAppResourceSpec as
// a AppManifestAppResourceUcSecurableSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *AppManifestAppResourceSpec) GetUcSecurableSpec(ctx context.Context) (AppManifestAppResourceUcSecurableSpec, bool) {
	var e AppManifestAppResourceUcSecurableSpec
	if m.UcSecurableSpec.IsNull() || m.UcSecurableSpec.IsUnknown() {
		return e, false
	}
	var v AppManifestAppResourceUcSecurableSpec
	d := m.UcSecurableSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUcSecurableSpec sets the value of the UcSecurableSpec field in AppManifestAppResourceSpec.
func (m *AppManifestAppResourceSpec) SetUcSecurableSpec(ctx context.Context, v AppManifestAppResourceUcSecurableSpec) {
	vs := v.ToObjectValue(ctx)
	m.UcSecurableSpec = vs
}

type AppManifestAppResourceSqlWarehouseSpec struct {
	// Permission to grant on the SQL warehouse. Supported permissions are:
	// "CAN_MANAGE", "CAN_USE", "IS_OWNER".
	Permission types.String `tfsdk:"permission"`
}

func (to *AppManifestAppResourceSqlWarehouseSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AppManifestAppResourceSqlWarehouseSpec) {
}

func (to *AppManifestAppResourceSqlWarehouseSpec) SyncFieldsDuringRead(ctx context.Context, from AppManifestAppResourceSqlWarehouseSpec) {
}

func (m AppManifestAppResourceSqlWarehouseSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission"] = attrs["permission"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppManifestAppResourceSqlWarehouseSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AppManifestAppResourceSqlWarehouseSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppManifestAppResourceSqlWarehouseSpec
// only implements ToObjectValue() and Type().
func (m AppManifestAppResourceSqlWarehouseSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission": m.Permission,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AppManifestAppResourceSqlWarehouseSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission": types.StringType,
		},
	}
}

type AppManifestAppResourceUcSecurableSpec struct {
	Permission types.String `tfsdk:"permission"`

	SecurableType types.String `tfsdk:"securable_type"`
}

func (to *AppManifestAppResourceUcSecurableSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AppManifestAppResourceUcSecurableSpec) {
}

func (to *AppManifestAppResourceUcSecurableSpec) SyncFieldsDuringRead(ctx context.Context, from AppManifestAppResourceUcSecurableSpec) {
}

func (m AppManifestAppResourceUcSecurableSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission"] = attrs["permission"].SetRequired()
	attrs["securable_type"] = attrs["securable_type"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppManifestAppResourceUcSecurableSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AppManifestAppResourceUcSecurableSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppManifestAppResourceUcSecurableSpec
// only implements ToObjectValue() and Type().
func (m AppManifestAppResourceUcSecurableSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission":     m.Permission,
			"securable_type": m.SecurableType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AppManifestAppResourceUcSecurableSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission":     types.StringType,
			"securable_type": types.StringType,
		},
	}
}

type AppPermission struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *AppPermission) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AppPermission) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (to *AppPermission) SyncFieldsDuringRead(ctx context.Context, from AppPermission) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (m AppPermission) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["inherited"] = attrs["inherited"].SetOptional()
	attrs["inherited_from_object"] = attrs["inherited_from_object"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppPermission.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AppPermission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppPermission
// only implements ToObjectValue() and Type().
func (m AppPermission) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             m.Inherited,
			"inherited_from_object": m.InheritedFromObject,
			"permission_level":      m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AppPermission) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"inherited": types.BoolType,
			"inherited_from_object": basetypes.ListType{
				ElemType: types.StringType,
			},
			"permission_level": types.StringType,
		},
	}
}

// GetInheritedFromObject returns the value of the InheritedFromObject field in AppPermission as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *AppPermission) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
	if m.InheritedFromObject.IsNull() || m.InheritedFromObject.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.InheritedFromObject.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInheritedFromObject sets the value of the InheritedFromObject field in AppPermission.
func (m *AppPermission) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InheritedFromObject = types.ListValueMust(t, vs)
}

type AppPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (to *AppPermissions) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AppPermissions) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *AppPermissions) SyncFieldsDuringRead(ctx context.Context, from AppPermissions) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m AppPermissions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["object_id"] = attrs["object_id"].SetOptional()
	attrs["object_type"] = attrs["object_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppPermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AppPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(AppAccessControlResponse{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppPermissions
// only implements ToObjectValue() and Type().
func (m AppPermissions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"object_id":           m.ObjectId,
			"object_type":         m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AppPermissions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: AppAccessControlResponse{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in AppPermissions as
// a slice of AppAccessControlResponse values.
// If the field is unknown or null, the boolean return value is false.
func (m *AppPermissions) GetAccessControlList(ctx context.Context) ([]AppAccessControlResponse, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []AppAccessControlResponse
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in AppPermissions.
func (m *AppPermissions) SetAccessControlList(ctx context.Context, v []AppAccessControlResponse) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

type AppPermissionsDescription struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *AppPermissionsDescription) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AppPermissionsDescription) {
}

func (to *AppPermissionsDescription) SyncFieldsDuringRead(ctx context.Context, from AppPermissionsDescription) {
}

func (m AppPermissionsDescription) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppPermissionsDescription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AppPermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppPermissionsDescription
// only implements ToObjectValue() and Type().
func (m AppPermissionsDescription) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      m.Description,
			"permission_level": m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AppPermissionsDescription) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type AppPermissionsRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// The app for which to get or manage permissions.
	AppName types.String `tfsdk:"-"`
}

func (to *AppPermissionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AppPermissionsRequest) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *AppPermissionsRequest) SyncFieldsDuringRead(ctx context.Context, from AppPermissionsRequest) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m AppPermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["app_name"] = attrs["app_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AppPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(AppAccessControlRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppPermissionsRequest
// only implements ToObjectValue() and Type().
func (m AppPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"app_name":            m.AppName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AppPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: AppAccessControlRequest{}.Type(ctx),
			},
			"app_name": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in AppPermissionsRequest as
// a slice of AppAccessControlRequest values.
// If the field is unknown or null, the boolean return value is false.
func (m *AppPermissionsRequest) GetAccessControlList(ctx context.Context) ([]AppAccessControlRequest, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []AppAccessControlRequest
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in AppPermissionsRequest.
func (m *AppPermissionsRequest) SetAccessControlList(ctx context.Context, v []AppAccessControlRequest) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

type AppResource struct {
	Database types.Object `tfsdk:"database"`
	// Description of the App Resource.
	Description types.String `tfsdk:"description"`

	Experiment types.Object `tfsdk:"experiment"`

	GenieSpace types.Object `tfsdk:"genie_space"`

	Job types.Object `tfsdk:"job"`
	// Name of the App Resource.
	Name types.String `tfsdk:"name"`

	Secret types.Object `tfsdk:"secret"`

	ServingEndpoint types.Object `tfsdk:"serving_endpoint"`

	SqlWarehouse types.Object `tfsdk:"sql_warehouse"`

	UcSecurable types.Object `tfsdk:"uc_securable"`
}

func (to *AppResource) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AppResource) {
	if !from.Database.IsNull() && !from.Database.IsUnknown() {
		if toDatabase, ok := to.GetDatabase(ctx); ok {
			if fromDatabase, ok := from.GetDatabase(ctx); ok {
				// Recursively sync the fields of Database
				toDatabase.SyncFieldsDuringCreateOrUpdate(ctx, fromDatabase)
				to.SetDatabase(ctx, toDatabase)
			}
		}
	}
	if !from.Experiment.IsNull() && !from.Experiment.IsUnknown() {
		if toExperiment, ok := to.GetExperiment(ctx); ok {
			if fromExperiment, ok := from.GetExperiment(ctx); ok {
				// Recursively sync the fields of Experiment
				toExperiment.SyncFieldsDuringCreateOrUpdate(ctx, fromExperiment)
				to.SetExperiment(ctx, toExperiment)
			}
		}
	}
	if !from.GenieSpace.IsNull() && !from.GenieSpace.IsUnknown() {
		if toGenieSpace, ok := to.GetGenieSpace(ctx); ok {
			if fromGenieSpace, ok := from.GetGenieSpace(ctx); ok {
				// Recursively sync the fields of GenieSpace
				toGenieSpace.SyncFieldsDuringCreateOrUpdate(ctx, fromGenieSpace)
				to.SetGenieSpace(ctx, toGenieSpace)
			}
		}
	}
	if !from.Job.IsNull() && !from.Job.IsUnknown() {
		if toJob, ok := to.GetJob(ctx); ok {
			if fromJob, ok := from.GetJob(ctx); ok {
				// Recursively sync the fields of Job
				toJob.SyncFieldsDuringCreateOrUpdate(ctx, fromJob)
				to.SetJob(ctx, toJob)
			}
		}
	}
	if !from.Secret.IsNull() && !from.Secret.IsUnknown() {
		if toSecret, ok := to.GetSecret(ctx); ok {
			if fromSecret, ok := from.GetSecret(ctx); ok {
				// Recursively sync the fields of Secret
				toSecret.SyncFieldsDuringCreateOrUpdate(ctx, fromSecret)
				to.SetSecret(ctx, toSecret)
			}
		}
	}
	if !from.ServingEndpoint.IsNull() && !from.ServingEndpoint.IsUnknown() {
		if toServingEndpoint, ok := to.GetServingEndpoint(ctx); ok {
			if fromServingEndpoint, ok := from.GetServingEndpoint(ctx); ok {
				// Recursively sync the fields of ServingEndpoint
				toServingEndpoint.SyncFieldsDuringCreateOrUpdate(ctx, fromServingEndpoint)
				to.SetServingEndpoint(ctx, toServingEndpoint)
			}
		}
	}
	if !from.SqlWarehouse.IsNull() && !from.SqlWarehouse.IsUnknown() {
		if toSqlWarehouse, ok := to.GetSqlWarehouse(ctx); ok {
			if fromSqlWarehouse, ok := from.GetSqlWarehouse(ctx); ok {
				// Recursively sync the fields of SqlWarehouse
				toSqlWarehouse.SyncFieldsDuringCreateOrUpdate(ctx, fromSqlWarehouse)
				to.SetSqlWarehouse(ctx, toSqlWarehouse)
			}
		}
	}
	if !from.UcSecurable.IsNull() && !from.UcSecurable.IsUnknown() {
		if toUcSecurable, ok := to.GetUcSecurable(ctx); ok {
			if fromUcSecurable, ok := from.GetUcSecurable(ctx); ok {
				// Recursively sync the fields of UcSecurable
				toUcSecurable.SyncFieldsDuringCreateOrUpdate(ctx, fromUcSecurable)
				to.SetUcSecurable(ctx, toUcSecurable)
			}
		}
	}
}

func (to *AppResource) SyncFieldsDuringRead(ctx context.Context, from AppResource) {
	if !from.Database.IsNull() && !from.Database.IsUnknown() {
		if toDatabase, ok := to.GetDatabase(ctx); ok {
			if fromDatabase, ok := from.GetDatabase(ctx); ok {
				toDatabase.SyncFieldsDuringRead(ctx, fromDatabase)
				to.SetDatabase(ctx, toDatabase)
			}
		}
	}
	if !from.Experiment.IsNull() && !from.Experiment.IsUnknown() {
		if toExperiment, ok := to.GetExperiment(ctx); ok {
			if fromExperiment, ok := from.GetExperiment(ctx); ok {
				toExperiment.SyncFieldsDuringRead(ctx, fromExperiment)
				to.SetExperiment(ctx, toExperiment)
			}
		}
	}
	if !from.GenieSpace.IsNull() && !from.GenieSpace.IsUnknown() {
		if toGenieSpace, ok := to.GetGenieSpace(ctx); ok {
			if fromGenieSpace, ok := from.GetGenieSpace(ctx); ok {
				toGenieSpace.SyncFieldsDuringRead(ctx, fromGenieSpace)
				to.SetGenieSpace(ctx, toGenieSpace)
			}
		}
	}
	if !from.Job.IsNull() && !from.Job.IsUnknown() {
		if toJob, ok := to.GetJob(ctx); ok {
			if fromJob, ok := from.GetJob(ctx); ok {
				toJob.SyncFieldsDuringRead(ctx, fromJob)
				to.SetJob(ctx, toJob)
			}
		}
	}
	if !from.Secret.IsNull() && !from.Secret.IsUnknown() {
		if toSecret, ok := to.GetSecret(ctx); ok {
			if fromSecret, ok := from.GetSecret(ctx); ok {
				toSecret.SyncFieldsDuringRead(ctx, fromSecret)
				to.SetSecret(ctx, toSecret)
			}
		}
	}
	if !from.ServingEndpoint.IsNull() && !from.ServingEndpoint.IsUnknown() {
		if toServingEndpoint, ok := to.GetServingEndpoint(ctx); ok {
			if fromServingEndpoint, ok := from.GetServingEndpoint(ctx); ok {
				toServingEndpoint.SyncFieldsDuringRead(ctx, fromServingEndpoint)
				to.SetServingEndpoint(ctx, toServingEndpoint)
			}
		}
	}
	if !from.SqlWarehouse.IsNull() && !from.SqlWarehouse.IsUnknown() {
		if toSqlWarehouse, ok := to.GetSqlWarehouse(ctx); ok {
			if fromSqlWarehouse, ok := from.GetSqlWarehouse(ctx); ok {
				toSqlWarehouse.SyncFieldsDuringRead(ctx, fromSqlWarehouse)
				to.SetSqlWarehouse(ctx, toSqlWarehouse)
			}
		}
	}
	if !from.UcSecurable.IsNull() && !from.UcSecurable.IsUnknown() {
		if toUcSecurable, ok := to.GetUcSecurable(ctx); ok {
			if fromUcSecurable, ok := from.GetUcSecurable(ctx); ok {
				toUcSecurable.SyncFieldsDuringRead(ctx, fromUcSecurable)
				to.SetUcSecurable(ctx, toUcSecurable)
			}
		}
	}
}

func (m AppResource) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["database"] = attrs["database"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["experiment"] = attrs["experiment"].SetOptional()
	attrs["genie_space"] = attrs["genie_space"].SetOptional()
	attrs["job"] = attrs["job"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["secret"] = attrs["secret"].SetOptional()
	attrs["serving_endpoint"] = attrs["serving_endpoint"].SetOptional()
	attrs["sql_warehouse"] = attrs["sql_warehouse"].SetOptional()
	attrs["uc_securable"] = attrs["uc_securable"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppResource.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AppResource) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database":         reflect.TypeOf(AppResourceDatabase{}),
		"experiment":       reflect.TypeOf(AppResourceExperiment{}),
		"genie_space":      reflect.TypeOf(AppResourceGenieSpace{}),
		"job":              reflect.TypeOf(AppResourceJob{}),
		"secret":           reflect.TypeOf(AppResourceSecret{}),
		"serving_endpoint": reflect.TypeOf(AppResourceServingEndpoint{}),
		"sql_warehouse":    reflect.TypeOf(AppResourceSqlWarehouse{}),
		"uc_securable":     reflect.TypeOf(AppResourceUcSecurable{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppResource
// only implements ToObjectValue() and Type().
func (m AppResource) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database":         m.Database,
			"description":      m.Description,
			"experiment":       m.Experiment,
			"genie_space":      m.GenieSpace,
			"job":              m.Job,
			"name":             m.Name,
			"secret":           m.Secret,
			"serving_endpoint": m.ServingEndpoint,
			"sql_warehouse":    m.SqlWarehouse,
			"uc_securable":     m.UcSecurable,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AppResource) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"database":         AppResourceDatabase{}.Type(ctx),
			"description":      types.StringType,
			"experiment":       AppResourceExperiment{}.Type(ctx),
			"genie_space":      AppResourceGenieSpace{}.Type(ctx),
			"job":              AppResourceJob{}.Type(ctx),
			"name":             types.StringType,
			"secret":           AppResourceSecret{}.Type(ctx),
			"serving_endpoint": AppResourceServingEndpoint{}.Type(ctx),
			"sql_warehouse":    AppResourceSqlWarehouse{}.Type(ctx),
			"uc_securable":     AppResourceUcSecurable{}.Type(ctx),
		},
	}
}

// GetDatabase returns the value of the Database field in AppResource as
// a AppResourceDatabase value.
// If the field is unknown or null, the boolean return value is false.
func (m *AppResource) GetDatabase(ctx context.Context) (AppResourceDatabase, bool) {
	var e AppResourceDatabase
	if m.Database.IsNull() || m.Database.IsUnknown() {
		return e, false
	}
	var v AppResourceDatabase
	d := m.Database.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatabase sets the value of the Database field in AppResource.
func (m *AppResource) SetDatabase(ctx context.Context, v AppResourceDatabase) {
	vs := v.ToObjectValue(ctx)
	m.Database = vs
}

// GetExperiment returns the value of the Experiment field in AppResource as
// a AppResourceExperiment value.
// If the field is unknown or null, the boolean return value is false.
func (m *AppResource) GetExperiment(ctx context.Context) (AppResourceExperiment, bool) {
	var e AppResourceExperiment
	if m.Experiment.IsNull() || m.Experiment.IsUnknown() {
		return e, false
	}
	var v AppResourceExperiment
	d := m.Experiment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExperiment sets the value of the Experiment field in AppResource.
func (m *AppResource) SetExperiment(ctx context.Context, v AppResourceExperiment) {
	vs := v.ToObjectValue(ctx)
	m.Experiment = vs
}

// GetGenieSpace returns the value of the GenieSpace field in AppResource as
// a AppResourceGenieSpace value.
// If the field is unknown or null, the boolean return value is false.
func (m *AppResource) GetGenieSpace(ctx context.Context) (AppResourceGenieSpace, bool) {
	var e AppResourceGenieSpace
	if m.GenieSpace.IsNull() || m.GenieSpace.IsUnknown() {
		return e, false
	}
	var v AppResourceGenieSpace
	d := m.GenieSpace.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGenieSpace sets the value of the GenieSpace field in AppResource.
func (m *AppResource) SetGenieSpace(ctx context.Context, v AppResourceGenieSpace) {
	vs := v.ToObjectValue(ctx)
	m.GenieSpace = vs
}

// GetJob returns the value of the Job field in AppResource as
// a AppResourceJob value.
// If the field is unknown or null, the boolean return value is false.
func (m *AppResource) GetJob(ctx context.Context) (AppResourceJob, bool) {
	var e AppResourceJob
	if m.Job.IsNull() || m.Job.IsUnknown() {
		return e, false
	}
	var v AppResourceJob
	d := m.Job.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetJob sets the value of the Job field in AppResource.
func (m *AppResource) SetJob(ctx context.Context, v AppResourceJob) {
	vs := v.ToObjectValue(ctx)
	m.Job = vs
}

// GetSecret returns the value of the Secret field in AppResource as
// a AppResourceSecret value.
// If the field is unknown or null, the boolean return value is false.
func (m *AppResource) GetSecret(ctx context.Context) (AppResourceSecret, bool) {
	var e AppResourceSecret
	if m.Secret.IsNull() || m.Secret.IsUnknown() {
		return e, false
	}
	var v AppResourceSecret
	d := m.Secret.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSecret sets the value of the Secret field in AppResource.
func (m *AppResource) SetSecret(ctx context.Context, v AppResourceSecret) {
	vs := v.ToObjectValue(ctx)
	m.Secret = vs
}

// GetServingEndpoint returns the value of the ServingEndpoint field in AppResource as
// a AppResourceServingEndpoint value.
// If the field is unknown or null, the boolean return value is false.
func (m *AppResource) GetServingEndpoint(ctx context.Context) (AppResourceServingEndpoint, bool) {
	var e AppResourceServingEndpoint
	if m.ServingEndpoint.IsNull() || m.ServingEndpoint.IsUnknown() {
		return e, false
	}
	var v AppResourceServingEndpoint
	d := m.ServingEndpoint.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServingEndpoint sets the value of the ServingEndpoint field in AppResource.
func (m *AppResource) SetServingEndpoint(ctx context.Context, v AppResourceServingEndpoint) {
	vs := v.ToObjectValue(ctx)
	m.ServingEndpoint = vs
}

// GetSqlWarehouse returns the value of the SqlWarehouse field in AppResource as
// a AppResourceSqlWarehouse value.
// If the field is unknown or null, the boolean return value is false.
func (m *AppResource) GetSqlWarehouse(ctx context.Context) (AppResourceSqlWarehouse, bool) {
	var e AppResourceSqlWarehouse
	if m.SqlWarehouse.IsNull() || m.SqlWarehouse.IsUnknown() {
		return e, false
	}
	var v AppResourceSqlWarehouse
	d := m.SqlWarehouse.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSqlWarehouse sets the value of the SqlWarehouse field in AppResource.
func (m *AppResource) SetSqlWarehouse(ctx context.Context, v AppResourceSqlWarehouse) {
	vs := v.ToObjectValue(ctx)
	m.SqlWarehouse = vs
}

// GetUcSecurable returns the value of the UcSecurable field in AppResource as
// a AppResourceUcSecurable value.
// If the field is unknown or null, the boolean return value is false.
func (m *AppResource) GetUcSecurable(ctx context.Context) (AppResourceUcSecurable, bool) {
	var e AppResourceUcSecurable
	if m.UcSecurable.IsNull() || m.UcSecurable.IsUnknown() {
		return e, false
	}
	var v AppResourceUcSecurable
	d := m.UcSecurable.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUcSecurable sets the value of the UcSecurable field in AppResource.
func (m *AppResource) SetUcSecurable(ctx context.Context, v AppResourceUcSecurable) {
	vs := v.ToObjectValue(ctx)
	m.UcSecurable = vs
}

type AppResourceDatabase struct {
	DatabaseName types.String `tfsdk:"database_name"`

	InstanceName types.String `tfsdk:"instance_name"`

	Permission types.String `tfsdk:"permission"`
}

func (to *AppResourceDatabase) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AppResourceDatabase) {
}

func (to *AppResourceDatabase) SyncFieldsDuringRead(ctx context.Context, from AppResourceDatabase) {
}

func (m AppResourceDatabase) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["database_name"] = attrs["database_name"].SetRequired()
	attrs["instance_name"] = attrs["instance_name"].SetRequired()
	attrs["permission"] = attrs["permission"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppResourceDatabase.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AppResourceDatabase) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppResourceDatabase
// only implements ToObjectValue() and Type().
func (m AppResourceDatabase) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_name": m.DatabaseName,
			"instance_name": m.InstanceName,
			"permission":    m.Permission,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AppResourceDatabase) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"database_name": types.StringType,
			"instance_name": types.StringType,
			"permission":    types.StringType,
		},
	}
}

type AppResourceExperiment struct {
	ExperimentId types.String `tfsdk:"experiment_id"`

	Permission types.String `tfsdk:"permission"`
}

func (to *AppResourceExperiment) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AppResourceExperiment) {
}

func (to *AppResourceExperiment) SyncFieldsDuringRead(ctx context.Context, from AppResourceExperiment) {
}

func (m AppResourceExperiment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["experiment_id"] = attrs["experiment_id"].SetRequired()
	attrs["permission"] = attrs["permission"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppResourceExperiment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AppResourceExperiment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppResourceExperiment
// only implements ToObjectValue() and Type().
func (m AppResourceExperiment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": m.ExperimentId,
			"permission":    m.Permission,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AppResourceExperiment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
			"permission":    types.StringType,
		},
	}
}

type AppResourceGenieSpace struct {
	Name types.String `tfsdk:"name"`

	Permission types.String `tfsdk:"permission"`

	SpaceId types.String `tfsdk:"space_id"`
}

func (to *AppResourceGenieSpace) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AppResourceGenieSpace) {
}

func (to *AppResourceGenieSpace) SyncFieldsDuringRead(ctx context.Context, from AppResourceGenieSpace) {
}

func (m AppResourceGenieSpace) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["permission"] = attrs["permission"].SetRequired()
	attrs["space_id"] = attrs["space_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppResourceGenieSpace.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AppResourceGenieSpace) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppResourceGenieSpace
// only implements ToObjectValue() and Type().
func (m AppResourceGenieSpace) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":       m.Name,
			"permission": m.Permission,
			"space_id":   m.SpaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AppResourceGenieSpace) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":       types.StringType,
			"permission": types.StringType,
			"space_id":   types.StringType,
		},
	}
}

type AppResourceJob struct {
	// Id of the job to grant permission on.
	Id types.String `tfsdk:"id"`
	// Permissions to grant on the Job. Supported permissions are: "CAN_MANAGE",
	// "IS_OWNER", "CAN_MANAGE_RUN", "CAN_VIEW".
	Permission types.String `tfsdk:"permission"`
}

func (to *AppResourceJob) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AppResourceJob) {
}

func (to *AppResourceJob) SyncFieldsDuringRead(ctx context.Context, from AppResourceJob) {
}

func (m AppResourceJob) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()
	attrs["permission"] = attrs["permission"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppResourceJob.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AppResourceJob) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppResourceJob
// only implements ToObjectValue() and Type().
func (m AppResourceJob) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":         m.Id,
			"permission": m.Permission,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AppResourceJob) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":         types.StringType,
			"permission": types.StringType,
		},
	}
}

type AppResourceSecret struct {
	// Key of the secret to grant permission on.
	Key types.String `tfsdk:"key"`
	// Permission to grant on the secret scope. For secrets, only one permission
	// is allowed. Permission must be one of: "READ", "WRITE", "MANAGE".
	Permission types.String `tfsdk:"permission"`
	// Scope of the secret to grant permission on.
	Scope types.String `tfsdk:"scope"`
}

func (to *AppResourceSecret) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AppResourceSecret) {
}

func (to *AppResourceSecret) SyncFieldsDuringRead(ctx context.Context, from AppResourceSecret) {
}

func (m AppResourceSecret) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetRequired()
	attrs["permission"] = attrs["permission"].SetRequired()
	attrs["scope"] = attrs["scope"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppResourceSecret.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AppResourceSecret) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppResourceSecret
// only implements ToObjectValue() and Type().
func (m AppResourceSecret) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":        m.Key,
			"permission": m.Permission,
			"scope":      m.Scope,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AppResourceSecret) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":        types.StringType,
			"permission": types.StringType,
			"scope":      types.StringType,
		},
	}
}

type AppResourceServingEndpoint struct {
	// Name of the serving endpoint to grant permission on.
	Name types.String `tfsdk:"name"`
	// Permission to grant on the serving endpoint. Supported permissions are:
	// "CAN_MANAGE", "CAN_QUERY", "CAN_VIEW".
	Permission types.String `tfsdk:"permission"`
}

func (to *AppResourceServingEndpoint) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AppResourceServingEndpoint) {
}

func (to *AppResourceServingEndpoint) SyncFieldsDuringRead(ctx context.Context, from AppResourceServingEndpoint) {
}

func (m AppResourceServingEndpoint) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["permission"] = attrs["permission"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppResourceServingEndpoint.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AppResourceServingEndpoint) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppResourceServingEndpoint
// only implements ToObjectValue() and Type().
func (m AppResourceServingEndpoint) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":       m.Name,
			"permission": m.Permission,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AppResourceServingEndpoint) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":       types.StringType,
			"permission": types.StringType,
		},
	}
}

type AppResourceSqlWarehouse struct {
	// Id of the SQL warehouse to grant permission on.
	Id types.String `tfsdk:"id"`
	// Permission to grant on the SQL warehouse. Supported permissions are:
	// "CAN_MANAGE", "CAN_USE", "IS_OWNER".
	Permission types.String `tfsdk:"permission"`
}

func (to *AppResourceSqlWarehouse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AppResourceSqlWarehouse) {
}

func (to *AppResourceSqlWarehouse) SyncFieldsDuringRead(ctx context.Context, from AppResourceSqlWarehouse) {
}

func (m AppResourceSqlWarehouse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()
	attrs["permission"] = attrs["permission"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppResourceSqlWarehouse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AppResourceSqlWarehouse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppResourceSqlWarehouse
// only implements ToObjectValue() and Type().
func (m AppResourceSqlWarehouse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id":         m.Id,
			"permission": m.Permission,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AppResourceSqlWarehouse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":         types.StringType,
			"permission": types.StringType,
		},
	}
}

type AppResourceUcSecurable struct {
	Permission types.String `tfsdk:"permission"`

	SecurableFullName types.String `tfsdk:"securable_full_name"`

	SecurableType types.String `tfsdk:"securable_type"`
}

func (to *AppResourceUcSecurable) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AppResourceUcSecurable) {
}

func (to *AppResourceUcSecurable) SyncFieldsDuringRead(ctx context.Context, from AppResourceUcSecurable) {
}

func (m AppResourceUcSecurable) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission"] = attrs["permission"].SetRequired()
	attrs["securable_full_name"] = attrs["securable_full_name"].SetRequired()
	attrs["securable_type"] = attrs["securable_type"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppResourceUcSecurable.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AppResourceUcSecurable) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppResourceUcSecurable
// only implements ToObjectValue() and Type().
func (m AppResourceUcSecurable) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission":          m.Permission,
			"securable_full_name": m.SecurableFullName,
			"securable_type":      m.SecurableType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AppResourceUcSecurable) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission":          types.StringType,
			"securable_full_name": types.StringType,
			"securable_type":      types.StringType,
		},
	}
}

type AppUpdate struct {
	BudgetPolicyId types.String `tfsdk:"budget_policy_id"`

	ComputeSize types.String `tfsdk:"compute_size"`

	Description types.String `tfsdk:"description"`

	GitRepository types.Object `tfsdk:"git_repository"`

	Resources types.List `tfsdk:"resources"`

	Status types.Object `tfsdk:"status"`

	UsagePolicyId types.String `tfsdk:"usage_policy_id"`

	UserApiScopes types.List `tfsdk:"user_api_scopes"`
}

func (to *AppUpdate) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AppUpdate) {
	if !from.GitRepository.IsNull() && !from.GitRepository.IsUnknown() {
		if toGitRepository, ok := to.GetGitRepository(ctx); ok {
			if fromGitRepository, ok := from.GetGitRepository(ctx); ok {
				// Recursively sync the fields of GitRepository
				toGitRepository.SyncFieldsDuringCreateOrUpdate(ctx, fromGitRepository)
				to.SetGitRepository(ctx, toGitRepository)
			}
		}
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
}

func (to *AppUpdate) SyncFieldsDuringRead(ctx context.Context, from AppUpdate) {
	if !from.GitRepository.IsNull() && !from.GitRepository.IsUnknown() {
		if toGitRepository, ok := to.GetGitRepository(ctx); ok {
			if fromGitRepository, ok := from.GetGitRepository(ctx); ok {
				toGitRepository.SyncFieldsDuringRead(ctx, fromGitRepository)
				to.SetGitRepository(ctx, toGitRepository)
			}
		}
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
}

func (m AppUpdate) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["budget_policy_id"] = attrs["budget_policy_id"].SetOptional()
	attrs["compute_size"] = attrs["compute_size"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["git_repository"] = attrs["git_repository"].SetOptional()
	attrs["resources"] = attrs["resources"].SetOptional()
	attrs["status"] = attrs["status"].SetComputed()
	attrs["usage_policy_id"] = attrs["usage_policy_id"].SetOptional()
	attrs["user_api_scopes"] = attrs["user_api_scopes"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppUpdate.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AppUpdate) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"git_repository":  reflect.TypeOf(GitRepository{}),
		"resources":       reflect.TypeOf(AppResource{}),
		"status":          reflect.TypeOf(AppUpdateUpdateStatus{}),
		"user_api_scopes": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppUpdate
// only implements ToObjectValue() and Type().
func (m AppUpdate) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budget_policy_id": m.BudgetPolicyId,
			"compute_size":     m.ComputeSize,
			"description":      m.Description,
			"git_repository":   m.GitRepository,
			"resources":        m.Resources,
			"status":           m.Status,
			"usage_policy_id":  m.UsagePolicyId,
			"user_api_scopes":  m.UserApiScopes,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AppUpdate) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget_policy_id": types.StringType,
			"compute_size":     types.StringType,
			"description":      types.StringType,
			"git_repository":   GitRepository{}.Type(ctx),
			"resources": basetypes.ListType{
				ElemType: AppResource{}.Type(ctx),
			},
			"status":          AppUpdateUpdateStatus{}.Type(ctx),
			"usage_policy_id": types.StringType,
			"user_api_scopes": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetGitRepository returns the value of the GitRepository field in AppUpdate as
// a GitRepository value.
// If the field is unknown or null, the boolean return value is false.
func (m *AppUpdate) GetGitRepository(ctx context.Context) (GitRepository, bool) {
	var e GitRepository
	if m.GitRepository.IsNull() || m.GitRepository.IsUnknown() {
		return e, false
	}
	var v GitRepository
	d := m.GitRepository.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGitRepository sets the value of the GitRepository field in AppUpdate.
func (m *AppUpdate) SetGitRepository(ctx context.Context, v GitRepository) {
	vs := v.ToObjectValue(ctx)
	m.GitRepository = vs
}

// GetResources returns the value of the Resources field in AppUpdate as
// a slice of AppResource values.
// If the field is unknown or null, the boolean return value is false.
func (m *AppUpdate) GetResources(ctx context.Context) ([]AppResource, bool) {
	if m.Resources.IsNull() || m.Resources.IsUnknown() {
		return nil, false
	}
	var v []AppResource
	d := m.Resources.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResources sets the value of the Resources field in AppUpdate.
func (m *AppUpdate) SetResources(ctx context.Context, v []AppResource) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["resources"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Resources = types.ListValueMust(t, vs)
}

// GetStatus returns the value of the Status field in AppUpdate as
// a AppUpdateUpdateStatus value.
// If the field is unknown or null, the boolean return value is false.
func (m *AppUpdate) GetStatus(ctx context.Context) (AppUpdateUpdateStatus, bool) {
	var e AppUpdateUpdateStatus
	if m.Status.IsNull() || m.Status.IsUnknown() {
		return e, false
	}
	var v AppUpdateUpdateStatus
	d := m.Status.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatus sets the value of the Status field in AppUpdate.
func (m *AppUpdate) SetStatus(ctx context.Context, v AppUpdateUpdateStatus) {
	vs := v.ToObjectValue(ctx)
	m.Status = vs
}

// GetUserApiScopes returns the value of the UserApiScopes field in AppUpdate as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *AppUpdate) GetUserApiScopes(ctx context.Context) ([]types.String, bool) {
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

// SetUserApiScopes sets the value of the UserApiScopes field in AppUpdate.
func (m *AppUpdate) SetUserApiScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["user_api_scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.UserApiScopes = types.ListValueMust(t, vs)
}

type AppUpdateUpdateStatus struct {
	Message types.String `tfsdk:"message"`

	State types.String `tfsdk:"state"`
}

func (to *AppUpdateUpdateStatus) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AppUpdateUpdateStatus) {
}

func (to *AppUpdateUpdateStatus) SyncFieldsDuringRead(ctx context.Context, from AppUpdateUpdateStatus) {
}

func (m AppUpdateUpdateStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["message"] = attrs["message"].SetComputed()
	attrs["state"] = attrs["state"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AppUpdateUpdateStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AppUpdateUpdateStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AppUpdateUpdateStatus
// only implements ToObjectValue() and Type().
func (m AppUpdateUpdateStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"message": m.Message,
			"state":   m.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AppUpdateUpdateStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"message": types.StringType,
			"state":   types.StringType,
		},
	}
}

type ApplicationStatus struct {
	// Application status message
	Message types.String `tfsdk:"message"`
	// State of the application.
	State types.String `tfsdk:"state"`
}

func (to *ApplicationStatus) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ApplicationStatus) {
}

func (to *ApplicationStatus) SyncFieldsDuringRead(ctx context.Context, from ApplicationStatus) {
}

func (m ApplicationStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["message"] = attrs["message"].SetComputed()
	attrs["state"] = attrs["state"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ApplicationStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ApplicationStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ApplicationStatus
// only implements ToObjectValue() and Type().
func (m ApplicationStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"message": m.Message,
			"state":   m.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ApplicationStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"message": types.StringType,
			"state":   types.StringType,
		},
	}
}

type AsyncUpdateAppRequest struct {
	App types.Object `tfsdk:"app"`

	AppName types.String `tfsdk:"-"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	UpdateMask types.String `tfsdk:"update_mask"`
}

func (to *AsyncUpdateAppRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AsyncUpdateAppRequest) {
	if !from.App.IsNull() && !from.App.IsUnknown() {
		if toApp, ok := to.GetApp(ctx); ok {
			if fromApp, ok := from.GetApp(ctx); ok {
				// Recursively sync the fields of App
				toApp.SyncFieldsDuringCreateOrUpdate(ctx, fromApp)
				to.SetApp(ctx, toApp)
			}
		}
	}
}

func (to *AsyncUpdateAppRequest) SyncFieldsDuringRead(ctx context.Context, from AsyncUpdateAppRequest) {
	if !from.App.IsNull() && !from.App.IsUnknown() {
		if toApp, ok := to.GetApp(ctx); ok {
			if fromApp, ok := from.GetApp(ctx); ok {
				toApp.SyncFieldsDuringRead(ctx, fromApp)
				to.SetApp(ctx, toApp)
			}
		}
	}
}

func (m AsyncUpdateAppRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["app"] = attrs["app"].SetOptional()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()
	attrs["app_name"] = attrs["app_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AsyncUpdateAppRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AsyncUpdateAppRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"app": reflect.TypeOf(App{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AsyncUpdateAppRequest
// only implements ToObjectValue() and Type().
func (m AsyncUpdateAppRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app":         m.App,
			"app_name":    m.AppName,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AsyncUpdateAppRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app":         App{}.Type(ctx),
			"app_name":    types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetApp returns the value of the App field in AsyncUpdateAppRequest as
// a App value.
// If the field is unknown or null, the boolean return value is false.
func (m *AsyncUpdateAppRequest) GetApp(ctx context.Context) (App, bool) {
	var e App
	if m.App.IsNull() || m.App.IsUnknown() {
		return e, false
	}
	var v App
	d := m.App.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetApp sets the value of the App field in AsyncUpdateAppRequest.
func (m *AsyncUpdateAppRequest) SetApp(ctx context.Context, v App) {
	vs := v.ToObjectValue(ctx)
	m.App = vs
}

type ComputeStatus struct {
	// The number of compute instances currently serving requests for this
	// application. An instance is considered active if it is reachable and
	// ready to handle requests.
	ActiveInstances types.Int64 `tfsdk:"active_instances"`
	// Compute status message
	Message types.String `tfsdk:"message"`
	// State of the app compute.
	State types.String `tfsdk:"state"`
}

func (to *ComputeStatus) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ComputeStatus) {
}

func (to *ComputeStatus) SyncFieldsDuringRead(ctx context.Context, from ComputeStatus) {
}

func (m ComputeStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["active_instances"] = attrs["active_instances"].SetComputed()
	attrs["message"] = attrs["message"].SetComputed()
	attrs["state"] = attrs["state"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ComputeStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ComputeStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ComputeStatus
// only implements ToObjectValue() and Type().
func (m ComputeStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"active_instances": m.ActiveInstances,
			"message":          m.Message,
			"state":            m.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ComputeStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"active_instances": types.Int64Type,
			"message":          types.StringType,
			"state":            types.StringType,
		},
	}
}

type CreateAppDeploymentRequest struct {
	// The app deployment configuration.
	AppDeployment types.Object `tfsdk:"app_deployment"`
	// The name of the app.
	AppName types.String `tfsdk:"-"`
}

func (to *CreateAppDeploymentRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateAppDeploymentRequest) {
	if !from.AppDeployment.IsNull() && !from.AppDeployment.IsUnknown() {
		if toAppDeployment, ok := to.GetAppDeployment(ctx); ok {
			if fromAppDeployment, ok := from.GetAppDeployment(ctx); ok {
				// Recursively sync the fields of AppDeployment
				toAppDeployment.SyncFieldsDuringCreateOrUpdate(ctx, fromAppDeployment)
				to.SetAppDeployment(ctx, toAppDeployment)
			}
		}
	}
}

func (to *CreateAppDeploymentRequest) SyncFieldsDuringRead(ctx context.Context, from CreateAppDeploymentRequest) {
	if !from.AppDeployment.IsNull() && !from.AppDeployment.IsUnknown() {
		if toAppDeployment, ok := to.GetAppDeployment(ctx); ok {
			if fromAppDeployment, ok := from.GetAppDeployment(ctx); ok {
				toAppDeployment.SyncFieldsDuringRead(ctx, fromAppDeployment)
				to.SetAppDeployment(ctx, toAppDeployment)
			}
		}
	}
}

func (m CreateAppDeploymentRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["app_deployment"] = attrs["app_deployment"].SetRequired()
	attrs["app_name"] = attrs["app_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateAppDeploymentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateAppDeploymentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"app_deployment": reflect.TypeOf(AppDeployment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateAppDeploymentRequest
// only implements ToObjectValue() and Type().
func (m CreateAppDeploymentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app_deployment": m.AppDeployment,
			"app_name":       m.AppName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateAppDeploymentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app_deployment": AppDeployment{}.Type(ctx),
			"app_name":       types.StringType,
		},
	}
}

// GetAppDeployment returns the value of the AppDeployment field in CreateAppDeploymentRequest as
// a AppDeployment value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateAppDeploymentRequest) GetAppDeployment(ctx context.Context) (AppDeployment, bool) {
	var e AppDeployment
	if m.AppDeployment.IsNull() || m.AppDeployment.IsUnknown() {
		return e, false
	}
	var v AppDeployment
	d := m.AppDeployment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAppDeployment sets the value of the AppDeployment field in CreateAppDeploymentRequest.
func (m *CreateAppDeploymentRequest) SetAppDeployment(ctx context.Context, v AppDeployment) {
	vs := v.ToObjectValue(ctx)
	m.AppDeployment = vs
}

type CreateAppRequest struct {
	App types.Object `tfsdk:"app"`
	// If true, the app will not be started after creation.
	NoCompute types.Bool `tfsdk:"-"`
}

func (to *CreateAppRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateAppRequest) {
	if !from.App.IsNull() && !from.App.IsUnknown() {
		if toApp, ok := to.GetApp(ctx); ok {
			if fromApp, ok := from.GetApp(ctx); ok {
				// Recursively sync the fields of App
				toApp.SyncFieldsDuringCreateOrUpdate(ctx, fromApp)
				to.SetApp(ctx, toApp)
			}
		}
	}
}

func (to *CreateAppRequest) SyncFieldsDuringRead(ctx context.Context, from CreateAppRequest) {
	if !from.App.IsNull() && !from.App.IsUnknown() {
		if toApp, ok := to.GetApp(ctx); ok {
			if fromApp, ok := from.GetApp(ctx); ok {
				toApp.SyncFieldsDuringRead(ctx, fromApp)
				to.SetApp(ctx, toApp)
			}
		}
	}
}

func (m CreateAppRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["app"] = attrs["app"].SetRequired()
	attrs["no_compute"] = attrs["no_compute"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateAppRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateAppRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"app": reflect.TypeOf(App{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateAppRequest
// only implements ToObjectValue() and Type().
func (m CreateAppRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app":        m.App,
			"no_compute": m.NoCompute,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateAppRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app":        App{}.Type(ctx),
			"no_compute": types.BoolType,
		},
	}
}

// GetApp returns the value of the App field in CreateAppRequest as
// a App value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateAppRequest) GetApp(ctx context.Context) (App, bool) {
	var e App
	if m.App.IsNull() || m.App.IsUnknown() {
		return e, false
	}
	var v App
	d := m.App.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetApp sets the value of the App field in CreateAppRequest.
func (m *CreateAppRequest) SetApp(ctx context.Context, v App) {
	vs := v.ToObjectValue(ctx)
	m.App = vs
}

type CreateCustomTemplateRequest struct {
	Template types.Object `tfsdk:"template"`
}

func (to *CreateCustomTemplateRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCustomTemplateRequest) {
	if !from.Template.IsNull() && !from.Template.IsUnknown() {
		if toTemplate, ok := to.GetTemplate(ctx); ok {
			if fromTemplate, ok := from.GetTemplate(ctx); ok {
				// Recursively sync the fields of Template
				toTemplate.SyncFieldsDuringCreateOrUpdate(ctx, fromTemplate)
				to.SetTemplate(ctx, toTemplate)
			}
		}
	}
}

func (to *CreateCustomTemplateRequest) SyncFieldsDuringRead(ctx context.Context, from CreateCustomTemplateRequest) {
	if !from.Template.IsNull() && !from.Template.IsUnknown() {
		if toTemplate, ok := to.GetTemplate(ctx); ok {
			if fromTemplate, ok := from.GetTemplate(ctx); ok {
				toTemplate.SyncFieldsDuringRead(ctx, fromTemplate)
				to.SetTemplate(ctx, toTemplate)
			}
		}
	}
}

func (m CreateCustomTemplateRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["template"] = attrs["template"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCustomTemplateRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateCustomTemplateRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"template": reflect.TypeOf(CustomTemplate{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCustomTemplateRequest
// only implements ToObjectValue() and Type().
func (m CreateCustomTemplateRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"template": m.Template,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateCustomTemplateRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"template": CustomTemplate{}.Type(ctx),
		},
	}
}

// GetTemplate returns the value of the Template field in CreateCustomTemplateRequest as
// a CustomTemplate value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateCustomTemplateRequest) GetTemplate(ctx context.Context) (CustomTemplate, bool) {
	var e CustomTemplate
	if m.Template.IsNull() || m.Template.IsUnknown() {
		return e, false
	}
	var v CustomTemplate
	d := m.Template.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTemplate sets the value of the Template field in CreateCustomTemplateRequest.
func (m *CreateCustomTemplateRequest) SetTemplate(ctx context.Context, v CustomTemplate) {
	vs := v.ToObjectValue(ctx)
	m.Template = vs
}

type CreateSpaceRequest struct {
	Space types.Object `tfsdk:"space"`
}

func (to *CreateSpaceRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateSpaceRequest) {
	if !from.Space.IsNull() && !from.Space.IsUnknown() {
		if toSpace, ok := to.GetSpace(ctx); ok {
			if fromSpace, ok := from.GetSpace(ctx); ok {
				// Recursively sync the fields of Space
				toSpace.SyncFieldsDuringCreateOrUpdate(ctx, fromSpace)
				to.SetSpace(ctx, toSpace)
			}
		}
	}
}

func (to *CreateSpaceRequest) SyncFieldsDuringRead(ctx context.Context, from CreateSpaceRequest) {
	if !from.Space.IsNull() && !from.Space.IsUnknown() {
		if toSpace, ok := to.GetSpace(ctx); ok {
			if fromSpace, ok := from.GetSpace(ctx); ok {
				toSpace.SyncFieldsDuringRead(ctx, fromSpace)
				to.SetSpace(ctx, toSpace)
			}
		}
	}
}

func (m CreateSpaceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["space"] = attrs["space"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateSpaceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateSpaceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"space": reflect.TypeOf(Space{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateSpaceRequest
// only implements ToObjectValue() and Type().
func (m CreateSpaceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"space": m.Space,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateSpaceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"space": Space{}.Type(ctx),
		},
	}
}

// GetSpace returns the value of the Space field in CreateSpaceRequest as
// a Space value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateSpaceRequest) GetSpace(ctx context.Context) (Space, bool) {
	var e Space
	if m.Space.IsNull() || m.Space.IsUnknown() {
		return e, false
	}
	var v Space
	d := m.Space.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSpace sets the value of the Space field in CreateSpaceRequest.
func (m *CreateSpaceRequest) SetSpace(ctx context.Context, v Space) {
	vs := v.ToObjectValue(ctx)
	m.Space = vs
}

type CustomTemplate struct {
	Creator types.String `tfsdk:"creator"`
	// The description of the template.
	Description types.String `tfsdk:"description"`
	// The Git provider of the template.
	GitProvider types.String `tfsdk:"git_provider"`
	// The Git repository URL that the template resides in.
	GitRepo types.String `tfsdk:"git_repo"`
	// The manifest of the template. It defines fields and default values when
	// installing the template.
	Manifest types.Object `tfsdk:"manifest"`
	// The name of the template. It must contain only alphanumeric characters,
	// hyphens, underscores, and whitespaces. It must be unique within the
	// workspace.
	Name types.String `tfsdk:"name"`
	// The path to the template within the Git repository.
	Path types.String `tfsdk:"path"`
}

func (to *CustomTemplate) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CustomTemplate) {
	if !from.Manifest.IsNull() && !from.Manifest.IsUnknown() {
		if toManifest, ok := to.GetManifest(ctx); ok {
			if fromManifest, ok := from.GetManifest(ctx); ok {
				// Recursively sync the fields of Manifest
				toManifest.SyncFieldsDuringCreateOrUpdate(ctx, fromManifest)
				to.SetManifest(ctx, toManifest)
			}
		}
	}
}

func (to *CustomTemplate) SyncFieldsDuringRead(ctx context.Context, from CustomTemplate) {
	if !from.Manifest.IsNull() && !from.Manifest.IsUnknown() {
		if toManifest, ok := to.GetManifest(ctx); ok {
			if fromManifest, ok := from.GetManifest(ctx); ok {
				toManifest.SyncFieldsDuringRead(ctx, fromManifest)
				to.SetManifest(ctx, toManifest)
			}
		}
	}
}

func (m CustomTemplate) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["creator"] = attrs["creator"].SetComputed()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["git_provider"] = attrs["git_provider"].SetRequired()
	attrs["git_repo"] = attrs["git_repo"].SetRequired()
	attrs["manifest"] = attrs["manifest"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["path"] = attrs["path"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CustomTemplate.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CustomTemplate) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"manifest": reflect.TypeOf(AppManifest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CustomTemplate
// only implements ToObjectValue() and Type().
func (m CustomTemplate) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creator":      m.Creator,
			"description":  m.Description,
			"git_provider": m.GitProvider,
			"git_repo":     m.GitRepo,
			"manifest":     m.Manifest,
			"name":         m.Name,
			"path":         m.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CustomTemplate) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creator":      types.StringType,
			"description":  types.StringType,
			"git_provider": types.StringType,
			"git_repo":     types.StringType,
			"manifest":     AppManifest{}.Type(ctx),
			"name":         types.StringType,
			"path":         types.StringType,
		},
	}
}

// GetManifest returns the value of the Manifest field in CustomTemplate as
// a AppManifest value.
// If the field is unknown or null, the boolean return value is false.
func (m *CustomTemplate) GetManifest(ctx context.Context) (AppManifest, bool) {
	var e AppManifest
	if m.Manifest.IsNull() || m.Manifest.IsUnknown() {
		return e, false
	}
	var v AppManifest
	d := m.Manifest.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetManifest sets the value of the Manifest field in CustomTemplate.
func (m *CustomTemplate) SetManifest(ctx context.Context, v AppManifest) {
	vs := v.ToObjectValue(ctx)
	m.Manifest = vs
}

// Databricks Error that is returned by all Databricks APIs.
type DatabricksServiceExceptionWithDetailsProto struct {
	Details types.List `tfsdk:"details"`

	ErrorCode types.String `tfsdk:"error_code"`

	Message types.String `tfsdk:"message"`

	StackTrace types.String `tfsdk:"stack_trace"`
}

func (to *DatabricksServiceExceptionWithDetailsProto) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabricksServiceExceptionWithDetailsProto) {
	if !from.Details.IsNull() && !from.Details.IsUnknown() && to.Details.IsNull() && len(from.Details.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Details, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Details = from.Details
	}
}

func (to *DatabricksServiceExceptionWithDetailsProto) SyncFieldsDuringRead(ctx context.Context, from DatabricksServiceExceptionWithDetailsProto) {
	if !from.Details.IsNull() && !from.Details.IsUnknown() && to.Details.IsNull() && len(from.Details.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Details, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Details = from.Details
	}
}

func (m DatabricksServiceExceptionWithDetailsProto) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["details"] = attrs["details"].SetOptional()
	attrs["error_code"] = attrs["error_code"].SetOptional()
	attrs["message"] = attrs["message"].SetOptional()
	attrs["stack_trace"] = attrs["stack_trace"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatabricksServiceExceptionWithDetailsProto.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DatabricksServiceExceptionWithDetailsProto) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"details": reflect.TypeOf(types.Object{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabricksServiceExceptionWithDetailsProto
// only implements ToObjectValue() and Type().
func (m DatabricksServiceExceptionWithDetailsProto) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"details":     m.Details,
			"error_code":  m.ErrorCode,
			"message":     m.Message,
			"stack_trace": m.StackTrace,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DatabricksServiceExceptionWithDetailsProto) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"details": basetypes.ListType{
				ElemType: types.ObjectType{},
			},
			"error_code":  types.StringType,
			"message":     types.StringType,
			"stack_trace": types.StringType,
		},
	}
}

// GetDetails returns the value of the Details field in DatabricksServiceExceptionWithDetailsProto as
// a slice of types.Object values.
// If the field is unknown or null, the boolean return value is false.
func (m *DatabricksServiceExceptionWithDetailsProto) GetDetails(ctx context.Context) ([]types.Object, bool) {
	if m.Details.IsNull() || m.Details.IsUnknown() {
		return nil, false
	}
	var v []types.Object
	d := m.Details.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDetails sets the value of the Details field in DatabricksServiceExceptionWithDetailsProto.
func (m *DatabricksServiceExceptionWithDetailsProto) SetDetails(ctx context.Context, v []types.Object) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["details"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Details = types.ListValueMust(t, vs)
}

type DeleteAppRequest struct {
	// The name of the app.
	Name types.String `tfsdk:"-"`
}

func (to *DeleteAppRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteAppRequest) {
}

func (to *DeleteAppRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteAppRequest) {
}

func (m DeleteAppRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAppRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteAppRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAppRequest
// only implements ToObjectValue() and Type().
func (m DeleteAppRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteAppRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteCustomTemplateRequest struct {
	// The name of the custom template.
	Name types.String `tfsdk:"-"`
}

func (to *DeleteCustomTemplateRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteCustomTemplateRequest) {
}

func (to *DeleteCustomTemplateRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteCustomTemplateRequest) {
}

func (m DeleteCustomTemplateRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCustomTemplateRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteCustomTemplateRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCustomTemplateRequest
// only implements ToObjectValue() and Type().
func (m DeleteCustomTemplateRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteCustomTemplateRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteSpaceRequest struct {
	// The name of the app space.
	Name types.String `tfsdk:"-"`
}

func (to *DeleteSpaceRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteSpaceRequest) {
}

func (to *DeleteSpaceRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteSpaceRequest) {
}

func (m DeleteSpaceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteSpaceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteSpaceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSpaceRequest
// only implements ToObjectValue() and Type().
func (m DeleteSpaceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteSpaceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type EnvVar struct {
	// The name of the environment variable.
	Name types.String `tfsdk:"name"`
	// The value for the environment variable.
	Value types.String `tfsdk:"value"`
	// The name of an external Databricks resource that contains the value, such
	// as a secret or a database table.
	ValueFrom types.String `tfsdk:"value_from"`
}

func (to *EnvVar) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EnvVar) {
}

func (to *EnvVar) SyncFieldsDuringRead(ctx context.Context, from EnvVar) {
}

func (m EnvVar) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetOptional()
	attrs["value"] = attrs["value"].SetOptional()
	attrs["value_from"] = attrs["value_from"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EnvVar.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EnvVar) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnvVar
// only implements ToObjectValue() and Type().
func (m EnvVar) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":       m.Name,
			"value":      m.Value,
			"value_from": m.ValueFrom,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EnvVar) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":       types.StringType,
			"value":      types.StringType,
			"value_from": types.StringType,
		},
	}
}

type GetAppDeploymentRequest struct {
	// The name of the app.
	AppName types.String `tfsdk:"-"`
	// The unique id of the deployment.
	DeploymentId types.String `tfsdk:"-"`
}

func (to *GetAppDeploymentRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetAppDeploymentRequest) {
}

func (to *GetAppDeploymentRequest) SyncFieldsDuringRead(ctx context.Context, from GetAppDeploymentRequest) {
}

func (m GetAppDeploymentRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["app_name"] = attrs["app_name"].SetRequired()
	attrs["deployment_id"] = attrs["deployment_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAppDeploymentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetAppDeploymentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAppDeploymentRequest
// only implements ToObjectValue() and Type().
func (m GetAppDeploymentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app_name":      m.AppName,
			"deployment_id": m.DeploymentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetAppDeploymentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app_name":      types.StringType,
			"deployment_id": types.StringType,
		},
	}
}

type GetAppPermissionLevelsRequest struct {
	// The app for which to get or manage permissions.
	AppName types.String `tfsdk:"-"`
}

func (to *GetAppPermissionLevelsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetAppPermissionLevelsRequest) {
}

func (to *GetAppPermissionLevelsRequest) SyncFieldsDuringRead(ctx context.Context, from GetAppPermissionLevelsRequest) {
}

func (m GetAppPermissionLevelsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["app_name"] = attrs["app_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAppPermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetAppPermissionLevelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAppPermissionLevelsRequest
// only implements ToObjectValue() and Type().
func (m GetAppPermissionLevelsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app_name": m.AppName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetAppPermissionLevelsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app_name": types.StringType,
		},
	}
}

type GetAppPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (to *GetAppPermissionLevelsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetAppPermissionLevelsResponse) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (to *GetAppPermissionLevelsResponse) SyncFieldsDuringRead(ctx context.Context, from GetAppPermissionLevelsResponse) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (m GetAppPermissionLevelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission_levels"] = attrs["permission_levels"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAppPermissionLevelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetAppPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(AppPermissionsDescription{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAppPermissionLevelsResponse
// only implements ToObjectValue() and Type().
func (m GetAppPermissionLevelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": m.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetAppPermissionLevelsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: AppPermissionsDescription{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetAppPermissionLevelsResponse as
// a slice of AppPermissionsDescription values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetAppPermissionLevelsResponse) GetPermissionLevels(ctx context.Context) ([]AppPermissionsDescription, bool) {
	if m.PermissionLevels.IsNull() || m.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []AppPermissionsDescription
	d := m.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetAppPermissionLevelsResponse.
func (m *GetAppPermissionLevelsResponse) SetPermissionLevels(ctx context.Context, v []AppPermissionsDescription) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PermissionLevels = types.ListValueMust(t, vs)
}

type GetAppPermissionsRequest struct {
	// The app for which to get or manage permissions.
	AppName types.String `tfsdk:"-"`
}

func (to *GetAppPermissionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetAppPermissionsRequest) {
}

func (to *GetAppPermissionsRequest) SyncFieldsDuringRead(ctx context.Context, from GetAppPermissionsRequest) {
}

func (m GetAppPermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["app_name"] = attrs["app_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAppPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetAppPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAppPermissionsRequest
// only implements ToObjectValue() and Type().
func (m GetAppPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app_name": m.AppName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetAppPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app_name": types.StringType,
		},
	}
}

type GetAppRequest struct {
	// The name of the app.
	Name types.String `tfsdk:"-"`
}

func (to *GetAppRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetAppRequest) {
}

func (to *GetAppRequest) SyncFieldsDuringRead(ctx context.Context, from GetAppRequest) {
}

func (m GetAppRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAppRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetAppRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAppRequest
// only implements ToObjectValue() and Type().
func (m GetAppRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetAppRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetAppUpdateRequest struct {
	// The name of the app.
	AppName types.String `tfsdk:"-"`
}

func (to *GetAppUpdateRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetAppUpdateRequest) {
}

func (to *GetAppUpdateRequest) SyncFieldsDuringRead(ctx context.Context, from GetAppUpdateRequest) {
}

func (m GetAppUpdateRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["app_name"] = attrs["app_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAppUpdateRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetAppUpdateRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAppUpdateRequest
// only implements ToObjectValue() and Type().
func (m GetAppUpdateRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app_name": m.AppName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetAppUpdateRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app_name": types.StringType,
		},
	}
}

type GetCustomTemplateRequest struct {
	// The name of the custom template.
	Name types.String `tfsdk:"-"`
}

func (to *GetCustomTemplateRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetCustomTemplateRequest) {
}

func (to *GetCustomTemplateRequest) SyncFieldsDuringRead(ctx context.Context, from GetCustomTemplateRequest) {
}

func (m GetCustomTemplateRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetCustomTemplateRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetCustomTemplateRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCustomTemplateRequest
// only implements ToObjectValue() and Type().
func (m GetCustomTemplateRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetCustomTemplateRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetOperationRequest struct {
	// The name of the operation resource.
	Name types.String `tfsdk:"-"`
}

func (to *GetOperationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetOperationRequest) {
}

func (to *GetOperationRequest) SyncFieldsDuringRead(ctx context.Context, from GetOperationRequest) {
}

func (m GetOperationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetOperationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetOperationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetOperationRequest
// only implements ToObjectValue() and Type().
func (m GetOperationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetOperationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetSpaceRequest struct {
	// The name of the app space.
	Name types.String `tfsdk:"-"`
}

func (to *GetSpaceRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetSpaceRequest) {
}

func (to *GetSpaceRequest) SyncFieldsDuringRead(ctx context.Context, from GetSpaceRequest) {
}

func (m GetSpaceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetSpaceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetSpaceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSpaceRequest
// only implements ToObjectValue() and Type().
func (m GetSpaceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetSpaceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Git repository configuration specifying the location of the repository.
type GitRepository struct {
	// Git provider. Case insensitive. Supported values: gitHub,
	// gitHubEnterprise, bitbucketCloud, bitbucketServer, azureDevOpsServices,
	// gitLab, gitLabEnterpriseEdition, awsCodeCommit.
	Provider types.String `tfsdk:"provider"`
	// URL of the Git repository.
	Url types.String `tfsdk:"url"`
}

func (to *GitRepository) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GitRepository) {
}

func (to *GitRepository) SyncFieldsDuringRead(ctx context.Context, from GitRepository) {
}

func (m GitRepository) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["provider"] = attrs["provider"].SetRequired()
	attrs["url"] = attrs["url"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GitRepository.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GitRepository) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GitRepository
// only implements ToObjectValue() and Type().
func (m GitRepository) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"provider": m.Provider,
			"url":      m.Url,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GitRepository) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"provider": types.StringType,
			"url":      types.StringType,
		},
	}
}

// Complete git source specification including repository location and
// reference.
type GitSource struct {
	// Git branch to checkout.
	Branch types.String `tfsdk:"branch"`
	// Git commit SHA to checkout.
	Commit types.String `tfsdk:"commit"`
	// Git repository configuration. Populated from the app's git_repository
	// configuration.
	GitRepository types.Object `tfsdk:"git_repository"`
	// The resolved commit SHA that was actually used for the deployment. This
	// is populated by the system after resolving the reference (branch, tag, or
	// commit). If commit is specified directly, this will match commit. If a
	// branch or tag is specified, this contains the commit SHA that the branch
	// or tag pointed to at deployment time.
	ResolvedCommit types.String `tfsdk:"resolved_commit"`
	// Relative path to the app source code within the Git repository. If not
	// specified, the root of the repository is used.
	SourceCodePath types.String `tfsdk:"source_code_path"`
	// Git tag to checkout.
	Tag types.String `tfsdk:"tag"`
}

func (to *GitSource) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GitSource) {
	if !from.GitRepository.IsNull() && !from.GitRepository.IsUnknown() {
		if toGitRepository, ok := to.GetGitRepository(ctx); ok {
			if fromGitRepository, ok := from.GetGitRepository(ctx); ok {
				// Recursively sync the fields of GitRepository
				toGitRepository.SyncFieldsDuringCreateOrUpdate(ctx, fromGitRepository)
				to.SetGitRepository(ctx, toGitRepository)
			}
		}
	}
}

func (to *GitSource) SyncFieldsDuringRead(ctx context.Context, from GitSource) {
	if !from.GitRepository.IsNull() && !from.GitRepository.IsUnknown() {
		if toGitRepository, ok := to.GetGitRepository(ctx); ok {
			if fromGitRepository, ok := from.GetGitRepository(ctx); ok {
				toGitRepository.SyncFieldsDuringRead(ctx, fromGitRepository)
				to.SetGitRepository(ctx, toGitRepository)
			}
		}
	}
}

func (m GitSource) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["branch"] = attrs["branch"].SetOptional()
	attrs["commit"] = attrs["commit"].SetOptional()
	attrs["git_repository"] = attrs["git_repository"].SetComputed()
	attrs["resolved_commit"] = attrs["resolved_commit"].SetComputed()
	attrs["source_code_path"] = attrs["source_code_path"].SetOptional()
	attrs["tag"] = attrs["tag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GitSource.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GitSource) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"git_repository": reflect.TypeOf(GitRepository{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GitSource
// only implements ToObjectValue() and Type().
func (m GitSource) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch":           m.Branch,
			"commit":           m.Commit,
			"git_repository":   m.GitRepository,
			"resolved_commit":  m.ResolvedCommit,
			"source_code_path": m.SourceCodePath,
			"tag":              m.Tag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GitSource) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch":           types.StringType,
			"commit":           types.StringType,
			"git_repository":   GitRepository{}.Type(ctx),
			"resolved_commit":  types.StringType,
			"source_code_path": types.StringType,
			"tag":              types.StringType,
		},
	}
}

// GetGitRepository returns the value of the GitRepository field in GitSource as
// a GitRepository value.
// If the field is unknown or null, the boolean return value is false.
func (m *GitSource) GetGitRepository(ctx context.Context) (GitRepository, bool) {
	var e GitRepository
	if m.GitRepository.IsNull() || m.GitRepository.IsUnknown() {
		return e, false
	}
	var v GitRepository
	d := m.GitRepository.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGitRepository sets the value of the GitRepository field in GitSource.
func (m *GitSource) SetGitRepository(ctx context.Context, v GitRepository) {
	vs := v.ToObjectValue(ctx)
	m.GitRepository = vs
}

type ListAppDeploymentsRequest struct {
	// The name of the app.
	AppName types.String `tfsdk:"-"`
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page of apps. Requests first page if
	// absent.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListAppDeploymentsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListAppDeploymentsRequest) {
}

func (to *ListAppDeploymentsRequest) SyncFieldsDuringRead(ctx context.Context, from ListAppDeploymentsRequest) {
}

func (m ListAppDeploymentsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["app_name"] = attrs["app_name"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAppDeploymentsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListAppDeploymentsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAppDeploymentsRequest
// only implements ToObjectValue() and Type().
func (m ListAppDeploymentsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app_name":   m.AppName,
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListAppDeploymentsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app_name":   types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListAppDeploymentsResponse struct {
	// Deployment history of the app.
	AppDeployments types.List `tfsdk:"app_deployments"`
	// Pagination token to request the next page of apps.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListAppDeploymentsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListAppDeploymentsResponse) {
	if !from.AppDeployments.IsNull() && !from.AppDeployments.IsUnknown() && to.AppDeployments.IsNull() && len(from.AppDeployments.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AppDeployments, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AppDeployments = from.AppDeployments
	}
}

func (to *ListAppDeploymentsResponse) SyncFieldsDuringRead(ctx context.Context, from ListAppDeploymentsResponse) {
	if !from.AppDeployments.IsNull() && !from.AppDeployments.IsUnknown() && to.AppDeployments.IsNull() && len(from.AppDeployments.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AppDeployments, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AppDeployments = from.AppDeployments
	}
}

func (m ListAppDeploymentsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["app_deployments"] = attrs["app_deployments"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAppDeploymentsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListAppDeploymentsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"app_deployments": reflect.TypeOf(AppDeployment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAppDeploymentsResponse
// only implements ToObjectValue() and Type().
func (m ListAppDeploymentsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app_deployments": m.AppDeployments,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListAppDeploymentsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app_deployments": basetypes.ListType{
				ElemType: AppDeployment{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetAppDeployments returns the value of the AppDeployments field in ListAppDeploymentsResponse as
// a slice of AppDeployment values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListAppDeploymentsResponse) GetAppDeployments(ctx context.Context) ([]AppDeployment, bool) {
	if m.AppDeployments.IsNull() || m.AppDeployments.IsUnknown() {
		return nil, false
	}
	var v []AppDeployment
	d := m.AppDeployments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAppDeployments sets the value of the AppDeployments field in ListAppDeploymentsResponse.
func (m *ListAppDeploymentsResponse) SetAppDeployments(ctx context.Context, v []AppDeployment) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["app_deployments"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AppDeployments = types.ListValueMust(t, vs)
}

type ListAppsRequest struct {
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page of apps. Requests first page if
	// absent.
	PageToken types.String `tfsdk:"-"`
	// Filter apps by app space name. When specified, only apps belonging to
	// this space are returned.
	Space types.String `tfsdk:"-"`
}

func (to *ListAppsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListAppsRequest) {
}

func (to *ListAppsRequest) SyncFieldsDuringRead(ctx context.Context, from ListAppsRequest) {
}

func (m ListAppsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["space"] = attrs["space"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAppsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListAppsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAppsRequest
// only implements ToObjectValue() and Type().
func (m ListAppsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"space":      m.Space,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListAppsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"space":      types.StringType,
		},
	}
}

type ListAppsResponse struct {
	Apps types.List `tfsdk:"apps"`
	// Pagination token to request the next page of apps.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListAppsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListAppsResponse) {
	if !from.Apps.IsNull() && !from.Apps.IsUnknown() && to.Apps.IsNull() && len(from.Apps.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Apps, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Apps = from.Apps
	}
}

func (to *ListAppsResponse) SyncFieldsDuringRead(ctx context.Context, from ListAppsResponse) {
	if !from.Apps.IsNull() && !from.Apps.IsUnknown() && to.Apps.IsNull() && len(from.Apps.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Apps, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Apps = from.Apps
	}
}

func (m ListAppsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["apps"] = attrs["apps"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAppsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListAppsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"apps": reflect.TypeOf(App{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAppsResponse
// only implements ToObjectValue() and Type().
func (m ListAppsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"apps":            m.Apps,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListAppsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"apps": basetypes.ListType{
				ElemType: App{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetApps returns the value of the Apps field in ListAppsResponse as
// a slice of App values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListAppsResponse) GetApps(ctx context.Context) ([]App, bool) {
	if m.Apps.IsNull() || m.Apps.IsUnknown() {
		return nil, false
	}
	var v []App
	d := m.Apps.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetApps sets the value of the Apps field in ListAppsResponse.
func (m *ListAppsResponse) SetApps(ctx context.Context, v []App) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["apps"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Apps = types.ListValueMust(t, vs)
}

type ListCustomTemplatesRequest struct {
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page of custom templates. Requests
	// first page if absent.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListCustomTemplatesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListCustomTemplatesRequest) {
}

func (to *ListCustomTemplatesRequest) SyncFieldsDuringRead(ctx context.Context, from ListCustomTemplatesRequest) {
}

func (m ListCustomTemplatesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCustomTemplatesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListCustomTemplatesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCustomTemplatesRequest
// only implements ToObjectValue() and Type().
func (m ListCustomTemplatesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListCustomTemplatesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListCustomTemplatesResponse struct {
	// Pagination token to request the next page of custom templates.
	NextPageToken types.String `tfsdk:"next_page_token"`

	Templates types.List `tfsdk:"templates"`
}

func (to *ListCustomTemplatesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListCustomTemplatesResponse) {
	if !from.Templates.IsNull() && !from.Templates.IsUnknown() && to.Templates.IsNull() && len(from.Templates.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Templates, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Templates = from.Templates
	}
}

func (to *ListCustomTemplatesResponse) SyncFieldsDuringRead(ctx context.Context, from ListCustomTemplatesResponse) {
	if !from.Templates.IsNull() && !from.Templates.IsUnknown() && to.Templates.IsNull() && len(from.Templates.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Templates, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Templates = from.Templates
	}
}

func (m ListCustomTemplatesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["templates"] = attrs["templates"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCustomTemplatesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListCustomTemplatesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"templates": reflect.TypeOf(CustomTemplate{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCustomTemplatesResponse
// only implements ToObjectValue() and Type().
func (m ListCustomTemplatesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"templates":       m.Templates,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListCustomTemplatesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"templates": basetypes.ListType{
				ElemType: CustomTemplate{}.Type(ctx),
			},
		},
	}
}

// GetTemplates returns the value of the Templates field in ListCustomTemplatesResponse as
// a slice of CustomTemplate values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListCustomTemplatesResponse) GetTemplates(ctx context.Context) ([]CustomTemplate, bool) {
	if m.Templates.IsNull() || m.Templates.IsUnknown() {
		return nil, false
	}
	var v []CustomTemplate
	d := m.Templates.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTemplates sets the value of the Templates field in ListCustomTemplatesResponse.
func (m *ListCustomTemplatesResponse) SetTemplates(ctx context.Context, v []CustomTemplate) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["templates"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Templates = types.ListValueMust(t, vs)
}

type ListSpacesRequest struct {
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page of app spaces. Requests first
	// page if absent.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListSpacesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSpacesRequest) {
}

func (to *ListSpacesRequest) SyncFieldsDuringRead(ctx context.Context, from ListSpacesRequest) {
}

func (m ListSpacesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSpacesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListSpacesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSpacesRequest
// only implements ToObjectValue() and Type().
func (m ListSpacesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSpacesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListSpacesResponse struct {
	// Pagination token to request the next page of app spaces.
	NextPageToken types.String `tfsdk:"next_page_token"`

	Spaces types.List `tfsdk:"spaces"`
}

func (to *ListSpacesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSpacesResponse) {
	if !from.Spaces.IsNull() && !from.Spaces.IsUnknown() && to.Spaces.IsNull() && len(from.Spaces.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Spaces, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Spaces = from.Spaces
	}
}

func (to *ListSpacesResponse) SyncFieldsDuringRead(ctx context.Context, from ListSpacesResponse) {
	if !from.Spaces.IsNull() && !from.Spaces.IsUnknown() && to.Spaces.IsNull() && len(from.Spaces.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Spaces, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Spaces = from.Spaces
	}
}

func (m ListSpacesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["spaces"] = attrs["spaces"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSpacesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListSpacesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"spaces": reflect.TypeOf(Space{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSpacesResponse
// only implements ToObjectValue() and Type().
func (m ListSpacesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"spaces":          m.Spaces,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSpacesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"spaces": basetypes.ListType{
				ElemType: Space{}.Type(ctx),
			},
		},
	}
}

// GetSpaces returns the value of the Spaces field in ListSpacesResponse as
// a slice of Space values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListSpacesResponse) GetSpaces(ctx context.Context) ([]Space, bool) {
	if m.Spaces.IsNull() || m.Spaces.IsUnknown() {
		return nil, false
	}
	var v []Space
	d := m.Spaces.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSpaces sets the value of the Spaces field in ListSpacesResponse.
func (m *ListSpacesResponse) SetSpaces(ctx context.Context, v []Space) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["spaces"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Spaces = types.ListValueMust(t, vs)
}

// This resource represents a long-running operation that is the result of a
// network API call.
type Operation struct {
	// If the value is `false`, it means the operation is still in progress. If
	// `true`, the operation is completed, and either `error` or `response` is
	// available.
	Done types.Bool `tfsdk:"done"`
	// The error result of the operation in case of failure or cancellation.
	Error types.Object `tfsdk:"error"`
	// Service-specific metadata associated with the operation. It typically
	// contains progress information and common metadata such as create time.
	// Some services might not provide such metadata.
	Metadata types.Object `tfsdk:"metadata"`
	// The server-assigned name, which is only unique within the same service
	// that originally returns it. If you use the default HTTP mapping, the
	// `name` should be a resource name ending with `operations/{unique_id}`.
	Name types.String `tfsdk:"name"`
	// The normal, successful response of the operation.
	Response types.Object `tfsdk:"response"`
}

func (to *Operation) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Operation) {
	if !from.Error.IsNull() && !from.Error.IsUnknown() {
		if toError, ok := to.GetError(ctx); ok {
			if fromError, ok := from.GetError(ctx); ok {
				// Recursively sync the fields of Error
				toError.SyncFieldsDuringCreateOrUpdate(ctx, fromError)
				to.SetError(ctx, toError)
			}
		}
	}
}

func (to *Operation) SyncFieldsDuringRead(ctx context.Context, from Operation) {
	if !from.Error.IsNull() && !from.Error.IsUnknown() {
		if toError, ok := to.GetError(ctx); ok {
			if fromError, ok := from.GetError(ctx); ok {
				toError.SyncFieldsDuringRead(ctx, fromError)
				to.SetError(ctx, toError)
			}
		}
	}
}

func (m Operation) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["done"] = attrs["done"].SetOptional()
	attrs["error"] = attrs["error"].SetOptional()
	attrs["metadata"] = attrs["metadata"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["response"] = attrs["response"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Operation.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Operation) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"error": reflect.TypeOf(DatabricksServiceExceptionWithDetailsProto{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Operation
// only implements ToObjectValue() and Type().
func (m Operation) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"done":     m.Done,
			"error":    m.Error,
			"metadata": m.Metadata,
			"name":     m.Name,
			"response": m.Response,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Operation) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"done":     types.BoolType,
			"error":    DatabricksServiceExceptionWithDetailsProto{}.Type(ctx),
			"metadata": types.ObjectType{},
			"name":     types.StringType,
			"response": types.ObjectType{},
		},
	}
}

// GetError returns the value of the Error field in Operation as
// a DatabricksServiceExceptionWithDetailsProto value.
// If the field is unknown or null, the boolean return value is false.
func (m *Operation) GetError(ctx context.Context) (DatabricksServiceExceptionWithDetailsProto, bool) {
	var e DatabricksServiceExceptionWithDetailsProto
	if m.Error.IsNull() || m.Error.IsUnknown() {
		return e, false
	}
	var v DatabricksServiceExceptionWithDetailsProto
	d := m.Error.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetError sets the value of the Error field in Operation.
func (m *Operation) SetError(ctx context.Context, v DatabricksServiceExceptionWithDetailsProto) {
	vs := v.ToObjectValue(ctx)
	m.Error = vs
}

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
	UserApiScopes types.List `tfsdk:"user_api_scopes"`
}

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
}

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
}

func (m Space) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["creator"] = attrs["creator"].SetComputed()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["effective_usage_policy_id"] = attrs["effective_usage_policy_id"].SetComputed()
	attrs["effective_user_api_scopes"] = attrs["effective_user_api_scopes"].SetComputed()
	attrs["id"] = attrs["id"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
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

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Space.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Space) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"effective_user_api_scopes": reflect.TypeOf(types.String{}),
		"resources":                 reflect.TypeOf(AppResource{}),
		"status":                    reflect.TypeOf(SpaceStatus{}),
		"user_api_scopes":           reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Space
// only implements ToObjectValue() and Type().
func (m Space) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time":                 m.CreateTime,
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
		})
}

// Type implements basetypes.ObjectValuable.
func (m Space) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":               timetypes.RFC3339{}.Type(ctx),
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
				ElemType: AppResource{}.Type(ctx),
			},
			"service_principal_client_id": types.StringType,
			"service_principal_id":        types.Int64Type,
			"service_principal_name":      types.StringType,
			"status":                      SpaceStatus{}.Type(ctx),
			"update_time":                 timetypes.RFC3339{}.Type(ctx),
			"updater":                     types.StringType,
			"usage_policy_id":             types.StringType,
			"user_api_scopes": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
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
// a slice of AppResource values.
// If the field is unknown or null, the boolean return value is false.
func (m *Space) GetResources(ctx context.Context) ([]AppResource, bool) {
	if m.Resources.IsNull() || m.Resources.IsUnknown() {
		return nil, false
	}
	var v []AppResource
	d := m.Resources.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResources sets the value of the Resources field in Space.
func (m *Space) SetResources(ctx context.Context, v []AppResource) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["resources"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Resources = types.ListValueMust(t, vs)
}

// GetStatus returns the value of the Status field in Space as
// a SpaceStatus value.
// If the field is unknown or null, the boolean return value is false.
func (m *Space) GetStatus(ctx context.Context) (SpaceStatus, bool) {
	var e SpaceStatus
	if m.Status.IsNull() || m.Status.IsUnknown() {
		return e, false
	}
	var v SpaceStatus
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
func (m *Space) SetStatus(ctx context.Context, v SpaceStatus) {
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

type SpaceStatus struct {
	// Message providing context about the current state.
	Message types.String `tfsdk:"message"`
	// The state of the app space.
	State types.String `tfsdk:"state"`
}

func (to *SpaceStatus) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SpaceStatus) {
}

func (to *SpaceStatus) SyncFieldsDuringRead(ctx context.Context, from SpaceStatus) {
}

func (m SpaceStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["message"] = attrs["message"].SetComputed()
	attrs["state"] = attrs["state"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SpaceStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SpaceStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SpaceStatus
// only implements ToObjectValue() and Type().
func (m SpaceStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"message": m.Message,
			"state":   m.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SpaceStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"message": types.StringType,
			"state":   types.StringType,
		},
	}
}

// Tracks app space update information.
type SpaceUpdate struct {
	Description types.String `tfsdk:"description"`

	Resources types.List `tfsdk:"resources"`

	Status types.Object `tfsdk:"status"`

	UsagePolicyId types.String `tfsdk:"usage_policy_id"`

	UserApiScopes types.List `tfsdk:"user_api_scopes"`
}

func (to *SpaceUpdate) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SpaceUpdate) {
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
}

func (to *SpaceUpdate) SyncFieldsDuringRead(ctx context.Context, from SpaceUpdate) {
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
}

func (m SpaceUpdate) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["resources"] = attrs["resources"].SetOptional()
	attrs["status"] = attrs["status"].SetComputed()
	attrs["usage_policy_id"] = attrs["usage_policy_id"].SetOptional()
	attrs["user_api_scopes"] = attrs["user_api_scopes"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SpaceUpdate.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SpaceUpdate) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"resources":       reflect.TypeOf(AppResource{}),
		"status":          reflect.TypeOf(SpaceUpdateStatus{}),
		"user_api_scopes": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SpaceUpdate
// only implements ToObjectValue() and Type().
func (m SpaceUpdate) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":     m.Description,
			"resources":       m.Resources,
			"status":          m.Status,
			"usage_policy_id": m.UsagePolicyId,
			"user_api_scopes": m.UserApiScopes,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SpaceUpdate) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"resources": basetypes.ListType{
				ElemType: AppResource{}.Type(ctx),
			},
			"status":          SpaceUpdateStatus{}.Type(ctx),
			"usage_policy_id": types.StringType,
			"user_api_scopes": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetResources returns the value of the Resources field in SpaceUpdate as
// a slice of AppResource values.
// If the field is unknown or null, the boolean return value is false.
func (m *SpaceUpdate) GetResources(ctx context.Context) ([]AppResource, bool) {
	if m.Resources.IsNull() || m.Resources.IsUnknown() {
		return nil, false
	}
	var v []AppResource
	d := m.Resources.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResources sets the value of the Resources field in SpaceUpdate.
func (m *SpaceUpdate) SetResources(ctx context.Context, v []AppResource) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["resources"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Resources = types.ListValueMust(t, vs)
}

// GetStatus returns the value of the Status field in SpaceUpdate as
// a SpaceUpdateStatus value.
// If the field is unknown or null, the boolean return value is false.
func (m *SpaceUpdate) GetStatus(ctx context.Context) (SpaceUpdateStatus, bool) {
	var e SpaceUpdateStatus
	if m.Status.IsNull() || m.Status.IsUnknown() {
		return e, false
	}
	var v SpaceUpdateStatus
	d := m.Status.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatus sets the value of the Status field in SpaceUpdate.
func (m *SpaceUpdate) SetStatus(ctx context.Context, v SpaceUpdateStatus) {
	vs := v.ToObjectValue(ctx)
	m.Status = vs
}

// GetUserApiScopes returns the value of the UserApiScopes field in SpaceUpdate as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *SpaceUpdate) GetUserApiScopes(ctx context.Context) ([]types.String, bool) {
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

// SetUserApiScopes sets the value of the UserApiScopes field in SpaceUpdate.
func (m *SpaceUpdate) SetUserApiScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["user_api_scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.UserApiScopes = types.ListValueMust(t, vs)
}

// Status of an app space update operation
type SpaceUpdateStatus struct {
	Message types.String `tfsdk:"message"`

	State types.String `tfsdk:"state"`
}

func (to *SpaceUpdateStatus) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SpaceUpdateStatus) {
}

func (to *SpaceUpdateStatus) SyncFieldsDuringRead(ctx context.Context, from SpaceUpdateStatus) {
}

func (m SpaceUpdateStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["message"] = attrs["message"].SetComputed()
	attrs["state"] = attrs["state"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SpaceUpdateStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SpaceUpdateStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SpaceUpdateStatus
// only implements ToObjectValue() and Type().
func (m SpaceUpdateStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"message": m.Message,
			"state":   m.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SpaceUpdateStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"message": types.StringType,
			"state":   types.StringType,
		},
	}
}

type StartAppRequest struct {
	// The name of the app.
	Name types.String `tfsdk:"-"`
}

func (to *StartAppRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StartAppRequest) {
}

func (to *StartAppRequest) SyncFieldsDuringRead(ctx context.Context, from StartAppRequest) {
}

func (m StartAppRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StartAppRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m StartAppRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StartAppRequest
// only implements ToObjectValue() and Type().
func (m StartAppRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m StartAppRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type StopAppRequest struct {
	// The name of the app.
	Name types.String `tfsdk:"-"`
}

func (to *StopAppRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StopAppRequest) {
}

func (to *StopAppRequest) SyncFieldsDuringRead(ctx context.Context, from StopAppRequest) {
}

func (m StopAppRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StopAppRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m StopAppRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StopAppRequest
// only implements ToObjectValue() and Type().
func (m StopAppRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m StopAppRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type UpdateAppRequest struct {
	App types.Object `tfsdk:"app"`
	// The name of the app. The name must contain only lowercase alphanumeric
	// characters and hyphens. It must be unique within the workspace.
	Name types.String `tfsdk:"-"`
}

func (to *UpdateAppRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateAppRequest) {
	if !from.App.IsNull() && !from.App.IsUnknown() {
		if toApp, ok := to.GetApp(ctx); ok {
			if fromApp, ok := from.GetApp(ctx); ok {
				// Recursively sync the fields of App
				toApp.SyncFieldsDuringCreateOrUpdate(ctx, fromApp)
				to.SetApp(ctx, toApp)
			}
		}
	}
}

func (to *UpdateAppRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateAppRequest) {
	if !from.App.IsNull() && !from.App.IsUnknown() {
		if toApp, ok := to.GetApp(ctx); ok {
			if fromApp, ok := from.GetApp(ctx); ok {
				toApp.SyncFieldsDuringRead(ctx, fromApp)
				to.SetApp(ctx, toApp)
			}
		}
	}
}

func (m UpdateAppRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["app"] = attrs["app"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateAppRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateAppRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"app": reflect.TypeOf(App{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAppRequest
// only implements ToObjectValue() and Type().
func (m UpdateAppRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app":  m.App,
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateAppRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app":  App{}.Type(ctx),
			"name": types.StringType,
		},
	}
}

// GetApp returns the value of the App field in UpdateAppRequest as
// a App value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateAppRequest) GetApp(ctx context.Context) (App, bool) {
	var e App
	if m.App.IsNull() || m.App.IsUnknown() {
		return e, false
	}
	var v App
	d := m.App.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetApp sets the value of the App field in UpdateAppRequest.
func (m *UpdateAppRequest) SetApp(ctx context.Context, v App) {
	vs := v.ToObjectValue(ctx)
	m.App = vs
}

type UpdateCustomTemplateRequest struct {
	// The name of the template. It must contain only alphanumeric characters,
	// hyphens, underscores, and whitespaces. It must be unique within the
	// workspace.
	Name types.String `tfsdk:"-"`

	Template types.Object `tfsdk:"template"`
}

func (to *UpdateCustomTemplateRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateCustomTemplateRequest) {
	if !from.Template.IsNull() && !from.Template.IsUnknown() {
		if toTemplate, ok := to.GetTemplate(ctx); ok {
			if fromTemplate, ok := from.GetTemplate(ctx); ok {
				// Recursively sync the fields of Template
				toTemplate.SyncFieldsDuringCreateOrUpdate(ctx, fromTemplate)
				to.SetTemplate(ctx, toTemplate)
			}
		}
	}
}

func (to *UpdateCustomTemplateRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateCustomTemplateRequest) {
	if !from.Template.IsNull() && !from.Template.IsUnknown() {
		if toTemplate, ok := to.GetTemplate(ctx); ok {
			if fromTemplate, ok := from.GetTemplate(ctx); ok {
				toTemplate.SyncFieldsDuringRead(ctx, fromTemplate)
				to.SetTemplate(ctx, toTemplate)
			}
		}
	}
}

func (m UpdateCustomTemplateRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["template"] = attrs["template"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCustomTemplateRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateCustomTemplateRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"template": reflect.TypeOf(CustomTemplate{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCustomTemplateRequest
// only implements ToObjectValue() and Type().
func (m UpdateCustomTemplateRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":     m.Name,
			"template": m.Template,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateCustomTemplateRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":     types.StringType,
			"template": CustomTemplate{}.Type(ctx),
		},
	}
}

// GetTemplate returns the value of the Template field in UpdateCustomTemplateRequest as
// a CustomTemplate value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateCustomTemplateRequest) GetTemplate(ctx context.Context) (CustomTemplate, bool) {
	var e CustomTemplate
	if m.Template.IsNull() || m.Template.IsUnknown() {
		return e, false
	}
	var v CustomTemplate
	d := m.Template.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTemplate sets the value of the Template field in UpdateCustomTemplateRequest.
func (m *UpdateCustomTemplateRequest) SetTemplate(ctx context.Context, v CustomTemplate) {
	vs := v.ToObjectValue(ctx)
	m.Template = vs
}

type UpdateSpaceRequest struct {
	// The name of the app space. The name must contain only lowercase
	// alphanumeric characters and hyphens. It must be unique within the
	// workspace.
	Name types.String `tfsdk:"-"`

	Space types.Object `tfsdk:"space"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateSpaceRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateSpaceRequest) {
	if !from.Space.IsNull() && !from.Space.IsUnknown() {
		if toSpace, ok := to.GetSpace(ctx); ok {
			if fromSpace, ok := from.GetSpace(ctx); ok {
				// Recursively sync the fields of Space
				toSpace.SyncFieldsDuringCreateOrUpdate(ctx, fromSpace)
				to.SetSpace(ctx, toSpace)
			}
		}
	}
}

func (to *UpdateSpaceRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateSpaceRequest) {
	if !from.Space.IsNull() && !from.Space.IsUnknown() {
		if toSpace, ok := to.GetSpace(ctx); ok {
			if fromSpace, ok := from.GetSpace(ctx); ok {
				toSpace.SyncFieldsDuringRead(ctx, fromSpace)
				to.SetSpace(ctx, toSpace)
			}
		}
	}
}

func (m UpdateSpaceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["space"] = attrs["space"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateSpaceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateSpaceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"space": reflect.TypeOf(Space{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateSpaceRequest
// only implements ToObjectValue() and Type().
func (m UpdateSpaceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":        m.Name,
			"space":       m.Space,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateSpaceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":        types.StringType,
			"space":       Space{}.Type(ctx),
			"update_mask": types.StringType,
		},
	}
}

// GetSpace returns the value of the Space field in UpdateSpaceRequest as
// a Space value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateSpaceRequest) GetSpace(ctx context.Context) (Space, bool) {
	var e Space
	if m.Space.IsNull() || m.Space.IsUnknown() {
		return e, false
	}
	var v Space
	d := m.Space.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSpace sets the value of the Space field in UpdateSpaceRequest.
func (m *UpdateSpaceRequest) SetSpace(ctx context.Context, v Space) {
	vs := v.ToObjectValue(ctx)
	m.Space = vs
}
