package terraformtest

import (
	"bytes"
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"github.com/databricks/terraform-provider-databricks/qa/httpfixture"
)

func TestRecordingRequiresExactSelectionAndRejectsCI(t *testing.T) {
	for _, name := range []string{"CI", "GITHUB_ACTIONS", "TF_BUILD", "BUILDKITE", "JENKINS_URL"} {
		t.Setenv(name, "")
	}
	files := []testFile{{name: "One"}, {name: "Two"}}
	for _, filter := range []string{"", "TestReplay/One", "^TestReplay$/.*", "^TestReplay$/^(One|Two)$"} {
		require.ErrorContains(t, recordingSelection("TestReplay", files, "dev", filter), "one exact anchored")
	}
	require.NoError(t, recordingSelection("TestReplay", files, "dev", "^TestReplay$/^One$"))
	t.Setenv("CI", "true")
	require.ErrorContains(t, recordingSelection("TestReplay", files, "dev", "^TestReplay$/^One$"), "forbidden in CI")
	require.NoError(t, recordingSelection("TestReplay", files, "", ""))
}

func TestBasicRecordingRejectsSetupBeforeLoadingProfile(t *testing.T) {
	for _, setup := range []string{"initial_state: {checkpoint: missing.tfstate}\n", "recording: {environment_template: ''}\n"} {
		directory := t.TempDir()
		original := minimalDefinition + setup
		writeDiscoveryFile(t, directory, "One.yaml", original)
		output, err := recorderSubprocess(t, directory, "missing-profile")
		require.Error(t, err, output)
		require.Contains(t, output, "basic recording does not support")
		after, err := os.ReadFile(filepath.Join(directory, "One.yaml"))
		require.NoError(t, err)
		require.Equal(t, original, string(after))
	}
}

func TestRecordingProfileIgnoresAmbientCredentials(t *testing.T) {
	directory := t.TempDir()
	writeDiscoveryFile(t, directory, "profiles", "[selected]\nhost = http://localhost:9999\ntoken = selected-token\nworkspace_id = 9876543210123456\n")
	t.Setenv("DATABRICKS_CONFIG_FILE", filepath.Join(directory, "profiles"))
	t.Setenv("DATABRICKS_HOST", "https://ambient.invalid")
	t.Setenv("DATABRICKS_TOKEN", "ambient-token")
	profile, err := loadRecordingProfile("selected")
	require.NoError(t, err)
	cfg := &config.Config{Host: "https://yaml.invalid", Token: "yaml-token"}
	require.NoError(t, profile.customize(http.DefaultTransport)(cfg))
	require.Equal(t, "http://localhost:9999", cfg.Host)
	require.Equal(t, "selected-token", cfg.Token)
	require.Equal(t, "9876543210123456", cfg.WorkspaceID)
	_, err = loadRecordingProfile("missing")
	require.ErrorContains(t, err, "cannot load recorder profile")
	require.NotContains(t, err.Error(), "selected-token")
}

func TestRecordingNormalizesKnownIDsAndRedactsPayloads(t *testing.T) {
	d := &testDefinition{HostMetadata: map[string]any{"workspace_id": "12345", "account_id": "replay-account"}}
	n, err := newRecordingNormalizer(d, &config.Config{WorkspaceID: "9876543210123456", AccountID: "real-account-id", Token: "private-token"})
	require.NoError(t, err)
	data := []byte("workspace_id: 9876543210123456\naccount_id: real-account-id\nurl: '/9876543210123456/9876543210123456?account=real-account-id'\nother: x9876543210123456x\npassword: secret-password\nmessage: 'Bearer private-token'\nbase64: " + base64.StdEncoding.EncodeToString([]byte("private-token 9876543210123456")) + "\n")
	normalized, err := n.transform(data, false, true)
	require.NoError(t, err)
	var values map[string]any
	require.NoError(t, yaml.Unmarshal(normalized, &values))
	require.Equal(t, 12345, values["workspace_id"])
	require.Equal(t, "replay-account", values["account_id"])
	require.Equal(t, "/12345/12345?account=replay-account", values["url"])
	require.Equal(t, "x9876543210123456x", values["other"])
	require.Equal(t, "[REDACTED]", values["password"])
	require.Equal(t, "Bearer [REDACTED]", values["message"])
	binary, err := base64.StdEncoding.DecodeString(values["base64"].(string))
	require.NoError(t, err)
	require.Equal(t, "[REDACTED] 12345", string(binary))
	live, err := n.transform(normalized, true, false)
	require.NoError(t, err)
	require.Contains(t, string(live), "workspace_id: 9876543210123456")
	_, err = newRecordingNormalizer(d, &config.Config{WorkspaceID: "12345"})
	require.ErrorContains(t, err, "different from the recorder profile")
}

