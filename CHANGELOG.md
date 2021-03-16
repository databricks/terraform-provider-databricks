# Version changelog

## 0.3.2

* Fix incorrect escaping of notebook names ([#566](https://github.com/databrickslabs/terraform-provider-databricks/pull/566))

## 0.3.1

* Added `databricks_global_init_script` resource to configure global init scripts ([#487](https://github.com/databrickslabs/terraform-provider-databricks/issues/487)).
* Added `databricks_sql_endpoint` resource ([#498](https://github.com/databrickslabs/terraform-provider-databricks/pull/498))
* Added [experimental resource exporter](https://github.com/databrickslabs/terraform-provider-databricks/blob/master/docs/guides/experimental-exporter.md) to generate configuration for entire workspace.
* Improved user-facing documentaiton ([#508](https://github.com/databrickslabs/terraform-provider-databricks/pull/508/files), [#516](https://github.com/databrickslabs/terraform-provider-databricks/pull/516), [#511](https://github.com/databrickslabs/terraform-provider-databricks/pull/511), [#504](https://github.com/databrickslabs/terraform-provider-databricks/pull/504), [#492]([Update docs in various places](https://github.com/databrickslabs/terraform-provider-databricks/pull/492)))
* Simplified authentication issues debugging ([#490](https://github.com/databrickslabs/terraform-provider-databricks/pull/490))
* Made cleaner error message for no config profile ([#491](https://github.com/databrickslabs/terraform-provider-databricks/pull/491))
* Allow tokens without comment or expiration ([#495](https://github.com/databrickslabs/terraform-provider-databricks/pull/495/files))
* Ensured consistent slashes in notebook paths for different OSes ([#500](https://github.com/databrickslabs/terraform-provider-databricks/pull/500))
* Fix error message panic in command result parsing ([#502](https://github.com/databrickslabs/terraform-provider-databricks/pull/502))
* Updated `databricks_group` data resource to allow non-alphanumeric characters in group name filter ([#507](https://github.com/databrickslabs/terraform-provider-databricks/pull/507/files))

**Behavior changes**

* Assigning any permission to `admins` would result in an error, so that behavior is consistent ([#486](https://github.com/databrickslabs/terraform-provider-databricks/issues/486)).

Updated dependency versions:

* github.com/zclconf/go-cty from 1.2.1 to 1.7.1
* github.com/Azure/go-autorest/autorest/azure/auth from 0.5.6 to 0.5.7 
* github.com/hashicorp/hcl/v2 from 2.3.0 to 2.8.2
* github.com/aws/aws-sdk-go from 1.37.1 to 1.37.11
* github.com/Azure/go-autorest/autorest from 0.11.17 to 0.11.18

## 0.3.0

* Added configurable provisioning timeout for `databricks_mws_workspaces`, so that local DNS cache issues would be more tolerated.
* Added [databricks_current_user] to simplify applying the same Terraform configuration by different users in the shared workspace for testing purposes. 
* Added client-side rate limiting to release the pressure on backend APIs and prevent client blocking ([#465](https://github.com/databrickslabs/terraform-provider-databricks/pull/465))
* Member usernames, group names and instance profile names in `databricks_group` data source are now sorted and providing consistent behavior between runs ([#449](https://github.com/databrickslabs/terraform-provider-databricks/issues/449))
* Fixed redundant multiple mounting clusters ([#445](https://github.com/databrickslabs/terraform-provider-databricks/issues/445))
* Added optional parameter azure_environment to provider config which defaults to public ([#437](https://github.com/databrickslabs/terraform-provider-databricks/pull/437)).
* Added [databricks_service_principal](https://github.com/databrickslabs/terraform-provider-databricks/pull/386) resource.
* `skip_validation` from `databricks_instance_profile` was removed and is always set to `true`.
* Added propagation of terraform version to `User-Agent` header, along with type of resource used.
* `databricks_notebook` & `databricks_dbfs_file` got new `source` field to specify location of a local file.
* `databricks_notebook` can have `language` field optional, as long as `source` is set to a file with `.py`, `.scala`, `.sql`, or `.r` extension.
* `databricks_me` data source was added to represent `user_name`, `home` & `id` of the caller user (or service principal).
* Added validation for secret scope name in `databricks_secret`, `databricks_secret_scope` and `databricks_secret_acl`. Non-compliant names may cause errors.
* Added [databricks_spark_version](https://github.com/databrickslabs/terraform-provider-databricks/issues/347) data source.
* Fixed support for [single node clusters](https://docs.databricks.com/clusters/single-node.html) support by allowing [`num_workers` to be `0`](https://github.com/databrickslabs/terraform-provider-databricks/pull/454).
* Fixed bug in destruction of IP access lists ([#426](https://github.com/databrickslabs/terraform-provider-databricks/issues/426)).
* All resource imports are now making call to corresponding Databricks API by default ([#471](https://github.com/databrickslabs/terraform-provider-databricks/issues/471)).

**Behavior changes**
* Removed deprecated `library_jar`, `library_egg`, `library_whl`, `library_pypi`, `library_cran`, and `library_maven` from `databricks_cluster` and `databricks_job` in favor of more API-transparent [library](https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs/resources/cluster#library-configuration-block) configuration block.
* Removed deprecated `notebook_path` and `notebook_base_parameters` from `databricks_job` in favor of [notebook_task](https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs/resources/job#notebook_task-configuration-block) configuration block.
* Removed deprecated `jar_uri`, `jar_main_class_name`, and `jar_parameters` from `databricks_job` in favor of [spark_jar_task](https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs/resources/job#spark_jar_task-configuration-block) configuration block.
* Removed deprecated `python_file` and `python_parameters` from `databricks_job` in favor of [spark_python_task](https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs/resources/job#spark_python_task-configuration-block) configuration block.
* Removed deprecated `spark_submit_parameters` from `databricks_job` in favor of [spark_submit_task](https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs/resources/job#spark_submit_task-configuration-block) configuration block.
* Removed deprecated `databricks_scim_user` resource in favor of [databricks_user](https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs/resources/user).
* Removed deprecated `databricks_scim_group` resource in favor of [databricks_group](https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs/resources/group).
* Removed deprecated `databricks_default_user_roles` data source in favor of [databricks_group](https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs/data-sources/group#attribute-reference) data source.
* Removed deprecated `basic_auth` and `azure_auth` provider configuration blocks in favor of [documented authentication methods](https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs).
* `format`, `overwrite`, and `mkdirs` were removed from `databricks_notebook`. To follow expected behavior of Terraform, notebooks are always overwritten.
* `skip_validation` from `databricks_instance_profile` was removed and is always set to `true` for subsequent requests.
* `databricks_mws_workspace` got `verify_workspace_runnning` removed and now validates all every deployment. In case deployment failed, it removes workspace that failed and returns error message with explanation.
* `default_tags` were removed from `databricks_instance_pool`. `disk_spec` got new attribute `disk_type`, that contains `azure_disk_volume_type` and `ebs_volume_type`. This change is made to closer reflect API structure.
* `databricks_notebook` & `databricks_dbfs_file` got `content` attribute renamed to `content_base64` and now share the same logic to work with local files.

## 0.2.9

* Fixed documentation issues.
* Added missing resource importers and test to cover it.
* Migrated build from TravisCI to GitHub Actions.
* Fixed custom `config_file` issue configuration handling ([#420](https://github.com/databrickslabs/terraform-provider-databricks/issues/420)).

**Deprecations**
* `databricks_notebook` has got `overwrite`, `mkdirs` and `format` parameters, that always have to be set to certain values in order to follow expected behavior of terraform. These fields would be removed in 0.3 and always set to proper values.
* `databricks_notebook` & `databricks_dbfs_file` field `content` is deprecated and would be renamed to `content_base64` to further increase clarity. 
* `databricks_dbfs_file` has got `content`, `content_b64_md5`, `overwrite`, `mkdirs`, `validate_remote_file` fields deprecated and they will be removed in the next version, where critical code path will be shared with `databricks_notebook`.
* `network_error_messages` and `verify_workspace_runnning` from `databricks_mws_workspaces` is deprecated and going to be removed in 0.3.
* `error_messages` from `databricks_mws_networks` are deprecated and would be removed in 0.3.
* `ebs_volume_type` and `azure_disk_volume_type` from `databricks_instance_pool` is going to be moved to `disk_type` sub-block in 0.3, which means you'll slightly have to modify configuration while migrating to 0.3. Computed field `default_tags` is going to be removed from resource. This is done to further increase maintainability of provider in the future.

Updated dependency versions:

* github.com/aws/aws-sdk-go 35.36
* github.com/hashicorp/go-retryablehttp 0.6.8
* github.com/Azure/go-autorest/autorest 0.11.12

**Behavior changes**
* `min_idle_instances` for `databricks_instance_pool` is now optional.
* `skip_validation` for `databricks_instance_profile` is going to be removed in 0.3.

## 0.2.8

* Added [databricks_workspace_conf](https://github.com/databrickslabs/terraform-provider-databricks/pull/398) resource.
* Added [databricks_mws_log_delivery](https://github.com/databrickslabs/terraform-provider-databricks/pull/343) resource for billable usage & audit logs consumption.
* Added [databricks_node_type](https://github.com/databrickslabs/terraform-provider-databricks/pull/376) data source for simpler selection of node types across AWS & Azure.
* Added [Azure Key Vault support](https://github.com/databrickslabs/terraform-provider-databricks/pull/381) for `databricks_secret_scope` for Azure CLI authenticated users.
* Added [is_pinned](https://github.com/databrickslabs/terraform-provider-databricks/pull/348) support for `databricks_cluster` resource.
* Fixed restarting cluster on changes in cluster configuration aren't related to the cluster configuration ([issue #379](https://github.com/databrickslabs/terraform-provider-databricks/issues/379))
* Fixed issue [#383](https://github.com/databrickslabs/terraform-provider-databricks/issues/383) by cleaning up clusters that fail to start.
* Fixed issue [#382](https://github.com/databrickslabs/terraform-provider-databricks/issues/382) by ignoring any incoming changes to deployment name of `databricks_mws_workspaces`, as well as propagating the right error messages.
* Internal: API for retrieval of the cluster events.
* Increased code coverage to 71%.

Updated dependency versions:

* github.com/Azure/go-autorest/autorest v0.11.9
* github.com/Azure/go-autorest/autorest/adal v0.9.5
* github.com/Azure/go-autorest/autorest/azure/auth v0.5.3
* github.com/Azure/go-autorest/autorest/azure/cli v0.4.2
* gopkg.in/ini.v1 1.62.0

**Deprecations**
* `network_error_messages` from `databricks_mws_workspaces` is deprecated and going to be removed in 0.3

## 0.2.7

* Small fixes

## 0.2.6

* Added support for [customer managed keys](https://github.com/databrickslabs/terraform-provider-databricks/pull/332) for Accounts API.
* Added `databricks_user` resource.
* Added `databricks_user_instance_profile` resource.
* Added `databricks_group` data source.

Updated dependency versions:

* github.com/Azure/go-autorest/autorest v0.11.6
* github.com/Azure/go-autorest/autorest/adal v0.9.4
* github.com/Azure/go-autorest/autorest/azure/auth v0.5.2
* github.com/Azure/go-autorest/autorest/azure/cli v0.4.1
* gopkg.in/ini.v1 v1.61.0

**Deprecations**
* `databricks_scim_user` is no longer receiving fixes and will be removed in `0.3`, please rewrite using the `databricks_user` resource, which has more consistent semantics with `databricks_group` and works better with identity provider SCIM sync.
* `databricks_scim_group` is no longer receiving fixes and will be removed in `0.3`. Please rewrite using the `databricks_group` resource.
* `databricks_default_user_roles` is no longer receiving fixes and will be removed in `0.3`, please rewrite using `databricks_user` & `databricks_group` resources.

**Behavior changes**
* State changes to legacy `spark.databricks.delta.preview.enabled` config option are [now ignored](https://github.com/databrickslabs/terraform-provider-databricks/pull/334) by `databricks_job` & `databricks_cluster`
* Libraries, which are installed on all clusters and are not part of cluster resource definition, won't be waited for INSTALLED status
* Fixed "[Secret scope ACL is MANAGE for all users by default](https://github.com/databrickslabs/terraform-provider-databricks/pull/326)" ([issue 322](https://github.com/databrickslabs/terraform-provider-databricks/issues/322)).  If you were relying on setting `MANAGE` permission to all users by default, you need to add `initial_manage_principal = "users"` to your `resource "databricks_secret_scope"` declaration. 

## 0.2.5

* Added support for [local disk encryption](https://github.com/databrickslabs/terraform-provider-databricks/pull/313)
* Added more reliable [indication about Azure environment](https://github.com/databrickslabs/terraform-provider-databricks/pull/312) and fixed azure auth issue for Terraform 0.13
* Updated [databricks_aws_crossaccount_policy](https://github.com/databrickslabs/terraform-provider-databricks/pull/311) to latest rules
* Fixed missing importers for [databricks_scim_*](https://github.com/databrickslabs/terraform-provider-databricks/pull/290) resources
* Updated [Terraform Plugin SDK](https://github.com/databrickslabs/terraform-provider-databricks/pull/279) to latest version along with transitive dependencies.
* Added support disclaimers
* Increased code coverage to 65%

## 0.2.4

* Added [Azure CLI authentication](https://github.com/databrickslabs/terraform-provider-databricks/blob/master/docs/index.md#authenticating-with-azure-cli) to bridge the gap of local development workflows and let more people use the provider.
* All authentication is completely [lazy-initialized](https://github.com/databrickslabs/terraform-provider-databricks/pull/270), which makes it provider overall more stable.
* Significantly increased [unit test coverage](https://codecov.io/gh/databrickslabs/terraform-provider-databricks), which runs before every merge of a pull request.
* Introduced constantly running integration test suite environments: azsp, azcli & awsmt
* Numerous stability improvements for clusters, mounts, libraries, notebooks, files, authentication and TLS connectivity.
* Added ability to mount storage without explicitly defining a cluster, though it will still launch auto-terminating `terraform-mount` cluster to perform the mount.
* `databricks_cluster` & `databricks_job` now share significant portion of configuration wiring code, therefore increasing the stability of codebase.
* Added support for Terraform 0.13 [local builds](https://github.com/databrickslabs/terraform-provider-databricks/pull/281) for those who develop or cannot wait for next release.
* Added AWS IAM Policy [data helpers](https://github.com/databrickslabs/terraform-provider-databricks/pull/255) to simplify new deployments.
* [Migrated all documentation](https://github.com/databrickslabs/terraform-provider-databricks/pull/250) to Terraform Registry format, therefore having a single always-accurate place for end-user guides.
* Internally, codebase [has been split](https://github.com/databrickslabs/terraform-provider-databricks/pull/224) into multiple packages, which should make further contributions simpler.

Updated dependency versions:

* github.com/Azure/go-autorest/autorest v0.11.4
* github.com/Azure/go-autorest/autorest/adal v0.9.2
* github.com/Azure/go-autorest/autorest/azure/auth v0.5.1
* github.com/aws/aws-sdk-go v1.34.13
* gopkg.in/ini.v1 v1.60.2

**Deprecations**
* `library_*` is no longer receiving fixes and will be removed in `0.3`, please rewrite cluster & job resources to use [`library` configuration block](https://github.com/databrickslabs/terraform-provider-databricks/blob/master/docs/resources/cluster.md#library-configuration-block).
* `basic_auth` provider block is no longer receiving fixesand will be removed in `0.3`, please use `username` and `password` options
* `azure_auth` provider block is no longer receiving fixesand will be removed in `0.3`, please use `azure_*` options 

**Behavior changes**
* Previously, mounts code paths were different functions. This release unifies them to be a single testable codebase with different configuration options & re-use of the critical code paths. For maintainability reasons, there's no longer check performed on container & storage account names, but rather on high-level *mount source uri*.
