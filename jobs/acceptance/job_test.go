package acceptance

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"

	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/internal/compute"
	"github.com/databricks/terraform-provider-databricks/jobs"
	"github.com/databricks/terraform-provider-databricks/libraries"
	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAwsAccJobsCreate(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}

	client := common.NewClientFromEnvironment()
	jobsAPI := jobs.NewJobsAPI(context.Background(), client)
	clustersAPI := clusters.NewClustersAPI(context.Background(), client)
	sparkVersion := clustersAPI.LatestSparkVersionOrDefault(clusters.SparkVersionRequest{Latest: true, LongTermSupport: true})

	jobSettings := jobs.JobSettings{
		NewCluster: &clusters.Cluster{
			NumWorkers:   2,
			SparkVersion: sparkVersion,
			SparkConf:    nil,
			AwsAttributes: &clusters.AwsAttributes{
				Availability: "ON_DEMAND",
			},
			NodeTypeID: clustersAPI.GetSmallestNodeType(clusters.NodeTypeRequest{
				LocalDisk: true,
			}),
		},
		NotebookTask: &jobs.NotebookTask{
			NotebookPath: "/tf-test/demo-terraform/demo-notebook",
		},
		Name: "1-test-job",
		Libraries: []libraries.Library{
			{
				Maven: &libraries.Maven{
					Coordinates: "org.jsoup:jsoup:1.7.2",
				},
			},
		},
		EmailNotifications: &jobs.EmailNotifications{
			OnStart:   []string{},
			OnSuccess: []string{},
			OnFailure: []string{},
		},
		TimeoutSeconds: 3600,
		MaxRetries:     1,
		Schedule: &jobs.CronSchedule{
			QuartzCronExpression: "0 15 22 ? * *",
			TimezoneID:           "America/Los_Angeles",
		},
		MaxConcurrentRuns: 1,
	}

	job, err := jobsAPI.Create(jobSettings)
	require.NoError(t, err, err)
	id := job.ID()
	defer func() {
		err := jobsAPI.Delete(id)
		assert.NoError(t, err, err)
	}()
	t.Log(id)
	job, err = jobsAPI.Read(id)
	assert.NoError(t, err, err)
	assert.True(t, job.Settings.NewCluster.SparkVersion == sparkVersion, "Something is wrong with spark version")

	newSparkVersion := clustersAPI.LatestSparkVersionOrDefault(clusters.SparkVersionRequest{Latest: true})
	jobSettings.NewCluster.SparkVersion = newSparkVersion

	err = jobsAPI.Update(id, jobSettings)
	assert.NoError(t, err, err)

	job, err = jobsAPI.Read(id)
	assert.NoError(t, err, err)
	assert.True(t, job.Settings.NewCluster.SparkVersion == newSparkVersion, "Something is wrong with spark version")
}

