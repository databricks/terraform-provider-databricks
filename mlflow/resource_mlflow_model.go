package mlflow

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/mlflow"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// ModelVersion defines a MLFlow model version as returned by the API
type ModelVersion struct {
	Name                 string `json:"name"`
	Version              string `json:"version"`
	CreationTimestamp    int64  `json:"creation_timestamp,omitempty"`
	LastUpdatedTimestamp int64  `json:"last_updated_timestamp,omitempty"`
	UserID               string `json:"user_id,omitempty"`
	CurrentStage         string `json:"current_stage,omitempty"`
	Source               string `json:"source,omitempty"`
	Status               string `json:"status,omitempty"`
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

func ResourceMlflowModel() *schema.Resource {
	s := common.StructToSchema(
		Model{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			delete(m, "latest_versions")
			return m
		})

	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			model, err := w.RegisteredModels.Create(ctx, mlflow.CreateRegisteredModelRequest{
				Description: d.Get("description").(string),
				Name:        d.Get("name").(string),
				Tags:        d.Get("tags").([]mlflow.RegisteredModelTag),
			})
			if err != nil {
				return err
			}
			d.SetId(model.RegisteredModel.Name)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			model, err := w.RegisteredModels.GetByName(ctx, d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(model, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			err = w.RegisteredModels.Update(ctx, mlflow.UpdateRegisteredModelRequest{
				Description: d.Get("description").(string),
				Name:        d.Get("name").(string),
			})
			return err
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			err = w.RegisteredModels.Delete(ctx, mlflow.DeleteRegisteredModelRequest{
				Name: d.Get("name").(string),
			})
			return err
		},
		Schema: s,
	}.ToResource()
}
