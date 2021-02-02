package mws

import (
	"context"
	"fmt"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/databrickslabs/databricks-terraform/internal/util"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// AwsKeyInfo has information about the KMS key for BYOK
type AwsKeyInfo struct {
	KeyArn    string `json:"key_arn"`
	KeyAlias  string `json:"key_alias"`
	KeyRegion string `json:"key_region,omitempty" tf:"computed"`
}

// CustomerManagedKey contains key information and metadata for BYOK for E2
type CustomerManagedKey struct {
	CustomerManagedKeyID string      `json:"customer_managed_key_id,omitempty" tf:"computed"`
	AwsKeyInfo           *AwsKeyInfo `json:"aws_key_info"`
	AccountID            string      `json:"account_id"`
	CreationTime         int64       `json:"creation_time,omitempty" tf:"computed"`
}

// NewCustomerManagedKeysAPI creates CustomerManagedKeysAPI instance from provider meta
func NewCustomerManagedKeysAPI(ctx context.Context, m interface{}) CustomerManagedKeysAPI {
	return CustomerManagedKeysAPI{m.(*common.DatabricksClient), ctx}
}

// CustomerManagedKeysAPI exposes the mws customerManagedKeys API
type CustomerManagedKeysAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Create creates a set of MWS CustomerManagedKeys for the BYOVPC
func (a CustomerManagedKeysAPI) Create(cmk CustomerManagedKey) (k CustomerManagedKey, err error) {
	customerManagedKeysAPIPath := fmt.Sprintf("/accounts/%s/customer-managed-keys", cmk.AccountID)
	err = a.client.Post(a.context, customerManagedKeysAPIPath, cmk, &k)
	return
}

// Read returns the customer managed key object along with metadata
func (a CustomerManagedKeysAPI) Read(
	accountID, customerManagedKeyID string) (k CustomerManagedKey, err error) {
	err = a.client.Get(a.context, fmt.Sprintf("/accounts/%s/customer-managed-keys/%s",
		accountID, customerManagedKeyID), nil, &k)
	return
}

// Delete deletes the customer managed key object given a network id
func (a CustomerManagedKeysAPI) Delete(accountID, customerManagedKeyID string) error {
	return a.client.Delete(a.context, fmt.Sprintf("/accounts/%s/customer-managed-keys/%s",
		accountID, customerManagedKeyID), nil)
}

// List lists all the available customer managed key objects in the mws account
func (a CustomerManagedKeysAPI) List(accountID string) (kl []CustomerManagedKey, err error) {
	err = a.client.Get(a.context, fmt.Sprintf("/accounts/%s/customer-managed-keys", accountID), nil, &kl)
	return
}

// ResourceCustomerManagedKey ...
func ResourceCustomerManagedKey() *schema.Resource {
	s := internal.StructToSchema(CustomerManagedKey{},
		func(s map[string]*schema.Schema) map[string]*schema.Schema {
			s["aws_key_info"].ForceNew = true
			s["account_id"].ForceNew = true
			return s
		})
	p := util.NewPairSeparatedID("account_id", "customer_managed_key_id", "/")
	return util.CommonResource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var cmk CustomerManagedKey
			if err := internal.DataToStructPointer(d, s, &cmk); err != nil {
				return err
			}
			customerManagedKeyData, err := NewCustomerManagedKeysAPI(ctx, c).Create(cmk)
			if err != nil {
				return err
			}
			d.Set("customer_managed_key_id", customerManagedKeyData.CustomerManagedKeyID)
			p.Pack(d)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			accountID, cmkID, err := p.Unpack(d)
			if err != nil {
				return err
			}
			cmk, err := NewCustomerManagedKeysAPI(ctx, c).Read(accountID, cmkID)
			if err != nil {
				return err
			}
			return internal.StructToData(cmk, s, d)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			accountID, cmkID, err := p.Unpack(d)
			if err != nil {
				return err
			}
			return NewCustomerManagedKeysAPI(ctx, c).Delete(accountID, cmkID)
		},
		Schema: s,
	}.ToResource()
}
