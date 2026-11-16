// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package lakewatch_datasource

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/lakewatch"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/lakewatch_tf"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceName = "lakewatch_datasource"

var _ datasource.DataSourceWithConfigure = &DatasourceDataSource{}

func DataSourceDatasource() datasource.DataSource {
	return &DatasourceDataSource{}
}

type DatasourceDataSource struct {
	Client *autogen.DatabricksClient
}

// ProviderConfigData contains the fields to configure the provider.
type ProviderConfigData struct {
	WorkspaceID types.String `tfsdk:"workspace_id"`
}

// ApplySchemaCustomizations applies the schema customizations to the ProviderConfig type.
func (r ProviderConfigData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_id"] = attrs["workspace_id"].SetOptional()
	attrs["workspace_id"] = attrs["workspace_id"].SetComputed()

	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.StringAttributeBuilder).AddValidator(stringvalidator.LengthAtLeast(1))
	return attrs
}

// ProviderConfigDataWorkspaceIDPlanModifier is plan modifier for the workspace_id field.
// Resource requires replacement if the workspace_id changes from one non-empty value to another.
func ProviderConfigDataWorkspaceIDPlanModifier(ctx context.Context, req planmodifier.StringRequest, resp *stringplanmodifier.RequiresReplaceIfFuncResponse) {
	// Require replacement if workspace_id changes from one non-empty value to another
	oldValue := req.StateValue.ValueString()
	newValue := req.PlanValue.ValueString()

	if oldValue != "" && newValue != "" && oldValue != newValue {
		resp.RequiresReplace = true
	}
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// ProviderConfigData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (r ProviderConfigData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProviderConfigData
// only implements ToObjectValue() and Type().
func (r ProviderConfigData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		r.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": r.WorkspaceID,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (r ProviderConfigData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.StringType,
		},
	}
}

// DatasourceData extends the main model with additional fields.
type DatasourceData struct {
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
	// Required on create and can be updated within the same catalog and schema.
	// Get synchronizes this field with the current Unity Catalog name, also
	// reported in `effective_datasource_name`.
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
	// access checks. Required and non-blank on every create and update request
	// except a rename-only PATCH; there is no default warehouse. The requester
	// must be able to use the warehouse. It is not stored, is not returned on
	// reads, and does not affect the compute that runs the ingestion job.
	WarehouseId        types.String `tfsdk:"warehouse_id"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// DatasourceData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m DatasourceData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
		"provider_config":          reflect.TypeOf(ProviderConfigData{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatasourceData
// only implements ToObjectValue() and Type().
func (m DatasourceData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"annotations":                m.Annotations,
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

			"provider_config": m.ProviderConfigData,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m DatasourceData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"annotations": basetypes.MapType{
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

			"provider_config": ProviderConfigData{}.Type(ctx),
		},
	}
}

func (m DatasourceData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["annotations"] = attrs["annotations"].SetComputed()
	attrs["autoloader"] = attrs["autoloader"].SetComputed()
	attrs["catalog_name"] = attrs["catalog_name"].SetComputed()
	attrs["compute_mode"] = attrs["compute_mode"].SetComputed()
	attrs["compute_workspace_id"] = attrs["compute_workspace_id"].SetComputed()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["created_by"] = attrs["created_by"].SetComputed()
	attrs["datasource_id"] = attrs["datasource_id"].SetComputed()
	attrs["datasource_name"] = attrs["datasource_name"].SetComputed()
	attrs["description"] = attrs["description"].SetComputed()
	attrs["effective_catalog_name"] = attrs["effective_catalog_name"].SetComputed()
	attrs["effective_datasource_name"] = attrs["effective_datasource_name"].SetComputed()
	attrs["effective_run_as_principal"] = attrs["effective_run_as_principal"].SetComputed()
	attrs["effective_schema_name"] = attrs["effective_schema_name"].SetComputed()
	attrs["etag"] = attrs["etag"].SetComputed()
	attrs["existing_table"] = attrs["existing_table"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["owner"] = attrs["owner"].SetComputed()
	attrs["parsers"] = attrs["parsers"].SetComputed()
	attrs["run_as_principal"] = attrs["run_as_principal"].SetComputed()
	attrs["schedule"] = attrs["schedule"].SetComputed()
	attrs["schema_name"] = attrs["schema_name"].SetComputed()
	attrs["status"] = attrs["status"].SetComputed()
	attrs["unity_catalog_connection"] = attrs["unity_catalog_connection"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["updated_by"] = attrs["updated_by"].SetComputed()
	attrs["version"] = attrs["version"].SetComputed()
	attrs["warehouse_id"] = attrs["warehouse_id"].SetComputed()

	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

func (r *DatasourceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *DatasourceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, DatasourceData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Datasource",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *DatasourceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *DatasourceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config DatasourceData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest lakewatch.GetDatasourceRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var namespace ProviderConfigData
	resp.Diagnostics.Append(config.ProviderConfigData.As(ctx, &namespace, basetypes.ObjectAsOptions{
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
		resp.Diagnostics.AddError("failed to get lakewatch_datasource", err.Error())
		return
	}

	var newState DatasourceData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Preserve provider_config from config so state.Set has the correct type info
	newState.ProviderConfigData = config.ProviderConfigData

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
