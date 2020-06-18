#!/bin/bash
cd $(dirname "$0")

set -e

cd ../../
go build -mod vendor -v -o terraform-provider-databricks
rm -f -r $HOME/.terraform.d/plugins
mkdir -p $HOME/.terraform.d/plugins/
mv terraform-provider-databricks $HOME/.terraform.d/plugins/terraform-provider-databricks_v0.2
cd -

rm -f -r .terraform
rm -f terraform.state

terraform init
terraform apply -var-file ./terraform.tfvars