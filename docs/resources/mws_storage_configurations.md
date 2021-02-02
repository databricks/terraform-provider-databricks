# databricks_mws_storage_configurations Resource

-> **Note** This resource has an evolving API, which may change in future versions of the provider.

This resource to configure root bucket new workspaces within AWS.

It is important to understand that this will require you to configure your provider separately for the multiple workspaces resources. This will point to https://accounts.cloud.databricks.com for the HOST and it will use basic auth as that is the only authentication method available for multiple workspaces api. 

Please follow this [complete runnable example](https://github.com/databrickslabs/terraform-provider-databricks/blob/master/scripts/awsmt-integration/main.tf) with new VPC and new workspace setup. Please pay special attention to the fact that there you have two different instances of a databricks provider - one for deploying workspaces (with host=https://accounts.cloud.databricks.com/) and another for the workspace you've created with databricks_mws_workspaces resource. If you want both creation of workspaces & clusters within workspace within the same terraform module (essentially same directory), you should use the provider aliasing feature of Terraform. We strongly recommend having one terraform module for creation of workspace + PAT token and the rest in different modules.

## Example Usage

```hcl
resource "aws_s3_bucket" "root_storage_bucket" {
  bucket = "${var.prefix}-rootbucket"
  acl    = "private"
  versioning {
    enabled = false
  }
}

resource "databricks_mws_storage_configurations" "this" {
  provider                   = databricks.mws
  account_id                 = var.account_id
  storage_configuration_name = "${var.prefix}-storage"
  bucket_name                = aws_s3_bucket.root_storage_bucket.bucket
}
```

## Argument Reference

The following arguments are required:

* `bucket_name` - (Required) (String) name of AWS S3 bucket
* `account_id` - (Required) (String) master account id (also used for `sts:ExternaId` of `sts:AssumeRole`)
* `storage_configuration_name` - (Required) (String) name under which this storage configuration is stored

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the mws storage configurations.
* `storage_configuration_id` - (String) id of storage config to be used for `databricks_mws_workspace` resource.
