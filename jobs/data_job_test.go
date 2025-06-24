package jobs

import (
	"fmt"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func commonFixtures(name string) []qa.HTTPFixture {
	resource := "/api/2.1/jobs/list?expand_tasks=false&limit=25"
	if name != "" {
		resource = fmt.Sprintf("/api/2.1/jobs/list?expand_tasks=true&limit=25&name=%s", name)
	}
	return []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: resource,
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
		{
			Method:   "GET",
			Resource: resource,
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
	}

}
func TestDataSourceQueryableJobMatchesId(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/jobs/get?job_id=234",
				Response: Job{
					JobID: 234,
					Settings: &JobSettings{
						Name: "Second",
					},
				},
			},
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
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/jobs/get?job_id=234",
				Response: Job{
					JobID: 234,
					Settings: &JobSettings{
						Name: "Second",
					},
					CreatorUserName: "user@domain.com",
					RunAsUserName:   spID,
				},
			},
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
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/jobs/get?job_id=234",
				Response: Job{
					JobID: 234,
					Settings: &JobSettings{
						Name: "Second",
					},
					CreatorUserName: "user@domain.com",
					RunAsUserName:   "user@domain.com",
				},
			},
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
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/jobs/get?job_id=234",
				Response: Job{
					JobID: 234,
					Settings: &JobSettings{
						Name: "Second",
					},
					CreatorUserName: "user1@domain.com",
					RunAsUserName:   "user2@domain.com",
				},
			},
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
		Fixtures:    commonFixtures("First"),
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
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.1/jobs/list?expand_tasks=true&limit=25&name=Third",
				Response: JobListResponse{
					Jobs: []Job{},
				},
			},
		},
		Resource:    DataSourceJob(),
		Read:        true,
		NonWritable: true,
		HCL:         `job_name= "Third"`,
		ID:          "_",
	}.ExpectError(t, "no job found with specified name")
}

func TestDataSourceQueryableJobNoMatchId(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/jobs/get?job_id=567",
				Response: common.APIErrorBody{
					ErrorCode: "RESOURCE_DOES_NOT_EXIST",
					Message:   "Job 567 does not exist.",
				},
				Status: 400,
			},
		},
		Resource:    DataSourceJob(),
		Read:        true,
		NonWritable: true,
		HCL:         `id= "567"`,
		ID:          "_",
	}.ExpectError(t, "Job 567 does not exist.")
}
