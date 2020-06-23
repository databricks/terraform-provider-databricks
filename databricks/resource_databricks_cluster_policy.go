package databricks

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
)

func parsePolicyFromData(d *schema.ResourceData) (*model.ClusterPolicy, error) {
	clusterPolicy := new(model.ClusterPolicy)
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
	client := m.(*service.DBApiClient)
	clusterPolicy, err := parsePolicyFromData(d)
	if err != nil {
		return err
	}
	err = client.ClusterPolicies().Create(clusterPolicy)
	if err != nil {
		return err
	}
	d.SetId(clusterPolicy.PolicyID)
	return resourceClusterPolicyRead(d, m)
}

func resourceClusterPolicyRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)
	clusterPolicy, err := client.ClusterPolicies().Get(d.Id())
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
	return nil
}

func resourceClusterPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)
	clusterPolicy, err := parsePolicyFromData(d)
	if err != nil {
		return err
	}
	err = client.ClusterPolicies().Edit(clusterPolicy)
	if err != nil {
		return err
	}
	return resourceClusterPolicyRead(d, m)
}

func resourceClusterPolicyDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DBApiClient)
	return client.ClusterPolicies().Delete(id)
}

func resourceClusterPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceClusterPolicyCreate,
		Read:   resourceClusterPolicyRead,
		Update: resourceClusterPolicyUpdate,
		Delete: resourceClusterPolicyDelete,

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
