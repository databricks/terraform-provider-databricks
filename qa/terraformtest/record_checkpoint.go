package terraformtest

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
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

// Publish an immutable checkpoint first, then atomically update its YAML
// reference. A failed/interrupted YAML write cannot corrupt an older snapshot
// that this or another test still references.
func saveRecordedCheckpoint(path string, original, updated, checkpoint []byte, requested string) error {
	name := checkpointFilename(requested, checkpoint)
	var document yaml.Node
	if err := yaml.Unmarshal(updated, &document); err != nil {
		return err
	}
	if err := setYAMLField(yamlField(document.Content[0], "initial_state"), "checkpoint", filepath.ToSlash(name)); err != nil {
		return err
	}
	final, err := yaml.Marshal(&document)
	if err != nil {
		return err
	}
	target := filepath.Join(filepath.Dir(path), name)
	created, err := publishCheckpoint(target, checkpoint)
	if err != nil {
		return err
	}
	if err := replaceRecording(path, original, final); err != nil {
		if created {
			_ = os.Remove(target)
		}
		return err
	}
	return nil
}

func checkpointFilename(requested string, checkpoint []byte) string {
	ext := filepath.Ext(requested)
	stem := strings.TrimSuffix(requested, ext)
	// Refresh a previously generated name instead of appending hashes forever.
	if dot := strings.LastIndexByte(stem, '.'); dot >= 0 && len(stem)-dot-1 == 64 {
		if _, err := hex.DecodeString(stem[dot+1:]); err == nil {
			stem = stem[:dot]
		}
	}
	return fmt.Sprintf("%s.%x%s", stem, sha256.Sum256(checkpoint), ext)
}

func publishCheckpoint(path string, checkpoint []byte) (bool, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return false, err
	}
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
	if os.IsExist(err) {
		existing, readErr := os.ReadFile(path)
		if readErr == nil && !bytes.Equal(existing, checkpoint) {
			readErr = fmt.Errorf("immutable checkpoint already exists with different contents")
		}
		return false, readErr
	}
	if err != nil {
		return false, err
	}
	_, writeErr := file.Write(checkpoint)
	if writeErr == nil {
		writeErr = file.Sync()
	}
	closeErr := file.Close()
	if writeErr != nil || closeErr != nil {
		_ = os.Remove(path)
		return false, fmt.Errorf("could not publish initial checkpoint")
	}
	return true, nil
}
