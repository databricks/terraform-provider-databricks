package aws

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestDataAwsAssumeRolePolicy(t *testing.T) {
	d, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsAssumeRolePolicy(),
		NonWritable: true,
		ID:          ".",
		HCL:         `external_id = "abc"`,
	}.Apply(t)
	assert.NoError(t, err)
	j := d.Get("json")
	assert.Lenf(t, j, 299, "Strange length for policy: %s", j)
}

func TestDataAwsAssumeRolePolicyGov(t *testing.T) {
	d, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsAssumeRolePolicy(),
		NonWritable: true,
		ID:          ".",
		HCL: `
		aws_partition = "aws-us-gov"
		external_id = "abc"
		`,
	}.Apply(t)
	assert.NoError(t, err)
	j := d.Get("json")
	assert.Lenf(t, j, 306, "Strange length for policy: %s", j)
}

func TestDataAwsAssumeRolePolicyLogDelivery(t *testing.T) {
	d, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsAssumeRolePolicy(),
		NonWritable: true,
		ID:          ".",
		HCL: `
		external_id = "abc"
		for_log_delivery = true
		`,
	}.Apply(t)
	assert.NoError(t, err)
	j := d.Get("json")
	assert.Lenf(t, j, 347, "Strange length for policy: %s", j)
}

func TestDataAwsAssumeRolePolicyLogDeliveryGov(t *testing.T) {
	d, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsAssumeRolePolicy(),
		NonWritable: true,
		ID:          ".",
		HCL: `
		aws_partition = "aws-us-gov"
		external_id = "abc"
		for_log_delivery = true
		`,
	}.Apply(t)
	assert.NoError(t, err)
	j := d.Get("json")
	assert.Lenf(t, j, 362, "Strange length for policy: %s", j)
}
