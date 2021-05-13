package identity

import (
	"context"
	"fmt"
	"strings"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)


// DataSourceUser returns information about user specified by user name
func DataSourceUser() *schema.Resource {
	type entity struct {
		UserName             	string   `json:"user_name"`
		Home             	string   `json:"home" tf:"computed"`
		Alphanumeric             	string   `json:"alphanumeric" tf:"computed"`
	}

	s := common.StructToSchema(entity{}, func(
		s map[string]*schema.Schema) map[string]*schema.Schema {
		// nolint once SDKv2 has Diagnostics-returning validators, change
		s["user_name"].ValidateFunc = validation.StringIsNotEmpty
		return s
	})

	return &schema.Resource{
		Schema: s,
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			var this entity
			err := common.DataToStructPointer(d, s, &this)
			if err != nil {
				return diag.FromErr(err)
			}
			usersAPI := NewUsersAPI(ctx, m)
			userList, err := usersAPI.Filter(fmt.Sprintf("userName eq '%s'", this.UserName))
			if err != nil {
				return diag.FromErr(err)
			}
			if len(userList) == 0 {
				return diag.FromErr(fmt.Errorf("cannot find user %s", this.UserName))
			}
			d.Set("user_name", userList[0].UserName)
			d.Set("home", fmt.Sprintf("/Users/%s", userList[0].UserName))
			splits := strings.Split(userList[0].UserName, "@")
			norm := nonAlphanumeric.ReplaceAllLiteralString(splits[0], "_")
			norm = strings.ToLower(norm)
			d.Set("alphanumeric", norm)
			d.SetId(userList[0].ID)
			return nil
		},
	}
}
