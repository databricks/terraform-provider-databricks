# Version changelog

## 0.4.3

* Improved documentation with regards to public subnets in AWS quick start ([#1005](https://github.com/databrickslabs/terraform-provider-databricks/pull/1005)).
* Added `databricks_mount` code genration for [exporter](https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs/guides/experimental-exporter) tooling ([#1006](https://github.com/databrickslabs/terraform-provider-databricks/pull/1006)).
* Increase dependency check frequency ([#1007](https://github.com/databrickslabs/terraform-provider-databricks/pull/1007)).
* Added experimental resources.

## 0.4.2

* Added optional `auth_type` provider conf to enforce specific auth type to be used in very rare cases, where a single Terraform state manages Databricks workspaces on more than one cloud and `More than one authorization method configured` error is a false positive. Valid values are `pat`, `basic`, `azure-client-secret`, `azure-msi`, `azure-cli`, and `databricks-cli` ([#1000](https://github.com/databrickslabs/terraform-provider-databricks/pull/1000)).
* Added `DBC` format support for `databricks_notebook` ([#989](https://github.com/databrickslabs/terraform-provider-databricks/pull/989)). 
* Fixed creating new `databricks_mws_workspaces` with `token {}` block ([#994](https://github.com/databrickslabs/terraform-provider-databricks/issues/994)).
* Added automated documentation formatting with `make fmt-docs`, so that all HCL examples look consistent ([#999](https://github.com/databrickslabs/terraform-provider-databricks/pull/999)).
* Increased codebase unit test coverage to 91% to improve stability ([#996](https://github.com/databrickslabs/terraform-provider-databricks/pull/996), [#992](https://github.com/databrickslabs/terraform-provider-databricks/pull/992), [#991](https://github.com/databrickslabs/terraform-provider-databricks/pull/991), [#990](https://github.com/databrickslabs/terraform-provider-databricks/pull/990)).

Updated dependency versions:

* Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.10.0 to 2.10.1 

## 0.4.1

* Added `databricks_library` resource to install library on `databricks_cluster` ([#904](https://github.com/databrickslabs/terraform-provider-databricks/pull/904)).
* Added `databricks_clusters` data resource to list all clusters in the workspace, which might be used to install `databricks_library` on all clusters ([#955](https://github.com/databrickslabs/terraform-provider-databricks/pull/955)).
* Fixed refresh of `library` blocks on a stopped `databricks_cluster` ([#952](https://github.com/databrickslabs/terraform-provider-databricks/issues/952)).
* Whenever a library fails to get installed on a running `databricks_cluster`, we now automatically remove this library, so that the clean state of managed libraries is properly maintained. Without this fix users had to manually go to Clusters UI and remove library from a cluster, where it failed to install. Libraries add up to CREATE and UPDATE timeouts of `databricks_cluster` resource. ([#599](https://github.com/databrickslabs/terraform-provider-databricks/issues/599)).
* Added `token` block to `databricks_mws_workspaces` to avoid unnecessary provider aliasing ([#957](https://github.com/databrickslabs/terraform-provider-databricks/issues/957)).
* Fixed disabling `databricks_global_init_script` ([#958](https://github.com/databrickslabs/terraform-provider-databricks/issues/958)).
* Fixed configuration drift issues with `aws_attributes`, `azure_attributes`, `gcp_attributes`, and `email_notifications` configuration blocks in `databricks_cluster`, `databricks_job`, and `databricks_instance_pool` resources ([#981](https://github.com/databrickslabs/terraform-provider-databricks/pull/981)).
* Improved Databricks CLI auth by eagerly resolving `host`, `username`, `password`, and `token` from the specified `profile`. Added explicit logging of auth parameters in debug logs ([#965](https://github.com/databrickslabs/terraform-provider-databricks/pull/965)).
* TLS timeouts, which may occur during Azure MSI auth, are no longer failing API requests and retried within a normal policy ([#966](https://github.com/databrickslabs/terraform-provider-databricks/pull/966)).
* `debug_headers` provider conf is also logging the `Host` header to help troubleshooting auth issues ([#964](https://github.com/databrickslabs/terraform-provider-databricks/pull/964)).
* Added new experimental resources and increased test coverage.

Updated dependency versions:

* Bump github.com/golang-jwt/jwt/v4 from 4.1.0 to 4.2.0
* Bump google.golang.org/api from 0.60.0 to 0.63.0
* Bump github.com/Azure/go-autorest/autorest from 0.11.22 to 0.11.23
* Bump github.com/Azure/go-autorest/autorest/azure/auth from 0.5.9 to 0.5.10
* Bump gopkg.in/ini.v1 from 1.66.0 to 1.66.2
* Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.9.0 to 2.10.0

## 0.4.0

* Added `databricks_mlflow_model` and `databricks_mlflow_experiment` resources ([#931](https://github.com/databrickslabs/terraform-provider-databricks/pull/931)) 
* Added support for `repo_path` to `databricks_permissions` resource ([#875](https://github.com/databrickslabs/terraform-provider-databricks/issues/875)).
* Added `external_id` to `databricks_user` and `databricks_group` ([#927](https://github.com/databrickslabs/terraform-provider-databricks/pull/927)).
* Fixed `databricks_repo` creation corner cases on MS Windows OS ([#911](https://github.com/databrickslabs/terraform-provider-databricks/issues/911)).
* Fixed configuration drift for `databricks_cluster`.`aws_attributes`.`zone_id` with `auto`, which resulted in unwanted cluster restarts ([#937](https://github.com/databrickslabs/terraform-provider-databricks/pull/937)).
* Added new experimental resources, increased test coverage, and automated integration testing infrastructure.
* Multiple documentation improvements and new guides.

**Behavior changes**

* Renamed `allow_sql_analytics_access` to `databricks_sql_access` in `databricks_user`, `databricks_group`, and `databricks_service_principal` resources.
* Removed deprecated `azure_use_pat_for_spn`, `azure_use_pat_for_cli`, `azure_pat_token_duration_seconds` provider attributes.
* Removed deprecated `azure_workspace_name`, `azure_resource_group`, `azure_subscription_id` in favor of just using `azure_workspace_resource_id`.
* Renamed `DATABRICKS_AZURE_WORKSPACE_RESOURCE_ID` environment variable to `DATABRICKS_AZURE_RESOURCE_ID`.
* `DATABRICKS_AZURE_CLIENT_SECRET` environment variable is no longer having any effect in favor of just using `ARM_CLIENT_SECRET`.
* `DATABRICKS_AZURE_CLIENT_ID` environment variable is no longer having any effect in favor of just using `ARM_CLIENT_ID`.
* `DATABRICKS_AZURE_TENANT_ID` environment variable is no longer having any effect in favor of just using `ARM_TENANT_ID`.

Updated dependency versions:

* Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.7.1 to 2.9.0
* Bump github.com/Azure/go-autorest/autorest/adal from 0.9.16 to 0.9.17
* Bump github.com/golang-jwt/jwt/v4 from 4.0.0 to 4.1.0
* Bump github.com/zclconf/go-cty from 1.9.1 to 1.10.0
* Bump github.com/Azure/go-autorest/autorest from 0.11.21 to 0.11.22

## 0.3.11

* Added `databricks_sql_global_config` resource to provide global configuration for SQL Endpoints ([#855](https://github.com/databrickslabs/terraform-provider-databricks/issues/855))
* Added `databricks_mount` resource to mount arbitrary cloud storage ([#497](https://github.com/databrickslabs/terraform-provider-databricks/issues/497))
* Improved implementation of `databricks_repo` by creating the parent folder structure ([#895](https://github.com/databrickslabs/terraform-provider-databricks/pull/895))
* Fixed `databricks_job` error related [to randomized job IDs](https://docs.databricks.com/release-notes/product/2021/august.html#jobs-service-stability-and-scalability-improvements) ([#901](https://github.com/databrickslabs/terraform-provider-databricks/issues/901))
* Replace `databricks_group` on name change ([#890](https://github.com/databrickslabs/terraform-provider-databricks/pull/890))
* Names of scopes in `databricks_secret_scope` can have `/` characters in them ([#892](https://github.com/databrickslabs/terraform-provider-databricks/pull/892))

**Deprecations**
* `databricks_aws_s3_mount`, `databricks_azure_adls_gen1_mount`, `databricks_azure_adls_gen2_mount`, and `databricks_azure_blob_mount` are deprecated in favor of `databricks_mount`.

Updated dependency versions:

* Bump google.golang.org/api from 0.59.0 to 0.60.0

## 0.3.10

* Added `private_access_level` and `allowed_vpc_endpoint_ids` to `databricks_mws_private_access_settings` resource, which is also now updatable ([#867](https://github.com/databrickslabs/terraform-provider-databricks/issues/867)).
* Fixed missing diff skip for `skip_validation` in `databricks_instance_profile` ([#860](https://github.com/databrickslabs/terraform-provider-databricks/issues/860)).
* Added support for `pipeline_task` ([871](https://github.com/databrickslabs/terraform-provider-databricks/pull/871)) and `python_wheel_task` ([#872](https://github.com/databrickslabs/terraform-provider-databricks/pull/872)) to `databricks_job`.
* Improved enterprise HTTPS proxy support for creating workspaces in PrivateLink environments ([#882](https://github.com/databrickslabs/terraform-provider-databricks/pull/882)).
* Added `hostname` attribute to `odbc_params` in `databricks_sql_endpoint` ([#868](https://github.com/databrickslabs/terraform-provider-databricks/issues/868)).
* Improved documentation ([#858](https://github.com/databrickslabs/terraform-provider-databricks/pull/858), [#870](https://github.com/databrickslabs/terraform-provider-databricks/pull/870)).

Updated dependency versions:

* Bumped google.golang.org/api from 0.58.0 to 0.59.0

## 0.3.9

* Added initial support for multiple task orchestration in `databricks_job` [#853](https://github.com/databrickslabs/terraform-provider-databricks/pull/853)
* Fixed provider crash for new terraform states related to bug [#813](https://github.com/hashicorp/terraform-plugin-sdk/issues/813) in Terraform SDK v2.8.0 ([#854](https://github.com/databrickslabs/terraform-provider-databricks/issues/854))
* Re-added `skip_validation` to `databricks_instance_profile` resource [#762](https://github.com/databrickslabs/terraform-provider-databricks/issues/762)
* Removed direct dependency on `aws-sdk-go`.

Updated dependency versions:

* Reverted github.com/hashicorp/terraform-plugin-sdk/v2 from 2.8.0 to 2.7.0

## 0.3.8

* Added `databricks_repo` resource to manage [Databricks Repos](https://docs.databricks.com/repos.html) ([#771](https://github.com/databrickslabs/terraform-provider-databricks/pull/771))
* Added support for Azure MSI authentication ([#743](https://github.com/databrickslabs/terraform-provider-databricks/pull/743))
* Added support to create `databricks_user` on the account level ([#818](https://github.com/databrickslabs/terraform-provider-databricks/issues/818))
* Already deleted `databricks_token` don't fail the apply ([#808](https://github.com/databrickslabs/terraform-provider-databricks/pull/808))
* Default `terraform-mount` clusters created for mounting for `databricks_aws_s3_mount`, `databricks_azure_adls_gen1_mount`, `databricks_azure_adls_gen2_mount`, and `databricks_azure_blob_mount` have now `spark.scheduler.mode` as `FIFO` ([#828](https://github.com/databrickslabs/terraform-provider-databricks/pull/828))
* Fixed crash when using non-Azure authentication to mount Azure resources ([#831](https://github.com/databrickslabs/terraform-provider-databricks/issues/831))
* Fixed replacement of `instance_pool_id` in `databricks_cluster`, when `driver_instance_pool_id` was not explicitly specified ([#824](https://github.com/databrickslabs/terraform-provider-databricks/issues/824))
* Ingorning diff customization for permissions resource, so that new workspace deployments won't fail without explicit dependency on a workspace resource
* Multiple documentation improvements

**Deprecations**
* `azure_workspace_name`, `azure_resource_group`, `azure_subscription_id`, and `azure_workspace_resource_id` are deprecated and would be removed in v0.4.0. Please rewrite provider configuration with `host = data.azurerm_databricks_workspace.example.workspace_url` to achieve the same effect. Please check [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/data-sources/databricks_workspace#workspace_url) resource documentation for details.
* `azure_use_pat_for_spn`, `azure_use_pat_for_cli`, and `azure_pat_token_duration_seconds` are deprecated to fully switch to AAD token authentication in the near future.
* `DATABRICKS_AZURE_CLIENT_SECRET` environment variable is deprecated in favor of just using `ARM_CLIENT_SECRET`.
* `DATABRICKS_AZURE_CLIENT_ID` environment variable is deprecated in favor of just using `ARM_CLIENT_ID`.
* `DATABRICKS_AZURE_TENANT_ID` environment variable is deprecated in favor of just using `ARM_TENANT_ID`.

Updated dependency versions:

* Bump github.com/aws/aws-sdk-go from 1.40.19 to 1.40.54
* Bump github.com/Azure/go-autorest/autorest from 0.11.19 to 0.11.21
* Bump github.com/Azure/go-autorest/autorest/azure/cli from 0.4.2 to 0.4.3
* Bump github.com/Azure/go-autorest/autorest/adal from 0.9.14 to 0.9.16
* Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.7.0 to 2.8.0
* Bump github.com/zclconf/go-cty from 1.9.0 to 1.9.1
* Bump golang.org/x/mod from 0.4.2 to 0.5.1
* Bump google.golang.org/api from 0.52.0 to 0.58.0
* Bump gopkg.in/ini.v1 from 1.62.0 to 1.63.2

## 0.3.7

* Added `databricks_obo_token` resource to create On-Behalf-Of tokens for a Service Principal in Databricks workspaces on AWS. It is very useful, when you want to provision resources within a workspace through narrowly-scoped service principal, that has no access to other workspaces within the same Databricks Account ([#736](https://github.com/databrickslabs/terraform-provider-databricks/pull/736))
* Added support for [IAM credential passthrough](https://docs.databricks.com/security/credential-passthrough/iam-passthrough.html) with `is_meta_instance_profile` property for `databricks_instance_profile` ([#745](https://github.com/databrickslabs/terraform-provider-databricks/pull/745))
* Fixed incorrect workspace update bug and added more validation error messaging ([#649](https://github.com/databrickslabs/terraform-provider-databricks/pull/649))
* Clarify network modification procedure on active workspaces ([#732](https://github.com/databrickslabs/terraform-provider-databricks/issues/732))
* Updated AWS IAM policy templates version to `2012-10-17` (`databricks_aws_bucket_policy`, `databricks_aws_assume_role_policy`, and `databricks_aws_crossaccount_policy`) ([#688](https://github.com/databrickslabs/terraform-provider-databricks/issues/688))
* Various bug fixes in Databricks SQL resources

Updated dependency versions:

* Bump github.com/aws/aws-sdk-go to v1.40.12
* Bump github.com/hashicorp/hcl/v2 to v2.10.1
* Bump github.com/zclconf/go-cty to v1.9.0
* Bump golang.org/x/time to v0.0.0-20210723032227-1f47c861a9ac
* Bump golang.org/x/tools to v0.1.5

## 0.3.6

* Added support for hybrid pools ([#689](https://github.com/databrickslabs/terraform-provider-databricks/pull/689))
* Added support for `always_running` jobs, which are restarted on resource updates ([#715](https://github.com/databrickslabs/terraform-provider-databricks/pull/715))
* Azure CLI auth is now forcing JSON output ([#717](https://github.com/databrickslabs/terraform-provider-databricks/pull/717))
* `databricks_permissions` are getting validation on `terraform plan` stage ([#706](https://github.com/databrickslabs/terraform-provider-databricks/pull/706))
* Added `databricks_directory` resource ([#690](https://github.com/databrickslabs/terraform-provider-databricks/pull/690))
* Added `run_as_role` field to `databricks_sql_query` ([#684](https://github.com/databrickslabs/terraform-provider-databricks/pull/684))
* Added `user_id` attribute for `databricks_user` data resource, so that it's possible to dynamically create resources based on members of the group ([#714](https://github.com/databrickslabs/terraform-provider-databricks/pull/714))
* Added more selectors to `databricks_node_type` data source ([#723](https://github.com/databrickslabs/terraform-provider-databricks/pull/723))
* Azure auth with SPN now uses AAD token by default instead of PAT. Previous behavior (using PAT) could be restored by setting `azure_use_pat_for_spn` to `true` ([#721](https://github.com/databrickslabs/terraform-provider-databricks/pull/721))
* `deployment_name` for `databricks_mws_workspaces` is now optional, how it should have been. This enables creation of Databricks workspaces without an account prefix.
* To replicate default behavior of Databricks SQL UI, `enable_photon` is now `true` by default for `databricks_sql_endpoint`.
* Various documentation and bugfixes

Updated dependency versions:

* Bump github.com/aws/aws-sdk-go from 1.38.51 to 1.38.71
* Bump github.com/Azure/go-autorest/autorest/azure/auth from 0.5.7 to 0.5.8
* Bump github.com/Azure/go-autorest/autorest from 0.11.18 to 0.11.19
* Bump github.com/Azure/go-autorest/autorest/adal from 0.9.13 to 0.9.14
* Bump github.com/zclconf/go-cty from 1.8.3 to 1.8.4 
* Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.6.1 to 2.7.0

## 0.3.5

* Fixed setting of permissions for SQLA endpoints ([#661](https://github.com/databrickslabs/terraform-provider-databricks/issues/661))
* Added support for preloading of Docker images into instance pools ([#663](https://github.com/databrickslabs/terraform-provider-databricks/issues/663))
* Added the `databricks_user` data source ([#648](https://github.com/databrickslabs/terraform-provider-databricks/pull/648))
* Fixed support for `spot_instance_policy` in SQLA Endpoints ([#665](https://github.com/databrickslabs/terraform-provider-databricks/issues/665))
* Added documentation for `databricks_pipeline` resource ([#673](https://github.com/databrickslabs/terraform-provider-databricks/pull/673))
* Fixed mapping for `databricks_service_principal` on AWS ([#656](https://github.com/databrickslabs/terraform-provider-databricks/issues/656))
* Made preview environment tests to run on a release basis

Updated dependency versions:

* Bump github.com/zclconf/go-cty from 1.8.2 to 1.8.3
* Bump github.com/aws/aws-sdk-go from 1.38.30 to 1.38.51

## 0.3.4

* Fixed state refresh bugs in `databricks_sql_permissions` ([#620](https://github.com/databrickslabs/terraform-provider-databricks/issues/620), [#619](https://github.com/databrickslabs/terraform-provider-databricks/issues/620))
* Fixed `workspace_ids_filter` mapping for `databricks_mws_log_delivery` ([#635](https://github.com/databrickslabs/terraform-provider-databricks/issues/635))
* Multiple documentation improvements ([#597](https://github.com/databrickslabs/terraform-provider-databricks/issues/597), [eb60d10](https://github.com/databrickslabs/terraform-provider-databricks/commit/eb60d103ea63221a1eb0069723ba3a0af45dbe3b), [edcd4b1](https://github.com/databrickslabs/terraform-provider-databricks/commit/edcd4b121254e3ff3130bed9c4ef9d849d342561), [404bdab](https://github.com/databrickslabs/terraform-provider-databricks/commit/404bdab637c0a4a15b6a4b6a77567166315955ca), [#615](https://github.com/databrickslabs/terraform-provider-databricks/pull/615), [f14b825](https://github.com/databrickslabs/terraform-provider-databricks/commit/f14b825e9cb11d75e9ad077b35c7e9c410fd8351), [e615c3a](https://github.com/databrickslabs/terraform-provider-databricks/commit/e615c3a68d1ad45f91453ec448b55ca7b204fb97), [#612](https://github.com/databrickslabs/terraform-provider-databricks/pull/612))
* Mounting clusters are recreated now, even when they are deleted ([#637](https://github.com/databrickslabs/terraform-provider-databricks/issues/637))
* Fixed handling of empty blocks for clusters/jobs/instance pools ([22cdf2f](https://github.com/databrickslabs/terraform-provider-databricks/commit/22cdf2fc9d50f67b14b49d11e7fbaacce0f52399))
* Mark instance pool attributes as ForceNew when it's requited ([#629](https://github.com/databrickslabs/terraform-provider-databricks/issues/629))
* Switched to use https://staticcheck.io/ for static code analysis ([#602](https://github.com/databrickslabs/terraform-provider-databricks/issues/602))

**Behavior changes**

* The `customer_managed_key_id` field in `databricks_mws_workspaces` resource is deprecated and should be replaced with `managed_services_customer_managed_key_id` (and optionally `storage_customer_managed_key_id`). `databricks_mws_customer_managed_keys` now requires the parameter `use_cases` ([#642](https://github.com/databrickslabs/terraform-provider-databricks/pull/642)). *If you've used the resource before, please add `use_cases = ["MANAGED_SERVICES"]` to keep the behaviour.*

Updated dependency versions:

* Bump github.com/aws/aws-sdk-go to v1.38.30
* Bump github.com/hashicorp/go-retryablehttp to v0.7.0
* Bump github.com/hashicorp/hcl/v2 to v2.10.0
* Bump github.com/hashicorp/terraform-plugin-sdk/v2 to v2.6.1
* Bump github.com/zclconf/go-cty to v1.8.2

## 0.3.3

* Added resources for SQL Analytics queries and dashboards: `databricks_sql_query`, `databricks_sql_visualization`, `databricks_sql_dashboard`, `databricks_sql_widget` ([#553](https://github.com/databrickslabs/terraform-provider-databricks/pull/553))
* Added `databricks_sql_permissions` resource ([#545](https://github.com/databrickslabs/terraform-provider-databricks/pull/545/files))
* Fixed documentation bugs ([#603](https://github.com/databrickslabs/terraform-provider-databricks/issues/603))
* Improved resource exporter ([#593](https://github.com/databrickslabs/terraform-provider-databricks/issues/593))
* Added missing properties to `databricks_mws_private_access_settings` ([#590](https://github.com/databrickslabs/terraform-provider-databricks/issues/590))
* Include SQLA data source ID in `databricks_sql_endpoint` state ([#601](https://github.com/databrickslabs/terraform-provider-databricks/issues/601))
* Apply `debug_truncate_bytes` also for response dumps ([#589](https://github.com/databrickslabs/terraform-provider-databricks/issues/589))
* More verbose logging of `databricks_cluster` termination reason ([#588](https://github.com/databrickslabs/terraform-provider-databricks/issues/588))
* Move non-auth provider config documentation into separate section ([#587](https://github.com/databrickslabs/terraform-provider-databricks/pull/587))


## 0.3.2

* Fixed minor issues to add support for GCP ([#558](https://github.com/databrickslabs/terraform-provider-databricks/pull/558))
* Fixed `databricks_permissions` for SQL Analytics Entities ([#535](https://github.com/databrickslabs/terraform-provider-databricks/issues/535))
* Fixed incorrect HTTP 404 handling on create ([#564](https://github.com/databrickslabs/terraform-provider-databricks/issues/564), [#576](https://github.com/databrickslabs/terraform-provider-databricks/issues/576))
* Fixed incorrect escaping of notebook names ([#566](https://github.com/databrickslabs/terraform-provider-databricks/pull/566))
* Fixed entitlements for databricks_group ([#549](https://github.com/databrickslabs/terraform-provider-databricks/pull/549))
* Fixed rate limiting to perform more than 1 request per second ([#577](https://github.com/databrickslabs/terraform-provider-databricks/pull/577))
* Added support for spot instances on Azure ([#571](https://github.com/databrickslabs/terraform-provider-databricks/pull/571))
* Added job schedules support for `pause_status` as a optional field. ([#575](https://github.com/databrickslabs/terraform-provider-databricks/pull/575))
* Fixed minor documentation issues.

Updated dependency versions:

* Bump github.com/aws/aws-sdk-go from 1.37.20 to 1.38.10
* Bump github.com/hashicorp/hcl/v2 from 2.9.0 to 2.9.1 
* Bump github.com/zclconf/go-cty from 1.8.0 to 1.8.1
* Bump github.com/google/go-querystring from 1.0.0 to 1.1.0

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
