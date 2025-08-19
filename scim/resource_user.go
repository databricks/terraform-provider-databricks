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

type userResource struct {
	entitlements
	UserName              string `json:"user_name" tf:"force_new"`
	DisplayName           string `json:"display_name,omitempty" tf:"computed"`
	Active                bool   `json:"active,omitempty"`
	ExternalID            string `json:"external_id,omitempty" tf:"suppress_diff"`
	Force                 bool   `json:"force,omitempty"`
	Home                  string `json:"home,omitempty" tf:"computed"`
	Repos                 string `json:"repos,omitempty" tf:"computed"`
	ForceDeleteRepos      bool   `json:"force_delete_repos,omitempty"`
	ForceDeleteHomeDir    bool   `json:"force_delete_home_dir,omitempty"`
	DisableAsUserDeletion bool   `json:"disable_as_user_deletion,omitempty"`
	AclPrincipalID        string `json:"acl_principal_id,omitempty" tf:"computed"`
}

func getUserHomeDir(userName string) string {
	return fmt.Sprintf("/Users/%s", userName)
}

func getUserReposDir(userName string) string {
	return fmt.Sprintf("/Repos/%s", userName)
}

// ResourceUser manages users within workspace
func ResourceUser() common.Resource {
	userSchema := common.StructToSchema(userResource{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["user_name"].DiffSuppressFunc = common.EqualFoldDiffSuppress
			m["active"].Default = true
			return customizeEntitlementsSchema(m)
		})
	return common.Resource{
		Schema: userSchema,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var userResource userResource
			common.DataToStructPointer(d, userSchema, &userResource)
			u := User{
				UserName:     userResource.UserName,
				DisplayName:  userResource.DisplayName,
				Active:       userResource.Active,
				Entitlements: userResource.entitlements.toComplexValueList(),
				ExternalID:   userResource.ExternalID,
			}
			usersAPI := NewUsersAPI(ctx, c)
			user, err := usersAPI.Create(u)
			if err != nil {
				if !userResource.Force {
					return err
				}
				return createForceOverridesManuallyAddedUser(err, d, usersAPI, u)
			}
			d.SetId(user.ID)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			user, err := NewUsersAPI(ctx, c).Read(d.Id(), userAttributes)
			if err != nil {
				return err
			}
			userResource := userResource{
				DisplayName:    user.DisplayName,
				Active:         user.Active,
				ExternalID:     user.ExternalID,
				Home:           getUserHomeDir(user.UserName),
				Repos:          getUserReposDir(user.UserName),
				UserName:       user.UserName,
				AclPrincipalID: fmt.Sprintf("users/%s", user.UserName),
				entitlements:   newEntitlements(ctx, user.Entitlements),
			}
			return common.StructToData(userResource, userSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var userResource userResource
			common.DataToStructPointer(d, userSchema, &userResource)
			u := User{
				UserName:     userResource.UserName,
				DisplayName:  userResource.DisplayName,
				Active:       userResource.Active,
				Entitlements: userResource.entitlements.toComplexValueList(),
				ExternalID:   userResource.ExternalID,
			}
			return NewUsersAPI(ctx, c).Update(d.Id(), u)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var userResource userResource
			common.DataToStructPointer(d, userSchema, &userResource)
			user := NewUsersAPI(ctx, c)
			var err error = nil
			isAccount := c.Config.IsAccountClient() && c.Config.AccountID != ""
			isForceDeleteRepos := userResource.ForceDeleteRepos
			isForceDeleteHomeDir := userResource.ForceDeleteHomeDir
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
			if !isAccount && !isDisable {
				if isForceDeleteRepos {
					err = workspace.NewNotebooksAPI(ctx, c).Delete(getUserReposDir(userResource.UserName), true)
					if err != nil && !apierr.IsMissing(err) {
						return fmt.Errorf("force_delete_repos: %s", err.Error())
					}
				}
				if isForceDeleteHomeDir {
					err = workspace.NewNotebooksAPI(ctx, c).Delete(getUserHomeDir(userResource.UserName), true)
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
