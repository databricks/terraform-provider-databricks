package mlflow

import (
	"context"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/mlflow/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// MLFlowModel defines the parameters that can be set in the resource.
type Model struct {
	Name        string    `json:"name"`
	Tags        []api.Tag `json:"tags,omitempty" tf:"force_new"`
	Description string    `json:"description,omitempty"`
}

func (d *Model) toAPIObject(schema map[string]*schema.Schema, data *schema.ResourceData) (*api.Model, error) {
	// Extract from ResourceData.
	if err := common.DataToStructPointer(data, schema, d); err != nil {
		return nil, err
	}

	// Copy to API object.
	var ad api.Model
	ad.Name = d.Name
	ad.Tags = d.Tags
	ad.Description = d.Description

	return &ad, nil
}

func (d *Model) fromAPIObject(ad *api.Model, schema map[string]*schema.Schema, data *schema.ResourceData) error {
	// Copy from API object.
	d.Name = ad.Name
	d.Tags = ad.Tags
	d.Description = ad.Description

	// Pass to ResourceData.
	if err := common.StructToData(*d, schema, data); err != nil {
		return err
	}

	return nil
}

// ResourceDashboard ...
func ResourceMLFlowModel() *schema.Resource {
	s := common.StructToSchema(
		Model{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			return m
		})

	return common.Resource{
		Create: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var d Model
			ad, err := d.toAPIObject(s, data)
			if err != nil {
				return err
			}

			err = api.NewModelAPI(ctx, c).Create(ad)
			if err != nil {
				return err
			}

			// No need to set anything because the resource is going to be
			// read immediately after being created.
			data.SetId(ad.Name)
			return nil
		},
		Read: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			ad, err := api.NewModelAPI(ctx, c).Read(data.Id())
			if err != nil {
				return err
			}

			var d Model
			return d.fromAPIObject(ad, s, data)
		},
		Update: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var d Model
			ad, err := d.toAPIObject(s, data)
			if err != nil {
				return err
			}

			return api.NewModelAPI(ctx, c).Update(ad)
		},
		Delete: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var d Model
			ad, err := d.toAPIObject(s, data)
			if err != nil {
				return err
			}
			return api.NewModelAPI(ctx, c).Delete(ad)
		},
		Schema: s,
	}.ToResource()
}
