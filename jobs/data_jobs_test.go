package jobs

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestJobsData(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/jobs/list?limit=100",
				Response: JobListResponse{
					Jobs: []Job{
						{
							JobID: 123,
							Settings: &JobSettings{
								Name: "First",
							},
						},
						{
							JobID: 234,
							Settings: &JobSettings{
								Name: "Second",
							},
						},
					},
				},
			},
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
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/jobs/list?limit=100",
				Response: JobListResponse{
					Jobs: []Job{
						{
							JobID: 123,
							Settings: &JobSettings{
								Name: "First",
							},
						},
						{
							JobID: 234,
							Settings: &JobSettings{
								Name: "Second",
							},
						},
					},
				},
			},
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
