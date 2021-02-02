# Migration from 0.2.x to 0.3.x

Certain resources undergone changes in order to ensure consistency with REST API and standard expected Terraform behavior. You can upgrade provider with `terraform init -upgrade`.

## provider

* Rewrite `basic_auth` block with `username` and `password` fields, as specified in [main document](https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs#authenticating-with-hostname-username-and-password).
* Rewrite `azure_auth` block with appropriate [Azure configuration](https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs#special-configurations-for-azure).

## databricks_job

* Rewrite `spark_submit_parameters` with [spark_submit_task](https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs/resources/job#spark_submit_task-configuration-block) configuration block.
* Rewrite `python_file` and `python_parameters` with [spark_python_task](https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs/resources/job#spark_python_task-configuration-block) configuration block.
* Rewrite `jar_uri`, `jar_main_class_name`, and `jar_parameters` with [spark_jar_task](https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs/resources/job#spark_jar_task-configuration-block) configuration block.
* Rewrite `notebook_path` and `notebook_base_parameters` with [notebook_task](https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs/resources/job#notebook_task-configuration-block) configuration block.
* Rewrite `library_jar`, `library_egg`, `library_whl`, `library_pypi`, `library_cran`, and `library_maven` with [library](https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs/resources/cluster#library-configuration-block) configuration block.

## databricks_dbfs_file

* Rename `content` to `content_base64`, as this closer represents actual data within the field and simplifies internal code reusability.
* Remove `overwrite` attribute. Starting from v0.3.0 it behaves as if it is set to `true`.
* Remove `mkdirs` attribute. Starting from v0.3.0 it behaves as if it is set to `true`.
* Remove `validate_remote_file` attribute. Due to performance reasons, starting from v0.3.0 it doesn't fetch the contents of remote file to verify the checksum. 
* If you've relied on internal `content_b64_md5` attribute, please remove it. Starting from v0.3.0 its behavior is internalized.

DBFS files would only be changed, if Terraform stage did change. This means that any manual changes to managed file won't be overwritten by Terraform, if there's no local change. 

## databricks_notebook

* Rename `content` to `content_base64`, as this closer represents actual data within the field and simplifies internal code reusability.
* Remove `format` attribute. Starting from v0.3.0 it behaves as if it is set to `SOURCE`.
* Remove `overwrite` attribute. Starting from v0.3.0 it behaves as if it is set to `true`.
* Remove `mkdirs` attributes. Starting from v0.3.0 it behaves as if it is set to `true`.

After changing the code, `terraform apply` would replace managed notebooks.

Notebook on Databricks workspace would only be changed, if Terraform stage did change. This means that any manual changes to managed notebook won't be overwritten by Terraform, if there's no local change to notebook sources. Notebooks are identified by their path, so changing notebook's name manually on the workspace and then applying Terraform state would result in creation of notebook from Terraform state.

## databricks_cluster

* Rewrite `library_jar`, `library_egg`, `library_whl`, `library_pypi`, `library_cran`, and `library_maven` with [library](https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs/resources/cluster#library-configuration-block) configuration block.

## databricks_instance_profile

* Remove `skip_validation` from all `databricks_instance_profile` resources. In order to ensure consistency, all AWS EC2 profiles are now checked to work before returning state to main Terraform process.

## databricks_mws_workspaces

* Remove `verify_workspace_runnning` attribute from all `databricks_mws_workspaces` resources. All workspaces are verified to be running automatically as of [this change](https://github.com/databrickslabs/terraform-provider-databricks/commit/ef64b5d26daa23ff2532f1076a0db01864e4f73c).

## databricks_instance_pool

* Remove `default_tags`.

## databricks_scim_user

* This resource was removed as deprecated. Please rewrite using [databricks_user](../resources/user.md).

## databricks_scim_group

* This resource was removed as deprecated. Please rewrite using [databricks_group](../resources/group.md).

## databricks_default_user_roles

* This data source was removed as deprecated. Please use [databricks_group](../data-sources/group.md) data source for performing equivalent tasks.