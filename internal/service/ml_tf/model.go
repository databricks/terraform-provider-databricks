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
	"fmt"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// Activity recorded for the action.
type Activity struct {
	// Type of activity. Valid values are: * `APPLIED_TRANSITION`: User applied
	// the corresponding stage transition.
	//
	// * `REQUESTED_TRANSITION`: User requested the corresponding stage
	// transition.
	//
	// * `CANCELLED_REQUEST`: User cancelled an existing transition request.
	//
	// * `APPROVED_REQUEST`: User approved the corresponding stage transition.
	//
	// * `REJECTED_REQUEST`: User rejected the coressponding stage transition.
	//
	// * `SYSTEM_TRANSITION`: For events performed as a side effect, such as
	// archiving existing model versions in a stage.
	ActivityType types.String `tfsdk:"activity_type" tf:"optional"`
	// User-provided comment associated with the activity.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp" tf:"optional"`
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
	FromStage types.String `tfsdk:"from_stage" tf:"optional"`
	// Unique identifier for the object.
	Id types.String `tfsdk:"id" tf:"optional"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp" tf:"optional"`
	// Comment made by system, for example explaining an activity of type
	// `SYSTEM_TRANSITION`. It usually describes a side effect, such as a
	// version being archived as part of another version's stage transition, and
	// may not be returned for some activity types.
	SystemComment types.String `tfsdk:"system_comment" tf:"optional"`
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
	ToStage types.String `tfsdk:"to_stage" tf:"optional"`
	// The username of the user that created the object.
	UserId types.String `tfsdk:"user_id" tf:"optional"`
}

func (newState *Activity) SyncEffectiveFieldsDuringCreateOrUpdate(plan Activity) {
}

func (newState *Activity) SyncEffectiveFieldsDuringRead(existingState Activity) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = Activity{}

// Equal implements basetypes.ObjectValuable.
func (o Activity) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o Activity) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o Activity) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o Activity) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o Activity) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o Activity) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type ApproveTransitionRequest struct {
	// Specifies whether to archive all current model versions in the target
	// stage.
	ArchiveExistingVersions types.Bool `tfsdk:"archive_existing_versions" tf:""`
	// User-provided comment on the action.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Name of the model.
	Name types.String `tfsdk:"name" tf:""`
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage types.String `tfsdk:"stage" tf:""`
	// Version of the model.
	Version types.String `tfsdk:"version" tf:""`
}

func (newState *ApproveTransitionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ApproveTransitionRequest) {
}

func (newState *ApproveTransitionRequest) SyncEffectiveFieldsDuringRead(existingState ApproveTransitionRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ApproveTransitionRequest{}

// Equal implements basetypes.ObjectValuable.
func (o ApproveTransitionRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ApproveTransitionRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ApproveTransitionRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ApproveTransitionRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ApproveTransitionRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ApproveTransitionRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	// Activity recorded for the action.
	Activity types.List `tfsdk:"activity" tf:"optional,object"`
}

func (newState *ApproveTransitionRequestResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ApproveTransitionRequestResponse) {
}

func (newState *ApproveTransitionRequestResponse) SyncEffectiveFieldsDuringRead(existingState ApproveTransitionRequestResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ApproveTransitionRequestResponse{}

// Equal implements basetypes.ObjectValuable.
func (o ApproveTransitionRequestResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ApproveTransitionRequestResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ApproveTransitionRequestResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ApproveTransitionRequestResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ApproveTransitionRequestResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ApproveTransitionRequestResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o ApproveTransitionRequestResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"activity": basetypes.ListType{
				ElemType: Activity{}.Type(ctx),
			},
		},
	}
}

// Comment details.
type CommentObject struct {
	// Array of actions on the activity allowed for the current viewer.
	AvailableActions types.List `tfsdk:"available_actions" tf:"optional"`
	// User-provided comment on the action.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp" tf:"optional"`
	// Comment ID
	Id types.String `tfsdk:"id" tf:"optional"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp" tf:"optional"`
	// The username of the user that created the object.
	UserId types.String `tfsdk:"user_id" tf:"optional"`
}

func (newState *CommentObject) SyncEffectiveFieldsDuringCreateOrUpdate(plan CommentObject) {
}

func (newState *CommentObject) SyncEffectiveFieldsDuringRead(existingState CommentObject) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CommentObject{}

// Equal implements basetypes.ObjectValuable.
func (o CommentObject) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CommentObject) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CommentObject) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CommentObject) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CommentObject) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CommentObject) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type CreateComment struct {
	// User-provided comment on the action.
	Comment types.String `tfsdk:"comment" tf:""`
	// Name of the model.
	Name types.String `tfsdk:"name" tf:""`
	// Version of the model.
	Version types.String `tfsdk:"version" tf:""`
}

func (newState *CreateComment) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateComment) {
}

func (newState *CreateComment) SyncEffectiveFieldsDuringRead(existingState CreateComment) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateComment{}

// Equal implements basetypes.ObjectValuable.
func (o CreateComment) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateComment) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateComment) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateComment) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateComment) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateComment) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	// Comment details.
	Comment types.List `tfsdk:"comment" tf:"optional,object"`
}

func (newState *CreateCommentResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCommentResponse) {
}

func (newState *CreateCommentResponse) SyncEffectiveFieldsDuringRead(existingState CreateCommentResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateCommentResponse{}

// Equal implements basetypes.ObjectValuable.
func (o CreateCommentResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateCommentResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateCommentResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateCommentResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateCommentResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateCommentResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o CreateCommentResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment": basetypes.ListType{
				ElemType: CommentObject{}.Type(ctx),
			},
		},
	}
}

type CreateExperiment struct {
	// Location where all artifacts for the experiment are stored. If not
	// provided, the remote server will select an appropriate default.
	ArtifactLocation types.String `tfsdk:"artifact_location" tf:"optional"`
	// Experiment name.
	Name types.String `tfsdk:"name" tf:""`
	// A collection of tags to set on the experiment. Maximum tag size and
	// number of tags per request depends on the storage backend. All storage
	// backends are guaranteed to support tag keys up to 250 bytes in size and
	// tag values up to 5000 bytes in size. All storage backends are also
	// guaranteed to support up to 20 tags per request.
	Tags types.List `tfsdk:"tags" tf:"optional"`
}

func (newState *CreateExperiment) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateExperiment) {
}

func (newState *CreateExperiment) SyncEffectiveFieldsDuringRead(existingState CreateExperiment) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateExperiment{}

// Equal implements basetypes.ObjectValuable.
func (o CreateExperiment) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateExperiment) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateExperiment) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateExperiment) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateExperiment) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateExperiment) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type CreateExperimentResponse struct {
	// Unique identifier for the experiment.
	ExperimentId types.String `tfsdk:"experiment_id" tf:"optional"`
}

func (newState *CreateExperimentResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateExperimentResponse) {
}

func (newState *CreateExperimentResponse) SyncEffectiveFieldsDuringRead(existingState CreateExperimentResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateExperimentResponse{}

// Equal implements basetypes.ObjectValuable.
func (o CreateExperimentResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateExperimentResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateExperimentResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateExperimentResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateExperimentResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateExperimentResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o CreateExperimentResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
		},
	}
}

type CreateModelRequest struct {
	// Optional description for registered model.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Register models under this name
	Name types.String `tfsdk:"name" tf:""`
	// Additional metadata for registered model.
	Tags types.List `tfsdk:"tags" tf:"optional"`
}

func (newState *CreateModelRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateModelRequest) {
}

func (newState *CreateModelRequest) SyncEffectiveFieldsDuringRead(existingState CreateModelRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateModelRequest{}

// Equal implements basetypes.ObjectValuable.
func (o CreateModelRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateModelRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateModelRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateModelRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateModelRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateModelRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type CreateModelResponse struct {
	RegisteredModel types.List `tfsdk:"registered_model" tf:"optional,object"`
}

func (newState *CreateModelResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateModelResponse) {
}

func (newState *CreateModelResponse) SyncEffectiveFieldsDuringRead(existingState CreateModelResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateModelResponse{}

// Equal implements basetypes.ObjectValuable.
func (o CreateModelResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateModelResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateModelResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateModelResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateModelResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateModelResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o CreateModelResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"registered_model": basetypes.ListType{
				ElemType: Model{}.Type(ctx),
			},
		},
	}
}

type CreateModelVersionRequest struct {
	// Optional description for model version.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Register model under this name
	Name types.String `tfsdk:"name" tf:""`
	// MLflow run ID for correlation, if `source` was generated by an experiment
	// run in MLflow tracking server
	RunId types.String `tfsdk:"run_id" tf:"optional"`
	// MLflow run link - this is the exact link of the run that generated this
	// model version, potentially hosted at another instance of MLflow.
	RunLink types.String `tfsdk:"run_link" tf:"optional"`
	// URI indicating the location of the model artifacts.
	Source types.String `tfsdk:"source" tf:""`
	// Additional metadata for model version.
	Tags types.List `tfsdk:"tags" tf:"optional"`
}

func (newState *CreateModelVersionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateModelVersionRequest) {
}

func (newState *CreateModelVersionRequest) SyncEffectiveFieldsDuringRead(existingState CreateModelVersionRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateModelVersionRequest{}

// Equal implements basetypes.ObjectValuable.
func (o CreateModelVersionRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateModelVersionRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateModelVersionRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateModelVersionRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateModelVersionRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateModelVersionRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type CreateModelVersionResponse struct {
	// Return new version number generated for this model in registry.
	ModelVersion types.List `tfsdk:"model_version" tf:"optional,object"`
}

func (newState *CreateModelVersionResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateModelVersionResponse) {
}

func (newState *CreateModelVersionResponse) SyncEffectiveFieldsDuringRead(existingState CreateModelVersionResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateModelVersionResponse{}

// Equal implements basetypes.ObjectValuable.
func (o CreateModelVersionResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateModelVersionResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateModelVersionResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateModelVersionResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateModelVersionResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateModelVersionResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o CreateModelVersionResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_version": basetypes.ListType{
				ElemType: ModelVersion{}.Type(ctx),
			},
		},
	}
}

type CreateRegistryWebhook struct {
	// User-specified description for the webhook.
	Description types.String `tfsdk:"description" tf:"optional"`
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
	Events types.List `tfsdk:"events" tf:""`

	HttpUrlSpec types.List `tfsdk:"http_url_spec" tf:"optional,object"`

	JobSpec types.List `tfsdk:"job_spec" tf:"optional,object"`
	// Name of the model whose events would trigger this webhook.
	ModelName types.String `tfsdk:"model_name" tf:"optional"`
	// Enable or disable triggering the webhook, or put the webhook into test
	// mode. The default is `ACTIVE`: * `ACTIVE`: Webhook is triggered when an
	// associated event happens.
	//
	// * `DISABLED`: Webhook is not triggered.
	//
	// * `TEST_MODE`: Webhook can be triggered through the test endpoint, but is
	// not triggered on a real event.
	Status types.String `tfsdk:"status" tf:"optional"`
}

func (newState *CreateRegistryWebhook) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateRegistryWebhook) {
}

func (newState *CreateRegistryWebhook) SyncEffectiveFieldsDuringRead(existingState CreateRegistryWebhook) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateRegistryWebhook{}

// Equal implements basetypes.ObjectValuable.
func (o CreateRegistryWebhook) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateRegistryWebhook) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateRegistryWebhook) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateRegistryWebhook) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateRegistryWebhook) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateRegistryWebhook) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o CreateRegistryWebhook) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"events": basetypes.ListType{
				ElemType: types.StringType,
			},
			"http_url_spec": basetypes.ListType{
				ElemType: HttpUrlSpec{}.Type(ctx),
			},
			"job_spec": basetypes.ListType{
				ElemType: JobSpec{}.Type(ctx),
			},
			"model_name": types.StringType,
			"status":     types.StringType,
		},
	}
}

