#!/bin/bash

if ! type jq &> /dev/null; then
    >&2 echo "[-] Please install jq (https://stedolan.github.io/jq/) to run this command."
    exit 1
fi

if ! type terraform &> /dev/null; then
    >&2 echo "[-] Please install Terraform (https://www.terraform.io/downloads.html) to run this command."
    exit 1
fi

if [ "" == "$1" ]; then
    echo "Just exporting environment variables:"
    echo "$0 <env-name> --export"
    echo ""
    echo "Running cloud-specific tests:"
    echo "$0 <env-name> '^(TestAcc|TestAzureAcc)' [--destroy] [--quick] [--debug]"
    echo "--destroy is optional flag to destroy environment after tests are done"
    echo "--debug is optional flag to log DEBUG lines into env-integration/tf.log"
    exit
fi

OLDPWD="$PWD"
DIR=$(dirname "$0")
TARGET="${DIR}/$1-integration"
STATE="${DIR}/$1-integration/terraform.tfstate"
JQ='to_entries|map("\(.key|ascii_upcase)=\(.value.value|tostring)")|.[]'

if ! [ -d "$TARGET" ]; then
    >&2 echo "[-] Cannot find $1 integration tests."
    exit 1
fi

# ./run.sh mws --export to print exportable env
if [ "--export" == "$2" ]; then
    if ! [ -f "$STATE" ]; then
        >&2 echo "[-] $1 didn't provision environment yet."
        exit 1
    fi
    terraform output -state=$STATE --json | jq -r $JQ
    exit
fi

# check for environment variables early
if [ -f "$TARGET/require_env" ]; then
    M=0
    for var in $(cat $TARGET/require_env); do
        if [ "" == "${!var}" ]; then
            >&2 echo -e "[-] Missing $var variable."
            ((M++))
        fi
    done
    if [[ "$M" -gt 0 ]]; then
        >&2 echo "[-] Please set $M env variables and restart."
        exit 1
    fi
fi

cd $TARGET

if [[ $@ == *"--destroy"* ]]; then
    if [ -f "$STATE" ]; then
        function cleanup()
        {
            echo "[*] Cleanup with destroy"
            terraform destroy -auto-approve
            rm -f *.tfstate*
            rm -f .terraform
        }
        trap cleanup EXIT
    fi
fi

if [ -f "main.tf" ]; then
    rm -fr .terraform .terraform.lock.hcl
    terraform init -upgrade
    terraform apply -auto-approve
    export $(terraform output --json | jq -r $JQ) >/dev/null 2>&1
else
    echo "[*] $1 has no specific Terraform environment."
fi

if [[ $@ == *"--debug"* ]]; then
    export TF_LOG="DEBUG"
    export TF_LOG_PATH=$PWD/tf.log
    export TF_ACC_LOG_PATH=$PWD/tf.log
    echo "[*] To see debug logs: tail -f $PWD/tf.log"
fi

function go_test {
    # set tmp dir to short one
    export TMPDIR=/tmp
    TF_ACC=1 gotestsum \
    --format short-verbose \
    --raw-command go test -v \
    -json -coverprofile=coverage.out \
    -test.timeout 35m \
    -run $1 ../../...
}

if [[ $@ == *"--tee"* ]]; then
    go_test $2 2>&1 | tee out.log
    echo "✓ To output of existing tests: less $PWD/out.log"

    FAILURES=$(grep "\-\-\- FAIL" out.log | sed 's/--- FAIL: / \* \[ \] /g' | sort)
    PASSES=$(grep PASS out.log | grep Test | sort | sed 's/PASS/ \* \[x\]/')

cat <<-EOF > test-report.log
$1
---
${FAILURES}
${PASSES}
EOF
else
    go_test $2
fi
