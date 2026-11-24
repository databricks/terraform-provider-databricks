package mlflow_model_test

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

// The four calls that databricks_mlflow_model makes.
const (
	createModelCall = "POST /api/2.0/mlflow/registered-models/create"
	getModelCall    = "GET /api/2.0/mlflow/databricks/registered-models/get"
	updateModelCall = "PATCH /api/2.0/mlflow/registered-models/update"
	deleteModelCall = "DELETE /api/2.0/mlflow/registered-models/delete"
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

// userAgentExtraChild marks the test process that TestMlflowModelUserAgentExtra starts.
const userAgentExtraChild = "DATABRICKS_TF_TEST_USER_AGENT_EXTRA_CHILD"

// sdkGoVariable opts resources in to the modular SDK, and sdkGoResource is the name by
// which it opts databricks_mlflow_model in.
const (
	sdkGoVariable = "USE_SDK_GO_RESOURCES"
	sdkGoResource = "databricks_mlflow_model"
)

// sdkGoUserAgent is what the User-Agent of a registered model call gains on the modular
// SDK path: the modules of the modular SDK that make the call, at their versions in
// go.mod.
const sdkGoUserAgent = " sdk-go-core/0.1.0 sdk-go-modelregistry/0.7.0"

// onBothPaths runs test on the classic SDK path of databricks_mlflow_model and on its
// modular SDK path. The classic run clears the opt-in, so that the environment of the
// test process can't switch the path.
func onBothPaths(t *testing.T, test func(t *testing.T)) {
	t.Run("classic", func(t *testing.T) {
		t.Setenv(sdkGoVariable, "")
		test(t)
	})
	t.Run("sdkgo", func(t *testing.T) {
		t.Setenv(sdkGoVariable, sdkGoResource)
		test(t)
	})
}

// TestMlflowModelUserAgent pins the two User-Agent dimensions that attribute usage of
// the resource: the provider version and the resource name, on each of its four calls.
// YAML header matchers compare exact values, so the dimensions are checked in Go.
func TestMlflowModelUserAgent(t *testing.T) {
	onBothPaths(t, testMlflowModelUserAgent)
}

func testMlflowModelUserAgent(t *testing.T) {
	requests := runModelJourney(t, "native_testdata/user_agent.yaml")
	requireEveryModelCall(t, requests)
	for _, r := range requests {
		userAgents := r.header.Values("User-Agent")
		require.Len(t, userAgents, 1, r.call)
		dimensions := strings.Fields(userAgents[0])
		assert.Contains(t, dimensions, "databricks-tf-provider/"+common.Version(), r.call)
		assert.Contains(t, dimensions, "resource/mlflow_model", r.call)
	}
}

// TestMlflowModelUserAgentDimensions pins the exact User-Agent of every registered
// model call for each environment the classic SDK detects.
func TestMlflowModelUserAgentDimensions(t *testing.T) {
	onBothPaths(t, testMlflowModelUserAgentDimensions)
}

func testMlflowModelUserAgentDimensions(t *testing.T) {
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
			requests := runModelJourney(t, "native_testdata/user_agent.yaml")
			want := wantUserAgent(t, nil, tc.upstream, tc.detected)
			assertUserAgent(t, requests, want)
		})
	}
}

// TestMlflowModelUserAgentExtra pins DATABRICKS_USER_AGENT_EXTRA on every registered
// model call. The provider reads it once at start-up, so the test runs itself in a new
// process.
func TestMlflowModelUserAgentExtra(t *testing.T) {
	onBothPaths(t, testMlflowModelUserAgentExtra)
}

