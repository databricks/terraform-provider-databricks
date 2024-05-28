package aws

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestDataAwsUnityCatalogAssumeRolePolicy(t *testing.T) {
	d, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsUnityCatalogAssumeRolePolicy(),
		NonWritable: true,
		ID:          ".",
		HCL: `
        aws_account_id = "123456789098"
        unity_catalog_iam_arn = "arn:aws:iam::414351767826:role/unity-catalog-prod-UCMasterRole-14S5ZJVKOTYTL"
        role_name = "databricks-role"
        external_id = "12345"
        `,
	}.Apply(t)
	assert.NoError(t, err)
	j := d.Get("json").(string)
	p := `{
          "Version": "2012-10-17",
          "Statement": [
            {
              "Sid": "UnityCatalogAssumeRole",
              "Effect": "Allow",
              "Action": "sts:AssumeRole",
              "Principal": {
                "AWS": "arn:aws:iam::414351767826:role/unity-catalog-prod-UCMasterRole-14S5ZJVKOTYTL"
              },
              "Condition": {
                "StringEquals": {
                  "sts:ExternalId": "12345"
                }
              }
            },
            {
              "Sid": "ExplicitSelfRoleAssumption",
              "Effect": "Allow",
              "Action": "sts:AssumeRole",
              "Principal": {
                "AWS": "arn:aws:iam::123456789098:root"
              },
              "Condition": {
                "ArnLike": {
                  "aws:PrincipalArn": "arn:aws:iam::123456789098:role/databricks-role"
                }
              }
            }
          ]
        }`
	compareJSON(t, j, p)
}

func TestDataAwsUnityCatalogAssumeRolePolicyWithoutUcArn(t *testing.T) {
	d, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsUnityCatalogAssumeRolePolicy(),
		NonWritable: true,
		ID:          ".",
		HCL: `
        aws_account_id = "123456789098"
        role_name = "databricks-role"
        external_id = "12345"
        `,
	}.Apply(t)
	assert.NoError(t, err)
	j := d.Get("json").(string)
	p := `{
          "Version": "2012-10-17",
          "Statement": [
            {
              "Sid": "UnityCatalogAssumeRole",
              "Effect": "Allow",
              "Action": "sts:AssumeRole",
              "Principal": {
                "AWS": "arn:aws:iam::414351767826:role/unity-catalog-prod-UCMasterRole-14S5ZJVKOTYTL"
              },
              "Condition": {
                "StringEquals": {
                  "sts:ExternalId": "12345"
                }
              }
            },
            {
              "Sid": "ExplicitSelfRoleAssumption",
              "Effect": "Allow",
              "Action": "sts:AssumeRole",
              "Principal": {
                "AWS": "arn:aws:iam::123456789098:root"
              },
              "Condition": {
                "ArnLike": {
                  "aws:PrincipalArn": "arn:aws:iam::123456789098:role/databricks-role"
                }
              }
            }
          ]
        }`
	compareJSON(t, j, p)
}
