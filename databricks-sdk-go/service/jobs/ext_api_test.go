package jobs

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/qa"

	"github.com/stretchr/testify/assert"
)

func TestGetRun(t *testing.T) {
	ctx := context.Background()

	t.Run("run with no pagination", func(t *testing.T) {
		var requestMocks qa.HTTPFixtures = []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/runs/get?run_id=514594995218126",
				Response: Run{
					Iterations: []RunTask{},
					Tasks: []RunTask{
						{
							RunId:   123,
							TaskKey: "task1",
						},
						{
							RunId:   1234,
							TaskKey: "task2",
						},
					},
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.1/jobs/runs/get?run_id=514594995218126",
				Response: Run{
					Iterations: []RunTask{},
					Tasks: []RunTask{
						{
							RunId:   123,
							TaskKey: "task1",
						},
						{
							RunId:   1234,
							TaskKey: "task2",
						},
					},
				},
			},
		}
		client, server := requestMocks.Client(t)
		defer server.Close()

		mockJobsImpl := &jobsImpl{
			client: client,
		}
		api := &JobsAPI{jobsImpl: *mockJobsImpl}

		request := GetRunRequest{RunId: 514594995218126}
		run, err := api.GetRun(ctx, request)

		assert.NoError(t, err)
		assert.Equal(t, 2, len(run.Tasks))
		assert.Empty(t, run.Iterations)
		assert.EqualValues(t, 123, run.Tasks[0].RunId)
		assert.EqualValues(t, 1234, run.Tasks[1].RunId)
	})

	t.Run("run with two tasks pages", func(t *testing.T) {
		var requestMocks qa.HTTPFixtures = []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.2/jobs/runs/get?run_id=111222333",
				Response: Run{
					Iterations: []RunTask{},
					Tasks: []RunTask{
						{
							RunId: 123,
						},
						{
							RunId: 1234,
						},
					},
					JobClusters: []JobCluster{
						{
							JobClusterKey: "cluster1",
						},
						{
							JobClusterKey: "cluster2",
						},
					},
					NextPageToken: "token1",
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.2/jobs/runs/get?page_token=token1&run_id=111222333",
				Response: Run{
					Iterations: []RunTask{},
					Tasks: []RunTask{
						{
							RunId: 222,
						},
						{
							RunId: 333,
						},
					},
					JobClusters: []JobCluster{
						{
							JobClusterKey: "cluster1",
						},
						{
							JobClusterKey: "cluster2",
						},
					},
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.1/jobs/runs/get?run_id=111222333",
				Response: Run{
					Iterations: []RunTask{},
					Tasks: []RunTask{
						{
							RunId: 123,
						},
						{
							RunId: 1234,
						},
						{
							RunId: 222,
						},
						{
							RunId: 333,
						},
					},
					JobClusters: []JobCluster{
						{
							JobClusterKey: "cluster1",
						},
						{
							JobClusterKey: "cluster2",
						},
					},
				},
			},
		}
		client, server := requestMocks.Client(t)
		defer server.Close()

		mockJobsImpl := &jobsImpl{
			client: client,
		}
		api := &JobsAPI{jobsImpl: *mockJobsImpl}

		request := GetRunRequest{RunId: 111222333}
		run, err := api.GetRun(ctx, request)

		assert.NoError(t, err)
		assert.Equal(t, 4, len(run.Tasks))
		assert.Empty(t, run.Iterations)
		assert.Empty(t, run.NextPageToken)
		expected := []RunTask{
			{RunId: 123, ForceSendFields: []string{"RunId", "TaskKey"}},
			{RunId: 1234, ForceSendFields: []string{"RunId", "TaskKey"}},
			{RunId: 222, ForceSendFields: []string{"RunId", "TaskKey"}},
			{RunId: 333, ForceSendFields: []string{"RunId", "TaskKey"}},
		}
		assert.Equal(t, expected, run.Tasks)
	})

	t.Run("clusters array is also paginated", func(t *testing.T) {
		var requestMocks qa.HTTPFixtures = []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.2/jobs/runs/get?run_id=111222333",
				Response: Run{
					Iterations: []RunTask{},
					Tasks: []RunTask{
						{
							RunId: 123,
						},
						{
							RunId: 1234,
						},
					},
					JobClusters: []JobCluster{
						{
							JobClusterKey: "cluster1",
						},
						{
							JobClusterKey: "cluster2",
						},
					},
					JobParameters: []JobParameter{
						{
							Name:  "key1",
							Value: "value1",
						},
						{
							Name:  "key2",
							Value: "value2",
						},
					},
					NextPageToken: "token1",
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.2/jobs/runs/get?page_token=token1&run_id=111222333",
				Response: Run{
					Iterations: []RunTask{},
					Tasks: []RunTask{
						{
							RunId: 222,
						},
						{
							RunId: 333,
						},
					},
					JobClusters: []JobCluster{
						{
							JobClusterKey: "cluster3",
						},
					},
					JobParameters: []JobParameter{
						{
							Name:  "key3",
							Value: "value3",
						},
						{
							Name:  "key4",
							Value: "value4",
						},
					},
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.1/jobs/runs/get?run_id=111222333",
				Response: Run{
					Iterations: []RunTask{},
					Tasks: []RunTask{
						{
							RunId: 123,
						},
						{
							RunId: 1234,
						},
						{
							RunId: 222,
						},
						{
							RunId: 333,
						},
					},
					JobClusters: []JobCluster{
						{
							JobClusterKey: "cluster1",
						},
						{
							JobClusterKey: "cluster2",
						},
						{
							JobClusterKey: "cluster3",
						},
					},
					JobParameters: []JobParameter{
						{
							Name:  "key1",
							Value: "value1",
						},
						{
							Name:  "key2",
							Value: "value2",
						},
						{
							Name:  "key3",
							Value: "value3",
						},
						{
							Name:  "key4",
							Value: "value4",
						},
					},
				},
			},
		}
		client, server := requestMocks.Client(t)
		defer server.Close()

		mockJobsImpl := &jobsImpl{
			client: client,
		}
		api := &JobsAPI{jobsImpl: *mockJobsImpl}

		request := GetRunRequest{RunId: 111222333}
		run, err := api.GetRun(ctx, request)

		assert.NoError(t, err)
		assert.Equal(t, 4, len(run.Tasks))
		assert.Equal(t, 3, len(run.JobClusters))
		assert.Equal(t, 4, len(run.JobParameters))
		assert.Equal(t, "cluster1", run.JobClusters[0].JobClusterKey)
		assert.Equal(t, "cluster2", run.JobClusters[1].JobClusterKey)
		assert.Equal(t, "cluster3", run.JobClusters[2].JobClusterKey)
		assert.Equal(t, "key1", run.JobParameters[0].Name)
		assert.Equal(t, "value1", run.JobParameters[0].Value)
		assert.Equal(t, "key4", run.JobParameters[3].Name)
		assert.Equal(t, "value4", run.JobParameters[3].Value)
	})

	t.Run("run with two iterations pages", func(t *testing.T) {
		var requestMocks qa.HTTPFixtures = []qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.2/jobs/runs/get?run_id=4444",
				Response: Run{
					Iterations: []RunTask{
						{
							RunId: 123,
						},
						{
							RunId: 1234,
						},
					},
					Tasks: []RunTask{
						{
							RunId: 999,
						},
					},
					NextPageToken: "token1",
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.2/jobs/runs/get?page_token=token1&run_id=4444",
				Response: Run{
					Iterations: []RunTask{
						{
							RunId: 222,
						},
						{
							RunId: 333,
						},
					},
					Tasks: []RunTask{
						{
							RunId: 999,
						},
					},
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.1/jobs/runs/get?run_id=4444",
				Response: Run{
					Iterations: []RunTask{
						{
							RunId: 123,
						},
						{
							RunId: 1234,
						},
						{
							RunId: 222,
						},
						{
							RunId: 333,
						},
					},
					Tasks: []RunTask{
						{
							RunId: 999,
						},
					},
				},
			},
		}
		client, server := requestMocks.Client(t)
		defer server.Close()

		mockJobsImpl := &jobsImpl{
			client: client,
		}
		api := &JobsAPI{jobsImpl: *mockJobsImpl}

		request := GetRunRequest{RunId: 4444}
		run, err := api.GetRun(ctx, request)

		assert.NoError(t, err)
		assert.Equal(t, 4, len(run.Iterations))
		assert.Equal(t, 1, len(run.Tasks))
		assert.Empty(t, run.NextPageToken)
		expected := []RunTask{
			{RunId: 123, ForceSendFields: []string{"RunId", "TaskKey"}},
			{RunId: 1234, ForceSendFields: []string{"RunId", "TaskKey"}},
			{RunId: 222, ForceSendFields: []string{"RunId", "TaskKey"}},
			{RunId: 333, ForceSendFields: []string{"RunId", "TaskKey"}},
		}
		assert.Equal(t, expected, run.Iterations)
		assert.EqualValues(t, 999, run.Tasks[0].RunId)
	})
}

