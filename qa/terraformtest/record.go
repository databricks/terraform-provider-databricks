package terraformtest

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa/httpfixture"
)

const recorderProfileEnv = "TEST_RECORDER_PROFILE_NAME"

// Recording is deliberately selected before loading credentials. An inherited
// environment variable must not turn a suite or a CI job into live API traffic.
func recordingSelection(parent string, files []testFile, profile, filter string) error {
	if profile == "" {
		return nil
	}
	for _, name := range []string{"CI", "GITHUB_ACTIONS", "TF_BUILD", "BUILDKITE", "JENKINS_URL"} {
		if value := os.Getenv(name); value != "" && value != "false" && value != "0" {
			return fmt.Errorf("recording is forbidden in CI (%s)", name)
		}
	}
	for _, name := range []string{"TF_LOG", "TF_LOG_CORE", "TF_LOG_PROVIDER", "TF_ACC_LOG"} {
		if os.Getenv(name) != "" {
			return fmt.Errorf("disable %s before recording; live payloads must not be written to debug logs", name)
		}
	}
	for _, file := range files {
		parts := strings.Split(parent+"/"+file.name, "/")
		for i := range parts {
			parts[i] = "^" + regexp.QuoteMeta(parts[i]) + "$"
		}
		if filter == strings.Join(parts, "/") {
			return nil
		}
	}
	return fmt.Errorf("recording requires -run with one exact anchored test path, e.g. '^%s$/^%s$'", parent, files[0].name)
}

type recordingHTTP struct{ capture *httpfixture.CaptureTransport }

func (r recordingHTTP) StartStep(label string, _ []httpfixture.Call) error {
	return r.capture.StartStep(label)
}

func (r recordingHTTP) Verify() error { return r.capture.Verify() }

func recordFile(t *testing.T, path, profileName string) {
	t.Helper()
	original, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	definition, err := decodeDefinition(bytes.NewReader(original))
	if err != nil {
		t.Fatal(err)
	}
	if err := validateRecordingSetup(definition); err != nil {
		t.Fatal(err)
	}
	profile, err := loadRecordingProfile(profileName)
	if err != nil {
		t.Fatal(err)
	}
	recordLoadedFile(t, path, original, definition, profile)
}

func recordLoadedFile(t *testing.T, path string, original []byte, definition *testDefinition, profile *recordingProfile) {
	t.Helper()
	normalizer, err := newRecordingNormalizer(definition, profile.config)
	if err != nil {
		t.Fatal(err)
	}
	live, err := recordingDefinition(path, original, normalizer)
	if err != nil {
		t.Fatal(err)
	}
	if definition.InitialState != nil || definition.Recording != nil {
		recordWithSetup(t, path, original, definition, live, profile, normalizer)
		return
	}
	updated := captureRecording(t, original, definition, live, profile, normalizer)
	if updated == nil || !verifyRecordedFile(t, path, updated) {
		return
	}
	if err := replaceRecording(path, original, updated); err != nil {
		t.Fatal(err)
	}
}

func captureRecording(t *testing.T, original []byte, definition, live *testDefinition, profile *recordingProfile, normalizer *recordingNormalizer) []byte {
	t.Helper()
	capture := httpfixture.NewCaptureTransport(profile.transport())
	// resource.Test uses Fatal/Goexit on unexpected errors. A child test lets
	// cleanup finish and gives us a success result before touching the YAML.
	if !t.Run("record", func(t *testing.T) {
		runTerraform(t, live, recordingHTTP{capture}, profile.customize(capture), nil)
	}) {
		return nil
	}
	steps, err := capture.Snapshot()
	if err != nil {
		t.Fatal(err)
	}
	updated, err := refreshRecording(original, definition, steps, normalizer)
	if err != nil {
		t.Fatal(err)
	}
	return updated
}

func verifyRecordedFile(t *testing.T, path string, updated []byte) bool {
	t.Helper()
	return t.Run("replay", func(t *testing.T) {
		replayed, err := decodeDefinition(bytes.NewReader(updated))
		if err != nil {
			t.Fatal(err)
		}
		if err := replayed.loadFiles(filepath.Dir(path)); err != nil {
			t.Fatal(err)
		}
		run(t, replayed)
	})
}

func testRunFilter() string {
	if value := flag.Lookup("test.run"); value != nil {
		return value.Value.String()
	}
	return ""
}
