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

NEW_SCHEMA=$(generate_schema | tail -n1)
checkout $BASE
CURRENT_SCHEMA=$(generate_schema | tail -n1)
checkout $CURRENT_BRANCH

set +e
jd -color "$CURRENT_SCHEMA" "$NEW_SCHEMA"
RES=$?
set -e
if [ $RES -eq 0 ]; then
    echo "No schema changes detected."
fi