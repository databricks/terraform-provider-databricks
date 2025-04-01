package aws

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestDataAwsCrossAccountDatabricksManagedPolicy(t *testing.T) {
	d, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsCrossaccountPolicy(),
		NonWritable: true,
		ID:          ".",
	}.Apply(t)
	assert.NoError(t, err)
	j := d.Get("json")
	assert.Lenf(t, j, 3171, "Strange length for policy: %s", j)
}

func TestDataAwsCrossAccountCustomerManagedPolicy(t *testing.T) {
	d, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsCrossaccountPolicy(),
		NonWritable: true,
		HCL:         `policy_type = "customer"`,
		ID:          ".",
	}.Apply(t)
	assert.NoError(t, err)
	j := d.Get("json")
	assert.Lenf(t, j, 2328, "Strange length for policy: %s", j)
}

func TestDataAwsCrossAccountPolicy_WithPassRoles(t *testing.T) {
	d, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsCrossaccountPolicy(),
		NonWritable: true,
		HCL:         `pass_roles = ["a", "b", "c"]`,
		ID:          ".",
	}.Apply(t)
	assert.NoError(t, err)
	j := d.Get("json")
	assert.Lenf(t, j, 3307, "Strange length for policy: %s", j)
}

func TestDataAwsCrossAccountManagedPolicyRoles(t *testing.T) {
	expectedJSON := `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "ec2:AssignPrivateIpAddresses",
        "ec2:CancelSpotInstanceRequests",
        "ec2:DescribeAvailabilityZones",
        "ec2:DescribeIamInstanceProfileAssociations",
        "ec2:DescribeInstanceStatus",
        "ec2:DescribeInstances",
        "ec2:DescribeInternetGateways",
        "ec2:DescribeNatGateways",
        "ec2:DescribeNetworkAcls",
        "ec2:DescribePrefixLists",
        "ec2:DescribeReservedInstancesOfferings",
        "ec2:DescribeRouteTables",
        "ec2:DescribeSecurityGroups",
        "ec2:DescribeSpotInstanceRequests",
        "ec2:DescribeSpotPriceHistory",
        "ec2:DescribeSubnets",
        "ec2:DescribeVolumes",
        "ec2:DescribeVpcAttribute",
        "ec2:DescribeVpcs",
        "ec2:CreateTags",
        "ec2:DeleteTags",
        "ec2:GetSpotPlacementScores",
        "ec2:RequestSpotInstances",
        "ec2:DescribeFleetHistory",
        "ec2:ModifyFleet",
        "ec2:DeleteFleets",
        "ec2:DescribeFleetInstances",
        "ec2:DescribeFleets",
        "ec2:CreateFleet",
        "ec2:DeleteLaunchTemplate",
        "ec2:GetLaunchTemplateData",
        "ec2:CreateLaunchTemplate",
        "ec2:DescribeLaunchTemplates",
        "ec2:DescribeLaunchTemplateVersions",
        "ec2:ModifyLaunchTemplate",
        "ec2:DeleteLaunchTemplateVersions",
        "ec2:CreateLaunchTemplateVersion",
        "ec2:AssociateIamInstanceProfile",
        "ec2:AttachVolume",
        "ec2:AuthorizeSecurityGroupEgress",
        "ec2:AuthorizeSecurityGroupIngress",
        "ec2:CreateVolume",
        "ec2:DeleteVolume",
        "ec2:DetachVolume",
        "ec2:DisassociateIamInstanceProfile",
        "ec2:ReplaceIamInstanceProfileAssociation",
        "ec2:RevokeSecurityGroupEgress",
        "ec2:RevokeSecurityGroupIngress",
        "ec2:RunInstances",
        "ec2:TerminateInstances",
        "ec2:AttachInternetGateway",
        "ec2:AllocateAddress",
        "ec2:AssociateDhcpOptions",
        "ec2:AssociateRouteTable",
        "ec2:CreateDhcpOptions",
        "ec2:CreateInternetGateway",
        "ec2:CreateNatGateway",
        "ec2:CreateRoute",
        "ec2:CreateRouteTable",
        "ec2:CreateSecurityGroup",
        "ec2:CreateSubnet",
        "ec2:CreateVpc",
        "ec2:CreateVpcEndpoint",
        "ec2:DeleteDhcpOptions",
        "ec2:DeleteInternetGateway",
        "ec2:DeleteNatGateway",
        "ec2:DeleteRoute",
        "ec2:DeleteRouteTable",
        "ec2:DeleteSecurityGroup",
        "ec2:DeleteSubnet",
        "ec2:DeleteVpc",
        "ec2:DeleteVpcEndpoints",
        "ec2:DetachInternetGateway",
        "ec2:DisassociateRouteTable",
        "ec2:ModifyVpcAttribute",
        "ec2:ReleaseAddress"
      ],
      "Resource": "*"
    },
    {
      "Effect": "Allow",
      "Action": [
        "iam:CreateServiceLinkedRole",
        "iam:PutRolePolicy"
      ],
      "Resource": "arn:aws:iam::*:role/aws-service-role/spot.amazonaws.com/AWSServiceRoleForEC2Spot",
      "Condition": {
        "StringLike": {
          "iam:AWSServiceName": "spot.amazonaws.com"
        }
      }
    }
  ]
}`

	d, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsCrossaccountPolicy(),
		NonWritable: true,
		HCL:         `policy_type = "managed"`,
		ID:          ".",
	}.Apply(t)
	assert.NoError(t, err)
	actualJSON := d.Get("json").(string)
	assert.Equal(t, expectedJSON, actualJSON)

	// Negative test: ensure that customer policy is not equal to customer policy
	managedD, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsCrossaccountPolicy(),
		NonWritable: true,
		HCL:         `policy_type = "customer"`,
		ID:          ".",
	}.Apply(t)
	assert.NoError(t, err)
	managedJSON := managedD.Get("json").(string)
	assert.NotEqual(t, actualJSON, managedJSON)
}

