// Package config parses and validates testframeworkV2's test.yaml schema
// (DESIGN.md §4). The package owns:
//
//   - the public Go types (TestSpec, Step, Requires) that the runner
//     consumes,
//   - YAML parsing with strict-unknown-field rejection so typos surface
//     loud rather than silently turning into defaults,
//   - schema validation: required fields, slug shape, version syntax,
//     `expect: failure` ⇒ ≥1 of error_substring/error_regex, regex
//     compilation, enum membership.
//
// The package does NOT make any cloud calls. Profile-existence checks
// happen in the config layer because they're a parse-time concern, but
// the actual file IO is delegated to internal/profile.
package config

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/profile"
)

// LocalVersion is the literal version string a step uses to opt into a
// local build (DESIGN.md §4 / §8). Lifted into config so the runner can
// dispatch to providercache.Resolve("local", ...) without re-checking
// against the constant elsewhere.
const LocalVersion = "local"

// DefaultStepTimeout is the per-step timeout used when test.yaml omits
// the `timeout` field (DESIGN.md §4 schema rules).
const DefaultStepTimeout = 10 * time.Minute

// DefaultCleanup is the default value of the top-level `cleanup` flag.
const DefaultCleanup = true

// Cloud and Level are wire-friendly enum types lifted into the config
// schema. Their values intentionally match internal/profile's enum
// values so the runner can do a direct comparison
// (config.Cloud == profile.Cloud).
type Cloud string

const (
	CloudAWS   Cloud = "aws"
	CloudAzure Cloud = "azure"
	CloudGCP   Cloud = "gcp"
	CloudAny   Cloud = "any"
)

type Level string

const (
	LevelWorkspace Level = "workspace"
	LevelAccount   Level = "account"
	LevelUCWS      Level = "ucws"
	LevelUCAcct    Level = "ucacct"
)

// Command is the terraform subcommand a step invokes.
type Command string

const (
	CommandPlan    Command = "plan"
	CommandApply   Command = "apply"
	CommandDestroy Command = "destroy"
)

// Expect is whether a step should pass or fail.
type Expect string

const (
	ExpectSuccess Expect = "success"
	ExpectFailure Expect = "failure"
)

// Requires gates whether a test runs against a given host.
type Requires struct {
	Cloud Cloud `yaml:"cloud"`
	Level Level `yaml:"level"`
}

// Step is one atomic unit of work in a test.yaml.
type Step struct {
	Name           string         `yaml:"name"`
	Version        string         `yaml:"version"`
	Command        Command        `yaml:"command"`
	Expect         Expect         `yaml:"expect"`
	ErrorSubstring string         `yaml:"error_substring"`
	ErrorRegex     string         `yaml:"error_regex"`
	Timeout        time.Duration  `yaml:"-"`
	TimeoutRaw     string         `yaml:"timeout"`
	CompiledRegex  *regexp.Regexp `yaml:"-"`
}

// TestSpec is the full parsed test.yaml.
type TestSpec struct {
	Name           string   `yaml:"name"`
	Profile        string   `yaml:"profile"`
	Cleanup        *bool    `yaml:"cleanup"`
	Requires       Requires `yaml:"requires"`
	PassthroughEnv []string `yaml:"passthrough_env"`
	Steps          []Step   `yaml:"steps"`
}

// CleanupEnabled reports whether destroy-on-completion is on. Honours the
// pointer-vs-default distinction: an explicit `cleanup: false` MUST stay
// false, while an omitted field defaults to DefaultCleanup.
func (s *TestSpec) CleanupEnabled() bool {
	if s.Cleanup == nil {
		return DefaultCleanup
	}
	return *s.Cleanup
}

// Load parses and validates the file at path. It also confirms the named
// profile exists in DefaultPath() (or DATABRICKS_CONFIG_FILE override).
// Returning a TestSpec implies "ready to run" — the runner doesn't have
// to validate again.
func Load(path string) (*TestSpec, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("config: open %s: %w", path, err)
	}
	defer f.Close()
	return load(f, path, profile.DefaultPath())
}

// LoadWithProfilePath is identical to Load except the caller supplies the
// path to ~/.databrickscfg (used by tests that want to point at a
// fixture file rather than the developer's real config).
func LoadWithProfilePath(testYamlPath, databricksCfgPath string) (*TestSpec, error) {
	f, err := os.Open(testYamlPath)
	if err != nil {
		return nil, fmt.Errorf("config: open %s: %w", testYamlPath, err)
	}
	defer f.Close()
	return load(f, testYamlPath, databricksCfgPath)
}

