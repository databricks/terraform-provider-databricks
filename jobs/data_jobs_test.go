package jobs

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestJobsData(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			iterator := listing.SliceIterator[jobs.BaseJob]{
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
			}
			w.GetMockJobsAPI().EXPECT().
				List(mock.Anything, jobs.ListJobsRequest{ExpandTasks: false, Limit: 100}).
				Return(&iterator)
		},
		Resource:    DataSourceJobs(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"ids": map[string]any{
			"First":  "123",
			"Second": "234",
		},
	})
}

func TestJobsDataWithFilter(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			iterator := listing.SliceIterator[jobs.BaseJob]{
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
			}
			w.GetMockJobsAPI().EXPECT().
				List(mock.Anything, jobs.ListJobsRequest{ExpandTasks: false, Limit: 100}).
				Return(&iterator)
		},
		Resource:    DataSourceJobs(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		HCL:         `job_name_contains = "Sec"`,
	}.ApplyAndExpectData(t, map[string]any{
		"ids": map[string]any{
			"Second": "234",
		},
	})
}

func TestJobsDataWithKeyID(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			iterator := listing.SliceIterator[jobs.BaseJob]{
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
			}
			w.GetMockJobsAPI().EXPECT().
				List(mock.Anything, jobs.ListJobsRequest{ExpandTasks: false, Limit: 100}).
				Return(&iterator)
		},
		Resource:    DataSourceJobs(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		HCL:         `key = "id"`,
	}.ApplyAndExpectData(t, map[string]any{
		"ids": map[string]any{
			"123": "123",
			"234": "234",
		},
	})
}

func TestJobsDataWithDuplicateNames(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			iterator := listing.SliceIterator[jobs.BaseJob]{
				{
					JobId: 123,
					Settings: &jobs.JobSettings{
						Name: "Duplicate",
					},
				},
				{
					JobId: 234,
					Settings: &jobs.JobSettings{
						Name: "Duplicate",
					},
				},
			}
			w.GetMockJobsAPI().EXPECT().
				List(mock.Anything, jobs.ListJobsRequest{ExpandTasks: false, Limit: 100}).
				Return(&iterator)
		},
		Resource:    DataSourceJobs(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "duplicate job name detected: Duplicate")
}

func TestJobsDataWithInvalidKey(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(w *mocks.MockWorkspaceClient) {
			iterator := listing.SliceIterator[jobs.BaseJob]{
				{
					JobId: 123,
					Settings: &jobs.JobSettings{
						Name: "First",
					},
				},
			}
			w.GetMockJobsAPI().EXPECT().
				List(mock.Anything, jobs.ListJobsRequest{ExpandTasks: false, Limit: 100}).
				Return(&iterator)
		},
		Resource:    DataSourceJobs(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		HCL:         `key = "invalid"`,
	}.ExpectError(t, "unsupported key invalid, must be one of name or id")
}
