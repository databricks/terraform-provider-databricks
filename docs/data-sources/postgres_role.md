---
subcategory: "Postgres"
---
# databricks_postgres_role Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

This data source retrieves a single Postgres role.


## Example Usage
### Retrieve Role by Name

```hcl
data "databricks_postgres_role" "this" {
  name = "projects/my-project/branches/main/roles/jane"
}

output "role_postgres_name" {
  value = data.databricks_postgres_role.this.status.postgres_role
}

output "role_identity_type" {
  value = data.databricks_postgres_role.this.status.identity_type
}

output "role_auth_method" {
  value = data.databricks_postgres_role.this.status.auth_method
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - Output only. The full resource path of the role.
  Format: projects/{project_id}/branches/{branch_id}/roles/{role_id}
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `create_time` (string)
* `name` (string) - Output only. The full resource path of the role.
  Format: projects/{project_id}/branches/{branch_id}/roles/{role_id}
* `parent` (string) - The Branch where this Role exists.
  Format: projects/{project_id}/branches/{branch_id}
* `spec` (RoleRoleSpec) - The spec contains the role configuration, including identity type, authentication method, and role attributes
* `status` (RoleRoleStatus) - Current status of the role, including its identity type, authentication method, and role attributes
* `update_time` (string)

### RoleAttributes
* `bypassrls` (boolean)
* `createdb` (boolean)
* `createrole` (boolean)

### RoleRoleSpec
* `attributes` (RoleAttributes) - The desired API-exposed Postgres role attribute to associate with the role. Optional
* `auth_method` (string) - If auth_method is left unspecified, a meaningful authentication method is derived from the identity_type:
  * For the managed identities, OAUTH is used.
  * For the regular postgres roles, authentication based on postgres passwords is used.
  
  NOTE: for the Databricks identity type GROUP, LAKEBASE_OAUTH_V1
  is the default auth method (group can login as well). Possible values are: `LAKEBASE_OAUTH_V1`, `NO_LOGIN`, `PG_PASSWORD_SCRAM_SHA_256`
* `identity_type` (string) - The type of role.
  When specifying a managed-identity, the chosen role_id must be a valid:
  
  * application ID for SERVICE_PRINCIPAL
  * user email for USER
  * group name for GROUP. Possible values are: `GROUP`, `SERVICE_PRINCIPAL`, `USER`
* `membership_roles` (list of string) - An enum value for a standard role that this role is a member of
* `postgres_role` (string) - The name of the Postgres role.
  
  This expects a valid Postgres identifier as specified in the link below.
  https://www.postgresql.org/docs/current/sql-syntax-lexical.html#SQL-SYNTAX-IDENTIFIERS
  
  Required when creating the Role.
  
  If you wish to create a Postgres Role backed by a managed Databricks identity, then postgres_role
  must be one of the following:
  
  1. user email for IdentityType.USER
  2. app ID for IdentityType.SERVICE_PRINCIPAL
  2. group name for IdentityType.GROUP

### RoleRoleStatus
* `attributes` (RoleAttributes) - The PG role attributes associated with the role
* `auth_method` (string) - Possible values are: `LAKEBASE_OAUTH_V1`, `NO_LOGIN`, `PG_PASSWORD_SCRAM_SHA_256`
* `identity_type` (string) - The type of the role. Possible values are: `GROUP`, `SERVICE_PRINCIPAL`, `USER`
* `membership_roles` (list of string) - An enum value for a standard role that this role is a member of
* `postgres_role` (string) - The name of the Postgres role
* `role_id` (string) - The short identifier of the role, suitable for showing to the users.
  For a role with name `projects/my-project/branches/my-branch/roles/my-role`,
  the role_id is `my-role`.
  
  Use this field when building UI components that display roles to users (e.g., a drop-down
  selector). Prefer showing `role_id` instead of the full resource name from `Role.name`,
  which follows the `projects/{project_id}/branches/{branch_id}/roles/{role_id}` format
  and is not user-friendly