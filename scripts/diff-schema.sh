#!/usr/bin/env bash
set -eo pipefail

source scripts/libschema.sh

BASE=${BASE_COMMIT:-"main"}
# Capture the actual SHA so we can return here reliably. `git rev-parse --abbrev-ref HEAD`
# returns the literal string "HEAD" when we're in a detached-HEAD state (as we are during
# GitHub Actions PR checkouts), which makes the final restore-checkout a silent no-op.
HEAD_SHA=$(git rev-parse HEAD)

checkout() {
    local ref=$1
    echo "Checking out ref: $ref"
    git checkout $ref
}

if [ -n "$(git status --porcelain)" ]; then
    echo "There are uncommitted changes. Please commit them before running this script."
    exit 1
fi

# Build the classifier from HEAD *before* any base/head checkouts, so the build is
# guaranteed to see the HEAD source tree. The base may not contain the classifier
# directory (e.g., the very PR that introduces it), and the final restore-checkout
# below also depends on the build artifact already existing in /tmp.
# -mod=mod skips the vendor consistency check; the classifier itself is stdlib-only.
go build -mod=mod -o /tmp/schema-classifier ./scripts/schema-classifier

NEW_SCHEMA=$(generate_schema)
checkout $BASE
CURRENT_SCHEMA=$(generate_schema)
checkout $HEAD_SHA

set +e
jd -color "$CURRENT_SCHEMA" "$NEW_SCHEMA"
RES=$?
set -e
if [ $RES -eq 0 ]; then
    echo "No schema changes detected."
fi

# Classify changes as breaking vs non-breaking. The classifier exits non-zero on any
# breaking change unless ALLOW_BREAKING=1 is set in the environment (for local dev or
# when the PR carries the ALLOW_SCHEMA_BREAKING_CHANGE=true bypass directive).
echo
echo "===== Breaking-change classification ====="

CLASSIFIER_FLAGS=()
if [ "${ALLOW_BREAKING:-0}" = "1" ] || [ "${ALLOW_BREAKING:-}" = "true" ]; then
    CLASSIFIER_FLAGS+=(--allow-breaking)
fi

# Always print the human-readable text view to the log. Use the `|| CLASSIFIER_EXIT=$?`
# pattern so `set -e` (at the top of this script) doesn't kill us when the classifier
# legitimately exits 1 on breaking changes.
CLASSIFIER_EXIT=0
/tmp/schema-classifier --base "$CURRENT_SCHEMA" --head "$NEW_SCHEMA" "${CLASSIFIER_FLAGS[@]}" || CLASSIFIER_EXIT=$?

exit $CLASSIFIER_EXIT