func TestDataAwsCrossAccountCustomerManagedPolicyRoles(t *testing.T) {
	expectedJSON := `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "ec2:AssignPrivateIpAddresses",
        "ec2:CancelSpotInstanceRequests",
        "ec2:DescribeAvailabilityZones",
        "ec2:DescribeIamInstanceProfileAssociations",
        "ec2:DescribeInstanceStatus",
        "ec2:DescribeInstances",
        "ec2:DescribeInternetGateways",
        "ec2:DescribeNatGateways",
        "ec2:DescribeNetworkAcls",
        "ec2:DescribePrefixLists",
        "ec2:DescribeReservedInstancesOfferings",
        "ec2:DescribeRouteTables",
        "ec2:DescribeSecurityGroups",
        "ec2:DescribeSpotInstanceRequests",
        "ec2:DescribeSpotPriceHistory",
        "ec2:DescribeSubnets",
        "ec2:DescribeVolumes",
        "ec2:DescribeVpcAttribute",
        "ec2:DescribeVpcs",
        "ec2:CreateTags",
        "ec2:DeleteTags",
        "ec2:GetSpotPlacementScores",
        "ec2:RequestSpotInstances",
        "ec2:DescribeFleetHistory",
        "ec2:ModifyFleet",
        "ec2:DeleteFleets",
        "ec2:DescribeFleetInstances",
        "ec2:DescribeFleets",
        "ec2:CreateFleet",
        "ec2:DeleteLaunchTemplate",
        "ec2:GetLaunchTemplateData",
        "ec2:CreateLaunchTemplate",
        "ec2:DescribeLaunchTemplates",
        "ec2:DescribeLaunchTemplateVersions",
        "ec2:ModifyLaunchTemplate",
        "ec2:DeleteLaunchTemplateVersions",
        "ec2:CreateLaunchTemplateVersion",
        "ec2:AssociateIamInstanceProfile",
        "ec2:AttachVolume",
        "ec2:AuthorizeSecurityGroupEgress",
        "ec2:AuthorizeSecurityGroupIngress",
        "ec2:CreateVolume",
        "ec2:DeleteVolume",
        "ec2:DetachVolume",
        "ec2:DisassociateIamInstanceProfile",
        "ec2:ReplaceIamInstanceProfileAssociation",
        "ec2:RevokeSecurityGroupEgress",
        "ec2:RevokeSecurityGroupIngress",
        "ec2:RunInstances",
        "ec2:TerminateInstances"
      ],
      "Resource": "*"
    },
    {
      "Effect": "Allow",
      "Action": [
        "iam:CreateServiceLinkedRole",
        "iam:PutRolePolicy"
      ],
      "Resource": "arn:aws:iam::*:role/aws-service-role/spot.amazonaws.com/AWSServiceRoleForEC2Spot",
      "Condition": {
        "StringLike": {
          "iam:AWSServiceName": "spot.amazonaws.com"
        }
      }
    }
  ]
}`

	d, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsCrossaccountPolicy(),
		NonWritable: true,
		HCL:         `policy_type = "customer"`,
		ID:          ".",
	}.Apply(t)
	assert.NoError(t, err)
	actualJSON := d.Get("json").(string)
	assert.Equal(t, expectedJSON, actualJSON)

	// Negative test: ensure that customer policy is not equal to managed policy
	managedD, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsCrossaccountPolicy(),
		NonWritable: true,
		HCL:         `policy_type = "managed"`,
		ID:          ".",
	}.Apply(t)
	assert.NoError(t, err)
	managedJSON := managedD.Get("json").(string)
	assert.NotEqual(t, actualJSON, managedJSON)
}

