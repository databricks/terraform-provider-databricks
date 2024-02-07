package aws

import (
	"strings"
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestDataAwsUnityCatalogPolicy(t *testing.T) {
	d, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsUnityCatalogPolicy(),
		NonWritable: true,
		ID:          ".",
		HCL: `
        aws_account_id = "123456789098"
        bucket_name = "databricks-bucket"
        role_name = "databricks-role"
        kms_name = "databricks-kms"
        `,
	}.Apply(t)
	assert.NoError(t, err)
	j := strings.ReplaceAll(d.Get("json").(string), " ", "")
	p := strings.ReplaceAll(`{
          "Version": "2012-10-17",
          "Statement": [
            {
              "Effect": "Allow",
              "Action": [
                "s3:GetObject",
                "s3:PutObject",
                "s3:DeleteObject",
                "s3:ListBucket",
                "s3:GetBucketLocation"
              ],
              "Resource": [
                "arn:aws:s3:::databricks-bucket/*",
                "arn:aws:s3:::databricks-bucket"
              ]
            },
            {
              "Effect": "Allow",
              "Action": [
                "sts:AssumeRole"
              ],
              "Resource": [
                "arn:aws:iam::123456789098:role/databricks-role"
              ]
            },
            {
              "Effect": "Allow",
              "Action": [
                "kms:Decrypt",
                "kms:Encrypt",
                "kms:GenerateDataKey*"
              ],
              "Resource": [
                "arn:aws:kms:databricks-kms"
              ]
            }
          ]
        }`, " ", "")
	assert.EqualValues(t, j, p)
}

func TestDataAwsUnityCatalogPolicyWithoutKMS(t *testing.T) {
	d, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsUnityCatalogPolicy(),
		NonWritable: true,
		ID:          ".",
		HCL: `
        aws_account_id = "123456789098"
        bucket_name = "databricks-bucket"
        role_name = "databricks-role"
        `,
	}.Apply(t)
	assert.NoError(t, err)
	j := strings.ReplaceAll(d.Get("json").(string), " ", "")
	p := strings.ReplaceAll(`{
          "Version": "2012-10-17",
          "Statement": [
            {
              "Effect": "Allow",
              "Action": [
                "s3:GetObject",
                "s3:PutObject",
                "s3:DeleteObject",
                "s3:ListBucket",
                "s3:GetBucketLocation"
              ],
              "Resource": [
                "arn:aws:s3:::databricks-bucket/*",
                "arn:aws:s3:::databricks-bucket"
              ]
            },
            {
              "Effect": "Allow",
              "Action": [
                "sts:AssumeRole"
              ],
              "Resource": [
                "arn:aws:iam::123456789098:role/databricks-role"
              ]
            }
          ]
        }`, " ", "")
	assert.EqualValues(t, j, p)
}
