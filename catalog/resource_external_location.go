package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/catalog/bindings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// This structure contains the fields of both catalog.UpdateExternalLocation and catalog.CreateExternalLocation
type ExternalLocationInfo struct {
	catalog.ExternalLocationInfo
	SkipValidation bool `json:"skip_validation,omitempty"`
}

func ResourceExternalLocation() common.Resource {
	s := common.StructToSchema(ExternalLocationInfo{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["force_destroy"] = &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			}
			m["force_update"] = &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			}
			m["skip_validation"].DiffSuppressFunc = func(k, old, new string, d *schema.ResourceData) bool {
				return old == "false" && new == "true"
			}
			common.CustomizeSchemaPath(m, "url").SetRequired().SetCustomSuppressDiff(ucDirectoryPathSlashOnlySuppressDiff)
			common.CustomizeSchemaPath(m, "name").SetRequired().SetForceNew().SetCustomSuppressDiff(common.EqualFoldDiffSuppress)
			common.CustomizeSchemaPath(m, "credential_name").SetRequired()
			common.CustomizeSchemaPath(m, "isolation_mode").SetComputed()
			common.CustomizeSchemaPath(m, "owner").SetComputed()
			common.CustomizeSchemaPath(m, "metastore_id").SetComputed()
			for _, key := range []string{"created_at", "created_by", "credential_id", "updated_at", "updated_by", "browse_only"} {
				common.CustomizeSchemaPath(m, key).SetReadOnly()
			}
			// customize file event queue
			supportedQueues := []string{"managed_pubsub", "managed_aqs", "managed_sqs", "provided_pubsub", "provided_aqs", "provided_sqs"}
			for _, key := range supportedQueues {
				common.CustomizeSchemaPath(m, "file_event_queue", key, "managed_resource_id").SetReadOnly()
				conflicts := make([]string, 0, len(supportedQueues)-1)
				for _, otherKey := range supportedQueues {
					if key != otherKey {
						conflicts = append(conflicts, "file_event_queue.0."+otherKey)
					}
				}
				common.CustomizeSchemaPath(m, "file_event_queue", key).SetConflictsWith(conflicts)
			}
			common.CustomizeSchemaPath(m, "file_event_queue", "provided_pubsub", "subscription_name").SetRequired()
			common.CustomizeSchemaPath(m, "file_event_queue", "provided_aqs", "queue_url").SetRequired()
			common.CustomizeSchemaPath(m, "file_event_queue", "provided_sqs", "queue_url").SetRequired()
			common.CustomizeSchemaPath(m, "file_event_queue", "managed_aqs", "resource_group").SetRequired()
			common.CustomizeSchemaPath(m, "file_event_queue", "managed_aqs", "subscription_id").SetRequired()
			common.CustomizeSchemaPath(m, "file_event_queue").SetMaxItems(1)

			return m
		})
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			err = validateMetastoreId(ctx, w, d.Get("metastore_id").(string))
			if err != nil {
				return err
			}
			var createExternalLocationRequest catalog.CreateExternalLocation
			common.DataToStructPointer(d, s, &createExternalLocationRequest)
			el, err := w.ExternalLocations.Create(ctx, createExternalLocationRequest)
			if err != nil {
				return err
			}
			d.SetId(el.Name)

			// Update owner or isolation mode if it is provided
			if !updateRequired(d, []string{"owner", "isolation_mode"}) {
				return nil
			}

			var updateExternalLocationRequest catalog.UpdateExternalLocation
			common.DataToStructPointer(d, s, &updateExternalLocationRequest)
			updateExternalLocationRequest.Name = d.Id()
			_, err = w.ExternalLocations.Update(ctx, updateExternalLocationRequest)
			if err != nil {
				return err
			}

			// Bind the current workspace if the external location is isolated, otherwise the read will fail
			return bindings.AddCurrentWorkspaceBindings(ctx, d, w, el.Name, bindings.BindingsSecurableTypeExternalLocation)
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			el, err := w.ExternalLocations.GetByName(ctx, d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(el, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			force := d.Get("force_update").(bool)
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			err = validateMetastoreId(ctx, w, d.Get("metastore_id").(string))
			if err != nil {
				return err
			}
			var updateExternalLocationRequest catalog.UpdateExternalLocation
			common.DataToStructPointer(d, s, &updateExternalLocationRequest)
			updateExternalLocationRequest.Name = d.Id()
			updateExternalLocationRequest.Force = force
			if d.HasChange("owner") {
				_, err = w.ExternalLocations.Update(ctx, catalog.UpdateExternalLocation{
					Name:  updateExternalLocationRequest.Name,
					Owner: updateExternalLocationRequest.Owner,
				})
				if err != nil {
					return err
				}
			}

			if !d.HasChangeExcept("owner") {
				return nil
			}
			booleanFields := map[string]string{"read_only": "ReadOnly", "fallback": "Fallback", "enable_file_events": "EnableFileEvents"}
			for key, value := range booleanFields {
				if d.HasChange(key) {
					updateExternalLocationRequest.ForceSendFields = append(updateExternalLocationRequest.ForceSendFields, value)
				}
			}
			// ugly hack until API is fixed
			if updateExternalLocationRequest.FileEventQueue != nil {
				if updateExternalLocationRequest.FileEventQueue.ManagedAqs != nil {
					updateExternalLocationRequest.FileEventQueue.ManagedAqs.ManagedResourceId = ""
				}
				if updateExternalLocationRequest.FileEventQueue.ManagedPubsub != nil {
					updateExternalLocationRequest.FileEventQueue.ManagedPubsub.ManagedResourceId = ""
				}
				if updateExternalLocationRequest.FileEventQueue.ManagedSqs != nil {
					updateExternalLocationRequest.FileEventQueue.ManagedSqs.ManagedResourceId = ""
				}
				if updateExternalLocationRequest.FileEventQueue.ProvidedPubsub != nil {
					updateExternalLocationRequest.FileEventQueue.ProvidedPubsub.ManagedResourceId = ""
				}
				if updateExternalLocationRequest.FileEventQueue.ProvidedAqs != nil {
					updateExternalLocationRequest.FileEventQueue.ProvidedAqs.ManagedResourceId = ""
				}
				if updateExternalLocationRequest.FileEventQueue.ProvidedSqs != nil {
					updateExternalLocationRequest.FileEventQueue.ProvidedSqs.ManagedResourceId = ""
				}
			}
			// if file events are disabled we shouldn't send FileEventQueue
			if !updateExternalLocationRequest.EnableFileEvents {
				updateExternalLocationRequest.FileEventQueue = nil
			}
			updateExternalLocationRequest.Owner = ""
			_, err = w.ExternalLocations.Update(ctx, updateExternalLocationRequest)
			if err != nil {
				if d.HasChange("owner") {
					// Rollback
					old, new := d.GetChange("owner")
					_, rollbackErr := w.ExternalLocations.Update(ctx, catalog.UpdateExternalLocation{
						Name:  updateExternalLocationRequest.Name,
						Owner: old.(string),
					})
					if rollbackErr != nil {
						return common.OwnerRollbackError(err, rollbackErr, old.(string), new.(string))
					}
				}
				return err
			}
			// Bind the current workspace if the external location is isolated, otherwise the read will fail
			return bindings.AddCurrentWorkspaceBindings(ctx, d, w, updateExternalLocationRequest.Name, bindings.BindingsSecurableTypeExternalLocation)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			force := d.Get("force_destroy").(bool)
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			err = validateMetastoreId(ctx, w, d.Get("metastore_id").(string))
			if err != nil {
				return err
			}
			return w.ExternalLocations.Delete(ctx, catalog.DeleteExternalLocationRequest{
				Name:  d.Id(),
				Force: force,
			})
		},
	}
}
