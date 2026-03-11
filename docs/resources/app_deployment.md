---
subcategory: "Apps"
---
# databricks_app_deployment Resource

-> **Note** This resource is in Public Preview.

Deploys code to a [Databricks App](https://docs.databricks.com/en/dev-tools/databricks-apps/index.html). Deployments can be created from either workspace filesystem paths or Git repositories.

-> This resource can only be used with a workspace-level provider!

## Example Usage

### Deploy from Workspace Filesystem

```hcl
resource "databricks_workspace_file" "app_code" {
  source = "${path.module}/app.py"
  path   = "/Workspace/apps/my-app/app.py"
}

resource "databricks_app" "this" {
  name = "my-app"
}

resource "databricks_app_deployment" "this" {
  depends_on = [databricks_workspace_file.app_code]

  app_name         = databricks_app.this.name
  source_code_path = "/Workspace/apps/my-app"

  triggers = {
    # Force redeploy when source code changes
    source_hash = sha1(join("", [for f in fileset("${path.module}/app", "**/*") : filesha1("${path.module}/app/${f}")]))
  }
}
```

### Deploy from Git Branch

-> **Note** Git-based deployments require the `git_repository` block on [`databricks_app`](app.md) as the Git repository URL is defined at the app level.

```hcl
resource "databricks_git_credential" "github" {
  git_provider          = "gitHub"
  git_username          = "myuser"
  personal_access_token = var.github_token
}

resource "databricks_app" "this" {
  name = "my-app"

  # Configure Git repository at app level
  git_repository = {
    provider = "gitHub"
    url      = "https://github.com/myorg/myrepo"
  }
}

resource "databricks_app_deployment" "this" {
  depends_on = [databricks_git_credential.github]

  app_name = databricks_app.this.name

  git_source = {
    branch = "main"
  }

  triggers = {
    deploy_version = "v1"  # Change to force redeploy
  }
}
```

### Deploy from Git Branch with Subdirectory

```hcl
resource "databricks_app_deployment" "this" {
  depends_on = [databricks_git_credential.github]

  app_name = databricks_app.this.name

  git_source = {
    branch           = "main"
    source_code_path = "/"  # Use "/" for root, or "/apps/myapp" for subdirectory
  }

  triggers = {
    deploy_version = "v2"
  }
}
```

### Deploy from Git Tag

```hcl
resource "databricks_app_deployment" "this" {
  depends_on = [databricks_git_credential.github]

  app_name = databricks_app.this.name

  git_source = {
    tag = "v1.2.3"
  }

  triggers = {
    version = "v1.2.3"
  }
}
```

### Deploy from Git Commit

```hcl
resource "databricks_app_deployment" "this" {
  depends_on = [databricks_git_credential.github]

  app_name = databricks_app.this.name

  git_source = {
    commit = "abc123def456789"
  }

  triggers = {
    commit = "abc123def456789"
  }
}
```

## Argument Reference

The following arguments are required:

* `app_name` - (Required) The name of the app to deploy to. Changing this forces a new deployment.

Exactly one of the following is required:

* `source_code_path` - (Optional) The workspace filesystem path of the source code to deploy (e.g., `/Workspace/apps/my-app`). Conflicts with `git_source`. Changing this forces a new deployment.
* `git_source` - (Optional) Git source configuration for the deployment. Conflicts with `source_code_path`. See [git_source Configuration](#git_source-configuration) below. Changing this forces a new deployment.

The following arguments are optional:

* `mode` - (Optional) The deployment mode. Allowed values are `SNAPSHOT` (default) and `AUTO_SYNC`. Changing this forces a new deployment.
* `triggers` - (Optional) A map of arbitrary string key/value pairs that, when changed, will force a new deployment. This can be used to force redeployment when source code changes.

### git_source Configuration

The `git_source` block requires exactly one of the following:

* `branch` - (Optional) Git branch to checkout (e.g., `main`). Exactly one of `branch`, `tag`, or `commit` must be specified.
* `tag` - (Optional) Git tag to checkout (e.g., `v1.0.0`). Exactly one of `branch`, `tag`, or `commit` must be specified.
* `commit` - (Optional) Git commit SHA to checkout (e.g., `abc123def456`). Exactly one of `branch`, `tag`, or `commit` must be specified.

The following is optional:

* `source_code_path` - (Optional) Relative path to the app source code within the Git repository. Use `"/"` for the root of the repository, or a path like `"/apps/myapp"` for a subdirectory. If not specified, defaults to the root. **Note:** This field may not be available in all workspaces.

-> **Note** The `git_repository` (provider and URL) must be configured at the app level using the `databricks_app` resource's `git_repository` block. See the examples above.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `deployment_id` - The unique ID of the deployment.
* `status` - The status of the deployment (e.g., `SUCCEEDED`, `FAILED`, `IN_PROGRESS`).
* `create_time` - The creation time of the deployment in ISO 8601 format.
* `git_source.resolved_commit` - (For Git deployments) The actual commit SHA that was deployed. For branch or tag deployments, this shows the commit that the reference pointed to at deployment time.

## Import

This resource can be imported using the app name and deployment ID:

```sh
terraform import databricks_app_deployment.this <app_name>/<deployment_id>
```

For example:

```sh
terraform import databricks_app_deployment.this my-app/01f1234567890abc
```

## Git Deployment Requirements

To deploy from Git repositories, you must:

1. **Configure Git credentials** at the workspace level using `databricks_git_credential`:
   ```hcl
   resource "databricks_git_credential" "github" {
     git_provider          = "gitHub"
     git_username          = "myuser"
     personal_access_token = var.github_token
   }
   ```

2. **Configure the Git repository** at the app level in `databricks_app`:
   ```hcl
   resource "databricks_app" "this" {
     name = "my-app"
     git_repository = {
       provider = "gitHub"
       url      = "https://github.com/myorg/myrepo"
     }
   }
   ```

3. **Reference the Git source** in your deployment:
   ```hcl
   resource "databricks_app_deployment" "this" {
     app_name = databricks_app.this.name
     git_source = {
       branch = "main"
     }
   }
   ```

### Supported Git Providers

The following Git providers are supported:
- `gitHub` - GitHub
- `gitHubEnterprise` - GitHub Enterprise Server
- `gitLab` - GitLab
- `gitLabEnterpriseEdition` - GitLab Enterprise Edition
- `bitbucketCloud` - Bitbucket Cloud
- `bitbucketServer` - Bitbucket Server / Data Center
- `azureDevOpsServices` - Azure DevOps Services
- `awsCodeCommit` - AWS CodeCommit

## Deployment Behavior

* **Immutable deployments**: All deployment properties force replacement. To update a deployment, Terraform will create a new deployment and remove the old one from state.
* **No API deletion**: Deployments cannot be deleted via the Databricks API. Removing a deployment from Terraform state does not delete it from Databricks.
* **Deployment wait**: The resource waits for the deployment to reach a terminal state (`SUCCEEDED` or `FAILED`) before completing.

## Related Resources

* [databricks_app](app.md) to manage Databricks Apps.
* [databricks_git_credential](git_credential.md) to manage Git credentials for private repositories.
