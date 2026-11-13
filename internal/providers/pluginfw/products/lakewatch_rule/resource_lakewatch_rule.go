// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package lakewatch_rule

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

const resourceName = "lakewatch_rule"

var _ resource.ResourceWithConfigure = &RuleResource{}
var _ resource.ResourceWithModifyPlan = &RuleResource{}

func ResourceRule() resource.Resource {
	return &RuleResource{}
}

type RuleResource struct {
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

// Rule extends the main model with additional fields.
type Rule struct {
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
	Version        types.Int64  `tfsdk:"version"`
	ProviderConfig types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// Rule struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m Rule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
		"provider_config":       reflect.TypeOf(ProviderConfig{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Rule
// only implements ToObjectValue() and Type().
func (m Rule) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"annotations": m.Annotations,
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

			"provider_config": m.ProviderConfig,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m Rule) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"annotations": basetypes.MapType{
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

			"provider_config": ProviderConfig{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *Rule) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Rule) {
	if !from.Batch.IsNull() && !from.Batch.IsUnknown() {
		if toBatch, ok := to.GetBatch(ctx); ok {
			if fromBatch, ok := from.GetBatch(ctx); ok {
				// Recursively sync the fields of Batch
				toBatch.SyncFieldsDuringCreateOrUpdate(ctx, fromBatch)
				to.SetBatch(ctx, toBatch)
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
	if !from.Mitre.IsNull() && !from.Mitre.IsUnknown() && to.Mitre.IsNull() && len(from.Mitre.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Mitre, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Mitre = from.Mitre
	}
	if !from.Mitre.IsNull() && !from.Mitre.IsUnknown() {
		if toMitre, ok := to.GetMitre(ctx); ok {
			if fromMitre, ok := from.GetMitre(ctx); ok {
				// Recursively sync the fields of each Mitre element by position.
				for i := range toMitre {
					if i < len(fromMitre) {
						toMitre[i].SyncFieldsDuringCreateOrUpdate(ctx, fromMitre[i])
					}
				}
				to.SetMitre(ctx, toMitre)
			}
		}
	}
	if !from.OriginalContentPack.IsNull() && !from.OriginalContentPack.IsUnknown() {
		if toOriginalContentPack, ok := to.GetOriginalContentPack(ctx); ok {
			if fromOriginalContentPack, ok := from.GetOriginalContentPack(ctx); ok {
				// Recursively sync the fields of OriginalContentPack
				toOriginalContentPack.SyncFieldsDuringCreateOrUpdate(ctx, fromOriginalContentPack)
				to.SetOriginalContentPack(ctx, toOriginalContentPack)
			}
		}
	}
	if !from.Output.IsNull() && !from.Output.IsUnknown() {
		if toOutput, ok := to.GetOutput(ctx); ok {
			if fromOutput, ok := from.GetOutput(ctx); ok {
				// Recursively sync the fields of Output
				toOutput.SyncFieldsDuringCreateOrUpdate(ctx, fromOutput)
				to.SetOutput(ctx, toOutput)
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
	if !from.References.IsNull() && !from.References.IsUnknown() && to.References.IsNull() && len(from.References.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for References, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.References = from.References
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
	if !from.Timerange.IsNull() && !from.Timerange.IsUnknown() {
		if toTimerange, ok := to.GetTimerange(ctx); ok {
			if fromTimerange, ok := from.GetTimerange(ctx); ok {
				// Recursively sync the fields of Timerange
				toTimerange.SyncFieldsDuringCreateOrUpdate(ctx, fromTimerange)
				to.SetTimerange(ctx, toTimerange)
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
	to.ProviderConfig = from.ProviderConfig

}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *Rule) SyncFieldsDuringRead(ctx context.Context, from Rule) {
	if !from.Batch.IsNull() && !from.Batch.IsUnknown() {
		if toBatch, ok := to.GetBatch(ctx); ok {
			if fromBatch, ok := from.GetBatch(ctx); ok {
				toBatch.SyncFieldsDuringRead(ctx, fromBatch)
				to.SetBatch(ctx, toBatch)
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
	if !from.Mitre.IsNull() && !from.Mitre.IsUnknown() && to.Mitre.IsNull() && len(from.Mitre.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Mitre, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Mitre = from.Mitre
	}
	if !from.Mitre.IsNull() && !from.Mitre.IsUnknown() {
		if toMitre, ok := to.GetMitre(ctx); ok {
			if fromMitre, ok := from.GetMitre(ctx); ok {
				for i := range toMitre {
					if i < len(fromMitre) {
						toMitre[i].SyncFieldsDuringRead(ctx, fromMitre[i])
					}
				}
				to.SetMitre(ctx, toMitre)
			}
		}
	}
	if !from.OriginalContentPack.IsNull() && !from.OriginalContentPack.IsUnknown() {
		if toOriginalContentPack, ok := to.GetOriginalContentPack(ctx); ok {
			if fromOriginalContentPack, ok := from.GetOriginalContentPack(ctx); ok {
				toOriginalContentPack.SyncFieldsDuringRead(ctx, fromOriginalContentPack)
				to.SetOriginalContentPack(ctx, toOriginalContentPack)
			}
		}
	}
	if !from.Output.IsNull() && !from.Output.IsUnknown() {
		if toOutput, ok := to.GetOutput(ctx); ok {
			if fromOutput, ok := from.GetOutput(ctx); ok {
				toOutput.SyncFieldsDuringRead(ctx, fromOutput)
				to.SetOutput(ctx, toOutput)
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
	if !from.References.IsNull() && !from.References.IsUnknown() && to.References.IsNull() && len(from.References.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for References, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.References = from.References
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
	if !from.Timerange.IsNull() && !from.Timerange.IsUnknown() {
		if toTimerange, ok := to.GetTimerange(ctx); ok {
			if fromTimerange, ok := from.GetTimerange(ctx); ok {
				toTimerange.SyncFieldsDuringRead(ctx, fromTimerange)
				to.SetTimerange(ctx, toTimerange)
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
	to.ProviderConfig = from.ProviderConfig

}

func (m Rule) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["annotations"] = attrs["annotations"].SetOptional()
	attrs["batch"] = attrs["batch"].SetOptional()
	attrs["catalog_name"] = attrs["catalog_name"].SetRequired()
	attrs["catalog_name"] = attrs["catalog_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["category"] = attrs["category"].SetOptional()
	attrs["compute_workspace_id"] = attrs["compute_workspace_id"].SetComputed()
	attrs["content_resource_id"] = attrs["content_resource_id"].SetOptional()
	attrs["content_resource_id"] = attrs["content_resource_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["created_by"] = attrs["created_by"].SetComputed()
	attrs["deduplication"] = attrs["deduplication"].SetOptional()
	attrs["deduplication_time_window"] = attrs["deduplication_time_window"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["effective_catalog_name"] = attrs["effective_catalog_name"].SetComputed()
	attrs["effective_rule_name"] = attrs["effective_rule_name"].SetComputed()
	attrs["effective_run_as_principal"] = attrs["effective_run_as_principal"].SetComputed()
	attrs["effective_schema_name"] = attrs["effective_schema_name"].SetComputed()
	attrs["etag"] = attrs["etag"].SetComputed()
	attrs["fidelity"] = attrs["fidelity"].SetRequired()
	attrs["mitre"] = attrs["mitre"].SetOptional()
	attrs["name"] = attrs["name"].SetComputed()
	attrs["objective"] = attrs["objective"].SetOptional()
	attrs["original_content_pack"] = attrs["original_content_pack"].SetComputed()
	attrs["output"] = attrs["output"].SetRequired()
	attrs["owner"] = attrs["owner"].SetComputed()
	attrs["references"] = attrs["references"].SetOptional()
	attrs["rule_id"] = attrs["rule_id"].SetComputed()
	attrs["rule_name"] = attrs["rule_name"].SetRequired()
	attrs["run_as_principal"] = attrs["run_as_principal"].SetOptional()
	attrs["runbook"] = attrs["runbook"].SetOptional()
	attrs["schedule"] = attrs["schedule"].SetRequired()
	attrs["schema_name"] = attrs["schema_name"].SetRequired()
	attrs["schema_name"] = attrs["schema_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["severity"] = attrs["severity"].SetRequired()
	attrs["status"] = attrs["status"].SetComputed()
	attrs["suppress_until_time"] = attrs["suppress_until_time"].SetOptional()
	attrs["timerange"] = attrs["timerange"].SetOptional()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["updated_by"] = attrs["updated_by"].SetComputed()
	attrs["version"] = attrs["version"].SetComputed()

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	attrs["provider_config"] = attrs["provider_config"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(tfschema.ProviderConfigPlanModifier{})

	return attrs
}

// GetAnnotations returns the value of the Annotations field in Rule as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *Rule) GetAnnotations(ctx context.Context) (map[string]types.String, bool) {
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

// SetAnnotations sets the value of the Annotations field in Rule.
func (m *Rule) SetAnnotations(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["annotations"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Annotations = types.MapValueMust(t, vs)
}

// GetBatch returns the value of the Batch field in Rule as
// a lakewatch_tf.RuleInputBatch value.
// If the field is unknown or null, the boolean return value is false.
func (m *Rule) GetBatch(ctx context.Context) (lakewatch_tf.RuleInputBatch, bool) {
	var e lakewatch_tf.RuleInputBatch
	if m.Batch.IsNull() || m.Batch.IsUnknown() {
		return e, false
	}
	var v lakewatch_tf.RuleInputBatch
	d := m.Batch.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBatch sets the value of the Batch field in Rule.
func (m *Rule) SetBatch(ctx context.Context, v lakewatch_tf.RuleInputBatch) {
	vs := v.ToObjectValue(ctx)
	m.Batch = vs
}

// GetCreatedBy returns the value of the CreatedBy field in Rule as
// a lakewatch_tf.Principal value.
// If the field is unknown or null, the boolean return value is false.
func (m *Rule) GetCreatedBy(ctx context.Context) (lakewatch_tf.Principal, bool) {
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

// SetCreatedBy sets the value of the CreatedBy field in Rule.
func (m *Rule) SetCreatedBy(ctx context.Context, v lakewatch_tf.Principal) {
	vs := v.ToObjectValue(ctx)
	m.CreatedBy = vs
}

// GetMitre returns the value of the Mitre field in Rule as
// a slice of lakewatch_tf.RuleMitreMapping values.
// If the field is unknown or null, the boolean return value is false.
func (m *Rule) GetMitre(ctx context.Context) ([]lakewatch_tf.RuleMitreMapping, bool) {
	if m.Mitre.IsNull() || m.Mitre.IsUnknown() {
		return nil, false
	}
	var v []lakewatch_tf.RuleMitreMapping
	d := m.Mitre.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMitre sets the value of the Mitre field in Rule.
func (m *Rule) SetMitre(ctx context.Context, v []lakewatch_tf.RuleMitreMapping) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["mitre"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Mitre = types.ListValueMust(t, vs)
}

// GetOriginalContentPack returns the value of the OriginalContentPack field in Rule as
// a lakewatch_tf.ContentPackReference value.
// If the field is unknown or null, the boolean return value is false.
func (m *Rule) GetOriginalContentPack(ctx context.Context) (lakewatch_tf.ContentPackReference, bool) {
	var e lakewatch_tf.ContentPackReference
	if m.OriginalContentPack.IsNull() || m.OriginalContentPack.IsUnknown() {
		return e, false
	}
	var v lakewatch_tf.ContentPackReference
	d := m.OriginalContentPack.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOriginalContentPack sets the value of the OriginalContentPack field in Rule.
func (m *Rule) SetOriginalContentPack(ctx context.Context, v lakewatch_tf.ContentPackReference) {
	vs := v.ToObjectValue(ctx)
	m.OriginalContentPack = vs
}

// GetOutput returns the value of the Output field in Rule as
// a lakewatch_tf.RuleOutput value.
// If the field is unknown or null, the boolean return value is false.
func (m *Rule) GetOutput(ctx context.Context) (lakewatch_tf.RuleOutput, bool) {
	var e lakewatch_tf.RuleOutput
	if m.Output.IsNull() || m.Output.IsUnknown() {
		return e, false
	}
	var v lakewatch_tf.RuleOutput
	d := m.Output.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOutput sets the value of the Output field in Rule.
func (m *Rule) SetOutput(ctx context.Context, v lakewatch_tf.RuleOutput) {
	vs := v.ToObjectValue(ctx)
	m.Output = vs
}

// GetOwner returns the value of the Owner field in Rule as
// a lakewatch_tf.Principal value.
// If the field is unknown or null, the boolean return value is false.
func (m *Rule) GetOwner(ctx context.Context) (lakewatch_tf.Principal, bool) {
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

// SetOwner sets the value of the Owner field in Rule.
func (m *Rule) SetOwner(ctx context.Context, v lakewatch_tf.Principal) {
	vs := v.ToObjectValue(ctx)
	m.Owner = vs
}

// GetReferences returns the value of the References field in Rule as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *Rule) GetReferences(ctx context.Context) ([]types.String, bool) {
	if m.References.IsNull() || m.References.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.References.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetReferences sets the value of the References field in Rule.
func (m *Rule) SetReferences(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["references"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.References = types.ListValueMust(t, vs)
}

// GetSchedule returns the value of the Schedule field in Rule as
// a lakewatch_tf.RuleSchedule value.
// If the field is unknown or null, the boolean return value is false.
func (m *Rule) GetSchedule(ctx context.Context) (lakewatch_tf.RuleSchedule, bool) {
	var e lakewatch_tf.RuleSchedule
	if m.Schedule.IsNull() || m.Schedule.IsUnknown() {
		return e, false
	}
	var v lakewatch_tf.RuleSchedule
	d := m.Schedule.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSchedule sets the value of the Schedule field in Rule.
func (m *Rule) SetSchedule(ctx context.Context, v lakewatch_tf.RuleSchedule) {
	vs := v.ToObjectValue(ctx)
	m.Schedule = vs
}

// GetStatus returns the value of the Status field in Rule as
// a lakewatch_tf.ResourceStatus value.
// If the field is unknown or null, the boolean return value is false.
func (m *Rule) GetStatus(ctx context.Context) (lakewatch_tf.ResourceStatus, bool) {
	var e lakewatch_tf.ResourceStatus
	if m.Status.IsNull() || m.Status.IsUnknown() {
		return e, false
	}
	var v lakewatch_tf.ResourceStatus
	d := m.Status.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStatus sets the value of the Status field in Rule.
func (m *Rule) SetStatus(ctx context.Context, v lakewatch_tf.ResourceStatus) {
	vs := v.ToObjectValue(ctx)
	m.Status = vs
}

// GetTimerange returns the value of the Timerange field in Rule as
// a lakewatch_tf.RuleTimeRange value.
// If the field is unknown or null, the boolean return value is false.
func (m *Rule) GetTimerange(ctx context.Context) (lakewatch_tf.RuleTimeRange, bool) {
	var e lakewatch_tf.RuleTimeRange
	if m.Timerange.IsNull() || m.Timerange.IsUnknown() {
		return e, false
	}
	var v lakewatch_tf.RuleTimeRange
	d := m.Timerange.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTimerange sets the value of the Timerange field in Rule.
func (m *Rule) SetTimerange(ctx context.Context, v lakewatch_tf.RuleTimeRange) {
	vs := v.ToObjectValue(ctx)
	m.Timerange = vs
}

// GetUpdatedBy returns the value of the UpdatedBy field in Rule as
// a lakewatch_tf.Principal value.
// If the field is unknown or null, the boolean return value is false.
func (m *Rule) GetUpdatedBy(ctx context.Context) (lakewatch_tf.Principal, bool) {
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

// SetUpdatedBy sets the value of the UpdatedBy field in Rule.
func (m *Rule) SetUpdatedBy(ctx context.Context, v lakewatch_tf.Principal) {
	vs := v.ToObjectValue(ctx)
	m.UpdatedBy = vs
}

func (r *RuleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *RuleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, Rule{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks lakewatch_rule",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *RuleResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *RuleResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Skip entirely on destroy (no plan state).
	if req.Plan.Raw.IsNull() {
		return
	}
	if r.Client == nil {
		return
	}
	tfschema.WorkspaceDriftDetection(ctx, r.Client, req, resp)
}

func (r *RuleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan Rule
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var rule lakewatch.Rule

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &rule)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := lakewatch.CreateRuleRequest{
		Rule: rule,
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

	response, err := client.Lakewatch.CreateRule(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create lakewatch_rule", err.Error())
		return
	}

	var newState Rule

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

func (r *RuleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState Rule
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest lakewatch.GetRuleRequest
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
	response, err := client.Lakewatch.GetRule(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to get lakewatch_rule", err.Error())
		return
	}

	var newState Rule
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

func (r *RuleResource) update(ctx context.Context, plan Rule, diags *diag.Diagnostics, state *tfsdk.State) {
	var rule lakewatch.Rule

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &rule)...)
	if diags.HasError() {
		return
	}

	updateRequest := lakewatch.UpdateRuleRequest{
		Rule:       rule,
		Name:       plan.Name.ValueString(),
		UpdateMask: *fieldmask.New(strings.Split("annotations,batch,category,deduplication,deduplication_time_window,description,fidelity,mitre,objective,output,references,rule_name,run_as_principal,runbook,schedule,severity,suppress_until_time,timerange", ",")),
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
	response, err := client.Lakewatch.UpdateRule(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update lakewatch_rule", err.Error())
		return
	}

	var newState Rule

	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *RuleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan Rule
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *RuleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state Rule
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest lakewatch.DeleteRuleRequest
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

	err := client.Lakewatch.DeleteRule(ctx, deleteRequest)
	if !declarative.IsDeleteError(err) {
		err = nil
	}
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete lakewatch_rule", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &RuleResource{}

func (r *RuleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
