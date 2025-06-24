---
subcategory: "Deployment"
---
# databricks_aws_crossaccount_policy Data Source

This data source constructs necessary AWS cross-account policy for you, which is based on [official documentation](https://docs.databricks.com/administration-guide/account-api/iam-role.html#language-Your%C2%A0VPC,%C2%A0default).

-> This data source can be used with an account or workspace-level provider.

## Example Usage

For more detailed usage please see [databricks_aws_assume_role_policy](aws_assume_role_policy.md) or [databricks_aws_s3_mount](../resources/mount.md) pages.

```hcl
data "databricks_aws_crossaccount_policy" "this" {}
```

## Argument Reference

* `policy_type` (Optional) The type of cross account policy to generated: `managed` for Databricks-managed VPC and `customer` for customer-managed VPC, `restricted` for customer-managed VPC with policy restrictions
* `pass_roles` (Optional) (List) List of Data IAM role ARNs that are explicitly granted `iam:PassRole` action.
The below arguments are only valid for `restricted` policy type
* `aws_account_id` — Your AWS account ID, which is a number.
* `aws_partition` - (Optional) AWS partition. The options are `aws`, `aws-us-gov`, or `aws-us-gov-dod`. Defaults to `aws`
* `vpc_id` — ID of the AWS VPC where you want to launch workspaces.
* `region` — AWS Region name for your VPC deployment, for example `us-west-2`.
* `security_group_id` — ID of your AWS security group. When you add a security group restriction, you cannot reuse the cross-account IAM role or reference a credentials ID (`credentials_id`) for any other workspaces. For those other workspaces, you must create separate roles, policies, and credentials objects.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `json` - AWS IAM Policy JSON document

## Related Resources

The following resources are used in the same context:

* [Provisioning AWS Databricks workspaces with a Hub & Spoke firewall for data exfiltration protection](../guides/aws-e2-firewall-hub-and-spoke.md) guide
* [databricks_aws_assume_role_policy](aws_assume_role_policy.md) data to construct the necessary AWS STS assume role policy.
* [databricks_aws_bucket_policy](aws_bucket_policy.md) data to configure a simple access policy for AWS S3 buckets, so that Databricks can access data in it.
* [databricks_instance_profile](../resources/instance_profile.md) to manage AWS EC2 instance profiles that users can launch [databricks_cluster](../resources/cluster.md) and access data, like [databricks_mount](../resources/mount.md).