func TestDataAwsCrossAccountRestrictedPolicyRoles(t *testing.T) {
	expectedJSON := `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "ec2:AssignPrivateIpAddresses",
        "ec2:CancelSpotInstanceRequests",
        "ec2:DescribeAvailabilityZones",
        "ec2:DescribeIamInstanceProfileAssociations",
        "ec2:DescribeInstanceStatus",
        "ec2:DescribeInstances",
        "ec2:DescribeInternetGateways",
        "ec2:DescribeNatGateways",
        "ec2:DescribeNetworkAcls",
        "ec2:DescribePrefixLists",
        "ec2:DescribeReservedInstancesOfferings",
        "ec2:DescribeRouteTables",
        "ec2:DescribeSecurityGroups",
        "ec2:DescribeSpotInstanceRequests",
        "ec2:DescribeSpotPriceHistory",
        "ec2:DescribeSubnets",
        "ec2:DescribeVolumes",
        "ec2:DescribeVpcAttribute",
        "ec2:DescribeVpcs",
        "ec2:CreateTags",
        "ec2:DeleteTags",
        "ec2:GetSpotPlacementScores",
        "ec2:RequestSpotInstances",
        "ec2:DescribeFleetHistory",
        "ec2:ModifyFleet",
        "ec2:DeleteFleets",
        "ec2:DescribeFleetInstances",
        "ec2:DescribeFleets",
        "ec2:CreateFleet",
        "ec2:DeleteLaunchTemplate",
        "ec2:GetLaunchTemplateData",
        "ec2:CreateLaunchTemplate",
        "ec2:DescribeLaunchTemplates",
        "ec2:DescribeLaunchTemplateVersions",
        "ec2:ModifyLaunchTemplate",
        "ec2:DeleteLaunchTemplateVersions",
        "ec2:CreateLaunchTemplateVersion"
      ],
      "Resource": "*"
    },
    {
      "Effect": "Allow",
      "Action": [
        "iam:CreateServiceLinkedRole",
        "iam:PutRolePolicy"
      ],
      "Resource": "arn:aws:iam::*:role/aws-service-role/spot.amazonaws.com/AWSServiceRoleForEC2Spot",
      "Condition": {
        "StringLike": {
          "iam:AWSServiceName": "spot.amazonaws.com"
        }
      }
    },
    {
      "Sid": "InstancePoolsSupport",
      "Effect": "Allow",
      "Action": [
        "ec2:AssociateIamInstanceProfile",
        "ec2:DisassociateIamInstanceProfile",
        "ec2:ReplaceIamInstanceProfileAssociation"
      ],
      "Resource": "arn:aws:ec2:us-west-2:123456789012:instance/*",
      "Condition": {
        "StringEquals": {
          "ec2:ResourceTag/Vendor": "Databricks"
        }
      }
    },
    {
      "Sid": "AllowEc2RunInstancePerTag",
      "Effect": "Allow",
      "Action": "ec2:RunInstances",
      "Resource": [
        "arn:aws:ec2:us-west-2:123456789012:volume/*",
        "arn:aws:ec2:us-west-2:123456789012:instance/*"
      ],
      "Condition": {
        "StringEquals": {
          "aws:RequestTag/Vendor": "Databricks"
        }
      }
    },
    {
      "Sid": "AllowEc2RunInstanceImagePerTag",
      "Effect": "Allow",
      "Action": "ec2:RunInstances",
      "Resource": "arn:aws:ec2:us-west-2:123456789012:image/*",
      "Condition": {
        "StringEquals": {
          "aws:ResourceTag/Vendor": "Databricks"
        }
      }
    },
    {
      "Sid": "AllowEc2RunInstancePerVPCid",
      "Effect": "Allow",
      "Action": "ec2:RunInstances",
      "Resource": [
        "arn:aws:ec2:us-west-2:123456789012:network-interface/*",
        "arn:aws:ec2:us-west-2:123456789012:subnet/*",
        "arn:aws:ec2:us-west-2:123456789012:security-group/*"
      ],
      "Condition": {
        "StringEquals": {
          "ec2:vpc": "arn:aws:ec2:us-west-2:123456789012:vpc/vpc-abcdefg12345"
        }
      }
    },
    {
      "Sid": "AllowEc2RunInstanceOtherResources",
      "Effect": "Allow",
      "Action": "ec2:RunInstances",
      "NotResource": [
        "arn:aws:ec2:us-west-2:123456789012:image/*",
        "arn:aws:ec2:us-west-2:123456789012:network-interface/*",
        "arn:aws:ec2:us-west-2:123456789012:subnet/*",
        "arn:aws:ec2:us-west-2:123456789012:security-group/*",
        "arn:aws:ec2:us-west-2:123456789012:volume/*",
        "arn:aws:ec2:us-west-2:123456789012:instance/*"
      ]
    },
    {
      "Sid": "EC2TerminateInstancesTag",
      "Effect": "Allow",
      "Action": "ec2:TerminateInstances",
      "Resource": "arn:aws:ec2:us-west-2:123456789012:instance/*",
      "Condition": {
        "StringEquals": {
          "ec2:ResourceTag/Vendor": "Databricks"
        }
      }
    },
    {
      "Sid": "EC2AttachDetachVolumeTag",
      "Effect": "Allow",
      "Action": [
        "ec2:AttachVolume",
        "ec2:DetachVolume"
      ],
      "Resource": [
        "arn:aws:ec2:us-west-2:123456789012:instance/*",
        "arn:aws:ec2:us-west-2:123456789012:volume/*"
      ],
      "Condition": {
        "StringEquals": {
          "ec2:ResourceTag/Vendor": "Databricks"
        }
      }
    },
    {
      "Sid": "EC2CreateVolumeByTag",
      "Effect": "Allow",
      "Action": "ec2:CreateVolume",
      "Resource": "arn:aws:ec2:us-west-2:123456789012:volume/*",
      "Condition": {
        "StringEquals": {
          "aws:RequestTag/Vendor": "Databricks"
        }
      }
    },
    {
      "Sid": "EC2DeleteVolumeByTag",
      "Effect": "Allow",
      "Action": "ec2:DeleteVolume",
      "Resource": [
        "arn:aws:ec2:us-west-2:123456789012:volume/*"
      ],
      "Condition": {
        "StringEquals": {
          "ec2:ResourceTag/Vendor": "Databricks"
        }
      }
    },
    {
      "Sid": "VpcNonresourceSpecificActions",
      "Effect": "Allow",
      "Action": [
        "ec2:AuthorizeSecurityGroupEgress",
        "ec2:AuthorizeSecurityGroupIngress",
        "ec2:RevokeSecurityGroupEgress",
        "ec2:RevokeSecurityGroupIngress"
      ],
      "Resource": "arn:aws:ec2:us-west-2:123456789012:security-group/sg-12345678",
      "Condition": {
        "StringEquals": {
          "ec2:vpc": "arn:aws:ec2:us-west-2:123456789012:vpc/vpc-abcdefg12345"
        }
      }
    }
  ]
}`

	d, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsCrossaccountPolicy(),
		NonWritable: true,
		HCL: `
policy_type       = "restricted"
aws_account_id    = "123456789012"
vpc_id            = "vpc-abcdefg12345"
region            = "us-west-2"
security_group_id = "sg-12345678"
`,
		ID: ".",
	}.Apply(t)
	assert.NoError(t, err)
	actualJSON := d.Get("json").(string)
	assert.Equal(t, expectedJSON, actualJSON)

	// Negative test: ensure that restricted policy is not equal to managed policy
	managedD, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsCrossaccountPolicy(),
		NonWritable: true,
		HCL:         `policy_type = "managed"`,
		ID:          ".",
	}.Apply(t)
	assert.NoError(t, err)
	managedJSON := managedD.Get("json").(string)
	assert.NotEqual(t, actualJSON, managedJSON)

	// Negative test: ensure that restricted policy is not equal to customer policy
	customerD, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsCrossaccountPolicy(),
		NonWritable: true,
		HCL:         `policy_type = "customer"`,
		ID:          ".",
	}.Apply(t)
	assert.NoError(t, err)
	customerJSON := customerD.Get("json").(string)
	assert.NotEqual(t, actualJSON, customerJSON)
}

