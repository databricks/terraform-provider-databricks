package compute

import (
	"log"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ClusterPoliciesAPI  allows you to create, list, and edit cluster policies.
//
// Creation and editing is available to admins only.
// NewClusterPoliciesAPI creates ClusterPoliciesAPI instance from provider meta
func NewClusterPoliciesAPI(m interface{}) ClusterPoliciesAPI {
	return ClusterPoliciesAPI{client: m.(*common.DatabricksClient)}
}

// Listing can be performed by any user and is limited to policies accessible by that user.
type ClusterPoliciesAPI struct {
	client *common.DatabricksClient
}

type policyIDWrapper struct {
	PolicyID string `json:"policy_id,omitempty" url:"policy_id,omitempty"`
}

// Create creates new cluster policy and sets PolicyID
func (a ClusterPoliciesAPI) Create(clusterPolicy *ClusterPolicy) error {
	var policyIDResponse = policyIDWrapper{}
	err := a.client.Post("/policies/clusters/create", clusterPolicy, &policyIDResponse)
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
	return a.client.Post("/policies/clusters/edit", clusterPolicy, nil)
}

// Get returns cluster policy
func (a ClusterPoliciesAPI) Get(policyID string) (policy ClusterPolicy, err error) {
	err = a.client.Get("/policies/clusters/get", policyIDWrapper{policyID}, &policy)
	return
}

// Delete removes cluster policy
func (a ClusterPoliciesAPI) Delete(policyID string) error {
	return a.client.Post("/policies/clusters/delete", policyIDWrapper{policyID}, nil)
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

func resourceClusterPolicyCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*common.DatabricksClient)
	clusterPolicy, err := parsePolicyFromData(d)
	if err != nil {
		return err
	}
	err = NewClusterPoliciesAPI(client).Create(clusterPolicy)
	if err != nil {
		return err
	}
	d.SetId(clusterPolicy.PolicyID)
	return resourceClusterPolicyRead(d, m)
}

func resourceClusterPolicyRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*common.DatabricksClient)
	clusterPolicy, err := NewClusterPoliciesAPI(client).Get(d.Id())
	if e, ok := err.(common.APIError); ok && e.IsMissing() {
		log.Printf("[ERROR] missing resource due to error: %v\n", e)
		d.SetId("")
		return nil
	}
	if err != nil {
		return err
	}
	err = d.Set("name", clusterPolicy.Name)
	if err != nil {
		return err
	}
	err = d.Set("definition", clusterPolicy.Definition)
	if err != nil {
		return err
	}
	err = d.Set("policy_id", clusterPolicy.PolicyID)
	if err != nil {
		return err
	}
	return nil
}

func resourceClusterPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*common.DatabricksClient)
	clusterPolicy, err := parsePolicyFromData(d)
	if err != nil {
		return err
	}
	err = NewClusterPoliciesAPI(client).Edit(clusterPolicy)
	if err != nil {
		return err
	}
	return resourceClusterPolicyRead(d, m)
}

func resourceClusterPolicyDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*common.DatabricksClient)
	return NewClusterPoliciesAPI(client).Delete(id)
}

func ResourceClusterPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceClusterPolicyCreate,
		Read:   resourceClusterPolicyRead,
		Update: resourceClusterPolicyUpdate,
		Delete: resourceClusterPolicyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
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
	}
}
