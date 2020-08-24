# databricks_aws_bucket_policy Data Source

This datasource configures simple access policy for AWS S3 buckets, so that Databricks can access data in it. 

## Example Usage

```hcl
resource "aws_s3_bucket" "this" {
  bucket = "things"
  region = "eu-west-1"
  acl    = "private"
  force_destroy = true
}

data "databricks_aws_bucket_policy" "stuff" {
  bucket_name = aws_s3_bucket.this.bucket
}

resource "aws_s3_bucket_policy" "this" {
  bucket     = aws_s3_bucket.this.id
  policy     = data.databricks_aws_bucket_policy.this.json
}
```
## Argument Reference

* `bucket_name` - (Required) AWS S3 Bucket name for which to generate policy document.
* `json` - (Read-only) AWS IAM Policy JSON document to grand Databricks full access to bucket.
