package mlflow

import (
	"context"
	"fmt"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Model defines the parameters that can be set in the resource.
type Model struct {
	Name        string `json:"name" tf:"force_new"`
	Tags        []Tag  `json:"tags,omitempty" tf:"force_new"`
	Description string `json:"description,omitempty"`
}

// ModelDto defines the response object from the API
type ModelDto struct {
	Name                 string   `json:"name"`
	CreationTimestamp    int64    `json:"creation_timestamp,omitempty" tf:"computed"`
	LastUpdatedTimestamp int64    `json:"last_updated_timestamp,omitempty" tf:"computed"`
	UserID               string   `json:"user_id,omitempty" tf:"computed"`
	LatestVersions       []string `json:"latest_versions,omitempty" tf:"computed"`
	Description          string   `json:"description,omitempty"`
	Tags                 []Tag    `json:"tags,omitempty"`
}

// RegisteredModel defines response from GET API op
type RegisteredModelDto struct {
	RegisteredModel ModelDto `json:"registered_model"`
}

// ModelAPI ...
type ModelAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// NewModelAPI ...
func NewModelAPI(ctx context.Context, m interface{}) ModelAPI {
	return ModelAPI{m.(*common.DatabricksClient), ctx}
}

// Create ...
func (a ModelAPI) Create(d *ModelDto) error {
	return a.client.Post(a.context, "/mlflow/registered-models/create", d, &d)
}

// Read ...
func (a ModelAPI) Read(modelName string) (*ModelDto, error) {
	var d RegisteredModelDto
	err := a.client.Get(a.context, fmt.Sprintf("/mlflow/registered-models/get?name=%s", modelName), nil, &d)
	if err != nil {
		return nil, err
	}

	return &d.RegisteredModel, nil
}

// Update ...
func (a ModelAPI) Update(d *ModelDto) error {
	return a.client.Patch(a.context, "/mlflow/registered-models/update", d)
}

// Delete ...
func (a ModelAPI) Delete(d *ModelDto) error {
	return a.client.Delete(a.context, "/mlflow/registered-models/delete", d)
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
			var ad ModelDto
			if err := common.DataToStructPointer(data, s, &ad); err != nil {
				return nil
			}
			if err := NewModelAPI(ctx, c).Create(&ad); err != nil {
				return err
			}
			// No need to set anything because the resource is going to be
			// read immediately after being created.
			data.SetId(ad.Name)
			return nil
		},
		Read: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var d Model
			ad, err := NewModelAPI(ctx, c).Read(data.Id())
			if err != nil {
				return err
			}

			if err := common.StructToData(d, s, data); err != nil {

				return err
			}
			data.Set("name", ad.Name)
			data.Set("description", ad.Description)

			return nil
		},
		Update: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var ad ModelDto
			if err := common.DataToStructPointer(data, s, &ad); err != nil {
				return nil
			}

			return NewModelAPI(ctx, c).Update(&ad)
		},
		Delete: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var ad ModelDto
			if err := common.DataToStructPointer(data, s, &ad); err != nil {
				return nil
			}
			return NewModelAPI(ctx, c).Delete(&ad)
		},
		Schema: s,
	}.ToResource()
}
