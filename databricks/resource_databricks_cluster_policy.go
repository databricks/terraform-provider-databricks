package databricks

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

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
		err := clusterPolicy.ParseDefinition(data.(string))
		if err != nil {
			return nil, err
		}
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

	// err = clusterPolicy.ParseDefinition(clusterPolicy.Definition)
	// if err != nil {
	// 	return err
	// }
	// policies := clusterPolicy.AttributePoliciesState()
	// err = d.Set("attribute_policy", policies)
	// if err != nil {
	// 	return err
	// }
	// TODO: add definitions!!!

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
			},
			"definition": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Policy definition JSON document expressed in\n" +
					"Databricks Policy Definition Language.",
				ConflictsWith: []string{"attribute_policy"},
			},
			"attribute_policy": {
				Type:          schema.TypeList,
				Optional:      true,
				MinItems:      1,
				ConflictsWith: []string{"definition"},
				ConfigMode:    schema.SchemaConfigModeAttr,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"path": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": { // type: fixed
							Type:     schema.TypeString,
							Optional: true,
						},
						"default_value": { // type: limiting
							Type:     schema.TypeString,
							Optional: true,
						},
						"values": { // type: limiting
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"hidden": { // type: fixed
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"is_optional": { // type: limiting
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"pattern": { // type: regex
							Type:     schema.TypeString,
							Optional: true,
						},
						"min_value": { // type: range
							Type:     schema.TypeInt,
							Optional: true,
						},
						"max_value": { // type: range
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
		},
	}
}
