package compute

import (
	"context"

	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceNodeType returns smallest node depedning on the cloud
func DataSourceNodeType() *schema.Resource {
	s := internal.StructToSchema(NodeTypeRequest{}, func(
		s map[string]*schema.Schema) map[string]*schema.Schema {
		return s
	})
	return &schema.Resource{
		Schema: s,
		ReadContext: func(ctx context.Context, d *schema.ResourceData,
			m interface{}) diag.Diagnostics {
			var this NodeTypeRequest
			err := internal.DataToStructPointer(d, s, &this)
			if err != nil {
				return diag.FromErr(err)
			}
			clustersAPI := NewClustersAPI(ctx, m)
			d.SetId(clustersAPI.GetSmallestNodeType(this))
			return nil
		},
	}
}
