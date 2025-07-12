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

func (newState *Activity) SyncEffectiveFieldsDuringCreateOrUpdate(plan Activity) {
}

func (newState *Activity) SyncEffectiveFieldsDuringRead(existingState Activity) {
}

func (c Activity) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Activity) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Activity
// only implements ToObjectValue() and Type().
func (o Activity) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o Activity) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ApproveTransitionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ApproveTransitionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ApproveTransitionRequest
// only implements ToObjectValue() and Type().
func (o ApproveTransitionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ApproveTransitionRequest) Type(ctx context.Context) attr.Type {
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

func (newState *ApproveTransitionRequestResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ApproveTransitionRequestResponse) {
}

func (newState *ApproveTransitionRequestResponse) SyncEffectiveFieldsDuringRead(existingState ApproveTransitionRequestResponse) {
}

func (c ApproveTransitionRequestResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ApproveTransitionRequestResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"activity": reflect.TypeOf(Activity{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ApproveTransitionRequestResponse
// only implements ToObjectValue() and Type().
func (o ApproveTransitionRequestResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"activity": o.Activity,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ApproveTransitionRequestResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"activity": Activity{}.Type(ctx),
		},
	}
}

// GetActivity returns the value of the Activity field in ApproveTransitionRequestResponse as
// a Activity value.
// If the field is unknown or null, the boolean return value is false.
func (o *ApproveTransitionRequestResponse) GetActivity(ctx context.Context) (Activity, bool) {
	var e Activity
	if o.Activity.IsNull() || o.Activity.IsUnknown() {
		return e, false
	}
	var v []Activity
	d := o.Activity.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetActivity sets the value of the Activity field in ApproveTransitionRequestResponse.
func (o *ApproveTransitionRequestResponse) SetActivity(ctx context.Context, v Activity) {
	vs := v.ToObjectValue(ctx)
	o.Activity = vs
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

func (newState *CommentObject) SyncEffectiveFieldsDuringCreateOrUpdate(plan CommentObject) {
}

func (newState *CommentObject) SyncEffectiveFieldsDuringRead(existingState CommentObject) {
}

func (c CommentObject) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CommentObject) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"available_actions": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CommentObject
// only implements ToObjectValue() and Type().
func (o CommentObject) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CommentObject) Type(ctx context.Context) attr.Type {
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
func (o *CommentObject) GetAvailableActions(ctx context.Context) ([]types.String, bool) {
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

// SetAvailableActions sets the value of the AvailableActions field in CommentObject.
func (o *CommentObject) SetAvailableActions(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["available_actions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AvailableActions = types.ListValueMust(t, vs)
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateComment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateComment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateComment
// only implements ToObjectValue() and Type().
func (o CreateComment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment": o.Comment,
			"name":    o.Name,
			"version": o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateComment) Type(ctx context.Context) attr.Type {
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

func (newState *CreateCommentResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCommentResponse) {
}

func (newState *CreateCommentResponse) SyncEffectiveFieldsDuringRead(existingState CreateCommentResponse) {
}

func (c CreateCommentResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateCommentResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"comment": reflect.TypeOf(CommentObject{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCommentResponse
// only implements ToObjectValue() and Type().
func (o CreateCommentResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment": o.Comment,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCommentResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment": CommentObject{}.Type(ctx),
		},
	}
}

// GetComment returns the value of the Comment field in CreateCommentResponse as
// a CommentObject value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCommentResponse) GetComment(ctx context.Context) (CommentObject, bool) {
	var e CommentObject
	if o.Comment.IsNull() || o.Comment.IsUnknown() {
		return e, false
	}
	var v []CommentObject
	d := o.Comment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetComment sets the value of the Comment field in CreateCommentResponse.
func (o *CreateCommentResponse) SetComment(ctx context.Context, v CommentObject) {
	vs := v.ToObjectValue(ctx)
	o.Comment = vs
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateExperiment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateExperiment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(ExperimentTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateExperiment
// only implements ToObjectValue() and Type().
func (o CreateExperiment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"artifact_location": o.ArtifactLocation,
			"name":              o.Name,
			"tags":              o.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateExperiment) Type(ctx context.Context) attr.Type {
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
func (o *CreateExperiment) GetTags(ctx context.Context) ([]ExperimentTag, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []ExperimentTag
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in CreateExperiment.
func (o *CreateExperiment) SetTags(ctx context.Context, v []ExperimentTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type CreateExperimentResponse struct {
	// Unique identifier for the experiment.
	ExperimentId types.String `tfsdk:"experiment_id"`
}

func (newState *CreateExperimentResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateExperimentResponse) {
}

func (newState *CreateExperimentResponse) SyncEffectiveFieldsDuringRead(existingState CreateExperimentResponse) {
}

func (c CreateExperimentResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateExperimentResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateExperimentResponse
// only implements ToObjectValue() and Type().
func (o CreateExperimentResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": o.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateExperimentResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
		},
	}
}

type CreateFeatureTagRequest struct {
	FeatureName types.String `tfsdk:"-"`

	FeatureTag types.Object `tfsdk:"feature_tag"`

	TableName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateFeatureTagRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateFeatureTagRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"feature_tag": reflect.TypeOf(FeatureTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateFeatureTagRequest
// only implements ToObjectValue() and Type().
func (o CreateFeatureTagRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_name": o.FeatureName,
			"feature_tag":  o.FeatureTag,
			"table_name":   o.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateFeatureTagRequest) Type(ctx context.Context) attr.Type {
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
func (o *CreateFeatureTagRequest) GetFeatureTag(ctx context.Context) (FeatureTag, bool) {
	var e FeatureTag
	if o.FeatureTag.IsNull() || o.FeatureTag.IsUnknown() {
		return e, false
	}
	var v []FeatureTag
	d := o.FeatureTag.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFeatureTag sets the value of the FeatureTag field in CreateFeatureTagRequest.
func (o *CreateFeatureTagRequest) SetFeatureTag(ctx context.Context, v FeatureTag) {
	vs := v.ToObjectValue(ctx)
	o.FeatureTag = vs
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateForecastingExperimentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateForecastingExperimentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (o CreateForecastingExperimentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateForecastingExperimentRequest) Type(ctx context.Context) attr.Type {
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
func (o *CreateForecastingExperimentRequest) GetHolidayRegions(ctx context.Context) ([]types.String, bool) {
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

// SetHolidayRegions sets the value of the HolidayRegions field in CreateForecastingExperimentRequest.
func (o *CreateForecastingExperimentRequest) SetHolidayRegions(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["holiday_regions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.HolidayRegions = types.ListValueMust(t, vs)
}

// GetIncludeFeatures returns the value of the IncludeFeatures field in CreateForecastingExperimentRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateForecastingExperimentRequest) GetIncludeFeatures(ctx context.Context) ([]types.String, bool) {
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

// SetIncludeFeatures sets the value of the IncludeFeatures field in CreateForecastingExperimentRequest.
func (o *CreateForecastingExperimentRequest) SetIncludeFeatures(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["include_features"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.IncludeFeatures = types.ListValueMust(t, vs)
}

// GetTimeseriesIdentifierColumns returns the value of the TimeseriesIdentifierColumns field in CreateForecastingExperimentRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateForecastingExperimentRequest) GetTimeseriesIdentifierColumns(ctx context.Context) ([]types.String, bool) {
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

// SetTimeseriesIdentifierColumns sets the value of the TimeseriesIdentifierColumns field in CreateForecastingExperimentRequest.
func (o *CreateForecastingExperimentRequest) SetTimeseriesIdentifierColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["timeseries_identifier_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.TimeseriesIdentifierColumns = types.ListValueMust(t, vs)
}

// GetTrainingFrameworks returns the value of the TrainingFrameworks field in CreateForecastingExperimentRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateForecastingExperimentRequest) GetTrainingFrameworks(ctx context.Context) ([]types.String, bool) {
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

// SetTrainingFrameworks sets the value of the TrainingFrameworks field in CreateForecastingExperimentRequest.
func (o *CreateForecastingExperimentRequest) SetTrainingFrameworks(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["training_frameworks"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.TrainingFrameworks = types.ListValueMust(t, vs)
}

type CreateForecastingExperimentResponse struct {
	// The unique ID of the created forecasting experiment
	ExperimentId types.String `tfsdk:"experiment_id"`
}

func (newState *CreateForecastingExperimentResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateForecastingExperimentResponse) {
}

func (newState *CreateForecastingExperimentResponse) SyncEffectiveFieldsDuringRead(existingState CreateForecastingExperimentResponse) {
}

func (c CreateForecastingExperimentResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateForecastingExperimentResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateForecastingExperimentResponse
// only implements ToObjectValue() and Type().
func (o CreateForecastingExperimentResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": o.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateForecastingExperimentResponse) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateLoggedModelRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateLoggedModelRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"params": reflect.TypeOf(LoggedModelParameter{}),
		"tags":   reflect.TypeOf(LoggedModelTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateLoggedModelRequest
// only implements ToObjectValue() and Type().
func (o CreateLoggedModelRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateLoggedModelRequest) Type(ctx context.Context) attr.Type {
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
func (o *CreateLoggedModelRequest) GetParams(ctx context.Context) ([]LoggedModelParameter, bool) {
	if o.Params.IsNull() || o.Params.IsUnknown() {
		return nil, false
	}
	var v []LoggedModelParameter
	d := o.Params.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParams sets the value of the Params field in CreateLoggedModelRequest.
func (o *CreateLoggedModelRequest) SetParams(ctx context.Context, v []LoggedModelParameter) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Params = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in CreateLoggedModelRequest as
// a slice of LoggedModelTag values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateLoggedModelRequest) GetTags(ctx context.Context) ([]LoggedModelTag, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []LoggedModelTag
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in CreateLoggedModelRequest.
func (o *CreateLoggedModelRequest) SetTags(ctx context.Context, v []LoggedModelTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type CreateLoggedModelResponse struct {
	// The newly created logged model.
	Model types.Object `tfsdk:"model"`
}

func (newState *CreateLoggedModelResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateLoggedModelResponse) {
}

func (newState *CreateLoggedModelResponse) SyncEffectiveFieldsDuringRead(existingState CreateLoggedModelResponse) {
}

func (c CreateLoggedModelResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateLoggedModelResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model": reflect.TypeOf(LoggedModel{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateLoggedModelResponse
// only implements ToObjectValue() and Type().
func (o CreateLoggedModelResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model": o.Model,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateLoggedModelResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model": LoggedModel{}.Type(ctx),
		},
	}
}

// GetModel returns the value of the Model field in CreateLoggedModelResponse as
// a LoggedModel value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateLoggedModelResponse) GetModel(ctx context.Context) (LoggedModel, bool) {
	var e LoggedModel
	if o.Model.IsNull() || o.Model.IsUnknown() {
		return e, false
	}
	var v []LoggedModel
	d := o.Model.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetModel sets the value of the Model field in CreateLoggedModelResponse.
func (o *CreateLoggedModelResponse) SetModel(ctx context.Context, v LoggedModel) {
	vs := v.ToObjectValue(ctx)
	o.Model = vs
}

type CreateModelRequest struct {
	// Optional description for registered model.
	Description types.String `tfsdk:"description"`
	// Register models under this name
	Name types.String `tfsdk:"name"`
	// Additional metadata for registered model.
	Tags types.List `tfsdk:"tags"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateModelRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateModelRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(ModelTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateModelRequest
// only implements ToObjectValue() and Type().
func (o CreateModelRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description": o.Description,
			"name":        o.Name,
			"tags":        o.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateModelRequest) Type(ctx context.Context) attr.Type {
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
func (o *CreateModelRequest) GetTags(ctx context.Context) ([]ModelTag, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []ModelTag
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in CreateModelRequest.
func (o *CreateModelRequest) SetTags(ctx context.Context, v []ModelTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type CreateModelResponse struct {
	RegisteredModel types.Object `tfsdk:"registered_model"`
}

func (newState *CreateModelResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateModelResponse) {
}

func (newState *CreateModelResponse) SyncEffectiveFieldsDuringRead(existingState CreateModelResponse) {
}

func (c CreateModelResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateModelResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"registered_model": reflect.TypeOf(Model{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateModelResponse
// only implements ToObjectValue() and Type().
func (o CreateModelResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"registered_model": o.RegisteredModel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateModelResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"registered_model": Model{}.Type(ctx),
		},
	}
}

// GetRegisteredModel returns the value of the RegisteredModel field in CreateModelResponse as
// a Model value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateModelResponse) GetRegisteredModel(ctx context.Context) (Model, bool) {
	var e Model
	if o.RegisteredModel.IsNull() || o.RegisteredModel.IsUnknown() {
		return e, false
	}
	var v []Model
	d := o.RegisteredModel.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRegisteredModel sets the value of the RegisteredModel field in CreateModelResponse.
func (o *CreateModelResponse) SetRegisteredModel(ctx context.Context, v Model) {
	vs := v.ToObjectValue(ctx)
	o.RegisteredModel = vs
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateModelVersionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateModelVersionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(ModelVersionTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateModelVersionRequest
// only implements ToObjectValue() and Type().
func (o CreateModelVersionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateModelVersionRequest) Type(ctx context.Context) attr.Type {
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
func (o *CreateModelVersionRequest) GetTags(ctx context.Context) ([]ModelVersionTag, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []ModelVersionTag
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in CreateModelVersionRequest.
func (o *CreateModelVersionRequest) SetTags(ctx context.Context, v []ModelVersionTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type CreateModelVersionResponse struct {
	// Return new version number generated for this model in registry.
	ModelVersion types.Object `tfsdk:"model_version"`
}

func (newState *CreateModelVersionResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateModelVersionResponse) {
}

func (newState *CreateModelVersionResponse) SyncEffectiveFieldsDuringRead(existingState CreateModelVersionResponse) {
}

func (c CreateModelVersionResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateModelVersionResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model_version": reflect.TypeOf(ModelVersion{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateModelVersionResponse
// only implements ToObjectValue() and Type().
func (o CreateModelVersionResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_version": o.ModelVersion,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateModelVersionResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_version": ModelVersion{}.Type(ctx),
		},
	}
}

// GetModelVersion returns the value of the ModelVersion field in CreateModelVersionResponse as
// a ModelVersion value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateModelVersionResponse) GetModelVersion(ctx context.Context) (ModelVersion, bool) {
	var e ModelVersion
	if o.ModelVersion.IsNull() || o.ModelVersion.IsUnknown() {
		return e, false
	}
	var v []ModelVersion
	d := o.ModelVersion.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetModelVersion sets the value of the ModelVersion field in CreateModelVersionResponse.
func (o *CreateModelVersionResponse) SetModelVersion(ctx context.Context, v ModelVersion) {
	vs := v.ToObjectValue(ctx)
	o.ModelVersion = vs
}

type CreateOnlineStoreRequest struct {
	// Online store to create.
	OnlineStore types.Object `tfsdk:"online_store"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateOnlineStoreRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateOnlineStoreRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"online_store": reflect.TypeOf(OnlineStore{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateOnlineStoreRequest
// only implements ToObjectValue() and Type().
func (o CreateOnlineStoreRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"online_store": o.OnlineStore,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateOnlineStoreRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"online_store": OnlineStore{}.Type(ctx),
		},
	}
}

// GetOnlineStore returns the value of the OnlineStore field in CreateOnlineStoreRequest as
// a OnlineStore value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateOnlineStoreRequest) GetOnlineStore(ctx context.Context) (OnlineStore, bool) {
	var e OnlineStore
	if o.OnlineStore.IsNull() || o.OnlineStore.IsUnknown() {
		return e, false
	}
	var v []OnlineStore
	d := o.OnlineStore.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOnlineStore sets the value of the OnlineStore field in CreateOnlineStoreRequest.
func (o *CreateOnlineStoreRequest) SetOnlineStore(ctx context.Context, v OnlineStore) {
	vs := v.ToObjectValue(ctx)
	o.OnlineStore = vs
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateRegistryWebhook.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateRegistryWebhook) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"events":        reflect.TypeOf(types.String{}),
		"http_url_spec": reflect.TypeOf(HttpUrlSpec{}),
		"job_spec":      reflect.TypeOf(JobSpec{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateRegistryWebhook
// only implements ToObjectValue() and Type().
func (o CreateRegistryWebhook) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateRegistryWebhook) Type(ctx context.Context) attr.Type {
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
func (o *CreateRegistryWebhook) GetEvents(ctx context.Context) ([]types.String, bool) {
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

// SetEvents sets the value of the Events field in CreateRegistryWebhook.
func (o *CreateRegistryWebhook) SetEvents(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["events"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Events = types.ListValueMust(t, vs)
}

// GetHttpUrlSpec returns the value of the HttpUrlSpec field in CreateRegistryWebhook as
// a HttpUrlSpec value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateRegistryWebhook) GetHttpUrlSpec(ctx context.Context) (HttpUrlSpec, bool) {
	var e HttpUrlSpec
	if o.HttpUrlSpec.IsNull() || o.HttpUrlSpec.IsUnknown() {
		return e, false
	}
	var v []HttpUrlSpec
	d := o.HttpUrlSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetHttpUrlSpec sets the value of the HttpUrlSpec field in CreateRegistryWebhook.
func (o *CreateRegistryWebhook) SetHttpUrlSpec(ctx context.Context, v HttpUrlSpec) {
	vs := v.ToObjectValue(ctx)
	o.HttpUrlSpec = vs
}

// GetJobSpec returns the value of the JobSpec field in CreateRegistryWebhook as
// a JobSpec value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateRegistryWebhook) GetJobSpec(ctx context.Context) (JobSpec, bool) {
	var e JobSpec
	if o.JobSpec.IsNull() || o.JobSpec.IsUnknown() {
		return e, false
	}
	var v []JobSpec
	d := o.JobSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetJobSpec sets the value of the JobSpec field in CreateRegistryWebhook.
func (o *CreateRegistryWebhook) SetJobSpec(ctx context.Context, v JobSpec) {
	vs := v.ToObjectValue(ctx)
	o.JobSpec = vs
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateRun.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateRun) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(RunTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateRun
// only implements ToObjectValue() and Type().
func (o CreateRun) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateRun) Type(ctx context.Context) attr.Type {
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
func (o *CreateRun) GetTags(ctx context.Context) ([]RunTag, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []RunTag
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in CreateRun.
func (o *CreateRun) SetTags(ctx context.Context, v []RunTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type CreateRunResponse struct {
	// The newly created run.
	Run types.Object `tfsdk:"run"`
}

func (newState *CreateRunResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateRunResponse) {
}

func (newState *CreateRunResponse) SyncEffectiveFieldsDuringRead(existingState CreateRunResponse) {
}

func (c CreateRunResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateRunResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"run": reflect.TypeOf(Run{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateRunResponse
// only implements ToObjectValue() and Type().
func (o CreateRunResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run": o.Run,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateRunResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run": Run{}.Type(ctx),
		},
	}
}

// GetRun returns the value of the Run field in CreateRunResponse as
// a Run value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateRunResponse) GetRun(ctx context.Context) (Run, bool) {
	var e Run
	if o.Run.IsNull() || o.Run.IsUnknown() {
		return e, false
	}
	var v []Run
	d := o.Run.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRun sets the value of the Run field in CreateRunResponse.
func (o *CreateRunResponse) SetRun(ctx context.Context, v Run) {
	vs := v.ToObjectValue(ctx)
	o.Run = vs
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateTransitionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateTransitionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateTransitionRequest
// only implements ToObjectValue() and Type().
func (o CreateTransitionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateTransitionRequest) Type(ctx context.Context) attr.Type {
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

func (newState *CreateTransitionRequestResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateTransitionRequestResponse) {
}

func (newState *CreateTransitionRequestResponse) SyncEffectiveFieldsDuringRead(existingState CreateTransitionRequestResponse) {
}

func (c CreateTransitionRequestResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateTransitionRequestResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"request": reflect.TypeOf(TransitionRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateTransitionRequestResponse
// only implements ToObjectValue() and Type().
func (o CreateTransitionRequestResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"request": o.Request,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateTransitionRequestResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"request": TransitionRequest{}.Type(ctx),
		},
	}
}

// GetRequest returns the value of the Request field in CreateTransitionRequestResponse as
// a TransitionRequest value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateTransitionRequestResponse) GetRequest(ctx context.Context) (TransitionRequest, bool) {
	var e TransitionRequest
	if o.Request.IsNull() || o.Request.IsUnknown() {
		return e, false
	}
	var v []TransitionRequest
	d := o.Request.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRequest sets the value of the Request field in CreateTransitionRequestResponse.
func (o *CreateTransitionRequestResponse) SetRequest(ctx context.Context, v TransitionRequest) {
	vs := v.ToObjectValue(ctx)
	o.Request = vs
}

type CreateWebhookResponse struct {
	Webhook types.Object `tfsdk:"webhook"`
}

func (newState *CreateWebhookResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateWebhookResponse) {
}

func (newState *CreateWebhookResponse) SyncEffectiveFieldsDuringRead(existingState CreateWebhookResponse) {
}

func (c CreateWebhookResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateWebhookResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"webhook": reflect.TypeOf(RegistryWebhook{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateWebhookResponse
// only implements ToObjectValue() and Type().
func (o CreateWebhookResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"webhook": o.Webhook,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateWebhookResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"webhook": RegistryWebhook{}.Type(ctx),
		},
	}
}

// GetWebhook returns the value of the Webhook field in CreateWebhookResponse as
// a RegistryWebhook value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateWebhookResponse) GetWebhook(ctx context.Context) (RegistryWebhook, bool) {
	var e RegistryWebhook
	if o.Webhook.IsNull() || o.Webhook.IsUnknown() {
		return e, false
	}
	var v []RegistryWebhook
	d := o.Webhook.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWebhook sets the value of the Webhook field in CreateWebhookResponse.
func (o *CreateWebhookResponse) SetWebhook(ctx context.Context, v RegistryWebhook) {
	vs := v.ToObjectValue(ctx)
	o.Webhook = vs
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

func (newState *Dataset) SyncEffectiveFieldsDuringCreateOrUpdate(plan Dataset) {
}

func (newState *Dataset) SyncEffectiveFieldsDuringRead(existingState Dataset) {
}

func (c Dataset) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Dataset) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Dataset
// only implements ToObjectValue() and Type().
func (o Dataset) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o Dataset) Type(ctx context.Context) attr.Type {
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

func (newState *DatasetInput) SyncEffectiveFieldsDuringCreateOrUpdate(plan DatasetInput) {
}

func (newState *DatasetInput) SyncEffectiveFieldsDuringRead(existingState DatasetInput) {
}

func (c DatasetInput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DatasetInput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dataset": reflect.TypeOf(Dataset{}),
		"tags":    reflect.TypeOf(InputTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatasetInput
// only implements ToObjectValue() and Type().
func (o DatasetInput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dataset": o.Dataset,
			"tags":    o.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DatasetInput) Type(ctx context.Context) attr.Type {
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
func (o *DatasetInput) GetDataset(ctx context.Context) (Dataset, bool) {
	var e Dataset
	if o.Dataset.IsNull() || o.Dataset.IsUnknown() {
		return e, false
	}
	var v []Dataset
	d := o.Dataset.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDataset sets the value of the Dataset field in DatasetInput.
func (o *DatasetInput) SetDataset(ctx context.Context, v Dataset) {
	vs := v.ToObjectValue(ctx)
	o.Dataset = vs
}

// GetTags returns the value of the Tags field in DatasetInput as
// a slice of InputTag values.
// If the field is unknown or null, the boolean return value is false.
func (o *DatasetInput) GetTags(ctx context.Context) ([]InputTag, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []InputTag
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in DatasetInput.
func (o *DatasetInput) SetTags(ctx context.Context, v []InputTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type DeleteCommentRequest struct {
	// Unique identifier of an activity
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCommentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteCommentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCommentRequest
// only implements ToObjectValue() and Type().
func (o DeleteCommentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCommentRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteCommentResponse struct {
}

func (newState *DeleteCommentResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteCommentResponse) {
}

func (newState *DeleteCommentResponse) SyncEffectiveFieldsDuringRead(existingState DeleteCommentResponse) {
}

func (c DeleteCommentResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCommentResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteCommentResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCommentResponse
// only implements ToObjectValue() and Type().
func (o DeleteCommentResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCommentResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteExperiment struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"experiment_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteExperiment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteExperiment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteExperiment
// only implements ToObjectValue() and Type().
func (o DeleteExperiment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": o.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteExperiment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
		},
	}
}

type DeleteExperimentResponse struct {
}

func (newState *DeleteExperimentResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteExperimentResponse) {
}

func (newState *DeleteExperimentResponse) SyncEffectiveFieldsDuringRead(existingState DeleteExperimentResponse) {
}

func (c DeleteExperimentResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteExperimentResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteExperimentResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteExperimentResponse
// only implements ToObjectValue() and Type().
func (o DeleteExperimentResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteExperimentResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteFeatureTagRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteFeatureTagRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteFeatureTagRequest
// only implements ToObjectValue() and Type().
func (o DeleteFeatureTagRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_name": o.FeatureName,
			"key":          o.Key,
			"table_name":   o.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteFeatureTagRequest) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteLoggedModelRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteLoggedModelRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteLoggedModelRequest
// only implements ToObjectValue() and Type().
func (o DeleteLoggedModelRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_id": o.ModelId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteLoggedModelRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_id": types.StringType,
		},
	}
}

type DeleteLoggedModelResponse struct {
}

func (newState *DeleteLoggedModelResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteLoggedModelResponse) {
}

func (newState *DeleteLoggedModelResponse) SyncEffectiveFieldsDuringRead(existingState DeleteLoggedModelResponse) {
}

func (c DeleteLoggedModelResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteLoggedModelResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteLoggedModelResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteLoggedModelResponse
// only implements ToObjectValue() and Type().
func (o DeleteLoggedModelResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteLoggedModelResponse) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteLoggedModelTagRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteLoggedModelTagRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteLoggedModelTagRequest
// only implements ToObjectValue() and Type().
func (o DeleteLoggedModelTagRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_id": o.ModelId,
			"tag_key":  o.TagKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteLoggedModelTagRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_id": types.StringType,
			"tag_key":  types.StringType,
		},
	}
}

type DeleteLoggedModelTagResponse struct {
}

func (newState *DeleteLoggedModelTagResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteLoggedModelTagResponse) {
}

func (newState *DeleteLoggedModelTagResponse) SyncEffectiveFieldsDuringRead(existingState DeleteLoggedModelTagResponse) {
}

func (c DeleteLoggedModelTagResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteLoggedModelTagResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteLoggedModelTagResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteLoggedModelTagResponse
// only implements ToObjectValue() and Type().
func (o DeleteLoggedModelTagResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteLoggedModelTagResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteModelRequest struct {
	// Registered model unique name identifier.
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteModelRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteModelRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelRequest
// only implements ToObjectValue() and Type().
func (o DeleteModelRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteModelRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteModelResponse struct {
}

func (newState *DeleteModelResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteModelResponse) {
}

func (newState *DeleteModelResponse) SyncEffectiveFieldsDuringRead(existingState DeleteModelResponse) {
}

func (c DeleteModelResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteModelResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteModelResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelResponse
// only implements ToObjectValue() and Type().
func (o DeleteModelResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteModelResponse) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteModelTagRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteModelTagRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelTagRequest
// only implements ToObjectValue() and Type().
func (o DeleteModelTagRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":  o.Key,
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteModelTagRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":  types.StringType,
			"name": types.StringType,
		},
	}
}

type DeleteModelTagResponse struct {
}

func (newState *DeleteModelTagResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteModelTagResponse) {
}

func (newState *DeleteModelTagResponse) SyncEffectiveFieldsDuringRead(existingState DeleteModelTagResponse) {
}

func (c DeleteModelTagResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteModelTagResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteModelTagResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelTagResponse
// only implements ToObjectValue() and Type().
func (o DeleteModelTagResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteModelTagResponse) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteModelVersionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteModelVersionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelVersionRequest
// only implements ToObjectValue() and Type().
func (o DeleteModelVersionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":    o.Name,
			"version": o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteModelVersionRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":    types.StringType,
			"version": types.StringType,
		},
	}
}

type DeleteModelVersionResponse struct {
}

func (newState *DeleteModelVersionResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteModelVersionResponse) {
}

func (newState *DeleteModelVersionResponse) SyncEffectiveFieldsDuringRead(existingState DeleteModelVersionResponse) {
}

func (c DeleteModelVersionResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteModelVersionResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteModelVersionResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelVersionResponse
// only implements ToObjectValue() and Type().
func (o DeleteModelVersionResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteModelVersionResponse) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteModelVersionTagRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteModelVersionTagRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelVersionTagRequest
// only implements ToObjectValue() and Type().
func (o DeleteModelVersionTagRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":     o.Key,
			"name":    o.Name,
			"version": o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteModelVersionTagRequest) Type(ctx context.Context) attr.Type {
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

func (newState *DeleteModelVersionTagResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteModelVersionTagResponse) {
}

func (newState *DeleteModelVersionTagResponse) SyncEffectiveFieldsDuringRead(existingState DeleteModelVersionTagResponse) {
}

func (c DeleteModelVersionTagResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteModelVersionTagResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteModelVersionTagResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteModelVersionTagResponse
// only implements ToObjectValue() and Type().
func (o DeleteModelVersionTagResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteModelVersionTagResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteOnlineStoreRequest struct {
	// Name of the online store to delete.
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteOnlineStoreRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteOnlineStoreRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteOnlineStoreRequest
// only implements ToObjectValue() and Type().
func (o DeleteOnlineStoreRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteOnlineStoreRequest) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteRun.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteRun) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRun
// only implements ToObjectValue() and Type().
func (o DeleteRun) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run_id": o.RunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteRun) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run_id": types.StringType,
		},
	}
}

type DeleteRunResponse struct {
}

func (newState *DeleteRunResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteRunResponse) {
}

func (newState *DeleteRunResponse) SyncEffectiveFieldsDuringRead(existingState DeleteRunResponse) {
}

func (c DeleteRunResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteRunResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteRunResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRunResponse
// only implements ToObjectValue() and Type().
func (o DeleteRunResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteRunResponse) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteRuns.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteRuns) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRuns
// only implements ToObjectValue() and Type().
func (o DeleteRuns) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id":        o.ExperimentId,
			"max_runs":             o.MaxRuns,
			"max_timestamp_millis": o.MaxTimestampMillis,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteRuns) Type(ctx context.Context) attr.Type {
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

func (newState *DeleteRunsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteRunsResponse) {
}

func (newState *DeleteRunsResponse) SyncEffectiveFieldsDuringRead(existingState DeleteRunsResponse) {
}

func (c DeleteRunsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteRunsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRunsResponse
// only implements ToObjectValue() and Type().
func (o DeleteRunsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"runs_deleted": o.RunsDeleted,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteRunsResponse) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteTag.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteTag) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteTag
// only implements ToObjectValue() and Type().
func (o DeleteTag) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":    o.Key,
			"run_id": o.RunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteTag) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":    types.StringType,
			"run_id": types.StringType,
		},
	}
}

type DeleteTagResponse struct {
}

func (newState *DeleteTagResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteTagResponse) {
}

func (newState *DeleteTagResponse) SyncEffectiveFieldsDuringRead(existingState DeleteTagResponse) {
}

func (c DeleteTagResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteTagResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteTagResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteTagResponse
// only implements ToObjectValue() and Type().
func (o DeleteTagResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteTagResponse) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteTransitionRequestRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteTransitionRequestRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteTransitionRequestRequest
// only implements ToObjectValue() and Type().
func (o DeleteTransitionRequestRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o DeleteTransitionRequestRequest) Type(ctx context.Context) attr.Type {
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

func (newState *DeleteTransitionRequestResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteTransitionRequestResponse) {
}

func (newState *DeleteTransitionRequestResponse) SyncEffectiveFieldsDuringRead(existingState DeleteTransitionRequestResponse) {
}

func (c DeleteTransitionRequestResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteTransitionRequestResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"activity": reflect.TypeOf(Activity{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteTransitionRequestResponse
// only implements ToObjectValue() and Type().
func (o DeleteTransitionRequestResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"activity": o.Activity,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteTransitionRequestResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"activity": Activity{}.Type(ctx),
		},
	}
}

// GetActivity returns the value of the Activity field in DeleteTransitionRequestResponse as
// a Activity value.
// If the field is unknown or null, the boolean return value is false.
func (o *DeleteTransitionRequestResponse) GetActivity(ctx context.Context) (Activity, bool) {
	var e Activity
	if o.Activity.IsNull() || o.Activity.IsUnknown() {
		return e, false
	}
	var v []Activity
	d := o.Activity.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetActivity sets the value of the Activity field in DeleteTransitionRequestResponse.
func (o *DeleteTransitionRequestResponse) SetActivity(ctx context.Context, v Activity) {
	vs := v.ToObjectValue(ctx)
	o.Activity = vs
}

type DeleteWebhookRequest struct {
	// Webhook ID required to delete a registry webhook.
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteWebhookRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteWebhookRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWebhookRequest
// only implements ToObjectValue() and Type().
func (o DeleteWebhookRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteWebhookRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteWebhookResponse struct {
}

func (newState *DeleteWebhookResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteWebhookResponse) {
}

func (newState *DeleteWebhookResponse) SyncEffectiveFieldsDuringRead(existingState DeleteWebhookResponse) {
}

func (c DeleteWebhookResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteWebhookResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteWebhookResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteWebhookResponse
// only implements ToObjectValue() and Type().
func (o DeleteWebhookResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteWebhookResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
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

func (newState *Experiment) SyncEffectiveFieldsDuringCreateOrUpdate(plan Experiment) {
}

func (newState *Experiment) SyncEffectiveFieldsDuringRead(existingState Experiment) {
}

func (c Experiment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Experiment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(ExperimentTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Experiment
// only implements ToObjectValue() and Type().
func (o Experiment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o Experiment) Type(ctx context.Context) attr.Type {
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
func (o *Experiment) GetTags(ctx context.Context) ([]ExperimentTag, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []ExperimentTag
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in Experiment.
func (o *Experiment) SetTags(ctx context.Context, v []ExperimentTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
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

func (newState *ExperimentAccessControlRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExperimentAccessControlRequest) {
}

func (newState *ExperimentAccessControlRequest) SyncEffectiveFieldsDuringRead(existingState ExperimentAccessControlRequest) {
}

func (c ExperimentAccessControlRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ExperimentAccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExperimentAccessControlRequest
// only implements ToObjectValue() and Type().
func (o ExperimentAccessControlRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ExperimentAccessControlRequest) Type(ctx context.Context) attr.Type {
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

func (newState *ExperimentAccessControlResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExperimentAccessControlResponse) {
}

func (newState *ExperimentAccessControlResponse) SyncEffectiveFieldsDuringRead(existingState ExperimentAccessControlResponse) {
}

func (c ExperimentAccessControlResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ExperimentAccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(ExperimentPermission{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExperimentAccessControlResponse
// only implements ToObjectValue() and Type().
func (o ExperimentAccessControlResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ExperimentAccessControlResponse) Type(ctx context.Context) attr.Type {
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
func (o *ExperimentAccessControlResponse) GetAllPermissions(ctx context.Context) ([]ExperimentPermission, bool) {
	if o.AllPermissions.IsNull() || o.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []ExperimentPermission
	d := o.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in ExperimentAccessControlResponse.
func (o *ExperimentAccessControlResponse) SetAllPermissions(ctx context.Context, v []ExperimentPermission) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllPermissions = types.ListValueMust(t, vs)
}

type ExperimentPermission struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (newState *ExperimentPermission) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExperimentPermission) {
}

func (newState *ExperimentPermission) SyncEffectiveFieldsDuringRead(existingState ExperimentPermission) {
}

func (c ExperimentPermission) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ExperimentPermission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExperimentPermission
// only implements ToObjectValue() and Type().
func (o ExperimentPermission) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             o.Inherited,
			"inherited_from_object": o.InheritedFromObject,
			"permission_level":      o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExperimentPermission) Type(ctx context.Context) attr.Type {
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
func (o *ExperimentPermission) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
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

// SetInheritedFromObject sets the value of the InheritedFromObject field in ExperimentPermission.
func (o *ExperimentPermission) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InheritedFromObject = types.ListValueMust(t, vs)
}

type ExperimentPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (newState *ExperimentPermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExperimentPermissions) {
}

func (newState *ExperimentPermissions) SyncEffectiveFieldsDuringRead(existingState ExperimentPermissions) {
}

func (c ExperimentPermissions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ExperimentPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(ExperimentAccessControlResponse{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExperimentPermissions
// only implements ToObjectValue() and Type().
func (o ExperimentPermissions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExperimentPermissions) Type(ctx context.Context) attr.Type {
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
func (o *ExperimentPermissions) GetAccessControlList(ctx context.Context) ([]ExperimentAccessControlResponse, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []ExperimentAccessControlResponse
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in ExperimentPermissions.
func (o *ExperimentPermissions) SetAccessControlList(ctx context.Context, v []ExperimentAccessControlResponse) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type ExperimentPermissionsDescription struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (newState *ExperimentPermissionsDescription) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExperimentPermissionsDescription) {
}

func (newState *ExperimentPermissionsDescription) SyncEffectiveFieldsDuringRead(existingState ExperimentPermissionsDescription) {
}

func (c ExperimentPermissionsDescription) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ExperimentPermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExperimentPermissionsDescription
// only implements ToObjectValue() and Type().
func (o ExperimentPermissionsDescription) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      o.Description,
			"permission_level": o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExperimentPermissionsDescription) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExperimentPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ExperimentPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(ExperimentAccessControlRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExperimentPermissionsRequest
// only implements ToObjectValue() and Type().
func (o ExperimentPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"experiment_id":       o.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExperimentPermissionsRequest) Type(ctx context.Context) attr.Type {
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
func (o *ExperimentPermissionsRequest) GetAccessControlList(ctx context.Context) ([]ExperimentAccessControlRequest, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []ExperimentAccessControlRequest
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in ExperimentPermissionsRequest.
func (o *ExperimentPermissionsRequest) SetAccessControlList(ctx context.Context, v []ExperimentAccessControlRequest) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

// A tag for an experiment.
type ExperimentTag struct {
	// The tag key.
	Key types.String `tfsdk:"key"`
	// The tag value.
	Value types.String `tfsdk:"value"`
}

func (newState *ExperimentTag) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExperimentTag) {
}

func (newState *ExperimentTag) SyncEffectiveFieldsDuringRead(existingState ExperimentTag) {
}

func (c ExperimentTag) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ExperimentTag) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExperimentTag
// only implements ToObjectValue() and Type().
func (o ExperimentTag) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExperimentTag) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

// Feature for model version.
type Feature struct {
	// Feature name
	FeatureName types.String `tfsdk:"feature_name"`
	// Feature table id
	FeatureTableId types.String `tfsdk:"feature_table_id"`
	// Feature table name
	FeatureTableName types.String `tfsdk:"feature_table_name"`
}

func (newState *Feature) SyncEffectiveFieldsDuringCreateOrUpdate(plan Feature) {
}

func (newState *Feature) SyncEffectiveFieldsDuringRead(existingState Feature) {
}

func (c Feature) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["feature_name"] = attrs["feature_name"].SetOptional()
	attrs["feature_table_id"] = attrs["feature_table_id"].SetOptional()
	attrs["feature_table_name"] = attrs["feature_table_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Feature.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Feature) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Feature
// only implements ToObjectValue() and Type().
func (o Feature) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_name":       o.FeatureName,
			"feature_table_id":   o.FeatureTableId,
			"feature_table_name": o.FeatureTableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Feature) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"feature_name":       types.StringType,
			"feature_table_id":   types.StringType,
			"feature_table_name": types.StringType,
		},
	}
}

type FeatureLineage struct {
	// List of feature specs that contain this feature.
	FeatureSpecs types.List `tfsdk:"feature_specs"`
	// List of Unity Catalog models that were trained on this feature.
	Models types.List `tfsdk:"models"`
	// List of online features that use this feature as source.
	OnlineFeatures types.List `tfsdk:"online_features"`
}

func (newState *FeatureLineage) SyncEffectiveFieldsDuringCreateOrUpdate(plan FeatureLineage) {
}

func (newState *FeatureLineage) SyncEffectiveFieldsDuringRead(existingState FeatureLineage) {
}

func (c FeatureLineage) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a FeatureLineage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"feature_specs":   reflect.TypeOf(FeatureLineageFeatureSpec{}),
		"models":          reflect.TypeOf(FeatureLineageModel{}),
		"online_features": reflect.TypeOf(FeatureLineageOnlineFeature{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FeatureLineage
// only implements ToObjectValue() and Type().
func (o FeatureLineage) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_specs":   o.FeatureSpecs,
			"models":          o.Models,
			"online_features": o.OnlineFeatures,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FeatureLineage) Type(ctx context.Context) attr.Type {
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
func (o *FeatureLineage) GetFeatureSpecs(ctx context.Context) ([]FeatureLineageFeatureSpec, bool) {
	if o.FeatureSpecs.IsNull() || o.FeatureSpecs.IsUnknown() {
		return nil, false
	}
	var v []FeatureLineageFeatureSpec
	d := o.FeatureSpecs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFeatureSpecs sets the value of the FeatureSpecs field in FeatureLineage.
func (o *FeatureLineage) SetFeatureSpecs(ctx context.Context, v []FeatureLineageFeatureSpec) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["feature_specs"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.FeatureSpecs = types.ListValueMust(t, vs)
}

// GetModels returns the value of the Models field in FeatureLineage as
// a slice of FeatureLineageModel values.
// If the field is unknown or null, the boolean return value is false.
func (o *FeatureLineage) GetModels(ctx context.Context) ([]FeatureLineageModel, bool) {
	if o.Models.IsNull() || o.Models.IsUnknown() {
		return nil, false
	}
	var v []FeatureLineageModel
	d := o.Models.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModels sets the value of the Models field in FeatureLineage.
func (o *FeatureLineage) SetModels(ctx context.Context, v []FeatureLineageModel) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["models"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Models = types.ListValueMust(t, vs)
}

// GetOnlineFeatures returns the value of the OnlineFeatures field in FeatureLineage as
// a slice of FeatureLineageOnlineFeature values.
// If the field is unknown or null, the boolean return value is false.
func (o *FeatureLineage) GetOnlineFeatures(ctx context.Context) ([]FeatureLineageOnlineFeature, bool) {
	if o.OnlineFeatures.IsNull() || o.OnlineFeatures.IsUnknown() {
		return nil, false
	}
	var v []FeatureLineageOnlineFeature
	d := o.OnlineFeatures.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOnlineFeatures sets the value of the OnlineFeatures field in FeatureLineage.
func (o *FeatureLineage) SetOnlineFeatures(ctx context.Context, v []FeatureLineageOnlineFeature) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["online_features"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OnlineFeatures = types.ListValueMust(t, vs)
}

type FeatureLineageFeatureSpec struct {
	// The full name of the feature spec in Unity Catalog.
	Name types.String `tfsdk:"name"`
}

func (newState *FeatureLineageFeatureSpec) SyncEffectiveFieldsDuringCreateOrUpdate(plan FeatureLineageFeatureSpec) {
}

func (newState *FeatureLineageFeatureSpec) SyncEffectiveFieldsDuringRead(existingState FeatureLineageFeatureSpec) {
}

func (c FeatureLineageFeatureSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a FeatureLineageFeatureSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FeatureLineageFeatureSpec
// only implements ToObjectValue() and Type().
func (o FeatureLineageFeatureSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FeatureLineageFeatureSpec) Type(ctx context.Context) attr.Type {
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

func (newState *FeatureLineageModel) SyncEffectiveFieldsDuringCreateOrUpdate(plan FeatureLineageModel) {
}

func (newState *FeatureLineageModel) SyncEffectiveFieldsDuringRead(existingState FeatureLineageModel) {
}

func (c FeatureLineageModel) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a FeatureLineageModel) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FeatureLineageModel
// only implements ToObjectValue() and Type().
func (o FeatureLineageModel) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":    o.Name,
			"version": o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FeatureLineageModel) Type(ctx context.Context) attr.Type {
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

func (newState *FeatureLineageOnlineFeature) SyncEffectiveFieldsDuringCreateOrUpdate(plan FeatureLineageOnlineFeature) {
}

func (newState *FeatureLineageOnlineFeature) SyncEffectiveFieldsDuringRead(existingState FeatureLineageOnlineFeature) {
}

func (c FeatureLineageOnlineFeature) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a FeatureLineageOnlineFeature) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FeatureLineageOnlineFeature
// only implements ToObjectValue() and Type().
func (o FeatureLineageOnlineFeature) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_name": o.FeatureName,
			"table_name":   o.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FeatureLineageOnlineFeature) Type(ctx context.Context) attr.Type {
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

func (newState *FeatureList) SyncEffectiveFieldsDuringCreateOrUpdate(plan FeatureList) {
}

func (newState *FeatureList) SyncEffectiveFieldsDuringRead(existingState FeatureList) {
}

func (c FeatureList) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a FeatureList) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"features": reflect.TypeOf(Feature{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FeatureList
// only implements ToObjectValue() and Type().
func (o FeatureList) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"features": o.Features,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FeatureList) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"features": basetypes.ListType{
				ElemType: Feature{}.Type(ctx),
			},
		},
	}
}

// GetFeatures returns the value of the Features field in FeatureList as
// a slice of Feature values.
// If the field is unknown or null, the boolean return value is false.
func (o *FeatureList) GetFeatures(ctx context.Context) ([]Feature, bool) {
	if o.Features.IsNull() || o.Features.IsUnknown() {
		return nil, false
	}
	var v []Feature
	d := o.Features.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFeatures sets the value of the Features field in FeatureList.
func (o *FeatureList) SetFeatures(ctx context.Context, v []Feature) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["features"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Features = types.ListValueMust(t, vs)
}

// Represents a tag on a feature in a feature table.
type FeatureTag struct {
	Key types.String `tfsdk:"key"`

	Value types.String `tfsdk:"value"`
}

func (newState *FeatureTag) SyncEffectiveFieldsDuringCreateOrUpdate(plan FeatureTag) {
}

func (newState *FeatureTag) SyncEffectiveFieldsDuringRead(existingState FeatureTag) {
}

func (c FeatureTag) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a FeatureTag) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FeatureTag
// only implements ToObjectValue() and Type().
func (o FeatureTag) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FeatureTag) Type(ctx context.Context) attr.Type {
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

func (newState *FileInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan FileInfo) {
}

func (newState *FileInfo) SyncEffectiveFieldsDuringRead(existingState FileInfo) {
}

func (c FileInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a FileInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FileInfo
// only implements ToObjectValue() and Type().
func (o FileInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"file_size": o.FileSize,
			"is_dir":    o.IsDir,
			"path":      o.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FileInfo) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FinalizeLoggedModelRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a FinalizeLoggedModelRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FinalizeLoggedModelRequest
// only implements ToObjectValue() and Type().
func (o FinalizeLoggedModelRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_id": o.ModelId,
			"status":   o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FinalizeLoggedModelRequest) Type(ctx context.Context) attr.Type {
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

func (newState *FinalizeLoggedModelResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan FinalizeLoggedModelResponse) {
}

func (newState *FinalizeLoggedModelResponse) SyncEffectiveFieldsDuringRead(existingState FinalizeLoggedModelResponse) {
}

func (c FinalizeLoggedModelResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a FinalizeLoggedModelResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model": reflect.TypeOf(LoggedModel{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FinalizeLoggedModelResponse
// only implements ToObjectValue() and Type().
func (o FinalizeLoggedModelResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model": o.Model,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FinalizeLoggedModelResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model": LoggedModel{}.Type(ctx),
		},
	}
}

// GetModel returns the value of the Model field in FinalizeLoggedModelResponse as
// a LoggedModel value.
// If the field is unknown or null, the boolean return value is false.
func (o *FinalizeLoggedModelResponse) GetModel(ctx context.Context) (LoggedModel, bool) {
	var e LoggedModel
	if o.Model.IsNull() || o.Model.IsUnknown() {
		return e, false
	}
	var v []LoggedModel
	d := o.Model.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetModel sets the value of the Model field in FinalizeLoggedModelResponse.
func (o *FinalizeLoggedModelResponse) SetModel(ctx context.Context, v LoggedModel) {
	vs := v.ToObjectValue(ctx)
	o.Model = vs
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

func (newState *ForecastingExperiment) SyncEffectiveFieldsDuringCreateOrUpdate(plan ForecastingExperiment) {
}

func (newState *ForecastingExperiment) SyncEffectiveFieldsDuringRead(existingState ForecastingExperiment) {
}

func (c ForecastingExperiment) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ForecastingExperiment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ForecastingExperiment
// only implements ToObjectValue() and Type().
func (o ForecastingExperiment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id":       o.ExperimentId,
			"experiment_page_url": o.ExperimentPageUrl,
			"state":               o.State,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ForecastingExperiment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id":       types.StringType,
			"experiment_page_url": types.StringType,
			"state":               types.StringType,
		},
	}
}

type GetByNameRequest struct {
	// Name of the associated experiment.
	ExperimentName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetByNameRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetByNameRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetByNameRequest
// only implements ToObjectValue() and Type().
func (o GetByNameRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_name": o.ExperimentName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetByNameRequest) Type(ctx context.Context) attr.Type {
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

func (newState *GetExperimentByNameResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetExperimentByNameResponse) {
}

func (newState *GetExperimentByNameResponse) SyncEffectiveFieldsDuringRead(existingState GetExperimentByNameResponse) {
}

func (c GetExperimentByNameResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetExperimentByNameResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"experiment": reflect.TypeOf(Experiment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExperimentByNameResponse
// only implements ToObjectValue() and Type().
func (o GetExperimentByNameResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment": o.Experiment,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetExperimentByNameResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment": Experiment{}.Type(ctx),
		},
	}
}

// GetExperiment returns the value of the Experiment field in GetExperimentByNameResponse as
// a Experiment value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetExperimentByNameResponse) GetExperiment(ctx context.Context) (Experiment, bool) {
	var e Experiment
	if o.Experiment.IsNull() || o.Experiment.IsUnknown() {
		return e, false
	}
	var v []Experiment
	d := o.Experiment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetExperiment sets the value of the Experiment field in GetExperimentByNameResponse.
func (o *GetExperimentByNameResponse) SetExperiment(ctx context.Context, v Experiment) {
	vs := v.ToObjectValue(ctx)
	o.Experiment = vs
}

type GetExperimentPermissionLevelsRequest struct {
	// The experiment for which to get or manage permissions.
	ExperimentId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetExperimentPermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetExperimentPermissionLevelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExperimentPermissionLevelsRequest
// only implements ToObjectValue() and Type().
func (o GetExperimentPermissionLevelsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": o.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetExperimentPermissionLevelsRequest) Type(ctx context.Context) attr.Type {
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

func (newState *GetExperimentPermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetExperimentPermissionLevelsResponse) {
}

func (newState *GetExperimentPermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetExperimentPermissionLevelsResponse) {
}

func (c GetExperimentPermissionLevelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetExperimentPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(ExperimentPermissionsDescription{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExperimentPermissionLevelsResponse
// only implements ToObjectValue() and Type().
func (o GetExperimentPermissionLevelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": o.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetExperimentPermissionLevelsResponse) Type(ctx context.Context) attr.Type {
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
func (o *GetExperimentPermissionLevelsResponse) GetPermissionLevels(ctx context.Context) ([]ExperimentPermissionsDescription, bool) {
	if o.PermissionLevels.IsNull() || o.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []ExperimentPermissionsDescription
	d := o.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetExperimentPermissionLevelsResponse.
func (o *GetExperimentPermissionLevelsResponse) SetPermissionLevels(ctx context.Context, v []ExperimentPermissionsDescription) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PermissionLevels = types.ListValueMust(t, vs)
}

type GetExperimentPermissionsRequest struct {
	// The experiment for which to get or manage permissions.
	ExperimentId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetExperimentPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetExperimentPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExperimentPermissionsRequest
// only implements ToObjectValue() and Type().
func (o GetExperimentPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": o.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetExperimentPermissionsRequest) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetExperimentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetExperimentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExperimentRequest
// only implements ToObjectValue() and Type().
func (o GetExperimentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": o.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetExperimentRequest) Type(ctx context.Context) attr.Type {
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

func (newState *GetExperimentResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetExperimentResponse) {
}

func (newState *GetExperimentResponse) SyncEffectiveFieldsDuringRead(existingState GetExperimentResponse) {
}

func (c GetExperimentResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetExperimentResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"experiment": reflect.TypeOf(Experiment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetExperimentResponse
// only implements ToObjectValue() and Type().
func (o GetExperimentResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment": o.Experiment,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetExperimentResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment": Experiment{}.Type(ctx),
		},
	}
}

// GetExperiment returns the value of the Experiment field in GetExperimentResponse as
// a Experiment value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetExperimentResponse) GetExperiment(ctx context.Context) (Experiment, bool) {
	var e Experiment
	if o.Experiment.IsNull() || o.Experiment.IsUnknown() {
		return e, false
	}
	var v []Experiment
	d := o.Experiment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetExperiment sets the value of the Experiment field in GetExperimentResponse.
func (o *GetExperimentResponse) SetExperiment(ctx context.Context, v Experiment) {
	vs := v.ToObjectValue(ctx)
	o.Experiment = vs
}

type GetFeatureLineageRequest struct {
	// The name of the feature.
	FeatureName types.String `tfsdk:"-"`
	// The full name of the feature table in Unity Catalog.
	TableName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetFeatureLineageRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetFeatureLineageRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetFeatureLineageRequest
// only implements ToObjectValue() and Type().
func (o GetFeatureLineageRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_name": o.FeatureName,
			"table_name":   o.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetFeatureLineageRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"feature_name": types.StringType,
			"table_name":   types.StringType,
		},
	}
}

type GetFeatureTagRequest struct {
	FeatureName types.String `tfsdk:"-"`

	Key types.String `tfsdk:"-"`

	TableName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetFeatureTagRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetFeatureTagRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetFeatureTagRequest
// only implements ToObjectValue() and Type().
func (o GetFeatureTagRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_name": o.FeatureName,
			"key":          o.Key,
			"table_name":   o.TableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetFeatureTagRequest) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetForecastingExperimentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetForecastingExperimentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetForecastingExperimentRequest
// only implements ToObjectValue() and Type().
func (o GetForecastingExperimentRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": o.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetForecastingExperimentRequest) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetHistoryRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetHistoryRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetHistoryRequest
// only implements ToObjectValue() and Type().
func (o GetHistoryRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o GetHistoryRequest) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetLatestVersionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetLatestVersionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"stages": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLatestVersionsRequest
// only implements ToObjectValue() and Type().
func (o GetLatestVersionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":   o.Name,
			"stages": o.Stages,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetLatestVersionsRequest) Type(ctx context.Context) attr.Type {
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
func (o *GetLatestVersionsRequest) GetStages(ctx context.Context) ([]types.String, bool) {
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

// SetStages sets the value of the Stages field in GetLatestVersionsRequest.
func (o *GetLatestVersionsRequest) SetStages(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["stages"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Stages = types.ListValueMust(t, vs)
}

type GetLatestVersionsResponse struct {
	// Latest version models for each requests stage. Only return models with
	// current `READY` status. If no `stages` provided, returns the latest
	// version for each stage, including `"None"`.
	ModelVersions types.List `tfsdk:"model_versions"`
}

func (newState *GetLatestVersionsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetLatestVersionsResponse) {
}

func (newState *GetLatestVersionsResponse) SyncEffectiveFieldsDuringRead(existingState GetLatestVersionsResponse) {
}

func (c GetLatestVersionsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetLatestVersionsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model_versions": reflect.TypeOf(ModelVersion{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLatestVersionsResponse
// only implements ToObjectValue() and Type().
func (o GetLatestVersionsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_versions": o.ModelVersions,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetLatestVersionsResponse) Type(ctx context.Context) attr.Type {
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
func (o *GetLatestVersionsResponse) GetModelVersions(ctx context.Context) ([]ModelVersion, bool) {
	if o.ModelVersions.IsNull() || o.ModelVersions.IsUnknown() {
		return nil, false
	}
	var v []ModelVersion
	d := o.ModelVersions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModelVersions sets the value of the ModelVersions field in GetLatestVersionsResponse.
func (o *GetLatestVersionsResponse) SetModelVersions(ctx context.Context, v []ModelVersion) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["model_versions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ModelVersions = types.ListValueMust(t, vs)
}

type GetLoggedModelRequest struct {
	// The ID of the logged model to retrieve.
	ModelId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetLoggedModelRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetLoggedModelRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLoggedModelRequest
// only implements ToObjectValue() and Type().
func (o GetLoggedModelRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_id": o.ModelId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetLoggedModelRequest) Type(ctx context.Context) attr.Type {
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

func (newState *GetLoggedModelResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetLoggedModelResponse) {
}

func (newState *GetLoggedModelResponse) SyncEffectiveFieldsDuringRead(existingState GetLoggedModelResponse) {
}

func (c GetLoggedModelResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetLoggedModelResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model": reflect.TypeOf(LoggedModel{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLoggedModelResponse
// only implements ToObjectValue() and Type().
func (o GetLoggedModelResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model": o.Model,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetLoggedModelResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model": LoggedModel{}.Type(ctx),
		},
	}
}

// GetModel returns the value of the Model field in GetLoggedModelResponse as
// a LoggedModel value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetLoggedModelResponse) GetModel(ctx context.Context) (LoggedModel, bool) {
	var e LoggedModel
	if o.Model.IsNull() || o.Model.IsUnknown() {
		return e, false
	}
	var v []LoggedModel
	d := o.Model.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetModel sets the value of the Model field in GetLoggedModelResponse.
func (o *GetLoggedModelResponse) SetModel(ctx context.Context, v LoggedModel) {
	vs := v.ToObjectValue(ctx)
	o.Model = vs
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

func (newState *GetMetricHistoryResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetMetricHistoryResponse) {
}

func (newState *GetMetricHistoryResponse) SyncEffectiveFieldsDuringRead(existingState GetMetricHistoryResponse) {
}

func (c GetMetricHistoryResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetMetricHistoryResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metrics": reflect.TypeOf(Metric{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetMetricHistoryResponse
// only implements ToObjectValue() and Type().
func (o GetMetricHistoryResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metrics":         o.Metrics,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetMetricHistoryResponse) Type(ctx context.Context) attr.Type {
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
func (o *GetMetricHistoryResponse) GetMetrics(ctx context.Context) ([]Metric, bool) {
	if o.Metrics.IsNull() || o.Metrics.IsUnknown() {
		return nil, false
	}
	var v []Metric
	d := o.Metrics.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMetrics sets the value of the Metrics field in GetMetricHistoryResponse.
func (o *GetMetricHistoryResponse) SetMetrics(ctx context.Context, v []Metric) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["metrics"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Metrics = types.ListValueMust(t, vs)
}

type GetModelRequest struct {
	// Registered model unique name identifier.
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetModelRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetModelRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetModelRequest
// only implements ToObjectValue() and Type().
func (o GetModelRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetModelRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetModelResponse struct {
	RegisteredModelDatabricks types.Object `tfsdk:"registered_model_databricks"`
}

func (newState *GetModelResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetModelResponse) {
}

func (newState *GetModelResponse) SyncEffectiveFieldsDuringRead(existingState GetModelResponse) {
}

func (c GetModelResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetModelResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"registered_model_databricks": reflect.TypeOf(ModelDatabricks{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetModelResponse
// only implements ToObjectValue() and Type().
func (o GetModelResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"registered_model_databricks": o.RegisteredModelDatabricks,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetModelResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"registered_model_databricks": ModelDatabricks{}.Type(ctx),
		},
	}
}

// GetRegisteredModelDatabricks returns the value of the RegisteredModelDatabricks field in GetModelResponse as
// a ModelDatabricks value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetModelResponse) GetRegisteredModelDatabricks(ctx context.Context) (ModelDatabricks, bool) {
	var e ModelDatabricks
	if o.RegisteredModelDatabricks.IsNull() || o.RegisteredModelDatabricks.IsUnknown() {
		return e, false
	}
	var v []ModelDatabricks
	d := o.RegisteredModelDatabricks.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRegisteredModelDatabricks sets the value of the RegisteredModelDatabricks field in GetModelResponse.
func (o *GetModelResponse) SetRegisteredModelDatabricks(ctx context.Context, v ModelDatabricks) {
	vs := v.ToObjectValue(ctx)
	o.RegisteredModelDatabricks = vs
}

type GetModelVersionDownloadUriRequest struct {
	// Name of the registered model
	Name types.String `tfsdk:"-"`
	// Model version number
	Version types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetModelVersionDownloadUriRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetModelVersionDownloadUriRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetModelVersionDownloadUriRequest
// only implements ToObjectValue() and Type().
func (o GetModelVersionDownloadUriRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":    o.Name,
			"version": o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetModelVersionDownloadUriRequest) Type(ctx context.Context) attr.Type {
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

func (newState *GetModelVersionDownloadUriResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetModelVersionDownloadUriResponse) {
}

func (newState *GetModelVersionDownloadUriResponse) SyncEffectiveFieldsDuringRead(existingState GetModelVersionDownloadUriResponse) {
}

func (c GetModelVersionDownloadUriResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetModelVersionDownloadUriResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetModelVersionDownloadUriResponse
// only implements ToObjectValue() and Type().
func (o GetModelVersionDownloadUriResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"artifact_uri": o.ArtifactUri,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetModelVersionDownloadUriResponse) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetModelVersionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetModelVersionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetModelVersionRequest
// only implements ToObjectValue() and Type().
func (o GetModelVersionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":    o.Name,
			"version": o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetModelVersionRequest) Type(ctx context.Context) attr.Type {
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

func (newState *GetModelVersionResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetModelVersionResponse) {
}

func (newState *GetModelVersionResponse) SyncEffectiveFieldsDuringRead(existingState GetModelVersionResponse) {
}

func (c GetModelVersionResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetModelVersionResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model_version": reflect.TypeOf(ModelVersion{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetModelVersionResponse
// only implements ToObjectValue() and Type().
func (o GetModelVersionResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_version": o.ModelVersion,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetModelVersionResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_version": ModelVersion{}.Type(ctx),
		},
	}
}

// GetModelVersion returns the value of the ModelVersion field in GetModelVersionResponse as
// a ModelVersion value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetModelVersionResponse) GetModelVersion(ctx context.Context) (ModelVersion, bool) {
	var e ModelVersion
	if o.ModelVersion.IsNull() || o.ModelVersion.IsUnknown() {
		return e, false
	}
	var v []ModelVersion
	d := o.ModelVersion.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetModelVersion sets the value of the ModelVersion field in GetModelVersionResponse.
func (o *GetModelVersionResponse) SetModelVersion(ctx context.Context, v ModelVersion) {
	vs := v.ToObjectValue(ctx)
	o.ModelVersion = vs
}

type GetOnlineStoreRequest struct {
	// Name of the online store to get.
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetOnlineStoreRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetOnlineStoreRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetOnlineStoreRequest
// only implements ToObjectValue() and Type().
func (o GetOnlineStoreRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetOnlineStoreRequest) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRegisteredModelPermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetRegisteredModelPermissionLevelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRegisteredModelPermissionLevelsRequest
// only implements ToObjectValue() and Type().
func (o GetRegisteredModelPermissionLevelsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"registered_model_id": o.RegisteredModelId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetRegisteredModelPermissionLevelsRequest) Type(ctx context.Context) attr.Type {
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

func (newState *GetRegisteredModelPermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRegisteredModelPermissionLevelsResponse) {
}

func (newState *GetRegisteredModelPermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetRegisteredModelPermissionLevelsResponse) {
}

func (c GetRegisteredModelPermissionLevelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetRegisteredModelPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(RegisteredModelPermissionsDescription{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRegisteredModelPermissionLevelsResponse
// only implements ToObjectValue() and Type().
func (o GetRegisteredModelPermissionLevelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": o.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetRegisteredModelPermissionLevelsResponse) Type(ctx context.Context) attr.Type {
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
func (o *GetRegisteredModelPermissionLevelsResponse) GetPermissionLevels(ctx context.Context) ([]RegisteredModelPermissionsDescription, bool) {
	if o.PermissionLevels.IsNull() || o.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []RegisteredModelPermissionsDescription
	d := o.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetRegisteredModelPermissionLevelsResponse.
func (o *GetRegisteredModelPermissionLevelsResponse) SetPermissionLevels(ctx context.Context, v []RegisteredModelPermissionsDescription) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PermissionLevels = types.ListValueMust(t, vs)
}

type GetRegisteredModelPermissionsRequest struct {
	// The registered model for which to get or manage permissions.
	RegisteredModelId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRegisteredModelPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetRegisteredModelPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRegisteredModelPermissionsRequest
// only implements ToObjectValue() and Type().
func (o GetRegisteredModelPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"registered_model_id": o.RegisteredModelId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetRegisteredModelPermissionsRequest) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRunRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetRunRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRunRequest
// only implements ToObjectValue() and Type().
func (o GetRunRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run_id":   o.RunId,
			"run_uuid": o.RunUuid,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetRunRequest) Type(ctx context.Context) attr.Type {
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

func (newState *GetRunResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRunResponse) {
}

func (newState *GetRunResponse) SyncEffectiveFieldsDuringRead(existingState GetRunResponse) {
}

func (c GetRunResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetRunResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"run": reflect.TypeOf(Run{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRunResponse
// only implements ToObjectValue() and Type().
func (o GetRunResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run": o.Run,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetRunResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run": Run{}.Type(ctx),
		},
	}
}

// GetRun returns the value of the Run field in GetRunResponse as
// a Run value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetRunResponse) GetRun(ctx context.Context) (Run, bool) {
	var e Run
	if o.Run.IsNull() || o.Run.IsUnknown() {
		return e, false
	}
	var v []Run
	d := o.Run.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRun sets the value of the Run field in GetRunResponse.
func (o *GetRunResponse) SetRun(ctx context.Context, v Run) {
	vs := v.ToObjectValue(ctx)
	o.Run = vs
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

func (newState *HttpUrlSpec) SyncEffectiveFieldsDuringCreateOrUpdate(plan HttpUrlSpec) {
}

func (newState *HttpUrlSpec) SyncEffectiveFieldsDuringRead(existingState HttpUrlSpec) {
}

func (c HttpUrlSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a HttpUrlSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, HttpUrlSpec
// only implements ToObjectValue() and Type().
func (o HttpUrlSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o HttpUrlSpec) Type(ctx context.Context) attr.Type {
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

func (newState *HttpUrlSpecWithoutSecret) SyncEffectiveFieldsDuringCreateOrUpdate(plan HttpUrlSpecWithoutSecret) {
}

func (newState *HttpUrlSpecWithoutSecret) SyncEffectiveFieldsDuringRead(existingState HttpUrlSpecWithoutSecret) {
}

func (c HttpUrlSpecWithoutSecret) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a HttpUrlSpecWithoutSecret) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, HttpUrlSpecWithoutSecret
// only implements ToObjectValue() and Type().
func (o HttpUrlSpecWithoutSecret) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enable_ssl_verification": o.EnableSslVerification,
			"url":                     o.Url,
		})
}

// Type implements basetypes.ObjectValuable.
func (o HttpUrlSpecWithoutSecret) Type(ctx context.Context) attr.Type {
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

func (newState *InputTag) SyncEffectiveFieldsDuringCreateOrUpdate(plan InputTag) {
}

func (newState *InputTag) SyncEffectiveFieldsDuringRead(existingState InputTag) {
}

func (c InputTag) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a InputTag) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, InputTag
// only implements ToObjectValue() and Type().
func (o InputTag) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o InputTag) Type(ctx context.Context) attr.Type {
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

func (newState *JobSpec) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobSpec) {
}

func (newState *JobSpec) SyncEffectiveFieldsDuringRead(existingState JobSpec) {
}

func (c JobSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a JobSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, JobSpec
// only implements ToObjectValue() and Type().
func (o JobSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_token":  o.AccessToken,
			"job_id":        o.JobId,
			"workspace_url": o.WorkspaceUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (o JobSpec) Type(ctx context.Context) attr.Type {
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

func (newState *JobSpecWithoutSecret) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobSpecWithoutSecret) {
}

func (newState *JobSpecWithoutSecret) SyncEffectiveFieldsDuringRead(existingState JobSpecWithoutSecret) {
}

func (c JobSpecWithoutSecret) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a JobSpecWithoutSecret) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, JobSpecWithoutSecret
// only implements ToObjectValue() and Type().
func (o JobSpecWithoutSecret) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"job_id":        o.JobId,
			"workspace_url": o.WorkspaceUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (o JobSpecWithoutSecret) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"job_id":        types.StringType,
			"workspace_url": types.StringType,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListArtifactsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListArtifactsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListArtifactsRequest
// only implements ToObjectValue() and Type().
func (o ListArtifactsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListArtifactsRequest) Type(ctx context.Context) attr.Type {
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

func (newState *ListArtifactsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListArtifactsResponse) {
}

func (newState *ListArtifactsResponse) SyncEffectiveFieldsDuringRead(existingState ListArtifactsResponse) {
}

func (c ListArtifactsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListArtifactsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"files": reflect.TypeOf(FileInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListArtifactsResponse
// only implements ToObjectValue() and Type().
func (o ListArtifactsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"files":           o.Files,
			"next_page_token": o.NextPageToken,
			"root_uri":        o.RootUri,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListArtifactsResponse) Type(ctx context.Context) attr.Type {
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
func (o *ListArtifactsResponse) GetFiles(ctx context.Context) ([]FileInfo, bool) {
	if o.Files.IsNull() || o.Files.IsUnknown() {
		return nil, false
	}
	var v []FileInfo
	d := o.Files.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFiles sets the value of the Files field in ListArtifactsResponse.
func (o *ListArtifactsResponse) SetFiles(ctx context.Context, v []FileInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["files"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Files = types.ListValueMust(t, vs)
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListExperimentsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListExperimentsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExperimentsRequest
// only implements ToObjectValue() and Type().
func (o ListExperimentsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results": o.MaxResults,
			"page_token":  o.PageToken,
			"view_type":   o.ViewType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListExperimentsRequest) Type(ctx context.Context) attr.Type {
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

func (newState *ListExperimentsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListExperimentsResponse) {
}

func (newState *ListExperimentsResponse) SyncEffectiveFieldsDuringRead(existingState ListExperimentsResponse) {
}

func (c ListExperimentsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListExperimentsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"experiments": reflect.TypeOf(Experiment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListExperimentsResponse
// only implements ToObjectValue() and Type().
func (o ListExperimentsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiments":     o.Experiments,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListExperimentsResponse) Type(ctx context.Context) attr.Type {
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
func (o *ListExperimentsResponse) GetExperiments(ctx context.Context) ([]Experiment, bool) {
	if o.Experiments.IsNull() || o.Experiments.IsUnknown() {
		return nil, false
	}
	var v []Experiment
	d := o.Experiments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExperiments sets the value of the Experiments field in ListExperimentsResponse.
func (o *ListExperimentsResponse) SetExperiments(ctx context.Context, v []Experiment) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["experiments"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Experiments = types.ListValueMust(t, vs)
}

type ListFeatureTagsRequest struct {
	FeatureName types.String `tfsdk:"-"`
	// The maximum number of results to return.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page based on a previous query.
	PageToken types.String `tfsdk:"-"`

	TableName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListFeatureTagsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListFeatureTagsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFeatureTagsRequest
// only implements ToObjectValue() and Type().
func (o ListFeatureTagsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListFeatureTagsRequest) Type(ctx context.Context) attr.Type {
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

func (newState *ListFeatureTagsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListFeatureTagsResponse) {
}

func (newState *ListFeatureTagsResponse) SyncEffectiveFieldsDuringRead(existingState ListFeatureTagsResponse) {
}

func (c ListFeatureTagsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListFeatureTagsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"feature_tags": reflect.TypeOf(FeatureTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFeatureTagsResponse
// only implements ToObjectValue() and Type().
func (o ListFeatureTagsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"feature_tags":    o.FeatureTags,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListFeatureTagsResponse) Type(ctx context.Context) attr.Type {
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
func (o *ListFeatureTagsResponse) GetFeatureTags(ctx context.Context) ([]FeatureTag, bool) {
	if o.FeatureTags.IsNull() || o.FeatureTags.IsUnknown() {
		return nil, false
	}
	var v []FeatureTag
	d := o.FeatureTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFeatureTags sets the value of the FeatureTags field in ListFeatureTagsResponse.
func (o *ListFeatureTagsResponse) SetFeatureTags(ctx context.Context, v []FeatureTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["feature_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.FeatureTags = types.ListValueMust(t, vs)
}

type ListModelsRequest struct {
	// Maximum number of registered models desired. Max threshold is 1000.
	MaxResults types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page based on a previous query.
	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListModelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListModelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListModelsRequest
// only implements ToObjectValue() and Type().
func (o ListModelsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"max_results": o.MaxResults,
			"page_token":  o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListModelsRequest) Type(ctx context.Context) attr.Type {
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

func (newState *ListModelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListModelsResponse) {
}

func (newState *ListModelsResponse) SyncEffectiveFieldsDuringRead(existingState ListModelsResponse) {
}

func (c ListModelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListModelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"registered_models": reflect.TypeOf(Model{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListModelsResponse
// only implements ToObjectValue() and Type().
func (o ListModelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":   o.NextPageToken,
			"registered_models": o.RegisteredModels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListModelsResponse) Type(ctx context.Context) attr.Type {
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
func (o *ListModelsResponse) GetRegisteredModels(ctx context.Context) ([]Model, bool) {
	if o.RegisteredModels.IsNull() || o.RegisteredModels.IsUnknown() {
		return nil, false
	}
	var v []Model
	d := o.RegisteredModels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRegisteredModels sets the value of the RegisteredModels field in ListModelsResponse.
func (o *ListModelsResponse) SetRegisteredModels(ctx context.Context, v []Model) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["registered_models"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RegisteredModels = types.ListValueMust(t, vs)
}

type ListOnlineStoresRequest struct {
	// The maximum number of results to return. Defaults to 100 if not
	// specified.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page based on a previous query.
	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListOnlineStoresRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListOnlineStoresRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListOnlineStoresRequest
// only implements ToObjectValue() and Type().
func (o ListOnlineStoresRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListOnlineStoresRequest) Type(ctx context.Context) attr.Type {
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

func (newState *ListOnlineStoresResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListOnlineStoresResponse) {
}

func (newState *ListOnlineStoresResponse) SyncEffectiveFieldsDuringRead(existingState ListOnlineStoresResponse) {
}

func (c ListOnlineStoresResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListOnlineStoresResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"online_stores": reflect.TypeOf(OnlineStore{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListOnlineStoresResponse
// only implements ToObjectValue() and Type().
func (o ListOnlineStoresResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"online_stores":   o.OnlineStores,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListOnlineStoresResponse) Type(ctx context.Context) attr.Type {
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
func (o *ListOnlineStoresResponse) GetOnlineStores(ctx context.Context) ([]OnlineStore, bool) {
	if o.OnlineStores.IsNull() || o.OnlineStores.IsUnknown() {
		return nil, false
	}
	var v []OnlineStore
	d := o.OnlineStores.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOnlineStores sets the value of the OnlineStores field in ListOnlineStoresResponse.
func (o *ListOnlineStoresResponse) SetOnlineStores(ctx context.Context, v []OnlineStore) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["online_stores"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OnlineStores = types.ListValueMust(t, vs)
}

type ListRegistryWebhooks struct {
	// Token that can be used to retrieve the next page of artifact results
	NextPageToken types.String `tfsdk:"next_page_token"`
	// Array of registry webhooks.
	Webhooks types.List `tfsdk:"webhooks"`
}

func (newState *ListRegistryWebhooks) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListRegistryWebhooks) {
}

func (newState *ListRegistryWebhooks) SyncEffectiveFieldsDuringRead(existingState ListRegistryWebhooks) {
}

func (c ListRegistryWebhooks) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListRegistryWebhooks) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"webhooks": reflect.TypeOf(RegistryWebhook{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListRegistryWebhooks
// only implements ToObjectValue() and Type().
func (o ListRegistryWebhooks) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"webhooks":        o.Webhooks,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListRegistryWebhooks) Type(ctx context.Context) attr.Type {
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
func (o *ListRegistryWebhooks) GetWebhooks(ctx context.Context) ([]RegistryWebhook, bool) {
	if o.Webhooks.IsNull() || o.Webhooks.IsUnknown() {
		return nil, false
	}
	var v []RegistryWebhook
	d := o.Webhooks.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWebhooks sets the value of the Webhooks field in ListRegistryWebhooks.
func (o *ListRegistryWebhooks) SetWebhooks(ctx context.Context, v []RegistryWebhook) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["webhooks"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Webhooks = types.ListValueMust(t, vs)
}

type ListTransitionRequestsRequest struct {
	// Name of the registered model.
	Name types.String `tfsdk:"-"`
	// Version of the model.
	Version types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListTransitionRequestsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListTransitionRequestsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTransitionRequestsRequest
// only implements ToObjectValue() and Type().
func (o ListTransitionRequestsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":    o.Name,
			"version": o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListTransitionRequestsRequest) Type(ctx context.Context) attr.Type {
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

func (newState *ListTransitionRequestsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListTransitionRequestsResponse) {
}

func (newState *ListTransitionRequestsResponse) SyncEffectiveFieldsDuringRead(existingState ListTransitionRequestsResponse) {
}

func (c ListTransitionRequestsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListTransitionRequestsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"requests": reflect.TypeOf(Activity{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTransitionRequestsResponse
// only implements ToObjectValue() and Type().
func (o ListTransitionRequestsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"requests": o.Requests,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListTransitionRequestsResponse) Type(ctx context.Context) attr.Type {
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
func (o *ListTransitionRequestsResponse) GetRequests(ctx context.Context) ([]Activity, bool) {
	if o.Requests.IsNull() || o.Requests.IsUnknown() {
		return nil, false
	}
	var v []Activity
	d := o.Requests.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRequests sets the value of the Requests field in ListTransitionRequestsResponse.
func (o *ListTransitionRequestsResponse) SetRequests(ctx context.Context, v []Activity) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["requests"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Requests = types.ListValueMust(t, vs)
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListWebhooksRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListWebhooksRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"events": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListWebhooksRequest
// only implements ToObjectValue() and Type().
func (o ListWebhooksRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListWebhooksRequest) Type(ctx context.Context) attr.Type {
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
func (o *ListWebhooksRequest) GetEvents(ctx context.Context) ([]types.String, bool) {
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

// SetEvents sets the value of the Events field in ListWebhooksRequest.
func (o *ListWebhooksRequest) SetEvents(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["events"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Events = types.ListValueMust(t, vs)
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogBatch.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a LogBatch) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metrics": reflect.TypeOf(Metric{}),
		"params":  reflect.TypeOf(Param{}),
		"tags":    reflect.TypeOf(RunTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogBatch
// only implements ToObjectValue() and Type().
func (o LogBatch) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o LogBatch) Type(ctx context.Context) attr.Type {
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
func (o *LogBatch) GetMetrics(ctx context.Context) ([]Metric, bool) {
	if o.Metrics.IsNull() || o.Metrics.IsUnknown() {
		return nil, false
	}
	var v []Metric
	d := o.Metrics.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMetrics sets the value of the Metrics field in LogBatch.
func (o *LogBatch) SetMetrics(ctx context.Context, v []Metric) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["metrics"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Metrics = types.ListValueMust(t, vs)
}

// GetParams returns the value of the Params field in LogBatch as
// a slice of Param values.
// If the field is unknown or null, the boolean return value is false.
func (o *LogBatch) GetParams(ctx context.Context) ([]Param, bool) {
	if o.Params.IsNull() || o.Params.IsUnknown() {
		return nil, false
	}
	var v []Param
	d := o.Params.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParams sets the value of the Params field in LogBatch.
func (o *LogBatch) SetParams(ctx context.Context, v []Param) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Params = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in LogBatch as
// a slice of RunTag values.
// If the field is unknown or null, the boolean return value is false.
func (o *LogBatch) GetTags(ctx context.Context) ([]RunTag, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []RunTag
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in LogBatch.
func (o *LogBatch) SetTags(ctx context.Context, v []RunTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type LogBatchResponse struct {
}

func (newState *LogBatchResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogBatchResponse) {
}

func (newState *LogBatchResponse) SyncEffectiveFieldsDuringRead(existingState LogBatchResponse) {
}

func (c LogBatchResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogBatchResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a LogBatchResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogBatchResponse
// only implements ToObjectValue() and Type().
func (o LogBatchResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o LogBatchResponse) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogInputs.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a LogInputs) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"datasets": reflect.TypeOf(DatasetInput{}),
		"models":   reflect.TypeOf(ModelInput{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogInputs
// only implements ToObjectValue() and Type().
func (o LogInputs) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"datasets": o.Datasets,
			"models":   o.Models,
			"run_id":   o.RunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LogInputs) Type(ctx context.Context) attr.Type {
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
func (o *LogInputs) GetDatasets(ctx context.Context) ([]DatasetInput, bool) {
	if o.Datasets.IsNull() || o.Datasets.IsUnknown() {
		return nil, false
	}
	var v []DatasetInput
	d := o.Datasets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatasets sets the value of the Datasets field in LogInputs.
func (o *LogInputs) SetDatasets(ctx context.Context, v []DatasetInput) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["datasets"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Datasets = types.ListValueMust(t, vs)
}

// GetModels returns the value of the Models field in LogInputs as
// a slice of ModelInput values.
// If the field is unknown or null, the boolean return value is false.
func (o *LogInputs) GetModels(ctx context.Context) ([]ModelInput, bool) {
	if o.Models.IsNull() || o.Models.IsUnknown() {
		return nil, false
	}
	var v []ModelInput
	d := o.Models.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModels sets the value of the Models field in LogInputs.
func (o *LogInputs) SetModels(ctx context.Context, v []ModelInput) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["models"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Models = types.ListValueMust(t, vs)
}

type LogInputsResponse struct {
}

func (newState *LogInputsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogInputsResponse) {
}

func (newState *LogInputsResponse) SyncEffectiveFieldsDuringRead(existingState LogInputsResponse) {
}

func (c LogInputsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogInputsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a LogInputsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogInputsResponse
// only implements ToObjectValue() and Type().
func (o LogInputsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o LogInputsResponse) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogLoggedModelParamsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a LogLoggedModelParamsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"params": reflect.TypeOf(LoggedModelParameter{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogLoggedModelParamsRequest
// only implements ToObjectValue() and Type().
func (o LogLoggedModelParamsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_id": o.ModelId,
			"params":   o.Params,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LogLoggedModelParamsRequest) Type(ctx context.Context) attr.Type {
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
func (o *LogLoggedModelParamsRequest) GetParams(ctx context.Context) ([]LoggedModelParameter, bool) {
	if o.Params.IsNull() || o.Params.IsUnknown() {
		return nil, false
	}
	var v []LoggedModelParameter
	d := o.Params.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParams sets the value of the Params field in LogLoggedModelParamsRequest.
func (o *LogLoggedModelParamsRequest) SetParams(ctx context.Context, v []LoggedModelParameter) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Params = types.ListValueMust(t, vs)
}

type LogLoggedModelParamsRequestResponse struct {
}

func (newState *LogLoggedModelParamsRequestResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogLoggedModelParamsRequestResponse) {
}

func (newState *LogLoggedModelParamsRequestResponse) SyncEffectiveFieldsDuringRead(existingState LogLoggedModelParamsRequestResponse) {
}

func (c LogLoggedModelParamsRequestResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogLoggedModelParamsRequestResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a LogLoggedModelParamsRequestResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogLoggedModelParamsRequestResponse
// only implements ToObjectValue() and Type().
func (o LogLoggedModelParamsRequestResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o LogLoggedModelParamsRequestResponse) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogMetric.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a LogMetric) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogMetric
// only implements ToObjectValue() and Type().
func (o LogMetric) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o LogMetric) Type(ctx context.Context) attr.Type {
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

func (newState *LogMetricResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogMetricResponse) {
}

func (newState *LogMetricResponse) SyncEffectiveFieldsDuringRead(existingState LogMetricResponse) {
}

func (c LogMetricResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogMetricResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a LogMetricResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogMetricResponse
// only implements ToObjectValue() and Type().
func (o LogMetricResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o LogMetricResponse) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogModel.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a LogModel) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogModel
// only implements ToObjectValue() and Type().
func (o LogModel) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_json": o.ModelJson,
			"run_id":     o.RunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LogModel) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_json": types.StringType,
			"run_id":     types.StringType,
		},
	}
}

type LogModelResponse struct {
}

func (newState *LogModelResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogModelResponse) {
}

func (newState *LogModelResponse) SyncEffectiveFieldsDuringRead(existingState LogModelResponse) {
}

func (c LogModelResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogModelResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a LogModelResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogModelResponse
// only implements ToObjectValue() and Type().
func (o LogModelResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o LogModelResponse) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogOutputsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a LogOutputsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"models": reflect.TypeOf(ModelOutput{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogOutputsRequest
// only implements ToObjectValue() and Type().
func (o LogOutputsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"models": o.Models,
			"run_id": o.RunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LogOutputsRequest) Type(ctx context.Context) attr.Type {
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
func (o *LogOutputsRequest) GetModels(ctx context.Context) ([]ModelOutput, bool) {
	if o.Models.IsNull() || o.Models.IsUnknown() {
		return nil, false
	}
	var v []ModelOutput
	d := o.Models.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModels sets the value of the Models field in LogOutputsRequest.
func (o *LogOutputsRequest) SetModels(ctx context.Context, v []ModelOutput) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["models"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Models = types.ListValueMust(t, vs)
}

type LogOutputsResponse struct {
}

func (newState *LogOutputsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogOutputsResponse) {
}

func (newState *LogOutputsResponse) SyncEffectiveFieldsDuringRead(existingState LogOutputsResponse) {
}

func (c LogOutputsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogOutputsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a LogOutputsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogOutputsResponse
// only implements ToObjectValue() and Type().
func (o LogOutputsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o LogOutputsResponse) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogParam.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a LogParam) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogParam
// only implements ToObjectValue() and Type().
func (o LogParam) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o LogParam) Type(ctx context.Context) attr.Type {
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

func (newState *LogParamResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogParamResponse) {
}

func (newState *LogParamResponse) SyncEffectiveFieldsDuringRead(existingState LogParamResponse) {
}

func (c LogParamResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogParamResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a LogParamResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogParamResponse
// only implements ToObjectValue() and Type().
func (o LogParamResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o LogParamResponse) Type(ctx context.Context) attr.Type {
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

func (newState *LoggedModel) SyncEffectiveFieldsDuringCreateOrUpdate(plan LoggedModel) {
}

func (newState *LoggedModel) SyncEffectiveFieldsDuringRead(existingState LoggedModel) {
}

func (c LoggedModel) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a LoggedModel) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data": reflect.TypeOf(LoggedModelData{}),
		"info": reflect.TypeOf(LoggedModelInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LoggedModel
// only implements ToObjectValue() and Type().
func (o LoggedModel) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data": o.Data,
			"info": o.Info,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LoggedModel) Type(ctx context.Context) attr.Type {
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
func (o *LoggedModel) GetData(ctx context.Context) (LoggedModelData, bool) {
	var e LoggedModelData
	if o.Data.IsNull() || o.Data.IsUnknown() {
		return e, false
	}
	var v []LoggedModelData
	d := o.Data.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetData sets the value of the Data field in LoggedModel.
func (o *LoggedModel) SetData(ctx context.Context, v LoggedModelData) {
	vs := v.ToObjectValue(ctx)
	o.Data = vs
}

// GetInfo returns the value of the Info field in LoggedModel as
// a LoggedModelInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *LoggedModel) GetInfo(ctx context.Context) (LoggedModelInfo, bool) {
	var e LoggedModelInfo
	if o.Info.IsNull() || o.Info.IsUnknown() {
		return e, false
	}
	var v []LoggedModelInfo
	d := o.Info.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInfo sets the value of the Info field in LoggedModel.
func (o *LoggedModel) SetInfo(ctx context.Context, v LoggedModelInfo) {
	vs := v.ToObjectValue(ctx)
	o.Info = vs
}

// A LoggedModelData message includes logged model params and linked metrics.
type LoggedModelData struct {
	// Performance metrics linked to the model.
	Metrics types.List `tfsdk:"metrics"`
	// Immutable string key-value pairs of the model.
	Params types.List `tfsdk:"params"`
}

func (newState *LoggedModelData) SyncEffectiveFieldsDuringCreateOrUpdate(plan LoggedModelData) {
}

func (newState *LoggedModelData) SyncEffectiveFieldsDuringRead(existingState LoggedModelData) {
}

func (c LoggedModelData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a LoggedModelData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metrics": reflect.TypeOf(Metric{}),
		"params":  reflect.TypeOf(LoggedModelParameter{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LoggedModelData
// only implements ToObjectValue() and Type().
func (o LoggedModelData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metrics": o.Metrics,
			"params":  o.Params,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LoggedModelData) Type(ctx context.Context) attr.Type {
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
func (o *LoggedModelData) GetMetrics(ctx context.Context) ([]Metric, bool) {
	if o.Metrics.IsNull() || o.Metrics.IsUnknown() {
		return nil, false
	}
	var v []Metric
	d := o.Metrics.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMetrics sets the value of the Metrics field in LoggedModelData.
func (o *LoggedModelData) SetMetrics(ctx context.Context, v []Metric) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["metrics"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Metrics = types.ListValueMust(t, vs)
}

// GetParams returns the value of the Params field in LoggedModelData as
// a slice of LoggedModelParameter values.
// If the field is unknown or null, the boolean return value is false.
func (o *LoggedModelData) GetParams(ctx context.Context) ([]LoggedModelParameter, bool) {
	if o.Params.IsNull() || o.Params.IsUnknown() {
		return nil, false
	}
	var v []LoggedModelParameter
	d := o.Params.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParams sets the value of the Params field in LoggedModelData.
func (o *LoggedModelData) SetParams(ctx context.Context, v []LoggedModelParameter) {
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

func (newState *LoggedModelInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan LoggedModelInfo) {
}

func (newState *LoggedModelInfo) SyncEffectiveFieldsDuringRead(existingState LoggedModelInfo) {
}

func (c LoggedModelInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a LoggedModelInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(LoggedModelTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LoggedModelInfo
// only implements ToObjectValue() and Type().
func (o LoggedModelInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o LoggedModelInfo) Type(ctx context.Context) attr.Type {
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
func (o *LoggedModelInfo) GetTags(ctx context.Context) ([]LoggedModelTag, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []LoggedModelTag
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in LoggedModelInfo.
func (o *LoggedModelInfo) SetTags(ctx context.Context, v []LoggedModelTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

// Parameter associated with a LoggedModel.
type LoggedModelParameter struct {
	// The key identifying this param.
	Key types.String `tfsdk:"key"`
	// The value of this param.
	Value types.String `tfsdk:"value"`
}

func (newState *LoggedModelParameter) SyncEffectiveFieldsDuringCreateOrUpdate(plan LoggedModelParameter) {
}

func (newState *LoggedModelParameter) SyncEffectiveFieldsDuringRead(existingState LoggedModelParameter) {
}

func (c LoggedModelParameter) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a LoggedModelParameter) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LoggedModelParameter
// only implements ToObjectValue() and Type().
func (o LoggedModelParameter) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LoggedModelParameter) Type(ctx context.Context) attr.Type {
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

func (newState *LoggedModelTag) SyncEffectiveFieldsDuringCreateOrUpdate(plan LoggedModelTag) {
}

func (newState *LoggedModelTag) SyncEffectiveFieldsDuringRead(existingState LoggedModelTag) {
}

func (c LoggedModelTag) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a LoggedModelTag) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LoggedModelTag
// only implements ToObjectValue() and Type().
func (o LoggedModelTag) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LoggedModelTag) Type(ctx context.Context) attr.Type {
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

func (newState *Metric) SyncEffectiveFieldsDuringCreateOrUpdate(plan Metric) {
}

func (newState *Metric) SyncEffectiveFieldsDuringRead(existingState Metric) {
}

func (c Metric) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Metric) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Metric
// only implements ToObjectValue() and Type().
func (o Metric) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o Metric) Type(ctx context.Context) attr.Type {
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

func (newState *Model) SyncEffectiveFieldsDuringCreateOrUpdate(plan Model) {
}

func (newState *Model) SyncEffectiveFieldsDuringRead(existingState Model) {
}

func (c Model) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Model) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"latest_versions": reflect.TypeOf(ModelVersion{}),
		"tags":            reflect.TypeOf(ModelTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Model
// only implements ToObjectValue() and Type().
func (o Model) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o Model) Type(ctx context.Context) attr.Type {
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
func (o *Model) GetLatestVersions(ctx context.Context) ([]ModelVersion, bool) {
	if o.LatestVersions.IsNull() || o.LatestVersions.IsUnknown() {
		return nil, false
	}
	var v []ModelVersion
	d := o.LatestVersions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLatestVersions sets the value of the LatestVersions field in Model.
func (o *Model) SetLatestVersions(ctx context.Context, v []ModelVersion) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["latest_versions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.LatestVersions = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in Model as
// a slice of ModelTag values.
// If the field is unknown or null, the boolean return value is false.
func (o *Model) GetTags(ctx context.Context) ([]ModelTag, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []ModelTag
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in Model.
func (o *Model) SetTags(ctx context.Context, v []ModelTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
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

func (newState *ModelDatabricks) SyncEffectiveFieldsDuringCreateOrUpdate(plan ModelDatabricks) {
}

func (newState *ModelDatabricks) SyncEffectiveFieldsDuringRead(existingState ModelDatabricks) {
}

func (c ModelDatabricks) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ModelDatabricks) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"latest_versions": reflect.TypeOf(ModelVersion{}),
		"tags":            reflect.TypeOf(ModelTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelDatabricks
// only implements ToObjectValue() and Type().
func (o ModelDatabricks) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ModelDatabricks) Type(ctx context.Context) attr.Type {
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
func (o *ModelDatabricks) GetLatestVersions(ctx context.Context) ([]ModelVersion, bool) {
	if o.LatestVersions.IsNull() || o.LatestVersions.IsUnknown() {
		return nil, false
	}
	var v []ModelVersion
	d := o.LatestVersions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLatestVersions sets the value of the LatestVersions field in ModelDatabricks.
func (o *ModelDatabricks) SetLatestVersions(ctx context.Context, v []ModelVersion) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["latest_versions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.LatestVersions = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in ModelDatabricks as
// a slice of ModelTag values.
// If the field is unknown or null, the boolean return value is false.
func (o *ModelDatabricks) GetTags(ctx context.Context) ([]ModelTag, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []ModelTag
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in ModelDatabricks.
func (o *ModelDatabricks) SetTags(ctx context.Context, v []ModelTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

// Represents a LoggedModel or Registered Model Version input to a Run.
type ModelInput struct {
	// The unique identifier of the model.
	ModelId types.String `tfsdk:"model_id"`
}

func (newState *ModelInput) SyncEffectiveFieldsDuringCreateOrUpdate(plan ModelInput) {
}

func (newState *ModelInput) SyncEffectiveFieldsDuringRead(existingState ModelInput) {
}

func (c ModelInput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ModelInput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelInput
// only implements ToObjectValue() and Type().
func (o ModelInput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_id": o.ModelId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ModelInput) Type(ctx context.Context) attr.Type {
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

func (newState *ModelOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan ModelOutput) {
}

func (newState *ModelOutput) SyncEffectiveFieldsDuringRead(existingState ModelOutput) {
}

func (c ModelOutput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ModelOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelOutput
// only implements ToObjectValue() and Type().
func (o ModelOutput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_id": o.ModelId,
			"step":     o.Step,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ModelOutput) Type(ctx context.Context) attr.Type {
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

func (newState *ModelTag) SyncEffectiveFieldsDuringCreateOrUpdate(plan ModelTag) {
}

func (newState *ModelTag) SyncEffectiveFieldsDuringRead(existingState ModelTag) {
}

func (c ModelTag) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ModelTag) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelTag
// only implements ToObjectValue() and Type().
func (o ModelTag) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ModelTag) Type(ctx context.Context) attr.Type {
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

func (newState *ModelVersion) SyncEffectiveFieldsDuringCreateOrUpdate(plan ModelVersion) {
}

func (newState *ModelVersion) SyncEffectiveFieldsDuringRead(existingState ModelVersion) {
}

func (c ModelVersion) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ModelVersion) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(ModelVersionTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelVersion
// only implements ToObjectValue() and Type().
func (o ModelVersion) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ModelVersion) Type(ctx context.Context) attr.Type {
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
func (o *ModelVersion) GetTags(ctx context.Context) ([]ModelVersionTag, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []ModelVersionTag
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in ModelVersion.
func (o *ModelVersion) SetTags(ctx context.Context, v []ModelVersionTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
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

func (newState *ModelVersionDatabricks) SyncEffectiveFieldsDuringCreateOrUpdate(plan ModelVersionDatabricks) {
}

func (newState *ModelVersionDatabricks) SyncEffectiveFieldsDuringRead(existingState ModelVersionDatabricks) {
}

func (c ModelVersionDatabricks) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ModelVersionDatabricks) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"feature_list":  reflect.TypeOf(FeatureList{}),
		"open_requests": reflect.TypeOf(Activity{}),
		"tags":          reflect.TypeOf(ModelVersionTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelVersionDatabricks
// only implements ToObjectValue() and Type().
func (o ModelVersionDatabricks) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ModelVersionDatabricks) Type(ctx context.Context) attr.Type {
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
func (o *ModelVersionDatabricks) GetFeatureList(ctx context.Context) (FeatureList, bool) {
	var e FeatureList
	if o.FeatureList.IsNull() || o.FeatureList.IsUnknown() {
		return e, false
	}
	var v []FeatureList
	d := o.FeatureList.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFeatureList sets the value of the FeatureList field in ModelVersionDatabricks.
func (o *ModelVersionDatabricks) SetFeatureList(ctx context.Context, v FeatureList) {
	vs := v.ToObjectValue(ctx)
	o.FeatureList = vs
}

// GetOpenRequests returns the value of the OpenRequests field in ModelVersionDatabricks as
// a slice of Activity values.
// If the field is unknown or null, the boolean return value is false.
func (o *ModelVersionDatabricks) GetOpenRequests(ctx context.Context) ([]Activity, bool) {
	if o.OpenRequests.IsNull() || o.OpenRequests.IsUnknown() {
		return nil, false
	}
	var v []Activity
	d := o.OpenRequests.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOpenRequests sets the value of the OpenRequests field in ModelVersionDatabricks.
func (o *ModelVersionDatabricks) SetOpenRequests(ctx context.Context, v []Activity) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["open_requests"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OpenRequests = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in ModelVersionDatabricks as
// a slice of ModelVersionTag values.
// If the field is unknown or null, the boolean return value is false.
func (o *ModelVersionDatabricks) GetTags(ctx context.Context) ([]ModelVersionTag, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []ModelVersionTag
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in ModelVersionDatabricks.
func (o *ModelVersionDatabricks) SetTags(ctx context.Context, v []ModelVersionTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type ModelVersionTag struct {
	// The tag key.
	Key types.String `tfsdk:"key"`
	// The tag value.
	Value types.String `tfsdk:"value"`
}

func (newState *ModelVersionTag) SyncEffectiveFieldsDuringCreateOrUpdate(plan ModelVersionTag) {
}

func (newState *ModelVersionTag) SyncEffectiveFieldsDuringRead(existingState ModelVersionTag) {
}

func (c ModelVersionTag) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ModelVersionTag) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelVersionTag
// only implements ToObjectValue() and Type().
func (o ModelVersionTag) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ModelVersionTag) Type(ctx context.Context) attr.Type {
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

func (newState *OnlineStore) SyncEffectiveFieldsDuringCreateOrUpdate(plan OnlineStore) {
}

func (newState *OnlineStore) SyncEffectiveFieldsDuringRead(existingState OnlineStore) {
}

func (c OnlineStore) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a OnlineStore) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, OnlineStore
// only implements ToObjectValue() and Type().
func (o OnlineStore) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o OnlineStore) Type(ctx context.Context) attr.Type {
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

func (newState *Param) SyncEffectiveFieldsDuringCreateOrUpdate(plan Param) {
}

func (newState *Param) SyncEffectiveFieldsDuringRead(existingState Param) {
}

func (c Param) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Param) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Param
// only implements ToObjectValue() and Type().
func (o Param) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Param) Type(ctx context.Context) attr.Type {
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
	// source table. Defaults to TRIGGERED if not specified. All publish modes
	// require the source table to have Change Data Feed (CDF) enabled.
	PublishMode types.String `tfsdk:"publish_mode"`
}

func (newState *PublishSpec) SyncEffectiveFieldsDuringCreateOrUpdate(plan PublishSpec) {
}

func (newState *PublishSpec) SyncEffectiveFieldsDuringRead(existingState PublishSpec) {
}

func (c PublishSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["online_store"] = attrs["online_store"].SetRequired()
	attrs["online_table_name"] = attrs["online_table_name"].SetRequired()
	attrs["publish_mode"] = attrs["publish_mode"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PublishSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PublishSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PublishSpec
// only implements ToObjectValue() and Type().
func (o PublishSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"online_store":      o.OnlineStore,
			"online_table_name": o.OnlineTableName,
			"publish_mode":      o.PublishMode,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PublishSpec) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PublishTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PublishTableRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"publish_spec": reflect.TypeOf(PublishSpec{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PublishTableRequest
// only implements ToObjectValue() and Type().
func (o PublishTableRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"publish_spec":      o.PublishSpec,
			"source_table_name": o.SourceTableName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PublishTableRequest) Type(ctx context.Context) attr.Type {
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
func (o *PublishTableRequest) GetPublishSpec(ctx context.Context) (PublishSpec, bool) {
	var e PublishSpec
	if o.PublishSpec.IsNull() || o.PublishSpec.IsUnknown() {
		return e, false
	}
	var v []PublishSpec
	d := o.PublishSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPublishSpec sets the value of the PublishSpec field in PublishTableRequest.
func (o *PublishTableRequest) SetPublishSpec(ctx context.Context, v PublishSpec) {
	vs := v.ToObjectValue(ctx)
	o.PublishSpec = vs
}

type PublishTableResponse struct {
	// The full three-part (catalog, schema, table) name of the online table.
	OnlineTableName types.String `tfsdk:"online_table_name"`
	// The ID of the pipeline that syncs the online table with the source table.
	PipelineId types.String `tfsdk:"pipeline_id"`
}

func (newState *PublishTableResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan PublishTableResponse) {
}

func (newState *PublishTableResponse) SyncEffectiveFieldsDuringRead(existingState PublishTableResponse) {
}

func (c PublishTableResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PublishTableResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PublishTableResponse
// only implements ToObjectValue() and Type().
func (o PublishTableResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"online_table_name": o.OnlineTableName,
			"pipeline_id":       o.PipelineId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PublishTableResponse) Type(ctx context.Context) attr.Type {
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

func (newState *RegisteredModelAccessControlRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegisteredModelAccessControlRequest) {
}

func (newState *RegisteredModelAccessControlRequest) SyncEffectiveFieldsDuringRead(existingState RegisteredModelAccessControlRequest) {
}

func (c RegisteredModelAccessControlRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RegisteredModelAccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegisteredModelAccessControlRequest
// only implements ToObjectValue() and Type().
func (o RegisteredModelAccessControlRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o RegisteredModelAccessControlRequest) Type(ctx context.Context) attr.Type {
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

func (newState *RegisteredModelAccessControlResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegisteredModelAccessControlResponse) {
}

func (newState *RegisteredModelAccessControlResponse) SyncEffectiveFieldsDuringRead(existingState RegisteredModelAccessControlResponse) {
}

func (c RegisteredModelAccessControlResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RegisteredModelAccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(RegisteredModelPermission{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegisteredModelAccessControlResponse
// only implements ToObjectValue() and Type().
func (o RegisteredModelAccessControlResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o RegisteredModelAccessControlResponse) Type(ctx context.Context) attr.Type {
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
func (o *RegisteredModelAccessControlResponse) GetAllPermissions(ctx context.Context) ([]RegisteredModelPermission, bool) {
	if o.AllPermissions.IsNull() || o.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []RegisteredModelPermission
	d := o.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in RegisteredModelAccessControlResponse.
func (o *RegisteredModelAccessControlResponse) SetAllPermissions(ctx context.Context, v []RegisteredModelPermission) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllPermissions = types.ListValueMust(t, vs)
}

type RegisteredModelPermission struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (newState *RegisteredModelPermission) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegisteredModelPermission) {
}

func (newState *RegisteredModelPermission) SyncEffectiveFieldsDuringRead(existingState RegisteredModelPermission) {
}

func (c RegisteredModelPermission) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RegisteredModelPermission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegisteredModelPermission
// only implements ToObjectValue() and Type().
func (o RegisteredModelPermission) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             o.Inherited,
			"inherited_from_object": o.InheritedFromObject,
			"permission_level":      o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RegisteredModelPermission) Type(ctx context.Context) attr.Type {
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
func (o *RegisteredModelPermission) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
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

// SetInheritedFromObject sets the value of the InheritedFromObject field in RegisteredModelPermission.
func (o *RegisteredModelPermission) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InheritedFromObject = types.ListValueMust(t, vs)
}

type RegisteredModelPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (newState *RegisteredModelPermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegisteredModelPermissions) {
}

func (newState *RegisteredModelPermissions) SyncEffectiveFieldsDuringRead(existingState RegisteredModelPermissions) {
}

func (c RegisteredModelPermissions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RegisteredModelPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(RegisteredModelAccessControlResponse{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegisteredModelPermissions
// only implements ToObjectValue() and Type().
func (o RegisteredModelPermissions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RegisteredModelPermissions) Type(ctx context.Context) attr.Type {
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
func (o *RegisteredModelPermissions) GetAccessControlList(ctx context.Context) ([]RegisteredModelAccessControlResponse, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []RegisteredModelAccessControlResponse
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in RegisteredModelPermissions.
func (o *RegisteredModelPermissions) SetAccessControlList(ctx context.Context, v []RegisteredModelAccessControlResponse) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type RegisteredModelPermissionsDescription struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (newState *RegisteredModelPermissionsDescription) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegisteredModelPermissionsDescription) {
}

func (newState *RegisteredModelPermissionsDescription) SyncEffectiveFieldsDuringRead(existingState RegisteredModelPermissionsDescription) {
}

func (c RegisteredModelPermissionsDescription) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RegisteredModelPermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegisteredModelPermissionsDescription
// only implements ToObjectValue() and Type().
func (o RegisteredModelPermissionsDescription) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      o.Description,
			"permission_level": o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RegisteredModelPermissionsDescription) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RegisteredModelPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RegisteredModelPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(RegisteredModelAccessControlRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegisteredModelPermissionsRequest
// only implements ToObjectValue() and Type().
func (o RegisteredModelPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"registered_model_id": o.RegisteredModelId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RegisteredModelPermissionsRequest) Type(ctx context.Context) attr.Type {
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
func (o *RegisteredModelPermissionsRequest) GetAccessControlList(ctx context.Context) ([]RegisteredModelAccessControlRequest, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []RegisteredModelAccessControlRequest
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in RegisteredModelPermissionsRequest.
func (o *RegisteredModelPermissionsRequest) SetAccessControlList(ctx context.Context, v []RegisteredModelAccessControlRequest) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
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

func (newState *RegistryWebhook) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegistryWebhook) {
}

func (newState *RegistryWebhook) SyncEffectiveFieldsDuringRead(existingState RegistryWebhook) {
}

func (c RegistryWebhook) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RegistryWebhook) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"events":        reflect.TypeOf(types.String{}),
		"http_url_spec": reflect.TypeOf(HttpUrlSpecWithoutSecret{}),
		"job_spec":      reflect.TypeOf(JobSpecWithoutSecret{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RegistryWebhook
// only implements ToObjectValue() and Type().
func (o RegistryWebhook) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o RegistryWebhook) Type(ctx context.Context) attr.Type {
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
func (o *RegistryWebhook) GetEvents(ctx context.Context) ([]types.String, bool) {
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

// SetEvents sets the value of the Events field in RegistryWebhook.
func (o *RegistryWebhook) SetEvents(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["events"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Events = types.ListValueMust(t, vs)
}

// GetHttpUrlSpec returns the value of the HttpUrlSpec field in RegistryWebhook as
// a HttpUrlSpecWithoutSecret value.
// If the field is unknown or null, the boolean return value is false.
func (o *RegistryWebhook) GetHttpUrlSpec(ctx context.Context) (HttpUrlSpecWithoutSecret, bool) {
	var e HttpUrlSpecWithoutSecret
	if o.HttpUrlSpec.IsNull() || o.HttpUrlSpec.IsUnknown() {
		return e, false
	}
	var v []HttpUrlSpecWithoutSecret
	d := o.HttpUrlSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetHttpUrlSpec sets the value of the HttpUrlSpec field in RegistryWebhook.
func (o *RegistryWebhook) SetHttpUrlSpec(ctx context.Context, v HttpUrlSpecWithoutSecret) {
	vs := v.ToObjectValue(ctx)
	o.HttpUrlSpec = vs
}

// GetJobSpec returns the value of the JobSpec field in RegistryWebhook as
// a JobSpecWithoutSecret value.
// If the field is unknown or null, the boolean return value is false.
func (o *RegistryWebhook) GetJobSpec(ctx context.Context) (JobSpecWithoutSecret, bool) {
	var e JobSpecWithoutSecret
	if o.JobSpec.IsNull() || o.JobSpec.IsUnknown() {
		return e, false
	}
	var v []JobSpecWithoutSecret
	d := o.JobSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetJobSpec sets the value of the JobSpec field in RegistryWebhook.
func (o *RegistryWebhook) SetJobSpec(ctx context.Context, v JobSpecWithoutSecret) {
	vs := v.ToObjectValue(ctx)
	o.JobSpec = vs
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RejectTransitionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RejectTransitionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RejectTransitionRequest
// only implements ToObjectValue() and Type().
func (o RejectTransitionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o RejectTransitionRequest) Type(ctx context.Context) attr.Type {
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

func (newState *RejectTransitionRequestResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RejectTransitionRequestResponse) {
}

func (newState *RejectTransitionRequestResponse) SyncEffectiveFieldsDuringRead(existingState RejectTransitionRequestResponse) {
}

func (c RejectTransitionRequestResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RejectTransitionRequestResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"activity": reflect.TypeOf(Activity{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RejectTransitionRequestResponse
// only implements ToObjectValue() and Type().
func (o RejectTransitionRequestResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"activity": o.Activity,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RejectTransitionRequestResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"activity": Activity{}.Type(ctx),
		},
	}
}

// GetActivity returns the value of the Activity field in RejectTransitionRequestResponse as
// a Activity value.
// If the field is unknown or null, the boolean return value is false.
func (o *RejectTransitionRequestResponse) GetActivity(ctx context.Context) (Activity, bool) {
	var e Activity
	if o.Activity.IsNull() || o.Activity.IsUnknown() {
		return e, false
	}
	var v []Activity
	d := o.Activity.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetActivity sets the value of the Activity field in RejectTransitionRequestResponse.
func (o *RejectTransitionRequestResponse) SetActivity(ctx context.Context, v Activity) {
	vs := v.ToObjectValue(ctx)
	o.Activity = vs
}

type RenameModelRequest struct {
	// Registered model unique name identifier.
	Name types.String `tfsdk:"name"`
	// If provided, updates the name for this `registered_model`.
	NewName types.String `tfsdk:"new_name"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RenameModelRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RenameModelRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RenameModelRequest
// only implements ToObjectValue() and Type().
func (o RenameModelRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":     o.Name,
			"new_name": o.NewName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RenameModelRequest) Type(ctx context.Context) attr.Type {
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

func (newState *RenameModelResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RenameModelResponse) {
}

func (newState *RenameModelResponse) SyncEffectiveFieldsDuringRead(existingState RenameModelResponse) {
}

func (c RenameModelResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RenameModelResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"registered_model": reflect.TypeOf(Model{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RenameModelResponse
// only implements ToObjectValue() and Type().
func (o RenameModelResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"registered_model": o.RegisteredModel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RenameModelResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"registered_model": Model{}.Type(ctx),
		},
	}
}

// GetRegisteredModel returns the value of the RegisteredModel field in RenameModelResponse as
// a Model value.
// If the field is unknown or null, the boolean return value is false.
func (o *RenameModelResponse) GetRegisteredModel(ctx context.Context) (Model, bool) {
	var e Model
	if o.RegisteredModel.IsNull() || o.RegisteredModel.IsUnknown() {
		return e, false
	}
	var v []Model
	d := o.RegisteredModel.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRegisteredModel sets the value of the RegisteredModel field in RenameModelResponse.
func (o *RenameModelResponse) SetRegisteredModel(ctx context.Context, v Model) {
	vs := v.ToObjectValue(ctx)
	o.RegisteredModel = vs
}

type RestoreExperiment struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"experiment_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestoreExperiment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RestoreExperiment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestoreExperiment
// only implements ToObjectValue() and Type().
func (o RestoreExperiment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": o.ExperimentId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RestoreExperiment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
		},
	}
}

type RestoreExperimentResponse struct {
}

func (newState *RestoreExperimentResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RestoreExperimentResponse) {
}

func (newState *RestoreExperimentResponse) SyncEffectiveFieldsDuringRead(existingState RestoreExperimentResponse) {
}

func (c RestoreExperimentResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestoreExperimentResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RestoreExperimentResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestoreExperimentResponse
// only implements ToObjectValue() and Type().
func (o RestoreExperimentResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o RestoreExperimentResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type RestoreRun struct {
	// ID of the run to restore.
	RunId types.String `tfsdk:"run_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestoreRun.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RestoreRun) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestoreRun
// only implements ToObjectValue() and Type().
func (o RestoreRun) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run_id": o.RunId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RestoreRun) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run_id": types.StringType,
		},
	}
}

type RestoreRunResponse struct {
}

func (newState *RestoreRunResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RestoreRunResponse) {
}

func (newState *RestoreRunResponse) SyncEffectiveFieldsDuringRead(existingState RestoreRunResponse) {
}

func (c RestoreRunResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestoreRunResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RestoreRunResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestoreRunResponse
// only implements ToObjectValue() and Type().
func (o RestoreRunResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o RestoreRunResponse) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestoreRuns.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RestoreRuns) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestoreRuns
// only implements ToObjectValue() and Type().
func (o RestoreRuns) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id":        o.ExperimentId,
			"max_runs":             o.MaxRuns,
			"min_timestamp_millis": o.MinTimestampMillis,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RestoreRuns) Type(ctx context.Context) attr.Type {
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

func (newState *RestoreRunsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RestoreRunsResponse) {
}

func (newState *RestoreRunsResponse) SyncEffectiveFieldsDuringRead(existingState RestoreRunsResponse) {
}

func (c RestoreRunsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RestoreRunsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestoreRunsResponse
// only implements ToObjectValue() and Type().
func (o RestoreRunsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"runs_restored": o.RunsRestored,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RestoreRunsResponse) Type(ctx context.Context) attr.Type {
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

func (newState *Run) SyncEffectiveFieldsDuringCreateOrUpdate(plan Run) {
}

func (newState *Run) SyncEffectiveFieldsDuringRead(existingState Run) {
}

func (c Run) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Run) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data":   reflect.TypeOf(RunData{}),
		"info":   reflect.TypeOf(RunInfo{}),
		"inputs": reflect.TypeOf(RunInputs{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Run
// only implements ToObjectValue() and Type().
func (o Run) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data":   o.Data,
			"info":   o.Info,
			"inputs": o.Inputs,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Run) Type(ctx context.Context) attr.Type {
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
func (o *Run) GetData(ctx context.Context) (RunData, bool) {
	var e RunData
	if o.Data.IsNull() || o.Data.IsUnknown() {
		return e, false
	}
	var v []RunData
	d := o.Data.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetData sets the value of the Data field in Run.
func (o *Run) SetData(ctx context.Context, v RunData) {
	vs := v.ToObjectValue(ctx)
	o.Data = vs
}

// GetInfo returns the value of the Info field in Run as
// a RunInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *Run) GetInfo(ctx context.Context) (RunInfo, bool) {
	var e RunInfo
	if o.Info.IsNull() || o.Info.IsUnknown() {
		return e, false
	}
	var v []RunInfo
	d := o.Info.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInfo sets the value of the Info field in Run.
func (o *Run) SetInfo(ctx context.Context, v RunInfo) {
	vs := v.ToObjectValue(ctx)
	o.Info = vs
}

// GetInputs returns the value of the Inputs field in Run as
// a RunInputs value.
// If the field is unknown or null, the boolean return value is false.
func (o *Run) GetInputs(ctx context.Context) (RunInputs, bool) {
	var e RunInputs
	if o.Inputs.IsNull() || o.Inputs.IsUnknown() {
		return e, false
	}
	var v []RunInputs
	d := o.Inputs.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInputs sets the value of the Inputs field in Run.
func (o *Run) SetInputs(ctx context.Context, v RunInputs) {
	vs := v.ToObjectValue(ctx)
	o.Inputs = vs
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

func (newState *RunData) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunData) {
}

func (newState *RunData) SyncEffectiveFieldsDuringRead(existingState RunData) {
}

func (c RunData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RunData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"metrics": reflect.TypeOf(Metric{}),
		"params":  reflect.TypeOf(Param{}),
		"tags":    reflect.TypeOf(RunTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RunData
// only implements ToObjectValue() and Type().
func (o RunData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"metrics": o.Metrics,
			"params":  o.Params,
			"tags":    o.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RunData) Type(ctx context.Context) attr.Type {
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
func (o *RunData) GetMetrics(ctx context.Context) ([]Metric, bool) {
	if o.Metrics.IsNull() || o.Metrics.IsUnknown() {
		return nil, false
	}
	var v []Metric
	d := o.Metrics.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMetrics sets the value of the Metrics field in RunData.
func (o *RunData) SetMetrics(ctx context.Context, v []Metric) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["metrics"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Metrics = types.ListValueMust(t, vs)
}

// GetParams returns the value of the Params field in RunData as
// a slice of Param values.
// If the field is unknown or null, the boolean return value is false.
func (o *RunData) GetParams(ctx context.Context) ([]Param, bool) {
	if o.Params.IsNull() || o.Params.IsUnknown() {
		return nil, false
	}
	var v []Param
	d := o.Params.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParams sets the value of the Params field in RunData.
func (o *RunData) SetParams(ctx context.Context, v []Param) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Params = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in RunData as
// a slice of RunTag values.
// If the field is unknown or null, the boolean return value is false.
func (o *RunData) GetTags(ctx context.Context) ([]RunTag, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []RunTag
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in RunData.
func (o *RunData) SetTags(ctx context.Context, v []RunTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
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

func (newState *RunInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunInfo) {
}

func (newState *RunInfo) SyncEffectiveFieldsDuringRead(existingState RunInfo) {
}

func (c RunInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RunInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RunInfo
// only implements ToObjectValue() and Type().
func (o RunInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o RunInfo) Type(ctx context.Context) attr.Type {
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

func (newState *RunInputs) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunInputs) {
}

func (newState *RunInputs) SyncEffectiveFieldsDuringRead(existingState RunInputs) {
}

func (c RunInputs) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RunInputs) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dataset_inputs": reflect.TypeOf(DatasetInput{}),
		"model_inputs":   reflect.TypeOf(ModelInput{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RunInputs
// only implements ToObjectValue() and Type().
func (o RunInputs) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dataset_inputs": o.DatasetInputs,
			"model_inputs":   o.ModelInputs,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RunInputs) Type(ctx context.Context) attr.Type {
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
func (o *RunInputs) GetDatasetInputs(ctx context.Context) ([]DatasetInput, bool) {
	if o.DatasetInputs.IsNull() || o.DatasetInputs.IsUnknown() {
		return nil, false
	}
	var v []DatasetInput
	d := o.DatasetInputs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatasetInputs sets the value of the DatasetInputs field in RunInputs.
func (o *RunInputs) SetDatasetInputs(ctx context.Context, v []DatasetInput) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dataset_inputs"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DatasetInputs = types.ListValueMust(t, vs)
}

// GetModelInputs returns the value of the ModelInputs field in RunInputs as
// a slice of ModelInput values.
// If the field is unknown or null, the boolean return value is false.
func (o *RunInputs) GetModelInputs(ctx context.Context) ([]ModelInput, bool) {
	if o.ModelInputs.IsNull() || o.ModelInputs.IsUnknown() {
		return nil, false
	}
	var v []ModelInput
	d := o.ModelInputs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModelInputs sets the value of the ModelInputs field in RunInputs.
func (o *RunInputs) SetModelInputs(ctx context.Context, v []ModelInput) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["model_inputs"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ModelInputs = types.ListValueMust(t, vs)
}

// Tag for a run.
type RunTag struct {
	// The tag key.
	Key types.String `tfsdk:"key"`
	// The tag value.
	Value types.String `tfsdk:"value"`
}

func (newState *RunTag) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunTag) {
}

func (newState *RunTag) SyncEffectiveFieldsDuringRead(existingState RunTag) {
}

func (c RunTag) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RunTag) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RunTag
// only implements ToObjectValue() and Type().
func (o RunTag) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RunTag) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SearchExperiments.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SearchExperiments) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"order_by": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchExperiments
// only implements ToObjectValue() and Type().
func (o SearchExperiments) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o SearchExperiments) Type(ctx context.Context) attr.Type {
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
func (o *SearchExperiments) GetOrderBy(ctx context.Context) ([]types.String, bool) {
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

// SetOrderBy sets the value of the OrderBy field in SearchExperiments.
func (o *SearchExperiments) SetOrderBy(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["order_by"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OrderBy = types.ListValueMust(t, vs)
}

type SearchExperimentsResponse struct {
	// Experiments that match the search criteria
	Experiments types.List `tfsdk:"experiments"`
	// Token that can be used to retrieve the next page of experiments. An empty
	// token means that no more experiments are available for retrieval.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (newState *SearchExperimentsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan SearchExperimentsResponse) {
}

func (newState *SearchExperimentsResponse) SyncEffectiveFieldsDuringRead(existingState SearchExperimentsResponse) {
}

func (c SearchExperimentsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SearchExperimentsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"experiments": reflect.TypeOf(Experiment{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchExperimentsResponse
// only implements ToObjectValue() and Type().
func (o SearchExperimentsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiments":     o.Experiments,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SearchExperimentsResponse) Type(ctx context.Context) attr.Type {
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
func (o *SearchExperimentsResponse) GetExperiments(ctx context.Context) ([]Experiment, bool) {
	if o.Experiments.IsNull() || o.Experiments.IsUnknown() {
		return nil, false
	}
	var v []Experiment
	d := o.Experiments.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExperiments sets the value of the Experiments field in SearchExperimentsResponse.
func (o *SearchExperimentsResponse) SetExperiments(ctx context.Context, v []Experiment) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["experiments"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Experiments = types.ListValueMust(t, vs)
}

type SearchLoggedModelsDataset struct {
	// The digest of the dataset.
	DatasetDigest types.String `tfsdk:"dataset_digest"`
	// The name of the dataset.
	DatasetName types.String `tfsdk:"dataset_name"`
}

func (newState *SearchLoggedModelsDataset) SyncEffectiveFieldsDuringCreateOrUpdate(plan SearchLoggedModelsDataset) {
}

func (newState *SearchLoggedModelsDataset) SyncEffectiveFieldsDuringRead(existingState SearchLoggedModelsDataset) {
}

func (c SearchLoggedModelsDataset) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SearchLoggedModelsDataset) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchLoggedModelsDataset
// only implements ToObjectValue() and Type().
func (o SearchLoggedModelsDataset) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dataset_digest": o.DatasetDigest,
			"dataset_name":   o.DatasetName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SearchLoggedModelsDataset) Type(ctx context.Context) attr.Type {
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

func (newState *SearchLoggedModelsOrderBy) SyncEffectiveFieldsDuringCreateOrUpdate(plan SearchLoggedModelsOrderBy) {
}

func (newState *SearchLoggedModelsOrderBy) SyncEffectiveFieldsDuringRead(existingState SearchLoggedModelsOrderBy) {
}

func (c SearchLoggedModelsOrderBy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SearchLoggedModelsOrderBy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchLoggedModelsOrderBy
// only implements ToObjectValue() and Type().
func (o SearchLoggedModelsOrderBy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o SearchLoggedModelsOrderBy) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SearchLoggedModelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SearchLoggedModelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"datasets":       reflect.TypeOf(SearchLoggedModelsDataset{}),
		"experiment_ids": reflect.TypeOf(types.String{}),
		"order_by":       reflect.TypeOf(SearchLoggedModelsOrderBy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchLoggedModelsRequest
// only implements ToObjectValue() and Type().
func (o SearchLoggedModelsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o SearchLoggedModelsRequest) Type(ctx context.Context) attr.Type {
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
func (o *SearchLoggedModelsRequest) GetDatasets(ctx context.Context) ([]SearchLoggedModelsDataset, bool) {
	if o.Datasets.IsNull() || o.Datasets.IsUnknown() {
		return nil, false
	}
	var v []SearchLoggedModelsDataset
	d := o.Datasets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatasets sets the value of the Datasets field in SearchLoggedModelsRequest.
func (o *SearchLoggedModelsRequest) SetDatasets(ctx context.Context, v []SearchLoggedModelsDataset) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["datasets"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Datasets = types.ListValueMust(t, vs)
}

// GetExperimentIds returns the value of the ExperimentIds field in SearchLoggedModelsRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *SearchLoggedModelsRequest) GetExperimentIds(ctx context.Context) ([]types.String, bool) {
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

// SetExperimentIds sets the value of the ExperimentIds field in SearchLoggedModelsRequest.
func (o *SearchLoggedModelsRequest) SetExperimentIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["experiment_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ExperimentIds = types.ListValueMust(t, vs)
}

// GetOrderBy returns the value of the OrderBy field in SearchLoggedModelsRequest as
// a slice of SearchLoggedModelsOrderBy values.
// If the field is unknown or null, the boolean return value is false.
func (o *SearchLoggedModelsRequest) GetOrderBy(ctx context.Context) ([]SearchLoggedModelsOrderBy, bool) {
	if o.OrderBy.IsNull() || o.OrderBy.IsUnknown() {
		return nil, false
	}
	var v []SearchLoggedModelsOrderBy
	d := o.OrderBy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOrderBy sets the value of the OrderBy field in SearchLoggedModelsRequest.
func (o *SearchLoggedModelsRequest) SetOrderBy(ctx context.Context, v []SearchLoggedModelsOrderBy) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["order_by"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OrderBy = types.ListValueMust(t, vs)
}

type SearchLoggedModelsResponse struct {
	// Logged models that match the search criteria.
	Models types.List `tfsdk:"models"`
	// The token that can be used to retrieve the next page of logged models.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (newState *SearchLoggedModelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan SearchLoggedModelsResponse) {
}

func (newState *SearchLoggedModelsResponse) SyncEffectiveFieldsDuringRead(existingState SearchLoggedModelsResponse) {
}

func (c SearchLoggedModelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SearchLoggedModelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"models": reflect.TypeOf(LoggedModel{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchLoggedModelsResponse
// only implements ToObjectValue() and Type().
func (o SearchLoggedModelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"models":          o.Models,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SearchLoggedModelsResponse) Type(ctx context.Context) attr.Type {
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
func (o *SearchLoggedModelsResponse) GetModels(ctx context.Context) ([]LoggedModel, bool) {
	if o.Models.IsNull() || o.Models.IsUnknown() {
		return nil, false
	}
	var v []LoggedModel
	d := o.Models.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModels sets the value of the Models field in SearchLoggedModelsResponse.
func (o *SearchLoggedModelsResponse) SetModels(ctx context.Context, v []LoggedModel) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["models"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Models = types.ListValueMust(t, vs)
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SearchModelVersionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SearchModelVersionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"order_by": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchModelVersionsRequest
// only implements ToObjectValue() and Type().
func (o SearchModelVersionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o SearchModelVersionsRequest) Type(ctx context.Context) attr.Type {
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
func (o *SearchModelVersionsRequest) GetOrderBy(ctx context.Context) ([]types.String, bool) {
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

// SetOrderBy sets the value of the OrderBy field in SearchModelVersionsRequest.
func (o *SearchModelVersionsRequest) SetOrderBy(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["order_by"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OrderBy = types.ListValueMust(t, vs)
}

type SearchModelVersionsResponse struct {
	// Models that match the search criteria
	ModelVersions types.List `tfsdk:"model_versions"`
	// Pagination token to request next page of models for the same search
	// query.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (newState *SearchModelVersionsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan SearchModelVersionsResponse) {
}

func (newState *SearchModelVersionsResponse) SyncEffectiveFieldsDuringRead(existingState SearchModelVersionsResponse) {
}

func (c SearchModelVersionsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SearchModelVersionsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model_versions": reflect.TypeOf(ModelVersion{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchModelVersionsResponse
// only implements ToObjectValue() and Type().
func (o SearchModelVersionsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_versions":  o.ModelVersions,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SearchModelVersionsResponse) Type(ctx context.Context) attr.Type {
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
func (o *SearchModelVersionsResponse) GetModelVersions(ctx context.Context) ([]ModelVersion, bool) {
	if o.ModelVersions.IsNull() || o.ModelVersions.IsUnknown() {
		return nil, false
	}
	var v []ModelVersion
	d := o.ModelVersions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetModelVersions sets the value of the ModelVersions field in SearchModelVersionsResponse.
func (o *SearchModelVersionsResponse) SetModelVersions(ctx context.Context, v []ModelVersion) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["model_versions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ModelVersions = types.ListValueMust(t, vs)
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SearchModelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SearchModelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"order_by": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchModelsRequest
// only implements ToObjectValue() and Type().
func (o SearchModelsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o SearchModelsRequest) Type(ctx context.Context) attr.Type {
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
func (o *SearchModelsRequest) GetOrderBy(ctx context.Context) ([]types.String, bool) {
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

// SetOrderBy sets the value of the OrderBy field in SearchModelsRequest.
func (o *SearchModelsRequest) SetOrderBy(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["order_by"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OrderBy = types.ListValueMust(t, vs)
}

type SearchModelsResponse struct {
	// Pagination token to request the next page of models.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// Registered Models that match the search criteria.
	RegisteredModels types.List `tfsdk:"registered_models"`
}

func (newState *SearchModelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan SearchModelsResponse) {
}

func (newState *SearchModelsResponse) SyncEffectiveFieldsDuringRead(existingState SearchModelsResponse) {
}

func (c SearchModelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SearchModelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"registered_models": reflect.TypeOf(Model{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchModelsResponse
// only implements ToObjectValue() and Type().
func (o SearchModelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token":   o.NextPageToken,
			"registered_models": o.RegisteredModels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SearchModelsResponse) Type(ctx context.Context) attr.Type {
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
func (o *SearchModelsResponse) GetRegisteredModels(ctx context.Context) ([]Model, bool) {
	if o.RegisteredModels.IsNull() || o.RegisteredModels.IsUnknown() {
		return nil, false
	}
	var v []Model
	d := o.RegisteredModels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRegisteredModels sets the value of the RegisteredModels field in SearchModelsResponse.
func (o *SearchModelsResponse) SetRegisteredModels(ctx context.Context, v []Model) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["registered_models"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RegisteredModels = types.ListValueMust(t, vs)
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SearchRuns.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SearchRuns) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"experiment_ids": reflect.TypeOf(types.String{}),
		"order_by":       reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchRuns
// only implements ToObjectValue() and Type().
func (o SearchRuns) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o SearchRuns) Type(ctx context.Context) attr.Type {
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
func (o *SearchRuns) GetExperimentIds(ctx context.Context) ([]types.String, bool) {
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

// SetExperimentIds sets the value of the ExperimentIds field in SearchRuns.
func (o *SearchRuns) SetExperimentIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["experiment_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ExperimentIds = types.ListValueMust(t, vs)
}

// GetOrderBy returns the value of the OrderBy field in SearchRuns as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *SearchRuns) GetOrderBy(ctx context.Context) ([]types.String, bool) {
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

// SetOrderBy sets the value of the OrderBy field in SearchRuns.
func (o *SearchRuns) SetOrderBy(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["order_by"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OrderBy = types.ListValueMust(t, vs)
}

type SearchRunsResponse struct {
	// Token for the next page of runs.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// Runs that match the search criteria.
	Runs types.List `tfsdk:"runs"`
}

func (newState *SearchRunsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan SearchRunsResponse) {
}

func (newState *SearchRunsResponse) SyncEffectiveFieldsDuringRead(existingState SearchRunsResponse) {
}

func (c SearchRunsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SearchRunsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"runs": reflect.TypeOf(Run{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SearchRunsResponse
// only implements ToObjectValue() and Type().
func (o SearchRunsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"runs":            o.Runs,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SearchRunsResponse) Type(ctx context.Context) attr.Type {
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
func (o *SearchRunsResponse) GetRuns(ctx context.Context) ([]Run, bool) {
	if o.Runs.IsNull() || o.Runs.IsUnknown() {
		return nil, false
	}
	var v []Run
	d := o.Runs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRuns sets the value of the Runs field in SearchRunsResponse.
func (o *SearchRunsResponse) SetRuns(ctx context.Context, v []Run) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["runs"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Runs = types.ListValueMust(t, vs)
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetExperimentTag.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SetExperimentTag) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetExperimentTag
// only implements ToObjectValue() and Type().
func (o SetExperimentTag) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": o.ExperimentId,
			"key":           o.Key,
			"value":         o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SetExperimentTag) Type(ctx context.Context) attr.Type {
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

func (newState *SetExperimentTagResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetExperimentTagResponse) {
}

func (newState *SetExperimentTagResponse) SyncEffectiveFieldsDuringRead(existingState SetExperimentTagResponse) {
}

func (c SetExperimentTagResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetExperimentTagResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SetExperimentTagResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetExperimentTagResponse
// only implements ToObjectValue() and Type().
func (o SetExperimentTagResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o SetExperimentTagResponse) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetLoggedModelTagsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SetLoggedModelTagsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(LoggedModelTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetLoggedModelTagsRequest
// only implements ToObjectValue() and Type().
func (o SetLoggedModelTagsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_id": o.ModelId,
			"tags":     o.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SetLoggedModelTagsRequest) Type(ctx context.Context) attr.Type {
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
func (o *SetLoggedModelTagsRequest) GetTags(ctx context.Context) ([]LoggedModelTag, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []LoggedModelTag
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in SetLoggedModelTagsRequest.
func (o *SetLoggedModelTagsRequest) SetTags(ctx context.Context, v []LoggedModelTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type SetLoggedModelTagsResponse struct {
}

func (newState *SetLoggedModelTagsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetLoggedModelTagsResponse) {
}

func (newState *SetLoggedModelTagsResponse) SyncEffectiveFieldsDuringRead(existingState SetLoggedModelTagsResponse) {
}

func (c SetLoggedModelTagsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetLoggedModelTagsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SetLoggedModelTagsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetLoggedModelTagsResponse
// only implements ToObjectValue() and Type().
func (o SetLoggedModelTagsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o SetLoggedModelTagsResponse) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetModelTagRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SetModelTagRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetModelTagRequest
// only implements ToObjectValue() and Type().
func (o SetModelTagRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"name":  o.Name,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SetModelTagRequest) Type(ctx context.Context) attr.Type {
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

func (newState *SetModelTagResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetModelTagResponse) {
}

func (newState *SetModelTagResponse) SyncEffectiveFieldsDuringRead(existingState SetModelTagResponse) {
}

func (c SetModelTagResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetModelTagResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SetModelTagResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetModelTagResponse
// only implements ToObjectValue() and Type().
func (o SetModelTagResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o SetModelTagResponse) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetModelVersionTagRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SetModelVersionTagRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetModelVersionTagRequest
// only implements ToObjectValue() and Type().
func (o SetModelVersionTagRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o SetModelVersionTagRequest) Type(ctx context.Context) attr.Type {
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

func (newState *SetModelVersionTagResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetModelVersionTagResponse) {
}

func (newState *SetModelVersionTagResponse) SyncEffectiveFieldsDuringRead(existingState SetModelVersionTagResponse) {
}

func (c SetModelVersionTagResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetModelVersionTagResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SetModelVersionTagResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetModelVersionTagResponse
// only implements ToObjectValue() and Type().
func (o SetModelVersionTagResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o SetModelVersionTagResponse) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetTag.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SetTag) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetTag
// only implements ToObjectValue() and Type().
func (o SetTag) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o SetTag) Type(ctx context.Context) attr.Type {
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

func (newState *SetTagResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetTagResponse) {
}

func (newState *SetTagResponse) SyncEffectiveFieldsDuringRead(existingState SetTagResponse) {
}

func (c SetTagResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetTagResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SetTagResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetTagResponse
// only implements ToObjectValue() and Type().
func (o SetTagResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o SetTagResponse) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TestRegistryWebhookRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TestRegistryWebhookRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TestRegistryWebhookRequest
// only implements ToObjectValue() and Type().
func (o TestRegistryWebhookRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"event": o.Event,
			"id":    o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TestRegistryWebhookRequest) Type(ctx context.Context) attr.Type {
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

func (newState *TestRegistryWebhookResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan TestRegistryWebhookResponse) {
}

func (newState *TestRegistryWebhookResponse) SyncEffectiveFieldsDuringRead(existingState TestRegistryWebhookResponse) {
}

func (c TestRegistryWebhookResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a TestRegistryWebhookResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TestRegistryWebhookResponse
// only implements ToObjectValue() and Type().
func (o TestRegistryWebhookResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"body":        o.Body,
			"status_code": o.StatusCode,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TestRegistryWebhookResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"body":        types.StringType,
			"status_code": types.Int64Type,
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TransitionModelVersionStageDatabricks.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TransitionModelVersionStageDatabricks) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TransitionModelVersionStageDatabricks
// only implements ToObjectValue() and Type().
func (o TransitionModelVersionStageDatabricks) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o TransitionModelVersionStageDatabricks) Type(ctx context.Context) attr.Type {
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

func (newState *TransitionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan TransitionRequest) {
}

func (newState *TransitionRequest) SyncEffectiveFieldsDuringRead(existingState TransitionRequest) {
}

func (c TransitionRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a TransitionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"available_actions": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TransitionRequest
// only implements ToObjectValue() and Type().
func (o TransitionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o TransitionRequest) Type(ctx context.Context) attr.Type {
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
func (o *TransitionRequest) GetAvailableActions(ctx context.Context) ([]types.String, bool) {
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

// SetAvailableActions sets the value of the AvailableActions field in TransitionRequest.
func (o *TransitionRequest) SetAvailableActions(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["available_actions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AvailableActions = types.ListValueMust(t, vs)
}

type TransitionStageResponse struct {
	// Updated model version
	ModelVersionDatabricks types.Object `tfsdk:"model_version_databricks"`
}

func (newState *TransitionStageResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan TransitionStageResponse) {
}

func (newState *TransitionStageResponse) SyncEffectiveFieldsDuringRead(existingState TransitionStageResponse) {
}

func (c TransitionStageResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a TransitionStageResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model_version_databricks": reflect.TypeOf(ModelVersionDatabricks{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TransitionStageResponse
// only implements ToObjectValue() and Type().
func (o TransitionStageResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_version_databricks": o.ModelVersionDatabricks,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TransitionStageResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_version_databricks": ModelVersionDatabricks{}.Type(ctx),
		},
	}
}

// GetModelVersionDatabricks returns the value of the ModelVersionDatabricks field in TransitionStageResponse as
// a ModelVersionDatabricks value.
// If the field is unknown or null, the boolean return value is false.
func (o *TransitionStageResponse) GetModelVersionDatabricks(ctx context.Context) (ModelVersionDatabricks, bool) {
	var e ModelVersionDatabricks
	if o.ModelVersionDatabricks.IsNull() || o.ModelVersionDatabricks.IsUnknown() {
		return e, false
	}
	var v []ModelVersionDatabricks
	d := o.ModelVersionDatabricks.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetModelVersionDatabricks sets the value of the ModelVersionDatabricks field in TransitionStageResponse.
func (o *TransitionStageResponse) SetModelVersionDatabricks(ctx context.Context, v ModelVersionDatabricks) {
	vs := v.ToObjectValue(ctx)
	o.ModelVersionDatabricks = vs
}

// Details required to edit a comment on a model version.
type UpdateComment struct {
	// User-provided comment on the action.
	Comment types.String `tfsdk:"comment"`
	// Unique identifier of an activity
	Id types.String `tfsdk:"id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateComment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateComment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateComment
// only implements ToObjectValue() and Type().
func (o UpdateComment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment": o.Comment,
			"id":      o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateComment) Type(ctx context.Context) attr.Type {
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

func (newState *UpdateCommentResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateCommentResponse) {
}

func (newState *UpdateCommentResponse) SyncEffectiveFieldsDuringRead(existingState UpdateCommentResponse) {
}

func (c UpdateCommentResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateCommentResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"comment": reflect.TypeOf(CommentObject{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCommentResponse
// only implements ToObjectValue() and Type().
func (o UpdateCommentResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment": o.Comment,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateCommentResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment": CommentObject{}.Type(ctx),
		},
	}
}

// GetComment returns the value of the Comment field in UpdateCommentResponse as
// a CommentObject value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateCommentResponse) GetComment(ctx context.Context) (CommentObject, bool) {
	var e CommentObject
	if o.Comment.IsNull() || o.Comment.IsUnknown() {
		return e, false
	}
	var v []CommentObject
	d := o.Comment.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetComment sets the value of the Comment field in UpdateCommentResponse.
func (o *UpdateCommentResponse) SetComment(ctx context.Context, v CommentObject) {
	vs := v.ToObjectValue(ctx)
	o.Comment = vs
}

type UpdateExperiment struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"experiment_id"`
	// If provided, the experiment's name is changed to the new name. The new
	// name must be unique.
	NewName types.String `tfsdk:"new_name"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateExperiment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateExperiment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateExperiment
// only implements ToObjectValue() and Type().
func (o UpdateExperiment) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"experiment_id": o.ExperimentId,
			"new_name":      o.NewName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateExperiment) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
			"new_name":      types.StringType,
		},
	}
}

type UpdateExperimentResponse struct {
}

func (newState *UpdateExperimentResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateExperimentResponse) {
}

func (newState *UpdateExperimentResponse) SyncEffectiveFieldsDuringRead(existingState UpdateExperimentResponse) {
}

func (c UpdateExperimentResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateExperimentResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateExperimentResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateExperimentResponse
// only implements ToObjectValue() and Type().
func (o UpdateExperimentResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateExperimentResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateFeatureTagRequest struct {
	FeatureName types.String `tfsdk:"-"`

	FeatureTag types.Object `tfsdk:"feature_tag"`

	Key types.String `tfsdk:"-"`

	TableName types.String `tfsdk:"-"`
	// The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateFeatureTagRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateFeatureTagRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"feature_tag": reflect.TypeOf(FeatureTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateFeatureTagRequest
// only implements ToObjectValue() and Type().
func (o UpdateFeatureTagRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o UpdateFeatureTagRequest) Type(ctx context.Context) attr.Type {
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
func (o *UpdateFeatureTagRequest) GetFeatureTag(ctx context.Context) (FeatureTag, bool) {
	var e FeatureTag
	if o.FeatureTag.IsNull() || o.FeatureTag.IsUnknown() {
		return e, false
	}
	var v []FeatureTag
	d := o.FeatureTag.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFeatureTag sets the value of the FeatureTag field in UpdateFeatureTagRequest.
func (o *UpdateFeatureTagRequest) SetFeatureTag(ctx context.Context, v FeatureTag) {
	vs := v.ToObjectValue(ctx)
	o.FeatureTag = vs
}

type UpdateModelRequest struct {
	// If provided, updates the description for this `registered_model`.
	Description types.String `tfsdk:"description"`
	// Registered model unique name identifier.
	Name types.String `tfsdk:"name"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateModelRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateModelRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateModelRequest
// only implements ToObjectValue() and Type().
func (o UpdateModelRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description": o.Description,
			"name":        o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateModelRequest) Type(ctx context.Context) attr.Type {
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

func (newState *UpdateModelResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateModelResponse) {
}

func (newState *UpdateModelResponse) SyncEffectiveFieldsDuringRead(existingState UpdateModelResponse) {
}

func (c UpdateModelResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateModelResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"registered_model": reflect.TypeOf(Model{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateModelResponse
// only implements ToObjectValue() and Type().
func (o UpdateModelResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"registered_model": o.RegisteredModel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateModelResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"registered_model": Model{}.Type(ctx),
		},
	}
}

// GetRegisteredModel returns the value of the RegisteredModel field in UpdateModelResponse as
// a Model value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateModelResponse) GetRegisteredModel(ctx context.Context) (Model, bool) {
	var e Model
	if o.RegisteredModel.IsNull() || o.RegisteredModel.IsUnknown() {
		return e, false
	}
	var v []Model
	d := o.RegisteredModel.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRegisteredModel sets the value of the RegisteredModel field in UpdateModelResponse.
func (o *UpdateModelResponse) SetRegisteredModel(ctx context.Context, v Model) {
	vs := v.ToObjectValue(ctx)
	o.RegisteredModel = vs
}

type UpdateModelVersionRequest struct {
	// If provided, updates the description for this `registered_model`.
	Description types.String `tfsdk:"description"`
	// Name of the registered model
	Name types.String `tfsdk:"name"`
	// Model version number
	Version types.String `tfsdk:"version"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateModelVersionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateModelVersionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateModelVersionRequest
// only implements ToObjectValue() and Type().
func (o UpdateModelVersionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description": o.Description,
			"name":        o.Name,
			"version":     o.Version,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateModelVersionRequest) Type(ctx context.Context) attr.Type {
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

func (newState *UpdateModelVersionResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateModelVersionResponse) {
}

func (newState *UpdateModelVersionResponse) SyncEffectiveFieldsDuringRead(existingState UpdateModelVersionResponse) {
}

func (c UpdateModelVersionResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateModelVersionResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"model_version": reflect.TypeOf(ModelVersion{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateModelVersionResponse
// only implements ToObjectValue() and Type().
func (o UpdateModelVersionResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_version": o.ModelVersion,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateModelVersionResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_version": ModelVersion{}.Type(ctx),
		},
	}
}

// GetModelVersion returns the value of the ModelVersion field in UpdateModelVersionResponse as
// a ModelVersion value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateModelVersionResponse) GetModelVersion(ctx context.Context) (ModelVersion, bool) {
	var e ModelVersion
	if o.ModelVersion.IsNull() || o.ModelVersion.IsUnknown() {
		return e, false
	}
	var v []ModelVersion
	d := o.ModelVersion.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetModelVersion sets the value of the ModelVersion field in UpdateModelVersionResponse.
func (o *UpdateModelVersionResponse) SetModelVersion(ctx context.Context, v ModelVersion) {
	vs := v.ToObjectValue(ctx)
	o.ModelVersion = vs
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateOnlineStoreRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateOnlineStoreRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"online_store": reflect.TypeOf(OnlineStore{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateOnlineStoreRequest
// only implements ToObjectValue() and Type().
func (o UpdateOnlineStoreRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":         o.Name,
			"online_store": o.OnlineStore,
			"update_mask":  o.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateOnlineStoreRequest) Type(ctx context.Context) attr.Type {
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
func (o *UpdateOnlineStoreRequest) GetOnlineStore(ctx context.Context) (OnlineStore, bool) {
	var e OnlineStore
	if o.OnlineStore.IsNull() || o.OnlineStore.IsUnknown() {
		return e, false
	}
	var v []OnlineStore
	d := o.OnlineStore.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOnlineStore sets the value of the OnlineStore field in UpdateOnlineStoreRequest.
func (o *UpdateOnlineStoreRequest) SetOnlineStore(ctx context.Context, v OnlineStore) {
	vs := v.ToObjectValue(ctx)
	o.OnlineStore = vs
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateRegistryWebhook.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateRegistryWebhook) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"events":        reflect.TypeOf(types.String{}),
		"http_url_spec": reflect.TypeOf(HttpUrlSpec{}),
		"job_spec":      reflect.TypeOf(JobSpec{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRegistryWebhook
// only implements ToObjectValue() and Type().
func (o UpdateRegistryWebhook) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o UpdateRegistryWebhook) Type(ctx context.Context) attr.Type {
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
func (o *UpdateRegistryWebhook) GetEvents(ctx context.Context) ([]types.String, bool) {
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

// SetEvents sets the value of the Events field in UpdateRegistryWebhook.
func (o *UpdateRegistryWebhook) SetEvents(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["events"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Events = types.ListValueMust(t, vs)
}

// GetHttpUrlSpec returns the value of the HttpUrlSpec field in UpdateRegistryWebhook as
// a HttpUrlSpec value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateRegistryWebhook) GetHttpUrlSpec(ctx context.Context) (HttpUrlSpec, bool) {
	var e HttpUrlSpec
	if o.HttpUrlSpec.IsNull() || o.HttpUrlSpec.IsUnknown() {
		return e, false
	}
	var v []HttpUrlSpec
	d := o.HttpUrlSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetHttpUrlSpec sets the value of the HttpUrlSpec field in UpdateRegistryWebhook.
func (o *UpdateRegistryWebhook) SetHttpUrlSpec(ctx context.Context, v HttpUrlSpec) {
	vs := v.ToObjectValue(ctx)
	o.HttpUrlSpec = vs
}

// GetJobSpec returns the value of the JobSpec field in UpdateRegistryWebhook as
// a JobSpec value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateRegistryWebhook) GetJobSpec(ctx context.Context) (JobSpec, bool) {
	var e JobSpec
	if o.JobSpec.IsNull() || o.JobSpec.IsUnknown() {
		return e, false
	}
	var v []JobSpec
	d := o.JobSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetJobSpec sets the value of the JobSpec field in UpdateRegistryWebhook.
func (o *UpdateRegistryWebhook) SetJobSpec(ctx context.Context, v JobSpec) {
	vs := v.ToObjectValue(ctx)
	o.JobSpec = vs
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateRun.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateRun) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRun
// only implements ToObjectValue() and Type().
func (o UpdateRun) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o UpdateRun) Type(ctx context.Context) attr.Type {
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

func (newState *UpdateRunResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateRunResponse) {
}

func (newState *UpdateRunResponse) SyncEffectiveFieldsDuringRead(existingState UpdateRunResponse) {
}

func (c UpdateRunResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateRunResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"run_info": reflect.TypeOf(RunInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRunResponse
// only implements ToObjectValue() and Type().
func (o UpdateRunResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run_info": o.RunInfo,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateRunResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run_info": RunInfo{}.Type(ctx),
		},
	}
}

// GetRunInfo returns the value of the RunInfo field in UpdateRunResponse as
// a RunInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateRunResponse) GetRunInfo(ctx context.Context) (RunInfo, bool) {
	var e RunInfo
	if o.RunInfo.IsNull() || o.RunInfo.IsUnknown() {
		return e, false
	}
	var v []RunInfo
	d := o.RunInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRunInfo sets the value of the RunInfo field in UpdateRunResponse.
func (o *UpdateRunResponse) SetRunInfo(ctx context.Context, v RunInfo) {
	vs := v.ToObjectValue(ctx)
	o.RunInfo = vs
}

type UpdateWebhookResponse struct {
	Webhook types.Object `tfsdk:"webhook"`
}

func (newState *UpdateWebhookResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateWebhookResponse) {
}

func (newState *UpdateWebhookResponse) SyncEffectiveFieldsDuringRead(existingState UpdateWebhookResponse) {
}

func (c UpdateWebhookResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateWebhookResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"webhook": reflect.TypeOf(RegistryWebhook{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWebhookResponse
// only implements ToObjectValue() and Type().
func (o UpdateWebhookResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"webhook": o.Webhook,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateWebhookResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"webhook": RegistryWebhook{}.Type(ctx),
		},
	}
}

// GetWebhook returns the value of the Webhook field in UpdateWebhookResponse as
// a RegistryWebhook value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateWebhookResponse) GetWebhook(ctx context.Context) (RegistryWebhook, bool) {
	var e RegistryWebhook
	if o.Webhook.IsNull() || o.Webhook.IsUnknown() {
		return e, false
	}
	var v []RegistryWebhook
	d := o.Webhook.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWebhook sets the value of the Webhook field in UpdateWebhookResponse.
func (o *UpdateWebhookResponse) SetWebhook(ctx context.Context, v RegistryWebhook) {
	vs := v.ToObjectValue(ctx)
	o.Webhook = vs
}
