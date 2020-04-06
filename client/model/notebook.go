package model

//go:generate easytags $GOFILE

type Language string
type ObjectType string
type ExportFormat string

const (
	Source  ExportFormat = "SOURCE"
	Html    ExportFormat = "HTML"
	Jupyter ExportFormat = "JUPYTER"
	DBC     ExportFormat = "DBC"
)

const (
	Scala  Language = "SCALA"
	Python Language = "PYTHON"
	SQL    Language = "SQL"
	R      Language = "R"
)

const (
	Notebook  ObjectType = "NOTEBOOK"
	Directory ObjectType = "DIRECTORY"
	Library   ObjectType = "LIBRARY"
)

type NotebookInfo struct {
	ObjectId   int32      `json:"object_id,omitempty"`
	ObjectType ObjectType `json:"object_type,omitempty"`
	Path       string     `json:"path,omitempty"`
	Language   Language   `json:"language,omitempty"`
}
