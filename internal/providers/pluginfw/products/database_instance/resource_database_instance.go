// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package database_instance

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/database"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/database_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const resourceName = "database_instance"

var _ resource.ResourceWithConfigure = &DatabaseInstanceResource{}

func ResourceDatabaseInstance() resource.Resource {
	return &DatabaseInstanceResource{}
}

type DatabaseInstanceResource struct {
	Client *autogen.DatabricksClient
}

// DatabaseInstance extends the main model with additional fields.
type DatabaseInstance struct {
	// The sku of the instance. Valid values are "CU_1", "CU_2", "CU_4", "CU_8".
	Capacity types.String `tfsdk:"capacity"`
	// The refs of the child instances. This is only available if the instance
	// is parent instance.
	ChildInstanceRefs types.List `tfsdk:"child_instance_refs"`
	// The timestamp when the instance was created.
	CreationTime types.String `tfsdk:"creation_time"`
	// The email of the creator of the instance.
	Creator types.String `tfsdk:"creator"`
	// Custom tags associated with the instance. This field is only included on
	// create and update responses.
	CustomTags types.List `tfsdk:"custom_tags"`
	// Deprecated. The sku of the instance; this field will always match the
	// value of capacity.
	EffectiveCapacity types.String `tfsdk:"effective_capacity"`
	// The recorded custom tags associated with the instance.
	EffectiveCustomTags types.List `tfsdk:"effective_custom_tags"`
	// Whether the instance has PG native password login enabled.
	EffectiveEnablePgNativeLogin types.Bool `tfsdk:"effective_enable_pg_native_login"`
	// Whether secondaries serving read-only traffic are enabled. Defaults to
	// false.
	EffectiveEnableReadableSecondaries types.Bool `tfsdk:"effective_enable_readable_secondaries"`
	// The number of nodes in the instance, composed of 1 primary and 0 or more
	// secondaries. Defaults to 1 primary and 0 secondaries.
	EffectiveNodeCount types.Int64 `tfsdk:"effective_node_count"`
	// The retention window for the instance. This is the time window in days
	// for which the historical data is retained.
	EffectiveRetentionWindowInDays types.Int64 `tfsdk:"effective_retention_window_in_days"`
	// Whether the instance is stopped.
	EffectiveStopped types.Bool `tfsdk:"effective_stopped"`
	// The policy that is applied to the instance.
	EffectiveUsagePolicyId types.String `tfsdk:"effective_usage_policy_id"`
	// Whether to enable PG native password login on the instance. Defaults to
	// false.
	EnablePgNativeLogin types.Bool `tfsdk:"enable_pg_native_login"`
	// Whether to enable secondaries to serve read-only traffic. Defaults to
	// false.
	EnableReadableSecondaries types.Bool `tfsdk:"enable_readable_secondaries"`
	// The name of the instance. This is the unique identifier for the instance.
	Name types.String `tfsdk:"name"`
	// The number of nodes in the instance, composed of 1 primary and 0 or more
	// secondaries. Defaults to 1 primary and 0 secondaries. This field is input
	// only, see effective_node_count for the output.
	NodeCount types.Int64 `tfsdk:"node_count"`
	// The ref of the parent instance. This is only available if the instance is
	// child instance. Input: For specifying the parent instance to create a
	// child instance. Optional. Output: Only populated if provided as input to
	// create a child instance.
	ParentInstanceRef types.Object `tfsdk:"parent_instance_ref"`
	// The version of Postgres running on the instance.
	PgVersion types.String `tfsdk:"pg_version"`
	// Purge the resource on delete
	PurgeOnDelete types.Bool `tfsdk:"purge_on_delete"`
	// The DNS endpoint to connect to the instance for read only access. This is
	// only available if enable_readable_secondaries is true.
	ReadOnlyDns types.String `tfsdk:"read_only_dns"`
	// The DNS endpoint to connect to the instance for read+write access.
	ReadWriteDns types.String `tfsdk:"read_write_dns"`
	// The retention window for the instance. This is the time window in days
	// for which the historical data is retained. The default value is 7 days.
	// Valid values are 2 to 35 days.
	RetentionWindowInDays types.Int64 `tfsdk:"retention_window_in_days"`
	// The current state of the instance.
	State types.String `tfsdk:"state"`
	// Whether to stop the instance. An input only param, see effective_stopped
	// for the output.
	Stopped types.Bool `tfsdk:"stopped"`
	// An immutable UUID identifier for the instance.
	Uid types.String `tfsdk:"uid"`
	// The desired usage policy to associate with the instance.
	UsagePolicyId types.String `tfsdk:"usage_policy_id"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// DatabaseInstance struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m DatabaseInstance) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"child_instance_refs":   reflect.TypeOf(database_tf.DatabaseInstanceRef{}),
		"custom_tags":           reflect.TypeOf(database_tf.CustomTag{}),
		"effective_custom_tags": reflect.TypeOf(database_tf.CustomTag{}),
		"parent_instance_ref":   reflect.TypeOf(database_tf.DatabaseInstanceRef{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseInstance
// only implements ToObjectValue() and Type().
func (m DatabaseInstance) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"capacity": m.Capacity,
			"child_instance_refs":                   m.ChildInstanceRefs,
			"creation_time":                         m.CreationTime,
			"creator":                               m.Creator,
			"custom_tags":                           m.CustomTags,
			"effective_capacity":                    m.EffectiveCapacity,
			"effective_custom_tags":                 m.EffectiveCustomTags,
			"effective_enable_pg_native_login":      m.EffectiveEnablePgNativeLogin,
			"effective_enable_readable_secondaries": m.EffectiveEnableReadableSecondaries,
			"effective_node_count":                  m.EffectiveNodeCount,
			"effective_retention_window_in_days":    m.EffectiveRetentionWindowInDays,
			"effective_stopped":                     m.EffectiveStopped,
			"effective_usage_policy_id":             m.EffectiveUsagePolicyId,
			"enable_pg_native_login":                m.EnablePgNativeLogin,
			"enable_readable_secondaries":           m.EnableReadableSecondaries,
			"name":                                  m.Name,
			"node_count":                            m.NodeCount,
			"parent_instance_ref":                   m.ParentInstanceRef,
			"pg_version":                            m.PgVersion,
			"purge_on_delete":                       m.PurgeOnDelete,
			"read_only_dns":                         m.ReadOnlyDns,
			"read_write_dns":                        m.ReadWriteDns,
			"retention_window_in_days":              m.RetentionWindowInDays,
			"state":                                 m.State,
			"stopped":                               m.Stopped,
			"uid":                                   m.Uid,
			"usage_policy_id":                       m.UsagePolicyId,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m DatabaseInstance) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"capacity": types.StringType,
			"child_instance_refs": basetypes.ListType{
				ElemType: database_tf.DatabaseInstanceRef{}.Type(ctx),
			},
			"creation_time": types.StringType,
			"creator":       types.StringType,
			"custom_tags": basetypes.ListType{
				ElemType: database_tf.CustomTag{}.Type(ctx),
			},
			"effective_capacity": types.StringType,
			"effective_custom_tags": basetypes.ListType{
				ElemType: database_tf.CustomTag{}.Type(ctx),
			},
			"effective_enable_pg_native_login":      types.BoolType,
			"effective_enable_readable_secondaries": types.BoolType,
			"effective_node_count":                  types.Int64Type,
			"effective_retention_window_in_days":    types.Int64Type,
			"effective_stopped":                     types.BoolType,
			"effective_usage_policy_id":             types.StringType,
			"enable_pg_native_login":                types.BoolType,
			"enable_readable_secondaries":           types.BoolType,
			"name":                                  types.StringType,
			"node_count":                            types.Int64Type,
			"parent_instance_ref":                   database_tf.DatabaseInstanceRef{}.Type(ctx),
			"pg_version":                            types.StringType,
			"purge_on_delete":                       types.BoolType,
			"read_only_dns":                         types.StringType,
			"read_write_dns":                        types.StringType,
			"retention_window_in_days":              types.Int64Type,
			"state":                                 types.StringType,
			"stopped":                               types.BoolType,
			"uid":                                   types.StringType,
			"usage_policy_id":                       types.StringType,
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *DatabaseInstance) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabaseInstance) {
	if !from.ChildInstanceRefs.IsNull() && !from.ChildInstanceRefs.IsUnknown() && to.ChildInstanceRefs.IsNull() && len(from.ChildInstanceRefs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ChildInstanceRefs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ChildInstanceRefs = from.ChildInstanceRefs
	}
	if !from.CustomTags.IsUnknown() && !from.CustomTags.IsNull() {
		// CustomTags is an input only field and not returned by the service, so we keep the value from the prior state.
		to.CustomTags = from.CustomTags
	}
	if !from.CustomTags.IsNull() && !from.CustomTags.IsUnknown() && to.CustomTags.IsNull() && len(from.CustomTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CustomTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CustomTags = from.CustomTags
	}
	if !from.EffectiveCustomTags.IsNull() && !from.EffectiveCustomTags.IsUnknown() && to.EffectiveCustomTags.IsNull() && len(from.EffectiveCustomTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EffectiveCustomTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EffectiveCustomTags = from.EffectiveCustomTags
	}
	if !from.EnablePgNativeLogin.IsUnknown() && !from.EnablePgNativeLogin.IsNull() {
		// EnablePgNativeLogin is an input only field and not returned by the service, so we keep the value from the prior state.
		to.EnablePgNativeLogin = from.EnablePgNativeLogin
	}
	if !from.EnableReadableSecondaries.IsUnknown() && !from.EnableReadableSecondaries.IsNull() {
		// EnableReadableSecondaries is an input only field and not returned by the service, so we keep the value from the prior state.
		to.EnableReadableSecondaries = from.EnableReadableSecondaries
	}
	if !from.NodeCount.IsUnknown() && !from.NodeCount.IsNull() {
		// NodeCount is an input only field and not returned by the service, so we keep the value from the prior state.
		to.NodeCount = from.NodeCount
	}
	if !from.ParentInstanceRef.IsNull() && !from.ParentInstanceRef.IsUnknown() {
		if toParentInstanceRef, ok := to.GetParentInstanceRef(ctx); ok {
			if fromParentInstanceRef, ok := from.GetParentInstanceRef(ctx); ok {
				// Recursively sync the fields of ParentInstanceRef
				toParentInstanceRef.SyncFieldsDuringCreateOrUpdate(ctx, fromParentInstanceRef)
				to.SetParentInstanceRef(ctx, toParentInstanceRef)
			}
		}
	}
	to.PurgeOnDelete = from.PurgeOnDelete
	if !from.RetentionWindowInDays.IsUnknown() && !from.RetentionWindowInDays.IsNull() {
		// RetentionWindowInDays is an input only field and not returned by the service, so we keep the value from the prior state.
		to.RetentionWindowInDays = from.RetentionWindowInDays
	}
	if !from.Stopped.IsUnknown() && !from.Stopped.IsNull() {
		// Stopped is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Stopped = from.Stopped
	}
	if !from.UsagePolicyId.IsUnknown() && !from.UsagePolicyId.IsNull() {
		// UsagePolicyId is an input only field and not returned by the service, so we keep the value from the prior state.
		to.UsagePolicyId = from.UsagePolicyId
	}
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *DatabaseInstance) SyncFieldsDuringRead(ctx context.Context, from DatabaseInstance) {
	if !from.ChildInstanceRefs.IsNull() && !from.ChildInstanceRefs.IsUnknown() && to.ChildInstanceRefs.IsNull() && len(from.ChildInstanceRefs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ChildInstanceRefs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ChildInstanceRefs = from.ChildInstanceRefs
	}
	if !from.CustomTags.IsUnknown() && !from.CustomTags.IsNull() {
		// CustomTags is an input only field and not returned by the service, so we keep the value from the prior state.
		to.CustomTags = from.CustomTags
	}
	if !from.CustomTags.IsNull() && !from.CustomTags.IsUnknown() && to.CustomTags.IsNull() && len(from.CustomTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CustomTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CustomTags = from.CustomTags
	}
	if !from.EffectiveCustomTags.IsNull() && !from.EffectiveCustomTags.IsUnknown() && to.EffectiveCustomTags.IsNull() && len(from.EffectiveCustomTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EffectiveCustomTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EffectiveCustomTags = from.EffectiveCustomTags
	}
	if !from.EnablePgNativeLogin.IsUnknown() && !from.EnablePgNativeLogin.IsNull() {
		// EnablePgNativeLogin is an input only field and not returned by the service, so we keep the value from the prior state.
		to.EnablePgNativeLogin = from.EnablePgNativeLogin
	}
	if !from.EnableReadableSecondaries.IsUnknown() && !from.EnableReadableSecondaries.IsNull() {
		// EnableReadableSecondaries is an input only field and not returned by the service, so we keep the value from the prior state.
		to.EnableReadableSecondaries = from.EnableReadableSecondaries
	}
	if !from.NodeCount.IsUnknown() && !from.NodeCount.IsNull() {
		// NodeCount is an input only field and not returned by the service, so we keep the value from the prior state.
		to.NodeCount = from.NodeCount
	}
	if !from.ParentInstanceRef.IsNull() && !from.ParentInstanceRef.IsUnknown() {
		if toParentInstanceRef, ok := to.GetParentInstanceRef(ctx); ok {
			if fromParentInstanceRef, ok := from.GetParentInstanceRef(ctx); ok {
				toParentInstanceRef.SyncFieldsDuringRead(ctx, fromParentInstanceRef)
				to.SetParentInstanceRef(ctx, toParentInstanceRef)
			}
		}
	}
	to.PurgeOnDelete = from.PurgeOnDelete
	if !from.RetentionWindowInDays.IsUnknown() && !from.RetentionWindowInDays.IsNull() {
		// RetentionWindowInDays is an input only field and not returned by the service, so we keep the value from the prior state.
		to.RetentionWindowInDays = from.RetentionWindowInDays
	}
	if !from.Stopped.IsUnknown() && !from.Stopped.IsNull() {
		// Stopped is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Stopped = from.Stopped
	}
	if !from.UsagePolicyId.IsUnknown() && !from.UsagePolicyId.IsNull() {
		// UsagePolicyId is an input only field and not returned by the service, so we keep the value from the prior state.
		to.UsagePolicyId = from.UsagePolicyId
	}
}

func (m DatabaseInstance) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["capacity"] = attrs["capacity"].SetOptional()
	attrs["child_instance_refs"] = attrs["child_instance_refs"].SetComputed()
	attrs["creation_time"] = attrs["creation_time"].SetComputed()
	attrs["creator"] = attrs["creator"].SetComputed()
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["custom_tags"] = attrs["custom_tags"].SetComputed()
	attrs["custom_tags"] = attrs["custom_tags"].(tfschema.ListNestedAttributeBuilder).AddPlanModifier(listplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["effective_capacity"] = attrs["effective_capacity"].SetComputed()
	attrs["effective_custom_tags"] = attrs["effective_custom_tags"].SetComputed()
	attrs["effective_enable_pg_native_login"] = attrs["effective_enable_pg_native_login"].SetComputed()
	attrs["effective_enable_readable_secondaries"] = attrs["effective_enable_readable_secondaries"].SetComputed()
	attrs["effective_node_count"] = attrs["effective_node_count"].SetComputed()
	attrs["effective_retention_window_in_days"] = attrs["effective_retention_window_in_days"].SetComputed()
	attrs["effective_stopped"] = attrs["effective_stopped"].SetComputed()
	attrs["effective_usage_policy_id"] = attrs["effective_usage_policy_id"].SetComputed()
	attrs["enable_pg_native_login"] = attrs["enable_pg_native_login"].SetOptional()
	attrs["enable_pg_native_login"] = attrs["enable_pg_native_login"].SetComputed()
	attrs["enable_pg_native_login"] = attrs["enable_pg_native_login"].(tfschema.BoolAttributeBuilder).AddPlanModifier(boolplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["enable_readable_secondaries"] = attrs["enable_readable_secondaries"].SetOptional()
	attrs["enable_readable_secondaries"] = attrs["enable_readable_secondaries"].SetComputed()
	attrs["enable_readable_secondaries"] = attrs["enable_readable_secondaries"].(tfschema.BoolAttributeBuilder).AddPlanModifier(boolplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["node_count"] = attrs["node_count"].SetOptional()
	attrs["node_count"] = attrs["node_count"].SetComputed()
	attrs["node_count"] = attrs["node_count"].(tfschema.Int64AttributeBuilder).AddPlanModifier(int64planmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["parent_instance_ref"] = attrs["parent_instance_ref"].SetOptional()
	attrs["parent_instance_ref"] = attrs["parent_instance_ref"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["pg_version"] = attrs["pg_version"].SetComputed()
	attrs["read_only_dns"] = attrs["read_only_dns"].SetComputed()
	attrs["read_write_dns"] = attrs["read_write_dns"].SetComputed()
	attrs["retention_window_in_days"] = attrs["retention_window_in_days"].SetOptional()
	attrs["retention_window_in_days"] = attrs["retention_window_in_days"].SetComputed()
	attrs["retention_window_in_days"] = attrs["retention_window_in_days"].(tfschema.Int64AttributeBuilder).AddPlanModifier(int64planmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["state"] = attrs["state"].SetComputed()
	attrs["stopped"] = attrs["stopped"].SetOptional()
	attrs["stopped"] = attrs["stopped"].SetComputed()
	attrs["stopped"] = attrs["stopped"].(tfschema.BoolAttributeBuilder).AddPlanModifier(boolplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["uid"] = attrs["uid"].SetComputed()
	attrs["usage_policy_id"] = attrs["usage_policy_id"].SetOptional()
	attrs["usage_policy_id"] = attrs["usage_policy_id"].SetComputed()
	attrs["usage_policy_id"] = attrs["usage_policy_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["purge_on_delete"] = attrs["purge_on_delete"].SetOptional()

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	return attrs
}

// GetChildInstanceRefs returns the value of the ChildInstanceRefs field in DatabaseInstance as
// a slice of database_tf.DatabaseInstanceRef values.
// If the field is unknown or null, the boolean return value is false.
func (m *DatabaseInstance) GetChildInstanceRefs(ctx context.Context) ([]database_tf.DatabaseInstanceRef, bool) {
	if m.ChildInstanceRefs.IsNull() || m.ChildInstanceRefs.IsUnknown() {
		return nil, false
	}
	var v []database_tf.DatabaseInstanceRef
	d := m.ChildInstanceRefs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetChildInstanceRefs sets the value of the ChildInstanceRefs field in DatabaseInstance.
func (m *DatabaseInstance) SetChildInstanceRefs(ctx context.Context, v []database_tf.DatabaseInstanceRef) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["child_instance_refs"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ChildInstanceRefs = types.ListValueMust(t, vs)
}

// GetCustomTags returns the value of the CustomTags field in DatabaseInstance as
// a slice of database_tf.CustomTag values.
// If the field is unknown or null, the boolean return value is false.
func (m *DatabaseInstance) GetCustomTags(ctx context.Context) ([]database_tf.CustomTag, bool) {
	if m.CustomTags.IsNull() || m.CustomTags.IsUnknown() {
		return nil, false
	}
	var v []database_tf.CustomTag
	d := m.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in DatabaseInstance.
func (m *DatabaseInstance) SetCustomTags(ctx context.Context, v []database_tf.CustomTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CustomTags = types.ListValueMust(t, vs)
}

// GetEffectiveCustomTags returns the value of the EffectiveCustomTags field in DatabaseInstance as
// a slice of database_tf.CustomTag values.
// If the field is unknown or null, the boolean return value is false.
func (m *DatabaseInstance) GetEffectiveCustomTags(ctx context.Context) ([]database_tf.CustomTag, bool) {
	if m.EffectiveCustomTags.IsNull() || m.EffectiveCustomTags.IsUnknown() {
		return nil, false
	}
	var v []database_tf.CustomTag
	d := m.EffectiveCustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveCustomTags sets the value of the EffectiveCustomTags field in DatabaseInstance.
func (m *DatabaseInstance) SetEffectiveCustomTags(ctx context.Context, v []database_tf.CustomTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EffectiveCustomTags = types.ListValueMust(t, vs)
}

// GetParentInstanceRef returns the value of the ParentInstanceRef field in DatabaseInstance as
// a database_tf.DatabaseInstanceRef value.
// If the field is unknown or null, the boolean return value is false.
func (m *DatabaseInstance) GetParentInstanceRef(ctx context.Context) (database_tf.DatabaseInstanceRef, bool) {
	var e database_tf.DatabaseInstanceRef
	if m.ParentInstanceRef.IsNull() || m.ParentInstanceRef.IsUnknown() {
		return e, false
	}
	var v database_tf.DatabaseInstanceRef
	d := m.ParentInstanceRef.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParentInstanceRef sets the value of the ParentInstanceRef field in DatabaseInstance.
func (m *DatabaseInstance) SetParentInstanceRef(ctx context.Context, v database_tf.DatabaseInstanceRef) {
	vs := v.ToObjectValue(ctx)
	m.ParentInstanceRef = vs
}

func (r *DatabaseInstanceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *DatabaseInstanceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, DatabaseInstance{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks database_instance",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *DatabaseInstanceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *DatabaseInstanceResource) update(ctx context.Context, plan DatabaseInstance, diags *diag.Diagnostics, state *tfsdk.State) {
	var database_instance database.DatabaseInstance

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &database_instance)...)
	if diags.HasError() {
		return
	}

	updateRequest := database.UpdateDatabaseInstanceRequest{
		DatabaseInstance: database_instance,
		Name:             plan.Name.ValueString(),
		UpdateMask:       "capacity,custom_tags,enable_pg_native_login,enable_readable_secondaries,node_count,retention_window_in_days,stopped,usage_policy_id",
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}
	response, err := client.Database.UpdateDatabaseInstance(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update database_instance", err.Error())
		return
	}

	var newState DatabaseInstance

	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *DatabaseInstanceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan DatabaseInstance
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var database_instance database.DatabaseInstance

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &database_instance)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := database.CreateDatabaseInstanceRequest{
		DatabaseInstance: database_instance,
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.Database.CreateDatabaseInstance(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create database_instance", err.Error())
		return
	}

	var newState DatabaseInstance

	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response.Response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	waitResponse, err := response.Get()
	if err != nil {
		resp.Diagnostics.AddError("error waiting for database_instance to be ready", err.Error())
		return
	}

	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, waitResponse, &newState)...)

	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *DatabaseInstanceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState DatabaseInstance
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest database.GetDatabaseInstanceRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}
	response, err := client.Database.GetDatabaseInstance(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get database_instance", err.Error())
		return
	}

	var newState DatabaseInstance
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *DatabaseInstanceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan DatabaseInstance
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *DatabaseInstanceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state DatabaseInstance
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest database.DeleteDatabaseInstanceRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if !state.PurgeOnDelete.IsNull() && !state.PurgeOnDelete.IsUnknown() {
		deleteRequest.Purge = state.PurgeOnDelete.ValueBool()
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := client.Database.DeleteDatabaseInstance(ctx, deleteRequest)
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete database_instance", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &DatabaseInstanceResource{}

func (r *DatabaseInstanceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
