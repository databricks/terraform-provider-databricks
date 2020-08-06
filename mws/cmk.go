package mws

import (
	"fmt"

	"github.com/databrickslabs/databricks-terraform/common"
)

// NewMWSCustomerManagedKeysAPI creates MWSCustomerManagedKeysAPI instance from provider meta
func NewMWSCustomerManagedKeysAPI(m interface{}) MWSCustomerManagedKeysAPI {
	return MWSCustomerManagedKeysAPI{client: m.(*common.DatabricksClient)}
}

// MWSCustomerManagedKeysAPI exposes the mws customerManagedKeys API
type MWSCustomerManagedKeysAPI struct {
	client *common.DatabricksClient
}

// Create creates a set of MWS CustomerManagedKeys for the BYOVPC
func (a MWSCustomerManagedKeysAPI) Create(mwsAcctID, keyArn, keyAlias, keyRegion string) (k MWSCustomerManagedKey, err error) {
	customerManagedKeysAPIPath := fmt.Sprintf("/accounts/%s/customer-managed-keys", mwsAcctID)
	err = a.client.Post(customerManagedKeysAPIPath, MWSCustomerManagedKey{
		AwsKeyInfo: &AwsKeyInfo{
			KeyArn:    keyArn,
			KeyAlias:  keyAlias,
			KeyRegion: keyRegion,
		},
	}, &k)
	return
}

// Read returns the customer managed key object along with metadata
func (a MWSCustomerManagedKeysAPI) Read(
	mwsAcctID, customerManagedKeysID string) (k MWSCustomerManagedKey, err error) {
	err = a.client.Get(fmt.Sprintf("/accounts/%s/customer-managed-keys/%s",
		mwsAcctID, customerManagedKeysID), nil, &k)
	return
}

// Delete deletes the customer managed key object given a network id
func (a MWSCustomerManagedKeysAPI) Delete(customerManagedKeysID string) error {
	return fmt.Errorf("delete is not yet supported")
}

// List lists all the available customer managed key objects in the mws account
func (a MWSCustomerManagedKeysAPI) List(mwsAcctID string) (kl []MWSCustomerManagedKey, err error) {
	err = a.client.Get(fmt.Sprintf("/accounts/%s/customer-managed-keys", mwsAcctID), nil, &kl)
	return
}
