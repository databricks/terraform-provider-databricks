#!/bin/bash

BASE_BRANCH="master"

checkout_branch() {
    local branch=$1
    echo "Checking out branch: $branch"
    git checkout $branch
}


# Function to generate provider schema
generate_schema() {
  local TMPDIR=$1
  make install
  version=$(./terraform-provider-databricks version)

  echo "Generating provider schema for $branch..."
  set -ex
  pushd $TMPDIR
  cat > main.tf <<EOF
terraform {
  required_providers {
    databricks = {
      source = "databricks/databricks"
      version = "$version"
    }
  }
}
EOF
  terraform init
  terraform providers schema -json > schema.json
  popd
}

if [ -n "$(git status --porcelain)" ]; then
    echo "There are uncommitted changes. Please commit them before running this script."
    exit 1
fi

CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
NEW_TMP=$(mktemp -d -t schema-new)
generate_schema "$NEW_TMP"
checkout_branch $BASE_BRANCH
CURRENT_TMP=$(mktemp -d -t schema-current)
generate_schema "$CURRENT_TMP"
checkout_branch $CURRENT_BRANCH

jd -color "$CURRENT_TMP/schema.json" "$NEW_TMP/schema.json"