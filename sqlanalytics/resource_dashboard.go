package sqlanalytics

import (
	"context"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/sqlanalytics/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DashboardEntity defines the parameters that can be set in the resource.
type DashboardEntity struct {
	Name string   `json:"name"`
	Tags []string `json:"tags,omitempty"`
}

func (d *DashboardEntity) toAPIObject(schema map[string]*schema.Schema, data *schema.ResourceData) (*api.Dashboard, error) {
	// Extract from ResourceData.
	if err := common.DataToStructPointer(data, schema, d); err != nil {
		return nil, err
	}

	// Copy to API object.
	var ad api.Dashboard
	ad.ID = data.Id()
	ad.Name = d.Name
	ad.Tags = append([]string{}, d.Tags...)

	return &ad, nil
}

func (d *DashboardEntity) fromAPIObject(ad *api.Dashboard, schema map[string]*schema.Schema, data *schema.ResourceData) error {
	// Copy from API object.
	d.Name = ad.Name
	d.Tags = append([]string{}, ad.Tags...)

	// Pass to ResourceData.
	if err := common.StructToData(*d, schema, data); err != nil {
		return err
	}

	// Overwrite `tags` in case they're empty on the server side.
	// This would have been skipped by `common.StructToData` because of slice emptiness.
	// Ideally, the reflection code also sets empty values, but we'd risk
	// clobbering values we actually want to keep around in existing code.
	data.Set("tags", ad.Tags)
	return nil
}

// ResourceDashboard ...
func ResourceDashboard() *schema.Resource {
	s := common.StructToSchema(
		DashboardEntity{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			return m
		})

	return common.Resource{
		Create: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var d DashboardEntity
			ad, err := d.toAPIObject(s, data)
			if err != nil {
				return err
			}

			adNew, err := api.NewWrapper(ctx, c).CreateDashboard(ad)
			if err != nil {
				return err
			}

			// No need to set anything because the resource is going to be
			// read immediately after being created.
			data.SetId(adNew.ID)
			return nil
		},
		Read: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var d DashboardEntity
			ad, err := d.toAPIObject(s, data)
			if err != nil {
				return err
			}

			adNew, err := api.NewWrapper(ctx, c).ReadDashboard(ad)
			if err != nil {
				return err
			}

			return d.fromAPIObject(adNew, s, data)
		},
		Update: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var d DashboardEntity
			ad, err := d.toAPIObject(s, data)
			if err != nil {
				return err
			}

			_, err = api.NewWrapper(ctx, c).UpdateDashboard(ad)
			if err != nil {
				return err
			}

			// No need to set anything because the resource is going to be
			// read immediately after being created.
			return nil
		},
		Delete: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var d DashboardEntity
			ad, err := d.toAPIObject(s, data)
			if err != nil {
				return err
			}

			return api.NewWrapper(ctx, c).DeleteDashboard(ad)
		},
		Schema: s,
	}.ToResource()
}
