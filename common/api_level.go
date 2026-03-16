package common

import (
	"context"
	"fmt"

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

// apiLevelContextKey is a context key for passing the API level override to
// request visitors like scimVisitor.
type apiLevelContextKey struct{}

// ContextWithApiLevel returns a new context with the API level set.
func ContextWithApiLevel(ctx context.Context, level string) context.Context {
	return context.WithValue(ctx, apiLevelContextKey{}, level)
}

// ApiLevelFromContext returns the API level from the context, or empty string if not set.
func ApiLevelFromContext(ctx context.Context) string {
	v, _ := ctx.Value(apiLevelContextKey{}).(string)
	return v
}

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

// IsAccountLevel determines whether a resource should use account-level APIs.
// It checks the `api` field first. If set, it takes precedence. Otherwise, it
// falls back to the provider's host type.
func IsAccountLevel(d *schema.ResourceData, c *DatabricksClient) bool {
	return isAccountLevelFromApiField(d, c)
}

// isAccountLevelFromApiField checks the `api` field and falls back to host type.
func isAccountLevelFromApiField(d *schema.ResourceData, c *DatabricksClient) bool {
	if apiLevel, ok := d.GetOk("api"); ok {
		return apiLevel.(string) == ApiLevelAccount
	}
	return c.Config.HostType() == config.AccountHost
}

// ContextWithApiLevelFromData returns a context with the API level set from the
// resource data's `api` field. This is used to pass the API level to request
// visitors (e.g., scimVisitor) through the context.
func ContextWithApiLevelFromData(ctx context.Context, d *schema.ResourceData) context.Context {
	if apiLevel, ok := d.GetOk("api"); ok {
		return ContextWithApiLevel(ctx, apiLevel.(string))
	}
	return ctx
}

// AccountOrWorkspaceRequestWithApiLevel routes the request to account or workspace
// callbacks based on the `api` field in the resource data, falling back to host type.
func (c *DatabricksClient) AccountOrWorkspaceRequestWithApiLevel(
	d *schema.ResourceData,
	accCallback func(*DatabricksClient) error,
	wsCallback func(*DatabricksClient) error,
) error {
	if IsAccountLevel(d, c) {
		return accCallback(c)
	}
	return wsCallback(c)
}

// ValidateApiField validates the `api` field value during plan phase.
// Returns an error if the value is set but invalid.
func ValidateApiField(d *schema.ResourceData) error {
	if apiLevel, ok := d.GetOk("api"); ok {
		level := apiLevel.(string)
		if level != ApiLevelAccount && level != ApiLevelWorkspace {
			return fmt.Errorf("invalid value for `api`: %q, must be %q or %q",
				level, ApiLevelAccount, ApiLevelWorkspace)
		}
	}
	return nil
}
