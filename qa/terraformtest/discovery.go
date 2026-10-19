package terraformtest

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Run discovers .yaml and .yml files directly in directory, relative to the
// test package, and runs each as a subtest named after its filename stem.
// Each lifecycle has independent Terraform state and HTTP fixtures. Subtests
// run sequentially in filename order and support normal go test -run filtering.
// Terraform must be installed; neither the wrapper nor its ancestors may use
// t.Parallel because CLI discovery uses process environment variables.
func Run(t *testing.T, directory string) {
	t.Helper()
	files, err := discoverFiles(directory)
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range files {
		t.Run(file.name, func(t *testing.T) {
			definition, err := loadDefinition(file.path)
			if err != nil {
				t.Fatalf("%s: %v", file.path, err)
			}
			run(t, definition)
		})
	}
}

type testFile struct {
	name string
	path string
}

// discoverFiles rejects empty suites and ambiguous names so a misplaced file
// or duplicate extension cannot silently reduce coverage or rename a subtest.
func discoverFiles(directory string) ([]testFile, error) {
	entries, err := os.ReadDir(directory)
	if err != nil {
		return nil, fmt.Errorf("discover Terraform tests in %s: %w", directory, err)
	}
	var files []testFile
	names := make(map[string]string)
	// os.ReadDir sorts by filename, keeping registration stable across runs.
	for _, entry := range entries {
		ext := filepath.Ext(entry.Name())
		if entry.IsDir() || (ext != ".yaml" && ext != ".yml") {
			continue
		}
		name := strings.TrimSuffix(entry.Name(), ext)
		path := filepath.Join(directory, entry.Name())
		if strings.TrimSpace(name) == "" {
			return nil, fmt.Errorf("%s: Terraform test filename must have a nonempty stem", path)
		}
		if previous, exists := names[name]; exists {
			return nil, fmt.Errorf("duplicate Terraform test name %q: %s and %s", name, previous, path)
		}
		names[name] = path
		files = append(files, testFile{name: name, path: path})
	}
	if len(files) == 0 {
		return nil, fmt.Errorf("no Terraform tests found in %s; expected .yaml or .yml files", directory)
	}
	return files, nil
}