type CreateRun struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"experiment_id" tf:"optional"`
	// Unix timestamp in milliseconds of when the run started.
	StartTime types.Int64 `tfsdk:"start_time" tf:"optional"`
	// Additional metadata for run.
	Tags types.List `tfsdk:"tags" tf:"optional"`
	// ID of the user executing the run. This field is deprecated as of MLflow
	// 1.0, and will be removed in a future MLflow release. Use 'mlflow.user'
	// tag instead.
	UserId types.String `tfsdk:"user_id" tf:"optional"`
}

func (newState *CreateRun) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateRun) {
}

func (newState *CreateRun) SyncEffectiveFieldsDuringRead(existingState CreateRun) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateRun{}

// Equal implements basetypes.ObjectValuable.
func (o CreateRun) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateRun) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateRun) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateRun) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateRun) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateRun) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o CreateRun) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
			"start_time":    types.Int64Type,
			"tags": basetypes.ListType{
				ElemType: RunTag{}.Type(ctx),
			},
			"user_id": types.StringType,
		},
	}
}

type CreateRunResponse struct {
	// The newly created run.
	Run types.List `tfsdk:"run" tf:"optional,object"`
}

func (newState *CreateRunResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateRunResponse) {
}

func (newState *CreateRunResponse) SyncEffectiveFieldsDuringRead(existingState CreateRunResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateRunResponse{}

// Equal implements basetypes.ObjectValuable.
func (o CreateRunResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateRunResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateRunResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateRunResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateRunResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateRunResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o CreateRunResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run": basetypes.ListType{
				ElemType: Run{}.Type(ctx),
			},
		},
	}
}

type CreateTransitionRequest struct {
	// User-provided comment on the action.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Name of the model.
	Name types.String `tfsdk:"name" tf:""`
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage types.String `tfsdk:"stage" tf:""`
	// Version of the model.
	Version types.String `tfsdk:"version" tf:""`
}

func (newState *CreateTransitionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateTransitionRequest) {
}

func (newState *CreateTransitionRequest) SyncEffectiveFieldsDuringRead(existingState CreateTransitionRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateTransitionRequest{}

// Equal implements basetypes.ObjectValuable.
func (o CreateTransitionRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateTransitionRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateTransitionRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateTransitionRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateTransitionRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateTransitionRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	// Transition request details.
	Request types.List `tfsdk:"request" tf:"optional,object"`
}

func (newState *CreateTransitionRequestResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateTransitionRequestResponse) {
}

func (newState *CreateTransitionRequestResponse) SyncEffectiveFieldsDuringRead(existingState CreateTransitionRequestResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateTransitionRequestResponse{}

// Equal implements basetypes.ObjectValuable.
func (o CreateTransitionRequestResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateTransitionRequestResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateTransitionRequestResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateTransitionRequestResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateTransitionRequestResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateTransitionRequestResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o CreateTransitionRequestResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"request": basetypes.ListType{
				ElemType: TransitionRequest{}.Type(ctx),
			},
		},
	}
}

type CreateWebhookResponse struct {
	Webhook types.List `tfsdk:"webhook" tf:"optional,object"`
}

func (newState *CreateWebhookResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateWebhookResponse) {
}

func (newState *CreateWebhookResponse) SyncEffectiveFieldsDuringRead(existingState CreateWebhookResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = CreateWebhookResponse{}

// Equal implements basetypes.ObjectValuable.
func (o CreateWebhookResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o CreateWebhookResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o CreateWebhookResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o CreateWebhookResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o CreateWebhookResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o CreateWebhookResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o CreateWebhookResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"webhook": basetypes.ListType{
				ElemType: RegistryWebhook{}.Type(ctx),
			},
		},
	}
}

type Dataset struct {
	// Dataset digest, e.g. an md5 hash of the dataset that uniquely identifies
	// it within datasets of the same name.
	Digest types.String `tfsdk:"digest" tf:"optional"`
	// The name of the dataset. E.g. “my.uc.table@2” “nyc-taxi-dataset”,
	// “fantastic-elk-3”
	Name types.String `tfsdk:"name" tf:"optional"`
	// The profile of the dataset. Summary statistics for the dataset, such as
	// the number of rows in a table, the mean / std / mode of each column in a
	// table, or the number of elements in an array.
	Profile types.String `tfsdk:"profile" tf:"optional"`
	// The schema of the dataset. E.g., MLflow ColSpec JSON for a dataframe,
	// MLflow TensorSpec JSON for an ndarray, or another schema format.
	Schema types.String `tfsdk:"schema" tf:"optional"`
	// The type of the dataset source, e.g. ‘databricks-uc-table’,
	// ‘DBFS’, ‘S3’, ...
	Source types.String `tfsdk:"source" tf:"optional"`
	// Source information for the dataset. Note that the source may not exactly
	// reproduce the dataset if it was transformed / modified before use with
	// MLflow.
	SourceType types.String `tfsdk:"source_type" tf:"optional"`
}

func (newState *Dataset) SyncEffectiveFieldsDuringCreateOrUpdate(plan Dataset) {
}

func (newState *Dataset) SyncEffectiveFieldsDuringRead(existingState Dataset) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = Dataset{}

// Equal implements basetypes.ObjectValuable.
func (o Dataset) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o Dataset) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o Dataset) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o Dataset) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o Dataset) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o Dataset) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type DatasetInput struct {
	// The dataset being used as a Run input.
	Dataset types.List `tfsdk:"dataset" tf:"optional,object"`
	// A list of tags for the dataset input, e.g. a “context” tag with value
	// “training”
	Tags types.List `tfsdk:"tags" tf:"optional"`
}

func (newState *DatasetInput) SyncEffectiveFieldsDuringCreateOrUpdate(plan DatasetInput) {
}

func (newState *DatasetInput) SyncEffectiveFieldsDuringRead(existingState DatasetInput) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DatasetInput{}

// Equal implements basetypes.ObjectValuable.
func (o DatasetInput) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DatasetInput) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DatasetInput) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DatasetInput) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DatasetInput) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DatasetInput) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o DatasetInput) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dataset": basetypes.ListType{
				ElemType: Dataset{}.Type(ctx),
			},
			"tags": basetypes.ListType{
				ElemType: InputTag{}.Type(ctx),
			},
		},
	}
}

// Delete a comment
type DeleteCommentRequest struct {
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteCommentRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteCommentRequest) {
}

func (newState *DeleteCommentRequest) SyncEffectiveFieldsDuringRead(existingState DeleteCommentRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteCommentRequest{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteCommentRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteCommentRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteCommentRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteCommentRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteCommentRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteCommentRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteCommentResponse{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteCommentResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteCommentResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteCommentResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteCommentResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteCommentResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteCommentResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCommentResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteExperiment struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"experiment_id" tf:""`
}

func (newState *DeleteExperiment) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteExperiment) {
}

func (newState *DeleteExperiment) SyncEffectiveFieldsDuringRead(existingState DeleteExperiment) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteExperiment{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteExperiment) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteExperiment) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteExperiment) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteExperiment) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteExperiment) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteExperiment) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteExperimentResponse{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteExperimentResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteExperimentResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteExperimentResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteExperimentResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteExperimentResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteExperimentResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o DeleteExperimentResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a model
type DeleteModelRequest struct {
	// Registered model unique name identifier.
	Name types.String `tfsdk:"-"`
}

func (newState *DeleteModelRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteModelRequest) {
}

func (newState *DeleteModelRequest) SyncEffectiveFieldsDuringRead(existingState DeleteModelRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteModelRequest{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteModelRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteModelRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteModelRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteModelRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteModelRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteModelRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteModelResponse{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteModelResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteModelResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteModelResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteModelResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteModelResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteModelResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o DeleteModelResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a model tag
type DeleteModelTagRequest struct {
	// Name of the tag. The name must be an exact match; wild-card deletion is
	// not supported. Maximum size is 250 bytes.
	Key types.String `tfsdk:"-"`
	// Name of the registered model that the tag was logged under.
	Name types.String `tfsdk:"-"`
}

func (newState *DeleteModelTagRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteModelTagRequest) {
}

func (newState *DeleteModelTagRequest) SyncEffectiveFieldsDuringRead(existingState DeleteModelTagRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteModelTagRequest{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteModelTagRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteModelTagRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteModelTagRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteModelTagRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteModelTagRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteModelTagRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteModelTagResponse{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteModelTagResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteModelTagResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteModelTagResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteModelTagResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteModelTagResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteModelTagResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o DeleteModelTagResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a model version.
type DeleteModelVersionRequest struct {
	// Name of the registered model
	Name types.String `tfsdk:"-"`
	// Model version number
	Version types.String `tfsdk:"-"`
}

func (newState *DeleteModelVersionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteModelVersionRequest) {
}

func (newState *DeleteModelVersionRequest) SyncEffectiveFieldsDuringRead(existingState DeleteModelVersionRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteModelVersionRequest{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteModelVersionRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteModelVersionRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteModelVersionRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteModelVersionRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteModelVersionRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteModelVersionRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteModelVersionResponse{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteModelVersionResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteModelVersionResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteModelVersionResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteModelVersionResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteModelVersionResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteModelVersionResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o DeleteModelVersionResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a model version tag
type DeleteModelVersionTagRequest struct {
	// Name of the tag. The name must be an exact match; wild-card deletion is
	// not supported. Maximum size is 250 bytes.
	Key types.String `tfsdk:"-"`
	// Name of the registered model that the tag was logged under.
	Name types.String `tfsdk:"-"`
	// Model version number that the tag was logged under.
	Version types.String `tfsdk:"-"`
}

func (newState *DeleteModelVersionTagRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteModelVersionTagRequest) {
}

func (newState *DeleteModelVersionTagRequest) SyncEffectiveFieldsDuringRead(existingState DeleteModelVersionTagRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteModelVersionTagRequest{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteModelVersionTagRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteModelVersionTagRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteModelVersionTagRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteModelVersionTagRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteModelVersionTagRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteModelVersionTagRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteModelVersionTagResponse{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteModelVersionTagResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteModelVersionTagResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteModelVersionTagResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteModelVersionTagResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteModelVersionTagResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteModelVersionTagResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o DeleteModelVersionTagResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteRun struct {
	// ID of the run to delete.
	RunId types.String `tfsdk:"run_id" tf:""`
}

func (newState *DeleteRun) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteRun) {
}

func (newState *DeleteRun) SyncEffectiveFieldsDuringRead(existingState DeleteRun) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteRun{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteRun) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteRun) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteRun) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteRun) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteRun) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteRun) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteRunResponse{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteRunResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteRunResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteRunResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteRunResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteRunResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteRunResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o DeleteRunResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteRuns struct {
	// The ID of the experiment containing the runs to delete.
	ExperimentId types.String `tfsdk:"experiment_id" tf:""`
	// An optional positive integer indicating the maximum number of runs to
	// delete. The maximum allowed value for max_runs is 10000.
	MaxRuns types.Int64 `tfsdk:"max_runs" tf:"optional"`
	// The maximum creation timestamp in milliseconds since the UNIX epoch for
	// deleting runs. Only runs created prior to or at this timestamp are
	// deleted.
	MaxTimestampMillis types.Int64 `tfsdk:"max_timestamp_millis" tf:""`
}

func (newState *DeleteRuns) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteRuns) {
}

func (newState *DeleteRuns) SyncEffectiveFieldsDuringRead(existingState DeleteRuns) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteRuns{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteRuns) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteRuns) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteRuns) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteRuns) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteRuns) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteRuns) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	RunsDeleted types.Int64 `tfsdk:"runs_deleted" tf:"optional"`
}

func (newState *DeleteRunsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteRunsResponse) {
}

func (newState *DeleteRunsResponse) SyncEffectiveFieldsDuringRead(existingState DeleteRunsResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteRunsResponse{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteRunsResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteRunsResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteRunsResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteRunsResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteRunsResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteRunsResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	Key types.String `tfsdk:"key" tf:""`
	// ID of the run that the tag was logged under. Must be provided.
	RunId types.String `tfsdk:"run_id" tf:""`
}

func (newState *DeleteTag) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteTag) {
}

func (newState *DeleteTag) SyncEffectiveFieldsDuringRead(existingState DeleteTag) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteTag{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteTag) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteTag) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteTag) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteTag) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteTag) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteTag) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteTagResponse{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteTagResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteTagResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteTagResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteTagResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteTagResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteTagResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o DeleteTagResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a transition request
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

func (newState *DeleteTransitionRequestRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteTransitionRequestRequest) {
}

func (newState *DeleteTransitionRequestRequest) SyncEffectiveFieldsDuringRead(existingState DeleteTransitionRequestRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteTransitionRequestRequest{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteTransitionRequestRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteTransitionRequestRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteTransitionRequestRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteTransitionRequestRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteTransitionRequestRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteTransitionRequestRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
}

func (newState *DeleteTransitionRequestResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteTransitionRequestResponse) {
}

func (newState *DeleteTransitionRequestResponse) SyncEffectiveFieldsDuringRead(existingState DeleteTransitionRequestResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteTransitionRequestResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteTransitionRequestResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteTransitionRequestResponse{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteTransitionRequestResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteTransitionRequestResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteTransitionRequestResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteTransitionRequestResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteTransitionRequestResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteTransitionRequestResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o DeleteTransitionRequestResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a webhook
type DeleteWebhookRequest struct {
	// Webhook ID required to delete a registry webhook.
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteWebhookRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteWebhookRequest) {
}

func (newState *DeleteWebhookRequest) SyncEffectiveFieldsDuringRead(existingState DeleteWebhookRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteWebhookRequest{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteWebhookRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteWebhookRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteWebhookRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteWebhookRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteWebhookRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteWebhookRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = DeleteWebhookResponse{}

// Equal implements basetypes.ObjectValuable.
func (o DeleteWebhookResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o DeleteWebhookResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o DeleteWebhookResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o DeleteWebhookResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o DeleteWebhookResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o DeleteWebhookResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o DeleteWebhookResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type Experiment struct {
	// Location where artifacts for the experiment are stored.
	ArtifactLocation types.String `tfsdk:"artifact_location" tf:"optional"`
	// Creation time
	CreationTime types.Int64 `tfsdk:"creation_time" tf:"optional"`
	// Unique identifier for the experiment.
	ExperimentId types.String `tfsdk:"experiment_id" tf:"optional"`
	// Last update time
	LastUpdateTime types.Int64 `tfsdk:"last_update_time" tf:"optional"`
	// Current life cycle stage of the experiment: "active" or "deleted".
	// Deleted experiments are not returned by APIs.
	LifecycleStage types.String `tfsdk:"lifecycle_stage" tf:"optional"`
	// Human readable name that identifies the experiment.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Tags: Additional metadata key-value pairs.
	Tags types.List `tfsdk:"tags" tf:"optional"`
}

func (newState *Experiment) SyncEffectiveFieldsDuringCreateOrUpdate(plan Experiment) {
}

func (newState *Experiment) SyncEffectiveFieldsDuringRead(existingState Experiment) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = Experiment{}

// Equal implements basetypes.ObjectValuable.
func (o Experiment) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o Experiment) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o Experiment) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o Experiment) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o Experiment) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o Experiment) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type ExperimentAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *ExperimentAccessControlRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExperimentAccessControlRequest) {
}

func (newState *ExperimentAccessControlRequest) SyncEffectiveFieldsDuringRead(existingState ExperimentAccessControlRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ExperimentAccessControlRequest{}

// Equal implements basetypes.ObjectValuable.
func (o ExperimentAccessControlRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ExperimentAccessControlRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ExperimentAccessControlRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ExperimentAccessControlRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ExperimentAccessControlRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ExperimentAccessControlRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	AllPermissions types.List `tfsdk:"all_permissions" tf:"optional"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *ExperimentAccessControlResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExperimentAccessControlResponse) {
}

func (newState *ExperimentAccessControlResponse) SyncEffectiveFieldsDuringRead(existingState ExperimentAccessControlResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ExperimentAccessControlResponse{}

// Equal implements basetypes.ObjectValuable.
func (o ExperimentAccessControlResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ExperimentAccessControlResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ExperimentAccessControlResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ExperimentAccessControlResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ExperimentAccessControlResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ExperimentAccessControlResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type ExperimentPermission struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *ExperimentPermission) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExperimentPermission) {
}

func (newState *ExperimentPermission) SyncEffectiveFieldsDuringRead(existingState ExperimentPermission) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ExperimentPermission{}

// Equal implements basetypes.ObjectValuable.
func (o ExperimentPermission) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ExperimentPermission) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ExperimentPermission) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ExperimentPermission) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ExperimentPermission) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ExperimentPermission) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type ExperimentPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

func (newState *ExperimentPermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExperimentPermissions) {
}

func (newState *ExperimentPermissions) SyncEffectiveFieldsDuringRead(existingState ExperimentPermissions) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ExperimentPermissions{}

// Equal implements basetypes.ObjectValuable.
func (o ExperimentPermissions) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ExperimentPermissions) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ExperimentPermissions) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ExperimentPermissions) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ExperimentPermissions) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ExperimentPermissions) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type ExperimentPermissionsDescription struct {
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *ExperimentPermissionsDescription) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExperimentPermissionsDescription) {
}

func (newState *ExperimentPermissionsDescription) SyncEffectiveFieldsDuringRead(existingState ExperimentPermissionsDescription) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ExperimentPermissionsDescription{}

// Equal implements basetypes.ObjectValuable.
func (o ExperimentPermissionsDescription) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ExperimentPermissionsDescription) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ExperimentPermissionsDescription) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ExperimentPermissionsDescription) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ExperimentPermissionsDescription) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ExperimentPermissionsDescription) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`
	// The experiment for which to get or manage permissions.
	ExperimentId types.String `tfsdk:"-"`
}

func (newState *ExperimentPermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExperimentPermissionsRequest) {
}

func (newState *ExperimentPermissionsRequest) SyncEffectiveFieldsDuringRead(existingState ExperimentPermissionsRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ExperimentPermissionsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o ExperimentPermissionsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ExperimentPermissionsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ExperimentPermissionsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ExperimentPermissionsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ExperimentPermissionsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ExperimentPermissionsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type ExperimentTag struct {
	// The tag key.
	Key types.String `tfsdk:"key" tf:"optional"`
	// The tag value.
	Value types.String `tfsdk:"value" tf:"optional"`
}

func (newState *ExperimentTag) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExperimentTag) {
}

func (newState *ExperimentTag) SyncEffectiveFieldsDuringRead(existingState ExperimentTag) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ExperimentTag{}

// Equal implements basetypes.ObjectValuable.
func (o ExperimentTag) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ExperimentTag) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ExperimentTag) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ExperimentTag) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ExperimentTag) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ExperimentTag) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type FileInfo struct {
	// Size in bytes. Unset for directories.
	FileSize types.Int64 `tfsdk:"file_size" tf:"optional"`
	// Whether the path is a directory.
	IsDir types.Bool `tfsdk:"is_dir" tf:"optional"`
	// Path relative to the root artifact directory run.
	Path types.String `tfsdk:"path" tf:"optional"`
}

