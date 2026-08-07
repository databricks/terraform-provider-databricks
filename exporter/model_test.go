package exporter

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResourceApproaximationGet(t *testing.T) {
	_, found := (&resourceApproximation{}).Get("test")
	assert.False(t, found)

	v, found := (&resourceApproximation{
		Instances: []instanceApproximation{
			{Attributes: map[string]any{"test": "42"}},
		},
	}).Get("test")
	require.True(t, found)
	assert.Equal(t, "42", v.(string))
}

func TestExtraDataGet(t *testing.T) {
	r := &resource{}
	_, found := r.GetExtraData("test")
	assert.False(t, found)

	r.AddExtraData("test", "42")
	v, found := r.GetExtraData("test")
	require.True(t, found)
	assert.Equal(t, "42", v.(string))
}

// TestImportCommandShellEscaping is a regression test for SEC-21613
// (H1-3750124): attacker-controlled workspace object paths flow into r.ID and
// are interpolated into the generated import.sh. The ID must be POSIX
// single-quote-escaped so that shell metacharacters are treated as literals.
func TestImportCommandShellEscaping(t *testing.T) {
	ic := &importContext{}
	cases := []struct {
		name string
		id   string
		want string
	}{
		{
			name: "benign path",
			id:   "/Users/me/notebook",
			want: `terraform import databricks_notebook.foo '/Users/me/notebook'`,
		},
		{
			name: "command substitution",
			id:   "/Shared/$(touch pwned).py",
			want: `terraform import databricks_notebook.foo '/Shared/$(touch pwned).py'`,
		},
		{
			name: "backtick substitution",
			id:   "/a/`touch pwned`.py",
			want: "terraform import databricks_notebook.foo '/a/`touch pwned`.py'",
		},
		{
			name: "variable expansion",
			id:   "/a/${HOME}.py",
			want: `terraform import databricks_notebook.foo '/a/${HOME}.py'`,
		},
		{
			name: "double quotes",
			id:   `/a/"b".py`,
			want: `terraform import databricks_notebook.foo '/a/"b".py'`,
		},
		{
			name: "embedded single quotes",
			id:   `/a/'b'.py`,
			want: `terraform import databricks_notebook.foo '/a/'\''b'\''.py'`,
		},
		{
			name: "backslash",
			id:   `/a/\b.py`,
			want: `terraform import databricks_notebook.foo '/a/\b.py'`,
		},
		{
			name: "newline",
			id:   "/a/\nrm -rf x\n.py",
			want: "terraform import databricks_notebook.foo '/a/\nrm -rf x\n.py'",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r := &resource{Resource: "databricks_notebook", Name: "foo", ID: c.id}
			assert.Equal(t, c.want, r.ImportCommand(ic))
		})
	}

	// The module prefix must still be emitted for the resource address.
	r := &resource{Resource: "databricks_notebook", Name: "foo", ID: "/a/b"}
	icMod := &importContext{Module: "mymod"}
	assert.Equal(t, `terraform import mymod.databricks_notebook.foo '/a/b'`, r.ImportCommand(icMod))
}

// TestImportCommandNoShellInjectionWhenExecuted executes a generated import.sh
// containing malicious workspace paths and asserts that no marker file is
// created, i.e. the embedded commands do not run. Regression test for SEC-21613.
func TestImportCommandNoShellInjectionWhenExecuted(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("import.sh is a POSIX shell script")
	}
	dir := t.TempDir()

	// Stub `terraform` on PATH so the script's `set -e` doesn't abort and no
	// real CLI is required; the stub is a no-op.
	stub := filepath.Join(dir, "terraform")
	require.NoError(t, os.WriteFile(stub, []byte("#!/bin/sh\nexit 0\n"), 0o755))
	t.Setenv("PATH", dir+string(os.PathListSeparator)+os.Getenv("PATH"))

	dollarMarker := filepath.Join(dir, "pwned_dollar")
	backtickMarker := filepath.Join(dir, "pwned_backtick")

	ic := &importContext{}
	payloads := []string{
		"/Shared/$(touch " + dollarMarker + ").py",
		"/Shared/`touch " + backtickMarker + "`.py",
	}
	var lines []string
	for _, p := range payloads {
		r := &resource{Resource: "databricks_notebook", Name: "foo", ID: p}
		lines = append(lines, r.ImportCommand(ic))
	}
	// Mirror the real import.sh header (context.go).
	script := "#!/bin/sh\n\nset -e\n\n" + strings.Join(lines, "\n") + "\n"
	shPath := filepath.Join(dir, "import.sh")
	require.NoError(t, os.WriteFile(shPath, []byte(script), 0o755))

	out, err := exec.Command("/bin/sh", shPath).CombinedOutput()
	require.NoError(t, err, "import.sh failed to run: %s", out)

	markers, err := filepath.Glob(filepath.Join(dir, "pwned_*"))
	require.NoError(t, err)
	assert.Empty(t, markers, "command injection executed; marker files were created: %v", markers)
}
