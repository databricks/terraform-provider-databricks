package api

import (
	"context"
	"fmt"

	"github.com/databrickslabs/terraform-provider-databricks/common"
)

// Experiment defines the response object from the API
type Experiment struct {
	ExperimentId     string `json:"experiment_id"`
	Name             string `json:"name"`
	ArtifactLocation string `json:"artifact_location,omitempty"`
	LifecycleStage   string `json:"lifecycle_stage,omitempty"`
	LastUpdateTime   int64  `json:"last_update_time,omitempty"`
	CreationTime     int64  `json:"creation_time,omitempty"`
	Tags             []Tag  `json:"tags,omitempty"`
}

type ExperimentUpdate struct {
	ExperimentId string `json:"experiment_id"`
	NewName      string `json:"new_name"`
}

type Experiments struct {
	Experiment Experiment `json:"experiment"`
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
func (a ExperimentAPI) Create(d *Experiment) error {
	return a.client.Post(a.context, "/mlflow/experiments/create", d, &d)
}

// Read ...
func (a ExperimentAPI) Read(experimentId string) (*Experiment, error) {
	var d Experiments
	err := a.client.Get(a.context, fmt.Sprintf("/mlflow/experiments/get?experiment_id=%s", experimentId), nil, &d)
	if err != nil {
		return nil, err
	}
	return &d.Experiment, nil
}

// Update ...
func (a ExperimentAPI) Update(d *ExperimentUpdate) error {
	return a.client.Post(a.context, "/mlflow/experiments/update", d, &d)
}

// Delete ...
func (a ExperimentAPI) Delete(d *Experiment) error {
	return a.client.Post(a.context, "/mlflow/experiments/delete", d, &d)
}
