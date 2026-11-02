// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package computemigrations_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// A request to approve migration execution at the MIGRATION_REVIEW gate.
type ApproveMigrationExecutionRequest struct {
	// Optional free-text feedback or instructions the user provides along with
	// the approval, carried into the next stage of the migration.
	AdditionalFeedback types.String `tfsdk:"additional_feedback"`
	// Optional resource version for optimistic concurrency control. If
	// supplied, it must match the current server value.
	Etag types.String `tfsdk:"etag"`
	// The ID of the migration to approve for execution.
	MigrationId types.String `tfsdk:"-"`
}

func (to *ApproveMigrationExecutionRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ApproveMigrationExecutionRequest) {
}

func (to *ApproveMigrationExecutionRequest) SyncFieldsDuringRead(ctx context.Context, from ApproveMigrationExecutionRequest) {
}

func (m ApproveMigrationExecutionRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["additional_feedback"] = attrs["additional_feedback"].SetOptional()
	attrs["etag"] = attrs["etag"].SetOptional()
	attrs["migration_id"] = attrs["migration_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ApproveMigrationExecutionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ApproveMigrationExecutionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ApproveMigrationExecutionRequest
// only implements ToObjectValue() and Type().
func (m ApproveMigrationExecutionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"additional_feedback": m.AdditionalFeedback,
			"etag":                m.Etag,
			"migration_id":        m.MigrationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ApproveMigrationExecutionRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"additional_feedback": types.StringType,
			"etag":                types.StringType,
			"migration_id":        types.StringType,
		},
	}
}

// A request to approve the working copy at the TRANSLATION_READY gate.
type ApproveMigrationTranslationRequest struct {
	// Optional free-text feedback or instructions the user provides along with
	// the approval, carried into the next stage of the migration.
	AdditionalFeedback types.String `tfsdk:"additional_feedback"`
	// Optional resource version for optimistic concurrency control. If
	// supplied, it must match the current server value.
	Etag types.String `tfsdk:"etag"`
	// The ID of the migration to approve.
	MigrationId types.String `tfsdk:"-"`
}

func (to *ApproveMigrationTranslationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ApproveMigrationTranslationRequest) {
}

func (to *ApproveMigrationTranslationRequest) SyncFieldsDuringRead(ctx context.Context, from ApproveMigrationTranslationRequest) {
}

func (m ApproveMigrationTranslationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["additional_feedback"] = attrs["additional_feedback"].SetOptional()
	attrs["etag"] = attrs["etag"].SetOptional()
	attrs["migration_id"] = attrs["migration_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ApproveMigrationTranslationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ApproveMigrationTranslationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ApproveMigrationTranslationRequest
// only implements ToObjectValue() and Type().
func (m ApproveMigrationTranslationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"additional_feedback": m.AdditionalFeedback,
			"etag":                m.Etag,
			"migration_id":        m.MigrationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ApproveMigrationTranslationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"additional_feedback": types.StringType,
			"etag":                types.StringType,
			"migration_id":        types.StringType,
		},
	}
}

// Configuration for the baking window (post-execution production monitoring).
type BakingConfig struct {
	// Number of successful production runs required before the migration
	// reaches MIGRATION_COMPLETE. Default applied by the server (configurable).
	RunCount types.Int64 `tfsdk:"run_count"`
	// Maximum baking window duration in hours. If the required run_count has
	// not been reached within this window, the migration transitions to FAILED.
	// Default applied by the server (configurable).
	TimeWindowHours types.Int64 `tfsdk:"time_window_hours"`
}

func (to *BakingConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BakingConfig) {
}

func (to *BakingConfig) SyncFieldsDuringRead(ctx context.Context, from BakingConfig) {
}

func (m BakingConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["run_count"] = attrs["run_count"].SetOptional()
	attrs["time_window_hours"] = attrs["time_window_hours"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BakingConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m BakingConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BakingConfig
// only implements ToObjectValue() and Type().
func (m BakingConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"run_count":         m.RunCount,
			"time_window_hours": m.TimeWindowHours,
		})
}

// Type implements basetypes.ObjectValuable.
func (m BakingConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"run_count":         types.Int64Type,
			"time_window_hours": types.Int64Type,
		},
	}
}

// Summary of baking metrics after the baking window completes.
type BakingReport struct {
	// Number of successful runs.
	SuccessfulRuns types.Int64 `tfsdk:"successful_runs"`
	// Number of runs observed during the baking window.
	TotalRuns types.Int64 `tfsdk:"total_runs"`
}

func (to *BakingReport) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BakingReport) {
}

func (to *BakingReport) SyncFieldsDuringRead(ctx context.Context, from BakingReport) {
}

func (m BakingReport) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["successful_runs"] = attrs["successful_runs"].SetComputed()
	attrs["total_runs"] = attrs["total_runs"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BakingReport.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m BakingReport) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BakingReport
// only implements ToObjectValue() and Type().
func (m BakingReport) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"successful_runs": m.SuccessfulRuns,
			"total_runs":      m.TotalRuns,
		})
}

// Type implements basetypes.ObjectValuable.
func (m BakingReport) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"successful_runs": types.Int64Type,
			"total_runs":      types.Int64Type,
		},
	}
}

// A request to cancel a migration from any non-terminal state.
type CancelMigrationRequest struct {
	// Optional free-text reason for the cancellation.
	CancellationReason types.String `tfsdk:"cancellation_reason"`
	// Optional resource version for optimistic concurrency control. If
	// supplied, it must match the current server value.
	Etag types.String `tfsdk:"etag"`
	// The ID of the migration to cancel.
	MigrationId types.String `tfsdk:"-"`
}

