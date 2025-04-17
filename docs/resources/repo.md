---
subcategory: "Workspace"
---
# databricks_repo Resource

This resource allows you to manage [Databricks Git folders](https://docs.databricks.com/en/repos/index.html) (formerly known as Databricks Repos).

-> This resource can only be used with a workspace-level provider!

-> To create a Git folder from a private repository you need to configure Git token as described in the [documentation](https://docs.databricks.com/en/repos/index.html#configure-your-git-integration-with-databricks).  To set this token you can use [databricks_git_credential](git_credential.md) resource.

## Example Usage

You can declare Terraform-managed Git folder by specifying `url` attribute of Git repository. In addition to that you may need to specify `git_provider` attribute if Git provider doesn't belong to cloud Git providers (Github, GitLab, ...).  If `path` attribute isn't provided, then Git folder will be created in the default location:


```hcl
resource "databricks_repo" "nutter_in_home" {
  url = "https://github.com/user/demo.git"
}
```

## Argument Reference

-> Git folder in Databricks workspace would only be changed, if Terraform stage did change. This means that any manual changes to managed repository won't be overwritten by Terraform, if there's no local changes to configuration. If Git folder in Databricks workspace is modified, application of configuration changes will fail.

The following arguments are supported:

* `url` -  (Required) The URL of the Git Repository to clone from. If the value changes, Git folder is re-created.
* `git_provider` - (Optional, if it's possible to detect Git provider by host name) case insensitive name of the Git provider.  Following values are supported right now (could be a subject for a change, consult [Repos API documentation](https://docs.databricks.com/dev-tools/api/latest/repos.html)): `gitHub`, `gitHubEnterprise`, `bitbucketCloud`, `bitbucketServer`, `azureDevOpsServices`, `gitLab`, `gitLabEnterpriseEdition`, `awsCodeCommit`.
* `path` - (Optional) path to put the checked out Git folder. If not specified, , then the Git folder will be created in the default location.  If the value changes, Git folder is re-created.
* `branch` - (Optional) name of the branch for initial checkout. If not specified, the default branch of the repository will be used.  Conflicts with `tag`.  If `branch` is removed, and `tag` isn't specified, then the repository will stay at the previously checked out state.
* `tag` - (Optional) name of the tag for initial checkout.  Conflicts with `branch`.

### sparse_checkout

Optional `sparse_checkout` configuration block contains attributes related to [sparse checkout feature](https://docs.databricks.com/repos/git-operations-with-repos.html#configure-sparse-checkout-mode) in Databricks Git folders.  It supports following attributes:

* `patterns` - array of paths (directories) that will be used for sparse checkout.  List of patterns could be updated in-place.

Addition or removal of the `sparse_checkout` configuration block will lead to recreation of the Git folder.


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` -  Git folder identifier
* `commit_hash` - Hash of the HEAD commit at time of the last executed operation. It won't change if you manually perform pull operation via UI or API
* `workspace_path` - path on Workspace File System (WSFS) in form of `/Workspace` + `path`

## Access Control

* [databricks_permissions](permissions.md#Repos-usage) can control which groups or individual users can access repos.

## Import

The resource can be imported using the Git folder ID (obtained via UI or using API)

```bash
$ terraform import databricks_repo.this repo_id
```

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_git_credential](git_credential.md) to manage Git credentials.
* [databricks_directory](directory.md) to manage directories in [Databricks Workpace](https://docs.databricks.com/workspace/workspace-objects.html).
* [databricks_pipeline](pipeline.md) to deploy [Delta Live Tables](https://docs.databricks.com/data-engineering/delta-live-tables/index.html). 
* [databricks_secret](secret.md) to manage [secrets](https://docs.databricks.com/security/secrets/index.html#secrets-user-guide) in Databricks workspace.
* [databricks_secret_acl](secret_acl.md) to manage access to [secrets](https://docs.databricks.com/security/secrets/index.html#secrets-user-guide) in Databricks workspace.
* [databricks_secret_scope](secret_scope.md) to create [secret scopes](https://docs.databricks.com/security/secrets/index.html#secrets-user-guide) in Databricks workspace.
* [databricks_workspace_conf](workspace_conf.md) to manage workspace configuration for expert usage.
