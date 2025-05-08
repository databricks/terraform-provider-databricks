package aws

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// DataAwsBucketPolicy ...
func DataAwsBucketPolicy() common.Resource {
	return common.Resource{
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			bucket := d.Get("bucket").(string)
			awsPartition := d.Get("aws_partition").(string)
			databricksAwsAccountId := AwsConfig[awsPartition]["accountId"]
			awsNamespace := AwsConfig[awsPartition]["awsNamespace"]

			if databricksAwsAccountId == "" {
				databricksAwsAccountId = AwsConfig[awsPartition]["accountId"]
			}

			policy := awsIamPolicy{
				Version: "2012-10-17",
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
							fmt.Sprintf("arn:%s:s3:::%s/*", awsNamespace, bucket),
							fmt.Sprintf("arn:%s:s3:::%s", awsNamespace, bucket),
						},
						Principal: map[string]string{
							"AWS": fmt.Sprintf("arn:%s:iam::%s:root", awsNamespace, databricksAwsAccountId),
						},
					},
				},
			}
			e2AccountId := d.Get("databricks_e2_account_id").(string)
			if e2AccountId != "" {
				policy.Statements[0].Condition = map[string]map[string]string{
					"StringEquals": {
						"aws:PrincipalTag/DatabricksAccountId": e2AccountId,
					},
				}
			}
			if v, ok := d.GetOk("full_access_role"); ok {
				policy.Statements[0].Principal["AWS"] = v.(string)
			}
			policyJSON, err := json.MarshalIndent(policy, "", "  ")
			if err != nil {
				return err
			}
			d.SetId(bucket)
			// nolint
			d.Set("json", string(policyJSON))
			return nil
		},
		Schema: map[string]*schema.Schema{
			"aws_partition": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(AwsPartitions, false),
				Default:      "aws",
			},
			"databricks_account_id": {
				Type:       schema.TypeString,
				Optional:   true,
				Deprecated: "databricks_account_id will be will be removed in the next major release.",
			},
			"databricks_e2_account_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"full_access_role": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"bucket": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringMatch(AwsBucketNameRegex, AwsBucketNameRegexError),
			},
			"json": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}
