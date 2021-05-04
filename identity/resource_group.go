package identity

import (
	"context"
	"log"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceGroup manages user groups
func ResourceGroup() *schema.Resource {
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
		if err = d.Set("allow_cluster_create", isGroupClusterCreateEntitled(&group)); err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("allow_sql_analytics_access", isGroupSQLAnalyticsAccessEntitled(&group)); err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("allow_instance_pool_create", isGroupInstancePoolCreateEntitled(&group)); err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("allow_workspace_access", isGroupWorkspaceAccessEntitled(&group)); err != nil {
			return diag.FromErr(err)
		}
		d.Set("url", m.(*common.DatabricksClient).FormatURL("#setting/accounts/groups/", d.Id()))
		return nil
	}
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			groupName := d.Get("display_name").(string)
			allowClusterCreate := d.Get("allow_cluster_create").(bool)
			allowInstancePoolCreate := d.Get("allow_instance_pool_create").(bool)
			allowSQLAnalyticsAccess := d.Get("allow_sql_analytics_access").(bool)
			allowWorkspaceAccess := d.Get("allow_workspace_access").(bool)

			// If entitlement flags are set to be true
			var entitlementsList []string
			if allowClusterCreate {
				entitlementsList = append(entitlementsList, string(AllowClusterCreateEntitlement))
			}
			if allowSQLAnalyticsAccess {
				entitlementsList = append(entitlementsList, string(AllowSQLAnalyticsAccessEntitlement))
			}
			if allowInstancePoolCreate {
				entitlementsList = append(entitlementsList, string(AllowInstancePoolCreateEntitlement))
			}
			if allowWorkspaceAccess {
				entitlementsList = append(entitlementsList, string(AllowWorkspaceAccessEntitlement))
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
			// If allow_cluster_create has changed
			if d.HasChange("allow_cluster_create") {
				allowClusterCreate := d.Get("allow_cluster_create").(bool)
				// Changed to true
				if allowClusterCreate {
					entitlementsAddList = append(entitlementsAddList, string(AllowClusterCreateEntitlement))
				} else {
					// Changed to false
					entitlementsRemoveList = append(entitlementsRemoveList, string(AllowClusterCreateEntitlement))
				}
			}
			// If allow_sql_analytics_access has changed
			if d.HasChange("allow_sql_analytics_access") {
				allowSQLAnalyticsAccess := d.Get("allow_sql_analytics_access").(bool)
				// Changed to true
				if allowSQLAnalyticsAccess {
					entitlementsAddList = append(entitlementsAddList, string(AllowSQLAnalyticsAccessEntitlement))
				} else {
					// Changed to false
					entitlementsRemoveList = append(entitlementsRemoveList, string(AllowSQLAnalyticsAccessEntitlement))
				}
			}
			// If allow_instance_pool_create has changed
			if d.HasChange("allow_instance_pool_create") {
				allowClusterCreate := d.Get("allow_instance_pool_create").(bool)
				// Changed to true
				if allowClusterCreate {
					entitlementsAddList = append(entitlementsAddList, string(AllowInstancePoolCreateEntitlement))
				} else {
					// Changed to false
					entitlementsRemoveList = append(entitlementsRemoveList, string(AllowInstancePoolCreateEntitlement))
				}
			}
			// If allow_instance_pool_create has changed
			if d.HasChange("allow_workspace_access") {
				allowClusterCreate := d.Get("allow_workspace_access").(bool)
				// Changed to true
				if allowClusterCreate {
					entitlementsAddList = append(entitlementsAddList, string(AllowWorkspaceAccessEntitlement))
				} else {
					// Changed to false
					entitlementsRemoveList = append(entitlementsRemoveList, string(AllowWorkspaceAccessEntitlement))
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
			"allow_cluster_create": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"allow_sql_analytics_access": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"allow_instance_pool_create": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"allow_workspace_access": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func isGroupClusterCreateEntitled(group *ScimGroup) bool {
	for _, entitlement := range group.Entitlements {
		if entitlement.Value == AllowClusterCreateEntitlement {
			return true
		}
	}
	return false
}

func isGroupSQLAnalyticsAccessEntitled(group *ScimGroup) bool {
	for _, entitlement := range group.Entitlements {
		if entitlement.Value == AllowSQLAnalyticsAccessEntitlement {
			return true
		}
	}
	return false
}

func isGroupInstancePoolCreateEntitled(group *ScimGroup) bool {
	for _, entitlement := range group.Entitlements {
		if entitlement.Value == AllowInstancePoolCreateEntitlement {
			return true
		}
	}
	return false
}

func isGroupWorkspaceAccessEntitled(group *ScimGroup) bool {
	for _, entitlement := range group.Entitlements {
		if entitlement.Value == AllowWorkspaceAccessEntitlement {
			return true
		}
	}
	return false
}
