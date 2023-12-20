#!/bin/bash

BASE_BRANCH="master"

checkout_branch() {
    local branch=$1
    echo "Checking out branch: $branch"
    git checkout $branch
}


# Function to generate provider schema
generate_schema() {
  local schema_name=$1
  make install
  version=$(./terraform-provider-databricks version)

  echo "Generating provider schema for $branch..."
  set -ex
  cd /tmp
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
  terraform providers schema -json > $schema_name.json
  cd - # Go back to the original directory
}

if [ -n "$(git status --porcelain)" ]; then
    echo "There are uncommitted changes. Please commit them before running this script."
    exit 1
fi

CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
generate_schema "schema-new"
checkout_branch $BASE_BRANCH
generate_schema "schema-current"
checkout_branch $CURRENT_BRANCH

jd -color /tmp/schema-current.json /tmp/schema-new.json