package aws

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

type awsIamPolicy struct {
	Version    string                   `json:"Version,omitempty"`
	ID         string                   `json:"Id,omitempty"`
	Statements []*awsIamPolicyStatement `json:"Statement"`
}

type awsIamPolicyStatement struct {
	Sid          string                       `json:"Sid,omitempty"`
	Effect       string                       `json:"Effect,omitempty"`
	Actions      any                          `json:"Action,omitempty"`
	NotActions   any                          `json:"NotAction,omitempty"`
	Resources    any                          `json:"Resource,omitempty"`
	NotResources any                          `json:"NotResource,omitempty"`
	Principal    map[string]string            `json:"Principal,omitempty"`
	Condition    map[string]map[string]string `json:"Condition,omitempty"`
}

// DataAwsAssumeRolePolicy ...
func DataAwsAssumeRolePolicy() common.Resource {
	return common.Resource{
		Read: func(ctx context.Context, d *schema.ResourceData, m *common.DatabricksClient) error {
			externalID := d.Get("external_id").(string)
			awsPartition := d.Get("aws_partition").(string)
			databricksAwsAccountId := d.Get("databricks_account_id").(string)

			if databricksAwsAccountId == "" {
				databricksAwsAccountId = AwsConfig[awsPartition]["accountId"]
			}

			policy := awsIamPolicy{
				Version: "2012-10-17",
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
							"AWS": fmt.Sprintf("arn:%s:iam::%s:root", awsPartition, databricksAwsAccountId),
						},
					},
				},
			}
			if v, ok := d.GetOk("for_log_delivery"); ok {
				if v.(bool) {
					policy.Statements[0].Principal["AWS"] = AwsConfig[awsPartition]["logDeliveryIamArn"]
				}
			}
			policyJSON, err := json.MarshalIndent(policy, "", "  ")
			if err != nil {
				return err
			}
			d.SetId(externalID)
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
			"for_log_delivery": {
				Type:        schema.TypeBool,
				Description: "Grant AssumeRole to Databricks SaasUsageDeliveryRole instead of root account",
				Optional:    true,
				Default:     false,
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
