// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package softwareartifacts

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/marshal"
)

type DeleteSoftwareArtifactRequest struct {
	// Force deletion even if the software artifact has versions, deleting them
	// with it. Default is false.
	Force bool `json:"-" url:"force,omitempty"`
	// Resource name of the software artifact, of the form
	// `software-artifacts/{catalog}.{schema}.{artifact}`.
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteSoftwareArtifactRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteSoftwareArtifactRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteSoftwareArtifactVersionRequest struct {
	// Resource name of the version, of the form
	// `software-artifacts/{catalog}.{schema}.{artifact}/versions/{version}`,
	// where `{version}` is the version's content digest, for example
	// `sha256:<hex>`.
	Name string `json:"-" url:"-"`
}

func (s *DeleteSoftwareArtifactVersionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetSoftwareArtifactRequest struct {
	// Resource name of the artifact, of the form
	// `software-artifacts/{catalog}.{schema}.{artifact}`.
	Name string `json:"-" url:"-"`
}

func (s *GetSoftwareArtifactRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetSoftwareArtifactVersionRequest struct {
	// Resource name of the version, of the form
	// `software-artifacts/{catalog}.{schema}.{artifact}/versions/{version}`,
	// where `{version}` is the version's content digest.
	Name string `json:"-" url:"-"`
}

func (s *GetSoftwareArtifactVersionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Image acceleration information for a software artifact version.
type ImageAccelerationDetail struct {
	// Image acceleration status for the version.
	Status ImageAccelerationDetailStatus `json:"status,omitempty"`
}

func (s *ImageAccelerationDetail) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Image acceleration status for a software artifact version.
type ImageAccelerationDetailStatus string

const ImageAccelerationDetailStatusSuccess ImageAccelerationDetailStatus = `SUCCESS`

const ImageAccelerationDetailStatusUnspecified ImageAccelerationDetailStatus = `UNSPECIFIED`

// String representation for [fmt.Print]
func (f *ImageAccelerationDetailStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ImageAccelerationDetailStatus) Set(v string) error {
	switch v {
	case `SUCCESS`, `UNSPECIFIED`:
		*f = ImageAccelerationDetailStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "SUCCESS", "UNSPECIFIED"`, v)
	}
}

// Values returns all possible values for ImageAccelerationDetailStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *ImageAccelerationDetailStatus) Values() []ImageAccelerationDetailStatus {
	return []ImageAccelerationDetailStatus{
		ImageAccelerationDetailStatusSuccess,
		ImageAccelerationDetailStatusUnspecified,
	}
}

// Type always returns ImageAccelerationDetailStatus to satisfy [pflag.Value] interface
func (f *ImageAccelerationDetailStatus) Type() string {
	return "ImageAccelerationDetailStatus"
}

type ListSoftwareArtifactVersionsRequest struct {
	// Maximum number of versions to return. If omitted or set to 0, the page
	// size defaults to 50. The maximum value is 50; values above 50 or below 0
	// are rejected.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Opaque continuation token from a previous response's next_page_token.
	// Omit for the first page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Resource name of the parent artifact, of the form
	// `software-artifacts/{catalog}.{schema}.{artifact}`.
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSoftwareArtifactVersionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSoftwareArtifactVersionsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response containing a page of software artifact versions.
type ListSoftwareArtifactVersionsResponse struct {
	// Opaque token to pass as the next request's page_token. Empty when there
	// are no more pages.
	NextPageToken string `json:"next_page_token,omitempty"`
	// The versions of the requested software artifact.
	SoftwareArtifactVersions []SoftwareArtifactVersion `json:"software_artifact_versions,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSoftwareArtifactVersionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSoftwareArtifactVersionsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSoftwareArtifactsRequest struct {
	// Maximum number of software artifacts to return. If omitted or set to 0,
	// the page size defaults to 50. The maximum value is 50; values above 50 or
	// below 0 are rejected.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Opaque continuation token from a previous response's next_page_token.
	// Omit for the first page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Resource name of the parent schema, of the form
	// `schemas/{catalog}.{schema}`.
	Parent string `json:"-" url:"parent"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSoftwareArtifactsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSoftwareArtifactsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response containing a page of software artifacts.
type ListSoftwareArtifactsResponse struct {
	// Opaque token to pass as the next request's page_token. Empty when there
	// are no more pages.
	NextPageToken string `json:"next_page_token,omitempty"`
	// The software artifacts under the requested schema.
	SoftwareArtifacts []SoftwareArtifact `json:"software_artifacts,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSoftwareArtifactsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSoftwareArtifactsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Request to resolve a software artifact version from its tag alias.
type ResolveSoftwareArtifactVersionRequest struct {
	// Image tag alias selecting the version. The alias must be a valid OCI
	// image tag of at most 128 characters, made of letters, digits,
	// underscores, periods, and hyphens, and must not start with a period or
	// hyphen.
	Alias string `json:"alias"`
	// Resource name of the software artifact containing the alias, of the form
	// `software-artifacts/{catalog}.{schema}.{artifact}`.
	Parent string `json:"-" url:"-"`
}

func (s *ResolveSoftwareArtifactVersionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Response containing the resolved software artifact version.
type ResolveSoftwareArtifactVersionResponse struct {
	// The software artifact version selected by the alias.
	SoftwareArtifactVersion *SoftwareArtifactVersion `json:"software_artifact_version,omitempty"`
}

func (s *ResolveSoftwareArtifactVersionResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A software artifact is a named container image stored in a Unity Catalog
// schema and governed by Unity Catalog privileges. Images are pushed through
// Databricks Artifact Registry; each pushed image digest becomes an immutable
// version of the artifact, exposed as a SoftwareArtifactVersion. See
// https://docs.databricks.com/aws/en/artifact-registry/ for concepts, naming
// rules, and the privileges each operation requires.
type SoftwareArtifact struct {
	// Time at which the artifact was created, formatted as an RFC 3339
	// timestamp in JSON.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Principal name of the software artifact creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Resource name of the artifact, of the form
	// `software-artifacts/{catalog}.{schema}.{artifact}`.
	Name string `json:"name,omitempty"`
	// Principal that owns the artifact, given as a user's email address, a
	// service principal's application ID, or a group name.
	Owner string `json:"owner,omitempty"`
	// Time at which the artifact was last updated, formatted as an RFC 3339
	// timestamp in JSON.
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// Principal name that last updated this software artifact.
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SoftwareArtifact) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SoftwareArtifact) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// An immutable version of a software artifact, identified by its content
// digest.
type SoftwareArtifactVersion struct {
	// Time at which this software artifact version was created, formatted as an
	// RFC 3339 timestamp in JSON.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Principal name of the software artifact version creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Content digest of the version's image manifest, of the form
	// `{algorithm}:{hex}`.
	Digest string `json:"digest,omitempty"`
	// Image acceleration information for this version.
	ImageAccelerationDetails *ImageAccelerationDetail `json:"image_acceleration_details,omitempty"`
	// Resource name of the version, of the form
	// `software-artifacts/{catalog}.{schema}.{artifact}/versions/{version}`,
	// where `{version}` is the version's content digest.
	Name string `json:"name,omitempty"`
	// Tags (aliases) that currently point at this version, for example `latest`
	// or `v1.2.3`. Bounded at 10 per version; a version pointed at by more tags
	// returns an arbitrary 10 of them.
	Tags []string `json:"tags,omitempty"`
	// Time at which this software artifact version was last updated, formatted
	// as an RFC 3339 timestamp in JSON.
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// Principal name that last updated this software artifact version.
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SoftwareArtifactVersion) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SoftwareArtifactVersion) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateSoftwareArtifactRequest struct {
	// Resource name of the artifact, of the form
	// `software-artifacts/{catalog}.{schema}.{artifact}`.
	Name string `json:"-" url:"-"`
	// The software artifact to update. Its name identifies the existing
	// artifact; only owner may be changed.
	SoftwareArtifact SoftwareArtifact `json:"software_artifact"`
	// Fields to update, relative to `software_artifact`. Accepts `owner` or
	// `*`; owner is the only mutable field, so both select the same update.
	// Output-only values in the resource are ignored.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateSoftwareArtifactRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}
