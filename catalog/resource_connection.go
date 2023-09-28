package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/exp/slices"
)

// This structure contains the fields of catalog.UpdateConnection and catalog.CreateConnection
// We need to create this because we need Owner, FullNameArg, SchemaName and CatalogName which aren't present in a single of them.
// We also need to annotate tf:"computed" for the Owner field.
type ConnectionInfo struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty" tf:"force_new"`
	// The type of connection.
	ConnectionType string `json:"connection_type" tf:"force_new"`
	// Unique identifier of parent metastore.
	MetastoreId string `json:"metastore_id,omitempty" tf:"computed"`
	// Name of the connection.
	Name string `json:"name"`
	// Name of the connection.
	NameArg string `json:"-" url:"-"`
	// A map of key-value properties attached to the securable.
	Options map[string]string `json:"options" tf:"sensitive"`
	// An object containing map of key-value properties attached to the
	// connection.
	Properties map[string]string `json:"properties,omitempty" tf:"force_new"`
	// If the connection is read only.
	ReadOnly bool `json:"read_only,omitempty" tf:"force_new,computed"`
}

var sensitiveOptions = []string{"user", "password", "personalAccessToken", "access_token", "client_secret", "OAuthPvtKey"}

func ResourceConnection() *schema.Resource {
	s := common.StructToSchema(ConnectionInfo{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["owner"] = &schema.Schema{
				Type:       schema.TypeString,
				Optional:   true,
				Deprecated: `owner field is deprecated due the Catalog API changes. Owner cannot be specified when creating a connection.`,
			}
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
			var createConnectionRequest catalog.CreateConnection
			common.DataToStructPointer(d, s, &createConnectionRequest)
			conn, err := w.Connections.Create(ctx, createConnectionRequest)
			if err != nil {
				return err
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
			conn, err := w.Connections.GetByNameArg(ctx, connName)
			if err != nil {
				return err
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
			var updateConnectionRequest catalog.UpdateConnection
			common.DataToStructPointer(d, s, &updateConnectionRequest)
			_, connName, err := pi.Unpack(d)
			updateConnectionRequest.NameArg = connName
			if err != nil {
				return err
			}
			conn, err := w.Connections.Update(ctx, updateConnectionRequest)
			if err != nil {
				return err
			}
			// We need to repack the Id as the name may have changed
			d.Set("name", conn.Name)
			pi.Pack(d)
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
			return w.Connections.DeleteByNameArg(ctx, connName)
		},
	}.ToResource()
}
