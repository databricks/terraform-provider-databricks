package catalog

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type RegisteredModelSchemaStruct struct {
	catalog.CreateRegisteredModelRequest
	common.Namespace
}

func ResourceRegisteredModel() common.Resource {
	s := common.StructToSchema(
		RegisteredModelSchemaStruct{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			caseInsensitiveFields := []string{"name", "catalog_name", "schema_name"}
			for _, field := range caseInsensitiveFields {
				m[field].DiffSuppressFunc = common.EqualFoldDiffSuppress
			}
			m["name"].ForceNew = true
			m["catalog_name"].ForceNew = true
			m["schema_name"].ForceNew = true
			m["storage_location"].ForceNew = true
			m["storage_location"].Computed = true
			m["owner"] = &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			}
			// Mark read-only fields as Computed
			readOnlyFields := []string{"created_at", "created_by", "full_name", "metastore_id", "updated_at", "updated_by"}
			for _, field := range readOnlyFields {
				if m[field] != nil {
					m[field].Computed = true
				}
			}
			// Customize aliases schema
			if aliasesSchema, ok := m["aliases"]; ok {
				aliasesSchema.Optional = true
				aliasesSchema.Computed = true
				if aliasElem, ok := aliasesSchema.Elem.(*schema.Resource); ok {
					for _, f := range []string{"catalog_name", "schema_name", "model_name", "id"} {
						if af, ok := aliasElem.Schema[f]; ok {
							af.Optional = true
							af.Computed = true
						}
					}
				}
			}
			common.NamespaceCustomizeSchemaMap(m)
			return m
		})

	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			var m catalog.CreateRegisteredModelRequest
			common.DataToStructPointer(d, s, &m)
			aliases := m.Aliases
			m.Aliases = nil
			model, err := w.RegisteredModels.Create(ctx, m)
			if err != nil {
				return err
			}
			d.SetId(model.FullName)
			// Don't update owner if it is not provided
			if d.Get("owner") != "" {
				var update catalog.UpdateRegisteredModelRequest
				common.DataToStructPointer(d, s, &update)
				update.FullName = d.Id()
				_, err = w.RegisteredModels.Update(ctx, update)
				if err != nil {
					return err
				}
			}
			return setRegisteredModelAliases(ctx, d.Id(), aliases, w)
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			model, err := w.RegisteredModels.Get(ctx, catalog.GetRegisteredModelRequest{
				FullName:       d.Id(),
				IncludeAliases: true,
			})
			if err != nil {
				return err
			}
			return common.StructToData(*model, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}

			var u catalog.UpdateRegisteredModelRequest
			common.DataToStructPointer(d, s, &u)
			u.FullName = d.Id()

			if d.HasChange("owner") {
				_, err := w.RegisteredModels.Update(ctx, catalog.UpdateRegisteredModelRequest{
					FullName: u.FullName,
					Owner:    u.Owner,
				})
				if err != nil {
					return err
				}
			}

			if d.HasChange("aliases") {
				if err := updateRegisteredModelAliases(ctx, d, w); err != nil {
					return err
				}
			}

			if !d.HasChangeExcept("owner") || !d.HasChange("comment") {
				return nil
			}

			if u.Comment == "" {
				u.ForceSendFields = append(u.ForceSendFields, "Comment")
			}
			u.Owner = ""
			u.Aliases = nil
			_, err = w.RegisteredModels.Update(ctx, u)
			if err != nil {
				if d.HasChange("owner") {
					// Rollback
					old, new := d.GetChange("owner")
					_, rollbackErr := w.RegisteredModels.Update(ctx, catalog.UpdateRegisteredModelRequest{
						FullName: u.FullName,
						Owner:    old.(string),
					})
					if rollbackErr != nil {
						return common.OwnerRollbackError(err, rollbackErr, old.(string), new.(string))
					}
				}
				return err
			}
			_, err = w.RegisteredModels.Update(ctx, u)
			return err
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			return w.RegisteredModels.DeleteByFullName(ctx, d.Id())
		},
		StateUpgraders: []schema.StateUpgrader{},
		Schema:         s,
		SchemaVersion:  0,
	}
}

func setRegisteredModelAliases(ctx context.Context, fullName string, aliases []catalog.RegisteredModelAlias, w *databricks.WorkspaceClient) error {
	for _, alias := range aliases {
		_, err := w.RegisteredModels.SetAlias(ctx, catalog.SetRegisteredModelAliasRequest{
			FullName:   fullName,
			Alias:      alias.AliasName,
			VersionNum: alias.VersionNum,
		})
		if err != nil {
			return fmt.Errorf("failed to set alias %q: %w", alias.AliasName, err)
		}
	}
	return nil
}

func updateRegisteredModelAliases(ctx context.Context, d *schema.ResourceData, w *databricks.WorkspaceClient) error {
	old, new := d.GetChange("aliases")
	oldAliases := old.([]interface{})
	newAliases := new.([]interface{})

	oldMap := make(map[string]int)
	for _, a := range oldAliases {
		alias := a.(map[string]interface{})
		oldMap[alias["alias_name"].(string)] = alias["version_num"].(int)
	}

	newMap := make(map[string]int)
	for _, a := range newAliases {
		alias := a.(map[string]interface{})
		newMap[alias["alias_name"].(string)] = alias["version_num"].(int)
	}

	// Delete aliases that were removed
	for name := range oldMap {
		if _, exists := newMap[name]; !exists {
			err := w.RegisteredModels.DeleteAlias(ctx, catalog.DeleteAliasRequest{
				FullName: d.Id(),
				Alias:    name,
			})
			if err != nil {
				return fmt.Errorf("failed to delete alias %q: %w", name, err)
			}
		}
	}

	// Set aliases that were added or changed
	for name, version := range newMap {
		oldVersion, exists := oldMap[name]
		if !exists || oldVersion != version {
			_, err := w.RegisteredModels.SetAlias(ctx, catalog.SetRegisteredModelAliasRequest{
				FullName:   d.Id(),
				Alias:      name,
				VersionNum: version,
			})
			if err != nil {
				return fmt.Errorf("failed to set alias %q: %w", name, err)
			}
		}
	}

	return nil
}
