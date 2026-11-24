package terraformtest

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"github.com/databricks/terraform-provider-databricks/qa/httpfixture"
)

func TestFaultReplay(t *testing.T) {
	Run(t, "testdata/faults")
}

// A faulted call followed by a normal one pins a retry; a faulted call followed
// by nothing pins that there is none. Each journey fails if that changes.
func TestFaultJourneysPinRetries(t *testing.T) {
	retried, err := loadDefinition("testdata/faults/reset.yaml")
	require.NoError(t, err)
	successfulPOST := retried.Steps[0].Calls[1]
	for _, tc := range []struct {
		name, journey, message string
		edit                   func(*step)
	}{
		{"reset without its retry", "reset", "step 1 (create_retries_after_reset): unexpected HTTP call: POST /api/2.0/preview/scim/v2/Groups",
			func(s *step) { s.Calls = append(s.Calls[:1], s.Calls[2:]...) }},
		{"close with a retry", "close", "step 1 (create_fails_with_eof): unconsumed call 2: POST /api/2.0/preview/scim/v2/Groups",
			func(s *step) { s.Calls = append(s.Calls, successfulPOST) }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			d, err := loadDefinition("testdata/faults/" + tc.journey + ".yaml")
			require.NoError(t, err)
			tc.edit(&d.Steps[0])
			contents, err := yaml.Marshal(d)
			require.NoError(t, err)
			directory := t.TempDir()
			writeDiscoveryFile(t, directory, tc.journey+".yaml", string(contents))
			output, err := discoverySubprocess(t, directory, "^TestDirectoryRunnerHelper$")
			require.Error(t, err, output)
			require.Contains(t, output, tc.message)
		})
	}
}

const faultedCall = `      - request: {method: GET, path: /faulted}
        response: {fault: timeout}
`

func TestFaultValidationNamesFileStepAndCall(t *testing.T) {
	invalidDestroy := strings.Replace(faultedCall, "{fault: timeout}", "{status: 200, delay: soon}", 1)
	for _, tc := range []struct {
		name, from, to, message string
	}{
		{"step", "    calls: []\n", "    calls:\n" + faultedCall,
			`Fault.yaml: step 1 (create): call 1 response: unknown fault "timeout"`},
		{"destroy", "destroy:\n  calls: []\n", "destroy:\n  calls:\n" + invalidDestroy,
			`Fault.yaml: destroy: call 1 response: delay "soon" must be a positive duration`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			directory := t.TempDir()
			writeDiscoveryFile(t, directory, "Fault.yaml", strings.Replace(minimalDefinition, tc.from, tc.to, 1))
			output, err := discoverySubprocess(t, directory, "^TestDirectoryRunnerHelper$")
			require.Error(t, err, output)
			require.Contains(t, output, tc.message)
		})
	}
}

func TestRecordingRejectsScriptedWireBehavior(t *testing.T) {
	for _, field := range []string{"fault: close", "delay: 1s", "status_line: HTTP/1.1 200"} {
		t.Run(field, func(t *testing.T) {
			response := "{" + field + "}"
			if field != "fault: close" {
				response = "{status: 200, " + field + "}"
			}
			update := "  - name: update\n    config: 'provider \"databricks\" {}'\n    calls:\n" +
				"      - request: {method: GET, path: /scripted}\n        response: " + response + "\ndestroy:\n"
			d, err := decodeDefinition(strings.NewReader(strings.Replace(minimalDefinition, "destroy:\n", update, 1)))
			require.NoError(t, err)
			require.ErrorContains(t, validateRecordableCalls(d),
				"step 2 (update): call 1: recording cannot reproduce a response delay, fault or status_line")
		})
	}
	d, err := loadDefinition("testdata/faults/reset.yaml")
	require.NoError(t, err)
	d.Steps[0].Calls = d.Steps[0].Calls[1:]
	require.NoError(t, validateRecordableCalls(d))
}

// The replay transport opens one connection per request, so net/http cannot
// resend a faulted request on a fresh connection before the SDK sees the fault.
func TestFixtureTransportReportsFaultsAfterEarlierCalls(t *testing.T) {
	ok := "ok"
	handler := httpfixture.NewHandler()
	require.NoError(t, handler.StartStep("faults", []httpfixture.Call{
		{Request: httpfixture.Request{Method: "GET", Path: "/first"}, Response: httpfixture.Response{Status: 200, Body: &httpfixture.Body{Text: &ok}}},
		{Request: httpfixture.Request{Method: "GET", Path: "/close"}, Response: httpfixture.Response{Fault: httpfixture.FaultClose}},
	}))
	server := httptest.NewServer(handler)
	defer server.Close()
	cfg := &config.Config{}
	require.NoError(t, fixtureConfig(server.URL)(cfg))
	client := &http.Client{Transport: cfg.HTTPTransport}
	resp, err := client.Get(server.URL + "/first")
	require.NoError(t, err)
	_, _ = io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	_, err = client.Get(server.URL + "/close")
	require.ErrorIs(t, err, io.EOF)
	require.NoError(t, handler.Verify())
}

// Recording must reject scripted wire behavior before it loads a profile, so a
// faulted journey never reaches a live workspace or has its YAML overwritten.
func TestRecordingRejectsScriptedWireBehaviorBeforeLoadingProfile(t *testing.T) {
	directory := t.TempDir()
	original := strings.Replace(minimalDefinition, "    calls: []\n",
		"    calls:\n      - request: {method: GET, path: /scripted}\n        response: {fault: close}\n", 1)
	writeDiscoveryFile(t, directory, "One.yaml", original)
	output, err := recorderSubprocess(t, directory, "missing-profile")
	require.Error(t, err, output)
	require.Contains(t, output, "step 1 (create): call 1: recording cannot reproduce a response delay, fault or status_line")
	require.NotContains(t, output, "cannot load recorder profile")
	after, err := os.ReadFile(filepath.Join(directory, "One.yaml"))
	require.NoError(t, err)
	require.Equal(t, original, string(after))
}
