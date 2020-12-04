package mws

import (
	"context"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
)

func TestMwsAccCustomerManagedKeys(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	acctID := qa.GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID")
	kmsKeyArn := qa.GetEnvOrSkipTest(t, "TEST_KMS_KEY_ARN")
	kmsKeyAlias := qa.GetEnvOrSkipTest(t, "TEST_KMS_KEY_ALIAS")
	client := common.CommonEnvironmentClient()
	cmkApi := NewCustomerManagedKeysAPI(context.Background(), client)
	cmkList, err := cmkApi.List(acctID)
	assert.NoError(t, err, err)
	t.Log(cmkList)

	keyInfo, err := cmkApi.Create(CustomerManagedKey{
		AwsKeyInfo: &AwsKeyInfo{
			KeyArn:   kmsKeyArn,
			KeyAlias: kmsKeyAlias,
		},
		AccountID: acctID,
	})
	assert.NoError(t, err, err)

	keyID := keyInfo.CustomerManagedKeyID

	defer func() {
		err := cmkApi.Delete(acctID, keyID)
		assert.NoError(t, err, err)
	}()

	getKeyInfo, err := cmkApi.Read(acctID, keyID)
	assert.NoError(t, err, err)
	assert.NotNil(t, getKeyInfo, "key info should not be nil")
}

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
					CreationTime: 123,
				},
			},
		},
		Resource: ResourceCustomerManagedKey(),
		HCL: `
			account_id = "abc"

			aws_key_info {
				key_arn   = "key-arn"
				key_alias = "key-alias"
			}
		`,
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
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
				},
				Response: common.APIErrorBody{
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
					AccountID:    "abc",
					CreationTime: 123,
				},
			},
		},
		Resource: ResourceCustomerManagedKey(),
		HCL: `
			account_id = "abc"

			aws_key_info {
				key_arn   = "key-arn"
				key_alias = "key-alias"
			}
		`,
		Create: true,
	}.Apply(t)
	assert.Error(t, err, err)
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
					AccountID:    "abc",
					CreationTime: 123,
				},
			},
		},
		Resource: ResourceCustomerManagedKey(),
		HCL: `
			account_id = "abc"

			aws_key_info {
				key_arn   = "key-arn"
				key_alias = "key-alias"
			}
		`,
		ID:   "abc/cmkid",
		Read: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/cmkid", d.Id())
	assert.Equal(t, "key-arn", d.Get("aws_key_info.0.key_arn"))
	assert.Equal(t, "key-alias", d.Get("aws_key_info.0.key_alias"))
	assert.Equal(t, "us-east-1", d.Get("aws_key_info.0.key_region"))
	assert.Equal(t, "abc", d.Get("account_id"))
	assert.Equal(t, 123, d.Get("creation_time"))
	assert.Equal(t, "cmkid", d.Get("customer_managed_key_id"))
}

func TestResourceCustomerManagedKeyRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/customer-managed-keys/cmkid",
				Response: common.APIErrorBody{
					Message: "Invalid endpoint",
				},
				Status: 404,
			},
		},
		Resource: ResourceCustomerManagedKey(),
		HCL: `
			account_id = "abc"

			aws_key_info {
				key_arn   = "key-arn"
				key_alias = "key-alias"
			}
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
		Resource: ResourceCustomerManagedKey(),
		HCL: `
			account_id = "abc"

			aws_key_info {
				key_arn   = "key-arn"
				key_alias = "key-alias"
			}
		`,
		ID:     "abc/cmkid",
		Delete: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/cmkid", d.Id())
}
