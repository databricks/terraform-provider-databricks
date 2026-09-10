package catalog

import (
	"context"
	"log"
	"strings"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/exp/slices"
)

var sensitiveOptions = []string{"user", "password", "personalAccessToken", "access_token", "client_secret",
	"pem_private_key", "OAuthPvtKey", "GoogleServiceAccountKeyJson", "bearer_token"}

var computedOptions = []string{"pem_private_key_expiration_epoch_sec", "access_token_expiration"}

func suppressComputedFields(k, old, new string, d *schema.ResourceData) bool {
	for _, option := range computedOptions {
		if k == "options."+option {
			log.Printf("[INFO] Suppressing diff on %s", k)
			return true
		}
	}
	return false
}

// schemaParentFromFullName rebuilds a schema-level connection's parent
// ("schemas/{catalog}.{schema}") from its 3-part full_name. A metastore-level
// (L1) connection has a 1-part full_name and therefore no parent.
func schemaParentFromFullName(fullName string) string {
	parts := strings.Split(fullName, ".")
	if len(parts) == 3 {
		return "schemas/" + parts[0] + "." + parts[1]
	}
	return ""
}

type ConnectionSchemaStruct struct {
	catalog.ConnectionInfo
	// Parent schema for schema-level (L3) connections, in format
	// "schemas/{catalog}.{schema}". Absent for metastore-level (L1) connections.
	Parent string `json:"parent,omitempty"`
	common.Namespace
}

func ResourceConnection() common.Resource {
	s := common.StructToSchema(ConnectionSchemaStruct{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			for _, v := range []string{"url", "metastore_id", "credential_type", "connection_id",
				"created_at", "created_by", "full_name", "provisioning_info", "securable_type", "updated_at", "updated_by"} {
				common.CustomizeSchemaPath(m, v).SetReadOnly()
			}
			for _, v := range []string{"owner", "read_only"} {
				common.CustomizeSchemaPath(m, v).SetComputed()
			}
			for _, v := range []string{"read_only", "properties", "comment", "connection_type", "parent"} {
				common.CustomizeSchemaPath(m, v).SetForceNew()
			}
			common.CustomizeSchemaPath(m, "options").SetSensitive().SetCustomSuppressDiff(suppressComputedFields)
			common.CustomizeSchemaPath(m, "name").SetCustomSuppressDiff(common.EqualFoldDiffSuppress)
			common.NamespaceCustomizeSchemaMap(m)
			return m
		})
	pi := common.NewPairID("metastore_id", "full_name").Schema(
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			return s
		})
	return common.Resource{
		Schema:        s,
		CustomizeDiff: common.NamespaceCustomizeDiffNoForceNew,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			err = validateMetastoreId(ctx, w, d.Get("metastore_id").(string))
			if err != nil {
				return err
			}
			var createConnectionRequest catalog.CreateConnection
			common.DataToStructPointer(d, s, &createConnectionRequest)
			conn, err := w.Connections.Create(ctx, createConnectionRequest)
			if err != nil {
				return err
			}
			// Update owner if it is provided
			if d.Get("owner") != "" {
				var updateConnectionRequest catalog.UpdateConnection
				common.DataToStructPointer(d, s, &updateConnectionRequest)
				updateConnectionRequest.Name = createConnectionRequest.Name
				conn, err = w.Connections.Update(ctx, updateConnectionRequest)
				if err != nil {
					return err
				}
			}
			d.Set("metastore_id", conn.MetastoreId)
			// Address connections by full_name: for L3 it is catalog.schema.name, and for L1 it
			// equals the name, so this preserves the existing metastore_id|name id for L1 (no migration).
			d.SetId(conn.MetastoreId + "|" + conn.FullName)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			_, connName, err := pi.Unpack(d)
			if err != nil {
				return err
			}
			conn, err := w.Connections.GetByName(ctx, connName)
			if err != nil {
				return err
			}
			// If there are no options returned, need to initialize the map
			if conn.Options == nil {
				conn.Options = map[string]string{}
			}
			// remove not necessary parameters for builtin HMS to avoid configuration drift
			if val, exists := conn.Options["builtin"]; exists && val == "true" {
				delete(conn.Options, "host")
				delete(conn.Options, "port")
				delete(conn.Options, "home_workspace_id")
				delete(conn.Options, "database")
			}
			// We need to preserve original sensitive options as API doesn't return them
			var cOrig catalog.CreateConnection
			common.DataToStructPointer(d, s, &cOrig)
			for key, element := range cOrig.Options {
				if slices.Contains(sensitiveOptions, key) {
					conn.Options[key] = element
				}
			}
			if err := common.StructToData(conn, s, d); err != nil {
				return err
			}
			// The API returns only full_name, not parent; rebuild the caller's parent from it so a
			// schema-level connection round-trips without a spurious diff (empty for L1 connections).
			return d.Set("parent", schemaParentFromFullName(conn.FullName))
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			err = validateMetastoreId(ctx, w, d.Get("metastore_id").(string))
			if err != nil {
				return err
			}
			var updateConnectionRequest catalog.UpdateConnection
			common.DataToStructPointer(d, s, &updateConnectionRequest)
			_, connName, err := pi.Unpack(d)
			if err != nil {
				return err
			}
			updateConnectionRequest.Name = connName

			if d.HasChange("owner") {
				_, err = w.Connections.Update(ctx, catalog.UpdateConnection{
					Name:  updateConnectionRequest.Name,
					Owner: updateConnectionRequest.Owner,
				})
				if err != nil {
					return err
				}
			}

			updateConnectionRequest.Owner = ""
			for _, option := range computedOptions {
				delete(updateConnectionRequest.Options, option)
			}
			_, err = w.Connections.Update(ctx, updateConnectionRequest)
			if err != nil {
				if d.HasChange("owner") {
					// Rollback
					old, new := d.GetChange("owner")
					_, rollbackErr := w.Connections.Update(ctx, catalog.UpdateConnection{
						Name:  updateConnectionRequest.Name,
						Owner: old.(string),
					})
					if rollbackErr != nil {
						return common.OwnerRollbackError(err, rollbackErr, old.(string), new.(string))
					}
				}
				return err
			}
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			_, connName, err := pi.Unpack(d)
			if err != nil {
				return err
			}
			return w.Connections.DeleteByName(ctx, connName)
		},
	}
}
