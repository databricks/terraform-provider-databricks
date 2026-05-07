// Package profile parses ~/.databrickscfg and infers cloud + level for
// testframeworkV2's `requires.cloud` / `requires.level` skip checks.
//
// We deliberately do NOT use the Databricks SDK or
// `internal/acceptance.LoadProfileEnv` here:
//
//   - The SDK validates ALL fields aggressively, which means it errors
//     out on profiles that are perfectly fine for our purposes
//     (DESIGN.md §10/G9).
//   - `internal/acceptance.init()` calls os.Setenv("TF_LOG", "DEBUG")
//     which leaks into every subsequent terraform invocation
//     (DESIGN.md §10/G7 / B6 / F).
//
// We do JUST enough parsing to (a) confirm the named section exists in
// the file, and (b) infer Cloud + Level from the host string. Everything
// else — token, account_id, google_service_account, etc. — is left to the
// SDK at terraform-run time, which is the only place that needs to
// validate cloud-specific fields.
package profile

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Cloud identifies the Databricks cloud a profile points at. CloudUnknown
// signals "could not infer" — the runner treats it as a non-match for any
// requires.cloud value other than `any`.
type Cloud string

const (
	CloudAWS     Cloud = "aws"
	CloudAzure   Cloud = "azure"
	CloudGCP     Cloud = "gcp"
	CloudUnknown Cloud = ""
)

// Level identifies whether a profile is workspace-level or account-level.
type Level string

const (
	LevelWorkspace Level = "workspace"
	LevelAccount   Level = "account"
)

// Profile is the parsed view of a single ~/.databrickscfg section. Raw
// retains every key/value pair the section declared, for debugging.
type Profile struct {
	// Name is the section name, e.g. "ACCOUNT_AWS".
	Name string

	// Host is the value of the `host` key, verbatim.
	Host string

	// AccountID is the value of `account_id`, or "" if unset.
	AccountID string

	// AuthType is the value of `auth_type`, or "" if unset. The runner
	// does not interpret this — it's surfaced for debugging only.
	AuthType string

	// Cloud is inferred from Host (DESIGN.md §10 G9). CloudUnknown when
	// the host shape doesn't match any known pattern.
	Cloud Cloud

	// Level is inferred from Host prefix and AccountID presence. A host
	// starting with "accounts." OR a non-empty account_id field implies
	// LevelAccount; otherwise LevelWorkspace.
	Level Level

	// Raw is the verbatim key/value map for the section.
	Raw map[string]string
}

// DefaultPath returns the conventional ~/.databrickscfg path. Honors
// DATABRICKS_CONFIG_FILE if set in the parent env. The runner uses this
// for both Load() and for setting DATABRICKS_CONFIG_FILE on the terraform
// subprocess (subprocenv.Build).
func DefaultPath() string {
	if p := strings.TrimSpace(os.Getenv("DATABRICKS_CONFIG_FILE")); p != "" {
		return p
	}
	home, err := os.UserHomeDir()
	if err != nil || home == "" {
		// Best-effort fallback: $HOME isn't set on some CI runners, but
		// the runner usually checks the file exists later anyway.
		return filepath.Join(os.Getenv("HOME"), ".databrickscfg")
	}
	return filepath.Join(home, ".databrickscfg")
}

// ErrSectionNotFound is returned by Load / LoadFromPath when the named
// profile section does not exist in the file.
var ErrSectionNotFound = errors.New("profile section not found in ~/.databrickscfg")

// Load reads DefaultPath() and returns the named profile. Wraps
// LoadFromPath for the common case where DATABRICKS_CONFIG_FILE governs.
func Load(name string) (*Profile, error) {
	return LoadFromPath(DefaultPath(), name)
}

// LoadFromPath reads the .databrickscfg-format file at path and returns
// the section named `name`. Returns ErrSectionNotFound when the section
// is missing — the runner uses errors.Is to differentiate from I/O errors
// when surfacing a friendly preflight message.
func LoadFromPath(path, name string) (*Profile, error) {
	if name == "" {
		return nil, errors.New("profile: name is required")
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("profile: open %s: %w", path, err)
	}
	defer f.Close()

	sections, err := parseINI(f)
	if err != nil {
		return nil, fmt.Errorf("profile: parse %s: %w", path, err)
	}
	raw, ok := sections[name]
	if !ok {
		return nil, fmt.Errorf("%w: %q in %s", ErrSectionNotFound, name, path)
	}
	return profileFromRaw(name, raw), nil
}

// profileFromRaw assembles a Profile from a parsed section's key/value
// map. Cloud/Level inference happens here so callers always get a
// fully-populated value.
func profileFromRaw(name string, raw map[string]string) *Profile {
	host := raw["host"]
	accountID := raw["account_id"]
	cloud, level := inferCloudLevel(host, accountID)
	return &Profile{
		Name:      name,
		Host:      host,
		AccountID: accountID,
		AuthType:  raw["auth_type"],
		Cloud:     cloud,
		Level:     level,
		Raw:       raw,
	}
}

