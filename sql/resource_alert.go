package sql

import (
	"context"
	"log"
	"strings"

	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

type Alert struct {
	sql.Alert
	common.Namespace
}

func ResourceAlert() common.Resource {
	s := common.StructToSchema(Alert{}, func(m map[string]*schema.Schema) map[string]*schema.Schema {
		common.CustomizeSchemaPath(m, "display_name").SetRequired()
		common.CustomizeSchemaPath(m, "query_id").SetRequired()
		common.CustomizeSchemaPath(m, "condition").SetRequired()
		// TODO: can we automatically generate it from SDK? Or should we avoid validation at all?
		common.CustomizeSchemaPath(m, "condition", "op").SetRequired().SetValidateFunc(validation.StringInSlice([]string{
			"GREATER_THAN", "GREATER_THAN_OR_EQUAL", "LESS_THAN", "LESS_THAN_OR_EQUAL", "EQUAL", "NOT_EQUAL", "IS_NULL"}, true))
		common.CustomizeSchemaPath(m, "parent_path").SetCustomSuppressDiff(common.WorkspaceOrEmptyPathPrefixDiffSuppress).SetForceNew()
		common.CustomizeSchemaPath(m, "condition", "operand").SetRequired()
		common.CustomizeSchemaPath(m, "condition", "operand", "column").SetRequired()
		common.CustomizeSchemaPath(m, "condition", "operand", "column", "name").SetRequired()
		common.CustomizeSchemaPath(m, "condition", "empty_result_state").SetValidateFunc(
			validation.StringInSlice([]string{"UNKNOWN", "OK", "TRIGGERED"}, true))
		// We may not need it for some conditions
		// common.CustomizeSchemaPath(m, "condition", "threshold").SetRequired()
		common.CustomizeSchemaPath(m, "condition", "threshold", "value").SetRequired()
		alof := []string{
			"condition.0.threshold.0.value.0.string_value",
			"condition.0.threshold.0.value.0.double_value",
			"condition.0.threshold.0.value.0.bool_value",
		}
		for _, f := range alof {
			common.CustomizeSchemaPath(m, "condition", "threshold", "value",
				strings.TrimPrefix(f, "condition.0.threshold.0.value.0.")).SetExactlyOneOf(alof)
		}
		common.CustomizeSchemaPath(m, "owner_user_name").SetSuppressDiff()
		common.CustomizeSchemaPath(m, "notify_on_ok").SetDefault(true)
		common.CustomizeSchemaPath(m, "id").SetReadOnly()
		common.CustomizeSchemaPath(m, "create_time").SetReadOnly()
		common.CustomizeSchemaPath(m, "lifecycle_state").SetReadOnly()
		common.CustomizeSchemaPath(m, "state").SetReadOnly()
		common.CustomizeSchemaPath(m, "trigger_time").SetReadOnly()
		common.CustomizeSchemaPath(m, "update_time").SetReadOnly()
		common.NamespaceCustomizeSchemaMap(m)
		return m
	})

	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			var a sql.CreateAlertRequestAlert
			common.DataToStructPointer(d, s, &a)
			apiAlert, err := w.Alerts.Create(ctx, sql.CreateAlertRequest{
				AutoResolveDisplayName: false,
				Alert:                  &a,
				ForceSendFields:        []string{"AutoResolveDisplayName"},
			})
			if err != nil {
				return err
			}
			d.SetId(apiAlert.Id)
			owner := d.Get("owner_user_name").(string)
			if owner != "" {
				_, err = w.Alerts.Update(ctx, sql.UpdateAlertRequest{
					Alert: &sql.UpdateAlertRequestAlert{
						OwnerUserName: owner,
					},
					Id:         apiAlert.Id,
					UpdateMask: "owner_user_name",
				})
			}
			return err
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			apiAlert, err := w.Alerts.GetById(ctx, d.Id())
			if err != nil {
				log.Printf("[WARN] error getting alert by ID: %v", err)
				return err
			}
			parentPath := d.Get("parent_path").(string)
			if parentPath != "" && strings.HasPrefix(apiAlert.ParentPath, "/Workspace") && !strings.HasPrefix(parentPath, "/Workspace") {
				apiAlert.ParentPath = strings.TrimPrefix(parentPath, "/Workspace")
			}
			return common.StructToData(apiAlert, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			var a sql.UpdateAlertRequestAlert
			updateMask := "display_name,query_id,seconds_to_retrigger,condition,custom_body,custom_subject"
			if d.HasChange("owner_user_name") {
				updateMask += ",owner_user_name"
			}
			if d.HasChange("notify_on_ok") {
				updateMask += ",notify_on_ok"
				a.ForceSendFields = append(a.ForceSendFields, "NotifyOnOk")
			}
			common.DataToStructPointer(d, s, &a)
			_, err = w.Alerts.Update(ctx, sql.UpdateAlertRequest{
				Alert:                  &a,
				Id:                     d.Id(),
				UpdateMask:             updateMask,
				AutoResolveDisplayName: false,
				ForceSendFields:        []string{"AutoResolveDisplayName"},
			})
			return err
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			return w.Alerts.DeleteById(ctx, d.Id())
		},
		Schema: s,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff) error {
			return common.NamespaceCustomizeDiff(d)
		},
	}
}
