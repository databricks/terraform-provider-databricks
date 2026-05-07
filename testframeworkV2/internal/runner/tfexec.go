package runner

import (
	"context"
	"fmt"
	"io"

	"github.com/hashicorp/terraform-exec/tfexec"
)

// tfExec is the slice of *tfexec.Terraform's surface area the runner
// needs, behind an interface so unit tests can mock it. The interface
// is non-variadic because Go's type-assertion semantics don't allow a
// variadic method to satisfy a non-variadic one — the realTF wrapper
// below bridges that gap by binding the relevant per-command options
// once at construction.
type tfExec interface {
	Init(ctx context.Context) error
	Plan(ctx context.Context) (bool, error)
	Apply(ctx context.Context) error
	Destroy(ctx context.Context) error

	// SetEnv replaces the subprocess env entirely (it does NOT inherit
	// os.Environ()). The runner relies on this replacement semantics
	// for B5 leak protection (DESIGN.md §10/G6).
	SetEnv(env map[string]string) error

	// SetStdout / SetStderr direct the terraform subprocess's streams
	// to the per-step log files (DESIGN.md §3 — runtime tree).
	SetStdout(w io.Writer)
	SetStderr(w io.Writer)
}

// tfFactory creates a tfExec for the given workDir + binary path. The
// runner calls this once per step (a fresh tfexec.Terraform per step
// matches tfexec's documented usage; the cost is negligible because
// SetEnv / SetStdout / SetStderr are O(1)).
type tfFactory func(workDir, bin string) (tfExec, error)

// realTF wraps a *tfexec.Terraform and binds the framework-default
// per-command options (auto-approve on Apply / Destroy). Runtime
// callers don't need to know which options the framework prefers —
// they just call Apply(ctx) and the wrapper does the right thing.
type realTF struct {
	inner *tfexec.Terraform
}

// newRealTF is the default tfFactory. M5's CLI wires it up; M4 unit
// tests inject their own factory via WithTFFactory.
func newRealTF(workDir, bin string) (tfExec, error) {
	tf, err := tfexec.NewTerraform(workDir, bin)
	if err != nil {
		return nil, fmt.Errorf("runner: tfexec.NewTerraform(%q, %q): %w", workDir, bin, err)
	}
	return &realTF{inner: tf}, nil
}

func (r *realTF) Init(ctx context.Context) error {
	// The framework deliberately does NOT pass tfexec.Reconfigure(true)
	// here — §7.1.c removes .terraform/ entirely before each init,
	// which makes Reconfigure unnecessary.
	return r.inner.Init(ctx)
}

func (r *realTF) Plan(ctx context.Context) (bool, error) {
	// Plan returns (hasChanges, error). The runner discards hasChanges
	// — we only care about success/failure for assertion semantics
	// (DESIGN.md §7).
	return r.inner.Plan(ctx)
}

func (r *realTF) Apply(ctx context.Context) error {
	return r.inner.Apply(ctx)
}

func (r *realTF) Destroy(ctx context.Context) error {
	return r.inner.Destroy(ctx)
}

func (r *realTF) SetEnv(env map[string]string) error {
	return r.inner.SetEnv(env)
}

func (r *realTF) SetStdout(w io.Writer) {
	r.inner.SetStdout(w)
}

func (r *realTF) SetStderr(w io.Writer) {
	r.inner.SetStderr(w)
}
