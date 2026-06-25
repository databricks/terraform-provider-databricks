// Tests for the shared inert-path classifier. Run with: node --test .github/scripts/
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
