package service

import (
	"encoding/json"
	"net/http"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

// ClusterPoliciesAPI  allows you to create, list, and edit cluster policies.
//
// Creation and editing is available to admins only.
// Listing can be performed by any user and is limited to policies accessible by that user.
type ClusterPoliciesAPI struct {
	Client *DBApiClient
}

type policyIDWrapper struct {
	PolicyID string `json:"policy_id,omitempty" url:"policy_id,omitempty"`
}

// Create creates new cluster policy and sets PolicyID
func (a ClusterPoliciesAPI) Create(clusterPolicy *model.ClusterPolicy) error {
	//clusterPolicyWrapper := &policyWrapper{clusterPolicy.Name, clusterPolicy.Definition}
	resp, err := a.Client.performQuery(http.MethodPost, "/policies/clusters/create", "2.0", nil, clusterPolicy, nil)
	if err != nil {
		return err
	}
	var policyIDResponse = new(policyIDWrapper)
	err = json.Unmarshal(resp, &policyIDResponse)
	clusterPolicy.PolicyID = policyIDResponse.PolicyID
	return err
}

// Edit will update an existing policy.
// This may make some clusters governed by this policy invalid.
// For such clusters the next cluster edit must provide a confirming configuration,
// but otherwise they can continue to run.
func (a ClusterPoliciesAPI) Edit(clusterPolicy *model.ClusterPolicy) error {
	_, err := a.Client.performQuery(http.MethodPost, "/policies/clusters/edit", "2.0", nil, clusterPolicy, nil)
	return err
}

// Get returns cluster policy
func (a ClusterPoliciesAPI) Get(policyID string) (*model.ClusterPolicy, error) {
	var clusterPolicy model.ClusterPolicy
	resp, err := a.Client.performQuery(http.MethodGet, "/policies/clusters/get", "2.0", nil, policyIDWrapper{policyID}, nil)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &clusterPolicy)
	return &clusterPolicy, err
}

// Delete removes cluster policy
func (a ClusterPoliciesAPI) Delete(policyID string) error {
	_, err := a.Client.performQuery(http.MethodPost, "/policies/clusters/delete", "2.0", nil, policyIDWrapper{policyID}, nil)
	return err
}
