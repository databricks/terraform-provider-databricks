package terraformtest

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"github.com/databricks/terraform-provider-databricks/qa/httpfixture"
)

func TestInitialStateRecordingGeneratesAndReplaysNativeCheckpoint(t *testing.T) {
	original, err := os.ReadFile("testdata/initial_state/rejected_update.yaml")
	require.NoError(t, err)
	var document yaml.Node
	require.NoError(t, yaml.Unmarshal(original, &document))
	require.NoError(t, setYAMLField(document.Content[0], "recording", recordingSetup{
		InitialStateTemplate: "resource \"databricks_group\" \"this\" {\n  display_name = \"Example group\"\n}\n",
	}))
	original, err = yaml.Marshal(&document)
	require.NoError(t, err)
	peer := recordingGroupPeer(t)
	server := httptest.NewServer(peer)
	t.Cleanup(server.Close)
	directory := localRecordingFiles(t, original, server.URL)
	// Existing historical snapshots are never edited or replaced in place.
	writeDiscoveryFile(t, directory, "checkpoints/group.tfstate", "historical-checkpoint")
	output, err := recorderSubprocess(t, directory, "local")
	require.NoError(t, err, output)
	updated, err := loadDefinition(filepath.Join(directory, "One.yaml"))
	require.NoError(t, err)
	require.NotEqual(t, "checkpoints/group.tfstate", updated.InitialState.Checkpoint)
	require.NotContains(t, string(updated.checkpoint), "9876543210123456")
	require.NotContains(t, string(updated.checkpoint), "recorder-secret")
	require.Contains(t, string(updated.checkpoint), `"workspace_id": "12345"`)
	prior, err := os.ReadFile(filepath.Join(directory, "checkpoints/group.tfstate"))
	require.NoError(t, err)
	require.Equal(t, "historical-checkpoint", string(prior))
	server.Close()
	run(t, updated)
}

func TestRecordingEnvironmentOwnsSeparateStateAndCleansUpAfterJourney(t *testing.T) {
	for _, fail := range []bool{false, true} {
		t.Run(fmt.Sprintf("failed_check_%t", fail), func(t *testing.T) {
			peer := recordingGroupPeer(t)
			var mu sync.Mutex
			var events []string
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				mu.Lock()
				events = append(events, req.Method+" "+req.URL.Path)
				mu.Unlock()
				peer.ServeHTTP(w, req)
			}))
			t.Cleanup(server.Close)
			// A real environment owns the group; the tested journey only reads it.
			// If state were shared, the data-only apply could delete the group.
			source := `version: 1
host_metadata: {workspace_id: "12345"}
recording:
  environment_template: |
    resource "databricks_group" "background" {
      display_name = "Example group"
    }
steps:
  - name: read_background
    config: |
      data "databricks_group" "this" {
        display_name = "Example group"
      }
    calls: []
    state:
      resources:
        data.databricks_group.this:
          attributes:
            - path: [id]
              equals: group-id
destroy: {calls: []}
`
			if fail {
				source = string(bytes.ReplaceAll([]byte(source), []byte("equals: group-id"), []byte("equals: incorrect")))
			}
			directory := localRecordingFiles(t, []byte(source), server.URL)
			output, recordErr := recorderSubprocess(t, directory, "local")
			require.NoError(t, peer.Verify(), output)
			mu.Lock()
			got := append([]string(nil), events...)
			mu.Unlock()
			require.Equal(t, "DELETE /api/2.0/preview/scim/v2/Groups/group-id", got[len(got)-1], output)
			updated, err := os.ReadFile(filepath.Join(directory, "One.yaml"))
			require.NoError(t, err)
			if fail {
				require.Error(t, recordErr, output)
				require.Equal(t, source, string(updated))
				return
			}
			require.NoError(t, recordErr, output)
			d, err := loadDefinition(filepath.Join(directory, "One.yaml"))
			require.NoError(t, err)
			for _, call := range d.Destroy.Calls {
				require.NotEqual(t, "DELETE", call.Request.Method)
			}
			server.Close()
			run(t, d)
		})
	}
}

func TestCheckpointPublicationPreservesOldSourceOnConcurrentEdits(t *testing.T) {
	directory := t.TempDir()
	path := filepath.Join(directory, "test.yaml")
	updated := []byte(minimalDefinition + "initial_state: {checkpoint: snapshots/group.tfstate}\n")
	require.NoError(t, os.WriteFile(path, []byte("concurrent edit"), 0600))
	err := saveRecordedCheckpoint(path, []byte("original"), updated, []byte("new snapshot"), "snapshots/group.tfstate")
	require.ErrorContains(t, err, "changed while recording")
	files, err := os.ReadDir(filepath.Join(directory, "snapshots"))
	require.NoError(t, err)
	require.Empty(t, files)
	name := checkpointFilename("group.tfstate", []byte("snapshot"))
	require.Equal(t, name, checkpointFilename(name, []byte("snapshot")))
}

// Local fixtures stand in for Databricks so record-time ownership and generated
// checkpoints can be tested without running a live integration test.
func recordingGroupPeer(t *testing.T) *httpfixture.Handler {
	t.Helper()
	d, err := loadDefinition("testdata/matchers/group_create_body.yml")
	require.NoError(t, err)
	d.HostMetadata["workspace_id"] = "9876543210123456"
	calls := append(d.metadataCalls(), d.Steps[0].Calls...)
	calls = append(calls, d.Destroy.Calls...)
	calls = append(calls, httpfixture.Call{Request: httpfixture.Request{Method: "GET", Path: "/api/2.0/preview/scim/v2/Groups",
		Query: map[string][]string{"filter": {`displayName eq "Example group"`}}},
		Response: httpfixture.Response{Status: 200, Body: &httpfixture.Body{JSON: map[string]any{
			"Resources": []any{map[string]any{"id": "group-id", "displayName": "Example group"}}, "totalResults": 1}}}})
	peer := httpfixture.NewHandler()
	require.NoError(t, peer.SetReusableFixtures(calls))
	return peer
}

func localRecordingFiles(t *testing.T, original []byte, host string) string {
	t.Helper()
	directory := t.TempDir()
	writeDiscoveryFile(t, directory, "One.yaml", string(original))
	writeDiscoveryFile(t, directory, "profiles", "[local]\nhost = "+host+"\ntoken = recorder-secret\nworkspace_id = 9876543210123456\n")
	t.Setenv("DATABRICKS_CONFIG_FILE", filepath.Join(directory, "profiles"))
	return directory
}
