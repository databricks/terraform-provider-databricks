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

func isRetriable(err error) *retry.RetryError {
	var apiErr *apierr.APIError
	if !errors.As(err, &apiErr) {
		return retry.NonRetryableError(err)
	}
	if apiErr.StatusCode == 404 {
		return retry.RetryableError(err)
	} else {
		return retry.NonRetryableError(err)
	}
}

var ipAclTimeout = 10 * time.Minute

// ResourceIPAccessList manages IP access lists
func ResourceIPAccessList() common.Resource {
	s := common.StructToSchema(settings.IpAccessListInfo{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		for _, required := range []string{"list_type", "label"} {
			common.CustomizeSchemaPath(s, required).SetRequired()
		}

		for _, computed := range []string{"address_count", "created_at", "created_by", "list_id", "updated_at", "updated_by"} {
			common.CustomizeSchemaPath(s, computed).SetComputed()
		}

		common.CustomizeSchemaPath(s, "enabled").SetDefault(true)

		common.CustomizeSchemaPath(s, "list_type").SetValidateFunc(validation.StringInSlice([]string{"ALLOW", "BLOCK"}, false))
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
				retry.RetryContext(ctx, ipAclTimeout, func() *retry.RetryError {
					_, err := acc.IpAccessLists.GetByIpAccessListId(ctx, ipAclId)
					return isRetriable(err)
				})
				//need to enable the IP Access List with update, retry if 404 is returned due to eventual consistency
				if d.Get("enabled").(bool) {
					updateIacl.IpAccessListId = ipAclId
					retry.RetryContext(ctx, ipAclTimeout, func() *retry.RetryError {
						err = acc.IpAccessLists.Update(ctx, updateIacl)
						return isRetriable(err)
					})
				}
				d.SetId(ipAclId)
				return common.StructToData(status.IpAccessList, s, d)
			}, func(w *databricks.WorkspaceClient) error {
				status, err := w.IpAccessLists.Create(ctx, iacl)
				if err != nil {
					return err
				}
				d.SetId(status.IpAccessList.ListId)
				return common.StructToData(status.IpAccessList, s, d)
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
