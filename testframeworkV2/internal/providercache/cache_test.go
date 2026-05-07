package providercache

import (
	"bytes"
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
)

const testVersion = "1.114.0"

func testTarget() Target { return Target{OS: "darwin", Arch: "arm64"} }

// fakeZipContent returns deterministic-per-call random bytes that stand in
// for a real provider zip. providercache never opens the archive in M1, so
// the bytes only need to round-trip from server to disk faithfully.
func fakeZipContent(t *testing.T) []byte {
	t.Helper()
	b := make([]byte, 4096)
	if _, err := rand.Read(b); err != nil {
		t.Fatalf("rand.Read: %v", err)
	}
	return b
}

// releasePath returns the expected URL path under the GitHub releases prefix
// for a given (version, target).
func releasePath(version, target string) string {
	return fmt.Sprintf("/v%s/terraform-provider-databricks_%s_%s.zip", version, version, target)
}

// expectedZipPath is where the cache should land a successful (version,
// target) download under c.Root().
func expectedZipPath(c *Cache, version, target string) string {
	return filepath.Join(
		c.Root(),
		"registry.terraform.io", "databricks", "databricks",
		fmt.Sprintf("terraform-provider-databricks_%s_%s.zip", version, target),
	)
}

// servingHandler returns a handler that serves body at the canonical release
// path for (version, target) and t.Errorf-fails any unrelated request.
// hits, when non-nil, counts successful serves.
func servingHandler(t *testing.T, version, target string, body []byte, hits *int32) http.Handler {
	t.Helper()
	mux := http.NewServeMux()
	mux.HandleFunc(releasePath(version, target), func(w http.ResponseWriter, r *http.Request) {
		if hits != nil {
			atomic.AddInt32(hits, 1)
		}
		w.Header().Set("Content-Type", "application/zip")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(body)
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
		http.Error(w, "not found", http.StatusNotFound)
	})
	return mux
}

func newCacheServer(t *testing.T, h http.Handler) *Cache {
	t.Helper()
	s := httptest.NewServer(h)
	t.Cleanup(s.Close)
	return New(t.TempDir(), WithBaseURL(s.URL))
}

// TestTarget_String pins the canonical "<os>_<arch>" formatting that the
// rest of the framework, .terraformrc generation, and the cache layout all
// depend on.
func TestTarget_String(t *testing.T) {
	if got, want := (Target{OS: "darwin", Arch: "arm64"}).String(), "darwin_arm64"; got != want {
		t.Errorf("Target.String: got %q want %q", got, want)
	}
}

// TestHostTarget_MatchesRuntime ensures HostTarget tracks runtime.GOOS /
// runtime.GOARCH. This is light insurance against accidentally hardcoding a
// fixed target during refactors.
func TestHostTarget_MatchesRuntime(t *testing.T) {
	got := HostTarget()
	if got.OS != runtime.GOOS || got.Arch != runtime.GOARCH {
		t.Errorf("HostTarget = %s, want %s_%s", got, runtime.GOOS, runtime.GOARCH)
	}
}

// TestResolve_DownloadsAndCaches covers the M1 happy path: a fresh cache
// downloads the asset on the first Resolve, and a second Resolve at the same
// (version, target) returns the cached path without re-fetching.
func TestResolve_DownloadsAndCaches(t *testing.T) {
	expected := fakeZipContent(t)
	var hits int32
	c := newCacheServer(t, servingHandler(t, testVersion, "darwin_arm64", expected, &hits))

	// First call: triggers a download.
	gotPath, gotVer, err := c.Resolve(context.Background(), testVersion, testTarget())
	if err != nil {
		t.Fatalf("Resolve: %v", err)
	}
	if gotVer != testVersion {
		t.Errorf("synthetic version: got %q want %q", gotVer, testVersion)
	}
	if want := expectedZipPath(c, testVersion, "darwin_arm64"); gotPath != want {
		t.Errorf("path mismatch: got %q want %q", gotPath, want)
	}
	got, err := os.ReadFile(gotPath)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	if !bytes.Equal(got, expected) {
		t.Errorf("downloaded bytes do not match (len got=%d want=%d)", len(got), len(expected))
	}

	// Second call: cache hit, no additional HTTP request.
	gotPath2, _, err := c.Resolve(context.Background(), testVersion, testTarget())
	if err != nil {
		t.Fatalf("Resolve(cached): %v", err)
	}
	if gotPath2 != gotPath {
		t.Errorf("cached path mismatch: got %q want %q", gotPath2, gotPath)
	}
	if h := atomic.LoadInt32(&hits); h != 1 {
		t.Errorf("expected 1 HTTP hit total, got %d", h)
	}
}

