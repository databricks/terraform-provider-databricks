// Shared inert-path classifier — the single source of truth for which changed
// files are "inert" (documentation and repository metadata that can never
// affect an integration-test outcome). It is consumed in three places:
//
//   1. integration-tests.yml `detect-changes` (this repo, via github-script)
//   2. integration-tests.yml `carry-forward`  (this repo, via github-script)
//   3. The internal fork-PR / manual-dispatch integration-test workflow, which
//      sparse-checks this exact file out of `main` and imports it so those runs
//      share the same allowlist.
//
// Keeping the allowlist and the matching logic here — and nowhere else — stops
// the call sites from drifting apart. The file is intentionally dependency-free
// (pure Node, no npm) so every caller can run it without an install step.

import { readFileSync } from "node:fs";
import { pathToFileURL } from "node:url";

// Inert allowlist. A PR is inert-only when every changed file matches one of
// these globs. Globs use the small vocabulary understood by globToRegExp below.
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

// Translate the three glob shapes the allowlist uses into anchored regexes.
// Anything outside that vocabulary throws so callers can fail open (treat the
// change as non-inert and run the full suite) rather than silently mismatch.
export const globToRegExp = (glob) => {
  let m = glob.match(/^([^*?{}\[\]]+)\/\*\*$/); // "<dir>/**"
  if (m) return new RegExp("^" + escapeRegex(m[1]) + "/");
  m = glob.match(/^\*(\.[A-Za-z0-9.]+)$/); // root "*.<ext>"
  if (m) return new RegExp("^[^/]*" + escapeRegex(m[1]) + "$");
  if (!/[*?{}\[\]]/.test(glob)) return new RegExp("^" + escapeRegex(glob) + "$"); // literal
  throw new Error(`Unsupported glob: ${glob}`);
};

const INERT_MATCHERS = INERT_GLOBS.map(globToRegExp);

// True when `path` matches any inert glob.
export const isInert = (path) => INERT_MATCHERS.some((re) => re.test(path));

// True only when there is at least one path and every path is inert. An empty
// list is treated as non-inert: with nothing to classify we cannot prove the
// change is safe to skip.
export const allInert = (paths) => paths.length > 0 && paths.every(isInert);

// CLI entrypoint for non-JS callers and tests:
//   node classify_inert.mjs <changed_files.txt>
//   node classify_inert.mjs            # reads the list from stdin
// Reads a newline-delimited file list and prints exactly `inert_only=true` or
// `inert_only=false`. A parse error (e.g. an unsupported glob) throws and exits
// non-zero so the caller can fail open.
function main(argv) {
  const arg = argv[2];
  const raw = readFileSync(arg ?? 0, "utf8");
  const paths = raw
    .split("\n")
    .map((line) => line.trim())
    .filter(Boolean);
  process.stdout.write(`inert_only=${allInert(paths)}\n`);
}

if (import.meta.url === pathToFileURL(process.argv[1] ?? "").href) {
  main(process.argv);
}
