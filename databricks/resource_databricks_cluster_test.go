package databricks

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/stretchr/testify/assert"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
)

func TestClusterRead_Missing(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			ExpectedRequest: model.ClusterIDRequest{
				ClusterID: "abc",
			},
			Response: service.APIErrorBody{
				ErrorCode: "SOME_CODE",
				Message:   "Cluster abc does not exist",
			},
			Status: 404,
		},
	}, resourceCluster, nil, func(d *schema.ResourceData, c interface{}) error {
		d.SetId("abc")
		return resourceClusterRead(d, c)
	})
	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id())
}

func TestClusterRead_JustErr(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			ExpectedRequest: model.ClusterIDRequest{
				ClusterID: "abc",
			},
			Response: service.APIErrorBody{
				ErrorCode: "SOME_CODE",
				Message:   "Bad request",
			},
			Status: 400,
		},
	}, resourceCluster, nil, func(d *schema.ResourceData, c interface{}) error {
		d.SetId("abc")
		return resourceClusterRead(d, c)
	})
	assert.Error(t, err)
	assert.Equal(t, "abc", d.Id())
}

func TestClusterRead_LibrariesReturnError(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			ExpectedRequest: model.ClusterIDRequest{
				ClusterID: "abc",
			},
			Response: model.ClusterInfo{
				NumWorkers:  3,
				ClusterName: "Shared Cluster",
			},
		},
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/libraries/cluster-status?cluster_id=abc",
			ExpectedRequest: model.ClusterIDRequest{
				ClusterID: "abc",
			},
			Response: service.APIErrorBody{
				ErrorCode: "SOME_CODE",
				Message:   "Bad request",
			},
			Status: 400,
		},
	}, resourceCluster, nil, func(d *schema.ResourceData, c interface{}) error {
		d.SetId("abc")
		return resourceClusterRead(d, c)
	})
	assert.Error(t, err, err)
	assert.Equal(t, "abc", d.Id())
}

func TestClusterRead_Some(t *testing.T) {
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			ExpectedRequest: model.ClusterIDRequest{
				ClusterID: "abc",
			},
			Response: model.ClusterInfo{
				ClusterName: "Shared Cluster",
				NumWorkers:  10, // doesn't make sense, but let's try
				Autoscale: &model.AutoScale{
					MinWorkers: 10,
					MaxWorkers: 45,
				},
				AutoterminationMinutes: 20,
				SparkVersion:           "x.y.z",
				SparkConf: map[string]string{
					"spark.sql.shuffle.partitions": "1000",
				},
				DriverNodeTypeID: "i3.xlarge",
				NodeTypeID:       "i3.xlarge",
				SSHPublicKeys: []string{
					"abc",
					"bcd",
				},
				CustomTags: map[string]string{
					"Team": "EMEA",
				},
				ClusterLogConf: &model.StorageInfo{
					S3: &model.S3StorageInfo{
						Destination:      "s3a://foo",
						EnableEncryption: true,
					},
				},
				InitScripts: []model.StorageInfo{
					{
						Dbfs: &model.DbfsStorageInfo{
							Destination: "dbfs://foo/bar.sh",
						},
					},
					{
						S3: &model.S3StorageInfo{
							Destination: "s3a://baz/foo.sh",
							CannedACL:   "abc",
						},
					},
				},
				DockerImage: &model.DockerImage{
					URL: "https://foo.bar",
					BasicAuth: &model.DockerBasicAuth{
						Username: "foo",
						Password: "br",
					},
				},
				SparkEnvVars: map[string]string{
					"FOO": "BAR",
				},
				ClusterID: "abc",
			},
		},
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/libraries/cluster-status?cluster_id=abc",
			ExpectedRequest: model.ClusterIDRequest{
				ClusterID: "abc",
			},
			Response: model.ClusterLibraryStatuses{
				ClusterID: "abc",
				LibraryStatuses: []model.LibraryStatus{
					{
						Library: &model.Library{
							Maven: &model.Maven{
								Coordinates: "foo:bar.baz",
							},
						},
						Status: "INSTALLED",
					},
				},
			},
		},
	}, resourceCluster, nil, func(d *schema.ResourceData, c interface{}) error {
		d.SetId("abc")
		return resourceClusterRead(d, c)
	})
	assert.NoError(t, err, err)
	assert.Equal(t, map[string]string{
		"autoscale.#":                               "1",
		"autoscale.901678046.max_workers":           "45",
		"autoscale.901678046.min_workers":           "10",
		"autotermination_minutes":                   "20",
		"cluster_id":                                "abc",
		"cluster_log_conf.#":                        "1",
		"cluster_log_conf.0.dbfs.#":                 "0",
		"cluster_log_conf.0.s3.#":                   "1",
		"cluster_log_conf.0.s3.0.canned_acl":        "",
		"cluster_log_conf.0.s3.0.destination":       "s3a://foo",
		"cluster_log_conf.0.s3.0.enable_encryption": "true",
		"cluster_log_conf.0.s3.0.encryption_type":   "",
		"cluster_log_conf.0.s3.0.endpoint":          "",
		"cluster_log_conf.0.s3.0.kms_key":           "",
		"cluster_log_conf.0.s3.0.region":            "",
		"cluster_name":                              "Shared Cluster",
		"custom_tags.%":                             "1",
		"custom_tags.Team":                          "EMEA",
		"default_tags.%":                            "0",
		"docker_image.#":                            "1",
		"docker_image.0.basic_auth.#":               "1",
		"docker_image.0.basic_auth.0.password":      "br",
		"docker_image.0.basic_auth.0.username":      "foo",
		"docker_image.0.url":                        "https://foo.bar",
		"driver_node_type_id":                       "i3.xlarge",
		"enable_elastic_disk":                       "false",
		"id":                                        "abc",
		"init_scripts.#":                            "2",
		"init_scripts.0.dbfs.#":                     "1",
		"init_scripts.0.dbfs.0.destination":         "dbfs://foo/bar.sh",
		"init_scripts.0.s3.#":                       "0",
		"init_scripts.1.dbfs.#":                     "0",
		"init_scripts.1.s3.#":                       "1",
		"init_scripts.1.s3.0.canned_acl":            "abc",
		"init_scripts.1.s3.0.destination":           "s3a://baz/foo.sh",
		"init_scripts.1.s3.0.enable_encryption":     "false",
		"init_scripts.1.s3.0.encryption_type":       "",
		"init_scripts.1.s3.0.endpoint":              "",
		"init_scripts.1.s3.0.kms_key":               "",
		"init_scripts.1.s3.0.region":                "",
		"instance_pool_id":                          "",
		"library_maven.#":                           "1",
		"library_maven.2048034540.coordinates":      "foo:bar.baz",
		"library_maven.2048034540.messages.#":       "0",
		"library_maven.2048034540.repo":             "",
		"library_maven.2048034540.status":           "INSTALLED",
		"node_type_id":                              "i3.xlarge",
		"num_workers":                               "10",
		"single_user_name":                          "",
		"spark_conf.%":                              "1",
		"spark_conf.spark.sql.shuffle.partitions":   "1000",
		"spark_env_vars.%":                          "1",
		"spark_env_vars.FOO":                        "BAR",
		"spark_version":                             "x.y.z",
		"ssh_public_keys.#":                         "2",
		"ssh_public_keys.1247859306":                "bcd",
		"ssh_public_keys.374767988":                 "abc",
		"state":                                     "",
		"state_message":                             "",
	}, d.State().Attributes)
}

