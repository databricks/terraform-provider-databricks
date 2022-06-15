package aws

import (
	"context"
	"encoding/json"
	
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataAwsCrossAccountPolicy defines the cross-account policy
func DataAwsCrossAccountPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			policy := awsIamPolicy{
				Version: "2012-10-17",
				Statements: []*awsIamPolicyStatement{
					{
						Effect: "Allow",
						Actions: []string{
							"ec2:AssociateDhcpOptions",
							"ec2:AssociateIamInstanceProfile",
							"ec2:AssociateRouteTable",
							"ec2:AttachInternetGateway",
							"ec2:AttachVolume",
							"ec2:AuthorizeSecurityGroupEgress",
							"ec2:AuthorizeSecurityGroupIngress",
							"ec2:CancelSpotInstanceRequests",
							"ec2:CreateDhcpOptions",
							"ec2:CreateInternetGateway",
							"ec2:CreateKeyPair",
							"ec2:CreateRoute",
							"ec2:CreateSecurityGroup",
							"ec2:CreateSubnet",
							"ec2:CreateTags",
							"ec2:CreateVolume",
							"ec2:CreateVpc",
							"ec2:DeleteInternetGateway",
							"ec2:DeleteKeyPair",
							"ec2:DeleteRoute",
							"ec2:DeleteRouteTable",
							"ec2:DeleteSecurityGroup",
							"ec2:DeleteSubnet",
							"ec2:DeleteTags",
							"ec2:DeleteVolume",
							"ec2:DeleteVpc",
							"ec2:DescribeAvailabilityZones",
							"ec2:DescribeNetworkAcls",
							"ec2:DescribeInternetGateways",
							"ec2:DescribeVpcAttribute",
							"ec2:DescribeIamInstanceProfileAssociations",
							"ec2:DescribeInstanceStatus",
							"ec2:DescribeInstances",
							"ec2:DescribePrefixLists",
							"ec2:DescribeReservedInstancesOfferings",
							"ec2:DescribeRouteTables",
							"ec2:DescribeSecurityGroups",
							"ec2:DescribeSpotInstanceRequests",
							"ec2:DescribeSpotPriceHistory",
							"ec2:DescribeSubnets",
							"ec2:DescribeVolumes",
							"ec2:DescribeVpcs",
							"ec2:DetachInternetGateway",
							"ec2:DisassociateIamInstanceProfile",
							"ec2:ModifyVpcAttribute",
							"ec2:ReplaceIamInstanceProfileAssociation",
							"ec2:RequestSpotInstances",
							"ec2:RevokeSecurityGroupEgress",
							"ec2:RevokeSecurityGroupIngress",
							"ec2:RunInstances",
							"ec2:TerminateInstances",
							"ec2:CreatePlacementGroup",
							"ec2:DeletePlacementGroup",
							"ec2:DescribePlacementGroups",
							"ec2:AllocateAddress",
							"ec2:CreateNatGateway",
							"ec2:CreateRouteTable",
							"ec2:CreateVpcEndpoint",
							"ec2:DeleteDhcpOptions",
							"ec2:DeleteNatGateway",
							"ec2:DeleteVpcEndpoints",
							"ec2:DescribeNatGateways",
							"ec2:DisassociateRouteTable",
							"ec2:ReleaseAddress",
							"ec2:DetachVolume",
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
				return diag.FromErr(err)
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
	}
}
