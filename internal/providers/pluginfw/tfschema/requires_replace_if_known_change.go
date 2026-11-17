// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
package tfschema

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

// CreationInputsAppliedKey records a successful Create or Update, including
// intentionally omitted inputs that cannot be distinguished from import nulls.
const CreationInputsAppliedKey = "creation_inputs_applied"

// RequiresReplaceIfKnownChange is a RequiresReplaceIf predicate that triggers
// resource replacement when a known prior value differs from the plan. It
// tolerates missing imported state, but not unknown plans: a reference to a
// replaced parent must also replace its child.
//
// This is the right plan modifier for fields that are conceptually immutable
// but whose prior state may legitimately be null - for example, fields that
// the API does not echo back on Read, which leaves them null in state after
// `terraform import`. A plain RequiresReplace() would treat the post-import
// null -> configured-value transition as a destructive change.
func RequiresReplaceIfKnownChange(ctx context.Context, req planmodifier.StringRequest, resp *stringplanmodifier.RequiresReplaceIfFuncResponse) {
	if req.StateValue.IsUnknown() {
		return
	}
	if req.StateValue.IsNull() {
		if req.ConfigValue.IsNull() || req.PlanValue.IsNull() {
			return
		}
		applied, diags := req.Private.GetKey(ctx, CreationInputsAppliedKey)
		resp.Diagnostics.Append(diags...)
		resp.RequiresReplace = string(applied) == "true"
		return
	}
	if !req.StateValue.Equal(req.PlanValue) {
		resp.RequiresReplace = true
	}
}
