package terraformtest

import (
	"errors"
	"fmt"
	"maps"
	"os"
	"os/exec"
	"slices"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// userAgentVariables are the variables besides DATABRICKS_* from which the SDK
// derives User-Agent dimensions. The SDK keeps its lists private, so
// applyEnvironment checks that nothing is detected once these are cleared.
var userAgentVariables = []string{
	// Coding agents.
	"AMP_CURRENT_THREAD_ID", "ANTIGRAVITY_AGENT", "AUGMENT_AGENT", "CLAUDECODE",
	"CLINE_ACTIVE", "CODEX_CI", "COPILOT_CLI", "CURSOR_AGENT", "GEMINI_CLI",
	"GOOSE_TERMINAL", "KIRO", "OPENCLAW_SHELL", "OPENCODE", "PI_CODING_AGENT",
	"VSCODE_AGENT", "WINDSURF_AGENT", "AGENT", "AI_AGENT",
	// Agent meta-harnesses.
	"OMNIGENT",
	// CI/CD providers.
	"GITHUB_ACTIONS", "GITLAB_CI", "JENKINS_URL", "TF_BUILD", "CIRCLECI", "TRAVIS",
	"BITBUCKET_BUILD_NUMBER", "PROJECT_ID", "BUILD_ID", "PROJECT_NUMBER", "LOCATION",
	"CODEBUILD_BUILD_ARN", "TFC_RUN_ID",
}

// rolloutVariables choose the provider's resource implementations. The
// DATABRICKS_TF_* rollout variables are cleared with the other DATABRICKS_* ones.
var rolloutVariables = []string{"USE_SDK_V2_RESOURCES", "USE_SDK_V2_DATA_SOURCES"}

// processStartVariables are read once, when the provider package initializes, so
// t.Setenv cannot change them. A journey that sets one runs in its own process.
var processStartVariables = []string{"DATABRICKS_USER_AGENT_EXTRA"}

// journeyProcessEnv names the journey that a test process was started to run.
const journeyProcessEnv = "TERRAFORMTEST_JOURNEY_PROCESS"

// configurationVariables are the variables the SDK reads configuration from.
// Replay ignores them, because the fixture loader replaces the SDK's loaders.
func configurationVariables() map[string]bool {
	names := map[string]bool{}
	for _, attr := range config.ConfigAttributes {
		for _, name := range attr.EnvVars {
			names[name] = true
		}
	}
	return names
}

func validateEnvironment(env map[string]string) error {
	configuration := configurationVariables()
	for _, name := range slices.Sorted(maps.Keys(env)) {
		switch {
		case name == "" || strings.ContainsAny(name, "=\x00"):
			return fmt.Errorf("environment: invalid variable name %q", name)
		case strings.ContainsRune(env[name], 0):
			return fmt.Errorf("environment.%s: value must not contain NUL", name)
		case strings.HasPrefix(name, "TF_ACC_") || slices.Contains([]string{
			"CHECKPOINT_DISABLE", "TF_IN_AUTOMATION", recorderProfileEnv,
			journeyProcessEnv}, name):
			return fmt.Errorf("environment.%s: the runner sets this variable", name)
		case configuration[name]:
			return fmt.Errorf("environment.%s: replay ignores configuration variables; "+
				"set credentials with auth and other attributes in the provider block", name)
		}
	}
	return nil
}

// clearedVariables are the variables that change the User-Agent or the client's
// configuration, which applyEnvironment clears before setting a journey's.
func clearedVariables() []string {
	names := slices.Concat(userAgentVariables, rolloutVariables,
		slices.Collect(maps.Keys(configurationVariables())))
	for _, entry := range os.Environ() {
		if name, _, _ := strings.Cut(entry, "="); strings.HasPrefix(name, "DATABRICKS_") {
			names = append(names, name)
		}
	}
	return names
}

// applyEnvironment gives a journey with an environment the same variables under any
// coding agent, CI runner or shell: it clears the variables that change the
// User-Agent or the configuration, then sets the journey's for the test.
func applyEnvironment(t *testing.T, env map[string]string) error {
	t.Helper()
	if env == nil {
		return nil
	}
	if !inOwnProcess(t) {
		for _, name := range processStartVariables {
			if value := os.Getenv(name); value != "" {
				return fmt.Errorf("unset %s=%q: the provider read it when the test "+
					"process started, so every User-Agent contains it", name, value)
			}
		}
	}
	// The SDK caches each detection for the life of the process. Registered first,
	// this cleanup runs after every variable is restored.
	t.Cleanup(useragent.ClearCache)
	for _, name := range clearedVariables() {
		t.Setenv(name, "")
		if err := os.Unsetenv(name); err != nil {
			return err
		}
	}
	useragent.ClearCache()
	if err := undetectedEnvironment(); err != nil {
		return err
	}
	for name, value := range env {
		t.Setenv(name, value)
	}
	useragent.ClearCache()
	return nil
}

// undetectedEnvironment fails if the SDK still detects a User-Agent dimension, as
// it would from a variable the SDK added after userAgentVariables was written.
func undetectedEnvironment() error {
	for _, detected := range [][2]string{
		{"agent", useragent.AgentProvider()},
		{"meta-harness", useragent.MetaHarnessProvider()},
		{"cicd", useragent.CiCdProvider()},
		{"runtime", useragent.Runtime()},
	} {
		if detected[1] != "" {
			return fmt.Errorf("the SDK detects %s/%s from a variable that the runner "+
				"does not clear; add it to userAgentVariables", detected[0], detected[1])
		}
	}
	return nil
}

// processStartVariable returns a variable the journey sets that the provider reads
// only when its process starts, or an empty string.
func (d *testDefinition) processStartVariable() string {
	for _, name := range processStartVariables {
		if _, ok := d.Environment[name]; ok {
			return name
		}
	}
	return ""
}

// inOwnProcess reports whether this test process was started to run t's journey.
func inOwnProcess(t *testing.T) bool {
	return os.Getenv(journeyProcessEnv) == t.Name()
}

// validateNative rejects process-start variables in RunNative, whose Go checks run
// in the calling process, so the journey cannot start its own.
func (d *testDefinition) validateNative() error {
	if name := d.processStartVariable(); name != "" {
		return fmt.Errorf("environment.%s: the provider reads it when its process "+
			"starts, which only Run supports", name)
	}
	return nil
}

// runOwnProcess runs the journey t names in a new test process that starts with the
// journey's environment, as Terraform starts the provider in a new process.
func runOwnProcess(t *testing.T, env map[string]string) {
	t.Helper()
	executable, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	args := []string{"-test.v", "-test.run=" + exactRunPattern(t.Name())}
	if testing.Short() {
		args = append(args, "-test.short")
	}
	cmd := exec.CommandContext(t.Context(), executable, args...)
	cleared := clearedVariables()
	for _, entry := range os.Environ() {
		if name, _, _ := strings.Cut(entry, "="); !slices.Contains(cleared, name) {
			cmd.Env = append(cmd.Env, entry)
		}
	}
	cmd.Env = append(cmd.Env, journeyProcessEnv+"="+t.Name())
	for name, value := range env {
		cmd.Env = append(cmd.Env, name+"="+value)
	}
	output, err := cmd.CombinedOutput()
	if err == nil && !strings.Contains(string(output), "--- PASS: "+t.Name()+" ") {
		err = errors.New("the journey did not run")
	}
	if err != nil {
		t.Fatalf("journey process: %v\n%s", err, output)
	}
}
