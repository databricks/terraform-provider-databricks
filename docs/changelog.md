# Version changelog

## 0.2.6

* Added `databricks_user` resource.

**Deprecations**
* `databricks_scim_user` is no longer receiving fixes and will be removed in `0.3`, please rewrite using `databricks_user` resource, which has more consistent semantics with `databricks_group` and works better with identity provider SCIM sync.
* `databricks_scim_group` is no longer receiving fixes and will be removed in `0.3`, please rewrite using `databricks_group` resource.

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
