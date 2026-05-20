---
subcategory: "Postgres"
---
# databricks_postgres_role Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

### Lakebase Autoscaling Terraform Behavior

This resource uses Lakebase Autoscaling Terraform semantics. For complete details on how spec/status fields work, drift detection behavior, and state management requirements, see the `databricks_postgres_project` resource documentation.

### Overview

A Postgres role is an authentication and authorization principal inside a Postgres branch. A role can be a plain Postgres role (authenticating with a password) or can be backed by a Databricks managed identity (authenticating via OAuth).

### Hierarchy Context

Roles exist within the Lakebase Autoscaling resource hierarchy:
- A **role** belongs to a **branch** within a **project**
- A **role** can own one or more **databases** in the same branch
- A branch can contain multiple roles

### Identity Types

`spec.identity_type` controls whether the role is a plain Postgres role or is mapped to a Databricks managed identity:

- **Unset (default)**: Plain Postgres role. `spec.postgres_role` is a Postgres identifier of your choice.
- **`USER`**: Role is backed by a Databricks workspace user. `spec.postgres_role` must be the user's email (e.g., `jane@example.com`).
- **`SERVICE_PRINCIPAL`**: Role is backed by a Databricks service principal. `spec.postgres_role` must be the service principal's application ID.
- **`GROUP`**: Role is backed by a Databricks group. `spec.postgres_role` must be the group name. `GROUP` roles can log in directly via OAuth.

### Authentication Methods

`spec.auth_method` controls how the role authenticates to Postgres. If left unset, a reasonable default is inferred from `identity_type`:

- **`PG_PASSWORD_SCRAM_SHA_256`**: Default for plain Postgres roles (no identity type).
- **`LAKEBASE_OAUTH_V1`**: Default for managed identities (`USER`, `SERVICE_PRINCIPAL`, `GROUP`). The role logs in via Databricks OAuth.
- **`NO_LOGIN`**: The role cannot be used for interactive access. Useful for roles that only hold privileges and serve as group targets.

### Attributes

`spec.attributes` exposes a subset of Postgres role attributes:

- `createdb` — role can create databases
- `createrole` — role can create other roles
- `bypassrls` — role bypasses row-level security policies