func (to *CancelMigrationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CancelMigrationRequest) {
}

func (to *CancelMigrationRequest) SyncFieldsDuringRead(ctx context.Context, from CancelMigrationRequest) {
}

func (m CancelMigrationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cancellation_reason"] = attrs["cancellation_reason"].SetOptional()
	attrs["etag"] = attrs["etag"].SetOptional()
	attrs["migration_id"] = attrs["migration_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CancelMigrationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CancelMigrationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CancelMigrationRequest
// only implements ToObjectValue() and Type().
func (m CancelMigrationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cancellation_reason": m.CancellationReason,
			"etag":                m.Etag,
			"migration_id":        m.MigrationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CancelMigrationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cancellation_reason": types.StringType,
			"etag":                types.StringType,
			"migration_id":        types.StringType,
		},
	}
}

type CreateMigrationRequest struct {
	// The migration to create. The server generates `migration_id` and sets
	// `state` to PLANNING_PENDING. Required fields: `asset_type`, `asset_uri`.
	Migration types.Object `tfsdk:"migration"`
	// Client-generated request ID for idempotent request deduplication. The
	// server de-duplicates requests with the same request_id for at least 24
	// hours, returning the existing migration instead of creating a duplicate.
	RequestId types.String `tfsdk:"-"`
}

func (to *CreateMigrationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateMigrationRequest) {
	if !from.Migration.IsNull() && !from.Migration.IsUnknown() {
		if toMigration, ok := to.GetMigration(ctx); ok {
			if fromMigration, ok := from.GetMigration(ctx); ok {
				// Recursively sync the fields of Migration
				toMigration.SyncFieldsDuringCreateOrUpdate(ctx, fromMigration)
				to.SetMigration(ctx, toMigration)
			}
		}
	}
}

func (to *CreateMigrationRequest) SyncFieldsDuringRead(ctx context.Context, from CreateMigrationRequest) {
	if !from.Migration.IsNull() && !from.Migration.IsUnknown() {
		if toMigration, ok := to.GetMigration(ctx); ok {
			if fromMigration, ok := from.GetMigration(ctx); ok {
				toMigration.SyncFieldsDuringRead(ctx, fromMigration)
				to.SetMigration(ctx, toMigration)
			}
		}
	}
}

func (m CreateMigrationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["migration"] = attrs["migration"].SetRequired()
	attrs["request_id"] = attrs["request_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateMigrationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateMigrationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"migration": reflect.TypeOf(Migration{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateMigrationRequest
// only implements ToObjectValue() and Type().
func (m CreateMigrationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"migration":  m.Migration,
			"request_id": m.RequestId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateMigrationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"migration":  Migration{}.Type(ctx),
			"request_id": types.StringType,
		},
	}
}

// GetMigration returns the value of the Migration field in CreateMigrationRequest as
// a Migration value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateMigrationRequest) GetMigration(ctx context.Context) (Migration, bool) {
	var e Migration
	if m.Migration.IsNull() || m.Migration.IsUnknown() {
		return e, false
	}
	var v Migration
	d := m.Migration.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMigration sets the value of the Migration field in CreateMigrationRequest.
func (m *CreateMigrationRequest) SetMigration(ctx context.Context, v Migration) {
	vs := v.ToObjectValue(ctx)
	m.Migration = vs
}

type DeleteMigrationRequest struct {
	// If nonempty, must match the current resource etag or the request fails
	// with ABORTED.
	Etag types.String `tfsdk:"-"`
	// The ID of the migration to cancel.
	MigrationId types.String `tfsdk:"-"`
}

func (to *DeleteMigrationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteMigrationRequest) {
}

func (to *DeleteMigrationRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteMigrationRequest) {
}

func (m DeleteMigrationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["migration_id"] = attrs["migration_id"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteMigrationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteMigrationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteMigrationRequest
// only implements ToObjectValue() and Type().
func (m DeleteMigrationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag":         m.Etag,
			"migration_id": m.MigrationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteMigrationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag":         types.StringType,
			"migration_id": types.StringType,
		},
	}
}

type GetMigrationRequest struct {
	// The ID of the migration to retrieve.
	MigrationId types.String `tfsdk:"-"`
}

func (to *GetMigrationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetMigrationRequest) {
}

func (to *GetMigrationRequest) SyncFieldsDuringRead(ctx context.Context, from GetMigrationRequest) {
}

func (m GetMigrationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["migration_id"] = attrs["migration_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetMigrationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetMigrationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetMigrationRequest
// only implements ToObjectValue() and Type().
func (m GetMigrationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"migration_id": m.MigrationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetMigrationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"migration_id": types.StringType,
		},
	}
}

type ListMigrationsRequest struct {
	// Filter by asset type.
	AssetType types.String `tfsdk:"-"`
	// Maximum number of migrations to return. Default: 50, maximum: 1000.
	PageSize types.Int64 `tfsdk:"-"`
	// Page token from a previous ListMigrations call. When paginating, all
	// other parameters must match the call that provided the page token.
	PageToken types.String `tfsdk:"-"`
	// Filter by migration state. Returns migrations in any of the specified
	// states.
	States types.List `tfsdk:"-"`
}

func (to *ListMigrationsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListMigrationsRequest) {
	if !from.States.IsNull() && !from.States.IsUnknown() && to.States.IsNull() && len(from.States.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for States, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.States = from.States
	}
}

func (to *ListMigrationsRequest) SyncFieldsDuringRead(ctx context.Context, from ListMigrationsRequest) {
	if !from.States.IsNull() && !from.States.IsUnknown() && to.States.IsNull() && len(from.States.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for States, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.States = from.States
	}
}

