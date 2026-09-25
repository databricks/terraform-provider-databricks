---
subcategory: "Postgres"
---
# databricks_postgres_database_credential Ephemeral Resource

[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

The `databricks_postgres_database_credential` ephemeral resource generates a short-lived OAuth credential for a Lakebase Postgres endpoint. The credential is available only during the current Terraform phase and is not stored in plan or state files.

This ephemeral resource requires Terraform 1.10 or later.

## Example Usage

The token can be supplied as the password for the [`cyrilgdn/postgresql`](https://registry.terraform.io/providers/cyrilgdn/postgresql/latest/docs) provider. The Databricks provider and the Postgres role named by `username` must represent the same Databricks user or service principal.

```hcl
data "databricks_postgres_endpoint" "this" {
  name = "projects/my-project/branches/production/endpoints/primary"
}

ephemeral "databricks_postgres_database_credential" "this" {
  endpoint = data.databricks_postgres_endpoint.this.name
}

provider "postgresql" {
  host             = data.databricks_postgres_endpoint.this.status.hosts.host
  port             = 5432
  database         = "application"
  username         = var.service_principal_application_id
  password         = ephemeral.databricks_postgres_database_credential.this.token
  sslmode          = "verify-full"
  superuser        = false
  expected_version = "17"
}
```

Lakebase credentials cannot be renewed without changing the token. A Terraform operation that opens a new Postgres connection after the credential expires can fail. Use the maximum one-hour lifetime for provider authentication and keep each Terraform operation within that lifetime.

## Arguments

The following arguments are supported:

* `endpoint` (string, required) - The full Lakebase endpoint resource name. Format: `projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}`.
* `ttl` (string, optional) - The requested credential lifetime as a Go duration. Must be between 5 minutes and 1 hour. Defaults to `1h`.

## Attributes

The following attributes are exported:

* `token` (string, sensitive) - The OAuth token to use as the Postgres password.
* `expire_time` (string) - The UTC timestamp when the credential expires.
