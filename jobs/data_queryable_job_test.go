package jobs

import (
	"github.com/databricks/terraform-provider-databricks/qa"
	"testing"
)

var fixtures = []qa.HTTPFixture{
	{
		Method:   "GET",
		Resource: "/api/2.1/jobs/list",
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
}

func TestDataSourceQueryableJobMatchesId(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    fixtures,
		Resource:    DataSourceQueryableJob(),
		Read:        true,
		NonWritable: true,
		HCL:         `job_id = "234"`,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"ids": []string{
			"234",
		},
	})
}

func TestDataSourceQueryableJobMatchesName(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    fixtures,
		Resource:    DataSourceQueryableJob(),
		Read:        true,
		NonWritable: true,
		HCL:         `job_name= "First"`,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"ids": []string{
			"123",
		},
	})
}

func TestDataSourceQueryableJobNoMatchName(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    fixtures,
		Resource:    DataSourceQueryableJob(),
		Read:        true,
		NonWritable: true,
		HCL:         `job_name= "Third"`,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"ids": []string{},
	})
}

func TestDataSourceQueryableJobNoMatchId(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    fixtures,
		Resource:    DataSourceQueryableJob(),
		Read:        true,
		NonWritable: true,
		HCL:         `job_id= "567"`,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"ids": []string{},
	})
}
