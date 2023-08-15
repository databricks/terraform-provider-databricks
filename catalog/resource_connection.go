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
	// Username of current owner of the connection.
	Owner string `json:"owner,omitempty" tf:"force_new,suppress_diff"`
	// An object containing map of key-value properties attached to the
	// connection.
	Properties map[string]string `json:"properties,omitempty" tf:"force_new"`
	// If the connection is read only.
	ReadOnly bool `json:"read_only,omitempty" tf:"force_new,computed"`
}

// suppress diff for sensitive options, which are not returned by the server
func suppressSensitiveOptions(k, old, new string, d *schema.ResourceData) bool {
	//this list will expand as other auth may have different sensitive options
	sensitiveOptions := []string{"user", "password"}
	o, n := d.GetChange("options")
	oldOpt := o.(map[string]any)
	newOpt := n.(map[string]any)
	//loop through the map and ignore diff for sensitive options
	for key, element := range newOpt {
		if slices.Contains(sensitiveOptions, key) {
			continue
		}
		if oldOpt[key] != element {
			return false
		}
	}
	return true
}

func ResourceConnection() *schema.Resource {
	s := common.StructToSchema(ConnectionInfo{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["options"].DiffSuppressFunc = suppressSensitiveOptions
			return m
		})
	pi := common.NewPairID("name", "metastore_id").Schema(
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
			connName, _, err := pi.Unpack(d)
			if err != nil {
				return err
			}
			conn, err := w.Connections.GetByNameArg(ctx, connName)
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
			common.DataToStructPointer(d, s, &updateConnectionRequest)
			connName, _, err := pi.Unpack(d)
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
			connName, _, err := pi.Unpack(d)
			if err != nil {
				return err
			}
			return w.Connections.DeleteByNameArg(ctx, connName)
		},
	}.ToResource()
}
