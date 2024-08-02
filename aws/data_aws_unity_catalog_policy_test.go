package aws

import (
	"encoding/json"
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
	j := d.Get("json").(string)
	p := `{
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
        }`
	compareJSON(t, j, p)
}

func TestDataAwsUnityCatalogPolicyFullKms(t *testing.T) {
	d, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsUnityCatalogPolicy(),
		NonWritable: true,
		ID:          ".",
		HCL: `
        aws_account_id = "123456789098"
        bucket_name = "databricks-bucket"
        role_name = "databricks-role"
        kms_name = "arn:aws:kms:us-west-2:111122223333:key/databricks-kms"
        `,
	}.Apply(t)
	assert.NoError(t, err)
	j := d.Get("json").(string)
	p := `{
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
                "arn:aws:kms:us-west-2:111122223333:key/databricks-kms"
              ]
            }
          ]
        }`
	compareJSON(t, j, p)
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
	j := d.Get("json").(string)
	p := `{
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
        }`
	compareJSON(t, j, p)
}

func compareJSON(t *testing.T, json1 string, json2 string) {
	var i1 interface{}
	var i2 interface{}
	err := json.Unmarshal([]byte(json1), &i1)
	assert.NoError(t, err, "error while unmarshalling")
	err = json.Unmarshal([]byte(json2), &i2)
	assert.NoError(t, err, "error while unmarshalling")
	assert.Equal(t, i1, i2)
}
