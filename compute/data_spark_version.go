package compute

import (
	"context"

	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceSparkVersion returns DBR version matching to the specification
func DataSourceSparkVersion() *schema.Resource {
	s := internal.StructToSchema(SparkVersionRequest{}, func(
		s map[string]*schema.Schema) map[string]*schema.Schema {
		return s
	})

	return &schema.Resource{
		Schema: s,
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			var this SparkVersionRequest
			err := internal.DataToStructPointer(d, s, &this)
			if err != nil {
				return diag.FromErr(err)
			}
			version, err := NewClustersAPI(ctx, m).LatestSparkVersion(this)
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId(version)
			return nil
		},
	}
}
