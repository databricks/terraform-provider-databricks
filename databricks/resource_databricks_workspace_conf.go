package databricks

// Preview feature: https://docs.databricks.com/security/network/ip-access-list.html
// REST API: https://docs.databricks.com/dev-tools/api/latest/ip-access-list.html#operation/create-list

import (
	"strconv"

	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceWorkspaceConf() *schema.Resource {
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
	client := m.(*service.DBApiClient)

	wsConfMap := map[string]string{
		tfNameToJsonName["enable_ip_access_lists"]: strconv.FormatBool(d.Get("enable_ip_access_lists").(bool)),
	}
	err := client.WorkspaceConfigurations().Update(wsConfMap)
	if err != nil {
		return err
	}

	d.SetId("workspace_configs")

	return resourceWorkspaceConfRead(d, m)
}

func resourceWorkspaceConfRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)
	resp, err := client.WorkspaceConfigurations().Read(tfNameToJsonName["enable_ip_access_lists"])
	if err != nil {
		return err
	}

	val, e2 := strconv.ParseBool(resp[tfNameToJsonName["enable_ip_access_lists"]])
	if e2 != nil {
		val = false
	}
	err = d.Set("enable_ip_access_lists", val)
	if err != nil {
		return err
	}

	return err
}

func resourceWorkspaceConfDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)

	// For IP Access Lists, you can't set to null or "" once it is set.  Only true/false allowed
	wsConfMap := map[string]string{
		tfNameToJsonName["enable_ip_access_lists"]: "false",
	}
	return client.WorkspaceConfigurations().Update(wsConfMap)
}
