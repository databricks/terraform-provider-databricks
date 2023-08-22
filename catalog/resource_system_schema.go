package catalog

import (
	"context"
	"fmt"
	"log"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/exp/slices"
)

func ResourceSystemSchema() *schema.Resource {
	create := func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
		o, n := d.GetChange("system_schema")
		old, okOld := o.([]any)
		new, okNew := n.([]any)
		if !okNew || !okOld {
			return fmt.Errorf("internal type casting error")
		}
		var oldSchemas, newSchemas []string
		for _, schema := range old {
			oldSchemas = append(oldSchemas, schema.(string))
		}
		for _, schema := range new {
			newSchemas = append(newSchemas, schema.(string))
		}
		log.Printf("[DEBUG] Old system schemas: %v, new: %v", oldSchemas, newSchemas)
		w, err := c.WorkspaceClient()
		if err != nil {
			return err
		}
		metastoreSummary, err := w.Metastores.Summary(ctx)
		if err != nil {
			return err
		}
		//enable new schemas that is not enabled
		for _, schema := range newSchemas {
			if !slices.Contains(oldSchemas, schema) {
				err := w.SystemSchemas.Enable(ctx, catalog.EnableRequest{
					MetastoreId: metastoreSummary.MetastoreId,
					SchemaName:  catalog.EnableSchemaName(schema),
				})
				if err != nil {
					return err
				}
			}
		}
		//disable old schemas that is not needed
		for _, schema := range oldSchemas {
			if !slices.Contains(newSchemas, schema) {
				err := w.SystemSchemas.DisableByMetastoreIdAndSchemaName(ctx, metastoreSummary.MetastoreId, catalog.DisableSchemaName(schema))
				if err != nil {
					return err
				}
			}
		}
		d.SetId(metastoreSummary.MetastoreId)
		return nil
	}
	return common.Resource{
		Schema: map[string]*schema.Schema{
			"system_schema": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
		Create: create,
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			metastoreSummary, err := w.Metastores.Summary(ctx)
			if err != nil {
				return err
			}
			systemSchemaInfo, err := w.SystemSchemas.ListByMetastoreId(ctx, metastoreSummary.MetastoreId)
			if err != nil {
				return err
			}
			//only collect schemas that have been enabled
			var schemaEnabled []string
			for _, schema := range systemSchemaInfo.Schemas {
				if schema.State == catalog.SystemSchemaInfoStateEnableCompleted {
					schemaEnabled = append(schemaEnabled, schema.Schema)
				}
			}
			d.Set("system_schema", schemaEnabled)
			return nil
		},
		Update: create,
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			systemSchemas := d.Get("system_schema").([]any)
			metastoreSummary, err := w.Metastores.Summary(ctx)
			if err != nil {
				return err
			}
			for _, schema := range systemSchemas {
				err := w.SystemSchemas.DisableByMetastoreIdAndSchemaName(ctx, metastoreSummary.MetastoreId, catalog.DisableSchemaName(schema.(string)))
				if err != nil {
					return err
				}
			}
			return nil
		},
	}.ToResource()
}