func testMlflowModelUserAgentExtra(t *testing.T) {
	const extra = "partner-a/1.2.3 team_b/0.0.1 partner-a/1.2.3 partner-a/2.0.0 " +
		"resource/mlflow_model"
	if os.Getenv(userAgentExtraChild) == "" {
		requireNoUserAgentExtra(t)
		// The child runs only this subtest, which is the path of its parent.
		child := exec.Command(os.Args[0], "-test.v",
			"-test.run=^"+strings.ReplaceAll(t.Name(), "/", "$/^")+"$")
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
	requests := runModelJourney(t, "native_testdata/user_agent.yaml")
	// Repeated pairs are dropped. A pair repeating a per-request dimension is kept.
	parsed := []string{
		"partner-a/1.2.3", "team_b/0.0.1", "partner-a/2.0.0", "resource/mlflow_model",
	}
	assertUserAgent(t, requests, wantUserAgent(t, parsed, nil, nil))
}

// TestMlflowModelTraceparent pins the W3C traceparent header of every registered model
// call, including both attempts of each retried call.
func TestMlflowModelTraceparent(t *testing.T) {
	onBothPaths(t, testMlflowModelTraceparent)
}

func testMlflowModelTraceparent(t *testing.T) {
	requests := runModelJourney(t, "native_testdata/traceparent_retry.yaml")
	var calls []string
	for _, r := range requests {
		calls = append(calls, r.call)
	}
	require.Equal(t, []string{
		createModelCall, createModelCall, getModelCall, getModelCall, getModelCall,
		getModelCall, updateModelCall, updateModelCall, getModelCall, getModelCall,
		getModelCall, deleteModelCall, deleteModelCall,
	}, calls)
	// Version 00 with the sampled flag set. Every attempt, a retry included, starts a
	// new trace with a new parent ID.
	syntax := regexp.MustCompile(`^00-([0-9a-f]{32})-([0-9a-f]{16})-01$`)
	traceIDs, parentIDs := map[string]bool{}, map[string]bool{}
	for i, r := range requests {
		values := r.header.Values("traceparent")
		require.Len(t, values, 1, "request %d: %s", i, r.call)
		match := syntax.FindStringSubmatch(values[0])
		require.NotNil(t, match, "request %d: %s: %s", i, r.call, values[0])
		assert.NotEqual(t, strings.Repeat("0", 32), match[1], r.call)
		assert.NotEqual(t, strings.Repeat("0", 16), match[2], r.call)
		assert.False(t, traceIDs[match[1]], "%s reuses trace %s", r.call, match[1])
		assert.False(t, parentIDs[match[2]], "%s reuses parent %s", r.call, match[2])
		traceIDs[match[1]], parentIDs[match[2]] = true, true
		assert.Empty(t, r.header.Values("tracestate"), r.call)
	}
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

// wantUserAgent returns the User-Agent of a registered model call with the given extra,
// upstream and detected dimensions, on the path that the test runs on.
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
	dimensions = append(dimensions, "sdk/sdkv2", "resource/mlflow_model", "auth/pat")
	userAgent := strings.Join(append(dimensions, detected...), " ")
	if os.Getenv(sdkGoVariable) == sdkGoResource {
		userAgent += sdkGoUserAgent
	}
	return userAgent
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

// assertUserAgent checks that every registered model call sends exactly the wanted
// User-Agent, and that the journey made each of the four calls.
func assertUserAgent(t *testing.T, requests []modelRequest, want string) {
	t.Helper()
	requireEveryModelCall(t, requests)
	for _, r := range requests {
		assert.Equal(t, []string{want}, r.header.Values("User-Agent"), r.call)
	}
}

// requireEveryModelCall fails unless the journey made each of the four registered model
// calls, so that a check over its requests covers all of them.
func requireEveryModelCall(t *testing.T, requests []modelRequest) {
	t.Helper()
	made := map[string]bool{}
	for _, r := range requests {
		made[r.call] = true
	}
	require.Equal(t, map[string]bool{
		createModelCall: true, getModelCall: true, updateModelCall: true,
		deleteModelCall: true,
	}, made)
}

// modelRequest is one registered model request as the server received it.
type modelRequest struct {
	// call is the method and path, such as createModelCall.
	call   string
	header http.Header
}

type modelRequestRecorder struct {
	mu       sync.Mutex
	requests []modelRequest
}

// runModelJourney replays one native fixture and returns its registered model requests
// in the order the provider sent them.
func runModelJourney(t *testing.T, path string) []modelRequest {
	t.Helper()
	recorder := &modelRequestRecorder{}
	terraformtest.RunNative(t, path,
		func(tc *resource.TestCase, customize func(*config.Config) error) {
			tc.ProtoV6ProviderFactories = terraformtest.ProviderFactories(
				func(c *config.Config) error {
					if err := customize(c); err != nil {
						return err
					}
					c.HTTPTransport = modelRequestTransport{recorder, c.HTTPTransport}
					return nil
				})
		})
	recorder.mu.Lock()
	defer recorder.mu.Unlock()
	return recorder.requests
}

type modelRequestTransport struct {
	recorder  *modelRequestRecorder
	transport http.RoundTripper
}

func (m modelRequestTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if strings.HasPrefix(req.URL.Path, "/api/2.0/mlflow/") {
		request := modelRequest{req.Method + " " + req.URL.Path, req.Header.Clone()}
		m.recorder.mu.Lock()
		m.recorder.requests = append(m.recorder.requests, request)
		m.recorder.mu.Unlock()
	}
	return m.transport.RoundTrip(req)
}
