---
subcategory: "Postgres"
---
# databricks_postgres_role Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


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
* `workspace_id` (string,required) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

### RoleAttributes
* `bypassrls` (boolean, optional)
* `createdb` (boolean, optional)
* `createrole` (boolean, optional)

### RoleRoleSpec
* `attributes` (RoleAttributes, optional) - The desired API-exposed Postgres role attribute to associate with the role. Optional
* `auth_method` (string, optional) - If auth_method is left unspecified, a meaningful authentication method is derived from the identity_type:
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