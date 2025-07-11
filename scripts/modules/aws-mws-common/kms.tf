variable "databricks_aws_account_id" {
  default = "414351767826"
}

resource "aws_kms_key" "customer_managed_key" {
}

resource "aws_kms_grant" "databricks-grant" {
  name              = "databricks-grant"
  key_id            = aws_kms_key.customer_managed_key.key_id
  grantee_principal = "arn:aws:iam::${var.databricks_aws_account_id}:root"

  operations = ["Encrypt", "Decrypt", "DescribeKey",
    "GenerateDataKey", "ReEncryptFrom", "ReEncryptTo",
    "GenerateDataKeyWithoutPlaintext"]
}


output "kms_key_arn" {
  value = aws_kms_key.customer_managed_key.arn
}

resource "aws_kms_alias" "customer_managed_key_alias" {
  name          = "alias/${var.prefix}-customer-key-alias"
  target_key_id = aws_kms_key.customer_managed_key.key_id
}

output "kms_key_alias" {
  value = aws_kms_alias.customer_managed_key_alias.name
}
