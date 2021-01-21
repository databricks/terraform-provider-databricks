# Migration from 0.2.x to 0.3.x

Certain resources undergone changes in order to ensure consistency with 

## databricks_mws_workspaces

* Remove `verify_workspace_runnning` attributes from all `databricks_mws_workspaces` resources. All workspaces are verified to be running automatically as of [this change](https://github.com/databrickslabs/terraform-provider-databricks/commit/ef64b5d26daa23ff2532f1076a0db01864e4f73c).