func (m ListMigrationsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["asset_type"] = attrs["asset_type"].SetOptional()
	attrs["states"] = attrs["states"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListMigrationsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListMigrationsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"states": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListMigrationsRequest
// only implements ToObjectValue() and Type().
func (m ListMigrationsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"asset_type": m.AssetType,
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"states":     m.States,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListMigrationsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"asset_type": types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"states": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetStates returns the value of the States field in ListMigrationsRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListMigrationsRequest) GetStates(ctx context.Context) ([]types.String, bool) {
	if m.States.IsNull() || m.States.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.States.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStates sets the value of the States field in ListMigrationsRequest.
func (m *ListMigrationsRequest) SetStates(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["states"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.States = types.ListValueMust(t, vs)
}

// Response containing a paginated list of migrations matching the request
// filters.
type ListMigrationsResponse struct {
	// The migrations matching the request filters.
	Migrations types.List `tfsdk:"migrations"`
	// Token for the next page. Omitted if there are no more results.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListMigrationsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListMigrationsResponse) {
	if !from.Migrations.IsNull() && !from.Migrations.IsUnknown() && to.Migrations.IsNull() && len(from.Migrations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Migrations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Migrations = from.Migrations
	}
	if !from.Migrations.IsNull() && !from.Migrations.IsUnknown() {
		if toMigrations, ok := to.GetMigrations(ctx); ok {
			if fromMigrations, ok := from.GetMigrations(ctx); ok {
				// Recursively sync the fields of each Migrations element by position.
				for i := range toMigrations {
					if i < len(fromMigrations) {
						toMigrations[i].SyncFieldsDuringCreateOrUpdate(ctx, fromMigrations[i])
					}
				}
				to.SetMigrations(ctx, toMigrations)
			}
		}
	}
}

func (to *ListMigrationsResponse) SyncFieldsDuringRead(ctx context.Context, from ListMigrationsResponse) {
	if !from.Migrations.IsNull() && !from.Migrations.IsUnknown() && to.Migrations.IsNull() && len(from.Migrations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Migrations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Migrations = from.Migrations
	}
	if !from.Migrations.IsNull() && !from.Migrations.IsUnknown() {
		if toMigrations, ok := to.GetMigrations(ctx); ok {
			if fromMigrations, ok := from.GetMigrations(ctx); ok {
				for i := range toMigrations {
					if i < len(fromMigrations) {
						toMigrations[i].SyncFieldsDuringRead(ctx, fromMigrations[i])
					}
				}
				to.SetMigrations(ctx, toMigrations)
			}
		}
	}
}

func (m ListMigrationsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["migrations"] = attrs["migrations"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListMigrationsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListMigrationsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"migrations": reflect.TypeOf(Migration{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListMigrationsResponse
// only implements ToObjectValue() and Type().
func (m ListMigrationsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"migrations":      m.Migrations,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListMigrationsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"migrations": basetypes.ListType{
				ElemType: Migration{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetMigrations returns the value of the Migrations field in ListMigrationsResponse as
// a slice of Migration values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListMigrationsResponse) GetMigrations(ctx context.Context) ([]Migration, bool) {
	if m.Migrations.IsNull() || m.Migrations.IsUnknown() {
		return nil, false
	}
	var v []Migration
	d := m.Migrations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMigrations sets the value of the Migrations field in ListMigrationsResponse.
func (m *ListMigrationsResponse) SetMigrations(ctx context.Context, v []Migration) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["migrations"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Migrations = types.ListValueMust(t, vs)
}

// A persistent object tracking the full lifecycle of migrating one Classic
// compute asset to Serverless. Each Migration has a well-defined state machine
// and is the unit of orchestration, approval, and rollback.
type Migration struct {
	// Type of Classic compute asset being migrated. Cannot be changed after
	// creation.
	AssetType types.String `tfsdk:"asset_type"`
	// URI of the source asset (e.g., "jobs/12345"). Cannot be changed after
	// creation.
	AssetUri types.String `tfsdk:"asset_uri"`
	// When true, the server skips the translation review (TRANSLATION_READY)
	// instead of waiting for an explicit ApproveMigrationTranslation.
	// MIGRATION_REVIEW is never skipped. Unset/false means the translation
	// review requires explicit user action. Set at creation; cannot be changed
	// afterward. Does not suppress cancellation or terminal failures.
	AutoApprove types.Bool `tfsdk:"auto_approve"`
	// Baking window configuration. Mutable via UpdateMigration before baking
	// starts.
	BakingConfig types.Object `tfsdk:"baking_config"`
	// Summary of baking metrics after the baking window completes.
	BakingReport types.Object `tfsdk:"baking_report"`
	// Timestamp when the migration was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// Resource version for optimistic concurrency control. Clients read this
	// from GetMigration or ListMigrations and pass it back as the request
	// `etag` on update, delete, approve, reject, and cancel methods.
	Etag types.String `tfsdk:"etag"`
	// The user who initiated this migration and their role. Set at creation.
	Initiator types.Object `tfsdk:"initiator"`
	// Server-generated unique identifier for this migration.
	MigrationId types.String `tfsdk:"migration_id"`
	// Agent-produced planning report, presented to the user at the review
	// gates.
	MigrationReport types.Object `tfsdk:"migration_report"`
	// When true, Databricks runs the working copy to validate it after the
	// translation review. Unset/false skips the run, so the migration goes from
	// the translation review straight to execution review. Whether the
	// translation review waits for the user is controlled by `auto_approve`,
	// not this field. Set at creation; cannot be changed afterward.
	RunValidation types.Bool `tfsdk:"run_validation"`
	// Current state of the migration in the state machine. Managed by the
	// server.
	State types.String `tfsdk:"state"`
	// History of state transitions, ordered chronologically. Each entry records
	// the state entered and the timestamp of the transition.
	StateHistory types.List `tfsdk:"state_history"`
	// Agent-produced validation results from sandbox execution.
	ValidationReport types.Object `tfsdk:"validation_report"`
}

func (to *Migration) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Migration) {
	if !from.BakingConfig.IsNull() && !from.BakingConfig.IsUnknown() {
		if toBakingConfig, ok := to.GetBakingConfig(ctx); ok {
			if fromBakingConfig, ok := from.GetBakingConfig(ctx); ok {
				// Recursively sync the fields of BakingConfig
				toBakingConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromBakingConfig)
				to.SetBakingConfig(ctx, toBakingConfig)
			}
		}
	}
	if !from.BakingReport.IsNull() && !from.BakingReport.IsUnknown() {
		if toBakingReport, ok := to.GetBakingReport(ctx); ok {
			if fromBakingReport, ok := from.GetBakingReport(ctx); ok {
				// Recursively sync the fields of BakingReport
				toBakingReport.SyncFieldsDuringCreateOrUpdate(ctx, fromBakingReport)
				to.SetBakingReport(ctx, toBakingReport)
			}
		}
	}
	if !from.Initiator.IsNull() && !from.Initiator.IsUnknown() {
		if toInitiator, ok := to.GetInitiator(ctx); ok {
			if fromInitiator, ok := from.GetInitiator(ctx); ok {
				// Recursively sync the fields of Initiator
				toInitiator.SyncFieldsDuringCreateOrUpdate(ctx, fromInitiator)
				to.SetInitiator(ctx, toInitiator)
			}
		}
	}
	if !from.MigrationReport.IsNull() && !from.MigrationReport.IsUnknown() {
		if toMigrationReport, ok := to.GetMigrationReport(ctx); ok {
			if fromMigrationReport, ok := from.GetMigrationReport(ctx); ok {
				// Recursively sync the fields of MigrationReport
				toMigrationReport.SyncFieldsDuringCreateOrUpdate(ctx, fromMigrationReport)
				to.SetMigrationReport(ctx, toMigrationReport)
			}
		}
	}
	if !from.StateHistory.IsNull() && !from.StateHistory.IsUnknown() && to.StateHistory.IsNull() && len(from.StateHistory.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for StateHistory, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.StateHistory = from.StateHistory
	}
	if !from.StateHistory.IsNull() && !from.StateHistory.IsUnknown() {
		if toStateHistory, ok := to.GetStateHistory(ctx); ok {
			if fromStateHistory, ok := from.GetStateHistory(ctx); ok {
				// Recursively sync the fields of each StateHistory element by position.
				for i := range toStateHistory {
					if i < len(fromStateHistory) {
						toStateHistory[i].SyncFieldsDuringCreateOrUpdate(ctx, fromStateHistory[i])
					}
				}
				to.SetStateHistory(ctx, toStateHistory)
			}
		}
	}
	if !from.ValidationReport.IsNull() && !from.ValidationReport.IsUnknown() {
		if toValidationReport, ok := to.GetValidationReport(ctx); ok {
			if fromValidationReport, ok := from.GetValidationReport(ctx); ok {
				// Recursively sync the fields of ValidationReport
				toValidationReport.SyncFieldsDuringCreateOrUpdate(ctx, fromValidationReport)
				to.SetValidationReport(ctx, toValidationReport)
			}
		}
	}
}

func (to *Migration) SyncFieldsDuringRead(ctx context.Context, from Migration) {
	if !from.BakingConfig.IsNull() && !from.BakingConfig.IsUnknown() {
		if toBakingConfig, ok := to.GetBakingConfig(ctx); ok {
			if fromBakingConfig, ok := from.GetBakingConfig(ctx); ok {
				toBakingConfig.SyncFieldsDuringRead(ctx, fromBakingConfig)
				to.SetBakingConfig(ctx, toBakingConfig)
			}
		}
	}
	if !from.BakingReport.IsNull() && !from.BakingReport.IsUnknown() {
		if toBakingReport, ok := to.GetBakingReport(ctx); ok {
			if fromBakingReport, ok := from.GetBakingReport(ctx); ok {
				toBakingReport.SyncFieldsDuringRead(ctx, fromBakingReport)
				to.SetBakingReport(ctx, toBakingReport)
			}
		}
	}
	if !from.Initiator.IsNull() && !from.Initiator.IsUnknown() {
		if toInitiator, ok := to.GetInitiator(ctx); ok {
			if fromInitiator, ok := from.GetInitiator(ctx); ok {
				toInitiator.SyncFieldsDuringRead(ctx, fromInitiator)
				to.SetInitiator(ctx, toInitiator)
			}
		}
	}
	if !from.MigrationReport.IsNull() && !from.MigrationReport.IsUnknown() {
		if toMigrationReport, ok := to.GetMigrationReport(ctx); ok {
			if fromMigrationReport, ok := from.GetMigrationReport(ctx); ok {
				toMigrationReport.SyncFieldsDuringRead(ctx, fromMigrationReport)
				to.SetMigrationReport(ctx, toMigrationReport)
			}
		}
	}
	if !from.StateHistory.IsNull() && !from.StateHistory.IsUnknown() && to.StateHistory.IsNull() && len(from.StateHistory.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for StateHistory, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.StateHistory = from.StateHistory
	}
	if !from.StateHistory.IsNull() && !from.StateHistory.IsUnknown() {
		if toStateHistory, ok := to.GetStateHistory(ctx); ok {
			if fromStateHistory, ok := from.GetStateHistory(ctx); ok {
				for i := range toStateHistory {
					if i < len(fromStateHistory) {
						toStateHistory[i].SyncFieldsDuringRead(ctx, fromStateHistory[i])
					}
				}
				to.SetStateHistory(ctx, toStateHistory)
			}
		}
	}
	if !from.ValidationReport.IsNull() && !from.ValidationReport.IsUnknown() {
		if toValidationReport, ok := to.GetValidationReport(ctx); ok {
			if fromValidationReport, ok := from.GetValidationReport(ctx); ok {
				toValidationReport.SyncFieldsDuringRead(ctx, fromValidationReport)
				to.SetValidationReport(ctx, toValidationReport)
			}
		}
	}
}

func (m Migration) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["asset_type"] = attrs["asset_type"].SetRequired()
	attrs["asset_type"] = attrs["asset_type"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["asset_uri"] = attrs["asset_uri"].SetRequired()
	attrs["asset_uri"] = attrs["asset_uri"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["auto_approve"] = attrs["auto_approve"].SetOptional()
	attrs["auto_approve"] = attrs["auto_approve"].(tfschema.BoolAttributeBuilder).AddPlanModifier(boolplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["baking_config"] = attrs["baking_config"].SetOptional()
	attrs["baking_report"] = attrs["baking_report"].SetComputed()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["etag"] = attrs["etag"].SetComputed()
	attrs["initiator"] = attrs["initiator"].SetComputed()
	attrs["migration_id"] = attrs["migration_id"].SetOptional()
	attrs["migration_id"] = attrs["migration_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["migration_report"] = attrs["migration_report"].SetComputed()
	attrs["run_validation"] = attrs["run_validation"].SetOptional()
	attrs["run_validation"] = attrs["run_validation"].(tfschema.BoolAttributeBuilder).AddPlanModifier(boolplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["state"] = attrs["state"].SetComputed()
	attrs["state_history"] = attrs["state_history"].SetComputed()
	attrs["validation_report"] = attrs["validation_report"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Migration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Migration) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"baking_config":     reflect.TypeOf(BakingConfig{}),
		"baking_report":     reflect.TypeOf(BakingReport{}),
		"initiator":         reflect.TypeOf(MigrationInitiator{}),
		"migration_report":  reflect.TypeOf(MigrationReport{}),
		"state_history":     reflect.TypeOf(MigrationStateTransition{}),
		"validation_report": reflect.TypeOf(ValidationReport{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Migration
// only implements ToObjectValue() and Type().
func (m Migration) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"asset_type":        m.AssetType,
			"asset_uri":         m.AssetUri,
			"auto_approve":      m.AutoApprove,
			"baking_config":     m.BakingConfig,
			"baking_report":     m.BakingReport,
			"create_time":       m.CreateTime,
			"etag":              m.Etag,
			"initiator":         m.Initiator,
			"migration_id":      m.MigrationId,
			"migration_report":  m.MigrationReport,
			"run_validation":    m.RunValidation,
			"state":             m.State,
			"state_history":     m.StateHistory,
			"validation_report": m.ValidationReport,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Migration) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"asset_type":       types.StringType,
			"asset_uri":        types.StringType,
			"auto_approve":     types.BoolType,
			"baking_config":    BakingConfig{}.Type(ctx),
			"baking_report":    BakingReport{}.Type(ctx),
			"create_time":      timetypes.RFC3339{}.Type(ctx),
			"etag":             types.StringType,
			"initiator":        MigrationInitiator{}.Type(ctx),
			"migration_id":     types.StringType,
			"migration_report": MigrationReport{}.Type(ctx),
			"run_validation":   types.BoolType,
			"state":            types.StringType,
			"state_history": basetypes.ListType{
				ElemType: MigrationStateTransition{}.Type(ctx),
			},
			"validation_report": ValidationReport{}.Type(ctx),
		},
	}
}

// GetBakingConfig returns the value of the BakingConfig field in Migration as
// a BakingConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *Migration) GetBakingConfig(ctx context.Context) (BakingConfig, bool) {
	var e BakingConfig
	if m.BakingConfig.IsNull() || m.BakingConfig.IsUnknown() {
		return e, false
	}
	var v BakingConfig
	d := m.BakingConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBakingConfig sets the value of the BakingConfig field in Migration.
func (m *Migration) SetBakingConfig(ctx context.Context, v BakingConfig) {
	vs := v.ToObjectValue(ctx)
	m.BakingConfig = vs
}

// GetBakingReport returns the value of the BakingReport field in Migration as
// a BakingReport value.
// If the field is unknown or null, the boolean return value is false.
func (m *Migration) GetBakingReport(ctx context.Context) (BakingReport, bool) {
	var e BakingReport
	if m.BakingReport.IsNull() || m.BakingReport.IsUnknown() {
		return e, false
	}
	var v BakingReport
	d := m.BakingReport.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBakingReport sets the value of the BakingReport field in Migration.
func (m *Migration) SetBakingReport(ctx context.Context, v BakingReport) {
	vs := v.ToObjectValue(ctx)
	m.BakingReport = vs
}

// GetInitiator returns the value of the Initiator field in Migration as
// a MigrationInitiator value.
// If the field is unknown or null, the boolean return value is false.
func (m *Migration) GetInitiator(ctx context.Context) (MigrationInitiator, bool) {
	var e MigrationInitiator
	if m.Initiator.IsNull() || m.Initiator.IsUnknown() {
		return e, false
	}
	var v MigrationInitiator
	d := m.Initiator.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInitiator sets the value of the Initiator field in Migration.
func (m *Migration) SetInitiator(ctx context.Context, v MigrationInitiator) {
	vs := v.ToObjectValue(ctx)
	m.Initiator = vs
}

// GetMigrationReport returns the value of the MigrationReport field in Migration as
// a MigrationReport value.
// If the field is unknown or null, the boolean return value is false.
func (m *Migration) GetMigrationReport(ctx context.Context) (MigrationReport, bool) {
	var e MigrationReport
	if m.MigrationReport.IsNull() || m.MigrationReport.IsUnknown() {
		return e, false
	}
	var v MigrationReport
	d := m.MigrationReport.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMigrationReport sets the value of the MigrationReport field in Migration.
func (m *Migration) SetMigrationReport(ctx context.Context, v MigrationReport) {
	vs := v.ToObjectValue(ctx)
	m.MigrationReport = vs
}

// GetStateHistory returns the value of the StateHistory field in Migration as
// a slice of MigrationStateTransition values.
// If the field is unknown or null, the boolean return value is false.
func (m *Migration) GetStateHistory(ctx context.Context) ([]MigrationStateTransition, bool) {
	if m.StateHistory.IsNull() || m.StateHistory.IsUnknown() {
		return nil, false
	}
	var v []MigrationStateTransition
	d := m.StateHistory.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStateHistory sets the value of the StateHistory field in Migration.
func (m *Migration) SetStateHistory(ctx context.Context, v []MigrationStateTransition) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["state_history"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.StateHistory = types.ListValueMust(t, vs)
}

// GetValidationReport returns the value of the ValidationReport field in Migration as
// a ValidationReport value.
// If the field is unknown or null, the boolean return value is false.
func (m *Migration) GetValidationReport(ctx context.Context) (ValidationReport, bool) {
	var e ValidationReport
	if m.ValidationReport.IsNull() || m.ValidationReport.IsUnknown() {
		return e, false
	}
	var v ValidationReport
	d := m.ValidationReport.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetValidationReport sets the value of the ValidationReport field in Migration.
func (m *Migration) SetValidationReport(ctx context.Context, v ValidationReport) {
	vs := v.ToObjectValue(ctx)
	m.ValidationReport = vs
}

// A working copy (job or notebook) the migration agent created for the
// migration.
type MigrationAsset struct {
	// Type of asset this working copy represents.
	AssetType types.String `tfsdk:"asset_type"`
	// Identifier of the asset.
	Id types.String `tfsdk:"id"`
	// Link to the asset, e.g. "jobs/12345".
	Uri types.String `tfsdk:"uri"`
}

func (to *MigrationAsset) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from MigrationAsset) {
}

func (to *MigrationAsset) SyncFieldsDuringRead(ctx context.Context, from MigrationAsset) {
}

func (m MigrationAsset) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["asset_type"] = attrs["asset_type"].SetComputed()
	attrs["id"] = attrs["id"].SetComputed()
	attrs["uri"] = attrs["uri"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MigrationAsset.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m MigrationAsset) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MigrationAsset
// only implements ToObjectValue() and Type().
func (m MigrationAsset) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"asset_type": m.AssetType,
			"id":         m.Id,
			"uri":        m.Uri,
		})
}

// Type implements basetypes.ObjectValuable.
func (m MigrationAsset) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"asset_type": types.StringType,
			"id":         types.StringType,
			"uri":        types.StringType,
		},
	}
}

// The initiator of a migration, including their role and user identity.
type MigrationInitiator struct {
	// The role of the initiator, determining approval authority. Set by the
	// server: the role is resolved from the caller's identity and permissions
	// on the asset, never taken from the request, since a caller must not be
	// able to assert its own approval authority.
	Role types.String `tfsdk:"role"`
	// The user ID of the initiator. Set by the server from the request context.
	UserId types.Int64 `tfsdk:"user_id"`
}

func (to *MigrationInitiator) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from MigrationInitiator) {
}

func (to *MigrationInitiator) SyncFieldsDuringRead(ctx context.Context, from MigrationInitiator) {
}

func (m MigrationInitiator) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["role"] = attrs["role"].SetComputed()
	attrs["user_id"] = attrs["user_id"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MigrationInitiator.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m MigrationInitiator) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MigrationInitiator
// only implements ToObjectValue() and Type().
func (m MigrationInitiator) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"role":    m.Role,
			"user_id": m.UserId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m MigrationInitiator) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"role":    types.StringType,
			"user_id": types.Int64Type,
		},
	}
}

// Agent-produced planning report, presented to the user at the review gates.
type MigrationReport struct {
	// The primary asset the agent created — the entry point / validation
	// target.
	PrimaryAsset types.Object `tfsdk:"primary_asset"`
	// Additional assets the agent created (e.g. notebooks or jobs the primary
	// asset invokes).
	SecondaryAssets types.List `tfsdk:"secondary_assets"`
}

func (to *MigrationReport) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from MigrationReport) {
	if !from.PrimaryAsset.IsNull() && !from.PrimaryAsset.IsUnknown() {
		if toPrimaryAsset, ok := to.GetPrimaryAsset(ctx); ok {
			if fromPrimaryAsset, ok := from.GetPrimaryAsset(ctx); ok {
				// Recursively sync the fields of PrimaryAsset
				toPrimaryAsset.SyncFieldsDuringCreateOrUpdate(ctx, fromPrimaryAsset)
				to.SetPrimaryAsset(ctx, toPrimaryAsset)
			}
		}
	}
	if !from.SecondaryAssets.IsNull() && !from.SecondaryAssets.IsUnknown() && to.SecondaryAssets.IsNull() && len(from.SecondaryAssets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SecondaryAssets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SecondaryAssets = from.SecondaryAssets
	}
	if !from.SecondaryAssets.IsNull() && !from.SecondaryAssets.IsUnknown() {
		if toSecondaryAssets, ok := to.GetSecondaryAssets(ctx); ok {
			if fromSecondaryAssets, ok := from.GetSecondaryAssets(ctx); ok {
				// Recursively sync the fields of each SecondaryAssets element by position.
				for i := range toSecondaryAssets {
					if i < len(fromSecondaryAssets) {
						toSecondaryAssets[i].SyncFieldsDuringCreateOrUpdate(ctx, fromSecondaryAssets[i])
					}
				}
				to.SetSecondaryAssets(ctx, toSecondaryAssets)
			}
		}
	}
}

func (to *MigrationReport) SyncFieldsDuringRead(ctx context.Context, from MigrationReport) {
	if !from.PrimaryAsset.IsNull() && !from.PrimaryAsset.IsUnknown() {
		if toPrimaryAsset, ok := to.GetPrimaryAsset(ctx); ok {
			if fromPrimaryAsset, ok := from.GetPrimaryAsset(ctx); ok {
				toPrimaryAsset.SyncFieldsDuringRead(ctx, fromPrimaryAsset)
				to.SetPrimaryAsset(ctx, toPrimaryAsset)
			}
		}
	}
	if !from.SecondaryAssets.IsNull() && !from.SecondaryAssets.IsUnknown() && to.SecondaryAssets.IsNull() && len(from.SecondaryAssets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SecondaryAssets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SecondaryAssets = from.SecondaryAssets
	}
	if !from.SecondaryAssets.IsNull() && !from.SecondaryAssets.IsUnknown() {
		if toSecondaryAssets, ok := to.GetSecondaryAssets(ctx); ok {
			if fromSecondaryAssets, ok := from.GetSecondaryAssets(ctx); ok {
				for i := range toSecondaryAssets {
					if i < len(fromSecondaryAssets) {
						toSecondaryAssets[i].SyncFieldsDuringRead(ctx, fromSecondaryAssets[i])
					}
				}
				to.SetSecondaryAssets(ctx, toSecondaryAssets)
			}
		}
	}
}

func (m MigrationReport) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["primary_asset"] = attrs["primary_asset"].SetComputed()
	attrs["secondary_assets"] = attrs["secondary_assets"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MigrationReport.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m MigrationReport) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"primary_asset":    reflect.TypeOf(MigrationAsset{}),
		"secondary_assets": reflect.TypeOf(MigrationAsset{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MigrationReport
// only implements ToObjectValue() and Type().
func (m MigrationReport) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"primary_asset":    m.PrimaryAsset,
			"secondary_assets": m.SecondaryAssets,
		})
}

// Type implements basetypes.ObjectValuable.
func (m MigrationReport) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"primary_asset": MigrationAsset{}.Type(ctx),
			"secondary_assets": basetypes.ListType{
				ElemType: MigrationAsset{}.Type(ctx),
			},
		},
	}
}

