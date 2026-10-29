package terraformtest

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// validateFiles protects Terraform configuration, state and framework files
// from being overwritten by resource inputs. Sources are relative to the YAML;
// destinations are filenames in Terraform's working directory.
func validateFiles(files map[string]string) error {
	for name, source := range files {
		if name == "" || name == "." || name == ".." || filepath.Base(name) != name || strings.ContainsAny(name, `/\\`) {
			return fmt.Errorf("files destination %q must be a filename without directories", name)
		}
		if strings.HasPrefix(name, ".terraform") || strings.HasPrefix(name, "terraform.") ||
			strings.HasPrefix(name, "terraform_plugin_test") || strings.HasSuffix(name, ".tf") ||
			strings.HasSuffix(name, ".tf.json") || strings.HasSuffix(name, ".tfvars") || strings.HasSuffix(name, ".tfvars.json") {
			return fmt.Errorf("files destination %q is reserved for Terraform", name)
		}
		if strings.TrimSpace(source) == "" || filepath.IsAbs(source) {
			return fmt.Errorf("files[%q] must be a nonempty path relative to the YAML file", name)
		}
	}
	return nil
}

// loadFiles reads every step's inputs before Terraform starts, so missing files
// in a later step cannot interrupt a partially applied lifecycle.
func (d *testDefinition) loadFiles(directory string) error {
	for i := range d.Steps {
		s := &d.Steps[i]
		s.fileContents = make(map[string][]byte, len(s.Files))
		for name, source := range s.Files {
			body, err := os.ReadFile(filepath.Join(directory, source))
			if err != nil {
				return fmt.Errorf("%s files[%q]: %w", stepLabel(i, s.Name), name, err)
			}
			s.fileContents[name] = body
		}
	}
	return nil
}

// configureStepFiles stages inputs alongside HCL so HashiCorp's ConfigDirectory
// API copies them into Terraform's isolated working directory. This lets resource
// source paths refer to local files without modifying the checked-in fixtures.
// Recopying cumulative inputs also restores JSON files removed by config cleanup.
// A changed source overwrites the same destination; omitted inputs persist
// through later steps and destroy, matching ordinary local-file workflows.
func configureStepFiles(t *testing.T, definition *testDefinition, testCase *resource.TestCase) {
	t.Helper()
	files := map[string][]byte{}
	for i := range testCase.Steps {
		if i < len(definition.Steps) {
			for name, body := range definition.Steps[i].fileContents {
				files[name] = body
			}
		}
		if len(files) == 0 {
			continue
		}
		s := &testCase.Steps[i]
		directory := stepFilesDirectory(t, s)
		for name, body := range files {
			if err := os.WriteFile(filepath.Join(directory, name), body, 0600); err != nil {
				t.Fatal(err)
			}
		}
	}
	if len(files) > 0 {
		useStepProviderFactories(testCase)
	}
}

// stepFilesDirectory prepares a staging directory, not Terraform's execution
// directory. Each step gets its own HCL and input snapshot, which HashiCorp later
// copies into the lifecycle's shared working directory before running Terraform.
func stepFilesDirectory(t *testing.T, s *resource.TestStep) string {
	t.Helper()
	if s.ConfigDirectory != nil {
		// The initial-checkpoint step already owns a staging directory.
		return s.ConfigDirectory.Exec(config.TestStepConfigRequest{})
	}
	directory := t.TempDir()
	hcl, err := checkpointConfig(s.Config)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(directory, "main.tf"), []byte(hcl), 0600); err != nil {
		t.Fatal(err)
	}
	s.Config, s.ConfigDirectory = "", config.StaticDirectory(directory)
	return directory
}

// ConfigDirectory may contain provider blocks; HashiCorp requires step-level
// provider factories in that case. Preserve factories already moved by seeding.
func useStepProviderFactories(testCase *resource.TestCase) {
	if testCase.ProtoV6ProviderFactories == nil {
		return
	}
	for i := range testCase.Steps {
		testCase.Steps[i].ProtoV6ProviderFactories = testCase.ProtoV6ProviderFactories
	}
	testCase.ProtoV6ProviderFactories = nil
}