func (newState *FileInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan FileInfo) {
}

func (newState *FileInfo) SyncEffectiveFieldsDuringRead(existingState FileInfo) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = FileInfo{}

// Equal implements basetypes.ObjectValuable.
func (o FileInfo) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o FileInfo) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o FileInfo) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o FileInfo) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o FileInfo) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o FileInfo) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// Get metadata
type GetByNameRequest struct {
	// Name of the associated experiment.
	ExperimentName types.String `tfsdk:"-"`
}

func (newState *GetByNameRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetByNameRequest) {
}

func (newState *GetByNameRequest) SyncEffectiveFieldsDuringRead(existingState GetByNameRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetByNameRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetByNameRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetByNameRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetByNameRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetByNameRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetByNameRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetByNameRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GetByNameRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_name": types.StringType,
		},
	}
}

// Get experiment permission levels
type GetExperimentPermissionLevelsRequest struct {
	// The experiment for which to get or manage permissions.
	ExperimentId types.String `tfsdk:"-"`
}

func (newState *GetExperimentPermissionLevelsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetExperimentPermissionLevelsRequest) {
}

func (newState *GetExperimentPermissionLevelsRequest) SyncEffectiveFieldsDuringRead(existingState GetExperimentPermissionLevelsRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetExperimentPermissionLevelsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetExperimentPermissionLevelsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetExperimentPermissionLevelsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetExperimentPermissionLevelsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetExperimentPermissionLevelsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetExperimentPermissionLevelsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetExperimentPermissionLevelsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	PermissionLevels types.List `tfsdk:"permission_levels" tf:"optional"`
}

func (newState *GetExperimentPermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetExperimentPermissionLevelsResponse) {
}

func (newState *GetExperimentPermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetExperimentPermissionLevelsResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetExperimentPermissionLevelsResponse{}

// Equal implements basetypes.ObjectValuable.
func (o GetExperimentPermissionLevelsResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetExperimentPermissionLevelsResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetExperimentPermissionLevelsResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetExperimentPermissionLevelsResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetExperimentPermissionLevelsResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetExperimentPermissionLevelsResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// Get experiment permissions
type GetExperimentPermissionsRequest struct {
	// The experiment for which to get or manage permissions.
	ExperimentId types.String `tfsdk:"-"`
}

func (newState *GetExperimentPermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetExperimentPermissionsRequest) {
}

func (newState *GetExperimentPermissionsRequest) SyncEffectiveFieldsDuringRead(existingState GetExperimentPermissionsRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetExperimentPermissionsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetExperimentPermissionsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetExperimentPermissionsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetExperimentPermissionsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetExperimentPermissionsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetExperimentPermissionsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetExperimentPermissionsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GetExperimentPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment_id": types.StringType,
		},
	}
}

