package catalog

import (
	"context"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type SchemasAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

func NewSchemasAPI(ctx context.Context, m interface{}) SchemasAPI {
	return SchemasAPI{m.(*common.DatabricksClient), ctx}
}

type SchemaInfo struct {
	Name        string            `json:"name"`
	CatalogName string            `json:"catalog_name"`
	Comment     string            `json:"comment,omitempty"`
	Properties  map[string]string `json:"properties,omitempty"`
	Owner       string            `json:"owner,omitempty" tf:"computed"`
	MetastoreID string            `json:"metastore_id,omitempty" tf:"computed"`
	FullName    string            `json:"full_name,omitempty" tf:"computed"`
}

func (a SchemasAPI) createSchema(si *SchemaInfo) error {
	return a.client.Post(a.context, "/unity-catalog/schemas", si, si)
}

func (a SchemasAPI) getSchema(name string) (si SchemaInfo, err error) {
	err = a.client.Get(a.context, "/unity-catalog/schemas/"+name, nil, &si)
	return
}

func (a SchemasAPI) updateSchema(si SchemaInfo) error {
	return a.client.Patch(a.context, "/unity-catalog/schemas/"+si.FullName, si)
}

func (a SchemasAPI) deleteSchema(name string) error {
	return a.client.Delete(a.context, "/unity-catalog/schemas/"+name, nil)
}

func ResourceSchema() *schema.Resource {
	s := common.StructToSchema(SchemaInfo{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			delete(m, "full_name")
			return m
		})
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var si SchemaInfo
			if err := common.DataToStructPointer(d, s, &si); err != nil {
				return err
			}
			if err := NewSchemasAPI(ctx, c).createSchema(&si); err != nil {
				return err
			}
			d.SetId(si.FullName)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			si, err := NewSchemasAPI(ctx, c).getSchema(d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(si, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var si SchemaInfo
			if err := common.DataToStructPointer(d, s, &si); err != nil {
				return err
			}
			si.FullName = d.Id()
			return NewSchemasAPI(ctx, c).updateSchema(si)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewSchemasAPI(ctx, c).deleteSchema(d.Id())
		},
	}.ToResource()
}
