// Package stateassert evaluates `terraform show -json` output against
// the v2-mode `assert:` block declared on a test.yaml step.
//
// The package is the runner's hook for "after a successful command,
// look at the resulting state and report structured failures". It is
// deliberately simple: spawn `terraform show -json` (the runner has
// already located the binary), parse the resulting state, and walk
// `values.root_module.resources[*]` for each assertion. Failures are
// returned as `result.AssertionFailure` so the runner can surface them
// uniformly with the rest of the step result (DESIGN.md §17.5 / §17.8).
//
// Three classes of failure exist:
//
//   - Resource not present when `present: true` (or omitted) was
//     expected. Field == "", Reason == "expected present, not found in
//     state".
//   - Resource present when `present: false` was expected. Field == "",
//     Reason == "expected absent, found in state".
//   - Per-attribute value mismatch. Field == the attribute key (or
//     dot-path), Reason == "value mismatch", Expected/Actual filled
//     from YAML and JSON sides respectively.
//
// Numerics are normalized to float64 before comparison because YAML's
// untagged-int decoder produces `int` while `terraform show -json`'s
// JSON decoder produces `float64`. Without coercion a YAML `3600`
// would never match a JSON `3600` despite being semantically equal.
//
// Sensitive attributes are reported as `"(sensitive)"` by terraform
// and treated as a hard mismatch — DESIGN.md §17.7 rule 7 says the
// runner should reject these at parse time eventually, but until
// then we surface a clear failure so the user knows what happened.
package stateassert

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"reflect"
	"sort"
	"strings"

	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/config"
	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/result"
)

// Run is the package entry point. Spawns `terraform show -json` in
// workdir with env (which must already include TF_CLI_CONFIG_FILE +
// DATABRICKS_CONFIG_PROFILE — same shape the runner hands to tfexec),
// parses the output, and evaluates each assertion.
//
// The error return is reserved for infrastructure failures (terraform
// failed to launch, returned non-zero, output didn't parse). Per-
// assertion mismatches go into the []AssertionFailure return value
// and never produce an error — the caller decides whether to flip
// step.Status to fail.
//
// Assertions are evaluated in YAML order; within a single Assertion,
// missing-resource and per-attribute checks run sequentially. Run
// never short-circuits — collect-all-failures is the documented
// invariant (DESIGN.md §17.5).
func Run(ctx context.Context, workdir, terraformBin string, env map[string]string, assertions []config.Assertion) ([]result.AssertionFailure, error) {
	if len(assertions) == 0 {
		return nil, nil
	}
	state, err := loadState(ctx, workdir, terraformBin, env)
	if err != nil {
		return nil, err
	}
	resources := indexResources(state)
	failures := []result.AssertionFailure{}
	for _, a := range assertions {
		failures = append(failures, evaluate(resources, a)...)
	}
	return failures, nil
}

// loadState invokes `terraform show -json` and unmarshals the output.
// Only the fields under `values.root_module.resources` (and nested
// `child_modules`) are relevant; everything else is opaque.
func loadState(ctx context.Context, workdir, terraformBin string, env map[string]string) (*tfState, error) {
	cmd := exec.CommandContext(ctx, terraformBin, "show", "-json")
	cmd.Dir = workdir
	cmd.Env = mapToSliceEnv(env)
	out, err := cmd.Output()
	if err != nil {
		var ee *exec.ExitError
		if errors.As(err, &ee) {
			return nil, fmt.Errorf("stateassert: terraform show -json: %w (stderr: %s)", err, string(ee.Stderr))
		}
		return nil, fmt.Errorf("stateassert: terraform show -json: %w", err)
	}
	var st tfState
	if err := json.Unmarshal(out, &st); err != nil {
		return nil, fmt.Errorf("stateassert: parse terraform show -json output: %w", err)
	}
	return &st, nil
}

// tfState is the partial schema we care about from `terraform show
// -json`. Many other fields are present in the real output; we ignore
// them with json's default skip-unknown behaviour.
type tfState struct {
	Values struct {
		RootModule tfModule `json:"root_module"`
	} `json:"values"`
}

type tfModule struct {
	Resources    []tfResource `json:"resources"`
	ChildModules []tfModule   `json:"child_modules"`
}

type tfResource struct {
	Address string         `json:"address"`
	Values  map[string]any `json:"values"`
}

// indexResources flattens the module tree into an address-keyed map.
// Module-scoped addresses (`module.X.Y.Z`) are exposed in the index
// but the v2-launch validation rejects them at parse time — included
// here so a future `module.` address support doesn't require a
// stateassert change.
func indexResources(s *tfState) map[string]map[string]any {
	out := map[string]map[string]any{}
	walkModule(out, s.Values.RootModule)
	return out
}

