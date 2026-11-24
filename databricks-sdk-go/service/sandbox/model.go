// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sandbox

import (
	"fmt"
	"io"

	"github.com/databricks/databricks-sdk-go/common/types/duration"
	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/marshal"
)

// A single command execution.
type Command struct {
	// Arguments passed to the program.
	Args []string `json:"args,omitempty"`
	// The program that was executed.
	Cmd string `json:"cmd,omitempty"`
	// Stable identifier for this command.
	CommandId string `json:"command_id,omitempty"`
	// Process exit code. Only present when finished is true and the process
	// exited normally (not killed by signal or failed to start).
	ExitCode int `json:"exit_code,omitempty"`
	// Whether the command has finished executing.
	Finished bool `json:"finished,omitempty"`
	// PID of the spawned process. Absent if the process failed to start.
	Pid int64 `json:"pid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Command) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Command) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ComputeSpec struct {
	// Idle time after which the sandbox is automatically stopped. Unset or 0
	// disables inactivity auto-stop. Negative values are rejected. Can be
	// changed on an existing sandbox with UpdateSandbox.
	InactivityTimeout *duration.Duration `json:"inactivity_timeout,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ComputeSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ComputeSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Selects a sandbox directory to create and whether to create its parents.
type CreateDirectoryRequest struct {
	Parent string `json:"-" url:"-"`
	// Defaults to true, which also succeeds when the directory already exists.
	Parents bool `json:"parents,omitempty"`
	// Absolute UTF-8 path of the directory.
	Path string `json:"path"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateDirectoryRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateDirectoryRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Confirms that the requested directory exists.
type CreateDirectoryResponse struct {
}

func (s *CreateDirectoryResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateSandboxRequest struct {
	// The sandbox to create. An empty body creates a sandbox with the default
	// configuration.
	Sandbox Sandbox `json:"sandbox"`
	// Client-supplied ID that becomes the final path segment of the resource
	// name. It must start with a lowercase letter, contain only lowercase
	// letters, digits and hyphens, end with a letter or digit, and be at most
	// 63 characters long. Values shaped like a UUID are rejected.
	SandboxId string `json:"-" url:"sandbox_id"`
}

func (s *CreateSandboxRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Selects a sandbox path to delete and whether to delete its contents.
type DeletePathRequest struct {
	Parent string `json:"-" url:"-"`
	// Absolute UTF-8 path of the entry to delete.
	Path string `json:"path"`
	// Required for a non-empty directory.
	Recursive bool `json:"recursive,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeletePathRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeletePathRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Confirms that the requested path is absent.
type DeletePathResponse struct {
}

func (s *DeletePathResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteSandboxRequest struct {
	// Resource name of the sandbox to delete, in the form
	// `sandboxes/{sandbox_id}`.
	Name string `json:"-" url:"-"`
}

func (s *DeleteSandboxRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DownloadFileRequest struct {
	// A single byte range, such as bytes=0-1023 or bytes=-1024. Equivalent to
	// the Range header.
	ByteRange string `json:"-" url:"byte_range,omitempty"`
	// Entity tags or *. A mismatch returns 412 Precondition Failed. Equivalent
	// to the If-Match header.
	IfMatch string `json:"-" url:"if_match,omitempty"`
	// An HTTP-date. Returns 304 if the file has not changed since this time.
	// Ignored when If-None-Match is present. Equivalent to the
	// If-Modified-Since header.
	IfModifiedSince string `json:"-" url:"if_modified_since,omitempty"`
	// Entity tags or *. A match returns 304 Not Modified. Equivalent to the
	// If-None-Match header.
	IfNoneMatch string `json:"-" url:"if_none_match,omitempty"`
	// A strong entity tag. A mismatch, weak entity tag, or HTTP-date causes a
	// full download. Equivalent to the If-Range header.
	IfRange string `json:"-" url:"if_range,omitempty"`
	// An HTTP-date. Returns 412 if the file has changed since this time.
	// Ignored when If-Match is present. Equivalent to the If-Unmodified-Since
	// header.
	IfUnmodifiedSince string `json:"-" url:"if_unmodified_since,omitempty"`
	// The sandbox resource name, in the form sandboxes/{sandbox_id}.
	Parent string `json:"-" url:"-"`
	// The absolute UTF-8 path of the file inside the sandbox.
	Path string `json:"-" url:"path"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DownloadFileRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DownloadFileRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DownloadFileResponse struct {
	// The supported range unit: bytes.
	AcceptRanges string `json:"-" url:"-" header:"Accept-Ranges,omitempty"`
	// The length of the response body in bytes.
	ContentLength int64 `json:"-" url:"-" header:"Content-Length,omitempty"`
	// The returned byte range and full file length, for a partial response.
	ContentRange string `json:"-" url:"-" header:"Content-Range,omitempty"`
	// The media type of the response body: application/octet-stream.
	ContentType string `json:"-" url:"-" header:"Content-Type,omitempty"`

	Contents io.ReadCloser `json:"-"`
	// The entity tag used for conditional requests.
	Etag string `json:"-" url:"-" header:"etag,omitempty"`
	// The modification time as an HTTP-date, when representable.
	LastModified string `json:"-" url:"-" header:"Last-Modified,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DownloadFileResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DownloadFileResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EnvironmentSpec struct {
	// The Unity Catalog container image to run the sandbox from, in the form
	// `catalog.schema.artifact:tag`, or
	// `catalog.schema.artifact@sha256:<digest>` to pin an exact version. When
	// omitted, the sandbox runs the default sandbox image. See [custom images].
	//
	// [custom images]: https://docs.databricks.com/en/compute/serverless/sandbox-custom-image
	ImageUri string `json:"image_uri,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EnvironmentSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EnvironmentSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Terminal status of a unary command execution.
type ExecuteCommandStatus string

const ExecuteCommandStatusExecuteCommandStatusCompleted ExecuteCommandStatus = `EXECUTE_COMMAND_STATUS_COMPLETED`

const ExecuteCommandStatusExecuteCommandStatusFailed ExecuteCommandStatus = `EXECUTE_COMMAND_STATUS_FAILED`

const ExecuteCommandStatusExecuteCommandStatusTimedOut ExecuteCommandStatus = `EXECUTE_COMMAND_STATUS_TIMED_OUT`

// String representation for [fmt.Print]
func (f *ExecuteCommandStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ExecuteCommandStatus) Set(v string) error {
	switch v {
	case `EXECUTE_COMMAND_STATUS_COMPLETED`, `EXECUTE_COMMAND_STATUS_FAILED`, `EXECUTE_COMMAND_STATUS_TIMED_OUT`:
		*f = ExecuteCommandStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EXECUTE_COMMAND_STATUS_COMPLETED", "EXECUTE_COMMAND_STATUS_FAILED", "EXECUTE_COMMAND_STATUS_TIMED_OUT"`, v)
	}
}

// Values returns all possible values for ExecuteCommandStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *ExecuteCommandStatus) Values() []ExecuteCommandStatus {
	return []ExecuteCommandStatus{
		ExecuteCommandStatusExecuteCommandStatusCompleted,
		ExecuteCommandStatusExecuteCommandStatusFailed,
		ExecuteCommandStatusExecuteCommandStatusTimedOut,
	}
}

// Type always returns ExecuteCommandStatus to satisfy [pflag.Value] interface
func (f *ExecuteCommandStatus) Type() string {
	return "ExecuteCommandStatus"
}

// Request to run a command in the given sandbox and wait for it to finish.
type ExecuteCommandSyncRequest struct {
	// Arguments passed to `cmd`.
	Args []string `json:"args,omitempty"`
	// Executable or command to run (e.g. `/bin/echo`, `python3`).
	Cmd string `json:"cmd"`
	// Extra environment variables for the command's process, merged over the
	// sandbox's default environment.
	Envs map[string]string `json:"envs,omitempty"`
	// Maximum time to wait for the command to finish. When it elapses the
	// command is terminated and the response carries status `TIMED_OUT`. The
	// server applies a default when unset and clamps to an upper bound;
	// negative or otherwise invalid durations are rejected with
	// `INVALID_ARGUMENT`.
	ExecutionTimeout *duration.Duration `json:"execution_timeout,omitempty"`
	// Resource name of the sandbox to run the command in, in the form
	// `sandboxes/{sandbox_id}`. Bound from the URL path.
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExecuteCommandSyncRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExecuteCommandSyncRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Result of a completed unary command execution: captured output, exit code,
// and terminal status.
type ExecuteCommandSyncResponse struct {
	// Daemon-generated identifier for this command execution, for correlation
	// (for example in `ListCommands`).
	CommandId string `json:"command_id,omitempty"`
	// Process exit code. Unset when the process was terminated by a signal
	// (e.g. on `TIMED_OUT`) or never started (`FAILED`) rather than exiting
	// normally.
	ExitCode int `json:"exit_code,omitempty"`
	// Terminal status of the command execution. Always set on a successful
	// response; never `EXECUTE_COMMAND_STATUS_UNSPECIFIED`.
	Status ExecuteCommandStatus `json:"status,omitempty"`
	// Captured standard error, with the same UTF-8 semantics as `stdout`.
	Stderr string `json:"stderr,omitempty"`
	// Captured standard output as UTF-8 text. Invalid UTF-8 bytes are replaced
	// with the Unicode replacement character (U+FFFD).
	Stdout string `json:"stdout,omitempty"`
	// True when `stdout` / `stderr` were truncated because the captured output
	// exceeded the server's per-response size cap. The dropped output is not
	// included in this response and is not recoverable through this unary API.
	Truncated bool `json:"truncated,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExecuteCommandSyncResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExecuteCommandSyncResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetMetadataRequest struct {
	Parent string `json:"-" url:"-"`
	// Absolute UTF-8 path of the entry to inspect.
	Path string `json:"-" url:"path"`
}

func (s *GetMetadataRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Contains metadata for the requested path without following its final symlink.
type GetMetadataResponse struct {
	File *SandboxFileInfo `json:"file,omitempty"`
}

func (s *GetMetadataResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetSandboxRequest struct {
	// Resource name of the sandbox to retrieve, in the form
	// `sandboxes/{sandbox_id}`.
	Name string `json:"-" url:"-"`
}

func (s *GetSandboxRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ListCommandsRequest struct {
	// Maximum number of commands to return. The server may return fewer. If
	// unspecified, the server returns all commands.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Page token returned by a previous ListCommands call. Use this to retrieve
	// the next page of results.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The sandbox whose commands to list, in the form `sandboxes/{sandbox_id}`.
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListCommandsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCommandsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response listing tracked command executions.
type ListCommandsResponse struct {
	// Commands in this page of results.
	Commands []Command `json:"commands,omitempty"`
	// Token to retrieve the next page. Empty when there are no more results.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListCommandsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCommandsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListDirectoryContentsRequest struct {
	// Zero uses the default of 1000; values above 1000 are clamped to 1000.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// The token from a prior response. Keep parent and path unchanged between
	// pages.
	PageToken string `json:"-" url:"page_token,omitempty"`

	Parent string `json:"-" url:"-"`
	// Absolute UTF-8 path of the directory.
	Path string `json:"-" url:"path"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDirectoryContentsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDirectoryContentsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Contains a page of immediate directory entries and an optional continuation
// token. An empty token marks the final page. Pages reflect concurrent
// directory changes.
type ListDirectoryContentsResponse struct {
	Contents []SandboxFileInfo `json:"contents,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDirectoryContentsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDirectoryContentsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSandboxesRequest struct {
	// Maximum number of sandboxes to return in one page. When unset or 0 the
	// server returns up to 1000 sandboxes; values above 5000 are reduced to
	// 5000.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Page token from the `next_page_token` of a previous ListSandboxes
	// response. Omit it to retrieve the first page.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSandboxesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSandboxesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A list of Sandboxes.
type ListSandboxesResponse struct {
	// Token for retrieving the next page of results. Empty when there are no
	// more results.
	NextPageToken string `json:"next_page_token,omitempty"`
	// Sandboxes in this page of results.
	Sandboxes []Sandbox `json:"sandboxes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSandboxesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSandboxesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Network configuration for the sandbox.
type NetworkSpec struct {
	// Whether the sandbox can reach the public internet. Defaults to true.
	// Control-plane, PrivateLink, and workspace/UC/Delta-shared storage stay
	// reachable either way.
	InternetAccess bool `json:"internet_access,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *NetworkSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NetworkSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Identifies the source and exact destination for a sandbox path rename.
type RenamePathRequest struct {
	// Exact absolute UTF-8 destination; it must not already exist.
	Destination string `json:"destination"`
	// Sandbox resource name, for example `sandboxes/my-sandbox`.
	Parent string `json:"-" url:"-"`
	// Absolute UTF-8 path to move.
	Source string `json:"source"`
}

func (s *RenamePathRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Confirms that the sandbox path was renamed.
type RenamePathResponse struct {
}

func (s *RenamePathResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// An isolated, pre-configured, low-latency Serverless compute environment for
// running code. Create a sandbox, wait until its state is
// `SANDBOX_STATE_RUNNING`, then run commands in it with ExecuteCommandSync. The
// data in its home directory persists across stop and start and is removed when
// the sandbox is deleted. See [Databricks Sandbox].
//
// [Databricks Sandbox]: https://docs.databricks.com/en/compute/serverless/sandbox
type Sandbox struct {
	// Output only. The creation time of the sandbox.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Human-readable display label for the sandbox. At most 256 characters.
	DisplayName string `json:"display_name,omitempty"`
	// Tags and other metadata associated with the sandbox.
	Metadata *SandboxMetadata `json:"metadata,omitempty"`
	// The sandbox resource name, in the form `sandboxes/{sandbox_id}`. Derived
	// from `sandbox_id`; any value supplied in a create or update request body
	// is ignored.
	Name string `json:"name,omitempty"`
	// The desired configuration of the sandbox, supplied by the caller at
	// creation time.
	Spec *SandboxSpec `json:"spec,omitempty"`
	// The observed runtime state of the sandbox, populated by the server.
	Status *SandboxStatus `json:"status,omitempty"`
	// Output only. The last update time of the sandbox metadata and spec.
	UpdateTime *time.Time `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Sandbox) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Sandbox) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Metadata for a sandbox filesystem entry, describing a final symlink itself.
type SandboxFileInfo struct {
	FileType SandboxFileType `json:"file_type,omitempty"`

	Group string `json:"group,omitempty"`

	Mode int `json:"mode,omitempty"`
	// Nanosecond precision is preserved on the wire.
	ModifyTime *time.Time `json:"modify_time,omitempty"`

	Name string `json:"name,omitempty"`

	Owner string `json:"owner,omitempty"`

	Path string `json:"path,omitempty"`

	Permissions string `json:"permissions,omitempty"`

	Size int64 `json:"size,omitempty"`

	SymlinkTarget string `json:"symlink_target,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SandboxFileInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SandboxFileInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Type of a sandbox filesystem entry without following its final symlink.
type SandboxFileType string

const SandboxFileTypeSandboxFileTypeDirectory SandboxFileType = `SANDBOX_FILE_TYPE_DIRECTORY`

const SandboxFileTypeSandboxFileTypeFile SandboxFileType = `SANDBOX_FILE_TYPE_FILE`

const SandboxFileTypeSandboxFileTypeOther SandboxFileType = `SANDBOX_FILE_TYPE_OTHER`

const SandboxFileTypeSandboxFileTypeSymlink SandboxFileType = `SANDBOX_FILE_TYPE_SYMLINK`

// String representation for [fmt.Print]
func (f *SandboxFileType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SandboxFileType) Set(v string) error {
	switch v {
	case `SANDBOX_FILE_TYPE_DIRECTORY`, `SANDBOX_FILE_TYPE_FILE`, `SANDBOX_FILE_TYPE_OTHER`, `SANDBOX_FILE_TYPE_SYMLINK`:
		*f = SandboxFileType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "SANDBOX_FILE_TYPE_DIRECTORY", "SANDBOX_FILE_TYPE_FILE", "SANDBOX_FILE_TYPE_OTHER", "SANDBOX_FILE_TYPE_SYMLINK"`, v)
	}
}

// Values returns all possible values for SandboxFileType.
//
// There is no guarantee on the order of the values in the slice.
func (f *SandboxFileType) Values() []SandboxFileType {
	return []SandboxFileType{
		SandboxFileTypeSandboxFileTypeDirectory,
		SandboxFileTypeSandboxFileTypeFile,
		SandboxFileTypeSandboxFileTypeOther,
		SandboxFileTypeSandboxFileTypeSymlink,
	}
}

// Type always returns SandboxFileType to satisfy [pflag.Value] interface
func (f *SandboxFileType) Type() string {
	return "SandboxFileType"
}

// Metadata associated with a sandbox.
type SandboxMetadata struct {
	// Key-value tags associated with the sandbox. Supports up to 54 tags. Keys
	// must be between 1 and 512 characters, and values must not exceed 256
	// characters.
	CustomTags map[string]string `json:"custom_tags,omitempty"`
}

func (s *SandboxMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type SandboxSpec struct {
	// Compute settings requested for the sandbox, such as the inactivity
	// timeout.
	Compute *ComputeSpec `json:"compute,omitempty"`
	// Software environment for the sandbox, such as the container image it
	// runs.
	Environment *EnvironmentSpec `json:"environment,omitempty"`
	// Network configuration requested for the sandbox.
	Network *NetworkSpec `json:"network,omitempty"`
}

func (s *SandboxSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Lifecycle state of a Sandbox resource. STOPPING is the transient state while
// the sandbox is being stopped -- by a Stop request or inactivity
// auto-termination -- and settles to STOPPED once the operation completes.
type SandboxState string

const SandboxStateSandboxStatePending SandboxState = `SANDBOX_STATE_PENDING`

const SandboxStateSandboxStateRunning SandboxState = `SANDBOX_STATE_RUNNING`

const SandboxStateSandboxStateStopped SandboxState = `SANDBOX_STATE_STOPPED`

const SandboxStateSandboxStateStopping SandboxState = `SANDBOX_STATE_STOPPING`

// String representation for [fmt.Print]
func (f *SandboxState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SandboxState) Set(v string) error {
	switch v {
	case `SANDBOX_STATE_PENDING`, `SANDBOX_STATE_RUNNING`, `SANDBOX_STATE_STOPPED`, `SANDBOX_STATE_STOPPING`:
		*f = SandboxState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "SANDBOX_STATE_PENDING", "SANDBOX_STATE_RUNNING", "SANDBOX_STATE_STOPPED", "SANDBOX_STATE_STOPPING"`, v)
	}
}

// Values returns all possible values for SandboxState.
//
// There is no guarantee on the order of the values in the slice.
func (f *SandboxState) Values() []SandboxState {
	return []SandboxState{
		SandboxStateSandboxStatePending,
		SandboxStateSandboxStateRunning,
		SandboxStateSandboxStateStopped,
		SandboxStateSandboxStateStopping,
	}
}

// Type always returns SandboxState to satisfy [pflag.Value] interface
func (f *SandboxState) Type() string {
	return "SandboxState"
}

type SandboxStatus struct {
	// Lifecycle state of the sandbox.
	//
	// - `SANDBOX_STATE_PENDING`: The sandbox is being provisioned after a
	// create or start request and cannot run commands yet. -
	// `SANDBOX_STATE_RUNNING`: The sandbox is ready and commands can run in it.
	// - `SANDBOX_STATE_STOPPING`: The sandbox is shutting down. It settles to
	// `SANDBOX_STATE_STOPPED`, or disappears if it is being deleted. -
	// `SANDBOX_STATE_STOPPED`: The sandbox is stopped, by a stop request,
	// inactivity auto-stop, or a failed provisioning. Its home directory is
	// kept and it can be started again with StartSandbox.
	State SandboxState `json:"state,omitempty"`
}

func (s *SandboxStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A request to start a Sandbox.
type StartSandboxRequest struct {
	// Resource name of the sandbox to start, in the form
	// `sandboxes/{sandbox_id}`.
	Name string `json:"-" url:"-"`
}

func (s *StartSandboxRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A request to stop a Sandbox.
type StopSandboxRequest struct {
	// Resource name of the sandbox to stop, in the form
	// `sandboxes/{sandbox_id}`.
	Name string `json:"-" url:"-"`
}

func (s *StopSandboxRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateSandboxRequest struct {
	// Resource name of the sandbox to update, in the form
	// `sandboxes/{sandbox_id}`.
	Name string `json:"-" url:"-"`
	// The Sandbox resource carrying new field values. Only fields named in
	// `update_mask` are read; unmasked fields are ignored.
	Sandbox Sandbox `json:"sandbox"`
	// Field paths to update. Must be a non-empty subset of: - display_name -
	// spec.compute.inactivity_timeout Any other path returns
	// INVALID_PARAMETER_VALUE.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateSandboxRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UploadFileRequest struct {
	// The complete file contents. An empty body creates an empty file.
	Contents io.ReadCloser `json:"-"`
	// Whether to replace an existing file. Defaults to true when omitted. If
	// false, the upload atomically creates the file only if the destination
	// does not exist, equivalent to If-None-Match: *. An existing entry returns
	// 412 Precondition Failed. Explicit true conflicts with If-None-Match: *.
	Overwrite bool `json:"-" url:"overwrite,omitempty"`
	// The sandbox resource name, in the form sandboxes/{sandbox_id}.
	Parent string `json:"-" url:"-"`
	// The absolute UTF-8 destination path inside the sandbox.
	Path string `json:"-" url:"path"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UploadFileRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UploadFileRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
