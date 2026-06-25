// Shared inert-path classifier: the single source of truth for which changed
// files are "inert" (docs and repo metadata that can't affect an integration
// test). Imported by integration-tests.yml's `detect-changes` and
// `carry-forward` jobs and by the internal fork/manual integration-test
// workflow (which checks this file out of `main`), so the allowlist lives in
// one place. Dependency-free so every caller can import it without an install.

// A PR is inert-only when every changed file matches one of these globs.
export const INERT_GLOBS = [
  "docs/**",
  "*.md",
  ".github/ISSUE_TEMPLATE/**",
  ".github/PULL_REQUEST_TEMPLATE.md",
  ".github/dependabot.yml",
  "CODEOWNERS",
  "LICENSE",
];

const escapeRegex = (s) => s.replace(/[.*+?^${}()|[\]\\]/g, "\\$&");

// Translate the glob shapes the allowlist uses into anchored regexes. Anything
// outside that vocabulary throws, so callers fail open rather than mismatch.
export const globToRegExp = (glob) => {
  let m = glob.match(/^([^*?{}\[\]]+)\/\*\*$/); // "<dir>/**"
  if (m) return new RegExp("^" + escapeRegex(m[1]) + "/");
  m = glob.match(/^\*(\.[A-Za-z0-9.]+)$/); // root "*.<ext>"
  if (m) return new RegExp("^[^/]*" + escapeRegex(m[1]) + "$");
  if (!/[*?{}\[\]]/.test(glob)) return new RegExp("^" + escapeRegex(glob) + "$"); // literal
  throw new Error(`Unsupported glob: ${glob}`);
};

const INERT_MATCHERS = INERT_GLOBS.map(globToRegExp);

export const isInert = (path) => INERT_MATCHERS.some((re) => re.test(path));

// Empty list is non-inert: with nothing to classify we can't prove it's safe to skip.
export const allInert = (paths) => paths.length > 0 && paths.every(isInert);
