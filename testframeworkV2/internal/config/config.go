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

	// v2-mode fields. A non-empty Config OR a non-empty Assert flips the
	// containing TestSpec into v2 mode; the runner then wipes
	// non-framework `*.tf` from the workdir and copies Step.Config in
	// before each step's `terraform init` (rather than the v1 model
	// where the user's full `*.tf` set is copied once at run start).
	//
	// Config is a path RELATIVE to the test directory containing
	// test.yaml — e.g. "step1_create.tf" if test.yaml lives in
	// `tests/token_lifecycle_v2/`.
	//
	// Assert runs against `terraform show -json` after the step's
	// command succeeds. Assertions are only consulted when
	// Expect == ExpectSuccess (failure-path steps don't have meaningful
	// state to inspect).
	Config string      `yaml:"config"`
	Assert []Assertion `yaml:"assert"`

	// Plan-content matchers (DESIGN.md §17.10). Both fields require
	// `command: plan` AND `expect: success`; the runner evaluates them
	// against terraform plan's stdout (which contains the diff
	// annotations: "Plan: X to add...", "# forces replacement",
	// "No changes.", etc.).
	//
	// ExpectNonEmptyPlan asserts the plan is NOT empty — i.e. stdout
	// does NOT contain "No changes.". Useful for fixtures that should
	// always show a diff (regression-guard against silent acceptance).
	//
	// PlanMatch is a Go RE2 string matched against plan stdout with
	// the multiline `(?s)` flag implicit (so `.` spans newlines). Used
	// to anchor on stable phrases like "# forces replacement" or
	// "Plan: 1 to add, 0 to change, 1 to destroy" without needing a
	// post-step manual log grep.
	//
	// Both AND together when both set. Both default off so existing
	// fixtures keep their behaviour exactly.
	ExpectNonEmptyPlan bool           `yaml:"expect_non_empty_plan"`
	PlanMatch          string         `yaml:"plan_match"`
	CompiledPlanMatch  *regexp.Regexp `yaml:"-"`
}

// Assertion is one (resource address, presence, attributes) bundle —
// the v2 mode's structured replacement for stderr-regex matching.
//
// Resource is the canonical Terraform address: `<type>.<name>` for
// managed resources (e.g. `databricks_token.pat`) and
// `data.<type>.<name>` for data sources (e.g.
// `data.databricks_mws_workspaces.all`). Module-scoped addresses are
// out of scope for v2 launch — root module only.
//
// Present is a pointer-bool so a YAML-omitted field is distinguishable
// from an explicit `present: false`. Default (omitted) is true; the
// PresentValue helper centralizes the defaulting logic.
//
// Attrs is the (key, expected-value) map evaluated only when Present is
// true (or omitted). Values use Go-typed YAML scalars / sequences /
// mappings; the runner normalizes numerics to float64 (matching
// `terraform show -json` decoding) before comparing.
type Assertion struct {
	Resource string         `yaml:"resource"`
	Present  *bool          `yaml:"present"`
	Attrs    map[string]any `yaml:"attrs"`
}

