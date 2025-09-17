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

func (toState *Activity_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Activity_SdkV2) {
}

func (toState *Activity_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState Activity_SdkV2) {
}

func (c Activity_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Activity_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Activity_SdkV2
// only implements ToObjectValue() and Type().
func (o Activity_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"activity_type":          o.ActivityType,
			"comment":                o.Comment,
			"creation_timestamp":     o.CreationTimestamp,
			"from_stage":             o.FromStage,
			"id":                     o.Id,
			"last_updated_timestamp": o.LastUpdatedTimestamp,
			"system_comment":         o.SystemComment,
			"to_stage":               o.ToStage,
			"user_id":                o.UserId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Activity_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *ApproveTransitionRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ApproveTransitionRequest_SdkV2) {
}

func (toState *ApproveTransitionRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ApproveTransitionRequest_SdkV2) {
}

func (c ApproveTransitionRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ApproveTransitionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ApproveTransitionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ApproveTransitionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"archive_existing_versions": o.ArchiveExistingVersions,
			"comment":                   o.Comment,
			"name":                      o.Name,
			"stage":                     o.Stage,
			"version":                   o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ApproveTransitionRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *ApproveTransitionRequestResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ApproveTransitionRequestResponse_SdkV2) {
	if !fromPlan.Activity.IsNull() && !fromPlan.Activity.IsUnknown() {
		if toStateActivity, ok := toState.GetActivity(ctx); ok {
			if fromPlanActivity, ok := fromPlan.GetActivity(ctx); ok {
				toStateActivity.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanActivity)
				toState.SetActivity(ctx, toStateActivity)
			}
		}
	}
}

func (toState *ApproveTransitionRequestResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ApproveTransitionRequestResponse_SdkV2) {
	if !fromState.Activity.IsNull() && !fromState.Activity.IsUnknown() {
		if toStateActivity, ok := toState.GetActivity(ctx); ok {
			if fromStateActivity, ok := fromState.GetActivity(ctx); ok {
				toStateActivity.SyncFieldsDuringRead(ctx, fromStateActivity)
				toState.SetActivity(ctx, toStateActivity)
			}
		}
	}
}

