package api

import (
	"context"
	"fmt"

	"github.com/databrickslabs/terraform-provider-databricks/common"
)

// MLFlowModelAPI defines the response object from the API
type Model struct {
	Name                 string   `json:"name"`
	CreationTimestamp    int64    `json:"creation_timestamp,omitempty" tf:"computed"`
	LastUpdatedTimestamp int64    `json:"last_updated_timestamp,omitempty" tf:"computed"`
	UserID               string   `json:"user_id,omitempty" tf:"computed"`
	LatestVersions       []string `json:"latest_versions,omitempty" tf:"computed"`
	Description          string   `json:"description,omitempty"`
	Tags                 []Tag    `json:"tags,omitempty"`
}

type RegisteredModel struct {
	RegisteredModel Model `json:"registered_model"`
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
func (a ModelAPI) Create(d *Model) error {
	return a.client.Post(a.context, "/mlflow/registered-models/create", d, &d)
}

// Read ...
func (a ModelAPI) Read(modelName string) (*Model, error) {
	var d RegisteredModel
	err := a.client.Get(a.context, fmt.Sprintf("/mlflow/registered-models/get?name=%s", modelName), nil, &d)
	if err != nil {
		return nil, err
	}

	return &d.RegisteredModel, nil
}

// Update ...
func (a ModelAPI) Update(d *Model) error {
	return a.client.Patch(a.context, "/mlflow/registered-models/update", d)
}

// Delete ...
func (a ModelAPI) Delete(d *Model) error {
	return a.client.Delete(a.context, "/mlflow/registered-models/delete", d)
}
