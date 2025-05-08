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
