package tfschema

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// UpgradeProviderConfigListToObject is a StateUpgrader that converts the
// provider_config field from a list (the v1.113.0 ListNestedBlock encoding
// produced by Namespace_SdkV2) into the object encoding now used by
// Namespace (SingleNestedAttribute):
//
//   - "provider_config": []                          -> "provider_config": null
//   - "provider_config": null                        -> "provider_config": null
//   - "provider_config": [{...}]                     -> "provider_config": {...}
//
// It is the prior-version-0 upgrader for resources whose provider_config
// shape changed in #5582 (databricks_library, databricks_share,
// databricks_quality_monitor). All other fields pass through unchanged.
//
// Issue: https://github.com/databricks/terraform-provider-databricks/issues/5669
func UpgradeProviderConfigListToObject(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	if req.RawState == nil || req.RawState.JSON == nil {
		return
	}

	var state map[string]any
	if err := json.Unmarshal(req.RawState.JSON, &state); err != nil {
		resp.Diagnostics.Append(diag.NewErrorDiagnostic(
			"Failed to parse prior state JSON for upgrade",
			err.Error(),
		))
		return
	}

	if pc, present := state["provider_config"]; present {
		switch v := pc.(type) {
		case nil:
			// already null — no change.
		case []any:
			if len(v) == 0 {
				state["provider_config"] = nil
			} else if obj, ok := v[0].(map[string]any); ok {
				state["provider_config"] = obj
			} else {
				// list of unexpected element type — treat as null rather than fail.
				state["provider_config"] = nil
			}
		}
	}

	upgraded, err := json.Marshal(state)
	if err != nil {
		resp.Diagnostics.Append(diag.NewErrorDiagnostic(
			"Failed to re-encode upgraded state JSON",
			err.Error(),
		))
		return
	}

	schemaType := resp.State.Schema.Type().TerraformType(ctx)
	rawValue, err := (&tfprotov6.RawState{JSON: upgraded}).UnmarshalWithOpts(
		schemaType,
		tfprotov6.UnmarshalOpts{
			ValueFromJSONOpts: tftypes.ValueFromJSONOpts{
				IgnoreUndefinedAttributes: true,
			},
		},
	)
	if err != nil {
		resp.Diagnostics.Append(diag.NewErrorDiagnostic(
			"Failed to decode upgraded state with current schema",
			err.Error(),
		))
		return
	}

	resp.State.Raw = rawValue
}
