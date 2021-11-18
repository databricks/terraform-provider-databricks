package mlflow

import (
	"context"
	"fmt"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// MLFlowExperiment defines the parameters that can be set in the resource.
type MLFlowExperiment struct {
	Name             string `json:"name"`
	ArtifactLocation string `json:"artifact_location,omitempty"`
	Tags             []Tag  `json:"tags,omitempty" tf:"force_new"`
	Description      string `json:"description,omitempty"`
}

// MLFlowExperimentAPI defines the response object from the API
type MLFlowExperimentAPI struct {
	ExperimentId     string `json:"experiment_id"`
	Name             string `json:"name"`
	ArtifactLocation string `json:"artifact_location,omitempty"`
	LifecycleStage   string `json:"lifecycle_stage,omitempty"`
	LastUpdateTime   int64  `json:"last_update_time,omitempty"`
	CreationTime     int64  `json:"creation_time,omitempty"`
	Tags             []Tag  `json:"tags,omitempty" tf:"force_new"`
}

type MLFlowExperimentUpdateAPI struct {
	ExperimentId string `json:"experiment_id"`
	NewName      string `json:"new_name"`
}

type MLFlowExperimentsAPI struct {
	Experiment MLFlowExperimentAPI `json:"experiment"`
}

func (d *MLFlowExperiment) toAPIObject(schema map[string]*schema.Schema, data *schema.ResourceData) (*MLFlowExperimentAPI, error) {
	// Extract from ResourceData.
	if err := common.DataToStructPointer(data, schema, d); err != nil {
		return nil, err
	}

	// Copy to API object.
	var ad MLFlowExperimentAPI
	ad.Name = d.Name
	ad.ExperimentId = data.Id()
	ad.Tags = d.Tags
	ad.ArtifactLocation = d.ArtifactLocation

	return &ad, nil
}

func (d *MLFlowExperiment) fromAPIObject(ad *MLFlowExperimentAPI, schema map[string]*schema.Schema, data *schema.ResourceData) error {
	// Copy from API object.
	d.Name = ad.Name
	d.Tags = ad.Tags

	// Pass to ResourceData.
	if err := common.StructToData(*d, schema, data); err != nil {
		return err
	}

	// Overwrite `tags` in case they're empty on the server side.
	// This would have been skipped by `common.StructToData` because of slice emptiness.
	// Ideally, the reflection code also sets empty values, but we'd risk
	// clobbering values we actually want to keep around in existing code.
	data.Set("tags", ad.Tags)
	data.Set("name", ad.Name)
	data.SetId(ad.ExperimentId)

	return nil
}

// NewMLFlowExperimentAPI ...
func NewMLFlowExperimentAPI(ctx context.Context, m interface{}) MLFlowExpAPI {
	return MLFlowExpAPI{m.(*common.DatabricksClient), ctx}
}

// MLFlowExpAPI ...
type MLFlowExpAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Create ...
func (a MLFlowExpAPI) Create(d *MLFlowExperimentAPI) error {
	return a.client.Post(a.context, "/mlflow/experiments/create", d, &d)
}

// Read ...
func (a MLFlowExpAPI) Read(experimentId string) (*MLFlowExperimentAPI, error) {
	var d MLFlowExperimentsAPI
	err := a.client.Get(a.context, fmt.Sprintf("/mlflow/experiments/get?experiment_id=%s", experimentId), nil, &d)
	if err != nil {
		return nil, err
	}
	return &d.Experiment, nil
}

// Update ...
func (a MLFlowExpAPI) Update(d *MLFlowExperimentUpdateAPI) error {
	return a.client.Post(a.context, "/mlflow/experiments/update", d, &d)
}

// Delete ...
func (a MLFlowExpAPI) Delete(d *MLFlowExperimentAPI) error {
	return a.client.Post(a.context, "/mlflow/experiments/delete", d, &d)
}

///func ResourceMLFlowExperiment() {}
func ResourceMLFlowExperiment() *schema.Resource {
	s := common.StructToSchema(
		MLFlowExperiment{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			return m
		})

	return common.Resource{
		Create: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var d MLFlowExperiment
			ad, err := d.toAPIObject(s, data)
			if err != nil {
				return err
			}

			err = NewMLFlowExperimentAPI(ctx, c).Create(ad)
			if err != nil {
				return err
			}

			// No need to set anything because the resource is going to be
			// read immediately after being created.
			data.SetId(ad.ExperimentId)
			data.Set("name", ad.Name)

			return nil
		},
		Read: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			ad, err := NewMLFlowExperimentAPI(ctx, c).Read(data.Id())
			if err != nil {
				return err
			}

			var d MLFlowExperiment
			return d.fromAPIObject(ad, s, data)
		},
		Update: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var d MLFlowExperiment
			ad, err := d.toAPIObject(s, data)
			if err != nil {
				return err
			}
			updateDoc := MLFlowExperimentUpdateAPI{
				ExperimentId: ad.ExperimentId,
				NewName:      ad.Name,
			}
			return NewMLFlowExperimentAPI(ctx, c).Update(&updateDoc)
		},
		Delete: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var d MLFlowExperiment
			ad, err := d.toAPIObject(s, data)
			if err != nil {
				return err
			}
			return NewMLFlowExperimentAPI(ctx, c).Delete(ad)
		},
		Schema: s,
	}.ToResource()
}
