package sql

import (
	"context"
	"log"

	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceAlert() common.Resource {
	s := common.StructToSchema(sql.Alert{}, func(m map[string]*schema.Schema) map[string]*schema.Schema {
		common.CustomizeSchemaPath(m, "condition").SetRequired()
		// TODO: can we automatically generate it from SDK?
		common.CustomizeSchemaPath(m, "condition", "op").SetRequired().SetValidateFunc(validation.StringInSlice([]string{
			"GREATER_THAN", "GREATER_THAN_OR_EQUAL", "LESS_THAN", "LESS_THAN_OR_EQUAL", "EQUAL", "NOT_EQUAL", "IS_NULL"}, true))
		common.CustomizeSchemaPath(m, "condition", "operand").SetRequired()
		common.CustomizeSchemaPath(m, "condition", "operand", "column").SetRequired()
		common.CustomizeSchemaPath(m, "condition", "operand", "column", "name").SetRequired()
		// We may not need it for some conditions
		// common.CustomizeSchemaPath(m, "condition", "threshold").SetRequired()
		// common.CustomizeSchemaPath(m, "condition", "threshold", "value").SetRequired()
		// alof := []string{"string_value", "double_value", "bool_value"}
		// for _, f := range alof {
		// 	common.CustomizeSchemaPath(m, "condition", "threshold", "value", f).SetAtLeastOneOf(alof)
		// }
		common.CustomizeSchemaPath(m, "condition", "op").SetRequired()
		common.CustomizeSchemaPath(m, "id").SetReadOnly()
		common.CustomizeSchemaPath(m, "create_time").SetReadOnly()
		common.CustomizeSchemaPath(m, "lifecycle_state").SetReadOnly()
		common.CustomizeSchemaPath(m, "state").SetReadOnly()
		common.CustomizeSchemaPath(m, "trigger_time").SetReadOnly()
		common.CustomizeSchemaPath(m, "update_time").SetReadOnly()
		common.CustomizeSchemaPath(m, "owner_user_name").SetSuppressDiff()
		common.CustomizeSchemaPath(m, "parent_path").SetSuppressDiff().SetForceNew()
		common.CustomizeSchemaPath(m, "display_name").SetRequired()
		common.CustomizeSchemaPath(m, "query_id").SetRequired()
		return m
	})

	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var a sql.CreateAlertRequestAlert
			common.DataToStructPointer(d, s, &a)
			apiAlert, err := w.Alerts.Create(ctx, sql.CreateAlertRequest{
				Alert: &a,
			})
			if err != nil {
				return err
			}
			d.SetId(apiAlert.Id)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			apiAlert, err := w.Alerts.GetById(ctx, d.Id())
			if err != nil {
				log.Printf("[WARN] error getting alert by ID: %v", err)
				return err
			}
			return common.StructToData(apiAlert, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var a sql.UpdateAlertRequestAlert
			updateMask := "display_name,query_id,seconds_to_retrigger,condition,custom_body,custom_subject"
			if d.HasChange("owner_user_name") {
				updateMask += ",owner_user_name"
			}
			common.DataToStructPointer(d, s, &a)
			_, err = w.Alerts.Update(ctx, sql.UpdateAlertRequest{
				Alert:      &a,
				Id:         d.Id(),
				UpdateMask: updateMask,
			})
			return err
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return w.Alerts.DeleteById(ctx, d.Id())
		},
		Schema: s,
	}
}
