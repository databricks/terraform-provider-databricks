package catalog

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const onlineTableDefaultProvisionTimeout = 45 * time.Minute

func waitForOnlineTable(w *databricks.WorkspaceClient, ctx context.Context, onlineTableName string) error {
	return retry.RetryContext(ctx, lakehouseMonitorDefaultProvisionTimeout, func() *retry.RetryError {
		endpoint, err := w.OnlineTables.GetByName(ctx, onlineTableName)
		if err != nil {
			return retry.NonRetryableError(err)
		}
		switch endpoint.Status.DetailedState {
		case catalog.OnlineTableStateOnline, catalog.OnlineTableStateOnlineContinuousUpdate,
			catalog.OnlineTableStateOnlineNoPendingUpdate, catalog.OnlineTableStateOnlineTriggeredUpdate:
			return nil

		// does catalog.OnlineTableStateOffline means that it's failed?
		case catalog.OnlineTableStateOfflineFailed, catalog.OnlineTableStateOnlinePipelineFailed:
			return retry.NonRetryableError(fmt.Errorf("online table status returned %s for online table: %s",
				endpoint.Status.DetailedState.String(), onlineTableName))
		}
		return retry.RetryableError(fmt.Errorf("online table %s is still pending", onlineTableName))
	})
}

func ResourceOnlineTable() common.Resource {
	s := common.StructToSchema(catalog.OnlineTable{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			common.CustomizeSchemaPath(m, "name").SetRequired().SetForceNew()
			common.CustomizeSchemaPath(m, "status").SetReadOnly()
			common.CustomizeSchemaPath(m, "spec", "pipeline_id").SetReadOnly()

			runTypes := []string{"spec.0.run_triggered", "spec.0.run_continuously"}
			common.CustomizeSchemaPath(m, "spec", "run_triggered").SetAtLeastOneOf(runTypes).SetSuppressDiff()
			common.CustomizeSchemaPath(m, "spec", "run_continuously").SetAtLeastOneOf(runTypes).SetSuppressDiff()
			return m
		})

	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var req catalog.ViewData
			common.DataToStructPointer(d, s, &req)
			res, err := w.OnlineTables.Create(ctx, req)
			if err != nil {
				return err
			}
			// this should be specified in the API Spec - filed a ticket to add it
			err = waitForOnlineTable(w, ctx, res.Name)
			if err != nil {

				return err
			}
			d.SetId(res.Name)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			table, err := w.OnlineTables.GetByName(ctx, d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(*table, s, d)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			// TODO: talk with the team on how fast it's deleted - should we wait for it to be deleted?
			return w.OnlineTables.DeleteByName(ctx, d.Id())
		},
		StateUpgraders: []schema.StateUpgrader{},
		Schema:         s,
		SchemaVersion:  0,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(onlineTableDefaultProvisionTimeout),
		},
	}
}