// Record against a deterministic local peer, then replay with the peer closed.
// This exercises real Terraform, provider configuration, capture, YAML refresh,
// matcher preservation, ID normalization and mandatory cleanup together.
func TestRecorderRoundTripAndFailurePreservesOriginal(t *testing.T) {
	for _, fail := range []bool{false, true} {
		t.Run(map[bool]string{false: "round_trip", true: "failed_check"}[fail], func(t *testing.T) {
			original, err := os.ReadFile("testdata/matchers/group_create_body.yml")
			require.NoError(t, err)
			original = bytes.ReplaceAll(original, []byte("group-id"), []byte("captured-id"))
			if fail {
				original = bytes.Replace(original, []byte("destroy:\n"), []byte("    state:\n      resources:\n        databricks_group.this:\n          attributes:\n            - path: [display_name]\n              equals: incorrect\ndestroy:\n"), 1)
			}
			d, err := decodeDefinition(bytes.NewReader(original))
			require.NoError(t, err)
			peer := httpfixture.NewHandler()
			d.HostMetadata["workspace_id"] = "9876543210123456"
			calls := append(d.metadataCalls(), d.Steps[0].Calls...)
			calls = append(calls, d.Destroy.Calls...)
			require.NoError(t, peer.SetReusableFixtures(calls))
			var deletes atomic.Int32
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				if req.Method == "DELETE" {
					deletes.Add(1)
				}
				peer.ServeHTTP(w, req)
			}))
			t.Cleanup(server.Close)
			directory := t.TempDir()
			writeDiscoveryFile(t, directory, "One.yaml", string(original))
			writeDiscoveryFile(t, directory, "profiles", "[local]\nhost = "+server.URL+"\ntoken = recorder-secret\nworkspace_id = 9876543210123456\n")
			t.Setenv("DATABRICKS_CONFIG_FILE", filepath.Join(directory, "profiles"))
			output, recordErr := recorderSubprocess(t, directory, "local")
			require.Positive(t, deletes.Load(), output)
			updated, err := os.ReadFile(filepath.Join(directory, "One.yaml"))
			require.NoError(t, err)
			if fail {
				require.Error(t, recordErr, output)
				require.Equal(t, original, updated)
				return
			}
			require.NoError(t, recordErr, output)
			require.Contains(t, string(updated), "any_string")
			require.Contains(t, string(updated), "# JSON body matcher showcase")
			require.NotContains(t, string(updated), "9876543210123456")
			require.NotContains(t, string(updated), "recorder-secret")
			server.Close()
			replayed, err := loadDefinition(filepath.Join(directory, "One.yaml"))
			require.NoError(t, err)
			run(t, replayed)
		})
	}
}

func TestReplaceRecordingIsAtomicAndRefusesConcurrentEdits(t *testing.T) {
	path := filepath.Join(t.TempDir(), "test.yaml")
	require.NoError(t, os.WriteFile(path, []byte("before"), 0600))
	require.ErrorContains(t, replaceRecording(path, []byte("stale"), []byte("after")), "changed while recording")
	data, err := os.ReadFile(path)
	require.NoError(t, err)
	require.Equal(t, "before", string(data))
	require.NoError(t, replaceRecording(path, data, []byte("after")))
	data, err = os.ReadFile(path)
	require.NoError(t, err)
	require.Equal(t, "after", string(data))
	entries, err := os.ReadDir(filepath.Dir(path))
	require.NoError(t, err)
	require.Len(t, entries, 1)
}

func recorderSubprocess(t *testing.T, directory, profile string) (string, error) {
	t.Helper()
	executable, err := os.Executable()
	require.NoError(t, err)
	cmd := exec.Command(executable, "-test.v", "-test.run=^TestDirectoryRunnerHelper$/^One$")
	cmd.Env = append(os.Environ(), "TERRAFORMTEST_DISCOVERY_DIRECTORY="+directory, recorderProfileEnv+"="+profile)
	// Simulate a developer invocation inside CI. The selected profile points
	// only to this test's local peer; production entry points still reject CI.
	for _, name := range []string{"CI", "GITHUB_ACTIONS", "TF_BUILD", "BUILDKITE", "JENKINS_URL"} {
		cmd.Env = append(cmd.Env, name+"=")
	}
	output, err := cmd.CombinedOutput()
	return strings.TrimSpace(string(output)), err
}
