package jobs

import (
	"github.com/databricks/terraform-provider-databricks/qa"
	"testing"
)

func commonFixtures() []qa.HTTPFixture {
	return []qa.HTTPFixture{
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
	}

}
func TestDataSourceQueryableJobMatchesId(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    commonFixtures(),
		Resource:    DataSourceJob(),
		Read:        true,
		New:         true,
		NonWritable: true,
		HCL:         `job_id = "234"`,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"job_id":                         "234",
		"job_settings.0.settings.0.name": "Second",
	})
}

func TestDataSourceQueryableJobMatchesName(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    commonFixtures(),
		Resource:    DataSourceJob(),
		Read:        true,
		NonWritable: true,
		HCL:         `job_name = "First"`,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"job_id":                         "123",
		"job_settings.0.settings.0.name": "First",
	})
}

func TestDataSourceQueryableJobNoMatchName(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    commonFixtures(),
		Resource:    DataSourceJob(),
		Read:        true,
		NonWritable: true,
		HCL:         `job_name= "Third"`,
		ID:          "_",
	}.ExpectError(t, "no job found with specified name or id")
}

func TestDataSourceQueryableJobNoMatchId(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    commonFixtures(),
		Resource:    DataSourceJob(),
		Read:        true,
		NonWritable: true,
		HCL:         `job_id= "567"`,
		ID:          "_",
	}.ExpectError(t, "no job found with specified name or id")
}
