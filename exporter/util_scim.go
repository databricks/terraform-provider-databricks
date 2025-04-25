package exporter

import (
	"fmt"
	"log"
	"strings"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/scim"

	"github.com/databricks/databricks-sdk-go/service/iam"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	nonExistingUserOrSp = "__USER_OR_SPN_DOES_NOT_EXIST__"
)

func (ic *importContext) emitListOfUsers(users []string) {
	for _, user := range users {
		if user != "" {
			ic.Emit(&resource{
				Resource:  "databricks_user",
				Attribute: "user_name",
				Value:     user,
			})
		}
	}
}

func (ic *importContext) emitUserOrServicePrincipal(userOrSPName string) {
	if userOrSPName == "" || !ic.isServiceEnabled("users") {
		return
	}
	// Cache check here to avoid emitting
	ic.emittedUsersMutex.RLock()
	_, exists := ic.emittedUsers[userOrSPName]
	ic.emittedUsersMutex.RUnlock()
	if exists {
		// log.Printf("[DEBUG] user or SP %s already emitted...", userOrSPName)
		return
	}
	if common.StringIsUUID(userOrSPName) {
		user, err := ic.findSpnByAppID(userOrSPName, false)
		if err != nil {
			log.Printf("[ERROR] Can't find SP with application ID %s", userOrSPName)
			ic.addIgnoredResource(fmt.Sprintf("databricks_service_principal. application_id=%s", userOrSPName))
		} else {
			ic.Emit(&resource{
				Resource: "databricks_service_principal",
				ID:       user.ID,
			})
		}
	} else {
		user, err := ic.findUserByName(strings.ToLower(userOrSPName), false)
		if err != nil {
			log.Printf("[ERROR] Can't find user with name %s", userOrSPName)
			ic.addIgnoredResource(fmt.Sprintf("databricks_user. user_name=%s", userOrSPName))
		} else {
			ic.Emit(&resource{
				Resource: "databricks_user",
				ID:       user.ID,
			})
		}
	}
	ic.emittedUsersMutex.Lock()
	ic.emittedUsers[userOrSPName] = struct{}{}
	ic.emittedUsersMutex.Unlock()
}

func getUserOrSpNameAndDirectory(path, prefix string) (string, string) {
	if !strings.HasPrefix(path, prefix) {
		return "", ""
	}
	pathLen := len(path)
	prefixLen := len(prefix)
	searchStart := prefixLen + 1
	if pathLen <= searchStart {
		return "", ""
	}
	pos := strings.Index(path[searchStart:pathLen], "/")
	if pos == -1 { // we have only user directory...
		return path[searchStart:pathLen], path
	}
	return path[searchStart : pos+searchStart], path[0 : pos+searchStart]
}

func (ic *importContext) emitUserOrServicePrincipalForPath(path, prefix string) {
	userOrSpName, _ := getUserOrSpNameAndDirectory(path, prefix)
	if userOrSpName != "" {
		ic.emitUserOrServicePrincipal(userOrSpName)
	}
}

func (ic *importContext) IsUserOrServicePrincipalDirectory(path, prefix string, strict bool) bool {
	userOrSPName, userDir := getUserOrSpNameAndDirectory(path, prefix)
	if userOrSPName == "" {
		return false
	}
	// strict mode means that it should be exactly user dir, maybe with trailing `/`
	if strict && !(len(path) == len(userDir) || (len(path) == len(userDir)+1 && path[len(path)-1] == '/')) {
		return false
	}
	ic.userOrSpDirectoriesMutex.RLock()
	result, exists := ic.userOrSpDirectories[userDir]
	ic.userOrSpDirectoriesMutex.RUnlock()
	if exists {
		// log.Printf("[DEBUG] Directory %s already checked. Result=%v", userDir, result)
		return result
	}
	var err error
	if common.StringIsUUID(userOrSPName) {
		_, err = ic.findSpnByAppID(userOrSPName, true)
		if err != nil {
			ic.addIgnoredResource(fmt.Sprintf("databricks_service_principal. application_id=%s", userOrSPName))
		}
	} else {
		_, err = ic.findUserByName(strings.ToLower(userOrSPName), true)
		if err != nil {
			ic.addIgnoredResource(fmt.Sprintf("databricks_user. user_name=%s", userOrSPName))
		}
	}
	ic.userOrSpDirectoriesMutex.Lock()
	ic.userOrSpDirectories[userDir] = (err == nil)
	ic.userOrSpDirectoriesMutex.Unlock()
	return err == nil
}

