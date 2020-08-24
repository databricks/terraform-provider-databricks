# Version changelog

## 0.2.4

* Significantly increased unit test coverage, which runs before every merge of a pull request.
* Introduced constantly running integration test suite environments: azsp, azcli & awsmt
* Added ability to mount storage without explicitly defining a cluster, though it will still launch auto-terminating `terraform-mount` cluster to perform the mount.
* `databricks_cluster` & `databricks_job` now share significant portion of configuration wiring code, therefore increasing the stability of codebase.

**Deprecations**
* `library_*` is no longer receiving fixes and will be removed in `0.3`, please rewrite cluster & job resources to use `libraries` config block.
* `basic_auth` provider block is no longer receiving fixesand will be removed in `0.3`, please use `username` and `password` options
* `azure_auth` provider block is no longer receiving fixesand will be removed in `0.3`, please use `azure_*` options 

**Behavior changes**
* Previously, mounts code paths were different functions. This release unifies them to be a single testable codebase with different configuration options & re-use of the critical code paths. For maintainability reasons, there's no longer check performed on container & storage account names, but rather on high-level *mount source uri*.