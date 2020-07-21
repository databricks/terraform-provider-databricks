package service

import (
	"github.com/databrickslabs/databricks-terraform/client/model"
)

// JobsAPI exposes the Jobs API
type JobsAPI struct {
	client *DatabricksClient
}

// Create creates a job on the workspace given the job settings
func (a JobsAPI) Create(jobSettings model.JobSettings) (model.Job, error) {
	var job model.Job
	err := a.client.post("/jobs/create", jobSettings, &job)
	return job, err
}

// Update updates a job given the id and a new set of job settings
func (a JobsAPI) Update(jobID int64, jobSettings model.JobSettings) error {
	jobResetRequest := struct {
		JobID       int64              `json:"job_id,omitempty" url:"job_id,omitempty"`
		NewSettings *model.JobSettings `json:"new_settings,omitempty" url:"new_settings,omitempty"`
	}{JobID: jobID, NewSettings: &jobSettings}
	return a.client.post("/jobs/reset", jobResetRequest, nil)
}

// Read returns the job object with all the attributes
func (a JobsAPI) Read(jobID int64) (model.Job, error) {
	err := a.client.get("/jobs/get", map[string]int64{
		"job_id": jobID,
	}, &job)
	return job, err
}

// Delete deletes the job given a job id
func (a JobsAPI) Delete(jobID int64) error {
	return a.client.post("/jobs/delete", map[string]int64{
		"job_id": jobID,
	}, nil)
}
