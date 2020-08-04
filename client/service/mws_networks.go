package service

import (
	"fmt"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

// MWSNetworksAPI exposes the mws networks API
type MWSNetworksAPI struct {
	client *DatabricksClient
}

// Create creates a set of MWS Networks for the BYOVPC
func (a MWSNetworksAPI) Create(mwsAcctID, networkName string, vpcID string, subnetIds []string, securityGroupIds []string) (model.MWSNetwork, error) {
	var mwsNetwork model.MWSNetwork
	networksAPIPath := fmt.Sprintf("/accounts/%s/networks", mwsAcctID)
	err := a.client.post(networksAPIPath, model.MWSNetwork{
		NetworkName:      networkName,
		VPCID:            vpcID,
		SubnetIds:        subnetIds,
		SecurityGroupIds: securityGroupIds,
	}, &mwsNetwork)
	return mwsNetwork, err
}

// Read returns the network object along with metadata and any additional errors when attaching to workspace
func (a MWSNetworksAPI) Read(mwsAcctID, networksID string) (model.MWSNetwork, error) {
	var mwsNetwork model.MWSNetwork
	networksAPIPath := fmt.Sprintf("/accounts/%s/networks/%s", mwsAcctID, networksID)
	err := a.client.get(networksAPIPath, nil, &mwsNetwork)
	return mwsNetwork, err
}

// Delete deletes the network object given a network id
func (a MWSNetworksAPI) Delete(mwsAcctID, networksID string) error {
	networksAPIPath := fmt.Sprintf("/accounts/%s/networks/%s", mwsAcctID, networksID)
	return a.client.delete(networksAPIPath, nil)
}

// List lists all the available network objects in the mws account
func (a MWSNetworksAPI) List(mwsAcctID string) ([]model.MWSNetwork, error) {
	var mwsNetworkList []model.MWSNetwork
	networksAPIPath := fmt.Sprintf("/accounts/%s/networks", mwsAcctID)
	err := a.client.get(networksAPIPath, nil, &mwsNetworkList)
	return mwsNetworkList, err
}
