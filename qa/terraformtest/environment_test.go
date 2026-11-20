package terraformtest

import (
	"os"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/useragent"
	"github.com/stretchr/testify/require"
)

const environmentDefinition = `version: 1
environment:
  CLAUDECODE: "1"
steps:
  - name: create
    config: 'provider "databricks" {}'
    calls: []
destroy:
  calls: []
`

func TestEnvironmentReplay(t *testing.T) {
	Run(t, "testdata/environment")
}

func TestDecodeEnvironmentRejectsInvalidVariables(t *testing.T) {
	for _, tc := range []struct {
		name, variable, message string
	}{
		{"runner Terraform setting", "TF_ACC_TERRAFORM_PATH",
			"environment.TF_ACC_TERRAFORM_PATH: the runner sets this variable"},
		{"runner checkpoint setting", "CHECKPOINT_DISABLE",
			"environment.CHECKPOINT_DISABLE: the runner sets this variable"},
		{"recorder profile", "TEST_RECORDER_PROFILE_NAME",
			"environment.TEST_RECORDER_PROFILE_NAME: the runner sets this variable"},
		{"journey process", journeyProcessEnv,
			"environment." + journeyProcessEnv + ": the runner sets this variable"},
		{"Databricks configuration", "DATABRICKS_TOKEN", "environment.DATABRICKS_TOKEN: " +
			"replay ignores configuration variables; set credentials with auth and other " +
			"attributes in the provider block"},
		{"Azure configuration", "ARM_CLIENT_ID",
			"environment.ARM_CLIENT_ID: replay ignores configuration variables"},
		{"equals sign", `"A=B"`, `environment: invalid variable name "A=B"`},
		{"empty name", `""`, `environment: invalid variable name ""`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			input := strings.Replace(environmentDefinition, "CLAUDECODE", tc.variable, 1)
			_, err := decodeDefinition(strings.NewReader(input))
			require.ErrorContains(t, err, tc.message)
		})
	}
}

// setAmbientEnvironment sets variables a coding agent, CI runner or shell could set.
func setAmbientEnvironment(t *testing.T) {
	t.Helper()
	for name, value := range map[string]string{"CLAUDECODE": "1",
		"GITHUB_ACTIONS": "true", "DATABRICKS_HOST": "https://example.com",
		"DATABRICKS_TF_ENABLED_PF_RESOURCES": "databricks_share",
		"USE_SDK_V2_RESOURCES":               "databricks_share", "ARM_CLIENT_ID": "ambient-sp",
		"GOOGLE_CREDENTIALS": "{}"} {
		t.Setenv(name, value)
	}
	t.Cleanup(useragent.ClearCache)
}

func TestApplyEnvironmentClearsThenSets(t *testing.T) {
	setAmbientEnvironment(t)
	t.Run("journey", func(t *testing.T) {
		require.NoError(t, applyEnvironment(t, map[string]string{"CURSOR_AGENT": "",
			"DATABRICKS_RUNTIME_VERSION": "15.4"}))
		for _, name := range []string{"CLAUDECODE", "GITHUB_ACTIONS", "DATABRICKS_HOST",
			"DATABRICKS_TF_ENABLED_PF_RESOURCES", "USE_SDK_V2_RESOURCES", "ARM_CLIENT_ID",
			"GOOGLE_CREDENTIALS"} {
			_, set := os.LookupEnv(name)
			require.False(t, set, name)
		}
		value, set := os.LookupEnv("CURSOR_AGENT")
		require.True(t, set)
		require.Empty(t, value)
		require.Equal(t, "cursor", useragent.AgentProvider())
		require.Empty(t, useragent.CiCdProvider())
		require.Equal(t, "15.4", useragent.Runtime())
	})
	require.Equal(t, "1", os.Getenv("CLAUDECODE"))
	require.Equal(t, "https://example.com", os.Getenv("DATABRICKS_HOST"))
	_, set := os.LookupEnv("CURSOR_AGENT")
	require.False(t, set)
	// The detections are computed again from the restored variables.
	require.Equal(t, "github", useragent.CiCdProvider())
	require.Empty(t, useragent.Runtime())
}

func TestApplyEnvironmentWithoutEnvironmentKeepsVariables(t *testing.T) {
	setAmbientEnvironment(t)
	t.Run("journey", func(t *testing.T) {
		require.NoError(t, applyEnvironment(t, nil))
		require.Equal(t, "1", os.Getenv("CLAUDECODE"))
	})
	t.Run("empty environment", func(t *testing.T) {
		require.NoError(t, applyEnvironment(t, map[string]string{}))
		_, set := os.LookupEnv("CLAUDECODE")
		require.False(t, set)
	})
}

func TestApplyEnvironmentRejectsProcessStartVariables(t *testing.T) {
	t.Setenv("DATABRICKS_USER_AGENT_EXTRA", "ambient/1.0")
	err := applyEnvironment(t, map[string]string{})
	require.EqualError(t, err, `unset DATABRICKS_USER_AGENT_EXTRA="ambient/1.0": the `+
		"provider read it when the test process started, so every User-Agent contains it")
	// The journey's own process starts with the variable.
	t.Setenv(journeyProcessEnv, t.Name())
	require.NoError(t, applyEnvironment(t, map[string]string{}))
}

func TestUndetectedEnvironmentReportsVariablesLeftSet(t *testing.T) {
	t.Setenv("CLAUDECODE", "1")
	useragent.ClearCache()
	t.Cleanup(useragent.ClearCache)
	require.EqualError(t, undetectedEnvironment(), "the SDK detects agent/claude-code "+
		"from a variable that the runner does not clear; add it to userAgentVariables")
}

func TestRunNativeRejectsProcessStartVariables(t *testing.T) {
	d := &testDefinition{Environment: map[string]string{
		"DATABRICKS_USER_AGENT_EXTRA": "partner/1.0"}}
	require.EqualError(t, d.validateNative(), "environment.DATABRICKS_USER_AGENT_EXTRA: "+
		"the provider reads it when its process starts, which only Run supports")
	d.Environment = map[string]string{"CLAUDECODE": "1"}
	require.NoError(t, d.validateNative())
}

// A failure in the journey's own process fails the journey with that process's output.
func TestOwnProcessReportsFailures(t *testing.T) {
	contents, err := os.ReadFile("testdata/environment/user_agent_extra.yaml")
	require.NoError(t, err)
	fixture := strings.Replace(string(contents), "partner-a/1.2.3 team_b/0.0.1",
		"partner-a/9.9.9 team_b/0.0.1", 1)
	require.NotEqual(t, string(contents), fixture)
	directory := t.TempDir()
	writeDiscoveryFile(t, directory, "user_agent_extra.yaml", fixture)
	output, err := discoverySubprocess(t, directory,
		"^TestDirectoryRunnerHelper$/^user_agent_extra$")
	require.Error(t, err, output)
	require.Contains(t, output, "journey process: exit status 1")
	require.Contains(t, output, "step 1 (create): unexpected HTTP call: "+
		"POST /api/2.0/workspace/mkdirs")
}
