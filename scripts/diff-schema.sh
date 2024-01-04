#!/usr/bin/env bash

set -x

source scripts/libschema.sh

BASE_BRANCH="master"

checkout_branch() {
    local branch=$1
    echo "Checking out branch: $branch"
    git checkout $branch
}

if [ -n "$(git status --porcelain)" ]; then
    echo "There are uncommitted changes. Please commit them before running this script."
    exit 1
fi

CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
NEW_SCHEMA=$(generate_schema)
checkout_branch $BASE_BRANCH
CURRENT_SCHEMA=$(generate_schema)
checkout_branch $CURRENT_BRANCH

set +e
jd -color "$CURRENT_SCHEMA" "$NEW_SCHEMA"
RES=$?
set -e
if [ $RES -eq 0 ]; then
    echo "No schema changes detected."
fi