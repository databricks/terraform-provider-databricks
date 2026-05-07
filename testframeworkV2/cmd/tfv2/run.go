package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/config"
	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/profile"
	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/result"
	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/runner"
	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/terraform"
)

// runRun handles the `tfv2 run <test-dir>` subcommand. It returns the
// exit code main should pass to os.Exit. Errors during setup return
// exitCodeUsage / exitCodeFailed depending on whether they're caused
// by the user's invocation or by the test outcome.
func runRun(args []string) int {
	f, err := parseRunFlags(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "tfv2 run: %v\n\n", err)
		printUsage(os.Stderr)
		return exitCodeUsage
	}
	ctx, cancel := signalContext()
	defer cancel()

	res, err := runOnce(ctx, f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "tfv2 run: %v\n", err)
		return exitCodeFailed
	}
	printRunResult(os.Stdout, res)
	if !res.AllPassed() {
		return exitCodeFailed
	}
	return exitCodeOK
}

// runOnce performs the full test execution: locate terraform, load
// config, load profile, build runner, Run. Factored out of runRun so
// tests can call it directly with a synthetic runFlags + injected
// runner Options.
func runOnce(ctx context.Context, f runFlags) (result.RunResult, error) {
	absTestDir, err := filepath.Abs(f.testDir)
	if err != nil {
		return result.RunResult{}, fmt.Errorf("resolve test-dir %q: %w", f.testDir, err)
	}
	bin, _, err := terraform.LocateAndCheck(ctx, f.terraformBin)
	if err != nil {
		return result.RunResult{}, err
	}
	spec, err := config.LoadDir(absTestDir)
	if err != nil {
		return result.RunResult{}, err
	}
	prof, err := profile.Load(spec.Profile)
	if err != nil {
		// Profile-existence is also checked at config-load time. If we
		// reach here the .databrickscfg contents changed between
		// validation and load, which is rare but worth surfacing.
		if errors.Is(err, profile.ErrSectionNotFound) {
			return result.RunResult{}, fmt.Errorf("profile %q not found in ~/.databrickscfg (set DATABRICKS_CONFIG_FILE if your config lives elsewhere)", spec.Profile)
		}
		return result.RunResult{}, err
	}
	r := runner.New(spec, prof, optionsFromFlags(f, bin, absTestDir))
	return r.Run(ctx)
}

// optionsFromFlags maps the CLI flag struct into runner.Options.
// Defaults for cache-dir and run-dir are filled in from the user's
// home directory.
func optionsFromFlags(f runFlags, terraformBin, sourceDir string) runner.Options {
	cacheDir := f.cacheDir
	if cacheDir == "" {
		cacheDir = defaultCacheDir()
	}
	runDir := f.runDir
	if runDir == "" {
		runDir = defaultRunDir()
	}
	return runner.Options{
		SourceDir:    sourceDir,
		CacheDir:     cacheDir,
		RunRoot:      runDir,
		TerraformBin: terraformBin,
		RepoRoot:     f.repoRoot,
		NoCleanup:    f.noCleanup,
	}
}