// Get an experiment
type GetExperimentRequest struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"-"`
}

func (newState *GetExperimentRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetExperimentRequest) {
}

func (newState *GetExperimentRequest) SyncEffectiveFieldsDuringRead(existingState GetExperimentRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetExperimentRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetExperimentRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetExperimentRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetExperimentRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetExperimentRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetExperimentRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetExperimentRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	Experiment types.List `tfsdk:"experiment" tf:"optional,object"`
}

func (newState *GetExperimentResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetExperimentResponse) {
}

func (newState *GetExperimentResponse) SyncEffectiveFieldsDuringRead(existingState GetExperimentResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetExperimentResponse{}

// Equal implements basetypes.ObjectValuable.
func (o GetExperimentResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetExperimentResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetExperimentResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetExperimentResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetExperimentResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetExperimentResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GetExperimentResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"experiment": basetypes.ListType{
				ElemType: Experiment{}.Type(ctx),
			},
		},
	}
}

// Get history of a given metric within a run
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
	// [Deprecated, use run_id instead] ID of the run from which to fetch metric
	// values. This field will be removed in a future MLflow version.
	RunUuid types.String `tfsdk:"-"`
}

func (newState *GetHistoryRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetHistoryRequest) {
}

func (newState *GetHistoryRequest) SyncEffectiveFieldsDuringRead(existingState GetHistoryRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetHistoryRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetHistoryRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetHistoryRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetHistoryRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetHistoryRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetHistoryRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetHistoryRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	Name types.String `tfsdk:"name" tf:""`
	// List of stages.
	Stages types.List `tfsdk:"stages" tf:"optional"`
}

func (newState *GetLatestVersionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetLatestVersionsRequest) {
}

func (newState *GetLatestVersionsRequest) SyncEffectiveFieldsDuringRead(existingState GetLatestVersionsRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetLatestVersionsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetLatestVersionsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetLatestVersionsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetLatestVersionsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetLatestVersionsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetLatestVersionsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetLatestVersionsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type GetLatestVersionsResponse struct {
	// Latest version models for each requests stage. Only return models with
	// current `READY` status. If no `stages` provided, returns the latest
	// version for each stage, including `"None"`.
	ModelVersions types.List `tfsdk:"model_versions" tf:"optional"`
}

func (newState *GetLatestVersionsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetLatestVersionsResponse) {
}

func (newState *GetLatestVersionsResponse) SyncEffectiveFieldsDuringRead(existingState GetLatestVersionsResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetLatestVersionsResponse{}

// Equal implements basetypes.ObjectValuable.
func (o GetLatestVersionsResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetLatestVersionsResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetLatestVersionsResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetLatestVersionsResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetLatestVersionsResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetLatestVersionsResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type GetMetricHistoryResponse struct {
	// All logged values for this metric.
	Metrics types.List `tfsdk:"metrics" tf:"optional"`
	// Token that can be used to retrieve the next page of metric history
	// results
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *GetMetricHistoryResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetMetricHistoryResponse) {
}

func (newState *GetMetricHistoryResponse) SyncEffectiveFieldsDuringRead(existingState GetMetricHistoryResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetMetricHistoryResponse{}

// Equal implements basetypes.ObjectValuable.
func (o GetMetricHistoryResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetMetricHistoryResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetMetricHistoryResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetMetricHistoryResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetMetricHistoryResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetMetricHistoryResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// Get model
type GetModelRequest struct {
	// Registered model unique name identifier.
	Name types.String `tfsdk:"-"`
}

func (newState *GetModelRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetModelRequest) {
}

func (newState *GetModelRequest) SyncEffectiveFieldsDuringRead(existingState GetModelRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetModelRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetModelRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetModelRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetModelRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetModelRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetModelRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetModelRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	RegisteredModelDatabricks types.List `tfsdk:"registered_model_databricks" tf:"optional,object"`
}

func (newState *GetModelResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetModelResponse) {
}

func (newState *GetModelResponse) SyncEffectiveFieldsDuringRead(existingState GetModelResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetModelResponse{}

// Equal implements basetypes.ObjectValuable.
func (o GetModelResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetModelResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetModelResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetModelResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetModelResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetModelResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GetModelResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"registered_model_databricks": basetypes.ListType{
				ElemType: ModelDatabricks{}.Type(ctx),
			},
		},
	}
}

// Get a model version URI
type GetModelVersionDownloadUriRequest struct {
	// Name of the registered model
	Name types.String `tfsdk:"-"`
	// Model version number
	Version types.String `tfsdk:"-"`
}

func (newState *GetModelVersionDownloadUriRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetModelVersionDownloadUriRequest) {
}

func (newState *GetModelVersionDownloadUriRequest) SyncEffectiveFieldsDuringRead(existingState GetModelVersionDownloadUriRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetModelVersionDownloadUriRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetModelVersionDownloadUriRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetModelVersionDownloadUriRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetModelVersionDownloadUriRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetModelVersionDownloadUriRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetModelVersionDownloadUriRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetModelVersionDownloadUriRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	ArtifactUri types.String `tfsdk:"artifact_uri" tf:"optional"`
}

func (newState *GetModelVersionDownloadUriResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetModelVersionDownloadUriResponse) {
}

func (newState *GetModelVersionDownloadUriResponse) SyncEffectiveFieldsDuringRead(existingState GetModelVersionDownloadUriResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetModelVersionDownloadUriResponse{}

// Equal implements basetypes.ObjectValuable.
func (o GetModelVersionDownloadUriResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetModelVersionDownloadUriResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetModelVersionDownloadUriResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetModelVersionDownloadUriResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetModelVersionDownloadUriResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetModelVersionDownloadUriResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GetModelVersionDownloadUriResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"artifact_uri": types.StringType,
		},
	}
}

// Get a model version
type GetModelVersionRequest struct {
	// Name of the registered model
	Name types.String `tfsdk:"-"`
	// Model version number
	Version types.String `tfsdk:"-"`
}

func (newState *GetModelVersionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetModelVersionRequest) {
}

func (newState *GetModelVersionRequest) SyncEffectiveFieldsDuringRead(existingState GetModelVersionRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetModelVersionRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetModelVersionRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetModelVersionRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetModelVersionRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetModelVersionRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetModelVersionRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetModelVersionRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	ModelVersion types.List `tfsdk:"model_version" tf:"optional,object"`
}

func (newState *GetModelVersionResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetModelVersionResponse) {
}

func (newState *GetModelVersionResponse) SyncEffectiveFieldsDuringRead(existingState GetModelVersionResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetModelVersionResponse{}

// Equal implements basetypes.ObjectValuable.
func (o GetModelVersionResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetModelVersionResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetModelVersionResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetModelVersionResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetModelVersionResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetModelVersionResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GetModelVersionResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_version": basetypes.ListType{
				ElemType: ModelVersion{}.Type(ctx),
			},
		},
	}
}

// Get registered model permission levels
type GetRegisteredModelPermissionLevelsRequest struct {
	// The registered model for which to get or manage permissions.
	RegisteredModelId types.String `tfsdk:"-"`
}

func (newState *GetRegisteredModelPermissionLevelsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRegisteredModelPermissionLevelsRequest) {
}

func (newState *GetRegisteredModelPermissionLevelsRequest) SyncEffectiveFieldsDuringRead(existingState GetRegisteredModelPermissionLevelsRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetRegisteredModelPermissionLevelsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetRegisteredModelPermissionLevelsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetRegisteredModelPermissionLevelsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetRegisteredModelPermissionLevelsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetRegisteredModelPermissionLevelsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetRegisteredModelPermissionLevelsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetRegisteredModelPermissionLevelsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	PermissionLevels types.List `tfsdk:"permission_levels" tf:"optional"`
}

func (newState *GetRegisteredModelPermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRegisteredModelPermissionLevelsResponse) {
}

func (newState *GetRegisteredModelPermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetRegisteredModelPermissionLevelsResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetRegisteredModelPermissionLevelsResponse{}

// Equal implements basetypes.ObjectValuable.
func (o GetRegisteredModelPermissionLevelsResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetRegisteredModelPermissionLevelsResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetRegisteredModelPermissionLevelsResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetRegisteredModelPermissionLevelsResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetRegisteredModelPermissionLevelsResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetRegisteredModelPermissionLevelsResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// Get registered model permissions
type GetRegisteredModelPermissionsRequest struct {
	// The registered model for which to get or manage permissions.
	RegisteredModelId types.String `tfsdk:"-"`
}

func (newState *GetRegisteredModelPermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRegisteredModelPermissionsRequest) {
}

func (newState *GetRegisteredModelPermissionsRequest) SyncEffectiveFieldsDuringRead(existingState GetRegisteredModelPermissionsRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetRegisteredModelPermissionsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetRegisteredModelPermissionsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetRegisteredModelPermissionsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetRegisteredModelPermissionsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetRegisteredModelPermissionsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetRegisteredModelPermissionsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetRegisteredModelPermissionsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GetRegisteredModelPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"registered_model_id": types.StringType,
		},
	}
}

// Get a run
type GetRunRequest struct {
	// ID of the run to fetch. Must be provided.
	RunId types.String `tfsdk:"-"`
	// [Deprecated, use run_id instead] ID of the run to fetch. This field will
	// be removed in a future MLflow version.
	RunUuid types.String `tfsdk:"-"`
}

func (newState *GetRunRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRunRequest) {
}

func (newState *GetRunRequest) SyncEffectiveFieldsDuringRead(existingState GetRunRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetRunRequest{}

// Equal implements basetypes.ObjectValuable.
func (o GetRunRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetRunRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetRunRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetRunRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetRunRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetRunRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	Run types.List `tfsdk:"run" tf:"optional,object"`
}

func (newState *GetRunResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRunResponse) {
}

func (newState *GetRunResponse) SyncEffectiveFieldsDuringRead(existingState GetRunResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = GetRunResponse{}

// Equal implements basetypes.ObjectValuable.
func (o GetRunResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o GetRunResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o GetRunResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o GetRunResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o GetRunResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o GetRunResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o GetRunResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run": basetypes.ListType{
				ElemType: Run{}.Type(ctx),
			},
		},
	}
}

type HttpUrlSpec struct {
	// Value of the authorization header that should be sent in the request sent
	// by the wehbook. It should be of the form `"<auth type> <credentials>"`.
	// If set to an empty string, no authorization header will be included in
	// the request.
	Authorization types.String `tfsdk:"authorization" tf:"optional"`
	// Enable/disable SSL certificate validation. Default is true. For
	// self-signed certificates, this field must be false AND the destination
	// server must disable certificate validation as well. For security
	// purposes, it is encouraged to perform secret validation with the
	// HMAC-encoded portion of the payload and acknowledge the risk associated
	// with disabling hostname validation whereby it becomes more likely that
	// requests can be maliciously routed to an unintended host.
	EnableSslVerification types.Bool `tfsdk:"enable_ssl_verification" tf:"optional"`
	// Shared secret required for HMAC encoding payload. The HMAC-encoded
	// payload will be sent in the header as: { "X-Databricks-Signature":
	// $encoded_payload }.
	Secret types.String `tfsdk:"secret" tf:"optional"`
	// External HTTPS URL called on event trigger (by using a POST request).
	Url types.String `tfsdk:"url" tf:""`
}

func (newState *HttpUrlSpec) SyncEffectiveFieldsDuringCreateOrUpdate(plan HttpUrlSpec) {
}

func (newState *HttpUrlSpec) SyncEffectiveFieldsDuringRead(existingState HttpUrlSpec) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = HttpUrlSpec{}

// Equal implements basetypes.ObjectValuable.
func (o HttpUrlSpec) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o HttpUrlSpec) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o HttpUrlSpec) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o HttpUrlSpec) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o HttpUrlSpec) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o HttpUrlSpec) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	EnableSslVerification types.Bool `tfsdk:"enable_ssl_verification" tf:"optional"`
	// External HTTPS URL called on event trigger (by using a POST request).
	Url types.String `tfsdk:"url" tf:"optional"`
}

func (newState *HttpUrlSpecWithoutSecret) SyncEffectiveFieldsDuringCreateOrUpdate(plan HttpUrlSpecWithoutSecret) {
}

func (newState *HttpUrlSpecWithoutSecret) SyncEffectiveFieldsDuringRead(existingState HttpUrlSpecWithoutSecret) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = HttpUrlSpecWithoutSecret{}

// Equal implements basetypes.ObjectValuable.
func (o HttpUrlSpecWithoutSecret) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o HttpUrlSpecWithoutSecret) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o HttpUrlSpecWithoutSecret) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o HttpUrlSpecWithoutSecret) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o HttpUrlSpecWithoutSecret) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o HttpUrlSpecWithoutSecret) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type InputTag struct {
	// The tag key.
	Key types.String `tfsdk:"key" tf:"optional"`
	// The tag value.
	Value types.String `tfsdk:"value" tf:"optional"`
}

func (newState *InputTag) SyncEffectiveFieldsDuringCreateOrUpdate(plan InputTag) {
}

func (newState *InputTag) SyncEffectiveFieldsDuringRead(existingState InputTag) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = InputTag{}

// Equal implements basetypes.ObjectValuable.
func (o InputTag) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o InputTag) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o InputTag) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o InputTag) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o InputTag) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o InputTag) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	AccessToken types.String `tfsdk:"access_token" tf:""`
	// ID of the job that the webhook runs.
	JobId types.String `tfsdk:"job_id" tf:""`
	// URL of the workspace containing the job that this webhook runs. If not
	// specified, the job’s workspace URL is assumed to be the same as the
	// workspace where the webhook is created.
	WorkspaceUrl types.String `tfsdk:"workspace_url" tf:"optional"`
}

func (newState *JobSpec) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobSpec) {
}

func (newState *JobSpec) SyncEffectiveFieldsDuringRead(existingState JobSpec) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = JobSpec{}

// Equal implements basetypes.ObjectValuable.
func (o JobSpec) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o JobSpec) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o JobSpec) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o JobSpec) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o JobSpec) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o JobSpec) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	JobId types.String `tfsdk:"job_id" tf:"optional"`
	// URL of the workspace containing the job that this webhook runs. Defaults
	// to the workspace URL in which the webhook is created. If not specified,
	// the job’s workspace is assumed to be the same as the webhook’s.
	WorkspaceUrl types.String `tfsdk:"workspace_url" tf:"optional"`
}

func (newState *JobSpecWithoutSecret) SyncEffectiveFieldsDuringCreateOrUpdate(plan JobSpecWithoutSecret) {
}

func (newState *JobSpecWithoutSecret) SyncEffectiveFieldsDuringRead(existingState JobSpecWithoutSecret) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = JobSpecWithoutSecret{}

// Equal implements basetypes.ObjectValuable.
func (o JobSpecWithoutSecret) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o JobSpecWithoutSecret) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o JobSpecWithoutSecret) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o JobSpecWithoutSecret) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o JobSpecWithoutSecret) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o JobSpecWithoutSecret) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// Get all artifacts
type ListArtifactsRequest struct {
	// Token indicating the page of artifact results to fetch. `page_token` is
	// not supported when listing artifacts in UC Volumes. A maximum of 1000
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
	// [Deprecated, use run_id instead] ID of the run whose artifacts to list.
	// This field will be removed in a future MLflow version.
	RunUuid types.String `tfsdk:"-"`
}

func (newState *ListArtifactsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListArtifactsRequest) {
}

func (newState *ListArtifactsRequest) SyncEffectiveFieldsDuringRead(existingState ListArtifactsRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListArtifactsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o ListArtifactsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListArtifactsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListArtifactsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListArtifactsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListArtifactsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListArtifactsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	// File location and metadata for artifacts.
	Files types.List `tfsdk:"files" tf:"optional"`
	// Token that can be used to retrieve the next page of artifact results
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// Root artifact directory for the run.
	RootUri types.String `tfsdk:"root_uri" tf:"optional"`
}

func (newState *ListArtifactsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListArtifactsResponse) {
}

func (newState *ListArtifactsResponse) SyncEffectiveFieldsDuringRead(existingState ListArtifactsResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListArtifactsResponse{}

// Equal implements basetypes.ObjectValuable.
func (o ListArtifactsResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListArtifactsResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListArtifactsResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListArtifactsResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListArtifactsResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListArtifactsResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// List experiments
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

func (newState *ListExperimentsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListExperimentsRequest) {
}

func (newState *ListExperimentsRequest) SyncEffectiveFieldsDuringRead(existingState ListExperimentsRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListExperimentsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o ListExperimentsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListExperimentsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListExperimentsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListExperimentsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListExperimentsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListExperimentsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	Experiments types.List `tfsdk:"experiments" tf:"optional"`
	// Token that can be used to retrieve the next page of experiments. Empty
	// token means no more experiment is available for retrieval.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListExperimentsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListExperimentsResponse) {
}

func (newState *ListExperimentsResponse) SyncEffectiveFieldsDuringRead(existingState ListExperimentsResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListExperimentsResponse{}

// Equal implements basetypes.ObjectValuable.
func (o ListExperimentsResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListExperimentsResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListExperimentsResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListExperimentsResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListExperimentsResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListExperimentsResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// List models
type ListModelsRequest struct {
	// Maximum number of registered models desired. Max threshold is 1000.
	MaxResults types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page based on a previous query.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListModelsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListModelsRequest) {
}

func (newState *ListModelsRequest) SyncEffectiveFieldsDuringRead(existingState ListModelsRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListModelsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o ListModelsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListModelsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListModelsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListModelsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListModelsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListModelsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	RegisteredModels types.List `tfsdk:"registered_models" tf:"optional"`
}

func (newState *ListModelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListModelsResponse) {
}

func (newState *ListModelsResponse) SyncEffectiveFieldsDuringRead(existingState ListModelsResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListModelsResponse{}

// Equal implements basetypes.ObjectValuable.
func (o ListModelsResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListModelsResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListModelsResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListModelsResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListModelsResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListModelsResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type ListRegistryWebhooks struct {
	// Token that can be used to retrieve the next page of artifact results
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// Array of registry webhooks.
	Webhooks types.List `tfsdk:"webhooks" tf:"optional"`
}

func (newState *ListRegistryWebhooks) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListRegistryWebhooks) {
}

func (newState *ListRegistryWebhooks) SyncEffectiveFieldsDuringRead(existingState ListRegistryWebhooks) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListRegistryWebhooks{}

// Equal implements basetypes.ObjectValuable.
func (o ListRegistryWebhooks) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListRegistryWebhooks) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListRegistryWebhooks) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListRegistryWebhooks) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListRegistryWebhooks) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListRegistryWebhooks) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// List transition requests
type ListTransitionRequestsRequest struct {
	// Name of the model.
	Name types.String `tfsdk:"-"`
	// Version of the model.
	Version types.String `tfsdk:"-"`
}

func (newState *ListTransitionRequestsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListTransitionRequestsRequest) {
}

func (newState *ListTransitionRequestsRequest) SyncEffectiveFieldsDuringRead(existingState ListTransitionRequestsRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListTransitionRequestsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o ListTransitionRequestsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListTransitionRequestsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListTransitionRequestsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListTransitionRequestsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListTransitionRequestsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListTransitionRequestsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	Requests types.List `tfsdk:"requests" tf:"optional"`
}

func (newState *ListTransitionRequestsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListTransitionRequestsResponse) {
}

func (newState *ListTransitionRequestsResponse) SyncEffectiveFieldsDuringRead(existingState ListTransitionRequestsResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListTransitionRequestsResponse{}

// Equal implements basetypes.ObjectValuable.
func (o ListTransitionRequestsResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListTransitionRequestsResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListTransitionRequestsResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListTransitionRequestsResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListTransitionRequestsResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListTransitionRequestsResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// List registry webhooks
type ListWebhooksRequest struct {
	// If `events` is specified, any webhook with one or more of the specified
	// trigger events is included in the output. If `events` is not specified,
	// webhooks of all event types are included in the output.
	Events types.List `tfsdk:"-"`
	// If not specified, all webhooks associated with the specified events are
	// listed, regardless of their associated model.
	ModelName types.String `tfsdk:"-"`
	// Token indicating the page of artifact results to fetch
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListWebhooksRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListWebhooksRequest) {
}

func (newState *ListWebhooksRequest) SyncEffectiveFieldsDuringRead(existingState ListWebhooksRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ListWebhooksRequest{}

// Equal implements basetypes.ObjectValuable.
func (o ListWebhooksRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ListWebhooksRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ListWebhooksRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ListWebhooksRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ListWebhooksRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ListWebhooksRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o ListWebhooksRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"events": basetypes.ListType{
				ElemType: types.StringType,
			},
			"model_name": types.StringType,
			"page_token": types.StringType,
		},
	}
}

type LogBatch struct {
	// Metrics to log. A single request can contain up to 1000 metrics, and up
	// to 1000 metrics, params, and tags in total.
	Metrics types.List `tfsdk:"metrics" tf:"optional"`
	// Params to log. A single request can contain up to 100 params, and up to
	// 1000 metrics, params, and tags in total.
	Params types.List `tfsdk:"params" tf:"optional"`
	// ID of the run to log under
	RunId types.String `tfsdk:"run_id" tf:"optional"`
	// Tags to log. A single request can contain up to 100 tags, and up to 1000
	// metrics, params, and tags in total.
	Tags types.List `tfsdk:"tags" tf:"optional"`
}

func (newState *LogBatch) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogBatch) {
}

func (newState *LogBatch) SyncEffectiveFieldsDuringRead(existingState LogBatch) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = LogBatch{}

// Equal implements basetypes.ObjectValuable.
func (o LogBatch) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o LogBatch) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o LogBatch) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o LogBatch) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o LogBatch) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o LogBatch) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type LogBatchResponse struct {
}

func (newState *LogBatchResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogBatchResponse) {
}

func (newState *LogBatchResponse) SyncEffectiveFieldsDuringRead(existingState LogBatchResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = LogBatchResponse{}

// Equal implements basetypes.ObjectValuable.
func (o LogBatchResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o LogBatchResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o LogBatchResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o LogBatchResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o LogBatchResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o LogBatchResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o LogBatchResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type LogInputs struct {
	// Dataset inputs
	Datasets types.List `tfsdk:"datasets" tf:"optional"`
	// ID of the run to log under
	RunId types.String `tfsdk:"run_id" tf:"optional"`
}

func (newState *LogInputs) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogInputs) {
}

func (newState *LogInputs) SyncEffectiveFieldsDuringRead(existingState LogInputs) {
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
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = LogInputs{}

// Equal implements basetypes.ObjectValuable.
func (o LogInputs) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o LogInputs) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o LogInputs) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o LogInputs) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o LogInputs) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o LogInputs) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o LogInputs) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"datasets": basetypes.ListType{
				ElemType: DatasetInput{}.Type(ctx),
			},
			"run_id": types.StringType,
		},
	}
}

type LogInputsResponse struct {
}

func (newState *LogInputsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogInputsResponse) {
}

func (newState *LogInputsResponse) SyncEffectiveFieldsDuringRead(existingState LogInputsResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = LogInputsResponse{}

// Equal implements basetypes.ObjectValuable.
func (o LogInputsResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o LogInputsResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o LogInputsResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o LogInputsResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o LogInputsResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o LogInputsResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o LogInputsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type LogMetric struct {
	// Name of the metric.
	Key types.String `tfsdk:"key" tf:""`
	// ID of the run under which to log the metric. Must be provided.
	RunId types.String `tfsdk:"run_id" tf:"optional"`
	// [Deprecated, use run_id instead] ID of the run under which to log the
	// metric. This field will be removed in a future MLflow version.
	RunUuid types.String `tfsdk:"run_uuid" tf:"optional"`
	// Step at which to log the metric
	Step types.Int64 `tfsdk:"step" tf:"optional"`
	// Unix timestamp in milliseconds at the time metric was logged.
	Timestamp types.Int64 `tfsdk:"timestamp" tf:""`
	// Double value of the metric being logged.
	Value types.Float64 `tfsdk:"value" tf:""`
}

func (newState *LogMetric) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogMetric) {
}

func (newState *LogMetric) SyncEffectiveFieldsDuringRead(existingState LogMetric) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = LogMetric{}

// Equal implements basetypes.ObjectValuable.
func (o LogMetric) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o LogMetric) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o LogMetric) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o LogMetric) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o LogMetric) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o LogMetric) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o LogMetric) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":       types.StringType,
			"run_id":    types.StringType,
			"run_uuid":  types.StringType,
			"step":      types.Int64Type,
			"timestamp": types.Int64Type,
			"value":     types.Float64Type,
		},
	}
}

type LogMetricResponse struct {
}

func (newState *LogMetricResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogMetricResponse) {
}

func (newState *LogMetricResponse) SyncEffectiveFieldsDuringRead(existingState LogMetricResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = LogMetricResponse{}

// Equal implements basetypes.ObjectValuable.
func (o LogMetricResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o LogMetricResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o LogMetricResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o LogMetricResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o LogMetricResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o LogMetricResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o LogMetricResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type LogModel struct {
	// MLmodel file in json format.
	ModelJson types.String `tfsdk:"model_json" tf:"optional"`
	// ID of the run to log under
	RunId types.String `tfsdk:"run_id" tf:"optional"`
}

func (newState *LogModel) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogModel) {
}

func (newState *LogModel) SyncEffectiveFieldsDuringRead(existingState LogModel) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = LogModel{}

// Equal implements basetypes.ObjectValuable.
func (o LogModel) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o LogModel) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o LogModel) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o LogModel) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o LogModel) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o LogModel) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = LogModelResponse{}

// Equal implements basetypes.ObjectValuable.
func (o LogModelResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o LogModelResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o LogModelResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o LogModelResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o LogModelResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o LogModelResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o LogModelResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type LogParam struct {
	// Name of the param. Maximum size is 255 bytes.
	Key types.String `tfsdk:"key" tf:""`
	// ID of the run under which to log the param. Must be provided.
	RunId types.String `tfsdk:"run_id" tf:"optional"`
	// [Deprecated, use run_id instead] ID of the run under which to log the
	// param. This field will be removed in a future MLflow version.
	RunUuid types.String `tfsdk:"run_uuid" tf:"optional"`
	// String value of the param being logged. Maximum size is 500 bytes.
	Value types.String `tfsdk:"value" tf:""`
}

func (newState *LogParam) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogParam) {
}

func (newState *LogParam) SyncEffectiveFieldsDuringRead(existingState LogParam) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = LogParam{}

// Equal implements basetypes.ObjectValuable.
func (o LogParam) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o LogParam) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o LogParam) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o LogParam) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o LogParam) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o LogParam) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = LogParamResponse{}

// Equal implements basetypes.ObjectValuable.
func (o LogParamResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o LogParamResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o LogParamResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o LogParamResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o LogParamResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o LogParamResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o LogParamResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type Metric struct {
	// Key identifying this metric.
	Key types.String `tfsdk:"key" tf:"optional"`
	// Step at which to log the metric.
	Step types.Int64 `tfsdk:"step" tf:"optional"`
	// The timestamp at which this metric was recorded.
	Timestamp types.Int64 `tfsdk:"timestamp" tf:"optional"`
	// Value associated with this metric.
	Value types.Float64 `tfsdk:"value" tf:"optional"`
}

func (newState *Metric) SyncEffectiveFieldsDuringCreateOrUpdate(plan Metric) {
}

func (newState *Metric) SyncEffectiveFieldsDuringRead(existingState Metric) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = Metric{}

// Equal implements basetypes.ObjectValuable.
func (o Metric) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o Metric) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o Metric) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o Metric) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o Metric) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o Metric) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o Metric) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":       types.StringType,
			"step":      types.Int64Type,
			"timestamp": types.Int64Type,
			"value":     types.Float64Type,
		},
	}
}

type Model struct {
	// Timestamp recorded when this `registered_model` was created.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp" tf:"optional"`
	// Description of this `registered_model`.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Timestamp recorded when metadata for this `registered_model` was last
	// updated.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp" tf:"optional"`
	// Collection of latest model versions for each stage. Only contains models
	// with current `READY` status.
	LatestVersions types.List `tfsdk:"latest_versions" tf:"optional"`
	// Unique name for the model.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Tags: Additional metadata key-value pairs for this `registered_model`.
	Tags types.List `tfsdk:"tags" tf:"optional"`
	// User that created this `registered_model`
	UserId types.String `tfsdk:"user_id" tf:"optional"`
}

