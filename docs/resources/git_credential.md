---
subcategory: "Workspace"
---
# databricks_git_credential Resource

This resource allows you to manage credentials for [Databricks Repos](https://docs.databricks.com/repos.html) using [Git Credentials API](https://docs.databricks.com/dev-tools/api/latest/gitcredentials.html).

-> This resource can only be used with a workspace-level provider!

## Example Usage

You can declare Terraform-managed Git credential using following code:

```hcl
resource "databricks_git_credential" "ado" {
  git_username          = "myuser"
  git_provider          = "azureDevOpsServices"
  personal_access_token = "sometoken"
}
```

## Argument Reference

The following arguments are supported:

* `personal_access_token` - (Optional, required for some Git providers) The personal access token used to authenticate to the corresponding Git provider. If value is not provided, it's sourced from the first environment variable of [`GITHUB_TOKEN`](https://registry.terraform.io/providers/integrations/github/latest/docs#oauth--personal-access-token), [`GITLAB_TOKEN`](https://registry.terraform.io/providers/gitlabhq/gitlab/latest/docs#required), or [`AZDO_PERSONAL_ACCESS_TOKEN`](https://registry.terraform.io/providers/microsoft/azuredevops/latest/docs#argument-reference), that has a non-empty value.
* `git_username` - (Optional, required for some Git providers) user name at Git provider.
* `git_provider` -  (Required) case insensitive name of the Git provider.  Following values are supported right now (could be a subject for a change, consult [Git Credentials API documentation](https://docs.databricks.com/dev-tools/api/latest/gitcredentials.html)): `gitHub`, `gitHubEnterprise`, `bitbucketCloud`, `bitbucketServer`, `azureDevOpsServices`, `gitLab`, `gitLabEnterpriseEdition`, `awsCodeCommit`, `azureDevOpsServicesAad`.
* `force` - (Optional) specify if settings need to be enforced - right now, Databricks allows only single Git credential, so if it's already configured, the apply operation will fail.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - identifier of specific Git credential

## Import

The resource cluster can be imported using ID of Git credential that could be obtained via REST API:

```hcl
import {
  to = databricks_git_credential.this
  id = "<git-credential-id>"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_git_credential.this <git-credential-id>
```

## Related Resources

The following resources are often used in the same context:

* [databricks_repo](repo.md) to manage Databricks Repos.
