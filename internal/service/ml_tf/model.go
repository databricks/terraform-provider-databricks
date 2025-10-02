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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// For activities, this contains the activity recorded for the action. For
// comments, this contains the comment details. For transition requests, this
// contains the transition request details.
type Activity struct {
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

func (to *Activity) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Activity) {
}

func (to *Activity) SyncFieldsDuringRead(ctx context.Context, from Activity) {
}

func (m Activity) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Activity) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Activity
// only implements ToObjectValue() and Type().
func (m Activity) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m Activity) Type(ctx context.Context) attr.Type {
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
type ApproveTransitionRequest struct {
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

func (to *ApproveTransitionRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ApproveTransitionRequest) {
}

func (to *ApproveTransitionRequest) SyncFieldsDuringRead(ctx context.Context, from ApproveTransitionRequest) {
}

func (m ApproveTransitionRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ApproveTransitionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ApproveTransitionRequest
// only implements ToObjectValue() and Type().
func (m ApproveTransitionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ApproveTransitionRequest) Type(ctx context.Context) attr.Type {
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

type ApproveTransitionRequestResponse struct {
	// New activity generated as a result of this operation.
	Activity types.Object `tfsdk:"activity"`
}

func (to *ApproveTransitionRequestResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ApproveTransitionRequestResponse) {
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

func (to *ApproveTransitionRequestResponse) SyncFieldsDuringRead(ctx context.Context, from ApproveTransitionRequestResponse) {
	if !from.Activity.IsNull() && !from.Activity.IsUnknown() {
		if toActivity, ok := to.GetActivity(ctx); ok {
			if fromActivity, ok := from.GetActivity(ctx); ok {
				toActivity.SyncFieldsDuringRead(ctx, fromActivity)
				to.SetActivity(ctx, toActivity)
			}
		}
	}
}

func (m ApproveTransitionRequestResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["activity"] = attrs["activity"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ApproveTransitionRequestResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ApproveTransitionRequestResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"activity": reflect.TypeOf(Activity{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ApproveTransitionRequestResponse
// only implements ToObjectValue() and Type().
func (m ApproveTransitionRequestResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"activity": m.Activity,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ApproveTransitionRequestResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"activity": Activity{}.Type(ctx),
		},
	}
}

// GetActivity returns the value of the Activity field in ApproveTransitionRequestResponse as
// a Activity value.
// If the field is unknown or null, the boolean return value is false.
func (m *ApproveTransitionRequestResponse) GetActivity(ctx context.Context) (Activity, bool) {
	var e Activity
	if m.Activity.IsNull() || m.Activity.IsUnknown() {
		return e, false
	}
	var v Activity
	d := m.Activity.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetActivity sets the value of the Activity field in ApproveTransitionRequestResponse.
func (m *ApproveTransitionRequestResponse) SetActivity(ctx context.Context, v Activity) {
	vs := v.ToObjectValue(ctx)
	m.Activity = vs
}

// For activities, this contains the activity recorded for the action. For
// comments, this contains the comment details. For transition requests, this
// contains the transition request details.
type CommentObject struct {
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

func (to *CommentObject) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CommentObject) {
	if !from.AvailableActions.IsNull() && !from.AvailableActions.IsUnknown() && to.AvailableActions.IsNull() && len(from.AvailableActions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AvailableActions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AvailableActions = from.AvailableActions
	}
}

func (to *CommentObject) SyncFieldsDuringRead(ctx context.Context, from CommentObject) {
	if !from.AvailableActions.IsNull() && !from.AvailableActions.IsUnknown() && to.AvailableActions.IsNull() && len(from.AvailableActions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AvailableActions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AvailableActions = from.AvailableActions
	}
}

func (m CommentObject) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CommentObject) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"available_actions": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CommentObject
// only implements ToObjectValue() and Type().
func (m CommentObject) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m CommentObject) Type(ctx context.Context) attr.Type {
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

// GetAvailableActions returns the value of the AvailableActions field in CommentObject as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CommentObject) GetAvailableActions(ctx context.Context) ([]types.String, bool) {
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

// SetAvailableActions sets the value of the AvailableActions field in CommentObject.
func (m *CommentObject) SetAvailableActions(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["available_actions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AvailableActions = types.ListValueMust(t, vs)
}

// Details required to create a comment on a model version.
type CreateComment struct {
	// User-provided comment on the action.
	Comment types.String `tfsdk:"comment"`
	// Name of the model.
	Name types.String `tfsdk:"name"`
	// Version of the model.
	Version types.String `tfsdk:"version"`
}

func (to *CreateComment) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateComment) {
}

func (to *CreateComment) SyncFieldsDuringRead(ctx context.Context, from CreateComment) {
}

func (m CreateComment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateComment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateComment
// only implements ToObjectValue() and Type().
func (m CreateComment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment": m.Comment,
			"name":    m.Name,
			"version": m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateComment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment": types.StringType,
			"name":    types.StringType,
			"version": types.StringType,
		},
	}
}

type CreateCommentResponse struct {
	// New comment object
	Comment types.Object `tfsdk:"comment"`
}

func (to *CreateCommentResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCommentResponse) {
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

func (to *CreateCommentResponse) SyncFieldsDuringRead(ctx context.Context, from CreateCommentResponse) {
	if !from.Comment.IsNull() && !from.Comment.IsUnknown() {
		if toComment, ok := to.GetComment(ctx); ok {
			if fromComment, ok := from.GetComment(ctx); ok {
				toComment.SyncFieldsDuringRead(ctx, fromComment)
				to.SetComment(ctx, toComment)
			}
		}
	}
}

func (m CreateCommentResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCommentResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateCommentResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"comment": reflect.TypeOf(CommentObject{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCommentResponse
// only implements ToObjectValue() and Type().
func (m CreateCommentResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment": m.Comment,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateCommentResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment": CommentObject{}.Type(ctx),
		},
	}
}

// GetComment returns the value of the Comment field in CreateCommentResponse as
// a CommentObject value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateCommentResponse) GetComment(ctx context.Context) (CommentObject, bool) {
	var e CommentObject
	if m.Comment.IsNull() || m.Comment.IsUnknown() {
		return e, false
	}
	var v CommentObject
	d := m.Comment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetComment sets the value of the Comment field in CreateCommentResponse.
func (m *CreateCommentResponse) SetComment(ctx context.Context, v CommentObject) {
	vs := v.ToObjectValue(ctx)
	m.Comment = vs
}

type CreateExperiment struct {
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

func (to *CreateExperiment) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateExperiment) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *CreateExperiment) SyncFieldsDuringRead(ctx context.Context, from CreateExperiment) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m CreateExperiment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateExperiment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(ExperimentTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateExperiment
// only implements ToObjectValue() and Type().
func (m CreateExperiment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"artifact_location": m.ArtifactLocation,
			"name":              m.Name,
			"tags":              m.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateExperiment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"artifact_location": types.StringType,
			"name":              types.StringType,
			"tags": basetypes.ListType{
				ElemType: ExperimentTag{}.Type(ctx),
			},
		},
	}
}

// GetTags returns the value of the Tags field in CreateExperiment as
// a slice of ExperimentTag values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateExperiment) GetTags(ctx context.Context) ([]ExperimentTag, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []ExperimentTag
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in CreateExperiment.
func (m *CreateExperiment) SetTags(ctx context.Context, v []ExperimentTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type CreateExperimentResponse struct {
	// Unique identifier for the experiment.
	ExperimentId types.String `tfsdk:"experiment_id"`
}

func (to *CreateExperimentResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateExperimentResponse) {
}

func (to *CreateExperimentResponse) SyncFieldsDuringRead(ctx context.Context, from CreateExperimentResponse) {
}

func (m CreateExperimentResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateExperimentResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateExperimentResponse
// only implements ToObjectValue() and Type().
func (m CreateExperimentResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": m.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateExperimentResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
		},
	}
}

type CreateFeatureRequest struct {
	// Feature to create.
	Feature types.Object `tfsdk:"feature"`
}

func (to *CreateFeatureRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateFeatureRequest) {
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

func (to *CreateFeatureRequest) SyncFieldsDuringRead(ctx context.Context, from CreateFeatureRequest) {
	if !from.Feature.IsNull() && !from.Feature.IsUnknown() {
		if toFeature, ok := to.GetFeature(ctx); ok {
			if fromFeature, ok := from.GetFeature(ctx); ok {
				toFeature.SyncFieldsDuringRead(ctx, fromFeature)
				to.SetFeature(ctx, toFeature)
			}
		}
	}
}

func (m CreateFeatureRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["feature"] = attrs["feature"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateFeatureRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateFeatureRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"feature": reflect.TypeOf(Feature{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateFeatureRequest
// only implements ToObjectValue() and Type().
func (m CreateFeatureRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature": m.Feature,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateFeatureRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"feature": Feature{}.Type(ctx),
		},
	}
}

// GetFeature returns the value of the Feature field in CreateFeatureRequest as
// a Feature value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateFeatureRequest) GetFeature(ctx context.Context) (Feature, bool) {
	var e Feature
	if m.Feature.IsNull() || m.Feature.IsUnknown() {
		return e, false
	}
	var v Feature
	d := m.Feature.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFeature sets the value of the Feature field in CreateFeatureRequest.
func (m *CreateFeatureRequest) SetFeature(ctx context.Context, v Feature) {
	vs := v.ToObjectValue(ctx)
	m.Feature = vs
}

type CreateFeatureTagRequest struct {
	FeatureName types.String `tfsdk:"-"`

	FeatureTag types.Object `tfsdk:"feature_tag"`

	TableName types.String `tfsdk:"-"`
}

func (to *CreateFeatureTagRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateFeatureTagRequest) {
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

func (to *CreateFeatureTagRequest) SyncFieldsDuringRead(ctx context.Context, from CreateFeatureTagRequest) {
	if !from.FeatureTag.IsNull() && !from.FeatureTag.IsUnknown() {
		if toFeatureTag, ok := to.GetFeatureTag(ctx); ok {
			if fromFeatureTag, ok := from.GetFeatureTag(ctx); ok {
				toFeatureTag.SyncFieldsDuringRead(ctx, fromFeatureTag)
				to.SetFeatureTag(ctx, toFeatureTag)
			}
		}
	}
}

func (m CreateFeatureTagRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["feature_tag"] = attrs["feature_tag"].SetRequired()
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
func (m CreateFeatureTagRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"feature_tag": reflect.TypeOf(FeatureTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateFeatureTagRequest
// only implements ToObjectValue() and Type().
func (m CreateFeatureTagRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_name": m.FeatureName,
			"feature_tag":  m.FeatureTag,
			"table_name":   m.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateFeatureTagRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"feature_name": types.StringType,
			"feature_tag":  FeatureTag{}.Type(ctx),
			"table_name":   types.StringType,
		},
	}
}

// GetFeatureTag returns the value of the FeatureTag field in CreateFeatureTagRequest as
// a FeatureTag value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateFeatureTagRequest) GetFeatureTag(ctx context.Context) (FeatureTag, bool) {
	var e FeatureTag
	if m.FeatureTag.IsNull() || m.FeatureTag.IsUnknown() {
		return e, false
	}
	var v FeatureTag
	d := m.FeatureTag.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFeatureTag sets the value of the FeatureTag field in CreateFeatureTagRequest.
func (m *CreateFeatureTagRequest) SetFeatureTag(ctx context.Context, v FeatureTag) {
	vs := v.ToObjectValue(ctx)
	m.FeatureTag = vs
}

type CreateForecastingExperimentRequest struct {
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

func (to *CreateForecastingExperimentRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateForecastingExperimentRequest) {
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

func (to *CreateForecastingExperimentRequest) SyncFieldsDuringRead(ctx context.Context, from CreateForecastingExperimentRequest) {
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

func (m CreateForecastingExperimentRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateForecastingExperimentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"holiday_regions":               reflect.TypeOf(types.String{}),
		"include_features":              reflect.TypeOf(types.String{}),
		"timeseries_identifier_columns": reflect.TypeOf(types.String{}),
		"training_frameworks":           reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateForecastingExperimentRequest
// only implements ToObjectValue() and Type().
func (m CreateForecastingExperimentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m CreateForecastingExperimentRequest) Type(ctx context.Context) attr.Type {
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

// GetHolidayRegions returns the value of the HolidayRegions field in CreateForecastingExperimentRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateForecastingExperimentRequest) GetHolidayRegions(ctx context.Context) ([]types.String, bool) {
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

// SetHolidayRegions sets the value of the HolidayRegions field in CreateForecastingExperimentRequest.
func (m *CreateForecastingExperimentRequest) SetHolidayRegions(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["holiday_regions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.HolidayRegions = types.ListValueMust(t, vs)
}

// GetIncludeFeatures returns the value of the IncludeFeatures field in CreateForecastingExperimentRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateForecastingExperimentRequest) GetIncludeFeatures(ctx context.Context) ([]types.String, bool) {
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

// SetIncludeFeatures sets the value of the IncludeFeatures field in CreateForecastingExperimentRequest.
func (m *CreateForecastingExperimentRequest) SetIncludeFeatures(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["include_features"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.IncludeFeatures = types.ListValueMust(t, vs)
}

// GetTimeseriesIdentifierColumns returns the value of the TimeseriesIdentifierColumns field in CreateForecastingExperimentRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateForecastingExperimentRequest) GetTimeseriesIdentifierColumns(ctx context.Context) ([]types.String, bool) {
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

// SetTimeseriesIdentifierColumns sets the value of the TimeseriesIdentifierColumns field in CreateForecastingExperimentRequest.
func (m *CreateForecastingExperimentRequest) SetTimeseriesIdentifierColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["timeseries_identifier_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.TimeseriesIdentifierColumns = types.ListValueMust(t, vs)
}

// GetTrainingFrameworks returns the value of the TrainingFrameworks field in CreateForecastingExperimentRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateForecastingExperimentRequest) GetTrainingFrameworks(ctx context.Context) ([]types.String, bool) {
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

// SetTrainingFrameworks sets the value of the TrainingFrameworks field in CreateForecastingExperimentRequest.
func (m *CreateForecastingExperimentRequest) SetTrainingFrameworks(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["training_frameworks"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.TrainingFrameworks = types.ListValueMust(t, vs)
}

type CreateForecastingExperimentResponse struct {
	// The unique ID of the created forecasting experiment
	ExperimentId types.String `tfsdk:"experiment_id"`
}

func (to *CreateForecastingExperimentResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateForecastingExperimentResponse) {
}

func (to *CreateForecastingExperimentResponse) SyncFieldsDuringRead(ctx context.Context, from CreateForecastingExperimentResponse) {
}

func (m CreateForecastingExperimentResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateForecastingExperimentResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateForecastingExperimentResponse
// only implements ToObjectValue() and Type().
func (m CreateForecastingExperimentResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": m.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateForecastingExperimentResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
		},
	}
}

type CreateLoggedModelRequest struct {
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

func (to *CreateLoggedModelRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateLoggedModelRequest) {
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

func (to *CreateLoggedModelRequest) SyncFieldsDuringRead(ctx context.Context, from CreateLoggedModelRequest) {
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

func (m CreateLoggedModelRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateLoggedModelRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"params": reflect.TypeOf(LoggedModelParameter{}),
		"tags":   reflect.TypeOf(LoggedModelTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateLoggedModelRequest
// only implements ToObjectValue() and Type().
func (m CreateLoggedModelRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m CreateLoggedModelRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
			"model_type":    types.StringType,
			"name":          types.StringType,
			"params": basetypes.ListType{
				ElemType: LoggedModelParameter{}.Type(ctx),
			},
			"source_run_id": types.StringType,
			"tags": basetypes.ListType{
				ElemType: LoggedModelTag{}.Type(ctx),
			},
		},
	}
}

// GetParams returns the value of the Params field in CreateLoggedModelRequest as
// a slice of LoggedModelParameter values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateLoggedModelRequest) GetParams(ctx context.Context) ([]LoggedModelParameter, bool) {
	if m.Params.IsNull() || m.Params.IsUnknown() {
		return nil, false
	}
	var v []LoggedModelParameter
	d := m.Params.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParams sets the value of the Params field in CreateLoggedModelRequest.
func (m *CreateLoggedModelRequest) SetParams(ctx context.Context, v []LoggedModelParameter) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["params"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Params = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in CreateLoggedModelRequest as
// a slice of LoggedModelTag values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateLoggedModelRequest) GetTags(ctx context.Context) ([]LoggedModelTag, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []LoggedModelTag
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in CreateLoggedModelRequest.
func (m *CreateLoggedModelRequest) SetTags(ctx context.Context, v []LoggedModelTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type CreateLoggedModelResponse struct {
	// The newly created logged model.
	Model types.Object `tfsdk:"model"`
}

func (to *CreateLoggedModelResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateLoggedModelResponse) {
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

func (to *CreateLoggedModelResponse) SyncFieldsDuringRead(ctx context.Context, from CreateLoggedModelResponse) {
	if !from.Model.IsNull() && !from.Model.IsUnknown() {
		if toModel, ok := to.GetModel(ctx); ok {
			if fromModel, ok := from.GetModel(ctx); ok {
				toModel.SyncFieldsDuringRead(ctx, fromModel)
				to.SetModel(ctx, toModel)
			}
		}
	}
}

func (m CreateLoggedModelResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["model"] = attrs["model"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateLoggedModelResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateLoggedModelResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model": reflect.TypeOf(LoggedModel{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateLoggedModelResponse
// only implements ToObjectValue() and Type().
func (m CreateLoggedModelResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model": m.Model,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateLoggedModelResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model": LoggedModel{}.Type(ctx),
		},
	}
}

// GetModel returns the value of the Model field in CreateLoggedModelResponse as
// a LoggedModel value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateLoggedModelResponse) GetModel(ctx context.Context) (LoggedModel, bool) {
	var e LoggedModel
	if m.Model.IsNull() || m.Model.IsUnknown() {
		return e, false
	}
	var v LoggedModel
	d := m.Model.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModel sets the value of the Model field in CreateLoggedModelResponse.
func (m *CreateLoggedModelResponse) SetModel(ctx context.Context, v LoggedModel) {
	vs := v.ToObjectValue(ctx)
	m.Model = vs
}

type CreateModelRequest struct {
	// Optional description for registered model.
	Description types.String `tfsdk:"description"`
	// Register models under this name
	Name types.String `tfsdk:"name"`
	// Additional metadata for registered model.
	Tags types.List `tfsdk:"tags"`
}

func (to *CreateModelRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateModelRequest) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *CreateModelRequest) SyncFieldsDuringRead(ctx context.Context, from CreateModelRequest) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m CreateModelRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateModelRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(ModelTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateModelRequest
// only implements ToObjectValue() and Type().
func (m CreateModelRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description": m.Description,
			"name":        m.Name,
			"tags":        m.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateModelRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"name":        types.StringType,
			"tags": basetypes.ListType{
				ElemType: ModelTag{}.Type(ctx),
			},
		},
	}
}

// GetTags returns the value of the Tags field in CreateModelRequest as
// a slice of ModelTag values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateModelRequest) GetTags(ctx context.Context) ([]ModelTag, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []ModelTag
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in CreateModelRequest.
func (m *CreateModelRequest) SetTags(ctx context.Context, v []ModelTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type CreateModelResponse struct {
	RegisteredModel types.Object `tfsdk:"registered_model"`
}

func (to *CreateModelResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateModelResponse) {
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

func (to *CreateModelResponse) SyncFieldsDuringRead(ctx context.Context, from CreateModelResponse) {
	if !from.RegisteredModel.IsNull() && !from.RegisteredModel.IsUnknown() {
		if toRegisteredModel, ok := to.GetRegisteredModel(ctx); ok {
			if fromRegisteredModel, ok := from.GetRegisteredModel(ctx); ok {
				toRegisteredModel.SyncFieldsDuringRead(ctx, fromRegisteredModel)
				to.SetRegisteredModel(ctx, toRegisteredModel)
			}
		}
	}
}

func (m CreateModelResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["registered_model"] = attrs["registered_model"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateModelResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateModelResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"registered_model": reflect.TypeOf(Model{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateModelResponse
// only implements ToObjectValue() and Type().
func (m CreateModelResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"registered_model": m.RegisteredModel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateModelResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"registered_model": Model{}.Type(ctx),
		},
	}
}

// GetRegisteredModel returns the value of the RegisteredModel field in CreateModelResponse as
// a Model value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateModelResponse) GetRegisteredModel(ctx context.Context) (Model, bool) {
	var e Model
	if m.RegisteredModel.IsNull() || m.RegisteredModel.IsUnknown() {
		return e, false
	}
	var v Model
	d := m.RegisteredModel.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRegisteredModel sets the value of the RegisteredModel field in CreateModelResponse.
func (m *CreateModelResponse) SetRegisteredModel(ctx context.Context, v Model) {
	vs := v.ToObjectValue(ctx)
	m.RegisteredModel = vs
}

type CreateModelVersionRequest struct {
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

func (to *CreateModelVersionRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateModelVersionRequest) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *CreateModelVersionRequest) SyncFieldsDuringRead(ctx context.Context, from CreateModelVersionRequest) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m CreateModelVersionRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateModelVersionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(ModelVersionTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateModelVersionRequest
// only implements ToObjectValue() and Type().
func (m CreateModelVersionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m CreateModelVersionRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"name":        types.StringType,
			"run_id":      types.StringType,
			"run_link":    types.StringType,
			"source":      types.StringType,
			"tags": basetypes.ListType{
				ElemType: ModelVersionTag{}.Type(ctx),
			},
		},
	}
}

// GetTags returns the value of the Tags field in CreateModelVersionRequest as
// a slice of ModelVersionTag values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateModelVersionRequest) GetTags(ctx context.Context) ([]ModelVersionTag, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []ModelVersionTag
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in CreateModelVersionRequest.
func (m *CreateModelVersionRequest) SetTags(ctx context.Context, v []ModelVersionTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type CreateModelVersionResponse struct {
	// Return new version number generated for this model in registry.
	ModelVersion types.Object `tfsdk:"model_version"`
}

func (to *CreateModelVersionResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateModelVersionResponse) {
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

func (to *CreateModelVersionResponse) SyncFieldsDuringRead(ctx context.Context, from CreateModelVersionResponse) {
	if !from.ModelVersion.IsNull() && !from.ModelVersion.IsUnknown() {
		if toModelVersion, ok := to.GetModelVersion(ctx); ok {
			if fromModelVersion, ok := from.GetModelVersion(ctx); ok {
				toModelVersion.SyncFieldsDuringRead(ctx, fromModelVersion)
				to.SetModelVersion(ctx, toModelVersion)
			}
		}
	}
}

func (m CreateModelVersionResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["model_version"] = attrs["model_version"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateModelVersionResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateModelVersionResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model_version": reflect.TypeOf(ModelVersion{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateModelVersionResponse
// only implements ToObjectValue() and Type().
func (m CreateModelVersionResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_version": m.ModelVersion,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateModelVersionResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_version": ModelVersion{}.Type(ctx),
		},
	}
}

// GetModelVersion returns the value of the ModelVersion field in CreateModelVersionResponse as
// a ModelVersion value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateModelVersionResponse) GetModelVersion(ctx context.Context) (ModelVersion, bool) {
	var e ModelVersion
	if m.ModelVersion.IsNull() || m.ModelVersion.IsUnknown() {
		return e, false
	}
	var v ModelVersion
	d := m.ModelVersion.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModelVersion sets the value of the ModelVersion field in CreateModelVersionResponse.
func (m *CreateModelVersionResponse) SetModelVersion(ctx context.Context, v ModelVersion) {
	vs := v.ToObjectValue(ctx)
	m.ModelVersion = vs
}

type CreateOnlineStoreRequest struct {
	// Online store to create.
	OnlineStore types.Object `tfsdk:"online_store"`
}

func (to *CreateOnlineStoreRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateOnlineStoreRequest) {
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

func (to *CreateOnlineStoreRequest) SyncFieldsDuringRead(ctx context.Context, from CreateOnlineStoreRequest) {
	if !from.OnlineStore.IsNull() && !from.OnlineStore.IsUnknown() {
		if toOnlineStore, ok := to.GetOnlineStore(ctx); ok {
			if fromOnlineStore, ok := from.GetOnlineStore(ctx); ok {
				toOnlineStore.SyncFieldsDuringRead(ctx, fromOnlineStore)
				to.SetOnlineStore(ctx, toOnlineStore)
			}
		}
	}
}

func (m CreateOnlineStoreRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["online_store"] = attrs["online_store"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateOnlineStoreRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateOnlineStoreRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"online_store": reflect.TypeOf(OnlineStore{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateOnlineStoreRequest
// only implements ToObjectValue() and Type().
func (m CreateOnlineStoreRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"online_store": m.OnlineStore,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateOnlineStoreRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"online_store": OnlineStore{}.Type(ctx),
		},
	}
}

// GetOnlineStore returns the value of the OnlineStore field in CreateOnlineStoreRequest as
// a OnlineStore value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateOnlineStoreRequest) GetOnlineStore(ctx context.Context) (OnlineStore, bool) {
	var e OnlineStore
	if m.OnlineStore.IsNull() || m.OnlineStore.IsUnknown() {
		return e, false
	}
	var v OnlineStore
	d := m.OnlineStore.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOnlineStore sets the value of the OnlineStore field in CreateOnlineStoreRequest.
func (m *CreateOnlineStoreRequest) SetOnlineStore(ctx context.Context, v OnlineStore) {
	vs := v.ToObjectValue(ctx)
	m.OnlineStore = vs
}

// Details required to create a registry webhook.
type CreateRegistryWebhook struct {
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
	HttpUrlSpec types.Object `tfsdk:"http_url_spec"`
	// ID of the job that the webhook runs.
	JobSpec types.Object `tfsdk:"job_spec"`
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

func (to *CreateRegistryWebhook) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateRegistryWebhook) {
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

func (to *CreateRegistryWebhook) SyncFieldsDuringRead(ctx context.Context, from CreateRegistryWebhook) {
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

func (m CreateRegistryWebhook) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["events"] = attrs["events"].SetRequired()
	attrs["http_url_spec"] = attrs["http_url_spec"].SetOptional()
	attrs["job_spec"] = attrs["job_spec"].SetOptional()
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
func (m CreateRegistryWebhook) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"events":        reflect.TypeOf(types.String{}),
		"http_url_spec": reflect.TypeOf(HttpUrlSpec{}),
		"job_spec":      reflect.TypeOf(JobSpec{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateRegistryWebhook
// only implements ToObjectValue() and Type().
func (m CreateRegistryWebhook) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m CreateRegistryWebhook) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"events": basetypes.ListType{
				ElemType: types.StringType,
			},
			"http_url_spec": HttpUrlSpec{}.Type(ctx),
			"job_spec":      JobSpec{}.Type(ctx),
			"model_name":    types.StringType,
			"status":        types.StringType,
		},
	}
}

// GetEvents returns the value of the Events field in CreateRegistryWebhook as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateRegistryWebhook) GetEvents(ctx context.Context) ([]types.String, bool) {
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

// SetEvents sets the value of the Events field in CreateRegistryWebhook.
func (m *CreateRegistryWebhook) SetEvents(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["events"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Events = types.ListValueMust(t, vs)
}

// GetHttpUrlSpec returns the value of the HttpUrlSpec field in CreateRegistryWebhook as
// a HttpUrlSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateRegistryWebhook) GetHttpUrlSpec(ctx context.Context) (HttpUrlSpec, bool) {
	var e HttpUrlSpec
	if m.HttpUrlSpec.IsNull() || m.HttpUrlSpec.IsUnknown() {
		return e, false
	}
	var v HttpUrlSpec
	d := m.HttpUrlSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetHttpUrlSpec sets the value of the HttpUrlSpec field in CreateRegistryWebhook.
func (m *CreateRegistryWebhook) SetHttpUrlSpec(ctx context.Context, v HttpUrlSpec) {
	vs := v.ToObjectValue(ctx)
	m.HttpUrlSpec = vs
}

// GetJobSpec returns the value of the JobSpec field in CreateRegistryWebhook as
// a JobSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateRegistryWebhook) GetJobSpec(ctx context.Context) (JobSpec, bool) {
	var e JobSpec
	if m.JobSpec.IsNull() || m.JobSpec.IsUnknown() {
		return e, false
	}
	var v JobSpec
	d := m.JobSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetJobSpec sets the value of the JobSpec field in CreateRegistryWebhook.
func (m *CreateRegistryWebhook) SetJobSpec(ctx context.Context, v JobSpec) {
	vs := v.ToObjectValue(ctx)
	m.JobSpec = vs
}

type CreateRun struct {
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

func (to *CreateRun) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateRun) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *CreateRun) SyncFieldsDuringRead(ctx context.Context, from CreateRun) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m CreateRun) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateRun) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(RunTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateRun
// only implements ToObjectValue() and Type().
func (m CreateRun) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m CreateRun) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
			"run_name":      types.StringType,
			"start_time":    types.Int64Type,
			"tags": basetypes.ListType{
				ElemType: RunTag{}.Type(ctx),
			},
			"user_id": types.StringType,
		},
	}
}

// GetTags returns the value of the Tags field in CreateRun as
// a slice of RunTag values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateRun) GetTags(ctx context.Context) ([]RunTag, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []RunTag
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in CreateRun.
func (m *CreateRun) SetTags(ctx context.Context, v []RunTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type CreateRunResponse struct {
	// The newly created run.
	Run types.Object `tfsdk:"run"`
}

func (to *CreateRunResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateRunResponse) {
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

func (to *CreateRunResponse) SyncFieldsDuringRead(ctx context.Context, from CreateRunResponse) {
	if !from.Run.IsNull() && !from.Run.IsUnknown() {
		if toRun, ok := to.GetRun(ctx); ok {
			if fromRun, ok := from.GetRun(ctx); ok {
				toRun.SyncFieldsDuringRead(ctx, fromRun)
				to.SetRun(ctx, toRun)
			}
		}
	}
}

func (m CreateRunResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["run"] = attrs["run"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateRunResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateRunResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"run": reflect.TypeOf(Run{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateRunResponse
// only implements ToObjectValue() and Type().
func (m CreateRunResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run": m.Run,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateRunResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run": Run{}.Type(ctx),
		},
	}
}

// GetRun returns the value of the Run field in CreateRunResponse as
// a Run value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateRunResponse) GetRun(ctx context.Context) (Run, bool) {
	var e Run
	if m.Run.IsNull() || m.Run.IsUnknown() {
		return e, false
	}
	var v Run
	d := m.Run.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRun sets the value of the Run field in CreateRunResponse.
func (m *CreateRunResponse) SetRun(ctx context.Context, v Run) {
	vs := v.ToObjectValue(ctx)
	m.Run = vs
}

// Details required to create a model version stage transition request.
type CreateTransitionRequest struct {
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

func (to *CreateTransitionRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateTransitionRequest) {
}

func (to *CreateTransitionRequest) SyncFieldsDuringRead(ctx context.Context, from CreateTransitionRequest) {
}

func (m CreateTransitionRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateTransitionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateTransitionRequest
// only implements ToObjectValue() and Type().
func (m CreateTransitionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m CreateTransitionRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment": types.StringType,
			"name":    types.StringType,
			"stage":   types.StringType,
			"version": types.StringType,
		},
	}
}

type CreateTransitionRequestResponse struct {
	// New activity generated for stage transition request.
	Request types.Object `tfsdk:"request"`
}

func (to *CreateTransitionRequestResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateTransitionRequestResponse) {
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

func (to *CreateTransitionRequestResponse) SyncFieldsDuringRead(ctx context.Context, from CreateTransitionRequestResponse) {
	if !from.Request.IsNull() && !from.Request.IsUnknown() {
		if toRequest, ok := to.GetRequest(ctx); ok {
			if fromRequest, ok := from.GetRequest(ctx); ok {
				toRequest.SyncFieldsDuringRead(ctx, fromRequest)
				to.SetRequest(ctx, toRequest)
			}
		}
	}
}

func (m CreateTransitionRequestResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["request"] = attrs["request"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateTransitionRequestResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateTransitionRequestResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"request": reflect.TypeOf(TransitionRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateTransitionRequestResponse
// only implements ToObjectValue() and Type().
func (m CreateTransitionRequestResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"request": m.Request,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateTransitionRequestResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"request": TransitionRequest{}.Type(ctx),
		},
	}
}

// GetRequest returns the value of the Request field in CreateTransitionRequestResponse as
// a TransitionRequest value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateTransitionRequestResponse) GetRequest(ctx context.Context) (TransitionRequest, bool) {
	var e TransitionRequest
	if m.Request.IsNull() || m.Request.IsUnknown() {
		return e, false
	}
	var v TransitionRequest
	d := m.Request.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRequest sets the value of the Request field in CreateTransitionRequestResponse.
func (m *CreateTransitionRequestResponse) SetRequest(ctx context.Context, v TransitionRequest) {
	vs := v.ToObjectValue(ctx)
	m.Request = vs
}

type CreateWebhookResponse struct {
	Webhook types.Object `tfsdk:"webhook"`
}

func (to *CreateWebhookResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateWebhookResponse) {
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

func (to *CreateWebhookResponse) SyncFieldsDuringRead(ctx context.Context, from CreateWebhookResponse) {
	if !from.Webhook.IsNull() && !from.Webhook.IsUnknown() {
		if toWebhook, ok := to.GetWebhook(ctx); ok {
			if fromWebhook, ok := from.GetWebhook(ctx); ok {
				toWebhook.SyncFieldsDuringRead(ctx, fromWebhook)
				to.SetWebhook(ctx, toWebhook)
			}
		}
	}
}

func (m CreateWebhookResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["webhook"] = attrs["webhook"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateWebhookResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateWebhookResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"webhook": reflect.TypeOf(RegistryWebhook{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateWebhookResponse
// only implements ToObjectValue() and Type().
func (m CreateWebhookResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"webhook": m.Webhook,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateWebhookResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"webhook": RegistryWebhook{}.Type(ctx),
		},
	}
}

// GetWebhook returns the value of the Webhook field in CreateWebhookResponse as
// a RegistryWebhook value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateWebhookResponse) GetWebhook(ctx context.Context) (RegistryWebhook, bool) {
	var e RegistryWebhook
	if m.Webhook.IsNull() || m.Webhook.IsUnknown() {
		return e, false
	}
	var v RegistryWebhook
	d := m.Webhook.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWebhook sets the value of the Webhook field in CreateWebhookResponse.
func (m *CreateWebhookResponse) SetWebhook(ctx context.Context, v RegistryWebhook) {
	vs := v.ToObjectValue(ctx)
	m.Webhook = vs
}

type DataSource struct {
	DeltaTableSource types.Object `tfsdk:"delta_table_source"`
}

func (to *DataSource) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DataSource) {
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

func (to *DataSource) SyncFieldsDuringRead(ctx context.Context, from DataSource) {
	if !from.DeltaTableSource.IsNull() && !from.DeltaTableSource.IsUnknown() {
		if toDeltaTableSource, ok := to.GetDeltaTableSource(ctx); ok {
			if fromDeltaTableSource, ok := from.GetDeltaTableSource(ctx); ok {
				toDeltaTableSource.SyncFieldsDuringRead(ctx, fromDeltaTableSource)
				to.SetDeltaTableSource(ctx, toDeltaTableSource)
			}
		}
	}
}

func (m DataSource) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["delta_table_source"] = attrs["delta_table_source"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DataSource.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DataSource) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"delta_table_source": reflect.TypeOf(DeltaTableSource{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DataSource
// only implements ToObjectValue() and Type().
func (m DataSource) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"delta_table_source": m.DeltaTableSource,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DataSource) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"delta_table_source": DeltaTableSource{}.Type(ctx),
		},
	}
}

// GetDeltaTableSource returns the value of the DeltaTableSource field in DataSource as
// a DeltaTableSource value.
// If the field is unknown or null, the boolean return value is false.
func (m *DataSource) GetDeltaTableSource(ctx context.Context) (DeltaTableSource, bool) {
	var e DeltaTableSource
	if m.DeltaTableSource.IsNull() || m.DeltaTableSource.IsUnknown() {
		return e, false
	}
	var v DeltaTableSource
	d := m.DeltaTableSource.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDeltaTableSource sets the value of the DeltaTableSource field in DataSource.
func (m *DataSource) SetDeltaTableSource(ctx context.Context, v DeltaTableSource) {
	vs := v.ToObjectValue(ctx)
	m.DeltaTableSource = vs
}

// Dataset. Represents a reference to data used for training, testing, or
// evaluation during the model development process.
type Dataset struct {
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

func (to *Dataset) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Dataset) {
}

func (to *Dataset) SyncFieldsDuringRead(ctx context.Context, from Dataset) {
}

func (m Dataset) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Dataset) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Dataset
// only implements ToObjectValue() and Type().
func (m Dataset) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m Dataset) Type(ctx context.Context) attr.Type {
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
type DatasetInput struct {
	// The dataset being used as a Run input.
	Dataset types.Object `tfsdk:"dataset"`
	// A list of tags for the dataset input, e.g. a “context” tag with value
	// “training”
	Tags types.List `tfsdk:"tags"`
}

func (to *DatasetInput) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatasetInput) {
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

func (to *DatasetInput) SyncFieldsDuringRead(ctx context.Context, from DatasetInput) {
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

func (m DatasetInput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dataset"] = attrs["dataset"].SetRequired()
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
func (m DatasetInput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dataset": reflect.TypeOf(Dataset{}),
		"tags":    reflect.TypeOf(InputTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatasetInput
// only implements ToObjectValue() and Type().
func (m DatasetInput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dataset": m.Dataset,
			"tags":    m.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DatasetInput) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dataset": Dataset{}.Type(ctx),
			"tags": basetypes.ListType{
				ElemType: InputTag{}.Type(ctx),
			},
		},
	}
}

// GetDataset returns the value of the Dataset field in DatasetInput as
// a Dataset value.
// If the field is unknown or null, the boolean return value is false.
func (m *DatasetInput) GetDataset(ctx context.Context) (Dataset, bool) {
	var e Dataset
	if m.Dataset.IsNull() || m.Dataset.IsUnknown() {
		return e, false
	}
	var v Dataset
	d := m.Dataset.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDataset sets the value of the Dataset field in DatasetInput.
func (m *DatasetInput) SetDataset(ctx context.Context, v Dataset) {
	vs := v.ToObjectValue(ctx)
	m.Dataset = vs
}

// GetTags returns the value of the Tags field in DatasetInput as
// a slice of InputTag values.
// If the field is unknown or null, the boolean return value is false.
func (m *DatasetInput) GetTags(ctx context.Context) ([]InputTag, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []InputTag
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in DatasetInput.
func (m *DatasetInput) SetTags(ctx context.Context, v []InputTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type DeleteCommentRequest struct {
	// Unique identifier of an activity
	Id types.String `tfsdk:"-"`
}

func (to *DeleteCommentRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteCommentRequest) {
}

func (to *DeleteCommentRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteCommentRequest) {
}

func (m DeleteCommentRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteCommentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCommentRequest
// only implements ToObjectValue() and Type().
func (m DeleteCommentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteCommentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteCommentResponse struct {
}

func (to *DeleteCommentResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteCommentResponse) {
}

func (to *DeleteCommentResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteCommentResponse) {
}

func (m DeleteCommentResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCommentResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteCommentResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCommentResponse
// only implements ToObjectValue() and Type().
func (m DeleteCommentResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteCommentResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteExperiment struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"experiment_id"`
}

func (to *DeleteExperiment) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteExperiment) {
}

func (to *DeleteExperiment) SyncFieldsDuringRead(ctx context.Context, from DeleteExperiment) {
}

func (m DeleteExperiment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteExperiment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteExperiment
// only implements ToObjectValue() and Type().
func (m DeleteExperiment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": m.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteExperiment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
		},
	}
}

type DeleteExperimentResponse struct {
}

func (to *DeleteExperimentResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteExperimentResponse) {
}

func (to *DeleteExperimentResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteExperimentResponse) {
}

func (m DeleteExperimentResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteExperimentResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteExperimentResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteExperimentResponse
// only implements ToObjectValue() and Type().
func (m DeleteExperimentResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteExperimentResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteFeatureRequest struct {
	// Name of the feature to delete.
	FullName types.String `tfsdk:"-"`
}

func (to *DeleteFeatureRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteFeatureRequest) {
}

func (to *DeleteFeatureRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteFeatureRequest) {
}

func (m DeleteFeatureRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteFeatureRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteFeatureRequest
// only implements ToObjectValue() and Type().
func (m DeleteFeatureRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"full_name": m.FullName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteFeatureRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_name": types.StringType,
		},
	}
}

type DeleteFeatureTagRequest struct {
	// The name of the feature within the feature table.
	FeatureName types.String `tfsdk:"-"`
	// The key of the tag to delete.
	Key types.String `tfsdk:"-"`
	// The name of the feature table.
	TableName types.String `tfsdk:"-"`
}

func (to *DeleteFeatureTagRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteFeatureTagRequest) {
}

func (to *DeleteFeatureTagRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteFeatureTagRequest) {
}

func (m DeleteFeatureTagRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteFeatureTagRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteFeatureTagRequest
// only implements ToObjectValue() and Type().
func (m DeleteFeatureTagRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_name": m.FeatureName,
			"key":          m.Key,
			"table_name":   m.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteFeatureTagRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"feature_name": types.StringType,
			"key":          types.StringType,
			"table_name":   types.StringType,
		},
	}
}

type DeleteLoggedModelRequest struct {
	// The ID of the logged model to delete.
	ModelId types.String `tfsdk:"-"`
}

func (to *DeleteLoggedModelRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteLoggedModelRequest) {
}

func (to *DeleteLoggedModelRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteLoggedModelRequest) {
}

func (m DeleteLoggedModelRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteLoggedModelRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteLoggedModelRequest
// only implements ToObjectValue() and Type().
func (m DeleteLoggedModelRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_id": m.ModelId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteLoggedModelRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_id": types.StringType,
		},
	}
}

type DeleteLoggedModelResponse struct {
}

func (to *DeleteLoggedModelResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteLoggedModelResponse) {
}

func (to *DeleteLoggedModelResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteLoggedModelResponse) {
}

func (m DeleteLoggedModelResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteLoggedModelResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteLoggedModelResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteLoggedModelResponse
// only implements ToObjectValue() and Type().
func (m DeleteLoggedModelResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteLoggedModelResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteLoggedModelTagRequest struct {
	// The ID of the logged model to delete the tag from.
	ModelId types.String `tfsdk:"-"`
	// The tag key.
	TagKey types.String `tfsdk:"-"`
}

func (to *DeleteLoggedModelTagRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteLoggedModelTagRequest) {
}

func (to *DeleteLoggedModelTagRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteLoggedModelTagRequest) {
}

func (m DeleteLoggedModelTagRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteLoggedModelTagRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteLoggedModelTagRequest
// only implements ToObjectValue() and Type().
func (m DeleteLoggedModelTagRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_id": m.ModelId,
			"tag_key":  m.TagKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteLoggedModelTagRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_id": types.StringType,
			"tag_key":  types.StringType,
		},
	}
}

type DeleteLoggedModelTagResponse struct {
}

func (to *DeleteLoggedModelTagResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteLoggedModelTagResponse) {
}

func (to *DeleteLoggedModelTagResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteLoggedModelTagResponse) {
}

func (m DeleteLoggedModelTagResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteLoggedModelTagResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteLoggedModelTagResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteLoggedModelTagResponse
// only implements ToObjectValue() and Type().
func (m DeleteLoggedModelTagResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteLoggedModelTagResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteModelRequest struct {
	// Registered model unique name identifier.
	Name types.String `tfsdk:"-"`
}

func (to *DeleteModelRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteModelRequest) {
}

func (to *DeleteModelRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteModelRequest) {
}

func (m DeleteModelRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteModelRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelRequest
// only implements ToObjectValue() and Type().
func (m DeleteModelRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteModelRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteModelResponse struct {
}

func (to *DeleteModelResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteModelResponse) {
}

func (to *DeleteModelResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteModelResponse) {
}

func (m DeleteModelResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteModelResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteModelResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelResponse
// only implements ToObjectValue() and Type().
func (m DeleteModelResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteModelResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteModelTagRequest struct {
	// Name of the tag. The name must be an exact match; wild-card deletion is
	// not supported. Maximum size is 250 bytes.
	Key types.String `tfsdk:"-"`
	// Name of the registered model that the tag was logged under.
	Name types.String `tfsdk:"-"`
}

func (to *DeleteModelTagRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteModelTagRequest) {
}

func (to *DeleteModelTagRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteModelTagRequest) {
}

func (m DeleteModelTagRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteModelTagRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelTagRequest
// only implements ToObjectValue() and Type().
func (m DeleteModelTagRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":  m.Key,
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteModelTagRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":  types.StringType,
			"name": types.StringType,
		},
	}
}

type DeleteModelTagResponse struct {
}

func (to *DeleteModelTagResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteModelTagResponse) {
}

func (to *DeleteModelTagResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteModelTagResponse) {
}

func (m DeleteModelTagResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteModelTagResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteModelTagResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelTagResponse
// only implements ToObjectValue() and Type().
func (m DeleteModelTagResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteModelTagResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteModelVersionRequest struct {
	// Name of the registered model
	Name types.String `tfsdk:"-"`
	// Model version number
	Version types.String `tfsdk:"-"`
}

func (to *DeleteModelVersionRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteModelVersionRequest) {
}

func (to *DeleteModelVersionRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteModelVersionRequest) {
}

func (m DeleteModelVersionRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteModelVersionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelVersionRequest
// only implements ToObjectValue() and Type().
func (m DeleteModelVersionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":    m.Name,
			"version": m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteModelVersionRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":    types.StringType,
			"version": types.StringType,
		},
	}
}

type DeleteModelVersionResponse struct {
}

func (to *DeleteModelVersionResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteModelVersionResponse) {
}

func (to *DeleteModelVersionResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteModelVersionResponse) {
}

func (m DeleteModelVersionResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteModelVersionResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteModelVersionResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelVersionResponse
// only implements ToObjectValue() and Type().
func (m DeleteModelVersionResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteModelVersionResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteModelVersionTagRequest struct {
	// Name of the tag. The name must be an exact match; wild-card deletion is
	// not supported. Maximum size is 250 bytes.
	Key types.String `tfsdk:"-"`
	// Name of the registered model that the tag was logged under.
	Name types.String `tfsdk:"-"`
	// Model version number that the tag was logged under.
	Version types.String `tfsdk:"-"`
}

func (to *DeleteModelVersionTagRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteModelVersionTagRequest) {
}

func (to *DeleteModelVersionTagRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteModelVersionTagRequest) {
}

func (m DeleteModelVersionTagRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteModelVersionTagRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelVersionTagRequest
// only implements ToObjectValue() and Type().
func (m DeleteModelVersionTagRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":     m.Key,
			"name":    m.Name,
			"version": m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteModelVersionTagRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":     types.StringType,
			"name":    types.StringType,
			"version": types.StringType,
		},
	}
}

type DeleteModelVersionTagResponse struct {
}

func (to *DeleteModelVersionTagResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteModelVersionTagResponse) {
}

func (to *DeleteModelVersionTagResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteModelVersionTagResponse) {
}

func (m DeleteModelVersionTagResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteModelVersionTagResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteModelVersionTagResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelVersionTagResponse
// only implements ToObjectValue() and Type().
func (m DeleteModelVersionTagResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteModelVersionTagResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteOnlineStoreRequest struct {
	// Name of the online store to delete.
	Name types.String `tfsdk:"-"`
}

func (to *DeleteOnlineStoreRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteOnlineStoreRequest) {
}

func (to *DeleteOnlineStoreRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteOnlineStoreRequest) {
}

func (m DeleteOnlineStoreRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteOnlineStoreRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteOnlineStoreRequest
// only implements ToObjectValue() and Type().
func (m DeleteOnlineStoreRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteOnlineStoreRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteRun struct {
	// ID of the run to delete.
	RunId types.String `tfsdk:"run_id"`
}

func (to *DeleteRun) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteRun) {
}

func (to *DeleteRun) SyncFieldsDuringRead(ctx context.Context, from DeleteRun) {
}

func (m DeleteRun) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteRun) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRun
// only implements ToObjectValue() and Type().
func (m DeleteRun) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run_id": m.RunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteRun) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run_id": types.StringType,
		},
	}
}

type DeleteRunResponse struct {
}

func (to *DeleteRunResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteRunResponse) {
}

func (to *DeleteRunResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteRunResponse) {
}

func (m DeleteRunResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteRunResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteRunResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRunResponse
// only implements ToObjectValue() and Type().
func (m DeleteRunResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteRunResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteRuns struct {
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

func (to *DeleteRuns) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteRuns) {
}

func (to *DeleteRuns) SyncFieldsDuringRead(ctx context.Context, from DeleteRuns) {
}

func (m DeleteRuns) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteRuns) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRuns
// only implements ToObjectValue() and Type().
func (m DeleteRuns) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id":        m.ExperimentId,
			"max_runs":             m.MaxRuns,
			"max_timestamp_millis": m.MaxTimestampMillis,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteRuns) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id":        types.StringType,
			"max_runs":             types.Int64Type,
			"max_timestamp_millis": types.Int64Type,
		},
	}
}

type DeleteRunsResponse struct {
	// The number of runs deleted.
	RunsDeleted types.Int64 `tfsdk:"runs_deleted"`
}

func (to *DeleteRunsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteRunsResponse) {
}

func (to *DeleteRunsResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteRunsResponse) {
}

func (m DeleteRunsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteRunsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRunsResponse
// only implements ToObjectValue() and Type().
func (m DeleteRunsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"runs_deleted": m.RunsDeleted,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteRunsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"runs_deleted": types.Int64Type,
		},
	}
}

type DeleteTag struct {
	// Name of the tag. Maximum size is 255 bytes. Must be provided.
	Key types.String `tfsdk:"key"`
	// ID of the run that the tag was logged under. Must be provided.
	RunId types.String `tfsdk:"run_id"`
}

func (to *DeleteTag) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteTag) {
}

func (to *DeleteTag) SyncFieldsDuringRead(ctx context.Context, from DeleteTag) {
}

func (m DeleteTag) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteTag) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteTag
// only implements ToObjectValue() and Type().
func (m DeleteTag) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":    m.Key,
			"run_id": m.RunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteTag) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":    types.StringType,
			"run_id": types.StringType,
		},
	}
}

type DeleteTagResponse struct {
}

func (to *DeleteTagResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteTagResponse) {
}

func (to *DeleteTagResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteTagResponse) {
}

func (m DeleteTagResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteTagResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteTagResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteTagResponse
// only implements ToObjectValue() and Type().
func (m DeleteTagResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteTagResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteTransitionRequestRequest struct {
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

func (to *DeleteTransitionRequestRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteTransitionRequestRequest) {
}

func (to *DeleteTransitionRequestRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteTransitionRequestRequest) {
}

func (m DeleteTransitionRequestRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteTransitionRequestRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteTransitionRequestRequest
// only implements ToObjectValue() and Type().
func (m DeleteTransitionRequestRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m DeleteTransitionRequestRequest) Type(ctx context.Context) attr.Type {
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

type DeleteTransitionRequestResponse struct {
	// New activity generated as a result of this operation.
	Activity types.Object `tfsdk:"activity"`
}

func (to *DeleteTransitionRequestResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteTransitionRequestResponse) {
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

func (to *DeleteTransitionRequestResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteTransitionRequestResponse) {
	if !from.Activity.IsNull() && !from.Activity.IsUnknown() {
		if toActivity, ok := to.GetActivity(ctx); ok {
			if fromActivity, ok := from.GetActivity(ctx); ok {
				toActivity.SyncFieldsDuringRead(ctx, fromActivity)
				to.SetActivity(ctx, toActivity)
			}
		}
	}
}

func (m DeleteTransitionRequestResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["activity"] = attrs["activity"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteTransitionRequestResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteTransitionRequestResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"activity": reflect.TypeOf(Activity{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteTransitionRequestResponse
// only implements ToObjectValue() and Type().
func (m DeleteTransitionRequestResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"activity": m.Activity,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteTransitionRequestResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"activity": Activity{}.Type(ctx),
		},
	}
}

// GetActivity returns the value of the Activity field in DeleteTransitionRequestResponse as
// a Activity value.
// If the field is unknown or null, the boolean return value is false.
func (m *DeleteTransitionRequestResponse) GetActivity(ctx context.Context) (Activity, bool) {
	var e Activity
	if m.Activity.IsNull() || m.Activity.IsUnknown() {
		return e, false
	}
	var v Activity
	d := m.Activity.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetActivity sets the value of the Activity field in DeleteTransitionRequestResponse.
func (m *DeleteTransitionRequestResponse) SetActivity(ctx context.Context, v Activity) {
	vs := v.ToObjectValue(ctx)
	m.Activity = vs
}

type DeleteWebhookRequest struct {
	// Webhook ID required to delete a registry webhook.
	Id types.String `tfsdk:"-"`
}

func (to *DeleteWebhookRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteWebhookRequest) {
}

func (to *DeleteWebhookRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteWebhookRequest) {
}

func (m DeleteWebhookRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteWebhookRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWebhookRequest
// only implements ToObjectValue() and Type().
func (m DeleteWebhookRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteWebhookRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteWebhookResponse struct {
}

func (to *DeleteWebhookResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteWebhookResponse) {
}

func (to *DeleteWebhookResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteWebhookResponse) {
}

func (m DeleteWebhookResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteWebhookResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteWebhookResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWebhookResponse
// only implements ToObjectValue() and Type().
func (m DeleteWebhookResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteWebhookResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeltaTableSource struct {
	// The entity columns of the Delta table.
	EntityColumns types.List `tfsdk:"entity_columns"`
	// The full three-part (catalog, schema, table) name of the Delta table.
	FullName types.String `tfsdk:"full_name"`
	// The timeseries column of the Delta table.
	TimeseriesColumn types.String `tfsdk:"timeseries_column"`
}

func (to *DeltaTableSource) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeltaTableSource) {
}

func (to *DeltaTableSource) SyncFieldsDuringRead(ctx context.Context, from DeltaTableSource) {
}

func (m DeltaTableSource) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeltaTableSource) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"entity_columns": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeltaTableSource
// only implements ToObjectValue() and Type().
func (m DeltaTableSource) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"entity_columns":    m.EntityColumns,
			"full_name":         m.FullName,
			"timeseries_column": m.TimeseriesColumn,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeltaTableSource) Type(ctx context.Context) attr.Type {
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

// GetEntityColumns returns the value of the EntityColumns field in DeltaTableSource as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *DeltaTableSource) GetEntityColumns(ctx context.Context) ([]types.String, bool) {
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

// SetEntityColumns sets the value of the EntityColumns field in DeltaTableSource.
func (m *DeltaTableSource) SetEntityColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["entity_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EntityColumns = types.ListValueMust(t, vs)
}

// An experiment and its metadata.
type Experiment struct {
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

func (to *Experiment) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Experiment) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *Experiment) SyncFieldsDuringRead(ctx context.Context, from Experiment) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m Experiment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Experiment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(ExperimentTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Experiment
// only implements ToObjectValue() and Type().
func (m Experiment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m Experiment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"artifact_location": types.StringType,
			"creation_time":     types.Int64Type,
			"experiment_id":     types.StringType,
			"last_update_time":  types.Int64Type,
			"lifecycle_stage":   types.StringType,
			"name":              types.StringType,
			"tags": basetypes.ListType{
				ElemType: ExperimentTag{}.Type(ctx),
			},
		},
	}
}

// GetTags returns the value of the Tags field in Experiment as
// a slice of ExperimentTag values.
// If the field is unknown or null, the boolean return value is false.
func (m *Experiment) GetTags(ctx context.Context) ([]ExperimentTag, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []ExperimentTag
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in Experiment.
func (m *Experiment) SetTags(ctx context.Context, v []ExperimentTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type ExperimentAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`

	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (to *ExperimentAccessControlRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExperimentAccessControlRequest) {
}

func (to *ExperimentAccessControlRequest) SyncFieldsDuringRead(ctx context.Context, from ExperimentAccessControlRequest) {
}

func (m ExperimentAccessControlRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ExperimentAccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExperimentAccessControlRequest
// only implements ToObjectValue() and Type().
func (m ExperimentAccessControlRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ExperimentAccessControlRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type ExperimentAccessControlResponse struct {
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

func (to *ExperimentAccessControlResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExperimentAccessControlResponse) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (to *ExperimentAccessControlResponse) SyncFieldsDuringRead(ctx context.Context, from ExperimentAccessControlResponse) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (m ExperimentAccessControlResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ExperimentAccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(ExperimentPermission{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExperimentAccessControlResponse
// only implements ToObjectValue() and Type().
func (m ExperimentAccessControlResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ExperimentAccessControlResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: ExperimentPermission{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

// GetAllPermissions returns the value of the AllPermissions field in ExperimentAccessControlResponse as
// a slice of ExperimentPermission values.
// If the field is unknown or null, the boolean return value is false.
func (m *ExperimentAccessControlResponse) GetAllPermissions(ctx context.Context) ([]ExperimentPermission, bool) {
	if m.AllPermissions.IsNull() || m.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []ExperimentPermission
	d := m.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in ExperimentAccessControlResponse.
func (m *ExperimentAccessControlResponse) SetAllPermissions(ctx context.Context, v []ExperimentPermission) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllPermissions = types.ListValueMust(t, vs)
}

type ExperimentPermission struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *ExperimentPermission) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExperimentPermission) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (to *ExperimentPermission) SyncFieldsDuringRead(ctx context.Context, from ExperimentPermission) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (m ExperimentPermission) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ExperimentPermission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExperimentPermission
// only implements ToObjectValue() and Type().
func (m ExperimentPermission) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             m.Inherited,
			"inherited_from_object": m.InheritedFromObject,
			"permission_level":      m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExperimentPermission) Type(ctx context.Context) attr.Type {
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

// GetInheritedFromObject returns the value of the InheritedFromObject field in ExperimentPermission as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ExperimentPermission) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
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

// SetInheritedFromObject sets the value of the InheritedFromObject field in ExperimentPermission.
func (m *ExperimentPermission) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InheritedFromObject = types.ListValueMust(t, vs)
}

type ExperimentPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (to *ExperimentPermissions) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExperimentPermissions) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *ExperimentPermissions) SyncFieldsDuringRead(ctx context.Context, from ExperimentPermissions) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m ExperimentPermissions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ExperimentPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(ExperimentAccessControlResponse{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExperimentPermissions
// only implements ToObjectValue() and Type().
func (m ExperimentPermissions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"object_id":           m.ObjectId,
			"object_type":         m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExperimentPermissions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: ExperimentAccessControlResponse{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in ExperimentPermissions as
// a slice of ExperimentAccessControlResponse values.
// If the field is unknown or null, the boolean return value is false.
func (m *ExperimentPermissions) GetAccessControlList(ctx context.Context) ([]ExperimentAccessControlResponse, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []ExperimentAccessControlResponse
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in ExperimentPermissions.
func (m *ExperimentPermissions) SetAccessControlList(ctx context.Context, v []ExperimentAccessControlResponse) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

type ExperimentPermissionsDescription struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *ExperimentPermissionsDescription) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExperimentPermissionsDescription) {
}

func (to *ExperimentPermissionsDescription) SyncFieldsDuringRead(ctx context.Context, from ExperimentPermissionsDescription) {
}

func (m ExperimentPermissionsDescription) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ExperimentPermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExperimentPermissionsDescription
// only implements ToObjectValue() and Type().
func (m ExperimentPermissionsDescription) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      m.Description,
			"permission_level": m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExperimentPermissionsDescription) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type ExperimentPermissionsRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// The experiment for which to get or manage permissions.
	ExperimentId types.String `tfsdk:"-"`
}

func (to *ExperimentPermissionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExperimentPermissionsRequest) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *ExperimentPermissionsRequest) SyncFieldsDuringRead(ctx context.Context, from ExperimentPermissionsRequest) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m ExperimentPermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ExperimentPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(ExperimentAccessControlRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExperimentPermissionsRequest
// only implements ToObjectValue() and Type().
func (m ExperimentPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"experiment_id":       m.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExperimentPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: ExperimentAccessControlRequest{}.Type(ctx),
			},
			"experiment_id": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in ExperimentPermissionsRequest as
// a slice of ExperimentAccessControlRequest values.
// If the field is unknown or null, the boolean return value is false.
func (m *ExperimentPermissionsRequest) GetAccessControlList(ctx context.Context) ([]ExperimentAccessControlRequest, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []ExperimentAccessControlRequest
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in ExperimentPermissionsRequest.
func (m *ExperimentPermissionsRequest) SetAccessControlList(ctx context.Context, v []ExperimentAccessControlRequest) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

// A tag for an experiment.
type ExperimentTag struct {
	// The tag key.
	Key types.String `tfsdk:"key"`
	// The tag value.
	Value types.String `tfsdk:"value"`
}

func (to *ExperimentTag) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExperimentTag) {
}

func (to *ExperimentTag) SyncFieldsDuringRead(ctx context.Context, from ExperimentTag) {
}

func (m ExperimentTag) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ExperimentTag) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExperimentTag
// only implements ToObjectValue() and Type().
func (m ExperimentTag) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExperimentTag) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

type Feature struct {
	// The description of the feature.
	Description types.String `tfsdk:"description"`
	// The full three-part name (catalog, schema, name) of the feature.
	FullName types.String `tfsdk:"full_name"`
	// The function by which the feature is computed.
	Function types.Object `tfsdk:"function"`
	// The input columns from which the feature is computed.
	Inputs types.List `tfsdk:"inputs"`
	// The data source of the feature.
	Source types.Object `tfsdk:"source"`
	// The time window in which the feature is computed.
	TimeWindow types.Object `tfsdk:"time_window"`
}

func (to *Feature) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Feature) {
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

func (to *Feature) SyncFieldsDuringRead(ctx context.Context, from Feature) {
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

func (m Feature) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["full_name"] = attrs["full_name"].SetRequired()
	attrs["full_name"] = attrs["full_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["function"] = attrs["function"].SetRequired()
	attrs["function"] = attrs["function"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["inputs"] = attrs["inputs"].SetRequired()
	attrs["inputs"] = attrs["inputs"].(tfschema.ListAttributeBuilder).AddPlanModifier(listplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["source"] = attrs["source"].SetRequired()
	attrs["source"] = attrs["source"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["time_window"] = attrs["time_window"].SetRequired()
	attrs["time_window"] = attrs["time_window"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Feature.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Feature) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"function":    reflect.TypeOf(Function{}),
		"inputs":      reflect.TypeOf(types.String{}),
		"source":      reflect.TypeOf(DataSource{}),
		"time_window": reflect.TypeOf(TimeWindow{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Feature
// only implements ToObjectValue() and Type().
func (m Feature) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m Feature) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"full_name":   types.StringType,
			"function":    Function{}.Type(ctx),
			"inputs": basetypes.ListType{
				ElemType: types.StringType,
			},
			"source":      DataSource{}.Type(ctx),
			"time_window": TimeWindow{}.Type(ctx),
		},
	}
}

// GetFunction returns the value of the Function field in Feature as
// a Function value.
// If the field is unknown or null, the boolean return value is false.
func (m *Feature) GetFunction(ctx context.Context) (Function, bool) {
	var e Function
	if m.Function.IsNull() || m.Function.IsUnknown() {
		return e, false
	}
	var v Function
	d := m.Function.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFunction sets the value of the Function field in Feature.
func (m *Feature) SetFunction(ctx context.Context, v Function) {
	vs := v.ToObjectValue(ctx)
	m.Function = vs
}

// GetInputs returns the value of the Inputs field in Feature as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *Feature) GetInputs(ctx context.Context) ([]types.String, bool) {
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

// SetInputs sets the value of the Inputs field in Feature.
func (m *Feature) SetInputs(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["inputs"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Inputs = types.ListValueMust(t, vs)
}

// GetSource returns the value of the Source field in Feature as
// a DataSource value.
// If the field is unknown or null, the boolean return value is false.
func (m *Feature) GetSource(ctx context.Context) (DataSource, bool) {
	var e DataSource
	if m.Source.IsNull() || m.Source.IsUnknown() {
		return e, false
	}
	var v DataSource
	d := m.Source.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSource sets the value of the Source field in Feature.
func (m *Feature) SetSource(ctx context.Context, v DataSource) {
	vs := v.ToObjectValue(ctx)
	m.Source = vs
}

// GetTimeWindow returns the value of the TimeWindow field in Feature as
// a TimeWindow value.
// If the field is unknown or null, the boolean return value is false.
func (m *Feature) GetTimeWindow(ctx context.Context) (TimeWindow, bool) {
	var e TimeWindow
	if m.TimeWindow.IsNull() || m.TimeWindow.IsUnknown() {
		return e, false
	}
	var v TimeWindow
	d := m.TimeWindow.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTimeWindow sets the value of the TimeWindow field in Feature.
func (m *Feature) SetTimeWindow(ctx context.Context, v TimeWindow) {
	vs := v.ToObjectValue(ctx)
	m.TimeWindow = vs
}

type FeatureLineage struct {
	// List of feature specs that contain this feature.
	FeatureSpecs types.List `tfsdk:"feature_specs"`
	// List of Unity Catalog models that were trained on this feature.
	Models types.List `tfsdk:"models"`
	// List of online features that use this feature as source.
	OnlineFeatures types.List `tfsdk:"online_features"`
}

func (to *FeatureLineage) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FeatureLineage) {
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

func (to *FeatureLineage) SyncFieldsDuringRead(ctx context.Context, from FeatureLineage) {
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

func (m FeatureLineage) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m FeatureLineage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"feature_specs":   reflect.TypeOf(FeatureLineageFeatureSpec{}),
		"models":          reflect.TypeOf(FeatureLineageModel{}),
		"online_features": reflect.TypeOf(FeatureLineageOnlineFeature{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FeatureLineage
// only implements ToObjectValue() and Type().
func (m FeatureLineage) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_specs":   m.FeatureSpecs,
			"models":          m.Models,
			"online_features": m.OnlineFeatures,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FeatureLineage) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"feature_specs": basetypes.ListType{
				ElemType: FeatureLineageFeatureSpec{}.Type(ctx),
			},
			"models": basetypes.ListType{
				ElemType: FeatureLineageModel{}.Type(ctx),
			},
			"online_features": basetypes.ListType{
				ElemType: FeatureLineageOnlineFeature{}.Type(ctx),
			},
		},
	}
}

// GetFeatureSpecs returns the value of the FeatureSpecs field in FeatureLineage as
// a slice of FeatureLineageFeatureSpec values.
// If the field is unknown or null, the boolean return value is false.
func (m *FeatureLineage) GetFeatureSpecs(ctx context.Context) ([]FeatureLineageFeatureSpec, bool) {
	if m.FeatureSpecs.IsNull() || m.FeatureSpecs.IsUnknown() {
		return nil, false
	}
	var v []FeatureLineageFeatureSpec
	d := m.FeatureSpecs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFeatureSpecs sets the value of the FeatureSpecs field in FeatureLineage.
func (m *FeatureLineage) SetFeatureSpecs(ctx context.Context, v []FeatureLineageFeatureSpec) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["feature_specs"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.FeatureSpecs = types.ListValueMust(t, vs)
}

// GetModels returns the value of the Models field in FeatureLineage as
// a slice of FeatureLineageModel values.
// If the field is unknown or null, the boolean return value is false.
func (m *FeatureLineage) GetModels(ctx context.Context) ([]FeatureLineageModel, bool) {
	if m.Models.IsNull() || m.Models.IsUnknown() {
		return nil, false
	}
	var v []FeatureLineageModel
	d := m.Models.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModels sets the value of the Models field in FeatureLineage.
func (m *FeatureLineage) SetModels(ctx context.Context, v []FeatureLineageModel) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["models"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Models = types.ListValueMust(t, vs)
}

// GetOnlineFeatures returns the value of the OnlineFeatures field in FeatureLineage as
// a slice of FeatureLineageOnlineFeature values.
// If the field is unknown or null, the boolean return value is false.
func (m *FeatureLineage) GetOnlineFeatures(ctx context.Context) ([]FeatureLineageOnlineFeature, bool) {
	if m.OnlineFeatures.IsNull() || m.OnlineFeatures.IsUnknown() {
		return nil, false
	}
	var v []FeatureLineageOnlineFeature
	d := m.OnlineFeatures.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOnlineFeatures sets the value of the OnlineFeatures field in FeatureLineage.
func (m *FeatureLineage) SetOnlineFeatures(ctx context.Context, v []FeatureLineageOnlineFeature) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["online_features"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.OnlineFeatures = types.ListValueMust(t, vs)
}

type FeatureLineageFeatureSpec struct {
	// The full name of the feature spec in Unity Catalog.
	Name types.String `tfsdk:"name"`
}

func (to *FeatureLineageFeatureSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FeatureLineageFeatureSpec) {
}

func (to *FeatureLineageFeatureSpec) SyncFieldsDuringRead(ctx context.Context, from FeatureLineageFeatureSpec) {
}

func (m FeatureLineageFeatureSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m FeatureLineageFeatureSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FeatureLineageFeatureSpec
// only implements ToObjectValue() and Type().
func (m FeatureLineageFeatureSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FeatureLineageFeatureSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type FeatureLineageModel struct {
	// The full name of the model in Unity Catalog.
	Name types.String `tfsdk:"name"`
	// The version of the model.
	Version types.Int64 `tfsdk:"version"`
}

func (to *FeatureLineageModel) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FeatureLineageModel) {
}

func (to *FeatureLineageModel) SyncFieldsDuringRead(ctx context.Context, from FeatureLineageModel) {
}

func (m FeatureLineageModel) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m FeatureLineageModel) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FeatureLineageModel
// only implements ToObjectValue() and Type().
func (m FeatureLineageModel) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":    m.Name,
			"version": m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FeatureLineageModel) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":    types.StringType,
			"version": types.Int64Type,
		},
	}
}

type FeatureLineageOnlineFeature struct {
	// The name of the online feature (column name).
	FeatureName types.String `tfsdk:"feature_name"`
	// The full name of the online table in Unity Catalog.
	TableName types.String `tfsdk:"table_name"`
}

func (to *FeatureLineageOnlineFeature) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FeatureLineageOnlineFeature) {
}

func (to *FeatureLineageOnlineFeature) SyncFieldsDuringRead(ctx context.Context, from FeatureLineageOnlineFeature) {
}

func (m FeatureLineageOnlineFeature) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m FeatureLineageOnlineFeature) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FeatureLineageOnlineFeature
// only implements ToObjectValue() and Type().
func (m FeatureLineageOnlineFeature) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_name": m.FeatureName,
			"table_name":   m.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FeatureLineageOnlineFeature) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"feature_name": types.StringType,
			"table_name":   types.StringType,
		},
	}
}

// Feature list wrap all the features for a model version
type FeatureList struct {
	Features types.List `tfsdk:"features"`
}

func (to *FeatureList) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FeatureList) {
	if !from.Features.IsNull() && !from.Features.IsUnknown() && to.Features.IsNull() && len(from.Features.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Features, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Features = from.Features
	}
}

func (to *FeatureList) SyncFieldsDuringRead(ctx context.Context, from FeatureList) {
	if !from.Features.IsNull() && !from.Features.IsUnknown() && to.Features.IsNull() && len(from.Features.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Features, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Features = from.Features
	}
}

func (m FeatureList) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m FeatureList) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"features": reflect.TypeOf(LinkedFeature{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FeatureList
// only implements ToObjectValue() and Type().
func (m FeatureList) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"features": m.Features,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FeatureList) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"features": basetypes.ListType{
				ElemType: LinkedFeature{}.Type(ctx),
			},
		},
	}
}

// GetFeatures returns the value of the Features field in FeatureList as
// a slice of LinkedFeature values.
// If the field is unknown or null, the boolean return value is false.
func (m *FeatureList) GetFeatures(ctx context.Context) ([]LinkedFeature, bool) {
	if m.Features.IsNull() || m.Features.IsUnknown() {
		return nil, false
	}
	var v []LinkedFeature
	d := m.Features.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFeatures sets the value of the Features field in FeatureList.
func (m *FeatureList) SetFeatures(ctx context.Context, v []LinkedFeature) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["features"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Features = types.ListValueMust(t, vs)
}

// Represents a tag on a feature in a feature table.
type FeatureTag struct {
	Key types.String `tfsdk:"key"`

	Value types.String `tfsdk:"value"`
}

func (to *FeatureTag) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FeatureTag) {
}

func (to *FeatureTag) SyncFieldsDuringRead(ctx context.Context, from FeatureTag) {
}

func (m FeatureTag) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m FeatureTag) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FeatureTag
// only implements ToObjectValue() and Type().
func (m FeatureTag) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FeatureTag) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

// Metadata of a single artifact file or directory.
type FileInfo struct {
	// The size in bytes of the file. Unset for directories.
	FileSize types.Int64 `tfsdk:"file_size"`
	// Whether the path is a directory.
	IsDir types.Bool `tfsdk:"is_dir"`
	// The path relative to the root artifact directory run.
	Path types.String `tfsdk:"path"`
}

func (to *FileInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FileInfo) {
}

func (to *FileInfo) SyncFieldsDuringRead(ctx context.Context, from FileInfo) {
}

func (m FileInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m FileInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FileInfo
// only implements ToObjectValue() and Type().
func (m FileInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_size": m.FileSize,
			"is_dir":    m.IsDir,
			"path":      m.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FileInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"file_size": types.Int64Type,
			"is_dir":    types.BoolType,
			"path":      types.StringType,
		},
	}
}

type FinalizeLoggedModelRequest struct {
	// The ID of the logged model to finalize.
	ModelId types.String `tfsdk:"-"`
	// Whether or not the model is ready for use.
	// ``"LOGGED_MODEL_UPLOAD_FAILED"`` indicates that something went wrong when
	// logging the model weights / agent code.
	Status types.String `tfsdk:"status"`
}

func (to *FinalizeLoggedModelRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FinalizeLoggedModelRequest) {
}

func (to *FinalizeLoggedModelRequest) SyncFieldsDuringRead(ctx context.Context, from FinalizeLoggedModelRequest) {
}

func (m FinalizeLoggedModelRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m FinalizeLoggedModelRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FinalizeLoggedModelRequest
// only implements ToObjectValue() and Type().
func (m FinalizeLoggedModelRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_id": m.ModelId,
			"status":   m.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FinalizeLoggedModelRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_id": types.StringType,
			"status":   types.StringType,
		},
	}
}

type FinalizeLoggedModelResponse struct {
	// The updated logged model.
	Model types.Object `tfsdk:"model"`
}

func (to *FinalizeLoggedModelResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FinalizeLoggedModelResponse) {
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

func (to *FinalizeLoggedModelResponse) SyncFieldsDuringRead(ctx context.Context, from FinalizeLoggedModelResponse) {
	if !from.Model.IsNull() && !from.Model.IsUnknown() {
		if toModel, ok := to.GetModel(ctx); ok {
			if fromModel, ok := from.GetModel(ctx); ok {
				toModel.SyncFieldsDuringRead(ctx, fromModel)
				to.SetModel(ctx, toModel)
			}
		}
	}
}

func (m FinalizeLoggedModelResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["model"] = attrs["model"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FinalizeLoggedModelResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m FinalizeLoggedModelResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model": reflect.TypeOf(LoggedModel{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FinalizeLoggedModelResponse
// only implements ToObjectValue() and Type().
func (m FinalizeLoggedModelResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model": m.Model,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FinalizeLoggedModelResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model": LoggedModel{}.Type(ctx),
		},
	}
}

// GetModel returns the value of the Model field in FinalizeLoggedModelResponse as
// a LoggedModel value.
// If the field is unknown or null, the boolean return value is false.
func (m *FinalizeLoggedModelResponse) GetModel(ctx context.Context) (LoggedModel, bool) {
	var e LoggedModel
	if m.Model.IsNull() || m.Model.IsUnknown() {
		return e, false
	}
	var v LoggedModel
	d := m.Model.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModel sets the value of the Model field in FinalizeLoggedModelResponse.
func (m *FinalizeLoggedModelResponse) SetModel(ctx context.Context, v LoggedModel) {
	vs := v.ToObjectValue(ctx)
	m.Model = vs
}

// Represents a forecasting experiment with its unique identifier, URL, and
// state.
type ForecastingExperiment struct {
	// The unique ID for the forecasting experiment.
	ExperimentId types.String `tfsdk:"experiment_id"`
	// The URL to the forecasting experiment page.
	ExperimentPageUrl types.String `tfsdk:"experiment_page_url"`
	// The current state of the forecasting experiment.
	State types.String `tfsdk:"state"`
}

func (to *ForecastingExperiment) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ForecastingExperiment) {
}

func (to *ForecastingExperiment) SyncFieldsDuringRead(ctx context.Context, from ForecastingExperiment) {
}

func (m ForecastingExperiment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ForecastingExperiment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ForecastingExperiment
// only implements ToObjectValue() and Type().
func (m ForecastingExperiment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id":       m.ExperimentId,
			"experiment_page_url": m.ExperimentPageUrl,
			"state":               m.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ForecastingExperiment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id":       types.StringType,
			"experiment_page_url": types.StringType,
			"state":               types.StringType,
		},
	}
}

type Function struct {
	// Extra parameters for parameterized functions.
	ExtraParameters types.List `tfsdk:"extra_parameters"`
	// The type of the function.
	FunctionType types.String `tfsdk:"function_type"`
}

func (to *Function) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Function) {
	if !from.ExtraParameters.IsNull() && !from.ExtraParameters.IsUnknown() && to.ExtraParameters.IsNull() && len(from.ExtraParameters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ExtraParameters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ExtraParameters = from.ExtraParameters
	}
}

func (to *Function) SyncFieldsDuringRead(ctx context.Context, from Function) {
	if !from.ExtraParameters.IsNull() && !from.ExtraParameters.IsUnknown() && to.ExtraParameters.IsNull() && len(from.ExtraParameters.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ExtraParameters, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ExtraParameters = from.ExtraParameters
	}
}

func (m Function) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Function) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"extra_parameters": reflect.TypeOf(FunctionExtraParameter{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Function
// only implements ToObjectValue() and Type().
func (m Function) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"extra_parameters": m.ExtraParameters,
			"function_type":    m.FunctionType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Function) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"extra_parameters": basetypes.ListType{
				ElemType: FunctionExtraParameter{}.Type(ctx),
			},
			"function_type": types.StringType,
		},
	}
}

// GetExtraParameters returns the value of the ExtraParameters field in Function as
// a slice of FunctionExtraParameter values.
// If the field is unknown or null, the boolean return value is false.
func (m *Function) GetExtraParameters(ctx context.Context) ([]FunctionExtraParameter, bool) {
	if m.ExtraParameters.IsNull() || m.ExtraParameters.IsUnknown() {
		return nil, false
	}
	var v []FunctionExtraParameter
	d := m.ExtraParameters.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExtraParameters sets the value of the ExtraParameters field in Function.
func (m *Function) SetExtraParameters(ctx context.Context, v []FunctionExtraParameter) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["extra_parameters"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ExtraParameters = types.ListValueMust(t, vs)
}

type FunctionExtraParameter struct {
	// The name of the parameter.
	Key types.String `tfsdk:"key"`
	// The value of the parameter.
	Value types.String `tfsdk:"value"`
}

func (to *FunctionExtraParameter) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FunctionExtraParameter) {
}

func (to *FunctionExtraParameter) SyncFieldsDuringRead(ctx context.Context, from FunctionExtraParameter) {
}

func (m FunctionExtraParameter) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m FunctionExtraParameter) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FunctionExtraParameter
// only implements ToObjectValue() and Type().
func (m FunctionExtraParameter) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FunctionExtraParameter) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

type GetByNameRequest struct {
	// Name of the associated experiment.
	ExperimentName types.String `tfsdk:"-"`
}

func (to *GetByNameRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetByNameRequest) {
}

func (to *GetByNameRequest) SyncFieldsDuringRead(ctx context.Context, from GetByNameRequest) {
}

func (m GetByNameRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetByNameRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetByNameRequest
// only implements ToObjectValue() and Type().
func (m GetByNameRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_name": m.ExperimentName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetByNameRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_name": types.StringType,
		},
	}
}

type GetExperimentByNameResponse struct {
	// Experiment details.
	Experiment types.Object `tfsdk:"experiment"`
}

func (to *GetExperimentByNameResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetExperimentByNameResponse) {
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

func (to *GetExperimentByNameResponse) SyncFieldsDuringRead(ctx context.Context, from GetExperimentByNameResponse) {
	if !from.Experiment.IsNull() && !from.Experiment.IsUnknown() {
		if toExperiment, ok := to.GetExperiment(ctx); ok {
			if fromExperiment, ok := from.GetExperiment(ctx); ok {
				toExperiment.SyncFieldsDuringRead(ctx, fromExperiment)
				to.SetExperiment(ctx, toExperiment)
			}
		}
	}
}

func (m GetExperimentByNameResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["experiment"] = attrs["experiment"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetExperimentByNameResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetExperimentByNameResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"experiment": reflect.TypeOf(Experiment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExperimentByNameResponse
// only implements ToObjectValue() and Type().
func (m GetExperimentByNameResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment": m.Experiment,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetExperimentByNameResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment": Experiment{}.Type(ctx),
		},
	}
}

// GetExperiment returns the value of the Experiment field in GetExperimentByNameResponse as
// a Experiment value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetExperimentByNameResponse) GetExperiment(ctx context.Context) (Experiment, bool) {
	var e Experiment
	if m.Experiment.IsNull() || m.Experiment.IsUnknown() {
		return e, false
	}
	var v Experiment
	d := m.Experiment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExperiment sets the value of the Experiment field in GetExperimentByNameResponse.
func (m *GetExperimentByNameResponse) SetExperiment(ctx context.Context, v Experiment) {
	vs := v.ToObjectValue(ctx)
	m.Experiment = vs
}

type GetExperimentPermissionLevelsRequest struct {
	// The experiment for which to get or manage permissions.
	ExperimentId types.String `tfsdk:"-"`
}

func (to *GetExperimentPermissionLevelsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetExperimentPermissionLevelsRequest) {
}

func (to *GetExperimentPermissionLevelsRequest) SyncFieldsDuringRead(ctx context.Context, from GetExperimentPermissionLevelsRequest) {
}

func (m GetExperimentPermissionLevelsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetExperimentPermissionLevelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExperimentPermissionLevelsRequest
// only implements ToObjectValue() and Type().
func (m GetExperimentPermissionLevelsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": m.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetExperimentPermissionLevelsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
		},
	}
}

type GetExperimentPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (to *GetExperimentPermissionLevelsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetExperimentPermissionLevelsResponse) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (to *GetExperimentPermissionLevelsResponse) SyncFieldsDuringRead(ctx context.Context, from GetExperimentPermissionLevelsResponse) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (m GetExperimentPermissionLevelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetExperimentPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(ExperimentPermissionsDescription{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExperimentPermissionLevelsResponse
// only implements ToObjectValue() and Type().
func (m GetExperimentPermissionLevelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": m.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetExperimentPermissionLevelsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: ExperimentPermissionsDescription{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetExperimentPermissionLevelsResponse as
// a slice of ExperimentPermissionsDescription values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetExperimentPermissionLevelsResponse) GetPermissionLevels(ctx context.Context) ([]ExperimentPermissionsDescription, bool) {
	if m.PermissionLevels.IsNull() || m.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []ExperimentPermissionsDescription
	d := m.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetExperimentPermissionLevelsResponse.
func (m *GetExperimentPermissionLevelsResponse) SetPermissionLevels(ctx context.Context, v []ExperimentPermissionsDescription) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PermissionLevels = types.ListValueMust(t, vs)
}

type GetExperimentPermissionsRequest struct {
	// The experiment for which to get or manage permissions.
	ExperimentId types.String `tfsdk:"-"`
}

func (to *GetExperimentPermissionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetExperimentPermissionsRequest) {
}

func (to *GetExperimentPermissionsRequest) SyncFieldsDuringRead(ctx context.Context, from GetExperimentPermissionsRequest) {
}

func (m GetExperimentPermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetExperimentPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExperimentPermissionsRequest
// only implements ToObjectValue() and Type().
func (m GetExperimentPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": m.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetExperimentPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
		},
	}
}

type GetExperimentRequest struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"-"`
}

func (to *GetExperimentRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetExperimentRequest) {
}

func (to *GetExperimentRequest) SyncFieldsDuringRead(ctx context.Context, from GetExperimentRequest) {
}

func (m GetExperimentRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetExperimentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExperimentRequest
// only implements ToObjectValue() and Type().
func (m GetExperimentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": m.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetExperimentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
		},
	}
}

type GetExperimentResponse struct {
	// Experiment details.
	Experiment types.Object `tfsdk:"experiment"`
}

func (to *GetExperimentResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetExperimentResponse) {
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

func (to *GetExperimentResponse) SyncFieldsDuringRead(ctx context.Context, from GetExperimentResponse) {
	if !from.Experiment.IsNull() && !from.Experiment.IsUnknown() {
		if toExperiment, ok := to.GetExperiment(ctx); ok {
			if fromExperiment, ok := from.GetExperiment(ctx); ok {
				toExperiment.SyncFieldsDuringRead(ctx, fromExperiment)
				to.SetExperiment(ctx, toExperiment)
			}
		}
	}
}

func (m GetExperimentResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["experiment"] = attrs["experiment"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetExperimentResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetExperimentResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"experiment": reflect.TypeOf(Experiment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExperimentResponse
// only implements ToObjectValue() and Type().
func (m GetExperimentResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment": m.Experiment,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetExperimentResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment": Experiment{}.Type(ctx),
		},
	}
}

// GetExperiment returns the value of the Experiment field in GetExperimentResponse as
// a Experiment value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetExperimentResponse) GetExperiment(ctx context.Context) (Experiment, bool) {
	var e Experiment
	if m.Experiment.IsNull() || m.Experiment.IsUnknown() {
		return e, false
	}
	var v Experiment
	d := m.Experiment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExperiment sets the value of the Experiment field in GetExperimentResponse.
func (m *GetExperimentResponse) SetExperiment(ctx context.Context, v Experiment) {
	vs := v.ToObjectValue(ctx)
	m.Experiment = vs
}

type GetFeatureLineageRequest struct {
	// The name of the feature.
	FeatureName types.String `tfsdk:"-"`
	// The full name of the feature table in Unity Catalog.
	TableName types.String `tfsdk:"-"`
}

func (to *GetFeatureLineageRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetFeatureLineageRequest) {
}

func (to *GetFeatureLineageRequest) SyncFieldsDuringRead(ctx context.Context, from GetFeatureLineageRequest) {
}

func (m GetFeatureLineageRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetFeatureLineageRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetFeatureLineageRequest
// only implements ToObjectValue() and Type().
func (m GetFeatureLineageRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_name": m.FeatureName,
			"table_name":   m.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetFeatureLineageRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"feature_name": types.StringType,
			"table_name":   types.StringType,
		},
	}
}

type GetFeatureRequest struct {
	// Name of the feature to get.
	FullName types.String `tfsdk:"-"`
}

func (to *GetFeatureRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetFeatureRequest) {
}

func (to *GetFeatureRequest) SyncFieldsDuringRead(ctx context.Context, from GetFeatureRequest) {
}

func (m GetFeatureRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetFeatureRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetFeatureRequest
// only implements ToObjectValue() and Type().
func (m GetFeatureRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"full_name": m.FullName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetFeatureRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"full_name": types.StringType,
		},
	}
}

type GetFeatureTagRequest struct {
	FeatureName types.String `tfsdk:"-"`

	Key types.String `tfsdk:"-"`

	TableName types.String `tfsdk:"-"`
}

func (to *GetFeatureTagRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetFeatureTagRequest) {
}

func (to *GetFeatureTagRequest) SyncFieldsDuringRead(ctx context.Context, from GetFeatureTagRequest) {
}

func (m GetFeatureTagRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetFeatureTagRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetFeatureTagRequest
// only implements ToObjectValue() and Type().
func (m GetFeatureTagRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_name": m.FeatureName,
			"key":          m.Key,
			"table_name":   m.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetFeatureTagRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"feature_name": types.StringType,
			"key":          types.StringType,
			"table_name":   types.StringType,
		},
	}
}

type GetForecastingExperimentRequest struct {
	// The unique ID of a forecasting experiment
	ExperimentId types.String `tfsdk:"-"`
}

func (to *GetForecastingExperimentRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetForecastingExperimentRequest) {
}

func (to *GetForecastingExperimentRequest) SyncFieldsDuringRead(ctx context.Context, from GetForecastingExperimentRequest) {
}

func (m GetForecastingExperimentRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetForecastingExperimentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetForecastingExperimentRequest
// only implements ToObjectValue() and Type().
func (m GetForecastingExperimentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": m.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetForecastingExperimentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
		},
	}
}

type GetHistoryRequest struct {
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

func (to *GetHistoryRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetHistoryRequest) {
}

func (to *GetHistoryRequest) SyncFieldsDuringRead(ctx context.Context, from GetHistoryRequest) {
}

func (m GetHistoryRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetHistoryRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetHistoryRequest
// only implements ToObjectValue() and Type().
func (m GetHistoryRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m GetHistoryRequest) Type(ctx context.Context) attr.Type {
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

type GetLatestVersionsRequest struct {
	// Registered model unique name identifier.
	Name types.String `tfsdk:"name"`
	// List of stages.
	Stages types.List `tfsdk:"stages"`
}

func (to *GetLatestVersionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetLatestVersionsRequest) {
	if !from.Stages.IsNull() && !from.Stages.IsUnknown() && to.Stages.IsNull() && len(from.Stages.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Stages, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Stages = from.Stages
	}
}

func (to *GetLatestVersionsRequest) SyncFieldsDuringRead(ctx context.Context, from GetLatestVersionsRequest) {
	if !from.Stages.IsNull() && !from.Stages.IsUnknown() && to.Stages.IsNull() && len(from.Stages.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Stages, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Stages = from.Stages
	}
}

func (m GetLatestVersionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetLatestVersionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"stages": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLatestVersionsRequest
// only implements ToObjectValue() and Type().
func (m GetLatestVersionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":   m.Name,
			"stages": m.Stages,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetLatestVersionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"stages": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetStages returns the value of the Stages field in GetLatestVersionsRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetLatestVersionsRequest) GetStages(ctx context.Context) ([]types.String, bool) {
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

// SetStages sets the value of the Stages field in GetLatestVersionsRequest.
func (m *GetLatestVersionsRequest) SetStages(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["stages"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Stages = types.ListValueMust(t, vs)
}

type GetLatestVersionsResponse struct {
	// Latest version models for each requests stage. Only return models with
	// current `READY` status. If no `stages` provided, returns the latest
	// version for each stage, including `"None"`.
	ModelVersions types.List `tfsdk:"model_versions"`
}

func (to *GetLatestVersionsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetLatestVersionsResponse) {
	if !from.ModelVersions.IsNull() && !from.ModelVersions.IsUnknown() && to.ModelVersions.IsNull() && len(from.ModelVersions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ModelVersions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ModelVersions = from.ModelVersions
	}
}

func (to *GetLatestVersionsResponse) SyncFieldsDuringRead(ctx context.Context, from GetLatestVersionsResponse) {
	if !from.ModelVersions.IsNull() && !from.ModelVersions.IsUnknown() && to.ModelVersions.IsNull() && len(from.ModelVersions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ModelVersions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ModelVersions = from.ModelVersions
	}
}

func (m GetLatestVersionsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetLatestVersionsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model_versions": reflect.TypeOf(ModelVersion{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLatestVersionsResponse
// only implements ToObjectValue() and Type().
func (m GetLatestVersionsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_versions": m.ModelVersions,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetLatestVersionsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_versions": basetypes.ListType{
				ElemType: ModelVersion{}.Type(ctx),
			},
		},
	}
}

// GetModelVersions returns the value of the ModelVersions field in GetLatestVersionsResponse as
// a slice of ModelVersion values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetLatestVersionsResponse) GetModelVersions(ctx context.Context) ([]ModelVersion, bool) {
	if m.ModelVersions.IsNull() || m.ModelVersions.IsUnknown() {
		return nil, false
	}
	var v []ModelVersion
	d := m.ModelVersions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModelVersions sets the value of the ModelVersions field in GetLatestVersionsResponse.
func (m *GetLatestVersionsResponse) SetModelVersions(ctx context.Context, v []ModelVersion) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["model_versions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ModelVersions = types.ListValueMust(t, vs)
}

type GetLoggedModelRequest struct {
	// The ID of the logged model to retrieve.
	ModelId types.String `tfsdk:"-"`
}

func (to *GetLoggedModelRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetLoggedModelRequest) {
}

func (to *GetLoggedModelRequest) SyncFieldsDuringRead(ctx context.Context, from GetLoggedModelRequest) {
}

func (m GetLoggedModelRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetLoggedModelRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLoggedModelRequest
// only implements ToObjectValue() and Type().
func (m GetLoggedModelRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_id": m.ModelId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetLoggedModelRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_id": types.StringType,
		},
	}
}

type GetLoggedModelResponse struct {
	// The retrieved logged model.
	Model types.Object `tfsdk:"model"`
}

func (to *GetLoggedModelResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetLoggedModelResponse) {
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

func (to *GetLoggedModelResponse) SyncFieldsDuringRead(ctx context.Context, from GetLoggedModelResponse) {
	if !from.Model.IsNull() && !from.Model.IsUnknown() {
		if toModel, ok := to.GetModel(ctx); ok {
			if fromModel, ok := from.GetModel(ctx); ok {
				toModel.SyncFieldsDuringRead(ctx, fromModel)
				to.SetModel(ctx, toModel)
			}
		}
	}
}

func (m GetLoggedModelResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["model"] = attrs["model"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetLoggedModelResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetLoggedModelResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model": reflect.TypeOf(LoggedModel{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLoggedModelResponse
// only implements ToObjectValue() and Type().
func (m GetLoggedModelResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model": m.Model,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetLoggedModelResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model": LoggedModel{}.Type(ctx),
		},
	}
}

// GetModel returns the value of the Model field in GetLoggedModelResponse as
// a LoggedModel value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetLoggedModelResponse) GetModel(ctx context.Context) (LoggedModel, bool) {
	var e LoggedModel
	if m.Model.IsNull() || m.Model.IsUnknown() {
		return e, false
	}
	var v LoggedModel
	d := m.Model.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModel sets the value of the Model field in GetLoggedModelResponse.
func (m *GetLoggedModelResponse) SetModel(ctx context.Context, v LoggedModel) {
	vs := v.ToObjectValue(ctx)
	m.Model = vs
}

type GetMetricHistoryResponse struct {
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

func (to *GetMetricHistoryResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetMetricHistoryResponse) {
	if !from.Metrics.IsNull() && !from.Metrics.IsUnknown() && to.Metrics.IsNull() && len(from.Metrics.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Metrics, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Metrics = from.Metrics
	}
}

func (to *GetMetricHistoryResponse) SyncFieldsDuringRead(ctx context.Context, from GetMetricHistoryResponse) {
	if !from.Metrics.IsNull() && !from.Metrics.IsUnknown() && to.Metrics.IsNull() && len(from.Metrics.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Metrics, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Metrics = from.Metrics
	}
}

func (m GetMetricHistoryResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetMetricHistoryResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metrics": reflect.TypeOf(Metric{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetMetricHistoryResponse
// only implements ToObjectValue() and Type().
func (m GetMetricHistoryResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metrics":         m.Metrics,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetMetricHistoryResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metrics": basetypes.ListType{
				ElemType: Metric{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetMetrics returns the value of the Metrics field in GetMetricHistoryResponse as
// a slice of Metric values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetMetricHistoryResponse) GetMetrics(ctx context.Context) ([]Metric, bool) {
	if m.Metrics.IsNull() || m.Metrics.IsUnknown() {
		return nil, false
	}
	var v []Metric
	d := m.Metrics.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMetrics sets the value of the Metrics field in GetMetricHistoryResponse.
func (m *GetMetricHistoryResponse) SetMetrics(ctx context.Context, v []Metric) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["metrics"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Metrics = types.ListValueMust(t, vs)
}

type GetModelRequest struct {
	// Registered model unique name identifier.
	Name types.String `tfsdk:"-"`
}

func (to *GetModelRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetModelRequest) {
}

func (to *GetModelRequest) SyncFieldsDuringRead(ctx context.Context, from GetModelRequest) {
}

func (m GetModelRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetModelRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetModelRequest
// only implements ToObjectValue() and Type().
func (m GetModelRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetModelRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetModelResponse struct {
	RegisteredModelDatabricks types.Object `tfsdk:"registered_model_databricks"`
}

func (to *GetModelResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetModelResponse) {
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

func (to *GetModelResponse) SyncFieldsDuringRead(ctx context.Context, from GetModelResponse) {
	if !from.RegisteredModelDatabricks.IsNull() && !from.RegisteredModelDatabricks.IsUnknown() {
		if toRegisteredModelDatabricks, ok := to.GetRegisteredModelDatabricks(ctx); ok {
			if fromRegisteredModelDatabricks, ok := from.GetRegisteredModelDatabricks(ctx); ok {
				toRegisteredModelDatabricks.SyncFieldsDuringRead(ctx, fromRegisteredModelDatabricks)
				to.SetRegisteredModelDatabricks(ctx, toRegisteredModelDatabricks)
			}
		}
	}
}

func (m GetModelResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["registered_model_databricks"] = attrs["registered_model_databricks"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetModelResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetModelResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"registered_model_databricks": reflect.TypeOf(ModelDatabricks{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetModelResponse
// only implements ToObjectValue() and Type().
func (m GetModelResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"registered_model_databricks": m.RegisteredModelDatabricks,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetModelResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"registered_model_databricks": ModelDatabricks{}.Type(ctx),
		},
	}
}

// GetRegisteredModelDatabricks returns the value of the RegisteredModelDatabricks field in GetModelResponse as
// a ModelDatabricks value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetModelResponse) GetRegisteredModelDatabricks(ctx context.Context) (ModelDatabricks, bool) {
	var e ModelDatabricks
	if m.RegisteredModelDatabricks.IsNull() || m.RegisteredModelDatabricks.IsUnknown() {
		return e, false
	}
	var v ModelDatabricks
	d := m.RegisteredModelDatabricks.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRegisteredModelDatabricks sets the value of the RegisteredModelDatabricks field in GetModelResponse.
func (m *GetModelResponse) SetRegisteredModelDatabricks(ctx context.Context, v ModelDatabricks) {
	vs := v.ToObjectValue(ctx)
	m.RegisteredModelDatabricks = vs
}

type GetModelVersionDownloadUriRequest struct {
	// Name of the registered model
	Name types.String `tfsdk:"-"`
	// Model version number
	Version types.String `tfsdk:"-"`
}

func (to *GetModelVersionDownloadUriRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetModelVersionDownloadUriRequest) {
}

func (to *GetModelVersionDownloadUriRequest) SyncFieldsDuringRead(ctx context.Context, from GetModelVersionDownloadUriRequest) {
}

func (m GetModelVersionDownloadUriRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetModelVersionDownloadUriRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetModelVersionDownloadUriRequest
// only implements ToObjectValue() and Type().
func (m GetModelVersionDownloadUriRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":    m.Name,
			"version": m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetModelVersionDownloadUriRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":    types.StringType,
			"version": types.StringType,
		},
	}
}

type GetModelVersionDownloadUriResponse struct {
	// URI corresponding to where artifacts for this model version are stored.
	ArtifactUri types.String `tfsdk:"artifact_uri"`
}

func (to *GetModelVersionDownloadUriResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetModelVersionDownloadUriResponse) {
}

func (to *GetModelVersionDownloadUriResponse) SyncFieldsDuringRead(ctx context.Context, from GetModelVersionDownloadUriResponse) {
}

func (m GetModelVersionDownloadUriResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetModelVersionDownloadUriResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetModelVersionDownloadUriResponse
// only implements ToObjectValue() and Type().
func (m GetModelVersionDownloadUriResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"artifact_uri": m.ArtifactUri,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetModelVersionDownloadUriResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"artifact_uri": types.StringType,
		},
	}
}

type GetModelVersionRequest struct {
	// Name of the registered model
	Name types.String `tfsdk:"-"`
	// Model version number
	Version types.String `tfsdk:"-"`
}

func (to *GetModelVersionRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetModelVersionRequest) {
}

func (to *GetModelVersionRequest) SyncFieldsDuringRead(ctx context.Context, from GetModelVersionRequest) {
}

func (m GetModelVersionRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetModelVersionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetModelVersionRequest
// only implements ToObjectValue() and Type().
func (m GetModelVersionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":    m.Name,
			"version": m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetModelVersionRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":    types.StringType,
			"version": types.StringType,
		},
	}
}

type GetModelVersionResponse struct {
	ModelVersion types.Object `tfsdk:"model_version"`
}

func (to *GetModelVersionResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetModelVersionResponse) {
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

func (to *GetModelVersionResponse) SyncFieldsDuringRead(ctx context.Context, from GetModelVersionResponse) {
	if !from.ModelVersion.IsNull() && !from.ModelVersion.IsUnknown() {
		if toModelVersion, ok := to.GetModelVersion(ctx); ok {
			if fromModelVersion, ok := from.GetModelVersion(ctx); ok {
				toModelVersion.SyncFieldsDuringRead(ctx, fromModelVersion)
				to.SetModelVersion(ctx, toModelVersion)
			}
		}
	}
}

func (m GetModelVersionResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["model_version"] = attrs["model_version"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetModelVersionResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetModelVersionResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model_version": reflect.TypeOf(ModelVersion{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetModelVersionResponse
// only implements ToObjectValue() and Type().
func (m GetModelVersionResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_version": m.ModelVersion,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetModelVersionResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_version": ModelVersion{}.Type(ctx),
		},
	}
}

// GetModelVersion returns the value of the ModelVersion field in GetModelVersionResponse as
// a ModelVersion value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetModelVersionResponse) GetModelVersion(ctx context.Context) (ModelVersion, bool) {
	var e ModelVersion
	if m.ModelVersion.IsNull() || m.ModelVersion.IsUnknown() {
		return e, false
	}
	var v ModelVersion
	d := m.ModelVersion.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModelVersion sets the value of the ModelVersion field in GetModelVersionResponse.
func (m *GetModelVersionResponse) SetModelVersion(ctx context.Context, v ModelVersion) {
	vs := v.ToObjectValue(ctx)
	m.ModelVersion = vs
}

type GetOnlineStoreRequest struct {
	// Name of the online store to get.
	Name types.String `tfsdk:"-"`
}

func (to *GetOnlineStoreRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetOnlineStoreRequest) {
}

func (to *GetOnlineStoreRequest) SyncFieldsDuringRead(ctx context.Context, from GetOnlineStoreRequest) {
}

func (m GetOnlineStoreRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetOnlineStoreRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetOnlineStoreRequest
// only implements ToObjectValue() and Type().
func (m GetOnlineStoreRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetOnlineStoreRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetRegisteredModelPermissionLevelsRequest struct {
	// The registered model for which to get or manage permissions.
	RegisteredModelId types.String `tfsdk:"-"`
}

func (to *GetRegisteredModelPermissionLevelsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetRegisteredModelPermissionLevelsRequest) {
}

func (to *GetRegisteredModelPermissionLevelsRequest) SyncFieldsDuringRead(ctx context.Context, from GetRegisteredModelPermissionLevelsRequest) {
}

func (m GetRegisteredModelPermissionLevelsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetRegisteredModelPermissionLevelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRegisteredModelPermissionLevelsRequest
// only implements ToObjectValue() and Type().
func (m GetRegisteredModelPermissionLevelsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"registered_model_id": m.RegisteredModelId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetRegisteredModelPermissionLevelsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"registered_model_id": types.StringType,
		},
	}
}

type GetRegisteredModelPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (to *GetRegisteredModelPermissionLevelsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetRegisteredModelPermissionLevelsResponse) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (to *GetRegisteredModelPermissionLevelsResponse) SyncFieldsDuringRead(ctx context.Context, from GetRegisteredModelPermissionLevelsResponse) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (m GetRegisteredModelPermissionLevelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetRegisteredModelPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(RegisteredModelPermissionsDescription{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRegisteredModelPermissionLevelsResponse
// only implements ToObjectValue() and Type().
func (m GetRegisteredModelPermissionLevelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": m.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetRegisteredModelPermissionLevelsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: RegisteredModelPermissionsDescription{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetRegisteredModelPermissionLevelsResponse as
// a slice of RegisteredModelPermissionsDescription values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetRegisteredModelPermissionLevelsResponse) GetPermissionLevels(ctx context.Context) ([]RegisteredModelPermissionsDescription, bool) {
	if m.PermissionLevels.IsNull() || m.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []RegisteredModelPermissionsDescription
	d := m.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetRegisteredModelPermissionLevelsResponse.
func (m *GetRegisteredModelPermissionLevelsResponse) SetPermissionLevels(ctx context.Context, v []RegisteredModelPermissionsDescription) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PermissionLevels = types.ListValueMust(t, vs)
}

type GetRegisteredModelPermissionsRequest struct {
	// The registered model for which to get or manage permissions.
	RegisteredModelId types.String `tfsdk:"-"`
}

func (to *GetRegisteredModelPermissionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetRegisteredModelPermissionsRequest) {
}

func (to *GetRegisteredModelPermissionsRequest) SyncFieldsDuringRead(ctx context.Context, from GetRegisteredModelPermissionsRequest) {
}

func (m GetRegisteredModelPermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetRegisteredModelPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRegisteredModelPermissionsRequest
// only implements ToObjectValue() and Type().
func (m GetRegisteredModelPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"registered_model_id": m.RegisteredModelId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetRegisteredModelPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"registered_model_id": types.StringType,
		},
	}
}

type GetRunRequest struct {
	// ID of the run to fetch. Must be provided.
	RunId types.String `tfsdk:"-"`
	// [Deprecated, use `run_id` instead] ID of the run to fetch. This field
	// will be removed in a future MLflow version.
	RunUuid types.String `tfsdk:"-"`
}

func (to *GetRunRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetRunRequest) {
}

func (to *GetRunRequest) SyncFieldsDuringRead(ctx context.Context, from GetRunRequest) {
}

func (m GetRunRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetRunRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRunRequest
// only implements ToObjectValue() and Type().
func (m GetRunRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run_id":   m.RunId,
			"run_uuid": m.RunUuid,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetRunRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run_id":   types.StringType,
			"run_uuid": types.StringType,
		},
	}
}

type GetRunResponse struct {
	// Run metadata (name, start time, etc) and data (metrics, params, and
	// tags).
	Run types.Object `tfsdk:"run"`
}

func (to *GetRunResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetRunResponse) {
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

func (to *GetRunResponse) SyncFieldsDuringRead(ctx context.Context, from GetRunResponse) {
	if !from.Run.IsNull() && !from.Run.IsUnknown() {
		if toRun, ok := to.GetRun(ctx); ok {
			if fromRun, ok := from.GetRun(ctx); ok {
				toRun.SyncFieldsDuringRead(ctx, fromRun)
				to.SetRun(ctx, toRun)
			}
		}
	}
}

func (m GetRunResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["run"] = attrs["run"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRunResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetRunResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"run": reflect.TypeOf(Run{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRunResponse
// only implements ToObjectValue() and Type().
func (m GetRunResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run": m.Run,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetRunResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run": Run{}.Type(ctx),
		},
	}
}

// GetRun returns the value of the Run field in GetRunResponse as
// a Run value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetRunResponse) GetRun(ctx context.Context) (Run, bool) {
	var e Run
	if m.Run.IsNull() || m.Run.IsUnknown() {
		return e, false
	}
	var v Run
	d := m.Run.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRun sets the value of the Run field in GetRunResponse.
func (m *GetRunResponse) SetRun(ctx context.Context, v Run) {
	vs := v.ToObjectValue(ctx)
	m.Run = vs
}

type HttpUrlSpec struct {
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

func (to *HttpUrlSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from HttpUrlSpec) {
}

func (to *HttpUrlSpec) SyncFieldsDuringRead(ctx context.Context, from HttpUrlSpec) {
}

func (m HttpUrlSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m HttpUrlSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, HttpUrlSpec
// only implements ToObjectValue() and Type().
func (m HttpUrlSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m HttpUrlSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"authorization":           types.StringType,
			"enable_ssl_verification": types.BoolType,
			"secret":                  types.StringType,
			"url":                     types.StringType,
		},
	}
}

type HttpUrlSpecWithoutSecret struct {
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

func (to *HttpUrlSpecWithoutSecret) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from HttpUrlSpecWithoutSecret) {
}

func (to *HttpUrlSpecWithoutSecret) SyncFieldsDuringRead(ctx context.Context, from HttpUrlSpecWithoutSecret) {
}

func (m HttpUrlSpecWithoutSecret) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m HttpUrlSpecWithoutSecret) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, HttpUrlSpecWithoutSecret
// only implements ToObjectValue() and Type().
func (m HttpUrlSpecWithoutSecret) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enable_ssl_verification": m.EnableSslVerification,
			"url":                     m.Url,
		})
}

// Type implements basetypes.ObjectValuable.
func (m HttpUrlSpecWithoutSecret) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enable_ssl_verification": types.BoolType,
			"url":                     types.StringType,
		},
	}
}

// Tag for a dataset input.
type InputTag struct {
	// The tag key.
	Key types.String `tfsdk:"key"`
	// The tag value.
	Value types.String `tfsdk:"value"`
}

func (to *InputTag) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from InputTag) {
}

func (to *InputTag) SyncFieldsDuringRead(ctx context.Context, from InputTag) {
}

func (m InputTag) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m InputTag) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InputTag
// only implements ToObjectValue() and Type().
func (m InputTag) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m InputTag) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

type JobSpec struct {
	// The personal access token used to authorize webhook's job runs.
	AccessToken types.String `tfsdk:"access_token"`
	// ID of the job that the webhook runs.
	JobId types.String `tfsdk:"job_id"`
	// URL of the workspace containing the job that this webhook runs. If not
	// specified, the job’s workspace URL is assumed to be the same as the
	// workspace where the webhook is created.
	WorkspaceUrl types.String `tfsdk:"workspace_url"`
}

func (to *JobSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from JobSpec) {
}

func (to *JobSpec) SyncFieldsDuringRead(ctx context.Context, from JobSpec) {
}

func (m JobSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m JobSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, JobSpec
// only implements ToObjectValue() and Type().
func (m JobSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_token":  m.AccessToken,
			"job_id":        m.JobId,
			"workspace_url": m.WorkspaceUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (m JobSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_token":  types.StringType,
			"job_id":        types.StringType,
			"workspace_url": types.StringType,
		},
	}
}

type JobSpecWithoutSecret struct {
	// ID of the job that the webhook runs.
	JobId types.String `tfsdk:"job_id"`
	// URL of the workspace containing the job that this webhook runs. If not
	// specified, the job’s workspace URL is assumed to be the same as the
	// workspace where the webhook is created.
	WorkspaceUrl types.String `tfsdk:"workspace_url"`
}

func (to *JobSpecWithoutSecret) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from JobSpecWithoutSecret) {
}

func (to *JobSpecWithoutSecret) SyncFieldsDuringRead(ctx context.Context, from JobSpecWithoutSecret) {
}

func (m JobSpecWithoutSecret) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m JobSpecWithoutSecret) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, JobSpecWithoutSecret
// only implements ToObjectValue() and Type().
func (m JobSpecWithoutSecret) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"job_id":        m.JobId,
			"workspace_url": m.WorkspaceUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (m JobSpecWithoutSecret) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"job_id":        types.StringType,
			"workspace_url": types.StringType,
		},
	}
}

// Feature for model version. ([ML-57150] Renamed from Feature to LinkedFeature)
type LinkedFeature struct {
	// Feature name
	FeatureName types.String `tfsdk:"feature_name"`
	// Feature table id
	FeatureTableId types.String `tfsdk:"feature_table_id"`
	// Feature table name
	FeatureTableName types.String `tfsdk:"feature_table_name"`
}

func (to *LinkedFeature) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LinkedFeature) {
}

func (to *LinkedFeature) SyncFieldsDuringRead(ctx context.Context, from LinkedFeature) {
}

func (m LinkedFeature) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m LinkedFeature) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LinkedFeature
// only implements ToObjectValue() and Type().
func (m LinkedFeature) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_name":       m.FeatureName,
			"feature_table_id":   m.FeatureTableId,
			"feature_table_name": m.FeatureTableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LinkedFeature) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"feature_name":       types.StringType,
			"feature_table_id":   types.StringType,
			"feature_table_name": types.StringType,
		},
	}
}

type ListArtifactsRequest struct {
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

func (to *ListArtifactsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListArtifactsRequest) {
}

func (to *ListArtifactsRequest) SyncFieldsDuringRead(ctx context.Context, from ListArtifactsRequest) {
}

func (m ListArtifactsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListArtifactsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListArtifactsRequest
// only implements ToObjectValue() and Type().
func (m ListArtifactsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ListArtifactsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_token": types.StringType,
			"path":       types.StringType,
			"run_id":     types.StringType,
			"run_uuid":   types.StringType,
		},
	}
}

type ListArtifactsResponse struct {
	// The file location and metadata for artifacts.
	Files types.List `tfsdk:"files"`
	// The token that can be used to retrieve the next page of artifact results.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// The root artifact directory for the run.
	RootUri types.String `tfsdk:"root_uri"`
}

func (to *ListArtifactsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListArtifactsResponse) {
	if !from.Files.IsNull() && !from.Files.IsUnknown() && to.Files.IsNull() && len(from.Files.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Files, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Files = from.Files
	}
}

func (to *ListArtifactsResponse) SyncFieldsDuringRead(ctx context.Context, from ListArtifactsResponse) {
	if !from.Files.IsNull() && !from.Files.IsUnknown() && to.Files.IsNull() && len(from.Files.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Files, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Files = from.Files
	}
}

func (m ListArtifactsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListArtifactsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"files": reflect.TypeOf(FileInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListArtifactsResponse
// only implements ToObjectValue() and Type().
func (m ListArtifactsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"files":           m.Files,
			"next_page_token": m.NextPageToken,
			"root_uri":        m.RootUri,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListArtifactsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"files": basetypes.ListType{
				ElemType: FileInfo{}.Type(ctx),
			},
			"next_page_token": types.StringType,
			"root_uri":        types.StringType,
		},
	}
}

// GetFiles returns the value of the Files field in ListArtifactsResponse as
// a slice of FileInfo values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListArtifactsResponse) GetFiles(ctx context.Context) ([]FileInfo, bool) {
	if m.Files.IsNull() || m.Files.IsUnknown() {
		return nil, false
	}
	var v []FileInfo
	d := m.Files.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFiles sets the value of the Files field in ListArtifactsResponse.
func (m *ListArtifactsResponse) SetFiles(ctx context.Context, v []FileInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["files"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Files = types.ListValueMust(t, vs)
}

type ListExperimentsRequest struct {
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

func (to *ListExperimentsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListExperimentsRequest) {
}

func (to *ListExperimentsRequest) SyncFieldsDuringRead(ctx context.Context, from ListExperimentsRequest) {
}

func (m ListExperimentsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListExperimentsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExperimentsRequest
// only implements ToObjectValue() and Type().
func (m ListExperimentsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results": m.MaxResults,
			"page_token":  m.PageToken,
			"view_type":   m.ViewType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListExperimentsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_results": types.Int64Type,
			"page_token":  types.StringType,
			"view_type":   types.StringType,
		},
	}
}

type ListExperimentsResponse struct {
	// Paginated Experiments beginning with the first item on the requested
	// page.
	Experiments types.List `tfsdk:"experiments"`
	// Token that can be used to retrieve the next page of experiments. Empty
	// token means no more experiment is available for retrieval.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListExperimentsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListExperimentsResponse) {
	if !from.Experiments.IsNull() && !from.Experiments.IsUnknown() && to.Experiments.IsNull() && len(from.Experiments.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Experiments, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Experiments = from.Experiments
	}
}

func (to *ListExperimentsResponse) SyncFieldsDuringRead(ctx context.Context, from ListExperimentsResponse) {
	if !from.Experiments.IsNull() && !from.Experiments.IsUnknown() && to.Experiments.IsNull() && len(from.Experiments.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Experiments, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Experiments = from.Experiments
	}
}

func (m ListExperimentsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListExperimentsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"experiments": reflect.TypeOf(Experiment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExperimentsResponse
// only implements ToObjectValue() and Type().
func (m ListExperimentsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiments":     m.Experiments,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListExperimentsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiments": basetypes.ListType{
				ElemType: Experiment{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetExperiments returns the value of the Experiments field in ListExperimentsResponse as
// a slice of Experiment values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListExperimentsResponse) GetExperiments(ctx context.Context) ([]Experiment, bool) {
	if m.Experiments.IsNull() || m.Experiments.IsUnknown() {
		return nil, false
	}
	var v []Experiment
	d := m.Experiments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExperiments sets the value of the Experiments field in ListExperimentsResponse.
func (m *ListExperimentsResponse) SetExperiments(ctx context.Context, v []Experiment) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["experiments"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Experiments = types.ListValueMust(t, vs)
}

type ListFeatureTagsRequest struct {
	FeatureName types.String `tfsdk:"-"`
	// The maximum number of results to return.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page based on a previous query.
	PageToken types.String `tfsdk:"-"`

	TableName types.String `tfsdk:"-"`
}

func (to *ListFeatureTagsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListFeatureTagsRequest) {
}

func (to *ListFeatureTagsRequest) SyncFieldsDuringRead(ctx context.Context, from ListFeatureTagsRequest) {
}

func (m ListFeatureTagsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListFeatureTagsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFeatureTagsRequest
// only implements ToObjectValue() and Type().
func (m ListFeatureTagsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ListFeatureTagsRequest) Type(ctx context.Context) attr.Type {
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
type ListFeatureTagsResponse struct {
	FeatureTags types.List `tfsdk:"feature_tags"`
	// Pagination token to request the next page of results for this query.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListFeatureTagsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListFeatureTagsResponse) {
	if !from.FeatureTags.IsNull() && !from.FeatureTags.IsUnknown() && to.FeatureTags.IsNull() && len(from.FeatureTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FeatureTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FeatureTags = from.FeatureTags
	}
}

func (to *ListFeatureTagsResponse) SyncFieldsDuringRead(ctx context.Context, from ListFeatureTagsResponse) {
	if !from.FeatureTags.IsNull() && !from.FeatureTags.IsUnknown() && to.FeatureTags.IsNull() && len(from.FeatureTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FeatureTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FeatureTags = from.FeatureTags
	}
}

func (m ListFeatureTagsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListFeatureTagsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"feature_tags": reflect.TypeOf(FeatureTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFeatureTagsResponse
// only implements ToObjectValue() and Type().
func (m ListFeatureTagsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_tags":    m.FeatureTags,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListFeatureTagsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"feature_tags": basetypes.ListType{
				ElemType: FeatureTag{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetFeatureTags returns the value of the FeatureTags field in ListFeatureTagsResponse as
// a slice of FeatureTag values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListFeatureTagsResponse) GetFeatureTags(ctx context.Context) ([]FeatureTag, bool) {
	if m.FeatureTags.IsNull() || m.FeatureTags.IsUnknown() {
		return nil, false
	}
	var v []FeatureTag
	d := m.FeatureTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFeatureTags sets the value of the FeatureTags field in ListFeatureTagsResponse.
func (m *ListFeatureTagsResponse) SetFeatureTags(ctx context.Context, v []FeatureTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["feature_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.FeatureTags = types.ListValueMust(t, vs)
}

type ListFeaturesRequest struct {
	// The maximum number of results to return.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page based on a previous query.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListFeaturesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListFeaturesRequest) {
}

func (to *ListFeaturesRequest) SyncFieldsDuringRead(ctx context.Context, from ListFeaturesRequest) {
}

func (m ListFeaturesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListFeaturesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFeaturesRequest
// only implements ToObjectValue() and Type().
func (m ListFeaturesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListFeaturesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListFeaturesResponse struct {
	// List of features.
	Features types.List `tfsdk:"features"`
	// Pagination token to request the next page of results for this query.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListFeaturesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListFeaturesResponse) {
	if !from.Features.IsNull() && !from.Features.IsUnknown() && to.Features.IsNull() && len(from.Features.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Features, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Features = from.Features
	}
}

func (to *ListFeaturesResponse) SyncFieldsDuringRead(ctx context.Context, from ListFeaturesResponse) {
	if !from.Features.IsNull() && !from.Features.IsUnknown() && to.Features.IsNull() && len(from.Features.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Features, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Features = from.Features
	}
}

func (m ListFeaturesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListFeaturesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"features": reflect.TypeOf(Feature{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFeaturesResponse
// only implements ToObjectValue() and Type().
func (m ListFeaturesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"features":        m.Features,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListFeaturesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"features": basetypes.ListType{
				ElemType: Feature{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetFeatures returns the value of the Features field in ListFeaturesResponse as
// a slice of Feature values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListFeaturesResponse) GetFeatures(ctx context.Context) ([]Feature, bool) {
	if m.Features.IsNull() || m.Features.IsUnknown() {
		return nil, false
	}
	var v []Feature
	d := m.Features.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFeatures sets the value of the Features field in ListFeaturesResponse.
func (m *ListFeaturesResponse) SetFeatures(ctx context.Context, v []Feature) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["features"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Features = types.ListValueMust(t, vs)
}

type ListModelsRequest struct {
	// Maximum number of registered models desired. Max threshold is 1000.
	MaxResults types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page based on a previous query.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListModelsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListModelsRequest) {
}

func (to *ListModelsRequest) SyncFieldsDuringRead(ctx context.Context, from ListModelsRequest) {
}

func (m ListModelsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListModelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListModelsRequest
// only implements ToObjectValue() and Type().
func (m ListModelsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results": m.MaxResults,
			"page_token":  m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListModelsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"max_results": types.Int64Type,
			"page_token":  types.StringType,
		},
	}
}

type ListModelsResponse struct {
	// Pagination token to request next page of models for the same query.
	NextPageToken types.String `tfsdk:"next_page_token"`

	RegisteredModels types.List `tfsdk:"registered_models"`
}

func (to *ListModelsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListModelsResponse) {
	if !from.RegisteredModels.IsNull() && !from.RegisteredModels.IsUnknown() && to.RegisteredModels.IsNull() && len(from.RegisteredModels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RegisteredModels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RegisteredModels = from.RegisteredModels
	}
}

func (to *ListModelsResponse) SyncFieldsDuringRead(ctx context.Context, from ListModelsResponse) {
	if !from.RegisteredModels.IsNull() && !from.RegisteredModels.IsUnknown() && to.RegisteredModels.IsNull() && len(from.RegisteredModels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RegisteredModels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RegisteredModels = from.RegisteredModels
	}
}

func (m ListModelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListModelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"registered_models": reflect.TypeOf(Model{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListModelsResponse
// only implements ToObjectValue() and Type().
func (m ListModelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":   m.NextPageToken,
			"registered_models": m.RegisteredModels,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListModelsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"registered_models": basetypes.ListType{
				ElemType: Model{}.Type(ctx),
			},
		},
	}
}

// GetRegisteredModels returns the value of the RegisteredModels field in ListModelsResponse as
// a slice of Model values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListModelsResponse) GetRegisteredModels(ctx context.Context) ([]Model, bool) {
	if m.RegisteredModels.IsNull() || m.RegisteredModels.IsUnknown() {
		return nil, false
	}
	var v []Model
	d := m.RegisteredModels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRegisteredModels sets the value of the RegisteredModels field in ListModelsResponse.
func (m *ListModelsResponse) SetRegisteredModels(ctx context.Context, v []Model) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["registered_models"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.RegisteredModels = types.ListValueMust(t, vs)
}

type ListOnlineStoresRequest struct {
	// The maximum number of results to return. Defaults to 100 if not
	// specified.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page based on a previous query.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListOnlineStoresRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListOnlineStoresRequest) {
}

func (to *ListOnlineStoresRequest) SyncFieldsDuringRead(ctx context.Context, from ListOnlineStoresRequest) {
}

func (m ListOnlineStoresRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListOnlineStoresRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListOnlineStoresRequest
// only implements ToObjectValue() and Type().
func (m ListOnlineStoresRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListOnlineStoresRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListOnlineStoresResponse struct {
	// Pagination token to request the next page of results for this query.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// List of online stores.
	OnlineStores types.List `tfsdk:"online_stores"`
}

func (to *ListOnlineStoresResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListOnlineStoresResponse) {
	if !from.OnlineStores.IsNull() && !from.OnlineStores.IsUnknown() && to.OnlineStores.IsNull() && len(from.OnlineStores.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OnlineStores, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OnlineStores = from.OnlineStores
	}
}

func (to *ListOnlineStoresResponse) SyncFieldsDuringRead(ctx context.Context, from ListOnlineStoresResponse) {
	if !from.OnlineStores.IsNull() && !from.OnlineStores.IsUnknown() && to.OnlineStores.IsNull() && len(from.OnlineStores.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OnlineStores, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OnlineStores = from.OnlineStores
	}
}

func (m ListOnlineStoresResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListOnlineStoresResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"online_stores": reflect.TypeOf(OnlineStore{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListOnlineStoresResponse
// only implements ToObjectValue() and Type().
func (m ListOnlineStoresResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"online_stores":   m.OnlineStores,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListOnlineStoresResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"online_stores": basetypes.ListType{
				ElemType: OnlineStore{}.Type(ctx),
			},
		},
	}
}

// GetOnlineStores returns the value of the OnlineStores field in ListOnlineStoresResponse as
// a slice of OnlineStore values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListOnlineStoresResponse) GetOnlineStores(ctx context.Context) ([]OnlineStore, bool) {
	if m.OnlineStores.IsNull() || m.OnlineStores.IsUnknown() {
		return nil, false
	}
	var v []OnlineStore
	d := m.OnlineStores.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOnlineStores sets the value of the OnlineStores field in ListOnlineStoresResponse.
func (m *ListOnlineStoresResponse) SetOnlineStores(ctx context.Context, v []OnlineStore) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["online_stores"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.OnlineStores = types.ListValueMust(t, vs)
}

type ListRegistryWebhooks struct {
	// Token that can be used to retrieve the next page of artifact results
	NextPageToken types.String `tfsdk:"next_page_token"`
	// Array of registry webhooks.
	Webhooks types.List `tfsdk:"webhooks"`
}

func (to *ListRegistryWebhooks) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListRegistryWebhooks) {
	if !from.Webhooks.IsNull() && !from.Webhooks.IsUnknown() && to.Webhooks.IsNull() && len(from.Webhooks.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Webhooks, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Webhooks = from.Webhooks
	}
}

func (to *ListRegistryWebhooks) SyncFieldsDuringRead(ctx context.Context, from ListRegistryWebhooks) {
	if !from.Webhooks.IsNull() && !from.Webhooks.IsUnknown() && to.Webhooks.IsNull() && len(from.Webhooks.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Webhooks, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Webhooks = from.Webhooks
	}
}

func (m ListRegistryWebhooks) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListRegistryWebhooks) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"webhooks": reflect.TypeOf(RegistryWebhook{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListRegistryWebhooks
// only implements ToObjectValue() and Type().
func (m ListRegistryWebhooks) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"webhooks":        m.Webhooks,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListRegistryWebhooks) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"webhooks": basetypes.ListType{
				ElemType: RegistryWebhook{}.Type(ctx),
			},
		},
	}
}

// GetWebhooks returns the value of the Webhooks field in ListRegistryWebhooks as
// a slice of RegistryWebhook values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListRegistryWebhooks) GetWebhooks(ctx context.Context) ([]RegistryWebhook, bool) {
	if m.Webhooks.IsNull() || m.Webhooks.IsUnknown() {
		return nil, false
	}
	var v []RegistryWebhook
	d := m.Webhooks.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWebhooks sets the value of the Webhooks field in ListRegistryWebhooks.
func (m *ListRegistryWebhooks) SetWebhooks(ctx context.Context, v []RegistryWebhook) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["webhooks"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Webhooks = types.ListValueMust(t, vs)
}

type ListTransitionRequestsRequest struct {
	// Name of the registered model.
	Name types.String `tfsdk:"-"`
	// Version of the model.
	Version types.String `tfsdk:"-"`
}

func (to *ListTransitionRequestsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListTransitionRequestsRequest) {
}

func (to *ListTransitionRequestsRequest) SyncFieldsDuringRead(ctx context.Context, from ListTransitionRequestsRequest) {
}

func (m ListTransitionRequestsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListTransitionRequestsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTransitionRequestsRequest
// only implements ToObjectValue() and Type().
func (m ListTransitionRequestsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":    m.Name,
			"version": m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListTransitionRequestsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":    types.StringType,
			"version": types.StringType,
		},
	}
}

type ListTransitionRequestsResponse struct {
	// Array of open transition requests.
	Requests types.List `tfsdk:"requests"`
}

func (to *ListTransitionRequestsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListTransitionRequestsResponse) {
	if !from.Requests.IsNull() && !from.Requests.IsUnknown() && to.Requests.IsNull() && len(from.Requests.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Requests, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Requests = from.Requests
	}
}

func (to *ListTransitionRequestsResponse) SyncFieldsDuringRead(ctx context.Context, from ListTransitionRequestsResponse) {
	if !from.Requests.IsNull() && !from.Requests.IsUnknown() && to.Requests.IsNull() && len(from.Requests.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Requests, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Requests = from.Requests
	}
}

func (m ListTransitionRequestsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListTransitionRequestsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"requests": reflect.TypeOf(Activity{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTransitionRequestsResponse
// only implements ToObjectValue() and Type().
func (m ListTransitionRequestsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"requests": m.Requests,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListTransitionRequestsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"requests": basetypes.ListType{
				ElemType: Activity{}.Type(ctx),
			},
		},
	}
}

// GetRequests returns the value of the Requests field in ListTransitionRequestsResponse as
// a slice of Activity values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListTransitionRequestsResponse) GetRequests(ctx context.Context) ([]Activity, bool) {
	if m.Requests.IsNull() || m.Requests.IsUnknown() {
		return nil, false
	}
	var v []Activity
	d := m.Requests.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRequests sets the value of the Requests field in ListTransitionRequestsResponse.
func (m *ListTransitionRequestsResponse) SetRequests(ctx context.Context, v []Activity) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["requests"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Requests = types.ListValueMust(t, vs)
}

type ListWebhooksRequest struct {
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

func (to *ListWebhooksRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListWebhooksRequest) {
	if !from.Events.IsNull() && !from.Events.IsUnknown() && to.Events.IsNull() && len(from.Events.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Events, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Events = from.Events
	}
}

func (to *ListWebhooksRequest) SyncFieldsDuringRead(ctx context.Context, from ListWebhooksRequest) {
	if !from.Events.IsNull() && !from.Events.IsUnknown() && to.Events.IsNull() && len(from.Events.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Events, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Events = from.Events
	}
}

func (m ListWebhooksRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListWebhooksRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"events": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWebhooksRequest
// only implements ToObjectValue() and Type().
func (m ListWebhooksRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ListWebhooksRequest) Type(ctx context.Context) attr.Type {
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

// GetEvents returns the value of the Events field in ListWebhooksRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListWebhooksRequest) GetEvents(ctx context.Context) ([]types.String, bool) {
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

// SetEvents sets the value of the Events field in ListWebhooksRequest.
func (m *ListWebhooksRequest) SetEvents(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["events"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Events = types.ListValueMust(t, vs)
}

type LogBatch struct {
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

func (to *LogBatch) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogBatch) {
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

func (to *LogBatch) SyncFieldsDuringRead(ctx context.Context, from LogBatch) {
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

func (m LogBatch) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m LogBatch) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metrics": reflect.TypeOf(Metric{}),
		"params":  reflect.TypeOf(Param{}),
		"tags":    reflect.TypeOf(RunTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogBatch
// only implements ToObjectValue() and Type().
func (m LogBatch) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m LogBatch) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metrics": basetypes.ListType{
				ElemType: Metric{}.Type(ctx),
			},
			"params": basetypes.ListType{
				ElemType: Param{}.Type(ctx),
			},
			"run_id": types.StringType,
			"tags": basetypes.ListType{
				ElemType: RunTag{}.Type(ctx),
			},
		},
	}
}

// GetMetrics returns the value of the Metrics field in LogBatch as
// a slice of Metric values.
// If the field is unknown or null, the boolean return value is false.
func (m *LogBatch) GetMetrics(ctx context.Context) ([]Metric, bool) {
	if m.Metrics.IsNull() || m.Metrics.IsUnknown() {
		return nil, false
	}
	var v []Metric
	d := m.Metrics.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMetrics sets the value of the Metrics field in LogBatch.
func (m *LogBatch) SetMetrics(ctx context.Context, v []Metric) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["metrics"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Metrics = types.ListValueMust(t, vs)
}

// GetParams returns the value of the Params field in LogBatch as
// a slice of Param values.
// If the field is unknown or null, the boolean return value is false.
func (m *LogBatch) GetParams(ctx context.Context) ([]Param, bool) {
	if m.Params.IsNull() || m.Params.IsUnknown() {
		return nil, false
	}
	var v []Param
	d := m.Params.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParams sets the value of the Params field in LogBatch.
func (m *LogBatch) SetParams(ctx context.Context, v []Param) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["params"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Params = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in LogBatch as
// a slice of RunTag values.
// If the field is unknown or null, the boolean return value is false.
func (m *LogBatch) GetTags(ctx context.Context) ([]RunTag, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []RunTag
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in LogBatch.
func (m *LogBatch) SetTags(ctx context.Context, v []RunTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type LogBatchResponse struct {
}

func (to *LogBatchResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogBatchResponse) {
}

func (to *LogBatchResponse) SyncFieldsDuringRead(ctx context.Context, from LogBatchResponse) {
}

func (m LogBatchResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogBatchResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LogBatchResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogBatchResponse
// only implements ToObjectValue() and Type().
func (m LogBatchResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m LogBatchResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type LogInputs struct {
	// Dataset inputs
	Datasets types.List `tfsdk:"datasets"`
	// Model inputs
	Models types.List `tfsdk:"models"`
	// ID of the run to log under
	RunId types.String `tfsdk:"run_id"`
}

func (to *LogInputs) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogInputs) {
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

func (to *LogInputs) SyncFieldsDuringRead(ctx context.Context, from LogInputs) {
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

func (m LogInputs) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m LogInputs) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"datasets": reflect.TypeOf(DatasetInput{}),
		"models":   reflect.TypeOf(ModelInput{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogInputs
// only implements ToObjectValue() and Type().
func (m LogInputs) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"datasets": m.Datasets,
			"models":   m.Models,
			"run_id":   m.RunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LogInputs) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"datasets": basetypes.ListType{
				ElemType: DatasetInput{}.Type(ctx),
			},
			"models": basetypes.ListType{
				ElemType: ModelInput{}.Type(ctx),
			},
			"run_id": types.StringType,
		},
	}
}

// GetDatasets returns the value of the Datasets field in LogInputs as
// a slice of DatasetInput values.
// If the field is unknown or null, the boolean return value is false.
func (m *LogInputs) GetDatasets(ctx context.Context) ([]DatasetInput, bool) {
	if m.Datasets.IsNull() || m.Datasets.IsUnknown() {
		return nil, false
	}
	var v []DatasetInput
	d := m.Datasets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatasets sets the value of the Datasets field in LogInputs.
func (m *LogInputs) SetDatasets(ctx context.Context, v []DatasetInput) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["datasets"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Datasets = types.ListValueMust(t, vs)
}

// GetModels returns the value of the Models field in LogInputs as
// a slice of ModelInput values.
// If the field is unknown or null, the boolean return value is false.
func (m *LogInputs) GetModels(ctx context.Context) ([]ModelInput, bool) {
	if m.Models.IsNull() || m.Models.IsUnknown() {
		return nil, false
	}
	var v []ModelInput
	d := m.Models.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModels sets the value of the Models field in LogInputs.
func (m *LogInputs) SetModels(ctx context.Context, v []ModelInput) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["models"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Models = types.ListValueMust(t, vs)
}

type LogInputsResponse struct {
}

func (to *LogInputsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogInputsResponse) {
}

func (to *LogInputsResponse) SyncFieldsDuringRead(ctx context.Context, from LogInputsResponse) {
}

func (m LogInputsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogInputsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LogInputsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogInputsResponse
// only implements ToObjectValue() and Type().
func (m LogInputsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m LogInputsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type LogLoggedModelParamsRequest struct {
	// The ID of the logged model to log params for.
	ModelId types.String `tfsdk:"-"`
	// Parameters to attach to the model.
	Params types.List `tfsdk:"params"`
}

func (to *LogLoggedModelParamsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogLoggedModelParamsRequest) {
	if !from.Params.IsNull() && !from.Params.IsUnknown() && to.Params.IsNull() && len(from.Params.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Params, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Params = from.Params
	}
}

func (to *LogLoggedModelParamsRequest) SyncFieldsDuringRead(ctx context.Context, from LogLoggedModelParamsRequest) {
	if !from.Params.IsNull() && !from.Params.IsUnknown() && to.Params.IsNull() && len(from.Params.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Params, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Params = from.Params
	}
}

func (m LogLoggedModelParamsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m LogLoggedModelParamsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"params": reflect.TypeOf(LoggedModelParameter{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogLoggedModelParamsRequest
// only implements ToObjectValue() and Type().
func (m LogLoggedModelParamsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_id": m.ModelId,
			"params":   m.Params,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LogLoggedModelParamsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_id": types.StringType,
			"params": basetypes.ListType{
				ElemType: LoggedModelParameter{}.Type(ctx),
			},
		},
	}
}

// GetParams returns the value of the Params field in LogLoggedModelParamsRequest as
// a slice of LoggedModelParameter values.
// If the field is unknown or null, the boolean return value is false.
func (m *LogLoggedModelParamsRequest) GetParams(ctx context.Context) ([]LoggedModelParameter, bool) {
	if m.Params.IsNull() || m.Params.IsUnknown() {
		return nil, false
	}
	var v []LoggedModelParameter
	d := m.Params.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParams sets the value of the Params field in LogLoggedModelParamsRequest.
func (m *LogLoggedModelParamsRequest) SetParams(ctx context.Context, v []LoggedModelParameter) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["params"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Params = types.ListValueMust(t, vs)
}

type LogLoggedModelParamsRequestResponse struct {
}

func (to *LogLoggedModelParamsRequestResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogLoggedModelParamsRequestResponse) {
}

func (to *LogLoggedModelParamsRequestResponse) SyncFieldsDuringRead(ctx context.Context, from LogLoggedModelParamsRequestResponse) {
}

func (m LogLoggedModelParamsRequestResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogLoggedModelParamsRequestResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LogLoggedModelParamsRequestResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogLoggedModelParamsRequestResponse
// only implements ToObjectValue() and Type().
func (m LogLoggedModelParamsRequestResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m LogLoggedModelParamsRequestResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type LogMetric struct {
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

func (to *LogMetric) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogMetric) {
}

func (to *LogMetric) SyncFieldsDuringRead(ctx context.Context, from LogMetric) {
}

func (m LogMetric) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m LogMetric) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogMetric
// only implements ToObjectValue() and Type().
func (m LogMetric) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m LogMetric) Type(ctx context.Context) attr.Type {
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

type LogMetricResponse struct {
}

func (to *LogMetricResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogMetricResponse) {
}

func (to *LogMetricResponse) SyncFieldsDuringRead(ctx context.Context, from LogMetricResponse) {
}

func (m LogMetricResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogMetricResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LogMetricResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogMetricResponse
// only implements ToObjectValue() and Type().
func (m LogMetricResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m LogMetricResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type LogModel struct {
	// MLmodel file in json format.
	ModelJson types.String `tfsdk:"model_json"`
	// ID of the run to log under
	RunId types.String `tfsdk:"run_id"`
}

func (to *LogModel) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogModel) {
}

func (to *LogModel) SyncFieldsDuringRead(ctx context.Context, from LogModel) {
}

func (m LogModel) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m LogModel) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogModel
// only implements ToObjectValue() and Type().
func (m LogModel) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_json": m.ModelJson,
			"run_id":     m.RunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LogModel) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_json": types.StringType,
			"run_id":     types.StringType,
		},
	}
}

type LogModelResponse struct {
}

func (to *LogModelResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogModelResponse) {
}

func (to *LogModelResponse) SyncFieldsDuringRead(ctx context.Context, from LogModelResponse) {
}

func (m LogModelResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogModelResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LogModelResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogModelResponse
// only implements ToObjectValue() and Type().
func (m LogModelResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m LogModelResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type LogOutputsRequest struct {
	// The model outputs from the Run.
	Models types.List `tfsdk:"models"`
	// The ID of the Run from which to log outputs.
	RunId types.String `tfsdk:"run_id"`
}

func (to *LogOutputsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogOutputsRequest) {
	if !from.Models.IsNull() && !from.Models.IsUnknown() && to.Models.IsNull() && len(from.Models.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Models, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Models = from.Models
	}
}

func (to *LogOutputsRequest) SyncFieldsDuringRead(ctx context.Context, from LogOutputsRequest) {
	if !from.Models.IsNull() && !from.Models.IsUnknown() && to.Models.IsNull() && len(from.Models.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Models, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Models = from.Models
	}
}

func (m LogOutputsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m LogOutputsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"models": reflect.TypeOf(ModelOutput{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogOutputsRequest
// only implements ToObjectValue() and Type().
func (m LogOutputsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"models": m.Models,
			"run_id": m.RunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LogOutputsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"models": basetypes.ListType{
				ElemType: ModelOutput{}.Type(ctx),
			},
			"run_id": types.StringType,
		},
	}
}

// GetModels returns the value of the Models field in LogOutputsRequest as
// a slice of ModelOutput values.
// If the field is unknown or null, the boolean return value is false.
func (m *LogOutputsRequest) GetModels(ctx context.Context) ([]ModelOutput, bool) {
	if m.Models.IsNull() || m.Models.IsUnknown() {
		return nil, false
	}
	var v []ModelOutput
	d := m.Models.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModels sets the value of the Models field in LogOutputsRequest.
func (m *LogOutputsRequest) SetModels(ctx context.Context, v []ModelOutput) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["models"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Models = types.ListValueMust(t, vs)
}

type LogOutputsResponse struct {
}

func (to *LogOutputsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogOutputsResponse) {
}

func (to *LogOutputsResponse) SyncFieldsDuringRead(ctx context.Context, from LogOutputsResponse) {
}

func (m LogOutputsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogOutputsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LogOutputsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogOutputsResponse
// only implements ToObjectValue() and Type().
func (m LogOutputsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m LogOutputsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type LogParam struct {
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

func (to *LogParam) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogParam) {
}

func (to *LogParam) SyncFieldsDuringRead(ctx context.Context, from LogParam) {
}

func (m LogParam) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m LogParam) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogParam
// only implements ToObjectValue() and Type().
func (m LogParam) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m LogParam) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":      types.StringType,
			"run_id":   types.StringType,
			"run_uuid": types.StringType,
			"value":    types.StringType,
		},
	}
}

type LogParamResponse struct {
}

func (to *LogParamResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogParamResponse) {
}

func (to *LogParamResponse) SyncFieldsDuringRead(ctx context.Context, from LogParamResponse) {
}

func (m LogParamResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogParamResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LogParamResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogParamResponse
// only implements ToObjectValue() and Type().
func (m LogParamResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m LogParamResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// A logged model message includes logged model attributes, tags, registration
// info, params, and linked run metrics.
type LoggedModel struct {
	// The params and metrics attached to the logged model.
	Data types.Object `tfsdk:"data"`
	// The logged model attributes such as model ID, status, tags, etc.
	Info types.Object `tfsdk:"info"`
}

func (to *LoggedModel) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LoggedModel) {
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

func (to *LoggedModel) SyncFieldsDuringRead(ctx context.Context, from LoggedModel) {
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

func (m LoggedModel) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["data"] = attrs["data"].SetOptional()
	attrs["info"] = attrs["info"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LoggedModel.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LoggedModel) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data": reflect.TypeOf(LoggedModelData{}),
		"info": reflect.TypeOf(LoggedModelInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LoggedModel
// only implements ToObjectValue() and Type().
func (m LoggedModel) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data": m.Data,
			"info": m.Info,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LoggedModel) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"data": LoggedModelData{}.Type(ctx),
			"info": LoggedModelInfo{}.Type(ctx),
		},
	}
}

// GetData returns the value of the Data field in LoggedModel as
// a LoggedModelData value.
// If the field is unknown or null, the boolean return value is false.
func (m *LoggedModel) GetData(ctx context.Context) (LoggedModelData, bool) {
	var e LoggedModelData
	if m.Data.IsNull() || m.Data.IsUnknown() {
		return e, false
	}
	var v LoggedModelData
	d := m.Data.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetData sets the value of the Data field in LoggedModel.
func (m *LoggedModel) SetData(ctx context.Context, v LoggedModelData) {
	vs := v.ToObjectValue(ctx)
	m.Data = vs
}

// GetInfo returns the value of the Info field in LoggedModel as
// a LoggedModelInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *LoggedModel) GetInfo(ctx context.Context) (LoggedModelInfo, bool) {
	var e LoggedModelInfo
	if m.Info.IsNull() || m.Info.IsUnknown() {
		return e, false
	}
	var v LoggedModelInfo
	d := m.Info.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInfo sets the value of the Info field in LoggedModel.
func (m *LoggedModel) SetInfo(ctx context.Context, v LoggedModelInfo) {
	vs := v.ToObjectValue(ctx)
	m.Info = vs
}

// A LoggedModelData message includes logged model params and linked metrics.
type LoggedModelData struct {
	// Performance metrics linked to the model.
	Metrics types.List `tfsdk:"metrics"`
	// Immutable string key-value pairs of the model.
	Params types.List `tfsdk:"params"`
}

func (to *LoggedModelData) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LoggedModelData) {
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

func (to *LoggedModelData) SyncFieldsDuringRead(ctx context.Context, from LoggedModelData) {
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

func (m LoggedModelData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m LoggedModelData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metrics": reflect.TypeOf(Metric{}),
		"params":  reflect.TypeOf(LoggedModelParameter{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LoggedModelData
// only implements ToObjectValue() and Type().
func (m LoggedModelData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metrics": m.Metrics,
			"params":  m.Params,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LoggedModelData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metrics": basetypes.ListType{
				ElemType: Metric{}.Type(ctx),
			},
			"params": basetypes.ListType{
				ElemType: LoggedModelParameter{}.Type(ctx),
			},
		},
	}
}

// GetMetrics returns the value of the Metrics field in LoggedModelData as
// a slice of Metric values.
// If the field is unknown or null, the boolean return value is false.
func (m *LoggedModelData) GetMetrics(ctx context.Context) ([]Metric, bool) {
	if m.Metrics.IsNull() || m.Metrics.IsUnknown() {
		return nil, false
	}
	var v []Metric
	d := m.Metrics.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMetrics sets the value of the Metrics field in LoggedModelData.
func (m *LoggedModelData) SetMetrics(ctx context.Context, v []Metric) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["metrics"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Metrics = types.ListValueMust(t, vs)
}

// GetParams returns the value of the Params field in LoggedModelData as
// a slice of LoggedModelParameter values.
// If the field is unknown or null, the boolean return value is false.
func (m *LoggedModelData) GetParams(ctx context.Context) ([]LoggedModelParameter, bool) {
	if m.Params.IsNull() || m.Params.IsUnknown() {
		return nil, false
	}
	var v []LoggedModelParameter
	d := m.Params.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParams sets the value of the Params field in LoggedModelData.
func (m *LoggedModelData) SetParams(ctx context.Context, v []LoggedModelParameter) {
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
type LoggedModelInfo struct {
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

func (to *LoggedModelInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LoggedModelInfo) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *LoggedModelInfo) SyncFieldsDuringRead(ctx context.Context, from LoggedModelInfo) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m LoggedModelInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m LoggedModelInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(LoggedModelTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LoggedModelInfo
// only implements ToObjectValue() and Type().
func (m LoggedModelInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m LoggedModelInfo) Type(ctx context.Context) attr.Type {
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
				ElemType: LoggedModelTag{}.Type(ctx),
			},
		},
	}
}

// GetTags returns the value of the Tags field in LoggedModelInfo as
// a slice of LoggedModelTag values.
// If the field is unknown or null, the boolean return value is false.
func (m *LoggedModelInfo) GetTags(ctx context.Context) ([]LoggedModelTag, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []LoggedModelTag
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in LoggedModelInfo.
func (m *LoggedModelInfo) SetTags(ctx context.Context, v []LoggedModelTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

// Parameter associated with a LoggedModel.
type LoggedModelParameter struct {
	// The key identifying this param.
	Key types.String `tfsdk:"key"`
	// The value of this param.
	Value types.String `tfsdk:"value"`
}

func (to *LoggedModelParameter) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LoggedModelParameter) {
}

func (to *LoggedModelParameter) SyncFieldsDuringRead(ctx context.Context, from LoggedModelParameter) {
}

func (m LoggedModelParameter) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m LoggedModelParameter) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LoggedModelParameter
// only implements ToObjectValue() and Type().
func (m LoggedModelParameter) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LoggedModelParameter) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

// Tag for a LoggedModel.
type LoggedModelTag struct {
	// The tag key.
	Key types.String `tfsdk:"key"`
	// The tag value.
	Value types.String `tfsdk:"value"`
}

func (to *LoggedModelTag) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LoggedModelTag) {
}

func (to *LoggedModelTag) SyncFieldsDuringRead(ctx context.Context, from LoggedModelTag) {
}

func (m LoggedModelTag) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m LoggedModelTag) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LoggedModelTag
// only implements ToObjectValue() and Type().
func (m LoggedModelTag) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LoggedModelTag) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

// Metric associated with a run, represented as a key-value pair.
type Metric struct {
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

func (to *Metric) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Metric) {
}

func (to *Metric) SyncFieldsDuringRead(ctx context.Context, from Metric) {
}

func (m Metric) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Metric) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Metric
// only implements ToObjectValue() and Type().
func (m Metric) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m Metric) Type(ctx context.Context) attr.Type {
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

type Model struct {
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

func (to *Model) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Model) {
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

func (to *Model) SyncFieldsDuringRead(ctx context.Context, from Model) {
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

func (m Model) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Model) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"latest_versions": reflect.TypeOf(ModelVersion{}),
		"tags":            reflect.TypeOf(ModelTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Model
// only implements ToObjectValue() and Type().
func (m Model) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m Model) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creation_timestamp":     types.Int64Type,
			"description":            types.StringType,
			"last_updated_timestamp": types.Int64Type,
			"latest_versions": basetypes.ListType{
				ElemType: ModelVersion{}.Type(ctx),
			},
			"name": types.StringType,
			"tags": basetypes.ListType{
				ElemType: ModelTag{}.Type(ctx),
			},
			"user_id": types.StringType,
		},
	}
}

// GetLatestVersions returns the value of the LatestVersions field in Model as
// a slice of ModelVersion values.
// If the field is unknown or null, the boolean return value is false.
func (m *Model) GetLatestVersions(ctx context.Context) ([]ModelVersion, bool) {
	if m.LatestVersions.IsNull() || m.LatestVersions.IsUnknown() {
		return nil, false
	}
	var v []ModelVersion
	d := m.LatestVersions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLatestVersions sets the value of the LatestVersions field in Model.
func (m *Model) SetLatestVersions(ctx context.Context, v []ModelVersion) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["latest_versions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.LatestVersions = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in Model as
// a slice of ModelTag values.
// If the field is unknown or null, the boolean return value is false.
func (m *Model) GetTags(ctx context.Context) ([]ModelTag, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []ModelTag
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in Model.
func (m *Model) SetTags(ctx context.Context, v []ModelTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type ModelDatabricks struct {
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

func (to *ModelDatabricks) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ModelDatabricks) {
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

func (to *ModelDatabricks) SyncFieldsDuringRead(ctx context.Context, from ModelDatabricks) {
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

func (m ModelDatabricks) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ModelDatabricks) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"latest_versions": reflect.TypeOf(ModelVersion{}),
		"tags":            reflect.TypeOf(ModelTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelDatabricks
// only implements ToObjectValue() and Type().
func (m ModelDatabricks) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ModelDatabricks) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creation_timestamp":     types.Int64Type,
			"description":            types.StringType,
			"id":                     types.StringType,
			"last_updated_timestamp": types.Int64Type,
			"latest_versions": basetypes.ListType{
				ElemType: ModelVersion{}.Type(ctx),
			},
			"name":             types.StringType,
			"permission_level": types.StringType,
			"tags": basetypes.ListType{
				ElemType: ModelTag{}.Type(ctx),
			},
			"user_id": types.StringType,
		},
	}
}

// GetLatestVersions returns the value of the LatestVersions field in ModelDatabricks as
// a slice of ModelVersion values.
// If the field is unknown or null, the boolean return value is false.
func (m *ModelDatabricks) GetLatestVersions(ctx context.Context) ([]ModelVersion, bool) {
	if m.LatestVersions.IsNull() || m.LatestVersions.IsUnknown() {
		return nil, false
	}
	var v []ModelVersion
	d := m.LatestVersions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLatestVersions sets the value of the LatestVersions field in ModelDatabricks.
func (m *ModelDatabricks) SetLatestVersions(ctx context.Context, v []ModelVersion) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["latest_versions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.LatestVersions = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in ModelDatabricks as
// a slice of ModelTag values.
// If the field is unknown or null, the boolean return value is false.
func (m *ModelDatabricks) GetTags(ctx context.Context) ([]ModelTag, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []ModelTag
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in ModelDatabricks.
func (m *ModelDatabricks) SetTags(ctx context.Context, v []ModelTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

// Represents a LoggedModel or Registered Model Version input to a Run.
type ModelInput struct {
	// The unique identifier of the model.
	ModelId types.String `tfsdk:"model_id"`
}

func (to *ModelInput) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ModelInput) {
}

func (to *ModelInput) SyncFieldsDuringRead(ctx context.Context, from ModelInput) {
}

func (m ModelInput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ModelInput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelInput
// only implements ToObjectValue() and Type().
func (m ModelInput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_id": m.ModelId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ModelInput) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_id": types.StringType,
		},
	}
}

// Represents a LoggedModel output of a Run.
type ModelOutput struct {
	// The unique identifier of the model.
	ModelId types.String `tfsdk:"model_id"`
	// The step at which the model was produced.
	Step types.Int64 `tfsdk:"step"`
}

func (to *ModelOutput) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ModelOutput) {
}

func (to *ModelOutput) SyncFieldsDuringRead(ctx context.Context, from ModelOutput) {
}

func (m ModelOutput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ModelOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelOutput
// only implements ToObjectValue() and Type().
func (m ModelOutput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_id": m.ModelId,
			"step":     m.Step,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ModelOutput) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_id": types.StringType,
			"step":     types.Int64Type,
		},
	}
}

// Tag for a registered model
type ModelTag struct {
	// The tag key.
	Key types.String `tfsdk:"key"`
	// The tag value.
	Value types.String `tfsdk:"value"`
}

func (to *ModelTag) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ModelTag) {
}

func (to *ModelTag) SyncFieldsDuringRead(ctx context.Context, from ModelTag) {
}

func (m ModelTag) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ModelTag) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelTag
// only implements ToObjectValue() and Type().
func (m ModelTag) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ModelTag) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

type ModelVersion struct {
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

func (to *ModelVersion) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ModelVersion) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *ModelVersion) SyncFieldsDuringRead(ctx context.Context, from ModelVersion) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m ModelVersion) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ModelVersion) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(ModelVersionTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelVersion
// only implements ToObjectValue() and Type().
func (m ModelVersion) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ModelVersion) Type(ctx context.Context) attr.Type {
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
				ElemType: ModelVersionTag{}.Type(ctx),
			},
			"user_id": types.StringType,
			"version": types.StringType,
		},
	}
}

// GetTags returns the value of the Tags field in ModelVersion as
// a slice of ModelVersionTag values.
// If the field is unknown or null, the boolean return value is false.
func (m *ModelVersion) GetTags(ctx context.Context) ([]ModelVersionTag, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []ModelVersionTag
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in ModelVersion.
func (m *ModelVersion) SetTags(ctx context.Context, v []ModelVersionTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type ModelVersionDatabricks struct {
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
	FeatureList types.Object `tfsdk:"feature_list"`
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

func (to *ModelVersionDatabricks) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ModelVersionDatabricks) {
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

func (to *ModelVersionDatabricks) SyncFieldsDuringRead(ctx context.Context, from ModelVersionDatabricks) {
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

func (m ModelVersionDatabricks) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["creation_timestamp"] = attrs["creation_timestamp"].SetOptional()
	attrs["current_stage"] = attrs["current_stage"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["email_subscription_status"] = attrs["email_subscription_status"].SetOptional()
	attrs["feature_list"] = attrs["feature_list"].SetOptional()
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
func (m ModelVersionDatabricks) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"feature_list":  reflect.TypeOf(FeatureList{}),
		"open_requests": reflect.TypeOf(Activity{}),
		"tags":          reflect.TypeOf(ModelVersionTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelVersionDatabricks
// only implements ToObjectValue() and Type().
func (m ModelVersionDatabricks) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ModelVersionDatabricks) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creation_timestamp":        types.Int64Type,
			"current_stage":             types.StringType,
			"description":               types.StringType,
			"email_subscription_status": types.StringType,
			"feature_list":              FeatureList{}.Type(ctx),
			"last_updated_timestamp":    types.Int64Type,
			"name":                      types.StringType,
			"open_requests": basetypes.ListType{
				ElemType: Activity{}.Type(ctx),
			},
			"permission_level": types.StringType,
			"run_id":           types.StringType,
			"run_link":         types.StringType,
			"source":           types.StringType,
			"status":           types.StringType,
			"status_message":   types.StringType,
			"tags": basetypes.ListType{
				ElemType: ModelVersionTag{}.Type(ctx),
			},
			"user_id": types.StringType,
			"version": types.StringType,
		},
	}
}

// GetFeatureList returns the value of the FeatureList field in ModelVersionDatabricks as
// a FeatureList value.
// If the field is unknown or null, the boolean return value is false.
func (m *ModelVersionDatabricks) GetFeatureList(ctx context.Context) (FeatureList, bool) {
	var e FeatureList
	if m.FeatureList.IsNull() || m.FeatureList.IsUnknown() {
		return e, false
	}
	var v FeatureList
	d := m.FeatureList.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFeatureList sets the value of the FeatureList field in ModelVersionDatabricks.
func (m *ModelVersionDatabricks) SetFeatureList(ctx context.Context, v FeatureList) {
	vs := v.ToObjectValue(ctx)
	m.FeatureList = vs
}

// GetOpenRequests returns the value of the OpenRequests field in ModelVersionDatabricks as
// a slice of Activity values.
// If the field is unknown or null, the boolean return value is false.
func (m *ModelVersionDatabricks) GetOpenRequests(ctx context.Context) ([]Activity, bool) {
	if m.OpenRequests.IsNull() || m.OpenRequests.IsUnknown() {
		return nil, false
	}
	var v []Activity
	d := m.OpenRequests.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOpenRequests sets the value of the OpenRequests field in ModelVersionDatabricks.
func (m *ModelVersionDatabricks) SetOpenRequests(ctx context.Context, v []Activity) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["open_requests"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.OpenRequests = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in ModelVersionDatabricks as
// a slice of ModelVersionTag values.
// If the field is unknown or null, the boolean return value is false.
func (m *ModelVersionDatabricks) GetTags(ctx context.Context) ([]ModelVersionTag, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []ModelVersionTag
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in ModelVersionDatabricks.
func (m *ModelVersionDatabricks) SetTags(ctx context.Context, v []ModelVersionTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type ModelVersionTag struct {
	// The tag key.
	Key types.String `tfsdk:"key"`
	// The tag value.
	Value types.String `tfsdk:"value"`
}

func (to *ModelVersionTag) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ModelVersionTag) {
}

func (to *ModelVersionTag) SyncFieldsDuringRead(ctx context.Context, from ModelVersionTag) {
}

func (m ModelVersionTag) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ModelVersionTag) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelVersionTag
// only implements ToObjectValue() and Type().
func (m ModelVersionTag) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ModelVersionTag) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

// An OnlineStore is a logical database instance that stores and serves features
// online.
type OnlineStore struct {
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

func (to *OnlineStore) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from OnlineStore) {
}

func (to *OnlineStore) SyncFieldsDuringRead(ctx context.Context, from OnlineStore) {
}

func (m OnlineStore) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m OnlineStore) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, OnlineStore
// only implements ToObjectValue() and Type().
func (m OnlineStore) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m OnlineStore) Type(ctx context.Context) attr.Type {
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
type Param struct {
	// Key identifying this param.
	Key types.String `tfsdk:"key"`
	// Value associated with this param.
	Value types.String `tfsdk:"value"`
}

func (to *Param) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Param) {
}

func (to *Param) SyncFieldsDuringRead(ctx context.Context, from Param) {
}

func (m Param) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m Param) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Param
// only implements ToObjectValue() and Type().
func (m Param) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Param) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

type PublishSpec struct {
	// The name of the target online store.
	OnlineStore types.String `tfsdk:"online_store"`
	// The full three-part (catalog, schema, table) name of the online table.
	OnlineTableName types.String `tfsdk:"online_table_name"`
	// The publish mode of the pipeline that syncs the online table with the
	// source table.
	PublishMode types.String `tfsdk:"publish_mode"`
}

func (to *PublishSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PublishSpec) {
}

func (to *PublishSpec) SyncFieldsDuringRead(ctx context.Context, from PublishSpec) {
}

func (m PublishSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PublishSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PublishSpec
// only implements ToObjectValue() and Type().
func (m PublishSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"online_store":      m.OnlineStore,
			"online_table_name": m.OnlineTableName,
			"publish_mode":      m.PublishMode,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PublishSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"online_store":      types.StringType,
			"online_table_name": types.StringType,
			"publish_mode":      types.StringType,
		},
	}
}

type PublishTableRequest struct {
	// The specification for publishing the online table from the source table.
	PublishSpec types.Object `tfsdk:"publish_spec"`
	// The full three-part (catalog, schema, table) name of the source table.
	SourceTableName types.String `tfsdk:"-"`
}

func (to *PublishTableRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PublishTableRequest) {
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

func (to *PublishTableRequest) SyncFieldsDuringRead(ctx context.Context, from PublishTableRequest) {
	if !from.PublishSpec.IsNull() && !from.PublishSpec.IsUnknown() {
		if toPublishSpec, ok := to.GetPublishSpec(ctx); ok {
			if fromPublishSpec, ok := from.GetPublishSpec(ctx); ok {
				toPublishSpec.SyncFieldsDuringRead(ctx, fromPublishSpec)
				to.SetPublishSpec(ctx, toPublishSpec)
			}
		}
	}
}

func (m PublishTableRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["publish_spec"] = attrs["publish_spec"].SetRequired()
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
func (m PublishTableRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"publish_spec": reflect.TypeOf(PublishSpec{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PublishTableRequest
// only implements ToObjectValue() and Type().
func (m PublishTableRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"publish_spec":      m.PublishSpec,
			"source_table_name": m.SourceTableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PublishTableRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"publish_spec":      PublishSpec{}.Type(ctx),
			"source_table_name": types.StringType,
		},
	}
}

// GetPublishSpec returns the value of the PublishSpec field in PublishTableRequest as
// a PublishSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *PublishTableRequest) GetPublishSpec(ctx context.Context) (PublishSpec, bool) {
	var e PublishSpec
	if m.PublishSpec.IsNull() || m.PublishSpec.IsUnknown() {
		return e, false
	}
	var v PublishSpec
	d := m.PublishSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPublishSpec sets the value of the PublishSpec field in PublishTableRequest.
func (m *PublishTableRequest) SetPublishSpec(ctx context.Context, v PublishSpec) {
	vs := v.ToObjectValue(ctx)
	m.PublishSpec = vs
}

type PublishTableResponse struct {
	// The full three-part (catalog, schema, table) name of the online table.
	OnlineTableName types.String `tfsdk:"online_table_name"`
	// The ID of the pipeline that syncs the online table with the source table.
	PipelineId types.String `tfsdk:"pipeline_id"`
}

func (to *PublishTableResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PublishTableResponse) {
}

func (to *PublishTableResponse) SyncFieldsDuringRead(ctx context.Context, from PublishTableResponse) {
}

func (m PublishTableResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PublishTableResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PublishTableResponse
// only implements ToObjectValue() and Type().
func (m PublishTableResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"online_table_name": m.OnlineTableName,
			"pipeline_id":       m.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PublishTableResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"online_table_name": types.StringType,
			"pipeline_id":       types.StringType,
		},
	}
}

type RegisteredModelAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`

	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (to *RegisteredModelAccessControlRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RegisteredModelAccessControlRequest) {
}

func (to *RegisteredModelAccessControlRequest) SyncFieldsDuringRead(ctx context.Context, from RegisteredModelAccessControlRequest) {
}

func (m RegisteredModelAccessControlRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RegisteredModelAccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegisteredModelAccessControlRequest
// only implements ToObjectValue() and Type().
func (m RegisteredModelAccessControlRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m RegisteredModelAccessControlRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type RegisteredModelAccessControlResponse struct {
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

func (to *RegisteredModelAccessControlResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RegisteredModelAccessControlResponse) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (to *RegisteredModelAccessControlResponse) SyncFieldsDuringRead(ctx context.Context, from RegisteredModelAccessControlResponse) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (m RegisteredModelAccessControlResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RegisteredModelAccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(RegisteredModelPermission{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegisteredModelAccessControlResponse
// only implements ToObjectValue() and Type().
func (m RegisteredModelAccessControlResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m RegisteredModelAccessControlResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: RegisteredModelPermission{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

// GetAllPermissions returns the value of the AllPermissions field in RegisteredModelAccessControlResponse as
// a slice of RegisteredModelPermission values.
// If the field is unknown or null, the boolean return value is false.
func (m *RegisteredModelAccessControlResponse) GetAllPermissions(ctx context.Context) ([]RegisteredModelPermission, bool) {
	if m.AllPermissions.IsNull() || m.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []RegisteredModelPermission
	d := m.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in RegisteredModelAccessControlResponse.
func (m *RegisteredModelAccessControlResponse) SetAllPermissions(ctx context.Context, v []RegisteredModelPermission) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllPermissions = types.ListValueMust(t, vs)
}

type RegisteredModelPermission struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *RegisteredModelPermission) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RegisteredModelPermission) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (to *RegisteredModelPermission) SyncFieldsDuringRead(ctx context.Context, from RegisteredModelPermission) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (m RegisteredModelPermission) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RegisteredModelPermission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegisteredModelPermission
// only implements ToObjectValue() and Type().
func (m RegisteredModelPermission) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             m.Inherited,
			"inherited_from_object": m.InheritedFromObject,
			"permission_level":      m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RegisteredModelPermission) Type(ctx context.Context) attr.Type {
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

// GetInheritedFromObject returns the value of the InheritedFromObject field in RegisteredModelPermission as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *RegisteredModelPermission) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
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

// SetInheritedFromObject sets the value of the InheritedFromObject field in RegisteredModelPermission.
func (m *RegisteredModelPermission) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InheritedFromObject = types.ListValueMust(t, vs)
}

type RegisteredModelPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (to *RegisteredModelPermissions) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RegisteredModelPermissions) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *RegisteredModelPermissions) SyncFieldsDuringRead(ctx context.Context, from RegisteredModelPermissions) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m RegisteredModelPermissions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RegisteredModelPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(RegisteredModelAccessControlResponse{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegisteredModelPermissions
// only implements ToObjectValue() and Type().
func (m RegisteredModelPermissions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"object_id":           m.ObjectId,
			"object_type":         m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RegisteredModelPermissions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: RegisteredModelAccessControlResponse{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in RegisteredModelPermissions as
// a slice of RegisteredModelAccessControlResponse values.
// If the field is unknown or null, the boolean return value is false.
func (m *RegisteredModelPermissions) GetAccessControlList(ctx context.Context) ([]RegisteredModelAccessControlResponse, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []RegisteredModelAccessControlResponse
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in RegisteredModelPermissions.
func (m *RegisteredModelPermissions) SetAccessControlList(ctx context.Context, v []RegisteredModelAccessControlResponse) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

type RegisteredModelPermissionsDescription struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *RegisteredModelPermissionsDescription) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RegisteredModelPermissionsDescription) {
}

func (to *RegisteredModelPermissionsDescription) SyncFieldsDuringRead(ctx context.Context, from RegisteredModelPermissionsDescription) {
}

func (m RegisteredModelPermissionsDescription) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RegisteredModelPermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegisteredModelPermissionsDescription
// only implements ToObjectValue() and Type().
func (m RegisteredModelPermissionsDescription) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      m.Description,
			"permission_level": m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RegisteredModelPermissionsDescription) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type RegisteredModelPermissionsRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// The registered model for which to get or manage permissions.
	RegisteredModelId types.String `tfsdk:"-"`
}

func (to *RegisteredModelPermissionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RegisteredModelPermissionsRequest) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *RegisteredModelPermissionsRequest) SyncFieldsDuringRead(ctx context.Context, from RegisteredModelPermissionsRequest) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m RegisteredModelPermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RegisteredModelPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(RegisteredModelAccessControlRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegisteredModelPermissionsRequest
// only implements ToObjectValue() and Type().
func (m RegisteredModelPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"registered_model_id": m.RegisteredModelId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RegisteredModelPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: RegisteredModelAccessControlRequest{}.Type(ctx),
			},
			"registered_model_id": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in RegisteredModelPermissionsRequest as
// a slice of RegisteredModelAccessControlRequest values.
// If the field is unknown or null, the boolean return value is false.
func (m *RegisteredModelPermissionsRequest) GetAccessControlList(ctx context.Context) ([]RegisteredModelAccessControlRequest, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []RegisteredModelAccessControlRequest
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in RegisteredModelPermissionsRequest.
func (m *RegisteredModelPermissionsRequest) SetAccessControlList(ctx context.Context, v []RegisteredModelAccessControlRequest) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

type RegistryWebhook struct {
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

	HttpUrlSpec types.Object `tfsdk:"http_url_spec"`
	// Webhook ID
	Id types.String `tfsdk:"id"`

	JobSpec types.Object `tfsdk:"job_spec"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp"`
	// Name of the model whose events would trigger this webhook.
	ModelName types.String `tfsdk:"model_name"`

	Status types.String `tfsdk:"status"`
}

func (to *RegistryWebhook) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RegistryWebhook) {
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

func (to *RegistryWebhook) SyncFieldsDuringRead(ctx context.Context, from RegistryWebhook) {
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

func (m RegistryWebhook) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["creation_timestamp"] = attrs["creation_timestamp"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["events"] = attrs["events"].SetOptional()
	attrs["http_url_spec"] = attrs["http_url_spec"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["job_spec"] = attrs["job_spec"].SetOptional()
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
func (m RegistryWebhook) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"events":        reflect.TypeOf(types.String{}),
		"http_url_spec": reflect.TypeOf(HttpUrlSpecWithoutSecret{}),
		"job_spec":      reflect.TypeOf(JobSpecWithoutSecret{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegistryWebhook
// only implements ToObjectValue() and Type().
func (m RegistryWebhook) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m RegistryWebhook) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creation_timestamp": types.Int64Type,
			"description":        types.StringType,
			"events": basetypes.ListType{
				ElemType: types.StringType,
			},
			"http_url_spec":          HttpUrlSpecWithoutSecret{}.Type(ctx),
			"id":                     types.StringType,
			"job_spec":               JobSpecWithoutSecret{}.Type(ctx),
			"last_updated_timestamp": types.Int64Type,
			"model_name":             types.StringType,
			"status":                 types.StringType,
		},
	}
}

// GetEvents returns the value of the Events field in RegistryWebhook as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *RegistryWebhook) GetEvents(ctx context.Context) ([]types.String, bool) {
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

// SetEvents sets the value of the Events field in RegistryWebhook.
func (m *RegistryWebhook) SetEvents(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["events"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Events = types.ListValueMust(t, vs)
}

// GetHttpUrlSpec returns the value of the HttpUrlSpec field in RegistryWebhook as
// a HttpUrlSpecWithoutSecret value.
// If the field is unknown or null, the boolean return value is false.
func (m *RegistryWebhook) GetHttpUrlSpec(ctx context.Context) (HttpUrlSpecWithoutSecret, bool) {
	var e HttpUrlSpecWithoutSecret
	if m.HttpUrlSpec.IsNull() || m.HttpUrlSpec.IsUnknown() {
		return e, false
	}
	var v HttpUrlSpecWithoutSecret
	d := m.HttpUrlSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetHttpUrlSpec sets the value of the HttpUrlSpec field in RegistryWebhook.
func (m *RegistryWebhook) SetHttpUrlSpec(ctx context.Context, v HttpUrlSpecWithoutSecret) {
	vs := v.ToObjectValue(ctx)
	m.HttpUrlSpec = vs
}

// GetJobSpec returns the value of the JobSpec field in RegistryWebhook as
// a JobSpecWithoutSecret value.
// If the field is unknown or null, the boolean return value is false.
func (m *RegistryWebhook) GetJobSpec(ctx context.Context) (JobSpecWithoutSecret, bool) {
	var e JobSpecWithoutSecret
	if m.JobSpec.IsNull() || m.JobSpec.IsUnknown() {
		return e, false
	}
	var v JobSpecWithoutSecret
	d := m.JobSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetJobSpec sets the value of the JobSpec field in RegistryWebhook.
func (m *RegistryWebhook) SetJobSpec(ctx context.Context, v JobSpecWithoutSecret) {
	vs := v.ToObjectValue(ctx)
	m.JobSpec = vs
}

// Details required to identify and reject a model version stage transition
// request.
type RejectTransitionRequest struct {
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

func (to *RejectTransitionRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RejectTransitionRequest) {
}

func (to *RejectTransitionRequest) SyncFieldsDuringRead(ctx context.Context, from RejectTransitionRequest) {
}

func (m RejectTransitionRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RejectTransitionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RejectTransitionRequest
// only implements ToObjectValue() and Type().
func (m RejectTransitionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m RejectTransitionRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment": types.StringType,
			"name":    types.StringType,
			"stage":   types.StringType,
			"version": types.StringType,
		},
	}
}

type RejectTransitionRequestResponse struct {
	// New activity generated as a result of this operation.
	Activity types.Object `tfsdk:"activity"`
}

func (to *RejectTransitionRequestResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RejectTransitionRequestResponse) {
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

func (to *RejectTransitionRequestResponse) SyncFieldsDuringRead(ctx context.Context, from RejectTransitionRequestResponse) {
	if !from.Activity.IsNull() && !from.Activity.IsUnknown() {
		if toActivity, ok := to.GetActivity(ctx); ok {
			if fromActivity, ok := from.GetActivity(ctx); ok {
				toActivity.SyncFieldsDuringRead(ctx, fromActivity)
				to.SetActivity(ctx, toActivity)
			}
		}
	}
}

func (m RejectTransitionRequestResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["activity"] = attrs["activity"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RejectTransitionRequestResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RejectTransitionRequestResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"activity": reflect.TypeOf(Activity{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RejectTransitionRequestResponse
// only implements ToObjectValue() and Type().
func (m RejectTransitionRequestResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"activity": m.Activity,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RejectTransitionRequestResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"activity": Activity{}.Type(ctx),
		},
	}
}

// GetActivity returns the value of the Activity field in RejectTransitionRequestResponse as
// a Activity value.
// If the field is unknown or null, the boolean return value is false.
func (m *RejectTransitionRequestResponse) GetActivity(ctx context.Context) (Activity, bool) {
	var e Activity
	if m.Activity.IsNull() || m.Activity.IsUnknown() {
		return e, false
	}
	var v Activity
	d := m.Activity.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetActivity sets the value of the Activity field in RejectTransitionRequestResponse.
func (m *RejectTransitionRequestResponse) SetActivity(ctx context.Context, v Activity) {
	vs := v.ToObjectValue(ctx)
	m.Activity = vs
}

type RenameModelRequest struct {
	// Registered model unique name identifier.
	Name types.String `tfsdk:"name"`
	// If provided, updates the name for this `registered_model`.
	NewName types.String `tfsdk:"new_name"`
}

func (to *RenameModelRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RenameModelRequest) {
}

func (to *RenameModelRequest) SyncFieldsDuringRead(ctx context.Context, from RenameModelRequest) {
}

func (m RenameModelRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RenameModelRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RenameModelRequest
// only implements ToObjectValue() and Type().
func (m RenameModelRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":     m.Name,
			"new_name": m.NewName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RenameModelRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":     types.StringType,
			"new_name": types.StringType,
		},
	}
}

type RenameModelResponse struct {
	RegisteredModel types.Object `tfsdk:"registered_model"`
}

func (to *RenameModelResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RenameModelResponse) {
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

func (to *RenameModelResponse) SyncFieldsDuringRead(ctx context.Context, from RenameModelResponse) {
	if !from.RegisteredModel.IsNull() && !from.RegisteredModel.IsUnknown() {
		if toRegisteredModel, ok := to.GetRegisteredModel(ctx); ok {
			if fromRegisteredModel, ok := from.GetRegisteredModel(ctx); ok {
				toRegisteredModel.SyncFieldsDuringRead(ctx, fromRegisteredModel)
				to.SetRegisteredModel(ctx, toRegisteredModel)
			}
		}
	}
}

func (m RenameModelResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["registered_model"] = attrs["registered_model"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RenameModelResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RenameModelResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"registered_model": reflect.TypeOf(Model{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RenameModelResponse
// only implements ToObjectValue() and Type().
func (m RenameModelResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"registered_model": m.RegisteredModel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RenameModelResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"registered_model": Model{}.Type(ctx),
		},
	}
}

// GetRegisteredModel returns the value of the RegisteredModel field in RenameModelResponse as
// a Model value.
// If the field is unknown or null, the boolean return value is false.
func (m *RenameModelResponse) GetRegisteredModel(ctx context.Context) (Model, bool) {
	var e Model
	if m.RegisteredModel.IsNull() || m.RegisteredModel.IsUnknown() {
		return e, false
	}
	var v Model
	d := m.RegisteredModel.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRegisteredModel sets the value of the RegisteredModel field in RenameModelResponse.
func (m *RenameModelResponse) SetRegisteredModel(ctx context.Context, v Model) {
	vs := v.ToObjectValue(ctx)
	m.RegisteredModel = vs
}

type RestoreExperiment struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"experiment_id"`
}

func (to *RestoreExperiment) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RestoreExperiment) {
}

func (to *RestoreExperiment) SyncFieldsDuringRead(ctx context.Context, from RestoreExperiment) {
}

func (m RestoreExperiment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RestoreExperiment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestoreExperiment
// only implements ToObjectValue() and Type().
func (m RestoreExperiment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": m.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RestoreExperiment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
		},
	}
}

type RestoreExperimentResponse struct {
}

func (to *RestoreExperimentResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RestoreExperimentResponse) {
}

func (to *RestoreExperimentResponse) SyncFieldsDuringRead(ctx context.Context, from RestoreExperimentResponse) {
}

func (m RestoreExperimentResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestoreExperimentResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RestoreExperimentResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestoreExperimentResponse
// only implements ToObjectValue() and Type().
func (m RestoreExperimentResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m RestoreExperimentResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type RestoreRun struct {
	// ID of the run to restore.
	RunId types.String `tfsdk:"run_id"`
}

func (to *RestoreRun) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RestoreRun) {
}

func (to *RestoreRun) SyncFieldsDuringRead(ctx context.Context, from RestoreRun) {
}

func (m RestoreRun) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RestoreRun) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestoreRun
// only implements ToObjectValue() and Type().
func (m RestoreRun) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run_id": m.RunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RestoreRun) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run_id": types.StringType,
		},
	}
}

type RestoreRunResponse struct {
}

func (to *RestoreRunResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RestoreRunResponse) {
}

func (to *RestoreRunResponse) SyncFieldsDuringRead(ctx context.Context, from RestoreRunResponse) {
}

func (m RestoreRunResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestoreRunResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RestoreRunResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestoreRunResponse
// only implements ToObjectValue() and Type().
func (m RestoreRunResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m RestoreRunResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type RestoreRuns struct {
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

func (to *RestoreRuns) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RestoreRuns) {
}

func (to *RestoreRuns) SyncFieldsDuringRead(ctx context.Context, from RestoreRuns) {
}

func (m RestoreRuns) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RestoreRuns) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestoreRuns
// only implements ToObjectValue() and Type().
func (m RestoreRuns) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id":        m.ExperimentId,
			"max_runs":             m.MaxRuns,
			"min_timestamp_millis": m.MinTimestampMillis,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RestoreRuns) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id":        types.StringType,
			"max_runs":             types.Int64Type,
			"min_timestamp_millis": types.Int64Type,
		},
	}
}

type RestoreRunsResponse struct {
	// The number of runs restored.
	RunsRestored types.Int64 `tfsdk:"runs_restored"`
}

func (to *RestoreRunsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RestoreRunsResponse) {
}

func (to *RestoreRunsResponse) SyncFieldsDuringRead(ctx context.Context, from RestoreRunsResponse) {
}

func (m RestoreRunsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RestoreRunsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestoreRunsResponse
// only implements ToObjectValue() and Type().
func (m RestoreRunsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"runs_restored": m.RunsRestored,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RestoreRunsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"runs_restored": types.Int64Type,
		},
	}
}

// A single run.
type Run struct {
	// Run data.
	Data types.Object `tfsdk:"data"`
	// Run metadata.
	Info types.Object `tfsdk:"info"`
	// Run inputs.
	Inputs types.Object `tfsdk:"inputs"`
}

func (to *Run) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Run) {
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

func (to *Run) SyncFieldsDuringRead(ctx context.Context, from Run) {
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

func (m Run) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["data"] = attrs["data"].SetOptional()
	attrs["info"] = attrs["info"].SetOptional()
	attrs["inputs"] = attrs["inputs"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Run.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Run) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data":   reflect.TypeOf(RunData{}),
		"info":   reflect.TypeOf(RunInfo{}),
		"inputs": reflect.TypeOf(RunInputs{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Run
// only implements ToObjectValue() and Type().
func (m Run) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data":   m.Data,
			"info":   m.Info,
			"inputs": m.Inputs,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Run) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"data":   RunData{}.Type(ctx),
			"info":   RunInfo{}.Type(ctx),
			"inputs": RunInputs{}.Type(ctx),
		},
	}
}

// GetData returns the value of the Data field in Run as
// a RunData value.
// If the field is unknown or null, the boolean return value is false.
func (m *Run) GetData(ctx context.Context) (RunData, bool) {
	var e RunData
	if m.Data.IsNull() || m.Data.IsUnknown() {
		return e, false
	}
	var v RunData
	d := m.Data.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetData sets the value of the Data field in Run.
func (m *Run) SetData(ctx context.Context, v RunData) {
	vs := v.ToObjectValue(ctx)
	m.Data = vs
}

// GetInfo returns the value of the Info field in Run as
// a RunInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *Run) GetInfo(ctx context.Context) (RunInfo, bool) {
	var e RunInfo
	if m.Info.IsNull() || m.Info.IsUnknown() {
		return e, false
	}
	var v RunInfo
	d := m.Info.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInfo sets the value of the Info field in Run.
func (m *Run) SetInfo(ctx context.Context, v RunInfo) {
	vs := v.ToObjectValue(ctx)
	m.Info = vs
}

// GetInputs returns the value of the Inputs field in Run as
// a RunInputs value.
// If the field is unknown or null, the boolean return value is false.
func (m *Run) GetInputs(ctx context.Context) (RunInputs, bool) {
	var e RunInputs
	if m.Inputs.IsNull() || m.Inputs.IsUnknown() {
		return e, false
	}
	var v RunInputs
	d := m.Inputs.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInputs sets the value of the Inputs field in Run.
func (m *Run) SetInputs(ctx context.Context, v RunInputs) {
	vs := v.ToObjectValue(ctx)
	m.Inputs = vs
}

// Run data (metrics, params, and tags).
type RunData struct {
	// Run metrics.
	Metrics types.List `tfsdk:"metrics"`
	// Run parameters.
	Params types.List `tfsdk:"params"`
	// Additional metadata key-value pairs.
	Tags types.List `tfsdk:"tags"`
}

func (to *RunData) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RunData) {
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

func (to *RunData) SyncFieldsDuringRead(ctx context.Context, from RunData) {
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

func (m RunData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RunData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metrics": reflect.TypeOf(Metric{}),
		"params":  reflect.TypeOf(Param{}),
		"tags":    reflect.TypeOf(RunTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RunData
// only implements ToObjectValue() and Type().
func (m RunData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metrics": m.Metrics,
			"params":  m.Params,
			"tags":    m.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RunData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"metrics": basetypes.ListType{
				ElemType: Metric{}.Type(ctx),
			},
			"params": basetypes.ListType{
				ElemType: Param{}.Type(ctx),
			},
			"tags": basetypes.ListType{
				ElemType: RunTag{}.Type(ctx),
			},
		},
	}
}

// GetMetrics returns the value of the Metrics field in RunData as
// a slice of Metric values.
// If the field is unknown or null, the boolean return value is false.
func (m *RunData) GetMetrics(ctx context.Context) ([]Metric, bool) {
	if m.Metrics.IsNull() || m.Metrics.IsUnknown() {
		return nil, false
	}
	var v []Metric
	d := m.Metrics.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMetrics sets the value of the Metrics field in RunData.
func (m *RunData) SetMetrics(ctx context.Context, v []Metric) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["metrics"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Metrics = types.ListValueMust(t, vs)
}

// GetParams returns the value of the Params field in RunData as
// a slice of Param values.
// If the field is unknown or null, the boolean return value is false.
func (m *RunData) GetParams(ctx context.Context) ([]Param, bool) {
	if m.Params.IsNull() || m.Params.IsUnknown() {
		return nil, false
	}
	var v []Param
	d := m.Params.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParams sets the value of the Params field in RunData.
func (m *RunData) SetParams(ctx context.Context, v []Param) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["params"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Params = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in RunData as
// a slice of RunTag values.
// If the field is unknown or null, the boolean return value is false.
func (m *RunData) GetTags(ctx context.Context) ([]RunTag, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []RunTag
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in RunData.
func (m *RunData) SetTags(ctx context.Context, v []RunTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

// Metadata of a single run.
type RunInfo struct {
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

func (to *RunInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RunInfo) {
}

func (to *RunInfo) SyncFieldsDuringRead(ctx context.Context, from RunInfo) {
}

func (m RunInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RunInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RunInfo
// only implements ToObjectValue() and Type().
func (m RunInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m RunInfo) Type(ctx context.Context) attr.Type {
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
type RunInputs struct {
	// Run metrics.
	DatasetInputs types.List `tfsdk:"dataset_inputs"`
	// Model inputs to the Run.
	ModelInputs types.List `tfsdk:"model_inputs"`
}

func (to *RunInputs) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RunInputs) {
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

func (to *RunInputs) SyncFieldsDuringRead(ctx context.Context, from RunInputs) {
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

func (m RunInputs) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RunInputs) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dataset_inputs": reflect.TypeOf(DatasetInput{}),
		"model_inputs":   reflect.TypeOf(ModelInput{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RunInputs
// only implements ToObjectValue() and Type().
func (m RunInputs) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dataset_inputs": m.DatasetInputs,
			"model_inputs":   m.ModelInputs,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RunInputs) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dataset_inputs": basetypes.ListType{
				ElemType: DatasetInput{}.Type(ctx),
			},
			"model_inputs": basetypes.ListType{
				ElemType: ModelInput{}.Type(ctx),
			},
		},
	}
}

// GetDatasetInputs returns the value of the DatasetInputs field in RunInputs as
// a slice of DatasetInput values.
// If the field is unknown or null, the boolean return value is false.
func (m *RunInputs) GetDatasetInputs(ctx context.Context) ([]DatasetInput, bool) {
	if m.DatasetInputs.IsNull() || m.DatasetInputs.IsUnknown() {
		return nil, false
	}
	var v []DatasetInput
	d := m.DatasetInputs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatasetInputs sets the value of the DatasetInputs field in RunInputs.
func (m *RunInputs) SetDatasetInputs(ctx context.Context, v []DatasetInput) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["dataset_inputs"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DatasetInputs = types.ListValueMust(t, vs)
}

// GetModelInputs returns the value of the ModelInputs field in RunInputs as
// a slice of ModelInput values.
// If the field is unknown or null, the boolean return value is false.
func (m *RunInputs) GetModelInputs(ctx context.Context) ([]ModelInput, bool) {
	if m.ModelInputs.IsNull() || m.ModelInputs.IsUnknown() {
		return nil, false
	}
	var v []ModelInput
	d := m.ModelInputs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModelInputs sets the value of the ModelInputs field in RunInputs.
func (m *RunInputs) SetModelInputs(ctx context.Context, v []ModelInput) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["model_inputs"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ModelInputs = types.ListValueMust(t, vs)
}

// Tag for a run.
type RunTag struct {
	// The tag key.
	Key types.String `tfsdk:"key"`
	// The tag value.
	Value types.String `tfsdk:"value"`
}

func (to *RunTag) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RunTag) {
}

func (to *RunTag) SyncFieldsDuringRead(ctx context.Context, from RunTag) {
}

func (m RunTag) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RunTag) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RunTag
// only implements ToObjectValue() and Type().
func (m RunTag) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RunTag) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

type SearchExperiments struct {
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

func (to *SearchExperiments) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SearchExperiments) {
	if !from.OrderBy.IsNull() && !from.OrderBy.IsUnknown() && to.OrderBy.IsNull() && len(from.OrderBy.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OrderBy, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OrderBy = from.OrderBy
	}
}

func (to *SearchExperiments) SyncFieldsDuringRead(ctx context.Context, from SearchExperiments) {
	if !from.OrderBy.IsNull() && !from.OrderBy.IsUnknown() && to.OrderBy.IsNull() && len(from.OrderBy.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OrderBy, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OrderBy = from.OrderBy
	}
}

func (m SearchExperiments) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SearchExperiments) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"order_by": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchExperiments
// only implements ToObjectValue() and Type().
func (m SearchExperiments) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m SearchExperiments) Type(ctx context.Context) attr.Type {
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

// GetOrderBy returns the value of the OrderBy field in SearchExperiments as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *SearchExperiments) GetOrderBy(ctx context.Context) ([]types.String, bool) {
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

// SetOrderBy sets the value of the OrderBy field in SearchExperiments.
func (m *SearchExperiments) SetOrderBy(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["order_by"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.OrderBy = types.ListValueMust(t, vs)
}

type SearchExperimentsResponse struct {
	// Experiments that match the search criteria
	Experiments types.List `tfsdk:"experiments"`
	// Token that can be used to retrieve the next page of experiments. An empty
	// token means that no more experiments are available for retrieval.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *SearchExperimentsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SearchExperimentsResponse) {
	if !from.Experiments.IsNull() && !from.Experiments.IsUnknown() && to.Experiments.IsNull() && len(from.Experiments.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Experiments, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Experiments = from.Experiments
	}
}

func (to *SearchExperimentsResponse) SyncFieldsDuringRead(ctx context.Context, from SearchExperimentsResponse) {
	if !from.Experiments.IsNull() && !from.Experiments.IsUnknown() && to.Experiments.IsNull() && len(from.Experiments.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Experiments, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Experiments = from.Experiments
	}
}

func (m SearchExperimentsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SearchExperimentsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"experiments": reflect.TypeOf(Experiment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchExperimentsResponse
// only implements ToObjectValue() and Type().
func (m SearchExperimentsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiments":     m.Experiments,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SearchExperimentsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiments": basetypes.ListType{
				ElemType: Experiment{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetExperiments returns the value of the Experiments field in SearchExperimentsResponse as
// a slice of Experiment values.
// If the field is unknown or null, the boolean return value is false.
func (m *SearchExperimentsResponse) GetExperiments(ctx context.Context) ([]Experiment, bool) {
	if m.Experiments.IsNull() || m.Experiments.IsUnknown() {
		return nil, false
	}
	var v []Experiment
	d := m.Experiments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExperiments sets the value of the Experiments field in SearchExperimentsResponse.
func (m *SearchExperimentsResponse) SetExperiments(ctx context.Context, v []Experiment) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["experiments"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Experiments = types.ListValueMust(t, vs)
}

type SearchLoggedModelsDataset struct {
	// The digest of the dataset.
	DatasetDigest types.String `tfsdk:"dataset_digest"`
	// The name of the dataset.
	DatasetName types.String `tfsdk:"dataset_name"`
}

func (to *SearchLoggedModelsDataset) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SearchLoggedModelsDataset) {
}

func (to *SearchLoggedModelsDataset) SyncFieldsDuringRead(ctx context.Context, from SearchLoggedModelsDataset) {
}

func (m SearchLoggedModelsDataset) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SearchLoggedModelsDataset) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchLoggedModelsDataset
// only implements ToObjectValue() and Type().
func (m SearchLoggedModelsDataset) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dataset_digest": m.DatasetDigest,
			"dataset_name":   m.DatasetName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SearchLoggedModelsDataset) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dataset_digest": types.StringType,
			"dataset_name":   types.StringType,
		},
	}
}

type SearchLoggedModelsOrderBy struct {
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

func (to *SearchLoggedModelsOrderBy) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SearchLoggedModelsOrderBy) {
}

func (to *SearchLoggedModelsOrderBy) SyncFieldsDuringRead(ctx context.Context, from SearchLoggedModelsOrderBy) {
}

func (m SearchLoggedModelsOrderBy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SearchLoggedModelsOrderBy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchLoggedModelsOrderBy
// only implements ToObjectValue() and Type().
func (m SearchLoggedModelsOrderBy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m SearchLoggedModelsOrderBy) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ascending":      types.BoolType,
			"dataset_digest": types.StringType,
			"dataset_name":   types.StringType,
			"field_name":     types.StringType,
		},
	}
}

type SearchLoggedModelsRequest struct {
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

func (to *SearchLoggedModelsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SearchLoggedModelsRequest) {
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

func (to *SearchLoggedModelsRequest) SyncFieldsDuringRead(ctx context.Context, from SearchLoggedModelsRequest) {
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

func (m SearchLoggedModelsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SearchLoggedModelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"datasets":       reflect.TypeOf(SearchLoggedModelsDataset{}),
		"experiment_ids": reflect.TypeOf(types.String{}),
		"order_by":       reflect.TypeOf(SearchLoggedModelsOrderBy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchLoggedModelsRequest
// only implements ToObjectValue() and Type().
func (m SearchLoggedModelsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m SearchLoggedModelsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"datasets": basetypes.ListType{
				ElemType: SearchLoggedModelsDataset{}.Type(ctx),
			},
			"experiment_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"filter":      types.StringType,
			"max_results": types.Int64Type,
			"order_by": basetypes.ListType{
				ElemType: SearchLoggedModelsOrderBy{}.Type(ctx),
			},
			"page_token": types.StringType,
		},
	}
}

// GetDatasets returns the value of the Datasets field in SearchLoggedModelsRequest as
// a slice of SearchLoggedModelsDataset values.
// If the field is unknown or null, the boolean return value is false.
func (m *SearchLoggedModelsRequest) GetDatasets(ctx context.Context) ([]SearchLoggedModelsDataset, bool) {
	if m.Datasets.IsNull() || m.Datasets.IsUnknown() {
		return nil, false
	}
	var v []SearchLoggedModelsDataset
	d := m.Datasets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatasets sets the value of the Datasets field in SearchLoggedModelsRequest.
func (m *SearchLoggedModelsRequest) SetDatasets(ctx context.Context, v []SearchLoggedModelsDataset) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["datasets"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Datasets = types.ListValueMust(t, vs)
}

// GetExperimentIds returns the value of the ExperimentIds field in SearchLoggedModelsRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *SearchLoggedModelsRequest) GetExperimentIds(ctx context.Context) ([]types.String, bool) {
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

// SetExperimentIds sets the value of the ExperimentIds field in SearchLoggedModelsRequest.
func (m *SearchLoggedModelsRequest) SetExperimentIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["experiment_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ExperimentIds = types.ListValueMust(t, vs)
}

// GetOrderBy returns the value of the OrderBy field in SearchLoggedModelsRequest as
// a slice of SearchLoggedModelsOrderBy values.
// If the field is unknown or null, the boolean return value is false.
func (m *SearchLoggedModelsRequest) GetOrderBy(ctx context.Context) ([]SearchLoggedModelsOrderBy, bool) {
	if m.OrderBy.IsNull() || m.OrderBy.IsUnknown() {
		return nil, false
	}
	var v []SearchLoggedModelsOrderBy
	d := m.OrderBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOrderBy sets the value of the OrderBy field in SearchLoggedModelsRequest.
func (m *SearchLoggedModelsRequest) SetOrderBy(ctx context.Context, v []SearchLoggedModelsOrderBy) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["order_by"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.OrderBy = types.ListValueMust(t, vs)
}

type SearchLoggedModelsResponse struct {
	// Logged models that match the search criteria.
	Models types.List `tfsdk:"models"`
	// The token that can be used to retrieve the next page of logged models.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *SearchLoggedModelsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SearchLoggedModelsResponse) {
	if !from.Models.IsNull() && !from.Models.IsUnknown() && to.Models.IsNull() && len(from.Models.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Models, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Models = from.Models
	}
}

func (to *SearchLoggedModelsResponse) SyncFieldsDuringRead(ctx context.Context, from SearchLoggedModelsResponse) {
	if !from.Models.IsNull() && !from.Models.IsUnknown() && to.Models.IsNull() && len(from.Models.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Models, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Models = from.Models
	}
}

func (m SearchLoggedModelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SearchLoggedModelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"models": reflect.TypeOf(LoggedModel{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchLoggedModelsResponse
// only implements ToObjectValue() and Type().
func (m SearchLoggedModelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"models":          m.Models,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SearchLoggedModelsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"models": basetypes.ListType{
				ElemType: LoggedModel{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetModels returns the value of the Models field in SearchLoggedModelsResponse as
// a slice of LoggedModel values.
// If the field is unknown or null, the boolean return value is false.
func (m *SearchLoggedModelsResponse) GetModels(ctx context.Context) ([]LoggedModel, bool) {
	if m.Models.IsNull() || m.Models.IsUnknown() {
		return nil, false
	}
	var v []LoggedModel
	d := m.Models.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModels sets the value of the Models field in SearchLoggedModelsResponse.
func (m *SearchLoggedModelsResponse) SetModels(ctx context.Context, v []LoggedModel) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["models"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Models = types.ListValueMust(t, vs)
}

type SearchModelVersionsRequest struct {
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

func (to *SearchModelVersionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SearchModelVersionsRequest) {
	if !from.OrderBy.IsNull() && !from.OrderBy.IsUnknown() && to.OrderBy.IsNull() && len(from.OrderBy.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OrderBy, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OrderBy = from.OrderBy
	}
}

func (to *SearchModelVersionsRequest) SyncFieldsDuringRead(ctx context.Context, from SearchModelVersionsRequest) {
	if !from.OrderBy.IsNull() && !from.OrderBy.IsUnknown() && to.OrderBy.IsNull() && len(from.OrderBy.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OrderBy, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OrderBy = from.OrderBy
	}
}

func (m SearchModelVersionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SearchModelVersionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"order_by": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchModelVersionsRequest
// only implements ToObjectValue() and Type().
func (m SearchModelVersionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m SearchModelVersionsRequest) Type(ctx context.Context) attr.Type {
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

// GetOrderBy returns the value of the OrderBy field in SearchModelVersionsRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *SearchModelVersionsRequest) GetOrderBy(ctx context.Context) ([]types.String, bool) {
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

// SetOrderBy sets the value of the OrderBy field in SearchModelVersionsRequest.
func (m *SearchModelVersionsRequest) SetOrderBy(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["order_by"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.OrderBy = types.ListValueMust(t, vs)
}

type SearchModelVersionsResponse struct {
	// Models that match the search criteria
	ModelVersions types.List `tfsdk:"model_versions"`
	// Pagination token to request next page of models for the same search
	// query.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *SearchModelVersionsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SearchModelVersionsResponse) {
	if !from.ModelVersions.IsNull() && !from.ModelVersions.IsUnknown() && to.ModelVersions.IsNull() && len(from.ModelVersions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ModelVersions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ModelVersions = from.ModelVersions
	}
}

func (to *SearchModelVersionsResponse) SyncFieldsDuringRead(ctx context.Context, from SearchModelVersionsResponse) {
	if !from.ModelVersions.IsNull() && !from.ModelVersions.IsUnknown() && to.ModelVersions.IsNull() && len(from.ModelVersions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ModelVersions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ModelVersions = from.ModelVersions
	}
}

func (m SearchModelVersionsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SearchModelVersionsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model_versions": reflect.TypeOf(ModelVersion{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchModelVersionsResponse
// only implements ToObjectValue() and Type().
func (m SearchModelVersionsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_versions":  m.ModelVersions,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SearchModelVersionsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_versions": basetypes.ListType{
				ElemType: ModelVersion{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetModelVersions returns the value of the ModelVersions field in SearchModelVersionsResponse as
// a slice of ModelVersion values.
// If the field is unknown or null, the boolean return value is false.
func (m *SearchModelVersionsResponse) GetModelVersions(ctx context.Context) ([]ModelVersion, bool) {
	if m.ModelVersions.IsNull() || m.ModelVersions.IsUnknown() {
		return nil, false
	}
	var v []ModelVersion
	d := m.ModelVersions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModelVersions sets the value of the ModelVersions field in SearchModelVersionsResponse.
func (m *SearchModelVersionsResponse) SetModelVersions(ctx context.Context, v []ModelVersion) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["model_versions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ModelVersions = types.ListValueMust(t, vs)
}

type SearchModelsRequest struct {
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

func (to *SearchModelsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SearchModelsRequest) {
	if !from.OrderBy.IsNull() && !from.OrderBy.IsUnknown() && to.OrderBy.IsNull() && len(from.OrderBy.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OrderBy, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OrderBy = from.OrderBy
	}
}

func (to *SearchModelsRequest) SyncFieldsDuringRead(ctx context.Context, from SearchModelsRequest) {
	if !from.OrderBy.IsNull() && !from.OrderBy.IsUnknown() && to.OrderBy.IsNull() && len(from.OrderBy.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OrderBy, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OrderBy = from.OrderBy
	}
}

func (m SearchModelsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SearchModelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"order_by": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchModelsRequest
// only implements ToObjectValue() and Type().
func (m SearchModelsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m SearchModelsRequest) Type(ctx context.Context) attr.Type {
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

// GetOrderBy returns the value of the OrderBy field in SearchModelsRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *SearchModelsRequest) GetOrderBy(ctx context.Context) ([]types.String, bool) {
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

// SetOrderBy sets the value of the OrderBy field in SearchModelsRequest.
func (m *SearchModelsRequest) SetOrderBy(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["order_by"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.OrderBy = types.ListValueMust(t, vs)
}

type SearchModelsResponse struct {
	// Pagination token to request the next page of models.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// Registered Models that match the search criteria.
	RegisteredModels types.List `tfsdk:"registered_models"`
}

func (to *SearchModelsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SearchModelsResponse) {
	if !from.RegisteredModels.IsNull() && !from.RegisteredModels.IsUnknown() && to.RegisteredModels.IsNull() && len(from.RegisteredModels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RegisteredModels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RegisteredModels = from.RegisteredModels
	}
}

func (to *SearchModelsResponse) SyncFieldsDuringRead(ctx context.Context, from SearchModelsResponse) {
	if !from.RegisteredModels.IsNull() && !from.RegisteredModels.IsUnknown() && to.RegisteredModels.IsNull() && len(from.RegisteredModels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RegisteredModels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RegisteredModels = from.RegisteredModels
	}
}

func (m SearchModelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SearchModelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"registered_models": reflect.TypeOf(Model{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchModelsResponse
// only implements ToObjectValue() and Type().
func (m SearchModelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":   m.NextPageToken,
			"registered_models": m.RegisteredModels,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SearchModelsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"registered_models": basetypes.ListType{
				ElemType: Model{}.Type(ctx),
			},
		},
	}
}

// GetRegisteredModels returns the value of the RegisteredModels field in SearchModelsResponse as
// a slice of Model values.
// If the field is unknown or null, the boolean return value is false.
func (m *SearchModelsResponse) GetRegisteredModels(ctx context.Context) ([]Model, bool) {
	if m.RegisteredModels.IsNull() || m.RegisteredModels.IsUnknown() {
		return nil, false
	}
	var v []Model
	d := m.RegisteredModels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRegisteredModels sets the value of the RegisteredModels field in SearchModelsResponse.
func (m *SearchModelsResponse) SetRegisteredModels(ctx context.Context, v []Model) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["registered_models"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.RegisteredModels = types.ListValueMust(t, vs)
}

type SearchRuns struct {
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

func (to *SearchRuns) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SearchRuns) {
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

func (to *SearchRuns) SyncFieldsDuringRead(ctx context.Context, from SearchRuns) {
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

func (m SearchRuns) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SearchRuns) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"experiment_ids": reflect.TypeOf(types.String{}),
		"order_by":       reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchRuns
// only implements ToObjectValue() and Type().
func (m SearchRuns) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m SearchRuns) Type(ctx context.Context) attr.Type {
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

// GetExperimentIds returns the value of the ExperimentIds field in SearchRuns as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *SearchRuns) GetExperimentIds(ctx context.Context) ([]types.String, bool) {
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

// SetExperimentIds sets the value of the ExperimentIds field in SearchRuns.
func (m *SearchRuns) SetExperimentIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["experiment_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ExperimentIds = types.ListValueMust(t, vs)
}

// GetOrderBy returns the value of the OrderBy field in SearchRuns as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *SearchRuns) GetOrderBy(ctx context.Context) ([]types.String, bool) {
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

// SetOrderBy sets the value of the OrderBy field in SearchRuns.
func (m *SearchRuns) SetOrderBy(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["order_by"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.OrderBy = types.ListValueMust(t, vs)
}

type SearchRunsResponse struct {
	// Token for the next page of runs.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// Runs that match the search criteria.
	Runs types.List `tfsdk:"runs"`
}

func (to *SearchRunsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SearchRunsResponse) {
	if !from.Runs.IsNull() && !from.Runs.IsUnknown() && to.Runs.IsNull() && len(from.Runs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Runs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Runs = from.Runs
	}
}

func (to *SearchRunsResponse) SyncFieldsDuringRead(ctx context.Context, from SearchRunsResponse) {
	if !from.Runs.IsNull() && !from.Runs.IsUnknown() && to.Runs.IsNull() && len(from.Runs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Runs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Runs = from.Runs
	}
}

func (m SearchRunsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SearchRunsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"runs": reflect.TypeOf(Run{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchRunsResponse
// only implements ToObjectValue() and Type().
func (m SearchRunsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"runs":            m.Runs,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SearchRunsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"runs": basetypes.ListType{
				ElemType: Run{}.Type(ctx),
			},
		},
	}
}

// GetRuns returns the value of the Runs field in SearchRunsResponse as
// a slice of Run values.
// If the field is unknown or null, the boolean return value is false.
func (m *SearchRunsResponse) GetRuns(ctx context.Context) ([]Run, bool) {
	if m.Runs.IsNull() || m.Runs.IsUnknown() {
		return nil, false
	}
	var v []Run
	d := m.Runs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRuns sets the value of the Runs field in SearchRunsResponse.
func (m *SearchRunsResponse) SetRuns(ctx context.Context, v []Run) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["runs"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Runs = types.ListValueMust(t, vs)
}

type SetExperimentTag struct {
	// ID of the experiment under which to log the tag. Must be provided.
	ExperimentId types.String `tfsdk:"experiment_id"`
	// Name of the tag. Keys up to 250 bytes in size are supported.
	Key types.String `tfsdk:"key"`
	// String value of the tag being logged. Values up to 64KB in size are
	// supported.
	Value types.String `tfsdk:"value"`
}

func (to *SetExperimentTag) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SetExperimentTag) {
}

func (to *SetExperimentTag) SyncFieldsDuringRead(ctx context.Context, from SetExperimentTag) {
}

func (m SetExperimentTag) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SetExperimentTag) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetExperimentTag
// only implements ToObjectValue() and Type().
func (m SetExperimentTag) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": m.ExperimentId,
			"key":           m.Key,
			"value":         m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SetExperimentTag) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
			"key":           types.StringType,
			"value":         types.StringType,
		},
	}
}

type SetExperimentTagResponse struct {
}

func (to *SetExperimentTagResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SetExperimentTagResponse) {
}

func (to *SetExperimentTagResponse) SyncFieldsDuringRead(ctx context.Context, from SetExperimentTagResponse) {
}

func (m SetExperimentTagResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetExperimentTagResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SetExperimentTagResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetExperimentTagResponse
// only implements ToObjectValue() and Type().
func (m SetExperimentTagResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m SetExperimentTagResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type SetLoggedModelTagsRequest struct {
	// The ID of the logged model to set the tags on.
	ModelId types.String `tfsdk:"-"`
	// The tags to set on the logged model.
	Tags types.List `tfsdk:"tags"`
}

func (to *SetLoggedModelTagsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SetLoggedModelTagsRequest) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *SetLoggedModelTagsRequest) SyncFieldsDuringRead(ctx context.Context, from SetLoggedModelTagsRequest) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m SetLoggedModelTagsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SetLoggedModelTagsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(LoggedModelTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetLoggedModelTagsRequest
// only implements ToObjectValue() and Type().
func (m SetLoggedModelTagsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_id": m.ModelId,
			"tags":     m.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SetLoggedModelTagsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_id": types.StringType,
			"tags": basetypes.ListType{
				ElemType: LoggedModelTag{}.Type(ctx),
			},
		},
	}
}

// GetTags returns the value of the Tags field in SetLoggedModelTagsRequest as
// a slice of LoggedModelTag values.
// If the field is unknown or null, the boolean return value is false.
func (m *SetLoggedModelTagsRequest) GetTags(ctx context.Context) ([]LoggedModelTag, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []LoggedModelTag
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in SetLoggedModelTagsRequest.
func (m *SetLoggedModelTagsRequest) SetTags(ctx context.Context, v []LoggedModelTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type SetLoggedModelTagsResponse struct {
}

func (to *SetLoggedModelTagsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SetLoggedModelTagsResponse) {
}

func (to *SetLoggedModelTagsResponse) SyncFieldsDuringRead(ctx context.Context, from SetLoggedModelTagsResponse) {
}

func (m SetLoggedModelTagsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetLoggedModelTagsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SetLoggedModelTagsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetLoggedModelTagsResponse
// only implements ToObjectValue() and Type().
func (m SetLoggedModelTagsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m SetLoggedModelTagsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type SetModelTagRequest struct {
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

func (to *SetModelTagRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SetModelTagRequest) {
}

func (to *SetModelTagRequest) SyncFieldsDuringRead(ctx context.Context, from SetModelTagRequest) {
}

func (m SetModelTagRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SetModelTagRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetModelTagRequest
// only implements ToObjectValue() and Type().
func (m SetModelTagRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"name":  m.Name,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SetModelTagRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"name":  types.StringType,
			"value": types.StringType,
		},
	}
}

type SetModelTagResponse struct {
}

func (to *SetModelTagResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SetModelTagResponse) {
}

func (to *SetModelTagResponse) SyncFieldsDuringRead(ctx context.Context, from SetModelTagResponse) {
}

func (m SetModelTagResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetModelTagResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SetModelTagResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetModelTagResponse
// only implements ToObjectValue() and Type().
func (m SetModelTagResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m SetModelTagResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type SetModelVersionTagRequest struct {
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

func (to *SetModelVersionTagRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SetModelVersionTagRequest) {
}

func (to *SetModelVersionTagRequest) SyncFieldsDuringRead(ctx context.Context, from SetModelVersionTagRequest) {
}

func (m SetModelVersionTagRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SetModelVersionTagRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetModelVersionTagRequest
// only implements ToObjectValue() and Type().
func (m SetModelVersionTagRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m SetModelVersionTagRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":     types.StringType,
			"name":    types.StringType,
			"value":   types.StringType,
			"version": types.StringType,
		},
	}
}

type SetModelVersionTagResponse struct {
}

func (to *SetModelVersionTagResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SetModelVersionTagResponse) {
}

func (to *SetModelVersionTagResponse) SyncFieldsDuringRead(ctx context.Context, from SetModelVersionTagResponse) {
}

func (m SetModelVersionTagResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetModelVersionTagResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SetModelVersionTagResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetModelVersionTagResponse
// only implements ToObjectValue() and Type().
func (m SetModelVersionTagResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m SetModelVersionTagResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type SetTag struct {
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

func (to *SetTag) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SetTag) {
}

func (to *SetTag) SyncFieldsDuringRead(ctx context.Context, from SetTag) {
}

func (m SetTag) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SetTag) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetTag
// only implements ToObjectValue() and Type().
func (m SetTag) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m SetTag) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":      types.StringType,
			"run_id":   types.StringType,
			"run_uuid": types.StringType,
			"value":    types.StringType,
		},
	}
}

type SetTagResponse struct {
}

func (to *SetTagResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SetTagResponse) {
}

func (to *SetTagResponse) SyncFieldsDuringRead(ctx context.Context, from SetTagResponse) {
}

func (m SetTagResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetTagResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SetTagResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetTagResponse
// only implements ToObjectValue() and Type().
func (m SetTagResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m SetTagResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Details required to test a registry webhook.
type TestRegistryWebhookRequest struct {
	// If `event` is specified, the test trigger uses the specified event. If
	// `event` is not specified, the test trigger uses a randomly chosen event
	// associated with the webhook.
	Event types.String `tfsdk:"event"`
	// Webhook ID
	Id types.String `tfsdk:"id"`
}

func (to *TestRegistryWebhookRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TestRegistryWebhookRequest) {
}

func (to *TestRegistryWebhookRequest) SyncFieldsDuringRead(ctx context.Context, from TestRegistryWebhookRequest) {
}

func (m TestRegistryWebhookRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m TestRegistryWebhookRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TestRegistryWebhookRequest
// only implements ToObjectValue() and Type().
func (m TestRegistryWebhookRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"event": m.Event,
			"id":    m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TestRegistryWebhookRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"event": types.StringType,
			"id":    types.StringType,
		},
	}
}

type TestRegistryWebhookResponse struct {
	// Body of the response from the webhook URL
	Body types.String `tfsdk:"body"`
	// Status code returned by the webhook URL
	StatusCode types.Int64 `tfsdk:"status_code"`
}

func (to *TestRegistryWebhookResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TestRegistryWebhookResponse) {
}

func (to *TestRegistryWebhookResponse) SyncFieldsDuringRead(ctx context.Context, from TestRegistryWebhookResponse) {
}

func (m TestRegistryWebhookResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m TestRegistryWebhookResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TestRegistryWebhookResponse
// only implements ToObjectValue() and Type().
func (m TestRegistryWebhookResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"body":        m.Body,
			"status_code": m.StatusCode,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TestRegistryWebhookResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"body":        types.StringType,
			"status_code": types.Int64Type,
		},
	}
}

type TimeWindow struct {
	// The duration of the time window.
	Duration types.String `tfsdk:"duration"`
	// The offset of the time window.
	Offset types.String `tfsdk:"offset"`
}

func (to *TimeWindow) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TimeWindow) {
}

func (to *TimeWindow) SyncFieldsDuringRead(ctx context.Context, from TimeWindow) {
}

func (m TimeWindow) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m TimeWindow) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TimeWindow
// only implements ToObjectValue() and Type().
func (m TimeWindow) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"duration": m.Duration,
			"offset":   m.Offset,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TimeWindow) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"duration": types.StringType,
			"offset":   types.StringType,
		},
	}
}

// Details required to transition a model version's stage.
type TransitionModelVersionStageDatabricks struct {
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

func (to *TransitionModelVersionStageDatabricks) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TransitionModelVersionStageDatabricks) {
}

func (to *TransitionModelVersionStageDatabricks) SyncFieldsDuringRead(ctx context.Context, from TransitionModelVersionStageDatabricks) {
}

func (m TransitionModelVersionStageDatabricks) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m TransitionModelVersionStageDatabricks) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TransitionModelVersionStageDatabricks
// only implements ToObjectValue() and Type().
func (m TransitionModelVersionStageDatabricks) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m TransitionModelVersionStageDatabricks) Type(ctx context.Context) attr.Type {
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
type TransitionRequest struct {
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

func (to *TransitionRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TransitionRequest) {
	if !from.AvailableActions.IsNull() && !from.AvailableActions.IsUnknown() && to.AvailableActions.IsNull() && len(from.AvailableActions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AvailableActions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AvailableActions = from.AvailableActions
	}
}

func (to *TransitionRequest) SyncFieldsDuringRead(ctx context.Context, from TransitionRequest) {
	if !from.AvailableActions.IsNull() && !from.AvailableActions.IsUnknown() && to.AvailableActions.IsNull() && len(from.AvailableActions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AvailableActions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AvailableActions = from.AvailableActions
	}
}

func (m TransitionRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m TransitionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"available_actions": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TransitionRequest
// only implements ToObjectValue() and Type().
func (m TransitionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m TransitionRequest) Type(ctx context.Context) attr.Type {
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

// GetAvailableActions returns the value of the AvailableActions field in TransitionRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *TransitionRequest) GetAvailableActions(ctx context.Context) ([]types.String, bool) {
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

// SetAvailableActions sets the value of the AvailableActions field in TransitionRequest.
func (m *TransitionRequest) SetAvailableActions(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["available_actions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AvailableActions = types.ListValueMust(t, vs)
}

type TransitionStageResponse struct {
	// Updated model version
	ModelVersionDatabricks types.Object `tfsdk:"model_version_databricks"`
}

func (to *TransitionStageResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TransitionStageResponse) {
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

func (to *TransitionStageResponse) SyncFieldsDuringRead(ctx context.Context, from TransitionStageResponse) {
	if !from.ModelVersionDatabricks.IsNull() && !from.ModelVersionDatabricks.IsUnknown() {
		if toModelVersionDatabricks, ok := to.GetModelVersionDatabricks(ctx); ok {
			if fromModelVersionDatabricks, ok := from.GetModelVersionDatabricks(ctx); ok {
				toModelVersionDatabricks.SyncFieldsDuringRead(ctx, fromModelVersionDatabricks)
				to.SetModelVersionDatabricks(ctx, toModelVersionDatabricks)
			}
		}
	}
}

func (m TransitionStageResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["model_version_databricks"] = attrs["model_version_databricks"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TransitionStageResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m TransitionStageResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model_version_databricks": reflect.TypeOf(ModelVersionDatabricks{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TransitionStageResponse
// only implements ToObjectValue() and Type().
func (m TransitionStageResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_version_databricks": m.ModelVersionDatabricks,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TransitionStageResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_version_databricks": ModelVersionDatabricks{}.Type(ctx),
		},
	}
}

// GetModelVersionDatabricks returns the value of the ModelVersionDatabricks field in TransitionStageResponse as
// a ModelVersionDatabricks value.
// If the field is unknown or null, the boolean return value is false.
func (m *TransitionStageResponse) GetModelVersionDatabricks(ctx context.Context) (ModelVersionDatabricks, bool) {
	var e ModelVersionDatabricks
	if m.ModelVersionDatabricks.IsNull() || m.ModelVersionDatabricks.IsUnknown() {
		return e, false
	}
	var v ModelVersionDatabricks
	d := m.ModelVersionDatabricks.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModelVersionDatabricks sets the value of the ModelVersionDatabricks field in TransitionStageResponse.
func (m *TransitionStageResponse) SetModelVersionDatabricks(ctx context.Context, v ModelVersionDatabricks) {
	vs := v.ToObjectValue(ctx)
	m.ModelVersionDatabricks = vs
}

// Details required to edit a comment on a model version.
type UpdateComment struct {
	// User-provided comment on the action.
	Comment types.String `tfsdk:"comment"`
	// Unique identifier of an activity
	Id types.String `tfsdk:"id"`
}

func (to *UpdateComment) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateComment) {
}

func (to *UpdateComment) SyncFieldsDuringRead(ctx context.Context, from UpdateComment) {
}

func (m UpdateComment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UpdateComment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateComment
// only implements ToObjectValue() and Type().
func (m UpdateComment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment": m.Comment,
			"id":      m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateComment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment": types.StringType,
			"id":      types.StringType,
		},
	}
}

type UpdateCommentResponse struct {
	// Updated comment object
	Comment types.Object `tfsdk:"comment"`
}

func (to *UpdateCommentResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateCommentResponse) {
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

func (to *UpdateCommentResponse) SyncFieldsDuringRead(ctx context.Context, from UpdateCommentResponse) {
	if !from.Comment.IsNull() && !from.Comment.IsUnknown() {
		if toComment, ok := to.GetComment(ctx); ok {
			if fromComment, ok := from.GetComment(ctx); ok {
				toComment.SyncFieldsDuringRead(ctx, fromComment)
				to.SetComment(ctx, toComment)
			}
		}
	}
}

func (m UpdateCommentResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCommentResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateCommentResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"comment": reflect.TypeOf(CommentObject{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCommentResponse
// only implements ToObjectValue() and Type().
func (m UpdateCommentResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment": m.Comment,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateCommentResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment": CommentObject{}.Type(ctx),
		},
	}
}

// GetComment returns the value of the Comment field in UpdateCommentResponse as
// a CommentObject value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateCommentResponse) GetComment(ctx context.Context) (CommentObject, bool) {
	var e CommentObject
	if m.Comment.IsNull() || m.Comment.IsUnknown() {
		return e, false
	}
	var v CommentObject
	d := m.Comment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetComment sets the value of the Comment field in UpdateCommentResponse.
func (m *UpdateCommentResponse) SetComment(ctx context.Context, v CommentObject) {
	vs := v.ToObjectValue(ctx)
	m.Comment = vs
}

type UpdateExperiment struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"experiment_id"`
	// If provided, the experiment's name is changed to the new name. The new
	// name must be unique.
	NewName types.String `tfsdk:"new_name"`
}

func (to *UpdateExperiment) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateExperiment) {
}

func (to *UpdateExperiment) SyncFieldsDuringRead(ctx context.Context, from UpdateExperiment) {
}

func (m UpdateExperiment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UpdateExperiment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateExperiment
// only implements ToObjectValue() and Type().
func (m UpdateExperiment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": m.ExperimentId,
			"new_name":      m.NewName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateExperiment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
			"new_name":      types.StringType,
		},
	}
}

type UpdateExperimentResponse struct {
}

func (to *UpdateExperimentResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateExperimentResponse) {
}

func (to *UpdateExperimentResponse) SyncFieldsDuringRead(ctx context.Context, from UpdateExperimentResponse) {
}

func (m UpdateExperimentResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateExperimentResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateExperimentResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateExperimentResponse
// only implements ToObjectValue() and Type().
func (m UpdateExperimentResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateExperimentResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateFeatureRequest struct {
	// Feature to update.
	Feature types.Object `tfsdk:"feature"`
	// The full three-part name (catalog, schema, name) of the feature.
	FullName types.String `tfsdk:"-"`
	// The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateFeatureRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateFeatureRequest) {
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

func (to *UpdateFeatureRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateFeatureRequest) {
	if !from.Feature.IsNull() && !from.Feature.IsUnknown() {
		if toFeature, ok := to.GetFeature(ctx); ok {
			if fromFeature, ok := from.GetFeature(ctx); ok {
				toFeature.SyncFieldsDuringRead(ctx, fromFeature)
				to.SetFeature(ctx, toFeature)
			}
		}
	}
}

func (m UpdateFeatureRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["feature"] = attrs["feature"].SetRequired()
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
func (m UpdateFeatureRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"feature": reflect.TypeOf(Feature{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateFeatureRequest
// only implements ToObjectValue() and Type().
func (m UpdateFeatureRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature":     m.Feature,
			"full_name":   m.FullName,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateFeatureRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"feature":     Feature{}.Type(ctx),
			"full_name":   types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetFeature returns the value of the Feature field in UpdateFeatureRequest as
// a Feature value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateFeatureRequest) GetFeature(ctx context.Context) (Feature, bool) {
	var e Feature
	if m.Feature.IsNull() || m.Feature.IsUnknown() {
		return e, false
	}
	var v Feature
	d := m.Feature.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFeature sets the value of the Feature field in UpdateFeatureRequest.
func (m *UpdateFeatureRequest) SetFeature(ctx context.Context, v Feature) {
	vs := v.ToObjectValue(ctx)
	m.Feature = vs
}

type UpdateFeatureTagRequest struct {
	FeatureName types.String `tfsdk:"-"`

	FeatureTag types.Object `tfsdk:"feature_tag"`

	Key types.String `tfsdk:"-"`

	TableName types.String `tfsdk:"-"`
	// The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateFeatureTagRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateFeatureTagRequest) {
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

func (to *UpdateFeatureTagRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateFeatureTagRequest) {
	if !from.FeatureTag.IsNull() && !from.FeatureTag.IsUnknown() {
		if toFeatureTag, ok := to.GetFeatureTag(ctx); ok {
			if fromFeatureTag, ok := from.GetFeatureTag(ctx); ok {
				toFeatureTag.SyncFieldsDuringRead(ctx, fromFeatureTag)
				to.SetFeatureTag(ctx, toFeatureTag)
			}
		}
	}
}

func (m UpdateFeatureTagRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["feature_tag"] = attrs["feature_tag"].SetRequired()
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
func (m UpdateFeatureTagRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"feature_tag": reflect.TypeOf(FeatureTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateFeatureTagRequest
// only implements ToObjectValue() and Type().
func (m UpdateFeatureTagRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m UpdateFeatureTagRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"feature_name": types.StringType,
			"feature_tag":  FeatureTag{}.Type(ctx),
			"key":          types.StringType,
			"table_name":   types.StringType,
			"update_mask":  types.StringType,
		},
	}
}

// GetFeatureTag returns the value of the FeatureTag field in UpdateFeatureTagRequest as
// a FeatureTag value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateFeatureTagRequest) GetFeatureTag(ctx context.Context) (FeatureTag, bool) {
	var e FeatureTag
	if m.FeatureTag.IsNull() || m.FeatureTag.IsUnknown() {
		return e, false
	}
	var v FeatureTag
	d := m.FeatureTag.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFeatureTag sets the value of the FeatureTag field in UpdateFeatureTagRequest.
func (m *UpdateFeatureTagRequest) SetFeatureTag(ctx context.Context, v FeatureTag) {
	vs := v.ToObjectValue(ctx)
	m.FeatureTag = vs
}

type UpdateModelRequest struct {
	// If provided, updates the description for this `registered_model`.
	Description types.String `tfsdk:"description"`
	// Registered model unique name identifier.
	Name types.String `tfsdk:"name"`
}

func (to *UpdateModelRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateModelRequest) {
}

func (to *UpdateModelRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateModelRequest) {
}

func (m UpdateModelRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UpdateModelRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateModelRequest
// only implements ToObjectValue() and Type().
func (m UpdateModelRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description": m.Description,
			"name":        m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateModelRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"name":        types.StringType,
		},
	}
}

type UpdateModelResponse struct {
	RegisteredModel types.Object `tfsdk:"registered_model"`
}

func (to *UpdateModelResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateModelResponse) {
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

func (to *UpdateModelResponse) SyncFieldsDuringRead(ctx context.Context, from UpdateModelResponse) {
	if !from.RegisteredModel.IsNull() && !from.RegisteredModel.IsUnknown() {
		if toRegisteredModel, ok := to.GetRegisteredModel(ctx); ok {
			if fromRegisteredModel, ok := from.GetRegisteredModel(ctx); ok {
				toRegisteredModel.SyncFieldsDuringRead(ctx, fromRegisteredModel)
				to.SetRegisteredModel(ctx, toRegisteredModel)
			}
		}
	}
}

func (m UpdateModelResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["registered_model"] = attrs["registered_model"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateModelResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateModelResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"registered_model": reflect.TypeOf(Model{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateModelResponse
// only implements ToObjectValue() and Type().
func (m UpdateModelResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"registered_model": m.RegisteredModel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateModelResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"registered_model": Model{}.Type(ctx),
		},
	}
}

// GetRegisteredModel returns the value of the RegisteredModel field in UpdateModelResponse as
// a Model value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateModelResponse) GetRegisteredModel(ctx context.Context) (Model, bool) {
	var e Model
	if m.RegisteredModel.IsNull() || m.RegisteredModel.IsUnknown() {
		return e, false
	}
	var v Model
	d := m.RegisteredModel.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRegisteredModel sets the value of the RegisteredModel field in UpdateModelResponse.
func (m *UpdateModelResponse) SetRegisteredModel(ctx context.Context, v Model) {
	vs := v.ToObjectValue(ctx)
	m.RegisteredModel = vs
}

type UpdateModelVersionRequest struct {
	// If provided, updates the description for this `registered_model`.
	Description types.String `tfsdk:"description"`
	// Name of the registered model
	Name types.String `tfsdk:"name"`
	// Model version number
	Version types.String `tfsdk:"version"`
}

func (to *UpdateModelVersionRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateModelVersionRequest) {
}

func (to *UpdateModelVersionRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateModelVersionRequest) {
}

func (m UpdateModelVersionRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UpdateModelVersionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateModelVersionRequest
// only implements ToObjectValue() and Type().
func (m UpdateModelVersionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description": m.Description,
			"name":        m.Name,
			"version":     m.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateModelVersionRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"name":        types.StringType,
			"version":     types.StringType,
		},
	}
}

type UpdateModelVersionResponse struct {
	// Return new version number generated for this model in registry.
	ModelVersion types.Object `tfsdk:"model_version"`
}

func (to *UpdateModelVersionResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateModelVersionResponse) {
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

func (to *UpdateModelVersionResponse) SyncFieldsDuringRead(ctx context.Context, from UpdateModelVersionResponse) {
	if !from.ModelVersion.IsNull() && !from.ModelVersion.IsUnknown() {
		if toModelVersion, ok := to.GetModelVersion(ctx); ok {
			if fromModelVersion, ok := from.GetModelVersion(ctx); ok {
				toModelVersion.SyncFieldsDuringRead(ctx, fromModelVersion)
				to.SetModelVersion(ctx, toModelVersion)
			}
		}
	}
}

func (m UpdateModelVersionResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["model_version"] = attrs["model_version"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateModelVersionResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateModelVersionResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model_version": reflect.TypeOf(ModelVersion{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateModelVersionResponse
// only implements ToObjectValue() and Type().
func (m UpdateModelVersionResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_version": m.ModelVersion,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateModelVersionResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_version": ModelVersion{}.Type(ctx),
		},
	}
}

// GetModelVersion returns the value of the ModelVersion field in UpdateModelVersionResponse as
// a ModelVersion value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateModelVersionResponse) GetModelVersion(ctx context.Context) (ModelVersion, bool) {
	var e ModelVersion
	if m.ModelVersion.IsNull() || m.ModelVersion.IsUnknown() {
		return e, false
	}
	var v ModelVersion
	d := m.ModelVersion.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModelVersion sets the value of the ModelVersion field in UpdateModelVersionResponse.
func (m *UpdateModelVersionResponse) SetModelVersion(ctx context.Context, v ModelVersion) {
	vs := v.ToObjectValue(ctx)
	m.ModelVersion = vs
}

type UpdateOnlineStoreRequest struct {
	// The name of the online store. This is the unique identifier for the
	// online store.
	Name types.String `tfsdk:"-"`
	// Online store to update.
	OnlineStore types.Object `tfsdk:"online_store"`
	// The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateOnlineStoreRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateOnlineStoreRequest) {
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

func (to *UpdateOnlineStoreRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateOnlineStoreRequest) {
	if !from.OnlineStore.IsNull() && !from.OnlineStore.IsUnknown() {
		if toOnlineStore, ok := to.GetOnlineStore(ctx); ok {
			if fromOnlineStore, ok := from.GetOnlineStore(ctx); ok {
				toOnlineStore.SyncFieldsDuringRead(ctx, fromOnlineStore)
				to.SetOnlineStore(ctx, toOnlineStore)
			}
		}
	}
}

func (m UpdateOnlineStoreRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["online_store"] = attrs["online_store"].SetRequired()
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
func (m UpdateOnlineStoreRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"online_store": reflect.TypeOf(OnlineStore{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateOnlineStoreRequest
// only implements ToObjectValue() and Type().
func (m UpdateOnlineStoreRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":         m.Name,
			"online_store": m.OnlineStore,
			"update_mask":  m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateOnlineStoreRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":         types.StringType,
			"online_store": OnlineStore{}.Type(ctx),
			"update_mask":  types.StringType,
		},
	}
}

// GetOnlineStore returns the value of the OnlineStore field in UpdateOnlineStoreRequest as
// a OnlineStore value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateOnlineStoreRequest) GetOnlineStore(ctx context.Context) (OnlineStore, bool) {
	var e OnlineStore
	if m.OnlineStore.IsNull() || m.OnlineStore.IsUnknown() {
		return e, false
	}
	var v OnlineStore
	d := m.OnlineStore.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOnlineStore sets the value of the OnlineStore field in UpdateOnlineStoreRequest.
func (m *UpdateOnlineStoreRequest) SetOnlineStore(ctx context.Context, v OnlineStore) {
	vs := v.ToObjectValue(ctx)
	m.OnlineStore = vs
}

// Details required to update a registry webhook. Only the fields that need to
// be updated should be specified, and both `http_url_spec` and `job_spec`
// should not be specified in the same request.
type UpdateRegistryWebhook struct {
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

	HttpUrlSpec types.Object `tfsdk:"http_url_spec"`
	// Webhook ID
	Id types.String `tfsdk:"id"`

	JobSpec types.Object `tfsdk:"job_spec"`

	Status types.String `tfsdk:"status"`
}

func (to *UpdateRegistryWebhook) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateRegistryWebhook) {
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

func (to *UpdateRegistryWebhook) SyncFieldsDuringRead(ctx context.Context, from UpdateRegistryWebhook) {
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

func (m UpdateRegistryWebhook) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["events"] = attrs["events"].SetOptional()
	attrs["http_url_spec"] = attrs["http_url_spec"].SetOptional()
	attrs["id"] = attrs["id"].SetRequired()
	attrs["job_spec"] = attrs["job_spec"].SetOptional()
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
func (m UpdateRegistryWebhook) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"events":        reflect.TypeOf(types.String{}),
		"http_url_spec": reflect.TypeOf(HttpUrlSpec{}),
		"job_spec":      reflect.TypeOf(JobSpec{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRegistryWebhook
// only implements ToObjectValue() and Type().
func (m UpdateRegistryWebhook) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m UpdateRegistryWebhook) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"events": basetypes.ListType{
				ElemType: types.StringType,
			},
			"http_url_spec": HttpUrlSpec{}.Type(ctx),
			"id":            types.StringType,
			"job_spec":      JobSpec{}.Type(ctx),
			"status":        types.StringType,
		},
	}
}

// GetEvents returns the value of the Events field in UpdateRegistryWebhook as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateRegistryWebhook) GetEvents(ctx context.Context) ([]types.String, bool) {
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

// SetEvents sets the value of the Events field in UpdateRegistryWebhook.
func (m *UpdateRegistryWebhook) SetEvents(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["events"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Events = types.ListValueMust(t, vs)
}

// GetHttpUrlSpec returns the value of the HttpUrlSpec field in UpdateRegistryWebhook as
// a HttpUrlSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateRegistryWebhook) GetHttpUrlSpec(ctx context.Context) (HttpUrlSpec, bool) {
	var e HttpUrlSpec
	if m.HttpUrlSpec.IsNull() || m.HttpUrlSpec.IsUnknown() {
		return e, false
	}
	var v HttpUrlSpec
	d := m.HttpUrlSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetHttpUrlSpec sets the value of the HttpUrlSpec field in UpdateRegistryWebhook.
func (m *UpdateRegistryWebhook) SetHttpUrlSpec(ctx context.Context, v HttpUrlSpec) {
	vs := v.ToObjectValue(ctx)
	m.HttpUrlSpec = vs
}

// GetJobSpec returns the value of the JobSpec field in UpdateRegistryWebhook as
// a JobSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateRegistryWebhook) GetJobSpec(ctx context.Context) (JobSpec, bool) {
	var e JobSpec
	if m.JobSpec.IsNull() || m.JobSpec.IsUnknown() {
		return e, false
	}
	var v JobSpec
	d := m.JobSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetJobSpec sets the value of the JobSpec field in UpdateRegistryWebhook.
func (m *UpdateRegistryWebhook) SetJobSpec(ctx context.Context, v JobSpec) {
	vs := v.ToObjectValue(ctx)
	m.JobSpec = vs
}

type UpdateRun struct {
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

func (to *UpdateRun) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateRun) {
}

func (to *UpdateRun) SyncFieldsDuringRead(ctx context.Context, from UpdateRun) {
}

func (m UpdateRun) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UpdateRun) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRun
// only implements ToObjectValue() and Type().
func (m UpdateRun) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m UpdateRun) Type(ctx context.Context) attr.Type {
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

type UpdateRunResponse struct {
	// Updated metadata of the run.
	RunInfo types.Object `tfsdk:"run_info"`
}

func (to *UpdateRunResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateRunResponse) {
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

func (to *UpdateRunResponse) SyncFieldsDuringRead(ctx context.Context, from UpdateRunResponse) {
	if !from.RunInfo.IsNull() && !from.RunInfo.IsUnknown() {
		if toRunInfo, ok := to.GetRunInfo(ctx); ok {
			if fromRunInfo, ok := from.GetRunInfo(ctx); ok {
				toRunInfo.SyncFieldsDuringRead(ctx, fromRunInfo)
				to.SetRunInfo(ctx, toRunInfo)
			}
		}
	}
}

func (m UpdateRunResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["run_info"] = attrs["run_info"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateRunResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateRunResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"run_info": reflect.TypeOf(RunInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRunResponse
// only implements ToObjectValue() and Type().
func (m UpdateRunResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run_info": m.RunInfo,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateRunResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run_info": RunInfo{}.Type(ctx),
		},
	}
}

// GetRunInfo returns the value of the RunInfo field in UpdateRunResponse as
// a RunInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateRunResponse) GetRunInfo(ctx context.Context) (RunInfo, bool) {
	var e RunInfo
	if m.RunInfo.IsNull() || m.RunInfo.IsUnknown() {
		return e, false
	}
	var v RunInfo
	d := m.RunInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRunInfo sets the value of the RunInfo field in UpdateRunResponse.
func (m *UpdateRunResponse) SetRunInfo(ctx context.Context, v RunInfo) {
	vs := v.ToObjectValue(ctx)
	m.RunInfo = vs
}

type UpdateWebhookResponse struct {
	Webhook types.Object `tfsdk:"webhook"`
}

func (to *UpdateWebhookResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateWebhookResponse) {
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

func (to *UpdateWebhookResponse) SyncFieldsDuringRead(ctx context.Context, from UpdateWebhookResponse) {
	if !from.Webhook.IsNull() && !from.Webhook.IsUnknown() {
		if toWebhook, ok := to.GetWebhook(ctx); ok {
			if fromWebhook, ok := from.GetWebhook(ctx); ok {
				toWebhook.SyncFieldsDuringRead(ctx, fromWebhook)
				to.SetWebhook(ctx, toWebhook)
			}
		}
	}
}

func (m UpdateWebhookResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["webhook"] = attrs["webhook"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateWebhookResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateWebhookResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"webhook": reflect.TypeOf(RegistryWebhook{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWebhookResponse
// only implements ToObjectValue() and Type().
func (m UpdateWebhookResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"webhook": m.Webhook,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateWebhookResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"webhook": RegistryWebhook{}.Type(ctx),
		},
	}
}

// GetWebhook returns the value of the Webhook field in UpdateWebhookResponse as
// a RegistryWebhook value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateWebhookResponse) GetWebhook(ctx context.Context) (RegistryWebhook, bool) {
	var e RegistryWebhook
	if m.Webhook.IsNull() || m.Webhook.IsUnknown() {
		return e, false
	}
	var v RegistryWebhook
	d := m.Webhook.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWebhook sets the value of the Webhook field in UpdateWebhookResponse.
func (m *UpdateWebhookResponse) SetWebhook(ctx context.Context, v RegistryWebhook) {
	vs := v.ToObjectValue(ctx)
	m.Webhook = vs
}
