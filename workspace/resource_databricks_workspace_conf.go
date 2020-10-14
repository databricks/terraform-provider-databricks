package workspace

// Preview feature: https://docs.databricks.com/security/network/ip-access-list.html
// REST API: https://docs.databricks.com/dev-tools/api/latest/ip-access-list.html#operation/create-list

import (
	"strconv"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceWorkspaceConf() *schema.Resource {
	return &schema.Resource{
		Create: resourceWorkspaceConfCreateOrUpdate,
		Read:   resourceWorkspaceConfRead,
		Update: resourceWorkspaceConfCreateOrUpdate,
		Delete: resourceWorkspaceConfDelete,

		Schema: map[string]*schema.Schema{
			"enable_ip_access_lists": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

var (
	// Keeping the model simple (non-existent) with the client operating on map[string]string
	// This is a map from tf name to json name
	tfNameToJsonName = map[string]string{
		"enable_ip_access_lists": "enableIpAccessLists",
	}
)

func resourceWorkspaceConfCreateOrUpdate(d *schema.ResourceData, m interface{}) error {
	wsConfMap := map[string]string{
		tfNameToJsonName["enable_ip_access_lists"]: strconv.FormatBool(d.Get("enable_ip_access_lists").(bool)), // fmt.sprintf TODO
	}
	err := NewWorkspaceConfAPI(m).Update(wsConfMap)
	// 404 check not needed as this only return 400, 401, and 500 on error
	if err != nil {
		return err
	}

	d.SetId("workspace_configs") // Setting to a single id--expected only one config object

	return resourceWorkspaceConfRead(d, m)
}

func resourceWorkspaceConfRead(d *schema.ResourceData, m interface{}) (err error) {
	resp, err := NewWorkspaceConfAPI(m).Read(tfNameToJsonName["enable_ip_access_lists"])
	// 404 check not required as the service only return 400 errors
	if e, ok := err.(common.APIError); ok && e.IsMissing() {
		d.SetId("")
		return nil
	}
	if err != nil {
		// check 404 (missing) and set id to empty string for tf
		return
	}

	val, e2 := strconv.ParseBool(resp[tfNameToJsonName["enable_ip_access_lists"]])
	if e2 != nil {
		val = false
	}
	err = d.Set("enable_ip_access_lists", val)
	if err != nil {
		return
	}

	return
}

func resourceWorkspaceConfDelete(d *schema.ResourceData, m interface{}) (_ error) {
	// For IP Access Lists, you can't set to null or "" once it is set.  Only true/false allowed
	wsConfMap := map[string]string{
		tfNameToJsonName["enable_ip_access_lists"]: "false",
	}
	return NewWorkspaceConfAPI(m).Update(wsConfMap)
}
