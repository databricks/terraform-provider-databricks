variable "databricks_mws_aws_acct_id" {
  type    = string
}

variable "databricks_mws_acct_id" {
  type    = string
}

provider "aws" {
}

provider "random" {
  version = "~> 2.2"
}

resource "random_string" "naming" {
  special = false
  upper   = false
  length  = 6
}

data "template_file" "cross_account_role_policy" {
  template = "${file("${path.module}/templates/cross_account_role_policy.tpl")}"
}

data "template_file" "cross_account_role_assume_policy" {
  template = "${file("${path.module}/templates/cross_account_role_assume_policy.tpl")}"
  vars = {
    databricks_app_external_id = var.databricks_mws_acct_id
    databricks_aws_account_id = var.databricks_mws_aws_acct_id
  }
}

resource "aws_iam_role" "cross_account_role" {
  name = "tf_test_cross_acct_role_${random_string.naming.result}"
  assume_role_policy = data.template_file.cross_account_role_assume_policy.rendered
}

resource "aws_iam_policy" "cross_account_role_policy" {
  name = "tf_test_cross_acct_role_${random_string.naming.result}_policy"
  description = "E2 Workspace Cross account role policy policy"
  policy = data.template_file.cross_account_role_policy.rendered
}

resource "aws_iam_role_policy_attachment" "cross_account_role_policy_attach" {
  role       = aws_iam_role.cross_account_role.name
  policy_arn = aws_iam_policy.cross_account_role_policy.arn
}

data "template_file" "storage_bucket_policy" {
  template = "${file("${path.module}/templates/storage_bucket_policy.tpl")}"
  vars = {
    bucket_name = aws_s3_bucket.root_storage_bucket.bucket
    databricks_aws_account_id = var.databricks_mws_aws_acct_id
  }
}


resource "aws_s3_bucket" "root_storage_bucket" {
  bucket = "tf-test-root-bucket-${random_string.naming.result}"
  acl    = "private"
  versioning {
    enabled = false
  }
  force_destroy = true

  tags = {
    Name        = "tf-test-bucket"
    Environment = "Dev"
    Owner = "test@databricks.com"
  }
}


resource "aws_s3_bucket_policy" "root_bucket_policy" {
  bucket = aws_s3_bucket.root_storage_bucket.id
  policy = data.template_file.storage_bucket_policy.rendered
}

resource "aws_vpc" "main" {
  cidr_block       = "10.0.0.0/16"
  enable_dns_hostnames = true

  tags = {
    Name = "tf-test-mws-vpc"
  }
}

resource "aws_subnet" "public" {
  vpc_id     = aws_vpc.main.id
  cidr_block = "10.0.1.0/24"
  availability_zone = "us-east-1b"

  tags = {
    Name = "public-subnet"
  }
}

resource "aws_subnet" "private" {
  vpc_id     = aws_vpc.main.id
  cidr_block = "10.0.2.0/24"
  availability_zone = "us-east-1a"
  tags = {
    Name = "private-subnet"
  }
}

resource "aws_internet_gateway" "gw" {
  vpc_id = aws_vpc.main.id
  tags = {
    Name = "test-igw"
  }
}

resource "aws_route" "r" {
  route_table_id            = aws_vpc.main.default_route_table_id
  destination_cidr_block    = "0.0.0.0/0"
  gateway_id = aws_internet_gateway.gw.id

  depends_on = [aws_internet_gateway.gw, aws_vpc.main]
}

resource "aws_security_group" "test_sg" {
  name        = "all all"
  description = "Allow inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "All"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "test_sg"
  }
}


output "aws_s3_bucket_name" {
  value = aws_s3_bucket.root_storage_bucket.bucket
}

output "aws_cross_acct_role_arn" {
  value = aws_iam_role.cross_account_role.arn
}

output "aws_vpc_id" {
  value = aws_vpc.main.id
}

output "aws_subnet_1" {
  value = aws_subnet.public.id
}

output "aws_subnet_2" {
  value = aws_subnet.private.id
}

output "aws_sg" {
  value = aws_security_group.test_sg.id
}
