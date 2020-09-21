resource "aws_kms_key" "customer_managed_key" {
}

resource "aws_kms_grant" "databricks-grant" {
  name = "databricks-grant"
  key_id  = aws_kms_key.customer_managed_key.key_id
  // TODO: maybe this should be an env variable
  grantee_principal = "arn:aws:iam::414351767826:root"

  operations = ["Encrypt", "Decrypt"]
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