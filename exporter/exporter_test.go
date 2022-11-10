package exporter

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/databricks/terraform-provider-databricks/access"
	"github.com/databricks/terraform-provider-databricks/aws"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/commands"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/jobs"
	"github.com/databricks/terraform-provider-databricks/libraries"
	"github.com/databricks/terraform-provider-databricks/pipelines"
	"github.com/databricks/terraform-provider-databricks/policies"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/databricks/terraform-provider-databricks/repos"
	"github.com/databricks/terraform-provider-databricks/scim"
	"github.com/databricks/terraform-provider-databricks/secrets"
	"github.com/databricks/terraform-provider-databricks/sql"
	"github.com/databricks/terraform-provider-databricks/workspace"
	"github.com/hashicorp/hcl/v2/hclwrite"

	"github.com/stretchr/testify/assert"
)

// nolint
func getJSONObject(filename string) any {
	data, _ := ioutil.ReadFile(filename)
	var obj map[string]any
	err := json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Printf("[ERROR] error! err=%v\n", err)
		fmt.Printf("[ERROR] data=%s\n", string(data))
	}
	return obj
}

func TestImportingMounts(t *testing.T) {
	qa.HTTPFixturesApply(t,
		[]qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Me",
				Response: scim.User{},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/list",
				Response: clusters.ClusterList{
					Clusters: []clusters.ClusterInfo{
						{
							ClusterName: "terraform-mount",
							ClusterID:   "mount",
						},
						{
							ClusterName: "terraform-mount-shard-s3-access",
							ClusterID:   "mount",
						},
					},
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/get?cluster_id=mount",
				Response: clusters.ClusterInfo{
					State:       "RUNNING",
					ClusterID:   "mount",
					ClusterName: "dummy",
				},
			},
			{
				Method:       "POST",
				ReuseRequest: true,
				Resource:     "/api/1.2/contexts/create",
				Response: commands.Command{
					ID: "context",
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/1.2/contexts/status?clusterId=mount&contextId=context",
				Response: commands.Command{
					Status: "Running",
				},
			},
			{
				Method:       "POST",
				ReuseRequest: true,
				Resource:     "/api/1.2/commands/execute",
				Response: commands.Command{
					ID: "run",
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/1.2/commands/status?clusterId=mount&commandId=run&contextId=context",
				Response: commands.Command{
					Status: "Finished",
					Results: &common.CommandResults{
						ResultType: "text",
						Data: `{"foo": "s3a://foo", "bar": "abfss://bar@baz.com/thing", "third": "adls://foo3.bar.com/path", "fourth":"wasbs://bar4@baz4.com/dir", "fifth":"gs://foo5", "sixth":"abc://foo5"}
					and some chatty messages`,
					},
				},
			},
			{
				Method:       "POST",
				Resource:     "/api/1.2/contexts/destroy",
				ReuseRequest: true,
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/instance-profiles/list",
				Response: aws.InstanceProfileList{
					InstanceProfiles: []aws.InstanceProfileInfo{
						{
							InstanceProfileArn: "arn:aws:iam::12345:instance-profile/shard-s3-access",
						},
					},
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/spark-versions",
				Response: clusters.SparkVersionsList{
					SparkVersions: []clusters.SparkVersion{
						{
							Version: "Foo LTS",
						},
					},
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/list-node-types",
				Response: clusters.NodeTypeList{
					NodeTypes: []clusters.NodeType{
						{
							NodeTypeID: "m5d.large",
						},
					},
				},
			},
			{
				Method:       "POST",
				ReuseRequest: true,
				Resource:     "/api/2.0/clusters/events",
				Response: clusters.EventsResponse{
					Events: []clusters.ClusterEvent{},
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/libraries/cluster-status?cluster_id=mount",
				Response: libraries.ClusterLibraryList{
					Libraries: []libraries.Library{},
				},
			},
		}, func(ctx context.Context, client *common.DatabricksClient) {
			ic := newImportContext(client)
			ic.services = "mounts"
			ic.listing = "mounts"
			ic.mounts = true

			err := ic.Importables["databricks_mount"].List(ic)
			assert.NoError(t, err)

			for i := 0; i < len(ic.Scope); i++ {
				err = ic.Importables["databricks_mount"].Body(ic,
					hclwrite.NewEmptyFile().Body(), ic.Scope[i])
				assert.NoError(t, err)
			}
		})
}

var meAdminFixture = qa.HTTPFixture{
	Method:       "GET",
	ReuseRequest: true,
	Resource:     "/api/2.0/preview/scim/v2/Me",
	Response: scim.User{
		Groups: []scim.ComplexValue{
			{
				Display: "admins",
			},
		},
	},
}

var emptyPipelines = qa.HTTPFixture{
	Method:       "GET",
	ReuseRequest: true,
	Resource:     "/api/2.0/pipelines?max_results=50",
	Response:     pipelines.PipelineListResponse{},
}

var emptyRepos = qa.HTTPFixture{
	Method:       "GET",
	ReuseRequest: true,
	Resource:     "/api/2.0/repos?",
	Response:     repos.ReposListResponse{},
}

var emptyGitCredentials = qa.HTTPFixture{
	Method:   http.MethodGet,
	Resource: "/api/2.0/git-credentials",
	Response: repos.GitCredentialList{},
}

var emptyIpAccessLIst = qa.HTTPFixture{
	Method:   http.MethodGet,
	Resource: "/api/2.0/ip-access-lists",
	Response: map[string]any{},
}

var emptyWorkspace = qa.HTTPFixture{
	Method:   "GET",
	Resource: "/api/2.0/workspace/list?path=%2F",
	Response: workspace.ObjectList{},
}

var emptySqlEndpoints = qa.HTTPFixture{
	Method:       "GET",
	Resource:     "/api/2.0/sql/warehouses",
	Response:     map[string]any{},
	ReuseRequest: true,
}

var emptySqlDashboards = qa.HTTPFixture{
	Method:       "GET",
	Resource:     "/api/2.0/preview/sql/dashboards",
	Response:     map[string]any{},
	ReuseRequest: true,
}

var emptySqlQueries = qa.HTTPFixture{
	Method:       "GET",
	Resource:     "/api/2.0/preview/sql/queries",
	Response:     map[string]any{},
	ReuseRequest: true,
}

func TestImportingUsersGroupsSecretScopes(t *testing.T) {
	qa.HTTPFixturesApply(t,
		[]qa.HTTPFixture{
			meAdminFixture,
			emptyRepos,
			emptyGitCredentials,
			emptyWorkspace,
			emptyIpAccessLIst,
			emptySqlDashboards,
			emptySqlEndpoints,
			emptySqlQueries,
			emptyPipelines,
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups?",
				Response: scim.GroupList{
					Resources: []scim.Group{
						// TODO: add another user for which there is no filter resut
						{ID: "a", DisplayName: "admins",
							Members: []scim.ComplexValue{
								{Display: "test@test.com", Value: "123", Ref: "Users/123"},
								{Display: "Test group", Value: "f", Ref: "Groups/f"},
							},
						},
						{ID: "b", DisplayName: "users"},
						{ID: "c", DisplayName: "test"},
					},
				},
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/global-init-scripts",
				ReuseRequest: true,
				Response: map[string]any{
					"scripts": []map[string]any{},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/a",
				Response: scim.Group{ID: "a", DisplayName: "admins",
					Members: []scim.ComplexValue{
						{Display: "test@test.com", Value: "123", Ref: "Users/123"},
						{Display: "Test group", Value: "f", Ref: "Groups/f"},
					},
				},
				ReuseRequest: true,
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/b",
				Response: scim.Group{ID: "b", DisplayName: "users"},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/c",
				Response: scim.Group{ID: "c", DisplayName: "test",
					Groups: []scim.ComplexValue{
						{Display: "admins", Value: "a", Ref: "Groups/a", Type: "direct"},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups/f",
				Response: scim.Group{ID: "f", DisplayName: "nested"},
			},
			// TODO: add groups to the output
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Users/123",
				Response: scim.User{ID: "123", DisplayName: "test@test.com", UserName: "test@test.com"},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Users?filter=userName%20eq%20%27test%40test.com%27",
				Response: scim.UserList{
					Resources: []scim.User{
						{ID: "123", DisplayName: "test@test.com", UserName: "test@test.com"},
					},
					StartIndex:   1,
					TotalResults: 1,
					ItemsPerPage: 1,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/jobs/list?expand_tasks=false&limit=25&offset=0",
				Response: jobs.JobListResponse{},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/list",
				Response: clusters.ClusterList{},
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/secrets/scopes/list",
				ReuseRequest: true,
				Response: secrets.SecretScopeList{
					Scopes: []secrets.SecretScope{
						{Name: "a"},
					},
				},
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/secrets/list?scope=a",
				ReuseRequest: true,
				Response: secrets.SecretsList{
					Secrets: []secrets.SecretMetadata{
						{Key: "b"},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/secrets/acls/list?scope=a",
				Response: secrets.SecretScopeACL{
					Items: []secrets.ACLItem{
						{Permission: "MANAGE", Principal: "test"},
						{Permission: "READ", Principal: "users"},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/secrets/acls/list?scope=a",
				Response: secrets.SecretScopeACL{
					Items: []secrets.ACLItem{
						{Permission: "MANAGE", Principal: "test"},
						{Permission: "READ", Principal: "users"},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/secrets/acls/get?principal=test&scope=a",
				Response: secrets.ACLItem{Permission: "MANAGE", Principal: "test"},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/secrets/acls/get?principal=users&scope=a",
				Response: secrets.ACLItem{Permission: "READ", Principal: "users"},
			},
		}, func(ctx context.Context, client *common.DatabricksClient) {
			tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
			defer os.RemoveAll(tmpDir)

			ic := newImportContext(client)
			ic.Directory = tmpDir
			services, listing := ic.allServicesAndListing()
			ic.services = services
			ic.listing = listing

			err := ic.Run()
			assert.NoError(t, err)
		})
}

func TestImportingNoResourcesError(t *testing.T) {
	qa.HTTPFixturesApply(t,
		[]qa.HTTPFixture{
			meAdminFixture,
			emptyRepos,
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups?",
				Response: scim.GroupList{Resources: []scim.Group{}},
			},
			emptyGitCredentials,
			emptyIpAccessLIst,
			emptyWorkspace,
			emptySqlEndpoints,
			emptySqlQueries,
			emptySqlDashboards,
			emptyPipelines,
			{
				Method:       "GET",
				Resource:     "/api/2.0/global-init-scripts",
				ReuseRequest: true,
				Response: map[string]any{
					"scripts": []map[string]any{},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/jobs/list?expand_tasks=false&limit=25&offset=0",
				Response: jobs.JobListResponse{},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/list",
				Response: clusters.ClusterList{},
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/secrets/scopes/list",
				ReuseRequest: true,
				Response: secrets.SecretScopeList{
					Scopes: []secrets.SecretScope{},
				},
			},
		}, func(ctx context.Context, client *common.DatabricksClient) {
			tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
			defer os.RemoveAll(tmpDir)

			ic := newImportContext(client)
			ic.Directory = tmpDir
			services, listing := ic.allServicesAndListing()
			ic.listing = listing
			ic.services = services

			err := ic.Run()
			assert.EqualError(t, err, "no resources to import")
		})
}

func TestImportingClusters(t *testing.T) {
	qa.HTTPFixturesApply(t,
		[]qa.HTTPFixture{
			meAdminFixture,
			emptyRepos,
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups?",
				Response: scim.GroupList{Resources: []scim.Group{}},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/jobs/list?expand_tasks=false&limit=25&offset=0",
				Response: jobs.JobListResponse{},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/list",
				Response: getJSONObject("test-data/clusters-list-response.json"),
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=test1",
				Response: getJSONObject("test-data/get-cluster-test1-response.json"),
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/clusters/events",
				Response: clusters.EventDetails{},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/libraries/cluster-status?cluster_id=test1",
				Response: getJSONObject("test-data/libraries-cluster-status-test1.json"),
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/permissions/clusters/test1",
				Response: getJSONObject("test-data/get-cluster-permissions-test1-response.json"),
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/dbfs/get-status?path=dbfs%3A%2FFileStore%2Fjars%2Ftest.jar",
				ReuseRequest: true,
				Response:     getJSONObject("test-data/get-dbfs-library-status.json"),
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/dbfs/read?length=1000000&path=dbfs%3A%2FFileStore%2Fjars%2Ftest.jar",
				ReuseRequest: true,
				Response:     getJSONObject("test-data/get-dbfs-library-data.json"),
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=test2",
				Response: getJSONObject("test-data/get-cluster-test2-response.json"),
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/clusters/events",
				ExpectedRequest: clusters.EventsRequest{
					ClusterID:  "test2",
					Order:      "DESC",
					EventTypes: []clusters.ClusterEventType{"PINNED", "UNPINNED"},
					Limit:      1,
				},
				Response: clusters.EventDetails{},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/libraries/cluster-status?cluster_id=test2",
				Response: getJSONObject("test-data/libraries-cluster-status-test2.json"),
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/permissions/clusters/test2",
				Response: getJSONObject("test-data/get-cluster-permissions-test2-response.json"),
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/policies/clusters/get?policy_id=123",
				Response: getJSONObject("test-data/get-cluster-policy.json"),
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/permissions/cluster-policies/123",
				Response: getJSONObject("test-data/get-cluster-policy-permissions.json"),
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/get?cluster_id=awscluster",
				Response: getJSONObject("test-data/get-cluster-awscluster-response.json"),
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/clusters/events",
				ExpectedRequest: clusters.EventsRequest{
					ClusterID:  "awscluster",
					Order:      "DESC",
					EventTypes: []clusters.ClusterEventType{"PINNED", "UNPINNED"},
					Limit:      1,
				},
				Response: clusters.EventDetails{},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/libraries/cluster-status?cluster_id=awscluster",
				Response: getJSONObject("test-data/libraries-cluster-status-test2.json"),
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/permissions/clusters/awscluster",
				Response: getJSONObject("test-data/get-cluster-permissions-awscluster-response.json"),
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/instance-profiles/list",
				Response: getJSONObject("test-data/list-instance-profiles.json"),
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/preview/scim/v2/Me",
				ReuseRequest: true,
				Response:     scim.User{ID: "a", DisplayName: "test@test.com"},
			},
		},
		func(ctx context.Context, client *common.DatabricksClient) {
			tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
			defer os.RemoveAll(tmpDir)

			ic := newImportContext(client)
			ic.Directory = tmpDir
			ic.listing = "compute"
			services, _ := ic.allServicesAndListing()
			ic.services = services

			err := ic.Run()
			assert.NoError(t, err)
		})
}

func TestImportingJobs_JobList(t *testing.T) {
	nowSeconds := time.Now().Unix()
	jobRuns := jobs.JobRunsList{
		Runs: []jobs.JobRun{
			{
				StartTime: nowSeconds * 1000,
			},
		},
	}
	qa.HTTPFixturesApply(t,
		[]qa.HTTPFixture{
			meAdminFixture,
			emptyRepos,
			{
				Method:   "GET",
				Resource: "/api/2.1/jobs/list?expand_tasks=false&limit=25&offset=0",
				Response: jobs.JobListResponse{
					Jobs: []jobs.Job{
						{
							JobID: 14,
							Settings: &jobs.JobSettings{
								Name: "Demo job",
							},
						},
						{
							JobID: 15,
							Settings: &jobs.JobSettings{
								Name: "Demo job",
							},
						},
						{
							JobID: 16,
							Settings: &jobs.JobSettings{
								Name: "Demo job",
							},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/permissions/jobs/14",
				Response: getJSONObject("test-data/get-job-14-permissions.json"),
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/dbfs/get-status?path=dbfs%3A%2FFileStore%2Fjars%2Ftest.jar",
				ReuseRequest: true,
				Response:     getJSONObject("test-data/get-dbfs-library-status.json"),
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/dbfs/read?length=1000000&path=dbfs%3A%2FFileStore%2Fjars%2Ftest.jar",
				ReuseRequest: true,
				Response:     getJSONObject("test-data/get-dbfs-library-data.json"),
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/instance-pools/get?instance_pool_id=pool1",
				ReuseRequest: true,
				Response:     getJSONObject("test-data/get-instance-pool1.json"),
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/permissions/instance-pools/pool1",
				ReuseRequest: true,
				Response:     getJSONObject("test-data/get-job-14-permissions.json"),
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/jobs/get?job_id=14",
				Response: jobs.Job{
					JobID: 14,
					Settings: &jobs.JobSettings{
						RetryOnTimeout: true,
						Libraries: []libraries.Library{
							{Jar: "dbfs:/FileStore/jars/test.jar"},
						},
						Name: "Dummy",
						NewCluster: &clusters.Cluster{
							InstancePoolID: "pool1",
							NumWorkers:     2,
							SparkVersion:   "6.4.x-scala2.11",
							PolicyID:       "123",
						},
						SparkJarTask: &jobs.SparkJarTask{
							JarURI:        "dbfs:/FileStore/jars/test.jar",
							MainClassName: "com.databricks.examples.ProjectDriver",
						},
						SparkPythonTask: &jobs.SparkPythonTask{
							// this makes no sense for prod, but does for tests ;-)
							PythonFile: "/foo/bar.py",
							Parameters: []string{
								"dbfs:/FileStore/jars/test.jar",
								"etc",
							},
						},
						NotebookTask: &jobs.NotebookTask{
							NotebookPath: "/Test",
						},
						PipelineTask: &jobs.PipelineTask{
							PipelineID: "123",
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/policies/clusters/get?policy_id=123",
				Response: policies.ClusterPolicy{
					PolicyID: "123",
					Name:     "dummy",
					Definition: `{
						"aws_attributes.instance_profile_arn": {
							"type": "fixed",
							"value": "arn:aws:iam::12345:instance-profile/shard-s3-access",
							"hidden": true
						},
						"instance_pool_id": {
							"type": "fixed",
							"value": "pool1",
							"hidden": true
						}
					}`,
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/instance-profiles/list",
				Response: aws.InstanceProfileList{
					InstanceProfiles: []aws.InstanceProfileInfo{
						{
							InstanceProfileArn: "arn:aws:iam::12345:instance-profile/shard-s3-access",
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/permissions/cluster-policies/123",
				Response: getJSONObject("test-data/get-cluster-policy-permissions.json"),
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/instance-profiles/list",
				Response: getJSONObject("test-data/list-instance-profiles.json"),
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/instance-pools/get?instance_pool_id=pool1",
				ReuseRequest: true,
				Response:     getJSONObject("test-data/get-instance-pool1.json"),
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/permissions/instance-pools/pool1",
				ReuseRequest: true,
				Response:     getJSONObject("test-data/get-job-14-permissions.json"),
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/jobs/runs/list?completed_only=true&job_id=14&limit=1",
				Response: jobRuns,
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/jobs/runs/list?completed_only=true&job_id=15&limit=1",
				Response: jobs.JobRunsList{
					Runs: []jobs.JobRun{},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/jobs/runs/list?completed_only=true&job_id=16&limit=1",
				Response: jobs.JobRunsList{
					Runs: []jobs.JobRun{
						{
							StartTime: 0,
						},
					},
				},
			},
		},
		func(ctx context.Context, client *common.DatabricksClient) {
			ic := newImportContext(client)
			ic.services = "jobs,access,storage,clusters"
			ic.listing = "jobs"
			ic.mounts = true
			ic.meAdmin = true
			tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
			defer os.RemoveAll(tmpDir)
			ic.Directory = tmpDir

			err := ic.Importables["databricks_job"].List(ic)
			assert.NoError(t, err)

			for _, res := range ic.Scope {
				if res.Resource != "databricks_job" {
					continue
				}
				// simulate complex HCL write
				err = ic.dataToHcl(
					ic.Importables["databricks_job"],
					[]string{},
					ic.Resources["databricks_job"],
					res.Data,
					hclwrite.NewEmptyFile().Body())

				assert.NoError(t, err)
			}
		})
}

func TestImportingJobs_JobListMultiTask(t *testing.T) {
	nowSeconds := time.Now().Unix()
	jobRuns := jobs.JobRunsList{
		Runs: []jobs.JobRun{
			{
				StartTime: nowSeconds * 1000,
			},
		},
	}
	qa.HTTPFixturesApply(t,
		[]qa.HTTPFixture{
			meAdminFixture,
			emptyRepos,
			{
				Method:   "GET",
				Resource: "/api/2.1/jobs/list?expand_tasks=false&limit=25&offset=0",
				Response: jobs.JobListResponse{
					Jobs: []jobs.Job{
						{
							JobID: 14,
							Settings: &jobs.JobSettings{
								Name: "Demo job",
							},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/permissions/jobs/14",
				Response: getJSONObject("test-data/get-job-14-permissions.json"),
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/dbfs/get-status?path=dbfs%3A%2FFileStore%2Fjars%2Ftest.jar",
				ReuseRequest: true,
				Response:     getJSONObject("test-data/get-dbfs-library-status.json"),
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/dbfs/read?length=1000000&path=dbfs%3A%2FFileStore%2Fjars%2Ftest.jar",
				ReuseRequest: true,
				Response:     getJSONObject("test-data/get-dbfs-library-data.json"),
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/instance-pools/get?instance_pool_id=pool1",
				ReuseRequest: true,
				Response:     getJSONObject("test-data/get-instance-pool1.json"),
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/permissions/instance-pools/pool1",
				ReuseRequest: true,
				Response:     getJSONObject("test-data/get-job-14-permissions.json"),
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/jobs/get?job_id=14",
				Response: jobs.Job{
					JobID: 14,
					Settings: &jobs.JobSettings{
						RetryOnTimeout: true,
						Tasks: []jobs.JobTaskSettings{
							{
								TaskKey: "dummy",
								Libraries: []libraries.Library{
									{Jar: "dbfs:/FileStore/jars/test.jar"},
								},
								NewCluster: &clusters.Cluster{
									InstancePoolID:       "pool1",
									DriverInstancePoolID: "pool1",
									NumWorkers:           2,
									SparkVersion:         "6.4.x-scala2.11",
									PolicyID:             "123",
								},
								SparkJarTask: &jobs.SparkJarTask{
									JarURI:        "dbfs:/FileStore/jars/test.jar",
									MainClassName: "com.databricks.examples.ProjectDriver",
								},
								SparkPythonTask: &jobs.SparkPythonTask{
									// this makes no sense for prod, but does for tests ;-)
									PythonFile: "/foo/bar.py",
									Parameters: []string{
										"dbfs:/FileStore/jars/test.jar",
										"etc",
									},
								},
								NotebookTask: &jobs.NotebookTask{
									NotebookPath: "/Test",
								},
								PipelineTask: &jobs.PipelineTask{
									PipelineID: "123",
								},
								SqlTask: &jobs.SqlTask{
									Dashboard: &jobs.SqlDashboardTask{
										DashboardID: "123",
									},
									WarehouseID: "123",
								},
								DbtTask: &jobs.DbtTask{
									WarehouseId: "123",
									Commands:    []string{"dbt init"},
								},
							},
							{
								TaskKey: "dummy2",
								SqlTask: &jobs.SqlTask{
									Query: &jobs.SqlQueryTask{
										QueryID: "123",
									},
								},
							},
						},
						Name:   "Dummy",
						Format: "MULTI_TASK",
						JobClusters: []jobs.JobCluster{
							{
								JobClusterKey: "shared",
								NewCluster: &clusters.Cluster{
									InstancePoolID: "pool1",
									NumWorkers:     2,
									SparkVersion:   "6.4.x-scala2.11",
									PolicyID:       "123",
								},
							},
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/policies/clusters/get?policy_id=123",
				Response: policies.ClusterPolicy{
					PolicyID: "123",
					Name:     "dummy",
					Definition: `{
						"aws_attributes.instance_profile_arn": {
							"type": "fixed",
							"value": "arn:aws:iam::12345:instance-profile/shard-s3-access",
							"hidden": true
						},
						"instance_pool_id": {
							"type": "fixed",
							"value": "pool1",
							"hidden": true
						}
					}`,
				},
			},
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/instance-profiles/list",
				Response: aws.InstanceProfileList{
					InstanceProfiles: []aws.InstanceProfileInfo{
						{
							InstanceProfileArn: "arn:aws:iam::12345:instance-profile/shard-s3-access",
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/permissions/cluster-policies/123",
				Response: getJSONObject("test-data/get-cluster-policy-permissions.json"),
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/instance-profiles/list",
				Response: getJSONObject("test-data/list-instance-profiles.json"),
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/instance-pools/get?instance_pool_id=pool1",
				ReuseRequest: true,
				Response:     getJSONObject("test-data/get-instance-pool1.json"),
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/permissions/instance-pools/pool1",
				ReuseRequest: true,
				Response:     getJSONObject("test-data/get-job-14-permissions.json"),
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/jobs/runs/list?completed_only=true&job_id=14&limit=1",
				Response: jobRuns,
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/jobs/runs/list?completed_only=true&job_id=15&limit=1",
				Response: jobs.JobRunsList{
					Runs: []jobs.JobRun{},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/jobs/runs/list?completed_only=true&job_id=16&limit=1",
				Response: jobs.JobRunsList{
					Runs: []jobs.JobRun{
						{
							StartTime: 0,
						},
					},
				},
			},
		},
		func(ctx context.Context, client *common.DatabricksClient) {
			ic := newImportContext(client)
			ic.services = "jobs,access,storage,clusters"
			ic.listing = "jobs"
			ic.mounts = true
			ic.meAdmin = true
			tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
			defer os.RemoveAll(tmpDir)
			ic.Directory = tmpDir

			err := ic.Importables["databricks_job"].List(ic)
			assert.NoError(t, err)

			for _, res := range ic.Scope {
				if res.Resource != "databricks_job" {
					continue
				}
				// simulate complex HCL write
				err = ic.dataToHcl(
					ic.Importables["databricks_job"],
					[]string{},
					ic.Resources["databricks_job"],
					res.Data,
					hclwrite.NewEmptyFile().Body())

				assert.NoError(t, err)
			}
		})
}

func TestImportingWithError(t *testing.T) {
	err := Run("-directory", "/bin/sh", "-services", "groups,users", "-skip-interactive")
	assert.EqualError(t, err, "the path /bin/sh is not a directory")

	err = Run("-directory", "/bin/abcd", "-services", "groups,users", "-prefix", "abc", "-skip-interactive")
	assert.EqualError(t, err, "can't create directory /bin/abcd")
}

func TestImportingSecrets(t *testing.T) {
	qa.HTTPFixturesApply(t,
		[]qa.HTTPFixture{
			meAdminFixture,
			emptyRepos,
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups?",
				Response: scim.GroupList{Resources: []scim.Group{}},
			},
			{
				Method:   "GET",
				Resource: "/api/2.1/jobs/list?expand_tasks=false&limit=25&offset=0",
				Response: jobs.JobListResponse{},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/clusters/list",
				Response: clusters.ClusterList{},
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/secrets/scopes/list",
				ReuseRequest: true,
				Response:     getJSONObject("test-data/secret-scopes-response.json"),
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/secrets/list?scope=some-kv-scope",
				ReuseRequest: true,
				Response:     getJSONObject("test-data/secret-scopes-list-scope-response.json"),
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/secrets/acls/list?scope=some-kv-scope",
				ReuseRequest: true,
				Response:     getJSONObject("test-data/secret-scopes-list-scope-acls-response.json"),
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/secrets/acls/get?principal=test%40test.com&scope=some-kv-scope",
				ReuseRequest: true,
				Response:     getJSONObject("test-data/secret-scopes-get-principal-response.json"),
			},
		}, func(ctx context.Context, client *common.DatabricksClient) {
			tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
			defer os.RemoveAll(tmpDir)

			ic := newImportContext(client)
			ic.Directory = tmpDir
			ic.listing = "secrets"
			services, _ := ic.allServicesAndListing()
			ic.services = services
			ic.generateDeclaration = true

			err := ic.Run()
			assert.NoError(t, err)
		})
}

func TestResourceName(t *testing.T) {
	ic := newImportContext(&common.DatabricksClient{})
	norm := ic.ResourceName(&resource{
		Name: "9721431b_bcd3_4526_b90f_f5de2befec8c-dbutils_extensions_2_11_0_0_1-18dc8.jar",
	})
	assert.Equal(t, "dbutils_extensions_jar", norm)

	norm = ic.ResourceName(&resource{
		Name: "9721431b_bcd3_4526_b90f_f5de2befec8c|8737798193",
	})
	assert.Equal(t, "r7322b058678", norm)

	norm = ic.ResourceName(&resource{
		Name: "General Policy - All Users",
	})
	assert.Equal(t, "general_policy_all_users", norm)
}

func TestImportingGlobalInitScripts(t *testing.T) {
	qa.HTTPFixturesApply(t,
		[]qa.HTTPFixture{
			meAdminFixture,
			emptyRepos,
			{
				Method:       "GET",
				Resource:     "/api/2.0/global-init-scripts",
				ReuseRequest: true,
				Response:     getJSONObject("test-data/global-init-scripts-list.json"),
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/global-init-scripts/C39FD6BAC8088BBC",
				ReuseRequest: true,
				Response:     getJSONObject("test-data/global-init-script-get1.json"),
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/global-init-scripts/F931E63C248C1D8C",
				ReuseRequest: true,
				Response:     getJSONObject("test-data/global-init-script-get2.json"),
			},
		}, func(ctx context.Context, client *common.DatabricksClient) {
			tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
			defer os.RemoveAll(tmpDir)

			ic := newImportContext(client)
			ic.Directory = tmpDir
			ic.listing = "workspace"
			services, _ := ic.allServicesAndListing()
			ic.services = services
			ic.generateDeclaration = true

			err := ic.Run()
			assert.NoError(t, err)
		})
}

func TestImportingUser(t *testing.T) {
	qa.HTTPFixturesApply(t,
		[]qa.HTTPFixture{
			{
				Method:       "GET",
				ReuseRequest: true,
				Resource:     "/api/2.0/preview/scim/v2/Users?filter=userName%20eq%20%27me%27",
				Response: scim.UserList{
					Resources: []scim.User{
						{
							ID:       "123",
							UserName: "me",
							Groups: []scim.ComplexValue{
								{
									Value: "abc",
									Type:  "direct",
								},
							},
						},
					},
				},
			},
		}, func(ctx context.Context, client *common.DatabricksClient) {
			ic := newImportContext(client)
			err := resourcesMap["databricks_user"].Search(ic, &resource{
				Resource: "databricks_user",
				Value:    "me",
			})
			assert.NoError(t, err)

			d := ic.Resources["databricks_user"].TestResourceData()
			d.Set("user_name", "me")
			err = resourcesMap["databricks_user"].Import(ic, &resource{
				Resource: "databricks_user",
				Data:     d,
			})
			assert.NoError(t, err)
		})
}

func TestEitherString(t *testing.T) {
	assert.Equal(t, "a", eitherString("a", nil))
	assert.Equal(t, "a", eitherString(nil, "a"))
	assert.Equal(t, "", eitherString(nil, nil))
}

func TestImportingRepos(t *testing.T) {
	resp := repos.ReposInformation{
		ID:           121232342,
		Url:          "https://github.com/user/test.git",
		Provider:     "gitHub",
		Path:         "/Repos/user@domain/test",
		HeadCommitID: "1124323423abc23424",
		Branch:       "releases",
	}

	qa.HTTPFixturesApply(t,
		[]qa.HTTPFixture{
			meAdminFixture,
			{
				Method:   "GET",
				Resource: "/api/2.0/repos?",
				Response: repos.ReposListResponse{
					Repos: []repos.ReposInformation{
						resp,
					},
				},
			},
			emptyGitCredentials,
			{
				Method:   "GET",
				Resource: "/api/2.0/repos/121232342",
				Response: resp,
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/permissions/repos/121232342",
				Response: getJSONObject("test-data/get-repo-permissions.json"),
			},
		},
		func(ctx context.Context, client *common.DatabricksClient) {
			tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
			defer os.RemoveAll(tmpDir)

			ic := newImportContext(client)
			ic.Directory = tmpDir
			ic.listing = "repos"
			ic.services = "repos"

			err := ic.Run()
			assert.NoError(t, err)
		})
}

func TestImportingIPAccessLists(t *testing.T) {
	resp := access.IpAccessListStatus{
		ListID:       "123",
		Label:        "block_list",
		ListType:     "BLOCK",
		IPAddresses:  []string{"1.2.3.4"},
		AddressCount: 2,
		Enabled:      true,
	}
	resp2 := resp
	resp2.IPAddresses = []string{}
	resp2.ListID = "124"
	qa.HTTPFixturesApply(t,
		[]qa.HTTPFixture{
			meAdminFixture,
			emptyRepos,
			{
				Method:   "GET",
				Resource: "/api/2.0/global-init-scripts",
				Response: map[string]any{},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/ip-access-lists",
				Response: access.ListIPAccessListsResponse{
					ListIPAccessListsResponse: []access.IpAccessListStatus{resp, resp2},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/ip-access-lists/123",
				Response: access.IpAccessListStatusWrapper{
					IPAccessList: resp,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/ip-access-lists/124",
				Response: access.IpAccessListStatusWrapper{
					IPAccessList: resp2,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/workspace-conf?keys=enableIpAccessLists%2CenableTokensConfig%2CmaxTokenLifetimeDays",
				Response: map[string]any{
					"enableIpAccessLists":  "true",
					"maxTokenLifetimeDays": "90",
					"enableTokensConfig":   "true",
				},
				ReuseRequest: true,
			},
		},
		func(ctx context.Context, client *common.DatabricksClient) {
			tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
			defer os.RemoveAll(tmpDir)

			ic := newImportContext(client)
			ic.Directory = tmpDir
			ic.listing = "workspace,access"
			ic.services = "workspace,access"

			err := ic.Run()
			assert.NoError(t, err)
		})
}

func TestImportingSqlObjects(t *testing.T) {
	qa.HTTPFixturesApply(t,
		[]qa.HTTPFixture{
			meAdminFixture,
			emptyRepos,
			emptyIpAccessLIst,
			{
				Method:   "GET",
				Resource: "/api/2.0/global-init-scripts",
				Response: map[string]any{},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/sql/warehouses",
				Response: getJSONObject("test-data/get-sql-endpoints.json"),
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/sql/warehouses/f562046bc1272886",
				Response: getJSONObject("test-data/get-sql-endpoint.json"),
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/data_sources",
				Response: []sql.DataSource{
					{
						ID:         "147164a6-8316-4a9d-beff-f57261801374",
						EndpointID: "f562046bc1272886",
					},
				},
				ReuseRequest: true,
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/permissions/sql/warehouses/f562046bc1272886",
				Response: getJSONObject("test-data/get-sql-endpoint-permissions.json"),
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/preview/sql/dashboards",
				Response:     getJSONObject("test-data/get-sql-dashboards.json"),
				ReuseRequest: true,
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/preview/sql/dashboards/9cb0c8f5-6262-4a1f-a741-2181de76028f",
				Response:     getJSONObject("test-data/get-sql-dashboard.json"),
				ReuseRequest: true,
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/preview/sql/queries",
				Response:     getJSONObject("test-data/get-sql-queries.json"),
				ReuseRequest: true,
			},
			{
				Method:       "GET",
				Resource:     "/api/2.0/preview/sql/queries/16c4f969-eea0-4aad-8f82-03d79b078dcc",
				Response:     getJSONObject("test-data/get-sql-query.json"),
				ReuseRequest: true,
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/permissions/queries/16c4f969-eea0-4aad-8f82-03d79b078dcc",
				Response: getJSONObject("test-data/get-sql-query-permissions.json"),
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/sql/permissions/dashboards/9cb0c8f5-6262-4a1f-a741-2181de76028f",
				Response: getJSONObject("test-data/get-sql-dashboard-permissions.json"),
			},
		},
		func(ctx context.Context, client *common.DatabricksClient) {
			tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
			defer os.RemoveAll(tmpDir)

			ic := newImportContext(client)
			ic.Directory = tmpDir
			ic.listing = "sql,access"
			ic.services = "sql,access"

			err := ic.Run()
			assert.NoError(t, err)
		})
}

func TestImportingDLTPipelines(t *testing.T) {
	qa.HTTPFixturesApply(t,
		[]qa.HTTPFixture{
			meAdminFixture,
			emptyRepos,
			emptyIpAccessLIst,
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines?max_results=50",

				Response: pipelines.PipelineListResponse{
					Statuses: []pipelines.PipelineStateInfo{
						{
							PipelineID: "123",
							Name:       "Pipeline1",
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines/123",
				Response: getJSONObject("test-data/get-dlt-pipeline.json"),
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/permissions/pipelines/123",
				Response: getJSONObject("test-data/get-pipeline-permissions.json"),
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/workspace/get-status?path=%2FUsers%2Fuser%40domain.com%2FTest%20DLT",
				Response: workspace.ObjectStatus{
					Language:   workspace.Python,
					ObjectID:   123,
					ObjectType: workspace.Notebook,
					Path:       "/Users/user@domain.com/Test DLT",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/workspace/export?format=SOURCE&path=%2FUsers%2Fuser%40domain.com%2FTest+DLT",
				Response: workspace.ExportPath{
					Content: "spark.range(10)",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/instance-profiles/list",
				Response: getJSONObject("test-data/list-instance-profiles.json"),
			},
		},
		func(ctx context.Context, client *common.DatabricksClient) {
			tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
			defer os.RemoveAll(tmpDir)

			ic := newImportContext(client)
			ic.Directory = tmpDir
			ic.listing = "dlt"
			ic.services = "dlt,access,notebooks"

			err := ic.Run()
			assert.NoError(t, err)
		})
}

func TestImportingDLTPipelinesMatchingOnly(t *testing.T) {
	qa.HTTPFixturesApply(t,
		[]qa.HTTPFixture{
			meAdminFixture,
			emptyRepos,
			emptyIpAccessLIst,
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines?max_results=50",

				Response: pipelines.PipelineListResponse{
					Statuses: []pipelines.PipelineStateInfo{
						{
							PipelineID: "123",
							Name:       "Pipeline1 test",
						},
						{
							PipelineID: "124",
							Name:       "Pipeline1",
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines/123",
				Response: getJSONObject("test-data/get-dlt-pipeline.json"),
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/permissions/pipelines/123",
				Response: getJSONObject("test-data/get-pipeline-permissions.json"),
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/instance-profiles/list",
				Response: getJSONObject("test-data/list-instance-profiles.json"),
			},
		},
		func(ctx context.Context, client *common.DatabricksClient) {
			tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
			defer os.RemoveAll(tmpDir)

			ic := newImportContext(client)
			ic.Directory = tmpDir
			ic.match = "test"
			ic.listing = "dlt"
			ic.services = "dlt,access"

			err := ic.Run()
			assert.NoError(t, err)
		})
}
