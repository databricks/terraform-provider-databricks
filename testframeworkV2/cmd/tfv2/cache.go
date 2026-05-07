package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// runCache dispatches `tfv2 cache <list|prune>`. Sub-actions are kept
// in this single file because they share the same flag plumbing.
func runCache(args []string) int {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "tfv2 cache: missing action (list|prune)")
		return exitCodeUsage
	}
	switch args[0] {
	case "list":
		return runCacheList(args[1:])
	case "prune":
		return runCachePrune(args[1:])
	default:
		fmt.Fprintf(os.Stderr, "tfv2 cache: unknown action %q (expected list|prune)\n", args[0])
		return exitCodeUsage
	}
}

// cacheCommonFlags is shared by `cache list` and `cache prune`. Lifts
// --cache-dir + TFV2_CACHE_DIR resolution out of the action funcs.
type cacheCommonFlags struct {
	cacheDir string
}

func parseCacheCommon(name string, args []string) (cacheCommonFlags, error) {
	fs := flag.NewFlagSet(name, flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	var c cacheCommonFlags
	fs.StringVar(&c.cacheDir, "cache-dir", os.Getenv("TFV2_CACHE_DIR"), "override ~/.testframeworkv2/providers")
	if err := fs.Parse(args); err != nil {
		return cacheCommonFlags{}, err
	}
	if fs.NArg() != 0 {
		return cacheCommonFlags{}, fmt.Errorf("unexpected positional argument %q", fs.Arg(0))
	}
	if c.cacheDir == "" {
		c.cacheDir = defaultCacheDir()
	}
	return c, nil
}

// runCacheList walks the cache root and prints one line per cached
// version. Output shape:
//
//	1.113.0 darwin_arm64 packed   /Users/.../.zip   65 MB
//	1.114.0 darwin_arm64 packed   /Users/.../.zip   65 MB
//	99.0.0-local darwin_arm64 unpacked  /Users/...  62 MB
//
// We deliberately don't pretty-print into columns — `column -t` is
// the user's friend and avoids us tracking column widths.
func runCacheList(args []string) int {
	c, err := parseCacheCommon("cache list", args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "tfv2 cache list: %v\n", err)
		return exitCodeUsage
	}
	entries, err := scanCache(c.cacheDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "tfv2 cache list: %v\n", err)
		return exitCodeFailed
	}
	if len(entries) == 0 {
		fmt.Fprintf(os.Stdout, "(empty cache at %s)\n", c.cacheDir)
		return exitCodeOK
	}
	for _, e := range entries {
		fmt.Fprintf(os.Stdout, "%s %s %s %s %s\n",
			e.Version, e.Target, e.Layout, e.Path, formatBytes(e.SizeBytes))
	}
	return exitCodeOK
}

// runCachePrune removes the entire cache tree. Idempotent: a missing
// cache directory exits OK.
func runCachePrune(args []string) int {
	c, err := parseCacheCommon("cache prune", args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "tfv2 cache prune: %v\n", err)
		return exitCodeUsage
	}
	if _, err := os.Stat(c.cacheDir); os.IsNotExist(err) {
		fmt.Fprintf(os.Stdout, "cache already empty: %s\n", c.cacheDir)
		return exitCodeOK
	}
	if err := os.RemoveAll(c.cacheDir); err != nil {
		fmt.Fprintf(os.Stderr, "tfv2 cache prune: %v\n", err)
		return exitCodeFailed
	}
	fmt.Fprintf(os.Stdout, "pruned: %s\n", c.cacheDir)
	return exitCodeOK
}

// CacheEntry is one cached provider build. Layout is "packed" for
// released-version zips and "unpacked" for local builds. Exported
// from the test-only callable scanCache so tests can assert on
// scanCache output without pulling in the printer.
type CacheEntry struct {
	Version   string
	Target    string
	Layout    string
	Path      string
	SizeBytes int64
}

// scanCache discovers every cached provider build under root. The
// scan is filesystem-only — no validation that the binary is a real
// terraform provider, since corruption would surface at terraform
// init time anyway.
func scanCache(root string) ([]CacheEntry, error) {
	provDir := filepath.Join(root, "registry.terraform.io", "databricks", "databricks")
	if _, err := os.Stat(provDir); os.IsNotExist(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	var out []CacheEntry
	err := filepath.WalkDir(provDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		entry, ok := classifyCacheEntry(provDir, path)
		if !ok {
			return nil
		}
		info, err := d.Info()
		if err != nil {
			return nil
		}
		entry.SizeBytes = info.Size()
		out = append(out, entry)
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Version != out[j].Version {
			return out[i].Version < out[j].Version
		}
		return out[i].Target < out[j].Target
	})
	return out, nil
}

// classifyCacheEntry returns a CacheEntry for path if it matches one
// of the expected layouts (packed zip OR unpacked binary), and false
// otherwise. The provenance JSON files we wrote ourselves are
// deliberately filtered out.
func classifyCacheEntry(provDir, path string) (CacheEntry, bool) {
	rel, err := filepath.Rel(provDir, path)
	if err != nil {
		return CacheEntry{}, false
	}
	parts := strings.Split(filepath.ToSlash(rel), "/")
	switch {
	case len(parts) == 1 && strings.HasPrefix(parts[0], "terraform-provider-databricks_") && strings.HasSuffix(parts[0], ".zip"):
		// packed: terraform-provider-databricks_<version>_<target>.zip
		v, t, ok := parsePackedZipName(parts[0])
		if !ok {
			return CacheEntry{}, false
		}
		return CacheEntry{Version: v, Target: t, Layout: "packed", Path: path}, true
	case len(parts) == 3 && strings.HasPrefix(parts[2], "terraform-provider-databricks_v"):
		// unpacked: <version>/<target>/terraform-provider-databricks_v<version>
		return CacheEntry{Version: parts[0], Target: parts[1], Layout: "unpacked", Path: path}, true
	}
	return CacheEntry{}, false
}

// parsePackedZipName extracts (version, target) from a packed zip
// basename. Returns ("", "", false) for any unexpected shape.
func parsePackedZipName(name string) (version, target string, ok bool) {
	const prefix = "terraform-provider-databricks_"
	const suffix = ".zip"
	if !strings.HasPrefix(name, prefix) || !strings.HasSuffix(name, suffix) {
		return "", "", false
	}
	body := strings.TrimSuffix(strings.TrimPrefix(name, prefix), suffix)
	// body is "<version>_<os>_<arch>". The version itself can contain
	// "-prerelease" but not "_", so the split-on-underscore trick
	// works as long as we take the LAST two underscore-separated
	// fields as the target.
	idx := strings.LastIndex(body, "_")
	if idx < 0 {
		return "", "", false
	}
	rest := body[:idx]
	arch := body[idx+1:]
	idx = strings.LastIndex(rest, "_")
	if idx < 0 {
		return "", "", false
	}
	version = rest[:idx]
	osName := rest[idx+1:]
	return version, osName + "_" + arch, true
}

// formatBytes returns a short human-friendly size string.
func formatBytes(n int64) string {
	const (
		kb = 1024
		mb = 1024 * kb
		gb = 1024 * mb
	)
	switch {
	case n >= gb:
		return fmt.Sprintf("%.1f GB", float64(n)/float64(gb))
	case n >= mb:
		return fmt.Sprintf("%d MB", n/mb)
	case n >= kb:
		return fmt.Sprintf("%d KB", n/kb)
	default:
		return fmt.Sprintf("%d B", n)
	}
}