// TestResolve_CacheLayoutPacked pins the on-disk layout against DESIGN.md
// §6 ("Cache layout — Released versions"). Breaking the layout breaks
// Terraform's filesystem_mirror discovery.
func TestResolve_CacheLayoutPacked(t *testing.T) {
	body := fakeZipContent(t)
	c := newCacheServer(t, servingHandler(t, testVersion, "linux_amd64", body, nil))

	gotPath, _, err := c.Resolve(context.Background(), testVersion, Target{OS: "linux", Arch: "amd64"})
	if err != nil {
		t.Fatalf("Resolve: %v", err)
	}

	want := filepath.Join(
		c.Root(),
		"registry.terraform.io", "databricks", "databricks",
		"terraform-provider-databricks_1.114.0_linux_amd64.zip",
	)
	if gotPath != want {
		t.Errorf("packed mirror path:\n got  %q\n want %q", gotPath, want)
	}

	// No spurious .partial.* siblings should remain in the provider dir.
	entries, err := os.ReadDir(filepath.Dir(want))
	if err != nil {
		t.Fatalf("ReadDir: %v", err)
	}
	for _, e := range entries {
		if strings.Contains(e.Name(), ".partial") {
			t.Errorf("unexpected leftover partial file: %s", e.Name())
		}
	}
}

// TestResolve_NormalizesLeadingV verifies that "v1.114.0" and "1.114.0"
// resolve to the same on-disk path and synthetic version. GitHub release tags
// use the v-prefix while Terraform's filesystem_mirror filenames don't.
func TestResolve_NormalizesLeadingV(t *testing.T) {
	body := fakeZipContent(t)
	c := newCacheServer(t, servingHandler(t, testVersion, "darwin_arm64", body, nil))

	gotPath, gotVer, err := c.Resolve(context.Background(), "v"+testVersion, testTarget())
	if err != nil {
		t.Fatalf("Resolve: %v", err)
	}
	if gotVer != testVersion {
		t.Errorf("synthetic version: got %q want %q", gotVer, testVersion)
	}
	if !strings.HasSuffix(gotPath, "_1.114.0_darwin_arm64.zip") {
		t.Errorf("path should normalize leading v: got %q", gotPath)
	}
}

// TestResolve_NotFound_ReturnsError exercises the 404 path: the GitHub
// release does not exist (typoed version, non-existent target). The cache
// must surface a clear error and leave nothing on disk.
func TestResolve_NotFound_ReturnsError(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})
	s := httptest.NewServer(mux)
	t.Cleanup(s.Close)
	c := New(t.TempDir(), WithBaseURL(s.URL))

	_, _, err := c.Resolve(context.Background(), "9.9.9", testTarget())
	if err == nil {
		t.Fatal("Resolve: expected error, got nil")
	}
	if !strings.Contains(err.Error(), "HTTP 404") {
		t.Errorf("error should mention HTTP 404: %v", err)
	}

	// Cache directory should be empty (no .zip and no .partial leftover).
	zip := expectedZipPath(c, "9.9.9", "darwin_arm64")
	if _, err := os.Stat(zip); !errors.Is(err, os.ErrNotExist) {
		t.Errorf("expected no zip on disk, got err=%v", err)
	}
	if entries, _ := os.ReadDir(filepath.Dir(zip)); len(entries) != 0 {
		var names []string
		for _, e := range entries {
			names = append(names, e.Name())
		}
		t.Errorf("expected empty provider dir, got entries: %v", names)
	}
}