func walkModule(out map[string]map[string]any, m tfModule) {
	for _, r := range m.Resources {
		out[r.Address] = r.Values
	}
	for _, child := range m.ChildModules {
		walkModule(out, child)
	}
}

// evaluate runs one assertion against the indexed state, returning a
// slice of AssertionFailures (zero, one, or many). The slice form
// keeps the per-attribute fan-out legible: an assertion with three
// attrs that all mismatched produces three failures rather than one
// concatenated string.
func evaluate(resources map[string]map[string]any, a config.Assertion) []result.AssertionFailure {
	values, found := resources[a.Resource]
	switch {
	case a.PresentValue() && !found:
		return []result.AssertionFailure{{
			Address: a.Resource,
			Reason:  "expected present, not found in state",
		}}
	case !a.PresentValue() && found:
		return []result.AssertionFailure{{
			Address: a.Resource,
			Reason:  "expected absent, found in state",
		}}
	case !a.PresentValue() && !found:
		return nil
	}
	return evaluateAttrs(a.Resource, values, a.Attrs)
}

// evaluateAttrs walks each (key, expected) pair in the YAML attrs
// map and reports a failure for any mismatch. Iteration order is
// sorted for stable test output.
func evaluateAttrs(address string, values map[string]any, attrs map[string]any) []result.AssertionFailure {
	if len(attrs) == 0 {
		return nil
	}
	keys := make([]string, 0, len(attrs))
	for k := range attrs {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var failures []result.AssertionFailure
	for _, k := range keys {
		expected := attrs[k]
		got, ok := dotWalk(values, k)
		if !ok {
			failures = append(failures, result.AssertionFailure{
				Address: address, Field: k,
				Reason:   "attribute not found",
				Expected: expected,
			})
			continue
		}
		if isSensitiveSentinel(got) {
			failures = append(failures, result.AssertionFailure{
				Address: address, Field: k,
				Reason: "attribute is marked sensitive in state — assert against a non-sensitive proxy (DESIGN.md §17.6)",
				Actual: got,
			})
			continue
		}
		if !equal(expected, got) {
			failures = append(failures, result.AssertionFailure{
				Address: address, Field: k,
				Reason:   "value mismatch",
				Expected: expected,
				Actual:   got,
			})
		}
	}
	return failures
}

// isSensitiveSentinel reports whether terraform reported this value
// as opaque-sensitive ("(sensitive)" string). DESIGN.md §17.7 rule 7
// says we should reject these at parse time eventually; until then
// we surface a useful failure pointing the user at the proxy
// pattern.
func isSensitiveSentinel(v any) bool {
	s, ok := v.(string)
	return ok && s == "(sensitive)"
}

// dotWalk navigates a `terraform show -json` resource's `values` map
// using a dot-delimited attribute path (e.g. "comment" or
// "provider_config.workspace_id"). Returns (value, true) on success;
// (nil, false) when any segment is missing or the path tries to
// recurse into a non-map.
func dotWalk(values map[string]any, attr string) (any, bool) {
	parts := strings.Split(attr, ".")
	var current any = values
	for _, p := range parts {
		m, ok := current.(map[string]any)
		if !ok {
			return nil, false
		}
		next, ok := m[p]
		if !ok {
			return nil, false
		}
		current = next
	}
	return current, true
}

// equal compares an expected (YAML-decoded) value against an actual
// (JSON-decoded) value. JSON decodes all numbers as float64 while
// YAML's int decoder produces int, so we normalize integer expected
// values to float64 before reflect.DeepEqual.
func equal(expected, got any) bool {
	expected = normalizeNumeric(expected)
	got = normalizeNumeric(got)
	return reflect.DeepEqual(expected, got)
}

// normalizeNumeric coerces int / int64 / int32 / float32 to float64
// to match JSON's default decoded numeric type. Other types pass
// through unchanged. Recursive on slices and maps so nested expected
// values with int leaves still compare equal to JSON-decoded float64
// leaves.
func normalizeNumeric(v any) any {
	switch x := v.(type) {
	case int:
		return float64(x)
	case int32:
		return float64(x)
	case int64:
		return float64(x)
	case float32:
		return float64(x)
	case []any:
		out := make([]any, len(x))
		for i, item := range x {
			out[i] = normalizeNumeric(item)
		}
		return out
	case map[string]any:
		out := make(map[string]any, len(x))
		for k, item := range x {
			out[k] = normalizeNumeric(item)
		}
		return out
	default:
		return v
	}
}

// mapToSliceEnv converts a map[string]string env into the os/exec
// "KEY=VALUE" slice form. Same shape used by tfexec internally.
func mapToSliceEnv(env map[string]string) []string {
	if env == nil {
		return nil
	}
	out := make([]string, 0, len(env))
	for k, v := range env {
		out = append(out, k+"="+v)
	}
	return out
}
