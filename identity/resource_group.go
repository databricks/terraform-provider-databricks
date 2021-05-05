package identity

import (
	"context"
	"log"
	"sort"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceGroup manages user groups
func ResourceGroup() *schema.Resource {
	entitlementsMap := map[string]Entitlement{
		"allow_cluster_create":       AllowClusterCreateEntitlement,
		"allow_instance_pool_create": AllowInstancePoolCreateEntitlement,
		"allow_sql_analytics_access": AllowSQLAnalyticsAccessEntitlement,
		"allow_workspace_access":     AllowWorkspaceAccessEntitlement,
	}
	//To make sure the order of fields are in consistent sorted order and make the testing accurate
	entitlementFields := getSortedKeys(entitlementsMap)
	readContext := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		group, err := NewGroupsAPI(ctx, m).Read(d.Id())
		if err != nil {
			if e, ok := err.(common.APIError); ok && e.IsMissing() {
				log.Printf("missing resource due to error: %v\n", e)
				d.SetId("")
				return nil
			}
			return diag.FromErr(err)
		}
		if err = d.Set("display_name", group.DisplayName); err != nil {
			return diag.FromErr(err)
		}
		for _, entitlementField := range entitlementFields {
			if err = d.Set(entitlementField, groupEntitlementExists(&group, entitlementsMap[entitlementField])); err != nil {
				return diag.FromErr(err)
			}
		}
		d.Set("url", m.(*common.DatabricksClient).FormatURL("#setting/accounts/groups/", d.Id()))
		return nil
	}
	groupResource := &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			groupName := d.Get("display_name").(string)
			var entitlementsList []string
			for _, entitlementField := range entitlementFields {
				fieldValue := d.Get(entitlementField).(bool)
				if fieldValue {
					entitlementsList = append(entitlementsList, string(entitlementsMap[entitlementField]))
				}
			}
			group, err := NewGroupsAPI(ctx, m).Create(groupName, nil, nil, entitlementsList)
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId(group.ID)
			return readContext(ctx, d, m)
		},
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			// Handle entitlements update
			var entitlementsAddList []string
			var entitlementsRemoveList []string
			for _, entitlementField := range entitlementFields {
				// If entitlement field has changed
				if d.HasChange(entitlementField) {
					fieldValue := d.Get(entitlementField).(bool)
					// Changed to true
					if fieldValue {
						entitlementsAddList = append(entitlementsAddList, string(entitlementsMap[entitlementField]))
					} else {
						// Changed to false
						entitlementsRemoveList = append(entitlementsRemoveList, string(entitlementsMap[entitlementField]))
					}
				}
			}
			// TODO: not currently possible to update group display name
			if entitlementsAddList != nil || entitlementsRemoveList != nil {
				if err := NewGroupsAPI(ctx, m).Patch(d.Id(),
					entitlementsAddList, entitlementsRemoveList,
					GroupEntitlementsPath); err != nil {
					return diag.FromErr(err)
				}
			}
			return nil
		},
		ReadContext: readContext,
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			if err := NewGroupsAPI(ctx, m).Delete(d.Id()); err != nil {
				return diag.FromErr(err)
			}
			return nil
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
	//Add entitlements to the schema
	for _, entitlementField := range entitlementFields {
		groupResource.Schema[entitlementField] = &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		}
	}
	return groupResource
}

func groupEntitlementExists(group *ScimGroup, entitlement Entitlement) bool {
	for _, groupEntitlement := range group.Entitlements {
		if groupEntitlement.Value == entitlement {
			return true
		}
	}
	return false
}

func getSortedKeys(m map[string]Entitlement) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys
}
