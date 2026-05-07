package stateassert

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/config"
	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/result"
)

// goldenState is a hand-crafted `terraform show -json` payload
// covering the shapes the package's evaluator has to handle:
//   - top-level scalar (`comment`)
//   - top-level numeric (`lifetime_seconds`) — YAML int vs JSON float64
//   - top-level list of strings (`tags`)
//   - nested map (`provider_config.workspace_id`)
//   - sensitive sentinel (`token_value`)
//   - resource in a child module
//
// Lives as a Go literal rather than a fixture file so the test is
// hermetic — no external file IO.
var goldenState = tfState{
	Values: struct {
		RootModule tfModule `json:"root_module"`
	}{
		RootModule: tfModule{
			Resources: []tfResource{{
				Address: "databricks_token.pat",
				Values: map[string]any{
					"comment":          "tfv2-token-lifecycle-step-1",
					"lifetime_seconds": float64(3600),
					"tags":             []any{"alpha", "beta"},
					"token_value":      "(sensitive)",
					"provider_config": map[string]any{
						"workspace_id": "1234567890",
					},
				},
			}},
			ChildModules: []tfModule{{
				Resources: []tfResource{{
					Address: "module.nested.databricks_secret.s",
					Values:  map[string]any{"key": "ABC"},
				}},
			}},
		},
	},
}

// indexed returns the address->values map evaluate() walks. Tests
// use this directly to skip the terraform-show step.
func indexed() map[string]map[string]any {
	return indexResources(&goldenState)
}

func ptrTrue() *bool  { b := true; return &b }
func ptrFalse() *bool { b := false; return &b }

func TestIndexResources_FlattensChildModules(t *testing.T) {
	got := indexed()
	for _, want := range []string{"databricks_token.pat", "module.nested.databricks_secret.s"} {
		if _, ok := got[want]; !ok {
			t.Errorf("expected %q in index", want)
		}
	}
}

// TestEvaluate_PresentDefaultPasses confirms the common case: an
// assertion with a present resource and an attrs match returns
// nothing.
func TestEvaluate_PresentDefaultPasses(t *testing.T) {
	failures := evaluate(indexed(), config.Assertion{
		Resource: "databricks_token.pat",
		Attrs: map[string]any{
			"comment":          "tfv2-token-lifecycle-step-1",
			"lifetime_seconds": 3600, // YAML int → coerced to float64 for compare
		},
	})
	if len(failures) != 0 {
		t.Errorf("expected pass, got %v", failures)
	}
}

func TestEvaluate_PresentExplicitPasses(t *testing.T) {
	failures := evaluate(indexed(), config.Assertion{
		Resource: "databricks_token.pat",
		Present:  ptrTrue(),
		Attrs:    map[string]any{"comment": "tfv2-token-lifecycle-step-1"},
	})
	if len(failures) != 0 {
		t.Errorf("expected pass, got %v", failures)
	}
}

// TestEvaluate_PresentTrue_NotFound covers the "expected present,
// missing from state" failure shape.
func TestEvaluate_PresentTrue_NotFound(t *testing.T) {
	failures := evaluate(indexed(), config.Assertion{
		Resource: "databricks_nonexistent.x",
	})
	if len(failures) != 1 {
		t.Fatalf("expected 1 failure, got %v", failures)
	}
	if failures[0].Reason != "expected present, not found in state" {
		t.Errorf("Reason: %q", failures[0].Reason)
	}
	if failures[0].Field != "" {
		t.Errorf("Field should be empty for presence failure, got %q", failures[0].Field)
	}
}

// TestEvaluate_PresentFalse_NotFound is the absence-assertion happy
// path: "I expect no databricks_token after destroy".
func TestEvaluate_PresentFalse_NotFound(t *testing.T) {
	failures := evaluate(indexed(), config.Assertion{
		Resource: "databricks_nonexistent.x",
		Present:  ptrFalse(),
	})
	if len(failures) != 0 {
		t.Errorf("absence-assertion should pass when resource not in state, got %v", failures)
	}
}

func TestEvaluate_PresentFalse_FoundFails(t *testing.T) {
	failures := evaluate(indexed(), config.Assertion{
		Resource: "databricks_token.pat",
		Present:  ptrFalse(),
	})
	if len(failures) != 1 {
		t.Fatalf("expected 1 failure, got %v", failures)
	}
	if failures[0].Reason != "expected absent, found in state" {
		t.Errorf("Reason: %q", failures[0].Reason)
	}
}

func TestEvaluate_AttrMismatch(t *testing.T) {
	failures := evaluate(indexed(), config.Assertion{
		Resource: "databricks_token.pat",
		Attrs:    map[string]any{"comment": "wrong-value"},
	})
	if len(failures) != 1 {
		t.Fatalf("expected 1 failure, got %v", failures)
	}
	if failures[0].Field != "comment" {
		t.Errorf("Field: %q", failures[0].Field)
	}
	if failures[0].Reason != "value mismatch" {
		t.Errorf("Reason: %q", failures[0].Reason)
	}
	if failures[0].Actual != "tfv2-token-lifecycle-step-1" {
		t.Errorf("Actual: %v", failures[0].Actual)
	}
}