func TestDataAwsCrossAccountRestrictedPolicy(t *testing.T) {
	d, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsCrossaccountPolicy(),
		NonWritable: true,
		HCL: `
		policy_type = "restricted"
		aws_account_id = "123456789012"
		vpc_id = "vpc-12345678"
		region = "us-west-2"
		security_group_id = "sg-12345678"`,
		ID: ".",
	}.Apply(t)
	assert.NoError(t, err)
	j := d.Get("json")
	assert.Lenf(t, j, 5725, "Strange length for policy: %s", j)
}

func TestDataAwsCrossAccountRestrictedPolicyPartitionGov(t *testing.T) {
	d, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsCrossaccountPolicy(),
		NonWritable: true,
		HCL: `
		policy_type = "restricted"
		aws_account_id = "123456789012"
    aws_partition = "aws-us-gov"
		vpc_id = "vpc-12345678"
		region = "us-gov-west-1"
		security_group_id = "sg-12345678"`,
		ID: ".",
	}.Apply(t)
	assert.NoError(t, err)
	j := d.Get("json")
	assert.Lenf(t, j, 5963, "Strange length for policy: %s", j)
}

func TestDataAwsCrossAccountRestrictedPolicyPartitionGovDoD(t *testing.T) {
	d, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsCrossaccountPolicy(),
		NonWritable: true,
		HCL: `
		policy_type = "restricted"
		aws_account_id = "123456789012"
    aws_partition = "aws-us-gov-dod"
		vpc_id = "vpc-12345678"
		region = "us-gov-west-1"
		security_group_id = "sg-12345678"`,
		ID: ".",
	}.Apply(t)
	assert.NoError(t, err)
	j := d.Get("json")
	assert.Lenf(t, j, 5963, "Strange length for policy: %s", j)
}

