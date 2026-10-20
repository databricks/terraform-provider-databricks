package terraformtest

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// seedInitialCheckpoint uses HashiCorp's directory-copy API to restore the
// original state before the first step. Only that step gets the checkpoint;
// later steps and destroy must see the state produced by preceding operations.
func seedInitialCheckpoint(t *testing.T, definition *testDefinition, testCase *resource.TestCase) {
	t.Helper()
	for i, s := range definition.Steps {
		if _, err := checkpointConfig(s.Config); err != nil {
			t.Fatalf("%s: %v", stepLabel(i, s.Name), err)
		}
	}
	directory := t.TempDir()
	if err := os.WriteFile(filepath.Join(directory, "terraform.tfstate"), definition.checkpoint, 0600); err != nil {
		t.Fatal(err)
	}
	first := &testCase.Steps[0]
	hcl, err := checkpointConfig(first.Config)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(directory, "main.tf"), []byte(hcl), 0600); err != nil {
		t.Fatal(err)
	}
	first.Config = ""
	first.ConfigDirectory = config.StaticDirectory(directory)
	// ConfigDirectory can contain provider blocks; HashiCorp requires provider
	// factories at step level rather than case level when that is the case.
	for i := range testCase.Steps {
		testCase.Steps[i].ProtoV6ProviderFactories = testCase.ProtoV6ProviderFactories
	}
	testCase.ProtoV6ProviderFactories = nil
}

// ConfigDirectory expects complete provider requirements, unlike Config where
// HashiCorp supplies them. Keep the same default without duplicating an author's
// explicit terraform block or modifying any checkpoint contents.
func checkpointConfig(source string) (string, error) {
	file, diagnostics := hclwrite.ParseConfig([]byte(source), "config.tf", hcl.InitialPos)
	if diagnostics.HasErrors() {
		return "", fmt.Errorf("initial_state step configuration: %s", diagnostics.Error())
	}
	hasTerraformBlock := false
	for _, block := range file.Body().Blocks() {
		if block.Type() == "terraform" {
			hasTerraformBlock = true
			for _, setting := range block.Body().Blocks() {
				if setting.Type() == "backend" || setting.Type() == "cloud" {
					return "", fmt.Errorf("initial_state requires the default local backend; backend and cloud blocks are not supported")
				}
			}
		}
	}
	if hasTerraformBlock {
		return source, nil
	}
	return `terraform {
  required_providers {
    databricks = { source = "databricks/databricks" }
  }
}
` + source, nil
}
