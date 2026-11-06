package terraformtest

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"gopkg.in/yaml.v3"
)

// Terraform generates the checkpoint from sanitized setup responses using the
// real provider. Copy its bytes at the setup boundary; never construct state
// JSON, rewrite private provider data, or sanitize a live checkpoint in place.
func replayCheckpoint(t *testing.T, d *testDefinition) []byte {
	t.Helper()
	configureTerraform(t)
	handler, host := replayBoundary(t, d)
	native := terraformCase(t, d, handler, fixtureConfig(host), nil)
	parent := t.TempDir()
	native.WorkingDir = parent
	var checkpoint []byte
	verified := native.Steps[0].PostApplyFunc
	native.Steps[0].PostApplyFunc = func() {
		verified()
		files, err := filepath.Glob(filepath.Join(parent, "work*", "terraform.tfstate"))
		if err != nil || len(files) != 1 {
			t.Fatal("expected one Terraform-generated setup checkpoint")
		}
		checkpoint, err = os.ReadFile(files[0])
		if err != nil {
			t.Fatal(err)
		}
	}
	resource.Test(t, native)
	if len(checkpoint) == 0 {
		t.Fatal("initial_state_template did not produce a Terraform checkpoint")
	}
	return checkpoint
}

// Write the test's checkpoint before updating the YAML. If the YAML write fails,
// rerun recording to refresh both files; no checkpoint rollback is needed.
func saveRecordedCheckpoint(path string, original, updated, checkpoint []byte) error {
	name := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	relative := filepath.Join("checkpoints", name+".tfstate")
	var document yaml.Node
	if err := yaml.Unmarshal(updated, &document); err != nil {
		return err
	}
	if err := setYAMLField(yamlField(document.Content[0], "initial_state"), "checkpoint", filepath.ToSlash(relative)); err != nil {
		return err
	}
	final, err := yaml.Marshal(&document)
	if err != nil {
		return err
	}
	target := filepath.Join(filepath.Dir(path), relative)
	if err := os.MkdirAll(filepath.Dir(target), 0700); err != nil {
		return err
	}
	if err := os.WriteFile(target, checkpoint, 0600); err != nil {
		return err
	}
	return replaceRecording(path, original, final)
}
