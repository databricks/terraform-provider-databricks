// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dataclassification

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/marshal"
)

// Auto-tagging configuration for a classification tag. When enabled, detected
// columns are automatically tagged with Unity Catalog tags.
type AutoTaggingConfig struct {
	// Whether auto-tagging is enabled or disabled for this classification tag.
	AutoTaggingMode AutoTaggingConfigAutoTaggingMode `json:"auto_tagging_mode"`
	// The Classification Tag. For built-in classes this is a system tag (e.g.,
	// "class.name", "class.location"); for custom classes it is a user-defined
	// governance tag key. The built-in tags are listed at
	// https://docs.databricks.com/aws/en/data-governance/unity-catalog/data-classification-tags.
	ClassificationTag string `json:"classification_tag"`
}

func (s *AutoTaggingConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Auto-tagging mode.
type AutoTaggingConfigAutoTaggingMode string

const AutoTaggingConfigAutoTaggingModeAutoTaggingDisabled AutoTaggingConfigAutoTaggingMode = `AUTO_TAGGING_DISABLED`

const AutoTaggingConfigAutoTaggingModeAutoTaggingEnabled AutoTaggingConfigAutoTaggingMode = `AUTO_TAGGING_ENABLED`

// String representation for [fmt.Print]
func (f *AutoTaggingConfigAutoTaggingMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AutoTaggingConfigAutoTaggingMode) Set(v string) error {
	switch v {
	case `AUTO_TAGGING_DISABLED`, `AUTO_TAGGING_ENABLED`:
		*f = AutoTaggingConfigAutoTaggingMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AUTO_TAGGING_DISABLED", "AUTO_TAGGING_ENABLED"`, v)
	}
}

// Values returns all possible values for AutoTaggingConfigAutoTaggingMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *AutoTaggingConfigAutoTaggingMode) Values() []AutoTaggingConfigAutoTaggingMode {
	return []AutoTaggingConfigAutoTaggingMode{
		AutoTaggingConfigAutoTaggingModeAutoTaggingDisabled,
		AutoTaggingConfigAutoTaggingModeAutoTaggingEnabled,
	}
}

// Type always returns AutoTaggingConfigAutoTaggingMode to satisfy [pflag.Value] interface
func (f *AutoTaggingConfigAutoTaggingMode) Type() string {
	return "AutoTaggingConfigAutoTaggingMode"
}

// Data Classification configuration for a Unity Catalog catalog. This message
// follows the "At Most One Resource" pattern: at most one CatalogConfig exists
// per catalog. - Full CRUD operations are supported: Create enables Data
// Classification, Delete disables it - It has no unique identifier of its own
// and uses its parent catalog's identifier (catalog_name)
type CatalogConfig struct {
	// Auto-tagging configurations that apply to this catalog. Each entry
	// overrides the metastore-level auto-tagging setting for its classification
	// tag; tags without an entry inherit the metastore-level setting.
	AutoTagConfigs []AutoTaggingConfig `json:"auto_tag_configs,omitempty"`
	// Schemas to exclude from the scan, each named relative to the parent
	// catalog. If specified, all schemas except the specified ones will be
	// scanned, including schemas created later. Mutually exclusive with
	// `included_schemas`: only one may be set per request. If neither
	// `included_schemas` nor `excluded_schemas` is set, all schemas are
	// scanned.
	ExcludedSchemas *CatalogConfigSchemaNames `json:"excluded_schemas,omitempty"`
	// Schemas to include in the scan, each named relative to the parent
	// catalog. If specified, only listed schemas will be scanned. Schemas
	// created later are not scanned until you add them to the list. Mutually
	// exclusive with `excluded_schemas`: only one may be set per request. If
	// neither `included_schemas` nor `excluded_schemas` is set, all schemas are
	// scanned.
	IncludedSchemas *CatalogConfigSchemaNames `json:"included_schemas,omitempty"`
	// Resource name in the format: catalogs/{catalog_name}/config.
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CatalogConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CatalogConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Wrapper message for a list of schema names.
type CatalogConfigSchemaNames struct {
	// Schema names, each relative to the parent catalog. Must not be empty.
	Names []string `json:"names"`
}

func (s *CatalogConfigSchemaNames) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateCatalogConfigRequest struct {
	// The configuration to create.
	CatalogConfig CatalogConfig `json:"catalog_config"`
	// Name of the catalog to enable, in the format catalogs/{catalog_name}.
	Parent string `json:"-" url:"-"`
}

func (s *CreateCatalogConfigRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteCatalogConfigRequest struct {
	// Resource name of the config, in the format
	// catalogs/{catalog_name}/config.
	Name string `json:"-" url:"-"`
}

func (s *DeleteCatalogConfigRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetCatalogConfigRequest struct {
	// Resource name of the config, in the format
	// catalogs/{catalog_name}/config.
	Name string `json:"-" url:"-"`
}

func (s *GetCatalogConfigRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Request to rescan a catalog that has Data Classification enabled.
type RescanCatalogRequest struct {
	// Resource name of the catalog, in the format: catalogs/{catalog_name}.
	Name string `json:"-" url:"-"`
}

func (s *RescanCatalogRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Response for RescanCatalog. Empty for now; a named message lets later
// versions add fields.
type RescanCatalogResponse struct {
}

func (s *RescanCatalogResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateCatalogConfigRequest struct {
	// The configuration to apply to the catalog. The name field in
	// catalog_config identifies which resource to update.
	CatalogConfig CatalogConfig `json:"catalog_config"`
	// Resource name in the format: catalogs/{catalog_name}/config.
	Name string `json:"-" url:"-"`
	// Comma-separated list of the fields of catalog_config to update, for
	// example `included_schemas` or `auto_tag_configs`.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateCatalogConfigRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}
