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
	retryTimeout := func(seconds string) func(*step) {
		return func(s *step) {
			s.Config = strings.Replace(s.Config, "retry_timeout_seconds = 5", seconds, 1)
		}
	}
	const (
		groups = `/api/2\.0/preview/scim/v2/Groups`
		create = `call 1 \(POST ` + groups + `\)`
	)
	for _, tc := range []struct {
		name, journey, failure string
		edit                   func(*step)
	}{
		{"429 twice in a 2 s retry timeout", "retry_429_until_timeout", "",
			func(s *step) {
				retryTimeout("retry_timeout_seconds = 2")(s)
				s.Calls[0].Times = httpfixture.Exactly(2)
			}},
		{"429 retried past the range without a retry timeout", "retry_429_until_timeout",
			`step 1 \(create_gives_up_at_retry_timeout\): ` +
				`unexpected HTTP call: POST ` + groups,
			retryTimeout("")},
		{"retry later than max_gap", "retry_503",
			`step 1 \(create_retries_503\): call 2: POST ` + groups +
				` arrived [0-9.]+s after ` + create + `, above max_gap 1s`,
			func(s *step) { s.Calls[1].MinGap, s.Calls[1].MaxGap = "", "1s" }},
		{"rate_limit delays the read after the create", "retry_503", "", withRateLimit},
	} {
		t.Run(tc.name, func(t *testing.T) {
			d, err := loadDefinition("testdata/repeats/" + tc.journey + ".yaml")
			require.NoError(t, err)
			tc.edit(&d.Steps[0])
			contents, err := yaml.Marshal(d)
			require.NoError(t, err)
			directory := t.TempDir()
			writeDiscoveryFile(t, directory, tc.journey+".yaml", string(contents))
			output, err := discoverySubprocess(t, directory,
				"^TestDirectoryRunnerHelper$")
			if tc.failure == "" {
				require.NoError(t, err, output)
				return
			}
			require.Error(t, err, output)
			require.Regexp(t, tc.failure, output)
		})
	}
}

// withRateLimit sets rate_limit = 1, drops the 503s and expects the read at least
// 900ms after the create, which the SDK's rate limiter only does when the runner
// keeps rate_limit. A slow machine only lengthens the gap.
func withRateLimit(s *step) {
	s.Config = "provider \"databricks\" {\n  rate_limit = 1\n}\n" + s.Config
	s.Calls = s.Calls[1:]
	s.Calls[0].MinGap, s.Calls[0].MaxGap = "", ""
	s.Calls[1].MinGap = "900ms"
}

func TestRecordingRejectsRepeatFields(t *testing.T) {
	for _, field := range []string{"times: 2", "min_gap: 1s", "max_gap: 1s"} {
		t.Run(field, func(t *testing.T) {
			update := "  - name: update\n" +
				"    config: 'provider \"databricks\" {}'\n" +
				"    calls:\n" +
				"      - {request: {method: GET, path: /first}, " +
				"response: {status: 200}}\n" +
				"      - {request: {method: GET, path: /repeated}, " +
				"response: {status: 200}, " + field + "}\n" +
				"destroy:\n"
			definition := strings.Replace(minimalDefinition, "destroy:\n", update, 1)
			d, err := decodeDefinition(strings.NewReader(definition))
			require.NoError(t, err)
			require.ErrorContains(t, validateRecordableCalls(d),
				"step 2 (update): call 2: recording cannot reproduce times, "+
					"min_gap or max_gap")
		})
	}
}

func TestRepeatValidationNamesFileStepAndCall(t *testing.T) {
	calls := "    calls:\n" +
		"      - {request: {method: GET, path: /retry}, response: {status: 429}, " +
		"times: {min: 3, max: 2}}\n"
	directory := t.TempDir()
	definition := strings.Replace(minimalDefinition, "    calls: []\n", calls, 1)
	writeDiscoveryFile(t, directory, "Repeat.yaml", definition)
	output, err := discoverySubprocess(t, directory, "^TestDirectoryRunnerHelper$")
	require.Error(t, err, output)
	require.Contains(t, output, "Repeat.yaml: step 1 (create): call 1: "+
		"times {min: 3, max: 2} needs 0 <= min <= max and max >= 1")
}
