package aws

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func generateReadContext(ctx context.Context, d *schema.ResourceData, m *common.DatabricksClient) error {
	bucket := d.Get("bucket_name").(string)
	awsAccountId := d.Get("aws_account_id").(string)
	awsPartition := d.Get("aws_partition").(string)
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
				},
				Resources: []string{
					fmt.Sprintf("arn:%s:s3:::%s/*", awsPartition, bucket),
					fmt.Sprintf("arn:%s:s3:::%s", awsPartition, bucket),
				},
			},
			{
				Effect: "Allow",
				Actions: []string{
					"sts:AssumeRole",
				},
				Resources: []string{
					fmt.Sprintf("arn:%s:iam::%s:role/%s", awsPartition, awsAccountId, roleName),
				},
			},
		},
	}
	if kmsKey, ok := d.GetOk("kms_name"); ok {
		kmsArn := fmt.Sprintf("arn:%s:kms:%s", awsPartition, kmsKey)
		if strings.HasPrefix(kmsKey.(string), fmt.Sprintf("arn:%s", awsPartition)) {
			kmsArn = kmsKey.(string)
		}
		policy.Statements = append(policy.Statements, &awsIamPolicyStatement{
			Effect: "Allow",
			Actions: []string{
				"kms:Decrypt",
				"kms:Encrypt",
				"kms:GenerateDataKey*",
			},
			Resources: []string{kmsArn},
		})
	}
	policy.Statements = append(policy.Statements, &awsIamPolicyStatement{
		Sid:    "ManagedFileEventsSetupStatement",
		Effect: "Allow",
		Actions: []string{
			"s3:GetBucketNotification",
			"s3:PutBucketNotification",
			"sns:ListSubscriptionsByTopic",
			"sns:GetTopicAttributes",
			"sns:SetTopicAttributes",
			"sns:CreateTopic",
			"sns:TagResource",
			"sns:Publish",
			"sns:Subscribe",
			"sqs:CreateQueue",
			"sqs:DeleteMessage",
			"sqs:ReceiveMessage",
			"sqs:SendMessage",
			"sqs:GetQueueUrl",
			"sqs:GetQueueAttributes",
			"sqs:SetQueueAttributes",
			"sqs:TagQueue",
			"sqs:ChangeMessageVisibility",
			"sqs:PurgeQueue",
		},
		Resources: []string{
			fmt.Sprintf("arn:%s:s3:::%s", awsPartition, bucket),
			fmt.Sprintf("arn:%s:sqs:*:*:csms-*", awsPartition),
			fmt.Sprintf("arn:%s:sns:*:*:csms-*", awsPartition),
		},
	},
		&awsIamPolicyStatement{
			Sid:    "ManagedFileEventsListStatement",
			Effect: "Allow",
			Actions: []string{
				"sqs:ListQueues",
				"sqs:ListQueueTags",
				"sns:ListTopics",
			},
			Resources: []string{
				fmt.Sprintf("arn:%s:sqs:*:*:csms-*", awsPartition),
				fmt.Sprintf("arn:%s:sns:*:*:csms-*", awsPartition),
			},
		},
		&awsIamPolicyStatement{
			Sid:    "ManagedFileEventsTeardownStatement",
			Effect: "Allow",
			Actions: []string{
				"sns:Unsubscribe",
				"sns:DeleteTopic",
				"sqs:DeleteQueue",
			},
			Resources: []string{
				fmt.Sprintf("arn:%s:sqs:*:*:csms-*", awsPartition),
				fmt.Sprintf("arn:%s:sns:*:*:csms-*", awsPartition),
			},
		})
	policyJSON, err := json.MarshalIndent(policy, "", "  ")
	if err != nil {
		return err
	}
	d.SetId(fmt.Sprintf("%s-%s-%s", bucket, awsAccountId, roleName))
	err = d.Set("json", string(policyJSON))
	if err != nil {
		return err
	}
	return nil
}

func validateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"kms_name": {
			Type:     schema.TypeString,
			Optional: true,
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
		"aws_partition": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice(AwsPartitions, false),
			Default:      "aws",
		},
		"json": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func DataAwsUnityCatalogPolicy() common.Resource {
	return common.Resource{
		Read:   generateReadContext,
		Schema: validateSchema(),
	}
}
