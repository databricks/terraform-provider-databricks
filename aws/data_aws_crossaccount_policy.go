package aws

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"slices"

	"github.com/databricks/terraform-provider-databricks/common"
)

// DataAwsCrossaccountPolicy defines the cross-account policy
func DataAwsCrossaccountPolicy() common.Resource {
	type AwsCrossAccountPolicy struct {
		PolicyType      string   `json:"policy_type,omitempty" tf:"default:managed"`
		PassRole        []string `json:"pass_roles,omitempty"`
		JSON            string   `json:"json" tf:"computed"`
		AwsAccountId    string   `json:"aws_account_id,omitempty"`
		AwsPartition    string   `json:"aws_partition,omitempty" tf:"default:aws"`
		VpcId           string   `json:"vpc_id,omitempty"`
		Region          string   `json:"region,omitempty"`
		SecurityGroupId string   `json:"security_group_id,omitempty"`
	}
	return common.NoClientData(func(ctx context.Context, data *AwsCrossAccountPolicy) error {
		if !slices.Contains(AwsPartitions, data.AwsPartition) {
			return errors.New(AwsPartitionsValidationError)
		}

		if !slices.Contains([]string{"managed", "customer", "restricted"}, data.PolicyType) {
			return fmt.Errorf("policy_type must be either 'managed', 'customer' or 'restricted'")
		}

		if data.PolicyType == "restricted" {
			match, _ := regexp.MatchString(`^\d{12}$`, data.AwsAccountId)
			if !match {
				return fmt.Errorf("aws_account_id must be a 12 digit number")
			}
			match, _ = regexp.MatchString(`^vpc-.*$`, data.VpcId)
			if !match {
				return fmt.Errorf("vpc_id must begin with 'vpc-'")
			}
			if data.Region == "" {
				return fmt.Errorf("region must be set")
			}
			match, _ = regexp.MatchString(`^sg-.*$`, data.SecurityGroupId)
			if !match {
				return fmt.Errorf("security_group_id must begin with 'sg-'")
			}
		}
		awsNamespace := AwsConfig[data.AwsPartition]["awsNamespace"]
		// non resource-based permissions
		actions := []string{
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
		}
		if data.PolicyType != "restricted" {
			actions = append(actions, []string{
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
			}...)
		}
		// additional permissions for Databricks-managed VPC policy
		if data.PolicyType == "managed" {
			actions = append(actions, []string{
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
				"ec2:ReleaseAddress",
			}...)
		}
		policy := awsIamPolicy{
			Version: "2012-10-17",
			Statements: []*awsIamPolicyStatement{
				{
					Effect:    "Allow",
					Actions:   actions,
					Resources: "*",
				},
				{
					Effect: "Allow",
					Actions: []string{
						"iam:CreateServiceLinkedRole",
						"iam:PutRolePolicy",
					},
					Resources: fmt.Sprintf("arn:%s:iam::*:role/aws-service-role/spot.amazonaws.com/AWSServiceRoleForEC2Spot", awsNamespace),
					Condition: map[string]map[string]string{
						"StringLike": {
							"iam:AWSServiceName": "spot.amazonaws.com",
						},
					},
				},
			},
		}
		// pass role permissions
		if len(data.PassRole) > 0 {
			policy.Statements = append(policy.Statements,
				&awsIamPolicyStatement{
					Effect:    "Allow",
					Actions:   "iam:PassRole",
					Resources: data.PassRole,
				})
		}

		// resource-based permissions
		if data.PolicyType == "restricted" {
			region := data.Region
			aws_account_id := data.AwsAccountId
			vpc_id := data.VpcId
			security_group_id := data.SecurityGroupId
			policy.Statements = append(policy.Statements,
				&awsIamPolicyStatement{
					Sid:    "InstancePoolsSupport",
					Effect: "Allow",
					Actions: []string{
						"ec2:AssociateIamInstanceProfile",
						"ec2:DisassociateIamInstanceProfile",
						"ec2:ReplaceIamInstanceProfileAssociation",
					},
					Resources: fmt.Sprintf("arn:%s:ec2:%s:%s:instance/*", awsNamespace, region, aws_account_id),
					Condition: map[string]map[string]string{
						"StringEquals": {
							"ec2:ResourceTag/Vendor": "Databricks",
						},
					},
				},
				&awsIamPolicyStatement{
					Sid:     "AllowEc2RunInstancePerTag",
					Effect:  "Allow",
					Actions: "ec2:RunInstances",
					Resources: []string{
						fmt.Sprintf("arn:%s:ec2:%s:%s:volume/*", awsNamespace, region, aws_account_id),
						fmt.Sprintf("arn:%s:ec2:%s:%s:instance/*", awsNamespace, region, aws_account_id),
					},
					Condition: map[string]map[string]string{
						"StringEquals": {
							"aws:RequestTag/Vendor": "Databricks",
						},
					},
				},
				&awsIamPolicyStatement{
					Sid:       "AllowEc2RunInstanceImagePerTag",
					Effect:    "Allow",
					Actions:   "ec2:RunInstances",
					Resources: fmt.Sprintf("arn:%s:ec2:%s:%s:image/*", awsNamespace, region, aws_account_id),
					Condition: map[string]map[string]string{
						"StringEquals": {
							"aws:ResourceTag/Vendor": "Databricks",
						},
					},
				},
				&awsIamPolicyStatement{
					Sid:     "AllowEc2RunInstancePerVPCid",
					Effect:  "Allow",
					Actions: "ec2:RunInstances",
					Resources: []string{
						fmt.Sprintf("arn:%s:ec2:%s:%s:network-interface/*", awsNamespace, region, aws_account_id),
						fmt.Sprintf("arn:%s:ec2:%s:%s:subnet/*", awsNamespace, region, aws_account_id),
						fmt.Sprintf("arn:%s:ec2:%s:%s:security-group/*", awsNamespace, region, aws_account_id),
					},
					Condition: map[string]map[string]string{
						"StringEquals": {
							"ec2:vpc": fmt.Sprintf("arn:%s:ec2:%s:%s:vpc/%s", awsNamespace, region, aws_account_id, vpc_id),
						},
					},
				},
				&awsIamPolicyStatement{
					Sid:     "AllowEc2RunInstanceOtherResources",
					Effect:  "Allow",
					Actions: "ec2:RunInstances",
					NotResources: []string{
						fmt.Sprintf("arn:%s:ec2:%s:%s:image/*", awsNamespace, region, aws_account_id),
						fmt.Sprintf("arn:%s:ec2:%s:%s:network-interface/*", awsNamespace, region, aws_account_id),
						fmt.Sprintf("arn:%s:ec2:%s:%s:subnet/*", awsNamespace, region, aws_account_id),
						fmt.Sprintf("arn:%s:ec2:%s:%s:security-group/*", awsNamespace, region, aws_account_id),
						fmt.Sprintf("arn:%s:ec2:%s:%s:volume/*", awsNamespace, region, aws_account_id),
						fmt.Sprintf("arn:%s:ec2:%s:%s:instance/*", awsNamespace, region, aws_account_id),
					},
				},
				&awsIamPolicyStatement{
					Sid:       "EC2TerminateInstancesTag",
					Effect:    "Allow",
					Actions:   "ec2:TerminateInstances",
					Resources: fmt.Sprintf("arn:%s:ec2:%s:%s:instance/*", awsNamespace, region, aws_account_id),
					Condition: map[string]map[string]string{
						"StringEquals": {
							"ec2:ResourceTag/Vendor": "Databricks",
						},
					},
				},
				&awsIamPolicyStatement{
					Sid:    "EC2AttachDetachVolumeTag",
					Effect: "Allow",
					Actions: []string{
						"ec2:AttachVolume",
						"ec2:DetachVolume",
					},
					Resources: []string{
						fmt.Sprintf("arn:%s:ec2:%s:%s:instance/*", awsNamespace, region, aws_account_id),
						fmt.Sprintf("arn:%s:ec2:%s:%s:volume/*", awsNamespace, region, aws_account_id),
					},
					Condition: map[string]map[string]string{
						"StringEquals": {
							"ec2:ResourceTag/Vendor": "Databricks",
						},
					},
				},
				&awsIamPolicyStatement{
					Sid:       "EC2CreateVolumeByTag",
					Effect:    "Allow",
					Actions:   "ec2:CreateVolume",
					Resources: fmt.Sprintf("arn:%s:ec2:%s:%s:volume/*", awsNamespace, region, aws_account_id),
					Condition: map[string]map[string]string{
						"StringEquals": {
							"aws:RequestTag/Vendor": "Databricks",
						},
					},
				},
				&awsIamPolicyStatement{
					Sid:     "EC2DeleteVolumeByTag",
					Effect:  "Allow",
					Actions: "ec2:DeleteVolume",
					Resources: []string{
						fmt.Sprintf("arn:%s:ec2:%s:%s:volume/*", awsNamespace, region, aws_account_id),
					},
					Condition: map[string]map[string]string{
						"StringEquals": {
							"ec2:ResourceTag/Vendor": "Databricks",
						},
					},
				},
				&awsIamPolicyStatement{
					Sid:    "VpcNonresourceSpecificActions",
					Effect: "Allow",
					Actions: []string{
						"ec2:AuthorizeSecurityGroupEgress",
						"ec2:AuthorizeSecurityGroupIngress",
						"ec2:RevokeSecurityGroupEgress",
						"ec2:RevokeSecurityGroupIngress",
					},
					Resources: fmt.Sprintf("arn:%s:ec2:%s:%s:security-group/%s", awsNamespace, region, aws_account_id, security_group_id),
					Condition: map[string]map[string]string{
						"StringEquals": {
							"ec2:vpc": fmt.Sprintf("arn:%s:ec2:%s:%s:vpc/%s", awsNamespace, region, aws_account_id, vpc_id),
						},
					},
				},
			)
		}
		policyJSON, err := json.MarshalIndent(policy, "", "  ")
		if err != nil {
			return err
		}
		data.JSON = string(policyJSON)
		return nil
	})
}
