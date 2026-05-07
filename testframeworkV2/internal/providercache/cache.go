// Package providercache caches Databricks Terraform provider binaries on the
// local filesystem in the layout that Terraform's `filesystem_mirror`
// provider_installation block expects.
//
// See DESIGN.md §6 ("Cache layout") for the overall design.
//
// In M1 the cache supports released versions only:
//
//	<root>/registry.terraform.io/databricks/databricks/terraform-provider-databricks_<version>_<os>_<arch>.zip
//
// Local-build (unpacked) layout is wired up in M6 (DESIGN.md §15). The
// SyntheticVersionLocal sentinel and the early-exit error in Resolve preserve
// the API shape so the runner does not need to change when local lands.
package providercache

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// Target identifies a provider build target (OS + Arch). It is the value
// Terraform calls "<os>_<arch>" in filesystem_mirror layouts and in release
// asset filenames (e.g. "darwin_arm64").
type Target struct {
	OS   string
	Arch string
}

// String returns the canonical "<os>_<arch>" form Terraform uses in
// filesystem_mirror filenames.
func (t Target) String() string {
	return t.OS + "_" + t.Arch
}

// HostTarget returns the Target for the current process. The runner uses
// this when no explicit target is configured.
func HostTarget() Target {
	return Target{OS: runtime.GOOS, Arch: runtime.GOARCH}
}

// SyntheticVersionLocal is the placeholder version used by the framework for
// locally-built providers (DESIGN.md §6 / §8). M1 does not implement local
// builds; this constant is exported so that callers and future milestones
// share a single source of truth.
const SyntheticVersionLocal = "99.0.0-local"

// LocalVersionInput is the literal string a test.yaml step uses to opt into a
// local build (DESIGN.md §4 / §8).
const LocalVersionInput = "local"

// DefaultDownloadBaseURL is the GitHub releases prefix for the published
// terraform-provider-databricks zip assets. Tests substitute a httptest.Server
// URL via WithBaseURL.
const DefaultDownloadBaseURL = "https://github.com/databricks/terraform-provider-databricks/releases/download"

// defaultDownloadTimeout caps a single asset download. The published zips are
// ~65 MB; 10 minutes is the same ceiling go-getter and other Terraform
// tooling use, and is generous enough for slow networks without letting a
// stuck connection block a CI job indefinitely.
const defaultDownloadTimeout = 10 * time.Minute

// Cache manages downloaded provider zips on the local filesystem. A single
// Cache value is safe for concurrent use across goroutines and processes:
// every download writes to a unique "<dst>.partial.<rand>" file and then
// atomically renames into place (DESIGN.md §6 "Cache atomicity"), so two
// parallel runs racing on the same uncached version both succeed and end with
// identical content at <dst>.
type Cache struct {
	root string

	httpClient *http.Client
	baseURL    string
}

// Option configures a Cache at construction time.
type Option func(*Cache)

// WithHTTPClient overrides the default *http.Client used for downloads.
// Useful for tests that want to inject a custom Transport.
func WithHTTPClient(c *http.Client) Option {
	return func(ca *Cache) {
		if c != nil {
			ca.httpClient = c
		}
	}
}

// WithBaseURL overrides the GitHub releases base URL used for downloads.
// Primarily useful for unit tests that point the cache at httptest.Server.
func WithBaseURL(u string) Option {
	return func(ca *Cache) {
		if u != "" {
			ca.baseURL = strings.TrimRight(u, "/")
		}
	}
}

// New constructs a Cache rooted at root. The directory is created lazily on
// the first download — New does not touch the filesystem.
//
// root is typically `~/.testframeworkv2/providers` (DESIGN.md §3).
func New(root string, opts ...Option) *Cache {
	c := &Cache{
		root:       root,
		httpClient: &http.Client{Timeout: defaultDownloadTimeout},
		baseURL:    DefaultDownloadBaseURL,
	}
	for _, o := range opts {
		o(c)
	}
	return c
}

// Root returns the filesystem path the cache is rooted at. Callers point
// `filesystem_mirror.path` at this directory in the generated `.terraformrc`.
func (c *Cache) Root() string { return c.root }

// Resolve returns a filesystem path that Terraform's filesystem_mirror can
// serve `databricks/databricks` from at the given (version, target). It also
// returns the version string the framework should pin in the generated
// `_tfv2_versions_override.tf`. For released versions the synthetic version
// equals the (normalized) input; for local builds (M6) it is the
// SyntheticVersionLocal placeholder.
//
// Released-version flow (M1):
//
//  1. compute the canonical zip path under c.root,
//  2. if cached, return it immediately,
//  3. otherwise download the asset and atomically install it.
//
// Versions starting with "v" are normalized to drop the prefix — GitHub
// release tags use "v1.114.0" while Terraform's filesystem_mirror filenames
// use "1.114.0".
func (c *Cache) Resolve(ctx context.Context, version string, target Target) (path, syntheticVersion string, err error) {
	if version == "" {
		return "", "", errors.New("providercache: version is required")
	}
	if version == LocalVersionInput {
		// M6 wires up local builds. Surfacing a clear error is much more
		// helpful than the bewildering 404 we would otherwise hit while
		// trying to GET v"local" from GitHub.
		return "", "", fmt.Errorf("providercache: version %q is not yet implemented (DESIGN.md §15 M6)", LocalVersionInput)
	}
	if target.OS == "" || target.Arch == "" {
		return "", "", fmt.Errorf("providercache: target requires both OS and Arch (got %q)", target.String())
	}

	v := normalizeVersion(version)
	zip := c.zipPath(v, target)

	if _, statErr := os.Stat(zip); statErr == nil {
		return zip, v, nil
	} else if !errors.Is(statErr, os.ErrNotExist) {
		return "", "", fmt.Errorf("providercache: stat %s: %w", zip, statErr)
	}

	if err := c.downloadZip(ctx, v, target, zip); err != nil {
		return "", "", err
	}
	return zip, v, nil
}

// zipPath returns the canonical packed-mirror filesystem path for a given
// (version, target) pair, irrespective of whether the file currently exists
// on disk. Exposed for tests and for the future `tfv2 cache list` subcommand.
func (c *Cache) zipPath(version string, target Target) string {
	return filepath.Join(c.providerDir(), zipFilename(version, target))
}

// providerDir returns the registry.terraform.io/databricks/databricks subtree
// — the directory that holds either packed zips (released versions) or the
// `<version>/<target>/` unpacked layout (local builds, M6).
func (c *Cache) providerDir() string {
	return filepath.Join(c.root, "registry.terraform.io", "databricks", "databricks")
}

// zipFilename matches Terraform's packed filesystem_mirror naming convention:
// terraform-provider-<TYPE>_<VERSION>_<TARGET>.zip.
func zipFilename(version string, target Target) string {
	return fmt.Sprintf("terraform-provider-databricks_%s_%s.zip", version, target.String())
}

// normalizeVersion strips a leading "v" if present. GitHub release tags use
// "v1.114.0"; Terraform's filesystem_mirror filenames use "1.114.0".
func normalizeVersion(version string) string {
	return strings.TrimPrefix(version, "v")
}
