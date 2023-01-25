package policies

import (
	"context"

	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ClusterPolicy defines cluster policy
type ClusterPolicy struct {
	PolicyID           string `json:"policy_id,omitempty"`
	Name               string `json:"name"`
	Definition         string `json:"definition"`
	CreatedAtTimeStamp int64  `json:"created_at_timestamp"`
	MaxClustersPerUser int64  `json:"max_clusters_per_user,omitempty"`
}

// ClusterPolicyCreate is the endity used for request
type ClusterPolicyCreate struct {
	Name       string `json:"name"`
	Definition string `json:"definition"`
}

type ClusterPolicyList struct {
	Policies []ClusterPolicy `json:"policies"`
}

// NewClusterPoliciesAPI creates ClusterPoliciesAPI instance from provider meta
// Creation and editing is available to admins only.
func NewClusterPoliciesAPI(ctx context.Context, m any) ClusterPoliciesAPI {
	return ClusterPoliciesAPI{m.(*common.DatabricksClient), ctx}
}

// ClusterPoliciesAPI struct for cluster policies API
type ClusterPoliciesAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

type policyIDWrapper struct {
	PolicyID string `json:"policy_id,omitempty" url:"policy_id,omitempty"`
}

// Create creates new cluster policy and sets PolicyID
func (a ClusterPoliciesAPI) Create(clusterPolicy *ClusterPolicy) error {
	var policyIDResponse = policyIDWrapper{}
	err := a.client.Post(a.context, "/policies/clusters/create", clusterPolicy, &policyIDResponse)
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
func (a ClusterPoliciesAPI) Edit(clusterPolicy *ClusterPolicy) error {
	return a.client.Post(a.context, "/policies/clusters/edit", clusterPolicy, nil)
}

// Get returns cluster policy
func (a ClusterPoliciesAPI) Get(policyID string) (policy ClusterPolicy, err error) {
	err = a.client.Get(a.context, "/policies/clusters/get", policyIDWrapper{policyID}, &policy)
	return
}

// Delete removes cluster policy
func (a ClusterPoliciesAPI) Delete(policyID string) error {
	return a.client.Post(a.context, "/policies/clusters/delete", policyIDWrapper{policyID}, nil)
}

// Get returns cluster policy
func (a ClusterPoliciesAPI) List() ([]ClusterPolicy, error) {
	var lst ClusterPolicyList
	err := a.client.Get(a.context, "/policies/clusters/list", nil, &lst)
	if err != nil {
		return []ClusterPolicy{}, err
	}
	return lst.Policies, nil
}

func parsePolicyFromData(d *schema.ResourceData) (*ClusterPolicy, error) {
	clusterPolicy := new(ClusterPolicy)
	clusterPolicy.PolicyID = d.Id()
	if name, ok := d.GetOk("name"); ok {
		clusterPolicy.Name = name.(string)
	}
	if data, ok := d.GetOk("definition"); ok {
		clusterPolicy.Definition = data.(string)
	}
	return clusterPolicy, nil
}

// ResourceClusterPolicy ...
func ResourceClusterPolicy() *schema.Resource {
	return common.Resource{
		Schema: map[string]*schema.Schema{
			"policy_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Description: "Cluster policy name. This must be unique.\n" +
					"Length must be between 1 and 100 characters.",
				ValidateFunc: validation.StringLenBetween(1, 100),
			},
			"definition": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Policy definition JSON document expressed in\n" +
					"Databricks Policy Definition Language.",
				ValidateFunc: validation.StringIsJSON,
			},
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			clusterPolicy, err := parsePolicyFromData(d)
			if err != nil {
				return err
			}
			if err = NewClusterPoliciesAPI(ctx, c).Create(clusterPolicy); err != nil {
				return err
			}
			d.SetId(clusterPolicy.PolicyID)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			clusterPolicy, err := NewClusterPoliciesAPI(ctx, c).Get(d.Id())
			if err != nil {
				return err
			}
			if err = d.Set("name", clusterPolicy.Name); err != nil {
				return err
			}
			if err = d.Set("definition", clusterPolicy.Definition); err != nil {
				return err
			}
			if err = d.Set("policy_id", clusterPolicy.PolicyID); err != nil {
				return err
			}
			return nil
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			clusterPolicy, err := parsePolicyFromData(d)
			if err != nil {
				return err
			}
			return NewClusterPoliciesAPI(ctx, c).Edit(clusterPolicy)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewClusterPoliciesAPI(ctx, c).Delete(d.Id())
		},
	}.ToResource()
}
