---
subcategory: "AI/BI"
---
# databricks_genie_agent Resource
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)

A Genie Agent is a no-code, AI/BI-powered analytics experience that lets business users ask questions of their data in natural language. Analysts configure a Genie Agent with a SQL warehouse and a `definition` attribute (the tables, instructions, and sample questions that ground the agent); business users then query it conversationally.

This resource manages the full lifecycle of a Genie Agent — create, update, and delete. The service returns `definition` exactly as it was applied, so the value never drifts; a definition the service would store differently is rejected with `INVALID_PARAMETER_VALUE`. In particular: write multi-line text (such as SQL or an instruction) as a single string, omit empty strings rather than setting them, list sample questions in `id` order, benchmark questions in descending `id` order, and MCP connections sorted by `mcp_connection`, and give every entry that has an `id` a unique 32-character lowercase hexadecimal value. Each column config must name a column that exists in its table, and a column config may set `display_name` only where column display names are enabled for the workspace; otherwise the apply is rejected with `INVALID_PARAMETER_VALUE`.

### Requirements

Genie uses data registered in Unity Catalog and requires at least CAN USE permission on a Pro or Serverless SQL warehouse. Databricks Assistant must be enabled for the workspace.

### Upgrading from an earlier provider version

Earlier versions of this resource took the definition as a JSON string in `serialized_genie_agent`. This version replaces it with the `definition` attribute. To get the definition of an existing Genie Agent, read it with the Genie Agent data source, or import the resource with `terraform import` and generate its configuration with `terraform plan -generate-config-out`. The format of `serialized_space` exported from the Genie UI or the Genie Spaces API is not accepted by `definition`; read the space through this resource or its data source instead, since every Genie space is also a Genie Agent.


## Example Usage
The following example creates a Genie Agent in a workspace folder, backed by a Pro or Serverless SQL warehouse. The `definition` attribute defines the tables, instructions, and sample questions that ground the agent.

```hcl
resource "databricks_genie_agent" "this" {
  display_name = "Sales Analytics Agent"
  description  = "Answers natural-language questions about sales data."
  warehouse_id = "abcd1234567890ef"
  parent_path  = "/Users/analyst@example.com/genie"

  definition = {
    config = {
      sample_questions = [
        {
          id       = "01f0a1b2c3d4e5f60718293a4b5c6d7e"
          question = "What were total sales last month?"
        },
      ]
    }
    data_sources = {
      tables = [
        {
          identifier  = "main.sales.orders"
          description = "One row per order, with order date, amount, and customer."
          column_configs = [
            {
              column                   = "region"
              enable_format_assistance = true
              enable_entity_matching   = true
            },
          ]
        },
        { identifier = "main.sales.customers" },
      ]
    }
    instructions = {
      text_instructions = [
        {
          id      = "01f0a1b2c3d4e5f60718293a4b5c6d7f"
          content = <<-EOT
            Report revenue as the sum of order_amount, rounded to 2 decimal places.
            Use fiscal quarters for quarterly questions.
          EOT
        },
      ]
      join_specs = [
        {
          id                = "01f0a1b2c3d4e5f60718293a4b5c6d80"
          left              = { identifier = "main.sales.orders", alias = "orders" }
          right             = { identifier = "main.sales.customers", alias = "customers" }
          sql               = "orders.customer_id = customers.customer_id"
          relationship_type = "GENIE_AGENT_JOIN_RELATIONSHIP_TYPE_MANY_TO_ONE"
        },
      ]
    }
  }
}
```


## Arguments
The following arguments are supported:
* `definition` (GenieAgentDefinition, required) - The definition of the Genie Agent: its sample questions, data sources, instructions, and
  benchmarks. Must set at least one field. Returned on get and on list with
  `GENIE_AGENT_VIEW_FULL`. Create and update return a submitted definition unchanged; one the
  service would store differently, or with a column config the workspace can't keep (unknown
  column, or a display name while those are disabled), is rejected with INVALID_PARAMETER_VALUE
* `display_name` (string, required) - Display name of the Genie Agent. Must be unique within its folder (effective_parent_path): a
  create, rename, or move that would collide with an existing Genie Agent's name in the target
  folder is rejected with ALREADY_EXISTS.
  
  The maximum length is 63 characters. Requests with a longer value are rejected with
  INVALID_PARAMETER_VALUE
* `warehouse_id` (string, required) - Warehouse associated with the Genie Agent
* `description` (string, optional) - Description of the Genie Agent.
  
  The maximum length is 4,096 characters. Requests with a longer value are rejected with
  INVALID_PARAMETER_VALUE