func TestDataAwsCrossAccountInvalidPolicy(t *testing.T) {
	qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsCrossaccountPolicy(),
		NonWritable: true,
		HCL:         `policy_type = "something"`,
		ID:          ".",
	}.ExpectError(t, "policy_type must be either 'managed', 'customer' or 'restricted'")
}

func TestDataAwsCrossAccountInvalidAccountId(t *testing.T) {
	qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsCrossaccountPolicy(),
		NonWritable: true,
		HCL: `
		policy_type = "restricted"
		aws_account_id = "12345678901212"`,
		ID: ".",
	}.ExpectError(t, "aws_account_id must be a 12 digit number")
}

func TestDataAwsCrossAccountInvalidPartition(t *testing.T) {
	qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsCrossaccountPolicy(),
		NonWritable: true,
		HCL:         `aws_partition = "something"`,
		ID:          ".",
	}.ExpectError(t, AwsPartitionsValidationError)
}

func TestDataAwsCrossAccountInvalidVpcId(t *testing.T) {
	qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsCrossaccountPolicy(),
		NonWritable: true,
		HCL: `
		policy_type = "restricted"
		aws_account_id = "123456789012"
		vpc_id = "2345678"`,
		ID: ".",
	}.ExpectError(t, "vpc_id must begin with 'vpc-'")
}

func TestDataAwsCrossAccountMissingRegion(t *testing.T) {
	qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsCrossaccountPolicy(),
		NonWritable: true,
		HCL: `
		policy_type = "restricted"
		aws_account_id = "123456789012"
		vpc_id = "vpc-12345678"
		security_group_id = "sg-12345678"`,
		ID: ".",
	}.ExpectError(t, "region must be set")
}

func TestDataAwsCrossAccountInvalidSgGroup(t *testing.T) {
	qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsCrossaccountPolicy(),
		NonWritable: true,
		HCL: `
		policy_type = "restricted"
		aws_account_id = "123456789012"
		vpc_id = "vpc-12345678"
		region = "us-west-2"
		security_group_id = "12345678"`,
		ID: ".",
	}.ExpectError(t, "security_group_id must begin with 'sg-'")
}
