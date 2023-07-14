package aws

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"regexp"
)

func DataAwsUnityCatalogPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
			bucket := d.Get("bucket_name").(string)
			kmsKey := d.Get("kms_name").(string)
			awsAccountId := d.Get("aws_account_id").(string)
			roleName := d.Get("role_name").(string)
			policy := awsIamPolicy{
				Version: "2012-10-17",
				Statements: []*awsIamPolicyStatement{
					{
						Effect: "Allow",
						Actions: []string{
							"s3:GetObject",
							"s3:PutObject",
							"s3:DeleteObject",
							"s3:ListBucket",
							"s3:GetBucketLocation",
							"s3:GetLifecycleConfiguration",
							"s3:PutLifecycleConfiguration",
						},
						Resources: []string{
							fmt.Sprintf("arn:aws:s3:::%s/*", bucket),
							fmt.Sprintf("arn:aws:s3:::%s", bucket),
						},
					},
					{
						Effect: "Allow",
						Actions: []string{
							"sts:AssumeRole",
						},
						Resources: []string{
							fmt.Sprintf("arn:aws:iam::%s:role/%s", awsAccountId, roleName),
						},
					},
				},
			}
			if _, ok := d.GetOk("kms_name"); ok {
				policy.Statements = append(policy.Statements, &awsIamPolicyStatement{
					Effect: "Allow",
					Actions: []string{
						"kms:Decrypt",
						"kms:Encrypt",
						"kms:GenerateDataKey*",
					},
					Resources: []string{
						fmt.Sprintf("arn:aws:kms:%s", kmsKey),
					},
				})
			}
			policyJSON, err := json.MarshalIndent(policy, "", "  ")
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId("unity-catalog")
			// nolint
			d.Set("json", string(policyJSON))
			return nil
		},
		Schema: map[string]*schema.Schema{
			"kms_name": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringMatch(
					regexp.MustCompile(`^[0-9a-zA-Z/_-]+$`),
					"must contain only alphanumeric, hyphens, forward slashes, and underscores characters"),
			},
			"bucket_name": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringMatch(
					regexp.MustCompile(`^[0-9a-zA-Z_-]+$`),
					"must contain only alphanumeric, underscore, and hyphen characters"),
			},
			"role_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"aws_account_id": {
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