// GetPrimaryAsset returns the value of the PrimaryAsset field in MigrationReport as
// a MigrationAsset value.
// If the field is unknown or null, the boolean return value is false.
func (m *MigrationReport) GetPrimaryAsset(ctx context.Context) (MigrationAsset, bool) {
	var e MigrationAsset
	if m.PrimaryAsset.IsNull() || m.PrimaryAsset.IsUnknown() {
		return e, false
	}
	var v MigrationAsset
	d := m.PrimaryAsset.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPrimaryAsset sets the value of the PrimaryAsset field in MigrationReport.
func (m *MigrationReport) SetPrimaryAsset(ctx context.Context, v MigrationAsset) {
	vs := v.ToObjectValue(ctx)
	m.PrimaryAsset = vs
}

// GetSecondaryAssets returns the value of the SecondaryAssets field in MigrationReport as
// a slice of MigrationAsset values.
// If the field is unknown or null, the boolean return value is false.
func (m *MigrationReport) GetSecondaryAssets(ctx context.Context) ([]MigrationAsset, bool) {
	if m.SecondaryAssets.IsNull() || m.SecondaryAssets.IsUnknown() {
		return nil, false
	}
	var v []MigrationAsset
	d := m.SecondaryAssets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSecondaryAssets sets the value of the SecondaryAssets field in MigrationReport.
func (m *MigrationReport) SetSecondaryAssets(ctx context.Context, v []MigrationAsset) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["secondary_assets"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SecondaryAssets = types.ListValueMust(t, vs)
}

// A record of a state transition in the migration lifecycle.
type MigrationStateTransition struct {
	// Free-text the service recorded for this transition — e.g. an
	// internally-determined cancellation or failure reason. Empty for
	// user-driven transitions.
	InternalComment types.String `tfsdk:"internal_comment"`
	// The state that was entered.
	State types.String `tfsdk:"state"`
	// Timestamp when this state was entered.
	TransitionTime timetypes.RFC3339 `tfsdk:"transition_time"`
	// Free-text the initiating user supplied at a user-driven transition —
	// reviewer feedback on an approval, or the reason for a rejection or
	// cancellation. Empty for transitions the service drives internally.
	UserComment types.String `tfsdk:"user_comment"`
}

func (to *MigrationStateTransition) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from MigrationStateTransition) {
}

func (to *MigrationStateTransition) SyncFieldsDuringRead(ctx context.Context, from MigrationStateTransition) {
}

func (m MigrationStateTransition) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["internal_comment"] = attrs["internal_comment"].SetComputed()
	attrs["state"] = attrs["state"].SetComputed()
	attrs["transition_time"] = attrs["transition_time"].SetComputed()
	attrs["user_comment"] = attrs["user_comment"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MigrationStateTransition.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m MigrationStateTransition) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MigrationStateTransition
// only implements ToObjectValue() and Type().
func (m MigrationStateTransition) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"internal_comment": m.InternalComment,
			"state":            m.State,
			"transition_time":  m.TransitionTime,
			"user_comment":     m.UserComment,
		})
}