func TestGetJob(t *testing.T) {
	ctx := context.Background()

	t.Run("job with no pagination", func(t *testing.T) {
		var requestMocks qa.HTTPFixtures = []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/get?job_id=514594995218126",
				Response: Job{
					Settings: &JobSettings{
						Tasks: []Task{
							{
								TaskKey: "task1",
							},
							{
								TaskKey: "task2",
							},
							{
								TaskKey: "task3",
							},
							{
								TaskKey: "task4",
							},
						},
					},
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.1/jobs/get?job_id=514594995218126",
				Response: Job{
					Settings: &JobSettings{
						Tasks: []Task{
							{
								TaskKey: "task1",
							},
							{
								TaskKey: "task2",
							},
							{
								TaskKey: "task3",
							},
							{
								TaskKey: "task4",
							},
						},
					},
				},
			},
		}
		client, server := requestMocks.Client(t)
		defer server.Close()

		mockJobsImpl := &jobsImpl{
			client: client,
		}
		api := &JobsAPI{jobsImpl: *mockJobsImpl}

		request := GetJobRequest{JobId: 514594995218126}
		job, err := api.Get(ctx, request)

		assert.NoError(t, err)
		assert.Equal(t, 4, len(job.Settings.Tasks))
		assert.EqualValues(t, "task1", job.Settings.Tasks[0].TaskKey)
		assert.EqualValues(t, "task4", job.Settings.Tasks[3].TaskKey)
	})

	t.Run("job with multiple pages", func(t *testing.T) {
		var requestMocks qa.HTTPFixtures = []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/get?job_id=514594995218126",
				Response: Job{
					Settings: &JobSettings{
						Tasks: []Task{
							{
								TaskKey: "task1",
							},
							{
								TaskKey: "task2",
							},
						},
						JobClusters: []JobCluster{
							{
								JobClusterKey: "cluster1",
							},
							{
								JobClusterKey: "cluster2",
							},
						},
						Parameters: []JobParameterDefinition{
							{
								Name:    "param1",
								Default: "default1",
							},
							{
								Name:    "param2",
								Default: "default2",
							},
						},
						Environments: []JobEnvironment{
							{
								EnvironmentKey: "env1",
							},
						},
					},
					NextPageToken: "token1",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/get?job_id=514594995218126&page_token=token1",
				Response: Job{
					Settings: &JobSettings{
						Tasks: []Task{
							{
								TaskKey: "task3",
							},
							{
								TaskKey: "task4",
							},
						},
						JobClusters: []JobCluster{
							{
								JobClusterKey: "cluster3",
							},
							{
								JobClusterKey: "cluster4",
							},
						},
						Parameters: []JobParameterDefinition{
							{
								Name:    "param3",
								Default: "default3",
							},
						},
						Environments: []JobEnvironment{
							{
								EnvironmentKey: "env2",
							},
						},
					},
					NextPageToken: "token2",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/get?job_id=514594995218126&page_token=token2",
				Response: Job{
					Settings: &JobSettings{
						Tasks: []Task{
							{
								TaskKey: "task5",
							},
						},
						Environments: []JobEnvironment{
							{
								EnvironmentKey: "env3",
							},
						},
					},
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.1/jobs/get?job_id=514594995218126",
				Response: Job{
					Settings: &JobSettings{
						Tasks: []Task{
							{
								TaskKey: "task1",
							},
							{
								TaskKey: "task2",
							},
							{
								TaskKey: "task3",
							},
							{
								TaskKey: "task4",
							},
							{
								TaskKey: "task5",
							},
						},
						JobClusters: []JobCluster{
							{
								JobClusterKey: "cluster1",
							},
							{
								JobClusterKey: "cluster2",
							},
							{
								JobClusterKey: "cluster3",
							},
							{
								JobClusterKey: "cluster4",
							},
						},
						Parameters: []JobParameterDefinition{
							{
								Name:    "param1",
								Default: "default1",
							},
							{
								Name:    "param2",
								Default: "default2",
							},
							{
								Name:    "param3",
								Default: "default3",
							},
						},
						Environments: []JobEnvironment{
							{
								EnvironmentKey: "env1",
							},
							{
								EnvironmentKey: "env2",
							},
							{
								EnvironmentKey: "env3",
							},
						},
					},
				},
			},
		}
		client, server := requestMocks.Client(t)
		defer server.Close()

		mockJobsImpl := &jobsImpl{
			client: client,
		}
		api := &JobsAPI{jobsImpl: *mockJobsImpl}

		request := GetJobRequest{JobId: 514594995218126}
		job, err := api.Get(ctx, request)

		assert.NoError(t, err)
		assert.Equal(t, 5, len(job.Settings.Tasks))
		assert.Equal(t, 4, len(job.Settings.JobClusters))
		assert.Equal(t, 3, len(job.Settings.Parameters))
		assert.Equal(t, 3, len(job.Settings.Environments))
		assert.EqualValues(t, "task1", job.Settings.Tasks[0].TaskKey)
		assert.EqualValues(t, "task4", job.Settings.Tasks[3].TaskKey)
		assert.EqualValues(t, "task5", job.Settings.Tasks[4].TaskKey)
		assert.EqualValues(t, "cluster3", job.Settings.JobClusters[2].JobClusterKey)
		assert.EqualValues(t, "param3", job.Settings.Parameters[2].Name)
		assert.EqualValues(t, "env3", job.Settings.Environments[2].EnvironmentKey)
	})
}

