// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package genieagents

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/marshal"
)

type CreateGenieAgentRequest struct {
	// The Genie Agent to create.
	GenieAgent GenieAgent `json:"genie_agent"`
	// The ID for the new Genie Agent: 32 lowercase hexadecimal characters (a
	// UUID without dashes), not all zeros. Assigned by the server if omitted.
	// Fails with ALREADY_EXISTS if the ID is in use in the workspace, including
	// by a Genie Agent in the trash, or PERMISSION_DENIED if the caller cannot
	// view the Genie Agent using it.
	GenieAgentId string `json:"-" url:"genie_agent_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateGenieAgentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateGenieAgentRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteGenieAgentRequest struct {
	// Etag from a previous read. On mismatch the delete fails with ABORTED and
	// nothing is deleted.
	Etag string `json:"-" url:"etag,omitempty"`
	// The resource name of the Genie Agent to delete, in the form
	// `agents/{genie_agent_id}`.
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteGenieAgentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteGenieAgentRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A Genie Agent resource.
type GenieAgent struct {
	// Time when the Genie Agent was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The definition of the Genie Agent: its sample questions, data sources,
	// instructions, and benchmarks. Create requires at least one table, volume,
	// MCP connection, or sub-agent in `data_sources`. Returned on get and on
	// list with `GENIE_AGENT_VIEW_FULL`. Create and update return a submitted
	// definition unchanged; one the service would store differently, or with a
	// column config the workspace can't keep (unknown column, or a display name
	// while those are disabled), is rejected with INVALID_PARAMETER_VALUE.
	Definition GenieAgentDefinition `json:"definition"`
	// Description of the Genie Agent.
	//
	// The maximum length is 4,096 characters. Requests with a longer value are
	// rejected with INVALID_PARAMETER_VALUE.
	Description string `json:"description,omitempty"`
	// Display name of the Genie Agent. Must be unique within its folder
	// (effective_parent_path): a create, rename, or move that would collide
	// with an existing Genie Agent's name in the target folder is rejected with
	// ALREADY_EXISTS.
	//
	// The maximum length is 63 characters. Requests with a longer value are
	// rejected with INVALID_PARAMETER_VALUE.
	DisplayName string `json:"display_name"`
	// The workspace folder that contains the Genie Agent.
	EffectiveParentPath string `json:"effective_parent_path,omitempty"`
	// Weak etag of `definition` only; metadata changes do not affect it. Send
	// it on update or delete to detect concurrent definition changes.
	Etag string `json:"etag,omitempty"`
	// The resource name of the Genie Agent, in the form
	// `agents/{genie_agent_id}`. Ignored on create.
	Name string `json:"name,omitempty"`
	// The workspace folder to create the Genie Agent in. If omitted, the
	// caller's home folder is used. Returned exactly as submitted; see
	// `effective_parent_path` for the current folder. On update, setting it
	// moves the Genie Agent and clearing it does not. At most 4,096 characters.
	ParentPath string `json:"parent_path,omitempty"`
	// Time when the Genie Agent was last modified.
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// Warehouse associated with the Genie Agent.
	WarehouseId string `json:"warehouse_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieAgent) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieAgent) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The expected answer to a benchmark question.
type GenieAgentBenchmarkAnswer struct {
	// The answer.
	Content string `json:"content"`
	// The format of the answer.
	Format GenieAgentBenchmarkAnswerFormat `json:"format"`
}

func (s *GenieAgentBenchmarkAnswer) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The format of a benchmark answer.
type GenieAgentBenchmarkAnswerFormat string

const GenieAgentBenchmarkAnswerFormatGenieAgentBenchmarkAnswerFormatSql GenieAgentBenchmarkAnswerFormat = `GENIE_AGENT_BENCHMARK_ANSWER_FORMAT_SQL`

// String representation for [fmt.Print]
func (f *GenieAgentBenchmarkAnswerFormat) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GenieAgentBenchmarkAnswerFormat) Set(v string) error {
	switch v {
	case `GENIE_AGENT_BENCHMARK_ANSWER_FORMAT_SQL`:
		*f = GenieAgentBenchmarkAnswerFormat(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "GENIE_AGENT_BENCHMARK_ANSWER_FORMAT_SQL"`, v)
	}
}

// Values returns all possible values for GenieAgentBenchmarkAnswerFormat.
//
// There is no guarantee on the order of the values in the slice.
func (f *GenieAgentBenchmarkAnswerFormat) Values() []GenieAgentBenchmarkAnswerFormat {
	return []GenieAgentBenchmarkAnswerFormat{
		GenieAgentBenchmarkAnswerFormatGenieAgentBenchmarkAnswerFormatSql,
	}
}

// Type always returns GenieAgentBenchmarkAnswerFormat to satisfy [pflag.Value] interface
func (f *GenieAgentBenchmarkAnswerFormat) Type() string {
	return "GenieAgentBenchmarkAnswerFormat"
}

// A benchmark question with its expected answer.
type GenieAgentBenchmarkQuestion struct {
	// The expected answer. At most one answer is supported; omit it when the
	// question has no expected answer.
	Answer []GenieAgentBenchmarkAnswer `json:"answer,omitempty"`
	// Notes on how to evaluate answers to the question.
	EvaluationNote string `json:"evaluation_note,omitempty"`
	// Identifier of the question: 32 lowercase hexadecimal characters, unique
	// within the Genie Agent.
	Id string `json:"id"`
	// The question.
	Question string `json:"question"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieAgentBenchmarkQuestion) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieAgentBenchmarkQuestion) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Benchmarks for evaluating a Genie Agent.
type GenieAgentBenchmarks struct {
	// The benchmark questions, sorted by `id` in descending order.
	Questions []GenieAgentBenchmarkQuestion `json:"questions,omitempty"`
}

func (s *GenieAgentBenchmarks) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Settings for one column of a table or metric view.
type GenieAgentColumnConfig struct {
	// The name of the column. Must exist in the table or metric view.
	Column string `json:"column"`
	// Description of the column, overriding its Unity Catalog comment.
	Description string `json:"description,omitempty"`
	// Display name of the column in the Genie Agent. Accepted only where column
	// display names are enabled for the workspace.
	DisplayName string `json:"display_name,omitempty"`
	// Whether the Genie Agent indexes the column's values to match them in
	// questions.
	EnableEntityMatching bool `json:"enable_entity_matching,omitempty"`
	// Whether the Genie Agent samples example values from the column to learn
	// its format.
	EnableFormatAssistance bool `json:"enable_format_assistance,omitempty"`
	// Whether the column is hidden from the Genie Agent.
	Exclude bool `json:"exclude,omitempty"`
	// Synonyms for the column name.
	Synonyms []string `json:"synonyms,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieAgentColumnConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieAgentColumnConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// General settings of a Genie Agent. The warehouse, display name, and
// description are separate fields of the Genie Agent.
type GenieAgentConfig struct {
	// Sample questions shown to users of the Genie Agent, sorted by `id`.
	SampleQuestions []GenieAgentSampleQuestion `json:"sample_questions,omitempty"`
}

func (s *GenieAgentConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The definition of a Genie Agent.
type GenieAgentDefinition struct {
	// Benchmarks for evaluating the quality of the Genie Agent's answers.
	Benchmarks *GenieAgentBenchmarks `json:"benchmarks,omitempty"`
	// General settings of the Genie Agent.
	Config *GenieAgentConfig `json:"config,omitempty"`
	// Instructions and SQL that guide how the Genie Agent answers questions.
	Instructions *GenieAgentInstructions `json:"instructions,omitempty"`
	// The resources the Genie Agent can use: data, SQL functions, connections,
	// and sub-agents.
	Resources *GenieAgentResources `json:"resources,omitempty"`
}

func (s *GenieAgentDefinition) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// An example question paired with the SQL that answers it.
type GenieAgentExampleQuestionSql struct {
	// Identifier of the example: 32 lowercase hexadecimal characters, unique
	// within the Genie Agent.
	Id string `json:"id"`
	// Parameters used in the SQL, in the order given.
	Parameters []GenieAgentSqlParameter `json:"parameters,omitempty"`
	// The question.
	Question string `json:"question"`
	// The SQL that answers the question.
	Sql string `json:"sql"`
	// Guidance on when to use this example.
	UsageGuidance string `json:"usage_guidance,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieAgentExampleQuestionSql) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieAgentExampleQuestionSql) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Instructions and SQL that guide how a Genie Agent answers questions.
type GenieAgentInstructions struct {
	// Example questions paired with the SQL that answers them.
	ExampleQuestionSqls []GenieAgentExampleQuestionSql `json:"example_question_sqls,omitempty"`
	// Predefined joins between tables or metric views.
	JoinSpecs []GenieAgentJoinSpec `json:"join_specs,omitempty"`
	// Reusable SQL filters, expressions, and measures.
	SqlSnippets *GenieAgentSqlSnippets `json:"sql_snippets,omitempty"`
	// The general text instruction.
	TextInstruction *GenieAgentTextInstruction `json:"text_instruction,omitempty"`
}

func (s *GenieAgentInstructions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// How the rows of the two sides of a join relate.
type GenieAgentJoinRelationshipType string

const GenieAgentJoinRelationshipTypeGenieAgentJoinRelationshipTypeManyToMany GenieAgentJoinRelationshipType = `GENIE_AGENT_JOIN_RELATIONSHIP_TYPE_MANY_TO_MANY`

const GenieAgentJoinRelationshipTypeGenieAgentJoinRelationshipTypeManyToOne GenieAgentJoinRelationshipType = `GENIE_AGENT_JOIN_RELATIONSHIP_TYPE_MANY_TO_ONE`

const GenieAgentJoinRelationshipTypeGenieAgentJoinRelationshipTypeOneToMany GenieAgentJoinRelationshipType = `GENIE_AGENT_JOIN_RELATIONSHIP_TYPE_ONE_TO_MANY`

const GenieAgentJoinRelationshipTypeGenieAgentJoinRelationshipTypeOneToOne GenieAgentJoinRelationshipType = `GENIE_AGENT_JOIN_RELATIONSHIP_TYPE_ONE_TO_ONE`

// String representation for [fmt.Print]
func (f *GenieAgentJoinRelationshipType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GenieAgentJoinRelationshipType) Set(v string) error {
	switch v {
	case `GENIE_AGENT_JOIN_RELATIONSHIP_TYPE_MANY_TO_MANY`, `GENIE_AGENT_JOIN_RELATIONSHIP_TYPE_MANY_TO_ONE`, `GENIE_AGENT_JOIN_RELATIONSHIP_TYPE_ONE_TO_MANY`, `GENIE_AGENT_JOIN_RELATIONSHIP_TYPE_ONE_TO_ONE`:
		*f = GenieAgentJoinRelationshipType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "GENIE_AGENT_JOIN_RELATIONSHIP_TYPE_MANY_TO_MANY", "GENIE_AGENT_JOIN_RELATIONSHIP_TYPE_MANY_TO_ONE", "GENIE_AGENT_JOIN_RELATIONSHIP_TYPE_ONE_TO_MANY", "GENIE_AGENT_JOIN_RELATIONSHIP_TYPE_ONE_TO_ONE"`, v)
	}
}

// Values returns all possible values for GenieAgentJoinRelationshipType.
//
// There is no guarantee on the order of the values in the slice.
func (f *GenieAgentJoinRelationshipType) Values() []GenieAgentJoinRelationshipType {
	return []GenieAgentJoinRelationshipType{
		GenieAgentJoinRelationshipTypeGenieAgentJoinRelationshipTypeManyToMany,
		GenieAgentJoinRelationshipTypeGenieAgentJoinRelationshipTypeManyToOne,
		GenieAgentJoinRelationshipTypeGenieAgentJoinRelationshipTypeOneToMany,
		GenieAgentJoinRelationshipTypeGenieAgentJoinRelationshipTypeOneToOne,
	}
}

// Type always returns GenieAgentJoinRelationshipType to satisfy [pflag.Value] interface
func (f *GenieAgentJoinRelationshipType) Type() string {
	return "GenieAgentJoinRelationshipType"
}

// One side of a join.
type GenieAgentJoinSource struct {
	// The alias the join condition uses for this side.
	Alias string `json:"alias,omitempty"`
	// The full three-level name of the table or metric view.
	Identifier string `json:"identifier"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieAgentJoinSource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieAgentJoinSource) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A predefined join between two tables or metric views.
type GenieAgentJoinSpec struct {
	// Description of the join's purpose.
	Comment string `json:"comment,omitempty"`
	// Identifier of the join: 32 lowercase hexadecimal characters, unique
	// within the Genie Agent.
	Id string `json:"id"`
	// Guidance on when to use the join.
	Instruction string `json:"instruction,omitempty"`
	// The left side of the join.
	Left GenieAgentJoinSource `json:"left"`
	// How the rows of the two sides relate.
	RelationshipType GenieAgentJoinRelationshipType `json:"relationship_type,omitempty"`
	// The right side of the join.
	Right GenieAgentJoinSource `json:"right"`
	// The join condition, for example `orders.customer_id = customers.id`.
	Sql string `json:"sql"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieAgentJoinSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieAgentJoinSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A Unity Catalog MCP connection the Genie Agent can use.
type GenieAgentMcpConnection struct {
	// The name of the Unity Catalog connection, for example
	// `system.ai.google_calendar`.
	McpConnection string `json:"mcp_connection"`
}

func (s *GenieAgentMcpConnection) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The resources a Genie Agent can use.
type GenieAgentResources struct {
	// Unity Catalog MCP connections, sorted by `mcp_connection`.
	McpConnections []GenieAgentMcpConnection `json:"mcp_connections,omitempty"`
	// Unity Catalog SQL functions the generated SQL may call, sorted by `id`,
	// then `identifier`.
	SqlFunctions []GenieAgentSqlFunction `json:"sql_functions,omitempty"`
	// Unity Catalog tables, views, and metric views, sorted by `identifier`.
	Tables []GenieAgentTable `json:"tables,omitempty"`
	// Unity Catalog volumes, sorted by `path`.
	Volumes []GenieAgentVolume `json:"volumes,omitempty"`
}

func (s *GenieAgentResources) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A sample question shown to users of the Genie Agent.
type GenieAgentSampleQuestion struct {
	// Identifier of the sample question: 32 lowercase hexadecimal characters,
	// unique within the Genie Agent.
	Id string `json:"id"`
	// The text of the question.
	Question string `json:"question"`
}

func (s *GenieAgentSampleQuestion) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A reusable non-aggregated SQL expression (dimension).
type GenieAgentSqlExpression struct {
	// The alias the Genie Agent uses for the expression in generated SQL.
	Alias string `json:"alias,omitempty"`
	// Display name of the expression.
	DisplayName string `json:"display_name,omitempty"`
	// Identifier of the expression: 32 lowercase hexadecimal characters, unique
	// within the Genie Agent.
	Id string `json:"id"`
	// Guidance on when to use the expression.
	Instruction string `json:"instruction,omitempty"`
	// The SQL of the expression.
	Sql string `json:"sql"`
	// Synonyms for the expression.
	Synonyms []string `json:"synonyms,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieAgentSqlExpression) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieAgentSqlExpression) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A reusable SQL filter.
type GenieAgentSqlFilter struct {
	// Display name of the filter.
	DisplayName string `json:"display_name,omitempty"`
	// Identifier of the filter: 32 lowercase hexadecimal characters, unique
	// within the Genie Agent.
	Id string `json:"id"`
	// Guidance on when to use the filter.
	Instruction string `json:"instruction,omitempty"`
	// The SQL of the filter.
	Sql string `json:"sql"`
	// Synonyms for the filter.
	Synonyms []string `json:"synonyms,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieAgentSqlFilter) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieAgentSqlFilter) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A Unity Catalog SQL function the generated SQL may call.
type GenieAgentSqlFunction struct {
	// Identifier of the entry: 32 lowercase hexadecimal characters, unique
	// within the Genie Agent.
	Id string `json:"id"`
	// The full three-level name of the function, for example
	// `catalog.schema.function`.
	Identifier string `json:"identifier"`
}

func (s *GenieAgentSqlFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A reusable aggregated SQL measure.
type GenieAgentSqlMeasure struct {
	// The alias the Genie Agent uses for the measure in generated SQL.
	Alias string `json:"alias,omitempty"`
	// Display name of the measure.
	DisplayName string `json:"display_name,omitempty"`
	// Identifier of the measure: 32 lowercase hexadecimal characters, unique
	// within the Genie Agent.
	Id string `json:"id"`
	// Guidance on when to use the measure.
	Instruction string `json:"instruction,omitempty"`
	// The SQL of the measure.
	Sql string `json:"sql"`
	// Synonyms for the measure.
	Synonyms []string `json:"synonyms,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieAgentSqlMeasure) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieAgentSqlMeasure) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A parameter of a parameterized SQL example.
type GenieAgentSqlParameter struct {
	// The parameter's default value.
	DefaultValue *GenieAgentSqlParameterDefaultValue `json:"default_value,omitempty"`
	// Description of the parameter.
	Description string `json:"description,omitempty"`
	// The name of the parameter in the SQL.
	Name string `json:"name"`
	// The parameter's data type, for example `STRING` or `INTEGER`.
	TypeHint string `json:"type_hint,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieAgentSqlParameter) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieAgentSqlParameter) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A range of default values for a SQL parameter.
type GenieAgentSqlParameterDefaultRange struct {
	// The upper end of the range.
	Max string `json:"max,omitempty"`
	// The lower end of the range.
	Min string `json:"min,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieAgentSqlParameterDefaultRange) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieAgentSqlParameterDefaultRange) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The default value of a SQL parameter: a list of values or a range.
type GenieAgentSqlParameterDefaultValue struct {
	// The default range.
	ValueRange *GenieAgentSqlParameterDefaultRange `json:"value_range,omitempty"`
	// The default values. A single default is a one-element list.
	Values []string `json:"values,omitempty"`
}

func (s *GenieAgentSqlParameterDefaultValue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Reusable SQL snippets.
type GenieAgentSqlSnippets struct {
	// Non-aggregated expressions (dimensions), such as `price * units`.
	Expressions []GenieAgentSqlExpression `json:"expressions,omitempty"`
	// Filters, such as `state = 'CA'`.
	Filters []GenieAgentSqlFilter `json:"filters,omitempty"`
	// Aggregated measures, such as `SUM(price * units)`.
	Measures []GenieAgentSqlMeasure `json:"measures,omitempty"`
}

func (s *GenieAgentSqlSnippets) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A Unity Catalog table, view, or metric view the Genie Agent can access.
type GenieAgentTable struct {
	// Column settings for the table, sorted by `column`.
	ColumnConfigs []GenieAgentColumnConfig `json:"column_configs,omitempty"`
	// Description of the table.
	Description string `json:"description,omitempty"`
	// The full three-level name, for example `catalog.schema.table`.
	Identifier string `json:"identifier"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieAgentTable) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieAgentTable) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A general text instruction for the Genie Agent.
type GenieAgentTextInstruction struct {
	// The text of the instruction.
	Content string `json:"content"`
	// Identifier of the instruction: 32 lowercase hexadecimal characters,
	// unique within the Genie Agent.
	Id string `json:"id"`
}

func (s *GenieAgentTextInstruction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The fields returned for each Genie Agent in a list response.
type GenieAgentView string

const GenieAgentViewGenieAgentViewBasic GenieAgentView = `GENIE_AGENT_VIEW_BASIC`

const GenieAgentViewGenieAgentViewFull GenieAgentView = `GENIE_AGENT_VIEW_FULL`

// String representation for [fmt.Print]
func (f *GenieAgentView) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GenieAgentView) Set(v string) error {
	switch v {
	case `GENIE_AGENT_VIEW_BASIC`, `GENIE_AGENT_VIEW_FULL`:
		*f = GenieAgentView(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "GENIE_AGENT_VIEW_BASIC", "GENIE_AGENT_VIEW_FULL"`, v)
	}
}

// Values returns all possible values for GenieAgentView.
//
// There is no guarantee on the order of the values in the slice.
func (f *GenieAgentView) Values() []GenieAgentView {
	return []GenieAgentView{
		GenieAgentViewGenieAgentViewBasic,
		GenieAgentViewGenieAgentViewFull,
	}
}

// Type always returns GenieAgentView to satisfy [pflag.Value] interface
func (f *GenieAgentView) Type() string {
	return "GenieAgentView"
}

// A Unity Catalog volume the Genie Agent can access.
type GenieAgentVolume struct {
	// Description of the volume.
	Description string `json:"description,omitempty"`
	// The volume path, for example `/Volumes/catalog/schema/volume`. May name a
	// directory or file within the volume.
	Path string `json:"path"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieAgentVolume) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieAgentVolume) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetGenieAgentRequest struct {
	// The resource name of the Genie Agent to retrieve, in the form
	// `agents/{genie_agent_id}`.
	Name string `json:"-" url:"-"`
}

func (s *GetGenieAgentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ListGenieAgentsRequest struct {
	// Maximum number of Genie Agents to return per page. Defaults to 20 when
	// unset or 0; values above 100 are coerced to 100. Negative values are
	// rejected with INVALID_PARAMETER_VALUE. With `GENIE_AGENT_VIEW_FULL`, a
	// smaller per-page maximum applies.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token from a previous response's `next_page_token`. A token
	// that was not returned by this method is rejected with
	// INVALID_PARAMETER_VALUE.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The fields to return for each Genie Agent. Defaults to
	// `GENIE_AGENT_VIEW_BASIC`. Use `GENIE_AGENT_VIEW_FULL` only when the agent
	// bodies are needed, since it returns smaller pages.
	View GenieAgentView `json:"-" url:"view,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListGenieAgentsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListGenieAgentsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response for a list request: a page of Genie Agents plus the token for the
// next page.
type ListGenieAgentsResponse struct {
	// The Genie Agents, with the fields selected by the request's `view`.
	GenieAgents []GenieAgent `json:"genie_agents,omitempty"`
	// Token to get the next page of results.
	NextPageToken string `json:"next_page_token,omitempty"`
	// Resource names, in the form `agents/{genie_agent_id}`, of Genie Agents in
	// this page that could not be read and are omitted from `genie_agents`.
	// Populated only for `GENIE_AGENT_VIEW_FULL`. To see why, get the Genie
	// Agent by name; to retry, repeat the request. Other kinds of resource
	// names may be added to this field in the future.
	Unreachable []string `json:"unreachable,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListGenieAgentsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListGenieAgentsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateGenieAgentRequest struct {
	// Etag from a previous read. On mismatch the update fails with ABORTED
	// (metadata changes may already be applied); get the agent and retry. May
	// return TEMPORARILY_UNAVAILABLE.
	Etag string `json:"-" url:"etag,omitempty"`
	// The Genie Agent to update. Its `name` identifies the resource; fields
	// named in update_mask are applied.
	GenieAgent GenieAgent `json:"genie_agent"`
	// The resource name of the Genie Agent, in the form
	// `agents/{genie_agent_id}`. Ignored on create.
	Name string `json:"-" url:"-"`
	// The set of fields to update. The mask is required and must name at least
	// one field; an empty or absent mask is rejected with
	// INVALID_PARAMETER_VALUE, as is a path that is not an updatable field.
	//
	// A field named in the mask takes the value from `genie_agent`, even when
	// it is unset there: an unset `description` is cleared; an unset
	// `display_name`, `warehouse_id`, or `definition` is rejected with
	// INVALID_PARAMETER_VALUE; an unset `parent_path` is cleared without moving
	// the Genie Agent. `definition` is replaced as a whole.
	//
	// The one supported path inside `definition` is
	// `definition.resources.sub_agents`, and only alongside `definition`.
	// Naming it makes the sub-agent list exact, so an empty list detaches every
	// sub-agent. Without it, including with `*`, an empty list keeps them.
	//
	// The special value `*` replaces the whole Genie Agent, as if every
	// updatable field were named. It must be the only path in the mask. Prefer
	// naming fields explicitly: with `*`, any field added to Genie Agent later
	// is also replaced, so a client unaware of it may clear it.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateGenieAgentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateGenieAgentRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
