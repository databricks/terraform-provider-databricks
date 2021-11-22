package mlflow

import (
	"context"
	"fmt"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Experiment defines the parameters that can be set in the resource.
type Experiment struct {
	Name             string `json:"name"`
	ArtifactLocation string `json:"artifact_location,omitempty" tf:"force_new"`
	Description      string `json:"description,omitempty"`
}

// ExperimentDto defines the response object from the API
type ExperimentDto struct {
	ExperimentId     string `json:"experiment_id"`
	Name             string `json:"name"`
	ArtifactLocation string `json:"artifact_location,omitempty"`
	LifecycleStage   string `json:"lifecycle_stage,omitempty"`
	LastUpdateTime   int64  `json:"last_update_time,omitempty"`
	CreationTime     int64  `json:"creation_time,omitempty"`
}

type ExperimentUpdateDto struct {
	ExperimentId string `json:"experiment_id"`
	NewName      string `json:"new_name"`
}

type ExperimentsDto struct {
	Experiment ExperimentDto `json:"experiment"`
}

// ExperimentAPI ...
type ExperimentAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// NewExperimentAPI ...
func NewExperimentAPI(ctx context.Context, m interface{}) ExperimentAPI {
	return ExperimentAPI{m.(*common.DatabricksClient), ctx}
}

// Create ...
func (a ExperimentAPI) Create(d *ExperimentDto) error {
	return a.client.Post(a.context, "/mlflow/experiments/create", d, &d)
}

// Read ...
func (a ExperimentAPI) Read(experimentId string) (*ExperimentDto, error) {
	var d ExperimentsDto
	err := a.client.Get(a.context, fmt.Sprintf("/mlflow/experiments/get?experiment_id=%s", experimentId), nil, &d)
	if err != nil {
		return nil, err
	}
	return &d.Experiment, nil
}

// Update ...
func (a ExperimentAPI) Update(d *ExperimentUpdateDto) error {
	return a.client.Post(a.context, "/mlflow/experiments/update", d, &d)
}

// Delete ...
func (a ExperimentAPI) Delete(d *ExperimentDto) error {
	return a.client.Post(a.context, "/mlflow/experiments/delete", d, &d)
}

///func ResourceMLFlowExperiment() {}
func ResourceMLFlowExperiment() *schema.Resource {
	s := common.StructToSchema(
		Experiment{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			return m
		})

	return common.Resource{
		Create: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var ad ExperimentDto
			if err := common.DataToStructPointer(data, s, &ad); err != nil {
				return err
			}
			if err := NewExperimentAPI(ctx, c).Create(&ad); err != nil {
				return err
			}
			data.SetId(ad.ExperimentId)
			return nil
		},
		Read: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var d Experiment
			ad, err := NewExperimentAPI(ctx, c).Read(data.Id())
			if err != nil {
				return err
			}

			if err := common.StructToData(d, s, data); err != nil {
				return err
			}

			data.Set("name", ad.Name)
			data.SetId(ad.ExperimentId)
			return nil
		},
		Update: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var ad ExperimentDto
			if err := common.DataToStructPointer(data, s, &ad); err != nil {
				return err
			}
			updateDoc := ExperimentUpdateDto{ExperimentId: data.Id(), NewName: ad.Name}
			return NewExperimentAPI(ctx, c).Update(&updateDoc)
		},
		Delete: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			ad := ExperimentDto{ExperimentId: data.Id()}
			return NewExperimentAPI(ctx, c).Delete(&ad)
		},
		StateUpgraders: []schema.StateUpgrader{},
		Schema:         s,
		SchemaVersion:  0,
		Timeouts:       &schema.ResourceTimeout{},
	}.ToResource()
}
