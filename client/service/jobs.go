package service

import (
	"encoding/json"
	"net/http"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

// JobsAPI exposes the Jobs API
type JobsAPI struct {
	Client *DBApiClient
}

// Create creates a job on the workspace given the job settings
func (a JobsAPI) Create(jobSettings model.JobSettings) (model.Job, error) {
	var job model.Job
	resp, err := a.Client.performQuery(http.MethodPost, "/jobs/create", "2.0", nil, jobSettings, nil)
	if err != nil {
		return job, err
	}

	err = json.Unmarshal(resp, &job)
	return job, err
}

// Update updates a job given the id and a new set of job settings
func (a JobsAPI) Update(jobID int64, jobSettings model.JobSettings) error {
	jobResetRequest := struct {
		JobID       int64              `json:"job_id,omitempty" url:"job_id,omitempty"`
		NewSettings *model.JobSettings `json:"new_settings,omitempty" url:"new_settings,omitempty"`
	}{JobID: jobID, NewSettings: &jobSettings}
	_, err := a.Client.performQuery(http.MethodPost, "/jobs/reset", "2.0", nil, jobResetRequest, nil)
	return err
}

// Read returns the job object with all the attributes
func (a JobsAPI) Read(jobID int64) (model.Job, error) {
	jobGetRequest := struct {
		JobID int64 `json:"job_id,omitempty" url:"job_id,omitempty"`
	}{JobID: jobID}

	var job model.Job

	resp, err := a.Client.performQuery(http.MethodGet, "/jobs/get", "2.0", nil, jobGetRequest, nil)
	if err != nil {
		return job, err
	}

	err = json.Unmarshal(resp, &job)

	return job, err
}

// Delete deletes the job given a job id
func (a JobsAPI) Delete(jobID int64) error {
	jobDeleteRequest := struct {
		JobID int64 `json:"job_id,omitempty" url:"job_id,omitempty"`
	}{JobID: jobID}

	_, err := a.Client.performQuery(http.MethodPost, "/jobs/delete", "2.0", nil, jobDeleteRequest, nil)

	return err
}
