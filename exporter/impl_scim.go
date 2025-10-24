package exporter

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/databricks/terraform-provider-databricks/mws"
	"github.com/databricks/terraform-provider-databricks/scim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"golang.org/x/exp/maps"
)

func listGroups(ic *importContext) error {
	if err := ic.cacheGroups(); err != nil {
		return err
	}
	for offset, g := range ic.allGroups {
		if !ic.MatchesName(g.DisplayName) {
			log.Printf("[INFO] Group %s doesn't match %s filter", g.DisplayName, ic.match)
			continue
		}
		ic.Emit(&resource{
			Resource: "databricks_group",
			ID:       g.ID,
		})
		log.Printf("[INFO] Scanned %d of %d groups", offset+1, len(ic.allGroups))
	}
	return nil
}

func searchGroup(ic *importContext, r *resource) error {
	if r.Attribute != "display_name" {
		return fmt.Errorf("wrong search attribute '%s' for databricks_group", r.Attribute)
	}
	if err := ic.cacheGroups(); err != nil {
		return err
	}
	for _, g := range ic.allGroups {
		if g.DisplayName == r.Value {
			r.ID = g.ID
			return nil
		}
	}
	return nil
}

func importGroup(ic *importContext, r *resource) error {
	groupName := r.Data.Get("display_name").(string)
	if err := ic.cacheGroups(); err != nil {
		return err
	}
	var group *scim.Group
	for _, g := range ic.allGroups {
		if r.ID == g.ID {
			group = &g
			break
		}
	}
	if group == nil {
		return fmt.Errorf("group %s not found", r.ID)
	}
	isAccountLevelGroup := ic.currentMetastore != nil && group.Meta != nil && group.Meta.ResourceType == "Group"
	if (!ic.accountLevel && (isAccountLevelGroup || groupName == "admins" || groupName == "users")) ||
		(ic.accountLevel && groupName == "account users") {
		// Workspace admins & users or Account users are to be imported through "data block"
		r.Mode = "data"
		r.Data.Set("workspace_access", false)
		r.Data.Set("databricks_sql_access", false)
		r.Data.Set("allow_instance_pool_create", false)
		r.Data.Set("allow_cluster_create", false)
		r.Data.State().Set(&terraform.InstanceState{
			ID: r.ID,
			Attributes: map[string]string{
				"display_name": r.Name,
			},
		})
	} else if r.Data != nil {
		r.Data.Set("force", true)
	}
	// process group data
	ic.emitRoles("group", group.ID, group.Roles)
	builtInUserGroup := (ic.accountLevel && group.DisplayName == "account users") || (!ic.accountLevel && group.DisplayName == "users")
	if builtInUserGroup && !ic.importAllUsers {
		log.Printf("[INFO] Skipping import of entire user directory ...")
		return nil
	}
	// process group members only if they are on account level or is not an account level group
	if ic.accountLevel || !isAccountLevelGroup {
		if len(group.Members) > 10 {
			log.Printf("[INFO] Importing %d members of %s", len(group.Members), group.DisplayName)
		}
		for _, parent := range group.Groups {
			ic.Emit(&resource{
				Resource: "databricks_group",
				ID:       parent.Value,
			})
			if parent.Type == "direct" {
				id := fmt.Sprintf("%s|%s", parent.Value, group.ID)
				ic.Emit(&resource{
					Resource: "databricks_group_member",
					ID:       id,
					Name:     fmt.Sprintf("%s_%s_%s", parent.Display, parent.Value, group.DisplayName),
					Data:     ic.makeGroupMemberData(id, parent.Value, group.ID),
				})
			}
		}
		for i, x := range group.Members {
			if strings.HasPrefix(x.Ref, "Users/") {
				ic.Emit(&resource{
					Resource: "databricks_user",
					ID:       x.Value,
				})
				if !builtInUserGroup {
					id := fmt.Sprintf("%s|%s", group.ID, x.Value)
					ic.Emit(&resource{
						Resource: "databricks_group_member",
						ID:       id,
						Name:     fmt.Sprintf("%s_%s_%s_%s", group.DisplayName, group.ID, x.Display, x.Value),
						Data:     ic.makeGroupMemberData(id, group.ID, x.Value),
					})
				}
			}
			if strings.HasPrefix(x.Ref, "ServicePrincipals/") {
				ic.Emit(&resource{
					Resource: "databricks_service_principal",
					ID:       x.Value,
				})
				if !builtInUserGroup {
					id := fmt.Sprintf("%s|%s", group.ID, x.Value)
					ic.Emit(&resource{
						Resource: "databricks_group_member",
						ID:       id,
						Name:     fmt.Sprintf("%s_%s_%s_%s", group.DisplayName, group.ID, x.Display, x.Value),
						Data:     ic.makeGroupMemberData(id, group.ID, x.Value),
					})
				}
			}
			if strings.HasPrefix(x.Ref, "Groups/") {
				ic.Emit(&resource{
					Resource: "databricks_group",
					ID:       x.Value,
				})
				if !builtInUserGroup {
					id := fmt.Sprintf("%s|%s", group.ID, x.Value)
					ic.Emit(&resource{
						Resource: "databricks_group_member",
						ID:       id,
						Name:     fmt.Sprintf("%s_%s_%s_%s", group.DisplayName, group.ID, x.Display, x.Value),
						Data:     ic.makeGroupMemberData(id, group.ID, x.Value),
					})
				}
			}
			if len(group.Members) > 10 {
				log.Printf("[INFO] Imported %d of %d members of %s", i+1, len(group.Members), group.DisplayName)
			}
		}
	}

	if ic.accountLevel {
		ic.Emit(&resource{
			Resource: "databricks_access_control_rule_set",
			ID: fmt.Sprintf("accounts/%s/groups/%s/ruleSets/default",
				ic.Client.Config.AccountID, r.ID),
		})
	}

	return nil
}

