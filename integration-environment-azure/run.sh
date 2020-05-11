#!/bin/bash
set -e
cd $(dirname "$0")

function cleanup()
{
    echo -e "----> Destroy prereqs \n\n"
    # terraform destroy -auto-approve
}
trap cleanup EXIT

echo -e "----> Running Terraform to create prereqs in Azure \n\n"

# Setup Auth for Azure RM provider in terraform
export ARM_CLIENT_ID=$DATABRICKS_AZURE_CLIENT_ID
export ARM_CLIENT_SECRET=$DATABRICKS_AZURE_CLIENT_SECRET
export ARM_SUBSCRIPTION_ID=$DATABRICKS_AZURE_SUBSCRIPTION_ID
export ARM_TENANT_ID=$DATABRICKS_AZURE_TENANT_ID

# Add back in before push to ensure fresh env
# on each run of the integration tests
# rm *.tfstate

terraform init
terraform apply -auto-approve

export TEST_RESOURCE_GROUP=$(terraform output rg_name) 
export TEST_WORKSPACE_NAME=$(terraform output workspace_name)
export TEST_GEN2_ADAL_NAME=$(terraform output gen2_adal_name)
export TEST_MANAGED_RESOURCE_GROUP=$(terraform output workspace_managed_rg_name)
export TEST_LOCATION=$(terraform output location)


echo -e "----> Running Azure Acceptance Tests \n\n"
# Run all Azure integration tests
TF_ACC=1 gotestsum --format short-verbose --raw-command go test -v -json -short -coverprofile=coverage.out -test.timeout 15m -run 'TestAccAzure' ./../...