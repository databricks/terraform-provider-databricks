package access

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ResourceIPAccessList manages IP access lists
func ResourceIPAccessList() *schema.Resource {
	s := common.StructToSchema(struct {
		Label       string   `json:"label"`
		ListType    string   `json:"list_type"`
		IPAddresses []string `json:"ip_addresses"`
		Enabled     bool     `json:"enabled,omitempty" tf:"default:true"`
	}{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		// nolint
		s["list_type"].ValidateFunc = validation.StringInSlice([]string{"ALLOW", "BLOCK"}, false)
		s["ip_addresses"].Elem = &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.Any(validation.IsIPv4Address, validation.IsCIDR),
		}
		return s
	})
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
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
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			status, err := w.IpAccessLists.GetByIpAccessListId(ctx, d.Id())
			if err != nil {
				return err
			}
			common.StructToData(status, s, d)
			return nil
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var iacl settings.UpdateIpAccessList
			common.DataToStructPointer(d, s, &iacl)
			return w.IpAccessLists.Update(ctx, iacl)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return w.IpAccessLists.DeleteByIpAccessListId(ctx, d.Id())
		},
	}.ToResource()
}