// PresentValue returns the effective Present value, defaulting to true
// when the field was omitted in YAML. Mirrors TestSpec.CleanupEnabled().
func (a *Assertion) PresentValue() bool {
	if a.Present == nil {
		return true
	}
	return *a.Present
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

// Mode is the v1/v2 enum for a parsed TestSpec. The runner branches
// on this to decide between "copy *.tf once + run multi-step" (v1)
// and "wipe-and-copy per step" (v2).
type Mode string

const (
	ModeV1 Mode = "v1"
	ModeV2 Mode = "v2"
)

// Mode reports v1 vs v2. Determined by the FIRST step's Config field:
// because validateV2Consistency enforces all-or-none Config across
// steps, looking at steps[0] is sufficient and unambiguous after
// validation completes (DESIGN.md §17.2).
func (s *TestSpec) Mode() Mode {
	if len(s.Steps) > 0 && s.Steps[0].Config != "" {
		return ModeV2
	}
	return ModeV1
}

// IsV2 is a convenience predicate for callers that prefer a bool over
// the enum. Equivalent to Mode() == ModeV2.
func (s *TestSpec) IsV2() bool { return s.Mode() == ModeV2 }

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

// LoadDir is a convenience wrapper that loads `<dir>/test.yaml` and,
// for v2-mode specs, additionally verifies that every step's
// `config:` references an existing file under `dir`. The existence
// check is dir-aware (vs. `Load`, which only sees the test.yaml path)
// and so lives here. Surfacing typo'd `config:` paths at parse time —
// before any step runs — prevents a 5-step test from applying steps
// 1-4 (mutating real cloud resources) only to fail at step 5 with a
// "no such file or directory".
func LoadDir(dir string) (*TestSpec, error) {
	yamlPath := filepath.Join(dir, "test.yaml")
	spec, err := Load(yamlPath)
	if err != nil {
		return nil, err
	}
	if spec.IsV2() {
		for i, s := range spec.Steps {
			cfgPath := filepath.Join(dir, s.Config)
			if _, err := os.Stat(cfgPath); err != nil {
				return nil, fmt.Errorf("config: %s: steps[%d] (%s): config file %q not found in %s: %w", yamlPath, i, s.Name, s.Config, dir, err)
			}
		}
	}
	return spec, nil
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
		if s.PlanMatch != "" {
			// Multiline by default — plan output has line breaks and
			// callers anchor on phrases like "# forces replacement"
			// that follow a leading "  - resource ..." line.
			re, err := regexp.Compile(`(?s)` + s.PlanMatch)
			if err != nil {
				return fmt.Errorf("step %q: invalid plan_match %q: %w", s.Name, s.PlanMatch, err)
			}
			s.CompiledPlanMatch = re
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
	if err := validateV2Consistency(spec.Steps); err != nil {
		return err
	}
	return validateProfileExists(spec.Profile, databricksCfgPath)
}

// validateProfileExists is the parse-time profile-existence preflight
// (DESIGN.md §4 "Pre-flight validation"). Skipping it for the empty-
// string case is unreachable from Load (profile required above), but
// the explicit guard keeps the function safe for callers that bypass
// Load.
func validateProfileExists(profileName, databricksCfgPath string) error {
	if databricksCfgPath == "" {
		return nil
	}
	ok, err := profile.SectionExists(databricksCfgPath, profileName)
	if err != nil {
		return fmt.Errorf("profile: %w", err)
	}
	if !ok {
		// Wrap the sentinel so callers can detect "profile missing" via
		// errors.Is(err, profile.ErrSectionNotFound). fixtures_test.go
		// uses this to t.Skip rather than t.Fatal when the test
		// environment doesn't have the profile in question.
		return fmt.Errorf("profile: %q not found in %s: %w",
			profileName, databricksCfgPath, profile.ErrSectionNotFound)
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
//
// v2-mode invariants: state assertions (Step.Assert) only make sense
// when a step is expected to succeed AND the command leaves state
// behind (apply or destroy — plan doesn't mutate state). Reject
// Assert on expect=failure or command=plan as a likely
// misconfiguration.
func validateStepAssertion(i int, s Step) error {
	hasSub := s.ErrorSubstring != ""
	hasRe := s.ErrorRegex != ""
	switch s.Expect {
	case ExpectFailure:
		if !hasSub && !hasRe {
			return fmt.Errorf("steps[%d] (%s): expect=failure requires error_substring or error_regex", i, s.Name)
		}
		if len(s.Assert) > 0 {
			return fmt.Errorf("steps[%d] (%s): assert is only valid with expect=success", i, s.Name)
		}
	case ExpectSuccess:
		if hasSub || hasRe {
			return fmt.Errorf("steps[%d] (%s): error_substring/error_regex are only valid with expect=failure", i, s.Name)
		}
	}
	for j, a := range s.Assert {
		if a.Resource == "" {
			return fmt.Errorf("steps[%d] (%s): assert[%d].resource is required", i, s.Name, j)
		}
		// present:false + attrs:set is logically inconsistent — can't
		// inspect attrs of a resource we expect to NOT exist
		// (DESIGN.md §17.7 rule 6).
		if a.Present != nil && !*a.Present && len(a.Attrs) > 0 {
			return fmt.Errorf("steps[%d] (%s): assert[%d] cannot set attrs when present is false", i, s.Name, j)
		}
		// Resource address shape: managed = `<type>.<name>`; data
		// sources = `data.<type>.<name>`. v2 launch is root-module
		// only — module-scoped addresses are deferred (§17.5 / §17.9).
		if !validResourceAddress(a.Resource) {
			return fmt.Errorf("steps[%d] (%s): assert[%d].resource %q is not a valid root-module address (expected `type.name` or `data.type.name`)",
				i, s.Name, j, a.Resource)
		}
	}
	return validatePlanMatchers(i, s)
}

// validatePlanMatchers enforces the §17.10 invariants on the new
// plan-content matcher fields:
//
//   - `expect_non_empty_plan: true` requires `command: plan` AND
//     `expect: success`. plan-content matchers against an apply or
//     destroy stdout would be checking the wrong thing (terraform
//     prints different summary lines), and against a failed plan
//     they're noise (the cmdErr already explains the failure).
//   - `plan_match: <regex>` has the same two requirements.
//
// Both fields default off (false / empty); existing fixtures see no
// behaviour change.
func validatePlanMatchers(i int, s Step) error {
	if !s.ExpectNonEmptyPlan && s.PlanMatch == "" {
		return nil
	}
	field := "expect_non_empty_plan"
	if s.PlanMatch != "" {
		field = "plan_match"
	}
	if s.Command != CommandPlan {
		return fmt.Errorf("steps[%d] (%s): %s requires command: plan (got %q)", i, s.Name, field, s.Command)
	}
	if s.Expect != ExpectSuccess {
		return fmt.Errorf("steps[%d] (%s): %s requires expect: success (got %q)", i, s.Name, field, s.Expect)
	}
	return nil
}

// validResourceAddress reports whether addr is a root-module
// Terraform resource address (DESIGN.md §17.5). Two shapes:
//
//   - managed: `<type>.<name>` (e.g. `databricks_token.pat`)
//   - data:    `data.<type>.<name>` (e.g.
//     `data.databricks_mws_workspaces.all`)
//
// Validates structurally (split on `.`) rather than via a single
// regex because Go's RE2 lacks the negative lookahead needed to
// distinguish "data.X" (a 2-part malformed address) from "data.X.Y"
// (a valid 3-part data-source address).
func validResourceAddress(addr string) bool {
	parts := strings.Split(addr, ".")
	switch len(parts) {
	case 2:
		// Managed resource. Reject `data.<x>` as a 2-part address —
		// that's an incomplete data-source address (`data.<type>.<name>`
		// requires 3 parts).
		if parts[0] == "data" {
			return false
		}
		return v2AddrTypeRegexp.MatchString(parts[0]) && v2AddrNameRegexp.MatchString(parts[1])
	case 3:
		return parts[0] == "data" &&
			v2AddrTypeRegexp.MatchString(parts[1]) &&
			v2AddrNameRegexp.MatchString(parts[2])
	default:
		return false
	}
}

var (
	// v2AddrTypeRegexp constrains the resource type segment
	// (`databricks_token`, `databricks_mws_workspaces`).
	v2AddrTypeRegexp = regexp.MustCompile(`^[a-z][a-z0-9_]*$`)

	// v2AddrNameRegexp constrains the resource name segment, allowing
	// the slightly-broader Terraform identifier shape
	// (`pat`, `MyResource_42`, `tag-foo`).
	v2AddrNameRegexp = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_-]*$`)
)

// validateV2Consistency enforces the all-or-none invariant from
// DESIGN.md §17.2: a test.yaml is v2 iff every step has a non-empty
// `config:`, v1 iff no step has one. Mixed configurations are a
// parse-time error.
//
// Additional v2-mode rules enforced here:
//
//   - `config:` paths must be slug-shaped basenames ending in `.tf`
//     or `.tf.json`. No path traversal (`../escape.tf`), no
//     subdirectories, no hidden files.
//   - `config:` must NOT start with the framework's `_tfv2_` prefix
//     (collision risk with `_tfv2_versions_override.tf`; §17.4).
//   - `assert:` blocks may only appear in v2 specs. A v1 step (no
//     `config:`) with `assert:` set is a parse error (§17.7 rule 3).
//
// Per-Assertion shape rules (resource address, present+attrs
// consistency) live on validateStepAssertion.
func validateV2Consistency(steps []Step) error {
	configured := 0
	for _, s := range steps {
		if s.Config != "" {
			configured++
		}
	}
	mode := ModeV1
	switch configured {
	case 0:
		// pure v1; assert: not allowed.
		for i, s := range steps {
			if len(s.Assert) > 0 {
				return fmt.Errorf("steps[%d] (%s): assert: requires v2 mode (set config: on every step) — see DESIGN.md §17.7", i, s.Name)
			}
		}
		return nil
	case len(steps):
		mode = ModeV2
	default:
		// mixed.
		for i, s := range steps {
			if s.Config == "" {
				return fmt.Errorf("steps[%d] (%s): v2 mode requires every step to set `config:` (this test mixes v1 and v2 steps) — see DESIGN.md §17.2", i, s.Name)
			}
		}
	}
	_ = mode
	for i, s := range steps {
		if !v2ConfigPathRegexp.MatchString(s.Config) {
			return fmt.Errorf("steps[%d] (%s): config %q must be a slug-shaped .tf or .tf.json basename (e.g. step1_create.tf)", i, s.Name, s.Config)
		}
		if strings.HasPrefix(s.Config, "_tfv2_") {
			return fmt.Errorf("steps[%d] (%s): config %q must not start with `_tfv2_` (framework-reserved prefix) — see DESIGN.md §17.4", i, s.Name, s.Config)
		}
	}
	return nil
}

// v2ConfigPathRegexp constrains Step.Config to a slug-shaped basename
// ending in `.tf` or `.tf.json`. Matches `step1_create.tf`,
// `apply-with-tag.tf`, and `main.tf.json`; rejects `../escape.tf`,
// `subdir/file.tf`, hidden `.foo.tf`, and non-`.tf` extensions.
var v2ConfigPathRegexp = regexp.MustCompile(`^[a-zA-Z0-9_][a-zA-Z0-9_-]*\.(tf|tf\.json)$`)

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
