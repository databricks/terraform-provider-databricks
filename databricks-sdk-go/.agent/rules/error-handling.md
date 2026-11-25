---
description: "Error handling conventions for the Databricks SDK."
globs: "**/*.go"
alwaysApply: false
---

# Error Handling

- Wrap errors with **caller's context**: `fmt.Errorf("resolving workspace host: %w", err)`.
  Describe what *you* were doing, not what the callee does
- **Add context at the caller, not in utility functions** — the caller knows the
  operation's purpose; a utility function like `atomicWriteFile` should not prefix its
  errors when the caller can provide better context like `"storing token in local cache"`
- Don't add redundant wrapping — if the underlying error already contains the path/resource,
  don't repeat it: `write /home/user/.databricks/.tmp: permission denied` is sufficient
  without a `"write temp file: "` prefix
- Include the specific resource/URL that failed, not just the operation name
- Define sentinel errors as package-level `var` using `errors.New()` — never `fmt.Errorf`.
  They must be comparable with `errors.Is()`
- **Unexpected cases must error**: a `default` case in a `switch` should return an error
  (e.g., `fmt.Errorf("unknown host type: %v", h)`), not silently succeed. If silence is
  intentional, add a comment explaining why
- **Error messages should be actionable**: include the specific field or value that is
  wrong so the user knows what to fix (e.g., `"azure_client_id is required"`)
- Use `%q` (not `%s`) for strings in error/log messages — makes empty strings visible as `""`
