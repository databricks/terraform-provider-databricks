package scim

import (
	"context"
	"fmt"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/workspace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// userListAttrs are the SCIM attributes fetched when bulk-listing users for the read cache.
const userListAttrs = "id,userName,displayName,active,externalId,entitlements"

// globalUsersListCache caches a ListAll result per (host, apiLevel, accountID) so
// that N concurrent reads for N distinct databricks_user resources only issue one
// SCIM list call instead of N individual GET-by-ID calls.
//
// The cache key includes accountID so that two provider instances targeting the
// same API host but different Databricks accounts never share cached users.
var globalUsersListCache = newUsersListCache()

func newUsersListCache() *common.KeyedCache[string, map[string]User] {
	return common.NewKeyedCache[string, map[string]User]()
}

func usersCacheKey(api UsersAPI) string {
	return api.client.Config.Host + "|" + api.ApiLevel + "|" + api.client.Config.AccountID
}

func usersListCacheLookup(api UsersAPI, userID string) (User, error) {
	byID, err := globalUsersListCache.Get(usersCacheKey(api), func() (map[string]User, error) {
		users, err := api.ListAll(userListAttrs)
		if err != nil {
			return nil, err
		}
		m := make(map[string]User, len(users))
		for _, u := range users {
			m[u.ID] = u
		}
		return m, nil
	})
	if err != nil {
		return User{}, err
	}
	if u, ok := byID[userID]; ok {
		return u, nil
	}
	// Cache populated but user absent (e.g. created concurrently); fall back to a direct read.
	return api.Read(userID, userAttributes)
}

func userExistsErrorMessage(userName string, isAccount bool) string {
	if isAccount {
		return strings.ToLower(fmt.Sprintf("User with email %s already exists in this account", userName))
	} else {
		return strings.ToLower(fmt.Sprintf("User with username %s already exists.", userName))
	}
}

const (
	userAttributes = "userName,displayName,active,externalId,entitlements"
)

func setCommonUserFields(d *schema.ResourceData, user User, username string) {
	d.Set("display_name", user.DisplayName)
	d.Set("active", user.Active)
	d.Set("external_id", user.ExternalID)
	d.Set("home", fmt.Sprintf("/Users/%s", username))
	d.Set("repos", fmt.Sprintf("/Repos/%s", username))
}

// ResourceUser manages users within workspace
func ResourceUser() common.Resource {
	type entity struct {
		UserName    string `json:"user_name" tf:"force_new"`
		DisplayName string `json:"display_name,omitempty" tf:"computed"`
		Active      bool   `json:"active,omitempty"`
		ExternalID  string `json:"external_id,omitempty" tf:"suppress_diff"`
	}
	userSchema := common.StructToSchema(entity{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			addEntitlementsToSchema(m)
			common.AddApiField(m)
			m["user_name"].DiffSuppressFunc = common.EqualFoldDiffSuppress
			m["active"].Default = true
			m["force"] = &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			}
			m["home"] = &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			}
			m["repos"] = &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			}
			m["force_delete_repos"] = &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			}
			m["force_delete_home_dir"] = &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			}
			m["disable_as_user_deletion"] = &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			}
			m["acl_principal_id"] = &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			}
			return m
		})
	common.AddNamespaceInSchema(userSchema)
	common.NamespaceCustomizeSchemaMap(userSchema)
	scimUserFromData := func(d *schema.ResourceData) (user User, err error) {
		var u entity
		common.DataToStructPointer(d, userSchema, &u)
		return User{
			UserName:     u.UserName,
			DisplayName:  u.DisplayName,
			Active:       u.Active,
			Entitlements: readEntitlementsFromData(d),
			ExternalID:   u.ExternalID,
		}, nil
	}
	return common.Resource{
		IsDual: true,
		Schema: userSchema,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, c *common.DatabricksClient) error {
			return common.CustomizeDiffDualResources(ctx, d, c)
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			c, err := c.DatabricksClientForDualResource(ctx, d)
			if err != nil {
				return err
			}
			usersAPI := NewUsersAPI(ctx, c, common.GetApiLevel(d))
			defer globalUsersListCache.Invalidate(usersCacheKey(usersAPI))
			u, err := scimUserFromData(d)
			if err != nil {
				return err
			}
			user, err := usersAPI.Create(u)
			if err != nil {
				return createForceOverridesManuallyAddedUser(err, d, usersAPI, u)
			}
			d.SetId(user.ID)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			c, err := c.DatabricksClientForDualResource(ctx, d)
			if err != nil {
				return err
			}
			usersAPI := NewUsersAPI(ctx, c, common.GetApiLevel(d))
			user, err := usersListCacheLookup(usersAPI, d.Id())
			if err != nil {
				return err
			}
			setCommonUserFields(d, user, user.UserName)
			d.Set("user_name", user.UserName)
			d.Set("acl_principal_id", fmt.Sprintf("users/%s", user.UserName))
			return user.Entitlements.readIntoData(d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			c, err := c.DatabricksClientForDualResource(ctx, d)
			if err != nil {
				return err
			}
			usersAPI := NewUsersAPI(ctx, c, common.GetApiLevel(d))
			defer globalUsersListCache.Invalidate(usersCacheKey(usersAPI))
			u, err := scimUserFromData(d)
			if err != nil {
				return err
			}
			return usersAPI.Update(d.Id(), u)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			c, err := c.DatabricksClientForDualResource(ctx, d)
			if err != nil {
				return err
			}
			usersAPIForDelete := NewUsersAPI(ctx, c, common.GetApiLevel(d))
			defer globalUsersListCache.Invalidate(usersCacheKey(usersAPIForDelete))
			user := usersAPIForDelete
			userName := d.Get("user_name").(string)
			isAccount := common.IsAccountLevel(d, c)
			isForceDeleteRepos := d.Get("force_delete_repos").(bool)
			isForceDeleteHomeDir := d.Get("force_delete_home_dir").(bool)
			// Determine if disable or delete
			var isDisable bool
			if isDisableP, exists := d.GetOkExists("disable_as_user_deletion"); exists {
				isDisable = isDisableP.(bool)
			} else {
				// Default is true for Account SCIM, false otherwise
				isDisable = isAccount
			}
			// Validate input
			if !isAccount && isDisable && isForceDeleteRepos {
				return fmt.Errorf("force_delete_repos: cannot force delete if disable_as_user_deletion is set")
			}
			if !isAccount && isDisable && isForceDeleteHomeDir {
				return fmt.Errorf("force_delete_home_dir: cannot force delete if disable_as_user_deletion is set")
			}
			// Disable or delete
			if isDisable {
				r := PatchRequestWithValue("replace", "active", "false")
				err = user.Patch(d.Id(), r)
			} else {
				err = user.Delete(d.Id())
			}
			if err != nil {
				return err
			}
			// Handle force delete flags
			if !isAccount && !isDisable && err == nil {
				if isForceDeleteRepos {
					err = workspace.NewNotebooksAPI(ctx, c).Delete(fmt.Sprintf("/Repos/%v", userName), true)
					if err != nil && !apierr.IsMissing(err) {
						return fmt.Errorf("force_delete_repos: %s", err.Error())
					}
				}
				if isForceDeleteHomeDir {
					err = workspace.NewNotebooksAPI(ctx, c).Delete(fmt.Sprintf("/Users/%v", userName), true)
					if err != nil && !apierr.IsMissing(err) {
						return fmt.Errorf("force_delete_home_dir: %s", err.Error())
					}
				}
			}
			return nil
		},
	}
}

func createForceOverridesManuallyAddedUser(err error, d *schema.ResourceData, usersAPI UsersAPI, u User) error {
	forceCreate := d.Get("force").(bool)
	if !forceCreate {
		return err
	}
	// corner-case for overriding manually provisioned users
	userName := strings.ReplaceAll(u.UserName, "'", "")
	errStr := strings.ToLower(err.Error())
	if (!strings.HasPrefix(errStr, userExistsErrorMessage(userName, false))) &&
		(!strings.HasPrefix(errStr, userExistsErrorMessage(userName, true))) {
		return err
	}
	userList, err := usersAPI.Filter(fmt.Sprintf(`userName eq "%s"`, userName), true)
	if err != nil {
		return err
	}
	if len(userList) == 0 {
		return fmt.Errorf("cannot find %s for force import", userName)
	}
	user := userList[0]
	d.SetId(user.ID)
	return usersAPI.Update(d.Id(), u)
}
