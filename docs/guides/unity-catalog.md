---
page_title: "Unity Catalog setup on AWS"
---

```hcl
provider "aws" {
  region = local.region
}

resource "aws_s3_bucket" "metastore" {
  bucket = "${local.prefix}-metastore"
  acl    = "private"
  versioning {
    enabled = false
  }
  force_destroy = true
  tags = merge(local.tags, {
    Name = "${local.prefix}-metastore"
  })
}

resource "aws_s3_bucket_public_access_block" "root_storage_bucket" {
  bucket             = aws_s3_bucket.metastore.id
  ignore_public_acls = true
  depends_on         = [aws_s3_bucket.metastore]
}

data "aws_iam_policy_document" "passrole_for_uc" {
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRole"]
    principals {
      identifiers = ["arn:aws:iam::414351767826:role/unity-catalog-prod-UCMasterRole-14S5ZJVKOTYTL"]
      type        = "AWS"
    }
    condition {
      test     = "StringEquals"
      variable = "sts:ExternalId"
      values   = [local.databricks_account_id]
    }
  }
}

resource "aws_iam_role" "metastore_data_access" {
  name               = "${local.prefix}-uc-access"
  assume_role_policy = data.aws_iam_policy_document.passrole_for_uc.json
  tags               = local.tags
  inline_policy {
    name = "${aws_s3_bucket.metastore.id}-access"
    policy = jsonencode({
      Version = "2012-10-17"
      Statement = [{
        Action = [
          "s3:PutObjectAcl",
          "s3:PutObject",
          "s3:ListBucket",
          "s3:GetObjectVersion",
          "s3:GetObject",
          "s3:GetBucketLocation",
          "s3:DeleteObject",
        ]
        Effect = "Allow"
        Resource = [
          aws_s3_bucket.metastore.arn,
          "${aws_s3_bucket.metastore.arn}/*"
        ]
      }]
    })
  }
}

resource "databricks_metastore" "this" {
  name          = "primary"
  storage_root  = "s3://${aws_s3_bucket.metastore.id}/metastore"
  owner         = "uc admins"
  force_destroy = true
}

resource "databricks_metastore_data_access" "this" {
  metastore_id = databricks_metastore.this.id
  name         = aws_iam_role.metastore_data_access.name
  aws_iam_role {
    role_arn = aws_iam_role.metastore_data_access.arn
  }
  is_default = true
}

resource "databricks_metastore_assignment" "this" {
  metastore_id = databricks_metastore.this.id
  workspace_id = local.workspace_id
}

resource "databricks_catalog" "sandbox" {
  metastore_id = databricks_metastore.this.id
  name         = "sandbox"
  comment      = "this catalog is managed by terraform"
  properties = {
    purpose = "testing"
  }
}

resource "databricks_grants" "sandbox" {
  catalog = databricks_catalog.sandbox.name
  grant {
    principal  = "Data Scientists"
    privileges = ["USAGE", "CREATE"]
  }
  grant {
    principal  = "Data Engineers"
    privileges = ["USAGE"]
  }
}

resource "databricks_schema" "things" {
  catalog_name = databricks_catalog.sandbox.id
  name         = "things"
  comment      = "this database is managed by terraform"
  properties = {
    kind = "various"
  }
}

resource "databricks_grants" "things" {
  schema = databricks_schema.things.id
  grant {
    principal  = "Data Engineers"
    privileges = ["USAGE"]
  }
}

resource "aws_s3_bucket" "external" {
  bucket = "${local.prefix}-external"
  acl    = "private"
  versioning {
    enabled = false
  }
  force_destroy = true
  tags = merge(local.tags, {
    Name = "${local.prefix}-external"
  })
}

resource "aws_s3_bucket_public_access_block" "external" {
  bucket             = aws_s3_bucket.external.id
  ignore_public_acls = true
  depends_on         = [aws_s3_bucket.external]
}

resource "aws_iam_role" "external_data_access" {
  name               = "${local.prefix}-external-access"
  assume_role_policy = data.aws_iam_policy_document.passrole_for_uc.json
  tags               = local.tags
  inline_policy {
    name = "${aws_s3_bucket.external.id}-access"
    policy = jsonencode({
      Version = "2012-10-17"
      Statement = [{
        Action = [
          "s3:PutObjectAcl",
          "s3:PutObject",
          "s3:ListBucket",
          "s3:GetObjectVersion",
          "s3:GetObject",
          "s3:GetBucketLocation",
          "s3:DeleteObject",
        ]
        Effect = "Allow"
        Resource = [
          aws_s3_bucket.external.arn,
          "${aws_s3_bucket.external.arn}/*"
        ]
      }]
    })
  }
}

resource "databricks_storage_credential" "external" {
  name = aws_iam_role.external_data_access.name
  aws_iam_role {
    role_arn = aws_iam_role.external_data_access.arn
  }
  comment = "Managed by TF"
}

resource "databricks_grants" "external_creds" {
  storage_credential = databricks_storage_credential.external.id
  grant {
    principal  = "Data Engineers"
    privileges = ["CREATE TABLE"]
  }
}

resource "databricks_external_location" "some" {
  name            = "external"
  url             = "s3://${aws_s3_bucket.external.id}/some"
  credential_name = databricks_storage_credential.external.id
  comment         = "Managed by TF"
}

resource "databricks_grants" "some" {
  external_location = databricks_external_location.some.id
  grant {
    principal  = "Data Engineers"
    privileges = ["CREATE TABLE", "READ FILES"]
  }
}
```
