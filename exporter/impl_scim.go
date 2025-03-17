package exporter

import (
	"fmt"
	"log"
	"strings"

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
	if (!ic.accountLevel && (groupName == "admins" || groupName == "users")) ||
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
	if err := ic.cacheGroups(); err != nil {
		return err
	}
	for _, g := range ic.allGroups {
		if r.ID != g.ID {
			continue
		}
		ic.emitRoles("group", g.ID, g.Roles)
		builtInUserGroup := (ic.accountLevel && g.DisplayName == "account users") || (!ic.accountLevel && g.DisplayName == "users")
		if builtInUserGroup && !ic.importAllUsers {
			log.Printf("[INFO] Skipping import of entire user directory ...")
			continue
		}
		if len(g.Members) > 10 {
			log.Printf("[INFO] Importing %d members of %s", len(g.Members), g.DisplayName)
		}
		for _, parent := range g.Groups {
			ic.Emit(&resource{
				Resource: "databricks_group",
				ID:       parent.Value,
			})
			if parent.Type == "direct" {
				id := fmt.Sprintf("%s|%s", parent.Value, g.ID)
				ic.Emit(&resource{
					Resource: "databricks_group_member",
					ID:       id,
					Name:     fmt.Sprintf("%s_%s_%s", parent.Display, parent.Value, g.DisplayName),
					Data:     ic.makeGroupMemberData(id, parent.Value, g.ID),
				})
			}
		}
		for i, x := range g.Members {
			if strings.HasPrefix(x.Ref, "Users/") {
				ic.Emit(&resource{
					Resource: "databricks_user",
					ID:       x.Value,
				})
				if !builtInUserGroup {
					id := fmt.Sprintf("%s|%s", g.ID, x.Value)
					ic.Emit(&resource{
						Resource: "databricks_group_member",
						ID:       id,
						Name:     fmt.Sprintf("%s_%s_%s_%s", g.DisplayName, g.ID, x.Display, x.Value),
						Data:     ic.makeGroupMemberData(id, g.ID, x.Value),
					})
				}
			}
			if strings.HasPrefix(x.Ref, "ServicePrincipals/") {
				ic.Emit(&resource{
					Resource: "databricks_service_principal",
					ID:       x.Value,
				})
				if !builtInUserGroup {
					id := fmt.Sprintf("%s|%s", g.ID, x.Value)
					ic.Emit(&resource{
						Resource: "databricks_group_member",
						ID:       id,
						Name:     fmt.Sprintf("%s_%s_%s_%s", g.DisplayName, g.ID, x.Display, x.Value),
						Data:     ic.makeGroupMemberData(id, g.ID, x.Value),
					})
				}
			}
			if strings.HasPrefix(x.Ref, "Groups/") {
				ic.Emit(&resource{
					Resource: "databricks_group",
					ID:       x.Value,
				})
				if !builtInUserGroup {
					id := fmt.Sprintf("%s|%s", g.ID, x.Value)
					ic.Emit(&resource{
						Resource: "databricks_group_member",
						ID:       id,
						Name:     fmt.Sprintf("%s_%s_%s_%s", g.DisplayName, g.ID, x.Display, x.Value),
						Data:     ic.makeGroupMemberData(id, g.ID, x.Value),
					})
				}
			}
			if len(g.Members) > 10 {
				log.Printf("[INFO] Imported %d of %d members of %s", i+1, len(g.Members), g.DisplayName)
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
	r.Data.Set("force", true)
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
	r.Data.Set("force", true)
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
