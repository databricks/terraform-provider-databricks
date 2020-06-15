package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

// MWSCustomerManagedKeysAPI exposes the mws customerManagedKeys API
type MWSCustomerManagedKeysAPI struct {
	Client *DBApiClient
}

// Create creates a set of MWS CustomerManagedKeys for the BYOVPC
func (a MWSCustomerManagedKeysAPI) Create(mwsAcctId, keyArn, keyAlias, keyRegion string) (model.MWSCustomerManagedKey, error) {
	var mwsCustomerManagedKey model.MWSCustomerManagedKey

	customerManagedKeysAPIPath := fmt.Sprintf("/accounts/%s/customer-managed-keys", mwsAcctId)
	mwsCustomerManagedKeysRequest := model.MWSCustomerManagedKey{
		AwsKeyInfo: &model.AwsKeyInfo{
			KeyArn:    keyArn,
			KeyAlias:  keyAlias,
			KeyRegion: keyRegion,
		},
	}
	resp, err := a.Client.performQuery(http.MethodPost, customerManagedKeysAPIPath, "2.0", nil, mwsCustomerManagedKeysRequest, nil)
	if err != nil {
		return mwsCustomerManagedKey, err
	}
	err = json.Unmarshal(resp, &mwsCustomerManagedKey)
	return mwsCustomerManagedKey, err
}

// Read returns the customer managed key object along with metadata
func (a MWSCustomerManagedKeysAPI) Read(mwsAcctId, customerManagedKeysID string) (model.MWSCustomerManagedKey, error) {
	var mwsCustomerManagedKey model.MWSCustomerManagedKey
	customerManagedKeysAPIPath := fmt.Sprintf("/accounts/%s/customer-managed-keys/%s", mwsAcctId, customerManagedKeysID)
	resp, err := a.Client.performQuery(http.MethodGet, customerManagedKeysAPIPath, "2.0", nil, nil, nil)
	if err != nil {
		return mwsCustomerManagedKey, err
	}
	err = json.Unmarshal(resp, &mwsCustomerManagedKey)
	return mwsCustomerManagedKey, err
}

// Delete deletes the customer managed key object given a network id
func (a MWSCustomerManagedKeysAPI) Delete(customerManagedKeysID string) error {
	//customerManagedKeysAPIPath := fmt.Sprintf("/accounts/%s/customer-managed-keys/%s", a.Client.Config.E2AcctID, customerManagedKeysID)
	//_, err := a.Client.performQuery(http.MethodDelete, customerManagedKeysAPIPath, "2.0", nil, nil, nil)
	//return err
	return errors.New("delete is not yet supported")
}

// List lists all the available customer managed key objects in the mws account
func (a MWSCustomerManagedKeysAPI) List(mwsAcctId string) ([]model.MWSCustomerManagedKey, error) {
	var mwsCustomerManagedKeyList []model.MWSCustomerManagedKey

	customerManagedKeysAPIPath := fmt.Sprintf("/accounts/%s/customer-managed-keys", mwsAcctId)

	resp, err := a.Client.performQuery(http.MethodGet, customerManagedKeysAPIPath, "2.0", nil, nil, nil)
	if err != nil {
		return mwsCustomerManagedKeyList, err
	}

	err = json.Unmarshal(resp, &mwsCustomerManagedKeyList)
	return mwsCustomerManagedKeyList, err
}
