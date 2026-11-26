// Package headercheck checks the User-Agent and traceparent headers of databricks_volume
// calls. The volume replay package runs the checks on the existing client, and
// volume_sdk_go_headers runs them with the volume on sdk-go uc/volumes/v1. Each check
// replays a fixture from the native_testdata directory of the package that calls it.
package headercheck

import (
	"cmp"
	"encoding/json"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/useragent"
	"github.com/databricks/databricks-sdk-go/version"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

// userAgentEnvironment lists every variable from which the classic SDK derives a
// User-Agent dimension per request: CI/CD, Databricks Runtime, agent, meta-harness and
// upstream detection.
var userAgentEnvironment = []string{
	"GITHUB_ACTIONS", "GITLAB_CI", "JENKINS_URL", "TF_BUILD", "CIRCLECI", "TRAVIS",
	"BITBUCKET_BUILD_NUMBER", "PROJECT_ID", "BUILD_ID", "PROJECT_NUMBER", "LOCATION",
	"CODEBUILD_BUILD_ARN", "TFC_RUN_ID",
	"DATABRICKS_RUNTIME_VERSION",
	"AMP_CURRENT_THREAD_ID", "ANTIGRAVITY_AGENT", "AUGMENT_AGENT", "CLAUDECODE",
	"CLINE_ACTIVE", "CODEX_CI", "COPILOT_CLI", "CURSOR_AGENT", "GEMINI_CLI",
	"GOOSE_TERMINAL", "KIRO", "OPENCLAW_SHELL", "OPENCODE", "PI_CODING_AGENT",
	"VSCODE_AGENT", "WINDSURF_AGENT", "AGENT", "AI_AGENT",
	"OMNIGENT",
	"DATABRICKS_SDK_UPSTREAM", "DATABRICKS_SDK_UPSTREAM_VERSION",
}

const userAgentExtraVariable = "DATABRICKS_USER_AGENT_EXTRA"

// userAgentExtraChild marks the test process that UserAgentExtra starts.
const userAgentExtraChild = "DATABRICKS_TF_TEST_USER_AGENT_EXTRA_CHILD"

// UserAgentDimensions pins the exact User-Agent of every volume call of
// native_testdata/user_agent.yaml for each environment the classic SDK detects. With
// sdkGo set, the volume calls go through sdk-go uc/volumes/v1, which append the sdk-go
// module segments.
func UserAgentDimensions(t *testing.T, sdkGo bool) {
	requireNoUserAgentExtra(t)
	longAgent := strings.Repeat("x", 60) + "/abc/def"
	testCases := []struct {
		name     string
		env      map[string]string
		upstream []string
		detected []string
	}{
		{name: "no_detected_environment"},
		{
			name: "upstream_product_and_version",
			env: map[string]string{
				"DATABRICKS_SDK_UPSTREAM":         "databricks-cli",
				"DATABRICKS_SDK_UPSTREAM_VERSION": "0.250.0",
			},
			upstream: []string{"upstream/databricks-cli", "upstream-version/0.250.0"},
		},
		{
			name: "upstream_product_without_version",
			env:  map[string]string{"DATABRICKS_SDK_UPSTREAM": "databricks-cli"},
		},
		{
			name: "upstream_version_without_product",
			env:  map[string]string{"DATABRICKS_SDK_UPSTREAM_VERSION": "0.250.0"},
		},
		{
			name: "upstream_values_are_not_sanitized",
			env: map[string]string{
				"DATABRICKS_SDK_UPSTREAM":         "my tool/x",
				"DATABRICKS_SDK_UPSTREAM_VERSION": "1.0 (beta)",
			},
			upstream: []string{"upstream/my tool/x", "upstream-version/1.0 (beta)"},
		},
		{
			name:     "cicd_github_actions",
			env:      map[string]string{"GITHUB_ACTIONS": "true"},
			detected: []string{"cicd/github"},
		},
		{
			name: "cicd_value_must_match",
			env:  map[string]string{"GITHUB_ACTIONS": "1"},
		},
		{
			name:     "cicd_presence_with_empty_value",
			env:      map[string]string{"JENKINS_URL": ""},
			detected: []string{"cicd/jenkins"},
		},
		{
			name:     "cicd_first_listed_provider_wins",
			env:      map[string]string{"GITLAB_CI": "true", "GITHUB_ACTIONS": "true"},
			detected: []string{"cicd/github"},
		},
		{
			name:     "runtime_version",
			env:      map[string]string{"DATABRICKS_RUNTIME_VERSION": "15.4"},
			detected: []string{"runtime/15.4"},
		},
		{
			name: "runtime_version_is_sanitized",
			env: map[string]string{
				"DATABRICKS_RUNTIME_VERSION": "15.4 LTS (Scala 2.12)",
			},
			detected: []string{"runtime/15.4-LTS--Scala-2.12-"},
		},
		{
			name:     "agent_known",
			env:      map[string]string{"CLAUDECODE": "1"},
			detected: []string{"agent/claude-code"},
		},
		{
			name:     "agent_known_with_empty_value",
			env:      map[string]string{"CURSOR_AGENT": ""},
			detected: []string{"agent/cursor"},
		},
		{
			name:     "agent_multiple_known",
			env:      map[string]string{"CLAUDECODE": "1", "CURSOR_AGENT": "1"},
			detected: []string{"agent/multiple"},
		},
		{
			name:     "agent_known_wins_over_agent_variable",
			env:      map[string]string{"CLAUDECODE": "1", "AGENT": "goose"},
			detected: []string{"agent/claude-code"},
		},
		{
			name:     "agent_fallback_is_sanitized",
			env:      map[string]string{"AGENT": "My Agent/2.0 (beta)"},
			detected: []string{"agent/My-Agent-2.0--beta-"},
		},
		{
			name:     "agent_fallback_prefers_agent_over_ai_agent",
			env:      map[string]string{"AGENT": "goose", "AI_AGENT": "v0"},
			detected: []string{"agent/goose"},
		},
		{
			name:     "agent_fallback_to_ai_agent_when_agent_is_empty",
			env:      map[string]string{"AGENT": "", "AI_AGENT": "v0"},
			detected: []string{"agent/v0"},
		},
		{
			name:     "agent_fallback_is_capped_at_64_characters",
			env:      map[string]string{"AGENT": longAgent},
			detected: []string{"agent/" + strings.Repeat("x", 60) + "-abc"},
		},
		{
			name:     "meta_harness_known",
			env:      map[string]string{"OMNIGENT": "1"},
			detected: []string{"meta-harness/omnigent"},
		},
		{
			name: "every_detected_dimension",
			env: map[string]string{
				"DATABRICKS_SDK_UPSTREAM":         "databricks-cli",
				"DATABRICKS_SDK_UPSTREAM_VERSION": "0.250.0",
				"GITHUB_ACTIONS":                  "true",
				"DATABRICKS_RUNTIME_VERSION":      "15.4",
				"CLAUDECODE":                      "1",
				"OMNIGENT":                        "1",
			},
			upstream: []string{"upstream/databricks-cli", "upstream-version/0.250.0"},
			detected: []string{
				"cicd/github", "runtime/15.4", "agent/claude-code",
				"meta-harness/omnigent",
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			setUserAgentEnvironment(t, tc.env)
			requests := runVolumeJourney(t, "user_agent.yaml", sdkGo)
			want := wantUserAgent(t, nil, tc.upstream, tc.detected)
			assertUserAgent(t, requests, want, sdkGo)
		})
	}
}

// UserAgentExtra pins DATABRICKS_USER_AGENT_EXTRA on every volume call of
// native_testdata/user_agent.yaml, through sdk-go uc/volumes/v1 if sdkGo is set. The
// provider reads it once at start-up, so the check runs the calling test again alone in
// a new process.
func UserAgentExtra(t *testing.T, sdkGo bool) {
	const extra = "partner-a/1.2.3 team_b/0.0.1 partner-a/1.2.3 partner-a/2.0.0 " +
		"resource/volume"
	if os.Getenv(userAgentExtraChild) == "" {
		requireNoUserAgentExtra(t)
		child := exec.Command(os.Args[0], "-test.v", "-test.run=^"+t.Name()+"$")
		child.Env = append(os.Environ(),
			userAgentExtraVariable+"="+extra, userAgentExtraChild+"=1")
		output, err := child.CombinedOutput()
		require.NoError(t, err, "%s", output)
		passed := "--- PASS: " + t.Name()
		assert.Contains(t, string(output), passed, "%s", output)
		return
	}
	require.Equal(t, extra, os.Getenv(userAgentExtraVariable))
	setUserAgentEnvironment(t, nil)
	requests := runVolumeJourney(t, "user_agent.yaml", sdkGo)
	// Repeated pairs are dropped. A pair repeating a per-request dimension is kept.
	parsed := []string{
		"partner-a/1.2.3", "team_b/0.0.1", "partner-a/2.0.0", "resource/volume",
	}
	assertUserAgent(t, requests, wantUserAgent(t, parsed, nil, nil), sdkGo)
}

// Traceparent pins the W3C traceparent header of every volume call of
// native_testdata/traceparent_retry.yaml, including both attempts of each retried call,
// through sdk-go uc/volumes/v1 if sdkGo is set.
func Traceparent(t *testing.T, sdkGo bool) {
	requests := runVolumeJourney(t, "traceparent_retry.yaml", sdkGo)
	var calls []string
	for _, r := range requests {
		calls = append(calls, r.method+" "+r.path)
	}
	path := "/api/2.1/unity-catalog/volumes"
	volume := path + "/main.replay.traceparent_volume"
	require.Equal(t, []string{
		"POST " + path, "POST " + path, "GET " + volume, "GET " + volume, "GET " + volume,
		"DELETE " + volume, "DELETE " + volume,
	}, calls)
	// Version 00 with the sampled flag set. Every attempt, a retry included, starts a
	// new trace with a new parent ID.
	syntax := regexp.MustCompile(`^00-([0-9a-f]{32})-([0-9a-f]{16})-01$`)
	traceIDs, parentIDs := map[string]bool{}, map[string]bool{}
	for i, r := range requests {
		values := r.header.Values("traceparent")
		require.Len(t, values, 1, calls[i])
		match := syntax.FindStringSubmatch(values[0])
		require.NotNil(t, match, "%s: %s", calls[i], values[0])
		assert.NotEqual(t, strings.Repeat("0", 32), match[1], calls[i])
		assert.NotEqual(t, strings.Repeat("0", 16), match[2], calls[i])
		assert.False(t, traceIDs[match[1]], "%s reuses trace %s", calls[i], match[1])
		assert.False(t, parentIDs[match[2]], "%s reuses parent %s", calls[i], match[2])
		traceIDs[match[1]], parentIDs[match[2]] = true, true
		assert.Empty(t, r.header.Values("tracestate"), calls[i])
	}
	assertVolumeClient(t, requests, sdkGo)
}

// requireNoUserAgentExtra fails a test whose process started with
// DATABRICKS_USER_AGENT_EXTRA, which every User-Agent would then contain.
func requireNoUserAgentExtra(t *testing.T) {
	t.Helper()
	if value := os.Getenv(userAgentExtraVariable); value != "" {
		t.Fatalf("unset %s=%q: the provider adds it to every User-Agent",
			userAgentExtraVariable, value)
	}
}

// setUserAgentEnvironment replaces the detected environment with env, so that the
// agent or CI runner running the test doesn't change the User-Agent.
func setUserAgentEnvironment(t *testing.T, env map[string]string) {
	t.Helper()
	for _, name := range userAgentEnvironment {
		t.Setenv(name, "")
		require.NoError(t, os.Unsetenv(name))
	}
	for name, value := range env {
		t.Setenv(name, value)
	}
	// The SDK caches each detection for the life of the process.
	useragent.ClearCache()
	t.Cleanup(useragent.ClearCache)
}

// wantUserAgent returns the User-Agent of a volume call with the given extra, upstream
// and detected dimensions.
func wantUserAgent(t *testing.T, extra, upstream, detected []string) string {
	t.Helper()
	dimensions := []string{
		"databricks-tf-provider/" + common.Version(),
		"databricks-sdk-go/" + version.Version,
		"go/" + goUserAgentVersion(),
		"os/" + runtime.GOOS,
	}
	dimensions = append(dimensions, extra...)
	dimensions = append(dimensions, "terraform/"+terraformCLIVersion(t))
	dimensions = append(dimensions, upstream...)
	dimensions = append(dimensions, "sdk/sdkv2", "resource/volume", "auth/pat")
	return strings.Join(append(dimensions, detected...), " ")
}

// goUserAgentVersion returns the toolchain version as the three-part semantic version
// that the SDK reports for release toolchains.
func goUserAgentVersion() string {
	v := strings.TrimPrefix(runtime.Version(), "go")
	if strings.Count(v, ".") == 1 {
		v += ".0"
	}
	return v
}

// terraformCLIVersion returns the version of the Terraform CLI that the runner uses.
func terraformCLIVersion(t *testing.T) string {
	t.Helper()
	path := cmp.Or(os.Getenv("TF_ACC_TERRAFORM_PATH"), "terraform")
	cli := exec.Command(path, "version", "-json")
	cli.Env = append(os.Environ(), "CHECKPOINT_DISABLE=1")
	output, err := cli.Output()
	require.NoError(t, err)
	var v struct {
		TerraformVersion string `json:"terraform_version"`
	}
	require.NoError(t, json.Unmarshal(output, &v))
	return v.TerraformVersion
}

// assertUserAgent checks that every volume call sends exactly the wanted User-Agent,
// and that the journey sent each method. With sdkGo set, the wanted User-Agent is
// followed by the sdk-go module segments, whose versions follow go.mod.
func assertUserAgent(t *testing.T, requests []volumeRequest, want string, sdkGo bool) {
	t.Helper()
	pattern := regexp.QuoteMeta(want)
	if sdkGo {
		pattern += ` sdk-go-core/\d+\.\d+\.\d+ sdk-go-uc-volumes/\d+\.\d+\.\d+`
	}
	exact := regexp.MustCompile("^" + pattern + "$")
	methods := map[string]bool{}
	for _, r := range requests {
		methods[r.method] = true
		call := r.method + " " + r.path
		values := r.header.Values("User-Agent")
		if assert.Len(t, values, 1, call) {
			assert.Regexp(t, exact, values[0], call)
		}
	}
	all := map[string]bool{"POST": true, "PATCH": true, "GET": true, "DELETE": true}
	assert.Equal(t, all, methods)
}

// assertVolumeClient checks that every volume call came from sdk-go uc/volumes/v1 if
// sdkGo is set, and from the existing client otherwise.
func assertVolumeClient(t *testing.T, requests []volumeRequest, sdkGo bool) {
	t.Helper()
	for _, r := range requests {
		userAgent := r.header.Get("User-Agent")
		assert.Equal(t, sdkGo, strings.Contains(userAgent, " sdk-go-uc-volumes/"),
			"%s %s: %s", r.method, r.path, userAgent)
	}
}

type volumeRequest struct {
	method string
	path   string
	header http.Header
}

type volumeRequestRecorder struct {
	mu       sync.Mutex
	requests []volumeRequest
}

// runVolumeJourney replays the native fixture with the given name and returns its volume
// requests as the server received them. With sdkGo set, the volume is on sdk-go
// uc/volumes/v1 in every step.
func runVolumeJourney(t *testing.T, name string, sdkGo bool) []volumeRequest {
	t.Helper()
	path := "native_testdata/" + name
	recorder := &volumeRequestRecorder{}
	terraformtest.RunNative(t, path,
		func(tc *resource.TestCase, customize func(*config.Config) error) {
			record := func(c *config.Config) error {
				if err := customize(c); err != nil {
					return err
				}
				c.HTTPTransport = volumeRequestTransport{recorder, c.HTTPTransport}
				return nil
			}
			tc.ProtoV6ProviderFactories = terraformtest.ProviderFactories(record)
			if sdkGo {
				tc.ProtoV6ProviderFactories = terraformtest.ProviderFactoriesWithSdkGo(
					record, []string{"databricks_volume"})
			}
		})
	recorder.mu.Lock()
	defer recorder.mu.Unlock()
	return recorder.requests
}

type volumeRequestTransport struct {
	recorder  *volumeRequestRecorder
	transport http.RoundTripper
}

func (v volumeRequestTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if strings.HasPrefix(req.URL.Path, "/api/2.1/unity-catalog/volumes") {
		request := volumeRequest{req.Method, req.URL.Path, req.Header.Clone()}
		v.recorder.mu.Lock()
		v.recorder.requests = append(v.recorder.requests, request)
		v.recorder.mu.Unlock()
	}
	return v.transport.RoundTrip(req)
}