func (newState *Model) SyncEffectiveFieldsDuringCreateOrUpdate(plan Model) {
}

func (newState *Model) SyncEffectiveFieldsDuringRead(existingState Model) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = Model{}

// Equal implements basetypes.ObjectValuable.
func (o Model) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o Model) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o Model) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o Model) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o Model) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o Model) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type ModelDatabricks struct {
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp" tf:"optional"`
	// User-specified description for the object.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Unique identifier for the object.
	Id types.String `tfsdk:"id" tf:"optional"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp" tf:"optional"`
	// Array of model versions, each the latest version for its stage.
	LatestVersions types.List `tfsdk:"latest_versions" tf:"optional"`
	// Name of the model.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Permission level of the requesting user on the object. For what is
	// allowed at each level, see [MLflow Model permissions](..).
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
	// Array of tags associated with the model.
	Tags types.List `tfsdk:"tags" tf:"optional"`
	// The username of the user that created the object.
	UserId types.String `tfsdk:"user_id" tf:"optional"`
}

func (newState *ModelDatabricks) SyncEffectiveFieldsDuringCreateOrUpdate(plan ModelDatabricks) {
}

func (newState *ModelDatabricks) SyncEffectiveFieldsDuringRead(existingState ModelDatabricks) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ModelDatabricks{}

// Equal implements basetypes.ObjectValuable.
func (o ModelDatabricks) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ModelDatabricks) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ModelDatabricks) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ModelDatabricks) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ModelDatabricks) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ModelDatabricks) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type ModelTag struct {
	// The tag key.
	Key types.String `tfsdk:"key" tf:"optional"`
	// The tag value.
	Value types.String `tfsdk:"value" tf:"optional"`
}

func (newState *ModelTag) SyncEffectiveFieldsDuringCreateOrUpdate(plan ModelTag) {
}

func (newState *ModelTag) SyncEffectiveFieldsDuringRead(existingState ModelTag) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ModelTag{}

// Equal implements basetypes.ObjectValuable.
func (o ModelTag) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ModelTag) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ModelTag) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ModelTag) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ModelTag) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ModelTag) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp" tf:"optional"`
	// Current stage for this `model_version`.
	CurrentStage types.String `tfsdk:"current_stage" tf:"optional"`
	// Description of this `model_version`.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Timestamp recorded when metadata for this `model_version` was last
	// updated.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp" tf:"optional"`
	// Unique name of the model
	Name types.String `tfsdk:"name" tf:"optional"`
	// MLflow run ID used when creating `model_version`, if `source` was
	// generated by an experiment run stored in MLflow tracking server.
	RunId types.String `tfsdk:"run_id" tf:"optional"`
	// Run Link: Direct link to the run that generated this version
	RunLink types.String `tfsdk:"run_link" tf:"optional"`
	// URI indicating the location of the source model artifacts, used when
	// creating `model_version`
	Source types.String `tfsdk:"source" tf:"optional"`
	// Current status of `model_version`
	Status types.String `tfsdk:"status" tf:"optional"`
	// Details on current `status`, if it is pending or failed.
	StatusMessage types.String `tfsdk:"status_message" tf:"optional"`
	// Tags: Additional metadata key-value pairs for this `model_version`.
	Tags types.List `tfsdk:"tags" tf:"optional"`
	// User that created this `model_version`.
	UserId types.String `tfsdk:"user_id" tf:"optional"`
	// Model's version number.
	Version types.String `tfsdk:"version" tf:"optional"`
}

func (newState *ModelVersion) SyncEffectiveFieldsDuringCreateOrUpdate(plan ModelVersion) {
}

func (newState *ModelVersion) SyncEffectiveFieldsDuringRead(existingState ModelVersion) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ModelVersion{}

// Equal implements basetypes.ObjectValuable.
func (o ModelVersion) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ModelVersion) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ModelVersion) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ModelVersion) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ModelVersion) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ModelVersion) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type ModelVersionDatabricks struct {
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp" tf:"optional"`
	// Stage of the model version. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	CurrentStage types.String `tfsdk:"current_stage" tf:"optional"`
	// User-specified description for the object.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp" tf:"optional"`
	// Name of the model.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Permission level of the requesting user on the object. For what is
	// allowed at each level, see [MLflow Model permissions](..).
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
	// Unique identifier for the MLflow tracking run associated with the source
	// model artifacts.
	RunId types.String `tfsdk:"run_id" tf:"optional"`
	// URL of the run associated with the model artifacts. This field is set at
	// model version creation time only for model versions whose source run is
	// from a tracking server that is different from the registry server.
	RunLink types.String `tfsdk:"run_link" tf:"optional"`
	// URI that indicates the location of the source model artifacts. This is
	// used when creating the model version.
	Source types.String `tfsdk:"source" tf:"optional"`
	// The status of the model version. Valid values are: *
	// `PENDING_REGISTRATION`: Request to register a new model version is
	// pending as server performs background tasks.
	//
	// * `FAILED_REGISTRATION`: Request to register a new model version has
	// failed.
	//
	// * `READY`: Model version is ready for use.
	Status types.String `tfsdk:"status" tf:"optional"`
	// Details on the current status, for example why registration failed.
	StatusMessage types.String `tfsdk:"status_message" tf:"optional"`
	// Array of tags that are associated with the model version.
	Tags types.List `tfsdk:"tags" tf:"optional"`
	// The username of the user that created the object.
	UserId types.String `tfsdk:"user_id" tf:"optional"`
	// Version of the model.
	Version types.String `tfsdk:"version" tf:"optional"`
}

func (newState *ModelVersionDatabricks) SyncEffectiveFieldsDuringCreateOrUpdate(plan ModelVersionDatabricks) {
}

func (newState *ModelVersionDatabricks) SyncEffectiveFieldsDuringRead(existingState ModelVersionDatabricks) {
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
		"tags": reflect.TypeOf(ModelVersionTag{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ModelVersionDatabricks{}

// Equal implements basetypes.ObjectValuable.
func (o ModelVersionDatabricks) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ModelVersionDatabricks) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ModelVersionDatabricks) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ModelVersionDatabricks) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ModelVersionDatabricks) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ModelVersionDatabricks) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o ModelVersionDatabricks) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creation_timestamp":     types.Int64Type,
			"current_stage":          types.StringType,
			"description":            types.StringType,
			"last_updated_timestamp": types.Int64Type,
			"name":                   types.StringType,
			"permission_level":       types.StringType,
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

type ModelVersionTag struct {
	// The tag key.
	Key types.String `tfsdk:"key" tf:"optional"`
	// The tag value.
	Value types.String `tfsdk:"value" tf:"optional"`
}

func (newState *ModelVersionTag) SyncEffectiveFieldsDuringCreateOrUpdate(plan ModelVersionTag) {
}

func (newState *ModelVersionTag) SyncEffectiveFieldsDuringRead(existingState ModelVersionTag) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = ModelVersionTag{}

// Equal implements basetypes.ObjectValuable.
func (o ModelVersionTag) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o ModelVersionTag) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o ModelVersionTag) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o ModelVersionTag) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o ModelVersionTag) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o ModelVersionTag) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type Param struct {
	// Key identifying this param.
	Key types.String `tfsdk:"key" tf:"optional"`
	// Value associated with this param.
	Value types.String `tfsdk:"value" tf:"optional"`
}

func (newState *Param) SyncEffectiveFieldsDuringCreateOrUpdate(plan Param) {
}

func (newState *Param) SyncEffectiveFieldsDuringRead(existingState Param) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = Param{}

// Equal implements basetypes.ObjectValuable.
func (o Param) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o Param) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o Param) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o Param) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o Param) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o Param) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type RegisteredModelAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *RegisteredModelAccessControlRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegisteredModelAccessControlRequest) {
}

