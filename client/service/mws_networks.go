package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

// MWSNetworksAPI exposes the mws networks API
type MWSNetworksAPI struct {
	Client *DBApiClient
}

// Create creates a set of MWS Networks for the BYOVPC
func (a MWSNetworksAPI) Create(mwsAcctId, networkName string, vpcID string, subnetIds []string, securityGroupIds []string) (model.MWSNetwork, error) {
	var mwsNetwork model.MWSNetwork
	networksAPIPath := fmt.Sprintf("/accounts/%s/networks", mwsAcctId)
	mwsNetworksRequest := model.MWSNetwork{
		NetworkName:      networkName,
		VPCID:            vpcID,
		SubnetIds:        subnetIds,
		SecurityGroupIds: securityGroupIds,
	}
	resp, err := a.Client.performQuery(http.MethodPost, networksAPIPath, "2.0", nil, mwsNetworksRequest, nil)
	if err != nil {
		return mwsNetwork, err
	}
	err = json.Unmarshal(resp, &mwsNetwork)
	return mwsNetwork, err
}

// Read returns the network object along with metadata and any additional errors when attaching to workspace
func (a MWSNetworksAPI) Read(mwsAcctId, networksID string) (model.MWSNetwork, error) {
	var mwsNetwork model.MWSNetwork
	networksAPIPath := fmt.Sprintf("/accounts/%s/networks/%s", mwsAcctId, networksID)
	resp, err := a.Client.performQuery(http.MethodGet, networksAPIPath, "2.0", nil, nil, nil)
	if err != nil {
		return mwsNetwork, err
	}
	err = json.Unmarshal(resp, &mwsNetwork)
	return mwsNetwork, err
}

// Delete deletes the network object given a network id
func (a MWSNetworksAPI) Delete(mwsAcctId, networksID string) error {
	networksAPIPath := fmt.Sprintf("/accounts/%s/networks/%s", mwsAcctId, networksID)
	_, err := a.Client.performQuery(http.MethodDelete, networksAPIPath, "2.0", nil, nil, nil)
	return err
}

// List lists all the available network objects in the mws account
func (a MWSNetworksAPI) List(mwsAcctId string) ([]model.MWSNetwork, error) {
	var mwsNetworkList []model.MWSNetwork

	networksAPIPath := fmt.Sprintf("/accounts/%s/networks", mwsAcctId)

	resp, err := a.Client.performQuery(http.MethodGet, networksAPIPath, "2.0", nil, nil, nil)
	if err != nil {
		return mwsNetworkList, err
	}

	err = json.Unmarshal(resp, &mwsNetworkList)
	return mwsNetworkList, err
}
