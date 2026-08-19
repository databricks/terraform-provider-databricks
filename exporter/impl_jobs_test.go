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

func TestJobsIgnoreManagedDeployment(t *testing.T) {
	ic := importContextForTest()
	d := tf_jobs.ResourceJob().ToResource().TestResourceData()
	d.SetId("12345")
	assert.NoError(t, d.Set("name", "system job"))
	assert.NoError(t, d.Set("task", []any{
		map[string]any{
			"task_key": "task",
			"notebook_task": []any{
				map[string]any{"notebook_path": "/Test"},
			},
		},
	}))
	assert.NoError(t, d.Set("deployment", []any{
		map[string]any{"kind": "SYSTEM_MANAGED"},
	}))
	r := &resource{ID: "12345", Data: d}

	assert.True(t, resourcesMap["databricks_job"].Ignore(ic, r))
	assert.Contains(t, ic.ignoredResources, "databricks_job. id=12345")
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

func TestImportTaskAlertTaskEmitsDependencies(t *testing.T) {
	ic := importContextForTest()
	ic.enableServices("jobs,alerts,sql-endpoints,settings,users")
	importTask(ic, sdk_jobs.Task{
		AlertTask: &sdk_jobs.AlertTask{
			AlertId:     "alert-id-1",
			WarehouseId: "warehouse-1",
			Subscribers: []sdk_jobs.AlertTaskSubscriber{
				{UserName: "subscriber@example.com"},
				{DestinationId: "nd-destination-1"},
			},
		},
	}, "test_job", "99")
	assert.Len(t, ic.testEmits, 4)
	assert.Contains(t, ic.testEmits, "databricks_alert_v2[<unknown>] (id: alert-id-1)")
	assert.Contains(t, ic.testEmits, "databricks_sql_endpoint[<unknown>] (id: warehouse-1)")
	assert.Contains(t, ic.testEmits, "databricks_user[<unknown>] (user_name: subscriber@example.com)")
	assert.Contains(t, ic.testEmits, "databricks_notification_destination[<unknown>] (id: nd-destination-1)")
}

func TestJobDependenciesIncludeTaskParameterReferences(t *testing.T) {
	dependencies := createJobDependencies()

	for _, expected := range []reference{
		{Path: "task.python_wheel_task.parameters", Resource: "databricks_file"},
		{Path: "task.run_job_task.job_parameters", Resource: "databricks_file"},
		{Path: "task.spark_python_task.parameters", Resource: "databricks_file"},
		{Path: "task.spark_python_task.parameters", Resource: "databricks_workspace_file", Match: "workspace_path"},
		{Path: "task.spark_python_task.parameters", Resource: "databricks_repo", Match: "workspace_path",
			MatchType: MatchPrefix, SearchValueTransformFunc: appendEndingSlashToDirName},
		{Path: "task.for_each_task.task.spark_python_task.parameters", Resource: "databricks_file"},
	} {
		assert.True(t, hasJobDependency(dependencies, expected), "missing dependency: %#v", expected)
	}
}

func hasJobDependency(dependencies []reference, expected reference) bool {
	for _, dependency := range dependencies {
		if dependency.Path == expected.Path &&
			dependency.Resource == expected.Resource &&
			dependency.Match == expected.Match &&
			dependency.MatchType == expected.MatchType {
			return true
		}
	}
	return false
}

func TestJobListNoNameMatchOrManaged(t *testing.T) {
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
							EditMode: sdk_jobs.JobEditModeUiLocked,
							Deployment: &sdk_jobs.JobDeployment{
								Kind: sdk_jobs.JobDeploymentKindBundle,
							},
						},
					},
					{
						Settings: &sdk_jobs.JobSettings{
							Name: "bcd",
							Deployment: &sdk_jobs.JobDeployment{
								Kind: sdk_jobs.JobDeploymentKindSystemManaged,
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
