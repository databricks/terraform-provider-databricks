// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
package tfschema

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// AdoptImportedCreateFields saves missing creation inputs without an API update
// when they are the only change. It returns true when the request is handled,
// including diagnostic errors. Other changes must use the normal update path.
func AdoptImportedCreateFields(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse, fieldNames ...string) bool {
	if req.State.Raw.IsNull() || req.Plan.Raw.IsNull() {
		return false
	}
	candidate := req.State
	adopted := false
	for _, name := range fieldNames {
		var prior, planned types.String
		fieldPath := path.Root(name)
		resp.Diagnostics.Append(req.State.GetAttribute(ctx, fieldPath, &prior)...)
		resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, fieldPath, &planned)...)
		if resp.Diagnostics.HasError() {
			return true
		}
		if !prior.IsNull() || planned.IsNull() || planned.IsUnknown() {
			continue
		}
		resp.Diagnostics.Append(candidate.SetAttribute(ctx, fieldPath, planned)...)
		if resp.Diagnostics.HasError() {
			return true
		}
		adopted = true
	}
	if !adopted {
		return false
	}

	if !knownPlanMatchesState(req.Plan.Raw, candidate.Raw) {
		return false
	}
	resp.State = candidate
	return true
}

// Unknown computed values may be filled from refreshed state. Set elements
// must be matched by their known contents, not by paths containing unknowns.
func knownPlanMatchesState(plan, state tftypes.Value) bool {
	if !plan.Type().Equal(state.Type()) {
		return false
	}
	if !plan.IsKnown() {
		return true
	}
	if plan.IsFullyKnown() || plan.IsNull() || state.IsNull() || !state.IsKnown() {
		return plan.Equal(state)
	}
	switch plan.Type().(type) {
	case tftypes.Object, tftypes.Map:
		var planned, actual map[string]tftypes.Value
		if plan.As(&planned) != nil || state.As(&actual) != nil || len(planned) != len(actual) {
			return false
		}
		for key, value := range planned {
			prior, ok := actual[key]
			if !ok || !knownPlanMatchesState(value, prior) {
				return false
			}
		}
		return true
	case tftypes.List, tftypes.Tuple, tftypes.Set:
		var planned, actual []tftypes.Value
		if plan.As(&planned) != nil || state.As(&actual) != nil || len(planned) != len(actual) {
			return false
		}
		if _, set := plan.Type().(tftypes.Set); set {
			return knownSetMatchesState(planned, actual)
		}
		for i, value := range planned {
			if !knownPlanMatchesState(value, actual[i]) {
				return false
			}
		}
		return true
	default:
		return false
	}
}

func knownSetMatchesState(planned, actual []tftypes.Value) bool {
	// Find a one-to-one match. Reassign earlier matches when unknown fields
	// allow multiple candidates, so neither ordering nor element reuse wins.
	matches := make(map[int]int, len(actual))
	var match func(int, map[int]bool) bool
	match = func(i int, visited map[int]bool) bool {
		for j, value := range actual {
			if visited[j] || !knownPlanMatchesState(planned[i], value) {
				continue
			}
			visited[j] = true
			previous, used := matches[j]
			if !used || match(previous, visited) {
				matches[j] = i
				return true
			}
		}
		return false
	}
	for i := range planned {
		if !match(i, make(map[int]bool)) {
			return false
		}
	}
	return true
}
