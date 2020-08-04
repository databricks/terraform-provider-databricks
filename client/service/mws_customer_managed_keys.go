package service

import (
	"errors"
	"fmt"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

// MWSCustomerManagedKeysAPI exposes the mws customerManagedKeys API
type MWSCustomerManagedKeysAPI struct {
	client *DatabricksClient
}

// Create creates a set of MWS CustomerManagedKeys for the BYOVPC
func (a MWSCustomerManagedKeysAPI) Create(mwsAcctID, keyArn, keyAlias, keyRegion string) (k model.MWSCustomerManagedKey, err error) {
	customerManagedKeysAPIPath := fmt.Sprintf("/accounts/%s/customer-managed-keys", mwsAcctID)
	err = a.client.post(customerManagedKeysAPIPath, model.MWSCustomerManagedKey{
		AwsKeyInfo: &model.AwsKeyInfo{
			KeyArn:    keyArn,
			KeyAlias:  keyAlias,
			KeyRegion: keyRegion,
		},
	}, &k)
	return
}

// Read returns the customer managed key object along with metadata
func (a MWSCustomerManagedKeysAPI) Read(
	mwsAcctID, customerManagedKeysID string) (k model.MWSCustomerManagedKey, err error) {
	err = a.client.get(fmt.Sprintf("/accounts/%s/customer-managed-keys/%s",
		mwsAcctID, customerManagedKeysID), nil, &k)
	return
}

// Delete deletes the customer managed key object given a network id
func (a MWSCustomerManagedKeysAPI) Delete(customerManagedKeysID string) error {
	return errors.New("delete is not yet supported")
}

// List lists all the available customer managed key objects in the mws account
func (a MWSCustomerManagedKeysAPI) List(mwsAcctID string) (kl []model.MWSCustomerManagedKey, err error) {
	err = a.client.get(fmt.Sprintf("/accounts/%s/customer-managed-keys", mwsAcctID), nil, &kl)
	return
}
