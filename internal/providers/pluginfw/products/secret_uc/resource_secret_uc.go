// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package secret_uc

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/declarative"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
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

const resourceName = "secret_uc"

var _ resource.ResourceWithConfigure = &SecretResource{}
var _ resource.ResourceWithModifyPlan = &SecretResource{}

func ResourceSecret() resource.Resource {
	return &SecretResource{}
}

type SecretResource struct {
	Client *autogen.DatabricksClient
}

// ProviderConfig contains the fields to configure the provider.
type ProviderConfig struct {
	WorkspaceID types.String `tfsdk:"workspace_id"`
}

// ApplySchemaCustomizations applies the schema customizations to the ProviderConfig type.
func (r ProviderConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_id"] = attrs["workspace_id"].SetOptional()
	attrs["workspace_id"] = attrs["workspace_id"].SetComputed()
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

// Secret extends the main model with additional fields.
type Secret struct {
	// Indicates whether the principal is limited to retrieving metadata for the
	// associated object through the **BROWSE** privilege when
	// **include_browse** is enabled in the request.
	BrowseOnly types.Bool `tfsdk:"browse_only"`
	// The name of the catalog where the schema and the secret reside.
	CatalogName types.String `tfsdk:"catalog_name"`
	// User-provided free-form text description of the secret.
	Comment types.String `tfsdk:"comment"`
	// The time at which this secret was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// The principal that created the secret.
	CreatedBy types.String `tfsdk:"created_by"`
	// The effective owner of the secret, which may differ from the directly-set
	// **owner** due to inheritance.
	EffectiveOwner types.String `tfsdk:"effective_owner"`
	// The secret value. Only populated in responses when you have the
	// **READ_SECRET** privilege and **include_value** is set to true in the
	// request. The maximum size is 60 KiB.
	EffectiveValue types.String `tfsdk:"effective_value"`
	// User-provided expiration time of the secret. This field indicates when
	// the secret should no longer be used and may be displayed as a warning in
	// the UI. It is purely informational and does not trigger any automatic
	// actions or affect the secret's lifecycle.
	ExpireTime timetypes.RFC3339 `tfsdk:"expire_time"`

	ExternalSecretId types.String `tfsdk:"external_secret_id"`
	// The three-level (fully qualified) name of the secret, in the form of
	// **catalog_name.schema_name.secret_name**.
	FullName types.String `tfsdk:"full_name"`
	// Unique identifier of the metastore hosting the secret.
	MetastoreId types.String `tfsdk:"metastore_id"`
	// The name of the secret, relative to its parent schema.
	Name types.String `tfsdk:"name"`
	// The owner of the secret. Defaults to the creating principal on creation.
	// Can be updated to transfer ownership of the secret to another principal.
	Owner types.String `tfsdk:"owner"`
	// The name of the schema where the secret resides.
	SchemaName types.String `tfsdk:"schema_name"`
	// The time at which this secret was last updated.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
	// The principal that last updated the secret.
	UpdatedBy types.String `tfsdk:"updated_by"`
	// The secret value to store. This field is input-only and is not returned
	// in responses — use the **effective_value** field (via GetSecret with
	// **include_value** set to true) to read the secret value. The maximum size
	// is 60 KiB (pre-encryption). Accepted content includes passwords, tokens,
	// keys, and other sensitive credential data.
	Value          types.String `tfsdk:"value"`
	ProviderConfig types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// Secret struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m Secret) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"provider_config": reflect.TypeOf(ProviderConfig{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Secret
// only implements ToObjectValue() and Type().
func (m Secret) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"browse_only": m.BrowseOnly,
			"catalog_name":       m.CatalogName,
			"comment":            m.Comment,
			"create_time":        m.CreateTime,
			"created_by":         m.CreatedBy,
			"effective_owner":    m.EffectiveOwner,
			"effective_value":    m.EffectiveValue,
			"expire_time":        m.ExpireTime,
			"external_secret_id": m.ExternalSecretId,
			"full_name":          m.FullName,
			"metastore_id":       m.MetastoreId,
			"name":               m.Name,
			"owner":              m.Owner,
			"schema_name":        m.SchemaName,
			"update_time":        m.UpdateTime,
			"updated_by":         m.UpdatedBy,
			"value":              m.Value,

			"provider_config": m.ProviderConfig,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m Secret) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"browse_only": types.BoolType,
			"catalog_name":       types.StringType,
			"comment":            types.StringType,
			"create_time":        timetypes.RFC3339{}.Type(ctx),
			"created_by":         types.StringType,
			"effective_owner":    types.StringType,
			"effective_value":    types.StringType,
			"expire_time":        timetypes.RFC3339{}.Type(ctx),
			"external_secret_id": types.StringType,
			"full_name":          types.StringType,
			"metastore_id":       types.StringType,
			"name":               types.StringType,
			"owner":              types.StringType,
			"schema_name":        types.StringType,
			"update_time":        timetypes.RFC3339{}.Type(ctx),
			"updated_by":         types.StringType,
			"value":              types.StringType,

			"provider_config": ProviderConfig{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *Secret) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Secret) {
	if !from.Owner.IsUnknown() && !from.Owner.IsNull() {
		// Owner is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Owner = from.Owner
	}
	if !from.Value.IsUnknown() && !from.Value.IsNull() {
		// Value is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Value = from.Value
	}
	to.ProviderConfig = from.ProviderConfig

}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *Secret) SyncFieldsDuringRead(ctx context.Context, from Secret) {
	if !from.Owner.IsUnknown() && !from.Owner.IsNull() {
		// Owner is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Owner = from.Owner
	}
	if !from.Value.IsUnknown() && !from.Value.IsNull() {
		// Value is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Value = from.Value
	}
	to.ProviderConfig = from.ProviderConfig

}

