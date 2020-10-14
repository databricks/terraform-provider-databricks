# AWS Databricks E2 workspaces within custom VPC

## data-security.tf

* creates S3 bucket, associated IAM role & instance profile
* exports `test_s3_bucket` & `test_ec2_instance_profile` variables

## main.tf

* creates & registers AWS cross-account role
* creates [common AWS](../modules/aws-mws-common) resources
* registers root bucket & VPC
* creates new databricks workspace
* creates new PAT token for that databricks workspace
