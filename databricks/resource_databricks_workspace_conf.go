package databricks

// Preview feature: https://docs.databricks.com/security/network/ip-access-list.html
// REST API: https://docs.databricks.com/dev-tools/api/latest/ip-access-list.html#operation/create-list

import (
	"strconv"

	"github.com/databrickslabs/databricks-terraform/client/model"
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
			// "sidebar_logo_active": {
			// 	Type:     schema.TypeBool,
			// 	Optional: true,
			// 	Default:  false,
			// },
			// "home_page_welcome_message": {
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// 	Default:  "",
			// },
			// "sidebar_logo_text": {
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// 	Default:  "",
			// },
			// "sidebar_logo_inactive": {
			// 	Type:     schema.TypeBool,
			// 	Optional: true,
			// 	Default:  false,
			// },
			// "home_page_logo": {
			// 	Type:         schema.TypeString,
			// 	Optional:     true,
			// 	Default:      "",
			// 	ValidateFunc: validation.IsURLWithHTTPorHTTPS,
			// },
			// "home_page_logo_width": {
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// 	Default:  "",
			// },
			// "product_name": {
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// 	Default:  "",
			// },
			// "login_logo": {
			// 	Type:         schema.TypeString,
			// 	Optional:     true,
			// 	Default:      "",
			// 	ValidateFunc: validation.IsURLWithHTTPorHTTPS,
			// },
			// "login_logo_width": {
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// 	Default:  "",
			// },
			// "custom_references": {
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// 	Default:  "",
			// },
		},
	}
}

func resourceWorkspaceConfCreateOrUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)

	wsConfMapBuilder := model.WorkspaceConfRequestMapBuilder{}
	wsConfMap := wsConfMapBuilder.Build(
		// model.WithProductName(d.Get("product_name").(string)),
		// model.WithSidebarLogoText(d.Get("sidebar_logo_text").(string)),
		// model.WithLoginLogo(d.Get("login_logo").(string)),
		// model.WithLoginLogoWidth(d.Get("login_logo_width").(string)),
		// model.WithHomePageWelcomeMessage(d.Get("home_page_welcome_message").(string)),
		// model.WithHomePageLogo(d.Get("home_page_logo").(string)),
		// model.WithHomePageLogoWidth(d.Get("home_page_logo_width").(string)),
		// model.WithSidebarLogoActive(strconv.FormatBool(d.Get("sidebar_logo_active").(bool))),
		// model.WithSidebarLogoInactive(strconv.FormatBool(d.Get("sidebar_logo_inactive").(bool))),
		// model.WithCustomReferences(d.Get("custom_references").(string)),
		model.WithEnableIpAccessLists(strconv.FormatBool(d.Get("enable_ip_access_lists").(bool))),
	)
	err := client.WorkspaceConfigurations().Update(wsConfMap)
	if err != nil {
		return err
	}

	d.SetId("workspace_configs")

	return resourceWorkspaceConfRead(d, m)
}

func resourceWorkspaceConfRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)
	resp, err := client.WorkspaceConfigurations().Read(model.AllWorkspaceConfKeys)
	if err != nil {
		return err
	}

	// err = d.Set("product_name", resp.ProductName)
	// if err != nil {
	// 	return err
	// }
	// err = d.Set("sidebar_logo_text", resp.SidebarLogoText)
	// if err != nil {
	// 	return err
	// }
	// err = d.Set("login_logo", resp.LoginLogo)
	// if err != nil {
	// 	return err
	// }
	// err = d.Set("login_logo_width", resp.LoginLogoWidth)
	// if err != nil {
	// 	return err
	// }
	// err = d.Set("home_page_welcome_message", resp.HomePageWelcomeMessage)
	// if err != nil {
	// 	return err
	// }
	// err = d.Set("home_page_logo", resp.HomePageLogo)
	// if err != nil {
	// 	return err
	// }
	// err = d.Set("home_page_logo_width", resp.HomePageLogoWidth)
	// if err != nil {
	// 	return err
	// }
	// err = d.Set("custom_references", resp.CustomReferences)
	// if err != nil {
	// 	return err
	// }

	// val, e2 := strconv.ParseBool(resp.SidebarLogoActive)
	// if e2 != nil {
	// 	val = false
	// }
	// err = d.Set("sidebar_logo_active", val)
	// if err != nil {
	// 	return err
	// }

	// val, e2 = strconv.ParseBool(resp.SidebarLogoInactive)
	// if e2 != nil {
	// 	val = false
	// }
	// err = d.Set("sidebar_logo_inactive", val)
	// if err != nil {
	// 	return err
	// }

	val, e2 := strconv.ParseBool(resp.EnableIpAccessLists)
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

	wsConfMapBuilder := model.WorkspaceConfRequestMapBuilder{}
	wsConfMap := wsConfMapBuilder.Build(
		// model.WithProductName(""),
		// model.WithSidebarLogoText(""),
		// model.WithLoginLogo(""),
		// model.WithLoginLogoWidth(""),
		// model.WithHomePageWelcomeMessage(""),
		// model.WithHomePageLogo(""),
		// model.WithHomePageLogoWidth(""),
		// model.WithSidebarLogoActive("false"),
		// model.WithSidebarLogoInactive("false"),
		// model.WithCustomReferences(""),
		model.WithEnableIpAccessLists("false"),
	)
	return client.WorkspaceConfigurations().Update(wsConfMap)
}
