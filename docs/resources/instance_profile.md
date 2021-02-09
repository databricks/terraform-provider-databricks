---
subcategory: "AWS"
---
# databricks_instance_profile Resource

This resource allows you to register or unregister EC2 instance profiles that users can launch [databricks_cluster](cluster.md) and access data, like [databricks_aws_s3_mount](aws_s3_mount.md). The following example demonstrates how to create an instance profile and create a cluster with it.

```hcl
variable "crossaccount_role_name" {
  type        = string
  description = "Role that you've specified on https://accounts.cloud.databricks.com/#aws"
}
data "aws_iam_policy_document" "assume_role_for_ec2" {
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRole"]
    principals {
      identifiers = ["ec2.amazonaws.com"]
      type        = "Service"
    }
  }
}
resource "aws_iam_role" "role_for_s3_access" {
  name               = "shared-ec2-role-for-s3"
  description        = "Role for shared access"
  assume_role_policy = data.aws_iam_policy_document.assume_role_for_ec2.json
}
data "aws_iam_policy_document" "pass_role_for_s3_access" {
  statement {
    effect    = "Allow"
    actions   = ["iam:PassRole"]
    resources = [aws_iam_role.role_for_s3_access.arn]
  }
}
resource "aws_iam_policy" "pass_role_for_s3_access" {
  name   = "shared-pass-role-for-s3-access"
  path   = "/"
  policy = data.aws_iam_policy_document.pass_role_for_s3_access.json
}
resource "aws_iam_role_policy_attachment" "cross_account" {
  policy_arn = aws_iam_policy.pass_role_for_s3_access.arn
  role       = var.crossaccount_role_name
}
resource "aws_iam_instance_profile" "shared" {
  name = "shared-instance-profile"
  role = aws_iam_role.role_for_s3_access.name
}
resource "databricks_instance_profile" "shared" {
  instance_profile_arn = aws_iam_instance_profile.shared.arn
}
resource "databricks_cluster" "this" {
  cluster_name            = "Shared Autoscaling"
  spark_version           = "6.6.x-scala2.11"
  node_type_id            = "i3.xlarge"
  autotermination_minutes = 20
  autoscale {
    min_workers = 1
    max_workers = 50
  }
  aws_attributes {
    instance_profile_arn    = databricks_instance_profile.shared.id
    availability            = "SPOT"
    zone_id                 = "us-east-1"
    first_on_demand         = 1
    spot_bid_price_percent  = 100
  }
}
```

## Usage with Cluster Policies

It is advised to keep all common configurations in [Cluster Policies](cluster_policy.md) to maintain control of the environments launched, so `databricks_cluster` above could be replaced with `databricks_cluster_policy`:

```hcl
resource "databricks_cluster_policy" "this" {
  name = "Policy with predefined instance profile"
  definition = jsonencode({
    # most likely policy might have way more things init.
    "aws_attributes.instance_profile_arn": {
      "type": "fixed",
      "value": databricks_instance_profile.shared.arn
    }
  })
}
```

## Granting access to all users

You can make instance profile available to all users by [associating it](group_instance_profile.md) with the special group called `users` through [databricks_group](../data-sources/group.md) data source.

```hcl
resource "databricks_instance_profile" "this" {
  instance_profile_arn = aws_iam_instance_profile.shared.arn
}

data "databricks_group" "users" {
  display_name = "users"
}

resource "databricks_group_instance_profile" "all" {
  group_id            = data.databricks_group.users.id
  instance_profile_id = databricks_instance_profile.this.id
}
```

## Argument Reference

The following arguments are supported:

* `instance_profile_arn` - (Required) `ARN` attribute of `aws_iam_instance_profile` output, the EC2 instance profile association to AWS IAM role.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ARN for EC2 Instance Profile, that is registered with Databricks.


## Import

The resource instance profile can be imported using the ARN of it

```bash
$ terraform import databricks_instance_profile.this <instance-profile-arn>
```
