package service

import (
	"github.com/databrickslabs/databricks-terraform/client/model"
)

// ClusterPoliciesAPI  allows you to create, list, and edit cluster policies.
//
// Creation and editing is available to admins only.
// Listing can be performed by any user and is limited to policies accessible by that user.
type ClusterPoliciesAPI struct {
	client *DatabricksClient
}

type policyIDWrapper struct {
	PolicyID string `json:"policy_id,omitempty" url:"policy_id,omitempty"`
}

// Create creates new cluster policy and sets PolicyID
func (a ClusterPoliciesAPI) Create(clusterPolicy *model.ClusterPolicy) error {
	var policyIDResponse = policyIDWrapper{}
	err := a.client.post("/policies/clusters/create", clusterPolicy, &policyIDResponse)
	if err != nil {
		return err
	}
	clusterPolicy.PolicyID = policyIDResponse.PolicyID
	return nil
}

// Edit will update an existing policy.
// This may make some clusters governed by this policy invalid.
// For such clusters the next cluster edit must provide a confirming configuration,
// but otherwise they can continue to run.
func (a ClusterPoliciesAPI) Edit(clusterPolicy *model.ClusterPolicy) error {
	return a.client.post("/policies/clusters/edit", clusterPolicy, nil)
}

// Get returns cluster policy
func (a ClusterPoliciesAPI) Get(policyID string) (policy *model.ClusterPolicy, err error) {
	err = a.client.get("/policies/clusters/get", policyIDWrapper{policyID}, policy)
	return
}

// Delete removes cluster policy
func (a ClusterPoliciesAPI) Delete(policyID string) error {
	return a.client.post("/policies/clusters/delete", policyIDWrapper{policyID}, nil)
}
