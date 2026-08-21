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

# require_valid_schema bails out cleanly when `terraform providers schema -json`
# produced an empty or non-JSON file (typically because the provider panicked
# during schema introspection). Without this, we'd run the classifier against a
# garbage schema and report misleading "breaking changes detected" results.
#
# When CLASSIFIER_REPORT is set, writes a markdown explanation there so the CI
# status job can post it as the sticky comment. When CLASSIFIER_GENERATION_ERROR_MARKER
# is set, touches that file so the status job can distinguish this case from a
# real breaking change (the ALLOW_SCHEMA_BREAKING_CHANGE bypass should NOT apply
# to generation failures — they're provider bugs, not intentional breaks).
require_valid_schema() {
    local side="$1"
    local path="$2"
    if [ -s "$path" ] && jq empty "$path" >/dev/null 2>&1; then
        return 0
    fi
    echo >&2
    echo "ERROR: \`terraform providers schema -json\` did not produce a valid schema for $side." >&2
    echo "       Path \"$path\" is empty or non-JSON. Look earlier in this log for a provider panic." >&2
    if [ -n "${CLASSIFIER_REPORT:-}" ]; then
        cat > "$CLASSIFIER_REPORT" <<EOF
## :warning: Schema generation failed on \`$side\`

The breaking-schema check couldn't compare schemas because \`terraform providers schema -json\` failed for \`$side\`. Typically this means the provider **panicked** during schema introspection — look for a \`panic:\` line in the **classify** job's log.

Common causes:
- A new Plugin Framework resource / data source whose Go struct embeds another struct with complex (\`types.Object\` / \`types.List\` / \`types.Set\` / \`types.Map\`) fields, but the parent's \`GetComplexFieldTypes()\` doesn't include them.
- Malformed \`tfsdk\` struct tags.
- A panic in schema-customization code.

**The \`ALLOW_SCHEMA_BREAKING_CHANGE\` bypass does NOT apply here** — this is a provider bug, not a breaking schema change. Fix the panic and push a new commit.
EOF
    fi
    if [ -n "${CLASSIFIER_GENERATION_ERROR_MARKER:-}" ]; then
        : > "$CLASSIFIER_GENERATION_ERROR_MARKER"
    fi
    # Best-effort restore so local dev doesn't end up stuck on the base ref.
    git checkout "$HEAD_SHA" >/dev/null 2>&1 || true
    exit 3
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
require_valid_schema "HEAD" "$NEW_SCHEMA"
checkout $BASE
CURRENT_SCHEMA=$(generate_schema)
require_valid_schema "BASE" "$CURRENT_SCHEMA"
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

# When CLASSIFIER_REPORT is set (CI uses this so the comment job can pick it up
# via an artifact), also write the markdown view to that path. --allow-breaking
# on this second invocation so the exit code we already captured above is the
# one that matters.
if [ -n "${CLASSIFIER_REPORT:-}" ]; then
    /tmp/schema-classifier --base "$CURRENT_SCHEMA" --head "$NEW_SCHEMA" \
        --format markdown --allow-breaking > "$CLASSIFIER_REPORT" || true
fi

exit $CLASSIFIER_EXIT
