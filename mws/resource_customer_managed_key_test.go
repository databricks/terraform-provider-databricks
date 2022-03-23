package mws

import (
	"context"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/common"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
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
	cmkAPI := NewCustomerManagedKeysAPI(context.Background(), client)
	cmkList, err := cmkAPI.List(acctID)
	assert.NoError(t, err, err)
	t.Log(cmkList)

	keyInfo, err := cmkAPI.Create(CustomerManagedKey{
		AwsKeyInfo: &AwsKeyInfo{
			KeyArn:   kmsKeyArn,
			KeyAlias: kmsKeyAlias,
		},
		AccountID: acctID,
		UseCases:  []string{"MANAGED_SERVICES"},
	})
	assert.NoError(t, err, err)

	keyID := keyInfo.CustomerManagedKeyID

	defer func() {
		err := cmkAPI.Delete(acctID, keyID)
		assert.NoError(t, err, err)
	}()

	getKeyInfo, err := cmkAPI.Read(acctID, keyID)
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
					ReuseKeyForClusterVolumes: false,
					UseCases:                  []string{"MANAGED_SERVICES"},
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
					AccountID:                 "abc",
					ReuseKeyForClusterVolumes: false,
					UseCases:                  []string{"MANAGED_SERVICES"},
					CreationTime:              123,
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
			use_cases = ["MANAGED_SERVICES"]
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
					UseCases: []string{"MANAGED_SERVICE"},
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
					UseCases:     []string{"MANAGED_SERVICE"},
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
			use_cases = ["MANAGED_SERVICE"]
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
					ReuseKeyForClusterVolumes: false,
					UseCases:                  []string{"MANAGED_SERVICES"},
					AccountID:                 "abc",
					CreationTime:              123,
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
			use_cases = ["MANAGED_SERVICES"]
		`,
		ID:   "abc/cmkid",
		Read: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/cmkid", d.Id())
	assert.Equal(t, "key-arn", d.Get("aws_key_info.0.key_arn"))
	assert.Equal(t, "key-alias", d.Get("aws_key_info.0.key_alias"))
	assert.Equal(t, "us-east-1", d.Get("aws_key_info.0.key_region"))
	assert.Equal(t, []interface{}{"MANAGED_SERVICES"}, d.Get("use_cases"))
	assert.Equal(t, "abc", d.Get("account_id"))
	assert.Equal(t, 123, d.Get("creation_time"))
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
		Resource: ResourceCustomerManagedKey(),
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
	assert.NoError(t, err, err)
	assert.Equal(t, "abc/cmkid", d.Id())
}

func TestCmkStateUpgrader(t *testing.T) {
	state, err := migrateResourceCustomerManagedKeyV0(context.Background(),
		map[string]interface{}{}, nil)
	assert.NoError(t, err)
	_, ok := state["use_cases"]
	assert.True(t, ok)
}
