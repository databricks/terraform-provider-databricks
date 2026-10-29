package terraformtest

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"gopkg.in/yaml.v3"

	"github.com/databricks/terraform-provider-databricks/qa/httpfixture"
)

func validateRecordingSetup(d *testDefinition) error {
	hasTemplate := d.Recording != nil && strings.TrimSpace(d.Recording.InitialStateTemplate) != ""
	if d.InitialState != nil && !hasTemplate {
		return fmt.Errorf("recording initial_state requires recording.initial_state_template")
	}
	if d.InitialState == nil && hasTemplate {
		return fmt.Errorf("recording.initial_state_template requires initial_state.checkpoint")
	}
	if d.Recording != nil {
		for _, template := range []string{d.Recording.EnvironmentTemplate, d.Recording.InitialStateTemplate} {
			if _, err := checkpointConfig(template); err != nil {
				return fmt.Errorf("recording setup: %w", err)
			}
		}
	}
	return nil
}

func recordWithSetup(t *testing.T, path string, original []byte, d, live *testDefinition, profile *recordingProfile, n *recordingNormalizer) {
	t.Helper()
	var captured []httpfixture.CapturedStep
	// Background resources own separate state, so the tested journey cannot
	// destroy them. Its cleanup finishes before environment teardown begins.
	if !recordingEnvironment(t, live, profile, func(t *testing.T) {
		capture := httpfixture.NewCaptureTransport(profile.transport())
		if !t.Run("record", func(t *testing.T) {
			runTerraform(t, withInitialTemplate(live, nil), recordingHTTP{capture}, profile.customize(capture), nil)
		}) {
			return
		}
		var err error
		captured, err = capture.Snapshot()
		if err != nil {
			t.Fatal(err)
		}
	}) || captured == nil {
		return
	}
	initialCalls, journey := splitSetupCapture(d, captured)
	updated, err := refreshRecording(original, d, journey, n)
	if err != nil {
		t.Fatal(err)
	}
	if d.InitialState == nil {
		if verifyRecordedFile(t, path, updated) {
			if err := replaceRecording(path, original, updated); err != nil {
				t.Fatal(err)
			}
		}
		return
	}
	finishInitialRecording(t, path, original, updated, initialCalls, n)
}

// Both record-time initial setup and checkpoint generation use ordinary steps
// in the journey's state. This is not a state-file edit or a separate state swap.
func withInitialTemplate(d *testDefinition, calls []httpfixture.Call) *testDefinition {
	copy := *d
	if d.InitialState != nil {
		first := step{Name: "recording_initial_state", Config: d.Recording.InitialStateTemplate, Calls: calls,
			Files: d.Steps[0].Files, fileContents: d.Steps[0].fileContents}
		copy.Steps = append([]step{first}, d.Steps...)
		copy.InitialState = nil
	}
	return &copy
}

func splitSetupCapture(d *testDefinition, captured []httpfixture.CapturedStep) ([]httpfixture.Call, []httpfixture.CapturedStep) {
	var initial []httpfixture.Call
	if d.InitialState != nil && len(captured) != 0 {
		initial, captured = captured[0].Calls, captured[1:]
	}
	// The recording-only setup precedes user steps. Saved fixture labels keep
	// user step indices, including the automatically numbered destroy.
	for i := range captured {
		name := "destroy"
		if i < len(d.Steps) {
			name = d.Steps[i].Name
		}
		captured[i].Label = stepLabel(i, name)
	}
	return initial, captured
}

func recordingEnvironment(t *testing.T, d *testDefinition, profile *recordingProfile, journey func(*testing.T)) bool {
	t.Helper()
	if d.Recording == nil || strings.TrimSpace(d.Recording.EnvironmentTemplate) == "" {
		return t.Run("journey", journey)
	}
	return t.Run("environment", func(t *testing.T) {
		configureTerraform(t)
		environment := &testDefinition{Steps: []step{{Name: "environment", Config: d.Recording.EnvironmentTemplate,
			Files: d.Steps[0].Files, fileContents: d.Steps[0].fileContents}}}
		capture := httpfixture.NewCaptureTransport(profile.transport())
		native := terraformCase(t, environment, recordingHTTP{capture}, profile.customize(capture), nil)
		verified := native.Steps[0].PostApplyFunc
		native.Steps[0].PostApplyFunc = func() {
			verified()
			t.Run("journey", journey)
		}
		resource.Test(t, native)
	})
}

func finishInitialRecording(t *testing.T, path string, original, updated []byte, initialCalls []httpfixture.Call, n *recordingNormalizer) {
	t.Helper()
	d, err := decodeDefinition(bytes.NewReader(updated))
	if err != nil {
		t.Fatal(err)
	}
	if err := d.loadFiles(filepath.Dir(path)); err != nil {
		t.Fatal(err)
	}
	initialCalls, err = normalizedSetupCalls(initialCalls, n)
	if err != nil {
		t.Fatal(err)
	}
	var checkpoint []byte
	if !t.Run("generate_checkpoint", func(t *testing.T) {
		checkpoint = replayCheckpoint(t, withInitialTemplate(d, initialCalls))
	}) {
		return
	}
	d.checkpoint = checkpoint
	if !t.Run("replay", func(t *testing.T) { run(t, d) }) {
		return
	}
	if err := saveRecordedCheckpoint(path, original, updated, checkpoint, d.InitialState.Checkpoint); err != nil {
		t.Fatal(err)
	}
}

func normalizedSetupCalls(initialCalls []httpfixture.Call, n *recordingNormalizer) ([]httpfixture.Call, error) {
	encoded, err := yaml.Marshal(initialCalls)
	if err != nil {
		return nil, err
	}
	encoded, err = n.transform(encoded, false, true)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(encoded, &initialCalls); err != nil {
		return nil, err
	}
	// Initial setup and the journey share the same reusable host metadata.
	var ignored map[string]any
	initialCalls, err = separateMetadata(initialCalls, &ignored)
	if err != nil {
		return nil, err
	}
	if err := httpfixture.PreserveMatchers(initialCalls, nil); err != nil {
		return nil, err
	}
	return initialCalls, nil
}
