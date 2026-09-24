package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/providercache"
)

// runBuild dispatches `tfv2 build <local>`. Reserved structure for
// future build modes (cross-arch, vendored deps, etc.).
func runBuild(args []string) int {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "tfv2 build: missing mode (local)")
		return exitCodeUsage
	}
	switch args[0] {
	case "local":
		return runBuildLocal(args[1:])
	default:
		fmt.Fprintf(os.Stderr, "tfv2 build: unknown mode %q (expected local)\n", args[0])
		return exitCodeUsage
	}
}

// buildLocalFlags holds parsed `tfv2 build local` options.
type buildLocalFlags struct {
	cacheDir string
	repoRoot string
}

func parseBuildLocal(args []string) (buildLocalFlags, error) {
	fs := flag.NewFlagSet("build local", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	var f buildLocalFlags
	fs.StringVar(&f.cacheDir, "cache-dir", os.Getenv("TFV2_CACHE_DIR"), "override ~/.testframeworkv2/providers")
	fs.StringVar(&f.repoRoot, "repo", "", "provider repo root (required)")
	if err := fs.Parse(args); err != nil {
		return buildLocalFlags{}, err
	}
	if fs.NArg() != 0 {
		return buildLocalFlags{}, fmt.Errorf("unexpected positional argument %q", fs.Arg(0))
	}
	if f.repoRoot == "" {
		return buildLocalFlags{}, fmt.Errorf("--repo is required")
	}
	if f.cacheDir == "" {
		f.cacheDir = defaultCacheDir()
	}
	return f, nil
}

// runBuildLocal eagerly compiles the provider from --repo into the
// cache. Useful for warming the cache before a CI run that sets
// version=local in many tests.
func runBuildLocal(args []string) int {
	f, err := parseBuildLocal(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "tfv2 build local: %v\n", err)
		return exitCodeUsage
	}
	ctx, cancel := signalContext()
	defer cancel()

	c := providercache.New(f.cacheDir)
	binDir, syn, prov, err := c.BuildLocal(ctx, f.repoRoot, providercache.HostTarget())
	if err != nil {
		fmt.Fprintf(os.Stderr, "tfv2 build local: %v\n", err)
		return exitCodeFailed
	}
	fmt.Fprintf(os.Stdout, "built %s\n", syn)
	fmt.Fprintf(os.Stdout, "  binary: %s\n", binDir)
	fmt.Fprintf(os.Stdout, "  git_sha: %s (dirty=%v)\n", prov.GitSHA, prov.Dirty)
	fmt.Fprintf(os.Stdout, "  go_version: %s os_arch: %s\n", prov.GoVersion, prov.OSArch)
	fmt.Fprintf(os.Stdout, "  built_at: %s\n", prov.BuiltAt.Format("2006-01-02T15:04:05Z"))
	return exitCodeOK
}
