package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// This structure contains the fields of catalog.UpdateConnection and catalog.CreateConnection
// We need to create this because we need Owner, FullNameArg, SchemaName and CatalogName which aren't present in a single of them.
// We also need to annotate tf:"computed" for the Owner field.
type ConnectionInfo struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty" tf:"force_new"`
	// The type of connection.
	ConnectionType string `json:"connection_type" tf:"force_new"`
	// Name of the connection.
	Name string `json:"name"`
	// Name of the connection.
	NameArg string `json:"-" url:"-"`
	// A map of key-value properties attached to the securable.
	OptionsKvpairs map[string]string `json:"options_kvpairs" tf:"alias:options,sensitive"`
	// Username of current owner of the connection.
	Owner string `json:"owner,omitempty" tf:"force_new"`
	// An object containing map of key-value properties attached to the
	// connection.
	PropertiesKvpairs map[string]string `json:"properties_kvpairs,omitempty" tf:"alias:properties,force_new"`
	// If the connection is read only.
	ReadOnly bool `json:"read_only,omitempty" tf:"force_new"`
}

func ResourceConnection() *schema.Resource {
	s := common.StructToSchema(ConnectionInfo{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			return m
		})
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var createConnectionRequest catalog.CreateConnection
			var alias ConnectionInfo
			common.DataToStructPointer(d, s, &createConnectionRequest)
			common.DataToStructPointer(d, s, &alias)
			//workaround as cannot set tf:"alias" for the Go SDK struct
			createConnectionRequest.OptionsKvpairs = alias.OptionsKvpairs
			createConnectionRequest.PropertiesKvpairs = alias.PropertiesKvpairs
			conn, err := w.Connections.Create(ctx, createConnectionRequest)
			if err != nil {
				return err
			}
			d.SetId(conn.Name)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			conn, err := w.Connections.GetByNameArg(ctx, d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(conn, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var updateConnectionRequest catalog.UpdateConnection
			var alias ConnectionInfo
			common.DataToStructPointer(d, s, &updateConnectionRequest)
			common.DataToStructPointer(d, s, &alias)
			//workaround as cannot set tf:"alias" for the Go SDK struct
			updateConnectionRequest.OptionsKvpairs = alias.OptionsKvpairs
			updateConnectionRequest.NameArg = d.Id()
			conn, err := w.Connections.Update(ctx, updateConnectionRequest)
			if err != nil {
				return err
			}
			// We need to update the resource Id because Name is updatable and FullName consists of Name,
			// So if we don't update the field then the requests would be made to old FullName which doesn't exists.
			d.SetId(conn.Name)
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return w.Connections.DeleteByNameArg(ctx, d.Id())
		},
	}.ToResource()
}
