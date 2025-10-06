package jobs

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestDataSourceQueryableJobMatchesId(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockJobsAPI().EXPECT().Get(mock.Anything, jobs.GetJobRequest{JobId: 234}).Return(&jobs.Job{
				JobId: 234,
				Settings: &jobs.JobSettings{
					Name: "Second",
				},
			}, nil)
		},
		Resource:    DataSourceJob(),
		Read:        true,
		New:         true,
		NonWritable: true,
		HCL:         `job_id = "234"`,
		ID:          "234",
	}.ApplyAndExpectData(t, map[string]any{
		"job_id":                         "234",
		"id":                             "234",
		"job_settings.0.settings.0.name": "Second",
	})
}

func TestDataSourceQueryableJobRunAsSP(t *testing.T) {
	spID := "3f670caf-9a4b-4479-8143-1a0878da8f57"
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockJobsAPI().EXPECT().Get(mock.Anything, jobs.GetJobRequest{JobId: 234}).Return(&jobs.Job{
				JobId: 234,
				Settings: &jobs.JobSettings{
					Name: "Second",
					RunAs: &jobs.JobRunAs{
						ServicePrincipalName: spID,
					},
				},
				CreatorUserName: "user@domain.com",
				RunAsUserName:   spID,
			}, nil)
		},
		Resource:    DataSourceJob(),
		Read:        true,
		New:         true,
		NonWritable: true,
		HCL:         `job_id = "234"`,
		ID:          "234",
	}.ApplyAndExpectData(t, map[string]any{
		"job_id":                         "234",
		"id":                             "234",
		"job_settings.0.settings.0.name": "Second",
		"job_settings.0.settings.0.run_as.0.service_principal_name": spID,
	})
}

func TestDataSourceQueryableJobRunAsSameUser(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockJobsAPI().EXPECT().Get(mock.Anything, jobs.GetJobRequest{JobId: 234}).Return(&jobs.Job{
				JobId: 234,
				Settings: &jobs.JobSettings{
					Name: "Second",
					RunAs: &jobs.JobRunAs{
						UserName: "user@domain.com",
					},
				},
				CreatorUserName: "user@domain.com",
				RunAsUserName:   "user@domain.com",
			}, nil)
		},
		Resource:    DataSourceJob(),
		Read:        true,
		New:         true,
		NonWritable: true,
		HCL:         `job_id = "234"`,
		ID:          "234",
	}.ApplyAndExpectData(t, map[string]any{
		"job_id":                         "234",
		"id":                             "234",
		"job_settings.0.settings.0.name": "Second",
		"job_settings.0.settings.0.run_as.0.user_name": "user@domain.com",
	})
}

func TestDataSourceQueryableJobRunAsAnotherUser(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockJobsAPI().EXPECT().Get(mock.Anything, jobs.GetJobRequest{JobId: 234}).Return(&jobs.Job{
				JobId: 234,
				Settings: &jobs.JobSettings{
					Name: "Second",
					RunAs: &jobs.JobRunAs{
						UserName: "user2@domain.com",
					},
				},
				CreatorUserName: "user1@domain.com",
				RunAsUserName:   "user2@domain.com",
			}, nil)
		},
		Resource:    DataSourceJob(),
		Read:        true,
		New:         true,
		NonWritable: true,
		HCL:         `job_id = "234"`,
		ID:          "234",
	}.ApplyAndExpectData(t, map[string]any{
		"job_id":                         "234",
		"id":                             "234",
		"job_settings.0.settings.0.name": "Second",
		"job_settings.0.settings.0.run_as.0.user_name": "user2@domain.com",
	})
}

func TestDataSourceQueryableJobMatchesName(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockJobsAPI().EXPECT().ListAll(mock.Anything, jobs.ListJobsRequest{
				ExpandTasks: false,
				Name:        "First",
				Limit:       100,
			}).Return([]jobs.BaseJob{
				{
					JobId: 123,
					Settings: &jobs.JobSettings{
						Name: "First",
					},
				},
				{
					JobId: 234,
					Settings: &jobs.JobSettings{
						Name: "Second",
					},
				},
			}, nil)
			w.GetMockJobsAPI().EXPECT().Get(mock.Anything, jobs.GetJobRequest{JobId: 123}).Return(&jobs.Job{
				JobId: 123,
				Settings: &jobs.JobSettings{
					Name: "First",
				},
			}, nil)
		},
		Resource:    DataSourceJob(),
		Read:        true,
		NonWritable: true,
		HCL:         `job_name = "First"`,
		ID:          "123",
	}.ApplyAndExpectData(t, map[string]any{
		"job_id":                         "123",
		"id":                             "123",
		"job_settings.0.settings.0.name": "First",
	})
}

func TestDataSourceQueryableJobNoMatchName(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockJobsAPI().EXPECT().ListAll(mock.Anything, jobs.ListJobsRequest{
				ExpandTasks: false,
				Name:        "Third",
				Limit:       100,
			}).Return([]jobs.BaseJob{}, nil)
		},
		Resource:    DataSourceJob(),
		Read:        true,
		NonWritable: true,
		HCL:         `job_name = "Third"`,
		ID:          "_",
	}.ExpectError(t, "no job found with specified name")
}

func TestDataSourceQueryableJobNoMatchId(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			w.GetMockJobsAPI().EXPECT().Get(mock.Anything, jobs.GetJobRequest{JobId: 567}).Return(nil, &apierr.APIError{
				ErrorCode:  "RESOURCE_DOES_NOT_EXIST",
				Message:    "Job 567 does not exist.",
				StatusCode: 404,
			})
		},
		Resource:    DataSourceJob(),
		Read:        true,
		NonWritable: true,
		HCL:         `id = "567"`,
		ID:          "_",
	}.ExpectError(t, "Job 567 does not exist.")
}
