package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

// MWSWorkspacesAPI exposes the mws workspaces API
type MWSWorkspacesAPI struct {
	Client *DBApiClient
}

// Create creates the workspace creation process
func (a MWSWorkspacesAPI) Create(mwsAcctId, workspaceName, deploymentName, awsRegion, credentialsID, storageConfigurationID, networkID, customerManagedKeyID string, isNoPublicIpEnabled bool) (model.MWSWorkspace, error) {
	var mwsWorkspace model.MWSWorkspace

	workspacesAPIPath := fmt.Sprintf("/accounts/%s/workspaces", mwsAcctId)

	mwsWorkspacesRequest := model.MWSWorkspace{
		WorkspaceName:          workspaceName,
		DeploymentName:         deploymentName,
		AwsRegion:              awsRegion,
		CredentialsID:          credentialsID,
		StorageConfigurationID: storageConfigurationID,
		IsNoPublicIpEnabled:    isNoPublicIpEnabled,
	}

	if !reflect.ValueOf(networkID).IsZero() {
		mwsWorkspacesRequest.NetworkID = networkID
	}

	if !reflect.ValueOf(customerManagedKeyID).IsZero() {
		mwsWorkspacesRequest.CustomerManagedKeyID = customerManagedKeyID
	}

	resp, err := a.Client.performQuery(http.MethodPost, workspacesAPIPath, "2.0", nil, mwsWorkspacesRequest, nil)
	if err != nil {
		return mwsWorkspace, err
	}

	err = json.Unmarshal(resp, &mwsWorkspace)
	return mwsWorkspace, err
}

// WaitForWorkspaceRunning will hold the main thread till the workspace is in a running state
func (a MWSWorkspacesAPI) WaitForWorkspaceRunning(mwsAcctId string, workspaceID int64, sleepDurationSeconds time.Duration, timeoutDurationMinutes time.Duration) error {
	errChan := make(chan error, 1)
	go func() {
		for {
			workspace, err := a.Read(mwsAcctId, workspaceID)
			if err != nil {
				errChan <- err
			}
			if workspace.WorkspaceStatus == model.WorkspaceStatusRunning {
				errChan <- nil
			} else if model.ContainsWorkspaceState(model.WorkspaceStatusesNonRunnable, workspace.WorkspaceStatus) {
				errChan <- errors.New("Workspace is in a non runnable state will not be able to transition to running, needs " +
					"to be created again. Current state: " + workspace.WorkspaceStatus)
			}
			log.Println("Waiting for workspace to go to running, current state is: " + workspace.WorkspaceStatus)
			time.Sleep(sleepDurationSeconds * time.Second)
		}
	}()
	select {
	case err := <-errChan:
		return err
	case <-time.After(timeoutDurationMinutes * time.Minute):
		return errors.New("Timed out workspace has not reached running state")
	}
}

// Patch will relaunch the mws workspace deployment TODO: may need to include customer managed key
func (a MWSWorkspacesAPI) Patch(mwsAcctId string, workspaceID int64, awsRegion, credentialsID, storageConfigurationID, networkID, customerManagedKeyID string, isNoPublicIpEnabled bool) error {
	workspacesAPIPath := fmt.Sprintf("/accounts/%s/workspaces/%d", mwsAcctId, workspaceID)

	mwsWorkspacesRequest := model.MWSWorkspace{
		AwsRegion:              awsRegion,
		CredentialsID:          credentialsID,
		StorageConfigurationID: storageConfigurationID,
		IsNoPublicIpEnabled:    isNoPublicIpEnabled,
	}

	if !reflect.ValueOf(networkID).IsZero() {
		mwsWorkspacesRequest.NetworkID = networkID
	}

	if !reflect.ValueOf(customerManagedKeyID).IsZero() {
		mwsWorkspacesRequest.CustomerManagedKeyID = customerManagedKeyID
	}

	_, err := a.Client.performQuery(http.MethodPatch, workspacesAPIPath, "2.0", nil, mwsWorkspacesRequest, nil)
	return err
}

// Read will return the mws workspace metadata and status of the workspace deployment
func (a MWSWorkspacesAPI) Read(mwsAcctId string, workspaceID int64) (model.MWSWorkspace, error) {
	var mwsWorkspace model.MWSWorkspace

	workspacesAPIPath := fmt.Sprintf("/accounts/%s/workspaces/%d", mwsAcctId, workspaceID)

	resp, err := a.Client.performQuery(http.MethodGet, workspacesAPIPath, "2.0", nil, nil, nil)
	if err != nil {
		return mwsWorkspace, err
	}

	err = json.Unmarshal(resp, &mwsWorkspace)
	return mwsWorkspace, err
}

// Delete will delete the configuration for the workspace given a workspace id and will not block. A follow up email
// will be sent when the workspace is fully deleted.
func (a MWSWorkspacesAPI) Delete(mwsAcctId string, workspaceID int64) error {
	workspacesAPIPath := fmt.Sprintf("/accounts/%s/workspaces/%d", mwsAcctId, workspaceID)

	_, err := a.Client.performQuery(http.MethodDelete, workspacesAPIPath, "2.0", nil, nil, nil)

	return err
}

// List will list all workspaces in a given mws account
func (a MWSWorkspacesAPI) List(mwsAcctId string) ([]model.MWSWorkspace, error) {
	var mwsWorkspacesList []model.MWSWorkspace

	workspacesAPIPath := fmt.Sprintf("/accounts/%s/workspaces", mwsAcctId)

	resp, err := a.Client.performQuery(http.MethodGet, workspacesAPIPath, "2.0", nil, nil, nil)
	if err != nil {
		return mwsWorkspacesList, err
	}

	err = json.Unmarshal(resp, &mwsWorkspacesList)
	return mwsWorkspacesList, err
}
