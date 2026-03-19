package common

import (
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const (
	// ApiLevelAccount indicates the resource should use account-level APIs.
	ApiLevelAccount = "account"
	// ApiLevelWorkspace indicates the resource should use workspace-level APIs.
	ApiLevelWorkspace = "workspace"
)

// AddApiField adds the `api` field to a resource schema. This field allows users
// to explicitly specify whether the resource should use account-level or
// workspace-level APIs. When set, it takes precedence over host-based inference.
func AddApiField(s map[string]*schema.Schema) map[string]*schema.Schema {
	s["api"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		ValidateFunc: validation.StringInSlice(
			[]string{ApiLevelAccount, ApiLevelWorkspace}, false,
		),
		Description: "Specifies whether to use account-level or workspace-level API. " +
			"Valid values are `account` and `workspace`. When not set, the API level " +
			"is inferred from the provider host.",
	}
	return s
}

// GetApiLevel returns the value of the `api` field from resource data,
// or empty string if not set.
func GetApiLevel(d *schema.ResourceData) string {
	if v, ok := d.GetOk("api"); ok {
		level := v.(string)
		if level == ApiLevelAccount || level == ApiLevelWorkspace {
			return level
		}
	}
	return ""
}

// IsAccountLevel determines whether a resource should use account-level APIs.
// It checks the `api` field first. If set, it takes precedence. Otherwise, it
// falls back to the provider's host type.
func IsAccountLevel(d *schema.ResourceData, c *DatabricksClient) bool {
	if apiLevel, ok := d.GetOk("api"); ok {
		switch apiLevel.(string) {
		case ApiLevelAccount:
			return true
		case ApiLevelWorkspace:
			return false
		default:
			return c.Config.HostType() == config.AccountHost
		}
	}
	return c.Config.HostType() == config.AccountHost
}
