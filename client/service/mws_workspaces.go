package service

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

// MWSWorkspacesAPI exposes the mws workspaces API
type MWSWorkspacesAPI struct {
	client *DatabricksClient
}

// Create creates the workspace creation process
func (a MWSWorkspacesAPI) Create(mwsAcctID, workspaceName, deploymentName, awsRegion, credentialsID, storageConfigurationID, networkID, customerManagedKeyID string, isNoPublicIPEnabled bool) (model.MWSWorkspace, error) {
	var mwsWorkspace model.MWSWorkspace
	workspacesAPIPath := fmt.Sprintf("/accounts/%s/workspaces", mwsAcctID)
	mwsWorkspacesRequest := model.MWSWorkspace{
		WorkspaceName:          workspaceName,
		DeploymentName:         deploymentName,
		AwsRegion:              awsRegion,
		CredentialsID:          credentialsID,
		StorageConfigurationID: storageConfigurationID,
		IsNoPublicIPEnabled:    isNoPublicIPEnabled,
	}
	if !reflect.ValueOf(networkID).IsZero() {
		mwsWorkspacesRequest.NetworkID = networkID
	}
	if !reflect.ValueOf(customerManagedKeyID).IsZero() {
		mwsWorkspacesRequest.CustomerManagedKeyID = customerManagedKeyID
	}
	err := a.client.post(workspacesAPIPath, mwsWorkspacesRequest, &mwsWorkspace)
	return mwsWorkspace, err
}

// WaitForWorkspaceRunning will hold the main thread till the workspace is in a running state
func (a MWSWorkspacesAPI) WaitForWorkspaceRunning(mwsAcctID string, workspaceID int64, sleepDurationSeconds time.Duration, timeoutDurationMinutes time.Duration) error {
	// TODO: move all resource awaiters from client to TF resource level, for sepration of concerns sake
	errChan := make(chan error, 1)
	go func() {
		for {
			workspace, err := a.Read(mwsAcctID, workspaceID)
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
func (a MWSWorkspacesAPI) Patch(mwsAcctID string, workspaceID int64, awsRegion, credentialsID, storageConfigurationID, networkID, customerManagedKeyID string, isNoPublicIPEnabled bool) error {
	workspacesAPIPath := fmt.Sprintf("/accounts/%s/workspaces/%d", mwsAcctID, workspaceID)
	mwsWorkspacesRequest := model.MWSWorkspace{
		AwsRegion:              awsRegion,
		CredentialsID:          credentialsID,
		StorageConfigurationID: storageConfigurationID,
		IsNoPublicIPEnabled:    isNoPublicIPEnabled,
	}
	if !reflect.ValueOf(networkID).IsZero() {
		mwsWorkspacesRequest.NetworkID = networkID
	}
	if !reflect.ValueOf(customerManagedKeyID).IsZero() {
		mwsWorkspacesRequest.CustomerManagedKeyID = customerManagedKeyID
	}
	return a.client.patch(workspacesAPIPath, mwsWorkspacesRequest)
}

// Read will return the mws workspace metadata and status of the workspace deployment
func (a MWSWorkspacesAPI) Read(mwsAcctID string, workspaceID int64) (model.MWSWorkspace, error) {
	var mwsWorkspace model.MWSWorkspace
	workspacesAPIPath := fmt.Sprintf("/accounts/%s/workspaces/%d", mwsAcctID, workspaceID)
	err := a.client.get(workspacesAPIPath, nil, &mwsWorkspace)
	return mwsWorkspace, err
}

// Delete will delete the configuration for the workspace given a workspace id and will not block. A follow up email
// will be sent when the workspace is fully deleted.
func (a MWSWorkspacesAPI) Delete(mwsAcctID string, workspaceID int64) error {
	workspacesAPIPath := fmt.Sprintf("/accounts/%s/workspaces/%d", mwsAcctID, workspaceID)
	return a.client.delete(workspacesAPIPath, nil)
}

// List will list all workspaces in a given mws account
func (a MWSWorkspacesAPI) List(mwsAcctID string) ([]model.MWSWorkspace, error) {
	var mwsWorkspacesList []model.MWSWorkspace
	workspacesAPIPath := fmt.Sprintf("/accounts/%s/workspaces", mwsAcctID)
	err := a.client.get(workspacesAPIPath, nil, &mwsWorkspacesList)
	return mwsWorkspacesList, err
}
