// Command tfv2 is the CLI entry point for testframeworkV2 — the
// multi-version Terraform test harness. Subcommands:
//
//	tfv2 run <test-dir> run a single test
//	tfv2 run -r <root> recursively run every test under root
//	tfv2 cache list show cached versions
//	tfv2 cache prune delete provider cache
//	tfv2 build local --repo <p> eagerly build local provider into cache
//	tfv2 version print version
//	tfv2 help show usage banner
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

// exitCodeOK and friends keep the exit code policy in one place. The
// CLI scripts that invoke tfv2 use these as their pass/fail oracle.
const (
	exitCodeOK     = 0 // every step passed (or the run was skipped)
	exitCodeFailed = 1 // at least one step did not pass
	exitCodeUsage  = 2 // bad flags or missing arguments
)

// version is set by the build (e.g., goreleaser ldflags). Defaults to
// "dev" so plain `go build` still produces a useful banner.
var version = "dev"

func main() {
	os.Exit(run(os.Args[1:]))
}

// run is the testable entry point — main delegates here so unit tests
// can exercise flag parsing without touching os.Exit. It returns the
// exit code the process should exit with.
func run(args []string) int {
	if len(args) == 0 {
		printUsage(os.Stderr)
		return exitCodeUsage
	}
	switch args[0] {
	case "run":
		return runRun(args[1:])
	case "cache":
		return runCache(args[1:])
	case "build":
		return runBuild(args[1:])
	case "version", "--version", "-v":
		fmt.Fprintf(os.Stdout, "tfv2 %s\n", version)
		return exitCodeOK
	case "help", "--help", "-h":
		printUsage(os.Stdout)
		return exitCodeOK
	default:
		fmt.Fprintf(os.Stderr, "tfv2: unknown subcommand %q\n\n", args[0])
		printUsage(os.Stderr)
		return exitCodeUsage
	}
}

// printUsage emits a top-level usage banner. Each subcommand
// (run / cache / build) has its own --help via the flag.FlagSet
// returned by parseRunFlags / parseCacheFlags / parseBuildFlags.
func printUsage(w *os.File) {
	fmt.Fprintln(w, `tfv2 — testframeworkV2 multi-version Terraform test harness

usage:
  tfv2 run <test-dir>             run a single test
  tfv2 run -r <root>              recursively run every test.yaml under root
  tfv2 cache list                 show cached provider versions
  tfv2 cache prune                delete the provider cache
  tfv2 build local --repo <path>  eagerly build local provider into cache
  tfv2 version                    print version
  tfv2 help                       show this banner

flags for `+"`tfv2 run`"+`:
  -r, --recursive                 walk <root> for nested test.yaml files
  --terraform-bin <path>          override terraform binary discovery
  --cache-dir <path>              override ~/.testframeworkv2/providers
  --run-dir <path>                override ~/.testframeworkv2/runs root
  --repo <path>                   provider repo root (for version=local)
  --no-cleanup                    disable cleanup destroy regardless of test.yaml
  --verbose                       print framework debug logs to stderr

env-var equivalents (CLI flags win):
  TFV2_TERRAFORM_BIN              = --terraform-bin
  TFV2_CACHE_DIR                  = --cache-dir
  T_NO_CLEANUP=1                  = --no-cleanup`)
}

// runFlags holds the parsed `tfv2 run` options. Kept as a struct so the
// dispatcher in runRun can be unit-tested by constructing one directly
// without a flag.FlagSet round-trip.
type runFlags struct {
	terraformBin string
	cacheDir     string
	runDir       string
	repoRoot     string
	noCleanup    bool
	verbose      bool
	recursive    bool
	testDir      string
}

// parseRunFlags handles the `tfv2 run` flag set. Returns a populated
// runFlags or an error suitable for stderr printing. Defaults pull
// from the documented env-var equivalents.
func parseRunFlags(args []string) (runFlags, error) {
	fs := flag.NewFlagSet("run", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)

	var f runFlags
	fs.StringVar(&f.terraformBin, "terraform-bin", os.Getenv("TFV2_TERRAFORM_BIN"), "override terraform binary discovery")
	fs.StringVar(&f.cacheDir, "cache-dir", os.Getenv("TFV2_CACHE_DIR"), "override ~/.testframeworkv2/providers")
	fs.StringVar(&f.runDir, "run-dir", "", "override ~/.testframeworkv2/runs root")
	fs.StringVar(&f.repoRoot, "repo", os.Getenv("TFV2_REPO"), "provider repo root for version=local (auto-discovered if empty)")
	fs.BoolVar(&f.noCleanup, "no-cleanup", false, "disable cleanup destroy")
	fs.BoolVar(&f.verbose, "verbose", false, "print framework debug logs")
	fs.BoolVar(&f.recursive, "r", false, "recursively walk <test-dir> for nested test.yaml files")
	fs.BoolVar(&f.recursive, "recursive", false, "recursively walk <test-dir> for nested test.yaml files (alias of -r)")
	if err := fs.Parse(args); err != nil {
		return runFlags{}, err
	}
	if fs.NArg() != 1 {
		return runFlags{}, fmt.Errorf("expected exactly one <test-dir> argument, got %d", fs.NArg())
	}
	f.testDir = fs.Arg(0)
	return f, nil
}

// signalContext returns a context that is cancelled on SIGINT / SIGTERM
// so the runner gracefully aborts a long-running step (e.g. an apply
// the user wants to interrupt). The cancellation propagates through
// tfexec via the per-step context.WithTimeout chain.
func signalContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		cancel()
	}()
	return ctx, cancel
}

// defaultCacheDir resolves the cache root the runner should use when
// the --cache-dir flag (and TFV2_CACHE_DIR env) are unset.
func defaultCacheDir() string {
	home, err := os.UserHomeDir()
	if err != nil || home == "" {
		// Fall back to CWD/.testframeworkv2 so the CLI still runs on
		// shells with no $HOME (rare but possible in CI).
		return filepath.Join(".testframeworkv2", "providers")
	}
	return filepath.Join(home, ".testframeworkv2", "providers")
}

// defaultRunDir is the analogous default for run roots.
func defaultRunDir() string {
	home, err := os.UserHomeDir()
	if err != nil || home == "" {
		return filepath.Join(".testframeworkv2", "runs")
	}
	return filepath.Join(home, ".testframeworkv2", "runs")
}
