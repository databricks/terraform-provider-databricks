package keyvault

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/keyvault/validate"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/network"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2018-02-14/keyvault"
	"github.com/hashicorp/go-azure-helpers/response"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	uuid "github.com/satori/go.uuid"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/set"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/locks"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tags"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/timeouts"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

// As can be seen in the API definition, the Sku Family only supports the value
// `A` and is a required field
// https://github.com/Azure/azure-rest-api-specs/blob/master/arm-keyvault/2015-06-01/swagger/keyvault.json#L239
var armKeyVaultSkuFamily = "A"

var keyVaultResourceName = "azurerm_key_vault"

func resourceArmKeyVault() *schema.Resource {
	return &schema.Resource{
		Create: resourceArmKeyVaultCreate,
		Read:   resourceArmKeyVaultRead,
		Update: resourceArmKeyVaultUpdate,
		Delete: resourceArmKeyVaultDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		MigrateState:  resourceAzureRMKeyVaultMigrateState,
		SchemaVersion: 1,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.KeyVaultName,
			},

			"location": azure.SchemaLocation(),

			"resource_group_name": azure.SchemaResourceGroupName(),

			"sku_name": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(keyvault.Standard),
					string(keyvault.Premium),
				}, false),
			},

			"tenant_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsUUID,
			},

			"access_policy": {
				Type:       schema.TypeList,
				ConfigMode: schema.SchemaConfigModeAttr,
				Optional:   true,
				Computed:   true,
				MaxItems:   1024,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tenant_id": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.IsUUID,
						},
						"object_id": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.IsUUID,
						},
						"application_id": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.IsUUID,
						},
						"certificate_permissions": azure.SchemaKeyVaultCertificatePermissions(),
						"key_permissions":         azure.SchemaKeyVaultKeyPermissions(),
						"secret_permissions":      azure.SchemaKeyVaultSecretPermissions(),
						"storage_permissions":     azure.SchemaKeyVaultStoragePermissions(),
					},
				},
			},

			"enabled_for_deployment": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"enabled_for_disk_encryption": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"enabled_for_template_deployment": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"network_acls": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default_action": {
							Type:     schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringInSlice([]string{
								string(keyvault.Allow),
								string(keyvault.Deny),
							}, false),
						},
						"bypass": {
							Type:     schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringInSlice([]string{
								string(keyvault.None),
								string(keyvault.AzureServices),
							}, false),
						},
						"ip_rules": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
						"virtual_network_subnet_ids": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      set.HashStringIgnoreCase,
						},
					},
				},
			},

			"purge_protection_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"soft_delete_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"tags": tags.Schema(),

			// Computed
			"vault_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceArmKeyVaultCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).KeyVault.VaultsClient
	ctx, cancel := timeouts.ForCreate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	name := d.Get("name").(string)
	resourceGroup := d.Get("resource_group_name").(string)
	location := azure.NormalizeLocation(d.Get("location").(string))

	// Locking this resource so we don't make modifications to it at the same time if there is a
	// key vault access policy trying to update it as well
	locks.ByName(name, keyVaultResourceName)
	defer locks.UnlockByName(name, keyVaultResourceName)

	// check for the presence of an existing, live one which should be imported into the state
	existing, err := client.Get(ctx, resourceGroup, name)
	if err != nil {
		if !utils.ResponseWasNotFound(existing.Response) {
			return fmt.Errorf("Error checking for presence of existing Key Vault %q (Resource Group %q): %s", name, resourceGroup, err)
		}
	}

	if existing.ID != nil && *existing.ID != "" {
		return tf.ImportAsExistsError("azurerm_key_vault", *existing.ID)
	}

	// before creating check to see if the key vault exists in the soft delete state
	softDeletedKeyVault, err := client.GetDeleted(ctx, name, location)
	if err != nil {
		if !utils.ResponseWasNotFound(softDeletedKeyVault.Response) {
			return fmt.Errorf("Error checking for the presence of an existing Soft-Deleted Key Vault %q (Location %q): %+v", name, location, err)
		}
	}

	// if so, does the user want us to recover it?
	recoverSoftDeletedKeyVault := false
	if !utils.ResponseWasNotFound(softDeletedKeyVault.Response) {
		if !meta.(*clients.Client).Features.KeyVault.RecoverSoftDeletedKeyVaults {
			// this exists but the users opted out so they must import this it out-of-band
			return fmt.Errorf(optedOutOfRecoveringSoftDeletedKeyVaultErrorFmt(name, location))
		}

		recoverSoftDeletedKeyVault = true
	}

	tenantUUID := uuid.FromStringOrNil(d.Get("tenant_id").(string))
	enabledForDeployment := d.Get("enabled_for_deployment").(bool)
	enabledForDiskEncryption := d.Get("enabled_for_disk_encryption").(bool)
	enabledForTemplateDeployment := d.Get("enabled_for_template_deployment").(bool)
	t := d.Get("tags").(map[string]interface{})

	policies := d.Get("access_policy").([]interface{})
	accessPolicies, err := azure.ExpandKeyVaultAccessPolicies(policies)
	if err != nil {
		return fmt.Errorf("Error expanding `access_policy`: %+v", err)
	}

	networkAclsRaw := d.Get("network_acls").([]interface{})
	networkAcls, subnetIds := expandKeyVaultNetworkAcls(networkAclsRaw)

	sku := keyvault.Sku{
		Family: &armKeyVaultSkuFamily,
		Name:   keyvault.SkuName(d.Get("sku_name").(string)),
	}

	parameters := keyvault.VaultCreateOrUpdateParameters{
		Location: &location,
		Properties: &keyvault.VaultProperties{
			TenantID:                     &tenantUUID,
			Sku:                          &sku,
			AccessPolicies:               accessPolicies,
			EnabledForDeployment:         &enabledForDeployment,
			EnabledForDiskEncryption:     &enabledForDiskEncryption,
			EnabledForTemplateDeployment: &enabledForTemplateDeployment,
			NetworkAcls:                  networkAcls,
		},
		Tags: tags.Expand(t),
	}

	// This settings can only be set if it is true, if set when value is false API returns errors
	if softDeleteEnabled := d.Get("soft_delete_enabled").(bool); softDeleteEnabled {
		parameters.Properties.EnableSoftDelete = utils.Bool(softDeleteEnabled)
	}
	if purgeProtectionEnabled := d.Get("purge_protection_enabled").(bool); purgeProtectionEnabled {
		parameters.Properties.EnablePurgeProtection = utils.Bool(purgeProtectionEnabled)
	}

	if recoverSoftDeletedKeyVault {
		parameters.Properties.CreateMode = keyvault.CreateModeRecover
	}

	// also lock on the Virtual Network ID's since modifications in the networking stack are exclusive
	virtualNetworkNames := make([]string, 0)
	for _, v := range subnetIds {
		id, err2 := azure.ParseAzureResourceID(v)
		if err2 != nil {
			return err2
		}

		virtualNetworkName := id.Path["virtualNetworks"]
		if !azure.SliceContainsValue(virtualNetworkNames, virtualNetworkName) {
			virtualNetworkNames = append(virtualNetworkNames, virtualNetworkName)
		}
	}

	locks.MultipleByName(&virtualNetworkNames, network.VirtualNetworkResourceName)
	defer locks.UnlockMultipleByName(&virtualNetworkNames, network.VirtualNetworkResourceName)

	if _, err := client.CreateOrUpdate(ctx, resourceGroup, name, parameters); err != nil {
		return fmt.Errorf("Error creating Key Vault %q (Resource Group %q): %+v", name, resourceGroup, err)
	}

	read, err := client.Get(ctx, resourceGroup, name)
	if err != nil {
		return fmt.Errorf("Error retrieving Key Vault %q (Resource Group %q): %+v", name, resourceGroup, err)
	}
	if read.ID == nil {
		return fmt.Errorf("Cannot read KeyVault %s (resource Group %q) ID", name, resourceGroup)
	}

	d.SetId(*read.ID)

	if props := read.Properties; props != nil {
		if vault := props.VaultURI; vault != nil {
			log.Printf("[DEBUG] Waiting for Key Vault %q (Resource Group %q) to become available", name, resourceGroup)
			stateConf := &resource.StateChangeConf{
				Pending:                   []string{"pending"},
				Target:                    []string{"available"},
				Refresh:                   keyVaultRefreshFunc(*vault),
				Delay:                     30 * time.Second,
				PollInterval:              10 * time.Second,
				ContinuousTargetOccurence: 10,
				Timeout:                   d.Timeout(schema.TimeoutCreate),
			}

			if _, err := stateConf.WaitForState(); err != nil {
				return fmt.Errorf("Error waiting for Key Vault %q (Resource Group %q) to become available: %s", name, resourceGroup, err)
			}
		}
	}

	return resourceArmKeyVaultRead(d, meta)
}

func resourceArmKeyVaultUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).KeyVault.VaultsClient
	ctx, cancel := timeouts.ForUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := azure.ParseAzureResourceID(d.Id())
	if err != nil {
		return err
	}
	resourceGroup := id.ResourceGroup
	name := id.Path["vaults"]

	// Locking this resource so we don't make modifications to it at the same time if there is a
	// key vault access policy trying to update it as well
	locks.ByName(name, keyVaultResourceName)
	defer locks.UnlockByName(name, keyVaultResourceName)

	d.Partial(true)

	// first pull the existing key vault since we need to lock on several bits of it's information
	existing, err := client.Get(ctx, resourceGroup, name)
	if err != nil {
		return fmt.Errorf("Error retrieving Key Vault %q (Resource Group %q): %+v", name, resourceGroup, err)
	}
	if existing.Properties == nil {
		return fmt.Errorf("Error retrieving Key Vault %q (Resource Group %q): `properties` was nil", name, resourceGroup)
	}

	update := keyvault.VaultPatchParameters{}

	if d.HasChange("access_policy") {
		if update.Properties == nil {
			update.Properties = &keyvault.VaultPatchProperties{}
		}

		policiesRaw := d.Get("access_policy").([]interface{})
		accessPolicies, err := azure.ExpandKeyVaultAccessPolicies(policiesRaw)
		if err != nil {
			return fmt.Errorf("Error expanding `access_policy`: %+v", err)
		}
		update.Properties.AccessPolicies = accessPolicies
	}

	if d.HasChange("enabled_for_deployment") {
		if update.Properties == nil {
			update.Properties = &keyvault.VaultPatchProperties{}
		}

		update.Properties.EnabledForDeployment = utils.Bool(d.Get("enabled_for_deployment").(bool))
	}

	if d.HasChange("enabled_for_disk_encryption") {
		if update.Properties == nil {
			update.Properties = &keyvault.VaultPatchProperties{}
		}

		update.Properties.EnabledForDiskEncryption = utils.Bool(d.Get("enabled_for_disk_encryption").(bool))
	}

	if d.HasChange("enabled_for_template_deployment") {
		if update.Properties == nil {
			update.Properties = &keyvault.VaultPatchProperties{}
		}

		update.Properties.EnabledForTemplateDeployment = utils.Bool(d.Get("enabled_for_template_deployment").(bool))
	}

	if d.HasChange("network_acls") {
		if update.Properties == nil {
			update.Properties = &keyvault.VaultPatchProperties{}
		}

		networkAclsRaw := d.Get("network_acls").([]interface{})
		networkAcls, subnetIds := expandKeyVaultNetworkAcls(networkAclsRaw)

		// also lock on the Virtual Network ID's since modifications in the networking stack are exclusive
		virtualNetworkNames := make([]string, 0)
		for _, v := range subnetIds {
			id, err2 := azure.ParseAzureResourceID(v)
			if err2 != nil {
				return err2
			}

			virtualNetworkName := id.Path["virtualNetworks"]
			if !azure.SliceContainsValue(virtualNetworkNames, virtualNetworkName) {
				virtualNetworkNames = append(virtualNetworkNames, virtualNetworkName)
			}
		}

		locks.MultipleByName(&virtualNetworkNames, network.VirtualNetworkResourceName)
		defer locks.UnlockMultipleByName(&virtualNetworkNames, network.VirtualNetworkResourceName)

		update.Properties.NetworkAcls = networkAcls
	}

	if d.HasChange("purge_protection_enabled") {
		if update.Properties == nil {
			update.Properties = &keyvault.VaultPatchProperties{}
		}

		newValue := d.Get("purge_protection_enabled").(bool)

		// existing.Properties guaranteed non-nil above
		oldValue := false
		if existing.Properties.EnablePurgeProtection != nil {
			oldValue = *existing.Properties.EnablePurgeProtection
		}

		// whilst this should have got caught in the customizeDiff this won't work if that fields interpolated
		// hence the double-checking here
		if oldValue && !newValue {
			return fmt.Errorf("Error updating Key Vault %q (Resource Group %q): once Purge Protection has been Enabled it's not possible to disable it", name, resourceGroup)
		}

		update.Properties.EnablePurgeProtection = utils.Bool(newValue)
	}

	if d.HasChange("sku_name") {
		if update.Properties == nil {
			update.Properties = &keyvault.VaultPatchProperties{}
		}

		update.Properties.Sku = &keyvault.Sku{
			Family: &armKeyVaultSkuFamily,
			Name:   keyvault.SkuName(d.Get("sku_name").(string)),
		}
	}

	if d.HasChange("soft_delete_enabled") {
		if update.Properties == nil {
			update.Properties = &keyvault.VaultPatchProperties{}
		}

		newValue := d.Get("soft_delete_enabled").(bool)

		// existing.Properties guaranteed non-nil above
		oldValue := false
		if existing.Properties.EnableSoftDelete != nil {
			oldValue = *existing.Properties.EnableSoftDelete
		}

		// whilst this should have got caught in the customizeDiff this won't work if that fields interpolated
		// hence the double-checking here
		if oldValue && !newValue {
			return fmt.Errorf("Error updating Key Vault %q (Resource Group %q): once Soft Delete has been Enabled it's not possible to disable it", name, resourceGroup)
		}

		update.Properties.EnableSoftDelete = utils.Bool(newValue)
	}

	if d.HasChange("tenant_id") {
		if update.Properties == nil {
			update.Properties = &keyvault.VaultPatchProperties{}
		}

		tenantUUID := uuid.FromStringOrNil(d.Get("tenant_id").(string))
		update.Properties.TenantID = &tenantUUID
	}

	if d.HasChange("tags") {
		t := d.Get("tags").(map[string]interface{})
		update.Tags = tags.Expand(t)
	}

	if _, err := client.Update(ctx, resourceGroup, name, update); err != nil {
		return fmt.Errorf("Error updating Key Vault %q (Resource Group %q): %+v", name, resourceGroup, err)
	}

	d.Partial(false)

	return resourceArmKeyVaultRead(d, meta)
}

func resourceArmKeyVaultRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).KeyVault.VaultsClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := azure.ParseAzureResourceID(d.Id())
	if err != nil {
		return err
	}
	resourceGroup := id.ResourceGroup
	name := id.Path["vaults"]

	resp, err := client.Get(ctx, resourceGroup, name)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			log.Printf("[DEBUG] Key Vault %q was not found in Resource Group %q - removing from state!", name, resourceGroup)
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error making Read request on KeyVault %q (Resource Group %q): %+v", name, resourceGroup, err)
	}

	d.Set("name", resp.Name)
	d.Set("resource_group_name", resourceGroup)
	if location := resp.Location; location != nil {
		d.Set("location", azure.NormalizeLocation(*location))
	}

	if props := resp.Properties; props != nil {
		d.Set("tenant_id", props.TenantID.String())
		d.Set("enabled_for_deployment", props.EnabledForDeployment)
		d.Set("enabled_for_disk_encryption", props.EnabledForDiskEncryption)
		d.Set("enabled_for_template_deployment", props.EnabledForTemplateDeployment)
		d.Set("soft_delete_enabled", props.EnableSoftDelete)
		d.Set("purge_protection_enabled", props.EnablePurgeProtection)
		d.Set("vault_uri", props.VaultURI)

		if sku := props.Sku; sku != nil {
			if err := d.Set("sku_name", string(sku.Name)); err != nil {
				return fmt.Errorf("Error setting 'sku_name' for KeyVault %q: %+v", *resp.Name, err)
			}
		} else {
			return fmt.Errorf("Error making Read request on KeyVault %q (Resource Group %q): Unable to retrieve 'sku' value", name, resourceGroup)
		}

		if err := d.Set("network_acls", flattenKeyVaultNetworkAcls(props.NetworkAcls)); err != nil {
			return fmt.Errorf("Error setting `network_acls` for KeyVault %q: %+v", *resp.Name, err)
		}

		flattenedPolicies := azure.FlattenKeyVaultAccessPolicies(props.AccessPolicies)
		if err := d.Set("access_policy", flattenedPolicies); err != nil {
			return fmt.Errorf("Error setting `access_policy` for KeyVault %q: %+v", *resp.Name, err)
		}
	}

	return tags.FlattenAndSet(d, resp.Tags)
}

func resourceArmKeyVaultDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).KeyVault.VaultsClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := azure.ParseAzureResourceID(d.Id())
	if err != nil {
		return err
	}

	resourceGroup := id.ResourceGroup
	name := id.Path["vaults"]
	location := d.Get("location").(string)

	locks.ByName(name, keyVaultResourceName)
	defer locks.UnlockByName(name, keyVaultResourceName)

	read, err := client.Get(ctx, resourceGroup, name)
	if err != nil {
		if utils.ResponseWasNotFound(read.Response) {
			return nil
		}

		return fmt.Errorf("Error retrieving Key Vault %q (Resource Group %q): %+v", name, resourceGroup, err)
	}

	if read.Properties == nil {
		return fmt.Errorf("Error retrieving Key Vault %q (Resource Group %q): `properties` was nil", name, resourceGroup)
	}

	// Check to see if purge protection is enabled or not...
	purgeProtectionEnabled := false
	if ppe := read.Properties.EnablePurgeProtection; ppe != nil {
		purgeProtectionEnabled = *ppe
	}
	softDeleteEnabled := false
	if sde := read.Properties.EnableSoftDelete; sde != nil {
		softDeleteEnabled = *sde
	}

	// ensure we lock on the latest network names, to ensure we handle Azure's networking layer being limited to one change at a time
	virtualNetworkNames := make([]string, 0)
	if props := read.Properties; props != nil {
		if acls := props.NetworkAcls; acls != nil {
			if rules := acls.VirtualNetworkRules; rules != nil {
				for _, v := range *rules {
					if v.ID == nil {
						continue
					}

					id, err2 := azure.ParseAzureResourceID(*v.ID)
					if err2 != nil {
						return err2
					}

					virtualNetworkName := id.Path["virtualNetworks"]
					if !azure.SliceContainsValue(virtualNetworkNames, virtualNetworkName) {
						virtualNetworkNames = append(virtualNetworkNames, virtualNetworkName)
					}
				}
			}
		}
	}

	locks.MultipleByName(&virtualNetworkNames, network.VirtualNetworkResourceName)
	defer locks.UnlockMultipleByName(&virtualNetworkNames, network.VirtualNetworkResourceName)

	resp, err := client.Delete(ctx, resourceGroup, name)
	if err != nil {
		if !response.WasNotFound(resp.Response) {
			return fmt.Errorf("Error retrieving Key Vault %q (Resource Group %q): %+v", name, resourceGroup, err)
		}
	}

	// Purge the soft deleted key vault permanently if the feature flag is enabled
	if meta.(*clients.Client).Features.KeyVault.PurgeSoftDeleteOnDestroy && softDeleteEnabled {
		// KeyVaults with Purge Protection Enabled cannot be deleted unless done by Azure
		if purgeProtectionEnabled {
			deletedInfo, err := getSoftDeletedStateForKeyVault(ctx, client, name, location)
			if err != nil {
				return fmt.Errorf("Error retrieving the Deletion Details for KeyVault %q: %+v", name, err)
			}

			// in the future it'd be nice to raise a warning, but this is the best we can do for now
			if deletedInfo != nil {
				log.Printf("[DEBUG] The Key Vault %q has Purge Protection Enabled and was deleted on %q. Azure will purge this on %q", name, deletedInfo.deleteDate, deletedInfo.purgeDate)
			} else {
				log.Printf("[DEBUG] The Key Vault %q has Purge Protection Enabled and will be purged automatically by Azure", name)
			}
			return nil
		}

		log.Printf("[DEBUG] KeyVault %q marked for purge - executing purge", name)
		future, err := client.PurgeDeleted(ctx, name, location)
		if err != nil {
			return err
		}

		log.Printf("[DEBUG] Waiting for purge of KeyVault %q..", name)
		err = future.WaitForCompletionRef(ctx, client.Client)
		if err != nil {
			return fmt.Errorf("Error purging Key Vault %q (Resource Group %q): %+v", name, resourceGroup, err)
		}
		log.Printf("[DEBUG] Purged KeyVault %q.", name)
	}

	return nil
}

