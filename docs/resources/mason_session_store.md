---
subcategory: "Agent Bricks"
---
# databricks_mason_session_store Resource
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `session_store_id` (string, required) - Caller-provided, workspace-unique session store ID. It must be 3-55 characters, begin with a
  lowercase letter, and contain only lowercase letters, digits, and hyphens
* `description` (string, optional) - Human-readable description of the session store
* `metadata` (object, optional) - Mutable caller-defined string labels
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - Time when the store was created
* `creator_user_id` (string) - Workspace-local user ID of the authenticated principal that created the store. This is
  immutable server-set attribution and does not grant access; authorization is evaluated from
  the authenticated request context
* `name` (string) - Resource name in the form `session-stores/{session_store_id}`
* `update_time` (string) - Time when the store was last updated

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_mason_session_store.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_mason_session_store.this "name"
```