// Type implements basetypes.ObjectValuable.
func (m MigrationStateTransition) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internal_comment": types.StringType,
			"state":            types.StringType,
			"transition_time":  timetypes.RFC3339{}.Type(ctx),
			"user_comment":     types.StringType,
		},
	}
}

// A request to reject migration execution at the MIGRATION_REVIEW gate, looping
// back to REVISION_PENDING.
type RejectMigrationExecutionRequest struct {
	// Optional resource version for optimistic concurrency control. If
	// supplied, it must match the current server value.
	Etag types.String `tfsdk:"etag"`
	// The ID of the migration to reject.
	MigrationId types.String `tfsdk:"-"`
	// Required free-text reason for the rejection, fed back into re-planning so
	// the next plan can address the user's concerns.
	RejectionReason types.String `tfsdk:"rejection_reason"`
}

func (to *RejectMigrationExecutionRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RejectMigrationExecutionRequest) {
}

func (to *RejectMigrationExecutionRequest) SyncFieldsDuringRead(ctx context.Context, from RejectMigrationExecutionRequest) {
}

func (m RejectMigrationExecutionRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()
	attrs["rejection_reason"] = attrs["rejection_reason"].SetRequired()
	attrs["migration_id"] = attrs["migration_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RejectMigrationExecutionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RejectMigrationExecutionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RejectMigrationExecutionRequest
// only implements ToObjectValue() and Type().
func (m RejectMigrationExecutionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag":             m.Etag,
			"migration_id":     m.MigrationId,
			"rejection_reason": m.RejectionReason,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RejectMigrationExecutionRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag":             types.StringType,
			"migration_id":     types.StringType,
			"rejection_reason": types.StringType,
		},
	}
}

