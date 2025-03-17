package catalog

import (
	"context"
	"log"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/exp/slices"
)

var sensitiveOptions = []string{"user", "password", "personalAccessToken", "access_token", "client_secret",
	"pem_private_key", "OAuthPvtKey", "GoogleServiceAccountKeyJson"}

func suppressPemPrivateKeyExpiration(k, old, new string, d *schema.ResourceData) bool {
	if k == "options.pem_private_key_expiration_epoch_sec" {
		log.Printf("[INFO] Suppressing diff on %s", k)
		return true
	}
	return false
}

func ResourceConnection() common.Resource {
	s := common.StructToSchema(catalog.ConnectionInfo{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			for _, v := range []string{"url", "metastore_id", "credential_type", "connection_id",
				"created_at", "created_by", "full_name", "provisioning_info", "securable_type", "updated_at", "updated_by"} {
				common.CustomizeSchemaPath(m, v).SetReadOnly()
			}
			for _, v := range []string{"owner", "read_only"} {
				common.CustomizeSchemaPath(m, v).SetComputed()
			}
			for _, v := range []string{"read_only", "properties", "comment", "connection_type"} {
				common.CustomizeSchemaPath(m, v).SetForceNew()
			}
			common.CustomizeSchemaPath(m, "options").SetSensitive().SetCustomSuppressDiff(suppressPemPrivateKeyExpiration)
			common.CustomizeSchemaPath(m, "name").SetCustomSuppressDiff(common.EqualFoldDiffSuppress)

			return m
		})
	pi := common.NewPairID("metastore_id", "name").Schema(
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			return s
		})
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
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
			pi.Pack(d)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
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
			return common.StructToData(conn, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
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
			delete(updateConnectionRequest.Options, "pem_private_key_expiration_epoch_sec")
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
			w, err := c.WorkspaceClient()
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
