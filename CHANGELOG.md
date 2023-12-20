# Version changelog

## 1.33.0
New Features and Improvements:
 * Added support for ownership changes for unity catalog resources ([#3029](https://github.com/databricks/terraform-provider-databricks/pull/3029)).

Exporter:
 * Ignore workspace assets of deleted users and service principals ([#2980](https://github.com/databricks/terraform-provider-databricks/pull/2980)).

## 1.32.0
New Features and Improvements:
 * Improved debugging of errors ([#744](https://github.com/databricks/databricks-sdk-go/pull/744))
 * Removed requirement to provide accountId in mws_credentials resource ([#3028](https://github.com/databricks/terraform-provider-databricks/pull/3028)).
 * Validated metastore_id field in worspace UC resources except metastore resource ([#3035](https://github.com/databricks/terraform-provider-databricks/pull/3035)), ([#3039](https://github.com/databricks/terraform-provider-databricks/pull/3039)).
 * Wait for resource to be updated during grants resource modification ([#3026](https://github.com/databricks/terraform-provider-databricks/pull/3026)).
 * Deprecated top level task attributes ([#2993](https://github.com/databricks/terraform-provider-databricks/pull/2993)).

Documentation Changes:
 * Added BQ example for [databricks_connection](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/connection) ([#3040](https://github.com/databricks/terraform-provider-databricks/pull/3040)).
 * Consistent header levels and same doc structure ([#3007](https://github.com/databricks/terraform-provider-databricks/pull/3007)).
 * Fixed two minor docs typos in node_type and cluster_policy ([#3038](https://github.com/databricks/terraform-provider-databricks/pull/3038)).

Exporter:
 * Removed dead code that reflects old Jobs API 2.0 structures ([#2981](https://github.com/databricks/terraform-provider-databricks/pull/2981)).

 Internal Changes:
 * Bumped Go SDK to 0.28.1 ([#3043](https://github.com/databricks/terraform-provider-databricks/pull/3043)).
 * Bumped github.com/hashicorp/terraform-plugin-sdk/v2 from 2.30.0 to 2.31.0 ([#3033](https://github.com/databricks/terraform-provider-databricks/pull/3033)).
 * Added acceptance tests for [databricks_grants](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/grants) ([#3034](https://github.com/databricks/terraform-provider-databricks/pull/3034)).
 * Migrated library create to use go-sdk ([#2998](https://github.com/databricks/terraform-provider-databricks/pull/2998)).

## 1.31.1

New Features and Improvements:
 * Added `workspace_path` attribute to [databricks_workspace_file](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/workspace_file) resource ([#2984](https://github.com/databricks/terraform-provider-databricks/pull/2984)).

Bugfixes:
 * Bugfix for azure authentication ([#732](https://github.com/databricks/databricks-sdk-go/pull/732))

Exporter:
 * Exporter: fix suppression of empty `webhook_notifications` blocks in tasks ([#2954](https://github.com/databricks/terraform-provider-databricks/pull/2954)).

Internal Changes:
 * Migrated library read to use Go SDK ([#2985](https://github.com/databricks/terraform-provider-databricks/pull/2985)).
 * Migrated [databricks_mws_credentials](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/mws_credentials) to Go SDK ([#2962](https://github.com/databricks/terraform-provider-databricks/pull/2962)).
 * Update Go SDK to v0.26.2


## 1.31.0

New Features and Improvements:
 * Add `webhooks_notifications` on task levels in [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) resource ([#2953](https://github.com/databricks/terraform-provider-databricks/pull/2953)).
 * Add delegation support for [databricks_storage_credential](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/storage_credential) ([#2973](https://github.com/databricks/terraform-provider-databricks/pull/2973)).
 * Add edit_mode to [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) ([#2918](https://github.com/databricks/terraform-provider-databricks/pull/2918)).
 * Added [databricks_instance_profiles](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/instance_profiles) data source ([#2891](https://github.com/databricks/terraform-provider-databricks/pull/2891)).
 * Allow to override built-in [databricks_cluster_policy](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/cluster_policy) resources ([#2941](https://github.com/databricks/terraform-provider-databricks/pull/2941)).
 * Add [databricks_current_config](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/current_config) data source ([#2890](https://github.com/databricks/terraform-provider-databricks/pull/2890)).

Bugfixes:
 * Do not persist state from failed updates in  databricks_workspace_conf  ([#2914](https://github.com/databricks/terraform-provider-databricks/pull/2914)).
 * Fix `force_destroy` behaviour for [databricks_schema](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/schema) ([#2907](https://github.com/databricks/terraform-provider-databricks/pull/2907)).
 * Fix databricks_workspace_conf nightly race ([#2952](https://github.com/databricks/terraform-provider-databricks/pull/2952)).
 * Force sending the `content` field when creating empty [databricks_workspace_file](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/workspace_file) resources ([#2931](https://github.com/databricks/terraform-provider-databricks/pull/2931)).
 * Generate correct `url` field for [databricks_group](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/group) resource on account level ([#2911](https://github.com/databricks/terraform-provider-databricks/pull/2911)).
 * Make default_namespace_setting resource name singular ([#2925](https://github.com/databricks/terraform-provider-databricks/pull/2925)).
 * Fix: resolved issue with run_as_role not being applied on databricks_sql_query resource ([#2940](https://github.com/databricks/terraform-provider-databricks/pull/2940)).
 * Remove diff suppression for a job's `run_as` block ([#2943](https://github.com/databricks/terraform-provider-databricks/pull/2943)).

Exporter:
 * Fix generation of `branch` & `tag` attributes in [databricks_repo](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/repo) ([#2929](https://github.com/databricks/terraform-provider-databricks/pull/2929)).
 * Generate correct code for inherited [databricks_cluster_policies](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/cluster_policies) ([#2936](https://github.com/databricks/terraform-provider-databricks/pull/2936)).
 * Generate references also for integer attributes ([#2946](https://github.com/databricks/terraform-provider-databricks/pull/2946)).

Documentation changes:
 * Add documentation for `condition_task` in [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) resource ([#2955](https://github.com/databricks/terraform-provider-databricks/pull/2955)).
 * Clarify that any_file also works for reading ([#2937](https://github.com/databricks/terraform-provider-databricks/pull/2937)).
 * Documented use of libraries in the [databricks_cluster_policy](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/cluster_policy) resource ([#2934](https://github.com/databricks/terraform-provider-databricks/pull/2934)).
 * Small databricks_catalog_workspace_binding docs change ([#2923](https://github.com/databricks/terraform-provider-databricks/pull/2923)).
 * [DOC] Refresh Unity Catalog guides with the latest updates ([#2917](https://github.com/databricks/terraform-provider-databricks/pull/2917)).
 * [JOBS-13079] Add description to job settings ([#2932](https://github.com/databricks/terraform-provider-databricks/pull/2932)).
 * doc fixes ([#2987](https://github.com/databricks/terraform-provider-databricks/pull/2987)).
 * removing outdated note from doc ([#2969](https://github.com/databricks/terraform-provider-databricks/pull/2969)).
  
Internal Changes:
 * Add integration test coverage for delta sharing parameters ([#2991](https://github.com/databricks/terraform-provider-databricks/pull/2991)).
 * Bump Go SDK to v0.26.1 ([#2968](https://github.com/databricks/terraform-provider-databricks/pull/2968)).
 * Fixed unit tests for the upcoming Go SDK release ([#2956](https://github.com/databricks/terraform-provider-databricks/pull/2956)).
 * Migrate delete to use go-sdk for library resource ([#2960](https://github.com/databricks/terraform-provider-databricks/pull/2960)).
 * Migrate jobs delete to go-sdk ([#2948](https://github.com/databricks/terraform-provider-databricks/pull/2948)).
 * Moved `contains` function to common package ([#2990](https://github.com/databricks/terraform-provider-databricks/pull/2990)).

## 1.30.0

New Features and Improvements:
 * Added DeltaSharingRecipientTokenLifetimeInSeconds to ForceSendFields if DeltaSharingScope is set in [databricks_metastore](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/metastore) ([#2846](https://github.com/databricks/terraform-provider-databricks/pull/2846)).
 * Added NumWorkers to ForceSendFields in [databricks_pipeline](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/pipeline) if `spark_conf` contains config for single-node cluster ([#2852](https://github.com/databricks/terraform-provider-databricks/pull/2852)).
 * Added Terraform Provider for DefaultNamespaceSettings ([#2710](https://github.com/databricks/terraform-provider-databricks/pull/2710)).
 * Added `binding_type` attribute to [databricks_catalog_workspace_binding](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/catalog_workspace_binding) resource ([#2803](https://github.com/databricks/terraform-provider-databricks/pull/2803)).
 * Added `run_job_task` into the list of supported tasks of [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) resource ([#2889](https://github.com/databricks/terraform-provider-databricks/pull/2889)).
 * Added support for `run_as_role` in [databricks_sql_dashboard](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_dashboard) resource ([#2756](https://github.com/databricks/terraform-provider-databricks/pull/2756)).
resource ([#2495](https://github.com/databricks/terraform-provider-databricks/pull/2495)).
 * Added [databricks_mlflow_model](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/mlflow_model) data source ([#2456](https://github.com/databricks/terraform-provider-databricks/pull/2456)).
 * Allow authentication during configure to fail ([#2892](https://github.com/databricks/terraform-provider-databricks/pull/2892)).
 * Authenticate client during configuration ([#2885](https://github.com/databricks/terraform-provider-databricks/pull/2885)).
 * Improve performance of [databricks_notebook](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/notebook) and [databricks_workspace_file](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/workspace_file) resources ([#2902](https://github.com/databricks/terraform-provider-databricks/pull/2902)).
 * Make `aws_key_info.key_alias` optional for [databricks_mws_customer_managed_keys](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/mws_customer_managed_keys) ([#2791](https://github.com/databricks/terraform-provider-databricks/pull/2791)).
 * Make `prefix` & `suffix` optional in `multiple` block of [databricks_sql_query](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_query) ([#2905](https://github.com/databricks/terraform-provider-databricks/pull/2905)).
 * Recreate metastore assignment when changing workspace id ([#2899](https://github.com/databricks/terraform-provider-databricks/pull/2899)).
 * Update connection types ([#2897](https://github.com/databricks/terraform-provider-databricks/pull/2897)).
 * Mark `storage_root` optional for [databricks_metastore](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/metastore) resource ([#2873](https://github.com/databricks/terraform-provider-databricks/pull/2873)).
 * Refactor [databricks_recipient](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/recipient) to Go SDK ([#2858](https://github.com/databricks/terraform-provider-databricks/pull/2858)).

Bugfixes:
 * Fixed `Update` method of [databricks_user](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/user) ([#2878](https://github.com/databricks/terraform-provider-databricks/pull/2878)).
 * Fixed `Read` method of the [databricks_mlflow_model](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/mlflow_model) resource ([#2853](https://github.com/databricks/terraform-provider-databricks/pull/2853)).

Exporter:
 * Fixed emitting of group membership 
([#2868](https://github.com/databricks/terraform-provider-databricks/pull/2868)).
 * Don't omit `display_name` in [databricks_user](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/user) if it's not equal to `user_name` ([#2851](https://github.com/databricks/terraform-provider-databricks/pull/2851)).
 * Follow recommendations of identity team on decreasing load on SCIM API ([#2773](https://github.com/databricks/terraform-provider-databricks/pull/2773)).
 * Add case-insensitive match type for dependencies ([#2854](https://github.com/databricks/terraform-provider-databricks/pull/2854)).
 * Disable generation of empty or default blocks for [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) and [databricks_cluster](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/cluster) ([#2857](https://github.com/databricks/terraform-provider-databricks/pull/2857)).

Documentation Changes:
 * Update [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) documentation with `queue` block ([#2872](https://github.com/databricks/terraform-provider-databricks/pull/2872)).
 * Use correct label in the issue template for documentation ([#2850](https://github.com/databricks/terraform-provider-databricks/pull/2850)).
 * Modified github issue template to identify regressions faster ([#2847](https://github.com/databricks/terraform-provider-databricks/pull/2847)).
 * Fixed documented default value for data_security_mode ([#2866](https://github.com/databricks/terraform-provider-databricks/pull/2866)).
 * Add a note about setting entitlements with [databricks_entitlements](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/entitlements) for account-level identities ([#2910](https://github.com/databricks/terraform-provider-databricks/pull/2910)).
 * Add the missing `region` field to [databricks_metastore](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/metastore) resource block ([#2859](https://github.com/databricks/terraform-provider-databricks/pull/2859)).

Internal Changes:
 * Skiped GCP for UC integration tests using system schemas or statement execution API ([#2877](https://github.com/databricks/terraform-provider-databricks/pull/2877)).
 * Added integration test for SQL warehouse permissions in [databricks_permissions](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/permissions) ([#2845](https://github.com/databricks/terraform-provider-databricks/pull/2845)).
 * Bumped github.com/databricks/databricks-sdk-go from 0.24.0 to 0.25.0 ([#2909](https://github.com/databricks/terraform-provider-databricks/pull/2909)).
 * Enabled and fix tests ([#2901](https://github.com/databricks/terraform-provider-databricks/pull/2901)).
 * Fix Acceptance Tests ([#2916](https://github.com/databricks/terraform-provider-databricks/pull/2916)).

## 1.29.0

New Features and Improvements:
 * Added [databricks_artifact_allowlist](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/artifact_allowlist) ([#2744](https://github.com/databricks/terraform-provider-databricks/pull/2744)).
 * Added deployment field to databricks_job ([#2821](https://github.com/databricks/terraform-provider-databricks/pull/2821)).
 * Added support for `warehouse_id`, `options`, `partitions` and liquid clustering for [databricks_sql_table](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_table) resource ([#2789](https://github.com/databricks/terraform-provider-databricks/pull/2789)).

Bugfixes:
 * Fixed DiffSuppressFunc for [databricks_model_serving](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/model_serving) resource ([#2813](https://github.com/databricks/terraform-provider-databricks/pull/2813)).
 * Fixed [databricks_catalog](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/catalog) `isolation_mode` ([#2805](https://github.com/databricks/terraform-provider-databricks/pull/2805)).

Exporter:
 * Exporter: Correct treatment of disallowed characters in notebooks/directory names ([#2831](https://github.com/databricks/terraform-provider-databricks/pull/2831)).
 * Exporter: Added support for exporting [databricks_access_control_rule_set](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/access_control_rule_set) on account level ([#2699](https://github.com/databricks/terraform-provider-databricks/pull/2699)).
 * Exporter: Adjusted generation & normalization of job names ([2836](https://github.com/databricks/terraform-provider-databricks/pull/2836)).
 * Exporter: Clarified behaviour of services not marked with listing ([#2785](https://github.com/databricks/terraform-provider-databricks/pull/2785)).
 * Exporter: Fixed generation of names for [databricks_library](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/library) resources ([#2832](https://github.com/databricks/terraform-provider-databricks/pull/2832)).
 * Clarified experimental status of the exporter ([#2816](https://github.com/databricks/terraform-provider-databricks/pull/2816)).
 * Exporter: Ensured that name of databricks_repo is unique ([#2842](https://github.com/databricks/terraform-provider-databricks/pull/2842)).

Documentation changes:
 * Added `account_id` parameter to provider blocks in the guides ([#2817](https://github.com/databricks/terraform-provider-databricks/pull/2817)).
 * Changed CREATE_TABLE to CREATE_EXTERNAL_TABLE ([#2812](https://github.com/databricks/terraform-provider-databricks/pull/2812)).
 * Corrected max number of pinned clusters in the [databricks_cluster](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/cluster) docs ([#2825](https://github.com/databricks/terraform-provider-databricks/pull/2825)).
 * Put `empty_result_state` under `options` in the [databricks_sql_alert](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_alert) resource ([#2834](https://github.com/databricks/terraform-provider-databricks/pull/2834)).
 * Corrected grammar & rewording for better readability (part 1) ([#2781](https://github.com/databricks/terraform-provider-databricks/pull/2781)).

Internal changes:
 * Added additional integration tests for condition_task tasks in databricks_job ([#2830](https://github.com/databricks/terraform-provider-databricks/pull/2830)).
 * Bump github.com/databricks/databricks-sdk-go from 0.23.0 to 0.24.0 ([#2835](https://github.com/databricks/terraform-provider-databricks/pull/2835)).
 * Bump github.com/hashicorp/hcl/v2 from 2.18.1 to 2.19.1 ([#2819](https://github.com/databricks/terraform-provider-databricks/pull/2819)).
 * Enabled Cluster Integration test ([#2815](https://github.com/databricks/terraform-provider-databricks/pull/2815)).
 * Changed Jobs List operations to use token-based pagination ([#2810](https://github.com/databricks/terraform-provider-databricks/pull/2810)).

## 1.28.1
 * Fixed read method for `databricks_storage_credential` resource ([#2804](https://github.com/databricks/terraform-provider-databricks/pull/2804)).
 

## 1.28.0
* Added `dashboard_filters_enabled` attribute to [databricks_sql_dashboard](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_dashboard) resource ([#2725](https://github.com/databricks/terraform-provider-databricks/pull/2725)).
 * Added `empty_result_state` attribute to the [databricks_sql_alert](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_alert) resource ([#2724](https://github.com/databricks/terraform-provider-databricks/pull/2724)).
 * Added enabled field for queueing ([#2741](https://github.com/databricks/terraform-provider-databricks/pull/2741)).
 * Added [databricks_registered_model](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/registered_model) resource ([#2771](https://github.com/databricks/terraform-provider-databricks/pull/2771)).
 * Added logging package and fixed issue with API calls not being shown in DEBUG or lower log levels ([#2747](https://github.com/databricks/terraform-provider-databricks/pull/2747)).
 * Added [databricks_system_schema](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/system_schema) resource ([#2606](https://github.com/databricks/terraform-provider-databricks/pull/2606)).
 * Don't rely on having `@` to check if it's user or SP ([#2765](https://github.com/databricks/terraform-provider-databricks/pull/2765)).
 * Forced recreation of UC Volume when `volume_type` and `storage_location` are changed ([#2734](https://github.com/databricks/terraform-provider-databricks/pull/2734)).
 * Improved Provider Logging ([#2801](https://github.com/databricks/terraform-provider-databricks/pull/2801)).
 * Marked attributes in the `run_as` block in [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) as `ExactlyOneOf` ([#2784](https://github.com/databricks/terraform-provider-databricks/pull/2784)).
 * Masked sensitive field ([#2755](https://github.com/databricks/terraform-provider-databricks/pull/2755)).
 * Removed deprecation warning from `cluster_mount_info` in [databricks_cluster](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/cluster), but mark it as experimental ([#2787](https://github.com/databricks/terraform-provider-databricks/pull/2787)).
 * Suppress diff for `user_name` in [databricks_user](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/user) when the changes only in character case ([#2786](https://github.com/databricks/terraform-provider-databricks/pull/2786)).
 * Refresh grant lists ([#2746](https://github.com/databricks/terraform-provider-databricks/pull/2746)).
 * Fixed run_as_role drift for databricks_sql_query resource ([#2799](https://github.com/databricks/terraform-provider-databricks/pull/2799)).
 * Fixed metastore read and add test ([#2795](https://github.com/databricks/terraform-provider-databricks/pull/2795)).

Exporter:
 * Exporter: fix a logic for omitting some fields ([#2774](https://github.com/databricks/terraform-provider-databricks/pull/2774)).
 * Exporter: improve exporting of [databricks_cluster_policy](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/cluster_policy) resource ([#2680](https://github.com/databricks/terraform-provider-databricks/pull/2680)).
 * Exporter: parallel export of resources ([#2742](https://github.com/databricks/terraform-provider-databricks/pull/2742)).

Documentation:
 * Updated [databricks_grants](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/grants) examples for [databricks_external_location](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/external_location) ([#2735](https://github.com/databricks/terraform-provider-databricks/pull/2735)).
 * Fixed documentation for [databricks_schema](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/schema) about default value for `storage_root` ([#2790](https://github.com/databricks/terraform-provider-databricks/pull/2790)).
 * Clarified possible values for `principal` attribute of [databricks_secret_acl](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/secret_acl) ([#2772](https://github.com/databricks/terraform-provider-databricks/pull/2772)).

Other Changes:
 * Bumped databricks-sdk-go dependency to 0.21.0 ([#2738](https://github.com/databricks/terraform-provider-databricks/pull/2738)).
 * Bumped github.com/databricks/databricks-sdk-go from 0.21.0 to 0.22.0 ([#2761](https://github.com/databricks/terraform-provider-databricks/pull/2761)).
 * Bumped github.com/databricks/databricks-sdk-go from 0.22.0 to 0.23.0 ([#2794](https://github.com/databricks/terraform-provider-databricks/pull/2794)).
 * Bumped github.com/hashicorp/hcl/v2 from 2.18.0 to 2.18.1 ([#2776](https://github.com/databricks/terraform-provider-databricks/pull/2776)).
 * Bumped github.com/zclconf/go-cty from 1.14.0 to 1.14.1 ([#2777](https://github.com/databricks/terraform-provider-databricks/pull/2777)).
 * Used `terraform-field-dev` as code owner instead of `field-dev-ecosystem` ([#2718](https://github.com/databricks/terraform-provider-databricks/pull/2718)).
 * GitHub Actions workflow to compute provider schema diff ([#2740](https://github.com/databricks/terraform-provider-databricks/pull/2740)).


## 1.27.0
 * Fixed [databricks_permissions](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/permissions) resource for correct permissions update. ([#2719](https://github.com/databricks/terraform-provider-databricks/pull/2719)).
 * Added `owner` & `force_destroy` to [databricks_metastore_data_access](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/metastore_data_access) ([#2713](https://github.com/databricks/terraform-provider-databricks/pull/2713)).
 * Fixed [databricks_metastore_data_access](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/metastore_data_access) and [databricks_storage_credential](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/storage_credential) creation on GCP ([#2712](https://github.com/databricks/terraform-provider-databricks/pull/2712)).
 * Fixed missing `registered_model_id` attribute in [databricks_mlflow_model](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/mlflow_model) resource ([#2732](https://github.com/databricks/terraform-provider-databricks/pull/2732)).
 * Fixed mount read failing if cluster no longer exists ([#2634](https://github.com/databricks/terraform-provider-databricks/pull/2634)).
 * Deprecated `cluster_mount_info` block in [databricks_cluster](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/cluster) resource ([#2703](https://github.com/databricks/terraform-provider-databricks/pull/2703)).

 Exporter:
 * Exporter: parallel listing of workspace objects ([#2691](https://github.com/databricks/terraform-provider-databricks/pull/2691)).
 * Exporter: removed `modifiedAt != 0` check in the incremental mode ([#2736](https://github.com/databricks/terraform-provider-databricks/pull/2736)).

 Other Changes:
 * Bumped github.com/databricks/databricks-sdk-go from 0.19.2 to 0.20.0 ([#2714](https://github.com/databricks/terraform-provider-databricks/pull/2714)).



## 1.26.0
 * Removed `computed` for `pause_status` in [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) resource ([#2696](https://github.com/databricks/terraform-provider-databricks/pull/2696)).
 * Added support for init scripts from Unity Catalog Volumes in [databricks_cluster](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/cluster) ([#2666](https://github.com/databricks/terraform-provider-databricks/pull/2666)).
 * Added deprecation warning for init scripts from DBFS in [databricks_cluster](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/cluster) ([#2667](https://github.com/databricks/terraform-provider-databricks/pull/2667)).
 * Deprecated `photon` and `graviton` selectors in [databricks_spark_version](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/spark_version) data source ([#2687](https://github.com/databricks/terraform-provider-databricks/pull/2687)).

Documentation:
 * Documentation updates for UC resources ([#2632](https://github.com/databricks/terraform-provider-databricks/pull/2632)).
 * Fixed incorrect reference in the [databricks_git_credential](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/git_credential) documentation ([#2674](https://github.com/databricks/terraform-provider-databricks/pull/2674)).
 * Fixed Service Principal reference ([#2694](https://github.com/databricks/terraform-provider-databricks/pull/2694)).

Exporter:
 * Exporter: export references in `run_as` block of [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) ([#2690](https://github.com/databricks/terraform-provider-databricks/pull/2690)).

## 1.25.1

 * Fixed the issue with `cannot reset nil reader` by bumping Go SDK to 0.19.2 ([#2684](https://github.com/databricks/terraform-provider-databricks/pull/2684)).
 * Changed validation for `max_concurrent_runs` in `databricks_job` to allow 0 value ([#2682](https://github.com/databricks/terraform-provider-databricks/pull/2682)).


## 1.25.0

 * Added `IS_OWNER` permission for SQL Warehouse ([#2600](https://github.com/databricks/terraform-provider-databricks/pull/2600)).
 * Added `managed_identity_id` to [databricks_storage_credential](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/storage_credential) to support user-assigned managed identities ([#2536](https://github.com/databricks/terraform-provider-databricks/pull/2536)).
 * Added `options` field to UC catalog resource to support foreign catalog creation ([#2616](https://github.com/databricks/terraform-provider-databricks/pull/2616)).
 * Detected run_as drift in job resource ([#2626](https://github.com/databricks/terraform-provider-databricks/pull/2626)).
 * Removed callback field from acceptance test framework ([#2649](https://github.com/databricks/terraform-provider-databricks/pull/2649)).
 * Fixed `force_new` attributes for [databricks_schema](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/schema) and [databricks_volume](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/volume) ([#2635](https://github.com/databricks/terraform-provider-databricks/pull/2635)).
 * Fixed do not attempt to delete default schema for foriegn catalogs ([#2622](https://github.com/databricks/terraform-provider-databricks/pull/2622)).

 Documentation:
 * Updated documentation for marketplace admin role ([#2638](https://github.com/databricks/terraform-provider-databricks/pull/2638)).
 * Updated documentation for share and share recipient ([#2641](https://github.com/databricks/terraform-provider-databricks/pull/2641)).
 * Updated documentation for group, service principal and user([#2644](https://github.com/databricks/terraform-provider-databricks/pull/2644)).

 Other Changes:
 * Bumped github.com/databricks/databricks-sdk-go from 0.17.0 to 0.19.1 ([#2660](https://github.com/databricks/terraform-provider-databricks/pull/2660)).
 * Bumped github.com/hashicorp/hcl/v2 from 2.17.0 to 2.18.0 ([#2636](https://github.com/databricks/terraform-provider-databricks/pull/2636)).
 * Added doc strings for ResourceFixtures ([#2633](https://github.com/databricks/terraform-provider-databricks/pull/2633)).


## 1.24.1

 * Fixed verification of workspace reachability by using scim/me which is always available  ([#2618](https://github.com/databricks/terraform-provider-databricks/pull/2618)).

## 1.24.0

 * Added account-level API support for Unity Catalog objects ([#2182](https://github.com/databricks/terraform-provider-databricks/pull/2182)).
 * Added [databricks_connection](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/connection) resource to support Lakehouse Federation ([#2528](https://github.com/databricks/terraform-provider-databricks/pull/2528)).
 * Added `owner` parameter to [databricks_share](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/share) resource ([#2594](https://github.com/databricks/terraform-provider-databricks/pull/2594)).
 * Added `acl_principal_id` to data sources: [databricks_user](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/user), [databricks_service_principal](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/service_principal), [databricks_group](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/group), [databricks_current_user](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/current_user) ([#2555](https://github.com/databricks/terraform-provider-databricks/pull/2555)).
 * Fixed creation of views with comments using [databricks_sql_table](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_table) ([#2589](https://github.com/databricks/terraform-provider-databricks/pull/2589)) for GCP.
 * Fixed reflection method marshallJSON for CMK in mws workspace ([#2605](https://github.com/databricks/terraform-provider-databricks/pull/2605)).
 * Fixed databricks_access_control_rule_set integration test in Azure ([#2591](https://github.com/databricks/terraform-provider-databricks/pull/2591)).
 * Fixed RunJobTask job_id type ([#2588](https://github.com/databricks/terraform-provider-databricks/pull/2588)).
 * Updated Go SDK to v0.17.0 ([#2599](https://github.com/databricks/terraform-provider-databricks/pull/2599)).

 Documentation:
 * Added troubleshooting guide for grants/permissions config drifts ([#2576](https://github.com/databricks/terraform-provider-databricks/pull/2576)).
 * Updated doc for [databricks_model_serving](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/model_serving) to include new fields ([#2579](https://github.com/databricks/terraform-provider-databricks/pull/2579)).
 * Added instructions for group rule set management ([#2561](https://github.com/databricks/terraform-provider-databricks/pull/2561)).
 * Added missing documentation for CMK support on GCP ([#2604](https://github.com/databricks/terraform-provider-databricks/pull/2604)).

 Exporter:
 * Exporter: Incremental export of notebooks, SQL objects and some other resources ([#2563](https://github.com/databricks/terraform-provider-databricks/pull/2563)).
 * Exporter: add List operation for [databricks_instance_pool](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/instance_pool) resource ([#2570](https://github.com/databricks/terraform-provider-databricks/pull/2570)).
 * Exporter: command-line option to control output format for notebooks ([#2569](https://github.com/databricks/terraform-provider-databricks/pull/2569)).
 * Exporter: Added exporting of [databricks_mlflow_webhook](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/mlflow_webhook) resources ([#2552](https://github.com/databricks/terraform-provider-databricks/pull/2552)).

 Other Changes:
 * Migrated [databricks_mlflow_model](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/mlflow_model) to go sdk ([#2257](https://github.com/databricks/terraform-provider-databricks/pull/2257)).
 * Migrated [databricks_mlflow_webhook](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/mlflow_webhook) to Go SDK ([#2560](https://github.com/databricks/terraform-provider-databricks/pull/2560)).
 * Refactored [databricks_external_location](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/external_location) to Go SDK ([#2546](https://github.com/databricks/terraform-provider-databricks/pull/2546)).
 * Refactored [databricks_schema](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/schema) to Go SDK ([#2572](https://github.com/databricks/terraform-provider-databricks/pull/2572)).
 * Refreshed [databricks_grants](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/grants) with latest permissible grants ([#2567](https://github.com/databricks/terraform-provider-databricks/pull/2567)).
 * Fixed UC acceptance test ([#2613](https://github.com/databricks/terraform-provider-databricks/pull/2613)).
 * Auto-assign engineering reviewers for TF ([#2564](https://github.com/databricks/terraform-provider-databricks/pull/2564)).


## 1.23.0

 * Added Terraform support for Job Parameters (Private Preview) ([#2509](https://github.com/databricks/terraform-provider-databricks/pull/2509)).
 * Added `gcp_attributes.local_ssd_count` to [databricks_cluster](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/cluster) resource ([#2422](https://github.com/databricks/terraform-provider-databricks/pull/2422)).
 * Added `gcp_attributes.local_ssd_count` to [databricks_instance_pool](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/instance_pool) resource ([#2558](https://github.com/databricks/terraform-provider-databricks/pull/2558)).
 * Extend RunJobTask with additional supported fields and document them in [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) ([#2314](https://github.com/databricks/terraform-provider-databricks/pull/2314), [#2562](https://github.com/databricks/terraform-provider-databricks/pull/2562)).
 * Fixed update from instance pool to node type in [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) ([#2549](https://github.com/databricks/terraform-provider-databricks/pull/2549)).

Exporter:
 * Exporter: add support for [databricks_model_serving](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/model_serving) ([#2512](https://github.com/databricks/terraform-provider-databricks/pull/2512)).

Documentation:
 * Documented missing environment variables for authentication ([#2541](https://github.com/databricks/terraform-provider-databricks/pull/2541)).
 * Documented run_if field and suppress diff when field is not present ([#2435](https://github.com/databricks/terraform-provider-databricks/pull/2435)).
 * Updated diagram with Databricks resources ([#2526](https://github.com/databricks/terraform-provider-databricks/pull/2526)).
 * Updated metastore.md ([#2547](https://github.com/databricks/terraform-provider-databricks/pull/2547)).

Other Changes:
 * Marked `gcp_attributes.use_preemptible_executors` as deprecated in [databricks_cluster](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/cluster) ([#2551](https://github.com/databricks/terraform-provider-databricks/pull/2551)).
 * Provided `application_id` when creating SP in [databricks_access_control_rule_set](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/access_control_rule_set) resource integration test (and temporarily disabled this test outside of AWS) ([#2542](https://github.com/databricks/terraform-provider-databricks/pull/2542)).

## 1.22.0

 * Added [databricks_access_control_rule_set](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/access_control_rule_set) resource for managing account-level access ([#2371](https://github.com/databricks/terraform-provider-databricks/pull/2371)).
 * Added READ_VOLUME and WRITE_VOLUME to [databricks_grants](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/grants) resources at the schema/catalog-level ([#2529](https://github.com/databricks/terraform-provider-databricks/pull/2529)).
 * Added `acl_principal_id` attribute to [databricks_user](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/user), [databricks_group](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/group) & [databricks_service_principal](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/service_principal) for easier use with [databricks_access_control_rule_set](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/access_control_rule_set) ([#2485](https://github.com/databricks/terraform-provider-databricks/pull/2485)).
 * Added `control_run_state` flag to the [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) resource for continuous jobs ([#2466](https://github.com/databricks/terraform-provider-databricks/pull/2466)).
 * Added `full_refresh` attribute to the `pipeline_task` in [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) ([#2444](https://github.com/databricks/terraform-provider-databricks/pull/2444)).
 * Added support for Unity Catalog [databricks_metastore](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/metastore) data source ([#2492](https://github.com/databricks/terraform-provider-databricks/pull/2492)).
 * Added support for Unity Catalog [databricks_metastores](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/metastores) data source  ([#2017](https://github.com/databricks/terraform-provider-databricks/pull/2017)).
 * Added support for `USE_MARKETPLACE_ASSETS` privilege to [databricks_metastore](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/metastore) ([#2505](https://github.com/databricks/terraform-provider-databricks/pull/2505)).
 * Added support for boolean values in [databricks_sql_alert](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_alert) alerts ([#2506](https://github.com/databricks/terraform-provider-databricks/pull/2506)).
 * Added late jobs support (aka health conditions) in [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) resource ([2496](https://github.com/databricks/terraform-provider-databricks/pull/2496)).
 * Allow support for searching SQL Warehouses by name in [databricks_sql_warehouse](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_warehouse) data source ([#2458](https://github.com/databricks/terraform-provider-databricks/pull/2458)).
 * Added suppress diff for `aws_attributes.zone_id` with value `auto` in [databricks_instance_pool](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/instance_pool) ([#2518](https://github.com/databricks/terraform-provider-databricks/pull/2518)).
 * Changed test to use random catalog name in SQL table integration tests ([#2473](https://github.com/databricks/terraform-provider-databricks/pull/2473)).
 * Fixed [databricks_ip_access_list](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/ip_access_list) read ([#2515](https://github.com/databricks/terraform-provider-databricks/pull/2515)).
 * Fixed [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) resource to clear instance-specific attributes when `instance_pool_id` is specified ([#2507](https://github.com/databricks/terraform-provider-databricks/pull/2507)).
 * Fixed handling of comments in [databricks_sql_table](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_table) resource ([#2472](https://github.com/databricks/terraform-provider-databricks/pull/2472)).
 * Fixed model serving integration test using pip if the cluster is already running ([#2470](https://github.com/databricks/terraform-provider-databricks/pull/2470)).
 * Fixed provider after updating SDK to 0.13 ([#2494](https://github.com/databricks/terraform-provider-databricks/pull/2494)).
 * Updated [databricks_user](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/user) with `force = true` to check for error message prefix ([#2510](https://github.com/databricks/terraform-provider-databricks/pull/2510)).

Exporter:
 * Added exporter for [databricks_workspace_file](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/workspace_file) resource ([#2493](https://github.com/databricks/terraform-provider-databricks/pull/2493)).
 * Made resource names more unique to avoid duplicate resources errors ([#2452](https://github.com/databricks/terraform-provider-databricks/pull/2452)).

Documentation updates:
 * Added documentation notes about legacy cluster type & data access ([#2437](https://github.com/databricks/terraform-provider-databricks/pull/2437)).
 * Added one more item to the troubleshooting guide ([#2477](https://github.com/databricks/terraform-provider-databricks/pull/2477)).
 * Added clarification that [databricks_schema](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/schema) and [databricks_sql_table](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_table) should be imported by their full name, not just by name ([#2491](https://github.com/databricks/terraform-provider-databricks/pull/2491)).
 * Added more common issues for troubleshooting ([#2486](https://github.com/databricks/terraform-provider-databricks/pull/2486)).
 * Added additional documentation ([#2516](https://github.com/databricks/terraform-provider-databricks/pull/2516)).
 * Linked model serving docs to top level README ([#2474](https://github.com/databricks/terraform-provider-databricks/pull/2474)).

Other notes:
 * Added code owners for terraform-provider-databricks ([#2498](https://github.com/databricks/terraform-provider-databricks/pull/2498)).
 * Added support for new Delve binary name format ([#2497](https://github.com/databricks/terraform-provider-databricks/pull/2497)).
 * Configured merge queue for the provider ([#2533](https://github.com/databricks/terraform-provider-databricks/pull/2533)).
 * Removed unused dlvLoadConfig configuration from settings.json ([#2499](https://github.com/databricks/terraform-provider-databricks/pull/2499)).

Updated dependency versions:
 * Bump github.com/databricks/databricks-sdk-go from 0.13.0 to 0.14.1 ([#2523](https://github.com/databricks/terraform-provider-databricks/pull/2523)).
 * Bump golang.org/x/mod from 0.11.0 to 0.12.0 ([#2462](https://github.com/databricks/terraform-provider-databricks/pull/2462)).

## 1.21.0

 * Added condition_task to the [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) resource (private preview) ([#2459](https://github.com/databricks/terraform-provider-databricks/pull/2459)).
 * Added `AccountData`, `AccountClient` and define generic databricks data utilites for defining workspace and account-level data sources ([#2429](https://github.com/databricks/terraform-provider-databricks/pull/2429)).
 * Added documentation link to existing Databricks Terraform modules ([#2439](https://github.com/databricks/terraform-provider-databricks/pull/2439)).
 * Added experimental compute field to [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) resource ([#2401](https://github.com/databricks/terraform-provider-databricks/pull/2401)).
 * Added import example to doc for [databricks_group_member](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/group_member) resource ([#2453](https://github.com/databricks/terraform-provider-databricks/pull/2453)).
 * Added support for subscriptions in dashboards & alert SQL tasks in [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) ([#2447](https://github.com/databricks/terraform-provider-databricks/pull/2447)).
 * Fixed model serving integration test ([#2460](https://github.com/databricks/terraform-provider-databricks/pull/2460), [#2461](https://github.com/databricks/terraform-provider-databricks/pull/2461)).
 * Fixed [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) resource file arrival trigger parameter name ([#2438](https://github.com/databricks/terraform-provider-databricks/pull/2438)).
 * Fixed catalog_workspace_binding_test ([#2463](https://github.com/databricks/terraform-provider-databricks/pull/2463), [#2451](https://github.com/databricks/terraform-provider-databricks/pull/2451)).

## 1.20.0

 * Added suppress diff for `storage_location` in [databricks_sql_table](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_table) and [databricks_volume](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/volume) resources ([#2408](https://github.com/databricks/terraform-provider-databricks/pull/2408)).
 * Added option disable_as_user_deletion to disable instead of delete [databricks_user](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/user) and [databricks_service_principal](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/service_principal), enabled by default for user management at the account level ([#2378](https://github.com/databricks/terraform-provider-databricks/pull/2378)).
 * Added support for Unity Catalog [databricks_catalog_workspace_binding](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/catalog_workspace_binding) resource ([#2364](https://github.com/databricks/terraform-provider-databricks/pull/2364)).
 * Allowed assigning GCP SA in [databricks_sql_global_config](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_global_config) resource ([#2405](https://github.com/databricks/terraform-provider-databricks/pull/2405)).
 * Allowed changing `custom_tags` in [databricks_instance_pool](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/instance_pool) resource without recreating a pool ([#2400](https://github.com/databricks/terraform-provider-databricks/pull/2400)).
 * Bumped Go SDK version to 0.12.0 ([#2442](https://github.com/databricks/terraform-provider-databricks/pull/2442), [#2414](https://github.com/databricks/terraform-provider-databricks/pull/2414)).
 * Bumped golang.org/x/mod from 0.10.0 to 0.11.0 ([#2406](https://github.com/databricks/terraform-provider-databricks/pull/2406)).
 * Enabled model serving acceptance tests ([#2420](https://github.com/databricks/terraform-provider-databricks/pull/2420)).
 * Improved export of users/groups at the account level ([#2398](https://github.com/databricks/terraform-provider-databricks/pull/2398)).
 * Fixed [databricks_entitlements](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/entitlements) resource edge behaviour ([#2409](https://github.com/databricks/terraform-provider-databricks/pull/2409)).
 * Improved descriptions for [databricks_volume](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/volume) resource ([#2413](https://github.com/databricks/terraform-provider-databricks/pull/2413)).
 * Skipped TestAccClusterResource_CreateClusterWithLibraries integration test ([#2404](https://github.com/databricks/terraform-provider-databricks/pull/2404)).
 * Updated [databricks_grant](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/grant) resource with latest set of permissions ([#2399](https://github.com/databricks/terraform-provider-databricks/pull/2399)).
 * Used separate suppress diff function for [databricks_sql_table](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_table) and [databricks_volume](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/volume) ([#2412](https://github.com/databricks/terraform-provider-databricks/pull/2412)).

## 1.19.0

 * Added `run_as` and `run_as_user_name` fields to [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) ([#2388](https://github.com/databricks/terraform-provider-databricks/pull/2388)).
 * Added more attributes to [databricks_cluster_policy](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/cluster_policy) data source ([#2351](https://github.com/databricks/terraform-provider-databricks/pull/2351)).
 * Added `force_destroy` to [databricks_external_location](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/external_location) resource ([#2353](https://github.com/databricks/terraform-provider-databricks/pull/2353)).
 * Added less permissive principal for self assuming UC role ([#2386](https://github.com/databricks/terraform-provider-databricks/pull/2386)).
 * Document `workload_type` block in [databricks_cluster](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/cluster) resource ([#2370](https://github.com/databricks/terraform-provider-databricks/pull/2370)).
 * Expose `read_only` option in [databricks_external_location](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/external_location) and [databricks_storage_credential](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/storage_credential) ([#2391](https://github.com/databricks/terraform-provider-databricks/pull/2391)).
 * Migrate [databricks_catalog](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/catalog) to Go SDK ([#2357](https://github.com/databricks/terraform-provider-databricks/pull/2357)).
 * Update `FullName` in [databricks_volume](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/volume) resource only if no error is returned. ([#2374](https://github.com/databricks/terraform-provider-databricks/pull/2374)).
 * Updated test to use databricks-sdk library as Faker library isn't being installed correctly ([#2397](https://github.com/databricks/terraform-provider-databricks/pull/2397)).
 * Explicit conversion of strings to int64. ([#2346](https://github.com/databricks/terraform-provider-databricks/pull/2346)).

## 1.18.0

 * Added [databricks_volume](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/volume) resource in Unity Catalog ([#2324](https://github.com/databricks/terraform-provider-databricks/pull/2324)).
 * Added [databricks_model_serving](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/model_serving) support to [databricks_permissions](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/permissions) ([#2331](https://github.com/databricks/terraform-provider-databricks/pull/2331)).
 * Added suppress diff for URL change that only changes `/` (remove or add) in UC resources ([#2336](https://github.com/databricks/terraform-provider-databricks/pull/2336)).
 * Fixed attributes typo in SCIM API ([#2344](https://github.com/databricks/terraform-provider-databricks/pull/2344)).
 * Fixed updates for [databricks_share](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/share) resource ([#2307](https://github.com/databricks/terraform-provider-databricks/pull/2307)).
 * Updated documentation for [databricks_service_principal_secret](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/service_principal_secret) ([#2332](https://github.com/databricks/terraform-provider-databricks/pull/2332)).
 * Updated documentation for troubleshooting guide with a typical error when creating groups/users on the Account level ([#2338](https://github.com/databricks/terraform-provider-databricks/pull/2338)).
 * Other testing infrastructure improvements ([#2350](https://github.com/databricks/terraform-provider-databricks/pull/2350), [#2355](https://github.com/databricks/terraform-provider-databricks/pull/2355), [#2358](https://github.com/databricks/terraform-provider-databricks/pull/2358)).

Updated dependency versions:

 * Bump github.com/hashicorp/hcl/v2 from 2.16.2 to 2.17.0 ([#2359](https://github.com/databricks/terraform-provider-databricks/pull/2359)).
 * Bump github.com/stretchr/testify from 1.8.3 to 1.8.4 ([#2354](https://github.com/databricks/terraform-provider-databricks/pull/2354)).
 * Bump github.com/zclconf/go-cty from 1.13.1 to 1.13.2 ([#2329](https://github.com/databricks/terraform-provider-databricks/pull/2329)).

## 1.17.0

 * **Removed support for releasing 32-bit binaries** ([#2315](https://github.com/databricks/terraform-provider-databricks/pull/2315), [#2320](https://github.com/databricks/terraform-provider-databricks/pull/2320)).
 * Added more information on impact of using a cluster policy in [databricks_cluster](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/cluster) resource ([#2313](https://github.com/databricks/terraform-provider-databricks/pull/2313)).
 * Added missing `serverless` option to [databricks_pipeline](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/pipeline) ([#2308](https://github.com/databricks/terraform-provider-databricks/pull/2308)).
 * Updated `channel` and `edition` values in [databricks_pipeline](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/pipeline) docs ([#2322](https://github.com/databricks/terraform-provider-databricks/pull/2322)).
 * Automatically add `CAN_MANAGE` permission on [databricks_instance_pool](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/instance_pool) for calling user ([#2298](https://github.com/databricks/terraform-provider-databricks/pull/2298)).
 * Migrated [databricks_ip_access_list](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/ip_access_list) resource to Go SDK ([#2306](https://github.com/databricks/terraform-provider-databricks/pull/2306)).

Updated dependency versions:

 * Bump github.com/stretchr/testify from 1.8.2 to 1.8.3 ([#2317](https://github.com/databricks/terraform-provider-databricks/pull/2317)).
 * Bump github.com/databricks/databricks-sdk-go from v0.8.1 to v0.9.0 ([#2327](https://github.com/databricks/terraform-provider-databricks/pull/2327)).

## 1.16.1

 * Added [databricks_service_principal_secret](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/service_principal_secret) documentation ([#2296](https://github.com/databricks/terraform-provider-databricks/pull/2296)).
 * Documentation about Unity Catalog integration in [databricks_pipeline](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/pipeline) ([#2289](https://github.com/databricks/terraform-provider-databricks/pull/2289)).
 * Exporter: don't emit [databricks_permissions](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/permissions) for `/Shared` directory ([#2288](https://github.com/databricks/terraform-provider-databricks/pull/2288)).
 * Fix dependency when exporting [databricks_service_principal_role](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/service_principal_role) ([#2285](https://github.com/databricks/terraform-provider-databricks/pull/2285)).
 * Removed reference to deprecated [databricks_group_instance_profile](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/group_instance_profile) from [databricks_group_role](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/group_role) docs ([#2290](https://github.com/databricks/terraform-provider-databricks/pull/2290)).
 * Updated documentation for [databricks_secret_scope](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/secret_scope) resource ([#2297](https://github.com/databricks/terraform-provider-databricks/pull/2297)).

## 1.16.0

 * Added [databricks_workspace_file](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/workspace_file) resource ([#2266](https://github.com/databricks/terraform-provider-databricks/pull/2266)).
 * Added `notification_settings` block to [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) resource ([#2276](https://github.com/databricks/terraform-provider-databricks/pull/2276)).
 * Added missing permission to [databricks_aws_crossaccount_policy](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/aws_crossaccount_policy) data source ([#2283](https://github.com/databricks/terraform-provider-databricks/pull/2283)).
 * Fixed [databricks_sql_table](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_table) doc ([#2281](https://github.com/databricks/terraform-provider-databricks/pull/2281)).
 * Fixed doc on tag propagation and tag conflict for [databricks_instance_pool](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/instance_pool) resource ([#2242](https://github.com/databricks/terraform-provider-databricks/pull/2242)).
 * Document & export [databricks_workspace_conf](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/workspace_conf) parameters for legacy init scripts ([#2280](https://github.com/databricks/terraform-provider-databricks/pull/2280)).
 * Removed deprecated `CREATE_VIEW` from code of [databricks_grants](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/grants) resource ([#2230](https://github.com/databricks/terraform-provider-databricks/pull/2230)).
 * Exporter: relax handling of problematic files/directories ([#2258](https://github.com/databricks/terraform-provider-databricks/pull/2258)).

## 1.15.0

 * Added _experimental_ [databricks_sql_table](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_table) resource to manage tables in the Unity Catalog ([#2213](https://github.com/databricks/terraform-provider-databricks/pull/2213)).
 * Added [databricks_pipelines](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/pipelines) data source ([#2202](https://github.com/databricks/terraform-provider-databricks/pull/2202)).
 * Added missing permissions to [databricks_grants](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/grants) resource ([#2260](https://github.com/databricks/terraform-provider-databricks/pull/2260)).
 * Added support for running [databricks_cluster](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/cluster) init scripts from workspace files ([#2251](https://github.com/databricks/terraform-provider-databricks/pull/2251)).
 * Added suppress diff for `run_if` in [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) ([#2229](https://github.com/databricks/terraform-provider-databricks/pull/2229)).
 * Fixed delete permission error for manually removed [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) ([#2225](https://github.com/databricks/terraform-provider-databricks/pull/2225)).
 * Documented `fleet` selector in the [databricks_node_type](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/node_type) data source ([#2261](https://github.com/databricks/terraform-provider-databricks/pull/2261)).
 * Documented example for [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) tags ([#2235](https://github.com/databricks/terraform-provider-databricks/pull/2235)).
 * Documented more Azure Databricks workspaces creation ([#2228](https://github.com/databricks/terraform-provider-databricks/pull/2228), [#2248](https://github.com/databricks/terraform-provider-databricks/pull/2248)).
 * Removed note about Azure KeyVault-based [databricks_secret_scope](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/secret_scope) and service principals ([#2231](https://github.com/databricks/terraform-provider-databricks/pull/2231)).
 * Migrated [databricks_cluster_policy](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/cluster_policy) to Go SDK ([#2250](https://github.com/databricks/terraform-provider-databricks/pull/2250)).
 * Updated [databricks_sql_endpoint](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_endpoint) docs for Azure DBSQL Serverless GA ([#2249](https://github.com/databricks/terraform-provider-databricks/pull/2249), [#2254](https://github.com/databricks/terraform-provider-databricks/pull/2254)).
 * Updated policy definition in [databricks_aws_crossaccount_policy](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/aws_crossaccount_policy) data source ([#2262](https://github.com/databricks/terraform-provider-databricks/pull/2262)).
 * Bumped `databricks-sdk-go` from 0.7.0 to 0.8.0 ([#2259](https://github.com/databricks/terraform-provider-databricks/pull/2259)).

## 1.14.3

 * Add exporter support for [databricks_sql_alert](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_alert) ([#2207](https://github.com/databricks/terraform-provider-databricks/pull/2207)).
 * Add support for notifications in [databricks_pipeline](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/pipeline) resource ([#2218](https://github.com/databricks/terraform-provider-databricks/pull/2218)).
 * Add suppress diff for `path` argument of [databricks_directory](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/directory) resource when it ends with `/` ([#2204](https://github.com/databricks/terraform-provider-databricks/pull/2204)).
 * Exporter: don't export `storage` when [databricks_pipeline](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/pipeline) was created with default value ([#2203](https://github.com/databricks/terraform-provider-databricks/pull/2203)).
 * Exporter: phase 1 of support for Account-level exports ([#2205](https://github.com/databricks/terraform-provider-databricks/pull/2205)).
 * Fixed regression in [databricks_group](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/group) data not to require workspace admin privileges ([#2210](https://github.com/databricks/terraform-provider-databricks/pull/2210)).
 * Suppress diff for `edition` attribute in [databricks_pipeline](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/pipeline) ([#2219](https://github.com/databricks/terraform-provider-databricks/pull/2219)).

## 1.14.2

 * Explicitly include SCIM attributes for [databricks_group](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/group), [databricks_user](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/user), [databricks_user_role](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/user_role), [databricks_group_role](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/group_role), [databricks_group_member](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/group_member), [databricks_group_instance_profile](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/group_instance_profile), [databricks_user](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/user) data, [databricks_group](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/group) data, and [databricks_entitlement](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/entitlement) resources ([#2200](https://github.com/databricks/terraform-provider-databricks/pull/2200)).
 * Added SQL File task to [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) resource ([#2199](https://github.com/databricks/terraform-provider-databricks/pull/2199)).
 * Fix diff detection for [databricks_share](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/share) resource ([#2197](https://github.com/databricks/terraform-provider-databricks/pull/2197)).
 * Guide to deploy Databricks on Azure with PrivateLink - Standard deployment ([#2066](https://github.com/databricks/terraform-provider-databricks/pull/2066)).
 * Remove `CREATE_VIEW` from [databricks_grants](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/grants) doc ([#2198](https://github.com/databricks/terraform-provider-databricks/pull/2198)).

Updated dependency versions:

 * Bump github.com/databricks/databricks-sdk-go from 0.6.0 to 0.7.0 ([#2195](https://github.com/databricks/terraform-provider-databricks/pull/2195)).
 * Bump golang.org/x/mod from 0.9.0 to 0.10.0 ([#2192](https://github.com/databricks/terraform-provider-databricks/pull/2192)).

## 1.14.1

 * Allow rotating `token` block in [databricks_mws_workspaces](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/mws_workspaces) resource by only changing `comment` field ([#2114](https://github.com/databricks/terraform-provider-databricks/pull/2114)).
 * Exclude roles in SCIM API list calls to reduce load on Databricks SCIM service ([#2181](https://github.com/databricks/terraform-provider-databricks/pull/2181)).
 * Fixed [databricks_service_principals](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/service_principals) data source issue with empty filter ([#2185](https://github.com/databricks/terraform-provider-databricks/pull/2185)).
 * Update Go SDK to v0.6.0 ([#2186](https://github.com/databricks/terraform-provider-databricks/pull/2186)).

## 1.14.0

 * Added caching of SCIM Me API call to [databricks_current_user](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/current_user) data source and [databricks_permissions](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/permissions) resource. ([#2170](https://github.com/databricks/terraform-provider-databricks/pull/2170)).
 * Added example for automatically rotating a token ([#2135](https://github.com/databricks/terraform-provider-databricks/pull/2135)).
 * Added `run_if` condition in [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) resource ([#2125](https://github.com/databricks/terraform-provider-databricks/pull/2125)).
 * Added File Arrival trigger to [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) resource ([#2142](https://github.com/databricks/terraform-provider-databricks/pull/2142)).
 * Added `source` parameter for `spark_python_task` in [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) ([#2157](https://github.com/databricks/terraform-provider-databricks/pull/2157)).
 * Automatically add `CAN_MANAGE` permission on [databricks_sql_endpoint](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_endpoint) for calling user ([#2168](https://github.com/databricks/terraform-provider-databricks/pull/2168)).
 * Deprecated `enable_serverless_compute` on [databricks_sql_global_config](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_global_config) resource ([#2139](https://github.com/databricks/terraform-provider-databricks/pull/2139)).
 * Document `enable_serverless_compute` API changes in [databricks_sql_endpoint](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_endpoint) resource ([#2137](https://github.com/databricks/terraform-provider-databricks/pull/2137)).
 * Fix edge cases for [databricks_permissions](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/permissions) resource ([#2158](https://github.com/databricks/terraform-provider-databricks/pull/2158)).
 * Migrate [databricks_cluster_policy](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/cluster_policy) data source to Go SDK ([#2155](https://github.com/databricks/terraform-provider-databricks/pull/2155)).
 * Update the description of `ip_addresses` parameter in [databricks_ip_access_list](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/ip_access_list) docs ([#2116](https://github.com/databricks/terraform-provider-databricks/pull/2116)).

Updated dependency versions:

 * Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.25.0 to 2.26.1 ([#2140](https://github.com/databricks/terraform-provider-databricks/pull/2140)).
 * Bump github.com/zclconf/go-cty from 1.13.0 to 1.13.1 ([#2124](https://github.com/databricks/terraform-provider-databricks/pull/2124)).

## 1.13.0

 * Added [databricks_sql_alert](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_alert) resource ([#2047](https://github.com/databricks/terraform-provider-databricks/pull/2047)).
 * Added experimental GCP support for more resources ([#2088](https://github.com/databricks/terraform-provider-databricks/pull/2088), [#2091](https://github.com/databricks/terraform-provider-databricks/pull/2091), [#2090](https://github.com/databricks/terraform-provider-databricks/pull/2090), [#2089](https://github.com/databricks/terraform-provider-databricks/pull/2089), [#2080](https://github.com/databricks/terraform-provider-databricks/pull/2080)).
 * Added suppress diff to `default_data_access_config_id` in [databricks_metastore](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/metastore) ([#2111](https://github.com/databricks/terraform-provider-databricks/pull/2111)).
 * Export `initialize_file_system` for ADLS Gen2 mounts ([#2107](https://github.com/databricks/terraform-provider-databricks/pull/2107)).
 * Enforce formatting on pull requests ([#2099](https://github.com/databricks/terraform-provider-databricks/pull/2099)).
 * Added `MatchRegexp` match type for partial matches ([#2102](https://github.com/databricks/terraform-provider-databricks/pull/2102)).
 * Skip model serving integration test on GCP ([#2098](https://github.com/databricks/terraform-provider-databricks/pull/2098)).
 * More doc fixes ([#2108](https://github.com/databricks/terraform-provider-databricks/pull/2108), [#2106](https://github.com/databricks/terraform-provider-databricks/pull/2106)).

Updated dependency versions:

 * Bump github.com/databricks/databricks-sdk-go from 0.4.0 to 0.5.0 ([#2096](https://github.com/databricks/terraform-provider-databricks/pull/2096), [#2112](https://github.com/databricks/terraform-provider-databricks/pull/2112)).
 * Bump github.com/hashicorp/hcl/v2 from 2.16.1 to 2.16.2 ([#2097](https://github.com/databricks/terraform-provider-databricks/pull/2097)).

### 1.12.0

 * Added [databricks_model_serving](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/model_serving) resource ([#2054](https://github.com/databricks/terraform-provider-databricks/pull/2054)).
 * Added Unity Catalog on GCP support ([#2000](https://github.com/databricks/terraform-provider-databricks/pull/2000)).
 * Deprecate `schedule` block in [databricks_sql_query](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_query) resource ([#2078](https://github.com/databricks/terraform-provider-databricks/pull/2078)).
 * Fixed InitScripts Type to work with GCS and ABFS in [databricks_cluster](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/cluster) resource ([#2067](https://github.com/databricks/terraform-provider-databricks/pull/2067)).
 * Added more testing for [databricks_tables](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/tables) data source ([#2075](https://github.com/databricks/terraform-provider-databricks/pull/2075)).
 * Added more testing for [databricks_schemas](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/schemas) data source ([#2074](https://github.com/databricks/terraform-provider-databricks/pull/2074)).
 * Migrated [databricks_node_type](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/node_type) data source to Go SDK ([#2070](https://github.com/databricks/terraform-provider-databricks/pull/2070)).
 * Migrated [databricks_schemas](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/schemas) data to Go SDK ([#2065](https://github.com/databricks/terraform-provider-databricks/pull/2065)).
 * Migrated [databricks_tables](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/tables) data source to SDK ([#2068](https://github.com/databricks/terraform-provider-databricks/pull/2068)).
 * Migrated [databricks_views](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/views) data source to Go SDK ([#2073](https://github.com/databricks/terraform-provider-databricks/pull/2073)).
 * Migrated [databricks_git_credential](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/git_credential) to Go SDK ([#2069](https://github.com/databricks/terraform-provider-databricks/pull/2069)).
 * Migrated [databricks_shares](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/shares) data source to Go SDK ([#2072](https://github.com/databricks/terraform-provider-databricks/pull/2072)).

Updated dependency versions:

 * Bump github.com/databricks/databricks-sdk-go from v0.3.3 to v0.4.0 ([#2086](https://github.com/databricks/terraform-provider-databricks/pull/2086)).
 * Bump golang.org/x/mod from 0.8.0 to 0.9.0 ([#2076](https://github.com/databricks/terraform-provider-databricks/pull/2076)).

### 1.11.1

 * Databricks on Azure with PrivateLink - Simplified deployment ([#1977](https://github.com/databricks/terraform-provider-databricks/pull/1977)).
 * Deterministic diff for [databricks_permissions](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/permissions) ([#2059](https://github.com/databricks/terraform-provider-databricks/pull/2059)).

### 1.11.0

 * Added `force_delete_home_dir` and `force_delete_repos` attributes to [databricks_user](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/user) and [databricks_service_principal](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/service_principal) resources ([#2032](https://github.com/databricks/terraform-provider-databricks/pull/2032)).
 * Added docs for `continuous` block in the [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) resource ([#2048](https://github.com/databricks/terraform-provider-databricks/pull/2048)).
 * Exporter: `databricks_permissions` for [databricks_notebook](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/notebook) & [databricks_directory](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/directory) ([#1908](https://github.com/databricks/terraform-provider-databricks/pull/1908)).
 * Improve error messages for [databricks_permissions](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/permissions) ([#2055](https://github.com/databricks/terraform-provider-databricks/pull/2055)).
 * Make reflect resource utility friendly with Go SDK ([#2051](https://github.com/databricks/terraform-provider-databricks/pull/2051)).

Updated dependency versions:

 * Bump github.com/stretchr/testify from 1.8.1 to 1.8.2 ([#2049](https://github.com/databricks/terraform-provider-databricks/pull/2049)).

### 1.10.1

 * Migrated [databricks_catalogs](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/catalogs) data to Go SDK ([#2038](https://github.com/databricks/terraform-provider-databricks/pull/2038)).
 * Migrated [databricks_current_user](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/current_user) data source to Go SDK ([#2037](https://github.com/databricks/terraform-provider-databricks/pull/2037)).
 * Migrated [databricks_workspace_conf](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/workspace_conf) to Go SDK ([#2035](https://github.com/databricks/terraform-provider-databricks/pull/2035)).
 * Added new feature ([#2015](https://github.com/databricks/terraform-provider-databricks/pull/2015)).
 * Minor documentation improvements ([#2039](https://github.com/databricks/terraform-provider-databricks/pull/2039)).

Updated dependency versions:

 * Bump github.com/databricks/databricks-sdk-go from v0.3.2 to v0.3.3 ([#2044](https://github.com/databricks/terraform-provider-databricks/pull/2044)).
 * Bump github.com/zclconf/go-cty from 1.12.1 to 1.13.0 ([#2043](https://github.com/databricks/terraform-provider-databricks/pull/2043)).

### 1.10.0

 * Migrated client and configuration layer to [`github.com/databricks/databricks-sdk-go`](http://github.com/databricks/databricks-sdk-go) ([#1848](https://github.com/databricks/terraform-provider-databricks/pull/1848)).
 * Added new attribute to [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) resource ([#1988](https://github.com/databricks/terraform-provider-databricks/pull/1988)).
 * Added `history_data_sharing_status`, `partition` and `status` fields to [databricks_share](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/share) ([#2013](https://github.com/databricks/terraform-provider-databricks/pull/2013)).
 * Added `start_version` and `cdf_enabled` support for [databricks_share](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/share) ([#2007](https://github.com/databricks/terraform-provider-databricks/pull/2007)).
 * Fixed [databricks_share](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/share) update logic ([#2022](https://github.com/databricks/terraform-provider-databricks/pull/2022)).
 * Updated GCP guide with new permissions required to deploy [databricks_mws_networks](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/mws_networks) ([#1999](https://github.com/databricks/terraform-provider-databricks/pull/1999)).
 * Minor doc changes to address [#1916](https://github.com/databricks/terraform-provider-databricks/issues/1916), [#1937](https://github.com/databricks/terraform-provider-databricks/issues/1937), [#2010](https://github.com/databricks/terraform-provider-databricks/issues/2010) ([#2012](https://github.com/databricks/terraform-provider-databricks/pull/2012)).
 * Various documentation improvements ([#1981](https://github.com/databricks/terraform-provider-databricks/pull/1981), [#1998](https://github.com/databricks/terraform-provider-databricks/pull/1998), [#2002](https://github.com/databricks/terraform-provider-databricks/pull/2002), [#2016](https://github.com/databricks/terraform-provider-databricks/pull/2016), [#2019](https://github.com/databricks/terraform-provider-databricks/pull/2019)).
 * Integration testing improvements ([#1935](https://github.com/databricks/terraform-provider-databricks/pull/1935), [#2024](https://github.com/databricks/terraform-provider-databricks/pull/2024)).

Updated dependency versions:

 * Removed unused `github.com/Azure/go-autorest/autorest` and `github.com/mitchellh/go-homedir` ([#2026](https://github.com/databricks/terraform-provider-databricks/pull/2026)).
 * Bump Go from 1.18 to 1.19 ([#2029](https://github.com/databricks/terraform-provider-databricks/pull/2029)).
 * Bump CodeQL to v2 ([#2027](https://github.com/databricks/terraform-provider-databricks/pull/2027)).
 * Bump github.com/golang-jwt/jwt/v4 from 4.4.3 to 4.5.0 ([#2030](https://github.com/databricks/terraform-provider-databricks/pull/2030)).
 * Bump github.com/hashicorp/hcl/v2 from 2.15.0 to 2.16.1 ([#1971](https://github.com/databricks/terraform-provider-databricks/pull/1971), [#1997](https://github.com/databricks/terraform-provider-databricks/pull/1997)).
 * Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.24.1 to 2.25.0 ([#2031](https://github.com/databricks/terraform-provider-databricks/pull/2031)).
 * Bump golang.org/x/mod from 0.7.0 to 0.8.0 ([#1996](https://github.com/databricks/terraform-provider-databricks/pull/1996)).
 * Bump golang.org/x/net from 0.6.0 to 0.7.0 ([#2028](https://github.com/databricks/terraform-provider-databricks/pull/2028)).
 * Bump google.golang.org/api from 0.108.0 to 0.110.0 ([#1984](https://github.com/databricks/terraform-provider-databricks/pull/1984), [#1995](https://github.com/databricks/terraform-provider-databricks/pull/1995)).

### 1.9.2

 * Added `file` library type to [databricks_pipeline](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/pipeline) resource ([#1975](https://github.com/databricks/terraform-provider-databricks/pull/1975)).
 * Added new node type selector to [databricks_node_type](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/node_type) data source ([#1957](https://github.com/databricks/terraform-provider-databricks/pull/1957)).
 * Removed check for creation of Azure Key Vault-based [databricks_secret_scope](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/secret_scope) using Service Principals ([#1965](https://github.com/databricks/terraform-provider-databricks/pull/1965)).
 * Reworked Graviton selection logic in [databricks_node_type](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/node_type) data source ([#1974](https://github.com/databricks/terraform-provider-databricks/pull/1974)).
 * Don't consider GPU-based nodes in [databricks_node_type](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/node_type) data source until explicitly requested ([#1978](https://github.com/databricks/terraform-provider-databricks/pull/1978)).
 
Documentation:

 * Added SQL example for [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) documentation ([#1967](https://github.com/databricks/terraform-provider-databricks/pull/1967)).
 * Added description to `max_clusters_per_user` in [databricks_cluster_policy](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/cluster_policy) ([#1959](https://github.com/databricks/terraform-provider-databricks/pull/1959)).
 * Clarify SQL warehouse types for `sql_task` in [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) ([#1964](https://github.com/databricks/terraform-provider-databricks/pull/1964)).
 * Moved tags to `default_tags` in AWS provider examples ([#1969](https://github.com/databricks/terraform-provider-databricks/pull/1969)).
 * Added documentation around `webhooks_notifications` for [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) ([#1968](https://github.com/databricks/terraform-provider-databricks/pull/1968)).
 * Removed deprecated params from bucket doc examples ([#1970](https://github.com/databricks/terraform-provider-databricks/pull/1970)).

### 1.9.1

* Added `max_clusters_per_user` to [databricks_cluster_policy](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/cluster_policy) resource ([#1955](https://github.com/databricks/terraform-provider-databricks/pull/1955)).
 * Added support for `cluster_mount_infos` in [databricks_cluster](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/cluster) ([#1945](https://github.com/databricks/terraform-provider-databricks/pull/1945)).
 * Added `parent` attribute to [databricks_sql_query](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_query) and [databricks_sql_dashboard](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_dashboard) ([#1949](https://github.com/databricks/terraform-provider-databricks/pull/1949)).
 * Added support for `iam_role_arn` in [databricks_instance_profile](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/instance_profile) ([#1943](https://github.com/databricks/terraform-provider-databricks/pull/1943)).
 * Use `zone_id = "auto"` for [databricks_mount](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/mount) on S3 ([#1932](https://github.com/databricks/terraform-provider-databricks/pull/1932)).
 * Fixed [databricks_cluster](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/cluster) resource missing library definitions issues ([#1925](https://github.com/databricks/terraform-provider-databricks/pull/1925)).
 * Remove `computed` annotation for `aws_attributes` ([#1933](https://github.com/databricks/terraform-provider-databricks/pull/1933)).

Updated dependency versions:

 * Bump github.com/Azure/go-autorest/autorest/adal from 0.9.21 to 0.9.22 ([#1953](https://github.com/databricks/terraform-provider-databricks/pull/1953)).
 * Bump github.com/Azure/go-autorest/autorest/azure/auth ([#1952](https://github.com/databricks/terraform-provider-databricks/pull/1952)).
 * Bump google.golang.org/api from 0.106.0 to 0.107.0 ([#1929](https://github.com/databricks/terraform-provider-databricks/pull/1929)).
 * Bump google.golang.org/api from 0.107.0 to 0.108.0 ([#1954](https://github.com/databricks/terraform-provider-databricks/pull/1954)).

### 1.9.0

 * Added [databricks_mws_credentials](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/mws_credentials) data source ([#1909](https://github.com/databricks/terraform-provider-databricks/pull/1909)).
 * Fixed incorrect diff for `num_clusters` and `min_num_clusters` for stopped [databricks_sql_endpoint](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_endpoint) ([#1927](https://github.com/databricks/terraform-provider-databricks/pull/1927)).
 * Fixed `privilege_assignment` in [databricks_sql_permissions](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_permissions) ([#1872](https://github.com/databricks/terraform-provider-databricks/pull/1872)).
 * Documentation improvements ([#1921](https://github.com/databricks/terraform-provider-databricks/pull/1921), [#1923](https://github.com/databricks/terraform-provider-databricks/pull/1923)).

Updated dependency versions:

 * Bump github.com/hashicorp/go-retryablehttp from 0.7.1 to 0.7.2 ([#1913](https://github.com/databricks/terraform-provider-databricks/pull/1913)).
 * Bump google.golang.org/api from 0.105.0 to 0.106.0 ([#1914](https://github.com/databricks/terraform-provider-databricks/pull/1914)).

### 1.8.0

 * Added support for GCP (Public Preview) with state upgraders for [databricks_mws_workspaces](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/mws_workspaces) ([#1871](https://github.com/databricks/terraform-provider-databricks/pull/1871), [#1886](https://github.com/databricks/terraform-provider-databricks/pull/1886), [#1748](https://github.com/databricks/terraform-provider-databricks/pull/1748), [#1879](https://github.com/databricks/terraform-provider-databricks/pull/1879)).
 * Added documentation for [databricks_permission_assignment](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/permission_assignment) ([#1880](https://github.com/databricks/terraform-provider-databricks/pull/1880)).
 * Added `azure_attributes` to clusters in [databricks_pipeline](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/pipeline) resource ([#1854](https://github.com/databricks/terraform-provider-databricks/pull/1854)).
 * Added [databricks_instance_pool](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/instance_pool) data source ([#1907](https://github.com/databricks/terraform-provider-databricks/pull/1907)).
 * Added [databricks_provider](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/provider) resource for Delta Sharing ([#1572](https://github.com/databricks/terraform-provider-databricks/pull/1572)).
 * Added acceptance tests for [databricks_cluster](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/cluster) data and [databricks_sql_endpoint](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/sql_endpoint) data ([#1882](https://github.com/databricks/terraform-provider-databricks/pull/1882)).
 * Added search by name for [databricks_cluster](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/cluster) data source ([#1901](https://github.com/databricks/terraform-provider-databricks/pull/1901)).
 * Added support for sparse checkouts in the [databricks_repo](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/repo) ([#1869](https://github.com/databricks/terraform-provider-databricks/pull/1869)).
 * Added `config_reference` attribute to [databricks_secret](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/secret) for easier reference from Spark configuration ([#1898](https://github.com/databricks/terraform-provider-databricks/pull/1898)).
 * Added `datarbicks_directory` data source ([#1902](https://github.com/databricks/terraform-provider-databricks/pull/1902)).
 * Extended [databricks_catalog](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/catalog) resource to support Delta Sharing Catalog ([#1887](https://github.com/databricks/terraform-provider-databricks/pull/1887)).
 * Improved [databricks_metastore_assignment](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/metastore_assignment) resource ([#1900](https://github.com/databricks/terraform-provider-databricks/pull/1900)).
 * Improved [databricks_grants](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/grants) error messages ([#1888](https://github.com/databricks/terraform-provider-databricks/pull/1888)).
 * Improved [databricks_group](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/group) documentation ([#1816](https://github.com/databricks/terraform-provider-databricks/pull/1816)).
 * Improved `data_security_mode` documentation for [databricks_cluster](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/cluster) ([#1830](https://github.com/databricks/terraform-provider-databricks/pull/1830)).
 * Replace the incorrect `id` attribute with the correct `sp_id` in service principal data source docs example ([#1873](https://github.com/databricks/terraform-provider-databricks/pull/1873)).
 * Use NVMe disk size for selection in the [databricks_node_type](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/node_type) data source ([#1856](https://github.com/databricks/terraform-provider-databricks/pull/1856)).
 * Allowed to override resource ID for specific data source ([#1885](https://github.com/databricks/terraform-provider-databricks/pull/1885)).

Experimental Resource Exporter

 * Added command-line flag to export all users and service principals ([#1893](https://github.com/databricks/terraform-provider-databricks/pull/1893)).
 * Added emit for secret scopes from Spark Conf and Env Vars ([#1897](https://github.com/databricks/terraform-provider-databricks/pull/1897)).
 * Fixed issue with exporting of resources with simple name ([#1891](https://github.com/databricks/terraform-provider-databricks/pull/1891)).
 * Generate references that are matching only to a prefix ([#1890](https://github.com/databricks/terraform-provider-databricks/pull/1890)).
 * Updated list of supported services and associated resources ([#1894](https://github.com/databricks/terraform-provider-databricks/pull/1894)).
 * Use [databricks_group_role](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/group_role) instead of deprecated [databricks_group_instance_profile](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/group_instance_profile) ([#1905](https://github.com/databricks/terraform-provider-databricks/pull/1905)).
 * Use `dbfs_path` as reference for init scripts in clusters ([#1892](https://github.com/databricks/terraform-provider-databricks/pull/1892)).
 * Logging improvements ([#1895](https://github.com/databricks/terraform-provider-databricks/pull/1895)).
 
Updated dependency versions:

 * Bump google.golang.org/api from 0.104.0 to 0.105.0 ([#1863](https://github.com/databricks/terraform-provider-databricks/pull/1863)).

### 1.7.0

 * Added support for Unity Catalog on GCP through `gcp_service_account_key` ([#1836](https://github.com/databricks/terraform-provider-databricks/pull/1836)).
 * Added support for Jupyter (`.ipynb`) files in [databricks_notebook](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/databricks_notebook) resource ([#1801](https://github.com/databricks/terraform-provider-databricks/pull/1801)).
 * Added [databricks_cluster_policy](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/cluster_policy) data source ([#1792](https://github.com/databricks/terraform-provider-databricks/pull/1792)).
 * Added `catalog` parameter to `dbt_task` in [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) resource ([#1774](https://github.com/databricks/terraform-provider-databricks/pull/1774)).
 * Added `home` and `repos` attributes to [databricks_user](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/user) and [databricks_service_principal](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/service_principal) ([#1791](https://github.com/databricks/terraform-provider-databricks/pull/1791)).
 * Added `storage_root` parameter in [databricks_catalog](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/catalog) and [databricks_schema](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/schema) resources ([#1784](https://github.com/databricks/terraform-provider-databricks/pull/1784)).
 * Added exporting of [databricks_service_principal](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/service_principal) resource ([#1811](https://github.com/databricks/terraform-provider-databricks/pull/1811)).
 * Added notebook `source` parameter for [databricks_jobs](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/jobs) data ([#1837](https://github.com/databricks/terraform-provider-databricks/pull/1837)).
 * Added support for init scripts on ADLS in [databricks_cluster](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/cluster) and related resources ([#1845](https://github.com/databricks/terraform-provider-databricks/pull/1845)).
 * Added `definiton` attribute to [databricks_cluster_policy](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/cluster_policy) data ([#1834](https://github.com/databricks/terraform-provider-databricks/pull/1834)).
 * Added validation for preview account API paths ([#1795](https://github.com/databricks/terraform-provider-databricks/pull/1795)).
 * Made `lifetime_seconds`  & `comment` optional for [databricks_obo_token](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/obo_token) ([#1844](https://github.com/databricks/terraform-provider-databricks/pull/1844)).
 * Force replace when changing `display_name` of [databricks_service_principal](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/service_principal) ([#1783](https://github.com/databricks/terraform-provider-databricks/pull/1783)).
 * We can now export `admin` and `users` groups as [databricks_group](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/group) data, not resource ([#1721](https://github.com/databricks/terraform-provider-databricks/pull/1721)).
 * We can now export predefined [databricks_cluster_policy](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/cluster_policy) as data sources ([#1815](https://github.com/databricks/terraform-provider-databricks/pull/1815)).
 * We can now export [databricks_sql_global_config](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_global_config) resource ([#1806](https://github.com/databricks/terraform-provider-databricks/pull/1806)).
 * Fixed reference to [databricks_cluster_policy](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/cluster_policy) in [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) during export ([#1813](https://github.com/databricks/terraform-provider-databricks/pull/1813)).
 * We can now explicitly normalize variable names during export ([#1835](https://github.com/databricks/terraform-provider-databricks/pull/1835)).
 * We can now fail fast in `import.sh` when can't import resource ([#1814](https://github.com/databricks/terraform-provider-databricks/pull/1814)).
 * Fixed multidomain [databricks_user](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/user) in exporter ([#1722](https://github.com/databricks/terraform-provider-databricks/pull/1722)).
 * Fixed missing attributes from [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/job) data source ([#1798](https://github.com/databricks/terraform-provider-databricks/pull/1798)).
 * Suppressed `network` diff for GCP workspaces with Databricks-managed VPC ([#1847](https://github.com/databricks/terraform-provider-databricks/pull/1847)).
 * Improved export of DBSQL objects ([#1788](https://github.com/databricks/terraform-provider-databricks/pull/1788)).
 * Documented `webhook_notifications` for [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) resource ([#1842](https://github.com/databricks/terraform-provider-databricks/pull/1842)).
 * Clarified [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job), `databricks_mws_*`, [databricks_cluster](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/cluster) docs ([#1796](https://github.com/databricks/terraform-provider-databricks/pull/1796)).
 
Updated dependency versions:

 * Bump github.com/golang-jwt/jwt/v4 from 4.4.2 to 4.4.3 ([#1817](https://github.com/databricks/terraform-provider-databricks/pull/1817)).
 * Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.24.0 to 2.24.1 ([#1761](https://github.com/databricks/terraform-provider-databricks/pull/1761)).
 * Bump golang.org/x/mod from 0.6.0 to 0.7.0 ([#1760](https://github.com/databricks/terraform-provider-databricks/pull/1760)).
 * Bump google.golang.org/api from 0.102.0 to 0.104.0 ([#1758](https://github.com/databricks/terraform-provider-databricks/pull/1758), [#1839](https://github.com/databricks/terraform-provider-databricks/pull/1839)).

### 1.6.5

 * Added `query_plan` attribute to `databricks_sql_visualization` ([#1752](https://github.com/databricks/terraform-provider-databricks/pull/1752)).
 * Fixed `member_id` reference of nested groups and [databricks_group_member](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/group_member) for exporter ([#1723](https://github.com/databricks/terraform-provider-databricks/pull/1723)).
 * Fixed auto-purged cluster behaviour for [databricks_library](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/library) ([#1745](https://github.com/databricks/terraform-provider-databricks/pull/1745)).
 * Use Jobs API 2.1 with name filter for [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/job) data source ([#1744](https://github.com/databricks/terraform-provider-databricks/pull/1744)).

### 1.6.4

 * Reverted `PRO` as default `warehouse_type` in [databricks_sql_endpoint](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_endpoint) ([#1741](https://github.com/databricks/terraform-provider-databricks/pull/1741)).
 * Moved testing of clusters, instance pools, and SQL Warehouses to internally configured test environment variables ([#1739](https://github.com/databricks/terraform-provider-databricks/pull/1739)).

Updated dependency versions:

 * Bump google.golang.org/api from 0.101.0 to 0.102.0 ([#1736](https://github.com/databricks/terraform-provider-databricks/pull/1736)).

### 1.6.3

 * Added `warehouse_type` parameter to [databricks_sql_endpoint](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_endpoint) to support PRO SKU ([#1728](https://github.com/databricks/terraform-provider-databricks/pull/1728)).
 * Correct exporting of the computed attributes for [databricks_cluster](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/cluster) ([#1711](https://github.com/databricks/terraform-provider-databricks/pull/1711)).
 * Escape database name in [databricks_sql_permissions](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_permissions) resource ([#1730](https://github.com/databricks/terraform-provider-databricks/pull/1730)).

## 1.6.2

 * Added `runtime_engine` to [databricks_cluster](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/cluster) ([#1686](https://github.com/databricks/terraform-provider-databricks/pull/1686)).
 * Added validation for `path` in [databricks_repo](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/repo) ([#1702](https://github.com/databricks/terraform-provider-databricks/pull/1702)).
 * Added auto-detection of AWS CodeCommit URLs in [databricks_repo](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/repo) ([#1704](https://github.com/databricks/terraform-provider-databricks/pull/1704)).
 * Restricting access to S3 bucket by custom tag on the IAM identity according to security email in [databricks_aws_bucket_policy](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/aws_bucket_policy) data resource ([#1694](https://github.com/databricks/terraform-provider-databricks/pull/1694)).
 * Update Azure Unity Catalog guide to use `azurerm_databricks_access_connector` ([#1685](https://github.com/databricks/terraform-provider-databricks/pull/1685)).
 * Clarify that [databricks_mws_permission_assignment](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/mws_permission_assignment) should be used for assigning account-level users/groups ([#1706](https://github.com/databricks/terraform-provider-databricks/pull/1706)).
 * Other documentation fixes ([#1696](https://github.com/databricks/terraform-provider-databricks/pull/1696), [#1692](https://github.com/databricks/terraform-provider-databricks/pull/1692)).

Updated dependency versions:

 * Bump github.com/stretchr/testify from 1.8.0 to 1.8.1 ([#1689](https://github.com/databricks/terraform-provider-databricks/pull/1689)).
 * Bump github.com/zclconf/go-cty from 1.11.1 to 1.12.0 ([#1714](https://github.com/databricks/terraform-provider-databricks/pull/1714)).
 * Bump golang.org/x/mod from 0.5.1 to 0.6.0 ([#1690](https://github.com/databricks/terraform-provider-databricks/pull/1690)).
 * Bump google.golang.org/api from 0.99.0 to 0.101.0 ([#1713](https://github.com/databricks/terraform-provider-databricks/pull/1713), [#1691](https://github.com/databricks/terraform-provider-databricks/pull/1691)).

## 1.6.1

 * Added `CAN_VIEW` permissions for [databricks_sql_dashboard](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_dashboard) and [databricks_sql_query](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_query) ([#1682](https://github.com/databricks/terraform-provider-databricks/pull/1682)).
 * Added `webhook_notifications` support to [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) ([#1674](https://github.com/databricks/terraform-provider-databricks/pull/1674)).
 * Allow updating `private_access_settings_id` for [databricks_mws_workspaces](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/mws_workspaces) ([#1668](https://github.com/databricks/terraform-provider-databricks/pull/1668)).
 * Changed [databricks_pipeline](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/pipeline) edition default value from `advanced` to `ADVANCED` ([#1683](https://github.com/databricks/terraform-provider-databricks/pull/1683)).
 * Fixed reference to `display_name_contains` in docs example ([#1677](https://github.com/databricks/terraform-provider-databricks/pull/1677)).

Updated dependency versions:

 * Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.23.0 to 2.24.0 ([#1671](https://github.com/databricks/terraform-provider-databricks/pull/1671)).
 * Bump github.com/zclconf/go-cty from 1.11.0 to 1.11.1 ([#1672](https://github.com/databricks/terraform-provider-databricks/pull/1672)).
 * Bump google.golang.org/api from 0.98.0 to 0.99.0 ([#1673](https://github.com/databricks/terraform-provider-databricks/pull/1673)).

## 1.6.0

 * Added [databricks_share](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/share) resource, [databricks_shares](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/shares) data source, and [databricks_share](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/share) data source ([#1664](https://github.com/databricks/terraform-provider-databricks/pull/1664)).
 * Documentation updates ([#1666](https://github.com/databricks/terraform-provider-databricks/pull/1666), [#1669](https://github.com/databricks/terraform-provider-databricks/pull/1669)).

## 1.5.0

 * Added `enable_serverless_compute` in the documentation for the [databricks_sql_global_config](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_global_config) ([#1655](https://github.com/databricks/terraform-provider-databricks/pull/1655)).
 * Update [databricks_grants](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/grants) with new permissions model ([#1657](https://github.com/databricks/terraform-provider-databricks/pull/1657)).
 * Exporter now adds the `force` flag to  users and groups ([#1661](https://github.com/databricks/terraform-provider-databricks/pull/1661)).

Updated dependency versions:

 * Bump google.golang.org/api from 0.97.0 to 0.98.0 ([#1652](https://github.com/databricks/terraform-provider-databricks/pull/1652)).

## 1.4.0

 * Added [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/job) data resource ([#1509](https://github.com/databricks/terraform-provider-databricks/pull/1509)).
 * Fixed handling of references in [databricks_permissions](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/permissions) resource exporter ([#1631](https://github.com/databricks/terraform-provider-databricks/pull/1631)).
 * Fixed stripping of `CAN_MANAGE` permission from caller of [databricks_sql_query](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_query) resource ([#1644](https://github.com/databricks/terraform-provider-databricks/pull/1644)).
 * Fixed provider debug mode ([#1560](https://github.com/databricks/terraform-provider-databricks/pull/1560)).

Documentation improvements:

 * Expanded documentation for [databricks_sql_visualization](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_visualization) resource ([#1630](https://github.com/databricks/terraform-provider-databricks/pull/1630)).

Updated dependency versions:

 * Bump github.com/hashicorp/hcl/v2 from 2.14.0 to 2.14.1 ([#1634](https://github.com/databricks/terraform-provider-databricks/pull/1634)).
 * Bump google.golang.org/api from 0.96.0 to 0.97.0 ([#1633](https://github.com/databricks/terraform-provider-databricks/pull/1633)).


## 1.3.1

 * Added autoscale `mode` configuration to [databricks_pipeline](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/pipeline) resource ([#1600](https://github.com/databricks/terraform-provider-databricks/pull/1600)).
 * Fixed `gcp_availability` field in [databricks_insance_pool](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/insance_pool) resource ([#1610](https://github.com/databricks/terraform-provider-databricks/pull/1610)).
 * Fixed `secret` field name in [databricks_mlflow_webhook](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/mlflow_webhook) resource ([#1625](https://github.com/databricks/terraform-provider-databricks/pull/1625)).
 * Further improvements of exporter ([#1602](https://github.com/databricks/terraform-provider-databricks/pull/1602)).

 Documentation improvements:
 
 * Updates for [databricks_table](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/table) and [databricks_mws_vpc_endpoint](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/mws_vpc_endpoint) ([#1616](https://github.com/databricks/terraform-provider-databricks/pull/1616)).
 * Document `http_timeout_seconds` from `common/client.go` ([#1599](https://github.com/databricks/terraform-provider-databricks/pull/1599)).
 * Expand `parameter` documentation for [databricks_sql_query](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_query) resource ([#1603](https://github.com/databricks/terraform-provider-databricks/pull/1603)).
 * Updated Unity Catalog docs ([#1605](https://github.com/databricks/terraform-provider-databricks/pull/1605), [#1626](https://github.com/databricks/terraform-provider-databricks/pull/1626), [#1620](https://github.com/databricks/terraform-provider-databricks/pull/1620)).

 Updated dependency versions:

 * Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.21.0 to 2.23.0 ([#1597](https://github.com/databricks/terraform-provider-databricks/pull/1597), [#1611](https://github.com/databricks/terraform-provider-databricks/pull/1611)).
 * Bump google.golang.org/api from 0.94.0 to 0.96.0 ([#1598](https://github.com/databricks/terraform-provider-databricks/pull/1598), [#1612](https://github.com/databricks/terraform-provider-databricks/pull/1612)).

## 1.3.0

 * Added `force_destroy` flag to [databricks_schema](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/schema) & [databricks_catalog](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/catalog) ([#1578](https://github.com/databricks/terraform-provider-databricks/pull/1578)).
 * Added [databricks_entitlements](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/entitlements) resource ([#1583](https://github.com/databricks/terraform-provider-databricks/pull/1583)).
 * Added [databricks_group_role](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/group_role) resource ([#1575](https://github.com/databricks/terraform-provider-databricks/pull/1575)).
 * Added [databricks_recipient](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/recipient) resource for Delta Sharing ([#1571](https://github.com/databricks/terraform-provider-databricks/pull/1571)).
 * Added `dbt_task` field to [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) ([#1537](https://github.com/databricks/terraform-provider-databricks/pull/1537)).
 * Fixed drift in `storage` for [databricks_pipeline](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/pipeline) ([#1574](https://github.com/databricks/terraform-provider-databricks/pull/1574)).
 * Fixed [databricks_mws_private_access_settings](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/mws_private_access_settings) defaults ([#1567](https://github.com/databricks/terraform-provider-databricks/pull/1567)).
 * Fixed [databricks_user](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/user) creation with `force` on account ([#1577](https://github.com/databricks/terraform-provider-databricks/pull/1577)).
 * Fixed [databricks_permissions](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/permissions) for calling user to `CAN_MANAGE` on [databricks_cluster](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/cluster) ([#1579](https://github.com/databricks/terraform-provider-databricks/pull/1579)).
 
 Documentation improvements:
 * Added `sql_task` configuration block in [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) resource ([#1589](https://github.com/databricks/terraform-provider-databricks/pull/1589)).
 * Added supported languages in Unity Shared clusters ([#1587](https://github.com/databricks/terraform-provider-databricks/pull/1587)).
 * Removed "Public Preview" marker from Unity Catalog resources docs ([#1570](https://github.com/databricks/terraform-provider-databricks/pull/1570)).
 * Fixed instance pool docs ([#1581](https://github.com/databricks/terraform-provider-databricks/pull/1581)).

Updated dependency versions:

 * Bump github.com/hashicorp/hcl/v2 from 2.13.0 to 2.14.0 ([#1582](https://github.com/databricks/terraform-provider-databricks/pull/1582)).
 * Bump google.golang.org/api from 0.93.0 to 0.94.0 ([#1569](https://github.com/databricks/terraform-provider-databricks/pull/1569)).

## 1.2.1

* Use resize API to scale `databricks_cluster` while it's running ([#1541](https://github.com/databricks/terraform-provider-databricks/pull/1541)).
* Updated to latest Unity Catalog privileges model ([#1556](https://github.com/databricks/terraform-provider-databricks/pull/1556)).
* Added cluster policy support for `databricks_pipeline` ([#1554](https://github.com/databricks/terraform-provider-databricks/pull/1554)).
* Fixed `databricks_node_type` by skipping nodes that aren't available in subscription/region ([#1534](https://github.com/databricks/terraform-provider-databricks/pull/1534)).
* Fixed sending of `active` flag in the `databricks_user` ([#1536](https://github.com/databricks/terraform-provider-databricks/pull/1536)).
* Fixed Azure CLI tests on Go 1.19 ([#1538](https://github.com/databricks/terraform-provider-databricks/pull/1538)).
* Various doc updates ([#1553](https://github.com/databricks/terraform-provider-databricks/pull/1553), [#1552](https://github.com/databricks/terraform-provider-databricks/pull/1552), [#1544](https://github.com/databricks/terraform-provider-databricks/pull/1544), [#1543](https://github.com/databricks/terraform-provider-databricks/pull/1543)).

Updated dependency versions:

* Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.20.0 to 2.21.0 ([#1540](https://github.com/databricks/terraform-provider-databricks/pull/1540)).
* Bump github.com/zclconf/go-cty from 1.10.0 to 1.11.0 ([#1558](https://github.com/databricks/terraform-provider-databricks/pull/1558)).
* Bump google.golang.org/api from 0.90.0 to 0.93.0 ([#1525](https://github.com/databricks/terraform-provider-databricks/pull/1525), [#1545](https://github.com/databricks/terraform-provider-databricks/pull/1545)).
* Bump gopkg.in/ini.v1 from 1.66.6 to 1.67.0 ([#1526](https://github.com/databricks/terraform-provider-databricks/pull/1526)).

## 1.2.0

 * Added [databricks_mws_permission_assignment](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/mws_permission_assignment) resource ([#1491](https://github.com/databricks/terraform-provider-databricks/pull/1491)).
 * Added [databricks_mws_workspaces](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/mws_workspaces) data resource ([#1497](https://github.com/databricks/terraform-provider-databricks/pull/1497)).
 * Added exporter for [databricks_pipeline](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/pipeline) and [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) ([#1484](https://github.com/databricks/terraform-provider-databricks/pull/1484)).
 * Fixed [databricks_permissions](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/permissions) for calling user to `CAN_MANAGE` on [databricks_mlflow_model](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/mlflow_model) ([#1507](https://github.com/databricks/terraform-provider-databricks/pull/1507)).
 * Fixed phantom `cluster` blocks in [databricks_pipeline](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/pipeline) ([#1508](https://github.com/databricks/terraform-provider-databricks/pull/1508)).
 * Improve test coverage of multitask [databricks_job](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job) and [databricks_pipeline](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/pipeline) ([#1493](https://github.com/databricks/terraform-provider-databricks/pull/1493)).
 * Minor stability improvements ([#1505](https://github.com/databricks/terraform-provider-databricks/pull/1505), [#1487](https://github.com/databricks/terraform-provider-databricks/pull/1487), [#1496](https://github.com/databricks/terraform-provider-databricks/pull/1496), [#1503](https://github.com/databricks/terraform-provider-databricks/pull/1503)).
 
Updated dependency versions:

 * Bump github.com/Azure/go-autorest/autorest from 0.11.27 to 0.11.28 ([#1513](https://github.com/databricks/terraform-provider-databricks/pull/1513)).
 * Bump github.com/Azure/go-autorest/autorest/adal from 0.9.20 to 0.9.21 ([#1512](https://github.com/databricks/terraform-provider-databricks/pull/1512)).
 * Bump github.com/Azure/go-autorest/autorest/azure/cli from 0.4.5 to 0.4.6 ([#1515](https://github.com/databricks/terraform-provider-databricks/pull/1515)).
 * Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.19.0 to 2.20.0 ([#1516](https://github.com/databricks/terraform-provider-databricks/pull/1516)).
 * Bump google.golang.org/api from 0.87.0 to 0.90.0 ([#1499](https://github.com/databricks/terraform-provider-databricks/pull/1499), [#1514](https://github.com/databricks/terraform-provider-databricks/pull/1514)).

## 1.1.0

* Added `databricks_sql_warehouses` data resource to list all warehouses in a workspace with a name filter ([#1460](https://github.com/databricks/terraform-provider-databricks/pull/1460)).
* Added `databricks_sql_warehouse` data resource to list SQL warehouse attributes based on single id ([#1460](https://github.com/databricks/terraform-provider-databricks/pull/1460)).
* Added `databricks_cluster` data resource to list cluster attributes based on single id ([#1460](https://github.com/databricks/terraform-provider-databricks/pull/1460)).
* Added Azure Managed Identity documentation examples ([#1471](https://github.com/databricks/terraform-provider-databricks/pull/1471)).
* Added more attributes to `databricks_cluster` ([#1459](https://github.com/databricks/terraform-provider-databricks/pull/1459)).
* Added more attributes to `databricks_instance_pool` ([#1463](https://github.com/databricks/terraform-provider-databricks/pull/1463)).
* Added feature request GitHub issue template ([#1482](https://github.com/databricks/terraform-provider-databricks/pull/1482)).
* Added `tf:optional` tag to simplify the code ([#1395](https://github.com/databricks/terraform-provider-databricks/pull/1395)).
* Fixed `databricks_pipeline` incorrect generation of `cluster` blocks ([#1416](https://github.com/databricks/terraform-provider-databricks/pull/1416)).
* Fixed `databricks_table` update for `column` block ([#1468](https://github.com/databricks/terraform-provider-databricks/pull/1468)).
* Fixed reads for `any_file` and `anonymous_function` in `databricks_sql_permissions` ([#1477](https://github.com/databricks/terraform-provider-databricks/pull/1477)).
* Tuned integration tests for `databricks_mws_*` ([#1483](https://github.com/databricks/terraform-provider-databricks/pull/1483)).
* Removed integration tests for `databricks_azure_adls_gen1_mount` ([#1461](https://github.com/databricks/terraform-provider-databricks/pull/1461)).

Updated dependency versions:

* Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.18.0 to 2.19.0 ([#1475](https://github.com/databricks/terraform-provider-databricks/pull/1475)).
* Bump google.golang.org/api from 0.86.0 to 0.87.0 ([#1476](https://github.com/databricks/terraform-provider-databricks/pull/1476)).

## 1.0.2

* Added `metastore` to `databricks_grants` ([#1447](https://github.com/databricks/terraform-provider-databricks/pull/1447)).
* Added update support for `databricks_mws_log_delivery` ([#1439](https://github.com/databricks/terraform-provider-databricks/pull/1439)).
* Fixed name generation for exported `databricks_notebook` ([#1435](https://github.com/databricks/terraform-provider-databricks/pull/1435)).
* Fixed `databricks_job` recreation on changed `docker_image` URL ([#1437](https://github.com/databricks/terraform-provider-databricks/pull/1437)).
* Fixed replace `databricks_mws_networks` on `vpc_endpoints` change ([#1453](https://github.com/databricks/terraform-provider-databricks/pull/1453)).
* Fixed diff suppress in `databricks_external_location` on `skip_validation` ([#1421](https://github.com/databricks/terraform-provider-databricks/pull/1421)).
* Switched to SQL warehouses API for `databricks_sql_endpoint` ([#1414](https://github.com/databricks/terraform-provider-databricks/pull/1414)).
* Improve docs for `databricks_permissions` and `databricks_obo_token` ([#1462](https://github.com/databricks/terraform-provider-databricks/pull/1462)).
* Improve docs for `abfs` block in `databricks_mount` ([#1446](https://github.com/databricks/terraform-provider-databricks/pull/1446)).
* Improve dev docs for `tf:"suppress_diff"`, `tf:"force_new"` and `tf:"sensitive"` tags ([#1465](https://github.com/databricks/terraform-provider-databricks/pull/1465)).
* Moved couple of acceptance tests out of preview ([#1433](https://github.com/databricks/terraform-provider-databricks/pull/1433)).
* Migrated to Go 1.18: `interface{}` -> `any` ([#1466](https://github.com/databricks/terraform-provider-databricks/pull/1466)).

Updated dependency versions:

* Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.17.0 to 2.18.0 ([#1450](https://github.com/databricks/terraform-provider-databricks/pull/1450)).

## 1.0.1

* Added better handling of missing objects during import ([#1417](https://github.com/databricks/terraform-provider-databricks/pull/1417)).
* Fix problem in the `databricks_node_type` data source ([#1400](https://github.com/databricks/terraform-provider-databricks/pull/1400)).
* Fix update for `databricks_storage_credential` resource ([#1403](https://github.com/databricks/terraform-provider-databricks/pull/1403)).
* Improve `databricks_job` resource ([#1428](https://github.com/databricks/terraform-provider-databricks/pull/1428)).
* Upgraded UC API from 2.0 to 2.1 ([#1418](https://github.com/databricks/terraform-provider-databricks/pull/1418)).
* Sanitized `host` provider field prior to auth flow ([#1385](https://github.com/databricks/terraform-provider-databricks/pull/1385)).
* Use `delta_sharing_scope` instead of `delta_sharing_enabled` ([#1398](https://github.com/databricks/terraform-provider-databricks/pull/1398)).
* Added instructions to create missing `.terraform.lock.hcl` ([#1397](https://github.com/databricks/terraform-provider-databricks/pull/1397)).
* Clarified `databricks_permissions` doc for service principals ([#1426](https://github.com/databricks/terraform-provider-databricks/pull/1426)).
* Clarified `databricks_cluster` `autotermination_minutes` default value ([#1419](https://github.com/databricks/terraform-provider-databricks/pull/1419)).
* Fully moved codebase from `databrickslabs` to `databricks` namespace ([#1429](https://github.com/databricks/terraform-provider-databricks/pull/1429)).
* Various integration testing improvements ([#1425](https://github.com/databricks/terraform-provider-databricks/pull/1425), [#1427](https://github.com/databricks/terraform-provider-databricks/pull/1427), [#1420](https://github.com/databricks/terraform-provider-databricks/pull/1420)).

Updated dependency versions:

* Bump Go from 1.16.x to 1.18.x ([#1413](https://github.com/databricks/terraform-provider-databricks/pull/1413)).
* Bump github.com/golang-jwt/jwt/v4 from 4.4.1 to 4.4.2 ([#1407](https://github.com/databricks/terraform-provider-databricks/pull/1407)).
* Bump github.com/hashicorp/hcl/v2 from 2.12.0 to 2.13.0 ([#1406](https://github.com/databricks/terraform-provider-databricks/pull/1406)).
* Bump github.com/stretchr/testify from 1.7.3 to 1.8.0 ([#1387](https://github.com/databricks/terraform-provider-databricks/pull/1387), [#1408](https://github.com/databricks/terraform-provider-databricks/pull/1408), [#1422](https://github.com/databricks/terraform-provider-databricks/pull/1422)).
* Bump google.golang.org/api from 0.84.0 to 0.86.0 ([#1386](https://github.com/databricks/terraform-provider-databricks/pull/1386), [#1423](https://github.com/databricks/terraform-provider-databricks/pull/1423)).

## 1.0.0

To make Databricks Terraform Provider generally available, we've moved it from [https://github.com/databrickslabs](https://github.com/databrickslabs) to [https://github.com/databricks](https://github.com/databricks). We've worked closely with the Terraform Registry team at Hashicorp to ensure a smooth migration. Existing terraform deployments continue to work as expected without any action from your side. We ask you to replace `databrickslabs/databricks` with `databricks/databricks` in all your `.tf` files. 

You should have .terraform.lock.hcl file in your state directory that is checked into source control. terraform init will give you the following warning.

```
Warning: Additional provider information from registry 

The remote registry returned warnings for registry.terraform.io/databrickslabs/databricks:
- For users on Terraform 0.13 or greater, this provider has moved to databricks/databricks. Please update your source in required_providers.
```

After you replace `databrickslabs/databricks` with `databricks/databricks` in the `required_providers` block, the warning will disappear. Do a global "search and replace" in `*.tf` files. Alternatively you can run `python3 -c "$(curl -Ls https://dbricks.co/updtfns)"` from the command-line, that would do all the boring work for you.

If you didn't check-in [`.terraform.lock.hcl`](https://www.terraform.io/language/files/dependency-lock#lock-file-location) to the source code version control, you may you may see `Failed to install provider` error. Please follow the simple steps described in the [troubleshooting guide](https://github.com/databrickslabs/terraform-provider-databricks/blob/master/docs/guides/troubleshooting.md).

## 0.6.2

* Add a warning in `databricks_permissions` token usage docs ([#1380](https://github.com/databricks/terraform-provider-databricks/pull/1380)).

Updated dependency versions:

* Bump google.golang.org/api from 0.83.0 to 0.84.0
* Bump github.com/stretchr/testify from 1.7.2 to 1.7.3

## 0.6.1

* Added `databricks_service_principal` and `databricks_service_principals` data resources ([#1370](https://github.com/databricks/terraform-provider-databricks/pull/1370)).
* Updated `databricks_pipeline` resource to match latest APIs ([#1368](https://github.com/databricks/terraform-provider-databricks/pull/1368)).
* Made `gcp_managed_network_config` optional in `databricks_mws_workspaces` ([#1365](https://github.com/databricks/terraform-provider-databricks/pull/1365)).
* Enforced consistent naming for resource files ([#1366](https://github.com/databricks/terraform-provider-databricks/pull/1366), [#1369](https://github.com/databricks/terraform-provider-databricks/pull/1369)).
* Updated resources diagram ([#1373](https://github.com/databricks/terraform-provider-databricks/pull/1373)).

Updated dependency versions:

* Bump google.golang.org/api from 0.82.0 to 0.83.0

## 0.6.0

* Added `delta_sharing_*` support to `databricks_metastore` ([#1334](https://github.com/databricks/terraform-provider-databricks/pull/1334)).
* Added `databricks_git_credentials` pat discovery from common environment variables ([#1353](https://github.com/databricks/terraform-provider-databricks/pull/1353)).
* Added `databricks_permissions` for `databricks_pipeline` ([#1361](https://github.com/databricks/terraform-provider-databricks/pull/1361)).
* Added `network_id` to `network` block in `databricks_mws_workspaces` for GCP ([#1360](https://github.com/databricks/terraform-provider-databricks/pull/1360)).
* Added `azure_managed_identity` block to `databricks_storage_credential` and `databricks_metastore_data_access` resources ([#1354](https://github.com/databricks/terraform-provider-databricks/pull/1354)).
* Update docs regarding importing of `databricks_sql_*` resources ([#1349](https://github.com/databricks/terraform-provider-databricks/pull/1349)).
* Apply ownership for UC objects during creation ([#1338](https://github.com/databricks/terraform-provider-databricks/pull/1338)).
* Re-create purged cluster for `databricks_mount` for Google Storage ([#1333](https://github.com/databricks/terraform-provider-databricks/pull/1333)).
* Various documentation fixes ([#1350](https://github.com/databricks/terraform-provider-databricks/pull/1350)).

Updated dependency versions:

* Bump google.golang.org/api from 0.80.0 to 0.81.0
* Bump gopkg.in/ini.v1 from 1.66.4 to 1.66.6
* Bump google.golang.org/api from 0.81.0 to 0.82.0
* Bump github.com/stretchr/testify from 1.7.1 to 1.7.2
* Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.16.0 to 2.17.0

## 0.5.9

* Added warning section for debug mode ([#1325](https://github.com/databricks/terraform-provider-databricks/pull/1325)).
* Added ability to specify tags for `databricks_job` ([#1337](https://github.com/databricks/terraform-provider-databricks/pull/1337)).
* Upgraded AWS provider for AWS guides. Added examples for account-level identities ([#1332](https://github.com/databricks/terraform-provider-databricks/pull/1332)).
* Updated docs to use `application_id` as privilege for `databricks_service_principal` ([#1336](https://github.com/databricks/terraform-provider-databricks/pull/1336)).
* Added `databricks_service_principal_role` resource ([#1340](https://github.com/databricks/terraform-provider-databricks/pull/1340)).
* Fixed itegration testing image ([#1342](https://github.com/databricks/terraform-provider-databricks/pull/1342), [#1343](https://github.com/databricks/terraform-provider-databricks/pull/1343)).
* Added `skip_validation` for `databricks_external_location` ([#1330](https://github.com/databricks/terraform-provider-databricks/pull/1330)).
* Added `alert_on_last_attempt` to `databricks_job` ([#1341](https://github.com/databricks/terraform-provider-databricks/pull/1341)).
* Skip `make test` on doc-only changes ([#1339](https://github.com/databricks/terraform-provider-databricks/pull/1339)).
* Improve common package test coverage ([#1344](https://github.com/databricks/terraform-provider-databricks/pull/1344)).
* Re-create purged cluster for `databricks_mount` for AWS S3 ([#1345](https://github.com/databricks/terraform-provider-databricks/pull/1345)).

Updated dependency versions:

* Bump google.golang.org/api from 0.79.0 to 0.80.0
* Bump github.com/Azure/go-autorest/autorest/adal from 0.9.19 to 0.9.20

## 0.5.8

* Update `aws_iam_policy_document` in `databricks_mws_customer_managed_keys` docs to restrict KMS policy to caller AWS account ([#1309](https://github.com/databricks/terraform-provider-databricks/pull/1309)).
* Added `gcs` destination to `init_scripts` in `databricks_cluster` ([#1308](https://github.com/databricks/terraform-provider-databricks/pull/1308)).
* Clarify optionality of `databricks_mws_workspaces`.`deployment_name` in docs and examples ([#1315](https://github.com/databricks/terraform-provider-databricks/pull/1315)).
* Update `databricks_mws_log_delivery` docs ([#1320](https://github.com/databricks/terraform-provider-databricks/pull/1320)).
* Fix updating `databricks_service_principal` on Azure ([#1322](https://github.com/databricks/terraform-provider-databricks/pull/1322)).
* Added `tf:suppress_diff` on `artifact_location` for `databricks_mlflow_experiment` ([#1323](https://github.com/databricks/terraform-provider-databricks/pull/1323)).

Updated dependency versions:

* Bump github.com/Azure/go-autorest/autorest/adal from 0.9.18 to 0.9.19
* Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.14.0 to 2.16.0
* Bump google.golang.org/api from 0.77.0 to 0.79.0

## 0.5.7

* Added `external_id` and `force` attributes to `databricks_service_principal` resource ([#1293](https://github.com/databricks/terraform-provider-databricks/pull/1293)).
* Added documentation for `databicks_git_credential` resource ([#1295](https://github.com/databricks/terraform-provider-databricks/pull/1295)).
* Added documentation for `git_source` in `databricks_job` ([#1297](https://github.com/databricks/terraform-provider-databricks/pull/1297)).
* Fix `job_cluster`.`num_workers` in `databricks_job` ([#1284](https://github.com/databricks/terraform-provider-databricks/pull/1284)).
* Various documentation improvements ([#1292](https://github.com/databricks/terraform-provider-databricks/pull/1292)), ([#1296](https://github.com/databricks/terraform-provider-databricks/pull/1296)), ([#1298](https://github.com/databricks/terraform-provider-databricks/pull/1298)).

Updated dependency versions:

* Bump google.golang.org/api from 0.75.0 to 0.77.0
* Removed github.com/pkg/errors dependency

## 0.5.6

* Added `databricks_views` data resource, making `databricks_tables` return only managed or external tables in Unity Catalog ([#1274](https://github.com/databricks/terraform-provider-databricks/issues/1274)).
* Added default timeout of 20m to `databricks_mount` ([#1280](https://github.com/databricks/terraform-provider-databricks/pull/1280)).
* Made `common.DataResource` deterministic ([#1279](https://github.com/databricks/terraform-provider-databricks/pull/1279)).
* Fixed exporting text-only widgets ([#1278](https://github.com/databricks/terraform-provider-databricks/pull/1278)).
* Updated devcontainer to support ARM ([#1256](https://github.com/databricks/terraform-provider-databricks/pull/1256)).
* Various documentation fixes ([#1285](https://github.com/databricks/terraform-provider-databricks/pull/1285), [#1282](https://github.com/databricks/terraform-provider-databricks/pull/1282), [#1281](https://github.com/databricks/terraform-provider-databricks/pull/1281), [#1276](https://github.com/databricks/terraform-provider-databricks/pull/1276)).

Updated dependency versions:

* Bump github.com/hashicorp/hcl/v2 from 2.11.1 to 2.12.0
* Bump github.com/Azure/go-autorest/autorest from 0.11.26 to 0.11.27

## 0.5.5

* Added configuration generators for `databricks_sql_*` resources in _experimental_ [Resource Exporter](https://asciinema.org/a/Rv8ZFJQpfrfp6ggWddjtyXaOy) ([#1199](https://github.com/databricks/terraform-provider-databricks/pull/1199)).
* Added `google_credentials` provider agument that has the same semantics as [`credentials` argument](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#credentials) in official [`google` provider](https://registry.terraform.io/providers/hashicorp/google/latest/docs) ([#1214](https://github.com/databricks/terraform-provider-databricks/pull/1214)).
* Fixed `databricks_grants` on UC external location empty list error ([#1202](https://github.com/databricks/terraform-provider-databricks/issues/1202)).
* Fixed errors in `databricks_permissions` resource for auto-purged `databricks_cluster` ([#1252](https://github.com/databricks/terraform-provider-databricks/commit/dac42524f5037c796187c77ba49367b964b03e9f)). 
* Various documentation fixes ([#1231](https://github.com/databricks/terraform-provider-databricks/pull/1231), [#1239](https://github.com/databricks/terraform-provider-databricks/pull/1239), [#1254](https://github.com/databricks/terraform-provider-databricks/pull/1254), [#1240](https://github.com/databricks/terraform-provider-databricks/commit/2dabfc90592d79249bd177bb975a84e0b98504f7)).

Updated dependency versions:

* Bump google.golang.org/api from 0.71.0 to 0.75.0
* Bump github.com/golang-jwt/jwt/v4 from 4.3.0 to 4.4.1
* Bump github.com/stretchr/testify from 1.7.0 to 1.7.1
* Bump github.com/hashicorp/go-retryablehttp from 0.7.0 to 0.7.1
* Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.11.0 to 2.14.0
* Bump github.com/Azure/go-autorest/autorest from 0.11.24 to 0.11.26

## 0.5.4

* Completely removed custom client-side validation in `databricks_service_principal` ([#1193](https://github.com/databricks/terraform-provider-databricks/issues/1193)).
* Added export functionality for Databricks SQL objects - endpoints, queries, dashboards, widgets, visualizations ([#1199](https://github.com/databricks/terraform-provider-databricks/pull/1199)).

Updated dependency versions:

* Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.10.1 to 2.11.0

## 0.5.3

* Failures in [exporter](https://asciinema.org/a/Rv8ZFJQpfrfp6ggWddjtyXaOy) resource listing no longer halt the entire command run ([#1166](https://github.com/databricks/terraform-provider-databricks/issues/1166)).
* Removed client-side validation in `databricks_service_principal` for `application_id`, that may not always be available in the planning stage ([#1165](https://github.com/databricks/terraform-provider-databricks/issues/1165)).
* Use correct HTTP verb for modifying `databricks_permissions` on `databricks_sql_endpoint` entities. Authorized user, assumingly part of `admins` group, is no longer sending `CAN_MANAGE` permission in the HTTP PUT request ([#1163](https://github.com/databricks/terraform-provider-databricks/issues/1163)).
* Added diff suppression for `min_num_clusters` field in `databricks_sql_endpoint` ([#1172](https://github.com/databricks/terraform-provider-databricks/pull/1172)).
* Added special case for handling `Cannot access cluster that was terminated or unpinned more than 30 days ago` error in `databricks_cluster` as an indication of resource removed on the platform side ([#1177](https://github.com/databricks/terraform-provider-databricks/issues/1177)).
* Fixed updating of `databricks_table` resources ([#1175](https://github.com/databricks/terraform-provider-databricks/issues/1175)).
* Fixed configuration drift in `databricks_grant` by reading existing permissions and removing them in subsequent update calls ([#1164](https://github.com/databricks/terraform-provider-databricks/issues/1164)).

Updated dependency versions:

* Bump google.golang.org/api from 0.70.0 to 0.71.0

## 0.5.2

* Added `databricks_catalogs`, `databricks_schemas`, and `databricks_tables` data resources ([#1155](https://github.com/databricks/terraform-provider-databricks/pull/1155)).
* Fixed `databricks_metastore_assignment` configuration drift by properly deleting metastore assignment and detecting manual changes from Account console. This also means that de-assigned metastore from a workspace would mark it as remotely removed. Manual assignment of different metastore would also trigger resource updates ([#1146](https://github.com/databricks/terraform-provider-databricks/issues/1146)).
* Fixed `databricks_table` creation in managed mode ([#1151](https://github.com/databricks/terraform-provider-databricks/issues/1151)).
* Fixed `databricks_sql_endpoint` timeout propagation ([#1142](https://github.com/databricks/terraform-provider-databricks/issues/1142)).
* Multiple documentation fixes.

Updated dependency versions:

* Bump google.golang.org/api from 0.69.0 to 0.70.0

## 0.5.1

* Added an extended documentation from provisioning AWS PrivateLink workspace ([#1084](https://github.com/databricks/terraform-provider-databricks/pull/1084)).
* Added `databricks_jobs` data resource to get a map of all job names and their ids ([#1138](https://github.com/databricks/terraform-provider-databricks/pull/1138)).

Updated dependency versions:

* Bump google.golang.org/api from 0.68.0 to 0.69.0

## 0.5.0

* Added `workspace_url` attribute to the `databricks_current_user` data source ([#1107](https://github.com/databricks/terraform-provider-databricks/pull/1107)).
* Fixed issue at `databricks_mount` where new cluster was created for S3 mount even when `cluster_id` was specified ([#1064](https://github.com/databricks/terraform-provider-databricks/issues/1064)).
* Allow to disable auto-termination for Databricks SQL endpoints ([#900](https://github.com/databricks/terraform-provider-databricks/pull/900)).
* Added new `gcp_attributes` to `databricks_cluster` and `databricks_instance_pool` ([#1126](https://github.com/databricks/terraform-provider-databricks/pull/1126)).
* Added exporter functionality for `databricks_ip_access_list` and `databricks_workspace_conf` ([#1125](https://github.com/databricks/terraform-provider-databricks/pull/1125)).
* Added `graviton` selector for `databricks_node_type` and `databricks_spark_version` data sources ([#1127](https://github.com/databricks/terraform-provider-databricks/pull/1127)).
* Added interactive mode to resource exporter ([#1010](https://github.com/databricks/terraform-provider-databricks/pull/1010)).
* Added preview support for `git_source` in `databricks_job` ([#1090](https://github.com/databricks/terraform-provider-databricks/pull/1090)).
* Multiple other fixes and documentation improvements.

Updated dependency versions:

* Bump github.com/golang-jwt/jwt/v4 from 4.2.0 to 4.3.0
* Bump google.golang.org/api from 0.67.0 to 0.68.0
* Bump gopkg.in/ini.v1 from 1.66.3 to 1.66.4

## 0.4.9

* Prevent creation of `databricks_group` with `users` and `admins` reserved names ([#1089](https://github.com/databricks/terraform-provider-databricks/issues/1089)).
* Added support for shared clusters in multi-task `databricks_job` ([#1082](https://github.com/databricks/terraform-provider-databricks/issues/1082)).
* Added diff suppression for `external_id` in `databricks_group` ([#1099](https://github.com/databricks/terraform-provider-databricks/issues/1099)).
* Added diff suppression for `external_id` in `databricks_user` ([#1097](https://github.com/databricks/terraform-provider-databricks/issues/1097)).
* Added `users`, `service_principals`, and `child_groups` exported properties to `databricks_group` data resource ([#1085](https://github.com/databricks/terraform-provider-databricks/issues/1085)).
* Added various documentation improvements.

**Deprecations**

* `databricks_group`.`members` is deprecated in favor of `users`, `service_principals`, and `child_groups` exported properties. Please do slight modifications of your configuration.

Updated dependency versions:

* Bump google.golang.org/api from 0.66.0 to 0.67.0

## 0.4.8

* Added support for `tf:suppress_diff` on primitive types ([#984](https://github.com/databricks/terraform-provider-databricks/issues/984)).
* Fixed issue arises when destroying `databricks_sql_global_config` with instance profile set ([#1076](https://github.com/databricks/terraform-provider-databricks/issues/1076)).
* Added setting of SQL configuration parameters in `databricks_sql_global_config` ([#1080](https://github.com/databricks/terraform-provider-databricks/pull/1080)).
* Added support for release channels in `databricks_sql_endpoint` configuration ([#1078](https://github.com/databricks/terraform-provider-databricks/pull/1078)).
* Added documentation for `databricks_mlflow_webhook` resource ([#1086](https://github.com/databricks/terraform-provider-databricks/pull/1086)).

Updated dependency versions:

* Bump google.golang.org/api from 0.65.0 to 0.66.0

## 0.4.7
* Added optional `force` argument to `databricks_group` resource to ignore `cannot create group: Group with name X already exists.` errors and implicitly import the specific group into Terraform state, enforcing entitlements defined in the instance of resource ([#1066](https://github.com/databricks/terraform-provider-databricks/pull/1066)).
* Added support to configure permissions for all MLflow models ([#1044](https://github.com/databricks/terraform-provider-databricks/issues/1044)).
* Fixed `databricks_service_principal` `display_name` update ([#1065](https://github.com/databricks/terraform-provider-databricks/issues/1065)).
* Added documentation for Unity Catalog resources.

Updated dependency versions:

* Bump gopkg.in/ini.v1 from 1.66.2 to 1.66.3

## 0.4.6

* Clarified error messages around `azure_workspace_resource_id` provider configuration ([#1049](https://github.com/databricks/terraform-provider-databricks/issues/1049)).
* Added optional `force` argument to `databricks_user` resource to ignore `cannot create user: User with username X already exists` errors and implicitly import the specific user into Terraform state, enforcing entitlements defined in the instance of resource ([#1048](https://github.com/databricks/terraform-provider-databricks/pull/1048)).
* Added `databricks_user_role` resource, that can assign roles on Databricks Account or `databricks_instance_profile` for data access ([#1047](https://github.com/databricks/terraform-provider-databricks/pull/1047)).

**Deprecations**

* `databricks_user_instance_profile` is deprecated in favor of `databricks_user_role`. Please do slight modifications of your configuration.

Updated dependency versions:

* Bump github.com/Azure/go-autorest/autorest/azure/auth from 0.5.10 to 0.5.11
* Bump github.com/Azure/go-autorest/autorest/azure/cli from 0.4.3 to 0.4.5
* Bump github.com/Azure/go-autorest/autorest from 0.11.23 to 0.11.24

## 0.4.5

* Cross-linked resource documentation ([#1027](https://github.com/databricks/terraform-provider-databricks/pull/1027)).
* Added Azure example for sql_global_config ([#1028](https://github.com/databricks/terraform-provider-databricks/pull/1028)).

Updated dependency versions:

* Bump google.golang.org/api from 0.63.0 to 0.65.0

## 0.4.4

* Added support for [running provider in a debug mode](https://www.terraform.io/plugin/sdkv2/debugging#running-terraform-with-a-provider-in-debug-mode) from Visual Studio Code through `Debug Provider` run configuration in order to troubleshoot more complicated issues.
* Allowed managing of libraries on `databricks_cluster` outside of Terraform state for resources without any `library` configuration blocks. This should simplify PaaS-like CI/CD workflows ([#1024](https://github.com/databricks/terraform-provider-databricks/pull/1024)).
* Added experimental resources.

**Behavior changes**

* Whenever library is installed on `databricks_cluster` without any [`library` configuration blocks](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/cluster#library-configuration-block), it won't be removed anymore.

## 0.4.3

* Added support for `databricks_permissions` for `databricks_mlflow_experiment` and `databricks_mlflow_model` ([#1013](https://github.com/databricks/terraform-provider-databricks/pull/1013)).
* Added `Using XXX auth` explanation to HTTP 403 errors, which should help troubleshooting misconfigured authentication or provider aliasing. Example error message now looks like: *cannot create group: /2.0/preview/scim/v2/Groups is only accessible by admins. Using databricks-cli auth: host=https://XXX.cloud.databricks.com/, token=`***REDACTED***`, profile=demo.* All sensitive configuration parameters (`token`, `password`, and `azure_client_secret`) are redacted and replaced with `***REDACTED***` ([#821](https://github.com/databricks/terraform-provider-databricks/issues/821)).
* Improved documentation with regards to public subnets in AWS quick start ([#1005](https://github.com/databricks/terraform-provider-databricks/pull/1005)).
* Added `databricks_mount` code genration for [exporter](https://registry.terraform.io/providers/databricks/databricks/latest/docs/guides/experimental-exporter) tooling ([#1006](https://github.com/databricks/terraform-provider-databricks/pull/1006)).
* Increase dependency check frequency ([#1007](https://github.com/databricks/terraform-provider-databricks/pull/1007)).
* Added experimental resources.

## 0.4.2

* Added optional `auth_type` provider conf to enforce specific auth type to be used in very rare cases, where a single Terraform state manages Databricks workspaces on more than one cloud and `More than one authorization method configured` error is a false positive. Valid values are `pat`, `basic`, `azure-client-secret`, `azure-msi`, `azure-cli`, and `databricks-cli` ([#1000](https://github.com/databricks/terraform-provider-databricks/pull/1000)).
* Added `DBC` format support for `databricks_notebook` ([#989](https://github.com/databricks/terraform-provider-databricks/pull/989)). 
* Fixed creating new `databricks_mws_workspaces` with `token {}` block ([#994](https://github.com/databricks/terraform-provider-databricks/issues/994)).
* Added automated documentation formatting with `make fmt-docs`, so that all HCL examples look consistent ([#999](https://github.com/databricks/terraform-provider-databricks/pull/999)).
* Increased codebase unit test coverage to 91% to improve stability ([#996](https://github.com/databricks/terraform-provider-databricks/pull/996), [#992](https://github.com/databricks/terraform-provider-databricks/pull/992), [#991](https://github.com/databricks/terraform-provider-databricks/pull/991), [#990](https://github.com/databricks/terraform-provider-databricks/pull/990)).

Updated dependency versions:

* Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.10.0 to 2.10.1 

## 0.4.1

* Added `databricks_library` resource to install library on `databricks_cluster` ([#904](https://github.com/databricks/terraform-provider-databricks/pull/904)).
* Added `databricks_clusters` data resource to list all clusters in the workspace, which might be used to install `databricks_library` on all clusters ([#955](https://github.com/databricks/terraform-provider-databricks/pull/955)).
* Fixed refresh of `library` blocks on a stopped `databricks_cluster` ([#952](https://github.com/databricks/terraform-provider-databricks/issues/952)).
* Whenever a library fails to get installed on a running `databricks_cluster`, we now automatically remove this library, so that the clean state of managed libraries is properly maintained. Without this fix users had to manually go to Clusters UI and remove library from a cluster, where it failed to install. Libraries add up to CREATE and UPDATE timeouts of `databricks_cluster` resource. ([#599](https://github.com/databricks/terraform-provider-databricks/issues/599)).
* Added `token` block to `databricks_mws_workspaces` to avoid unnecessary provider aliasing ([#957](https://github.com/databricks/terraform-provider-databricks/issues/957)).
* Fixed disabling `databricks_global_init_script` ([#958](https://github.com/databricks/terraform-provider-databricks/issues/958)).
* Fixed configuration drift issues with `aws_attributes`, `azure_attributes`, `gcp_attributes`, and `email_notifications` configuration blocks in `databricks_cluster`, `databricks_job`, and `databricks_instance_pool` resources ([#981](https://github.com/databricks/terraform-provider-databricks/pull/981)).
* Improved Databricks CLI auth by eagerly resolving `host`, `username`, `password`, and `token` from the specified `profile`. Added explicit logging of auth parameters in debug logs ([#965](https://github.com/databricks/terraform-provider-databricks/pull/965)).
* TLS timeouts, which may occur during Azure MSI auth, are no longer failing API requests and retried within a normal policy ([#966](https://github.com/databricks/terraform-provider-databricks/pull/966)).
* `debug_headers` provider conf is also logging the `Host` header to help troubleshooting auth issues ([#964](https://github.com/databricks/terraform-provider-databricks/pull/964)).
* Added new experimental resources and increased test coverage.

Updated dependency versions:

* Bump github.com/golang-jwt/jwt/v4 from 4.1.0 to 4.2.0
* Bump google.golang.org/api from 0.60.0 to 0.63.0
* Bump github.com/Azure/go-autorest/autorest from 0.11.22 to 0.11.23
* Bump github.com/Azure/go-autorest/autorest/azure/auth from 0.5.9 to 0.5.10
* Bump gopkg.in/ini.v1 from 1.66.0 to 1.66.2
* Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.9.0 to 2.10.0

## 0.4.0

* Added `databricks_mlflow_model` and `databricks_mlflow_experiment` resources ([#931](https://github.com/databricks/terraform-provider-databricks/pull/931)) 
* Added support for `repo_path` to `databricks_permissions` resource ([#875](https://github.com/databricks/terraform-provider-databricks/issues/875)).
* Added `external_id` to `databricks_user` and `databricks_group` ([#927](https://github.com/databricks/terraform-provider-databricks/pull/927)).
* Fixed `databricks_repo` creation corner cases on MS Windows OS ([#911](https://github.com/databricks/terraform-provider-databricks/issues/911)).
* Fixed configuration drift for `databricks_cluster`.`aws_attributes`.`zone_id` with `auto`, which resulted in unwanted cluster restarts ([#937](https://github.com/databricks/terraform-provider-databricks/pull/937)).
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

* Added `databricks_sql_global_config` resource to provide global configuration for SQL Endpoints ([#855](https://github.com/databricks/terraform-provider-databricks/issues/855))
* Added `databricks_mount` resource to mount arbitrary cloud storage ([#497](https://github.com/databricks/terraform-provider-databricks/issues/497))
* Improved implementation of `databricks_repo` by creating the parent folder structure ([#895](https://github.com/databricks/terraform-provider-databricks/pull/895))
* Fixed `databricks_job` error related [to randomized job IDs](https://docs.databricks.com/release-notes/product/2021/august.html#jobs-service-stability-and-scalability-improvements) ([#901](https://github.com/databricks/terraform-provider-databricks/issues/901))
* Replace `databricks_group` on name change ([#890](https://github.com/databricks/terraform-provider-databricks/pull/890))
* Names of scopes in `databricks_secret_scope` can have `/` characters in them ([#892](https://github.com/databricks/terraform-provider-databricks/pull/892))

**Deprecations**
* `databricks_aws_s3_mount`, `databricks_azure_adls_gen1_mount`, `databricks_azure_adls_gen2_mount`, and `databricks_azure_blob_mount` are deprecated in favor of `databricks_mount`.

Updated dependency versions:

* Bump google.golang.org/api from 0.59.0 to 0.60.0

## 0.3.10

* Added `private_access_level` and `allowed_vpc_endpoint_ids` to `databricks_mws_private_access_settings` resource, which is also now updatable ([#867](https://github.com/databricks/terraform-provider-databricks/issues/867)).
* Fixed missing diff skip for `skip_validation` in `databricks_instance_profile` ([#860](https://github.com/databricks/terraform-provider-databricks/issues/860)).
* Added support for `pipeline_task` ([871](https://github.com/databricks/terraform-provider-databricks/pull/871)) and `python_wheel_task` ([#872](https://github.com/databricks/terraform-provider-databricks/pull/872)) to `databricks_job`.
* Improved enterprise HTTPS proxy support for creating workspaces in PrivateLink environments ([#882](https://github.com/databricks/terraform-provider-databricks/pull/882)).
* Added `hostname` attribute to `odbc_params` in `databricks_sql_endpoint` ([#868](https://github.com/databricks/terraform-provider-databricks/issues/868)).
* Improved documentation ([#858](https://github.com/databricks/terraform-provider-databricks/pull/858), [#870](https://github.com/databricks/terraform-provider-databricks/pull/870)).

Updated dependency versions:

* Bumped google.golang.org/api from 0.58.0 to 0.59.0

## 0.3.9

* Added initial support for multiple task orchestration in `databricks_job` [#853](https://github.com/databricks/terraform-provider-databricks/pull/853)
* Fixed provider crash for new terraform states related to bug [#813](https://github.com/hashicorp/terraform-plugin-sdk/issues/813) in Terraform SDK v2.8.0 ([#854](https://github.com/databricks/terraform-provider-databricks/issues/854))
* Re-added `skip_validation` to `databricks_instance_profile` resource [#762](https://github.com/databricks/terraform-provider-databricks/issues/762)
* Removed direct dependency on `aws-sdk-go`.

Updated dependency versions:

* Reverted github.com/hashicorp/terraform-plugin-sdk/v2 from 2.8.0 to 2.7.0

## 0.3.8

* Added `databricks_repo` resource to manage [Databricks Repos](https://docs.databricks.com/repos.html) ([#771](https://github.com/databricks/terraform-provider-databricks/pull/771))
* Added support for Azure MSI authentication ([#743](https://github.com/databricks/terraform-provider-databricks/pull/743))
* Added support to create `databricks_user` on the account level ([#818](https://github.com/databricks/terraform-provider-databricks/issues/818))
* Already deleted `databricks_token` don't fail the apply ([#808](https://github.com/databricks/terraform-provider-databricks/pull/808))
* Default `terraform-mount` clusters created for mounting for `databricks_aws_s3_mount`, `databricks_azure_adls_gen1_mount`, `databricks_azure_adls_gen2_mount`, and `databricks_azure_blob_mount` have now `spark.scheduler.mode` as `FIFO` ([#828](https://github.com/databricks/terraform-provider-databricks/pull/828))
* Fixed crash when using non-Azure authentication to mount Azure resources ([#831](https://github.com/databricks/terraform-provider-databricks/issues/831))
* Fixed replacement of `instance_pool_id` in `databricks_cluster`, when `driver_instance_pool_id` was not explicitly specified ([#824](https://github.com/databricks/terraform-provider-databricks/issues/824))
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

* Added `databricks_obo_token` resource to create On-Behalf-Of tokens for a Service Principal in Databricks workspaces on AWS. It is very useful, when you want to provision resources within a workspace through narrowly-scoped service principal, that has no access to other workspaces within the same Databricks Account ([#736](https://github.com/databricks/terraform-provider-databricks/pull/736))
* Added support for [IAM credential passthrough](https://docs.databricks.com/security/credential-passthrough/iam-passthrough.html) with `is_meta_instance_profile` property for `databricks_instance_profile` ([#745](https://github.com/databricks/terraform-provider-databricks/pull/745))
* Fixed incorrect workspace update bug and added more validation error messaging ([#649](https://github.com/databricks/terraform-provider-databricks/pull/649))
* Clarify network modification procedure on active workspaces ([#732](https://github.com/databricks/terraform-provider-databricks/issues/732))
* Updated AWS IAM policy templates version to `2012-10-17` (`databricks_aws_bucket_policy`, `databricks_aws_assume_role_policy`, and `databricks_aws_crossaccount_policy`) ([#688](https://github.com/databricks/terraform-provider-databricks/issues/688))
* Various bug fixes in Databricks SQL resources

Updated dependency versions:

* Bump github.com/aws/aws-sdk-go to v1.40.12
* Bump github.com/hashicorp/hcl/v2 to v2.10.1
* Bump github.com/zclconf/go-cty to v1.9.0
* Bump golang.org/x/time to v0.0.0-20210723032227-1f47c861a9ac
* Bump golang.org/x/tools to v0.1.5

## 0.3.6

* Added support for hybrid pools ([#689](https://github.com/databricks/terraform-provider-databricks/pull/689))
* Added support for `always_running` jobs, which are restarted on resource updates ([#715](https://github.com/databricks/terraform-provider-databricks/pull/715))
* Azure CLI auth is now forcing JSON output ([#717](https://github.com/databricks/terraform-provider-databricks/pull/717))
* `databricks_permissions` are getting validation on `terraform plan` stage ([#706](https://github.com/databricks/terraform-provider-databricks/pull/706))
* Added `databricks_directory` resource ([#690](https://github.com/databricks/terraform-provider-databricks/pull/690))
* Added `run_as_role` field to `databricks_sql_query` ([#684](https://github.com/databricks/terraform-provider-databricks/pull/684))
* Added `user_id` attribute for `databricks_user` data resource, so that it's possible to dynamically create resources based on members of the group ([#714](https://github.com/databricks/terraform-provider-databricks/pull/714))
* Added more selectors to `databricks_node_type` data source ([#723](https://github.com/databricks/terraform-provider-databricks/pull/723))
* Azure auth with SPN now uses AAD token by default instead of PAT. Previous behavior (using PAT) could be restored by setting `azure_use_pat_for_spn` to `true` ([#721](https://github.com/databricks/terraform-provider-databricks/pull/721))
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

* Fixed setting of permissions for SQLA endpoints ([#661](https://github.com/databricks/terraform-provider-databricks/issues/661))
* Added support for preloading of Docker images into instance pools ([#663](https://github.com/databricks/terraform-provider-databricks/issues/663))
* Added the `databricks_user` data source ([#648](https://github.com/databricks/terraform-provider-databricks/pull/648))
* Fixed support for `spot_instance_policy` in SQLA Endpoints ([#665](https://github.com/databricks/terraform-provider-databricks/issues/665))
* Added documentation for `databricks_pipeline` resource ([#673](https://github.com/databricks/terraform-provider-databricks/pull/673))
* Fixed mapping for `databricks_service_principal` on AWS ([#656](https://github.com/databricks/terraform-provider-databricks/issues/656))
* Made preview environment tests to run on a release basis

Updated dependency versions:

* Bump github.com/zclconf/go-cty from 1.8.2 to 1.8.3
* Bump github.com/aws/aws-sdk-go from 1.38.30 to 1.38.51

## 0.3.4

* Fixed state refresh bugs in `databricks_sql_permissions` ([#620](https://github.com/databricks/terraform-provider-databricks/issues/620), [#619](https://github.com/databricks/terraform-provider-databricks/issues/620))
* Fixed `workspace_ids_filter` mapping for `databricks_mws_log_delivery` ([#635](https://github.com/databricks/terraform-provider-databricks/issues/635))
* Multiple documentation improvements ([#597](https://github.com/databricks/terraform-provider-databricks/issues/597), [eb60d10](https://github.com/databricks/terraform-provider-databricks/commit/eb60d103ea63221a1eb0069723ba3a0af45dbe3b), [edcd4b1](https://github.com/databricks/terraform-provider-databricks/commit/edcd4b121254e3ff3130bed9c4ef9d849d342561), [404bdab](https://github.com/databricks/terraform-provider-databricks/commit/404bdab637c0a4a15b6a4b6a77567166315955ca), [#615](https://github.com/databricks/terraform-provider-databricks/pull/615), [f14b825](https://github.com/databricks/terraform-provider-databricks/commit/f14b825e9cb11d75e9ad077b35c7e9c410fd8351), [e615c3a](https://github.com/databricks/terraform-provider-databricks/commit/e615c3a68d1ad45f91453ec448b55ca7b204fb97), [#612](https://github.com/databricks/terraform-provider-databricks/pull/612))
* Mounting clusters are recreated now, even when they are deleted ([#637](https://github.com/databricks/terraform-provider-databricks/issues/637))
* Fixed handling of empty blocks for clusters/jobs/instance pools ([22cdf2f](https://github.com/databricks/terraform-provider-databricks/commit/22cdf2fc9d50f67b14b49d11e7fbaacce0f52399))
* Mark instance pool attributes as ForceNew when it's requited ([#629](https://github.com/databricks/terraform-provider-databricks/issues/629))
* Switched to use https://staticcheck.io/ for static code analysis ([#602](https://github.com/databricks/terraform-provider-databricks/issues/602))

**Behavior changes**

* The `customer_managed_key_id` field in `databricks_mws_workspaces` resource is deprecated and should be replaced with `managed_services_customer_managed_key_id` (and optionally `storage_customer_managed_key_id`). `databricks_mws_customer_managed_keys` now requires the parameter `use_cases` ([#642](https://github.com/databricks/terraform-provider-databricks/pull/642)). *If you've used the resource before, please add `use_cases = ["MANAGED_SERVICES"]` to keep the behaviour.*

Updated dependency versions:

* Bump github.com/aws/aws-sdk-go to v1.38.30
* Bump github.com/hashicorp/go-retryablehttp to v0.7.0
* Bump github.com/hashicorp/hcl/v2 to v2.10.0
* Bump github.com/hashicorp/terraform-plugin-sdk/v2 to v2.6.1
* Bump github.com/zclconf/go-cty to v1.8.2

## 0.3.3

* Added resources for SQL Analytics queries and dashboards: `databricks_sql_query`, `databricks_sql_visualization`, `databricks_sql_dashboard`, `databricks_sql_widget` ([#553](https://github.com/databricks/terraform-provider-databricks/pull/553))
* Added `databricks_sql_permissions` resource ([#545](https://github.com/databricks/terraform-provider-databricks/pull/545/files))
* Fixed documentation bugs ([#603](https://github.com/databricks/terraform-provider-databricks/issues/603))
* Improved resource exporter ([#593](https://github.com/databricks/terraform-provider-databricks/issues/593))
* Added missing properties to `databricks_mws_private_access_settings` ([#590](https://github.com/databricks/terraform-provider-databricks/issues/590))
* Include SQLA data source ID in `databricks_sql_endpoint` state ([#601](https://github.com/databricks/terraform-provider-databricks/issues/601))
* Apply `debug_truncate_bytes` also for response dumps ([#589](https://github.com/databricks/terraform-provider-databricks/issues/589))
* More verbose logging of `databricks_cluster` termination reason ([#588](https://github.com/databricks/terraform-provider-databricks/issues/588))
* Move non-auth provider config documentation into separate section ([#587](https://github.com/databricks/terraform-provider-databricks/pull/587))


## 0.3.2

* Fixed minor issues to add support for GCP ([#558](https://github.com/databricks/terraform-provider-databricks/pull/558))
* Fixed `databricks_permissions` for SQL Analytics Entities ([#535](https://github.com/databricks/terraform-provider-databricks/issues/535))
* Fixed incorrect HTTP 404 handling on create ([#564](https://github.com/databricks/terraform-provider-databricks/issues/564), [#576](https://github.com/databricks/terraform-provider-databricks/issues/576))
* Fixed incorrect escaping of notebook names ([#566](https://github.com/databricks/terraform-provider-databricks/pull/566))
* Fixed entitlements for databricks_group ([#549](https://github.com/databricks/terraform-provider-databricks/pull/549))
* Fixed rate limiting to perform more than 1 request per second ([#577](https://github.com/databricks/terraform-provider-databricks/pull/577))
* Added support for spot instances on Azure ([#571](https://github.com/databricks/terraform-provider-databricks/pull/571))
* Added job schedules support for `pause_status` as a optional field. ([#575](https://github.com/databricks/terraform-provider-databricks/pull/575))
* Fixed minor documentation issues.

Updated dependency versions:

* Bump github.com/aws/aws-sdk-go from 1.37.20 to 1.38.10
* Bump github.com/hashicorp/hcl/v2 from 2.9.0 to 2.9.1 
* Bump github.com/zclconf/go-cty from 1.8.0 to 1.8.1
* Bump github.com/google/go-querystring from 1.0.0 to 1.1.0

## 0.3.1

* Added `databricks_global_init_script` resource to configure global init scripts ([#487](https://github.com/databricks/terraform-provider-databricks/issues/487)).
* Added `databricks_sql_endpoint` resource ([#498](https://github.com/databricks/terraform-provider-databricks/pull/498))
* Added [experimental resource exporter](https://github.com/databricks/terraform-provider-databricks/blob/master/docs/guides/experimental-exporter.md) to generate configuration for entire workspace.
* Improved user-facing documentaiton ([#508](https://github.com/databricks/terraform-provider-databricks/pull/508/files), [#516](https://github.com/databricks/terraform-provider-databricks/pull/516), [#511](https://github.com/databricks/terraform-provider-databricks/pull/511), [#504](https://github.com/databricks/terraform-provider-databricks/pull/504), [#492]([Update docs in various places](https://github.com/databricks/terraform-provider-databricks/pull/492)))
* Simplified authentication issues debugging ([#490](https://github.com/databricks/terraform-provider-databricks/pull/490))
* Made cleaner error message for no config profile ([#491](https://github.com/databricks/terraform-provider-databricks/pull/491))
* Allow tokens without comment or expiration ([#495](https://github.com/databricks/terraform-provider-databricks/pull/495/files))
* Ensured consistent slashes in notebook paths for different OSes ([#500](https://github.com/databricks/terraform-provider-databricks/pull/500))
* Fix error message panic in command result parsing ([#502](https://github.com/databricks/terraform-provider-databricks/pull/502))
* Updated `databricks_group` data resource to allow non-alphanumeric characters in group name filter ([#507](https://github.com/databricks/terraform-provider-databricks/pull/507/files))

**Behavior changes**

* Assigning any permission to `admins` would result in an error, so that behavior is consistent ([#486](https://github.com/databricks/terraform-provider-databricks/issues/486)).

Updated dependency versions:

* github.com/zclconf/go-cty from 1.2.1 to 1.7.1
* github.com/Azure/go-autorest/autorest/azure/auth from 0.5.6 to 0.5.7 
* github.com/hashicorp/hcl/v2 from 2.3.0 to 2.8.2
* github.com/aws/aws-sdk-go from 1.37.1 to 1.37.11
* github.com/Azure/go-autorest/autorest from 0.11.17 to 0.11.18

## 0.3.0

* Added configurable provisioning timeout for `databricks_mws_workspaces`, so that local DNS cache issues would be more tolerated.
* Added [databricks_current_user] to simplify applying the same Terraform configuration by different users in the shared workspace for testing purposes. 
* Added client-side rate limiting to release the pressure on backend APIs and prevent client blocking ([#465](https://github.com/databricks/terraform-provider-databricks/pull/465))
* Member usernames, group names and instance profile names in `databricks_group` data source are now sorted and providing consistent behavior between runs ([#449](https://github.com/databricks/terraform-provider-databricks/issues/449))
* Fixed redundant multiple mounting clusters ([#445](https://github.com/databricks/terraform-provider-databricks/issues/445))
* Added optional parameter azure_environment to provider config which defaults to public ([#437](https://github.com/databricks/terraform-provider-databricks/pull/437)).
* Added [databricks_service_principal](https://github.com/databricks/terraform-provider-databricks/pull/386) resource.
* `skip_validation` from `databricks_instance_profile` was removed and is always set to `true`.
* Added propagation of terraform version to `User-Agent` header, along with type of resource used.
* `databricks_notebook` & `databricks_dbfs_file` got new `source` field to specify location of a local file.
* `databricks_notebook` can have `language` field optional, as long as `source` is set to a file with `.py`, `.scala`, `.sql`, or `.r` extension.
* `databricks_me` data source was added to represent `user_name`, `home` & `id` of the caller user (or service principal).
* Added validation for secret scope name in `databricks_secret`, `databricks_secret_scope` and `databricks_secret_acl`. Non-compliant names may cause errors.
* Added [databricks_spark_version](https://github.com/databricks/terraform-provider-databricks/issues/347) data source.
* Fixed support for [single node clusters](https://docs.databricks.com/clusters/single-node.html) support by allowing [`num_workers` to be `0`](https://github.com/databricks/terraform-provider-databricks/pull/454).
* Fixed bug in destruction of IP access lists ([#426](https://github.com/databricks/terraform-provider-databricks/issues/426)).
* All resource imports are now making call to corresponding Databricks API by default ([#471](https://github.com/databricks/terraform-provider-databricks/issues/471)).

**Behavior changes**
* Removed deprecated `library_jar`, `library_egg`, `library_whl`, `library_pypi`, `library_cran`, and `library_maven` from `databricks_cluster` and `databricks_job` in favor of more API-transparent [library](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/cluster#library-configuration-block) configuration block.
* Removed deprecated `notebook_path` and `notebook_base_parameters` from `databricks_job` in favor of [notebook_task](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job#notebook_task-configuration-block) configuration block.
* Removed deprecated `jar_uri`, `jar_main_class_name`, and `jar_parameters` from `databricks_job` in favor of [spark_jar_task](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job#spark_jar_task-configuration-block) configuration block.
* Removed deprecated `python_file` and `python_parameters` from `databricks_job` in favor of [spark_python_task](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job#spark_python_task-configuration-block) configuration block.
* Removed deprecated `spark_submit_parameters` from `databricks_job` in favor of [spark_submit_task](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/job#spark_submit_task-configuration-block) configuration block.
* Removed deprecated `databricks_scim_user` resource in favor of [databricks_user](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/user).
* Removed deprecated `databricks_scim_group` resource in favor of [databricks_group](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/group).
* Removed deprecated `databricks_default_user_roles` data source in favor of [databricks_group](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/group#attribute-reference) data source.
* Removed deprecated `basic_auth` and `azure_auth` provider configuration blocks in favor of [documented authentication methods](https://registry.terraform.io/providers/databricks/databricks/latest/docs).
* `format`, `overwrite`, and `mkdirs` were removed from `databricks_notebook`. To follow expected behavior of Terraform, notebooks are always overwritten.
* `skip_validation` from `databricks_instance_profile` was removed and is always set to `true` for subsequent requests.
* `databricks_mws_workspace` got `verify_workspace_runnning` removed and now validates all every deployment. In case deployment failed, it removes workspace that failed and returns error message with explanation.
* `default_tags` were removed from `databricks_instance_pool`. `disk_spec` got new attribute `disk_type`, that contains `azure_disk_volume_type` and `ebs_volume_type`. This change is made to closer reflect API structure.
* `databricks_notebook` & `databricks_dbfs_file` got `content` attribute renamed to `content_base64` and now share the same logic to work with local files.

## 0.2.9

* Fixed documentation issues.
* Added missing resource importers and test to cover it.
* Migrated build from TravisCI to GitHub Actions.
* Fixed custom `config_file` issue configuration handling ([#420](https://github.com/databricks/terraform-provider-databricks/issues/420)).

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

* Added [databricks_workspace_conf](https://github.com/databricks/terraform-provider-databricks/pull/398) resource.
* Added [databricks_mws_log_delivery](https://github.com/databricks/terraform-provider-databricks/pull/343) resource for billable usage & audit logs consumption.
* Added [databricks_node_type](https://github.com/databricks/terraform-provider-databricks/pull/376) data source for simpler selection of node types across AWS & Azure.
* Added [Azure Key Vault support](https://github.com/databricks/terraform-provider-databricks/pull/381) for `databricks_secret_scope` for Azure CLI authenticated users.
* Added [is_pinned](https://github.com/databricks/terraform-provider-databricks/pull/348) support for `databricks_cluster` resource.
* Fixed restarting cluster on changes in cluster configuration aren't related to the cluster configuration ([issue #379](https://github.com/databricks/terraform-provider-databricks/issues/379))
* Fixed issue [#383](https://github.com/databricks/terraform-provider-databricks/issues/383) by cleaning up clusters that fail to start.
* Fixed issue [#382](https://github.com/databricks/terraform-provider-databricks/issues/382) by ignoring any incoming changes to deployment name of `databricks_mws_workspaces`, as well as propagating the right error messages.
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

* Added support for [customer managed keys](https://github.com/databricks/terraform-provider-databricks/pull/332) for Accounts API.
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
* State changes to legacy `spark.databricks.delta.preview.enabled` config option are [now ignored](https://github.com/databricks/terraform-provider-databricks/pull/334) by `databricks_job` & `databricks_cluster`
* Libraries, which are installed on all clusters and are not part of cluster resource definition, won't be waited for INSTALLED status
* Fixed "[Secret scope ACL is MANAGE for all users by default](https://github.com/databricks/terraform-provider-databricks/pull/326)" ([issue 322](https://github.com/databricks/terraform-provider-databricks/issues/322)).  If you were relying on setting `MANAGE` permission to all users by default, you need to add `initial_manage_principal = "users"` to your `resource "databricks_secret_scope"` declaration. 

## 0.2.5

* Added support for [local disk encryption](https://github.com/databricks/terraform-provider-databricks/pull/313)
* Added more reliable [indication about Azure environment](https://github.com/databricks/terraform-provider-databricks/pull/312) and fixed azure auth issue for Terraform 0.13
* Updated [databricks_aws_crossaccount_policy](https://github.com/databricks/terraform-provider-databricks/pull/311) to latest rules
* Fixed missing importers for [databricks_scim_*](https://github.com/databricks/terraform-provider-databricks/pull/290) resources
* Updated [Terraform Plugin SDK](https://github.com/databricks/terraform-provider-databricks/pull/279) to latest version along with transitive dependencies.
* Added support disclaimers
* Increased code coverage to 65%

## 0.2.4

* Added [Azure CLI authentication](https://github.com/databricks/terraform-provider-databricks/blob/master/docs/index.md#authenticating-with-azure-cli) to bridge the gap of local development workflows and let more people use the provider.
* All authentication is completely [lazy-initialized](https://github.com/databricks/terraform-provider-databricks/pull/270), which makes it provider overall more stable.
* Significantly increased [unit test coverage](https://codecov.io/gh/databricks/terraform-provider-databricks), which runs before every merge of a pull request.
* Introduced constantly running integration test suite environments: azsp, azcli & awsmt
* Numerous stability improvements for clusters, mounts, libraries, notebooks, files, authentication and TLS connectivity.
* Added ability to mount storage without explicitly defining a cluster, though it will still launch auto-terminating `terraform-mount` cluster to perform the mount.
* `databricks_cluster` & `databricks_job` now share significant portion of configuration wiring code, therefore increasing the stability of codebase.
* Added support for Terraform 0.13 [local builds](https://github.com/databricks/terraform-provider-databricks/pull/281) for those who develop or cannot wait for next release.
* Added AWS IAM Policy [data helpers](https://github.com/databricks/terraform-provider-databricks/pull/255) to simplify new deployments.
* [Migrated all documentation](https://github.com/databricks/terraform-provider-databricks/pull/250) to Terraform Registry format, therefore having a single always-accurate place for end-user guides.
* Internally, codebase [has been split](https://github.com/databricks/terraform-provider-databricks/pull/224) into multiple packages, which should make further contributions simpler.

Updated dependency versions:

* github.com/Azure/go-autorest/autorest v0.11.4
* github.com/Azure/go-autorest/autorest/adal v0.9.2
* github.com/Azure/go-autorest/autorest/azure/auth v0.5.1
* github.com/aws/aws-sdk-go v1.34.13
* gopkg.in/ini.v1 v1.60.2

**Deprecations**
* `library_*` is no longer receiving fixes and will be removed in `0.3`, please rewrite cluster & job resources to use [`library` configuration block](https://github.com/databricks/terraform-provider-databricks/blob/master/docs/resources/cluster.md#library-configuration-block).
* `basic_auth` provider block is no longer receiving fixesand will be removed in `0.3`, please use `username` and `password` options
* `azure_auth` provider block is no longer receiving fixesand will be removed in `0.3`, please use `azure_*` options 

**Behavior changes**
* Previously, mounts code paths were different functions. This release unifies them to be a single testable codebase with different configuration options & re-use of the critical code paths. For maintainability reasons, there's no longer check performed on container & storage account names, but rather on high-level *mount source uri*.
