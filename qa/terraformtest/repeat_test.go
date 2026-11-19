package terraformtest

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"github.com/databricks/terraform-provider-databricks/qa/httpfixture"
)

func TestRepeatReplay(t *testing.T) {
	Run(t, "testdata/repeats")
}

// Each edit runs a repeat journey in a subprocess. The passing ones pin that
// the runner keeps retry_timeout_seconds and rate_limit from the provider block,
// and the failing ones that the journeys catch a change in retries or timing.
func TestRepeatJourneysPinRetriesAndTiming(t *testing.T) {
	withRateLimit := func(s *step) {
		s.Config = "provider \"databricks\" {\n  rate_limit = 1\n}\n" + s.Config
		readAfterCreate(s)
	}
	for _, tc := range []struct {
		name, journey, message string
		edit                   func(*step)
	}{
		{"429 twice in a 2 s retry timeout", "retry_429_until_timeout", "", func(s *step) {
			s.Config = strings.Replace(s.Config, "retry_timeout_seconds = 5", "retry_timeout_seconds = 2", 1)
			s.Calls[0].Times = httpfixture.Exactly(2)
		}},
		{"429 retried past the range without a retry timeout", "retry_429_until_timeout",
			"step 1 (create_gives_up_at_retry_timeout): unexpected HTTP call: POST /api/2.0/preview/scim/v2/Groups",
			func(s *step) { s.Config = strings.Replace(s.Config, "retry_timeout_seconds = 5", "", 1) }},
		{"retry later than max_gap", "retry_503",
			"step 1 (create_retries_503): call 2: POST /api/2.0/preview/scim/v2/Groups arrived 2",
			func(s *step) { s.Calls[1].MinGap, s.Calls[1].MaxGap = "", "1s" }},
		{"rate_limit delays the read after the create", "retry_503", "", withRateLimit},
		{"read without rate_limit is sooner than min_gap", "retry_503",
			"step 1 (create_retries_503): call 2: GET /api/2.0/preview/scim/v2/Groups/group-id?attributes=displayName%2CexternalId%2Centitlements arrived",
			readAfterCreate},
	} {
		t.Run(tc.name, func(t *testing.T) {
			d, err := loadDefinition("testdata/repeats/" + tc.journey + ".yaml")
			require.NoError(t, err)
			tc.edit(&d.Steps[0])
			contents, err := yaml.Marshal(d)
			require.NoError(t, err)
			directory := t.TempDir()
			writeDiscoveryFile(t, directory, tc.journey+".yaml", string(contents))
			output, err := discoverySubprocess(t, directory, "^TestDirectoryRunnerHelper$")
			if tc.message == "" {
				require.NoError(t, err, output)
				return
			}
			require.Error(t, err, output)
			require.Contains(t, output, tc.message)
		})
	}
}

// readAfterCreate drops the 503s and expects the read at least 900ms after the
// create. The SDK's rate limiter spaces them only when rate_limit = 1 is kept.
func readAfterCreate(s *step) {
	s.Calls = s.Calls[1:]
	s.Calls[0].MinGap, s.Calls[0].MaxGap = "", ""
	s.Calls[1].MinGap = "900ms"
}

func TestRecordingRejectsRepeatFields(t *testing.T) {
	for _, field := range []string{"times: 2", "min_gap: 1s", "max_gap: 1s"} {
		t.Run(field, func(t *testing.T) {
			update := "  - name: update\n    config: 'provider \"databricks\" {}'\n    calls:\n" +
				"      - {request: {method: GET, path: /first}, response: {status: 200}}\n" +
				"      - {request: {method: GET, path: /repeated}, response: {status: 200}, " + field + "}\ndestroy:\n"
			d, err := decodeDefinition(strings.NewReader(strings.Replace(minimalDefinition, "destroy:\n", update, 1)))
			require.NoError(t, err)
			require.ErrorContains(t, validateRecordableCalls(d),
				"step 2 (update): call 2: recording cannot reproduce times, min_gap or max_gap")
		})
	}
}

func TestRepeatValidationNamesFileStepAndCall(t *testing.T) {
	calls := "    calls:\n      - {request: {method: GET, path: /retry}, response: {status: 429}, times: {min: 3, max: 2}}\n"
	directory := t.TempDir()
	writeDiscoveryFile(t, directory, "Repeat.yaml", strings.Replace(minimalDefinition, "    calls: []\n", calls, 1))
	output, err := discoverySubprocess(t, directory, "^TestDirectoryRunnerHelper$")
	require.Error(t, err, output)
	require.Contains(t, output, "Repeat.yaml: step 1 (create): call 1: times {min: 3, max: 2} needs 0 <= min <= max and max >= 1")
}
