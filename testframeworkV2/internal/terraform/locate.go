// Package terraform locates the `terraform` binary that the runner spawns
// and sanity-checks its version against testframeworkV2's minimum
// requirement (1.5.0).
//
// The framework deliberately does NOT use hc-install or any
// auto-installer / B7 (Hashicorp's release-signing
// key has expired, breaking installer flows). Instead we resolve a
// known-good binary the user has already provisioned.
package terraform

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// MinVersion is the lowest terraform release version testframeworkV2
// supports. Bumped here is a breaking change — agree on it across the
// team. 1.5.0 was chosen because it's the LTS-style baseline aligned with
// the parent provider's compatibility matrix.
const MinVersion = "1.5.0"

// EnvVarBin is the framework's environment variable for overriding the
// terraform binary location. Equivalent to the --terraform-bin CLI flag,
// but useful in CI matrices where flags are awkward.
const EnvVarBin = "TFV2_TERRAFORM_BIN"

// ErrNotFound is returned when no terraform binary can be located via the
// resolution order.
var ErrNotFound = errors.New("terraform binary not found; install terraform >= " + MinVersion + " or pass --terraform-bin")

// Locate resolves the terraform binary using the documented order
// :
//
// 1. override (typically wired up to the --terraform-bin CLI flag)
// 2. TFV2_TERRAFORM_BIN environment variable
// 3. exec.LookPath("terraform") — first match in the user's PATH
// 4. ErrNotFound
//
// The returned path is always absolute. Callers are expected to pass the
// path to tfexec.NewTerraform; sanity-checking the version is a separate
// step (see SanityCheckVersion or LocateAndCheck).
func Locate(override string) (string, error) {
	if p := strings.TrimSpace(override); p != "" {
		return validateBinary(p, "--terraform-bin")
	}
	if p := strings.TrimSpace(os.Getenv(EnvVarBin)); p != "" {
		return validateBinary(p, EnvVarBin)
	}
	p, err := exec.LookPath("terraform")
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrNotFound, err)
	}
	return validateBinary(p, "$PATH")
}

// validateBinary confirms that the resolved path exists and is a regular
// file (not a directory) before returning it. Returning a clear error
// here avoids passing a stat-failing path through to tfexec which fails
// with a more cryptic error.
func validateBinary(path, source string) (string, error) {
	info, err := os.Stat(path)
	if err != nil {
		return "", fmt.Errorf("terraform: %s points at %q which does not exist: %w", source, path, err)
	}
	if info.IsDir() {
		return "", fmt.Errorf("terraform: %s points at %q which is a directory", source, path)
	}
	return path, nil
}

// SanityCheckVersion runs `<bin> -version`, parses the first output line,
// and returns the version. The returned error wraps a clear message when
// the version is below MinVersion or the output is unrecognizable.
func SanityCheckVersion(ctx context.Context, bin string) (string, error) {
	cmd := exec.CommandContext(ctx, bin, "-version")
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("terraform: running %q -version: %w", bin, err)
	}
	v, err := ParseVersion(string(out))
	if err != nil {
		return "", err
	}
	cmp, err := compareSemver(v, MinVersion)
	if err != nil {
		// Unparseable version reaches here only for non-strict-semver
		// values like "0.13.0-rc1"; surface a useful message.
		return v, fmt.Errorf("terraform: cannot compare version %q against minimum %q: %w", v, MinVersion, err)
	}
	if cmp < 0 {
		return v, fmt.Errorf("terraform: version %q is below minimum %q", v, MinVersion)
	}
	return v, nil
}

// LocateAndCheck combines Locate and SanityCheckVersion. The runner uses
// this at startup; CLI subcommands that don't need terraform call only
// Locate.
func LocateAndCheck(ctx context.Context, override string) (path, version string, err error) {
	path, err = Locate(override)
	if err != nil {
		return "", "", err
	}
	version, err = SanityCheckVersion(ctx, path)
	if err != nil {
		return path, version, err
	}
	return path, version, nil
}

// ParseVersion extracts the version string (e.g. "1.7.0") from the
// `terraform -version` output. The first line follows the shape:
//
//	Terraform v1.7.0
//
// optionally followed by build metadata. We strip the "v" prefix and
// trim trailing whitespace. Any non-matching first line returns an
// error so we don't silently accept garbage.
func ParseVersion(output string) (string, error) {
	first, _, _ := strings.Cut(strings.TrimSpace(output), "\n")
	first = strings.TrimSpace(first)
	const prefix = "Terraform v"
	if !strings.HasPrefix(first, prefix) {
		return "", fmt.Errorf("terraform: cannot parse version from %q (expected %q prefix)", first, prefix)
	}
	v := strings.TrimPrefix(first, prefix)
	// `terraform -version` may print extra fields after the version on
	// some platforms; cut at the first whitespace so "1.7.0 (foo)"
	// parses as "1.7.0".
	if i := strings.IndexAny(v, " \t"); i >= 0 {
		v = v[:i]
	}
	if v == "" {
		return "", fmt.Errorf("terraform: empty version in %q", first)
	}
	return v, nil
}

// compareSemver compares two strict major.minor.patch versions. Returns
// -1 if a<b, 0 if equal, +1 if a>b. Any prerelease suffix is stripped
// before comparison (we don't support fine-grained prerelease ordering;
// the framework's MinVersion is a stable release).
func compareSemver(a, b string) (int, error) {
	ap, err := splitSemver(a)
	if err != nil {
		return 0, err
	}
	bp, err := splitSemver(b)
	if err != nil {
		return 0, err
	}
	for i := range 3 {
		if ap[i] < bp[i] {
			return -1, nil
		}
		if ap[i] > bp[i] {
			return 1, nil
		}
	}
	return 0, nil
}

// splitSemver returns the [major, minor, patch] integers from a strict
// version string. Any "-prerelease" or "+build" suffix is dropped
// before splitting (we don't compare prerelease ordering).
func splitSemver(v string) ([3]int, error) {
	core := v
	if i := strings.IndexAny(v, "-+"); i >= 0 {
		core = v[:i]
	}
	parts := strings.Split(core, ".")
	if len(parts) != 3 {
		return [3]int{}, fmt.Errorf("not a strict major.minor.patch version: %q", v)
	}
	var out [3]int
	for i, p := range parts {
		n, err := strconv.Atoi(p)
		if err != nil {
			return [3]int{}, fmt.Errorf("invalid integer %q in version %q: %w", p, v, err)
		}
		out[i] = n
	}
	return out, nil
}
