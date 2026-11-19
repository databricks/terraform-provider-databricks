---
subcategory: "AI/BI"
---
# databricks_genie_agents Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)

Lists the Genie Agents in the workspace. Use this data source to enumerate existing Genie Agents — for example, to discover their resource names or drive other resources from the collection.


## Example Usage
The following example lists all Genie Agents in the workspace.

```hcl
data "databricks_genie_agents" "all" {}
```


## Arguments
The following arguments are supported:
* `page_size` (integer, optional) - Maximum number of Genie Agents to return per page. Defaults to 20 when unset or 0; values above
  100 are coerced to 100. Negative values are rejected with INVALID_PARAMETER_VALUE. With
  `GENIE_AGENT_VIEW_FULL`, a smaller per-page maximum applies
* `view` (string, optional) - The fields to return for each Genie Agent. Defaults to `GENIE_AGENT_VIEW_BASIC`. Use
  `GENIE_AGENT_VIEW_FULL` only when the agent bodies are needed, since it returns smaller pages. Possible values are: `GENIE_AGENT_VIEW_BASIC`, `GENIE_AGENT_VIEW_FULL`
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `genie_agents`. It is a list of resources, each with the following attributes:
* `create_time` (string) - Time when the Genie Agent was created
* `definition` (GenieAgentDefinition) - The definition of the Genie Agent: its sample questions, data sources, instructions, and
  benchmarks. Must set at least one field. Returned on get and on list with
  `GENIE_AGENT_VIEW_FULL`. Create and update return a submitted definition unchanged; one the
  service would store differently, or with a column config the workspace can't keep (unknown
  column, or a display name while those are disabled), is rejected with INVALID_PARAMETER_VALUE
* `description` (string) - Description of the Genie Agent.
  
  The maximum length is 4,096 characters. Requests with a longer value are rejected with
  INVALID_PARAMETER_VALUE
* `display_name` (string) - Display name of the Genie Agent. Must be unique within its folder (effective_parent_path): a
  create, rename, or move that would collide with an existing Genie Agent's name in the target
  folder is rejected with ALREADY_EXISTS.
  
  The maximum length is 63 characters. Requests with a longer value are rejected with
  INVALID_PARAMETER_VALUE
* `effective_parent_path` (string) - The workspace folder that contains the Genie Agent
* `etag` (string) - Weak etag of `definition` only; metadata changes do not affect it. Send it on update or
  delete to detect concurrent definition changes
* `name` (string) - The resource name of the Genie Agent, in the form `agents/{genie_agent_id}`. Ignored on create
* `parent_path` (string) - The workspace folder to create the Genie Agent in. If omitted, the caller's home folder is
  used. Returned exactly as submitted; see `effective_parent_path` for the current folder. On
  update, setting it moves the Genie Agent and clearing it does not. At most 4,096 characters
* `update_time` (string) - Time when the Genie Agent was last modified
* `warehouse_id` (string) - Warehouse associated with the Genie Agent

### GenieAgentBenchmarkAnswer
* `content` (string) - The answer
* `format` (string) - The format of the answer. Possible values are: `GENIE_AGENT_BENCHMARK_ANSWER_FORMAT_SQL`

### GenieAgentBenchmarkQuestion
* `answer` (list of GenieAgentBenchmarkAnswer) - The expected answer. At most one answer is supported; omit it when the question has no
  expected answer
* `evaluation_note` (string) - Notes on how to evaluate answers to the question
* `id` (string) - Identifier of the question: 32 lowercase hexadecimal characters, unique within the Genie
  Agent
* `question` (string) - The question

### GenieAgentBenchmarks
* `questions` (list of GenieAgentBenchmarkQuestion) - The benchmark questions, sorted by `id` in descending order

### GenieAgentColumnConfig
* `column` (string) - The name of the column. Must exist in the table or metric view
* `description` (string) - Description of the column, overriding its Unity Catalog comment
* `display_name` (string) - Display name of the column in the Genie Agent. Accepted only where column display names are
  enabled for the workspace
* `enable_entity_matching` (boolean) - Whether the Genie Agent indexes the column's values to match them in questions
* `enable_format_assistance` (boolean) - Whether the Genie Agent samples example values from the column to learn its format
* `exclude` (boolean) - Whether the column is hidden from the Genie Agent
* `synonyms` (list of string) - Synonyms for the column name

### GenieAgentConfig
* `sample_questions` (list of GenieAgentSampleQuestion) - Sample questions shown to users of the Genie Agent, sorted by `id`

### GenieAgentDataSources
* `mcp_connections` (list of GenieAgentMcpConnection) - Unity Catalog MCP connections, sorted by `mcp_connection`
* `tables` (list of GenieAgentTable) - Unity Catalog tables, views, and metric views, in the order they were added
* `volumes` (list of GenieAgentVolume) - Unity Catalog volumes, in the order they were added

### GenieAgentDefinition
* `benchmarks` (GenieAgentBenchmarks) - Benchmarks for evaluating the quality of the Genie Agent's answers
* `config` (GenieAgentConfig) - General settings of the Genie Agent
* `data_sources` (GenieAgentDataSources) - The data the Genie Agent can access
* `instructions` (GenieAgentInstructions) - Instructions and SQL that guide how the Genie Agent answers questions

