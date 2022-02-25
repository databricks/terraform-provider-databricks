package jobs

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
)

func TestJobsData(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/jobs/list",
				Response: JobList{
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
	}.ApplyAndExpectData(t, map[string]interface{}{
		"ids": map[string]interface{}{
			"First":  "123",
			"Second": "234",
		},
	})
}
