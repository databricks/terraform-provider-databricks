variable "cidr_block" {}
variable "region" {}

locals {
  // see https://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html
  allows = {
    "us-east-1" : [
      "52.27.216.188/32",
      "54.156.226.103/32",
    ],
    "us-east-2" : [
      "52.27.216.188/32",
      "18.221.200.169/32",
    ],
    "us-west-2" : [
      "10.200.0.0/16"
    ],
    "eu-west-1" : [
      "46.137.47.49/32",
    ]
  }
}

resource "aws_vpc" "main" {
  cidr_block           = var.cidr_block
  enable_dns_hostnames = true

  tags = merge(var.tags, {
    Name = "${var.prefix}-vpc"
  })
}

output "vpc_id" {
  value = aws_vpc.main.id
}

resource "aws_subnet" "public" {
  vpc_id     = aws_vpc.main.id
  cidr_block = cidrsubnet(aws_vpc.main.cidr_block, 3, 0)
  //availability_zone = "${var.region}b"

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
  subnet_id      = aws_subnet.public.id
}

resource "aws_eip" "nat" {
  vpc        = true
  depends_on = [aws_internet_gateway.gw]
  tags = merge(var.tags, {
    Name = "${var.prefix}-eip"
  })
}

resource "aws_nat_gateway" "gw" {
  allocation_id = aws_eip.nat.id
  subnet_id     = aws_subnet.public.id
  tags = merge(var.tags, {
    Name = "${var.prefix}-nat"
  })
}

output "subnet_public" {
  value = aws_subnet.public.id
}

resource "aws_subnet" "private" {
  vpc_id     = aws_vpc.main.id
  cidr_block = cidrsubnet(aws_vpc.main.cidr_block, 3, 1)
  //availability_zone = "${var.region}a"

  tags = merge(var.tags, {
    Name = "${var.prefix}-private-sn"
  })
}

resource "aws_route_table" "private" {
  vpc_id = aws_vpc.main.id

  route {
    nat_gateway_id = aws_nat_gateway.gw.id
    cidr_block     = "0.0.0.0/0"
  }

  tags = merge(var.tags, {
    Name = "${var.prefix}-private-rt"
  })
}

output "private_rt" {
  value = aws_route_table.private.id
}

resource "aws_route_table_association" "private" {
  route_table_id = aws_route_table.private.id
  subnet_id      = aws_subnet.private.id
}

output "subnet_private" {
  value = aws_subnet.private.id
}

resource "aws_security_group" "test_sg" {
  name        = "all all"
  description = "Allow inbound traffic"
  vpc_id      = aws_vpc.main.id

  ingress {
    description = "All Internal TCP"
    protocol    = "tcp"
    from_port   = 0
    to_port     = 65535
    self        = true
  }

  ingress {
    description = "All Internal UDP"
    protocol    = "udp"
    from_port   = 0
    to_port     = 65535
    self        = true
  }

  ingress {
    cidr_blocks = local.allows[var.region]
    description = "ssh"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
  }

  ingress {
    cidr_blocks = local.allows[var.region]
    description = "proxy"
    from_port   = 5557
    to_port     = 5557
    protocol    = "tcp"
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

output "security_group" {
  value = aws_security_group.test_sg.id
}
