package exporter

import (
	"context"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"

	sdk_jobs "github.com/databricks/databricks-sdk-go/service/jobs"
	tf_jobs "github.com/databricks/terraform-provider-databricks/jobs"
)

func TestJobsIgnore(t *testing.T) {
	ic := importContextForTest()
	d := tf_jobs.ResourceJob().ToResource().TestResourceData()
	d.SetId("12345")
	r := &resource{ID: "12345", Data: d}
	// job without tasks
	assert.True(t, resourcesMap["databricks_job"].Ignore(ic, r))
	assert.Equal(t, 1, len(ic.ignoredResources))
}

func TestJobName(t *testing.T) {
	ic := importContextForTest()
	d := tf_jobs.ResourceJob().ToResource().TestResourceData()
	d.SetId("12345")
	// job without name
	assert.Equal(t, "job_12345", resourcesMap["databricks_job"].Name(ic, d))
	// job with name
	d.Set("name", "test@1pm")
	assert.Equal(t, "test_1pm_12345", resourcesMap["databricks_job"].Name(ic, d))
}

func TestJobListNoNameMatchOrFromBundles(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.2/jobs/list?limit=100",
			Response: sdk_jobs.ListJobsResponse{
				Jobs: []sdk_jobs.BaseJob{
					{
						Settings: &sdk_jobs.JobSettings{
							Name: "abc",
						},
					},
					{
						Settings: &sdk_jobs.JobSettings{
							Name:     "bcd",
							EditMode: "UI_LOCKED",
							Deployment: &sdk_jobs.JobDeployment{
								Kind: "BUNDLE",
							},
						},
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("jobs")
		ic.match = "bcd"
		err := resourcesMap["databricks_job"].List(ic)
		assert.NoError(t, err)
		assert.Equal(t, 0, len(ic.testEmits))
	})
}