// TestResolve_LocalVersion_NotImplemented locks in the M1 behaviour: a step
// asking for version=local must fail with a clear, actionable error pointing
// to the milestone that lands the feature, not a confusing 404 from
// "GET .../vlocal/...".
func TestResolve_LocalVersion_NotImplemented(t *testing.T) {
	c := New(t.TempDir())
	_, _, err := c.Resolve(context.Background(), LocalVersionInput, testTarget())
	if err == nil {
		t.Fatal("Resolve(local): expected error, got nil")
	}
	if !strings.Contains(err.Error(), "M6") {
		t.Errorf("error should mention M6 milestone: %v", err)
	}
}

// TestResolve_RejectsBadInputs ensures the cache fails fast on invalid
// arguments rather than emitting a malformed URL or path.
func TestResolve_RejectsBadInputs(t *testing.T) {
	c := New(t.TempDir())

	for _, tc := range []struct {
		name    string
		version string
		target  Target
		wantSub string
	}{
		{"empty version", "", testTarget(), "version is required"},
		{"missing arch", testVersion, Target{OS: "darwin"}, "OS and Arch"},
		{"missing OS", testVersion, Target{Arch: "arm64"}, "OS and Arch"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, _, err := c.Resolve(context.Background(), tc.version, tc.target)
			if err == nil {
				t.Fatal("expected error, got nil")
			}
			if !strings.Contains(err.Error(), tc.wantSub) {
				t.Errorf("error %q does not contain %q", err.Error(), tc.wantSub)
			}
		})
	}
}

// TestResolve_TruncatedStream_CleansUpPartial covers the crash-safety
// invariant from DESIGN.md §6: a broken transfer must leave no
// half-written zip and no leftover .partial.* artifacts.
func TestResolve_TruncatedStream_CleansUpPartial(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc(releasePath(testVersion, "darwin_arm64"), func(w http.ResponseWriter, r *http.Request) {
		// Lie about Content-Length so io.Copy treats a short read as an
		// error rather than a clean EOF.
		w.Header().Set("Content-Length", "4096")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("only-a-header"))

		hj, ok := w.(http.Hijacker)
		if !ok {
			t.Errorf("ResponseWriter does not support http.Hijacker")
			return
		}
		conn, _, err := hj.Hijack()
		if err != nil {
			t.Errorf("Hijack: %v", err)
			return
		}
		_ = conn.Close()
	})
	s := httptest.NewServer(mux)
	t.Cleanup(s.Close)
	c := New(t.TempDir(), WithBaseURL(s.URL))

	_, _, err := c.Resolve(context.Background(), testVersion, testTarget())
	if err == nil {
		t.Fatal("Resolve: expected error from truncated stream, got nil")
	}

	zip := expectedZipPath(c, testVersion, "darwin_arm64")
	if _, err := os.Stat(zip); !errors.Is(err, os.ErrNotExist) {
		t.Errorf("expected no zip on disk, got err=%v", err)
	}
	if entries, _ := os.ReadDir(filepath.Dir(zip)); len(entries) != 0 {
		var names []string
		for _, e := range entries {
			names = append(names, e.Name())
		}
		t.Errorf("expected empty provider dir after truncated download, got entries: %v", names)
	}
}

// TestResolve_ContextCanceled ensures Resolve respects the caller's context.
// We pre-cancel the context so we don't depend on a real network round-trip.
func TestResolve_ContextCanceled(t *testing.T) {
	c := New(t.TempDir(), WithBaseURL("http://127.0.0.1:1")) // unreachable
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, _, err := c.Resolve(ctx, testVersion, testTarget())
	if err == nil {
		t.Fatal("Resolve: expected error from canceled context, got nil")
	}
}

