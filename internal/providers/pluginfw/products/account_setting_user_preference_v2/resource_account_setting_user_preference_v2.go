// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package account_setting_user_preference_v2

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/settingsv2"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/settingsv2_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const resourceName = "account_setting_user_preference_v2"

var _ resource.ResourceWithConfigure = &UserPreferenceResource{}

func ResourceUserPreference() resource.Resource {
	return &UserPreferenceResource{}
}

type UserPreferenceResource struct {
	Client *autogen.DatabricksClient
}

// UserPreference extends the main model with additional fields.
type UserPreference struct {
	BooleanVal types.Object `tfsdk:"boolean_val"`

	EffectiveBooleanVal types.Object `tfsdk:"effective_boolean_val"`

	EffectiveStringVal types.Object `tfsdk:"effective_string_val"`
	// Name of the setting.
	Name types.String `tfsdk:"name"`

	StringVal types.Object `tfsdk:"string_val"`
	// User ID of the user.
	UserId types.String `tfsdk:"user_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// UserPreference struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m UserPreference) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"boolean_val":           reflect.TypeOf(settingsv2_tf.BooleanMessage{}),
		"effective_boolean_val": reflect.TypeOf(settingsv2_tf.BooleanMessage{}),
		"effective_string_val":  reflect.TypeOf(settingsv2_tf.StringMessage{}),
		"string_val":            reflect.TypeOf(settingsv2_tf.StringMessage{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UserPreference
// only implements ToObjectValue() and Type().
func (m UserPreference) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"boolean_val": m.BooleanVal,
			"effective_boolean_val": m.EffectiveBooleanVal,
			"effective_string_val":  m.EffectiveStringVal,
			"name":                  m.Name,
			"string_val":            m.StringVal,
			"user_id":               m.UserId,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m UserPreference) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"boolean_val": settingsv2_tf.BooleanMessage{}.Type(ctx),
			"effective_boolean_val": settingsv2_tf.BooleanMessage{}.Type(ctx),
			"effective_string_val":  settingsv2_tf.StringMessage{}.Type(ctx),
			"name":                  types.StringType,
			"string_val":            settingsv2_tf.StringMessage{}.Type(ctx),
			"user_id":               types.StringType,
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *UserPreference) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UserPreference) {
	if !from.BooleanVal.IsNull() && !from.BooleanVal.IsUnknown() {
		if toBooleanVal, ok := to.GetBooleanVal(ctx); ok {
			if fromBooleanVal, ok := from.GetBooleanVal(ctx); ok {
				// Recursively sync the fields of BooleanVal
				toBooleanVal.SyncFieldsDuringCreateOrUpdate(ctx, fromBooleanVal)
				to.SetBooleanVal(ctx, toBooleanVal)
			}
		}
	}
	if !from.EffectiveBooleanVal.IsNull() && !from.EffectiveBooleanVal.IsUnknown() {
		if toEffectiveBooleanVal, ok := to.GetEffectiveBooleanVal(ctx); ok {
			if fromEffectiveBooleanVal, ok := from.GetEffectiveBooleanVal(ctx); ok {
				// Recursively sync the fields of EffectiveBooleanVal
				toEffectiveBooleanVal.SyncFieldsDuringCreateOrUpdate(ctx, fromEffectiveBooleanVal)
				to.SetEffectiveBooleanVal(ctx, toEffectiveBooleanVal)
			}
		}
	}
	if !from.EffectiveStringVal.IsNull() && !from.EffectiveStringVal.IsUnknown() {
		if toEffectiveStringVal, ok := to.GetEffectiveStringVal(ctx); ok {
			if fromEffectiveStringVal, ok := from.GetEffectiveStringVal(ctx); ok {
				// Recursively sync the fields of EffectiveStringVal
				toEffectiveStringVal.SyncFieldsDuringCreateOrUpdate(ctx, fromEffectiveStringVal)
				to.SetEffectiveStringVal(ctx, toEffectiveStringVal)
			}
		}
	}
	if !from.StringVal.IsNull() && !from.StringVal.IsUnknown() {
		if toStringVal, ok := to.GetStringVal(ctx); ok {
			if fromStringVal, ok := from.GetStringVal(ctx); ok {
				// Recursively sync the fields of StringVal
				toStringVal.SyncFieldsDuringCreateOrUpdate(ctx, fromStringVal)
				to.SetStringVal(ctx, toStringVal)
			}
		}
	}
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *UserPreference) SyncFieldsDuringRead(ctx context.Context, from UserPreference) {
	if !from.BooleanVal.IsNull() && !from.BooleanVal.IsUnknown() {
		if toBooleanVal, ok := to.GetBooleanVal(ctx); ok {
			if fromBooleanVal, ok := from.GetBooleanVal(ctx); ok {
				toBooleanVal.SyncFieldsDuringRead(ctx, fromBooleanVal)
				to.SetBooleanVal(ctx, toBooleanVal)
			}
		}
	}
	if !from.EffectiveBooleanVal.IsNull() && !from.EffectiveBooleanVal.IsUnknown() {
		if toEffectiveBooleanVal, ok := to.GetEffectiveBooleanVal(ctx); ok {
			if fromEffectiveBooleanVal, ok := from.GetEffectiveBooleanVal(ctx); ok {
				toEffectiveBooleanVal.SyncFieldsDuringRead(ctx, fromEffectiveBooleanVal)
				to.SetEffectiveBooleanVal(ctx, toEffectiveBooleanVal)
			}
		}
	}
	if !from.EffectiveStringVal.IsNull() && !from.EffectiveStringVal.IsUnknown() {
		if toEffectiveStringVal, ok := to.GetEffectiveStringVal(ctx); ok {
			if fromEffectiveStringVal, ok := from.GetEffectiveStringVal(ctx); ok {
				toEffectiveStringVal.SyncFieldsDuringRead(ctx, fromEffectiveStringVal)
				to.SetEffectiveStringVal(ctx, toEffectiveStringVal)
			}
		}
	}
	if !from.StringVal.IsNull() && !from.StringVal.IsUnknown() {
		if toStringVal, ok := to.GetStringVal(ctx); ok {
			if fromStringVal, ok := from.GetStringVal(ctx); ok {
				toStringVal.SyncFieldsDuringRead(ctx, fromStringVal)
				to.SetStringVal(ctx, toStringVal)
			}
		}
	}
}

func (m UserPreference) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["boolean_val"] = attrs["boolean_val"].SetOptional()
	attrs["effective_boolean_val"] = attrs["effective_boolean_val"].SetComputed()
	attrs["effective_string_val"] = attrs["effective_string_val"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["string_val"] = attrs["string_val"].SetOptional()
	attrs["user_id"] = attrs["user_id"].SetOptional()

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["user_id"] = attrs["user_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	return attrs
}

// GetBooleanVal returns the value of the BooleanVal field in UserPreference as
// a settingsv2_tf.BooleanMessage value.
// If the field is unknown or null, the boolean return value is false.
func (m *UserPreference) GetBooleanVal(ctx context.Context) (settingsv2_tf.BooleanMessage, bool) {
	var e settingsv2_tf.BooleanMessage
	if m.BooleanVal.IsNull() || m.BooleanVal.IsUnknown() {
		return e, false
	}
	var v settingsv2_tf.BooleanMessage
	d := m.BooleanVal.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBooleanVal sets the value of the BooleanVal field in UserPreference.
func (m *UserPreference) SetBooleanVal(ctx context.Context, v settingsv2_tf.BooleanMessage) {
	vs := v.ToObjectValue(ctx)
	m.BooleanVal = vs
}

// GetEffectiveBooleanVal returns the value of the EffectiveBooleanVal field in UserPreference as
// a settingsv2_tf.BooleanMessage value.
// If the field is unknown or null, the boolean return value is false.
func (m *UserPreference) GetEffectiveBooleanVal(ctx context.Context) (settingsv2_tf.BooleanMessage, bool) {
	var e settingsv2_tf.BooleanMessage
	if m.EffectiveBooleanVal.IsNull() || m.EffectiveBooleanVal.IsUnknown() {
		return e, false
	}
	var v settingsv2_tf.BooleanMessage
	d := m.EffectiveBooleanVal.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveBooleanVal sets the value of the EffectiveBooleanVal field in UserPreference.
func (m *UserPreference) SetEffectiveBooleanVal(ctx context.Context, v settingsv2_tf.BooleanMessage) {
	vs := v.ToObjectValue(ctx)
	m.EffectiveBooleanVal = vs
}

// GetEffectiveStringVal returns the value of the EffectiveStringVal field in UserPreference as
// a settingsv2_tf.StringMessage value.
// If the field is unknown or null, the boolean return value is false.
func (m *UserPreference) GetEffectiveStringVal(ctx context.Context) (settingsv2_tf.StringMessage, bool) {
	var e settingsv2_tf.StringMessage
	if m.EffectiveStringVal.IsNull() || m.EffectiveStringVal.IsUnknown() {
		return e, false
	}
	var v settingsv2_tf.StringMessage
	d := m.EffectiveStringVal.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveStringVal sets the value of the EffectiveStringVal field in UserPreference.
func (m *UserPreference) SetEffectiveStringVal(ctx context.Context, v settingsv2_tf.StringMessage) {
	vs := v.ToObjectValue(ctx)
	m.EffectiveStringVal = vs
}

// GetStringVal returns the value of the StringVal field in UserPreference as
// a settingsv2_tf.StringMessage value.
// If the field is unknown or null, the boolean return value is false.
func (m *UserPreference) GetStringVal(ctx context.Context) (settingsv2_tf.StringMessage, bool) {
	var e settingsv2_tf.StringMessage
	if m.StringVal.IsNull() || m.StringVal.IsUnknown() {
		return e, false
	}
	var v settingsv2_tf.StringMessage
	d := m.StringVal.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStringVal sets the value of the StringVal field in UserPreference.
func (m *UserPreference) SetStringVal(ctx context.Context, v settingsv2_tf.StringMessage) {
	vs := v.ToObjectValue(ctx)
	m.StringVal = vs
}

func (r *UserPreferenceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *UserPreferenceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, UserPreference{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks account_setting_user_preference_v2",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *UserPreferenceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *UserPreferenceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan UserPreference
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *UserPreferenceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState UserPreference
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest settingsv2.GetPublicAccountUserPreferenceRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetAccountClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}
	response, err := client.SettingsV2.GetPublicAccountUserPreference(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get account_setting_user_preference_v2", err.Error())
		return
	}

	var newState UserPreference
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *UserPreferenceResource) update(ctx context.Context, plan UserPreference, diags *diag.Diagnostics, state *tfsdk.State) {
	var user_preference settingsv2.UserPreference

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &user_preference)...)
	if diags.HasError() {
		return
	}

	updateRequest := settingsv2.PatchPublicAccountUserPreferenceRequest{
		Setting: user_preference,
		Name:    plan.Name.ValueString(),
		UserId:  plan.UserId.ValueString(),
	}

	client, clientDiags := r.Client.GetAccountClient()

	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}
	response, err := client.SettingsV2.PatchPublicAccountUserPreference(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update account_setting_user_preference_v2", err.Error())
		return
	}

	var newState UserPreference

	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *UserPreferenceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan UserPreference
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *UserPreferenceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

var _ resource.ResourceWithImportState = &UserPreferenceResource{}

func (r *UserPreferenceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, ",")

	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf(
				"Expected import identifier with format: name,user_id. Got: %q",
				req.ID,
			),
		)
		return
	}

	name := parts[0]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), name)...)
	userId := parts[1]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("user_id"), userId)...)
}
