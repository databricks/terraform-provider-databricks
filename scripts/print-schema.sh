#!/usr/bin/env bash

source scripts/libschema.sh

cat $(generate_schema) | tail -n1 | jq