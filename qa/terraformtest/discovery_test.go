package terraformtest

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiscoverFilesSortedAndNonRecursive(t *testing.T) {
	directory := t.TempDir()
	for _, name := range []string{"Zulu.yaml", "Alpha.yml", "README.md", "nested/Ignored.yaml"} {
		writeDiscoveryFile(t, directory, name, "")
	}
	// A directory whose name looks like a test must not be registered.
	require.NoError(t, os.Mkdir(filepath.Join(directory, "Directory.yaml"), 0700))
	files, err := discoverFiles(directory)
	require.NoError(t, err)
	require.Equal(t, []testFile{
		{name: "Alpha", path: filepath.Join(directory, "Alpha.yml")},
		{name: "Zulu", path: filepath.Join(directory, "Zulu.yaml")},
	}, files)
}

func TestDiscoverFilesRejectsInvalidSuites(t *testing.T) {
	for _, tc := range []struct {
		name    string
		files   []string
		message string
	}{
		{"empty", nil, "no Terraform tests found"},
		{"unrelated files", []string{"README.md", "nested/Test.yaml"}, "no Terraform tests found"},
		{"duplicate stems", []string{"Test.yaml", "Test.yml"}, `duplicate Terraform test name "Test"`},
		{"empty stem", []string{".yaml"}, "nonempty stem"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			directory := t.TempDir()
			for _, name := range tc.files {
				writeDiscoveryFile(t, directory, name, "")
			}
			_, err := discoverFiles(directory)
			require.ErrorContains(t, err, tc.message)
			require.ErrorContains(t, err, directory)
		})
	}
	_, err := discoverFiles(filepath.Join(t.TempDir(), "missing"))
	require.ErrorContains(t, err, "discover Terraform tests")
}

func TestRunRegistersFilesAndIsolatesFailures(t *testing.T) {
	directory := t.TempDir()
	writeDiscoveryFile(t, directory, "First.yaml", "version: 2\n")
	writeDiscoveryFile(t, directory, "Second.yml", "version: 3\n")
	output, err := discoverySubprocess(t, directory, "^TestDirectoryRunnerHelper$")
	require.Error(t, err)
	require.Contains(t, output, "--- FAIL: TestDirectoryRunnerHelper/First")
	require.Contains(t, output, "--- FAIL: TestDirectoryRunnerHelper/Second")
	require.Contains(t, output, "First.yaml: unsupported version 2")
	require.Contains(t, output, "Second.yml: unsupported version 3")
	require.Less(t, strings.Index(output, "=== RUN   TestDirectoryRunnerHelper/First"),
		strings.Index(output, "=== RUN   TestDirectoryRunnerHelper/Second"))
}

func TestRunSupportsSubtestFiltering(t *testing.T) {
	directory := t.TempDir()
	writeDiscoveryFile(t, directory, "First.yaml", "version: 2\n")
	writeDiscoveryFile(t, directory, "Second.yml", "version: 3\n")
	output, err := discoverySubprocess(t, directory, "^TestDirectoryRunnerHelper$/^Second$")
	require.Error(t, err)
	require.Contains(t, output, "--- FAIL: TestDirectoryRunnerHelper/Second")
	require.Contains(t, output, "Second.yml: unsupported version 3")
	require.NotContains(t, output, "TestDirectoryRunnerHelper/First")
	require.NotContains(t, output, "unsupported version 2")
}

func TestRunRejectsEmptyDirectory(t *testing.T) {
	output, err := discoverySubprocess(t, t.TempDir(), "^TestDirectoryRunnerHelper$")
	require.Error(t, err)
	require.Contains(t, output, "no Terraform tests found")
}

func TestDirectoryRunnerHelper(t *testing.T) {
	directory := os.Getenv("TERRAFORMTEST_DISCOVERY_DIRECTORY")
	if directory == "" {
		t.Skip("subprocess helper for testing fatal errors and subtest registration")
	}
	Run(t, directory)
}

// Invalid definitions fail before CLI discovery, so registration and filtering
// tests do not depend on an installed Terraform executable.
func discoverySubprocess(t *testing.T, directory, filter string) (string, error) {
	t.Helper()
	executable, err := os.Executable()
	require.NoError(t, err)
	cmd := exec.Command(executable, "-test.v", "-test.run="+filter)
	cmd.Env = append(os.Environ(), "TERRAFORMTEST_DISCOVERY_DIRECTORY="+directory)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func writeDiscoveryFile(t *testing.T, directory, name, content string) {
	t.Helper()
	path := filepath.Join(directory, name)
	require.NoError(t, os.MkdirAll(filepath.Dir(path), 0700))
	require.NoError(t, os.WriteFile(path, []byte(content), 0600))
}
