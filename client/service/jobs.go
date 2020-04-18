package service

import (
	"encoding/json"
	"github.com/databrickslabs/databricks-terraform/client/model"
	"net/http"
)

// TokensAPI exposes the Secrets API
type JobsAPI struct {
	Client DBApiClient
}

func (a JobsAPI) Create(jobSettings model.JobSettings) (model.Job, error) {
	var job model.Job
	resp, err := a.Client.performQuery(http.MethodPost, "/jobs/create", "2.0", nil, jobSettings)
	if err != nil {
		return job, err
	}

	err = json.Unmarshal(resp, &job)
	return job, err
}

func (a JobsAPI) Update(jobId int64, jobSettings model.JobSettings) error {
	jobResetRequest := struct {
		JobId       int64              `json:"job_id,omitempty" url:"job_id,omitempty"`
		NewSettings *model.JobSettings `json:"new_settings,omitempty" url:"new_settings,omitempty"`
	}{JobId: jobId, NewSettings: &jobSettings}
	_, err := a.Client.performQuery(http.MethodPost, "/jobs/reset", "2.0", nil, jobResetRequest)
	return err
}

func (a JobsAPI) Read(jobId int64) (model.Job, error) {
	jobGetRequest := struct {
		JobId int64 `json:"job_id,omitempty" url:"job_id,omitempty"`
	}{JobId: jobId}

	var job model.Job

	resp, err := a.Client.performQuery(http.MethodGet, "/jobs/get", "2.0", nil, jobGetRequest)
	if err != nil {
		return job, err
	}

	err = json.Unmarshal(resp, &job)

	return job, err
}

func (a JobsAPI) Delete(jobId int64) error {
	jobDeleteRequest := struct {
		JobId int64 `json:"job_id,omitempty" url:"job_id,omitempty"`
	}{JobId: jobId}

	_, err := a.Client.performQuery(http.MethodPost, "/jobs/delete", "2.0", nil, jobDeleteRequest)

	return err
}
