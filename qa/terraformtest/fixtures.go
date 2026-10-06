package terraformtest

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/databricks/databricks-sdk-go/common/environment"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/terraform-provider-databricks/qa"
)

const (
	fixtureFormatVersion = 1
	destroyStepName      = "destroy"
)

type fixtureExchange struct {
	Version   int             `json:"version"`
	StepIndex *int            `json:"step_index"`
	StepName  string          `json:"step_name"`
	Sequence  *int            `json:"sequence"`
	Request   fixtureRequest  `json:"request"`
	Response  fixtureResponse `json:"response"`
}

type fixtureRequest struct {
	Method  string            `json:"method"`
	Path    string            `json:"path"`
	Headers map[string]string `json:"headers"`
	Body    json.RawMessage   `json:"body"`
}

type fixtureResponse struct {
	Status  int               `json:"status"`
	Headers map[string]string `json:"headers"`
	Body    json.RawMessage   `json:"body"`
}

// WorkspaceHostMetadata returns the standard discovery response for a mocked workspace host.
func WorkspaceHostMetadata(workspaceID string) *config.HostMetadata {
	return &config.HostMetadata{
		WorkspaceID: workspaceID,
		Cloud:       environment.CloudAWS,
		HostType:    config.WorkspaceHost,
	}
}

// withHostMetadataFixtures prepends the test-level discovery response to every fixture phase.
// Terraform configures fresh provider instances for plan and apply commands, so the fixture is
// reusable within each phase. Copying the response also lets the local server supply a valid OIDC
// endpoint without modifying the TestCase value.
func withHostMetadataFixtures(
	fixturePhases [][]qa.HTTPFixture,
	metadata *config.HostMetadata,
	defaultOIDCEndpoint string,
) [][]qa.HTTPFixture {
	if metadata == nil {
		return fixturePhases
	}
	metadataCopy := *metadata
	if metadataCopy.OIDCEndpoint == "" {
		metadataCopy.OIDCEndpoint = defaultOIDCEndpoint
	}
	metadataFixture := qa.HTTPFixture{
		Method:       http.MethodGet,
		Resource:     "/.well-known/databricks-config",
		Status:       http.StatusOK,
		Response:     metadataCopy,
		ReuseRequest: true,
	}
	result := make([][]qa.HTTPFixture, len(fixturePhases))
	for phaseIndex, fixtures := range fixturePhases {
		result[phaseIndex] = append([]qa.HTTPFixture{metadataFixture}, fixtures...)
	}
	return result
}

func loadFixtureFile(path string, steps []Step) ([][]qa.HTTPFixture, error) {
	stepNames, err := fixtureStepNames(steps)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fixtures := make([][]qa.HTTPFixture, len(stepNames))
	expectedSequences := make([]int, len(stepNames))
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 64*1024), 10*1024*1024)
	for lineNumber := 1; scanner.Scan(); lineNumber++ {
		line := bytes.TrimSpace(scanner.Bytes())
		if len(line) == 0 {
			continue
		}
		exchange, err := decodeFixtureExchange(line)
		if err != nil {
			return nil, fmt.Errorf("%s:%d: %w", path, lineNumber, err)
		}
		if err := validateFixtureExchange(exchange, stepNames, expectedSequences); err != nil {
			return nil, fmt.Errorf("%s:%d: %w", path, lineNumber, err)
		}

		stepIndex := *exchange.StepIndex
		fixtures[stepIndex] = append(fixtures[stepIndex], qa.HTTPFixture{
			Method:          exchange.Request.Method,
			Resource:        exchange.Request.Path,
			ExpectedHeaders: exchange.Request.Headers,
			ExpectedRequest: fixtureBody(exchange.Request.Body),
			Status:          exchange.Response.Status,
			ResponseHeaders: exchange.Response.Headers,
			Response:        fixtureBody(exchange.Response.Body),
		})
		expectedSequences[stepIndex]++
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}
	return fixtures, nil
}

func fixtureStepNames(steps []Step) ([]string, error) {
	if len(steps) == 0 {
		return nil, nil
	}
	names := make([]string, 0, len(steps)+1)
	seen := make(map[string]struct{}, len(steps)+1)
	for stepIndex, step := range steps {
		if step.Name == "" {
			return nil, fmt.Errorf("step %d has no name", stepIndex)
		}
		if step.Name == destroyStepName {
			return nil, fmt.Errorf("step %d uses reserved name %q", stepIndex, destroyStepName)
		}
		if _, ok := seen[step.Name]; ok {
			return nil, fmt.Errorf("step %d repeats name %q", stepIndex, step.Name)
		}
		seen[step.Name] = struct{}{}
		names = append(names, step.Name)
	}
	return append(names, destroyStepName), nil
}

func decodeFixtureExchange(line []byte) (fixtureExchange, error) {
	var exchange fixtureExchange
	decoder := json.NewDecoder(bytes.NewReader(line))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&exchange); err != nil {
		return fixtureExchange{}, err
	}
	var trailing any
	if err := decoder.Decode(&trailing); err != io.EOF {
		if err == nil {
			return fixtureExchange{}, fmt.Errorf("multiple JSON values on one line")
		}
		return fixtureExchange{}, err
	}
	return exchange, nil
}

func validateFixtureExchange(
	exchange fixtureExchange,
	stepNames []string,
	expectedSequences []int,
) error {
	if exchange.Version != fixtureFormatVersion {
		return fmt.Errorf("unsupported fixture version %d", exchange.Version)
	}
	if exchange.StepIndex == nil {
		return fmt.Errorf("missing step_index")
	}
	stepIndex := *exchange.StepIndex
	if stepIndex < 0 || stepIndex >= len(stepNames) {
		return fmt.Errorf("step_index %d is outside [0, %d)", stepIndex, len(stepNames))
	}
	if exchange.StepName != stepNames[stepIndex] {
		return fmt.Errorf(
			"step_name %q does not match step %d name %q",
			exchange.StepName,
			stepIndex,
			stepNames[stepIndex],
		)
	}
	if exchange.Sequence == nil {
		return fmt.Errorf("missing sequence")
	}
	if *exchange.Sequence != expectedSequences[stepIndex] {
		return fmt.Errorf(
			"step %q has sequence %d; expected %d",
			exchange.StepName,
			*exchange.Sequence,
			expectedSequences[stepIndex],
		)
	}
	if exchange.Request.Method == "" {
		return fmt.Errorf("missing request method")
	}
	if exchange.Request.Path == "" {
		return fmt.Errorf("missing request path")
	}
	if exchange.Response.Status < 100 || exchange.Response.Status > 599 {
		return fmt.Errorf("invalid response status %d", exchange.Response.Status)
	}
	return nil
}

func fixtureBody(body json.RawMessage) any {
	body = bytes.TrimSpace(body)
	if len(body) == 0 || bytes.Equal(body, []byte("null")) {
		return nil
	}
	return append(json.RawMessage(nil), body...)
}