func (m Secret) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["browse_only"] = attrs["browse_only"].SetComputed()
	attrs["catalog_name"] = attrs["catalog_name"].SetRequired()
	attrs["catalog_name"] = attrs["catalog_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["created_by"] = attrs["created_by"].SetComputed()
	attrs["effective_owner"] = attrs["effective_owner"].SetComputed()
	attrs["effective_value"] = attrs["effective_value"].SetComputed()
	attrs["expire_time"] = attrs["expire_time"].SetOptional()
	attrs["external_secret_id"] = attrs["external_secret_id"].SetComputed()
	attrs["full_name"] = attrs["full_name"].SetComputed()
	attrs["metastore_id"] = attrs["metastore_id"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["owner"] = attrs["owner"].SetComputed()
	attrs["owner"] = attrs["owner"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["schema_name"] = attrs["schema_name"].SetRequired()
	attrs["schema_name"] = attrs["schema_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["updated_by"] = attrs["updated_by"].SetComputed()
	attrs["value"] = attrs["value"].SetRequired()

	attrs["full_name"] = attrs["full_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	attrs["provider_config"] = attrs["provider_config"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(tfschema.ProviderConfigPlanModifier{})

	return attrs
}

func (r *SecretResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *SecretResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, Secret{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks secret_uc",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *SecretResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *SecretResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Skip entirely on destroy (no plan state).
	if req.Plan.Raw.IsNull() {
		return
	}
	if r.Client == nil {
		return
	}
	tfschema.WorkspaceDriftDetection(ctx, r.Client, req, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	tfschema.ValidateWorkspaceID(ctx, r.Client, req, resp)
}

func (r *SecretResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan Secret
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var secret catalog.Secret

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &secret)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := catalog.CreateSecretRequest{
		Secret: secret,
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

	response, err := client.SecretsUc.CreateSecret(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create secret_uc", err.Error())
		return
	}

	var newState Secret

	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInState(ctx, r.Client, plan.ProviderConfig, &resp.State)...)
}

func (r *SecretResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState Secret
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest catalog.GetSecretRequest
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
	response, err := client.SecretsUc.GetSecret(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to get secret_uc", err.Error())
		return
	}

	var newState Secret
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInState(ctx, r.Client, existingState.ProviderConfig, &resp.State)...)
}

func (r *SecretResource) update(ctx context.Context, plan Secret, diags *diag.Diagnostics, state *tfsdk.State) {
	var secret catalog.Secret

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &secret)...)
	if diags.HasError() {
		return
	}

	updateRequest := catalog.UpdateSecretRequest{
		Secret:     secret,
		FullName:   plan.FullName.ValueString(),
		UpdateMask: *fieldmask.New(strings.Split("comment,expire_time,owner,value", ",")),
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
	response, err := client.SecretsUc.UpdateSecret(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update secret_uc", err.Error())
		return
	}

	var newState Secret

	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *SecretResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan Secret
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *SecretResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state Secret
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest catalog.DeleteSecretRequest
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

	err := client.SecretsUc.DeleteSecret(ctx, deleteRequest)
	if !declarative.IsDeleteError(err) {
		err = nil
	}
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete secret_uc", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &SecretResource{}

func (r *SecretResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, ",")

	if len(parts) != 1 || parts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf(
				"Expected import identifier with format: full_name. Got: %q",
				req.ID,
			),
		)
		return
	}

	fullName := parts[0]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("full_name"), fullName)...)
}
