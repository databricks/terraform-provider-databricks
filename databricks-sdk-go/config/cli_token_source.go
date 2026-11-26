package config

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/logger"
	"golang.org/x/mod/semver"
	"golang.org/x/oauth2"
)

// databricksCliMinSize is the size at or above which a file is assumed to be
// the Go CLI without reading it.
const databricksCliMinSize = 1024 * 1024

// legacyCliMarker identifies the legacy CLI's pip entry point script. The
// imported name is omitted because it changed from `cli` to `main` in 0.17.8.
var legacyCliMarker = []byte("from databricks_cli.cli import ")

var (
	ErrCliNotFound       = errors.New("databricks CLI not found")
	ErrLegacyCliDetected = errors.New("legacy databricks CLI detected; upgrade to >= 0.100.0")
)

type cliTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Expiry      string `json:"expiry"`
}

// Minimum CLI versions for flag support. Versions are in the format accepted
// by [golang.org/x/mod/semver] (with leading "v").
const (
	cliVersionForProfile      = "v0.207.1" // https://github.com/databricks/cli/pull/855
	cliVersionForForceRefresh = "v0.296.0" // https://github.com/databricks/cli/pull/4767
)

// devBuildVersionPrefix is the sentinel the Databricks CLI uses when
// buildVersion wasn't injected by goreleaser (e.g. a plain `go build`).
// See https://github.com/databricks/cli/blob/main/internal/build/info.go.
const devBuildVersionPrefix = "v0.0.0-dev"

// getCliVersion runs "databricks version" and returns the parsed semver string
// (e.g., "v0.207.1"). Returns an empty string and an error if the version
// cannot be determined.
func getCliVersion(ctx context.Context, cliPath string) (string, error) {
	cmd := exec.CommandContext(ctx, cliPath, "version")
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("cannot get CLI version: %w", err)
	}
	return parseCliVersion(strings.TrimSpace(string(out)))
}

// parseCliVersion extracts the semver from an output like "Databricks CLI v0.207.1".
func parseCliVersion(s string) (string, error) {
	v := strings.TrimPrefix(s, "Databricks CLI ")
	if !semver.IsValid(v) {
		return "", fmt.Errorf("cannot parse CLI version %q", s)
	}
	return v, nil
}

// displayVersion returns a human-readable version string, or "unknown" when
// version detection failed (indicated by an empty string).
func displayVersion(ver string) string {
	if ver == "" {
		return "unknown"
	}
	return ver
}

// CliTokenSource fetches OAuth tokens by shelling out to the Databricks CLI.
type CliTokenSource struct {
	cmd []string
}

// NewCliTokenSource creates a [CliTokenSource] by detecting the installed CLI
// version and building the appropriate auth token command.
func NewCliTokenSource(ctx context.Context, cfg *Config) (*CliTokenSource, error) {
	cliPath, err := findDatabricksCli(cfg.DatabricksCliPath)
	if err != nil {
		return nil, err
	}
	cmd, err := resolveCliCommand(ctx, cliPath, cfg)
	if err != nil {
		return nil, err
	}
	return &CliTokenSource{cmd: cmd}, nil
}

// resolveCliCommand detects the CLI version and builds the command to execute.
func resolveCliCommand(ctx context.Context, cliPath string, cfg *Config) ([]string, error) {
	ver, err := getCliVersion(ctx, cliPath)
	if err != nil {
		logger.Warnf(ctx, "Failed to detect Databricks CLI version: %v. Falling back to conservative flag set.", err)
		ver = ""
	} else if strings.HasPrefix(ver, devBuildVersionPrefix) {
		// Dev builds always compare less than any real release, which would
		// disable every feature gate; surface an informational hint instead.
		logger.Infof(ctx,
			"Databricks CLI %s is a development build; feature detection will use conservative fallbacks. "+
				"Rebuild the CLI with an explicit version to enable capability-based flag selection, "+
				"e.g. `go build -ldflags '-X github.com/databricks/cli/internal/build.buildVersion=0.296.0'`.",
			ver)
	}
	cmd, err := buildCliCommand(ctx, cliPath, cfg, ver)
	if err != nil {
		return nil, fmt.Errorf("cannot configure CLI token source: %w", err)
	}
	return cmd, nil
}