* `genie_agent_id` (string, optional) - The ID for the new Genie Agent: 32 lowercase hexadecimal characters (a UUID without dashes), not
  all zeros. Assigned by the server if omitted. Fails with ALREADY_EXISTS if the ID is in use in
  the workspace, including by a Genie Agent in the trash, or PERMISSION_DENIED if the caller
  cannot view the Genie Agent using it
* `parent_path` (string, optional) - The workspace folder to create the Genie Agent in. If omitted, the caller's home folder is
  used. Returned exactly as submitted; see `effective_parent_path` for the current folder. On
  update, setting it moves the Genie Agent and clearing it does not. At most 4,096 characters
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

### GenieAgentBenchmarkAnswer
* `content` (string, required) - The answer
* `format` (string, required) - The format of the answer. Possible values are: `GENIE_AGENT_BENCHMARK_ANSWER_FORMAT_SQL`

### GenieAgentBenchmarkQuestion
* `id` (string, required) - Identifier of the question: 32 lowercase hexadecimal characters, unique within the Genie
  Agent
* `question` (string, required) - The question
* `answer` (list of GenieAgentBenchmarkAnswer, optional) - The expected answer. At most one answer is supported; omit it when the question has no
  expected answer
* `evaluation_note` (string, optional) - Notes on how to evaluate answers to the question

### GenieAgentBenchmarks
* `questions` (list of GenieAgentBenchmarkQuestion, optional) - The benchmark questions, sorted by `id` in descending order

### GenieAgentColumnConfig
* `column` (string, required) - The name of the column. Must exist in the table or metric view
* `description` (string, optional) - Description of the column, overriding its Unity Catalog comment
* `display_name` (string, optional) - Display name of the column in the Genie Agent. Accepted only where column display names are
  enabled for the workspace
* `enable_entity_matching` (boolean, optional) - Whether the Genie Agent indexes the column's values to match them in questions
* `enable_format_assistance` (boolean, optional) - Whether the Genie Agent samples example values from the column to learn its format
* `exclude` (boolean, optional) - Whether the column is hidden from the Genie Agent
* `synonyms` (list of string, optional) - Synonyms for the column name

### GenieAgentConfig
* `sample_questions` (list of GenieAgentSampleQuestion, optional) - Sample questions shown to users of the Genie Agent, sorted by `id`

### GenieAgentDataSources
* `mcp_connections` (list of GenieAgentMcpConnection, optional) - Unity Catalog MCP connections, sorted by `mcp_connection`
* `tables` (list of GenieAgentTable, optional) - Unity Catalog tables, views, and metric views, in the order they were added
* `volumes` (list of GenieAgentVolume, optional) - Unity Catalog volumes, in the order they were added

### GenieAgentDefinition
* `benchmarks` (GenieAgentBenchmarks, optional) - Benchmarks for evaluating the quality of the Genie Agent's answers
* `config` (GenieAgentConfig, optional) - General settings of the Genie Agent
* `data_sources` (GenieAgentDataSources, optional) - The data the Genie Agent can access
* `instructions` (GenieAgentInstructions, optional) - Instructions and SQL that guide how the Genie Agent answers questions

### GenieAgentExampleQuestionSql
* `id` (string, required) - Identifier of the example: 32 lowercase hexadecimal characters, unique within the Genie Agent
* `question` (string, required) - The question
* `sql` (string, required) - The SQL that answers the question
* `parameters` (list of GenieAgentSqlParameter, optional) - Parameters used in the SQL, in the order given
* `usage_guidance` (string, optional) - Guidance on when to use this example

### GenieAgentInstructions
* `example_question_sqls` (list of GenieAgentExampleQuestionSql, optional) - Example questions paired with the SQL that answers them
* `join_specs` (list of GenieAgentJoinSpec, optional) - Predefined joins between tables or metric views
* `sql_functions` (list of GenieAgentSqlFunction, optional) - Unity Catalog SQL functions the generated SQL may call
* `sql_snippets` (GenieAgentSqlSnippets, optional) - Reusable SQL filters, expressions, and measures
* `text_instructions` (list of GenieAgentTextInstruction, optional) - General text instructions. At most one text instruction is supported

### GenieAgentJoinSource
* `identifier` (string, required) - The full three-level name of the table or metric view
* `alias` (string, optional) - The alias the join condition uses for this side

