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

// WorkspaceObjectStatus contains information when doing a get request or list request on the workspace api
type WorkspaceObjectStatus struct {
	ObjectID   int64      `json:"object_id,omitempty"`
	ObjectType ObjectType `json:"object_type,omitempty"`
	Path       string     `json:"path,omitempty"`
	Language   Language   `json:"language,omitempty"`
}

// NotebookContent contains the base64 content of the notebook
type NotebookContent struct {
	Content string `json:"content,omitempty"`
}

// NotebookImportRequest contains the payload to import a notebook
type NotebookImportRequest struct {
	Content   string       `json:"content,omitempty"`
	Path      string       `json:"path,omitempty"`
	Language  Language     `json:"language,omitempty"`
	Overwrite bool         `json:"overwrite,omitempty"`
	Format    ExportFormat `json:"format,omitempty"`
}

// NotebookDeleteRequest contains the payload to delete a notebook
type NotebookDeleteRequest struct {
	Path      string `json:"path,omitempty"`
	Recursive bool   `json:"recursive,omitempty"`
}