// buildCliCommand constructs the full CLI command for fetching an auth
// token, including capability-gated flags. Returns a precise error
// describing which piece of configuration is missing when the command
// cannot be built.
func buildCliCommand(ctx context.Context, cliPath string, cfg *Config, ver string) ([]string, error) {
	cmd, err := buildCoreCliCommand(ctx, cliPath, cfg, ver)
	if err != nil {
		return nil, err
	}
	if semver.Compare(ver, cliVersionForForceRefresh) >= 0 {
		cmd = append(cmd, "--force-refresh")
	} else {
		logger.Warnf(ctx, "Databricks CLI %s does not support --force-refresh (requires >= %s). The CLI's token cache may provide stale tokens.",
			displayVersion(ver), cliVersionForForceRefresh)
	}
	return cmd, nil
}

// buildCoreCliCommand constructs the base `auth token` command without any
// capability-gated flags. Falls back to --host when --profile is either not
// configured or not supported by the installed CLI. Returns a precise error
// describing which piece of configuration is missing when the command
// cannot be built.
func buildCoreCliCommand(ctx context.Context, cliPath string, cfg *Config, ver string) ([]string, error) {
	if cfg.Profile == "" {
		cmd, err := buildHostCommand(cliPath, cfg)
		if err != nil {
			return nil, fmt.Errorf("neither profile nor host is configured: %w", err)
		}

		return cmd, nil
	}

	// Flag --profile is a global CLI flag and is recognized for all commands
	// even the ones that do not support it. Only use flag --profile in CLI
	// versions that are known to support it in command `auth token`.
	if semver.Compare(ver, cliVersionForProfile) < 0 {
		logger.Warnf(ctx, "Databricks CLI %s does not support --profile (requires >= %s). Falling back to --host.", displayVersion(ver), cliVersionForProfile)

		cmd, err := buildHostCommand(cliPath, cfg)
		if err != nil {
			return nil, fmt.Errorf("databricks CLI %s does not support --profile (requires >= %s) and no host fallback is configured: %w",
				displayVersion(ver), cliVersionForProfile, err)
		}

		return cmd, nil
	}

	return []string{cliPath, "auth", "token", "--profile", cfg.Profile}, nil
}

// buildHostCommand constructs the --host based auth token command.
// Returns an error when cfg.Host is unset.
func buildHostCommand(cliPath string, cfg *Config) ([]string, error) {
	if cfg.Host == "" {
		return nil, errors.New("host is not set")
	}
	cmd := []string{cliPath, "auth", "token", "--host", cfg.Host}
	if cfg.HostType() == AccountHost {
		cmd = append(cmd, "--account-id", cfg.AccountID)
	}
	return cmd, nil
}

// Token fetches an OAuth token by shelling out to the Databricks CLI.
func (c *CliTokenSource) Token(ctx context.Context) (*oauth2.Token, error) {
	return c.execCliCommand(ctx, c.cmd)
}

func (c *CliTokenSource) execCliCommand(ctx context.Context, args []string) (*oauth2.Token, error) {
	cmd := exec.CommandContext(ctx, args[0], args[1:]...)
	stdout, err := cmd.Output()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			// %s prints exitErr.Stderr (CLI's message). %w prints err.Error(),
			// which for *exec.ExitError is just "exit status N" — complementary,
			// not duplicated. %w also preserves the type for errors.As.
			return nil, fmt.Errorf("cannot get access token: %s: %w", strings.TrimSpace(string(exitErr.Stderr)), err)
		}
		return nil, fmt.Errorf("cannot get access token: %w", err)
	}
	var resp cliTokenResponse
	if err := json.Unmarshal(stdout, &resp); err != nil {
		return nil, fmt.Errorf("cannot parse CLI response: %w", err)
	}
	expiry, err := parseExpiry(resp.Expiry)
	if err != nil {
		return nil, fmt.Errorf("cannot parse token expiry: %w", err)
	}
	return &oauth2.Token{
		AccessToken: resp.AccessToken,
		TokenType:   resp.TokenType,
		Expiry:      expiry,
	}, nil
}