func (newState *RegisteredModelAccessControlRequest) SyncEffectiveFieldsDuringRead(existingState RegisteredModelAccessControlRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RegisteredModelAccessControlRequest{}

// Equal implements basetypes.ObjectValuable.
func (o RegisteredModelAccessControlRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RegisteredModelAccessControlRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RegisteredModelAccessControlRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RegisteredModelAccessControlRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RegisteredModelAccessControlRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RegisteredModelAccessControlRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	AllPermissions types.List `tfsdk:"all_permissions" tf:"optional"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *RegisteredModelAccessControlResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegisteredModelAccessControlResponse) {
}

func (newState *RegisteredModelAccessControlResponse) SyncEffectiveFieldsDuringRead(existingState RegisteredModelAccessControlResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RegisteredModelAccessControlResponse{}

// Equal implements basetypes.ObjectValuable.
func (o RegisteredModelAccessControlResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RegisteredModelAccessControlResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RegisteredModelAccessControlResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RegisteredModelAccessControlResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RegisteredModelAccessControlResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RegisteredModelAccessControlResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type RegisteredModelPermission struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *RegisteredModelPermission) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegisteredModelPermission) {
}

func (newState *RegisteredModelPermission) SyncEffectiveFieldsDuringRead(existingState RegisteredModelPermission) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RegisteredModelPermission{}

// Equal implements basetypes.ObjectValuable.
func (o RegisteredModelPermission) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RegisteredModelPermission) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RegisteredModelPermission) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RegisteredModelPermission) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RegisteredModelPermission) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RegisteredModelPermission) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type RegisteredModelPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

func (newState *RegisteredModelPermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegisteredModelPermissions) {
}

func (newState *RegisteredModelPermissions) SyncEffectiveFieldsDuringRead(existingState RegisteredModelPermissions) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RegisteredModelPermissions{}

// Equal implements basetypes.ObjectValuable.
func (o RegisteredModelPermissions) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RegisteredModelPermissions) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RegisteredModelPermissions) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RegisteredModelPermissions) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RegisteredModelPermissions) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RegisteredModelPermissions) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type RegisteredModelPermissionsDescription struct {
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *RegisteredModelPermissionsDescription) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegisteredModelPermissionsDescription) {
}

func (newState *RegisteredModelPermissionsDescription) SyncEffectiveFieldsDuringRead(existingState RegisteredModelPermissionsDescription) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RegisteredModelPermissionsDescription{}

// Equal implements basetypes.ObjectValuable.
func (o RegisteredModelPermissionsDescription) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RegisteredModelPermissionsDescription) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RegisteredModelPermissionsDescription) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RegisteredModelPermissionsDescription) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RegisteredModelPermissionsDescription) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RegisteredModelPermissionsDescription) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`
	// The registered model for which to get or manage permissions.
	RegisteredModelId types.String `tfsdk:"-"`
}

func (newState *RegisteredModelPermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegisteredModelPermissionsRequest) {
}

func (newState *RegisteredModelPermissionsRequest) SyncEffectiveFieldsDuringRead(existingState RegisteredModelPermissionsRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RegisteredModelPermissionsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o RegisteredModelPermissionsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RegisteredModelPermissionsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RegisteredModelPermissionsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RegisteredModelPermissionsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RegisteredModelPermissionsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RegisteredModelPermissionsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type RegistryWebhook struct {
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp" tf:"optional"`
	// User-specified description for the webhook.
	Description types.String `tfsdk:"description" tf:"optional"`
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
	Events types.List `tfsdk:"events" tf:"optional"`

	HttpUrlSpec types.List `tfsdk:"http_url_spec" tf:"optional,object"`
	// Webhook ID
	Id types.String `tfsdk:"id" tf:"optional"`

	JobSpec types.List `tfsdk:"job_spec" tf:"optional,object"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp" tf:"optional"`
	// Name of the model whose events would trigger this webhook.
	ModelName types.String `tfsdk:"model_name" tf:"optional"`
	// Enable or disable triggering the webhook, or put the webhook into test
	// mode. The default is `ACTIVE`: * `ACTIVE`: Webhook is triggered when an
	// associated event happens.
	//
	// * `DISABLED`: Webhook is not triggered.
	//
	// * `TEST_MODE`: Webhook can be triggered through the test endpoint, but is
	// not triggered on a real event.
	Status types.String `tfsdk:"status" tf:"optional"`
}

func (newState *RegistryWebhook) SyncEffectiveFieldsDuringCreateOrUpdate(plan RegistryWebhook) {
}

func (newState *RegistryWebhook) SyncEffectiveFieldsDuringRead(existingState RegistryWebhook) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RegistryWebhook{}

// Equal implements basetypes.ObjectValuable.
func (o RegistryWebhook) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RegistryWebhook) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RegistryWebhook) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RegistryWebhook) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RegistryWebhook) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RegistryWebhook) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
			"http_url_spec": basetypes.ListType{
				ElemType: HttpUrlSpecWithoutSecret{}.Type(ctx),
			},
			"id": types.StringType,
			"job_spec": basetypes.ListType{
				ElemType: JobSpecWithoutSecret{}.Type(ctx),
			},
			"last_updated_timestamp": types.Int64Type,
			"model_name":             types.StringType,
			"status":                 types.StringType,
		},
	}
}

type RejectTransitionRequest struct {
	// User-provided comment on the action.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Name of the model.
	Name types.String `tfsdk:"name" tf:""`
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage types.String `tfsdk:"stage" tf:""`
	// Version of the model.
	Version types.String `tfsdk:"version" tf:""`
}

func (newState *RejectTransitionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan RejectTransitionRequest) {
}

func (newState *RejectTransitionRequest) SyncEffectiveFieldsDuringRead(existingState RejectTransitionRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RejectTransitionRequest{}

// Equal implements basetypes.ObjectValuable.
func (o RejectTransitionRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RejectTransitionRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RejectTransitionRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RejectTransitionRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RejectTransitionRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RejectTransitionRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	// Activity recorded for the action.
	Activity types.List `tfsdk:"activity" tf:"optional,object"`
}

func (newState *RejectTransitionRequestResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RejectTransitionRequestResponse) {
}

func (newState *RejectTransitionRequestResponse) SyncEffectiveFieldsDuringRead(existingState RejectTransitionRequestResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RejectTransitionRequestResponse{}

// Equal implements basetypes.ObjectValuable.
func (o RejectTransitionRequestResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RejectTransitionRequestResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RejectTransitionRequestResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RejectTransitionRequestResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RejectTransitionRequestResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RejectTransitionRequestResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o RejectTransitionRequestResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"activity": basetypes.ListType{
				ElemType: Activity{}.Type(ctx),
			},
		},
	}
}

type RenameModelRequest struct {
	// Registered model unique name identifier.
	Name types.String `tfsdk:"name" tf:""`
	// If provided, updates the name for this `registered_model`.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
}

func (newState *RenameModelRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan RenameModelRequest) {
}

func (newState *RenameModelRequest) SyncEffectiveFieldsDuringRead(existingState RenameModelRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RenameModelRequest{}

// Equal implements basetypes.ObjectValuable.
func (o RenameModelRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RenameModelRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RenameModelRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RenameModelRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RenameModelRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RenameModelRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	RegisteredModel types.List `tfsdk:"registered_model" tf:"optional,object"`
}

func (newState *RenameModelResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RenameModelResponse) {
}

func (newState *RenameModelResponse) SyncEffectiveFieldsDuringRead(existingState RenameModelResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RenameModelResponse{}

// Equal implements basetypes.ObjectValuable.
func (o RenameModelResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RenameModelResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RenameModelResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RenameModelResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RenameModelResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RenameModelResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o RenameModelResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"registered_model": basetypes.ListType{
				ElemType: Model{}.Type(ctx),
			},
		},
	}
}

type RestoreExperiment struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"experiment_id" tf:""`
}

func (newState *RestoreExperiment) SyncEffectiveFieldsDuringCreateOrUpdate(plan RestoreExperiment) {
}

func (newState *RestoreExperiment) SyncEffectiveFieldsDuringRead(existingState RestoreExperiment) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RestoreExperiment{}

// Equal implements basetypes.ObjectValuable.
func (o RestoreExperiment) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RestoreExperiment) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RestoreExperiment) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RestoreExperiment) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RestoreExperiment) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RestoreExperiment) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RestoreExperimentResponse{}

// Equal implements basetypes.ObjectValuable.
func (o RestoreExperimentResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RestoreExperimentResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RestoreExperimentResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RestoreExperimentResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RestoreExperimentResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RestoreExperimentResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o RestoreExperimentResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type RestoreRun struct {
	// ID of the run to restore.
	RunId types.String `tfsdk:"run_id" tf:""`
}

func (newState *RestoreRun) SyncEffectiveFieldsDuringCreateOrUpdate(plan RestoreRun) {
}

func (newState *RestoreRun) SyncEffectiveFieldsDuringRead(existingState RestoreRun) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RestoreRun{}

// Equal implements basetypes.ObjectValuable.
func (o RestoreRun) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RestoreRun) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RestoreRun) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RestoreRun) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RestoreRun) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RestoreRun) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RestoreRunResponse{}

// Equal implements basetypes.ObjectValuable.
func (o RestoreRunResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RestoreRunResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RestoreRunResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RestoreRunResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RestoreRunResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RestoreRunResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o RestoreRunResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type RestoreRuns struct {
	// The ID of the experiment containing the runs to restore.
	ExperimentId types.String `tfsdk:"experiment_id" tf:""`
	// An optional positive integer indicating the maximum number of runs to
	// restore. The maximum allowed value for max_runs is 10000.
	MaxRuns types.Int64 `tfsdk:"max_runs" tf:"optional"`
	// The minimum deletion timestamp in milliseconds since the UNIX epoch for
	// restoring runs. Only runs deleted no earlier than this timestamp are
	// restored.
	MinTimestampMillis types.Int64 `tfsdk:"min_timestamp_millis" tf:""`
}

func (newState *RestoreRuns) SyncEffectiveFieldsDuringCreateOrUpdate(plan RestoreRuns) {
}

func (newState *RestoreRuns) SyncEffectiveFieldsDuringRead(existingState RestoreRuns) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RestoreRuns{}

// Equal implements basetypes.ObjectValuable.
func (o RestoreRuns) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RestoreRuns) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RestoreRuns) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RestoreRuns) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RestoreRuns) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RestoreRuns) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	RunsRestored types.Int64 `tfsdk:"runs_restored" tf:"optional"`
}

func (newState *RestoreRunsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RestoreRunsResponse) {
}

func (newState *RestoreRunsResponse) SyncEffectiveFieldsDuringRead(existingState RestoreRunsResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RestoreRunsResponse{}

// Equal implements basetypes.ObjectValuable.
func (o RestoreRunsResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RestoreRunsResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RestoreRunsResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RestoreRunsResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RestoreRunsResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RestoreRunsResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o RestoreRunsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"runs_restored": types.Int64Type,
		},
	}
}

type Run struct {
	// Run data.
	Data types.List `tfsdk:"data" tf:"optional,object"`
	// Run metadata.
	Info types.List `tfsdk:"info" tf:"optional,object"`
	// Run inputs.
	Inputs types.List `tfsdk:"inputs" tf:"optional,object"`
}

func (newState *Run) SyncEffectiveFieldsDuringCreateOrUpdate(plan Run) {
}

func (newState *Run) SyncEffectiveFieldsDuringRead(existingState Run) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = Run{}

// Equal implements basetypes.ObjectValuable.
func (o Run) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o Run) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o Run) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o Run) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o Run) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o Run) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o Run) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"data": basetypes.ListType{
				ElemType: RunData{}.Type(ctx),
			},
			"info": basetypes.ListType{
				ElemType: RunInfo{}.Type(ctx),
			},
			"inputs": basetypes.ListType{
				ElemType: RunInputs{}.Type(ctx),
			},
		},
	}
}

type RunData struct {
	// Run metrics.
	Metrics types.List `tfsdk:"metrics" tf:"optional"`
	// Run parameters.
	Params types.List `tfsdk:"params" tf:"optional"`
	// Additional metadata key-value pairs.
	Tags types.List `tfsdk:"tags" tf:"optional"`
}

func (newState *RunData) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunData) {
}

func (newState *RunData) SyncEffectiveFieldsDuringRead(existingState RunData) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RunData{}

// Equal implements basetypes.ObjectValuable.
func (o RunData) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RunData) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RunData) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RunData) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RunData) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RunData) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type RunInfo struct {
	// URI of the directory where artifacts should be uploaded. This can be a
	// local path (starting with "/"), or a distributed file system (DFS) path,
	// like `s3://bucket/directory` or `dbfs:/my/directory`. If not set, the
	// local `./mlruns` directory is chosen.
	ArtifactUri types.String `tfsdk:"artifact_uri" tf:"optional"`
	// Unix timestamp of when the run ended in milliseconds.
	EndTime types.Int64 `tfsdk:"end_time" tf:"optional"`
	// The experiment ID.
	ExperimentId types.String `tfsdk:"experiment_id" tf:"optional"`
	// Current life cycle stage of the experiment : OneOf("active", "deleted")
	LifecycleStage types.String `tfsdk:"lifecycle_stage" tf:"optional"`
	// Unique identifier for the run.
	RunId types.String `tfsdk:"run_id" tf:"optional"`
	// [Deprecated, use run_id instead] Unique identifier for the run. This
	// field will be removed in a future MLflow version.
	RunUuid types.String `tfsdk:"run_uuid" tf:"optional"`
	// Unix timestamp of when the run started in milliseconds.
	StartTime types.Int64 `tfsdk:"start_time" tf:"optional"`
	// Current status of the run.
	Status types.String `tfsdk:"status" tf:"optional"`
	// User who initiated the run. This field is deprecated as of MLflow 1.0,
	// and will be removed in a future MLflow release. Use 'mlflow.user' tag
	// instead.
	UserId types.String `tfsdk:"user_id" tf:"optional"`
}

func (newState *RunInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunInfo) {
}

func (newState *RunInfo) SyncEffectiveFieldsDuringRead(existingState RunInfo) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RunInfo{}

// Equal implements basetypes.ObjectValuable.
func (o RunInfo) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RunInfo) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RunInfo) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RunInfo) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RunInfo) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RunInfo) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
			"run_uuid":        types.StringType,
			"start_time":      types.Int64Type,
			"status":          types.StringType,
			"user_id":         types.StringType,
		},
	}
}

type RunInputs struct {
	// Run metrics.
	DatasetInputs types.List `tfsdk:"dataset_inputs" tf:"optional"`
}

func (newState *RunInputs) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunInputs) {
}

func (newState *RunInputs) SyncEffectiveFieldsDuringRead(existingState RunInputs) {
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
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RunInputs{}

// Equal implements basetypes.ObjectValuable.
func (o RunInputs) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RunInputs) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RunInputs) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RunInputs) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RunInputs) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RunInputs) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o RunInputs) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dataset_inputs": basetypes.ListType{
				ElemType: DatasetInput{}.Type(ctx),
			},
		},
	}
}

type RunTag struct {
	// The tag key.
	Key types.String `tfsdk:"key" tf:"optional"`
	// The tag value.
	Value types.String `tfsdk:"value" tf:"optional"`
}

func (newState *RunTag) SyncEffectiveFieldsDuringCreateOrUpdate(plan RunTag) {
}

func (newState *RunTag) SyncEffectiveFieldsDuringRead(existingState RunTag) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = RunTag{}

