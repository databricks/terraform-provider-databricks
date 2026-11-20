---
subcategory: "Postgres"
---
# databricks_postgres_branch Resource
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/postgres)

### Lakebase Autoscaling Terraform Behavior

These resources use a spec/status model: the `spec` field defines the resource configuration Terraform manages, and `status` reports the resource's current state.

Provider versions without drift detection do not detect changes made through the UI, CLI, or API. They apply updates based on changes to your Terraform configuration.

With a provider version that supports drift detection, `terraform plan` detects changes to managed `spec` fields made outside Terraform.

For help resolving provider inconsistency errors or recovering from a partially completed apply, see [Terraform state and drift detection for Lakebase](https://docs.databricks.com/aws/en/oltp/projects/terraform-state-and-drift). Use shared Terraform state storage when collaborating on the same resources.

### Upgrading to a Provider with Drift Detection

When upgrading from a provider version without drift detection to one that supports it, follow the steps below to resolve initial or persistent diffs.

Prefer a complete `spec` that describes the resource's full desired configuration. The examples below show the minimal valid specs for an existing resource, not the recommended end state.

#### Replace a Null Spec

Null specs can produce persistent differences against the object returned during refresh; repeated applies do not resolve those differences.

For a Branch that wasn't restored from a snapshot, replace `spec = null` with an empty object:

```hcl
spec = {}
```

For a Branch restored from a snapshot, replace `spec = null` with a spec containing its existing `source_snapshot`. Omitting this immutable value can cause Terraform to plan a replacement that deletes and recreates the Branch on apply:

```hcl
spec = {
  source_snapshot = "projects/my-project/snapshots/my-snapshot"
}
```

Use the Branch's existing snapshot name. Review and apply the plan, then run `terraform plan` again.

#### Reconcile an Initial Diff with an Existing Spec

A Branch created before Lakebase supported drift detection can show an initial in-place update even when its current values already match your configured spec. Review the plan against your intended configuration. For a Branch created through Terraform, retain any source fields already configured in its spec. Apply the reviewed plan, then run `terraform plan` again.

#### Resolve a Persistent Diff

If differences remain after applying, follow the [shared reconciliation steps](https://docs.databricks.com/aws/en/oltp/projects/terraform-state-and-drift#reconcile-existing-resources-after-rollout) to choose between the values in your configuration and the current values returned by Lakebase, then review the resulting plan.

### Import limitations

See [Import](#import) for import commands.

Configuring `source_branch`, `source_branch_lsn`, or `source_branch_time` on an imported Branch is currently unsupported. Omit these fields from its spec. This restriction does not apply to `source_snapshot`: for a Branch restored from a snapshot, include its existing snapshot name.

### Branch Expiration in Terraform

Prefer to use a fixed `expire_time` instead of relative `ttl`. Each update that sends `ttl` recalculates expiration as the update time plus that duration, so reapplying the same spec can extend the Branch's lifetime. Terraform configurations should describe a stable deadline.

For a complete example, see [Branch with an Absolute Expiration](#branch-with-expiration-ttl).

To preserve an existing deadline, use its current `status.expire_time`. Keep the timestamp fixed across plans; don't derive it from `timestamp()` or another value that advances on every run. Remove `ttl` and `no_expiry` when setting `expire_time`; these expiration options are mutually exclusive. Use `no_expiry = true` if you intend the Branch to have no expiration.

### Timestamp Formats

Use UTC timestamps ending in `Z`, such as `"2030-01-01T00:00:00Z"`, rather than numeric time-zone offsets. This applies to `expire_time` and, for Branches created through Terraform, `source_branch_time`.

### Overview

A Postgres branch is an independent database environment within a project that shares storage with its parent branch through copy-on-write. Branches allow you to create isolated development and testing environments, test applications against realistic data sets, and perform point-in-time recovery operations.

### Hierarchy Context

Branches exist within the Lakebase Autoscaling resource hierarchy:
- Each **branch** belongs to a **project**
- A **branch** can contain multiple **endpoints**, **databases**, and **roles**
- Branches can be created from other branches for point-in-time recovery

### Use Cases

- **Development environments**: Create isolated branches for feature development without affecting production
- **Testing**: Spin up temporary branches for testing changes before applying to production
- **Point-in-time recovery**: Create a branch from a specific point in time on another branch
- **Restore from a snapshot**: Create a branch whose data comes from an existing snapshot
- **Data exploration**: Safely query and analyze data without risking production workloads


## Example Usage
### Managing Implicitly Created Root Branch

A root branch named `production` is implicitly created for every project. Since Terraform is declarative, managing an already-existing resource requires `replace_existing = true`: it lets Terraform represent the implicitly created branch in Terraform state and immediately apply the provided configuration to it. Support for providing a custom `branch_id` will be available in later versions.

For this implicitly created root Branch, Terraform manages updates only. It does not control creation or deletion of the branch itself. Removing the resource from your Terraform configuration only removes it from Terraform state; the actual branch is unaffected, because its lifecycle is currently controlled by the parent project. The only way to remove the actual branch is to delete the project it belongs to.

```hcl
resource "databricks_postgres_project" "this" {
  project_id      = "my-project"
  spec = {
    pg_version   = 17
    display_name = "My Project"
  }

  lifecycle {
    prevent_destroy = true
  }
}

resource "databricks_postgres_branch" "production" {
  branch_id = "production"
  parent    = databricks_postgres_project.this.name
  spec = {
    no_expiry = true
  }
  replace_existing = true
}
```

### Basic Branch Creation

```hcl
resource "databricks_postgres_branch" "dev" {
  branch_id = "dev-branch"
  parent    = databricks_postgres_project.this.name
  spec = {
    no_expiry = true
  }
}
```

### Protected Branch

Only one branch per project can be protected at a time.

```hcl
resource "databricks_postgres_branch" "protected" {
  branch_id = "protected-branch"
  parent    = databricks_postgres_project.this.name
  spec = {
    is_protected = true
    no_expiry    = true
  }
}
```

### <a id="branch-with-expiration-ttl"></a>Branch with an Absolute Expiration

Prefer to use a fixed future `expire_time` instead of relative `ttl`. Reapplying a relative TTL recalculates the deadline from the update time and can extend the branch's lifetime. Replace the example timestamp with your intended deadline.

```hcl
resource "databricks_postgres_branch" "temporary" {
  branch_id = "temp-feature-test"
  parent    = databricks_postgres_project.this.name
  spec = {
    expire_time = "2030-01-01T00:00:00Z"
  }
}
```

### Branch from a Snapshot

Create a branch whose data comes from an existing snapshot instead of a source branch. The snapshot must be `AVAILABLE` and belong to the same project. `source_snapshot` is immutable and mutually exclusive with `source_branch`. The restored snapshot is reported back in `status.source_snapshot`.

```hcl
resource "databricks_postgres_branch" "from_snapshot" {
  branch_id = "restored-branch"
  parent    = databricks_postgres_project.this.name
  spec = {
    source_snapshot = "projects/my-project/snapshots/my-snapshot"
    no_expiry       = true
  }
}
```


## Arguments
The following arguments are supported:
* `branch_id` (string, required) - The part of the name, chosen by the user when the resource was created
* `parent` (string, required) - The project containing this branch (API resource hierarchy).
  Format: projects/{project_id}
  
  Note: This field indicates where the branch exists in the resource hierarchy.
  For point-in-time branching from another branch, see `status.source_branch`
* `purge_on_delete` (boolean, optional) - If true, permanently delete the branch; if false, soft delete
* `replace_existing` (boolean, optional) - If true, update the branch if it already exists instead of returning an error
* `spec` (BranchSpec, optional) - The spec contains the branch configuration.
  
  Read the resource's current values from `status`, rather than `spec`.
  `spec` tracks field ownership for infrastructure-as-code tools and may omit
  unmanaged fields
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

### BranchSpec
* `expire_time` (string, optional) - Absolute expiration timestamp. When set, the branch will expire at this time.
  Mutually exclusive with `ttl` and `no_expiry`
* `is_protected` (boolean, optional) - When set to true, protects the branch from deletion and reset. Associated compute endpoints and the project cannot be deleted while the branch is protected
* `no_expiry` (boolean, optional) - Explicitly disable expiration. When set to true, the branch will not expire.
  If set to false, the request is invalid; provide either ttl or expire_time instead.
  Mutually exclusive with `expire_time` and `ttl`
* `source_branch` (string, optional) - The name of the source branch from which this branch was created (data lineage for point-in-time recovery).
  If not specified, defaults to the project's default branch.
  Format: projects/{project_id}/branches/{branch_id}
* `source_branch_lsn` (string, optional) - The Log Sequence Number (LSN) on the source branch from which this branch was created
* `source_branch_time` (string, optional) - The point in time on the source branch from which this branch was created
* `source_snapshot` (string, optional) - The snapshot this branch was created from. When set, the branch's data
  comes from the snapshot rather than a source branch, so source_branch,
  source_branch_lsn, and source_branch_time must be empty. The snapshot must
  be AVAILABLE and belong to this branch's project.
  Format: projects/{project_id}/snapshots/{snapshot_id}
* `ttl` (string, optional) - Relative time-to-live duration. When set, the branch will expire at creation_time + ttl.
  Mutually exclusive with `expire_time` and `no_expiry`

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - A timestamp indicating when the branch was created
* `name` (string) - Output only. The full resource path of the branch.
  Format: projects/{project_id}/branches/{branch_id}
* `status` (BranchStatus) - The current status of a Branch
* `uid` (string) - System-generated unique ID for the branch
* `update_time` (string) - A timestamp indicating when the branch was last updated

### BranchStatus
* `branch_id` (string) - Part of the resource name
* `current_state` (string) - The branch's state, indicating if it is initializing, ready for use, or archived. Possible values are: `ARCHIVED`, `DELETED`, `IMPORTING`, `INIT`, `READY`, `RESETTING`
* `default` (boolean) - Whether the branch is the project's default branch
* `delete_time` (string) - A timestamp indicating when the branch was deleted.
  Empty if the branch is not deleted
* `expire_time` (string) - Absolute expiration time for the branch. Empty if expiration is disabled
* `is_protected` (boolean) - Whether the branch is protected
* `logical_size_bytes` (integer) - The logical size of the branch
* `pending_state` (string) - The pending state of the branch, if a state transition is in progress. Possible values are: `ARCHIVED`, `DELETED`, `IMPORTING`, `INIT`, `READY`, `RESETTING`
* `purge_time` (string) - A timestamp indicating when the branch is scheduled to be purged.
  Empty if the branch is not deleted, otherwise set to a timestamp in the future
* `source_branch` (string) - The name of the source branch from which this branch was created.
  Format: projects/{project_id}/branches/{branch_id}
* `source_branch_lsn` (string) - The Log Sequence Number (LSN) on the source branch from which this branch was created
* `source_branch_time` (string) - The point in time on the source branch from which this branch was created
* `source_snapshot` (string) - The snapshot this branch was restored from. Set only for branches created by
  restoring a snapshot; unset for all other branches.
  Format: projects/{project_id}/snapshots/{snapshot_id}
* `state_change_time` (string) - A timestamp indicating when the `current_state` began

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_postgres_branch.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_postgres_branch.this "name"
```