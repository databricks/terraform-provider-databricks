package main

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"

	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/config"
	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/profile"
	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/result"
	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/runner"
	"github.com/databricks/terraform-provider-databricks/testframeworkV2/internal/terraform"
)

// runRun handles the `tfv2 run <test-dir>` subcommand (and `-r` for
// recursive runs). Returns the exit code main should hand to
// os.Exit.
func runRun(args []string) int {
	f, err := parseRunFlags(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "tfv2 run: %v\n\n", err)
		printUsage(os.Stderr)
		return exitCodeUsage
	}
	ctx, cancel := signalContext()
	defer cancel()

	if f.recursive {
		return runRecursive(ctx, f)
	}
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

// runRecursive walks f.testDir for nested test.yaml files and runs
// each one in turn. We continue past per-test failures so the user
// sees the full picture in one invocation; the aggregate exit code is
// failed if ANY test failed.
//
// Discovery rule: any directory containing a regular `test.yaml` file
// counts as a test root. Subdirectories that themselves contain a
// `test.yaml` are also picked up (separate, independent tests).
func runRecursive(ctx context.Context, f runFlags) int {
	roots, err := discoverTests(f.testDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "tfv2 run -r: %v\n", err)
		return exitCodeFailed
	}
	if len(roots) == 0 {
		fmt.Fprintf(os.Stderr, "tfv2 run -r: no test.yaml found under %s\n", f.testDir)
		return exitCodeFailed
	}
	fmt.Fprintf(os.Stdout, "Discovered %d test(s) under %s\n", len(roots), f.testDir)
	allPassed := true
	for i, dir := range roots {
		fmt.Fprintf(os.Stdout, "\n== test %d/%d: %s ==\n", i+1, len(roots), dir)
		each := f
		each.testDir = dir
		each.recursive = false
		res, err := runOnce(ctx, each)
		if err != nil {
			fmt.Fprintf(os.Stderr, "tfv2 run -r: %s: %v\n", dir, err)
			allPassed = false
			continue
		}
		printRunResult(os.Stdout, res)
		if !res.AllPassed() {
			allPassed = false
		}
	}
	if !allPassed {
		return exitCodeFailed
	}
	return exitCodeOK
}

// discoverTests walks root and returns every directory containing a
// regular `test.yaml` file, sorted by path. Hidden directories
// (starting with ".") are skipped — that excludes .git, .terraform,
// and the per-run workdirs that the framework writes elsewhere.
func discoverTests(root string) ([]string, error) {
	var out []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d.IsDir() {
			name := filepath.Base(path)
			if path != root && len(name) > 0 && name[0] == '.' {
				return filepath.SkipDir
			}
			return nil
		}
		if d.Name() != "test.yaml" {
			return nil
		}
		out = append(out, filepath.Dir(path))
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Strings(out)
	return out, nil
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
