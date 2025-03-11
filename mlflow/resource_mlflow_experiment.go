package mlflow

import (
	"context"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Experiment defines the response object from the API
type Experiment struct {
	Name             string          `json:"name"`
	ArtifactLocation string          `json:"artifact_location,omitempty" tf:"force_new,suppress_diff"`
	Tags             []ExperimentTag `json:"tags,omitempty"`
	ExperimentId     string          `json:"experiment_id,omitempty" tf:"computed"`
	LifecycleStage   string          `json:"lifecycle_stage,omitempty" tf:"computed"`
	LastUpdateTime   int64           `json:"last_update_time,omitempty" tf:"computed"`
	CreationTime     int64           `json:"creation_time,omitempty" tf:"computed"`
}

// A tag for an experiment.
type ExperimentTag struct {
	// The tag key.
	Key string `json:"key"`
	// The tag value.
	Value string `json:"value"`
}

type experimentUpdate struct {
	ExperimentId string `json:"experiment_id"`
	NewName      string `json:"new_name"`
}

type experimentWrapper struct {
	Experiment Experiment `json:"experiment"`
}

// ExperimentsAPI ...
type ExperimentsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// NewExperimentsAPI ...
func NewExperimentsAPI(ctx context.Context, m any) ExperimentsAPI {
	return ExperimentsAPI{m.(*common.DatabricksClient), ctx}
}

// Create ...
func (a ExperimentsAPI) Create(d *Experiment) error {
	return a.client.Post(a.context, "/mlflow/experiments/create", d, &d)
}

// Read ...
func (a ExperimentsAPI) Read(experimentId string) (*Experiment, error) {
	var d experimentWrapper
	err := a.client.Get(a.context, "/mlflow/experiments/get", map[string]string{
		"experiment_id": experimentId,
	}, &d)
	if err != nil {
		return nil, err
	}
	return &d.Experiment, nil
}

// Update ...
func (a ExperimentsAPI) Update(e *experimentUpdate) error {
	return a.client.Post(a.context, "/mlflow/experiments/update", e, &e)
}

// Delete ...
func (a ExperimentsAPI) Delete(id string) error {
	return a.client.Post(a.context, "/mlflow/experiments/delete", map[string]string{
		"experiment_id": id,
	}, nil)
}

func ResourceMlflowExperiment() common.Resource {
	s := common.StructToSchema(
		Experiment{},
		common.NoCustomize)

	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var e Experiment
			common.DataToStructPointer(d, s, &e)
			if err := NewExperimentsAPI(ctx, c).Create(&e); err != nil {
				return err
			}
			d.SetId(e.ExperimentId)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			e, err := NewExperimentsAPI(ctx, c).Read(d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(*e, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var e Experiment
			common.DataToStructPointer(d, s, &e)
			updateDoc := experimentUpdate{ExperimentId: d.Id(), NewName: e.Name}
			return NewExperimentsAPI(ctx, c).Update(&updateDoc)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewExperimentsAPI(ctx, c).Delete(d.Id())
		},
		StateUpgraders: []schema.StateUpgrader{},
		Schema:         s,
		SchemaVersion:  0,
		Timeouts:       &schema.ResourceTimeout{},
	}
}
