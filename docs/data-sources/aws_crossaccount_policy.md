---
subcategory: "AWS"
---
# databricks_aws_crossaccount_policy Data Source

-> **Note** This resource has an evolving API, which may change in future versions of the provider. Please always consult [latest documentation](https://docs.databricks.com/administration-guide/account-api/iam-role.html#language-Your%C2%A0VPC,%C2%A0default) in case of any questions.

This data source constructs necessary AWS cross-account policy for you, which is based on [official documentation](https://docs.databricks.com/administration-guide/account-api/iam-role.html#language-Your%C2%A0VPC,%C2%A0default).

## Example Usage

For more detailed usage please see [databricks_aws_assume_role_policy](aws_assume_role_policy.md) or [databricks_aws_s3_mount](../resources/aws_s3_mount.md) pages.

```hcl
data "databricks_aws_crossaccount_policy" "this" {}
```

## Argument Reference

* `pass_roles` (Optional) (List) List of Data IAM role ARNs that are explicitly granted `iam:PassRole` action.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `json` - AWS IAM Policy JSON document

## Related Resources

The following resources are used in the same context:

* [Provisioning AWS Databricks E2 with a Hub & Spoke firewall for data exfiltration protection](../guides/aws-e2-firewall-hub-and-spoke.md) guide
* [databricks_aws_assume_role_policy](aws_assume_role_policy.md) data to construct the necessary AWS STS assume role policy.
* [databricks_aws_bucket_policy](aws_bucket_policy.md) data to configure a simple access policy for AWS S3 buckets, so that Databricks can access data in it.
* [databricks_instance_profile](../resources/instance_profile.md) to manage AWS EC2 instance profiles that users can launch [databricks_cluster](../resources/cluster.md) and access data, like [databricks_mount](../resources/mount.md).
