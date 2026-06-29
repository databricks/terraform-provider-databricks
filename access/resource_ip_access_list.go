package access

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

type ipAccessListUpdateRequest struct {
	Label       string            `json:"label"`
	ListType    settings.ListType `json:"list_type"`
	IpAddresses []string          `json:"ip_addresses"`
	Enabled     bool              `json:"enabled,omitempty" tf:"default:true"`
	common.Namespace
}

// updateIPAccessList issues the update (PATCH) for an IP access list. `enabled`
// is force-sent because the SDK marshaler omits zero-valued fields unless they
// are named in ForceSendFields; without it a configured `enabled = false` is
// dropped from the request and the list is left enabled.
func updateIPAccessList(ctx context.Context, w *databricks.WorkspaceClient, d *schema.ResourceData, s map[string]*schema.Schema) error {
	var iacl settings.UpdateIpAccessList
	common.DataToStructPointer(d, s, &iacl)
	iacl.IpAccessListId = d.Id()
	iacl.ForceSendFields = append(iacl.ForceSendFields, "Enabled")
	return w.IpAccessLists.Update(ctx, iacl)
}

// ResourceIPAccessList manages IP access lists
func ResourceIPAccessList() common.Resource {
	s := common.StructToSchema(ipAccessListUpdateRequest{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		// nolint
		s["list_type"].ValidateFunc = validation.StringInSlice([]string{"ALLOW", "BLOCK"}, false)
		s["ip_addresses"].Elem = &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.Any(validation.IsIPv4Address, validation.IsCIDR),
		}
		common.NamespaceCustomizeSchemaMap(s)
		return s
	})
	return common.Resource{
		Schema: s,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, c *common.DatabricksClient) error {
			return common.NamespaceCustomizeDiff(ctx, d, c)
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			var iacl settings.CreateIpAccessList
			common.DataToStructPointer(d, s, &iacl)
			status, err := w.IpAccessLists.Create(ctx, iacl)
			if err != nil {
				return err
			}
			d.SetId(status.IpAccessList.ListId)
			// The create API has no `enabled` field, so a new list is always
			// created enabled. Honor an explicit `enabled = false` with a
			// follow-up update. The list already exists, so the ID is kept and
			// the error is wrapped to make clear it was created but not disabled;
			// a subsequent apply reconciles the state through the update path.
			if !d.Get("enabled").(bool) {
				if err := updateIPAccessList(ctx, w, d, s); err != nil {
					return fmt.Errorf("IP access list %s created but could not be disabled: %w", d.Id(), err)
				}
			}
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			status, err := w.IpAccessLists.GetByIpAccessListId(ctx, d.Id())
			if err != nil {
				return err
			}
			common.StructToData(status.IpAccessList, s, d)
			return nil
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			return updateIPAccessList(ctx, w, d, s)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			return w.IpAccessLists.DeleteByIpAccessListId(ctx, d.Id())
		},
	}
}