// Equal implements basetypes.ObjectValuable.
func (o RunTag) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o RunTag) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o RunTag) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o RunTag) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o RunTag) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o RunTag) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	Filter types.String `tfsdk:"filter" tf:"optional"`
	// Maximum number of experiments desired. Max threshold is 3000.
	MaxResults types.Int64 `tfsdk:"max_results" tf:"optional"`
	// List of columns for ordering search results, which can include experiment
	// name and last updated timestamp with an optional "DESC" or "ASC"
	// annotation, where "ASC" is the default. Tiebreaks are done by experiment
	// id DESC.
	OrderBy types.List `tfsdk:"order_by" tf:"optional"`
	// Token indicating the page of experiments to fetch
	PageToken types.String `tfsdk:"page_token" tf:"optional"`
	// Qualifier for type of experiments to be returned. If unspecified, return
	// only active experiments.
	ViewType types.String `tfsdk:"view_type" tf:"optional"`
}

func (newState *SearchExperiments) SyncEffectiveFieldsDuringCreateOrUpdate(plan SearchExperiments) {
}

func (newState *SearchExperiments) SyncEffectiveFieldsDuringRead(existingState SearchExperiments) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = SearchExperiments{}

// Equal implements basetypes.ObjectValuable.
func (o SearchExperiments) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o SearchExperiments) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o SearchExperiments) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o SearchExperiments) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o SearchExperiments) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o SearchExperiments) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type SearchExperimentsResponse struct {
	// Experiments that match the search criteria
	Experiments types.List `tfsdk:"experiments" tf:"optional"`
	// Token that can be used to retrieve the next page of experiments. An empty
	// token means that no more experiments are available for retrieval.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *SearchExperimentsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan SearchExperimentsResponse) {
}

func (newState *SearchExperimentsResponse) SyncEffectiveFieldsDuringRead(existingState SearchExperimentsResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = SearchExperimentsResponse{}

// Equal implements basetypes.ObjectValuable.
func (o SearchExperimentsResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o SearchExperimentsResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o SearchExperimentsResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o SearchExperimentsResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o SearchExperimentsResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o SearchExperimentsResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// Searches model versions
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

func (newState *SearchModelVersionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan SearchModelVersionsRequest) {
}

func (newState *SearchModelVersionsRequest) SyncEffectiveFieldsDuringRead(existingState SearchModelVersionsRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = SearchModelVersionsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o SearchModelVersionsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o SearchModelVersionsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o SearchModelVersionsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o SearchModelVersionsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o SearchModelVersionsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o SearchModelVersionsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type SearchModelVersionsResponse struct {
	// Models that match the search criteria
	ModelVersions types.List `tfsdk:"model_versions" tf:"optional"`
	// Pagination token to request next page of models for the same search
	// query.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *SearchModelVersionsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan SearchModelVersionsResponse) {
}

func (newState *SearchModelVersionsResponse) SyncEffectiveFieldsDuringRead(existingState SearchModelVersionsResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = SearchModelVersionsResponse{}

// Equal implements basetypes.ObjectValuable.
func (o SearchModelVersionsResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o SearchModelVersionsResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o SearchModelVersionsResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o SearchModelVersionsResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o SearchModelVersionsResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o SearchModelVersionsResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// Search models
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

func (newState *SearchModelsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan SearchModelsRequest) {
}

func (newState *SearchModelsRequest) SyncEffectiveFieldsDuringRead(existingState SearchModelsRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = SearchModelsRequest{}

// Equal implements basetypes.ObjectValuable.
func (o SearchModelsRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o SearchModelsRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o SearchModelsRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o SearchModelsRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o SearchModelsRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o SearchModelsRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type SearchModelsResponse struct {
	// Pagination token to request the next page of models.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// Registered Models that match the search criteria.
	RegisteredModels types.List `tfsdk:"registered_models" tf:"optional"`
}

func (newState *SearchModelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan SearchModelsResponse) {
}

func (newState *SearchModelsResponse) SyncEffectiveFieldsDuringRead(existingState SearchModelsResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = SearchModelsResponse{}

// Equal implements basetypes.ObjectValuable.
func (o SearchModelsResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o SearchModelsResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o SearchModelsResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o SearchModelsResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o SearchModelsResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o SearchModelsResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type SearchRuns struct {
	// List of experiment IDs to search over.
	ExperimentIds types.List `tfsdk:"experiment_ids" tf:"optional"`
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
	Filter types.String `tfsdk:"filter" tf:"optional"`
	// Maximum number of runs desired. Max threshold is 50000
	MaxResults types.Int64 `tfsdk:"max_results" tf:"optional"`
	// List of columns to be ordered by, including attributes, params, metrics,
	// and tags with an optional "DESC" or "ASC" annotation, where "ASC" is the
	// default. Example: ["params.input DESC", "metrics.alpha ASC",
	// "metrics.rmse"] Tiebreaks are done by start_time DESC followed by run_id
	// for runs with the same start time (and this is the default ordering
	// criterion if order_by is not provided).
	OrderBy types.List `tfsdk:"order_by" tf:"optional"`
	// Token for the current page of runs.
	PageToken types.String `tfsdk:"page_token" tf:"optional"`
	// Whether to display only active, only deleted, or all runs. Defaults to
	// only active runs.
	RunViewType types.String `tfsdk:"run_view_type" tf:"optional"`
}

func (newState *SearchRuns) SyncEffectiveFieldsDuringCreateOrUpdate(plan SearchRuns) {
}

func (newState *SearchRuns) SyncEffectiveFieldsDuringRead(existingState SearchRuns) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = SearchRuns{}

// Equal implements basetypes.ObjectValuable.
func (o SearchRuns) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o SearchRuns) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o SearchRuns) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o SearchRuns) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o SearchRuns) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o SearchRuns) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type SearchRunsResponse struct {
	// Token for the next page of runs.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// Runs that match the search criteria.
	Runs types.List `tfsdk:"runs" tf:"optional"`
}

func (newState *SearchRunsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan SearchRunsResponse) {
}

func (newState *SearchRunsResponse) SyncEffectiveFieldsDuringRead(existingState SearchRunsResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = SearchRunsResponse{}

// Equal implements basetypes.ObjectValuable.
func (o SearchRunsResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o SearchRunsResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o SearchRunsResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o SearchRunsResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o SearchRunsResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o SearchRunsResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type SetExperimentTag struct {
	// ID of the experiment under which to log the tag. Must be provided.
	ExperimentId types.String `tfsdk:"experiment_id" tf:""`
	// Name of the tag. Maximum size depends on storage backend. All storage
	// backends are guaranteed to support key values up to 250 bytes in size.
	Key types.String `tfsdk:"key" tf:""`
	// String value of the tag being logged. Maximum size depends on storage
	// backend. All storage backends are guaranteed to support key values up to
	// 5000 bytes in size.
	Value types.String `tfsdk:"value" tf:""`
}

func (newState *SetExperimentTag) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetExperimentTag) {
}

func (newState *SetExperimentTag) SyncEffectiveFieldsDuringRead(existingState SetExperimentTag) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = SetExperimentTag{}

// Equal implements basetypes.ObjectValuable.
func (o SetExperimentTag) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o SetExperimentTag) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o SetExperimentTag) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o SetExperimentTag) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o SetExperimentTag) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o SetExperimentTag) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = SetExperimentTagResponse{}

// Equal implements basetypes.ObjectValuable.
func (o SetExperimentTagResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o SetExperimentTagResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o SetExperimentTagResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o SetExperimentTagResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o SetExperimentTagResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o SetExperimentTagResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o SetExperimentTagResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type SetModelTagRequest struct {
	// Name of the tag. Maximum size depends on storage backend. If a tag with
	// this name already exists, its preexisting value will be replaced by the
	// specified `value`. All storage backends are guaranteed to support key
	// values up to 250 bytes in size.
	Key types.String `tfsdk:"key" tf:""`
	// Unique name of the model.
	Name types.String `tfsdk:"name" tf:""`
	// String value of the tag being logged. Maximum size depends on storage
	// backend. All storage backends are guaranteed to support key values up to
	// 5000 bytes in size.
	Value types.String `tfsdk:"value" tf:""`
}

func (newState *SetModelTagRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetModelTagRequest) {
}

func (newState *SetModelTagRequest) SyncEffectiveFieldsDuringRead(existingState SetModelTagRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = SetModelTagRequest{}

// Equal implements basetypes.ObjectValuable.
func (o SetModelTagRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o SetModelTagRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o SetModelTagRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o SetModelTagRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o SetModelTagRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o SetModelTagRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = SetModelTagResponse{}

// Equal implements basetypes.ObjectValuable.
func (o SetModelTagResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o SetModelTagResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o SetModelTagResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o SetModelTagResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o SetModelTagResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o SetModelTagResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	Key types.String `tfsdk:"key" tf:""`
	// Unique name of the model.
	Name types.String `tfsdk:"name" tf:""`
	// String value of the tag being logged. Maximum size depends on storage
	// backend. All storage backends are guaranteed to support key values up to
	// 5000 bytes in size.
	Value types.String `tfsdk:"value" tf:""`
	// Model version number.
	Version types.String `tfsdk:"version" tf:""`
}

func (newState *SetModelVersionTagRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetModelVersionTagRequest) {
}

func (newState *SetModelVersionTagRequest) SyncEffectiveFieldsDuringRead(existingState SetModelVersionTagRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = SetModelVersionTagRequest{}

// Equal implements basetypes.ObjectValuable.
func (o SetModelVersionTagRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o SetModelVersionTagRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o SetModelVersionTagRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o SetModelVersionTagRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o SetModelVersionTagRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o SetModelVersionTagRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = SetModelVersionTagResponse{}

// Equal implements basetypes.ObjectValuable.
func (o SetModelVersionTagResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o SetModelVersionTagResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o SetModelVersionTagResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o SetModelVersionTagResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o SetModelVersionTagResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o SetModelVersionTagResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o SetModelVersionTagResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type SetTag struct {
	// Name of the tag. Maximum size depends on storage backend. All storage
	// backends are guaranteed to support key values up to 250 bytes in size.
	Key types.String `tfsdk:"key" tf:""`
	// ID of the run under which to log the tag. Must be provided.
	RunId types.String `tfsdk:"run_id" tf:"optional"`
	// [Deprecated, use run_id instead] ID of the run under which to log the
	// tag. This field will be removed in a future MLflow version.
	RunUuid types.String `tfsdk:"run_uuid" tf:"optional"`
	// String value of the tag being logged. Maximum size depends on storage
	// backend. All storage backends are guaranteed to support key values up to
	// 5000 bytes in size.
	Value types.String `tfsdk:"value" tf:""`
}

func (newState *SetTag) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetTag) {
}

func (newState *SetTag) SyncEffectiveFieldsDuringRead(existingState SetTag) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = SetTag{}

// Equal implements basetypes.ObjectValuable.
func (o SetTag) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o SetTag) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o SetTag) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o SetTag) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o SetTag) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o SetTag) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = SetTagResponse{}

// Equal implements basetypes.ObjectValuable.
func (o SetTagResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o SetTagResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o SetTagResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o SetTagResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o SetTagResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o SetTagResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o SetTagResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Test webhook response object.
type TestRegistryWebhook struct {
	// Body of the response from the webhook URL
	Body types.String `tfsdk:"body" tf:"optional"`
	// Status code returned by the webhook URL
	StatusCode types.Int64 `tfsdk:"status_code" tf:"optional"`
}

func (newState *TestRegistryWebhook) SyncEffectiveFieldsDuringCreateOrUpdate(plan TestRegistryWebhook) {
}

func (newState *TestRegistryWebhook) SyncEffectiveFieldsDuringRead(existingState TestRegistryWebhook) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TestRegistryWebhook.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TestRegistryWebhook) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = TestRegistryWebhook{}

// Equal implements basetypes.ObjectValuable.
func (o TestRegistryWebhook) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o TestRegistryWebhook) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o TestRegistryWebhook) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o TestRegistryWebhook) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o TestRegistryWebhook) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o TestRegistryWebhook) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o TestRegistryWebhook) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"body":        types.StringType,
			"status_code": types.Int64Type,
		},
	}
}

type TestRegistryWebhookRequest struct {
	// If `event` is specified, the test trigger uses the specified event. If
	// `event` is not specified, the test trigger uses a randomly chosen event
	// associated with the webhook.
	Event types.String `tfsdk:"event" tf:"optional"`
	// Webhook ID
	Id types.String `tfsdk:"id" tf:""`
}

func (newState *TestRegistryWebhookRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan TestRegistryWebhookRequest) {
}

func (newState *TestRegistryWebhookRequest) SyncEffectiveFieldsDuringRead(existingState TestRegistryWebhookRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = TestRegistryWebhookRequest{}

// Equal implements basetypes.ObjectValuable.
func (o TestRegistryWebhookRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o TestRegistryWebhookRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o TestRegistryWebhookRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o TestRegistryWebhookRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o TestRegistryWebhookRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o TestRegistryWebhookRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	// Test webhook response object.
	Webhook types.List `tfsdk:"webhook" tf:"optional,object"`
}

func (newState *TestRegistryWebhookResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan TestRegistryWebhookResponse) {
}

func (newState *TestRegistryWebhookResponse) SyncEffectiveFieldsDuringRead(existingState TestRegistryWebhookResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TestRegistryWebhookResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TestRegistryWebhookResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"webhook": reflect.TypeOf(TestRegistryWebhook{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = TestRegistryWebhookResponse{}

// Equal implements basetypes.ObjectValuable.
func (o TestRegistryWebhookResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o TestRegistryWebhookResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o TestRegistryWebhookResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o TestRegistryWebhookResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o TestRegistryWebhookResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o TestRegistryWebhookResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o TestRegistryWebhookResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"webhook": basetypes.ListType{
				ElemType: TestRegistryWebhook{}.Type(ctx),
			},
		},
	}
}

type TransitionModelVersionStageDatabricks struct {
	// Specifies whether to archive all current model versions in the target
	// stage.
	ArchiveExistingVersions types.Bool `tfsdk:"archive_existing_versions" tf:""`
	// User-provided comment on the action.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Name of the model.
	Name types.String `tfsdk:"name" tf:""`
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage types.String `tfsdk:"stage" tf:""`
	// Version of the model.
	Version types.String `tfsdk:"version" tf:""`
}

func (newState *TransitionModelVersionStageDatabricks) SyncEffectiveFieldsDuringCreateOrUpdate(plan TransitionModelVersionStageDatabricks) {
}

func (newState *TransitionModelVersionStageDatabricks) SyncEffectiveFieldsDuringRead(existingState TransitionModelVersionStageDatabricks) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = TransitionModelVersionStageDatabricks{}

// Equal implements basetypes.ObjectValuable.
func (o TransitionModelVersionStageDatabricks) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o TransitionModelVersionStageDatabricks) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o TransitionModelVersionStageDatabricks) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o TransitionModelVersionStageDatabricks) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o TransitionModelVersionStageDatabricks) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o TransitionModelVersionStageDatabricks) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// Transition request details.
type TransitionRequest struct {
	// Array of actions on the activity allowed for the current viewer.
	AvailableActions types.List `tfsdk:"available_actions" tf:"optional"`
	// User-provided comment associated with the transition request.
	Comment types.String `tfsdk:"comment" tf:"optional"`
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp" tf:"optional"`
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
	ToStage types.String `tfsdk:"to_stage" tf:"optional"`
	// The username of the user that created the object.
	UserId types.String `tfsdk:"user_id" tf:"optional"`
}