func (c ApproveTransitionRequestResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ApproveTransitionRequestResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"activity": reflect.TypeOf(Activity_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ApproveTransitionRequestResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ApproveTransitionRequestResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"activity": o.Activity,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ApproveTransitionRequestResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ApproveTransitionRequestResponse_SdkV2) GetActivity(ctx context.Context) (Activity_SdkV2, bool) {
	var e Activity_SdkV2
	if o.Activity.IsNull() || o.Activity.IsUnknown() {
		return e, false
	}
	var v []Activity_SdkV2
	d := o.Activity.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetActivity sets the value of the Activity field in ApproveTransitionRequestResponse_SdkV2.
func (o *ApproveTransitionRequestResponse_SdkV2) SetActivity(ctx context.Context, v Activity_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["activity"]
	o.Activity = types.ListValueMust(t, vs)
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

func (toState *CommentObject_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CommentObject_SdkV2) {
}

func (toState *CommentObject_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CommentObject_SdkV2) {
}

func (c CommentObject_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CommentObject_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"available_actions": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CommentObject_SdkV2
// only implements ToObjectValue() and Type().
func (o CommentObject_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"available_actions":      o.AvailableActions,
			"comment":                o.Comment,
			"creation_timestamp":     o.CreationTimestamp,
			"id":                     o.Id,
			"last_updated_timestamp": o.LastUpdatedTimestamp,
			"user_id":                o.UserId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CommentObject_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *CommentObject_SdkV2) GetAvailableActions(ctx context.Context) ([]types.String, bool) {
	if o.AvailableActions.IsNull() || o.AvailableActions.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.AvailableActions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAvailableActions sets the value of the AvailableActions field in CommentObject_SdkV2.
func (o *CommentObject_SdkV2) SetAvailableActions(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["available_actions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AvailableActions = types.ListValueMust(t, vs)
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

func (toState *CreateComment_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateComment_SdkV2) {
}

func (toState *CreateComment_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CreateComment_SdkV2) {
}

func (c CreateComment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateComment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateComment_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateComment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment": o.Comment,
			"name":    o.Name,
			"version": o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateComment_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *CreateCommentResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateCommentResponse_SdkV2) {
	if !fromPlan.Comment.IsNull() && !fromPlan.Comment.IsUnknown() {
		if toStateComment, ok := toState.GetComment(ctx); ok {
			if fromPlanComment, ok := fromPlan.GetComment(ctx); ok {
				toStateComment.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanComment)
				toState.SetComment(ctx, toStateComment)
			}
		}
	}
}

func (toState *CreateCommentResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CreateCommentResponse_SdkV2) {
	if !fromState.Comment.IsNull() && !fromState.Comment.IsUnknown() {
		if toStateComment, ok := toState.GetComment(ctx); ok {
			if fromStateComment, ok := fromState.GetComment(ctx); ok {
				toStateComment.SyncFieldsDuringRead(ctx, fromStateComment)
				toState.SetComment(ctx, toStateComment)
			}
		}
	}
}

func (c CreateCommentResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateCommentResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"comment": reflect.TypeOf(CommentObject_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCommentResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateCommentResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment": o.Comment,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCommentResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *CreateCommentResponse_SdkV2) GetComment(ctx context.Context) (CommentObject_SdkV2, bool) {
	var e CommentObject_SdkV2
	if o.Comment.IsNull() || o.Comment.IsUnknown() {
		return e, false
	}
	var v []CommentObject_SdkV2
	d := o.Comment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetComment sets the value of the Comment field in CreateCommentResponse_SdkV2.
func (o *CreateCommentResponse_SdkV2) SetComment(ctx context.Context, v CommentObject_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["comment"]
	o.Comment = types.ListValueMust(t, vs)
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

func (toState *CreateExperiment_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateExperiment_SdkV2) {
}

func (toState *CreateExperiment_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CreateExperiment_SdkV2) {
}

func (c CreateExperiment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateExperiment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(ExperimentTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateExperiment_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateExperiment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"artifact_location": o.ArtifactLocation,
			"name":              o.Name,
			"tags":              o.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateExperiment_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *CreateExperiment_SdkV2) GetTags(ctx context.Context) ([]ExperimentTag_SdkV2, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []ExperimentTag_SdkV2
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in CreateExperiment_SdkV2.
func (o *CreateExperiment_SdkV2) SetTags(ctx context.Context, v []ExperimentTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type CreateExperimentResponse_SdkV2 struct {
	// Unique identifier for the experiment.
	ExperimentId types.String `tfsdk:"experiment_id"`
}

func (toState *CreateExperimentResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateExperimentResponse_SdkV2) {
}

func (toState *CreateExperimentResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CreateExperimentResponse_SdkV2) {
}

func (c CreateExperimentResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateExperimentResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateExperimentResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateExperimentResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": o.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateExperimentResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
		},
	}
}

type CreateFeatureTagRequest_SdkV2 struct {
	FeatureName types.String `tfsdk:"-"`

	FeatureTag types.List `tfsdk:"feature_tag"`

	TableName types.String `tfsdk:"-"`
}

func (toState *CreateFeatureTagRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateFeatureTagRequest_SdkV2) {
	if !fromPlan.FeatureTag.IsNull() && !fromPlan.FeatureTag.IsUnknown() {
		if toStateFeatureTag, ok := toState.GetFeatureTag(ctx); ok {
			if fromPlanFeatureTag, ok := fromPlan.GetFeatureTag(ctx); ok {
				toStateFeatureTag.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanFeatureTag)
				toState.SetFeatureTag(ctx, toStateFeatureTag)
			}
		}
	}
}

func (toState *CreateFeatureTagRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CreateFeatureTagRequest_SdkV2) {
	if !fromState.FeatureTag.IsNull() && !fromState.FeatureTag.IsUnknown() {
		if toStateFeatureTag, ok := toState.GetFeatureTag(ctx); ok {
			if fromStateFeatureTag, ok := fromState.GetFeatureTag(ctx); ok {
				toStateFeatureTag.SyncFieldsDuringRead(ctx, fromStateFeatureTag)
				toState.SetFeatureTag(ctx, toStateFeatureTag)
			}
		}
	}
}

func (c CreateFeatureTagRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateFeatureTagRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"feature_tag": reflect.TypeOf(FeatureTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateFeatureTagRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateFeatureTagRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_name": o.FeatureName,
			"feature_tag":  o.FeatureTag,
			"table_name":   o.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateFeatureTagRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *CreateFeatureTagRequest_SdkV2) GetFeatureTag(ctx context.Context) (FeatureTag_SdkV2, bool) {
	var e FeatureTag_SdkV2
	if o.FeatureTag.IsNull() || o.FeatureTag.IsUnknown() {
		return e, false
	}
	var v []FeatureTag_SdkV2
	d := o.FeatureTag.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFeatureTag sets the value of the FeatureTag field in CreateFeatureTagRequest_SdkV2.
func (o *CreateFeatureTagRequest_SdkV2) SetFeatureTag(ctx context.Context, v FeatureTag_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["feature_tag"]
	o.FeatureTag = types.ListValueMust(t, vs)
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

func (toState *CreateForecastingExperimentRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateForecastingExperimentRequest_SdkV2) {
}

func (toState *CreateForecastingExperimentRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CreateForecastingExperimentRequest_SdkV2) {
}

func (c CreateForecastingExperimentRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateForecastingExperimentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (o CreateForecastingExperimentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"custom_weights_column":         o.CustomWeightsColumn,
			"experiment_path":               o.ExperimentPath,
			"forecast_granularity":          o.ForecastGranularity,
			"forecast_horizon":              o.ForecastHorizon,
			"future_feature_data_path":      o.FutureFeatureDataPath,
			"holiday_regions":               o.HolidayRegions,
			"include_features":              o.IncludeFeatures,
			"max_runtime":                   o.MaxRuntime,
			"prediction_data_path":          o.PredictionDataPath,
			"primary_metric":                o.PrimaryMetric,
			"register_to":                   o.RegisterTo,
			"split_column":                  o.SplitColumn,
			"target_column":                 o.TargetColumn,
			"time_column":                   o.TimeColumn,
			"timeseries_identifier_columns": o.TimeseriesIdentifierColumns,
			"train_data_path":               o.TrainDataPath,
			"training_frameworks":           o.TrainingFrameworks,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateForecastingExperimentRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *CreateForecastingExperimentRequest_SdkV2) GetHolidayRegions(ctx context.Context) ([]types.String, bool) {
	if o.HolidayRegions.IsNull() || o.HolidayRegions.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.HolidayRegions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetHolidayRegions sets the value of the HolidayRegions field in CreateForecastingExperimentRequest_SdkV2.
func (o *CreateForecastingExperimentRequest_SdkV2) SetHolidayRegions(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["holiday_regions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.HolidayRegions = types.ListValueMust(t, vs)
}

// GetIncludeFeatures returns the value of the IncludeFeatures field in CreateForecastingExperimentRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateForecastingExperimentRequest_SdkV2) GetIncludeFeatures(ctx context.Context) ([]types.String, bool) {
	if o.IncludeFeatures.IsNull() || o.IncludeFeatures.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.IncludeFeatures.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIncludeFeatures sets the value of the IncludeFeatures field in CreateForecastingExperimentRequest_SdkV2.
func (o *CreateForecastingExperimentRequest_SdkV2) SetIncludeFeatures(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["include_features"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.IncludeFeatures = types.ListValueMust(t, vs)
}

// GetTimeseriesIdentifierColumns returns the value of the TimeseriesIdentifierColumns field in CreateForecastingExperimentRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateForecastingExperimentRequest_SdkV2) GetTimeseriesIdentifierColumns(ctx context.Context) ([]types.String, bool) {
	if o.TimeseriesIdentifierColumns.IsNull() || o.TimeseriesIdentifierColumns.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.TimeseriesIdentifierColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTimeseriesIdentifierColumns sets the value of the TimeseriesIdentifierColumns field in CreateForecastingExperimentRequest_SdkV2.
func (o *CreateForecastingExperimentRequest_SdkV2) SetTimeseriesIdentifierColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["timeseries_identifier_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.TimeseriesIdentifierColumns = types.ListValueMust(t, vs)
}

// GetTrainingFrameworks returns the value of the TrainingFrameworks field in CreateForecastingExperimentRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateForecastingExperimentRequest_SdkV2) GetTrainingFrameworks(ctx context.Context) ([]types.String, bool) {
	if o.TrainingFrameworks.IsNull() || o.TrainingFrameworks.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.TrainingFrameworks.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTrainingFrameworks sets the value of the TrainingFrameworks field in CreateForecastingExperimentRequest_SdkV2.
func (o *CreateForecastingExperimentRequest_SdkV2) SetTrainingFrameworks(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["training_frameworks"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.TrainingFrameworks = types.ListValueMust(t, vs)
}

type CreateForecastingExperimentResponse_SdkV2 struct {
	// The unique ID of the created forecasting experiment
	ExperimentId types.String `tfsdk:"experiment_id"`
}

func (toState *CreateForecastingExperimentResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateForecastingExperimentResponse_SdkV2) {
}

func (toState *CreateForecastingExperimentResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CreateForecastingExperimentResponse_SdkV2) {
}

func (c CreateForecastingExperimentResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateForecastingExperimentResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateForecastingExperimentResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateForecastingExperimentResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": o.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateForecastingExperimentResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *CreateLoggedModelRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateLoggedModelRequest_SdkV2) {
}

func (toState *CreateLoggedModelRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CreateLoggedModelRequest_SdkV2) {
}

func (c CreateLoggedModelRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateLoggedModelRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"params": reflect.TypeOf(LoggedModelParameter_SdkV2{}),
		"tags":   reflect.TypeOf(LoggedModelTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateLoggedModelRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateLoggedModelRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": o.ExperimentId,
			"model_type":    o.ModelType,
			"name":          o.Name,
			"params":        o.Params,
			"source_run_id": o.SourceRunId,
			"tags":          o.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateLoggedModelRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *CreateLoggedModelRequest_SdkV2) GetParams(ctx context.Context) ([]LoggedModelParameter_SdkV2, bool) {
	if o.Params.IsNull() || o.Params.IsUnknown() {
		return nil, false
	}
	var v []LoggedModelParameter_SdkV2
	d := o.Params.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParams sets the value of the Params field in CreateLoggedModelRequest_SdkV2.
func (o *CreateLoggedModelRequest_SdkV2) SetParams(ctx context.Context, v []LoggedModelParameter_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Params = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in CreateLoggedModelRequest_SdkV2 as
// a slice of LoggedModelTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateLoggedModelRequest_SdkV2) GetTags(ctx context.Context) ([]LoggedModelTag_SdkV2, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []LoggedModelTag_SdkV2
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in CreateLoggedModelRequest_SdkV2.
func (o *CreateLoggedModelRequest_SdkV2) SetTags(ctx context.Context, v []LoggedModelTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type CreateLoggedModelResponse_SdkV2 struct {
	// The newly created logged model.
	Model types.List `tfsdk:"model"`
}

func (toState *CreateLoggedModelResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateLoggedModelResponse_SdkV2) {
	if !fromPlan.Model.IsNull() && !fromPlan.Model.IsUnknown() {
		if toStateModel, ok := toState.GetModel(ctx); ok {
			if fromPlanModel, ok := fromPlan.GetModel(ctx); ok {
				toStateModel.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanModel)
				toState.SetModel(ctx, toStateModel)
			}
		}
	}
}

func (toState *CreateLoggedModelResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CreateLoggedModelResponse_SdkV2) {
	if !fromState.Model.IsNull() && !fromState.Model.IsUnknown() {
		if toStateModel, ok := toState.GetModel(ctx); ok {
			if fromStateModel, ok := fromState.GetModel(ctx); ok {
				toStateModel.SyncFieldsDuringRead(ctx, fromStateModel)
				toState.SetModel(ctx, toStateModel)
			}
		}
	}
}

func (c CreateLoggedModelResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateLoggedModelResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model": reflect.TypeOf(LoggedModel_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateLoggedModelResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateLoggedModelResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model": o.Model,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateLoggedModelResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *CreateLoggedModelResponse_SdkV2) GetModel(ctx context.Context) (LoggedModel_SdkV2, bool) {
	var e LoggedModel_SdkV2
	if o.Model.IsNull() || o.Model.IsUnknown() {
		return e, false
	}
	var v []LoggedModel_SdkV2
	d := o.Model.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetModel sets the value of the Model field in CreateLoggedModelResponse_SdkV2.
func (o *CreateLoggedModelResponse_SdkV2) SetModel(ctx context.Context, v LoggedModel_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["model"]
	o.Model = types.ListValueMust(t, vs)
}

type CreateModelRequest_SdkV2 struct {
	// Optional description for registered model.
	Description types.String `tfsdk:"description"`
	// Register models under this name
	Name types.String `tfsdk:"name"`
	// Additional metadata for registered model.
	Tags types.List `tfsdk:"tags"`
}

func (toState *CreateModelRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateModelRequest_SdkV2) {
}

func (toState *CreateModelRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CreateModelRequest_SdkV2) {
}

func (c CreateModelRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateModelRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(ModelTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateModelRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateModelRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description": o.Description,
			"name":        o.Name,
			"tags":        o.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateModelRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *CreateModelRequest_SdkV2) GetTags(ctx context.Context) ([]ModelTag_SdkV2, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []ModelTag_SdkV2
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in CreateModelRequest_SdkV2.
func (o *CreateModelRequest_SdkV2) SetTags(ctx context.Context, v []ModelTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type CreateModelResponse_SdkV2 struct {
	RegisteredModel types.List `tfsdk:"registered_model"`
}

func (toState *CreateModelResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateModelResponse_SdkV2) {
	if !fromPlan.RegisteredModel.IsNull() && !fromPlan.RegisteredModel.IsUnknown() {
		if toStateRegisteredModel, ok := toState.GetRegisteredModel(ctx); ok {
			if fromPlanRegisteredModel, ok := fromPlan.GetRegisteredModel(ctx); ok {
				toStateRegisteredModel.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanRegisteredModel)
				toState.SetRegisteredModel(ctx, toStateRegisteredModel)
			}
		}
	}
}

func (toState *CreateModelResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CreateModelResponse_SdkV2) {
	if !fromState.RegisteredModel.IsNull() && !fromState.RegisteredModel.IsUnknown() {
		if toStateRegisteredModel, ok := toState.GetRegisteredModel(ctx); ok {
			if fromStateRegisteredModel, ok := fromState.GetRegisteredModel(ctx); ok {
				toStateRegisteredModel.SyncFieldsDuringRead(ctx, fromStateRegisteredModel)
				toState.SetRegisteredModel(ctx, toStateRegisteredModel)
			}
		}
	}
}

func (c CreateModelResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateModelResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"registered_model": reflect.TypeOf(Model_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateModelResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateModelResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"registered_model": o.RegisteredModel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateModelResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *CreateModelResponse_SdkV2) GetRegisteredModel(ctx context.Context) (Model_SdkV2, bool) {
	var e Model_SdkV2
	if o.RegisteredModel.IsNull() || o.RegisteredModel.IsUnknown() {
		return e, false
	}
	var v []Model_SdkV2
	d := o.RegisteredModel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRegisteredModel sets the value of the RegisteredModel field in CreateModelResponse_SdkV2.
func (o *CreateModelResponse_SdkV2) SetRegisteredModel(ctx context.Context, v Model_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["registered_model"]
	o.RegisteredModel = types.ListValueMust(t, vs)
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

func (toState *CreateModelVersionRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateModelVersionRequest_SdkV2) {
}

func (toState *CreateModelVersionRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CreateModelVersionRequest_SdkV2) {
}

func (c CreateModelVersionRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateModelVersionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(ModelVersionTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateModelVersionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateModelVersionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description": o.Description,
			"name":        o.Name,
			"run_id":      o.RunId,
			"run_link":    o.RunLink,
			"source":      o.Source,
			"tags":        o.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateModelVersionRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *CreateModelVersionRequest_SdkV2) GetTags(ctx context.Context) ([]ModelVersionTag_SdkV2, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []ModelVersionTag_SdkV2
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in CreateModelVersionRequest_SdkV2.
func (o *CreateModelVersionRequest_SdkV2) SetTags(ctx context.Context, v []ModelVersionTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type CreateModelVersionResponse_SdkV2 struct {
	// Return new version number generated for this model in registry.
	ModelVersion types.List `tfsdk:"model_version"`
}

func (toState *CreateModelVersionResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateModelVersionResponse_SdkV2) {
	if !fromPlan.ModelVersion.IsNull() && !fromPlan.ModelVersion.IsUnknown() {
		if toStateModelVersion, ok := toState.GetModelVersion(ctx); ok {
			if fromPlanModelVersion, ok := fromPlan.GetModelVersion(ctx); ok {
				toStateModelVersion.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanModelVersion)
				toState.SetModelVersion(ctx, toStateModelVersion)
			}
		}
	}
}

func (toState *CreateModelVersionResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CreateModelVersionResponse_SdkV2) {
	if !fromState.ModelVersion.IsNull() && !fromState.ModelVersion.IsUnknown() {
		if toStateModelVersion, ok := toState.GetModelVersion(ctx); ok {
			if fromStateModelVersion, ok := fromState.GetModelVersion(ctx); ok {
				toStateModelVersion.SyncFieldsDuringRead(ctx, fromStateModelVersion)
				toState.SetModelVersion(ctx, toStateModelVersion)
			}
		}
	}
}

func (c CreateModelVersionResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateModelVersionResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model_version": reflect.TypeOf(ModelVersion_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateModelVersionResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateModelVersionResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_version": o.ModelVersion,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateModelVersionResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *CreateModelVersionResponse_SdkV2) GetModelVersion(ctx context.Context) (ModelVersion_SdkV2, bool) {
	var e ModelVersion_SdkV2
	if o.ModelVersion.IsNull() || o.ModelVersion.IsUnknown() {
		return e, false
	}
	var v []ModelVersion_SdkV2
	d := o.ModelVersion.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetModelVersion sets the value of the ModelVersion field in CreateModelVersionResponse_SdkV2.
func (o *CreateModelVersionResponse_SdkV2) SetModelVersion(ctx context.Context, v ModelVersion_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["model_version"]
	o.ModelVersion = types.ListValueMust(t, vs)
}

type CreateOnlineStoreRequest_SdkV2 struct {
	// Online store to create.
	OnlineStore types.List `tfsdk:"online_store"`
}

func (toState *CreateOnlineStoreRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateOnlineStoreRequest_SdkV2) {
	if !fromPlan.OnlineStore.IsNull() && !fromPlan.OnlineStore.IsUnknown() {
		if toStateOnlineStore, ok := toState.GetOnlineStore(ctx); ok {
			if fromPlanOnlineStore, ok := fromPlan.GetOnlineStore(ctx); ok {
				toStateOnlineStore.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanOnlineStore)
				toState.SetOnlineStore(ctx, toStateOnlineStore)
			}
		}
	}
}

func (toState *CreateOnlineStoreRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CreateOnlineStoreRequest_SdkV2) {
	if !fromState.OnlineStore.IsNull() && !fromState.OnlineStore.IsUnknown() {
		if toStateOnlineStore, ok := toState.GetOnlineStore(ctx); ok {
			if fromStateOnlineStore, ok := fromState.GetOnlineStore(ctx); ok {
				toStateOnlineStore.SyncFieldsDuringRead(ctx, fromStateOnlineStore)
				toState.SetOnlineStore(ctx, toStateOnlineStore)
			}
		}
	}
}

func (c CreateOnlineStoreRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateOnlineStoreRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"online_store": reflect.TypeOf(OnlineStore_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateOnlineStoreRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateOnlineStoreRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"online_store": o.OnlineStore,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateOnlineStoreRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *CreateOnlineStoreRequest_SdkV2) GetOnlineStore(ctx context.Context) (OnlineStore_SdkV2, bool) {
	var e OnlineStore_SdkV2
	if o.OnlineStore.IsNull() || o.OnlineStore.IsUnknown() {
		return e, false
	}
	var v []OnlineStore_SdkV2
	d := o.OnlineStore.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOnlineStore sets the value of the OnlineStore field in CreateOnlineStoreRequest_SdkV2.
func (o *CreateOnlineStoreRequest_SdkV2) SetOnlineStore(ctx context.Context, v OnlineStore_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["online_store"]
	o.OnlineStore = types.ListValueMust(t, vs)
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

func (toState *CreateRegistryWebhook_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateRegistryWebhook_SdkV2) {
	if !fromPlan.HttpUrlSpec.IsNull() && !fromPlan.HttpUrlSpec.IsUnknown() {
		if toStateHttpUrlSpec, ok := toState.GetHttpUrlSpec(ctx); ok {
			if fromPlanHttpUrlSpec, ok := fromPlan.GetHttpUrlSpec(ctx); ok {
				toStateHttpUrlSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanHttpUrlSpec)
				toState.SetHttpUrlSpec(ctx, toStateHttpUrlSpec)
			}
		}
	}
	if !fromPlan.JobSpec.IsNull() && !fromPlan.JobSpec.IsUnknown() {
		if toStateJobSpec, ok := toState.GetJobSpec(ctx); ok {
			if fromPlanJobSpec, ok := fromPlan.GetJobSpec(ctx); ok {
				toStateJobSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanJobSpec)
				toState.SetJobSpec(ctx, toStateJobSpec)
			}
		}
	}
}

func (toState *CreateRegistryWebhook_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CreateRegistryWebhook_SdkV2) {
	if !fromState.HttpUrlSpec.IsNull() && !fromState.HttpUrlSpec.IsUnknown() {
		if toStateHttpUrlSpec, ok := toState.GetHttpUrlSpec(ctx); ok {
			if fromStateHttpUrlSpec, ok := fromState.GetHttpUrlSpec(ctx); ok {
				toStateHttpUrlSpec.SyncFieldsDuringRead(ctx, fromStateHttpUrlSpec)
				toState.SetHttpUrlSpec(ctx, toStateHttpUrlSpec)
			}
		}
	}
	if !fromState.JobSpec.IsNull() && !fromState.JobSpec.IsUnknown() {
		if toStateJobSpec, ok := toState.GetJobSpec(ctx); ok {
			if fromStateJobSpec, ok := fromState.GetJobSpec(ctx); ok {
				toStateJobSpec.SyncFieldsDuringRead(ctx, fromStateJobSpec)
				toState.SetJobSpec(ctx, toStateJobSpec)
			}
		}
	}
}

func (c CreateRegistryWebhook_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateRegistryWebhook_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"events":        reflect.TypeOf(types.String{}),
		"http_url_spec": reflect.TypeOf(HttpUrlSpec_SdkV2{}),
		"job_spec":      reflect.TypeOf(JobSpec_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateRegistryWebhook_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateRegistryWebhook_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":   o.Description,
			"events":        o.Events,
			"http_url_spec": o.HttpUrlSpec,
			"job_spec":      o.JobSpec,
			"model_name":    o.ModelName,
			"status":        o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateRegistryWebhook_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *CreateRegistryWebhook_SdkV2) GetEvents(ctx context.Context) ([]types.String, bool) {
	if o.Events.IsNull() || o.Events.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Events.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEvents sets the value of the Events field in CreateRegistryWebhook_SdkV2.
func (o *CreateRegistryWebhook_SdkV2) SetEvents(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["events"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Events = types.ListValueMust(t, vs)
}

// GetHttpUrlSpec returns the value of the HttpUrlSpec field in CreateRegistryWebhook_SdkV2 as
// a HttpUrlSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateRegistryWebhook_SdkV2) GetHttpUrlSpec(ctx context.Context) (HttpUrlSpec_SdkV2, bool) {
	var e HttpUrlSpec_SdkV2
	if o.HttpUrlSpec.IsNull() || o.HttpUrlSpec.IsUnknown() {
		return e, false
	}
	var v []HttpUrlSpec_SdkV2
	d := o.HttpUrlSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetHttpUrlSpec sets the value of the HttpUrlSpec field in CreateRegistryWebhook_SdkV2.
func (o *CreateRegistryWebhook_SdkV2) SetHttpUrlSpec(ctx context.Context, v HttpUrlSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["http_url_spec"]
	o.HttpUrlSpec = types.ListValueMust(t, vs)
}

// GetJobSpec returns the value of the JobSpec field in CreateRegistryWebhook_SdkV2 as
// a JobSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateRegistryWebhook_SdkV2) GetJobSpec(ctx context.Context) (JobSpec_SdkV2, bool) {
	var e JobSpec_SdkV2
	if o.JobSpec.IsNull() || o.JobSpec.IsUnknown() {
		return e, false
	}
	var v []JobSpec_SdkV2
	d := o.JobSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetJobSpec sets the value of the JobSpec field in CreateRegistryWebhook_SdkV2.
func (o *CreateRegistryWebhook_SdkV2) SetJobSpec(ctx context.Context, v JobSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["job_spec"]
	o.JobSpec = types.ListValueMust(t, vs)
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

func (toState *CreateRun_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateRun_SdkV2) {
}

func (toState *CreateRun_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CreateRun_SdkV2) {
}

func (c CreateRun_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateRun_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(RunTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateRun_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateRun_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": o.ExperimentId,
			"run_name":      o.RunName,
			"start_time":    o.StartTime,
			"tags":          o.Tags,
			"user_id":       o.UserId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateRun_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *CreateRun_SdkV2) GetTags(ctx context.Context) ([]RunTag_SdkV2, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []RunTag_SdkV2
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in CreateRun_SdkV2.
func (o *CreateRun_SdkV2) SetTags(ctx context.Context, v []RunTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type CreateRunResponse_SdkV2 struct {
	// The newly created run.
	Run types.List `tfsdk:"run"`
}

func (toState *CreateRunResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateRunResponse_SdkV2) {
	if !fromPlan.Run.IsNull() && !fromPlan.Run.IsUnknown() {
		if toStateRun, ok := toState.GetRun(ctx); ok {
			if fromPlanRun, ok := fromPlan.GetRun(ctx); ok {
				toStateRun.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanRun)
				toState.SetRun(ctx, toStateRun)
			}
		}
	}
}

func (toState *CreateRunResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CreateRunResponse_SdkV2) {
	if !fromState.Run.IsNull() && !fromState.Run.IsUnknown() {
		if toStateRun, ok := toState.GetRun(ctx); ok {
			if fromStateRun, ok := fromState.GetRun(ctx); ok {
				toStateRun.SyncFieldsDuringRead(ctx, fromStateRun)
				toState.SetRun(ctx, toStateRun)
			}
		}
	}
}

func (c CreateRunResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateRunResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"run": reflect.TypeOf(Run_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateRunResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateRunResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run": o.Run,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateRunResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *CreateRunResponse_SdkV2) GetRun(ctx context.Context) (Run_SdkV2, bool) {
	var e Run_SdkV2
	if o.Run.IsNull() || o.Run.IsUnknown() {
		return e, false
	}
	var v []Run_SdkV2
	d := o.Run.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRun sets the value of the Run field in CreateRunResponse_SdkV2.
func (o *CreateRunResponse_SdkV2) SetRun(ctx context.Context, v Run_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["run"]
	o.Run = types.ListValueMust(t, vs)
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

func (toState *CreateTransitionRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateTransitionRequest_SdkV2) {
}

func (toState *CreateTransitionRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CreateTransitionRequest_SdkV2) {
}

func (c CreateTransitionRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateTransitionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateTransitionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateTransitionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment": o.Comment,
			"name":    o.Name,
			"stage":   o.Stage,
			"version": o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateTransitionRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *CreateTransitionRequestResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateTransitionRequestResponse_SdkV2) {
	if !fromPlan.Request.IsNull() && !fromPlan.Request.IsUnknown() {
		if toStateRequest, ok := toState.GetRequest(ctx); ok {
			if fromPlanRequest, ok := fromPlan.GetRequest(ctx); ok {
				toStateRequest.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanRequest)
				toState.SetRequest(ctx, toStateRequest)
			}
		}
	}
}

func (toState *CreateTransitionRequestResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CreateTransitionRequestResponse_SdkV2) {
	if !fromState.Request.IsNull() && !fromState.Request.IsUnknown() {
		if toStateRequest, ok := toState.GetRequest(ctx); ok {
			if fromStateRequest, ok := fromState.GetRequest(ctx); ok {
				toStateRequest.SyncFieldsDuringRead(ctx, fromStateRequest)
				toState.SetRequest(ctx, toStateRequest)
			}
		}
	}
}

func (c CreateTransitionRequestResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateTransitionRequestResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"request": reflect.TypeOf(TransitionRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateTransitionRequestResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateTransitionRequestResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"request": o.Request,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateTransitionRequestResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *CreateTransitionRequestResponse_SdkV2) GetRequest(ctx context.Context) (TransitionRequest_SdkV2, bool) {
	var e TransitionRequest_SdkV2
	if o.Request.IsNull() || o.Request.IsUnknown() {
		return e, false
	}
	var v []TransitionRequest_SdkV2
	d := o.Request.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRequest sets the value of the Request field in CreateTransitionRequestResponse_SdkV2.
func (o *CreateTransitionRequestResponse_SdkV2) SetRequest(ctx context.Context, v TransitionRequest_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["request"]
	o.Request = types.ListValueMust(t, vs)
}

type CreateWebhookResponse_SdkV2 struct {
	Webhook types.List `tfsdk:"webhook"`
}

func (toState *CreateWebhookResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateWebhookResponse_SdkV2) {
	if !fromPlan.Webhook.IsNull() && !fromPlan.Webhook.IsUnknown() {
		if toStateWebhook, ok := toState.GetWebhook(ctx); ok {
			if fromPlanWebhook, ok := fromPlan.GetWebhook(ctx); ok {
				toStateWebhook.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanWebhook)
				toState.SetWebhook(ctx, toStateWebhook)
			}
		}
	}
}

func (toState *CreateWebhookResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CreateWebhookResponse_SdkV2) {
	if !fromState.Webhook.IsNull() && !fromState.Webhook.IsUnknown() {
		if toStateWebhook, ok := toState.GetWebhook(ctx); ok {
			if fromStateWebhook, ok := fromState.GetWebhook(ctx); ok {
				toStateWebhook.SyncFieldsDuringRead(ctx, fromStateWebhook)
				toState.SetWebhook(ctx, toStateWebhook)
			}
		}
	}
}

func (c CreateWebhookResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateWebhookResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"webhook": reflect.TypeOf(RegistryWebhook_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateWebhookResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateWebhookResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"webhook": o.Webhook,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateWebhookResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *CreateWebhookResponse_SdkV2) GetWebhook(ctx context.Context) (RegistryWebhook_SdkV2, bool) {
	var e RegistryWebhook_SdkV2
	if o.Webhook.IsNull() || o.Webhook.IsUnknown() {
		return e, false
	}
	var v []RegistryWebhook_SdkV2
	d := o.Webhook.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWebhook sets the value of the Webhook field in CreateWebhookResponse_SdkV2.
func (o *CreateWebhookResponse_SdkV2) SetWebhook(ctx context.Context, v RegistryWebhook_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["webhook"]
	o.Webhook = types.ListValueMust(t, vs)
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

func (toState *Dataset_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Dataset_SdkV2) {
}

func (toState *Dataset_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState Dataset_SdkV2) {
}

func (c Dataset_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Dataset_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Dataset_SdkV2
// only implements ToObjectValue() and Type().
func (o Dataset_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"digest":      o.Digest,
			"name":        o.Name,
			"profile":     o.Profile,
			"schema":      o.Schema,
			"source":      o.Source,
			"source_type": o.SourceType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Dataset_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *DatasetInput_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DatasetInput_SdkV2) {
	if !fromPlan.Dataset.IsNull() && !fromPlan.Dataset.IsUnknown() {
		if toStateDataset, ok := toState.GetDataset(ctx); ok {
			if fromPlanDataset, ok := fromPlan.GetDataset(ctx); ok {
				toStateDataset.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanDataset)
				toState.SetDataset(ctx, toStateDataset)
			}
		}
	}
}

func (toState *DatasetInput_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DatasetInput_SdkV2) {
	if !fromState.Dataset.IsNull() && !fromState.Dataset.IsUnknown() {
		if toStateDataset, ok := toState.GetDataset(ctx); ok {
			if fromStateDataset, ok := fromState.GetDataset(ctx); ok {
				toStateDataset.SyncFieldsDuringRead(ctx, fromStateDataset)
				toState.SetDataset(ctx, toStateDataset)
			}
		}
	}
}

func (c DatasetInput_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DatasetInput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dataset": reflect.TypeOf(Dataset_SdkV2{}),
		"tags":    reflect.TypeOf(InputTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatasetInput_SdkV2
// only implements ToObjectValue() and Type().
func (o DatasetInput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dataset": o.Dataset,
			"tags":    o.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DatasetInput_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *DatasetInput_SdkV2) GetDataset(ctx context.Context) (Dataset_SdkV2, bool) {
	var e Dataset_SdkV2
	if o.Dataset.IsNull() || o.Dataset.IsUnknown() {
		return e, false
	}
	var v []Dataset_SdkV2
	d := o.Dataset.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDataset sets the value of the Dataset field in DatasetInput_SdkV2.
func (o *DatasetInput_SdkV2) SetDataset(ctx context.Context, v Dataset_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dataset"]
	o.Dataset = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in DatasetInput_SdkV2 as
// a slice of InputTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *DatasetInput_SdkV2) GetTags(ctx context.Context) ([]InputTag_SdkV2, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []InputTag_SdkV2
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in DatasetInput_SdkV2.
func (o *DatasetInput_SdkV2) SetTags(ctx context.Context, v []InputTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type DeleteCommentRequest_SdkV2 struct {
	// Unique identifier of an activity
	Id types.String `tfsdk:"-"`
}

func (toState *DeleteCommentRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteCommentRequest_SdkV2) {
}

func (toState *DeleteCommentRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteCommentRequest_SdkV2) {
}

func (c DeleteCommentRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteCommentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCommentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteCommentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCommentRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteCommentResponse_SdkV2 struct {
}

func (toState *DeleteCommentResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteCommentResponse_SdkV2) {
}

func (toState *DeleteCommentResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteCommentResponse_SdkV2) {
}

func (c DeleteCommentResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCommentResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteCommentResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCommentResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteCommentResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCommentResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteExperiment_SdkV2 struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"experiment_id"`
}

func (toState *DeleteExperiment_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteExperiment_SdkV2) {
}

func (toState *DeleteExperiment_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteExperiment_SdkV2) {
}

func (c DeleteExperiment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteExperiment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteExperiment_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteExperiment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": o.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteExperiment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
		},
	}
}

type DeleteExperimentResponse_SdkV2 struct {
}

func (toState *DeleteExperimentResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteExperimentResponse_SdkV2) {
}

func (toState *DeleteExperimentResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteExperimentResponse_SdkV2) {
}

func (c DeleteExperimentResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteExperimentResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteExperimentResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteExperimentResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteExperimentResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteExperimentResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
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

func (toState *DeleteFeatureTagRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteFeatureTagRequest_SdkV2) {
}

func (toState *DeleteFeatureTagRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteFeatureTagRequest_SdkV2) {
}

func (c DeleteFeatureTagRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteFeatureTagRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteFeatureTagRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteFeatureTagRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_name": o.FeatureName,
			"key":          o.Key,
			"table_name":   o.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteFeatureTagRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *DeleteLoggedModelRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteLoggedModelRequest_SdkV2) {
}

func (toState *DeleteLoggedModelRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteLoggedModelRequest_SdkV2) {
}

func (c DeleteLoggedModelRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteLoggedModelRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteLoggedModelRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteLoggedModelRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_id": o.ModelId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteLoggedModelRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_id": types.StringType,
		},
	}
}

type DeleteLoggedModelResponse_SdkV2 struct {
}

func (toState *DeleteLoggedModelResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteLoggedModelResponse_SdkV2) {
}

func (toState *DeleteLoggedModelResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteLoggedModelResponse_SdkV2) {
}

func (c DeleteLoggedModelResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteLoggedModelResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteLoggedModelResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteLoggedModelResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteLoggedModelResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteLoggedModelResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *DeleteLoggedModelTagRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteLoggedModelTagRequest_SdkV2) {
}

func (toState *DeleteLoggedModelTagRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteLoggedModelTagRequest_SdkV2) {
}

func (c DeleteLoggedModelTagRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteLoggedModelTagRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteLoggedModelTagRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteLoggedModelTagRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_id": o.ModelId,
			"tag_key":  o.TagKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteLoggedModelTagRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_id": types.StringType,
			"tag_key":  types.StringType,
		},
	}
}

type DeleteLoggedModelTagResponse_SdkV2 struct {
}

func (toState *DeleteLoggedModelTagResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteLoggedModelTagResponse_SdkV2) {
}

func (toState *DeleteLoggedModelTagResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteLoggedModelTagResponse_SdkV2) {
}

func (c DeleteLoggedModelTagResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteLoggedModelTagResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteLoggedModelTagResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteLoggedModelTagResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteLoggedModelTagResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteLoggedModelTagResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteModelRequest_SdkV2 struct {
	// Registered model unique name identifier.
	Name types.String `tfsdk:"-"`
}

func (toState *DeleteModelRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteModelRequest_SdkV2) {
}

func (toState *DeleteModelRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteModelRequest_SdkV2) {
}

func (c DeleteModelRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteModelRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteModelRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteModelRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteModelResponse_SdkV2 struct {
}

func (toState *DeleteModelResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteModelResponse_SdkV2) {
}

func (toState *DeleteModelResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteModelResponse_SdkV2) {
}

func (c DeleteModelResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteModelResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteModelResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteModelResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteModelResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *DeleteModelTagRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteModelTagRequest_SdkV2) {
}

func (toState *DeleteModelTagRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteModelTagRequest_SdkV2) {
}

func (c DeleteModelTagRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteModelTagRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelTagRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteModelTagRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":  o.Key,
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteModelTagRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":  types.StringType,
			"name": types.StringType,
		},
	}
}

type DeleteModelTagResponse_SdkV2 struct {
}

func (toState *DeleteModelTagResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteModelTagResponse_SdkV2) {
}

func (toState *DeleteModelTagResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteModelTagResponse_SdkV2) {
}

func (c DeleteModelTagResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteModelTagResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteModelTagResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelTagResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteModelTagResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteModelTagResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *DeleteModelVersionRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteModelVersionRequest_SdkV2) {
}

func (toState *DeleteModelVersionRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteModelVersionRequest_SdkV2) {
}

func (c DeleteModelVersionRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteModelVersionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelVersionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteModelVersionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":    o.Name,
			"version": o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteModelVersionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":    types.StringType,
			"version": types.StringType,
		},
	}
}

type DeleteModelVersionResponse_SdkV2 struct {
}

func (toState *DeleteModelVersionResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteModelVersionResponse_SdkV2) {
}

func (toState *DeleteModelVersionResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteModelVersionResponse_SdkV2) {
}

func (c DeleteModelVersionResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteModelVersionResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteModelVersionResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelVersionResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteModelVersionResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteModelVersionResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *DeleteModelVersionTagRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteModelVersionTagRequest_SdkV2) {
}

func (toState *DeleteModelVersionTagRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteModelVersionTagRequest_SdkV2) {
}

func (c DeleteModelVersionTagRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteModelVersionTagRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelVersionTagRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteModelVersionTagRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":     o.Key,
			"name":    o.Name,
			"version": o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteModelVersionTagRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *DeleteModelVersionTagResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteModelVersionTagResponse_SdkV2) {
}

func (toState *DeleteModelVersionTagResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteModelVersionTagResponse_SdkV2) {
}

func (c DeleteModelVersionTagResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteModelVersionTagResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteModelVersionTagResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelVersionTagResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteModelVersionTagResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteModelVersionTagResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteOnlineStoreRequest_SdkV2 struct {
	// Name of the online store to delete.
	Name types.String `tfsdk:"-"`
}

func (toState *DeleteOnlineStoreRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteOnlineStoreRequest_SdkV2) {
}

func (toState *DeleteOnlineStoreRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteOnlineStoreRequest_SdkV2) {
}

func (c DeleteOnlineStoreRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteOnlineStoreRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteOnlineStoreRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteOnlineStoreRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteOnlineStoreRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *DeleteRun_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteRun_SdkV2) {
}

func (toState *DeleteRun_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteRun_SdkV2) {
}

func (c DeleteRun_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteRun_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRun_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteRun_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run_id": o.RunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteRun_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run_id": types.StringType,
		},
	}
}

type DeleteRunResponse_SdkV2 struct {
}

func (toState *DeleteRunResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteRunResponse_SdkV2) {
}

func (toState *DeleteRunResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteRunResponse_SdkV2) {
}

func (c DeleteRunResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteRunResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteRunResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRunResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteRunResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteRunResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *DeleteRuns_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteRuns_SdkV2) {
}

func (toState *DeleteRuns_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteRuns_SdkV2) {
}

func (c DeleteRuns_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteRuns_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRuns_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteRuns_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id":        o.ExperimentId,
			"max_runs":             o.MaxRuns,
			"max_timestamp_millis": o.MaxTimestampMillis,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteRuns_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *DeleteRunsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteRunsResponse_SdkV2) {
}

func (toState *DeleteRunsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteRunsResponse_SdkV2) {
}

func (c DeleteRunsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteRunsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRunsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteRunsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"runs_deleted": o.RunsDeleted,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteRunsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *DeleteTag_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteTag_SdkV2) {
}

func (toState *DeleteTag_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteTag_SdkV2) {
}

func (c DeleteTag_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteTag_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteTag_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteTag_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":    o.Key,
			"run_id": o.RunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteTag_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":    types.StringType,
			"run_id": types.StringType,
		},
	}
}

type DeleteTagResponse_SdkV2 struct {
}

func (toState *DeleteTagResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteTagResponse_SdkV2) {
}

func (toState *DeleteTagResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteTagResponse_SdkV2) {
}

func (c DeleteTagResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteTagResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteTagResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteTagResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteTagResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteTagResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *DeleteTransitionRequestRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteTransitionRequestRequest_SdkV2) {
}

func (toState *DeleteTransitionRequestRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteTransitionRequestRequest_SdkV2) {
}

func (c DeleteTransitionRequestRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteTransitionRequestRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteTransitionRequestRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteTransitionRequestRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment": o.Comment,
			"creator": o.Creator,
			"name":    o.Name,
			"stage":   o.Stage,
			"version": o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteTransitionRequestRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *DeleteTransitionRequestResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteTransitionRequestResponse_SdkV2) {
	if !fromPlan.Activity.IsNull() && !fromPlan.Activity.IsUnknown() {
		if toStateActivity, ok := toState.GetActivity(ctx); ok {
			if fromPlanActivity, ok := fromPlan.GetActivity(ctx); ok {
				toStateActivity.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanActivity)
				toState.SetActivity(ctx, toStateActivity)
			}
		}
	}
}

func (toState *DeleteTransitionRequestResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteTransitionRequestResponse_SdkV2) {
	if !fromState.Activity.IsNull() && !fromState.Activity.IsUnknown() {
		if toStateActivity, ok := toState.GetActivity(ctx); ok {
			if fromStateActivity, ok := fromState.GetActivity(ctx); ok {
				toStateActivity.SyncFieldsDuringRead(ctx, fromStateActivity)
				toState.SetActivity(ctx, toStateActivity)
			}
		}
	}
}

func (c DeleteTransitionRequestResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteTransitionRequestResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"activity": reflect.TypeOf(Activity_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteTransitionRequestResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteTransitionRequestResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"activity": o.Activity,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteTransitionRequestResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *DeleteTransitionRequestResponse_SdkV2) GetActivity(ctx context.Context) (Activity_SdkV2, bool) {
	var e Activity_SdkV2
	if o.Activity.IsNull() || o.Activity.IsUnknown() {
		return e, false
	}
	var v []Activity_SdkV2
	d := o.Activity.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetActivity sets the value of the Activity field in DeleteTransitionRequestResponse_SdkV2.
func (o *DeleteTransitionRequestResponse_SdkV2) SetActivity(ctx context.Context, v Activity_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["activity"]
	o.Activity = types.ListValueMust(t, vs)
}

type DeleteWebhookRequest_SdkV2 struct {
	// Webhook ID required to delete a registry webhook.
	Id types.String `tfsdk:"-"`
}

func (toState *DeleteWebhookRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteWebhookRequest_SdkV2) {
}

func (toState *DeleteWebhookRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteWebhookRequest_SdkV2) {
}

func (c DeleteWebhookRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteWebhookRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWebhookRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteWebhookRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteWebhookRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteWebhookResponse_SdkV2 struct {
}

func (toState *DeleteWebhookResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteWebhookResponse_SdkV2) {
}

func (toState *DeleteWebhookResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteWebhookResponse_SdkV2) {
}

func (c DeleteWebhookResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteWebhookResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteWebhookResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWebhookResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteWebhookResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteWebhookResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
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

func (toState *Experiment_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Experiment_SdkV2) {
}

func (toState *Experiment_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState Experiment_SdkV2) {
}

func (c Experiment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Experiment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(ExperimentTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Experiment_SdkV2
// only implements ToObjectValue() and Type().
func (o Experiment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"artifact_location": o.ArtifactLocation,
			"creation_time":     o.CreationTime,
			"experiment_id":     o.ExperimentId,
			"last_update_time":  o.LastUpdateTime,
			"lifecycle_stage":   o.LifecycleStage,
			"name":              o.Name,
			"tags":              o.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Experiment_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *Experiment_SdkV2) GetTags(ctx context.Context) ([]ExperimentTag_SdkV2, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []ExperimentTag_SdkV2
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in Experiment_SdkV2.
func (o *Experiment_SdkV2) SetTags(ctx context.Context, v []ExperimentTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
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

func (toState *ExperimentAccessControlRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ExperimentAccessControlRequest_SdkV2) {
}

func (toState *ExperimentAccessControlRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ExperimentAccessControlRequest_SdkV2) {
}

func (c ExperimentAccessControlRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ExperimentAccessControlRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExperimentAccessControlRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ExperimentAccessControlRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_name":             o.GroupName,
			"permission_level":       o.PermissionLevel,
			"service_principal_name": o.ServicePrincipalName,
			"user_name":              o.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExperimentAccessControlRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *ExperimentAccessControlResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ExperimentAccessControlResponse_SdkV2) {
}

func (toState *ExperimentAccessControlResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ExperimentAccessControlResponse_SdkV2) {
}

func (c ExperimentAccessControlResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ExperimentAccessControlResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(ExperimentPermission_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExperimentAccessControlResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ExperimentAccessControlResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"all_permissions":        o.AllPermissions,
			"display_name":           o.DisplayName,
			"group_name":             o.GroupName,
			"service_principal_name": o.ServicePrincipalName,
			"user_name":              o.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExperimentAccessControlResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ExperimentAccessControlResponse_SdkV2) GetAllPermissions(ctx context.Context) ([]ExperimentPermission_SdkV2, bool) {
	if o.AllPermissions.IsNull() || o.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []ExperimentPermission_SdkV2
	d := o.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in ExperimentAccessControlResponse_SdkV2.
func (o *ExperimentAccessControlResponse_SdkV2) SetAllPermissions(ctx context.Context, v []ExperimentPermission_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllPermissions = types.ListValueMust(t, vs)
}

type ExperimentPermission_SdkV2 struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (toState *ExperimentPermission_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ExperimentPermission_SdkV2) {
}

func (toState *ExperimentPermission_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ExperimentPermission_SdkV2) {
}

func (c ExperimentPermission_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ExperimentPermission_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExperimentPermission_SdkV2
// only implements ToObjectValue() and Type().
func (o ExperimentPermission_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             o.Inherited,
			"inherited_from_object": o.InheritedFromObject,
			"permission_level":      o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExperimentPermission_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ExperimentPermission_SdkV2) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
	if o.InheritedFromObject.IsNull() || o.InheritedFromObject.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.InheritedFromObject.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInheritedFromObject sets the value of the InheritedFromObject field in ExperimentPermission_SdkV2.
func (o *ExperimentPermission_SdkV2) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InheritedFromObject = types.ListValueMust(t, vs)
}

type ExperimentPermissions_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (toState *ExperimentPermissions_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ExperimentPermissions_SdkV2) {
}

func (toState *ExperimentPermissions_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ExperimentPermissions_SdkV2) {
}

func (c ExperimentPermissions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ExperimentPermissions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(ExperimentAccessControlResponse_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExperimentPermissions_SdkV2
// only implements ToObjectValue() and Type().
func (o ExperimentPermissions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExperimentPermissions_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ExperimentPermissions_SdkV2) GetAccessControlList(ctx context.Context) ([]ExperimentAccessControlResponse_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []ExperimentAccessControlResponse_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in ExperimentPermissions_SdkV2.
func (o *ExperimentPermissions_SdkV2) SetAccessControlList(ctx context.Context, v []ExperimentAccessControlResponse_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type ExperimentPermissionsDescription_SdkV2 struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (toState *ExperimentPermissionsDescription_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ExperimentPermissionsDescription_SdkV2) {
}

func (toState *ExperimentPermissionsDescription_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ExperimentPermissionsDescription_SdkV2) {
}

func (c ExperimentPermissionsDescription_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ExperimentPermissionsDescription_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExperimentPermissionsDescription_SdkV2
// only implements ToObjectValue() and Type().
func (o ExperimentPermissionsDescription_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      o.Description,
			"permission_level": o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExperimentPermissionsDescription_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *ExperimentPermissionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ExperimentPermissionsRequest_SdkV2) {
}

func (toState *ExperimentPermissionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ExperimentPermissionsRequest_SdkV2) {
}

func (c ExperimentPermissionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ExperimentPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(ExperimentAccessControlRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExperimentPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ExperimentPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"experiment_id":       o.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExperimentPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ExperimentPermissionsRequest_SdkV2) GetAccessControlList(ctx context.Context) ([]ExperimentAccessControlRequest_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []ExperimentAccessControlRequest_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in ExperimentPermissionsRequest_SdkV2.
func (o *ExperimentPermissionsRequest_SdkV2) SetAccessControlList(ctx context.Context, v []ExperimentAccessControlRequest_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

// A tag for an experiment.
type ExperimentTag_SdkV2 struct {
	// The tag key.
	Key types.String `tfsdk:"key"`
	// The tag value.
	Value types.String `tfsdk:"value"`
}

func (toState *ExperimentTag_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ExperimentTag_SdkV2) {
}

func (toState *ExperimentTag_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ExperimentTag_SdkV2) {
}

func (c ExperimentTag_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ExperimentTag_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExperimentTag_SdkV2
// only implements ToObjectValue() and Type().
func (o ExperimentTag_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExperimentTag_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

type FeatureLineage_SdkV2 struct {
	// List of feature specs that contain this feature.
	FeatureSpecs types.List `tfsdk:"feature_specs"`
	// List of Unity Catalog models that were trained on this feature.
	Models types.List `tfsdk:"models"`
	// List of online features that use this feature as source.
	OnlineFeatures types.List `tfsdk:"online_features"`
}

func (toState *FeatureLineage_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan FeatureLineage_SdkV2) {
}

func (toState *FeatureLineage_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState FeatureLineage_SdkV2) {
}

func (c FeatureLineage_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a FeatureLineage_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"feature_specs":   reflect.TypeOf(FeatureLineageFeatureSpec_SdkV2{}),
		"models":          reflect.TypeOf(FeatureLineageModel_SdkV2{}),
		"online_features": reflect.TypeOf(FeatureLineageOnlineFeature_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FeatureLineage_SdkV2
// only implements ToObjectValue() and Type().
func (o FeatureLineage_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_specs":   o.FeatureSpecs,
			"models":          o.Models,
			"online_features": o.OnlineFeatures,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FeatureLineage_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *FeatureLineage_SdkV2) GetFeatureSpecs(ctx context.Context) ([]FeatureLineageFeatureSpec_SdkV2, bool) {
	if o.FeatureSpecs.IsNull() || o.FeatureSpecs.IsUnknown() {
		return nil, false
	}
	var v []FeatureLineageFeatureSpec_SdkV2
	d := o.FeatureSpecs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFeatureSpecs sets the value of the FeatureSpecs field in FeatureLineage_SdkV2.
func (o *FeatureLineage_SdkV2) SetFeatureSpecs(ctx context.Context, v []FeatureLineageFeatureSpec_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["feature_specs"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.FeatureSpecs = types.ListValueMust(t, vs)
}

// GetModels returns the value of the Models field in FeatureLineage_SdkV2 as
// a slice of FeatureLineageModel_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *FeatureLineage_SdkV2) GetModels(ctx context.Context) ([]FeatureLineageModel_SdkV2, bool) {
	if o.Models.IsNull() || o.Models.IsUnknown() {
		return nil, false
	}
	var v []FeatureLineageModel_SdkV2
	d := o.Models.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModels sets the value of the Models field in FeatureLineage_SdkV2.
func (o *FeatureLineage_SdkV2) SetModels(ctx context.Context, v []FeatureLineageModel_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["models"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Models = types.ListValueMust(t, vs)
}

// GetOnlineFeatures returns the value of the OnlineFeatures field in FeatureLineage_SdkV2 as
// a slice of FeatureLineageOnlineFeature_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *FeatureLineage_SdkV2) GetOnlineFeatures(ctx context.Context) ([]FeatureLineageOnlineFeature_SdkV2, bool) {
	if o.OnlineFeatures.IsNull() || o.OnlineFeatures.IsUnknown() {
		return nil, false
	}
	var v []FeatureLineageOnlineFeature_SdkV2
	d := o.OnlineFeatures.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOnlineFeatures sets the value of the OnlineFeatures field in FeatureLineage_SdkV2.
func (o *FeatureLineage_SdkV2) SetOnlineFeatures(ctx context.Context, v []FeatureLineageOnlineFeature_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["online_features"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OnlineFeatures = types.ListValueMust(t, vs)
}

type FeatureLineageFeatureSpec_SdkV2 struct {
	// The full name of the feature spec in Unity Catalog.
	Name types.String `tfsdk:"name"`
}

func (toState *FeatureLineageFeatureSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan FeatureLineageFeatureSpec_SdkV2) {
}

func (toState *FeatureLineageFeatureSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState FeatureLineageFeatureSpec_SdkV2) {
}

func (c FeatureLineageFeatureSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a FeatureLineageFeatureSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FeatureLineageFeatureSpec_SdkV2
// only implements ToObjectValue() and Type().
func (o FeatureLineageFeatureSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FeatureLineageFeatureSpec_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *FeatureLineageModel_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan FeatureLineageModel_SdkV2) {
}

func (toState *FeatureLineageModel_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState FeatureLineageModel_SdkV2) {
}

func (c FeatureLineageModel_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a FeatureLineageModel_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FeatureLineageModel_SdkV2
// only implements ToObjectValue() and Type().
func (o FeatureLineageModel_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":    o.Name,
			"version": o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FeatureLineageModel_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *FeatureLineageOnlineFeature_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan FeatureLineageOnlineFeature_SdkV2) {
}

func (toState *FeatureLineageOnlineFeature_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState FeatureLineageOnlineFeature_SdkV2) {
}

func (c FeatureLineageOnlineFeature_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a FeatureLineageOnlineFeature_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FeatureLineageOnlineFeature_SdkV2
// only implements ToObjectValue() and Type().
func (o FeatureLineageOnlineFeature_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_name": o.FeatureName,
			"table_name":   o.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FeatureLineageOnlineFeature_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *FeatureList_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan FeatureList_SdkV2) {
}

func (toState *FeatureList_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState FeatureList_SdkV2) {
}

func (c FeatureList_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a FeatureList_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"features": reflect.TypeOf(LinkedFeature_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FeatureList_SdkV2
// only implements ToObjectValue() and Type().
func (o FeatureList_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"features": o.Features,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FeatureList_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *FeatureList_SdkV2) GetFeatures(ctx context.Context) ([]LinkedFeature_SdkV2, bool) {
	if o.Features.IsNull() || o.Features.IsUnknown() {
		return nil, false
	}
	var v []LinkedFeature_SdkV2
	d := o.Features.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFeatures sets the value of the Features field in FeatureList_SdkV2.
func (o *FeatureList_SdkV2) SetFeatures(ctx context.Context, v []LinkedFeature_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["features"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Features = types.ListValueMust(t, vs)
}

// Represents a tag on a feature in a feature table.
type FeatureTag_SdkV2 struct {
	Key types.String `tfsdk:"key"`

	Value types.String `tfsdk:"value"`
}

func (toState *FeatureTag_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan FeatureTag_SdkV2) {
}

func (toState *FeatureTag_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState FeatureTag_SdkV2) {
}

func (c FeatureTag_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a FeatureTag_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FeatureTag_SdkV2
// only implements ToObjectValue() and Type().
func (o FeatureTag_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FeatureTag_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *FileInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan FileInfo_SdkV2) {
}

func (toState *FileInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState FileInfo_SdkV2) {
}

func (c FileInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a FileInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FileInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o FileInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_size": o.FileSize,
			"is_dir":    o.IsDir,
			"path":      o.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FileInfo_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *FinalizeLoggedModelRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan FinalizeLoggedModelRequest_SdkV2) {
}

func (toState *FinalizeLoggedModelRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState FinalizeLoggedModelRequest_SdkV2) {
}

func (c FinalizeLoggedModelRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a FinalizeLoggedModelRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FinalizeLoggedModelRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o FinalizeLoggedModelRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_id": o.ModelId,
			"status":   o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FinalizeLoggedModelRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *FinalizeLoggedModelResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan FinalizeLoggedModelResponse_SdkV2) {
	if !fromPlan.Model.IsNull() && !fromPlan.Model.IsUnknown() {
		if toStateModel, ok := toState.GetModel(ctx); ok {
			if fromPlanModel, ok := fromPlan.GetModel(ctx); ok {
				toStateModel.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanModel)
				toState.SetModel(ctx, toStateModel)
			}
		}
	}
}

func (toState *FinalizeLoggedModelResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState FinalizeLoggedModelResponse_SdkV2) {
	if !fromState.Model.IsNull() && !fromState.Model.IsUnknown() {
		if toStateModel, ok := toState.GetModel(ctx); ok {
			if fromStateModel, ok := fromState.GetModel(ctx); ok {
				toStateModel.SyncFieldsDuringRead(ctx, fromStateModel)
				toState.SetModel(ctx, toStateModel)
			}
		}
	}
}

func (c FinalizeLoggedModelResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a FinalizeLoggedModelResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model": reflect.TypeOf(LoggedModel_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FinalizeLoggedModelResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o FinalizeLoggedModelResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model": o.Model,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FinalizeLoggedModelResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *FinalizeLoggedModelResponse_SdkV2) GetModel(ctx context.Context) (LoggedModel_SdkV2, bool) {
	var e LoggedModel_SdkV2
	if o.Model.IsNull() || o.Model.IsUnknown() {
		return e, false
	}
	var v []LoggedModel_SdkV2
	d := o.Model.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetModel sets the value of the Model field in FinalizeLoggedModelResponse_SdkV2.
func (o *FinalizeLoggedModelResponse_SdkV2) SetModel(ctx context.Context, v LoggedModel_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["model"]
	o.Model = types.ListValueMust(t, vs)
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

func (toState *ForecastingExperiment_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ForecastingExperiment_SdkV2) {
}

func (toState *ForecastingExperiment_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ForecastingExperiment_SdkV2) {
}

func (c ForecastingExperiment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ForecastingExperiment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ForecastingExperiment_SdkV2
// only implements ToObjectValue() and Type().
func (o ForecastingExperiment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id":       o.ExperimentId,
			"experiment_page_url": o.ExperimentPageUrl,
			"state":               o.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ForecastingExperiment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id":       types.StringType,
			"experiment_page_url": types.StringType,
			"state":               types.StringType,
		},
	}
}

type GetByNameRequest_SdkV2 struct {
	// Name of the associated experiment.
	ExperimentName types.String `tfsdk:"-"`
}

func (toState *GetByNameRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetByNameRequest_SdkV2) {
}

func (toState *GetByNameRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetByNameRequest_SdkV2) {
}

func (c GetByNameRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetByNameRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetByNameRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetByNameRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_name": o.ExperimentName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetByNameRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *GetExperimentByNameResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetExperimentByNameResponse_SdkV2) {
	if !fromPlan.Experiment.IsNull() && !fromPlan.Experiment.IsUnknown() {
		if toStateExperiment, ok := toState.GetExperiment(ctx); ok {
			if fromPlanExperiment, ok := fromPlan.GetExperiment(ctx); ok {
				toStateExperiment.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanExperiment)
				toState.SetExperiment(ctx, toStateExperiment)
			}
		}
	}
}

func (toState *GetExperimentByNameResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetExperimentByNameResponse_SdkV2) {
	if !fromState.Experiment.IsNull() && !fromState.Experiment.IsUnknown() {
		if toStateExperiment, ok := toState.GetExperiment(ctx); ok {
			if fromStateExperiment, ok := fromState.GetExperiment(ctx); ok {
				toStateExperiment.SyncFieldsDuringRead(ctx, fromStateExperiment)
				toState.SetExperiment(ctx, toStateExperiment)
			}
		}
	}
}

func (c GetExperimentByNameResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetExperimentByNameResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"experiment": reflect.TypeOf(Experiment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExperimentByNameResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetExperimentByNameResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment": o.Experiment,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetExperimentByNameResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *GetExperimentByNameResponse_SdkV2) GetExperiment(ctx context.Context) (Experiment_SdkV2, bool) {
	var e Experiment_SdkV2
	if o.Experiment.IsNull() || o.Experiment.IsUnknown() {
		return e, false
	}
	var v []Experiment_SdkV2
	d := o.Experiment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetExperiment sets the value of the Experiment field in GetExperimentByNameResponse_SdkV2.
func (o *GetExperimentByNameResponse_SdkV2) SetExperiment(ctx context.Context, v Experiment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["experiment"]
	o.Experiment = types.ListValueMust(t, vs)
}

type GetExperimentPermissionLevelsRequest_SdkV2 struct {
	// The experiment for which to get or manage permissions.
	ExperimentId types.String `tfsdk:"-"`
}

func (toState *GetExperimentPermissionLevelsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetExperimentPermissionLevelsRequest_SdkV2) {
}

func (toState *GetExperimentPermissionLevelsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetExperimentPermissionLevelsRequest_SdkV2) {
}

func (c GetExperimentPermissionLevelsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetExperimentPermissionLevelsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExperimentPermissionLevelsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetExperimentPermissionLevelsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": o.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetExperimentPermissionLevelsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *GetExperimentPermissionLevelsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetExperimentPermissionLevelsResponse_SdkV2) {
}

func (toState *GetExperimentPermissionLevelsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetExperimentPermissionLevelsResponse_SdkV2) {
}

func (c GetExperimentPermissionLevelsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetExperimentPermissionLevelsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(ExperimentPermissionsDescription_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExperimentPermissionLevelsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetExperimentPermissionLevelsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": o.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetExperimentPermissionLevelsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *GetExperimentPermissionLevelsResponse_SdkV2) GetPermissionLevels(ctx context.Context) ([]ExperimentPermissionsDescription_SdkV2, bool) {
	if o.PermissionLevels.IsNull() || o.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []ExperimentPermissionsDescription_SdkV2
	d := o.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetExperimentPermissionLevelsResponse_SdkV2.
func (o *GetExperimentPermissionLevelsResponse_SdkV2) SetPermissionLevels(ctx context.Context, v []ExperimentPermissionsDescription_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PermissionLevels = types.ListValueMust(t, vs)
}

type GetExperimentPermissionsRequest_SdkV2 struct {
	// The experiment for which to get or manage permissions.
	ExperimentId types.String `tfsdk:"-"`
}

func (toState *GetExperimentPermissionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetExperimentPermissionsRequest_SdkV2) {
}

func (toState *GetExperimentPermissionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetExperimentPermissionsRequest_SdkV2) {
}

func (c GetExperimentPermissionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetExperimentPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExperimentPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetExperimentPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": o.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetExperimentPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *GetExperimentRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetExperimentRequest_SdkV2) {
}

func (toState *GetExperimentRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetExperimentRequest_SdkV2) {
}

func (c GetExperimentRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetExperimentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExperimentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetExperimentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": o.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetExperimentRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *GetExperimentResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetExperimentResponse_SdkV2) {
	if !fromPlan.Experiment.IsNull() && !fromPlan.Experiment.IsUnknown() {
		if toStateExperiment, ok := toState.GetExperiment(ctx); ok {
			if fromPlanExperiment, ok := fromPlan.GetExperiment(ctx); ok {
				toStateExperiment.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanExperiment)
				toState.SetExperiment(ctx, toStateExperiment)
			}
		}
	}
}

func (toState *GetExperimentResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetExperimentResponse_SdkV2) {
	if !fromState.Experiment.IsNull() && !fromState.Experiment.IsUnknown() {
		if toStateExperiment, ok := toState.GetExperiment(ctx); ok {
			if fromStateExperiment, ok := fromState.GetExperiment(ctx); ok {
				toStateExperiment.SyncFieldsDuringRead(ctx, fromStateExperiment)
				toState.SetExperiment(ctx, toStateExperiment)
			}
		}
	}
}

func (c GetExperimentResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetExperimentResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"experiment": reflect.TypeOf(Experiment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExperimentResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetExperimentResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment": o.Experiment,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetExperimentResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *GetExperimentResponse_SdkV2) GetExperiment(ctx context.Context) (Experiment_SdkV2, bool) {
	var e Experiment_SdkV2
	if o.Experiment.IsNull() || o.Experiment.IsUnknown() {
		return e, false
	}
	var v []Experiment_SdkV2
	d := o.Experiment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetExperiment sets the value of the Experiment field in GetExperimentResponse_SdkV2.
func (o *GetExperimentResponse_SdkV2) SetExperiment(ctx context.Context, v Experiment_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["experiment"]
	o.Experiment = types.ListValueMust(t, vs)
}

type GetFeatureLineageRequest_SdkV2 struct {
	// The name of the feature.
	FeatureName types.String `tfsdk:"-"`
	// The full name of the feature table in Unity Catalog.
	TableName types.String `tfsdk:"-"`
}

func (toState *GetFeatureLineageRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetFeatureLineageRequest_SdkV2) {
}

func (toState *GetFeatureLineageRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetFeatureLineageRequest_SdkV2) {
}

func (c GetFeatureLineageRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetFeatureLineageRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetFeatureLineageRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetFeatureLineageRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_name": o.FeatureName,
			"table_name":   o.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetFeatureLineageRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"feature_name": types.StringType,
			"table_name":   types.StringType,
		},
	}
}

type GetFeatureTagRequest_SdkV2 struct {
	FeatureName types.String `tfsdk:"-"`

	Key types.String `tfsdk:"-"`

	TableName types.String `tfsdk:"-"`
}

func (toState *GetFeatureTagRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetFeatureTagRequest_SdkV2) {
}

func (toState *GetFeatureTagRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetFeatureTagRequest_SdkV2) {
}

func (c GetFeatureTagRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetFeatureTagRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetFeatureTagRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetFeatureTagRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_name": o.FeatureName,
			"key":          o.Key,
			"table_name":   o.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetFeatureTagRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *GetForecastingExperimentRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetForecastingExperimentRequest_SdkV2) {
}

func (toState *GetForecastingExperimentRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetForecastingExperimentRequest_SdkV2) {
}

func (c GetForecastingExperimentRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetForecastingExperimentRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetForecastingExperimentRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetForecastingExperimentRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": o.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetForecastingExperimentRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *GetHistoryRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetHistoryRequest_SdkV2) {
}

func (toState *GetHistoryRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetHistoryRequest_SdkV2) {
}

func (c GetHistoryRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetHistoryRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetHistoryRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetHistoryRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results": o.MaxResults,
			"metric_key":  o.MetricKey,
			"page_token":  o.PageToken,
			"run_id":      o.RunId,
			"run_uuid":    o.RunUuid,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetHistoryRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *GetLatestVersionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetLatestVersionsRequest_SdkV2) {
}

func (toState *GetLatestVersionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetLatestVersionsRequest_SdkV2) {
}

func (c GetLatestVersionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetLatestVersionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"stages": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLatestVersionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetLatestVersionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":   o.Name,
			"stages": o.Stages,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetLatestVersionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *GetLatestVersionsRequest_SdkV2) GetStages(ctx context.Context) ([]types.String, bool) {
	if o.Stages.IsNull() || o.Stages.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Stages.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStages sets the value of the Stages field in GetLatestVersionsRequest_SdkV2.
func (o *GetLatestVersionsRequest_SdkV2) SetStages(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["stages"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Stages = types.ListValueMust(t, vs)
}

type GetLatestVersionsResponse_SdkV2 struct {
	// Latest version models for each requests stage. Only return models with
	// current `READY` status. If no `stages` provided, returns the latest
	// version for each stage, including `"None"`.
	ModelVersions types.List `tfsdk:"model_versions"`
}

func (toState *GetLatestVersionsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetLatestVersionsResponse_SdkV2) {
}

func (toState *GetLatestVersionsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetLatestVersionsResponse_SdkV2) {
}

func (c GetLatestVersionsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetLatestVersionsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model_versions": reflect.TypeOf(ModelVersion_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLatestVersionsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetLatestVersionsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_versions": o.ModelVersions,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetLatestVersionsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *GetLatestVersionsResponse_SdkV2) GetModelVersions(ctx context.Context) ([]ModelVersion_SdkV2, bool) {
	if o.ModelVersions.IsNull() || o.ModelVersions.IsUnknown() {
		return nil, false
	}
	var v []ModelVersion_SdkV2
	d := o.ModelVersions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModelVersions sets the value of the ModelVersions field in GetLatestVersionsResponse_SdkV2.
func (o *GetLatestVersionsResponse_SdkV2) SetModelVersions(ctx context.Context, v []ModelVersion_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["model_versions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ModelVersions = types.ListValueMust(t, vs)
}

type GetLoggedModelRequest_SdkV2 struct {
	// The ID of the logged model to retrieve.
	ModelId types.String `tfsdk:"-"`
}

func (toState *GetLoggedModelRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetLoggedModelRequest_SdkV2) {
}

func (toState *GetLoggedModelRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetLoggedModelRequest_SdkV2) {
}

func (c GetLoggedModelRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetLoggedModelRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLoggedModelRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetLoggedModelRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_id": o.ModelId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetLoggedModelRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *GetLoggedModelResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetLoggedModelResponse_SdkV2) {
	if !fromPlan.Model.IsNull() && !fromPlan.Model.IsUnknown() {
		if toStateModel, ok := toState.GetModel(ctx); ok {
			if fromPlanModel, ok := fromPlan.GetModel(ctx); ok {
				toStateModel.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanModel)
				toState.SetModel(ctx, toStateModel)
			}
		}
	}
}

func (toState *GetLoggedModelResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetLoggedModelResponse_SdkV2) {
	if !fromState.Model.IsNull() && !fromState.Model.IsUnknown() {
		if toStateModel, ok := toState.GetModel(ctx); ok {
			if fromStateModel, ok := fromState.GetModel(ctx); ok {
				toStateModel.SyncFieldsDuringRead(ctx, fromStateModel)
				toState.SetModel(ctx, toStateModel)
			}
		}
	}
}

func (c GetLoggedModelResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetLoggedModelResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model": reflect.TypeOf(LoggedModel_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLoggedModelResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetLoggedModelResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model": o.Model,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetLoggedModelResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *GetLoggedModelResponse_SdkV2) GetModel(ctx context.Context) (LoggedModel_SdkV2, bool) {
	var e LoggedModel_SdkV2
	if o.Model.IsNull() || o.Model.IsUnknown() {
		return e, false
	}
	var v []LoggedModel_SdkV2
	d := o.Model.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetModel sets the value of the Model field in GetLoggedModelResponse_SdkV2.
func (o *GetLoggedModelResponse_SdkV2) SetModel(ctx context.Context, v LoggedModel_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["model"]
	o.Model = types.ListValueMust(t, vs)
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

func (toState *GetMetricHistoryResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetMetricHistoryResponse_SdkV2) {
}

func (toState *GetMetricHistoryResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetMetricHistoryResponse_SdkV2) {
}

func (c GetMetricHistoryResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetMetricHistoryResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metrics": reflect.TypeOf(Metric_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetMetricHistoryResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetMetricHistoryResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metrics":         o.Metrics,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetMetricHistoryResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *GetMetricHistoryResponse_SdkV2) GetMetrics(ctx context.Context) ([]Metric_SdkV2, bool) {
	if o.Metrics.IsNull() || o.Metrics.IsUnknown() {
		return nil, false
	}
	var v []Metric_SdkV2
	d := o.Metrics.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMetrics sets the value of the Metrics field in GetMetricHistoryResponse_SdkV2.
func (o *GetMetricHistoryResponse_SdkV2) SetMetrics(ctx context.Context, v []Metric_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["metrics"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Metrics = types.ListValueMust(t, vs)
}

type GetModelRequest_SdkV2 struct {
	// Registered model unique name identifier.
	Name types.String `tfsdk:"-"`
}

func (toState *GetModelRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetModelRequest_SdkV2) {
}

func (toState *GetModelRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetModelRequest_SdkV2) {
}

func (c GetModelRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetModelRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetModelRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetModelRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetModelRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetModelResponse_SdkV2 struct {
	RegisteredModelDatabricks types.List `tfsdk:"registered_model_databricks"`
}

func (toState *GetModelResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetModelResponse_SdkV2) {
	if !fromPlan.RegisteredModelDatabricks.IsNull() && !fromPlan.RegisteredModelDatabricks.IsUnknown() {
		if toStateRegisteredModelDatabricks, ok := toState.GetRegisteredModelDatabricks(ctx); ok {
			if fromPlanRegisteredModelDatabricks, ok := fromPlan.GetRegisteredModelDatabricks(ctx); ok {
				toStateRegisteredModelDatabricks.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanRegisteredModelDatabricks)
				toState.SetRegisteredModelDatabricks(ctx, toStateRegisteredModelDatabricks)
			}
		}
	}
}

func (toState *GetModelResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetModelResponse_SdkV2) {
	if !fromState.RegisteredModelDatabricks.IsNull() && !fromState.RegisteredModelDatabricks.IsUnknown() {
		if toStateRegisteredModelDatabricks, ok := toState.GetRegisteredModelDatabricks(ctx); ok {
			if fromStateRegisteredModelDatabricks, ok := fromState.GetRegisteredModelDatabricks(ctx); ok {
				toStateRegisteredModelDatabricks.SyncFieldsDuringRead(ctx, fromStateRegisteredModelDatabricks)
				toState.SetRegisteredModelDatabricks(ctx, toStateRegisteredModelDatabricks)
			}
		}
	}
}

func (c GetModelResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetModelResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"registered_model_databricks": reflect.TypeOf(ModelDatabricks_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetModelResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetModelResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"registered_model_databricks": o.RegisteredModelDatabricks,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetModelResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *GetModelResponse_SdkV2) GetRegisteredModelDatabricks(ctx context.Context) (ModelDatabricks_SdkV2, bool) {
	var e ModelDatabricks_SdkV2
	if o.RegisteredModelDatabricks.IsNull() || o.RegisteredModelDatabricks.IsUnknown() {
		return e, false
	}
	var v []ModelDatabricks_SdkV2
	d := o.RegisteredModelDatabricks.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRegisteredModelDatabricks sets the value of the RegisteredModelDatabricks field in GetModelResponse_SdkV2.
func (o *GetModelResponse_SdkV2) SetRegisteredModelDatabricks(ctx context.Context, v ModelDatabricks_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["registered_model_databricks"]
	o.RegisteredModelDatabricks = types.ListValueMust(t, vs)
}

type GetModelVersionDownloadUriRequest_SdkV2 struct {
	// Name of the registered model
	Name types.String `tfsdk:"-"`
	// Model version number
	Version types.String `tfsdk:"-"`
}

func (toState *GetModelVersionDownloadUriRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetModelVersionDownloadUriRequest_SdkV2) {
}

func (toState *GetModelVersionDownloadUriRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetModelVersionDownloadUriRequest_SdkV2) {
}

func (c GetModelVersionDownloadUriRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetModelVersionDownloadUriRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetModelVersionDownloadUriRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetModelVersionDownloadUriRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":    o.Name,
			"version": o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetModelVersionDownloadUriRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *GetModelVersionDownloadUriResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetModelVersionDownloadUriResponse_SdkV2) {
}

func (toState *GetModelVersionDownloadUriResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetModelVersionDownloadUriResponse_SdkV2) {
}

func (c GetModelVersionDownloadUriResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetModelVersionDownloadUriResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetModelVersionDownloadUriResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetModelVersionDownloadUriResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"artifact_uri": o.ArtifactUri,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetModelVersionDownloadUriResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *GetModelVersionRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetModelVersionRequest_SdkV2) {
}

func (toState *GetModelVersionRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetModelVersionRequest_SdkV2) {
}

func (c GetModelVersionRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetModelVersionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetModelVersionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetModelVersionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":    o.Name,
			"version": o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetModelVersionRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *GetModelVersionResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetModelVersionResponse_SdkV2) {
	if !fromPlan.ModelVersion.IsNull() && !fromPlan.ModelVersion.IsUnknown() {
		if toStateModelVersion, ok := toState.GetModelVersion(ctx); ok {
			if fromPlanModelVersion, ok := fromPlan.GetModelVersion(ctx); ok {
				toStateModelVersion.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanModelVersion)
				toState.SetModelVersion(ctx, toStateModelVersion)
			}
		}
	}
}

func (toState *GetModelVersionResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetModelVersionResponse_SdkV2) {
	if !fromState.ModelVersion.IsNull() && !fromState.ModelVersion.IsUnknown() {
		if toStateModelVersion, ok := toState.GetModelVersion(ctx); ok {
			if fromStateModelVersion, ok := fromState.GetModelVersion(ctx); ok {
				toStateModelVersion.SyncFieldsDuringRead(ctx, fromStateModelVersion)
				toState.SetModelVersion(ctx, toStateModelVersion)
			}
		}
	}
}

func (c GetModelVersionResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetModelVersionResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model_version": reflect.TypeOf(ModelVersion_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetModelVersionResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetModelVersionResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_version": o.ModelVersion,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetModelVersionResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *GetModelVersionResponse_SdkV2) GetModelVersion(ctx context.Context) (ModelVersion_SdkV2, bool) {
	var e ModelVersion_SdkV2
	if o.ModelVersion.IsNull() || o.ModelVersion.IsUnknown() {
		return e, false
	}
	var v []ModelVersion_SdkV2
	d := o.ModelVersion.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetModelVersion sets the value of the ModelVersion field in GetModelVersionResponse_SdkV2.
func (o *GetModelVersionResponse_SdkV2) SetModelVersion(ctx context.Context, v ModelVersion_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["model_version"]
	o.ModelVersion = types.ListValueMust(t, vs)
}

type GetOnlineStoreRequest_SdkV2 struct {
	// Name of the online store to get.
	Name types.String `tfsdk:"-"`
}

func (toState *GetOnlineStoreRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetOnlineStoreRequest_SdkV2) {
}

func (toState *GetOnlineStoreRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetOnlineStoreRequest_SdkV2) {
}

func (c GetOnlineStoreRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetOnlineStoreRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetOnlineStoreRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetOnlineStoreRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetOnlineStoreRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *GetRegisteredModelPermissionLevelsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetRegisteredModelPermissionLevelsRequest_SdkV2) {
}

func (toState *GetRegisteredModelPermissionLevelsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetRegisteredModelPermissionLevelsRequest_SdkV2) {
}

func (c GetRegisteredModelPermissionLevelsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetRegisteredModelPermissionLevelsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRegisteredModelPermissionLevelsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetRegisteredModelPermissionLevelsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"registered_model_id": o.RegisteredModelId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetRegisteredModelPermissionLevelsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *GetRegisteredModelPermissionLevelsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetRegisteredModelPermissionLevelsResponse_SdkV2) {
}

func (toState *GetRegisteredModelPermissionLevelsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetRegisteredModelPermissionLevelsResponse_SdkV2) {
}

func (c GetRegisteredModelPermissionLevelsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetRegisteredModelPermissionLevelsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(RegisteredModelPermissionsDescription_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRegisteredModelPermissionLevelsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetRegisteredModelPermissionLevelsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": o.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetRegisteredModelPermissionLevelsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *GetRegisteredModelPermissionLevelsResponse_SdkV2) GetPermissionLevels(ctx context.Context) ([]RegisteredModelPermissionsDescription_SdkV2, bool) {
	if o.PermissionLevels.IsNull() || o.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []RegisteredModelPermissionsDescription_SdkV2
	d := o.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetRegisteredModelPermissionLevelsResponse_SdkV2.
func (o *GetRegisteredModelPermissionLevelsResponse_SdkV2) SetPermissionLevels(ctx context.Context, v []RegisteredModelPermissionsDescription_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PermissionLevels = types.ListValueMust(t, vs)
}

type GetRegisteredModelPermissionsRequest_SdkV2 struct {
	// The registered model for which to get or manage permissions.
	RegisteredModelId types.String `tfsdk:"-"`
}

func (toState *GetRegisteredModelPermissionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetRegisteredModelPermissionsRequest_SdkV2) {
}

func (toState *GetRegisteredModelPermissionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetRegisteredModelPermissionsRequest_SdkV2) {
}

func (c GetRegisteredModelPermissionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetRegisteredModelPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRegisteredModelPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetRegisteredModelPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"registered_model_id": o.RegisteredModelId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetRegisteredModelPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *GetRunRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetRunRequest_SdkV2) {
}

func (toState *GetRunRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetRunRequest_SdkV2) {
}

func (c GetRunRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetRunRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRunRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetRunRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run_id":   o.RunId,
			"run_uuid": o.RunUuid,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetRunRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *GetRunResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetRunResponse_SdkV2) {
	if !fromPlan.Run.IsNull() && !fromPlan.Run.IsUnknown() {
		if toStateRun, ok := toState.GetRun(ctx); ok {
			if fromPlanRun, ok := fromPlan.GetRun(ctx); ok {
				toStateRun.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanRun)
				toState.SetRun(ctx, toStateRun)
			}
		}
	}
}

func (toState *GetRunResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetRunResponse_SdkV2) {
	if !fromState.Run.IsNull() && !fromState.Run.IsUnknown() {
		if toStateRun, ok := toState.GetRun(ctx); ok {
			if fromStateRun, ok := fromState.GetRun(ctx); ok {
				toStateRun.SyncFieldsDuringRead(ctx, fromStateRun)
				toState.SetRun(ctx, toStateRun)
			}
		}
	}
}

func (c GetRunResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetRunResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"run": reflect.TypeOf(Run_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRunResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetRunResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run": o.Run,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetRunResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *GetRunResponse_SdkV2) GetRun(ctx context.Context) (Run_SdkV2, bool) {
	var e Run_SdkV2
	if o.Run.IsNull() || o.Run.IsUnknown() {
		return e, false
	}
	var v []Run_SdkV2
	d := o.Run.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRun sets the value of the Run field in GetRunResponse_SdkV2.
func (o *GetRunResponse_SdkV2) SetRun(ctx context.Context, v Run_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["run"]
	o.Run = types.ListValueMust(t, vs)
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

func (toState *HttpUrlSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan HttpUrlSpec_SdkV2) {
}

func (toState *HttpUrlSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState HttpUrlSpec_SdkV2) {
}

func (c HttpUrlSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a HttpUrlSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, HttpUrlSpec_SdkV2
// only implements ToObjectValue() and Type().
func (o HttpUrlSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"authorization":           o.Authorization,
			"enable_ssl_verification": o.EnableSslVerification,
			"secret":                  o.Secret,
			"url":                     o.Url,
		})
}

// Type implements basetypes.ObjectValuable.
func (o HttpUrlSpec_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *HttpUrlSpecWithoutSecret_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan HttpUrlSpecWithoutSecret_SdkV2) {
}

func (toState *HttpUrlSpecWithoutSecret_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState HttpUrlSpecWithoutSecret_SdkV2) {
}

func (c HttpUrlSpecWithoutSecret_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a HttpUrlSpecWithoutSecret_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, HttpUrlSpecWithoutSecret_SdkV2
// only implements ToObjectValue() and Type().
func (o HttpUrlSpecWithoutSecret_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enable_ssl_verification": o.EnableSslVerification,
			"url":                     o.Url,
		})
}

// Type implements basetypes.ObjectValuable.
func (o HttpUrlSpecWithoutSecret_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *InputTag_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan InputTag_SdkV2) {
}

func (toState *InputTag_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState InputTag_SdkV2) {
}

func (c InputTag_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a InputTag_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InputTag_SdkV2
// only implements ToObjectValue() and Type().
func (o InputTag_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o InputTag_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *JobSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan JobSpec_SdkV2) {
}

func (toState *JobSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState JobSpec_SdkV2) {
}

func (c JobSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a JobSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, JobSpec_SdkV2
// only implements ToObjectValue() and Type().
func (o JobSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_token":  o.AccessToken,
			"job_id":        o.JobId,
			"workspace_url": o.WorkspaceUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (o JobSpec_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *JobSpecWithoutSecret_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan JobSpecWithoutSecret_SdkV2) {
}

func (toState *JobSpecWithoutSecret_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState JobSpecWithoutSecret_SdkV2) {
}

func (c JobSpecWithoutSecret_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a JobSpecWithoutSecret_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, JobSpecWithoutSecret_SdkV2
// only implements ToObjectValue() and Type().
func (o JobSpecWithoutSecret_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"job_id":        o.JobId,
			"workspace_url": o.WorkspaceUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (o JobSpecWithoutSecret_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *LinkedFeature_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan LinkedFeature_SdkV2) {
}

func (toState *LinkedFeature_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState LinkedFeature_SdkV2) {
}

func (c LinkedFeature_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a LinkedFeature_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LinkedFeature_SdkV2
// only implements ToObjectValue() and Type().
func (o LinkedFeature_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_name":       o.FeatureName,
			"feature_table_id":   o.FeatureTableId,
			"feature_table_name": o.FeatureTableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LinkedFeature_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *ListArtifactsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListArtifactsRequest_SdkV2) {
}

func (toState *ListArtifactsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListArtifactsRequest_SdkV2) {
}

func (c ListArtifactsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListArtifactsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListArtifactsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListArtifactsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_token": o.PageToken,
			"path":       o.Path,
			"run_id":     o.RunId,
			"run_uuid":   o.RunUuid,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListArtifactsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *ListArtifactsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListArtifactsResponse_SdkV2) {
}

func (toState *ListArtifactsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListArtifactsResponse_SdkV2) {
}

func (c ListArtifactsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListArtifactsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"files": reflect.TypeOf(FileInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListArtifactsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListArtifactsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"files":           o.Files,
			"next_page_token": o.NextPageToken,
			"root_uri":        o.RootUri,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListArtifactsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ListArtifactsResponse_SdkV2) GetFiles(ctx context.Context) ([]FileInfo_SdkV2, bool) {
	if o.Files.IsNull() || o.Files.IsUnknown() {
		return nil, false
	}
	var v []FileInfo_SdkV2
	d := o.Files.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFiles sets the value of the Files field in ListArtifactsResponse_SdkV2.
func (o *ListArtifactsResponse_SdkV2) SetFiles(ctx context.Context, v []FileInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["files"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Files = types.ListValueMust(t, vs)
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

func (toState *ListExperimentsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListExperimentsRequest_SdkV2) {
}

func (toState *ListExperimentsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListExperimentsRequest_SdkV2) {
}

func (c ListExperimentsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListExperimentsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExperimentsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListExperimentsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results": o.MaxResults,
			"page_token":  o.PageToken,
			"view_type":   o.ViewType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListExperimentsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *ListExperimentsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListExperimentsResponse_SdkV2) {
}

func (toState *ListExperimentsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListExperimentsResponse_SdkV2) {
}

func (c ListExperimentsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListExperimentsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"experiments": reflect.TypeOf(Experiment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExperimentsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListExperimentsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiments":     o.Experiments,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListExperimentsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ListExperimentsResponse_SdkV2) GetExperiments(ctx context.Context) ([]Experiment_SdkV2, bool) {
	if o.Experiments.IsNull() || o.Experiments.IsUnknown() {
		return nil, false
	}
	var v []Experiment_SdkV2
	d := o.Experiments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExperiments sets the value of the Experiments field in ListExperimentsResponse_SdkV2.
func (o *ListExperimentsResponse_SdkV2) SetExperiments(ctx context.Context, v []Experiment_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["experiments"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Experiments = types.ListValueMust(t, vs)
}

type ListFeatureTagsRequest_SdkV2 struct {
	FeatureName types.String `tfsdk:"-"`
	// The maximum number of results to return.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page based on a previous query.
	PageToken types.String `tfsdk:"-"`

	TableName types.String `tfsdk:"-"`
}

func (toState *ListFeatureTagsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListFeatureTagsRequest_SdkV2) {
}

func (toState *ListFeatureTagsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListFeatureTagsRequest_SdkV2) {
}

func (c ListFeatureTagsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListFeatureTagsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFeatureTagsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListFeatureTagsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_name": o.FeatureName,
			"page_size":    o.PageSize,
			"page_token":   o.PageToken,
			"table_name":   o.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListFeatureTagsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *ListFeatureTagsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListFeatureTagsResponse_SdkV2) {
}

func (toState *ListFeatureTagsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListFeatureTagsResponse_SdkV2) {
}

func (c ListFeatureTagsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListFeatureTagsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"feature_tags": reflect.TypeOf(FeatureTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFeatureTagsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListFeatureTagsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_tags":    o.FeatureTags,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListFeatureTagsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ListFeatureTagsResponse_SdkV2) GetFeatureTags(ctx context.Context) ([]FeatureTag_SdkV2, bool) {
	if o.FeatureTags.IsNull() || o.FeatureTags.IsUnknown() {
		return nil, false
	}
	var v []FeatureTag_SdkV2
	d := o.FeatureTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFeatureTags sets the value of the FeatureTags field in ListFeatureTagsResponse_SdkV2.
func (o *ListFeatureTagsResponse_SdkV2) SetFeatureTags(ctx context.Context, v []FeatureTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["feature_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.FeatureTags = types.ListValueMust(t, vs)
}

type ListModelsRequest_SdkV2 struct {
	// Maximum number of registered models desired. Max threshold is 1000.
	MaxResults types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page based on a previous query.
	PageToken types.String `tfsdk:"-"`
}

func (toState *ListModelsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListModelsRequest_SdkV2) {
}

func (toState *ListModelsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListModelsRequest_SdkV2) {
}

func (c ListModelsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListModelsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListModelsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListModelsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results": o.MaxResults,
			"page_token":  o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListModelsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *ListModelsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListModelsResponse_SdkV2) {
}

func (toState *ListModelsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListModelsResponse_SdkV2) {
}

func (c ListModelsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListModelsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"registered_models": reflect.TypeOf(Model_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListModelsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListModelsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":   o.NextPageToken,
			"registered_models": o.RegisteredModels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListModelsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ListModelsResponse_SdkV2) GetRegisteredModels(ctx context.Context) ([]Model_SdkV2, bool) {
	if o.RegisteredModels.IsNull() || o.RegisteredModels.IsUnknown() {
		return nil, false
	}
	var v []Model_SdkV2
	d := o.RegisteredModels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRegisteredModels sets the value of the RegisteredModels field in ListModelsResponse_SdkV2.
func (o *ListModelsResponse_SdkV2) SetRegisteredModels(ctx context.Context, v []Model_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["registered_models"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RegisteredModels = types.ListValueMust(t, vs)
}

type ListOnlineStoresRequest_SdkV2 struct {
	// The maximum number of results to return. Defaults to 100 if not
	// specified.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page based on a previous query.
	PageToken types.String `tfsdk:"-"`
}

func (toState *ListOnlineStoresRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListOnlineStoresRequest_SdkV2) {
}

func (toState *ListOnlineStoresRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListOnlineStoresRequest_SdkV2) {
}

func (c ListOnlineStoresRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListOnlineStoresRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListOnlineStoresRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListOnlineStoresRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListOnlineStoresRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *ListOnlineStoresResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListOnlineStoresResponse_SdkV2) {
}

func (toState *ListOnlineStoresResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListOnlineStoresResponse_SdkV2) {
}

func (c ListOnlineStoresResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListOnlineStoresResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"online_stores": reflect.TypeOf(OnlineStore_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListOnlineStoresResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListOnlineStoresResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"online_stores":   o.OnlineStores,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListOnlineStoresResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ListOnlineStoresResponse_SdkV2) GetOnlineStores(ctx context.Context) ([]OnlineStore_SdkV2, bool) {
	if o.OnlineStores.IsNull() || o.OnlineStores.IsUnknown() {
		return nil, false
	}
	var v []OnlineStore_SdkV2
	d := o.OnlineStores.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOnlineStores sets the value of the OnlineStores field in ListOnlineStoresResponse_SdkV2.
func (o *ListOnlineStoresResponse_SdkV2) SetOnlineStores(ctx context.Context, v []OnlineStore_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["online_stores"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OnlineStores = types.ListValueMust(t, vs)
}

type ListRegistryWebhooks_SdkV2 struct {
	// Token that can be used to retrieve the next page of artifact results
	NextPageToken types.String `tfsdk:"next_page_token"`
	// Array of registry webhooks.
	Webhooks types.List `tfsdk:"webhooks"`
}

func (toState *ListRegistryWebhooks_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListRegistryWebhooks_SdkV2) {
}

func (toState *ListRegistryWebhooks_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListRegistryWebhooks_SdkV2) {
}

func (c ListRegistryWebhooks_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListRegistryWebhooks_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"webhooks": reflect.TypeOf(RegistryWebhook_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListRegistryWebhooks_SdkV2
// only implements ToObjectValue() and Type().
func (o ListRegistryWebhooks_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"webhooks":        o.Webhooks,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListRegistryWebhooks_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ListRegistryWebhooks_SdkV2) GetWebhooks(ctx context.Context) ([]RegistryWebhook_SdkV2, bool) {
	if o.Webhooks.IsNull() || o.Webhooks.IsUnknown() {
		return nil, false
	}
	var v []RegistryWebhook_SdkV2
	d := o.Webhooks.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWebhooks sets the value of the Webhooks field in ListRegistryWebhooks_SdkV2.
func (o *ListRegistryWebhooks_SdkV2) SetWebhooks(ctx context.Context, v []RegistryWebhook_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["webhooks"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Webhooks = types.ListValueMust(t, vs)
}

type ListTransitionRequestsRequest_SdkV2 struct {
	// Name of the registered model.
	Name types.String `tfsdk:"-"`
	// Version of the model.
	Version types.String `tfsdk:"-"`
}

func (toState *ListTransitionRequestsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListTransitionRequestsRequest_SdkV2) {
}

func (toState *ListTransitionRequestsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListTransitionRequestsRequest_SdkV2) {
}

func (c ListTransitionRequestsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListTransitionRequestsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTransitionRequestsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListTransitionRequestsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":    o.Name,
			"version": o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListTransitionRequestsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *ListTransitionRequestsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListTransitionRequestsResponse_SdkV2) {
}

func (toState *ListTransitionRequestsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListTransitionRequestsResponse_SdkV2) {
}

func (c ListTransitionRequestsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListTransitionRequestsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"requests": reflect.TypeOf(Activity_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTransitionRequestsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListTransitionRequestsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"requests": o.Requests,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListTransitionRequestsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ListTransitionRequestsResponse_SdkV2) GetRequests(ctx context.Context) ([]Activity_SdkV2, bool) {
	if o.Requests.IsNull() || o.Requests.IsUnknown() {
		return nil, false
	}
	var v []Activity_SdkV2
	d := o.Requests.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRequests sets the value of the Requests field in ListTransitionRequestsResponse_SdkV2.
func (o *ListTransitionRequestsResponse_SdkV2) SetRequests(ctx context.Context, v []Activity_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["requests"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Requests = types.ListValueMust(t, vs)
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

func (toState *ListWebhooksRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListWebhooksRequest_SdkV2) {
}

func (toState *ListWebhooksRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListWebhooksRequest_SdkV2) {
}

func (c ListWebhooksRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListWebhooksRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"events": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWebhooksRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListWebhooksRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"events":      o.Events,
			"max_results": o.MaxResults,
			"model_name":  o.ModelName,
			"page_token":  o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListWebhooksRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ListWebhooksRequest_SdkV2) GetEvents(ctx context.Context) ([]types.String, bool) {
	if o.Events.IsNull() || o.Events.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Events.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEvents sets the value of the Events field in ListWebhooksRequest_SdkV2.
func (o *ListWebhooksRequest_SdkV2) SetEvents(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["events"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Events = types.ListValueMust(t, vs)
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

func (toState *LogBatch_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan LogBatch_SdkV2) {
}

func (toState *LogBatch_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState LogBatch_SdkV2) {
}

func (c LogBatch_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a LogBatch_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metrics": reflect.TypeOf(Metric_SdkV2{}),
		"params":  reflect.TypeOf(Param_SdkV2{}),
		"tags":    reflect.TypeOf(RunTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogBatch_SdkV2
// only implements ToObjectValue() and Type().
func (o LogBatch_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metrics": o.Metrics,
			"params":  o.Params,
			"run_id":  o.RunId,
			"tags":    o.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LogBatch_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *LogBatch_SdkV2) GetMetrics(ctx context.Context) ([]Metric_SdkV2, bool) {
	if o.Metrics.IsNull() || o.Metrics.IsUnknown() {
		return nil, false
	}
	var v []Metric_SdkV2
	d := o.Metrics.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMetrics sets the value of the Metrics field in LogBatch_SdkV2.
func (o *LogBatch_SdkV2) SetMetrics(ctx context.Context, v []Metric_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["metrics"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Metrics = types.ListValueMust(t, vs)
}

// GetParams returns the value of the Params field in LogBatch_SdkV2 as
// a slice of Param_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *LogBatch_SdkV2) GetParams(ctx context.Context) ([]Param_SdkV2, bool) {
	if o.Params.IsNull() || o.Params.IsUnknown() {
		return nil, false
	}
	var v []Param_SdkV2
	d := o.Params.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParams sets the value of the Params field in LogBatch_SdkV2.
func (o *LogBatch_SdkV2) SetParams(ctx context.Context, v []Param_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Params = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in LogBatch_SdkV2 as
// a slice of RunTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *LogBatch_SdkV2) GetTags(ctx context.Context) ([]RunTag_SdkV2, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []RunTag_SdkV2
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in LogBatch_SdkV2.
func (o *LogBatch_SdkV2) SetTags(ctx context.Context, v []RunTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type LogBatchResponse_SdkV2 struct {
}

func (toState *LogBatchResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan LogBatchResponse_SdkV2) {
}

func (toState *LogBatchResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState LogBatchResponse_SdkV2) {
}

func (c LogBatchResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogBatchResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a LogBatchResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogBatchResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o LogBatchResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o LogBatchResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *LogInputs_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan LogInputs_SdkV2) {
}

func (toState *LogInputs_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState LogInputs_SdkV2) {
}

func (c LogInputs_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a LogInputs_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"datasets": reflect.TypeOf(DatasetInput_SdkV2{}),
		"models":   reflect.TypeOf(ModelInput_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogInputs_SdkV2
// only implements ToObjectValue() and Type().
func (o LogInputs_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"datasets": o.Datasets,
			"models":   o.Models,
			"run_id":   o.RunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LogInputs_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *LogInputs_SdkV2) GetDatasets(ctx context.Context) ([]DatasetInput_SdkV2, bool) {
	if o.Datasets.IsNull() || o.Datasets.IsUnknown() {
		return nil, false
	}
	var v []DatasetInput_SdkV2
	d := o.Datasets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatasets sets the value of the Datasets field in LogInputs_SdkV2.
func (o *LogInputs_SdkV2) SetDatasets(ctx context.Context, v []DatasetInput_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["datasets"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Datasets = types.ListValueMust(t, vs)
}

// GetModels returns the value of the Models field in LogInputs_SdkV2 as
// a slice of ModelInput_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *LogInputs_SdkV2) GetModels(ctx context.Context) ([]ModelInput_SdkV2, bool) {
	if o.Models.IsNull() || o.Models.IsUnknown() {
		return nil, false
	}
	var v []ModelInput_SdkV2
	d := o.Models.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModels sets the value of the Models field in LogInputs_SdkV2.
func (o *LogInputs_SdkV2) SetModels(ctx context.Context, v []ModelInput_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["models"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Models = types.ListValueMust(t, vs)
}

type LogInputsResponse_SdkV2 struct {
}

func (toState *LogInputsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan LogInputsResponse_SdkV2) {
}

func (toState *LogInputsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState LogInputsResponse_SdkV2) {
}

func (c LogInputsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogInputsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a LogInputsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogInputsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o LogInputsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o LogInputsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *LogLoggedModelParamsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan LogLoggedModelParamsRequest_SdkV2) {
}

func (toState *LogLoggedModelParamsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState LogLoggedModelParamsRequest_SdkV2) {
}

func (c LogLoggedModelParamsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a LogLoggedModelParamsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"params": reflect.TypeOf(LoggedModelParameter_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogLoggedModelParamsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o LogLoggedModelParamsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_id": o.ModelId,
			"params":   o.Params,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LogLoggedModelParamsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *LogLoggedModelParamsRequest_SdkV2) GetParams(ctx context.Context) ([]LoggedModelParameter_SdkV2, bool) {
	if o.Params.IsNull() || o.Params.IsUnknown() {
		return nil, false
	}
	var v []LoggedModelParameter_SdkV2
	d := o.Params.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParams sets the value of the Params field in LogLoggedModelParamsRequest_SdkV2.
func (o *LogLoggedModelParamsRequest_SdkV2) SetParams(ctx context.Context, v []LoggedModelParameter_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Params = types.ListValueMust(t, vs)
}

type LogLoggedModelParamsRequestResponse_SdkV2 struct {
}

func (toState *LogLoggedModelParamsRequestResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan LogLoggedModelParamsRequestResponse_SdkV2) {
}

func (toState *LogLoggedModelParamsRequestResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState LogLoggedModelParamsRequestResponse_SdkV2) {
}

func (c LogLoggedModelParamsRequestResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogLoggedModelParamsRequestResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a LogLoggedModelParamsRequestResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogLoggedModelParamsRequestResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o LogLoggedModelParamsRequestResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o LogLoggedModelParamsRequestResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *LogMetric_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan LogMetric_SdkV2) {
}

func (toState *LogMetric_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState LogMetric_SdkV2) {
}

func (c LogMetric_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a LogMetric_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogMetric_SdkV2
// only implements ToObjectValue() and Type().
func (o LogMetric_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dataset_digest": o.DatasetDigest,
			"dataset_name":   o.DatasetName,
			"key":            o.Key,
			"model_id":       o.ModelId,
			"run_id":         o.RunId,
			"run_uuid":       o.RunUuid,
			"step":           o.Step,
			"timestamp":      o.Timestamp,
			"value":          o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LogMetric_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *LogMetricResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan LogMetricResponse_SdkV2) {
}

func (toState *LogMetricResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState LogMetricResponse_SdkV2) {
}

func (c LogMetricResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogMetricResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a LogMetricResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogMetricResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o LogMetricResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o LogMetricResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *LogModel_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan LogModel_SdkV2) {
}

func (toState *LogModel_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState LogModel_SdkV2) {
}

func (c LogModel_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a LogModel_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogModel_SdkV2
// only implements ToObjectValue() and Type().
func (o LogModel_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_json": o.ModelJson,
			"run_id":     o.RunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LogModel_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_json": types.StringType,
			"run_id":     types.StringType,
		},
	}
}

type LogModelResponse_SdkV2 struct {
}

func (toState *LogModelResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan LogModelResponse_SdkV2) {
}

func (toState *LogModelResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState LogModelResponse_SdkV2) {
}

func (c LogModelResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogModelResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a LogModelResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogModelResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o LogModelResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o LogModelResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *LogOutputsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan LogOutputsRequest_SdkV2) {
}

func (toState *LogOutputsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState LogOutputsRequest_SdkV2) {
}

func (c LogOutputsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a LogOutputsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"models": reflect.TypeOf(ModelOutput_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogOutputsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o LogOutputsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"models": o.Models,
			"run_id": o.RunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LogOutputsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *LogOutputsRequest_SdkV2) GetModels(ctx context.Context) ([]ModelOutput_SdkV2, bool) {
	if o.Models.IsNull() || o.Models.IsUnknown() {
		return nil, false
	}
	var v []ModelOutput_SdkV2
	d := o.Models.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModels sets the value of the Models field in LogOutputsRequest_SdkV2.
func (o *LogOutputsRequest_SdkV2) SetModels(ctx context.Context, v []ModelOutput_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["models"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Models = types.ListValueMust(t, vs)
}

type LogOutputsResponse_SdkV2 struct {
}

func (toState *LogOutputsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan LogOutputsResponse_SdkV2) {
}

func (toState *LogOutputsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState LogOutputsResponse_SdkV2) {
}

func (c LogOutputsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogOutputsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a LogOutputsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogOutputsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o LogOutputsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o LogOutputsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *LogParam_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan LogParam_SdkV2) {
}

func (toState *LogParam_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState LogParam_SdkV2) {
}

func (c LogParam_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a LogParam_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogParam_SdkV2
// only implements ToObjectValue() and Type().
func (o LogParam_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":      o.Key,
			"run_id":   o.RunId,
			"run_uuid": o.RunUuid,
			"value":    o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LogParam_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *LogParamResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan LogParamResponse_SdkV2) {
}

func (toState *LogParamResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState LogParamResponse_SdkV2) {
}

func (c LogParamResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogParamResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a LogParamResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogParamResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o LogParamResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o LogParamResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *LoggedModel_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan LoggedModel_SdkV2) {
	if !fromPlan.Data.IsNull() && !fromPlan.Data.IsUnknown() {
		if toStateData, ok := toState.GetData(ctx); ok {
			if fromPlanData, ok := fromPlan.GetData(ctx); ok {
				toStateData.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanData)
				toState.SetData(ctx, toStateData)
			}
		}
	}
	if !fromPlan.Info.IsNull() && !fromPlan.Info.IsUnknown() {
		if toStateInfo, ok := toState.GetInfo(ctx); ok {
			if fromPlanInfo, ok := fromPlan.GetInfo(ctx); ok {
				toStateInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanInfo)
				toState.SetInfo(ctx, toStateInfo)
			}
		}
	}
}

func (toState *LoggedModel_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState LoggedModel_SdkV2) {
	if !fromState.Data.IsNull() && !fromState.Data.IsUnknown() {
		if toStateData, ok := toState.GetData(ctx); ok {
			if fromStateData, ok := fromState.GetData(ctx); ok {
				toStateData.SyncFieldsDuringRead(ctx, fromStateData)
				toState.SetData(ctx, toStateData)
			}
		}
	}
	if !fromState.Info.IsNull() && !fromState.Info.IsUnknown() {
		if toStateInfo, ok := toState.GetInfo(ctx); ok {
			if fromStateInfo, ok := fromState.GetInfo(ctx); ok {
				toStateInfo.SyncFieldsDuringRead(ctx, fromStateInfo)
				toState.SetInfo(ctx, toStateInfo)
			}
		}
	}
}

func (c LoggedModel_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a LoggedModel_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data": reflect.TypeOf(LoggedModelData_SdkV2{}),
		"info": reflect.TypeOf(LoggedModelInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LoggedModel_SdkV2
// only implements ToObjectValue() and Type().
func (o LoggedModel_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data": o.Data,
			"info": o.Info,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LoggedModel_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *LoggedModel_SdkV2) GetData(ctx context.Context) (LoggedModelData_SdkV2, bool) {
	var e LoggedModelData_SdkV2
	if o.Data.IsNull() || o.Data.IsUnknown() {
		return e, false
	}
	var v []LoggedModelData_SdkV2
	d := o.Data.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetData sets the value of the Data field in LoggedModel_SdkV2.
func (o *LoggedModel_SdkV2) SetData(ctx context.Context, v LoggedModelData_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["data"]
	o.Data = types.ListValueMust(t, vs)
}

// GetInfo returns the value of the Info field in LoggedModel_SdkV2 as
// a LoggedModelInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *LoggedModel_SdkV2) GetInfo(ctx context.Context) (LoggedModelInfo_SdkV2, bool) {
	var e LoggedModelInfo_SdkV2
	if o.Info.IsNull() || o.Info.IsUnknown() {
		return e, false
	}
	var v []LoggedModelInfo_SdkV2
	d := o.Info.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInfo sets the value of the Info field in LoggedModel_SdkV2.
func (o *LoggedModel_SdkV2) SetInfo(ctx context.Context, v LoggedModelInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["info"]
	o.Info = types.ListValueMust(t, vs)
}

// A LoggedModelData message includes logged model params and linked metrics.
type LoggedModelData_SdkV2 struct {
	// Performance metrics linked to the model.
	Metrics types.List `tfsdk:"metrics"`
	// Immutable string key-value pairs of the model.
	Params types.List `tfsdk:"params"`
}

func (toState *LoggedModelData_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan LoggedModelData_SdkV2) {
}

func (toState *LoggedModelData_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState LoggedModelData_SdkV2) {
}

func (c LoggedModelData_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a LoggedModelData_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metrics": reflect.TypeOf(Metric_SdkV2{}),
		"params":  reflect.TypeOf(LoggedModelParameter_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LoggedModelData_SdkV2
// only implements ToObjectValue() and Type().
func (o LoggedModelData_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metrics": o.Metrics,
			"params":  o.Params,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LoggedModelData_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *LoggedModelData_SdkV2) GetMetrics(ctx context.Context) ([]Metric_SdkV2, bool) {
	if o.Metrics.IsNull() || o.Metrics.IsUnknown() {
		return nil, false
	}
	var v []Metric_SdkV2
	d := o.Metrics.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMetrics sets the value of the Metrics field in LoggedModelData_SdkV2.
func (o *LoggedModelData_SdkV2) SetMetrics(ctx context.Context, v []Metric_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["metrics"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Metrics = types.ListValueMust(t, vs)
}

// GetParams returns the value of the Params field in LoggedModelData_SdkV2 as
// a slice of LoggedModelParameter_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *LoggedModelData_SdkV2) GetParams(ctx context.Context) ([]LoggedModelParameter_SdkV2, bool) {
	if o.Params.IsNull() || o.Params.IsUnknown() {
		return nil, false
	}
	var v []LoggedModelParameter_SdkV2
	d := o.Params.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParams sets the value of the Params field in LoggedModelData_SdkV2.
func (o *LoggedModelData_SdkV2) SetParams(ctx context.Context, v []LoggedModelParameter_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Params = types.ListValueMust(t, vs)
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

func (toState *LoggedModelInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan LoggedModelInfo_SdkV2) {
}

func (toState *LoggedModelInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState LoggedModelInfo_SdkV2) {
}

func (c LoggedModelInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a LoggedModelInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(LoggedModelTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LoggedModelInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o LoggedModelInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"artifact_uri":              o.ArtifactUri,
			"creation_timestamp_ms":     o.CreationTimestampMs,
			"creator_id":                o.CreatorId,
			"experiment_id":             o.ExperimentId,
			"last_updated_timestamp_ms": o.LastUpdatedTimestampMs,
			"model_id":                  o.ModelId,
			"model_type":                o.ModelType,
			"name":                      o.Name,
			"source_run_id":             o.SourceRunId,
			"status":                    o.Status,
			"status_message":            o.StatusMessage,
			"tags":                      o.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LoggedModelInfo_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *LoggedModelInfo_SdkV2) GetTags(ctx context.Context) ([]LoggedModelTag_SdkV2, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []LoggedModelTag_SdkV2
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in LoggedModelInfo_SdkV2.
func (o *LoggedModelInfo_SdkV2) SetTags(ctx context.Context, v []LoggedModelTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

// Parameter associated with a LoggedModel.
type LoggedModelParameter_SdkV2 struct {
	// The key identifying this param.
	Key types.String `tfsdk:"key"`
	// The value of this param.
	Value types.String `tfsdk:"value"`
}

func (toState *LoggedModelParameter_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan LoggedModelParameter_SdkV2) {
}

func (toState *LoggedModelParameter_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState LoggedModelParameter_SdkV2) {
}

func (c LoggedModelParameter_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a LoggedModelParameter_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LoggedModelParameter_SdkV2
// only implements ToObjectValue() and Type().
func (o LoggedModelParameter_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LoggedModelParameter_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *LoggedModelTag_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan LoggedModelTag_SdkV2) {
}

func (toState *LoggedModelTag_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState LoggedModelTag_SdkV2) {
}

func (c LoggedModelTag_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a LoggedModelTag_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LoggedModelTag_SdkV2
// only implements ToObjectValue() and Type().
func (o LoggedModelTag_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LoggedModelTag_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *Metric_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Metric_SdkV2) {
}

func (toState *Metric_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState Metric_SdkV2) {
}

func (c Metric_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Metric_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Metric_SdkV2
// only implements ToObjectValue() and Type().
func (o Metric_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dataset_digest": o.DatasetDigest,
			"dataset_name":   o.DatasetName,
			"key":            o.Key,
			"model_id":       o.ModelId,
			"run_id":         o.RunId,
			"step":           o.Step,
			"timestamp":      o.Timestamp,
			"value":          o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Metric_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *Model_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Model_SdkV2) {
}

func (toState *Model_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState Model_SdkV2) {
}

func (c Model_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Model_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"latest_versions": reflect.TypeOf(ModelVersion_SdkV2{}),
		"tags":            reflect.TypeOf(ModelTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Model_SdkV2
// only implements ToObjectValue() and Type().
func (o Model_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creation_timestamp":     o.CreationTimestamp,
			"description":            o.Description,
			"last_updated_timestamp": o.LastUpdatedTimestamp,
			"latest_versions":        o.LatestVersions,
			"name":                   o.Name,
			"tags":                   o.Tags,
			"user_id":                o.UserId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Model_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *Model_SdkV2) GetLatestVersions(ctx context.Context) ([]ModelVersion_SdkV2, bool) {
	if o.LatestVersions.IsNull() || o.LatestVersions.IsUnknown() {
		return nil, false
	}
	var v []ModelVersion_SdkV2
	d := o.LatestVersions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLatestVersions sets the value of the LatestVersions field in Model_SdkV2.
func (o *Model_SdkV2) SetLatestVersions(ctx context.Context, v []ModelVersion_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["latest_versions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.LatestVersions = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in Model_SdkV2 as
// a slice of ModelTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *Model_SdkV2) GetTags(ctx context.Context) ([]ModelTag_SdkV2, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []ModelTag_SdkV2
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in Model_SdkV2.
func (o *Model_SdkV2) SetTags(ctx context.Context, v []ModelTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
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

func (toState *ModelDatabricks_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ModelDatabricks_SdkV2) {
}

func (toState *ModelDatabricks_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ModelDatabricks_SdkV2) {
}

func (c ModelDatabricks_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ModelDatabricks_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"latest_versions": reflect.TypeOf(ModelVersion_SdkV2{}),
		"tags":            reflect.TypeOf(ModelTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelDatabricks_SdkV2
// only implements ToObjectValue() and Type().
func (o ModelDatabricks_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creation_timestamp":     o.CreationTimestamp,
			"description":            o.Description,
			"id":                     o.Id,
			"last_updated_timestamp": o.LastUpdatedTimestamp,
			"latest_versions":        o.LatestVersions,
			"name":                   o.Name,
			"permission_level":       o.PermissionLevel,
			"tags":                   o.Tags,
			"user_id":                o.UserId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ModelDatabricks_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ModelDatabricks_SdkV2) GetLatestVersions(ctx context.Context) ([]ModelVersion_SdkV2, bool) {
	if o.LatestVersions.IsNull() || o.LatestVersions.IsUnknown() {
		return nil, false
	}
	var v []ModelVersion_SdkV2
	d := o.LatestVersions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLatestVersions sets the value of the LatestVersions field in ModelDatabricks_SdkV2.
func (o *ModelDatabricks_SdkV2) SetLatestVersions(ctx context.Context, v []ModelVersion_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["latest_versions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.LatestVersions = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in ModelDatabricks_SdkV2 as
// a slice of ModelTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ModelDatabricks_SdkV2) GetTags(ctx context.Context) ([]ModelTag_SdkV2, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []ModelTag_SdkV2
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in ModelDatabricks_SdkV2.
func (o *ModelDatabricks_SdkV2) SetTags(ctx context.Context, v []ModelTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

// Represents a LoggedModel or Registered Model Version input to a Run.
type ModelInput_SdkV2 struct {
	// The unique identifier of the model.
	ModelId types.String `tfsdk:"model_id"`
}

func (toState *ModelInput_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ModelInput_SdkV2) {
}

func (toState *ModelInput_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ModelInput_SdkV2) {
}

func (c ModelInput_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ModelInput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelInput_SdkV2
// only implements ToObjectValue() and Type().
func (o ModelInput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_id": o.ModelId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ModelInput_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *ModelOutput_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ModelOutput_SdkV2) {
}

func (toState *ModelOutput_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ModelOutput_SdkV2) {
}

func (c ModelOutput_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ModelOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelOutput_SdkV2
// only implements ToObjectValue() and Type().
func (o ModelOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_id": o.ModelId,
			"step":     o.Step,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ModelOutput_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *ModelTag_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ModelTag_SdkV2) {
}

func (toState *ModelTag_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ModelTag_SdkV2) {
}

func (c ModelTag_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ModelTag_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelTag_SdkV2
// only implements ToObjectValue() and Type().
func (o ModelTag_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ModelTag_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *ModelVersion_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ModelVersion_SdkV2) {
}

func (toState *ModelVersion_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ModelVersion_SdkV2) {
}

func (c ModelVersion_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ModelVersion_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(ModelVersionTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelVersion_SdkV2
// only implements ToObjectValue() and Type().
func (o ModelVersion_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creation_timestamp":     o.CreationTimestamp,
			"current_stage":          o.CurrentStage,
			"description":            o.Description,
			"last_updated_timestamp": o.LastUpdatedTimestamp,
			"name":                   o.Name,
			"run_id":                 o.RunId,
			"run_link":               o.RunLink,
			"source":                 o.Source,
			"status":                 o.Status,
			"status_message":         o.StatusMessage,
			"tags":                   o.Tags,
			"user_id":                o.UserId,
			"version":                o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ModelVersion_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ModelVersion_SdkV2) GetTags(ctx context.Context) ([]ModelVersionTag_SdkV2, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []ModelVersionTag_SdkV2
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in ModelVersion_SdkV2.
func (o *ModelVersion_SdkV2) SetTags(ctx context.Context, v []ModelVersionTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
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

func (toState *ModelVersionDatabricks_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ModelVersionDatabricks_SdkV2) {
	if !fromPlan.FeatureList.IsNull() && !fromPlan.FeatureList.IsUnknown() {
		if toStateFeatureList, ok := toState.GetFeatureList(ctx); ok {
			if fromPlanFeatureList, ok := fromPlan.GetFeatureList(ctx); ok {
				toStateFeatureList.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanFeatureList)
				toState.SetFeatureList(ctx, toStateFeatureList)
			}
		}
	}
}

func (toState *ModelVersionDatabricks_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ModelVersionDatabricks_SdkV2) {
	if !fromState.FeatureList.IsNull() && !fromState.FeatureList.IsUnknown() {
		if toStateFeatureList, ok := toState.GetFeatureList(ctx); ok {
			if fromStateFeatureList, ok := fromState.GetFeatureList(ctx); ok {
				toStateFeatureList.SyncFieldsDuringRead(ctx, fromStateFeatureList)
				toState.SetFeatureList(ctx, toStateFeatureList)
			}
		}
	}
}

func (c ModelVersionDatabricks_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ModelVersionDatabricks_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"feature_list":  reflect.TypeOf(FeatureList_SdkV2{}),
		"open_requests": reflect.TypeOf(Activity_SdkV2{}),
		"tags":          reflect.TypeOf(ModelVersionTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelVersionDatabricks_SdkV2
// only implements ToObjectValue() and Type().
func (o ModelVersionDatabricks_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creation_timestamp":        o.CreationTimestamp,
			"current_stage":             o.CurrentStage,
			"description":               o.Description,
			"email_subscription_status": o.EmailSubscriptionStatus,
			"feature_list":              o.FeatureList,
			"last_updated_timestamp":    o.LastUpdatedTimestamp,
			"name":                      o.Name,
			"open_requests":             o.OpenRequests,
			"permission_level":          o.PermissionLevel,
			"run_id":                    o.RunId,
			"run_link":                  o.RunLink,
			"source":                    o.Source,
			"status":                    o.Status,
			"status_message":            o.StatusMessage,
			"tags":                      o.Tags,
			"user_id":                   o.UserId,
			"version":                   o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ModelVersionDatabricks_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ModelVersionDatabricks_SdkV2) GetFeatureList(ctx context.Context) (FeatureList_SdkV2, bool) {
	var e FeatureList_SdkV2
	if o.FeatureList.IsNull() || o.FeatureList.IsUnknown() {
		return e, false
	}
	var v []FeatureList_SdkV2
	d := o.FeatureList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFeatureList sets the value of the FeatureList field in ModelVersionDatabricks_SdkV2.
func (o *ModelVersionDatabricks_SdkV2) SetFeatureList(ctx context.Context, v FeatureList_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["feature_list"]
	o.FeatureList = types.ListValueMust(t, vs)
}

// GetOpenRequests returns the value of the OpenRequests field in ModelVersionDatabricks_SdkV2 as
// a slice of Activity_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ModelVersionDatabricks_SdkV2) GetOpenRequests(ctx context.Context) ([]Activity_SdkV2, bool) {
	if o.OpenRequests.IsNull() || o.OpenRequests.IsUnknown() {
		return nil, false
	}
	var v []Activity_SdkV2
	d := o.OpenRequests.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOpenRequests sets the value of the OpenRequests field in ModelVersionDatabricks_SdkV2.
func (o *ModelVersionDatabricks_SdkV2) SetOpenRequests(ctx context.Context, v []Activity_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["open_requests"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OpenRequests = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in ModelVersionDatabricks_SdkV2 as
// a slice of ModelVersionTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ModelVersionDatabricks_SdkV2) GetTags(ctx context.Context) ([]ModelVersionTag_SdkV2, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []ModelVersionTag_SdkV2
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in ModelVersionDatabricks_SdkV2.
func (o *ModelVersionDatabricks_SdkV2) SetTags(ctx context.Context, v []ModelVersionTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type ModelVersionTag_SdkV2 struct {
	// The tag key.
	Key types.String `tfsdk:"key"`
	// The tag value.
	Value types.String `tfsdk:"value"`
}

func (toState *ModelVersionTag_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ModelVersionTag_SdkV2) {
}

func (toState *ModelVersionTag_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ModelVersionTag_SdkV2) {
}

func (c ModelVersionTag_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ModelVersionTag_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelVersionTag_SdkV2
// only implements ToObjectValue() and Type().
func (o ModelVersionTag_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ModelVersionTag_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *OnlineStore_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan OnlineStore_SdkV2) {
}

func (toState *OnlineStore_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState OnlineStore_SdkV2) {
}

func (c OnlineStore_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a OnlineStore_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, OnlineStore_SdkV2
// only implements ToObjectValue() and Type().
func (o OnlineStore_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"capacity":           o.Capacity,
			"creation_time":      o.CreationTime,
			"creator":            o.Creator,
			"name":               o.Name,
			"read_replica_count": o.ReadReplicaCount,
			"state":              o.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (o OnlineStore_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *Param_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Param_SdkV2) {
}

func (toState *Param_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState Param_SdkV2) {
}

func (c Param_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Param_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Param_SdkV2
// only implements ToObjectValue() and Type().
func (o Param_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Param_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *PublishSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PublishSpec_SdkV2) {
}

func (toState *PublishSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState PublishSpec_SdkV2) {
}

func (c PublishSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PublishSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PublishSpec_SdkV2
// only implements ToObjectValue() and Type().
func (o PublishSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"online_store":      o.OnlineStore,
			"online_table_name": o.OnlineTableName,
			"publish_mode":      o.PublishMode,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PublishSpec_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *PublishTableRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PublishTableRequest_SdkV2) {
	if !fromPlan.PublishSpec.IsNull() && !fromPlan.PublishSpec.IsUnknown() {
		if toStatePublishSpec, ok := toState.GetPublishSpec(ctx); ok {
			if fromPlanPublishSpec, ok := fromPlan.GetPublishSpec(ctx); ok {
				toStatePublishSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanPublishSpec)
				toState.SetPublishSpec(ctx, toStatePublishSpec)
			}
		}
	}
}

func (toState *PublishTableRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState PublishTableRequest_SdkV2) {
	if !fromState.PublishSpec.IsNull() && !fromState.PublishSpec.IsUnknown() {
		if toStatePublishSpec, ok := toState.GetPublishSpec(ctx); ok {
			if fromStatePublishSpec, ok := fromState.GetPublishSpec(ctx); ok {
				toStatePublishSpec.SyncFieldsDuringRead(ctx, fromStatePublishSpec)
				toState.SetPublishSpec(ctx, toStatePublishSpec)
			}
		}
	}
}

func (c PublishTableRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PublishTableRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"publish_spec": reflect.TypeOf(PublishSpec_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PublishTableRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o PublishTableRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"publish_spec":      o.PublishSpec,
			"source_table_name": o.SourceTableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PublishTableRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *PublishTableRequest_SdkV2) GetPublishSpec(ctx context.Context) (PublishSpec_SdkV2, bool) {
	var e PublishSpec_SdkV2
	if o.PublishSpec.IsNull() || o.PublishSpec.IsUnknown() {
		return e, false
	}
	var v []PublishSpec_SdkV2
	d := o.PublishSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPublishSpec sets the value of the PublishSpec field in PublishTableRequest_SdkV2.
func (o *PublishTableRequest_SdkV2) SetPublishSpec(ctx context.Context, v PublishSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["publish_spec"]
	o.PublishSpec = types.ListValueMust(t, vs)
}

type PublishTableResponse_SdkV2 struct {
	// The full three-part (catalog, schema, table) name of the online table.
	OnlineTableName types.String `tfsdk:"online_table_name"`
	// The ID of the pipeline that syncs the online table with the source table.
	PipelineId types.String `tfsdk:"pipeline_id"`
}

func (toState *PublishTableResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan PublishTableResponse_SdkV2) {
}

func (toState *PublishTableResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState PublishTableResponse_SdkV2) {
}

func (c PublishTableResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PublishTableResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PublishTableResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o PublishTableResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"online_table_name": o.OnlineTableName,
			"pipeline_id":       o.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PublishTableResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *RegisteredModelAccessControlRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan RegisteredModelAccessControlRequest_SdkV2) {
}

func (toState *RegisteredModelAccessControlRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState RegisteredModelAccessControlRequest_SdkV2) {
}

func (c RegisteredModelAccessControlRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RegisteredModelAccessControlRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegisteredModelAccessControlRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o RegisteredModelAccessControlRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_name":             o.GroupName,
			"permission_level":       o.PermissionLevel,
			"service_principal_name": o.ServicePrincipalName,
			"user_name":              o.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RegisteredModelAccessControlRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *RegisteredModelAccessControlResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan RegisteredModelAccessControlResponse_SdkV2) {
}

func (toState *RegisteredModelAccessControlResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState RegisteredModelAccessControlResponse_SdkV2) {
}

func (c RegisteredModelAccessControlResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RegisteredModelAccessControlResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(RegisteredModelPermission_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegisteredModelAccessControlResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o RegisteredModelAccessControlResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"all_permissions":        o.AllPermissions,
			"display_name":           o.DisplayName,
			"group_name":             o.GroupName,
			"service_principal_name": o.ServicePrincipalName,
			"user_name":              o.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RegisteredModelAccessControlResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *RegisteredModelAccessControlResponse_SdkV2) GetAllPermissions(ctx context.Context) ([]RegisteredModelPermission_SdkV2, bool) {
	if o.AllPermissions.IsNull() || o.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []RegisteredModelPermission_SdkV2
	d := o.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in RegisteredModelAccessControlResponse_SdkV2.
func (o *RegisteredModelAccessControlResponse_SdkV2) SetAllPermissions(ctx context.Context, v []RegisteredModelPermission_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllPermissions = types.ListValueMust(t, vs)
}

type RegisteredModelPermission_SdkV2 struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (toState *RegisteredModelPermission_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan RegisteredModelPermission_SdkV2) {
}

func (toState *RegisteredModelPermission_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState RegisteredModelPermission_SdkV2) {
}

func (c RegisteredModelPermission_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RegisteredModelPermission_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegisteredModelPermission_SdkV2
// only implements ToObjectValue() and Type().
func (o RegisteredModelPermission_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             o.Inherited,
			"inherited_from_object": o.InheritedFromObject,
			"permission_level":      o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RegisteredModelPermission_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *RegisteredModelPermission_SdkV2) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
	if o.InheritedFromObject.IsNull() || o.InheritedFromObject.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.InheritedFromObject.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInheritedFromObject sets the value of the InheritedFromObject field in RegisteredModelPermission_SdkV2.
func (o *RegisteredModelPermission_SdkV2) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InheritedFromObject = types.ListValueMust(t, vs)
}

type RegisteredModelPermissions_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (toState *RegisteredModelPermissions_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan RegisteredModelPermissions_SdkV2) {
}

func (toState *RegisteredModelPermissions_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState RegisteredModelPermissions_SdkV2) {
}

func (c RegisteredModelPermissions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RegisteredModelPermissions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(RegisteredModelAccessControlResponse_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegisteredModelPermissions_SdkV2
// only implements ToObjectValue() and Type().
func (o RegisteredModelPermissions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RegisteredModelPermissions_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *RegisteredModelPermissions_SdkV2) GetAccessControlList(ctx context.Context) ([]RegisteredModelAccessControlResponse_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []RegisteredModelAccessControlResponse_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in RegisteredModelPermissions_SdkV2.
func (o *RegisteredModelPermissions_SdkV2) SetAccessControlList(ctx context.Context, v []RegisteredModelAccessControlResponse_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type RegisteredModelPermissionsDescription_SdkV2 struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (toState *RegisteredModelPermissionsDescription_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan RegisteredModelPermissionsDescription_SdkV2) {
}

func (toState *RegisteredModelPermissionsDescription_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState RegisteredModelPermissionsDescription_SdkV2) {
}

func (c RegisteredModelPermissionsDescription_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RegisteredModelPermissionsDescription_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegisteredModelPermissionsDescription_SdkV2
// only implements ToObjectValue() and Type().
func (o RegisteredModelPermissionsDescription_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      o.Description,
			"permission_level": o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RegisteredModelPermissionsDescription_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *RegisteredModelPermissionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan RegisteredModelPermissionsRequest_SdkV2) {
}

func (toState *RegisteredModelPermissionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState RegisteredModelPermissionsRequest_SdkV2) {
}

func (c RegisteredModelPermissionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RegisteredModelPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(RegisteredModelAccessControlRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegisteredModelPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o RegisteredModelPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"registered_model_id": o.RegisteredModelId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RegisteredModelPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *RegisteredModelPermissionsRequest_SdkV2) GetAccessControlList(ctx context.Context) ([]RegisteredModelAccessControlRequest_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []RegisteredModelAccessControlRequest_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in RegisteredModelPermissionsRequest_SdkV2.
func (o *RegisteredModelPermissionsRequest_SdkV2) SetAccessControlList(ctx context.Context, v []RegisteredModelAccessControlRequest_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
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

func (toState *RegistryWebhook_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan RegistryWebhook_SdkV2) {
	if !fromPlan.HttpUrlSpec.IsNull() && !fromPlan.HttpUrlSpec.IsUnknown() {
		if toStateHttpUrlSpec, ok := toState.GetHttpUrlSpec(ctx); ok {
			if fromPlanHttpUrlSpec, ok := fromPlan.GetHttpUrlSpec(ctx); ok {
				toStateHttpUrlSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanHttpUrlSpec)
				toState.SetHttpUrlSpec(ctx, toStateHttpUrlSpec)
			}
		}
	}
	if !fromPlan.JobSpec.IsNull() && !fromPlan.JobSpec.IsUnknown() {
		if toStateJobSpec, ok := toState.GetJobSpec(ctx); ok {
			if fromPlanJobSpec, ok := fromPlan.GetJobSpec(ctx); ok {
				toStateJobSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanJobSpec)
				toState.SetJobSpec(ctx, toStateJobSpec)
			}
		}
	}
}

func (toState *RegistryWebhook_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState RegistryWebhook_SdkV2) {
	if !fromState.HttpUrlSpec.IsNull() && !fromState.HttpUrlSpec.IsUnknown() {
		if toStateHttpUrlSpec, ok := toState.GetHttpUrlSpec(ctx); ok {
			if fromStateHttpUrlSpec, ok := fromState.GetHttpUrlSpec(ctx); ok {
				toStateHttpUrlSpec.SyncFieldsDuringRead(ctx, fromStateHttpUrlSpec)
				toState.SetHttpUrlSpec(ctx, toStateHttpUrlSpec)
			}
		}
	}
	if !fromState.JobSpec.IsNull() && !fromState.JobSpec.IsUnknown() {
		if toStateJobSpec, ok := toState.GetJobSpec(ctx); ok {
			if fromStateJobSpec, ok := fromState.GetJobSpec(ctx); ok {
				toStateJobSpec.SyncFieldsDuringRead(ctx, fromStateJobSpec)
				toState.SetJobSpec(ctx, toStateJobSpec)
			}
		}
	}
}

func (c RegistryWebhook_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RegistryWebhook_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"events":        reflect.TypeOf(types.String{}),
		"http_url_spec": reflect.TypeOf(HttpUrlSpecWithoutSecret_SdkV2{}),
		"job_spec":      reflect.TypeOf(JobSpecWithoutSecret_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegistryWebhook_SdkV2
// only implements ToObjectValue() and Type().
func (o RegistryWebhook_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creation_timestamp":     o.CreationTimestamp,
			"description":            o.Description,
			"events":                 o.Events,
			"http_url_spec":          o.HttpUrlSpec,
			"id":                     o.Id,
			"job_spec":               o.JobSpec,
			"last_updated_timestamp": o.LastUpdatedTimestamp,
			"model_name":             o.ModelName,
			"status":                 o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RegistryWebhook_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *RegistryWebhook_SdkV2) GetEvents(ctx context.Context) ([]types.String, bool) {
	if o.Events.IsNull() || o.Events.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Events.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEvents sets the value of the Events field in RegistryWebhook_SdkV2.
func (o *RegistryWebhook_SdkV2) SetEvents(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["events"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Events = types.ListValueMust(t, vs)
}

// GetHttpUrlSpec returns the value of the HttpUrlSpec field in RegistryWebhook_SdkV2 as
// a HttpUrlSpecWithoutSecret_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RegistryWebhook_SdkV2) GetHttpUrlSpec(ctx context.Context) (HttpUrlSpecWithoutSecret_SdkV2, bool) {
	var e HttpUrlSpecWithoutSecret_SdkV2
	if o.HttpUrlSpec.IsNull() || o.HttpUrlSpec.IsUnknown() {
		return e, false
	}
	var v []HttpUrlSpecWithoutSecret_SdkV2
	d := o.HttpUrlSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetHttpUrlSpec sets the value of the HttpUrlSpec field in RegistryWebhook_SdkV2.
func (o *RegistryWebhook_SdkV2) SetHttpUrlSpec(ctx context.Context, v HttpUrlSpecWithoutSecret_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["http_url_spec"]
	o.HttpUrlSpec = types.ListValueMust(t, vs)
}

// GetJobSpec returns the value of the JobSpec field in RegistryWebhook_SdkV2 as
// a JobSpecWithoutSecret_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *RegistryWebhook_SdkV2) GetJobSpec(ctx context.Context) (JobSpecWithoutSecret_SdkV2, bool) {
	var e JobSpecWithoutSecret_SdkV2
	if o.JobSpec.IsNull() || o.JobSpec.IsUnknown() {
		return e, false
	}
	var v []JobSpecWithoutSecret_SdkV2
	d := o.JobSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetJobSpec sets the value of the JobSpec field in RegistryWebhook_SdkV2.
func (o *RegistryWebhook_SdkV2) SetJobSpec(ctx context.Context, v JobSpecWithoutSecret_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["job_spec"]
	o.JobSpec = types.ListValueMust(t, vs)
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

func (toState *RejectTransitionRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan RejectTransitionRequest_SdkV2) {
}

func (toState *RejectTransitionRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState RejectTransitionRequest_SdkV2) {
}

func (c RejectTransitionRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RejectTransitionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RejectTransitionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o RejectTransitionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment": o.Comment,
			"name":    o.Name,
			"stage":   o.Stage,
			"version": o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RejectTransitionRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *RejectTransitionRequestResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan RejectTransitionRequestResponse_SdkV2) {
	if !fromPlan.Activity.IsNull() && !fromPlan.Activity.IsUnknown() {
		if toStateActivity, ok := toState.GetActivity(ctx); ok {
			if fromPlanActivity, ok := fromPlan.GetActivity(ctx); ok {
				toStateActivity.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanActivity)
				toState.SetActivity(ctx, toStateActivity)
			}
		}
	}
}

func (toState *RejectTransitionRequestResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState RejectTransitionRequestResponse_SdkV2) {
	if !fromState.Activity.IsNull() && !fromState.Activity.IsUnknown() {
		if toStateActivity, ok := toState.GetActivity(ctx); ok {
			if fromStateActivity, ok := fromState.GetActivity(ctx); ok {
				toStateActivity.SyncFieldsDuringRead(ctx, fromStateActivity)
				toState.SetActivity(ctx, toStateActivity)
			}
		}
	}
}

func (c RejectTransitionRequestResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RejectTransitionRequestResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"activity": reflect.TypeOf(Activity_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RejectTransitionRequestResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o RejectTransitionRequestResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"activity": o.Activity,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RejectTransitionRequestResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *RejectTransitionRequestResponse_SdkV2) GetActivity(ctx context.Context) (Activity_SdkV2, bool) {
	var e Activity_SdkV2
	if o.Activity.IsNull() || o.Activity.IsUnknown() {
		return e, false
	}
	var v []Activity_SdkV2
	d := o.Activity.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetActivity sets the value of the Activity field in RejectTransitionRequestResponse_SdkV2.
func (o *RejectTransitionRequestResponse_SdkV2) SetActivity(ctx context.Context, v Activity_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["activity"]
	o.Activity = types.ListValueMust(t, vs)
}

type RenameModelRequest_SdkV2 struct {
	// Registered model unique name identifier.
	Name types.String `tfsdk:"name"`
	// If provided, updates the name for this `registered_model`.
	NewName types.String `tfsdk:"new_name"`
}

func (toState *RenameModelRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan RenameModelRequest_SdkV2) {
}

func (toState *RenameModelRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState RenameModelRequest_SdkV2) {
}

func (c RenameModelRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RenameModelRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RenameModelRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o RenameModelRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":     o.Name,
			"new_name": o.NewName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RenameModelRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *RenameModelResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan RenameModelResponse_SdkV2) {
	if !fromPlan.RegisteredModel.IsNull() && !fromPlan.RegisteredModel.IsUnknown() {
		if toStateRegisteredModel, ok := toState.GetRegisteredModel(ctx); ok {
			if fromPlanRegisteredModel, ok := fromPlan.GetRegisteredModel(ctx); ok {
				toStateRegisteredModel.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanRegisteredModel)
				toState.SetRegisteredModel(ctx, toStateRegisteredModel)
			}
		}
	}
}

func (toState *RenameModelResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState RenameModelResponse_SdkV2) {
	if !fromState.RegisteredModel.IsNull() && !fromState.RegisteredModel.IsUnknown() {
		if toStateRegisteredModel, ok := toState.GetRegisteredModel(ctx); ok {
			if fromStateRegisteredModel, ok := fromState.GetRegisteredModel(ctx); ok {
				toStateRegisteredModel.SyncFieldsDuringRead(ctx, fromStateRegisteredModel)
				toState.SetRegisteredModel(ctx, toStateRegisteredModel)
			}
		}
	}
}

func (c RenameModelResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RenameModelResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"registered_model": reflect.TypeOf(Model_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RenameModelResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o RenameModelResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"registered_model": o.RegisteredModel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RenameModelResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *RenameModelResponse_SdkV2) GetRegisteredModel(ctx context.Context) (Model_SdkV2, bool) {
	var e Model_SdkV2
	if o.RegisteredModel.IsNull() || o.RegisteredModel.IsUnknown() {
		return e, false
	}
	var v []Model_SdkV2
	d := o.RegisteredModel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRegisteredModel sets the value of the RegisteredModel field in RenameModelResponse_SdkV2.
func (o *RenameModelResponse_SdkV2) SetRegisteredModel(ctx context.Context, v Model_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["registered_model"]
	o.RegisteredModel = types.ListValueMust(t, vs)
}

type RestoreExperiment_SdkV2 struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"experiment_id"`
}

func (toState *RestoreExperiment_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan RestoreExperiment_SdkV2) {
}

func (toState *RestoreExperiment_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState RestoreExperiment_SdkV2) {
}

func (c RestoreExperiment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RestoreExperiment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestoreExperiment_SdkV2
// only implements ToObjectValue() and Type().
func (o RestoreExperiment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": o.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RestoreExperiment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
		},
	}
}

type RestoreExperimentResponse_SdkV2 struct {
}

func (toState *RestoreExperimentResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan RestoreExperimentResponse_SdkV2) {
}

func (toState *RestoreExperimentResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState RestoreExperimentResponse_SdkV2) {
}

func (c RestoreExperimentResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestoreExperimentResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RestoreExperimentResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestoreExperimentResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o RestoreExperimentResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o RestoreExperimentResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type RestoreRun_SdkV2 struct {
	// ID of the run to restore.
	RunId types.String `tfsdk:"run_id"`
}

func (toState *RestoreRun_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan RestoreRun_SdkV2) {
}

func (toState *RestoreRun_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState RestoreRun_SdkV2) {
}

func (c RestoreRun_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RestoreRun_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestoreRun_SdkV2
// only implements ToObjectValue() and Type().
func (o RestoreRun_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run_id": o.RunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RestoreRun_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run_id": types.StringType,
		},
	}
}

type RestoreRunResponse_SdkV2 struct {
}

func (toState *RestoreRunResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan RestoreRunResponse_SdkV2) {
}

func (toState *RestoreRunResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState RestoreRunResponse_SdkV2) {
}

func (c RestoreRunResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestoreRunResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RestoreRunResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestoreRunResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o RestoreRunResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o RestoreRunResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *RestoreRuns_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan RestoreRuns_SdkV2) {
}

func (toState *RestoreRuns_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState RestoreRuns_SdkV2) {
}

func (c RestoreRuns_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RestoreRuns_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestoreRuns_SdkV2
// only implements ToObjectValue() and Type().
func (o RestoreRuns_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id":        o.ExperimentId,
			"max_runs":             o.MaxRuns,
			"min_timestamp_millis": o.MinTimestampMillis,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RestoreRuns_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *RestoreRunsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan RestoreRunsResponse_SdkV2) {
}

func (toState *RestoreRunsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState RestoreRunsResponse_SdkV2) {
}

func (c RestoreRunsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RestoreRunsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestoreRunsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o RestoreRunsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"runs_restored": o.RunsRestored,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RestoreRunsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *Run_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan Run_SdkV2) {
	if !fromPlan.Data.IsNull() && !fromPlan.Data.IsUnknown() {
		if toStateData, ok := toState.GetData(ctx); ok {
			if fromPlanData, ok := fromPlan.GetData(ctx); ok {
				toStateData.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanData)
				toState.SetData(ctx, toStateData)
			}
		}
	}
	if !fromPlan.Info.IsNull() && !fromPlan.Info.IsUnknown() {
		if toStateInfo, ok := toState.GetInfo(ctx); ok {
			if fromPlanInfo, ok := fromPlan.GetInfo(ctx); ok {
				toStateInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanInfo)
				toState.SetInfo(ctx, toStateInfo)
			}
		}
	}
	if !fromPlan.Inputs.IsNull() && !fromPlan.Inputs.IsUnknown() {
		if toStateInputs, ok := toState.GetInputs(ctx); ok {
			if fromPlanInputs, ok := fromPlan.GetInputs(ctx); ok {
				toStateInputs.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanInputs)
				toState.SetInputs(ctx, toStateInputs)
			}
		}
	}
}

func (toState *Run_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState Run_SdkV2) {
	if !fromState.Data.IsNull() && !fromState.Data.IsUnknown() {
		if toStateData, ok := toState.GetData(ctx); ok {
			if fromStateData, ok := fromState.GetData(ctx); ok {
				toStateData.SyncFieldsDuringRead(ctx, fromStateData)
				toState.SetData(ctx, toStateData)
			}
		}
	}
	if !fromState.Info.IsNull() && !fromState.Info.IsUnknown() {
		if toStateInfo, ok := toState.GetInfo(ctx); ok {
			if fromStateInfo, ok := fromState.GetInfo(ctx); ok {
				toStateInfo.SyncFieldsDuringRead(ctx, fromStateInfo)
				toState.SetInfo(ctx, toStateInfo)
			}
		}
	}
	if !fromState.Inputs.IsNull() && !fromState.Inputs.IsUnknown() {
		if toStateInputs, ok := toState.GetInputs(ctx); ok {
			if fromStateInputs, ok := fromState.GetInputs(ctx); ok {
				toStateInputs.SyncFieldsDuringRead(ctx, fromStateInputs)
				toState.SetInputs(ctx, toStateInputs)
			}
		}
	}
}

func (c Run_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Run_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data":   reflect.TypeOf(RunData_SdkV2{}),
		"info":   reflect.TypeOf(RunInfo_SdkV2{}),
		"inputs": reflect.TypeOf(RunInputs_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Run_SdkV2
// only implements ToObjectValue() and Type().
func (o Run_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data":   o.Data,
			"info":   o.Info,
			"inputs": o.Inputs,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Run_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *Run_SdkV2) GetData(ctx context.Context) (RunData_SdkV2, bool) {
	var e RunData_SdkV2
	if o.Data.IsNull() || o.Data.IsUnknown() {
		return e, false
	}
	var v []RunData_SdkV2
	d := o.Data.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetData sets the value of the Data field in Run_SdkV2.
func (o *Run_SdkV2) SetData(ctx context.Context, v RunData_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["data"]
	o.Data = types.ListValueMust(t, vs)
}

// GetInfo returns the value of the Info field in Run_SdkV2 as
// a RunInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Run_SdkV2) GetInfo(ctx context.Context) (RunInfo_SdkV2, bool) {
	var e RunInfo_SdkV2
	if o.Info.IsNull() || o.Info.IsUnknown() {
		return e, false
	}
	var v []RunInfo_SdkV2
	d := o.Info.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInfo sets the value of the Info field in Run_SdkV2.
func (o *Run_SdkV2) SetInfo(ctx context.Context, v RunInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["info"]
	o.Info = types.ListValueMust(t, vs)
}

// GetInputs returns the value of the Inputs field in Run_SdkV2 as
// a RunInputs_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Run_SdkV2) GetInputs(ctx context.Context) (RunInputs_SdkV2, bool) {
	var e RunInputs_SdkV2
	if o.Inputs.IsNull() || o.Inputs.IsUnknown() {
		return e, false
	}
	var v []RunInputs_SdkV2
	d := o.Inputs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInputs sets the value of the Inputs field in Run_SdkV2.
func (o *Run_SdkV2) SetInputs(ctx context.Context, v RunInputs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inputs"]
	o.Inputs = types.ListValueMust(t, vs)
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

func (toState *RunData_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan RunData_SdkV2) {
}

func (toState *RunData_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState RunData_SdkV2) {
}

func (c RunData_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RunData_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metrics": reflect.TypeOf(Metric_SdkV2{}),
		"params":  reflect.TypeOf(Param_SdkV2{}),
		"tags":    reflect.TypeOf(RunTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RunData_SdkV2
// only implements ToObjectValue() and Type().
func (o RunData_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metrics": o.Metrics,
			"params":  o.Params,
			"tags":    o.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RunData_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *RunData_SdkV2) GetMetrics(ctx context.Context) ([]Metric_SdkV2, bool) {
	if o.Metrics.IsNull() || o.Metrics.IsUnknown() {
		return nil, false
	}
	var v []Metric_SdkV2
	d := o.Metrics.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMetrics sets the value of the Metrics field in RunData_SdkV2.
func (o *RunData_SdkV2) SetMetrics(ctx context.Context, v []Metric_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["metrics"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Metrics = types.ListValueMust(t, vs)
}

// GetParams returns the value of the Params field in RunData_SdkV2 as
// a slice of Param_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *RunData_SdkV2) GetParams(ctx context.Context) ([]Param_SdkV2, bool) {
	if o.Params.IsNull() || o.Params.IsUnknown() {
		return nil, false
	}
	var v []Param_SdkV2
	d := o.Params.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParams sets the value of the Params field in RunData_SdkV2.
func (o *RunData_SdkV2) SetParams(ctx context.Context, v []Param_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Params = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in RunData_SdkV2 as
// a slice of RunTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *RunData_SdkV2) GetTags(ctx context.Context) ([]RunTag_SdkV2, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []RunTag_SdkV2
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in RunData_SdkV2.
func (o *RunData_SdkV2) SetTags(ctx context.Context, v []RunTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
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

func (toState *RunInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan RunInfo_SdkV2) {
}

func (toState *RunInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState RunInfo_SdkV2) {
}

func (c RunInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RunInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RunInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o RunInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"artifact_uri":    o.ArtifactUri,
			"end_time":        o.EndTime,
			"experiment_id":   o.ExperimentId,
			"lifecycle_stage": o.LifecycleStage,
			"run_id":          o.RunId,
			"run_name":        o.RunName,
			"run_uuid":        o.RunUuid,
			"start_time":      o.StartTime,
			"status":          o.Status,
			"user_id":         o.UserId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RunInfo_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *RunInputs_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan RunInputs_SdkV2) {
}

func (toState *RunInputs_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState RunInputs_SdkV2) {
}

func (c RunInputs_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RunInputs_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dataset_inputs": reflect.TypeOf(DatasetInput_SdkV2{}),
		"model_inputs":   reflect.TypeOf(ModelInput_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RunInputs_SdkV2
// only implements ToObjectValue() and Type().
func (o RunInputs_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dataset_inputs": o.DatasetInputs,
			"model_inputs":   o.ModelInputs,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RunInputs_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *RunInputs_SdkV2) GetDatasetInputs(ctx context.Context) ([]DatasetInput_SdkV2, bool) {
	if o.DatasetInputs.IsNull() || o.DatasetInputs.IsUnknown() {
		return nil, false
	}
	var v []DatasetInput_SdkV2
	d := o.DatasetInputs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatasetInputs sets the value of the DatasetInputs field in RunInputs_SdkV2.
func (o *RunInputs_SdkV2) SetDatasetInputs(ctx context.Context, v []DatasetInput_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dataset_inputs"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DatasetInputs = types.ListValueMust(t, vs)
}

// GetModelInputs returns the value of the ModelInputs field in RunInputs_SdkV2 as
// a slice of ModelInput_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *RunInputs_SdkV2) GetModelInputs(ctx context.Context) ([]ModelInput_SdkV2, bool) {
	if o.ModelInputs.IsNull() || o.ModelInputs.IsUnknown() {
		return nil, false
	}
	var v []ModelInput_SdkV2
	d := o.ModelInputs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModelInputs sets the value of the ModelInputs field in RunInputs_SdkV2.
func (o *RunInputs_SdkV2) SetModelInputs(ctx context.Context, v []ModelInput_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["model_inputs"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ModelInputs = types.ListValueMust(t, vs)
}

// Tag for a run.
type RunTag_SdkV2 struct {
	// The tag key.
	Key types.String `tfsdk:"key"`
	// The tag value.
	Value types.String `tfsdk:"value"`
}

func (toState *RunTag_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan RunTag_SdkV2) {
}

func (toState *RunTag_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState RunTag_SdkV2) {
}

func (c RunTag_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RunTag_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RunTag_SdkV2
// only implements ToObjectValue() and Type().
func (o RunTag_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RunTag_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *SearchExperiments_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SearchExperiments_SdkV2) {
}

func (toState *SearchExperiments_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState SearchExperiments_SdkV2) {
}

func (c SearchExperiments_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SearchExperiments_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"order_by": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchExperiments_SdkV2
// only implements ToObjectValue() and Type().
func (o SearchExperiments_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter":      o.Filter,
			"max_results": o.MaxResults,
			"order_by":    o.OrderBy,
			"page_token":  o.PageToken,
			"view_type":   o.ViewType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SearchExperiments_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *SearchExperiments_SdkV2) GetOrderBy(ctx context.Context) ([]types.String, bool) {
	if o.OrderBy.IsNull() || o.OrderBy.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.OrderBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOrderBy sets the value of the OrderBy field in SearchExperiments_SdkV2.
func (o *SearchExperiments_SdkV2) SetOrderBy(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["order_by"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OrderBy = types.ListValueMust(t, vs)
}

type SearchExperimentsResponse_SdkV2 struct {
	// Experiments that match the search criteria
	Experiments types.List `tfsdk:"experiments"`
	// Token that can be used to retrieve the next page of experiments. An empty
	// token means that no more experiments are available for retrieval.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (toState *SearchExperimentsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SearchExperimentsResponse_SdkV2) {
}

func (toState *SearchExperimentsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState SearchExperimentsResponse_SdkV2) {
}

func (c SearchExperimentsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SearchExperimentsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"experiments": reflect.TypeOf(Experiment_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchExperimentsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o SearchExperimentsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiments":     o.Experiments,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SearchExperimentsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *SearchExperimentsResponse_SdkV2) GetExperiments(ctx context.Context) ([]Experiment_SdkV2, bool) {
	if o.Experiments.IsNull() || o.Experiments.IsUnknown() {
		return nil, false
	}
	var v []Experiment_SdkV2
	d := o.Experiments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExperiments sets the value of the Experiments field in SearchExperimentsResponse_SdkV2.
func (o *SearchExperimentsResponse_SdkV2) SetExperiments(ctx context.Context, v []Experiment_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["experiments"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Experiments = types.ListValueMust(t, vs)
}

type SearchLoggedModelsDataset_SdkV2 struct {
	// The digest of the dataset.
	DatasetDigest types.String `tfsdk:"dataset_digest"`
	// The name of the dataset.
	DatasetName types.String `tfsdk:"dataset_name"`
}

func (toState *SearchLoggedModelsDataset_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SearchLoggedModelsDataset_SdkV2) {
}

func (toState *SearchLoggedModelsDataset_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState SearchLoggedModelsDataset_SdkV2) {
}

func (c SearchLoggedModelsDataset_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SearchLoggedModelsDataset_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchLoggedModelsDataset_SdkV2
// only implements ToObjectValue() and Type().
func (o SearchLoggedModelsDataset_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dataset_digest": o.DatasetDigest,
			"dataset_name":   o.DatasetName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SearchLoggedModelsDataset_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *SearchLoggedModelsOrderBy_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SearchLoggedModelsOrderBy_SdkV2) {
}

func (toState *SearchLoggedModelsOrderBy_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState SearchLoggedModelsOrderBy_SdkV2) {
}

func (c SearchLoggedModelsOrderBy_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SearchLoggedModelsOrderBy_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchLoggedModelsOrderBy_SdkV2
// only implements ToObjectValue() and Type().
func (o SearchLoggedModelsOrderBy_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ascending":      o.Ascending,
			"dataset_digest": o.DatasetDigest,
			"dataset_name":   o.DatasetName,
			"field_name":     o.FieldName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SearchLoggedModelsOrderBy_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *SearchLoggedModelsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SearchLoggedModelsRequest_SdkV2) {
}

func (toState *SearchLoggedModelsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState SearchLoggedModelsRequest_SdkV2) {
}

func (c SearchLoggedModelsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SearchLoggedModelsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"datasets":       reflect.TypeOf(SearchLoggedModelsDataset_SdkV2{}),
		"experiment_ids": reflect.TypeOf(types.String{}),
		"order_by":       reflect.TypeOf(SearchLoggedModelsOrderBy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchLoggedModelsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o SearchLoggedModelsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"datasets":       o.Datasets,
			"experiment_ids": o.ExperimentIds,
			"filter":         o.Filter,
			"max_results":    o.MaxResults,
			"order_by":       o.OrderBy,
			"page_token":     o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SearchLoggedModelsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *SearchLoggedModelsRequest_SdkV2) GetDatasets(ctx context.Context) ([]SearchLoggedModelsDataset_SdkV2, bool) {
	if o.Datasets.IsNull() || o.Datasets.IsUnknown() {
		return nil, false
	}
	var v []SearchLoggedModelsDataset_SdkV2
	d := o.Datasets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatasets sets the value of the Datasets field in SearchLoggedModelsRequest_SdkV2.
func (o *SearchLoggedModelsRequest_SdkV2) SetDatasets(ctx context.Context, v []SearchLoggedModelsDataset_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["datasets"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Datasets = types.ListValueMust(t, vs)
}

// GetExperimentIds returns the value of the ExperimentIds field in SearchLoggedModelsRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *SearchLoggedModelsRequest_SdkV2) GetExperimentIds(ctx context.Context) ([]types.String, bool) {
	if o.ExperimentIds.IsNull() || o.ExperimentIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.ExperimentIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExperimentIds sets the value of the ExperimentIds field in SearchLoggedModelsRequest_SdkV2.
func (o *SearchLoggedModelsRequest_SdkV2) SetExperimentIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["experiment_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ExperimentIds = types.ListValueMust(t, vs)
}

// GetOrderBy returns the value of the OrderBy field in SearchLoggedModelsRequest_SdkV2 as
// a slice of SearchLoggedModelsOrderBy_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *SearchLoggedModelsRequest_SdkV2) GetOrderBy(ctx context.Context) ([]SearchLoggedModelsOrderBy_SdkV2, bool) {
	if o.OrderBy.IsNull() || o.OrderBy.IsUnknown() {
		return nil, false
	}
	var v []SearchLoggedModelsOrderBy_SdkV2
	d := o.OrderBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOrderBy sets the value of the OrderBy field in SearchLoggedModelsRequest_SdkV2.
func (o *SearchLoggedModelsRequest_SdkV2) SetOrderBy(ctx context.Context, v []SearchLoggedModelsOrderBy_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["order_by"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OrderBy = types.ListValueMust(t, vs)
}

type SearchLoggedModelsResponse_SdkV2 struct {
	// Logged models that match the search criteria.
	Models types.List `tfsdk:"models"`
	// The token that can be used to retrieve the next page of logged models.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (toState *SearchLoggedModelsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SearchLoggedModelsResponse_SdkV2) {
}

func (toState *SearchLoggedModelsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState SearchLoggedModelsResponse_SdkV2) {
}

func (c SearchLoggedModelsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SearchLoggedModelsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"models": reflect.TypeOf(LoggedModel_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchLoggedModelsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o SearchLoggedModelsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"models":          o.Models,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SearchLoggedModelsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *SearchLoggedModelsResponse_SdkV2) GetModels(ctx context.Context) ([]LoggedModel_SdkV2, bool) {
	if o.Models.IsNull() || o.Models.IsUnknown() {
		return nil, false
	}
	var v []LoggedModel_SdkV2
	d := o.Models.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModels sets the value of the Models field in SearchLoggedModelsResponse_SdkV2.
func (o *SearchLoggedModelsResponse_SdkV2) SetModels(ctx context.Context, v []LoggedModel_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["models"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Models = types.ListValueMust(t, vs)
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

func (toState *SearchModelVersionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SearchModelVersionsRequest_SdkV2) {
}

func (toState *SearchModelVersionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState SearchModelVersionsRequest_SdkV2) {
}

func (c SearchModelVersionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SearchModelVersionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"order_by": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchModelVersionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o SearchModelVersionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter":      o.Filter,
			"max_results": o.MaxResults,
			"order_by":    o.OrderBy,
			"page_token":  o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SearchModelVersionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *SearchModelVersionsRequest_SdkV2) GetOrderBy(ctx context.Context) ([]types.String, bool) {
	if o.OrderBy.IsNull() || o.OrderBy.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.OrderBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOrderBy sets the value of the OrderBy field in SearchModelVersionsRequest_SdkV2.
func (o *SearchModelVersionsRequest_SdkV2) SetOrderBy(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["order_by"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OrderBy = types.ListValueMust(t, vs)
}

type SearchModelVersionsResponse_SdkV2 struct {
	// Models that match the search criteria
	ModelVersions types.List `tfsdk:"model_versions"`
	// Pagination token to request next page of models for the same search
	// query.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (toState *SearchModelVersionsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SearchModelVersionsResponse_SdkV2) {
}

func (toState *SearchModelVersionsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState SearchModelVersionsResponse_SdkV2) {
}

func (c SearchModelVersionsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SearchModelVersionsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model_versions": reflect.TypeOf(ModelVersion_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchModelVersionsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o SearchModelVersionsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_versions":  o.ModelVersions,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SearchModelVersionsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *SearchModelVersionsResponse_SdkV2) GetModelVersions(ctx context.Context) ([]ModelVersion_SdkV2, bool) {
	if o.ModelVersions.IsNull() || o.ModelVersions.IsUnknown() {
		return nil, false
	}
	var v []ModelVersion_SdkV2
	d := o.ModelVersions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModelVersions sets the value of the ModelVersions field in SearchModelVersionsResponse_SdkV2.
func (o *SearchModelVersionsResponse_SdkV2) SetModelVersions(ctx context.Context, v []ModelVersion_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["model_versions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ModelVersions = types.ListValueMust(t, vs)
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

func (toState *SearchModelsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SearchModelsRequest_SdkV2) {
}

func (toState *SearchModelsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState SearchModelsRequest_SdkV2) {
}

func (c SearchModelsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SearchModelsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"order_by": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchModelsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o SearchModelsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"filter":      o.Filter,
			"max_results": o.MaxResults,
			"order_by":    o.OrderBy,
			"page_token":  o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SearchModelsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *SearchModelsRequest_SdkV2) GetOrderBy(ctx context.Context) ([]types.String, bool) {
	if o.OrderBy.IsNull() || o.OrderBy.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.OrderBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOrderBy sets the value of the OrderBy field in SearchModelsRequest_SdkV2.
func (o *SearchModelsRequest_SdkV2) SetOrderBy(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["order_by"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OrderBy = types.ListValueMust(t, vs)
}

type SearchModelsResponse_SdkV2 struct {
	// Pagination token to request the next page of models.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// Registered Models that match the search criteria.
	RegisteredModels types.List `tfsdk:"registered_models"`
}

func (toState *SearchModelsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SearchModelsResponse_SdkV2) {
}

func (toState *SearchModelsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState SearchModelsResponse_SdkV2) {
}

func (c SearchModelsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SearchModelsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"registered_models": reflect.TypeOf(Model_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchModelsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o SearchModelsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":   o.NextPageToken,
			"registered_models": o.RegisteredModels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SearchModelsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *SearchModelsResponse_SdkV2) GetRegisteredModels(ctx context.Context) ([]Model_SdkV2, bool) {
	if o.RegisteredModels.IsNull() || o.RegisteredModels.IsUnknown() {
		return nil, false
	}
	var v []Model_SdkV2
	d := o.RegisteredModels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRegisteredModels sets the value of the RegisteredModels field in SearchModelsResponse_SdkV2.
func (o *SearchModelsResponse_SdkV2) SetRegisteredModels(ctx context.Context, v []Model_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["registered_models"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RegisteredModels = types.ListValueMust(t, vs)
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

func (toState *SearchRuns_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SearchRuns_SdkV2) {
}

func (toState *SearchRuns_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState SearchRuns_SdkV2) {
}

func (c SearchRuns_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SearchRuns_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"experiment_ids": reflect.TypeOf(types.String{}),
		"order_by":       reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchRuns_SdkV2
// only implements ToObjectValue() and Type().
func (o SearchRuns_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_ids": o.ExperimentIds,
			"filter":         o.Filter,
			"max_results":    o.MaxResults,
			"order_by":       o.OrderBy,
			"page_token":     o.PageToken,
			"run_view_type":  o.RunViewType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SearchRuns_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *SearchRuns_SdkV2) GetExperimentIds(ctx context.Context) ([]types.String, bool) {
	if o.ExperimentIds.IsNull() || o.ExperimentIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.ExperimentIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExperimentIds sets the value of the ExperimentIds field in SearchRuns_SdkV2.
func (o *SearchRuns_SdkV2) SetExperimentIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["experiment_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ExperimentIds = types.ListValueMust(t, vs)
}

// GetOrderBy returns the value of the OrderBy field in SearchRuns_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *SearchRuns_SdkV2) GetOrderBy(ctx context.Context) ([]types.String, bool) {
	if o.OrderBy.IsNull() || o.OrderBy.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.OrderBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOrderBy sets the value of the OrderBy field in SearchRuns_SdkV2.
func (o *SearchRuns_SdkV2) SetOrderBy(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["order_by"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OrderBy = types.ListValueMust(t, vs)
}

type SearchRunsResponse_SdkV2 struct {
	// Token for the next page of runs.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// Runs that match the search criteria.
	Runs types.List `tfsdk:"runs"`
}

func (toState *SearchRunsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SearchRunsResponse_SdkV2) {
}

func (toState *SearchRunsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState SearchRunsResponse_SdkV2) {
}

func (c SearchRunsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SearchRunsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"runs": reflect.TypeOf(Run_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchRunsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o SearchRunsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"runs":            o.Runs,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SearchRunsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *SearchRunsResponse_SdkV2) GetRuns(ctx context.Context) ([]Run_SdkV2, bool) {
	if o.Runs.IsNull() || o.Runs.IsUnknown() {
		return nil, false
	}
	var v []Run_SdkV2
	d := o.Runs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRuns sets the value of the Runs field in SearchRunsResponse_SdkV2.
func (o *SearchRunsResponse_SdkV2) SetRuns(ctx context.Context, v []Run_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["runs"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Runs = types.ListValueMust(t, vs)
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

func (toState *SetExperimentTag_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SetExperimentTag_SdkV2) {
}

func (toState *SetExperimentTag_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState SetExperimentTag_SdkV2) {
}

func (c SetExperimentTag_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SetExperimentTag_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetExperimentTag_SdkV2
// only implements ToObjectValue() and Type().
func (o SetExperimentTag_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": o.ExperimentId,
			"key":           o.Key,
			"value":         o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SetExperimentTag_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *SetExperimentTagResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SetExperimentTagResponse_SdkV2) {
}

func (toState *SetExperimentTagResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState SetExperimentTagResponse_SdkV2) {
}

func (c SetExperimentTagResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetExperimentTagResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SetExperimentTagResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetExperimentTagResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o SetExperimentTagResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o SetExperimentTagResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *SetLoggedModelTagsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SetLoggedModelTagsRequest_SdkV2) {
}

func (toState *SetLoggedModelTagsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState SetLoggedModelTagsRequest_SdkV2) {
}

func (c SetLoggedModelTagsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SetLoggedModelTagsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(LoggedModelTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetLoggedModelTagsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o SetLoggedModelTagsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_id": o.ModelId,
			"tags":     o.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SetLoggedModelTagsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *SetLoggedModelTagsRequest_SdkV2) GetTags(ctx context.Context) ([]LoggedModelTag_SdkV2, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []LoggedModelTag_SdkV2
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in SetLoggedModelTagsRequest_SdkV2.
func (o *SetLoggedModelTagsRequest_SdkV2) SetTags(ctx context.Context, v []LoggedModelTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type SetLoggedModelTagsResponse_SdkV2 struct {
}

func (toState *SetLoggedModelTagsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SetLoggedModelTagsResponse_SdkV2) {
}

func (toState *SetLoggedModelTagsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState SetLoggedModelTagsResponse_SdkV2) {
}

func (c SetLoggedModelTagsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetLoggedModelTagsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SetLoggedModelTagsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetLoggedModelTagsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o SetLoggedModelTagsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o SetLoggedModelTagsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *SetModelTagRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SetModelTagRequest_SdkV2) {
}

func (toState *SetModelTagRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState SetModelTagRequest_SdkV2) {
}

func (c SetModelTagRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SetModelTagRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetModelTagRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o SetModelTagRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"name":  o.Name,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SetModelTagRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *SetModelTagResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SetModelTagResponse_SdkV2) {
}

func (toState *SetModelTagResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState SetModelTagResponse_SdkV2) {
}

func (c SetModelTagResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetModelTagResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SetModelTagResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetModelTagResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o SetModelTagResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o SetModelTagResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *SetModelVersionTagRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SetModelVersionTagRequest_SdkV2) {
}

func (toState *SetModelVersionTagRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState SetModelVersionTagRequest_SdkV2) {
}

func (c SetModelVersionTagRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SetModelVersionTagRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetModelVersionTagRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o SetModelVersionTagRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":     o.Key,
			"name":    o.Name,
			"value":   o.Value,
			"version": o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SetModelVersionTagRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *SetModelVersionTagResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SetModelVersionTagResponse_SdkV2) {
}

func (toState *SetModelVersionTagResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState SetModelVersionTagResponse_SdkV2) {
}

func (c SetModelVersionTagResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetModelVersionTagResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SetModelVersionTagResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetModelVersionTagResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o SetModelVersionTagResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o SetModelVersionTagResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *SetTag_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SetTag_SdkV2) {
}

func (toState *SetTag_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState SetTag_SdkV2) {
}

func (c SetTag_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SetTag_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetTag_SdkV2
// only implements ToObjectValue() and Type().
func (o SetTag_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":      o.Key,
			"run_id":   o.RunId,
			"run_uuid": o.RunUuid,
			"value":    o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SetTag_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *SetTagResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SetTagResponse_SdkV2) {
}

func (toState *SetTagResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState SetTagResponse_SdkV2) {
}

func (c SetTagResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetTagResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SetTagResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetTagResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o SetTagResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o SetTagResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *TestRegistryWebhookRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan TestRegistryWebhookRequest_SdkV2) {
}

func (toState *TestRegistryWebhookRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState TestRegistryWebhookRequest_SdkV2) {
}

func (c TestRegistryWebhookRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a TestRegistryWebhookRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TestRegistryWebhookRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o TestRegistryWebhookRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"event": o.Event,
			"id":    o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TestRegistryWebhookRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *TestRegistryWebhookResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan TestRegistryWebhookResponse_SdkV2) {
}

func (toState *TestRegistryWebhookResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState TestRegistryWebhookResponse_SdkV2) {
}

func (c TestRegistryWebhookResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a TestRegistryWebhookResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TestRegistryWebhookResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o TestRegistryWebhookResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"body":        o.Body,
			"status_code": o.StatusCode,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TestRegistryWebhookResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"body":        types.StringType,
			"status_code": types.Int64Type,
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

func (toState *TransitionModelVersionStageDatabricks_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan TransitionModelVersionStageDatabricks_SdkV2) {
}

func (toState *TransitionModelVersionStageDatabricks_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState TransitionModelVersionStageDatabricks_SdkV2) {
}

func (c TransitionModelVersionStageDatabricks_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a TransitionModelVersionStageDatabricks_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TransitionModelVersionStageDatabricks_SdkV2
// only implements ToObjectValue() and Type().
func (o TransitionModelVersionStageDatabricks_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"archive_existing_versions": o.ArchiveExistingVersions,
			"comment":                   o.Comment,
			"name":                      o.Name,
			"stage":                     o.Stage,
			"version":                   o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TransitionModelVersionStageDatabricks_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *TransitionRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan TransitionRequest_SdkV2) {
}

func (toState *TransitionRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState TransitionRequest_SdkV2) {
}

func (c TransitionRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a TransitionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"available_actions": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TransitionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o TransitionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"available_actions":  o.AvailableActions,
			"comment":            o.Comment,
			"creation_timestamp": o.CreationTimestamp,
			"to_stage":           o.ToStage,
			"user_id":            o.UserId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TransitionRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *TransitionRequest_SdkV2) GetAvailableActions(ctx context.Context) ([]types.String, bool) {
	if o.AvailableActions.IsNull() || o.AvailableActions.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.AvailableActions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAvailableActions sets the value of the AvailableActions field in TransitionRequest_SdkV2.
func (o *TransitionRequest_SdkV2) SetAvailableActions(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["available_actions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AvailableActions = types.ListValueMust(t, vs)
}

type TransitionStageResponse_SdkV2 struct {
	// Updated model version
	ModelVersionDatabricks types.List `tfsdk:"model_version_databricks"`
}

func (toState *TransitionStageResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan TransitionStageResponse_SdkV2) {
	if !fromPlan.ModelVersionDatabricks.IsNull() && !fromPlan.ModelVersionDatabricks.IsUnknown() {
		if toStateModelVersionDatabricks, ok := toState.GetModelVersionDatabricks(ctx); ok {
			if fromPlanModelVersionDatabricks, ok := fromPlan.GetModelVersionDatabricks(ctx); ok {
				toStateModelVersionDatabricks.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanModelVersionDatabricks)
				toState.SetModelVersionDatabricks(ctx, toStateModelVersionDatabricks)
			}
		}
	}
}

func (toState *TransitionStageResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState TransitionStageResponse_SdkV2) {
	if !fromState.ModelVersionDatabricks.IsNull() && !fromState.ModelVersionDatabricks.IsUnknown() {
		if toStateModelVersionDatabricks, ok := toState.GetModelVersionDatabricks(ctx); ok {
			if fromStateModelVersionDatabricks, ok := fromState.GetModelVersionDatabricks(ctx); ok {
				toStateModelVersionDatabricks.SyncFieldsDuringRead(ctx, fromStateModelVersionDatabricks)
				toState.SetModelVersionDatabricks(ctx, toStateModelVersionDatabricks)
			}
		}
	}
}

func (c TransitionStageResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a TransitionStageResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model_version_databricks": reflect.TypeOf(ModelVersionDatabricks_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TransitionStageResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o TransitionStageResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_version_databricks": o.ModelVersionDatabricks,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TransitionStageResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *TransitionStageResponse_SdkV2) GetModelVersionDatabricks(ctx context.Context) (ModelVersionDatabricks_SdkV2, bool) {
	var e ModelVersionDatabricks_SdkV2
	if o.ModelVersionDatabricks.IsNull() || o.ModelVersionDatabricks.IsUnknown() {
		return e, false
	}
	var v []ModelVersionDatabricks_SdkV2
	d := o.ModelVersionDatabricks.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetModelVersionDatabricks sets the value of the ModelVersionDatabricks field in TransitionStageResponse_SdkV2.
func (o *TransitionStageResponse_SdkV2) SetModelVersionDatabricks(ctx context.Context, v ModelVersionDatabricks_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["model_version_databricks"]
	o.ModelVersionDatabricks = types.ListValueMust(t, vs)
}

// Details required to edit a comment on a model version.
type UpdateComment_SdkV2 struct {
	// User-provided comment on the action.
	Comment types.String `tfsdk:"comment"`
	// Unique identifier of an activity
	Id types.String `tfsdk:"id"`
}

func (toState *UpdateComment_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan UpdateComment_SdkV2) {
}

func (toState *UpdateComment_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState UpdateComment_SdkV2) {
}

func (c UpdateComment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateComment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateComment_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateComment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment": o.Comment,
			"id":      o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateComment_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *UpdateCommentResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan UpdateCommentResponse_SdkV2) {
	if !fromPlan.Comment.IsNull() && !fromPlan.Comment.IsUnknown() {
		if toStateComment, ok := toState.GetComment(ctx); ok {
			if fromPlanComment, ok := fromPlan.GetComment(ctx); ok {
				toStateComment.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanComment)
				toState.SetComment(ctx, toStateComment)
			}
		}
	}
}

func (toState *UpdateCommentResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState UpdateCommentResponse_SdkV2) {
	if !fromState.Comment.IsNull() && !fromState.Comment.IsUnknown() {
		if toStateComment, ok := toState.GetComment(ctx); ok {
			if fromStateComment, ok := fromState.GetComment(ctx); ok {
				toStateComment.SyncFieldsDuringRead(ctx, fromStateComment)
				toState.SetComment(ctx, toStateComment)
			}
		}
	}
}

func (c UpdateCommentResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateCommentResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"comment": reflect.TypeOf(CommentObject_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCommentResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateCommentResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment": o.Comment,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateCommentResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *UpdateCommentResponse_SdkV2) GetComment(ctx context.Context) (CommentObject_SdkV2, bool) {
	var e CommentObject_SdkV2
	if o.Comment.IsNull() || o.Comment.IsUnknown() {
		return e, false
	}
	var v []CommentObject_SdkV2
	d := o.Comment.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetComment sets the value of the Comment field in UpdateCommentResponse_SdkV2.
func (o *UpdateCommentResponse_SdkV2) SetComment(ctx context.Context, v CommentObject_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["comment"]
	o.Comment = types.ListValueMust(t, vs)
}

type UpdateExperiment_SdkV2 struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"experiment_id"`
	// If provided, the experiment's name is changed to the new name. The new
	// name must be unique.
	NewName types.String `tfsdk:"new_name"`
}

func (toState *UpdateExperiment_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan UpdateExperiment_SdkV2) {
}

func (toState *UpdateExperiment_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState UpdateExperiment_SdkV2) {
}

func (c UpdateExperiment_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateExperiment_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateExperiment_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateExperiment_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": o.ExperimentId,
			"new_name":      o.NewName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateExperiment_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
			"new_name":      types.StringType,
		},
	}
}

type UpdateExperimentResponse_SdkV2 struct {
}

func (toState *UpdateExperimentResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan UpdateExperimentResponse_SdkV2) {
}

func (toState *UpdateExperimentResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState UpdateExperimentResponse_SdkV2) {
}

func (c UpdateExperimentResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateExperimentResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateExperimentResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateExperimentResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateExperimentResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateExperimentResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateFeatureTagRequest_SdkV2 struct {
	FeatureName types.String `tfsdk:"-"`

	FeatureTag types.List `tfsdk:"feature_tag"`

	Key types.String `tfsdk:"-"`

	TableName types.String `tfsdk:"-"`
	// The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
}

func (toState *UpdateFeatureTagRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan UpdateFeatureTagRequest_SdkV2) {
	if !fromPlan.FeatureTag.IsNull() && !fromPlan.FeatureTag.IsUnknown() {
		if toStateFeatureTag, ok := toState.GetFeatureTag(ctx); ok {
			if fromPlanFeatureTag, ok := fromPlan.GetFeatureTag(ctx); ok {
				toStateFeatureTag.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanFeatureTag)
				toState.SetFeatureTag(ctx, toStateFeatureTag)
			}
		}
	}
}

func (toState *UpdateFeatureTagRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState UpdateFeatureTagRequest_SdkV2) {
	if !fromState.FeatureTag.IsNull() && !fromState.FeatureTag.IsUnknown() {
		if toStateFeatureTag, ok := toState.GetFeatureTag(ctx); ok {
			if fromStateFeatureTag, ok := fromState.GetFeatureTag(ctx); ok {
				toStateFeatureTag.SyncFieldsDuringRead(ctx, fromStateFeatureTag)
				toState.SetFeatureTag(ctx, toStateFeatureTag)
			}
		}
	}
}

func (c UpdateFeatureTagRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateFeatureTagRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"feature_tag": reflect.TypeOf(FeatureTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateFeatureTagRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateFeatureTagRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_name": o.FeatureName,
			"feature_tag":  o.FeatureTag,
			"key":          o.Key,
			"table_name":   o.TableName,
			"update_mask":  o.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateFeatureTagRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *UpdateFeatureTagRequest_SdkV2) GetFeatureTag(ctx context.Context) (FeatureTag_SdkV2, bool) {
	var e FeatureTag_SdkV2
	if o.FeatureTag.IsNull() || o.FeatureTag.IsUnknown() {
		return e, false
	}
	var v []FeatureTag_SdkV2
	d := o.FeatureTag.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFeatureTag sets the value of the FeatureTag field in UpdateFeatureTagRequest_SdkV2.
func (o *UpdateFeatureTagRequest_SdkV2) SetFeatureTag(ctx context.Context, v FeatureTag_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["feature_tag"]
	o.FeatureTag = types.ListValueMust(t, vs)
}

type UpdateModelRequest_SdkV2 struct {
	// If provided, updates the description for this `registered_model`.
	Description types.String `tfsdk:"description"`
	// Registered model unique name identifier.
	Name types.String `tfsdk:"name"`
}

func (toState *UpdateModelRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan UpdateModelRequest_SdkV2) {
}

func (toState *UpdateModelRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState UpdateModelRequest_SdkV2) {
}

func (c UpdateModelRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateModelRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateModelRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateModelRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description": o.Description,
			"name":        o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateModelRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *UpdateModelResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan UpdateModelResponse_SdkV2) {
	if !fromPlan.RegisteredModel.IsNull() && !fromPlan.RegisteredModel.IsUnknown() {
		if toStateRegisteredModel, ok := toState.GetRegisteredModel(ctx); ok {
			if fromPlanRegisteredModel, ok := fromPlan.GetRegisteredModel(ctx); ok {
				toStateRegisteredModel.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanRegisteredModel)
				toState.SetRegisteredModel(ctx, toStateRegisteredModel)
			}
		}
	}
}

func (toState *UpdateModelResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState UpdateModelResponse_SdkV2) {
	if !fromState.RegisteredModel.IsNull() && !fromState.RegisteredModel.IsUnknown() {
		if toStateRegisteredModel, ok := toState.GetRegisteredModel(ctx); ok {
			if fromStateRegisteredModel, ok := fromState.GetRegisteredModel(ctx); ok {
				toStateRegisteredModel.SyncFieldsDuringRead(ctx, fromStateRegisteredModel)
				toState.SetRegisteredModel(ctx, toStateRegisteredModel)
			}
		}
	}
}

func (c UpdateModelResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateModelResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"registered_model": reflect.TypeOf(Model_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateModelResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateModelResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"registered_model": o.RegisteredModel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateModelResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *UpdateModelResponse_SdkV2) GetRegisteredModel(ctx context.Context) (Model_SdkV2, bool) {
	var e Model_SdkV2
	if o.RegisteredModel.IsNull() || o.RegisteredModel.IsUnknown() {
		return e, false
	}
	var v []Model_SdkV2
	d := o.RegisteredModel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRegisteredModel sets the value of the RegisteredModel field in UpdateModelResponse_SdkV2.
func (o *UpdateModelResponse_SdkV2) SetRegisteredModel(ctx context.Context, v Model_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["registered_model"]
	o.RegisteredModel = types.ListValueMust(t, vs)
}

type UpdateModelVersionRequest_SdkV2 struct {
	// If provided, updates the description for this `registered_model`.
	Description types.String `tfsdk:"description"`
	// Name of the registered model
	Name types.String `tfsdk:"name"`
	// Model version number
	Version types.String `tfsdk:"version"`
}

func (toState *UpdateModelVersionRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan UpdateModelVersionRequest_SdkV2) {
}

func (toState *UpdateModelVersionRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState UpdateModelVersionRequest_SdkV2) {
}

func (c UpdateModelVersionRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateModelVersionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateModelVersionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateModelVersionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description": o.Description,
			"name":        o.Name,
			"version":     o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateModelVersionRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *UpdateModelVersionResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan UpdateModelVersionResponse_SdkV2) {
	if !fromPlan.ModelVersion.IsNull() && !fromPlan.ModelVersion.IsUnknown() {
		if toStateModelVersion, ok := toState.GetModelVersion(ctx); ok {
			if fromPlanModelVersion, ok := fromPlan.GetModelVersion(ctx); ok {
				toStateModelVersion.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanModelVersion)
				toState.SetModelVersion(ctx, toStateModelVersion)
			}
		}
	}
}

func (toState *UpdateModelVersionResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState UpdateModelVersionResponse_SdkV2) {
	if !fromState.ModelVersion.IsNull() && !fromState.ModelVersion.IsUnknown() {
		if toStateModelVersion, ok := toState.GetModelVersion(ctx); ok {
			if fromStateModelVersion, ok := fromState.GetModelVersion(ctx); ok {
				toStateModelVersion.SyncFieldsDuringRead(ctx, fromStateModelVersion)
				toState.SetModelVersion(ctx, toStateModelVersion)
			}
		}
	}
}

func (c UpdateModelVersionResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateModelVersionResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model_version": reflect.TypeOf(ModelVersion_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateModelVersionResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateModelVersionResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_version": o.ModelVersion,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateModelVersionResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *UpdateModelVersionResponse_SdkV2) GetModelVersion(ctx context.Context) (ModelVersion_SdkV2, bool) {
	var e ModelVersion_SdkV2
	if o.ModelVersion.IsNull() || o.ModelVersion.IsUnknown() {
		return e, false
	}
	var v []ModelVersion_SdkV2
	d := o.ModelVersion.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetModelVersion sets the value of the ModelVersion field in UpdateModelVersionResponse_SdkV2.
func (o *UpdateModelVersionResponse_SdkV2) SetModelVersion(ctx context.Context, v ModelVersion_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["model_version"]
	o.ModelVersion = types.ListValueMust(t, vs)
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

func (toState *UpdateOnlineStoreRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan UpdateOnlineStoreRequest_SdkV2) {
	if !fromPlan.OnlineStore.IsNull() && !fromPlan.OnlineStore.IsUnknown() {
		if toStateOnlineStore, ok := toState.GetOnlineStore(ctx); ok {
			if fromPlanOnlineStore, ok := fromPlan.GetOnlineStore(ctx); ok {
				toStateOnlineStore.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanOnlineStore)
				toState.SetOnlineStore(ctx, toStateOnlineStore)
			}
		}
	}
}

func (toState *UpdateOnlineStoreRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState UpdateOnlineStoreRequest_SdkV2) {
	if !fromState.OnlineStore.IsNull() && !fromState.OnlineStore.IsUnknown() {
		if toStateOnlineStore, ok := toState.GetOnlineStore(ctx); ok {
			if fromStateOnlineStore, ok := fromState.GetOnlineStore(ctx); ok {
				toStateOnlineStore.SyncFieldsDuringRead(ctx, fromStateOnlineStore)
				toState.SetOnlineStore(ctx, toStateOnlineStore)
			}
		}
	}
}

func (c UpdateOnlineStoreRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateOnlineStoreRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"online_store": reflect.TypeOf(OnlineStore_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateOnlineStoreRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateOnlineStoreRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":         o.Name,
			"online_store": o.OnlineStore,
			"update_mask":  o.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateOnlineStoreRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *UpdateOnlineStoreRequest_SdkV2) GetOnlineStore(ctx context.Context) (OnlineStore_SdkV2, bool) {
	var e OnlineStore_SdkV2
	if o.OnlineStore.IsNull() || o.OnlineStore.IsUnknown() {
		return e, false
	}
	var v []OnlineStore_SdkV2
	d := o.OnlineStore.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOnlineStore sets the value of the OnlineStore field in UpdateOnlineStoreRequest_SdkV2.
func (o *UpdateOnlineStoreRequest_SdkV2) SetOnlineStore(ctx context.Context, v OnlineStore_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["online_store"]
	o.OnlineStore = types.ListValueMust(t, vs)
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

func (toState *UpdateRegistryWebhook_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan UpdateRegistryWebhook_SdkV2) {
	if !fromPlan.HttpUrlSpec.IsNull() && !fromPlan.HttpUrlSpec.IsUnknown() {
		if toStateHttpUrlSpec, ok := toState.GetHttpUrlSpec(ctx); ok {
			if fromPlanHttpUrlSpec, ok := fromPlan.GetHttpUrlSpec(ctx); ok {
				toStateHttpUrlSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanHttpUrlSpec)
				toState.SetHttpUrlSpec(ctx, toStateHttpUrlSpec)
			}
		}
	}
	if !fromPlan.JobSpec.IsNull() && !fromPlan.JobSpec.IsUnknown() {
		if toStateJobSpec, ok := toState.GetJobSpec(ctx); ok {
			if fromPlanJobSpec, ok := fromPlan.GetJobSpec(ctx); ok {
				toStateJobSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanJobSpec)
				toState.SetJobSpec(ctx, toStateJobSpec)
			}
		}
	}
}

func (toState *UpdateRegistryWebhook_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState UpdateRegistryWebhook_SdkV2) {
	if !fromState.HttpUrlSpec.IsNull() && !fromState.HttpUrlSpec.IsUnknown() {
		if toStateHttpUrlSpec, ok := toState.GetHttpUrlSpec(ctx); ok {
			if fromStateHttpUrlSpec, ok := fromState.GetHttpUrlSpec(ctx); ok {
				toStateHttpUrlSpec.SyncFieldsDuringRead(ctx, fromStateHttpUrlSpec)
				toState.SetHttpUrlSpec(ctx, toStateHttpUrlSpec)
			}
		}
	}
	if !fromState.JobSpec.IsNull() && !fromState.JobSpec.IsUnknown() {
		if toStateJobSpec, ok := toState.GetJobSpec(ctx); ok {
			if fromStateJobSpec, ok := fromState.GetJobSpec(ctx); ok {
				toStateJobSpec.SyncFieldsDuringRead(ctx, fromStateJobSpec)
				toState.SetJobSpec(ctx, toStateJobSpec)
			}
		}
	}
}

func (c UpdateRegistryWebhook_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateRegistryWebhook_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"events":        reflect.TypeOf(types.String{}),
		"http_url_spec": reflect.TypeOf(HttpUrlSpec_SdkV2{}),
		"job_spec":      reflect.TypeOf(JobSpec_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRegistryWebhook_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateRegistryWebhook_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":   o.Description,
			"events":        o.Events,
			"http_url_spec": o.HttpUrlSpec,
			"id":            o.Id,
			"job_spec":      o.JobSpec,
			"status":        o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateRegistryWebhook_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *UpdateRegistryWebhook_SdkV2) GetEvents(ctx context.Context) ([]types.String, bool) {
	if o.Events.IsNull() || o.Events.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Events.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEvents sets the value of the Events field in UpdateRegistryWebhook_SdkV2.
func (o *UpdateRegistryWebhook_SdkV2) SetEvents(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["events"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Events = types.ListValueMust(t, vs)
}

// GetHttpUrlSpec returns the value of the HttpUrlSpec field in UpdateRegistryWebhook_SdkV2 as
// a HttpUrlSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateRegistryWebhook_SdkV2) GetHttpUrlSpec(ctx context.Context) (HttpUrlSpec_SdkV2, bool) {
	var e HttpUrlSpec_SdkV2
	if o.HttpUrlSpec.IsNull() || o.HttpUrlSpec.IsUnknown() {
		return e, false
	}
	var v []HttpUrlSpec_SdkV2
	d := o.HttpUrlSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetHttpUrlSpec sets the value of the HttpUrlSpec field in UpdateRegistryWebhook_SdkV2.
func (o *UpdateRegistryWebhook_SdkV2) SetHttpUrlSpec(ctx context.Context, v HttpUrlSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["http_url_spec"]
	o.HttpUrlSpec = types.ListValueMust(t, vs)
}

// GetJobSpec returns the value of the JobSpec field in UpdateRegistryWebhook_SdkV2 as
// a JobSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateRegistryWebhook_SdkV2) GetJobSpec(ctx context.Context) (JobSpec_SdkV2, bool) {
	var e JobSpec_SdkV2
	if o.JobSpec.IsNull() || o.JobSpec.IsUnknown() {
		return e, false
	}
	var v []JobSpec_SdkV2
	d := o.JobSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetJobSpec sets the value of the JobSpec field in UpdateRegistryWebhook_SdkV2.
func (o *UpdateRegistryWebhook_SdkV2) SetJobSpec(ctx context.Context, v JobSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["job_spec"]
	o.JobSpec = types.ListValueMust(t, vs)
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

func (toState *UpdateRun_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan UpdateRun_SdkV2) {
}

func (toState *UpdateRun_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState UpdateRun_SdkV2) {
}

func (c UpdateRun_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateRun_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRun_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateRun_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"end_time": o.EndTime,
			"run_id":   o.RunId,
			"run_name": o.RunName,
			"run_uuid": o.RunUuid,
			"status":   o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateRun_SdkV2) Type(ctx context.Context) attr.Type {
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

func (toState *UpdateRunResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan UpdateRunResponse_SdkV2) {
	if !fromPlan.RunInfo.IsNull() && !fromPlan.RunInfo.IsUnknown() {
		if toStateRunInfo, ok := toState.GetRunInfo(ctx); ok {
			if fromPlanRunInfo, ok := fromPlan.GetRunInfo(ctx); ok {
				toStateRunInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanRunInfo)
				toState.SetRunInfo(ctx, toStateRunInfo)
			}
		}
	}
}

func (toState *UpdateRunResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState UpdateRunResponse_SdkV2) {
	if !fromState.RunInfo.IsNull() && !fromState.RunInfo.IsUnknown() {
		if toStateRunInfo, ok := toState.GetRunInfo(ctx); ok {
			if fromStateRunInfo, ok := fromState.GetRunInfo(ctx); ok {
				toStateRunInfo.SyncFieldsDuringRead(ctx, fromStateRunInfo)
				toState.SetRunInfo(ctx, toStateRunInfo)
			}
		}
	}
}

func (c UpdateRunResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateRunResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"run_info": reflect.TypeOf(RunInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRunResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateRunResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run_info": o.RunInfo,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateRunResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *UpdateRunResponse_SdkV2) GetRunInfo(ctx context.Context) (RunInfo_SdkV2, bool) {
	var e RunInfo_SdkV2
	if o.RunInfo.IsNull() || o.RunInfo.IsUnknown() {
		return e, false
	}
	var v []RunInfo_SdkV2
	d := o.RunInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRunInfo sets the value of the RunInfo field in UpdateRunResponse_SdkV2.
func (o *UpdateRunResponse_SdkV2) SetRunInfo(ctx context.Context, v RunInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["run_info"]
	o.RunInfo = types.ListValueMust(t, vs)
}

type UpdateWebhookResponse_SdkV2 struct {
	Webhook types.List `tfsdk:"webhook"`
}

func (toState *UpdateWebhookResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan UpdateWebhookResponse_SdkV2) {
	if !fromPlan.Webhook.IsNull() && !fromPlan.Webhook.IsUnknown() {
		if toStateWebhook, ok := toState.GetWebhook(ctx); ok {
			if fromPlanWebhook, ok := fromPlan.GetWebhook(ctx); ok {
				toStateWebhook.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanWebhook)
				toState.SetWebhook(ctx, toStateWebhook)
			}
		}
	}
}

func (toState *UpdateWebhookResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState UpdateWebhookResponse_SdkV2) {
	if !fromState.Webhook.IsNull() && !fromState.Webhook.IsUnknown() {
		if toStateWebhook, ok := toState.GetWebhook(ctx); ok {
			if fromStateWebhook, ok := fromState.GetWebhook(ctx); ok {
				toStateWebhook.SyncFieldsDuringRead(ctx, fromStateWebhook)
				toState.SetWebhook(ctx, toStateWebhook)
			}
		}
	}
}

func (c UpdateWebhookResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateWebhookResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"webhook": reflect.TypeOf(RegistryWebhook_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWebhookResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateWebhookResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"webhook": o.Webhook,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateWebhookResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *UpdateWebhookResponse_SdkV2) GetWebhook(ctx context.Context) (RegistryWebhook_SdkV2, bool) {
	var e RegistryWebhook_SdkV2
	if o.Webhook.IsNull() || o.Webhook.IsUnknown() {
		return e, false
	}
	var v []RegistryWebhook_SdkV2
	d := o.Webhook.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWebhook sets the value of the Webhook field in UpdateWebhookResponse_SdkV2.
func (o *UpdateWebhookResponse_SdkV2) SetWebhook(ctx context.Context, v RegistryWebhook_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["webhook"]
	o.Webhook = types.ListValueMust(t, vs)
}
