# databricks_mws_networks Resource

-> **Note** This resource has an evolving API, which may change in future versions of the provider.

Use this resource to [configure VPC](https://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html) & subnets for new workspaces within AWS. It is essential to understand that this will require you to configure your provider separately for the multiple workspaces resources.

* Databricks must have access to at least two subnets for each workspace, with each subnet in a different Availability Zone. You cannot specify more than one Databricks workspace subnet per Availability Zone in the Create network configuration API call. You can have more than one subnet per Availability Zone as part of your network setup, but you can choose only one subnet per Availability Zone for the Databricks workspace.
* Databricks assigns two IP addresses per node, one for management traffic and one for Spark applications. The total number of instances for each subnet is equal to half of the available IP addresses.
* Each subnet must have a netmask between /17 and /25.
* Subnets must be private.
* Subnets must have outbound access to the public network using a NAT gateway and internet gateway, or other similar customer-managed appliance infrastructure.
* The NAT gateway must be set up in its subnet that routes quad-zero (0.0.0.0/0) traffic to an internet gateway or other customer-managed appliance infrastructure.

Please follow this [complete runnable example](https://github.com/databrickslabs/terraform-provider-databricks/blob/master/scripts/awsmt-integration/main.tf) with new VPC and new workspace setup. Please pay special attention to the fact that there you have two different instances of a databricks provider - one for deploying workspaces (with host=https://accounts.cloud.databricks.com/) and another for the workspace you've created with databricks_mws_workspaces resource. If you want both creations of workspaces & clusters within the same Terraform module (essentially the same directory), you should use the provider aliasing feature of Terraform. We strongly recommend having one terraform module to create workspace + PAT token and the rest in different modules.

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

resource "aws_internet_gateway" "gw" {
  vpc_id = aws_vpc.main.id
  tags = merge(var.tags, {
    Name = "${var.prefix}-igw"
  })
}

resource "aws_route_table" "public" {
  vpc_id = aws_vpc.main.id
  
  route {
    gateway_id = aws_internet_gateway.gw.id
    cidr_block = "0.0.0.0/0"
  }

  tags = merge(var.tags, {
    Name = "${var.prefix}-public-rt"
  })
}

resource "aws_route_table_association" "public" {
  route_table_id = aws_route_table.public.id
  subnet_id = aws_subnet.public.id
}

resource "aws_eip" "nat" {
  vpc = true
  depends_on = [aws_internet_gateway.gw]
}

resource "aws_nat_gateway" "gw" {
  allocation_id = aws_eip.nat.id
  subnet_id     = aws_subnet.public.id
  tags = merge(var.tags, {
    Name = "${var.prefix}-nat"
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

resource "aws_route_table" "private" {
  vpc_id = aws_vpc.main.id

  route {
    nat_gateway_id = aws_nat_gateway.gw.id
    cidr_block = "0.0.0.0/0"
  }

  tags = merge(var.tags, {
    Name = "${var.prefix}-private-rt"
  })
}

resource "aws_route_table_association" "private" {
  route_table_id = aws_route_table.private.id
  subnet_id = aws_subnet.private.id
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

  lifecycle {
    # you may need this workaround until issue #382 is fixed:
    # https://github.com/databrickslabs/terraform-provider-databricks/issues/382
    ignore_changes = [
      deployment_name
    ]
  }
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
