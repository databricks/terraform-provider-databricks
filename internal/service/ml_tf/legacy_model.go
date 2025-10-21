// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package ml_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// For activities, this contains the activity recorded for the action. For
// comments, this contains the comment details. For transition requests, this
// contains the transition request details.
type Activity_SdkV2 struct {
	ActivityType types.String `tfsdk:"activity_type"`
	// User-provided comment associated with the activity, comment, or
	// transition request.
	Comment types.String `tfsdk:"comment"`
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp"`
	// Source stage of the transition (if the activity is stage transition
	// related). Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	FromStage types.String `tfsdk:"from_stage"`
	// Unique identifier for the object.
	Id types.String `tfsdk:"id"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp"`
	// Comment made by system, for example explaining an activity of type
	// `SYSTEM_TRANSITION`. It usually describes a side effect, such as a
	// version being archived as part of another version's stage transition, and
	// may not be returned for some activity types.
	SystemComment types.String `tfsdk:"system_comment"`
	// Target stage of the transition (if the activity is stage transition
	// related). Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	ToStage types.String `tfsdk:"to_stage"`
	// The username of the user that created the object.
	UserId types.String `tfsdk:"user_id"`
}

func (to *Activity_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Activity_SdkV2) {
}

func (to *Activity_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Activity_SdkV2) {
}

func (m Activity_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["activity_type"] = attrs["activity_type"].SetOptional()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["creation_timestamp"] = attrs["creation_timestamp"].SetOptional()
	attrs["from_stage"] = attrs["from_stage"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["last_updated_timestamp"] = attrs["last_updated_timestamp"].SetOptional()
	attrs["system_comment"] = attrs["system_comment"].SetOptional()
	attrs["to_stage"] = attrs["to_stage"].SetOptional()
	attrs["user_id"] = attrs["user_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Activity.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Activity_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Activity_SdkV2
// only implements ToObjectValue() and Type().
func (m Activity_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"activity_type":          m.ActivityType,
			"comment":                m.Comment,
			"creation_timestamp":     m.CreationTimestamp,
			"from_stage":             m.FromStage,
			"id":                     m.Id,
			"last_updated_timestamp": m.LastUpdatedTimestamp,
			"system_comment":         m.SystemComment,
			"to_stage":               m.ToStage,
			"user_id":                m.UserId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Activity_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"activity_type":          types.StringType,
			"comment":                types.StringType,
			"creation_timestamp":     types.Int64Type,
			"from_stage":             types.StringType,
			"id":                     types.StringType,
			"last_updated_timestamp": types.Int64Type,
			"system_comment":         types.StringType,
			"to_stage":               types.StringType,
			"user_id":                types.StringType,
		},
	}
}

// Details required to identify and approve a model version stage transition
// request.
type ApproveTransitionRequest_SdkV2 struct {
	// Specifies whether to archive all current model versions in the target
	// stage.
	ArchiveExistingVersions types.Bool `tfsdk:"archive_existing_versions"`
	// User-provided comment on the action.
	Comment types.String `tfsdk:"comment"`
	// Name of the model.
	Name types.String `tfsdk:"name"`
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage types.String `tfsdk:"stage"`
	// Version of the model.
	Version types.String `tfsdk:"version"`
}

func (to *ApproveTransitionRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ApproveTransitionRequest_SdkV2) {
}

func (to *ApproveTransitionRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ApproveTransitionRequest_SdkV2) {
}

func (m ApproveTransitionRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["archive_existing_versions"] = attrs["archive_existing_versions"].SetRequired()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["stage"] = attrs["stage"].SetRequired()
	attrs["version"] = attrs["version"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ApproveTransitionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ApproveTransitionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ApproveTransitionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ApproveTransitionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"archive_existing_versions": m.ArchiveExistingVersions,
			"comment":                   m.Comment,
			"name":                      m.Name,
			"stage":                     m.Stage,
			"version":                   m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ApproveTransitionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"archive_existing_versions": types.BoolType,
			"comment":                   types.StringType,
			"name":                      types.StringType,
			"stage":                     types.StringType,
			"version":                   types.StringType,
		},
	}
}

type ApproveTransitionRequestResponse_SdkV2 struct {
	// New activity generated as a result of this operation.
	Activity types.List `tfsdk:"activity"`
}

func (to *ApproveTransitionRequestResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ApproveTransitionRequestResponse_SdkV2) {
	if !from.Activity.IsNull() && !from.Activity.IsUnknown() {
		if toActivity, ok := to.GetActivity(ctx); ok {
			if fromActivity, ok := from.GetActivity(ctx); ok {
				// Recursively sync the fields of Activity
				toActivity.SyncFieldsDuringCreateOrUpdate(ctx, fromActivity)
				to.SetActivity(ctx, toActivity)
			}
		}
	}
}

func (to *ApproveTransitionRequestResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ApproveTransitionRequestResponse_SdkV2) {
	if !from.Activity.IsNull() && !from.Activity.IsUnknown() {
		if toActivity, ok := to.GetActivity(ctx); ok {
			if fromActivity, ok := from.GetActivity(ctx); ok {
				toActivity.SyncFieldsDuringRead(ctx, fromActivity)
				to.SetActivity(ctx, toActivity)
			}
		}
	}
}

func (m ApproveTransitionRequestResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["activity"] = attrs["activity"].SetOptional()
	attrs["activity"] = attrs["activity"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ApproveTransitionRequestResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ApproveTransitionRequestResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"activity": reflect.TypeOf(Activity_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ApproveTransitionRequestResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ApproveTransitionRequestResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"activity": m.Activity,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ApproveTransitionRequestResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"activity": basetypes.ListType{
				ElemType: Activity_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetActivity returns the value of the Activity field in ApproveTransitionRequestResponse_SdkV2 as
// a Activity_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ApproveTransitionRequestResponse_SdkV2) GetActivity(ctx context.Context) (Activity_SdkV2, bool) {
	var e Activity_SdkV2
	if m.Activity.IsNull() || m.Activity.IsUnknown() {
		return e, false
	}
	var v []Activity_SdkV2
	d := m.Activity.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetActivity sets the value of the Activity field in ApproveTransitionRequestResponse_SdkV2.
func (m *ApproveTransitionRequestResponse_SdkV2) SetActivity(ctx context.Context, v Activity_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["activity"]
	m.Activity = types.ListValueMust(t, vs)
}

// For activities, this contains the activity recorded for the action. For
// comments, this contains the comment details. For transition requests, this
// contains the transition request details.
type CommentObject_SdkV2 struct {
	// Array of actions on the activity allowed for the current viewer.
	AvailableActions types.List `tfsdk:"available_actions"`
	// User-provided comment associated with the activity, comment, or
	// transition request.
	Comment types.String `tfsdk:"comment"`
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp"`
	// Unique identifier for the object.
	Id types.String `tfsdk:"id"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp"`
	// The username of the user that created the object.
	UserId types.String `tfsdk:"user_id"`
}

func (to *CommentObject_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CommentObject_SdkV2) {
	if !from.AvailableActions.IsNull() && !from.AvailableActions.IsUnknown() && to.AvailableActions.IsNull() && len(from.AvailableActions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AvailableActions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AvailableActions = from.AvailableActions
	}
}

func (to *CommentObject_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CommentObject_SdkV2) {
	if !from.AvailableActions.IsNull() && !from.AvailableActions.IsUnknown() && to.AvailableActions.IsNull() && len(from.AvailableActions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AvailableActions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AvailableActions = from.AvailableActions
	}
}

func (m CommentObject_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["available_actions"] = attrs["available_actions"].SetOptional()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["creation_timestamp"] = attrs["creation_timestamp"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["last_updated_timestamp"] = attrs["last_updated_timestamp"].SetOptional()
	attrs["user_id"] = attrs["user_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CommentObject.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CommentObject_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"available_actions": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CommentObject_SdkV2
// only implements ToObjectValue() and Type().
func (m CommentObject_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"available_actions":      m.AvailableActions,
			"comment":                m.Comment,
			"creation_timestamp":     m.CreationTimestamp,
			"id":                     m.Id,
			"last_updated_timestamp": m.LastUpdatedTimestamp,
			"user_id":                m.UserId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CommentObject_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"available_actions": basetypes.ListType{
				ElemType: types.StringType,
			},
			"comment":                types.StringType,
			"creation_timestamp":     types.Int64Type,
			"id":                     types.StringType,
			"last_updated_timestamp": types.Int64Type,
			"user_id":                types.StringType,
		},
	}
}

// GetAvailableActions returns the value of the AvailableActions field in CommentObject_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CommentObject_SdkV2) GetAvailableActions(ctx context.Context) ([]types.String, bool) {
	if m.AvailableActions.IsNull() || m.AvailableActions.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.AvailableActions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAvailableActions sets the value of the AvailableActions field in CommentObject_SdkV2.
func (m *CommentObject_SdkV2) SetAvailableActions(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["available_actions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AvailableActions = types.ListValueMust(t, vs)
}

// Details required to create a comment on a model version.
type CreateComment_SdkV2 struct {
	// User-provided comment on the action.
	Comment types.String `tfsdk:"comment"`
	// Name of the model.
	Name types.String `tfsdk:"name"`
	// Version of the model.
	Version types.String `tfsdk:"version"`
}

func (to *CreateComment_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateComment_SdkV2) {
}

func (to *CreateComment_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateComment_SdkV2) {
}

func (m CreateComment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["version"] = attrs["version"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateComment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateComment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateComment_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateComment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment": m.Comment,
			"name":    m.Name,
			"version": m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateComment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment": types.StringType,
			"name":    types.StringType,
			"version": types.StringType,
		},
	}
}

type CreateCommentResponse_SdkV2 struct {
	// New comment object
	Comment types.List `tfsdk:"comment"`
}

func (to *CreateCommentResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCommentResponse_SdkV2) {
	if !from.Comment.IsNull() && !from.Comment.IsUnknown() {
		if toComment, ok := to.GetComment(ctx); ok {
			if fromComment, ok := from.GetComment(ctx); ok {
				// Recursively sync the fields of Comment
				toComment.SyncFieldsDuringCreateOrUpdate(ctx, fromComment)
				to.SetComment(ctx, toComment)
			}
		}
	}
}

func (to *CreateCommentResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateCommentResponse_SdkV2) {
	if !from.Comment.IsNull() && !from.Comment.IsUnknown() {
		if toComment, ok := to.GetComment(ctx); ok {
			if fromComment, ok := from.GetComment(ctx); ok {
				toComment.SyncFieldsDuringRead(ctx, fromComment)
				to.SetComment(ctx, toComment)
			}
		}
	}
}

func (m CreateCommentResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["comment"] = attrs["comment"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCommentResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateCommentResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"comment": reflect.TypeOf(CommentObject_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCommentResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateCommentResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment": m.Comment,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateCommentResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment": basetypes.ListType{
				ElemType: CommentObject_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetComment returns the value of the Comment field in CreateCommentResponse_SdkV2 as
// a CommentObject_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateCommentResponse_SdkV2) GetComment(ctx context.Context) (CommentObject_SdkV2, bool) {
	var e CommentObject_SdkV2
	if m.Comment.IsNull() || m.Comment.IsUnknown() {
		return e, false
	}
	var v []CommentObject_SdkV2
	d := m.Comment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetComment sets the value of the Comment field in CreateCommentResponse_SdkV2.
func (m *CreateCommentResponse_SdkV2) SetComment(ctx context.Context, v CommentObject_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["comment"]
	m.Comment = types.ListValueMust(t, vs)
}

type CreateExperiment_SdkV2 struct {
	// Location where all artifacts for the experiment are stored. If not
	// provided, the remote server will select an appropriate default.
	ArtifactLocation types.String `tfsdk:"artifact_location"`
	// Experiment name.
	Name types.String `tfsdk:"name"`
	// A collection of tags to set on the experiment. Maximum tag size and
	// number of tags per request depends on the storage backend. All storage
	// backends are guaranteed to support tag keys up to 250 bytes in size and
	// tag values up to 5000 bytes in size. All storage backends are also
	// guaranteed to support up to 20 tags per request.
	Tags types.List `tfsdk:"tags"`
}

func (to *CreateExperiment_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateExperiment_SdkV2) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *CreateExperiment_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateExperiment_SdkV2) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m CreateExperiment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["artifact_location"] = attrs["artifact_location"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["tags"] = attrs["tags"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateExperiment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateExperiment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(ExperimentTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateExperiment_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateExperiment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"artifact_location": m.ArtifactLocation,
			"name":              m.Name,
			"tags":              m.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateExperiment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"artifact_location": types.StringType,
			"name":              types.StringType,
			"tags": basetypes.ListType{
				ElemType: ExperimentTag_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetTags returns the value of the Tags field in CreateExperiment_SdkV2 as
// a slice of ExperimentTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateExperiment_SdkV2) GetTags(ctx context.Context) ([]ExperimentTag_SdkV2, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []ExperimentTag_SdkV2
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in CreateExperiment_SdkV2.
func (m *CreateExperiment_SdkV2) SetTags(ctx context.Context, v []ExperimentTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type CreateExperimentResponse_SdkV2 struct {
	// Unique identifier for the experiment.
	ExperimentId types.String `tfsdk:"experiment_id"`
}

func (to *CreateExperimentResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateExperimentResponse_SdkV2) {
}

func (to *CreateExperimentResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateExperimentResponse_SdkV2) {
}

func (m CreateExperimentResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["experiment_id"] = attrs["experiment_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateExperimentResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateExperimentResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateExperimentResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateExperimentResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": m.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateExperimentResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
		},
	}
}

type CreateFeatureRequest_SdkV2 struct {
	// Feature to create.
	Feature types.List `tfsdk:"feature"`
}

func (to *CreateFeatureRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateFeatureRequest_SdkV2) {
	if !from.Feature.IsNull() && !from.Feature.IsUnknown() {
		if toFeature, ok := to.GetFeature(ctx); ok {
			if fromFeature, ok := from.GetFeature(ctx); ok {
				// Recursively sync the fields of Feature
				toFeature.SyncFieldsDuringCreateOrUpdate(ctx, fromFeature)
				to.SetFeature(ctx, toFeature)
			}
		}
	}
}

func (to *CreateFeatureRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateFeatureRequest_SdkV2) {
	if !from.Feature.IsNull() && !from.Feature.IsUnknown() {
		if toFeature, ok := to.GetFeature(ctx); ok {
			if fromFeature, ok := from.GetFeature(ctx); ok {
				toFeature.SyncFieldsDuringRead(ctx, fromFeature)
				to.SetFeature(ctx, toFeature)
			}
		}
	}
}

func (m CreateFeatureRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["feature"] = attrs["feature"].SetRequired()
	attrs["feature"] = attrs["feature"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateFeatureRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateFeatureRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"feature": reflect.TypeOf(Feature_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateFeatureRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateFeatureRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature": m.Feature,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateFeatureRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"feature": basetypes.ListType{
				ElemType: Feature_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetFeature returns the value of the Feature field in CreateFeatureRequest_SdkV2 as
// a Feature_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateFeatureRequest_SdkV2) GetFeature(ctx context.Context) (Feature_SdkV2, bool) {
	var e Feature_SdkV2
	if m.Feature.IsNull() || m.Feature.IsUnknown() {
		return e, false
	}
	var v []Feature_SdkV2
	d := m.Feature.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFeature sets the value of the Feature field in CreateFeatureRequest_SdkV2.
func (m *CreateFeatureRequest_SdkV2) SetFeature(ctx context.Context, v Feature_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["feature"]
	m.Feature = types.ListValueMust(t, vs)
}

type CreateFeatureTagRequest_SdkV2 struct {
	FeatureName types.String `tfsdk:"-"`

	FeatureTag types.List `tfsdk:"feature_tag"`

	TableName types.String `tfsdk:"-"`
}

func (to *CreateFeatureTagRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateFeatureTagRequest_SdkV2) {
	if !from.FeatureTag.IsNull() && !from.FeatureTag.IsUnknown() {
		if toFeatureTag, ok := to.GetFeatureTag(ctx); ok {
			if fromFeatureTag, ok := from.GetFeatureTag(ctx); ok {
				// Recursively sync the fields of FeatureTag
				toFeatureTag.SyncFieldsDuringCreateOrUpdate(ctx, fromFeatureTag)
				to.SetFeatureTag(ctx, toFeatureTag)
			}
		}
	}
}

func (to *CreateFeatureTagRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateFeatureTagRequest_SdkV2) {
	if !from.FeatureTag.IsNull() && !from.FeatureTag.IsUnknown() {
		if toFeatureTag, ok := to.GetFeatureTag(ctx); ok {
			if fromFeatureTag, ok := from.GetFeatureTag(ctx); ok {
				toFeatureTag.SyncFieldsDuringRead(ctx, fromFeatureTag)
				to.SetFeatureTag(ctx, toFeatureTag)
			}
		}
	}
}

func (m CreateFeatureTagRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["feature_tag"] = attrs["feature_tag"].SetRequired()
	attrs["feature_tag"] = attrs["feature_tag"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["table_name"] = attrs["table_name"].SetRequired()
	attrs["feature_name"] = attrs["feature_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateFeatureTagRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateFeatureTagRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"feature_tag": reflect.TypeOf(FeatureTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateFeatureTagRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateFeatureTagRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_name": m.FeatureName,
			"feature_tag":  m.FeatureTag,
			"table_name":   m.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateFeatureTagRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"feature_name": types.StringType,
			"feature_tag": basetypes.ListType{
				ElemType: FeatureTag_SdkV2{}.Type(ctx),
			},
			"table_name": types.StringType,
		},
	}
}

// GetFeatureTag returns the value of the FeatureTag field in CreateFeatureTagRequest_SdkV2 as
// a FeatureTag_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateFeatureTagRequest_SdkV2) GetFeatureTag(ctx context.Context) (FeatureTag_SdkV2, bool) {
	var e FeatureTag_SdkV2
	if m.FeatureTag.IsNull() || m.FeatureTag.IsUnknown() {
		return e, false
	}
	var v []FeatureTag_SdkV2
	d := m.FeatureTag.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFeatureTag sets the value of the FeatureTag field in CreateFeatureTagRequest_SdkV2.
func (m *CreateFeatureTagRequest_SdkV2) SetFeatureTag(ctx context.Context, v FeatureTag_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["feature_tag"]
	m.FeatureTag = types.ListValueMust(t, vs)
}

type CreateForecastingExperimentRequest_SdkV2 struct {
	// The column in the training table used to customize weights for each time
	// series.
	CustomWeightsColumn types.String `tfsdk:"custom_weights_column"`
	// The path in the workspace to store the created experiment.
	ExperimentPath types.String `tfsdk:"experiment_path"`
	// The time interval between consecutive rows in the time series data.
	// Possible values include: '1 second', '1 minute', '5 minutes', '10
	// minutes', '15 minutes', '30 minutes', 'Hourly', 'Daily', 'Weekly',
	// 'Monthly', 'Quarterly', 'Yearly'.
	ForecastGranularity types.String `tfsdk:"forecast_granularity"`
	// The number of time steps into the future to make predictions, calculated
	// as a multiple of forecast_granularity. This value represents how far
	// ahead the model should forecast.
	ForecastHorizon types.Int64 `tfsdk:"forecast_horizon"`
	// The fully qualified path of a Unity Catalog table, formatted as
	// catalog_name.schema_name.table_name, used to store future feature data
	// for predictions.
	FutureFeatureDataPath types.String `tfsdk:"future_feature_data_path"`
	// The region code(s) to automatically add holiday features. Currently
	// supports only one region.
	HolidayRegions types.List `tfsdk:"holiday_regions"`
	// Specifies the list of feature columns to include in model training. These
	// columns must exist in the training data and be of type string, numerical,
	// or boolean. If not specified, no additional features will be included.
	// Note: Certain columns are automatically handled: - Automatically
	// excluded: split_column, target_column, custom_weights_column. -
	// Automatically included: time_column.
	IncludeFeatures types.List `tfsdk:"include_features"`
	// The maximum duration for the experiment in minutes. The experiment stops
	// automatically if it exceeds this limit.
	MaxRuntime types.Int64 `tfsdk:"max_runtime"`
	// The fully qualified path of a Unity Catalog table, formatted as
	// catalog_name.schema_name.table_name, used to store predictions.
	PredictionDataPath types.String `tfsdk:"prediction_data_path"`
	// The evaluation metric used to optimize the forecasting model.
	PrimaryMetric types.String `tfsdk:"primary_metric"`
	// The fully qualified path of a Unity Catalog model, formatted as
	// catalog_name.schema_name.model_name, used to store the best model.
	RegisterTo types.String `tfsdk:"register_to"`
	// // The column in the training table used for custom data splits. Values
	// must be 'train', 'validate', or 'test'.
	SplitColumn types.String `tfsdk:"split_column"`
	// The column in the input training table used as the prediction target for
	// model training. The values in this column are used as the ground truth
	// for model training.
	TargetColumn types.String `tfsdk:"target_column"`
	// The column in the input training table that represents each row's
	// timestamp.
	TimeColumn types.String `tfsdk:"time_column"`
	// The column in the training table used to group the dataset for predicting
	// individual time series.
	TimeseriesIdentifierColumns types.List `tfsdk:"timeseries_identifier_columns"`
	// The fully qualified path of a Unity Catalog table, formatted as
	// catalog_name.schema_name.table_name, used as training data for the
	// forecasting model.
	TrainDataPath types.String `tfsdk:"train_data_path"`
	// List of frameworks to include for model tuning. Possible values are
	// 'Prophet', 'ARIMA', 'DeepAR'. An empty list includes all supported
	// frameworks.
	TrainingFrameworks types.List `tfsdk:"training_frameworks"`
}

func (to *CreateForecastingExperimentRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateForecastingExperimentRequest_SdkV2) {
	if !from.HolidayRegions.IsNull() && !from.HolidayRegions.IsUnknown() && to.HolidayRegions.IsNull() && len(from.HolidayRegions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for HolidayRegions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.HolidayRegions = from.HolidayRegions
	}
	if !from.IncludeFeatures.IsNull() && !from.IncludeFeatures.IsUnknown() && to.IncludeFeatures.IsNull() && len(from.IncludeFeatures.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for IncludeFeatures, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.IncludeFeatures = from.IncludeFeatures
	}
	if !from.TimeseriesIdentifierColumns.IsNull() && !from.TimeseriesIdentifierColumns.IsUnknown() && to.TimeseriesIdentifierColumns.IsNull() && len(from.TimeseriesIdentifierColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for TimeseriesIdentifierColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.TimeseriesIdentifierColumns = from.TimeseriesIdentifierColumns
	}
	if !from.TrainingFrameworks.IsNull() && !from.TrainingFrameworks.IsUnknown() && to.TrainingFrameworks.IsNull() && len(from.TrainingFrameworks.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for TrainingFrameworks, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.TrainingFrameworks = from.TrainingFrameworks
	}
}

func (to *CreateForecastingExperimentRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateForecastingExperimentRequest_SdkV2) {
	if !from.HolidayRegions.IsNull() && !from.HolidayRegions.IsUnknown() && to.HolidayRegions.IsNull() && len(from.HolidayRegions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for HolidayRegions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.HolidayRegions = from.HolidayRegions
	}
	if !from.IncludeFeatures.IsNull() && !from.IncludeFeatures.IsUnknown() && to.IncludeFeatures.IsNull() && len(from.IncludeFeatures.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for IncludeFeatures, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.IncludeFeatures = from.IncludeFeatures
	}
	if !from.TimeseriesIdentifierColumns.IsNull() && !from.TimeseriesIdentifierColumns.IsUnknown() && to.TimeseriesIdentifierColumns.IsNull() && len(from.TimeseriesIdentifierColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for TimeseriesIdentifierColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.TimeseriesIdentifierColumns = from.TimeseriesIdentifierColumns
	}
	if !from.TrainingFrameworks.IsNull() && !from.TrainingFrameworks.IsUnknown() && to.TrainingFrameworks.IsNull() && len(from.TrainingFrameworks.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for TrainingFrameworks, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.TrainingFrameworks = from.TrainingFrameworks
	}
}

func (m CreateForecastingExperimentRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["custom_weights_column"] = attrs["custom_weights_column"].SetOptional()
	attrs["experiment_path"] = attrs["experiment_path"].SetOptional()
	attrs["forecast_granularity"] = attrs["forecast_granularity"].SetRequired()
	attrs["forecast_horizon"] = attrs["forecast_horizon"].SetRequired()
	attrs["future_feature_data_path"] = attrs["future_feature_data_path"].SetOptional()
	attrs["holiday_regions"] = attrs["holiday_regions"].SetOptional()
	attrs["include_features"] = attrs["include_features"].SetOptional()
	attrs["max_runtime"] = attrs["max_runtime"].SetOptional()
	attrs["prediction_data_path"] = attrs["prediction_data_path"].SetOptional()
	attrs["primary_metric"] = attrs["primary_metric"].SetOptional()
	attrs["register_to"] = attrs["register_to"].SetOptional()
	attrs["split_column"] = attrs["split_column"].SetOptional()
	attrs["target_column"] = attrs["target_column"].SetRequired()
	attrs["time_column"] = attrs["time_column"].SetRequired()
	attrs["timeseries_identifier_columns"] = attrs["timeseries_identifier_columns"].SetOptional()
	attrs["train_data_path"] = attrs["train_data_path"].SetRequired()
	attrs["training_frameworks"] = attrs["training_frameworks"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateForecastingExperimentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateForecastingExperimentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"holiday_regions":               reflect.TypeOf(types.String{}),
		"include_features":              reflect.TypeOf(types.String{}),
		"timeseries_identifier_columns": reflect.TypeOf(types.String{}),
		"training_frameworks":           reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateForecastingExperimentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateForecastingExperimentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"custom_weights_column":         m.CustomWeightsColumn,
			"experiment_path":               m.ExperimentPath,
			"forecast_granularity":          m.ForecastGranularity,
			"forecast_horizon":              m.ForecastHorizon,
			"future_feature_data_path":      m.FutureFeatureDataPath,
			"holiday_regions":               m.HolidayRegions,
			"include_features":              m.IncludeFeatures,
			"max_runtime":                   m.MaxRuntime,
			"prediction_data_path":          m.PredictionDataPath,
			"primary_metric":                m.PrimaryMetric,
			"register_to":                   m.RegisterTo,
			"split_column":                  m.SplitColumn,
			"target_column":                 m.TargetColumn,
			"time_column":                   m.TimeColumn,
			"timeseries_identifier_columns": m.TimeseriesIdentifierColumns,
			"train_data_path":               m.TrainDataPath,
			"training_frameworks":           m.TrainingFrameworks,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateForecastingExperimentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"custom_weights_column":    types.StringType,
			"experiment_path":          types.StringType,
			"forecast_granularity":     types.StringType,
			"forecast_horizon":         types.Int64Type,
			"future_feature_data_path": types.StringType,
			"holiday_regions": basetypes.ListType{
				ElemType: types.StringType,
			},
			"include_features": basetypes.ListType{
				ElemType: types.StringType,
			},
			"max_runtime":          types.Int64Type,
			"prediction_data_path": types.StringType,
			"primary_metric":       types.StringType,
			"register_to":          types.StringType,
			"split_column":         types.StringType,
			"target_column":        types.StringType,
			"time_column":          types.StringType,
			"timeseries_identifier_columns": basetypes.ListType{
				ElemType: types.StringType,
			},
			"train_data_path": types.StringType,
			"training_frameworks": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetHolidayRegions returns the value of the HolidayRegions field in CreateForecastingExperimentRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateForecastingExperimentRequest_SdkV2) GetHolidayRegions(ctx context.Context) ([]types.String, bool) {
	if m.HolidayRegions.IsNull() || m.HolidayRegions.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.HolidayRegions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetHolidayRegions sets the value of the HolidayRegions field in CreateForecastingExperimentRequest_SdkV2.
func (m *CreateForecastingExperimentRequest_SdkV2) SetHolidayRegions(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["holiday_regions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.HolidayRegions = types.ListValueMust(t, vs)
}

// GetIncludeFeatures returns the value of the IncludeFeatures field in CreateForecastingExperimentRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateForecastingExperimentRequest_SdkV2) GetIncludeFeatures(ctx context.Context) ([]types.String, bool) {
	if m.IncludeFeatures.IsNull() || m.IncludeFeatures.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.IncludeFeatures.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIncludeFeatures sets the value of the IncludeFeatures field in CreateForecastingExperimentRequest_SdkV2.
func (m *CreateForecastingExperimentRequest_SdkV2) SetIncludeFeatures(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["include_features"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.IncludeFeatures = types.ListValueMust(t, vs)
}

// GetTimeseriesIdentifierColumns returns the value of the TimeseriesIdentifierColumns field in CreateForecastingExperimentRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateForecastingExperimentRequest_SdkV2) GetTimeseriesIdentifierColumns(ctx context.Context) ([]types.String, bool) {
	if m.TimeseriesIdentifierColumns.IsNull() || m.TimeseriesIdentifierColumns.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.TimeseriesIdentifierColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTimeseriesIdentifierColumns sets the value of the TimeseriesIdentifierColumns field in CreateForecastingExperimentRequest_SdkV2.
func (m *CreateForecastingExperimentRequest_SdkV2) SetTimeseriesIdentifierColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["timeseries_identifier_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.TimeseriesIdentifierColumns = types.ListValueMust(t, vs)
}

// GetTrainingFrameworks returns the value of the TrainingFrameworks field in CreateForecastingExperimentRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateForecastingExperimentRequest_SdkV2) GetTrainingFrameworks(ctx context.Context) ([]types.String, bool) {
	if m.TrainingFrameworks.IsNull() || m.TrainingFrameworks.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.TrainingFrameworks.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTrainingFrameworks sets the value of the TrainingFrameworks field in CreateForecastingExperimentRequest_SdkV2.
func (m *CreateForecastingExperimentRequest_SdkV2) SetTrainingFrameworks(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["training_frameworks"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.TrainingFrameworks = types.ListValueMust(t, vs)
}

type CreateForecastingExperimentResponse_SdkV2 struct {
	// The unique ID of the created forecasting experiment
	ExperimentId types.String `tfsdk:"experiment_id"`
}

func (to *CreateForecastingExperimentResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateForecastingExperimentResponse_SdkV2) {
}

func (to *CreateForecastingExperimentResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateForecastingExperimentResponse_SdkV2) {
}

func (m CreateForecastingExperimentResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["experiment_id"] = attrs["experiment_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateForecastingExperimentResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateForecastingExperimentResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateForecastingExperimentResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateForecastingExperimentResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": m.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateForecastingExperimentResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
		},
	}
}

type CreateLoggedModelRequest_SdkV2 struct {
	// The ID of the experiment that owns the model.
	ExperimentId types.String `tfsdk:"experiment_id"`
	// The type of the model, such as ``"Agent"``, ``"Classifier"``, ``"LLM"``.
	ModelType types.String `tfsdk:"model_type"`
	// The name of the model (optional). If not specified one will be generated.
	Name types.String `tfsdk:"name"`
	// Parameters attached to the model.
	Params types.List `tfsdk:"params"`
	// The ID of the run that created the model.
	SourceRunId types.String `tfsdk:"source_run_id"`
	// Tags attached to the model.
	Tags types.List `tfsdk:"tags"`
}

func (to *CreateLoggedModelRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateLoggedModelRequest_SdkV2) {
	if !from.Params.IsNull() && !from.Params.IsUnknown() && to.Params.IsNull() && len(from.Params.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Params, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Params = from.Params
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *CreateLoggedModelRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateLoggedModelRequest_SdkV2) {
	if !from.Params.IsNull() && !from.Params.IsUnknown() && to.Params.IsNull() && len(from.Params.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Params, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Params = from.Params
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m CreateLoggedModelRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["experiment_id"] = attrs["experiment_id"].SetRequired()
	attrs["model_type"] = attrs["model_type"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["params"] = attrs["params"].SetOptional()
	attrs["source_run_id"] = attrs["source_run_id"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateLoggedModelRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateLoggedModelRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"params": reflect.TypeOf(LoggedModelParameter_SdkV2{}),
		"tags":   reflect.TypeOf(LoggedModelTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateLoggedModelRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateLoggedModelRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": m.ExperimentId,
			"model_type":    m.ModelType,
			"name":          m.Name,
			"params":        m.Params,
			"source_run_id": m.SourceRunId,
			"tags":          m.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateLoggedModelRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
			"model_type":    types.StringType,
			"name":          types.StringType,
			"params": basetypes.ListType{
				ElemType: LoggedModelParameter_SdkV2{}.Type(ctx),
			},
			"source_run_id": types.StringType,
			"tags": basetypes.ListType{
				ElemType: LoggedModelTag_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetParams returns the value of the Params field in CreateLoggedModelRequest_SdkV2 as
// a slice of LoggedModelParameter_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateLoggedModelRequest_SdkV2) GetParams(ctx context.Context) ([]LoggedModelParameter_SdkV2, bool) {
	if m.Params.IsNull() || m.Params.IsUnknown() {
		return nil, false
	}
	var v []LoggedModelParameter_SdkV2
	d := m.Params.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParams sets the value of the Params field in CreateLoggedModelRequest_SdkV2.
func (m *CreateLoggedModelRequest_SdkV2) SetParams(ctx context.Context, v []LoggedModelParameter_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["params"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Params = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in CreateLoggedModelRequest_SdkV2 as
// a slice of LoggedModelTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateLoggedModelRequest_SdkV2) GetTags(ctx context.Context) ([]LoggedModelTag_SdkV2, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []LoggedModelTag_SdkV2
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in CreateLoggedModelRequest_SdkV2.
func (m *CreateLoggedModelRequest_SdkV2) SetTags(ctx context.Context, v []LoggedModelTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type CreateLoggedModelResponse_SdkV2 struct {
	// The newly created logged model.
	Model types.List `tfsdk:"model"`
}

func (to *CreateLoggedModelResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateLoggedModelResponse_SdkV2) {
	if !from.Model.IsNull() && !from.Model.IsUnknown() {
		if toModel, ok := to.GetModel(ctx); ok {
			if fromModel, ok := from.GetModel(ctx); ok {
				// Recursively sync the fields of Model
				toModel.SyncFieldsDuringCreateOrUpdate(ctx, fromModel)
				to.SetModel(ctx, toModel)
			}
		}
	}
}

func (to *CreateLoggedModelResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateLoggedModelResponse_SdkV2) {
	if !from.Model.IsNull() && !from.Model.IsUnknown() {
		if toModel, ok := to.GetModel(ctx); ok {
			if fromModel, ok := from.GetModel(ctx); ok {
				toModel.SyncFieldsDuringRead(ctx, fromModel)
				to.SetModel(ctx, toModel)
			}
		}
	}
}

func (m CreateLoggedModelResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["model"] = attrs["model"].SetOptional()
	attrs["model"] = attrs["model"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateLoggedModelResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateLoggedModelResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model": reflect.TypeOf(LoggedModel_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateLoggedModelResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateLoggedModelResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model": m.Model,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateLoggedModelResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model": basetypes.ListType{
				ElemType: LoggedModel_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetModel returns the value of the Model field in CreateLoggedModelResponse_SdkV2 as
// a LoggedModel_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateLoggedModelResponse_SdkV2) GetModel(ctx context.Context) (LoggedModel_SdkV2, bool) {
	var e LoggedModel_SdkV2
	if m.Model.IsNull() || m.Model.IsUnknown() {
		return e, false
	}
	var v []LoggedModel_SdkV2
	d := m.Model.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetModel sets the value of the Model field in CreateLoggedModelResponse_SdkV2.
func (m *CreateLoggedModelResponse_SdkV2) SetModel(ctx context.Context, v LoggedModel_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["model"]
	m.Model = types.ListValueMust(t, vs)
}

type CreateModelRequest_SdkV2 struct {
	// Optional description for registered model.
	Description types.String `tfsdk:"description"`
	// Register models under this name
	Name types.String `tfsdk:"name"`
	// Additional metadata for registered model.
	Tags types.List `tfsdk:"tags"`
}

func (to *CreateModelRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateModelRequest_SdkV2) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *CreateModelRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateModelRequest_SdkV2) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m CreateModelRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["tags"] = attrs["tags"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateModelRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateModelRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(ModelTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateModelRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateModelRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description": m.Description,
			"name":        m.Name,
			"tags":        m.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateModelRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"name":        types.StringType,
			"tags": basetypes.ListType{
				ElemType: ModelTag_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetTags returns the value of the Tags field in CreateModelRequest_SdkV2 as
// a slice of ModelTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateModelRequest_SdkV2) GetTags(ctx context.Context) ([]ModelTag_SdkV2, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []ModelTag_SdkV2
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in CreateModelRequest_SdkV2.
func (m *CreateModelRequest_SdkV2) SetTags(ctx context.Context, v []ModelTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type CreateModelResponse_SdkV2 struct {
	RegisteredModel types.List `tfsdk:"registered_model"`
}

func (to *CreateModelResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateModelResponse_SdkV2) {
	if !from.RegisteredModel.IsNull() && !from.RegisteredModel.IsUnknown() {
		if toRegisteredModel, ok := to.GetRegisteredModel(ctx); ok {
			if fromRegisteredModel, ok := from.GetRegisteredModel(ctx); ok {
				// Recursively sync the fields of RegisteredModel
				toRegisteredModel.SyncFieldsDuringCreateOrUpdate(ctx, fromRegisteredModel)
				to.SetRegisteredModel(ctx, toRegisteredModel)
			}
		}
	}
}

func (to *CreateModelResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateModelResponse_SdkV2) {
	if !from.RegisteredModel.IsNull() && !from.RegisteredModel.IsUnknown() {
		if toRegisteredModel, ok := to.GetRegisteredModel(ctx); ok {
			if fromRegisteredModel, ok := from.GetRegisteredModel(ctx); ok {
				toRegisteredModel.SyncFieldsDuringRead(ctx, fromRegisteredModel)
				to.SetRegisteredModel(ctx, toRegisteredModel)
			}
		}
	}
}

func (m CreateModelResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["registered_model"] = attrs["registered_model"].SetOptional()
	attrs["registered_model"] = attrs["registered_model"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateModelResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateModelResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"registered_model": reflect.TypeOf(Model_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateModelResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateModelResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"registered_model": m.RegisteredModel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateModelResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"registered_model": basetypes.ListType{
				ElemType: Model_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetRegisteredModel returns the value of the RegisteredModel field in CreateModelResponse_SdkV2 as
// a Model_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateModelResponse_SdkV2) GetRegisteredModel(ctx context.Context) (Model_SdkV2, bool) {
	var e Model_SdkV2
	if m.RegisteredModel.IsNull() || m.RegisteredModel.IsUnknown() {
		return e, false
	}
	var v []Model_SdkV2
	d := m.RegisteredModel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRegisteredModel sets the value of the RegisteredModel field in CreateModelResponse_SdkV2.
func (m *CreateModelResponse_SdkV2) SetRegisteredModel(ctx context.Context, v Model_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["registered_model"]
	m.RegisteredModel = types.ListValueMust(t, vs)
}

type CreateModelVersionRequest_SdkV2 struct {
	// Optional description for model version.
	Description types.String `tfsdk:"description"`
	// Register model under this name
	Name types.String `tfsdk:"name"`
	// MLflow run ID for correlation, if `source` was generated by an experiment
	// run in MLflow tracking server
	RunId types.String `tfsdk:"run_id"`
	// MLflow run link - this is the exact link of the run that generated this
	// model version, potentially hosted at another instance of MLflow.
	RunLink types.String `tfsdk:"run_link"`
	// URI indicating the location of the model artifacts.
	Source types.String `tfsdk:"source"`
	// Additional metadata for model version.
	Tags types.List `tfsdk:"tags"`
}

func (to *CreateModelVersionRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateModelVersionRequest_SdkV2) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *CreateModelVersionRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateModelVersionRequest_SdkV2) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m CreateModelVersionRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["run_id"] = attrs["run_id"].SetOptional()
	attrs["run_link"] = attrs["run_link"].SetOptional()
	attrs["source"] = attrs["source"].SetRequired()
	attrs["tags"] = attrs["tags"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateModelVersionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateModelVersionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(ModelVersionTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateModelVersionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateModelVersionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description": m.Description,
			"name":        m.Name,
			"run_id":      m.RunId,
			"run_link":    m.RunLink,
			"source":      m.Source,
			"tags":        m.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateModelVersionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"name":        types.StringType,
			"run_id":      types.StringType,
			"run_link":    types.StringType,
			"source":      types.StringType,
			"tags": basetypes.ListType{
				ElemType: ModelVersionTag_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetTags returns the value of the Tags field in CreateModelVersionRequest_SdkV2 as
// a slice of ModelVersionTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateModelVersionRequest_SdkV2) GetTags(ctx context.Context) ([]ModelVersionTag_SdkV2, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []ModelVersionTag_SdkV2
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in CreateModelVersionRequest_SdkV2.
func (m *CreateModelVersionRequest_SdkV2) SetTags(ctx context.Context, v []ModelVersionTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type CreateModelVersionResponse_SdkV2 struct {
	// Return new version number generated for this model in registry.
	ModelVersion types.List `tfsdk:"model_version"`
}

func (to *CreateModelVersionResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateModelVersionResponse_SdkV2) {
	if !from.ModelVersion.IsNull() && !from.ModelVersion.IsUnknown() {
		if toModelVersion, ok := to.GetModelVersion(ctx); ok {
			if fromModelVersion, ok := from.GetModelVersion(ctx); ok {
				// Recursively sync the fields of ModelVersion
				toModelVersion.SyncFieldsDuringCreateOrUpdate(ctx, fromModelVersion)
				to.SetModelVersion(ctx, toModelVersion)
			}
		}
	}
}

func (to *CreateModelVersionResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateModelVersionResponse_SdkV2) {
	if !from.ModelVersion.IsNull() && !from.ModelVersion.IsUnknown() {
		if toModelVersion, ok := to.GetModelVersion(ctx); ok {
			if fromModelVersion, ok := from.GetModelVersion(ctx); ok {
				toModelVersion.SyncFieldsDuringRead(ctx, fromModelVersion)
				to.SetModelVersion(ctx, toModelVersion)
			}
		}
	}
}

func (m CreateModelVersionResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["model_version"] = attrs["model_version"].SetOptional()
	attrs["model_version"] = attrs["model_version"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateModelVersionResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateModelVersionResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model_version": reflect.TypeOf(ModelVersion_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateModelVersionResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateModelVersionResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_version": m.ModelVersion,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateModelVersionResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_version": basetypes.ListType{
				ElemType: ModelVersion_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetModelVersion returns the value of the ModelVersion field in CreateModelVersionResponse_SdkV2 as
// a ModelVersion_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateModelVersionResponse_SdkV2) GetModelVersion(ctx context.Context) (ModelVersion_SdkV2, bool) {
	var e ModelVersion_SdkV2
	if m.ModelVersion.IsNull() || m.ModelVersion.IsUnknown() {
		return e, false
	}
	var v []ModelVersion_SdkV2
	d := m.ModelVersion.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetModelVersion sets the value of the ModelVersion field in CreateModelVersionResponse_SdkV2.
func (m *CreateModelVersionResponse_SdkV2) SetModelVersion(ctx context.Context, v ModelVersion_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["model_version"]
	m.ModelVersion = types.ListValueMust(t, vs)
}

type CreateOnlineStoreRequest_SdkV2 struct {
	// Online store to create.
	OnlineStore types.List `tfsdk:"online_store"`
}

func (to *CreateOnlineStoreRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateOnlineStoreRequest_SdkV2) {
	if !from.OnlineStore.IsNull() && !from.OnlineStore.IsUnknown() {
		if toOnlineStore, ok := to.GetOnlineStore(ctx); ok {
			if fromOnlineStore, ok := from.GetOnlineStore(ctx); ok {
				// Recursively sync the fields of OnlineStore
				toOnlineStore.SyncFieldsDuringCreateOrUpdate(ctx, fromOnlineStore)
				to.SetOnlineStore(ctx, toOnlineStore)
			}
		}
	}
}

func (to *CreateOnlineStoreRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateOnlineStoreRequest_SdkV2) {
	if !from.OnlineStore.IsNull() && !from.OnlineStore.IsUnknown() {
		if toOnlineStore, ok := to.GetOnlineStore(ctx); ok {
			if fromOnlineStore, ok := from.GetOnlineStore(ctx); ok {
				toOnlineStore.SyncFieldsDuringRead(ctx, fromOnlineStore)
				to.SetOnlineStore(ctx, toOnlineStore)
			}
		}
	}
}

func (m CreateOnlineStoreRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["online_store"] = attrs["online_store"].SetRequired()
	attrs["online_store"] = attrs["online_store"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateOnlineStoreRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateOnlineStoreRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"online_store": reflect.TypeOf(OnlineStore_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateOnlineStoreRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateOnlineStoreRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"online_store": m.OnlineStore,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateOnlineStoreRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"online_store": basetypes.ListType{
				ElemType: OnlineStore_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetOnlineStore returns the value of the OnlineStore field in CreateOnlineStoreRequest_SdkV2 as
// a OnlineStore_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateOnlineStoreRequest_SdkV2) GetOnlineStore(ctx context.Context) (OnlineStore_SdkV2, bool) {
	var e OnlineStore_SdkV2
	if m.OnlineStore.IsNull() || m.OnlineStore.IsUnknown() {
		return e, false
	}
	var v []OnlineStore_SdkV2
	d := m.OnlineStore.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOnlineStore sets the value of the OnlineStore field in CreateOnlineStoreRequest_SdkV2.
func (m *CreateOnlineStoreRequest_SdkV2) SetOnlineStore(ctx context.Context, v OnlineStore_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["online_store"]
	m.OnlineStore = types.ListValueMust(t, vs)
}

// Details required to create a registry webhook.
type CreateRegistryWebhook_SdkV2 struct {
	// User-specified description for the webhook.
	Description types.String `tfsdk:"description"`
	// Events that can trigger a registry webhook: * `MODEL_VERSION_CREATED`: A
	// new model version was created for the associated model.
	//
	// * `MODEL_VERSION_TRANSITIONED_STAGE`: A model version’s stage was
	// changed.
	//
	// * `TRANSITION_REQUEST_CREATED`: A user requested a model version’s
	// stage be transitioned.
	//
	// * `COMMENT_CREATED`: A user wrote a comment on a registered model.
	//
	// * `REGISTERED_MODEL_CREATED`: A new registered model was created. This
	// event type can only be specified for a registry-wide webhook, which can
	// be created by not specifying a model name in the create request.
	//
	// * `MODEL_VERSION_TAG_SET`: A user set a tag on the model version.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_STAGING`: A model version was
	// transitioned to staging.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_PRODUCTION`: A model version was
	// transitioned to production.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_ARCHIVED`: A model version was archived.
	//
	// * `TRANSITION_REQUEST_TO_STAGING_CREATED`: A user requested a model
	// version be transitioned to staging.
	//
	// * `TRANSITION_REQUEST_TO_PRODUCTION_CREATED`: A user requested a model
	// version be transitioned to production.
	//
	// * `TRANSITION_REQUEST_TO_ARCHIVED_CREATED`: A user requested a model
	// version be archived.
	Events types.List `tfsdk:"events"`
	// External HTTPS URL called on event trigger (by using a POST request).
	HttpUrlSpec types.List `tfsdk:"http_url_spec"`
	// ID of the job that the webhook runs.
	JobSpec types.List `tfsdk:"job_spec"`
	// If model name is not specified, a registry-wide webhook is created that
	// listens for the specified events across all versions of all registered
	// models.
	ModelName types.String `tfsdk:"model_name"`
	// Enable or disable triggering the webhook, or put the webhook into test
	// mode. The default is `ACTIVE`: * `ACTIVE`: Webhook is triggered when an
	// associated event happens.
	//
	// * `DISABLED`: Webhook is not triggered.
	//
	// * `TEST_MODE`: Webhook can be triggered through the test endpoint, but is
	// not triggered on a real event.
	Status types.String `tfsdk:"status"`
}

func (to *CreateRegistryWebhook_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateRegistryWebhook_SdkV2) {
	if !from.HttpUrlSpec.IsNull() && !from.HttpUrlSpec.IsUnknown() {
		if toHttpUrlSpec, ok := to.GetHttpUrlSpec(ctx); ok {
			if fromHttpUrlSpec, ok := from.GetHttpUrlSpec(ctx); ok {
				// Recursively sync the fields of HttpUrlSpec
				toHttpUrlSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromHttpUrlSpec)
				to.SetHttpUrlSpec(ctx, toHttpUrlSpec)
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
}

func (to *CreateRegistryWebhook_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateRegistryWebhook_SdkV2) {
	if !from.HttpUrlSpec.IsNull() && !from.HttpUrlSpec.IsUnknown() {
		if toHttpUrlSpec, ok := to.GetHttpUrlSpec(ctx); ok {
			if fromHttpUrlSpec, ok := from.GetHttpUrlSpec(ctx); ok {
				toHttpUrlSpec.SyncFieldsDuringRead(ctx, fromHttpUrlSpec)
				to.SetHttpUrlSpec(ctx, toHttpUrlSpec)
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
}

func (m CreateRegistryWebhook_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["events"] = attrs["events"].SetRequired()
	attrs["http_url_spec"] = attrs["http_url_spec"].SetOptional()
	attrs["http_url_spec"] = attrs["http_url_spec"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["job_spec"] = attrs["job_spec"].SetOptional()
	attrs["job_spec"] = attrs["job_spec"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["model_name"] = attrs["model_name"].SetOptional()
	attrs["status"] = attrs["status"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateRegistryWebhook.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateRegistryWebhook_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"events":        reflect.TypeOf(types.String{}),
		"http_url_spec": reflect.TypeOf(HttpUrlSpec_SdkV2{}),
		"job_spec":      reflect.TypeOf(JobSpec_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateRegistryWebhook_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateRegistryWebhook_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":   m.Description,
			"events":        m.Events,
			"http_url_spec": m.HttpUrlSpec,
			"job_spec":      m.JobSpec,
			"model_name":    m.ModelName,
			"status":        m.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateRegistryWebhook_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"events": basetypes.ListType{
				ElemType: types.StringType,
			},
			"http_url_spec": basetypes.ListType{
				ElemType: HttpUrlSpec_SdkV2{}.Type(ctx),
			},
			"job_spec": basetypes.ListType{
				ElemType: JobSpec_SdkV2{}.Type(ctx),
			},
			"model_name": types.StringType,
			"status":     types.StringType,
		},
	}
}

// GetEvents returns the value of the Events field in CreateRegistryWebhook_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateRegistryWebhook_SdkV2) GetEvents(ctx context.Context) ([]types.String, bool) {
	if m.Events.IsNull() || m.Events.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Events.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEvents sets the value of the Events field in CreateRegistryWebhook_SdkV2.
func (m *CreateRegistryWebhook_SdkV2) SetEvents(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["events"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Events = types.ListValueMust(t, vs)
}

// GetHttpUrlSpec returns the value of the HttpUrlSpec field in CreateRegistryWebhook_SdkV2 as
// a HttpUrlSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateRegistryWebhook_SdkV2) GetHttpUrlSpec(ctx context.Context) (HttpUrlSpec_SdkV2, bool) {
	var e HttpUrlSpec_SdkV2
	if m.HttpUrlSpec.IsNull() || m.HttpUrlSpec.IsUnknown() {
		return e, false
	}
	var v []HttpUrlSpec_SdkV2
	d := m.HttpUrlSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetHttpUrlSpec sets the value of the HttpUrlSpec field in CreateRegistryWebhook_SdkV2.
func (m *CreateRegistryWebhook_SdkV2) SetHttpUrlSpec(ctx context.Context, v HttpUrlSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["http_url_spec"]
	m.HttpUrlSpec = types.ListValueMust(t, vs)
}

// GetJobSpec returns the value of the JobSpec field in CreateRegistryWebhook_SdkV2 as
// a JobSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateRegistryWebhook_SdkV2) GetJobSpec(ctx context.Context) (JobSpec_SdkV2, bool) {
	var e JobSpec_SdkV2
	if m.JobSpec.IsNull() || m.JobSpec.IsUnknown() {
		return e, false
	}
	var v []JobSpec_SdkV2
	d := m.JobSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetJobSpec sets the value of the JobSpec field in CreateRegistryWebhook_SdkV2.
func (m *CreateRegistryWebhook_SdkV2) SetJobSpec(ctx context.Context, v JobSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["job_spec"]
	m.JobSpec = types.ListValueMust(t, vs)
}

type CreateRun_SdkV2 struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"experiment_id"`
	// The name of the run.
	RunName types.String `tfsdk:"run_name"`
	// Unix timestamp in milliseconds of when the run started.
	StartTime types.Int64 `tfsdk:"start_time"`
	// Additional metadata for run.
	Tags types.List `tfsdk:"tags"`
	// ID of the user executing the run. This field is deprecated as of MLflow
	// 1.0, and will be removed in a future MLflow release. Use 'mlflow.user'
	// tag instead.
	UserId types.String `tfsdk:"user_id"`
}

func (to *CreateRun_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateRun_SdkV2) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *CreateRun_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateRun_SdkV2) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m CreateRun_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["experiment_id"] = attrs["experiment_id"].SetOptional()
	attrs["run_name"] = attrs["run_name"].SetOptional()
	attrs["start_time"] = attrs["start_time"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()
	attrs["user_id"] = attrs["user_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateRun.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateRun_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(RunTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateRun_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateRun_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": m.ExperimentId,
			"run_name":      m.RunName,
			"start_time":    m.StartTime,
			"tags":          m.Tags,
			"user_id":       m.UserId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateRun_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
			"run_name":      types.StringType,
			"start_time":    types.Int64Type,
			"tags": basetypes.ListType{
				ElemType: RunTag_SdkV2{}.Type(ctx),
			},
			"user_id": types.StringType,
		},
	}
}

// GetTags returns the value of the Tags field in CreateRun_SdkV2 as
// a slice of RunTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateRun_SdkV2) GetTags(ctx context.Context) ([]RunTag_SdkV2, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []RunTag_SdkV2
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in CreateRun_SdkV2.
func (m *CreateRun_SdkV2) SetTags(ctx context.Context, v []RunTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type CreateRunResponse_SdkV2 struct {
	// The newly created run.
	Run types.List `tfsdk:"run"`
}

func (to *CreateRunResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateRunResponse_SdkV2) {
	if !from.Run.IsNull() && !from.Run.IsUnknown() {
		if toRun, ok := to.GetRun(ctx); ok {
			if fromRun, ok := from.GetRun(ctx); ok {
				// Recursively sync the fields of Run
				toRun.SyncFieldsDuringCreateOrUpdate(ctx, fromRun)
				to.SetRun(ctx, toRun)
			}
		}
	}
}

func (to *CreateRunResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateRunResponse_SdkV2) {
	if !from.Run.IsNull() && !from.Run.IsUnknown() {
		if toRun, ok := to.GetRun(ctx); ok {
			if fromRun, ok := from.GetRun(ctx); ok {
				toRun.SyncFieldsDuringRead(ctx, fromRun)
				to.SetRun(ctx, toRun)
			}
		}
	}
}

func (m CreateRunResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["run"] = attrs["run"].SetOptional()
	attrs["run"] = attrs["run"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateRunResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateRunResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"run": reflect.TypeOf(Run_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateRunResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateRunResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run": m.Run,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateRunResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run": basetypes.ListType{
				ElemType: Run_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetRun returns the value of the Run field in CreateRunResponse_SdkV2 as
// a Run_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateRunResponse_SdkV2) GetRun(ctx context.Context) (Run_SdkV2, bool) {
	var e Run_SdkV2
	if m.Run.IsNull() || m.Run.IsUnknown() {
		return e, false
	}
	var v []Run_SdkV2
	d := m.Run.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRun sets the value of the Run field in CreateRunResponse_SdkV2.
func (m *CreateRunResponse_SdkV2) SetRun(ctx context.Context, v Run_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["run"]
	m.Run = types.ListValueMust(t, vs)
}

// Details required to create a model version stage transition request.
type CreateTransitionRequest_SdkV2 struct {
	// User-provided comment on the action.
	Comment types.String `tfsdk:"comment"`
	// Name of the model.
	Name types.String `tfsdk:"name"`
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage types.String `tfsdk:"stage"`
	// Version of the model.
	Version types.String `tfsdk:"version"`
}

func (to *CreateTransitionRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateTransitionRequest_SdkV2) {
}

func (to *CreateTransitionRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateTransitionRequest_SdkV2) {
}

func (m CreateTransitionRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["stage"] = attrs["stage"].SetRequired()
	attrs["version"] = attrs["version"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateTransitionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateTransitionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateTransitionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateTransitionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment": m.Comment,
			"name":    m.Name,
			"stage":   m.Stage,
			"version": m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateTransitionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment": types.StringType,
			"name":    types.StringType,
			"stage":   types.StringType,
			"version": types.StringType,
		},
	}
}

type CreateTransitionRequestResponse_SdkV2 struct {
	// New activity generated for stage transition request.
	Request types.List `tfsdk:"request"`
}

func (to *CreateTransitionRequestResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateTransitionRequestResponse_SdkV2) {
	if !from.Request.IsNull() && !from.Request.IsUnknown() {
		if toRequest, ok := to.GetRequest(ctx); ok {
			if fromRequest, ok := from.GetRequest(ctx); ok {
				// Recursively sync the fields of Request
				toRequest.SyncFieldsDuringCreateOrUpdate(ctx, fromRequest)
				to.SetRequest(ctx, toRequest)
			}
		}
	}
}

func (to *CreateTransitionRequestResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateTransitionRequestResponse_SdkV2) {
	if !from.Request.IsNull() && !from.Request.IsUnknown() {
		if toRequest, ok := to.GetRequest(ctx); ok {
			if fromRequest, ok := from.GetRequest(ctx); ok {
				toRequest.SyncFieldsDuringRead(ctx, fromRequest)
				to.SetRequest(ctx, toRequest)
			}
		}
	}
}

func (m CreateTransitionRequestResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["request"] = attrs["request"].SetOptional()
	attrs["request"] = attrs["request"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateTransitionRequestResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateTransitionRequestResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"request": reflect.TypeOf(TransitionRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateTransitionRequestResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateTransitionRequestResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"request": m.Request,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateTransitionRequestResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"request": basetypes.ListType{
				ElemType: TransitionRequest_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetRequest returns the value of the Request field in CreateTransitionRequestResponse_SdkV2 as
// a TransitionRequest_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateTransitionRequestResponse_SdkV2) GetRequest(ctx context.Context) (TransitionRequest_SdkV2, bool) {
	var e TransitionRequest_SdkV2
	if m.Request.IsNull() || m.Request.IsUnknown() {
		return e, false
	}
	var v []TransitionRequest_SdkV2
	d := m.Request.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRequest sets the value of the Request field in CreateTransitionRequestResponse_SdkV2.
func (m *CreateTransitionRequestResponse_SdkV2) SetRequest(ctx context.Context, v TransitionRequest_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["request"]
	m.Request = types.ListValueMust(t, vs)
}

type CreateWebhookResponse_SdkV2 struct {
	Webhook types.List `tfsdk:"webhook"`
}

func (to *CreateWebhookResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateWebhookResponse_SdkV2) {
	if !from.Webhook.IsNull() && !from.Webhook.IsUnknown() {
		if toWebhook, ok := to.GetWebhook(ctx); ok {
			if fromWebhook, ok := from.GetWebhook(ctx); ok {
				// Recursively sync the fields of Webhook
				toWebhook.SyncFieldsDuringCreateOrUpdate(ctx, fromWebhook)
				to.SetWebhook(ctx, toWebhook)
			}
		}
	}
}

func (to *CreateWebhookResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateWebhookResponse_SdkV2) {
	if !from.Webhook.IsNull() && !from.Webhook.IsUnknown() {
		if toWebhook, ok := to.GetWebhook(ctx); ok {
			if fromWebhook, ok := from.GetWebhook(ctx); ok {
				toWebhook.SyncFieldsDuringRead(ctx, fromWebhook)
				to.SetWebhook(ctx, toWebhook)
			}
		}
	}
}

func (m CreateWebhookResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["webhook"] = attrs["webhook"].SetOptional()
	attrs["webhook"] = attrs["webhook"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateWebhookResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateWebhookResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"webhook": reflect.TypeOf(RegistryWebhook_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateWebhookResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateWebhookResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"webhook": m.Webhook,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateWebhookResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"webhook": basetypes.ListType{
				ElemType: RegistryWebhook_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetWebhook returns the value of the Webhook field in CreateWebhookResponse_SdkV2 as
// a RegistryWebhook_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateWebhookResponse_SdkV2) GetWebhook(ctx context.Context) (RegistryWebhook_SdkV2, bool) {
	var e RegistryWebhook_SdkV2
	if m.Webhook.IsNull() || m.Webhook.IsUnknown() {
		return e, false
	}
	var v []RegistryWebhook_SdkV2
	d := m.Webhook.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWebhook sets the value of the Webhook field in CreateWebhookResponse_SdkV2.
func (m *CreateWebhookResponse_SdkV2) SetWebhook(ctx context.Context, v RegistryWebhook_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["webhook"]
	m.Webhook = types.ListValueMust(t, vs)
}

type DataSource_SdkV2 struct {
	DeltaTableSource types.List `tfsdk:"delta_table_source"`
}

func (to *DataSource_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DataSource_SdkV2) {
	if !from.DeltaTableSource.IsNull() && !from.DeltaTableSource.IsUnknown() {
		if toDeltaTableSource, ok := to.GetDeltaTableSource(ctx); ok {
			if fromDeltaTableSource, ok := from.GetDeltaTableSource(ctx); ok {
				// Recursively sync the fields of DeltaTableSource
				toDeltaTableSource.SyncFieldsDuringCreateOrUpdate(ctx, fromDeltaTableSource)
				to.SetDeltaTableSource(ctx, toDeltaTableSource)
			}
		}
	}
}

func (to *DataSource_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DataSource_SdkV2) {
	if !from.DeltaTableSource.IsNull() && !from.DeltaTableSource.IsUnknown() {
		if toDeltaTableSource, ok := to.GetDeltaTableSource(ctx); ok {
			if fromDeltaTableSource, ok := from.GetDeltaTableSource(ctx); ok {
				toDeltaTableSource.SyncFieldsDuringRead(ctx, fromDeltaTableSource)
				to.SetDeltaTableSource(ctx, toDeltaTableSource)
			}
		}
	}
}

func (m DataSource_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["delta_table_source"] = attrs["delta_table_source"].SetOptional()
	attrs["delta_table_source"] = attrs["delta_table_source"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DataSource.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DataSource_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"delta_table_source": reflect.TypeOf(DeltaTableSource_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DataSource_SdkV2
// only implements ToObjectValue() and Type().
func (m DataSource_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"delta_table_source": m.DeltaTableSource,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DataSource_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"delta_table_source": basetypes.ListType{
				ElemType: DeltaTableSource_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetDeltaTableSource returns the value of the DeltaTableSource field in DataSource_SdkV2 as
// a DeltaTableSource_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *DataSource_SdkV2) GetDeltaTableSource(ctx context.Context) (DeltaTableSource_SdkV2, bool) {
	var e DeltaTableSource_SdkV2
	if m.DeltaTableSource.IsNull() || m.DeltaTableSource.IsUnknown() {
		return e, false
	}
	var v []DeltaTableSource_SdkV2
	d := m.DeltaTableSource.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDeltaTableSource sets the value of the DeltaTableSource field in DataSource_SdkV2.
func (m *DataSource_SdkV2) SetDeltaTableSource(ctx context.Context, v DeltaTableSource_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["delta_table_source"]
	m.DeltaTableSource = types.ListValueMust(t, vs)
}

// Dataset. Represents a reference to data used for training, testing, or
// evaluation during the model development process.
type Dataset_SdkV2 struct {
	// Dataset digest, e.g. an md5 hash of the dataset that uniquely identifies
	// it within datasets of the same name.
	Digest types.String `tfsdk:"digest"`
	// The name of the dataset. E.g. “my.uc.table@2” “nyc-taxi-dataset”,
	// “fantastic-elk-3”
	Name types.String `tfsdk:"name"`
	// The profile of the dataset. Summary statistics for the dataset, such as
	// the number of rows in a table, the mean / std / mode of each column in a
	// table, or the number of elements in an array.
	Profile types.String `tfsdk:"profile"`
	// The schema of the dataset. E.g., MLflow ColSpec JSON for a dataframe,
	// MLflow TensorSpec JSON for an ndarray, or another schema format.
	Schema types.String `tfsdk:"schema"`
	// Source information for the dataset. Note that the source may not exactly
	// reproduce the dataset if it was transformed / modified before use with
	// MLflow.
	Source types.String `tfsdk:"source"`
	// The type of the dataset source, e.g. ‘databricks-uc-table’,
	// ‘DBFS’, ‘S3’, ...
	SourceType types.String `tfsdk:"source_type"`
}

func (to *Dataset_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Dataset_SdkV2) {
}

func (to *Dataset_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Dataset_SdkV2) {
}

func (m Dataset_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["digest"] = attrs["digest"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["profile"] = attrs["profile"].SetOptional()
	attrs["schema"] = attrs["schema"].SetOptional()
	attrs["source"] = attrs["source"].SetRequired()
	attrs["source_type"] = attrs["source_type"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Dataset.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Dataset_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Dataset_SdkV2
// only implements ToObjectValue() and Type().
func (m Dataset_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"digest":      m.Digest,
			"name":        m.Name,
			"profile":     m.Profile,
			"schema":      m.Schema,
			"source":      m.Source,
			"source_type": m.SourceType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Dataset_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"digest":      types.StringType,
			"name":        types.StringType,
			"profile":     types.StringType,
			"schema":      types.StringType,
			"source":      types.StringType,
			"source_type": types.StringType,
		},
	}
}

// DatasetInput. Represents a dataset and input tags.
type DatasetInput_SdkV2 struct {
	// The dataset being used as a Run input.
	Dataset types.List `tfsdk:"dataset"`
	// A list of tags for the dataset input, e.g. a “context” tag with value
	// “training”
	Tags types.List `tfsdk:"tags"`
}

func (to *DatasetInput_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatasetInput_SdkV2) {
	if !from.Dataset.IsNull() && !from.Dataset.IsUnknown() {
		if toDataset, ok := to.GetDataset(ctx); ok {
			if fromDataset, ok := from.GetDataset(ctx); ok {
				// Recursively sync the fields of Dataset
				toDataset.SyncFieldsDuringCreateOrUpdate(ctx, fromDataset)
				to.SetDataset(ctx, toDataset)
			}
		}
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *DatasetInput_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DatasetInput_SdkV2) {
	if !from.Dataset.IsNull() && !from.Dataset.IsUnknown() {
		if toDataset, ok := to.GetDataset(ctx); ok {
			if fromDataset, ok := from.GetDataset(ctx); ok {
				toDataset.SyncFieldsDuringRead(ctx, fromDataset)
				to.SetDataset(ctx, toDataset)
			}
		}
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m DatasetInput_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dataset"] = attrs["dataset"].SetRequired()
	attrs["dataset"] = attrs["dataset"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["tags"] = attrs["tags"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatasetInput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DatasetInput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dataset": reflect.TypeOf(Dataset_SdkV2{}),
		"tags":    reflect.TypeOf(InputTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatasetInput_SdkV2
// only implements ToObjectValue() and Type().
func (m DatasetInput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dataset": m.Dataset,
			"tags":    m.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DatasetInput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dataset": basetypes.ListType{
				ElemType: Dataset_SdkV2{}.Type(ctx),
			},
			"tags": basetypes.ListType{
				ElemType: InputTag_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetDataset returns the value of the Dataset field in DatasetInput_SdkV2 as
// a Dataset_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *DatasetInput_SdkV2) GetDataset(ctx context.Context) (Dataset_SdkV2, bool) {
	var e Dataset_SdkV2
	if m.Dataset.IsNull() || m.Dataset.IsUnknown() {
		return e, false
	}
	var v []Dataset_SdkV2
	d := m.Dataset.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDataset sets the value of the Dataset field in DatasetInput_SdkV2.
func (m *DatasetInput_SdkV2) SetDataset(ctx context.Context, v Dataset_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["dataset"]
	m.Dataset = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in DatasetInput_SdkV2 as
// a slice of InputTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *DatasetInput_SdkV2) GetTags(ctx context.Context) ([]InputTag_SdkV2, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []InputTag_SdkV2
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in DatasetInput_SdkV2.
func (m *DatasetInput_SdkV2) SetTags(ctx context.Context, v []InputTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type DeleteCommentRequest_SdkV2 struct {
	// Unique identifier of an activity
	Id types.String `tfsdk:"-"`
}

func (to *DeleteCommentRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteCommentRequest_SdkV2) {
}

func (to *DeleteCommentRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteCommentRequest_SdkV2) {
}

func (m DeleteCommentRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCommentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteCommentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCommentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteCommentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteCommentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteCommentResponse_SdkV2 struct {
}

func (to *DeleteCommentResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteCommentResponse_SdkV2) {
}

func (to *DeleteCommentResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteCommentResponse_SdkV2) {
}

func (m DeleteCommentResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCommentResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteCommentResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCommentResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteCommentResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteCommentResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteExperiment_SdkV2 struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"experiment_id"`
}

func (to *DeleteExperiment_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteExperiment_SdkV2) {
}

func (to *DeleteExperiment_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteExperiment_SdkV2) {
}

func (m DeleteExperiment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["experiment_id"] = attrs["experiment_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteExperiment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteExperiment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteExperiment_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteExperiment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": m.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteExperiment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
		},
	}
}

type DeleteExperimentResponse_SdkV2 struct {
}

func (to *DeleteExperimentResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteExperimentResponse_SdkV2) {
}

func (to *DeleteExperimentResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteExperimentResponse_SdkV2) {
}

func (m DeleteExperimentResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteExperimentResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteExperimentResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteExperimentResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteExperimentResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteExperimentResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteFeatureRequest_SdkV2 struct {
	// Name of the feature to delete.
	FullName types.String `tfsdk:"-"`
}

func (to *DeleteFeatureRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteFeatureRequest_SdkV2) {
}

func (to *DeleteFeatureRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteFeatureRequest_SdkV2) {
}

func (m DeleteFeatureRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["full_name"] = attrs["full_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteFeatureRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteFeatureRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteFeatureRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteFeatureRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"full_name": m.FullName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteFeatureRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_name": types.StringType,
		},
	}
}

type DeleteFeatureTagRequest_SdkV2 struct {
	// The name of the feature within the feature table.
	FeatureName types.String `tfsdk:"-"`
	// The key of the tag to delete.
	Key types.String `tfsdk:"-"`
	// The name of the feature table.
	TableName types.String `tfsdk:"-"`
}

func (to *DeleteFeatureTagRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteFeatureTagRequest_SdkV2) {
}

func (to *DeleteFeatureTagRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteFeatureTagRequest_SdkV2) {
}

func (m DeleteFeatureTagRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["table_name"] = attrs["table_name"].SetRequired()
	attrs["feature_name"] = attrs["feature_name"].SetRequired()
	attrs["key"] = attrs["key"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteFeatureTagRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteFeatureTagRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteFeatureTagRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteFeatureTagRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_name": m.FeatureName,
			"key":          m.Key,
			"table_name":   m.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteFeatureTagRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"feature_name": types.StringType,
			"key":          types.StringType,
			"table_name":   types.StringType,
		},
	}
}

type DeleteLoggedModelRequest_SdkV2 struct {
	// The ID of the logged model to delete.
	ModelId types.String `tfsdk:"-"`
}

func (to *DeleteLoggedModelRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteLoggedModelRequest_SdkV2) {
}

func (to *DeleteLoggedModelRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteLoggedModelRequest_SdkV2) {
}

func (m DeleteLoggedModelRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["model_id"] = attrs["model_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteLoggedModelRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteLoggedModelRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteLoggedModelRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteLoggedModelRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_id": m.ModelId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteLoggedModelRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_id": types.StringType,
		},
	}
}

type DeleteLoggedModelResponse_SdkV2 struct {
}

func (to *DeleteLoggedModelResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteLoggedModelResponse_SdkV2) {
}

func (to *DeleteLoggedModelResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteLoggedModelResponse_SdkV2) {
}

func (m DeleteLoggedModelResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteLoggedModelResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteLoggedModelResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteLoggedModelResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteLoggedModelResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteLoggedModelResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteLoggedModelTagRequest_SdkV2 struct {
	// The ID of the logged model to delete the tag from.
	ModelId types.String `tfsdk:"-"`
	// The tag key.
	TagKey types.String `tfsdk:"-"`
}

func (to *DeleteLoggedModelTagRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteLoggedModelTagRequest_SdkV2) {
}

func (to *DeleteLoggedModelTagRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteLoggedModelTagRequest_SdkV2) {
}

func (m DeleteLoggedModelTagRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["model_id"] = attrs["model_id"].SetRequired()
	attrs["tag_key"] = attrs["tag_key"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteLoggedModelTagRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteLoggedModelTagRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteLoggedModelTagRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteLoggedModelTagRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_id": m.ModelId,
			"tag_key":  m.TagKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteLoggedModelTagRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_id": types.StringType,
			"tag_key":  types.StringType,
		},
	}
}

type DeleteLoggedModelTagResponse_SdkV2 struct {
}

func (to *DeleteLoggedModelTagResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteLoggedModelTagResponse_SdkV2) {
}

func (to *DeleteLoggedModelTagResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteLoggedModelTagResponse_SdkV2) {
}

func (m DeleteLoggedModelTagResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteLoggedModelTagResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteLoggedModelTagResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteLoggedModelTagResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteLoggedModelTagResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteLoggedModelTagResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteModelRequest_SdkV2 struct {
	// Registered model unique name identifier.
	Name types.String `tfsdk:"-"`
}

func (to *DeleteModelRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteModelRequest_SdkV2) {
}

func (to *DeleteModelRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteModelRequest_SdkV2) {
}

func (m DeleteModelRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteModelRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteModelRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteModelRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteModelRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteModelResponse_SdkV2 struct {
}

func (to *DeleteModelResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteModelResponse_SdkV2) {
}

func (to *DeleteModelResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteModelResponse_SdkV2) {
}

func (m DeleteModelResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteModelResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteModelResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteModelResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteModelResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteModelTagRequest_SdkV2 struct {
	// Name of the tag. The name must be an exact match; wild-card deletion is
	// not supported. Maximum size is 250 bytes.
	Key types.String `tfsdk:"-"`
	// Name of the registered model that the tag was logged under.
	Name types.String `tfsdk:"-"`
}

func (to *DeleteModelTagRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteModelTagRequest_SdkV2) {
}

func (to *DeleteModelTagRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteModelTagRequest_SdkV2) {
}

func (m DeleteModelTagRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["key"] = attrs["key"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteModelTagRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteModelTagRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelTagRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteModelTagRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":  m.Key,
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteModelTagRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":  types.StringType,
			"name": types.StringType,
		},
	}
}

type DeleteModelTagResponse_SdkV2 struct {
}

func (to *DeleteModelTagResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteModelTagResponse_SdkV2) {
}

func (to *DeleteModelTagResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteModelTagResponse_SdkV2) {
}

func (m DeleteModelTagResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteModelTagResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteModelTagResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelTagResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteModelTagResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteModelTagResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteModelVersionRequest_SdkV2 struct {
	// Name of the registered model
	Name types.String `tfsdk:"-"`
	// Model version number
	Version types.String `tfsdk:"-"`
}

func (to *DeleteModelVersionRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteModelVersionRequest_SdkV2) {
}

func (to *DeleteModelVersionRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteModelVersionRequest_SdkV2) {
}

func (m DeleteModelVersionRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["version"] = attrs["version"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteModelVersionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteModelVersionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelVersionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteModelVersionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":    m.Name,
			"version": m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteModelVersionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":    types.StringType,
			"version": types.StringType,
		},
	}
}

type DeleteModelVersionResponse_SdkV2 struct {
}

func (to *DeleteModelVersionResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteModelVersionResponse_SdkV2) {
}

func (to *DeleteModelVersionResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteModelVersionResponse_SdkV2) {
}

func (m DeleteModelVersionResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteModelVersionResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteModelVersionResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelVersionResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteModelVersionResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteModelVersionResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteModelVersionTagRequest_SdkV2 struct {
	// Name of the tag. The name must be an exact match; wild-card deletion is
	// not supported. Maximum size is 250 bytes.
	Key types.String `tfsdk:"-"`
	// Name of the registered model that the tag was logged under.
	Name types.String `tfsdk:"-"`
	// Model version number that the tag was logged under.
	Version types.String `tfsdk:"-"`
}

func (to *DeleteModelVersionTagRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteModelVersionTagRequest_SdkV2) {
}

func (to *DeleteModelVersionTagRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteModelVersionTagRequest_SdkV2) {
}

func (m DeleteModelVersionTagRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["version"] = attrs["version"].SetRequired()
	attrs["key"] = attrs["key"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteModelVersionTagRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteModelVersionTagRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelVersionTagRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteModelVersionTagRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":     m.Key,
			"name":    m.Name,
			"version": m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteModelVersionTagRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":     types.StringType,
			"name":    types.StringType,
			"version": types.StringType,
		},
	}
}

type DeleteModelVersionTagResponse_SdkV2 struct {
}

func (to *DeleteModelVersionTagResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteModelVersionTagResponse_SdkV2) {
}

func (to *DeleteModelVersionTagResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteModelVersionTagResponse_SdkV2) {
}

func (m DeleteModelVersionTagResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteModelVersionTagResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteModelVersionTagResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelVersionTagResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteModelVersionTagResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteModelVersionTagResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteOnlineStoreRequest_SdkV2 struct {
	// Name of the online store to delete.
	Name types.String `tfsdk:"-"`
}

func (to *DeleteOnlineStoreRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteOnlineStoreRequest_SdkV2) {
}

func (to *DeleteOnlineStoreRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteOnlineStoreRequest_SdkV2) {
}

func (m DeleteOnlineStoreRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteOnlineStoreRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteOnlineStoreRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteOnlineStoreRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteOnlineStoreRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteOnlineStoreRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteRun_SdkV2 struct {
	// ID of the run to delete.
	RunId types.String `tfsdk:"run_id"`
}

func (to *DeleteRun_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteRun_SdkV2) {
}

func (to *DeleteRun_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteRun_SdkV2) {
}

func (m DeleteRun_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["run_id"] = attrs["run_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteRun.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteRun_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRun_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteRun_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run_id": m.RunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteRun_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run_id": types.StringType,
		},
	}
}

type DeleteRunResponse_SdkV2 struct {
}

func (to *DeleteRunResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteRunResponse_SdkV2) {
}

func (to *DeleteRunResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteRunResponse_SdkV2) {
}

func (m DeleteRunResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteRunResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteRunResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRunResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteRunResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteRunResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteRuns_SdkV2 struct {
	// The ID of the experiment containing the runs to delete.
	ExperimentId types.String `tfsdk:"experiment_id"`
	// An optional positive integer indicating the maximum number of runs to
	// delete. The maximum allowed value for max_runs is 10000.
	MaxRuns types.Int64 `tfsdk:"max_runs"`
	// The maximum creation timestamp in milliseconds since the UNIX epoch for
	// deleting runs. Only runs created prior to or at this timestamp are
	// deleted.
	MaxTimestampMillis types.Int64 `tfsdk:"max_timestamp_millis"`
}

func (to *DeleteRuns_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteRuns_SdkV2) {
}

func (to *DeleteRuns_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteRuns_SdkV2) {
}

func (m DeleteRuns_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["experiment_id"] = attrs["experiment_id"].SetRequired()
	attrs["max_runs"] = attrs["max_runs"].SetOptional()
	attrs["max_timestamp_millis"] = attrs["max_timestamp_millis"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteRuns.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteRuns_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRuns_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteRuns_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id":        m.ExperimentId,
			"max_runs":             m.MaxRuns,
			"max_timestamp_millis": m.MaxTimestampMillis,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteRuns_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id":        types.StringType,
			"max_runs":             types.Int64Type,
			"max_timestamp_millis": types.Int64Type,
		},
	}
}

type DeleteRunsResponse_SdkV2 struct {
	// The number of runs deleted.
	RunsDeleted types.Int64 `tfsdk:"runs_deleted"`
}

func (to *DeleteRunsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteRunsResponse_SdkV2) {
}

func (to *DeleteRunsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteRunsResponse_SdkV2) {
}

func (m DeleteRunsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["runs_deleted"] = attrs["runs_deleted"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteRunsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteRunsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRunsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteRunsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"runs_deleted": m.RunsDeleted,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteRunsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"runs_deleted": types.Int64Type,
		},
	}
}

type DeleteTag_SdkV2 struct {
	// Name of the tag. Maximum size is 255 bytes. Must be provided.
	Key types.String `tfsdk:"key"`
	// ID of the run that the tag was logged under. Must be provided.
	RunId types.String `tfsdk:"run_id"`
}

func (to *DeleteTag_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteTag_SdkV2) {
}

func (to *DeleteTag_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteTag_SdkV2) {
}

func (m DeleteTag_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetRequired()
	attrs["run_id"] = attrs["run_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteTag.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteTag_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteTag_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteTag_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":    m.Key,
			"run_id": m.RunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteTag_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":    types.StringType,
			"run_id": types.StringType,
		},
	}
}

type DeleteTagResponse_SdkV2 struct {
}

func (to *DeleteTagResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteTagResponse_SdkV2) {
}

func (to *DeleteTagResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteTagResponse_SdkV2) {
}

func (m DeleteTagResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteTagResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteTagResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteTagResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteTagResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteTagResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteTransitionRequestRequest_SdkV2 struct {
	// User-provided comment on the action.
	Comment types.String `tfsdk:"-"`
	// Username of the user who created this request. Of the transition requests
	// matching the specified details, only the one transition created by this
	// user will be deleted.
	Creator types.String `tfsdk:"-"`
	// Name of the model.
	Name types.String `tfsdk:"-"`
	// Target stage of the transition request. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage types.String `tfsdk:"-"`
	// Version of the model.
	Version types.String `tfsdk:"-"`
}

func (to *DeleteTransitionRequestRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteTransitionRequestRequest_SdkV2) {
}

func (to *DeleteTransitionRequestRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteTransitionRequestRequest_SdkV2) {
}

func (m DeleteTransitionRequestRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["version"] = attrs["version"].SetRequired()
	attrs["stage"] = attrs["stage"].SetRequired()
	attrs["creator"] = attrs["creator"].SetRequired()
	attrs["comment"] = attrs["comment"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteTransitionRequestRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteTransitionRequestRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteTransitionRequestRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteTransitionRequestRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment": m.Comment,
			"creator": m.Creator,
			"name":    m.Name,
			"stage":   m.Stage,
			"version": m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteTransitionRequestRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment": types.StringType,
			"creator": types.StringType,
			"name":    types.StringType,
			"stage":   types.StringType,
			"version": types.StringType,
		},
	}
}

type DeleteTransitionRequestResponse_SdkV2 struct {
	// New activity generated as a result of this operation.
	Activity types.List `tfsdk:"activity"`
}

func (to *DeleteTransitionRequestResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteTransitionRequestResponse_SdkV2) {
	if !from.Activity.IsNull() && !from.Activity.IsUnknown() {
		if toActivity, ok := to.GetActivity(ctx); ok {
			if fromActivity, ok := from.GetActivity(ctx); ok {
				// Recursively sync the fields of Activity
				toActivity.SyncFieldsDuringCreateOrUpdate(ctx, fromActivity)
				to.SetActivity(ctx, toActivity)
			}
		}
	}
}

func (to *DeleteTransitionRequestResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteTransitionRequestResponse_SdkV2) {
	if !from.Activity.IsNull() && !from.Activity.IsUnknown() {
		if toActivity, ok := to.GetActivity(ctx); ok {
			if fromActivity, ok := from.GetActivity(ctx); ok {
				toActivity.SyncFieldsDuringRead(ctx, fromActivity)
				to.SetActivity(ctx, toActivity)
			}
		}
	}
}

func (m DeleteTransitionRequestResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["activity"] = attrs["activity"].SetOptional()
	attrs["activity"] = attrs["activity"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteTransitionRequestResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteTransitionRequestResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"activity": reflect.TypeOf(Activity_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteTransitionRequestResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteTransitionRequestResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"activity": m.Activity,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteTransitionRequestResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"activity": basetypes.ListType{
				ElemType: Activity_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetActivity returns the value of the Activity field in DeleteTransitionRequestResponse_SdkV2 as
// a Activity_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *DeleteTransitionRequestResponse_SdkV2) GetActivity(ctx context.Context) (Activity_SdkV2, bool) {
	var e Activity_SdkV2
	if m.Activity.IsNull() || m.Activity.IsUnknown() {
		return e, false
	}
	var v []Activity_SdkV2
	d := m.Activity.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetActivity sets the value of the Activity field in DeleteTransitionRequestResponse_SdkV2.
func (m *DeleteTransitionRequestResponse_SdkV2) SetActivity(ctx context.Context, v Activity_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["activity"]
	m.Activity = types.ListValueMust(t, vs)
}

type DeleteWebhookRequest_SdkV2 struct {
	// Webhook ID required to delete a registry webhook.
	Id types.String `tfsdk:"-"`
}

func (to *DeleteWebhookRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteWebhookRequest_SdkV2) {
}

func (to *DeleteWebhookRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteWebhookRequest_SdkV2) {
}

func (m DeleteWebhookRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteWebhookRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteWebhookRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWebhookRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteWebhookRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteWebhookRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteWebhookResponse_SdkV2 struct {
}

func (to *DeleteWebhookResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteWebhookResponse_SdkV2) {
}

func (to *DeleteWebhookResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteWebhookResponse_SdkV2) {
}

func (m DeleteWebhookResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteWebhookResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteWebhookResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWebhookResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteWebhookResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteWebhookResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeltaTableSource_SdkV2 struct {
	// The entity columns of the Delta table.
	EntityColumns types.List `tfsdk:"entity_columns"`
	// The full three-part (catalog, schema, table) name of the Delta table.
	FullName types.String `tfsdk:"full_name"`
	// The timeseries column of the Delta table.
	TimeseriesColumn types.String `tfsdk:"timeseries_column"`
}

func (to *DeltaTableSource_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeltaTableSource_SdkV2) {
}

func (to *DeltaTableSource_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeltaTableSource_SdkV2) {
}

func (m DeltaTableSource_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["entity_columns"] = attrs["entity_columns"].SetRequired()
	attrs["full_name"] = attrs["full_name"].SetRequired()
	attrs["timeseries_column"] = attrs["timeseries_column"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeltaTableSource.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeltaTableSource_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"entity_columns": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeltaTableSource_SdkV2
// only implements ToObjectValue() and Type().
func (m DeltaTableSource_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"entity_columns":    m.EntityColumns,
			"full_name":         m.FullName,
			"timeseries_column": m.TimeseriesColumn,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeltaTableSource_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"entity_columns": basetypes.ListType{
				ElemType: types.StringType,
			},
			"full_name":         types.StringType,
			"timeseries_column": types.StringType,
		},
	}
}

// GetEntityColumns returns the value of the EntityColumns field in DeltaTableSource_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *DeltaTableSource_SdkV2) GetEntityColumns(ctx context.Context) ([]types.String, bool) {
	if m.EntityColumns.IsNull() || m.EntityColumns.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.EntityColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEntityColumns sets the value of the EntityColumns field in DeltaTableSource_SdkV2.
func (m *DeltaTableSource_SdkV2) SetEntityColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["entity_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EntityColumns = types.ListValueMust(t, vs)
}

// An experiment and its metadata.
type Experiment_SdkV2 struct {
	// Location where artifacts for the experiment are stored.
	ArtifactLocation types.String `tfsdk:"artifact_location"`
	// Creation time
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// Unique identifier for the experiment.
	ExperimentId types.String `tfsdk:"experiment_id"`
	// Last update time
	LastUpdateTime types.Int64 `tfsdk:"last_update_time"`
	// Current life cycle stage of the experiment: "active" or "deleted".
	// Deleted experiments are not returned by APIs.
	LifecycleStage types.String `tfsdk:"lifecycle_stage"`
	// Human readable name that identifies the experiment.
	Name types.String `tfsdk:"name"`
	// Tags: Additional metadata key-value pairs.
	Tags types.List `tfsdk:"tags"`
}

func (to *Experiment_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Experiment_SdkV2) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *Experiment_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Experiment_SdkV2) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m Experiment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["artifact_location"] = attrs["artifact_location"].SetOptional()
	attrs["creation_time"] = attrs["creation_time"].SetOptional()
	attrs["experiment_id"] = attrs["experiment_id"].SetOptional()
	attrs["last_update_time"] = attrs["last_update_time"].SetOptional()
	attrs["lifecycle_stage"] = attrs["lifecycle_stage"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Experiment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Experiment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(ExperimentTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Experiment_SdkV2
// only implements ToObjectValue() and Type().
func (m Experiment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"artifact_location": m.ArtifactLocation,
			"creation_time":     m.CreationTime,
			"experiment_id":     m.ExperimentId,
			"last_update_time":  m.LastUpdateTime,
			"lifecycle_stage":   m.LifecycleStage,
			"name":              m.Name,
			"tags":              m.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Experiment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"artifact_location": types.StringType,
			"creation_time":     types.Int64Type,
			"experiment_id":     types.StringType,
			"last_update_time":  types.Int64Type,
			"lifecycle_stage":   types.StringType,
			"name":              types.StringType,
			"tags": basetypes.ListType{
				ElemType: ExperimentTag_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetTags returns the value of the Tags field in Experiment_SdkV2 as
// a slice of ExperimentTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *Experiment_SdkV2) GetTags(ctx context.Context) ([]ExperimentTag_SdkV2, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []ExperimentTag_SdkV2
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in Experiment_SdkV2.
func (m *Experiment_SdkV2) SetTags(ctx context.Context, v []ExperimentTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type ExperimentAccessControlRequest_SdkV2 struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`

	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (to *ExperimentAccessControlRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExperimentAccessControlRequest_SdkV2) {
}

func (to *ExperimentAccessControlRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExperimentAccessControlRequest_SdkV2) {
}

func (m ExperimentAccessControlRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExperimentAccessControlRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ExperimentAccessControlRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExperimentAccessControlRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ExperimentAccessControlRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ExperimentAccessControlRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type ExperimentAccessControlResponse_SdkV2 struct {
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

func (to *ExperimentAccessControlResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExperimentAccessControlResponse_SdkV2) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (to *ExperimentAccessControlResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExperimentAccessControlResponse_SdkV2) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (m ExperimentAccessControlResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["all_permissions"] = attrs["all_permissions"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExperimentAccessControlResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ExperimentAccessControlResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(ExperimentPermission_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExperimentAccessControlResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ExperimentAccessControlResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ExperimentAccessControlResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: ExperimentPermission_SdkV2{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

// GetAllPermissions returns the value of the AllPermissions field in ExperimentAccessControlResponse_SdkV2 as
// a slice of ExperimentPermission_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ExperimentAccessControlResponse_SdkV2) GetAllPermissions(ctx context.Context) ([]ExperimentPermission_SdkV2, bool) {
	if m.AllPermissions.IsNull() || m.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []ExperimentPermission_SdkV2
	d := m.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in ExperimentAccessControlResponse_SdkV2.
func (m *ExperimentAccessControlResponse_SdkV2) SetAllPermissions(ctx context.Context, v []ExperimentPermission_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllPermissions = types.ListValueMust(t, vs)
}

type ExperimentPermission_SdkV2 struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *ExperimentPermission_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExperimentPermission_SdkV2) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (to *ExperimentPermission_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExperimentPermission_SdkV2) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (m ExperimentPermission_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["inherited"] = attrs["inherited"].SetOptional()
	attrs["inherited_from_object"] = attrs["inherited_from_object"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExperimentPermission.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ExperimentPermission_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExperimentPermission_SdkV2
// only implements ToObjectValue() and Type().
func (m ExperimentPermission_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             m.Inherited,
			"inherited_from_object": m.InheritedFromObject,
			"permission_level":      m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExperimentPermission_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetInheritedFromObject returns the value of the InheritedFromObject field in ExperimentPermission_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ExperimentPermission_SdkV2) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
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

// SetInheritedFromObject sets the value of the InheritedFromObject field in ExperimentPermission_SdkV2.
func (m *ExperimentPermission_SdkV2) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InheritedFromObject = types.ListValueMust(t, vs)
}

type ExperimentPermissions_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (to *ExperimentPermissions_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExperimentPermissions_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *ExperimentPermissions_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExperimentPermissions_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m ExperimentPermissions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["object_id"] = attrs["object_id"].SetOptional()
	attrs["object_type"] = attrs["object_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExperimentPermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ExperimentPermissions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(ExperimentAccessControlResponse_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExperimentPermissions_SdkV2
// only implements ToObjectValue() and Type().
func (m ExperimentPermissions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"object_id":           m.ObjectId,
			"object_type":         m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExperimentPermissions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: ExperimentAccessControlResponse_SdkV2{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in ExperimentPermissions_SdkV2 as
// a slice of ExperimentAccessControlResponse_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ExperimentPermissions_SdkV2) GetAccessControlList(ctx context.Context) ([]ExperimentAccessControlResponse_SdkV2, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []ExperimentAccessControlResponse_SdkV2
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in ExperimentPermissions_SdkV2.
func (m *ExperimentPermissions_SdkV2) SetAccessControlList(ctx context.Context, v []ExperimentAccessControlResponse_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

type ExperimentPermissionsDescription_SdkV2 struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *ExperimentPermissionsDescription_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExperimentPermissionsDescription_SdkV2) {
}

func (to *ExperimentPermissionsDescription_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExperimentPermissionsDescription_SdkV2) {
}

func (m ExperimentPermissionsDescription_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExperimentPermissionsDescription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ExperimentPermissionsDescription_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExperimentPermissionsDescription_SdkV2
// only implements ToObjectValue() and Type().
func (m ExperimentPermissionsDescription_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      m.Description,
			"permission_level": m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExperimentPermissionsDescription_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type ExperimentPermissionsRequest_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// The experiment for which to get or manage permissions.
	ExperimentId types.String `tfsdk:"-"`
}

func (to *ExperimentPermissionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExperimentPermissionsRequest_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *ExperimentPermissionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExperimentPermissionsRequest_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m ExperimentPermissionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["experiment_id"] = attrs["experiment_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExperimentPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ExperimentPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(ExperimentAccessControlRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExperimentPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ExperimentPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"experiment_id":       m.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExperimentPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: ExperimentAccessControlRequest_SdkV2{}.Type(ctx),
			},
			"experiment_id": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in ExperimentPermissionsRequest_SdkV2 as
// a slice of ExperimentAccessControlRequest_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ExperimentPermissionsRequest_SdkV2) GetAccessControlList(ctx context.Context) ([]ExperimentAccessControlRequest_SdkV2, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []ExperimentAccessControlRequest_SdkV2
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in ExperimentPermissionsRequest_SdkV2.
func (m *ExperimentPermissionsRequest_SdkV2) SetAccessControlList(ctx context.Context, v []ExperimentAccessControlRequest_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

// A tag for an experiment.
type ExperimentTag_SdkV2 struct {
	// The tag key.
	Key types.String `tfsdk:"key"`
	// The tag value.
	Value types.String `tfsdk:"value"`
}

func (to *ExperimentTag_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExperimentTag_SdkV2) {
}

func (to *ExperimentTag_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExperimentTag_SdkV2) {
}

func (m ExperimentTag_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetOptional()
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExperimentTag.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ExperimentTag_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExperimentTag_SdkV2
// only implements ToObjectValue() and Type().
func (m ExperimentTag_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExperimentTag_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

type Feature_SdkV2 struct {
	// The description of the feature.
	Description types.String `tfsdk:"description"`
	// The full three-part name (catalog, schema, name) of the feature.
	FullName types.String `tfsdk:"full_name"`
	// The function by which the feature is computed.
	Function types.List `tfsdk:"function"`
	// The input columns from which the feature is computed.
	Inputs types.List `tfsdk:"inputs"`
	// The data source of the feature.
	Source types.List `tfsdk:"source"`
	// The time window in which the feature is computed.
	TimeWindow types.List `tfsdk:"time_window"`
}

func (to *Feature_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Feature_SdkV2) {
	if !from.Function.IsNull() && !from.Function.IsUnknown() {
		if toFunction, ok := to.GetFunction(ctx); ok {
			if fromFunction, ok := from.GetFunction(ctx); ok {
				// Recursively sync the fields of Function
				toFunction.SyncFieldsDuringCreateOrUpdate(ctx, fromFunction)
				to.SetFunction(ctx, toFunction)
			}
		}
	}
	if !from.Source.IsNull() && !from.Source.IsUnknown() {
		if toSource, ok := to.GetSource(ctx); ok {
			if fromSource, ok := from.GetSource(ctx); ok {
				// Recursively sync the fields of Source
				toSource.SyncFieldsDuringCreateOrUpdate(ctx, fromSource)
				to.SetSource(ctx, toSource)
			}
		}
	}
	if !from.TimeWindow.IsNull() && !from.TimeWindow.IsUnknown() {
		if toTimeWindow, ok := to.GetTimeWindow(ctx); ok {
			if fromTimeWindow, ok := from.GetTimeWindow(ctx); ok {
				// Recursively sync the fields of TimeWindow
				toTimeWindow.SyncFieldsDuringCreateOrUpdate(ctx, fromTimeWindow)
				to.SetTimeWindow(ctx, toTimeWindow)
			}
		}
	}
}

func (to *Feature_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Feature_SdkV2) {
	if !from.Function.IsNull() && !from.Function.IsUnknown() {
		if toFunction, ok := to.GetFunction(ctx); ok {
			if fromFunction, ok := from.GetFunction(ctx); ok {
				toFunction.SyncFieldsDuringRead(ctx, fromFunction)
				to.SetFunction(ctx, toFunction)
			}
		}
	}
	if !from.Source.IsNull() && !from.Source.IsUnknown() {
		if toSource, ok := to.GetSource(ctx); ok {
			if fromSource, ok := from.GetSource(ctx); ok {
				toSource.SyncFieldsDuringRead(ctx, fromSource)
				to.SetSource(ctx, toSource)
			}
		}
	}
	if !from.TimeWindow.IsNull() && !from.TimeWindow.IsUnknown() {
		if toTimeWindow, ok := to.GetTimeWindow(ctx); ok {
			if fromTimeWindow, ok := from.GetTimeWindow(ctx); ok {
				toTimeWindow.SyncFieldsDuringRead(ctx, fromTimeWindow)
				to.SetTimeWindow(ctx, toTimeWindow)
			}
		}
	}
}

func (m Feature_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["full_name"] = attrs["full_name"].SetRequired()
	attrs["full_name"] = attrs["full_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["function"] = attrs["function"].SetRequired()
	attrs["function"] = attrs["function"].(tfschema.ListNestedAttributeBuilder).AddPlanModifier(listplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["function"] = attrs["function"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["inputs"] = attrs["inputs"].SetRequired()
	attrs["inputs"] = attrs["inputs"].(tfschema.ListAttributeBuilder).AddPlanModifier(listplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["source"] = attrs["source"].SetRequired()
	attrs["source"] = attrs["source"].(tfschema.ListNestedAttributeBuilder).AddPlanModifier(listplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["source"] = attrs["source"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["time_window"] = attrs["time_window"].SetRequired()
	attrs["time_window"] = attrs["time_window"].(tfschema.ListNestedAttributeBuilder).AddPlanModifier(listplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["time_window"] = attrs["time_window"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Feature.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Feature_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"function":    reflect.TypeOf(Function_SdkV2{}),
		"inputs":      reflect.TypeOf(types.String{}),
		"source":      reflect.TypeOf(DataSource_SdkV2{}),
		"time_window": reflect.TypeOf(TimeWindow_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Feature_SdkV2
// only implements ToObjectValue() and Type().
func (m Feature_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description": m.Description,
			"full_name":   m.FullName,
			"function":    m.Function,
			"inputs":      m.Inputs,
			"source":      m.Source,
			"time_window": m.TimeWindow,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Feature_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"full_name":   types.StringType,
			"function": basetypes.ListType{
				ElemType: Function_SdkV2{}.Type(ctx),
			},
			"inputs": basetypes.ListType{
				ElemType: types.StringType,
			},
			"source": basetypes.ListType{
				ElemType: DataSource_SdkV2{}.Type(ctx),
			},
			"time_window": basetypes.ListType{
				ElemType: TimeWindow_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetFunction returns the value of the Function field in Feature_SdkV2 as
// a Function_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Feature_SdkV2) GetFunction(ctx context.Context) (Function_SdkV2, bool) {
	var e Function_SdkV2
	if m.Function.IsNull() || m.Function.IsUnknown() {
		return e, false
	}
	var v []Function_SdkV2
	d := m.Function.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFunction sets the value of the Function field in Feature_SdkV2.
func (m *Feature_SdkV2) SetFunction(ctx context.Context, v Function_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["function"]
	m.Function = types.ListValueMust(t, vs)
}

// GetInputs returns the value of the Inputs field in Feature_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *Feature_SdkV2) GetInputs(ctx context.Context) ([]types.String, bool) {
	if m.Inputs.IsNull() || m.Inputs.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Inputs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInputs sets the value of the Inputs field in Feature_SdkV2.
func (m *Feature_SdkV2) SetInputs(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["inputs"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Inputs = types.ListValueMust(t, vs)
}

// GetSource returns the value of the Source field in Feature_SdkV2 as
// a DataSource_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Feature_SdkV2) GetSource(ctx context.Context) (DataSource_SdkV2, bool) {
	var e DataSource_SdkV2
	if m.Source.IsNull() || m.Source.IsUnknown() {
		return e, false
	}
	var v []DataSource_SdkV2
	d := m.Source.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSource sets the value of the Source field in Feature_SdkV2.
func (m *Feature_SdkV2) SetSource(ctx context.Context, v DataSource_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["source"]
	m.Source = types.ListValueMust(t, vs)
}

// GetTimeWindow returns the value of the TimeWindow field in Feature_SdkV2 as
// a TimeWindow_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Feature_SdkV2) GetTimeWindow(ctx context.Context) (TimeWindow_SdkV2, bool) {
	var e TimeWindow_SdkV2
	if m.TimeWindow.IsNull() || m.TimeWindow.IsUnknown() {
		return e, false
	}
	var v []TimeWindow_SdkV2
	d := m.TimeWindow.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTimeWindow sets the value of the TimeWindow field in Feature_SdkV2.
func (m *Feature_SdkV2) SetTimeWindow(ctx context.Context, v TimeWindow_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["time_window"]
	m.TimeWindow = types.ListValueMust(t, vs)
}

type FeatureLineage_SdkV2 struct {
	// List of feature specs that contain this feature.
	FeatureSpecs types.List `tfsdk:"feature_specs"`
	// List of Unity Catalog models that were trained on this feature.
	Models types.List `tfsdk:"models"`
	// List of online features that use this feature as source.
	OnlineFeatures types.List `tfsdk:"online_features"`
}

func (to *FeatureLineage_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FeatureLineage_SdkV2) {
	if !from.FeatureSpecs.IsNull() && !from.FeatureSpecs.IsUnknown() && to.FeatureSpecs.IsNull() && len(from.FeatureSpecs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FeatureSpecs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FeatureSpecs = from.FeatureSpecs
	}
	if !from.Models.IsNull() && !from.Models.IsUnknown() && to.Models.IsNull() && len(from.Models.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Models, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Models = from.Models
	}
	if !from.OnlineFeatures.IsNull() && !from.OnlineFeatures.IsUnknown() && to.OnlineFeatures.IsNull() && len(from.OnlineFeatures.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OnlineFeatures, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OnlineFeatures = from.OnlineFeatures
	}
}

func (to *FeatureLineage_SdkV2) SyncFieldsDuringRead(ctx context.Context, from FeatureLineage_SdkV2) {
	if !from.FeatureSpecs.IsNull() && !from.FeatureSpecs.IsUnknown() && to.FeatureSpecs.IsNull() && len(from.FeatureSpecs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FeatureSpecs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FeatureSpecs = from.FeatureSpecs
	}
	if !from.Models.IsNull() && !from.Models.IsUnknown() && to.Models.IsNull() && len(from.Models.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Models, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Models = from.Models
	}
	if !from.OnlineFeatures.IsNull() && !from.OnlineFeatures.IsUnknown() && to.OnlineFeatures.IsNull() && len(from.OnlineFeatures.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OnlineFeatures, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OnlineFeatures = from.OnlineFeatures
	}
}

func (m FeatureLineage_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["feature_specs"] = attrs["feature_specs"].SetOptional()
	attrs["models"] = attrs["models"].SetOptional()
	attrs["online_features"] = attrs["online_features"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FeatureLineage.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m FeatureLineage_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"feature_specs":   reflect.TypeOf(FeatureLineageFeatureSpec_SdkV2{}),
		"models":          reflect.TypeOf(FeatureLineageModel_SdkV2{}),
		"online_features": reflect.TypeOf(FeatureLineageOnlineFeature_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FeatureLineage_SdkV2
// only implements ToObjectValue() and Type().
func (m FeatureLineage_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_specs":   m.FeatureSpecs,
			"models":          m.Models,
			"online_features": m.OnlineFeatures,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FeatureLineage_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"feature_specs": basetypes.ListType{
				ElemType: FeatureLineageFeatureSpec_SdkV2{}.Type(ctx),
			},
			"models": basetypes.ListType{
				ElemType: FeatureLineageModel_SdkV2{}.Type(ctx),
			},
			"online_features": basetypes.ListType{
				ElemType: FeatureLineageOnlineFeature_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetFeatureSpecs returns the value of the FeatureSpecs field in FeatureLineage_SdkV2 as
// a slice of FeatureLineageFeatureSpec_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *FeatureLineage_SdkV2) GetFeatureSpecs(ctx context.Context) ([]FeatureLineageFeatureSpec_SdkV2, bool) {
	if m.FeatureSpecs.IsNull() || m.FeatureSpecs.IsUnknown() {
		return nil, false
	}
	var v []FeatureLineageFeatureSpec_SdkV2
	d := m.FeatureSpecs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFeatureSpecs sets the value of the FeatureSpecs field in FeatureLineage_SdkV2.
func (m *FeatureLineage_SdkV2) SetFeatureSpecs(ctx context.Context, v []FeatureLineageFeatureSpec_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["feature_specs"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.FeatureSpecs = types.ListValueMust(t, vs)
}

// GetModels returns the value of the Models field in FeatureLineage_SdkV2 as
// a slice of FeatureLineageModel_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *FeatureLineage_SdkV2) GetModels(ctx context.Context) ([]FeatureLineageModel_SdkV2, bool) {
	if m.Models.IsNull() || m.Models.IsUnknown() {
		return nil, false
	}
	var v []FeatureLineageModel_SdkV2
	d := m.Models.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModels sets the value of the Models field in FeatureLineage_SdkV2.
func (m *FeatureLineage_SdkV2) SetModels(ctx context.Context, v []FeatureLineageModel_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["models"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Models = types.ListValueMust(t, vs)
}

// GetOnlineFeatures returns the value of the OnlineFeatures field in FeatureLineage_SdkV2 as
// a slice of FeatureLineageOnlineFeature_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *FeatureLineage_SdkV2) GetOnlineFeatures(ctx context.Context) ([]FeatureLineageOnlineFeature_SdkV2, bool) {
	if m.OnlineFeatures.IsNull() || m.OnlineFeatures.IsUnknown() {
		return nil, false
	}
	var v []FeatureLineageOnlineFeature_SdkV2
	d := m.OnlineFeatures.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOnlineFeatures sets the value of the OnlineFeatures field in FeatureLineage_SdkV2.
func (m *FeatureLineage_SdkV2) SetOnlineFeatures(ctx context.Context, v []FeatureLineageOnlineFeature_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["online_features"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.OnlineFeatures = types.ListValueMust(t, vs)
}

type FeatureLineageFeatureSpec_SdkV2 struct {
	// The full name of the feature spec in Unity Catalog.
	Name types.String `tfsdk:"name"`
}

func (to *FeatureLineageFeatureSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FeatureLineageFeatureSpec_SdkV2) {
}

func (to *FeatureLineageFeatureSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, from FeatureLineageFeatureSpec_SdkV2) {
}

func (m FeatureLineageFeatureSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FeatureLineageFeatureSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m FeatureLineageFeatureSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FeatureLineageFeatureSpec_SdkV2
// only implements ToObjectValue() and Type().
func (m FeatureLineageFeatureSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FeatureLineageFeatureSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type FeatureLineageModel_SdkV2 struct {
	// The full name of the model in Unity Catalog.
	Name types.String `tfsdk:"name"`
	// The version of the model.
	Version types.Int64 `tfsdk:"version"`
}

func (to *FeatureLineageModel_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FeatureLineageModel_SdkV2) {
}

func (to *FeatureLineageModel_SdkV2) SyncFieldsDuringRead(ctx context.Context, from FeatureLineageModel_SdkV2) {
}

func (m FeatureLineageModel_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetOptional()
	attrs["version"] = attrs["version"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FeatureLineageModel.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m FeatureLineageModel_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FeatureLineageModel_SdkV2
// only implements ToObjectValue() and Type().
func (m FeatureLineageModel_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":    m.Name,
			"version": m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FeatureLineageModel_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":    types.StringType,
			"version": types.Int64Type,
		},
	}
}

type FeatureLineageOnlineFeature_SdkV2 struct {
	// The name of the online feature (column name).
	FeatureName types.String `tfsdk:"feature_name"`
	// The full name of the online table in Unity Catalog.
	TableName types.String `tfsdk:"table_name"`
}

func (to *FeatureLineageOnlineFeature_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FeatureLineageOnlineFeature_SdkV2) {
}

func (to *FeatureLineageOnlineFeature_SdkV2) SyncFieldsDuringRead(ctx context.Context, from FeatureLineageOnlineFeature_SdkV2) {
}

func (m FeatureLineageOnlineFeature_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["feature_name"] = attrs["feature_name"].SetOptional()
	attrs["table_name"] = attrs["table_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FeatureLineageOnlineFeature.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m FeatureLineageOnlineFeature_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FeatureLineageOnlineFeature_SdkV2
// only implements ToObjectValue() and Type().
func (m FeatureLineageOnlineFeature_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_name": m.FeatureName,
			"table_name":   m.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FeatureLineageOnlineFeature_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"feature_name": types.StringType,
			"table_name":   types.StringType,
		},
	}
}

// Feature list wrap all the features for a model version
type FeatureList_SdkV2 struct {
	Features types.List `tfsdk:"features"`
}

func (to *FeatureList_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FeatureList_SdkV2) {
	if !from.Features.IsNull() && !from.Features.IsUnknown() && to.Features.IsNull() && len(from.Features.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Features, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Features = from.Features
	}
}

func (to *FeatureList_SdkV2) SyncFieldsDuringRead(ctx context.Context, from FeatureList_SdkV2) {
	if !from.Features.IsNull() && !from.Features.IsUnknown() && to.Features.IsNull() && len(from.Features.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Features, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Features = from.Features
	}
}

func (m FeatureList_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["features"] = attrs["features"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FeatureList.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m FeatureList_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"features": reflect.TypeOf(LinkedFeature_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FeatureList_SdkV2
// only implements ToObjectValue() and Type().
func (m FeatureList_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"features": m.Features,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FeatureList_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"features": basetypes.ListType{
				ElemType: LinkedFeature_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetFeatures returns the value of the Features field in FeatureList_SdkV2 as
// a slice of LinkedFeature_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *FeatureList_SdkV2) GetFeatures(ctx context.Context) ([]LinkedFeature_SdkV2, bool) {
	if m.Features.IsNull() || m.Features.IsUnknown() {
		return nil, false
	}
	var v []LinkedFeature_SdkV2
	d := m.Features.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFeatures sets the value of the Features field in FeatureList_SdkV2.
func (m *FeatureList_SdkV2) SetFeatures(ctx context.Context, v []LinkedFeature_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["features"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Features = types.ListValueMust(t, vs)
}

// Represents a tag on a feature in a feature table.
type FeatureTag_SdkV2 struct {
	Key types.String `tfsdk:"key"`

	Value types.String `tfsdk:"value"`
}

func (to *FeatureTag_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FeatureTag_SdkV2) {
}

func (to *FeatureTag_SdkV2) SyncFieldsDuringRead(ctx context.Context, from FeatureTag_SdkV2) {
}

func (m FeatureTag_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetRequired()
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FeatureTag.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m FeatureTag_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FeatureTag_SdkV2
// only implements ToObjectValue() and Type().
func (m FeatureTag_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FeatureTag_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

// Metadata of a single artifact file or directory.
type FileInfo_SdkV2 struct {
	// The size in bytes of the file. Unset for directories.
	FileSize types.Int64 `tfsdk:"file_size"`
	// Whether the path is a directory.
	IsDir types.Bool `tfsdk:"is_dir"`
	// The path relative to the root artifact directory run.
	Path types.String `tfsdk:"path"`
}

func (to *FileInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FileInfo_SdkV2) {
}

func (to *FileInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from FileInfo_SdkV2) {
}

func (m FileInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["file_size"] = attrs["file_size"].SetOptional()
	attrs["is_dir"] = attrs["is_dir"].SetOptional()
	attrs["path"] = attrs["path"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FileInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m FileInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FileInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m FileInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_size": m.FileSize,
			"is_dir":    m.IsDir,
			"path":      m.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FileInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_size": types.Int64Type,
			"is_dir":    types.BoolType,
			"path":      types.StringType,
		},
	}
}

type FinalizeLoggedModelRequest_SdkV2 struct {
	// The ID of the logged model to finalize.
	ModelId types.String `tfsdk:"-"`
	// Whether or not the model is ready for use.
	// ``"LOGGED_MODEL_UPLOAD_FAILED"`` indicates that something went wrong when
	// logging the model weights / agent code.
	Status types.String `tfsdk:"status"`
}

func (to *FinalizeLoggedModelRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FinalizeLoggedModelRequest_SdkV2) {
}

func (to *FinalizeLoggedModelRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from FinalizeLoggedModelRequest_SdkV2) {
}

func (m FinalizeLoggedModelRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["status"] = attrs["status"].SetRequired()
	attrs["model_id"] = attrs["model_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FinalizeLoggedModelRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m FinalizeLoggedModelRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FinalizeLoggedModelRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m FinalizeLoggedModelRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_id": m.ModelId,
			"status":   m.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FinalizeLoggedModelRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_id": types.StringType,
			"status":   types.StringType,
		},
	}
}

type FinalizeLoggedModelResponse_SdkV2 struct {
	// The updated logged model.
	Model types.List `tfsdk:"model"`
}

func (to *FinalizeLoggedModelResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FinalizeLoggedModelResponse_SdkV2) {
	if !from.Model.IsNull() && !from.Model.IsUnknown() {
		if toModel, ok := to.GetModel(ctx); ok {
			if fromModel, ok := from.GetModel(ctx); ok {
				// Recursively sync the fields of Model
				toModel.SyncFieldsDuringCreateOrUpdate(ctx, fromModel)
				to.SetModel(ctx, toModel)
			}
		}
	}
}

func (to *FinalizeLoggedModelResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from FinalizeLoggedModelResponse_SdkV2) {
	if !from.Model.IsNull() && !from.Model.IsUnknown() {
		if toModel, ok := to.GetModel(ctx); ok {
			if fromModel, ok := from.GetModel(ctx); ok {
				toModel.SyncFieldsDuringRead(ctx, fromModel)
				to.SetModel(ctx, toModel)
			}
		}
	}
}

func (m FinalizeLoggedModelResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["model"] = attrs["model"].SetOptional()
	attrs["model"] = attrs["model"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FinalizeLoggedModelResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m FinalizeLoggedModelResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model": reflect.TypeOf(LoggedModel_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FinalizeLoggedModelResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m FinalizeLoggedModelResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model": m.Model,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FinalizeLoggedModelResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model": basetypes.ListType{
				ElemType: LoggedModel_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetModel returns the value of the Model field in FinalizeLoggedModelResponse_SdkV2 as
// a LoggedModel_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *FinalizeLoggedModelResponse_SdkV2) GetModel(ctx context.Context) (LoggedModel_SdkV2, bool) {
	var e LoggedModel_SdkV2
	if m.Model.IsNull() || m.Model.IsUnknown() {
		return e, false
	}
	var v []LoggedModel_SdkV2
	d := m.Model.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetModel sets the value of the Model field in FinalizeLoggedModelResponse_SdkV2.
func (m *FinalizeLoggedModelResponse_SdkV2) SetModel(ctx context.Context, v LoggedModel_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["model"]
	m.Model = types.ListValueMust(t, vs)
}

// Represents a forecasting experiment with its unique identifier, URL, and
// state.
type ForecastingExperiment_SdkV2 struct {
	// The unique ID for the forecasting experiment.
	ExperimentId types.String `tfsdk:"experiment_id"`
	// The URL to the forecasting experiment page.
	ExperimentPageUrl types.String `tfsdk:"experiment_page_url"`
	// The current state of the forecasting experiment.
	State types.String `tfsdk:"state"`
}

func (to *ForecastingExperiment_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ForecastingExperiment_SdkV2) {
}

func (to *ForecastingExperiment_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ForecastingExperiment_SdkV2) {
}

func (m ForecastingExperiment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["experiment_id"] = attrs["experiment_id"].SetOptional()
	attrs["experiment_page_url"] = attrs["experiment_page_url"].SetOptional()
	attrs["state"] = attrs["state"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ForecastingExperiment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ForecastingExperiment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ForecastingExperiment_SdkV2
// only implements ToObjectValue() and Type().
func (m ForecastingExperiment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id":       m.ExperimentId,
			"experiment_page_url": m.ExperimentPageUrl,
			"state":               m.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ForecastingExperiment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id":       types.StringType,
			"experiment_page_url": types.StringType,
			"state":               types.StringType,
		},
	}
}

type Function_SdkV2 struct {
	// Extra parameters for parameterized functions.
	ExtraParameters types.List `tfsdk:"extra_parameters"`
	// The type of the function.
	FunctionType types.String `tfsdk:"function_type"`
}

func (to *Function_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Function_SdkV2) {
	if !from.ExtraParameters.IsNull() && !from.ExtraParameters.IsUnknown() && to.ExtraParameters.IsNull() && len(from.ExtraParameters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ExtraParameters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ExtraParameters = from.ExtraParameters
	}
}

func (to *Function_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Function_SdkV2) {
	if !from.ExtraParameters.IsNull() && !from.ExtraParameters.IsUnknown() && to.ExtraParameters.IsNull() && len(from.ExtraParameters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ExtraParameters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ExtraParameters = from.ExtraParameters
	}
}

func (m Function_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["extra_parameters"] = attrs["extra_parameters"].SetOptional()
	attrs["function_type"] = attrs["function_type"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Function.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Function_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"extra_parameters": reflect.TypeOf(FunctionExtraParameter_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Function_SdkV2
// only implements ToObjectValue() and Type().
func (m Function_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"extra_parameters": m.ExtraParameters,
			"function_type":    m.FunctionType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Function_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"extra_parameters": basetypes.ListType{
				ElemType: FunctionExtraParameter_SdkV2{}.Type(ctx),
			},
			"function_type": types.StringType,
		},
	}
}

// GetExtraParameters returns the value of the ExtraParameters field in Function_SdkV2 as
// a slice of FunctionExtraParameter_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *Function_SdkV2) GetExtraParameters(ctx context.Context) ([]FunctionExtraParameter_SdkV2, bool) {
	if m.ExtraParameters.IsNull() || m.ExtraParameters.IsUnknown() {
		return nil, false
	}
	var v []FunctionExtraParameter_SdkV2
	d := m.ExtraParameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExtraParameters sets the value of the ExtraParameters field in Function_SdkV2.
func (m *Function_SdkV2) SetExtraParameters(ctx context.Context, v []FunctionExtraParameter_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["extra_parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ExtraParameters = types.ListValueMust(t, vs)
}

type FunctionExtraParameter_SdkV2 struct {
	// The name of the parameter.
	Key types.String `tfsdk:"key"`
	// The value of the parameter.
	Value types.String `tfsdk:"value"`
}

func (to *FunctionExtraParameter_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FunctionExtraParameter_SdkV2) {
}

func (to *FunctionExtraParameter_SdkV2) SyncFieldsDuringRead(ctx context.Context, from FunctionExtraParameter_SdkV2) {
}

func (m FunctionExtraParameter_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetRequired()
	attrs["value"] = attrs["value"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FunctionExtraParameter.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m FunctionExtraParameter_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FunctionExtraParameter_SdkV2
// only implements ToObjectValue() and Type().
func (m FunctionExtraParameter_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FunctionExtraParameter_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

type GetByNameRequest_SdkV2 struct {
	// Name of the associated experiment.
	ExperimentName types.String `tfsdk:"-"`
}

func (to *GetByNameRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetByNameRequest_SdkV2) {
}

func (to *GetByNameRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetByNameRequest_SdkV2) {
}

func (m GetByNameRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["experiment_name"] = attrs["experiment_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetByNameRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetByNameRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetByNameRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetByNameRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_name": m.ExperimentName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetByNameRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_name": types.StringType,
		},
	}
}

type GetExperimentByNameResponse_SdkV2 struct {
	// Experiment details.
	Experiment types.List `tfsdk:"experiment"`
}

func (to *GetExperimentByNameResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetExperimentByNameResponse_SdkV2) {
	if !from.Experiment.IsNull() && !from.Experiment.IsUnknown() {
		if toExperiment, ok := to.GetExperiment(ctx); ok {
			if fromExperiment, ok := from.GetExperiment(ctx); ok {
				// Recursively sync the fields of Experiment
				toExperiment.SyncFieldsDuringCreateOrUpdate(ctx, fromExperiment)
				to.SetExperiment(ctx, toExperiment)
			}
		}
	}
}

func (to *GetExperimentByNameResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetExperimentByNameResponse_SdkV2) {
	if !from.Experiment.IsNull() && !from.Experiment.IsUnknown() {
		if toExperiment, ok := to.GetExperiment(ctx); ok {
			if fromExperiment, ok := from.GetExperiment(ctx); ok {
				toExperiment.SyncFieldsDuringRead(ctx, fromExperiment)
				to.SetExperiment(ctx, toExperiment)
			}
		}
	}
}

func (m GetExperimentByNameResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["experiment"] = attrs["experiment"].SetOptional()
	attrs["experiment"] = attrs["experiment"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetExperimentByNameResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetExperimentByNameResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"experiment": reflect.TypeOf(Experiment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExperimentByNameResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetExperimentByNameResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment": m.Experiment,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetExperimentByNameResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment": basetypes.ListType{
				ElemType: Experiment_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetExperiment returns the value of the Experiment field in GetExperimentByNameResponse_SdkV2 as
// a Experiment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetExperimentByNameResponse_SdkV2) GetExperiment(ctx context.Context) (Experiment_SdkV2, bool) {
	var e Experiment_SdkV2
	if m.Experiment.IsNull() || m.Experiment.IsUnknown() {
		return e, false
	}
	var v []Experiment_SdkV2
	d := m.Experiment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetExperiment sets the value of the Experiment field in GetExperimentByNameResponse_SdkV2.
func (m *GetExperimentByNameResponse_SdkV2) SetExperiment(ctx context.Context, v Experiment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["experiment"]
	m.Experiment = types.ListValueMust(t, vs)
}

type GetExperimentPermissionLevelsRequest_SdkV2 struct {
	// The experiment for which to get or manage permissions.
	ExperimentId types.String `tfsdk:"-"`
}

func (to *GetExperimentPermissionLevelsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetExperimentPermissionLevelsRequest_SdkV2) {
}

func (to *GetExperimentPermissionLevelsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetExperimentPermissionLevelsRequest_SdkV2) {
}

func (m GetExperimentPermissionLevelsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["experiment_id"] = attrs["experiment_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetExperimentPermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetExperimentPermissionLevelsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExperimentPermissionLevelsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetExperimentPermissionLevelsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": m.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetExperimentPermissionLevelsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
		},
	}
}

type GetExperimentPermissionLevelsResponse_SdkV2 struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (to *GetExperimentPermissionLevelsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetExperimentPermissionLevelsResponse_SdkV2) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (to *GetExperimentPermissionLevelsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetExperimentPermissionLevelsResponse_SdkV2) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (m GetExperimentPermissionLevelsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission_levels"] = attrs["permission_levels"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetExperimentPermissionLevelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetExperimentPermissionLevelsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(ExperimentPermissionsDescription_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExperimentPermissionLevelsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetExperimentPermissionLevelsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": m.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetExperimentPermissionLevelsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: ExperimentPermissionsDescription_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetExperimentPermissionLevelsResponse_SdkV2 as
// a slice of ExperimentPermissionsDescription_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetExperimentPermissionLevelsResponse_SdkV2) GetPermissionLevels(ctx context.Context) ([]ExperimentPermissionsDescription_SdkV2, bool) {
	if m.PermissionLevels.IsNull() || m.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []ExperimentPermissionsDescription_SdkV2
	d := m.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetExperimentPermissionLevelsResponse_SdkV2.
func (m *GetExperimentPermissionLevelsResponse_SdkV2) SetPermissionLevels(ctx context.Context, v []ExperimentPermissionsDescription_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PermissionLevels = types.ListValueMust(t, vs)
}

type GetExperimentPermissionsRequest_SdkV2 struct {
	// The experiment for which to get or manage permissions.
	ExperimentId types.String `tfsdk:"-"`
}

func (to *GetExperimentPermissionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetExperimentPermissionsRequest_SdkV2) {
}

func (to *GetExperimentPermissionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetExperimentPermissionsRequest_SdkV2) {
}

func (m GetExperimentPermissionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["experiment_id"] = attrs["experiment_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetExperimentPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetExperimentPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExperimentPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetExperimentPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": m.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetExperimentPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
		},
	}
}

type GetExperimentRequest_SdkV2 struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"-"`
}

func (to *GetExperimentRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetExperimentRequest_SdkV2) {
}

func (to *GetExperimentRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetExperimentRequest_SdkV2) {
}

func (m GetExperimentRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["experiment_id"] = attrs["experiment_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetExperimentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetExperimentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExperimentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetExperimentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": m.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetExperimentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
		},
	}
}

type GetExperimentResponse_SdkV2 struct {
	// Experiment details.
	Experiment types.List `tfsdk:"experiment"`
}

func (to *GetExperimentResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetExperimentResponse_SdkV2) {
	if !from.Experiment.IsNull() && !from.Experiment.IsUnknown() {
		if toExperiment, ok := to.GetExperiment(ctx); ok {
			if fromExperiment, ok := from.GetExperiment(ctx); ok {
				// Recursively sync the fields of Experiment
				toExperiment.SyncFieldsDuringCreateOrUpdate(ctx, fromExperiment)
				to.SetExperiment(ctx, toExperiment)
			}
		}
	}
}

func (to *GetExperimentResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetExperimentResponse_SdkV2) {
	if !from.Experiment.IsNull() && !from.Experiment.IsUnknown() {
		if toExperiment, ok := to.GetExperiment(ctx); ok {
			if fromExperiment, ok := from.GetExperiment(ctx); ok {
				toExperiment.SyncFieldsDuringRead(ctx, fromExperiment)
				to.SetExperiment(ctx, toExperiment)
			}
		}
	}
}

func (m GetExperimentResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["experiment"] = attrs["experiment"].SetOptional()
	attrs["experiment"] = attrs["experiment"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetExperimentResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetExperimentResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"experiment": reflect.TypeOf(Experiment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExperimentResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetExperimentResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment": m.Experiment,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetExperimentResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment": basetypes.ListType{
				ElemType: Experiment_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetExperiment returns the value of the Experiment field in GetExperimentResponse_SdkV2 as
// a Experiment_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetExperimentResponse_SdkV2) GetExperiment(ctx context.Context) (Experiment_SdkV2, bool) {
	var e Experiment_SdkV2
	if m.Experiment.IsNull() || m.Experiment.IsUnknown() {
		return e, false
	}
	var v []Experiment_SdkV2
	d := m.Experiment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetExperiment sets the value of the Experiment field in GetExperimentResponse_SdkV2.
func (m *GetExperimentResponse_SdkV2) SetExperiment(ctx context.Context, v Experiment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["experiment"]
	m.Experiment = types.ListValueMust(t, vs)
}

type GetFeatureLineageRequest_SdkV2 struct {
	// The name of the feature.
	FeatureName types.String `tfsdk:"-"`
	// The full name of the feature table in Unity Catalog.
	TableName types.String `tfsdk:"-"`
}

func (to *GetFeatureLineageRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetFeatureLineageRequest_SdkV2) {
}

func (to *GetFeatureLineageRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetFeatureLineageRequest_SdkV2) {
}

func (m GetFeatureLineageRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["table_name"] = attrs["table_name"].SetRequired()
	attrs["feature_name"] = attrs["feature_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetFeatureLineageRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetFeatureLineageRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetFeatureLineageRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetFeatureLineageRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_name": m.FeatureName,
			"table_name":   m.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetFeatureLineageRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"feature_name": types.StringType,
			"table_name":   types.StringType,
		},
	}
}

type GetFeatureRequest_SdkV2 struct {
	// Name of the feature to get.
	FullName types.String `tfsdk:"-"`
}

func (to *GetFeatureRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetFeatureRequest_SdkV2) {
}

func (to *GetFeatureRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetFeatureRequest_SdkV2) {
}

func (m GetFeatureRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["full_name"] = attrs["full_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetFeatureRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetFeatureRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetFeatureRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetFeatureRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"full_name": m.FullName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetFeatureRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_name": types.StringType,
		},
	}
}

type GetFeatureTagRequest_SdkV2 struct {
	FeatureName types.String `tfsdk:"-"`

	Key types.String `tfsdk:"-"`

	TableName types.String `tfsdk:"-"`
}

func (to *GetFeatureTagRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetFeatureTagRequest_SdkV2) {
}

func (to *GetFeatureTagRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetFeatureTagRequest_SdkV2) {
}

func (m GetFeatureTagRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["table_name"] = attrs["table_name"].SetRequired()
	attrs["feature_name"] = attrs["feature_name"].SetRequired()
	attrs["key"] = attrs["key"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetFeatureTagRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetFeatureTagRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetFeatureTagRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetFeatureTagRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_name": m.FeatureName,
			"key":          m.Key,
			"table_name":   m.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetFeatureTagRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"feature_name": types.StringType,
			"key":          types.StringType,
			"table_name":   types.StringType,
		},
	}
}

type GetForecastingExperimentRequest_SdkV2 struct {
	// The unique ID of a forecasting experiment
	ExperimentId types.String `tfsdk:"-"`
}

func (to *GetForecastingExperimentRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetForecastingExperimentRequest_SdkV2) {
}

func (to *GetForecastingExperimentRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetForecastingExperimentRequest_SdkV2) {
}

func (m GetForecastingExperimentRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["experiment_id"] = attrs["experiment_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetForecastingExperimentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetForecastingExperimentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetForecastingExperimentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetForecastingExperimentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": m.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetForecastingExperimentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
		},
	}
}

type GetHistoryRequest_SdkV2 struct {
	// Maximum number of Metric records to return per paginated request. Default
	// is set to 25,000. If set higher than 25,000, a request Exception will be
	// raised.
	MaxResults types.Int64 `tfsdk:"-"`
	// Name of the metric.
	MetricKey types.String `tfsdk:"-"`
	// Token indicating the page of metric histories to fetch.
	PageToken types.String `tfsdk:"-"`
	// ID of the run from which to fetch metric values. Must be provided.
	RunId types.String `tfsdk:"-"`
	// [Deprecated, use `run_id` instead] ID of the run from which to fetch
	// metric values. This field will be removed in a future MLflow version.
	RunUuid types.String `tfsdk:"-"`
}

func (to *GetHistoryRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetHistoryRequest_SdkV2) {
}

func (to *GetHistoryRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetHistoryRequest_SdkV2) {
}

func (m GetHistoryRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["run_id"] = attrs["run_id"].SetOptional()
	attrs["run_uuid"] = attrs["run_uuid"].SetOptional()
	attrs["metric_key"] = attrs["metric_key"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["max_results"] = attrs["max_results"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetHistoryRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetHistoryRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetHistoryRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetHistoryRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results": m.MaxResults,
			"metric_key":  m.MetricKey,
			"page_token":  m.PageToken,
			"run_id":      m.RunId,
			"run_uuid":    m.RunUuid,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetHistoryRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_results": types.Int64Type,
			"metric_key":  types.StringType,
			"page_token":  types.StringType,
			"run_id":      types.StringType,
			"run_uuid":    types.StringType,
		},
	}
}

type GetLatestVersionsRequest_SdkV2 struct {
	// Registered model unique name identifier.
	Name types.String `tfsdk:"name"`
	// List of stages.
	Stages types.List `tfsdk:"stages"`
}

func (to *GetLatestVersionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetLatestVersionsRequest_SdkV2) {
	if !from.Stages.IsNull() && !from.Stages.IsUnknown() && to.Stages.IsNull() && len(from.Stages.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Stages, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Stages = from.Stages
	}
}

func (to *GetLatestVersionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetLatestVersionsRequest_SdkV2) {
	if !from.Stages.IsNull() && !from.Stages.IsUnknown() && to.Stages.IsNull() && len(from.Stages.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Stages, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Stages = from.Stages
	}
}

func (m GetLatestVersionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["stages"] = attrs["stages"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetLatestVersionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetLatestVersionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"stages": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLatestVersionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetLatestVersionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":   m.Name,
			"stages": m.Stages,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetLatestVersionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"stages": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetStages returns the value of the Stages field in GetLatestVersionsRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetLatestVersionsRequest_SdkV2) GetStages(ctx context.Context) ([]types.String, bool) {
	if m.Stages.IsNull() || m.Stages.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Stages.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStages sets the value of the Stages field in GetLatestVersionsRequest_SdkV2.
func (m *GetLatestVersionsRequest_SdkV2) SetStages(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["stages"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Stages = types.ListValueMust(t, vs)
}

type GetLatestVersionsResponse_SdkV2 struct {
	// Latest version models for each requests stage. Only return models with
	// current `READY` status. If no `stages` provided, returns the latest
	// version for each stage, including `"None"`.
	ModelVersions types.List `tfsdk:"model_versions"`
}

func (to *GetLatestVersionsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetLatestVersionsResponse_SdkV2) {
	if !from.ModelVersions.IsNull() && !from.ModelVersions.IsUnknown() && to.ModelVersions.IsNull() && len(from.ModelVersions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ModelVersions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ModelVersions = from.ModelVersions
	}
}

func (to *GetLatestVersionsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetLatestVersionsResponse_SdkV2) {
	if !from.ModelVersions.IsNull() && !from.ModelVersions.IsUnknown() && to.ModelVersions.IsNull() && len(from.ModelVersions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ModelVersions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ModelVersions = from.ModelVersions
	}
}

func (m GetLatestVersionsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["model_versions"] = attrs["model_versions"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetLatestVersionsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetLatestVersionsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model_versions": reflect.TypeOf(ModelVersion_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLatestVersionsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetLatestVersionsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_versions": m.ModelVersions,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetLatestVersionsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_versions": basetypes.ListType{
				ElemType: ModelVersion_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetModelVersions returns the value of the ModelVersions field in GetLatestVersionsResponse_SdkV2 as
// a slice of ModelVersion_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetLatestVersionsResponse_SdkV2) GetModelVersions(ctx context.Context) ([]ModelVersion_SdkV2, bool) {
	if m.ModelVersions.IsNull() || m.ModelVersions.IsUnknown() {
		return nil, false
	}
	var v []ModelVersion_SdkV2
	d := m.ModelVersions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModelVersions sets the value of the ModelVersions field in GetLatestVersionsResponse_SdkV2.
func (m *GetLatestVersionsResponse_SdkV2) SetModelVersions(ctx context.Context, v []ModelVersion_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["model_versions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ModelVersions = types.ListValueMust(t, vs)
}

type GetLoggedModelRequest_SdkV2 struct {
	// The ID of the logged model to retrieve.
	ModelId types.String `tfsdk:"-"`
}

func (to *GetLoggedModelRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetLoggedModelRequest_SdkV2) {
}

func (to *GetLoggedModelRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetLoggedModelRequest_SdkV2) {
}

func (m GetLoggedModelRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["model_id"] = attrs["model_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetLoggedModelRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetLoggedModelRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLoggedModelRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetLoggedModelRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_id": m.ModelId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetLoggedModelRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_id": types.StringType,
		},
	}
}

type GetLoggedModelResponse_SdkV2 struct {
	// The retrieved logged model.
	Model types.List `tfsdk:"model"`
}

func (to *GetLoggedModelResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetLoggedModelResponse_SdkV2) {
	if !from.Model.IsNull() && !from.Model.IsUnknown() {
		if toModel, ok := to.GetModel(ctx); ok {
			if fromModel, ok := from.GetModel(ctx); ok {
				// Recursively sync the fields of Model
				toModel.SyncFieldsDuringCreateOrUpdate(ctx, fromModel)
				to.SetModel(ctx, toModel)
			}
		}
	}
}

func (to *GetLoggedModelResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetLoggedModelResponse_SdkV2) {
	if !from.Model.IsNull() && !from.Model.IsUnknown() {
		if toModel, ok := to.GetModel(ctx); ok {
			if fromModel, ok := from.GetModel(ctx); ok {
				toModel.SyncFieldsDuringRead(ctx, fromModel)
				to.SetModel(ctx, toModel)
			}
		}
	}
}

func (m GetLoggedModelResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["model"] = attrs["model"].SetOptional()
	attrs["model"] = attrs["model"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetLoggedModelResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetLoggedModelResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model": reflect.TypeOf(LoggedModel_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLoggedModelResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetLoggedModelResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model": m.Model,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetLoggedModelResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model": basetypes.ListType{
				ElemType: LoggedModel_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetModel returns the value of the Model field in GetLoggedModelResponse_SdkV2 as
// a LoggedModel_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetLoggedModelResponse_SdkV2) GetModel(ctx context.Context) (LoggedModel_SdkV2, bool) {
	var e LoggedModel_SdkV2
	if m.Model.IsNull() || m.Model.IsUnknown() {
		return e, false
	}
	var v []LoggedModel_SdkV2
	d := m.Model.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetModel sets the value of the Model field in GetLoggedModelResponse_SdkV2.
func (m *GetLoggedModelResponse_SdkV2) SetModel(ctx context.Context, v LoggedModel_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["model"]
	m.Model = types.ListValueMust(t, vs)
}

type GetMetricHistoryResponse_SdkV2 struct {
	// All logged values for this metric if `max_results` is not specified in
	// the request or if the total count of metrics returned is less than the
	// service level pagination threshold. Otherwise, this is one page of
	// results.
	Metrics types.List `tfsdk:"metrics"`
	// A token that can be used to issue a query for the next page of metric
	// history values. A missing token indicates that no additional metrics are
	// available to fetch.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *GetMetricHistoryResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetMetricHistoryResponse_SdkV2) {
	if !from.Metrics.IsNull() && !from.Metrics.IsUnknown() && to.Metrics.IsNull() && len(from.Metrics.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Metrics, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Metrics = from.Metrics
	}
}

func (to *GetMetricHistoryResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetMetricHistoryResponse_SdkV2) {
	if !from.Metrics.IsNull() && !from.Metrics.IsUnknown() && to.Metrics.IsNull() && len(from.Metrics.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Metrics, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Metrics = from.Metrics
	}
}

func (m GetMetricHistoryResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["metrics"] = attrs["metrics"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetMetricHistoryResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetMetricHistoryResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metrics": reflect.TypeOf(Metric_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetMetricHistoryResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetMetricHistoryResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metrics":         m.Metrics,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetMetricHistoryResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metrics": basetypes.ListType{
				ElemType: Metric_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetMetrics returns the value of the Metrics field in GetMetricHistoryResponse_SdkV2 as
// a slice of Metric_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetMetricHistoryResponse_SdkV2) GetMetrics(ctx context.Context) ([]Metric_SdkV2, bool) {
	if m.Metrics.IsNull() || m.Metrics.IsUnknown() {
		return nil, false
	}
	var v []Metric_SdkV2
	d := m.Metrics.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMetrics sets the value of the Metrics field in GetMetricHistoryResponse_SdkV2.
func (m *GetMetricHistoryResponse_SdkV2) SetMetrics(ctx context.Context, v []Metric_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["metrics"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Metrics = types.ListValueMust(t, vs)
}

type GetModelRequest_SdkV2 struct {
	// Registered model unique name identifier.
	Name types.String `tfsdk:"-"`
}

func (to *GetModelRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetModelRequest_SdkV2) {
}

func (to *GetModelRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetModelRequest_SdkV2) {
}

func (m GetModelRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetModelRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetModelRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetModelRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetModelRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetModelRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetModelResponse_SdkV2 struct {
	RegisteredModelDatabricks types.List `tfsdk:"registered_model_databricks"`
}

func (to *GetModelResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetModelResponse_SdkV2) {
	if !from.RegisteredModelDatabricks.IsNull() && !from.RegisteredModelDatabricks.IsUnknown() {
		if toRegisteredModelDatabricks, ok := to.GetRegisteredModelDatabricks(ctx); ok {
			if fromRegisteredModelDatabricks, ok := from.GetRegisteredModelDatabricks(ctx); ok {
				// Recursively sync the fields of RegisteredModelDatabricks
				toRegisteredModelDatabricks.SyncFieldsDuringCreateOrUpdate(ctx, fromRegisteredModelDatabricks)
				to.SetRegisteredModelDatabricks(ctx, toRegisteredModelDatabricks)
			}
		}
	}
}

func (to *GetModelResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetModelResponse_SdkV2) {
	if !from.RegisteredModelDatabricks.IsNull() && !from.RegisteredModelDatabricks.IsUnknown() {
		if toRegisteredModelDatabricks, ok := to.GetRegisteredModelDatabricks(ctx); ok {
			if fromRegisteredModelDatabricks, ok := from.GetRegisteredModelDatabricks(ctx); ok {
				toRegisteredModelDatabricks.SyncFieldsDuringRead(ctx, fromRegisteredModelDatabricks)
				to.SetRegisteredModelDatabricks(ctx, toRegisteredModelDatabricks)
			}
		}
	}
}

func (m GetModelResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["registered_model_databricks"] = attrs["registered_model_databricks"].SetOptional()
	attrs["registered_model_databricks"] = attrs["registered_model_databricks"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetModelResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetModelResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"registered_model_databricks": reflect.TypeOf(ModelDatabricks_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetModelResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetModelResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"registered_model_databricks": m.RegisteredModelDatabricks,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetModelResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"registered_model_databricks": basetypes.ListType{
				ElemType: ModelDatabricks_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetRegisteredModelDatabricks returns the value of the RegisteredModelDatabricks field in GetModelResponse_SdkV2 as
// a ModelDatabricks_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetModelResponse_SdkV2) GetRegisteredModelDatabricks(ctx context.Context) (ModelDatabricks_SdkV2, bool) {
	var e ModelDatabricks_SdkV2
	if m.RegisteredModelDatabricks.IsNull() || m.RegisteredModelDatabricks.IsUnknown() {
		return e, false
	}
	var v []ModelDatabricks_SdkV2
	d := m.RegisteredModelDatabricks.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRegisteredModelDatabricks sets the value of the RegisteredModelDatabricks field in GetModelResponse_SdkV2.
func (m *GetModelResponse_SdkV2) SetRegisteredModelDatabricks(ctx context.Context, v ModelDatabricks_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["registered_model_databricks"]
	m.RegisteredModelDatabricks = types.ListValueMust(t, vs)
}

type GetModelVersionDownloadUriRequest_SdkV2 struct {
	// Name of the registered model
	Name types.String `tfsdk:"-"`
	// Model version number
	Version types.String `tfsdk:"-"`
}

func (to *GetModelVersionDownloadUriRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetModelVersionDownloadUriRequest_SdkV2) {
}

func (to *GetModelVersionDownloadUriRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetModelVersionDownloadUriRequest_SdkV2) {
}

func (m GetModelVersionDownloadUriRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["version"] = attrs["version"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetModelVersionDownloadUriRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetModelVersionDownloadUriRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetModelVersionDownloadUriRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetModelVersionDownloadUriRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":    m.Name,
			"version": m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetModelVersionDownloadUriRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":    types.StringType,
			"version": types.StringType,
		},
	}
}

type GetModelVersionDownloadUriResponse_SdkV2 struct {
	// URI corresponding to where artifacts for this model version are stored.
	ArtifactUri types.String `tfsdk:"artifact_uri"`
}

func (to *GetModelVersionDownloadUriResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetModelVersionDownloadUriResponse_SdkV2) {
}

func (to *GetModelVersionDownloadUriResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetModelVersionDownloadUriResponse_SdkV2) {
}

func (m GetModelVersionDownloadUriResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["artifact_uri"] = attrs["artifact_uri"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetModelVersionDownloadUriResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetModelVersionDownloadUriResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetModelVersionDownloadUriResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetModelVersionDownloadUriResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"artifact_uri": m.ArtifactUri,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetModelVersionDownloadUriResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"artifact_uri": types.StringType,
		},
	}
}

type GetModelVersionRequest_SdkV2 struct {
	// Name of the registered model
	Name types.String `tfsdk:"-"`
	// Model version number
	Version types.String `tfsdk:"-"`
}

func (to *GetModelVersionRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetModelVersionRequest_SdkV2) {
}

func (to *GetModelVersionRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetModelVersionRequest_SdkV2) {
}

func (m GetModelVersionRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["version"] = attrs["version"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetModelVersionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetModelVersionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetModelVersionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetModelVersionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":    m.Name,
			"version": m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetModelVersionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":    types.StringType,
			"version": types.StringType,
		},
	}
}

type GetModelVersionResponse_SdkV2 struct {
	ModelVersion types.List `tfsdk:"model_version"`
}

func (to *GetModelVersionResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetModelVersionResponse_SdkV2) {
	if !from.ModelVersion.IsNull() && !from.ModelVersion.IsUnknown() {
		if toModelVersion, ok := to.GetModelVersion(ctx); ok {
			if fromModelVersion, ok := from.GetModelVersion(ctx); ok {
				// Recursively sync the fields of ModelVersion
				toModelVersion.SyncFieldsDuringCreateOrUpdate(ctx, fromModelVersion)
				to.SetModelVersion(ctx, toModelVersion)
			}
		}
	}
}

func (to *GetModelVersionResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetModelVersionResponse_SdkV2) {
	if !from.ModelVersion.IsNull() && !from.ModelVersion.IsUnknown() {
		if toModelVersion, ok := to.GetModelVersion(ctx); ok {
			if fromModelVersion, ok := from.GetModelVersion(ctx); ok {
				toModelVersion.SyncFieldsDuringRead(ctx, fromModelVersion)
				to.SetModelVersion(ctx, toModelVersion)
			}
		}
	}
}

func (m GetModelVersionResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["model_version"] = attrs["model_version"].SetOptional()
	attrs["model_version"] = attrs["model_version"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetModelVersionResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetModelVersionResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model_version": reflect.TypeOf(ModelVersion_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetModelVersionResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetModelVersionResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_version": m.ModelVersion,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetModelVersionResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_version": basetypes.ListType{
				ElemType: ModelVersion_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetModelVersion returns the value of the ModelVersion field in GetModelVersionResponse_SdkV2 as
// a ModelVersion_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetModelVersionResponse_SdkV2) GetModelVersion(ctx context.Context) (ModelVersion_SdkV2, bool) {
	var e ModelVersion_SdkV2
	if m.ModelVersion.IsNull() || m.ModelVersion.IsUnknown() {
		return e, false
	}
	var v []ModelVersion_SdkV2
	d := m.ModelVersion.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetModelVersion sets the value of the ModelVersion field in GetModelVersionResponse_SdkV2.
func (m *GetModelVersionResponse_SdkV2) SetModelVersion(ctx context.Context, v ModelVersion_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["model_version"]
	m.ModelVersion = types.ListValueMust(t, vs)
}

type GetOnlineStoreRequest_SdkV2 struct {
	// Name of the online store to get.
	Name types.String `tfsdk:"-"`
}

func (to *GetOnlineStoreRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetOnlineStoreRequest_SdkV2) {
}

func (to *GetOnlineStoreRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetOnlineStoreRequest_SdkV2) {
}

func (m GetOnlineStoreRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetOnlineStoreRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetOnlineStoreRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetOnlineStoreRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetOnlineStoreRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetOnlineStoreRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetRegisteredModelPermissionLevelsRequest_SdkV2 struct {
	// The registered model for which to get or manage permissions.
	RegisteredModelId types.String `tfsdk:"-"`
}

func (to *GetRegisteredModelPermissionLevelsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetRegisteredModelPermissionLevelsRequest_SdkV2) {
}

func (to *GetRegisteredModelPermissionLevelsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetRegisteredModelPermissionLevelsRequest_SdkV2) {
}

func (m GetRegisteredModelPermissionLevelsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["registered_model_id"] = attrs["registered_model_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRegisteredModelPermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetRegisteredModelPermissionLevelsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRegisteredModelPermissionLevelsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetRegisteredModelPermissionLevelsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"registered_model_id": m.RegisteredModelId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetRegisteredModelPermissionLevelsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"registered_model_id": types.StringType,
		},
	}
}

type GetRegisteredModelPermissionLevelsResponse_SdkV2 struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (to *GetRegisteredModelPermissionLevelsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetRegisteredModelPermissionLevelsResponse_SdkV2) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (to *GetRegisteredModelPermissionLevelsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetRegisteredModelPermissionLevelsResponse_SdkV2) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (m GetRegisteredModelPermissionLevelsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission_levels"] = attrs["permission_levels"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRegisteredModelPermissionLevelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetRegisteredModelPermissionLevelsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(RegisteredModelPermissionsDescription_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRegisteredModelPermissionLevelsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetRegisteredModelPermissionLevelsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": m.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetRegisteredModelPermissionLevelsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: RegisteredModelPermissionsDescription_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetRegisteredModelPermissionLevelsResponse_SdkV2 as
// a slice of RegisteredModelPermissionsDescription_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetRegisteredModelPermissionLevelsResponse_SdkV2) GetPermissionLevels(ctx context.Context) ([]RegisteredModelPermissionsDescription_SdkV2, bool) {
	if m.PermissionLevels.IsNull() || m.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []RegisteredModelPermissionsDescription_SdkV2
	d := m.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetRegisteredModelPermissionLevelsResponse_SdkV2.
func (m *GetRegisteredModelPermissionLevelsResponse_SdkV2) SetPermissionLevels(ctx context.Context, v []RegisteredModelPermissionsDescription_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PermissionLevels = types.ListValueMust(t, vs)
}

type GetRegisteredModelPermissionsRequest_SdkV2 struct {
	// The registered model for which to get or manage permissions.
	RegisteredModelId types.String `tfsdk:"-"`
}

func (to *GetRegisteredModelPermissionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetRegisteredModelPermissionsRequest_SdkV2) {
}

func (to *GetRegisteredModelPermissionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetRegisteredModelPermissionsRequest_SdkV2) {
}

func (m GetRegisteredModelPermissionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["registered_model_id"] = attrs["registered_model_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRegisteredModelPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetRegisteredModelPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRegisteredModelPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetRegisteredModelPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"registered_model_id": m.RegisteredModelId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetRegisteredModelPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"registered_model_id": types.StringType,
		},
	}
}

type GetRunRequest_SdkV2 struct {
	// ID of the run to fetch. Must be provided.
	RunId types.String `tfsdk:"-"`
	// [Deprecated, use `run_id` instead] ID of the run to fetch. This field
	// will be removed in a future MLflow version.
	RunUuid types.String `tfsdk:"-"`
}

func (to *GetRunRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetRunRequest_SdkV2) {
}

func (to *GetRunRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetRunRequest_SdkV2) {
}

func (m GetRunRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["run_id"] = attrs["run_id"].SetRequired()
	attrs["run_uuid"] = attrs["run_uuid"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRunRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetRunRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRunRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetRunRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run_id":   m.RunId,
			"run_uuid": m.RunUuid,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetRunRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run_id":   types.StringType,
			"run_uuid": types.StringType,
		},
	}
}

type GetRunResponse_SdkV2 struct {
	// Run metadata (name, start time, etc) and data (metrics, params, and
	// tags).
	Run types.List `tfsdk:"run"`
}

func (to *GetRunResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetRunResponse_SdkV2) {
	if !from.Run.IsNull() && !from.Run.IsUnknown() {
		if toRun, ok := to.GetRun(ctx); ok {
			if fromRun, ok := from.GetRun(ctx); ok {
				// Recursively sync the fields of Run
				toRun.SyncFieldsDuringCreateOrUpdate(ctx, fromRun)
				to.SetRun(ctx, toRun)
			}
		}
	}
}

func (to *GetRunResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetRunResponse_SdkV2) {
	if !from.Run.IsNull() && !from.Run.IsUnknown() {
		if toRun, ok := to.GetRun(ctx); ok {
			if fromRun, ok := from.GetRun(ctx); ok {
				toRun.SyncFieldsDuringRead(ctx, fromRun)
				to.SetRun(ctx, toRun)
			}
		}
	}
}

func (m GetRunResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["run"] = attrs["run"].SetOptional()
	attrs["run"] = attrs["run"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRunResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetRunResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"run": reflect.TypeOf(Run_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRunResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetRunResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run": m.Run,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetRunResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run": basetypes.ListType{
				ElemType: Run_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetRun returns the value of the Run field in GetRunResponse_SdkV2 as
// a Run_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetRunResponse_SdkV2) GetRun(ctx context.Context) (Run_SdkV2, bool) {
	var e Run_SdkV2
	if m.Run.IsNull() || m.Run.IsUnknown() {
		return e, false
	}
	var v []Run_SdkV2
	d := m.Run.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRun sets the value of the Run field in GetRunResponse_SdkV2.
func (m *GetRunResponse_SdkV2) SetRun(ctx context.Context, v Run_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["run"]
	m.Run = types.ListValueMust(t, vs)
}

type HttpUrlSpec_SdkV2 struct {
	// Value of the authorization header that should be sent in the request sent
	// by the wehbook. It should be of the form `"<auth type> <credentials>"`.
	// If set to an empty string, no authorization header will be included in
	// the request.
	Authorization types.String `tfsdk:"authorization"`
	// Enable/disable SSL certificate validation. Default is true. For
	// self-signed certificates, this field must be false AND the destination
	// server must disable certificate validation as well. For security
	// purposes, it is encouraged to perform secret validation with the
	// HMAC-encoded portion of the payload and acknowledge the risk associated
	// with disabling hostname validation whereby it becomes more likely that
	// requests can be maliciously routed to an unintended host.
	EnableSslVerification types.Bool `tfsdk:"enable_ssl_verification"`
	// Shared secret required for HMAC encoding payload. The HMAC-encoded
	// payload will be sent in the header as: { "X-Databricks-Signature":
	// $encoded_payload }.
	Secret types.String `tfsdk:"secret"`
	// External HTTPS URL called on event trigger (by using a POST request).
	Url types.String `tfsdk:"url"`
}

func (to *HttpUrlSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from HttpUrlSpec_SdkV2) {
}

func (to *HttpUrlSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, from HttpUrlSpec_SdkV2) {
}

func (m HttpUrlSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["authorization"] = attrs["authorization"].SetOptional()
	attrs["enable_ssl_verification"] = attrs["enable_ssl_verification"].SetOptional()
	attrs["secret"] = attrs["secret"].SetOptional()
	attrs["url"] = attrs["url"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in HttpUrlSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m HttpUrlSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, HttpUrlSpec_SdkV2
// only implements ToObjectValue() and Type().
func (m HttpUrlSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"authorization":           m.Authorization,
			"enable_ssl_verification": m.EnableSslVerification,
			"secret":                  m.Secret,
			"url":                     m.Url,
		})
}

// Type implements basetypes.ObjectValuable.
func (m HttpUrlSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"authorization":           types.StringType,
			"enable_ssl_verification": types.BoolType,
			"secret":                  types.StringType,
			"url":                     types.StringType,
		},
	}
}

type HttpUrlSpecWithoutSecret_SdkV2 struct {
	// Enable/disable SSL certificate validation. Default is true. For
	// self-signed certificates, this field must be false AND the destination
	// server must disable certificate validation as well. For security
	// purposes, it is encouraged to perform secret validation with the
	// HMAC-encoded portion of the payload and acknowledge the risk associated
	// with disabling hostname validation whereby it becomes more likely that
	// requests can be maliciously routed to an unintended host.
	EnableSslVerification types.Bool `tfsdk:"enable_ssl_verification"`
	// External HTTPS URL called on event trigger (by using a POST request).
	Url types.String `tfsdk:"url"`
}

func (to *HttpUrlSpecWithoutSecret_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from HttpUrlSpecWithoutSecret_SdkV2) {
}

func (to *HttpUrlSpecWithoutSecret_SdkV2) SyncFieldsDuringRead(ctx context.Context, from HttpUrlSpecWithoutSecret_SdkV2) {
}

func (m HttpUrlSpecWithoutSecret_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["enable_ssl_verification"] = attrs["enable_ssl_verification"].SetOptional()
	attrs["url"] = attrs["url"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in HttpUrlSpecWithoutSecret.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m HttpUrlSpecWithoutSecret_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, HttpUrlSpecWithoutSecret_SdkV2
// only implements ToObjectValue() and Type().
func (m HttpUrlSpecWithoutSecret_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enable_ssl_verification": m.EnableSslVerification,
			"url":                     m.Url,
		})
}

// Type implements basetypes.ObjectValuable.
func (m HttpUrlSpecWithoutSecret_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enable_ssl_verification": types.BoolType,
			"url":                     types.StringType,
		},
	}
}

// Tag for a dataset input.
type InputTag_SdkV2 struct {
	// The tag key.
	Key types.String `tfsdk:"key"`
	// The tag value.
	Value types.String `tfsdk:"value"`
}

func (to *InputTag_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from InputTag_SdkV2) {
}

func (to *InputTag_SdkV2) SyncFieldsDuringRead(ctx context.Context, from InputTag_SdkV2) {
}

func (m InputTag_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetRequired()
	attrs["value"] = attrs["value"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in InputTag.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m InputTag_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InputTag_SdkV2
// only implements ToObjectValue() and Type().
func (m InputTag_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m InputTag_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

type JobSpec_SdkV2 struct {
	// The personal access token used to authorize webhook's job runs.
	AccessToken types.String `tfsdk:"access_token"`
	// ID of the job that the webhook runs.
	JobId types.String `tfsdk:"job_id"`
	// URL of the workspace containing the job that this webhook runs. If not
	// specified, the job’s workspace URL is assumed to be the same as the
	// workspace where the webhook is created.
	WorkspaceUrl types.String `tfsdk:"workspace_url"`
}

func (to *JobSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from JobSpec_SdkV2) {
}

func (to *JobSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, from JobSpec_SdkV2) {
}

func (m JobSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_token"] = attrs["access_token"].SetRequired()
	attrs["job_id"] = attrs["job_id"].SetRequired()
	attrs["workspace_url"] = attrs["workspace_url"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m JobSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, JobSpec_SdkV2
// only implements ToObjectValue() and Type().
func (m JobSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_token":  m.AccessToken,
			"job_id":        m.JobId,
			"workspace_url": m.WorkspaceUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (m JobSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_token":  types.StringType,
			"job_id":        types.StringType,
			"workspace_url": types.StringType,
		},
	}
}

type JobSpecWithoutSecret_SdkV2 struct {
	// ID of the job that the webhook runs.
	JobId types.String `tfsdk:"job_id"`
	// URL of the workspace containing the job that this webhook runs. If not
	// specified, the job’s workspace URL is assumed to be the same as the
	// workspace where the webhook is created.
	WorkspaceUrl types.String `tfsdk:"workspace_url"`
}

func (to *JobSpecWithoutSecret_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from JobSpecWithoutSecret_SdkV2) {
}

func (to *JobSpecWithoutSecret_SdkV2) SyncFieldsDuringRead(ctx context.Context, from JobSpecWithoutSecret_SdkV2) {
}

func (m JobSpecWithoutSecret_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["job_id"] = attrs["job_id"].SetOptional()
	attrs["workspace_url"] = attrs["workspace_url"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JobSpecWithoutSecret.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m JobSpecWithoutSecret_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, JobSpecWithoutSecret_SdkV2
// only implements ToObjectValue() and Type().
func (m JobSpecWithoutSecret_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"job_id":        m.JobId,
			"workspace_url": m.WorkspaceUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (m JobSpecWithoutSecret_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"job_id":        types.StringType,
			"workspace_url": types.StringType,
		},
	}
}

// Feature for model version. ([ML-57150] Renamed from Feature to LinkedFeature)
type LinkedFeature_SdkV2 struct {
	// Feature name
	FeatureName types.String `tfsdk:"feature_name"`
	// Feature table id
	FeatureTableId types.String `tfsdk:"feature_table_id"`
	// Feature table name
	FeatureTableName types.String `tfsdk:"feature_table_name"`
}

func (to *LinkedFeature_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LinkedFeature_SdkV2) {
}

func (to *LinkedFeature_SdkV2) SyncFieldsDuringRead(ctx context.Context, from LinkedFeature_SdkV2) {
}

func (m LinkedFeature_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["feature_name"] = attrs["feature_name"].SetOptional()
	attrs["feature_table_id"] = attrs["feature_table_id"].SetOptional()
	attrs["feature_table_name"] = attrs["feature_table_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LinkedFeature.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LinkedFeature_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LinkedFeature_SdkV2
// only implements ToObjectValue() and Type().
func (m LinkedFeature_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_name":       m.FeatureName,
			"feature_table_id":   m.FeatureTableId,
			"feature_table_name": m.FeatureTableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LinkedFeature_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"feature_name":       types.StringType,
			"feature_table_id":   types.StringType,
			"feature_table_name": types.StringType,
		},
	}
}

type ListArtifactsRequest_SdkV2 struct {
	// The token indicating the page of artifact results to fetch. `page_token`
	// is not supported when listing artifacts in UC Volumes. A maximum of 1000
	// artifacts will be retrieved for UC Volumes. Please call
	// `/api/2.0/fs/directories{directory_path}` for listing artifacts in UC
	// Volumes, which supports pagination. See [List directory contents | Files
	// API](/api/workspace/files/listdirectorycontents).
	PageToken types.String `tfsdk:"-"`
	// Filter artifacts matching this path (a relative path from the root
	// artifact directory).
	Path types.String `tfsdk:"-"`
	// ID of the run whose artifacts to list. Must be provided.
	RunId types.String `tfsdk:"-"`
	// [Deprecated, use `run_id` instead] ID of the run whose artifacts to list.
	// This field will be removed in a future MLflow version.
	RunUuid types.String `tfsdk:"-"`
}

func (to *ListArtifactsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListArtifactsRequest_SdkV2) {
}

func (to *ListArtifactsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListArtifactsRequest_SdkV2) {
}

func (m ListArtifactsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["run_id"] = attrs["run_id"].SetOptional()
	attrs["run_uuid"] = attrs["run_uuid"].SetOptional()
	attrs["path"] = attrs["path"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListArtifactsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListArtifactsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListArtifactsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListArtifactsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_token": m.PageToken,
			"path":       m.Path,
			"run_id":     m.RunId,
			"run_uuid":   m.RunUuid,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListArtifactsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_token": types.StringType,
			"path":       types.StringType,
			"run_id":     types.StringType,
			"run_uuid":   types.StringType,
		},
	}
}

type ListArtifactsResponse_SdkV2 struct {
	// The file location and metadata for artifacts.
	Files types.List `tfsdk:"files"`
	// The token that can be used to retrieve the next page of artifact results.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// The root artifact directory for the run.
	RootUri types.String `tfsdk:"root_uri"`
}

func (to *ListArtifactsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListArtifactsResponse_SdkV2) {
	if !from.Files.IsNull() && !from.Files.IsUnknown() && to.Files.IsNull() && len(from.Files.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Files, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Files = from.Files
	}
}

func (to *ListArtifactsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListArtifactsResponse_SdkV2) {
	if !from.Files.IsNull() && !from.Files.IsUnknown() && to.Files.IsNull() && len(from.Files.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Files, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Files = from.Files
	}
}

func (m ListArtifactsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["files"] = attrs["files"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["root_uri"] = attrs["root_uri"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListArtifactsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListArtifactsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"files": reflect.TypeOf(FileInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListArtifactsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListArtifactsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"files":           m.Files,
			"next_page_token": m.NextPageToken,
			"root_uri":        m.RootUri,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListArtifactsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"files": basetypes.ListType{
				ElemType: FileInfo_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
			"root_uri":        types.StringType,
		},
	}
}

// GetFiles returns the value of the Files field in ListArtifactsResponse_SdkV2 as
// a slice of FileInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListArtifactsResponse_SdkV2) GetFiles(ctx context.Context) ([]FileInfo_SdkV2, bool) {
	if m.Files.IsNull() || m.Files.IsUnknown() {
		return nil, false
	}
	var v []FileInfo_SdkV2
	d := m.Files.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFiles sets the value of the Files field in ListArtifactsResponse_SdkV2.
func (m *ListArtifactsResponse_SdkV2) SetFiles(ctx context.Context, v []FileInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["files"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Files = types.ListValueMust(t, vs)
}

type ListExperimentsRequest_SdkV2 struct {
	// Maximum number of experiments desired. If `max_results` is unspecified,
	// return all experiments. If `max_results` is too large, it'll be
	// automatically capped at 1000. Callers of this endpoint are encouraged to
	// pass max_results explicitly and leverage page_token to iterate through
	// experiments.
	MaxResults types.Int64 `tfsdk:"-"`
	// Token indicating the page of experiments to fetch
	PageToken types.String `tfsdk:"-"`
	// Qualifier for type of experiments to be returned. If unspecified, return
	// only active experiments.
	ViewType types.String `tfsdk:"-"`
}

func (to *ListExperimentsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListExperimentsRequest_SdkV2) {
}

func (to *ListExperimentsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListExperimentsRequest_SdkV2) {
}

func (m ListExperimentsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["view_type"] = attrs["view_type"].SetOptional()
	attrs["max_results"] = attrs["max_results"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListExperimentsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListExperimentsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExperimentsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListExperimentsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results": m.MaxResults,
			"page_token":  m.PageToken,
			"view_type":   m.ViewType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListExperimentsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_results": types.Int64Type,
			"page_token":  types.StringType,
			"view_type":   types.StringType,
		},
	}
}

type ListExperimentsResponse_SdkV2 struct {
	// Paginated Experiments beginning with the first item on the requested
	// page.
	Experiments types.List `tfsdk:"experiments"`
	// Token that can be used to retrieve the next page of experiments. Empty
	// token means no more experiment is available for retrieval.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListExperimentsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListExperimentsResponse_SdkV2) {
	if !from.Experiments.IsNull() && !from.Experiments.IsUnknown() && to.Experiments.IsNull() && len(from.Experiments.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Experiments, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Experiments = from.Experiments
	}
}

func (to *ListExperimentsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListExperimentsResponse_SdkV2) {
	if !from.Experiments.IsNull() && !from.Experiments.IsUnknown() && to.Experiments.IsNull() && len(from.Experiments.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Experiments, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Experiments = from.Experiments
	}
}

func (m ListExperimentsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["experiments"] = attrs["experiments"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListExperimentsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListExperimentsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"experiments": reflect.TypeOf(Experiment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExperimentsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListExperimentsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiments":     m.Experiments,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListExperimentsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiments": basetypes.ListType{
				ElemType: Experiment_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetExperiments returns the value of the Experiments field in ListExperimentsResponse_SdkV2 as
// a slice of Experiment_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListExperimentsResponse_SdkV2) GetExperiments(ctx context.Context) ([]Experiment_SdkV2, bool) {
	if m.Experiments.IsNull() || m.Experiments.IsUnknown() {
		return nil, false
	}
	var v []Experiment_SdkV2
	d := m.Experiments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExperiments sets the value of the Experiments field in ListExperimentsResponse_SdkV2.
func (m *ListExperimentsResponse_SdkV2) SetExperiments(ctx context.Context, v []Experiment_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["experiments"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Experiments = types.ListValueMust(t, vs)
}

type ListFeatureTagsRequest_SdkV2 struct {
	FeatureName types.String `tfsdk:"-"`
	// The maximum number of results to return.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page based on a previous query.
	PageToken types.String `tfsdk:"-"`

	TableName types.String `tfsdk:"-"`
}

func (to *ListFeatureTagsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListFeatureTagsRequest_SdkV2) {
}

func (to *ListFeatureTagsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListFeatureTagsRequest_SdkV2) {
}

func (m ListFeatureTagsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["table_name"] = attrs["table_name"].SetRequired()
	attrs["feature_name"] = attrs["feature_name"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListFeatureTagsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListFeatureTagsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFeatureTagsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListFeatureTagsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_name": m.FeatureName,
			"page_size":    m.PageSize,
			"page_token":   m.PageToken,
			"table_name":   m.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListFeatureTagsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"feature_name": types.StringType,
			"page_size":    types.Int64Type,
			"page_token":   types.StringType,
			"table_name":   types.StringType,
		},
	}
}

// Response message for ListFeatureTag.
type ListFeatureTagsResponse_SdkV2 struct {
	FeatureTags types.List `tfsdk:"feature_tags"`
	// Pagination token to request the next page of results for this query.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListFeatureTagsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListFeatureTagsResponse_SdkV2) {
	if !from.FeatureTags.IsNull() && !from.FeatureTags.IsUnknown() && to.FeatureTags.IsNull() && len(from.FeatureTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FeatureTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FeatureTags = from.FeatureTags
	}
}

func (to *ListFeatureTagsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListFeatureTagsResponse_SdkV2) {
	if !from.FeatureTags.IsNull() && !from.FeatureTags.IsUnknown() && to.FeatureTags.IsNull() && len(from.FeatureTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FeatureTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FeatureTags = from.FeatureTags
	}
}

func (m ListFeatureTagsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["feature_tags"] = attrs["feature_tags"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListFeatureTagsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListFeatureTagsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"feature_tags": reflect.TypeOf(FeatureTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFeatureTagsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListFeatureTagsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_tags":    m.FeatureTags,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListFeatureTagsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"feature_tags": basetypes.ListType{
				ElemType: FeatureTag_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetFeatureTags returns the value of the FeatureTags field in ListFeatureTagsResponse_SdkV2 as
// a slice of FeatureTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListFeatureTagsResponse_SdkV2) GetFeatureTags(ctx context.Context) ([]FeatureTag_SdkV2, bool) {
	if m.FeatureTags.IsNull() || m.FeatureTags.IsUnknown() {
		return nil, false
	}
	var v []FeatureTag_SdkV2
	d := m.FeatureTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFeatureTags sets the value of the FeatureTags field in ListFeatureTagsResponse_SdkV2.
func (m *ListFeatureTagsResponse_SdkV2) SetFeatureTags(ctx context.Context, v []FeatureTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["feature_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.FeatureTags = types.ListValueMust(t, vs)
}

type ListFeaturesRequest_SdkV2 struct {
	// The maximum number of results to return.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page based on a previous query.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListFeaturesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListFeaturesRequest_SdkV2) {
}

func (to *ListFeaturesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListFeaturesRequest_SdkV2) {
}

func (m ListFeaturesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListFeaturesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListFeaturesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFeaturesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListFeaturesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListFeaturesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListFeaturesResponse_SdkV2 struct {
	// List of features.
	Features types.List `tfsdk:"features"`
	// Pagination token to request the next page of results for this query.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListFeaturesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListFeaturesResponse_SdkV2) {
	if !from.Features.IsNull() && !from.Features.IsUnknown() && to.Features.IsNull() && len(from.Features.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Features, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Features = from.Features
	}
}

func (to *ListFeaturesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListFeaturesResponse_SdkV2) {
	if !from.Features.IsNull() && !from.Features.IsUnknown() && to.Features.IsNull() && len(from.Features.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Features, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Features = from.Features
	}
}

func (m ListFeaturesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["features"] = attrs["features"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListFeaturesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListFeaturesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"features": reflect.TypeOf(Feature_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFeaturesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListFeaturesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"features":        m.Features,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListFeaturesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"features": basetypes.ListType{
				ElemType: Feature_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetFeatures returns the value of the Features field in ListFeaturesResponse_SdkV2 as
// a slice of Feature_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListFeaturesResponse_SdkV2) GetFeatures(ctx context.Context) ([]Feature_SdkV2, bool) {
	if m.Features.IsNull() || m.Features.IsUnknown() {
		return nil, false
	}
	var v []Feature_SdkV2
	d := m.Features.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFeatures sets the value of the Features field in ListFeaturesResponse_SdkV2.
func (m *ListFeaturesResponse_SdkV2) SetFeatures(ctx context.Context, v []Feature_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["features"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Features = types.ListValueMust(t, vs)
}

type ListModelsRequest_SdkV2 struct {
	// Maximum number of registered models desired. Max threshold is 1000.
	MaxResults types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page based on a previous query.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListModelsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListModelsRequest_SdkV2) {
}

func (to *ListModelsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListModelsRequest_SdkV2) {
}

func (m ListModelsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["max_results"] = attrs["max_results"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListModelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListModelsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListModelsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListModelsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results": m.MaxResults,
			"page_token":  m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListModelsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_results": types.Int64Type,
			"page_token":  types.StringType,
		},
	}
}

type ListModelsResponse_SdkV2 struct {
	// Pagination token to request next page of models for the same query.
	NextPageToken types.String `tfsdk:"next_page_token"`

	RegisteredModels types.List `tfsdk:"registered_models"`
}

func (to *ListModelsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListModelsResponse_SdkV2) {
	if !from.RegisteredModels.IsNull() && !from.RegisteredModels.IsUnknown() && to.RegisteredModels.IsNull() && len(from.RegisteredModels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RegisteredModels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RegisteredModels = from.RegisteredModels
	}
}

func (to *ListModelsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListModelsResponse_SdkV2) {
	if !from.RegisteredModels.IsNull() && !from.RegisteredModels.IsUnknown() && to.RegisteredModels.IsNull() && len(from.RegisteredModels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RegisteredModels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RegisteredModels = from.RegisteredModels
	}
}

func (m ListModelsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["registered_models"] = attrs["registered_models"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListModelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListModelsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"registered_models": reflect.TypeOf(Model_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListModelsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListModelsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":   m.NextPageToken,
			"registered_models": m.RegisteredModels,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListModelsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"registered_models": basetypes.ListType{
				ElemType: Model_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetRegisteredModels returns the value of the RegisteredModels field in ListModelsResponse_SdkV2 as
// a slice of Model_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListModelsResponse_SdkV2) GetRegisteredModels(ctx context.Context) ([]Model_SdkV2, bool) {
	if m.RegisteredModels.IsNull() || m.RegisteredModels.IsUnknown() {
		return nil, false
	}
	var v []Model_SdkV2
	d := m.RegisteredModels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRegisteredModels sets the value of the RegisteredModels field in ListModelsResponse_SdkV2.
func (m *ListModelsResponse_SdkV2) SetRegisteredModels(ctx context.Context, v []Model_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["registered_models"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.RegisteredModels = types.ListValueMust(t, vs)
}

type ListOnlineStoresRequest_SdkV2 struct {
	// The maximum number of results to return. Defaults to 100 if not
	// specified.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page based on a previous query.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListOnlineStoresRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListOnlineStoresRequest_SdkV2) {
}

func (to *ListOnlineStoresRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListOnlineStoresRequest_SdkV2) {
}

func (m ListOnlineStoresRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListOnlineStoresRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListOnlineStoresRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListOnlineStoresRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListOnlineStoresRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListOnlineStoresRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListOnlineStoresResponse_SdkV2 struct {
	// Pagination token to request the next page of results for this query.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// List of online stores.
	OnlineStores types.List `tfsdk:"online_stores"`
}

func (to *ListOnlineStoresResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListOnlineStoresResponse_SdkV2) {
	if !from.OnlineStores.IsNull() && !from.OnlineStores.IsUnknown() && to.OnlineStores.IsNull() && len(from.OnlineStores.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OnlineStores, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OnlineStores = from.OnlineStores
	}
}

func (to *ListOnlineStoresResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListOnlineStoresResponse_SdkV2) {
	if !from.OnlineStores.IsNull() && !from.OnlineStores.IsUnknown() && to.OnlineStores.IsNull() && len(from.OnlineStores.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OnlineStores, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OnlineStores = from.OnlineStores
	}
}

func (m ListOnlineStoresResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["online_stores"] = attrs["online_stores"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListOnlineStoresResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListOnlineStoresResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"online_stores": reflect.TypeOf(OnlineStore_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListOnlineStoresResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListOnlineStoresResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"online_stores":   m.OnlineStores,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListOnlineStoresResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"online_stores": basetypes.ListType{
				ElemType: OnlineStore_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetOnlineStores returns the value of the OnlineStores field in ListOnlineStoresResponse_SdkV2 as
// a slice of OnlineStore_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListOnlineStoresResponse_SdkV2) GetOnlineStores(ctx context.Context) ([]OnlineStore_SdkV2, bool) {
	if m.OnlineStores.IsNull() || m.OnlineStores.IsUnknown() {
		return nil, false
	}
	var v []OnlineStore_SdkV2
	d := m.OnlineStores.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOnlineStores sets the value of the OnlineStores field in ListOnlineStoresResponse_SdkV2.
func (m *ListOnlineStoresResponse_SdkV2) SetOnlineStores(ctx context.Context, v []OnlineStore_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["online_stores"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.OnlineStores = types.ListValueMust(t, vs)
}

type ListRegistryWebhooks_SdkV2 struct {
	// Token that can be used to retrieve the next page of artifact results
	NextPageToken types.String `tfsdk:"next_page_token"`
	// Array of registry webhooks.
	Webhooks types.List `tfsdk:"webhooks"`
}

func (to *ListRegistryWebhooks_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListRegistryWebhooks_SdkV2) {
	if !from.Webhooks.IsNull() && !from.Webhooks.IsUnknown() && to.Webhooks.IsNull() && len(from.Webhooks.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Webhooks, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Webhooks = from.Webhooks
	}
}

func (to *ListRegistryWebhooks_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListRegistryWebhooks_SdkV2) {
	if !from.Webhooks.IsNull() && !from.Webhooks.IsUnknown() && to.Webhooks.IsNull() && len(from.Webhooks.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Webhooks, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Webhooks = from.Webhooks
	}
}

func (m ListRegistryWebhooks_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["webhooks"] = attrs["webhooks"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListRegistryWebhooks.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListRegistryWebhooks_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"webhooks": reflect.TypeOf(RegistryWebhook_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListRegistryWebhooks_SdkV2
// only implements ToObjectValue() and Type().
func (m ListRegistryWebhooks_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"webhooks":        m.Webhooks,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListRegistryWebhooks_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"webhooks": basetypes.ListType{
				ElemType: RegistryWebhook_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetWebhooks returns the value of the Webhooks field in ListRegistryWebhooks_SdkV2 as
// a slice of RegistryWebhook_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListRegistryWebhooks_SdkV2) GetWebhooks(ctx context.Context) ([]RegistryWebhook_SdkV2, bool) {
	if m.Webhooks.IsNull() || m.Webhooks.IsUnknown() {
		return nil, false
	}
	var v []RegistryWebhook_SdkV2
	d := m.Webhooks.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWebhooks sets the value of the Webhooks field in ListRegistryWebhooks_SdkV2.
func (m *ListRegistryWebhooks_SdkV2) SetWebhooks(ctx context.Context, v []RegistryWebhook_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["webhooks"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Webhooks = types.ListValueMust(t, vs)
}

type ListTransitionRequestsRequest_SdkV2 struct {
	// Name of the registered model.
	Name types.String `tfsdk:"-"`
	// Version of the model.
	Version types.String `tfsdk:"-"`
}

func (to *ListTransitionRequestsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListTransitionRequestsRequest_SdkV2) {
}

func (to *ListTransitionRequestsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListTransitionRequestsRequest_SdkV2) {
}

func (m ListTransitionRequestsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["version"] = attrs["version"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListTransitionRequestsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListTransitionRequestsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTransitionRequestsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListTransitionRequestsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":    m.Name,
			"version": m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListTransitionRequestsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":    types.StringType,
			"version": types.StringType,
		},
	}
}

type ListTransitionRequestsResponse_SdkV2 struct {
	// Array of open transition requests.
	Requests types.List `tfsdk:"requests"`
}

func (to *ListTransitionRequestsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListTransitionRequestsResponse_SdkV2) {
	if !from.Requests.IsNull() && !from.Requests.IsUnknown() && to.Requests.IsNull() && len(from.Requests.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Requests, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Requests = from.Requests
	}
}

func (to *ListTransitionRequestsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListTransitionRequestsResponse_SdkV2) {
	if !from.Requests.IsNull() && !from.Requests.IsUnknown() && to.Requests.IsNull() && len(from.Requests.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Requests, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Requests = from.Requests
	}
}

func (m ListTransitionRequestsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["requests"] = attrs["requests"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListTransitionRequestsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListTransitionRequestsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"requests": reflect.TypeOf(Activity_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTransitionRequestsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListTransitionRequestsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"requests": m.Requests,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListTransitionRequestsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"requests": basetypes.ListType{
				ElemType: Activity_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetRequests returns the value of the Requests field in ListTransitionRequestsResponse_SdkV2 as
// a slice of Activity_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListTransitionRequestsResponse_SdkV2) GetRequests(ctx context.Context) ([]Activity_SdkV2, bool) {
	if m.Requests.IsNull() || m.Requests.IsUnknown() {
		return nil, false
	}
	var v []Activity_SdkV2
	d := m.Requests.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRequests sets the value of the Requests field in ListTransitionRequestsResponse_SdkV2.
func (m *ListTransitionRequestsResponse_SdkV2) SetRequests(ctx context.Context, v []Activity_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["requests"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Requests = types.ListValueMust(t, vs)
}

type ListWebhooksRequest_SdkV2 struct {
	// Events that trigger the webhook. * `MODEL_VERSION_CREATED`: A new model
	// version was created for the associated model.
	//
	// * `MODEL_VERSION_TRANSITIONED_STAGE`: A model version’s stage was
	// changed.
	//
	// * `TRANSITION_REQUEST_CREATED`: A user requested a model version’s
	// stage be transitioned.
	//
	// * `COMMENT_CREATED`: A user wrote a comment on a registered model.
	//
	// * `REGISTERED_MODEL_CREATED`: A new registered model was created. This
	// event type can only be specified for a registry-wide webhook, which can
	// be created by not specifying a model name in the create request.
	//
	// * `MODEL_VERSION_TAG_SET`: A user set a tag on the model version.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_STAGING`: A model version was
	// transitioned to staging.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_PRODUCTION`: A model version was
	// transitioned to production.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_ARCHIVED`: A model version was archived.
	//
	// * `TRANSITION_REQUEST_TO_STAGING_CREATED`: A user requested a model
	// version be transitioned to staging.
	//
	// * `TRANSITION_REQUEST_TO_PRODUCTION_CREATED`: A user requested a model
	// version be transitioned to production.
	//
	// * `TRANSITION_REQUEST_TO_ARCHIVED_CREATED`: A user requested a model
	// version be archived.
	//
	// If `events` is specified, any webhook with one or more of the specified
	// trigger events is included in the output. If `events` is not specified,
	// webhooks of all event types are included in the output.
	Events types.List `tfsdk:"-"`

	MaxResults types.Int64 `tfsdk:"-"`
	// Registered model name If not specified, all webhooks associated with the
	// specified events are listed, regardless of their associated model.
	ModelName types.String `tfsdk:"-"`
	// Token indicating the page of artifact results to fetch
	PageToken types.String `tfsdk:"-"`
}

func (to *ListWebhooksRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListWebhooksRequest_SdkV2) {
	if !from.Events.IsNull() && !from.Events.IsUnknown() && to.Events.IsNull() && len(from.Events.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Events, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Events = from.Events
	}
}

func (to *ListWebhooksRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListWebhooksRequest_SdkV2) {
	if !from.Events.IsNull() && !from.Events.IsUnknown() && to.Events.IsNull() && len(from.Events.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Events, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Events = from.Events
	}
}

func (m ListWebhooksRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["model_name"] = attrs["model_name"].SetOptional()
	attrs["events"] = attrs["events"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["max_results"] = attrs["max_results"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListWebhooksRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListWebhooksRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"events": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWebhooksRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListWebhooksRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"events":      m.Events,
			"max_results": m.MaxResults,
			"model_name":  m.ModelName,
			"page_token":  m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListWebhooksRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"events": basetypes.ListType{
				ElemType: types.StringType,
			},
			"max_results": types.Int64Type,
			"model_name":  types.StringType,
			"page_token":  types.StringType,
		},
	}
}

// GetEvents returns the value of the Events field in ListWebhooksRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListWebhooksRequest_SdkV2) GetEvents(ctx context.Context) ([]types.String, bool) {
	if m.Events.IsNull() || m.Events.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Events.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEvents sets the value of the Events field in ListWebhooksRequest_SdkV2.
func (m *ListWebhooksRequest_SdkV2) SetEvents(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["events"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Events = types.ListValueMust(t, vs)
}

type LogBatch_SdkV2 struct {
	// Metrics to log. A single request can contain up to 1000 metrics, and up
	// to 1000 metrics, params, and tags in total.
	Metrics types.List `tfsdk:"metrics"`
	// Params to log. A single request can contain up to 100 params, and up to
	// 1000 metrics, params, and tags in total.
	Params types.List `tfsdk:"params"`
	// ID of the run to log under
	RunId types.String `tfsdk:"run_id"`
	// Tags to log. A single request can contain up to 100 tags, and up to 1000
	// metrics, params, and tags in total.
	Tags types.List `tfsdk:"tags"`
}

func (to *LogBatch_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogBatch_SdkV2) {
	if !from.Metrics.IsNull() && !from.Metrics.IsUnknown() && to.Metrics.IsNull() && len(from.Metrics.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Metrics, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Metrics = from.Metrics
	}
	if !from.Params.IsNull() && !from.Params.IsUnknown() && to.Params.IsNull() && len(from.Params.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Params, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Params = from.Params
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *LogBatch_SdkV2) SyncFieldsDuringRead(ctx context.Context, from LogBatch_SdkV2) {
	if !from.Metrics.IsNull() && !from.Metrics.IsUnknown() && to.Metrics.IsNull() && len(from.Metrics.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Metrics, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Metrics = from.Metrics
	}
	if !from.Params.IsNull() && !from.Params.IsUnknown() && to.Params.IsNull() && len(from.Params.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Params, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Params = from.Params
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m LogBatch_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["metrics"] = attrs["metrics"].SetOptional()
	attrs["params"] = attrs["params"].SetOptional()
	attrs["run_id"] = attrs["run_id"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogBatch.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LogBatch_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metrics": reflect.TypeOf(Metric_SdkV2{}),
		"params":  reflect.TypeOf(Param_SdkV2{}),
		"tags":    reflect.TypeOf(RunTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogBatch_SdkV2
// only implements ToObjectValue() and Type().
func (m LogBatch_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metrics": m.Metrics,
			"params":  m.Params,
			"run_id":  m.RunId,
			"tags":    m.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LogBatch_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metrics": basetypes.ListType{
				ElemType: Metric_SdkV2{}.Type(ctx),
			},
			"params": basetypes.ListType{
				ElemType: Param_SdkV2{}.Type(ctx),
			},
			"run_id": types.StringType,
			"tags": basetypes.ListType{
				ElemType: RunTag_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetMetrics returns the value of the Metrics field in LogBatch_SdkV2 as
// a slice of Metric_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *LogBatch_SdkV2) GetMetrics(ctx context.Context) ([]Metric_SdkV2, bool) {
	if m.Metrics.IsNull() || m.Metrics.IsUnknown() {
		return nil, false
	}
	var v []Metric_SdkV2
	d := m.Metrics.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMetrics sets the value of the Metrics field in LogBatch_SdkV2.
func (m *LogBatch_SdkV2) SetMetrics(ctx context.Context, v []Metric_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["metrics"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Metrics = types.ListValueMust(t, vs)
}

// GetParams returns the value of the Params field in LogBatch_SdkV2 as
// a slice of Param_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *LogBatch_SdkV2) GetParams(ctx context.Context) ([]Param_SdkV2, bool) {
	if m.Params.IsNull() || m.Params.IsUnknown() {
		return nil, false
	}
	var v []Param_SdkV2
	d := m.Params.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParams sets the value of the Params field in LogBatch_SdkV2.
func (m *LogBatch_SdkV2) SetParams(ctx context.Context, v []Param_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["params"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Params = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in LogBatch_SdkV2 as
// a slice of RunTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *LogBatch_SdkV2) GetTags(ctx context.Context) ([]RunTag_SdkV2, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []RunTag_SdkV2
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in LogBatch_SdkV2.
func (m *LogBatch_SdkV2) SetTags(ctx context.Context, v []RunTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type LogBatchResponse_SdkV2 struct {
}

func (to *LogBatchResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogBatchResponse_SdkV2) {
}

func (to *LogBatchResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from LogBatchResponse_SdkV2) {
}

func (m LogBatchResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogBatchResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LogBatchResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogBatchResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m LogBatchResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m LogBatchResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type LogInputs_SdkV2 struct {
	// Dataset inputs
	Datasets types.List `tfsdk:"datasets"`
	// Model inputs
	Models types.List `tfsdk:"models"`
	// ID of the run to log under
	RunId types.String `tfsdk:"run_id"`
}

func (to *LogInputs_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogInputs_SdkV2) {
	if !from.Datasets.IsNull() && !from.Datasets.IsUnknown() && to.Datasets.IsNull() && len(from.Datasets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Datasets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Datasets = from.Datasets
	}
	if !from.Models.IsNull() && !from.Models.IsUnknown() && to.Models.IsNull() && len(from.Models.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Models, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Models = from.Models
	}
}

func (to *LogInputs_SdkV2) SyncFieldsDuringRead(ctx context.Context, from LogInputs_SdkV2) {
	if !from.Datasets.IsNull() && !from.Datasets.IsUnknown() && to.Datasets.IsNull() && len(from.Datasets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Datasets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Datasets = from.Datasets
	}
	if !from.Models.IsNull() && !from.Models.IsUnknown() && to.Models.IsNull() && len(from.Models.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Models, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Models = from.Models
	}
}

func (m LogInputs_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["datasets"] = attrs["datasets"].SetOptional()
	attrs["models"] = attrs["models"].SetOptional()
	attrs["run_id"] = attrs["run_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogInputs.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LogInputs_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"datasets": reflect.TypeOf(DatasetInput_SdkV2{}),
		"models":   reflect.TypeOf(ModelInput_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogInputs_SdkV2
// only implements ToObjectValue() and Type().
func (m LogInputs_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"datasets": m.Datasets,
			"models":   m.Models,
			"run_id":   m.RunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LogInputs_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"datasets": basetypes.ListType{
				ElemType: DatasetInput_SdkV2{}.Type(ctx),
			},
			"models": basetypes.ListType{
				ElemType: ModelInput_SdkV2{}.Type(ctx),
			},
			"run_id": types.StringType,
		},
	}
}

// GetDatasets returns the value of the Datasets field in LogInputs_SdkV2 as
// a slice of DatasetInput_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *LogInputs_SdkV2) GetDatasets(ctx context.Context) ([]DatasetInput_SdkV2, bool) {
	if m.Datasets.IsNull() || m.Datasets.IsUnknown() {
		return nil, false
	}
	var v []DatasetInput_SdkV2
	d := m.Datasets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatasets sets the value of the Datasets field in LogInputs_SdkV2.
func (m *LogInputs_SdkV2) SetDatasets(ctx context.Context, v []DatasetInput_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["datasets"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Datasets = types.ListValueMust(t, vs)
}

// GetModels returns the value of the Models field in LogInputs_SdkV2 as
// a slice of ModelInput_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *LogInputs_SdkV2) GetModels(ctx context.Context) ([]ModelInput_SdkV2, bool) {
	if m.Models.IsNull() || m.Models.IsUnknown() {
		return nil, false
	}
	var v []ModelInput_SdkV2
	d := m.Models.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModels sets the value of the Models field in LogInputs_SdkV2.
func (m *LogInputs_SdkV2) SetModels(ctx context.Context, v []ModelInput_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["models"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Models = types.ListValueMust(t, vs)
}

type LogInputsResponse_SdkV2 struct {
}

func (to *LogInputsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogInputsResponse_SdkV2) {
}

func (to *LogInputsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from LogInputsResponse_SdkV2) {
}

func (m LogInputsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogInputsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LogInputsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogInputsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m LogInputsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m LogInputsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type LogLoggedModelParamsRequest_SdkV2 struct {
	// The ID of the logged model to log params for.
	ModelId types.String `tfsdk:"-"`
	// Parameters to attach to the model.
	Params types.List `tfsdk:"params"`
}

func (to *LogLoggedModelParamsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogLoggedModelParamsRequest_SdkV2) {
	if !from.Params.IsNull() && !from.Params.IsUnknown() && to.Params.IsNull() && len(from.Params.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Params, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Params = from.Params
	}
}

func (to *LogLoggedModelParamsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from LogLoggedModelParamsRequest_SdkV2) {
	if !from.Params.IsNull() && !from.Params.IsUnknown() && to.Params.IsNull() && len(from.Params.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Params, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Params = from.Params
	}
}

func (m LogLoggedModelParamsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["params"] = attrs["params"].SetOptional()
	attrs["model_id"] = attrs["model_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogLoggedModelParamsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LogLoggedModelParamsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"params": reflect.TypeOf(LoggedModelParameter_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogLoggedModelParamsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m LogLoggedModelParamsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_id": m.ModelId,
			"params":   m.Params,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LogLoggedModelParamsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_id": types.StringType,
			"params": basetypes.ListType{
				ElemType: LoggedModelParameter_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetParams returns the value of the Params field in LogLoggedModelParamsRequest_SdkV2 as
// a slice of LoggedModelParameter_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *LogLoggedModelParamsRequest_SdkV2) GetParams(ctx context.Context) ([]LoggedModelParameter_SdkV2, bool) {
	if m.Params.IsNull() || m.Params.IsUnknown() {
		return nil, false
	}
	var v []LoggedModelParameter_SdkV2
	d := m.Params.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParams sets the value of the Params field in LogLoggedModelParamsRequest_SdkV2.
func (m *LogLoggedModelParamsRequest_SdkV2) SetParams(ctx context.Context, v []LoggedModelParameter_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["params"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Params = types.ListValueMust(t, vs)
}

type LogLoggedModelParamsRequestResponse_SdkV2 struct {
}

func (to *LogLoggedModelParamsRequestResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogLoggedModelParamsRequestResponse_SdkV2) {
}

func (to *LogLoggedModelParamsRequestResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from LogLoggedModelParamsRequestResponse_SdkV2) {
}

func (m LogLoggedModelParamsRequestResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogLoggedModelParamsRequestResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LogLoggedModelParamsRequestResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogLoggedModelParamsRequestResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m LogLoggedModelParamsRequestResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m LogLoggedModelParamsRequestResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type LogMetric_SdkV2 struct {
	// Dataset digest of the dataset associated with the metric, e.g. an md5
	// hash of the dataset that uniquely identifies it within datasets of the
	// same name.
	DatasetDigest types.String `tfsdk:"dataset_digest"`
	// The name of the dataset associated with the metric. E.g.
	// “my.uc.table@2” “nyc-taxi-dataset”, “fantastic-elk-3”
	DatasetName types.String `tfsdk:"dataset_name"`
	// Name of the metric.
	Key types.String `tfsdk:"key"`
	// ID of the logged model associated with the metric, if applicable
	ModelId types.String `tfsdk:"model_id"`
	// ID of the run under which to log the metric. Must be provided.
	RunId types.String `tfsdk:"run_id"`
	// [Deprecated, use `run_id` instead] ID of the run under which to log the
	// metric. This field will be removed in a future MLflow version.
	RunUuid types.String `tfsdk:"run_uuid"`
	// Step at which to log the metric
	Step types.Int64 `tfsdk:"step"`
	// Unix timestamp in milliseconds at the time metric was logged.
	Timestamp types.Int64 `tfsdk:"timestamp"`
	// Double value of the metric being logged.
	Value types.Float64 `tfsdk:"value"`
}

func (to *LogMetric_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogMetric_SdkV2) {
}

func (to *LogMetric_SdkV2) SyncFieldsDuringRead(ctx context.Context, from LogMetric_SdkV2) {
}

func (m LogMetric_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dataset_digest"] = attrs["dataset_digest"].SetOptional()
	attrs["dataset_name"] = attrs["dataset_name"].SetOptional()
	attrs["key"] = attrs["key"].SetRequired()
	attrs["model_id"] = attrs["model_id"].SetOptional()
	attrs["run_id"] = attrs["run_id"].SetOptional()
	attrs["run_uuid"] = attrs["run_uuid"].SetOptional()
	attrs["step"] = attrs["step"].SetOptional()
	attrs["timestamp"] = attrs["timestamp"].SetRequired()
	attrs["value"] = attrs["value"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogMetric.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LogMetric_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogMetric_SdkV2
// only implements ToObjectValue() and Type().
func (m LogMetric_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dataset_digest": m.DatasetDigest,
			"dataset_name":   m.DatasetName,
			"key":            m.Key,
			"model_id":       m.ModelId,
			"run_id":         m.RunId,
			"run_uuid":       m.RunUuid,
			"step":           m.Step,
			"timestamp":      m.Timestamp,
			"value":          m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LogMetric_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dataset_digest": types.StringType,
			"dataset_name":   types.StringType,
			"key":            types.StringType,
			"model_id":       types.StringType,
			"run_id":         types.StringType,
			"run_uuid":       types.StringType,
			"step":           types.Int64Type,
			"timestamp":      types.Int64Type,
			"value":          types.Float64Type,
		},
	}
}

type LogMetricResponse_SdkV2 struct {
}

func (to *LogMetricResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogMetricResponse_SdkV2) {
}

func (to *LogMetricResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from LogMetricResponse_SdkV2) {
}

func (m LogMetricResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogMetricResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LogMetricResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogMetricResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m LogMetricResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m LogMetricResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type LogModel_SdkV2 struct {
	// MLmodel file in json format.
	ModelJson types.String `tfsdk:"model_json"`
	// ID of the run to log under
	RunId types.String `tfsdk:"run_id"`
}

func (to *LogModel_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogModel_SdkV2) {
}

func (to *LogModel_SdkV2) SyncFieldsDuringRead(ctx context.Context, from LogModel_SdkV2) {
}

func (m LogModel_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["model_json"] = attrs["model_json"].SetOptional()
	attrs["run_id"] = attrs["run_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogModel.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LogModel_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogModel_SdkV2
// only implements ToObjectValue() and Type().
func (m LogModel_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_json": m.ModelJson,
			"run_id":     m.RunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LogModel_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_json": types.StringType,
			"run_id":     types.StringType,
		},
	}
}

type LogModelResponse_SdkV2 struct {
}

func (to *LogModelResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogModelResponse_SdkV2) {
}

func (to *LogModelResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from LogModelResponse_SdkV2) {
}

func (m LogModelResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogModelResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LogModelResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogModelResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m LogModelResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m LogModelResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type LogOutputsRequest_SdkV2 struct {
	// The model outputs from the Run.
	Models types.List `tfsdk:"models"`
	// The ID of the Run from which to log outputs.
	RunId types.String `tfsdk:"run_id"`
}

func (to *LogOutputsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogOutputsRequest_SdkV2) {
	if !from.Models.IsNull() && !from.Models.IsUnknown() && to.Models.IsNull() && len(from.Models.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Models, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Models = from.Models
	}
}

func (to *LogOutputsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from LogOutputsRequest_SdkV2) {
	if !from.Models.IsNull() && !from.Models.IsUnknown() && to.Models.IsNull() && len(from.Models.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Models, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Models = from.Models
	}
}

func (m LogOutputsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["models"] = attrs["models"].SetOptional()
	attrs["run_id"] = attrs["run_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogOutputsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LogOutputsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"models": reflect.TypeOf(ModelOutput_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogOutputsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m LogOutputsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"models": m.Models,
			"run_id": m.RunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LogOutputsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"models": basetypes.ListType{
				ElemType: ModelOutput_SdkV2{}.Type(ctx),
			},
			"run_id": types.StringType,
		},
	}
}

// GetModels returns the value of the Models field in LogOutputsRequest_SdkV2 as
// a slice of ModelOutput_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *LogOutputsRequest_SdkV2) GetModels(ctx context.Context) ([]ModelOutput_SdkV2, bool) {
	if m.Models.IsNull() || m.Models.IsUnknown() {
		return nil, false
	}
	var v []ModelOutput_SdkV2
	d := m.Models.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModels sets the value of the Models field in LogOutputsRequest_SdkV2.
func (m *LogOutputsRequest_SdkV2) SetModels(ctx context.Context, v []ModelOutput_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["models"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Models = types.ListValueMust(t, vs)
}

type LogOutputsResponse_SdkV2 struct {
}

func (to *LogOutputsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogOutputsResponse_SdkV2) {
}

func (to *LogOutputsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from LogOutputsResponse_SdkV2) {
}

func (m LogOutputsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogOutputsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LogOutputsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogOutputsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m LogOutputsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m LogOutputsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type LogParam_SdkV2 struct {
	// Name of the param. Maximum size is 255 bytes.
	Key types.String `tfsdk:"key"`
	// ID of the run under which to log the param. Must be provided.
	RunId types.String `tfsdk:"run_id"`
	// [Deprecated, use `run_id` instead] ID of the run under which to log the
	// param. This field will be removed in a future MLflow version.
	RunUuid types.String `tfsdk:"run_uuid"`
	// String value of the param being logged. Maximum size is 500 bytes.
	Value types.String `tfsdk:"value"`
}

func (to *LogParam_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogParam_SdkV2) {
}

func (to *LogParam_SdkV2) SyncFieldsDuringRead(ctx context.Context, from LogParam_SdkV2) {
}

func (m LogParam_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetRequired()
	attrs["run_id"] = attrs["run_id"].SetOptional()
	attrs["run_uuid"] = attrs["run_uuid"].SetOptional()
	attrs["value"] = attrs["value"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogParam.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LogParam_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogParam_SdkV2
// only implements ToObjectValue() and Type().
func (m LogParam_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":      m.Key,
			"run_id":   m.RunId,
			"run_uuid": m.RunUuid,
			"value":    m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LogParam_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":      types.StringType,
			"run_id":   types.StringType,
			"run_uuid": types.StringType,
			"value":    types.StringType,
		},
	}
}

type LogParamResponse_SdkV2 struct {
}

func (to *LogParamResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogParamResponse_SdkV2) {
}

func (to *LogParamResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from LogParamResponse_SdkV2) {
}

func (m LogParamResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogParamResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LogParamResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogParamResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m LogParamResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m LogParamResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// A logged model message includes logged model attributes, tags, registration
// info, params, and linked run metrics.
type LoggedModel_SdkV2 struct {
	// The params and metrics attached to the logged model.
	Data types.List `tfsdk:"data"`
	// The logged model attributes such as model ID, status, tags, etc.
	Info types.List `tfsdk:"info"`
}

func (to *LoggedModel_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LoggedModel_SdkV2) {
	if !from.Data.IsNull() && !from.Data.IsUnknown() {
		if toData, ok := to.GetData(ctx); ok {
			if fromData, ok := from.GetData(ctx); ok {
				// Recursively sync the fields of Data
				toData.SyncFieldsDuringCreateOrUpdate(ctx, fromData)
				to.SetData(ctx, toData)
			}
		}
	}
	if !from.Info.IsNull() && !from.Info.IsUnknown() {
		if toInfo, ok := to.GetInfo(ctx); ok {
			if fromInfo, ok := from.GetInfo(ctx); ok {
				// Recursively sync the fields of Info
				toInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromInfo)
				to.SetInfo(ctx, toInfo)
			}
		}
	}
}

func (to *LoggedModel_SdkV2) SyncFieldsDuringRead(ctx context.Context, from LoggedModel_SdkV2) {
	if !from.Data.IsNull() && !from.Data.IsUnknown() {
		if toData, ok := to.GetData(ctx); ok {
			if fromData, ok := from.GetData(ctx); ok {
				toData.SyncFieldsDuringRead(ctx, fromData)
				to.SetData(ctx, toData)
			}
		}
	}
	if !from.Info.IsNull() && !from.Info.IsUnknown() {
		if toInfo, ok := to.GetInfo(ctx); ok {
			if fromInfo, ok := from.GetInfo(ctx); ok {
				toInfo.SyncFieldsDuringRead(ctx, fromInfo)
				to.SetInfo(ctx, toInfo)
			}
		}
	}
}

func (m LoggedModel_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["data"] = attrs["data"].SetOptional()
	attrs["data"] = attrs["data"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["info"] = attrs["info"].SetOptional()
	attrs["info"] = attrs["info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LoggedModel.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LoggedModel_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data": reflect.TypeOf(LoggedModelData_SdkV2{}),
		"info": reflect.TypeOf(LoggedModelInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LoggedModel_SdkV2
// only implements ToObjectValue() and Type().
func (m LoggedModel_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data": m.Data,
			"info": m.Info,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LoggedModel_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"data": basetypes.ListType{
				ElemType: LoggedModelData_SdkV2{}.Type(ctx),
			},
			"info": basetypes.ListType{
				ElemType: LoggedModelInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetData returns the value of the Data field in LoggedModel_SdkV2 as
// a LoggedModelData_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *LoggedModel_SdkV2) GetData(ctx context.Context) (LoggedModelData_SdkV2, bool) {
	var e LoggedModelData_SdkV2
	if m.Data.IsNull() || m.Data.IsUnknown() {
		return e, false
	}
	var v []LoggedModelData_SdkV2
	d := m.Data.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetData sets the value of the Data field in LoggedModel_SdkV2.
func (m *LoggedModel_SdkV2) SetData(ctx context.Context, v LoggedModelData_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["data"]
	m.Data = types.ListValueMust(t, vs)
}

// GetInfo returns the value of the Info field in LoggedModel_SdkV2 as
// a LoggedModelInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *LoggedModel_SdkV2) GetInfo(ctx context.Context) (LoggedModelInfo_SdkV2, bool) {
	var e LoggedModelInfo_SdkV2
	if m.Info.IsNull() || m.Info.IsUnknown() {
		return e, false
	}
	var v []LoggedModelInfo_SdkV2
	d := m.Info.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInfo sets the value of the Info field in LoggedModel_SdkV2.
func (m *LoggedModel_SdkV2) SetInfo(ctx context.Context, v LoggedModelInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["info"]
	m.Info = types.ListValueMust(t, vs)
}

// A LoggedModelData message includes logged model params and linked metrics.
type LoggedModelData_SdkV2 struct {
	// Performance metrics linked to the model.
	Metrics types.List `tfsdk:"metrics"`
	// Immutable string key-value pairs of the model.
	Params types.List `tfsdk:"params"`
}

func (to *LoggedModelData_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LoggedModelData_SdkV2) {
	if !from.Metrics.IsNull() && !from.Metrics.IsUnknown() && to.Metrics.IsNull() && len(from.Metrics.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Metrics, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Metrics = from.Metrics
	}
	if !from.Params.IsNull() && !from.Params.IsUnknown() && to.Params.IsNull() && len(from.Params.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Params, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Params = from.Params
	}
}

func (to *LoggedModelData_SdkV2) SyncFieldsDuringRead(ctx context.Context, from LoggedModelData_SdkV2) {
	if !from.Metrics.IsNull() && !from.Metrics.IsUnknown() && to.Metrics.IsNull() && len(from.Metrics.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Metrics, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Metrics = from.Metrics
	}
	if !from.Params.IsNull() && !from.Params.IsUnknown() && to.Params.IsNull() && len(from.Params.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Params, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Params = from.Params
	}
}

func (m LoggedModelData_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["metrics"] = attrs["metrics"].SetOptional()
	attrs["params"] = attrs["params"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LoggedModelData.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LoggedModelData_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metrics": reflect.TypeOf(Metric_SdkV2{}),
		"params":  reflect.TypeOf(LoggedModelParameter_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LoggedModelData_SdkV2
// only implements ToObjectValue() and Type().
func (m LoggedModelData_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metrics": m.Metrics,
			"params":  m.Params,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LoggedModelData_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metrics": basetypes.ListType{
				ElemType: Metric_SdkV2{}.Type(ctx),
			},
			"params": basetypes.ListType{
				ElemType: LoggedModelParameter_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetMetrics returns the value of the Metrics field in LoggedModelData_SdkV2 as
// a slice of Metric_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *LoggedModelData_SdkV2) GetMetrics(ctx context.Context) ([]Metric_SdkV2, bool) {
	if m.Metrics.IsNull() || m.Metrics.IsUnknown() {
		return nil, false
	}
	var v []Metric_SdkV2
	d := m.Metrics.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMetrics sets the value of the Metrics field in LoggedModelData_SdkV2.
func (m *LoggedModelData_SdkV2) SetMetrics(ctx context.Context, v []Metric_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["metrics"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Metrics = types.ListValueMust(t, vs)
}

// GetParams returns the value of the Params field in LoggedModelData_SdkV2 as
// a slice of LoggedModelParameter_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *LoggedModelData_SdkV2) GetParams(ctx context.Context) ([]LoggedModelParameter_SdkV2, bool) {
	if m.Params.IsNull() || m.Params.IsUnknown() {
		return nil, false
	}
	var v []LoggedModelParameter_SdkV2
	d := m.Params.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParams sets the value of the Params field in LoggedModelData_SdkV2.
func (m *LoggedModelData_SdkV2) SetParams(ctx context.Context, v []LoggedModelParameter_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["params"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Params = types.ListValueMust(t, vs)
}

// A LoggedModelInfo includes logged model attributes, tags, and registration
// info.
type LoggedModelInfo_SdkV2 struct {
	// The URI of the directory where model artifacts are stored.
	ArtifactUri types.String `tfsdk:"artifact_uri"`
	// The timestamp when the model was created in milliseconds since the UNIX
	// epoch.
	CreationTimestampMs types.Int64 `tfsdk:"creation_timestamp_ms"`
	// The ID of the user or principal that created the model.
	CreatorId types.Int64 `tfsdk:"creator_id"`
	// The ID of the experiment that owns the model.
	ExperimentId types.String `tfsdk:"experiment_id"`
	// The timestamp when the model was last updated in milliseconds since the
	// UNIX epoch.
	LastUpdatedTimestampMs types.Int64 `tfsdk:"last_updated_timestamp_ms"`
	// The unique identifier for the logged model.
	ModelId types.String `tfsdk:"model_id"`
	// The type of model, such as ``"Agent"``, ``"Classifier"``, ``"LLM"``.
	ModelType types.String `tfsdk:"model_type"`
	// The name of the model.
	Name types.String `tfsdk:"name"`
	// The ID of the run that created the model.
	SourceRunId types.String `tfsdk:"source_run_id"`
	// The status of whether or not the model is ready for use.
	Status types.String `tfsdk:"status"`
	// Details on the current model status.
	StatusMessage types.String `tfsdk:"status_message"`
	// Mutable string key-value pairs set on the model.
	Tags types.List `tfsdk:"tags"`
}

func (to *LoggedModelInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LoggedModelInfo_SdkV2) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *LoggedModelInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from LoggedModelInfo_SdkV2) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m LoggedModelInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["artifact_uri"] = attrs["artifact_uri"].SetOptional()
	attrs["creation_timestamp_ms"] = attrs["creation_timestamp_ms"].SetOptional()
	attrs["creator_id"] = attrs["creator_id"].SetOptional()
	attrs["experiment_id"] = attrs["experiment_id"].SetOptional()
	attrs["last_updated_timestamp_ms"] = attrs["last_updated_timestamp_ms"].SetOptional()
	attrs["model_id"] = attrs["model_id"].SetOptional()
	attrs["model_type"] = attrs["model_type"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["source_run_id"] = attrs["source_run_id"].SetOptional()
	attrs["status"] = attrs["status"].SetOptional()
	attrs["status_message"] = attrs["status_message"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LoggedModelInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LoggedModelInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(LoggedModelTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LoggedModelInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m LoggedModelInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"artifact_uri":              m.ArtifactUri,
			"creation_timestamp_ms":     m.CreationTimestampMs,
			"creator_id":                m.CreatorId,
			"experiment_id":             m.ExperimentId,
			"last_updated_timestamp_ms": m.LastUpdatedTimestampMs,
			"model_id":                  m.ModelId,
			"model_type":                m.ModelType,
			"name":                      m.Name,
			"source_run_id":             m.SourceRunId,
			"status":                    m.Status,
			"status_message":            m.StatusMessage,
			"tags":                      m.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LoggedModelInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"artifact_uri":              types.StringType,
			"creation_timestamp_ms":     types.Int64Type,
			"creator_id":                types.Int64Type,
			"experiment_id":             types.StringType,
			"last_updated_timestamp_ms": types.Int64Type,
			"model_id":                  types.StringType,
			"model_type":                types.StringType,
			"name":                      types.StringType,
			"source_run_id":             types.StringType,
			"status":                    types.StringType,
			"status_message":            types.StringType,
			"tags": basetypes.ListType{
				ElemType: LoggedModelTag_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetTags returns the value of the Tags field in LoggedModelInfo_SdkV2 as
// a slice of LoggedModelTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *LoggedModelInfo_SdkV2) GetTags(ctx context.Context) ([]LoggedModelTag_SdkV2, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []LoggedModelTag_SdkV2
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in LoggedModelInfo_SdkV2.
func (m *LoggedModelInfo_SdkV2) SetTags(ctx context.Context, v []LoggedModelTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

// Parameter associated with a LoggedModel.
type LoggedModelParameter_SdkV2 struct {
	// The key identifying this param.
	Key types.String `tfsdk:"key"`
	// The value of this param.
	Value types.String `tfsdk:"value"`
}

func (to *LoggedModelParameter_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LoggedModelParameter_SdkV2) {
}

func (to *LoggedModelParameter_SdkV2) SyncFieldsDuringRead(ctx context.Context, from LoggedModelParameter_SdkV2) {
}

func (m LoggedModelParameter_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetOptional()
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LoggedModelParameter.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LoggedModelParameter_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LoggedModelParameter_SdkV2
// only implements ToObjectValue() and Type().
func (m LoggedModelParameter_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LoggedModelParameter_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

// Tag for a LoggedModel.
type LoggedModelTag_SdkV2 struct {
	// The tag key.
	Key types.String `tfsdk:"key"`
	// The tag value.
	Value types.String `tfsdk:"value"`
}

func (to *LoggedModelTag_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LoggedModelTag_SdkV2) {
}

func (to *LoggedModelTag_SdkV2) SyncFieldsDuringRead(ctx context.Context, from LoggedModelTag_SdkV2) {
}

func (m LoggedModelTag_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetOptional()
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LoggedModelTag.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LoggedModelTag_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LoggedModelTag_SdkV2
// only implements ToObjectValue() and Type().
func (m LoggedModelTag_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LoggedModelTag_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

// Metric associated with a run, represented as a key-value pair.
type Metric_SdkV2 struct {
	// The dataset digest of the dataset associated with the metric, e.g. an md5
	// hash of the dataset that uniquely identifies it within datasets of the
	// same name.
	DatasetDigest types.String `tfsdk:"dataset_digest"`
	// The name of the dataset associated with the metric. E.g.
	// “my.uc.table@2” “nyc-taxi-dataset”, “fantastic-elk-3”
	DatasetName types.String `tfsdk:"dataset_name"`
	// The key identifying the metric.
	Key types.String `tfsdk:"key"`
	// The ID of the logged model or registered model version associated with
	// the metric, if applicable.
	ModelId types.String `tfsdk:"model_id"`
	// The ID of the run containing the metric.
	RunId types.String `tfsdk:"run_id"`
	// The step at which the metric was logged.
	Step types.Int64 `tfsdk:"step"`
	// The timestamp at which the metric was recorded.
	Timestamp types.Int64 `tfsdk:"timestamp"`
	// The value of the metric.
	Value types.Float64 `tfsdk:"value"`
}

func (to *Metric_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Metric_SdkV2) {
}

func (to *Metric_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Metric_SdkV2) {
}

func (m Metric_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dataset_digest"] = attrs["dataset_digest"].SetOptional()
	attrs["dataset_name"] = attrs["dataset_name"].SetOptional()
	attrs["key"] = attrs["key"].SetOptional()
	attrs["model_id"] = attrs["model_id"].SetOptional()
	attrs["run_id"] = attrs["run_id"].SetOptional()
	attrs["step"] = attrs["step"].SetOptional()
	attrs["timestamp"] = attrs["timestamp"].SetOptional()
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Metric.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Metric_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Metric_SdkV2
// only implements ToObjectValue() and Type().
func (m Metric_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dataset_digest": m.DatasetDigest,
			"dataset_name":   m.DatasetName,
			"key":            m.Key,
			"model_id":       m.ModelId,
			"run_id":         m.RunId,
			"step":           m.Step,
			"timestamp":      m.Timestamp,
			"value":          m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Metric_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dataset_digest": types.StringType,
			"dataset_name":   types.StringType,
			"key":            types.StringType,
			"model_id":       types.StringType,
			"run_id":         types.StringType,
			"step":           types.Int64Type,
			"timestamp":      types.Int64Type,
			"value":          types.Float64Type,
		},
	}
}

type Model_SdkV2 struct {
	// Timestamp recorded when this `registered_model` was created.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp"`
	// Description of this `registered_model`.
	Description types.String `tfsdk:"description"`
	// Timestamp recorded when metadata for this `registered_model` was last
	// updated.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp"`
	// Collection of latest model versions for each stage. Only contains models
	// with current `READY` status.
	LatestVersions types.List `tfsdk:"latest_versions"`
	// Unique name for the model.
	Name types.String `tfsdk:"name"`
	// Tags: Additional metadata key-value pairs for this `registered_model`.
	Tags types.List `tfsdk:"tags"`
	// User that created this `registered_model`
	UserId types.String `tfsdk:"user_id"`
}

func (to *Model_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Model_SdkV2) {
	if !from.LatestVersions.IsNull() && !from.LatestVersions.IsUnknown() && to.LatestVersions.IsNull() && len(from.LatestVersions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for LatestVersions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.LatestVersions = from.LatestVersions
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *Model_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Model_SdkV2) {
	if !from.LatestVersions.IsNull() && !from.LatestVersions.IsUnknown() && to.LatestVersions.IsNull() && len(from.LatestVersions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for LatestVersions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.LatestVersions = from.LatestVersions
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m Model_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["creation_timestamp"] = attrs["creation_timestamp"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["last_updated_timestamp"] = attrs["last_updated_timestamp"].SetOptional()
	attrs["latest_versions"] = attrs["latest_versions"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()
	attrs["user_id"] = attrs["user_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Model.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Model_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"latest_versions": reflect.TypeOf(ModelVersion_SdkV2{}),
		"tags":            reflect.TypeOf(ModelTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Model_SdkV2
// only implements ToObjectValue() and Type().
func (m Model_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creation_timestamp":     m.CreationTimestamp,
			"description":            m.Description,
			"last_updated_timestamp": m.LastUpdatedTimestamp,
			"latest_versions":        m.LatestVersions,
			"name":                   m.Name,
			"tags":                   m.Tags,
			"user_id":                m.UserId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Model_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creation_timestamp":     types.Int64Type,
			"description":            types.StringType,
			"last_updated_timestamp": types.Int64Type,
			"latest_versions": basetypes.ListType{
				ElemType: ModelVersion_SdkV2{}.Type(ctx),
			},
			"name": types.StringType,
			"tags": basetypes.ListType{
				ElemType: ModelTag_SdkV2{}.Type(ctx),
			},
			"user_id": types.StringType,
		},
	}
}

// GetLatestVersions returns the value of the LatestVersions field in Model_SdkV2 as
// a slice of ModelVersion_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *Model_SdkV2) GetLatestVersions(ctx context.Context) ([]ModelVersion_SdkV2, bool) {
	if m.LatestVersions.IsNull() || m.LatestVersions.IsUnknown() {
		return nil, false
	}
	var v []ModelVersion_SdkV2
	d := m.LatestVersions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLatestVersions sets the value of the LatestVersions field in Model_SdkV2.
func (m *Model_SdkV2) SetLatestVersions(ctx context.Context, v []ModelVersion_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["latest_versions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.LatestVersions = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in Model_SdkV2 as
// a slice of ModelTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *Model_SdkV2) GetTags(ctx context.Context) ([]ModelTag_SdkV2, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []ModelTag_SdkV2
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in Model_SdkV2.
func (m *Model_SdkV2) SetTags(ctx context.Context, v []ModelTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type ModelDatabricks_SdkV2 struct {
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp"`
	// User-specified description for the object.
	Description types.String `tfsdk:"description"`
	// Unique identifier for the object.
	Id types.String `tfsdk:"id"`
	// Last update time of the object, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp"`
	// Array of model versions, each the latest version for its stage.
	LatestVersions types.List `tfsdk:"latest_versions"`
	// Name of the model.
	Name types.String `tfsdk:"name"`
	// Permission level granted for the requesting user on this registered model
	PermissionLevel types.String `tfsdk:"permission_level"`
	// Array of tags associated with the model.
	Tags types.List `tfsdk:"tags"`
	// The username of the user that created the object.
	UserId types.String `tfsdk:"user_id"`
}

func (to *ModelDatabricks_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ModelDatabricks_SdkV2) {
	if !from.LatestVersions.IsNull() && !from.LatestVersions.IsUnknown() && to.LatestVersions.IsNull() && len(from.LatestVersions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for LatestVersions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.LatestVersions = from.LatestVersions
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *ModelDatabricks_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ModelDatabricks_SdkV2) {
	if !from.LatestVersions.IsNull() && !from.LatestVersions.IsUnknown() && to.LatestVersions.IsNull() && len(from.LatestVersions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for LatestVersions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.LatestVersions = from.LatestVersions
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m ModelDatabricks_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["creation_timestamp"] = attrs["creation_timestamp"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["last_updated_timestamp"] = attrs["last_updated_timestamp"].SetOptional()
	attrs["latest_versions"] = attrs["latest_versions"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()
	attrs["user_id"] = attrs["user_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ModelDatabricks.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ModelDatabricks_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"latest_versions": reflect.TypeOf(ModelVersion_SdkV2{}),
		"tags":            reflect.TypeOf(ModelTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelDatabricks_SdkV2
// only implements ToObjectValue() and Type().
func (m ModelDatabricks_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creation_timestamp":     m.CreationTimestamp,
			"description":            m.Description,
			"id":                     m.Id,
			"last_updated_timestamp": m.LastUpdatedTimestamp,
			"latest_versions":        m.LatestVersions,
			"name":                   m.Name,
			"permission_level":       m.PermissionLevel,
			"tags":                   m.Tags,
			"user_id":                m.UserId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ModelDatabricks_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creation_timestamp":     types.Int64Type,
			"description":            types.StringType,
			"id":                     types.StringType,
			"last_updated_timestamp": types.Int64Type,
			"latest_versions": basetypes.ListType{
				ElemType: ModelVersion_SdkV2{}.Type(ctx),
			},
			"name":             types.StringType,
			"permission_level": types.StringType,
			"tags": basetypes.ListType{
				ElemType: ModelTag_SdkV2{}.Type(ctx),
			},
			"user_id": types.StringType,
		},
	}
}

// GetLatestVersions returns the value of the LatestVersions field in ModelDatabricks_SdkV2 as
// a slice of ModelVersion_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ModelDatabricks_SdkV2) GetLatestVersions(ctx context.Context) ([]ModelVersion_SdkV2, bool) {
	if m.LatestVersions.IsNull() || m.LatestVersions.IsUnknown() {
		return nil, false
	}
	var v []ModelVersion_SdkV2
	d := m.LatestVersions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLatestVersions sets the value of the LatestVersions field in ModelDatabricks_SdkV2.
func (m *ModelDatabricks_SdkV2) SetLatestVersions(ctx context.Context, v []ModelVersion_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["latest_versions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.LatestVersions = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in ModelDatabricks_SdkV2 as
// a slice of ModelTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ModelDatabricks_SdkV2) GetTags(ctx context.Context) ([]ModelTag_SdkV2, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []ModelTag_SdkV2
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in ModelDatabricks_SdkV2.
func (m *ModelDatabricks_SdkV2) SetTags(ctx context.Context, v []ModelTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

// Represents a LoggedModel or Registered Model Version input to a Run.
type ModelInput_SdkV2 struct {
	// The unique identifier of the model.
	ModelId types.String `tfsdk:"model_id"`
}

func (to *ModelInput_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ModelInput_SdkV2) {
}

func (to *ModelInput_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ModelInput_SdkV2) {
}

func (m ModelInput_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["model_id"] = attrs["model_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ModelInput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ModelInput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelInput_SdkV2
// only implements ToObjectValue() and Type().
func (m ModelInput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_id": m.ModelId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ModelInput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_id": types.StringType,
		},
	}
}

// Represents a LoggedModel output of a Run.
type ModelOutput_SdkV2 struct {
	// The unique identifier of the model.
	ModelId types.String `tfsdk:"model_id"`
	// The step at which the model was produced.
	Step types.Int64 `tfsdk:"step"`
}

func (to *ModelOutput_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ModelOutput_SdkV2) {
}

func (to *ModelOutput_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ModelOutput_SdkV2) {
}

func (m ModelOutput_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["model_id"] = attrs["model_id"].SetRequired()
	attrs["step"] = attrs["step"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ModelOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ModelOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelOutput_SdkV2
// only implements ToObjectValue() and Type().
func (m ModelOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_id": m.ModelId,
			"step":     m.Step,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ModelOutput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_id": types.StringType,
			"step":     types.Int64Type,
		},
	}
}

// Tag for a registered model
type ModelTag_SdkV2 struct {
	// The tag key.
	Key types.String `tfsdk:"key"`
	// The tag value.
	Value types.String `tfsdk:"value"`
}

func (to *ModelTag_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ModelTag_SdkV2) {
}

func (to *ModelTag_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ModelTag_SdkV2) {
}

func (m ModelTag_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetOptional()
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ModelTag.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ModelTag_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelTag_SdkV2
// only implements ToObjectValue() and Type().
func (m ModelTag_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ModelTag_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

type ModelVersion_SdkV2 struct {
	// Timestamp recorded when this `model_version` was created.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp"`
	// Current stage for this `model_version`.
	CurrentStage types.String `tfsdk:"current_stage"`
	// Description of this `model_version`.
	Description types.String `tfsdk:"description"`
	// Timestamp recorded when metadata for this `model_version` was last
	// updated.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp"`
	// Unique name of the model
	Name types.String `tfsdk:"name"`
	// MLflow run ID used when creating `model_version`, if `source` was
	// generated by an experiment run stored in MLflow tracking server.
	RunId types.String `tfsdk:"run_id"`
	// Run Link: Direct link to the run that generated this version
	RunLink types.String `tfsdk:"run_link"`
	// URI indicating the location of the source model artifacts, used when
	// creating `model_version`
	Source types.String `tfsdk:"source"`
	// Current status of `model_version`
	Status types.String `tfsdk:"status"`
	// Details on current `status`, if it is pending or failed.
	StatusMessage types.String `tfsdk:"status_message"`
	// Tags: Additional metadata key-value pairs for this `model_version`.
	Tags types.List `tfsdk:"tags"`
	// User that created this `model_version`.
	UserId types.String `tfsdk:"user_id"`
	// Model's version number.
	Version types.String `tfsdk:"version"`
}

func (to *ModelVersion_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ModelVersion_SdkV2) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *ModelVersion_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ModelVersion_SdkV2) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m ModelVersion_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["creation_timestamp"] = attrs["creation_timestamp"].SetOptional()
	attrs["current_stage"] = attrs["current_stage"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["last_updated_timestamp"] = attrs["last_updated_timestamp"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["run_id"] = attrs["run_id"].SetOptional()
	attrs["run_link"] = attrs["run_link"].SetOptional()
	attrs["source"] = attrs["source"].SetOptional()
	attrs["status"] = attrs["status"].SetOptional()
	attrs["status_message"] = attrs["status_message"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()
	attrs["user_id"] = attrs["user_id"].SetOptional()
	attrs["version"] = attrs["version"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ModelVersion.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ModelVersion_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(ModelVersionTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelVersion_SdkV2
// only implements ToObjectValue() and Type().
func (m ModelVersion_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creation_timestamp":     m.CreationTimestamp,
			"current_stage":          m.CurrentStage,
			"description":            m.Description,
			"last_updated_timestamp": m.LastUpdatedTimestamp,
			"name":                   m.Name,
			"run_id":                 m.RunId,
			"run_link":               m.RunLink,
			"source":                 m.Source,
			"status":                 m.Status,
			"status_message":         m.StatusMessage,
			"tags":                   m.Tags,
			"user_id":                m.UserId,
			"version":                m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ModelVersion_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creation_timestamp":     types.Int64Type,
			"current_stage":          types.StringType,
			"description":            types.StringType,
			"last_updated_timestamp": types.Int64Type,
			"name":                   types.StringType,
			"run_id":                 types.StringType,
			"run_link":               types.StringType,
			"source":                 types.StringType,
			"status":                 types.StringType,
			"status_message":         types.StringType,
			"tags": basetypes.ListType{
				ElemType: ModelVersionTag_SdkV2{}.Type(ctx),
			},
			"user_id": types.StringType,
			"version": types.StringType,
		},
	}
}

// GetTags returns the value of the Tags field in ModelVersion_SdkV2 as
// a slice of ModelVersionTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ModelVersion_SdkV2) GetTags(ctx context.Context) ([]ModelVersionTag_SdkV2, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []ModelVersionTag_SdkV2
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in ModelVersion_SdkV2.
func (m *ModelVersion_SdkV2) SetTags(ctx context.Context, v []ModelVersionTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type ModelVersionDatabricks_SdkV2 struct {
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp"`

	CurrentStage types.String `tfsdk:"current_stage"`
	// User-specified description for the object.
	Description types.String `tfsdk:"description"`
	// Email Subscription Status: This is the subscription status of the user to
	// the model version Users get subscribed by interacting with the model
	// version.
	EmailSubscriptionStatus types.String `tfsdk:"email_subscription_status"`
	// Feature lineage of `model_version`.
	FeatureList types.List `tfsdk:"feature_list"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp"`
	// Name of the model.
	Name types.String `tfsdk:"name"`
	// Open requests for this `model_versions`. Gap in sequence number is
	// intentional and is done in order to match field sequence numbers of
	// `ModelVersion` proto message
	OpenRequests types.List `tfsdk:"open_requests"`

	PermissionLevel types.String `tfsdk:"permission_level"`
	// Unique identifier for the MLflow tracking run associated with the source
	// model artifacts.
	RunId types.String `tfsdk:"run_id"`
	// URL of the run associated with the model artifacts. This field is set at
	// model version creation time only for model versions whose source run is
	// from a tracking server that is different from the registry server.
	RunLink types.String `tfsdk:"run_link"`
	// URI that indicates the location of the source model artifacts. This is
	// used when creating the model version.
	Source types.String `tfsdk:"source"`

	Status types.String `tfsdk:"status"`
	// Details on the current status, for example why registration failed.
	StatusMessage types.String `tfsdk:"status_message"`
	// Array of tags that are associated with the model version.
	Tags types.List `tfsdk:"tags"`
	// The username of the user that created the object.
	UserId types.String `tfsdk:"user_id"`
	// Version of the model.
	Version types.String `tfsdk:"version"`
}

func (to *ModelVersionDatabricks_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ModelVersionDatabricks_SdkV2) {
	if !from.FeatureList.IsNull() && !from.FeatureList.IsUnknown() {
		if toFeatureList, ok := to.GetFeatureList(ctx); ok {
			if fromFeatureList, ok := from.GetFeatureList(ctx); ok {
				// Recursively sync the fields of FeatureList
				toFeatureList.SyncFieldsDuringCreateOrUpdate(ctx, fromFeatureList)
				to.SetFeatureList(ctx, toFeatureList)
			}
		}
	}
	if !from.OpenRequests.IsNull() && !from.OpenRequests.IsUnknown() && to.OpenRequests.IsNull() && len(from.OpenRequests.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OpenRequests, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OpenRequests = from.OpenRequests
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *ModelVersionDatabricks_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ModelVersionDatabricks_SdkV2) {
	if !from.FeatureList.IsNull() && !from.FeatureList.IsUnknown() {
		if toFeatureList, ok := to.GetFeatureList(ctx); ok {
			if fromFeatureList, ok := from.GetFeatureList(ctx); ok {
				toFeatureList.SyncFieldsDuringRead(ctx, fromFeatureList)
				to.SetFeatureList(ctx, toFeatureList)
			}
		}
	}
	if !from.OpenRequests.IsNull() && !from.OpenRequests.IsUnknown() && to.OpenRequests.IsNull() && len(from.OpenRequests.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OpenRequests, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OpenRequests = from.OpenRequests
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m ModelVersionDatabricks_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["creation_timestamp"] = attrs["creation_timestamp"].SetOptional()
	attrs["current_stage"] = attrs["current_stage"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["email_subscription_status"] = attrs["email_subscription_status"].SetOptional()
	attrs["feature_list"] = attrs["feature_list"].SetOptional()
	attrs["feature_list"] = attrs["feature_list"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["last_updated_timestamp"] = attrs["last_updated_timestamp"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["open_requests"] = attrs["open_requests"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()
	attrs["run_id"] = attrs["run_id"].SetOptional()
	attrs["run_link"] = attrs["run_link"].SetOptional()
	attrs["source"] = attrs["source"].SetOptional()
	attrs["status"] = attrs["status"].SetOptional()
	attrs["status_message"] = attrs["status_message"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()
	attrs["user_id"] = attrs["user_id"].SetOptional()
	attrs["version"] = attrs["version"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ModelVersionDatabricks.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ModelVersionDatabricks_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"feature_list":  reflect.TypeOf(FeatureList_SdkV2{}),
		"open_requests": reflect.TypeOf(Activity_SdkV2{}),
		"tags":          reflect.TypeOf(ModelVersionTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelVersionDatabricks_SdkV2
// only implements ToObjectValue() and Type().
func (m ModelVersionDatabricks_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creation_timestamp":        m.CreationTimestamp,
			"current_stage":             m.CurrentStage,
			"description":               m.Description,
			"email_subscription_status": m.EmailSubscriptionStatus,
			"feature_list":              m.FeatureList,
			"last_updated_timestamp":    m.LastUpdatedTimestamp,
			"name":                      m.Name,
			"open_requests":             m.OpenRequests,
			"permission_level":          m.PermissionLevel,
			"run_id":                    m.RunId,
			"run_link":                  m.RunLink,
			"source":                    m.Source,
			"status":                    m.Status,
			"status_message":            m.StatusMessage,
			"tags":                      m.Tags,
			"user_id":                   m.UserId,
			"version":                   m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ModelVersionDatabricks_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creation_timestamp":        types.Int64Type,
			"current_stage":             types.StringType,
			"description":               types.StringType,
			"email_subscription_status": types.StringType,
			"feature_list": basetypes.ListType{
				ElemType: FeatureList_SdkV2{}.Type(ctx),
			},
			"last_updated_timestamp": types.Int64Type,
			"name":                   types.StringType,
			"open_requests": basetypes.ListType{
				ElemType: Activity_SdkV2{}.Type(ctx),
			},
			"permission_level": types.StringType,
			"run_id":           types.StringType,
			"run_link":         types.StringType,
			"source":           types.StringType,
			"status":           types.StringType,
			"status_message":   types.StringType,
			"tags": basetypes.ListType{
				ElemType: ModelVersionTag_SdkV2{}.Type(ctx),
			},
			"user_id": types.StringType,
			"version": types.StringType,
		},
	}
}

// GetFeatureList returns the value of the FeatureList field in ModelVersionDatabricks_SdkV2 as
// a FeatureList_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ModelVersionDatabricks_SdkV2) GetFeatureList(ctx context.Context) (FeatureList_SdkV2, bool) {
	var e FeatureList_SdkV2
	if m.FeatureList.IsNull() || m.FeatureList.IsUnknown() {
		return e, false
	}
	var v []FeatureList_SdkV2
	d := m.FeatureList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFeatureList sets the value of the FeatureList field in ModelVersionDatabricks_SdkV2.
func (m *ModelVersionDatabricks_SdkV2) SetFeatureList(ctx context.Context, v FeatureList_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["feature_list"]
	m.FeatureList = types.ListValueMust(t, vs)
}

// GetOpenRequests returns the value of the OpenRequests field in ModelVersionDatabricks_SdkV2 as
// a slice of Activity_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ModelVersionDatabricks_SdkV2) GetOpenRequests(ctx context.Context) ([]Activity_SdkV2, bool) {
	if m.OpenRequests.IsNull() || m.OpenRequests.IsUnknown() {
		return nil, false
	}
	var v []Activity_SdkV2
	d := m.OpenRequests.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOpenRequests sets the value of the OpenRequests field in ModelVersionDatabricks_SdkV2.
func (m *ModelVersionDatabricks_SdkV2) SetOpenRequests(ctx context.Context, v []Activity_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["open_requests"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.OpenRequests = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in ModelVersionDatabricks_SdkV2 as
// a slice of ModelVersionTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ModelVersionDatabricks_SdkV2) GetTags(ctx context.Context) ([]ModelVersionTag_SdkV2, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []ModelVersionTag_SdkV2
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in ModelVersionDatabricks_SdkV2.
func (m *ModelVersionDatabricks_SdkV2) SetTags(ctx context.Context, v []ModelVersionTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type ModelVersionTag_SdkV2 struct {
	// The tag key.
	Key types.String `tfsdk:"key"`
	// The tag value.
	Value types.String `tfsdk:"value"`
}

func (to *ModelVersionTag_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ModelVersionTag_SdkV2) {
}

func (to *ModelVersionTag_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ModelVersionTag_SdkV2) {
}

func (m ModelVersionTag_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetOptional()
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ModelVersionTag.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ModelVersionTag_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelVersionTag_SdkV2
// only implements ToObjectValue() and Type().
func (m ModelVersionTag_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ModelVersionTag_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

// An OnlineStore is a logical database instance that stores and serves features
// online.
type OnlineStore_SdkV2 struct {
	// The capacity of the online store. Valid values are "CU_1", "CU_2",
	// "CU_4", "CU_8".
	Capacity types.String `tfsdk:"capacity"`
	// The timestamp when the online store was created.
	CreationTime types.String `tfsdk:"creation_time"`
	// The email of the creator of the online store.
	Creator types.String `tfsdk:"creator"`
	// The name of the online store. This is the unique identifier for the
	// online store.
	Name types.String `tfsdk:"name"`
	// The number of read replicas for the online store. Defaults to 0.
	ReadReplicaCount types.Int64 `tfsdk:"read_replica_count"`
	// The current state of the online store.
	State types.String `tfsdk:"state"`
}

func (to *OnlineStore_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from OnlineStore_SdkV2) {
}

func (to *OnlineStore_SdkV2) SyncFieldsDuringRead(ctx context.Context, from OnlineStore_SdkV2) {
}

func (m OnlineStore_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["capacity"] = attrs["capacity"].SetRequired()
	attrs["creation_time"] = attrs["creation_time"].SetComputed()
	attrs["creator"] = attrs["creator"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["read_replica_count"] = attrs["read_replica_count"].SetOptional()
	attrs["state"] = attrs["state"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in OnlineStore.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m OnlineStore_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, OnlineStore_SdkV2
// only implements ToObjectValue() and Type().
func (m OnlineStore_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"capacity":           m.Capacity,
			"creation_time":      m.CreationTime,
			"creator":            m.Creator,
			"name":               m.Name,
			"read_replica_count": m.ReadReplicaCount,
			"state":              m.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (m OnlineStore_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"capacity":           types.StringType,
			"creation_time":      types.StringType,
			"creator":            types.StringType,
			"name":               types.StringType,
			"read_replica_count": types.Int64Type,
			"state":              types.StringType,
		},
	}
}

// Param associated with a run.
type Param_SdkV2 struct {
	// Key identifying this param.
	Key types.String `tfsdk:"key"`
	// Value associated with this param.
	Value types.String `tfsdk:"value"`
}

func (to *Param_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Param_SdkV2) {
}

func (to *Param_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Param_SdkV2) {
}

func (m Param_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetOptional()
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Param.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Param_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Param_SdkV2
// only implements ToObjectValue() and Type().
func (m Param_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Param_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

type PublishSpec_SdkV2 struct {
	// The name of the target online store.
	OnlineStore types.String `tfsdk:"online_store"`
	// The full three-part (catalog, schema, table) name of the online table.
	OnlineTableName types.String `tfsdk:"online_table_name"`
	// The publish mode of the pipeline that syncs the online table with the
	// source table.
	PublishMode types.String `tfsdk:"publish_mode"`
}

func (to *PublishSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PublishSpec_SdkV2) {
}

func (to *PublishSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PublishSpec_SdkV2) {
}

func (m PublishSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["online_store"] = attrs["online_store"].SetRequired()
	attrs["online_table_name"] = attrs["online_table_name"].SetRequired()
	attrs["publish_mode"] = attrs["publish_mode"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PublishSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PublishSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PublishSpec_SdkV2
// only implements ToObjectValue() and Type().
func (m PublishSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"online_store":      m.OnlineStore,
			"online_table_name": m.OnlineTableName,
			"publish_mode":      m.PublishMode,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PublishSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"online_store":      types.StringType,
			"online_table_name": types.StringType,
			"publish_mode":      types.StringType,
		},
	}
}

type PublishTableRequest_SdkV2 struct {
	// The specification for publishing the online table from the source table.
	PublishSpec types.List `tfsdk:"publish_spec"`
	// The full three-part (catalog, schema, table) name of the source table.
	SourceTableName types.String `tfsdk:"-"`
}

func (to *PublishTableRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PublishTableRequest_SdkV2) {
	if !from.PublishSpec.IsNull() && !from.PublishSpec.IsUnknown() {
		if toPublishSpec, ok := to.GetPublishSpec(ctx); ok {
			if fromPublishSpec, ok := from.GetPublishSpec(ctx); ok {
				// Recursively sync the fields of PublishSpec
				toPublishSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromPublishSpec)
				to.SetPublishSpec(ctx, toPublishSpec)
			}
		}
	}
}

func (to *PublishTableRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PublishTableRequest_SdkV2) {
	if !from.PublishSpec.IsNull() && !from.PublishSpec.IsUnknown() {
		if toPublishSpec, ok := to.GetPublishSpec(ctx); ok {
			if fromPublishSpec, ok := from.GetPublishSpec(ctx); ok {
				toPublishSpec.SyncFieldsDuringRead(ctx, fromPublishSpec)
				to.SetPublishSpec(ctx, toPublishSpec)
			}
		}
	}
}

func (m PublishTableRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["publish_spec"] = attrs["publish_spec"].SetRequired()
	attrs["publish_spec"] = attrs["publish_spec"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["source_table_name"] = attrs["source_table_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PublishTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PublishTableRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"publish_spec": reflect.TypeOf(PublishSpec_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PublishTableRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m PublishTableRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"publish_spec":      m.PublishSpec,
			"source_table_name": m.SourceTableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PublishTableRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"publish_spec": basetypes.ListType{
				ElemType: PublishSpec_SdkV2{}.Type(ctx),
			},
			"source_table_name": types.StringType,
		},
	}
}

// GetPublishSpec returns the value of the PublishSpec field in PublishTableRequest_SdkV2 as
// a PublishSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *PublishTableRequest_SdkV2) GetPublishSpec(ctx context.Context) (PublishSpec_SdkV2, bool) {
	var e PublishSpec_SdkV2
	if m.PublishSpec.IsNull() || m.PublishSpec.IsUnknown() {
		return e, false
	}
	var v []PublishSpec_SdkV2
	d := m.PublishSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPublishSpec sets the value of the PublishSpec field in PublishTableRequest_SdkV2.
func (m *PublishTableRequest_SdkV2) SetPublishSpec(ctx context.Context, v PublishSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["publish_spec"]
	m.PublishSpec = types.ListValueMust(t, vs)
}

type PublishTableResponse_SdkV2 struct {
	// The full three-part (catalog, schema, table) name of the online table.
	OnlineTableName types.String `tfsdk:"online_table_name"`
	// The ID of the pipeline that syncs the online table with the source table.
	PipelineId types.String `tfsdk:"pipeline_id"`
}

func (to *PublishTableResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PublishTableResponse_SdkV2) {
}

func (to *PublishTableResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PublishTableResponse_SdkV2) {
}

func (m PublishTableResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["online_table_name"] = attrs["online_table_name"].SetOptional()
	attrs["pipeline_id"] = attrs["pipeline_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PublishTableResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PublishTableResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PublishTableResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m PublishTableResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"online_table_name": m.OnlineTableName,
			"pipeline_id":       m.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PublishTableResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"online_table_name": types.StringType,
			"pipeline_id":       types.StringType,
		},
	}
}

type RegisteredModelAccessControlRequest_SdkV2 struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`

	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (to *RegisteredModelAccessControlRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RegisteredModelAccessControlRequest_SdkV2) {
}

func (to *RegisteredModelAccessControlRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RegisteredModelAccessControlRequest_SdkV2) {
}

func (m RegisteredModelAccessControlRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RegisteredModelAccessControlRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RegisteredModelAccessControlRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegisteredModelAccessControlRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m RegisteredModelAccessControlRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m RegisteredModelAccessControlRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type RegisteredModelAccessControlResponse_SdkV2 struct {
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

func (to *RegisteredModelAccessControlResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RegisteredModelAccessControlResponse_SdkV2) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (to *RegisteredModelAccessControlResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RegisteredModelAccessControlResponse_SdkV2) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (m RegisteredModelAccessControlResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["all_permissions"] = attrs["all_permissions"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RegisteredModelAccessControlResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RegisteredModelAccessControlResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(RegisteredModelPermission_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegisteredModelAccessControlResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m RegisteredModelAccessControlResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m RegisteredModelAccessControlResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: RegisteredModelPermission_SdkV2{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

// GetAllPermissions returns the value of the AllPermissions field in RegisteredModelAccessControlResponse_SdkV2 as
// a slice of RegisteredModelPermission_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *RegisteredModelAccessControlResponse_SdkV2) GetAllPermissions(ctx context.Context) ([]RegisteredModelPermission_SdkV2, bool) {
	if m.AllPermissions.IsNull() || m.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []RegisteredModelPermission_SdkV2
	d := m.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in RegisteredModelAccessControlResponse_SdkV2.
func (m *RegisteredModelAccessControlResponse_SdkV2) SetAllPermissions(ctx context.Context, v []RegisteredModelPermission_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllPermissions = types.ListValueMust(t, vs)
}

type RegisteredModelPermission_SdkV2 struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *RegisteredModelPermission_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RegisteredModelPermission_SdkV2) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (to *RegisteredModelPermission_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RegisteredModelPermission_SdkV2) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (m RegisteredModelPermission_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["inherited"] = attrs["inherited"].SetOptional()
	attrs["inherited_from_object"] = attrs["inherited_from_object"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RegisteredModelPermission.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RegisteredModelPermission_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegisteredModelPermission_SdkV2
// only implements ToObjectValue() and Type().
func (m RegisteredModelPermission_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             m.Inherited,
			"inherited_from_object": m.InheritedFromObject,
			"permission_level":      m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RegisteredModelPermission_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetInheritedFromObject returns the value of the InheritedFromObject field in RegisteredModelPermission_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *RegisteredModelPermission_SdkV2) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
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

// SetInheritedFromObject sets the value of the InheritedFromObject field in RegisteredModelPermission_SdkV2.
func (m *RegisteredModelPermission_SdkV2) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InheritedFromObject = types.ListValueMust(t, vs)
}

type RegisteredModelPermissions_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (to *RegisteredModelPermissions_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RegisteredModelPermissions_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *RegisteredModelPermissions_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RegisteredModelPermissions_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m RegisteredModelPermissions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["object_id"] = attrs["object_id"].SetOptional()
	attrs["object_type"] = attrs["object_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RegisteredModelPermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RegisteredModelPermissions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(RegisteredModelAccessControlResponse_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegisteredModelPermissions_SdkV2
// only implements ToObjectValue() and Type().
func (m RegisteredModelPermissions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"object_id":           m.ObjectId,
			"object_type":         m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RegisteredModelPermissions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: RegisteredModelAccessControlResponse_SdkV2{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in RegisteredModelPermissions_SdkV2 as
// a slice of RegisteredModelAccessControlResponse_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *RegisteredModelPermissions_SdkV2) GetAccessControlList(ctx context.Context) ([]RegisteredModelAccessControlResponse_SdkV2, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []RegisteredModelAccessControlResponse_SdkV2
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in RegisteredModelPermissions_SdkV2.
func (m *RegisteredModelPermissions_SdkV2) SetAccessControlList(ctx context.Context, v []RegisteredModelAccessControlResponse_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

type RegisteredModelPermissionsDescription_SdkV2 struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *RegisteredModelPermissionsDescription_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RegisteredModelPermissionsDescription_SdkV2) {
}

func (to *RegisteredModelPermissionsDescription_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RegisteredModelPermissionsDescription_SdkV2) {
}

func (m RegisteredModelPermissionsDescription_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RegisteredModelPermissionsDescription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RegisteredModelPermissionsDescription_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegisteredModelPermissionsDescription_SdkV2
// only implements ToObjectValue() and Type().
func (m RegisteredModelPermissionsDescription_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      m.Description,
			"permission_level": m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RegisteredModelPermissionsDescription_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type RegisteredModelPermissionsRequest_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// The registered model for which to get or manage permissions.
	RegisteredModelId types.String `tfsdk:"-"`
}

func (to *RegisteredModelPermissionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RegisteredModelPermissionsRequest_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *RegisteredModelPermissionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RegisteredModelPermissionsRequest_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m RegisteredModelPermissionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["registered_model_id"] = attrs["registered_model_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RegisteredModelPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RegisteredModelPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(RegisteredModelAccessControlRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegisteredModelPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m RegisteredModelPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"registered_model_id": m.RegisteredModelId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RegisteredModelPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: RegisteredModelAccessControlRequest_SdkV2{}.Type(ctx),
			},
			"registered_model_id": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in RegisteredModelPermissionsRequest_SdkV2 as
// a slice of RegisteredModelAccessControlRequest_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *RegisteredModelPermissionsRequest_SdkV2) GetAccessControlList(ctx context.Context) ([]RegisteredModelAccessControlRequest_SdkV2, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []RegisteredModelAccessControlRequest_SdkV2
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in RegisteredModelPermissionsRequest_SdkV2.
func (m *RegisteredModelPermissionsRequest_SdkV2) SetAccessControlList(ctx context.Context, v []RegisteredModelAccessControlRequest_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

type RegistryWebhook_SdkV2 struct {
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp"`
	// User-specified description for the webhook.
	Description types.String `tfsdk:"description"`
	// Events that can trigger a registry webhook: * `MODEL_VERSION_CREATED`: A
	// new model version was created for the associated model.
	//
	// * `MODEL_VERSION_TRANSITIONED_STAGE`: A model version’s stage was
	// changed.
	//
	// * `TRANSITION_REQUEST_CREATED`: A user requested a model version’s
	// stage be transitioned.
	//
	// * `COMMENT_CREATED`: A user wrote a comment on a registered model.
	//
	// * `REGISTERED_MODEL_CREATED`: A new registered model was created. This
	// event type can only be specified for a registry-wide webhook, which can
	// be created by not specifying a model name in the create request.
	//
	// * `MODEL_VERSION_TAG_SET`: A user set a tag on the model version.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_STAGING`: A model version was
	// transitioned to staging.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_PRODUCTION`: A model version was
	// transitioned to production.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_ARCHIVED`: A model version was archived.
	//
	// * `TRANSITION_REQUEST_TO_STAGING_CREATED`: A user requested a model
	// version be transitioned to staging.
	//
	// * `TRANSITION_REQUEST_TO_PRODUCTION_CREATED`: A user requested a model
	// version be transitioned to production.
	//
	// * `TRANSITION_REQUEST_TO_ARCHIVED_CREATED`: A user requested a model
	// version be archived.
	Events types.List `tfsdk:"events"`

	HttpUrlSpec types.List `tfsdk:"http_url_spec"`
	// Webhook ID
	Id types.String `tfsdk:"id"`

	JobSpec types.List `tfsdk:"job_spec"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp"`
	// Name of the model whose events would trigger this webhook.
	ModelName types.String `tfsdk:"model_name"`

	Status types.String `tfsdk:"status"`
}

func (to *RegistryWebhook_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RegistryWebhook_SdkV2) {
	if !from.Events.IsNull() && !from.Events.IsUnknown() && to.Events.IsNull() && len(from.Events.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Events, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Events = from.Events
	}
	if !from.HttpUrlSpec.IsNull() && !from.HttpUrlSpec.IsUnknown() {
		if toHttpUrlSpec, ok := to.GetHttpUrlSpec(ctx); ok {
			if fromHttpUrlSpec, ok := from.GetHttpUrlSpec(ctx); ok {
				// Recursively sync the fields of HttpUrlSpec
				toHttpUrlSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromHttpUrlSpec)
				to.SetHttpUrlSpec(ctx, toHttpUrlSpec)
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
}

func (to *RegistryWebhook_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RegistryWebhook_SdkV2) {
	if !from.Events.IsNull() && !from.Events.IsUnknown() && to.Events.IsNull() && len(from.Events.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Events, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Events = from.Events
	}
	if !from.HttpUrlSpec.IsNull() && !from.HttpUrlSpec.IsUnknown() {
		if toHttpUrlSpec, ok := to.GetHttpUrlSpec(ctx); ok {
			if fromHttpUrlSpec, ok := from.GetHttpUrlSpec(ctx); ok {
				toHttpUrlSpec.SyncFieldsDuringRead(ctx, fromHttpUrlSpec)
				to.SetHttpUrlSpec(ctx, toHttpUrlSpec)
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
}

func (m RegistryWebhook_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["creation_timestamp"] = attrs["creation_timestamp"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["events"] = attrs["events"].SetOptional()
	attrs["http_url_spec"] = attrs["http_url_spec"].SetOptional()
	attrs["http_url_spec"] = attrs["http_url_spec"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["id"] = attrs["id"].SetOptional()
	attrs["job_spec"] = attrs["job_spec"].SetOptional()
	attrs["job_spec"] = attrs["job_spec"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["last_updated_timestamp"] = attrs["last_updated_timestamp"].SetOptional()
	attrs["model_name"] = attrs["model_name"].SetOptional()
	attrs["status"] = attrs["status"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RegistryWebhook.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RegistryWebhook_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"events":        reflect.TypeOf(types.String{}),
		"http_url_spec": reflect.TypeOf(HttpUrlSpecWithoutSecret_SdkV2{}),
		"job_spec":      reflect.TypeOf(JobSpecWithoutSecret_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegistryWebhook_SdkV2
// only implements ToObjectValue() and Type().
func (m RegistryWebhook_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creation_timestamp":     m.CreationTimestamp,
			"description":            m.Description,
			"events":                 m.Events,
			"http_url_spec":          m.HttpUrlSpec,
			"id":                     m.Id,
			"job_spec":               m.JobSpec,
			"last_updated_timestamp": m.LastUpdatedTimestamp,
			"model_name":             m.ModelName,
			"status":                 m.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RegistryWebhook_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creation_timestamp": types.Int64Type,
			"description":        types.StringType,
			"events": basetypes.ListType{
				ElemType: types.StringType,
			},
			"http_url_spec": basetypes.ListType{
				ElemType: HttpUrlSpecWithoutSecret_SdkV2{}.Type(ctx),
			},
			"id": types.StringType,
			"job_spec": basetypes.ListType{
				ElemType: JobSpecWithoutSecret_SdkV2{}.Type(ctx),
			},
			"last_updated_timestamp": types.Int64Type,
			"model_name":             types.StringType,
			"status":                 types.StringType,
		},
	}
}

// GetEvents returns the value of the Events field in RegistryWebhook_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *RegistryWebhook_SdkV2) GetEvents(ctx context.Context) ([]types.String, bool) {
	if m.Events.IsNull() || m.Events.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Events.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEvents sets the value of the Events field in RegistryWebhook_SdkV2.
func (m *RegistryWebhook_SdkV2) SetEvents(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["events"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Events = types.ListValueMust(t, vs)
}

// GetHttpUrlSpec returns the value of the HttpUrlSpec field in RegistryWebhook_SdkV2 as
// a HttpUrlSpecWithoutSecret_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *RegistryWebhook_SdkV2) GetHttpUrlSpec(ctx context.Context) (HttpUrlSpecWithoutSecret_SdkV2, bool) {
	var e HttpUrlSpecWithoutSecret_SdkV2
	if m.HttpUrlSpec.IsNull() || m.HttpUrlSpec.IsUnknown() {
		return e, false
	}
	var v []HttpUrlSpecWithoutSecret_SdkV2
	d := m.HttpUrlSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetHttpUrlSpec sets the value of the HttpUrlSpec field in RegistryWebhook_SdkV2.
func (m *RegistryWebhook_SdkV2) SetHttpUrlSpec(ctx context.Context, v HttpUrlSpecWithoutSecret_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["http_url_spec"]
	m.HttpUrlSpec = types.ListValueMust(t, vs)
}

// GetJobSpec returns the value of the JobSpec field in RegistryWebhook_SdkV2 as
// a JobSpecWithoutSecret_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *RegistryWebhook_SdkV2) GetJobSpec(ctx context.Context) (JobSpecWithoutSecret_SdkV2, bool) {
	var e JobSpecWithoutSecret_SdkV2
	if m.JobSpec.IsNull() || m.JobSpec.IsUnknown() {
		return e, false
	}
	var v []JobSpecWithoutSecret_SdkV2
	d := m.JobSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetJobSpec sets the value of the JobSpec field in RegistryWebhook_SdkV2.
func (m *RegistryWebhook_SdkV2) SetJobSpec(ctx context.Context, v JobSpecWithoutSecret_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["job_spec"]
	m.JobSpec = types.ListValueMust(t, vs)
}

// Details required to identify and reject a model version stage transition
// request.
type RejectTransitionRequest_SdkV2 struct {
	// User-provided comment on the action.
	Comment types.String `tfsdk:"comment"`
	// Name of the model.
	Name types.String `tfsdk:"name"`
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage types.String `tfsdk:"stage"`
	// Version of the model.
	Version types.String `tfsdk:"version"`
}

func (to *RejectTransitionRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RejectTransitionRequest_SdkV2) {
}

func (to *RejectTransitionRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RejectTransitionRequest_SdkV2) {
}

func (m RejectTransitionRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["stage"] = attrs["stage"].SetRequired()
	attrs["version"] = attrs["version"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RejectTransitionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RejectTransitionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RejectTransitionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m RejectTransitionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment": m.Comment,
			"name":    m.Name,
			"stage":   m.Stage,
			"version": m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RejectTransitionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment": types.StringType,
			"name":    types.StringType,
			"stage":   types.StringType,
			"version": types.StringType,
		},
	}
}

type RejectTransitionRequestResponse_SdkV2 struct {
	// New activity generated as a result of this operation.
	Activity types.List `tfsdk:"activity"`
}

func (to *RejectTransitionRequestResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RejectTransitionRequestResponse_SdkV2) {
	if !from.Activity.IsNull() && !from.Activity.IsUnknown() {
		if toActivity, ok := to.GetActivity(ctx); ok {
			if fromActivity, ok := from.GetActivity(ctx); ok {
				// Recursively sync the fields of Activity
				toActivity.SyncFieldsDuringCreateOrUpdate(ctx, fromActivity)
				to.SetActivity(ctx, toActivity)
			}
		}
	}
}

func (to *RejectTransitionRequestResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RejectTransitionRequestResponse_SdkV2) {
	if !from.Activity.IsNull() && !from.Activity.IsUnknown() {
		if toActivity, ok := to.GetActivity(ctx); ok {
			if fromActivity, ok := from.GetActivity(ctx); ok {
				toActivity.SyncFieldsDuringRead(ctx, fromActivity)
				to.SetActivity(ctx, toActivity)
			}
		}
	}
}

func (m RejectTransitionRequestResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["activity"] = attrs["activity"].SetOptional()
	attrs["activity"] = attrs["activity"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RejectTransitionRequestResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RejectTransitionRequestResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"activity": reflect.TypeOf(Activity_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RejectTransitionRequestResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m RejectTransitionRequestResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"activity": m.Activity,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RejectTransitionRequestResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"activity": basetypes.ListType{
				ElemType: Activity_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetActivity returns the value of the Activity field in RejectTransitionRequestResponse_SdkV2 as
// a Activity_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *RejectTransitionRequestResponse_SdkV2) GetActivity(ctx context.Context) (Activity_SdkV2, bool) {
	var e Activity_SdkV2
	if m.Activity.IsNull() || m.Activity.IsUnknown() {
		return e, false
	}
	var v []Activity_SdkV2
	d := m.Activity.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetActivity sets the value of the Activity field in RejectTransitionRequestResponse_SdkV2.
func (m *RejectTransitionRequestResponse_SdkV2) SetActivity(ctx context.Context, v Activity_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["activity"]
	m.Activity = types.ListValueMust(t, vs)
}

type RenameModelRequest_SdkV2 struct {
	// Registered model unique name identifier.
	Name types.String `tfsdk:"name"`
	// If provided, updates the name for this `registered_model`.
	NewName types.String `tfsdk:"new_name"`
}

func (to *RenameModelRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RenameModelRequest_SdkV2) {
}

func (to *RenameModelRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RenameModelRequest_SdkV2) {
}

func (m RenameModelRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["new_name"] = attrs["new_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RenameModelRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RenameModelRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RenameModelRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m RenameModelRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":     m.Name,
			"new_name": m.NewName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RenameModelRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":     types.StringType,
			"new_name": types.StringType,
		},
	}
}

type RenameModelResponse_SdkV2 struct {
	RegisteredModel types.List `tfsdk:"registered_model"`
}

func (to *RenameModelResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RenameModelResponse_SdkV2) {
	if !from.RegisteredModel.IsNull() && !from.RegisteredModel.IsUnknown() {
		if toRegisteredModel, ok := to.GetRegisteredModel(ctx); ok {
			if fromRegisteredModel, ok := from.GetRegisteredModel(ctx); ok {
				// Recursively sync the fields of RegisteredModel
				toRegisteredModel.SyncFieldsDuringCreateOrUpdate(ctx, fromRegisteredModel)
				to.SetRegisteredModel(ctx, toRegisteredModel)
			}
		}
	}
}

func (to *RenameModelResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RenameModelResponse_SdkV2) {
	if !from.RegisteredModel.IsNull() && !from.RegisteredModel.IsUnknown() {
		if toRegisteredModel, ok := to.GetRegisteredModel(ctx); ok {
			if fromRegisteredModel, ok := from.GetRegisteredModel(ctx); ok {
				toRegisteredModel.SyncFieldsDuringRead(ctx, fromRegisteredModel)
				to.SetRegisteredModel(ctx, toRegisteredModel)
			}
		}
	}
}

func (m RenameModelResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["registered_model"] = attrs["registered_model"].SetOptional()
	attrs["registered_model"] = attrs["registered_model"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RenameModelResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RenameModelResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"registered_model": reflect.TypeOf(Model_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RenameModelResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m RenameModelResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"registered_model": m.RegisteredModel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RenameModelResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"registered_model": basetypes.ListType{
				ElemType: Model_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetRegisteredModel returns the value of the RegisteredModel field in RenameModelResponse_SdkV2 as
// a Model_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *RenameModelResponse_SdkV2) GetRegisteredModel(ctx context.Context) (Model_SdkV2, bool) {
	var e Model_SdkV2
	if m.RegisteredModel.IsNull() || m.RegisteredModel.IsUnknown() {
		return e, false
	}
	var v []Model_SdkV2
	d := m.RegisteredModel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRegisteredModel sets the value of the RegisteredModel field in RenameModelResponse_SdkV2.
func (m *RenameModelResponse_SdkV2) SetRegisteredModel(ctx context.Context, v Model_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["registered_model"]
	m.RegisteredModel = types.ListValueMust(t, vs)
}

type RestoreExperiment_SdkV2 struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"experiment_id"`
}

func (to *RestoreExperiment_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RestoreExperiment_SdkV2) {
}

func (to *RestoreExperiment_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RestoreExperiment_SdkV2) {
}

func (m RestoreExperiment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["experiment_id"] = attrs["experiment_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestoreExperiment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RestoreExperiment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestoreExperiment_SdkV2
// only implements ToObjectValue() and Type().
func (m RestoreExperiment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": m.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RestoreExperiment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
		},
	}
}

type RestoreExperimentResponse_SdkV2 struct {
}

func (to *RestoreExperimentResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RestoreExperimentResponse_SdkV2) {
}

func (to *RestoreExperimentResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RestoreExperimentResponse_SdkV2) {
}

func (m RestoreExperimentResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestoreExperimentResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RestoreExperimentResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestoreExperimentResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m RestoreExperimentResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m RestoreExperimentResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type RestoreRun_SdkV2 struct {
	// ID of the run to restore.
	RunId types.String `tfsdk:"run_id"`
}

func (to *RestoreRun_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RestoreRun_SdkV2) {
}

func (to *RestoreRun_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RestoreRun_SdkV2) {
}

func (m RestoreRun_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["run_id"] = attrs["run_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestoreRun.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RestoreRun_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestoreRun_SdkV2
// only implements ToObjectValue() and Type().
func (m RestoreRun_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run_id": m.RunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RestoreRun_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run_id": types.StringType,
		},
	}
}

type RestoreRunResponse_SdkV2 struct {
}

func (to *RestoreRunResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RestoreRunResponse_SdkV2) {
}

func (to *RestoreRunResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RestoreRunResponse_SdkV2) {
}

func (m RestoreRunResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestoreRunResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RestoreRunResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestoreRunResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m RestoreRunResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m RestoreRunResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type RestoreRuns_SdkV2 struct {
	// The ID of the experiment containing the runs to restore.
	ExperimentId types.String `tfsdk:"experiment_id"`
	// An optional positive integer indicating the maximum number of runs to
	// restore. The maximum allowed value for max_runs is 10000.
	MaxRuns types.Int64 `tfsdk:"max_runs"`
	// The minimum deletion timestamp in milliseconds since the UNIX epoch for
	// restoring runs. Only runs deleted no earlier than this timestamp are
	// restored.
	MinTimestampMillis types.Int64 `tfsdk:"min_timestamp_millis"`
}

func (to *RestoreRuns_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RestoreRuns_SdkV2) {
}

func (to *RestoreRuns_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RestoreRuns_SdkV2) {
}

func (m RestoreRuns_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["experiment_id"] = attrs["experiment_id"].SetRequired()
	attrs["max_runs"] = attrs["max_runs"].SetOptional()
	attrs["min_timestamp_millis"] = attrs["min_timestamp_millis"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestoreRuns.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RestoreRuns_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestoreRuns_SdkV2
// only implements ToObjectValue() and Type().
func (m RestoreRuns_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id":        m.ExperimentId,
			"max_runs":             m.MaxRuns,
			"min_timestamp_millis": m.MinTimestampMillis,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RestoreRuns_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id":        types.StringType,
			"max_runs":             types.Int64Type,
			"min_timestamp_millis": types.Int64Type,
		},
	}
}

type RestoreRunsResponse_SdkV2 struct {
	// The number of runs restored.
	RunsRestored types.Int64 `tfsdk:"runs_restored"`
}

func (to *RestoreRunsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RestoreRunsResponse_SdkV2) {
}

func (to *RestoreRunsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RestoreRunsResponse_SdkV2) {
}

func (m RestoreRunsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["runs_restored"] = attrs["runs_restored"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestoreRunsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RestoreRunsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestoreRunsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m RestoreRunsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"runs_restored": m.RunsRestored,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RestoreRunsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"runs_restored": types.Int64Type,
		},
	}
}

// A single run.
type Run_SdkV2 struct {
	// Run data.
	Data types.List `tfsdk:"data"`
	// Run metadata.
	Info types.List `tfsdk:"info"`
	// Run inputs.
	Inputs types.List `tfsdk:"inputs"`
}

func (to *Run_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Run_SdkV2) {
	if !from.Data.IsNull() && !from.Data.IsUnknown() {
		if toData, ok := to.GetData(ctx); ok {
			if fromData, ok := from.GetData(ctx); ok {
				// Recursively sync the fields of Data
				toData.SyncFieldsDuringCreateOrUpdate(ctx, fromData)
				to.SetData(ctx, toData)
			}
		}
	}
	if !from.Info.IsNull() && !from.Info.IsUnknown() {
		if toInfo, ok := to.GetInfo(ctx); ok {
			if fromInfo, ok := from.GetInfo(ctx); ok {
				// Recursively sync the fields of Info
				toInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromInfo)
				to.SetInfo(ctx, toInfo)
			}
		}
	}
	if !from.Inputs.IsNull() && !from.Inputs.IsUnknown() {
		if toInputs, ok := to.GetInputs(ctx); ok {
			if fromInputs, ok := from.GetInputs(ctx); ok {
				// Recursively sync the fields of Inputs
				toInputs.SyncFieldsDuringCreateOrUpdate(ctx, fromInputs)
				to.SetInputs(ctx, toInputs)
			}
		}
	}
}

func (to *Run_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Run_SdkV2) {
	if !from.Data.IsNull() && !from.Data.IsUnknown() {
		if toData, ok := to.GetData(ctx); ok {
			if fromData, ok := from.GetData(ctx); ok {
				toData.SyncFieldsDuringRead(ctx, fromData)
				to.SetData(ctx, toData)
			}
		}
	}
	if !from.Info.IsNull() && !from.Info.IsUnknown() {
		if toInfo, ok := to.GetInfo(ctx); ok {
			if fromInfo, ok := from.GetInfo(ctx); ok {
				toInfo.SyncFieldsDuringRead(ctx, fromInfo)
				to.SetInfo(ctx, toInfo)
			}
		}
	}
	if !from.Inputs.IsNull() && !from.Inputs.IsUnknown() {
		if toInputs, ok := to.GetInputs(ctx); ok {
			if fromInputs, ok := from.GetInputs(ctx); ok {
				toInputs.SyncFieldsDuringRead(ctx, fromInputs)
				to.SetInputs(ctx, toInputs)
			}
		}
	}
}

func (m Run_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["data"] = attrs["data"].SetOptional()
	attrs["data"] = attrs["data"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["info"] = attrs["info"].SetOptional()
	attrs["info"] = attrs["info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["inputs"] = attrs["inputs"].SetOptional()
	attrs["inputs"] = attrs["inputs"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Run.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Run_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data":   reflect.TypeOf(RunData_SdkV2{}),
		"info":   reflect.TypeOf(RunInfo_SdkV2{}),
		"inputs": reflect.TypeOf(RunInputs_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Run_SdkV2
// only implements ToObjectValue() and Type().
func (m Run_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data":   m.Data,
			"info":   m.Info,
			"inputs": m.Inputs,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Run_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"data": basetypes.ListType{
				ElemType: RunData_SdkV2{}.Type(ctx),
			},
			"info": basetypes.ListType{
				ElemType: RunInfo_SdkV2{}.Type(ctx),
			},
			"inputs": basetypes.ListType{
				ElemType: RunInputs_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetData returns the value of the Data field in Run_SdkV2 as
// a RunData_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Run_SdkV2) GetData(ctx context.Context) (RunData_SdkV2, bool) {
	var e RunData_SdkV2
	if m.Data.IsNull() || m.Data.IsUnknown() {
		return e, false
	}
	var v []RunData_SdkV2
	d := m.Data.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetData sets the value of the Data field in Run_SdkV2.
func (m *Run_SdkV2) SetData(ctx context.Context, v RunData_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["data"]
	m.Data = types.ListValueMust(t, vs)
}

// GetInfo returns the value of the Info field in Run_SdkV2 as
// a RunInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Run_SdkV2) GetInfo(ctx context.Context) (RunInfo_SdkV2, bool) {
	var e RunInfo_SdkV2
	if m.Info.IsNull() || m.Info.IsUnknown() {
		return e, false
	}
	var v []RunInfo_SdkV2
	d := m.Info.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInfo sets the value of the Info field in Run_SdkV2.
func (m *Run_SdkV2) SetInfo(ctx context.Context, v RunInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["info"]
	m.Info = types.ListValueMust(t, vs)
}

// GetInputs returns the value of the Inputs field in Run_SdkV2 as
// a RunInputs_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Run_SdkV2) GetInputs(ctx context.Context) (RunInputs_SdkV2, bool) {
	var e RunInputs_SdkV2
	if m.Inputs.IsNull() || m.Inputs.IsUnknown() {
		return e, false
	}
	var v []RunInputs_SdkV2
	d := m.Inputs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInputs sets the value of the Inputs field in Run_SdkV2.
func (m *Run_SdkV2) SetInputs(ctx context.Context, v RunInputs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["inputs"]
	m.Inputs = types.ListValueMust(t, vs)
}

// Run data (metrics, params, and tags).
type RunData_SdkV2 struct {
	// Run metrics.
	Metrics types.List `tfsdk:"metrics"`
	// Run parameters.
	Params types.List `tfsdk:"params"`
	// Additional metadata key-value pairs.
	Tags types.List `tfsdk:"tags"`
}

func (to *RunData_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RunData_SdkV2) {
	if !from.Metrics.IsNull() && !from.Metrics.IsUnknown() && to.Metrics.IsNull() && len(from.Metrics.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Metrics, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Metrics = from.Metrics
	}
	if !from.Params.IsNull() && !from.Params.IsUnknown() && to.Params.IsNull() && len(from.Params.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Params, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Params = from.Params
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *RunData_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RunData_SdkV2) {
	if !from.Metrics.IsNull() && !from.Metrics.IsUnknown() && to.Metrics.IsNull() && len(from.Metrics.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Metrics, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Metrics = from.Metrics
	}
	if !from.Params.IsNull() && !from.Params.IsUnknown() && to.Params.IsNull() && len(from.Params.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Params, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Params = from.Params
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m RunData_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["metrics"] = attrs["metrics"].SetOptional()
	attrs["params"] = attrs["params"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RunData.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RunData_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metrics": reflect.TypeOf(Metric_SdkV2{}),
		"params":  reflect.TypeOf(Param_SdkV2{}),
		"tags":    reflect.TypeOf(RunTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RunData_SdkV2
// only implements ToObjectValue() and Type().
func (m RunData_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metrics": m.Metrics,
			"params":  m.Params,
			"tags":    m.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RunData_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metrics": basetypes.ListType{
				ElemType: Metric_SdkV2{}.Type(ctx),
			},
			"params": basetypes.ListType{
				ElemType: Param_SdkV2{}.Type(ctx),
			},
			"tags": basetypes.ListType{
				ElemType: RunTag_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetMetrics returns the value of the Metrics field in RunData_SdkV2 as
// a slice of Metric_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *RunData_SdkV2) GetMetrics(ctx context.Context) ([]Metric_SdkV2, bool) {
	if m.Metrics.IsNull() || m.Metrics.IsUnknown() {
		return nil, false
	}
	var v []Metric_SdkV2
	d := m.Metrics.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMetrics sets the value of the Metrics field in RunData_SdkV2.
func (m *RunData_SdkV2) SetMetrics(ctx context.Context, v []Metric_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["metrics"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Metrics = types.ListValueMust(t, vs)
}

// GetParams returns the value of the Params field in RunData_SdkV2 as
// a slice of Param_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *RunData_SdkV2) GetParams(ctx context.Context) ([]Param_SdkV2, bool) {
	if m.Params.IsNull() || m.Params.IsUnknown() {
		return nil, false
	}
	var v []Param_SdkV2
	d := m.Params.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParams sets the value of the Params field in RunData_SdkV2.
func (m *RunData_SdkV2) SetParams(ctx context.Context, v []Param_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["params"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Params = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in RunData_SdkV2 as
// a slice of RunTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *RunData_SdkV2) GetTags(ctx context.Context) ([]RunTag_SdkV2, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []RunTag_SdkV2
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in RunData_SdkV2.
func (m *RunData_SdkV2) SetTags(ctx context.Context, v []RunTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

// Metadata of a single run.
type RunInfo_SdkV2 struct {
	// URI of the directory where artifacts should be uploaded. This can be a
	// local path (starting with "/"), or a distributed file system (DFS) path,
	// like ``s3://bucket/directory`` or ``dbfs:/my/directory``. If not set, the
	// local ``./mlruns`` directory is chosen.
	ArtifactUri types.String `tfsdk:"artifact_uri"`
	// Unix timestamp of when the run ended in milliseconds.
	EndTime types.Int64 `tfsdk:"end_time"`
	// The experiment ID.
	ExperimentId types.String `tfsdk:"experiment_id"`
	// Current life cycle stage of the experiment : OneOf("active", "deleted")
	LifecycleStage types.String `tfsdk:"lifecycle_stage"`
	// Unique identifier for the run.
	RunId types.String `tfsdk:"run_id"`
	// The name of the run.
	RunName types.String `tfsdk:"run_name"`
	// [Deprecated, use run_id instead] Unique identifier for the run. This
	// field will be removed in a future MLflow version.
	RunUuid types.String `tfsdk:"run_uuid"`
	// Unix timestamp of when the run started in milliseconds.
	StartTime types.Int64 `tfsdk:"start_time"`
	// Current status of the run.
	Status types.String `tfsdk:"status"`
	// User who initiated the run. This field is deprecated as of MLflow 1.0,
	// and will be removed in a future MLflow release. Use 'mlflow.user' tag
	// instead.
	UserId types.String `tfsdk:"user_id"`
}

func (to *RunInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RunInfo_SdkV2) {
}

func (to *RunInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RunInfo_SdkV2) {
}

func (m RunInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["artifact_uri"] = attrs["artifact_uri"].SetOptional()
	attrs["end_time"] = attrs["end_time"].SetOptional()
	attrs["experiment_id"] = attrs["experiment_id"].SetOptional()
	attrs["lifecycle_stage"] = attrs["lifecycle_stage"].SetOptional()
	attrs["run_id"] = attrs["run_id"].SetOptional()
	attrs["run_name"] = attrs["run_name"].SetOptional()
	attrs["run_uuid"] = attrs["run_uuid"].SetOptional()
	attrs["start_time"] = attrs["start_time"].SetOptional()
	attrs["status"] = attrs["status"].SetOptional()
	attrs["user_id"] = attrs["user_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RunInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RunInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RunInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m RunInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"artifact_uri":    m.ArtifactUri,
			"end_time":        m.EndTime,
			"experiment_id":   m.ExperimentId,
			"lifecycle_stage": m.LifecycleStage,
			"run_id":          m.RunId,
			"run_name":        m.RunName,
			"run_uuid":        m.RunUuid,
			"start_time":      m.StartTime,
			"status":          m.Status,
			"user_id":         m.UserId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RunInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"artifact_uri":    types.StringType,
			"end_time":        types.Int64Type,
			"experiment_id":   types.StringType,
			"lifecycle_stage": types.StringType,
			"run_id":          types.StringType,
			"run_name":        types.StringType,
			"run_uuid":        types.StringType,
			"start_time":      types.Int64Type,
			"status":          types.StringType,
			"user_id":         types.StringType,
		},
	}
}

// Run inputs.
type RunInputs_SdkV2 struct {
	// Run metrics.
	DatasetInputs types.List `tfsdk:"dataset_inputs"`
	// Model inputs to the Run.
	ModelInputs types.List `tfsdk:"model_inputs"`
}

func (to *RunInputs_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RunInputs_SdkV2) {
	if !from.DatasetInputs.IsNull() && !from.DatasetInputs.IsUnknown() && to.DatasetInputs.IsNull() && len(from.DatasetInputs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DatasetInputs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DatasetInputs = from.DatasetInputs
	}
	if !from.ModelInputs.IsNull() && !from.ModelInputs.IsUnknown() && to.ModelInputs.IsNull() && len(from.ModelInputs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ModelInputs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ModelInputs = from.ModelInputs
	}
}

func (to *RunInputs_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RunInputs_SdkV2) {
	if !from.DatasetInputs.IsNull() && !from.DatasetInputs.IsUnknown() && to.DatasetInputs.IsNull() && len(from.DatasetInputs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DatasetInputs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DatasetInputs = from.DatasetInputs
	}
	if !from.ModelInputs.IsNull() && !from.ModelInputs.IsUnknown() && to.ModelInputs.IsNull() && len(from.ModelInputs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ModelInputs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ModelInputs = from.ModelInputs
	}
}

func (m RunInputs_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dataset_inputs"] = attrs["dataset_inputs"].SetOptional()
	attrs["model_inputs"] = attrs["model_inputs"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RunInputs.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RunInputs_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dataset_inputs": reflect.TypeOf(DatasetInput_SdkV2{}),
		"model_inputs":   reflect.TypeOf(ModelInput_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RunInputs_SdkV2
// only implements ToObjectValue() and Type().
func (m RunInputs_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dataset_inputs": m.DatasetInputs,
			"model_inputs":   m.ModelInputs,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RunInputs_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dataset_inputs": basetypes.ListType{
				ElemType: DatasetInput_SdkV2{}.Type(ctx),
			},
			"model_inputs": basetypes.ListType{
				ElemType: ModelInput_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetDatasetInputs returns the value of the DatasetInputs field in RunInputs_SdkV2 as
// a slice of DatasetInput_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *RunInputs_SdkV2) GetDatasetInputs(ctx context.Context) ([]DatasetInput_SdkV2, bool) {
	if m.DatasetInputs.IsNull() || m.DatasetInputs.IsUnknown() {
		return nil, false
	}
	var v []DatasetInput_SdkV2
	d := m.DatasetInputs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatasetInputs sets the value of the DatasetInputs field in RunInputs_SdkV2.
func (m *RunInputs_SdkV2) SetDatasetInputs(ctx context.Context, v []DatasetInput_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["dataset_inputs"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DatasetInputs = types.ListValueMust(t, vs)
}

// GetModelInputs returns the value of the ModelInputs field in RunInputs_SdkV2 as
// a slice of ModelInput_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *RunInputs_SdkV2) GetModelInputs(ctx context.Context) ([]ModelInput_SdkV2, bool) {
	if m.ModelInputs.IsNull() || m.ModelInputs.IsUnknown() {
		return nil, false
	}
	var v []ModelInput_SdkV2
	d := m.ModelInputs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModelInputs sets the value of the ModelInputs field in RunInputs_SdkV2.
func (m *RunInputs_SdkV2) SetModelInputs(ctx context.Context, v []ModelInput_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["model_inputs"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ModelInputs = types.ListValueMust(t, vs)
}

// Tag for a run.
type RunTag_SdkV2 struct {
	// The tag key.
	Key types.String `tfsdk:"key"`
	// The tag value.
	Value types.String `tfsdk:"value"`
}

func (to *RunTag_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RunTag_SdkV2) {
}

func (to *RunTag_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RunTag_SdkV2) {
}

func (m RunTag_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetOptional()
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RunTag.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RunTag_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RunTag_SdkV2
// only implements ToObjectValue() and Type().
func (m RunTag_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RunTag_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

type SearchExperiments_SdkV2 struct {
	// String representing a SQL filter condition (e.g. "name ILIKE
	// 'my-experiment%'")
	Filter types.String `tfsdk:"filter"`
	// Maximum number of experiments desired. Max threshold is 3000.
	MaxResults types.Int64 `tfsdk:"max_results"`
	// List of columns for ordering search results, which can include experiment
	// name and last updated timestamp with an optional "DESC" or "ASC"
	// annotation, where "ASC" is the default. Tiebreaks are done by experiment
	// id DESC.
	OrderBy types.List `tfsdk:"order_by"`
	// Token indicating the page of experiments to fetch
	PageToken types.String `tfsdk:"page_token"`
	// Qualifier for type of experiments to be returned. If unspecified, return
	// only active experiments.
	ViewType types.String `tfsdk:"view_type"`
}

func (to *SearchExperiments_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SearchExperiments_SdkV2) {
	if !from.OrderBy.IsNull() && !from.OrderBy.IsUnknown() && to.OrderBy.IsNull() && len(from.OrderBy.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OrderBy, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OrderBy = from.OrderBy
	}
}

func (to *SearchExperiments_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SearchExperiments_SdkV2) {
	if !from.OrderBy.IsNull() && !from.OrderBy.IsUnknown() && to.OrderBy.IsNull() && len(from.OrderBy.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OrderBy, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OrderBy = from.OrderBy
	}
}

func (m SearchExperiments_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["filter"] = attrs["filter"].SetOptional()
	attrs["max_results"] = attrs["max_results"].SetOptional()
	attrs["order_by"] = attrs["order_by"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["view_type"] = attrs["view_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SearchExperiments.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SearchExperiments_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"order_by": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchExperiments_SdkV2
// only implements ToObjectValue() and Type().
func (m SearchExperiments_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter":      m.Filter,
			"max_results": m.MaxResults,
			"order_by":    m.OrderBy,
			"page_token":  m.PageToken,
			"view_type":   m.ViewType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SearchExperiments_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter":      types.StringType,
			"max_results": types.Int64Type,
			"order_by": basetypes.ListType{
				ElemType: types.StringType,
			},
			"page_token": types.StringType,
			"view_type":  types.StringType,
		},
	}
}

// GetOrderBy returns the value of the OrderBy field in SearchExperiments_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *SearchExperiments_SdkV2) GetOrderBy(ctx context.Context) ([]types.String, bool) {
	if m.OrderBy.IsNull() || m.OrderBy.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.OrderBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOrderBy sets the value of the OrderBy field in SearchExperiments_SdkV2.
func (m *SearchExperiments_SdkV2) SetOrderBy(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["order_by"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.OrderBy = types.ListValueMust(t, vs)
}

type SearchExperimentsResponse_SdkV2 struct {
	// Experiments that match the search criteria
	Experiments types.List `tfsdk:"experiments"`
	// Token that can be used to retrieve the next page of experiments. An empty
	// token means that no more experiments are available for retrieval.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *SearchExperimentsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SearchExperimentsResponse_SdkV2) {
	if !from.Experiments.IsNull() && !from.Experiments.IsUnknown() && to.Experiments.IsNull() && len(from.Experiments.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Experiments, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Experiments = from.Experiments
	}
}

func (to *SearchExperimentsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SearchExperimentsResponse_SdkV2) {
	if !from.Experiments.IsNull() && !from.Experiments.IsUnknown() && to.Experiments.IsNull() && len(from.Experiments.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Experiments, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Experiments = from.Experiments
	}
}

func (m SearchExperimentsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["experiments"] = attrs["experiments"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SearchExperimentsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SearchExperimentsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"experiments": reflect.TypeOf(Experiment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchExperimentsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m SearchExperimentsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiments":     m.Experiments,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SearchExperimentsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiments": basetypes.ListType{
				ElemType: Experiment_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetExperiments returns the value of the Experiments field in SearchExperimentsResponse_SdkV2 as
// a slice of Experiment_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *SearchExperimentsResponse_SdkV2) GetExperiments(ctx context.Context) ([]Experiment_SdkV2, bool) {
	if m.Experiments.IsNull() || m.Experiments.IsUnknown() {
		return nil, false
	}
	var v []Experiment_SdkV2
	d := m.Experiments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExperiments sets the value of the Experiments field in SearchExperimentsResponse_SdkV2.
func (m *SearchExperimentsResponse_SdkV2) SetExperiments(ctx context.Context, v []Experiment_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["experiments"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Experiments = types.ListValueMust(t, vs)
}

type SearchLoggedModelsDataset_SdkV2 struct {
	// The digest of the dataset.
	DatasetDigest types.String `tfsdk:"dataset_digest"`
	// The name of the dataset.
	DatasetName types.String `tfsdk:"dataset_name"`
}

func (to *SearchLoggedModelsDataset_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SearchLoggedModelsDataset_SdkV2) {
}

func (to *SearchLoggedModelsDataset_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SearchLoggedModelsDataset_SdkV2) {
}

func (m SearchLoggedModelsDataset_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dataset_digest"] = attrs["dataset_digest"].SetOptional()
	attrs["dataset_name"] = attrs["dataset_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SearchLoggedModelsDataset.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SearchLoggedModelsDataset_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchLoggedModelsDataset_SdkV2
// only implements ToObjectValue() and Type().
func (m SearchLoggedModelsDataset_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dataset_digest": m.DatasetDigest,
			"dataset_name":   m.DatasetName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SearchLoggedModelsDataset_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dataset_digest": types.StringType,
			"dataset_name":   types.StringType,
		},
	}
}

type SearchLoggedModelsOrderBy_SdkV2 struct {
	// Whether the search results order is ascending or not.
	Ascending types.Bool `tfsdk:"ascending"`
	// If ``field_name`` refers to a metric, this field specifies the digest of
	// the dataset associated with the metric. Only metrics associated with the
	// specified dataset name and digest will be considered for ordering. This
	// field may only be set if ``dataset_name`` is also set.
	DatasetDigest types.String `tfsdk:"dataset_digest"`
	// If ``field_name`` refers to a metric, this field specifies the name of
	// the dataset associated with the metric. Only metrics associated with the
	// specified dataset name will be considered for ordering. This field may
	// only be set if ``field_name`` refers to a metric.
	DatasetName types.String `tfsdk:"dataset_name"`
	// The name of the field to order by, e.g. "metrics.accuracy".
	FieldName types.String `tfsdk:"field_name"`
}

func (to *SearchLoggedModelsOrderBy_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SearchLoggedModelsOrderBy_SdkV2) {
}

func (to *SearchLoggedModelsOrderBy_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SearchLoggedModelsOrderBy_SdkV2) {
}

func (m SearchLoggedModelsOrderBy_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["ascending"] = attrs["ascending"].SetOptional()
	attrs["dataset_digest"] = attrs["dataset_digest"].SetOptional()
	attrs["dataset_name"] = attrs["dataset_name"].SetOptional()
	attrs["field_name"] = attrs["field_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SearchLoggedModelsOrderBy.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SearchLoggedModelsOrderBy_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchLoggedModelsOrderBy_SdkV2
// only implements ToObjectValue() and Type().
func (m SearchLoggedModelsOrderBy_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ascending":      m.Ascending,
			"dataset_digest": m.DatasetDigest,
			"dataset_name":   m.DatasetName,
			"field_name":     m.FieldName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SearchLoggedModelsOrderBy_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ascending":      types.BoolType,
			"dataset_digest": types.StringType,
			"dataset_name":   types.StringType,
			"field_name":     types.StringType,
		},
	}
}

type SearchLoggedModelsRequest_SdkV2 struct {
	// List of datasets on which to apply the metrics filter clauses. For
	// example, a filter with `metrics.accuracy > 0.9` and dataset info with
	// name "test_dataset" means we will return all logged models with accuracy
	// > 0.9 on the test_dataset. Metric values from ANY dataset matching the
	// criteria are considered. If no datasets are specified, then metrics
	// across all datasets are considered in the filter.
	Datasets types.List `tfsdk:"datasets"`
	// The IDs of the experiments in which to search for logged models.
	ExperimentIds types.List `tfsdk:"experiment_ids"`
	// A filter expression over logged model info and data that allows returning
	// a subset of logged models. The syntax is a subset of SQL that supports
	// AND'ing together binary operations.
	//
	// Example: ``params.alpha < 0.3 AND metrics.accuracy > 0.9``.
	Filter types.String `tfsdk:"filter"`
	// The maximum number of Logged Models to return. The maximum limit is 50.
	MaxResults types.Int64 `tfsdk:"max_results"`
	// The list of columns for ordering the results, with additional fields for
	// sorting criteria.
	OrderBy types.List `tfsdk:"order_by"`
	// The token indicating the page of logged models to fetch.
	PageToken types.String `tfsdk:"page_token"`
}

func (to *SearchLoggedModelsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SearchLoggedModelsRequest_SdkV2) {
	if !from.Datasets.IsNull() && !from.Datasets.IsUnknown() && to.Datasets.IsNull() && len(from.Datasets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Datasets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Datasets = from.Datasets
	}
	if !from.ExperimentIds.IsNull() && !from.ExperimentIds.IsUnknown() && to.ExperimentIds.IsNull() && len(from.ExperimentIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ExperimentIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ExperimentIds = from.ExperimentIds
	}
	if !from.OrderBy.IsNull() && !from.OrderBy.IsUnknown() && to.OrderBy.IsNull() && len(from.OrderBy.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OrderBy, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OrderBy = from.OrderBy
	}
}

func (to *SearchLoggedModelsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SearchLoggedModelsRequest_SdkV2) {
	if !from.Datasets.IsNull() && !from.Datasets.IsUnknown() && to.Datasets.IsNull() && len(from.Datasets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Datasets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Datasets = from.Datasets
	}
	if !from.ExperimentIds.IsNull() && !from.ExperimentIds.IsUnknown() && to.ExperimentIds.IsNull() && len(from.ExperimentIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ExperimentIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ExperimentIds = from.ExperimentIds
	}
	if !from.OrderBy.IsNull() && !from.OrderBy.IsUnknown() && to.OrderBy.IsNull() && len(from.OrderBy.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OrderBy, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OrderBy = from.OrderBy
	}
}

func (m SearchLoggedModelsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["datasets"] = attrs["datasets"].SetOptional()
	attrs["experiment_ids"] = attrs["experiment_ids"].SetOptional()
	attrs["filter"] = attrs["filter"].SetOptional()
	attrs["max_results"] = attrs["max_results"].SetOptional()
	attrs["order_by"] = attrs["order_by"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SearchLoggedModelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SearchLoggedModelsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"datasets":       reflect.TypeOf(SearchLoggedModelsDataset_SdkV2{}),
		"experiment_ids": reflect.TypeOf(types.String{}),
		"order_by":       reflect.TypeOf(SearchLoggedModelsOrderBy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchLoggedModelsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m SearchLoggedModelsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"datasets":       m.Datasets,
			"experiment_ids": m.ExperimentIds,
			"filter":         m.Filter,
			"max_results":    m.MaxResults,
			"order_by":       m.OrderBy,
			"page_token":     m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SearchLoggedModelsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"datasets": basetypes.ListType{
				ElemType: SearchLoggedModelsDataset_SdkV2{}.Type(ctx),
			},
			"experiment_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"filter":      types.StringType,
			"max_results": types.Int64Type,
			"order_by": basetypes.ListType{
				ElemType: SearchLoggedModelsOrderBy_SdkV2{}.Type(ctx),
			},
			"page_token": types.StringType,
		},
	}
}

// GetDatasets returns the value of the Datasets field in SearchLoggedModelsRequest_SdkV2 as
// a slice of SearchLoggedModelsDataset_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *SearchLoggedModelsRequest_SdkV2) GetDatasets(ctx context.Context) ([]SearchLoggedModelsDataset_SdkV2, bool) {
	if m.Datasets.IsNull() || m.Datasets.IsUnknown() {
		return nil, false
	}
	var v []SearchLoggedModelsDataset_SdkV2
	d := m.Datasets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatasets sets the value of the Datasets field in SearchLoggedModelsRequest_SdkV2.
func (m *SearchLoggedModelsRequest_SdkV2) SetDatasets(ctx context.Context, v []SearchLoggedModelsDataset_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["datasets"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Datasets = types.ListValueMust(t, vs)
}

// GetExperimentIds returns the value of the ExperimentIds field in SearchLoggedModelsRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *SearchLoggedModelsRequest_SdkV2) GetExperimentIds(ctx context.Context) ([]types.String, bool) {
	if m.ExperimentIds.IsNull() || m.ExperimentIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.ExperimentIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExperimentIds sets the value of the ExperimentIds field in SearchLoggedModelsRequest_SdkV2.
func (m *SearchLoggedModelsRequest_SdkV2) SetExperimentIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["experiment_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ExperimentIds = types.ListValueMust(t, vs)
}

// GetOrderBy returns the value of the OrderBy field in SearchLoggedModelsRequest_SdkV2 as
// a slice of SearchLoggedModelsOrderBy_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *SearchLoggedModelsRequest_SdkV2) GetOrderBy(ctx context.Context) ([]SearchLoggedModelsOrderBy_SdkV2, bool) {
	if m.OrderBy.IsNull() || m.OrderBy.IsUnknown() {
		return nil, false
	}
	var v []SearchLoggedModelsOrderBy_SdkV2
	d := m.OrderBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOrderBy sets the value of the OrderBy field in SearchLoggedModelsRequest_SdkV2.
func (m *SearchLoggedModelsRequest_SdkV2) SetOrderBy(ctx context.Context, v []SearchLoggedModelsOrderBy_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["order_by"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.OrderBy = types.ListValueMust(t, vs)
}

type SearchLoggedModelsResponse_SdkV2 struct {
	// Logged models that match the search criteria.
	Models types.List `tfsdk:"models"`
	// The token that can be used to retrieve the next page of logged models.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *SearchLoggedModelsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SearchLoggedModelsResponse_SdkV2) {
	if !from.Models.IsNull() && !from.Models.IsUnknown() && to.Models.IsNull() && len(from.Models.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Models, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Models = from.Models
	}
}

func (to *SearchLoggedModelsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SearchLoggedModelsResponse_SdkV2) {
	if !from.Models.IsNull() && !from.Models.IsUnknown() && to.Models.IsNull() && len(from.Models.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Models, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Models = from.Models
	}
}

func (m SearchLoggedModelsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["models"] = attrs["models"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SearchLoggedModelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SearchLoggedModelsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"models": reflect.TypeOf(LoggedModel_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchLoggedModelsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m SearchLoggedModelsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"models":          m.Models,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SearchLoggedModelsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"models": basetypes.ListType{
				ElemType: LoggedModel_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetModels returns the value of the Models field in SearchLoggedModelsResponse_SdkV2 as
// a slice of LoggedModel_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *SearchLoggedModelsResponse_SdkV2) GetModels(ctx context.Context) ([]LoggedModel_SdkV2, bool) {
	if m.Models.IsNull() || m.Models.IsUnknown() {
		return nil, false
	}
	var v []LoggedModel_SdkV2
	d := m.Models.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModels sets the value of the Models field in SearchLoggedModelsResponse_SdkV2.
func (m *SearchLoggedModelsResponse_SdkV2) SetModels(ctx context.Context, v []LoggedModel_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["models"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Models = types.ListValueMust(t, vs)
}

type SearchModelVersionsRequest_SdkV2 struct {
	// String filter condition, like "name='my-model-name'". Must be a single
	// boolean condition, with string values wrapped in single quotes.
	Filter types.String `tfsdk:"-"`
	// Maximum number of models desired. Max threshold is 10K.
	MaxResults types.Int64 `tfsdk:"-"`
	// List of columns to be ordered by including model name, version, stage
	// with an optional "DESC" or "ASC" annotation, where "ASC" is the default.
	// Tiebreaks are done by latest stage transition timestamp, followed by name
	// ASC, followed by version DESC.
	OrderBy types.List `tfsdk:"-"`
	// Pagination token to go to next page based on previous search query.
	PageToken types.String `tfsdk:"-"`
}

func (to *SearchModelVersionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SearchModelVersionsRequest_SdkV2) {
	if !from.OrderBy.IsNull() && !from.OrderBy.IsUnknown() && to.OrderBy.IsNull() && len(from.OrderBy.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OrderBy, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OrderBy = from.OrderBy
	}
}

func (to *SearchModelVersionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SearchModelVersionsRequest_SdkV2) {
	if !from.OrderBy.IsNull() && !from.OrderBy.IsUnknown() && to.OrderBy.IsNull() && len(from.OrderBy.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OrderBy, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OrderBy = from.OrderBy
	}
}

func (m SearchModelVersionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["filter"] = attrs["filter"].SetOptional()
	attrs["max_results"] = attrs["max_results"].SetOptional()
	attrs["order_by"] = attrs["order_by"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SearchModelVersionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SearchModelVersionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"order_by": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchModelVersionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m SearchModelVersionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter":      m.Filter,
			"max_results": m.MaxResults,
			"order_by":    m.OrderBy,
			"page_token":  m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SearchModelVersionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter":      types.StringType,
			"max_results": types.Int64Type,
			"order_by": basetypes.ListType{
				ElemType: types.StringType,
			},
			"page_token": types.StringType,
		},
	}
}

// GetOrderBy returns the value of the OrderBy field in SearchModelVersionsRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *SearchModelVersionsRequest_SdkV2) GetOrderBy(ctx context.Context) ([]types.String, bool) {
	if m.OrderBy.IsNull() || m.OrderBy.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.OrderBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOrderBy sets the value of the OrderBy field in SearchModelVersionsRequest_SdkV2.
func (m *SearchModelVersionsRequest_SdkV2) SetOrderBy(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["order_by"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.OrderBy = types.ListValueMust(t, vs)
}

type SearchModelVersionsResponse_SdkV2 struct {
	// Models that match the search criteria
	ModelVersions types.List `tfsdk:"model_versions"`
	// Pagination token to request next page of models for the same search
	// query.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *SearchModelVersionsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SearchModelVersionsResponse_SdkV2) {
	if !from.ModelVersions.IsNull() && !from.ModelVersions.IsUnknown() && to.ModelVersions.IsNull() && len(from.ModelVersions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ModelVersions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ModelVersions = from.ModelVersions
	}
}

func (to *SearchModelVersionsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SearchModelVersionsResponse_SdkV2) {
	if !from.ModelVersions.IsNull() && !from.ModelVersions.IsUnknown() && to.ModelVersions.IsNull() && len(from.ModelVersions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ModelVersions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ModelVersions = from.ModelVersions
	}
}

func (m SearchModelVersionsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["model_versions"] = attrs["model_versions"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SearchModelVersionsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SearchModelVersionsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model_versions": reflect.TypeOf(ModelVersion_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchModelVersionsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m SearchModelVersionsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_versions":  m.ModelVersions,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SearchModelVersionsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_versions": basetypes.ListType{
				ElemType: ModelVersion_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetModelVersions returns the value of the ModelVersions field in SearchModelVersionsResponse_SdkV2 as
// a slice of ModelVersion_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *SearchModelVersionsResponse_SdkV2) GetModelVersions(ctx context.Context) ([]ModelVersion_SdkV2, bool) {
	if m.ModelVersions.IsNull() || m.ModelVersions.IsUnknown() {
		return nil, false
	}
	var v []ModelVersion_SdkV2
	d := m.ModelVersions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModelVersions sets the value of the ModelVersions field in SearchModelVersionsResponse_SdkV2.
func (m *SearchModelVersionsResponse_SdkV2) SetModelVersions(ctx context.Context, v []ModelVersion_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["model_versions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ModelVersions = types.ListValueMust(t, vs)
}

type SearchModelsRequest_SdkV2 struct {
	// String filter condition, like "name LIKE 'my-model-name'". Interpreted in
	// the backend automatically as "name LIKE '%my-model-name%'". Single
	// boolean condition, with string values wrapped in single quotes.
	Filter types.String `tfsdk:"-"`
	// Maximum number of models desired. Default is 100. Max threshold is 1000.
	MaxResults types.Int64 `tfsdk:"-"`
	// List of columns for ordering search results, which can include model name
	// and last updated timestamp with an optional "DESC" or "ASC" annotation,
	// where "ASC" is the default. Tiebreaks are done by model name ASC.
	OrderBy types.List `tfsdk:"-"`
	// Pagination token to go to the next page based on a previous search query.
	PageToken types.String `tfsdk:"-"`
}

func (to *SearchModelsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SearchModelsRequest_SdkV2) {
	if !from.OrderBy.IsNull() && !from.OrderBy.IsUnknown() && to.OrderBy.IsNull() && len(from.OrderBy.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OrderBy, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OrderBy = from.OrderBy
	}
}

func (to *SearchModelsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SearchModelsRequest_SdkV2) {
	if !from.OrderBy.IsNull() && !from.OrderBy.IsUnknown() && to.OrderBy.IsNull() && len(from.OrderBy.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OrderBy, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OrderBy = from.OrderBy
	}
}

func (m SearchModelsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["filter"] = attrs["filter"].SetOptional()
	attrs["max_results"] = attrs["max_results"].SetOptional()
	attrs["order_by"] = attrs["order_by"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SearchModelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SearchModelsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"order_by": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchModelsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m SearchModelsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter":      m.Filter,
			"max_results": m.MaxResults,
			"order_by":    m.OrderBy,
			"page_token":  m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SearchModelsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"filter":      types.StringType,
			"max_results": types.Int64Type,
			"order_by": basetypes.ListType{
				ElemType: types.StringType,
			},
			"page_token": types.StringType,
		},
	}
}

// GetOrderBy returns the value of the OrderBy field in SearchModelsRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *SearchModelsRequest_SdkV2) GetOrderBy(ctx context.Context) ([]types.String, bool) {
	if m.OrderBy.IsNull() || m.OrderBy.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.OrderBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOrderBy sets the value of the OrderBy field in SearchModelsRequest_SdkV2.
func (m *SearchModelsRequest_SdkV2) SetOrderBy(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["order_by"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.OrderBy = types.ListValueMust(t, vs)
}

type SearchModelsResponse_SdkV2 struct {
	// Pagination token to request the next page of models.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// Registered Models that match the search criteria.
	RegisteredModels types.List `tfsdk:"registered_models"`
}

func (to *SearchModelsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SearchModelsResponse_SdkV2) {
	if !from.RegisteredModels.IsNull() && !from.RegisteredModels.IsUnknown() && to.RegisteredModels.IsNull() && len(from.RegisteredModels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RegisteredModels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RegisteredModels = from.RegisteredModels
	}
}

func (to *SearchModelsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SearchModelsResponse_SdkV2) {
	if !from.RegisteredModels.IsNull() && !from.RegisteredModels.IsUnknown() && to.RegisteredModels.IsNull() && len(from.RegisteredModels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RegisteredModels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RegisteredModels = from.RegisteredModels
	}
}

func (m SearchModelsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["registered_models"] = attrs["registered_models"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SearchModelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SearchModelsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"registered_models": reflect.TypeOf(Model_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchModelsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m SearchModelsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":   m.NextPageToken,
			"registered_models": m.RegisteredModels,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SearchModelsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"registered_models": basetypes.ListType{
				ElemType: Model_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetRegisteredModels returns the value of the RegisteredModels field in SearchModelsResponse_SdkV2 as
// a slice of Model_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *SearchModelsResponse_SdkV2) GetRegisteredModels(ctx context.Context) ([]Model_SdkV2, bool) {
	if m.RegisteredModels.IsNull() || m.RegisteredModels.IsUnknown() {
		return nil, false
	}
	var v []Model_SdkV2
	d := m.RegisteredModels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRegisteredModels sets the value of the RegisteredModels field in SearchModelsResponse_SdkV2.
func (m *SearchModelsResponse_SdkV2) SetRegisteredModels(ctx context.Context, v []Model_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["registered_models"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.RegisteredModels = types.ListValueMust(t, vs)
}

type SearchRuns_SdkV2 struct {
	// List of experiment IDs to search over.
	ExperimentIds types.List `tfsdk:"experiment_ids"`
	// A filter expression over params, metrics, and tags, that allows returning
	// a subset of runs. The syntax is a subset of SQL that supports ANDing
	// together binary operations between a param, metric, or tag and a
	// constant.
	//
	// Example: `metrics.rmse < 1 and params.model_class = 'LogisticRegression'`
	//
	// You can select columns with special characters (hyphen, space, period,
	// etc.) by using double quotes: `metrics."model class" = 'LinearRegression'
	// and tags."user-name" = 'Tomas'`
	//
	// Supported operators are `=`, `!=`, `>`, `>=`, `<`, and `<=`.
	Filter types.String `tfsdk:"filter"`
	// Maximum number of runs desired. Max threshold is 50000
	MaxResults types.Int64 `tfsdk:"max_results"`
	// List of columns to be ordered by, including attributes, params, metrics,
	// and tags with an optional `"DESC"` or `"ASC"` annotation, where `"ASC"`
	// is the default. Example: `["params.input DESC", "metrics.alpha ASC",
	// "metrics.rmse"]`. Tiebreaks are done by start_time `DESC` followed by
	// `run_id` for runs with the same start time (and this is the default
	// ordering criterion if order_by is not provided).
	OrderBy types.List `tfsdk:"order_by"`
	// Token for the current page of runs.
	PageToken types.String `tfsdk:"page_token"`
	// Whether to display only active, only deleted, or all runs. Defaults to
	// only active runs.
	RunViewType types.String `tfsdk:"run_view_type"`
}

func (to *SearchRuns_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SearchRuns_SdkV2) {
	if !from.ExperimentIds.IsNull() && !from.ExperimentIds.IsUnknown() && to.ExperimentIds.IsNull() && len(from.ExperimentIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ExperimentIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ExperimentIds = from.ExperimentIds
	}
	if !from.OrderBy.IsNull() && !from.OrderBy.IsUnknown() && to.OrderBy.IsNull() && len(from.OrderBy.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OrderBy, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OrderBy = from.OrderBy
	}
}

func (to *SearchRuns_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SearchRuns_SdkV2) {
	if !from.ExperimentIds.IsNull() && !from.ExperimentIds.IsUnknown() && to.ExperimentIds.IsNull() && len(from.ExperimentIds.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ExperimentIds, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ExperimentIds = from.ExperimentIds
	}
	if !from.OrderBy.IsNull() && !from.OrderBy.IsUnknown() && to.OrderBy.IsNull() && len(from.OrderBy.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OrderBy, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OrderBy = from.OrderBy
	}
}

func (m SearchRuns_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["experiment_ids"] = attrs["experiment_ids"].SetOptional()
	attrs["filter"] = attrs["filter"].SetOptional()
	attrs["max_results"] = attrs["max_results"].SetOptional()
	attrs["order_by"] = attrs["order_by"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["run_view_type"] = attrs["run_view_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SearchRuns.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SearchRuns_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"experiment_ids": reflect.TypeOf(types.String{}),
		"order_by":       reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchRuns_SdkV2
// only implements ToObjectValue() and Type().
func (m SearchRuns_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_ids": m.ExperimentIds,
			"filter":         m.Filter,
			"max_results":    m.MaxResults,
			"order_by":       m.OrderBy,
			"page_token":     m.PageToken,
			"run_view_type":  m.RunViewType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SearchRuns_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"filter":      types.StringType,
			"max_results": types.Int64Type,
			"order_by": basetypes.ListType{
				ElemType: types.StringType,
			},
			"page_token":    types.StringType,
			"run_view_type": types.StringType,
		},
	}
}

// GetExperimentIds returns the value of the ExperimentIds field in SearchRuns_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *SearchRuns_SdkV2) GetExperimentIds(ctx context.Context) ([]types.String, bool) {
	if m.ExperimentIds.IsNull() || m.ExperimentIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.ExperimentIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExperimentIds sets the value of the ExperimentIds field in SearchRuns_SdkV2.
func (m *SearchRuns_SdkV2) SetExperimentIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["experiment_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ExperimentIds = types.ListValueMust(t, vs)
}

// GetOrderBy returns the value of the OrderBy field in SearchRuns_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *SearchRuns_SdkV2) GetOrderBy(ctx context.Context) ([]types.String, bool) {
	if m.OrderBy.IsNull() || m.OrderBy.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.OrderBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOrderBy sets the value of the OrderBy field in SearchRuns_SdkV2.
func (m *SearchRuns_SdkV2) SetOrderBy(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["order_by"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.OrderBy = types.ListValueMust(t, vs)
}

type SearchRunsResponse_SdkV2 struct {
	// Token for the next page of runs.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// Runs that match the search criteria.
	Runs types.List `tfsdk:"runs"`
}

func (to *SearchRunsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SearchRunsResponse_SdkV2) {
	if !from.Runs.IsNull() && !from.Runs.IsUnknown() && to.Runs.IsNull() && len(from.Runs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Runs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Runs = from.Runs
	}
}

func (to *SearchRunsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SearchRunsResponse_SdkV2) {
	if !from.Runs.IsNull() && !from.Runs.IsUnknown() && to.Runs.IsNull() && len(from.Runs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Runs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Runs = from.Runs
	}
}

func (m SearchRunsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["runs"] = attrs["runs"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SearchRunsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SearchRunsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"runs": reflect.TypeOf(Run_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchRunsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m SearchRunsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"runs":            m.Runs,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SearchRunsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"runs": basetypes.ListType{
				ElemType: Run_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetRuns returns the value of the Runs field in SearchRunsResponse_SdkV2 as
// a slice of Run_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *SearchRunsResponse_SdkV2) GetRuns(ctx context.Context) ([]Run_SdkV2, bool) {
	if m.Runs.IsNull() || m.Runs.IsUnknown() {
		return nil, false
	}
	var v []Run_SdkV2
	d := m.Runs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRuns sets the value of the Runs field in SearchRunsResponse_SdkV2.
func (m *SearchRunsResponse_SdkV2) SetRuns(ctx context.Context, v []Run_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["runs"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Runs = types.ListValueMust(t, vs)
}

type SetExperimentTag_SdkV2 struct {
	// ID of the experiment under which to log the tag. Must be provided.
	ExperimentId types.String `tfsdk:"experiment_id"`
	// Name of the tag. Keys up to 250 bytes in size are supported.
	Key types.String `tfsdk:"key"`
	// String value of the tag being logged. Values up to 64KB in size are
	// supported.
	Value types.String `tfsdk:"value"`
}

func (to *SetExperimentTag_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SetExperimentTag_SdkV2) {
}

func (to *SetExperimentTag_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SetExperimentTag_SdkV2) {
}

func (m SetExperimentTag_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["experiment_id"] = attrs["experiment_id"].SetRequired()
	attrs["key"] = attrs["key"].SetRequired()
	attrs["value"] = attrs["value"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetExperimentTag.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SetExperimentTag_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetExperimentTag_SdkV2
// only implements ToObjectValue() and Type().
func (m SetExperimentTag_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": m.ExperimentId,
			"key":           m.Key,
			"value":         m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SetExperimentTag_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
			"key":           types.StringType,
			"value":         types.StringType,
		},
	}
}

type SetExperimentTagResponse_SdkV2 struct {
}

func (to *SetExperimentTagResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SetExperimentTagResponse_SdkV2) {
}

func (to *SetExperimentTagResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SetExperimentTagResponse_SdkV2) {
}

func (m SetExperimentTagResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetExperimentTagResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SetExperimentTagResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetExperimentTagResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m SetExperimentTagResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m SetExperimentTagResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type SetLoggedModelTagsRequest_SdkV2 struct {
	// The ID of the logged model to set the tags on.
	ModelId types.String `tfsdk:"-"`
	// The tags to set on the logged model.
	Tags types.List `tfsdk:"tags"`
}

func (to *SetLoggedModelTagsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SetLoggedModelTagsRequest_SdkV2) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *SetLoggedModelTagsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SetLoggedModelTagsRequest_SdkV2) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m SetLoggedModelTagsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["tags"] = attrs["tags"].SetOptional()
	attrs["model_id"] = attrs["model_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetLoggedModelTagsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SetLoggedModelTagsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(LoggedModelTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetLoggedModelTagsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m SetLoggedModelTagsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_id": m.ModelId,
			"tags":     m.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SetLoggedModelTagsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_id": types.StringType,
			"tags": basetypes.ListType{
				ElemType: LoggedModelTag_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetTags returns the value of the Tags field in SetLoggedModelTagsRequest_SdkV2 as
// a slice of LoggedModelTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *SetLoggedModelTagsRequest_SdkV2) GetTags(ctx context.Context) ([]LoggedModelTag_SdkV2, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []LoggedModelTag_SdkV2
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in SetLoggedModelTagsRequest_SdkV2.
func (m *SetLoggedModelTagsRequest_SdkV2) SetTags(ctx context.Context, v []LoggedModelTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type SetLoggedModelTagsResponse_SdkV2 struct {
}

func (to *SetLoggedModelTagsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SetLoggedModelTagsResponse_SdkV2) {
}

func (to *SetLoggedModelTagsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SetLoggedModelTagsResponse_SdkV2) {
}

func (m SetLoggedModelTagsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetLoggedModelTagsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SetLoggedModelTagsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetLoggedModelTagsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m SetLoggedModelTagsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m SetLoggedModelTagsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type SetModelTagRequest_SdkV2 struct {
	// Name of the tag. Maximum size depends on storage backend. If a tag with
	// this name already exists, its preexisting value will be replaced by the
	// specified `value`. All storage backends are guaranteed to support key
	// values up to 250 bytes in size.
	Key types.String `tfsdk:"key"`
	// Unique name of the model.
	Name types.String `tfsdk:"name"`
	// String value of the tag being logged. Maximum size depends on storage
	// backend. All storage backends are guaranteed to support key values up to
	// 5000 bytes in size.
	Value types.String `tfsdk:"value"`
}

func (to *SetModelTagRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SetModelTagRequest_SdkV2) {
}

func (to *SetModelTagRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SetModelTagRequest_SdkV2) {
}

func (m SetModelTagRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["value"] = attrs["value"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetModelTagRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SetModelTagRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetModelTagRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m SetModelTagRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"name":  m.Name,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SetModelTagRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"name":  types.StringType,
			"value": types.StringType,
		},
	}
}

type SetModelTagResponse_SdkV2 struct {
}

func (to *SetModelTagResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SetModelTagResponse_SdkV2) {
}

func (to *SetModelTagResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SetModelTagResponse_SdkV2) {
}

func (m SetModelTagResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetModelTagResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SetModelTagResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetModelTagResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m SetModelTagResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m SetModelTagResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type SetModelVersionTagRequest_SdkV2 struct {
	// Name of the tag. Maximum size depends on storage backend. If a tag with
	// this name already exists, its preexisting value will be replaced by the
	// specified `value`. All storage backends are guaranteed to support key
	// values up to 250 bytes in size.
	Key types.String `tfsdk:"key"`
	// Unique name of the model.
	Name types.String `tfsdk:"name"`
	// String value of the tag being logged. Maximum size depends on storage
	// backend. All storage backends are guaranteed to support key values up to
	// 5000 bytes in size.
	Value types.String `tfsdk:"value"`
	// Model version number.
	Version types.String `tfsdk:"version"`
}

func (to *SetModelVersionTagRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SetModelVersionTagRequest_SdkV2) {
}

func (to *SetModelVersionTagRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SetModelVersionTagRequest_SdkV2) {
}

func (m SetModelVersionTagRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["value"] = attrs["value"].SetRequired()
	attrs["version"] = attrs["version"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetModelVersionTagRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SetModelVersionTagRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetModelVersionTagRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m SetModelVersionTagRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":     m.Key,
			"name":    m.Name,
			"value":   m.Value,
			"version": m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SetModelVersionTagRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":     types.StringType,
			"name":    types.StringType,
			"value":   types.StringType,
			"version": types.StringType,
		},
	}
}

type SetModelVersionTagResponse_SdkV2 struct {
}

func (to *SetModelVersionTagResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SetModelVersionTagResponse_SdkV2) {
}

func (to *SetModelVersionTagResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SetModelVersionTagResponse_SdkV2) {
}

func (m SetModelVersionTagResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetModelVersionTagResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SetModelVersionTagResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetModelVersionTagResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m SetModelVersionTagResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m SetModelVersionTagResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type SetTag_SdkV2 struct {
	// Name of the tag. Keys up to 250 bytes in size are supported.
	Key types.String `tfsdk:"key"`
	// ID of the run under which to log the tag. Must be provided.
	RunId types.String `tfsdk:"run_id"`
	// [Deprecated, use `run_id` instead] ID of the run under which to log the
	// tag. This field will be removed in a future MLflow version.
	RunUuid types.String `tfsdk:"run_uuid"`
	// String value of the tag being logged. Values up to 64KB in size are
	// supported.
	Value types.String `tfsdk:"value"`
}

func (to *SetTag_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SetTag_SdkV2) {
}

func (to *SetTag_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SetTag_SdkV2) {
}

func (m SetTag_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetRequired()
	attrs["run_id"] = attrs["run_id"].SetOptional()
	attrs["run_uuid"] = attrs["run_uuid"].SetOptional()
	attrs["value"] = attrs["value"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetTag.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SetTag_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetTag_SdkV2
// only implements ToObjectValue() and Type().
func (m SetTag_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":      m.Key,
			"run_id":   m.RunId,
			"run_uuid": m.RunUuid,
			"value":    m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SetTag_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":      types.StringType,
			"run_id":   types.StringType,
			"run_uuid": types.StringType,
			"value":    types.StringType,
		},
	}
}

type SetTagResponse_SdkV2 struct {
}

func (to *SetTagResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SetTagResponse_SdkV2) {
}

func (to *SetTagResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SetTagResponse_SdkV2) {
}

func (m SetTagResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetTagResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SetTagResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetTagResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m SetTagResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m SetTagResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Details required to test a registry webhook.
type TestRegistryWebhookRequest_SdkV2 struct {
	// If `event` is specified, the test trigger uses the specified event. If
	// `event` is not specified, the test trigger uses a randomly chosen event
	// associated with the webhook.
	Event types.String `tfsdk:"event"`
	// Webhook ID
	Id types.String `tfsdk:"id"`
}

func (to *TestRegistryWebhookRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TestRegistryWebhookRequest_SdkV2) {
}

func (to *TestRegistryWebhookRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TestRegistryWebhookRequest_SdkV2) {
}

func (m TestRegistryWebhookRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["event"] = attrs["event"].SetOptional()
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TestRegistryWebhookRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m TestRegistryWebhookRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TestRegistryWebhookRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m TestRegistryWebhookRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"event": m.Event,
			"id":    m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TestRegistryWebhookRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"event": types.StringType,
			"id":    types.StringType,
		},
	}
}

type TestRegistryWebhookResponse_SdkV2 struct {
	// Body of the response from the webhook URL
	Body types.String `tfsdk:"body"`
	// Status code returned by the webhook URL
	StatusCode types.Int64 `tfsdk:"status_code"`
}

func (to *TestRegistryWebhookResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TestRegistryWebhookResponse_SdkV2) {
}

func (to *TestRegistryWebhookResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TestRegistryWebhookResponse_SdkV2) {
}

func (m TestRegistryWebhookResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["body"] = attrs["body"].SetOptional()
	attrs["status_code"] = attrs["status_code"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TestRegistryWebhookResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m TestRegistryWebhookResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TestRegistryWebhookResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m TestRegistryWebhookResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"body":        m.Body,
			"status_code": m.StatusCode,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TestRegistryWebhookResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"body":        types.StringType,
			"status_code": types.Int64Type,
		},
	}
}

type TimeWindow_SdkV2 struct {
	// The duration of the time window.
	Duration types.String `tfsdk:"duration"`
	// The offset of the time window.
	Offset types.String `tfsdk:"offset"`
}

func (to *TimeWindow_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TimeWindow_SdkV2) {
}

func (to *TimeWindow_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TimeWindow_SdkV2) {
}

func (m TimeWindow_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["duration"] = attrs["duration"].SetRequired()
	attrs["offset"] = attrs["offset"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TimeWindow.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m TimeWindow_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TimeWindow_SdkV2
// only implements ToObjectValue() and Type().
func (m TimeWindow_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"duration": m.Duration,
			"offset":   m.Offset,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TimeWindow_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"duration": types.StringType,
			"offset":   types.StringType,
		},
	}
}

// Details required to transition a model version's stage.
type TransitionModelVersionStageDatabricks_SdkV2 struct {
	// Specifies whether to archive all current model versions in the target
	// stage.
	ArchiveExistingVersions types.Bool `tfsdk:"archive_existing_versions"`
	// User-provided comment on the action.
	Comment types.String `tfsdk:"comment"`
	// Name of the model.
	Name types.String `tfsdk:"name"`
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage types.String `tfsdk:"stage"`
	// Version of the model.
	Version types.String `tfsdk:"version"`
}

func (to *TransitionModelVersionStageDatabricks_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TransitionModelVersionStageDatabricks_SdkV2) {
}

func (to *TransitionModelVersionStageDatabricks_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TransitionModelVersionStageDatabricks_SdkV2) {
}

func (m TransitionModelVersionStageDatabricks_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["archive_existing_versions"] = attrs["archive_existing_versions"].SetRequired()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["stage"] = attrs["stage"].SetRequired()
	attrs["version"] = attrs["version"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TransitionModelVersionStageDatabricks.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m TransitionModelVersionStageDatabricks_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TransitionModelVersionStageDatabricks_SdkV2
// only implements ToObjectValue() and Type().
func (m TransitionModelVersionStageDatabricks_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"archive_existing_versions": m.ArchiveExistingVersions,
			"comment":                   m.Comment,
			"name":                      m.Name,
			"stage":                     m.Stage,
			"version":                   m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TransitionModelVersionStageDatabricks_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"archive_existing_versions": types.BoolType,
			"comment":                   types.StringType,
			"name":                      types.StringType,
			"stage":                     types.StringType,
			"version":                   types.StringType,
		},
	}
}

// For activities, this contains the activity recorded for the action. For
// comments, this contains the comment details. For transition requests, this
// contains the transition request details.
type TransitionRequest_SdkV2 struct {
	// Array of actions on the activity allowed for the current viewer.
	AvailableActions types.List `tfsdk:"available_actions"`
	// User-provided comment associated with the activity, comment, or
	// transition request.
	Comment types.String `tfsdk:"comment"`
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp"`
	// Target stage of the transition (if the activity is stage transition
	// related). Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	ToStage types.String `tfsdk:"to_stage"`
	// The username of the user that created the object.
	UserId types.String `tfsdk:"user_id"`
}

func (to *TransitionRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TransitionRequest_SdkV2) {
	if !from.AvailableActions.IsNull() && !from.AvailableActions.IsUnknown() && to.AvailableActions.IsNull() && len(from.AvailableActions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AvailableActions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AvailableActions = from.AvailableActions
	}
}

func (to *TransitionRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TransitionRequest_SdkV2) {
	if !from.AvailableActions.IsNull() && !from.AvailableActions.IsUnknown() && to.AvailableActions.IsNull() && len(from.AvailableActions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AvailableActions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AvailableActions = from.AvailableActions
	}
}

func (m TransitionRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["available_actions"] = attrs["available_actions"].SetOptional()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["creation_timestamp"] = attrs["creation_timestamp"].SetOptional()
	attrs["to_stage"] = attrs["to_stage"].SetOptional()
	attrs["user_id"] = attrs["user_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TransitionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m TransitionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"available_actions": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TransitionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m TransitionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"available_actions":  m.AvailableActions,
			"comment":            m.Comment,
			"creation_timestamp": m.CreationTimestamp,
			"to_stage":           m.ToStage,
			"user_id":            m.UserId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TransitionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"available_actions": basetypes.ListType{
				ElemType: types.StringType,
			},
			"comment":            types.StringType,
			"creation_timestamp": types.Int64Type,
			"to_stage":           types.StringType,
			"user_id":            types.StringType,
		},
	}
}

// GetAvailableActions returns the value of the AvailableActions field in TransitionRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *TransitionRequest_SdkV2) GetAvailableActions(ctx context.Context) ([]types.String, bool) {
	if m.AvailableActions.IsNull() || m.AvailableActions.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.AvailableActions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAvailableActions sets the value of the AvailableActions field in TransitionRequest_SdkV2.
func (m *TransitionRequest_SdkV2) SetAvailableActions(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["available_actions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AvailableActions = types.ListValueMust(t, vs)
}

type TransitionStageResponse_SdkV2 struct {
	// Updated model version
	ModelVersionDatabricks types.List `tfsdk:"model_version_databricks"`
}

func (to *TransitionStageResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TransitionStageResponse_SdkV2) {
	if !from.ModelVersionDatabricks.IsNull() && !from.ModelVersionDatabricks.IsUnknown() {
		if toModelVersionDatabricks, ok := to.GetModelVersionDatabricks(ctx); ok {
			if fromModelVersionDatabricks, ok := from.GetModelVersionDatabricks(ctx); ok {
				// Recursively sync the fields of ModelVersionDatabricks
				toModelVersionDatabricks.SyncFieldsDuringCreateOrUpdate(ctx, fromModelVersionDatabricks)
				to.SetModelVersionDatabricks(ctx, toModelVersionDatabricks)
			}
		}
	}
}

func (to *TransitionStageResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TransitionStageResponse_SdkV2) {
	if !from.ModelVersionDatabricks.IsNull() && !from.ModelVersionDatabricks.IsUnknown() {
		if toModelVersionDatabricks, ok := to.GetModelVersionDatabricks(ctx); ok {
			if fromModelVersionDatabricks, ok := from.GetModelVersionDatabricks(ctx); ok {
				toModelVersionDatabricks.SyncFieldsDuringRead(ctx, fromModelVersionDatabricks)
				to.SetModelVersionDatabricks(ctx, toModelVersionDatabricks)
			}
		}
	}
}

func (m TransitionStageResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["model_version_databricks"] = attrs["model_version_databricks"].SetOptional()
	attrs["model_version_databricks"] = attrs["model_version_databricks"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TransitionStageResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m TransitionStageResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model_version_databricks": reflect.TypeOf(ModelVersionDatabricks_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TransitionStageResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m TransitionStageResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_version_databricks": m.ModelVersionDatabricks,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TransitionStageResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_version_databricks": basetypes.ListType{
				ElemType: ModelVersionDatabricks_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetModelVersionDatabricks returns the value of the ModelVersionDatabricks field in TransitionStageResponse_SdkV2 as
// a ModelVersionDatabricks_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *TransitionStageResponse_SdkV2) GetModelVersionDatabricks(ctx context.Context) (ModelVersionDatabricks_SdkV2, bool) {
	var e ModelVersionDatabricks_SdkV2
	if m.ModelVersionDatabricks.IsNull() || m.ModelVersionDatabricks.IsUnknown() {
		return e, false
	}
	var v []ModelVersionDatabricks_SdkV2
	d := m.ModelVersionDatabricks.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetModelVersionDatabricks sets the value of the ModelVersionDatabricks field in TransitionStageResponse_SdkV2.
func (m *TransitionStageResponse_SdkV2) SetModelVersionDatabricks(ctx context.Context, v ModelVersionDatabricks_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["model_version_databricks"]
	m.ModelVersionDatabricks = types.ListValueMust(t, vs)
}

// Details required to edit a comment on a model version.
type UpdateComment_SdkV2 struct {
	// User-provided comment on the action.
	Comment types.String `tfsdk:"comment"`
	// Unique identifier of an activity
	Id types.String `tfsdk:"id"`
}

func (to *UpdateComment_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateComment_SdkV2) {
}

func (to *UpdateComment_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateComment_SdkV2) {
}

func (m UpdateComment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetRequired()
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateComment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateComment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateComment_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateComment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment": m.Comment,
			"id":      m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateComment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment": types.StringType,
			"id":      types.StringType,
		},
	}
}

type UpdateCommentResponse_SdkV2 struct {
	// Updated comment object
	Comment types.List `tfsdk:"comment"`
}

func (to *UpdateCommentResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateCommentResponse_SdkV2) {
	if !from.Comment.IsNull() && !from.Comment.IsUnknown() {
		if toComment, ok := to.GetComment(ctx); ok {
			if fromComment, ok := from.GetComment(ctx); ok {
				// Recursively sync the fields of Comment
				toComment.SyncFieldsDuringCreateOrUpdate(ctx, fromComment)
				to.SetComment(ctx, toComment)
			}
		}
	}
}

func (to *UpdateCommentResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateCommentResponse_SdkV2) {
	if !from.Comment.IsNull() && !from.Comment.IsUnknown() {
		if toComment, ok := to.GetComment(ctx); ok {
			if fromComment, ok := from.GetComment(ctx); ok {
				toComment.SyncFieldsDuringRead(ctx, fromComment)
				to.SetComment(ctx, toComment)
			}
		}
	}
}

func (m UpdateCommentResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["comment"] = attrs["comment"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCommentResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateCommentResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"comment": reflect.TypeOf(CommentObject_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCommentResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateCommentResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment": m.Comment,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateCommentResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment": basetypes.ListType{
				ElemType: CommentObject_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetComment returns the value of the Comment field in UpdateCommentResponse_SdkV2 as
// a CommentObject_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateCommentResponse_SdkV2) GetComment(ctx context.Context) (CommentObject_SdkV2, bool) {
	var e CommentObject_SdkV2
	if m.Comment.IsNull() || m.Comment.IsUnknown() {
		return e, false
	}
	var v []CommentObject_SdkV2
	d := m.Comment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetComment sets the value of the Comment field in UpdateCommentResponse_SdkV2.
func (m *UpdateCommentResponse_SdkV2) SetComment(ctx context.Context, v CommentObject_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["comment"]
	m.Comment = types.ListValueMust(t, vs)
}

type UpdateExperiment_SdkV2 struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"experiment_id"`
	// If provided, the experiment's name is changed to the new name. The new
	// name must be unique.
	NewName types.String `tfsdk:"new_name"`
}

func (to *UpdateExperiment_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateExperiment_SdkV2) {
}

func (to *UpdateExperiment_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateExperiment_SdkV2) {
}

func (m UpdateExperiment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["experiment_id"] = attrs["experiment_id"].SetRequired()
	attrs["new_name"] = attrs["new_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateExperiment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateExperiment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateExperiment_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateExperiment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": m.ExperimentId,
			"new_name":      m.NewName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateExperiment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
			"new_name":      types.StringType,
		},
	}
}

type UpdateExperimentResponse_SdkV2 struct {
}

func (to *UpdateExperimentResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateExperimentResponse_SdkV2) {
}

func (to *UpdateExperimentResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateExperimentResponse_SdkV2) {
}

func (m UpdateExperimentResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateExperimentResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateExperimentResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateExperimentResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateExperimentResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateExperimentResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateFeatureRequest_SdkV2 struct {
	// Feature to update.
	Feature types.List `tfsdk:"feature"`
	// The full three-part name (catalog, schema, name) of the feature.
	FullName types.String `tfsdk:"-"`
	// The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateFeatureRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateFeatureRequest_SdkV2) {
	if !from.Feature.IsNull() && !from.Feature.IsUnknown() {
		if toFeature, ok := to.GetFeature(ctx); ok {
			if fromFeature, ok := from.GetFeature(ctx); ok {
				// Recursively sync the fields of Feature
				toFeature.SyncFieldsDuringCreateOrUpdate(ctx, fromFeature)
				to.SetFeature(ctx, toFeature)
			}
		}
	}
}

func (to *UpdateFeatureRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateFeatureRequest_SdkV2) {
	if !from.Feature.IsNull() && !from.Feature.IsUnknown() {
		if toFeature, ok := to.GetFeature(ctx); ok {
			if fromFeature, ok := from.GetFeature(ctx); ok {
				toFeature.SyncFieldsDuringRead(ctx, fromFeature)
				to.SetFeature(ctx, toFeature)
			}
		}
	}
}

func (m UpdateFeatureRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["feature"] = attrs["feature"].SetRequired()
	attrs["feature"] = attrs["feature"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["full_name"] = attrs["full_name"].SetRequired()
	attrs["full_name"] = attrs["full_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateFeatureRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateFeatureRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"feature": reflect.TypeOf(Feature_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateFeatureRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateFeatureRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature":     m.Feature,
			"full_name":   m.FullName,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateFeatureRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"feature": basetypes.ListType{
				ElemType: Feature_SdkV2{}.Type(ctx),
			},
			"full_name":   types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetFeature returns the value of the Feature field in UpdateFeatureRequest_SdkV2 as
// a Feature_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateFeatureRequest_SdkV2) GetFeature(ctx context.Context) (Feature_SdkV2, bool) {
	var e Feature_SdkV2
	if m.Feature.IsNull() || m.Feature.IsUnknown() {
		return e, false
	}
	var v []Feature_SdkV2
	d := m.Feature.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFeature sets the value of the Feature field in UpdateFeatureRequest_SdkV2.
func (m *UpdateFeatureRequest_SdkV2) SetFeature(ctx context.Context, v Feature_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["feature"]
	m.Feature = types.ListValueMust(t, vs)
}

type UpdateFeatureTagRequest_SdkV2 struct {
	FeatureName types.String `tfsdk:"-"`

	FeatureTag types.List `tfsdk:"feature_tag"`

	Key types.String `tfsdk:"-"`

	TableName types.String `tfsdk:"-"`
	// The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateFeatureTagRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateFeatureTagRequest_SdkV2) {
	if !from.FeatureTag.IsNull() && !from.FeatureTag.IsUnknown() {
		if toFeatureTag, ok := to.GetFeatureTag(ctx); ok {
			if fromFeatureTag, ok := from.GetFeatureTag(ctx); ok {
				// Recursively sync the fields of FeatureTag
				toFeatureTag.SyncFieldsDuringCreateOrUpdate(ctx, fromFeatureTag)
				to.SetFeatureTag(ctx, toFeatureTag)
			}
		}
	}
}

func (to *UpdateFeatureTagRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateFeatureTagRequest_SdkV2) {
	if !from.FeatureTag.IsNull() && !from.FeatureTag.IsUnknown() {
		if toFeatureTag, ok := to.GetFeatureTag(ctx); ok {
			if fromFeatureTag, ok := from.GetFeatureTag(ctx); ok {
				toFeatureTag.SyncFieldsDuringRead(ctx, fromFeatureTag)
				to.SetFeatureTag(ctx, toFeatureTag)
			}
		}
	}
}

func (m UpdateFeatureTagRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["feature_tag"] = attrs["feature_tag"].SetRequired()
	attrs["feature_tag"] = attrs["feature_tag"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["table_name"] = attrs["table_name"].SetRequired()
	attrs["feature_name"] = attrs["feature_name"].SetRequired()
	attrs["key"] = attrs["key"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateFeatureTagRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateFeatureTagRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"feature_tag": reflect.TypeOf(FeatureTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateFeatureTagRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateFeatureTagRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_name": m.FeatureName,
			"feature_tag":  m.FeatureTag,
			"key":          m.Key,
			"table_name":   m.TableName,
			"update_mask":  m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateFeatureTagRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"feature_name": types.StringType,
			"feature_tag": basetypes.ListType{
				ElemType: FeatureTag_SdkV2{}.Type(ctx),
			},
			"key":         types.StringType,
			"table_name":  types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetFeatureTag returns the value of the FeatureTag field in UpdateFeatureTagRequest_SdkV2 as
// a FeatureTag_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateFeatureTagRequest_SdkV2) GetFeatureTag(ctx context.Context) (FeatureTag_SdkV2, bool) {
	var e FeatureTag_SdkV2
	if m.FeatureTag.IsNull() || m.FeatureTag.IsUnknown() {
		return e, false
	}
	var v []FeatureTag_SdkV2
	d := m.FeatureTag.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFeatureTag sets the value of the FeatureTag field in UpdateFeatureTagRequest_SdkV2.
func (m *UpdateFeatureTagRequest_SdkV2) SetFeatureTag(ctx context.Context, v FeatureTag_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["feature_tag"]
	m.FeatureTag = types.ListValueMust(t, vs)
}

type UpdateModelRequest_SdkV2 struct {
	// If provided, updates the description for this `registered_model`.
	Description types.String `tfsdk:"description"`
	// Registered model unique name identifier.
	Name types.String `tfsdk:"name"`
}

func (to *UpdateModelRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateModelRequest_SdkV2) {
}

func (to *UpdateModelRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateModelRequest_SdkV2) {
}

func (m UpdateModelRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateModelRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateModelRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateModelRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateModelRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description": m.Description,
			"name":        m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateModelRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"name":        types.StringType,
		},
	}
}

type UpdateModelResponse_SdkV2 struct {
	RegisteredModel types.List `tfsdk:"registered_model"`
}

func (to *UpdateModelResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateModelResponse_SdkV2) {
	if !from.RegisteredModel.IsNull() && !from.RegisteredModel.IsUnknown() {
		if toRegisteredModel, ok := to.GetRegisteredModel(ctx); ok {
			if fromRegisteredModel, ok := from.GetRegisteredModel(ctx); ok {
				// Recursively sync the fields of RegisteredModel
				toRegisteredModel.SyncFieldsDuringCreateOrUpdate(ctx, fromRegisteredModel)
				to.SetRegisteredModel(ctx, toRegisteredModel)
			}
		}
	}
}

func (to *UpdateModelResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateModelResponse_SdkV2) {
	if !from.RegisteredModel.IsNull() && !from.RegisteredModel.IsUnknown() {
		if toRegisteredModel, ok := to.GetRegisteredModel(ctx); ok {
			if fromRegisteredModel, ok := from.GetRegisteredModel(ctx); ok {
				toRegisteredModel.SyncFieldsDuringRead(ctx, fromRegisteredModel)
				to.SetRegisteredModel(ctx, toRegisteredModel)
			}
		}
	}
}

func (m UpdateModelResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["registered_model"] = attrs["registered_model"].SetOptional()
	attrs["registered_model"] = attrs["registered_model"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateModelResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateModelResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"registered_model": reflect.TypeOf(Model_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateModelResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateModelResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"registered_model": m.RegisteredModel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateModelResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"registered_model": basetypes.ListType{
				ElemType: Model_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetRegisteredModel returns the value of the RegisteredModel field in UpdateModelResponse_SdkV2 as
// a Model_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateModelResponse_SdkV2) GetRegisteredModel(ctx context.Context) (Model_SdkV2, bool) {
	var e Model_SdkV2
	if m.RegisteredModel.IsNull() || m.RegisteredModel.IsUnknown() {
		return e, false
	}
	var v []Model_SdkV2
	d := m.RegisteredModel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRegisteredModel sets the value of the RegisteredModel field in UpdateModelResponse_SdkV2.
func (m *UpdateModelResponse_SdkV2) SetRegisteredModel(ctx context.Context, v Model_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["registered_model"]
	m.RegisteredModel = types.ListValueMust(t, vs)
}

type UpdateModelVersionRequest_SdkV2 struct {
	// If provided, updates the description for this `registered_model`.
	Description types.String `tfsdk:"description"`
	// Name of the registered model
	Name types.String `tfsdk:"name"`
	// Model version number
	Version types.String `tfsdk:"version"`
}

func (to *UpdateModelVersionRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateModelVersionRequest_SdkV2) {
}

func (to *UpdateModelVersionRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateModelVersionRequest_SdkV2) {
}

func (m UpdateModelVersionRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["version"] = attrs["version"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateModelVersionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateModelVersionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateModelVersionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateModelVersionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description": m.Description,
			"name":        m.Name,
			"version":     m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateModelVersionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"name":        types.StringType,
			"version":     types.StringType,
		},
	}
}

type UpdateModelVersionResponse_SdkV2 struct {
	// Return new version number generated for this model in registry.
	ModelVersion types.List `tfsdk:"model_version"`
}

func (to *UpdateModelVersionResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateModelVersionResponse_SdkV2) {
	if !from.ModelVersion.IsNull() && !from.ModelVersion.IsUnknown() {
		if toModelVersion, ok := to.GetModelVersion(ctx); ok {
			if fromModelVersion, ok := from.GetModelVersion(ctx); ok {
				// Recursively sync the fields of ModelVersion
				toModelVersion.SyncFieldsDuringCreateOrUpdate(ctx, fromModelVersion)
				to.SetModelVersion(ctx, toModelVersion)
			}
		}
	}
}

func (to *UpdateModelVersionResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateModelVersionResponse_SdkV2) {
	if !from.ModelVersion.IsNull() && !from.ModelVersion.IsUnknown() {
		if toModelVersion, ok := to.GetModelVersion(ctx); ok {
			if fromModelVersion, ok := from.GetModelVersion(ctx); ok {
				toModelVersion.SyncFieldsDuringRead(ctx, fromModelVersion)
				to.SetModelVersion(ctx, toModelVersion)
			}
		}
	}
}

func (m UpdateModelVersionResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["model_version"] = attrs["model_version"].SetOptional()
	attrs["model_version"] = attrs["model_version"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateModelVersionResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateModelVersionResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model_version": reflect.TypeOf(ModelVersion_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateModelVersionResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateModelVersionResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_version": m.ModelVersion,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateModelVersionResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_version": basetypes.ListType{
				ElemType: ModelVersion_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetModelVersion returns the value of the ModelVersion field in UpdateModelVersionResponse_SdkV2 as
// a ModelVersion_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateModelVersionResponse_SdkV2) GetModelVersion(ctx context.Context) (ModelVersion_SdkV2, bool) {
	var e ModelVersion_SdkV2
	if m.ModelVersion.IsNull() || m.ModelVersion.IsUnknown() {
		return e, false
	}
	var v []ModelVersion_SdkV2
	d := m.ModelVersion.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetModelVersion sets the value of the ModelVersion field in UpdateModelVersionResponse_SdkV2.
func (m *UpdateModelVersionResponse_SdkV2) SetModelVersion(ctx context.Context, v ModelVersion_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["model_version"]
	m.ModelVersion = types.ListValueMust(t, vs)
}

type UpdateOnlineStoreRequest_SdkV2 struct {
	// The name of the online store. This is the unique identifier for the
	// online store.
	Name types.String `tfsdk:"-"`
	// Online store to update.
	OnlineStore types.List `tfsdk:"online_store"`
	// The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateOnlineStoreRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateOnlineStoreRequest_SdkV2) {
	if !from.OnlineStore.IsNull() && !from.OnlineStore.IsUnknown() {
		if toOnlineStore, ok := to.GetOnlineStore(ctx); ok {
			if fromOnlineStore, ok := from.GetOnlineStore(ctx); ok {
				// Recursively sync the fields of OnlineStore
				toOnlineStore.SyncFieldsDuringCreateOrUpdate(ctx, fromOnlineStore)
				to.SetOnlineStore(ctx, toOnlineStore)
			}
		}
	}
}

func (to *UpdateOnlineStoreRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateOnlineStoreRequest_SdkV2) {
	if !from.OnlineStore.IsNull() && !from.OnlineStore.IsUnknown() {
		if toOnlineStore, ok := to.GetOnlineStore(ctx); ok {
			if fromOnlineStore, ok := from.GetOnlineStore(ctx); ok {
				toOnlineStore.SyncFieldsDuringRead(ctx, fromOnlineStore)
				to.SetOnlineStore(ctx, toOnlineStore)
			}
		}
	}
}

func (m UpdateOnlineStoreRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["online_store"] = attrs["online_store"].SetRequired()
	attrs["online_store"] = attrs["online_store"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateOnlineStoreRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateOnlineStoreRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"online_store": reflect.TypeOf(OnlineStore_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateOnlineStoreRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateOnlineStoreRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":         m.Name,
			"online_store": m.OnlineStore,
			"update_mask":  m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateOnlineStoreRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"online_store": basetypes.ListType{
				ElemType: OnlineStore_SdkV2{}.Type(ctx),
			},
			"update_mask": types.StringType,
		},
	}
}

// GetOnlineStore returns the value of the OnlineStore field in UpdateOnlineStoreRequest_SdkV2 as
// a OnlineStore_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateOnlineStoreRequest_SdkV2) GetOnlineStore(ctx context.Context) (OnlineStore_SdkV2, bool) {
	var e OnlineStore_SdkV2
	if m.OnlineStore.IsNull() || m.OnlineStore.IsUnknown() {
		return e, false
	}
	var v []OnlineStore_SdkV2
	d := m.OnlineStore.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOnlineStore sets the value of the OnlineStore field in UpdateOnlineStoreRequest_SdkV2.
func (m *UpdateOnlineStoreRequest_SdkV2) SetOnlineStore(ctx context.Context, v OnlineStore_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["online_store"]
	m.OnlineStore = types.ListValueMust(t, vs)
}

// Details required to update a registry webhook. Only the fields that need to
// be updated should be specified, and both `http_url_spec` and `job_spec`
// should not be specified in the same request.
type UpdateRegistryWebhook_SdkV2 struct {
	// User-specified description for the webhook.
	Description types.String `tfsdk:"description"`
	// Events that can trigger a registry webhook: * `MODEL_VERSION_CREATED`: A
	// new model version was created for the associated model.
	//
	// * `MODEL_VERSION_TRANSITIONED_STAGE`: A model version’s stage was
	// changed.
	//
	// * `TRANSITION_REQUEST_CREATED`: A user requested a model version’s
	// stage be transitioned.
	//
	// * `COMMENT_CREATED`: A user wrote a comment on a registered model.
	//
	// * `REGISTERED_MODEL_CREATED`: A new registered model was created. This
	// event type can only be specified for a registry-wide webhook, which can
	// be created by not specifying a model name in the create request.
	//
	// * `MODEL_VERSION_TAG_SET`: A user set a tag on the model version.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_STAGING`: A model version was
	// transitioned to staging.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_PRODUCTION`: A model version was
	// transitioned to production.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_ARCHIVED`: A model version was archived.
	//
	// * `TRANSITION_REQUEST_TO_STAGING_CREATED`: A user requested a model
	// version be transitioned to staging.
	//
	// * `TRANSITION_REQUEST_TO_PRODUCTION_CREATED`: A user requested a model
	// version be transitioned to production.
	//
	// * `TRANSITION_REQUEST_TO_ARCHIVED_CREATED`: A user requested a model
	// version be archived.
	Events types.List `tfsdk:"events"`

	HttpUrlSpec types.List `tfsdk:"http_url_spec"`
	// Webhook ID
	Id types.String `tfsdk:"id"`

	JobSpec types.List `tfsdk:"job_spec"`

	Status types.String `tfsdk:"status"`
}

func (to *UpdateRegistryWebhook_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateRegistryWebhook_SdkV2) {
	if !from.Events.IsNull() && !from.Events.IsUnknown() && to.Events.IsNull() && len(from.Events.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Events, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Events = from.Events
	}
	if !from.HttpUrlSpec.IsNull() && !from.HttpUrlSpec.IsUnknown() {
		if toHttpUrlSpec, ok := to.GetHttpUrlSpec(ctx); ok {
			if fromHttpUrlSpec, ok := from.GetHttpUrlSpec(ctx); ok {
				// Recursively sync the fields of HttpUrlSpec
				toHttpUrlSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromHttpUrlSpec)
				to.SetHttpUrlSpec(ctx, toHttpUrlSpec)
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
}

func (to *UpdateRegistryWebhook_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateRegistryWebhook_SdkV2) {
	if !from.Events.IsNull() && !from.Events.IsUnknown() && to.Events.IsNull() && len(from.Events.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Events, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Events = from.Events
	}
	if !from.HttpUrlSpec.IsNull() && !from.HttpUrlSpec.IsUnknown() {
		if toHttpUrlSpec, ok := to.GetHttpUrlSpec(ctx); ok {
			if fromHttpUrlSpec, ok := from.GetHttpUrlSpec(ctx); ok {
				toHttpUrlSpec.SyncFieldsDuringRead(ctx, fromHttpUrlSpec)
				to.SetHttpUrlSpec(ctx, toHttpUrlSpec)
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
}

func (m UpdateRegistryWebhook_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["events"] = attrs["events"].SetOptional()
	attrs["http_url_spec"] = attrs["http_url_spec"].SetOptional()
	attrs["http_url_spec"] = attrs["http_url_spec"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["id"] = attrs["id"].SetRequired()
	attrs["job_spec"] = attrs["job_spec"].SetOptional()
	attrs["job_spec"] = attrs["job_spec"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["status"] = attrs["status"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateRegistryWebhook.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateRegistryWebhook_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"events":        reflect.TypeOf(types.String{}),
		"http_url_spec": reflect.TypeOf(HttpUrlSpec_SdkV2{}),
		"job_spec":      reflect.TypeOf(JobSpec_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRegistryWebhook_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateRegistryWebhook_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":   m.Description,
			"events":        m.Events,
			"http_url_spec": m.HttpUrlSpec,
			"id":            m.Id,
			"job_spec":      m.JobSpec,
			"status":        m.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateRegistryWebhook_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"events": basetypes.ListType{
				ElemType: types.StringType,
			},
			"http_url_spec": basetypes.ListType{
				ElemType: HttpUrlSpec_SdkV2{}.Type(ctx),
			},
			"id": types.StringType,
			"job_spec": basetypes.ListType{
				ElemType: JobSpec_SdkV2{}.Type(ctx),
			},
			"status": types.StringType,
		},
	}
}

// GetEvents returns the value of the Events field in UpdateRegistryWebhook_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateRegistryWebhook_SdkV2) GetEvents(ctx context.Context) ([]types.String, bool) {
	if m.Events.IsNull() || m.Events.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Events.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEvents sets the value of the Events field in UpdateRegistryWebhook_SdkV2.
func (m *UpdateRegistryWebhook_SdkV2) SetEvents(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["events"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Events = types.ListValueMust(t, vs)
}

// GetHttpUrlSpec returns the value of the HttpUrlSpec field in UpdateRegistryWebhook_SdkV2 as
// a HttpUrlSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateRegistryWebhook_SdkV2) GetHttpUrlSpec(ctx context.Context) (HttpUrlSpec_SdkV2, bool) {
	var e HttpUrlSpec_SdkV2
	if m.HttpUrlSpec.IsNull() || m.HttpUrlSpec.IsUnknown() {
		return e, false
	}
	var v []HttpUrlSpec_SdkV2
	d := m.HttpUrlSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetHttpUrlSpec sets the value of the HttpUrlSpec field in UpdateRegistryWebhook_SdkV2.
func (m *UpdateRegistryWebhook_SdkV2) SetHttpUrlSpec(ctx context.Context, v HttpUrlSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["http_url_spec"]
	m.HttpUrlSpec = types.ListValueMust(t, vs)
}

// GetJobSpec returns the value of the JobSpec field in UpdateRegistryWebhook_SdkV2 as
// a JobSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateRegistryWebhook_SdkV2) GetJobSpec(ctx context.Context) (JobSpec_SdkV2, bool) {
	var e JobSpec_SdkV2
	if m.JobSpec.IsNull() || m.JobSpec.IsUnknown() {
		return e, false
	}
	var v []JobSpec_SdkV2
	d := m.JobSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetJobSpec sets the value of the JobSpec field in UpdateRegistryWebhook_SdkV2.
func (m *UpdateRegistryWebhook_SdkV2) SetJobSpec(ctx context.Context, v JobSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["job_spec"]
	m.JobSpec = types.ListValueMust(t, vs)
}

type UpdateRun_SdkV2 struct {
	// Unix timestamp in milliseconds of when the run ended.
	EndTime types.Int64 `tfsdk:"end_time"`
	// ID of the run to update. Must be provided.
	RunId types.String `tfsdk:"run_id"`
	// Updated name of the run.
	RunName types.String `tfsdk:"run_name"`
	// [Deprecated, use `run_id` instead] ID of the run to update. This field
	// will be removed in a future MLflow version.
	RunUuid types.String `tfsdk:"run_uuid"`
	// Updated status of the run.
	Status types.String `tfsdk:"status"`
}

func (to *UpdateRun_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateRun_SdkV2) {
}

func (to *UpdateRun_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateRun_SdkV2) {
}

func (m UpdateRun_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["end_time"] = attrs["end_time"].SetOptional()
	attrs["run_id"] = attrs["run_id"].SetOptional()
	attrs["run_name"] = attrs["run_name"].SetOptional()
	attrs["run_uuid"] = attrs["run_uuid"].SetOptional()
	attrs["status"] = attrs["status"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateRun.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateRun_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRun_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateRun_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"end_time": m.EndTime,
			"run_id":   m.RunId,
			"run_name": m.RunName,
			"run_uuid": m.RunUuid,
			"status":   m.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateRun_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"end_time": types.Int64Type,
			"run_id":   types.StringType,
			"run_name": types.StringType,
			"run_uuid": types.StringType,
			"status":   types.StringType,
		},
	}
}

type UpdateRunResponse_SdkV2 struct {
	// Updated metadata of the run.
	RunInfo types.List `tfsdk:"run_info"`
}

func (to *UpdateRunResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateRunResponse_SdkV2) {
	if !from.RunInfo.IsNull() && !from.RunInfo.IsUnknown() {
		if toRunInfo, ok := to.GetRunInfo(ctx); ok {
			if fromRunInfo, ok := from.GetRunInfo(ctx); ok {
				// Recursively sync the fields of RunInfo
				toRunInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromRunInfo)
				to.SetRunInfo(ctx, toRunInfo)
			}
		}
	}
}

func (to *UpdateRunResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateRunResponse_SdkV2) {
	if !from.RunInfo.IsNull() && !from.RunInfo.IsUnknown() {
		if toRunInfo, ok := to.GetRunInfo(ctx); ok {
			if fromRunInfo, ok := from.GetRunInfo(ctx); ok {
				toRunInfo.SyncFieldsDuringRead(ctx, fromRunInfo)
				to.SetRunInfo(ctx, toRunInfo)
			}
		}
	}
}

func (m UpdateRunResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["run_info"] = attrs["run_info"].SetOptional()
	attrs["run_info"] = attrs["run_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateRunResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateRunResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"run_info": reflect.TypeOf(RunInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRunResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateRunResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run_info": m.RunInfo,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateRunResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run_info": basetypes.ListType{
				ElemType: RunInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetRunInfo returns the value of the RunInfo field in UpdateRunResponse_SdkV2 as
// a RunInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateRunResponse_SdkV2) GetRunInfo(ctx context.Context) (RunInfo_SdkV2, bool) {
	var e RunInfo_SdkV2
	if m.RunInfo.IsNull() || m.RunInfo.IsUnknown() {
		return e, false
	}
	var v []RunInfo_SdkV2
	d := m.RunInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRunInfo sets the value of the RunInfo field in UpdateRunResponse_SdkV2.
func (m *UpdateRunResponse_SdkV2) SetRunInfo(ctx context.Context, v RunInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["run_info"]
	m.RunInfo = types.ListValueMust(t, vs)
}

type UpdateWebhookResponse_SdkV2 struct {
	Webhook types.List `tfsdk:"webhook"`
}

func (to *UpdateWebhookResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateWebhookResponse_SdkV2) {
	if !from.Webhook.IsNull() && !from.Webhook.IsUnknown() {
		if toWebhook, ok := to.GetWebhook(ctx); ok {
			if fromWebhook, ok := from.GetWebhook(ctx); ok {
				// Recursively sync the fields of Webhook
				toWebhook.SyncFieldsDuringCreateOrUpdate(ctx, fromWebhook)
				to.SetWebhook(ctx, toWebhook)
			}
		}
	}
}

func (to *UpdateWebhookResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateWebhookResponse_SdkV2) {
	if !from.Webhook.IsNull() && !from.Webhook.IsUnknown() {
		if toWebhook, ok := to.GetWebhook(ctx); ok {
			if fromWebhook, ok := from.GetWebhook(ctx); ok {
				toWebhook.SyncFieldsDuringRead(ctx, fromWebhook)
				to.SetWebhook(ctx, toWebhook)
			}
		}
	}
}

func (m UpdateWebhookResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["webhook"] = attrs["webhook"].SetOptional()
	attrs["webhook"] = attrs["webhook"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateWebhookResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateWebhookResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"webhook": reflect.TypeOf(RegistryWebhook_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWebhookResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateWebhookResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"webhook": m.Webhook,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateWebhookResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"webhook": basetypes.ListType{
				ElemType: RegistryWebhook_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetWebhook returns the value of the Webhook field in UpdateWebhookResponse_SdkV2 as
// a RegistryWebhook_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateWebhookResponse_SdkV2) GetWebhook(ctx context.Context) (RegistryWebhook_SdkV2, bool) {
	var e RegistryWebhook_SdkV2
	if m.Webhook.IsNull() || m.Webhook.IsUnknown() {
		return e, false
	}
	var v []RegistryWebhook_SdkV2
	d := m.Webhook.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWebhook sets the value of the Webhook field in UpdateWebhookResponse_SdkV2.
func (m *UpdateWebhookResponse_SdkV2) SetWebhook(ctx context.Context, v RegistryWebhook_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["webhook"]
	m.Webhook = types.ListValueMust(t, vs)
}
