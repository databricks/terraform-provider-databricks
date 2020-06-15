package model

// Language is a custom type for language types in Databricks notebooks
type Language string

// ObjectType is a custom type for object types in Databricks workspaces
type ObjectType string

// ExportFormat is a custom type for formats in which you can export Databricks workspace components
type ExportFormat string

// Different types of export formats available on Databricks
const (
	Source  ExportFormat = "SOURCE"
	HTML    ExportFormat = "HTML"
	Jupyter ExportFormat = "JUPYTER"
	DBC     ExportFormat = "DBC"
)

// Different types of language formats available on Databricks
const (
	Scala  Language = "SCALA"
	Python Language = "PYTHON"
	SQL    Language = "SQL"
	R      Language = "R"
)

// Different types of export formats available on Databricks
const (
	Notebook      ObjectType = "NOTEBOOK"
	Directory     ObjectType = "DIRECTORY"
	LibraryObject ObjectType = "LIBRARY"
)

// NotebookInfo contains information when doing a get request or list request on the workspace api
type NotebookInfo struct {
	ObjectID   int64      `json:"object_id,omitempty"`
	ObjectType ObjectType `json:"object_type,omitempty"`
	Path       string     `json:"path,omitempty"`
	Language   Language   `json:"language,omitempty"`
}
