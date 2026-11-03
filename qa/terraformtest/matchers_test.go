package terraformtest

import (
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"github.com/databricks/terraform-provider-databricks/qa/httpfixture"
)

const matcherCalls = `      - request:
          method: POST
          path: /requests
          headers: {X-Request-ID: captured-header}
          query: {request_id: [captured-query], fixed: [exact]}
          body:
            json:
              request_id: captured-body
              items: [{token: null, name: fixed}]
          matchers:
            headers: {X-Request-ID: any_string}
            query: {request_id: any_string}
            json:
              - path: [request_id]
                type: any_string
              - path: [items, 0, token]
                type: any_string
        response:
          status: 200
          body:
            json: {request_id: any_string}
`

func TestRequestMatcherReplay(t *testing.T) {
	Run(t, "testdata/matchers")
}

func TestRequestMatcherShowcaseRequiresMatcher(t *testing.T) {
	for _, tc := range []struct {
		name, signature string
	}{
		{"body", "POST /api/2.0/preview/scim/v2/Groups"},
		{"header", "POST /api/2.0/preview/scim/v2/Groups"},
		{"query", "GET /api/2.0/preview/scim/v2/Groups/group-id"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			d, err := loadDefinition("testdata/matchers/group_create_" + tc.name + ".yml")
			require.NoError(t, err)
			annotations := 0
			for i := range d.Steps[0].Calls {
				request := &d.Steps[0].Calls[i].Request
				annotations += len(request.Matchers.Headers) + len(request.Matchers.Query) + len(request.Matchers.JSON)
				request.Matchers = httpfixture.Matchers{}
			}
			require.Equal(t, 1, annotations, "each showcase exercises exactly one matcher")
			fixture, err := yaml.Marshal(d)
			require.NoError(t, err)
			directory := t.TempDir()
			writeDiscoveryFile(t, directory, "group_create.yml", string(fixture))
			output, err := discoverySubprocess(t, directory, "^TestDirectoryRunnerHelper$")
			require.Error(t, err, output)
			require.Contains(t, output, "step 1 (create): unexpected HTTP call: "+tc.signature)
		})
	}
}

func TestDecodeAndReplayRequestMatchers(t *testing.T) {
	input := strings.Replace(minimalDefinition, "    calls: []\n", "    calls:\n"+matcherCalls, 1)
	d, err := decodeDefinition(strings.NewReader(input))
	require.NoError(t, err)
	h := httpfixture.NewHandler()
	require.NoError(t, h.StartStep("create", d.Steps[0].Calls))
	req := httptest.NewRequest("POST", "/requests?fixed=exact&request_id=new-query", strings.NewReader(
		`{"request_id":"new-body","items":[{"token":"new-token","name":"fixed"}]}`))
	req.Header.Set("X-Request-ID", "new-header")
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	require.Equal(t, 200, w.Code, w.Body.String())
	require.JSONEq(t, `{"request_id":"any_string"}`, w.Body.String(), "responses remain literal")
	require.NoError(t, h.Verify())
}

func TestYAMLMatcherEmptyPathSelectsWholeBody(t *testing.T) {
	input := strings.Replace(minimalDefinition, "    calls: []\n", `    calls:
      - request:
          method: POST
          path: /value
          body: {json: captured-string}
          matchers:
            json: [{path: [], type: any_string}]
        response: {status: 204}
`, 1)
	d, err := decodeDefinition(strings.NewReader(input))
	require.NoError(t, err)
	h := httpfixture.NewHandler()
	require.NoError(t, h.StartStep("create", d.Steps[0].Calls))
	w := httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest("POST", "/value", strings.NewReader(`"new-string"`)))
	require.Equal(t, 204, w.Code)
	require.NoError(t, h.Verify())
}

func TestYAMLRejectsInvalidMatcherFields(t *testing.T) {
	for _, tc := range []struct {
		name, before, after, message string
	}{
		{"unknown location", "            headers: {", "            header: {", "field header not found"},
		{"unknown matcher", "query: {request_id: any_string}", "query: {request_id: any_number}", "expected a nonempty key and matcher any_string"},
		{"missing JSON path", "path: [request_id]", "path: null", "explicit path and type any_string"},
		{"unknown JSON field", "type: any_string", "match: any_string", "field match not found"},
		{"wrong index type", "items, 0, token", "items, 0.5, token", "path segments must be"},
		{"missing sample", "path: [request_id]", "path: [absent]", "is missing from sample"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			calls := strings.Replace(matcherCalls, tc.before, tc.after, 1)
			require.NotEqual(t, matcherCalls, calls)
			input := strings.Replace(minimalDefinition, "    calls: []\n", "    calls:\n"+calls, 1)
			_, err := decodeDefinition(strings.NewReader(input))
			require.ErrorContains(t, err, tc.message)
		})
	}
}

func TestMatchersValidateLaterStepsAndDestroyBeforeTerraform(t *testing.T) {
	invalid := strings.Replace(matcherCalls, "path: [request_id]", "path: [absent]", 1)
	for _, tc := range []struct {
		name, suffix, label string
	}{
		{"later step", "  - name: later\n    config: |\n      provider \"databricks\" {}\n    calls:\n" + invalid + "destroy:\n  calls: []\n", "step 2 (later)"},
		{"destroy", "destroy:\n  calls:\n" + invalid, "destroy"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			input := strings.Replace(minimalDefinition, "destroy:\n  calls: []\n", tc.suffix, 1)
			_, err := decodeDefinition(strings.NewReader(input))
			require.ErrorContains(t, err, tc.label+": call 1 request matchers:")
			require.ErrorContains(t, err, "is missing from sample")
		})
	}
}

func TestRequestMatcherWithRealProvider(t *testing.T) {
	contents, err := os.ReadFile("testdata/expected_errors/api_error.yaml")
	require.NoError(t, err)
	// The SDK supplies a version-dependent User-Agent. Require its presence
	// without coupling this replay to the current provider or SDK version.
	input := strings.Replace(string(contents), "          body:\n", `          headers:
            User-Agent: captured-on-another-version
          matchers:
            headers:
              User-Agent: any_string
          body:
`, 1)
	require.NotEqual(t, string(contents), input)
	d, err := decodeDefinition(strings.NewReader(input))
	require.NoError(t, err)
	run(t, d)
}
