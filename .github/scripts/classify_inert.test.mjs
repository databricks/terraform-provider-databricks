// Tests for the shared inert-path classifier.
// Run with: node --test .github/scripts/classify_inert.test.mjs
import { test } from "node:test";
import assert from "node:assert/strict";

import { isInert, allInert, globToRegExp, INERT_GLOBS } from "./classify_inert.mjs";

test("docs-only change is inert", () => {
  assert.equal(allInert(["docs/index.md", "docs/resources/job.md"]), true);
});

test("mixed docs + go is not inert", () => {
  assert.equal(allInert(["docs/index.md", "internal/service/jobs.go"]), false);
});

test("root CHANGELOG.md is inert", () => {
  assert.equal(isInert("CHANGELOG.md"), true);
  assert.equal(allInert(["CHANGELOG.md"]), true);
});

test("internal go file is not inert", () => {
  assert.equal(isInert("internal/foo.go"), false);
});

test(".github/dependabot.yml is inert", () => {
  assert.equal(isInert(".github/dependabot.yml"), true);
});

test("empty list is not inert", () => {
  assert.equal(allInert([]), false);
});

test("a path a naive prefix matcher would mis-accept is not inert", () => {
  // "docs/**" must require the directory boundary: "docsrc/..." and a file
  // literally named "docs" must NOT match. A naive startsWith("docs") would.
  assert.equal(isInert("docsrc/secret.go"), false);
  assert.equal(isInert("docs"), false);
  // root "*.md" must not match markdown nested in a non-docs subtree, and must
  // not match a file that merely contains ".md" before another extension.
  assert.equal(isInert("internal/notes.md"), false);
  assert.equal(isInert("notes.md.go"), false);
});

test("issue templates and PR template are inert", () => {
  assert.equal(isInert(".github/ISSUE_TEMPLATE/bug.md"), true);
  assert.equal(isInert(".github/PULL_REQUEST_TEMPLATE.md"), true);
});

test("every allowlist glob compiles to a regex", () => {
  for (const glob of INERT_GLOBS) {
    assert.ok(globToRegExp(glob) instanceof RegExp, glob);
  }
});

test("unsupported glob throws", () => {
  assert.throws(() => globToRegExp("src/**/*.go"));
});

// Renames: the workflows classify both the new and old path of a renamed file
// (`[f.filename, f.previous_filename]`) through this same `allInert`, so a file
// renamed OUT OF a code path into docs must still count as non-inert, while a
// rename that stays entirely within inert paths is inert.
test("a rename from a non-inert path is not inert", () => {
  // previous_filename "internal/x.go" -> filename "docs/x.md"
  assert.equal(allInert(["docs/x.md", "internal/x.go"]), false);
});

test("a rename within inert paths is inert", () => {
  // previous_filename "docs/old.md" -> filename "docs/new.md"
  assert.equal(allInert(["docs/new.md", "docs/old.md"]), true);
});

// The >300-file cap is a workflow-level fail-open guard (the github-script
// callers force non-inert above the cap before this module is consulted), not
// module logic. Below the cap, a large all-inert set still classifies as inert.
test("a large all-inert file set is inert", () => {
  const many = Array.from({ length: 250 }, (_, i) => `docs/page-${i}.md`);
  assert.equal(allInert(many), true);
});