// A request to reject the working copy at the TRANSLATION_READY gate, looping
// back to REVISION_PENDING.
type RejectMigrationTranslationRequest struct {
	// Optional resource version for optimistic concurrency control. If
	// supplied, it must match the current server value.
	Etag types.String `tfsdk:"etag"`
	// The ID of the migration to reject.
	MigrationId types.String `tfsdk:"-"`
	// Required free-text reason for the rejection, fed back into re-planning so
	// the next plan can address the user's concerns.
	RejectionReason types.String `tfsdk:"rejection_reason"`
}

func (to *RejectMigrationTranslationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RejectMigrationTranslationRequest) {
}

func (to *RejectMigrationTranslationRequest) SyncFieldsDuringRead(ctx context.Context, from RejectMigrationTranslationRequest) {
}

func (m RejectMigrationTranslationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()
	attrs["rejection_reason"] = attrs["rejection_reason"].SetRequired()
	attrs["migration_id"] = attrs["migration_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RejectMigrationTranslationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RejectMigrationTranslationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RejectMigrationTranslationRequest
// only implements ToObjectValue() and Type().
func (m RejectMigrationTranslationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag":             m.Etag,
			"migration_id":     m.MigrationId,
			"rejection_reason": m.RejectionReason,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RejectMigrationTranslationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag":             types.StringType,
			"migration_id":     types.StringType,
			"rejection_reason": types.StringType,
		},
	}
}

