# databricks_instance_profile Resource

This resource allows you to register or unregisters EC2 instance profiles that users can launch Databricks clusters with on AWS. The following example demonstrates how to create an instance profile and create cluster with it.

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
  skip_validation      = false
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

## Argument Reference

The following arguments are supported:

* `instance_profile_arn` - (Required) `ARN` attribute of `aws_iam_instance_profile` output, the EC2 instance profile association to AWS IAM role.
* `skip_validation` - (Required) whether or not to apply validation for 

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ARN for EC2 Instance Profile, that is registered with Databricks.


## Import

Importing this resource is not currently supported.