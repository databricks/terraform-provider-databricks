// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package lakewatch_datasource

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/service/lakewatch"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/declarative"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/lakewatch_tf"
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

const resourceName = "lakewatch_datasource"

var _ resource.ResourceWithConfigure = &DatasourceResource{}
var _ resource.ResourceWithModifyPlan = &DatasourceResource{}

func ResourceDatasource() resource.Resource {
	return &DatasourceResource{}
}

type DatasourceResource struct {
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

// Datasource extends the main model with additional fields.
type Datasource struct {
	// Free-form key-value annotations on the datasource, managed by the client.
	Annotations types.Map `tfsdk:"annotations"`

	Autoloader types.Object `tfsdk:"autoloader"`
	// Name of the Unity Catalog catalog that contains the datasource. Required
	// on create and cannot be changed afterward. Reads return the value as
	// supplied; the catalog's current name is reported in
	// `effective_catalog_name`.
	CatalogName types.String `tfsdk:"catalog_name"`
	// Compute mode for the datasource's ingestion job. If unset, the
	// `default_compute_mode` of the catalog's datasource configuration applies,
	// or `DS_COMPUTE_PERFORMANCE_OPTIMIZED` if the catalog has none.
	ComputeMode types.String `tfsdk:"compute_mode"`
	// ID of the workspace whose compute runs this datasource's ingestion job
	// and owns its associated resources, such as the storage the ingested data
	// lands in. Server-assigned.
	ComputeWorkspaceId types.String `tfsdk:"compute_workspace_id"`

	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`

	CreatedBy types.Object `tfsdk:"created_by"`
	// Unity Catalog securable ID (a UUID) of the datasource. It never changes,
	// including when the datasource is renamed in Unity Catalog, and is the
	// final segment of `name`.
	DatasourceId types.String `tfsdk:"datasource_id"`
	// Name of the datasource in Unity Catalog, unique within its schema.
	// Required on create and cannot be changed afterward. Reads return the
	// value as supplied; the datasource's current Unity Catalog name, which
	// reflects renames, is reported in `effective_datasource_name`.
	DatasourceName types.String `tfsdk:"datasource_name"`
	// Free-form description of the datasource.
	Description types.String `tfsdk:"description"`
	// Current name of the Unity Catalog catalog that contains the datasource,
	// which reflects catalog renames without a change to `catalog_name`.
	// Populated on get, create, and update responses; list responses omit it.
	EffectiveCatalogName types.String `tfsdk:"effective_catalog_name"`
	// Current UC leaf name of the datasource (server-owned). Tracks UC renames
	// so callers can observe the live name without having to update the
	// client-owned `datasource_name`; unset until resolved.
	EffectiveDatasourceName types.String `tfsdk:"effective_datasource_name"`
	// The identity the datasource runs as: the resolved `run_as_principal`, or
	// the requester's identity if `run_as_principal` was omitted on create or
	// cleared on update. Given as a user email or service principal application
	// ID, or as the numeric principal ID if the identity can no longer be
	// resolved. Populated on get, create, and update responses; list responses
	// omit it.
	EffectiveRunAsPrincipal types.String `tfsdk:"effective_run_as_principal"`
	// Current name of the Unity Catalog schema that contains the datasource,
	// which reflects schema renames without a change to `schema_name`.
	// Populated on get, create, and update responses; list responses omit it.
	EffectiveSchemaName types.String `tfsdk:"effective_schema_name"`
	// Opaque entity tag for optimistic concurrency control. Returned by get,
	// create, and update responses and by list responses in the
	// `DATASOURCE_VIEW_FULL` view. To make an update conditional on the
	// datasource not having changed since it was read, pass this value as the
	// update request's `etag`. Changes to `status` do not change the etag.
	Etag types.String `tfsdk:"etag"`

	ExistingTable types.Object `tfsdk:"existing_table"`
	// Resource name of the datasource, in the format
	// `datasources/{datasource_id}`. Unlike the datasource's Unity Catalog
	// location (`catalog_name`, `schema_name`, and `datasource_name`), the name
	// does not change when the datasource or its catalog or schema is renamed
	// in Unity Catalog.
	Name types.String `tfsdk:"name"`
	// Unity Catalog owner of the datasource, assigned by the server. Create
	// responses omit it.
	Owner types.Object `tfsdk:"owner"`
	// Parsers that schematize the ingested data. Each record is matched against
	// the parsers in order and written to the parsed table of the first parser
	// that it matches. Existing-table and Unity Catalog connection datasources
	// require at least one parser; an Auto Loader datasource with no parsers
	// writes only its raw table.
	Parsers types.List `tfsdk:"parsers"`
	// The user or service principal that the datasource runs as, given as a
	// user email or service principal application ID. When it is supplied on
	// create or update, the requester must be authorized to run the datasource
	// as that principal, or the request is rejected. When it is omitted on
	// create, the datasource runs as the requester. On update, a mask that
	// includes `run_as_principal` without a value clears it, so the datasource
	// runs as the requester; a mask that omits it keeps the current identity.
	// Reads return the value as supplied; the identity the datasource runs as
	// is reported in `effective_run_as_principal`.
	RunAsPrincipal types.String `tfsdk:"run_as_principal"`
	// Schedule on which the datasource's ingestion job runs. A datasource
	// cannot be paused, so a request that sets `schedule.enabled` to `false` is
	// rejected. Continuous schedules are not supported for Unity Catalog
	// connection datasources.
	Schedule types.Object `tfsdk:"schedule"`
	// Name of the Unity Catalog schema that contains the datasource. Required
	// on create and cannot be changed afterward. Reads return the value as
	// supplied; the schema's current name is reported in
	// `effective_schema_name`.
	SchemaName types.String `tfsdk:"schema_name"`
	// Runtime status of the jobs and pipelines that run this datasource.
	Status types.Object `tfsdk:"status"`

	UnityCatalogConnection types.Object `tfsdk:"unity_catalog_connection"`

	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`

	UpdatedBy types.Object `tfsdk:"updated_by"`
	// Version number of the datasource, incremented on each update.
	// Informational only; use `etag` for optimistic concurrency control.
	Version types.Int64 `tfsdk:"version"`
	// ID of the SQL warehouse that runs the checks Lakewatch performs to
	// validate a create or update request, such as source and destination
	// access checks. Required and non-blank on every create and update request;
	// there is no default warehouse. The requester must be able to use the
	// warehouse. Used only for those checks: it is not stored, is not returned
	// on reads, and does not affect the compute that runs the ingestion job.
	WarehouseId    types.String `tfsdk:"warehouse_id"`
	ProviderConfig types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// Datasource struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m Datasource) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"annotations":              reflect.TypeOf(types.String{}),
		"autoloader":               reflect.TypeOf(lakewatch_tf.AutoLoader{}),
		"created_by":               reflect.TypeOf(lakewatch_tf.Principal{}),
		"existing_table":           reflect.TypeOf(lakewatch_tf.ExistingTable{}),
		"owner":                    reflect.TypeOf(lakewatch_tf.Principal{}),
		"parsers":                  reflect.TypeOf(lakewatch_tf.ParserRef{}),
		"schedule":                 reflect.TypeOf(lakewatch_tf.Schedule{}),
		"status":                   reflect.TypeOf(lakewatch_tf.DatasourceStatus{}),
		"unity_catalog_connection": reflect.TypeOf(lakewatch_tf.UnityCatalogConnection{}),
		"updated_by":               reflect.TypeOf(lakewatch_tf.Principal{}),
		"provider_config":          reflect.TypeOf(ProviderConfig{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Datasource
// only implements ToObjectValue() and Type().
func (m Datasource) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"annotations": m.Annotations,
			"autoloader":                 m.Autoloader,
			"catalog_name":               m.CatalogName,
			"compute_mode":               m.ComputeMode,
			"compute_workspace_id":       m.ComputeWorkspaceId,
			"create_time":                m.CreateTime,
			"created_by":                 m.CreatedBy,
			"datasource_id":              m.DatasourceId,
			"datasource_name":            m.DatasourceName,
			"description":                m.Description,
			"effective_catalog_name":     m.EffectiveCatalogName,
			"effective_datasource_name":  m.EffectiveDatasourceName,
			"effective_run_as_principal": m.EffectiveRunAsPrincipal,
			"effective_schema_name":      m.EffectiveSchemaName,
			"etag":                       m.Etag,
			"existing_table":             m.ExistingTable,
			"name":                       m.Name,
			"owner":                      m.Owner,
			"parsers":                    m.Parsers,
			"run_as_principal":           m.RunAsPrincipal,
			"schedule":                   m.Schedule,
			"schema_name":                m.SchemaName,
			"status":                     m.Status,
			"unity_catalog_connection":   m.UnityCatalogConnection,
			"update_time":                m.UpdateTime,
			"updated_by":                 m.UpdatedBy,
			"version":                    m.Version,
			"warehouse_id":               m.WarehouseId,

			"provider_config": m.ProviderConfig,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m Datasource) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"annotations": basetypes.MapType{
			ElemType: types.StringType,
		},
			"autoloader":                 lakewatch_tf.AutoLoader{}.Type(ctx),
			"catalog_name":               types.StringType,
			"compute_mode":               types.StringType,
			"compute_workspace_id":       types.StringType,
			"create_time":                timetypes.RFC3339{}.Type(ctx),
			"created_by":                 lakewatch_tf.Principal{}.Type(ctx),
			"datasource_id":              types.StringType,
			"datasource_name":            types.StringType,
			"description":                types.StringType,
			"effective_catalog_name":     types.StringType,
			"effective_datasource_name":  types.StringType,
			"effective_run_as_principal": types.StringType,
			"effective_schema_name":      types.StringType,
			"etag":                       types.StringType,
			"existing_table":             lakewatch_tf.ExistingTable{}.Type(ctx),
			"name":                       types.StringType,
			"owner":                      lakewatch_tf.Principal{}.Type(ctx),
			"parsers": basetypes.ListType{
				ElemType: lakewatch_tf.ParserRef{}.Type(ctx),
			},
			"run_as_principal":         types.StringType,
			"schedule":                 lakewatch_tf.Schedule{}.Type(ctx),
			"schema_name":              types.StringType,
			"status":                   lakewatch_tf.DatasourceStatus{}.Type(ctx),
			"unity_catalog_connection": lakewatch_tf.UnityCatalogConnection{}.Type(ctx),
			"update_time":              timetypes.RFC3339{}.Type(ctx),
			"updated_by":               lakewatch_tf.Principal{}.Type(ctx),
			"version":                  types.Int64Type,
			"warehouse_id":             types.StringType,

			"provider_config": ProviderConfig{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *Datasource) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Datasource) {
	if !from.Autoloader.IsNull() && !from.Autoloader.IsUnknown() {
		if toAutoloader, ok := to.GetAutoloader(ctx); ok {
			if fromAutoloader, ok := from.GetAutoloader(ctx); ok {
				// Recursively sync the fields of Autoloader
				toAutoloader.SyncFieldsDuringCreateOrUpdate(ctx, fromAutoloader)
				to.SetAutoloader(ctx, toAutoloader)
			}
		}
	}
	if !from.CreatedBy.IsNull() && !from.CreatedBy.IsUnknown() {
		if toCreatedBy, ok := to.GetCreatedBy(ctx); ok {
			if fromCreatedBy, ok := from.GetCreatedBy(ctx); ok {
				// Recursively sync the fields of CreatedBy
				toCreatedBy.SyncFieldsDuringCreateOrUpdate(ctx, fromCreatedBy)
				to.SetCreatedBy(ctx, toCreatedBy)
			}
		}
	}
	if !from.ExistingTable.IsNull() && !from.ExistingTable.IsUnknown() {
		if toExistingTable, ok := to.GetExistingTable(ctx); ok {
			if fromExistingTable, ok := from.GetExistingTable(ctx); ok {
				// Recursively sync the fields of ExistingTable
				toExistingTable.SyncFieldsDuringCreateOrUpdate(ctx, fromExistingTable)
				to.SetExistingTable(ctx, toExistingTable)
			}
		}
	}
	if !from.Owner.IsNull() && !from.Owner.IsUnknown() {
		if toOwner, ok := to.GetOwner(ctx); ok {
			if fromOwner, ok := from.GetOwner(ctx); ok {
				// Recursively sync the fields of Owner
				toOwner.SyncFieldsDuringCreateOrUpdate(ctx, fromOwner)
				to.SetOwner(ctx, toOwner)
			}
		}
	}
	if !from.Parsers.IsNull() && !from.Parsers.IsUnknown() && to.Parsers.IsNull() && len(from.Parsers.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Parsers, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Parsers = from.Parsers
	}
	if !from.Parsers.IsNull() && !from.Parsers.IsUnknown() {
		if toParsers, ok := to.GetParsers(ctx); ok {
			if fromParsers, ok := from.GetParsers(ctx); ok {
				// Recursively sync the fields of each Parsers element by position.
				for i := range toParsers {
					if i < len(fromParsers) {
						toParsers[i].SyncFieldsDuringCreateOrUpdate(ctx, fromParsers[i])
					}
				}
				to.SetParsers(ctx, toParsers)
			}
		}
	}
	if !from.Schedule.IsNull() && !from.Schedule.IsUnknown() {
		if toSchedule, ok := to.GetSchedule(ctx); ok {
			if fromSchedule, ok := from.GetSchedule(ctx); ok {
				// Recursively sync the fields of Schedule
				toSchedule.SyncFieldsDuringCreateOrUpdate(ctx, fromSchedule)
				to.SetSchedule(ctx, toSchedule)
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
	if !from.UnityCatalogConnection.IsNull() && !from.UnityCatalogConnection.IsUnknown() {
		if toUnityCatalogConnection, ok := to.GetUnityCatalogConnection(ctx); ok {
			if fromUnityCatalogConnection, ok := from.GetUnityCatalogConnection(ctx); ok {
				// Recursively sync the fields of UnityCatalogConnection
				toUnityCatalogConnection.SyncFieldsDuringCreateOrUpdate(ctx, fromUnityCatalogConnection)
				to.SetUnityCatalogConnection(ctx, toUnityCatalogConnection)
			}
		}
	}
	if !from.UpdatedBy.IsNull() && !from.UpdatedBy.IsUnknown() {
		if toUpdatedBy, ok := to.GetUpdatedBy(ctx); ok {
			if fromUpdatedBy, ok := from.GetUpdatedBy(ctx); ok {
				// Recursively sync the fields of UpdatedBy
				toUpdatedBy.SyncFieldsDuringCreateOrUpdate(ctx, fromUpdatedBy)
				to.SetUpdatedBy(ctx, toUpdatedBy)
			}
		}
	}
	if !from.WarehouseId.IsUnknown() && !from.WarehouseId.IsNull() {
		// WarehouseId is an input only field and not returned by the service, so we keep the value from the prior state.
		to.WarehouseId = from.WarehouseId
	}
	to.ProviderConfig = from.ProviderConfig

}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *Datasource) SyncFieldsDuringRead(ctx context.Context, from Datasource) {
	if !from.Autoloader.IsNull() && !from.Autoloader.IsUnknown() {
		if toAutoloader, ok := to.GetAutoloader(ctx); ok {
			if fromAutoloader, ok := from.GetAutoloader(ctx); ok {
				toAutoloader.SyncFieldsDuringRead(ctx, fromAutoloader)
				to.SetAutoloader(ctx, toAutoloader)
			}
		}
	}
	if !from.CreatedBy.IsNull() && !from.CreatedBy.IsUnknown() {
		if toCreatedBy, ok := to.GetCreatedBy(ctx); ok {
			if fromCreatedBy, ok := from.GetCreatedBy(ctx); ok {
				toCreatedBy.SyncFieldsDuringRead(ctx, fromCreatedBy)
				to.SetCreatedBy(ctx, toCreatedBy)
			}
		}
	}
	if !from.ExistingTable.IsNull() && !from.ExistingTable.IsUnknown() {
		if toExistingTable, ok := to.GetExistingTable(ctx); ok {
			if fromExistingTable, ok := from.GetExistingTable(ctx); ok {
				toExistingTable.SyncFieldsDuringRead(ctx, fromExistingTable)
				to.SetExistingTable(ctx, toExistingTable)
			}
		}
	}
	if !from.Owner.IsNull() && !from.Owner.IsUnknown() {
		if toOwner, ok := to.GetOwner(ctx); ok {
			if fromOwner, ok := from.GetOwner(ctx); ok {
				toOwner.SyncFieldsDuringRead(ctx, fromOwner)
				to.SetOwner(ctx, toOwner)
			}
		}
	}
	if !from.Parsers.IsNull() && !from.Parsers.IsUnknown() && to.Parsers.IsNull() && len(from.Parsers.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Parsers, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Parsers = from.Parsers
	}
	if !from.Parsers.IsNull() && !from.Parsers.IsUnknown() {
		if toParsers, ok := to.GetParsers(ctx); ok {
			if fromParsers, ok := from.GetParsers(ctx); ok {
				for i := range toParsers {
					if i < len(fromParsers) {
						toParsers[i].SyncFieldsDuringRead(ctx, fromParsers[i])
					}
				}
				to.SetParsers(ctx, toParsers)
			}
		}
	}
	if !from.Schedule.IsNull() && !from.Schedule.IsUnknown() {
		if toSchedule, ok := to.GetSchedule(ctx); ok {
			if fromSchedule, ok := from.GetSchedule(ctx); ok {
				toSchedule.SyncFieldsDuringRead(ctx, fromSchedule)
				to.SetSchedule(ctx, toSchedule)
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
	if !from.UnityCatalogConnection.IsNull() && !from.UnityCatalogConnection.IsUnknown() {
		if toUnityCatalogConnection, ok := to.GetUnityCatalogConnection(ctx); ok {
			if fromUnityCatalogConnection, ok := from.GetUnityCatalogConnection(ctx); ok {
				toUnityCatalogConnection.SyncFieldsDuringRead(ctx, fromUnityCatalogConnection)
				to.SetUnityCatalogConnection(ctx, toUnityCatalogConnection)
			}
		}
	}
	if !from.UpdatedBy.IsNull() && !from.UpdatedBy.IsUnknown() {
		if toUpdatedBy, ok := to.GetUpdatedBy(ctx); ok {
			if fromUpdatedBy, ok := from.GetUpdatedBy(ctx); ok {
				toUpdatedBy.SyncFieldsDuringRead(ctx, fromUpdatedBy)
				to.SetUpdatedBy(ctx, toUpdatedBy)
			}
		}
	}
	if !from.WarehouseId.IsUnknown() && !from.WarehouseId.IsNull() {
		// WarehouseId is an input only field and not returned by the service, so we keep the value from the prior state.
		to.WarehouseId = from.WarehouseId
	}
	to.ProviderConfig = from.ProviderConfig

}

func (m Datasource) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["annotations"] = attrs["annotations"].SetOptional()
	attrs["autoloader"] = attrs["autoloader"].SetOptional()
	attrs["catalog_name"] = attrs["catalog_name"].SetRequired()
	attrs["catalog_name"] = attrs["catalog_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["compute_mode"] = attrs["compute_mode"].SetOptional()
	attrs["compute_workspace_id"] = attrs["compute_workspace_id"].SetComputed()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["created_by"] = attrs["created_by"].SetComputed()
	attrs["datasource_id"] = attrs["datasource_id"].SetComputed()
	attrs["datasource_name"] = attrs["datasource_name"].SetRequired()
	attrs["datasource_name"] = attrs["datasource_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["description"] = attrs["description"].SetOptional()
	attrs["effective_catalog_name"] = attrs["effective_catalog_name"].SetComputed()
	attrs["effective_datasource_name"] = attrs["effective_datasource_name"].SetComputed()
	attrs["effective_run_as_principal"] = attrs["effective_run_as_principal"].SetComputed()
	attrs["effective_schema_name"] = attrs["effective_schema_name"].SetComputed()
	attrs["etag"] = attrs["etag"].SetComputed()
	attrs["existing_table"] = attrs["existing_table"].SetOptional()
	attrs["name"] = attrs["name"].SetComputed()
	attrs["owner"] = attrs["owner"].SetComputed()
	attrs["parsers"] = attrs["parsers"].SetOptional()
	attrs["run_as_principal"] = attrs["run_as_principal"].SetOptional()
	attrs["schedule"] = attrs["schedule"].SetRequired()
	attrs["schema_name"] = attrs["schema_name"].SetRequired()
	attrs["schema_name"] = attrs["schema_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["status"] = attrs["status"].SetComputed()
	attrs["unity_catalog_connection"] = attrs["unity_catalog_connection"].SetOptional()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["updated_by"] = attrs["updated_by"].SetComputed()
	attrs["version"] = attrs["version"].SetComputed()
	attrs["warehouse_id"] = attrs["warehouse_id"].SetRequired()

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	attrs["provider_config"] = attrs["provider_config"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(tfschema.ProviderConfigPlanModifier{})

	return attrs
}

// GetAnnotations returns the value of the Annotations field in Datasource as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *Datasource) GetAnnotations(ctx context.Context) (map[string]types.String, bool) {
	if m.Annotations.IsNull() || m.Annotations.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.Annotations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAnnotations sets the value of the Annotations field in Datasource.
func (m *Datasource) SetAnnotations(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["annotations"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Annotations = types.MapValueMust(t, vs)
}

// GetAutoloader returns the value of the Autoloader field in Datasource as
// a lakewatch_tf.AutoLoader value.
// If the field is unknown or null, the boolean return value is false.
func (m *Datasource) GetAutoloader(ctx context.Context) (lakewatch_tf.AutoLoader, bool) {
	var e lakewatch_tf.AutoLoader
	if m.Autoloader.IsNull() || m.Autoloader.IsUnknown() {
		return e, false
	}
	var v lakewatch_tf.AutoLoader
	d := m.Autoloader.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAutoloader sets the value of the Autoloader field in Datasource.
func (m *Datasource) SetAutoloader(ctx context.Context, v lakewatch_tf.AutoLoader) {
	vs := v.ToObjectValue(ctx)
	m.Autoloader = vs
}

// GetCreatedBy returns the value of the CreatedBy field in Datasource as
// a lakewatch_tf.Principal value.
// If the field is unknown or null, the boolean return value is false.
func (m *Datasource) GetCreatedBy(ctx context.Context) (lakewatch_tf.Principal, bool) {
	var e lakewatch_tf.Principal
	if m.CreatedBy.IsNull() || m.CreatedBy.IsUnknown() {
		return e, false
	}
	var v lakewatch_tf.Principal
	d := m.CreatedBy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCreatedBy sets the value of the CreatedBy field in Datasource.
func (m *Datasource) SetCreatedBy(ctx context.Context, v lakewatch_tf.Principal) {
	vs := v.ToObjectValue(ctx)
	m.CreatedBy = vs
}

// GetExistingTable returns the value of the ExistingTable field in Datasource as
// a lakewatch_tf.ExistingTable value.
// If the field is unknown or null, the boolean return value is false.
func (m *Datasource) GetExistingTable(ctx context.Context) (lakewatch_tf.ExistingTable, bool) {
	var e lakewatch_tf.ExistingTable
	if m.ExistingTable.IsNull() || m.ExistingTable.IsUnknown() {
		return e, false
	}
	var v lakewatch_tf.ExistingTable
	d := m.ExistingTable.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExistingTable sets the value of the ExistingTable field in Datasource.
func (m *Datasource) SetExistingTable(ctx context.Context, v lakewatch_tf.ExistingTable) {
	vs := v.ToObjectValue(ctx)
	m.ExistingTable = vs
}

// GetOwner returns the value of the Owner field in Datasource as
// a lakewatch_tf.Principal value.
// If the field is unknown or null, the boolean return value is false.
func (m *Datasource) GetOwner(ctx context.Context) (lakewatch_tf.Principal, bool) {
	var e lakewatch_tf.Principal
	if m.Owner.IsNull() || m.Owner.IsUnknown() {
		return e, false
	}
	var v lakewatch_tf.Principal
	d := m.Owner.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOwner sets the value of the Owner field in Datasource.
func (m *Datasource) SetOwner(ctx context.Context, v lakewatch_tf.Principal) {
	vs := v.ToObjectValue(ctx)
	m.Owner = vs
}

// GetParsers returns the value of the Parsers field in Datasource as
// a slice of lakewatch_tf.ParserRef values.
// If the field is unknown or null, the boolean return value is false.
func (m *Datasource) GetParsers(ctx context.Context) ([]lakewatch_tf.ParserRef, bool) {
	if m.Parsers.IsNull() || m.Parsers.IsUnknown() {
		return nil, false
	}
	var v []lakewatch_tf.ParserRef
	d := m.Parsers.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParsers sets the value of the Parsers field in Datasource.
func (m *Datasource) SetParsers(ctx context.Context, v []lakewatch_tf.ParserRef) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["parsers"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Parsers = types.ListValueMust(t, vs)
}

// GetSchedule returns the value of the Schedule field in Datasource as
// a lakewatch_tf.Schedule value.
// If the field is unknown or null, the boolean return value is false.
func (m *Datasource) GetSchedule(ctx context.Context) (lakewatch_tf.Schedule, bool) {
	var e lakewatch_tf.Schedule
	if m.Schedule.IsNull() || m.Schedule.IsUnknown() {
		return e, false
	}
	var v lakewatch_tf.Schedule
	d := m.Schedule.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSchedule sets the value of the Schedule field in Datasource.
func (m *Datasource) SetSchedule(ctx context.Context, v lakewatch_tf.Schedule) {
	vs := v.ToObjectValue(ctx)
	m.Schedule = vs
}

// GetStatus returns the value of the Status field in Datasource as
// a lakewatch_tf.DatasourceStatus value.
// If the field is unknown or null, the boolean return value is false.
func (m *Datasource) GetStatus(ctx context.Context) (lakewatch_tf.DatasourceStatus, bool) {
	var e lakewatch_tf.DatasourceStatus
	if m.Status.IsNull() || m.Status.IsUnknown() {
		return e, false
	}
	var v lakewatch_tf.DatasourceStatus
	d := m.Status.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatus sets the value of the Status field in Datasource.
func (m *Datasource) SetStatus(ctx context.Context, v lakewatch_tf.DatasourceStatus) {
	vs := v.ToObjectValue(ctx)
	m.Status = vs
}

// GetUnityCatalogConnection returns the value of the UnityCatalogConnection field in Datasource as
// a lakewatch_tf.UnityCatalogConnection value.
// If the field is unknown or null, the boolean return value is false.
func (m *Datasource) GetUnityCatalogConnection(ctx context.Context) (lakewatch_tf.UnityCatalogConnection, bool) {
	var e lakewatch_tf.UnityCatalogConnection
	if m.UnityCatalogConnection.IsNull() || m.UnityCatalogConnection.IsUnknown() {
		return e, false
	}
	var v lakewatch_tf.UnityCatalogConnection
	d := m.UnityCatalogConnection.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUnityCatalogConnection sets the value of the UnityCatalogConnection field in Datasource.
func (m *Datasource) SetUnityCatalogConnection(ctx context.Context, v lakewatch_tf.UnityCatalogConnection) {
	vs := v.ToObjectValue(ctx)
	m.UnityCatalogConnection = vs
}

// GetUpdatedBy returns the value of the UpdatedBy field in Datasource as
// a lakewatch_tf.Principal value.
// If the field is unknown or null, the boolean return value is false.
func (m *Datasource) GetUpdatedBy(ctx context.Context) (lakewatch_tf.Principal, bool) {
	var e lakewatch_tf.Principal
	if m.UpdatedBy.IsNull() || m.UpdatedBy.IsUnknown() {
		return e, false
	}
	var v lakewatch_tf.Principal
	d := m.UpdatedBy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUpdatedBy sets the value of the UpdatedBy field in Datasource.
func (m *Datasource) SetUpdatedBy(ctx context.Context, v lakewatch_tf.Principal) {
	vs := v.ToObjectValue(ctx)
	m.UpdatedBy = vs
}

func (r *DatasourceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *DatasourceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, Datasource{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks lakewatch_datasource",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *DatasourceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *DatasourceResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Skip entirely on destroy (no plan state).
	if req.Plan.Raw.IsNull() {
		return
	}
	if r.Client == nil {
		return
	}
	tfschema.WorkspaceDriftDetection(ctx, r.Client, req, resp)
}

func (r *DatasourceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan Datasource
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var datasource lakewatch.Datasource

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &datasource)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := lakewatch.CreateDatasourceRequest{
		Datasource: datasource,
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

	response, err := client.Lakewatch.CreateDatasource(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create lakewatch_datasource", err.Error())
		return
	}

	var newState Datasource

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

func (r *DatasourceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState Datasource
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest lakewatch.GetDatasourceRequest
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
	response, err := client.Lakewatch.GetDatasource(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to get lakewatch_datasource", err.Error())
		return
	}

	var newState Datasource
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

func (r *DatasourceResource) update(ctx context.Context, plan Datasource, diags *diag.Diagnostics, state *tfsdk.State) {
	var datasource lakewatch.Datasource

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &datasource)...)
	if diags.HasError() {
		return
	}

	updateRequest := lakewatch.UpdateDatasourceRequest{
		Datasource: datasource,
		Name:       plan.Name.ValueString(),
		UpdateMask: *fieldmask.New(strings.Split("annotations,autoloader,compute_mode,description,existing_table,parsers,run_as_principal,schedule,unity_catalog_connection,warehouse_id", ",")),
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
	response, err := client.Lakewatch.UpdateDatasource(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update lakewatch_datasource", err.Error())
		return
	}

	var newState Datasource

	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *DatasourceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan Datasource
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *DatasourceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state Datasource
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest lakewatch.DeleteDatasourceRequest
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

	err := client.Lakewatch.DeleteDatasource(ctx, deleteRequest)
	if !declarative.IsDeleteError(err) {
		err = nil
	}
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete lakewatch_datasource", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &DatasourceResource{}

func (r *DatasourceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