func TestPreviewAccJobTasks(t *testing.T) {
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			data "databricks_current_user" "me" {}
			data "databricks_spark_version" "latest" {}
			data "databricks_node_type" "smallest" {
				local_disk = true
			}

			resource "databricks_notebook" "this" {
				path     = "${data.databricks_current_user.me.home}/Terraform{var.RANDOM}"
				language = "PYTHON"
				content_base64 = base64encode(<<-EOT
					# created from ${abspath(path.module)}
					display(spark.range(10))
					EOT
				)
			}

			resource "databricks_job" "this" {
				name = "{var.RANDOM}"

				job_cluster {
					job_cluster_key = "j"
					new_cluster {
						num_workers   = 20
						spark_version = data.databricks_spark_version.latest.id
						node_type_id  = data.databricks_node_type.smallest.id
					}
				}

				task {
					task_key = "a"

					new_cluster {
						num_workers   = 1
						spark_version = data.databricks_spark_version.latest.id
						node_type_id  = data.databricks_node_type.smallest.id
					}

					notebook_task {
						notebook_path = databricks_notebook.this.path
					}
				}

				task {
					task_key = "b"

					depends_on {
						task_key = "a"
					}

					new_cluster {
						num_workers   = 8
						spark_version = data.databricks_spark_version.latest.id
						node_type_id  = data.databricks_node_type.smallest.id
					}

					notebook_task {
						notebook_path = databricks_notebook.this.path
					}
				}

				task {
					task_key = "c"
					
					job_cluster_key = "j"

					depends_on {
						task_key = "b"
					}

					notebook_task {
						notebook_path = databricks_notebook.this.path
					}
				}
			}`,
		},
	})
}

func TestAccJobResource(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	t.Parallel()
	clustersAPI := clusters.NewClustersAPI(context.Background(), common.CommonEnvironmentClient())
	sparkVersion := clustersAPI.LatestSparkVersionOrDefault(clusters.SparkVersionRequest{Latest: true, LongTermSupport: true})
	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`resource "databricks_job" "this" {
					new_cluster  {
					  autoscale  {
						min_workers = 2
						max_workers = 3
					  }
					  instance_pool_id = "%s"
					  spark_version = "%s"
					}
					notebook_task {
						notebook_path = "/Production/MakeFeatures"
					}
					email_notifications {
						no_alert_for_skipped_runs = true
					}
					name = "%s"
					timeout_seconds = 3600
					max_retries = 1
					max_concurrent_runs = 1
				  }`, compute.CommonInstancePoolID(), sparkVersion, qa.RandomLongName()),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					acceptance.ResourceCheck("databricks_job.this",
						func(ctx context.Context, client *common.DatabricksClient, id string) error {
							job, err := jobs.NewJobsAPI(ctx, client).Read(id)
							assert.NoError(t, err)
							assert.NotNil(t, job.Settings)
							assert.NotNil(t, job.Settings.NewCluster)
							assert.NotNil(t, job.Settings.NewCluster.Autoscale)
							assert.NotNil(t, job.Settings.NotebookTask)
							assert.Equal(t, 2, int(job.Settings.NewCluster.Autoscale.MinWorkers))
							assert.Equal(t, 3, int(job.Settings.NewCluster.Autoscale.MaxWorkers))
							assert.Equal(t, sparkVersion, job.Settings.NewCluster.SparkVersion)
							assert.Equal(t, "/Production/MakeFeatures", job.Settings.NotebookTask.NotebookPath)
							assert.Equal(t, 3600, int(job.Settings.TimeoutSeconds))
							assert.Equal(t, 1, int(job.Settings.MaxRetries))
							assert.Equal(t, 1, int(job.Settings.MaxConcurrentRuns))
							return nil
						}),
				),
			},
		},
	})
}

func TestAwsAccJobResource_NoInstancePool(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	clustersAPI := clusters.NewClustersAPI(context.Background(), common.CommonEnvironmentClient())
	sparkVersion := clustersAPI.LatestSparkVersionOrDefault(clusters.SparkVersionRequest{Latest: true, LongTermSupport: true})
	randomStr := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	instanceProfileARN := fmt.Sprintf("arn:aws:iam::999999999999:instance-profile/tf-test-%s", randomStr)
	acceptance.AccTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`resource "databricks_job" "this" {
					new_cluster  {
					  num_workers = 1
					  aws_attributes {
 					    zone_id = "eu-central-1"
		                spot_bid_price_percent = "100"
					    instance_profile_arn = "%s"
					    first_on_demand = 1
					    ebs_volume_type = "GENERAL_PURPOSE_SSD"
					    ebs_volume_count = 1
					    ebs_volume_size = 32
					  }
					  node_type_id = "m4.large"
					  spark_version = "%s"
					}
					notebook_task {
						notebook_path = "/Production/MakeFeatures"
					}
					library {
						pypi {
							package = "networkx"
						}
					}
					email_notifications {
						no_alert_for_skipped_runs = true
					}
					name = "%s"
					timeout_seconds = 3600
					max_retries = 1
					max_concurrent_runs = 1
				  }`, instanceProfileARN, sparkVersion,
					qa.RandomLongName()),
				// compose a basic test, checking both remote and local values
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					acceptance.ResourceCheck("databricks_job.this",
						func(ctx context.Context, client *common.DatabricksClient, id string) error {
							job, err := jobs.NewJobsAPI(ctx, client).Read(id)
							assert.NoError(t, err)
							assert.NotNil(t, job.Settings)
							assert.NotNil(t, job.Settings.NewCluster)
							assert.NotNil(t, job.Settings.NewCluster.AwsAttributes)
							return nil
						}),
				),
			},
		},
	})
}
