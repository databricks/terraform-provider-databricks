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
	assert.Lenf(t, j, 571, "Strange length for policy: %s", j)
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
	assert.Lenf(t, j, 544, "Strange length for policy: %s", j)
}
