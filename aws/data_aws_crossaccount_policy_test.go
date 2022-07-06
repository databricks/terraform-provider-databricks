package aws

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestDataAwsCrossAccountPolicy(t *testing.T) {
	d, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsCrossaccountPolicy(),
		NonWritable: true,
		ID:          ".",
	}.Apply(t)
	assert.NoError(t, err)
	j := d.Get("json")
	assert.Lenf(t, j, 2759, "Strange length for policy: %s", j)
}

func TestDataAwsCrossAccountPolicy_WithPassRoles(t *testing.T) {
	d, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsCrossaccountPolicy(),
		NonWritable: true,
		HCL:         `pass_roles = ["a", "b", "c"]`,
		ID:          ".",
	}.Apply(t)
	assert.NoError(t, err)
	j := d.Get("json")
	assert.Lenf(t, j, 2895, "Strange length for policy: %s", j)
}
