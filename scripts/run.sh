#!/bin/bash

if ! type jq &> /dev/null; then
    >&2 echo "Please install jq (https://stedolan.github.io/jq/) to run this command."
    exit 1
fi

if ! type terraform &> /dev/null; then
    >&2 echo "Please install Terraform (https://www.terraform.io/downloads.html) to run this command."
    exit 1
fi

OLDPWD="$PWD"
DIR=$(dirname "$0")
TARGET="${DIR}/$1-integration"
STATE="${DIR}/$1-integration/terraform.tfstate"

if ! [ -d "$TARGET" ]; then
    >&2 echo "Cannot find $1 integration tests."
    exit 1
fi

# ./run.sh mws --export to print exportable env
if [ "--export" == "$2" ]; then 
    if ! [ -f "$STATE" ]; then
        >&2 echo "$1 didn't provision environment yet."
        exit 1
    fi
    terraform output -state=$STATE --json | jq -r 'to_entries|map("\(.key|ascii_upcase)=\(.value.value|tostring)")|.[]'
    exit
fi

# check for environment variables early
if [ -f "$TARGET/require_env" ]; then
    M=0
    for var in $(cat $TARGET/require_env); do
        if [ "" == "${!var}" ]; then
            >&2 echo -e "Missing $var variable."
            ((M++))
        fi
    done
    if [[ "$M" -gt 0 ]]; then
        >&2 echo "Please set $M env variables and restart."
        exit 1
    fi
fi



# cd $TARGET

# terraform init
# terraform apply
# TestAccMWSWorkspaces

# $2 = -run '^(TestAcc|TestAzureAcc)' ../../...

TF_ACC=1 gotestsum \
    --format short-verbose \
    --raw-command go test -v \
    -json -coverprofile=coverage.out \
    -test.timeout 35m \
    -run $2 ...