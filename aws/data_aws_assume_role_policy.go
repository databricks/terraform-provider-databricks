package aws

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
			predictedRoleArn := d.Get("predicted_role_arn").(string)
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
							"AWS": fmt.Sprintf("arn:aws:iam::%s:root", d.Get("databricks_account_id").(string)),
						},
					},
				},
			}
			if (len(predictedRoleArn) > 0) {
				rootArn := predictedRoleArn[:strings.IndexByte(predictedRoleArn, '/')] + "/root";
				elementToAdd := awsIamPolicyStatement{
					Sid: "ExplicitSelfRoleAssumption",
					Effect:  "Allow",
					Actions: "sts:AssumeRole",
					Condition: map[string]map[string]string{
						"ArnLike": {
							"aws:PrincipalArn": predictedRoleArn,
						},
					},
					Principal: map[string]string{
						"AWS": rootArn,
					},
				};
				policy.Statements = append(policy.Statements, &elementToAdd);
			}
			if v, ok := d.GetOk("for_log_delivery"); ok {
				if v.(bool) {
					// this is production UsageDelivery IAM role, that is considered a constant
					logDeliveryARN := "arn:aws:iam::414351767826:role/SaasUsageDeliveryRole-prod-IAMRole-3PLHICCRR1TK"
					policy.Statements[0].Principal["AWS"] = logDeliveryARN
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
			"databricks_account_id": {
				Type:     schema.TypeString,
				Default:  "414351767826",
				Optional: true,
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
			"predicted_role_arn": {
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Default:  "",
			},
			"json": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}
