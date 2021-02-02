# Integration Testing

The `run.sh` script will setup the required resources in specific clouds, if they are defined in `main.tf` and then pass these as environment variables to the `golang` integration tests. You can create `require_env` files with one line per environment variable and expect to have those definitely set for integration test run because of validation performed inside of it.

# Destruction

By default, we don't encourage creation/destruction of infrastructure multiple times per day, because in some cases it may take up to 30 minutes just to boot everything up. Therefore only `--destroy` flag in `run.sh` will trigger `terraform destroy -auto-approve`. `make test-*` tasks won't explicitly request destruction by default.

# Conventions

* `azcli` - Azure authenticated with `az login` command. No `require_env` file needed. Runnable test name prefixes are `TestAcc` and `TestAzureAcc`. By far, the simplest way to develop provider's functionality.
* `azsp` - Azure authenticated with Service Principal's ID/Secret pairs. Runnable test name prefixes are `TestAcc` and `TestAzureAcc`. Service pricipal must have `Storage Blob Data Contributor` role on ADLS account used. `ARM_SUBSCRIPTION_ID`, `ARM_CLIENT_SECRET`, `ARM_CLIENT_ID`, `ARM_TENANT_ID`, `OWNER` environment vars required. Note that these integration tests will use service principal based auth. Even though it is using a service principal, it will still be generating a personal access token to perform creation of resources.

* `mws` - AWS with Databricks Multiworkspace API. Runnable test name prefix is `TestMws`. Please [check if you're able to use it](https://docs.databricks.com/administration-guide/multiworkspace/new-workspace-aws.html). Required variables are `DATABRICKS_ACCOUNT_ID`, `DATABRICKS_USERNAME`, `DATABRICKS_PASSWORD` (something you use for https://accounts.cloud.databricks.com/), `TEST_REGION`, `TEST_CIDR`, `OWNER`. Only multiworkspace resources are tested.
* `awsst` - `DATABRICKS_CONFIG_PROFILE` (section within Databricks CLI `~/.databrickscfg` file) & `CLOUD_ENV=AWS`. In case you want to test provider on existing development single-tenant shard. Runnable test name prefixes are `TestAcc` and `TestAwsAcc`.
* `awsmt` - AWS with Databricks Multitenant Workspace. Currently work in progress and the test environment cannot be fully started.
* most of the tests should aim to be cloud-agnostic. Though, in case of specific branching needed, you can check `CLOUD_ENV` value (possible values are `Azure`, `AWS` & `MWS`).
* all environment variables are used by *DatabricksClient*, *provider integration tests* and *terraform configuration*.
* **each `output` becomes an environment variable** with the case changed to upper. This gives an easy way to manage the complexity of the testing environment. This is what gives those variables for `export $(scripts/run.sh azcli --export)` under the hood.
* `qa.EnvironmentTemplate` must be used to make readable templates with environment variable presence validation.
* `OWNER` is a variable name that holds your email address. It's propagated down to all resourced on the cloud.
* One must aim to write integration tests that will run on all clouds without causing panic under any circumstance.

# Development loop

```bash
# keep credentials outside of repository
source ~/path/to-env-files.sh

# export all related environment vars - a.k.a. activating environment
export $(scripts/run.sh azcli --export)

# launch VSCode with those environment vars, so that debugging is more seamless
code .

# run all tests
scripts/run.sh azcli '^(TestAcc|TestAzureAcc)' --debug

# run individual test
scripts/run.sh azcli TestAccContext --debug

# run all tests
make test-azcli
```

# Code

Example of environment-aware test code is:

```golang
func TestMwsAccStorageConfigurations(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	if cloudEnv != "MWS" {
		t.Skip("Cannot run test on non-MWS environment")
	}
	var bucket MWSStorageConfigurations
	config := qa.EnvironmentTemplate(t, `
	provider "databricks" {
		host     = "{env.DATABRICKS_HOST}"
		username = "{env.DATABRICKS_USERNAME}"
		password = "{env.DATABRICKS_PASSWORD}"
	}
	resource "databricks_mws_storage_configurations" "this" {
		account_id                 = "{env.DATABRICKS_ACCOUNT_ID}"
		storage_configuration_name = "terraform-{var.RANDOM}"
		bucket_name                = "terraform-{var.RANDOM}"
	  }
	`)
	bucketName := qa.FirstKeyValue(t, config, "bucket_name")
    configName := qa.FirstKeyValue(t, config, "storage_configuration_name")
    ...
```

## Cloud Specific testing

Basic cloud-integration test should have a prefix `TestAcc` if it is supposed to run on both clouds. Client must be created with `NewClientFromEnvironment()` as described in the following snippet:

```go
func TestAccListClustersIntegration(t *testing.T) {
	cloud := os.Getenv("CLOUD_ENV")
	if cloud == "" {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	client := common.NewClientFromEnvironment()

	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	cluster := Cluster{
		NumWorkers:  1,
		ClusterName: "my-cluster-" + randomName,
		SparkVersion:           "6.2.x-scala2.11",
		NodeTypeID:             qa.GetCloudInstanceType(client),
		DriverNodeTypeID:       qa.GetCloudInstanceType(client),
		IdempotencyToken:       "cluster-" + randomName,
		AutoterminationMinutes: 15,
	}

	if cloud == "AWS" {
		cluster.AwsAttributes = &AwsAttributes{
			EbsVolumeType:  EbsVolumeTypeGeneralPurposeSsd,
			EbsVolumeCount: 1,
			EbsVolumeSize:  32,
		}
	}

	// ...
}
```