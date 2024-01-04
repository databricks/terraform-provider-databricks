#!/usr/bin/env bash
set -eo pipefail

source scripts/libschema.sh

BASE=${BASE_COMMIT:-"master"}
CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
HEAD=${HEAD_COMMIT:-$CURRENT_BRANCH}

checkout() {
    local ref=$1
    echo "Checking out ref: $ref"
    git checkout $ref
}

if [ -n "$(git status --porcelain)" ]; then
    echo "There are uncommitted changes. Please commit them before running this script."
    exit 1
fi

checkout $HEAD
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