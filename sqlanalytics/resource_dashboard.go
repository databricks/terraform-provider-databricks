package sqlanalytics

import (
	"context"
	"fmt"

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
	common.DataToStructPointer(data, schema, d)

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

// NewDashboardAPI ...
func NewDashboardAPI(ctx context.Context, m interface{}) DashboardAPI {
	return DashboardAPI{m.(*common.DatabricksClient), ctx}
}

// DashboardAPI ...
type DashboardAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Create ...
func (a DashboardAPI) Create(d *api.Dashboard) error {
	return a.client.Post(a.context, "/preview/sql/dashboards", d, &d)
}

// Read ...
func (a DashboardAPI) Read(dashboardID string) (*api.Dashboard, error) {
	var d api.Dashboard
	err := a.client.Get(a.context, fmt.Sprintf("/preview/sql/dashboards/%s", dashboardID), nil, &d)
	if err != nil {
		return nil, err
	}

	return &d, nil
}

// Update ...
func (a DashboardAPI) Update(dashboardID string, d *api.Dashboard) error {
	return a.client.Post(a.context, fmt.Sprintf("/preview/sql/dashboards/%s", dashboardID), d, nil)
}

// Delete ...
func (a DashboardAPI) Delete(dashboardID string) error {
	return a.client.Delete(a.context, fmt.Sprintf("/preview/sql/dashboards/%s", dashboardID), nil)
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

			err = NewDashboardAPI(ctx, c).Create(ad)
			if err != nil {
				return err
			}

			// No need to set anything because the resource is going to be
			// read immediately after being created.
			data.SetId(ad.ID)
			return nil
		},
		Read: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			ad, err := NewDashboardAPI(ctx, c).Read(data.Id())
			if err != nil {
				return err
			}

			var d DashboardEntity
			return d.fromAPIObject(ad, s, data)
		},
		Update: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var d DashboardEntity
			ad, err := d.toAPIObject(s, data)
			if err != nil {
				return err
			}

			return NewDashboardAPI(ctx, c).Update(data.Id(), ad)
		},
		Delete: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			return NewDashboardAPI(ctx, c).Delete(data.Id())
		},
		Schema: s,
	}.ToResource()
}