// SectionExists is a thin convenience wrapper for the config-layer
// preflight check (DESIGN.md §4 "Profile existence"). It does NOT load
// the section — useful for confirming a profile name is well-formed
// before the runner spends time on cache lookups and run-dir setup.
func SectionExists(path, name string) (bool, error) {
	if name == "" {
		return false, errors.New("profile: name is required")
	}
	f, err := os.Open(path)
	if err != nil {
		return false, fmt.Errorf("profile: open %s: %w", path, err)
	}
	defer f.Close()
	sections, err := parseINI(f)
	if err != nil {
		return false, fmt.Errorf("profile: parse %s: %w", path, err)
	}
	_, ok := sections[name]
	return ok, nil
}

// parseINI is a deliberately tiny parser for the .databrickscfg subset of
// INI: section headers `[Name]`, `key = value` lines, comments starting
// with `#` or `;`, and blank lines. We don't pull in `gopkg.in/ini.v1`
// because (a) the format we accept is trivial, and (b) one fewer dep is
// one fewer attack surface for a separate go.mod whose whole point is
// minimalism.
func parseINI(r io.Reader) (map[string]map[string]string, error) {
	sections := map[string]map[string]string{}
	current := ""
	scanner := bufio.NewScanner(r)
	// Allow generous line lengths — Databricks PATs and JWT tokens can be
	// long, and 1 MiB is plenty for any reasonable .databrickscfg.
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	lineno := 0
	for scanner.Scan() {
		lineno++
		if err := parseINILine(scanner.Text(), lineno, sections, &current); err != nil {
			return nil, err
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan: %w", err)
	}
	return sections, nil
}

// parseINILine handles a single line of the .databrickscfg-format file.
// Mutates sections and *current. Lines that are blank or comments are
// no-ops; section headers update *current; key=value lines populate the
// active section.
func parseINILine(raw string, lineno int, sections map[string]map[string]string, current *string) error {
	line := strings.TrimSpace(raw)
	if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
		return nil
	}
	if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
		name := strings.TrimSpace(line[1 : len(line)-1])
		*current = name
		if _, ok := sections[name]; !ok {
			sections[name] = map[string]string{}
		}
		return nil
	}
	if *current == "" {
		// Implicit DEFAULT — Databricks CLI tolerates pre-section keys.
		*current = "DEFAULT"
		if _, ok := sections[*current]; !ok {
			sections[*current] = map[string]string{}
		}
	}
	k, v, ok := strings.Cut(line, "=")
	if !ok {
		return fmt.Errorf("line %d: expected `key = value`, got %q", lineno, raw)
	}
	k = strings.TrimSpace(k)
	if k == "" {
		return fmt.Errorf("line %d: empty key", lineno)
	}
	sections[*current][k] = strings.TrimSpace(v)
	return nil
}

// inferCloudLevel maps a host (and optional account_id) to (Cloud, Level)
// using the rules from DESIGN.md §10 G9.
//
// Cloud inference is keyed on substrings of the host because customer
// accounts can have arbitrary subdomain prefixes (e.g.
// "dbc-abcd1234.cloud.databricks.com" is a workspace; "accounts.cloud.
// databricks.com" is the AWS account control plane).
//
// Level inference: a host starting with "accounts." is the account
// control plane; or, presence of an `account_id` field indicates
// account-level (workspace profiles never set account_id).
func inferCloudLevel(host, accountID string) (Cloud, Level) {
	h := normalizeHost(host)

	level := LevelWorkspace
	if strings.HasPrefix(h, "accounts.") || strings.TrimSpace(accountID) != "" {
		level = LevelAccount
	}

	switch {
	case strings.Contains(h, "azuredatabricks.net"):
		return CloudAzure, level
	case strings.Contains(h, ".gcp.databricks.com"):
		return CloudGCP, level
	case strings.Contains(h, ".cloud.databricks.com"):
		return CloudAWS, level
	default:
		return CloudUnknown, level
	}
}

// normalizeHost returns the host portion of a Databricks URL — strips
// scheme, port, and path; lowercases the result. Matching is done
// substring-wise downstream, so bullet-proof URL parsing is overkill.
func normalizeHost(host string) string {
	h := strings.ToLower(strings.TrimSpace(host))
	h = strings.TrimPrefix(h, "https://")
	h = strings.TrimPrefix(h, "http://")
	if i := strings.IndexByte(h, '/'); i >= 0 {
		h = h[:i]
	}
	if i := strings.IndexByte(h, ':'); i >= 0 {
		h = h[:i]
	}
	return h
}
