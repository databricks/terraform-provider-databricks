#!/usr/bin/env bash

source scripts/libschema.sh

cat $(generate_schema) | jq