// TestResolve_Concurrent exercises the per-writer .partial discipline. With
// N goroutines racing on the same uncached (version, target), every Resolve
// must succeed and the cache must end with a single zip file containing the
// correct bytes — no leftover .partial.* entries.
func TestResolve_Concurrent(t *testing.T) {
	expected := fakeZipContent(t)
	var hits int32
	mux := http.NewServeMux()
	mux.HandleFunc(releasePath(testVersion, "darwin_arm64"), func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&hits, 1)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(expected)
	})
	s := httptest.NewServer(mux)
	t.Cleanup(s.Close)
	c := New(t.TempDir(), WithBaseURL(s.URL))

	const goroutines = 8
	var wg sync.WaitGroup
	wg.Add(goroutines)
	errs := make(chan error, goroutines)
	for range goroutines {
		go func() {
			defer wg.Done()
			_, _, err := c.Resolve(context.Background(), testVersion, testTarget())
			errs <- err
		}()
	}
	wg.Wait()
	close(errs)
	for err := range errs {
		if err != nil {
			t.Errorf("concurrent Resolve: %v", err)
		}
	}

	// Final state: exactly one zip, no .partial.* siblings, correct bytes.
	zip := expectedZipPath(c, testVersion, "darwin_arm64")
	got, err := os.ReadFile(zip)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	if !bytes.Equal(got, expected) {
		t.Errorf("cached bytes do not match expected (len got=%d want=%d)", len(got), len(expected))
	}
	entries, err := os.ReadDir(filepath.Dir(zip))
	if err != nil {
		t.Fatalf("ReadDir: %v", err)
	}
	zipCount := 0
	for _, e := range entries {
		if strings.Contains(e.Name(), ".partial") {
			t.Errorf("unexpected leftover partial: %s", e.Name())
		}
		if filepath.Ext(e.Name()) == ".zip" {
			zipCount++
		}
	}
	if zipCount != 1 {
		t.Errorf("expected exactly 1 cached zip, got %d", zipCount)
	}
	if h := atomic.LoadInt32(&hits); h < 1 {
		t.Errorf("expected at least 1 HTTP hit, got %d", h)
	}
}

// TestResolve_ReleaseURL pins the URL shape against DESIGN.md task
// description (the GitHub releases canonical asset URL). A regression here
// would silently start hitting the wrong endpoint.
func TestResolve_ReleaseURL(t *testing.T) {
	var capturedURL string
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		capturedURL = r.URL.Path
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(fakeZipContent(t))
	})
	s := httptest.NewServer(mux)
	t.Cleanup(s.Close)
	c := New(t.TempDir(), WithBaseURL(s.URL))

	_, _, err := c.Resolve(context.Background(), testVersion, Target{OS: "linux", Arch: "amd64"})
	if err != nil {
		t.Fatalf("Resolve: %v", err)
	}
	want := "/v1.114.0/terraform-provider-databricks_1.114.0_linux_amd64.zip"
	if capturedURL != want {
		t.Errorf("URL path: got %q want %q", capturedURL, want)
	}
}

// TestResolve_BaseURLTrailingSlash documents that the WithBaseURL option
// strips trailing slashes — otherwise a "//" sneaks into the GET URL and
// a strict server (or a future signature scheme) would reject it.
func TestResolve_BaseURLTrailingSlash(t *testing.T) {
	body := fakeZipContent(t)
	mux := http.NewServeMux()
	hit := false
	mux.HandleFunc(releasePath(testVersion, "darwin_arm64"), func(w http.ResponseWriter, r *http.Request) {
		hit = true
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(body)
	})
	s := httptest.NewServer(mux)
	t.Cleanup(s.Close)
	c := New(t.TempDir(), WithBaseURL(s.URL+"/"))

	if _, _, err := c.Resolve(context.Background(), testVersion, testTarget()); err != nil {
		t.Fatalf("Resolve: %v", err)
	}
	if !hit {
		t.Error("expected handler to be hit at canonical release path")
	}
}

// TestResolve_HTTPClientOverride confirms that callers may inject a custom
// *http.Client via WithHTTPClient — important for the future addition of a
// retrying or rate-limited client.
func TestResolve_HTTPClientOverride(t *testing.T) {
	body := fakeZipContent(t)
	c := newCacheServer(t, servingHandler(t, testVersion, "darwin_arm64", body, nil))
	// Replace the default client with one whose Transport always returns 503.
	custom := &http.Client{Transport: roundTripperFunc(func(_ *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusServiceUnavailable,
			Body:       io.NopCloser(strings.NewReader("nope")),
			Header:     http.Header{},
		}, nil
	})}
	WithHTTPClient(custom)(c)

	_, _, err := c.Resolve(context.Background(), testVersion, testTarget())
	if err == nil {
		t.Fatal("expected error from injected 503 transport, got nil")
	}
	if !strings.Contains(err.Error(), "HTTP 503") {
		t.Errorf("expected HTTP 503 in error, got: %v", err)
	}
}

// roundTripperFunc adapts a function to http.RoundTripper for the
// HTTPClientOverride test.
type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }
