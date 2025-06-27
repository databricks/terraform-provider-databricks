package sql

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/sql/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DashboardEntity defines the parameters that can be set in the resource.
type DashboardEntity struct {
	Name                    string   `json:"name"`
	Tags                    []string `json:"tags,omitempty"`
	Parent                  string   `json:"parent,omitempty" tf:"suppress_diff,force_new"`
	CreatedAt               string   `json:"created_at,omitempty" tf:"computed"`
	UpdatedAt               string   `json:"updated_at,omitempty" tf:"computed"`
	RunAsRole               string   `json:"run_as_role,omitempty" tf:"suppress_diff"`
	DashboardFiltersEnabled bool     `json:"dashboard_filters_enabled,omitempty"`
}

func (d *DashboardEntity) toAPIObject(schema map[string]*schema.Schema, data *schema.ResourceData) (*api.Dashboard, error) {
	// Extract from ResourceData.
	common.DataToStructPointer(data, schema, d)

	// Copy to API object.
	var ad api.Dashboard
	ad.ID = data.Id()
	ad.Name = d.Name
	ad.Tags = append([]string{}, d.Tags...)
	ad.Parent = d.Parent
	ad.DashboardFiltersEnabled = d.DashboardFiltersEnabled
	ad.RunAsRole = d.RunAsRole

	return &ad, nil
}

func (d *DashboardEntity) fromAPIObject(ad *api.Dashboard, schema map[string]*schema.Schema, data *schema.ResourceData) error {
	// Copy from API object.
	d.Name = ad.Name
	d.Tags = append([]string{}, ad.Tags...)
	d.Parent = ad.Parent
	d.UpdatedAt = ad.UpdatedAt
	d.CreatedAt = ad.CreatedAt
	d.DashboardFiltersEnabled = ad.DashboardFiltersEnabled
	d.RunAsRole = ad.RunAsRole

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
func NewDashboardAPI(ctx context.Context, m any) DashboardAPI {
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

func ResourceSqlDashboard() common.Resource {
	s := common.StructToSchema(
		DashboardEntity{},
		common.NoCustomize)

	return common.Resource{
		DeprecationMessage: "This resource is deprecated and will be removed in the future. Please use the `databricks_dashboard` resource.",
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
	}
}