func TestListJobs(t *testing.T) {
	ctx := context.Background()

	t.Run("jobs list with no task expansion", func(t *testing.T) {
		var requestMocks qa.HTTPFixtures = []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/list?",
				Response: ListJobsResponse{
					Jobs: []BaseJob{
						{
							JobId: 100,
							Settings: &JobSettings{
								Name: "job_100",
							},
						},
						{
							JobId: 200,
							Settings: &JobSettings{
								Name: "job_200",
							},
						},
					},
					NextPageToken: "token1",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/list?page_token=token1",
				Response: ListJobsResponse{
					Jobs: []BaseJob{
						{
							JobId: 300,
							Settings: &JobSettings{
								Name: "job_300",
							},
						},
						{
							JobId: 400,
							Settings: &JobSettings{
								Name: "job_400",
							},
						},
					},
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.1/jobs/list?",
				Response: ListJobsResponse{
					Jobs: []BaseJob{
						{
							JobId: 100,
							Settings: &JobSettings{
								Name: "job_100",
							},
						},
						{
							JobId: 200,
							Settings: &JobSettings{
								Name: "job_200",
							},
						},
					},
					NextPageToken: "token1",
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.1/jobs/list?page_token=token1",
				Response: ListJobsResponse{
					Jobs: []BaseJob{
						{
							JobId: 300,
							Settings: &JobSettings{
								Name: "job_300",
							},
						},
						{
							JobId: 400,
							Settings: &JobSettings{
								Name: "job_400",
							},
						},
					},
				},
			},
		}
		client, server := requestMocks.Client(t)
		defer server.Close()

		mockJobsImpl := &jobsImpl{
			client: client,
		}
		api := &JobsAPI{jobsImpl: *mockJobsImpl}

		jobsList := api.List(ctx, ListJobsRequest{})
		var allJobs []BaseJob
		for jobsList.HasNext(ctx) {
			job, err := jobsList.Next(ctx)
			assert.NoError(t, err)
			assert.NotEmpty(t, job.JobId)
			allJobs = append(allJobs, job)
		}

		assert.EqualValues(t, len(allJobs), 4)
		assert.EqualValues(t, allJobs[0].JobId, 100)
		assert.EqualValues(t, allJobs[2].JobId, 300)
		assert.EqualValues(t, allJobs[3].JobId, 400)
		assert.EqualValues(t, allJobs[3].Settings.Name, "job_400")
	})

	t.Run("jobs list with task expansion", func(t *testing.T) {
		var requestMocks qa.HTTPFixtures = []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/list?expand_tasks=true",
				Response: ListJobsResponse{
					Jobs: []BaseJob{
						{
							JobId: 100,
							Settings: &JobSettings{
								Name: "job_100",
								Tasks: []Task{
									{
										TaskKey: "job100_task1",
									},
									{
										TaskKey: "job100_task2",
									},
								},
								JobClusters: []JobCluster{
									{
										JobClusterKey: "job100_cluster1",
									},
									{
										JobClusterKey: "job100_cluster2",
									},
								},
								Parameters: []JobParameterDefinition{
									{
										Name:    "param1",
										Default: "default1",
									},
									{
										Name:    "param2",
										Default: "default2",
									},
								},
								Environments: []JobEnvironment{
									{
										EnvironmentKey: "env1",
									},
								},
							},
							HasMore: true,
						},
						{
							JobId: 200,
							Settings: &JobSettings{
								Name: "job_200",
								Tasks: []Task{
									{
										TaskKey: "job200_task1",
									},
									{
										TaskKey: "job200_task2",
									},
								},
								JobClusters: []JobCluster{
									{
										JobClusterKey: "job200_cluster1",
									},
								},
								Environments: []JobEnvironment{
									{
										EnvironmentKey: "env1",
									},
								},
							},
							HasMore: true,
						},
					},
					NextPageToken: "token1",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/list?expand_tasks=true&page_token=token1",
				Response: ListJobsResponse{
					Jobs: []BaseJob{
						{
							JobId: 300,
							Settings: &JobSettings{
								Name: "job_300",
								Tasks: []Task{
									{
										TaskKey: "job300_task1",
									},
								},
							},
						},
						{
							JobId: 400,
							Settings: &JobSettings{
								Name: "job_400",
								Tasks: []Task{
									{
										TaskKey: "job400_task3",
									},
									{
										TaskKey: "job400_task1",
									},
								},
								Parameters: []JobParameterDefinition{
									{
										Name:    "param1",
										Default: "default1",
									},
									{
										Name:    "param2",
										Default: "default2",
									},
								},
							},
							HasMore: true,
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/get?job_id=100",
				Response: Job{
					Settings: &JobSettings{
						Tasks: []Task{
							{
								TaskKey: "job100_task1",
							},
							{
								TaskKey: "job100_task2",
							},
						},
						JobClusters: []JobCluster{
							{
								JobClusterKey: "job100_cluster1",
							},
							{
								JobClusterKey: "job100_cluster2",
							},
						},
						Parameters: []JobParameterDefinition{
							{
								Name:    "param1",
								Default: "default1",
							},
							{
								Name:    "param2",
								Default: "default2",
							},
						},
						Environments: []JobEnvironment{
							{
								EnvironmentKey: "env1",
							},
						},
					},
					NextPageToken: "token1",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/get?job_id=100&page_token=token1",
				Response: Job{
					Settings: &JobSettings{
						Tasks: []Task{
							{
								TaskKey: "job100_task3",
							},
							{
								TaskKey: "job100_task4",
							},
						},
						JobClusters: []JobCluster{
							{
								JobClusterKey: "job100_cluster3",
							},
							{
								JobClusterKey: "job100_cluster4",
							},
						},
						Parameters: []JobParameterDefinition{
							{
								Name:    "param3",
								Default: "default3",
							},
						},
						Environments: []JobEnvironment{
							{
								EnvironmentKey: "env2",
							},
						},
					},
					NextPageToken: "token2",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/get?job_id=100&page_token=token2",
				Response: Job{
					Settings: &JobSettings{
						Tasks: []Task{
							{
								TaskKey: "job100_task5",
							},
						},
						Environments: []JobEnvironment{
							{
								EnvironmentKey: "env3",
							},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/get?job_id=200",
				Response: Job{
					Settings: &JobSettings{
						Tasks: []Task{
							{
								TaskKey: "job200_task1",
							},
							{
								TaskKey: "job200_task2",
							},
						},
						JobClusters: []JobCluster{
							{
								JobClusterKey: "job200_cluster1",
							},
						},
						Environments: []JobEnvironment{
							{
								EnvironmentKey: "env1",
							},
						},
					},
					NextPageToken: "token1",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/get?job_id=200&page_token=token1",
				Response: Job{
					Settings: &JobSettings{
						Tasks: []Task{
							{
								TaskKey: "job200_task3",
							},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/get?job_id=400",
				Response: Job{
					Settings: &JobSettings{
						Tasks: []Task{
							{
								TaskKey: "job400_task1",
							},
							{
								TaskKey: "job400_task2",
							},
						},
						Parameters: []JobParameterDefinition{
							{
								Name:    "param1",
								Default: "default1",
							},
							{
								Name:    "param2",
								Default: "default2",
							},
						},
					},
					NextPageToken: "token1",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/get?job_id=400&page_token=token1",
				Response: Job{
					Settings: &JobSettings{
						Tasks: []Task{
							{
								TaskKey: "job400_task3",
							},
							{
								TaskKey: "job400_task4",
							},
						},
						Parameters: []JobParameterDefinition{
							{
								Name:    "param3",
								Default: "default3",
							},
							{
								Name:    "param4",
								Default: "default4",
							},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/jobs/list?expand_tasks=true",
				Response: ListJobsResponse{
					Jobs: []BaseJob{
						{
							JobId: 100,
							Settings: &JobSettings{
								Name: "job_100",
								Tasks: []Task{
									{
										TaskKey: "job100_task1",
									},
									{
										TaskKey: "job100_task2",
									},
									{
										TaskKey: "job100_task3",
									},
									{
										TaskKey: "job100_task4",
									},
									{
										TaskKey: "job100_task5",
									},
								},
								JobClusters: []JobCluster{
									{
										JobClusterKey: "job100_cluster1",
									},
									{
										JobClusterKey: "job100_cluster2",
									},
									{
										JobClusterKey: "job100_cluster3",
									},
									{
										JobClusterKey: "job100_cluster4",
									},
								},
								Parameters: []JobParameterDefinition{
									{
										Name:    "param1",
										Default: "default1",
									},
									{
										Name:    "param2",
										Default: "default2",
									},
									{
										Name:    "param3",
										Default: "default3",
									},
								},
								Environments: []JobEnvironment{
									{
										EnvironmentKey: "env1",
									},
									{
										EnvironmentKey: "env2",
									},
									{
										EnvironmentKey: "env3",
									},
								},
							},
						},
						{
							JobId: 200,
							Settings: &JobSettings{
								Name: "job_200",
								Tasks: []Task{
									{
										TaskKey: "job200_task1",
									},
									{
										TaskKey: "job200_task2",
									},
									{
										TaskKey: "job200_task3",
									},
								},
								JobClusters: []JobCluster{
									{
										JobClusterKey: "job200_cluster1",
									},
								},
								Environments: []JobEnvironment{
									{
										EnvironmentKey: "env1",
									},
								},
							},
						},
					},
					NextPageToken: "token1",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/jobs/list?expand_tasks=true&page_token=token1",
				Response: ListJobsResponse{
					Jobs: []BaseJob{
						{
							JobId: 300,
							Settings: &JobSettings{
								Name: "job_300",
								Tasks: []Task{
									{
										TaskKey: "job300_task1",
									},
								},
							},
						},
						{
							JobId: 400,
							Settings: &JobSettings{
								Name: "job_400",
								Tasks: []Task{
									{
										TaskKey: "job400_task1",
									},
									{
										TaskKey: "job400_task2",
									},
									{
										TaskKey: "job400_task4",
									},
									{
										TaskKey: "job400_task5",
									},
								},
								Parameters: []JobParameterDefinition{
									{
										Name:    "param1",
										Default: "default1",
									},
									{
										Name:    "param2",
										Default: "default2",
									},
									{
										Name:    "param3",
										Default: "default3",
									},
									{
										Name:    "param4",
										Default: "default4",
									},
								},
							},
						},
					},
				},
			},
		}
		client, server := requestMocks.Client(t)
		defer server.Close()

		mockJobsImpl := &jobsImpl{
			client: client,
		}
		api := &JobsAPI{jobsImpl: *mockJobsImpl}

		jobsList := api.List(ctx, ListJobsRequest{ExpandTasks: true})
		var allJobs []BaseJob
		for jobsList.HasNext(ctx) {
			job, err := jobsList.Next(ctx)
			assert.NoError(t, err)
			assert.NotEmpty(t, job.JobId)
			assert.Empty(t, job.HasMore)
			allJobs = append(allJobs, job)
		}

		assert.Equal(t, 4, len(allJobs))
		assert.Equal(t, 5, len(allJobs[0].Settings.Tasks))
		assert.Equal(t, 4, len(allJobs[0].Settings.JobClusters))
		assert.Equal(t, 3, len(allJobs[0].Settings.Parameters))
		assert.Equal(t, 3, len(allJobs[0].Settings.Environments))
		assert.Equal(t, 3, len(allJobs[1].Settings.Tasks))
		assert.Equal(t, 1, len(allJobs[1].Settings.JobClusters))
		assert.Empty(t, allJobs[1].Settings.Parameters)
		assert.Equal(t, 1, len(allJobs[1].Settings.Environments))
		assert.Equal(t, 1, len(allJobs[2].Settings.Tasks))
		assert.Empty(t, allJobs[2].Settings.JobClusters)
		assert.Empty(t, allJobs[2].Settings.Parameters)
		assert.Empty(t, allJobs[2].Settings.Environments)
		assert.EqualValues(t, 100, allJobs[0].JobId)
		assert.EqualValues(t, 300, allJobs[2].JobId)
		assert.EqualValues(t, 400, allJobs[3].JobId)
		assert.EqualValues(t, "job_400", allJobs[3].Settings.Name)
	})
}

func TestListRuns(t *testing.T) {
	t.Run("runs list with no task expansion", func(t *testing.T) {
		ctx := context.Background()

		var requestMocks qa.HTTPFixtures = []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/runs/list?",
				Response: ListRunsResponse{
					Runs: []BaseRun{
						{
							RunId:   100,
							RunName: "run100",
						},
						{
							RunId:   200,
							RunName: "run200",
							JobParameters: []JobParameter{
								{
									Name:    "param1",
									Default: "default1",
								},
								{
									Name:    "param2",
									Default: "default2",
								},
							},
						},
						{
							RunId:   300,
							RunName: "run300",
						},
					},
					NextPageToken: "tokenToSecondPage",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/runs/list?page_token=tokenToSecondPage",
				Response: ListRunsResponse{
					Runs: []BaseRun{
						{
							RunId:   400,
							RunName: "run400",
							RepairHistory: []RepairHistoryItem{
								{
									Id: 410,
								},
								{
									Id: 411,
								},
							},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/jobs/runs/list?",
				Response: ListRunsResponse{
					Runs: []BaseRun{
						{
							RunId:   100,
							RunName: "run100",
						},
						{
							RunId:   200,
							RunName: "run200",
							JobParameters: []JobParameter{
								{
									Name:    "param1",
									Default: "default1",
								},
								{
									Name:    "param2",
									Default: "default2",
								},
							},
						},
						{
							RunId:   300,
							RunName: "run300",
						},
					},
					NextPageToken: "tokenToSecondPage",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/jobs/runs/list?page_token=tokenToSecondPage",
				Response: ListRunsResponse{
					Runs: []BaseRun{
						{
							RunId:   400,
							RunName: "run400",
							RepairHistory: []RepairHistoryItem{
								{
									Id: 410,
								},
								{
									Id: 411,
								},
							},
						},
					},
				},
			},
		}

		client, server := requestMocks.Client(t)
		defer server.Close()

		mockJobsImpl := &jobsImpl{
			client: client,
		}
		api := &JobsAPI{jobsImpl: *mockJobsImpl}

		runsList := api.ListRuns(ctx, ListRunsRequest{})
		var allRuns []BaseRun
		for runsList.HasNext(ctx) {
			run, err := runsList.Next(ctx)
			assert.NoError(t, err)
			assert.NotEmpty(t, run.RunId)
			assert.Empty(t, run.HasMore)
			allRuns = append(allRuns, run)
		}

		assert.EqualValues(t, len(allRuns), 4)
		assert.EqualValues(t, allRuns[0].RunId, 100)
		assert.EqualValues(t, allRuns[2].RunId, 300)
		assert.EqualValues(t, allRuns[3].RunId, 400)
		assert.EqualValues(t, allRuns[3].RunName, "run400")
	})

	t.Run("runs list with with task expansion", func(t *testing.T) {
		ctx := context.Background()

		var requestMocks qa.HTTPFixtures = []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/runs/list?expand_tasks=true",
				Response: ListRunsResponse{
					Runs: []BaseRun{
						{
							RunId: 100,
							Tasks: []RunTask{
								{TaskKey: "taskkey101"},
								{TaskKey: "taskkey102"},
							},
							HasMore: true,
						},
						{
							RunId: 200,
							Tasks: []RunTask{
								{TaskKey: "taskkey201"},
							},
						},
						{
							RunId: 300,
							Tasks: []RunTask{
								{TaskKey: "taskkey301"},
							},
						},
					},
					NextPageToken: "tokenToSecondPage",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/runs/list?expand_tasks=true&page_token=tokenToSecondPage",
				Response: ListRunsResponse{
					Runs: []BaseRun{
						{
							RunId: 400,
							Tasks: []RunTask{
								{TaskKey: "taskkey401"},
								{TaskKey: "taskkey402"},
							},
							HasMore: true,
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/runs/get?run_id=100",
				Response: Run{
					RunId: 100,
					Tasks: []RunTask{
						{TaskKey: "taskkey101"},
						{TaskKey: "taskkey102"},
					},
					NextPageToken: "tokenToSecondPage_100",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/runs/get?page_token=tokenToSecondPage_100&run_id=100",
				Response: Run{
					RunId: 100,
					Tasks: []RunTask{
						{TaskKey: "taskkey103"},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/runs/get?run_id=400",
				Response: Run{
					RunId: 400,
					Tasks: []RunTask{
						{TaskKey: "taskkey401"},
						{TaskKey: "taskkey403"},
					},
					NextPageToken: "tokenToSecondPage_400",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.2/jobs/runs/get?page_token=tokenToSecondPage_400&run_id=400",
				Response: Run{
					RunId: 400,
					Tasks: []RunTask{
						{TaskKey: "taskkey402"},
						{TaskKey: "taskkey404"},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/jobs/runs/list?expand_tasks=true",
				Response: ListRunsResponse{
					Runs: []BaseRun{
						{
							RunId: 100,
							Tasks: []RunTask{
								{TaskKey: "taskkey101"},
								{TaskKey: "taskkey102"},
								{TaskKey: "taskkey103"},
							},
						},
						{
							RunId: 200,
							Tasks: []RunTask{
								{TaskKey: "taskkey201"},
							},
						},
						{
							RunId: 300,
							Tasks: []RunTask{
								{TaskKey: "taskkey301"},
							},
						},
						{
							RunId: 400,
							Tasks: []RunTask{
								{TaskKey: "taskkey401"},
								{TaskKey: "taskkey403"},
								{TaskKey: "taskkey402"},
								{TaskKey: "taskkey404"},
							},
						},
					},
				},
			},
		}

		client, server := requestMocks.Client(t)
		defer server.Close()

		mockJobsImpl := &jobsImpl{
			client: client,
		}
		api := &JobsAPI{jobsImpl: *mockJobsImpl}

		runsList := api.ListRuns(ctx, ListRunsRequest{ExpandTasks: true})
		var allRuns []BaseRun
		for runsList.HasNext(ctx) {
			run, err := runsList.Next(ctx)
			assert.NoError(t, err)
			assert.NotEmpty(t, run.RunId)
			assert.Empty(t, run.HasMore)
			allRuns = append(allRuns, run)
		}

		assert.EqualValues(t, 4, len(allRuns))
		assert.EqualValues(t, 100, allRuns[0].RunId)
		assert.EqualValues(t, 300, allRuns[2].RunId)
		assert.EqualValues(t, 400, allRuns[3].RunId)
		assert.Equal(t, 3, len(allRuns[0].Tasks))
		assert.EqualValues(t, "taskkey401", allRuns[3].Tasks[0].TaskKey)
		assert.EqualValues(t, "taskkey403", allRuns[3].Tasks[1].TaskKey)
		assert.EqualValues(t, "taskkey402", allRuns[3].Tasks[2].TaskKey)
		assert.EqualValues(t, "taskkey404", allRuns[3].Tasks[3].TaskKey)
	})
}
