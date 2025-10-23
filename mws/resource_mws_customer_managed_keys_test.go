package mws

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestResourceCustomerManagedKeyCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/abc/customer-managed-keys",
				ExpectedRequest: CustomerManagedKey{
					AccountID: "abc",
					AwsKeyInfo: &AwsKeyInfo{
						KeyArn:   "key-arn",
						KeyAlias: "key-alias",
					},
					UseCases: []string{"MANAGED_SERVICES"},
				},
				Response: CustomerManagedKey{
					CustomerManagedKeyID: "cmkid",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/customer-managed-keys/cmkid",
				Response: CustomerManagedKey{
					CustomerManagedKeyID: "cmkid",
					AwsKeyInfo: &AwsKeyInfo{
						KeyArn:    "key-arn",
						KeyAlias:  "key-alias",
						KeyRegion: "us-east-1",
					},
					AccountID:    "abc",
					UseCases:     []string{"MANAGED_SERVICES"},
					CreationTime: 123,
				},
			},
		},
		Resource: ResourceMwsCustomerManagedKeys(),
		HCL: `
			account_id = "abc"

			aws_key_info {
				key_arn   = "key-arn"
				key_alias = "key-alias"
			}
			use_cases = ["MANAGED_SERVICES"]
		`,
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc/cmkid", d.Id())
	assert.Equal(t, "key-arn", d.Get("aws_key_info.0.key_arn"))
	assert.Equal(t, "key-alias", d.Get("aws_key_info.0.key_alias"))
}

func TestResourceCustomerManagedKeyCreate_Error(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/abc/customer-managed-keys",
				ExpectedRequest: CustomerManagedKey{
					AccountID: "abc",
					AwsKeyInfo: &AwsKeyInfo{
						KeyArn:   "key-arn",
						KeyAlias: "key-alias",
					},
					UseCases: []string{"MANAGED_SERVICE"},
				},
				Response: apierr.APIError{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/customer-managed-keys/cmkid",
				Response: CustomerManagedKey{
					CustomerManagedKeyID: "cmkid",
					AwsKeyInfo: &AwsKeyInfo{
						KeyArn:    "key-arn",
						KeyAlias:  "key-alias",
						KeyRegion: "us-east-1",
					},
					UseCases:     []string{"MANAGED_SERVICE"},
					AccountID:    "abc",
					CreationTime: 123,
				},
			},
		},
		Resource: ResourceMwsCustomerManagedKeys(),
		HCL: `
			account_id = "abc"

			aws_key_info {
				key_arn   = "key-arn"
				key_alias = "key-alias"
			}
			use_cases = ["MANAGED_SERVICE"]
		`,
		Create: true,
	}.Apply(t)
	assert.Error(t, err)
}

func TestResourceCustomerManagedKeyRead(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/customer-managed-keys/cmkid",
				Response: CustomerManagedKey{
					CustomerManagedKeyID: "cmkid",
					AwsKeyInfo: &AwsKeyInfo{
						KeyArn:    "key-arn",
						KeyAlias:  "key-alias",
						KeyRegion: "us-east-1",
					},
					UseCases:     []string{"MANAGED_SERVICES"},
					AccountID:    "abc",
					CreationTime: 123,
				},
			},
		},
		Resource: ResourceMwsCustomerManagedKeys(),
		HCL: `
			account_id = "abc"

			aws_key_info {
				key_arn   = "key-arn"
				key_alias = "key-alias"
			}
			use_cases = ["MANAGED_SERVICES"]
		`,
		ID:   "abc/cmkid",
		Read: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc/cmkid", d.Id())
	assert.Equal(t, "key-arn", d.Get("aws_key_info.0.key_arn"))
	assert.Equal(t, "key-alias", d.Get("aws_key_info.0.key_alias"))
	assert.Equal(t, "us-east-1", d.Get("aws_key_info.0.key_region"))
	assert.Equal(t, []any{"MANAGED_SERVICES"}, d.Get("use_cases"))
	assert.Equal(t, "abc", d.Get("account_id"))
	assert.Equal(t, 123, d.Get("creation_time"))
}

func TestResourceCustomerManagedKeyRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/customer-managed-keys/cmkid",
				Response: apierr.APIError{
					Message: "Invalid endpoint",
				},
				Status: 404,
			},
		},
		Resource: ResourceMwsCustomerManagedKeys(),
		HCL: `
			account_id = "abc"

			aws_key_info {
				key_arn   = "key-arn"
				key_alias = "key-alias"
			}
			use_cases = ["MANAGED_SERVICES"]
		`,
		ID:      "abc/cmkid",
		Read:    true,
		Removed: true,
	}.ApplyNoError(t)
}

func TestResourceCustomerManagedKeyDelete(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "DELETE",
				Resource: "/api/2.0/accounts/abc/customer-managed-keys/cmkid",
				Status:   200,
			},
		},
		Resource: ResourceMwsCustomerManagedKeys(),
		HCL: `
			account_id = "abc"

			aws_key_info {
				key_arn   = "key-arn"
				key_alias = "key-alias"
			}
			use_cases = ["MANAGED_SERVICES"]
		`,
		ID:     "abc/cmkid",
		Delete: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc/cmkid", d.Id())
}

func TestCmkStateUpgrader(t *testing.T) {
	state, err := migrateResourceCustomerManagedKeyV0(context.Background(),
		map[string]any{}, nil)
	assert.NoError(t, err)
	_, ok := state["use_cases"]
	assert.True(t, ok)
}

func TestAwsKeyInfoKeyAliasOptional(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/abc/customer-managed-keys",
				ExpectedRequest: CustomerManagedKey{
					AccountID: "abc",
					AwsKeyInfo: &AwsKeyInfo{
						KeyArn: "key-arn",
					},
					UseCases: []string{"MANAGED_SERVICES"},
				},
				Response: CustomerManagedKey{
					CustomerManagedKeyID: "cmkid",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/customer-managed-keys/cmkid",
				Response: CustomerManagedKey{
					CustomerManagedKeyID: "cmkid",
					AwsKeyInfo: &AwsKeyInfo{
						KeyArn:    "key-arn",
						KeyRegion: "us-east-1",
					},
					AccountID:    "abc",
					UseCases:     []string{"MANAGED_SERVICES"},
					CreationTime: 123,
				},
			},
		},
		Resource: ResourceMwsCustomerManagedKeys(),
		HCL: `
			account_id = "abc"

			aws_key_info {
				key_arn   = "key-arn"
			}
			use_cases = ["MANAGED_SERVICES"]
		`,
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "abc/cmkid", d.Id())
	assert.Equal(t, "key-arn", d.Get("aws_key_info.0.key_arn"))
}

func TestResourceCustomerManagedKeyCreate_NoAccountIDInResource(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/accounts/abc/customer-managed-keys",
				ExpectedRequest: CustomerManagedKey{
					AccountID: "abc",
					AwsKeyInfo: &AwsKeyInfo{
						KeyArn:   "key-arn",
						KeyAlias: "key-alias",
					},
					UseCases: []string{"MANAGED_SERVICES"},
				},
				Response: CustomerManagedKey{
					CustomerManagedKeyID: "cmkid",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/customer-managed-keys/cmkid",
				Response: CustomerManagedKey{
					CustomerManagedKeyID: "cmkid",
					AwsKeyInfo: &AwsKeyInfo{
						KeyArn:    "key-arn",
						KeyAlias:  "key-alias",
						KeyRegion: "us-east-1",
					},
					AccountID:    "abc",
					UseCases:     []string{"MANAGED_SERVICES"},
					CreationTime: 123,
				},
			},
		},
		Resource: ResourceMwsCustomerManagedKeys(),
		HCL: `
			aws_key_info {
				key_arn   = "key-arn"
				key_alias = "key-alias"
			}
			use_cases = ["MANAGED_SERVICES"]
		`,
		AccountID: "abc",
		Create:    true,
	}.ApplyAndExpectData(t, map[string]any{
		"account_id": "abc",
		"id":         "abc/cmkid",
	})
}

func TestResourceCustomerManagedKeyCreate_NoAccountID(t *testing.T) {
	qa.ResourceFixture{
		Resource: ResourceMwsCustomerManagedKeys(),
		HCL: `
			aws_key_info {
				key_arn   = "key-arn"
				key_alias = "key-alias"
			}
			use_cases = ["MANAGED_SERVICES"]
		`,
		Create: true,
	}.ExpectError(t, "account_id is required in the provider block or in the resource")
}