// load is the inner parser, factored so tests can supply an io.Reader
// directly. The path argument is used only for error messages and for
// resolving the test's home directory.
func load(r io.Reader, path, databricksCfgPath string) (*TestSpec, error) {
	dec := yaml.NewDecoder(r)
	dec.KnownFields(true) // unknown fields → error (catches typos like `cleanups: false`)
	var spec TestSpec
	if err := dec.Decode(&spec); err != nil {
		return nil, fmt.Errorf("config: parse %s: %w", path, err)
	}
	if err := dec.Decode(&struct{}{}); err != io.EOF {
		// More than one YAML document — we expect exactly one.
		return nil, fmt.Errorf("config: %s contains multiple YAML documents (only one allowed)", path)
	}
	if err := normalize(&spec); err != nil {
		return nil, fmt.Errorf("config: %s: %w", path, err)
	}
	if err := validate(&spec, databricksCfgPath); err != nil {
		return nil, fmt.Errorf("config: %s: %w", path, err)
	}
	return &spec, nil
}

// LoadDir is a convenience wrapper that loads `<dir>/test.yaml`.
func LoadDir(dir string) (*TestSpec, error) {
	return Load(filepath.Join(dir, "test.yaml"))
}

// normalize fills in defaults and parses string-encoded values that yaml
// itself doesn't decode for us (durations are the main one — yaml.v3
// doesn't natively decode a Go time.Duration without a yaml unmarshaler).
func normalize(spec *TestSpec) error {
	if spec.Requires.Cloud == "" {
		spec.Requires.Cloud = CloudAny
	}
	if spec.Requires.Level == "" {
		spec.Requires.Level = LevelWorkspace
	}
	for i := range spec.Steps {
		s := &spec.Steps[i]
		if s.Command == "" {
			s.Command = CommandApply
		}
		if s.Expect == "" {
			s.Expect = ExpectSuccess
		}
		if s.TimeoutRaw == "" {
			s.Timeout = DefaultStepTimeout
		} else {
			d, err := time.ParseDuration(s.TimeoutRaw)
			if err != nil {
				return fmt.Errorf("step %q: invalid timeout %q: %w", s.Name, s.TimeoutRaw, err)
			}
			if d <= 0 {
				return fmt.Errorf("step %q: timeout must be positive, got %s", s.Name, d)
			}
			s.Timeout = d
		}
		if s.ErrorRegex != "" {
			re, err := regexp.Compile(s.ErrorRegex)
			if err != nil {
				return fmt.Errorf("step %q: invalid error_regex %q: %w", s.Name, s.ErrorRegex, err)
			}
			s.CompiledRegex = re
		}
	}
	return nil
}

// slugRegexp matches the slug shape from DESIGN.md §4 schema rules
// (lowercase letters, digits, underscore, hyphen).
var slugRegexp = regexp.MustCompile(`^[a-z0-9_-]+$`)

// strictSemverRegexp matches `X.Y.Z` and `X.Y.Z-prerelease` strict semver.
// The regex is conservative: prerelease can include letters, digits,
// dots, and dashes (the SemVer 2.0 pre-release subset), and we forbid
// leading zeros to mirror Terraform's stricter parser.
var strictSemverRegexp = regexp.MustCompile(`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-[0-9A-Za-z.-]+)?$`)

// validate enforces every rule in DESIGN.md §4 + §10. Returning the
// first error keeps messages surgical; the caller invariably stops on
// the first issue anyway.
func validate(spec *TestSpec, databricksCfgPath string) error {
	if spec.Name == "" {
		return errors.New("name: required")
	}
	if !slugRegexp.MatchString(spec.Name) {
		return fmt.Errorf("name: %q is not a valid slug (must match %s)", spec.Name, slugRegexp)
	}
	if spec.Profile == "" {
		return errors.New("profile: required")
	}
	if !validCloud(spec.Requires.Cloud) {
		return fmt.Errorf("requires.cloud: %q is not one of [aws azure gcp any]", spec.Requires.Cloud)
	}
	if !validLevel(spec.Requires.Level) {
		return fmt.Errorf("requires.level: %q is not one of [workspace account ucws ucacct]", spec.Requires.Level)
	}
	if err := validatePassthroughEnv(spec.PassthroughEnv); err != nil {
		return err
	}
	if len(spec.Steps) == 0 {
		return errors.New("steps: at least one step is required")
	}
	if err := validateSteps(spec.Steps); err != nil {
		return err
	}
	// Profile existence is the last preflight check. Skipping it for
	// the empty-string case is unreachable (profile required above);
	// keeping the explicit guard so a future caller that bypasses Load
	// can't land here with an empty path.
	if databricksCfgPath != "" {
		ok, err := profile.SectionExists(databricksCfgPath, spec.Profile)
		if err != nil {
			return fmt.Errorf("profile: %w", err)
		}
		if !ok {
			return fmt.Errorf("profile: %q not found in %s", spec.Profile, databricksCfgPath)
		}
	}
	return nil
}

