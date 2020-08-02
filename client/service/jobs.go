package service

import (
	"fmt"
	"strconv"
	"strings"

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
func (a JobsAPI) Update(id string, jobSettings model.JobSettings) error {
	jobID, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return err
	}
	return wrapMissingJobError(a.client.post("/jobs/reset", model.UpdateJobRequest{
		JobID:       jobID,
		NewSettings: &jobSettings,
	}, nil), id)
}

// Read returns the job object with all the attributes
func (a JobsAPI) Read(id string) (job model.Job, err error) {
	jobID, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return
	}
	err = wrapMissingJobError(a.client.get("/jobs/get", map[string]int64{
		"job_id": jobID,
	}, &job), id)
	return
}

// Delete deletes the job given a job id
func (a JobsAPI) Delete(id string) error {
	jobID, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return err
	}
	return wrapMissingJobError(a.client.post("/jobs/delete", map[string]int64{
		"job_id": jobID,
	}, nil), id)
}

func wrapMissingJobError(err error, id string) error {
	if err == nil {
		return nil
	}
	apiErr, ok := err.(APIError)
	if !ok {
		return err
	}
	if apiErr.IsMissing() {
		return err
	}
	// fix non-compliant error code
	if strings.Contains(apiErr.Message,
		fmt.Sprintf("Job %s does not exist.", id)) {
		apiErr.StatusCode = 404
		return err
	}
	return err
}