func keyVaultRefreshFunc(vaultUri string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] Checking to see if KeyVault %q is available..", vaultUri)

		var PTransport = &http.Transport{Proxy: http.ProxyFromEnvironment}

		client := &http.Client{
			Transport: PTransport,
		}

		conn, err := client.Get(vaultUri)
		if err != nil {
			log.Printf("[DEBUG] Didn't find KeyVault at %q", vaultUri)
			return nil, "pending", fmt.Errorf("Error connecting to %q: %s", vaultUri, err)
		}

		defer conn.Body.Close()

		log.Printf("[DEBUG] Found KeyVault at %q", vaultUri)
		return "available", "available", nil
	}
}

func expandKeyVaultNetworkAcls(input []interface{}) (*keyvault.NetworkRuleSet, []string) {
	subnetIds := make([]string, 0)
	if len(input) == 0 {
		return nil, subnetIds
	}

	v := input[0].(map[string]interface{})

	bypass := v["bypass"].(string)
	defaultAction := v["default_action"].(string)

	ipRulesRaw := v["ip_rules"].(*schema.Set)
	ipRules := make([]keyvault.IPRule, 0)

	for _, v := range ipRulesRaw.List() {
		rule := keyvault.IPRule{
			Value: utils.String(v.(string)),
		}
		ipRules = append(ipRules, rule)
	}

	networkRulesRaw := v["virtual_network_subnet_ids"].(*schema.Set)
	networkRules := make([]keyvault.VirtualNetworkRule, 0)
	for _, v := range networkRulesRaw.List() {
		rawId := v.(string)
		subnetIds = append(subnetIds, rawId)
		rule := keyvault.VirtualNetworkRule{
			ID: utils.String(rawId),
		}
		networkRules = append(networkRules, rule)
	}

	ruleSet := keyvault.NetworkRuleSet{
		Bypass:              keyvault.NetworkRuleBypassOptions(bypass),
		DefaultAction:       keyvault.NetworkRuleAction(defaultAction),
		IPRules:             &ipRules,
		VirtualNetworkRules: &networkRules,
	}
	return &ruleSet, subnetIds
}