func TestEvaluate_AttrNotFound(t *testing.T) {
	failures := evaluate(indexed(), config.Assertion{
		Resource: "databricks_token.pat",
		Attrs:    map[string]any{"no_such_field": "x"},
	})
	if len(failures) != 1 {
		t.Fatalf("expected 1 failure, got %v", failures)
	}
	if failures[0].Reason != "attribute not found" {
		t.Errorf("Reason: %q", failures[0].Reason)
	}
}

func TestEvaluate_NestedAttrMatch(t *testing.T) {
	failures := evaluate(indexed(), config.Assertion{
		Resource: "databricks_token.pat",
		Attrs:    map[string]any{"provider_config.workspace_id": "1234567890"},
	})
	if len(failures) != 0 {
		t.Errorf("nested match should pass, got %v", failures)
	}
}

// TestEvaluate_CollectAllFailures pins the §17.5 invariant: never
// short-circuit, accumulate all per-attr failures.
func TestEvaluate_CollectAllFailures(t *testing.T) {
	failures := evaluate(indexed(), config.Assertion{
		Resource: "databricks_token.pat",
		Attrs: map[string]any{
			"comment":          "wrong",
			"lifetime_seconds": 9999, // wrong
			"missing":          "x",
		},
	})
	if len(failures) != 3 {
		t.Errorf("expected 3 failures (collect-all-failures), got %d: %v", len(failures), failures)
	}
}

// TestEvaluate_SensitiveAttrFails confirms an assertion against a
// (sensitive)-marked attribute surfaces a clear failure pointing at
// the proxy-attribute pattern.
func TestEvaluate_SensitiveAttrFails(t *testing.T) {
	failures := evaluate(indexed(), config.Assertion{
		Resource: "databricks_token.pat",
		Attrs:    map[string]any{"token_value": "anything"},
	})
	if len(failures) != 1 {
		t.Fatalf("expected 1 failure, got %v", failures)
	}
	if !contains(failures[0].Reason, "sensitive") {
		t.Errorf("Reason should mention sensitive: %q", failures[0].Reason)
	}
}

func TestEvaluate_ChildModuleResource(t *testing.T) {
	failures := evaluate(indexed(), config.Assertion{
		Resource: "module.nested.databricks_secret.s",
		Attrs:    map[string]any{"key": "ABC"},
	})
	if len(failures) != 0 {
		t.Errorf("child-module resource match should pass, got %v", failures)
	}
}

// TestEvaluate_ListMatch covers the YAML []any expected vs JSON
// []any actual case (e.g. tags lists).
func TestEvaluate_ListMatch(t *testing.T) {
	failures := evaluate(indexed(), config.Assertion{
		Resource: "databricks_token.pat",
		Attrs:    map[string]any{"tags": []any{"alpha", "beta"}},
	})
	if len(failures) != 0 {
		t.Errorf("list match should pass, got %v", failures)
	}
}

// TestRun_NoAssertions confirms the early-return: an empty list
// means stateassert never spawns terraform.
func TestRun_NoAssertions(t *testing.T) {
	failures, err := Run(t.Context(), "/nonexistent", "/nonexistent", nil, nil)
	if err != nil {
		t.Errorf("Run with no assertions should not call terraform: %v", err)
	}
	if failures != nil {
		t.Errorf("expected nil failures, got %v", failures)
	}
}

func TestNormalizeNumeric_Recursive(t *testing.T) {
	in := map[string]any{
		"a": int(1),
		"b": []any{int(2), int(3)},
		"c": map[string]any{"d": int64(4)},
	}
	got := normalizeNumeric(in)
	want := map[string]any{
		"a": float64(1),
		"b": []any{float64(2), float64(3)},
		"c": map[string]any{"d": float64(4)},
	}
	if !equal(want, got) {
		t.Errorf("normalize did not recurse:\n got=%#v\nwant=%#v", got, want)
	}
}

// TestAssertionFailureString covers result.AssertionFailure.String —
// asserts it's stable across with-Field and without-Field cases.
func TestAssertionFailureString(t *testing.T) {
	withField := result.AssertionFailure{Address: "x.y", Field: "z", Reason: "value mismatch", Expected: "a", Actual: "b"}
	want := "x.y.z: value mismatch (expected=a, actual=b)"
	if got := withField.String(); got != want {
		t.Errorf("with field:\n got %q\nwant %q", got, want)
	}
	noField := result.AssertionFailure{Address: "x.y", Reason: "expected present, not found in state"}
	want = "x.y: expected present, not found in state"
	if got := noField.String(); got != want {
		t.Errorf("no field:\n got %q\nwant %q", got, want)
	}
}

// TestMapToSliceEnv covers the env conversion. nil → nil; populated
// → KEY=VALUE entries.
func TestMapToSliceEnv(t *testing.T) {
	if got := mapToSliceEnv(nil); got != nil {
		t.Errorf("nil map should yield nil slice, got %v", got)
	}
	got := mapToSliceEnv(map[string]string{"FOO": "bar", "BAZ": "qux"})
	if len(got) != 2 {
		t.Errorf("expected 2 entries, got %v", got)
	}
}

// contains is a substring helper — avoiding a strings import for one
// call.
func contains(s, sub string) bool {
	return len(s) >= len(sub) && (func() bool {
		for i := 0; i+len(sub) <= len(s); i++ {
			if s[i:i+len(sub)] == sub {
				return true
			}
		}
		return false
	}())
}
