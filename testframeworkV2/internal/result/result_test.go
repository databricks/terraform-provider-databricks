package result

import (
	"strings"
	"testing"
	"time"
)

// TestRunResult_AllPassed_HappyPath confirms a run with all steps
// passing reports AllPassed=true.
func TestRunResult_AllPassed_HappyPath(t *testing.T) {
	r := RunResult{Steps: []StepResult{{Status: StatusPass}, {Status: StatusPass}}}
	if !r.AllPassed() {
		t.Errorf("expected AllPassed=true")
	}
}

// TestRunResult_AllPassed_OneFailFlipsResult covers the negative path:
// any non-pass step (fail, skipped) flips the result.
func TestRunResult_AllPassed_OneFailFlipsResult(t *testing.T) {
	for _, status := range []Status{StatusFail, StatusSkipped} {
		r := RunResult{Steps: []StepResult{{Status: StatusPass}, {Status: status}}}
		if r.AllPassed() {
			t.Errorf("status %q should NOT be all-passed", status)
		}
	}
}

// TestRunResult_AllPassed_SkippedRun documents that a skipped run
// trivially returns true (no steps were attempted).
func TestRunResult_AllPassed_SkippedRun(t *testing.T) {
	r := RunResult{Skipped: true, Reason: "requires.cloud=gcp but profile is aws"}
	if !r.AllPassed() {
		t.Errorf("skipped run should be AllPassed=true")
	}
}

func TestRunResult_FailedSteps(t *testing.T) {
	r := RunResult{
		Steps: []StepResult{
			{Index: 0, Status: StatusPass},
			{Index: 1, Status: StatusFail},
			{Index: 2, Status: StatusPass},
			{Index: 3, Status: StatusSkipped},
		},
	}
	got := r.FailedSteps()
	want := []int{1, 3}
	if len(got) != len(want) {
		t.Fatalf("FailedSteps len: got %d want %d", len(got), len(want))
	}
	for i, n := range want {
		if got[i] != n {
			t.Errorf("FailedSteps[%d]: got %d want %d", i, got[i], n)
		}
	}
}

func TestRunResult_String_Skipped(t *testing.T) {
	r := RunResult{Test: "t1", Skipped: true, Reason: "no GCP profile"}
	got := r.String()
	for _, want := range []string{"t1", "SKIPPED", "no GCP profile"} {
		if !strings.Contains(got, want) {
			t.Errorf("String() should contain %q, got %q", want, got)
		}
	}
}

func TestRunResult_String_PassFail(t *testing.T) {
	r := RunResult{
		Test:     "t1",
		Steps:    []StepResult{{Status: StatusPass}, {Status: StatusFail}, {Status: StatusPass}},
		Duration: 1500 * time.Millisecond,
	}
	got := r.String()
	for _, want := range []string{"t1", "FAIL", "2/3", "1.5s"} {
		if !strings.Contains(got, want) {
			t.Errorf("String() should contain %q, got %q", want, got)
		}
	}
}

func TestDurationString(t *testing.T) {
	for _, tc := range []struct {
		d    time.Duration
		want string
	}{
		{500 * time.Millisecond, "500ms"},
		{1500 * time.Millisecond, "1.5s"},
		{2 * time.Second, "2s"},
		{30 * time.Second, "30s"},
	} {
		if got := durationString(tc.d); got != tc.want {
			t.Errorf("durationString(%s) = %q, want %q", tc.d, got, tc.want)
		}
	}
}