func (newState *TransitionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan TransitionRequest) {
}

func (newState *TransitionRequest) SyncEffectiveFieldsDuringRead(existingState TransitionRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = TransitionRequest{}

// Equal implements basetypes.ObjectValuable.
func (o TransitionRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o TransitionRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o TransitionRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o TransitionRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o TransitionRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o TransitionRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

type TransitionStageResponse struct {
	ModelVersion types.List `tfsdk:"model_version" tf:"optional,object"`
}

func (newState *TransitionStageResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan TransitionStageResponse) {
}

func (newState *TransitionStageResponse) SyncEffectiveFieldsDuringRead(existingState TransitionStageResponse) {
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
		"model_version": reflect.TypeOf(ModelVersionDatabricks{}),
	}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = TransitionStageResponse{}

// Equal implements basetypes.ObjectValuable.
func (o TransitionStageResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o TransitionStageResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o TransitionStageResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o TransitionStageResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o TransitionStageResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o TransitionStageResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o TransitionStageResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_version": basetypes.ListType{
				ElemType: ModelVersionDatabricks{}.Type(ctx),
			},
		},
	}
}

type UpdateComment struct {
	// User-provided comment on the action.
	Comment types.String `tfsdk:"comment" tf:""`
	// Unique identifier of an activity
	Id types.String `tfsdk:"id" tf:""`
}

func (newState *UpdateComment) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateComment) {
}

func (newState *UpdateComment) SyncEffectiveFieldsDuringRead(existingState UpdateComment) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UpdateComment{}

// Equal implements basetypes.ObjectValuable.
func (o UpdateComment) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UpdateComment) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UpdateComment) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UpdateComment) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UpdateComment) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UpdateComment) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
	// Comment details.
	Comment types.List `tfsdk:"comment" tf:"optional,object"`
}

func (newState *UpdateCommentResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateCommentResponse) {
}

func (newState *UpdateCommentResponse) SyncEffectiveFieldsDuringRead(existingState UpdateCommentResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UpdateCommentResponse{}

// Equal implements basetypes.ObjectValuable.
func (o UpdateCommentResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UpdateCommentResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UpdateCommentResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UpdateCommentResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UpdateCommentResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UpdateCommentResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o UpdateCommentResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment": basetypes.ListType{
				ElemType: CommentObject{}.Type(ctx),
			},
		},
	}
}

type UpdateExperiment struct {
	// ID of the associated experiment.
	ExperimentId types.String `tfsdk:"experiment_id" tf:""`
	// If provided, the experiment's name is changed to the new name. The new
	// name must be unique.
	NewName types.String `tfsdk:"new_name" tf:"optional"`
}

func (newState *UpdateExperiment) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateExperiment) {
}

func (newState *UpdateExperiment) SyncEffectiveFieldsDuringRead(existingState UpdateExperiment) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UpdateExperiment{}

// Equal implements basetypes.ObjectValuable.
func (o UpdateExperiment) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UpdateExperiment) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UpdateExperiment) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UpdateExperiment) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UpdateExperiment) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UpdateExperiment) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UpdateExperimentResponse{}

// Equal implements basetypes.ObjectValuable.
func (o UpdateExperimentResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UpdateExperimentResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UpdateExperimentResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UpdateExperimentResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UpdateExperimentResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UpdateExperimentResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o UpdateExperimentResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateModelRequest struct {
	// If provided, updates the description for this `registered_model`.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Registered model unique name identifier.
	Name types.String `tfsdk:"name" tf:""`
}

func (newState *UpdateModelRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateModelRequest) {
}

func (newState *UpdateModelRequest) SyncEffectiveFieldsDuringRead(existingState UpdateModelRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UpdateModelRequest{}

// Equal implements basetypes.ObjectValuable.
func (o UpdateModelRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UpdateModelRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UpdateModelRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UpdateModelRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UpdateModelRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UpdateModelRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
}

func (newState *UpdateModelResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateModelResponse) {
}

func (newState *UpdateModelResponse) SyncEffectiveFieldsDuringRead(existingState UpdateModelResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateModelResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateModelResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UpdateModelResponse{}

// Equal implements basetypes.ObjectValuable.
func (o UpdateModelResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UpdateModelResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UpdateModelResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UpdateModelResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UpdateModelResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UpdateModelResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o UpdateModelResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateModelVersionRequest struct {
	// If provided, updates the description for this `registered_model`.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Name of the registered model
	Name types.String `tfsdk:"name" tf:""`
	// Model version number
	Version types.String `tfsdk:"version" tf:""`
}

func (newState *UpdateModelVersionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateModelVersionRequest) {
}

func (newState *UpdateModelVersionRequest) SyncEffectiveFieldsDuringRead(existingState UpdateModelVersionRequest) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UpdateModelVersionRequest{}

// Equal implements basetypes.ObjectValuable.
func (o UpdateModelVersionRequest) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UpdateModelVersionRequest) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UpdateModelVersionRequest) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UpdateModelVersionRequest) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UpdateModelVersionRequest) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UpdateModelVersionRequest) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
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
}

func (newState *UpdateModelVersionResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateModelVersionResponse) {
}

func (newState *UpdateModelVersionResponse) SyncEffectiveFieldsDuringRead(existingState UpdateModelVersionResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateModelVersionResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateModelVersionResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UpdateModelVersionResponse{}

// Equal implements basetypes.ObjectValuable.
func (o UpdateModelVersionResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UpdateModelVersionResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UpdateModelVersionResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UpdateModelVersionResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UpdateModelVersionResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UpdateModelVersionResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o UpdateModelVersionResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateRegistryWebhook struct {
	// User-specified description for the webhook.
	Description types.String `tfsdk:"description" tf:"optional"`
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
	Events types.List `tfsdk:"events" tf:"optional"`

	HttpUrlSpec types.List `tfsdk:"http_url_spec" tf:"optional,object"`
	// Webhook ID
	Id types.String `tfsdk:"id" tf:""`

	JobSpec types.List `tfsdk:"job_spec" tf:"optional,object"`
	// Enable or disable triggering the webhook, or put the webhook into test
	// mode. The default is `ACTIVE`: * `ACTIVE`: Webhook is triggered when an
	// associated event happens.
	//
	// * `DISABLED`: Webhook is not triggered.
	//
	// * `TEST_MODE`: Webhook can be triggered through the test endpoint, but is
	// not triggered on a real event.
	Status types.String `tfsdk:"status" tf:"optional"`
}

func (newState *UpdateRegistryWebhook) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateRegistryWebhook) {
}

func (newState *UpdateRegistryWebhook) SyncEffectiveFieldsDuringRead(existingState UpdateRegistryWebhook) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UpdateRegistryWebhook{}

// Equal implements basetypes.ObjectValuable.
func (o UpdateRegistryWebhook) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UpdateRegistryWebhook) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UpdateRegistryWebhook) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UpdateRegistryWebhook) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UpdateRegistryWebhook) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UpdateRegistryWebhook) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o UpdateRegistryWebhook) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description": types.StringType,
			"events": basetypes.ListType{
				ElemType: types.StringType,
			},
			"http_url_spec": basetypes.ListType{
				ElemType: HttpUrlSpec{}.Type(ctx),
			},
			"id": types.StringType,
			"job_spec": basetypes.ListType{
				ElemType: JobSpec{}.Type(ctx),
			},
			"status": types.StringType,
		},
	}
}

type UpdateRun struct {
	// Unix timestamp in milliseconds of when the run ended.
	EndTime types.Int64 `tfsdk:"end_time" tf:"optional"`
	// ID of the run to update. Must be provided.
	RunId types.String `tfsdk:"run_id" tf:"optional"`
	// [Deprecated, use run_id instead] ID of the run to update.. This field
	// will be removed in a future MLflow version.
	RunUuid types.String `tfsdk:"run_uuid" tf:"optional"`
	// Updated status of the run.
	Status types.String `tfsdk:"status" tf:"optional"`
}

func (newState *UpdateRun) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateRun) {
}

func (newState *UpdateRun) SyncEffectiveFieldsDuringRead(existingState UpdateRun) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UpdateRun{}

// Equal implements basetypes.ObjectValuable.
func (o UpdateRun) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UpdateRun) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UpdateRun) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UpdateRun) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UpdateRun) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UpdateRun) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o UpdateRun) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"end_time": types.Int64Type,
			"run_id":   types.StringType,
			"run_uuid": types.StringType,
			"status":   types.StringType,
		},
	}
}

type UpdateRunResponse struct {
	// Updated metadata of the run.
	RunInfo types.List `tfsdk:"run_info" tf:"optional,object"`
}

func (newState *UpdateRunResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateRunResponse) {
}

func (newState *UpdateRunResponse) SyncEffectiveFieldsDuringRead(existingState UpdateRunResponse) {
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

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UpdateRunResponse{}

// Equal implements basetypes.ObjectValuable.
func (o UpdateRunResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UpdateRunResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UpdateRunResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UpdateRunResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UpdateRunResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UpdateRunResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o UpdateRunResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run_info": basetypes.ListType{
				ElemType: RunInfo{}.Type(ctx),
			},
		},
	}
}

type UpdateWebhookResponse struct {
}

func (newState *UpdateWebhookResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateWebhookResponse) {
}

func (newState *UpdateWebhookResponse) SyncEffectiveFieldsDuringRead(existingState UpdateWebhookResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateWebhookResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateWebhookResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types also implement the ObjectValuable interface, so they can be used directly as objects
// and as elements in lists and maps.
var _ basetypes.ObjectValuable = UpdateWebhookResponse{}

// Equal implements basetypes.ObjectValuable.
func (o UpdateWebhookResponse) Equal(v attr.Value) bool {
	ov, d := o.ToObjectValue(context.Background())
	if d.HasError() {
		return false
	}
	return ov.Equal(v)
}

// IsNull implements basetypes.ObjectValuable.
func (o UpdateWebhookResponse) IsNull() bool {
	// TF SDK structures are never null.
	return false
}

// IsUnknown implements basetypes.ObjectValuable.
func (o UpdateWebhookResponse) IsUnknown() bool {
	// TF SDK structures are never unknown.
	return false
}

// String implements basetypes.ObjectValuable.
func (o UpdateWebhookResponse) String() string {
	return fmt.Sprintf("%#v", o)
}

// ToObjectValue implements basetypes.ObjectValuable.
func (o UpdateWebhookResponse) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		o,
	)
}

// ToTerraformValue implements basetypes.ObjectValuable.
func (o UpdateWebhookResponse) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	ov, d := o.ToObjectValue(ctx)
	if d.HasError() {
		return tftypes.Value{}, fmt.Errorf("error converting to object value: %s", pluginfwcommon.DiagToString(d))
	}
	return ov.ToTerraformValue(ctx)
}

// Type implements basetypes.ObjectValuable.
func (o UpdateWebhookResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// An action that a user (with sufficient permissions) could take on an
// activity. Valid values are: * `APPROVE_TRANSITION_REQUEST`: Approve a
// transition request
//
// * `REJECT_TRANSITION_REQUEST`: Reject a transition request
//
// * `CANCEL_TRANSITION_REQUEST`: Cancel (delete) a transition request

// Type of activity. Valid values are: * `APPLIED_TRANSITION`: User applied the
// corresponding stage transition.
//
// * `REQUESTED_TRANSITION`: User requested the corresponding stage transition.
//
// * `CANCELLED_REQUEST`: User cancelled an existing transition request.
//
// * `APPROVED_REQUEST`: User approved the corresponding stage transition.
//
// * `REJECTED_REQUEST`: User rejected the coressponding stage transition.
//
// * `SYSTEM_TRANSITION`: For events performed as a side effect, such as
// archiving existing model versions in a stage.

// An action that a user (with sufficient permissions) could take on a comment.
// Valid values are: * `EDIT_COMMENT`: Edit the comment
//
// * `DELETE_COMMENT`: Delete the comment

// Permission level

// Current status of `model_version`

// Permission level of the requesting user on the object. For what is allowed at
// each level, see [MLflow Model permissions](..).

// Permission level

// Enable or disable triggering the webhook, or put the webhook into test mode.
// The default is `ACTIVE`: * `ACTIVE`: Webhook is triggered when an associated
// event happens.
//
// * `DISABLED`: Webhook is not triggered.
//
// * `TEST_MODE`: Webhook can be triggered through the test endpoint, but is not
// triggered on a real event.

// Current status of the run.

// Qualifier for type of experiments to be returned. If unspecified, return only
// active experiments.

// Whether to display only active, only deleted, or all runs. Defaults to only
// active runs.

// Stage of the model version. Valid values are:
//
// * `None`: The initial stage of a model version.
//
// * `Staging`: Staging or pre-production stage.
//
// * `Production`: Production stage.
//
// * `Archived`: Archived stage.

// The status of the model version. Valid values are: * `PENDING_REGISTRATION`:
// Request to register a new model version is pending as server performs
// background tasks.
//
// * `FAILED_REGISTRATION`: Request to register a new model version has failed.
//
// * `READY`: Model version is ready for use.

// Updated status of the run.