// parseExpiry parses an expiry time string in multiple formats for cross-SDK compatibility.
//
// Supported formats:
//  1. RFC 3339 with nanoseconds (e.g., "2024-03-20T10:30:00.123456789Z")
//  2. RFC 3339 (e.g., "2024-03-20T10:30:00Z" or "2024-03-20T10:30:00+01:00")
//  3. Space-separated without fractional seconds (e.g., "2024-03-20 10:30:00")
//  4. Space-separated with fractional seconds (e.g., "2024-03-20 10:30:00.123")
//
// The Databricks CLI outputs time.RFC3339Nano format.
//
// Formats 3 and 4 are supported for parity with the Python and Java SDKs, which accept
// space-separated datetime formats and assume the local timezone.
func parseExpiry(expiry string) (time.Time, error) {
	if t, err := time.Parse(time.RFC3339Nano, expiry); err == nil {
		return t, nil
	}
	if t, err := time.Parse(time.RFC3339, expiry); err == nil {
		return t, nil
	}

	if t, err := time.Parse("2006-01-02 15:04:05", expiry); err == nil {
		return t, nil
	}

	layouts := []string{
		"2006-01-02 15:04:05.9",
		"2006-01-02 15:04:05.99",
		"2006-01-02 15:04:05.999",
		"2006-01-02 15:04:05.9999",
		"2006-01-02 15:04:05.99999",
		"2006-01-02 15:04:05.999999",
		"2006-01-02 15:04:05.9999999",
		"2006-01-02 15:04:05.99999999",
		"2006-01-02 15:04:05.999999999",
	}
	for _, layout := range layouts {
		if t, err := time.Parse(layout, expiry); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("cannot parse expiry %q", expiry)
}

func findDatabricksCli(cliPath string) (string, error) {
	if cliPath == "" {
		path, err := findInPath("databricks")
		if err != nil && runtime.GOOS == "windows" {
			return findInPath("databricks.exe")
		}
		return path, err
	}
	if strings.Contains(cliPath, string(filepath.Separator)) || strings.Contains(cliPath, "/") {
		// An explicit path is trusted without the legacy CLI check, so that a
		// wrapper the check would reject can still be used.
		if _, err := statCliPath(cliPath); err != nil {
			return "", err
		}
		return cliPath, nil
	}
	return findInPath(cliPath)
}

func findInPath(name string) (string, error) {
	pathEnv := os.Getenv("PATH")
	if pathEnv == "" {
		return "", ErrCliNotFound
	}
	lastErr := ErrCliNotFound
	for _, dir := range filepath.SplitList(pathEnv) {
		validPath, err := validateCliPath(filepath.Join(dir, name))
		if err == nil {
			return validPath, nil
		}
		if errors.Is(err, ErrLegacyCliDetected) {
			lastErr = err
		}
	}
	return "", lastErr
}

// statCliPath returns the file info of path if it is a regular file.
func statCliPath(path string) (os.FileInfo, error) {
	info, err := os.Stat(path)
	// Only regular files are accepted; reading a FIFO, for example, can block.
	if os.IsNotExist(err) || (err == nil && !info.Mode().IsRegular()) {
		return nil, ErrCliNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("cannot stat CLI path: %w", err)
	}
	return info, nil
}

func validateCliPath(path string) (string, error) {
	info, err := statCliPath(path)
	if err != nil {
		return "", err
	}
	if info.Size() >= databricksCliMinSize {
		return path, nil
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	if bytes.Contains(b, legacyCliMarker) {
		return "", ErrLegacyCliDetected
	}
	return path, nil
}