func flattenKeyVaultNetworkAcls(input *keyvault.NetworkRuleSet) []interface{} {
	if input == nil {
		return []interface{}{
			map[string]interface{}{
				"bypass":                     string(keyvault.AzureServices),
				"default_action":             string(keyvault.Allow),
				"ip_rules":                   schema.NewSet(schema.HashString, []interface{}{}),
				"virtual_network_subnet_ids": schema.NewSet(schema.HashString, []interface{}{}),
			},
		}
	}

	output := make(map[string]interface{})

	output["bypass"] = string(input.Bypass)
	output["default_action"] = string(input.DefaultAction)

	ipRules := make([]interface{}, 0)
	if input.IPRules != nil {
		for _, v := range *input.IPRules {
			if v.Value == nil {
				continue
			}

			ipRules = append(ipRules, *v.Value)
		}
	}
	output["ip_rules"] = schema.NewSet(schema.HashString, ipRules)

	virtualNetworkRules := make([]interface{}, 0)
	if input.VirtualNetworkRules != nil {
		for _, v := range *input.VirtualNetworkRules {
			if v.ID == nil {
				continue
			}

			virtualNetworkRules = append(virtualNetworkRules, *v.ID)
		}
	}
	output["virtual_network_subnet_ids"] = schema.NewSet(schema.HashString, virtualNetworkRules)

	return []interface{}{output}
}

func optedOutOfRecoveringSoftDeletedKeyVaultErrorFmt(name, location string) string {
	return fmt.Sprintf(`
An existing soft-deleted Key Vault exists with the Name %q in the location %q, however
automatically recovering this KeyVault has been disabled via the "features" block.

Terraform can automatically recover the soft-deleted Key Vault when this behaviour is
enabled within the "features" block (located within the "provider" block) - more
information can be found here:

https://www.terraform.io/docs/providers/azurerm/index.html#features

Alternatively you can manually recover this (e.g. using the Azure CLI) and then import
this into Terraform via "terraform import", or pick a different name/location.
`, name, location)
}

type keyVaultDeletionStatus struct {
	deleteDate string
	purgeDate  string
}

func getSoftDeletedStateForKeyVault(ctx context.Context, client *keyvault.VaultsClient, name string, location string) (*keyVaultDeletionStatus, error) {
	softDel, err := client.GetDeleted(ctx, name, location)
	if err != nil {
		return nil, err
	}

	// we found an existing key vault that is not soft deleted
	if softDel.Properties == nil {
		return nil, nil
	}

	// the logic is this way because the GetDeleted call will return an existing key vault
	// that is not soft deleted, but the Deleted Vault properties will be nil
	props := *softDel.Properties

	result := keyVaultDeletionStatus{}
	if props.DeletionDate != nil {
		result.deleteDate = (*props.DeletionDate).Format(time.RFC3339)
	}
	if props.ScheduledPurgeDate != nil {
		result.purgeDate = (*props.ScheduledPurgeDate).Format(time.RFC3339)
	}

	if result.deleteDate == "" && result.purgeDate == "" {
		return nil, nil
	}

	return &result, nil
}