// Validation results from running the translated workload in the sandbox.
type ValidationReport struct {
	// Duration of the sandbox validation run in milliseconds. Pair with
	// start_time to derive the end time.
	DurationMs types.Int64 `tfsdk:"duration_ms"`
	// Databricks-internal recommendation on whether to proceed with the
	// migration, based on the validation run. May differ from run_status: a run
	// can succeed while the recommendation is negative, or vice versa.
	Passed types.Bool `tfsdk:"passed"`
	// Status of the validation run itself (the objective run outcome), as
	// reported by the run system, e.g. "SUCCESS" / "FAILED" / "TIMED_OUT".
	RunStatus types.String `tfsdk:"run_status"`
	// Link to the validation run (job or notebook).
	RunUrl types.String `tfsdk:"run_url"`
	// Timestamp when validation started in the sandbox.
	StartTime timetypes.RFC3339 `tfsdk:"start_time"`
}

func (to *ValidationReport) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ValidationReport) {
}

func (to *ValidationReport) SyncFieldsDuringRead(ctx context.Context, from ValidationReport) {
}

func (m ValidationReport) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["duration_ms"] = attrs["duration_ms"].SetComputed()
	attrs["passed"] = attrs["passed"].SetComputed()
	attrs["run_status"] = attrs["run_status"].SetComputed()
	attrs["run_url"] = attrs["run_url"].SetComputed()
	attrs["start_time"] = attrs["start_time"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ValidationReport.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ValidationReport) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ValidationReport
// only implements ToObjectValue() and Type().
func (m ValidationReport) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"duration_ms": m.DurationMs,
			"passed":      m.Passed,
			"run_status":  m.RunStatus,
			"run_url":     m.RunUrl,
			"start_time":  m.StartTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ValidationReport) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"duration_ms": types.Int64Type,
			"passed":      types.BoolType,
			"run_status":  types.StringType,
			"run_url":     types.StringType,
			"start_time":  timetypes.RFC3339{}.Type(ctx),
		},
	}
}
