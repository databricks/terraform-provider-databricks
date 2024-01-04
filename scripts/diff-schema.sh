#!/usr/bin/env bash

source scripts/libschema.sh

BASE=${BASE_COMMIT:-"master"}
CURRENT_BRANCH=$(git rev-parse --abbrev-ref $HEAD)
HEAD=${HEAD_COMMIT:-$CURRENT_BRANCH}

checkout_branch() {
    local branch=$1
    echo "Checking out branch: $branch"
    git checkout $branch
}

if [ -n "$(git status --porcelain)" ]; then
    echo "There are uncommitted changes. Please commit them before running this script."
    exit 1
fi

NEW_SCHEMA=$(generate_schema)
checkout_branch $BASE
CURRENT_SCHEMA=$(generate_schema)
checkout_branch $HEAD

set +e
jd -color "$CURRENT_SCHEMA" "$NEW_SCHEMA"
RES=$?
set -e
if [ $RES -eq 0 ]; then
    echo "No schema changes detected."
fi