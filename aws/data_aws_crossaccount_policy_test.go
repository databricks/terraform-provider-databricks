package aws

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestDataAwsCrossAccountDatabricksManagedPolicy(t *testing.T) {
	d, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsCrossaccountPolicy(),
		NonWritable: true,
		ID:          ".",
	}.Apply(t)
	assert.NoError(t, err)
	j := d.Get("json")
	assert.Lenf(t, j, 3032, "Strange length for policy: %s", j)
}

func TestDataAwsCrossAccountCustomerManagedPolicy(t *testing.T) {
	d, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsCrossaccountPolicy(),
		NonWritable: true,
		HCL:         `policy_type = "customer"`,
		ID:          ".",
	}.Apply(t)
	assert.NoError(t, err)
	j := d.Get("json")
	assert.Lenf(t, j, 2328, "Strange length for policy: %s", j)
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
	assert.Lenf(t, j, 3168, "Strange length for policy: %s", j)
}

func TestDataAwsCrossAccountRestrictedPolicy(t *testing.T) {
	d, err := qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsCrossaccountPolicy(),
		NonWritable: true,
		HCL: `
		policy_type = "restricted"
		aws_account_id = "123456789012"
		vpc_id = "vpc-12345678"
		region = "us-west-2"
		security_group_id = "sg-12345678"`,
		ID: ".",
	}.Apply(t)
	assert.NoError(t, err)
	j := d.Get("json")
	assert.Lenf(t, j, 5725, "Strange length for policy: %s", j)
}

func TestDataAwsCrossAccountInvalidPolicy(t *testing.T) {
	qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsCrossaccountPolicy(),
		NonWritable: true,
		HCL:         `policy_type = "something"`,
		ID:          ".",
	}.ExpectError(t, "policy_type must be either 'managed', 'customer' or 'restricted'")
}

func TestDataAwsCrossAccountInvalidAccountId(t *testing.T) {
	qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsCrossaccountPolicy(),
		NonWritable: true,
		HCL: `
		policy_type = "restricted"
		aws_account_id = "12345678901212"`,
		ID: ".",
	}.ExpectError(t, "aws_account_id must be a 12 digit number")
}

func TestDataAwsCrossAccountInvalidVpcId(t *testing.T) {
	qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsCrossaccountPolicy(),
		NonWritable: true,
		HCL: `
		policy_type = "restricted"
		aws_account_id = "123456789012"
		vpc_id = "2345678"`,
		ID: ".",
	}.ExpectError(t, "vpc_id must begin with 'vpc-'")
}

func TestDataAwsCrossAccountMissingRegion(t *testing.T) {
	qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsCrossaccountPolicy(),
		NonWritable: true,
		HCL: `
		policy_type = "restricted"
		aws_account_id = "123456789012"
		vpc_id = "vpc-12345678"
		security_group_id = "sg-12345678"`,
		ID: ".",
	}.ExpectError(t, "region must be set")
}

func TestDataAwsCrossAccountInvalidSgGroup(t *testing.T) {
	qa.ResourceFixture{
		Read:        true,
		Resource:    DataAwsCrossaccountPolicy(),
		NonWritable: true,
		HCL: `
		policy_type = "restricted"
		aws_account_id = "123456789012"
		vpc_id = "vpc-12345678"
		region = "us-west-2"
		security_group_id = "12345678"`,
		ID: ".",
	}.ExpectError(t, "security_group_id must begin with 'sg-'")
}
