// Package terraformtest runs hermetic Terraform lifecycles from YAML fixtures.
package terraformtest

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/databricks/terraform-provider-databricks/qa/httpfixture"
)

type testDefinition struct {
	Version      int            `yaml:"version"`
	HostMetadata map[string]any `yaml:"host_metadata,omitempty"`
	Steps        []step         `yaml:"steps"`
	Destroy      destroy        `yaml:"destroy"`
}

type step struct {
	Name      string             `yaml:"name"`
	Terraform string             `yaml:"terraform"`
	Calls     []httpfixture.Call `yaml:"calls"`
	State     *expectedState     `yaml:"state,omitempty"`
}

type destroy struct {
	Calls []httpfixture.Call `yaml:"calls"`
}

type expectedState struct {
	Resources map[string]expectedResource `yaml:"resources,omitempty"`
}

type expectedResource struct {
	Attributes map[string]any `yaml:"attributes"`
}

func loadDefinition(path string) (*testDefinition, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return decodeDefinition(file)
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
	for i, s := range d.Steps {
		label := stepLabel(i, s.Name)
		if strings.TrimSpace(s.Name) == "" || strings.TrimSpace(s.Terraform) == "" {
			return fmt.Errorf("%s: name and terraform configuration are required", label)
		}
		if err := validateState(s.State); err != nil {
			return fmt.Errorf("%s: %w", label, err)
		}
	}
	return d.validateCalls()
}

func validateState(state *expectedState) error {
	if state == nil {
		return nil
	}
	for address, resource := range state.Resources {
		if address == "" {
			return fmt.Errorf("state resource address must not be empty")
		}
		if _, err := json.Marshal(resource.Attributes); err != nil {
			return fmt.Errorf("%s: attributes must contain JSON-compatible values: %w", address, err)
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
