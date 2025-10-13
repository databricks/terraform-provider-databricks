package access

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

type ipAccessListUpdateRequest struct {
	common.Namespace
	Label       string            `json:"label"`
	ListType    settings.ListType `json:"list_type"`
	IpAddresses []string          `json:"ip_addresses"`
	Enabled     bool              `json:"enabled,omitempty" tf:"default:true"`
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
			var iacl settings.UpdateIpAccessList
			common.DataToStructPointer(d, s, &iacl)
			iacl.IpAccessListId = d.Id()
			return w.IpAccessLists.Update(ctx, iacl)
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
