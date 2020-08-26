package access

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

type awsIamPolicy struct {
	Version    string                   `json:"Version,omitempty"`
	ID         string                   `json:"Id,omitempty"`
	Statements []*awsIamPolicyStatement `json:"Statement"`
}

type awsIamPolicyStatement struct {
	Sid          string                       `json:"Sid,omitempty"`
	Effect       string                       `json:"Effect,omitempty"`
	Actions      interface{}                  `json:"Action,omitempty"`
	NotActions   interface{}                  `json:"NotAction,omitempty"`
	Resources    interface{}                  `json:"Resource,omitempty"`
	NotResources interface{}                  `json:"NotResource,omitempty"`
	Principal    map[string]string            `json:"Principal,omitempty"`
	Condition    map[string]map[string]string `json:"Condition,omitempty"`
}

func DataAwsCrossAccountRolicy() *schema.Resource {
	return &schema.Resource{
		Read: func(d *schema.ResourceData, m interface{}) error {
			policy := awsIamPolicy{
				Version: "2008-10-17",
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
			return d.Set("json", string(policyJSON))
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

func DataAwsAssumeRolePolicy() *schema.Resource {
	return &schema.Resource{
		Read: func(d *schema.ResourceData, m interface{}) error {
			externalID := d.Get("external_id").(string)
			policyJSON, err := json.MarshalIndent(awsIamPolicy{
				Version: "2008-10-17",
				Statements: []*awsIamPolicyStatement{
					{
						Effect:  "Allow",
						Actions: "sts:AssumeRole",
						Condition: map[string]map[string]string{
							"StringEquals": {
								"sts:ExternalId": externalID,
							},
						},
						Principal: map[string]string{
							"AWS": fmt.Sprintf("arn:aws:iam::%s:root", d.Get("databricks_account_id").(string)),
						},
					},
				},
			}, "", "  ")
			if err != nil {
				return err
			}
			d.SetId(externalID)
			return d.Set("json", string(policyJSON))
		},
		Schema: map[string]*schema.Schema{
			"databricks_account_id": {
				Type:     schema.TypeString,
				Default:  "414351767826",
				Optional: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"json": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func DataAwsBucketPolicy() *schema.Resource {
	return &schema.Resource{
		Read: func(d *schema.ResourceData, m interface{}) error {
			bucket := d.Get("bucket").(string)
			policy := awsIamPolicy{
				Version: "2008-10-17",
				Statements: []*awsIamPolicyStatement{
					{
						Effect: "Allow",
						Actions: []string{
							"s3:GetObject",
							"s3:GetObjectVersion",
							"s3:PutObject",
							"s3:DeleteObject",
							"s3:ListBucket",
							"s3:GetBucketLocation",
						},
						Resources: []string{
							fmt.Sprintf("arn:aws:s3:::%s/*", bucket),
							fmt.Sprintf("arn:aws:s3:::%s", bucket),
						},
						Principal: map[string]string{
							"AWS": fmt.Sprintf("arn:aws:iam::%s:root", d.Get("databricks_account_id").(string)),
						},
					},
				},
			}
			if v, ok := d.GetOk("full_access_role"); ok {
				policy.Statements[0].Principal["AWS"] = v.(string)
			}
			policyJSON, err := json.MarshalIndent(policy, "", "  ")
			if err != nil {
				return err
			}
			d.SetId(bucket)
			return d.Set("json", string(policyJSON))
		},
		Schema: map[string]*schema.Schema{
			"databricks_account_id": {
				Type:     schema.TypeString,
				Default:  "414351767826",
				Optional: true,
			},
			"full_access_role": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"bucket": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringMatch(
					regexp.MustCompile(`^[0-9a-zA-Z_-]+$`),
					"must contain only alphanumeric, underscore, and hyphen characters"),
			},
			"json": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}
