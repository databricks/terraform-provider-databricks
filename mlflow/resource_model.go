package mlflow

import (
	"context"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// ModelVersion defines a MLFlow model version as returned by the API
type ModelVersion struct {
	Name                 string `json:"name" tf:"computed"`
	Version              string `json:"version" tf:"computed"`
	CreationTimestamp    int64  `json:"creation_timestamp,omitempty" tf:"computed"`
	LastUpdatedTimestamp int64  `json:"last_updated_timestamp,omitempty" tf:"computed"`
	UserID               string `json:"user_id,omitempty" tf:"computed"`
	CurrentStage         string `json:"current_stage,omitempty" tf:"computed"`
	Source               string `json:"source,omitempty" tf:"computed"`
	Status               string `json:"status,omitempty" tf:"computed"`
}

// Model defines the response object from the API
type Model struct {
	Name                 string         `json:"name" tf:"force_new"`
	CreationTimestamp    int64          `json:"creation_timestamp,omitempty" tf:"computed"`
	LastUpdatedTimestamp int64          `json:"last_updated_timestamp,omitempty" tf:"computed"`
	UserID               string         `json:"user_id,omitempty" tf:"computed"`
	LatestVersions       []ModelVersion `json:"latest_versions,omitempty" tf:"computed"`
	Description          string         `json:"description,omitempty"`
	Tags                 []Tag          `json:"tags,omitempty"`
	RegisteredModelID    string         `json:"id,omitempty" tf:"computed,alias:registered_model_id"`
}

// registeredModel defines response from GET API op
type registeredModel struct {
	RegisteredModelDatabricks Model `json:"registered_model_databricks"`
}

type ModelsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

func NewModelsAPI(ctx context.Context, m interface{}) ModelsAPI {
	return ModelsAPI{m.(*common.DatabricksClient), ctx}
}

func (a ModelsAPI) Create(m *Model) error {
	return a.client.Post(a.context, "/mlflow/registered-models/create", m, m)
}

func (a ModelsAPI) Read(name string) (*Model, error) {
	var m registeredModel
	err := a.client.Get(a.context, "/mlflow/databricks/registered-models/get", map[string]string{
		"name": name,
	}, &m)
	if err != nil {
		return nil, err
	}
	return &m.RegisteredModelDatabricks, nil
}

// Update the model entity
func (a ModelsAPI) Update(m *Model) error {
	return a.client.Patch(a.context, "/mlflow/registered-models/update", m)
}

// Delete removes the model by its name
func (a ModelsAPI) Delete(name string) error {
	return a.client.Delete(a.context, "/mlflow/registered-models/delete", map[string]string{
		"name": name,
	})
}

func ResourceMLFlowModel() *schema.Resource {
	s := common.StructToSchema(
		Model{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			delete(m, "latest_versions")
			return m
		})

	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var m Model
			common.DataToStructPointer(d, s, &m)
			if err := NewModelsAPI(ctx, c).Create(&m); err != nil {
				return err
			}
			d.SetId(m.Name)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			m, err := NewModelsAPI(ctx, c).Read(d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(*m, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var m Model
			common.DataToStructPointer(d, s, &m)
			return NewModelsAPI(ctx, c).Update(&m)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewModelsAPI(ctx, c).Delete(d.Id())
		},
		Schema: s,
	}.ToResource()
}
