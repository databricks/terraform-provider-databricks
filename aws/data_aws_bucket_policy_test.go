package aws

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestDataAwsBucketPolicy(t *testing.T) {
	d, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsBucketPolicy(),
		NonWritable: true,
		ID:          ".",
		HCL: `
		bucket = "abc"
		`,
	}.Apply(t)
	assert.NoError(t, err)
	j := d.Get("json")
	assert.Lenf(t, j, 440, "Strange length for policy: %s", j)
}

func TestDataAwsBucketPolicy_FullAccessRole(t *testing.T) {
	d, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsBucketPolicy(),
		NonWritable: true,
		ID:          ".",
		HCL: `
		bucket = "abc"
		full_access_role = "bcd"
		`,
	}.Apply(t)
	assert.NoError(t, err)
	j := d.Get("json")
	assert.Lenf(t, j, 413, "Strange length for policy: %s", j)
}

func TestDataAwsBucketPolicyConfusedDeputyProblem(t *testing.T) {
	d, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsBucketPolicy(),
		NonWritable: true,
		ID:          ".",
		HCL: `
		bucket = "abc"
		databricks_e2_account_id = "my_e2_account_id"
		`,
	}.Apply(t)
	assert.NoError(t, err)
	j := d.Get("json")
	assert.Lenf(t, j, 575, "Strange length for policy: %s", j)
}

func TestDataAwsBucketPolicyPartitionGov(t *testing.T) {
	d, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsBucketPolicy(),
		NonWritable: true,
		ID:          ".",
		HCL: `
		bucket = "abc"
		aws_partition = "aws-us-gov"
		`,
	}.Apply(t)
	assert.NoError(t, err)
	j := d.Get("json")
	assert.Lenf(t, j, 461, "Strange length for policy: %s", j)
}

func TestDataAwsBucketPolicyPartitionGovDoD(t *testing.T) {
	d, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsBucketPolicy(),
		NonWritable: true,
		ID:          ".",
		HCL: `
		bucket = "abc"
		aws_partition = "aws-us-gov-dod"
		`,
	}.Apply(t)
	assert.NoError(t, err)
	j := d.Get("json")
	assert.Lenf(t, j, 461, "Strange length for policy: %s", j)
}
