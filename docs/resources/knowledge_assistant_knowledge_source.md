---
subcategory: "Agent Bricks"
---
# databricks_knowledge_assistant_knowledge_source Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `description` (string, required) - Description of the knowledge source.
  Required when creating a Knowledge Source.
  When updating a Knowledge Source, optional unless included in update_mask
* `display_name` (string, required) - Human-readable display name of the knowledge source.
  Required when creating a Knowledge Source.
  When updating a Knowledge Source, optional unless included in update_mask
* `parent` (string, required) - Parent resource where this source will be created.
  Format: knowledge-assistants/{knowledge_assistant_id}
* `source_type` (string, required) - The type of the source: "index", "files", or "file_table".
  Required when creating a Knowledge Source.
  When updating a Knowledge Source, this field is ignored
* `file_table` (FileTableSpec, optional)
* `files` (FilesSpec, optional)
* `index` (IndexSpec, optional)
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

### FileTableSpec
* `file_col` (string, required) - The name of the column containing BINARY file content to be indexed
* `table_name` (string, required) - Full UC name of the table, in the format of {CATALOG}.{SCHEMA}.{TABLE_NAME}

### FilesSpec
* `path` (string, required) - A UC volume path that includes a list of files

### IndexSpec
* `doc_uri_col` (string, required) - The column that specifies a link or reference to where the information came from
* `index_name` (string, required) - Full UC name of the vector search index, in the format of {CATALOG}.{SCHEMA}.{INDEX_NAME}
* `text_col` (string, required) - The column that includes the document text for retrieval

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - Timestamp when this knowledge source was created
* `id` (string)
* `knowledge_cutoff_time` (string) - Timestamp representing the cutoff before which content in this knowledge source is being ingested
* `name` (string) - Full resource name:
  knowledge-assistants/{knowledge_assistant_id}/knowledge-sources/{knowledge_source_id}
* `state` (string) - Possible values are: `FAILED_UPDATE`, `UPDATED`, `UPDATING`

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_knowledge_assistant_knowledge_source.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_knowledge_assistant_knowledge_source.this "name"
```