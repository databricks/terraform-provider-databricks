package service

import (
	"net/http"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

func TestJobsAPI_Create(t *testing.T) {
	type args model.JobSettings

	tests := []struct {
		name           string
		response       string
		responseStatus int
		args           args
		want           interface{}
		wantErr        bool
	}{
		{
			name: "Create test",
			response: `{
						  "job_id": 1
						}`,
			responseStatus: http.StatusOK,
			args: args{
				ExistingClusterID: "my-cluster-id",
			},
			want: model.Job{
				JobID: 1,
			},
			wantErr: false,
		},
		{
			name:           "Create faulure test",
			response:       "",
			responseStatus: http.StatusBadRequest,
			args: args{
				ExistingClusterID: "my-cluster-id",
			},
			want:    model.Job{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/jobs/create", &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return client.Jobs().Create(model.JobSettings(tt.args))
			})
		})
	}
}

func TestJobsAPI_Update(t *testing.T) {
	type args struct {
		JobID       int64             `json:"job_id,omitempty" url:"job_id,omitempty"`
		NewSettings model.JobSettings `json:"new_settings,omitempty" url:"new_settings,omitempty"`
	}

	tests := []struct {
		name           string
		response       string
		responseStatus int
		args           args
		want           interface{}
		wantErr        bool
	}{
		{
			name: "Update test",
			response: `{
						  "job_id": 1
						}`,
			responseStatus: http.StatusOK,
			args: args{
				JobID:       1,
				NewSettings: model.JobSettings{},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name:           "Update faulure test",
			response:       "",
			responseStatus: http.StatusBadRequest,
			args: args{
				JobID:       0,
				NewSettings: model.JobSettings{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/jobs/reset", &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return nil, client.Jobs().Update(tt.args.JobID, tt.args.NewSettings)
			})
		})
	}
}

func TestJobsAPI_Delete(t *testing.T) {
	type args struct {
		JobID int64 `json:"job_id,omitempty" url:"job_id,omitempty"`
	}

	tests := []struct {
		name           string
		response       string
		responseStatus int
		args           args
		wantErr        bool
	}{
		{
			name:           "Delete Test",
			response:       "",
			responseStatus: http.StatusOK,
			args: args{
				JobID: 0,
			},
			wantErr: false,
		},
		{
			name:           "Delete failure Test",
			response:       "",
			responseStatus: http.StatusBadRequest,
			args: args{
				JobID: 0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/jobs/delete", &input, tt.response, tt.responseStatus, nil, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return nil, client.Jobs().Delete(tt.args.JobID)
			})
		})
	}
}

func TestJobsAPI_Read(t *testing.T) {
	type args struct {
		JobID int64 `json:"job_id,omitempty" url:"job_id,omitempty"`
	}
	tests := []struct {
		name           string
		response       string
		responseStatus int
		args           args
		wantURI        string
		want           interface{}
		wantErr        bool
	}{
		{
			name: "Read test",
			response: `{
						  "job_id": 1,
						  "settings": {
							"name": "Nightly model training",
							"new_cluster": {
							  "spark_version": "5.3.x-scala2.11",
							  "node_type_id": "r3.xlarge",
							  "aws_attributes": {
								"availability": "ON_DEMAND"
							  },
							  "num_workers": 10
							},
							"libraries": [
							  {
								"jar": "dbfs:/my-jar.jar"
							  },
							  {
								"maven": {
								  "coordinates": "org.jsoup:jsoup:1.7.2"
								}
							  }
							],
							"email_notifications": {
							  "on_start": [],
							  "on_success": [],
							  "on_failure": []
							},
							"timeout_seconds": 100000000,
							"max_retries": 1,
							"schedule": {
							  "quartz_cron_expression": "0 15 22 ? * *",
							  "timezone_id": "America/Los_Angeles"
							},
							"spark_jar_task": {
							  "main_class_name": "com.databricks.ComputeModels"
							}
						  },
						  "created_time": 1457570074236
						}`,
			responseStatus: http.StatusOK,
			args: args{
				JobID: 1,
			},
			wantURI: "/api/2.0/jobs/get?job_id=1",
			want: model.Job{
				JobID: 1,
				Settings: &model.JobSettings{
					NewCluster: &model.Cluster{
						NumWorkers:   10,
						SparkVersion: "5.3.x-scala2.11",
						AwsAttributes: &model.AwsAttributes{
							Availability: "ON_DEMAND",
						},
						NodeTypeID: "r3.xlarge",
					},
					SparkJarTask: &model.SparkJarTask{
						MainClassName: "com.databricks.ComputeModels",
					},
					Name: "Nightly model training",
					Libraries: []model.Library{
						{
							Jar: "dbfs:/my-jar.jar",
						},
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
					TimeoutSeconds: 100000000,
					MaxRetries:     1,
					Schedule: &model.CronSchedule{
						QuartzCronExpression: "0 15 22 ? * *",
						TimezoneID:           "America/Los_Angeles",
					},
				},
				CreatedTime: 1457570074236,
			},
			wantErr: false,
		},
		{
			name:           "Read failure test",
			response:       ``,
			responseStatus: http.StatusBadRequest,
			args: args{
				JobID: 1,
			},
			wantURI: "/api/2.0/jobs/get?job_id=1",
			want:    model.Job{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, tt.args, http.MethodGet, tt.wantURI, &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return client.Jobs().Read(tt.args.JobID)
			})
		})
	}
}
