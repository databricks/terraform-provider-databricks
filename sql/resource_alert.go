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

// AlertStruct embeds the SDK Alert with ProviderConfig
type AlertStruct struct {
	sql.Alert
	common.ProviderConfig
}

func (AlertStruct) CustomizeSchema(s *common.CustomizableSchema) *common.CustomizableSchema {
	s.SchemaPath("display_name").SetRequired()
	s.SchemaPath("query_id").SetRequired()
	s.SchemaPath("condition").SetRequired()
	// TODO: can we automatically generate it from SDK? Or should we avoid validation at all?
	s.SchemaPath("condition", "op").SetRequired().SetValidateFunc(validation.StringInSlice([]string{
		"GREATER_THAN", "GREATER_THAN_OR_EQUAL", "LESS_THAN", "LESS_THAN_OR_EQUAL", "EQUAL", "NOT_EQUAL", "IS_NULL"}, true))
	s.SchemaPath("parent_path").SetCustomSuppressDiff(common.WorkspaceOrEmptyPathPrefixDiffSuppress).SetForceNew()
	s.SchemaPath("condition", "operand").SetRequired()
	s.SchemaPath("condition", "operand", "column").SetRequired()
	s.SchemaPath("condition", "operand", "column", "name").SetRequired()
	s.SchemaPath("condition", "empty_result_state").SetValidateFunc(
		validation.StringInSlice([]string{"UNKNOWN", "OK", "TRIGGERED"}, true))
	// We may not need it for some conditions
	// s.SchemaPath("condition", "threshold").SetRequired()
	s.SchemaPath("condition", "threshold", "value").SetRequired()
	alof := []string{
		"condition.0.threshold.0.value.0.string_value",
		"condition.0.threshold.0.value.0.double_value",
		"condition.0.threshold.0.value.0.bool_value",
	}
	for _, f := range alof {
		s.SchemaPath("condition", "threshold", "value",
			strings.TrimPrefix(f, "condition.0.threshold.0.value.0.")).SetExactlyOneOf(alof)
	}
	s.SchemaPath("owner_user_name").SetSuppressDiff()
	s.SchemaPath("notify_on_ok").SetDefault(true)
	s.SchemaPath("id").SetReadOnly()
	s.SchemaPath("create_time").SetReadOnly()
	s.SchemaPath("lifecycle_state").SetReadOnly()
	s.SchemaPath("state").SetReadOnly()
	s.SchemaPath("trigger_time").SetReadOnly()
	s.SchemaPath("update_time").SetReadOnly()

	s.SchemaPath("provider_config").SetOptional()
	s.SchemaPath("provider_config", "workspace_id").SetRequired()

	return s
}

func ResourceAlert() common.Resource {
	s := common.StructToSchema(AlertStruct{}, nil)

	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
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
			w, err := c.WorkspaceClient()
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
			w, err := c.WorkspaceClient()
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
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return w.Alerts.DeleteById(ctx, d.Id())
		},
		Schema: s,
	}
}
