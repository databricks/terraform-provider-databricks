package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

// MWSCredentialsAPI exposes the mws credentials API
type MWSCredentialsAPI struct {
	Client *DBApiClient
}

// Create creates a set of MWS Credentials for the cross account role
func (a MWSCredentialsAPI) Create(mwsAcctId, credentialsName string, roleArn string) (model.MWSCredentials, error) {
	var mwsCreds model.MWSCredentials

	credentialsAPIPath := fmt.Sprintf("/accounts/%s/credentials", mwsAcctId)

	mwsCredentialsRequest := model.MWSCredentials{
		CredentialsName: credentialsName,
		AwsCredentials: &model.AwsCredentials{
			StsRole: &model.StsRole{
				RoleArn: roleArn,
			},
		},
	}

	resp, err := a.Client.performQuery(http.MethodPost, credentialsAPIPath, "2.0", nil, mwsCredentialsRequest, nil)
	if err != nil {
		return mwsCreds, err
	}

	err = json.Unmarshal(resp, &mwsCreds)
	return mwsCreds, err
}

// Read returns the credentials object along with metadata
func (a MWSCredentialsAPI) Read(mwsAcctId, credentialsID string) (model.MWSCredentials, error) {
	var mwsCreds model.MWSCredentials

	credentialsAPIPath := fmt.Sprintf("/accounts/%s/credentials/%s", mwsAcctId, credentialsID)

	resp, err := a.Client.performQuery(http.MethodGet, credentialsAPIPath, "2.0", nil, nil, nil)
	if err != nil {
		return mwsCreds, err
	}

	err = json.Unmarshal(resp, &mwsCreds)
	return mwsCreds, err
}

// Delete deletes the credentials object given a credentials id
func (a MWSCredentialsAPI) Delete(mwsAcctId, credentialsID string) error {
	credentialsAPIPath := fmt.Sprintf("/accounts/%s/credentials/%s", mwsAcctId, credentialsID)

	_, err := a.Client.performQuery(http.MethodDelete, credentialsAPIPath, "2.0", nil, nil, nil)

	return err
}

// List lists all the available credentials object in the mws account
func (a MWSCredentialsAPI) List(mwsAcctId string) ([]model.MWSCredentials, error) {
	var mwsCredsList []model.MWSCredentials

	credentialsAPIPath := fmt.Sprintf("/accounts/%s/credentials", mwsAcctId)

	resp, err := a.Client.performQuery(http.MethodGet, credentialsAPIPath, "2.0", nil, nil, nil)
	if err != nil {
		return mwsCredsList, err
	}

	err = json.Unmarshal(resp, &mwsCredsList)
	return mwsCredsList, err
}
