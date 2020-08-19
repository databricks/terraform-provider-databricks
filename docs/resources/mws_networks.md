# databricks_mws_networks Resource

**This resource has evolving API, which may change in future versions of provider.**

This resource to configure VPC & subnets for new workspaces within AWS.

It is important to understand that this will require you to configure your provider separately for the multiple workspaces resources. This will point to https://accounts.cloud.databricks.com for the HOST and it will use basic auth as that is the only authentication method available for multiple workspaces api. 

Please follow this [complete runnable example](https://github.com/databrickslabs/terraform-provider-databricks/blob/master/scripts/awsmt-integration/main.tf) with new VPC and new workspace setup. Please pay special attention to the fact that there you have two different instances of a databricks provider - one for deploying workspaces (with host=https://accounts.cloud.databricks.com/) and another for the workspace you've created with databricks_mws_workspaces resource. If you want both creation of workspaces & clusters within workspace within the same terraform module (essentially same directory), you should use the provider aliasing feature of Terraform. We strongly recommend having one terraform module for creation of workspace + PAT token and the rest in different modules.

## Example Usage

```hcl
resource "aws_vpc" "main" {
  cidr_block           = data.external.env.result.TEST_CIDR
  enable_dns_hostnames = true

  tags = merge(var.tags, {
    Name = "${var.prefix}-vpc"
  })
}

resource "aws_subnet" "public" {
  vpc_id            = aws_vpc.main.id
  cidr_block        = cidrsubnet(aws_vpc.main.cidr_block, 3, 0)
  availability_zone = "${data.external.env.result.TEST_REGION}b"

  tags = merge(var.tags, {
    Name = "${var.prefix}-public-sn"
  })
}

resource "aws_subnet" "private" {
  vpc_id            = aws_vpc.main.id
  cidr_block        = cidrsubnet(aws_vpc.main.cidr_block, 3, 1)
  availability_zone = "${data.external.env.result.TEST_REGION}a"

  tags = merge(var.tags, {
    Name = "${var.prefix}-private-sn"
  })
}

resource "aws_internet_gateway" "gw" {
  vpc_id = aws_vpc.main.id
  tags = merge(var.tags, {
    Name = "${var.prefix}-igw"
  })
}

resource "aws_route" "r" {
  route_table_id            = aws_vpc.main.default_route_table_id
  destination_cidr_block    = "0.0.0.0/0"
  gateway_id = aws_internet_gateway.gw.id
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

  tags = merge(var.tags, {
    Name = "${var.prefix}-sg"
  })
}

resource "databricks_mws_networks" "this" {
  provider     = databricks.mws
  account_id   = var.account_id
  network_name = "${var.prefix}-network"
  vpc_id       = aws_vpc.main.id
  subnet_ids   = [aws_subnet.public.id, aws_subnet.private.id]
  security_group_ids = [aws_security_group.test_sg.id]
}
```

## Argument Reference

The following arguments are required:

* `account_id` - (Required) (String) master account id (also used for `sts:ExternaId` of `sts:AssumeRole`)
* `network_name` - (Required) (String) name under which this network is regisstered
* `vpc_id` - (Required) (String) AWS VPC id
* `subnet_ids` - (Required) (Set) ids of AWS VPC subnets
* `security_group_ids` - (Required) (Set) ids of AWS Security Groups

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the mws networks.
* `network_id` - (String) id of network to be used for `databricks_mws_workspace` resource.
* `vpc_status` - (String) VPC attachment status
* `workspace_id` - (Integer) id of associated workspace
