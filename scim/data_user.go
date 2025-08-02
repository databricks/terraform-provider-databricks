package scim

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func NewUsersCache() *UsersCache {
	return &UsersCache{
		cache:     map[string]User{},
		mutex:     sync.RWMutex{},
		populated: false,
	}
}

type UsersCache struct {
	mutex     sync.RWMutex
	cache     map[string]User
	populated bool
}

var usersCache *UsersCache = NewUsersCache()

func (c *UsersCache) populate(api UsersAPI) error {
	var users UserList
	req := map[string]string{}
	req["excludedAttributes"] = "roles"
	err := api.client.Scim(api.context, http.MethodGet, "/preview/scim/v2/Users", req, &users)
	if err != nil {
		return err
	}
	u := users.Resources
	for _, user := range u {
		tflog.Debug(api.context, fmt.Sprintf("Caching user %s", user.UserName))
		c.cache[strings.ToLower(user.UserName)] = user
	}
	c.populated = true
	return nil
}

func (c *UsersCache) Get(api UsersAPI, userName string) (User, error) {
	// Databricks' search is case-insensitive
	userName = strings.ToLower(userName)

	c.mutex.RLock()
	if c.populated {
		user, ok := c.cache[userName]
		c.mutex.RUnlock()
		tflog.Debug(api.context, fmt.Sprintf("User %s found in cache: %t", userName, ok))
		if ok {
			return user, nil
		}
		return User{}, fmt.Errorf("cannot find user %s", userName)
	}
	c.mutex.RUnlock()

	c.mutex.Lock()
	defer c.mutex.Unlock()

	if !c.populated {
		if err := c.populate(api); err != nil {
			return User{}, err
		}
	}

	tflog.Debug(api.context, fmt.Sprintf("Getting user %s from cache", userName))
	user, ok := c.cache[userName]
	tflog.Debug(api.context, fmt.Sprintf("User %s found in cache: %t", userName, ok))
	if ok {
		return user, nil
	}

	return User{}, fmt.Errorf("cannot find user %s", userName)
}

func getUser(usersAPI UsersAPI, id, name string) (user User, err error) {
	if id != "" {
		return usersAPI.Read(id, "userName,displayName,externalId,applicationId")
	}
	user, err = usersCache.Get(usersAPI, name)
	return
}

// DataSourceUser returns information about user specified by user name
func DataSourceUser() common.Resource {
	return common.Resource{
		Schema: map[string]*schema.Schema{
			"user_name": {
				Type:         schema.TypeString,
				ExactlyOneOf: []string{"user_name", "user_id"},
				Optional:     true,
			},
			"user_id": {
				Type:         schema.TypeString,
				ExactlyOneOf: []string{"user_name", "user_id"},
				Optional:     true,
			},
			"home": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"repos": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"alphanumeric": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"application_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"acl_principal_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"active": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
		Read: func(ctx context.Context, d *schema.ResourceData, m *common.DatabricksClient) error {
			usersAPI := NewUsersAPI(ctx, m)
			user, err := getUser(usersAPI, d.Get("user_id").(string), d.Get("user_name").(string))
			if err != nil {
				return err
			}
			d.Set("user_name", user.UserName)
			d.Set("display_name", user.DisplayName)
			d.Set("home", fmt.Sprintf("/Users/%s", user.UserName))
			d.Set("repos", fmt.Sprintf("/Repos/%s", user.UserName))
			d.Set("acl_principal_id", fmt.Sprintf("users/%s", user.UserName))
			d.Set("external_id", user.ExternalID)
			d.Set("application_id", user.ApplicationID)
			d.Set("active", user.Active)
			splits := strings.Split(user.UserName, "@")
			norm := nonAlphanumeric.ReplaceAllLiteralString(splits[0], "_")
			norm = strings.ToLower(norm)
			d.Set("alphanumeric", norm)
			d.SetId(user.ID)
			return nil
		},
	}
}
