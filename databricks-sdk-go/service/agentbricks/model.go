// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package agentbricks

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type CancelCustomLlmOptimizationRunRequest struct {
	Id string `json:"-" url:"-"`
}

func (s *CancelCustomLlmOptimizationRunRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateCustomLlmRequest struct {
	// This will soon be deprecated!! Optional: UC path for agent artifacts. If
	// you are using a dataset that you only have read permissions, please
	// provide a destination path where you have write permissions. Please
	// provide this in catalog.schema format.
	AgentArtifactPath string `json:"agent_artifact_path,omitempty"`
	// Datasets used for training and evaluating the model, not for inference.
	// Currently, only 1 dataset is accepted.
	Datasets []Dataset `json:"datasets,omitempty"`
	// Guidelines for the custom LLM to adhere to
	Guidelines []string `json:"guidelines,omitempty"`
	// Instructions for the custom LLM to follow
	Instructions string `json:"instructions"`
	// Name of the custom LLM. Only alphanumeric characters and dashes allowed.
	Name string `json:"name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateCustomLlmRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateCustomLlmRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CustomLlm struct {
	AgentArtifactPath string `json:"agent_artifact_path,omitempty"`
	// Creation timestamp of the custom LLM
	CreationTime string `json:"creation_time,omitempty"`
	// Creator of the custom LLM
	Creator string `json:"creator,omitempty"`
	// Datasets used for training and evaluating the model, not for inference
	Datasets []Dataset `json:"datasets,omitempty"`
	// Name of the endpoint that will be used to serve the custom LLM
	EndpointName string `json:"endpoint_name,omitempty"`
	// Guidelines for the custom LLM to adhere to
	Guidelines []string `json:"guidelines,omitempty"`

	Id string `json:"id,omitempty"`
	// Instructions for the custom LLM to follow
	Instructions string `json:"instructions"`
	// Name of the custom LLM
	Name string `json:"name"`
	// If optimization is kicked off, tracks the state of the custom LLM
	OptimizationState State `json:"optimization_state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CustomLlm) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CustomLlm) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Dataset struct {
	Table Table `json:"table"`
}

func (s *Dataset) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteCustomLlmRequest struct {
	// The id of the custom llm
	Id string `json:"-" url:"-"`
}

func (s *DeleteCustomLlmRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetCustomLlmRequest struct {
	// The id of the custom llm
	Id string `json:"-" url:"-"`
}

func (s *GetCustomLlmRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type StartCustomLlmOptimizationRunRequest struct {
	// The Id of the tile.
	Id string `json:"-" url:"-"`
}

func (s *StartCustomLlmOptimizationRunRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// States of Custom LLM optimization lifecycle.
type State string

const StateCancelled State = `CANCELLED`

const StateCompleted State = `COMPLETED`

const StateCreated State = `CREATED`

const StateFailed State = `FAILED`

const StatePending State = `PENDING`

const StateRunning State = `RUNNING`

// String representation for [fmt.Print]
func (f *State) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *State) Set(v string) error {
	switch v {
	case `CANCELLED`, `COMPLETED`, `CREATED`, `FAILED`, `PENDING`, `RUNNING`:
		*f = State(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELLED", "COMPLETED", "CREATED", "FAILED", "PENDING", "RUNNING"`, v)
	}
}

// Values returns all possible values for State.
//
// There is no guarantee on the order of the values in the slice.
func (f *State) Values() []State {
	return []State{
		StateCancelled,
		StateCompleted,
		StateCreated,
		StateFailed,
		StatePending,
		StateRunning,
	}
}

// Type always returns State to satisfy [pflag.Value] interface
func (f *State) Type() string {
	return "State"
}

type Table struct {
	// Name of the request column
	RequestCol string `json:"request_col"`
	// Optional: Name of the response column if the data is labeled
	ResponseCol string `json:"response_col,omitempty"`
	// Full UC table path in catalog.schema.table_name format
	TablePath string `json:"table_path"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Table) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Table) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateCustomLlmRequest struct {
	// The CustomLlm containing the fields which should be updated.
	CustomLlm CustomLlm `json:"custom_llm"`
	// The id of the custom llm
	Id string `json:"-" url:"-"`
	// The list of the CustomLlm fields to update. These should correspond to
	// the values (or lack thereof) present in `custom_llm`.
	//
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	UpdateMask string `json:"update_mask"`
}

func (s *UpdateCustomLlmRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}
