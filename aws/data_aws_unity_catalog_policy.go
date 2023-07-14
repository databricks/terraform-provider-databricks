package aws

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
							"kms:Decrypt",
							"kms:Encrypt",
							"kms:GenerateDataKey*",
						},
						Resources: []string{
							fmt.Sprintf("arn:aws:kms:%s", kmsKey),
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
			"json": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}
