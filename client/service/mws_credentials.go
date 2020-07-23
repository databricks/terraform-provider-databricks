package service

import (
	"fmt"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

// MWSCredentialsAPI exposes the mws credentials API
type MWSCredentialsAPI struct {
	client *DatabricksClient
}

// TODO: move mwsAcctID into provider configuration...

// Create creates a set of MWS Credentials for the cross account role
func (a MWSCredentialsAPI) Create(mwsAcctID, credentialsName string, roleArn string) (model.MWSCredentials, error) {
	var mwsCreds model.MWSCredentials
	credentialsAPIPath := fmt.Sprintf("/accounts/%s/credentials", mwsAcctID)
	err := a.client.post(credentialsAPIPath, model.MWSCredentials{
		CredentialsName: credentialsName,
		AwsCredentials: &model.AwsCredentials{
			StsRole: &model.StsRole{
				RoleArn: roleArn,
			},
		},
	}, &mwsCreds)
	return mwsCreds, err
}

// Read returns the credentials object along with metadata
func (a MWSCredentialsAPI) Read(mwsAcctID, credentialsID string) (model.MWSCredentials, error) {
	var mwsCreds model.MWSCredentials
	credentialsAPIPath := fmt.Sprintf("/accounts/%s/credentials/%s", mwsAcctID, credentialsID)
	err := a.client.get(credentialsAPIPath, nil, &mwsCreds)
	return mwsCreds, err
}

// Delete deletes the credentials object given a credentials id
func (a MWSCredentialsAPI) Delete(mwsAcctID, credentialsID string) error {
	credentialsAPIPath := fmt.Sprintf("/accounts/%s/credentials/%s", mwsAcctID, credentialsID)
	return a.client.delete(credentialsAPIPath, nil)
}

// List lists all the available credentials object in the mws account
func (a MWSCredentialsAPI) List(mwsAcctID string) ([]model.MWSCredentials, error) {
	var mwsCredsList []model.MWSCredentials
	credentialsAPIPath := fmt.Sprintf("/accounts/%s/credentials", mwsAcctID)
	err := a.client.get(credentialsAPIPath, nil, &mwsCredsList)
	return mwsCredsList, err
}
