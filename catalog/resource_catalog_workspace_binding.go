package catalog

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var getSecurableName = func(d *schema.ResourceData) string {
	securableName, ok := d.GetOk("catalog_name")
	if !ok {
		securableName = d.Get("securable_name")
	}
	return securableName.(string)
}

func ResourceCatalogWorkspaceBinding() *schema.Resource {
	p := common.NewPairID("workspace_id", fmt.Sprintf("%s|%s", "securable_type", "securable_name"))
	workspaceBindingSchema := common.StructToSchema(catalog.WorkspaceBinding{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["catalog_name"] = &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{"catalog_name", "securable_name"},
				Deprecated:   "Please use 'securable_name' and 'securable_type instead.",
			}
			m["securable_name"] = &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{"catalog_name", "securable_name"},
			}
			m["securable_type"] = &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "catalog",
			}
			m["binding_type"].Default = catalog.WorkspaceBindingBindingTypeBindingTypeReadWrite
			m["binding_type"].ValidateFunc = validation.StringInSlice([]string{
				string(catalog.WorkspaceBindingBindingTypeBindingTypeReadWrite),
				string(catalog.WorkspaceBindingBindingTypeBindingTypeReadOnly),
			}, false)
			return m
		},
	)
	return common.Resource{
		Schema: workspaceBindingSchema,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var update catalog.WorkspaceBinding
			common.DataToStructPointer(d, workspaceBindingSchema, &update)
			_, err = w.WorkspaceBindings.UpdateBindings(ctx, catalog.UpdateWorkspaceBindingsParameters{
				Add:           []catalog.WorkspaceBinding{update},
				SecurableName: getSecurableName(d),
				SecurableType: d.Get("securable_type").(string),
			})
			d.Set("securable_name", getSecurableName(d))
			p.Pack(d)
			return err
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			workspaceId := int64(d.Get("workspace_id").(int))
			bindings, err := w.WorkspaceBindings.GetBindings(ctx, catalog.GetBindingsRequest{
				SecurableName: getSecurableName(d),
				SecurableType: d.Get("securable_type").(string),
			})
			if err != nil {
				return err
			}
			for _, binding := range bindings.Bindings {
				if binding.WorkspaceId == workspaceId {
					return common.StructToData(binding, workspaceBindingSchema, d)
				}
			}
			return apierr.NotFound("Catalog has no binding to this workspace")
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var update catalog.WorkspaceBinding
			common.DataToStructPointer(d, workspaceBindingSchema, &update)
			_, err = w.WorkspaceBindings.UpdateBindings(ctx, catalog.UpdateWorkspaceBindingsParameters{
				Remove:        []catalog.WorkspaceBinding{update},
				SecurableName: getSecurableName(d),
				SecurableType: d.Get("securable_type").(string),
			})
			return err
		},
	}.ToResource()
}
