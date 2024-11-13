package access

import (
	"context"
	"errors"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ResourceIPAccessList manages IP access lists
func ResourceIPAccessList() common.Resource {
	s := common.StructToSchema(settings.CreateIpAccessList{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		// nolint
		s["enabled"] = &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		}
		s["list_type"].ValidateFunc = validation.StringInSlice([]string{"ALLOW", "BLOCK"}, false)
		s["ip_addresses"].Elem = &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.Any(validation.IsIPv4Address, validation.IsCIDR),
		}
		return s
	})
	return common.Resource{
		Schema: s,
		CanSkipReadAfterCreateAndUpdate: func(d *schema.ResourceData) bool {
			return true
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var iacl settings.CreateIpAccessList
			var updateIacl settings.UpdateIpAccessList
			common.DataToStructPointer(d, s, &iacl)
			common.DataToStructPointer(d, s, &updateIacl)
			return c.AccountOrWorkspaceRequest(func(acc *databricks.AccountClient) error {
				status, err := acc.IpAccessLists.Create(ctx, iacl)
				if err != nil {
					return err
				}
				ipAclId := status.IpAccessList.ListId
				// need to wait until the ip access list is available from get
				retry.RetryContext(ctx, 1*time.Minute, func() *retry.RetryError {
					_, err := acc.IpAccessLists.GetByIpAccessListId(ctx, ipAclId)
					var apiErr *apierr.APIError
					if !errors.As(err, &apiErr) {
						return retry.NonRetryableError(err)
					}
					if apiErr.StatusCode == 404 {
						return retry.RetryableError(err)
					} else {
						return retry.NonRetryableError(err)
					}
				})
				//need to enable the IP Access List with update
				if d.Get("enabled").(bool) {
					updateIacl.IpAccessListId = ipAclId
					err = acc.IpAccessLists.Update(ctx, updateIacl)
					if err != nil {
						return err
					}
				}
				d.SetId(ipAclId)
				return nil
			}, func(w *databricks.WorkspaceClient) error {
				status, err := w.IpAccessLists.Create(ctx, iacl)
				if err != nil {
					return err
				}
				d.SetId(status.IpAccessList.ListId)
				return nil
			})
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return c.AccountOrWorkspaceRequest(func(acc *databricks.AccountClient) error {
				status, err := acc.IpAccessLists.GetByIpAccessListId(ctx, d.Id())
				if err != nil {
					return err
				}
				common.StructToData(status.IpAccessList, s, d)
				return nil
			}, func(w *databricks.WorkspaceClient) error {
				status, err := w.IpAccessLists.GetByIpAccessListId(ctx, d.Id())
				if err != nil {
					return err
				}
				common.StructToData(status.IpAccessList, s, d)
				return nil
			})
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var iacl settings.UpdateIpAccessList
			common.DataToStructPointer(d, s, &iacl)
			iacl.IpAccessListId = d.Id()
			return c.AccountOrWorkspaceRequest(func(acc *databricks.AccountClient) error {
				return acc.IpAccessLists.Update(ctx, iacl)
			}, func(w *databricks.WorkspaceClient) error {
				return w.IpAccessLists.Update(ctx, iacl)
			})
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return c.AccountOrWorkspaceRequest(func(acc *databricks.AccountClient) error {
				return acc.IpAccessLists.DeleteByIpAccessListId(ctx, d.Id())
			}, func(w *databricks.WorkspaceClient) error {
				return w.IpAccessLists.DeleteByIpAccessListId(ctx, d.Id())
			})
		},
	}
}