// validatePassthroughEnv enforces the §10/G6 invariant: no DATABRICKS_*
// names allowed (would defeat the profile mechanism). Empty names also
// rejected as a likely typo.
func validatePassthroughEnv(names []string) error {
	for _, n := range names {
		if strings.TrimSpace(n) == "" {
			return errors.New("passthrough_env: empty name not allowed")
		}
		if strings.HasPrefix(n, "DATABRICKS_") {
			return fmt.Errorf("passthrough_env: %q starts with DATABRICKS_ — use the `profile` field instead (DESIGN.md §4)", n)
		}
	}
	return nil
}

// validateSteps enforces step-level invariants and uniqueness of step
// names within the test. Per-step validation is split into validateStep
// to keep this loop short and the assertion logic isolated.
func validateSteps(steps []Step) error {
	seen := make(map[string]struct{}, len(steps))
	for i, s := range steps {
		if _, dup := seen[s.Name]; dup {
			return fmt.Errorf("steps[%d].name: duplicate step name %q", i, s.Name)
		}
		if err := validateStep(i, s); err != nil {
			return err
		}
		seen[s.Name] = struct{}{}
	}
	return nil
}

// validateStep checks a single step against the §4 schema rules.
func validateStep(i int, s Step) error {
	if s.Name == "" {
		return fmt.Errorf("steps[%d].name: required", i)
	}
	if !slugRegexp.MatchString(s.Name) {
		return fmt.Errorf("steps[%d].name: %q is not a valid slug", i, s.Name)
	}
	if s.Version == "" {
		return fmt.Errorf("steps[%d] (%s): version is required", i, s.Name)
	}
	if s.Version != LocalVersion && !strictSemverRegexp.MatchString(s.Version) {
		return fmt.Errorf("steps[%d] (%s): version %q must be %q or strict semver X.Y.Z[-prerelease]",
			i, s.Name, s.Version, LocalVersion)
	}
	if !validCommand(s.Command) {
		return fmt.Errorf("steps[%d] (%s): command %q is not one of [plan apply destroy]", i, s.Name, s.Command)
	}
	if !validExpect(s.Expect) {
		return fmt.Errorf("steps[%d] (%s): expect %q is not one of [success failure]", i, s.Name, s.Expect)
	}
	return validateStepAssertion(i, s)
}

// validateStepAssertion enforces the failure-assertion completeness rule
// from DESIGN.md §4: when expect=failure, at least one of error_substring
// / error_regex MUST be set; conversely, on expect=success the error_*
// fields are rejected as a likely copy-paste mistake.
func validateStepAssertion(i int, s Step) error {
	hasSub := s.ErrorSubstring != ""
	hasRe := s.ErrorRegex != ""
	switch s.Expect {
	case ExpectFailure:
		if !hasSub && !hasRe {
			return fmt.Errorf("steps[%d] (%s): expect=failure requires error_substring or error_regex", i, s.Name)
		}
	case ExpectSuccess:
		if hasSub || hasRe {
			return fmt.Errorf("steps[%d] (%s): error_substring/error_regex are only valid with expect=failure", i, s.Name)
		}
	}
	return nil
}

func validCloud(c Cloud) bool {
	return slices.Contains([]Cloud{CloudAWS, CloudAzure, CloudGCP, CloudAny}, c)
}

func validLevel(l Level) bool {
	return slices.Contains([]Level{LevelWorkspace, LevelAccount, LevelUCWS, LevelUCAcct}, l)
}

func validCommand(c Command) bool {
	return slices.Contains([]Command{CommandPlan, CommandApply, CommandDestroy}, c)
}

func validExpect(e Expect) bool {
	return e == ExpectSuccess || e == ExpectFailure
}
