package policy

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/policy"
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/structure"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/policy/parse"
	azSchema "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/schema"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/timeouts"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func resourceArmPolicyDefinition() *schema.Resource {
	return &schema.Resource{
		Create: resourceArmPolicyDefinitionCreateUpdate,
		Update: resourceArmPolicyDefinitionCreateUpdate,
		Read:   resourceArmPolicyDefinitionRead,
		Delete: resourceArmPolicyDefinitionDelete,

		Importer: azSchema.ValidateResourceIDPriorToImport(func(id string) error {
			_, err := parse.PolicyDefinitionID(id)
			return err
		}),

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"policy_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(policy.TypeBuiltIn),
					string(policy.TypeCustom),
					string(policy.TypeNotSpecified),
				}, true)},

			"mode": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(policy.All),
					string(policy.Indexed),
					string(policy.NotSpecified),
				}, true),
			},

			"management_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"policy_rule": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateFunc:     validation.StringIsJSON,
				DiffSuppressFunc: structure.SuppressJsonDiff,
			},

			"parameters": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateFunc:     validation.StringIsJSON,
				DiffSuppressFunc: structure.SuppressJsonDiff,
			},

			"metadata": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     validation.StringIsJSON,
				DiffSuppressFunc: structure.SuppressJsonDiff,
			},
		},
	}
}

func resourceArmPolicyDefinitionCreateUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Policy.DefinitionsClient
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	name := d.Get("name").(string)
	policyType := d.Get("policy_type").(string)
	mode := d.Get("mode").(string)
	displayName := d.Get("display_name").(string)
	description := d.Get("description").(string)
	managementGroupID := d.Get("management_group_id").(string)

	if d.IsNewResource() {
		existing, err := getPolicyDefinition(ctx, client, name, managementGroupID)
		if err != nil {
			if !utils.ResponseWasNotFound(existing.Response) {
				return fmt.Errorf("Error checking for presence of existing Policy Definition %q: %s", name, err)
			}
		}

		if existing.ID != nil && *existing.ID != "" {
			return tf.ImportAsExistsError("azurerm_policy_definition", *existing.ID)
		}
	}

	properties := policy.DefinitionProperties{
		PolicyType:  policy.Type(policyType),
		Mode:        policy.Mode(mode),
		DisplayName: utils.String(displayName),
		Description: utils.String(description),
	}

	if policyRuleString := d.Get("policy_rule").(string); policyRuleString != "" {
		policyRule, err := structure.ExpandJsonFromString(policyRuleString)
		if err != nil {
			return fmt.Errorf("unable to parse policy_rule: %s", err)
		}
		properties.PolicyRule = &policyRule
	}

	if metaDataString := d.Get("metadata").(string); metaDataString != "" {
		metaData, err := structure.ExpandJsonFromString(metaDataString)
		if err != nil {
			return fmt.Errorf("unable to parse metadata: %s", err)
		}
		properties.Metadata = &metaData
	}

	if parametersString := d.Get("parameters").(string); parametersString != "" {
		parameters, err := structure.ExpandJsonFromString(parametersString)
		if err != nil {
			return fmt.Errorf("unable to parse parameters: %s", err)
		}
		properties.Parameters = &parameters
	}

	definition := policy.Definition{
		Name:                 utils.String(name),
		DefinitionProperties: &properties,
	}

	var err error

	if managementGroupID == "" {
		_, err = client.CreateOrUpdate(ctx, name, definition)
	} else {
		_, err = client.CreateOrUpdateAtManagementGroup(ctx, name, definition, managementGroupID)
	}

	if err != nil {
		return err
	}

	// Policy Definitions are eventually consistent; wait for them to stabilize
	log.Printf("[DEBUG] Waiting for Policy Definition %q to become available", name)
	stateConf := &resource.StateChangeConf{
		Pending:                   []string{"404"},
		Target:                    []string{"200"},
		Refresh:                   policyDefinitionRefreshFunc(ctx, client, name, managementGroupID),
		MinTimeout:                10 * time.Second,
		ContinuousTargetOccurence: 10,
	}

	if d.IsNewResource() {
		stateConf.Timeout = d.Timeout(schema.TimeoutCreate)
	} else {
		stateConf.Timeout = d.Timeout(schema.TimeoutUpdate)
	}

	if _, err = stateConf.WaitForState(); err != nil {
		return fmt.Errorf("Error waiting for Policy Definition %q to become available: %s", name, err)
	}

	resp, err := getPolicyDefinition(ctx, client, name, managementGroupID)
	if err != nil {
		return err
	}

	d.SetId(*resp.ID)

	return resourceArmPolicyDefinitionRead(d, meta)
}

func resourceArmPolicyDefinitionRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Policy.DefinitionsClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.PolicyDefinitionID(d.Id())
	if err != nil {
		return err
	}

	managementGroupID := ""
	switch scopeId := id.PolicyScopeId.(type) {
	case parse.ScopeAtManagementGroup:
		managementGroupID = scopeId.ManagementGroupId
	}

	resp, err := getPolicyDefinition(ctx, client, id.Name, managementGroupID)

	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			log.Printf("[INFO] Error reading Policy Definition %q - removing from state", d.Id())
			d.SetId("")
			return nil
		}

		return fmt.Errorf("Error reading Policy Definition %+v", err)
	}

	d.Set("name", resp.Name)
	d.Set("management_group_id", managementGroupID)

	if props := resp.DefinitionProperties; props != nil {
		d.Set("policy_type", props.PolicyType)
		d.Set("mode", props.Mode)
		d.Set("display_name", props.DisplayName)
		d.Set("description", props.Description)

		if policyRuleStr := flattenJSON(props.PolicyRule); policyRuleStr != "" {
			d.Set("policy_rule", policyRuleStr)
		}

		if metadataStr := flattenJSON(props.Metadata); metadataStr != "" {
			d.Set("metadata", metadataStr)
		}

		if parametersStr := flattenJSON(props.Parameters); parametersStr != "" {
			d.Set("parameters", parametersStr)
		}
	}

	return nil
}

func resourceArmPolicyDefinitionDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Policy.DefinitionsClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.PolicyDefinitionID(d.Id())
	if err != nil {
		return err
	}

	managementGroupID := ""
	switch scopeId := id.PolicyScopeId.(type) {
	case parse.ScopeAtManagementGroup:
		managementGroupID = scopeId.ManagementGroupId
	}

	var resp autorest.Response
	if managementGroupID == "" {
		resp, err = client.Delete(ctx, id.Name)
	} else {
		resp, err = client.DeleteAtManagementGroup(ctx, id.Name, managementGroupID)
	}

	if err != nil {
		if utils.ResponseWasNotFound(resp) {
			return nil
		}

		return fmt.Errorf("Error deleting Policy Definition %q: %+v", id.Name, err)
	}

	return nil
}

func policyDefinitionRefreshFunc(ctx context.Context, client *policy.DefinitionsClient, name string, managementGroupID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		res, err := getPolicyDefinition(ctx, client, name, managementGroupID)

		if err != nil {
			return nil, strconv.Itoa(res.StatusCode), fmt.Errorf("Error issuing read request in policyAssignmentRefreshFunc for Policy Assignment %q: %s", name, err)
		}

		return res, strconv.Itoa(res.StatusCode), nil
	}
}

func getPolicyDefinition(ctx context.Context, client *policy.DefinitionsClient, name string, managementGroupID string) (res policy.Definition, err error) {
	if managementGroupID == "" {
		res, err = client.Get(ctx, name)
	} else {
		res, err = client.GetAtManagementGroup(ctx, name, managementGroupID)
	}

	return res, err
}

func flattenJSON(stringMap interface{}) string {
	if stringMap != nil {
		value := stringMap.(map[string]interface{})
		jsonString, err := structure.FlattenJsonToString(value)
		if err == nil {
			return jsonString
		}
	}

	return ""
}
