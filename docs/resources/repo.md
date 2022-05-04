---
subcategory: "Workspace"
---
# databricks_repo Resource

This resource allows you to manage [Databricks Repos](https://docs.databricks.com/repos.html).

-> **Note** To create a Repo from a private repository you need to configure Git token as described in the [documentation](https://docs.databricks.com/repos.html#configure-your-git-integration-with-databricks).  To set this token you can use [databricks_git_credential](git_credential.md) resource.

## Example Usage

You can declare Terraform-managed Repo by specifying `url` attribute of Git repository. In addition to that you may need to specify `git_provider` attribute if Git provider doesn't belong to cloud Git providers (Github, GitLab, ...).  If `path` attribute isn't provided, then repo will be created in the user's repo directory (`/Repos/<username>/...`):


```hcl
resource "databricks_repo" "nutter_in_home" {
  url = "https://github.com/user/demo.git"
}
```

## Argument Reference

-> **Note** Repo in Databricks workspace would only be changed, if Terraform stage did change. This means that any manual changes to managed repository won't be overwritten by Terraform, if there's no local changes to configuration. If Repo in Databricks workspace is modifying, application of configuration changes will fail.

The following arguments are supported:

* `url` -  (Required) The URL of the Git Repository to clone from. If value changes, repo is re-created.
* `git_provider` - (Optional, if it's possible to detect Git provider by host name) case insensitive name of the Git provider.  Following values are supported right now (could be a subject for a change, consult [Repos API documentation](https://docs.databricks.com/dev-tools/api/latest/repos.html)): `gitHub`, `gitHubEnterprise`, `bitbucketCloud`, `bitbucketServer`, `azureDevOpsServices`, `gitLab`, `gitLabEnterpriseEdition`, , `awsCodeCommit`.
* `path` - (Optional) path to put the checked out Repo. If not specified, then repo will be created in the user's repo directory (`/Repos/<username>/...`).  If value changes, repo is re-created.
* `branch` - (Optional) name of the branch for initial checkout. If not specified, the default branch of the repository will be used.  Conflicts with `tag`.  If `branch` is removed, and `tag` isn't specified, then the repository will stay at the previously checked out state.
* `tag` - (Optional) name of the tag for initial checkout.  Conflicts with `branch`.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` -  Repo identifier
* `commit_hash` - Hash of the HEAD commit at time of the last executed operation. It won't change if you manually perform pull operation via UI or API

## Access Control

* [databricks_permissions](permissions.md#Repos-usage) can control which groups or individual users can access repos.

## Import

The resource Repo can be imported using the Repo ID (obtained via UI or using API)

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
