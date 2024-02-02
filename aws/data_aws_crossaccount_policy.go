package aws

import (
	"context"
	"encoding/json"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataAwsCrossaccountPolicy defines the cross-account policy
func DataAwsCrossaccountPolicy() common.Resource {
	return common.Resource{
		Read: func(ctx context.Context, d *schema.ResourceData, m *common.DatabricksClient) error {
			policy := awsIamPolicy{
				Version: "2012-10-17",
				Statements: []*awsIamPolicyStatement{
					{
						Effect: "Allow",
						Actions: []string{
							"ec2:AllocateAddress",
							"ec2:AssignPrivateIpAddresses",
							"ec2:AssociateDhcpOptions",
							"ec2:AssociateIamInstanceProfile",
							"ec2:AssociateRouteTable",
							"ec2:AttachInternetGateway",
							"ec2:AttachVolume",
							"ec2:AuthorizeSecurityGroupEgress",
							"ec2:AuthorizeSecurityGroupIngress",
							"ec2:CancelSpotInstanceRequests",
							"ec2:CreateDhcpOptions",
							"ec2:CreateFleet",
							"ec2:CreateInternetGateway",
							"ec2:CreateKeyPair",
							"ec2:CreateLaunchTemplate",
							"ec2:CreateLaunchTemplateVersion",
							"ec2:CreateNatGateway",
							"ec2:CreatePlacementGroup",
							"ec2:CreateRoute",
							"ec2:CreateRouteTable",
							"ec2:CreateSecurityGroup",
							"ec2:CreateSubnet",
							"ec2:CreateTags",
							"ec2:CreateVolume",
							"ec2:CreateVpc",
							"ec2:CreateVpcEndpoint",
							"ec2:DeleteDhcpOptions",
							"ec2:DeleteFleets",
							"ec2:DeleteInternetGateway",
							"ec2:DeleteKeyPair",
							"ec2:DeleteLaunchTemplate",
							"ec2:DeleteLaunchTemplateVersions",
							"ec2:DeleteNatGateway",
							"ec2:DeletePlacementGroup",
							"ec2:DeleteRoute",
							"ec2:DeleteRouteTable",
							"ec2:DeleteSecurityGroup",
							"ec2:DeleteSubnet",
							"ec2:DeleteTags",
							"ec2:DeleteVolume",
							"ec2:DeleteVpc",
							"ec2:DeleteVpcEndpoints",
							"ec2:DescribeAvailabilityZones",
							"ec2:DescribeFleetHistory",
							"ec2:DescribeFleetInstances",
							"ec2:DescribeFleets",
							"ec2:DescribeIamInstanceProfileAssociations",
							"ec2:DescribeInstanceStatus",
							"ec2:DescribeInstances",
							"ec2:DescribeInternetGateways",
							"ec2:DescribeLaunchTemplates",
							"ec2:DescribeLaunchTemplateVersions",
							"ec2:DescribeNatGateways",
							"ec2:DescribeNetworkAcls",
							"ec2:DescribePlacementGroups",
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
							"ec2:DetachInternetGateway",
							"ec2:DetachVolume",
							"ec2:DisassociateIamInstanceProfile",
							"ec2:DisassociateRouteTable",
							"ec2:GetLaunchTemplateData",
							"ec2:GetSpotPlacementScores",
							"ec2:ModifyFleet",
							"ec2:ModifyLaunchTemplate",
							"ec2:ModifyVpcAttribute",
							"ec2:ReleaseAddress",
							"ec2:ReplaceIamInstanceProfileAssociation",
							"ec2:RequestSpotInstances",
							"ec2:RevokeSecurityGroupEgress",
							"ec2:RevokeSecurityGroupIngress",
							"ec2:RunInstances",
							"ec2:TerminateInstances",
						},
						Resources: "*",
					},
					{
						Effect: "Allow",
						Actions: []string{
							"iam:CreateServiceLinkedRole",
							"iam:PutRolePolicy",
						},
						Resources: "arn:aws:iam::*:role/aws-service-role/spot.amazonaws.com/AWSServiceRoleForEC2Spot",
						Condition: map[string]map[string]string{
							"StringLike": {
								"iam:AWSServiceName": "spot.amazonaws.com",
							},
						},
					},
				},
			}
			if passRoleARNs, ok := d.GetOk("pass_roles"); ok {
				policy.Statements = append(policy.Statements, &awsIamPolicyStatement{
					Effect:    "Allow",
					Actions:   "iam:PassRole",
					Resources: passRoleARNs,
				})
			}
			policyJSON, err := json.MarshalIndent(policy, "", "  ")
			if err != nil {
				return err
			}
			d.SetId("cross-account")
			// nolint
			d.Set("json", string(policyJSON))
			return nil
		},
		Schema: map[string]*schema.Schema{
			"pass_roles": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"json": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
		},
		StrictProviderLevelResource: true,
	}
}
