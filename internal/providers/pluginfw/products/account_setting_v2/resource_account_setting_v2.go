// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package account_setting_v2

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/settingsv2"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/settingsv2_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const resourceName = "account_setting_v2"

var _ resource.ResourceWithConfigure = &SettingResource{}

func ResourceSetting() resource.Resource {
	return &SettingResource{}
}

type SettingResource struct {
	Client *autogen.DatabricksClient
}

// Setting extends the main model with additional fields.
type Setting struct {
	AibiDashboardEmbeddingAccessPolicy types.Object `tfsdk:"aibi_dashboard_embedding_access_policy"`

	AibiDashboardEmbeddingApprovedDomains types.Object `tfsdk:"aibi_dashboard_embedding_approved_domains"`

	AutomaticClusterUpdateWorkspace types.Object `tfsdk:"automatic_cluster_update_workspace"`

	BooleanVal types.Object `tfsdk:"boolean_val"`

	EffectiveAibiDashboardEmbeddingAccessPolicy types.Object `tfsdk:"effective_aibi_dashboard_embedding_access_policy"`

	EffectiveAibiDashboardEmbeddingApprovedDomains types.Object `tfsdk:"effective_aibi_dashboard_embedding_approved_domains"`

	EffectiveAutomaticClusterUpdateWorkspace types.Object `tfsdk:"effective_automatic_cluster_update_workspace"`

	EffectiveBooleanVal types.Object `tfsdk:"effective_boolean_val"`

	EffectiveIntegerVal types.Object `tfsdk:"effective_integer_val"`

	EffectivePersonalCompute types.Object `tfsdk:"effective_personal_compute"`

	EffectiveRestrictWorkspaceAdmins types.Object `tfsdk:"effective_restrict_workspace_admins"`

	EffectiveStringVal types.Object `tfsdk:"effective_string_val"`

	IntegerVal types.Object `tfsdk:"integer_val"`
	// Name of the setting.
	Name types.String `tfsdk:"name"`

	PersonalCompute types.Object `tfsdk:"personal_compute"`

	RestrictWorkspaceAdmins types.Object `tfsdk:"restrict_workspace_admins"`

	StringVal types.Object `tfsdk:"string_val"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// Setting struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m Setting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aibi_dashboard_embedding_access_policy":              reflect.TypeOf(settingsv2_tf.AibiDashboardEmbeddingAccessPolicy{}),
		"aibi_dashboard_embedding_approved_domains":           reflect.TypeOf(settingsv2_tf.AibiDashboardEmbeddingApprovedDomains{}),
		"automatic_cluster_update_workspace":                  reflect.TypeOf(settingsv2_tf.ClusterAutoRestartMessage{}),
		"boolean_val":                                         reflect.TypeOf(settingsv2_tf.BooleanMessage{}),
		"effective_aibi_dashboard_embedding_access_policy":    reflect.TypeOf(settingsv2_tf.AibiDashboardEmbeddingAccessPolicy{}),
		"effective_aibi_dashboard_embedding_approved_domains": reflect.TypeOf(settingsv2_tf.AibiDashboardEmbeddingApprovedDomains{}),
		"effective_automatic_cluster_update_workspace":        reflect.TypeOf(settingsv2_tf.ClusterAutoRestartMessage{}),
		"effective_boolean_val":                               reflect.TypeOf(settingsv2_tf.BooleanMessage{}),
		"effective_integer_val":                               reflect.TypeOf(settingsv2_tf.IntegerMessage{}),
		"effective_personal_compute":                          reflect.TypeOf(settingsv2_tf.PersonalComputeMessage{}),
		"effective_restrict_workspace_admins":                 reflect.TypeOf(settingsv2_tf.RestrictWorkspaceAdminsMessage{}),
		"effective_string_val":                                reflect.TypeOf(settingsv2_tf.StringMessage{}),
		"integer_val":                                         reflect.TypeOf(settingsv2_tf.IntegerMessage{}),
		"personal_compute":                                    reflect.TypeOf(settingsv2_tf.PersonalComputeMessage{}),
		"restrict_workspace_admins":                           reflect.TypeOf(settingsv2_tf.RestrictWorkspaceAdminsMessage{}),
		"string_val":                                          reflect.TypeOf(settingsv2_tf.StringMessage{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Setting
// only implements ToObjectValue() and Type().
func (m Setting) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"aibi_dashboard_embedding_access_policy": m.AibiDashboardEmbeddingAccessPolicy,
			"aibi_dashboard_embedding_approved_domains":           m.AibiDashboardEmbeddingApprovedDomains,
			"automatic_cluster_update_workspace":                  m.AutomaticClusterUpdateWorkspace,
			"boolean_val":                                         m.BooleanVal,
			"effective_aibi_dashboard_embedding_access_policy":    m.EffectiveAibiDashboardEmbeddingAccessPolicy,
			"effective_aibi_dashboard_embedding_approved_domains": m.EffectiveAibiDashboardEmbeddingApprovedDomains,
			"effective_automatic_cluster_update_workspace":        m.EffectiveAutomaticClusterUpdateWorkspace,
			"effective_boolean_val":                               m.EffectiveBooleanVal,
			"effective_integer_val":                               m.EffectiveIntegerVal,
			"effective_personal_compute":                          m.EffectivePersonalCompute,
			"effective_restrict_workspace_admins":                 m.EffectiveRestrictWorkspaceAdmins,
			"effective_string_val":                                m.EffectiveStringVal,
			"integer_val":                                         m.IntegerVal,
			"name":                                                m.Name,
			"personal_compute":                                    m.PersonalCompute,
			"restrict_workspace_admins":                           m.RestrictWorkspaceAdmins,
			"string_val":                                          m.StringVal,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m Setting) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"aibi_dashboard_embedding_access_policy": settingsv2_tf.AibiDashboardEmbeddingAccessPolicy{}.Type(ctx),
			"aibi_dashboard_embedding_approved_domains":           settingsv2_tf.AibiDashboardEmbeddingApprovedDomains{}.Type(ctx),
			"automatic_cluster_update_workspace":                  settingsv2_tf.ClusterAutoRestartMessage{}.Type(ctx),
			"boolean_val":                                         settingsv2_tf.BooleanMessage{}.Type(ctx),
			"effective_aibi_dashboard_embedding_access_policy":    settingsv2_tf.AibiDashboardEmbeddingAccessPolicy{}.Type(ctx),
			"effective_aibi_dashboard_embedding_approved_domains": settingsv2_tf.AibiDashboardEmbeddingApprovedDomains{}.Type(ctx),
			"effective_automatic_cluster_update_workspace":        settingsv2_tf.ClusterAutoRestartMessage{}.Type(ctx),
			"effective_boolean_val":                               settingsv2_tf.BooleanMessage{}.Type(ctx),
			"effective_integer_val":                               settingsv2_tf.IntegerMessage{}.Type(ctx),
			"effective_personal_compute":                          settingsv2_tf.PersonalComputeMessage{}.Type(ctx),
			"effective_restrict_workspace_admins":                 settingsv2_tf.RestrictWorkspaceAdminsMessage{}.Type(ctx),
			"effective_string_val":                                settingsv2_tf.StringMessage{}.Type(ctx),
			"integer_val":                                         settingsv2_tf.IntegerMessage{}.Type(ctx),
			"name":                                                types.StringType,
			"personal_compute":                                    settingsv2_tf.PersonalComputeMessage{}.Type(ctx),
			"restrict_workspace_admins":                           settingsv2_tf.RestrictWorkspaceAdminsMessage{}.Type(ctx),
			"string_val":                                          settingsv2_tf.StringMessage{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *Setting) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Setting) {
	if !from.AibiDashboardEmbeddingAccessPolicy.IsNull() && !from.AibiDashboardEmbeddingAccessPolicy.IsUnknown() {
		if toAibiDashboardEmbeddingAccessPolicy, ok := to.GetAibiDashboardEmbeddingAccessPolicy(ctx); ok {
			if fromAibiDashboardEmbeddingAccessPolicy, ok := from.GetAibiDashboardEmbeddingAccessPolicy(ctx); ok {
				// Recursively sync the fields of AibiDashboardEmbeddingAccessPolicy
				toAibiDashboardEmbeddingAccessPolicy.SyncFieldsDuringCreateOrUpdate(ctx, fromAibiDashboardEmbeddingAccessPolicy)
				to.SetAibiDashboardEmbeddingAccessPolicy(ctx, toAibiDashboardEmbeddingAccessPolicy)
			}
		}
	}
	if !from.AibiDashboardEmbeddingApprovedDomains.IsNull() && !from.AibiDashboardEmbeddingApprovedDomains.IsUnknown() {
		if toAibiDashboardEmbeddingApprovedDomains, ok := to.GetAibiDashboardEmbeddingApprovedDomains(ctx); ok {
			if fromAibiDashboardEmbeddingApprovedDomains, ok := from.GetAibiDashboardEmbeddingApprovedDomains(ctx); ok {
				// Recursively sync the fields of AibiDashboardEmbeddingApprovedDomains
				toAibiDashboardEmbeddingApprovedDomains.SyncFieldsDuringCreateOrUpdate(ctx, fromAibiDashboardEmbeddingApprovedDomains)
				to.SetAibiDashboardEmbeddingApprovedDomains(ctx, toAibiDashboardEmbeddingApprovedDomains)
			}
		}
	}
	if !from.AutomaticClusterUpdateWorkspace.IsNull() && !from.AutomaticClusterUpdateWorkspace.IsUnknown() {
		if toAutomaticClusterUpdateWorkspace, ok := to.GetAutomaticClusterUpdateWorkspace(ctx); ok {
			if fromAutomaticClusterUpdateWorkspace, ok := from.GetAutomaticClusterUpdateWorkspace(ctx); ok {
				// Recursively sync the fields of AutomaticClusterUpdateWorkspace
				toAutomaticClusterUpdateWorkspace.SyncFieldsDuringCreateOrUpdate(ctx, fromAutomaticClusterUpdateWorkspace)
				to.SetAutomaticClusterUpdateWorkspace(ctx, toAutomaticClusterUpdateWorkspace)
			}
		}
	}
	if !from.BooleanVal.IsNull() && !from.BooleanVal.IsUnknown() {
		if toBooleanVal, ok := to.GetBooleanVal(ctx); ok {
			if fromBooleanVal, ok := from.GetBooleanVal(ctx); ok {
				// Recursively sync the fields of BooleanVal
				toBooleanVal.SyncFieldsDuringCreateOrUpdate(ctx, fromBooleanVal)
				to.SetBooleanVal(ctx, toBooleanVal)
			}
		}
	}
	if !from.EffectiveAibiDashboardEmbeddingAccessPolicy.IsNull() && !from.EffectiveAibiDashboardEmbeddingAccessPolicy.IsUnknown() {
		if toEffectiveAibiDashboardEmbeddingAccessPolicy, ok := to.GetEffectiveAibiDashboardEmbeddingAccessPolicy(ctx); ok {
			if fromEffectiveAibiDashboardEmbeddingAccessPolicy, ok := from.GetEffectiveAibiDashboardEmbeddingAccessPolicy(ctx); ok {
				// Recursively sync the fields of EffectiveAibiDashboardEmbeddingAccessPolicy
				toEffectiveAibiDashboardEmbeddingAccessPolicy.SyncFieldsDuringCreateOrUpdate(ctx, fromEffectiveAibiDashboardEmbeddingAccessPolicy)
				to.SetEffectiveAibiDashboardEmbeddingAccessPolicy(ctx, toEffectiveAibiDashboardEmbeddingAccessPolicy)
			}
		}
	}
	if !from.EffectiveAibiDashboardEmbeddingApprovedDomains.IsNull() && !from.EffectiveAibiDashboardEmbeddingApprovedDomains.IsUnknown() {
		if toEffectiveAibiDashboardEmbeddingApprovedDomains, ok := to.GetEffectiveAibiDashboardEmbeddingApprovedDomains(ctx); ok {
			if fromEffectiveAibiDashboardEmbeddingApprovedDomains, ok := from.GetEffectiveAibiDashboardEmbeddingApprovedDomains(ctx); ok {
				// Recursively sync the fields of EffectiveAibiDashboardEmbeddingApprovedDomains
				toEffectiveAibiDashboardEmbeddingApprovedDomains.SyncFieldsDuringCreateOrUpdate(ctx, fromEffectiveAibiDashboardEmbeddingApprovedDomains)
				to.SetEffectiveAibiDashboardEmbeddingApprovedDomains(ctx, toEffectiveAibiDashboardEmbeddingApprovedDomains)
			}
		}
	}
	if !from.EffectiveAutomaticClusterUpdateWorkspace.IsNull() && !from.EffectiveAutomaticClusterUpdateWorkspace.IsUnknown() {
		if toEffectiveAutomaticClusterUpdateWorkspace, ok := to.GetEffectiveAutomaticClusterUpdateWorkspace(ctx); ok {
			if fromEffectiveAutomaticClusterUpdateWorkspace, ok := from.GetEffectiveAutomaticClusterUpdateWorkspace(ctx); ok {
				// Recursively sync the fields of EffectiveAutomaticClusterUpdateWorkspace
				toEffectiveAutomaticClusterUpdateWorkspace.SyncFieldsDuringCreateOrUpdate(ctx, fromEffectiveAutomaticClusterUpdateWorkspace)
				to.SetEffectiveAutomaticClusterUpdateWorkspace(ctx, toEffectiveAutomaticClusterUpdateWorkspace)
			}
		}
	}
	if !from.EffectiveBooleanVal.IsNull() && !from.EffectiveBooleanVal.IsUnknown() {
		if toEffectiveBooleanVal, ok := to.GetEffectiveBooleanVal(ctx); ok {
			if fromEffectiveBooleanVal, ok := from.GetEffectiveBooleanVal(ctx); ok {
				// Recursively sync the fields of EffectiveBooleanVal
				toEffectiveBooleanVal.SyncFieldsDuringCreateOrUpdate(ctx, fromEffectiveBooleanVal)
				to.SetEffectiveBooleanVal(ctx, toEffectiveBooleanVal)
			}
		}
	}
	if !from.EffectiveIntegerVal.IsNull() && !from.EffectiveIntegerVal.IsUnknown() {
		if toEffectiveIntegerVal, ok := to.GetEffectiveIntegerVal(ctx); ok {
			if fromEffectiveIntegerVal, ok := from.GetEffectiveIntegerVal(ctx); ok {
				// Recursively sync the fields of EffectiveIntegerVal
				toEffectiveIntegerVal.SyncFieldsDuringCreateOrUpdate(ctx, fromEffectiveIntegerVal)
				to.SetEffectiveIntegerVal(ctx, toEffectiveIntegerVal)
			}
		}
	}
	if !from.EffectivePersonalCompute.IsNull() && !from.EffectivePersonalCompute.IsUnknown() {
		if toEffectivePersonalCompute, ok := to.GetEffectivePersonalCompute(ctx); ok {
			if fromEffectivePersonalCompute, ok := from.GetEffectivePersonalCompute(ctx); ok {
				// Recursively sync the fields of EffectivePersonalCompute
				toEffectivePersonalCompute.SyncFieldsDuringCreateOrUpdate(ctx, fromEffectivePersonalCompute)
				to.SetEffectivePersonalCompute(ctx, toEffectivePersonalCompute)
			}
		}
	}
	if !from.EffectiveRestrictWorkspaceAdmins.IsNull() && !from.EffectiveRestrictWorkspaceAdmins.IsUnknown() {
		if toEffectiveRestrictWorkspaceAdmins, ok := to.GetEffectiveRestrictWorkspaceAdmins(ctx); ok {
			if fromEffectiveRestrictWorkspaceAdmins, ok := from.GetEffectiveRestrictWorkspaceAdmins(ctx); ok {
				// Recursively sync the fields of EffectiveRestrictWorkspaceAdmins
				toEffectiveRestrictWorkspaceAdmins.SyncFieldsDuringCreateOrUpdate(ctx, fromEffectiveRestrictWorkspaceAdmins)
				to.SetEffectiveRestrictWorkspaceAdmins(ctx, toEffectiveRestrictWorkspaceAdmins)
			}
		}
	}
	if !from.EffectiveStringVal.IsNull() && !from.EffectiveStringVal.IsUnknown() {
		if toEffectiveStringVal, ok := to.GetEffectiveStringVal(ctx); ok {
			if fromEffectiveStringVal, ok := from.GetEffectiveStringVal(ctx); ok {
				// Recursively sync the fields of EffectiveStringVal
				toEffectiveStringVal.SyncFieldsDuringCreateOrUpdate(ctx, fromEffectiveStringVal)
				to.SetEffectiveStringVal(ctx, toEffectiveStringVal)
			}
		}
	}
	if !from.IntegerVal.IsNull() && !from.IntegerVal.IsUnknown() {
		if toIntegerVal, ok := to.GetIntegerVal(ctx); ok {
			if fromIntegerVal, ok := from.GetIntegerVal(ctx); ok {
				// Recursively sync the fields of IntegerVal
				toIntegerVal.SyncFieldsDuringCreateOrUpdate(ctx, fromIntegerVal)
				to.SetIntegerVal(ctx, toIntegerVal)
			}
		}
	}
	if !from.PersonalCompute.IsNull() && !from.PersonalCompute.IsUnknown() {
		if toPersonalCompute, ok := to.GetPersonalCompute(ctx); ok {
			if fromPersonalCompute, ok := from.GetPersonalCompute(ctx); ok {
				// Recursively sync the fields of PersonalCompute
				toPersonalCompute.SyncFieldsDuringCreateOrUpdate(ctx, fromPersonalCompute)
				to.SetPersonalCompute(ctx, toPersonalCompute)
			}
		}
	}
	if !from.RestrictWorkspaceAdmins.IsNull() && !from.RestrictWorkspaceAdmins.IsUnknown() {
		if toRestrictWorkspaceAdmins, ok := to.GetRestrictWorkspaceAdmins(ctx); ok {
			if fromRestrictWorkspaceAdmins, ok := from.GetRestrictWorkspaceAdmins(ctx); ok {
				// Recursively sync the fields of RestrictWorkspaceAdmins
				toRestrictWorkspaceAdmins.SyncFieldsDuringCreateOrUpdate(ctx, fromRestrictWorkspaceAdmins)
				to.SetRestrictWorkspaceAdmins(ctx, toRestrictWorkspaceAdmins)
			}
		}
	}
	if !from.StringVal.IsNull() && !from.StringVal.IsUnknown() {
		if toStringVal, ok := to.GetStringVal(ctx); ok {
			if fromStringVal, ok := from.GetStringVal(ctx); ok {
				// Recursively sync the fields of StringVal
				toStringVal.SyncFieldsDuringCreateOrUpdate(ctx, fromStringVal)
				to.SetStringVal(ctx, toStringVal)
			}
		}
	}
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *Setting) SyncFieldsDuringRead(ctx context.Context, from Setting) {
	if !from.AibiDashboardEmbeddingAccessPolicy.IsNull() && !from.AibiDashboardEmbeddingAccessPolicy.IsUnknown() {
		if toAibiDashboardEmbeddingAccessPolicy, ok := to.GetAibiDashboardEmbeddingAccessPolicy(ctx); ok {
			if fromAibiDashboardEmbeddingAccessPolicy, ok := from.GetAibiDashboardEmbeddingAccessPolicy(ctx); ok {
				toAibiDashboardEmbeddingAccessPolicy.SyncFieldsDuringRead(ctx, fromAibiDashboardEmbeddingAccessPolicy)
				to.SetAibiDashboardEmbeddingAccessPolicy(ctx, toAibiDashboardEmbeddingAccessPolicy)
			}
		}
	}
	if !from.AibiDashboardEmbeddingApprovedDomains.IsNull() && !from.AibiDashboardEmbeddingApprovedDomains.IsUnknown() {
		if toAibiDashboardEmbeddingApprovedDomains, ok := to.GetAibiDashboardEmbeddingApprovedDomains(ctx); ok {
			if fromAibiDashboardEmbeddingApprovedDomains, ok := from.GetAibiDashboardEmbeddingApprovedDomains(ctx); ok {
				toAibiDashboardEmbeddingApprovedDomains.SyncFieldsDuringRead(ctx, fromAibiDashboardEmbeddingApprovedDomains)
				to.SetAibiDashboardEmbeddingApprovedDomains(ctx, toAibiDashboardEmbeddingApprovedDomains)
			}
		}
	}
	if !from.AutomaticClusterUpdateWorkspace.IsNull() && !from.AutomaticClusterUpdateWorkspace.IsUnknown() {
		if toAutomaticClusterUpdateWorkspace, ok := to.GetAutomaticClusterUpdateWorkspace(ctx); ok {
			if fromAutomaticClusterUpdateWorkspace, ok := from.GetAutomaticClusterUpdateWorkspace(ctx); ok {
				toAutomaticClusterUpdateWorkspace.SyncFieldsDuringRead(ctx, fromAutomaticClusterUpdateWorkspace)
				to.SetAutomaticClusterUpdateWorkspace(ctx, toAutomaticClusterUpdateWorkspace)
			}
		}
	}
	if !from.BooleanVal.IsNull() && !from.BooleanVal.IsUnknown() {
		if toBooleanVal, ok := to.GetBooleanVal(ctx); ok {
			if fromBooleanVal, ok := from.GetBooleanVal(ctx); ok {
				toBooleanVal.SyncFieldsDuringRead(ctx, fromBooleanVal)
				to.SetBooleanVal(ctx, toBooleanVal)
			}
		}
	}
	if !from.EffectiveAibiDashboardEmbeddingAccessPolicy.IsNull() && !from.EffectiveAibiDashboardEmbeddingAccessPolicy.IsUnknown() {
		if toEffectiveAibiDashboardEmbeddingAccessPolicy, ok := to.GetEffectiveAibiDashboardEmbeddingAccessPolicy(ctx); ok {
			if fromEffectiveAibiDashboardEmbeddingAccessPolicy, ok := from.GetEffectiveAibiDashboardEmbeddingAccessPolicy(ctx); ok {
				toEffectiveAibiDashboardEmbeddingAccessPolicy.SyncFieldsDuringRead(ctx, fromEffectiveAibiDashboardEmbeddingAccessPolicy)
				to.SetEffectiveAibiDashboardEmbeddingAccessPolicy(ctx, toEffectiveAibiDashboardEmbeddingAccessPolicy)
			}
		}
	}
	if !from.EffectiveAibiDashboardEmbeddingApprovedDomains.IsNull() && !from.EffectiveAibiDashboardEmbeddingApprovedDomains.IsUnknown() {
		if toEffectiveAibiDashboardEmbeddingApprovedDomains, ok := to.GetEffectiveAibiDashboardEmbeddingApprovedDomains(ctx); ok {
			if fromEffectiveAibiDashboardEmbeddingApprovedDomains, ok := from.GetEffectiveAibiDashboardEmbeddingApprovedDomains(ctx); ok {
				toEffectiveAibiDashboardEmbeddingApprovedDomains.SyncFieldsDuringRead(ctx, fromEffectiveAibiDashboardEmbeddingApprovedDomains)
				to.SetEffectiveAibiDashboardEmbeddingApprovedDomains(ctx, toEffectiveAibiDashboardEmbeddingApprovedDomains)
			}
		}
	}
	if !from.EffectiveAutomaticClusterUpdateWorkspace.IsNull() && !from.EffectiveAutomaticClusterUpdateWorkspace.IsUnknown() {
		if toEffectiveAutomaticClusterUpdateWorkspace, ok := to.GetEffectiveAutomaticClusterUpdateWorkspace(ctx); ok {
			if fromEffectiveAutomaticClusterUpdateWorkspace, ok := from.GetEffectiveAutomaticClusterUpdateWorkspace(ctx); ok {
				toEffectiveAutomaticClusterUpdateWorkspace.SyncFieldsDuringRead(ctx, fromEffectiveAutomaticClusterUpdateWorkspace)
				to.SetEffectiveAutomaticClusterUpdateWorkspace(ctx, toEffectiveAutomaticClusterUpdateWorkspace)
			}
		}
	}
	if !from.EffectiveBooleanVal.IsNull() && !from.EffectiveBooleanVal.IsUnknown() {
		if toEffectiveBooleanVal, ok := to.GetEffectiveBooleanVal(ctx); ok {
			if fromEffectiveBooleanVal, ok := from.GetEffectiveBooleanVal(ctx); ok {
				toEffectiveBooleanVal.SyncFieldsDuringRead(ctx, fromEffectiveBooleanVal)
				to.SetEffectiveBooleanVal(ctx, toEffectiveBooleanVal)
			}
		}
	}
	if !from.EffectiveIntegerVal.IsNull() && !from.EffectiveIntegerVal.IsUnknown() {
		if toEffectiveIntegerVal, ok := to.GetEffectiveIntegerVal(ctx); ok {
			if fromEffectiveIntegerVal, ok := from.GetEffectiveIntegerVal(ctx); ok {
				toEffectiveIntegerVal.SyncFieldsDuringRead(ctx, fromEffectiveIntegerVal)
				to.SetEffectiveIntegerVal(ctx, toEffectiveIntegerVal)
			}
		}
	}
	if !from.EffectivePersonalCompute.IsNull() && !from.EffectivePersonalCompute.IsUnknown() {
		if toEffectivePersonalCompute, ok := to.GetEffectivePersonalCompute(ctx); ok {
			if fromEffectivePersonalCompute, ok := from.GetEffectivePersonalCompute(ctx); ok {
				toEffectivePersonalCompute.SyncFieldsDuringRead(ctx, fromEffectivePersonalCompute)
				to.SetEffectivePersonalCompute(ctx, toEffectivePersonalCompute)
			}
		}
	}
	if !from.EffectiveRestrictWorkspaceAdmins.IsNull() && !from.EffectiveRestrictWorkspaceAdmins.IsUnknown() {
		if toEffectiveRestrictWorkspaceAdmins, ok := to.GetEffectiveRestrictWorkspaceAdmins(ctx); ok {
			if fromEffectiveRestrictWorkspaceAdmins, ok := from.GetEffectiveRestrictWorkspaceAdmins(ctx); ok {
				toEffectiveRestrictWorkspaceAdmins.SyncFieldsDuringRead(ctx, fromEffectiveRestrictWorkspaceAdmins)
				to.SetEffectiveRestrictWorkspaceAdmins(ctx, toEffectiveRestrictWorkspaceAdmins)
			}
		}
	}
	if !from.EffectiveStringVal.IsNull() && !from.EffectiveStringVal.IsUnknown() {
		if toEffectiveStringVal, ok := to.GetEffectiveStringVal(ctx); ok {
			if fromEffectiveStringVal, ok := from.GetEffectiveStringVal(ctx); ok {
				toEffectiveStringVal.SyncFieldsDuringRead(ctx, fromEffectiveStringVal)
				to.SetEffectiveStringVal(ctx, toEffectiveStringVal)
			}
		}
	}
	if !from.IntegerVal.IsNull() && !from.IntegerVal.IsUnknown() {
		if toIntegerVal, ok := to.GetIntegerVal(ctx); ok {
			if fromIntegerVal, ok := from.GetIntegerVal(ctx); ok {
				toIntegerVal.SyncFieldsDuringRead(ctx, fromIntegerVal)
				to.SetIntegerVal(ctx, toIntegerVal)
			}
		}
	}
	if !from.PersonalCompute.IsNull() && !from.PersonalCompute.IsUnknown() {
		if toPersonalCompute, ok := to.GetPersonalCompute(ctx); ok {
			if fromPersonalCompute, ok := from.GetPersonalCompute(ctx); ok {
				toPersonalCompute.SyncFieldsDuringRead(ctx, fromPersonalCompute)
				to.SetPersonalCompute(ctx, toPersonalCompute)
			}
		}
	}
	if !from.RestrictWorkspaceAdmins.IsNull() && !from.RestrictWorkspaceAdmins.IsUnknown() {
		if toRestrictWorkspaceAdmins, ok := to.GetRestrictWorkspaceAdmins(ctx); ok {
			if fromRestrictWorkspaceAdmins, ok := from.GetRestrictWorkspaceAdmins(ctx); ok {
				toRestrictWorkspaceAdmins.SyncFieldsDuringRead(ctx, fromRestrictWorkspaceAdmins)
				to.SetRestrictWorkspaceAdmins(ctx, toRestrictWorkspaceAdmins)
			}
		}
	}
	if !from.StringVal.IsNull() && !from.StringVal.IsUnknown() {
		if toStringVal, ok := to.GetStringVal(ctx); ok {
			if fromStringVal, ok := from.GetStringVal(ctx); ok {
				toStringVal.SyncFieldsDuringRead(ctx, fromStringVal)
				to.SetStringVal(ctx, toStringVal)
			}
		}
	}
}

func (m Setting) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aibi_dashboard_embedding_access_policy"] = attrs["aibi_dashboard_embedding_access_policy"].SetOptional()
	attrs["aibi_dashboard_embedding_approved_domains"] = attrs["aibi_dashboard_embedding_approved_domains"].SetOptional()
	attrs["automatic_cluster_update_workspace"] = attrs["automatic_cluster_update_workspace"].SetOptional()
	attrs["boolean_val"] = attrs["boolean_val"].SetOptional()
	attrs["effective_aibi_dashboard_embedding_access_policy"] = attrs["effective_aibi_dashboard_embedding_access_policy"].SetOptional()
	attrs["effective_aibi_dashboard_embedding_approved_domains"] = attrs["effective_aibi_dashboard_embedding_approved_domains"].SetOptional()
	attrs["effective_automatic_cluster_update_workspace"] = attrs["effective_automatic_cluster_update_workspace"].SetOptional()
	attrs["effective_boolean_val"] = attrs["effective_boolean_val"].SetComputed()
	attrs["effective_integer_val"] = attrs["effective_integer_val"].SetComputed()
	attrs["effective_personal_compute"] = attrs["effective_personal_compute"].SetOptional()
	attrs["effective_restrict_workspace_admins"] = attrs["effective_restrict_workspace_admins"].SetOptional()
	attrs["effective_string_val"] = attrs["effective_string_val"].SetComputed()
	attrs["integer_val"] = attrs["integer_val"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["personal_compute"] = attrs["personal_compute"].SetOptional()
	attrs["restrict_workspace_admins"] = attrs["restrict_workspace_admins"].SetOptional()
	attrs["string_val"] = attrs["string_val"].SetOptional()

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	return attrs
}

// GetAibiDashboardEmbeddingAccessPolicy returns the value of the AibiDashboardEmbeddingAccessPolicy field in Setting as
// a settingsv2_tf.AibiDashboardEmbeddingAccessPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (m *Setting) GetAibiDashboardEmbeddingAccessPolicy(ctx context.Context) (settingsv2_tf.AibiDashboardEmbeddingAccessPolicy, bool) {
	var e settingsv2_tf.AibiDashboardEmbeddingAccessPolicy
	if m.AibiDashboardEmbeddingAccessPolicy.IsNull() || m.AibiDashboardEmbeddingAccessPolicy.IsUnknown() {
		return e, false
	}
	var v settingsv2_tf.AibiDashboardEmbeddingAccessPolicy
	d := m.AibiDashboardEmbeddingAccessPolicy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAibiDashboardEmbeddingAccessPolicy sets the value of the AibiDashboardEmbeddingAccessPolicy field in Setting.
func (m *Setting) SetAibiDashboardEmbeddingAccessPolicy(ctx context.Context, v settingsv2_tf.AibiDashboardEmbeddingAccessPolicy) {
	vs := v.ToObjectValue(ctx)
	m.AibiDashboardEmbeddingAccessPolicy = vs
}

// GetAibiDashboardEmbeddingApprovedDomains returns the value of the AibiDashboardEmbeddingApprovedDomains field in Setting as
// a settingsv2_tf.AibiDashboardEmbeddingApprovedDomains value.
// If the field is unknown or null, the boolean return value is false.
func (m *Setting) GetAibiDashboardEmbeddingApprovedDomains(ctx context.Context) (settingsv2_tf.AibiDashboardEmbeddingApprovedDomains, bool) {
	var e settingsv2_tf.AibiDashboardEmbeddingApprovedDomains
	if m.AibiDashboardEmbeddingApprovedDomains.IsNull() || m.AibiDashboardEmbeddingApprovedDomains.IsUnknown() {
		return e, false
	}
	var v settingsv2_tf.AibiDashboardEmbeddingApprovedDomains
	d := m.AibiDashboardEmbeddingApprovedDomains.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAibiDashboardEmbeddingApprovedDomains sets the value of the AibiDashboardEmbeddingApprovedDomains field in Setting.
func (m *Setting) SetAibiDashboardEmbeddingApprovedDomains(ctx context.Context, v settingsv2_tf.AibiDashboardEmbeddingApprovedDomains) {
	vs := v.ToObjectValue(ctx)
	m.AibiDashboardEmbeddingApprovedDomains = vs
}

// GetAutomaticClusterUpdateWorkspace returns the value of the AutomaticClusterUpdateWorkspace field in Setting as
// a settingsv2_tf.ClusterAutoRestartMessage value.
// If the field is unknown or null, the boolean return value is false.
func (m *Setting) GetAutomaticClusterUpdateWorkspace(ctx context.Context) (settingsv2_tf.ClusterAutoRestartMessage, bool) {
	var e settingsv2_tf.ClusterAutoRestartMessage
	if m.AutomaticClusterUpdateWorkspace.IsNull() || m.AutomaticClusterUpdateWorkspace.IsUnknown() {
		return e, false
	}
	var v settingsv2_tf.ClusterAutoRestartMessage
	d := m.AutomaticClusterUpdateWorkspace.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAutomaticClusterUpdateWorkspace sets the value of the AutomaticClusterUpdateWorkspace field in Setting.
func (m *Setting) SetAutomaticClusterUpdateWorkspace(ctx context.Context, v settingsv2_tf.ClusterAutoRestartMessage) {
	vs := v.ToObjectValue(ctx)
	m.AutomaticClusterUpdateWorkspace = vs
}

// GetBooleanVal returns the value of the BooleanVal field in Setting as
// a settingsv2_tf.BooleanMessage value.
// If the field is unknown or null, the boolean return value is false.
func (m *Setting) GetBooleanVal(ctx context.Context) (settingsv2_tf.BooleanMessage, bool) {
	var e settingsv2_tf.BooleanMessage
	if m.BooleanVal.IsNull() || m.BooleanVal.IsUnknown() {
		return e, false
	}
	var v settingsv2_tf.BooleanMessage
	d := m.BooleanVal.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBooleanVal sets the value of the BooleanVal field in Setting.
func (m *Setting) SetBooleanVal(ctx context.Context, v settingsv2_tf.BooleanMessage) {
	vs := v.ToObjectValue(ctx)
	m.BooleanVal = vs
}

// GetEffectiveAibiDashboardEmbeddingAccessPolicy returns the value of the EffectiveAibiDashboardEmbeddingAccessPolicy field in Setting as
// a settingsv2_tf.AibiDashboardEmbeddingAccessPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (m *Setting) GetEffectiveAibiDashboardEmbeddingAccessPolicy(ctx context.Context) (settingsv2_tf.AibiDashboardEmbeddingAccessPolicy, bool) {
	var e settingsv2_tf.AibiDashboardEmbeddingAccessPolicy
	if m.EffectiveAibiDashboardEmbeddingAccessPolicy.IsNull() || m.EffectiveAibiDashboardEmbeddingAccessPolicy.IsUnknown() {
		return e, false
	}
	var v settingsv2_tf.AibiDashboardEmbeddingAccessPolicy
	d := m.EffectiveAibiDashboardEmbeddingAccessPolicy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveAibiDashboardEmbeddingAccessPolicy sets the value of the EffectiveAibiDashboardEmbeddingAccessPolicy field in Setting.
func (m *Setting) SetEffectiveAibiDashboardEmbeddingAccessPolicy(ctx context.Context, v settingsv2_tf.AibiDashboardEmbeddingAccessPolicy) {
	vs := v.ToObjectValue(ctx)
	m.EffectiveAibiDashboardEmbeddingAccessPolicy = vs
}

// GetEffectiveAibiDashboardEmbeddingApprovedDomains returns the value of the EffectiveAibiDashboardEmbeddingApprovedDomains field in Setting as
// a settingsv2_tf.AibiDashboardEmbeddingApprovedDomains value.
// If the field is unknown or null, the boolean return value is false.
func (m *Setting) GetEffectiveAibiDashboardEmbeddingApprovedDomains(ctx context.Context) (settingsv2_tf.AibiDashboardEmbeddingApprovedDomains, bool) {
	var e settingsv2_tf.AibiDashboardEmbeddingApprovedDomains
	if m.EffectiveAibiDashboardEmbeddingApprovedDomains.IsNull() || m.EffectiveAibiDashboardEmbeddingApprovedDomains.IsUnknown() {
		return e, false
	}
	var v settingsv2_tf.AibiDashboardEmbeddingApprovedDomains
	d := m.EffectiveAibiDashboardEmbeddingApprovedDomains.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveAibiDashboardEmbeddingApprovedDomains sets the value of the EffectiveAibiDashboardEmbeddingApprovedDomains field in Setting.
func (m *Setting) SetEffectiveAibiDashboardEmbeddingApprovedDomains(ctx context.Context, v settingsv2_tf.AibiDashboardEmbeddingApprovedDomains) {
	vs := v.ToObjectValue(ctx)
	m.EffectiveAibiDashboardEmbeddingApprovedDomains = vs
}

// GetEffectiveAutomaticClusterUpdateWorkspace returns the value of the EffectiveAutomaticClusterUpdateWorkspace field in Setting as
// a settingsv2_tf.ClusterAutoRestartMessage value.
// If the field is unknown or null, the boolean return value is false.
func (m *Setting) GetEffectiveAutomaticClusterUpdateWorkspace(ctx context.Context) (settingsv2_tf.ClusterAutoRestartMessage, bool) {
	var e settingsv2_tf.ClusterAutoRestartMessage
	if m.EffectiveAutomaticClusterUpdateWorkspace.IsNull() || m.EffectiveAutomaticClusterUpdateWorkspace.IsUnknown() {
		return e, false
	}
	var v settingsv2_tf.ClusterAutoRestartMessage
	d := m.EffectiveAutomaticClusterUpdateWorkspace.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveAutomaticClusterUpdateWorkspace sets the value of the EffectiveAutomaticClusterUpdateWorkspace field in Setting.
func (m *Setting) SetEffectiveAutomaticClusterUpdateWorkspace(ctx context.Context, v settingsv2_tf.ClusterAutoRestartMessage) {
	vs := v.ToObjectValue(ctx)
	m.EffectiveAutomaticClusterUpdateWorkspace = vs
}

// GetEffectiveBooleanVal returns the value of the EffectiveBooleanVal field in Setting as
// a settingsv2_tf.BooleanMessage value.
// If the field is unknown or null, the boolean return value is false.
func (m *Setting) GetEffectiveBooleanVal(ctx context.Context) (settingsv2_tf.BooleanMessage, bool) {
	var e settingsv2_tf.BooleanMessage
	if m.EffectiveBooleanVal.IsNull() || m.EffectiveBooleanVal.IsUnknown() {
		return e, false
	}
	var v settingsv2_tf.BooleanMessage
	d := m.EffectiveBooleanVal.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveBooleanVal sets the value of the EffectiveBooleanVal field in Setting.
func (m *Setting) SetEffectiveBooleanVal(ctx context.Context, v settingsv2_tf.BooleanMessage) {
	vs := v.ToObjectValue(ctx)
	m.EffectiveBooleanVal = vs
}

// GetEffectiveIntegerVal returns the value of the EffectiveIntegerVal field in Setting as
// a settingsv2_tf.IntegerMessage value.
// If the field is unknown or null, the boolean return value is false.
func (m *Setting) GetEffectiveIntegerVal(ctx context.Context) (settingsv2_tf.IntegerMessage, bool) {
	var e settingsv2_tf.IntegerMessage
	if m.EffectiveIntegerVal.IsNull() || m.EffectiveIntegerVal.IsUnknown() {
		return e, false
	}
	var v settingsv2_tf.IntegerMessage
	d := m.EffectiveIntegerVal.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveIntegerVal sets the value of the EffectiveIntegerVal field in Setting.
func (m *Setting) SetEffectiveIntegerVal(ctx context.Context, v settingsv2_tf.IntegerMessage) {
	vs := v.ToObjectValue(ctx)
	m.EffectiveIntegerVal = vs
}

// GetEffectivePersonalCompute returns the value of the EffectivePersonalCompute field in Setting as
// a settingsv2_tf.PersonalComputeMessage value.
// If the field is unknown or null, the boolean return value is false.
func (m *Setting) GetEffectivePersonalCompute(ctx context.Context) (settingsv2_tf.PersonalComputeMessage, bool) {
	var e settingsv2_tf.PersonalComputeMessage
	if m.EffectivePersonalCompute.IsNull() || m.EffectivePersonalCompute.IsUnknown() {
		return e, false
	}
	var v settingsv2_tf.PersonalComputeMessage
	d := m.EffectivePersonalCompute.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectivePersonalCompute sets the value of the EffectivePersonalCompute field in Setting.
func (m *Setting) SetEffectivePersonalCompute(ctx context.Context, v settingsv2_tf.PersonalComputeMessage) {
	vs := v.ToObjectValue(ctx)
	m.EffectivePersonalCompute = vs
}

// GetEffectiveRestrictWorkspaceAdmins returns the value of the EffectiveRestrictWorkspaceAdmins field in Setting as
// a settingsv2_tf.RestrictWorkspaceAdminsMessage value.
// If the field is unknown or null, the boolean return value is false.
func (m *Setting) GetEffectiveRestrictWorkspaceAdmins(ctx context.Context) (settingsv2_tf.RestrictWorkspaceAdminsMessage, bool) {
	var e settingsv2_tf.RestrictWorkspaceAdminsMessage
	if m.EffectiveRestrictWorkspaceAdmins.IsNull() || m.EffectiveRestrictWorkspaceAdmins.IsUnknown() {
		return e, false
	}
	var v settingsv2_tf.RestrictWorkspaceAdminsMessage
	d := m.EffectiveRestrictWorkspaceAdmins.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveRestrictWorkspaceAdmins sets the value of the EffectiveRestrictWorkspaceAdmins field in Setting.
func (m *Setting) SetEffectiveRestrictWorkspaceAdmins(ctx context.Context, v settingsv2_tf.RestrictWorkspaceAdminsMessage) {
	vs := v.ToObjectValue(ctx)
	m.EffectiveRestrictWorkspaceAdmins = vs
}

// GetEffectiveStringVal returns the value of the EffectiveStringVal field in Setting as
// a settingsv2_tf.StringMessage value.
// If the field is unknown or null, the boolean return value is false.
func (m *Setting) GetEffectiveStringVal(ctx context.Context) (settingsv2_tf.StringMessage, bool) {
	var e settingsv2_tf.StringMessage
	if m.EffectiveStringVal.IsNull() || m.EffectiveStringVal.IsUnknown() {
		return e, false
	}
	var v settingsv2_tf.StringMessage
	d := m.EffectiveStringVal.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveStringVal sets the value of the EffectiveStringVal field in Setting.
func (m *Setting) SetEffectiveStringVal(ctx context.Context, v settingsv2_tf.StringMessage) {
	vs := v.ToObjectValue(ctx)
	m.EffectiveStringVal = vs
}

// GetIntegerVal returns the value of the IntegerVal field in Setting as
// a settingsv2_tf.IntegerMessage value.
// If the field is unknown or null, the boolean return value is false.
func (m *Setting) GetIntegerVal(ctx context.Context) (settingsv2_tf.IntegerMessage, bool) {
	var e settingsv2_tf.IntegerMessage
	if m.IntegerVal.IsNull() || m.IntegerVal.IsUnknown() {
		return e, false
	}
	var v settingsv2_tf.IntegerMessage
	d := m.IntegerVal.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIntegerVal sets the value of the IntegerVal field in Setting.
func (m *Setting) SetIntegerVal(ctx context.Context, v settingsv2_tf.IntegerMessage) {
	vs := v.ToObjectValue(ctx)
	m.IntegerVal = vs
}

// GetPersonalCompute returns the value of the PersonalCompute field in Setting as
// a settingsv2_tf.PersonalComputeMessage value.
// If the field is unknown or null, the boolean return value is false.
func (m *Setting) GetPersonalCompute(ctx context.Context) (settingsv2_tf.PersonalComputeMessage, bool) {
	var e settingsv2_tf.PersonalComputeMessage
	if m.PersonalCompute.IsNull() || m.PersonalCompute.IsUnknown() {
		return e, false
	}
	var v settingsv2_tf.PersonalComputeMessage
	d := m.PersonalCompute.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPersonalCompute sets the value of the PersonalCompute field in Setting.
func (m *Setting) SetPersonalCompute(ctx context.Context, v settingsv2_tf.PersonalComputeMessage) {
	vs := v.ToObjectValue(ctx)
	m.PersonalCompute = vs
}

// GetRestrictWorkspaceAdmins returns the value of the RestrictWorkspaceAdmins field in Setting as
// a settingsv2_tf.RestrictWorkspaceAdminsMessage value.
// If the field is unknown or null, the boolean return value is false.
func (m *Setting) GetRestrictWorkspaceAdmins(ctx context.Context) (settingsv2_tf.RestrictWorkspaceAdminsMessage, bool) {
	var e settingsv2_tf.RestrictWorkspaceAdminsMessage
	if m.RestrictWorkspaceAdmins.IsNull() || m.RestrictWorkspaceAdmins.IsUnknown() {
		return e, false
	}
	var v settingsv2_tf.RestrictWorkspaceAdminsMessage
	d := m.RestrictWorkspaceAdmins.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRestrictWorkspaceAdmins sets the value of the RestrictWorkspaceAdmins field in Setting.
func (m *Setting) SetRestrictWorkspaceAdmins(ctx context.Context, v settingsv2_tf.RestrictWorkspaceAdminsMessage) {
	vs := v.ToObjectValue(ctx)
	m.RestrictWorkspaceAdmins = vs
}

// GetStringVal returns the value of the StringVal field in Setting as
// a settingsv2_tf.StringMessage value.
// If the field is unknown or null, the boolean return value is false.
func (m *Setting) GetStringVal(ctx context.Context) (settingsv2_tf.StringMessage, bool) {
	var e settingsv2_tf.StringMessage
	if m.StringVal.IsNull() || m.StringVal.IsUnknown() {
		return e, false
	}
	var v settingsv2_tf.StringMessage
	d := m.StringVal.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStringVal sets the value of the StringVal field in Setting.
func (m *Setting) SetStringVal(ctx context.Context, v settingsv2_tf.StringMessage) {
	vs := v.ToObjectValue(ctx)
	m.StringVal = vs
}

func (r *SettingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *SettingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, Setting{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks account_setting_v2",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *SettingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *SettingResource) update(ctx context.Context, plan Setting, diags *diag.Diagnostics, state *tfsdk.State) {
	client, clientDiags := r.Client.GetAccountClient()
	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}

	var setting settingsv2.Setting

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &setting)...)
	if diags.HasError() {
		return
	}

	updateRequest := settingsv2.PatchPublicAccountSettingRequest{
		Setting: setting,
		Name:    plan.Name.ValueString(),
	}

	response, err := client.SettingsV2.PatchPublicAccountSetting(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update account_setting_v2", err.Error())
		return
	}

	var newState Setting
	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *SettingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan Setting
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *SettingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetAccountClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var existingState Setting
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest settingsv2.GetPublicAccountSettingRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.SettingsV2.GetPublicAccountSetting(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get account_setting_v2", err.Error())
		return
	}

	var newState Setting
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *SettingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan Setting
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *SettingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

var _ resource.ResourceWithImportState = &SettingResource{}

func (r *SettingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
