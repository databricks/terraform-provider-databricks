#!/usr/bin/env bash

# Function to generate provider schema
generate_schema() {
  >&2 make install 
  version=$(./terraform-provider-databricks version)

  local TMPDIR=$(mktemp -d)
  >&2 echo "Generating provider schema in $TMPDIR..."
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
  >&2 echo "Provider schema available in $TMPDIR/schema.json"
  echo "$TMPDIR/schema.json"
} 