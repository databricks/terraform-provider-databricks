// Package terraformtest runs hermetic Terraform lifecycles from YAML fixtures.
package terraformtest

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/databricks/terraform-provider-databricks/qa/httpfixture"
)

type testDefinition struct {
	Version          int            `yaml:"version"`
	HostMetadata     map[string]any `yaml:"host_metadata,omitempty"`
	InitialState     *initialState  `yaml:"initial_state,omitempty"`
	EnforceCallOrder bool           `yaml:"enforce_call_order,omitempty"`
	Steps            []step         `yaml:"steps"`
	Destroy          destroy        `yaml:"destroy"`
	checkpoint       []byte
}

type initialState struct {
	Checkpoint string `yaml:"checkpoint"`
}

type step struct {
	Name        string             `yaml:"name"`
	Config      string             `yaml:"config"`
	Calls       []httpfixture.Call `yaml:"calls"`
	State       *expectedState     `yaml:"state,omitempty"`
	ExpectError *string            `yaml:"expect_error,omitempty"`
	Plan        *expectedPlan      `yaml:"plan,omitempty"`
	PlanOnly    bool               `yaml:"plan_only,omitempty"`
}

type destroy struct {
	Calls []httpfixture.Call `yaml:"calls"`
}

type expectedState struct {
	Resources map[string]expectedResource `yaml:"resources,omitempty"`
}

type expectedResource struct {
	Attributes []expectedAttribute `yaml:"attributes"`
}

type expectedAttribute struct {
	Path []any `yaml:"path"`
	// A node distinguishes explicit null from an accidentally omitted expectation.
	Equals yaml.Node `yaml:"equals,omitempty"`
}

func loadDefinition(path string) (*testDefinition, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	definition, err := decodeDefinition(file)
	if err != nil || definition.InitialState == nil {
		return definition, err
	}
	path = filepath.Join(filepath.Dir(path), definition.InitialState.Checkpoint)
	definition.checkpoint, err = os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("initial_state checkpoint %s: %w", path, err)
	}
	if len(strings.TrimSpace(string(definition.checkpoint))) == 0 {
		return nil, fmt.Errorf("initial_state checkpoint %s must not be empty", path)
	}
	return definition, nil
}

// decodeDefinition rejects misspelled or unsupported fields before Terraform
// runs; accepting them silently could skip the assertions a test intended.
func decodeDefinition(reader io.Reader) (*testDefinition, error) {
	decoder := yaml.NewDecoder(reader)
	decoder.KnownFields(true)
	var definition testDefinition
	if err := decoder.Decode(&definition); err != nil {
		return nil, fmt.Errorf("invalid or unsupported YAML field: %w", err)
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		return nil, fmt.Errorf("expected a single YAML document")
	}
	if err := definition.validate(); err != nil {
		return nil, err
	}
	return &definition, nil
}

func (d *testDefinition) validate() error {
	if d.Version != 1 {
		return fmt.Errorf("unsupported version %d; expected 1", d.Version)
	}
	if len(d.Steps) == 0 {
		return fmt.Errorf("at least one step is required")
	}
	if d.InitialState != nil {
		if strings.TrimSpace(d.InitialState.Checkpoint) == "" || filepath.IsAbs(d.InitialState.Checkpoint) {
			return fmt.Errorf("initial_state.checkpoint must be a nonempty path relative to the YAML file")
		}
	}
	for i, s := range d.Steps {
		label := stepLabel(i, s.Name)
		if strings.TrimSpace(s.Name) == "" || strings.TrimSpace(s.Config) == "" {
			return fmt.Errorf("%s: name and config are required", label)
		}
		if err := validateState(s.State); err != nil {
			return fmt.Errorf("%s: %w", label, err)
		}
		if err := s.validateExpectedError(); err != nil {
			return fmt.Errorf("%s: %w", label, err)
		}
		if err := s.validatePlan(); err != nil {
			return fmt.Errorf("%s: %w", label, err)
		}
	}
	return d.validateCalls()
}

// validateExpectedError prevents broad empty patterns and assertions that the
// testing framework would skip after an expected failure.
func (s step) validateExpectedError() error {
	if s.ExpectError == nil {
		return nil
	}
	if strings.TrimSpace(*s.ExpectError) == "" {
		return fmt.Errorf("expect_error must be a nonempty regular expression")
	}
	if _, err := regexp.Compile(*s.ExpectError); err != nil {
		return fmt.Errorf("invalid expect_error regular expression: %w", err)
	}
	if s.State != nil {
		return fmt.Errorf("state checks on expected-error steps are not supported; check state in a subsequent successful step")
	}
	return nil
}

func validateState(state *expectedState) error {
	if state == nil {
		return nil
	}
	for address, resource := range state.Resources {
		if address == "" {
			return fmt.Errorf("state resource address must not be empty")
		}
		for i, attribute := range resource.Attributes {
			if _, err := attribute.attributePath(); err != nil {
				return fmt.Errorf("%s: attribute check %d: %w", address, i+1, err)
			}
			if _, err := attribute.valueCheck(); err != nil {
				return fmt.Errorf("%s: attribute check %d: %w", address, i+1, err)
			}
		}
	}
	return nil
}

// validateCalls checks every fixture set up front, including later steps and
// destroy, so an invalid cleanup fixture cannot interrupt a running lifecycle.
func (d *testDefinition) validateCalls() error {
	metadataHandler := httpfixture.NewHandler()
	if err := metadataHandler.SetReusableFixtures(d.metadataCalls()); err != nil {
		return err
	}
	for i, s := range d.Steps {
		// Independent handlers validate the fixtures without requiring consumption.
		handler := httpfixture.NewHandler()
		if err := handler.StartStep(stepLabel(i, s.Name), s.Calls); err != nil {
			return err
		}
	}
	return httpfixture.NewHandler().StartStep("destroy", d.Destroy.Calls)
}

// metadataCalls keeps host discovery reusable because Terraform configures
// fresh provider clients multiple times within a single step.
func (d *testDefinition) metadataCalls() []httpfixture.Call {
	if d.HostMetadata == nil {
		return nil
	}
	return []httpfixture.Call{{
		Request:  httpfixture.Request{Method: "GET", Path: "/.well-known/databricks-config"},
		Response: httpfixture.Response{Status: 200, Body: &httpfixture.Body{JSON: d.HostMetadata}},
	}}
}

func stepLabel(index int, name string) string {
	return fmt.Sprintf("step %d (%s)", index+1, name)
}