func TestClusterCreate_Some(t *testing.T) {
	info := model.ClusterInfo{
		State:       "RUNNING",
		ClusterID:   "abc",
		ClusterName: "Shared Cluster",
		Autoscale: &model.AutoScale{
			MinWorkers: 10,
			MaxWorkers: 45,
		},
		AutoterminationMinutes: 20,
		SparkVersion:           "x.y.z",
		SparkConf: map[string]string{
			"spark.sql.shuffle.partitions": "1000",
		},
		SSHPublicKeys: []string{
			"bcd",
			"abc",
		},
		CustomTags: map[string]string{
			"Team": "EMEA",
		},
		ClusterLogConf: &model.StorageInfo{
			S3: &model.S3StorageInfo{
				Destination:      "s3a://foo",
				EnableEncryption: true,
			},
		},
		InitScripts: []model.StorageInfo{
			{
				Dbfs: &model.DbfsStorageInfo{
					Destination: "dbfs://foo/bar.sh",
				},
			},
			{
				S3: &model.S3StorageInfo{
					Destination: "s3a://baz/foo.sh",
					CannedACL:   "abc",
				},
			},
		},
		SparkEnvVars: map[string]string{
			"FOO": "BAR",
		},
		InstancePoolID: "xxx-lkll-uyhr023",
	}
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   http.MethodPost,
			Resource: "/api/2.0/clusters/create",
			ExpectedRequest: model.Cluster{
				ClusterName: "Shared Cluster",
				Autoscale: &model.AutoScale{
					MinWorkers: 10,
					MaxWorkers: 45,
				},
				AutoterminationMinutes: 20,
				SparkVersion:           "x.y.z",
				SparkConf: map[string]string{
					"spark.sql.shuffle.partitions": "1000",
				},
				SSHPublicKeys: []string{
					"bcd",
					"abc",
				},
				CustomTags: map[string]string{
					"Team": "EMEA",
				},
				ClusterLogConf: &model.StorageInfo{
					S3: &model.S3StorageInfo{
						Destination:      "s3a://foo",
						EnableEncryption: true,
					},
				},
				InitScripts: []model.StorageInfo{
					{
						Dbfs: &model.DbfsStorageInfo{
							Destination: "dbfs://foo/bar.sh",
						},
					},
					{
						S3: &model.S3StorageInfo{
							Destination: "s3a://baz/foo.sh",
							CannedACL:   "abc",
						},
					},
				},
				SparkEnvVars: map[string]string{
					"FOO": "BAR",
				},
				InstancePoolID: "xxx-lkll-uyhr023",
			},
			Response: info,
		},
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/libraries/cluster-status?cluster_id=abc",
			ExpectedRequest: model.ClusterIDRequest{
				ClusterID: "abc",
			},
			Response: model.ClusterLibraryStatuses{
				ClusterID: "abc",
				LibraryStatuses: []model.LibraryStatus{
					{
						Library: &model.Library{
							Maven: &model.Maven{
								Coordinates: "foo:bar.baz",
							},
						},
						Status: "INSTALLED",
					},
				},
			},
		},
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			ExpectedRequest: model.ClusterIDRequest{
				ClusterID: "abc",
			},
			Response: info,
		},
		{
			Method:   http.MethodPost,
			Resource: "/api/2.0/libraries/install",
			ExpectedRequest: model.ClusterLibraryList{
				ClusterID: "abc",
				Libraries: []model.Library{
					{
						Maven: &model.Maven{
							Coordinates: "foo:bar.baz",
						},
					},
				},
			},
		},
	}, resourceCluster, map[string]interface{}{
		"autoscale": []interface{}{
			map[string]interface{}{
				"min_workers": 10,
				"max_workers": 45,
			},
		},
		"autotermination_minutes": 20,
		"cluster_log_conf": []interface{}{
			map[string]interface{}{
				"s3": []interface{}{
					map[string]interface{}{
						"destination":       "s3a://foo",
						"enable_encryption": true,
					},
				},
			},
		},
		"cluster_name": "Shared Cluster",
		"custom_tags": map[string]interface{}{
			"Team": "EMEA",
		},
		"init_scripts": []interface{}{
			map[string]interface{}{
				"dbfs": []interface{}{
					map[string]interface{}{
						"destination": "dbfs://foo/bar.sh",
					},
				},
			},
			map[string]interface{}{
				"s3": []interface{}{
					map[string]interface{}{
						"destination": "s3a://baz/foo.sh",
						"canned_acl":  "abc",
					},
				},
			},
		},
		"instance_pool_id": "xxx-lkll-uyhr023",
		"library_maven": []interface{}{
			map[string]interface{}{
				"coordinates": "foo:bar.baz",
			},
		},
		"spark_conf": map[string]interface{}{
			"spark.sql.shuffle.partitions": "1000",
		},
		"spark_env_vars": map[string]interface{}{
			"FOO": "BAR",
		},
		"spark_version": "x.y.z",
		"ssh_public_keys": []interface{}{
			"abc",
			"bcd",
		},
	}, resourceClusterCreate)
	assert.NoError(t, err, err)
	assert.Equal(t, "...", d.Id())
}

