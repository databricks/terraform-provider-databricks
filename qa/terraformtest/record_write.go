package terraformtest

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/databricks/terraform-provider-databricks/qa/httpfixture"
)

// Refresh only captured data. The document tree retains HCL, comments, checks,
// ordering choices and unsupported future recording settings verbatim in value.
func refreshRecording(original []byte, d *testDefinition, captured []httpfixture.CapturedStep, n *recordingNormalizer) ([]byte, error) {
	if len(captured) != len(d.Steps)+1 || captured[len(captured)-1].Label != stepLabel(len(d.Steps), "destroy") {
		return nil, fmt.Errorf("recording did not complete all steps and mandatory destroy; YAML unchanged")
	}
	for _, phase := range captured {
		redactSecretReads(phase.Calls)
	}
	encoded, err := yaml.Marshal(captured)
	if err != nil {
		return nil, err
	}
	encoded, err = n.transform(encoded, false, true)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(encoded, &captured); err != nil {
		return nil, err
	}
	var document yaml.Node
	if err := yaml.Unmarshal(original, &document); err != nil {
		return nil, err
	}
	root := document.Content[0]
	metadata, err := refreshStepCalls(root, d, captured)
	if err != nil {
		return nil, err
	}
	if metadata != nil {
		if err := setYAMLField(root, "host_metadata", metadata); err != nil {
			return nil, err
		}
	}
	var out bytes.Buffer
	encoder := yaml.NewEncoder(&out)
	encoder.SetIndent(2)
	if err := encoder.Encode(&document); err != nil {
		return nil, err
	}
	if _, err := decodeDefinition(bytes.NewReader(out.Bytes())); err != nil {
		return nil, fmt.Errorf("refreshed recording is invalid: %w", err)
	}
	return out.Bytes(), nil
}

// Secrets/get uses a generic "value" field rather than a secret-labelled key.
// Keep its wire encoding valid so provider decoding still runs during replay.
func redactSecretReads(calls []httpfixture.Call) {
	for _, call := range calls {
		if !strings.HasSuffix(call.Request.Path, "/secrets/get") || call.Response.Body == nil {
			continue
		}
		if value, ok := call.Response.Body.JSON.(map[string]any); ok {
			if _, exists := value["value"]; exists {
				value["value"] = base64.StdEncoding.EncodeToString([]byte("[REDACTED]"))
			}
		}
	}
}

func refreshStepCalls(root *yaml.Node, d *testDefinition, captured []httpfixture.CapturedStep) (map[string]any, error) {
	steps := yamlField(root, "steps")
	var metadata map[string]any
	for i, phase := range captured {
		label := stepLabel(i, "destroy")
		previous, target := d.Destroy.Calls, yamlField(root, "destroy")
		if i < len(d.Steps) {
			label, previous, target = stepLabel(i, d.Steps[i].Name), d.Steps[i].Calls, steps.Content[i]
		}
		if phase.Label != label {
			return nil, fmt.Errorf("recorded step %d does not match the YAML step name", i+1)
		}
		calls, err := separateMetadata(phase.Calls, &metadata)
		if err != nil {
			return nil, err
		}
		if err := httpfixture.PreserveMatchers(calls, previous); err != nil {
			return nil, fmt.Errorf("%s: %w", label, err)
		}
		if target == nil {
			if err := setYAMLField(root, "destroy", map[string]any{}); err != nil {
				return nil, err
			}
			target = yamlField(root, "destroy")
		}
		if err := setYAMLField(target, "calls", calls); err != nil {
			return nil, err
		}
	}
	return metadata, nil
}

func separateMetadata(calls []httpfixture.Call, metadata *map[string]any) ([]httpfixture.Call, error) {
	result := []httpfixture.Call{}
	for _, call := range calls {
		if call.Request.Method == "GET" && call.Request.Path == "/.well-known/databricks-config" &&
			call.Response.Status == 200 && call.Response.Body != nil {
			if value, ok := call.Response.Body.JSON.(map[string]any); ok {
				if *metadata != nil && !reflect.DeepEqual(*metadata, value) {
					return nil, fmt.Errorf("host metadata changed during recording; cannot make it reusable")
				}
				*metadata = value
				continue
			}
		}
		result = append(result, call)
	}
	return result, nil
}

func yamlField(node *yaml.Node, key string) *yaml.Node {
	if node == nil {
		return nil
	}
	if node.Kind == yaml.AliasNode {
		return yamlField(node.Alias, key)
	}
	for i := 0; i+1 < len(node.Content); i += 2 {
		if node.Content[i].Value == key {
			return node.Content[i+1]
		}
	}
	return nil
}

func setYAMLField(node *yaml.Node, key string, value any) error {
	var replacement yaml.Node
	if err := replacement.Encode(value); err != nil {
		return err
	}
	if node.Kind == yaml.AliasNode {
		// A refreshed phase must not overwrite another phase through an anchor.
		var resolved any
		if err := node.Decode(&resolved); err != nil {
			return err
		}
		if err := node.Encode(resolved); err != nil {
			return err
		}
	}
	if old := yamlField(node, key); old != nil {
		replacement.HeadComment, replacement.LineComment, replacement.FootComment = old.HeadComment, old.LineComment, old.FootComment
		// Existing aliases may refer to this node. Keep the node's identity and
		// anchor while replacing its contents, rather than leaving dangling aliases.
		replacement.Anchor = old.Anchor
		*old = replacement
	} else {
		node.Content = append(node.Content, &yaml.Node{Kind: yaml.ScalarNode, Tag: "!!str", Value: key}, &replacement)
	}
	return nil
}

// Use a same-directory rename so readers see either complete version. Recheck
// the source before replacing it to avoid overwriting edits made during a run.
func replaceRecording(path string, original, updated []byte) error {
	current, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if !bytes.Equal(current, original) {
		return fmt.Errorf("YAML changed while recording; refusing to overwrite %s", path)
	}
	file, err := os.CreateTemp(filepath.Dir(path), ".recording-*")
	if err != nil {
		return err
	}
	defer os.Remove(file.Name())
	defer file.Close()
	if _, err := file.Write(updated); err != nil {
		return err
	}
	if err := file.Sync(); err != nil {
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}
	return os.Rename(file.Name(), path)
}
