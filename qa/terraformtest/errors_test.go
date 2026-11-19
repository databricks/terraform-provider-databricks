package terraformtest

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/qa/httpfixture"
)

func TestExpectedErrorReplay(t *testing.T) {
	Run(t, "testdata/expected_errors")
}

func TestExpectedErrorContinuationAndCleanup(t *testing.T) {
	Run(t, "testdata/expected_errors/engine")
}

func TestDecodeExpectedErrors(t *testing.T) {
	for _, tc := range []struct {
		name, value, message string
	}{
		{"valid", "'specific provider error'", ""},
		{"invalid pattern", "'['", "invalid expect_error regular expression"},
		{"empty pattern", "''", "nonempty regular expression"},
		{"whitespace", "'   '", "nonempty regular expression"},
		{"wrong type", "{}", "cannot unmarshal"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			input := "version: 1\nsteps:\n  - name: rejected\n    config: 'provider \"databricks\" {}'\n    expect_error: " + tc.value + "\n"
			d, err := decodeDefinition(strings.NewReader(input))
			if tc.message != "" {
				require.ErrorContains(t, err, tc.message)
				return
			}
			require.NoError(t, err)
			require.Equal(t, "specific provider error", *d.Steps[0].ExpectError)
			steps := testSteps(t, d, httpfixture.NewHandler())
			require.True(t, steps[0].ExpectError.MatchString("specific provider error"))
			require.Nil(t, steps[1].ExpectError)
			require.True(t, steps[1].Destroy)
		})
	}
}

func TestExpectedErrorDoesNotSilentlySkipStateChecks(t *testing.T) {
	input := strings.Replace(minimalDefinition, "    calls: []", "    expect_error: 'rejected'\n    calls: []", 1)
	_, err := decodeDefinition(strings.NewReader(input))
	require.ErrorContains(t, err, "step 1 (create): state checks on expected-error steps are not supported")
}

func TestExpectedErrorCannotHideFixtureFailures(t *testing.T) {
	contents, err := os.ReadFile("testdata/expected_errors/api_error.yaml")
	require.NoError(t, err)
	for _, tc := range []struct {
		name, fixture, message string
	}{
		{"wrong request", strings.Replace(string(contents), "displayName: Example group", "displayName: Wrong group", 1), "unexpected HTTP call"},
		{"unused call", string(contents), "unconsumed call"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			fixture := strings.Replace(tc.fixture, "'group rejected by fixture'", "'.*'", 1)
			if tc.name == "unused call" {
				fixture = strings.Replace(fixture, "destroy:\n", `      - request: {method: GET, path: /unused}
        response: {status: 200}
destroy:
`, 1)
			}
			directory := t.TempDir()
			writeDiscoveryFile(t, directory, "Error.yaml", fixture)
			output, err := discoverySubprocess(t, directory, "^TestDirectoryRunnerHelper$")
			require.Error(t, err, output)
			require.Contains(t, output, tc.message)
		})
	}
}

func TestExpectedErrorMustOccurAndMatch(t *testing.T) {
	for _, tc := range []struct {
		name, hcl, pattern, message string
	}{
		{"missing error", `provider "databricks" {}`, "must fail", "expected an error but got none"},
		{"wrong error", `resource "databricks_group" "this" { display_name = "admins" }`, "different diagnostic", "no match on"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			directory := t.TempDir()
			writeDiscoveryFile(t, directory, "Error.yaml", "version: 1\nsteps:\n  - name: rejected\n    config: '"+tc.hcl+"'\n    expect_error: '"+tc.pattern+"'\ndestroy:\n  calls: []\n")
			output, err := discoverySubprocess(t, directory, "^TestDirectoryRunnerHelper$")
			require.Error(t, err, output)
			require.Contains(t, output, tc.message)
		})
	}
}