func listUsers(ic *importContext) error {
	ic.getUsersMapping()
	ic.allUsersMutex.RLocker().Lock()
	userMapping := maps.Clone(ic.allUsersMapping)
	ic.allUsersMutex.RLocker().Unlock()
	for userName, userScimId := range userMapping {
		log.Printf("[TRACE] Emitting user %s, SCIM id=%s", userName, userScimId)
		ic.Emit(&resource{
			Resource: "databricks_user",
			ID:       userScimId,
		})
	}
	return nil
}

func searchUser(ic *importContext, r *resource) error {
	u, err := ic.findUserByName(r.Value, false)
	if err != nil {
		return err
	}
	r.ID = u.ID
	return nil
}

func importUser(ic *importContext, r *resource) error {
	username := r.Data.Get("user_name").(string)
	if ic.currentMetastore != nil {
		// Users are maintained on account level and are referenced via data sources
		r.Mode = "data"
		r.Data.Set("workspace_access", false)
		r.Data.Set("databricks_sql_access", false)
		r.Data.Set("allow_instance_pool_create", false)
		r.Data.Set("allow_cluster_create", false)
		r.Data.Set("display_name", "")
		r.Data.Set("external_id", "")
		r.Data.State().Set(&terraform.InstanceState{
			ID: r.ID,
			Attributes: map[string]string{
				"user_name": username,
			},
		})
	} else if r.Data != nil {
		r.Data.Set("force", true)
	}
	u, err := ic.findUserByName(username, false)
	if err != nil {
		return err
	}
	ic.emitGroups(*u)
	ic.emitRoles("user", u.ID, u.Roles)
	return nil
}