func (ic *importContext) emitGroups(u scim.User) {
	for _, g := range u.Groups {
		if g.Type != "direct" {
			log.Printf("[DEBUG] Skipping non-direct group %s/%s for %s", g.Value, g.Display, u.DisplayName)
			continue
		}
		ic.Emit(&resource{
			Resource: "databricks_group",
			ID:       g.Value,
		})
		id := fmt.Sprintf("%s|%s", g.Value, u.ID)
		ic.Emit(&resource{
			Resource: "databricks_group_member",
			ID:       id,
			Name:     fmt.Sprintf("%s_%s_%s_%s", g.Display, g.Value, u.DisplayName, u.ID),
			Data:     ic.makeGroupMemberData(id, g.Value, u.ID),
		})
	}
}

func (ic *importContext) emitRoles(objType string, id string, roles []scim.ComplexValue) {
	log.Printf("[DEBUG] emitting roles for object type: %s, ID: %s, roles: %v", objType, id, roles)
	for _, role := range roles {
		if role.Type != "direct" {
			continue
		}
		if !ic.accountLevel {
			ic.Emit(&resource{
				Resource: "databricks_instance_profile",
				ID:       role.Value,
			})
		}
		ic.Emit(&resource{
			Resource: fmt.Sprintf("databricks_%s_role", objType),
			ID:       fmt.Sprintf("%s|%s", id, role.Value),
		})
	}
}

func (ic *importContext) cacheGroups() error {
	ic.groupsMutex.Lock()
	defer ic.groupsMutex.Unlock()
	if ic.allGroups == nil {
		log.Printf("[INFO] Caching groups in memory ...")
		var groups *[]iam.Group
		var err error
		err = runWithRetries(func() error {
			var grps []iam.Group
			var err error
			if ic.accountLevel {
				grps, err = ic.accountClient.Groups.ListAll(ic.Context, iam.ListAccountGroupsRequest{
					Attributes: "id",
				})
			} else {
				grps, err = ic.workspaceClient.Groups.ListAll(ic.Context, iam.ListGroupsRequest{
					Attributes: "id",
				})
			}
			if err != nil {
				return err
			}
			groups = &grps
			return nil
		}, "error fetching full list of groups")
		if err != nil {
			log.Printf("[ERROR] can't fetch list of groups. Error: %v", err)
			return err
		}
		api := scim.NewGroupsAPI(ic.Context, ic.Client)
		groupsCount := len(*groups)
		ic.allGroups = make([]scim.Group, 0, groupsCount)
		for i, g := range *groups {
			err = runWithRetries(func() error {
				group, err := api.Read(g.Id, "id,displayName,active,externalId,entitlements,groups,roles,members,meta")
				if err != nil {
					return err
				}
				ic.allGroups = append(ic.allGroups, group)
				return nil
			}, "error reading group with ID "+g.Id)
			if err != nil {
				log.Printf("[ERROR] Error reading group with ID %s: %v", g.Id, err)
				continue
			}
			if (i+1)%10 == 0 {
				log.Printf("[DEBUG] Read %d out of %d groups", i+1, groupsCount)
			}
		}
		log.Printf("[INFO] Cached %d groups", len(ic.allGroups))
	}
	return nil
}

func (ic *importContext) getUsersMapping() {
	ic.allUsersMutex.RLocker().Lock()
	userMapping := ic.allUsersMapping
	ic.allUsersMutex.RLocker().Unlock()
	if userMapping == nil {
		ic.allUsersMutex.Lock()
		defer ic.allUsersMutex.Unlock()
		if ic.allUsersMapping != nil {
			return
		}
		ic.allUsersMapping = make(map[string]string)
		err := runWithRetries(func() error {
			var users []iam.User
			var err error
			if ic.accountLevel {
				users, err = ic.accountClient.Users.ListAll(ic.Context, iam.ListAccountUsersRequest{
					Attributes: "id,userName",
				})
			} else {
				users, err = ic.workspaceClient.Users.ListAll(ic.Context, iam.ListUsersRequest{
					Attributes: "id,userName",
				})
			}
			if err != nil {
				return err
			}
			for _, user := range users {
				ic.allUsersMapping[user.UserName] = user.Id
			}
			log.Printf("[DEBUG] all %d users are copied", len(users))
			return nil
		}, "error fetching full list of users")
		if err != nil {
			log.Panicf("[ERROR] can't fetch list of users after few retries: error=%v", err)
		}
	}
}