### GenieAgentExampleQuestionSql
* `id` (string) - Identifier of the example: 32 lowercase hexadecimal characters, unique within the Genie Agent
* `parameters` (list of GenieAgentSqlParameter) - Parameters used in the SQL, in the order given
* `question` (string) - The question
* `sql` (string) - The SQL that answers the question
* `usage_guidance` (string) - Guidance on when to use this example

### GenieAgentInstructions
* `example_question_sqls` (list of GenieAgentExampleQuestionSql) - Example questions paired with the SQL that answers them
* `join_specs` (list of GenieAgentJoinSpec) - Predefined joins between tables or metric views
* `sql_functions` (list of GenieAgentSqlFunction) - Unity Catalog SQL functions the generated SQL may call
* `sql_snippets` (GenieAgentSqlSnippets) - Reusable SQL filters, expressions, and measures
* `text_instructions` (list of GenieAgentTextInstruction) - General text instructions. At most one text instruction is supported

### GenieAgentJoinSource
* `alias` (string) - The alias the join condition uses for this side
* `identifier` (string) - The full three-level name of the table or metric view

### GenieAgentJoinSpec
* `comment` (string) - Description of the join's purpose
* `id` (string) - Identifier of the join: 32 lowercase hexadecimal characters, unique within the Genie Agent
* `instruction` (string) - Guidance on when to use the join
* `left` (GenieAgentJoinSource) - The left side of the join
* `relationship_type` (string) - How the rows of the two sides relate. Possible values are: `GENIE_AGENT_JOIN_RELATIONSHIP_TYPE_MANY_TO_MANY`, `GENIE_AGENT_JOIN_RELATIONSHIP_TYPE_MANY_TO_ONE`, `GENIE_AGENT_JOIN_RELATIONSHIP_TYPE_ONE_TO_MANY`, `GENIE_AGENT_JOIN_RELATIONSHIP_TYPE_ONE_TO_ONE`
* `right` (GenieAgentJoinSource) - The right side of the join
* `sql` (string) - The join condition, for example `orders.customer_id = customers.id`

### GenieAgentMcpConnection
* `mcp_connection` (string) - The name of the Unity Catalog connection, for example `system.ai.google_calendar`

### GenieAgentSampleQuestion
* `id` (string) - Identifier of the sample question: 32 lowercase hexadecimal characters, unique within the
  Genie Agent
* `question` (string) - The text of the question

### GenieAgentSqlExpression
* `alias` (string) - The alias the Genie Agent uses for the expression in generated SQL
* `display_name` (string) - Display name of the expression
* `id` (string) - Identifier of the expression: 32 lowercase hexadecimal characters, unique within the Genie
  Agent
* `instruction` (string) - Guidance on when to use the expression
* `sql` (string) - The SQL of the expression
* `synonyms` (list of string) - Synonyms for the expression

### GenieAgentSqlFilter
* `display_name` (string) - Display name of the filter
* `id` (string) - Identifier of the filter: 32 lowercase hexadecimal characters, unique within the Genie Agent
* `instruction` (string) - Guidance on when to use the filter
* `sql` (string) - The SQL of the filter
* `synonyms` (list of string) - Synonyms for the filter

### GenieAgentSqlFunction
* `id` (string) - Identifier of the entry: 32 lowercase hexadecimal characters, unique within the Genie Agent
* `identifier` (string) - The full three-level name of the function, for example `catalog.schema.function`

### GenieAgentSqlMeasure
* `alias` (string) - The alias the Genie Agent uses for the measure in generated SQL
* `display_name` (string) - Display name of the measure
* `id` (string) - Identifier of the measure: 32 lowercase hexadecimal characters, unique within the Genie Agent
* `instruction` (string) - Guidance on when to use the measure
* `sql` (string) - The SQL of the measure
* `synonyms` (list of string) - Synonyms for the measure

### GenieAgentSqlParameter
* `default_value` (GenieAgentSqlParameterDefaultValue) - The parameter's default value
* `description` (string) - Description of the parameter
* `name` (string) - The name of the parameter in the SQL
* `type_hint` (string) - The parameter's data type, for example `STRING` or `INTEGER`

### GenieAgentSqlParameterDefaultRange
* `max` (string) - The upper end of the range
* `min` (string) - The lower end of the range

### GenieAgentSqlParameterDefaultValue
* `value_range` (GenieAgentSqlParameterDefaultRange) - The default range
* `values` (list of string) - The default values. A single default is a one-element list

### GenieAgentSqlSnippets
* `expressions` (list of GenieAgentSqlExpression) - Non-aggregated expressions (dimensions), such as `price * units`
* `filters` (list of GenieAgentSqlFilter) - Filters, such as `state = 'CA'`
* `measures` (list of GenieAgentSqlMeasure) - Aggregated measures, such as `SUM(price * units)`

### GenieAgentTable
* `column_configs` (list of GenieAgentColumnConfig) - Column settings for the table, sorted by `column`
* `description` (string) - Description of the table
* `identifier` (string) - The full three-level name, for example `catalog.schema.table`

### GenieAgentTextInstruction
* `content` (string) - The text of the instruction
* `id` (string) - Identifier of the instruction: 32 lowercase hexadecimal characters, unique within the Genie
  Agent

### GenieAgentVolume
* `description` (string) - Description of the volume
* `path` (string) - The volume path, for example `/Volumes/catalog/schema/volume`. May name a directory or file
  within the volume