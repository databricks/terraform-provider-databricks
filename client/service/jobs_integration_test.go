package service

import (
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/stretchr/testify/assert"
)

func TestJobsCreate(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client := GetIntegrationDBAPIClient()

	jobSettings := model.JobSettings{
		NewCluster: &model.Cluster{
			NumWorkers:   2,
			SparkVersion: "6.4.x-scala2.11",
			SparkConf:    nil,
			AwsAttributes: &model.AwsAttributes{
				Availability: "ON_DEMAND",
			},
			NodeTypeID: "r3.xlarge",
		},
		NotebookTask: &model.NotebookTask{
			NotebookPath: "/Users/sri.tikkireddy@databricks.com/demo-terraform/demo-notebook",
		},
		Name: "1-sri-test-job",
		Libraries: []model.Library{
			{
				Maven: &model.Maven{
					Coordinates: "org.jsoup:jsoup:1.7.2",
				},
			},
		},
		EmailNotifications: &model.JobEmailNotifications{
			OnStart:   []string{},
			OnSuccess: []string{},
			OnFailure: []string{},
		},
		TimeoutSeconds: 3600,
		MaxRetries:     1,
		Schedule: &model.CronSchedule{
			QuartzCronExpression: "0 15 22 ? * *",
			TimezoneID:           "America/Los_Angeles",
		},
		MaxConcurrentRuns: 1,
	}

	job, err := client.Jobs().Create(jobSettings)
	assert.NoError(t, err, err)
	id := job.JobID
	defer func() {
		err := client.Jobs().Delete(id)
		assert.NoError(t, err, err)
	}()
	t.Log(id)
	job, err = client.Jobs().Read(id)
	assert.NoError(t, err, err)
	assert.True(t, job.Settings.NewCluster.SparkVersion == "6.4.x-scala2.11",
		"Something is wrong with spark version")

	jobSettings.NewCluster.SparkVersion = "6.1.x-scala2.11"

	err = client.Jobs().Update(id, jobSettings)
	assert.NoError(t, err, err)

	job, err = client.Jobs().Read(id)
	assert.NoError(t, err, err)
	assert.True(t, job.Settings.NewCluster.SparkVersion == "6.1.x-scala2.11",
		"Something is wrong with spark version")
}
