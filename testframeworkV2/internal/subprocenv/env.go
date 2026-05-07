// Package subprocenv builds the environment variable list passed to every
// `terraform` subprocess in a testframeworkV2 run.
//
// The list is a strict allowlist (DESIGN.md §10/G6 / B5). It exists because
// `os.Environ()` can leak DATABRICKS_HOST or DATABRICKS_TOKEN from the
// developer's shell into the SDK auth chain, which silently routes the test
// against the wrong host. The runner is expected to call this function and
// then `tfexec.Terraform.SetEnv(...)` — `SetEnv` REPLACES the subprocess env
// entirely (it does not inherit os.Environ()), which is the linchpin of
// the leak protection.
package subprocenv

import (
	"os"
	"path/filepath"
	"strings"
)

// allowedKeys is the curated list of parent-environment variables that
// pass through to the terraform subprocess as-is. Order is grouped by
// rationale (basics, locale, TLS, proxy) for human readability.
//
// IMPORTANT: this list is exhaustive — anything not here is stripped.
// Adding to it widens the attack surface for env-var leaks, so keep
// additions justified in DESIGN.md.
var allowedKeys = []string{
	// Basics — required for any subprocess to execute.
	"PATH",
	"HOME",
	"USER",
	"TMPDIR", // CI runners often have custom temp dirs.

	// Locale — terraform formats dates and parses strings using locale.
	"LANG",
	"LC_ALL",
	"LC_CTYPE",

	// Corporate CA bundle — needed for HTTPS on locked-down dev / CI.
	"SSL_CERT_FILE",
	"SSL_CERT_DIR",

	// Corporate proxy — terraform init contacts the registry; the
	// framework also fetches GitHub releases.
	"HTTPS_PROXY",
	"HTTP_PROXY",
	"NO_PROXY",
}

// Build returns the env slice the runner passes to tfexec.Terraform.SetEnv.
//
// Inputs:
//
//   - profile        Databricks profile name (becomes
//     DATABRICKS_CONFIG_PROFILE; the SDK's auth chain reads it).
//   - tfrcPath       absolute path to the per-run .terraformrc the runner
//     wrote (becomes TF_CLI_CONFIG_FILE; this is the
//     mechanism that makes terraform ignore the user's
//     ~/.terraformrc — DESIGN.md §5 / G1).
//   - runDir         per-run root (TF_PLUGIN_CACHE_DIR is set to
//     <runDir>/plugins, enabling terraform's hardlink
//     optimization within the run — F2).
//   - passthrough    optional list of env-var names from test.yaml's
//     passthrough_env field. Names whose value is empty
//     in the parent environment are silently dropped.
//     DATABRICKS_* names are explicitly REJECTED here —
//     the profile mechanism is the only sanctioned auth
//     channel (DESIGN.md §4 / §10 G6).
//
// The DATABRICKS_CONFIG_FILE override always points at $HOME/.databrickscfg
// so the SDK's auth chain reads from a deterministic location regardless of
// the developer's shell config.
func Build(profile, tfrcPath, runDir string, passthrough []string) []string {
	env := make([]string, 0, len(allowedKeys)+5+len(passthrough))
	env = appendAllowlist(env)
	env = appendFrameworkControlled(env, profile, tfrcPath, runDir)
	env = appendPassthrough(env, passthrough)
	return env
}

// appendAllowlist copies allowed parent-env entries (PATH, HOME, LANG,
// proxy, etc.). Empty values are dropped — exec inherits them as empty
// strings which is wasteful at best and (for LC_ALL etc.) misleading.
func appendAllowlist(env []string) []string {
	for _, k := range allowedKeys {
		if v, ok := os.LookupEnv(k); ok && v != "" {
			env = append(env, k+"="+v)
		}
	}
	return env
}

// appendFrameworkControlled writes the variables the framework owns:
// TF_CLI_CONFIG_FILE / TF_PLUGIN_CACHE_DIR — and the
// DATABRICKS_CONFIG_PROFILE / DATABRICKS_CONFIG_FILE pair which is the
// linchpin of the no-leak invariant (DESIGN.md §10 G6).
//
// Note: TF_IN_AUTOMATION is intentionally NOT added here. tfexec
// auto-manages it (terraform-exec@v0.25.0/tfexec/cmd.go:178
// unconditionally sets env[automationEnvVar] = "1" in buildEnv) and
// rejects callers who try to set it via tfexec.Terraform.SetEnv —
// `manual setting of env var "TF_IN_AUTOMATION" detected`. The
// subprocess still receives it because tfexec sets it; we just must
// not duplicate. The same prohibition applies to TF_CLI_ARGS,
// TF_CLI_ARGS_*, TF_INPUT, TF_LOG, TF_LOG_CORE, TF_LOG_PATH,
// TF_LOG_PROVIDER, TF_REATTACH_PROVIDERS, TF_APPEND_USER_AGENT,
// TF_WORKSPACE, TF_DISABLE_PLUGIN_TLS, TF_SKIP_PROVIDER_VERIFY, and
// TF_VAR_* — none of which we set; the regression tests in
// env_test.go enumerate all of them so a future contributor can't
// silently re-introduce one.
func appendFrameworkControlled(env []string, profile, tfrcPath, runDir string) []string {
	return append(env,
		"TF_CLI_CONFIG_FILE="+tfrcPath,
		"TF_PLUGIN_CACHE_DIR="+filepath.Join(runDir, "plugins"),
		"DATABRICKS_CONFIG_PROFILE="+profile,
		"DATABRICKS_CONFIG_FILE="+filepath.Join(os.Getenv("HOME"), ".databrickscfg"),
	)
}

// appendPassthrough adds the names from test.yaml's passthrough_env list
// when they have a non-empty value in the parent env. DATABRICKS_*
// names are silently dropped here — the config layer rejects them at
// parse time, but this is defense-in-depth.
func appendPassthrough(env, passthrough []string) []string {
	for _, name := range passthrough {
		if name == "" || strings.HasPrefix(name, "DATABRICKS_") {
			continue
		}
		if v, ok := os.LookupEnv(name); ok && v != "" {
			env = append(env, name+"="+v)
		}
	}
	return env
}

// AllowedKeys returns a copy of the curated parent-env allowlist. Useful
// for tests and for the future `tfv2 env` debug subcommand.
func AllowedKeys() []string {
	out := make([]string, len(allowedKeys))
	copy(out, allowedKeys)
	return out
}
