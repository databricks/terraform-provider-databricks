---
subcategory: "Agent Bricks"
---
# databricks_knowledge_assistant_knowledge_sources Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `parent` (string, required) - Parent resource to list from.
  Format: knowledge-assistants/{knowledge_assistant_id}
* `page_size` (integer, optional)
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `knowledge_sources`. It is a list of resources, each with the following attributes:
* `create_time` (string) - Timestamp when this knowledge source was created
* `description` (string) - Description of the knowledge source.
  Required when creating a Knowledge Source.
  When updating a Knowledge Source, optional unless included in update_mask
* `display_name` (string) - Human-readable display name of the knowledge source.
  Required when creating a Knowledge Source.
  When updating a Knowledge Source, optional unless included in update_mask
* `file_table` (FileTableSpec)
* `files` (FilesSpec)
* `id` (string)
* `index` (IndexSpec)
* `knowledge_cutoff_time` (string) - Timestamp representing the cutoff before which content in this knowledge source is being ingested
* `name` (string) - Full resource name:
  knowledge-assistants/{knowledge_assistant_id}/knowledge-sources/{knowledge_source_id}
* `source_type` (string) - The type of the source: "index", "files", or "file_table".
  Required when creating a Knowledge Source.
  When updating a Knowledge Source, this field is ignored
* `state` (string) - Possible values are: `FAILED_UPDATE`, `UPDATED`, `UPDATING`

### FileTableSpec
* `file_col` (string) - The name of the column containing BINARY file content to be indexed
* `table_name` (string) - Full UC name of the table, in the format of {CATALOG}.{SCHEMA}.{TABLE_NAME}

### FilesSpec
* `path` (string) - A UC volume path that includes a list of files

### IndexSpec
* `doc_uri_col` (string) - The column that specifies a link or reference to where the information came from
* `index_name` (string) - Full UC name of the vector search index, in the format of {CATALOG}.{SCHEMA}.{INDEX_NAME}
* `text_col` (string) - The column that includes the document text for retrieval