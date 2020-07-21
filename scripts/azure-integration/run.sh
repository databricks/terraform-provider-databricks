#!/bin/bash
set -e
cd $(dirname "$0")

function cleanup()
{
    if [ -z "$SKIP_CLEANUP" ]; then
        echo "[*] Destroy test infra"
        terraform destroy -auto-approve
        echo "\$SKIP_CLEANUP isn't set so removing any pre-existing terraform state"
        rm -f *.tfstate*
        rm -f .terraform
    else
        echo "\$SKIP_CLEANUP is set so 'terraform destroy' not run. Warning: Resources left in subscription."
    fi
}
trap cleanup EXIT

echo "[*] Creating test infrastructure"

if [ -n "$TEST_LOG_LEVEL" ]; then
    export TF_LOG=$TEST_LOG_LEVEL
    # Output debug log to file while tests run
    export TF_LOG_PATH=$PWD/tf.log
fi

terraform init
terraform apply -auto-approve

export $(terraform output --json | jq -r 'to_entries|map("\(.key|ascii_upcase)=\(.value.value|tostring)")|.[]')

# Run all Azure integration tests
TF_ACC=1 gotestsum \
    --format short-verbose \
    --raw-command go test -v \
    -json -coverprofile=coverage.out \
    -test.timeout 35m \
    -run '^(TestAcc|TestAzureAcc)' ../../...