func listServicePrincipals(ic *importContext) error {
	ic.getSpsMapping()
	ic.spsMutex.RLock()
	spsMapping := maps.Clone(ic.allSpsMapping)
	ic.spsMutex.RLocker().Unlock()
	for applicationId, appScimId := range spsMapping {
		log.Printf("[TRACE] Emitting service principal %s, SCIM id=%s", applicationId, appScimId)
		ic.Emit(&resource{
			Resource: "databricks_service_principal",
			ID:       appScimId,
		})
	}
	return nil
}

func searchServicePrincipal(ic *importContext, r *resource) error {
	u, err := ic.findSpnByAppID(r.Value, false)
	if err != nil {
		return err
	}
	r.ID = u.ID
	return nil
}

func importServicePrincipal(ic *importContext, r *resource) error {
	applicationID := r.Data.Get("application_id").(string)
	if ic.currentMetastore != nil {
		// Users are maintained on account level and are referenced via data sources
		r.Mode = "data"
		r.Data.Set("workspace_access", false)
		r.Data.Set("databricks_sql_access", false)
		r.Data.Set("allow_instance_pool_create", false)
		r.Data.Set("allow_cluster_create", false)
		r.Data.Set("display_name", "")
		r.Data.Set("external_id", "")
		r.Data.State().Set(&terraform.InstanceState{
			ID: r.ID,
			Attributes: map[string]string{
				"application_id": applicationID,
			},
		})
	} else if r.Data != nil {
		r.Data.Set("force", true)
	}
	u, err := ic.findSpnByAppID(applicationID, false)
	if err != nil {
		return err
	}
	ic.emitGroups(*u)
	ic.emitRoles("service_principal", u.ID, u.Roles)
	if ic.accountLevel {
		ic.Emit(&resource{
			Resource: "databricks_access_control_rule_set",
			ID: fmt.Sprintf("accounts/%s/servicePrincipals/%s/ruleSets/default",
				ic.Client.Config.AccountID, applicationID),
		})
	}
	return nil
}

func emitIdfedAndUsersSpsGroups(ic *importContext, workspaceId int64) error {
	log.Printf("[DEBUG] Emitting permission assignments for workspace %d", workspaceId)
	pas, err := ic.accountClient.WorkspaceAssignment.ListByWorkspaceId(ic.Context, workspaceId)
	if err != nil {
		log.Printf("[ERROR] listing workspace permission assignments for workspace %d: %s",
			workspaceId, err.Error())
		return err
	}
	for _, pa := range pas.PermissionAssignments {
		perm := "unknown"
		if len(pa.Permissions) > 0 {
			perm = pa.Permissions[0].String()
		}
		nm := fmt.Sprintf("mws_pa_%d_%s_%s_%d", workspaceId, pa.Principal.DisplayName,
			perm, pa.Principal.PrincipalId)
		// We  generate Data directly to avoid calling APIs
		data := mws.ResourceMwsPermissionAssignment().ToResource().TestResourceData()
		paId := fmt.Sprintf("%d|%d", workspaceId, pa.Principal.PrincipalId)
		data = ic.generateNewData(data, "databricks_mws_permission_assignment", paId, pa)
		data.Set("workspace_id", workspaceId)
		data.Set("principal_id", pa.Principal.PrincipalId)
		ic.Emit(&resource{
			Resource: "databricks_mws_permission_assignment",
			ID:       paId,
			Name:     nameNormalizationRegex.ReplaceAllString(nm, "_"),
			Data:     data,
		})
		// Emit principals
		strPrincipalId := strconv.FormatInt(pa.Principal.PrincipalId, 10)
		if pa.Principal.ServicePrincipalName != "" {
			ic.Emit(&resource{
				Resource: "databricks_service_principal",
				ID:       strPrincipalId,
			})
		} else if pa.Principal.UserName != "" {
			ic.Emit(&resource{
				Resource: "databricks_user",
				ID:       strPrincipalId,
			})
		} else if pa.Principal.GroupName != "" {
			ic.Emit(&resource{
				Resource: "databricks_group",
				ID:       strPrincipalId,
			})
		}
	}

	return nil
}
