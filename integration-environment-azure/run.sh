#!/bin/bash
set -e
cd $(dirname "$0")

function cleanup()
{
    echo -e "----> Destroy prereqs \n\n"
    if [ -z "$SKIP_CLEANUP" ]
    then
        terraform destroy -auto-approve
    else
        echo "\$SKIP_CLEANUP is set so 'terraform destroy' not run. Warning: Resources left in subscription."
    fi
}
trap cleanup EXIT

echo -e "----> Running Terraform to create prereqs in Azure \n\n"

# Setup Auth for Azure RM provider in terraform
export ARM_CLIENT_ID=$DATABRICKS_AZURE_CLIENT_ID
export ARM_CLIENT_SECRET=$DATABRICKS_AZURE_CLIENT_SECRET
export ARM_SUBSCRIPTION_ID=$DATABRICKS_AZURE_SUBSCRIPTION_ID
export ARM_TENANT_ID=$DATABRICKS_AZURE_TENANT_ID

# Remove any old state unless SKIP_CLEANUP set
if [ -z "$SKIP_CLEANUP" ]
then
    echo "\$SKIP_CLEANUP isn't set so removing any pre-existing terraform state"
    rm -f *.tfstate
fi

terraform init
terraform apply -auto-approve

export TEST_RESOURCE_GROUP=$(terraform output rg_name) 
export TEST_WORKSPACE_NAME=$(terraform output workspace_name)
export TEST_GEN2_ADAL_NAME=$(terraform output gen2_adal_name)
export TEST_MANAGED_RESOURCE_GROUP=$(terraform output workspace_managed_rg_name)
export TEST_LOCATION=$(terraform output location)
export TEST_STORAGE_ACCOUNT_KEY=$(terraform output blob_storage_key)
export TEST_STORAGE_ACCOUNT_NAME=$(terraform output blob_storage_name)

echo -e "----> Running Azure Acceptance Tests \n\n"
# Output debug log to file while tests run
export TF_LOG_PATH=$PWD/tf.log
# Run all Azure integration tests
TF_LOG=debug TF_ACC=1 gotestsum --format short-verbose --raw-command go test -v -json -short -coverprofile=coverage.out -test.timeout 35m -run 'TestAccAzure' ./../...