func TestAccCluster_BasicLifecycle(t *testing.T) {
	var clusterID *string
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// create a cluster
				Config: fmt.Sprintf(`resource "databricks_cluster" "this" {
					cluster_name = "Terraform %[1]s"
					spark_version = "6.6.x-scala2.11"
					autotermination_minutes = 15
					node_type_id = "i3.xlarge"
					num_workers = 1
				}`, randomName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("databricks_cluster.this",
						"cluster_name", "Terraform "+randomName),
					testAccIDCallback(t, "databricks_cluster.this",
						func(client *service.DBApiClient, id string) error {
							clusterID = &id
							return nil
						}),
				),
			},
			{
				// and then resize it
				Config: fmt.Sprintf(`resource "databricks_cluster" "this" {
					cluster_name = "Terraform %[1]s"
					spark_version = "6.6.x-scala2.11"
					autotermination_minutes = 15
					node_type_id = "i3.xlarge"
					num_workers = 2
				}`, randomName),
				Check: resource.TestCheckResourceAttr("databricks_cluster.this", "num_workers", "2"),
			},
			{
				PreConfig: func() {
					// and then really delete it to emulate 404
					err := testAccProvider.Meta().(*service.DBApiClient).Clusters().PermanentDelete(*clusterID)
					assert.NoError(t, err)
				},
				// so that it's re-created again with different id
				Config: fmt.Sprintf(`resource "databricks_cluster" "this" {
					cluster_name = "Terraform %[1]s"
					spark_version = "6.6.x-scala2.11"
					autotermination_minutes = 15
					node_type_id = "i3.xlarge"
					num_workers = 1
				}`, randomName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("databricks_cluster.this", "num_workers", "1"),
					testAccIDCallback(t, "databricks_cluster.this",
						func(client *service.DBApiClient, id string) error {
							assert.NotEqual(t, id, *clusterID)
							return nil
						}),
				),
			},
		},
	})
}
