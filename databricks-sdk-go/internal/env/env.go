package env

import (
	"os"
	"strings"
	"testing"
)

// CleanupEnvironment sets up a pristine environment containing only $PATH and $HOME.
// The original environment is restored upon test completion.
// Note: use of this function is incompatible with parallel execution.
func CleanupEnvironment(t *testing.T) {
	// Restore environment when test finishes.
	environ := os.Environ()
	t.Cleanup(func() {
		// Restore original environment.
		for _, kv := range environ {
			kvs := strings.SplitN(kv, "=", 2)
			os.Setenv(kvs[0], kvs[1])
		}
	})

	path := os.Getenv("PATH")
	pwd := os.Getenv("PWD")
	os.Clearenv()

	// We use t.Setenv instead of os.Setenv because the former actively
	// prevents a test being run with t.Parallel. Modifying the environment
	// within a test is not compatible with running tests in parallel
	// because of isolation; the environment is scoped to the process.
	t.Setenv("PATH", path)
	t.Setenv("HOME", pwd)
}
