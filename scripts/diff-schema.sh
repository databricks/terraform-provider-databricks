#!/usr/bin/env bash
set -eo pipefail

source scripts/libschema.sh

BASE=${BASE_COMMIT:-"main"}
CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)

checkout() {
    local ref=$1
    echo "Checking out ref: $ref"
    git checkout $ref
}

if [ -n "$(git status --porcelain)" ]; then
    echo "There are uncommitted changes. Please commit them before running this script."
    exit 1
fi

NEW_SCHEMA=$(generate_schema)
checkout $BASE
CURRENT_SCHEMA=$(generate_schema)
checkout $CURRENT_BRANCH

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
# -mod=mod skips the vendor consistency check. We need this because this script does
# base/head checkouts in sequence, and the vendor/ directory may match neither at the
# point we build the classifier. The classifier itself is stdlib-only.
go build -mod=mod -o /tmp/schema-classifier ./scripts/schema-classifier

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
