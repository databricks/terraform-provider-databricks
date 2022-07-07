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
