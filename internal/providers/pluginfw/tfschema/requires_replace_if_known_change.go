// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
package tfschema

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

// RequiresReplaceIfKnownChange is a RequiresReplaceIf predicate that triggers
// resource replacement only when both the prior state and the planned value
// are known and they differ. It tolerates null/unknown prior state and
// unknown plan values.
//
// This is the right plan modifier for fields that are conceptually immutable
// but whose prior state may legitimately be null - for example, fields that
// the API does not echo back on Read, which leaves them null in state after
// `terraform import`. A plain RequiresReplace() would treat the post-import
// null -> configured-value transition as a destructive change.
func RequiresReplaceIfKnownChange(_ context.Context, req planmodifier.StringRequest, resp *stringplanmodifier.RequiresReplaceIfFuncResponse) {
	if req.StateValue.IsNull() || req.StateValue.IsUnknown() {
		return
	}
	if req.PlanValue.IsUnknown() {
		return
	}
	if req.StateValue.ValueString() != req.PlanValue.ValueString() {
		resp.RequiresReplace = true
	}
}