func (ic *importContext) findUserByName(name string, fastCheck bool) (u *scim.User, err error) {
	log.Printf("[DEBUG] Looking for user %s", name)
	ic.usersMutex.RLocker().Lock()
	user, exists := ic.allUsers[name]
	ic.usersMutex.RLocker().Unlock()
	if exists {
		if user.UserName == nonExistingUserOrSp {
			log.Printf("[DEBUG] non-existing user %s is found in the cache", name)
			err = fmt.Errorf("user %s is not found", name)
		} else {
			log.Printf("[DEBUG] existing user %s is found in the cache", name)
			u = &user
		}
		return
	}
	ic.getUsersMapping()
	ic.allUsersMutex.RLocker().Lock()
	userId, exists := ic.allUsersMapping[name]
	ic.allUsersMutex.RLocker().Unlock()
	if !exists {
		err = fmt.Errorf("there is no user '%s'", name)
		u = &scim.User{UserName: nonExistingUserOrSp}
	} else {
		if fastCheck {
			return &scim.User{UserName: name}, nil
		}
		a := scim.NewUsersAPI(ic.Context, ic.Client)
		err = runWithRetries(func() error {
			usr, err := a.Read(userId, "id,userName,displayName,active,externalId,entitlements,groups,roles")
			if err != nil {
				return err
			}
			u = &usr
			return nil
		}, fmt.Sprintf("error reading user with name '%s', user ID: %s", name, userId))
		if err != nil {
			log.Printf("[WARN] error reading user with name '%s', user ID: %s", name, userId)
			u = &scim.User{UserName: nonExistingUserOrSp}
		}
	}
	ic.usersMutex.Lock()
	defer ic.usersMutex.Unlock()
	ic.allUsers[name] = *u
	return
}

func (ic *importContext) getSpsMapping() {
	ic.spsMutex.Lock()
	defer ic.spsMutex.Unlock()
	if ic.allSpsMapping == nil {
		ic.allSpsMapping = make(map[string]string)
		err := runWithRetries(func() error {
			var sps []iam.ServicePrincipal
			var err error
			if ic.accountLevel {
				sps, err = ic.accountClient.ServicePrincipals.ListAll(ic.Context, iam.ListAccountServicePrincipalsRequest{
					Attributes: "id,userName",
				})
			} else {
				sps, err = ic.workspaceClient.ServicePrincipals.ListAll(ic.Context, iam.ListServicePrincipalsRequest{
					Attributes: "id,userName",
				})
			}
			if err != nil {
				return err
			}
			for _, sp := range sps {
				ic.allSpsMapping[sp.ApplicationId] = sp.Id
			}
			return nil
		}, "error fetching full list of service principals")
		if err != nil {
			log.Fatalf("[ERROR] can't fetch list of service principals after few retries: error=%v", err)
		}
	}
}

func (ic *importContext) findSpnByAppID(applicationID string, fastCheck bool) (u *scim.User, err error) {
	log.Printf("[DEBUG] Looking for SP %s", applicationID)
	ic.spsMutex.RLocker().Lock()
	sp, exists := ic.allSps[applicationID]
	ic.spsMutex.RLocker().Unlock()
	if exists {
		if sp.ApplicationID == nonExistingUserOrSp {
			log.Printf("[DEBUG] non-existing SP %s is found in the cache", applicationID)
			err = fmt.Errorf("service principal %s is not found", applicationID)
		} else {
			log.Printf("[DEBUG] existing SP %s is found in the cache", applicationID)
			u = &sp
		}
		return
	}
	ic.getSpsMapping()
	ic.spsMutex.RLocker().Lock()
	spId, exists := ic.allSpsMapping[applicationID]
	ic.spsMutex.RLocker().Unlock()
	if !exists {
		err = fmt.Errorf("there is no service principal '%s'", applicationID)
		u = &scim.User{ApplicationID: nonExistingUserOrSp}
	} else {
		if fastCheck {
			return &scim.User{ApplicationID: applicationID}, nil
		}
		a := scim.NewServicePrincipalsAPI(ic.Context, ic.Client)
		err = runWithRetries(func() error {
			usr, err := a.Read(spId, "userName,displayName,active,externalId,entitlements,groups,roles")
			if err != nil {
				return err
			}
			u = &usr
			return nil
		}, fmt.Sprintf("error reading service principal with AppID '%s', SP ID: %s", applicationID, spId))
		if err != nil {
			log.Printf("[WARN] error reading service principal with AppID '%s', SP ID: %s", applicationID, spId)
			u = &scim.User{ApplicationID: nonExistingUserOrSp}
		}
	}
	ic.spsMutex.Lock()
	defer ic.spsMutex.Unlock()
	ic.allSps[applicationID] = *u

	return
}

func (ic *importContext) makeGroupMemberData(id, groupId, memberId string) *schema.ResourceData {
	data := scim.ResourceGroupMember().ToResource().TestResourceData()
	data.MarkNewResource()
	data.SetId(id)
	data.Set("group_id", groupId)
	data.Set("member_id", memberId)
	return data
}
