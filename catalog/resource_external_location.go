package catalog

import (
	"context"
	"slices"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/databricks/terraform-provider-databricks/catalog/bindings"
	"github.com/databricks/terraform-provider-databricks/common"
)

// This structure contains the fields of both catalog.UpdateExternalLocation and catalog.CreateExternalLocation
type ExternalLocationInfo struct {
	catalog.ExternalLocationInfo
	SkipValidation bool `json:"skip_validation,omitempty"`
	common.Namespace
}

// fileEventsDisableRequested reports whether `enable_file_events = false` is set explicitly while
// file events are effectively enabled. This happens when the server enabled file events by default
// at creation: the API then doesn't return `enable_file_events`, so both state and configuration
// hold false and the attribute itself has no diff.
func fileEventsDisableRequested(d interface {
	GetRawConfigAt(cty.Path) (cty.Value, diag.Diagnostics)
	GetChange(string) (any, any)
}) bool {
	v, diags := d.GetRawConfigAt(cty.GetAttrPath("enable_file_events"))
	if diags.HasError() || !v.Type().Equals(cty.Bool) || v.IsNull() || !v.IsKnown() || v.True() {
		return false
	}
	effective, _ := d.GetChange("effective_enable_file_events")
	enabled, _ := effective.(bool)
	return enabled
}

func externalLocationCustomizeDiff(ctx context.Context, d *schema.ResourceDiff, c *common.DatabricksClient) error {
	// Plan an update to disable file events. effective_enable_file_events is computed, so marking it
	// as unknown is allowed, and Update sends the explicit `enable_file_events = false`.
	if d.Id() != "" && fileEventsDisableRequested(d) {
		if err := d.SetNewComputed("effective_enable_file_events"); err != nil {
			return err
		}
	}
	return common.NamespaceCustomizeDiffNoForceNew(ctx, d, c)
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
			for _, key := range []string{"created_at", "created_by", "credential_id", "updated_at", "updated_by", "browse_only", "effective_enable_file_events"} {
				common.CustomizeSchemaPath(m, key).SetReadOnly()
			}
			// Matches the pattern used for `effective_predictive_optimization_flag` and
			// `provisioning_info` on `databricks_catalog`.
			// SetReadOnly() alone causes a perpetual `(known after apply)` plan because the field
			// isn't enabled on the server side yet, and SetReadOnly().SetSuppressDiff() is rejected
			// by SDKv2's InternalValidate (DiffSuppressFunc requires a config source, which a field
			// with Optional=false doesn't have).
			common.CustomizeSchemaPath(m, "effective_file_event_queue").SetComputed().SetSuppressDiff()
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

			common.NamespaceCustomizeSchemaMap(m)
			return m
		})
	return common.Resource{
		Schema:        s,
		CustomizeDiff: externalLocationCustomizeDiff,
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
			// SDK marshals these bools with omitempty; force-send so an explicit
			// false from HCL reaches the server instead of being elided.
			common.SetForceSendFields(&createExternalLocationRequest, d,
				[]string{"read_only", "fallback", "enable_file_events"})
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

			// A requested disable is a change too: it's planned through the computed
			// effective_enable_file_events, which HasChange doesn't report while it's unknown.
			disableFileEvents := fileEventsDisableRequested(d)
			if !d.HasChangeExcept("owner") && !disableFileEvents {
				return nil
			}
			booleanFields := map[string]string{"read_only": "ReadOnly", "fallback": "Fallback", "enable_file_events": "EnableFileEvents"}
			for key, value := range booleanFields {
				if d.HasChange(key) {
					updateExternalLocationRequest.ForceSendFields = append(updateExternalLocationRequest.ForceSendFields, value)
				}
			}
			if disableFileEvents && !slices.Contains(updateExternalLocationRequest.ForceSendFields, "EnableFileEvents") {
				updateExternalLocationRequest.ForceSendFields = append(updateExternalLocationRequest.ForceSendFields, "EnableFileEvents")
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
			// If the external location is isolated, re-bind the current workspace so it's accessible for deletion.
			err = bindings.AddCurrentWorkspaceBindings(ctx, d, w, d.Id(), bindings.BindingsSecurableTypeExternalLocation)
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
