#!/bin/bash
set -e
cd $(dirname "$0")

export AWS_DEFAULT_REGION=$DATABRICKS_MWS_AWS_REGION
export DATABRICKS_USERNAME=$DATABRICKS_MWS_USERNAME
export DATABRICKS_PASSWORD=$DATABRICKS_MWS_PASSWORD

echo "Working in region: -- $AWS_DEFAULT_REGION"


function cleanup()
{
    echo -e "----> Destroy prereqs \n\n"
    if [ -z "$SKIP_CLEANUP" ]
    then
        terraform destroy -auto-approve \
          -var 'databricks_mws_aws_acct_id=414351767826' \
          -var 'databricks_mws_acct_id=$DATABRICKS_MWS_ACCT_ID'
    else
        echo "\$SKIP_CLEANUP is set so 'terraform destroy' not run. Warning: Resources left in aws account."
    fi
}
trap cleanup EXIT

echo -e "----> Running Terraform to create prereqs in AWS Account for MWS \n\n"

# Remove any old state unless SKIP_CLEANUP set
if [ -z "$SKIP_CLEANUP" ]
then
    echo "\$SKIP_CLEANUP isn't set so removing any pre-existing terraform state"
    rm -f *.tfstate
fi

terraform init
terraform apply -auto-approve \
  -var "databricks_mws_aws_acct_id=414351767826" \
  -var "databricks_mws_acct_id=$DATABRICKS_MWS_ACCT_ID"

export TEST_MWS_CROSS_ACCT_ROLE=$(terraform output aws_cross_acct_role_arn)
export TEST_MWS_ROOT_BUCKET=$(terraform output aws_s3_bucket_name)

export TEST_MWS_VPC_ID=$(terraform output aws_vpc_id)
export TEST_MWS_SUBNET_1=$(terraform output aws_subnet_1)
export TEST_MWS_SUBNET_2=$(terraform output aws_subnet_2)
export TEST_MWS_SG=$(terraform output aws_sg)


echo -e "----> Running AWS Multiple Workspaces Acceptance Tests \n\n"
# Output debug log to file while tests run
#export TF_LOG_PATH=$PWD/tf.log
# Run all AWS Multipleworkspace integration tests
TF_ACC=1 gotestsum --format short-verbose --raw-command go test -v -json -short -coverprofile=coverage.out -test.timeout 35m -run 'TestAccMWS' ./../...