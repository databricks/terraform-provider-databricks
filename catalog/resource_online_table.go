package catalog

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const onlineTableDefaultProvisionTimeout = 90 * time.Minute

func waitForOnlineTableCreation(w *databricks.WorkspaceClient, ctx context.Context, onlineTableName string) error {
	return retry.RetryContext(ctx, onlineTableDefaultProvisionTimeout, func() *retry.RetryError {
		endpoint, err := w.OnlineTables.GetByName(ctx, onlineTableName)
		if err != nil {
			return retry.NonRetryableError(err)
		}
		if endpoint.Status == nil {
			return retry.RetryableError(fmt.Errorf("online table status is not available yet"))
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

func waitForOnlineTableDeletion(w *databricks.WorkspaceClient, ctx context.Context, onlineTableName string) error {
	return retry.RetryContext(ctx, onlineTableDefaultProvisionTimeout, func() *retry.RetryError {
		_, err := w.OnlineTables.GetByName(ctx, onlineTableName)
		if err == nil {
			return retry.RetryableError(fmt.Errorf("online table %s is still not deleted", onlineTableName))
		}
		if errors.Is(err, apierr.ErrResourceDoesNotExist) || errors.Is(err, apierr.ErrNotFound) {
			return nil
		}
		return retry.NonRetryableError(fmt.Errorf("online table status returned %w", err))
	})
}

func ResourceOnlineTable() common.Resource {
	s := common.StructToSchema(catalog.OnlineTable{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["name"].DiffSuppressFunc = common.EqualFoldDiffSuppress
			common.CustomizeSchemaPath(m, "spec", "source_table_full_name").SetCustomSuppressDiff(common.EqualFoldDiffSuppress)
			common.CustomizeSchemaPath(m, "name").SetRequired().SetForceNew()
			common.CustomizeSchemaPath(m, "status").SetReadOnly()
			common.CustomizeSchemaPath(m, "table_serving_url").SetReadOnly()
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
			var req catalog.CreateOnlineTableRequest
			common.DataToStructPointer(d, s, &req)
			res, err := w.OnlineTables.Create(ctx, req)
			if err != nil {
				return err
			}
			// Note: We should set the id right after creation and before waiting for online table to be available.
			// This is because in case when online table isn't availabe, we still should have that resource in the state.
			d.SetId(res.Name)
			// this should be specified in the API Spec - filed a ticket to add it
			err = waitForOnlineTableCreation(w, ctx, res.Name)
			if err != nil {
				return err
			}
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
			err = w.OnlineTables.DeleteByName(ctx, d.Id())
			if err != nil {
				return err
			}
			return waitForOnlineTableDeletion(w, ctx, d.Id())
		},
		StateUpgraders: []schema.StateUpgrader{},
		Schema:         s,
		SchemaVersion:  0,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(onlineTableDefaultProvisionTimeout),
		},
	}
}
