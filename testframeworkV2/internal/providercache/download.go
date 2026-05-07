package providercache

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// downloadZip fetches the released zip for (version, target) and writes it
// to dst atomically. The work is split into two helpers — fetchAsset (GET
// and status check) and writeAtomic (temp-file plus rename) — so each piece
// stays small enough to read at a glance and so writeAtomic is reusable for
// future asset types (e.g. _SHA256SUMS in v2 / F4).
//
// See DESIGN.md §6 "Cache atomicity" and F3 (Atomic cache writes).
func (c *Cache) downloadZip(ctx context.Context, version string, target Target, dst string) error {
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return fmt.Errorf("providercache: mkdir %s: %w", filepath.Dir(dst), err)
	}
	body, err := c.fetchAsset(ctx, c.releaseURL(version, target))
	if err != nil {
		return err
	}
	defer body.Close()
	return writeAtomic(dst, body)
}

// fetchAsset performs the HTTP GET for a release URL and returns the body
// reader on a 200 response. Non-200 responses are surfaced as errors with a
// short body snippet for diagnostics — GitHub returns informative HTML on
// 404 which can help users spot a typo'd version.
func (c *Cache) fetchAsset(ctx context.Context, url string) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("providercache: build request %s: %w", url, err)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("providercache: GET %s: %w", url, err)
	}
	if resp.StatusCode != http.StatusOK {
		snippet, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
		_ = resp.Body.Close()
		return nil, fmt.Errorf("providercache: GET %s: HTTP %d: %s", url, resp.StatusCode, string(snippet))
	}
	return resp.Body, nil
}

// writeAtomic copies src to dst via a unique "<dst>.partial.<rand>" temp
// file plus fsync plus rename. The temp file is created with os.CreateTemp
// in the same directory as dst, which guarantees:
//
//   - a unique, freshly-created file we own exclusively (no two concurrent
//     writers stomp on each other's bytes), and
//   - a same-filesystem path so os.Rename is atomic.
//
// On any error path we close the temp file and remove it, so a crashed
// download leaves no .partial.* leftovers.
func writeAtomic(dst string, src io.Reader) (retErr error) {
	dir := filepath.Dir(dst)
	tmp, err := os.CreateTemp(dir, filepath.Base(dst)+".partial.*")
	if err != nil {
		return fmt.Errorf("providercache: create temp in %s: %w", dir, err)
	}
	tmpPath := tmp.Name()
	defer func() {
		if retErr != nil {
			// Close is idempotent (returns "file already closed" after the
			// happy-path Close below); ignoring the error is fine because
			// we only run this on the failure path.
			_ = tmp.Close()
			_ = os.Remove(tmpPath)
		}
	}()

	if _, err := io.Copy(tmp, src); err != nil {
		return fmt.Errorf("providercache: write %s: %w", tmpPath, err)
	}
	if err := tmp.Sync(); err != nil {
		return fmt.Errorf("providercache: fsync %s: %w", tmpPath, err)
	}
	if err := tmp.Close(); err != nil {
		return fmt.Errorf("providercache: close %s: %w", tmpPath, err)
	}
	if err := os.Rename(tmpPath, dst); err != nil {
		return fmt.Errorf("providercache: rename %s to %s: %w", tmpPath, dst, err)
	}
	return nil
}

// releaseURL constructs the GitHub releases asset URL for (version, target):
//
//	<baseURL>/v<version>/terraform-provider-databricks_<version>_<target>.zip
//
// version must already be normalized (no leading "v") — Resolve handles
// normalization before calling.
func (c *Cache) releaseURL(version string, target Target) string {
	return fmt.Sprintf(
		"%s/v%s/terraform-provider-databricks_%s_%s.zip",
		c.baseURL, version, version, target.String(),
	)
}
