package access

import (
	"testing"

	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
)

func TestDataAwsCrossAccountRolicy(t *testing.T) {
	d, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsCrossAccountRolicy(),
		NonWritable: true,
		ID:          ".",
	}.Apply(t)
	assert.NoError(t, err)
	j := d.Get("json")
	assert.Lenf(t, j, 2759, "Strange length for policy: %s", j)
}

func TestDataAwsCrossAccountRolicy_WithPassRoles(t *testing.T) {
	d, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsCrossAccountRolicy(),
		NonWritable: true,
		HCL:         `pass_roles = ["a", "b", "c"]`,
		ID:          ".",
	}.Apply(t)
	assert.NoError(t, err)
	j := d.Get("json")
	assert.Lenf(t, j, 2895, "Strange length for policy: %s", j)
}

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