See the [Postgres `CREATE ROLE` documentation](https://www.postgresql.org/docs/16/sql-createrole.html) for the authoritative semantics.

### Membership Roles

`spec.membership_roles` lets a role inherit privileges from Databricks-managed standard roles. The currently supported value is `DATABRICKS_SUPERUSER`, the highest set of privileges exposed to customers.

### Use Cases

- **Per-application service account**: Create a `SERVICE_PRINCIPAL`-backed role so an application authenticates to Postgres via OAuth without managing passwords
- **Human access**: Create a `USER`-backed role for each analyst or developer who needs direct SQL access
- **Team access**: Create a `GROUP`-backed role so a Databricks group can sign in collectively
- **Plain Postgres role**: Create a password-authenticated role for legacy workloads or tooling that cannot use OAuth


## Example Usage
### Role Backed by a Databricks User Identity

Create a role that is authenticated as a specific Databricks workspace user via OAuth. `auth_method` is left unset and defaults to `LAKEBASE_OAUTH_V1` for managed identities.

```hcl
resource "databricks_postgres_project" "this" {
  project_id = "my-project"
  spec = {
    pg_version   = 17
    display_name = "My Project"
  }
}

resource "databricks_postgres_branch" "main" {
  branch_id = "main"
  parent    = databricks_postgres_project.this.name
  spec = {
    no_expiry = true
  }
}

resource "databricks_postgres_role" "jane" {
  role_id = "jane"
  parent  = databricks_postgres_branch.main.name
  spec = {
    identity_type = "USER"
    postgres_role = "jane@databricks.com"   # Email of the user
  }
}
```

### Service Principal with `DATABRICKS_SUPERUSER` Membership

Create a role that is authenticated as a Databricks service principal via OAuth and grant it the highest customer-exposed privilege set via `DATABRICKS_SUPERUSER` membership.

```hcl
resource "databricks_postgres_role" "admin_sp" {
  role_id = "admin-sp"
  parent  = databricks_postgres_branch.main.name
  spec = {
    identity_type    = "SERVICE_PRINCIPAL"
    postgres_role    = "00000000-0000-0000-0000-000000000000" # application ID
    auth_method      = "LAKEBASE_OAUTH_V1"
    membership_roles = ["DATABRICKS_SUPERUSER"]
  }
}
```

### Multiple roles in a branch

By default, Terraform creates resources in parallel if the dependency graph allows. However, Lakebase
doesn't allow executing parallel manipulations inside a single branch. Only one of these resources can
be created or updated at a time:

- Role
- Database
- Endpoint

If you try to create resources in parallel, you'll see a conflict error like:

> Your project already has conflicting operations in progress. Please wait until they are complete, and then try again.

Terraform serializes execution automatically when one resource references another.
For example, when a database names a role as its owner via `spec.role`, Terraform creates the role before the database.
For resources that don't reference each other, like two sibling roles in the same branch, add `depends_on` so
Terraform knows to wait for creation of the first one to finish, before scheduling the creation of the second one.

For example:

```hcl
resource "databricks_postgres_role" "schema_owner" {
  role_id = "schemamigrator"
  parent  = databricks_postgres_branch.test.name  # previously created branch, omitted for compactness
  spec = {
    postgres_role = "schemamigrator"
    membership_roles = ["DATABRICKS_SUPERUSER"]
  }
}

resource "databricks_postgres_database" "application" {
  database_id = "application"
  parent      = databricks_postgres_branch.test.name
  spec = {
    postgres_database = "application"
    role              = databricks_postgres_role.schema_owner.name
  }
}

resource "databricks_postgres_role" "application" {
  role_id = "application"
  parent  = databricks_postgres_branch.test.name  # previously created branch, omitted for compactness
  spec = {
    postgres_role = "application"
  }
  
  depends_on = [ databricks_postgres_database.application ]
}
```

Note: in a real setup, the `application` role would also need `GRANT` privileges, but that's out of scope for this example.


## Arguments
The following arguments are supported:
* `parent` (string, required) - The Branch where this Role exists.
  Format: projects/{project_id}/branches/{branch_id}
* `role_id` (string, optional) - The ID to use for the Role, which will become the final component of
  the role's resource name.
  This ID becomes the role in Postgres.
  
  This value should be 4-63 characters, and valid characters
  are lowercase letters, numbers, and hyphens, as defined by RFC 1123.
  
  If role_id is not specified in the request, it is generated automatically
* `spec` (RoleRoleSpec, optional) - The spec contains the role configuration, including identity type, authentication method, and role attributes
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

### RoleAttributes
* `bypassrls` (boolean, optional)
* `createdb` (boolean, optional)
* `createrole` (boolean, optional)

### RoleRoleSpec
* `attributes` (RoleAttributes, optional) - The desired API-exposed Postgres role attribute to associate with the role. Optional
* `auth_method` (string, optional) - Controls how the Postgres role authenticates when a client opens a database
  connection. Supported values:
  
  * LAKEBASE_OAUTH_V1: the role authenticates by presenting a Databricks
  OAuth access token derived from the backing managed identity (the
  Databricks user, service principal, or group named by the role's
  `postgres_role`). No static password exists for roles using this method.
  * PG_PASSWORD_SCRAM_SHA_256: the role authenticates with a Postgres
  password verified server-side using the SCRAM-SHA-256 mechanism.
  Lakebase generates a password for the role.
  * NO_LOGIN: the role cannot open a Postgres session at all. Useful for
  roles that exist only to own objects or to aggregate privileges that
  are then granted to other, loginable roles.
  
  If auth_method is left unspecified, a meaningful authentication method is derived from the identity_type:
  * For the managed identities, OAUTH is used.
  * For the regular postgres roles, authentication based on postgres passwords is used.
  
  NOTE: for the Databricks identity type GROUP, LAKEBASE_OAUTH_V1
  is the default auth method (group can login as well). Possible values are: `LAKEBASE_OAUTH_V1`, `NO_LOGIN`, `PG_PASSWORD_SCRAM_SHA_256`
* `identity_type` (string, optional) - The type of role.
  When specifying a managed-identity, the chosen role_id must be a valid:
  
  * application ID for SERVICE_PRINCIPAL
  * user email for USER
  * group name for GROUP. Possible values are: `GROUP`, `SERVICE_PRINCIPAL`, `USER`
* `membership_roles` (list of string, optional) - An enum value for a standard role that this role is a member of
* `postgres_role` (string, optional) - The name of the Postgres role.
  
  This expects a valid Postgres identifier as specified in the link below.
  https://www.postgresql.org/docs/current/sql-syntax-lexical.html#SQL-SYNTAX-IDENTIFIERS
  
  Required when creating the Role.
  
  If you wish to create a Postgres Role backed by a managed Databricks identity, then postgres_role
  must be one of the following:
  
  1. user email for IdentityType.USER
  2. app ID for IdentityType.SERVICE_PRINCIPAL
  2. group name for IdentityType.GROUP

### RoleRoleStatus
* `attributes` (RoleAttributes, optional) - The PG role attributes associated with the role
* `auth_method` (string, optional) - Possible values are: `LAKEBASE_OAUTH_V1`, `NO_LOGIN`, `PG_PASSWORD_SCRAM_SHA_256`
* `identity_type` (string, optional) - The type of the role. Possible values are: `GROUP`, `SERVICE_PRINCIPAL`, `USER`
* `membership_roles` (list of string, optional) - An enum value for a standard role that this role is a member of
* `postgres_role` (string, optional) - The name of the Postgres role

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string)
* `name` (string) - Output only. The full resource path of the role.
  Format: projects/{project_id}/branches/{branch_id}/roles/{role_id}
* `status` (RoleRoleStatus) - Current status of the role, including its identity type, authentication method, and role attributes
* `update_time` (string)

### RoleRoleStatus
* `role_id` (string) - The short identifier of the role, suitable for showing to the users.
  For a role with name `projects/my-project/branches/my-branch/roles/my-role`,
  the role_id is `my-role`.
  
  Use this field when building UI components that display roles to users (e.g., a drop-down
  selector). Prefer showing `role_id` instead of the full resource name from `Role.name`,
  which follows the `projects/{project_id}/branches/{branch_id}/roles/{role_id}` format
  and is not user-friendly

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_postgres_role.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_postgres_role.this "name"
```