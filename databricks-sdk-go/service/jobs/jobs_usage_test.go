// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package jobs_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/databricks-sdk-go/service/jobs"
)

func ExampleJobsAPI_CancelAllRuns_jobsApiFullIntegration() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	notebookPath := func() string {
		me, err := w.CurrentUser.Me(ctx, iam.MeRequest{})
		if err != nil {
			panic(err)
		}
		return filepath.Join("/Users", me.UserName, fmt.Sprintf("sdk-%x", time.Now().UnixNano()))
	}()

	clusterId := func() string {
		clusterId := os.Getenv("DATABRICKS_CLUSTER_ID")
		err := w.Clusters.EnsureClusterIsRunning(ctx, clusterId)
		if err != nil {
			panic(err)
		}
		return clusterId
	}()

	createdJob, err := w.Jobs.Create(ctx, jobs.CreateJob{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Tasks: []jobs.Task{{
			Description:       "test",
			ExistingClusterId: clusterId,
			NotebookTask: &jobs.NotebookTask{
				NotebookPath: notebookPath,
			},
			TaskKey:        "test",
			TimeoutSeconds: 0,
		}},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", createdJob)

	err = w.Jobs.CancelAllRuns(ctx, jobs.CancelAllRuns{
		JobId: createdJob.JobId,
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Jobs.DeleteByJobId(ctx, createdJob.JobId)
	if err != nil {
		panic(err)
	}

}

func ExampleJobsAPI_CancelRun_jobsApiFullIntegration() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	notebookPath := func() string {
		me, err := w.CurrentUser.Me(ctx, iam.MeRequest{})
		if err != nil {
			panic(err)
		}
		return filepath.Join("/Users", me.UserName, fmt.Sprintf("sdk-%x", time.Now().UnixNano()))
	}()

	clusterId := func() string {
		clusterId := os.Getenv("DATABRICKS_CLUSTER_ID")
		err := w.Clusters.EnsureClusterIsRunning(ctx, clusterId)
		if err != nil {
			panic(err)
		}
		return clusterId
	}()

	createdJob, err := w.Jobs.Create(ctx, jobs.CreateJob{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Tasks: []jobs.Task{{
			Description:       "test",
			ExistingClusterId: clusterId,
			NotebookTask: &jobs.NotebookTask{
				NotebookPath: notebookPath,
			},
			TaskKey:        "test",
			TimeoutSeconds: 0,
		}},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", createdJob)

	runNowResponse, err := w.Jobs.RunNow(ctx, jobs.RunNow{
		JobId: createdJob.JobId,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", runNowResponse)

	cancelledRun, err := w.Jobs.CancelRunAndWait(ctx, jobs.CancelRun{
		RunId: runNowResponse.Response.RunId,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", cancelledRun)

	// cleanup

	err = w.Jobs.DeleteByJobId(ctx, createdJob.JobId)
	if err != nil {
		panic(err)
	}

}

func ExampleJobsAPI_Create_jobsApiFullIntegration() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	notebookPath := func() string {
		me, err := w.CurrentUser.Me(ctx, iam.MeRequest{})
		if err != nil {
			panic(err)
		}
		return filepath.Join("/Users", me.UserName, fmt.Sprintf("sdk-%x", time.Now().UnixNano()))
	}()

	clusterId := func() string {
		clusterId := os.Getenv("DATABRICKS_CLUSTER_ID")
		err := w.Clusters.EnsureClusterIsRunning(ctx, clusterId)
		if err != nil {
			panic(err)
		}
		return clusterId
	}()

	createdJob, err := w.Jobs.Create(ctx, jobs.CreateJob{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Tasks: []jobs.Task{{
			Description:       "test",
			ExistingClusterId: clusterId,
			NotebookTask: &jobs.NotebookTask{
				NotebookPath: notebookPath,
			},
			TaskKey:        "test",
			TimeoutSeconds: 0,
		}},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", createdJob)

	// cleanup

	err = w.Jobs.DeleteByJobId(ctx, createdJob.JobId)
	if err != nil {
		panic(err)
	}

}

func ExampleJobsAPI_ExportRun_jobsApiFullIntegration() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	notebookPath := func() string {
		me, err := w.CurrentUser.Me(ctx, iam.MeRequest{})
		if err != nil {
			panic(err)
		}
		return filepath.Join("/Users", me.UserName, fmt.Sprintf("sdk-%x", time.Now().UnixNano()))
	}()

	clusterId := func() string {
		clusterId := os.Getenv("DATABRICKS_CLUSTER_ID")
		err := w.Clusters.EnsureClusterIsRunning(ctx, clusterId)
		if err != nil {
			panic(err)
		}
		return clusterId
	}()

	createdJob, err := w.Jobs.Create(ctx, jobs.CreateJob{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Tasks: []jobs.Task{{
			Description:       "test",
			ExistingClusterId: clusterId,
			NotebookTask: &jobs.NotebookTask{
				NotebookPath: notebookPath,
			},
			TaskKey:        "test",
			TimeoutSeconds: 0,
		}},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", createdJob)

	runById, err := w.Jobs.RunNowAndWait(ctx, jobs.RunNow{
		JobId: createdJob.JobId,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", runById)

	exportedView, err := w.Jobs.ExportRun(ctx, jobs.ExportRunRequest{
		RunId:         runById.Tasks[0].RunId,
		ViewsToExport: "CODE",
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", exportedView)

	// cleanup

	err = w.Jobs.DeleteByJobId(ctx, createdJob.JobId)
	if err != nil {
		panic(err)
	}

}

func ExampleJobsAPI_Get_jobsApiFullIntegration() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	notebookPath := func() string {
		me, err := w.CurrentUser.Me(ctx, iam.MeRequest{})
		if err != nil {
			panic(err)
		}
		return filepath.Join("/Users", me.UserName, fmt.Sprintf("sdk-%x", time.Now().UnixNano()))
	}()

	clusterId := func() string {
		clusterId := os.Getenv("DATABRICKS_CLUSTER_ID")
		err := w.Clusters.EnsureClusterIsRunning(ctx, clusterId)
		if err != nil {
			panic(err)
		}
		return clusterId
	}()

	createdJob, err := w.Jobs.Create(ctx, jobs.CreateJob{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Tasks: []jobs.Task{{
			Description:       "test",
			ExistingClusterId: clusterId,
			NotebookTask: &jobs.NotebookTask{
				NotebookPath: notebookPath,
			},
			TaskKey:        "test",
			TimeoutSeconds: 0,
		}},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", createdJob)

	byId, err := w.Jobs.GetByJobId(ctx, createdJob.JobId)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", byId)

	// cleanup

	err = w.Jobs.DeleteByJobId(ctx, createdJob.JobId)
	if err != nil {
		panic(err)
	}

}

func ExampleJobsAPI_GetRunOutput_jobsApiFullIntegration() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	notebookPath := func() string {
		me, err := w.CurrentUser.Me(ctx, iam.MeRequest{})
		if err != nil {
			panic(err)
		}
		return filepath.Join("/Users", me.UserName, fmt.Sprintf("sdk-%x", time.Now().UnixNano()))
	}()

	clusterId := func() string {
		clusterId := os.Getenv("DATABRICKS_CLUSTER_ID")
		err := w.Clusters.EnsureClusterIsRunning(ctx, clusterId)
		if err != nil {
			panic(err)
		}
		return clusterId
	}()

	run, err := w.Jobs.SubmitAndWait(ctx, jobs.SubmitRun{
		RunName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Tasks: []jobs.SubmitTask{{
			ExistingClusterId: clusterId,
			NotebookTask: &jobs.NotebookTask{
				NotebookPath: notebookPath,
			},
			TaskKey: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		}},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", run)

	output, err := w.Jobs.GetRunOutputByRunId(ctx, run.Tasks[0].RunId)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", output)

	// cleanup

	err = w.Jobs.DeleteRunByRunId(ctx, run.RunId)
	if err != nil {
		panic(err)
	}

}

func ExampleJobsAPI_List_jobsApiFullIntegration() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	jobList, err := w.Jobs.ListAll(ctx, jobs.ListJobsRequest{
		ExpandTasks: false,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", jobList)

}

func ExampleJobsAPI_ListRuns_jobsApiFullIntegration() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	notebookPath := func() string {
		me, err := w.CurrentUser.Me(ctx, iam.MeRequest{})
		if err != nil {
			panic(err)
		}
		return filepath.Join("/Users", me.UserName, fmt.Sprintf("sdk-%x", time.Now().UnixNano()))
	}()

	clusterId := func() string {
		clusterId := os.Getenv("DATABRICKS_CLUSTER_ID")
		err := w.Clusters.EnsureClusterIsRunning(ctx, clusterId)
		if err != nil {
			panic(err)
		}
		return clusterId
	}()

	createdJob, err := w.Jobs.Create(ctx, jobs.CreateJob{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Tasks: []jobs.Task{{
			Description:       "test",
			ExistingClusterId: clusterId,
			NotebookTask: &jobs.NotebookTask{
				NotebookPath: notebookPath,
			},
			TaskKey:        "test",
			TimeoutSeconds: 0,
		}},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", createdJob)

	runList, err := w.Jobs.ListRunsAll(ctx, jobs.ListRunsRequest{
		JobId: createdJob.JobId,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", runList)

	// cleanup

	err = w.Jobs.DeleteByJobId(ctx, createdJob.JobId)
	if err != nil {
		panic(err)
	}

}

func ExampleJobsAPI_RepairRun_jobsApiFullIntegration() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	notebookPath := func() string {
		me, err := w.CurrentUser.Me(ctx, iam.MeRequest{})
		if err != nil {
			panic(err)
		}
		return filepath.Join("/Users", me.UserName, fmt.Sprintf("sdk-%x", time.Now().UnixNano()))
	}()

	clusterId := func() string {
		clusterId := os.Getenv("DATABRICKS_CLUSTER_ID")
		err := w.Clusters.EnsureClusterIsRunning(ctx, clusterId)
		if err != nil {
			panic(err)
		}
		return clusterId
	}()

	createdJob, err := w.Jobs.Create(ctx, jobs.CreateJob{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Tasks: []jobs.Task{{
			Description:       "test",
			ExistingClusterId: clusterId,
			NotebookTask: &jobs.NotebookTask{
				NotebookPath: notebookPath,
			},
			TaskKey:        "test",
			TimeoutSeconds: 0,
		}},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", createdJob)

	runNowResponse, err := w.Jobs.RunNow(ctx, jobs.RunNow{
		JobId: createdJob.JobId,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", runNowResponse)

	cancelledRun, err := w.Jobs.CancelRunAndWait(ctx, jobs.CancelRun{
		RunId: runNowResponse.Response.RunId,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", cancelledRun)

	repairedRun, err := w.Jobs.RepairRunAndWait(ctx, jobs.RepairRun{
		RerunTasks: []string{cancelledRun.Tasks[0].TaskKey},
		RunId:      runNowResponse.Response.RunId,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", repairedRun)

	// cleanup

	err = w.Jobs.DeleteByJobId(ctx, createdJob.JobId)
	if err != nil {
		panic(err)
	}

}

func ExampleJobsAPI_Reset_jobsApiFullIntegration() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	notebookPath := func() string {
		me, err := w.CurrentUser.Me(ctx, iam.MeRequest{})
		if err != nil {
			panic(err)
		}
		return filepath.Join("/Users", me.UserName, fmt.Sprintf("sdk-%x", time.Now().UnixNano()))
	}()

	clusterId := func() string {
		clusterId := os.Getenv("DATABRICKS_CLUSTER_ID")
		err := w.Clusters.EnsureClusterIsRunning(ctx, clusterId)
		if err != nil {
			panic(err)
		}
		return clusterId
	}()

	createdJob, err := w.Jobs.Create(ctx, jobs.CreateJob{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Tasks: []jobs.Task{{
			Description:       "test",
			ExistingClusterId: clusterId,
			NotebookTask: &jobs.NotebookTask{
				NotebookPath: notebookPath,
			},
			TaskKey:        "test",
			TimeoutSeconds: 0,
		}},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", createdJob)

	newName := fmt.Sprintf("sdk-%x", time.Now().UnixNano())

	byId, err := w.Jobs.GetByJobId(ctx, createdJob.JobId)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", byId)

	err = w.Jobs.Reset(ctx, jobs.ResetJob{
		JobId: byId.JobId,
		NewSettings: jobs.JobSettings{
			Name:  newName,
			Tasks: byId.Settings.Tasks,
		},
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Jobs.DeleteByJobId(ctx, createdJob.JobId)
	if err != nil {
		panic(err)
	}

}

func ExampleJobsAPI_RunNow_jobsApiFullIntegration() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	notebookPath := func() string {
		me, err := w.CurrentUser.Me(ctx, iam.MeRequest{})
		if err != nil {
			panic(err)
		}
		return filepath.Join("/Users", me.UserName, fmt.Sprintf("sdk-%x", time.Now().UnixNano()))
	}()

	clusterId := func() string {
		clusterId := os.Getenv("DATABRICKS_CLUSTER_ID")
		err := w.Clusters.EnsureClusterIsRunning(ctx, clusterId)
		if err != nil {
			panic(err)
		}
		return clusterId
	}()

	createdJob, err := w.Jobs.Create(ctx, jobs.CreateJob{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Tasks: []jobs.Task{{
			Description:       "test",
			ExistingClusterId: clusterId,
			NotebookTask: &jobs.NotebookTask{
				NotebookPath: notebookPath,
			},
			TaskKey:        "test",
			TimeoutSeconds: 0,
		}},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", createdJob)

	runById, err := w.Jobs.RunNowAndWait(ctx, jobs.RunNow{
		JobId: createdJob.JobId,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", runById)

	// cleanup

	err = w.Jobs.DeleteByJobId(ctx, createdJob.JobId)
	if err != nil {
		panic(err)
	}

}

func ExampleJobsAPI_Submit_jobsApiFullIntegration() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	notebookPath := func() string {
		me, err := w.CurrentUser.Me(ctx, iam.MeRequest{})
		if err != nil {
			panic(err)
		}
		return filepath.Join("/Users", me.UserName, fmt.Sprintf("sdk-%x", time.Now().UnixNano()))
	}()

	clusterId := func() string {
		clusterId := os.Getenv("DATABRICKS_CLUSTER_ID")
		err := w.Clusters.EnsureClusterIsRunning(ctx, clusterId)
		if err != nil {
			panic(err)
		}
		return clusterId
	}()

	run, err := w.Jobs.SubmitAndWait(ctx, jobs.SubmitRun{
		RunName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Tasks: []jobs.SubmitTask{{
			ExistingClusterId: clusterId,
			NotebookTask: &jobs.NotebookTask{
				NotebookPath: notebookPath,
			},
			TaskKey: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		}},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", run)

	// cleanup

	err = w.Jobs.DeleteRunByRunId(ctx, run.RunId)
	if err != nil {
		panic(err)
	}

}

func ExampleJobsAPI_Update_jobsApiFullIntegration() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	notebookPath := func() string {
		me, err := w.CurrentUser.Me(ctx, iam.MeRequest{})
		if err != nil {
			panic(err)
		}
		return filepath.Join("/Users", me.UserName, fmt.Sprintf("sdk-%x", time.Now().UnixNano()))
	}()

	clusterId := func() string {
		clusterId := os.Getenv("DATABRICKS_CLUSTER_ID")
		err := w.Clusters.EnsureClusterIsRunning(ctx, clusterId)
		if err != nil {
			panic(err)
		}
		return clusterId
	}()

	newName := fmt.Sprintf("sdk-%x", time.Now().UnixNano())

	createdJob, err := w.Jobs.Create(ctx, jobs.CreateJob{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Tasks: []jobs.Task{{
			Description:       "test",
			ExistingClusterId: clusterId,
			NotebookTask: &jobs.NotebookTask{
				NotebookPath: notebookPath,
			},
			TaskKey:        "test",
			TimeoutSeconds: 0,
		}},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", createdJob)

	err = w.Jobs.Update(ctx, jobs.UpdateJob{
		JobId: createdJob.JobId,
		NewSettings: &jobs.JobSettings{
			Name:              newName,
			MaxConcurrentRuns: 5,
		},
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Jobs.DeleteByJobId(ctx, createdJob.JobId)
	if err != nil {
		panic(err)
	}

}