### GenieAgentJoinSpec
* `id` (string, required) - Identifier of the join: 32 lowercase hexadecimal characters, unique within the Genie Agent
* `left` (GenieAgentJoinSource, required) - The left side of the join
* `right` (GenieAgentJoinSource, required) - The right side of the join
* `sql` (string, required) - The join condition, for example `orders.customer_id = customers.id`
* `comment` (string, optional) - Description of the join's purpose
* `instruction` (string, optional) - Guidance on when to use the join
* `relationship_type` (string, optional) - How the rows of the two sides relate. Possible values are: `GENIE_AGENT_JOIN_RELATIONSHIP_TYPE_MANY_TO_MANY`, `GENIE_AGENT_JOIN_RELATIONSHIP_TYPE_MANY_TO_ONE`, `GENIE_AGENT_JOIN_RELATIONSHIP_TYPE_ONE_TO_MANY`, `GENIE_AGENT_JOIN_RELATIONSHIP_TYPE_ONE_TO_ONE`

### GenieAgentMcpConnection
* `mcp_connection` (string, required) - The name of the Unity Catalog connection, for example `system.ai.google_calendar`

### GenieAgentSampleQuestion
* `id` (string, required) - Identifier of the sample question: 32 lowercase hexadecimal characters, unique within the
  Genie Agent
* `question` (string, required) - The text of the question

### GenieAgentSqlExpression
* `id` (string, required) - Identifier of the expression: 32 lowercase hexadecimal characters, unique within the Genie
  Agent
* `sql` (string, required) - The SQL of the expression
* `alias` (string, optional) - The alias the Genie Agent uses for the expression in generated SQL
* `display_name` (string, optional) - Display name of the expression
* `instruction` (string, optional) - Guidance on when to use the expression
* `synonyms` (list of string, optional) - Synonyms for the expression

### GenieAgentSqlFilter
* `id` (string, required) - Identifier of the filter: 32 lowercase hexadecimal characters, unique within the Genie Agent
* `sql` (string, required) - The SQL of the filter
* `display_name` (string, optional) - Display name of the filter
* `instruction` (string, optional) - Guidance on when to use the filter
* `synonyms` (list of string, optional) - Synonyms for the filter

### GenieAgentSqlFunction
* `id` (string, required) - Identifier of the entry: 32 lowercase hexadecimal characters, unique within the Genie Agent
* `identifier` (string, required) - The full three-level name of the function, for example `catalog.schema.function`

### GenieAgentSqlMeasure
* `id` (string, required) - Identifier of the measure: 32 lowercase hexadecimal characters, unique within the Genie Agent
* `sql` (string, required) - The SQL of the measure
* `alias` (string, optional) - The alias the Genie Agent uses for the measure in generated SQL
* `display_name` (string, optional) - Display name of the measure
* `instruction` (string, optional) - Guidance on when to use the measure
* `synonyms` (list of string, optional) - Synonyms for the measure

### GenieAgentSqlParameter
* `name` (string, required) - The name of the parameter in the SQL
* `default_value` (GenieAgentSqlParameterDefaultValue, optional) - The parameter's default value
* `description` (string, optional) - Description of the parameter
* `type_hint` (string, optional) - The parameter's data type, for example `STRING` or `INTEGER`

### GenieAgentSqlParameterDefaultRange
* `max` (string, optional) - The upper end of the range
* `min` (string, optional) - The lower end of the range

### GenieAgentSqlParameterDefaultValue
* `value_range` (GenieAgentSqlParameterDefaultRange, optional) - The default range
* `values` (list of string, optional) - The default values. A single default is a one-element list

### GenieAgentSqlSnippets
* `expressions` (list of GenieAgentSqlExpression, optional) - Non-aggregated expressions (dimensions), such as `price * units`
* `filters` (list of GenieAgentSqlFilter, optional) - Filters, such as `state = 'CA'`
* `measures` (list of GenieAgentSqlMeasure, optional) - Aggregated measures, such as `SUM(price * units)`

### GenieAgentTable
* `identifier` (string, required) - The full three-level name, for example `catalog.schema.table`
* `column_configs` (list of GenieAgentColumnConfig, optional) - Column settings for the table, sorted by `column`
* `description` (string, optional) - Description of the table

### GenieAgentTextInstruction
* `content` (string, required) - The text of the instruction
* `id` (string, required) - Identifier of the instruction: 32 lowercase hexadecimal characters, unique within the Genie
  Agent

### GenieAgentVolume
* `path` (string, required) - The volume path, for example `/Volumes/catalog/schema/volume`. May name a directory or file
  within the volume
* `description` (string, optional) - Description of the volume

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - Time when the Genie Agent was created
* `effective_parent_path` (string) - The workspace folder that contains the Genie Agent
* `etag` (string) - Weak etag of `definition` only; metadata changes do not affect it. Send it on update or
  delete to detect concurrent definition changes
* `name` (string) - The resource name of the Genie Agent, in the form `agents/{genie_agent_id}`. Ignored on create
* `update_time` (string) - Time when the Genie Agent was last modified

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_genie_agent.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_genie_agent.this "name"
```