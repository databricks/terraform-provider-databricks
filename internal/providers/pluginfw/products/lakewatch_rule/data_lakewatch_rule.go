// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package lakewatch_rule

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

const dataSourceName = "lakewatch_rule"

var _ datasource.DataSourceWithConfigure = &RuleDataSource{}

func DataSourceRule() datasource.DataSource {
	return &RuleDataSource{}
}

type RuleDataSource struct {
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

// RuleData extends the main model with additional fields.
type RuleData struct {
	// Free-form client-managed key/value annotations.
	Annotations types.Map `tfsdk:"annotations"`
	// Scheduled detection (also known as batch detection) input.
	Batch types.Object `tfsdk:"batch"`
	// Name of the Unity Catalog catalog that contains the rule. Required.
	// Cannot be changed after creation. Reads return the value as supplied; the
	// catalog's current name is reported in `effective_catalog_name`.
	CatalogName types.String `tfsdk:"catalog_name"`
	// Free-form rule category.
	Category types.String `tfsdk:"category"`
	// ID of the workspace whose compute runs this rule's jobs and owns its
	// associated resources, such as the SQL warehouse a scheduled detection
	// runs on. Server-assigned; a rule's warehouse can only be changed from
	// this workspace.
	ComputeWorkspaceId types.String `tfsdk:"compute_workspace_id"`
	// Immutable UUIDv4 identifying the authored content resource installed as
	// this rule.
	ContentResourceId types.String `tfsdk:"content_resource_id"`
	// Time the rule was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// Principal that created the rule.
	CreatedBy types.Object `tfsdk:"created_by"`
	// Signal deduplication mode.
	Deduplication types.String `tfsdk:"deduplication"`
	// Time window for deduplication, such as `1h`, `24h`, `7d`, or `automatic`.
	DeduplicationTimeWindow types.String `tfsdk:"deduplication_time_window"`
	// Free-form description.
	Description types.String `tfsdk:"description"`
	// Current name of the Unity Catalog catalog that contains the rule.
	// Reflects catalog renames, which do not change `catalog_name`. Unset until
	// resolved.
	EffectiveCatalogName types.String `tfsdk:"effective_catalog_name"`
	// Current UC leaf name of the rule (server-owned). Tracks UC renames so
	// callers can observe the live name without having to update the
	// client-owned `rule_name`; unset until resolved.
	EffectiveRuleName types.String `tfsdk:"effective_rule_name"`
	// The identity this rule's job runs as: the resolved `run_as_principal`, or
	// the requester's identity when it was omitted on create or cleared on
	// update. The value is a user email or service-principal application ID,
	// falling back to the numeric principal ID if that identity can no longer
	// be resolved. Populated on single-rule reads (get/create/update); list
	// responses omit it.
	EffectiveRunAsPrincipal types.String `tfsdk:"effective_run_as_principal"`
	// Current name of the Unity Catalog schema that contains the rule. Reflects
	// schema renames, which do not change `schema_name`. Unset until resolved.
	EffectiveSchemaName types.String `tfsdk:"effective_schema_name"`
	// Opaque entity tag for optimistic concurrency. Content-version ETags
	// change on user edits and remain stable across runtime status updates.
	Etag types.String `tfsdk:"etag"`
	// Signal quality for this rule.
	Fidelity types.String `tfsdk:"fidelity"`
	// Up to 64 MITRE ATT&CK mappings, producing at most 192 flattened alert
	// tags.
	Mitre types.List `tfsdk:"mitre"`
	// Resource name, in the form `rules/{rule_id}`. Assigned by the server on
	// create; required on update.
	Name types.String `tfsdk:"name"`
	// Long-form description of what this rule detects.
	Objective types.String `tfsdk:"objective"`
	// Reference to the content pack this rule was originally installed from.
	OriginalContentPack types.Object `tfsdk:"original_content_pack"`
	// What to record in each signal.
	Output types.Object `tfsdk:"output"`
	// Unity Catalog owner of the rule. Assigned by the server. Populated by
	// `GetRule`.
	Owner types.Object `tfsdk:"owner"`
	// Ordered reference URLs carried from the authored rule metadata. A maximum
	// of three references of up to 512 characters each is supported. Requests
	// exceeding these limits are rejected.
	References types.List `tfsdk:"references"`
	// Unity Catalog securable UUID of the rule, and the last segment of `name`.
	// Never changes, including when the rule's catalog or schema is renamed.
	RuleId types.String `tfsdk:"rule_id"`
	// Name of the rule within its Unity Catalog schema. Required on create; can
	// be updated within the same catalog and schema. Get synchronizes this
	// field with the current Unity Catalog name, also reported in
	// `effective_rule_name`.
	RuleName types.String `tfsdk:"rule_name"`
	// The principal whose identity this rule's job runs as, given as a user
	// email or service-principal application ID. When supplied on
	// create/update, Lakewatch resolves the principal, records it on the rule,
	// and creates the delegation to run as that principal — the requester
	// must be authorized to delegate to it, or the request is rejected. When
	// omitted on create, the requester's own identity is used. An update whose
	// mask does not include `run_as_principal` keeps the current identity, and
	// the requester must still be authorized to delegate to it; an editor who
	// is not must supply a run_as_principal they can delegate to (for example,
	// their own identity), or clear it, either of which reassigns the rule. An
	// update whose mask includes `run_as_principal` but supplies no value
	// clears it: the rule then runs as the requester, exactly as when it is
	// omitted on create. Reads return the value as supplied; the identity the
	// rule runs as is reported in `effective_run_as_principal`.
	RunAsPrincipal types.String `tfsdk:"run_as_principal"`
	// Long-form operational guidance carried from the authored rule metadata.
	// Supports up to 4096 characters. Requests exceeding this limit are
	// rejected. Supports Markdown; analysts and the investigation agent follow
	// it when triaging this rule's alerts.
	Runbook types.String `tfsdk:"runbook"`
	// When the rule executes.
	Schedule types.Object `tfsdk:"schedule"`
	// Name of the Unity Catalog schema that contains the rule. Required. Cannot
	// be changed after creation. Reads return the value as supplied; the
	// schema's current name is reported in `effective_schema_name`.
	SchemaName types.String `tfsdk:"schema_name"`
	// Free-form severity label, such as `Informational`, `Medium`, `High`, or
	// `Critical`.
	Severity types.String `tfsdk:"severity"`
	// Read-only lifecycle status.
	Status types.Object `tfsdk:"status"`
	// If set, this rule's signals do not open new alerts until the given time.
	SuppressUntilTime timetypes.RFC3339 `tfsdk:"suppress_until_time"`
	// Time-range configuration for scheduled detections. If unset, the
	// execution time range is extended by a fixed `15m` lookback.
	Timerange types.Object `tfsdk:"timerange"`
	// Time the rule was last updated.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
	// Principal that last updated the rule.
	UpdatedBy types.Object `tfsdk:"updated_by"`
	// Monotonic version, bumped on each update. Informational; use `etag` for
	// optimistic concurrency.
	Version            types.Int64  `tfsdk:"version"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// RuleData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m RuleData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"annotations":           reflect.TypeOf(types.String{}),
		"batch":                 reflect.TypeOf(lakewatch_tf.RuleInputBatch{}),
		"created_by":            reflect.TypeOf(lakewatch_tf.Principal{}),
		"mitre":                 reflect.TypeOf(lakewatch_tf.RuleMitreMapping{}),
		"original_content_pack": reflect.TypeOf(lakewatch_tf.ContentPackReference{}),
		"output":                reflect.TypeOf(lakewatch_tf.RuleOutput{}),
		"owner":                 reflect.TypeOf(lakewatch_tf.Principal{}),
		"references":            reflect.TypeOf(types.String{}),
		"schedule":              reflect.TypeOf(lakewatch_tf.RuleSchedule{}),
		"status":                reflect.TypeOf(lakewatch_tf.ResourceStatus{}),
		"timerange":             reflect.TypeOf(lakewatch_tf.RuleTimeRange{}),
		"updated_by":            reflect.TypeOf(lakewatch_tf.Principal{}),
		"provider_config":       reflect.TypeOf(ProviderConfigData{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RuleData
// only implements ToObjectValue() and Type().
func (m RuleData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"annotations":                m.Annotations,
			"batch":                      m.Batch,
			"catalog_name":               m.CatalogName,
			"category":                   m.Category,
			"compute_workspace_id":       m.ComputeWorkspaceId,
			"content_resource_id":        m.ContentResourceId,
			"create_time":                m.CreateTime,
			"created_by":                 m.CreatedBy,
			"deduplication":              m.Deduplication,
			"deduplication_time_window":  m.DeduplicationTimeWindow,
			"description":                m.Description,
			"effective_catalog_name":     m.EffectiveCatalogName,
			"effective_rule_name":        m.EffectiveRuleName,
			"effective_run_as_principal": m.EffectiveRunAsPrincipal,
			"effective_schema_name":      m.EffectiveSchemaName,
			"etag":                       m.Etag,
			"fidelity":                   m.Fidelity,
			"mitre":                      m.Mitre,
			"name":                       m.Name,
			"objective":                  m.Objective,
			"original_content_pack":      m.OriginalContentPack,
			"output":                     m.Output,
			"owner":                      m.Owner,
			"references":                 m.References,
			"rule_id":                    m.RuleId,
			"rule_name":                  m.RuleName,
			"run_as_principal":           m.RunAsPrincipal,
			"runbook":                    m.Runbook,
			"schedule":                   m.Schedule,
			"schema_name":                m.SchemaName,
			"severity":                   m.Severity,
			"status":                     m.Status,
			"suppress_until_time":        m.SuppressUntilTime,
			"timerange":                  m.Timerange,
			"update_time":                m.UpdateTime,
			"updated_by":                 m.UpdatedBy,
			"version":                    m.Version,

			"provider_config": m.ProviderConfigData,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m RuleData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"annotations": basetypes.MapType{
				ElemType: types.StringType,
			},
			"batch":                      lakewatch_tf.RuleInputBatch{}.Type(ctx),
			"catalog_name":               types.StringType,
			"category":                   types.StringType,
			"compute_workspace_id":       types.StringType,
			"content_resource_id":        types.StringType,
			"create_time":                timetypes.RFC3339{}.Type(ctx),
			"created_by":                 lakewatch_tf.Principal{}.Type(ctx),
			"deduplication":              types.StringType,
			"deduplication_time_window":  types.StringType,
			"description":                types.StringType,
			"effective_catalog_name":     types.StringType,
			"effective_rule_name":        types.StringType,
			"effective_run_as_principal": types.StringType,
			"effective_schema_name":      types.StringType,
			"etag":                       types.StringType,
			"fidelity":                   types.StringType,
			"mitre": basetypes.ListType{
				ElemType: lakewatch_tf.RuleMitreMapping{}.Type(ctx),
			},
			"name":                  types.StringType,
			"objective":             types.StringType,
			"original_content_pack": lakewatch_tf.ContentPackReference{}.Type(ctx),
			"output":                lakewatch_tf.RuleOutput{}.Type(ctx),
			"owner":                 lakewatch_tf.Principal{}.Type(ctx),
			"references": basetypes.ListType{
				ElemType: types.StringType,
			},
			"rule_id":             types.StringType,
			"rule_name":           types.StringType,
			"run_as_principal":    types.StringType,
			"runbook":             types.StringType,
			"schedule":            lakewatch_tf.RuleSchedule{}.Type(ctx),
			"schema_name":         types.StringType,
			"severity":            types.StringType,
			"status":              lakewatch_tf.ResourceStatus{}.Type(ctx),
			"suppress_until_time": timetypes.RFC3339{}.Type(ctx),
			"timerange":           lakewatch_tf.RuleTimeRange{}.Type(ctx),
			"update_time":         timetypes.RFC3339{}.Type(ctx),
			"updated_by":          lakewatch_tf.Principal{}.Type(ctx),
			"version":             types.Int64Type,

			"provider_config": ProviderConfigData{}.Type(ctx),
		},
	}
}

func (m RuleData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["annotations"] = attrs["annotations"].SetComputed()
	attrs["batch"] = attrs["batch"].SetComputed()
	attrs["catalog_name"] = attrs["catalog_name"].SetComputed()
	attrs["category"] = attrs["category"].SetComputed()
	attrs["compute_workspace_id"] = attrs["compute_workspace_id"].SetComputed()
	attrs["content_resource_id"] = attrs["content_resource_id"].SetComputed()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["created_by"] = attrs["created_by"].SetComputed()
	attrs["deduplication"] = attrs["deduplication"].SetComputed()
	attrs["deduplication_time_window"] = attrs["deduplication_time_window"].SetComputed()
	attrs["description"] = attrs["description"].SetComputed()
	attrs["effective_catalog_name"] = attrs["effective_catalog_name"].SetComputed()
	attrs["effective_rule_name"] = attrs["effective_rule_name"].SetComputed()
	attrs["effective_run_as_principal"] = attrs["effective_run_as_principal"].SetComputed()
	attrs["effective_schema_name"] = attrs["effective_schema_name"].SetComputed()
	attrs["etag"] = attrs["etag"].SetComputed()
	attrs["fidelity"] = attrs["fidelity"].SetComputed()
	attrs["mitre"] = attrs["mitre"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["objective"] = attrs["objective"].SetComputed()
	attrs["original_content_pack"] = attrs["original_content_pack"].SetComputed()
	attrs["output"] = attrs["output"].SetComputed()
	attrs["owner"] = attrs["owner"].SetComputed()
	attrs["references"] = attrs["references"].SetComputed()
	attrs["rule_id"] = attrs["rule_id"].SetComputed()
	attrs["rule_name"] = attrs["rule_name"].SetComputed()
	attrs["run_as_principal"] = attrs["run_as_principal"].SetComputed()
	attrs["runbook"] = attrs["runbook"].SetComputed()
	attrs["schedule"] = attrs["schedule"].SetComputed()
	attrs["schema_name"] = attrs["schema_name"].SetComputed()
	attrs["severity"] = attrs["severity"].SetComputed()
	attrs["status"] = attrs["status"].SetComputed()
	attrs["suppress_until_time"] = attrs["suppress_until_time"].SetComputed()
	attrs["timerange"] = attrs["timerange"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["updated_by"] = attrs["updated_by"].SetComputed()
	attrs["version"] = attrs["version"].SetComputed()

	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

func (r *RuleDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *RuleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, RuleData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Rule",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *RuleDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *RuleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config RuleData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest lakewatch.GetRuleRequest
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

	response, err := client.Lakewatch.GetRule(ctx, readRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to get lakewatch_rule", err.Error())
		return
	}

	var newState RuleData
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
